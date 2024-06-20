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
        if False:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wvefrkb"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwssl")))
        elif True:
            return False

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "domj"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_0_i0_ in range(default__.abs(367))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vg"))))

    @staticmethod
    def fm2(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, True, True, False])) + (_dafny.SeqWithoutIsStrInference([not(False)]))) + ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([not(True)])))

    @staticmethod
    def fm3(p0, p1, globalState):
        return ((0) - (len(_dafny.SeqWithoutIsStrInference([True])))) * ((len(_dafny.SeqWithoutIsStrInference([not(False), not(False)]))) - (598))

    @staticmethod
    def fm15(p0, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference([519 for d_1_i0_ in range(default__.abs(882))]): True})

    @staticmethod
    def fm16(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-357, 330):
                d_2_v0_: int = compr_0_
                if ((-357) <= (d_2_v0_)) and ((d_2_v0_) < (330)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_2_v0_, len(_dafny.SeqWithoutIsStrInference([False])))]))
            return _dafny.Set(coll0_)
        return ((_dafny.Set({892, (0) - (len(_dafny.Set({False}))), len(_dafny.Set({353, -515}))}) if True else _dafny.Set({len(_dafny.Set({-214}))}))) | ((iife0_()
        ) - (_dafny.Set({len(_dafny.Map({799: True})), len(_dafny.Map({52: True}))})))

    @staticmethod
    def fm17(p0, p1, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: _dafny.Map
            for compr_1_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prkc")))})]))).Elements:
                d_3_v0_: _dafny.Map = compr_1_
                if (d_3_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prkc")))})]))):
                    coll1_ = coll1_.union(_dafny.Set([d_3_v0_]))
            return _dafny.Set(coll1_)
        source0_ = D1_DC4(len(iife1_()
))
        if source0_.is_DC5:
            d_4___mcc_h0_ = source0_.cf5
            d_5_cf5_ = d_4___mcc_h0_
            return _dafny.Set({d_5_cf5_})
        elif True:
            d_6___mcc_h1_ = source0_.cf4
            d_7_cf4_ = d_6___mcc_h1_
            return _dafny.Set({True})

    @staticmethod
    def fm18(p0, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({946: len(_dafny.SeqWithoutIsStrInference([461]))}))])).cardinality: len(_dafny.Set({True}))}): True})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "begfq"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcfe"))), -714])): False})

    @staticmethod
    def fm19(p0, p1, globalState):
        return (_dafny.Map({True: False})) | (_dafny.Map({True: not(False)}))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        source1_ = D2_DC6(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_8_i0_ in range(default__.abs(110))]))
        if source1_.is_DC7:
            d_9___mcc_h0_ = source1_.cf7
            d_10___mcc_h1_ = source1_.cf8
            d_11_cf8_ = d_10___mcc_h1_
            d_12_cf7_ = d_9___mcc_h0_
            return _dafny.CodePoint('y')
        elif source1_.is_DC8:
            d_13___mcc_h2_ = source1_.cf9
            d_14_cf9_ = d_13___mcc_h2_
            return _dafny.CodePoint('i')
        elif source1_.is_DC6:
            d_15___mcc_h3_ = source1_.cf6
            d_16_cf6_ = d_15___mcc_h3_
            return _dafny.CodePoint('o')
        elif True:
            d_17___mcc_h4_ = source1_.cf10
            d_18_cf10_ = d_17___mcc_h4_
            if True:
                return _dafny.CodePoint('e')
            elif True:
                return _dafny.CodePoint('m')

    @staticmethod
    def fm23(globalState):
        return D0_DC1()

    @staticmethod
    def fm24(globalState):
        if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shj"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfuu"))):
            return D0_DC0(False)
        elif True:
            return D0_DC0(not(not(False)))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        if (False) not in (_dafny.Map({True: not(True)})):
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(-970, -460):
                    d_19_v0_: int = compr_2_
                    if ((-970) <= (d_19_v0_)) and ((d_19_v0_) < (-460)):
                        coll2_[(d_19_v0_) * (802)] = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thvvevq")))), -477, 803, (_dafny.MultiSet([False])).cardinality])
                return _dafny.Map(coll2_)
            return iife2_()
            
        elif True:
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(488, 518):
                    d_20_v1_: int = compr_3_
                    if ((488) <= (d_20_v1_)) and ((d_20_v1_) < (518)):
                        coll3_[default__.safeDivisionInt(d_20_v1_, 567)] = _dafny.SeqWithoutIsStrInference([-734 for d_21_i0_ in range(default__.abs(-215))])
                return _dafny.Map(coll3_)
            return iife3_()
            

    @staticmethod
    def fm26(globalState):
        return D1_DC5((_dafny.CodePoint('f')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ero"))))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        def iife4_():
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(14, 706):
                    d_24_v1_: int = compr_5_
                    if ((14) <= (d_24_v1_)) and ((d_24_v1_) < (706)):
                        coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_24_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_25_i1_ in range(default__.abs(-948))])))]))
                return _dafny.Set(coll5_)
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.SeqWithoutIsStrInference([-977 for d_22_i0_ in range(default__.abs(236))])).Elements:
                d_23_v0_: int = compr_4_
                if (d_23_v0_) in (_dafny.SeqWithoutIsStrInference([-977 for d_22_i0_ in range(default__.abs(236))])):
                    coll4_[default__.safeDivisionInt(d_23_v0_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpiay"))), len(_dafny.SeqWithoutIsStrInference([len(iife5_()
                    ), 659])), 310})))] = len(_dafny.Set({False, False, not(True)}))
            return _dafny.Map(coll4_)
        return (iife4_()
        ) | (_dafny.Map({-232: 13}))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        source2_ = D21_DC46(True, False, 966)
        if source2_.is_DC46:
            d_26___mcc_h0_ = source2_.cf61
            d_27___mcc_h1_ = source2_.cf62
            d_28___mcc_h2_ = source2_.cf63
            d_29_cf63_ = d_28___mcc_h2_
            d_30_cf62_ = d_27___mcc_h1_
            d_31_cf61_ = d_26___mcc_h0_
            return _dafny.Map({not(True): d_29_cf63_})
        elif source2_.is_DC47:
            d_32___mcc_h3_ = source2_.cf64
            d_33_cf64_ = d_32___mcc_h3_
            return _dafny.Map({d_33_cf64_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgerly")))})
        elif True:
            d_34___mcc_h4_ = source2_.cf60
            d_35_cf60_ = d_34___mcc_h4_
            return _dafny.Map({False: 872})

    @staticmethod
    def fm31(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfhxej")))}))])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality, -270])).cardinality]))) + ((_dafny.SeqWithoutIsStrInference([-916, 191])) + (_dafny.SeqWithoutIsStrInference([401, 4, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwex")))])))

    @staticmethod
    def fm32(globalState):
        if False:
            return (_dafny.MultiSet([-763])) - (_dafny.MultiSet([len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slnptiny"))})), 872, (_dafny.MultiSet([467])).cardinality]))
        elif True:
            return _dafny.MultiSet([(_dafny.MultiSet([_dafny.CodePoint('u'), _dafny.CodePoint('v')])).cardinality, 601])

    @staticmethod
    def fm33(globalState):
        source3_ = D12_DC21(_dafny.Map({False: False}))
        if source3_.is_DC22:
            d_36___mcc_h0_ = source3_.cf24
            d_37___mcc_h1_ = source3_.cf25
            d_38___mcc_h2_ = source3_.cf26
            d_39___mcc_h3_ = source3_.cf27
            d_40___mcc_h4_ = source3_.cf28
            d_41_cf28_ = d_40___mcc_h4_
            d_42_cf27_ = d_39___mcc_h3_
            d_43_cf26_ = d_38___mcc_h2_
            d_44_cf25_ = d_37___mcc_h1_
            d_45_cf24_ = d_36___mcc_h0_
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in (_dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([934])), len(_dafny.Set({d_42_cf27_})), -580, len(d_41_cf28_), (0) - (-570)}))])).Elements:
                    d_46_v0_: int = compr_6_
                    if (d_46_v0_) in (_dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([934])), len(_dafny.Set({d_42_cf27_})), -580, len(d_41_cf28_), (0) - (-570)}))])):
                        coll6_[default__.safeDivisionInt(d_46_v0_, -820)] = d_42_cf27_
                return _dafny.Map(coll6_)
            return (iife6_()
            ) | (_dafny.Map({len(d_41_cf28_): False}))
        elif source3_.is_DC23:
            d_47___mcc_h5_ = source3_.cf29
            d_48_cf29_ = d_47___mcc_h5_
            return _dafny.Map({-711: True})
        elif True:
            d_49___mcc_h6_ = source3_.cf23
            d_50_cf23_ = d_49___mcc_h6_
            return _dafny.Map({34: True})

    @staticmethod
    def fm34(p0, globalState):
        return (_dafny.MultiSet([False])).intersection(_dafny.MultiSet([False]))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(379, -580):
                d_51_v0_: int = compr_7_
                if ((379) <= (d_51_v0_)) and ((d_51_v0_) < (-580)):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_51_v0_, 624)]))
            return _dafny.Set(coll7_)
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(-9, 42):
                d_52_v1_: int = compr_8_
                if ((-9) <= (d_52_v1_)) and ((d_52_v1_) < (42)):
                    coll8_ = coll8_.union(_dafny.Set([default__.safeModuloInt(d_52_v1_, -429)]))
            return _dafny.Set(coll8_)
        return (iife7_()
        ).intersection(iife8_()
        )

    @staticmethod
    def fm36(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('h')

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        if False:
            return _dafny.SeqWithoutIsStrInference([-517, 898, len(_dafny.SeqWithoutIsStrInference([not(True)]))])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([668])) + (_dafny.SeqWithoutIsStrInference([452]))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-299]), _dafny.SeqWithoutIsStrInference([-751]), _dafny.SeqWithoutIsStrInference([854, (0) - (len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwgq")))): False})))]), _dafny.SeqWithoutIsStrInference([356])])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([493, (_dafny.MultiSet([not(not(False))])).cardinality])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-902])])))

    @staticmethod
    def fm39(p0, p1, p2, p3, globalState):
        if False:
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return _dafny.SeqWithoutIsStrInference([False, not(False)])

    @staticmethod
    def fm40(globalState):
        return _dafny.Map({_dafny.Set({True, True, True}): 649})

    @staticmethod
    def fm41(p0, p1, p2, globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({False: not(False)}))) | (_dafny.Map({True: False}))

    @staticmethod
    def fm43(p0, p1, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(853, 3):
                d_54_v0_: int = compr_9_
                if ((853) <= (d_54_v0_)) and ((d_54_v0_) < (3)):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_54_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsah"))))]))
            return _dafny.Set(coll9_)
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(-232, 274):
                d_55_v1_: int = compr_10_
                if ((-232) <= (d_55_v1_)) and ((d_55_v1_) < (274)):
                    coll10_ = coll10_.union(_dafny.Set([(d_55_v1_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jynfus"))))]))
            return _dafny.Set(coll10_)
        return (_dafny.Set({(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_53_i0_ in range(default__.abs(791))]))))})) | ((iife9_()
        ).intersection(iife10_()
        ))

    @staticmethod
    def fm44(p0, globalState):
        return _dafny.MultiSet([_dafny.CodePoint('q'), _dafny.CodePoint('s')])

    @staticmethod
    def fm45(globalState):
        return ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(True), True, True])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_56_i0_ in range(default__.abs(-220))])), len(_dafny.SeqWithoutIsStrInference([236])), (_dafny.MultiSet([971, -640, (_dafny.MultiSet([_dafny.Set({not(True), True})])).cardinality, len(_dafny.Set({False, True}))])).cardinality]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True]): _dafny.CodePoint('x')})) for d_57_i1_ in range(default__.abs(-804))])))) | (_dafny.MultiSet([473, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmbwmks")))]))

    @staticmethod
    def fm46(p0, globalState):
        if not((False if False else not(False))):
            return (_dafny.Map({False: True})) | (_dafny.Map({False: True}))
        elif True:
            return (_dafny.Map({True: True})) | (_dafny.Map({False: True}))

    @staticmethod
    def fm47(p0, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_58_i0_ in range(default__.abs(305))])): _dafny.Map({False: True})})

    @staticmethod
    def fm48(p0, p1, globalState):
        return _dafny.MultiSet([not(not(True)), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ju"))), (not(not(True))) in (_dafny.SeqWithoutIsStrInference([False, not(True)]))])

    @staticmethod
    def fm49(globalState):
        if True:
            return _dafny.Set({True})
        elif not(True):
            return _dafny.Set({True, not(False), True, True, True})
        elif True:
            return _dafny.Set({True, True})

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        def iife11_():
            coll11_ = _dafny.Map()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(247, 643):
                d_59_v0_: int = compr_11_
                if ((247) <= (d_59_v0_)) and ((d_59_v0_) < (643)):
                    coll11_[(d_59_v0_) + (454)] = not(True)
            return _dafny.Map(coll11_)
        return _dafny.Set({iife11_()
        })

    @staticmethod
    def fm51(p0, p1, p2, globalState):
        return _dafny.Set({_dafny.CodePoint('r')})

    @staticmethod
    def fm52(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-816 for d_60_i0_ in range(default__.abs(398))]), (D6_DC13(_dafny.SeqWithoutIsStrInference([247 for d_61_i1_ in range(default__.abs(772))]))).cf14, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-873])).cardinality, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmsk"))))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({262: (_dafny.MultiSet([not(not(True)), True])).cardinality})) for d_62_i2_ in range(default__.abs(513))])])

    @staticmethod
    def fm53(p0, p1, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm54(globalState):
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwfty"))})).Elements:
                d_63_v0_: _dafny.Seq = compr_12_
                if (d_63_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwfty"))})):
                    coll12_ = coll12_.union(_dafny.Set([d_63_v0_]))
            return _dafny.Set(coll12_)
        return D15_DC28(False, not(not(not((48) == (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([D19_DC41(D19_DC38(iife12_()
)), D19_DC41(D19_DC41(D19_DC40(True, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wkltshb")))), _dafny.Map({467: _dafny.SeqWithoutIsStrInference([111, 871])}))))]): True})))))))

    @staticmethod
    def fm55(p0, p1, globalState):
        def iife13_():
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.MultiSet([True])])).cardinality])).Elements:
                    d_64_v1_: int = compr_15_
                    if (d_64_v1_) in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.MultiSet([True])])).cardinality])):
                        coll15_[default__.safeModuloInt(d_64_v1_, 67)] = True
                return _dafny.Map(coll15_)
            coll13_ = _dafny.Map()
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.MultiSet([True])])).cardinality])).Elements:
                    d_64_v1_: int = compr_14_
                    if (d_64_v1_) in (_dafny.MultiSet([(_dafny.MultiSet([_dafny.MultiSet([True])])).cardinality])):
                        coll14_[default__.safeModuloInt(d_64_v1_, 67)] = True
                return _dafny.Map(coll14_)
            compr_13_: int
            for compr_13_ in (_dafny.SeqWithoutIsStrInference([len(iife14_()
            )])).Elements:
                d_65_v0_: int = compr_13_
                if (d_65_v0_) in (_dafny.SeqWithoutIsStrInference([len(iife15_()
                )])):
                    coll13_[(d_65_v0_) - (520)] = _dafny.SeqWithoutIsStrInference([False])
            return _dafny.Map(coll13_)
        return ((_dafny.Map({479: _dafny.SeqWithoutIsStrInference([True])})) | (iife13_()
        )) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([D16_DC30(_dafny.SeqWithoutIsStrInference([_dafny.Set({False, False})])) for d_66_i0_ in range(default__.abs(979))])): _dafny.SeqWithoutIsStrInference([True, True, True])}))

    @staticmethod
    def fm56(globalState):
        def iife16_():
            coll16_ = _dafny.Set()
            compr_16_: str
            for compr_16_ in (_dafny.MultiSet([_dafny.CodePoint('v')])).Elements:
                d_67_v0_: str = compr_16_
                if (d_67_v0_) in (_dafny.MultiSet([_dafny.CodePoint('v')])):
                    coll16_ = coll16_.union(_dafny.Set([d_67_v0_]))
            return _dafny.Set(coll16_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: D12_DC21(_dafny.Map({False: not(False)}))}))])) + (_dafny.SeqWithoutIsStrInference([len(iife16_()
        ), (0) - ((_dafny.MultiSet([373])).cardinality)]))) + ((_dafny.SeqWithoutIsStrInference([829])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xibg"))) for d_68_i0_ in range(default__.abs(361))])))

    @staticmethod
    def fm57(p0, p1, globalState):
        return (_dafny.Map({True: 350})) | ((_dafny.Map({not(True): 81})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eauvvtm"))]))})))

    @staticmethod
    def fm58(globalState):
        if (D26_DC57(True)).cf77:
            return D11_DC20(False, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugle"))): 317})))
        elif not(False):
            return D11_DC20(False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('a'): _dafny.MultiSet([_dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('u')])})), 711])))
        elif True:
            return D11_DC20(True, -365)

    @staticmethod
    def fm59(globalState):
        source4_ = D0_DC0((D26_DC57(not(True))).cf77)
        if source4_.is_DC1:
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: str
                for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('y')])).Elements:
                    d_69_v0_: str = compr_17_
                    if (d_69_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('y')])):
                        coll17_[d_69_v0_] = not(True)
                return _dafny.Map(coll17_)
            return (_dafny.Map({_dafny.CodePoint('v'): True})) | (iife17_()
            )
        elif source4_.is_DC2:
            d_70___mcc_h0_ = source4_.cf1
            d_71___mcc_h1_ = source4_.cf2
            d_72_cf2_ = d_71___mcc_h1_
            d_73_cf1_ = d_70___mcc_h0_
            return (_dafny.Map({_dafny.CodePoint('h'): True})) | (_dafny.Map({d_72_cf2_: False}))
        elif source4_.is_DC0:
            d_74___mcc_h2_ = source4_.cf0
            d_75_cf0_ = d_74___mcc_h2_
            return (_dafny.Map({_dafny.CodePoint('u'): d_75_cf0_})) | (_dafny.Map({_dafny.CodePoint('o'): d_75_cf0_}))
        elif True:
            d_76___mcc_h3_ = source4_.cf3
            d_77_cf3_ = d_76___mcc_h3_
            return _dafny.Map({_dafny.CodePoint('o'): True})

    @staticmethod
    def fm60(globalState):
        return D14_DC26((True) and (True), ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxkxihvnx")))))) + (113), _dafny.CodePoint('v'))

    @staticmethod
    def fm61(p0, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jemqeaufg")), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_78_i0_ in range(default__.abs(723))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyqd"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yk"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "acxl")), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqlctexte"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alohcbeea")))])

    @staticmethod
    def fm62(p0, p1, p2, p3, globalState):
        source5_ = D11_DC19(_dafny.CodePoint('g'))
        if source5_.is_DC20:
            d_79___mcc_h0_ = source5_.cf21
            d_80___mcc_h1_ = source5_.cf22
            d_81_cf22_ = d_80___mcc_h1_
            d_82_cf21_ = d_79___mcc_h0_
            return D2_DC8(d_82_cf21_)
        elif True:
            d_83___mcc_h2_ = source5_.cf20
            d_84_cf20_ = d_83___mcc_h2_
            return D2_DC8(False)

    @staticmethod
    def fm63(p0, globalState):
        return _dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_85_i0_ in range(default__.abs(736))])})

    @staticmethod
    def fm64(p0, globalState):
        return D6_DC14((True) == (False))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        return D17_DC34(_dafny.CodePoint('r'))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: D0 = D0.default()()
        r1: bool = False
        d_86_v0_: int
        d_86_v0_ = 842
        d_87_v1_: _dafny.Seq
        d_87_v1_ = _dafny.SeqWithoutIsStrInference([d_86_v0_])
        d_88_v2_: _dafny.Map
        d_88_v2_ = _dafny.Map({d_86_v0_: d_87_v1_})
        d_89_v3_: D19
        d_89_v3_ = D19_DC40(not((not(p2)) == (p2)), d_86_v0_, d_88_v2_)
        d_89_v3_ = d_89_v3_
        (globalState).f1 = default__.safeModuloInt(d_86_v0_, 239)
        d_90_i0_: int
        d_90_i0_ = 0
        with _dafny.label("0"):
            while default__.fm0(p2, globalState):
                with _dafny.c_label("0"):
                    if (d_90_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_90_i0_ = (d_90_i0_) + (1)
                    d_91_v4_: _dafny.Array
                    nw0_ = _dafny.Array(int(0), 13)
                    d_91_v4_ = nw0_
                    d_92_v5_: _dafny.Set
                    d_92_v5_ = _dafny.Set({d_91_v4_})
                    d_93_v6_: _dafny.Map
                    d_93_v6_ = _dafny.Map({len(d_92_v5_): D21_DC47(p2)})
                    d_94_v7_: _dafny.Map
                    d_94_v7_ = _dafny.Map({d_93_v6_: p2})
                    d_95_v8_: _dafny.Map
                    d_95_v8_ = _dafny.Map({p3: d_93_v6_})
                    r1 = ((d_94_v7_)[((d_95_v8_)[p3] if (p3) in (d_95_v8_) else d_93_v6_)] if (((d_95_v8_)[p3] if (p3) in (d_95_v8_) else d_93_v6_)) in (d_94_v7_) else not(p2))
                    d_96_v9_: _dafny.Seq
                    d_96_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ioa"))
                    rhs0_ = (0) - (default__.fm3(p2, p2, globalState))
                    rhs1_ = p2
                    rhs2_ = (d_87_v1_)[default__.safeIndex((d_86_v0_) + (-259), len(d_87_v1_))]
                    rhs3_ = p2
                    rhs4_ = d_96_v9_
                    lhs0_ = globalState
                    lhs1_ = globalState
                    lhs0_.f1 = rhs0_
                    r1 = rhs1_
                    lhs1_.f1 = rhs2_
                    r1 = rhs3_
                    d_96_v9_ = rhs4_
                    (globalState).f1 = default__.safeDivisionInt(d_86_v0_, d_86_v0_)
                    r1 = p2
                    pass
            pass
        (globalState).f1 = (0) - ((d_86_v0_) - (d_86_v0_))
        d_97_v10_: D21
        d_97_v10_ = D21_DC46(p2, p2, d_86_v0_)
        def lambda0_(source6_):
            if source6_.is_DC46:
                d_98___mcc_h0_ = source6_.cf61
                d_99___mcc_h1_ = source6_.cf62
                d_100___mcc_h2_ = source6_.cf63
                d_101_cf63_ = d_100___mcc_h2_
                d_102_cf62_ = d_99___mcc_h1_
                d_103_cf61_ = d_98___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcyxtburc")))
            elif source6_.is_DC47:
                d_104___mcc_h3_ = source6_.cf64
                d_105_cf64_ = d_104___mcc_h3_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
            elif True:
                d_106___mcc_h4_ = source6_.cf60
                d_107_cf60_ = d_106___mcc_h4_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oesqkuq"))

        (globalState).f1 = len(lambda0_(d_97_v10_))
        d_108_i1_: int
        d_108_i1_ = 0
        with _dafny.label("1"):
            while (d_86_v0_) != (d_86_v0_):
                with _dafny.c_label("1"):
                    if (d_108_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_108_i1_ = (d_108_i1_) + (1)
                    d_109_v11_: _dafny.Array
                    nw1_ = _dafny.Array(_dafny.MultiSet({}), 24)
                    d_109_v11_ = nw1_
                    d_110_v12_: _dafny.MultiSet
                    d_110_v12_ = _dafny.MultiSet([d_86_v0_, d_86_v0_])
                    index0_ = default__.safeIndex(938, (d_109_v11_).length(0))
                    (d_109_v11_)[index0_] = (d_110_v12_) | ((d_110_v12_).set(145, default__.abs(d_86_v0_)))
                    index1_ = default__.safeIndex(938, (d_109_v11_).length(0))
                    (d_109_v11_)[index1_] = d_110_v12_
                    if p2:
                        d_111_v13_: _dafny.Seq
                        d_111_v13_ = _dafny.SeqWithoutIsStrInference([True, p2])
                        d_112_v14_: _dafny.Map
                        d_112_v14_ = _dafny.Map({(d_111_v13_)[default__.safeIndex(d_86_v0_, len(d_111_v13_))]: p2})
                        r1 = (d_111_v13_)[default__.safeIndex(len(d_112_v14_), len(d_111_v13_))]
                        d_113_v15_: D11
                        d_113_v15_ = D11_DC20(p2, ((d_109_v11_)[default__.safeIndex(938, (d_109_v11_).length(0))]).cardinality)
                        d_114_v16_: _dafny.Map
                        d_114_v16_ = _dafny.Map({(d_113_v15_).cf21: p0})
                        d_115_v17_: _dafny.Set
                        d_115_v17_ = _dafny.Set({d_86_v0_, d_86_v0_})
                        d_116_v18_: _dafny.Map
                        d_116_v18_ = _dafny.Map({d_115_v17_: p0})
                        d_117_v19_: _dafny.Array
                        nw2_ = _dafny.Array(None, 25)
                        nw2_[int(0)] = p0
                        nw2_[int(1)] = p0
                        nw2_[int(2)] = ((d_114_v16_)[not(p2)] if (not(p2)) in (d_114_v16_) else p0)
                        nw2_[int(3)] = p0
                        nw2_[int(4)] = p0
                        nw2_[int(5)] = p0
                        nw2_[int(6)] = p0
                        nw2_[int(7)] = p0
                        nw2_[int(8)] = p0
                        nw2_[int(9)] = p0
                        nw2_[int(10)] = p0
                        nw2_[int(11)] = p0
                        nw2_[int(12)] = p0
                        nw2_[int(13)] = p0
                        nw2_[int(14)] = p0
                        nw2_[int(15)] = p0
                        nw2_[int(16)] = p0
                        nw2_[int(17)] = ((d_116_v18_)[_dafny.Set({148, d_86_v0_})] if (_dafny.Set({148, d_86_v0_})) in (d_116_v18_) else p0)
                        nw2_[int(18)] = (p0 if p2 else p0)
                        nw2_[int(19)] = p0
                        nw2_[int(20)] = p0
                        nw2_[int(21)] = p0
                        nw2_[int(22)] = p0
                        nw2_[int(23)] = p0
                        nw2_[int(24)] = p0
                        d_117_v19_ = nw2_
                        index2_ = default__.safeIndex(689, (d_117_v19_).length(0))
                        (d_117_v19_)[index2_] = p0
                        index3_ = default__.safeIndex(689, (d_117_v19_).length(0))
                        (d_117_v19_)[index3_] = p0
                        d_118_v20_: _dafny.Array
                        def lambda1_(d_119_v0_):
                            def lambda2_(d_120_i2_):
                                return (d_120_i2_) + (d_119_v0_)

                            return lambda2_

                        init0_ = lambda1_(d_86_v0_)
                        nw3_ = _dafny.Array(None, 28)
                        for i0_0_ in range(nw3_.length(0)):
                            nw3_[i0_0_] = init0_(i0_0_)
                        d_118_v20_ = nw3_
                        d_121_v21_: _dafny.Array
                        nw4_ = _dafny.Array(None, 22)
                        nw4_[int(0)] = d_118_v20_
                        nw4_[int(1)] = d_118_v20_
                        nw4_[int(2)] = d_118_v20_
                        nw4_[int(3)] = d_118_v20_
                        nw4_[int(4)] = d_118_v20_
                        nw4_[int(5)] = d_118_v20_
                        nw4_[int(6)] = d_118_v20_
                        nw4_[int(7)] = d_118_v20_
                        nw4_[int(8)] = d_118_v20_
                        nw4_[int(9)] = d_118_v20_
                        nw4_[int(10)] = d_118_v20_
                        nw4_[int(11)] = d_118_v20_
                        nw4_[int(12)] = d_118_v20_
                        nw4_[int(13)] = (d_118_v20_)
                        nw4_[int(14)] = d_118_v20_
                        nw4_[int(15)] = d_118_v20_
                        nw4_[int(16)] = d_118_v20_
                        nw4_[int(17)] = d_118_v20_
                        nw4_[int(18)] = d_118_v20_
                        nw4_[int(19)] = d_118_v20_
                        nw4_[int(20)] = (d_118_v20_ if p2 else d_118_v20_)
                        nw4_[int(21)] = d_118_v20_
                        d_121_v21_ = nw4_
                        index4_ = default__.safeIndex(338, (d_121_v21_).length(0))
                        (d_121_v21_)[index4_] = d_118_v20_
                        index5_ = default__.safeIndex(338, (d_121_v21_).length(0))
                        (d_121_v21_)[index5_] = d_118_v20_
                        d_122_v22_: _dafny.Map
                        d_122_v22_ = _dafny.Map({p0: p2})
                        d_122_v22_ = (d_122_v22_).set((d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))], p2)
                        arr0_ = (d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]
                        index6_ = default__.safeIndex(943, ((d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]).length(0))
                        arr0_[index6_] = p2
                        arr1_ = (d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]
                        index7_ = default__.safeIndex(943, ((d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]).length(0))
                        rhs5_ = (0) - (d_86_v0_)
                        rhs6_ = (-243) - (d_86_v0_)
                        rhs7_ = True
                        rhs8_ = (d_111_v13_)[default__.safeIndex((0) - (d_86_v0_), len(d_111_v13_))]
                        lhs2_ = globalState
                        lhs3_ = (d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]
                        lhs4_ = default__.safeIndex(943, ((d_117_v19_)[default__.safeIndex(689, (d_117_v19_).length(0))]).length(0))
                        lhs2_.f1 = rhs5_
                        d_86_v0_ = rhs6_
                        lhs3_[lhs4_] = rhs7_
                        r1 = rhs8_
                    elif True:
                        (globalState).f1 = (0) - (d_86_v0_)
                        d_123_v23_: _dafny.Array
                        def lambda3_(d_124_p2_):
                            def lambda4_(d_125_i3_):
                                return (_dafny.SeqWithoutIsStrInference([d_124_p2_, d_124_p2_])) + (_dafny.SeqWithoutIsStrInference([True, d_124_p2_]))

                            return lambda4_

                        init1_ = lambda3_(p2)
                        nw5_ = _dafny.Array(None, 19)
                        for i0_1_ in range(nw5_.length(0)):
                            nw5_[i0_1_] = init1_(i0_1_)
                        d_123_v23_ = nw5_
                        d_126_v24_: _dafny.Seq
                        d_126_v24_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                        d_127_v25_: _dafny.Seq
                        d_127_v25_ = _dafny.SeqWithoutIsStrInference([(d_126_v24_)[default__.safeIndex(len(d_87_v1_), len(d_126_v24_))], p2])
                        index8_ = default__.safeIndex(940, (d_123_v23_).length(0))
                        (d_123_v23_)[index8_] = d_126_v24_
                        index9_ = default__.safeIndex(940, (d_123_v23_).length(0))
                        (d_123_v23_)[index9_] = (d_127_v25_) + (d_126_v24_)
                        d_128_v26_: D14
                        d_128_v26_ = D14_DC26(True, d_86_v0_, p3)
                        d_129_v27_: C7
                        nw6_ = C7()
                        nw6_.ctor__((112) + (d_86_v0_), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oabcxl"))), d_86_v0_), p3, (d_128_v26_).cf32)
                        d_129_v27_ = nw6_
                        d_130_v28_: _dafny.Set
                        d_130_v28_ = _dafny.Set({(d_109_v11_)[default__.safeIndex(938, (d_109_v11_).length(0))], d_110_v12_, (d_109_v11_)[default__.safeIndex(938, (d_109_v11_).length(0))]})
                        r1 = ((d_130_v28_) | (d_130_v28_)) == ((d_130_v28_) | (d_130_v28_))
                        d_131_v29_: _dafny.Map
                        d_131_v29_ = _dafny.Map({-444: d_129_v27_.f10})
                        d_132_v30_: _dafny.Set
                        d_132_v30_ = _dafny.Set({p2, not(p2)})
                        d_133_v31_: _dafny.Seq
                        d_133_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycbrppf"))
                        d_134_v32_: _dafny.MultiSet
                        d_134_v32_ = _dafny.MultiSet([d_133_v31_])
                        d_135_v33_: _dafny.Map
                        d_135_v33_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_136_i4_ in range(default__.abs(402))])})
                        d_137_v34_: _dafny.Array
                        nw7_ = _dafny.Array(None, 18)
                        nw7_[int(0)] = d_86_v0_
                        nw7_[int(1)] = d_86_v0_
                        nw7_[int(2)] = d_129_v27_.f10
                        nw7_[int(3)] = len(d_131_v29_)
                        nw7_[int(4)] = (d_129_v27_).f9
                        nw7_[int(5)] = d_86_v0_
                        nw7_[int(6)] = len(_dafny.SeqWithoutIsStrInference([d_129_v27_.f10, len(d_132_v30_), (d_134_v32_).cardinality, len(d_135_v33_)]))
                        nw7_[int(7)] = d_129_v27_.f10
                        nw7_[int(8)] = 552
                        nw7_[int(9)] = d_86_v0_
                        nw7_[int(10)] = (d_129_v27_).f9
                        nw7_[int(11)] = d_86_v0_
                        nw7_[int(12)] = d_86_v0_
                        nw7_[int(13)] = 60
                        nw7_[int(14)] = (d_129_v27_).f9
                        nw7_[int(15)] = ((d_109_v11_)[default__.safeIndex(938, (d_109_v11_).length(0))]).cardinality
                        nw7_[int(16)] = (d_129_v27_).f9
                        nw7_[int(17)] = d_129_v27_.f10
                        d_137_v34_ = nw7_
                        d_138_v35_: _dafny.Map
                        d_138_v35_ = _dafny.Map({d_137_v34_: (d_129_v27_).f9})
                        d_139_v36_: _dafny.Map
                        d_139_v36_ = _dafny.Map({p2: p2})
                        d_138_v35_ = (d_138_v35_).set(d_137_v34_, (d_129_v27_.f10) - (len(d_139_v36_)))
                    d_140_v37_: _dafny.Seq
                    d_140_v37_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_141_v38_: _dafny.Seq
                    d_141_v38_ = _dafny.SeqWithoutIsStrInference([p2, p2, (d_140_v37_)[default__.safeIndex(d_86_v0_, len(d_140_v37_))]])
                    d_86_v0_ = (len((d_140_v37_) + (d_141_v38_)) if (p2 if p2 else p2) else d_86_v0_)
                    if (default__.safeDivisionInt((d_87_v1_)[default__.safeIndex(d_86_v0_, len(d_87_v1_))], d_86_v0_)) >= ((d_86_v0_) - (-634)):
                        d_142_v39_: _dafny.Array
                        nw8_ = _dafny.Array(None, 6)
                        nw8_[int(0)] = p1
                        nw8_[int(1)] = p1
                        nw8_[int(2)] = p3
                        nw8_[int(3)] = p1
                        nw8_[int(4)] = p1
                        nw8_[int(5)] = p1
                        d_142_v39_ = nw8_
                        d_143_v40_: _dafny.MultiSet
                        d_143_v40_ = _dafny.MultiSet([d_142_v39_, d_142_v39_])
                        d_144_v41_: T3
                        nw9_ = C7()
                        nw9_.ctor__(d_86_v0_, d_86_v0_, p3, p2)
                        d_144_v41_ = nw9_
                        d_145_v42_: _dafny.Map
                        d_145_v42_ = _dafny.Map({p2: d_144_v41_})
                        d_146_v43_: _dafny.Map
                        d_146_v43_ = _dafny.Map({d_143_v40_: (d_145_v42_) | (d_145_v42_)})
                        d_147_v44_: _dafny.Seq
                        d_147_v44_ = _dafny.SeqWithoutIsStrInference([d_145_v42_, d_145_v42_, d_145_v42_])
                        d_146_v43_ = (d_146_v43_).set((d_143_v40_) | (d_143_v40_), (d_147_v44_)[default__.safeIndex(d_86_v0_, len(d_147_v44_))])
                        d_148_v45_: _dafny.Array
                        nw10_ = _dafny.Array(_dafny.Array(None, 0), 18)
                        d_148_v45_ = nw10_
                        index10_ = default__.safeIndex(825, (p0).length(0))
                        (p0)[index10_] = (865) != (d_86_v0_)
                        d_149_v46_: D26
                        d_149_v46_ = D26_DC55(d_148_v45_)
                        index11_ = default__.safeIndex(825, (p0).length(0))
                        rhs9_ = (d_149_v46_).cf74
                        rhs10_ = p2
                        lhs5_ = p0
                        lhs6_ = default__.safeIndex(825, (p0).length(0))
                        d_148_v45_ = rhs9_
                        lhs5_[lhs6_] = rhs10_
                        d_150_v47_: _dafny.Map
                        d_150_v47_ = _dafny.Map({p1: d_86_v0_})
                        d_151_v48_: _dafny.Set
                        d_151_v48_ = _dafny.Set({642})
                        d_152_v49_: _dafny.Set
                        d_152_v49_ = _dafny.Set({d_151_v48_, d_151_v48_})
                        index12_ = default__.safeIndex(825, (p0).length(0))
                        index13_ = default__.safeIndex(825, (p0).length(0))
                        rhs11_ = (d_150_v47_).set(p3, (0) - (d_86_v0_))
                        rhs12_ = (d_144_v41_).f6
                        rhs13_ = (d_152_v49_).ispropersubset(_dafny.Set({d_151_v48_}))
                        lhs7_ = p0
                        lhs8_ = default__.safeIndex(825, (p0).length(0))
                        lhs9_ = p0
                        lhs10_ = default__.safeIndex(825, (p0).length(0))
                        d_150_v47_ = rhs11_
                        lhs7_[lhs8_] = rhs12_
                        lhs9_[lhs10_] = rhs13_
                        d_153_v50_: _dafny.Map
                        d_153_v50_ = _dafny.Map({(d_144_v41_).f6: d_87_v1_})
                        d_154_v51_: _dafny.MultiSet
                        d_154_v51_ = _dafny.MultiSet([p2])
                        d_155_v52_: _dafny.Seq
                        d_155_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swp"))
                        d_153_v50_ = (d_153_v50_).set(((d_154_v51_).cardinality) > (len(d_155_v52_)), d_87_v1_)
                        d_156_v53_: _dafny.MultiSet
                        d_156_v53_ = _dafny.MultiSet([d_155_v52_])
                        d_157_v54_: _dafny.Array
                        def lambda5_(d_158_v41_, d_159_v0_):
                            def lambda6_(d_160_i5_):
                                return (d_160_i5_) - (len(_dafny.Map({(d_158_v41_).f6: d_159_v0_})))

                            return lambda6_

                        init2_ = lambda5_(d_144_v41_, d_86_v0_)
                        nw11_ = _dafny.Array(None, 19)
                        for i0_2_ in range(nw11_.length(0)):
                            nw11_[i0_2_] = init2_(i0_2_)
                        d_157_v54_ = nw11_
                        d_161_v55_: _dafny.Map
                        d_161_v55_ = _dafny.Map({d_157_v54_: d_156_v53_})
                        index14_ = default__.safeIndex(825, (p0).length(0))
                        index15_ = default__.safeIndex(825, (p0).length(0))
                        rhs14_ = not(not((d_144_v41_).f6))
                        rhs15_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyay"))) + (d_155_v52_)) not in ((d_156_v53_ if (p0)[default__.safeIndex(825, (p0).length(0))] else ((d_161_v55_)[d_157_v54_] if (d_157_v54_) in (d_161_v55_) else _dafny.MultiSet([d_155_v52_]))))
                        rhs16_ = d_86_v0_
                        rhs17_ = ((d_87_v1_)[default__.safeIndex(d_86_v0_, len(d_87_v1_))]) - (d_86_v0_)
                        rhs18_ = default__.fm0(not(p2), globalState)
                        lhs11_ = p0
                        lhs12_ = default__.safeIndex(825, (p0).length(0))
                        lhs13_ = globalState
                        lhs14_ = globalState
                        lhs15_ = p0
                        lhs16_ = default__.safeIndex(825, (p0).length(0))
                        r1 = rhs14_
                        lhs11_[lhs12_] = rhs15_
                        lhs13_.f1 = rhs16_
                        lhs14_.f1 = rhs17_
                        lhs15_[lhs16_] = rhs18_
                    elif True:
                        d_162_v56_: _dafny.Seq
                        d_162_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xarvj"))
                        d_162_v56_ = d_162_v56_
                        d_163_v57_: C4
                        nw12_ = C4()
                        nw12_.ctor__()
                        d_163_v57_ = nw12_
                        index16_ = default__.safeIndex(213, (p0).length(0))
                        (p0)[index16_] = p2
                        index17_ = default__.safeIndex(213, (p0).length(0))
                        (p0)[index17_] = True
                        (globalState).f1 = (d_86_v0_ if (p0)[default__.safeIndex(213, (p0).length(0))] else d_86_v0_)
                        d_86_v0_ = (0) - (d_86_v0_)
                    pass
            pass
        d_164_v58_: _dafny.Map
        d_164_v58_ = _dafny.Map({p0: p3})
        d_165_v59_: D0
        d_165_v59_ = D0_DC2((d_164_v58_).set(p0, p3), p1)
        r0 = d_165_v59_
        d_166_v60_: _dafny.Set
        d_166_v60_ = _dafny.Set({_dafny.MultiSet(d_87_v1_)})
        r1 = ((d_166_v60_) == (d_166_v60_)) == (default__.fm0(p2, globalState))
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_167_v0_: int
        d_167_v0_ = 647
        d_168_v1_: _dafny.Seq
        d_168_v1_ = _dafny.SeqWithoutIsStrInference([d_167_v0_, d_167_v0_])
        d_169_v2_: _dafny.Seq
        d_169_v2_ = _dafny.SeqWithoutIsStrInference([d_168_v1_])
        d_170_v3_: bool
        d_170_v3_ = True
        d_171_v4_: D0
        d_171_v4_ = D0_DC0(d_170_v3_)
        d_172_v5_: _dafny.Seq
        d_172_v5_ = _dafny.SeqWithoutIsStrInference([d_170_v3_, d_170_v3_, True, d_170_v3_, d_170_v3_])
        d_173_globalState_: GlobalState
        nw13_ = GlobalState()
        nw13_.ctor__(_dafny.CodePoint('i'), 10, 430, d_169_v2_, _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_170_v3_, (d_171_v4_).cf0])) + (d_172_v5_)))
        d_173_globalState_ = nw13_
        d_174_i0_: int
        d_174_i0_ = 0
        with _dafny.label("2"):
            while d_170_v3_:
                with _dafny.c_label("2"):
                    if (d_174_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_174_i0_ = (d_174_i0_) + (1)
                    d_175_v6_: _dafny.Array
                    def lambda7_(d_176_v3_):
                        def lambda8_(d_177_i1_):
                            return d_176_v3_

                        return lambda8_

                    init3_ = lambda7_(d_170_v3_)
                    nw14_ = _dafny.Array(None, 16)
                    for i0_3_ in range(nw14_.length(0)):
                        nw14_[i0_3_] = init3_(i0_3_)
                    d_175_v6_ = nw14_
                    d_178_v7_: str
                    d_178_v7_ = _dafny.CodePoint('w')
                    d_179_v8_: D0
                    d_180_v9_: bool
                    out0_: D0
                    out1_: bool
                    out0_, out1_ = default__.m0(d_175_v6_, d_178_v7_, d_170_v3_, d_178_v7_, d_173_globalState_)
                    d_179_v8_ = out0_
                    d_180_v9_ = out1_
                    d_181_v10_: _dafny.Array
                    nw15_ = _dafny.Array(int(0), 7)
                    d_181_v10_ = nw15_
                    index18_ = default__.safeIndex(975, (d_181_v10_).length(0))
                    (d_181_v10_)[index18_] = d_167_v0_
                    index19_ = default__.safeIndex(975, (d_181_v10_).length(0))
                    (d_181_v10_)[index19_] = (d_167_v0_) + ((d_167_v0_ if d_180_v9_ else d_167_v0_))
                    d_182_v11_: _dafny.Map
                    d_182_v11_ = _dafny.Map({d_167_v0_: d_179_v8_})
                    d_183_v12_: _dafny.Map
                    d_183_v12_ = _dafny.Map({325: d_167_v0_})
                    d_184_v13_: _dafny.Map
                    d_184_v13_ = _dafny.Map({d_183_v12_: d_182_v11_})
                    d_185_v14_: _dafny.Map
                    d_185_v14_ = _dafny.Map({((d_181_v10_)[default__.safeIndex(975, (d_181_v10_).length(0))]) * (d_167_v0_): _dafny.Set({d_182_v11_, ((d_184_v13_)[d_183_v12_] if (d_183_v12_) in (d_184_v13_) else d_182_v11_), d_182_v11_})})
                    d_186_v15_: _dafny.Set
                    d_186_v15_ = _dafny.Set({d_182_v11_})
                    d_185_v14_ = (d_185_v14_).set(default__.safeDivisionInt((d_181_v10_)[default__.safeIndex(975, (d_181_v10_).length(0))], (d_181_v10_)[default__.safeIndex(975, (d_181_v10_).length(0))]), d_186_v15_)
                    d_187_v16_: _dafny.MultiSet
                    d_187_v16_ = _dafny.MultiSet([d_167_v0_, 493])
                    d_188_v17_: _dafny.Map
                    d_188_v17_ = _dafny.Map({default__.fm0(d_170_v3_, d_173_globalState_): default__.safeDivisionInt(379, (d_187_v16_).cardinality)})
                    d_188_v17_ = (d_188_v17_).set(d_170_v3_, d_167_v0_)
                    pass
            pass
        d_189_v18_: D0
        d_189_v18_ = D0_DC3(D0_DC0((d_172_v5_)[default__.safeIndex(d_167_v0_, len(d_172_v5_))]))
        d_189_v18_ = d_189_v18_
        d_190_v19_: _dafny.Array
        nw16_ = _dafny.Array(None, 7)
        nw16_[int(0)] = d_170_v3_
        nw16_[int(1)] = True
        nw16_[int(2)] = d_170_v3_
        nw16_[int(3)] = d_170_v3_
        nw16_[int(4)] = d_170_v3_
        nw16_[int(5)] = d_170_v3_
        nw16_[int(6)] = d_170_v3_
        d_190_v19_ = nw16_
        d_191_v20_: str
        d_191_v20_ = _dafny.CodePoint('j')
        d_192_v21_: D0
        d_193_v22_: bool
        out2_: D0
        out3_: bool
        out2_, out3_ = default__.m0(d_190_v19_, d_191_v20_, d_170_v3_, d_191_v20_, d_173_globalState_)
        d_192_v21_ = out2_
        d_193_v22_ = out3_
        d_194_v23_: _dafny.Map
        d_194_v23_ = _dafny.Map({d_191_v20_: d_167_v0_})
        d_194_v23_ = (d_194_v23_) | ((d_194_v23_) | (d_194_v23_))
        d_191_v20_ = d_191_v20_
        d_195_v24_: _dafny.Array
        nw17_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
        d_195_v24_ = nw17_
        index20_ = default__.safeIndex(731, (d_195_v24_).length(0))
        (d_195_v24_)[index20_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebdgugv"))
        d_196_v25_: _dafny.Seq
        d_196_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ds"))
        index21_ = default__.safeIndex(731, (d_195_v24_).length(0))
        (d_195_v24_)[index21_] = d_196_v25_
        d_197_v26_: _dafny.Map
        d_197_v26_ = _dafny.Map({d_190_v19_: d_191_v20_})
        d_198_v27_: _dafny.Array
        nw18_ = _dafny.Array(None, 4)
        nw18_[int(0)] = d_192_v21_
        nw18_[int(1)] = D0_DC2(d_197_v26_, d_191_v20_)
        nw18_[int(2)] = d_192_v21_
        nw18_[int(3)] = d_192_v21_
        d_198_v27_ = nw18_
        d_199_v28_: _dafny.Seq
        d_199_v28_ = _dafny.SeqWithoutIsStrInference([d_198_v27_])
        d_198_v27_ = (d_199_v28_)[default__.safeIndex(d_167_v0_, len(d_199_v28_))]
        if d_193_v22_:
            index22_ = default__.safeIndex(54, (d_190_v19_).length(0))
            (d_190_v19_)[index22_] = d_170_v3_
            index23_ = default__.safeIndex(54, (d_190_v19_).length(0))
            (d_190_v19_)[index23_] = not(((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]) == (default__.fm1(d_167_v0_, (d_172_v5_)[default__.safeIndex(d_167_v0_, len(d_172_v5_))], 606, d_173_globalState_)))
            d_200_v29_: _dafny.Array
            nw19_ = _dafny.Array(int(0), 9)
            d_200_v29_ = nw19_
            d_200_v29_ = d_200_v29_
            d_201_v31_: _dafny.Map
            d_201_v31_ = _dafny.Map({default__.fm2(d_170_v3_, d_193_v22_, d_173_globalState_): d_168_v1_})
            d_202_v32_: _dafny.Map
            d_202_v32_ = _dafny.Map({d_168_v1_: ((d_201_v31_)[_dafny.SeqWithoutIsStrInference([(d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))], (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]])] if (_dafny.SeqWithoutIsStrInference([(d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))], (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]])) in (d_201_v31_) else (d_169_v2_)[default__.safeIndex(len((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]), len(d_169_v2_))])})
            d_203_v33_: _dafny.MultiSet
            d_203_v33_ = _dafny.MultiSet([d_191_v20_, d_191_v20_])
            d_204_v34_: _dafny.Map
            d_204_v34_ = _dafny.Map({(d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]: d_203_v33_})
            d_205_v35_: D1
            d_205_v35_ = D1_DC4(len(d_204_v34_))
            d_206_v36_: D1
            d_206_v36_ = D1_DC5(not((d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]))
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: _dafny.Seq
                for compr_18_ in ((d_202_v32_).set(d_168_v1_, d_168_v1_)).keys.Elements:
                    d_207_v30_: _dafny.Seq = compr_18_
                    if (d_207_v30_) in ((d_202_v32_).set(d_168_v1_, d_168_v1_)):
                        coll18_[d_207_v30_] = d_193_v22_
                return _dafny.Map(coll18_)
            rhs19_ = -242
            rhs20_ = (len(iife18_()
            )) + ((d_205_v35_).cf4)
            rhs21_ = (d_206_v36_).cf5
            lhs17_ = d_173_globalState_
            lhs17_.f1 = rhs19_
            d_167_v0_ = rhs20_
            d_170_v3_ = rhs21_
            d_208_v37_: _dafny.Set
            d_208_v37_ = _dafny.Set({(0) - ((370 if (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))] else d_167_v0_)), d_167_v0_, d_167_v0_, d_167_v0_})
            def iife19_():
                coll19_ = _dafny.Set()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(270, -534):
                    d_209_v38_: int = compr_19_
                    if ((270) <= (d_209_v38_)) and ((d_209_v38_) < (-534)):
                        coll19_ = coll19_.union(_dafny.Set([(d_209_v38_) * (d_167_v0_)]))
                return _dafny.Set(coll19_)
            rhs22_ = ((0) - (len(((d_169_v2_)[default__.safeIndex(d_167_v0_, len(d_169_v2_))]) + (d_168_v1_)))) + (default__.fm3((d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))], d_170_v3_, d_173_globalState_))
            rhs23_ = (iife19_()
            ) | (d_208_v37_)
            lhs18_ = d_173_globalState_
            lhs18_.f1 = rhs22_
            d_208_v37_ = rhs23_
            d_210_v39_: D2
            d_210_v39_ = D2_DC6((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))])
            if ((d_196_v25_) + ((d_210_v39_).cf6)) != (d_196_v25_):
                d_211_v40_: _dafny.Map
                d_211_v40_ = _dafny.Map({d_167_v0_: (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]})
                d_212_v41_: _dafny.Map
                d_212_v41_ = _dafny.Map({d_170_v3_: ((d_211_v40_)[d_167_v0_] if (d_167_v0_) in (d_211_v40_) else d_170_v3_)})
                d_213_v42_: C7
                nw20_ = C7()
                nw20_.ctor__(d_167_v0_, len((d_212_v41_) | (d_212_v41_)), d_191_v20_, (d_170_v3_ if (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))] else (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]))
                d_213_v42_ = nw20_
                index24_ = default__.safeIndex(164, (d_200_v29_).length(0))
                (d_200_v29_)[index24_] = (d_213_v42_).f9
                index25_ = default__.safeIndex(164, (d_200_v29_).length(0))
                (d_200_v29_)[index25_] = d_167_v0_
                d_214_v43_: _dafny.Array
                nw21_ = _dafny.Array(_dafny.Seq({}), 15)
                d_214_v43_ = nw21_
                index26_ = default__.safeIndex(946, (d_214_v43_).length(0))
                (d_214_v43_)[index26_] = default__.fm2(d_193_v22_, not(d_170_v3_), d_173_globalState_)
                index27_ = default__.safeIndex(946, (d_214_v43_).length(0))
                (d_214_v43_)[index27_] = (_dafny.SeqWithoutIsStrInference([(d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))]])) + (d_172_v5_)
                index28_ = default__.safeIndex(54, (d_190_v19_).length(0))
                index29_ = default__.safeIndex(164, (d_200_v29_).length(0))
                index30_ = default__.safeIndex(54, (d_190_v19_).length(0))
                rhs24_ = d_193_v22_
                rhs25_ = -536
                rhs26_ = (d_213_v42_).fm8(not(((d_213_v42_).fm7(d_193_v22_, d_173_globalState_)) <= (d_167_v0_)), (d_190_v19_)[default__.safeIndex(54, (d_190_v19_).length(0))], d_193_v22_, d_173_globalState_)
                rhs27_ = d_170_v3_
                lhs19_ = d_190_v19_
                lhs20_ = default__.safeIndex(54, (d_190_v19_).length(0))
                lhs21_ = d_200_v29_
                lhs22_ = default__.safeIndex(164, (d_200_v29_).length(0))
                lhs23_ = d_190_v19_
                lhs24_ = default__.safeIndex(54, (d_190_v19_).length(0))
                lhs19_[lhs20_] = rhs24_
                lhs21_[lhs22_] = rhs25_
                d_167_v0_ = rhs26_
                lhs23_[lhs24_] = rhs27_
                index31_ = default__.safeIndex(54, (d_190_v19_).length(0))
                (d_190_v19_)[index31_] = d_193_v22_
            elif True:
                index32_ = default__.safeIndex(54, (d_190_v19_).length(0))
                (d_190_v19_)[index32_] = d_170_v3_
                (d_173_globalState_).f1 = default__.fm3(d_193_v22_, d_193_v22_, d_173_globalState_)
                d_215_v44_: C6
                nw22_ = C6()
                nw22_.ctor__()
                d_215_v44_ = nw22_
                d_216_v45_: _dafny.Array
                nw23_ = _dafny.Array(False, 9)
                d_216_v45_ = nw23_
                d_217_v46_: _dafny.Seq
                d_217_v46_ = _dafny.SeqWithoutIsStrInference([d_216_v45_])
                d_218_v47_: _dafny.Array
                nw24_ = _dafny.Array(None, 7)
                nw24_[int(0)] = (d_216_v45_ if d_193_v22_ else d_216_v45_)
                nw24_[int(1)] = d_216_v45_
                nw24_[int(2)] = (d_217_v46_)[default__.safeIndex(d_167_v0_, len(d_217_v46_))]
                nw24_[int(3)] = d_216_v45_
                nw24_[int(4)] = d_216_v45_
                nw24_[int(5)] = d_216_v45_
                nw24_[int(6)] = d_216_v45_
                d_218_v47_ = nw24_
                index33_ = default__.safeIndex(176, (d_218_v47_).length(0))
                (d_218_v47_)[index33_] = d_216_v45_
                index34_ = default__.safeIndex(176, (d_218_v47_).length(0))
                (d_218_v47_)[index34_] = (d_217_v46_)[default__.safeIndex((0) - (d_167_v0_), len(d_217_v46_))]
                index35_ = default__.safeIndex(54, (d_190_v19_).length(0))
                (d_190_v19_)[index35_] = not(not(d_170_v3_))
        elif True:
            d_219_v48_: _dafny.Set
            d_219_v48_ = _dafny.Set({d_167_v0_, len((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))])})
            d_167_v0_ = default__.fm3((False) and (d_193_v22_), (d_219_v48_).ispropersubset(_dafny.Set({306, d_167_v0_})), d_173_globalState_)
            if (d_170_v3_) or (default__.fm0(d_193_v22_, d_173_globalState_)):
                d_220_v49_: _dafny.Array
                def lambda9_(d_221_v20_, d_222_v0_):
                    def lambda10_(d_223_i2_):
                        return _dafny.Map({d_221_v20_: d_222_v0_})

                    return lambda10_

                init4_ = lambda9_(d_191_v20_, d_167_v0_)
                nw25_ = _dafny.Array(None, 21)
                for i0_4_ in range(nw25_.length(0)):
                    nw25_[i0_4_] = init4_(i0_4_)
                d_220_v49_ = nw25_
                d_224_v50_: _dafny.Seq
                d_224_v50_ = _dafny.SeqWithoutIsStrInference([d_220_v49_, d_220_v49_, d_220_v49_, d_220_v49_, d_220_v49_])
                d_225_v51_: _dafny.Map
                d_225_v51_ = _dafny.Map({d_167_v0_: len(d_172_v5_)})
                d_226_v52_: _dafny.Array
                nw26_ = _dafny.Array(None, 5)
                nw26_[int(0)] = d_220_v49_
                nw26_[int(1)] = d_220_v49_
                nw26_[int(2)] = (d_224_v50_)[default__.safeIndex(len(d_225_v51_), len(d_224_v50_))]
                nw26_[int(3)] = d_220_v49_
                nw26_[int(4)] = d_220_v49_
                d_226_v52_ = nw26_
                index36_ = default__.safeIndex(812, (d_226_v52_).length(0))
                (d_226_v52_)[index36_] = d_220_v49_
                index37_ = default__.safeIndex(812, (d_226_v52_).length(0))
                rhs28_ = default__.safeModuloInt(((0) - (d_167_v0_)) * (d_167_v0_), d_167_v0_)
                rhs29_ = d_220_v49_
                lhs25_ = d_173_globalState_
                lhs26_ = d_226_v52_
                lhs27_ = default__.safeIndex(812, (d_226_v52_).length(0))
                lhs25_.f1 = rhs28_
                lhs26_[lhs27_] = rhs29_
                d_227_v53_: C6
                nw27_ = C6()
                nw27_.ctor__()
                d_227_v53_ = nw27_
                d_228_v54_: _dafny.Map
                d_228_v54_ = _dafny.Map({d_227_v53_: d_193_v22_})
                d_229_v55_: _dafny.Map
                d_229_v55_ = _dafny.Map({(d_228_v54_) | (_dafny.Map({d_227_v53_: d_170_v3_})): 719})
                (d_173_globalState_).f1 = len(d_229_v55_)
                d_167_v0_ = (0) - (d_167_v0_)
                d_170_v3_ = not (default__.fm0((d_227_v53_).fm6(_dafny.SeqWithoutIsStrInference([d_167_v0_]), d_173_globalState_), d_173_globalState_)) or (d_170_v3_)
                d_230_v56_: int
                out4_: int
                out4_ = (d_227_v53_).m3(d_173_globalState_)
                d_230_v56_ = out4_
            elif True:
                d_219_v48_ = d_219_v48_
                d_231_v57_: _dafny.Map
                d_231_v57_ = _dafny.Map({d_193_v22_: d_167_v0_})
                d_232_v58_: _dafny.MultiSet
                d_232_v58_ = _dafny.MultiSet([d_231_v57_, d_231_v57_])
                d_233_v60_: _dafny.MultiSet
                d_233_v60_ = _dafny.MultiSet([len(d_231_v57_)])
                d_234_v61_: _dafny.Map
                def iife20_():
                    coll20_ = _dafny.Map()
                    compr_20_: int
                    for compr_20_ in (d_233_v60_).Elements:
                        d_235_v59_: int = compr_20_
                        if (d_235_v59_) in (d_233_v60_):
                            coll20_[(d_235_v59_) + (d_167_v0_)] = d_193_v22_
                    return _dafny.Map(coll20_)
                d_234_v61_ = iife20_()
                
                d_236_v62_: _dafny.Map
                d_236_v62_ = _dafny.Map({d_234_v61_: d_170_v3_})
                d_237_v63_: _dafny.Seq
                d_237_v63_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_170_v3_: len(d_236_v62_)}), d_231_v57_])
                d_238_v64_: _dafny.Seq
                d_238_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_237_v63_)])
                d_239_v65_: _dafny.Map
                d_239_v65_ = _dafny.Map({d_167_v0_: d_232_v58_})
                rhs30_ = (d_232_v58_) | (((d_238_v64_)[default__.safeIndex((0) - (d_167_v0_), len(d_238_v64_))] if d_193_v22_ else ((d_239_v65_)[d_167_v0_] if (d_167_v0_) in (d_239_v65_) else d_232_v58_)))
                rhs31_ = (len(d_172_v5_)) * (d_167_v0_)
                d_232_v58_ = rhs30_
                d_167_v0_ = rhs31_
                d_240_v66_: _dafny.Array
                def lambda11_(d_241_v20_):
                    def lambda12_(d_242_i3_):
                        return d_241_v20_

                    return lambda12_

                init5_ = lambda11_(d_191_v20_)
                nw28_ = _dafny.Array(None, 20)
                for i0_5_ in range(nw28_.length(0)):
                    nw28_[i0_5_] = init5_(i0_5_)
                d_240_v66_ = nw28_
                d_243_v67_: _dafny.Map
                d_243_v67_ = _dafny.Map({d_240_v66_: d_170_v3_})
                d_243_v67_ = (d_243_v67_).set(d_240_v66_, (d_167_v0_) >= ((d_233_v60_).cardinality))
                d_244_v68_: C6
                nw29_ = C6()
                nw29_.ctor__()
                d_244_v68_ = nw29_
                d_245_v69_: _dafny.Map
                d_245_v69_ = _dafny.Map({d_170_v3_: d_170_v3_})
                d_246_v70_: _dafny.MultiSet
                d_246_v70_ = _dafny.MultiSet([(d_245_v69_) | (d_245_v69_), (d_245_v69_).set(d_170_v3_, d_170_v3_)])
                index38_ = default__.safeIndex(671, (d_190_v19_).length(0))
                (d_190_v19_)[index38_] = (d_170_v3_) not in (d_245_v69_)
                d_247_v71_: _dafny.Map
                d_247_v71_ = _dafny.Map({d_167_v0_: len(d_219_v48_)})
                index39_ = default__.safeIndex(671, (d_190_v19_).length(0))
                rhs32_ = _dafny.MultiSet([d_245_v69_, ((_dafny.Map({d_193_v22_: d_170_v3_})).set(d_170_v3_, True)).set(d_193_v22_, d_193_v22_), d_245_v69_])
                rhs33_ = (d_168_v1_) + (d_168_v1_)
                rhs34_ = ((d_170_v3_) == (d_170_v3_) if d_193_v22_ else d_170_v3_)
                rhs35_ = (0) - (default__.safeDivisionInt(d_167_v0_, ((d_247_v71_)[929] if (929) in (d_247_v71_) else (0) - (d_167_v0_))))
                lhs28_ = d_190_v19_
                lhs29_ = default__.safeIndex(671, (d_190_v19_).length(0))
                lhs30_ = d_173_globalState_
                d_246_v70_ = rhs32_
                d_168_v1_ = rhs33_
                lhs28_[lhs29_] = rhs34_
                lhs30_.f1 = rhs35_
            d_248_v72_: _dafny.Map
            d_248_v72_ = _dafny.Map({d_167_v0_: d_170_v3_})
            d_249_v73_: _dafny.MultiSet
            d_249_v73_ = _dafny.MultiSet([len(d_248_v72_), len(d_168_v1_)])
            d_250_v74_: D0
            d_251_v75_: bool
            out5_: D0
            out6_: bool
            out5_, out6_ = default__.m0(d_190_v19_, d_191_v20_, (d_249_v73_).issubset(_dafny.MultiSet([(0) - (len((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))])), default__.fm3(d_193_v22_, d_170_v3_, d_173_globalState_)])), d_191_v20_, d_173_globalState_)
            d_250_v74_ = out5_
            d_251_v75_ = out6_
            d_248_v72_ = (d_248_v72_).set(253, d_251_v75_)
            d_252_v76_: _dafny.Map
            d_252_v76_ = _dafny.Map({d_167_v0_: -98})
            d_253_v77_: D0
            d_254_v78_: bool
            out7_: D0
            out8_: bool
            out7_, out8_ = default__.m0(d_190_v19_, d_191_v20_, (d_170_v3_) and (d_193_v22_), default__.fm36(d_167_v0_, d_167_v0_, d_167_v0_, d_252_v76_, d_173_globalState_), d_173_globalState_)
            d_253_v77_ = out7_
            d_254_v78_ = out8_
        d_167_v0_ = (d_167_v0_ if d_170_v3_ else d_167_v0_)
        d_255_v79_: _dafny.Map
        d_255_v79_ = _dafny.Map({d_170_v3_: d_189_v18_})
        d_256_v80_: _dafny.Map
        d_256_v80_ = _dafny.Map({not (d_193_v22_) or (d_170_v3_): d_255_v79_})
        d_257_v81_: _dafny.Array
        nw30_ = _dafny.Array(int(0), 7)
        d_257_v81_ = nw30_
        d_258_v82_: _dafny.Seq
        d_258_v82_ = _dafny.SeqWithoutIsStrInference([d_257_v81_, d_257_v81_, d_257_v81_, d_257_v81_])
        d_259_v83_: C8
        nw31_ = C8()
        nw31_.ctor__(d_190_v19_, d_258_v82_, d_191_v20_, d_193_v22_)
        d_259_v83_ = nw31_
        d_260_v84_: _dafny.MultiSet
        d_260_v84_ = _dafny.MultiSet([d_170_v3_])
        d_261_v85_: D22
        d_261_v85_ = D22_DC48(d_259_v83_)
        rhs36_ = (d_256_v80_ if d_170_v3_ else d_256_v80_)
        rhs37_ = (((d_260_v84_).cardinality) - (d_167_v0_)) + (d_167_v0_)
        rhs38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibtmhviyp"))
        rhs39_ = (d_261_v85_).cf65
        d_256_v80_ = rhs36_
        d_167_v0_ = rhs37_
        d_196_v25_ = rhs38_
        d_259_v83_ = rhs39_
        d_262_v86_: _dafny.Map
        d_262_v86_ = _dafny.Map({(0) - (d_167_v0_): _dafny.SeqWithoutIsStrInference([598, d_167_v0_, d_167_v0_, d_167_v0_])})
        d_263_v87_: D19
        d_263_v87_ = D19_DC40(d_170_v3_, (d_259_v83_).fm8(d_193_v22_, d_170_v3_, d_193_v22_, d_173_globalState_), d_262_v86_)
        d_264_v88_: _dafny.Map
        d_264_v88_ = _dafny.Map({d_170_v3_: d_167_v0_})
        d_265_v89_: _dafny.MultiSet
        d_265_v89_ = _dafny.MultiSet([d_264_v88_])
        if ((d_263_v87_).cf52) >= ((d_265_v89_).cardinality):
            d_266_v90_: _dafny.Set
            d_266_v90_ = _dafny.Set({d_257_v81_, d_257_v81_})
            d_267_v91_: D22
            d_267_v91_ = D22_DC49(d_266_v90_)
            source7_ = d_267_v91_
            if source7_.is_DC49:
                d_268___mcc_h0_ = source7_.cf66
                d_269_cf66_ = d_268___mcc_h0_
                d_170_v3_ = d_193_v22_
                d_270_v92_: _dafny.Set
                d_270_v92_ = _dafny.Set({d_170_v3_})
                d_271_v93_: _dafny.Map
                d_271_v93_ = _dafny.Map({d_270_v92_: len((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))])})
                d_272_v94_: bool
                d_273_v95_: _dafny.Seq
                out9_: bool
                out10_: _dafny.Seq
                out9_, out10_ = (d_259_v83_).m2(d_193_v22_, (len(d_271_v93_)) + (d_167_v0_), d_196_v25_, d_173_globalState_)
                d_272_v94_ = out9_
                d_273_v95_ = out10_
                (d_173_globalState_).f1 = (d_167_v0_) + (d_167_v0_)
                (d_173_globalState_).f1 = d_167_v0_
            elif True:
                d_274___mcc_h1_ = source7_.cf65
                d_275_cf65_ = d_274___mcc_h1_
                d_276_v96_: _dafny.Set
                d_276_v96_ = _dafny.Set({((0) - (d_167_v0_)) - (d_167_v0_)})
                d_277_v97_: _dafny.MultiSet
                d_277_v97_ = _dafny.MultiSet([len(_dafny.Set({d_167_v0_}))])
                d_276_v96_ = _dafny.Set({((d_277_v97_).cardinality if d_170_v3_ else ((d_260_v84_)[d_193_v22_] if (d_193_v22_) in (d_260_v84_) else (d_259_v83_).fm8(d_170_v3_, d_170_v3_, d_193_v22_, d_173_globalState_))), d_167_v0_, d_167_v0_, (d_275_cf65_).fm8(d_170_v3_, d_170_v3_, d_170_v3_, d_173_globalState_)})
                (d_275_cf65_).m5(default__.fm1(d_167_v0_, d_170_v3_, d_167_v0_, d_173_globalState_), d_167_v0_, d_173_globalState_)
                (d_173_globalState_).f1 = (default__.safeDivisionInt(d_167_v0_, (0) - (d_167_v0_))) + (len(((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]) + (d_196_v25_)))
                d_278_v98_: C0
                nw32_ = C0()
                nw32_.ctor__(d_193_v22_, (d_169_v2_) + (d_169_v2_))
                d_278_v98_ = nw32_
            d_168_v1_ = d_168_v1_
            index40_ = default__.safeIndex(742, (d_257_v81_).length(0))
            (d_257_v81_)[index40_] = d_167_v0_
            d_279_v99_: _dafny.Map
            d_279_v99_ = _dafny.Map({d_167_v0_: d_193_v22_})
            d_280_v100_: _dafny.Set
            d_280_v100_ = _dafny.Set({d_193_v22_})
            d_281_v101_: C5
            nw33_ = C5()
            nw33_.ctor__(len(d_280_v100_), d_191_v20_, d_193_v22_)
            d_281_v101_ = nw33_
            d_282_v102_: _dafny.Seq
            d_282_v102_ = _dafny.SeqWithoutIsStrInference([d_281_v101_, d_281_v101_])
            d_283_v103_: _dafny.Map
            d_283_v103_ = _dafny.Map({((d_279_v99_)[d_167_v0_] if (d_167_v0_) in (d_279_v99_) else not(d_193_v22_)): (d_282_v102_).set(default__.safeIndex((0) - (d_167_v0_), len(d_282_v102_)), d_281_v101_)})
            index41_ = default__.safeIndex(742, (d_257_v81_).length(0))
            rhs40_ = (0) - (d_167_v0_)
            rhs41_ = (d_168_v1_) + ((d_168_v1_) + (d_168_v1_))
            rhs42_ = (d_283_v103_) | ((d_283_v103_) | ((d_283_v103_).set(d_193_v22_, d_282_v102_)))
            rhs43_ = d_167_v0_
            lhs31_ = d_257_v81_
            lhs32_ = default__.safeIndex(742, (d_257_v81_).length(0))
            lhs33_ = d_281_v101_
            lhs31_[lhs32_] = rhs40_
            d_168_v1_ = rhs41_
            d_283_v103_ = rhs42_
            lhs33_.f15 = rhs43_
            d_284_v104_: C0
            nw34_ = C0()
            nw34_.ctor__((642) > (((d_264_v88_)[not(d_193_v22_)] if (not(d_193_v22_)) in (d_264_v88_) else d_281_v101_.f15)), ((d_169_v2_) + (d_169_v2_)).set(default__.safeIndex(d_281_v101_.f15, len((d_169_v2_) + (d_169_v2_))), d_168_v1_))
            d_284_v104_ = nw34_
            index42_ = default__.safeIndex(731, (d_195_v24_).length(0))
            rhs44_ = d_284_v104_
            rhs45_ = (0) - (d_167_v0_)
            rhs46_ = d_196_v25_
            lhs34_ = d_173_globalState_
            lhs35_ = d_195_v24_
            lhs36_ = default__.safeIndex(731, (d_195_v24_).length(0))
            d_284_v104_ = rhs44_
            lhs34_.f1 = rhs45_
            lhs35_[lhs36_] = rhs46_
            d_285_v105_: _dafny.MultiSet
            d_285_v105_ = _dafny.MultiSet([d_191_v20_])
            d_286_v106_: _dafny.Array
            nw35_ = _dafny.Array(None, 29)
            nw35_[int(0)] = d_281_v101_.f15
            nw35_[int(1)] = (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))]
            nw35_[int(2)] = (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))]
            nw35_[int(3)] = d_281_v101_.f15
            nw35_[int(4)] = d_167_v0_
            nw35_[int(5)] = (0) - (len(d_168_v1_))
            nw35_[int(6)] = d_281_v101_.f15
            nw35_[int(7)] = d_281_v101_.f15
            nw35_[int(8)] = len((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))])
            nw35_[int(9)] = d_281_v101_.f15
            nw35_[int(10)] = d_167_v0_
            nw35_[int(11)] = len(_dafny.SeqWithoutIsStrInference([False]))
            nw35_[int(12)] = d_281_v101_.f15
            nw35_[int(13)] = ((d_285_v105_)[d_191_v20_] if (d_191_v20_) in (d_285_v105_) else 73)
            nw35_[int(14)] = d_281_v101_.f15
            nw35_[int(15)] = (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))]
            nw35_[int(16)] = d_281_v101_.f15
            nw35_[int(17)] = 864
            nw35_[int(18)] = 933
            nw35_[int(19)] = d_281_v101_.f15
            nw35_[int(20)] = len(d_279_v99_)
            nw35_[int(21)] = (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))]
            nw35_[int(22)] = len(_dafny.Set({len(d_168_v1_)}))
            nw35_[int(23)] = d_281_v101_.f15
            nw35_[int(24)] = len(d_168_v1_)
            nw35_[int(25)] = (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))]
            nw35_[int(26)] = d_281_v101_.f15
            nw35_[int(27)] = d_281_v101_.f15
            nw35_[int(28)] = (d_281_v101_).fm8(d_170_v3_, d_170_v3_, not(d_193_v22_), d_173_globalState_)
            d_286_v106_ = nw35_
            d_287_v107_: D2
            d_288_v108_: _dafny.MultiSet
            out11_: D2
            out12_: _dafny.MultiSet
            out11_, out12_ = (d_259_v83_).m4(d_286_v106_, d_172_v5_, (d_284_v104_).f11, (d_257_v81_)[default__.safeIndex(742, (d_257_v81_).length(0))], d_173_globalState_)
            d_287_v107_ = out11_
            d_288_v108_ = out12_
        elif True:
            d_289_v109_: _dafny.Set
            d_289_v109_ = _dafny.Set({d_259_v83_.f7})
            index43_ = default__.safeIndex(265, (d_190_v19_).length(0))
            (d_190_v19_)[index43_] = (d_259_v83_.f7) not in (d_289_v109_)
            index44_ = default__.safeIndex(265, (d_190_v19_).length(0))
            (d_190_v19_)[index44_] = d_170_v3_
            source8_ = d_192_v21_
            if source8_.is_DC1:
                (d_173_globalState_).f1 = (d_167_v0_) + ((0) - ((d_167_v0_) + ((0) - (d_167_v0_))))
                d_290_v110_: _dafny.Map
                d_290_v110_ = _dafny.Map({d_193_v22_: d_191_v20_})
                d_291_v111_: _dafny.Set
                d_291_v111_ = _dafny.Set({d_290_v110_})
                (d_259_v83_).m5(d_196_v25_, len(d_291_v111_), d_173_globalState_)
                d_292_v112_: _dafny.Map
                d_292_v112_ = _dafny.Map({d_170_v3_: d_193_v22_})
                d_293_v113_: _dafny.MultiSet
                d_293_v113_ = _dafny.MultiSet([d_292_v112_])
                (d_173_globalState_).f1 = ((d_293_v113_).cardinality) * (d_167_v0_)
                index45_ = default__.safeIndex(265, (d_190_v19_).length(0))
                (d_190_v19_)[index45_] = False
            elif source8_.is_DC2:
                d_294___mcc_h2_ = source8_.cf1
                d_295___mcc_h3_ = source8_.cf2
                d_296_cf2_ = d_295___mcc_h3_
                d_297_cf1_ = d_294___mcc_h2_
                d_298_v114_: C3
                nw36_ = C3()
                nw36_.ctor__((d_196_v25_) <= ((d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]), True)
                d_298_v114_ = nw36_
                d_299_v115_: _dafny.Set
                d_299_v115_ = _dafny.Set({d_170_v3_})
                d_300_v116_: _dafny.Map
                d_300_v116_ = _dafny.Map({d_167_v0_: d_299_v115_})
                d_301_v117_: _dafny.Map
                d_301_v117_ = _dafny.Map({_dafny.Set({d_193_v22_}): d_170_v3_})
                rhs47_ = default__.safeModuloInt(d_167_v0_, default__.safeDivisionInt(d_167_v0_, d_167_v0_))
                rhs48_ = default__.fm22((d_299_v115_) - ((((d_300_v116_)[d_167_v0_] if (d_167_v0_) in (d_300_v116_) else d_299_v115_))), default__.safeDivisionInt((d_259_v83_).fm7((d_298_v114_).f14, d_173_globalState_), d_167_v0_), (d_259_v83_).fm10(d_260_v84_, d_301_v117_, d_167_v0_, d_173_globalState_), (default__.fm48(d_167_v0_, not(True), d_173_globalState_)).ispropersubset(_dafny.MultiSet([True])), d_173_globalState_)
                rhs49_ = d_167_v0_
                lhs37_ = d_173_globalState_
                lhs38_ = d_173_globalState_
                lhs37_.f1 = rhs47_
                d_296_cf2_ = rhs48_
                lhs38_.f1 = rhs49_
                rhs50_ = d_170_v3_
                rhs51_ = (d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]
                rhs52_ = d_167_v0_
                lhs39_ = d_173_globalState_
                d_193_v22_ = rhs50_
                d_196_v25_ = rhs51_
                lhs39_.f1 = rhs52_
                d_302_v118_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.Array(None, 0), 1)
                d_302_v118_ = nw37_
                d_302_v118_ = d_302_v118_
            elif source8_.is_DC0:
                d_303___mcc_h4_ = source8_.cf0
                d_304_cf0_ = d_303___mcc_h4_
                d_305_v119_: T1
                nw38_ = C4()
                nw38_.ctor__()
                d_305_v119_ = nw38_
                d_306_v120_: _dafny.Seq
                d_306_v120_ = _dafny.SeqWithoutIsStrInference([d_305_v119_])
                d_307_v121_: _dafny.MultiSet
                d_307_v121_ = _dafny.MultiSet([(_dafny.MultiSet(d_172_v5_)).cardinality, d_167_v0_])
                d_308_v122_: _dafny.Set
                d_308_v122_ = _dafny.Set({d_167_v0_, (d_307_v121_).cardinality})
                d_309_v123_: _dafny.Array
                nw39_ = _dafny.Array(None, 21)
                nw39_[int(0)] = d_305_v119_
                nw39_[int(1)] = d_305_v119_
                nw39_[int(2)] = d_305_v119_
                nw39_[int(3)] = d_305_v119_
                nw39_[int(4)] = d_305_v119_
                nw39_[int(5)] = d_305_v119_
                nw39_[int(6)] = d_305_v119_
                nw39_[int(7)] = (d_306_v120_)[default__.safeIndex(len(d_308_v122_), len(d_306_v120_))]
                nw39_[int(8)] = d_305_v119_
                nw39_[int(9)] = d_305_v119_
                nw39_[int(10)] = d_305_v119_
                nw39_[int(11)] = d_305_v119_
                nw39_[int(12)] = d_305_v119_
                nw39_[int(13)] = d_305_v119_
                nw39_[int(14)] = d_305_v119_
                nw39_[int(15)] = d_305_v119_
                nw39_[int(16)] = d_305_v119_
                nw39_[int(17)] = d_305_v119_
                nw39_[int(18)] = d_305_v119_
                nw39_[int(19)] = d_305_v119_
                nw39_[int(20)] = d_305_v119_
                d_309_v123_ = nw39_
                index46_ = default__.safeIndex(899, (d_309_v123_).length(0))
                (d_309_v123_)[index46_] = (d_305_v119_ if d_170_v3_ else d_305_v119_)
                index47_ = default__.safeIndex(899, (d_309_v123_).length(0))
                nw40_ = C4()
                nw40_.ctor__()
                (d_309_v123_)[index47_] = nw40_
                d_310_v124_: _dafny.Seq
                d_311_v125_: bool
                d_312_v126_: bool
                out13_: _dafny.Seq
                out14_: bool
                out15_: bool
                out13_, out14_, out15_ = (d_259_v83_).m6(True, (d_259_v83_).fm10(d_260_v84_, _dafny.Map({_dafny.Set({d_193_v22_}): (D1_DC5(d_193_v22_)).cf5}), d_167_v0_, d_173_globalState_), d_173_globalState_)
                d_310_v124_ = out13_
                d_311_v125_ = out14_
                d_312_v126_ = out15_
                d_313_v127_: _dafny.MultiSet
                d_313_v127_ = _dafny.MultiSet([d_172_v5_, d_172_v5_])
                index48_ = default__.safeIndex(265, (d_190_v19_).length(0))
                (d_190_v19_)[index48_] = (d_172_v5_) in ((d_313_v127_).set(d_172_v5_, default__.abs(d_167_v0_)))
                index49_ = default__.safeIndex(614, (d_257_v81_).length(0))
                (d_257_v81_)[index49_] = d_167_v0_
                index50_ = default__.safeIndex(614, (d_257_v81_).length(0))
                (d_257_v81_)[index50_] = default__.safeModuloInt((110 if d_312_v126_ else 438), d_167_v0_)
            elif True:
                d_314___mcc_h5_ = source8_.cf3
                d_315_cf3_ = d_314___mcc_h5_
                d_316_v128_: _dafny.Map
                d_316_v128_ = _dafny.Map({d_257_v81_: d_170_v3_})
                d_317_v129_: C5
                nw41_ = C5()
                nw41_.ctor__(len(((d_316_v128_).set(d_257_v81_, d_193_v22_) if (d_190_v19_)[default__.safeIndex(265, (d_190_v19_).length(0))] else d_316_v128_)), (d_196_v25_)[default__.safeIndex((0) - (d_167_v0_), len(d_196_v25_))], not(d_170_v3_))
                d_317_v129_ = nw41_
                d_318_v130_: bool
                d_319_v131_: _dafny.Seq
                out16_: bool
                out17_: _dafny.Seq
                out16_, out17_ = (d_259_v83_).m2(d_193_v22_, d_317_v129_.f15, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylqpcjtj")), d_173_globalState_)
                d_318_v130_ = out16_
                d_319_v131_ = out17_
                index51_ = default__.safeIndex(265, (d_190_v19_).length(0))
                (d_190_v19_)[index51_] = d_170_v3_
                index52_ = default__.safeIndex(892, (d_257_v81_).length(0))
                (d_257_v81_)[index52_] = len((d_172_v5_) + (d_172_v5_))
                index53_ = default__.safeIndex(892, (d_257_v81_).length(0))
                (d_257_v81_)[index53_] = (d_167_v0_) * ((d_317_v129_.f15) * (d_167_v0_))
            d_191_v20_ = d_191_v20_
            nw42_ = C8()
            nw42_.ctor__(d_259_v83_.f7, d_258_v82_, d_191_v20_, d_170_v3_)
            d_259_v83_ = nw42_
            d_193_v22_ = (_dafny.CodePoint('f')) not in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asmrl"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjcwgjki"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asmrl")))), d_191_v20_))
        hi0_ = (d_167_v0_) - (d_167_v0_)
        for d_320_i4_ in range(d_167_v0_, hi0_):
            d_167_v0_ = d_167_v0_
            index54_ = default__.safeIndex(71, (d_190_v19_).length(0))
            (d_190_v19_)[index54_] = d_170_v3_
            d_321_v132_: _dafny.MultiSet
            d_321_v132_ = _dafny.MultiSet([(d_320_i4_) + ((d_260_v84_).cardinality)])
            d_322_v133_: T3
            nw43_ = C5()
            nw43_.ctor__((0) - (d_167_v0_), d_191_v20_, False)
            d_322_v133_ = nw43_
            d_323_v134_: _dafny.Map
            d_323_v134_ = _dafny.Map({d_320_i4_: d_322_v133_})
            d_324_v135_: _dafny.Set
            d_324_v135_ = _dafny.Set({d_322_v133_, d_322_v133_, d_322_v133_, d_322_v133_, ((d_323_v134_)[-341] if (-341) in (d_323_v134_) else d_322_v133_)})
            d_325_v136_: _dafny.Seq
            d_325_v136_ = _dafny.SeqWithoutIsStrInference([d_321_v132_, d_321_v132_])
            index55_ = default__.safeIndex(71, (d_190_v19_).length(0))
            rhs53_ = d_170_v3_
            rhs54_ = ((_dafny.MultiSet([(d_259_v83_).fm4(d_172_v5_, d_320_i4_, d_320_i4_, _dafny.MultiSet([d_167_v0_]), d_173_globalState_), d_320_i4_, d_167_v0_, (0) - (d_167_v0_), len(d_324_v135_)])) | (d_321_v132_)) | ((d_325_v136_)[default__.safeIndex(len(d_196_v25_), len(d_325_v136_))])
            rhs55_ = ((0) - ((0) - (default__.safeDivisionInt(983, len(_dafny.SeqWithoutIsStrInference([(d_322_v133_).f5 for d_326_i5_ in range(default__.abs(374))]))))) if (d_322_v133_).f6 else (d_322_v133_).fm7(d_193_v22_, d_173_globalState_))
            lhs40_ = d_190_v19_
            lhs41_ = default__.safeIndex(71, (d_190_v19_).length(0))
            lhs42_ = d_173_globalState_
            lhs40_[lhs41_] = rhs53_
            d_321_v132_ = rhs54_
            lhs42_.f1 = rhs55_
            d_327_v137_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.MultiSet({}), 27)
            d_327_v137_ = nw44_
            index56_ = default__.safeIndex(329, (d_327_v137_).length(0))
            (d_327_v137_)[index56_] = d_321_v132_
            index57_ = default__.safeIndex(329, (d_327_v137_).length(0))
            rhs56_ = (d_321_v132_) - (d_321_v132_)
            rhs57_ = len((default__.fm1(347, (d_190_v19_)[default__.safeIndex(71, (d_190_v19_).length(0))], d_167_v0_, d_173_globalState_)) + (d_196_v25_))
            rhs58_ = (d_259_v83_).fm4(d_172_v5_, d_320_i4_, d_167_v0_, d_321_v132_, d_173_globalState_)
            lhs43_ = d_327_v137_
            lhs44_ = default__.safeIndex(329, (d_327_v137_).length(0))
            lhs45_ = d_173_globalState_
            lhs43_[lhs44_] = rhs56_
            lhs45_.f1 = rhs57_
            d_167_v0_ = rhs58_
            d_328_v138_: _dafny.Map
            d_328_v138_ = _dafny.Map({d_320_i4_: d_320_i4_})
            d_329_v139_: _dafny.Map
            d_329_v139_ = _dafny.Map({d_328_v138_: d_172_v5_})
            index58_ = default__.safeIndex(439, (d_257_v81_).length(0))
            (d_257_v81_)[index58_] = 782
            d_330_v140_: C3
            nw45_ = C3()
            nw45_.ctor__(d_193_v22_, True)
            d_330_v140_ = nw45_
            d_331_v141_: _dafny.Map
            d_331_v141_ = _dafny.Map({d_171_v4_: d_330_v140_})
            d_332_v142_: _dafny.Map
            d_332_v142_ = d_329_v139_
            d_333_v143_: D20
            d_333_v143_ = D20_DC43(692, (d_190_v19_)[default__.safeIndex(71, (d_190_v19_).length(0))], d_167_v0_)
            index59_ = default__.safeIndex(439, (d_257_v81_).length(0))
            index60_ = default__.safeIndex(731, (d_195_v24_).length(0))
            rhs59_ = (0) - (d_320_i4_)
            rhs60_ = (d_167_v0_) != (len(d_331_v141_))
            rhs61_ = (d_332_v142_)
            rhs62_ = (d_333_v143_).cf56
            rhs63_ = _dafny.SeqWithoutIsStrInference([(d_322_v133_).f5 for d_334_i6_ in range(default__.abs(448))])
            lhs46_ = d_173_globalState_
            lhs47_ = d_257_v81_
            lhs48_ = default__.safeIndex(439, (d_257_v81_).length(0))
            lhs49_ = d_195_v24_
            lhs50_ = default__.safeIndex(731, (d_195_v24_).length(0))
            lhs46_.f1 = rhs59_
            d_193_v22_ = rhs60_
            d_329_v139_ = rhs61_
            lhs47_[lhs48_] = rhs62_
            lhs49_[lhs50_] = rhs63_
        if (d_167_v0_) > (d_167_v0_):
            d_335_v144_: C2
            nw46_ = C2()
            nw46_.ctor__(d_191_v20_, d_170_v3_)
            d_335_v144_ = nw46_
            d_335_v144_ = d_335_v144_
            (d_173_globalState_).f1 = (len(d_196_v25_)) - (d_167_v0_)
            d_336_v145_: _dafny.Seq
            d_336_v145_ = _dafny.SeqWithoutIsStrInference([(d_195_v24_)[default__.safeIndex(731, (d_195_v24_).length(0))]])
            index61_ = default__.safeIndex(793, (d_257_v81_).length(0))
            (d_257_v81_)[index61_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + ((d_336_v145_)[default__.safeIndex((d_335_v144_).fm7(d_193_v22_, d_173_globalState_), len(d_336_v145_))]))
            index62_ = default__.safeIndex(793, (d_257_v81_).length(0))
            rhs64_ = (0) - ((d_167_v0_) - (d_167_v0_))
            rhs65_ = d_167_v0_
            rhs66_ = (d_335_v144_).fm9(d_167_v0_, d_171_v4_, False, d_167_v0_, d_173_globalState_)
            rhs67_ = (d_167_v0_) - ((d_259_v83_).fm8(d_193_v22_, True, d_170_v3_, d_173_globalState_))
            lhs51_ = d_173_globalState_
            lhs52_ = d_257_v81_
            lhs53_ = default__.safeIndex(793, (d_257_v81_).length(0))
            lhs51_.f1 = rhs64_
            d_167_v0_ = rhs65_
            d_193_v22_ = rhs66_
            lhs52_[lhs53_] = rhs67_
            d_167_v0_ = 358
            d_193_v22_ = True
        elif True:
            (d_173_globalState_).f1 = d_167_v0_
            d_170_v3_ = d_193_v22_
            d_337_v146_: C3
            nw47_ = C3()
            nw47_.ctor__(not(d_170_v3_), (d_172_v5_) == (d_172_v5_))
            d_337_v146_ = nw47_
            d_193_v22_ = not((d_337_v146_).f14)
            d_195_v24_ = (d_195_v24_ if ((d_337_v146_).f13) and (d_193_v22_) else d_195_v24_)
        d_338_v147_: _dafny.Map
        d_338_v147_ = _dafny.Map({d_167_v0_: d_167_v0_})
        d_170_v3_ = (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_191_v20_ for d_339_i7_ in range(default__.abs(151))])), d_167_v0_)) <= ((0) - ((d_167_v0_) - (len(d_338_v147_))))
        d_340_v148_: _dafny.Map
        d_340_v148_ = _dafny.Map({d_193_v22_: d_170_v3_})
        d_341_v149_: C4
        nw48_ = C4()
        nw48_.ctor__()
        d_341_v149_ = nw48_
        d_342_v150_: _dafny.Seq
        d_342_v150_ = _dafny.SeqWithoutIsStrInference([d_341_v149_])
        d_343_v151_: _dafny.MultiSet
        d_343_v151_ = _dafny.MultiSet([829, d_167_v0_, len(d_342_v150_), d_167_v0_, d_167_v0_])
        d_344_i8_: int
        d_344_i8_ = 0
        with _dafny.label("3"):
            while ((d_259_v83_).fm4(d_172_v5_, len(d_340_v148_), (0) - (d_167_v0_), d_343_v151_, d_173_globalState_)) < (d_167_v0_):
                with _dafny.c_label("3"):
                    if (d_344_i8_) >= (100):
                        raise _dafny.Break("3")
                    d_344_i8_ = (d_344_i8_) + (1)
                    d_345_v152_: _dafny.Array
                    nw49_ = _dafny.Array(_dafny.Array(None, 0), 19)
                    d_345_v152_ = nw49_
                    index63_ = default__.safeIndex(363, (d_345_v152_).length(0))
                    (d_345_v152_)[index63_] = d_190_v19_
                    index64_ = default__.safeIndex(363, (d_345_v152_).length(0))
                    (d_345_v152_)[index64_] = d_259_v83_.f7
                    d_346_v153_: _dafny.Set
                    d_346_v153_ = _dafny.Set({d_167_v0_})
                    d_347_v154_: _dafny.Map
                    d_347_v154_ = _dafny.Map({d_346_v153_: ((d_338_v147_)[d_167_v0_] if (d_167_v0_) in (d_338_v147_) else d_167_v0_)})
                    d_348_v155_: _dafny.Set
                    d_348_v155_ = _dafny.Set({d_171_v4_})
                    (d_173_globalState_).f1 = ((d_260_v84_)[True] if (True) in (d_260_v84_) else (d_167_v0_) - (((d_347_v154_)[_dafny.Set({d_167_v0_, d_167_v0_})] if (_dafny.Set({d_167_v0_, d_167_v0_})) in (d_347_v154_) else len(d_348_v155_))))
                    d_349_v156_: D2
                    d_350_v157_: _dafny.MultiSet
                    out18_: D2
                    out19_: _dafny.MultiSet
                    out18_, out19_ = (d_259_v83_).m4(d_257_v81_, _dafny.SeqWithoutIsStrInference([d_193_v22_]), d_193_v22_, ((d_260_v84_)[not(False)] if (not(False)) in (d_260_v84_) else (d_259_v83_).fm4(d_172_v5_, d_167_v0_, 903, d_343_v151_, d_173_globalState_)), d_173_globalState_)
                    d_349_v156_ = out18_
                    d_350_v157_ = out19_
                    rhs68_ = (d_170_v3_ if default__.fm0((d_172_v5_)[default__.safeIndex(d_167_v0_, len(d_172_v5_))], d_173_globalState_) else (d_343_v151_).issubset(d_343_v151_))
                    rhs69_ = (_dafny.MultiSet(d_168_v1_)).intersection(d_343_v151_)
                    d_170_v3_ = rhs68_
                    d_343_v151_ = rhs69_
                    pass
            pass
        d_351_i9_: int
        d_351_i9_ = 0
        with _dafny.label("4"):
            while d_193_v22_:
                with _dafny.c_label("4"):
                    if (d_351_i9_) >= (100):
                        raise _dafny.Break("4")
                    d_351_i9_ = (d_351_i9_) + (1)
                    d_341_v149_ = d_341_v149_
                    (d_173_globalState_).f1 = (d_343_v151_).cardinality
                    (d_173_globalState_).f1 = (d_167_v0_) * (d_167_v0_)
                    index65_ = default__.safeIndex(458, (d_257_v81_).length(0))
                    (d_257_v81_)[index65_] = (default__.fm3(False, d_170_v3_, d_173_globalState_)) - ((0) - (d_167_v0_))
                    d_352_v158_: _dafny.Array
                    nw50_ = _dafny.Array(_dafny.MultiSet({}), 13)
                    d_352_v158_ = nw50_
                    d_353_v159_: D16
                    d_353_v159_ = D16_DC31(d_193_v22_)
                    d_354_v160_: _dafny.MultiSet
                    d_354_v160_ = _dafny.MultiSet([d_353_v159_])
                    index66_ = default__.safeIndex(729, (d_352_v158_).length(0))
                    (d_352_v158_)[index66_] = d_354_v160_
                    d_355_v161_: D25
                    d_355_v161_ = D25_DC52(_dafny.MultiSet([d_353_v159_]))
                    d_356_v162_: _dafny.Map
                    d_356_v162_ = _dafny.Map({d_167_v0_: (d_355_v161_).cf69})
                    index67_ = default__.safeIndex(458, (d_257_v81_).length(0))
                    index68_ = default__.safeIndex(729, (d_352_v158_).length(0))
                    rhs70_ = d_167_v0_
                    rhs71_ = (((d_356_v162_)[491] if (491) in (d_356_v162_) else d_354_v160_)).intersection((d_354_v160_).set(d_353_v159_, default__.abs(d_167_v0_)))
                    rhs72_ = d_190_v19_
                    rhs73_ = d_167_v0_
                    lhs54_ = d_257_v81_
                    lhs55_ = default__.safeIndex(458, (d_257_v81_).length(0))
                    lhs56_ = d_352_v158_
                    lhs57_ = default__.safeIndex(729, (d_352_v158_).length(0))
                    lhs58_ = d_259_v83_
                    lhs54_[lhs55_] = rhs70_
                    lhs56_[lhs57_] = rhs71_
                    lhs58_.f7 = rhs72_
                    d_167_v0_ = rhs73_
                    pass
            pass
        _dafny.print(_dafny.string_of(d_167_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_v1_) == (_dafny.SeqWithoutIsStrInference([647, 647, 647, 647, 647, 647]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v2_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([647, 647])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_170_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v4_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v5_) == (_dafny.SeqWithoutIsStrInference([True, True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_173_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_173_globalState_).f3) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([647, 647])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_globalState_.f4) == (_dafny.MultiSet([True, True, True, True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_174_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_189_v18_).cf3).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_190_v19_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_191_v20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_192_v21_).cf1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_192_v21_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_193_v22_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v23_) == (_dafny.Map({_dafny.CodePoint('j'): 647}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_195_v24_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_196_v25_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_197_v26_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_198_v27_)[0]).cf1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v27_)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_198_v27_)[1]).cf1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v27_)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_198_v27_)[2]).cf1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v27_)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_198_v27_)[3]).cf1)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_v27_)[3]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_199_v28_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_255_v79_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_256_v80_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v81_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_258_v82_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v83_.f7)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_259_v83_).f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v84_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65.f7)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_261_v85_).cf65).f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_261_v85_).cf65).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v86_) == (_dafny.Map({-1: _dafny.SeqWithoutIsStrInference([598, 1, 1, 1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v87_).cf51))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v87_).cf52))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_263_v87_).cf53) == (_dafny.Map({-1: _dafny.SeqWithoutIsStrInference([598, 1, 1, 1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v88_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v89_) == (_dafny.MultiSet([_dafny.Map({True: 1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_338_v147_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v148_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_342_v150_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_343_v151_) == (_dafny.MultiSet([829, 1, 1, 1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_344_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_351_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
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

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
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
        return lambda: D1_DC5(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC5(D1, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D2_DC7(int(0), int(0))
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

class D2_DC7(D2, NamedTuple('DC7', [('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf6.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC9(D2, NamedTuple('DC9', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC9({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC9) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)

class D3_DC10(D3, NamedTuple('DC10', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC12(D5, NamedTuple('DC12', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)

class D7_DC15(D7, NamedTuple('DC15', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D8_DC16)

class D8_DC16(D8, NamedTuple('DC16', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC16({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC16) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D9_DC17)

class D9_DC17(D9, NamedTuple('DC17', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC17({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC17) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D10_DC18)

class D10_DC18(D10, NamedTuple('DC18', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC18({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC18) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC20(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D11_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D11_DC19)

class D11_DC20(D11, NamedTuple('DC20', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC20({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC20) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC19(D11, NamedTuple('DC19', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC19({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC19) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC22(_dafny.Map({}), _dafny.Seq({}), _dafny.CodePoint('D'), False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D12_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D12_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D12_DC21)

class D12_DC22(D12, NamedTuple('DC22', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC22({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {self.cf28.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC22) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC23(D12, NamedTuple('DC23', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC23({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC23) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC21(D12, NamedTuple('DC21', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC21({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC21) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D13_DC24)

class D13_DC24(D13, NamedTuple('DC24', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC24({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC24) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC26(False, int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D14_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D14_DC25)

class D14_DC26(D14, NamedTuple('DC26', [('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC26({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC26) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC25(D14, NamedTuple('DC25', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC25({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC25) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC28(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D15_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D15_DC27)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D15_DC29)

class D15_DC28(D15, NamedTuple('DC28', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC28({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC28) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC27(D15, NamedTuple('DC27', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC27({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC27) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC29(D15, NamedTuple('DC29', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC29({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC29) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC31(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D16_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D16_DC32)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D16_DC30)

class D16_DC31(D16, NamedTuple('DC31', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC31({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC31) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC32(D16, NamedTuple('DC32', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC32({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC32) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC30(D16, NamedTuple('DC30', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC30({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC30) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC34(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D17_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D17_DC33)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D17_DC35)

class D17_DC34(D17, NamedTuple('DC34', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC34({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC34) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC33(D17, NamedTuple('DC33', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC33({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC33) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC35(D17, NamedTuple('DC35', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC35({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC35) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC37(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D18_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D18_DC36)

class D18_DC37(D18, NamedTuple('DC37', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC37({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC37) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC36(D18, NamedTuple('DC36', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC36({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC36) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC39(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D19_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D19_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D19_DC38)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D19_DC41)

class D19_DC39(D19, NamedTuple('DC39', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC39({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC39) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC40(D19, NamedTuple('DC40', [('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC40({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC40) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC38(D19, NamedTuple('DC38', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC38({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC38) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC41(D19, NamedTuple('DC41', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC41({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC41) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC43(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D20_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D20_DC42)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D20_DC44)

class D20_DC43(D20, NamedTuple('DC43', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC43({_dafny.string_of(self.cf56)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC43) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC42(D20, NamedTuple('DC42', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC42({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC42) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC44(D20, NamedTuple('DC44', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC44({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC44) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC46(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D21_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D21_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D21_DC45)

class D21_DC46(D21, NamedTuple('DC46', [('cf61', Any), ('cf62', Any), ('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC46({_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC46) and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC47(D21, NamedTuple('DC47', [('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC47({_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC47) and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC45(D21, NamedTuple('DC45', [('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC45({_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC45) and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC49(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D22_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D22_DC48)

class D22_DC49(D22, NamedTuple('DC49', [('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC49({_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC49) and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC48(D22, NamedTuple('DC48', [('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC48({_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC48) and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D23_DC50)

class D23_DC50(D23, NamedTuple('DC50', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC50({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC50) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D24_DC51)

class D24_DC51(D24, NamedTuple('DC51', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC51({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC51) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC53(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D25_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D25_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D25_DC52)

class D25_DC53(D25, NamedTuple('DC53', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC53({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC53) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC54(D25, NamedTuple('DC54', [('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC54({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC54) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC52(D25, NamedTuple('DC52', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC52({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC52) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC56(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D26_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D26_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D26_DC58)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D26_DC55)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D26_DC59)

class D26_DC56(D26, NamedTuple('DC56', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC56({_dafny.string_of(self.cf75)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC56) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC57(D26, NamedTuple('DC57', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC57({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC57) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC58(D26, NamedTuple('DC58', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC58({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC58) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC55(D26, NamedTuple('DC55', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC55({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC55) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC59(D26, NamedTuple('DC59', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC59({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC59) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, globalState):
        pass


class T2:
    pass
    def m2(self, p0, p1, p2, globalState):
        pass

    def m3(self, globalState):
        pass


class T3:
    pass

class T4:
    pass
    def m4(self, p0, p1, p2, p3, globalState):
        pass

    def m5(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: int = int(0)
        self.f4: _dafny.MultiSet = _dafny.MultiSet({})
        self._f0: str = _dafny.CodePoint('D')
        self._f2: int = int(0)
        self._f3: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3

class C0(T0):
    def  __init__(self):
        self.f12: _dafny.Seq = _dafny.Seq({})
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f11, f12):
        (self)._f11 = f11
        (self).f12 = f12

    def fm4(self, p0, p1, p2, p3, globalState):
        return len((_dafny.Map({-765: (self).f11})))

    def fm5(self, p0, globalState):
        return D1_DC4(default__.safeModuloInt(31, -303))

    @property
    def f11(self):
        return self._f11

class C1(T2, T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, globalState):
        if not(not(not(False))):
            return (_dafny.MultiSet([-16, 436, 81, (0) - (len(_dafny.SeqWithoutIsStrInference([645 for d_357_i0_ in range(default__.abs(824))]))), -807])).issubset(_dafny.MultiSet([484, 447]))
        elif True:
            return (True) == (not(not(False)))

    def fm7(self, p0, globalState):
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: D1
            for compr_21_ in (_dafny.SeqWithoutIsStrInference([D1_DC5(True)])).Elements:
                d_358_v0_: D1 = compr_21_
                if (d_358_v0_) in (_dafny.SeqWithoutIsStrInference([D1_DC5(True)])):
                    coll21_ = coll21_.union(_dafny.Set([d_358_v0_]))
            return _dafny.Set(coll21_)
        return default__.safeDivisionInt(-399, len(iife21_()
        ))

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: int
            for compr_22_ in (_dafny.Map({907: True})).keys.Elements:
                d_359_v0_: int = compr_22_
                if (d_359_v0_) in (_dafny.Map({907: True})):
                    coll22_ = coll22_.union(_dafny.Set([(d_359_v0_) - (85)]))
            return _dafny.Set(coll22_)
        return len(_dafny.SeqWithoutIsStrInference([not((_dafny.Set({len(iife22_()
        )})).issubset(_dafny.Set({len(_dafny.Map({46: not(True)})), 873})))]))

    def fm5(self, p0, globalState):
        return D1_DC4(-49)

    def fm14(self, globalState):
        return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aam"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_360_i0_ in range(default__.abs(143))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hteems")))))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r0 = (self).fm6(_dafny.SeqWithoutIsStrInference([-615 for d_361_i0_ in range(default__.abs(42))]), globalState)
        d_362_v0_: _dafny.Seq
        d_362_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
        d_362_v0_ = (_dafny.SeqWithoutIsStrInference([p0])) + (d_362_v0_)
        d_363_v1_: _dafny.MultiSet
        d_363_v1_ = _dafny.MultiSet([(self).fm7(p0, globalState)])
        rhs74_ = p0
        rhs75_ = (self).fm4(_dafny.SeqWithoutIsStrInference([p0]), 325, default__.safeModuloInt(p1, p1), d_363_v1_, globalState)
        lhs59_ = globalState
        r0 = rhs74_
        lhs59_.f1 = rhs75_
        d_364_v2_: _dafny.Array
        nw51_ = _dafny.Array(None, 9)
        nw51_[int(0)] = default__.fm0(p0, globalState)
        nw51_[int(1)] = not(p0)
        nw51_[int(2)] = not(not((self).fm6(_dafny.SeqWithoutIsStrInference([p1 for d_365_i1_ in range(default__.abs(578))]), globalState)))
        nw51_[int(3)] = p0
        nw51_[int(4)] = p0
        nw51_[int(5)] = p0
        nw51_[int(6)] = p0
        nw51_[int(7)] = p0
        nw51_[int(8)] = False
        d_364_v2_ = nw51_
        d_366_v3_: _dafny.MultiSet
        d_366_v3_ = _dafny.MultiSet([d_364_v2_])
        d_367_v4_: D0
        d_367_v4_ = D0_DC0(p0)
        pat_let_tv0_ = p2
        pat_let_tv1_ = p2
        pat_let_tv2_ = p2
        pat_let_tv3_ = p2
        pat_let_tv4_ = p2
        pat_let_tv5_ = p2
        def lambda13_(source9_):
            if source9_.is_DC1:
                return pat_let_tv0_
            elif source9_.is_DC2:
                d_368___mcc_h0_ = source9_.cf1
                d_369___mcc_h1_ = source9_.cf2
                d_370_cf2_ = d_369___mcc_h1_
                d_371_cf1_ = d_368___mcc_h0_
                return pat_let_tv1_
            elif source9_.is_DC0:
                d_372___mcc_h2_ = source9_.cf0
                d_373_cf0_ = d_372___mcc_h2_
                return (pat_let_tv2_) + (pat_let_tv3_)
            elif True:
                d_374___mcc_h3_ = source9_.cf3
                d_375_cf3_ = d_374___mcc_h3_
                return (pat_let_tv4_) + (pat_let_tv5_)

        rhs76_ = _dafny.MultiSet([d_364_v2_, d_364_v2_, d_364_v2_])
        rhs77_ = len(lambda13_(d_367_v4_))
        lhs60_ = globalState
        d_366_v3_ = rhs76_
        lhs60_.f1 = rhs77_
        d_376_v5_: _dafny.Array
        nw52_ = _dafny.Array(_dafny.Map({}), 23)
        d_376_v5_ = nw52_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_376_v5_).length(0)):
            d_377_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_377_i2_)) and ((d_377_i2_) < ((d_376_v5_).length(0)))):
                (d_376_v5_)[(d_377_i2_)] = (_dafny.Map({179: p1}) if p0 else _dafny.Map({p1: len(_dafny.Set({p0, p0}))}))
        d_378_v6_: _dafny.Array
        nw53_ = _dafny.Array(int(0), 21)
        d_378_v6_ = nw53_
        d_379_v7_: D2
        d_379_v7_ = D2_DC7(p1, (0) - ((self).fm14(globalState)))
        index69_ = default__.safeIndex(308, (d_378_v6_).length(0))
        (d_378_v6_)[index69_] = (p1) - ((d_379_v7_).cf7)
        index70_ = default__.safeIndex(308, (d_378_v6_).length(0))
        (d_378_v6_)[index70_] = (d_366_v3_).cardinality
        r0 = p0
        r1 = p2
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        d_380_v0_: bool
        d_380_v0_ = True
        d_381_v1_: _dafny.Seq
        d_381_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mof"))
        d_382_v2_: _dafny.Map
        d_382_v2_ = _dafny.Map({(not(d_380_v0_)) == (d_380_v0_): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (d_381_v1_)})
        d_382_v2_ = d_382_v2_
        d_383_v3_: int
        d_383_v3_ = -319
        d_384_v4_: _dafny.Map
        d_384_v4_ = _dafny.Map({d_383_v3_: d_383_v3_})
        d_385_v5_: _dafny.Map
        d_385_v5_ = _dafny.Map({d_380_v0_: d_384_v4_})
        d_386_v6_: _dafny.Map
        d_386_v6_ = _dafny.Map({(d_385_v5_) | (d_385_v5_): True})
        if ((d_386_v6_)[_dafny.Map({d_380_v0_: d_384_v4_})] if (_dafny.Map({d_380_v0_: d_384_v4_})) in (d_386_v6_) else not(d_380_v0_)):
            d_387_v7_: _dafny.MultiSet
            d_387_v7_ = _dafny.MultiSet([d_380_v0_])
            d_380_v0_ = (not(d_380_v0_) if (d_387_v7_).ispropersubset(d_387_v7_) else d_380_v0_)
            (globalState).f1 = d_383_v3_
            d_388_v8_: _dafny.Array
            nw54_ = _dafny.Array(int(0), 10)
            d_388_v8_ = nw54_
            rhs78_ = default__.safeModuloInt(d_383_v3_, d_383_v3_)
            rhs79_ = d_388_v8_
            d_383_v3_ = rhs78_
            d_388_v8_ = rhs79_
            d_389_v9_: _dafny.Seq
            d_389_v9_ = _dafny.SeqWithoutIsStrInference([d_380_v0_])
            if ((d_389_v9_) + (d_389_v9_)) != (d_389_v9_):
                d_390_v10_: str
                d_390_v10_ = _dafny.CodePoint('k')
                d_390_v10_ = d_390_v10_
                d_391_v11_: _dafny.Array
                nw55_ = _dafny.Array(False, 29)
                d_391_v11_ = nw55_
                index71_ = default__.safeIndex(4, (d_391_v11_).length(0))
                (d_391_v11_)[index71_] = d_380_v0_
                index72_ = default__.safeIndex(4, (d_391_v11_).length(0))
                (d_391_v11_)[index72_] = (d_389_v9_)[default__.safeIndex((0) - (d_383_v3_), len(d_389_v9_))]
                d_390_v10_ = d_390_v10_
                d_392_v12_: _dafny.Map
                d_392_v12_ = _dafny.Map({(d_391_v11_)[default__.safeIndex(4, (d_391_v11_).length(0))]: d_380_v0_})
                index73_ = default__.safeIndex(4, (d_391_v11_).length(0))
                (d_391_v11_)[index73_] = ((d_392_v12_)[False] if (False) in (d_392_v12_) else not(not((d_383_v3_) < (d_383_v3_))))
                index74_ = default__.safeIndex(349, (d_388_v8_).length(0))
                (d_388_v8_)[index74_] = d_383_v3_
                index75_ = default__.safeIndex(349, (d_388_v8_).length(0))
                (d_388_v8_)[index75_] = (0) - ((0) - (d_383_v3_))
            elif True:
                d_380_v0_ = not (d_380_v0_) or (d_380_v0_)
                d_393_v13_: _dafny.Array
                nw56_ = _dafny.Array(D2.default()(), 27)
                d_393_v13_ = nw56_
                d_393_v13_ = d_393_v13_
                d_380_v0_ = not(d_380_v0_)
                d_380_v0_ = not((d_380_v0_) and (d_380_v0_))
                d_394_v14_: str
                d_394_v14_ = _dafny.CodePoint('k')
                d_394_v14_ = d_394_v14_
            d_380_v0_ = True
        elif True:
            d_385_v5_ = d_385_v5_
            (globalState).f1 = (d_383_v3_) * (d_383_v3_)
            d_381_v1_ = d_381_v1_
            d_395_v15_: _dafny.Map
            d_395_v15_ = _dafny.Map({d_380_v0_: -385})
            d_396_v16_: _dafny.Seq
            d_396_v16_ = _dafny.SeqWithoutIsStrInference([-967, len(d_395_v15_)])
            d_397_v17_: _dafny.Map
            d_397_v17_ = _dafny.Map({(d_396_v16_).set(default__.safeIndex(d_383_v3_, len(d_396_v16_)), d_383_v3_): d_380_v0_})
            d_398_v18_: _dafny.Set
            d_398_v18_ = _dafny.Set({False, not(d_380_v0_)})
            d_399_v19_: _dafny.Map
            d_399_v19_ = _dafny.Map({(d_397_v17_) | (d_397_v17_): not((d_398_v18_).ispropersubset(d_398_v18_))})
            d_399_v19_ = (d_399_v19_).set((d_397_v17_) | (default__.fm15((0) - (d_383_v3_), globalState)), d_380_v0_)
            d_400_v20_: _dafny.Array
            nw57_ = _dafny.Array(int(0), 8)
            d_400_v20_ = nw57_
            d_401_v21_: _dafny.Seq
            d_401_v21_ = _dafny.SeqWithoutIsStrInference([d_396_v16_, d_396_v16_, _dafny.SeqWithoutIsStrInference([d_383_v3_ for d_402_i0_ in range(default__.abs(-573))])])
            d_403_v22_: T0
            nw58_ = C0()
            nw58_.ctor__(d_380_v0_, (d_401_v21_) + (d_401_v21_))
            d_403_v22_ = nw58_
            d_404_v23_: _dafny.Seq
            d_404_v23_ = _dafny.SeqWithoutIsStrInference([d_380_v0_])
            rhs80_ = d_380_v0_
            rhs81_ = d_400_v20_
            rhs82_ = (d_404_v23_)[default__.safeIndex(d_383_v3_, len(d_404_v23_))]
            rhs83_ = d_403_v22_
            d_380_v0_ = rhs80_
            d_400_v20_ = rhs81_
            d_380_v0_ = rhs82_
            d_403_v22_ = rhs83_
        d_405_v24_: _dafny.Seq
        d_405_v24_ = _dafny.SeqWithoutIsStrInference([d_380_v0_, d_380_v0_, (d_383_v3_) != (-789), (d_383_v3_) >= (d_383_v3_)])
        if (d_405_v24_)[default__.safeIndex((0) - (d_383_v3_), len(d_405_v24_))]:
            d_406_v25_: _dafny.Seq
            d_406_v25_ = _dafny.SeqWithoutIsStrInference([len(d_384_v4_), 547, 858, d_383_v3_, 288])
            d_407_v26_: _dafny.Seq
            d_407_v26_ = _dafny.SeqWithoutIsStrInference([d_406_v25_, d_406_v25_])
            d_408_v27_: C0
            nw59_ = C0()
            nw59_.ctor__(d_380_v0_, d_407_v26_)
            d_408_v27_ = nw59_
            d_408_v27_ = d_408_v27_
            d_409_v28_: _dafny.MultiSet
            d_409_v28_ = _dafny.MultiSet([d_380_v0_, d_380_v0_, not(d_380_v0_)])
            (globalState).f4 = d_409_v28_
            d_410_v29_: _dafny.Array
            nw60_ = _dafny.Array(False, 28)
            d_410_v29_ = nw60_
            index76_ = default__.safeIndex(915, (d_410_v29_).length(0))
            (d_410_v29_)[index76_] = ((d_408_v27_).f11 if d_380_v0_ else d_380_v0_)
            d_411_v30_: D1
            d_411_v30_ = D1_DC5((d_408_v27_).f11)
            index77_ = default__.safeIndex(915, (d_410_v29_).length(0))
            rhs84_ = (d_411_v30_).cf5
            rhs85_ = d_380_v0_
            rhs86_ = (d_408_v27_).f11
            lhs61_ = d_410_v29_
            lhs62_ = default__.safeIndex(915, (d_410_v29_).length(0))
            lhs61_[lhs62_] = rhs84_
            d_380_v0_ = rhs85_
            d_380_v0_ = rhs86_
            index78_ = default__.safeIndex(915, (d_410_v29_).length(0))
            (d_410_v29_)[index78_] = (self).fm6((d_408_v27_.f12)[default__.safeIndex(d_383_v3_, len(d_408_v27_.f12))], globalState)
            d_380_v0_ = d_380_v0_
        elif True:
            d_412_v31_: _dafny.Seq
            d_412_v31_ = _dafny.SeqWithoutIsStrInference([d_405_v24_])
            d_413_v32_: _dafny.Seq
            d_413_v32_ = _dafny.SeqWithoutIsStrInference([d_383_v3_])
            d_384_v4_ = (d_384_v4_).set((0) - (len(d_412_v31_)), (len(_dafny.SeqWithoutIsStrInference([d_413_v32_]))) + ((0) - (d_383_v3_)))
            d_414_v33_: _dafny.Array
            def lambda14_(d_415_v0_):
                def lambda15_(d_416_i1_):
                    return d_415_v0_

                return lambda15_

            init6_ = lambda14_(d_380_v0_)
            nw61_ = _dafny.Array(None, 7)
            for i0_6_ in range(nw61_.length(0)):
                nw61_[i0_6_] = init6_(i0_6_)
            d_414_v33_ = nw61_
            index79_ = default__.safeIndex(798, (d_414_v33_).length(0))
            (d_414_v33_)[index79_] = d_380_v0_
            index80_ = default__.safeIndex(798, (d_414_v33_).length(0))
            rhs87_ = not (not(d_380_v0_)) or (d_380_v0_)
            rhs88_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_417_i2_ in range(default__.abs(242))])) + (d_381_v1_)
            rhs89_ = d_381_v1_
            rhs90_ = 535
            lhs63_ = d_414_v33_
            lhs64_ = default__.safeIndex(798, (d_414_v33_).length(0))
            lhs63_[lhs64_] = rhs87_
            d_381_v1_ = rhs88_
            d_381_v1_ = rhs89_
            d_383_v3_ = rhs90_
            d_418_v34_: _dafny.Array
            nw62_ = _dafny.Array(int(0), 18)
            d_418_v34_ = nw62_
            index81_ = default__.safeIndex(456, (d_418_v34_).length(0))
            (d_418_v34_)[index81_] = (d_383_v3_) + (d_383_v3_)
            d_419_v35_: _dafny.Map
            d_419_v35_ = _dafny.Map({(d_414_v33_)[default__.safeIndex(798, (d_414_v33_).length(0))]: (0) - (d_383_v3_)})
            d_420_v36_: _dafny.MultiSet
            d_420_v36_ = _dafny.MultiSet([d_384_v4_, d_384_v4_, _dafny.Map({d_383_v3_: -861})])
            d_421_v37_: str
            d_421_v37_ = _dafny.CodePoint('e')
            d_422_v38_: _dafny.Map
            d_422_v38_ = _dafny.Map({d_421_v37_: d_383_v3_})
            index82_ = default__.safeIndex(456, (d_418_v34_).length(0))
            (d_418_v34_)[index82_] = ((d_419_v35_)[(d_420_v36_).isdisjoint(d_420_v36_)] if ((d_420_v36_).isdisjoint(d_420_v36_)) in (d_419_v35_) else (((d_422_v38_)[d_421_v37_] if (d_421_v37_) in (d_422_v38_) else d_383_v3_)) - (d_383_v3_))
            index83_ = default__.safeIndex(798, (d_414_v33_).length(0))
            (d_414_v33_)[index83_] = (d_414_v33_)[default__.safeIndex(798, (d_414_v33_).length(0))]
            d_423_v39_: D1
            d_423_v39_ = D1_DC5((d_414_v33_)[default__.safeIndex(798, (d_414_v33_).length(0))])
            d_424_v40_: _dafny.Map
            d_424_v40_ = _dafny.Map({d_383_v3_: d_423_v39_})
            d_425_v41_: _dafny.Map
            d_425_v41_ = _dafny.Map({(d_418_v34_)[default__.safeIndex(456, (d_418_v34_).length(0))]: d_381_v1_})
            d_424_v40_ = _dafny.Map({len(d_425_v41_): d_423_v39_})
        d_426_v42_: _dafny.Seq
        d_426_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-622 for d_427_i3_ in range(default__.abs(299))])])
        d_428_v43_: C0
        nw63_ = C0()
        nw63_.ctor__(d_380_v0_, (d_426_v42_) + (d_426_v42_))
        d_428_v43_ = nw63_
        d_380_v0_ = False
        d_429_i4_: int
        d_429_i4_ = 0
        with _dafny.label("5"):
            while not(default__.fm0(not((d_428_v43_).f11), globalState)):
                with _dafny.c_label("5"):
                    if (d_429_i4_) >= (100):
                        raise _dafny.Break("5")
                    d_429_i4_ = (d_429_i4_) + (1)
                    (globalState).f1 = d_383_v3_
                    d_380_v0_ = d_380_v0_
                    if (d_428_v43_).f11:
                        d_430_v44_: _dafny.Array
                        nw64_ = _dafny.Array(False, 14)
                        d_430_v44_ = nw64_
                        d_431_v45_: _dafny.MultiSet
                        d_431_v45_ = _dafny.MultiSet([d_430_v44_, d_430_v44_, d_430_v44_])
                        d_380_v0_ = ((d_431_v45_) | (d_431_v45_)).issubset((_dafny.MultiSet([d_430_v44_])) - (d_431_v45_))
                        index84_ = default__.safeIndex(684, (d_430_v44_).length(0))
                        (d_430_v44_)[index84_] = False
                        d_432_v46_: _dafny.Set
                        d_432_v46_ = _dafny.Set({(d_428_v43_).f11})
                        index85_ = default__.safeIndex(684, (d_430_v44_).length(0))
                        (d_430_v44_)[index85_] = ((d_432_v46_).isdisjoint(d_432_v46_) if (d_428_v43_).f11 else d_380_v0_)
                        rhs91_ = d_381_v1_
                        rhs92_ = d_430_v44_
                        rhs93_ = -241
                        lhs65_ = globalState
                        d_381_v1_ = rhs91_
                        d_430_v44_ = rhs92_
                        lhs65_.f1 = rhs93_
                        d_433_v47_: C0
                        nw65_ = C0()
                        nw65_.ctor__(d_380_v0_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_383_v3_, d_383_v3_]) for d_434_i5_ in range(default__.abs(437))]))
                        d_433_v47_ = nw65_
                        d_383_v3_ = d_383_v3_
                    elif True:
                        d_435_v48_: str
                        d_435_v48_ = _dafny.CodePoint('p')
                        d_436_v49_: _dafny.Map
                        d_436_v49_ = _dafny.Map({d_435_v48_: d_383_v3_})
                        (globalState).f1 = ((d_436_v49_)[d_435_v48_] if (d_435_v48_) in (d_436_v49_) else d_383_v3_)
                        d_435_v48_ = d_435_v48_
                        r0 = (self).fm7((d_428_v43_).f11, globalState)
                        r0 = (self).fm14(globalState)
                        d_437_v50_: _dafny.Array
                        nw66_ = _dafny.Array(False, 20)
                        d_437_v50_ = nw66_
                        d_438_v51_: _dafny.Map
                        d_438_v51_ = _dafny.Map({d_437_v50_: _dafny.CodePoint('w')})
                        d_439_v52_: D0
                        d_439_v52_ = D0_DC2(d_438_v51_, d_435_v48_)
                        d_439_v52_ = d_439_v52_
                    d_428_v43_ = d_428_v43_
                    pass
            pass
        r0 = d_383_v3_
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_440_v0_: bool
        d_440_v0_ = True
        d_440_v0_ = d_440_v0_
        d_441_v1_: _dafny.Seq
        d_441_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvlac"))
        d_441_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_442_i0_ in range(default__.abs(444))])
        if p2:
            d_443_v2_: int
            d_443_v2_ = 188
            d_444_v3_: _dafny.Set
            d_444_v3_ = _dafny.Set({d_443_v2_})
            d_445_v4_: _dafny.Seq
            d_445_v4_ = _dafny.SeqWithoutIsStrInference([p2, d_440_v0_, False])
            d_446_v5_: _dafny.Map
            d_446_v5_ = _dafny.Map({True: d_443_v2_})
            d_447_v6_: _dafny.Array
            def lambda16_(d_448_i1_):
                return True

            init7_ = lambda16_
            nw67_ = _dafny.Array(None, 22)
            for i0_7_ in range(nw67_.length(0)):
                nw67_[i0_7_] = init7_(i0_7_)
            d_447_v6_ = nw67_
            d_449_v7_: _dafny.Map
            d_449_v7_ = _dafny.Map({p1: d_447_v6_})
            d_450_v8_: _dafny.Map
            d_450_v8_ = _dafny.Map({len(d_446_v5_): d_449_v7_})
            d_451_v9_: _dafny.Map
            d_451_v9_ = _dafny.Map({d_443_v2_: d_443_v2_})
            d_452_v10_: _dafny.Seq
            d_452_v10_ = _dafny.SeqWithoutIsStrInference([d_445_v4_])
            d_453_v11_: _dafny.Array
            nw68_ = _dafny.Array(None, 22)
            nw68_[int(0)] = p0
            nw68_[int(1)] = False
            nw68_[int(2)] = p0
            nw68_[int(3)] = p0
            nw68_[int(4)] = p0
            nw68_[int(5)] = (d_444_v3_) == (_dafny.Set({d_443_v2_}))
            nw68_[int(6)] = p0
            nw68_[int(7)] = p2
            nw68_[int(8)] = p0
            nw68_[int(9)] = (default__.fm16(71, globalState)).isdisjoint(d_444_v3_)
            nw68_[int(10)] = d_440_v0_
            nw68_[int(11)] = not((d_440_v0_ if p2 else (d_445_v4_)[default__.safeIndex(d_443_v2_, len(d_445_v4_))]))
            nw68_[int(12)] = p0
            nw68_[int(13)] = d_440_v0_
            nw68_[int(14)] = (False) or (p2)
            nw68_[int(15)] = (p1) not in (((d_450_v8_)[((d_451_v9_)[d_443_v2_] if (d_443_v2_) in (d_451_v9_) else d_443_v2_)] if (((d_451_v9_)[d_443_v2_] if (d_443_v2_) in (d_451_v9_) else d_443_v2_)) in (d_450_v8_) else _dafny.Map({p1: d_447_v6_})))
            nw68_[int(16)] = not(p0)
            nw68_[int(17)] = d_440_v0_
            nw68_[int(18)] = default__.fm0(False, globalState)
            nw68_[int(19)] = (d_445_v4_) <= ((d_452_v10_)[default__.safeIndex(148, len(d_452_v10_))])
            nw68_[int(20)] = p0
            nw68_[int(21)] = p0
            d_453_v11_ = nw68_
            index86_ = default__.safeIndex(612, (d_453_v11_).length(0))
            (d_453_v11_)[index86_] = True
            index87_ = default__.safeIndex(612, (d_453_v11_).length(0))
            (d_453_v11_)[index87_] = d_440_v0_
            d_454_v12_: _dafny.Array
            nw69_ = _dafny.Array(int(0), 2)
            d_454_v12_ = nw69_
            index88_ = default__.safeIndex(511, (d_454_v12_).length(0))
            (d_454_v12_)[index88_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cf")))) * (len(_dafny.SeqWithoutIsStrInference([d_443_v2_, d_443_v2_, d_443_v2_])))
            index89_ = default__.safeIndex(511, (d_454_v12_).length(0))
            (d_454_v12_)[index89_] = d_443_v2_
            (globalState).f1 = (d_454_v12_)[default__.safeIndex(511, (d_454_v12_).length(0))]
            d_455_v13_: str
            d_455_v13_ = _dafny.CodePoint('k')
            d_441_v1_ = (((d_441_v1_) + (d_441_v1_)) + (d_441_v1_)).set(default__.safeIndex(d_443_v2_, len(((d_441_v1_) + (d_441_v1_)) + (d_441_v1_))), d_455_v13_)
            d_456_v14_: D1
            d_456_v14_ = D1_DC5(p2)
            d_457_v15_: _dafny.Map
            d_457_v15_ = _dafny.Map({d_456_v14_: p2})
            d_458_v16_: _dafny.MultiSet
            d_458_v16_ = _dafny.MultiSet([d_455_v13_, d_455_v13_, d_455_v13_])
            index90_ = default__.safeIndex(511, (d_454_v12_).length(0))
            index91_ = default__.safeIndex(612, (d_453_v11_).length(0))
            rhs94_ = ((d_458_v16_).cardinality) * ((d_454_v12_)[default__.safeIndex(511, (d_454_v12_).length(0))])
            rhs95_ = (0) - (default__.safeDivisionInt(d_443_v2_, ((d_454_v12_)[default__.safeIndex(511, (d_454_v12_).length(0))]) - (len(d_451_v9_))))
            rhs96_ = (d_457_v15_) | (_dafny.Map({d_456_v14_: p2}))
            rhs97_ = d_440_v0_
            lhs66_ = globalState
            lhs67_ = d_454_v12_
            lhs68_ = default__.safeIndex(511, (d_454_v12_).length(0))
            lhs69_ = d_453_v11_
            lhs70_ = default__.safeIndex(612, (d_453_v11_).length(0))
            lhs66_.f1 = rhs94_
            lhs67_[lhs68_] = rhs95_
            d_457_v15_ = rhs96_
            lhs69_[lhs70_] = rhs97_
        elif True:
            d_459_v17_: _dafny.MultiSet
            d_459_v17_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_460_i2_ in range(default__.abs(-754))])])
            d_440_v0_ = (d_459_v17_).issubset(d_459_v17_)
            d_461_v18_: _dafny.Set
            d_461_v18_ = _dafny.Set({d_440_v0_})
            d_462_v19_: _dafny.Array
            nw70_ = _dafny.Array(False, 3)
            d_462_v19_ = nw70_
            index92_ = default__.safeIndex(470, (d_462_v19_).length(0))
            (d_462_v19_)[index92_] = p0
            d_463_v20_: int
            d_463_v20_ = 192
            d_464_v21_: D2
            d_464_v21_ = D2_DC8(p0)
            index93_ = default__.safeIndex(470, (d_462_v19_).length(0))
            def iife23_(_pat_let0_0):
                def iife24_(d_465_dt__update__tmp_h0_):
                    def iife25_(_pat_let1_0):
                        def iife26_(d_466_dt__update_hcf9_h0_):
                            return D2_DC8(d_466_dt__update_hcf9_h0_)
                        return iife26_(_pat_let1_0)
                    return iife25_(True)
                return iife24_(_pat_let0_0)
            rhs98_ = not(d_440_v0_)
            rhs99_ = (d_461_v18_) - ((d_461_v18_) - (default__.fm17(d_463_v20_, d_440_v0_, globalState)))
            rhs100_ = not (d_440_v0_) or ((d_463_v20_) < (d_463_v20_))
            rhs101_ = ((p2 if (iife23_(d_464_v21_)).cf9 else p0) if d_440_v0_ else p2)
            rhs102_ = (0) - (default__.fm3(d_440_v0_, p2, globalState))
            lhs71_ = d_462_v19_
            lhs72_ = default__.safeIndex(470, (d_462_v19_).length(0))
            lhs73_ = globalState
            d_440_v0_ = rhs98_
            d_461_v18_ = rhs99_
            lhs71_[lhs72_] = rhs100_
            d_440_v0_ = rhs101_
            lhs73_.f1 = rhs102_
            d_467_v22_: _dafny.Map
            d_467_v22_ = _dafny.Map({d_463_v20_: d_463_v20_})
            d_468_v23_: str
            d_468_v23_ = _dafny.CodePoint('e')
            d_469_v24_: _dafny.Map
            d_469_v24_ = _dafny.Map({default__.fm1(len(d_441_v1_), False, d_463_v20_, globalState): (d_441_v1_).set(default__.safeIndex(d_463_v20_, len(d_441_v1_)), d_468_v23_)})
            d_467_v22_ = _dafny.Map({default__.safeModuloInt(d_463_v20_, len(((d_469_v24_)[d_441_v1_] if (d_441_v1_) in (d_469_v24_) else d_441_v1_))): d_463_v20_})
            d_470_v25_: _dafny.Map
            d_470_v25_ = _dafny.Map({211: p2})
            d_471_v26_: _dafny.Map
            d_471_v26_ = d_470_v25_
            d_472_v27_: _dafny.Array
            nw71_ = _dafny.Array(None, 28)
            nw71_[int(0)] = d_471_v26_
            nw71_[int(1)] = d_470_v25_
            nw71_[int(2)] = d_471_v26_
            nw71_[int(3)] = d_470_v25_
            nw71_[int(4)] = d_470_v25_
            nw71_[int(5)] = d_471_v26_
            nw71_[int(6)] = d_471_v26_
            nw71_[int(7)] = d_471_v26_
            nw71_[int(8)] = d_471_v26_
            nw71_[int(9)] = default__.fm18(d_463_v20_, globalState)
            nw71_[int(10)] = d_471_v26_
            nw71_[int(11)] = default__.fm18(d_463_v20_, globalState)
            nw71_[int(12)] = d_470_v25_
            nw71_[int(13)] = d_470_v25_
            nw71_[int(14)] = d_471_v26_
            nw71_[int(15)] = default__.fm18(len(d_441_v1_), globalState)
            nw71_[int(16)] = d_470_v25_
            nw71_[int(17)] = d_471_v26_
            nw71_[int(18)] = d_471_v26_
            nw71_[int(19)] = d_471_v26_
            nw71_[int(20)] = d_471_v26_
            nw71_[int(21)] = d_471_v26_
            nw71_[int(22)] = d_471_v26_
            nw71_[int(23)] = d_471_v26_
            nw71_[int(24)] = d_471_v26_
            nw71_[int(25)] = d_471_v26_
            nw71_[int(26)] = (d_471_v26_ if d_440_v0_ else d_471_v26_)
            nw71_[int(27)] = d_471_v26_
            d_472_v27_ = nw71_
            d_472_v27_ = d_472_v27_
            d_473_v28_: _dafny.Seq
            d_473_v28_ = _dafny.SeqWithoutIsStrInference([d_463_v20_, d_463_v20_, (d_463_v20_) - (default__.fm3(not(p2), d_440_v0_, globalState))])
            d_474_v29_: _dafny.Seq
            d_474_v29_ = _dafny.SeqWithoutIsStrInference([d_473_v28_, d_473_v28_])
            d_475_v30_: T0
            nw72_ = C0()
            nw72_.ctor__((self).fm6(d_473_v28_, globalState), d_474_v29_)
            d_475_v30_ = nw72_
            index94_ = default__.safeIndex(470, (d_462_v19_).length(0))
            index95_ = default__.safeIndex(470, (d_462_v19_).length(0))
            rhs103_ = d_440_v0_
            rhs104_ = ((d_473_v28_) + (d_473_v28_)).set(default__.safeIndex(d_463_v20_, len((d_473_v28_) + (d_473_v28_))), (0) - (len(d_441_v1_)))
            rhs105_ = d_475_v30_
            rhs106_ = (p1).issubset(p1)
            lhs74_ = d_462_v19_
            lhs75_ = default__.safeIndex(470, (d_462_v19_).length(0))
            lhs76_ = d_462_v19_
            lhs77_ = default__.safeIndex(470, (d_462_v19_).length(0))
            lhs74_[lhs75_] = rhs103_
            d_473_v28_ = rhs104_
            d_475_v30_ = rhs105_
            lhs76_[lhs77_] = rhs106_
        d_476_v31_: int
        d_476_v31_ = 174
        d_477_v32_: _dafny.Seq
        d_477_v32_ = _dafny.SeqWithoutIsStrInference([d_476_v31_, (p1).cardinality])
        d_478_v33_: _dafny.Seq
        d_478_v33_ = _dafny.SeqWithoutIsStrInference([d_477_v32_])
        d_479_v34_: C0
        nw73_ = C0()
        nw73_.ctor__(p0, d_478_v33_)
        d_479_v34_ = nw73_
        d_480_v35_: _dafny.Array
        nw74_ = _dafny.Array(_dafny.Array(None, 0), 12)
        d_480_v35_ = nw74_
        d_481_v36_: _dafny.Map
        d_481_v36_ = _dafny.Map({not(d_440_v0_): d_441_v1_})
        d_482_v37_: _dafny.Array
        nw75_ = _dafny.Array(None, 8)
        nw75_[int(0)] = d_441_v1_
        nw75_[int(1)] = ((d_481_v36_)[p2] if (p2) in (d_481_v36_) else d_441_v1_)
        nw75_[int(2)] = d_441_v1_
        nw75_[int(3)] = d_441_v1_
        nw75_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cucmgtuti"))
        nw75_[int(5)] = d_441_v1_
        nw75_[int(6)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_483_i3_ in range(default__.abs(717))]) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))
        nw75_[int(7)] = default__.fm1(d_476_v31_, p2, d_476_v31_, globalState)
        d_482_v37_ = nw75_
        index96_ = default__.safeIndex(549, (d_482_v37_).length(0))
        (d_482_v37_)[index96_] = d_441_v1_
        d_484_v38_: _dafny.Map
        d_484_v38_ = _dafny.Map({d_440_v0_: False})
        index97_ = default__.safeIndex(549, (d_482_v37_).length(0))
        rhs107_ = not((not ((d_479_v34_).f11) or (d_440_v0_)) not in (d_484_v38_))
        rhs108_ = d_480_v35_
        rhs109_ = d_441_v1_
        rhs110_ = (0) - (default__.safeDivisionInt(d_476_v31_, d_476_v31_))
        lhs78_ = d_482_v37_
        lhs79_ = default__.safeIndex(549, (d_482_v37_).length(0))
        d_440_v0_ = rhs107_
        d_480_v35_ = rhs108_
        lhs78_[lhs79_] = rhs109_
        d_476_v31_ = rhs110_
        d_485_v39_: _dafny.Array
        nw76_ = _dafny.Array(_dafny.Array(None, 0), 26)
        d_485_v39_ = nw76_
        d_486_v40_: _dafny.Array
        def lambda17_(d_487_i4_):
            return (d_487_i4_) - (410)

        init8_ = lambda17_
        nw77_ = _dafny.Array(None, 3)
        for i0_8_ in range(nw77_.length(0)):
            nw77_[i0_8_] = init8_(i0_8_)
        d_486_v40_ = nw77_
        index98_ = default__.safeIndex(840, (d_485_v39_).length(0))
        (d_485_v39_)[index98_] = d_486_v40_
        index99_ = default__.safeIndex(840, (d_485_v39_).length(0))
        def lambda18_(d_488_v31_):
            def lambda19_(d_489_i5_):
                return (d_489_i5_) - (d_488_v31_)

            return lambda19_

        init9_ = lambda18_(d_476_v31_)
        nw78_ = _dafny.Array(None, 26)
        for i0_9_ in range(nw78_.length(0)):
            nw78_[i0_9_] = init9_(i0_9_)
        (d_485_v39_)[index99_] = nw78_


class C2(T4, T3, T2, T1, T0):
    def  __init__(self):
        self._f5: str = _dafny.CodePoint('D')
        self._f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f5, f6):
        (self)._f5 = f5
        (self)._f6 = f6

    def fm9(self, p0, p1, p2, p3, globalState):
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlmbv"))))) != ((619) - (len(_dafny.SeqWithoutIsStrInference([721]))))

    def fm10(self, p0, p1, p2, globalState):
        return (self).f6

    def fm8(self, p0, p1, p2, globalState):
        return 292

    def fm6(self, p0, globalState):
        return (self).f6

    def fm7(self, p0, globalState):
        return (588) * (737)

    def fm4(self, p0, p1, p2, p3, globalState):
        return ((0) - ((212) - (663))) + (len((_dafny.Set({(self).f5})) | (_dafny.Set({(self).f5}))))

    def fm5(self, p0, globalState):
        source10_ = (_dafny.Map({len(_dafny.Map({63: len(_dafny.SeqWithoutIsStrInference([False]))})): (self).f6}) if True else _dafny.Map({136: (self).f6}))
        d_490___mcc_h0_ = source10_
        d_491_cf12_ = d_490___mcc_h0_
        return D1_DC4(len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])))

    def fm20(self, p0, p1, globalState):
        return (_dafny.MultiSet([(self).f6])) - (_dafny.MultiSet([(self).f6]))

    def fm21(self, globalState):
        return _dafny.Set({False})

    def m4(self, p0, p1, p2, p3, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_492_i0_: int
        d_492_i0_ = 0
        with _dafny.label("6"):
            while not(p2):
                with _dafny.c_label("6"):
                    if (d_492_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_492_i0_ = (d_492_i0_) + (1)
                    if (self).f6:
                        d_493_v0_: _dafny.Seq
                        d_493_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljybtapti"))
                        d_494_v1_: D2
                        d_494_v1_ = D2_DC8(p2)
                        d_495_v2_: _dafny.Set
                        d_495_v2_ = _dafny.Set({(self).f6})
                        d_496_v3_: _dafny.Array
                        nw79_ = _dafny.Array(None, 9)
                        nw79_[int(0)] = (self).f6
                        nw79_[int(1)] = ((self).f5) not in (d_493_v0_)
                        nw79_[int(2)] = p2
                        nw79_[int(3)] = (d_494_v1_).cf9
                        nw79_[int(4)] = (self).f6
                        nw79_[int(5)] = (p1)[default__.safeIndex(len(default__.fm2((self).f6, (self).f6, globalState)), len(p1))]
                        nw79_[int(6)] = (default__.fm22(d_495_v2_, 102, (self).f6, (self).f6, globalState)) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fllshammk")))
                        nw79_[int(7)] = not((self).f6)
                        nw79_[int(8)] = p2
                        d_496_v3_ = nw79_
                        index100_ = default__.safeIndex(178, (d_496_v3_).length(0))
                        (d_496_v3_)[index100_] = p2
                        d_497_v4_: _dafny.Map
                        d_497_v4_ = _dafny.Map({p2: p3})
                        d_498_v5_: _dafny.Seq
                        d_498_v5_ = _dafny.SeqWithoutIsStrInference([d_497_v4_])
                        index101_ = default__.safeIndex(178, (d_496_v3_).length(0))
                        (d_496_v3_)[index101_] = ((self).f6) not in ((d_498_v5_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcgthjfh")))), len(d_498_v5_))])
                        (globalState).f1 = p3
                        d_499_v6_: _dafny.Map
                        d_499_v6_ = _dafny.Map({p3: p0})
                        d_500_v7_: _dafny.Map
                        d_500_v7_ = _dafny.Map({((d_499_v6_)[p3] if (p3) in (d_499_v6_) else p0): (self).f6})
                        d_501_v8_: _dafny.Map
                        d_501_v8_ = _dafny.Map({d_495_v2_: (d_496_v3_)[default__.safeIndex(178, (d_496_v3_).length(0))]})
                        d_502_v9_: _dafny.MultiSet
                        d_502_v9_ = _dafny.MultiSet([p2, p2, (self).f6, (self).fm10(_dafny.MultiSet(p1), (d_501_v8_).set(d_495_v2_, p2), p3, globalState), (d_496_v3_)[default__.safeIndex(178, (d_496_v3_).length(0))]])
                        d_500_v7_ = (d_500_v7_).set(p0, (self).fm10(d_502_v9_, d_501_v8_, p3, globalState))
                        d_493_v0_ = (d_493_v0_) + (default__.fm1(228, (d_496_v3_)[default__.safeIndex(178, (d_496_v3_).length(0))], 343, globalState))
                        d_503_v10_: _dafny.Seq
                        d_503_v10_ = _dafny.SeqWithoutIsStrInference([d_496_v3_, d_496_v3_])
                        index102_ = default__.safeIndex(178, (d_496_v3_).length(0))
                        (d_496_v3_)[index102_] = (d_503_v10_) <= (_dafny.SeqWithoutIsStrInference([d_496_v3_, d_496_v3_]))
                    elif True:
                        d_504_v11_: _dafny.MultiSet
                        d_504_v11_ = _dafny.MultiSet([p3, 53, p3])
                        index103_ = default__.safeIndex(525, (p0).length(0))
                        (p0)[index103_] = (0) - (default__.safeDivisionInt(len(_dafny.Map({_dafny.Map({not(p2): (d_504_v11_).cardinality}): (self).f6})), p3))
                        index104_ = default__.safeIndex(525, (p0).length(0))
                        (p0)[index104_] = 325
                        index105_ = default__.safeIndex(525, (p0).length(0))
                        (p0)[index105_] = (p0)[default__.safeIndex(525, (p0).length(0))]
                        d_505_v12_: _dafny.Seq
                        d_505_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxu"))
                        d_506_v13_: D0
                        d_506_v13_ = D0_DC1()
                        d_507_v14_: _dafny.Array
                        nw80_ = _dafny.Array(None, 23)
                        nw80_[int(0)] = d_506_v13_
                        nw80_[int(1)] = d_506_v13_
                        nw80_[int(2)] = d_506_v13_
                        nw80_[int(3)] = d_506_v13_
                        nw80_[int(4)] = d_506_v13_
                        nw80_[int(5)] = d_506_v13_
                        nw80_[int(6)] = d_506_v13_
                        nw80_[int(7)] = d_506_v13_
                        nw80_[int(8)] = default__.fm23(globalState)
                        nw80_[int(9)] = d_506_v13_
                        nw80_[int(10)] = d_506_v13_
                        nw80_[int(11)] = (d_506_v13_ if not((self).f6) else d_506_v13_)
                        nw80_[int(12)] = d_506_v13_
                        nw80_[int(13)] = d_506_v13_
                        nw80_[int(14)] = d_506_v13_
                        nw80_[int(15)] = default__.fm23(globalState)
                        nw80_[int(16)] = d_506_v13_
                        nw80_[int(17)] = (d_506_v13_ if p2 else d_506_v13_)
                        nw80_[int(18)] = D0_DC1()
                        nw80_[int(19)] = d_506_v13_
                        nw80_[int(20)] = d_506_v13_
                        nw80_[int(21)] = d_506_v13_
                        nw80_[int(22)] = d_506_v13_
                        d_507_v14_ = nw80_
                        index106_ = default__.safeIndex(686, (d_507_v14_).length(0))
                        (d_507_v14_)[index106_] = d_506_v13_
                        d_508_v15_: bool
                        d_508_v15_ = True
                        index107_ = default__.safeIndex(686, (d_507_v14_).length(0))
                        index108_ = default__.safeIndex(525, (p0).length(0))
                        rhs111_ = d_505_v12_
                        rhs112_ = (p0)[default__.safeIndex(525, (p0).length(0))]
                        rhs113_ = d_506_v13_
                        rhs114_ = d_508_v15_
                        rhs115_ = (p3) - (p3)
                        lhs80_ = globalState
                        lhs81_ = d_507_v14_
                        lhs82_ = default__.safeIndex(686, (d_507_v14_).length(0))
                        lhs83_ = p0
                        lhs84_ = default__.safeIndex(525, (p0).length(0))
                        d_505_v12_ = rhs111_
                        lhs80_.f1 = rhs112_
                        lhs81_[lhs82_] = rhs113_
                        d_508_v15_ = rhs114_
                        lhs83_[lhs84_] = rhs115_
                        d_509_v16_: _dafny.Seq
                        d_509_v16_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukntgl")))])
                        d_510_v17_: _dafny.Seq
                        d_510_v17_ = _dafny.SeqWithoutIsStrInference([d_509_v16_])
                        d_511_v18_: C0
                        nw81_ = C0()
                        nw81_.ctor__(d_508_v15_, d_510_v17_)
                        d_511_v18_ = nw81_
                        d_508_v15_ = (d_511_v18_).f11
                    d_512_v19_: bool
                    d_512_v19_ = False
                    d_513_v20_: _dafny.Map
                    d_513_v20_ = _dafny.Map({d_512_v19_: not(d_512_v19_)})
                    d_514_v21_: _dafny.Map
                    d_514_v21_ = _dafny.Map({d_512_v19_: p3})
                    d_515_v22_: _dafny.MultiSet
                    d_515_v22_ = _dafny.MultiSet([p3])
                    d_516_v23_: _dafny.Seq
                    d_516_v23_ = _dafny.SeqWithoutIsStrInference([d_515_v22_])
                    d_517_v24_: _dafny.Seq
                    d_517_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvirq"))
                    d_518_v25_: D1
                    d_518_v25_ = D1_DC5(False)
                    d_519_v26_: _dafny.Seq
                    d_519_v26_ = _dafny.SeqWithoutIsStrInference([p3, p3, p3, p3])
                    d_520_v27_: _dafny.Seq
                    d_520_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p3 for d_521_i1_ in range(default__.abs(615))]), d_519_v26_])
                    d_522_v28_: C0
                    nw82_ = C0()
                    nw82_.ctor__((self).f6, d_520_v27_)
                    d_522_v28_ = nw82_
                    d_523_v29_: _dafny.Map
                    d_523_v29_ = _dafny.Map({p2: d_522_v28_})
                    d_524_v30_: _dafny.MultiSet
                    d_524_v30_ = _dafny.MultiSet([d_523_v29_])
                    d_525_v31_: _dafny.Array
                    nw83_ = _dafny.Array(False, 17)
                    d_525_v31_ = nw83_
                    d_526_v32_: _dafny.MultiSet
                    d_526_v32_ = _dafny.MultiSet([d_525_v31_, d_525_v31_])
                    d_527_v33_: _dafny.Array
                    nw84_ = _dafny.Array(None, 21)
                    nw84_[int(0)] = d_512_v19_
                    nw84_[int(1)] = ((d_513_v20_)[d_512_v19_] if (d_512_v19_) in (d_513_v20_) else p2)
                    nw84_[int(2)] = ((d_513_v20_)[(self).f6] if ((self).f6) in (d_513_v20_) else (self).f6)
                    nw84_[int(3)] = not (d_512_v19_) or (d_512_v19_)
                    nw84_[int(4)] = (len(p1)) >= (-89)
                    nw84_[int(5)] = False
                    nw84_[int(6)] = not(default__.fm0(p2, globalState))
                    nw84_[int(7)] = ((self).fm8(p2, (self).f6, p2, globalState)) > (p3)
                    nw84_[int(8)] = False
                    nw84_[int(9)] = (_dafny.Map({p2: p3})) == (d_514_v21_)
                    nw84_[int(10)] = p2
                    nw84_[int(11)] = (d_515_v22_).issubset(((d_516_v23_)[default__.safeIndex(len(d_517_v24_), len(d_516_v23_))]).set(p3, default__.abs(108)))
                    nw84_[int(12)] = (d_518_v25_).cf5
                    nw84_[int(13)] = (False if not(d_512_v19_) else (self).f6)
                    nw84_[int(14)] = (p3) <= (p3)
                    nw84_[int(15)] = (d_523_v29_) not in (d_524_v30_)
                    nw84_[int(16)] = d_512_v19_
                    nw84_[int(17)] = d_512_v19_
                    nw84_[int(18)] = (d_522_v28_).f11
                    nw84_[int(19)] = not(((d_526_v32_).cardinality) < (p3))
                    nw84_[int(20)] = ((d_515_v22_).set(p3, default__.abs(p3))).isdisjoint(d_515_v22_)
                    d_527_v33_ = nw84_
                    index109_ = default__.safeIndex(282, (d_527_v33_).length(0))
                    (d_527_v33_)[index109_] = (d_522_v28_).f11
                    index110_ = default__.safeIndex(282, (d_527_v33_).length(0))
                    rhs116_ = d_512_v19_
                    rhs117_ = ((d_512_v19_) and ((d_522_v28_).f11)) or (d_512_v19_)
                    rhs118_ = p3
                    rhs119_ = p2
                    rhs120_ = (len((d_517_v24_ if (self).f6 else d_517_v24_))) + (p3)
                    lhs85_ = d_527_v33_
                    lhs86_ = default__.safeIndex(282, (d_527_v33_).length(0))
                    lhs87_ = globalState
                    lhs88_ = globalState
                    d_512_v19_ = rhs116_
                    lhs85_[lhs86_] = rhs117_
                    lhs87_.f1 = rhs118_
                    d_512_v19_ = rhs119_
                    lhs88_.f1 = rhs120_
                    index111_ = default__.safeIndex(282, (d_527_v33_).length(0))
                    (d_527_v33_)[index111_] = (d_522_v28_).f11
                    if (self).f6:
                        (globalState).f1 = p3
                        d_528_v34_: C1
                        nw85_ = C1()
                        nw85_.ctor__()
                        d_528_v34_ = nw85_
                        (globalState).f1 = ((p3) + (len((_dafny.SeqWithoutIsStrInference([(self).f6])).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference([(self).f6]))), (self).f6)))) * (p3)
                        d_529_v35_: _dafny.Array
                        def lambda20_(d_530_v24_):
                            def lambda21_(d_531_i2_):
                                return d_530_v24_

                            return lambda21_

                        init10_ = lambda20_(d_517_v24_)
                        nw86_ = _dafny.Array(None, 14)
                        for i0_10_ in range(nw86_.length(0)):
                            nw86_[i0_10_] = init10_(i0_10_)
                        d_529_v35_ = nw86_
                        d_532_v36_: D0
                        d_532_v36_ = D0_DC0(d_512_v19_)
                        d_533_v37_: _dafny.Map
                        d_533_v37_ = _dafny.Map({d_532_v36_: d_512_v19_})
                        index112_ = default__.safeIndex(3, (d_529_v35_).length(0))
                        (d_529_v35_)[index112_] = (d_517_v24_ if ((d_533_v37_)[d_532_v36_] if (d_532_v36_) in (d_533_v37_) else (d_527_v33_)[default__.safeIndex(282, (d_527_v33_).length(0))]) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "btwvxfgd")))
                        index113_ = default__.safeIndex(3, (d_529_v35_).length(0))
                        (d_529_v35_)[index113_] = d_517_v24_
                        d_534_v38_: _dafny.Map
                        d_534_v38_ = _dafny.Map({p0: p3})
                        d_534_v38_ = d_534_v38_
                    elif True:
                        d_535_v39_: str
                        d_535_v39_ = _dafny.CodePoint('u')
                        d_535_v39_ = (self).f5
                        d_536_v40_: _dafny.Set
                        d_536_v40_ = _dafny.Set({(d_527_v33_)[default__.safeIndex(282, (d_527_v33_).length(0))]})
                        d_537_v41_: D0
                        d_538_v42_: bool
                        out20_: D0
                        out21_: bool
                        out20_, out21_ = default__.m0(d_527_v33_, default__.fm22(d_536_v40_, p3, p2, (d_522_v28_).f11, globalState), (d_522_v28_).f11, (self).f5, globalState)
                        d_537_v41_ = out20_
                        d_538_v42_ = out21_
                        d_539_v43_: _dafny.MultiSet
                        d_539_v43_ = _dafny.MultiSet([(d_522_v28_).f11, d_538_v42_])
                        d_540_v44_: _dafny.MultiSet
                        d_540_v44_ = _dafny.MultiSet([((d_539_v43_).set((d_522_v28_).f11, default__.abs(p3))).set((d_522_v28_).f11, default__.abs((self).fm4(p1, (d_539_v43_).cardinality, p3, d_515_v22_, globalState))), (self).fm20(_dafny.SeqWithoutIsStrInference([d_535_v39_ for d_541_i3_ in range(default__.abs(-7))]), False, globalState), d_539_v43_])
                        d_540_v44_ = ((d_540_v44_) - (d_540_v44_)) - ((d_540_v44_) | (d_540_v44_))
                        index114_ = default__.safeIndex(282, (d_527_v33_).length(0))
                        (d_527_v33_)[index114_] = True
                        d_542_v45_: C1
                        nw87_ = C1()
                        nw87_.ctor__()
                        d_542_v45_ = nw87_
                    pass
            pass
        d_543_v46_: _dafny.Array
        nw88_ = _dafny.Array(_dafny.MultiSet({}), 15)
        d_543_v46_ = nw88_
        d_544_v47_: _dafny.Array
        nw89_ = _dafny.Array(_dafny.CodePoint('D'), 21)
        d_544_v47_ = nw89_
        index115_ = default__.safeIndex(794, (d_543_v46_).length(0))
        (d_543_v46_)[index115_] = _dafny.MultiSet([d_544_v47_])
        d_545_v48_: _dafny.MultiSet
        d_545_v48_ = _dafny.MultiSet([d_544_v47_, d_544_v47_, d_544_v47_])
        index116_ = default__.safeIndex(794, (d_543_v46_).length(0))
        (d_543_v46_)[index116_] = ((_dafny.MultiSet([d_544_v47_])).set(d_544_v47_, default__.abs(default__.fm3(p2, p2, globalState)))) - (d_545_v48_)
        d_546_v49_: _dafny.Array
        nw90_ = _dafny.Array(D0.default()(), 6)
        d_546_v49_ = nw90_
        d_546_v49_ = d_546_v49_
        d_547_v50_: bool
        d_547_v50_ = False
        d_548_v51_: _dafny.Seq
        d_548_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlqrprnjt"))
        d_549_v52_: _dafny.Map
        d_549_v52_ = _dafny.Map({d_547_v50_: d_547_v50_})
        d_550_v53_: _dafny.Set
        d_550_v53_ = _dafny.Set({(0) - (p3)})
        rhs121_ = (d_549_v52_) != (d_549_v52_)
        rhs122_ = ((self).fm8((self).f6, p2, p2, globalState)) in ((d_550_v53_).intersection(_dafny.Set({p3})))
        rhs123_ = (d_548_v51_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hs")))
        rhs124_ = ((p3 if (self).f6 else p3)) <= ((0) - ((900 if (self).f6 else p3)))
        rhs125_ = (self).f6
        d_547_v50_ = rhs121_
        d_547_v50_ = rhs122_
        d_548_v51_ = rhs123_
        d_547_v50_ = rhs124_
        d_547_v50_ = rhs125_
        d_551_v54_: _dafny.Map
        d_551_v54_ = _dafny.Map({(self).fm20(d_548_v51_, d_547_v50_, globalState): (self).f6})
        d_552_v55_: _dafny.MultiSet
        d_552_v55_ = _dafny.MultiSet([(self).f6, (self).f6, p2])
        d_553_v56_: _dafny.Seq
        d_553_v56_ = _dafny.SeqWithoutIsStrInference([((d_552_v55_).set(default__.fm0(p2, globalState), default__.abs(p3))).set(d_547_v50_, default__.abs(p3)), d_552_v55_])
        d_551_v54_ = (d_551_v54_).set((d_553_v56_)[default__.safeIndex(p3, len(d_553_v56_))], True)
        d_547_v50_ = (p3) > (len(d_548_v51_))
        r0 = D2_DC8((default__.fm24(globalState)).cf0)
        r1 = (d_552_v55_) | (d_552_v55_)
        return r0, r1

    def m5(self, p0, p1, globalState):
        d_554_v0_: _dafny.Seq
        d_554_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        hi1_ = (d_554_v0_)[default__.safeIndex(p1, len(d_554_v0_))]
        for d_555_i0_ in range((0) - (p1), hi1_):
            if (self).f6:
                d_556_v1_: _dafny.Map
                d_556_v1_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqwmkiqrg")): d_555_i0_})
                d_556_v1_ = (d_556_v1_).set(p0, d_555_i0_)
                d_557_v2_: _dafny.Array
                def lambda22_(d_558_i1_):
                    return (self).f6

                init11_ = lambda22_
                nw91_ = _dafny.Array(None, 17)
                for i0_11_ in range(nw91_.length(0)):
                    nw91_[i0_11_] = init11_(i0_11_)
                d_557_v2_ = nw91_
                d_559_v3_: D0
                d_560_v4_: bool
                out22_: D0
                out23_: bool
                out22_, out23_ = default__.m0(d_557_v2_, (self).f5, (self).f6, (self).f5, globalState)
                d_559_v3_ = out22_
                d_560_v4_ = out23_
                d_561_v5_: _dafny.Map
                d_561_v5_ = _dafny.Map({(self).f6: d_555_i0_})
                d_562_v6_: D0
                d_562_v6_ = D0_DC0(not(d_560_v4_))
                d_560_v4_ = (self).fm9((d_555_i0_) * (len(d_561_v5_)), d_562_v6_, (self).f6, p1, globalState)
                d_563_v7_: D0
                d_564_v8_: bool
                out24_: D0
                out25_: bool
                out24_, out25_ = default__.m0(d_557_v2_, _dafny.CodePoint('y'), (d_560_v4_) in ((self).fm20(p0, not((self).f6), globalState)), (self).f5, globalState)
                d_563_v7_ = out24_
                d_564_v8_ = out25_
                d_565_v9_: _dafny.Map
                d_565_v9_ = _dafny.Map({d_555_i0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycs"))})
                d_566_v10_: _dafny.Seq
                d_566_v10_ = _dafny.SeqWithoutIsStrInference([(self).f6, d_564_v8_, (self).f6])
                d_567_v11_: _dafny.Array
                nw92_ = _dafny.Array(None, 29)
                nw92_[int(0)] = d_566_v10_
                nw92_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_564_v8_])) + (d_566_v10_)
                nw92_[int(2)] = d_566_v10_
                nw92_[int(3)] = (d_566_v10_).set(default__.safeIndex(d_555_i0_, len(d_566_v10_)), (self).f6)
                nw92_[int(4)] = d_566_v10_
                nw92_[int(5)] = d_566_v10_
                nw92_[int(6)] = _dafny.SeqWithoutIsStrInference([d_564_v8_])
                nw92_[int(7)] = d_566_v10_
                nw92_[int(8)] = d_566_v10_
                nw92_[int(9)] = d_566_v10_
                nw92_[int(10)] = d_566_v10_
                nw92_[int(11)] = d_566_v10_
                nw92_[int(12)] = d_566_v10_
                nw92_[int(13)] = d_566_v10_
                nw92_[int(14)] = d_566_v10_
                nw92_[int(15)] = d_566_v10_
                nw92_[int(16)] = d_566_v10_
                nw92_[int(17)] = d_566_v10_
                nw92_[int(18)] = (d_566_v10_) + (d_566_v10_)
                nw92_[int(19)] = d_566_v10_
                nw92_[int(20)] = (d_566_v10_) + (d_566_v10_)
                nw92_[int(21)] = d_566_v10_
                nw92_[int(22)] = d_566_v10_
                nw92_[int(23)] = (d_566_v10_) + (d_566_v10_)
                nw92_[int(24)] = (d_566_v10_) + (d_566_v10_)
                nw92_[int(25)] = d_566_v10_
                nw92_[int(26)] = _dafny.SeqWithoutIsStrInference([(self).f6])
                nw92_[int(27)] = d_566_v10_
                nw92_[int(28)] = (d_566_v10_) + (default__.fm2(d_560_v4_, d_560_v4_, globalState))
                d_567_v11_ = nw92_
                d_568_v12_: _dafny.Seq
                d_568_v12_ = _dafny.SeqWithoutIsStrInference([d_567_v11_])
                rhs126_ = d_565_v9_
                rhs127_ = (d_568_v12_)[default__.safeIndex(len(_dafny.Map({d_555_i0_: p1})), len(d_568_v12_))]
                rhs128_ = p1
                lhs89_ = globalState
                d_565_v9_ = rhs126_
                d_567_v11_ = rhs127_
                lhs89_.f1 = rhs128_
            elif True:
                d_569_v13_: _dafny.Array
                nw93_ = _dafny.Array(False, 12)
                d_569_v13_ = nw93_
                d_570_v14_: D0
                d_571_v15_: bool
                out26_: D0
                out27_: bool
                out26_, out27_ = default__.m0(d_569_v13_, (self).f5, (self).f6, _dafny.CodePoint('x'), globalState)
                d_570_v14_ = out26_
                d_571_v15_ = out27_
                d_572_v16_: _dafny.Seq
                d_572_v16_ = _dafny.SeqWithoutIsStrInference([d_569_v13_, d_569_v13_])
                d_573_v17_: _dafny.Map
                d_573_v17_ = _dafny.Map({p1: d_569_v13_})
                d_574_v18_: _dafny.MultiSet
                d_574_v18_ = _dafny.MultiSet([(d_572_v16_)[default__.safeIndex(len(d_554_v0_), len(d_572_v16_))], d_569_v13_, ((d_573_v17_)[d_555_i0_] if (d_555_i0_) in (d_573_v17_) else d_569_v13_)])
                d_575_v19_: _dafny.Set
                d_575_v19_ = _dafny.Set({False})
                d_574_v18_ = (d_574_v18_).set(d_569_v13_, default__.abs((0) - (len(d_575_v19_))))
                d_569_v13_ = d_569_v13_
                d_576_v20_: _dafny.Seq
                d_576_v20_ = _dafny.SeqWithoutIsStrInference([(d_555_i0_) > (d_555_i0_), (self).f6, d_571_v15_, not((self).f6), (self).f6])
                d_576_v20_ = (d_576_v20_) + (d_576_v20_)
                d_577_v21_: _dafny.Map
                d_577_v21_ = _dafny.Map({(False) == (True): len(p0)})
                d_577_v21_ = d_577_v21_
            d_578_v22_: str
            d_578_v22_ = _dafny.CodePoint('a')
            d_578_v22_ = d_578_v22_
            d_579_v23_: _dafny.Map
            d_579_v23_ = _dafny.Map({p1: (self).f6})
            d_580_v24_: _dafny.Map
            d_580_v24_ = _dafny.Map({len((d_554_v0_) + (_dafny.SeqWithoutIsStrInference([len(d_579_v23_), len(p0)]))): p0})
            d_580_v24_ = d_580_v24_
            d_581_v25_: _dafny.Array
            nw94_ = _dafny.Array(int(0), 1)
            d_581_v25_ = nw94_
            index117_ = default__.safeIndex(864, (d_581_v25_).length(0))
            (d_581_v25_)[index117_] = len(p0)
            index118_ = default__.safeIndex(864, (d_581_v25_).length(0))
            (d_581_v25_)[index118_] = (p1) + ((0) - (len(_dafny.SeqWithoutIsStrInference([d_578_v22_ for d_582_i2_ in range(default__.abs(-825))]))))
        d_583_v26_: bool
        d_583_v26_ = False
        d_583_v26_ = d_583_v26_
        d_584_v27_: _dafny.Map
        d_584_v27_ = _dafny.Map({((self).f6) or ((self).f6): (p1 if (self).f6 else p1)})
        d_585_v28_: _dafny.Set
        d_585_v28_ = _dafny.Set({p0})
        d_586_v29_: _dafny.Array
        def lambda23_(d_587_v26_):
            def lambda24_(d_588_i3_):
                return (d_588_i3_) + (len(_dafny.SeqWithoutIsStrInference([d_587_v26_, d_587_v26_, (self).f6])))

            return lambda24_

        init12_ = lambda23_(d_583_v26_)
        nw95_ = _dafny.Array(None, 14)
        for i0_12_ in range(nw95_.length(0)):
            nw95_[i0_12_] = init12_(i0_12_)
        d_586_v29_ = nw95_
        d_589_v30_: _dafny.Array
        d_589_v30_ = d_586_v29_
        d_590_v31_: _dafny.Map
        d_590_v31_ = _dafny.Map({d_589_v30_: p1})
        d_584_v27_ = (d_584_v27_).set((d_585_v28_).isdisjoint(d_585_v28_), len(d_590_v31_))
        d_591_v32_: _dafny.MultiSet
        d_591_v32_ = _dafny.MultiSet([d_586_v29_, d_586_v29_, d_586_v29_])
        d_592_i4_: int
        d_592_i4_ = 0
        with _dafny.label("7"):
            while ((d_591_v32_) - (d_591_v32_)).issubset(d_591_v32_):
                with _dafny.c_label("7"):
                    if (d_592_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_592_i4_ = (d_592_i4_) + (1)
                    d_593_v33_: _dafny.Set
                    d_593_v33_ = _dafny.Set({(self).f6})
                    d_594_v34_: _dafny.Array
                    nw96_ = _dafny.Array(None, 5)
                    nw96_[int(0)] = (d_593_v33_).intersection(d_593_v33_)
                    nw96_[int(1)] = (d_593_v33_).intersection(d_593_v33_)
                    nw96_[int(2)] = _dafny.Set({False, d_583_v26_})
                    nw96_[int(3)] = (d_593_v33_) - (d_593_v33_)
                    nw96_[int(4)] = (self).fm21(globalState)
                    d_594_v34_ = nw96_
                    index119_ = default__.safeIndex(143, (d_594_v34_).length(0))
                    (d_594_v34_)[index119_] = (d_593_v33_) - (d_593_v33_)
                    index120_ = default__.safeIndex(143, (d_594_v34_).length(0))
                    rhs129_ = (d_593_v33_) | (d_593_v33_)
                    rhs130_ = p1
                    lhs90_ = d_594_v34_
                    lhs91_ = default__.safeIndex(143, (d_594_v34_).length(0))
                    lhs92_ = globalState
                    lhs90_[lhs91_] = rhs129_
                    lhs92_.f1 = rhs130_
                    d_595_v35_: D0
                    d_595_v35_ = D0_DC0(d_583_v26_)
                    d_596_v36_: _dafny.Map
                    d_596_v36_ = _dafny.Map({p1: p1})
                    d_597_v37_: _dafny.Map
                    d_597_v37_ = _dafny.Map({default__.fm25((self).f6, p1, d_595_v35_, d_596_v36_, globalState): d_583_v26_})
                    d_598_v38_: _dafny.Map
                    d_598_v38_ = _dafny.Map({p1: d_554_v0_})
                    d_599_v39_: _dafny.MultiSet
                    d_599_v39_ = _dafny.MultiSet([d_583_v26_, (self).f6])
                    d_597_v37_ = (d_597_v37_).set((d_598_v38_) | (d_598_v38_), (p1) >= ((d_599_v39_).cardinality))
                    source11_ = d_586_v29_
                    d_600___mcc_h0_ = source11_
                    d_601_cf13_ = d_600___mcc_h0_
                    d_583_v26_ = d_583_v26_
                    d_602_v40_: _dafny.Map
                    d_602_v40_ = _dafny.Map({(0) - (p1): d_583_v26_})
                    (globalState).f1 = (0) - ((0) - (len((d_602_v40_) | (d_602_v40_))))
                    d_584_v27_ = (d_584_v27_).set((self).f6, default__.safeModuloInt(default__.fm3(d_583_v26_, default__.fm0(d_583_v26_, globalState), globalState), p1))
                    d_603_v43_: _dafny.Map
                    d_603_v43_ = _dafny.Map({_dafny.Set({True}): p1})
                    def iife27_():
                        def iife29_():
                            coll25_ = _dafny.Map()
                            compr_25_: _dafny.Set
                            for compr_25_ in (d_603_v43_).keys.Elements:
                                d_604_v42_: _dafny.Set = compr_25_
                                if (d_604_v42_) in (d_603_v43_):
                                    coll25_[d_604_v42_] = (self).f6
                            return _dafny.Map(coll25_)
                        coll23_ = _dafny.Map()
                        def iife28_():
                            coll24_ = _dafny.Map()
                            compr_24_: _dafny.Set
                            for compr_24_ in (d_603_v43_).keys.Elements:
                                d_604_v42_: _dafny.Set = compr_24_
                                if (d_604_v42_) in (d_603_v43_):
                                    coll24_[d_604_v42_] = (self).f6
                            return _dafny.Map(coll24_)
                        compr_23_: _dafny.Set
                        for compr_23_ in (iife28_()
                        ).keys.Elements:
                            d_605_v41_: _dafny.Set = compr_23_
                            if (d_605_v41_) in (iife29_()
                            ):
                                coll23_[d_605_v41_] = (self).f6
                        return _dafny.Map(coll23_)
                    d_583_v26_ = (self).fm10(d_599_v39_, iife27_()
                    , p1, globalState)
                    d_606_v45_: _dafny.Map
                    def iife30_():
                        coll26_ = _dafny.Set()
                        compr_26_: int
                        for compr_26_ in _dafny.IntegerRange(572, 521):
                            d_607_v44_: int = compr_26_
                            if ((572) <= (d_607_v44_)) and ((d_607_v44_) < (521)):
                                coll26_ = coll26_.union(_dafny.Set([(d_607_v44_) + (611)]))
                        return _dafny.Set(coll26_)
                    d_606_v45_ = _dafny.Map({iife30_()
                    : p1})
                    d_608_v46_: _dafny.Map
                    d_608_v46_ = _dafny.Map({(d_599_v39_).cardinality: (self).f6})
                    d_609_v47_: _dafny.Set
                    d_609_v47_ = _dafny.Set({len(_dafny.Map({976: d_608_v46_}))})
                    (globalState).f1 = (0) - (((d_606_v45_)[d_609_v47_] if (d_609_v47_) in (d_606_v45_) else (p1) + (p1)))
                    pass
            pass
        source12_ = d_589_v30_
        d_610___mcc_h1_ = source12_
        d_611_cf13_ = d_610___mcc_h1_
        d_612_v48_: _dafny.Seq
        d_612_v48_ = _dafny.SeqWithoutIsStrInference([d_554_v0_])
        d_613_v49_: C0
        nw97_ = C0()
        nw97_.ctor__(((self).f6) and (d_583_v26_), d_612_v48_)
        d_613_v49_ = nw97_
        d_614_v50_: _dafny.Map
        d_614_v50_ = _dafny.Map({p1: (self).f6})
        d_614_v50_ = _dafny.Map({p1: d_583_v26_})
        d_615_v51_: _dafny.Set
        d_615_v51_ = _dafny.Set({(self).f6, (d_613_v49_).f11})
        d_616_v52_: _dafny.Array
        nw98_ = _dafny.Array(None, 6)
        nw98_[int(0)] = d_583_v26_
        nw98_[int(1)] = not (d_583_v26_) or (not((d_613_v49_).f11))
        nw98_[int(2)] = (d_613_v49_).f11
        nw98_[int(3)] = True
        nw98_[int(4)] = not((_dafny.Set({(d_613_v49_).f11, d_583_v26_, (d_613_v49_).f11, (d_613_v49_).f11, (d_613_v49_).f11})).ispropersubset(d_615_v51_))
        nw98_[int(5)] = ((d_613_v49_).f11) or (False)
        d_616_v52_ = nw98_
        index121_ = default__.safeIndex(150, (d_616_v52_).length(0))
        (d_616_v52_)[index121_] = (d_613_v49_).f11
        d_617_v53_: _dafny.Array
        nw99_ = _dafny.Array(_dafny.CodePoint('D'), 28)
        d_617_v53_ = nw99_
        d_618_v54_: _dafny.Seq
        d_618_v54_ = _dafny.SeqWithoutIsStrInference([d_583_v26_, d_583_v26_, (self).f6])
        index122_ = default__.safeIndex(150, (d_616_v52_).length(0))
        rhs131_ = p1
        rhs132_ = (d_583_v26_ if (self).f6 else (p1) >= (len(d_618_v54_)))
        rhs133_ = (p1) < ((0) - (p1))
        rhs134_ = default__.fm3(default__.fm0(d_583_v26_, globalState), (d_583_v26_ if True else d_583_v26_), globalState)
        rhs135_ = d_617_v53_
        lhs93_ = globalState
        lhs94_ = d_616_v52_
        lhs95_ = default__.safeIndex(150, (d_616_v52_).length(0))
        lhs96_ = globalState
        lhs93_.f1 = rhs131_
        d_583_v26_ = rhs132_
        lhs94_[lhs95_] = rhs133_
        lhs96_.f1 = rhs134_
        d_617_v53_ = rhs135_
        d_619_v55_: _dafny.Map
        d_619_v55_ = _dafny.Map({d_583_v26_: d_586_v29_})
        d_620_v56_: _dafny.Array
        nw100_ = _dafny.Array(None, 20)
        nw100_[int(0)] = d_586_v29_
        nw100_[int(1)] = ((d_619_v55_)[(self).f6] if ((self).f6) in (d_619_v55_) else d_611_cf13_)
        nw100_[int(2)] = d_611_cf13_
        nw100_[int(3)] = d_586_v29_
        nw100_[int(4)] = (d_586_v29_ if (d_616_v52_)[default__.safeIndex(150, (d_616_v52_).length(0))] else d_586_v29_)
        nw100_[int(5)] = (d_589_v30_)
        nw100_[int(6)] = d_611_cf13_
        nw100_[int(7)] = d_611_cf13_
        nw100_[int(8)] = d_586_v29_
        nw100_[int(9)] = d_611_cf13_
        nw100_[int(10)] = d_611_cf13_
        nw100_[int(11)] = d_611_cf13_
        nw100_[int(12)] = d_586_v29_
        nw100_[int(13)] = d_611_cf13_
        nw100_[int(14)] = d_586_v29_
        nw100_[int(15)] = d_586_v29_
        nw100_[int(16)] = d_611_cf13_
        nw100_[int(17)] = d_611_cf13_
        nw100_[int(18)] = d_611_cf13_
        nw100_[int(19)] = d_611_cf13_
        d_620_v56_ = nw100_
        d_620_v56_ = d_620_v56_
        if d_583_v26_:
            d_583_v26_ = (d_583_v26_ if (self).f6 else not (not(True)) or ((self).f6))
            (globalState).f1 = p1
            d_621_v57_: _dafny.Set
            d_621_v57_ = _dafny.Set({p1, p1})
            d_583_v26_ = (d_583_v26_) or ((d_621_v57_).issubset(d_621_v57_))
            if (default__.fm26(globalState)).cf5:
                d_622_v58_: D2
                d_622_v58_ = D2_DC7(p1, p1)
                d_623_v59_: _dafny.Map
                d_623_v59_ = _dafny.Map({d_622_v58_: p0})
                d_623_v59_ = (d_623_v59_).set(d_622_v58_, p0)
                d_624_v60_: C0
                nw101_ = C0()
                nw101_.ctor__((self).f6, _dafny.SeqWithoutIsStrInference([d_554_v0_]))
                d_624_v60_ = nw101_
                d_624_v60_ = d_624_v60_
                d_625_v61_: _dafny.Array
                def lambda25_(d_626_v26_):
                    def lambda26_(d_627_i5_):
                        return d_626_v26_

                    return lambda26_

                init13_ = lambda25_(d_583_v26_)
                nw102_ = _dafny.Array(None, 21)
                for i0_13_ in range(nw102_.length(0)):
                    nw102_[i0_13_] = init13_(i0_13_)
                d_625_v61_ = nw102_
                nw103_ = _dafny.Array(False, 24)
                d_625_v61_ = nw103_
                d_628_v62_: _dafny.Seq
                d_628_v62_ = _dafny.SeqWithoutIsStrInference([not((p1) in (_dafny.Map({p1: p1}))), (self).f6, d_583_v26_])
                rhs136_ = (d_628_v62_) + (_dafny.SeqWithoutIsStrInference([(d_624_v60_).f11, (self).f6, d_583_v26_]))
                rhs137_ = p1
                rhs138_ = p1
                rhs139_ = ((d_554_v0_) + ((d_554_v0_) + (d_554_v0_))).set(default__.safeIndex((0) - (p1), len((d_554_v0_) + ((d_554_v0_) + (d_554_v0_)))), p1)
                lhs97_ = globalState
                lhs98_ = globalState
                d_628_v62_ = rhs136_
                lhs97_.f1 = rhs137_
                lhs98_.f1 = rhs138_
                d_554_v0_ = rhs139_
                d_629_v63_: _dafny.MultiSet
                d_629_v63_ = _dafny.MultiSet([(p1) + (-917)])
                d_629_v63_ = d_629_v63_
            elif True:
                (globalState).f1 = (p1) + (p1)
                d_630_v64_: _dafny.Array
                nw104_ = _dafny.Array(False, 27)
                d_630_v64_ = nw104_
                d_631_v65_: _dafny.Seq
                d_631_v65_ = _dafny.SeqWithoutIsStrInference([d_584_v27_])
                d_632_v66_: _dafny.MultiSet
                d_632_v66_ = _dafny.MultiSet([d_583_v26_])
                d_633_v67_: _dafny.Set
                d_633_v67_ = _dafny.Set({d_583_v26_, (self).f6})
                d_634_v68_: _dafny.Map
                d_634_v68_ = _dafny.Map({d_633_v67_: d_583_v26_})
                rhs140_ = d_630_v64_
                rhs141_ = (default__.safeModuloInt(len(p0), p1)) + (349)
                rhs142_ = (p1) + (len((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: p1}) for d_635_i6_ in range(default__.abs(24))])) + (d_631_v65_)))
                rhs143_ = (self).fm10(d_632_v66_, d_634_v68_, (p1) + (p1), globalState)
                lhs99_ = globalState
                lhs100_ = globalState
                d_630_v64_ = rhs140_
                lhs99_.f1 = rhs141_
                lhs100_.f1 = rhs142_
                d_583_v26_ = rhs143_
                (globalState).f1 = (0) - ((self).fm7(not((d_554_v0_) < (d_554_v0_)), globalState))
                d_583_v26_ = (self).f6
                d_636_v69_: C1
                nw105_ = C1()
                nw105_.ctor__()
                d_636_v69_ = nw105_
            index123_ = default__.safeIndex(21, (d_586_v29_).length(0))
            (d_586_v29_)[index123_] = p1
            index124_ = default__.safeIndex(21, (d_586_v29_).length(0))
            (d_586_v29_)[index124_] = (p1) + (p1)
        elif True:
            d_637_v70_: _dafny.Array
            nw106_ = _dafny.Array(False, 6)
            d_637_v70_ = nw106_
            d_638_v71_: D0
            d_638_v71_ = D0_DC0(False)
            d_639_v72_: D0
            d_639_v72_ = D0_DC3(d_638_v71_)
            rhs144_ = d_637_v70_
            rhs145_ = (p1) >= ((d_554_v0_)[default__.safeIndex(p1, len(d_554_v0_))])
            rhs146_ = D0_DC3(d_638_v71_)
            d_637_v70_ = rhs144_
            d_583_v26_ = rhs145_
            d_639_v72_ = rhs146_
            d_640_v73_: _dafny.Array
            def lambda27_(d_641_v26_, d_642_p0_):
                def lambda28_(d_643_i7_):
                    return (d_642_p0_ if d_641_v26_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqdfgekt")))

                return lambda28_

            init14_ = lambda27_(d_583_v26_, p0)
            nw107_ = _dafny.Array(None, 2)
            for i0_14_ in range(nw107_.length(0)):
                nw107_[i0_14_] = init14_(i0_14_)
            d_640_v73_ = nw107_
            rhs147_ = (0) - (p1)
            rhs148_ = ((p1) * (p1)) - ((len(p0)) + (p1))
            rhs149_ = ((p1) * (p1)) * (default__.safeDivisionInt(-35, p1))
            rhs150_ = (d_640_v73_ if not((self).f6) else (d_640_v73_ if (self).f6 else d_640_v73_))
            lhs101_ = globalState
            lhs102_ = globalState
            lhs103_ = globalState
            lhs101_.f1 = rhs147_
            lhs102_.f1 = rhs148_
            lhs103_.f1 = rhs149_
            d_640_v73_ = rhs150_
            d_644_v74_: _dafny.MultiSet
            d_644_v74_ = _dafny.MultiSet([(self).f6])
            d_645_v75_: _dafny.Set
            d_645_v75_ = _dafny.Set({True, (self).f6, d_583_v26_, d_583_v26_})
            d_646_v76_: _dafny.Map
            d_646_v76_ = _dafny.Map({d_645_v75_: True})
            d_647_v77_: _dafny.Seq
            d_647_v77_ = _dafny.SeqWithoutIsStrInference([(False) or ((self).fm10(d_644_v74_, d_646_v76_, p1, globalState)), True, d_583_v26_])
            if (d_647_v77_)[default__.safeIndex(default__.safeDivisionInt(p1, len(d_554_v0_)), len(d_647_v77_))]:
                d_648_v78_: _dafny.Seq
                d_648_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oomt"))
                d_648_v78_ = p0
                d_649_v80_: _dafny.Set
                d_649_v80_ = _dafny.Set({p1})
                def iife31_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in (_dafny.SeqWithoutIsStrInference([29 for d_650_i8_ in range(default__.abs(-802))])).Elements:
                        d_651_v79_: int = compr_27_
                        if (d_651_v79_) in (_dafny.SeqWithoutIsStrInference([29 for d_650_i8_ in range(default__.abs(-802))])):
                            coll27_ = coll27_.union(_dafny.Set([(d_651_v79_) - (-659)]))
                    return _dafny.Set(coll27_)
                d_583_v26_ = (d_649_v80_).issubset(iife31_()
                )
                d_652_v81_: str
                d_652_v81_ = _dafny.CodePoint('v')
                d_653_v82_: D1
                d_653_v82_ = D1_DC4(683)
                rhs151_ = d_637_v70_
                rhs152_ = d_637_v70_
                rhs153_ = _dafny.CodePoint('k')
                rhs154_ = (d_653_v82_).cf4
                lhs104_ = globalState
                d_637_v70_ = rhs151_
                d_637_v70_ = rhs152_
                d_652_v81_ = rhs153_
                lhs104_.f1 = rhs154_
                index125_ = default__.safeIndex(633, (d_637_v70_).length(0))
                (d_637_v70_)[index125_] = (d_583_v26_) == (d_583_v26_)
                index126_ = default__.safeIndex(633, (d_637_v70_).length(0))
                (d_637_v70_)[index126_] = True
                d_583_v26_ = ((self).fm21(globalState)).isdisjoint(_dafny.Set({(self).f6, (self).f6}))
            elif True:
                d_654_v83_: C1
                nw108_ = C1()
                nw108_.ctor__()
                d_654_v83_ = nw108_
                d_655_v84_: _dafny.Map
                d_655_v84_ = _dafny.Map({d_583_v26_: d_654_v83_})
                d_655_v84_ = d_655_v84_
                d_656_v85_: _dafny.Array
                nw109_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_656_v85_ = nw109_
                d_656_v85_ = d_656_v85_
                (globalState).f1 = 961
                d_583_v26_ = (self).f6
                d_657_v86_: C1
                nw110_ = C1()
                nw110_.ctor__()
                d_657_v86_ = nw110_
            (globalState).f1 = p1
            (globalState).f1 = p1

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r0 = (p2) != (p2)
        d_658_v0_: _dafny.Array
        nw111_ = _dafny.Array(int(0), 5)
        d_658_v0_ = nw111_
        index127_ = default__.safeIndex(165, (d_658_v0_).length(0))
        (d_658_v0_)[index127_] = p1
        index128_ = default__.safeIndex(165, (d_658_v0_).length(0))
        (d_658_v0_)[index128_] = default__.safeDivisionInt(p1, p1)
        d_659_v1_: _dafny.Seq
        d_659_v1_ = _dafny.SeqWithoutIsStrInference([p0, (self).f6])
        d_660_v2_: D1
        d_660_v2_ = D1_DC4(len(d_659_v1_))
        pat_let_tv8_ = d_658_v0_
        pat_let_tv9_ = d_658_v0_
        def iife36_(_pat_let4_0):
            def iife37_(d_670_dt__update__tmp_h0_):
                def iife38_(_pat_let5_0):
                    def iife39_(d_671_dt__update_hcf4_h0_):
                        return D1_DC4(d_671_dt__update_hcf4_h0_)
                    return iife39_(_pat_let5_0)
                return iife38_((pat_let_tv9_)[default__.safeIndex(165, (pat_let_tv8_).length(0))])
            return iife37_(_pat_let4_0)
        hi2_ = p1
        for d_661_i0_ in range((iife36_(d_660_v2_)).cf4, hi2_):
            d_662_v3_: _dafny.Array
            nw112_ = _dafny.Array(False, 16)
            d_662_v3_ = nw112_
            index129_ = default__.safeIndex(660, (d_662_v3_).length(0))
            (d_662_v3_)[index129_] = (self).f6
            index130_ = default__.safeIndex(660, (d_662_v3_).length(0))
            (d_662_v3_)[index130_] = (self).f6
            r0 = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocyq"))) + (p2))) >= ((p1) - ((0) - ((0) - (d_661_i0_))))
            d_663_v4_: _dafny.Seq
            d_663_v4_ = _dafny.SeqWithoutIsStrInference([(d_658_v0_)[default__.safeIndex(165, (d_658_v0_).length(0))]])
            d_664_v5_: D6
            d_664_v5_ = D6_DC13(d_663_v4_)
            r0 = ((d_664_v5_).cf14) != (d_663_v4_)
            d_665_v6_: _dafny.Array
            nw113_ = _dafny.Array(D0.default()(), 4)
            d_665_v6_ = nw113_
            d_666_v7_: _dafny.Map
            d_666_v7_ = _dafny.Map({d_662_v3_: (self).f5})
            d_667_v8_: D0
            d_667_v8_ = D0_DC2(d_666_v7_, (self).f5)
            index131_ = default__.safeIndex(868, (d_665_v6_).length(0))
            (d_665_v6_)[index131_] = (d_667_v8_ if p0 else d_667_v8_)
            pat_let_tv6_ = d_666_v7_
            pat_let_tv7_ = d_666_v7_
            index132_ = default__.safeIndex(165, (d_658_v0_).length(0))
            index133_ = default__.safeIndex(660, (d_662_v3_).length(0))
            index134_ = default__.safeIndex(660, (d_662_v3_).length(0))
            index135_ = default__.safeIndex(868, (d_665_v6_).length(0))
            def iife32_(_pat_let2_0):
                def iife33_(d_668_dt__update__tmp_h1_):
                    def iife34_(_pat_let3_0):
                        def iife35_(d_669_dt__update_hcf1_h0_):
                            return D0_DC2(d_669_dt__update_hcf1_h0_, (d_668_dt__update__tmp_h1_).cf2)
                        return iife35_(_pat_let3_0)
                    return iife34_((pat_let_tv6_) | (pat_let_tv7_))
                return iife33_(_pat_let2_0)
            rhs155_ = (p1) + ((d_658_v0_)[default__.safeIndex(165, (d_658_v0_).length(0))])
            rhs156_ = not((not((d_661_i0_) != (len(d_659_v1_))) if (self).f6 else ((d_662_v3_)[default__.safeIndex(660, (d_662_v3_).length(0))] if (self).f6 else (self).f6)))
            rhs157_ = (d_661_i0_) == (p1)
            rhs158_ = iife32_(d_667_v8_)
            rhs159_ = (0) - ((0) - (p1))
            lhs105_ = d_658_v0_
            lhs106_ = default__.safeIndex(165, (d_658_v0_).length(0))
            lhs107_ = d_662_v3_
            lhs108_ = default__.safeIndex(660, (d_662_v3_).length(0))
            lhs109_ = d_662_v3_
            lhs110_ = default__.safeIndex(660, (d_662_v3_).length(0))
            lhs111_ = d_665_v6_
            lhs112_ = default__.safeIndex(868, (d_665_v6_).length(0))
            lhs113_ = globalState
            lhs105_[lhs106_] = rhs155_
            lhs107_[lhs108_] = rhs156_
            lhs109_[lhs110_] = rhs157_
            lhs111_[lhs112_] = rhs158_
            lhs113_.f1 = rhs159_
        d_672_v9_: _dafny.MultiSet
        d_672_v9_ = _dafny.MultiSet([p0])
        d_673_v10_: _dafny.Set
        d_673_v10_ = _dafny.Set({p0, (self).f6})
        d_674_v11_: _dafny.Map
        d_674_v11_ = _dafny.Map({d_673_v10_: p0})
        d_675_v12_: _dafny.Map
        d_675_v12_ = _dafny.Map({(self).f6: p1})
        d_676_v13_: _dafny.Map
        d_676_v13_ = _dafny.Map({(self).f5: len(d_675_v12_)})
        rhs160_ = (self).fm10(d_672_v9_, (d_674_v11_) | (d_674_v11_), (len(d_676_v13_) if True else (d_658_v0_)[default__.safeIndex(165, (d_658_v0_).length(0))]), globalState)
        rhs161_ = d_658_v0_
        rhs162_ = (self).f6
        r0 = rhs160_
        d_658_v0_ = rhs161_
        r0 = rhs162_
        d_677_v14_: _dafny.Seq
        d_677_v14_ = _dafny.SeqWithoutIsStrInference([(d_658_v0_)[default__.safeIndex(165, (d_658_v0_).length(0))]])
        d_678_v15_: _dafny.Map
        d_678_v15_ = _dafny.Map({len(d_677_v14_): p0})
        d_679_v16_: _dafny.Seq
        d_679_v16_ = _dafny.SeqWithoutIsStrInference([len(d_678_v15_), p1, len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_680_i1_ in range(default__.abs(-294))])), p1])
        r0 = (d_679_v16_) == (d_677_v14_)
        hi3_ = p1
        for d_681_i2_ in range(p1, hi3_):
            r0 = ((self).f6) in (d_675_v12_)
            d_682_v17_: _dafny.Array
            def lambda29_(d_683_i3_):
                return not((self).f6)

            init15_ = lambda29_
            nw114_ = _dafny.Array(None, 2)
            for i0_15_ in range(nw114_.length(0)):
                nw114_[i0_15_] = init15_(i0_15_)
            d_682_v17_ = nw114_
            d_684_v18_: D0
            d_684_v18_ = D0_DC0((self).f6)
            index136_ = default__.safeIndex(387, (d_682_v17_).length(0))
            (d_682_v17_)[index136_] = (self).fm9(d_681_i2_, d_684_v18_, (self).f6, 236, globalState)
            index137_ = default__.safeIndex(387, (d_682_v17_).length(0))
            (d_682_v17_)[index137_] = (d_659_v1_)[default__.safeIndex((d_658_v0_)[default__.safeIndex(165, (d_658_v0_).length(0))], len(d_659_v1_))]
            d_685_v19_: _dafny.Map
            d_685_v19_ = _dafny.Map({(self).f5: d_658_v0_})
            d_658_v0_ = ((d_685_v19_)[(self).f5] if ((self).f5) in (d_685_v19_) else (d_658_v0_ if not(p0) else d_658_v0_))
            d_686_v20_: _dafny.Array
            d_686_v20_ = d_658_v0_
            index138_ = default__.safeIndex(387, (d_682_v17_).length(0))
            rhs163_ = (d_672_v9_).isdisjoint(d_672_v9_)
            rhs164_ = d_686_v20_
            lhs114_ = d_682_v17_
            lhs115_ = default__.safeIndex(387, (d_682_v17_).length(0))
            lhs114_[lhs115_] = rhs163_
            d_686_v20_ = rhs164_
        r0 = (_dafny.Set({p0, p0, p0})).issubset((d_673_v10_).intersection(d_673_v10_))
        r1 = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_687_i4_ in range(default__.abs(425))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jr")))
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        d_688_v0_: bool
        d_688_v0_ = False
        d_689_v1_: _dafny.Array
        nw115_ = _dafny.Array(D2.default()(), 10)
        d_689_v1_ = nw115_
        d_690_v2_: _dafny.Set
        d_690_v2_ = _dafny.Set({d_689_v1_})
        d_688_v0_ = (d_690_v2_).ispropersubset(d_690_v2_)
        d_691_v3_: int
        d_691_v3_ = -976
        d_692_v4_: _dafny.Map
        d_692_v4_ = _dafny.Map({(self).fm8(d_688_v0_, d_688_v0_, (self).f6, globalState): d_691_v3_})
        d_693_v5_: _dafny.Map
        d_693_v5_ = d_692_v4_
        d_694_v6_: _dafny.Seq
        d_694_v6_ = _dafny.SeqWithoutIsStrInference([(self).f5])
        d_693_v5_ = default__.fm27(d_694_v6_, d_688_v0_, d_688_v0_, globalState)
        r0 = d_691_v3_
        d_695_v8_: D2
        d_695_v8_ = D2_DC7(d_691_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_696_i0_ in range(default__.abs(488))])))
        d_697_v9_: _dafny.Set
        d_697_v9_ = _dafny.Set({d_695_v8_})
        d_698_v10_: _dafny.Map
        def iife40_():
            coll28_ = _dafny.Map()
            compr_28_: D2
            for compr_28_ in (d_697_v9_).Elements:
                d_699_v7_: D2 = compr_28_
                if (d_699_v7_) in (d_697_v9_):
                    coll28_[d_699_v7_] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfqu"))])).cardinality
            return _dafny.Map(coll28_)
        d_698_v10_ = _dafny.Map({(self).f5: iife40_()
        })
        d_700_v11_: _dafny.Map
        d_700_v11_ = _dafny.Map({d_695_v8_: d_691_v3_})
        d_698_v10_ = (d_698_v10_).set((self).f5, d_700_v11_)
        d_701_v12_: C1
        nw116_ = C1()
        nw116_.ctor__()
        d_701_v12_ = nw116_
        d_702_v13_: _dafny.Set
        d_702_v13_ = _dafny.Set({d_701_v12_})
        d_703_i1_: int
        d_703_i1_ = 0
        with _dafny.label("8"):
            while (d_702_v13_).issubset(d_702_v13_):
                with _dafny.c_label("8"):
                    if (d_703_i1_) >= (100):
                        raise _dafny.Break("8")
                    d_703_i1_ = (d_703_i1_) + (1)
                    d_704_v14_: _dafny.Map
                    d_704_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f5 for d_705_i2_ in range(default__.abs(308))]): (self).f6})
                    d_704_v14_ = d_704_v14_
                    r0 = d_691_v3_
                    d_688_v0_ = (True) or (d_688_v0_)
                    d_706_v15_: str
                    d_706_v15_ = _dafny.CodePoint('w')
                    d_706_v15_ = d_706_v15_
                    pass
            pass
        d_707_i3_: int
        d_707_i3_ = 0
        with _dafny.label("9"):
            while d_688_v0_:
                with _dafny.c_label("9"):
                    if (d_707_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_707_i3_ = (d_707_i3_) + (1)
                    if (self).f6:
                        d_708_v16_: _dafny.Array
                        def lambda30_(d_709_v0_):
                            def lambda31_(d_710_i4_):
                                return d_709_v0_

                            return lambda31_

                        init16_ = lambda30_(d_688_v0_)
                        nw117_ = _dafny.Array(None, 15)
                        for i0_16_ in range(nw117_.length(0)):
                            nw117_[i0_16_] = init16_(i0_16_)
                        d_708_v16_ = nw117_
                        d_711_v17_: _dafny.Map
                        d_711_v17_ = _dafny.Map({(self).f5: d_708_v16_})
                        (globalState).f1 = len(d_711_v17_)
                        d_712_v18_: _dafny.Array
                        nw118_ = _dafny.Array(int(0), 13)
                        d_712_v18_ = nw118_
                        d_712_v18_ = d_712_v18_
                        d_713_v19_: _dafny.MultiSet
                        d_713_v19_ = _dafny.MultiSet([d_708_v16_, d_708_v16_])
                        d_688_v0_ = ((d_713_v19_).cardinality) >= (len(d_694_v6_))
                        d_714_v20_: _dafny.Seq
                        d_714_v20_ = _dafny.SeqWithoutIsStrInference([d_688_v0_, (self).f6])
                        r0 = (default__.safeDivisionInt((0) - (d_691_v3_), d_691_v3_)) - (default__.fm3((d_714_v20_)[default__.safeIndex(len(d_714_v20_), len(d_714_v20_))], default__.fm0(False, globalState), globalState))
                        d_715_v21_: _dafny.MultiSet
                        d_715_v21_ = _dafny.MultiSet([d_701_v12_])
                        d_716_v22_: _dafny.Seq
                        d_716_v22_ = _dafny.SeqWithoutIsStrInference([d_691_v3_])
                        d_717_v23_: _dafny.MultiSet
                        d_717_v23_ = _dafny.MultiSet([d_691_v3_])
                        index139_ = default__.safeIndex(972, (d_712_v18_).length(0))
                        (d_712_v18_)[index139_] = (0) - (((d_715_v21_)[d_701_v12_] if (d_701_v12_) in (d_715_v21_) else (d_701_v12_).fm4((d_714_v20_).set(default__.safeIndex(len(d_714_v20_), len(d_714_v20_)), (d_701_v12_).fm6(d_716_v22_, globalState)), d_691_v3_, d_691_v3_, d_717_v23_, globalState)))
                        d_718_v24_: _dafny.Seq
                        d_718_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f5 for d_719_i5_ in range(default__.abs(761))]), d_694_v6_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ea")), d_694_v6_, d_694_v6_])
                        d_720_v25_: _dafny.Seq
                        d_720_v25_ = _dafny.SeqWithoutIsStrInference([(d_718_v24_) + (_dafny.SeqWithoutIsStrInference([d_694_v6_ for d_721_i6_ in range(default__.abs(93))])), d_718_v24_])
                        index140_ = default__.safeIndex(972, (d_712_v18_).length(0))
                        (d_712_v18_)[index140_] = len((d_720_v25_)[default__.safeIndex(d_691_v3_, len(d_720_v25_))])
                    elif True:
                        d_722_v26_: _dafny.Array
                        def lambda32_(d_723_v6_):
                            def lambda33_(d_724_i7_):
                                return (d_724_i7_) - (len(_dafny.Set({d_723_v6_, d_723_v6_})))

                            return lambda33_

                        init17_ = lambda32_(d_694_v6_)
                        nw119_ = _dafny.Array(None, 9)
                        for i0_17_ in range(nw119_.length(0)):
                            nw119_[i0_17_] = init17_(i0_17_)
                        d_722_v26_ = nw119_
                        index141_ = default__.safeIndex(211, (d_722_v26_).length(0))
                        (d_722_v26_)[index141_] = len(d_694_v6_)
                        index142_ = default__.safeIndex(211, (d_722_v26_).length(0))
                        (d_722_v26_)[index142_] = (d_691_v3_ if (d_691_v3_) >= ((0) - (d_691_v3_)) else default__.safeDivisionInt(len(d_694_v6_), len(d_692_v4_)))
                        d_722_v26_ = d_722_v26_
                        d_725_v27_: _dafny.Map
                        d_725_v27_ = _dafny.Map({d_688_v0_: (self).f6})
                        d_726_v28_: _dafny.Seq
                        d_726_v28_ = _dafny.SeqWithoutIsStrInference([(self).fm7(False, globalState)])
                        d_725_v27_ = (d_725_v27_).set((self).f6, ((self).fm6(d_726_v28_, globalState)) or (d_688_v0_))
                        d_727_v29_: _dafny.Seq
                        d_727_v29_ = _dafny.SeqWithoutIsStrInference([d_688_v0_])
                        index143_ = default__.safeIndex(211, (d_722_v26_).length(0))
                        (d_722_v26_)[index143_] = len(((d_727_v29_) + (d_727_v29_)) + ((d_727_v29_).set(default__.safeIndex(d_691_v3_, len(d_727_v29_)), (self).f6)))
                        d_728_v30_: _dafny.Array
                        nw120_ = _dafny.Array(False, 20)
                        d_728_v30_ = nw120_
                        d_729_v31_: str
                        d_729_v31_ = _dafny.CodePoint('c')
                        index144_ = default__.safeIndex(211, (d_722_v26_).length(0))
                        rhs165_ = (d_726_v28_)[default__.safeIndex(d_691_v3_, len(d_726_v28_))]
                        rhs166_ = d_694_v6_
                        rhs167_ = d_728_v30_
                        rhs168_ = (self).f5
                        rhs169_ = (self).f5
                        lhs116_ = d_722_v26_
                        lhs117_ = default__.safeIndex(211, (d_722_v26_).length(0))
                        lhs116_[lhs117_] = rhs165_
                        d_694_v6_ = rhs166_
                        d_728_v30_ = rhs167_
                        d_729_v31_ = rhs168_
                        d_729_v31_ = rhs169_
                    d_730_v32_: _dafny.Array
                    nw121_ = _dafny.Array(False, 8)
                    d_730_v32_ = nw121_
                    index145_ = default__.safeIndex(547, (d_730_v32_).length(0))
                    (d_730_v32_)[index145_] = d_688_v0_
                    index146_ = default__.safeIndex(547, (d_730_v32_).length(0))
                    (d_730_v32_)[index146_] = d_688_v0_
                    (globalState).f1 = (0) - (d_691_v3_)
                    d_731_v33_: _dafny.Set
                    d_731_v33_ = _dafny.Set({d_688_v0_})
                    d_732_v34_: _dafny.Array
                    nw122_ = _dafny.Array(int(0), 12)
                    d_732_v34_ = nw122_
                    index147_ = default__.safeIndex(889, (d_732_v34_).length(0))
                    (d_732_v34_)[index147_] = len((d_694_v6_) + (_dafny.SeqWithoutIsStrInference([(self).f5 for d_733_i8_ in range(default__.abs(422))])))
                    d_734_v35_: _dafny.Seq
                    d_734_v35_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                    index148_ = default__.safeIndex(889, (d_732_v34_).length(0))
                    rhs170_ = (d_731_v33_) - (_dafny.Set({(d_730_v32_)[default__.safeIndex(547, (d_730_v32_).length(0))], d_688_v0_}))
                    rhs171_ = default__.safeModuloInt((d_691_v3_) + (len(d_734_v35_)), d_691_v3_)
                    lhs118_ = d_732_v34_
                    lhs119_ = default__.safeIndex(889, (d_732_v34_).length(0))
                    d_731_v33_ = rhs170_
                    lhs118_[lhs119_] = rhs171_
                    pass
            pass
        r0 = (0) - (d_691_v3_)
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_735_v0_: _dafny.Seq
        d_735_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
        d_735_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
        d_736_v1_: bool
        d_736_v1_ = False
        d_737_v2_: _dafny.Set
        d_737_v2_ = _dafny.Set({p2, d_736_v1_})
        d_736_v1_ = (d_737_v2_).ispropersubset(d_737_v2_)
        d_738_v3_: int
        d_738_v3_ = 153
        d_739_v4_: _dafny.Map
        d_739_v4_ = _dafny.Map({_dafny.CodePoint('m'): d_738_v3_})
        d_740_v5_: _dafny.Array
        nw123_ = _dafny.Array(None, 6)
        nw123_[int(0)] = d_738_v3_
        nw123_[int(1)] = (0) - (d_738_v3_)
        nw123_[int(2)] = d_738_v3_
        nw123_[int(3)] = default__.safeModuloInt(((d_739_v4_)[(self).f5] if ((self).f5) in (d_739_v4_) else d_738_v3_), d_738_v3_)
        nw123_[int(4)] = (0) - (d_738_v3_)
        nw123_[int(5)] = 632
        d_740_v5_ = nw123_
        index149_ = default__.safeIndex(784, (d_740_v5_).length(0))
        (d_740_v5_)[index149_] = (0) - (d_738_v3_)
        d_741_v6_: _dafny.Set
        d_741_v6_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f5 for d_742_i0_ in range(default__.abs(-352))]), d_735_v0_, d_735_v0_, d_735_v0_, (d_735_v0_) + (d_735_v0_)})
        index150_ = default__.safeIndex(784, (d_740_v5_).length(0))
        rhs172_ = (d_736_v1_) and (p0)
        rhs173_ = len(d_741_v6_)
        lhs120_ = d_740_v5_
        lhs121_ = default__.safeIndex(784, (d_740_v5_).length(0))
        d_736_v1_ = rhs172_
        lhs120_[lhs121_] = rhs173_
        d_743_v7_: _dafny.Array
        nw124_ = _dafny.Array(_dafny.Array(None, 0), 13)
        d_743_v7_ = nw124_
        d_744_v8_: _dafny.Array
        d_744_v8_ = d_740_v5_
        index151_ = default__.safeIndex(274, (d_743_v7_).length(0))
        (d_743_v7_)[index151_] = d_744_v8_
        index152_ = default__.safeIndex(274, (d_743_v7_).length(0))
        (d_743_v7_)[index152_] = d_740_v5_
        d_738_v3_ = d_738_v3_
        d_745_v9_: _dafny.Seq
        d_745_v9_ = _dafny.SeqWithoutIsStrInference([d_736_v1_])
        d_746_i1_: int
        d_746_i1_ = 0
        with _dafny.label("10"):
            while ((d_745_v9_)[default__.safeIndex((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))], len(d_745_v9_))]) == (p0):
                with _dafny.c_label("10"):
                    if (d_746_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_746_i1_ = (d_746_i1_) + (1)
                    d_747_v10_: _dafny.Map
                    d_747_v10_ = _dafny.Map({True: not(p2)})
                    d_748_v11_: _dafny.Array
                    def lambda34_(d_749_v3_):
                        def lambda35_(d_750_i2_):
                            return (_dafny.MultiSet([d_749_v3_, 4])) | (_dafny.MultiSet([d_749_v3_]))

                        return lambda35_

                    init18_ = lambda34_(d_738_v3_)
                    nw125_ = _dafny.Array(None, 5)
                    for i0_18_ in range(nw125_.length(0)):
                        nw125_[i0_18_] = init18_(i0_18_)
                    d_748_v11_ = nw125_
                    d_751_v12_: _dafny.Seq
                    d_751_v12_ = _dafny.SeqWithoutIsStrInference([d_738_v3_])
                    d_752_v13_: _dafny.Map
                    d_752_v13_ = _dafny.Map({d_738_v3_: len(d_735_v0_)})
                    d_753_v14_: _dafny.Set
                    d_753_v14_ = _dafny.Set({len(d_752_v13_)})
                    d_754_v15_: _dafny.MultiSet
                    d_754_v15_ = _dafny.MultiSet([(d_751_v12_)[default__.safeIndex(len(d_753_v14_), len(d_751_v12_))], (0) - (len(d_751_v12_))])
                    index153_ = default__.safeIndex(194, (d_748_v11_).length(0))
                    (d_748_v11_)[index153_] = d_754_v15_
                    index154_ = default__.safeIndex(194, (d_748_v11_).length(0))
                    rhs174_ = d_747_v10_
                    rhs175_ = d_735_v0_
                    rhs176_ = (0) - (d_738_v3_)
                    rhs177_ = default__.safeModuloInt((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))], d_738_v3_)
                    rhs178_ = (d_754_v15_).set(d_738_v3_, default__.abs(d_738_v3_))
                    lhs122_ = d_748_v11_
                    lhs123_ = default__.safeIndex(194, (d_748_v11_).length(0))
                    d_747_v10_ = rhs174_
                    d_735_v0_ = rhs175_
                    d_738_v3_ = rhs176_
                    d_738_v3_ = rhs177_
                    lhs122_[lhs123_] = rhs178_
                    (globalState).f1 = d_738_v3_
                    d_755_v16_: _dafny.Map
                    d_755_v16_ = _dafny.Map({d_737_v2_: (self).f6})
                    d_756_v17_: _dafny.Map
                    d_756_v17_ = _dafny.Map({d_738_v3_: p0})
                    d_757_v18_: _dafny.Map
                    d_757_v18_ = _dafny.Map({d_751_v12_: 316})
                    d_758_v19_: _dafny.Seq
                    d_758_v19_ = _dafny.SeqWithoutIsStrInference([d_757_v18_, (d_757_v18_).set(d_751_v12_, len(_dafny.Map({d_736_v1_: (d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]})))])
                    d_759_v20_: _dafny.Seq
                    d_759_v20_ = _dafny.SeqWithoutIsStrInference([(d_758_v19_)[default__.safeIndex(d_738_v3_, len(d_758_v19_))]])
                    d_760_v21_: D0
                    d_760_v21_ = D0_DC0((self).fm10(_dafny.MultiSet(d_745_v9_), d_755_v16_, (d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))], globalState))
                    d_761_v22_: _dafny.Array
                    nw126_ = _dafny.Array(None, 16)
                    nw126_[int(0)] = (self).fm6(_dafny.SeqWithoutIsStrInference([(0) - ((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]) for d_762_i3_ in range(default__.abs(388))]), globalState)
                    nw126_[int(1)] = p0
                    nw126_[int(2)] = (self).f6
                    nw126_[int(3)] = True
                    nw126_[int(4)] = not((d_745_v9_)[default__.safeIndex(len(d_735_v0_), len(d_745_v9_))])
                    nw126_[int(5)] = ((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]) == (468)
                    nw126_[int(6)] = (self).fm10(_dafny.MultiSet([p2, d_736_v1_]), d_755_v16_, d_738_v3_, globalState)
                    nw126_[int(7)] = (self).f6
                    nw126_[int(8)] = p0
                    nw126_[int(9)] = (self).f6
                    nw126_[int(10)] = p2
                    nw126_[int(11)] = not (((d_756_v17_)[(d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]] if ((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]) in (d_756_v17_) else d_736_v1_)) or (False)
                    nw126_[int(12)] = not(not(not(((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))]) >= (len(default__.fm28(d_736_v1_, p0, (d_758_v19_)[default__.safeIndex((d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))], len(d_758_v19_))], globalState))))))
                    nw126_[int(13)] = (d_760_v21_).cf0
                    nw126_[int(14)] = ((self).f6) or (not((self).f6))
                    nw126_[int(15)] = (686) != (53)
                    d_761_v22_ = nw126_
                    d_761_v22_ = d_761_v22_
                    d_763_v23_: D0
                    d_763_v23_ = D0_DC0(p2)
                    d_764_v24_: D0
                    d_764_v24_ = D0_DC3(d_763_v23_)
                    d_765_v25_: D2
                    d_765_v25_ = D2_DC7(d_738_v3_, (d_740_v5_)[default__.safeIndex(784, (d_740_v5_).length(0))])
                    d_766_v26_: _dafny.Map
                    d_766_v26_ = _dafny.Map({d_764_v24_: (d_765_v25_).cf8})
                    pat_let_tv10_ = d_763_v23_
                    def iife41_(_pat_let6_0):
                        def iife42_(d_767_dt__update__tmp_h1_):
                            def iife43_(_pat_let7_0):
                                def iife44_(d_768_dt__update_hcf3_h0_):
                                    return D0_DC3(d_768_dt__update_hcf3_h0_)
                                return iife44_(_pat_let7_0)
                            return iife43_(pat_let_tv10_)
                        return iife42_(_pat_let6_0)
                    d_766_v26_ = (d_766_v26_).set(iife41_(d_764_v24_), (d_751_v12_)[default__.safeIndex((_dafny.MultiSet([d_738_v3_])).cardinality, len(d_751_v12_))])
                    pass
            pass


class C3(T0, T1):
    def  __init__(self):
        self._f13: bool = False
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f13, f14):
        (self)._f13 = f13
        (self)._f14 = f14

    def fm4(self, p0, p1, p2, p3, globalState):
        def iife45_():
            coll29_ = _dafny.Set()
            compr_29_: int
            for compr_29_ in _dafny.IntegerRange(403, 547):
                d_769_v0_: int = compr_29_
                if ((403) <= (d_769_v0_)) and ((d_769_v0_) < (547)):
                    coll29_ = coll29_.union(_dafny.Set([(d_769_v0_) + (-565)]))
            return _dafny.Set(coll29_)
        def iife46_():
            coll30_ = _dafny.Set()
            compr_30_: int
            for compr_30_ in (_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))): (self).f13})): (self).f14})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ueuscelno")))])).Elements:
                d_770_v1_: int = compr_30_
                if (d_770_v1_) in (_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))): (self).f13})): (self).f14})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ueuscelno")))])):
                    coll30_ = coll30_.union(_dafny.Set([(d_770_v1_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_771_i0_ in range(default__.abs(603))])))]))
            return _dafny.Set(coll30_)
        return (0) - (default__.safeDivisionInt(42, len((_dafny.SeqWithoutIsStrInference([iife45_()
        , _dafny.Set({46, 880}), iife46_()
        ]) if (self).f14 else _dafny.SeqWithoutIsStrInference([_dafny.Set({(_dafny.MultiSet([(self).f13])).cardinality})])))))

    def fm5(self, p0, globalState):
        return D1_DC4((D2_DC7(135, len(_dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])))).cf7)

    def m1(self, p0, p1, p2, globalState):
        d_772_v0_: _dafny.Array
        def lambda36_(d_773_p1_):
            def lambda37_(d_774_i0_):
                return (d_774_i0_) + (len(_dafny.SeqWithoutIsStrInference([(0) - (-324), (d_773_p1_).cardinality])))

            return lambda37_

        init19_ = lambda36_(p1)
        nw127_ = _dafny.Array(None, 17)
        for i0_19_ in range(nw127_.length(0)):
            nw127_[i0_19_] = init19_(i0_19_)
        d_772_v0_ = nw127_
        d_775_v1_: int
        d_775_v1_ = -263
        d_776_v2_: _dafny.Seq
        d_776_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_777_v3_: _dafny.Set
        d_777_v3_ = _dafny.Set({114, d_775_v1_, len(d_776_v2_)})
        index155_ = default__.safeIndex(248, (d_772_v0_).length(0))
        (d_772_v0_)[index155_] = len(d_777_v3_)
        index156_ = default__.safeIndex(248, (d_772_v0_).length(0))
        (d_772_v0_)[index156_] = d_775_v1_
        d_778_v4_: bool
        d_778_v4_ = False
        d_778_v4_ = (d_776_v2_)[default__.safeIndex((d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))], len(d_776_v2_))]
        d_779_v5_: _dafny.Seq
        d_779_v5_ = _dafny.SeqWithoutIsStrInference([(d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))], d_775_v1_])
        d_780_v6_: _dafny.Seq
        d_780_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
        d_781_v7_: _dafny.Seq
        d_781_v7_ = _dafny.SeqWithoutIsStrInference([d_780_v6_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_782_i1_ in range(default__.abs(707))])])
        d_783_v8_: D2
        d_783_v8_ = D2_DC6(default__.fm1(897, True, (d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))], globalState))
        d_784_v9_: D2
        d_784_v9_ = D2_DC9(d_783_v8_)
        d_785_v10_: D1
        d_785_v10_ = D1_DC4(len(d_781_v7_))
        d_786_v11_: _dafny.Array
        nw128_ = _dafny.Array(False, 28)
        d_786_v11_ = nw128_
        d_787_v12_: _dafny.Array
        d_787_v12_ = d_786_v11_
        d_788_v13_: _dafny.Map
        d_788_v13_ = _dafny.Map({default__.fm0(d_778_v4_, globalState): d_787_v12_})
        d_789_v14_: _dafny.Seq
        d_789_v14_ = _dafny.SeqWithoutIsStrInference([d_788_v13_])
        d_790_v15_: _dafny.Map
        d_790_v15_ = _dafny.Map({d_775_v1_: (d_789_v14_)[default__.safeIndex(len(d_780_v6_), len(d_789_v14_))]})
        d_791_v16_: _dafny.Map
        d_791_v16_ = _dafny.Map({d_790_v15_: not(p2)})
        rhs179_ = default__.fm37(not((self).f13), default__.safeDivisionInt(len(d_779_v5_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))), d_785_v10_, d_775_v1_, globalState)
        rhs180_ = d_781_v7_
        rhs181_ = d_784_v9_
        rhs182_ = ((d_791_v16_)[d_790_v15_] if (d_790_v15_) in (d_791_v16_) else (self).f14)
        rhs183_ = d_775_v1_
        lhs124_ = globalState
        d_779_v5_ = rhs179_
        d_781_v7_ = rhs180_
        d_784_v9_ = rhs181_
        d_778_v4_ = rhs182_
        lhs124_.f1 = rhs183_
        d_792_v17_: _dafny.Map
        d_792_v17_ = _dafny.Map({d_775_v1_: (d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))]})
        d_793_v18_: _dafny.Set
        d_793_v18_ = _dafny.Set({default__.fm0((self).f14, globalState), (self).f13, (self).f14})
        (globalState).f1 = default__.safeModuloInt(len(d_792_v17_), len(d_793_v18_))
        if (self).f13:
            d_794_v19_: D0
            d_794_v19_ = D0_DC0(d_778_v4_)
            d_778_v4_ = (d_794_v19_).cf0
            d_795_v20_: str
            d_795_v20_ = _dafny.CodePoint('a')
            d_796_v21_: _dafny.Map
            d_796_v21_ = _dafny.Map({d_786_v11_: d_795_v20_})
            d_797_v22_: D0
            d_797_v22_ = D0_DC2(d_796_v21_, d_795_v20_)
            pat_let_tv11_ = d_786_v11_
            pat_let_tv12_ = d_795_v20_
            def iife47_(_pat_let8_0):
                def iife48_(d_798_dt__update__tmp_h0_):
                    def iife49_(_pat_let9_0):
                        def iife50_(d_799_dt__update_hcf1_h0_):
                            return D0_DC2(d_799_dt__update_hcf1_h0_, (d_798_dt__update__tmp_h0_).cf2)
                        return iife50_(_pat_let9_0)
                    return iife49_(_dafny.Map({pat_let_tv11_: pat_let_tv12_}))
                return iife48_(_pat_let8_0)
            d_795_v20_ = (iife47_(d_797_v22_)).cf2
            d_778_v4_ = (58) < (d_775_v1_)
            d_778_v4_ = (p2) == ((243) == (d_775_v1_))
            d_800_v23_: _dafny.Array
            nw129_ = _dafny.Array(None, 23)
            nw129_[int(0)] = d_795_v20_
            nw129_[int(1)] = d_795_v20_
            nw129_[int(2)] = d_795_v20_
            nw129_[int(3)] = d_795_v20_
            nw129_[int(4)] = d_795_v20_
            nw129_[int(5)] = d_795_v20_
            nw129_[int(6)] = _dafny.CodePoint('x')
            nw129_[int(7)] = d_795_v20_
            nw129_[int(8)] = (d_795_v20_ if p0 else d_795_v20_)
            nw129_[int(9)] = d_795_v20_
            nw129_[int(10)] = d_795_v20_
            nw129_[int(11)] = d_795_v20_
            nw129_[int(12)] = d_795_v20_
            nw129_[int(13)] = d_795_v20_
            nw129_[int(14)] = d_795_v20_
            nw129_[int(15)] = d_795_v20_
            nw129_[int(16)] = d_795_v20_
            nw129_[int(17)] = d_795_v20_
            nw129_[int(18)] = d_795_v20_
            nw129_[int(19)] = d_795_v20_
            nw129_[int(20)] = d_795_v20_
            nw129_[int(21)] = d_795_v20_
            nw129_[int(22)] = (d_780_v6_)[default__.safeIndex(len(d_780_v6_), len(d_780_v6_))]
            d_800_v23_ = nw129_
            index157_ = default__.safeIndex(535, (d_800_v23_).length(0))
            (d_800_v23_)[index157_] = _dafny.CodePoint('d')
            index158_ = default__.safeIndex(535, (d_800_v23_).length(0))
            (d_800_v23_)[index158_] = d_795_v20_
        elif True:
            index159_ = default__.safeIndex(248, (d_772_v0_).length(0))
            (d_772_v0_)[index159_] = (d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))]
            index160_ = default__.safeIndex(248, (d_772_v0_).length(0))
            rhs184_ = not((d_776_v2_)[default__.safeIndex(d_775_v1_, len(d_776_v2_))])
            rhs185_ = (d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))]
            rhs186_ = d_776_v2_
            lhs125_ = d_772_v0_
            lhs126_ = default__.safeIndex(248, (d_772_v0_).length(0))
            d_778_v4_ = rhs184_
            lhs125_[lhs126_] = rhs185_
            d_776_v2_ = rhs186_
            d_801_v24_: _dafny.MultiSet
            d_801_v24_ = _dafny.MultiSet([(d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))]])
            (globalState).f1 = default__.safeModuloInt(default__.safeModuloInt(872, (self).fm4(d_776_v2_, d_775_v1_, 176, d_801_v24_, globalState)), (0) - ((781) - (d_775_v1_)))
            d_802_v25_: str
            d_802_v25_ = _dafny.CodePoint('p')
            d_803_v26_: D0
            d_804_v27_: bool
            out28_: D0
            out29_: bool
            out28_, out29_ = default__.m0(d_786_v11_, d_802_v25_, p0, d_802_v25_, globalState)
            d_803_v26_ = out28_
            d_804_v27_ = out29_
            d_802_v25_ = _dafny.CodePoint('x')
        (globalState).f1 = (0) - ((d_772_v0_)[default__.safeIndex(248, (d_772_v0_).length(0))])

    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C4(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm4(self, p0, p1, p2, p3, globalState):
        if (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).issubset(_dafny.MultiSet([True])):
            return (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lsyykqbc")))) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ol")))))
        elif True:
            return len(_dafny.Map({_dafny.MultiSet([338]): False}))

    def fm5(self, p0, globalState):
        if True:
            def iife51_():
                coll31_ = _dafny.Map()
                compr_31_: int
                for compr_31_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfa"))), (0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))]))).Elements:
                    d_805_v0_: int = compr_31_
                    if (d_805_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qfa"))), (0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))]))):
                        coll31_[default__.safeDivisionInt(d_805_v0_, len(_dafny.Map({_dafny.CodePoint('x'): _dafny.SeqWithoutIsStrInference([-454])})))] = -177
                return _dafny.Map(coll31_)
            return D1_DC4(len(iife51_()
))
        elif True:
            return D1_DC4(368)

    def fm30(self, p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_806_i0_ in range(default__.abs(-327))]) if False else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))))

    def m1(self, p0, p1, p2, globalState):
        d_807_v0_: bool
        d_807_v0_ = True
        d_807_v0_ = False
        d_808_v1_: C0
        nw130_ = C0()
        nw130_.ctor__(p0, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([822, (_dafny.MultiSet([645, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-537]))).cardinality])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_809_i1_ in range(default__.abs(652))])), 830, 979]) for d_810_i0_ in range(default__.abs(-580))]))
        d_808_v1_ = nw130_
        d_807_v0_ = d_807_v0_
        d_811_v2_: _dafny.Seq
        d_811_v2_ = _dafny.SeqWithoutIsStrInference([d_807_v0_])
        d_812_i2_: int
        d_812_i2_ = 0
        with _dafny.label("11"):
            while not((d_811_v2_)[default__.safeIndex(-736, len(d_811_v2_))]):
                with _dafny.c_label("11"):
                    if (d_812_i2_) >= (100):
                        raise _dafny.Break("11")
                    d_812_i2_ = (d_812_i2_) + (1)
                    d_813_v3_: _dafny.Seq
                    d_813_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iygwbkys"))
                    d_814_v4_: _dafny.Seq
                    d_814_v4_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkbpmdi"))])
                    d_813_v3_ = ((d_814_v4_)[default__.safeIndex(-365, len(d_814_v4_))]) + (d_813_v3_)
                    d_815_v5_: int
                    d_815_v5_ = 723
                    d_816_v6_: D2
                    d_816_v6_ = D2_DC6(d_813_v3_)
                    d_817_v7_: _dafny.Set
                    d_817_v7_ = _dafny.Set({(d_808_v1_).f11})
                    d_818_v8_: _dafny.Map
                    d_818_v8_ = _dafny.Map({p0: d_817_v7_})
                    d_807_v0_ = (_dafny.MultiSet(((default__.fm31(p0, d_815_v5_, globalState)).set(default__.safeIndex(len((d_816_v6_).cf6), len(default__.fm31(p0, d_815_v5_, globalState))), 888)).set(default__.safeIndex(len(d_818_v8_), len((default__.fm31(p0, d_815_v5_, globalState)).set(default__.safeIndex(len((d_816_v6_).cf6), len(default__.fm31(p0, d_815_v5_, globalState))), 888))), len(d_813_v3_)))).isdisjoint(default__.fm32(globalState))
                    d_815_v5_ = d_815_v5_
                    (globalState).f1 = d_815_v5_
                    pass
            pass
        d_819_v9_: int
        d_819_v9_ = 140
        d_820_v10_: _dafny.Map
        d_820_v10_ = _dafny.Map({d_819_v9_: p2})
        d_821_v11_: _dafny.Map
        d_821_v11_ = _dafny.Map({(d_819_v9_) * (d_819_v9_): (default__.fm33(globalState)) | (d_820_v10_)})
        d_820_v10_ = ((d_821_v11_)[d_819_v9_] if (d_819_v9_) in (d_821_v11_) else d_820_v10_)
        d_822_i3_: int
        d_822_i3_ = 0
        with _dafny.label("12"):
            while (default__.fm0(d_807_v0_, globalState) if not (d_807_v0_) or (False) else d_807_v0_):
                with _dafny.c_label("12"):
                    if (d_822_i3_) >= (100):
                        raise _dafny.Break("12")
                    d_822_i3_ = (d_822_i3_) + (1)
                    d_823_v12_: C1
                    nw131_ = C1()
                    nw131_.ctor__()
                    d_823_v12_ = nw131_
                    d_824_v13_: _dafny.Array
                    nw132_ = _dafny.Array(False, 18)
                    d_824_v13_ = nw132_
                    d_825_v14_: _dafny.Array
                    d_825_v14_ = d_824_v13_
                    d_826_v15_: _dafny.Map
                    d_826_v15_ = _dafny.Map({p0: d_824_v13_})
                    d_827_v16_: _dafny.Array
                    nw133_ = _dafny.Array(None, 13)
                    nw133_[int(0)] = d_824_v13_
                    nw133_[int(1)] = d_824_v13_
                    nw133_[int(2)] = d_824_v13_
                    nw133_[int(3)] = d_824_v13_
                    nw133_[int(4)] = d_824_v13_
                    nw133_[int(5)] = d_824_v13_
                    nw133_[int(6)] = (d_825_v14_)
                    nw133_[int(7)] = d_824_v13_
                    nw133_[int(8)] = d_824_v13_
                    nw133_[int(9)] = ((d_826_v15_)[(d_808_v1_).f11] if ((d_808_v1_).f11) in (d_826_v15_) else d_824_v13_)
                    nw133_[int(10)] = d_824_v13_
                    nw133_[int(11)] = d_824_v13_
                    nw133_[int(12)] = d_824_v13_
                    d_827_v16_ = nw133_
                    index161_ = default__.safeIndex(610, (d_827_v16_).length(0))
                    (d_827_v16_)[index161_] = d_824_v13_
                    d_828_v17_: _dafny.Seq
                    d_828_v17_ = _dafny.SeqWithoutIsStrInference([d_824_v13_])
                    index162_ = default__.safeIndex(610, (d_827_v16_).length(0))
                    (d_827_v16_)[index162_] = (d_828_v17_)[default__.safeIndex(d_819_v9_, len(d_828_v17_))]
                    d_829_v18_: D2
                    d_829_v18_ = D2_DC8(True)
                    d_807_v0_ = ((d_829_v18_).cf9 if d_807_v0_ else (d_808_v1_).f11)
                    if (d_819_v9_) == (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0, False, d_807_v0_])): d_819_v9_}))):
                        d_807_v0_ = (d_811_v2_)[default__.safeIndex((d_819_v9_) + (d_819_v9_), len(d_811_v2_))]
                        d_830_v19_: D2
                        d_830_v19_ = D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))
                        d_830_v19_ = d_830_v19_
                        d_831_v20_: _dafny.Array
                        nw134_ = _dafny.Array(_dafny.Map({}), 15)
                        d_831_v20_ = nw134_
                        d_832_v21_: _dafny.Map
                        d_832_v21_ = _dafny.Map({d_824_v13_: ((d_808_v1_).fm5(d_829_v18_, globalState)).cf4})
                        index163_ = default__.safeIndex(927, (d_831_v20_).length(0))
                        (d_831_v20_)[index163_] = d_832_v21_
                        d_833_v22_: _dafny.Seq
                        d_833_v22_ = _dafny.SeqWithoutIsStrInference([d_832_v21_, d_832_v21_, d_832_v21_, d_832_v21_, d_832_v21_])
                        d_834_v23_: _dafny.Map
                        d_834_v23_ = _dafny.Map({(d_808_v1_).f11: default__.fm34((d_808_v1_).f11, globalState)})
                        index164_ = default__.safeIndex(927, (d_831_v20_).length(0))
                        rhs187_ = (d_833_v22_)[default__.safeIndex(d_819_v9_, len(d_833_v22_))]
                        rhs188_ = ((((d_834_v23_)[p0] if (p0) in (d_834_v23_) else p1)) == (_dafny.MultiSet([p0]))) and ((d_808_v1_).f11)
                        rhs189_ = (d_808_v1_).f11
                        lhs127_ = d_831_v20_
                        lhs128_ = default__.safeIndex(927, (d_831_v20_).length(0))
                        lhs127_[lhs128_] = rhs187_
                        d_807_v0_ = rhs188_
                        d_807_v0_ = rhs189_
                        d_807_v0_ = True
                        d_835_v24_: _dafny.MultiSet
                        d_835_v24_ = _dafny.MultiSet([d_819_v9_])
                        d_836_v25_: str
                        d_836_v25_ = _dafny.CodePoint('d')
                        d_837_v26_: _dafny.Map
                        d_837_v26_ = _dafny.Map({d_836_v25_: d_819_v9_})
                        d_838_v27_: _dafny.Map
                        d_838_v27_ = _dafny.Map({_dafny.Map({p2: (d_835_v24_).cardinality}): len(d_837_v26_)})
                        d_839_v28_: _dafny.MultiSet
                        d_839_v28_ = _dafny.MultiSet([len(d_838_v27_), d_819_v9_, d_819_v9_, (0) - (d_819_v9_), 741])
                        d_840_v29_: _dafny.Map
                        d_840_v29_ = _dafny.Map({d_808_v1_: d_839_v28_})
                        d_841_v30_: _dafny.Map
                        d_841_v30_ = _dafny.Map({d_819_v9_: (self).fm4(d_811_v2_, 574, d_819_v9_, ((d_840_v29_)[d_808_v1_] if (d_808_v1_) in (d_840_v29_) else d_839_v28_), globalState)})
                        d_842_v31_: _dafny.Set
                        d_842_v31_ = _dafny.Set({p0})
                        d_841_v30_ = (d_841_v30_).set(d_819_v9_, len(default__.fm31(d_807_v0_, len(d_842_v31_), globalState)))
                    elif True:
                        d_843_v32_: C1
                        nw135_ = C1()
                        nw135_.ctor__()
                        d_843_v32_ = nw135_
                        d_823_v12_ = d_823_v12_
                        (globalState).f1 = (0) - (d_819_v9_)
                        d_844_v34_: _dafny.Map
                        def iife52_():
                            coll32_ = _dafny.Map()
                            compr_32_: int
                            for compr_32_ in _dafny.IntegerRange(599, 961):
                                d_845_v33_: int = compr_32_
                                if ((599) <= (d_845_v33_)) and ((d_845_v33_) < (961)):
                                    coll32_[(d_845_v33_) + (d_819_v9_)] = _dafny.CodePoint('q')
                            return _dafny.Map(coll32_)
                        d_844_v34_ = _dafny.Map({(d_808_v1_).f11: (len(iife52_()
                        )) + (d_819_v9_)})
                        d_844_v34_ = (d_844_v34_).set(not (d_807_v0_) or ((d_808_v1_).f11), d_819_v9_)
                        d_807_v0_ = (default__.safeModuloInt((0) - (d_819_v9_), (0) - (d_819_v9_))) == (d_819_v9_)
                    pass
            pass

    def m10(self, p0, globalState):
        r0: C0 = None
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        if False:
            d_846_v0_: bool
            d_846_v0_ = False
            d_847_v1_: int
            d_847_v1_ = 479
            d_848_v2_: _dafny.Map
            d_848_v2_ = _dafny.Map({d_846_v0_: d_847_v1_})
            d_849_v3_: _dafny.Seq
            d_849_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_847_v1_]), default__.fm31(d_846_v0_, d_847_v1_, globalState), _dafny.SeqWithoutIsStrInference([len(d_848_v2_), d_847_v1_, d_847_v1_])])
            d_850_v4_: C0
            nw136_ = C0()
            nw136_.ctor__(not (d_846_v0_) or (d_846_v0_), d_849_v3_)
            d_850_v4_ = nw136_
            d_851_v5_: str
            d_851_v5_ = _dafny.CodePoint('d')
            d_851_v5_ = d_851_v5_
            d_852_v6_: _dafny.Seq
            d_852_v6_ = _dafny.SeqWithoutIsStrInference([(d_850_v4_).f11, d_846_v0_, False])
            d_853_v7_: D1
            d_853_v7_ = D1_DC4(d_847_v1_)
            d_854_v8_: _dafny.MultiSet
            d_854_v8_ = _dafny.MultiSet([(0) - (d_847_v1_)])
            (globalState).f1 = (d_850_v4_).fm4(d_852_v6_, d_847_v1_, (d_853_v7_).cf4, d_854_v8_, globalState)
            (globalState).f1 = ((d_847_v1_ if d_846_v0_ else d_847_v1_)) + (460)
            index165_ = default__.safeIndex(92, (p0).length(0))
            (p0)[index165_] = (d_850_v4_).f11
            d_855_v9_: _dafny.MultiSet
            d_855_v9_ = _dafny.MultiSet([d_846_v0_])
            d_856_v10_: _dafny.Map
            d_856_v10_ = _dafny.Map({(d_855_v9_).cardinality: d_846_v0_})
            d_857_v11_: _dafny.Map
            d_857_v11_ = _dafny.Map({d_850_v4_: ((d_856_v10_)[(0) - (len(_dafny.Map({d_847_v1_: (d_850_v4_).f11})))] if ((0) - (len(_dafny.Map({d_847_v1_: (d_850_v4_).f11})))) in (d_856_v10_) else d_846_v0_)})
            index166_ = default__.safeIndex(850, (p0).length(0))
            (p0)[index166_] = (d_850_v4_) in (d_857_v11_)
            index167_ = default__.safeIndex(92, (p0).length(0))
            index168_ = default__.safeIndex(850, (p0).length(0))
            rhs190_ = default__.fm0((d_850_v4_).f11, globalState)
            rhs191_ = d_846_v0_
            rhs192_ = ((d_846_v0_ if d_846_v0_ else (d_850_v4_).f11) if d_846_v0_ else d_846_v0_)
            lhs129_ = p0
            lhs130_ = default__.safeIndex(92, (p0).length(0))
            lhs131_ = p0
            lhs132_ = default__.safeIndex(850, (p0).length(0))
            lhs129_[lhs130_] = rhs190_
            lhs131_[lhs132_] = rhs191_
            d_846_v0_ = rhs192_
        elif True:
            d_858_v12_: bool
            d_858_v12_ = True
            d_858_v12_ = default__.fm0(d_858_v12_, globalState)
            d_859_v13_: _dafny.Seq
            d_859_v13_ = _dafny.SeqWithoutIsStrInference([p0, (p0 if d_858_v12_ else p0)])
            d_860_v14_: int
            d_860_v14_ = 497
            d_859_v13_ = (d_859_v13_).set(default__.safeIndex(d_860_v14_, len(d_859_v13_)), p0)
            d_861_v15_: _dafny.Seq
            d_861_v15_ = _dafny.SeqWithoutIsStrInference([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([d_858_v12_]))))])
            d_861_v15_ = (d_861_v15_) + (_dafny.SeqWithoutIsStrInference([389 for d_862_i0_ in range(default__.abs(-694))]))
            d_863_v16_: _dafny.Seq
            d_863_v16_ = _dafny.SeqWithoutIsStrInference([d_861_v15_, d_861_v15_])
            d_864_v17_: _dafny.Seq
            d_864_v17_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_861_v15_]), d_863_v16_])
            d_865_v18_: C0
            nw137_ = C0()
            nw137_.ctor__((default__.fm0(d_858_v12_, globalState) if False else not(d_858_v12_)), (d_863_v16_) + ((d_864_v17_)[default__.safeIndex(d_860_v14_, len(d_864_v17_))]))
            d_865_v18_ = nw137_
            (globalState).f1 = d_860_v14_
        d_866_v19_: int
        d_866_v19_ = 470
        d_867_v20_: _dafny.Seq
        d_867_v20_ = _dafny.SeqWithoutIsStrInference([d_866_v19_])
        hi4_ = (d_867_v20_)[default__.safeIndex(d_866_v19_, len(d_867_v20_))]
        for d_868_i1_ in range(default__.safeModuloInt(d_866_v19_, d_866_v19_), hi4_):
            d_869_v21_: bool
            d_869_v21_ = False
            index169_ = default__.safeIndex(600, (p0).length(0))
            (p0)[index169_] = d_869_v21_
            index170_ = default__.safeIndex(600, (p0).length(0))
            (p0)[index170_] = d_869_v21_
            if (d_868_i1_) <= (d_868_i1_):
                d_870_v22_: _dafny.Seq
                d_870_v22_ = _dafny.SeqWithoutIsStrInference([d_867_v20_])
                d_871_v23_: C0
                nw138_ = C0()
                nw138_.ctor__(default__.fm0((p0)[default__.safeIndex(600, (p0).length(0))], globalState), (d_870_v22_ if (p0)[default__.safeIndex(600, (p0).length(0))] else _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(0) - (d_868_i1_) for d_872_i2_ in range(default__.abs(857))]), d_867_v20_])))
                d_871_v23_ = nw138_
                index171_ = default__.safeIndex(600, (p0).length(0))
                (p0)[index171_] = (d_871_v23_).f11
                d_873_v24_: _dafny.Seq
                d_873_v24_ = _dafny.SeqWithoutIsStrInference([not(d_869_v21_), default__.fm0(not((d_871_v23_).f11), globalState)])
                d_874_v25_: _dafny.Set
                d_874_v25_ = _dafny.Set({len(d_873_v24_), d_866_v19_, d_866_v19_})
                d_875_v27_: _dafny.Map
                d_875_v27_ = _dafny.Map({True: _dafny.Set({d_868_i1_, d_866_v19_, d_868_i1_, d_868_i1_, d_866_v19_})})
                d_876_v28_: _dafny.MultiSet
                d_876_v28_ = _dafny.MultiSet([d_868_i1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkfxh")))])
                d_877_v29_: _dafny.Map
                d_877_v29_ = _dafny.Map({(d_871_v23_).f11: (d_871_v23_).f11})
                d_878_v30_: _dafny.Map
                d_878_v30_ = _dafny.Map({-696: d_866_v19_})
                d_879_v31_: _dafny.Array
                nw139_ = _dafny.Array(None, 11)
                nw139_[int(0)] = d_874_v25_
                nw139_[int(1)] = d_874_v25_
                nw139_[int(2)] = (d_874_v25_) | (d_874_v25_)
                nw139_[int(3)] = (d_874_v25_ if (d_871_v23_).f11 else d_874_v25_)
                nw139_[int(4)] = (d_874_v25_) | (d_874_v25_)
                nw139_[int(5)] = d_874_v25_
                def iife53_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(347, 914):
                        d_880_v26_: int = compr_33_
                        if ((347) <= (d_880_v26_)) and ((d_880_v26_) < (914)):
                            coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_880_v26_, 690)]))
                    return _dafny.Set(coll33_)
                nw139_[int(6)] = iife53_()
                
                nw139_[int(7)] = d_874_v25_
                nw139_[int(8)] = d_874_v25_
                nw139_[int(9)] = ((d_875_v27_)[(d_871_v23_).f11] if ((d_871_v23_).f11) in (d_875_v27_) else d_874_v25_)
                nw139_[int(10)] = default__.fm35(default__.fm1(d_868_i1_, (p0)[default__.safeIndex(600, (p0).length(0))], d_868_i1_, globalState), d_876_v28_, default__.fm36(d_868_i1_, len(d_877_v29_), d_868_i1_, d_878_v30_, globalState), globalState)
                d_879_v31_ = nw139_
                d_879_v31_ = d_879_v31_
                d_867_v20_ = (_dafny.SeqWithoutIsStrInference([d_868_i1_ for d_881_i3_ in range(default__.abs(784))])) + (d_867_v20_)
                d_882_v32_: _dafny.Map
                d_882_v32_ = _dafny.Map({(p0)[default__.safeIndex(600, (p0).length(0))]: (d_873_v24_) + (d_873_v24_)})
                d_883_v33_: _dafny.Seq
                d_883_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwoqi"))
                d_884_v34_: str
                d_884_v34_ = _dafny.CodePoint('a')
                index172_ = default__.safeIndex(600, (p0).length(0))
                index173_ = default__.safeIndex(600, (p0).length(0))
                index174_ = default__.safeIndex(600, (p0).length(0))
                rhs193_ = not((d_871_v23_).f11)
                rhs194_ = default__.fm0((d_871_v23_).f11, globalState)
                rhs195_ = d_866_v19_
                rhs196_ = ((d_882_v32_)[(p0)[default__.safeIndex(600, (p0).length(0))]] if ((p0)[default__.safeIndex(600, (p0).length(0))]) in (d_882_v32_) else d_873_v24_)
                rhs197_ = ((d_883_v33_) + (d_883_v33_)) == ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_885_i4_ in range(default__.abs(211))])).set(default__.safeIndex(d_868_i1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_886_i4_ in range(default__.abs(211))]))), d_884_v34_))
                lhs133_ = p0
                lhs134_ = default__.safeIndex(600, (p0).length(0))
                lhs135_ = p0
                lhs136_ = default__.safeIndex(600, (p0).length(0))
                lhs137_ = globalState
                lhs138_ = p0
                lhs139_ = default__.safeIndex(600, (p0).length(0))
                lhs133_[lhs134_] = rhs193_
                lhs135_[lhs136_] = rhs194_
                lhs137_.f1 = rhs195_
                d_873_v24_ = rhs196_
                lhs138_[lhs139_] = rhs197_
            elif True:
                d_887_v35_: _dafny.Array
                nw140_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_887_v35_ = nw140_
                d_887_v35_ = d_887_v35_
                d_888_v36_: T1
                nw141_ = C3()
                nw141_.ctor__(True, d_869_v21_)
                d_888_v36_ = nw141_
                d_889_v37_: _dafny.Map
                d_889_v37_ = _dafny.Map({default__.fm3(d_869_v21_, (p0)[default__.safeIndex(600, (p0).length(0))], globalState): d_866_v19_})
                d_890_v38_: _dafny.Map
                d_890_v38_ = _dafny.Map({d_889_v37_: d_869_v21_})
                d_891_v39_: _dafny.Seq
                d_891_v39_ = _dafny.SeqWithoutIsStrInference([d_869_v21_, (p0)[default__.safeIndex(600, (p0).length(0))]])
                d_892_v40_: _dafny.MultiSet
                d_892_v40_ = _dafny.MultiSet([d_891_v39_, d_891_v39_, d_891_v39_])
                d_893_v41_: _dafny.Map
                d_893_v41_ = _dafny.Map({(self).fm30(d_890_v38_, d_892_v40_, globalState): d_888_v36_})
                d_888_v36_ = ((d_893_v41_)[(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_894_i5_ in range(default__.abs(318))])))] if ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_895_i5_ in range(default__.abs(318))])))) in (d_893_v41_) else d_888_v36_)
                d_896_v42_: _dafny.Array
                nw142_ = _dafny.Array(int(0), 6)
                d_896_v42_ = nw142_
                index175_ = default__.safeIndex(117, (d_896_v42_).length(0))
                (d_896_v42_)[index175_] = d_866_v19_
                index176_ = default__.safeIndex(117, (d_896_v42_).length(0))
                (d_896_v42_)[index176_] = d_866_v19_
                d_897_v43_: C3
                nw143_ = C3()
                nw143_.ctor__(False, True)
                d_897_v43_ = nw143_
                d_898_v44_: C2
                nw144_ = C2()
                nw144_.ctor__(_dafny.CodePoint('m'), d_869_v21_)
                d_898_v44_ = nw144_
            if not(default__.fm0(True, globalState)):
                d_899_v45_: C1
                nw145_ = C1()
                nw145_.ctor__()
                d_899_v45_ = nw145_
                def lambda38_(d_900_v21_, d_901_p0_, d_902_i1_, d_903_v19_):
                    def lambda39_(d_904_i6_):
                        return not((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_900_v21_, (d_901_p0_)[default__.safeIndex(600, (d_901_p0_).length(0))]])), d_902_i1_, 266, d_903_v19_])) != (_dafny.MultiSet([d_902_i1_])))

                    return lambda39_

                init20_ = lambda38_(d_869_v21_, p0, d_868_i1_, d_866_v19_)
                nw146_ = _dafny.Array(None, 26)
                for i0_20_ in range(nw146_.length(0)):
                    nw146_[i0_20_] = init20_(i0_20_)
                r2 = nw146_
                d_905_v46_: _dafny.Seq
                d_905_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xr"))
                index177_ = default__.safeIndex(600, (p0).length(0))
                (p0)[index177_] = (d_905_v46_) < (d_905_v46_)
                d_906_v47_: D0
                d_906_v47_ = D0_DC1()
                d_906_v47_ = d_906_v47_
                d_907_v48_: _dafny.Array
                def lambda40_(d_908_p0_):
                    def lambda41_(d_909_i7_):
                        return (d_908_p0_)[default__.safeIndex(600, (d_908_p0_).length(0))]

                    return lambda41_

                init21_ = lambda40_(p0)
                nw147_ = _dafny.Array(None, 25)
                for i0_21_ in range(nw147_.length(0)):
                    nw147_[i0_21_] = init21_(i0_21_)
                d_907_v48_ = nw147_
                d_910_v49_: _dafny.Seq
                d_910_v49_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(600, (p0).length(0))], False])
                d_911_v50_: _dafny.Seq
                d_911_v50_ = _dafny.SeqWithoutIsStrInference([True, (p0)[default__.safeIndex(600, (p0).length(0))], not((d_910_v49_)[default__.safeIndex((0) - (d_866_v19_), len(d_910_v49_))]), not(False)])
                d_912_v51_: _dafny.Map
                d_912_v51_ = _dafny.Map({d_907_v48_: (d_910_v49_)[default__.safeIndex(d_868_i1_, len(d_910_v49_))]})
                d_912_v51_ = (d_912_v51_).set(d_907_v48_, d_869_v21_)
            elif True:
                d_913_v52_: str
                d_913_v52_ = _dafny.CodePoint('x')
                d_914_v53_: _dafny.MultiSet
                d_914_v53_ = _dafny.MultiSet([(p0)[default__.safeIndex(600, (p0).length(0))]])
                d_915_v54_: C2
                nw148_ = C2()
                nw148_.ctor__(d_913_v52_, (d_914_v53_).issubset(d_914_v53_))
                d_915_v54_ = nw148_
                (globalState).f1 = d_868_i1_
                d_916_v55_: C2
                nw149_ = C2()
                nw149_.ctor__(d_913_v52_, (d_866_v19_) < (d_866_v19_))
                d_916_v55_ = nw149_
                d_917_v56_: D2
                d_917_v56_ = D2_DC7(default__.safeDivisionInt(d_868_i1_, len((d_867_v20_).set(default__.safeIndex(646, len(d_867_v20_)), (0) - (d_868_i1_)))), d_868_i1_)
                d_918_v57_: _dafny.Seq
                d_918_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tecdwogxu"))
                rhs198_ = d_866_v19_
                rhs199_ = d_917_v56_
                rhs200_ = d_918_v57_
                rhs201_ = d_869_v21_
                rhs202_ = (d_913_v52_ if d_869_v21_ else d_913_v52_)
                lhs140_ = globalState
                lhs140_.f1 = rhs198_
                d_917_v56_ = rhs199_
                d_918_v57_ = rhs200_
                d_869_v21_ = rhs201_
                d_913_v52_ = rhs202_
                d_913_v52_ = d_913_v52_
            d_919_v58_: _dafny.Array
            def lambda42_(d_920_v21_):
                def lambda43_(d_921_i8_):
                    return d_920_v21_

                return lambda43_

            init22_ = lambda42_(d_869_v21_)
            nw150_ = _dafny.Array(None, 6)
            for i0_22_ in range(nw150_.length(0)):
                nw150_[i0_22_] = init22_(i0_22_)
            d_919_v58_ = nw150_
            d_922_v59_: _dafny.Map
            d_922_v59_ = _dafny.Map({d_919_v58_: (p0)[default__.safeIndex(600, (p0).length(0))]})
            d_923_v60_: C3
            nw151_ = C3()
            nw151_.ctor__((p0)[default__.safeIndex(600, (p0).length(0))], ((d_922_v59_)[d_919_v58_] if (d_919_v58_) in (d_922_v59_) else (p0)[default__.safeIndex(600, (p0).length(0))]))
            d_923_v60_ = nw151_
        d_924_v61_: bool
        d_924_v61_ = False
        d_925_v62_: T1
        nw152_ = C3()
        nw152_.ctor__(d_924_v61_, d_924_v61_)
        d_925_v62_ = nw152_
        hi5_ = d_866_v19_
        for d_926_i9_ in range(len(_dafny.Map({d_925_v62_: _dafny.CodePoint('l')})), hi5_):
            d_927_v63_: _dafny.Seq
            d_927_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amn"))
            d_928_v64_: _dafny.Map
            d_928_v64_ = _dafny.Map({d_926_i9_: len(d_927_v63_)})
            d_929_v65_: _dafny.Map
            d_929_v65_ = _dafny.Map({d_927_v63_: d_926_i9_})
            (globalState).f1 = (len(d_928_v64_)) + (default__.safeModuloInt(d_926_i9_, len(d_929_v65_)))
            d_866_v19_ = d_866_v19_
            d_930_v66_: _dafny.Array
            def lambda44_(d_931_i10_):
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_932_i11_ in range(default__.abs(924))])

            init23_ = lambda44_
            nw153_ = _dafny.Array(None, 24)
            for i0_23_ in range(nw153_.length(0)):
                nw153_[i0_23_] = init23_(i0_23_)
            d_930_v66_ = nw153_
            index178_ = default__.safeIndex(547, (d_930_v66_).length(0))
            (d_930_v66_)[index178_] = d_927_v63_
            d_933_v67_: _dafny.Seq
            d_933_v67_ = _dafny.SeqWithoutIsStrInference([not(d_924_v61_)])
            d_934_v68_: _dafny.MultiSet
            d_934_v68_ = _dafny.MultiSet([d_866_v19_])
            index179_ = default__.safeIndex(547, (d_930_v66_).length(0))
            (d_930_v66_)[index179_] = default__.fm1((len(d_933_v67_)) + (((d_934_v68_)[d_926_i9_] if (d_926_i9_) in (d_934_v68_) else d_926_i9_)), d_924_v61_, d_926_i9_, globalState)
            d_935_v69_: _dafny.Map
            d_935_v69_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "buasua"))): False})
            d_936_v70_: _dafny.Map
            d_936_v70_ = _dafny.Map({d_924_v61_: True})
            d_937_v71_: str
            d_937_v71_ = _dafny.CodePoint('t')
            d_938_v72_: _dafny.Map
            d_938_v72_ = _dafny.Map({d_924_v61_: d_937_v71_})
            d_939_v73_: _dafny.Array
            nw154_ = _dafny.Array(None, 19)
            nw154_[int(0)] = d_926_i9_
            nw154_[int(1)] = d_926_i9_
            nw154_[int(2)] = d_866_v19_
            nw154_[int(3)] = d_866_v19_
            nw154_[int(4)] = len(d_867_v20_)
            nw154_[int(5)] = (0) - (d_926_i9_)
            nw154_[int(6)] = d_866_v19_
            nw154_[int(7)] = d_866_v19_
            nw154_[int(8)] = d_866_v19_
            nw154_[int(9)] = (d_925_v62_).fm4(_dafny.SeqWithoutIsStrInference([((d_935_v69_)[len(d_936_v70_)] if (len(d_936_v70_)) in (d_935_v69_) else d_924_v61_)]), d_926_i9_, 301, d_934_v68_, globalState)
            nw154_[int(10)] = d_926_i9_
            nw154_[int(11)] = len(d_938_v72_)
            nw154_[int(12)] = d_866_v19_
            nw154_[int(13)] = d_866_v19_
            nw154_[int(14)] = d_926_i9_
            nw154_[int(15)] = (0) - (d_926_i9_)
            nw154_[int(16)] = d_866_v19_
            nw154_[int(17)] = len(_dafny.SeqWithoutIsStrInference([True, d_924_v61_, True, d_924_v61_]))
            nw154_[int(18)] = d_926_i9_
            d_939_v73_ = nw154_
            d_940_v74_: _dafny.Seq
            d_940_v74_ = _dafny.SeqWithoutIsStrInference([(d_939_v73_ if not(d_924_v61_) else d_939_v73_), d_939_v73_, d_939_v73_, d_939_v73_])
            d_940_v74_ = ((d_940_v74_) + (d_940_v74_)) + (_dafny.SeqWithoutIsStrInference([d_939_v73_, d_939_v73_, (d_940_v74_)[default__.safeIndex(d_866_v19_, len(d_940_v74_))], d_939_v73_, d_939_v73_]))
        (globalState).f1 = (0) - (default__.safeModuloInt(d_866_v19_, d_866_v19_))
        if (d_924_v61_) or (d_924_v61_):
            d_941_v75_: _dafny.Seq
            d_941_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ig"))
            d_941_v75_ = d_941_v75_
            (globalState).f1 = d_866_v19_
            d_866_v19_ = default__.safeDivisionInt(default__.safeDivisionInt(d_866_v19_, d_866_v19_), d_866_v19_)
            (globalState).f1 = d_866_v19_
            d_942_v76_: _dafny.Set
            d_942_v76_ = _dafny.Set({d_924_v61_, ((0) - (d_866_v19_)) != (d_866_v19_)})
            d_942_v76_ = d_942_v76_
        elif True:
            d_924_v61_ = d_924_v61_
            d_943_v77_: _dafny.Map
            d_943_v77_ = _dafny.Map({d_866_v19_: d_866_v19_})
            d_944_v78_: _dafny.Map
            d_944_v78_ = _dafny.Map({d_943_v77_: d_924_v61_})
            d_945_v79_: _dafny.Seq
            d_945_v79_ = _dafny.SeqWithoutIsStrInference([False, d_924_v61_])
            d_946_v80_: _dafny.MultiSet
            d_946_v80_ = _dafny.MultiSet([d_945_v79_])
            d_866_v19_ = ((d_866_v19_) + ((0) - (d_866_v19_))) * ((self).fm30(d_944_v78_, d_946_v80_, globalState))
            d_947_v81_: _dafny.Seq
            d_947_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcfaexkop"))
            d_948_v82_: _dafny.Array
            def lambda45_(d_949_v19_):
                def lambda46_(d_950_i12_):
                    return (d_950_i12_) + (d_949_v19_)

                return lambda46_

            init24_ = lambda45_(d_866_v19_)
            nw155_ = _dafny.Array(None, 5)
            for i0_24_ in range(nw155_.length(0)):
                nw155_[i0_24_] = init24_(i0_24_)
            d_948_v82_ = nw155_
            d_951_v83_: _dafny.Array
            d_951_v83_ = d_948_v82_
            d_952_v84_: _dafny.Map
            d_952_v84_ = _dafny.Map({d_924_v61_: d_951_v83_})
            d_953_v85_: _dafny.MultiSet
            d_953_v85_ = _dafny.MultiSet([len(d_952_v84_)])
            d_954_v86_: D2
            d_954_v86_ = D2_DC7((self).fm4(d_945_v79_, d_866_v19_, len(d_945_v79_), (d_953_v85_).set(d_866_v19_, default__.abs(d_866_v19_)), globalState), d_866_v19_)
            d_955_v87_: _dafny.Map
            d_955_v87_ = _dafny.Map({d_947_v81_: (d_954_v86_).cf7})
            d_955_v87_ = (d_955_v87_).set(d_947_v81_, d_866_v19_)
            d_956_v88_: _dafny.Map
            d_956_v88_ = _dafny.Map({len((d_947_v81_) + (default__.fm1(d_866_v19_, d_924_v61_, 726, globalState))): True})
            d_956_v88_ = (d_956_v88_).set(d_866_v19_, d_924_v61_)
            (globalState).f1 = -63
        d_957_v89_: _dafny.Array
        nw156_ = _dafny.Array(None, 21)
        d_957_v89_ = nw156_
        d_957_v89_ = d_957_v89_
        d_958_v90_: str
        d_958_v90_ = _dafny.CodePoint('a')
        d_959_v91_: C0
        nw157_ = C0()
        nw157_.ctor__(d_924_v61_, (_dafny.SeqWithoutIsStrInference([d_867_v20_]) if True else default__.fm38(default__.fm0(d_924_v61_, globalState), d_924_v61_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bj"))), d_958_v90_, globalState)))
        d_959_v91_ = nw157_
        r0 = d_959_v91_
        r1 = p0
        r2 = p0
        return r0, r1, r2

    def m11(self, p0, p1, p2, globalState):
        d_960_v0_: _dafny.Seq
        d_960_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "latlvtxl"))
        d_961_v1_: D2
        d_961_v1_ = D2_DC6(d_960_v0_)
        source13_ = d_961_v1_
        if source13_.is_DC7:
            d_962___mcc_h0_ = source13_.cf7
            d_963___mcc_h1_ = source13_.cf8
            d_964_cf8_ = d_963___mcc_h1_
            d_965_cf7_ = d_962___mcc_h0_
            d_966_v2_: D2
            d_966_v2_ = D2_DC7(p1, (0) - (d_964_cf8_))
            (globalState).f1 = ((d_966_v2_).cf8) * (default__.safeDivisionInt(305, d_965_cf7_))
            d_967_v3_: _dafny.Array
            nw158_ = _dafny.Array(_dafny.Seq({}), 23)
            d_967_v3_ = nw158_
            d_967_v3_ = d_967_v3_
            d_968_v4_: D2
            d_968_v4_ = D2_DC8(p2)
            pat_let_tv13_ = d_964_cf8_
            def iife54_(_pat_let10_0):
                def iife55_(d_969_dt__update__tmp_h0_):
                    def iife56_(_pat_let11_0):
                        def iife57_(d_970_dt__update_hcf4_h0_):
                            return D1_DC4(d_970_dt__update_hcf4_h0_)
                        return iife57_(_pat_let11_0)
                    return iife56_(pat_let_tv13_)
                return iife55_(_pat_let10_0)
            source14_ = iife54_((self).fm5(d_968_v4_, globalState))
            if source14_.is_DC5:
                d_971___mcc_h5_ = source14_.cf5
                d_972_cf5_ = d_971___mcc_h5_
                (globalState).f1 = default__.fm3(p2, p2, globalState)
                rhs203_ = p2
                rhs204_ = default__.safeDivisionInt(d_965_cf7_, default__.safeModuloInt((0) - (d_965_cf7_), (0) - (len(d_960_v0_))))
                lhs141_ = globalState
                d_972_cf5_ = rhs203_
                lhs141_.f1 = rhs204_
                d_973_v5_: str
                d_973_v5_ = _dafny.CodePoint('c')
                d_974_v6_: _dafny.Map
                d_974_v6_ = _dafny.Map({d_973_v5_: d_964_cf8_})
                d_974_v6_ = d_974_v6_
                (globalState).f1 = p1
            elif True:
                d_975___mcc_h6_ = source14_.cf4
                d_976_cf4_ = d_975___mcc_h6_
                d_977_v7_: bool
                d_977_v7_ = False
                d_977_v7_ = p2
                d_978_v8_: D0
                d_978_v8_ = D0_DC0(p2)
                d_979_v9_: D0
                d_979_v9_ = D0_DC3(d_978_v8_)
                d_980_v10_: D0
                d_980_v10_ = D0_DC3(d_979_v9_)
                d_981_v11_: D0
                d_981_v11_ = D0_DC3(d_978_v8_)
                d_982_v12_: _dafny.Seq
                d_982_v12_ = _dafny.SeqWithoutIsStrInference([d_977_v7_, True, False])
                d_983_v13_: _dafny.Seq
                d_983_v13_ = _dafny.SeqWithoutIsStrInference([(D0_DC3(d_978_v8_)).cf3])
                d_984_v14_: _dafny.Array
                def lambda47_(d_985_i0_):
                    return default__.safeDivisionInt(d_985_i0_, 864)

                init25_ = lambda47_
                nw159_ = _dafny.Array(None, 10)
                for i0_25_ in range(nw159_.length(0)):
                    nw159_[i0_25_] = init25_(i0_25_)
                d_984_v14_ = nw159_
                d_986_v15_: _dafny.Map
                d_986_v15_ = _dafny.Map({d_984_v14_: p1})
                pat_let_tv14_ = d_983_v13_
                pat_let_tv15_ = d_976_cf4_
                pat_let_tv16_ = d_983_v13_
                def iife58_(_pat_let12_0):
                    def iife59_(d_987_dt__update__tmp_h1_):
                        def iife60_(_pat_let13_0):
                            def iife61_(d_988_dt__update_hcf3_h0_):
                                return D0_DC3(d_988_dt__update_hcf3_h0_)
                            return iife61_(_pat_let13_0)
                        return iife60_((pat_let_tv14_)[default__.safeIndex(pat_let_tv15_, len(pat_let_tv16_))])
                    return iife59_(_pat_let12_0)
                rhs205_ = iife58_(d_981_v11_)
                rhs206_ = ((0) - (len(d_986_v15_))) + (d_976_cf4_)
                rhs207_ = d_982_v12_
                d_981_v11_ = rhs205_
                d_964_cf8_ = rhs206_
                d_982_v12_ = rhs207_
                d_977_v7_ = p2
                d_960_v0_ = d_960_v0_
            d_989_v16_: D0
            d_989_v16_ = D0_DC1()
            d_990_v17_: _dafny.Array
            def lambda48_(d_991_v0_):
                def lambda49_(d_992_i1_):
                    return ((D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgtgm")))).cf6) + (d_991_v0_)

                return lambda49_

            init26_ = lambda48_(d_960_v0_)
            nw160_ = _dafny.Array(None, 14)
            for i0_26_ in range(nw160_.length(0)):
                nw160_[i0_26_] = init26_(i0_26_)
            d_990_v17_ = nw160_
            d_993_v18_: str
            d_993_v18_ = _dafny.CodePoint('o')
            index180_ = default__.safeIndex(605, (d_990_v17_).length(0))
            (d_990_v17_)[index180_] = ((default__.fm1(p1, p0, len(_dafny.SeqWithoutIsStrInference([(D11_DC19(_dafny.CodePoint('q'))).cf20 for d_994_i2_ in range(default__.abs(-826))])), globalState)).set(default__.safeIndex(p1, len(default__.fm1(p1, p0, len(_dafny.SeqWithoutIsStrInference([(D11_DC19(_dafny.CodePoint('q'))).cf20 for d_995_i2_ in range(default__.abs(-826))])), globalState))), d_993_v18_)) + (_dafny.SeqWithoutIsStrInference([d_993_v18_ for d_996_i3_ in range(default__.abs(699))]))
            d_997_v19_: _dafny.Map
            d_997_v19_ = _dafny.Map({d_964_cf8_: p2})
            index181_ = default__.safeIndex(605, (d_990_v17_).length(0))
            rhs208_ = (0) - (default__.safeModuloInt(default__.safeModuloInt(p1, d_965_cf7_), (len(d_997_v19_)) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfndmpit"))))))
            rhs209_ = d_989_v16_
            rhs210_ = d_960_v0_
            lhs142_ = globalState
            lhs143_ = d_990_v17_
            lhs144_ = default__.safeIndex(605, (d_990_v17_).length(0))
            lhs142_.f1 = rhs208_
            d_989_v16_ = rhs209_
            lhs143_[lhs144_] = rhs210_
        elif source13_.is_DC8:
            d_998___mcc_h2_ = source13_.cf9
            d_999_cf9_ = d_998___mcc_h2_
            d_1000_v20_: _dafny.Map
            d_1000_v20_ = _dafny.Map({d_999_cf9_: p2})
            if ((d_1000_v20_)[p2] if (p2) in (d_1000_v20_) else p2):
                d_1001_v21_: str
                d_1001_v21_ = _dafny.CodePoint('k')
                d_1002_v22_: _dafny.Map
                d_1002_v22_ = _dafny.Map({d_960_v0_: d_1001_v21_})
                (globalState).f1 = len(d_1002_v22_)
                d_1003_v23_: _dafny.Map
                d_1003_v23_ = _dafny.Map({p2: p1})
                d_1004_v24_: D11
                d_1004_v24_ = D11_DC19(d_1001_v21_)
                rhs211_ = (d_1003_v23_) | (d_1003_v23_)
                rhs212_ = D11_DC19(d_1001_v21_)
                rhs213_ = d_960_v0_
                rhs214_ = ((d_960_v0_) + ((D2_DC6(d_960_v0_)).cf6)).set(default__.safeIndex(p1, len((d_960_v0_) + ((D2_DC6(d_960_v0_)).cf6))), d_1001_v21_)
                d_1003_v23_ = rhs211_
                d_1004_v24_ = rhs212_
                d_960_v0_ = rhs213_
                d_960_v0_ = rhs214_
                d_1005_v25_: _dafny.Array
                def lambda50_(d_1006_p0_):
                    def lambda51_(d_1007_i4_):
                        return d_1006_p0_

                    return lambda51_

                init27_ = lambda50_(p0)
                nw161_ = _dafny.Array(None, 8)
                for i0_27_ in range(nw161_.length(0)):
                    nw161_[i0_27_] = init27_(i0_27_)
                d_1005_v25_ = nw161_
                d_1008_v26_: _dafny.MultiSet
                d_1008_v26_ = _dafny.MultiSet([False, p0, p0])
                index182_ = default__.safeIndex(297, (d_1005_v25_).length(0))
                (d_1005_v25_)[index182_] = (d_1008_v26_).ispropersubset(default__.fm34(p2, globalState))
                d_1009_v27_: _dafny.Map
                d_1009_v27_ = _dafny.Map({d_960_v0_: (len(d_1000_v20_)) != (p1)})
                index183_ = default__.safeIndex(297, (d_1005_v25_).length(0))
                (d_1005_v25_)[index183_] = ((d_1009_v27_)[_dafny.SeqWithoutIsStrInference([d_1001_v21_ for d_1010_i5_ in range(default__.abs(41))])] if (_dafny.SeqWithoutIsStrInference([d_1001_v21_ for d_1011_i5_ in range(default__.abs(41))])) in (d_1009_v27_) else d_999_cf9_)
                (globalState).f1 = (p1) * (p1)
                d_1012_v28_: _dafny.Map
                d_1012_v28_ = _dafny.Map({d_1005_v25_: d_1001_v21_})
                d_1013_v29_: D0
                d_1013_v29_ = D0_DC2(d_1012_v28_, d_1001_v21_)
                d_1014_v30_: _dafny.Map
                d_1014_v30_ = _dafny.Map({d_1013_v29_: d_999_cf9_})
                d_1014_v30_ = d_1014_v30_
            elif True:
                (globalState).f1 = default__.fm3(p0, d_999_cf9_, globalState)
                d_1015_v31_: _dafny.Array
                nw162_ = _dafny.Array(False, 19)
                d_1015_v31_ = nw162_
                d_1016_v32_: str
                d_1016_v32_ = _dafny.CodePoint('j')
                d_1017_v33_: D0
                d_1018_v34_: bool
                out30_: D0
                out31_: bool
                out30_, out31_ = default__.m0(d_1015_v31_, d_1016_v32_, p0, d_1016_v32_, globalState)
                d_1017_v33_ = out30_
                d_1018_v34_ = out31_
                (globalState).f1 = (0) - (p1)
                d_1019_v35_: _dafny.Map
                d_1019_v35_ = _dafny.Map({d_960_v0_: p0})
                d_1020_v36_: D0
                d_1021_v37_: bool
                out32_: D0
                out33_: bool
                out32_, out33_ = default__.m0(d_1015_v31_, d_1016_v32_, ((d_1019_v35_)[d_960_v0_] if (d_960_v0_) in (d_1019_v35_) else d_1018_v34_), d_1016_v32_, globalState)
                d_1020_v36_ = out32_
                d_1021_v37_ = out33_
                d_1022_v38_: _dafny.Map
                d_1022_v38_ = _dafny.Map({p2: 699})
                d_1022_v38_ = (d_1022_v38_).set(d_999_cf9_, len(d_960_v0_))
            d_1023_v39_: _dafny.Map
            d_1023_v39_ = _dafny.Map({p1: p1})
            d_1024_v40_: _dafny.Map
            d_1024_v40_ = d_1023_v39_
            d_1025_v41_: D6
            d_1025_v41_ = D6_DC14(False)
            d_1026_v42_: _dafny.Map
            d_1026_v42_ = _dafny.Map({(d_1024_v40_): (d_1025_v41_).cf15})
            d_1027_v43_: _dafny.Seq
            d_1027_v43_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1028_v44_: _dafny.MultiSet
            d_1028_v44_ = _dafny.MultiSet([d_1027_v43_])
            d_1029_v45_: _dafny.MultiSet
            d_1029_v45_ = _dafny.MultiSet([(self).fm30(d_1026_v42_, d_1028_v44_, globalState), p1])
            d_1030_v46_: _dafny.Array
            nw163_ = _dafny.Array(None, 14)
            nw163_[int(0)] = p1
            nw163_[int(1)] = p1
            nw163_[int(2)] = p1
            nw163_[int(3)] = p1
            nw163_[int(4)] = p1
            nw163_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rodcjqa")))
            nw163_[int(6)] = (0) - (p1)
            nw163_[int(7)] = 96
            nw163_[int(8)] = p1
            nw163_[int(9)] = p1
            nw163_[int(10)] = p1
            nw163_[int(11)] = len(d_960_v0_)
            nw163_[int(12)] = p1
            nw163_[int(13)] = p1
            d_1030_v46_ = nw163_
            d_1031_v47_: _dafny.Map
            d_1031_v47_ = _dafny.Map({d_1030_v46_: False})
            d_1029_v45_ = _dafny.MultiSet([len(d_1031_v47_)])
            d_1032_v48_: _dafny.Seq
            d_1032_v48_ = _dafny.SeqWithoutIsStrInference([d_960_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")), (d_960_v0_) + (d_960_v0_)])
            d_1032_v48_ = _dafny.SeqWithoutIsStrInference([(d_960_v0_).set(default__.safeIndex(p1, len(d_960_v0_)), _dafny.CodePoint('o')) for d_1033_i6_ in range(default__.abs(-320))])
            d_1034_v49_: _dafny.Seq
            d_1034_v49_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1034_v49_ = d_1034_v49_
        elif source13_.is_DC6:
            d_1035___mcc_h3_ = source13_.cf6
            d_1036_cf6_ = d_1035___mcc_h3_
            d_1037_v50_: bool
            d_1037_v50_ = True
            d_1037_v50_ = d_1037_v50_
            d_1038_v51_: _dafny.Map
            d_1038_v51_ = _dafny.Map({d_1037_v50_: p0})
            d_1039_v52_: D12
            d_1039_v52_ = D12_DC21(d_1038_v51_)
            d_1040_v53_: _dafny.Set
            d_1040_v53_ = _dafny.Set({d_1038_v51_, (d_1039_v52_).cf23})
            (globalState).f1 = len((d_1040_v53_).intersection(d_1040_v53_))
            d_1041_v54_: C3
            nw164_ = C3()
            nw164_.ctor__(p2, True)
            d_1041_v54_ = nw164_
            d_1042_v55_: _dafny.Seq
            d_1042_v55_ = _dafny.SeqWithoutIsStrInference([p2, (d_1041_v54_).f13, (d_1041_v54_).f14])
            d_1042_v55_ = (d_1042_v55_) + (d_1042_v55_)
        elif True:
            d_1043___mcc_h4_ = source13_.cf10
            d_1044_cf10_ = d_1043___mcc_h4_
            d_1045_v56_: _dafny.Seq
            d_1045_v56_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1045_v56_ = _dafny.SeqWithoutIsStrInference([p2, p0, p0, p2])
            d_1046_v57_: str
            d_1046_v57_ = _dafny.CodePoint('p')
            d_1047_v58_: _dafny.Array
            nw165_ = _dafny.Array(None, 9)
            nw165_[int(0)] = True
            nw165_[int(1)] = p0
            nw165_[int(2)] = p2
            nw165_[int(3)] = (p1) > (p1)
            nw165_[int(4)] = p0
            nw165_[int(5)] = not (True) or (p2)
            nw165_[int(6)] = not(not(not((d_1046_v57_) in (d_960_v0_))))
            nw165_[int(7)] = p2
            nw165_[int(8)] = (not(p0) if p0 else p0)
            d_1047_v58_ = nw165_
            index184_ = default__.safeIndex(384, (d_1047_v58_).length(0))
            (d_1047_v58_)[index184_] = p2
            index185_ = default__.safeIndex(384, (d_1047_v58_).length(0))
            rhs215_ = p0
            rhs216_ = d_1046_v57_
            lhs145_ = d_1047_v58_
            lhs146_ = default__.safeIndex(384, (d_1047_v58_).length(0))
            lhs145_[lhs146_] = rhs215_
            d_1046_v57_ = rhs216_
            index186_ = default__.safeIndex(384, (d_1047_v58_).length(0))
            (d_1047_v58_)[index186_] = (d_1045_v56_)[default__.safeIndex(p1, len(d_1045_v56_))]
            d_1048_v59_: _dafny.Seq
            d_1048_v59_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1049_v60_: _dafny.Seq
            d_1049_v60_ = _dafny.SeqWithoutIsStrInference([d_1048_v59_, _dafny.SeqWithoutIsStrInference([743])])
            d_1050_v61_: C0
            nw166_ = C0()
            nw166_.ctor__((d_1047_v58_)[default__.safeIndex(384, (d_1047_v58_).length(0))], d_1049_v60_)
            d_1050_v61_ = nw166_
        d_1051_i7_: int
        d_1051_i7_ = 0
        with _dafny.label("13"):
            while (p2) == (p0):
                with _dafny.c_label("13"):
                    if (d_1051_i7_) >= (100):
                        raise _dafny.Break("13")
                    d_1051_i7_ = (d_1051_i7_) + (1)
                    d_1052_v62_: bool
                    d_1052_v62_ = False
                    d_1052_v62_ = p2
                    d_1053_v63_: _dafny.MultiSet
                    d_1053_v63_ = _dafny.MultiSet([_dafny.Set({d_1052_v62_, p2, p0})])
                    d_1054_v64_: D11
                    d_1054_v64_ = D11_DC20(d_1052_v62_, p1)
                    d_1055_v65_: _dafny.Set
                    d_1055_v65_ = _dafny.Set({p1, (0) - ((d_1054_v64_).cf22)})
                    d_1056_v66_: _dafny.Map
                    d_1056_v66_ = _dafny.Map({(d_1053_v63_).isdisjoint(d_1053_v63_): len(d_1055_v65_)})
                    d_1056_v66_ = (d_1056_v66_).set(d_1052_v62_, p1)
                    d_961_v1_ = D2_DC6(d_960_v0_)
                    (globalState).f1 = p1
                    pass
            pass
        if not (p0) or (p2):
            d_1057_v67_: str
            d_1057_v67_ = _dafny.CodePoint('f')
            d_1058_v68_: C2
            nw167_ = C2()
            nw167_.ctor__(d_1057_v67_, True)
            d_1058_v68_ = nw167_
            d_1059_v69_: D0
            d_1059_v69_ = D0_DC0(True)
            d_1060_v70_: D0
            d_1060_v70_ = D0_DC3(d_1059_v69_)
            pat_let_tv17_ = d_1059_v69_
            def iife62_(_pat_let14_0):
                def iife63_(d_1061_dt__update__tmp_h2_):
                    def iife64_(_pat_let15_0):
                        def iife65_(d_1062_dt__update_hcf3_h1_):
                            return D0_DC3(d_1062_dt__update_hcf3_h1_)
                        return iife65_(_pat_let15_0)
                    return iife64_(pat_let_tv17_)
                return iife63_(_pat_let14_0)
            source15_ = iife62_(d_1060_v70_)
            if source15_.is_DC1:
                d_1063_v71_: _dafny.Map
                d_1063_v71_ = _dafny.Map({p2: p2})
                d_1064_v72_: D12
                d_1064_v72_ = D12_DC21(d_1063_v71_)
                d_1064_v72_ = d_1064_v72_
                d_1065_v73_: _dafny.Map
                d_1065_v73_ = _dafny.Map({p1: default__.fm35(d_960_v0_, _dafny.MultiSet([p1, (0) - (-114)]), d_1057_v67_, globalState)})
                d_1066_v74_: _dafny.Set
                d_1066_v74_ = _dafny.Set({p1, p1})
                d_1065_v73_ = (d_1065_v73_).set(len((d_960_v0_) + (d_960_v0_)), d_1066_v74_)
                d_1067_v75_: bool
                d_1067_v75_ = True
                d_1067_v75_ = ((p2 if p2 else p0)) or (p2)
                d_1068_v76_: _dafny.Array
                def lambda52_(d_1069_p1_):
                    def lambda53_(d_1070_i8_):
                        return default__.safeModuloInt(d_1070_i8_, (0) - (d_1069_p1_))

                    return lambda53_

                init28_ = lambda52_(p1)
                nw168_ = _dafny.Array(None, 18)
                for i0_28_ in range(nw168_.length(0)):
                    nw168_[i0_28_] = init28_(i0_28_)
                d_1068_v76_ = nw168_
                index187_ = default__.safeIndex(85, (d_1068_v76_).length(0))
                (d_1068_v76_)[index187_] = (0) - (p1)
                index188_ = default__.safeIndex(85, (d_1068_v76_).length(0))
                (d_1068_v76_)[index188_] = p1
            elif source15_.is_DC2:
                d_1071___mcc_h7_ = source15_.cf1
                d_1072___mcc_h8_ = source15_.cf2
                d_1073_cf2_ = d_1072___mcc_h8_
                d_1074_cf1_ = d_1071___mcc_h7_
                d_1075_v77_: _dafny.Array
                nw169_ = _dafny.Array(None, 7)
                nw169_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1073_cf2_ for d_1076_i9_ in range(default__.abs(421))])
                nw169_[int(1)] = d_960_v0_
                nw169_[int(2)] = d_960_v0_
                nw169_[int(3)] = (d_960_v0_).set(default__.safeIndex(-556, len(d_960_v0_)), d_1073_cf2_)
                nw169_[int(4)] = (d_960_v0_) + (d_960_v0_)
                nw169_[int(5)] = d_960_v0_
                nw169_[int(6)] = d_960_v0_
                d_1075_v77_ = nw169_
                index189_ = default__.safeIndex(441, (d_1075_v77_).length(0))
                (d_1075_v77_)[index189_] = (_dafny.SeqWithoutIsStrInference([d_1057_v67_ for d_1077_i10_ in range(default__.abs(-35))])) + (d_960_v0_)
                index190_ = default__.safeIndex(441, (d_1075_v77_).length(0))
                (d_1075_v77_)[index190_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "remd"))) + (d_960_v0_)
                d_1078_v78_: bool
                d_1078_v78_ = False
                d_1079_v79_: _dafny.Seq
                d_1079_v79_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1080_v80_: _dafny.Seq
                d_1080_v80_ = _dafny.SeqWithoutIsStrInference([not(default__.fm0(d_1078_v78_, globalState)), False, not(p0), d_1078_v78_])
                d_1081_v81_: _dafny.Array
                nw170_ = _dafny.Array(None, 13)
                nw170_[int(0)] = d_1079_v79_
                nw170_[int(1)] = d_1079_v79_
                nw170_[int(2)] = d_1079_v79_
                nw170_[int(3)] = d_1079_v79_
                nw170_[int(4)] = d_1079_v79_
                nw170_[int(5)] = d_1079_v79_
                nw170_[int(6)] = d_1079_v79_
                nw170_[int(7)] = d_1080_v80_
                nw170_[int(8)] = d_1079_v79_
                nw170_[int(9)] = default__.fm39(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xun"))), p0, p1, p1, globalState)
                nw170_[int(10)] = d_1079_v79_
                nw170_[int(11)] = d_1079_v79_
                nw170_[int(12)] = d_1079_v79_
                d_1081_v81_ = nw170_
                d_1082_v82_: _dafny.Set
                d_1082_v82_ = _dafny.Set({len((d_1075_v77_)[default__.safeIndex(441, (d_1075_v77_).length(0))])})
                rhs217_ = (d_1058_v68_).fm8(True, d_1078_v78_, (p1) == (len(d_1082_v82_)), globalState)
                rhs218_ = (p1) != ((p1) + (p1))
                rhs219_ = d_1081_v81_
                lhs147_ = globalState
                lhs147_.f1 = rhs217_
                d_1078_v78_ = rhs218_
                d_1081_v81_ = rhs219_
                d_1083_v83_: _dafny.Array
                nw171_ = _dafny.Array(int(0), 15)
                d_1083_v83_ = nw171_
                index191_ = default__.safeIndex(827, (d_1083_v83_).length(0))
                (d_1083_v83_)[index191_] = p1
                index192_ = default__.safeIndex(827, (d_1083_v83_).length(0))
                (d_1083_v83_)[index192_] = (len((d_1075_v77_)[default__.safeIndex(441, (d_1075_v77_).length(0))])) * (p1)
                d_1084_v84_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.Set({}), 4)
                d_1084_v84_ = nw172_
                index193_ = default__.safeIndex(140, (d_1084_v84_).length(0))
                (d_1084_v84_)[index193_] = d_1082_v82_
                d_1085_v86_: _dafny.Map
                d_1085_v86_ = _dafny.Map({(d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))]: (d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))]})
                d_1086_v87_: _dafny.Map
                d_1086_v87_ = _dafny.Map({d_1085_v86_: d_1078_v78_})
                d_1087_v88_: _dafny.Seq
                d_1087_v88_ = _dafny.SeqWithoutIsStrInference([d_1080_v80_])
                d_1088_v89_: _dafny.MultiSet
                d_1088_v89_ = _dafny.MultiSet([(d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))], p1, (d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))], p1, p1])
                d_1089_v90_: _dafny.Seq
                d_1089_v90_ = _dafny.SeqWithoutIsStrInference([(d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))], p1])
                index194_ = default__.safeIndex(827, (d_1083_v83_).length(0))
                index195_ = default__.safeIndex(140, (d_1084_v84_).length(0))
                def iife66_():
                    coll34_ = _dafny.Map()
                    compr_34_: _dafny.Map
                    for compr_34_ in (d_1086_v87_).keys.Elements:
                        d_1090_v85_: _dafny.Map = compr_34_
                        if (d_1090_v85_) in (d_1086_v87_):
                            coll34_[d_1090_v85_] = d_1078_v78_
                    return _dafny.Map(coll34_)
                rhs220_ = default__.safeModuloInt(p1, ((self).fm30(iife66_()
                , _dafny.MultiSet(d_1087_v88_), globalState)) * (362))
                rhs221_ = default__.fm35(d_960_v0_, d_1088_v89_, d_1057_v67_, globalState)
                rhs222_ = ((d_1088_v89_)[((d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))] if d_1078_v78_ else (d_1088_v89_).cardinality)] if (((d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))] if d_1078_v78_ else (d_1088_v89_).cardinality)) in (d_1088_v89_) else default__.safeDivisionInt(191, p1))
                rhs223_ = not(((d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))]) == ((d_1089_v90_)[default__.safeIndex((d_1083_v83_)[default__.safeIndex(827, (d_1083_v83_).length(0))], len(d_1089_v90_))]))
                lhs148_ = d_1083_v83_
                lhs149_ = default__.safeIndex(827, (d_1083_v83_).length(0))
                lhs150_ = d_1084_v84_
                lhs151_ = default__.safeIndex(140, (d_1084_v84_).length(0))
                lhs152_ = globalState
                lhs148_[lhs149_] = rhs220_
                lhs150_[lhs151_] = rhs221_
                lhs152_.f1 = rhs222_
                d_1078_v78_ = rhs223_
            elif source15_.is_DC0:
                d_1091___mcc_h9_ = source15_.cf0
                d_1092_cf0_ = d_1091___mcc_h9_
                d_1093_v91_: _dafny.Array
                nw173_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_1093_v91_ = nw173_
                d_1094_v92_: _dafny.Array
                def lambda54_(d_1095_p2_):
                    def lambda55_(d_1096_i11_):
                        return d_1095_p2_

                    return lambda55_

                init29_ = lambda54_(p2)
                nw174_ = _dafny.Array(None, 1)
                for i0_29_ in range(nw174_.length(0)):
                    nw174_[i0_29_] = init29_(i0_29_)
                d_1094_v92_ = nw174_
                index196_ = default__.safeIndex(969, (d_1093_v91_).length(0))
                (d_1093_v91_)[index196_] = d_1094_v92_
                index197_ = default__.safeIndex(969, (d_1093_v91_).length(0))
                nw175_ = _dafny.Array(False, 5)
                (d_1093_v91_)[index197_] = nw175_
                d_1097_v93_: D11
                d_1097_v93_ = D11_DC19(d_1057_v67_)
                d_1098_v94_: _dafny.Map
                d_1098_v94_ = _dafny.Map({d_1097_v93_: default__.safeModuloInt(p1, p1)})
                d_1099_v95_: _dafny.MultiSet
                d_1099_v95_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrxmykw"))), p1])
                d_1098_v94_ = (d_1098_v94_).set(d_1097_v93_, ((d_1099_v95_)[(self).fm4(_dafny.SeqWithoutIsStrInference([d_1092_cf0_, p0]), p1, p1, (d_1099_v95_).set(p1, default__.abs(p1)), globalState)] if ((self).fm4(_dafny.SeqWithoutIsStrInference([d_1092_cf0_, p0]), p1, p1, (d_1099_v95_).set(p1, default__.abs(p1)), globalState)) in (d_1099_v95_) else p1))
                d_1100_v96_: _dafny.Array
                def lambda56_(d_1101_p1_):
                    def lambda57_(d_1102_i12_):
                        return default__.safeModuloInt(d_1102_i12_, d_1101_p1_)

                    return lambda57_

                init30_ = lambda56_(p1)
                nw176_ = _dafny.Array(None, 29)
                for i0_30_ in range(nw176_.length(0)):
                    nw176_[i0_30_] = init30_(i0_30_)
                d_1100_v96_ = nw176_
                index198_ = default__.safeIndex(888, (d_1100_v96_).length(0))
                (d_1100_v96_)[index198_] = p1
                index199_ = default__.safeIndex(888, (d_1100_v96_).length(0))
                rhs224_ = default__.fm3(p2, p0, globalState)
                rhs225_ = (p1) > (970)
                rhs226_ = d_1100_v96_
                lhs153_ = d_1100_v96_
                lhs154_ = default__.safeIndex(888, (d_1100_v96_).length(0))
                lhs153_[lhs154_] = rhs224_
                d_1092_cf0_ = rhs225_
                d_1100_v96_ = rhs226_
                (globalState).f1 = len(default__.fm40(globalState))
            elif True:
                d_1103___mcc_h10_ = source15_.cf3
                d_1104_cf3_ = d_1103___mcc_h10_
                d_1105_v97_: _dafny.Set
                d_1105_v97_ = _dafny.Set({(0) - (p1)})
                d_1106_v98_: _dafny.Array
                nw177_ = _dafny.Array(None, 16)
                nw177_[int(0)] = len(d_960_v0_)
                nw177_[int(1)] = p1
                nw177_[int(2)] = p1
                nw177_[int(3)] = len((d_1105_v97_) - (d_1105_v97_))
                nw177_[int(4)] = p1
                nw177_[int(5)] = (p1) * (p1)
                nw177_[int(6)] = p1
                nw177_[int(7)] = p1
                nw177_[int(8)] = p1
                nw177_[int(9)] = p1
                nw177_[int(10)] = p1
                nw177_[int(11)] = p1
                nw177_[int(12)] = (p1) - (-580)
                nw177_[int(13)] = p1
                nw177_[int(14)] = p1
                nw177_[int(15)] = p1
                d_1106_v98_ = nw177_
                index200_ = default__.safeIndex(796, (d_1106_v98_).length(0))
                (d_1106_v98_)[index200_] = (0) - ((0) - (((0) - (p1) if not(True) else p1)))
                index201_ = default__.safeIndex(796, (d_1106_v98_).length(0))
                (d_1106_v98_)[index201_] = (p1) - (default__.safeDivisionInt(p1, p1))
                d_1107_v99_: bool
                d_1107_v99_ = True
                d_1108_v100_: _dafny.Seq
                d_1108_v100_ = _dafny.SeqWithoutIsStrInference([d_1107_v99_, p2])
                d_1107_v99_ = (len(d_1108_v100_)) < (589)
                d_1109_v101_: C2
                nw178_ = C2()
                nw178_.ctor__((d_960_v0_)[default__.safeIndex((d_1106_v98_)[default__.safeIndex(796, (d_1106_v98_).length(0))], len(d_960_v0_))], not(default__.fm0(p0, globalState)))
                d_1109_v101_ = nw178_
                d_1110_v102_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.Map({}), 8)
                d_1110_v102_ = nw179_
                d_1111_v103_: _dafny.Array
                nw180_ = _dafny.Array(None, 8)
                nw180_[int(0)] = p0
                nw180_[int(1)] = p2
                nw180_[int(2)] = p2
                nw180_[int(3)] = p2
                nw180_[int(4)] = p2
                nw180_[int(5)] = False
                nw180_[int(6)] = not(p0)
                nw180_[int(7)] = d_1107_v99_
                d_1111_v103_ = nw180_
                d_1112_v104_: _dafny.Map
                d_1112_v104_ = _dafny.Map({p0: d_1111_v103_})
                index202_ = default__.safeIndex(850, (d_1110_v102_).length(0))
                (d_1110_v102_)[index202_] = (d_1112_v104_) | (d_1112_v104_)
                index203_ = default__.safeIndex(850, (d_1110_v102_).length(0))
                (d_1110_v102_)[index203_] = d_1112_v104_
            d_1113_v105_: C0
            nw181_ = C0()
            nw181_.ctor__(p2, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1])]))
            d_1113_v105_ = nw181_
            d_1114_v106_: _dafny.Array
            nw182_ = _dafny.Array(False, 5)
            d_1114_v106_ = nw182_
            d_1115_v107_: _dafny.Map
            d_1115_v107_ = _dafny.Map({p1: d_1114_v106_})
            d_1116_v108_: _dafny.Seq
            d_1116_v108_ = _dafny.SeqWithoutIsStrInference([d_1114_v106_])
            d_1117_v109_: _dafny.Array
            d_1117_v109_ = ((d_1115_v107_)[p1] if (p1) in (d_1115_v107_) else (d_1116_v108_)[default__.safeIndex(p1, len(d_1116_v108_))])
            d_1118_v110_: _dafny.Map
            d_1118_v110_ = _dafny.Map({d_1117_v109_: d_1114_v106_})
            d_1119_v111_: _dafny.Map
            d_1119_v111_ = d_1118_v110_
            d_1118_v110_ = (d_1119_v111_)
            d_1120_v112_: _dafny.Array
            def lambda58_(d_1121_p1_):
                def lambda59_(d_1122_i13_):
                    return (d_1122_i13_) - (len(_dafny.SeqWithoutIsStrInference([d_1121_p1_])))

                return lambda59_

            init31_ = lambda58_(p1)
            nw183_ = _dafny.Array(None, 20)
            for i0_31_ in range(nw183_.length(0)):
                nw183_[i0_31_] = init31_(i0_31_)
            d_1120_v112_ = nw183_
            d_1123_v113_: _dafny.Map
            d_1123_v113_ = _dafny.Map({p2: d_1120_v112_})
            d_1120_v112_ = ((d_1123_v113_)[not(p0)] if (not(p0)) in (d_1123_v113_) else d_1120_v112_)
        elif True:
            d_1124_v114_: _dafny.MultiSet
            d_1124_v114_ = _dafny.MultiSet([p2])
            (globalState).f4 = d_1124_v114_
            d_1125_v115_: _dafny.MultiSet
            d_1125_v115_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1126_i14_ in range(default__.abs(663))])])
            d_1127_v116_: _dafny.Map
            d_1127_v116_ = _dafny.Map({p0: d_1125_v115_})
            d_1127_v116_ = (d_1127_v116_).set((d_960_v0_) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbc"))), (d_1125_v115_) | (d_1125_v115_))
            d_1128_v117_: D0
            d_1128_v117_ = D0_DC0(p0)
            d_1129_v118_: _dafny.Seq
            d_1129_v118_ = _dafny.SeqWithoutIsStrInference([d_1128_v117_])
            d_1130_v119_: _dafny.Seq
            d_1130_v119_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1131_v120_: _dafny.Map
            d_1131_v120_ = _dafny.Map({d_1129_v118_: _dafny.MultiSet([d_1130_v119_])})
            d_1132_v121_: _dafny.MultiSet
            d_1132_v121_ = _dafny.MultiSet([default__.fm2(p2, p0, globalState)])
            d_1131_v120_ = (d_1131_v120_).set(d_1129_v118_, (_dafny.MultiSet([d_1130_v119_, d_1130_v119_])).intersection(d_1132_v121_))
            d_1133_v122_: _dafny.Array
            def lambda60_(d_1134_p2_):
                def lambda61_(d_1135_i15_):
                    return not (d_1134_p2_) or (d_1134_p2_)

                return lambda61_

            init32_ = lambda60_(p2)
            nw184_ = _dafny.Array(None, 26)
            for i0_32_ in range(nw184_.length(0)):
                nw184_[i0_32_] = init32_(i0_32_)
            d_1133_v122_ = nw184_
            index204_ = default__.safeIndex(106, (d_1133_v122_).length(0))
            (d_1133_v122_)[index204_] = (p2 if p0 else p2)
            index205_ = default__.safeIndex(106, (d_1133_v122_).length(0))
            (d_1133_v122_)[index205_] = True
            d_1136_v123_: _dafny.Set
            d_1136_v123_ = _dafny.Set({(d_960_v0_) + (d_960_v0_)})
            d_1137_v124_: str
            d_1137_v124_ = _dafny.CodePoint('n')
            index206_ = default__.safeIndex(106, (d_1133_v122_).length(0))
            rhs227_ = d_1136_v123_
            rhs228_ = default__.fm0(((0) - ((0) - (p1))) >= (65), globalState)
            rhs229_ = d_1137_v124_
            lhs155_ = d_1133_v122_
            lhs156_ = default__.safeIndex(106, (d_1133_v122_).length(0))
            d_1136_v123_ = rhs227_
            lhs155_[lhs156_] = rhs228_
            d_1137_v124_ = rhs229_
        d_1138_v125_: _dafny.Array
        def lambda62_(d_1139_p1_):
            def lambda63_(d_1140_i16_):
                return (d_1140_i16_) - ((0) - (d_1139_p1_))

            return lambda63_

        init33_ = lambda62_(p1)
        nw185_ = _dafny.Array(None, 7)
        for i0_33_ in range(nw185_.length(0)):
            nw185_[i0_33_] = init33_(i0_33_)
        d_1138_v125_ = nw185_
        d_1141_v126_: _dafny.Map
        d_1141_v126_ = _dafny.Map({p1: p2})
        d_1142_v127_: _dafny.Map
        d_1142_v127_ = _dafny.Map({default__.fm0(p2, globalState): len(d_1141_v126_)})
        index207_ = default__.safeIndex(129, (d_1138_v125_).length(0))
        (d_1138_v125_)[index207_] = (len(d_1142_v127_)) - (p1)
        d_1143_v128_: _dafny.Map
        d_1143_v128_ = _dafny.Map({p1: d_960_v0_})
        d_1144_v129_: _dafny.Seq
        d_1144_v129_ = _dafny.SeqWithoutIsStrInference([p2])
        d_1145_v130_: _dafny.Map
        d_1145_v130_ = _dafny.Map({d_1144_v129_: 396})
        index208_ = default__.safeIndex(129, (d_1138_v125_).length(0))
        (d_1138_v125_)[index208_] = len(((d_960_v0_) + (((d_1143_v128_)[len(d_1145_v130_)] if (len(d_1145_v130_)) in (d_1143_v128_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1146_i17_ in range(default__.abs(367))])))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvaeqnmre"))) + (d_960_v0_)))
        d_1147_v131_: _dafny.Seq
        d_1147_v131_ = d_1144_v129_
        source16_ = d_1147_v131_
        d_1148___mcc_h11_ = source16_
        d_1149_cf16_ = d_1148___mcc_h11_
        d_1150_v132_: bool
        d_1150_v132_ = False
        d_1150_v132_ = p2
        d_1151_v133_: D11
        d_1151_v133_ = D11_DC20(p2, 730)
        index209_ = default__.safeIndex(129, (d_1138_v125_).length(0))
        (d_1138_v125_)[index209_] = default__.safeModuloInt((d_1151_v133_).cf22, (d_1138_v125_)[default__.safeIndex(129, (d_1138_v125_).length(0))])
        index210_ = default__.safeIndex(129, (d_1138_v125_).length(0))
        (d_1138_v125_)[index210_] = (d_1138_v125_)[default__.safeIndex(129, (d_1138_v125_).length(0))]
        d_1152_v134_: _dafny.Seq
        d_1152_v134_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1153_v135_: _dafny.Seq
        d_1153_v135_ = _dafny.SeqWithoutIsStrInference([d_1152_v134_, d_1152_v134_, d_1152_v134_])
        d_1154_v136_: C0
        nw186_ = C0()
        nw186_.ctor__(p2, d_1153_v135_)
        d_1154_v136_ = nw186_
        d_1155_v137_: _dafny.Seq
        d_1155_v137_ = _dafny.SeqWithoutIsStrInference([p1, 709, (d_1138_v125_)[default__.safeIndex(129, (d_1138_v125_).length(0))]])
        d_1156_v138_: _dafny.Seq
        d_1156_v138_ = _dafny.SeqWithoutIsStrInference([d_1155_v137_, d_1155_v137_])
        d_1157_v139_: C0
        nw187_ = C0()
        nw187_.ctor__((len(d_1155_v137_)) < ((d_1138_v125_)[default__.safeIndex(129, (d_1138_v125_).length(0))]), (d_1156_v138_) + (d_1156_v138_))
        d_1157_v139_ = nw187_


class C5(T3, T2, T1, T0):
    def  __init__(self):
        self._f5: str = _dafny.CodePoint('D')
        self._f6: bool = False
        self.f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f15, f5, f6):
        (self).f15 = f15
        (self)._f5 = f5
        (self)._f6 = f6

    def fm8(self, p0, p1, p2, globalState):
        return self.f15

    def fm6(self, p0, globalState):
        return (self).f6

    def fm7(self, p0, globalState):
        return len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbq")) if (self).f6 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pt"))))

    def fm4(self, p0, p1, p2, p3, globalState):
        return self.f15

    def fm5(self, p0, globalState):
        return D1_DC4((0) - (self.f15))

    def fm42(self, p0, p1, p2, p3, globalState):
        def iife67_():
            coll35_ = _dafny.Set()
            compr_35_: _dafny.Map
            for compr_35_ in (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1158_i0_ in range(default__.abs(385))])): self.f15}): (self).f5})).keys.Elements:
                d_1159_v0_: _dafny.Map = compr_35_
                if (d_1159_v0_) in (_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1158_i0_ in range(default__.abs(385))])): self.f15}): (self).f5})):
                    coll35_ = coll35_.union(_dafny.Set([d_1159_v0_]))
            return _dafny.Set(coll35_)
        def iife68_():
            coll36_ = _dafny.Map()
            compr_36_: int
            for compr_36_ in (_dafny.SeqWithoutIsStrInference([(D1_DC4(self.f15)).cf4])).Elements:
                d_1160_v1_: int = compr_36_
                if (d_1160_v1_) in (_dafny.SeqWithoutIsStrInference([(D1_DC4(self.f15)).cf4])):
                    coll36_[default__.safeDivisionInt(d_1160_v1_, self.f15)] = _dafny.Set({_dafny.Map({self.f15: self.f15}), _dafny.Map({self.f15: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")))})})
            return _dafny.Map(coll36_)
        def iife69_():
            coll37_ = _dafny.Map()
            compr_37_: int
            for compr_37_ in _dafny.IntegerRange(844, 86):
                d_1161_v2_: int = compr_37_
                if ((844) <= (d_1161_v2_)) and ((d_1161_v2_) < (86)):
                    def iife70_():
                        coll38_ = _dafny.Set()
                        compr_38_: _dafny.Map
                        for compr_38_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({415: self.f15}), _dafny.Map({len(_dafny.Map({(self).f6: self.f15})): self.f15})])).Elements:
                            d_1162_v3_: _dafny.Map = compr_38_
                            if (d_1162_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({415: self.f15}), _dafny.Map({len(_dafny.Map({(self).f6: self.f15})): self.f15})])):
                                coll38_ = coll38_.union(_dafny.Set([d_1162_v3_]))
                        return _dafny.Set(coll38_)
                    coll37_[default__.safeModuloInt(d_1161_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpe"))))] = iife70_()

            return _dafny.Map(coll37_)
        def iife71_():
            coll39_ = _dafny.Set()
            compr_39_: _dafny.Map
            for compr_39_ in (_dafny.Map({_dafny.Map({self.f15: self.f15}): self.f15})).keys.Elements:
                d_1163_v4_: _dafny.Map = compr_39_
                if (d_1163_v4_) in (_dafny.Map({_dafny.Map({self.f15: self.f15}): self.f15})):
                    coll39_ = coll39_.union(_dafny.Set([d_1163_v4_]))
            return _dafny.Set(coll39_)
        return ((_dafny.Map({self.f15: iife67_()
        })) | (iife68_()
        )) | ((iife69_()
        ) | (_dafny.Map({self.f15: iife71_()
        })))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1164_v0_: _dafny.Seq
        d_1164_v0_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({self.f15: True}))])
        hi6_ = p1
        for d_1165_i0_ in range((535) * (len(d_1164_v0_)), hi6_):
            r1 = (p2 if not((d_1165_i0_) <= (len(p2))) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpeabsjw")))
            d_1166_v1_: C1
            nw188_ = C1()
            nw188_.ctor__()
            d_1166_v1_ = nw188_
            d_1167_v2_: _dafny.Array
            def lambda64_(d_1168_i1_):
                return (self).f6

            init34_ = lambda64_
            nw189_ = _dafny.Array(None, 8)
            for i0_34_ in range(nw189_.length(0)):
                nw189_[i0_34_] = init34_(i0_34_)
            d_1167_v2_ = nw189_
            d_1169_v3_: _dafny.Map
            d_1169_v3_ = _dafny.Map({d_1166_v1_: d_1167_v2_})
            source17_ = d_1169_v3_
            d_1170___mcc_h0_ = source17_
            d_1171_cf18_ = d_1170___mcc_h0_
            r1 = (p2) + (p2)
            d_1172_v4_: _dafny.MultiSet
            d_1172_v4_ = _dafny.MultiSet([False])
            d_1173_v5_: _dafny.Seq
            d_1173_v5_ = _dafny.SeqWithoutIsStrInference([default__.fm0((self).f6, globalState)])
            (globalState).f1 = ((d_1172_v4_).cardinality) * (len(d_1173_v5_))
            r0 = p0
            d_1174_v6_: _dafny.MultiSet
            d_1174_v6_ = _dafny.MultiSet([self.f15])
            d_1175_v7_: _dafny.Map
            d_1175_v7_ = _dafny.Map({p0: d_1165_i0_})
            d_1176_v8_: _dafny.Map
            d_1176_v8_ = _dafny.Map({(0) - (-800): p1})
            d_1177_v9_: _dafny.Map
            d_1177_v9_ = _dafny.Map({d_1176_v8_: (self).f6})
            d_1178_v10_: _dafny.Map
            d_1178_v10_ = _dafny.Map({(d_1172_v4_).ispropersubset(d_1172_v4_): (d_1174_v6_).issubset(((d_1174_v6_).set(-251, default__.abs(len(d_1175_v7_)))).set(p1, default__.abs(len(d_1177_v9_))))})
            d_1178_v10_ = (d_1178_v10_).set(p0, p0)
            d_1179_v11_: D6
            d_1179_v11_ = D6_DC14((332) > (d_1165_i0_))
            d_1180_v12_: _dafny.Array
            nw190_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 5)
            d_1180_v12_ = nw190_
            index211_ = default__.safeIndex(153, (d_1180_v12_).length(0))
            (d_1180_v12_)[index211_] = p2
            d_1181_v13_: _dafny.Array
            nw191_ = _dafny.Array(_dafny.Set({}), 24)
            d_1181_v13_ = nw191_
            index212_ = default__.safeIndex(776, (d_1181_v13_).length(0))
            (d_1181_v13_)[index212_] = default__.fm43(not(p0), not((self).f6), globalState)
            d_1182_v14_: _dafny.Set
            d_1182_v14_ = _dafny.Set({p1, self.f15, 439, len(d_1164_v0_), p1})
            d_1183_v15_: _dafny.Seq
            d_1183_v15_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({self.f15})) - (d_1182_v14_)])
            d_1184_v16_: _dafny.MultiSet
            d_1184_v16_ = _dafny.MultiSet([(self.f15) * (self.f15)])
            d_1185_v17_: _dafny.Map
            d_1185_v17_ = _dafny.Map({(d_1166_v1_).fm7((self).f6, globalState): not((self).f6)})
            index213_ = default__.safeIndex(153, (d_1180_v12_).length(0))
            index214_ = default__.safeIndex(776, (d_1181_v13_).length(0))
            def iife72_():
                coll40_ = _dafny.Set()
                compr_40_: int
                for compr_40_ in (d_1185_v17_).keys.Elements:
                    d_1186_v18_: int = compr_40_
                    if (d_1186_v18_) in (d_1185_v17_):
                        coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_1186_v18_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1187_i2_ in range(default__.abs(905))])))]))
                return _dafny.Set(coll40_)
            rhs230_ = d_1179_v11_
            rhs231_ = ((d_1184_v16_)[(p1) * (self.f15)] if ((p1) * (self.f15)) in (d_1184_v16_) else d_1165_i0_)
            rhs232_ = (p2 if (d_1165_i0_) != (d_1165_i0_) else p2)
            rhs233_ = (d_1182_v14_) | ((d_1182_v14_) | (iife72_()
            ))
            rhs234_ = d_1183_v15_
            lhs157_ = globalState
            lhs158_ = d_1180_v12_
            lhs159_ = default__.safeIndex(153, (d_1180_v12_).length(0))
            lhs160_ = d_1181_v13_
            lhs161_ = default__.safeIndex(776, (d_1181_v13_).length(0))
            d_1179_v11_ = rhs230_
            lhs157_.f1 = rhs231_
            lhs158_[lhs159_] = rhs232_
            lhs160_[lhs161_] = rhs233_
            d_1183_v15_ = rhs234_
            index215_ = default__.safeIndex(480, (d_1167_v2_).length(0))
            (d_1167_v2_)[index215_] = p0
            index216_ = default__.safeIndex(480, (d_1167_v2_).length(0))
            (d_1167_v2_)[index216_] = ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_1188_i3_ in range(default__.abs(-713))])) != ((d_1180_v12_)[default__.safeIndex(153, (d_1180_v12_).length(0))]) if ((self).f6) or ((self).f6) else not(not (p0) or (False)))
        (self).f15 = p1
        d_1189_v19_: C3
        nw192_ = C3()
        nw192_.ctor__(p0, (self).fm6(d_1164_v0_, globalState))
        d_1189_v19_ = nw192_
        (self).f15 = (d_1164_v0_)[default__.safeIndex(self.f15, len(d_1164_v0_))]
        d_1190_v20_: _dafny.MultiSet
        d_1190_v20_ = _dafny.MultiSet([(d_1189_v19_).f13, False, p0, (d_1189_v19_).f14, (d_1189_v19_).f14])
        r0 = (_dafny.MultiSet([(d_1189_v19_).f13])).ispropersubset((d_1190_v20_ if (self).f6 else d_1190_v20_))
        d_1191_v21_: D0
        d_1191_v21_ = D0_DC1()
        d_1192_v22_: D0
        d_1192_v22_ = D0_DC3(d_1191_v21_)
        d_1193_v23_: _dafny.Map
        d_1193_v23_ = _dafny.Map({d_1192_v22_: (0) - (p1)})
        d_1194_v24_: _dafny.Seq
        d_1194_v24_ = _dafny.SeqWithoutIsStrInference([d_1193_v23_, d_1193_v23_, d_1193_v23_, d_1193_v23_])
        d_1195_v25_: _dafny.Array
        nw193_ = _dafny.Array(None, 3)
        nw193_[int(0)] = d_1194_v24_
        nw193_[int(1)] = d_1194_v24_
        nw193_[int(2)] = d_1194_v24_
        d_1195_v25_ = nw193_
        index217_ = default__.safeIndex(46, (d_1195_v25_).length(0))
        (d_1195_v25_)[index217_] = _dafny.SeqWithoutIsStrInference([d_1193_v23_])
        index218_ = default__.safeIndex(46, (d_1195_v25_).length(0))
        (d_1195_v25_)[index218_] = (d_1194_v24_) + (d_1194_v24_)
        r0 = (self).f6
        r1 = ((p2) + (((p2).set(default__.safeIndex(self.f15, len(p2)), (self).f5)).set(default__.safeIndex(p1, len((p2).set(default__.safeIndex(self.f15, len(p2)), (self).f5))), (self).f5))) + ((p2).set(default__.safeIndex(p1, len(p2)), (self).f5))
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        source18_ = D0_DC1()
        if source18_.is_DC1:
            d_1196_v0_: bool
            d_1196_v0_ = False
            d_1197_v1_: _dafny.Seq
            d_1197_v1_ = _dafny.SeqWithoutIsStrInference([(self).f6, False])
            d_1196_v0_ = (d_1197_v1_) <= (d_1197_v1_)
            d_1198_v2_: _dafny.Map
            d_1198_v2_ = _dafny.Map({True: (self.f15) < (455)})
            d_1198_v2_ = (d_1198_v2_).set((self).f6, (self).f6)
            d_1199_v3_: _dafny.Array
            nw194_ = _dafny.Array(False, 15)
            d_1199_v3_ = nw194_
            index219_ = default__.safeIndex(725, (d_1199_v3_).length(0))
            (d_1199_v3_)[index219_] = (d_1196_v0_) or (not(False))
            d_1200_v4_: _dafny.MultiSet
            d_1200_v4_ = _dafny.MultiSet([d_1196_v0_, d_1196_v0_])
            d_1201_v5_: D15
            d_1201_v5_ = D15_DC27(d_1200_v4_)
            index220_ = default__.safeIndex(725, (d_1199_v3_).length(0))
            rhs235_ = ((d_1200_v4_) | (d_1200_v4_)).intersection((d_1201_v5_).cf35)
            rhs236_ = ((0) - (self.f15)) != (self.f15)
            lhs162_ = globalState
            lhs163_ = d_1199_v3_
            lhs164_ = default__.safeIndex(725, (d_1199_v3_).length(0))
            lhs162_.f4 = rhs235_
            lhs163_[lhs164_] = rhs236_
            d_1202_v6_: _dafny.MultiSet
            d_1202_v6_ = _dafny.MultiSet([self.f15, self.f15])
            d_1203_v7_: _dafny.Map
            d_1203_v7_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wl"))): d_1202_v6_})
            d_1204_v8_: _dafny.Seq
            d_1204_v8_ = _dafny.SeqWithoutIsStrInference([758])
            d_1205_v9_: _dafny.Seq
            d_1205_v9_ = _dafny.SeqWithoutIsStrInference([d_1202_v6_])
            d_1206_v10_: _dafny.Array
            nw195_ = _dafny.Array(None, 14)
            nw195_[int(0)] = _dafny.MultiSet([(default__.fm44(not((d_1199_v3_)[default__.safeIndex(725, (d_1199_v3_).length(0))]), globalState)).cardinality])
            nw195_[int(1)] = default__.fm45(globalState)
            nw195_[int(2)] = ((d_1203_v7_)[len(default__.fm46((self).fm6(d_1204_v8_, globalState), globalState))] if (len(default__.fm46((self).fm6(d_1204_v8_, globalState), globalState))) in (d_1203_v7_) else _dafny.MultiSet([self.f15]))
            nw195_[int(3)] = (d_1202_v6_) | ((d_1202_v6_).set(self.f15, default__.abs(self.f15)))
            nw195_[int(4)] = d_1202_v6_
            nw195_[int(5)] = d_1202_v6_
            nw195_[int(6)] = d_1202_v6_
            nw195_[int(7)] = _dafny.MultiSet([len(d_1204_v8_)])
            nw195_[int(8)] = d_1202_v6_
            nw195_[int(9)] = d_1202_v6_
            nw195_[int(10)] = d_1202_v6_
            nw195_[int(11)] = d_1202_v6_
            nw195_[int(12)] = (d_1205_v9_)[default__.safeIndex(self.f15, len(d_1205_v9_))]
            nw195_[int(13)] = ((d_1202_v6_).set(self.f15, default__.abs(self.f15))).intersection(d_1202_v6_)
            d_1206_v10_ = nw195_
            index221_ = default__.safeIndex(9, (d_1206_v10_).length(0))
            (d_1206_v10_)[index221_] = d_1202_v6_
            index222_ = default__.safeIndex(9, (d_1206_v10_).length(0))
            (d_1206_v10_)[index222_] = d_1202_v6_
        elif source18_.is_DC2:
            d_1207___mcc_h0_ = source18_.cf1
            d_1208___mcc_h1_ = source18_.cf2
            d_1209_cf2_ = d_1208___mcc_h1_
            d_1210_cf1_ = d_1207___mcc_h0_
            d_1211_v11_: bool
            d_1211_v11_ = False
            d_1211_v11_ = (self).f6
            r0 = (self.f15 if ((self).f6 if (self).f6 else False) else self.f15)
            d_1212_v12_: _dafny.MultiSet
            d_1212_v12_ = _dafny.MultiSet([d_1211_v11_])
            source19_ = D15_DC27(d_1212_v12_)
            if source19_.is_DC28:
                d_1213___mcc_h4_ = source19_.cf36
                d_1214___mcc_h5_ = source19_.cf37
                d_1215_cf37_ = d_1214___mcc_h5_
                d_1216_cf36_ = d_1213___mcc_h4_
                (globalState).f1 = (self).fm7((self).f6, globalState)
                d_1217_v13_: _dafny.Array
                nw196_ = _dafny.Array(False, 7)
                d_1217_v13_ = nw196_
                d_1218_v14_: _dafny.Seq
                d_1218_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
                d_1219_v15_: _dafny.MultiSet
                d_1219_v15_ = _dafny.MultiSet([d_1218_v14_, d_1218_v14_, _dafny.SeqWithoutIsStrInference([d_1209_cf2_ for d_1220_i0_ in range(default__.abs(969))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjrtxt")), d_1218_v14_])
                d_1221_v17_: _dafny.Map
                def iife73_():
                    coll41_ = _dafny.Map()
                    compr_41_: int
                    for compr_41_ in _dafny.IntegerRange(717, -196):
                        d_1222_v16_: int = compr_41_
                        if ((717) <= (d_1222_v16_)) and ((d_1222_v16_) < (-196)):
                            coll41_[default__.safeDivisionInt(d_1222_v16_, (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, (self).f6])])).cardinality)] = 868
                    return _dafny.Map(coll41_)
                d_1221_v17_ = iife73_()
                
                d_1223_v18_: D12
                d_1223_v18_ = D12_DC22(d_1221_v17_, _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1211_v11_, d_1215_cf37_]) for d_1224_i1_ in range(default__.abs(549))]), d_1209_cf2_, d_1211_v11_, d_1218_v14_)
                nw197_ = _dafny.Array(None, 19)
                nw197_[int(0)] = (self).f6
                nw197_[int(1)] = (d_1219_v15_).issubset(d_1219_v15_)
                nw197_[int(2)] = d_1215_cf37_
                nw197_[int(3)] = (not(d_1211_v11_) if d_1211_v11_ else d_1216_cf36_)
                nw197_[int(4)] = d_1215_cf37_
                nw197_[int(5)] = False
                nw197_[int(6)] = d_1216_cf36_
                nw197_[int(7)] = d_1216_cf36_
                nw197_[int(8)] = d_1215_cf37_
                nw197_[int(9)] = d_1211_v11_
                nw197_[int(10)] = (d_1218_v14_) < ((d_1218_v14_).set(default__.safeIndex(self.f15, len(d_1218_v14_)), (d_1223_v18_).cf26))
                nw197_[int(11)] = (440) > (self.f15)
                nw197_[int(12)] = d_1211_v11_
                nw197_[int(13)] = d_1211_v11_
                nw197_[int(14)] = d_1215_cf37_
                nw197_[int(15)] = d_1216_cf36_
                nw197_[int(16)] = d_1216_cf36_
                nw197_[int(17)] = not(not(d_1215_cf37_))
                nw197_[int(18)] = d_1211_v11_
                d_1217_v13_ = nw197_
                (globalState).f1 = (self.f15) * (self.f15)
                d_1225_v19_: D11
                d_1225_v19_ = D11_DC20((d_1215_cf37_ if (self).f6 else d_1211_v11_), (self.f15) - (self.f15))
                d_1225_v19_ = d_1225_v19_
            elif source19_.is_DC27:
                d_1226___mcc_h6_ = source19_.cf35
                d_1227_cf35_ = d_1226___mcc_h6_
                d_1228_v20_: _dafny.Map
                d_1228_v20_ = _dafny.Map({(self).f6: (self).f6})
                d_1229_v21_: _dafny.Map
                d_1229_v21_ = _dafny.Map({self.f15: d_1228_v20_})
                d_1230_v22_: _dafny.Map
                d_1230_v22_ = _dafny.Map({d_1211_v11_: default__.fm47((self).f6, globalState)})
                (globalState).f1 = len((d_1229_v21_) | (((d_1230_v22_)[(self).f6] if ((self).f6) in (d_1230_v22_) else d_1229_v21_)))
                (self).f15 = self.f15
                d_1231_v23_: _dafny.Array
                nw198_ = _dafny.Array(int(0), 17)
                d_1231_v23_ = nw198_
                d_1232_v24_: _dafny.Array
                d_1232_v24_ = d_1231_v23_
                d_1233_v25_: _dafny.Array
                nw199_ = _dafny.Array(None, 28)
                nw199_[int(0)] = d_1231_v23_
                nw199_[int(1)] = (d_1232_v24_)
                nw199_[int(2)] = (d_1231_v23_ if d_1211_v11_ else d_1231_v23_)
                nw199_[int(3)] = d_1231_v23_
                nw199_[int(4)] = d_1231_v23_
                nw199_[int(5)] = (d_1232_v24_)
                nw199_[int(6)] = d_1231_v23_
                nw199_[int(7)] = d_1231_v23_
                nw199_[int(8)] = d_1231_v23_
                nw199_[int(9)] = d_1231_v23_
                nw199_[int(10)] = d_1231_v23_
                nw199_[int(11)] = d_1231_v23_
                nw199_[int(12)] = d_1231_v23_
                nw199_[int(13)] = d_1231_v23_
                nw199_[int(14)] = d_1231_v23_
                nw199_[int(15)] = d_1231_v23_
                nw199_[int(16)] = d_1231_v23_
                nw199_[int(17)] = d_1231_v23_
                nw199_[int(18)] = d_1231_v23_
                nw199_[int(19)] = d_1231_v23_
                nw199_[int(20)] = d_1231_v23_
                nw199_[int(21)] = d_1231_v23_
                nw199_[int(22)] = d_1231_v23_
                nw199_[int(23)] = d_1231_v23_
                nw199_[int(24)] = d_1231_v23_
                nw199_[int(25)] = d_1231_v23_
                nw199_[int(26)] = d_1231_v23_
                nw199_[int(27)] = d_1231_v23_
                d_1233_v25_ = nw199_
                index223_ = default__.safeIndex(514, (d_1233_v25_).length(0))
                (d_1233_v25_)[index223_] = (d_1231_v23_ if d_1211_v11_ else d_1231_v23_)
                index224_ = default__.safeIndex(514, (d_1233_v25_).length(0))
                (d_1233_v25_)[index224_] = d_1231_v23_
                (globalState).f1 = (self.f15 if (self).f6 else self.f15)
            elif True:
                d_1234___mcc_h7_ = source19_.cf38
                d_1235_cf38_ = d_1234___mcc_h7_
                d_1211_v11_ = (self).f6
                d_1236_v26_: _dafny.Map
                d_1236_v26_ = _dafny.Map({(self).f6: (d_1212_v12_).cardinality})
                d_1237_v27_: _dafny.Array
                def lambda65_(d_1238_i2_):
                    return (d_1238_i2_) + (self.f15)

                init35_ = lambda65_
                nw200_ = _dafny.Array(None, 15)
                for i0_35_ in range(nw200_.length(0)):
                    nw200_[i0_35_] = init35_(i0_35_)
                d_1237_v27_ = nw200_
                d_1239_v28_: _dafny.Map
                d_1239_v28_ = _dafny.Map({(_dafny.Map({d_1211_v11_: (0) - (self.f15)})) | (d_1236_v26_): d_1237_v27_})
                d_1240_v29_: _dafny.Set
                d_1240_v29_ = _dafny.Set({d_1236_v26_, d_1236_v26_, (_dafny.Map({(self).f6: self.f15})) | (d_1236_v26_), d_1236_v26_, (d_1236_v26_) | (d_1236_v26_)})
                index225_ = default__.safeIndex(494, (d_1237_v27_).length(0))
                (d_1237_v27_)[index225_] = self.f15
                index226_ = default__.safeIndex(494, (d_1237_v27_).length(0))
                rhs237_ = self.f15
                rhs238_ = (_dafny.Map({d_1236_v26_: d_1237_v27_})) | (d_1239_v28_)
                rhs239_ = d_1240_v29_
                rhs240_ = self.f15
                rhs241_ = 459
                lhs165_ = globalState
                lhs166_ = d_1237_v27_
                lhs167_ = default__.safeIndex(494, (d_1237_v27_).length(0))
                lhs168_ = globalState
                lhs165_.f1 = rhs237_
                d_1239_v28_ = rhs238_
                d_1240_v29_ = rhs239_
                lhs166_[lhs167_] = rhs240_
                lhs168_.f1 = rhs241_
                d_1241_v30_: _dafny.Array
                def lambda66_(d_1242_i3_):
                    return (self).f6

                init36_ = lambda66_
                nw201_ = _dafny.Array(None, 3)
                for i0_36_ in range(nw201_.length(0)):
                    nw201_[i0_36_] = init36_(i0_36_)
                d_1241_v30_ = nw201_
                d_1243_v31_: D0
                d_1244_v32_: bool
                out34_: D0
                out35_: bool
                out34_, out35_ = default__.m0(d_1241_v30_, (self).f5, not((self).f6), d_1209_cf2_, globalState)
                d_1243_v31_ = out34_
                d_1244_v32_ = out35_
                d_1245_v33_: _dafny.Map
                d_1245_v33_ = _dafny.Map({(0) - ((default__.fm48((d_1237_v27_)[default__.safeIndex(494, (d_1237_v27_).length(0))], True, globalState)).cardinality): d_1244_v32_})
                d_1246_v34_: _dafny.Map
                d_1246_v34_ = d_1245_v33_
                d_1247_v35_: _dafny.Map
                d_1247_v35_ = _dafny.Map({d_1244_v32_: d_1246_v34_})
                d_1247_v35_ = (d_1247_v35_).set(((self).f6 if d_1244_v32_ else d_1211_v11_), d_1246_v34_)
            d_1248_v36_: _dafny.Seq
            d_1248_v36_ = _dafny.SeqWithoutIsStrInference([self.f15])
            d_1249_v37_: _dafny.Map
            d_1249_v37_ = _dafny.Map({D14_DC26(d_1211_v11_, len(d_1248_v36_), d_1209_cf2_): (self).f5})
            def iife74_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in (d_1248_v36_).Elements:
                    d_1250_v38_: int = compr_42_
                    if (d_1250_v38_) in (d_1248_v36_):
                        def iife75_():
                            coll43_ = _dafny.Map()
                            compr_43_: D14
                            for compr_43_ in (_dafny.SeqWithoutIsStrInference([D14_DC26(d_1211_v11_, self.f15, d_1209_cf2_) for d_1251_i4_ in range(default__.abs(-187))])).Elements:
                                d_1252_v39_: D14 = compr_43_
                                if (d_1252_v39_) in (_dafny.SeqWithoutIsStrInference([D14_DC26(d_1211_v11_, self.f15, d_1209_cf2_) for d_1251_i4_ in range(default__.abs(-187))])):
                                    coll43_[d_1252_v39_] = _dafny.CodePoint('q')
                            return _dafny.Map(coll43_)
                        coll42_[(d_1250_v38_) + (self.f15)] = (iife75_()
                        ).set(D14_DC26(d_1211_v11_, self.f15, (self).f5), d_1209_cf2_)
                return _dafny.Map(coll42_)
            d_1211_v11_ = (_dafny.Map({self.f15: d_1249_v37_})) != (iife74_()
            )
        elif source18_.is_DC0:
            d_1253___mcc_h2_ = source18_.cf0
            d_1254_cf0_ = d_1253___mcc_h2_
            d_1255_v40_: _dafny.Array
            def lambda67_(d_1256_i5_):
                return (d_1256_i5_) - (len(_dafny.Map({not((self).f6): (D6_DC14((self).f6)).cf15})))

            init37_ = lambda67_
            nw202_ = _dafny.Array(None, 27)
            for i0_37_ in range(nw202_.length(0)):
                nw202_[i0_37_] = init37_(i0_37_)
            d_1255_v40_ = nw202_
            d_1257_v41_: _dafny.Seq
            d_1257_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epvurt"))
            d_1258_v42_: _dafny.Seq
            d_1258_v42_ = _dafny.SeqWithoutIsStrInference([len(d_1257_v41_)])
            d_1259_v43_: _dafny.Seq
            d_1259_v43_ = _dafny.SeqWithoutIsStrInference([len(d_1258_v42_)])
            d_1260_v44_: _dafny.Array
            nw203_ = _dafny.Array(None, 3)
            nw203_[int(0)] = d_1254_cf0_
            nw203_[int(1)] = (self).fm6(d_1259_v43_, globalState)
            nw203_[int(2)] = d_1254_cf0_
            d_1260_v44_ = nw203_
            d_1261_v45_: _dafny.Set
            d_1261_v45_ = _dafny.Set({797, 526, len(default__.fm49(globalState))})
            index227_ = default__.safeIndex(764, (d_1260_v44_).length(0))
            (d_1260_v44_)[index227_] = (len(d_1261_v45_)) in (d_1261_v45_)
            d_1262_v46_: _dafny.Seq
            d_1262_v46_ = _dafny.SeqWithoutIsStrInference([d_1254_cf0_, (self).f6])
            d_1263_v47_: _dafny.Seq
            d_1263_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(d_1262_v46_), 53, self.f15, self.f15])])
            d_1264_v48_: D6
            d_1264_v48_ = D6_DC13(d_1258_v42_)
            d_1265_v49_: _dafny.Map
            d_1265_v49_ = _dafny.Map({len(d_1261_v45_): d_1254_cf0_})
            d_1266_v50_: _dafny.Array
            nw204_ = _dafny.Array(None, 22)
            nw204_[int(0)] = ((d_1258_v42_).set(default__.safeIndex(-560, len(d_1258_v42_)), 501)) + (d_1258_v42_)
            nw204_[int(1)] = d_1258_v42_
            nw204_[int(2)] = d_1259_v43_
            nw204_[int(3)] = (d_1263_v47_)[default__.safeIndex(self.f15, len(d_1263_v47_))]
            nw204_[int(4)] = d_1259_v43_
            nw204_[int(5)] = _dafny.SeqWithoutIsStrInference([189 for d_1267_i6_ in range(default__.abs(370))])
            nw204_[int(6)] = _dafny.SeqWithoutIsStrInference([self.f15, len(d_1257_v41_), self.f15])
            nw204_[int(7)] = d_1258_v42_
            nw204_[int(8)] = (d_1259_v43_).set(default__.safeIndex(700, len(d_1259_v43_)), self.f15)
            nw204_[int(9)] = (_dafny.SeqWithoutIsStrInference([self.f15, self.f15])) + ((d_1264_v48_).cf14)
            nw204_[int(10)] = d_1258_v42_
            nw204_[int(11)] = d_1259_v43_
            nw204_[int(12)] = (_dafny.SeqWithoutIsStrInference([len(d_1265_v49_)])) + (_dafny.SeqWithoutIsStrInference([-85 for d_1268_i7_ in range(default__.abs(-2))]))
            nw204_[int(13)] = d_1259_v43_
            nw204_[int(14)] = (d_1258_v42_) + (d_1258_v42_)
            nw204_[int(15)] = (d_1258_v42_).set(default__.safeIndex((0) - (self.f15), len(d_1258_v42_)), len(_dafny.Map({d_1254_cf0_: self.f15})))
            nw204_[int(16)] = ((_dafny.SeqWithoutIsStrInference([self.f15 for d_1269_i8_ in range(default__.abs(794))])) + (d_1258_v42_)).set(default__.safeIndex((0) - (self.f15), len((_dafny.SeqWithoutIsStrInference([self.f15 for d_1270_i8_ in range(default__.abs(794))])) + (d_1258_v42_))), self.f15)
            nw204_[int(17)] = d_1259_v43_
            nw204_[int(18)] = d_1258_v42_
            nw204_[int(19)] = (d_1259_v43_) + (d_1259_v43_)
            nw204_[int(20)] = d_1258_v42_
            nw204_[int(21)] = _dafny.SeqWithoutIsStrInference([self.f15 for d_1271_i9_ in range(default__.abs(781))])
            d_1266_v50_ = nw204_
            index228_ = default__.safeIndex(637, (d_1266_v50_).length(0))
            (d_1266_v50_)[index228_] = _dafny.SeqWithoutIsStrInference([self.f15, self.f15])
            d_1272_v51_: _dafny.MultiSet
            d_1272_v51_ = _dafny.MultiSet([(self).f6])
            d_1273_v52_: _dafny.Map
            d_1273_v52_ = _dafny.Map({self.f15: d_1272_v51_})
            d_1274_v53_: _dafny.Seq
            d_1274_v53_ = _dafny.SeqWithoutIsStrInference([d_1272_v51_])
            d_1275_v54_: _dafny.Array
            nw205_ = _dafny.Array(None, 25)
            nw205_[int(0)] = (_dafny.MultiSet(d_1262_v46_)) | ((d_1272_v51_).set(True, default__.abs(self.f15)))
            nw205_[int(1)] = d_1272_v51_
            nw205_[int(2)] = (d_1272_v51_) - (d_1272_v51_)
            nw205_[int(3)] = (d_1272_v51_).intersection(d_1272_v51_)
            nw205_[int(4)] = d_1272_v51_
            nw205_[int(5)] = d_1272_v51_
            nw205_[int(6)] = d_1272_v51_
            nw205_[int(7)] = d_1272_v51_
            nw205_[int(8)] = (d_1272_v51_) - (d_1272_v51_)
            nw205_[int(9)] = d_1272_v51_
            nw205_[int(10)] = (d_1272_v51_).intersection(default__.fm48(self.f15, (self).f6, globalState))
            nw205_[int(11)] = (d_1272_v51_) | (d_1272_v51_)
            nw205_[int(12)] = d_1272_v51_
            nw205_[int(13)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(True))]))
            nw205_[int(14)] = (d_1272_v51_) | (d_1272_v51_)
            nw205_[int(15)] = d_1272_v51_
            nw205_[int(16)] = ((d_1273_v52_)[(self).fm7(d_1254_cf0_, globalState)] if ((self).fm7(d_1254_cf0_, globalState)) in (d_1273_v52_) else d_1272_v51_)
            nw205_[int(17)] = (d_1274_v53_)[default__.safeIndex(self.f15, len(d_1274_v53_))]
            nw205_[int(18)] = (d_1272_v51_) - (d_1272_v51_)
            nw205_[int(19)] = (_dafny.MultiSet(d_1262_v46_)).set((d_1262_v46_)[default__.safeIndex(777, len(d_1262_v46_))], default__.abs(self.f15))
            nw205_[int(20)] = (_dafny.MultiSet([d_1254_cf0_, d_1254_cf0_])) | (d_1272_v51_)
            nw205_[int(21)] = _dafny.MultiSet([d_1254_cf0_])
            nw205_[int(22)] = d_1272_v51_
            nw205_[int(23)] = d_1272_v51_
            nw205_[int(24)] = (d_1272_v51_) - (d_1272_v51_)
            d_1275_v54_ = nw205_
            d_1276_v55_: _dafny.Array
            nw206_ = _dafny.Array(_dafny.Array(None, 0), 21)
            d_1276_v55_ = nw206_
            d_1277_v57_: _dafny.MultiSet
            d_1277_v57_ = _dafny.MultiSet([len(d_1257_v41_)])
            d_1278_v59_: _dafny.Array
            nw207_ = _dafny.Array(None, 28)
            nw207_[int(0)] = d_1265_v49_
            nw207_[int(1)] = d_1265_v49_
            nw207_[int(2)] = d_1265_v49_
            nw207_[int(3)] = d_1265_v49_
            nw207_[int(4)] = d_1265_v49_
            nw207_[int(5)] = d_1265_v49_
            nw207_[int(6)] = d_1265_v49_
            nw207_[int(7)] = d_1265_v49_
            def iife76_():
                coll44_ = _dafny.Map()
                compr_44_: int
                for compr_44_ in (d_1277_v57_).Elements:
                    d_1279_v56_: int = compr_44_
                    if (d_1279_v56_) in (d_1277_v57_):
                        coll44_[(d_1279_v56_) + (len(d_1262_v46_))] = d_1254_cf0_
                return _dafny.Map(coll44_)
            nw207_[int(8)] = iife76_()
            
            nw207_[int(9)] = d_1265_v49_
            nw207_[int(10)] = d_1265_v49_
            nw207_[int(11)] = d_1265_v49_
            nw207_[int(12)] = d_1265_v49_
            nw207_[int(13)] = d_1265_v49_
            nw207_[int(14)] = d_1265_v49_
            def iife77_():
                coll45_ = _dafny.Map()
                compr_45_: int
                for compr_45_ in _dafny.IntegerRange(-991, 889):
                    d_1280_v58_: int = compr_45_
                    if ((-991) <= (d_1280_v58_)) and ((d_1280_v58_) < (889)):
                        coll45_[(d_1280_v58_) + (self.f15)] = (D14_DC26(False, self.f15, (self).f5)).cf32
                return _dafny.Map(coll45_)
            nw207_[int(15)] = iife77_()
            
            nw207_[int(16)] = d_1265_v49_
            nw207_[int(17)] = (d_1265_v49_).set(self.f15, (self).f6)
            nw207_[int(18)] = d_1265_v49_
            nw207_[int(19)] = d_1265_v49_
            nw207_[int(20)] = d_1265_v49_
            nw207_[int(21)] = d_1265_v49_
            nw207_[int(22)] = d_1265_v49_
            nw207_[int(23)] = d_1265_v49_
            nw207_[int(24)] = d_1265_v49_
            nw207_[int(25)] = _dafny.Map({self.f15: (self).f6})
            nw207_[int(26)] = d_1265_v49_
            nw207_[int(27)] = d_1265_v49_
            d_1278_v59_ = nw207_
            index229_ = default__.safeIndex(693, (d_1276_v55_).length(0))
            (d_1276_v55_)[index229_] = d_1278_v59_
            index230_ = default__.safeIndex(764, (d_1260_v44_).length(0))
            index231_ = default__.safeIndex(637, (d_1266_v50_).length(0))
            index232_ = default__.safeIndex(693, (d_1276_v55_).length(0))
            rhs242_ = d_1255_v40_
            rhs243_ = not((d_1261_v45_).issubset((d_1261_v45_ if not(not(True)) else d_1261_v45_)))
            rhs244_ = (d_1259_v43_).set(default__.safeIndex(self.f15, len(d_1259_v43_)), 777)
            rhs245_ = d_1275_v54_
            rhs246_ = d_1278_v59_
            lhs169_ = d_1260_v44_
            lhs170_ = default__.safeIndex(764, (d_1260_v44_).length(0))
            lhs171_ = d_1266_v50_
            lhs172_ = default__.safeIndex(637, (d_1266_v50_).length(0))
            lhs173_ = d_1276_v55_
            lhs174_ = default__.safeIndex(693, (d_1276_v55_).length(0))
            d_1255_v40_ = rhs242_
            lhs169_[lhs170_] = rhs243_
            lhs171_[lhs172_] = rhs244_
            d_1275_v54_ = rhs245_
            lhs173_[lhs174_] = rhs246_
            d_1281_v60_: _dafny.Map
            d_1281_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erhqhi")): d_1257_v41_})
            d_1281_v60_ = d_1281_v60_
            d_1282_v61_: _dafny.Map
            d_1282_v61_ = _dafny.Map({(d_1260_v44_)[default__.safeIndex(764, (d_1260_v44_).length(0))]: d_1255_v40_})
            d_1283_v62_: _dafny.Array
            d_1283_v62_ = (d_1255_v40_ if not(False) else ((d_1282_v61_)[(d_1262_v46_)[default__.safeIndex(self.f15, len(d_1262_v46_))]] if ((d_1262_v46_)[default__.safeIndex(self.f15, len(d_1262_v46_))]) in (d_1282_v61_) else d_1255_v40_))
            rhs247_ = (d_1261_v45_) != (d_1261_v45_)
            rhs248_ = d_1261_v45_
            rhs249_ = d_1283_v62_
            d_1254_cf0_ = rhs247_
            d_1261_v45_ = rhs248_
            d_1283_v62_ = rhs249_
            (self).f15 = 75
        elif True:
            d_1284___mcc_h3_ = source18_.cf3
            d_1285_cf3_ = d_1284___mcc_h3_
            if (self).f6:
                d_1286_v63_: bool
                d_1286_v63_ = True
                d_1287_v64_: _dafny.Map
                d_1287_v64_ = _dafny.Map({(self).f5: self.f15})
                d_1288_v65_: _dafny.Map
                d_1288_v65_ = _dafny.Map({d_1286_v63_: (self).f5})
                def iife78_():
                    coll46_ = _dafny.Map()
                    compr_46_: int
                    for compr_46_ in _dafny.IntegerRange(938, 755):
                        d_1289_v66_: int = compr_46_
                        if ((938) <= (d_1289_v66_)) and ((d_1289_v66_) < (755)):
                            def iife79_():
                                coll47_ = _dafny.Map()
                                compr_47_: str
                                for compr_47_ in (_dafny.Set({(self).f5, (self).f5, (self).f5})).Elements:
                                    d_1290_v67_: str = compr_47_
                                    if (d_1290_v67_) in (_dafny.Set({(self).f5, (self).f5, (self).f5})):
                                        coll47_[d_1290_v67_] = False
                                return _dafny.Map(coll47_)
                            coll46_[(d_1289_v66_) * (198)] = len(_dafny.Map({998: iife79_()
                            }))
                    return _dafny.Map(coll46_)
                d_1286_v63_ = (_dafny.Map({self.f15: ((d_1287_v64_)[((d_1288_v65_)[d_1286_v63_] if (d_1286_v63_) in (d_1288_v65_) else (self).f5)] if (((d_1288_v65_)[d_1286_v63_] if (d_1286_v63_) in (d_1288_v65_) else (self).f5)) in (d_1287_v64_) else self.f15)})) == (iife78_()
                )
                d_1286_v63_ = d_1286_v63_
                d_1291_v68_: _dafny.Array
                def lambda68_(d_1292_i10_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvpi"))

                init38_ = lambda68_
                nw208_ = _dafny.Array(None, 8)
                for i0_38_ in range(nw208_.length(0)):
                    nw208_[i0_38_] = init38_(i0_38_)
                d_1291_v68_ = nw208_
                d_1293_v69_: _dafny.Map
                d_1293_v69_ = _dafny.Map({d_1286_v63_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfuflx"))})
                d_1294_v70_: _dafny.Seq
                d_1294_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bc"))
                index233_ = default__.safeIndex(88, (d_1291_v68_).length(0))
                (d_1291_v68_)[index233_] = (((d_1293_v69_)[(self).f6] if ((self).f6) in (d_1293_v69_) else d_1294_v70_)) + (d_1294_v70_)
                index234_ = default__.safeIndex(88, (d_1291_v68_).length(0))
                (d_1291_v68_)[index234_] = d_1294_v70_
                d_1295_v71_: _dafny.Array
                nw209_ = _dafny.Array(False, 13)
                d_1295_v71_ = nw209_
                index235_ = default__.safeIndex(452, (d_1295_v71_).length(0))
                (d_1295_v71_)[index235_] = not(d_1286_v63_)
                index236_ = default__.safeIndex(452, (d_1295_v71_).length(0))
                (d_1295_v71_)[index236_] = d_1286_v63_
                d_1296_v72_: _dafny.Map
                d_1296_v72_ = _dafny.Map({871: (self).f6})
                d_1297_v73_: _dafny.Map
                d_1297_v73_ = d_1296_v72_
                d_1298_v74_: _dafny.Map
                d_1298_v74_ = _dafny.Map({_dafny.Set({d_1297_v73_}): self.f15})
                def iife80_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in _dafny.IntegerRange(477, -150):
                        d_1299_v75_: int = compr_48_
                        if ((477) <= (d_1299_v75_)) and ((d_1299_v75_) < (-150)):
                            coll48_[(d_1299_v75_) * (len((d_1291_v68_)[default__.safeIndex(88, (d_1291_v68_).length(0))]))] = self.f15
                    return _dafny.Map(coll48_)
                def iife81_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in _dafny.IntegerRange(477, -150):
                        d_1300_v75_: int = compr_49_
                        if ((477) <= (d_1300_v75_)) and ((d_1300_v75_) < (-150)):
                            coll49_[(d_1300_v75_) * (len((d_1291_v68_)[default__.safeIndex(88, (d_1291_v68_).length(0))]))] = self.f15
                    return _dafny.Map(coll49_)
                (globalState).f1 = ((d_1298_v74_)[default__.fm50(default__.fm0(d_1286_v63_, globalState), (self).f6, iife80_()
                , globalState)] if (default__.fm50(default__.fm0(d_1286_v63_, globalState), (self).f6, iife81_()
                , globalState)) in (d_1298_v74_) else self.f15)
            elif True:
                d_1301_v76_: bool
                d_1301_v76_ = False
                d_1302_v77_: _dafny.Map
                d_1302_v77_ = _dafny.Map({(self).f6: d_1301_v76_})
                d_1301_v76_ = ((len(d_1302_v77_)) - (self.f15)) == (self.f15)
                (globalState).f1 = (default__.fm3((self).f6, True, globalState)) + (default__.safeModuloInt(self.f15, self.f15))
                d_1303_v78_: _dafny.Seq
                d_1303_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paf"))
                d_1304_v79_: _dafny.Map
                d_1304_v79_ = _dafny.Map({(self).f6: self.f15})
                d_1305_v80_: _dafny.Map
                d_1305_v80_ = _dafny.Map({len(d_1304_v79_): self.f15})
                d_1306_v81_: _dafny.Seq
                d_1306_v81_ = _dafny.SeqWithoutIsStrInference([self.f15, self.f15, len(d_1305_v80_)])
                d_1307_v82_: _dafny.Array
                nw210_ = _dafny.Array(None, 25)
                nw210_[int(0)] = (0) - (self.f15)
                nw210_[int(1)] = self.f15
                nw210_[int(2)] = (self.f15) + (462)
                nw210_[int(3)] = len(d_1303_v78_)
                nw210_[int(4)] = self.f15
                nw210_[int(5)] = (self.f15) + (len(d_1303_v78_))
                nw210_[int(6)] = self.f15
                nw210_[int(7)] = (159 if True else len(d_1303_v78_))
                nw210_[int(8)] = 72
                nw210_[int(9)] = (d_1306_v81_)[default__.safeIndex(self.f15, len(d_1306_v81_))]
                nw210_[int(10)] = self.f15
                nw210_[int(11)] = self.f15
                nw210_[int(12)] = self.f15
                nw210_[int(13)] = self.f15
                nw210_[int(14)] = self.f15
                nw210_[int(15)] = self.f15
                nw210_[int(16)] = self.f15
                nw210_[int(17)] = (self.f15) * (self.f15)
                nw210_[int(18)] = (0) - (self.f15)
                nw210_[int(19)] = default__.safeModuloInt(self.f15, len(_dafny.Map({not(True): self.f15})))
                nw210_[int(20)] = self.f15
                nw210_[int(21)] = (d_1306_v81_)[default__.safeIndex(self.f15, len(d_1306_v81_))]
                nw210_[int(22)] = self.f15
                nw210_[int(23)] = self.f15
                nw210_[int(24)] = self.f15
                d_1307_v82_ = nw210_
                index237_ = default__.safeIndex(935, (d_1307_v82_).length(0))
                (d_1307_v82_)[index237_] = default__.safeDivisionInt(self.f15, (self).fm8(True, (self).f6, (self).f6, globalState))
                index238_ = default__.safeIndex(935, (d_1307_v82_).length(0))
                (d_1307_v82_)[index238_] = 149
                index239_ = default__.safeIndex(935, (d_1307_v82_).length(0))
                (d_1307_v82_)[index239_] = (d_1307_v82_)[default__.safeIndex(935, (d_1307_v82_).length(0))]
                d_1303_v78_ = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_1308_i11_ in range(default__.abs(51))])) + ((d_1303_v78_).set(default__.safeIndex((d_1307_v82_)[default__.safeIndex(935, (d_1307_v82_).length(0))], len(d_1303_v78_)), (self).f5))
            if (self).f6:
                (globalState).f1 = (355) * (self.f15)
                d_1309_v83_: bool
                d_1309_v83_ = False
                d_1309_v83_ = d_1309_v83_
                d_1310_v84_: _dafny.Seq
                d_1310_v84_ = _dafny.SeqWithoutIsStrInference([self.f15, self.f15, 904, self.f15])
                d_1311_v85_: _dafny.Seq
                d_1311_v85_ = _dafny.SeqWithoutIsStrInference([d_1310_v84_])
                d_1312_v86_: C0
                nw211_ = C0()
                nw211_.ctor__(d_1309_v83_, d_1311_v85_)
                d_1312_v86_ = nw211_
                d_1313_v87_: _dafny.Seq
                d_1313_v87_ = _dafny.SeqWithoutIsStrInference([(D14_DC26(d_1309_v83_, self.f15, (self).f5)).cf32, (self).f6, False])
                d_1314_v88_: _dafny.Map
                d_1314_v88_ = _dafny.Map({_dafny.Set({(d_1312_v86_).f11, d_1309_v83_}): (self).f6})
                d_1315_v89_: _dafny.Set
                d_1315_v89_ = _dafny.Set({(self).fm6(_dafny.SeqWithoutIsStrInference([self.f15 for d_1316_i12_ in range(default__.abs(672))]), globalState)})
                d_1317_v90_: _dafny.Seq
                d_1317_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdvhobhq"))
                d_1318_v91_: _dafny.Map
                d_1318_v91_ = _dafny.Map({(self).fm8(d_1309_v83_, (d_1312_v86_).f11, (d_1312_v86_).f11, globalState): (d_1312_v86_).f11})
                d_1319_v92_: _dafny.Map
                d_1319_v92_ = _dafny.Map({len(d_1310_v84_): d_1317_v90_})
                d_1320_v93_: _dafny.Array
                nw212_ = _dafny.Array(None, 26)
                nw212_[int(0)] = (self).f6
                nw212_[int(1)] = (d_1312_v86_).f11
                nw212_[int(2)] = d_1309_v83_
                nw212_[int(3)] = (self.f15) <= (self.f15)
                nw212_[int(4)] = (d_1312_v86_).f11
                nw212_[int(5)] = (default__.fm0(d_1309_v83_, globalState)) or ((d_1312_v86_).f11)
                nw212_[int(6)] = (d_1312_v86_).f11
                nw212_[int(7)] = (d_1313_v87_)[default__.safeIndex(self.f15, len(d_1313_v87_))]
                nw212_[int(8)] = not ((self).f6) or ((d_1312_v86_).f11)
                nw212_[int(9)] = False
                nw212_[int(10)] = not((d_1312_v86_).f11)
                nw212_[int(11)] = (self.f15) < (self.f15)
                nw212_[int(12)] = (d_1312_v86_).f11
                nw212_[int(13)] = (d_1312_v86_).f11
                nw212_[int(14)] = ((d_1314_v88_)[d_1315_v89_] if (d_1315_v89_) in (d_1314_v88_) else d_1309_v83_)
                nw212_[int(15)] = ((self).f5) in (d_1317_v90_)
                nw212_[int(16)] = (d_1312_v86_).f11
                nw212_[int(17)] = (d_1312_v86_).f11
                nw212_[int(18)] = (d_1312_v86_).f11
                nw212_[int(19)] = (self).f6
                nw212_[int(20)] = ((d_1318_v91_)[len(d_1319_v92_)] if (len(d_1319_v92_)) in (d_1318_v91_) else d_1309_v83_)
                nw212_[int(21)] = (self).f6
                nw212_[int(22)] = ((d_1312_v86_).f11) and (d_1309_v83_)
                nw212_[int(23)] = (self).f6
                nw212_[int(24)] = (d_1312_v86_).f11
                nw212_[int(25)] = True
                d_1320_v93_ = nw212_
                d_1321_v94_: D2
                d_1321_v94_ = D2_DC8(d_1309_v83_)
                rhs250_ = self.f15
                rhs251_ = (((_dafny.MultiSet([d_1309_v83_])).set((d_1312_v86_).f11, default__.abs(self.f15))).set((d_1312_v86_).f11, default__.abs(self.f15))).set(d_1309_v83_, default__.abs(self.f15))
                rhs252_ = (d_1320_v93_ if not((self).f6) else d_1320_v93_)
                rhs253_ = d_1321_v94_
                rhs254_ = default__.safeModuloInt(539, default__.safeModuloInt(len(default__.fm46((d_1312_v86_).f11, globalState)), self.f15))
                lhs175_ = self
                lhs176_ = globalState
                lhs175_.f15 = rhs250_
                lhs176_.f4 = rhs251_
                d_1320_v93_ = rhs252_
                d_1321_v94_ = rhs253_
                r0 = rhs254_
                d_1322_v95_: C3
                nw213_ = C3()
                nw213_.ctor__((default__.fm51(d_1309_v83_, len(d_1318_v91_), (self).f6, globalState)).issubset(_dafny.Set({(self).f5})), (self).f6)
                d_1322_v95_ = nw213_
            elif True:
                d_1323_v96_: _dafny.MultiSet
                d_1323_v96_ = _dafny.MultiSet([(self).f6])
                d_1324_v97_: C3
                nw214_ = C3()
                nw214_.ctor__((d_1323_v96_).issubset(d_1323_v96_), not((self).f6))
                d_1324_v97_ = nw214_
                d_1325_v98_: _dafny.Seq
                d_1325_v98_ = _dafny.SeqWithoutIsStrInference([(self).f5])
                d_1326_v99_: _dafny.Map
                d_1326_v99_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5, (self).f5])) + (d_1325_v98_): self.f15})
                d_1326_v99_ = (d_1326_v99_).set(d_1325_v98_, self.f15)
                (self).f15 = (self.f15) + (len((_dafny.Map({len(d_1325_v98_): (d_1324_v97_).f14})).set(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aknri"))), (d_1324_v97_).f13)))
                d_1327_v100_: _dafny.Array
                def lambda69_(d_1328_i13_):
                    return default__.safeDivisionInt(d_1328_i13_, self.f15)

                init39_ = lambda69_
                nw215_ = _dafny.Array(None, 4)
                for i0_39_ in range(nw215_.length(0)):
                    nw215_[i0_39_] = init39_(i0_39_)
                d_1327_v100_ = nw215_
                index240_ = default__.safeIndex(858, (d_1327_v100_).length(0))
                (d_1327_v100_)[index240_] = self.f15
                index241_ = default__.safeIndex(858, (d_1327_v100_).length(0))
                (d_1327_v100_)[index241_] = self.f15
                d_1329_v101_: D15
                d_1329_v101_ = D15_DC27(d_1323_v96_)
                d_1329_v101_ = d_1329_v101_
            d_1330_v102_: _dafny.Map
            d_1331_v103_: _dafny.Seq
            d_1332_v104_: _dafny.Set
            d_1333_v105_: bool
            out36_: _dafny.Map
            out37_: _dafny.Seq
            out38_: _dafny.Set
            out39_: bool
            out36_, out37_, out38_, out39_ = (self).m12(self.f15, (self).f6, self.f15, globalState)
            d_1330_v102_ = out36_
            d_1331_v103_ = out37_
            d_1332_v104_ = out38_
            d_1333_v105_ = out39_
            d_1334_v106_: _dafny.Array
            nw216_ = _dafny.Array(int(0), 10)
            d_1334_v106_ = nw216_
            index242_ = default__.safeIndex(477, (d_1334_v106_).length(0))
            (d_1334_v106_)[index242_] = self.f15
            d_1335_v107_: _dafny.Seq
            d_1335_v107_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, d_1333_v105_, d_1333_v105_, (self).f6])
            d_1336_v108_: _dafny.MultiSet
            d_1336_v108_ = _dafny.MultiSet([len(d_1331_v103_), self.f15])
            d_1337_v109_: _dafny.MultiSet
            d_1337_v109_ = _dafny.MultiSet([(0) - (self.f15), (self).fm8(False, (self).f6, d_1333_v105_, globalState), (self).fm4(d_1335_v107_, self.f15, self.f15, d_1336_v108_, globalState), self.f15, self.f15])
            d_1338_v110_: _dafny.Map
            d_1338_v110_ = _dafny.Map({(self).f6: self.f15})
            d_1339_v111_: _dafny.Seq
            d_1339_v111_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f15, ((d_1336_v108_)[len(_dafny.Map({self.f15: self.f15}))] if (len(_dafny.Map({self.f15: self.f15}))) in (d_1336_v108_) else self.f15)]): d_1338_v110_})])
            index243_ = default__.safeIndex(477, (d_1334_v106_).length(0))
            (d_1334_v106_)[index243_] = (self.f15) + (len(((d_1339_v111_)[default__.safeIndex(self.f15, len(d_1339_v111_))]) | (_dafny.Map({d_1331_v103_: d_1338_v110_}))))
        (self).f15 = ((self).fm7((self).f6, globalState)) - (self.f15)
        d_1340_v112_: _dafny.MultiSet
        d_1340_v112_ = _dafny.MultiSet([self.f15, 521])
        hi7_ = default__.safeModuloInt(self.f15, self.f15)
        for d_1341_i14_ in range(((d_1340_v112_)[self.f15] if (self.f15) in (d_1340_v112_) else self.f15), hi7_):
            d_1342_v113_: _dafny.Array
            def lambda70_(d_1343_i15_):
                return (d_1343_i15_) * (self.f15)

            init40_ = lambda70_
            nw217_ = _dafny.Array(None, 11)
            for i0_40_ in range(nw217_.length(0)):
                nw217_[i0_40_] = init40_(i0_40_)
            d_1342_v113_ = nw217_
            index244_ = default__.safeIndex(954, (d_1342_v113_).length(0))
            (d_1342_v113_)[index244_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jwrbku")))
            d_1344_v114_: _dafny.Set
            d_1344_v114_ = _dafny.Set({d_1341_i14_})
            d_1345_v115_: _dafny.Seq
            d_1345_v115_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, not ((self).f6) or ((self).f6), (len(d_1344_v114_)) <= (self.f15)])
            d_1346_v116_: _dafny.MultiSet
            d_1346_v116_ = _dafny.MultiSet([(self).f6])
            d_1347_v117_: D15
            d_1347_v117_ = D15_DC27(d_1346_v116_)
            index245_ = default__.safeIndex(954, (d_1342_v113_).length(0))
            rhs255_ = (0) - (self.f15)
            rhs256_ = d_1345_v115_
            rhs257_ = -813
            rhs258_ = D15_DC27(d_1346_v116_)
            lhs177_ = d_1342_v113_
            lhs178_ = default__.safeIndex(954, (d_1342_v113_).length(0))
            lhs179_ = self
            lhs177_[lhs178_] = rhs255_
            d_1345_v115_ = rhs256_
            lhs179_.f15 = rhs257_
            d_1347_v117_ = rhs258_
            index246_ = default__.safeIndex(954, (d_1342_v113_).length(0))
            (d_1342_v113_)[index246_] = 327
            d_1348_v118_: _dafny.Array
            nw218_ = _dafny.Array(D2.default()(), 1)
            d_1348_v118_ = nw218_
            d_1348_v118_ = d_1348_v118_
            d_1349_v119_: bool
            d_1349_v119_ = False
            d_1350_v120_: _dafny.MultiSet
            d_1350_v120_ = _dafny.MultiSet([_dafny.CodePoint('o'), (self).f5, (self).f5, (self).f5, _dafny.CodePoint('c')])
            d_1349_v119_ = ((0) - (d_1341_i14_)) < ((d_1350_v120_).cardinality)
        d_1351_v121_: bool
        d_1351_v121_ = False
        d_1351_v121_ = (d_1351_v121_) and (d_1351_v121_)
        d_1352_v122_: _dafny.Seq
        d_1352_v122_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqjcniels"))
        d_1352_v122_ = d_1352_v122_
        d_1353_v123_: _dafny.Array
        nw219_ = _dafny.Array(int(0), 11)
        d_1353_v123_ = nw219_
        d_1354_v124_: C1
        nw220_ = C1()
        nw220_.ctor__()
        d_1354_v124_ = nw220_
        d_1355_v125_: _dafny.Array
        nw221_ = _dafny.Array(False, 26)
        d_1355_v125_ = nw221_
        d_1356_v126_: _dafny.Map
        d_1356_v126_ = _dafny.Map({d_1354_v124_: d_1355_v125_})
        d_1357_v127_: _dafny.Map
        d_1357_v127_ = d_1356_v126_
        index247_ = default__.safeIndex(712, (d_1353_v123_).length(0))
        (d_1353_v123_)[index247_] = len(_dafny.SeqWithoutIsStrInference([d_1357_v127_]))
        d_1358_v128_: _dafny.Seq
        d_1358_v128_ = _dafny.SeqWithoutIsStrInference([self.f15, self.f15])
        d_1359_v129_: _dafny.Seq
        d_1359_v129_ = _dafny.SeqWithoutIsStrInference([d_1358_v128_])
        d_1360_v130_: _dafny.Set
        d_1360_v130_ = _dafny.Set({d_1351_v121_})
        d_1361_v131_: _dafny.Map
        d_1361_v131_ = _dafny.Map({(d_1359_v129_).set(default__.safeIndex((0) - (len(d_1360_v130_)), len(d_1359_v129_)), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utb"))) for d_1362_i16_ in range(default__.abs(764))])): d_1352_v122_})
        index248_ = default__.safeIndex(712, (d_1353_v123_).length(0))
        rhs259_ = self.f15
        rhs260_ = ((d_1361_v131_)[(d_1359_v129_) + (default__.fm52((self).f6, False, globalState))] if ((d_1359_v129_) + (default__.fm52((self).f6, False, globalState))) in (d_1361_v131_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amowydn")))
        lhs180_ = d_1353_v123_
        lhs181_ = default__.safeIndex(712, (d_1353_v123_).length(0))
        lhs180_[lhs181_] = rhs259_
        d_1352_v122_ = rhs260_
        r0 = (d_1358_v128_)[default__.safeIndex((d_1353_v123_)[default__.safeIndex(712, (d_1353_v123_).length(0))], len(d_1358_v128_))]
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_1363_v0_: _dafny.Map
        d_1363_v0_ = _dafny.Map({p0: self.f15})
        hi8_ = (self.f15) * ((0) - (self.f15))
        for d_1364_i0_ in range(((d_1363_v0_)[default__.fm0(p2, globalState)] if (default__.fm0(p2, globalState)) in (d_1363_v0_) else self.f15), hi8_):
            d_1365_v1_: _dafny.Seq
            d_1365_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "petsff"))
            d_1366_v2_: _dafny.Map
            d_1366_v2_ = _dafny.Map({self.f15: d_1365_v1_})
            d_1367_v3_: _dafny.Array
            nw222_ = _dafny.Array(int(0), 12)
            d_1367_v3_ = nw222_
            d_1368_v4_: _dafny.Map
            d_1368_v4_ = _dafny.Map({d_1367_v3_: d_1364_i0_})
            d_1369_v5_: _dafny.Map
            d_1369_v5_ = _dafny.Map({False: not((self).f6)})
            d_1370_v6_: _dafny.Array
            nw223_ = _dafny.Array(None, 25)
            nw223_[int(0)] = d_1364_i0_
            nw223_[int(1)] = self.f15
            nw223_[int(2)] = ((self).fm7(p2, globalState)) + ((0) - (d_1364_i0_))
            nw223_[int(3)] = self.f15
            nw223_[int(4)] = self.f15
            nw223_[int(5)] = (d_1364_i0_) * (self.f15)
            nw223_[int(6)] = (0) - (len((d_1365_v1_ if (self).f6 else d_1365_v1_)))
            nw223_[int(7)] = d_1364_i0_
            nw223_[int(8)] = (0) - (len(d_1366_v2_))
            nw223_[int(9)] = (0) - (d_1364_i0_)
            nw223_[int(10)] = (0) - ((0) - (((d_1368_v4_)[d_1367_v3_] if (d_1367_v3_) in (d_1368_v4_) else d_1364_i0_)))
            nw223_[int(11)] = (0) - ((0) - (d_1364_i0_))
            nw223_[int(12)] = d_1364_i0_
            nw223_[int(13)] = self.f15
            nw223_[int(14)] = (317) * (len(d_1365_v1_))
            nw223_[int(15)] = len(d_1365_v1_)
            nw223_[int(16)] = 362
            nw223_[int(17)] = (0) - (default__.safeModuloInt(self.f15, d_1364_i0_))
            nw223_[int(18)] = d_1364_i0_
            nw223_[int(19)] = self.f15
            nw223_[int(20)] = self.f15
            nw223_[int(21)] = self.f15
            nw223_[int(22)] = (0) - (d_1364_i0_)
            nw223_[int(23)] = self.f15
            nw223_[int(24)] = (self.f15) * (len(d_1369_v5_))
            d_1370_v6_ = nw223_
            d_1367_v3_ = d_1370_v6_
            if default__.fm0(not ((self).f6) or (p2), globalState):
                (globalState).f1 = ((0) - (d_1364_i0_)) + (self.f15)
                d_1371_v7_: bool
                d_1371_v7_ = True
                d_1371_v7_ = p2
                d_1372_v8_: _dafny.Array
                nw224_ = _dafny.Array(False, 6)
                d_1372_v8_ = nw224_
                index249_ = default__.safeIndex(256, (d_1372_v8_).length(0))
                (d_1372_v8_)[index249_] = d_1371_v7_
                d_1373_v9_: D11
                d_1373_v9_ = D11_DC20(d_1371_v7_, 77)
                index250_ = default__.safeIndex(256, (d_1372_v8_).length(0))
                (d_1372_v8_)[index250_] = (d_1373_v9_).cf21
                d_1374_v10_: _dafny.Seq
                d_1374_v10_ = _dafny.SeqWithoutIsStrInference([d_1372_v8_, d_1372_v8_, d_1372_v8_])
                d_1375_v11_: _dafny.Array
                nw225_ = _dafny.Array(None, 16)
                nw225_[int(0)] = d_1372_v8_
                nw225_[int(1)] = d_1372_v8_
                nw225_[int(2)] = d_1372_v8_
                nw225_[int(3)] = (d_1372_v8_ if (self).f6 else d_1372_v8_)
                nw225_[int(4)] = d_1372_v8_
                nw225_[int(5)] = d_1372_v8_
                nw225_[int(6)] = d_1372_v8_
                nw225_[int(7)] = d_1372_v8_
                nw225_[int(8)] = d_1372_v8_
                nw225_[int(9)] = d_1372_v8_
                nw225_[int(10)] = d_1372_v8_
                nw225_[int(11)] = d_1372_v8_
                nw225_[int(12)] = d_1372_v8_
                nw225_[int(13)] = d_1372_v8_
                nw225_[int(14)] = (d_1374_v10_)[default__.safeIndex(d_1364_i0_, len(d_1374_v10_))]
                nw225_[int(15)] = d_1372_v8_
                d_1375_v11_ = nw225_
                d_1375_v11_ = d_1375_v11_
                d_1376_v12_: _dafny.Map
                d_1376_v12_ = _dafny.Map({d_1364_i0_: len(_dafny.SeqWithoutIsStrInference([(d_1372_v8_)[default__.safeIndex(256, (d_1372_v8_).length(0))]]))})
                d_1377_v13_: _dafny.Map
                d_1377_v13_ = _dafny.Map({self.f15: len(d_1376_v12_)})
                d_1378_v14_: _dafny.MultiSet
                d_1378_v14_ = _dafny.MultiSet([(self).f5])
                d_1371_v7_ = (_dafny.MultiSet([(self).f5, default__.fm53(d_1365_v1_, (0) - (len(d_1377_v13_)), globalState)])).isdisjoint((d_1378_v14_) | (_dafny.MultiSet([(self).f5, (self).f5])))
            elif True:
                d_1379_v15_: bool
                d_1379_v15_ = True
                d_1379_v15_ = (d_1364_i0_) < (d_1364_i0_)
                d_1380_v16_: _dafny.Array
                def lambda71_(d_1381_p0_):
                    def lambda72_(d_1382_i1_):
                        return d_1381_p0_

                    return lambda72_

                init41_ = lambda71_(p0)
                nw226_ = _dafny.Array(None, 27)
                for i0_41_ in range(nw226_.length(0)):
                    nw226_[i0_41_] = init41_(i0_41_)
                d_1380_v16_ = nw226_
                d_1383_v17_: str
                d_1383_v17_ = _dafny.CodePoint('b')
                d_1384_v18_: _dafny.Seq
                d_1384_v18_ = _dafny.SeqWithoutIsStrInference([len(d_1365_v1_), d_1364_i0_])
                rhs261_ = self.f15
                rhs262_ = d_1380_v16_
                rhs263_ = d_1383_v17_
                rhs264_ = (self).fm6(d_1384_v18_, globalState)
                lhs182_ = globalState
                lhs182_.f1 = rhs261_
                d_1380_v16_ = rhs262_
                d_1383_v17_ = rhs263_
                d_1379_v15_ = rhs264_
                (globalState).f1 = self.f15
                (globalState).f4 = (p1).intersection(_dafny.MultiSet([p0]))
                d_1383_v17_ = d_1383_v17_
            d_1385_v19_: _dafny.Map
            d_1385_v19_ = _dafny.Map({self.f15: p2})
            d_1385_v19_ = (d_1385_v19_).set(default__.safeDivisionInt(self.f15, self.f15), (self).f6)
            d_1386_v20_: C3
            nw227_ = C3()
            nw227_.ctor__((self).f6, p2)
            d_1386_v20_ = nw227_
        d_1387_v21_: _dafny.Seq
        d_1387_v21_ = _dafny.SeqWithoutIsStrInference([(self).f6])
        d_1387_v21_ = (_dafny.SeqWithoutIsStrInference([(self).f6])) + ((d_1387_v21_).set(default__.safeIndex(self.f15, len(d_1387_v21_)), (self).f6))
        d_1388_v22_: _dafny.Map
        d_1388_v22_ = _dafny.Map({self.f15: self.f15})
        d_1389_v23_: _dafny.Map
        d_1389_v23_ = d_1388_v22_
        d_1390_v24_: _dafny.Seq
        d_1390_v24_ = _dafny.SeqWithoutIsStrInference([d_1387_v21_])
        d_1391_v25_: _dafny.Seq
        d_1391_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vinlx"))
        d_1392_v26_: D12
        d_1392_v26_ = D12_DC22(d_1389_v23_, d_1390_v24_, (d_1391_v25_)[default__.safeIndex(self.f15, len(d_1391_v25_))], p0, d_1391_v25_)
        source20_ = d_1392_v26_
        if source20_.is_DC22:
            d_1393___mcc_h0_ = source20_.cf24
            d_1394___mcc_h1_ = source20_.cf25
            d_1395___mcc_h2_ = source20_.cf26
            d_1396___mcc_h3_ = source20_.cf27
            d_1397___mcc_h4_ = source20_.cf28
            d_1398_cf28_ = d_1397___mcc_h4_
            d_1399_cf27_ = d_1396___mcc_h3_
            d_1400_cf26_ = d_1395___mcc_h2_
            d_1401_cf25_ = d_1394___mcc_h1_
            d_1402_cf24_ = d_1393___mcc_h0_
            d_1391_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnlpvun"))
            if (self).f6:
                d_1398_cf28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "stpkvw"))
                d_1403_v27_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.Map({}), 15)
                d_1403_v27_ = nw228_
                d_1404_v28_: _dafny.Array
                nw229_ = _dafny.Array(False, 12)
                d_1404_v28_ = nw229_
                d_1405_v29_: _dafny.Map
                d_1405_v29_ = _dafny.Map({d_1404_v28_: self.f15})
                index251_ = default__.safeIndex(497, (d_1403_v27_).length(0))
                (d_1403_v27_)[index251_] = d_1405_v29_
                index252_ = default__.safeIndex(497, (d_1403_v27_).length(0))
                (d_1403_v27_)[index252_] = d_1405_v29_
                d_1406_v30_: _dafny.Map
                d_1406_v30_ = _dafny.Map({d_1400_cf26_: self.f15})
                d_1406_v30_ = (d_1406_v30_).set(default__.fm53(d_1391_v25_, len(_dafny.Map({p2: not(d_1399_cf27_)})), globalState), self.f15)
                d_1407_v31_: _dafny.MultiSet
                d_1407_v31_ = _dafny.MultiSet([self.f15])
                d_1399_cf27_ = (d_1407_v31_).isdisjoint(d_1407_v31_)
                (self).f15 = 64
            elif True:
                d_1408_v32_: _dafny.Seq
                d_1408_v32_ = _dafny.SeqWithoutIsStrInference([self.f15])
                d_1409_v33_: _dafny.Seq
                d_1409_v33_ = _dafny.SeqWithoutIsStrInference([d_1408_v32_])
                d_1410_v34_: C0
                nw230_ = C0()
                nw230_.ctor__(not(not(d_1399_cf27_)), (d_1409_v33_) + (d_1409_v33_))
                d_1410_v34_ = nw230_
                d_1411_v35_: _dafny.Array
                nw231_ = _dafny.Array(False, 19)
                d_1411_v35_ = nw231_
                d_1412_v36_: _dafny.Set
                d_1412_v36_ = _dafny.Set({self.f15})
                index253_ = default__.safeIndex(58, (d_1411_v35_).length(0))
                (d_1411_v35_)[index253_] = (self.f15) not in (d_1412_v36_)
                d_1413_v37_: _dafny.Map
                d_1413_v37_ = _dafny.Map({d_1411_v35_: _dafny.Set({563})})
                index254_ = default__.safeIndex(58, (d_1411_v35_).length(0))
                rhs265_ = not((p0) == ((d_1387_v21_) != (d_1387_v21_)))
                rhs266_ = ((p1).intersection(_dafny.MultiSet([p2, p2, (self).f6]))) | (_dafny.MultiSet([p0]))
                rhs267_ = (_dafny.Map({d_1411_v35_: d_1412_v36_}) if (self).fm6(d_1408_v32_, globalState) else d_1413_v37_)
                lhs183_ = d_1411_v35_
                lhs184_ = default__.safeIndex(58, (d_1411_v35_).length(0))
                lhs185_ = globalState
                lhs183_[lhs184_] = rhs265_
                lhs185_.f4 = rhs266_
                d_1413_v37_ = rhs267_
                d_1414_v38_: C3
                nw232_ = C3()
                nw232_.ctor__(not(p0), (d_1411_v35_)[default__.safeIndex(58, (d_1411_v35_).length(0))])
                d_1414_v38_ = nw232_
                index255_ = default__.safeIndex(58, (d_1411_v35_).length(0))
                rhs268_ = (d_1411_v35_)[default__.safeIndex(58, (d_1411_v35_).length(0))]
                rhs269_ = (d_1410_v34_).f11
                lhs186_ = d_1411_v35_
                lhs187_ = default__.safeIndex(58, (d_1411_v35_).length(0))
                lhs186_[lhs187_] = rhs268_
                d_1399_cf27_ = rhs269_
                d_1415_v39_: _dafny.Map
                d_1415_v39_ = _dafny.Map({(d_1414_v38_).f14: p2})
                d_1416_v40_: C3
                nw233_ = C3()
                nw233_.ctor__(not((d_1415_v39_) == ((d_1415_v39_).set((self).f6, p0))), p2)
                d_1416_v40_ = nw233_
            d_1417_v41_: _dafny.Array
            nw234_ = _dafny.Array(_dafny.Set({}), 20)
            d_1417_v41_ = nw234_
            d_1418_v42_: _dafny.Seq
            d_1418_v42_ = _dafny.SeqWithoutIsStrInference([self.f15])
            d_1419_v43_: _dafny.Map
            d_1419_v43_ = _dafny.Map({(self).f6: (self).fm6(d_1418_v42_, globalState)})
            d_1420_v44_: _dafny.Map
            d_1420_v44_ = _dafny.Map({self.f15: d_1419_v43_})
            d_1421_v45_: _dafny.Set
            d_1421_v45_ = _dafny.Set({((d_1420_v44_)[self.f15] if (self.f15) in (d_1420_v44_) else d_1419_v43_)})
            index256_ = default__.safeIndex(740, (d_1417_v41_).length(0))
            (d_1417_v41_)[index256_] = d_1421_v45_
            d_1422_v46_: _dafny.Map
            d_1422_v46_ = _dafny.Map({p0: d_1421_v45_})
            d_1423_v47_: _dafny.Seq
            d_1423_v47_ = _dafny.SeqWithoutIsStrInference([d_1419_v43_])
            index257_ = default__.safeIndex(740, (d_1417_v41_).length(0))
            def iife82_():
                coll50_ = _dafny.Set()
                compr_50_: _dafny.Map
                for compr_50_ in (d_1423_v47_).Elements:
                    d_1424_v48_: _dafny.Map = compr_50_
                    if (d_1424_v48_) in (d_1423_v47_):
                        coll50_ = coll50_.union(_dafny.Set([d_1424_v48_]))
                return _dafny.Set(coll50_)
            (d_1417_v41_)[index257_] = ((d_1422_v46_)[(p0) and (p0)] if ((p0) and (p0)) in (d_1422_v46_) else iife82_()
            )
            d_1425_v49_: _dafny.Array
            nw235_ = _dafny.Array(False, 28)
            d_1425_v49_ = nw235_
            d_1426_v50_: _dafny.Map
            d_1426_v50_ = _dafny.Map({d_1425_v49_: default__.safeModuloInt(self.f15, self.f15)})
            d_1426_v50_ = (d_1426_v50_).set(d_1425_v49_, default__.safeDivisionInt(self.f15, self.f15))
        elif source20_.is_DC23:
            d_1427___mcc_h5_ = source20_.cf29
            d_1428_cf29_ = d_1427___mcc_h5_
            d_1429_v51_: D15
            d_1429_v51_ = D15_DC28((False) == (p0), ((D1_DC4(self.f15)).cf4) <= ((_dafny.MultiSet([_dafny.CodePoint('n'), (self).f5, _dafny.CodePoint('g'), (self).f5, (self).f5])).cardinality))
            pat_let_tv18_ = p0
            def iife83_(_pat_let16_0):
                def iife84_(d_1430_dt__update__tmp_h0_):
                    def iife85_(_pat_let17_0):
                        def iife86_(d_1431_dt__update_hcf37_h0_):
                            return D15_DC28((d_1430_dt__update__tmp_h0_).cf36, d_1431_dt__update_hcf37_h0_)
                        return iife86_(_pat_let17_0)
                    return iife85_(pat_let_tv18_)
                return iife84_(_pat_let16_0)
            d_1429_v51_ = iife83_(default__.fm54(globalState))
            (globalState).f1 = self.f15
            (globalState).f1 = self.f15
            d_1432_v52_: _dafny.Map
            d_1432_v52_ = _dafny.Map({self.f15: p2})
            d_1433_v53_: _dafny.Map
            d_1434_v54_: _dafny.Seq
            d_1435_v55_: _dafny.Set
            d_1436_v56_: bool
            out40_: _dafny.Map
            out41_: _dafny.Seq
            out42_: _dafny.Set
            out43_: bool
            out40_, out41_, out42_, out43_ = (self).m12(len(default__.fm55(self.f15, self.f15, globalState)), ((0) - (len(d_1432_v52_))) >= (self.f15), self.f15, globalState)
            d_1433_v53_ = out40_
            d_1434_v54_ = out41_
            d_1435_v55_ = out42_
            d_1436_v56_ = out43_
        elif True:
            d_1437___mcc_h6_ = source20_.cf23
            d_1438_cf23_ = d_1437___mcc_h6_
            d_1439_v57_: _dafny.Array
            def lambda73_(d_1440_i2_):
                return (self).f5

            init42_ = lambda73_
            nw236_ = _dafny.Array(None, 24)
            for i0_42_ in range(nw236_.length(0)):
                nw236_[i0_42_] = init42_(i0_42_)
            d_1439_v57_ = nw236_
            d_1439_v57_ = d_1439_v57_
            (globalState).f1 = 229
            d_1441_v58_: _dafny.Array
            nw237_ = _dafny.Array(_dafny.MultiSet({}), 28)
            d_1441_v58_ = nw237_
            index258_ = default__.safeIndex(320, (d_1441_v58_).length(0))
            (d_1441_v58_)[index258_] = (_dafny.MultiSet([(self).f6])).intersection(p1)
            index259_ = default__.safeIndex(320, (d_1441_v58_).length(0))
            (d_1441_v58_)[index259_] = p1
            d_1442_v59_: _dafny.Array
            nw238_ = _dafny.Array(False, 14)
            d_1442_v59_ = nw238_
            d_1443_v60_: _dafny.Map
            d_1443_v60_ = _dafny.Map({d_1442_v59_: ((self).f6 if (self).f6 else p0)})
            if ((d_1443_v60_)[d_1442_v59_] if (d_1442_v59_) in (d_1443_v60_) else True):
                (self).f15 = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yl")))
                d_1444_v61_: _dafny.Map
                d_1444_v61_ = _dafny.Map({self.f15: d_1442_v59_})
                d_1444_v61_ = (d_1444_v61_).set(self.f15, d_1442_v59_)
                d_1445_v62_: bool
                d_1445_v62_ = False
                d_1445_v62_ = p0
                (globalState).f1 = (0) - ((self.f15) + (self.f15))
                d_1446_v63_: _dafny.Seq
                d_1446_v63_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f6: (0) - (len(d_1391_v25_))})])
                d_1447_v64_: _dafny.Map
                d_1447_v64_ = _dafny.Map({d_1446_v63_: _dafny.CodePoint('p')})
                d_1447_v64_ = (d_1447_v64_).set((d_1446_v63_) + (_dafny.SeqWithoutIsStrInference([d_1363_v0_ for d_1448_i3_ in range(default__.abs(4))])), (self).f5)
            elif True:
                d_1449_v65_: bool
                d_1449_v65_ = False
                d_1450_v66_: _dafny.Array
                def lambda74_(d_1451_i4_):
                    return default__.safeDivisionInt(d_1451_i4_, self.f15)

                init43_ = lambda74_
                nw239_ = _dafny.Array(None, 22)
                for i0_43_ in range(nw239_.length(0)):
                    nw239_[i0_43_] = init43_(i0_43_)
                d_1450_v66_ = nw239_
                d_1452_v67_: _dafny.Map
                d_1452_v67_ = _dafny.Map({d_1450_v66_: p0})
                d_1453_v68_: _dafny.Seq
                d_1453_v68_ = _dafny.SeqWithoutIsStrInference([d_1450_v66_, d_1450_v66_])
                d_1454_v69_: D1
                d_1454_v69_ = D1_DC5(d_1449_v65_)
                d_1449_v65_ = not(not(((d_1452_v67_)[(d_1453_v68_)[default__.safeIndex(len(d_1387_v21_), len(d_1453_v68_))]] if ((d_1453_v68_)[default__.safeIndex(len(d_1387_v21_), len(d_1453_v68_))]) in (d_1452_v67_) else (d_1454_v69_).cf5)))
                (globalState).f1 = self.f15
                d_1455_v70_: _dafny.MultiSet
                d_1455_v70_ = _dafny.MultiSet([self.f15])
                d_1456_v71_: _dafny.Seq
                d_1456_v71_ = _dafny.SeqWithoutIsStrInference([len(d_1391_v25_)])
                d_1457_v72_: _dafny.Seq
                d_1457_v72_ = _dafny.SeqWithoutIsStrInference([((d_1455_v70_)[((p1)[d_1449_v65_] if (d_1449_v65_) in (p1) else self.f15)] if (((p1)[d_1449_v65_] if (d_1449_v65_) in (p1) else self.f15)) in (d_1455_v70_) else self.f15), len(d_1456_v71_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf"))), self.f15, self.f15])
                d_1458_v73_: _dafny.MultiSet
                d_1458_v73_ = _dafny.MultiSet([len(d_1456_v71_)])
                d_1459_v74_: _dafny.Map
                d_1459_v74_ = _dafny.Map({d_1455_v70_: self.f15})
                (self).f15 = (((d_1459_v74_)[d_1458_v73_] if (d_1458_v73_) in (d_1459_v74_) else self.f15) if ((self).f6 if p0 else p0) else self.f15)
                d_1460_v75_: str
                d_1460_v75_ = _dafny.CodePoint('k')
                d_1461_v76_: D1
                d_1461_v76_ = D1_DC4(self.f15)
                rhs270_ = not(p0)
                rhs271_ = (p2) and (not((p1).ispropersubset((d_1441_v58_)[default__.safeIndex(320, (d_1441_v58_).length(0))])))
                rhs272_ = not ((self).f6) or (not(default__.fm0((self).f6, globalState)))
                rhs273_ = d_1460_v75_
                rhs274_ = d_1461_v76_
                d_1449_v65_ = rhs270_
                d_1449_v65_ = rhs271_
                d_1449_v65_ = rhs272_
                d_1460_v75_ = rhs273_
                d_1461_v76_ = rhs274_
                d_1449_v65_ = (d_1449_v65_) == (p2)
        d_1462_v77_: D15
        d_1462_v77_ = D15_DC28(p2, True)
        d_1463_v78_: D15
        d_1463_v78_ = D15_DC29(d_1462_v77_)
        source21_ = d_1463_v78_
        if source21_.is_DC28:
            d_1464___mcc_h7_ = source21_.cf36
            d_1465___mcc_h8_ = source21_.cf37
            d_1466_cf37_ = d_1465___mcc_h8_
            d_1467_cf36_ = d_1464___mcc_h7_
            d_1468_v79_: _dafny.Array
            nw240_ = _dafny.Array(int(0), 3)
            d_1468_v79_ = nw240_
            index260_ = default__.safeIndex(574, (d_1468_v79_).length(0))
            (d_1468_v79_)[index260_] = self.f15
            d_1469_v80_: _dafny.Seq
            d_1469_v80_ = _dafny.SeqWithoutIsStrInference([-106])
            d_1470_v81_: _dafny.MultiSet
            d_1470_v81_ = _dafny.MultiSet([(self.f15) * (self.f15), (d_1469_v80_)[default__.safeIndex(851, len(d_1469_v80_))], self.f15])
            d_1471_v82_: _dafny.Map
            d_1471_v82_ = _dafny.Map({(self).f6: d_1467_cf36_})
            d_1472_v83_: D12
            d_1472_v83_ = D12_DC21(d_1471_v82_)
            d_1473_v84_: D2
            d_1473_v84_ = D2_DC7(self.f15, (self).fm7(default__.fm0(d_1466_cf37_, globalState), globalState))
            index261_ = default__.safeIndex(574, (d_1468_v79_).length(0))
            rhs275_ = p0
            rhs276_ = (0) - (self.f15)
            rhs277_ = ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_1391_v25_, (d_1391_v25_).set(default__.safeIndex(746, len(d_1391_v25_)), (self).f5)])), self.f15, 597])) | (d_1470_v81_)) | (_dafny.MultiSet([self.f15, -663, 804]))
            rhs278_ = (d_1473_v84_).cf7
            rhs279_ = d_1472_v83_
            lhs188_ = d_1468_v79_
            lhs189_ = default__.safeIndex(574, (d_1468_v79_).length(0))
            lhs190_ = globalState
            d_1467_cf36_ = rhs275_
            lhs188_[lhs189_] = rhs276_
            d_1470_v81_ = rhs277_
            lhs190_.f1 = rhs278_
            d_1472_v83_ = rhs279_
            d_1474_v85_: C3
            nw241_ = C3()
            nw241_.ctor__(d_1466_cf37_, p2)
            d_1474_v85_ = nw241_
            d_1475_v86_: D11
            d_1475_v86_ = D11_DC20((self).f6, (d_1468_v79_)[default__.safeIndex(574, (d_1468_v79_).length(0))])
            d_1476_v87_: T4
            nw242_ = C2()
            nw242_.ctor__((self).f5, (d_1475_v86_).cf21)
            d_1476_v87_ = nw242_
            d_1476_v87_ = d_1476_v87_
            d_1477_v88_: _dafny.Array
            nw243_ = _dafny.Array(False, 26)
            d_1477_v88_ = nw243_
            index262_ = default__.safeIndex(764, (d_1477_v88_).length(0))
            (d_1477_v88_)[index262_] = d_1467_cf36_
            d_1478_v89_: _dafny.Array
            nw244_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_1478_v89_ = nw244_
            d_1479_v90_: _dafny.Array
            nw245_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_1479_v90_ = nw245_
            index263_ = default__.safeIndex(771, (d_1478_v89_).length(0))
            (d_1478_v89_)[index263_] = d_1479_v90_
            index264_ = default__.safeIndex(764, (d_1477_v88_).length(0))
            index265_ = default__.safeIndex(771, (d_1478_v89_).length(0))
            rhs280_ = not(d_1467_cf36_)
            rhs281_ = d_1479_v90_
            lhs191_ = d_1477_v88_
            lhs192_ = default__.safeIndex(764, (d_1477_v88_).length(0))
            lhs193_ = d_1478_v89_
            lhs194_ = default__.safeIndex(771, (d_1478_v89_).length(0))
            lhs191_[lhs192_] = rhs280_
            lhs193_[lhs194_] = rhs281_
        elif source21_.is_DC27:
            d_1480___mcc_h9_ = source21_.cf35
            d_1481_cf35_ = d_1480___mcc_h9_
            d_1482_v91_: _dafny.Array
            nw246_ = _dafny.Array(False, 27)
            d_1482_v91_ = nw246_
            d_1483_v92_: _dafny.Array
            d_1483_v92_ = d_1482_v91_
            d_1484_v93_: _dafny.Seq
            d_1484_v93_ = _dafny.SeqWithoutIsStrInference([d_1483_v92_])
            d_1484_v93_ = (d_1484_v93_) + ((d_1484_v93_) + (d_1484_v93_))
            d_1485_v94_: C1
            nw247_ = C1()
            nw247_.ctor__()
            d_1485_v94_ = nw247_
            d_1486_v95_: D11
            d_1486_v95_ = D11_DC20(p0, self.f15)
            d_1487_v96_: _dafny.Seq
            d_1487_v96_ = _dafny.SeqWithoutIsStrInference([self.f15])
            d_1488_v97_: _dafny.MultiSet
            d_1488_v97_ = _dafny.MultiSet([(self).fm7(p0, globalState), self.f15])
            d_1489_v98_: _dafny.Array
            nw248_ = _dafny.Array(None, 21)
            nw248_[int(0)] = self.f15
            nw248_[int(1)] = self.f15
            nw248_[int(2)] = self.f15
            nw248_[int(3)] = (d_1486_v95_).cf22
            nw248_[int(4)] = (d_1487_v96_)[default__.safeIndex(self.f15, len(d_1487_v96_))]
            nw248_[int(5)] = 632
            nw248_[int(6)] = (0) - (self.f15)
            nw248_[int(7)] = 725
            nw248_[int(8)] = 223
            nw248_[int(9)] = 898
            nw248_[int(10)] = self.f15
            nw248_[int(11)] = self.f15
            nw248_[int(12)] = default__.safeDivisionInt(self.f15, self.f15)
            nw248_[int(13)] = self.f15
            nw248_[int(14)] = (d_1485_v94_).fm14(globalState)
            nw248_[int(15)] = (((d_1488_v97_)[(d_1485_v94_).fm4(d_1387_v21_, 832, self.f15, d_1488_v97_, globalState)] if ((d_1485_v94_).fm4(d_1387_v21_, 832, self.f15, d_1488_v97_, globalState)) in (d_1488_v97_) else (self).fm7(p2, globalState))) - (self.f15)
            nw248_[int(16)] = len((_dafny.SeqWithoutIsStrInference([(self).f5 for d_1490_i5_ in range(default__.abs(-5))])) + (d_1391_v25_))
            nw248_[int(17)] = self.f15
            nw248_[int(18)] = (self.f15) * (self.f15)
            nw248_[int(19)] = self.f15
            nw248_[int(20)] = (0) - ((d_1487_v96_)[default__.safeIndex(437, len(d_1487_v96_))])
            d_1489_v98_ = nw248_
            index266_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            (d_1489_v98_)[index266_] = -44
            d_1491_v99_: bool
            d_1491_v99_ = True
            index267_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            rhs282_ = ((self).fm8((self).f6, p2, p2, globalState)) - (-696)
            rhs283_ = (d_1485_v94_).fm7(p2, globalState)
            rhs284_ = (default__.fm3(p2, (self).f6, globalState)) != (self.f15)
            rhs285_ = (0) - ((d_1487_v96_)[default__.safeIndex(len(((d_1487_v96_) + (d_1487_v96_)).set(default__.safeIndex(self.f15, len((d_1487_v96_) + (d_1487_v96_))), self.f15)), len(d_1487_v96_))])
            lhs195_ = globalState
            lhs196_ = d_1489_v98_
            lhs197_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            lhs198_ = globalState
            lhs195_.f1 = rhs282_
            lhs196_[lhs197_] = rhs283_
            d_1491_v99_ = rhs284_
            lhs198_.f1 = rhs285_
            d_1492_v100_: _dafny.MultiSet
            d_1492_v100_ = _dafny.MultiSet([d_1482_v91_, d_1482_v91_, d_1482_v91_])
            index268_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            index269_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            rhs286_ = ((d_1492_v100_)[(d_1482_v91_ if (self).f6 else d_1482_v91_)] if ((d_1482_v91_ if (self).f6 else d_1482_v91_)) in (d_1492_v100_) else self.f15)
            rhs287_ = (d_1489_v98_)[default__.safeIndex(697, (d_1489_v98_).length(0))]
            lhs199_ = d_1489_v98_
            lhs200_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            lhs201_ = d_1489_v98_
            lhs202_ = default__.safeIndex(697, (d_1489_v98_).length(0))
            lhs199_[lhs200_] = rhs286_
            lhs201_[lhs202_] = rhs287_
        elif True:
            d_1493___mcc_h10_ = source21_.cf38
            d_1494_cf38_ = d_1493___mcc_h10_
            d_1495_v101_: bool
            d_1495_v101_ = True
            d_1496_v102_: _dafny.Seq
            d_1496_v102_ = _dafny.SeqWithoutIsStrInference([self.f15, self.f15])
            rhs288_ = (self).f6
            rhs289_ = self.f15
            rhs290_ = (len(d_1496_v102_)) * (default__.safeDivisionInt(self.f15, self.f15))
            rhs291_ = 148
            lhs203_ = globalState
            lhs204_ = globalState
            lhs205_ = globalState
            d_1495_v101_ = rhs288_
            lhs203_.f1 = rhs289_
            lhs204_.f1 = rhs290_
            lhs205_.f1 = rhs291_
            d_1497_v103_: _dafny.Set
            d_1497_v103_ = _dafny.Set({p2})
            d_1495_v101_ = (_dafny.Set({p0, d_1495_v101_})).ispropersubset(d_1497_v103_)
            d_1495_v101_ = (d_1387_v21_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(0) - (self.f15)])), len(d_1387_v21_))]
            d_1498_v104_: _dafny.MultiSet
            d_1498_v104_ = _dafny.MultiSet([self.f15, self.f15])
            d_1499_v105_: _dafny.Array
            nw249_ = _dafny.Array(None, 20)
            nw249_[int(0)] = d_1498_v104_
            nw249_[int(1)] = d_1498_v104_
            nw249_[int(2)] = d_1498_v104_
            nw249_[int(3)] = (d_1498_v104_).intersection(d_1498_v104_)
            nw249_[int(4)] = d_1498_v104_
            nw249_[int(5)] = _dafny.MultiSet((d_1496_v102_) + (_dafny.SeqWithoutIsStrInference([self.f15 for d_1500_i6_ in range(default__.abs(782))])))
            nw249_[int(6)] = d_1498_v104_
            nw249_[int(7)] = d_1498_v104_
            nw249_[int(8)] = (default__.fm45(globalState)) | (d_1498_v104_)
            nw249_[int(9)] = d_1498_v104_
            nw249_[int(10)] = d_1498_v104_
            nw249_[int(11)] = d_1498_v104_
            nw249_[int(12)] = d_1498_v104_
            nw249_[int(13)] = d_1498_v104_
            nw249_[int(14)] = d_1498_v104_
            nw249_[int(15)] = (_dafny.MultiSet([self.f15, self.f15])) | (d_1498_v104_)
            nw249_[int(16)] = (default__.fm45(globalState)).intersection(d_1498_v104_)
            nw249_[int(17)] = d_1498_v104_
            nw249_[int(18)] = d_1498_v104_
            nw249_[int(19)] = d_1498_v104_
            d_1499_v105_ = nw249_
            rhs292_ = False
            rhs293_ = -677
            rhs294_ = d_1499_v105_
            rhs295_ = d_1495_v101_
            lhs206_ = globalState
            d_1495_v101_ = rhs292_
            lhs206_.f1 = rhs293_
            d_1499_v105_ = rhs294_
            d_1495_v101_ = rhs295_
        d_1501_v106_: bool
        d_1501_v106_ = False
        d_1502_v107_: _dafny.Array
        nw250_ = _dafny.Array(int(0), 1)
        d_1502_v107_ = nw250_
        d_1503_v108_: _dafny.Map
        d_1503_v108_ = _dafny.Map({d_1502_v107_: (self.f15) + (self.f15)})
        d_1504_v109_: _dafny.Seq
        d_1504_v109_ = _dafny.SeqWithoutIsStrInference([self.f15])
        d_1505_v110_: D14
        d_1505_v110_ = D14_DC26((self).fm6(d_1504_v109_, globalState), self.f15, (self).f5)
        rhs296_ = (d_1505_v110_).cf33
        rhs297_ = ((d_1391_v25_) + (d_1391_v25_)) != (((((d_1391_v25_).set(default__.safeIndex(self.f15, len(d_1391_v25_)), (self).f5)) + (d_1391_v25_)).set(default__.safeIndex(self.f15, len(((d_1391_v25_).set(default__.safeIndex(self.f15, len(d_1391_v25_)), (self).f5)) + (d_1391_v25_))), (self).f5)).set(default__.safeIndex(self.f15, len((((d_1391_v25_).set(default__.safeIndex(self.f15, len(d_1391_v25_)), (self).f5)) + (d_1391_v25_)).set(default__.safeIndex(self.f15, len(((d_1391_v25_).set(default__.safeIndex(self.f15, len(d_1391_v25_)), (self).f5)) + (d_1391_v25_))), (self).f5))), (self).f5))
        rhs298_ = (d_1503_v108_) | (d_1503_v108_)
        lhs207_ = self
        lhs207_.f15 = rhs296_
        d_1501_v106_ = rhs297_
        d_1503_v108_ = rhs298_
        d_1506_v111_: D6
        d_1506_v111_ = D6_DC13(d_1504_v109_)
        pat_let_tv19_ = d_1504_v109_
        def iife87_(_pat_let18_0):
            def iife88_(d_1507_dt__update__tmp_h1_):
                def iife89_(_pat_let19_0):
                    def iife90_(d_1508_dt__update_hcf14_h0_):
                        return D6_DC13(d_1508_dt__update_hcf14_h0_)
                    return iife90_(_pat_let19_0)
                return iife89_(pat_let_tv19_)
            return iife88_(_pat_let18_0)
        d_1506_v111_ = iife87_(d_1506_v111_)

    def m12(self, p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_1509_v0_: _dafny.Array
        nw251_ = _dafny.Array(int(0), 28)
        d_1509_v0_ = nw251_
        d_1510_v1_: _dafny.Seq
        d_1510_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_1511_v2_: _dafny.Map
        d_1511_v2_ = _dafny.Map({True: len(d_1510_v1_)})
        d_1512_v3_: _dafny.Map
        d_1512_v3_ = _dafny.Map({self.f15: (self).f6})
        d_1513_v4_: D2
        d_1513_v4_ = D2_DC7(self.f15, len(d_1512_v3_))
        d_1514_v5_: _dafny.Seq
        d_1514_v5_ = _dafny.SeqWithoutIsStrInference([self.f15])
        d_1515_v6_: _dafny.MultiSet
        d_1515_v6_ = _dafny.MultiSet([(self).f6])
        d_1516_v7_: _dafny.Map
        d_1516_v7_ = _dafny.Map({(self).f5: p1})
        d_1517_v8_: _dafny.Array
        nw252_ = _dafny.Array(None, 24)
        nw252_[int(0)] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_1509_v0_, d_1509_v0_])))
        nw252_[int(1)] = p2
        nw252_[int(2)] = default__.safeDivisionInt(self.f15, p2)
        nw252_[int(3)] = (((d_1511_v2_)[(self).f6] if ((self).f6) in (d_1511_v2_) else (d_1513_v4_).cf8) if p1 else p2)
        nw252_[int(4)] = default__.fm3((self).f6, p1, globalState)
        nw252_[int(5)] = -81
        nw252_[int(6)] = p2
        nw252_[int(7)] = (d_1514_v5_)[default__.safeIndex(len(_dafny.Map({p2: p1})), len(d_1514_v5_))]
        nw252_[int(8)] = p2
        nw252_[int(9)] = self.f15
        nw252_[int(10)] = (((d_1515_v6_)[(self).f6] if ((self).f6) in (d_1515_v6_) else self.f15)) - (p2)
        nw252_[int(11)] = (d_1514_v5_)[default__.safeIndex(-258, len(d_1514_v5_))]
        nw252_[int(12)] = p2
        nw252_[int(13)] = (self.f15) * (len(d_1516_v7_))
        nw252_[int(14)] = ((d_1515_v6_)[False] if (False) in (d_1515_v6_) else self.f15)
        nw252_[int(15)] = self.f15
        nw252_[int(16)] = (p0 if p1 else p0)
        nw252_[int(17)] = 886
        nw252_[int(18)] = len(default__.fm1(p0, (self).f6, -623, globalState))
        nw252_[int(19)] = (p2) - (p2)
        nw252_[int(20)] = len(d_1511_v2_)
        nw252_[int(21)] = self.f15
        nw252_[int(22)] = (0) - (p0)
        nw252_[int(23)] = p2
        d_1517_v8_ = nw252_
        index270_ = default__.safeIndex(892, (d_1517_v8_).length(0))
        (d_1517_v8_)[index270_] = default__.safeModuloInt(p0, self.f15)
        index271_ = default__.safeIndex(892, (d_1517_v8_).length(0))
        (d_1517_v8_)[index271_] = (len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1518_i0_ in range(default__.abs(364))]))) - (p2)
        d_1519_i1_: int
        d_1519_i1_ = 0
        with _dafny.label("14"):
            while p1:
                with _dafny.c_label("14"):
                    if (d_1519_i1_) >= (100):
                        raise _dafny.Break("14")
                    d_1519_i1_ = (d_1519_i1_) + (1)
                    d_1520_v9_: _dafny.Seq
                    d_1520_v9_ = _dafny.SeqWithoutIsStrInference([d_1514_v5_])
                    d_1521_v10_: C0
                    nw253_ = C0()
                    nw253_.ctor__((self).f6, d_1520_v9_)
                    d_1521_v10_ = nw253_
                    r3 = True
                    r3 = (self).f6
                    d_1522_v11_: _dafny.Array
                    nw254_ = _dafny.Array(False, 24)
                    d_1522_v11_ = nw254_
                    index272_ = default__.safeIndex(272, (d_1522_v11_).length(0))
                    (d_1522_v11_)[index272_] = p1
                    d_1523_v13_: _dafny.Map
                    d_1523_v13_ = _dafny.Map({723: d_1510_v1_})
                    index273_ = default__.safeIndex(272, (d_1522_v11_).length(0))
                    def iife91_():
                        coll51_ = _dafny.Map()
                        compr_51_: int
                        for compr_51_ in (default__.fm56(globalState)).Elements:
                            d_1524_v12_: int = compr_51_
                            if (d_1524_v12_) in (default__.fm56(globalState)):
                                coll51_[default__.safeModuloInt(d_1524_v12_, p0)] = d_1510_v1_
                        return _dafny.Map(coll51_)
                    (d_1522_v11_)[index273_] = (iife91_()
                    ) == (d_1523_v13_)
                    pass
            pass
        index274_ = default__.safeIndex(892, (d_1517_v8_).length(0))
        def iife92_():
            coll52_ = _dafny.Set()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(804, 781):
                d_1525_v14_: int = compr_52_
                if ((804) <= (d_1525_v14_)) and ((d_1525_v14_) < (781)):
                    coll52_ = coll52_.union(_dafny.Set([(d_1525_v14_) * ((0) - (p2))]))
            return _dafny.Set(coll52_)
        (d_1517_v8_)[index274_] = default__.safeModuloInt(self.f15, default__.safeModuloInt(len(iife92_()
        ), p0))
        d_1526_v15_: _dafny.Array
        def lambda75_(d_1527_i2_):
            return (self).f6

        init44_ = lambda75_
        nw255_ = _dafny.Array(None, 25)
        for i0_44_ in range(nw255_.length(0)):
            nw255_[i0_44_] = init44_(i0_44_)
        d_1526_v15_ = nw255_
        d_1528_v16_: _dafny.Array
        nw256_ = _dafny.Array(None, 17)
        nw256_[int(0)] = (self).f5
        nw256_[int(1)] = (self).f5
        nw256_[int(2)] = (self).f5
        nw256_[int(3)] = _dafny.CodePoint('u')
        nw256_[int(4)] = (self).f5
        nw256_[int(5)] = (self).f5
        nw256_[int(6)] = (self).f5
        nw256_[int(7)] = (self).f5
        nw256_[int(8)] = (self).f5
        nw256_[int(9)] = (self).f5
        nw256_[int(10)] = (self).f5
        nw256_[int(11)] = (self).f5
        nw256_[int(12)] = (self).f5
        nw256_[int(13)] = (self).f5
        nw256_[int(14)] = (self).f5
        nw256_[int(15)] = (self).f5
        def iife93_(_pat_let20_0):
            def iife94_(d_1529_dt__update__tmp_h0_):
                def iife95_(_pat_let21_0):
                    def iife96_(d_1530_dt__update_hcf2_h0_):
                        return D0_DC2((d_1529_dt__update__tmp_h0_).cf1, d_1530_dt__update_hcf2_h0_)
                    return iife96_(_pat_let21_0)
                return iife95_((self).f5)
            return iife94_(_pat_let20_0)
        nw256_[int(16)] = (iife93_(D0_DC2(_dafny.Map({d_1526_v15_: (self).f5}), (self).f5))).cf2
        d_1528_v16_ = nw256_
        index275_ = default__.safeIndex(366, (d_1528_v16_).length(0))
        (d_1528_v16_)[index275_] = (self).f5
        index276_ = default__.safeIndex(366, (d_1528_v16_).length(0))
        (d_1528_v16_)[index276_] = (self).f5
        d_1531_i3_: int
        d_1531_i3_ = 0
        with _dafny.label("15"):
            while (self).f6:
                with _dafny.c_label("15"):
                    if (d_1531_i3_) >= (100):
                        raise _dafny.Break("15")
                    d_1531_i3_ = (d_1531_i3_) + (1)
                    index277_ = default__.safeIndex(366, (d_1528_v16_).length(0))
                    (d_1528_v16_)[index277_] = _dafny.CodePoint('p')
                    index278_ = default__.safeIndex(519, (d_1526_v15_).length(0))
                    (d_1526_v15_)[index278_] = ((self).f6 if True else p1)
                    index279_ = default__.safeIndex(519, (d_1526_v15_).length(0))
                    (d_1526_v15_)[index279_] = (self).f6
                    r3 = (d_1526_v15_)[default__.safeIndex(519, (d_1526_v15_).length(0))]
                    d_1532_v17_: _dafny.Set
                    d_1532_v17_ = _dafny.Set({(d_1517_v8_)[default__.safeIndex(892, (d_1517_v8_).length(0))], self.f15})
                    d_1533_v18_: _dafny.Set
                    d_1533_v18_ = _dafny.Set({(0) - ((len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1534_i4_ in range(default__.abs(857))]))) * (len(d_1532_v17_))), p0})
                    d_1532_v17_ = d_1533_v18_
                    pass
            pass
        d_1535_v19_: _dafny.Array
        def lambda76_(d_1536_v8_):
            def lambda77_(d_1537_i5_):
                return (_dafny.Set({(d_1536_v8_)[default__.safeIndex(892, (d_1536_v8_).length(0))]})).intersection(_dafny.Set({self.f15, len(_dafny.Map({(self).f6: (self).f6}))}))

            return lambda77_

        init45_ = lambda76_(d_1517_v8_)
        nw257_ = _dafny.Array(None, 23)
        for i0_45_ in range(nw257_.length(0)):
            nw257_[i0_45_] = init45_(i0_45_)
        d_1535_v19_ = nw257_
        index280_ = default__.safeIndex(892, (d_1517_v8_).length(0))
        rhs299_ = d_1535_v19_
        rhs300_ = p1
        rhs301_ = (self).f6
        rhs302_ = self.f15
        rhs303_ = len((_dafny.SeqWithoutIsStrInference([(d_1528_v16_)[default__.safeIndex(366, (d_1528_v16_).length(0))] for d_1538_i6_ in range(default__.abs(-24))])).set(default__.safeIndex((d_1517_v8_)[default__.safeIndex(892, (d_1517_v8_).length(0))], len(_dafny.SeqWithoutIsStrInference([(d_1528_v16_)[default__.safeIndex(366, (d_1528_v16_).length(0))] for d_1539_i6_ in range(default__.abs(-24))]))), ((self).f5 if p1 else (self).f5)))
        lhs208_ = d_1517_v8_
        lhs209_ = default__.safeIndex(892, (d_1517_v8_).length(0))
        lhs210_ = self
        d_1535_v19_ = rhs299_
        r3 = rhs300_
        r3 = rhs301_
        lhs208_[lhs209_] = rhs302_
        lhs210_.f15 = rhs303_
        d_1540_v20_: _dafny.Map
        d_1540_v20_ = _dafny.Map({default__.fm1(len(d_1510_v1_), (self).fm6(d_1514_v5_, globalState), (d_1517_v8_)[default__.safeIndex(892, (d_1517_v8_).length(0))], globalState): p0})
        r0 = (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vudfnvin")): p0})) | (d_1540_v20_)
        r1 = ((d_1514_v5_) + (d_1514_v5_) if p1 else d_1514_v5_)
        d_1541_v21_: _dafny.Set
        d_1541_v21_ = _dafny.Set({(self).f6})
        r2 = (d_1541_v21_) - (default__.fm49(globalState))
        d_1542_v22_: _dafny.Seq
        d_1542_v22_ = _dafny.SeqWithoutIsStrInference([d_1509_v0_])
        d_1543_v23_: _dafny.Seq
        d_1543_v23_ = _dafny.SeqWithoutIsStrInference([d_1510_v1_])
        d_1544_v24_: _dafny.Map
        d_1544_v24_ = _dafny.Map({(default__.fm1(138, (self).f6, len(d_1542_v22_), globalState)) + ((d_1543_v23_)[default__.safeIndex((d_1517_v8_)[default__.safeIndex(892, (d_1517_v8_).length(0))], len(d_1543_v23_))]): (self).f6})
        r3 = ((d_1544_v24_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdv"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdv"))) in (d_1544_v24_) else False)
        return r0, r1, r2, r3


class C6(T2, T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, globalState):
        return not (False) or (False)

    def fm7(self, p0, globalState):
        def iife97_():
            coll53_ = _dafny.Map()
            compr_53_: str
            for compr_53_ in (_dafny.Set({_dafny.CodePoint('k')})).Elements:
                d_1545_v0_: str = compr_53_
                if (d_1545_v0_) in (_dafny.Set({_dafny.CodePoint('k')})):
                    coll53_[d_1545_v0_] = True
            return _dafny.Map(coll53_)
        return ((982) - (571)) * (default__.safeDivisionInt((0) - (len(iife97_()
        )), 708))

    def fm4(self, p0, p1, p2, p3, globalState):
        return (default__.safeDivisionInt(len(_dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyudstnyx"))), 832])).cardinality: False})), (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1546_i0_ in range(default__.abs(145))]))))) + (111)

    def fm5(self, p0, globalState):
        return D1_DC4(default__.safeDivisionInt(276, 330))

    def fm29(self, p0, p1, globalState):
        return ((0) - (((_dafny.MultiSet([len(_dafny.Map({True: True}))])).cardinality) * (568))) - (215)

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r0 = not (not(p0)) or (p0)
        d_1547_v0_: D0
        d_1547_v0_ = D0_DC1()
        d_1548_v1_: _dafny.Set
        d_1548_v1_ = _dafny.Set({d_1547_v0_, d_1547_v0_, d_1547_v0_})
        d_1549_v2_: _dafny.Set
        d_1549_v2_ = d_1548_v1_
        source22_ = d_1549_v2_
        d_1550___mcc_h0_ = source22_
        d_1551_cf17_ = d_1550___mcc_h0_
        (globalState).f1 = p1
        r0 = (p1) <= (p1)
        d_1552_v3_: C1
        nw258_ = C1()
        nw258_.ctor__()
        d_1552_v3_ = nw258_
        d_1553_v4_: _dafny.Seq
        d_1553_v4_ = _dafny.SeqWithoutIsStrInference([p1])
        d_1554_v5_: _dafny.Map
        d_1554_v5_ = _dafny.Map({p0: p0})
        d_1555_v6_: _dafny.Array
        nw259_ = _dafny.Array(None, 28)
        nw259_[int(0)] = p0
        nw259_[int(1)] = p0
        nw259_[int(2)] = p0
        nw259_[int(3)] = p0
        nw259_[int(4)] = p0
        nw259_[int(5)] = p0
        nw259_[int(6)] = False
        nw259_[int(7)] = True
        nw259_[int(8)] = p0
        nw259_[int(9)] = p0
        nw259_[int(10)] = p0
        nw259_[int(11)] = default__.fm0(False, globalState)
        nw259_[int(12)] = p0
        nw259_[int(13)] = p0
        nw259_[int(14)] = p0
        nw259_[int(15)] = (self).fm6(d_1553_v4_, globalState)
        nw259_[int(16)] = p0
        nw259_[int(17)] = ((d_1554_v5_)[p0] if (p0) in (d_1554_v5_) else p0)
        nw259_[int(18)] = (d_1552_v3_).fm6(d_1553_v4_, globalState)
        nw259_[int(19)] = p0
        nw259_[int(20)] = p0
        nw259_[int(21)] = p0
        nw259_[int(22)] = p0
        nw259_[int(23)] = p0
        nw259_[int(24)] = p0
        nw259_[int(25)] = p0
        nw259_[int(26)] = p0
        nw259_[int(27)] = p0
        d_1555_v6_ = nw259_
        d_1556_v7_: _dafny.Map
        d_1556_v7_ = _dafny.Map({d_1552_v3_: d_1555_v6_})
        d_1557_v8_: _dafny.Map
        d_1557_v8_ = d_1556_v7_
        source23_ = d_1557_v8_
        d_1558___mcc_h1_ = source23_
        d_1559_cf18_ = d_1558___mcc_h1_
        d_1560_v9_: T1
        nw260_ = C4()
        nw260_.ctor__()
        d_1560_v9_ = nw260_
        d_1561_v10_: _dafny.MultiSet
        d_1561_v10_ = _dafny.MultiSet([d_1560_v9_])
        d_1562_v11_: _dafny.Map
        d_1562_v11_ = _dafny.Map({(_dafny.Map({p0: p0})) | (d_1554_v5_): d_1561_v10_})
        d_1563_v12_: _dafny.Seq
        d_1563_v12_ = _dafny.SeqWithoutIsStrInference([d_1547_v0_, d_1547_v0_])
        d_1564_v13_: _dafny.Seq
        d_1564_v13_ = _dafny.SeqWithoutIsStrInference([d_1561_v10_])
        rhs304_ = not(p0)
        rhs305_ = (_dafny.Map({d_1554_v5_: d_1561_v10_})).set(d_1554_v5_, (d_1564_v13_)[default__.safeIndex(p1, len(d_1564_v13_))])
        rhs306_ = d_1563_v12_
        r0 = rhs304_
        d_1562_v11_ = rhs305_
        d_1563_v12_ = rhs306_
        d_1565_v14_: str
        d_1565_v14_ = _dafny.CodePoint('q')
        d_1566_v15_: C3
        nw261_ = C3()
        nw261_.ctor__(default__.fm0(p0, globalState), (d_1565_v14_) not in (p2))
        d_1566_v15_ = nw261_
        d_1567_v16_: _dafny.Seq
        d_1567_v16_ = _dafny.SeqWithoutIsStrInference([not((d_1566_v15_).f14)])
        d_1568_v17_: _dafny.Array
        nw262_ = _dafny.Array(int(0), 14)
        d_1568_v17_ = nw262_
        d_1569_v18_: _dafny.Map
        d_1569_v18_ = _dafny.Map({d_1567_v16_: d_1568_v17_})
        d_1570_v19_: _dafny.Map
        d_1570_v19_ = _dafny.Map({d_1569_v18_: (d_1566_v15_).f14})
        d_1570_v19_ = (d_1570_v19_).set(d_1569_v18_, p0)
        d_1571_v20_: _dafny.Map
        d_1571_v20_ = _dafny.Map({d_1567_v16_: (d_1566_v15_).f13})
        d_1572_v21_: _dafny.Array
        nw263_ = _dafny.Array(None, 15)
        nw263_[int(0)] = (d_1554_v5_) | (_dafny.Map({(d_1566_v15_).f14: (d_1566_v15_).f14}))
        nw263_[int(1)] = d_1554_v5_
        nw263_[int(2)] = (_dafny.Map({(d_1566_v15_).f14: p0})) | (_dafny.Map({not(True): not(p0)}))
        nw263_[int(3)] = (d_1554_v5_) | ((default__.fm41(p0, True, p0, globalState)).set((d_1566_v15_).f14, (d_1566_v15_).f14))
        nw263_[int(4)] = d_1554_v5_
        nw263_[int(5)] = d_1554_v5_
        nw263_[int(6)] = d_1554_v5_
        nw263_[int(7)] = (d_1554_v5_) | (d_1554_v5_)
        nw263_[int(8)] = (d_1554_v5_).set((d_1566_v15_).f13, False)
        nw263_[int(9)] = d_1554_v5_
        nw263_[int(10)] = _dafny.Map({((d_1571_v20_)[d_1567_v16_] if (d_1567_v16_) in (d_1571_v20_) else p0): (d_1566_v15_).f13})
        nw263_[int(11)] = d_1554_v5_
        nw263_[int(12)] = d_1554_v5_
        nw263_[int(13)] = (_dafny.Map({(d_1566_v15_).f13: (d_1566_v15_).f14})) | (_dafny.Map({(d_1566_v15_).f13: p0}))
        nw263_[int(14)] = d_1554_v5_
        d_1572_v21_ = nw263_
        index281_ = default__.safeIndex(573, (d_1572_v21_).length(0))
        (d_1572_v21_)[index281_] = d_1554_v5_
        d_1573_v22_: _dafny.Set
        d_1573_v22_ = _dafny.Set({p1, (0) - (p1)})
        index282_ = default__.safeIndex(573, (d_1572_v21_).length(0))
        (d_1572_v21_)[index282_] = _dafny.Map({(d_1566_v15_).f14: (d_1573_v22_).issubset(d_1573_v22_)})
        d_1574_v23_: str
        d_1574_v23_ = _dafny.CodePoint('a')
        d_1575_v24_: _dafny.Map
        d_1575_v24_ = _dafny.Map({p0: p1})
        d_1576_v25_: C4
        nw264_ = C4()
        nw264_.ctor__()
        d_1576_v25_ = nw264_
        d_1577_v26_: D14
        d_1577_v26_ = D14_DC25(d_1576_v25_)
        d_1578_v27_: _dafny.Map
        d_1578_v27_ = _dafny.Map({p0: (d_1577_v26_).cf31})
        d_1579_v28_: _dafny.Seq
        d_1579_v28_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1580_v29_: _dafny.MultiSet
        d_1580_v29_ = _dafny.MultiSet([d_1574_v23_])
        d_1581_v30_: _dafny.Array
        nw265_ = _dafny.Array(None, 10)
        nw265_[int(0)] = (0) - (len((p2) + ((p2).set(default__.safeIndex(p1, len(p2)), d_1574_v23_))))
        nw265_[int(1)] = default__.safeDivisionInt(339, p1)
        nw265_[int(2)] = p1
        nw265_[int(3)] = p1
        nw265_[int(4)] = (d_1553_v4_)[default__.safeIndex(((d_1575_v24_)[p0] if (p0) in (d_1575_v24_) else len(d_1578_v27_)), len(d_1553_v4_))]
        nw265_[int(5)] = p1
        nw265_[int(6)] = (0) - ((0) - (((_dafny.MultiSet([len(d_1554_v5_), p1, (0) - (len(d_1579_v28_))])).cardinality) * (p1)))
        nw265_[int(7)] = p1
        nw265_[int(8)] = (p1) * (p1)
        nw265_[int(9)] = ((d_1580_v29_)[d_1574_v23_] if (d_1574_v23_) in (d_1580_v29_) else p1)
        d_1581_v30_ = nw265_
        index283_ = default__.safeIndex(298, (d_1581_v30_).length(0))
        (d_1581_v30_)[index283_] = p1
        d_1582_v31_: _dafny.Set
        d_1582_v31_ = _dafny.Set({p1, (0) - (p1)})
        d_1583_v32_: _dafny.Map
        d_1583_v32_ = _dafny.Map({p1: len(d_1582_v31_)})
        d_1584_v33_: _dafny.Map
        d_1584_v33_ = _dafny.Map({d_1583_v32_: p0})
        d_1585_v34_: _dafny.Seq
        d_1585_v34_ = _dafny.SeqWithoutIsStrInference([d_1579_v28_])
        d_1586_v35_: _dafny.Map
        d_1586_v35_ = _dafny.Map({p0: d_1579_v28_})
        index284_ = default__.safeIndex(298, (d_1581_v30_).length(0))
        (d_1581_v30_)[index284_] = len(default__.fm1(((d_1576_v25_).fm30(d_1584_v33_, _dafny.MultiSet(d_1585_v34_), globalState)) * (p1), p0, len(((d_1586_v35_)[p0] if (p0) in (d_1586_v35_) else d_1579_v28_)), globalState))
        d_1587_v36_: _dafny.Array
        nw266_ = _dafny.Array(_dafny.MultiSet({}), 20)
        d_1587_v36_ = nw266_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1587_v36_).length(0)):
            d_1588_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_1588_i0_)) and ((d_1588_i0_) < ((d_1587_v36_).length(0)))):
                (d_1587_v36_)[(d_1588_i0_)] = ((_dafny.MultiSet([p0, p0])).intersection(_dafny.MultiSet([p0]))) | ((_dafny.MultiSet([p0])).intersection(_dafny.MultiSet([p0, p0, p0])))
        hi9_ = len(p2)
        for d_1589_i1_ in range(p1, hi9_):
            r0 = p0
            d_1590_v37_: _dafny.Array
            def lambda78_(d_1591_p2_):
                def lambda79_(d_1592_i2_):
                    return (_dafny.CodePoint('p')) not in (d_1591_p2_)

                return lambda79_

            init46_ = lambda78_(p2)
            nw267_ = _dafny.Array(None, 7)
            for i0_46_ in range(nw267_.length(0)):
                nw267_[i0_46_] = init46_(i0_46_)
            d_1590_v37_ = nw267_
            d_1593_v38_: _dafny.Seq
            d_1593_v38_ = _dafny.SeqWithoutIsStrInference([False, p0, p0, True])
            d_1594_v39_: _dafny.Seq
            d_1594_v39_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({p0: p0}))])
            d_1595_v40_: _dafny.Set
            d_1595_v40_ = _dafny.Set({d_1594_v39_})
            d_1596_v41_: _dafny.Array
            nw268_ = _dafny.Array(None, 5)
            nw268_[int(0)] = p1
            nw268_[int(1)] = ((_dafny.MultiSet([d_1593_v38_])).cardinality if p0 else p1)
            nw268_[int(2)] = (d_1589_i1_) * (d_1589_i1_)
            nw268_[int(3)] = len((d_1593_v38_) + ((d_1593_v38_).set(default__.safeIndex(p1, len(d_1593_v38_)), p0)))
            nw268_[int(4)] = default__.safeModuloInt(len(d_1595_v40_), (0) - ((0) - (-942)))
            d_1596_v41_ = nw268_
            index285_ = default__.safeIndex(214, (d_1596_v41_).length(0))
            (d_1596_v41_)[index285_] = p1
            index286_ = default__.safeIndex(214, (d_1596_v41_).length(0))
            rhs307_ = d_1590_v37_
            rhs308_ = p1
            lhs211_ = d_1596_v41_
            lhs212_ = default__.safeIndex(214, (d_1596_v41_).length(0))
            d_1590_v37_ = rhs307_
            lhs211_[lhs212_] = rhs308_
            d_1597_v42_: _dafny.Map
            d_1597_v42_ = _dafny.Map({d_1589_i1_: len(d_1594_v39_)})
            r0 = not(not ((p1) > (default__.fm3(True, p0, globalState))) or ((((d_1597_v42_)[d_1589_i1_] if (d_1589_i1_) in (d_1597_v42_) else p1)) < ((d_1596_v41_)[default__.safeIndex(214, (d_1596_v41_).length(0))])))
            d_1598_v43_: str
            d_1598_v43_ = _dafny.CodePoint('f')
            d_1599_v44_: _dafny.Map
            d_1599_v44_ = _dafny.Map({len(p2): not(p0)})
            d_1600_v45_: _dafny.Map
            d_1600_v45_ = _dafny.Map({d_1589_i1_: ((d_1599_v44_)[(d_1596_v41_)[default__.safeIndex(214, (d_1596_v41_).length(0))]] if ((d_1596_v41_)[default__.safeIndex(214, (d_1596_v41_).length(0))]) in (d_1599_v44_) else p0)})
            d_1601_v46_: D0
            d_1602_v47_: bool
            out44_: D0
            out45_: bool
            out44_, out45_ = default__.m0(d_1590_v37_, d_1598_v43_, p0, default__.fm36(d_1589_i1_, p1, len(d_1600_v45_), d_1597_v42_, globalState), globalState)
            d_1601_v46_ = out44_
            d_1602_v47_ = out45_
        d_1603_v48_: D2
        d_1603_v48_ = D2_DC6(p2)
        d_1604_v49_: _dafny.Map
        d_1604_v49_ = _dafny.Map({p0: p0})
        d_1605_v50_: _dafny.Array
        nw269_ = _dafny.Array(None, 22)
        nw269_[int(0)] = p1
        nw269_[int(1)] = p1
        nw269_[int(2)] = p1
        nw269_[int(3)] = len(p2)
        nw269_[int(4)] = p1
        nw269_[int(5)] = (0) - (p1)
        nw269_[int(6)] = p1
        nw269_[int(7)] = p1
        nw269_[int(8)] = p1
        nw269_[int(9)] = p1
        nw269_[int(10)] = p1
        nw269_[int(11)] = p1
        nw269_[int(12)] = p1
        nw269_[int(13)] = len(d_1604_v49_)
        nw269_[int(14)] = p1
        nw269_[int(15)] = -618
        nw269_[int(16)] = p1
        nw269_[int(17)] = p1
        nw269_[int(18)] = p1
        nw269_[int(19)] = p1
        nw269_[int(20)] = p1
        nw269_[int(21)] = p1
        d_1605_v50_ = nw269_
        d_1606_v51_: _dafny.MultiSet
        d_1606_v51_ = _dafny.MultiSet([d_1605_v50_, d_1605_v50_])
        d_1607_v52_: _dafny.Seq
        d_1607_v52_ = _dafny.SeqWithoutIsStrInference([((d_1606_v51_)[d_1605_v50_] if (d_1605_v50_) in (d_1606_v51_) else p1)])
        hi10_ = (d_1607_v52_)[default__.safeIndex(p1, len(d_1607_v52_))]
        for d_1608_i3_ in range(default__.safeDivisionInt((self).fm29(p1, (d_1603_v48_).cf6, globalState), p1), hi10_):
            d_1609_v53_: str
            d_1609_v53_ = _dafny.CodePoint('k')
            d_1610_v54_: _dafny.Set
            d_1610_v54_ = _dafny.Set({d_1609_v53_, d_1609_v53_})
            d_1611_v55_: _dafny.Array
            nw270_ = _dafny.Array(None, 5)
            d_1611_v55_ = nw270_
            d_1612_v56_: C3
            nw271_ = C3()
            nw271_.ctor__(p0, p0)
            d_1612_v56_ = nw271_
            index287_ = default__.safeIndex(28, (d_1611_v55_).length(0))
            (d_1611_v55_)[index287_] = (d_1612_v56_ if p0 else d_1612_v56_)
            index288_ = default__.safeIndex(171, (d_1605_v50_).length(0))
            (d_1605_v50_)[index288_] = p1
            d_1613_v57_: _dafny.Array
            nw272_ = _dafny.Array(None, 13)
            nw272_[int(0)] = d_1609_v53_
            nw272_[int(1)] = d_1609_v53_
            nw272_[int(2)] = d_1609_v53_
            nw272_[int(3)] = d_1609_v53_
            nw272_[int(4)] = d_1609_v53_
            nw272_[int(5)] = d_1609_v53_
            nw272_[int(6)] = (d_1609_v53_ if p0 else d_1609_v53_)
            nw272_[int(7)] = _dafny.CodePoint('w')
            nw272_[int(8)] = d_1609_v53_
            nw272_[int(9)] = (d_1609_v53_ if (d_1612_v56_).f13 else d_1609_v53_)
            nw272_[int(10)] = d_1609_v53_
            nw272_[int(11)] = d_1609_v53_
            nw272_[int(12)] = d_1609_v53_
            d_1613_v57_ = nw272_
            index289_ = default__.safeIndex(977, (d_1613_v57_).length(0))
            (d_1613_v57_)[index289_] = d_1609_v53_
            d_1614_v58_: _dafny.Array
            def lambda80_(d_1615_i4_):
                return False

            init47_ = lambda80_
            nw273_ = _dafny.Array(None, 8)
            for i0_47_ in range(nw273_.length(0)):
                nw273_[i0_47_] = init47_(i0_47_)
            d_1614_v58_ = nw273_
            d_1616_v59_: _dafny.Map
            d_1616_v59_ = _dafny.Map({p0: d_1614_v58_})
            index290_ = default__.safeIndex(28, (d_1611_v55_).length(0))
            index291_ = default__.safeIndex(171, (d_1605_v50_).length(0))
            index292_ = default__.safeIndex(977, (d_1613_v57_).length(0))
            rhs309_ = d_1610_v54_
            rhs310_ = d_1612_v56_
            rhs311_ = d_1608_i3_
            rhs312_ = d_1609_v53_
            rhs313_ = d_1616_v59_
            lhs213_ = d_1611_v55_
            lhs214_ = default__.safeIndex(28, (d_1611_v55_).length(0))
            lhs215_ = d_1605_v50_
            lhs216_ = default__.safeIndex(171, (d_1605_v50_).length(0))
            lhs217_ = d_1613_v57_
            lhs218_ = default__.safeIndex(977, (d_1613_v57_).length(0))
            d_1610_v54_ = rhs309_
            lhs213_[lhs214_] = rhs310_
            lhs215_[lhs216_] = rhs311_
            lhs217_[lhs218_] = rhs312_
            d_1616_v59_ = rhs313_
            d_1617_v60_: C2
            nw274_ = C2()
            nw274_.ctor__(d_1609_v53_, False)
            d_1617_v60_ = nw274_
            index293_ = default__.safeIndex(171, (d_1605_v50_).length(0))
            (d_1605_v50_)[index293_] = default__.safeDivisionInt(d_1608_i3_, default__.safeModuloInt(p1, len(p2)))
            r0 = False
        d_1618_v61_: _dafny.Array
        nw275_ = _dafny.Array(_dafny.MultiSet({}), 14)
        d_1618_v61_ = nw275_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1618_v61_).length(0)):
            d_1619_i5_: int = guard_loop_2_
            if (True) and (((0) <= (d_1619_i5_)) and ((d_1619_i5_) < ((d_1618_v61_).length(0)))):
                (d_1618_v61_)[(d_1619_i5_)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p0, p0]))]))) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p1 for d_1620_i6_ in range(default__.abs(20))]))) if p0 else (_dafny.MultiSet([p1])).intersection(_dafny.MultiSet([p1, len(d_1607_v52_)])))
        r0 = p0
        d_1621_v62_: _dafny.Map
        d_1621_v62_ = _dafny.Map({691: p0})
        d_1622_v63_: str
        d_1622_v63_ = _dafny.CodePoint('x')
        r1 = (((p2) + (p2)) + ((p2) + (p2))).set(default__.safeIndex(len((d_1621_v62_).set(p1, p0)), len(((p2) + (p2)) + ((p2) + (p2)))), d_1622_v63_)
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        d_1623_v0_: _dafny.Seq
        d_1623_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojxcdwwr"))
        d_1623_v0_ = d_1623_v0_
        d_1624_v1_: bool
        d_1624_v1_ = True
        d_1625_i0_: int
        d_1625_i0_ = 0
        with _dafny.label("16"):
            while d_1624_v1_:
                with _dafny.c_label("16"):
                    if (d_1625_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_1625_i0_ = (d_1625_i0_) + (1)
                    d_1626_v2_: _dafny.Seq
                    d_1626_v2_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_])
                    d_1627_v3_: int
                    d_1627_v3_ = 715
                    d_1628_v4_: D11
                    d_1628_v4_ = D11_DC20(d_1624_v1_, (0) - (d_1627_v3_))
                    if (d_1626_v2_)[default__.safeIndex((d_1627_v3_ if (d_1628_v4_).cf21 else d_1627_v3_), len(d_1626_v2_))]:
                        d_1629_v5_: C4
                        nw276_ = C4()
                        nw276_.ctor__()
                        d_1629_v5_ = nw276_
                        d_1630_v6_: str
                        d_1630_v6_ = _dafny.CodePoint('o')
                        d_1630_v6_ = d_1630_v6_
                        d_1631_v7_: _dafny.Array
                        nw277_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                        d_1631_v7_ = nw277_
                        index294_ = default__.safeIndex(915, (d_1631_v7_).length(0))
                        (d_1631_v7_)[index294_] = (d_1623_v0_) + (d_1623_v0_)
                        index295_ = default__.safeIndex(915, (d_1631_v7_).length(0))
                        (d_1631_v7_)[index295_] = ((_dafny.SeqWithoutIsStrInference([d_1630_v6_ for d_1632_i1_ in range(default__.abs(-456))])) + (_dafny.SeqWithoutIsStrInference([d_1630_v6_ for d_1633_i2_ in range(default__.abs(-219))])) if not(not(d_1624_v1_)) else d_1623_v0_)
                        d_1634_v8_: _dafny.Array
                        nw278_ = _dafny.Array(_dafny.Map({}), 9)
                        d_1634_v8_ = nw278_
                        d_1635_v9_: _dafny.Map
                        d_1635_v9_ = _dafny.Map({d_1624_v1_: (0) - (d_1627_v3_)})
                        index296_ = default__.safeIndex(108, (d_1634_v8_).length(0))
                        (d_1634_v8_)[index296_] = d_1635_v9_
                        index297_ = default__.safeIndex(108, (d_1634_v8_).length(0))
                        rhs314_ = d_1627_v3_
                        rhs315_ = not(((0) - ((d_1627_v3_) * (d_1627_v3_))) != (d_1627_v3_))
                        rhs316_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "syeqbdag")))) + ((d_1623_v0_) + ((d_1631_v7_)[default__.safeIndex(915, (d_1631_v7_).length(0))]))
                        rhs317_ = d_1624_v1_
                        rhs318_ = (d_1635_v9_) | ((_dafny.Map({d_1624_v1_: d_1627_v3_})) | (d_1635_v9_))
                        lhs219_ = d_1634_v8_
                        lhs220_ = default__.safeIndex(108, (d_1634_v8_).length(0))
                        d_1627_v3_ = rhs314_
                        d_1624_v1_ = rhs315_
                        d_1623_v0_ = rhs316_
                        d_1624_v1_ = rhs317_
                        lhs219_[lhs220_] = rhs318_
                        d_1629_v5_ = d_1629_v5_
                    elif True:
                        d_1636_v10_: _dafny.Array
                        nw279_ = _dafny.Array(_dafny.CodePoint('D'), 7)
                        d_1636_v10_ = nw279_
                        d_1637_v11_: _dafny.Array
                        nw280_ = _dafny.Array(None, 13)
                        nw280_[int(0)] = d_1636_v10_
                        nw280_[int(1)] = d_1636_v10_
                        nw280_[int(2)] = d_1636_v10_
                        nw280_[int(3)] = d_1636_v10_
                        nw280_[int(4)] = d_1636_v10_
                        nw280_[int(5)] = d_1636_v10_
                        nw280_[int(6)] = d_1636_v10_
                        nw280_[int(7)] = d_1636_v10_
                        nw280_[int(8)] = d_1636_v10_
                        nw280_[int(9)] = d_1636_v10_
                        nw280_[int(10)] = d_1636_v10_
                        nw280_[int(11)] = d_1636_v10_
                        nw280_[int(12)] = d_1636_v10_
                        d_1637_v11_ = nw280_
                        d_1638_v12_: _dafny.Array
                        nw281_ = _dafny.Array(int(0), 4)
                        d_1638_v12_ = nw281_
                        d_1639_v13_: _dafny.Map
                        d_1639_v13_ = _dafny.Map({d_1637_v11_: d_1638_v12_})
                        d_1639_v13_ = (d_1639_v13_).set(d_1637_v11_, d_1638_v12_)
                        d_1640_v14_: str
                        d_1640_v14_ = _dafny.CodePoint('r')
                        d_1641_v15_: D11
                        d_1641_v15_ = D11_DC19(d_1640_v14_)
                        d_1641_v15_ = d_1641_v15_
                        d_1642_v16_: C4
                        nw282_ = C4()
                        nw282_.ctor__()
                        d_1642_v16_ = nw282_
                        d_1643_v17_: _dafny.Set
                        d_1643_v17_ = _dafny.Set({d_1624_v1_, d_1624_v1_})
                        r0 = default__.safeDivisionInt(d_1627_v3_, default__.safeDivisionInt(len(d_1643_v17_), d_1627_v3_))
                        index298_ = default__.safeIndex(251, (d_1638_v12_).length(0))
                        (d_1638_v12_)[index298_] = d_1627_v3_
                        d_1644_v18_: _dafny.Map
                        d_1644_v18_ = _dafny.Map({d_1627_v3_: d_1624_v1_})
                        d_1645_v19_: _dafny.Map
                        d_1645_v19_ = d_1644_v18_
                        d_1646_v20_: _dafny.Array
                        nw283_ = _dafny.Array(_dafny.Seq({}), 11)
                        d_1646_v20_ = nw283_
                        index299_ = default__.safeIndex(290, (d_1646_v20_).length(0))
                        (d_1646_v20_)[index299_] = d_1626_v2_
                        d_1647_v21_: _dafny.Seq
                        d_1647_v21_ = _dafny.SeqWithoutIsStrInference([298, d_1627_v3_, d_1627_v3_])
                        d_1648_v23_: _dafny.Seq
                        d_1648_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhq")), d_1623_v0_])
                        index300_ = default__.safeIndex(251, (d_1638_v12_).length(0))
                        index301_ = default__.safeIndex(290, (d_1646_v20_).length(0))
                        def iife98_():
                            def iife100_():
                                coll56_ = _dafny.Set()
                                compr_56_: _dafny.Seq
                                for compr_56_ in (d_1648_v23_).Elements:
                                    d_1651_v24_: _dafny.Seq = compr_56_
                                    if (d_1651_v24_) in (d_1648_v23_):
                                        coll56_ = coll56_.union(_dafny.Set([d_1651_v24_]))
                                return _dafny.Set(coll56_)
                            coll54_ = _dafny.Map()
                            def iife99_():
                                coll55_ = _dafny.Set()
                                compr_55_: _dafny.Seq
                                for compr_55_ in (d_1648_v23_).Elements:
                                    d_1649_v24_: _dafny.Seq = compr_55_
                                    if (d_1649_v24_) in (d_1648_v23_):
                                        coll55_ = coll55_.union(_dafny.Set([d_1649_v24_]))
                                return _dafny.Set(coll55_)
                            compr_54_: _dafny.Seq
                            for compr_54_ in (iife99_()
                            ).Elements:
                                d_1650_v22_: _dafny.Seq = compr_54_
                                if (d_1650_v22_) in (iife100_()
                                ):
                                    coll54_[d_1650_v22_] = d_1624_v1_
                            return _dafny.Map(coll54_)
                        def iife101_():
                            def iife103_():
                                coll59_ = _dafny.Set()
                                compr_59_: _dafny.Seq
                                for compr_59_ in (d_1648_v23_).Elements:
                                    d_1654_v24_: _dafny.Seq = compr_59_
                                    if (d_1654_v24_) in (d_1648_v23_):
                                        coll59_ = coll59_.union(_dafny.Set([d_1654_v24_]))
                                return _dafny.Set(coll59_)
                            coll57_ = _dafny.Map()
                            def iife102_():
                                coll58_ = _dafny.Set()
                                compr_58_: _dafny.Seq
                                for compr_58_ in (d_1648_v23_).Elements:
                                    d_1652_v24_: _dafny.Seq = compr_58_
                                    if (d_1652_v24_) in (d_1648_v23_):
                                        coll58_ = coll58_.union(_dafny.Set([d_1652_v24_]))
                                return _dafny.Set(coll58_)
                            compr_57_: _dafny.Seq
                            for compr_57_ in (iife102_()
                            ).Elements:
                                d_1653_v22_: _dafny.Seq = compr_57_
                                if (d_1653_v22_) in (iife103_()
                                ):
                                    coll57_[d_1653_v22_] = d_1624_v1_
                            return _dafny.Map(coll57_)
                        rhs319_ = (d_1627_v3_ if not((d_1624_v1_ if d_1624_v1_ else False)) else len(d_1647_v21_))
                        rhs320_ = d_1645_v19_
                        rhs321_ = ((((d_1626_v2_).set(default__.safeIndex(len(d_1644_v18_), len(d_1626_v2_)), False)).set(default__.safeIndex(len(iife98_()
                        ), len((d_1626_v2_).set(default__.safeIndex(len(d_1644_v18_), len(d_1626_v2_)), False))), d_1624_v1_)).set(default__.safeIndex(len(d_1626_v2_), len(((d_1626_v2_).set(default__.safeIndex(len(d_1644_v18_), len(d_1626_v2_)), False)).set(default__.safeIndex(len(iife101_()
                        ), len((d_1626_v2_).set(default__.safeIndex(len(d_1644_v18_), len(d_1626_v2_)), False))), d_1624_v1_))), d_1624_v1_)) + (d_1626_v2_)
                        lhs221_ = d_1638_v12_
                        lhs222_ = default__.safeIndex(251, (d_1638_v12_).length(0))
                        lhs223_ = d_1646_v20_
                        lhs224_ = default__.safeIndex(290, (d_1646_v20_).length(0))
                        lhs221_[lhs222_] = rhs319_
                        d_1645_v19_ = rhs320_
                        lhs223_[lhs224_] = rhs321_
                    d_1655_v25_: C4
                    nw284_ = C4()
                    nw284_.ctor__()
                    d_1655_v25_ = nw284_
                    d_1656_v26_: D14
                    d_1656_v26_ = D14_DC25(d_1655_v25_)
                    pat_let_tv20_ = d_1655_v25_
                    d_1657_v27_: _dafny.Array
                    nw285_ = _dafny.Array(None, 13)
                    nw285_[int(0)] = d_1656_v26_
                    nw285_[int(1)] = d_1656_v26_
                    nw285_[int(2)] = d_1656_v26_
                    nw285_[int(3)] = d_1656_v26_
                    nw285_[int(4)] = d_1656_v26_
                    nw285_[int(5)] = d_1656_v26_
                    nw285_[int(6)] = D14_DC25(d_1655_v25_)
                    nw285_[int(7)] = d_1656_v26_
                    nw285_[int(8)] = D14_DC25(d_1655_v25_)
                    nw285_[int(9)] = D14_DC25(d_1655_v25_)
                    nw285_[int(10)] = d_1656_v26_
                    nw285_[int(11)] = d_1656_v26_
                    def iife104_(_pat_let22_0):
                        def iife105_(d_1658_dt__update__tmp_h0_):
                            def iife106_(_pat_let23_0):
                                def iife107_(d_1659_dt__update_hcf31_h0_):
                                    return D14_DC25(d_1659_dt__update_hcf31_h0_)
                                return iife107_(_pat_let23_0)
                            return iife106_(pat_let_tv20_)
                        return iife105_(_pat_let22_0)
                    nw285_[int(12)] = iife104_(d_1656_v26_)
                    d_1657_v27_ = nw285_
                    pat_let_tv21_ = d_1655_v25_
                    index302_ = default__.safeIndex(105, (d_1657_v27_).length(0))
                    def iife108_(_pat_let24_0):
                        def iife109_(d_1660_dt__update__tmp_h1_):
                            def iife110_(_pat_let25_0):
                                def iife111_(d_1661_dt__update_hcf31_h1_):
                                    return D14_DC25(d_1661_dt__update_hcf31_h1_)
                                return iife111_(_pat_let25_0)
                            return iife110_(pat_let_tv21_)
                        return iife109_(_pat_let24_0)
                    (d_1657_v27_)[index302_] = iife108_(d_1656_v26_)
                    d_1662_v28_: _dafny.Array
                    nw286_ = _dafny.Array(_dafny.Seq({}), 8)
                    d_1662_v28_ = nw286_
                    index303_ = default__.safeIndex(618, (d_1662_v28_).length(0))
                    (d_1662_v28_)[index303_] = _dafny.SeqWithoutIsStrInference([d_1627_v3_])
                    d_1663_v29_: _dafny.Array
                    def lambda81_(d_1664_v2_, d_1665_v3_):
                        def lambda82_(d_1666_i3_):
                            return _dafny.Map({(_dafny.MultiSet(d_1664_v2_)).cardinality: d_1665_v3_})

                        return lambda82_

                    init48_ = lambda81_(d_1626_v2_, d_1627_v3_)
                    nw287_ = _dafny.Array(None, 10)
                    for i0_48_ in range(nw287_.length(0)):
                        nw287_[i0_48_] = init48_(i0_48_)
                    d_1663_v29_ = nw287_
                    d_1667_v30_: str
                    d_1667_v30_ = _dafny.CodePoint('b')
                    d_1668_v31_: D11
                    d_1668_v31_ = D11_DC19(d_1667_v30_)
                    d_1669_v32_: _dafny.Map
                    d_1669_v32_ = _dafny.Map({(_dafny.Map({not(d_1624_v1_): (d_1668_v31_).cf20})).set(d_1624_v1_, d_1667_v30_): 88})
                    d_1670_v33_: _dafny.Array
                    nw288_ = _dafny.Array(False, 6)
                    d_1670_v33_ = nw288_
                    d_1671_v34_: _dafny.Array
                    d_1671_v34_ = d_1670_v33_
                    d_1672_v35_: _dafny.Map
                    d_1672_v35_ = _dafny.Map({d_1671_v34_: d_1624_v1_})
                    d_1673_v36_: _dafny.Seq
                    d_1673_v36_ = _dafny.SeqWithoutIsStrInference([len(d_1669_v32_), len((d_1672_v35_ if default__.fm0(default__.fm0(not(d_1624_v1_), globalState), globalState) else d_1672_v35_)), d_1627_v3_, d_1627_v3_])
                    index304_ = default__.safeIndex(105, (d_1657_v27_).length(0))
                    index305_ = default__.safeIndex(618, (d_1662_v28_).length(0))
                    rhs322_ = d_1656_v26_
                    rhs323_ = d_1673_v36_
                    rhs324_ = not (d_1624_v1_) or (d_1624_v1_)
                    rhs325_ = 342
                    rhs326_ = d_1663_v29_
                    lhs225_ = d_1657_v27_
                    lhs226_ = default__.safeIndex(105, (d_1657_v27_).length(0))
                    lhs227_ = d_1662_v28_
                    lhs228_ = default__.safeIndex(618, (d_1662_v28_).length(0))
                    lhs229_ = globalState
                    lhs225_[lhs226_] = rhs322_
                    lhs227_[lhs228_] = rhs323_
                    d_1624_v1_ = rhs324_
                    lhs229_.f1 = rhs325_
                    d_1663_v29_ = rhs326_
                    d_1674_v37_: _dafny.Map
                    d_1674_v37_ = _dafny.Map({d_1624_v1_: d_1623_v0_})
                    d_1675_v38_: _dafny.MultiSet
                    d_1675_v38_ = _dafny.MultiSet([((d_1674_v37_)[d_1624_v1_] if (d_1624_v1_) in (d_1674_v37_) else d_1623_v0_)])
                    d_1676_v39_: _dafny.Map
                    d_1676_v39_ = _dafny.Map({True: d_1627_v3_})
                    d_1677_v40_: _dafny.Map
                    d_1677_v40_ = _dafny.Map({d_1624_v1_: d_1676_v39_})
                    d_1678_v42_: _dafny.MultiSet
                    d_1678_v42_ = _dafny.MultiSet([d_1624_v1_, d_1624_v1_, d_1624_v1_, d_1624_v1_])
                    d_1679_v43_: _dafny.MultiSet
                    d_1679_v43_ = _dafny.MultiSet([d_1627_v3_])
                    d_1680_v44_: _dafny.Array
                    nw289_ = _dafny.Array(None, 17)
                    nw289_[int(0)] = (-450 if d_1624_v1_ else d_1627_v3_)
                    nw289_[int(1)] = (0) - (len(d_1677_v40_))
                    nw289_[int(2)] = (d_1627_v3_) + (len(d_1623_v0_))
                    nw289_[int(3)] = len(default__.fm35(d_1623_v0_, _dafny.MultiSet((d_1662_v28_)[default__.safeIndex(618, (d_1662_v28_).length(0))]), d_1667_v30_, globalState))
                    nw289_[int(4)] = len(d_1623_v0_)
                    def iife112_():
                        coll60_ = _dafny.Set()
                        compr_60_: int
                        for compr_60_ in _dafny.IntegerRange(973, 8):
                            d_1681_v41_: int = compr_60_
                            if ((973) <= (d_1681_v41_)) and ((d_1681_v41_) < (8)):
                                coll60_ = coll60_.union(_dafny.Set([(d_1681_v41_) - (d_1627_v3_)]))
                        return _dafny.Set(coll60_)
                    nw289_[int(5)] = len(iife112_()
                    )
                    nw289_[int(6)] = ((0) - (len(d_1623_v0_)) if d_1624_v1_ else d_1627_v3_)
                    nw289_[int(7)] = d_1627_v3_
                    nw289_[int(8)] = d_1627_v3_
                    nw289_[int(9)] = 927
                    nw289_[int(10)] = d_1627_v3_
                    nw289_[int(11)] = default__.fm3(d_1624_v1_, d_1624_v1_, globalState)
                    nw289_[int(12)] = d_1627_v3_
                    nw289_[int(13)] = d_1627_v3_
                    nw289_[int(14)] = len((d_1623_v0_) + (_dafny.SeqWithoutIsStrInference([d_1667_v30_ for d_1682_i4_ in range(default__.abs(131))])))
                    nw289_[int(15)] = (d_1678_v42_).cardinality
                    nw289_[int(16)] = (d_1679_v43_).cardinality
                    d_1680_v44_ = nw289_
                    d_1683_v45_: _dafny.Array
                    d_1683_v45_ = d_1680_v44_
                    d_1684_v46_: D12
                    d_1684_v46_ = D12_DC23(d_1680_v44_)
                    d_1685_v47_: _dafny.Seq
                    d_1685_v47_ = _dafny.SeqWithoutIsStrInference([d_1684_v46_, d_1684_v46_])
                    d_1686_v48_: _dafny.Map
                    d_1686_v48_ = _dafny.Map({(264) * (len(d_1623_v0_)): (d_1685_v47_).set(default__.safeIndex(len(d_1673_v36_), len(d_1685_v47_)), d_1684_v46_)})
                    d_1687_v49_: _dafny.Seq
                    d_1687_v49_ = _dafny.SeqWithoutIsStrInference([d_1623_v0_, ((d_1623_v0_).set(default__.safeIndex(d_1627_v3_, len(d_1623_v0_)), d_1667_v30_)).set(default__.safeIndex(d_1627_v3_, len((d_1623_v0_).set(default__.safeIndex(d_1627_v3_, len(d_1623_v0_)), d_1667_v30_))), d_1667_v30_)])
                    d_1688_v50_: _dafny.Map
                    d_1688_v50_ = _dafny.Map({292: _dafny.Map({d_1627_v3_: d_1685_v47_})})
                    rhs327_ = (d_1627_v3_) * (d_1627_v3_)
                    rhs328_ = _dafny.MultiSet((d_1687_v49_) + (d_1687_v49_))
                    rhs329_ = ((0) - (default__.safeModuloInt(-808, d_1627_v3_))) - ((d_1627_v3_) - (393))
                    rhs330_ = (d_1683_v45_)
                    rhs331_ = (d_1686_v48_) | (((d_1688_v50_)[d_1627_v3_] if (d_1627_v3_) in (d_1688_v50_) else d_1686_v48_))
                    lhs230_ = globalState
                    d_1627_v3_ = rhs327_
                    d_1675_v38_ = rhs328_
                    lhs230_.f1 = rhs329_
                    d_1680_v44_ = rhs330_
                    d_1686_v48_ = rhs331_
                    (globalState).f1 = d_1627_v3_
                    pass
            pass
        if d_1624_v1_:
            d_1689_v53_: _dafny.Array
            def lambda83_(d_1690_v0_, d_1691_v1_):
                def lambda84_(d_1692_i5_):
                    def iife113_():
                        def iife115_():
                            coll63_ = _dafny.Set()
                            compr_63_: _dafny.Seq
                            for compr_63_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1693_i6_ in range(default__.abs(-862))]), d_1690_v0_])).Elements:
                                d_1696_v52_: _dafny.Seq = compr_63_
                                if (d_1696_v52_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1693_i6_ in range(default__.abs(-862))]), d_1690_v0_])):
                                    coll63_ = coll63_.union(_dafny.Set([d_1696_v52_]))
                            return _dafny.Set(coll63_)
                        coll61_ = _dafny.Map()
                        def iife114_():
                            coll62_ = _dafny.Set()
                            compr_62_: _dafny.Seq
                            for compr_62_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1693_i6_ in range(default__.abs(-862))]), d_1690_v0_])).Elements:
                                d_1694_v52_: _dafny.Seq = compr_62_
                                if (d_1694_v52_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1693_i6_ in range(default__.abs(-862))]), d_1690_v0_])):
                                    coll62_ = coll62_.union(_dafny.Set([d_1694_v52_]))
                            return _dafny.Set(coll62_)
                        compr_61_: _dafny.Seq
                        for compr_61_ in (iife114_()
                        ).Elements:
                            d_1695_v51_: _dafny.Seq = compr_61_
                            if (d_1695_v51_) in (iife115_()
                            ):
                                coll61_[d_1695_v51_] = _dafny.Map({not(not(d_1691_v1_)): 989})
                        return _dafny.Map(coll61_)
                    return (d_1690_v0_)[default__.safeIndex(len(_dafny.Set({len(iife113_()
                    )})), len(d_1690_v0_))]

                return lambda84_

            init49_ = lambda83_(d_1623_v0_, d_1624_v1_)
            nw290_ = _dafny.Array(None, 23)
            for i0_49_ in range(nw290_.length(0)):
                nw290_[i0_49_] = init49_(i0_49_)
            d_1689_v53_ = nw290_
            d_1689_v53_ = d_1689_v53_
            if not(d_1624_v1_):
                d_1697_v54_: int
                d_1697_v54_ = 487
                (globalState).f1 = (d_1697_v54_) * (d_1697_v54_)
                d_1698_v55_: _dafny.MultiSet
                d_1698_v55_ = _dafny.MultiSet([not(d_1624_v1_)])
                d_1697_v54_ = ((d_1698_v55_)[d_1624_v1_] if (d_1624_v1_) in (d_1698_v55_) else 94)
                d_1699_v56_: _dafny.Array
                def lambda85_(d_1700_v1_):
                    def lambda86_(d_1701_i7_):
                        return d_1700_v1_

                    return lambda86_

                init50_ = lambda85_(d_1624_v1_)
                nw291_ = _dafny.Array(None, 17)
                for i0_50_ in range(nw291_.length(0)):
                    nw291_[i0_50_] = init50_(i0_50_)
                d_1699_v56_ = nw291_
                index306_ = default__.safeIndex(116, (d_1699_v56_).length(0))
                (d_1699_v56_)[index306_] = d_1624_v1_
                d_1702_v57_: str
                d_1702_v57_ = _dafny.CodePoint('r')
                d_1703_v58_: _dafny.Map
                d_1703_v58_ = _dafny.Map({d_1699_v56_: _dafny.CodePoint('c')})
                d_1704_v59_: D0
                d_1704_v59_ = D0_DC2(d_1703_v58_, d_1702_v57_)
                d_1705_v60_: _dafny.Seq
                d_1705_v60_ = _dafny.SeqWithoutIsStrInference([d_1704_v59_, D0_DC2(d_1703_v58_, d_1702_v57_)])
                d_1706_v61_: _dafny.Seq
                d_1706_v61_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_, d_1624_v1_])
                d_1707_v62_: _dafny.Set
                d_1707_v62_ = _dafny.Set({d_1624_v1_})
                d_1708_v63_: _dafny.Seq
                d_1708_v63_ = _dafny.SeqWithoutIsStrInference([d_1697_v54_])
                d_1709_v64_: _dafny.MultiSet
                d_1709_v64_ = _dafny.MultiSet([len(d_1708_v63_)])
                d_1710_v65_: _dafny.Map
                d_1710_v65_ = _dafny.Map({d_1697_v54_: (self).fm4(((d_1706_v61_).set(default__.safeIndex(len(d_1707_v62_), len(d_1706_v61_)), d_1624_v1_)).set(default__.safeIndex(d_1697_v54_, len((d_1706_v61_).set(default__.safeIndex(len(d_1707_v62_), len(d_1706_v61_)), d_1624_v1_))), not(d_1624_v1_)), d_1697_v54_, (0) - (d_1697_v54_), d_1709_v64_, globalState)})
                index307_ = default__.safeIndex(116, (d_1699_v56_).length(0))
                rhs332_ = d_1624_v1_
                rhs333_ = (_dafny.SeqWithoutIsStrInference([d_1704_v59_])) < (d_1705_v60_)
                rhs334_ = default__.fm36(d_1697_v54_, d_1697_v54_, (d_1697_v54_) - (d_1697_v54_), d_1710_v65_, globalState)
                lhs231_ = d_1699_v56_
                lhs232_ = default__.safeIndex(116, (d_1699_v56_).length(0))
                lhs231_[lhs232_] = rhs332_
                d_1624_v1_ = rhs333_
                d_1702_v57_ = rhs334_
                d_1624_v1_ = d_1624_v1_
                d_1697_v54_ = d_1697_v54_
            elif True:
                d_1711_v66_: _dafny.Seq
                d_1711_v66_ = _dafny.SeqWithoutIsStrInference([default__.fm3(d_1624_v1_, d_1624_v1_, globalState)])
                d_1712_v67_: int
                d_1712_v67_ = -506
                d_1713_v68_: _dafny.Array
                nw292_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_1713_v68_ = nw292_
                d_1714_v69_: _dafny.Map
                d_1714_v69_ = _dafny.Map({(d_1711_v66_)[default__.safeIndex(d_1712_v67_, len(d_1711_v66_))]: d_1713_v68_})
                d_1714_v69_ = (d_1714_v69_).set(d_1712_v67_, d_1713_v68_)
                d_1715_v70_: _dafny.Array
                def lambda87_(d_1716_v67_):
                    def lambda88_(d_1717_i8_):
                        return (_dafny.MultiSet([d_1716_v67_, d_1716_v67_])) - (_dafny.MultiSet([d_1716_v67_, 459]))

                    return lambda88_

                init51_ = lambda87_(d_1712_v67_)
                nw293_ = _dafny.Array(None, 16)
                for i0_51_ in range(nw293_.length(0)):
                    nw293_[i0_51_] = init51_(i0_51_)
                d_1715_v70_ = nw293_
                d_1718_v71_: _dafny.MultiSet
                d_1718_v71_ = _dafny.MultiSet([d_1712_v67_])
                index308_ = default__.safeIndex(556, (d_1715_v70_).length(0))
                (d_1715_v70_)[index308_] = (d_1718_v71_).set(d_1712_v67_, default__.abs(d_1712_v67_))
                index309_ = default__.safeIndex(556, (d_1715_v70_).length(0))
                (d_1715_v70_)[index309_] = d_1718_v71_
                d_1719_v72_: str
                d_1719_v72_ = _dafny.CodePoint('k')
                d_1720_v73_: T3
                nw294_ = C5()
                nw294_.ctor__((0) - (d_1712_v67_), d_1719_v72_, d_1624_v1_)
                d_1720_v73_ = nw294_
                d_1721_v74_: _dafny.Array
                nw295_ = _dafny.Array(None, 2)
                nw295_[int(0)] = d_1720_v73_
                nw295_[int(1)] = d_1720_v73_
                d_1721_v74_ = nw295_
                index310_ = default__.safeIndex(422, (d_1721_v74_).length(0))
                (d_1721_v74_)[index310_] = d_1720_v73_
                index311_ = default__.safeIndex(422, (d_1721_v74_).length(0))
                (d_1721_v74_)[index311_] = d_1720_v73_
                d_1712_v67_ = (0) - ((-624) * (default__.safeDivisionInt(d_1712_v67_, d_1712_v67_)))
                d_1624_v1_ = (d_1720_v73_).f6
            if d_1624_v1_:
                d_1722_v75_: D2
                d_1722_v75_ = D2_DC6(d_1623_v0_)
                pat_let_tv22_ = d_1623_v0_
                def iife116_(_pat_let26_0):
                    def iife117_(d_1723_dt__update__tmp_h3_):
                        def iife118_(_pat_let27_0):
                            def iife119_(d_1724_dt__update_hcf6_h0_):
                                return D2_DC6(d_1724_dt__update_hcf6_h0_)
                            return iife119_(_pat_let27_0)
                        return iife118_(pat_let_tv22_)
                    return iife117_(_pat_let26_0)
                d_1722_v75_ = ((d_1722_v75_ if d_1624_v1_ else iife116_(d_1722_v75_)) if d_1624_v1_ else d_1722_v75_)
                d_1624_v1_ = d_1624_v1_
                d_1725_v76_: _dafny.Array
                nw296_ = _dafny.Array(None, 8)
                nw296_[int(0)] = d_1624_v1_
                nw296_[int(1)] = d_1624_v1_
                nw296_[int(2)] = d_1624_v1_
                nw296_[int(3)] = d_1624_v1_
                nw296_[int(4)] = d_1624_v1_
                nw296_[int(5)] = d_1624_v1_
                nw296_[int(6)] = False
                nw296_[int(7)] = d_1624_v1_
                d_1725_v76_ = nw296_
                d_1726_v77_: _dafny.Array
                d_1726_v77_ = d_1725_v76_
                d_1727_v78_: _dafny.Map
                d_1727_v78_ = _dafny.Map({d_1726_v77_: d_1725_v76_})
                d_1728_v79_: _dafny.Map
                d_1728_v79_ = d_1727_v78_
                d_1729_v80_: _dafny.Array
                nw297_ = _dafny.Array(None, 13)
                nw297_[int(0)] = d_1728_v79_
                nw297_[int(1)] = d_1728_v79_
                nw297_[int(2)] = d_1727_v78_
                nw297_[int(3)] = d_1727_v78_
                nw297_[int(4)] = (d_1728_v79_ if d_1624_v1_ else d_1728_v79_)
                nw297_[int(5)] = d_1728_v79_
                nw297_[int(6)] = d_1728_v79_
                nw297_[int(7)] = d_1727_v78_
                nw297_[int(8)] = d_1728_v79_
                nw297_[int(9)] = d_1728_v79_
                nw297_[int(10)] = d_1728_v79_
                nw297_[int(11)] = d_1727_v78_
                nw297_[int(12)] = d_1727_v78_
                d_1729_v80_ = nw297_
                index312_ = default__.safeIndex(288, (d_1729_v80_).length(0))
                (d_1729_v80_)[index312_] = d_1727_v78_
                index313_ = default__.safeIndex(288, (d_1729_v80_).length(0))
                (d_1729_v80_)[index313_] = d_1728_v79_
                d_1730_v81_: int
                d_1730_v81_ = 958
                d_1731_v82_: _dafny.Map
                d_1731_v82_ = _dafny.Map({(self).fm29(d_1730_v81_, d_1623_v0_, globalState): d_1725_v76_})
                def iife120_():
                    coll64_ = _dafny.Set()
                    compr_64_: int
                    for compr_64_ in _dafny.IntegerRange(342, 17):
                        d_1732_v83_: int = compr_64_
                        if ((342) <= (d_1732_v83_)) and ((d_1732_v83_) < (17)):
                            coll64_ = coll64_.union(_dafny.Set([default__.safeDivisionInt(d_1732_v83_, (0) - (d_1730_v81_))]))
                    return _dafny.Set(coll64_)
                d_1731_v82_ = (d_1731_v82_).set((0) - (default__.safeModuloInt(len(iife120_()
                ), d_1730_v81_)), d_1725_v76_)
                d_1733_v84_: _dafny.Map
                d_1733_v84_ = _dafny.Map({d_1730_v81_: d_1624_v1_})
                index314_ = default__.safeIndex(328, (d_1725_v76_).length(0))
                (d_1725_v76_)[index314_] = ((d_1733_v84_)[d_1730_v81_] if (d_1730_v81_) in (d_1733_v84_) else d_1624_v1_)
                index315_ = default__.safeIndex(328, (d_1725_v76_).length(0))
                (d_1725_v76_)[index315_] = False
            elif True:
                d_1734_v85_: int
                d_1734_v85_ = 168
                d_1735_v86_: _dafny.Map
                d_1735_v86_ = _dafny.Map({268: d_1734_v85_})
                d_1736_v87_: _dafny.Map
                d_1736_v87_ = d_1735_v86_
                d_1737_v88_: _dafny.Seq
                d_1737_v88_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_])
                d_1738_v89_: _dafny.Seq
                d_1738_v89_ = _dafny.SeqWithoutIsStrInference([d_1737_v88_])
                d_1739_v90_: str
                d_1739_v90_ = _dafny.CodePoint('k')
                d_1740_v91_: D12
                d_1740_v91_ = D12_DC22(d_1736_v87_, d_1738_v89_, d_1739_v90_, d_1624_v1_, d_1623_v0_)
                d_1741_v92_: _dafny.MultiSet
                d_1741_v92_ = _dafny.MultiSet([d_1740_v91_])
                d_1742_v93_: _dafny.Seq
                d_1742_v93_ = _dafny.SeqWithoutIsStrInference([(d_1734_v85_) >= ((d_1741_v92_).cardinality)])
                d_1624_v1_ = (d_1742_v93_)[default__.safeIndex(755, len(d_1742_v93_))]
                d_1743_v94_: _dafny.Array
                nw298_ = _dafny.Array(int(0), 13)
                d_1743_v94_ = nw298_
                index316_ = default__.safeIndex(819, (d_1743_v94_).length(0))
                (d_1743_v94_)[index316_] = 824
                index317_ = default__.safeIndex(819, (d_1743_v94_).length(0))
                (d_1743_v94_)[index317_] = len(d_1623_v0_)
                d_1744_v95_: _dafny.MultiSet
                d_1744_v95_ = _dafny.MultiSet([d_1624_v1_])
                d_1745_v96_: _dafny.Map
                d_1745_v96_ = _dafny.Map({d_1624_v1_: d_1744_v95_})
                d_1746_v97_: C2
                nw299_ = C2()
                nw299_.ctor__(d_1739_v90_, d_1624_v1_)
                d_1746_v97_ = nw299_
                d_1747_v98_: _dafny.Map
                d_1747_v98_ = _dafny.Map({(((d_1745_v96_)[d_1624_v1_] if (d_1624_v1_) in (d_1745_v96_) else d_1744_v95_)).ispropersubset(d_1744_v95_): d_1746_v97_})
                d_1748_v99_: _dafny.Seq
                d_1748_v99_ = _dafny.SeqWithoutIsStrInference([d_1734_v85_])
                d_1747_v98_ = (d_1747_v98_).set(not(((d_1748_v99_)[default__.safeIndex((d_1743_v94_)[default__.safeIndex(819, (d_1743_v94_).length(0))], len(d_1748_v99_))]) >= (d_1734_v85_)), d_1746_v97_)
                d_1624_v1_ = (not(not (not(d_1624_v1_)) or (d_1624_v1_))) and (d_1624_v1_)
                d_1749_v100_: C2
                nw300_ = C2()
                nw300_.ctor__(d_1739_v90_, (d_1624_v1_) or (d_1624_v1_))
                d_1749_v100_ = nw300_
            d_1750_v101_: int
            d_1750_v101_ = 904
            d_1751_v102_: _dafny.Map
            d_1751_v102_ = _dafny.Map({(0) - (d_1750_v101_): d_1624_v1_})
            d_1752_v103_: _dafny.Seq
            d_1752_v103_ = _dafny.SeqWithoutIsStrInference([d_1750_v101_, len(d_1751_v102_)])
            d_1753_v104_: _dafny.Seq
            d_1753_v104_ = _dafny.SeqWithoutIsStrInference([d_1752_v103_])
            d_1754_v105_: C0
            nw301_ = C0()
            nw301_.ctor__(True, d_1753_v104_)
            d_1754_v105_ = nw301_
            d_1755_v106_: _dafny.MultiSet
            d_1755_v106_ = _dafny.MultiSet([d_1754_v105_, d_1754_v105_])
            rhs335_ = default__.safeDivisionInt(((d_1755_v106_)[d_1754_v105_] if (d_1754_v105_) in (d_1755_v106_) else (0) - (d_1750_v101_)), d_1750_v101_)
            rhs336_ = (default__.safeDivisionInt(d_1750_v101_, d_1750_v101_)) == (d_1750_v101_)
            lhs233_ = globalState
            lhs233_.f1 = rhs335_
            d_1624_v1_ = rhs336_
            d_1756_v107_: _dafny.Seq
            d_1756_v107_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_])
            d_1757_v108_: _dafny.MultiSet
            d_1757_v108_ = _dafny.MultiSet([d_1756_v107_])
            d_1758_v109_: _dafny.Seq
            d_1758_v109_ = _dafny.SeqWithoutIsStrInference([d_1756_v107_])
            d_1757_v108_ = _dafny.MultiSet((d_1758_v109_) + (d_1758_v109_))
        elif True:
            if d_1624_v1_:
                d_1759_v110_: _dafny.Array
                nw302_ = _dafny.Array(None, 2)
                nw302_[int(0)] = d_1624_v1_
                nw302_[int(1)] = d_1624_v1_
                d_1759_v110_ = nw302_
                d_1760_v111_: _dafny.Map
                d_1760_v111_ = _dafny.Map({d_1759_v110_: -376})
                d_1761_v112_: _dafny.Seq
                d_1761_v112_ = _dafny.SeqWithoutIsStrInference([d_1759_v110_])
                d_1762_v113_: int
                d_1762_v113_ = 688
                d_1763_v114_: C2
                nw303_ = C2()
                nw303_.ctor__(_dafny.CodePoint('l'), d_1624_v1_)
                d_1763_v114_ = nw303_
                d_1764_v115_: _dafny.MultiSet
                d_1764_v115_ = _dafny.MultiSet([d_1763_v114_, d_1763_v114_])
                d_1765_v116_: _dafny.Seq
                d_1765_v116_ = _dafny.SeqWithoutIsStrInference([d_1760_v111_, d_1760_v111_])
                d_1766_v117_: _dafny.Map
                d_1766_v117_ = _dafny.Map({d_1762_v113_: d_1759_v110_})
                d_1767_v118_: str
                d_1767_v118_ = _dafny.CodePoint('h')
                d_1768_v119_: _dafny.Seq
                d_1768_v119_ = _dafny.SeqWithoutIsStrInference([d_1762_v113_])
                d_1769_v120_: _dafny.Map
                d_1769_v120_ = _dafny.Map({False: d_1768_v119_})
                d_1770_v121_: _dafny.Set
                d_1770_v121_ = _dafny.Set({False})
                d_1771_v122_: _dafny.Array
                nw304_ = _dafny.Array(None, 22)
                nw304_[int(0)] = d_1762_v113_
                nw304_[int(1)] = d_1762_v113_
                nw304_[int(2)] = d_1762_v113_
                nw304_[int(3)] = d_1762_v113_
                nw304_[int(4)] = d_1762_v113_
                nw304_[int(5)] = len(d_1623_v0_)
                nw304_[int(6)] = d_1762_v113_
                nw304_[int(7)] = d_1762_v113_
                nw304_[int(8)] = d_1762_v113_
                nw304_[int(9)] = d_1762_v113_
                nw304_[int(10)] = d_1762_v113_
                nw304_[int(11)] = d_1762_v113_
                nw304_[int(12)] = d_1762_v113_
                nw304_[int(13)] = d_1762_v113_
                nw304_[int(14)] = d_1762_v113_
                nw304_[int(15)] = len(d_1769_v120_)
                nw304_[int(16)] = d_1762_v113_
                nw304_[int(17)] = 105
                nw304_[int(18)] = len(d_1770_v121_)
                nw304_[int(19)] = 153
                nw304_[int(20)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toyypgf")))
                nw304_[int(21)] = d_1762_v113_
                d_1771_v122_ = nw304_
                d_1772_v123_: _dafny.Map
                d_1772_v123_ = _dafny.Map({d_1767_v118_: d_1771_v122_})
                d_1773_v124_: _dafny.Array
                nw305_ = _dafny.Array(None, 24)
                nw305_[int(0)] = (d_1760_v111_) | (d_1760_v111_)
                nw305_[int(1)] = d_1760_v111_
                nw305_[int(2)] = d_1760_v111_
                nw305_[int(3)] = (_dafny.Map({(d_1761_v112_)[default__.safeIndex(d_1762_v113_, len(d_1761_v112_))]: d_1762_v113_})) | (_dafny.Map({d_1759_v110_: (d_1764_v115_).cardinality}))
                nw305_[int(4)] = d_1760_v111_
                nw305_[int(5)] = (d_1760_v111_).set(d_1759_v110_, 948)
                nw305_[int(6)] = d_1760_v111_
                nw305_[int(7)] = (d_1765_v116_)[default__.safeIndex(d_1762_v113_, len(d_1765_v116_))]
                nw305_[int(8)] = d_1760_v111_
                nw305_[int(9)] = d_1760_v111_
                nw305_[int(10)] = _dafny.Map({((d_1766_v117_)[len(d_1623_v0_)] if (len(d_1623_v0_)) in (d_1766_v117_) else d_1759_v110_): d_1762_v113_})
                nw305_[int(11)] = d_1760_v111_
                nw305_[int(12)] = _dafny.Map({d_1759_v110_: d_1762_v113_})
                nw305_[int(13)] = d_1760_v111_
                nw305_[int(14)] = d_1760_v111_
                nw305_[int(15)] = _dafny.Map({d_1759_v110_: len(d_1772_v123_)})
                nw305_[int(16)] = (d_1760_v111_) | (d_1760_v111_)
                nw305_[int(17)] = _dafny.Map({d_1759_v110_: d_1762_v113_})
                nw305_[int(18)] = (d_1760_v111_) | (d_1760_v111_)
                nw305_[int(19)] = d_1760_v111_
                nw305_[int(20)] = d_1760_v111_
                nw305_[int(21)] = (d_1760_v111_) | (d_1760_v111_)
                nw305_[int(22)] = d_1760_v111_
                nw305_[int(23)] = _dafny.Map({d_1759_v110_: d_1762_v113_})
                d_1773_v124_ = nw305_
                index318_ = default__.safeIndex(422, (d_1773_v124_).length(0))
                (d_1773_v124_)[index318_] = d_1760_v111_
                d_1774_v125_: _dafny.MultiSet
                d_1774_v125_ = _dafny.MultiSet([d_1762_v113_, d_1762_v113_])
                index319_ = default__.safeIndex(422, (d_1773_v124_).length(0))
                (d_1773_v124_)[index319_] = ((d_1760_v111_).set(d_1759_v110_, d_1762_v113_)).set(d_1759_v110_, default__.safeModuloInt(d_1762_v113_, (d_1774_v125_).cardinality))
                rhs337_ = d_1767_v118_
                rhs338_ = (default__.safeDivisionInt(811, d_1762_v113_) if (False if d_1624_v1_ else d_1624_v1_) else d_1762_v113_)
                lhs234_ = globalState
                d_1767_v118_ = rhs337_
                lhs234_.f1 = rhs338_
                d_1775_v126_: D6
                d_1775_v126_ = D6_DC13(d_1768_v119_)
                d_1775_v126_ = d_1775_v126_
                d_1776_v127_: _dafny.Map
                d_1776_v127_ = _dafny.Map({d_1624_v1_: d_1762_v113_})
                d_1777_v128_: _dafny.Map
                d_1777_v128_ = (_dafny.Map({d_1762_v113_: len(d_1776_v127_)})).set(d_1762_v113_, -464)
                d_1778_v129_: _dafny.Seq
                d_1778_v129_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_])
                d_1779_v130_: _dafny.Seq
                d_1779_v130_ = _dafny.SeqWithoutIsStrInference([d_1778_v129_])
                d_1780_v131_: D12
                d_1780_v131_ = D12_DC22(d_1777_v128_, d_1779_v130_, d_1767_v118_, d_1624_v1_, d_1623_v0_)
                d_1781_v132_: _dafny.Set
                d_1781_v132_ = _dafny.Set({d_1759_v110_})
                d_1782_v133_: C2
                nw306_ = C2()
                nw306_.ctor__((d_1767_v118_ if d_1624_v1_ else (d_1780_v131_).cf26), (d_1759_v110_) in (d_1781_v132_))
                d_1782_v133_ = nw306_
                d_1783_v134_: _dafny.Set
                d_1783_v134_ = _dafny.Set({D11_DC19(_dafny.CodePoint('l'))})
                index320_ = default__.safeIndex(447, (d_1759_v110_).length(0))
                (d_1759_v110_)[index320_] = (d_1778_v129_)[default__.safeIndex(d_1762_v113_, len(d_1778_v129_))]
                d_1784_v135_: _dafny.Map
                d_1784_v135_ = _dafny.Map({d_1762_v113_: (d_1783_v134_) | (d_1783_v134_)})
                d_1785_v136_: _dafny.Map
                d_1785_v136_ = _dafny.Map({137: d_1624_v1_})
                d_1786_v137_: _dafny.Map
                d_1786_v137_ = _dafny.Map({d_1785_v136_: d_1762_v113_})
                index321_ = default__.safeIndex(447, (d_1759_v110_).length(0))
                rhs339_ = d_1624_v1_
                rhs340_ = ((d_1784_v135_)[default__.safeDivisionInt(((d_1786_v137_)[d_1785_v136_] if (d_1785_v136_) in (d_1786_v137_) else (_dafny.MultiSet([d_1759_v110_])).cardinality), d_1762_v113_)] if (default__.safeDivisionInt(((d_1786_v137_)[d_1785_v136_] if (d_1785_v136_) in (d_1786_v137_) else (_dafny.MultiSet([d_1759_v110_])).cardinality), d_1762_v113_)) in (d_1784_v135_) else d_1783_v134_)
                rhs341_ = d_1624_v1_
                lhs235_ = d_1759_v110_
                lhs236_ = default__.safeIndex(447, (d_1759_v110_).length(0))
                d_1624_v1_ = rhs339_
                d_1783_v134_ = rhs340_
                lhs235_[lhs236_] = rhs341_
            elif True:
                d_1787_v138_: int
                d_1787_v138_ = -987
                (globalState).f1 = (0) - (d_1787_v138_)
                d_1788_v139_: _dafny.Seq
                d_1788_v139_ = _dafny.SeqWithoutIsStrInference([d_1787_v138_, d_1787_v138_])
                d_1624_v1_ = (self).fm6(d_1788_v139_, globalState)
                d_1789_v140_: _dafny.Array
                nw307_ = _dafny.Array(False, 12)
                d_1789_v140_ = nw307_
                d_1790_v141_: _dafny.Array
                nw308_ = _dafny.Array(None, 13)
                nw308_[int(0)] = d_1789_v140_
                nw308_[int(1)] = d_1789_v140_
                nw308_[int(2)] = d_1789_v140_
                nw308_[int(3)] = d_1789_v140_
                nw308_[int(4)] = d_1789_v140_
                nw308_[int(5)] = d_1789_v140_
                nw308_[int(6)] = d_1789_v140_
                nw308_[int(7)] = d_1789_v140_
                nw308_[int(8)] = d_1789_v140_
                nw308_[int(9)] = d_1789_v140_
                nw308_[int(10)] = d_1789_v140_
                nw308_[int(11)] = d_1789_v140_
                nw308_[int(12)] = d_1789_v140_
                d_1790_v141_ = nw308_
                rhs342_ = d_1790_v141_
                rhs343_ = d_1624_v1_
                d_1790_v141_ = rhs342_
                d_1624_v1_ = rhs343_
                (globalState).f1 = d_1787_v138_
                (self).m9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdpjpt")), globalState)
            d_1791_v142_: int
            d_1791_v142_ = 790
            d_1792_v143_: _dafny.Map
            d_1792_v143_ = _dafny.Map({763: d_1791_v142_})
            (globalState).f1 = ((d_1792_v143_)[d_1791_v142_] if (d_1791_v142_) in (d_1792_v143_) else (d_1791_v142_) - (d_1791_v142_))
            d_1793_v144_: _dafny.Seq
            d_1793_v144_ = _dafny.SeqWithoutIsStrInference([d_1624_v1_])
            d_1794_v145_: _dafny.MultiSet
            d_1794_v145_ = _dafny.MultiSet([d_1793_v144_])
            d_1794_v145_ = _dafny.MultiSet([d_1793_v144_])
            d_1795_v146_: _dafny.Set
            d_1795_v146_ = _dafny.Set({d_1624_v1_})
            if (d_1795_v146_).issubset((_dafny.Set({d_1624_v1_, d_1624_v1_, d_1624_v1_, d_1624_v1_}) if d_1624_v1_ else d_1795_v146_)):
                d_1796_v147_: _dafny.Array
                nw309_ = _dafny.Array(False, 5)
                d_1796_v147_ = nw309_
                d_1797_v148_: _dafny.Map
                d_1797_v148_ = _dafny.Map({d_1624_v1_: d_1791_v142_})
                d_1798_v149_: D15
                d_1798_v149_ = D15_DC28(d_1624_v1_, d_1624_v1_)
                pat_let_tv23_ = d_1624_v1_
                pat_let_tv24_ = globalState
                pat_let_tv25_ = d_1792_v143_
                def iife121_(_pat_let28_0):
                    def iife122_(d_1799_dt__update__tmp_h4_):
                        def iife123_(_pat_let29_0):
                            def iife124_(d_1800_dt__update_hcf36_h0_):
                                return D15_DC28(d_1800_dt__update_hcf36_h0_, (d_1799_dt__update__tmp_h4_).cf37)
                            return iife124_(_pat_let29_0)
                        return iife123_((D11_DC20(default__.fm0(pat_let_tv23_, pat_let_tv24_), len(pat_let_tv25_))).cf21)
                    return iife122_(_pat_let28_0)
                rhs344_ = (iife121_(d_1798_v149_)).cf37
                rhs345_ = d_1791_v142_
                rhs346_ = (d_1795_v146_).intersection(d_1795_v146_)
                rhs347_ = d_1796_v147_
                rhs348_ = d_1797_v148_
                lhs237_ = globalState
                d_1624_v1_ = rhs344_
                lhs237_.f1 = rhs345_
                d_1795_v146_ = rhs346_
                d_1796_v147_ = rhs347_
                d_1797_v148_ = rhs348_
                d_1801_v150_: _dafny.MultiSet
                d_1801_v150_ = _dafny.MultiSet([d_1624_v1_])
                d_1802_v151_: _dafny.Seq
                d_1802_v151_ = _dafny.SeqWithoutIsStrInference([d_1791_v142_])
                d_1803_v152_: _dafny.Array
                nw310_ = _dafny.Array(None, 19)
                nw310_[int(0)] = d_1791_v142_
                nw310_[int(1)] = d_1791_v142_
                nw310_[int(2)] = (self).fm7(d_1624_v1_, globalState)
                nw310_[int(3)] = 413
                nw310_[int(4)] = d_1791_v142_
                nw310_[int(5)] = 342
                nw310_[int(6)] = d_1791_v142_
                nw310_[int(7)] = d_1791_v142_
                nw310_[int(8)] = (d_1801_v150_).cardinality
                nw310_[int(9)] = d_1791_v142_
                nw310_[int(10)] = len(d_1802_v151_)
                nw310_[int(11)] = d_1791_v142_
                nw310_[int(12)] = d_1791_v142_
                nw310_[int(13)] = d_1791_v142_
                nw310_[int(14)] = 166
                nw310_[int(15)] = d_1791_v142_
                nw310_[int(16)] = ((d_1797_v148_)[True] if (True) in (d_1797_v148_) else d_1791_v142_)
                nw310_[int(17)] = d_1791_v142_
                nw310_[int(18)] = (0) - (d_1791_v142_)
                d_1803_v152_ = nw310_
                d_1804_v153_: _dafny.Seq
                d_1804_v153_ = _dafny.SeqWithoutIsStrInference([d_1803_v152_, d_1803_v152_, d_1803_v152_])
                d_1805_v154_: _dafny.Map
                d_1805_v154_ = _dafny.Map({(0) - (d_1791_v142_): (d_1804_v153_)[default__.safeIndex(d_1791_v142_, len(d_1804_v153_))]})
                rhs349_ = ((len(d_1623_v0_) if d_1624_v1_ else d_1791_v142_) if d_1624_v1_ else (d_1791_v142_) - (d_1791_v142_))
                rhs350_ = d_1805_v154_
                lhs238_ = globalState
                lhs238_.f1 = rhs349_
                d_1805_v154_ = rhs350_
                d_1806_v155_: D0
                d_1806_v155_ = D0_DC1()
                d_1807_v156_: D0
                d_1807_v156_ = D0_DC3(d_1806_v155_)
                d_1808_v157_: str
                d_1808_v157_ = _dafny.CodePoint('u')
                d_1809_v158_: C2
                nw311_ = C2()
                nw311_.ctor__(d_1808_v157_, default__.fm0(d_1624_v1_, globalState))
                d_1809_v158_ = nw311_
                d_1810_v159_: _dafny.Map
                d_1810_v159_ = _dafny.Map({d_1807_v156_: d_1809_v158_})
                d_1810_v159_ = (d_1810_v159_).set(d_1807_v156_, d_1809_v158_)
                d_1811_v160_: D11
                d_1811_v160_ = D11_DC20((d_1809_v158_).fm6(d_1802_v151_, globalState), d_1791_v142_)
                d_1797_v148_ = (d_1797_v148_).set((d_1811_v160_).cf21, d_1791_v142_)
                d_1803_v152_ = d_1803_v152_
            elif True:
                d_1624_v1_ = not (d_1624_v1_) or (d_1624_v1_)
                d_1812_v161_: _dafny.Map
                d_1812_v161_ = _dafny.Map({d_1623_v0_: d_1624_v1_})
                d_1812_v161_ = (d_1812_v161_).set((d_1623_v0_) + (d_1623_v0_), not(d_1624_v1_))
                d_1813_v162_: C1
                nw312_ = C1()
                nw312_.ctor__()
                d_1813_v162_ = nw312_
                d_1814_v163_: _dafny.Map
                d_1814_v163_ = _dafny.Map({d_1624_v1_: d_1624_v1_})
                d_1815_v164_: D12
                d_1815_v164_ = D12_DC21(d_1814_v163_)
                d_1816_v165_: _dafny.Map
                d_1816_v165_ = _dafny.Map({d_1624_v1_: d_1815_v164_})
                d_1816_v165_ = d_1816_v165_
                d_1624_v1_ = d_1624_v1_
            (globalState).f1 = default__.safeModuloInt(d_1791_v142_, d_1791_v142_)
        d_1817_v166_: int
        d_1817_v166_ = 892
        d_1818_v167_: D1
        d_1818_v167_ = D1_DC4(d_1817_v166_)
        d_1819_v168_: D2
        d_1819_v168_ = D2_DC8(d_1624_v1_)
        pat_let_tv26_ = d_1817_v166_
        def iife125_(_pat_let30_0):
            def iife126_(d_1820_dt__update__tmp_h5_):
                def iife127_(_pat_let31_0):
                    def iife128_(d_1821_dt__update_hcf4_h0_):
                        return D1_DC4(d_1821_dt__update_hcf4_h0_)
                    return iife128_(_pat_let31_0)
                return iife127_(pat_let_tv26_)
            return iife126_(_pat_let30_0)
        d_1818_v167_ = iife125_((self).fm5(d_1819_v168_, globalState))
        d_1822_v169_: _dafny.Seq
        d_1822_v169_ = _dafny.SeqWithoutIsStrInference([-263])
        d_1823_v170_: _dafny.Seq
        d_1823_v170_ = _dafny.SeqWithoutIsStrInference([d_1822_v169_, d_1822_v169_, d_1822_v169_])
        d_1824_v171_: _dafny.Seq
        d_1824_v171_ = _dafny.SeqWithoutIsStrInference([d_1822_v169_, d_1822_v169_, d_1822_v169_, d_1822_v169_, (d_1823_v170_)[default__.safeIndex(d_1817_v166_, len(d_1823_v170_))]])
        d_1825_v172_: C0
        nw313_ = C0()
        nw313_.ctor__((d_1624_v1_ if d_1624_v1_ else not(d_1624_v1_)), (d_1823_v170_) + (d_1823_v170_))
        d_1825_v172_ = nw313_
        d_1826_v173_: _dafny.Map
        d_1826_v173_ = _dafny.Map({not((d_1825_v172_).f11): d_1817_v166_})
        d_1827_v174_: _dafny.MultiSet
        d_1827_v174_ = _dafny.MultiSet([d_1624_v1_])
        d_1828_v175_: _dafny.Seq
        d_1828_v175_ = _dafny.SeqWithoutIsStrInference([(d_1819_v168_).cf9, d_1624_v1_, (d_1825_v172_).f11])
        d_1829_v176_: _dafny.Map
        d_1829_v176_ = _dafny.Map({-512: len(_dafny.SeqWithoutIsStrInference([d_1624_v1_, False]))})
        d_1830_v177_: _dafny.Set
        d_1830_v177_ = _dafny.Set({(d_1825_v172_).f11, (d_1825_v172_).f11})
        d_1831_v178_: _dafny.Array
        nw314_ = _dafny.Array(None, 27)
        nw314_[int(0)] = d_1817_v166_
        nw314_[int(1)] = len((d_1826_v173_) | (default__.fm57((0) - (d_1817_v166_), (d_1827_v174_).cardinality, globalState)))
        nw314_[int(2)] = d_1817_v166_
        nw314_[int(3)] = (d_1817_v166_) + (d_1817_v166_)
        nw314_[int(4)] = d_1817_v166_
        nw314_[int(5)] = (0) - (d_1817_v166_)
        nw314_[int(6)] = d_1817_v166_
        nw314_[int(7)] = d_1817_v166_
        nw314_[int(8)] = d_1817_v166_
        nw314_[int(9)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmtyrbao"))), d_1817_v166_)
        nw314_[int(10)] = (0) - ((len(d_1822_v169_)) - (d_1817_v166_))
        nw314_[int(11)] = d_1817_v166_
        nw314_[int(12)] = len(d_1828_v175_)
        nw314_[int(13)] = default__.safeDivisionInt((0) - (d_1817_v166_), d_1817_v166_)
        nw314_[int(14)] = d_1817_v166_
        nw314_[int(15)] = ((d_1829_v176_)[((d_1829_v176_)[d_1817_v166_] if (d_1817_v166_) in (d_1829_v176_) else d_1817_v166_)] if (((d_1829_v176_)[d_1817_v166_] if (d_1817_v166_) in (d_1829_v176_) else d_1817_v166_)) in (d_1829_v176_) else len(_dafny.Map({d_1817_v166_: d_1817_v166_})))
        nw314_[int(16)] = d_1817_v166_
        nw314_[int(17)] = len(d_1822_v169_)
        nw314_[int(18)] = (d_1817_v166_ if False else len(d_1828_v175_))
        nw314_[int(19)] = d_1817_v166_
        nw314_[int(20)] = d_1817_v166_
        nw314_[int(21)] = len(d_1623_v0_)
        nw314_[int(22)] = default__.safeModuloInt(d_1817_v166_, d_1817_v166_)
        nw314_[int(23)] = len((d_1830_v177_ if False else d_1830_v177_))
        nw314_[int(24)] = (d_1817_v166_) + (d_1817_v166_)
        nw314_[int(25)] = d_1817_v166_
        nw314_[int(26)] = d_1817_v166_
        d_1831_v178_ = nw314_
        rhs351_ = d_1831_v178_
        rhs352_ = d_1624_v1_
        d_1831_v178_ = rhs351_
        d_1624_v1_ = rhs352_
        r0 = d_1817_v166_
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_1832_v0_: int
        d_1832_v0_ = 442
        d_1833_v1_: _dafny.Seq
        d_1833_v1_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(default__.fm3(p0, default__.fm0(p0, globalState), globalState), d_1832_v0_)])
        d_1833_v1_ = d_1833_v1_
        d_1834_v2_: _dafny.Seq
        d_1834_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqnx"))
        d_1832_v0_ = default__.safeModuloInt(len(d_1834_v2_), d_1832_v0_)
        d_1835_i0_: int
        d_1835_i0_ = 0
        with _dafny.label("17"):
            while p0:
                with _dafny.c_label("17"):
                    if (d_1835_i0_) >= (100):
                        raise _dafny.Break("17")
                    d_1835_i0_ = (d_1835_i0_) + (1)
                    d_1836_v3_: bool
                    d_1836_v3_ = True
                    d_1837_v4_: _dafny.Array
                    nw315_ = _dafny.Array(None, 5)
                    nw315_[int(0)] = p0
                    nw315_[int(1)] = (p0) and (p2)
                    nw315_[int(2)] = default__.fm0(p2, globalState)
                    nw315_[int(3)] = not(d_1836_v3_)
                    nw315_[int(4)] = p0
                    d_1837_v4_ = nw315_
                    d_1838_v5_: _dafny.Seq
                    d_1838_v5_ = _dafny.SeqWithoutIsStrInference([not(p0)])
                    index322_ = default__.safeIndex(646, (d_1837_v4_).length(0))
                    (d_1837_v4_)[index322_] = (d_1838_v5_)[default__.safeIndex(d_1832_v0_, len(d_1838_v5_))]
                    index323_ = default__.safeIndex(646, (d_1837_v4_).length(0))
                    rhs353_ = (d_1838_v5_)[default__.safeIndex(default__.safeDivisionInt(d_1832_v0_, d_1832_v0_), len(d_1838_v5_))]
                    rhs354_ = not (p0) or (p2)
                    lhs239_ = d_1837_v4_
                    lhs240_ = default__.safeIndex(646, (d_1837_v4_).length(0))
                    d_1836_v3_ = rhs353_
                    lhs239_[lhs240_] = rhs354_
                    d_1839_v6_: _dafny.Set
                    d_1839_v6_ = _dafny.Set({d_1836_v3_, p2})
                    d_1840_v7_: _dafny.Seq
                    d_1840_v7_ = _dafny.SeqWithoutIsStrInference([d_1839_v6_, d_1839_v6_, d_1839_v6_, default__.fm49(globalState)])
                    d_1841_v8_: D16
                    d_1841_v8_ = D16_DC30(d_1840_v7_)
                    (globalState).f1 = len((d_1841_v8_).cf39)
                    (globalState).f1 = (default__.fm58(globalState)).cf22
                    d_1842_v9_: D0
                    d_1842_v9_ = D0_DC1()
                    d_1843_v10_: _dafny.Set
                    d_1843_v10_ = _dafny.Set({d_1842_v9_, d_1842_v9_, d_1842_v9_})
                    d_1844_v11_: _dafny.Set
                    d_1844_v11_ = d_1843_v10_
                    d_1844_v11_ = d_1844_v11_
                    pass
            pass
        d_1845_v12_: _dafny.Array
        nw316_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_1845_v12_ = nw316_
        index324_ = default__.safeIndex(92, (d_1845_v12_).length(0))
        (d_1845_v12_)[index324_] = d_1834_v2_
        index325_ = default__.safeIndex(92, (d_1845_v12_).length(0))
        (d_1845_v12_)[index325_] = (d_1834_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))
        d_1846_v13_: bool
        d_1846_v13_ = False
        d_1847_v14_: _dafny.Seq
        d_1847_v14_ = _dafny.SeqWithoutIsStrInference([(d_1845_v12_)[default__.safeIndex(92, (d_1845_v12_).length(0))], (d_1845_v12_)[default__.safeIndex(92, (d_1845_v12_).length(0))]])
        d_1846_v13_ = (True if p2 else ((d_1845_v12_)[default__.safeIndex(92, (d_1845_v12_).length(0))]) not in (d_1847_v14_))
        d_1848_v15_: _dafny.Array
        nw317_ = _dafny.Array(None, 3)
        nw317_[int(0)] = p0
        nw317_[int(1)] = d_1846_v13_
        nw317_[int(2)] = p0
        d_1848_v15_ = nw317_
        index326_ = default__.safeIndex(125, (d_1848_v15_).length(0))
        (d_1848_v15_)[index326_] = d_1846_v13_
        d_1849_v16_: _dafny.Map
        d_1849_v16_ = _dafny.Map({d_1832_v0_: d_1832_v0_})
        d_1850_v17_: _dafny.Map
        d_1850_v17_ = d_1849_v16_
        pat_let_tv27_ = p0
        index327_ = default__.safeIndex(125, (d_1848_v15_).length(0))
        def lambda89_(source24_):
            d_1851___mcc_h0_ = source24_
            d_1852_cf11_ = d_1851___mcc_h0_
            return pat_let_tv27_

        (d_1848_v15_)[index327_] = not(lambda89_(d_1849_v16_))

    def m8(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: C1 = None
        d_1853_v0_: _dafny.Array
        nw318_ = _dafny.Array(False, 10)
        d_1853_v0_ = nw318_
        index328_ = default__.safeIndex(282, (d_1853_v0_).length(0))
        (d_1853_v0_)[index328_] = default__.fm0(p1, globalState)
        d_1854_v1_: int
        d_1854_v1_ = -792
        index329_ = default__.safeIndex(112, (p2).length(0))
        (p2)[index329_] = d_1854_v1_
        d_1855_v2_: _dafny.Seq
        d_1855_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyvrsduxe"))
        d_1856_v3_: _dafny.Seq
        d_1856_v3_ = _dafny.SeqWithoutIsStrInference([p0, p1])
        d_1857_v4_: _dafny.MultiSet
        d_1857_v4_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thordvy")))])
        d_1858_v5_: D1
        d_1858_v5_ = D1_DC4(d_1854_v1_)
        index330_ = default__.safeIndex(282, (d_1853_v0_).length(0))
        index331_ = default__.safeIndex(112, (p2).length(0))
        rhs355_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1859_i0_ in range(default__.abs(941))]))) == (-462)
        rhs356_ = len((d_1855_v2_) + (d_1855_v2_))
        rhs357_ = ((0) - ((d_1854_v1_) - (d_1854_v1_))) * (default__.safeDivisionInt(d_1854_v1_, d_1854_v1_))
        rhs358_ = ((self).fm4(d_1856_v3_, d_1854_v1_, d_1854_v1_, d_1857_v4_, globalState)) - ((d_1858_v5_).cf4)
        lhs241_ = d_1853_v0_
        lhs242_ = default__.safeIndex(282, (d_1853_v0_).length(0))
        lhs243_ = p2
        lhs244_ = default__.safeIndex(112, (p2).length(0))
        lhs245_ = globalState
        lhs246_ = globalState
        lhs241_[lhs242_] = rhs355_
        lhs243_[lhs244_] = rhs356_
        lhs245_.f1 = rhs357_
        lhs246_.f1 = rhs358_
        d_1855_v2_ = d_1855_v2_
        d_1860_v6_: _dafny.Map
        d_1860_v6_ = _dafny.Map({(p2)[default__.safeIndex(112, (p2).length(0))]: (p2)[default__.safeIndex(112, (p2).length(0))]})
        d_1861_v7_: str
        d_1861_v7_ = _dafny.CodePoint('u')
        d_1862_v8_: D0
        d_1863_v9_: bool
        out46_: D0
        out47_: bool
        out46_, out47_ = default__.m0(d_1853_v0_, default__.fm36(default__.fm3(p1, p1, globalState), (p2)[default__.safeIndex(112, (p2).length(0))], d_1854_v1_, d_1860_v6_, globalState), (d_1855_v2_) == (d_1855_v2_), d_1861_v7_, globalState)
        d_1862_v8_ = out46_
        d_1863_v9_ = out47_
        d_1861_v7_ = d_1861_v7_
        index332_ = default__.safeIndex(112, (p2).length(0))
        (p2)[index332_] = default__.safeModuloInt((self).fm7(p0, globalState), d_1854_v1_)
        d_1864_v10_: _dafny.Array
        d_1864_v10_ = p2
        d_1865_v11_: _dafny.Array
        nw319_ = _dafny.Array(None, 9)
        nw319_[int(0)] = p2
        nw319_[int(1)] = p2
        nw319_[int(2)] = p2
        nw319_[int(3)] = p2
        nw319_[int(4)] = p2
        nw319_[int(5)] = p2
        nw319_[int(6)] = (d_1864_v10_)
        nw319_[int(7)] = p2
        nw319_[int(8)] = p2
        d_1865_v11_ = nw319_
        index333_ = default__.safeIndex(760, (d_1865_v11_).length(0))
        (d_1865_v11_)[index333_] = p2
        d_1866_v12_: _dafny.Set
        d_1866_v12_ = _dafny.Set({d_1861_v7_, d_1861_v7_})
        d_1867_v13_: _dafny.Set
        d_1867_v13_ = _dafny.Set({(_dafny.Set({d_1861_v7_})) | (d_1866_v12_)})
        index334_ = default__.safeIndex(760, (d_1865_v11_).length(0))
        rhs359_ = p2
        rhs360_ = (d_1867_v13_) | (d_1867_v13_)
        lhs247_ = d_1865_v11_
        lhs248_ = default__.safeIndex(760, (d_1865_v11_).length(0))
        lhs247_[lhs248_] = rhs359_
        d_1867_v13_ = rhs360_
        r0 = (p2)[default__.safeIndex(112, (p2).length(0))]
        d_1868_v14_: C1
        nw320_ = C1()
        nw320_.ctor__()
        d_1868_v14_ = nw320_
        r1 = d_1868_v14_
        return r0, r1

    def m9(self, p0, globalState):
        d_1869_v0_: int
        d_1869_v0_ = -163
        d_1870_v1_: bool
        d_1870_v1_ = True
        d_1871_v2_: _dafny.Set
        d_1871_v2_ = _dafny.Set({p0})
        d_1872_v3_: _dafny.Seq
        d_1872_v3_ = _dafny.SeqWithoutIsStrInference([True])
        d_1873_v4_: _dafny.MultiSet
        d_1873_v4_ = _dafny.MultiSet([d_1869_v0_])
        d_1874_v5_: _dafny.Map
        d_1874_v5_ = _dafny.Map({d_1869_v0_: d_1869_v0_})
        d_1875_v6_: _dafny.Map
        d_1875_v6_ = _dafny.Map({d_1869_v0_: d_1870_v1_})
        d_1876_v7_: D16
        d_1876_v7_ = D16_DC32(((d_1875_v6_)[230] if (230) in (d_1875_v6_) else d_1870_v1_))
        d_1877_v8_: _dafny.Seq
        d_1877_v8_ = _dafny.SeqWithoutIsStrInference([d_1876_v7_, D16_DC32(not(d_1870_v1_)), d_1876_v7_])
        d_1878_v9_: _dafny.Seq
        d_1878_v9_ = _dafny.SeqWithoutIsStrInference([d_1869_v0_])
        d_1879_v10_: _dafny.Array
        nw321_ = _dafny.Array(None, 24)
        nw321_[int(0)] = 167
        nw321_[int(1)] = (d_1869_v0_) + (d_1869_v0_)
        nw321_[int(2)] = (d_1869_v0_ if d_1870_v1_ else d_1869_v0_)
        nw321_[int(3)] = d_1869_v0_
        nw321_[int(4)] = len(d_1871_v2_)
        nw321_[int(5)] = (0) - (len(default__.fm59(globalState)))
        nw321_[int(6)] = d_1869_v0_
        nw321_[int(7)] = d_1869_v0_
        nw321_[int(8)] = d_1869_v0_
        nw321_[int(9)] = 158
        nw321_[int(10)] = d_1869_v0_
        nw321_[int(11)] = d_1869_v0_
        nw321_[int(12)] = d_1869_v0_
        nw321_[int(13)] = default__.fm3(d_1870_v1_, d_1870_v1_, globalState)
        nw321_[int(14)] = (_dafny.MultiSet(d_1872_v3_)).cardinality
        nw321_[int(15)] = -478
        nw321_[int(16)] = default__.safeModuloInt(default__.fm3(d_1870_v1_, d_1870_v1_, globalState), ((d_1873_v4_)[d_1869_v0_] if (d_1869_v0_) in (d_1873_v4_) else (0) - (d_1869_v0_)))
        nw321_[int(17)] = ((d_1873_v4_) | ((d_1873_v4_).set(d_1869_v0_, default__.abs(d_1869_v0_)))).cardinality
        nw321_[int(18)] = d_1869_v0_
        nw321_[int(19)] = len(default__.fm49(globalState))
        nw321_[int(20)] = len(default__.fm57(len(d_1874_v5_), (0) - (d_1869_v0_), globalState))
        nw321_[int(21)] = (0) - (len(d_1877_v8_))
        nw321_[int(22)] = (d_1869_v0_) - (len(d_1878_v9_))
        nw321_[int(23)] = d_1869_v0_
        d_1879_v10_ = nw321_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1879_v10_).length(0)):
            d_1880_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_1880_i0_)) and ((d_1880_i0_) < ((d_1879_v10_).length(0)))):
                (d_1879_v10_)[(d_1880_i0_)] = default__.safeModuloInt(d_1880_i0_, ((d_1874_v5_)[d_1869_v0_] if (d_1869_v0_) in (d_1874_v5_) else len(_dafny.SeqWithoutIsStrInference([D15_DC29(D15_DC27(_dafny.MultiSet([d_1870_v1_, False])))]))))
        d_1870_v1_ = (d_1869_v0_) <= (d_1869_v0_)
        d_1881_i1_: int
        d_1881_i1_ = 0
        with _dafny.label("18"):
            while not(default__.fm0((self).fm6(d_1878_v9_, globalState), globalState)):
                with _dafny.c_label("18"):
                    if (d_1881_i1_) >= (100):
                        raise _dafny.Break("18")
                    d_1881_i1_ = (d_1881_i1_) + (1)
                    d_1882_v11_: _dafny.Map
                    d_1882_v11_ = d_1875_v6_
                    source25_ = d_1882_v11_
                    d_1883___mcc_h0_ = source25_
                    d_1884_cf12_ = d_1883___mcc_h0_
                    (globalState).f1 = d_1869_v0_
                    d_1885_v12_: str
                    d_1885_v12_ = _dafny.CodePoint('c')
                    d_1886_v13_: D14
                    d_1886_v13_ = D14_DC26(True, d_1869_v0_, d_1885_v12_)
                    d_1887_v14_: _dafny.MultiSet
                    d_1887_v14_ = _dafny.MultiSet([_dafny.Map({d_1870_v1_: -334})])
                    d_1888_v15_: _dafny.Array
                    nw322_ = _dafny.Array(None, 28)
                    nw322_[int(0)] = D14_DC26(d_1870_v1_, d_1869_v0_, d_1885_v12_)
                    nw322_[int(1)] = d_1886_v13_
                    nw322_[int(2)] = d_1886_v13_
                    nw322_[int(3)] = d_1886_v13_
                    nw322_[int(4)] = d_1886_v13_
                    nw322_[int(5)] = d_1886_v13_
                    nw322_[int(6)] = d_1886_v13_
                    nw322_[int(7)] = d_1886_v13_
                    nw322_[int(8)] = d_1886_v13_
                    nw322_[int(9)] = d_1886_v13_
                    nw322_[int(10)] = d_1886_v13_
                    nw322_[int(11)] = d_1886_v13_
                    nw322_[int(12)] = d_1886_v13_
                    nw322_[int(13)] = d_1886_v13_
                    nw322_[int(14)] = d_1886_v13_
                    nw322_[int(15)] = d_1886_v13_
                    nw322_[int(16)] = D14_DC26(d_1870_v1_, 538, d_1885_v12_)
                    nw322_[int(17)] = D14_DC26((d_1872_v3_)[default__.safeIndex((d_1887_v14_).cardinality, len(d_1872_v3_))], 245, d_1885_v12_)
                    nw322_[int(18)] = d_1886_v13_
                    nw322_[int(19)] = d_1886_v13_
                    nw322_[int(20)] = d_1886_v13_
                    nw322_[int(21)] = d_1886_v13_
                    nw322_[int(22)] = d_1886_v13_
                    nw322_[int(23)] = d_1886_v13_
                    nw322_[int(24)] = d_1886_v13_
                    nw322_[int(25)] = default__.fm60(globalState)
                    nw322_[int(26)] = (d_1886_v13_ if d_1870_v1_ else d_1886_v13_)
                    nw322_[int(27)] = default__.fm60(globalState)
                    d_1888_v15_ = nw322_
                    d_1888_v15_ = d_1888_v15_
                    d_1889_v16_: _dafny.Array
                    nw323_ = _dafny.Array(_dafny.Map({}), 10)
                    d_1889_v16_ = nw323_
                    index335_ = default__.safeIndex(333, (d_1889_v16_).length(0))
                    (d_1889_v16_)[index335_] = (d_1874_v5_) | (d_1874_v5_)
                    index336_ = default__.safeIndex(333, (d_1889_v16_).length(0))
                    (d_1889_v16_)[index336_] = d_1874_v5_
                    d_1890_v17_: T1
                    nw324_ = C3()
                    nw324_.ctor__(d_1870_v1_, d_1870_v1_)
                    d_1890_v17_ = nw324_
                    d_1891_v18_: _dafny.Map
                    d_1891_v18_ = _dafny.Map({d_1890_v17_: d_1870_v1_})
                    d_1870_v1_ = ((d_1891_v18_)[d_1890_v17_] if (d_1890_v17_) in (d_1891_v18_) else d_1870_v1_)
                    d_1870_v1_ = d_1870_v1_
                    d_1892_v19_: _dafny.Seq
                    d_1892_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1869_v0_ for d_1893_i2_ in range(default__.abs(-648))]), d_1878_v9_])
                    d_1894_v20_: C0
                    nw325_ = C0()
                    nw325_.ctor__(d_1870_v1_, d_1892_v19_)
                    d_1894_v20_ = nw325_
                    d_1895_v21_: C1
                    nw326_ = C1()
                    nw326_.ctor__()
                    d_1895_v21_ = nw326_
                    pass
            pass
        d_1896_i3_: int
        d_1896_i3_ = 0
        with _dafny.label("19"):
            while (default__.safeModuloInt(d_1869_v0_, d_1869_v0_)) <= (d_1869_v0_):
                with _dafny.c_label("19"):
                    if (d_1896_i3_) >= (100):
                        raise _dafny.Break("19")
                    d_1896_i3_ = (d_1896_i3_) + (1)
                    if d_1870_v1_:
                        d_1897_v22_: _dafny.Map
                        d_1897_v22_ = _dafny.Map({d_1870_v1_: not(d_1870_v1_)})
                        d_1897_v22_ = (d_1897_v22_).set((d_1878_v9_) < (d_1878_v9_), d_1870_v1_)
                        (globalState).f1 = (d_1869_v0_) - ((d_1869_v0_) * (d_1869_v0_))
                        d_1898_v23_: str
                        d_1898_v23_ = _dafny.CodePoint('p')
                        rhs361_ = (0) - (d_1869_v0_)
                        rhs362_ = d_1870_v1_
                        rhs363_ = d_1869_v0_
                        rhs364_ = d_1898_v23_
                        lhs249_ = globalState
                        lhs250_ = globalState
                        lhs249_.f1 = rhs361_
                        d_1870_v1_ = rhs362_
                        lhs250_.f1 = rhs363_
                        d_1898_v23_ = rhs364_
                        d_1869_v0_ = d_1869_v0_
                        d_1899_v24_: _dafny.Array
                        nw327_ = _dafny.Array(None, 1)
                        nw327_[int(0)] = d_1870_v1_
                        d_1899_v24_ = nw327_
                        d_1900_v25_: _dafny.Seq
                        d_1900_v25_ = _dafny.SeqWithoutIsStrInference([d_1899_v24_])
                        d_1901_v26_: _dafny.Seq
                        d_1901_v26_ = _dafny.SeqWithoutIsStrInference([(d_1900_v25_)[default__.safeIndex(d_1869_v0_, len(d_1900_v25_))], d_1899_v24_, d_1899_v24_])
                        d_1899_v24_ = (d_1901_v26_)[default__.safeIndex(d_1869_v0_, len(d_1901_v26_))]
                    elif True:
                        (globalState).f1 = d_1869_v0_
                        d_1902_v27_: _dafny.Map
                        d_1902_v27_ = _dafny.Map({(p0) + (p0): (self).fm29(d_1869_v0_, p0, globalState)})
                        d_1902_v27_ = (d_1902_v27_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqjlgq")), ((d_1874_v5_)[d_1869_v0_] if (d_1869_v0_) in (d_1874_v5_) else d_1869_v0_))
                        (globalState).f1 = d_1869_v0_
                        d_1903_v28_: C1
                        nw328_ = C1()
                        nw328_.ctor__()
                        d_1903_v28_ = nw328_
                        d_1873_v4_ = default__.fm32(globalState)
                    d_1904_v29_: _dafny.Array
                    nw329_ = _dafny.Array(False, 3)
                    d_1904_v29_ = nw329_
                    d_1905_v30_: _dafny.Array
                    d_1905_v30_ = d_1904_v29_
                    d_1906_v31_: _dafny.Map
                    d_1906_v31_ = _dafny.Map({d_1905_v30_: not(d_1870_v1_)})
                    d_1906_v31_ = (d_1906_v31_).set(d_1904_v29_, False)
                    d_1870_v1_ = d_1870_v1_
                    d_1907_v32_: C3
                    nw330_ = C3()
                    nw330_.ctor__(d_1870_v1_, d_1870_v1_)
                    d_1907_v32_ = nw330_
                    pass
            pass
        d_1908_v33_: str
        d_1908_v33_ = _dafny.CodePoint('v')
        d_1909_v34_: C5
        nw331_ = C5()
        nw331_.ctor__(d_1869_v0_, d_1908_v33_, not((d_1870_v1_) or (d_1870_v1_)))
        d_1909_v34_ = nw331_
        d_1910_v35_: _dafny.Map
        d_1910_v35_ = _dafny.Map({d_1879_v10_: d_1872_v3_})
        d_1910_v35_ = (d_1910_v35_).set(d_1879_v10_, (d_1872_v3_) + (d_1872_v3_))


class C7(T3, T2, T1, T0):
    def  __init__(self):
        self._f5: str = _dafny.CodePoint('D')
        self._f6: bool = False
        self.f10: int = int(0)
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f9, f10, f5, f6):
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f5 = f5
        (self)._f6 = f6

    def fm8(self, p0, p1, p2, globalState):
        return default__.safeDivisionInt((self).f9, self.f10)

    def fm6(self, p0, globalState):
        return not((self).f6)

    def fm7(self, p0, globalState):
        return (self).f9

    def fm4(self, p0, p1, p2, p3, globalState):
        return self.f10

    def fm5(self, p0, globalState):
        if (self).f6:
            return D1_DC4((self).f9)
        elif True:
            return D1_DC4(self.f10)

    def fm13(self, globalState):
        return ((self).f9) == (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mhlxx"))))

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1911_v0_: _dafny.Array
        def lambda90_(d_1912_p2_):
            def lambda91_(d_1913_i0_):
                return ((self).f5) not in (d_1912_p2_)

            return lambda91_

        init52_ = lambda90_(p2)
        nw332_ = _dafny.Array(None, 29)
        for i0_52_ in range(nw332_.length(0)):
            nw332_[i0_52_] = init52_(i0_52_)
        d_1911_v0_ = nw332_
        index337_ = default__.safeIndex(292, (d_1911_v0_).length(0))
        (d_1911_v0_)[index337_] = (not(p0) if not((self).f6) else (self).f6)
        index338_ = default__.safeIndex(292, (d_1911_v0_).length(0))
        (d_1911_v0_)[index338_] = p0
        d_1914_v1_: _dafny.MultiSet
        d_1914_v1_ = _dafny.MultiSet([p0])
        (self).f10 = (self.f10) - (default__.safeModuloInt((d_1914_v1_).cardinality, self.f10))
        d_1915_v2_: D2
        d_1915_v2_ = D2_DC6(_dafny.SeqWithoutIsStrInference([(self).f5 for d_1916_i1_ in range(default__.abs(563))]))
        source26_ = d_1915_v2_
        if source26_.is_DC7:
            d_1917___mcc_h0_ = source26_.cf7
            d_1918___mcc_h1_ = source26_.cf8
            d_1919_cf8_ = d_1918___mcc_h1_
            d_1920_cf7_ = d_1917___mcc_h0_
            d_1921_v3_: str
            d_1921_v3_ = _dafny.CodePoint('l')
            d_1921_v3_ = (self).f5
            d_1922_v4_: _dafny.Seq
            d_1922_v4_ = _dafny.SeqWithoutIsStrInference([self.f10, 929, d_1920_cf7_, (0) - ((self).f9)])
            (globalState).f1 = len(d_1922_v4_)
            d_1923_v5_: _dafny.Seq
            d_1923_v5_ = _dafny.SeqWithoutIsStrInference([(d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], (self).f6])
            d_1923_v5_ = default__.fm2((d_1923_v5_)[default__.safeIndex((0) - ((0) - (len(_dafny.Set({p1, self.f10})))), len(d_1923_v5_))], (d_1923_v5_)[default__.safeIndex(p1, len(d_1923_v5_))], globalState)
            r1 = p2
        elif source26_.is_DC8:
            d_1924___mcc_h2_ = source26_.cf9
            d_1925_cf9_ = d_1924___mcc_h2_
            d_1926_v6_: _dafny.Set
            d_1926_v6_ = _dafny.Set({True, (d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], p0, d_1925_cf9_})
            d_1927_v7_: _dafny.Seq
            d_1927_v7_ = _dafny.SeqWithoutIsStrInference([False])
            d_1928_v8_: _dafny.Map
            d_1928_v8_ = _dafny.Map({d_1927_v7_: p1})
            d_1929_v9_: _dafny.Seq
            d_1929_v9_ = _dafny.SeqWithoutIsStrInference([d_1927_v7_])
            d_1930_v10_: _dafny.Map
            d_1930_v10_ = _dafny.Map({not(d_1925_cf9_): True})
            d_1931_v11_: _dafny.Seq
            d_1931_v11_ = _dafny.SeqWithoutIsStrInference([p1, len(p2), p1, p1])
            d_1932_v12_: _dafny.Array
            nw333_ = _dafny.Array(None, 12)
            nw333_[int(0)] = len(d_1926_v6_)
            nw333_[int(1)] = (self.f10) + (((d_1928_v8_)[(d_1929_v9_)[default__.safeIndex(self.f10, len(d_1929_v9_))]] if ((d_1929_v9_)[default__.safeIndex(self.f10, len(d_1929_v9_))]) in (d_1928_v8_) else self.f10))
            nw333_[int(2)] = self.f10
            nw333_[int(3)] = self.f10
            nw333_[int(4)] = self.f10
            nw333_[int(5)] = self.f10
            nw333_[int(6)] = len(p2)
            nw333_[int(7)] = 534
            nw333_[int(8)] = default__.safeDivisionInt(len(d_1930_v10_), len(d_1931_v11_))
            nw333_[int(9)] = p1
            nw333_[int(10)] = 599
            nw333_[int(11)] = (self).fm7((d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], globalState)
            d_1932_v12_ = nw333_
            index339_ = default__.safeIndex(641, (d_1932_v12_).length(0))
            (d_1932_v12_)[index339_] = p1
            d_1933_v13_: _dafny.Map
            d_1933_v13_ = _dafny.Map({((d_1915_v2_).cf6 if (self).f6 else p2): (p1) not in (_dafny.Map({p1: (self).f6}))})
            index340_ = default__.safeIndex(641, (d_1932_v12_).length(0))
            rhs365_ = (self).f9
            rhs366_ = d_1933_v13_
            rhs367_ = (default__.safeModuloInt((self).f9, (0) - ((self).f9))) * ((0) - (default__.safeDivisionInt(953, -194)))
            lhs251_ = d_1932_v12_
            lhs252_ = default__.safeIndex(641, (d_1932_v12_).length(0))
            lhs253_ = self
            lhs251_[lhs252_] = rhs365_
            d_1933_v13_ = rhs366_
            lhs253_.f10 = rhs367_
            d_1927_v7_ = d_1927_v7_
            d_1934_v14_: C0
            nw334_ = C0()
            nw334_.ctor__((d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], _dafny.SeqWithoutIsStrInference([d_1931_v11_ for d_1935_i2_ in range(default__.abs(794))]))
            d_1934_v14_ = nw334_
            d_1936_v15_: C0
            nw335_ = C0()
            nw335_.ctor__(p0, d_1934_v14_.f12)
            d_1936_v15_ = nw335_
        elif source26_.is_DC6:
            d_1937___mcc_h3_ = source26_.cf6
            d_1938_cf6_ = d_1937___mcc_h3_
            (globalState).f1 = self.f10
            (globalState).f1 = self.f10
            (self).f10 = (self).f9
            index341_ = default__.safeIndex(292, (d_1911_v0_).length(0))
            (d_1911_v0_)[index341_] = not(((self).f9) < (default__.fm3((self).f6, p0, globalState)))
        elif True:
            d_1939___mcc_h4_ = source26_.cf10
            d_1940_cf10_ = d_1939___mcc_h4_
            d_1941_v16_: _dafny.Seq
            d_1941_v16_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cofnqf")), (_dafny.SeqWithoutIsStrInference([(self).f5 for d_1942_i3_ in range(default__.abs(932))])) + (p2)])
            d_1941_v16_ = d_1941_v16_
            d_1943_v17_: T0
            nw336_ = C0()
            nw336_.ctor__((d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(self).f6, (self).f6}))]) for d_1944_i4_ in range(default__.abs(748))]))
            d_1943_v17_ = nw336_
            d_1945_v18_: _dafny.Seq
            d_1945_v18_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1946_v19_: _dafny.Map
            d_1946_v19_ = _dafny.Map({d_1943_v17_: d_1945_v18_})
            d_1947_v20_: D2
            d_1947_v20_ = D2_DC8(p0)
            index342_ = default__.safeIndex(292, (d_1911_v0_).length(0))
            rhs368_ = ((self.f10) * (611)) <= (len(_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])))
            rhs369_ = ((d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))]) or ((d_1947_v20_).cf9)
            rhs370_ = ((d_1946_v19_).set(d_1943_v17_, d_1945_v18_)) | ((d_1946_v19_).set(d_1943_v17_, _dafny.SeqWithoutIsStrInference([(d_1914_v1_).cardinality])))
            lhs254_ = d_1911_v0_
            lhs255_ = default__.safeIndex(292, (d_1911_v0_).length(0))
            r0 = rhs368_
            lhs254_[lhs255_] = rhs369_
            d_1946_v19_ = rhs370_
            d_1948_v21_: _dafny.Map
            d_1948_v21_ = _dafny.Map({self.f10: self.f10})
            d_1949_v22_: _dafny.Map
            d_1949_v22_ = _dafny.Map({d_1948_v21_: (p1 if not(p0) else self.f10)})
            d_1949_v22_ = (d_1949_v22_).set(d_1948_v21_, (self.f10) + (len(d_1945_v18_)))
            d_1950_v23_: _dafny.Seq
            d_1950_v23_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1951_v24_: _dafny.Array
            nw337_ = _dafny.Array(int(0), 1)
            d_1951_v24_ = nw337_
            d_1952_v25_: _dafny.Array
            d_1952_v25_ = d_1951_v24_
            rhs371_ = (d_1950_v23_) + ((_dafny.SeqWithoutIsStrInference([p0])) + (d_1950_v23_))
            rhs372_ = (d_1915_v2_ if False else d_1915_v2_)
            rhs373_ = ((d_1952_v25_ if not(p0) else d_1952_v25_))
            d_1950_v23_ = rhs371_
            d_1915_v2_ = rhs372_
            d_1951_v24_ = rhs373_
        d_1953_v26_: _dafny.Seq
        d_1953_v26_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1954_v27_: _dafny.Seq
        d_1954_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm3(p0, (d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))], globalState)])
        d_1955_v28_: _dafny.Map
        d_1955_v28_ = _dafny.Map({d_1953_v26_: (self).fm6(d_1954_v27_, globalState)})
        d_1955_v28_ = (d_1955_v28_).set(d_1953_v26_, p0)
        d_1956_v29_: _dafny.Seq
        d_1956_v29_ = _dafny.SeqWithoutIsStrInference([p2, p2])
        (self).f10 = default__.safeDivisionInt(len((d_1956_v29_)[default__.safeIndex(self.f10, len(d_1956_v29_))]), p1)
        d_1957_v30_: C1
        nw338_ = C1()
        nw338_.ctor__()
        d_1957_v30_ = nw338_
        r0 = ((self).f6 if False else (d_1911_v0_)[default__.safeIndex(292, (d_1911_v0_).length(0))])
        r1 = p2
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        d_1958_v0_: _dafny.Map
        d_1958_v0_ = _dafny.Map({(self).f6: not((self).f6)})
        d_1959_v1_: _dafny.Map
        d_1959_v1_ = _dafny.Map({(self).f9: -162})
        d_1960_v2_: _dafny.Seq
        d_1960_v2_ = _dafny.SeqWithoutIsStrInference([d_1958_v0_, d_1958_v0_, d_1958_v0_])
        d_1961_v3_: _dafny.Array
        nw339_ = _dafny.Array(None, 14)
        nw339_[int(0)] = d_1958_v0_
        nw339_[int(1)] = d_1958_v0_
        nw339_[int(2)] = (d_1958_v0_) | (d_1958_v0_)
        nw339_[int(3)] = _dafny.Map({(self).f6: (self).f6})
        nw339_[int(4)] = _dafny.Map({not(not((self).f6)): (self).f6})
        nw339_[int(5)] = (d_1958_v0_).set((self).f6, (self).f6)
        nw339_[int(6)] = (d_1958_v0_) | (d_1958_v0_)
        nw339_[int(7)] = _dafny.Map({(self).f6: (self).f6})
        nw339_[int(8)] = default__.fm19(default__.fm0((self).f6, globalState), ((d_1959_v1_)[(self).f9] if ((self).f9) in (d_1959_v1_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nud")))), globalState)
        nw339_[int(9)] = d_1958_v0_
        nw339_[int(10)] = d_1958_v0_
        nw339_[int(11)] = (d_1958_v0_) | (d_1958_v0_)
        nw339_[int(12)] = d_1958_v0_
        nw339_[int(13)] = (d_1958_v0_) | ((d_1960_v2_)[default__.safeIndex(945, len(d_1960_v2_))])
        d_1961_v3_ = nw339_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1961_v3_).length(0)):
            d_1962_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_1962_i0_)) and ((d_1962_i0_) < ((d_1961_v3_).length(0)))):
                (d_1961_v3_)[(d_1962_i0_)] = d_1958_v0_
        d_1963_i1_: int
        d_1963_i1_ = 0
        with _dafny.label("20"):
            while (self).f6:
                with _dafny.c_label("20"):
                    if (d_1963_i1_) >= (100):
                        raise _dafny.Break("20")
                    d_1963_i1_ = (d_1963_i1_) + (1)
                    d_1964_v4_: bool
                    d_1964_v4_ = False
                    d_1964_v4_ = ((d_1958_v0_)[(self).f6] if ((self).f6) in (d_1958_v0_) else ((self).f6) and (d_1964_v4_))
                    (globalState).f1 = (self).f9
                    d_1965_v5_: _dafny.Map
                    d_1965_v5_ = _dafny.Map({(self).f6: (self).f5})
                    d_1966_v6_: _dafny.Seq
                    d_1966_v6_ = _dafny.SeqWithoutIsStrInference([d_1964_v4_])
                    d_1967_v7_: _dafny.Set
                    d_1967_v7_ = _dafny.Set({len(d_1966_v6_), (_dafny.MultiSet([d_1964_v4_, (self).f6])).cardinality})
                    d_1968_v8_: _dafny.Map
                    d_1968_v8_ = _dafny.Map({133: D1_DC4((self).f9)})
                    d_1969_v9_: _dafny.MultiSet
                    d_1969_v9_ = _dafny.MultiSet([d_1967_v7_, _dafny.Set({len(d_1968_v8_)})])
                    d_1970_v10_: _dafny.Array
                    nw340_ = _dafny.Array(int(0), 2)
                    d_1970_v10_ = nw340_
                    d_1971_v11_: _dafny.MultiSet
                    d_1971_v11_ = _dafny.MultiSet([d_1970_v10_, d_1970_v10_, d_1970_v10_, d_1970_v10_, d_1970_v10_])
                    d_1965_v5_ = (d_1965_v5_).set(((d_1969_v9_).cardinality) <= ((d_1971_v11_).cardinality), (self).f5)
                    d_1972_v12_: _dafny.Seq
                    d_1972_v12_ = _dafny.SeqWithoutIsStrInference([445, self.f10])
                    d_1973_v13_: _dafny.Seq
                    d_1973_v13_ = _dafny.SeqWithoutIsStrInference([d_1972_v12_, d_1972_v12_])
                    d_1974_v14_: C0
                    nw341_ = C0()
                    nw341_.ctor__(default__.fm0((self).f6, globalState), (d_1973_v13_) + (d_1973_v13_))
                    d_1974_v14_ = nw341_
                    pass
            pass
        if (self).f6:
            (globalState).f1 = self.f10
            d_1975_v15_: _dafny.Array
            nw342_ = _dafny.Array(int(0), 21)
            d_1975_v15_ = nw342_
            d_1976_v16_: _dafny.Map
            d_1976_v16_ = _dafny.Map({(self).f5: d_1975_v15_})
            index343_ = default__.safeIndex(106, (d_1975_v15_).length(0))
            (d_1975_v15_)[index343_] = ((self).f9) * ((self).f9)
            d_1977_v17_: bool
            d_1977_v17_ = False
            d_1978_v18_: _dafny.Map
            d_1978_v18_ = _dafny.Map({(self).f9: False})
            d_1979_v19_: _dafny.Map
            d_1979_v19_ = d_1978_v18_
            d_1980_v20_: _dafny.Map
            d_1980_v20_ = _dafny.Map({d_1979_v19_: _dafny.CodePoint('j')})
            index344_ = default__.safeIndex(106, (d_1975_v15_).length(0))
            rhs374_ = d_1976_v16_
            rhs375_ = self.f10
            rhs376_ = False
            rhs377_ = (self).fm8((d_1980_v20_) != (d_1980_v20_), (self).f6, False, globalState)
            lhs256_ = d_1975_v15_
            lhs257_ = default__.safeIndex(106, (d_1975_v15_).length(0))
            lhs258_ = self
            d_1976_v16_ = rhs374_
            lhs256_[lhs257_] = rhs375_
            d_1977_v17_ = rhs376_
            lhs258_.f10 = rhs377_
            d_1981_v21_: _dafny.Seq
            d_1981_v21_ = _dafny.SeqWithoutIsStrInference([(d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))], self.f10])
            d_1982_v22_: C0
            nw343_ = C0()
            nw343_.ctor__(d_1977_v17_, _dafny.SeqWithoutIsStrInference([d_1981_v21_]))
            d_1982_v22_ = nw343_
            d_1982_v22_ = d_1982_v22_
            if ((self).f9) < (default__.safeDivisionInt((self).f9, (d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))])):
                d_1983_v23_: str
                d_1983_v23_ = _dafny.CodePoint('a')
                d_1983_v23_ = d_1983_v23_
                d_1984_v24_: _dafny.Seq
                d_1984_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcwnadjh"))
                d_1985_v25_: _dafny.Set
                d_1985_v25_ = _dafny.Set({self.f10})
                d_1986_v26_: _dafny.Map
                d_1986_v26_ = _dafny.Map({(d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]: (self).f5})
                index345_ = default__.safeIndex(106, (d_1975_v15_).length(0))
                index346_ = default__.safeIndex(106, (d_1975_v15_).length(0))
                rhs378_ = (d_1982_v22_).f11
                rhs379_ = ((d_1984_v24_) + (d_1984_v24_) if d_1977_v17_ else d_1984_v24_)
                rhs380_ = ((d_1985_v25_) - (d_1985_v25_)).issubset(d_1985_v25_)
                rhs381_ = (0) - (default__.safeDivisionInt(len(d_1959_v1_), (d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]))
                rhs382_ = default__.safeModuloInt((0) - (len(((d_1984_v24_) + (d_1984_v24_)).set(default__.safeIndex((self).f9, len((d_1984_v24_) + (d_1984_v24_))), ((d_1986_v26_)[(d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]] if ((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]) in (d_1986_v26_) else (self).f5)))), (d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))])
                lhs259_ = d_1975_v15_
                lhs260_ = default__.safeIndex(106, (d_1975_v15_).length(0))
                lhs261_ = d_1975_v15_
                lhs262_ = default__.safeIndex(106, (d_1975_v15_).length(0))
                d_1977_v17_ = rhs378_
                d_1984_v24_ = rhs379_
                d_1977_v17_ = rhs380_
                lhs259_[lhs260_] = rhs381_
                lhs261_[lhs262_] = rhs382_
                d_1987_v27_: D2
                d_1987_v27_ = D2_DC6(d_1984_v24_)
                d_1987_v27_ = D2_DC6(d_1984_v24_)
                d_1988_v28_: _dafny.Seq
                d_1988_v28_ = _dafny.SeqWithoutIsStrInference([(d_1982_v22_).f11])
                d_1989_v29_: _dafny.Map
                d_1989_v29_ = _dafny.Map({(d_1988_v28_) + (d_1988_v28_): d_1977_v17_})
                d_1989_v29_ = (d_1989_v29_) | (d_1989_v29_)
                d_1990_v30_: _dafny.Set
                d_1990_v30_ = _dafny.Set({(self).f6})
                (globalState).f1 = (self.f10) - ((len(_dafny.Map({False: (d_1982_v22_).f11}))) - (len(d_1990_v30_)))
            elif True:
                r0 = (0) - ((0) - (default__.safeModuloInt(default__.safeDivisionInt((D1_DC4((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))])).cf4, (d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]), len(d_1982_v22_.f12))))
                (self).f10 = len(((d_1958_v0_ if not((self).f6) else _dafny.Map({(d_1982_v22_).f11: (d_1982_v22_).f11}))) | (d_1958_v0_))
                d_1991_v31_: _dafny.Seq
                d_1991_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]: not((self).f6)})])
                d_1977_v17_ = ((self).f9) == (len(d_1991_v31_))
                d_1992_v32_: _dafny.Array
                nw344_ = _dafny.Array(False, 23)
                d_1992_v32_ = nw344_
                d_1993_v33_: _dafny.Map
                d_1993_v33_ = _dafny.Map({(self).f9: d_1992_v32_})
                d_1994_v34_: _dafny.Map
                d_1994_v34_ = _dafny.Map({D1_DC4(self.f10): (d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]})
                rhs383_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "latagnlxg"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avl")))
                rhs384_ = ((d_1993_v33_)[len(d_1994_v34_)] if (len(d_1994_v34_)) in (d_1993_v33_) else d_1992_v32_)
                rhs385_ = (0) - (self.f10)
                d_1977_v17_ = rhs383_
                d_1992_v32_ = rhs384_
                r0 = rhs385_
                d_1995_v35_: _dafny.Seq
                d_1995_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywr"))
                d_1995_v35_ = (d_1995_v35_) + ((d_1995_v35_ if (self).f6 else _dafny.SeqWithoutIsStrInference([(self).f5 for d_1996_i2_ in range(default__.abs(136))])))
            d_1997_v36_: D0
            d_1997_v36_ = D0_DC3(D0_DC1())
            d_1998_v37_: _dafny.Array
            def lambda92_(d_1999_i3_):
                return (self).f6

            init53_ = lambda92_
            nw345_ = _dafny.Array(None, 6)
            for i0_53_ in range(nw345_.length(0)):
                nw345_[i0_53_] = init53_(i0_53_)
            d_1998_v37_ = nw345_
            d_2000_v38_: _dafny.Map
            d_2000_v38_ = _dafny.Map({d_1998_v37_: (self).f5})
            d_2001_v39_: D0
            d_2001_v39_ = D0_DC2(d_2000_v38_, _dafny.CodePoint('e'))
            d_2002_v40_: D0
            d_2002_v40_ = D0_DC3(d_2001_v39_)
            pat_let_tv28_ = d_2001_v39_
            def iife129_(_pat_let32_0):
                def iife130_(d_2003_dt__update__tmp_h0_):
                    def iife131_(_pat_let33_0):
                        def iife132_(d_2004_dt__update_hcf3_h0_):
                            return D0_DC3(d_2004_dt__update_hcf3_h0_)
                        return iife132_(_pat_let33_0)
                    return iife131_(pat_let_tv28_)
                return iife130_(_pat_let32_0)
            source27_ = iife129_(d_1997_v36_)
            if source27_.is_DC1:
                d_1977_v17_ = (d_1982_v22_).f11
                d_2005_v41_: D1
                d_2005_v41_ = D1_DC4(-663)
                d_2006_v42_: _dafny.Array
                nw346_ = _dafny.Array(None, 7)
                nw346_[int(0)] = d_2005_v41_
                nw346_[int(1)] = d_2005_v41_
                nw346_[int(2)] = D1_DC4(len(_dafny.Set({(self).f6})))
                nw346_[int(3)] = d_2005_v41_
                nw346_[int(4)] = d_2005_v41_
                nw346_[int(5)] = d_2005_v41_
                nw346_[int(6)] = d_2005_v41_
                d_2006_v42_ = nw346_
                index347_ = default__.safeIndex(382, (d_2006_v42_).length(0))
                (d_2006_v42_)[index347_] = d_2005_v41_
                index348_ = default__.safeIndex(382, (d_2006_v42_).length(0))
                (d_2006_v42_)[index348_] = d_2005_v41_
                d_2007_v43_: _dafny.Seq
                d_2007_v43_ = _dafny.SeqWithoutIsStrInference([d_1975_v15_, d_1975_v15_])
                d_2008_v44_: _dafny.Seq
                d_2008_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qucb"))
                rhs386_ = (0) - ((d_1981_v21_)[default__.safeIndex((self).f9, len(d_1981_v21_))])
                rhs387_ = ((self).f9) - ((0) - ((self.f10) - ((self).f9)))
                rhs388_ = _dafny.SeqWithoutIsStrInference([d_1975_v15_])
                rhs389_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhwkuxs"))
                lhs263_ = self
                lhs264_ = globalState
                lhs263_.f10 = rhs386_
                lhs264_.f1 = rhs387_
                d_2007_v43_ = rhs388_
                d_2008_v44_ = rhs389_
                d_2009_v45_: _dafny.MultiSet
                d_2009_v45_ = _dafny.MultiSet([default__.fm0((d_1982_v22_).f11, globalState), (d_1982_v22_).f11])
                d_2010_v46_: _dafny.Seq
                d_2010_v46_ = _dafny.SeqWithoutIsStrInference([(self).f6, not(d_1977_v17_)])
                d_2011_v47_: _dafny.Array
                nw347_ = _dafny.Array(None, 18)
                nw347_[int(0)] = _dafny.MultiSet([(d_1982_v22_).f11, (self).f6])
                nw347_[int(1)] = (d_2009_v45_) | (d_2009_v45_)
                nw347_[int(2)] = (_dafny.MultiSet([(d_1982_v22_).f11, (d_1982_v22_).f11])) - (d_2009_v45_)
                nw347_[int(3)] = d_2009_v45_
                nw347_[int(4)] = (d_2009_v45_) - (_dafny.MultiSet(d_2010_v46_))
                nw347_[int(5)] = d_2009_v45_
                nw347_[int(6)] = _dafny.MultiSet([(self).f6, False, (d_1982_v22_).f11, not((self).f6), (d_1982_v22_).f11])
                nw347_[int(7)] = (d_2009_v45_) | (_dafny.MultiSet(d_2010_v46_))
                nw347_[int(8)] = d_2009_v45_
                nw347_[int(9)] = d_2009_v45_
                nw347_[int(10)] = d_2009_v45_
                nw347_[int(11)] = d_2009_v45_
                nw347_[int(12)] = (d_2009_v45_) | (d_2009_v45_)
                nw347_[int(13)] = d_2009_v45_
                nw347_[int(14)] = d_2009_v45_
                nw347_[int(15)] = d_2009_v45_
                nw347_[int(16)] = d_2009_v45_
                nw347_[int(17)] = (_dafny.MultiSet([(self).f6])) | ((_dafny.MultiSet([(self).f6, (d_1982_v22_).f11, (self).f6])).set((d_1982_v22_).f11, default__.abs(((d_2006_v42_)[default__.safeIndex(382, (d_2006_v42_).length(0))]).cf4)))
                d_2011_v47_ = nw347_
                index349_ = default__.safeIndex(585, (d_2011_v47_).length(0))
                (d_2011_v47_)[index349_] = (d_2009_v45_) | (d_2009_v45_)
                index350_ = default__.safeIndex(585, (d_2011_v47_).length(0))
                (d_2011_v47_)[index350_] = d_2009_v45_
            elif source27_.is_DC2:
                d_2012___mcc_h0_ = source27_.cf1
                d_2013___mcc_h1_ = source27_.cf2
                d_2014_cf2_ = d_2013___mcc_h1_
                d_2015_cf1_ = d_2012___mcc_h0_
                d_2016_v48_: C1
                nw348_ = C1()
                nw348_.ctor__()
                d_2016_v48_ = nw348_
                d_1958_v0_ = (d_1958_v0_).set(((self).f9) > (self.f10), (d_1982_v22_).f11)
                d_1977_v17_ = ((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]) != (default__.safeModuloInt((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))], (self).f9))
                d_1977_v17_ = (d_1982_v22_).f11
            elif source27_.is_DC0:
                d_2017___mcc_h2_ = source27_.cf0
                d_2018_cf0_ = d_2017___mcc_h2_
                d_2019_v49_: _dafny.Seq
                d_2019_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psooqryel"))
                d_2018_cf0_ = (d_2019_v49_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udq")))
                d_2020_v50_: _dafny.Seq
                d_2020_v50_ = _dafny.SeqWithoutIsStrInference([d_1975_v15_])
                rhs390_ = d_2019_v49_
                rhs391_ = default__.safeModuloInt(-128, (self).f9)
                rhs392_ = True
                rhs393_ = d_2018_cf0_
                rhs394_ = (d_2020_v50_) < (d_2020_v50_)
                lhs265_ = globalState
                d_2019_v49_ = rhs390_
                lhs265_.f1 = rhs391_
                d_1977_v17_ = rhs392_
                d_1977_v17_ = rhs393_
                d_2018_cf0_ = rhs394_
                d_2021_v51_: _dafny.Seq
                d_2021_v51_ = _dafny.SeqWithoutIsStrInference([d_1977_v17_])
                d_1977_v17_ = (True) or ((d_2021_v51_)[default__.safeIndex((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))], len(d_2021_v51_))])
                d_1977_v17_ = not(((d_1982_v22_).f11) == (((d_1982_v22_).f11) not in (_dafny.MultiSet(d_2021_v51_))))
            elif True:
                d_2022___mcc_h3_ = source27_.cf3
                d_2023_cf3_ = d_2022___mcc_h3_
                d_2024_v52_: _dafny.Array
                d_2024_v52_ = d_1975_v15_
                d_2025_v53_: _dafny.Set
                d_2025_v53_ = _dafny.Set({d_2024_v52_, d_2024_v52_, d_2024_v52_, d_2024_v52_})
                d_2025_v53_ = d_2025_v53_
                d_2026_v54_: D0
                d_2026_v54_ = D0_DC0((d_1982_v22_).f11)
                d_2027_v55_: _dafny.Set
                d_2027_v55_ = _dafny.Set({d_2026_v54_, D0_DC0(d_1977_v17_)})
                d_1977_v17_ = (d_2027_v55_) != (d_2027_v55_)
                d_2028_v56_: D0
                d_2028_v56_ = D0_DC2(d_2000_v38_, (self).f5)
                d_2029_v57_: _dafny.Seq
                d_2029_v57_ = _dafny.SeqWithoutIsStrInference([d_2028_v56_])
                d_2030_v58_: T4
                nw349_ = C2()
                nw349_.ctor__((self).f5, (self).f6)
                d_2030_v58_ = nw349_
                d_2031_v59_: _dafny.MultiSet
                d_2031_v59_ = _dafny.MultiSet([d_2030_v58_])
                d_2032_v60_: _dafny.Map
                d_2032_v60_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2029_v57_]): d_2031_v59_})
                d_2033_v61_: _dafny.Seq
                d_2033_v61_ = _dafny.SeqWithoutIsStrInference([d_2029_v57_, d_2029_v57_])
                d_2032_v60_ = (d_2032_v60_).set((d_2033_v61_) + (d_2033_v61_), (d_2031_v59_).intersection(d_2031_v59_))
                d_2034_v62_: _dafny.Seq
                d_2034_v62_ = _dafny.SeqWithoutIsStrInference([True])
                d_1958_v0_ = (d_1958_v0_).set((d_2034_v62_)[default__.safeIndex(self.f10, len(d_2034_v62_))], ((d_1975_v15_)[default__.safeIndex(106, (d_1975_v15_).length(0))]) >= (self.f10))
        elif True:
            d_2035_v63_: bool
            d_2035_v63_ = True
            d_2036_v64_: _dafny.Set
            d_2036_v64_ = _dafny.Set({not(d_2035_v63_), d_2035_v63_})
            d_2037_v65_: _dafny.Map
            d_2037_v65_ = _dafny.Map({(self).f9: (d_2036_v64_).isdisjoint(d_2036_v64_)})
            d_2035_v63_ = ((d_2037_v65_)[(self).f9] if ((self).f9) in (d_2037_v65_) else (self).f6)
            d_2035_v63_ = d_2035_v63_
            d_2038_v66_: _dafny.Seq
            d_2038_v66_ = _dafny.SeqWithoutIsStrInference([False])
            d_2039_v67_: _dafny.Seq
            d_2039_v67_ = d_2038_v66_
            d_2040_v68_: _dafny.Seq
            d_2040_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekfgutbiw"))
            d_2041_v69_: _dafny.Seq
            d_2041_v69_ = _dafny.SeqWithoutIsStrInference([len(d_2040_v68_)])
            rhs395_ = (d_2039_v67_)
            rhs396_ = ((d_2036_v64_).issubset(d_2036_v64_)) == ((d_2038_v66_)[default__.safeIndex((0) - (self.f10), len(d_2038_v66_))])
            rhs397_ = (default__.safeDivisionInt(652, (0) - ((self).f9))) not in (d_2041_v69_)
            rhs398_ = (self).f6
            d_2038_v66_ = rhs395_
            d_2035_v63_ = rhs396_
            d_2035_v63_ = rhs397_
            d_2035_v63_ = rhs398_
            d_2042_v70_: _dafny.Array
            def lambda93_(d_2043_v65_):
                def lambda94_(d_2044_i4_):
                    return ((d_2043_v65_)[len(_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5]))] if (len(_dafny.SeqWithoutIsStrInference([(self).f5, (self).f5, (self).f5]))) in (d_2043_v65_) else True)

                return lambda94_

            init54_ = lambda93_(d_2037_v65_)
            nw350_ = _dafny.Array(None, 8)
            for i0_54_ in range(nw350_.length(0)):
                nw350_[i0_54_] = init54_(i0_54_)
            d_2042_v70_ = nw350_
            d_2042_v70_ = d_2042_v70_
            d_2045_v71_: _dafny.Array
            nw351_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_2045_v71_ = nw351_
            index351_ = default__.safeIndex(723, (d_2045_v71_).length(0))
            (d_2045_v71_)[index351_] = d_2042_v70_
            index352_ = default__.safeIndex(723, (d_2045_v71_).length(0))
            rhs399_ = (len(d_2040_v68_)) >= (self.f10)
            rhs400_ = self.f10
            rhs401_ = (self).f9
            rhs402_ = d_2042_v70_
            rhs403_ = 848
            lhs266_ = self
            lhs267_ = globalState
            lhs268_ = d_2045_v71_
            lhs269_ = default__.safeIndex(723, (d_2045_v71_).length(0))
            lhs270_ = globalState
            d_2035_v63_ = rhs399_
            lhs266_.f10 = rhs400_
            lhs267_.f1 = rhs401_
            lhs268_[lhs269_] = rhs402_
            lhs270_.f1 = rhs403_
        d_2046_v72_: bool
        d_2046_v72_ = False
        d_2046_v72_ = d_2046_v72_
        hi11_ = (0) - (default__.safeDivisionInt(self.f10, (self).f9))
        for d_2047_i5_ in range(-364, hi11_):
            d_2048_v73_: _dafny.Array
            nw352_ = _dafny.Array(False, 7)
            d_2048_v73_ = nw352_
            index353_ = default__.safeIndex(520, (d_2048_v73_).length(0))
            (d_2048_v73_)[index353_] = (self).f6
            index354_ = default__.safeIndex(520, (d_2048_v73_).length(0))
            (d_2048_v73_)[index354_] = not(d_2046_v72_)
            d_2049_v74_: C1
            nw353_ = C1()
            nw353_.ctor__()
            d_2049_v74_ = nw353_
            d_2050_v75_: _dafny.MultiSet
            d_2050_v75_ = _dafny.MultiSet([54])
            d_2051_v76_: _dafny.Seq
            d_2051_v76_ = _dafny.SeqWithoutIsStrInference([True])
            d_2052_v77_: _dafny.Map
            d_2052_v77_ = _dafny.Map({self.f10: d_2050_v75_})
            d_2053_v78_: _dafny.Seq
            d_2053_v78_ = _dafny.SeqWithoutIsStrInference([(self).f9, self.f10, (self).f9])
            d_2054_v79_: _dafny.Seq
            d_2054_v79_ = _dafny.SeqWithoutIsStrInference([d_2053_v78_, d_2053_v78_])
            d_2055_v80_: C0
            nw354_ = C0()
            nw354_.ctor__((d_2049_v74_).fm6((d_2053_v78_).set(default__.safeIndex(self.f10, len(d_2053_v78_)), d_2047_i5_), globalState), d_2054_v79_)
            d_2055_v80_ = nw354_
            d_2056_v81_: _dafny.MultiSet
            d_2056_v81_ = _dafny.MultiSet([d_2055_v80_, d_2055_v80_])
            index355_ = default__.safeIndex(520, (d_2048_v73_).length(0))
            rhs404_ = ((d_2052_v77_)[d_2047_i5_] if (d_2047_i5_) in (d_2052_v77_) else d_2050_v75_)
            rhs405_ = ((d_2056_v81_)[d_2055_v80_] if (d_2055_v80_) in (d_2056_v81_) else (self).f9)
            rhs406_ = d_2051_v76_
            rhs407_ = (self).fm6((d_2053_v78_) + (d_2053_v78_), globalState)
            lhs271_ = d_2048_v73_
            lhs272_ = default__.safeIndex(520, (d_2048_v73_).length(0))
            d_2050_v75_ = rhs404_
            r0 = rhs405_
            d_2051_v76_ = rhs406_
            lhs271_[lhs272_] = rhs407_
            index356_ = default__.safeIndex(520, (d_2048_v73_).length(0))
            (d_2048_v73_)[index356_] = d_2046_v72_
        d_2057_i6_: int
        d_2057_i6_ = 0
        with _dafny.label("21"):
            while default__.fm0(not(d_2046_v72_), globalState):
                with _dafny.c_label("21"):
                    if (d_2057_i6_) >= (100):
                        raise _dafny.Break("21")
                    d_2057_i6_ = (d_2057_i6_) + (1)
                    if (self).f6:
                        d_2058_v82_: _dafny.Set
                        d_2058_v82_ = _dafny.Set({(self).f5, (self).f5, (self).f5, (self).f5})
                        d_2059_v83_: _dafny.Array
                        def lambda95_(d_2060_i7_):
                            return (d_2060_i7_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukbkpwm"))))

                        init55_ = lambda95_
                        nw355_ = _dafny.Array(None, 10)
                        for i0_55_ in range(nw355_.length(0)):
                            nw355_[i0_55_] = init55_(i0_55_)
                        d_2059_v83_ = nw355_
                        index357_ = default__.safeIndex(862, (d_2059_v83_).length(0))
                        (d_2059_v83_)[index357_] = self.f10
                        d_2061_v84_: D2
                        d_2061_v84_ = D2_DC8(d_2046_v72_)
                        d_2062_v85_: _dafny.MultiSet
                        d_2062_v85_ = _dafny.MultiSet([(self).f9])
                        d_2063_v86_: _dafny.MultiSet
                        d_2063_v86_ = _dafny.MultiSet([(self).f9, (self).f9, self.f10, (d_2062_v85_).cardinality])
                        d_2064_v87_: _dafny.MultiSet
                        d_2064_v87_ = _dafny.MultiSet([self.f10, (d_2062_v85_).cardinality])
                        d_2065_v88_: D6
                        d_2065_v88_ = D6_DC14(d_2046_v72_)
                        d_2066_v89_: _dafny.Seq
                        d_2066_v89_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "caicnmft"))
                        d_2067_v90_: D2
                        d_2067_v90_ = D2_DC7(self.f10, len(d_2066_v89_))
                        d_2068_v91_: _dafny.Seq
                        d_2068_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2067_v90_).cf7, (self).f9])])
                        d_2069_v92_: C0
                        nw356_ = C0()
                        nw356_.ctor__((d_2065_v88_).cf15, d_2068_v91_)
                        d_2069_v92_ = nw356_
                        d_2070_v93_: _dafny.Map
                        d_2070_v93_ = _dafny.Map({d_2069_v92_: (self).f9})
                        d_2071_v94_: _dafny.Seq
                        d_2071_v94_ = _dafny.SeqWithoutIsStrInference([d_2046_v72_, (self).f6, False])
                        d_2072_v95_: _dafny.Array
                        nw357_ = _dafny.Array(None, 15)
                        nw357_[int(0)] = not ((self).f6) or (False)
                        nw357_[int(1)] = (self).f6
                        nw357_[int(2)] = (d_2061_v84_).cf9
                        nw357_[int(3)] = (d_2062_v85_).ispropersubset(d_2063_v86_)
                        nw357_[int(4)] = d_2046_v72_
                        nw357_[int(5)] = d_2046_v72_
                        nw357_[int(6)] = (self).f6
                        nw357_[int(7)] = ((self).f9) < (len(d_2070_v93_))
                        nw357_[int(8)] = (d_2069_v92_).f11
                        nw357_[int(9)] = (self).f6
                        nw357_[int(10)] = not((False) and (False))
                        nw357_[int(11)] = (self).f6
                        nw357_[int(12)] = (d_2069_v92_).f11
                        nw357_[int(13)] = (d_2071_v94_)[default__.safeIndex(self.f10, len(d_2071_v94_))]
                        nw357_[int(14)] = False
                        d_2072_v95_ = nw357_
                        index358_ = default__.safeIndex(875, (d_2072_v95_).length(0))
                        (d_2072_v95_)[index358_] = (self).f6
                        d_2073_v96_: _dafny.Seq
                        d_2073_v96_ = _dafny.SeqWithoutIsStrInference([(self).f9, len(d_2066_v89_), len(d_2071_v94_)])
                        index359_ = default__.safeIndex(862, (d_2059_v83_).length(0))
                        index360_ = default__.safeIndex(875, (d_2072_v95_).length(0))
                        rhs408_ = d_2058_v82_
                        rhs409_ = (self).f9
                        rhs410_ = (self.f10) <= ((self.f10 if (d_2069_v92_).f11 else self.f10))
                        rhs411_ = d_2046_v72_
                        rhs412_ = (d_2073_v96_) + (_dafny.SeqWithoutIsStrInference([len(d_2066_v89_) for d_2074_i8_ in range(default__.abs(975))]))
                        lhs273_ = d_2059_v83_
                        lhs274_ = default__.safeIndex(862, (d_2059_v83_).length(0))
                        lhs275_ = d_2072_v95_
                        lhs276_ = default__.safeIndex(875, (d_2072_v95_).length(0))
                        d_2058_v82_ = rhs408_
                        lhs273_[lhs274_] = rhs409_
                        d_2046_v72_ = rhs410_
                        lhs275_[lhs276_] = rhs411_
                        d_2073_v96_ = rhs412_
                        (globalState).f1 = ((self).f9) - ((d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))])
                        d_2075_v97_: _dafny.Array
                        nw358_ = _dafny.Array(_dafny.Array(None, 0), 2)
                        d_2075_v97_ = nw358_
                        d_2076_v98_: _dafny.Map
                        d_2076_v98_ = _dafny.Map({self.f10: d_2075_v97_})
                        d_2076_v98_ = (d_2076_v98_).set((0) - ((d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))]), d_2075_v97_)
                        d_2077_v99_: C1
                        nw359_ = C1()
                        nw359_.ctor__()
                        d_2077_v99_ = nw359_
                        d_2078_v100_: _dafny.MultiSet
                        d_2078_v100_ = _dafny.MultiSet([d_2072_v95_, d_2072_v95_, d_2072_v95_, d_2072_v95_])
                        d_2079_v101_: _dafny.Map
                        d_2079_v101_ = _dafny.Map({_dafny.CodePoint('j'): d_2078_v100_})
                        d_2080_v102_: _dafny.Map
                        d_2080_v102_ = _dafny.Map({d_2046_v72_: d_2072_v95_})
                        d_2081_v103_: _dafny.Map
                        d_2081_v103_ = _dafny.Map({(d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))]: d_2078_v100_})
                        d_2082_v104_: _dafny.Seq
                        d_2082_v104_ = _dafny.SeqWithoutIsStrInference([d_2078_v100_])
                        d_2083_v105_: _dafny.Array
                        nw360_ = _dafny.Array(None, 28)
                        nw360_[int(0)] = d_2078_v100_
                        nw360_[int(1)] = (_dafny.MultiSet([d_2072_v95_, d_2072_v95_])).intersection(d_2078_v100_)
                        nw360_[int(2)] = _dafny.MultiSet([d_2072_v95_])
                        nw360_[int(3)] = ((d_2079_v101_)[(self).f5] if ((self).f5) in (d_2079_v101_) else d_2078_v100_)
                        nw360_[int(4)] = (d_2078_v100_).set(d_2072_v95_, default__.abs((self).f9))
                        nw360_[int(5)] = d_2078_v100_
                        nw360_[int(6)] = _dafny.MultiSet([d_2072_v95_, d_2072_v95_, d_2072_v95_, d_2072_v95_, d_2072_v95_])
                        nw360_[int(7)] = d_2078_v100_
                        nw360_[int(8)] = (d_2078_v100_).intersection(_dafny.MultiSet([d_2072_v95_, d_2072_v95_, ((d_2080_v102_)[not(not((self).f6))] if (not(not((self).f6))) in (d_2080_v102_) else d_2072_v95_), d_2072_v95_]))
                        nw360_[int(9)] = ((d_2081_v103_)[(self).f9] if ((self).f9) in (d_2081_v103_) else d_2078_v100_)
                        nw360_[int(10)] = (d_2082_v104_)[default__.safeIndex((d_2077_v99_).fm4(d_2071_v94_, self.f10, (d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))], _dafny.MultiSet([-337]), globalState), len(d_2082_v104_))]
                        nw360_[int(11)] = d_2078_v100_
                        nw360_[int(12)] = (d_2078_v100_) - (d_2078_v100_)
                        nw360_[int(13)] = (d_2078_v100_) - (d_2078_v100_)
                        nw360_[int(14)] = d_2078_v100_
                        nw360_[int(15)] = d_2078_v100_
                        nw360_[int(16)] = ((d_2081_v103_)[(d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))]] if ((d_2059_v83_)[default__.safeIndex(862, (d_2059_v83_).length(0))]) in (d_2081_v103_) else d_2078_v100_)
                        nw360_[int(17)] = (((d_2081_v103_)[(self).f9] if ((self).f9) in (d_2081_v103_) else d_2078_v100_)).intersection(_dafny.MultiSet([d_2072_v95_]))
                        nw360_[int(18)] = _dafny.MultiSet([d_2072_v95_])
                        nw360_[int(19)] = _dafny.MultiSet([d_2072_v95_])
                        nw360_[int(20)] = d_2078_v100_
                        nw360_[int(21)] = d_2078_v100_
                        nw360_[int(22)] = (d_2078_v100_) - (_dafny.MultiSet([d_2072_v95_, d_2072_v95_]))
                        nw360_[int(23)] = _dafny.MultiSet([d_2072_v95_, d_2072_v95_, d_2072_v95_, d_2072_v95_])
                        nw360_[int(24)] = d_2078_v100_
                        nw360_[int(25)] = d_2078_v100_
                        nw360_[int(26)] = (d_2078_v100_) | (d_2078_v100_)
                        nw360_[int(27)] = d_2078_v100_
                        d_2083_v105_ = nw360_
                        index361_ = default__.safeIndex(236, (d_2083_v105_).length(0))
                        (d_2083_v105_)[index361_] = d_2078_v100_
                        index362_ = default__.safeIndex(236, (d_2083_v105_).length(0))
                        (d_2083_v105_)[index362_] = (d_2078_v100_).set(d_2072_v95_, default__.abs(len((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2084_i9_ in range(default__.abs(739))])) + (d_2066_v89_))))
                    elif True:
                        d_2085_v106_: _dafny.Map
                        d_2085_v106_ = _dafny.Map({(self).f6: (self).f5})
                        d_2085_v106_ = (d_2085_v106_).set(True, _dafny.CodePoint('s'))
                        d_2046_v72_ = (self).f6
                        d_2086_v107_: str
                        d_2086_v107_ = _dafny.CodePoint('e')
                        d_2086_v107_ = _dafny.CodePoint('t')
                        (self).f10 = (self).f9
                        r0 = self.f10
                    (globalState).f1 = self.f10
                    r0 = -304
                    d_2087_v108_: _dafny.Array
                    nw361_ = _dafny.Array(None, 26)
                    d_2087_v108_ = nw361_
                    d_2088_v109_: _dafny.Map
                    d_2088_v109_ = _dafny.Map({128: not(d_2046_v72_)})
                    d_2089_v110_: T4
                    nw362_ = C2()
                    nw362_.ctor__((self).f5, not(((d_2088_v109_)[(self).f9] if ((self).f9) in (d_2088_v109_) else (self).f6)))
                    d_2089_v110_ = nw362_
                    index363_ = default__.safeIndex(534, (d_2087_v108_).length(0))
                    (d_2087_v108_)[index363_] = d_2089_v110_
                    d_2090_v111_: _dafny.Map
                    d_2090_v111_ = _dafny.Map({((self).f9 if not(d_2046_v72_) else -367): d_2089_v110_})
                    index364_ = default__.safeIndex(534, (d_2087_v108_).length(0))
                    (d_2087_v108_)[index364_] = ((d_2090_v111_)[(self).f9] if ((self).f9) in (d_2090_v111_) else d_2089_v110_)
                    pass
            pass
        r0 = self.f10
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_2091_v0_: C1
        nw363_ = C1()
        nw363_.ctor__()
        d_2091_v0_ = nw363_
        d_2092_v1_: _dafny.Seq
        d_2092_v1_ = _dafny.SeqWithoutIsStrInference([self.f10])
        d_2093_v2_: _dafny.Array
        nw364_ = _dafny.Array(None, 11)
        nw364_[int(0)] = self.f10
        nw364_[int(1)] = ((self).f9) + (len(d_2092_v1_))
        nw364_[int(2)] = (self).f9
        nw364_[int(3)] = (self).f9
        nw364_[int(4)] = (0) - (self.f10)
        nw364_[int(5)] = self.f10
        nw364_[int(6)] = (self.f10) + ((self).f9)
        nw364_[int(7)] = 37
        nw364_[int(8)] = self.f10
        nw364_[int(9)] = self.f10
        nw364_[int(10)] = (self).f9
        d_2093_v2_ = nw364_
        index365_ = default__.safeIndex(896, (d_2093_v2_).length(0))
        (d_2093_v2_)[index365_] = default__.safeModuloInt((0) - ((0) - ((p1).cardinality)), self.f10)
        d_2094_v3_: bool
        d_2094_v3_ = True
        d_2095_v4_: _dafny.Array
        nw365_ = _dafny.Array(False, 15)
        d_2095_v4_ = nw365_
        d_2096_v5_: str
        d_2096_v5_ = _dafny.CodePoint('a')
        d_2097_v6_: _dafny.Seq
        d_2097_v6_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
        d_2098_v7_: _dafny.Seq
        d_2098_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({self.f10: len(d_2097_v6_)})])
        d_2099_v8_: _dafny.Map
        d_2099_v8_ = _dafny.Map({(d_2091_v0_).fm14(globalState): (self).f5})
        d_2100_v9_: D1
        d_2100_v9_ = D1_DC4(self.f10)
        index366_ = default__.safeIndex(896, (d_2093_v2_).length(0))
        def iife133_(_pat_let34_0):
            def iife134_(d_2101_dt__update__tmp_h1_):
                def iife135_(_pat_let35_0):
                    def iife136_(d_2102_dt__update_hcf4_h1_):
                        return D1_DC4(d_2102_dt__update_hcf4_h1_)
                    return iife136_(_pat_let35_0)
                return iife135_((self).f9)
            return iife134_(_pat_let34_0)
        def iife137_(_pat_let36_0):
            def iife138_(d_2103_dt__update__tmp_h0_):
                def iife139_(_pat_let37_0):
                    def iife140_(d_2104_dt__update_hcf4_h0_):
                        return D1_DC4(d_2104_dt__update_hcf4_h0_)
                    return iife140_(_pat_let37_0)
                return iife139_((self).f9)
            return iife138_(_pat_let36_0)
        rhs413_ = len((d_2098_v7_) + (d_2098_v7_))
        rhs414_ = True
        rhs415_ = d_2095_v4_
        rhs416_ = ((d_2099_v8_)[(iife133_(d_2100_v9_)).cf4] if ((iife137_(d_2100_v9_)).cf4) in (d_2099_v8_) else _dafny.CodePoint('b'))
        lhs277_ = d_2093_v2_
        lhs278_ = default__.safeIndex(896, (d_2093_v2_).length(0))
        lhs277_[lhs278_] = rhs413_
        d_2094_v3_ = rhs414_
        d_2095_v4_ = rhs415_
        d_2096_v5_ = rhs416_
        d_2094_v3_ = (((self).f9) * ((self).f9)) != (((0) - ((self).f9)) * (580))
        d_2105_v10_: _dafny.Seq
        d_2105_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrsms"))
        if ((d_2105_v10_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqhcfc")))) < (default__.fm1((self).f9, p2, 338, globalState)):
            if False:
                d_2106_v11_: _dafny.Array
                def lambda96_(d_2107_i0_):
                    return _dafny.Set({D0_DC1(), D0_DC1(), D0_DC1(), D0_DC1()})

                init56_ = lambda96_
                nw366_ = _dafny.Array(None, 4)
                for i0_56_ in range(nw366_.length(0)):
                    nw366_[i0_56_] = init56_(i0_56_)
                d_2106_v11_ = nw366_
                d_2108_v12_: D0
                d_2108_v12_ = D0_DC1()
                d_2109_v13_: _dafny.Set
                d_2109_v13_ = _dafny.Set({d_2108_v12_, d_2108_v12_})
                index367_ = default__.safeIndex(901, (d_2106_v11_).length(0))
                (d_2106_v11_)[index367_] = d_2109_v13_
                index368_ = default__.safeIndex(901, (d_2106_v11_).length(0))
                rhs417_ = (d_2109_v13_)
                rhs418_ = (self).f9
                rhs419_ = (d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]
                lhs279_ = d_2106_v11_
                lhs280_ = default__.safeIndex(901, (d_2106_v11_).length(0))
                lhs281_ = globalState
                lhs282_ = globalState
                lhs279_[lhs280_] = rhs417_
                lhs281_.f1 = rhs418_
                lhs282_.f1 = rhs419_
                (self).f10 = (default__.fm3(False, (self).f6, globalState)) * (self.f10)
                d_2110_v14_: _dafny.Map
                d_2110_v14_ = _dafny.Map({default__.fm27(d_2105_v10_, not(not(p2)), True, globalState): (self).f9})
                d_2111_v15_: _dafny.Map
                d_2111_v15_ = _dafny.Map({(d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]: len(_dafny.SeqWithoutIsStrInference([d_2093_v2_]))})
                d_2112_v16_: _dafny.Map
                d_2112_v16_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2096_v5_ for d_2113_i1_ in range(default__.abs(40))])): self.f10})
                d_2114_v17_: D2
                d_2114_v17_ = D2_DC8(d_2094_v3_)
                d_2110_v14_ = (d_2110_v14_).set(d_2112_v16_, len((default__.fm1(len(_dafny.Set({(self).f6})), (d_2114_v17_).cf9, (self).f9, globalState)) + (d_2105_v10_)))
                d_2115_v18_: C2
                nw367_ = C2()
                nw367_.ctor__((self).f5, d_2094_v3_)
                d_2115_v18_ = nw367_
                d_2116_v19_: _dafny.Array
                nw368_ = _dafny.Array(None, 29)
                nw368_[int(0)] = d_2115_v18_
                nw368_[int(1)] = d_2115_v18_
                nw368_[int(2)] = d_2115_v18_
                nw368_[int(3)] = d_2115_v18_
                nw368_[int(4)] = d_2115_v18_
                nw368_[int(5)] = d_2115_v18_
                nw368_[int(6)] = d_2115_v18_
                nw368_[int(7)] = d_2115_v18_
                nw368_[int(8)] = d_2115_v18_
                nw368_[int(9)] = d_2115_v18_
                nw368_[int(10)] = d_2115_v18_
                nw368_[int(11)] = d_2115_v18_
                nw368_[int(12)] = d_2115_v18_
                nw368_[int(13)] = d_2115_v18_
                nw368_[int(14)] = d_2115_v18_
                nw368_[int(15)] = d_2115_v18_
                nw368_[int(16)] = d_2115_v18_
                nw368_[int(17)] = d_2115_v18_
                nw368_[int(18)] = d_2115_v18_
                nw368_[int(19)] = d_2115_v18_
                nw368_[int(20)] = d_2115_v18_
                nw368_[int(21)] = d_2115_v18_
                nw368_[int(22)] = d_2115_v18_
                nw368_[int(23)] = d_2115_v18_
                nw368_[int(24)] = d_2115_v18_
                nw368_[int(25)] = d_2115_v18_
                nw368_[int(26)] = d_2115_v18_
                nw368_[int(27)] = d_2115_v18_
                nw368_[int(28)] = d_2115_v18_
                d_2116_v19_ = nw368_
                d_2117_v20_: _dafny.Seq
                d_2117_v20_ = _dafny.SeqWithoutIsStrInference([d_2116_v19_])
                d_2118_v21_: _dafny.Map
                d_2118_v21_ = _dafny.Map({d_2095_v4_: (d_2117_v20_)[default__.safeIndex((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], len(d_2117_v20_))]})
                d_2119_v22_: _dafny.Map
                d_2119_v22_ = _dafny.Map({p2: ((_dafny.Map({d_2095_v4_: d_2116_v19_})).set(d_2095_v4_, d_2116_v19_)).set(d_2095_v4_, d_2116_v19_)})
                d_2120_v23_: _dafny.Seq
                d_2120_v23_ = d_2097_v6_
                d_2121_v24_: _dafny.Set
                d_2121_v24_ = _dafny.Set({d_2097_v6_, d_2097_v6_, d_2097_v6_, ((d_2120_v23_)).set(default__.safeIndex((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], len((d_2120_v23_))), (self).f6)})
                d_2118_v21_ = ((d_2119_v22_)[(_dafny.Set({d_2097_v6_})).issubset(d_2121_v24_)] if ((_dafny.Set({d_2097_v6_})).issubset(d_2121_v24_)) in (d_2119_v22_) else d_2118_v21_)
                d_2122_v25_: _dafny.Map
                d_2122_v25_ = _dafny.Map({self.f10: not(d_2094_v3_)})
                d_2123_v26_: _dafny.Set
                d_2123_v26_ = _dafny.Set({d_2122_v25_, d_2122_v25_, d_2122_v25_, d_2122_v25_})
                d_2124_v28_: _dafny.Seq
                d_2124_v28_ = _dafny.SeqWithoutIsStrInference([d_2123_v26_, d_2123_v26_])
                def iife141_():
                    coll65_ = _dafny.Set()
                    compr_65_: _dafny.Map
                    for compr_65_ in (d_2123_v26_).Elements:
                        d_2125_v27_: _dafny.Map = compr_65_
                        if (d_2125_v27_) in (d_2123_v26_):
                            coll65_ = coll65_.union(_dafny.Set([d_2125_v27_]))
                    return _dafny.Set(coll65_)
                d_2123_v26_ = (d_2123_v26_) | ((iife141_()
                ).intersection((d_2124_v28_)[default__.safeIndex(self.f10, len(d_2124_v28_))]))
            elif True:
                index369_ = default__.safeIndex(896, (d_2093_v2_).length(0))
                (d_2093_v2_)[index369_] = len(d_2097_v6_)
                d_2094_v3_ = p0
                d_2094_v3_ = (self).f6
                (globalState).f1 = ((0) - ((self).fm7(True, globalState))) - (759)
                d_2126_v29_: _dafny.Seq
                d_2126_v29_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f9 for d_2127_i2_ in range(default__.abs(294))])])
                d_2128_v30_: T0
                nw369_ = C0()
                nw369_.ctor__(p0, d_2126_v29_)
                d_2128_v30_ = nw369_
                d_2129_v31_: _dafny.Seq
                d_2129_v31_ = _dafny.SeqWithoutIsStrInference([(d_2128_v30_ if p0 else d_2128_v30_), d_2128_v30_, d_2128_v30_, d_2128_v30_])
                d_2129_v31_ = ((d_2129_v31_) + (_dafny.SeqWithoutIsStrInference([d_2128_v30_]))) + (d_2129_v31_)
            index370_ = default__.safeIndex(896, (d_2093_v2_).length(0))
            (d_2093_v2_)[index370_] = (0) - (self.f10)
            d_2094_v3_ = True
            d_2130_v32_: _dafny.Seq
            d_2130_v32_ = _dafny.SeqWithoutIsStrInference([d_2092_v1_])
            d_2131_v33_: T0
            nw370_ = C0()
            nw370_.ctor__(p2, d_2130_v32_)
            d_2131_v33_ = nw370_
            d_2131_v33_ = d_2131_v33_
            if ((self).f9) != ((0) - ((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))])):
                d_2132_v34_: _dafny.Set
                d_2132_v34_ = _dafny.Set({(d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], self.f10})
                index371_ = default__.safeIndex(896, (d_2093_v2_).length(0))
                (d_2093_v2_)[index371_] = (self).fm8(p2, d_2094_v3_, (d_2132_v34_).isdisjoint(d_2132_v34_), globalState)
                (globalState).f4 = p1
                d_2133_v35_: _dafny.Array
                nw371_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_2133_v35_ = nw371_
                index372_ = default__.safeIndex(926, (d_2133_v35_).length(0))
                (d_2133_v35_)[index372_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2134_i3_ in range(default__.abs(448))])) + (d_2105_v10_)
                d_2135_v36_: _dafny.Map
                d_2135_v36_ = _dafny.Map({(0) - ((self).f9): d_2097_v6_})
                index373_ = default__.safeIndex(926, (d_2133_v35_).length(0))
                (d_2133_v35_)[index373_] = (_dafny.SeqWithoutIsStrInference([d_2096_v5_ for d_2136_i4_ in range(default__.abs(110))])) + ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2137_i5_ in range(default__.abs(-26))])) + (((d_2105_v10_).set(default__.safeIndex((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], len(d_2105_v10_)), (self).f5)).set(default__.safeIndex(len(d_2135_v36_), len((d_2105_v10_).set(default__.safeIndex((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], len(d_2105_v10_)), (self).f5))), _dafny.CodePoint('r'))))
                rhs420_ = d_2095_v4_
                rhs421_ = (self).f6
                d_2095_v4_ = rhs420_
                d_2094_v3_ = rhs421_
                d_2138_v37_: _dafny.Seq
                d_2138_v37_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2096_v5_ for d_2139_i6_ in range(default__.abs(-664))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))), d_2105_v10_])
                d_2138_v37_ = d_2138_v37_
            elif True:
                d_2140_v38_: _dafny.Map
                d_2140_v38_ = _dafny.Map({self.f10: (self).f9})
                rhs422_ = d_2094_v3_
                rhs423_ = p2
                rhs424_ = d_2140_v38_
                rhs425_ = not ((self.f10) == (len(d_2105_v10_))) or ((d_2105_v10_) == ((d_2105_v10_).set(default__.safeIndex(574, len(d_2105_v10_)), d_2096_v5_)))
                d_2094_v3_ = rhs422_
                d_2094_v3_ = rhs423_
                d_2140_v38_ = rhs424_
                d_2094_v3_ = rhs425_
                index374_ = default__.safeIndex(896, (d_2093_v2_).length(0))
                (d_2093_v2_)[index374_] = (d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]
                d_2141_v39_: _dafny.Map
                d_2141_v39_ = _dafny.Map({d_2091_v0_: d_2095_v4_})
                d_2142_v40_: _dafny.Map
                d_2142_v40_ = d_2141_v39_
                d_2143_v41_: _dafny.Map
                d_2143_v41_ = _dafny.Map({(d_2142_v40_): d_2105_v10_})
                d_2143_v41_ = (d_2143_v41_).set(((d_2141_v39_).set(d_2091_v0_, d_2095_v4_)) | (_dafny.Map({d_2091_v0_: d_2095_v4_})), (d_2105_v10_) + (d_2105_v10_))
                d_2144_v42_: _dafny.Array
                nw372_ = _dafny.Array(_dafny.Set({}), 21)
                d_2144_v42_ = nw372_
                d_2145_v43_: T1
                nw373_ = C3()
                nw373_.ctor__(p0, p2)
                d_2145_v43_ = nw373_
                d_2146_v44_: _dafny.Map
                d_2146_v44_ = _dafny.Map({d_2100_v9_: d_2145_v43_})
                d_2147_v45_: _dafny.Seq
                d_2147_v45_ = _dafny.SeqWithoutIsStrInference([d_2146_v44_])
                d_2148_v46_: _dafny.Map
                d_2148_v46_ = _dafny.Map({self.f10: (_dafny.MultiSet([p0])).cardinality})
                d_2149_v47_: _dafny.Set
                d_2149_v47_ = _dafny.Set({d_2146_v44_, (d_2147_v45_)[default__.safeIndex(((d_2148_v46_)[len(d_2105_v10_)] if (len(d_2105_v10_)) in (d_2148_v46_) else len(_dafny.Map({p2: (d_2097_v6_)[default__.safeIndex((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], len(d_2097_v6_))]}))), len(d_2147_v45_))]})
                index375_ = default__.safeIndex(878, (d_2144_v42_).length(0))
                (d_2144_v42_)[index375_] = d_2149_v47_
                index376_ = default__.safeIndex(878, (d_2144_v42_).length(0))
                (d_2144_v42_)[index376_] = ((d_2149_v47_).intersection(d_2149_v47_)) | (d_2149_v47_)
                d_2150_v48_: C3
                nw374_ = C3()
                nw374_.ctor__((self).f6, False)
                d_2150_v48_ = nw374_
        elif True:
            d_2151_v49_: C6
            nw375_ = C6()
            nw375_.ctor__()
            d_2151_v49_ = nw375_
            d_2151_v49_ = d_2151_v49_
            (self).f10 = (self).f9
            if not (p0) or (p0):
                d_2094_v3_ = True
                d_2152_v50_: _dafny.Map
                d_2152_v50_ = _dafny.Map({p2: d_2093_v2_})
                d_2152_v50_ = (d_2152_v50_).set(d_2094_v3_, d_2093_v2_)
                d_2153_v51_: _dafny.Map
                d_2153_v51_ = _dafny.Map({(d_2105_v10_) + (_dafny.SeqWithoutIsStrInference([d_2096_v5_ for d_2154_i7_ in range(default__.abs(-377))])): (self).f6})
                d_2153_v51_ = (d_2153_v51_).set(d_2105_v10_, False)
                d_2094_v3_ = not(True)
                d_2155_v52_: _dafny.Map
                d_2155_v52_ = _dafny.Map({65: d_2094_v3_})
                d_2156_v53_: _dafny.Map
                d_2156_v53_ = _dafny.Map({not((d_2097_v6_)[default__.safeIndex((0) - ((_dafny.MultiSet([127, (d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], (d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]])).cardinality), len(d_2097_v6_))]): (d_2155_v52_).set((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))], (self).f6)})
                d_2156_v53_ = ((d_2156_v53_) | (d_2156_v53_)).set(not((d_2092_v1_) <= (d_2092_v1_)), (d_2155_v52_) | (d_2155_v52_))
            elif True:
                index377_ = default__.safeIndex(896, (d_2093_v2_).length(0))
                rhs426_ = ((self).f6) and (False)
                rhs427_ = (0) - (default__.safeDivisionInt(default__.fm3((self).f6, d_2094_v3_, globalState), (self).f9))
                lhs283_ = d_2093_v2_
                lhs284_ = default__.safeIndex(896, (d_2093_v2_).length(0))
                d_2094_v3_ = rhs426_
                lhs283_[lhs284_] = rhs427_
                d_2157_v54_: _dafny.Array
                def lambda97_(d_2158_i8_):
                    return D11_DC19(_dafny.CodePoint('j'))

                init57_ = lambda97_
                nw376_ = _dafny.Array(None, 29)
                for i0_57_ in range(nw376_.length(0)):
                    nw376_[i0_57_] = init57_(i0_57_)
                d_2157_v54_ = nw376_
                d_2159_v55_: D11
                d_2159_v55_ = D11_DC19(d_2096_v5_)
                index378_ = default__.safeIndex(674, (d_2157_v54_).length(0))
                (d_2157_v54_)[index378_] = d_2159_v55_
                d_2160_v56_: _dafny.Seq
                d_2160_v56_ = _dafny.SeqWithoutIsStrInference([d_2092_v1_, d_2092_v1_])
                d_2161_v57_: C0
                nw377_ = C0()
                nw377_.ctor__((self).f6, d_2160_v56_)
                d_2161_v57_ = nw377_
                d_2162_v58_: _dafny.Seq
                d_2162_v58_ = _dafny.SeqWithoutIsStrInference([d_2161_v57_])
                d_2163_v59_: _dafny.Set
                d_2163_v59_ = _dafny.Set({102})
                d_2164_v60_: _dafny.Seq
                d_2164_v60_ = _dafny.SeqWithoutIsStrInference([d_2163_v59_])
                index379_ = default__.safeIndex(674, (d_2157_v54_).length(0))
                rhs428_ = d_2094_v3_
                rhs429_ = (self).f9
                rhs430_ = (_dafny.MultiSet(((((d_2162_v58_ if d_2094_v3_ else d_2162_v58_)) + ((d_2162_v58_) + (d_2162_v58_))).set(default__.safeIndex(default__.safeDivisionInt((self).f9, len(d_2164_v60_)), len(((d_2162_v58_ if d_2094_v3_ else d_2162_v58_)) + ((d_2162_v58_) + (d_2162_v58_)))), d_2161_v57_)).set(default__.safeIndex(self.f10, len((((d_2162_v58_ if d_2094_v3_ else d_2162_v58_)) + ((d_2162_v58_) + (d_2162_v58_))).set(default__.safeIndex(default__.safeDivisionInt((self).f9, len(d_2164_v60_)), len(((d_2162_v58_ if d_2094_v3_ else d_2162_v58_)) + ((d_2162_v58_) + (d_2162_v58_)))), d_2161_v57_))), d_2161_v57_))).cardinality
                rhs431_ = (self).f6
                rhs432_ = d_2159_v55_
                lhs285_ = globalState
                lhs286_ = globalState
                lhs287_ = d_2157_v54_
                lhs288_ = default__.safeIndex(674, (d_2157_v54_).length(0))
                d_2094_v3_ = rhs428_
                lhs285_.f1 = rhs429_
                lhs286_.f1 = rhs430_
                d_2094_v3_ = rhs431_
                lhs287_[lhs288_] = rhs432_
                d_2093_v2_ = d_2093_v2_
                d_2165_v61_: _dafny.Array
                def lambda98_(d_2166_v10_):
                    def lambda99_(d_2167_i9_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbknyswhp"))) + (d_2166_v10_)

                    return lambda99_

                init58_ = lambda98_(d_2105_v10_)
                nw378_ = _dafny.Array(None, 19)
                for i0_58_ in range(nw378_.length(0)):
                    nw378_[i0_58_] = init58_(i0_58_)
                d_2165_v61_ = nw378_
                index380_ = default__.safeIndex(905, (d_2165_v61_).length(0))
                (d_2165_v61_)[index380_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "penmrr"))) + (d_2105_v10_)
                index381_ = default__.safeIndex(905, (d_2165_v61_).length(0))
                (d_2165_v61_)[index381_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onejtsx")))
                d_2168_v62_: D2
                d_2168_v62_ = D2_DC8(d_2094_v3_)
                rhs433_ = (d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]
                rhs434_ = (d_2100_v9_ if ((d_2093_v2_)[default__.safeIndex(896, (d_2093_v2_).length(0))]) > (self.f10) else (self).fm5(d_2168_v62_, globalState))
                lhs289_ = globalState
                lhs289_.f1 = rhs433_
                d_2100_v9_ = rhs434_
            d_2096_v5_ = ((d_2096_v5_ if p2 else (self).f5) if p0 else (self).f5)
            d_2169_v63_: _dafny.Array
            d_2169_v63_ = d_2093_v2_
            d_2170_v64_: D12
            d_2170_v64_ = D12_DC23(d_2169_v63_)
            pat_let_tv29_ = d_2093_v2_
            pat_let_tv30_ = d_2169_v63_
            d_2171_v65_: _dafny.Array
            nw379_ = _dafny.Array(None, 13)
            nw379_[int(0)] = d_2170_v64_
            nw379_[int(1)] = d_2170_v64_
            nw379_[int(2)] = d_2170_v64_
            nw379_[int(3)] = d_2170_v64_
            nw379_[int(4)] = d_2170_v64_
            nw379_[int(5)] = d_2170_v64_
            nw379_[int(6)] = d_2170_v64_
            def iife143_(_pat_let39_0):
                def iife144_(d_2172_dt__update__tmp_h3_):
                    def iife145_(_pat_let40_0):
                        def iife146_(d_2173_dt__update_hcf29_h0_):
                            return D12_DC23(d_2173_dt__update_hcf29_h0_)
                        return iife146_(_pat_let40_0)
                    return iife145_(pat_let_tv29_)
                return iife144_(_pat_let39_0)
            def iife142_(_pat_let38_0):
                def iife147_(d_2174_dt__update__tmp_h4_):
                    def iife148_(_pat_let41_0):
                        def iife149_(d_2175_dt__update_hcf29_h1_):
                            return D12_DC23(d_2175_dt__update_hcf29_h1_)
                        return iife149_(_pat_let41_0)
                    return iife148_(pat_let_tv30_)
                return iife147_(_pat_let38_0)
            nw379_[int(7)] = iife142_(iife143_(d_2170_v64_))
            nw379_[int(8)] = d_2170_v64_
            nw379_[int(9)] = d_2170_v64_
            nw379_[int(10)] = D12_DC23(d_2169_v63_)
            nw379_[int(11)] = D12_DC23(d_2169_v63_)
            nw379_[int(12)] = d_2170_v64_
            d_2171_v65_ = nw379_
            d_2176_v66_: _dafny.Array
            nw380_ = _dafny.Array(None, 2)
            nw380_[int(0)] = d_2171_v65_
            nw380_[int(1)] = d_2171_v65_
            d_2176_v66_ = nw380_
            d_2176_v66_ = d_2176_v66_
        d_2177_v67_: _dafny.Set
        d_2177_v67_ = _dafny.Set({(self).f9})
        d_2178_v68_: _dafny.Map
        d_2178_v68_ = _dafny.Map({((D16_DC32(p2)).cf41) not in (_dafny.MultiSet([(self).fm6(d_2092_v1_, globalState)])): d_2177_v67_})
        d_2178_v68_ = (d_2178_v68_).set(False, d_2177_v67_)
        index382_ = default__.safeIndex(806, (d_2095_v4_).length(0))
        (d_2095_v4_)[index382_] = p2
        index383_ = default__.safeIndex(806, (d_2095_v4_).length(0))
        (d_2095_v4_)[index383_] = (self.f10) > ((d_2092_v1_)[default__.safeIndex(self.f10, len(d_2092_v1_))])

    def m7(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: _dafny.Seq = _dafny.Seq({})
        d_2179_v0_: D0
        d_2179_v0_ = D0_DC3(D0_DC0((self).f6))
        d_2179_v0_ = (d_2179_v0_ if (self).f6 else d_2179_v0_)
        if (self).f6:
            d_2180_v1_: _dafny.Map
            d_2180_v1_ = _dafny.Map({(self).f6: False})
            d_2180_v1_ = (d_2180_v1_) | (d_2180_v1_)
            d_2181_v2_: C3
            nw381_ = C3()
            nw381_.ctor__((self).f6, (self).f6)
            d_2181_v2_ = nw381_
            d_2182_v3_: _dafny.Seq
            d_2182_v3_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_2183_v4_: _dafny.MultiSet
            d_2183_v4_ = _dafny.MultiSet([(d_2182_v3_)[default__.safeIndex(len(d_2182_v3_), len(d_2182_v3_))], -856])
            d_2183_v4_ = d_2183_v4_
            d_2184_v5_: str
            d_2184_v5_ = _dafny.CodePoint('n')
            d_2184_v5_ = d_2184_v5_
            (globalState).f1 = (self.f10 if not((d_2181_v2_).f13) else -730)
        elif True:
            d_2185_v6_: _dafny.Seq
            d_2185_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
            if ((self).f5) not in ((D2_DC6(d_2185_v6_)).cf6):
                d_2186_v7_: bool
                d_2186_v7_ = False
                d_2186_v7_ = not(True)
                r0 = p0
                d_2187_v8_: _dafny.Array
                nw382_ = _dafny.Array(_dafny.Map({}), 6)
                d_2187_v8_ = nw382_
                d_2188_v9_: _dafny.Array
                nw383_ = _dafny.Array(False, 10)
                d_2188_v9_ = nw383_
                d_2189_v10_: _dafny.Seq
                d_2189_v10_ = _dafny.SeqWithoutIsStrInference([d_2188_v9_, d_2188_v9_, d_2188_v9_])
                d_2190_v11_: _dafny.Map
                d_2190_v11_ = _dafny.Map({d_2187_v8_: (d_2189_v10_)[default__.safeIndex((self).fm8(d_2186_v7_, d_2186_v7_, (self).f6, globalState), len(d_2189_v10_))]})
                d_2191_v12_: D17
                d_2191_v12_ = D17_DC33(d_2190_v11_)
                d_2190_v11_ = (d_2191_v12_).cf42
                d_2192_v13_: C6
                nw384_ = C6()
                nw384_.ctor__()
                d_2192_v13_ = nw384_
                d_2193_v14_: _dafny.Array
                nw385_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2193_v14_ = nw385_
                index384_ = default__.safeIndex(536, (d_2193_v14_).length(0))
                (d_2193_v14_)[index384_] = d_2188_v9_
                d_2194_v15_: _dafny.Map
                d_2194_v15_ = _dafny.Map({p0: d_2185_v6_})
                d_2195_v16_: _dafny.Seq
                d_2195_v16_ = _dafny.SeqWithoutIsStrInference([d_2185_v6_, d_2185_v6_])
                d_2196_v17_: _dafny.Map
                d_2196_v17_ = _dafny.Map({True: (d_2195_v16_)[default__.safeIndex((self).f9, len(d_2195_v16_))]})
                d_2197_v18_: _dafny.Map
                d_2197_v18_ = _dafny.Map({p0: p0})
                d_2198_v19_: _dafny.Map
                d_2198_v19_ = d_2197_v18_
                d_2199_v20_: _dafny.Seq
                d_2199_v20_ = _dafny.SeqWithoutIsStrInference([d_2186_v7_])
                d_2200_v21_: _dafny.Seq
                d_2200_v21_ = _dafny.SeqWithoutIsStrInference([d_2199_v20_, default__.fm2((self).f6, False, globalState), d_2199_v20_, d_2199_v20_])
                d_2201_v22_: D12
                d_2201_v22_ = D12_DC22(d_2198_v19_, d_2200_v21_, (self).f5, d_2186_v7_, default__.fm1(len(d_2185_v6_), d_2186_v7_, p0, globalState))
                d_2202_v23_: _dafny.Array
                nw386_ = _dafny.Array(None, 6)
                nw386_[int(0)] = (((d_2194_v15_)[p0] if (p0) in (d_2194_v15_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))
                nw386_[int(1)] = ((d_2196_v17_)[d_2186_v7_] if (d_2186_v7_) in (d_2196_v17_) else _dafny.SeqWithoutIsStrInference([(self).f5 for d_2203_i0_ in range(default__.abs(3))]))
                nw386_[int(2)] = (d_2185_v6_) + (d_2185_v6_)
                nw386_[int(3)] = d_2185_v6_
                nw386_[int(4)] = d_2185_v6_
                nw386_[int(5)] = (d_2201_v22_).cf28
                d_2202_v23_ = nw386_
                index385_ = default__.safeIndex(440, (d_2202_v23_).length(0))
                (d_2202_v23_)[index385_] = d_2185_v6_
                index386_ = default__.safeIndex(536, (d_2193_v14_).length(0))
                index387_ = default__.safeIndex(440, (d_2202_v23_).length(0))
                rhs435_ = d_2188_v9_
                rhs436_ = d_2185_v6_
                lhs290_ = d_2193_v14_
                lhs291_ = default__.safeIndex(536, (d_2193_v14_).length(0))
                lhs292_ = d_2202_v23_
                lhs293_ = default__.safeIndex(440, (d_2202_v23_).length(0))
                lhs290_[lhs291_] = rhs435_
                lhs292_[lhs293_] = rhs436_
            elif True:
                d_2204_v24_: C2
                nw387_ = C2()
                nw387_.ctor__((d_2185_v6_)[default__.safeIndex((self).f9, len(d_2185_v6_))], default__.fm0((self).f6, globalState))
                d_2204_v24_ = nw387_
                d_2205_v25_: bool
                d_2205_v25_ = True
                d_2205_v25_ = d_2205_v25_
                d_2206_v26_: _dafny.MultiSet
                d_2206_v26_ = _dafny.MultiSet([(self).f6])
                d_2207_v27_: _dafny.Seq
                d_2207_v27_ = _dafny.SeqWithoutIsStrInference([(self).f6, d_2205_v25_, d_2205_v25_])
                d_2208_v28_: _dafny.Array
                nw388_ = _dafny.Array(None, 20)
                nw388_[int(0)] = d_2206_v26_
                nw388_[int(1)] = d_2206_v26_
                nw388_[int(2)] = d_2206_v26_
                nw388_[int(3)] = d_2206_v26_
                nw388_[int(4)] = d_2206_v26_
                nw388_[int(5)] = d_2206_v26_
                nw388_[int(6)] = d_2206_v26_
                nw388_[int(7)] = d_2206_v26_
                nw388_[int(8)] = _dafny.MultiSet(d_2207_v27_)
                nw388_[int(9)] = d_2206_v26_
                nw388_[int(10)] = d_2206_v26_
                nw388_[int(11)] = d_2206_v26_
                nw388_[int(12)] = d_2206_v26_
                nw388_[int(13)] = d_2206_v26_
                nw388_[int(14)] = d_2206_v26_
                nw388_[int(15)] = d_2206_v26_
                nw388_[int(16)] = d_2206_v26_
                nw388_[int(17)] = d_2206_v26_
                nw388_[int(18)] = d_2206_v26_
                nw388_[int(19)] = d_2206_v26_
                d_2208_v28_ = nw388_
                d_2209_v29_: D18
                d_2209_v29_ = D18_DC36(d_2208_v28_)
                d_2210_v30_: _dafny.Array
                nw389_ = _dafny.Array(None, 20)
                nw389_[int(0)] = (d_2209_v29_).cf45
                nw389_[int(1)] = d_2208_v28_
                nw389_[int(2)] = d_2208_v28_
                nw389_[int(3)] = d_2208_v28_
                nw389_[int(4)] = d_2208_v28_
                nw389_[int(5)] = d_2208_v28_
                nw389_[int(6)] = d_2208_v28_
                nw389_[int(7)] = d_2208_v28_
                nw389_[int(8)] = (d_2208_v28_ if True else d_2208_v28_)
                nw389_[int(9)] = d_2208_v28_
                nw389_[int(10)] = d_2208_v28_
                nw389_[int(11)] = d_2208_v28_
                nw389_[int(12)] = d_2208_v28_
                nw389_[int(13)] = d_2208_v28_
                nw389_[int(14)] = d_2208_v28_
                nw389_[int(15)] = d_2208_v28_
                nw389_[int(16)] = d_2208_v28_
                nw389_[int(17)] = d_2208_v28_
                nw389_[int(18)] = d_2208_v28_
                nw389_[int(19)] = d_2208_v28_
                d_2210_v30_ = nw389_
                index388_ = default__.safeIndex(482, (d_2210_v30_).length(0))
                (d_2210_v30_)[index388_] = d_2208_v28_
                index389_ = default__.safeIndex(482, (d_2210_v30_).length(0))
                (d_2210_v30_)[index389_] = d_2208_v28_
                d_2205_v25_ = True
                rhs437_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_2205_v25_])), (0) - (p0))
                rhs438_ = 287
                lhs294_ = self
                r1 = rhs437_
                lhs294_.f10 = rhs438_
            d_2211_v31_: _dafny.Array
            nw390_ = _dafny.Array(None, 7)
            nw390_[int(0)] = (self).f9
            nw390_[int(1)] = (0) - ((self).f9)
            nw390_[int(2)] = 792
            nw390_[int(3)] = len(d_2185_v6_)
            nw390_[int(4)] = self.f10
            nw390_[int(5)] = (self).f9
            nw390_[int(6)] = self.f10
            d_2211_v31_ = nw390_
            index390_ = default__.safeIndex(572, (d_2211_v31_).length(0))
            (d_2211_v31_)[index390_] = (self).fm7((self).f6, globalState)
            index391_ = default__.safeIndex(572, (d_2211_v31_).length(0))
            (d_2211_v31_)[index391_] = (837) * ((self).f9)
            d_2212_v32_: C5
            nw391_ = C5()
            nw391_.ctor__(109, (self).f5, (self).f6)
            d_2212_v32_ = nw391_
            d_2212_v32_ = d_2212_v32_
            d_2213_v33_: C6
            nw392_ = C6()
            nw392_.ctor__()
            d_2213_v33_ = nw392_
            d_2214_v34_: _dafny.Array
            nw393_ = _dafny.Array(_dafny.Seq({}), 17)
            d_2214_v34_ = nw393_
            d_2214_v34_ = d_2214_v34_
        d_2215_v35_: _dafny.Set
        d_2215_v35_ = _dafny.Set({default__.fm0((self).f6, globalState)})
        d_2215_v35_ = d_2215_v35_
        d_2216_v36_: _dafny.Map
        d_2216_v36_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f5 for d_2217_i1_ in range(default__.abs(111))]): (self).f5})
        d_2218_v37_: _dafny.Seq
        d_2218_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adh"))
        d_2219_v38_: _dafny.Seq
        d_2219_v38_ = _dafny.SeqWithoutIsStrInference([d_2218_v37_, d_2218_v37_])
        d_2216_v36_ = (d_2216_v36_).set((d_2219_v38_)[default__.safeIndex(self.f10, len(d_2219_v38_))], _dafny.CodePoint('k'))
        r1 = p0
        d_2220_v39_: _dafny.Array
        nw394_ = _dafny.Array(False, 16)
        d_2220_v39_ = nw394_
        d_2221_v40_: _dafny.Map
        d_2221_v40_ = _dafny.Map({(self).f6: (self).f6})
        d_2222_v41_: D1
        d_2222_v41_ = D1_DC4(len(d_2221_v40_))
        d_2223_v42_: _dafny.Map
        d_2223_v42_ = _dafny.Map({d_2222_v41_: default__.fm0(False, globalState)})
        d_2224_v43_: D0
        d_2225_v44_: bool
        out48_: D0
        out49_: bool
        out48_, out49_ = default__.m0(d_2220_v39_, (self).f5, ((d_2223_v42_)[d_2222_v41_] if (d_2222_v41_) in (d_2223_v42_) else (self).f6), _dafny.CodePoint('i'), globalState)
        d_2224_v43_ = out48_
        d_2225_v44_ = out49_
        r0 = self.f10
        r1 = default__.safeDivisionInt(default__.safeDivisionInt(len(d_2218_v37_), self.f10), default__.fm3((self).f6, d_2225_v44_, globalState))
        d_2226_v45_: D6
        d_2226_v45_ = D6_DC14(d_2225_v44_)
        pat_let_tv31_ = p0
        def lambda100_(source28_):
            if source28_.is_DC14:
                d_2227___mcc_h0_ = source28_.cf15
                d_2228_cf15_ = d_2227___mcc_h0_
                if True:
                    return _dafny.SeqWithoutIsStrInference([(self).f9, self.f10, (self).f9, pat_let_tv31_])
                elif True:
                    return _dafny.SeqWithoutIsStrInference([self.f10 for d_2229_i2_ in range(default__.abs(571))])
            elif True:
                d_2230___mcc_h1_ = source28_.cf14
                d_2231_cf14_ = d_2230___mcc_h1_
                return d_2231_cf14_

        r2 = len(lambda100_(d_2226_v45_))
        d_2232_v46_: _dafny.Seq
        d_2232_v46_ = _dafny.SeqWithoutIsStrInference([d_2220_v39_])
        r3 = d_2232_v46_
        return r0, r1, r2, r3

    @property
    def f9(self):
        return self._f9

class C8(T4, T3, T2, T1, T0):
    def  __init__(self):
        self._f5: str = _dafny.CodePoint('D')
        self._f6: bool = False
        self.f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
    def ctor__(self, f7, f8, f5, f6):
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f5 = f5
        (self)._f6 = f6

    def fm9(self, p0, p1, p2, p3, globalState):
        def iife150_():
            coll66_ = _dafny.Set()
            compr_66_: D0
            for compr_66_ in (_dafny.SeqWithoutIsStrInference([D0_DC0(False), D0_DC0((self).f6), D0_DC0((self).f6), D0_DC0(True), D0_DC0((self).f6)])).Elements:
                d_2233_v0_: D0 = compr_66_
                if (d_2233_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC0(False), D0_DC0((self).f6), D0_DC0((self).f6), D0_DC0(True), D0_DC0((self).f6)])):
                    coll66_ = coll66_.union(_dafny.Set([d_2233_v0_]))
            return _dafny.Set(coll66_)
        def iife151_():
            def iife153_():
                coll69_ = _dafny.Set()
                compr_69_: D0
                for compr_69_ in (_dafny.MultiSet([D0_DC0(not((self).f6))])).Elements:
                    d_2236_v1_: D0 = compr_69_
                    if (d_2236_v1_) in (_dafny.MultiSet([D0_DC0(not((self).f6))])):
                        coll69_ = coll69_.union(_dafny.Set([d_2236_v1_]))
                return _dafny.Set(coll69_)
            coll67_ = _dafny.Set()
            def iife152_():
                coll68_ = _dafny.Set()
                compr_68_: D0
                for compr_68_ in (_dafny.MultiSet([D0_DC0(not((self).f6))])).Elements:
                    d_2234_v1_: D0 = compr_68_
                    if (d_2234_v1_) in (_dafny.MultiSet([D0_DC0(not((self).f6))])):
                        coll68_ = coll68_.union(_dafny.Set([d_2234_v1_]))
                return _dafny.Set(coll68_)
            compr_67_: D0
            for compr_67_ in (iife152_()
            ).Elements:
                d_2235_v2_: D0 = compr_67_
                if (d_2235_v2_) in (iife153_()
                ):
                    coll67_ = coll67_.union(_dafny.Set([d_2235_v2_]))
            return _dafny.Set(coll67_)
        return ((iife150_()
        ) | (_dafny.Set({D0_DC0(True), D0_DC0((self).f6), D0_DC0((self).f6), D0_DC0(not((self).f6))}))).isdisjoint(iife151_()
        )

    def fm10(self, p0, p1, p2, globalState):
        return (default__.safeModuloInt(-966, len(_dafny.SeqWithoutIsStrInference([(self).f6])))) != (((D2_DC7(len(_dafny.SeqWithoutIsStrInference([-661])), -667)).cf8) - (len(_dafny.Map({not(not((self).f6)): (self).f6}))))

    def fm8(self, p0, p1, p2, globalState):
        return 434

    def fm6(self, p0, globalState):
        return not((self).f6)

    def fm7(self, p0, globalState):
        source29_ = D2_DC7(899, (0) - (len(_dafny.Set({_dafny.Set({(self).f6, (self).f6})}))))
        if source29_.is_DC7:
            d_2237___mcc_h0_ = source29_.cf7
            d_2238___mcc_h1_ = source29_.cf8
            d_2239_cf8_ = d_2238___mcc_h1_
            d_2240_cf7_ = d_2237___mcc_h0_
            return len((_dafny.Set({(self).f6})) | (_dafny.Set({not((self).f6)})))
        elif source29_.is_DC8:
            d_2241___mcc_h2_ = source29_.cf9
            d_2242_cf9_ = d_2241___mcc_h2_
            def iife154_():
                coll70_ = _dafny.Map()
                compr_70_: int
                for compr_70_ in (_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngprv"))): False}))])).Elements:
                    d_2243_v0_: int = compr_70_
                    if (d_2243_v0_) in (_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngprv"))): False}))])):
                        coll70_[default__.safeDivisionInt(d_2243_v0_, 266)] = d_2242_cf9_
                return _dafny.Map(coll70_)
            return len(iife154_()
            )
        elif source29_.is_DC6:
            d_2244___mcc_h3_ = source29_.cf6
            d_2245_cf6_ = d_2244___mcc_h3_
            return (-427) + (94)
        elif True:
            d_2246___mcc_h4_ = source29_.cf10
            d_2247_cf10_ = d_2246___mcc_h4_
            return default__.safeDivisionInt(409, 250)

    def fm4(self, p0, p1, p2, p3, globalState):
        return len(((_dafny.SeqWithoutIsStrInference([not((self).f6)])) + (_dafny.SeqWithoutIsStrInference([(self).f6, (self).f6]))) + (_dafny.SeqWithoutIsStrInference([False])))

    def fm5(self, p0, globalState):
        return D1_DC4(default__.safeDivisionInt(113, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-300]))}))))

    def fm11(self, p0, p1, p2, globalState):
        if (_dafny.Set({(self).f6})).issubset(_dafny.Set({(self).f6, (self).f6, (self).f6})):
            return _dafny.Map({609: len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, (self).f6])).cardinality for d_2248_i0_ in range(default__.abs(-597))]))})
        elif True:
            def iife155_():
                coll71_ = _dafny.Set()
                compr_71_: int
                for compr_71_ in _dafny.IntegerRange(19, 9):
                    d_2249_v0_: int = compr_71_
                    if ((19) <= (d_2249_v0_)) and ((d_2249_v0_) < (9)):
                        coll71_ = coll71_.union(_dafny.Set([(d_2249_v0_) * ((_dafny.MultiSet([585])).cardinality)]))
                return _dafny.Set(coll71_)
            return ((_dafny.Map({len(iife155_()
            ): len(_dafny.Set({(self).f6, (self).f6}))}))) | (_dafny.Map({168: 385}))

    def fm12(self, globalState):
        return (self).f6

    def m4(self, p0, p1, p2, p3, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        index392_ = default__.safeIndex(174, (p0).length(0))
        (p0)[index392_] = p3
        index393_ = default__.safeIndex(174, (p0).length(0))
        (p0)[index393_] = default__.fm3(not(p2), (p2 if p2 else (self).f6), globalState)
        d_2250_v0_: bool
        d_2250_v0_ = True
        d_2251_v1_: _dafny.Map
        d_2251_v1_ = _dafny.Map({self.f7: (self).f5})
        d_2252_v2_: D0
        d_2252_v2_ = D0_DC2(d_2251_v1_, (self).f5)
        d_2253_v3_: _dafny.Seq
        d_2253_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghkphwmku"))
        d_2254_v4_: _dafny.MultiSet
        d_2254_v4_ = _dafny.MultiSet([p3, p3])
        d_2255_v5_: _dafny.Seq
        d_2255_v5_ = _dafny.SeqWithoutIsStrInference([d_2254_v4_, d_2254_v4_])
        d_2256_v6_: _dafny.MultiSet
        d_2256_v6_ = _dafny.MultiSet([(self).f6, p2, (d_2254_v4_).issubset((d_2255_v5_)[default__.safeIndex((p0)[default__.safeIndex(174, (p0).length(0))], len(d_2255_v5_))]), (d_2250_v0_) and ((self).f6)])
        rhs439_ = d_2250_v0_
        rhs440_ = d_2252_v2_
        rhs441_ = (d_2256_v6_).cardinality
        rhs442_ = d_2253_v3_
        lhs295_ = globalState
        d_2250_v0_ = rhs439_
        d_2252_v2_ = rhs440_
        lhs295_.f1 = rhs441_
        d_2253_v3_ = rhs442_
        d_2257_v7_: C1
        nw395_ = C1()
        nw395_.ctor__()
        d_2257_v7_ = nw395_
        d_2258_v8_: _dafny.Seq
        d_2258_v8_ = _dafny.SeqWithoutIsStrInference([p3, (p0)[default__.safeIndex(174, (p0).length(0))], (p0)[default__.safeIndex(174, (p0).length(0))]])
        d_2259_v9_: _dafny.Map
        d_2259_v9_ = _dafny.Map({(p0)[default__.safeIndex(174, (p0).length(0))]: d_2258_v8_})
        d_2260_v10_: D6
        d_2260_v10_ = D6_DC13(((d_2259_v9_)[(p0)[default__.safeIndex(174, (p0).length(0))]] if ((p0)[default__.safeIndex(174, (p0).length(0))]) in (d_2259_v9_) else d_2258_v8_))
        d_2261_v11_: _dafny.Map
        d_2261_v11_ = _dafny.Map({d_2260_v10_: p3})
        hi12_ = default__.safeModuloInt(p3, (0) - ((p0)[default__.safeIndex(174, (p0).length(0))]))
        for d_2262_i0_ in range(((d_2261_v11_)[d_2260_v10_] if (d_2260_v10_) in (d_2261_v11_) else 415), hi12_):
            index394_ = default__.safeIndex(174, (p0).length(0))
            (p0)[index394_] = (d_2258_v8_)[default__.safeIndex(-641, len(d_2258_v8_))]
            d_2263_v12_: _dafny.Array
            def lambda101_(d_2264_i1_):
                return (self).f5

            init59_ = lambda101_
            nw396_ = _dafny.Array(None, 9)
            for i0_59_ in range(nw396_.length(0)):
                nw396_[i0_59_] = init59_(i0_59_)
            d_2263_v12_ = nw396_
            nw397_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_2263_v12_ = nw397_
            d_2258_v8_ = d_2258_v8_
            d_2265_v13_: C3
            nw398_ = C3()
            nw398_.ctor__((self).f6, p2)
            d_2265_v13_ = nw398_
        if d_2250_v0_:
            if (p2) and (d_2250_v0_):
                d_2266_v14_: C7
                nw399_ = C7()
                nw399_.ctor__((0) - (p3), (p0)[default__.safeIndex(174, (p0).length(0))], (self).f5, d_2250_v0_)
                d_2266_v14_ = nw399_
                d_2267_v15_: str
                d_2267_v15_ = _dafny.CodePoint('q')
                d_2267_v15_ = d_2267_v15_
                d_2268_v16_: _dafny.Map
                d_2268_v16_ = _dafny.Map({p3: (d_2266_v14_).f9})
                d_2269_v17_: _dafny.MultiSet
                d_2269_v17_ = _dafny.MultiSet([d_2253_v3_])
                d_2250_v0_ = (d_2269_v17_).issubset(default__.fm61(d_2268_v16_, globalState))
                index395_ = default__.safeIndex(174, (p0).length(0))
                (p0)[index395_] = 567
                d_2270_v18_: _dafny.Map
                d_2270_v18_ = _dafny.Map({(self).fm12(globalState): d_2250_v0_})
                d_2270_v18_ = (d_2270_v18_).set(d_2250_v0_, d_2250_v0_)
            elif True:
                d_2271_v19_: _dafny.Array
                nw400_ = _dafny.Array(None, 20)
                d_2271_v19_ = nw400_
                d_2272_v20_: C2
                nw401_ = C2()
                nw401_.ctor__(default__.fm22(_dafny.Set({(self).f6}), (_dafny.MultiSet([True])).cardinality, p2, (p1)[default__.safeIndex(p3, len(p1))], globalState), not(d_2250_v0_))
                d_2272_v20_ = nw401_
                index396_ = default__.safeIndex(435, (d_2271_v19_).length(0))
                (d_2271_v19_)[index396_] = d_2272_v20_
                index397_ = default__.safeIndex(435, (d_2271_v19_).length(0))
                (d_2271_v19_)[index397_] = d_2272_v20_
                d_2273_v21_: _dafny.Array
                nw402_ = _dafny.Array(_dafny.Map({}), 21)
                d_2273_v21_ = nw402_
                d_2274_v22_: _dafny.Map
                d_2274_v22_ = _dafny.Map({(p0)[default__.safeIndex(174, (p0).length(0))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lyknjnvny"))})
                index398_ = default__.safeIndex(408, (d_2273_v21_).length(0))
                (d_2273_v21_)[index398_] = (d_2274_v22_).set(p3, d_2253_v3_)
                index399_ = default__.safeIndex(408, (d_2273_v21_).length(0))
                (d_2273_v21_)[index399_] = (d_2274_v22_) | (_dafny.Map({p3: d_2253_v3_}))
                d_2275_v23_: D11
                d_2275_v23_ = D11_DC20(d_2250_v0_, p3)
                rhs443_ = p2
                rhs444_ = (d_2275_v23_).cf21
                rhs445_ = p3
                lhs296_ = globalState
                d_2250_v0_ = rhs443_
                d_2250_v0_ = rhs444_
                lhs296_.f1 = rhs445_
                d_2250_v0_ = ((self).f6) == (d_2250_v0_)
                d_2276_v24_: _dafny.Map
                d_2276_v24_ = _dafny.Map({d_2250_v0_: (len(d_2253_v3_)) + (-220)})
                d_2276_v24_ = (d_2276_v24_).set((self).f6, (p0)[default__.safeIndex(174, (p0).length(0))])
            d_2277_v25_: _dafny.Map
            d_2277_v25_ = _dafny.Map({default__.safeDivisionInt((p0)[default__.safeIndex(174, (p0).length(0))], (0) - (p3)): p3})
            d_2277_v25_ = (d_2277_v25_).set((p0)[default__.safeIndex(174, (p0).length(0))], (p0)[default__.safeIndex(174, (p0).length(0))])
            if (d_2254_v4_).ispropersubset((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smcp"))), (p0)[default__.safeIndex(174, (p0).length(0))], (d_2256_v6_).cardinality])).set(len(_dafny.Map({p3: (self).f6})), default__.abs(p3))):
                d_2253_v3_ = d_2253_v3_
                d_2250_v0_ = False
                rhs446_ = (p3) * (p3)
                rhs447_ = (p3) + (((p0)[default__.safeIndex(174, (p0).length(0))]) * ((0) - (p3)))
                lhs297_ = globalState
                lhs298_ = globalState
                lhs297_.f1 = rhs446_
                lhs298_.f1 = rhs447_
                d_2278_v26_: _dafny.Map
                d_2278_v26_ = _dafny.Map({p2: (self).f6})
                d_2278_v26_ = d_2278_v26_
                (globalState).f1 = p3
            elif True:
                d_2279_v27_: _dafny.Array
                def lambda102_(d_2280_p3_):
                    def lambda103_(d_2281_i2_):
                        return (d_2281_i2_) - (d_2280_p3_)

                    return lambda103_

                init60_ = lambda102_(p3)
                nw403_ = _dafny.Array(None, 19)
                for i0_60_ in range(nw403_.length(0)):
                    nw403_[i0_60_] = init60_(i0_60_)
                d_2279_v27_ = nw403_
                rhs448_ = (p3) * (len((D6_DC13(d_2258_v8_)).cf14))
                rhs449_ = self.f7
                rhs450_ = (self).f6
                rhs451_ = p0
                lhs299_ = globalState
                lhs300_ = self
                lhs299_.f1 = rhs448_
                lhs300_.f7 = rhs449_
                d_2250_v0_ = rhs450_
                d_2279_v27_ = rhs451_
                d_2250_v0_ = (self).f6
                d_2282_v28_: _dafny.Set
                d_2282_v28_ = _dafny.Set({d_2250_v0_})
                d_2283_v29_: _dafny.Map
                d_2283_v29_ = _dafny.Map({len((d_2282_v28_) | (d_2282_v28_)): d_2250_v0_})
                d_2283_v29_ = (d_2283_v29_).set((-999) - (p3), p2)
                d_2284_v30_: _dafny.Set
                d_2284_v30_ = _dafny.Set({len(default__.fm1((p0)[default__.safeIndex(174, (p0).length(0))], not(p2), (p0)[default__.safeIndex(174, (p0).length(0))], globalState))})
                d_2284_v30_ = d_2284_v30_
                d_2250_v0_ = p2
            d_2285_v31_: str
            d_2285_v31_ = _dafny.CodePoint('g')
            d_2285_v31_ = (self).f5
            d_2250_v0_ = d_2250_v0_
        elif True:
            d_2286_v32_: C6
            nw404_ = C6()
            nw404_.ctor__()
            d_2286_v32_ = nw404_
            d_2287_v33_: D14
            d_2287_v33_ = D14_DC26((self).f6, p3, default__.fm53(d_2253_v3_, (p0)[default__.safeIndex(174, (p0).length(0))], globalState))
            (globalState).f1 = (d_2258_v8_)[default__.safeIndex((d_2287_v33_).cf33, len(d_2258_v8_))]
            d_2250_v0_ = (self).f6
            d_2288_v34_: _dafny.Set
            d_2288_v34_ = _dafny.Set({p2})
            d_2289_v35_: _dafny.Map
            d_2289_v35_ = _dafny.Map({self.f7: (d_2288_v34_).isdisjoint(d_2288_v34_)})
            d_2290_v36_: _dafny.Seq
            d_2290_v36_ = _dafny.SeqWithoutIsStrInference([d_2289_v35_, _dafny.Map({self.f7: p2})])
            rhs452_ = ((d_2289_v35_).set(self.f7, (self).f6)) | ((d_2290_v36_)[default__.safeIndex((p0)[default__.safeIndex(174, (p0).length(0))], len(d_2290_v36_))])
            rhs453_ = False
            rhs454_ = d_2250_v0_
            rhs455_ = len((_dafny.Map({(self).f6: p2})).set(default__.fm0((self).f6, globalState), (p1)[default__.safeIndex(p3, len(p1))]))
            lhs301_ = globalState
            d_2289_v35_ = rhs452_
            d_2250_v0_ = rhs453_
            d_2250_v0_ = rhs454_
            lhs301_.f1 = rhs455_
            index400_ = default__.safeIndex(174, (p0).length(0))
            (p0)[index400_] = (p0)[default__.safeIndex(174, (p0).length(0))]
        d_2291_v37_: D11
        d_2291_v37_ = D11_DC20(not(True), (p0)[default__.safeIndex(174, (p0).length(0))])
        d_2292_v38_: _dafny.Map
        d_2292_v38_ = _dafny.Map({d_2258_v8_: (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymtkf")))) >= ((d_2291_v37_).cf22)})
        (globalState).f1 = len(d_2292_v38_)
        d_2293_v39_: _dafny.Map
        d_2293_v39_ = _dafny.Map({p2: p3})
        d_2294_v40_: _dafny.Map
        d_2294_v40_ = _dafny.Map({((d_2293_v39_)[not(p2)] if (not(p2)) in (d_2293_v39_) else (p0)[default__.safeIndex(174, (p0).length(0))]): d_2250_v0_})
        r0 = default__.fm62(D15_DC27(_dafny.MultiSet([False, p2])), ((d_2294_v40_)[p3] if (p3) in (d_2294_v40_) else d_2250_v0_), d_2250_v0_, (d_2258_v8_) + (d_2258_v8_), globalState)
        r1 = d_2256_v6_
        return r0, r1

    def m5(self, p0, p1, globalState):
        d_2295_v0_: _dafny.Seq
        d_2295_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duo"))
        d_2295_v0_ = p0
        d_2296_v1_: bool
        d_2296_v1_ = False
        d_2296_v1_ = d_2296_v1_
        d_2297_v2_: _dafny.Map
        d_2297_v2_ = _dafny.Map({p1: d_2296_v1_})
        source30_ = (d_2297_v2_) | (d_2297_v2_)
        d_2298___mcc_h0_ = source30_
        d_2299_cf12_ = d_2298___mcc_h0_
        d_2296_v1_ = d_2296_v1_
        (globalState).f1 = p1
        d_2300_v3_: _dafny.Map
        d_2300_v3_ = _dafny.Map({default__.safeDivisionInt(p1, len(_dafny.Set({d_2296_v1_}))): p1})
        d_2300_v3_ = (d_2300_v3_).set(133, (0) - (p1))
        d_2301_v4_: _dafny.Array
        d_2301_v4_ = self.f7
        d_2302_v5_: _dafny.Map
        d_2302_v5_ = _dafny.Map({d_2301_v4_: self.f7})
        d_2303_v6_: _dafny.Map
        d_2303_v6_ = d_2302_v5_
        source31_ = d_2303_v6_
        d_2304___mcc_h1_ = source31_
        d_2305_cf30_ = d_2304___mcc_h1_
        rhs456_ = p0
        rhs457_ = (d_2300_v3_) | ((d_2300_v3_) | ((_dafny.Map({p1: -500})).set(p1, p1)))
        d_2295_v0_ = rhs456_
        d_2300_v3_ = rhs457_
        d_2296_v1_ = (self).f6
        d_2306_v7_: _dafny.Array
        def lambda104_(d_2307_p1_):
            def lambda105_(d_2308_i0_):
                return (d_2308_i0_) - (d_2307_p1_)

            return lambda105_

        init61_ = lambda104_(p1)
        nw405_ = _dafny.Array(None, 16)
        for i0_61_ in range(nw405_.length(0)):
            nw405_[i0_61_] = init61_(i0_61_)
        d_2306_v7_ = nw405_
        d_2309_v8_: _dafny.Seq
        d_2309_v8_ = _dafny.SeqWithoutIsStrInference([(self).f6, not(d_2296_v1_), (self).f6, (self).f6, True])
        index401_ = default__.safeIndex(221, (d_2306_v7_).length(0))
        (d_2306_v7_)[index401_] = (len(d_2309_v8_)) * (p1)
        arr2_ = self.f7
        index402_ = default__.safeIndex(685, (self.f7).length(0))
        arr2_[index402_] = False
        d_2310_v10_: _dafny.Seq
        d_2310_v10_ = _dafny.SeqWithoutIsStrInference([p1])
        index403_ = default__.safeIndex(221, (d_2306_v7_).length(0))
        arr3_ = self.f7
        index404_ = default__.safeIndex(685, (self.f7).length(0))
        def iife156_():
            coll72_ = _dafny.Map()
            compr_72_: int
            for compr_72_ in (d_2310_v10_).Elements:
                d_2311_v9_: int = compr_72_
                if (d_2311_v9_) in (d_2310_v10_):
                    coll72_[default__.safeDivisionInt(d_2311_v9_, p1)] = p1
            return _dafny.Map(coll72_)
        rhs458_ = p1
        rhs459_ = default__.safeModuloInt(p1, p1)
        rhs460_ = (default__.safeDivisionInt(p1, (0) - (len(iife156_()
        )))) + (p1)
        rhs461_ = (self).f6
        lhs302_ = globalState
        lhs303_ = d_2306_v7_
        lhs304_ = default__.safeIndex(221, (d_2306_v7_).length(0))
        lhs305_ = globalState
        lhs306_ = self.f7
        lhs307_ = default__.safeIndex(685, (self.f7).length(0))
        lhs302_.f1 = rhs458_
        lhs303_[lhs304_] = rhs459_
        lhs305_.f1 = rhs460_
        lhs306_[lhs307_] = rhs461_
        d_2312_v11_: _dafny.Map
        d_2312_v11_ = d_2300_v3_
        d_2313_v12_: _dafny.Seq
        d_2313_v12_ = _dafny.SeqWithoutIsStrInference([d_2309_v8_, d_2309_v8_])
        d_2314_v13_: D12
        d_2314_v13_ = D12_DC22(d_2312_v11_, d_2313_v12_, (self).f5, d_2296_v1_, d_2295_v0_)
        d_2315_v14_: _dafny.Seq
        d_2315_v14_ = _dafny.SeqWithoutIsStrInference([(d_2314_v13_).cf28, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aldrxfvuv"))) + (p0), d_2295_v0_, _dafny.SeqWithoutIsStrInference([(self).f5 for d_2316_i1_ in range(default__.abs(461))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toj"))])
        d_2315_v14_ = d_2315_v14_
        d_2296_v1_ = not(d_2296_v1_)
        (globalState).f1 = len((D19_DC38(default__.fm63(p1, globalState))).cf49)
        d_2317_v15_: _dafny.Array
        def lambda106_(d_2318_i3_):
            return (d_2318_i3_) - (len(_dafny.SeqWithoutIsStrInference([(self).f6])))

        init62_ = lambda106_
        nw406_ = _dafny.Array(None, 3)
        for i0_62_ in range(nw406_.length(0)):
            nw406_[i0_62_] = init62_(i0_62_)
        d_2317_v15_ = nw406_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2317_v15_).length(0)):
            d_2319_i2_: int = guard_loop_5_
            if (True) and (((0) <= (d_2319_i2_)) and ((d_2319_i2_) < ((d_2317_v15_).length(0)))):
                (d_2317_v15_)[(d_2319_i2_)] = (d_2319_i2_) + ((D14_DC26((self).f6, len(_dafny.Map({not((self).f6): False})), (self).f5)).cf33)

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2320_v0_: _dafny.MultiSet
        d_2320_v0_ = _dafny.MultiSet([(self).f6, (self).f6, (self).f6])
        hi13_ = p1
        for d_2321_i0_ in range((d_2320_v0_).cardinality, hi13_):
            d_2322_v1_: _dafny.Array
            def lambda107_(d_2323_p1_):
                def lambda108_(d_2324_i1_):
                    return (d_2324_i1_) - (d_2323_p1_)

                return lambda108_

            init63_ = lambda107_(p1)
            nw407_ = _dafny.Array(None, 5)
            for i0_63_ in range(nw407_.length(0)):
                nw407_[i0_63_] = init63_(i0_63_)
            d_2322_v1_ = nw407_
            index405_ = default__.safeIndex(842, (d_2322_v1_).length(0))
            (d_2322_v1_)[index405_] = d_2321_i0_
            d_2325_v2_: _dafny.Map
            d_2325_v2_ = _dafny.Map({(self).f6: 815})
            d_2326_v3_: D0
            d_2326_v3_ = D0_DC0((self).f6)
            d_2327_v5_: _dafny.Map
            d_2327_v5_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_2321_i0_])), p1])})
            d_2328_v6_: _dafny.Seq
            d_2328_v6_ = _dafny.SeqWithoutIsStrInference([p1, len(p2)])
            index406_ = default__.safeIndex(842, (d_2322_v1_).length(0))
            def iife157_():
                coll73_ = _dafny.Map()
                compr_73_: int
                for compr_73_ in (((d_2327_v5_)[(self).f6] if ((self).f6) in (d_2327_v5_) else d_2328_v6_)).Elements:
                    d_2329_v4_: int = compr_73_
                    if (d_2329_v4_) in (((d_2327_v5_)[(self).f6] if ((self).f6) in (d_2327_v5_) else d_2328_v6_)):
                        coll73_[(d_2329_v4_) * (d_2321_i0_)] = len(d_2325_v2_)
                return _dafny.Map(coll73_)
            rhs462_ = (p1) + (((d_2325_v2_)[p0] if (p0) in (d_2325_v2_) else (0) - (p1)))
            rhs463_ = default__.fm0((self).fm9(p1, d_2326_v3_, (self).f6, p1, globalState), globalState)
            rhs464_ = (((d_2320_v0_)[p0] if (p0) in (d_2320_v0_) else default__.fm3((self).f6, (self).f6, globalState))) >= (d_2321_i0_)
            rhs465_ = default__.safeDivisionInt((len(iife157_()
            )) + (p1), (d_2321_i0_) * (p1))
            lhs308_ = globalState
            lhs309_ = d_2322_v1_
            lhs310_ = default__.safeIndex(842, (d_2322_v1_).length(0))
            lhs308_.f1 = rhs462_
            r0 = rhs463_
            r0 = rhs464_
            lhs309_[lhs310_] = rhs465_
            d_2330_v7_: _dafny.Seq
            d_2330_v7_ = _dafny.SeqWithoutIsStrInference([p0])
            rhs466_ = (self).f6
            rhs467_ = d_2320_v0_
            rhs468_ = (not((self).f6)) == ((d_2330_v7_)[default__.safeIndex((d_2322_v1_)[default__.safeIndex(842, (d_2322_v1_).length(0))], len(d_2330_v7_))])
            rhs469_ = self.f7
            lhs311_ = globalState
            lhs312_ = self
            r0 = rhs466_
            lhs311_.f4 = rhs467_
            r0 = rhs468_
            lhs312_.f7 = rhs469_
            r0 = (self).f6
            index407_ = default__.safeIndex(842, (d_2322_v1_).length(0))
            (d_2322_v1_)[index407_] = (d_2322_v1_)[default__.safeIndex(842, (d_2322_v1_).length(0))]
        d_2331_i2_: int
        d_2331_i2_ = 0
        with _dafny.label("22"):
            while (self).f6:
                with _dafny.c_label("22"):
                    if (d_2331_i2_) >= (100):
                        raise _dafny.Break("22")
                    d_2331_i2_ = (d_2331_i2_) + (1)
                    (globalState).f1 = p1
                    d_2332_v8_: _dafny.Map
                    d_2332_v8_ = _dafny.Map({p0: (self).f5})
                    d_2333_v9_: D11
                    d_2333_v9_ = D11_DC20((self).f6, len(d_2332_v8_))
                    d_2334_v10_: _dafny.Array
                    nw408_ = _dafny.Array(None, 12)
                    nw408_[int(0)] = (d_2333_v9_).cf22
                    nw408_[int(1)] = (0) - (p1)
                    nw408_[int(2)] = p1
                    nw408_[int(3)] = (d_2320_v0_).cardinality
                    nw408_[int(4)] = 544
                    nw408_[int(5)] = len(p2)
                    nw408_[int(6)] = p1
                    nw408_[int(7)] = p1
                    nw408_[int(8)] = 669
                    nw408_[int(9)] = (0) - (p1)
                    nw408_[int(10)] = p1
                    nw408_[int(11)] = p1
                    d_2334_v10_ = nw408_
                    d_2335_v11_: C4
                    nw409_ = C4()
                    nw409_.ctor__()
                    d_2335_v11_ = nw409_
                    d_2336_v12_: _dafny.MultiSet
                    d_2336_v12_ = _dafny.MultiSet([d_2335_v11_, d_2335_v11_, d_2335_v11_, d_2335_v11_])
                    index408_ = default__.safeIndex(913, (d_2334_v10_).length(0))
                    (d_2334_v10_)[index408_] = ((d_2336_v12_)[d_2335_v11_] if (d_2335_v11_) in (d_2336_v12_) else len(p2))
                    d_2337_v13_: _dafny.Map
                    d_2337_v13_ = _dafny.Map({(self).f6: p2})
                    index409_ = default__.safeIndex(913, (d_2334_v10_).length(0))
                    (d_2334_v10_)[index409_] = (len(d_2337_v13_)) - (p1)
                    d_2338_v14_: _dafny.Array
                    nw410_ = _dafny.Array(_dafny.Array(None, 0), 5)
                    d_2338_v14_ = nw410_
                    d_2339_v15_: _dafny.Array
                    nw411_ = _dafny.Array(D2.default()(), 29)
                    d_2339_v15_ = nw411_
                    d_2340_v16_: _dafny.Array
                    nw412_ = _dafny.Array(None, 17)
                    nw412_[int(0)] = d_2339_v15_
                    nw412_[int(1)] = d_2339_v15_
                    nw412_[int(2)] = d_2339_v15_
                    nw412_[int(3)] = d_2339_v15_
                    nw412_[int(4)] = d_2339_v15_
                    nw412_[int(5)] = d_2339_v15_
                    nw412_[int(6)] = d_2339_v15_
                    nw412_[int(7)] = d_2339_v15_
                    nw412_[int(8)] = d_2339_v15_
                    nw412_[int(9)] = d_2339_v15_
                    nw412_[int(10)] = d_2339_v15_
                    nw412_[int(11)] = d_2339_v15_
                    nw412_[int(12)] = d_2339_v15_
                    nw412_[int(13)] = d_2339_v15_
                    nw412_[int(14)] = d_2339_v15_
                    nw412_[int(15)] = d_2339_v15_
                    nw412_[int(16)] = d_2339_v15_
                    d_2340_v16_ = nw412_
                    index410_ = default__.safeIndex(622, (d_2338_v14_).length(0))
                    (d_2338_v14_)[index410_] = d_2340_v16_
                    index411_ = default__.safeIndex(622, (d_2338_v14_).length(0))
                    (d_2338_v14_)[index411_] = d_2340_v16_
                    d_2341_v17_: T1
                    nw413_ = C6()
                    nw413_.ctor__()
                    d_2341_v17_ = nw413_
                    d_2341_v17_ = d_2341_v17_
                    pass
            pass
        d_2342_v18_: _dafny.Seq
        d_2342_v18_ = _dafny.SeqWithoutIsStrInference([p2, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))])
        d_2343_i3_: int
        d_2343_i3_ = 0
        with _dafny.label("23"):
            while (d_2342_v18_) < (_dafny.SeqWithoutIsStrInference([p2])):
                with _dafny.c_label("23"):
                    if (d_2343_i3_) >= (100):
                        raise _dafny.Break("23")
                    d_2343_i3_ = (d_2343_i3_) + (1)
                    d_2344_v19_: C6
                    nw414_ = C6()
                    nw414_.ctor__()
                    d_2344_v19_ = nw414_
                    d_2345_v20_: _dafny.Map
                    d_2345_v20_ = _dafny.Map({d_2344_v19_: p0})
                    d_2346_v22_: D14
                    def iife158_():
                        coll74_ = _dafny.Map()
                        compr_74_: int
                        for compr_74_ in _dafny.IntegerRange(137, 471):
                            d_2347_v21_: int = compr_74_
                            if ((137) <= (d_2347_v21_)) and ((d_2347_v21_) < (471)):
                                coll74_[default__.safeModuloInt(d_2347_v21_, p1)] = p1
                        return _dafny.Map(coll74_)
                    d_2346_v22_ = D14_DC26((self).f6, p1, default__.fm36(p1, p1, len(d_2345_v20_), iife158_()
, globalState))
                    source32_ = d_2346_v22_
                    if source32_.is_DC26:
                        d_2348___mcc_h0_ = source32_.cf32
                        d_2349___mcc_h1_ = source32_.cf33
                        d_2350___mcc_h2_ = source32_.cf34
                        d_2351_cf34_ = d_2350___mcc_h2_
                        d_2352_cf33_ = d_2349___mcc_h1_
                        d_2353_cf32_ = d_2348___mcc_h0_
                        d_2354_v23_: _dafny.Map
                        d_2354_v23_ = _dafny.Map({p0: p1})
                        d_2354_v23_ = (d_2354_v23_).set((d_2352_cf33_) >= (((d_2320_v0_)[False] if (False) in (d_2320_v0_) else p1)), p1)
                        (globalState).f1 = 209
                        d_2355_v24_: _dafny.Map
                        d_2355_v24_ = _dafny.Map({d_2353_cf32_: p2})
                        d_2352_cf33_ = ((d_2352_cf33_) * (p1)) * ((len(d_2355_v24_)) * (p1))
                        d_2356_v25_: _dafny.Map
                        d_2356_v25_ = _dafny.Map({p1: p1})
                        d_2357_v26_: _dafny.Seq
                        d_2357_v26_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "inqarh"))), len(d_2356_v25_)])
                        d_2358_v27_: _dafny.MultiSet
                        d_2358_v27_ = _dafny.MultiSet([(d_2357_v26_) + (d_2357_v26_), _dafny.SeqWithoutIsStrInference([p1]), d_2357_v26_, d_2357_v26_])
                        d_2358_v27_ = (d_2358_v27_) - (d_2358_v27_)
                    elif True:
                        d_2359___mcc_h3_ = source32_.cf31
                        d_2360_cf31_ = d_2359___mcc_h3_
                        d_2361_v28_: _dafny.Array
                        nw415_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                        d_2361_v28_ = nw415_
                        index412_ = default__.safeIndex(591, (d_2361_v28_).length(0))
                        (d_2361_v28_)[index412_] = _dafny.CodePoint('q')
                        index413_ = default__.safeIndex(591, (d_2361_v28_).length(0))
                        (d_2361_v28_)[index413_] = _dafny.CodePoint('t')
                        d_2362_v29_: _dafny.Array
                        def lambda109_(d_2363_p0_):
                            def lambda110_(d_2364_i4_):
                                return (d_2364_i4_) * ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_2363_p0_, (self).f6])), 218])).cardinality)

                            return lambda110_

                        init64_ = lambda109_(p0)
                        nw416_ = _dafny.Array(None, 28)
                        for i0_64_ in range(nw416_.length(0)):
                            nw416_[i0_64_] = init64_(i0_64_)
                        d_2362_v29_ = nw416_
                        d_2365_v30_: _dafny.Set
                        d_2365_v30_ = _dafny.Set({(self).f6, (self).f6})
                        index414_ = default__.safeIndex(384, (d_2362_v29_).length(0))
                        (d_2362_v29_)[index414_] = (len(d_2365_v30_)) - ((0) - (p1))
                        d_2366_v31_: _dafny.Seq
                        d_2366_v31_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                        d_2367_v32_: _dafny.Map
                        d_2367_v32_ = _dafny.Map({p0: p1})
                        d_2368_v33_: _dafny.Map
                        d_2368_v33_ = _dafny.Map({p1: p1})
                        d_2369_v34_: _dafny.MultiSet
                        d_2369_v34_ = _dafny.MultiSet([p1, len(d_2368_v33_)])
                        d_2370_v35_: _dafny.Map
                        d_2370_v35_ = _dafny.Map({len(p2): (d_2344_v19_).fm4(d_2366_v31_, 306, len(d_2367_v32_), d_2369_v34_, globalState)})
                        index415_ = default__.safeIndex(384, (d_2362_v29_).length(0))
                        (d_2362_v29_)[index415_] = (((d_2370_v35_)[p1] if (p1) in (d_2370_v35_) else p1)) - (default__.safeDivisionInt(30, p1))
                        d_2371_v36_: C2
                        nw417_ = C2()
                        nw417_.ctor__((_dafny.CodePoint('l') if p0 else _dafny.CodePoint('c')), (self).f6)
                        d_2371_v36_ = nw417_
                        d_2372_v37_: D20
                        d_2372_v37_ = D20_DC42(d_2344_v19_)
                        d_2373_v38_: _dafny.Map
                        d_2373_v38_ = _dafny.Map({((d_2371_v36_).fm8(True, (self).f6, p0, globalState)) + (p1): (d_2372_v37_).cf55})
                        d_2373_v38_ = (d_2373_v38_).set((d_2344_v19_).fm4(d_2366_v31_, p1, p1, d_2369_v34_, globalState), d_2344_v19_)
                    d_2374_v39_: _dafny.Map
                    d_2374_v39_ = _dafny.Map({(_dafny.MultiSet([not(p0), p0, (self).f6, (self).f6, (self).f6])).cardinality: p1})
                    d_2375_v40_: _dafny.Set
                    d_2375_v40_ = _dafny.Set({p1, len(_dafny.Map({p1: p0}))})
                    d_2376_v41_: _dafny.Map
                    d_2376_v41_ = _dafny.Map({155: _dafny.SeqWithoutIsStrInference([p1])})
                    d_2377_v42_: D19
                    d_2377_v42_ = D19_DC40(p0, len(d_2375_v40_), d_2376_v41_)
                    d_2374_v39_ = (d_2374_v39_).set(p1, (d_2377_v42_).cf52)
                    d_2378_v43_: _dafny.Array
                    nw418_ = _dafny.Array(_dafny.Array(None, 0), 13)
                    d_2378_v43_ = nw418_
                    d_2379_v44_: _dafny.Array
                    nw419_ = _dafny.Array(None, 1)
                    nw419_[int(0)] = p1
                    d_2379_v44_ = nw419_
                    index416_ = default__.safeIndex(481, (d_2378_v43_).length(0))
                    (d_2378_v43_)[index416_] = d_2379_v44_
                    index417_ = default__.safeIndex(16, (d_2379_v44_).length(0))
                    (d_2379_v44_)[index417_] = (797) * (p1)
                    d_2380_v45_: _dafny.Seq
                    d_2380_v45_ = _dafny.SeqWithoutIsStrInference([p0, (self).f6, p0])
                    index418_ = default__.safeIndex(481, (d_2378_v43_).length(0))
                    index419_ = default__.safeIndex(16, (d_2379_v44_).length(0))
                    rhs470_ = d_2379_v44_
                    rhs471_ = (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2381_i5_ in range(default__.abs(946))])) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhbetwjr"))) + (_dafny.SeqWithoutIsStrInference([(self).f5 for d_2382_i6_ in range(default__.abs(443))])))
                    rhs472_ = default__.safeDivisionInt((0) - (p1), p1)
                    rhs473_ = d_2380_v45_
                    lhs313_ = d_2378_v43_
                    lhs314_ = default__.safeIndex(481, (d_2378_v43_).length(0))
                    lhs315_ = d_2379_v44_
                    lhs316_ = default__.safeIndex(16, (d_2379_v44_).length(0))
                    lhs313_[lhs314_] = rhs470_
                    r0 = rhs471_
                    lhs315_[lhs316_] = rhs472_
                    d_2380_v45_ = rhs473_
                    arr4_ = self.f7
                    index420_ = default__.safeIndex(369, (self.f7).length(0))
                    arr4_[index420_] = (p0 if p0 else (d_2380_v45_)[default__.safeIndex((d_2379_v44_)[default__.safeIndex(16, (d_2379_v44_).length(0))], len(d_2380_v45_))])
                    arr5_ = self.f7
                    index421_ = default__.safeIndex(369, (self.f7).length(0))
                    arr5_[index421_] = p0
                    pass
            pass
        d_2383_v46_: _dafny.Seq
        d_2384_v47_: bool
        d_2385_v48_: bool
        out50_: _dafny.Seq
        out51_: bool
        out52_: bool
        out50_, out51_, out52_ = (self).m6(((self).f6) == ((self).f6), (p1) < (p1), globalState)
        d_2383_v46_ = out50_
        d_2384_v47_ = out51_
        d_2385_v48_ = out52_
        (globalState).f1 = p1
        r0 = d_2384_v47_
        r0 = d_2385_v48_
        r1 = p2
        return r0, r1

    def m3(self, globalState):
        r0: int = int(0)
        d_2386_v0_: int
        d_2386_v0_ = 210
        d_2387_v1_: _dafny.Map
        d_2387_v1_ = _dafny.Map({d_2386_v0_: (self).f6})
        d_2388_v2_: _dafny.Seq
        d_2388_v2_ = _dafny.SeqWithoutIsStrInference([default__.fm33(globalState), d_2387_v1_])
        d_2389_v3_: _dafny.Map
        d_2389_v3_ = (d_2388_v2_)[default__.safeIndex(d_2386_v0_, len(d_2388_v2_))]
        source33_ = d_2389_v3_
        d_2390___mcc_h0_ = source33_
        d_2391_cf12_ = d_2390___mcc_h0_
        d_2392_v4_: _dafny.Array
        nw420_ = _dafny.Array(None, 9)
        nw420_[int(0)] = d_2386_v0_
        nw420_[int(1)] = d_2386_v0_
        nw420_[int(2)] = d_2386_v0_
        nw420_[int(3)] = d_2386_v0_
        nw420_[int(4)] = d_2386_v0_
        nw420_[int(5)] = d_2386_v0_
        nw420_[int(6)] = d_2386_v0_
        nw420_[int(7)] = d_2386_v0_
        nw420_[int(8)] = d_2386_v0_
        d_2392_v4_ = nw420_
        d_2393_v5_: _dafny.Array
        d_2393_v5_ = d_2392_v4_
        d_2394_v6_: D12
        d_2394_v6_ = D12_DC23(d_2393_v5_)
        pat_let_tv32_ = d_2393_v5_
        pat_let_tv33_ = d_2392_v4_
        def iife160_(_pat_let43_0):
            def iife161_(d_2395_dt__update__tmp_h0_):
                def iife162_(_pat_let44_0):
                    def iife163_(d_2396_dt__update_hcf29_h0_):
                        return D12_DC23(d_2396_dt__update_hcf29_h0_)
                    return iife163_(_pat_let44_0)
                return iife162_(pat_let_tv32_)
            return iife161_(_pat_let43_0)
        def iife159_(_pat_let42_0):
            def iife164_(d_2397_dt__update__tmp_h1_):
                def iife165_(_pat_let45_0):
                    def iife166_(d_2398_dt__update_hcf29_h1_):
                        return D12_DC23(d_2398_dt__update_hcf29_h1_)
                    return iife166_(_pat_let45_0)
                return iife165_(pat_let_tv33_)
            return iife164_(_pat_let42_0)
        source34_ = iife159_((iife160_(d_2394_v6_) if (self).f6 else d_2394_v6_))
        if source34_.is_DC22:
            d_2399___mcc_h1_ = source34_.cf24
            d_2400___mcc_h2_ = source34_.cf25
            d_2401___mcc_h3_ = source34_.cf26
            d_2402___mcc_h4_ = source34_.cf27
            d_2403___mcc_h5_ = source34_.cf28
            d_2404_cf28_ = d_2403___mcc_h5_
            d_2405_cf27_ = d_2402___mcc_h4_
            d_2406_cf26_ = d_2401___mcc_h3_
            d_2407_cf25_ = d_2400___mcc_h2_
            d_2408_cf24_ = d_2399___mcc_h1_
            d_2409_v7_: _dafny.Map
            d_2409_v7_ = _dafny.Map({(0) - (d_2386_v0_): d_2386_v0_})
            d_2410_v8_: _dafny.Seq
            d_2410_v8_ = _dafny.SeqWithoutIsStrInference([len(d_2409_v7_)])
            d_2411_v9_: C3
            nw421_ = C3()
            nw421_.ctor__(True, d_2405_cf27_)
            d_2411_v9_ = nw421_
            d_2412_v10_: _dafny.Set
            d_2412_v10_ = _dafny.Set({d_2411_v9_, d_2411_v9_, d_2411_v9_})
            d_2413_v11_: _dafny.MultiSet
            d_2413_v11_ = _dafny.MultiSet([d_2410_v8_, d_2410_v8_])
            d_2414_v12_: _dafny.Map
            d_2414_v12_ = _dafny.Map({d_2413_v11_: d_2386_v0_})
            rhs474_ = (((d_2410_v8_) + ((d_2410_v8_).set(default__.safeIndex(d_2386_v0_, len(d_2410_v8_)), d_2386_v0_))).set(default__.safeIndex(((d_2414_v12_)[d_2413_v11_] if (d_2413_v11_) in (d_2414_v12_) else d_2386_v0_), len((d_2410_v8_) + ((d_2410_v8_).set(default__.safeIndex(d_2386_v0_, len(d_2410_v8_)), d_2386_v0_)))), 672) if (d_2412_v10_).issubset(d_2412_v10_) else d_2410_v8_)
            rhs475_ = d_2405_cf27_
            d_2410_v8_ = rhs474_
            d_2405_cf27_ = rhs475_
            d_2415_v13_: _dafny.Array
            nw422_ = _dafny.Array(_dafny.MultiSet({}), 12)
            d_2415_v13_ = nw422_
            d_2416_v14_: D18
            d_2416_v14_ = D18_DC36(d_2415_v13_)
            d_2417_v15_: _dafny.Map
            d_2417_v15_ = _dafny.Map({d_2416_v14_: (d_2386_v0_) - (-10)})
            d_2417_v15_ = (d_2417_v15_).set(d_2416_v14_, default__.safeDivisionInt(d_2386_v0_, d_2386_v0_))
            d_2386_v0_ = (len(_dafny.Map({d_2392_v4_: d_2386_v0_}))) + (d_2386_v0_)
            index422_ = default__.safeIndex(267, (d_2392_v4_).length(0))
            (d_2392_v4_)[index422_] = d_2386_v0_
            index423_ = default__.safeIndex(267, (d_2392_v4_).length(0))
            (d_2392_v4_)[index423_] = default__.safeDivisionInt(d_2386_v0_, default__.safeDivisionInt(len(d_2404_cf28_), len(d_2410_v8_)))
        elif source34_.is_DC23:
            d_2418___mcc_h6_ = source34_.cf29
            d_2419_cf29_ = d_2418___mcc_h6_
            d_2420_v16_: str
            d_2420_v16_ = _dafny.CodePoint('e')
            d_2420_v16_ = (self).f5
            d_2421_v17_: _dafny.Set
            d_2421_v17_ = _dafny.Set({(self).f6})
            d_2421_v17_ = d_2421_v17_
            d_2422_v18_: _dafny.Seq
            d_2422_v18_ = _dafny.SeqWithoutIsStrInference([d_2386_v0_, d_2386_v0_])
            d_2423_v19_: _dafny.MultiSet
            d_2423_v19_ = _dafny.MultiSet([(d_2422_v18_)[default__.safeIndex(d_2386_v0_, len(d_2422_v18_))]])
            d_2423_v19_ = default__.fm32(globalState)
            d_2424_v20_: bool
            d_2424_v20_ = True
            d_2425_v21_: _dafny.Map
            d_2425_v21_ = _dafny.Map({d_2386_v0_: d_2420_v16_})
            d_2426_v22_: _dafny.Seq
            d_2426_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryhpyo"))
            d_2427_v23_: _dafny.Map
            d_2427_v23_ = _dafny.Map({not((self).f6): ((d_2425_v21_)[len(d_2426_v22_)] if (len(d_2426_v22_)) in (d_2425_v21_) else (self).f5)})
            d_2428_v24_: _dafny.Set
            d_2428_v24_ = _dafny.Set({d_2427_v23_})
            d_2429_v25_: _dafny.Map
            d_2429_v25_ = _dafny.Map({d_2427_v23_: (self).f6})
            d_2430_v27_: _dafny.MultiSet
            d_2430_v27_ = _dafny.MultiSet([not((self).f6), (self).f6])
            def iife167_():
                coll75_ = _dafny.Set()
                compr_75_: _dafny.Map
                for compr_75_ in (d_2429_v25_).keys.Elements:
                    d_2431_v26_: _dafny.Map = compr_75_
                    if (d_2431_v26_) in (d_2429_v25_):
                        coll75_ = coll75_.union(_dafny.Set([d_2431_v26_]))
                return _dafny.Set(coll75_)
            rhs476_ = (self).f6
            rhs477_ = (iife167_()
            ) | (d_2428_v24_)
            rhs478_ = (d_2430_v27_) - (_dafny.MultiSet([not(default__.fm0(d_2424_v20_, globalState))]))
            lhs317_ = globalState
            d_2424_v20_ = rhs476_
            d_2428_v24_ = rhs477_
            lhs317_.f4 = rhs478_
        elif True:
            d_2432___mcc_h7_ = source34_.cf23
            d_2433_cf23_ = d_2432___mcc_h7_
            (globalState).f1 = d_2386_v0_
            d_2434_v28_: C5
            nw423_ = C5()
            nw423_.ctor__(len((_dafny.Map({(self).f6: (self).f6})) | (_dafny.Map({(self).f6: (self).f6}))), (self).f5, (self).f6)
            d_2434_v28_ = nw423_
            d_2435_v29_: D6
            d_2435_v29_ = D6_DC14((self).f6)
            d_2436_v30_: _dafny.Array
            def lambda111_(d_2437_i0_):
                return _dafny.MultiSet([(self).f6])

            init65_ = lambda111_
            nw424_ = _dafny.Array(None, 11)
            for i0_65_ in range(nw424_.length(0)):
                nw424_[i0_65_] = init65_(i0_65_)
            d_2436_v30_ = nw424_
            d_2438_v31_: D18
            d_2438_v31_ = D18_DC36(d_2436_v30_)
            d_2439_v32_: _dafny.MultiSet
            d_2439_v32_ = _dafny.MultiSet([d_2438_v31_, d_2438_v31_, D18_DC36(d_2436_v30_), d_2438_v31_, (d_2438_v31_ if (self).f6 else d_2438_v31_)])
            rhs479_ = default__.fm64((self).f5, globalState)
            rhs480_ = (_dafny.MultiSet([d_2438_v31_])) | (d_2439_v32_)
            d_2435_v29_ = rhs479_
            d_2439_v32_ = rhs480_
            d_2440_v33_: _dafny.Set
            d_2440_v33_ = _dafny.Set({d_2386_v0_, d_2434_v28_.f15})
            d_2441_v34_: _dafny.Map
            d_2441_v34_ = _dafny.Map({d_2386_v0_: d_2440_v33_})
            d_2442_v35_: _dafny.Seq
            d_2442_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wae"))
            d_2441_v34_ = (d_2441_v34_).set(d_2386_v0_, (d_2440_v33_) | (_dafny.Set({len(d_2442_v35_), d_2434_v28_.f15})))
        d_2443_v36_: bool
        d_2443_v36_ = True
        d_2444_v37_: _dafny.Seq
        d_2444_v37_ = _dafny.SeqWithoutIsStrInference([d_2386_v0_])
        d_2445_v38_: _dafny.MultiSet
        d_2445_v38_ = _dafny.MultiSet([True])
        d_2443_v36_ = (self).fm6((d_2444_v37_) + (_dafny.SeqWithoutIsStrInference([(d_2445_v38_).cardinality, d_2386_v0_])), globalState)
        arr6_ = self.f7
        index424_ = default__.safeIndex(173, (self.f7).length(0))
        arr6_[index424_] = d_2443_v36_
        arr7_ = self.f7
        index425_ = default__.safeIndex(173, (self.f7).length(0))
        arr7_[index425_] = (self).f6
        arr8_ = self.f7
        index426_ = default__.safeIndex(173, (self.f7).length(0))
        arr8_[index426_] = (self.f7)[default__.safeIndex(173, (self.f7).length(0))]
        d_2446_v39_: _dafny.MultiSet
        d_2446_v39_ = _dafny.MultiSet([(self).f6])
        d_2447_v40_: _dafny.Seq
        d_2447_v40_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6, (self).f6])
        d_2448_v41_: C6
        nw425_ = C6()
        nw425_.ctor__()
        d_2448_v41_ = nw425_
        d_2449_v42_: _dafny.Map
        d_2449_v42_ = _dafny.Map({d_2448_v41_: (self).f6})
        d_2450_v43_: _dafny.Map
        d_2450_v43_ = _dafny.Map({len(d_2449_v42_): d_2386_v0_})
        d_2451_v44_: _dafny.Set
        d_2451_v44_ = _dafny.Set({d_2446_v39_, _dafny.MultiSet(d_2447_v40_), _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([(self).f6])).set(default__.safeIndex(d_2386_v0_, len(_dafny.SeqWithoutIsStrInference([(self).f6]))), (self).f6)).set(default__.safeIndex(((d_2450_v43_)[-249] if (-249) in (d_2450_v43_) else d_2386_v0_), len((_dafny.SeqWithoutIsStrInference([(self).f6])).set(default__.safeIndex(d_2386_v0_, len(_dafny.SeqWithoutIsStrInference([(self).f6]))), (self).f6))), (self).f6))})
        d_2452_v45_: C7
        nw426_ = C7()
        nw426_.ctor__(len(d_2451_v44_), default__.safeModuloInt(d_2386_v0_, d_2386_v0_), (self).f5, (self).f6)
        d_2452_v45_ = nw426_
        if (946) > (800):
            d_2453_v46_: bool
            d_2453_v46_ = True
            d_2453_v46_ = ((self).f6) and ((self).f6)
            arr9_ = self.f7
            index427_ = default__.safeIndex(644, (self.f7).length(0))
            arr9_[index427_] = (self).f6
            d_2454_v47_: _dafny.Seq
            d_2454_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkmwkypd"))
            d_2455_v48_: _dafny.Seq
            d_2455_v48_ = _dafny.SeqWithoutIsStrInference([(d_2452_v45_).f9])
            d_2456_v49_: _dafny.Map
            d_2456_v49_ = _dafny.Map({d_2386_v0_: d_2455_v48_})
            d_2457_v50_: _dafny.MultiSet
            d_2457_v50_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_2386_v0_, d_2452_v45_.f10])), (d_2452_v45_).f9])
            arr10_ = self.f7
            index428_ = default__.safeIndex(644, (self.f7).length(0))
            rhs481_ = (d_2453_v46_ if True else (d_2448_v41_).fm6(d_2455_v48_, globalState))
            rhs482_ = 111
            rhs483_ = ((self).f6 if ((self).f6 if not((D19_DC40((self).f6, (d_2452_v45_).f9, d_2456_v49_)).cf51) else (self).f6) else not((self).f6))
            rhs484_ = ((d_2454_v47_) + ((_dafny.SeqWithoutIsStrInference([(self).f5 for d_2458_i1_ in range(default__.abs(113))])).set(default__.safeIndex((d_2452_v45_).fm4(d_2447_v40_, d_2386_v0_, d_2452_v45_.f10, d_2457_v50_, globalState), len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2459_i1_ in range(default__.abs(113))]))), (self).f5))) + (d_2454_v47_)
            lhs318_ = self.f7
            lhs319_ = default__.safeIndex(644, (self.f7).length(0))
            d_2453_v46_ = rhs481_
            d_2386_v0_ = rhs482_
            lhs318_[lhs319_] = rhs483_
            d_2454_v47_ = rhs484_
            d_2460_v51_: _dafny.Set
            d_2460_v51_ = _dafny.Set({d_2386_v0_, (d_2452_v45_).f9})
            d_2453_v46_ = (d_2460_v51_) == (d_2460_v51_)
            d_2461_v52_: _dafny.Map
            d_2461_v52_ = _dafny.Map({default__.fm0(d_2453_v46_, globalState): False})
            d_2462_v53_: D21
            d_2462_v53_ = D21_DC45(d_2457_v50_)
            d_2463_v54_: _dafny.Map
            d_2463_v54_ = _dafny.Map({(d_2446_v39_).set((self.f7)[default__.safeIndex(644, (self.f7).length(0))], default__.abs(len(d_2461_v52_))): ((d_2462_v53_).cf60).isdisjoint(d_2457_v50_)})
            d_2463_v54_ = (d_2463_v54_).set(d_2446_v39_, not(not((self.f7)[default__.safeIndex(644, (self.f7).length(0))])))
            d_2464_v55_: _dafny.Seq
            d_2465_v56_: bool
            d_2466_v57_: bool
            out53_: _dafny.Seq
            out54_: bool
            out55_: bool
            out53_, out54_, out55_ = (self).m6((self.f7)[default__.safeIndex(644, (self.f7).length(0))], (self.f7)[default__.safeIndex(644, (self.f7).length(0))], globalState)
            d_2464_v55_ = out53_
            d_2465_v56_ = out54_
            d_2466_v57_ = out55_
        elif True:
            if True:
                d_2467_v58_: _dafny.Map
                d_2467_v58_ = _dafny.Map({default__.fm1(d_2452_v45_.f10, not((self).f6), d_2452_v45_.f10, globalState): d_2386_v0_})
                d_2468_v59_: _dafny.Seq
                d_2468_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "deayq"))
                d_2469_v60_: _dafny.Map
                d_2469_v60_ = _dafny.Map({(self).f6: d_2452_v45_.f10})
                d_2467_v58_ = (d_2467_v58_).set((d_2468_v59_) + (d_2468_v59_), ((d_2469_v60_)[(self).f6] if ((self).f6) in (d_2469_v60_) else d_2386_v0_))
                d_2470_v61_: C1
                nw427_ = C1()
                nw427_.ctor__()
                d_2470_v61_ = nw427_
                d_2471_v62_: bool
                d_2471_v62_ = True
                d_2471_v62_ = not ((self).f6) or (not((self).f6))
                d_2472_v63_: _dafny.Seq
                d_2472_v63_ = _dafny.SeqWithoutIsStrInference([638])
                d_2473_v64_: _dafny.Seq
                d_2473_v64_ = _dafny.SeqWithoutIsStrInference([d_2472_v63_])
                rhs485_ = ((d_2473_v64_)[default__.safeIndex(len(d_2472_v63_), len(d_2473_v64_))]) not in (d_2473_v64_)
                rhs486_ = d_2386_v0_
                rhs487_ = (d_2452_v45_).f9
                lhs320_ = d_2452_v45_
                lhs321_ = globalState
                d_2471_v62_ = rhs485_
                lhs320_.f10 = rhs486_
                lhs321_.f1 = rhs487_
                d_2474_v65_: _dafny.Set
                d_2474_v65_ = _dafny.Set({(self).f6})
                d_2475_v66_: _dafny.Map
                d_2475_v66_ = _dafny.Map({(self).fm6((d_2472_v63_).set(default__.safeIndex(d_2452_v45_.f10, len(d_2472_v63_)), (d_2452_v45_).f9), globalState): d_2474_v65_})
                d_2476_v67_: _dafny.Seq
                d_2476_v67_ = _dafny.SeqWithoutIsStrInference([d_2474_v65_, ((d_2475_v66_)[True] if (True) in (d_2475_v66_) else d_2474_v65_)])
                d_2477_v68_: _dafny.Map
                d_2477_v68_ = _dafny.Map({(self).f5: self.f7})
                d_2478_v69_: _dafny.Map
                d_2478_v69_ = _dafny.Map({d_2476_v67_: len((d_2477_v68_) | (d_2477_v68_))})
                d_2478_v69_ = (d_2478_v69_).set(d_2476_v67_, d_2386_v0_)
            elif True:
                arr11_ = self.f7
                index429_ = default__.safeIndex(585, (self.f7).length(0))
                arr11_[index429_] = not(((d_2452_v45_).f9) >= ((d_2448_v41_).fm7((self).f6, globalState)))
                arr12_ = self.f7
                index430_ = default__.safeIndex(585, (self.f7).length(0))
                arr12_[index430_] = default__.fm0((self).f6, globalState)
                d_2479_v70_: _dafny.Seq
                d_2479_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "baymc"))
                d_2480_v71_: _dafny.Array
                nw428_ = _dafny.Array(False, 7)
                d_2480_v71_ = nw428_
                d_2481_v72_: _dafny.Map
                d_2481_v72_ = _dafny.Map({d_2452_v45_.f10: d_2480_v71_})
                d_2482_v73_: _dafny.Seq
                d_2482_v73_ = _dafny.SeqWithoutIsStrInference([d_2452_v45_.f10, d_2452_v45_.f10, (0) - (len(d_2479_v70_))])
                d_2483_v74_: _dafny.MultiSet
                d_2483_v74_ = _dafny.MultiSet([(d_2452_v45_).f9])
                d_2484_v75_: _dafny.Map
                d_2484_v75_ = _dafny.Map({d_2452_v45_.f10: _dafny.SeqWithoutIsStrInference([((d_2483_v74_)[d_2452_v45_.f10] if (d_2452_v45_.f10) in (d_2483_v74_) else (d_2452_v45_).f9)])})
                d_2485_v76_: _dafny.MultiSet
                d_2485_v76_ = _dafny.MultiSet([len(d_2484_v75_), -162])
                d_2486_v77_: _dafny.Array
                nw429_ = _dafny.Array(None, 24)
                nw429_[int(0)] = (0) - (((d_2452_v45_).f9 if (self).f6 else d_2452_v45_.f10))
                nw429_[int(1)] = (d_2452_v45_).f9
                nw429_[int(2)] = (d_2452_v45_).f9
                nw429_[int(3)] = (len((d_2479_v70_).set(default__.safeIndex((d_2452_v45_).f9, len(d_2479_v70_)), (self).f5))) * ((d_2452_v45_).f9)
                nw429_[int(4)] = len(((d_2481_v72_).set(d_2452_v45_.f10, d_2480_v71_)) | (d_2481_v72_))
                nw429_[int(5)] = ((d_2452_v45_).f9) - (-572)
                nw429_[int(6)] = len(d_2482_v73_)
                nw429_[int(7)] = (d_2452_v45_).f9
                nw429_[int(8)] = (d_2452_v45_).f9
                nw429_[int(9)] = (d_2452_v45_).f9
                nw429_[int(10)] = d_2452_v45_.f10
                nw429_[int(11)] = d_2452_v45_.f10
                nw429_[int(12)] = (d_2452_v45_).fm4(d_2447_v40_, ((d_2446_v39_)[(self).f6] if ((self).f6) in (d_2446_v39_) else len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_2447_v40_ for d_2487_i3_ in range(default__.abs(811))])) for d_2488_i2_ in range(default__.abs(162))]))), d_2452_v45_.f10, d_2485_v76_, globalState)
                nw429_[int(13)] = (d_2452_v45_).fm7((self.f7)[default__.safeIndex(585, (self.f7).length(0))], globalState)
                nw429_[int(14)] = (d_2483_v74_).cardinality
                nw429_[int(15)] = -739
                nw429_[int(16)] = default__.fm3((self).f6, (self).f6, globalState)
                nw429_[int(17)] = d_2452_v45_.f10
                nw429_[int(18)] = (d_2452_v45_).f9
                nw429_[int(19)] = 221
                nw429_[int(20)] = ((d_2452_v45_).f9) * (len(_dafny.Map({(self.f7)[default__.safeIndex(585, (self.f7).length(0))]: d_2452_v45_.f10})))
                nw429_[int(21)] = len(_dafny.SeqWithoutIsStrInference([(self).f5 for d_2489_i4_ in range(default__.abs(738))]))
                nw429_[int(22)] = (d_2452_v45_).f9
                nw429_[int(23)] = len(d_2479_v70_)
                d_2486_v77_ = nw429_
                d_2490_v78_: _dafny.Set
                d_2490_v78_ = _dafny.Set({d_2452_v45_.f10, d_2452_v45_.f10})
                d_2491_v79_: _dafny.Seq
                d_2491_v79_ = _dafny.SeqWithoutIsStrInference([d_2490_v78_, d_2490_v78_])
                index431_ = default__.safeIndex(747, (d_2486_v77_).length(0))
                (d_2486_v77_)[index431_] = len(d_2491_v79_)
                index432_ = default__.safeIndex(747, (d_2486_v77_).length(0))
                (d_2486_v77_)[index432_] = default__.safeModuloInt((d_2452_v45_).f9, d_2452_v45_.f10)
                d_2386_v0_ = 135
                d_2492_v80_: _dafny.Array
                d_2492_v80_ = d_2486_v77_
                d_2493_v81_: _dafny.Map
                d_2493_v81_ = _dafny.Map({len(d_2447_v40_): d_2492_v80_})
                (d_2452_v45_).f10 = default__.safeModuloInt(((d_2486_v77_)[default__.safeIndex(747, (d_2486_v77_).length(0))]) - ((d_2486_v77_)[default__.safeIndex(747, (d_2486_v77_).length(0))]), len(d_2493_v81_))
                index433_ = default__.safeIndex(747, (d_2486_v77_).length(0))
                (d_2486_v77_)[index433_] = d_2452_v45_.f10
            d_2494_v82_: _dafny.MultiSet
            d_2494_v82_ = _dafny.MultiSet([(d_2452_v45_).f9])
            d_2495_v83_: _dafny.Map
            d_2495_v83_ = _dafny.Map({True: d_2494_v82_})
            (globalState).f1 = len((d_2495_v83_) | ((d_2495_v83_) | (d_2495_v83_)))
            arr13_ = self.f7
            index434_ = default__.safeIndex(609, (self.f7).length(0))
            arr13_[index434_] = ((self).f6) and ((d_2447_v40_)[default__.safeIndex(d_2452_v45_.f10, len(d_2447_v40_))])
            arr14_ = self.f7
            index435_ = default__.safeIndex(609, (self.f7).length(0))
            arr14_[index435_] = False
            d_2496_v84_: _dafny.Map
            d_2496_v84_ = _dafny.Map({True: (d_2452_v45_).f9})
            d_2497_v85_: D6
            d_2497_v85_ = D6_DC13(_dafny.SeqWithoutIsStrInference([(d_2452_v45_).f9]))
            d_2498_v86_: _dafny.Seq
            d_2498_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuafj"))
            d_2499_v87_: _dafny.Map
            d_2499_v87_ = _dafny.Map({(d_2497_v85_).cf14: len(d_2498_v86_)})
            rhs488_ = default__.fm28((self.f7)[default__.safeIndex(609, (self.f7).length(0))], (self.f7)[default__.safeIndex(609, (self.f7).length(0))], d_2499_v87_, globalState)
            rhs489_ = (d_2452_v45_).f9
            lhs322_ = globalState
            d_2496_v84_ = rhs488_
            lhs322_.f1 = rhs489_
            d_2500_v88_: _dafny.Array
            def lambda112_(d_2501_v45_):
                def lambda113_(d_2502_i5_):
                    return (d_2502_i5_) * (d_2501_v45_.f10)

                return lambda113_

            init66_ = lambda112_(d_2452_v45_)
            nw430_ = _dafny.Array(None, 5)
            for i0_66_ in range(nw430_.length(0)):
                nw430_[i0_66_] = init66_(i0_66_)
            d_2500_v88_ = nw430_
            def lambda114_(d_2503_v45_):
                def lambda115_(d_2504_i6_):
                    return (d_2504_i6_) * (d_2503_v45_.f10)

                return lambda115_

            init67_ = lambda114_(d_2452_v45_)
            nw431_ = _dafny.Array(None, 1)
            for i0_67_ in range(nw431_.length(0)):
                nw431_[i0_67_] = init67_(i0_67_)
            d_2500_v88_ = nw431_
        d_2505_v89_: bool
        d_2505_v89_ = True
        d_2505_v89_ = d_2505_v89_
        d_2505_v89_ = d_2505_v89_
        d_2506_v90_: _dafny.Map
        d_2506_v90_ = _dafny.Map({d_2505_v89_: not(True)})
        d_2507_v91_: C5
        nw432_ = C5()
        nw432_.ctor__((d_2446_v39_).cardinality, (self).f5, not((d_2452_v45_.f10) not in ((d_2450_v43_).set((d_2452_v45_).f9, len(default__.fm1(len(d_2506_v90_), (self).f6, d_2386_v0_, globalState))))))
        d_2507_v91_ = nw432_
        r0 = (d_2386_v0_) * (d_2507_v91_.f15)
        return r0

    def m1(self, p0, p1, p2, globalState):
        d_2508_v0_: str
        d_2508_v0_ = _dafny.CodePoint('l')
        d_2508_v0_ = (self).f5
        d_2509_v1_: _dafny.Seq
        d_2509_v1_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2510_v2_: D15
        d_2510_v2_ = D15_DC28((self).f6, (d_2509_v1_)[default__.safeIndex((0) - ((0) - (len(d_2509_v1_))), len(d_2509_v1_))])
        source35_ = d_2510_v2_
        if source35_.is_DC28:
            d_2511___mcc_h0_ = source35_.cf36
            d_2512___mcc_h1_ = source35_.cf37
            d_2513_cf37_ = d_2512___mcc_h1_
            d_2514_cf36_ = d_2511___mcc_h0_
            d_2515_v3_: int
            d_2515_v3_ = -712
            d_2516_v4_: _dafny.Seq
            d_2516_v4_ = _dafny.SeqWithoutIsStrInference([d_2509_v1_])
            rhs490_ = (d_2515_v3_) <= ((d_2515_v3_) + (d_2515_v3_))
            rhs491_ = ((d_2509_v1_) + (d_2509_v1_)) + ((d_2516_v4_)[default__.safeIndex(d_2515_v3_, len(d_2516_v4_))])
            rhs492_ = (d_2514_cf36_) and (d_2513_cf37_)
            d_2514_cf36_ = rhs490_
            d_2509_v1_ = rhs491_
            d_2513_cf37_ = rhs492_
            d_2514_cf36_ = d_2513_cf37_
            arr15_ = self.f7
            index436_ = default__.safeIndex(312, (self.f7).length(0))
            arr15_[index436_] = p0
            d_2517_v5_: D0
            d_2517_v5_ = D0_DC0(p2)
            d_2518_v6_: _dafny.Map
            d_2518_v6_ = _dafny.Map({d_2515_v3_: _dafny.MultiSet([len(_dafny.Set({(self).fm9(d_2515_v3_, d_2517_v5_, d_2514_cf36_, d_2515_v3_, globalState)}))])})
            d_2519_v7_: _dafny.MultiSet
            d_2519_v7_ = _dafny.MultiSet([d_2515_v3_])
            d_2520_v8_: _dafny.Seq
            d_2520_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgibkeo"))
            d_2521_v9_: D6
            d_2521_v9_ = D6_DC13(_dafny.SeqWithoutIsStrInference([d_2515_v3_ for d_2522_i0_ in range(default__.abs(158))]))
            d_2523_v10_: _dafny.Map
            d_2523_v10_ = _dafny.Map({(self).f6: p0})
            d_2524_v11_: D20
            d_2524_v11_ = D20_DC43(len((d_2521_v9_).cf14), ((d_2523_v10_)[p0] if (p0) in (d_2523_v10_) else (self).f6), 88)
            d_2525_v12_: D20
            d_2525_v12_ = D20_DC44(d_2524_v11_)
            d_2526_v13_: _dafny.Set
            d_2526_v13_ = _dafny.Set({d_2525_v12_})
            arr16_ = self.f7
            index437_ = default__.safeIndex(312, (self.f7).length(0))
            arr16_[index437_] = (((d_2518_v6_)[(0) - (d_2515_v3_)] if ((0) - (d_2515_v3_)) in (d_2518_v6_) else (_dafny.MultiSet([len(d_2520_v8_)])).set(689, default__.abs(len(d_2526_v13_))))).issubset(((d_2518_v6_)[d_2515_v3_] if (d_2515_v3_) in (d_2518_v6_) else d_2519_v7_))
            (globalState).f1 = d_2515_v3_
        elif source35_.is_DC27:
            d_2527___mcc_h2_ = source35_.cf35
            d_2528_cf35_ = d_2527___mcc_h2_
            if p0:
                d_2529_v14_: _dafny.Seq
                d_2529_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "giybh"))
                d_2530_v15_: int
                d_2530_v15_ = 752
                def iife168_(_pat_let46_0):
                    def iife169_(d_2531_dt__update__tmp_h0_):
                        def iife170_(_pat_let47_0):
                            def iife171_(d_2532_dt__update_hcf43_h0_):
                                return D17_DC34(d_2532_dt__update_hcf43_h0_)
                            return iife171_(_pat_let47_0)
                        return iife170_((self).f5)
                    return iife169_(_pat_let46_0)
                d_2508_v0_ = (iife168_(default__.fm65((self).f6, len(d_2529_v14_), d_2530_v15_, globalState))).cf43
                d_2533_v16_: C5
                nw433_ = C5()
                nw433_.ctor__(d_2530_v15_, (self).f5, (-195) < ((0) - (d_2530_v15_)))
                d_2533_v16_ = nw433_
                arr17_ = self.f7
                index438_ = default__.safeIndex(397, (self.f7).length(0))
                arr17_[index438_] = not (p0) or (p0)
                arr18_ = self.f7
                index439_ = default__.safeIndex(397, (self.f7).length(0))
                arr18_[index439_] = not (False) or ((d_2509_v1_) < (d_2509_v1_))
                d_2534_v17_: _dafny.Seq
                d_2534_v17_ = _dafny.SeqWithoutIsStrInference([d_2533_v16_.f15, len(d_2529_v14_)])
                d_2535_v18_: _dafny.Seq
                d_2535_v18_ = _dafny.SeqWithoutIsStrInference([d_2534_v17_])
                d_2536_v19_: C0
                nw434_ = C0()
                nw434_.ctor__(p0, d_2535_v18_)
                d_2536_v19_ = nw434_
                d_2537_v20_: C1
                nw435_ = C1()
                nw435_.ctor__()
                d_2537_v20_ = nw435_
                d_2538_v21_: _dafny.Map
                d_2538_v21_ = _dafny.Map({len(d_2529_v14_): (d_2536_v19_).f11})
                d_2539_v22_: _dafny.Map
                d_2539_v22_ = _dafny.Map({d_2509_v1_: d_2538_v21_})
                d_2540_v23_: _dafny.Map
                d_2540_v23_ = _dafny.Map({(d_2539_v22_).set(d_2509_v1_, d_2538_v21_): d_2536_v19_})
                arr19_ = self.f7
                index440_ = default__.safeIndex(397, (self.f7).length(0))
                rhs493_ = ((d_2540_v23_)[(d_2539_v22_) | (d_2539_v22_)] if ((d_2539_v22_) | (d_2539_v22_)) in (d_2540_v23_) else d_2536_v19_)
                rhs494_ = (d_2509_v1_)[default__.safeIndex(d_2533_v16_.f15, len(d_2509_v1_))]
                rhs495_ = d_2537_v20_
                lhs323_ = self.f7
                lhs324_ = default__.safeIndex(397, (self.f7).length(0))
                d_2536_v19_ = rhs493_
                lhs323_[lhs324_] = rhs494_
                d_2537_v20_ = rhs495_
                d_2541_v24_: _dafny.Array
                nw436_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2541_v24_ = nw436_
                d_2542_v25_: _dafny.Array
                nw437_ = _dafny.Array(int(0), 28)
                d_2542_v25_ = nw437_
                index441_ = default__.safeIndex(555, (d_2541_v24_).length(0))
                (d_2541_v24_)[index441_] = d_2542_v25_
                d_2543_v26_: _dafny.Map
                d_2543_v26_ = _dafny.Map({(self.f7)[default__.safeIndex(397, (self.f7).length(0))]: -212})
                d_2544_v27_: _dafny.Map
                d_2544_v27_ = _dafny.Map({_dafny.CodePoint('h'): d_2543_v26_})
                index442_ = default__.safeIndex(555, (d_2541_v24_).length(0))
                rhs496_ = ((_dafny.SeqWithoutIsStrInference([d_2533_v16_.f15])) + ((_dafny.SeqWithoutIsStrInference([d_2530_v15_])) + (d_2534_v17_))).set(default__.safeIndex(len(d_2544_v27_), len((_dafny.SeqWithoutIsStrInference([d_2533_v16_.f15])) + ((_dafny.SeqWithoutIsStrInference([d_2530_v15_])) + (d_2534_v17_)))), (d_2533_v16_.f15) - (len(d_2534_v17_)))
                rhs497_ = d_2542_v25_
                lhs325_ = d_2541_v24_
                lhs326_ = default__.safeIndex(555, (d_2541_v24_).length(0))
                d_2534_v17_ = rhs496_
                lhs325_[lhs326_] = rhs497_
            elif True:
                d_2545_v28_: D21
                d_2545_v28_ = D21_DC47((self).f6)
                d_2546_v29_: bool
                d_2546_v29_ = False
                d_2547_v30_: _dafny.Array
                nw438_ = _dafny.Array(int(0), 2)
                d_2547_v30_ = nw438_
                d_2548_v31_: int
                d_2548_v31_ = 893
                index443_ = default__.safeIndex(463, (d_2547_v30_).length(0))
                (d_2547_v30_)[index443_] = default__.safeModuloInt(d_2548_v31_, d_2548_v31_)
                d_2549_v32_: _dafny.Map
                d_2549_v32_ = _dafny.Map({(0) - ((self).fm7(p0, globalState)): (self).f6})
                d_2550_v33_: _dafny.Seq
                d_2550_v33_ = _dafny.SeqWithoutIsStrInference([self.f7])
                d_2551_v34_: _dafny.Seq
                d_2551_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eu"))
                pat_let_tv34_ = d_2548_v31_
                pat_let_tv35_ = d_2548_v31_
                index444_ = default__.safeIndex(463, (d_2547_v30_).length(0))
                def iife172_(_pat_let48_0):
                    def iife173_(d_2552_dt__update__tmp_h1_):
                        def iife174_(_pat_let49_0):
                            def iife175_(d_2553_dt__update_hcf64_h0_):
                                return D21_DC47(d_2553_dt__update_hcf64_h0_)
                            return iife175_(_pat_let49_0)
                        return iife174_((pat_let_tv34_) <= ((0) - (pat_let_tv35_)))
                    return iife173_(_pat_let48_0)
                rhs498_ = (d_2548_v31_ if not(((d_2549_v32_)[d_2548_v31_] if (d_2548_v31_) in (d_2549_v32_) else d_2546_v29_)) else (p1).cardinality)
                rhs499_ = iife172_(d_2545_v28_)
                rhs500_ = d_2546_v29_
                rhs501_ = (d_2550_v33_)[default__.safeIndex(d_2548_v31_, len(d_2550_v33_))]
                rhs502_ = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([d_2551_v34_, d_2551_v34_, d_2551_v34_, d_2551_v34_, d_2551_v34_])), (self).fm8(p2, d_2546_v29_, not(True), globalState))
                lhs327_ = globalState
                lhs328_ = self
                lhs329_ = d_2547_v30_
                lhs330_ = default__.safeIndex(463, (d_2547_v30_).length(0))
                lhs327_.f1 = rhs498_
                d_2545_v28_ = rhs499_
                d_2546_v29_ = rhs500_
                lhs328_.f7 = rhs501_
                lhs329_[lhs330_] = rhs502_
                arr20_ = self.f7
                index445_ = default__.safeIndex(174, (self.f7).length(0))
                arr20_[index445_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2554_i1_ in range(default__.abs(419))])) <= (d_2551_v34_)
                arr21_ = self.f7
                index446_ = default__.safeIndex(174, (self.f7).length(0))
                arr21_[index446_] = ((d_2509_v1_) + (d_2509_v1_)) <= (d_2509_v1_)
                d_2548_v31_ = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ff"))]))
                d_2555_v35_: _dafny.Map
                d_2555_v35_ = _dafny.Map({d_2548_v31_: d_2551_v34_})
                d_2556_v36_: _dafny.Set
                d_2556_v36_ = _dafny.Set({(d_2551_v34_).set(default__.safeIndex(d_2548_v31_, len(d_2551_v34_)), _dafny.CodePoint('f'))})
                d_2557_v37_: _dafny.Map
                d_2557_v37_ = _dafny.Map({d_2547_v30_: (d_2509_v1_)[default__.safeIndex(d_2548_v31_, len(d_2509_v1_))]})
                d_2551_v34_ = ((((d_2555_v35_)[(d_2547_v30_)[default__.safeIndex(463, (d_2547_v30_).length(0))]] if ((d_2547_v30_)[default__.safeIndex(463, (d_2547_v30_).length(0))]) in (d_2555_v35_) else d_2551_v34_)).set(default__.safeIndex(len(d_2556_v36_), len(((d_2555_v35_)[(d_2547_v30_)[default__.safeIndex(463, (d_2547_v30_).length(0))]] if ((d_2547_v30_)[default__.safeIndex(463, (d_2547_v30_).length(0))]) in (d_2555_v35_) else d_2551_v34_))), d_2508_v0_)) + ((_dafny.SeqWithoutIsStrInference([d_2508_v0_ for d_2558_i2_ in range(default__.abs(628))])).set(default__.safeIndex(len(d_2557_v37_), len(_dafny.SeqWithoutIsStrInference([d_2508_v0_ for d_2559_i2_ in range(default__.abs(628))]))), (default__.fm65((self).f6, 266, (d_2547_v30_)[default__.safeIndex(463, (d_2547_v30_).length(0))], globalState)).cf43))
                arr22_ = self.f7
                index447_ = default__.safeIndex(174, (self.f7).length(0))
                arr22_[index447_] = d_2546_v29_
            if p2:
                d_2560_v38_: int
                d_2560_v38_ = 788
                (globalState).f1 = (d_2560_v38_) - ((self).fm8(p2, p2, p2, globalState))
                d_2561_v39_: _dafny.Map
                d_2561_v39_ = _dafny.Map({p0: d_2560_v38_})
                d_2561_v39_ = (d_2561_v39_).set(False, d_2560_v38_)
                d_2562_v40_: _dafny.Seq
                d_2562_v40_ = _dafny.SeqWithoutIsStrInference([d_2560_v38_])
                (globalState).f1 = (self).fm4(d_2509_v1_, len(d_2562_v40_), (0) - (d_2560_v38_), _dafny.MultiSet([d_2560_v38_]), globalState)
                d_2563_v41_: C3
                nw439_ = C3()
                nw439_.ctor__(p2, p0)
                d_2563_v41_ = nw439_
                d_2564_v42_: _dafny.MultiSet
                d_2564_v42_ = _dafny.MultiSet([(0) - (d_2560_v38_)])
                d_2565_v43_: C7
                nw440_ = C7()
                nw440_.ctor__(d_2560_v38_, (d_2564_v42_).cardinality, (self).f5, p2)
                d_2565_v43_ = nw440_
            elif True:
                d_2508_v0_ = (self).f5
                arr23_ = self.f7
                index448_ = default__.safeIndex(36, (self.f7).length(0))
                arr23_[index448_] = (p0) == (False)
                d_2566_v44_: _dafny.Map
                d_2566_v44_ = _dafny.Map({default__.fm49(globalState): p2})
                d_2567_v45_: int
                d_2567_v45_ = 120
                d_2568_v46_: _dafny.MultiSet
                d_2568_v46_ = _dafny.MultiSet([d_2567_v45_])
                d_2569_v47_: T3
                nw441_ = C5()
                nw441_.ctor__((345 if (self).fm10(d_2528_cf35_, d_2566_v44_, 697, globalState) else (d_2568_v46_).cardinality), (self).f5, p2)
                d_2569_v47_ = nw441_
                d_2570_v48_: _dafny.Seq
                d_2570_v48_ = _dafny.SeqWithoutIsStrInference([-920, d_2567_v45_, 128, d_2567_v45_])
                arr24_ = self.f7
                index449_ = default__.safeIndex(36, (self.f7).length(0))
                rhs503_ = p0
                rhs504_ = len(d_2509_v1_)
                rhs505_ = default__.safeModuloInt((d_2567_v45_) - ((d_2570_v48_)[default__.safeIndex((0) - ((0) - (d_2567_v45_)), len(d_2570_v48_))]), ((d_2528_cf35_)[p2] if (p2) in (d_2528_cf35_) else d_2567_v45_))
                rhs506_ = d_2569_v47_
                lhs331_ = self.f7
                lhs332_ = default__.safeIndex(36, (self.f7).length(0))
                lhs333_ = globalState
                lhs334_ = globalState
                lhs331_[lhs332_] = rhs503_
                lhs333_.f1 = rhs504_
                lhs334_.f1 = rhs505_
                d_2569_v47_ = rhs506_
                d_2571_v49_: _dafny.Seq
                d_2571_v49_ = _dafny.SeqWithoutIsStrInference([(d_2570_v48_) + (_dafny.SeqWithoutIsStrInference([205 for d_2572_i3_ in range(default__.abs(-385))])), (_dafny.SeqWithoutIsStrInference([default__.fm3((d_2569_v47_).f6, p2, globalState), -206])) + (d_2570_v48_)])
                d_2573_v50_: _dafny.Seq
                d_2573_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdtgl"))
                d_2571_v49_ = default__.fm38((len(d_2573_v50_)) != (d_2567_v45_), p0, len(d_2573_v50_), d_2508_v0_, globalState)
                d_2567_v45_ = len(d_2570_v48_)
                (globalState).f1 = ((0) - (d_2567_v45_)) * (d_2567_v45_)
            d_2574_v51_: int
            d_2574_v51_ = 710
            d_2575_v52_: _dafny.MultiSet
            d_2575_v52_ = _dafny.MultiSet([422])
            d_2576_v53_: _dafny.Seq
            d_2576_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_2577_v54_: _dafny.Map
            d_2577_v54_ = _dafny.Map({(self).fm4(_dafny.SeqWithoutIsStrInference([False]), d_2574_v51_, d_2574_v51_, d_2575_v52_, globalState): d_2576_v53_})
            (globalState).f1 = (0) - (default__.safeModuloInt(644, (0) - (default__.safeModuloInt(len(d_2577_v54_), d_2574_v51_))))
            d_2578_v55_: _dafny.Array
            nw442_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_2578_v55_ = nw442_
            d_2579_v56_: _dafny.Map
            d_2579_v56_ = _dafny.Map({d_2574_v51_: (self).f6})
            d_2580_v58_: _dafny.Set
            d_2580_v58_ = _dafny.Set({d_2574_v51_, d_2574_v51_})
            d_2581_v59_: _dafny.Array
            nw443_ = _dafny.Array(None, 25)
            nw443_[int(0)] = d_2579_v56_
            nw443_[int(1)] = d_2579_v56_
            def iife176_():
                coll76_ = _dafny.Map()
                compr_76_: int
                for compr_76_ in (d_2580_v58_).Elements:
                    d_2582_v57_: int = compr_76_
                    if (d_2582_v57_) in (d_2580_v58_):
                        coll76_[(d_2582_v57_) + ((0) - (d_2574_v51_))] = p2
                return _dafny.Map(coll76_)
            nw443_[int(2)] = iife176_()
            
            nw443_[int(3)] = d_2579_v56_
            nw443_[int(4)] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2574_v51_])): p0})
            nw443_[int(5)] = d_2579_v56_
            nw443_[int(6)] = d_2579_v56_
            nw443_[int(7)] = ((_dafny.Map({d_2574_v51_: (self).f6}))).set(d_2574_v51_, (self).f6)
            nw443_[int(8)] = d_2579_v56_
            nw443_[int(9)] = d_2579_v56_
            nw443_[int(10)] = d_2579_v56_
            nw443_[int(11)] = d_2579_v56_
            nw443_[int(12)] = d_2579_v56_
            nw443_[int(13)] = d_2579_v56_
            nw443_[int(14)] = _dafny.Map({d_2574_v51_: p2})
            nw443_[int(15)] = d_2579_v56_
            nw443_[int(16)] = d_2579_v56_
            nw443_[int(17)] = d_2579_v56_
            nw443_[int(18)] = d_2579_v56_
            nw443_[int(19)] = d_2579_v56_
            nw443_[int(20)] = _dafny.Map({d_2574_v51_: p2})
            nw443_[int(21)] = d_2579_v56_
            nw443_[int(22)] = d_2579_v56_
            nw443_[int(23)] = (d_2579_v56_).set(len(_dafny.Map({d_2574_v51_: p2})), p2)
            nw443_[int(24)] = d_2579_v56_
            d_2581_v59_ = nw443_
            index450_ = default__.safeIndex(32, (d_2578_v55_).length(0))
            (d_2578_v55_)[index450_] = d_2581_v59_
            index451_ = default__.safeIndex(32, (d_2578_v55_).length(0))
            (d_2578_v55_)[index451_] = d_2581_v59_
        elif True:
            d_2583___mcc_h3_ = source35_.cf38
            d_2584_cf38_ = d_2583___mcc_h3_
            d_2585_v60_: _dafny.Map
            d_2585_v60_ = _dafny.Map({-135: d_2509_v1_})
            d_2586_v61_: int
            d_2586_v61_ = 202
            (globalState).f1 = default__.safeModuloInt(len(d_2585_v60_), d_2586_v61_)
            d_2587_v62_: _dafny.Array
            nw444_ = _dafny.Array(int(0), 8)
            d_2587_v62_ = nw444_
            index452_ = default__.safeIndex(14, (d_2587_v62_).length(0))
            (d_2587_v62_)[index452_] = d_2586_v61_
            d_2588_v63_: D20
            d_2588_v63_ = D20_DC43(d_2586_v61_, p2, d_2586_v61_)
            index453_ = default__.safeIndex(14, (d_2587_v62_).length(0))
            rhs507_ = self.f7
            rhs508_ = d_2586_v61_
            rhs509_ = (d_2588_v63_).cf58
            lhs335_ = self
            lhs336_ = d_2587_v62_
            lhs337_ = default__.safeIndex(14, (d_2587_v62_).length(0))
            lhs335_.f7 = rhs507_
            lhs336_[lhs337_] = rhs508_
            d_2586_v61_ = rhs509_
            d_2586_v61_ = ((0) - ((self).fm7(True, globalState)) if p2 else 956)
            d_2589_v64_: _dafny.Map
            d_2589_v64_ = _dafny.Map({(0) - ((d_2587_v62_)[default__.safeIndex(14, (d_2587_v62_).length(0))]): (d_2509_v1_)[default__.safeIndex(d_2586_v61_, len(d_2509_v1_))]})
            if (p2) and (((d_2589_v64_)[176] if (176) in (d_2589_v64_) else p2)):
                (self).f7 = self.f7
                d_2509_v1_ = d_2509_v1_
                d_2590_v65_: bool
                d_2590_v65_ = False
                d_2590_v65_ = not(not(not((self).f6)))
                d_2591_v66_: _dafny.Array
                d_2591_v66_ = self.f7
                d_2591_v66_ = d_2591_v66_
                d_2590_v65_ = p2
            elif True:
                (globalState).f1 = d_2586_v61_
                arr25_ = self.f7
                index454_ = default__.safeIndex(996, (self.f7).length(0))
                arr25_[index454_] = p2
                d_2592_v67_: _dafny.Map
                d_2592_v67_ = _dafny.Map({(self).f6: 118})
                d_2593_v68_: _dafny.MultiSet
                d_2593_v68_ = _dafny.MultiSet([(self).f5, (self).f5])
                arr26_ = self.f7
                index455_ = default__.safeIndex(996, (self.f7).length(0))
                rhs510_ = (d_2509_v1_).set(default__.safeIndex(d_2586_v61_, len(d_2509_v1_)), (self).f6)
                rhs511_ = d_2508_v0_
                rhs512_ = p2
                rhs513_ = d_2592_v67_
                rhs514_ = len((self).fm11(d_2586_v61_, p0, (d_2593_v68_).cardinality, globalState))
                lhs338_ = self.f7
                lhs339_ = default__.safeIndex(996, (self.f7).length(0))
                lhs340_ = globalState
                d_2509_v1_ = rhs510_
                d_2508_v0_ = rhs511_
                lhs338_[lhs339_] = rhs512_
                d_2592_v67_ = rhs513_
                lhs340_.f1 = rhs514_
                (globalState).f1 = (d_2587_v62_)[default__.safeIndex(14, (d_2587_v62_).length(0))]
                d_2594_v69_: _dafny.Map
                d_2594_v69_ = _dafny.Map({(self.f7)[default__.safeIndex(996, (self.f7).length(0))]: (self.f7)[default__.safeIndex(996, (self.f7).length(0))]})
                d_2594_v69_ = d_2594_v69_
                (globalState).f1 = default__.safeDivisionInt((len(_dafny.SeqWithoutIsStrInference([(d_2587_v62_)[default__.safeIndex(14, (d_2587_v62_).length(0))]])) if p0 else d_2586_v61_), (p1).cardinality)
        d_2595_v70_: int
        d_2595_v70_ = 712
        (globalState).f1 = d_2595_v70_
        d_2596_v71_: _dafny.Seq
        d_2596_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fadnxqgcp"))
        d_2597_v72_: _dafny.Seq
        d_2597_v72_ = _dafny.SeqWithoutIsStrInference([(d_2596_v71_) + (d_2596_v71_), d_2596_v71_])
        d_2597_v72_ = (d_2597_v72_) + (d_2597_v72_)
        d_2598_v73_: _dafny.MultiSet
        d_2598_v73_ = _dafny.MultiSet([d_2596_v71_])
        d_2599_v74_: _dafny.Map
        d_2599_v74_ = _dafny.Map({not ((self).f6) or ((d_2509_v1_)[default__.safeIndex(d_2595_v70_, len(d_2509_v1_))]): _dafny.SeqWithoutIsStrInference([(self).f5 for d_2600_i4_ in range(default__.abs(261))])})
        rhs515_ = d_2595_v70_
        rhs516_ = ((d_2599_v74_)[(self).f6] if ((self).f6) in (d_2599_v74_) else d_2596_v71_)
        rhs517_ = (d_2595_v70_) - (d_2595_v70_)
        rhs518_ = (d_2598_v73_).intersection(((_dafny.MultiSet([d_2596_v71_, d_2596_v71_])).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anuxkuyu")), default__.abs(-918))) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbkjrf"))])))
        lhs341_ = globalState
        lhs341_.f1 = rhs515_
        d_2596_v71_ = rhs516_
        d_2595_v70_ = rhs517_
        d_2598_v73_ = rhs518_
        d_2601_v75_: _dafny.Set
        d_2601_v75_ = _dafny.Set({self.f7, self.f7})
        d_2601_v75_ = (d_2601_v75_ if p0 else _dafny.Set({self.f7}))

    def m6(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: bool = False
        d_2602_v0_: _dafny.Seq
        d_2602_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        r1 = (d_2602_v0_) != (d_2602_v0_)
        d_2603_v1_: int
        d_2603_v1_ = 556
        d_2604_v2_: D21
        d_2604_v2_ = D21_DC46(p0, p1, d_2603_v1_)
        r2 = (d_2604_v2_).cf61
        d_2605_v3_: _dafny.Array
        nw445_ = _dafny.Array(None, 2)
        nw445_[int(0)] = d_2603_v1_
        nw445_[int(1)] = d_2603_v1_
        d_2605_v3_ = nw445_
        index456_ = default__.safeIndex(402, (d_2605_v3_).length(0))
        (d_2605_v3_)[index456_] = d_2603_v1_
        index457_ = default__.safeIndex(402, (d_2605_v3_).length(0))
        (d_2605_v3_)[index457_] = default__.safeModuloInt((d_2603_v1_) + (d_2603_v1_), 33)
        r2 = (self).f6
        d_2606_v4_: _dafny.Map
        d_2606_v4_ = _dafny.Map({p1: p0})
        d_2607_v5_: D15
        d_2607_v5_ = D15_DC28(((d_2606_v4_)[default__.fm0(p1, globalState)] if (default__.fm0(p1, globalState)) in (d_2606_v4_) else (self).f6), (self).f6)
        d_2608_i0_: int
        d_2608_i0_ = 0
        with _dafny.label("24"):
            while (d_2607_v5_).cf37:
                with _dafny.c_label("24"):
                    if (d_2608_i0_) >= (100):
                        raise _dafny.Break("24")
                    d_2608_i0_ = (d_2608_i0_) + (1)
                    d_2609_v6_: D0
                    d_2609_v6_ = D0_DC0(p1)
                    d_2610_v7_: C3
                    nw446_ = C3()
                    nw446_.ctor__(p1, (self).fm9((0) - (default__.fm3(p1, p0, globalState)), d_2609_v6_, default__.fm0(p0, globalState), d_2603_v1_, globalState))
                    d_2610_v7_ = nw446_
                    d_2611_v8_: _dafny.Seq
                    d_2611_v8_ = _dafny.SeqWithoutIsStrInference([d_2610_v7_])
                    d_2612_v9_: _dafny.Seq
                    d_2612_v9_ = _dafny.SeqWithoutIsStrInference([(d_2610_v7_ if not(p0) else d_2610_v7_), (d_2611_v8_)[default__.safeIndex((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], len(d_2611_v8_))], d_2610_v7_])
                    d_2613_v10_: _dafny.Seq
                    d_2613_v10_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]), (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], d_2603_v1_, (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]])
                    d_2610_v7_ = (d_2612_v9_)[default__.safeIndex(len(d_2613_v10_), len(d_2612_v9_))]
                    d_2614_v11_: _dafny.MultiSet
                    d_2614_v11_ = _dafny.MultiSet([d_2603_v1_])
                    d_2615_v12_: _dafny.Set
                    d_2615_v12_ = _dafny.Set({(self).fm4(d_2602_v0_, 914, d_2603_v1_, d_2614_v11_, globalState)})
                    d_2613_v10_ = _dafny.SeqWithoutIsStrInference([d_2603_v1_, d_2603_v1_, len((d_2615_v12_) | (d_2615_v12_)), (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]])
                    index458_ = default__.safeIndex(402, (d_2605_v3_).length(0))
                    (d_2605_v3_)[index458_] = default__.fm3(not((d_2610_v7_).f14), p0, globalState)
                    r2 = (d_2615_v12_).isdisjoint(d_2615_v12_)
                    pass
            pass
        d_2616_v13_: D6
        d_2616_v13_ = D6_DC14((self).f6)
        d_2617_i1_: int
        d_2617_i1_ = 0
        with _dafny.label("25"):
            while ((D6_DC14(p1) if (self).f6 else d_2616_v13_)).cf15:
                with _dafny.c_label("25"):
                    if (d_2617_i1_) >= (100):
                        raise _dafny.Break("25")
                    d_2617_i1_ = (d_2617_i1_) + (1)
                    d_2618_v14_: D0
                    d_2618_v14_ = D0_DC1()
                    d_2619_v15_: _dafny.Map
                    d_2619_v15_ = _dafny.Map({d_2618_v14_: 25})
                    d_2619_v15_ = (d_2619_v15_) | (d_2619_v15_)
                    arr27_ = self.f7
                    index459_ = default__.safeIndex(222, (self.f7).length(0))
                    arr27_[index459_] = p1
                    arr28_ = self.f7
                    index460_ = default__.safeIndex(222, (self.f7).length(0))
                    arr28_[index460_] = ((self).f6 if not((self).f6) else p1)
                    d_2620_v16_: _dafny.Array
                    nw447_ = _dafny.Array(_dafny.Map({}), 22)
                    d_2620_v16_ = nw447_
                    d_2621_v17_: _dafny.Map
                    d_2621_v17_ = _dafny.Map({d_2620_v16_: not((self).f6)})
                    d_2621_v17_ = (d_2621_v17_).set(d_2620_v16_, p0)
                    d_2622_v18_: _dafny.MultiSet
                    d_2622_v18_ = _dafny.MultiSet([True, p0])
                    d_2623_v19_: D15
                    d_2623_v19_ = D15_DC27((d_2622_v18_) | (_dafny.MultiSet(d_2602_v0_)))
                    source36_ = d_2623_v19_
                    if source36_.is_DC28:
                        d_2624___mcc_h0_ = source36_.cf36
                        d_2625___mcc_h1_ = source36_.cf37
                        d_2626_cf37_ = d_2625___mcc_h1_
                        d_2627_cf36_ = d_2624___mcc_h0_
                        d_2628_v20_: _dafny.Map
                        d_2628_v20_ = _dafny.Map({d_2627_cf36_: _dafny.MultiSet([d_2603_v1_, d_2603_v1_, (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]])})
                        d_2629_v21_: _dafny.MultiSet
                        d_2629_v21_ = _dafny.MultiSet([len(d_2628_v20_)])
                        d_2630_v22_: _dafny.MultiSet
                        d_2630_v22_ = _dafny.MultiSet([(d_2629_v21_).cardinality])
                        d_2631_v24_: _dafny.Map
                        def iife177_():
                            coll77_ = _dafny.Map()
                            compr_77_: int
                            for compr_77_ in _dafny.IntegerRange(138, 86):
                                d_2632_v23_: int = compr_77_
                                if ((138) <= (d_2632_v23_)) and ((d_2632_v23_) < (86)):
                                    coll77_[(d_2632_v23_) - (-264)] = False
                            return _dafny.Map(coll77_)
                        d_2631_v24_ = iife177_()
                        
                        d_2633_v25_: _dafny.Map
                        d_2633_v25_ = _dafny.Map({(d_2603_v1_) in (d_2629_v21_): d_2631_v24_})
                        d_2633_v25_ = (d_2633_v25_).set(not(True), (d_2631_v24_ if p1 else d_2631_v24_))
                        d_2634_v26_: _dafny.Set
                        d_2634_v26_ = _dafny.Set({d_2626_cf37_})
                        d_2635_v27_: _dafny.Seq
                        d_2635_v27_ = _dafny.SeqWithoutIsStrInference([d_2634_v26_])
                        d_2636_v28_: _dafny.Map
                        d_2636_v28_ = _dafny.Map({d_2626_cf37_: len(d_2602_v0_)})
                        d_2637_v29_: _dafny.Array
                        nw448_ = _dafny.Array(None, 8)
                        nw448_[int(0)] = _dafny.Set({True})
                        nw448_[int(1)] = (_dafny.Set({p1})) - (d_2634_v26_)
                        nw448_[int(2)] = d_2634_v26_
                        nw448_[int(3)] = (d_2635_v27_)[default__.safeIndex(len(d_2636_v28_), len(d_2635_v27_))]
                        nw448_[int(4)] = default__.fm17(d_2603_v1_, p1, globalState)
                        nw448_[int(5)] = default__.fm17(d_2603_v1_, (self).f6, globalState)
                        nw448_[int(6)] = (d_2634_v26_) | (d_2634_v26_)
                        nw448_[int(7)] = default__.fm49(globalState)
                        d_2637_v29_ = nw448_
                        d_2638_v30_: _dafny.Array
                        nw449_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                        d_2638_v30_ = nw449_
                        index461_ = default__.safeIndex(403, (d_2638_v30_).length(0))
                        (d_2638_v30_)[index461_] = _dafny.SeqWithoutIsStrInference([(self).f5 for d_2639_i2_ in range(default__.abs(-409))])
                        d_2640_v31_: _dafny.Seq
                        d_2640_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ayjjit"))
                        index462_ = default__.safeIndex(403, (d_2638_v30_).length(0))
                        rhs519_ = d_2637_v29_
                        rhs520_ = d_2640_v31_
                        lhs342_ = d_2638_v30_
                        lhs343_ = default__.safeIndex(403, (d_2638_v30_).length(0))
                        d_2637_v29_ = rhs519_
                        lhs342_[lhs343_] = rhs520_
                        (globalState).f1 = ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))] if not((p0) == (d_2626_cf37_)) else default__.safeDivisionInt((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], len(default__.fm1(597, d_2627_cf36_, (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], globalState))))
                        index463_ = default__.safeIndex(403, (d_2638_v30_).length(0))
                        (d_2638_v30_)[index463_] = d_2640_v31_
                    elif source36_.is_DC27:
                        d_2641___mcc_h2_ = source36_.cf35
                        d_2642_cf35_ = d_2641___mcc_h2_
                        r1 = False
                        (globalState).f1 = (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]
                        d_2643_v32_: _dafny.Seq
                        d_2643_v32_ = _dafny.SeqWithoutIsStrInference([(d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]])
                        d_2644_v33_: _dafny.Array
                        nw450_ = _dafny.Array(None, 13)
                        nw450_[int(0)] = not((p1 if (self).f6 else p0))
                        nw450_[int(1)] = p1
                        nw450_[int(2)] = (d_2602_v0_)[default__.safeIndex(d_2603_v1_, len(d_2602_v0_))]
                        nw450_[int(3)] = p1
                        nw450_[int(4)] = p1
                        nw450_[int(5)] = (self.f7)[default__.safeIndex(222, (self.f7).length(0))]
                        nw450_[int(6)] = (d_2643_v32_) == (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_2603_v1_: (self).f5})) for d_2645_i3_ in range(default__.abs(368))]))
                        nw450_[int(7)] = (self).f6
                        nw450_[int(8)] = not ((self.f7)[default__.safeIndex(222, (self.f7).length(0))]) or (p0)
                        nw450_[int(9)] = (self.f7)[default__.safeIndex(222, (self.f7).length(0))]
                        nw450_[int(10)] = (self.f7)[default__.safeIndex(222, (self.f7).length(0))]
                        nw450_[int(11)] = p0
                        nw450_[int(12)] = (d_2603_v1_) not in (d_2643_v32_)
                        d_2644_v33_ = nw450_
                        d_2646_v34_: _dafny.Map
                        d_2646_v34_ = _dafny.Map({(self).f6: ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]) + (d_2603_v1_)})
                        index464_ = default__.safeIndex(402, (d_2605_v3_).length(0))
                        arr29_ = self.f7
                        index465_ = default__.safeIndex(222, (self.f7).length(0))
                        rhs521_ = d_2603_v1_
                        rhs522_ = d_2644_v33_
                        rhs523_ = d_2644_v33_
                        rhs524_ = d_2646_v34_
                        rhs525_ = (self).f6
                        lhs344_ = d_2605_v3_
                        lhs345_ = default__.safeIndex(402, (d_2605_v3_).length(0))
                        lhs346_ = self.f7
                        lhs347_ = default__.safeIndex(222, (self.f7).length(0))
                        lhs344_[lhs345_] = rhs521_
                        d_2644_v33_ = rhs522_
                        d_2644_v33_ = rhs523_
                        d_2646_v34_ = rhs524_
                        lhs346_[lhs347_] = rhs525_
                        d_2647_v35_: _dafny.Set
                        d_2647_v35_ = _dafny.Set({(self).f6})
                        d_2648_v36_: _dafny.Map
                        d_2648_v36_ = _dafny.Map({d_2647_v35_: (self).f6})
                        arr30_ = self.f7
                        index466_ = default__.safeIndex(222, (self.f7).length(0))
                        arr30_[index466_] = not((self).fm10(d_2622_v18_, d_2648_v36_, ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]) - (d_2603_v1_), globalState))
                    elif True:
                        d_2649___mcc_h3_ = source36_.cf38
                        d_2650_cf38_ = d_2649___mcc_h3_
                        d_2651_v37_: _dafny.Seq
                        d_2651_v37_ = _dafny.SeqWithoutIsStrInference([d_2603_v1_])
                        d_2652_v38_: _dafny.Map
                        d_2652_v38_ = _dafny.Map({709: (d_2651_v37_).set(default__.safeIndex((self).fm8(True, (self).f6, (self.f7)[default__.safeIndex(222, (self.f7).length(0))], globalState), len(d_2651_v37_)), (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))])})
                        d_2653_v39_: _dafny.Seq
                        d_2653_v39_ = _dafny.SeqWithoutIsStrInference([((d_2652_v38_)[(d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]] if ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]) in (d_2652_v38_) else d_2651_v37_), d_2651_v37_])
                        d_2654_v40_: _dafny.Map
                        d_2654_v40_ = _dafny.Map({len(d_2653_v39_): True})
                        d_2654_v40_ = (d_2654_v40_).set((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))], not((self.f7)[default__.safeIndex(222, (self.f7).length(0))]))
                        d_2655_v41_: _dafny.Set
                        d_2655_v41_ = _dafny.Set({d_2603_v1_})
                        d_2656_v42_: _dafny.Seq
                        d_2656_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxqw"))
                        d_2657_v43_: _dafny.Map
                        d_2657_v43_ = _dafny.Map({len(d_2655_v41_): (self).fm8((d_2602_v0_)[default__.safeIndex(len(d_2656_v42_), len(d_2602_v0_))], p1, (self.f7)[default__.safeIndex(222, (self.f7).length(0))], globalState)})
                        d_2657_v43_ = (d_2657_v43_).set(d_2603_v1_, (self).fm8((self).f6, False, True, globalState))
                        arr31_ = self.f7
                        index467_ = default__.safeIndex(222, (self.f7).length(0))
                        arr31_[index467_] = (self).f6
                        r1 = False
                    pass
            pass
        d_2658_v44_: _dafny.Seq
        d_2658_v44_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]), d_2603_v1_])
        r0 = (((d_2658_v44_) + (d_2658_v44_)) + (_dafny.SeqWithoutIsStrInference([(0) - (d_2603_v1_), (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]]))).set(default__.safeIndex((d_2603_v1_) * (d_2603_v1_), len(((d_2658_v44_) + (d_2658_v44_)) + (_dafny.SeqWithoutIsStrInference([(0) - (d_2603_v1_), (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))]])))), d_2603_v1_)
        d_2659_v45_: D11
        d_2659_v45_ = D11_DC20(p0, (d_2605_v3_)[default__.safeIndex(402, (d_2605_v3_).length(0))])
        r1 = (d_2659_v45_).cf21
        r2 = not(p1)
        return r0, r1, r2

    @property
    def f8(self):
        return self._f8
