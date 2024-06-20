// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gmfo"), ((!(true)) ? (_dafny.Seq.UnicodeFromString("kambpnln")) : (_dafny.Seq.UnicodeFromString("crnml"))))).length);
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(682),new BigNumber(629))).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber(497))).Elements) {
          let _0_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(497)), _0_v0)) {
            _coll0.push([(_0_v0).minus(new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber((_dafny.Seq.UnicodeFromString("kee")).length)]);
          }
        }
        return _coll0;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length),new BigNumber((_dafny.Seq.UnicodeFromString("oehpvis")).length)));
    };
    static fm2(p0, globalState) {
      return _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0)));
    };
    static fm3(p0, globalState) {
      if (!(!(new BigNumber((_dafny.Seq.UnicodeFromString("po")).length)).isEqualTo(new BigNumber(384)))) {
        return (_dafny.MultiSet.fromElements(!(false), (_module.D4.create_DC13(!(true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-529)), function (_1_i0) {
  return new BigNumber(41);
}))).dtor_cf21)).IsProperSubsetOf(_dafny.MultiSet.fromElements(true, true));
      } else {
        return !(new BigNumber(-543)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length));
      }
    };
    static fm4(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_2_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("bgvltt")).length);
      })).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(320))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(549))));
    };
    static fm5(p0, globalState) {
      return (_dafny.MultiSet.fromElements(!(false), true)).Intersect(_dafny.MultiSet.FromArray(((!(false)) ? (_dafny.Seq.of(false, false, true, !((_module.D13.create_DC34(false, true, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(923)), function (_3_i0) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length))).length))).cardinality()), _dafny.Seq.UnicodeFromString("g"))).dtor_cf54), false)) : (_dafny.Seq.of(false, !(false), false)))));
    };
    static fm6(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(false, true)).Difference(_dafny.Set.fromElements(!((_module.D25.create_DC63(new BigNumber(551), !(true), true, false)).dtor_cf99)))).Difference(((!(!(true))) ? (_dafny.Set.fromElements(true, false)) : (_dafny.Set.fromElements(true))));
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Seq.of(true);
    };
    static fm8(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(95), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(83)), function (_4_i0) {
        return _module.D9.create_DC26(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-684), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(624)), function (_5_i1) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("tx")).length);
})).length))).cardinality()));
      })).length), new BigNumber((_dafny.Seq.of(true)).length))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("csvnyvpqu")).length))))).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(897), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(_module.D9.create_DC27(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(747),false))).length)))).Elements) {
          let _6_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D9.create_DC27(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(747),false))).length))), _6_v0)) {
            _coll1.add(_6_v0);
          }
        }
        return _coll1;
      }()).length))).length))));
    };
    static fm9(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(false)));
    };
    static fm10(p0, p1, p2, globalState) {
      if (!(false)) {
        return _module.D2.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), true, true, new BigNumber(502));
      } else {
        return _module.D2.create_DC6(new _dafny.CodePoint('y'.codePointAt(0)), false, true, new BigNumber(787));
      }
    };
    static fm17(p0, globalState) {
      return _module.D3.create_DC11(_module.D3.create_DC9(new BigNumber(168), false));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(389), new BigNumber((_dafny.Seq.UnicodeFromString("gkid")).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(609), new BigNumber(108), new BigNumber(299), new BigNumber(356), new BigNumber(454))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(286))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(867))).length), new BigNumber(361), new BigNumber(155)));
    };
    static fm22(p0, p1, globalState) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC11(((true) ? (_module.D3.create_DC10(new BigNumber(868), new BigNumber((_dafny.Seq.of(true, false)).length))) : (_module.D3.create_DC9(new BigNumber(100), true))));
    };
    static fm24(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm31(p0, p1, p2, p3, globalState) {
      if ((new BigNumber(-383)).isEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qldwqoe")).length), new BigNumber(769))).length))) {
        if (!(true)) {
          return _dafny.Seq.of(new BigNumber(-827));
        } else {
          return _dafny.Seq.of(new BigNumber(-628));
        }
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-762)), function (_7_i0) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tcvlq")).length),new BigNumber(320)),false)).length)))).cardinality());
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(99)), function (_8_i1) {
          return (_dafny.ZERO).minus(new BigNumber(-517));
        }));
      }
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC9(new BigNumber(-426), false),new BigNumber(-56));
    };
    static fm33(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(590), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ncy")).length)), new BigNumber((_dafny.Seq.UnicodeFromString("md")).length))).Union(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(683), new BigNumber(668))) {
          let _9_v0 = _compr_2;
          if (((new BigNumber(683)).isLessThanOrEqualTo(_9_v0)) && ((_9_v0).isLessThan(new BigNumber(668)))) {
            _coll2.add((_9_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), function (_10_i0) {
              return new _dafny.CodePoint('u'.codePointAt(0));
            })).length)));
          }
        }
        return _coll2;
      }());
    };
    static fm34(p0, p1, globalState) {
      return (_module.D26.create_DC65(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, true)))).dtor_cf102;
    };
    static fm35(p0, globalState) {
      return _module.D0.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("vhndkcyca")).length));
    };
    static fm36(p0, globalState) {
      return _module.D6.create_DC18((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-147))).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-719)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true)).length)))).length), true, new _dafny.CodePoint('t'.codePointAt(0)), new BigNumber(428));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_module.D9.create_DC26(new BigNumber(38))).dtor_cf44,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-408)), function (_11_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(200)), function (_12_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), function (_13_i2) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(591),_dafny.Seq.UnicodeFromString("apixoryg"))));
    };
    static fm38(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('j'.codePointAt(0));
    };
    static fm39(globalState) {
      return _module.D3.create_DC9(_module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()), new BigNumber((_dafny.Seq.of(new BigNumber(477), new BigNumber((_dafny.Seq.of(new BigNumber(543), new BigNumber((_dafny.Seq.UnicodeFromString("ymxyj")).length))).length))).length)), !(false));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC11(((false) ? (_module.D3.create_DC10(new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-591)), function (_14_i0) {
    return new _dafny.CodePoint('w'.codePointAt(0));
  })).Elements) {
    let _15_v0 = _compr_3;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-591)), function (_14_i0) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    }), _15_v0)) {
      _coll3.push([_15_v0,new BigNumber((_dafny.Seq.of(true)).length)]);
    }
  }
  return _coll3;
}()).length), new BigNumber(294))) : (_module.D3.create_DC10(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(608), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(410)), function (_16_i1) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})).length))).length))).length), new BigNumber(952)))));
    };
    static fm41(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.MultiSet.fromElements(_module.D6.create_DC20(_module.D6.create_DC18(new BigNumber(707), new BigNumber((_dafny.Seq.UnicodeFromString("njoljke")).length), false, new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("fcpaxxcwa")).length))))).Elements) {
          let _17_v0 = _compr_4;
          if ((_dafny.MultiSet.fromElements(_module.D6.create_DC20(_module.D6.create_DC18(new BigNumber(707), new BigNumber((_dafny.Seq.UnicodeFromString("njoljke")).length), false, new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("fcpaxxcwa")).length))))).contains(_17_v0)) {
            _coll4.push([_17_v0,new BigNumber(715)]);
          }
        }
        return _coll4;
      }())).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC20(_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(929),true))),new BigNumber((_dafny.Seq.UnicodeFromString("lfk")).length)), _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC20(_module.D6.create_DC20(_module.D6.create_DC20(_module.D6.create_DC20(_module.D6.create_DC18(new BigNumber(899), new BigNumber(505), true, new _dafny.CodePoint('e'.codePointAt(0)), new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-690)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(851)), function (_18_i0) {
  return new _dafny.CodePoint('u'.codePointAt(0));
})).length))).cardinality())))))),new BigNumber(121))));
    };
    static fm42(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_module.__default.safeModuloInt(new BigNumber(32), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(346),true)).length)));
    };
    static fm43(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true));
    };
    static fm44(globalState) {
      return ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, !(true), true, false, true))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(true)))).Intersect(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true), false), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true))).Elements) {
          let _19_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true), false), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true)), _19_v0)) {
            _coll5.add(_19_v0);
          }
        }
        return _coll5;
      }());
    };
    static fm45(p0, p1, globalState) {
      let _source0 = _module.D18.create_DC44(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC34(!(true), true, new BigNumber(606), _dafny.Seq.UnicodeFromString("nagwkirpq")),false));
      if (_source0.is_DC45) {
        let _20___mcc_h0 = (_source0).cf75;
        let _21___mcc_h1 = (_source0).cf76;
        let _22___mcc_h2 = (_source0).cf77;
        let _23_cf77 = _22___mcc_h2;
        let _24_cf76 = _21___mcc_h1;
        let _25_cf75 = _20___mcc_h0;
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), function (_26_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_27_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })));
      } else if (_source0.is_DC46) {
        return _dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0))));
      } else if (_source0.is_DC44) {
        let _28___mcc_h3 = (_source0).cf74;
        let _29_cf74 = _28___mcc_h3;
        return _dafny.MultiSet.fromElements((_module.D14.create_DC38(_dafny.Seq.UnicodeFromString("cqrhbqnd"))).dtor_cf64, _dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0))));
      } else {
        let _30___mcc_h4 = (_source0).cf78;
        let _31_cf78 = _30___mcc_h4;
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0)))));
      }
    };
    static fm46(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(685)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), function (_32_i0) {
        return new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).cardinality());
      }))), (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bs")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC11(_module.D3.create_DC10(new BigNumber(323), new BigNumber(384))),true)).length)))).Difference(_dafny.MultiSet.fromElements((_module.D19.create_DC49(new BigNumber(325), new BigNumber(780))).dtor_cf81, new BigNumber((_dafny.Seq.UnicodeFromString("rkkho")).length))), _dafny.MultiSet.fromElements(new BigNumber(219), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, true, true, false), _dafny.Seq.of(true), _dafny.Seq.of(true), _dafny.Seq.of(false, !(true), true, true), _dafny.Seq.of(!(!(true))))).length))).length)));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_33_i0) {
        return _module.D8.create_DC24(new _dafny.CodePoint('f'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(true,false));
      });
    };
    static fm48(p0, p1, globalState) {
      return _dafny.Seq.of(_module.D6.create_DC18(new BigNumber(693), new BigNumber((_dafny.Seq.UnicodeFromString("gmkabar")).length), false, new _dafny.CodePoint('d'.codePointAt(0)), new BigNumber((_dafny.Seq.of(true)).length)), _module.D6.create_DC18(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mt")).length))).length), new BigNumber(972), !(true), new _dafny.CodePoint('e'.codePointAt(0)), new BigNumber(894)));
    };
    static fm49(p0, p1, p2, globalState) {
      if (false) {
        return _module.D6.create_DC17(true);
      } else {
        return _module.D6.create_DC17(true);
      }
    };
    static fm50(p0, p1, p2, globalState) {
      return _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(978))).length));
    };
    static fm51(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-656)), function (_34_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }),new BigNumber(805))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cxllmt"),new BigNumber(-687)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ovmiu"),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))));
    };
    static fm52(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.MultiSet.fromElements(true, false, !(false), false, false), (_dafny.MultiSet.fromElements(!(false))).Intersect(_dafny.MultiSet.fromElements(true)));
    };
    static fm53(p0, globalState) {
      let _source1 = _module.D4.create_DC13(!(!(true)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(127), new BigNumber(985))).cardinality()))).length)));
      if (_source1.is_DC13) {
        let _35___mcc_h0 = (_source1).cf21;
        let _36___mcc_h1 = (_source1).cf22;
        let _37_cf22 = _36___mcc_h1;
        let _38_cf21 = _35___mcc_h0;
        return (function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.MultiSet.fromElements(_module.D13.create_DC35(new BigNumber(195), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_38_cf21,_38_cf21)).length),new BigNumber((_37_cf22).length))).length))).cardinality()), _38_cf21, new BigNumber(-88), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
    let _39_v1 = _compr_7;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0))), _39_v1)) {
      _coll7.add(_39_v1);
    }
  }
  return _coll7;
}()).length),_dafny.Seq.UnicodeFromString("cgeiumb")))), _module.D13.create_DC35(new BigNumber(-979), new BigNumber(-35), _38_cf21, new BigNumber(-701), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-944),_dafny.Seq.UnicodeFromString("ohcyaelk")))), _module.D13.create_DC35(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_38_cf21))).cardinality()), new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_38_cf21)).length),new BigNumber(880)))).length), _38_cf21, new BigNumber((_dafny.Set.fromElements(_38_cf21)).length), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, !(_38_cf21))).cardinality()))).length),_dafny.Seq.UnicodeFromString("ddei")))))).Elements) {
            let _40_v0 = _compr_6;
            if ((_dafny.MultiSet.fromElements(_module.D13.create_DC35(new BigNumber(195), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_38_cf21,_38_cf21)).length),new BigNumber((_37_cf22).length))).length))).cardinality()), _38_cf21, new BigNumber(-88), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll8 = new _dafny.Set();
  for (const _compr_8 of (_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)))).Elements) {
    let _41_v1 = _compr_8;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0))), _41_v1)) {
      _coll8.add(_41_v1);
    }
  }
  return _coll8;
}()).length),_dafny.Seq.UnicodeFromString("cgeiumb")))), _module.D13.create_DC35(new BigNumber(-979), new BigNumber(-35), _38_cf21, new BigNumber(-701), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-944),_dafny.Seq.UnicodeFromString("ohcyaelk")))), _module.D13.create_DC35(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_38_cf21))).cardinality()), new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_38_cf21)).length),new BigNumber(880)))).length), _38_cf21, new BigNumber((_dafny.Set.fromElements(_38_cf21)).length), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, !(_38_cf21))).cardinality()))).length),_dafny.Seq.UnicodeFromString("ddei")))))).contains(_40_v0)) {
              _coll6.push([_40_v0,new _dafny.CodePoint('h'.codePointAt(0))]);
            }
          }
          return _coll6;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC35(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber(-211))).length), new BigNumber((_dafny.MultiSet.FromArray(_37_cf22)).cardinality()), _38_cf21, new BigNumber(833), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(596),_dafny.Seq.UnicodeFromString("qlorlrnk")))),new _dafny.CodePoint('t'.codePointAt(0))));
      } else {
        let _42___mcc_h2 = (_source1).cf20;
        let _43_cf20 = _42___mcc_h2;
        return (function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of (_dafny.Seq.of(_module.D13.create_DC35(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.of(new BigNumber(876), new BigNumber(41))).length), false, (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll10 = new _dafny.Map();
  for (const _compr_10 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), function (_44_i0) {
    return new BigNumber((_dafny.Seq.UnicodeFromString("vwdhkf")).length);
  })).Elements) {
    let _45_v3 = _compr_10;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), function (_44_i0) {
      return new BigNumber((_dafny.Seq.UnicodeFromString("vwdhkf")).length);
    }), _45_v3)) {
      _coll10.push([_module.__default.safeDivisionInt(_45_v3, new BigNumber(-591)),true]);
    }
  }
  return _coll10;
}()).length)), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(654),_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), function (_46_i1) {
  return (_module.D8.create_DC24(new _dafny.CodePoint('m'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(!(false),true))).dtor_cf41;
})))))).Elements) {
            let _47_v2 = _compr_9;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D13.create_DC35(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.of(new BigNumber(876), new BigNumber(41))).length), false, (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll11 = new _dafny.Map();
  for (const _compr_11 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), function (_44_i0) {
    return new BigNumber((_dafny.Seq.UnicodeFromString("vwdhkf")).length);
  })).Elements) {
    let _45_v3 = _compr_11;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), function (_44_i0) {
      return new BigNumber((_dafny.Seq.UnicodeFromString("vwdhkf")).length);
    }), _45_v3)) {
      _coll11.push([_module.__default.safeDivisionInt(_45_v3, new BigNumber(-591)),true]);
    }
  }
  return _coll11;
}()).length)), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(654),_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), function (_46_i1) {
  return (_module.D8.create_DC24(new _dafny.CodePoint('m'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(!(false),true))).dtor_cf41;
}))))), _47_v2)) {
              _coll9.push([_47_v2,new _dafny.CodePoint('e'.codePointAt(0))]);
            }
          }
          return _coll9;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC35(new BigNumber(694), new BigNumber(-984), true, new BigNumber(87), _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(559),_dafny.Seq.UnicodeFromString("d")))),new _dafny.CodePoint('t'.codePointAt(0))));
      }
    };
    static fm54(p0, p1, globalState) {
      let _source2 = _module.D3.create_DC10(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),false)).length));
      if (_source2.is_DC9) {
        let _48___mcc_h0 = (_source2).cf15;
        let _49___mcc_h1 = (_source2).cf16;
        let _50_cf16 = _49___mcc_h1;
        let _51_cf15 = _48___mcc_h0;
        return _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-500),_50_cf16));
      } else if (_source2.is_DC10) {
        let _52___mcc_h2 = (_source2).cf17;
        let _53___mcc_h3 = (_source2).cf18;
        let _54_cf18 = _53___mcc_h3;
        let _55_cf17 = _52___mcc_h2;
        if (!(false)) {
          return _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),false));
        } else {
          return _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(469),true));
        }
      } else if (_source2.is_DC8) {
        let _56___mcc_h4 = (_source2).cf14;
        let _57_cf14 = _56___mcc_h4;
        return _module.D6.create_DC16(function () {
  let _coll12 = new _dafny.Map();
  for (const _compr_12 of _dafny.IntegerRange(new BigNumber(961), new BigNumber(472))) {
    let _58_v0 = _compr_12;
    if (((new BigNumber(961)).isLessThanOrEqualTo(_58_v0)) && ((_58_v0).isLessThan(new BigNumber(472)))) {
      _coll12.push([(_58_v0).minus(new BigNumber(48)),true]);
    }
  }
  return _coll12;
}());
      } else {
        let _59___mcc_h5 = (_source2).cf19;
        let _60_cf19 = _59___mcc_h5;
        return _module.D6.create_DC16(function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of (_dafny.MultiSet.fromElements(new BigNumber(637))).Elements) {
    let _61_v1 = _compr_13;
    if ((_dafny.MultiSet.fromElements(new BigNumber(637))).contains(_61_v1)) {
      _coll13.push([_module.__default.safeModuloInt(_61_v1, new BigNumber(-791)),(_module.D13.create_DC34(true, false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(!(false)))))).cardinality()), _dafny.Seq.UnicodeFromString("nnysmsacj"))).dtor_cf53]);
    }
  }
  return _coll13;
}());
      }
    };
    static fm55(p0, p1, p2, globalState) {
      return _module.D2.create_DC5(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), function (_62_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}))).length)));
    };
    static fm56(p0, globalState) {
      return _module.D11.create_DC31(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(279), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(179),new BigNumber(887))).length))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(860),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(702), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(902))).length))).cardinality()))).length)))).length)))).cardinality()));
    };
    static fm57(p0, globalState) {
      return ((function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC23(new BigNumber((function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-74), new BigNumber(163))) {
    let _63_v0 = _compr_15;
    if (((new BigNumber(-74)).isLessThanOrEqualTo(_63_v0)) && ((_63_v0).isLessThan(new BigNumber(163)))) {
      _coll15.push([_module.__default.safeModuloInt(_63_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-969)), function (_64_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length)),new BigNumber(474)]);
    }
  }
  return _coll15;
}()).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(438))).length))).cardinality()), false),true)).length), new BigNumber(416))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length))).length))).length),false),!(true))).Keys.Elements) {
          let _65_v1 = _compr_14;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC23(new BigNumber((function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-74), new BigNumber(163))) {
    let _63_v0 = _compr_16;
    if (((new BigNumber(-74)).isLessThanOrEqualTo(_63_v0)) && ((_63_v0).isLessThan(new BigNumber(163)))) {
      _coll16.push([_module.__default.safeModuloInt(_63_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-969)), function (_64_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length)),new BigNumber(474)]);
    }
  }
  return _coll16;
}()).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(438))).length))).cardinality()), false),true)).length), new BigNumber(416))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length))).length))).length),false),!(true))).contains(_65_v1)) {
            _coll14.add(_65_v1);
          }
        }
        return _coll14;
      }()).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Seq.of(true))).length),!(true)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(-841), new BigNumber(918))) {
          let _66_v2 = _compr_17;
          if (((new BigNumber(-841)).isLessThanOrEqualTo(_66_v2)) && ((_66_v2).isLessThan(new BigNumber(918)))) {
            _coll17.push([(_66_v2).plus(new BigNumber(7)),new _dafny.CodePoint('p'.codePointAt(0))]);
          }
        }
        return _coll17;
      }()).length),_dafny.MultiSet.fromElements(true))).length))).length),true)))).Intersect(_dafny.Set.fromElements(function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(405), new BigNumber(398))) {
          let _67_v3 = _compr_18;
          if (((new BigNumber(405)).isLessThanOrEqualTo(_67_v3)) && ((_67_v3).isLessThan(new BigNumber(398)))) {
            _coll18.push([_module.__default.safeDivisionInt(_67_v3, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))),true]);
          }
        }
        return _coll18;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(621),true)));
    };
    static fm58(p0, p1, p2, globalState) {
      return (_module.D16.create_DC41(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-156)), function (_68_i0) {
  return _dafny.Seq.of(false);
}))).dtor_cf69;
    };
    static fm59(p0, globalState) {
      return _module.D8.create_DC23((new BigNumber((_dafny.Seq.UnicodeFromString("wjlhd")).length)).minus(new BigNumber((_dafny.Seq.of(!(!(true)))).length)), new BigNumber(360), !(!((_dafny.MultiSet.fromElements(false, !(!(false)))).IsSubsetOf(_dafny.MultiSet.fromElements(false)))));
    };
    static fm60(globalState) {
      if (!((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).IsDisjointFrom(_dafny.MultiSet.fromElements(false)))) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)));
      } else {
        return _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)));
      }
    };
    static fm61(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("shgscmrg")).length), new BigNumber(934)))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true)).length), new BigNumber(816)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(274)), function (_69_i0) {
        return new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, false, true, true)).length))).length);
      }), _dafny.Seq.of(new BigNumber(650))))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(970), new BigNumber(482)), _dafny.Seq.of(new BigNumber(550))));
    };
    static fm62(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("dgdwgnu"),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mobipcov"),false));
    };
    static fm63(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Difference(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true)))).Difference(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(!(true),!((_module.D3.create_DC9(new BigNumber((function () {
  let _coll19 = new _dafny.Map();
  for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(996)), function (_70_i0) {
    return new _dafny.CodePoint('e'.codePointAt(0));
  })).length),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).Keys.Elements) {
    let _71_v0 = _compr_19;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(996)), function (_70_i0) {
      return new _dafny.CodePoint('e'.codePointAt(0));
    })).length),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).contains(_71_v0)) {
      _coll19.push([_module.__default.safeModuloInt(_71_v0, new BigNumber(482)),true]);
    }
  }
  return _coll19;
}()).length), true)).dtor_cf16))));
    };
    static fm64(p0, p1, p2, globalState) {
      let _source3 = _module.D4.create_DC13(false, _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bpwp")).length), new BigNumber(-500), (_dafny.ZERO).minus(new BigNumber(((_module.D3.create_DC8(_dafny.Set.fromElements(false))).dtor_cf14).length))));
      if (_source3.is_DC13) {
        let _72___mcc_h0 = (_source3).cf21;
        let _73___mcc_h1 = (_source3).cf22;
        let _74_cf22 = _73___mcc_h1;
        let _75_cf21 = _72___mcc_h0;
        if (_75_cf21) {
          return function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_76_cf21) => function (_77_i0) {
                return _module.D13.create_DC34(_76_cf21, false, new BigNumber(725), _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_78_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
              })(_75_cf21))).Elements) {
                let _79_v1 = _compr_21;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_80_cf21) => function (_77_i0) {
                  return _module.D13.create_DC34(_80_cf21, false, new BigNumber(725), _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_78_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
                })(_75_cf21)), _79_v1)) {
                  _coll21.push([_79_v1,new BigNumber(-316)]);
                }
              }
              return _coll21;
            }()).Keys.Elements) {
              let _81_v0 = _compr_20;
              if ((function () {
                let _coll22 = new _dafny.Map();
                for (const _compr_22 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_82_cf21) => function (_77_i0) {
                  return _module.D13.create_DC34(_82_cf21, false, new BigNumber(725), _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_78_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
                })(_75_cf21))).Elements) {
                  let _79_v1 = _compr_22;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_83_cf21) => function (_77_i0) {
                    return _module.D13.create_DC34(_83_cf21, false, new BigNumber(725), _dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), function (_78_i1) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}));
                  })(_75_cf21)), _79_v1)) {
                    _coll22.push([_79_v1,new BigNumber(-316)]);
                  }
                }
                return _coll22;
              }()).contains(_81_v0)) {
                _coll20.push([_81_v0,_75_cf21]);
              }
            }
            return _coll20;
          }();
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC34(_75_cf21, _75_cf21, (_dafny.ZERO).minus(new BigNumber(-435)), _dafny.Seq.UnicodeFromString("ucbh")),_75_cf21);
        }
      } else {
        let _84___mcc_h2 = (_source3).cf20;
        let _85_cf20 = _84___mcc_h2;
        if (false) {
          return _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC34(false, !(false), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("f")).length),new BigNumber(221))).length), _dafny.Seq.UnicodeFromString("gxpbdad")),false);
        } else {
          return (_module.D18.create_DC44(_dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC34(false, !(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), _dafny.Seq.UnicodeFromString("lu")),false))).dtor_cf74;
        }
      }
    };
    static fm65(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC13(true, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xhn")).length), new BigNumber(-381), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(858)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-57)), function (_86_i0) {
  return new BigNumber(324);
})));
    };
    static fm66(p0, p1, p2, globalState) {
      let _source4 = _module.D8.create_DC23(new BigNumber(((_module.D8.create_DC24(new _dafny.CodePoint('n'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(false,true))).dtor_cf42).length), new BigNumber(462), true);
      if (_source4.is_DC23) {
        let _87___mcc_h0 = (_source4).cf38;
        let _88___mcc_h1 = (_source4).cf39;
        let _89___mcc_h2 = (_source4).cf40;
        let _90_cf40 = _89___mcc_h2;
        let _91_cf39 = _88___mcc_h1;
        let _92_cf38 = _87___mcc_h0;
        return _module.D8.create_DC24(new _dafny.CodePoint('m'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(true,_90_cf40));
      } else if (_source4.is_DC24) {
        let _93___mcc_h3 = (_source4).cf41;
        let _94___mcc_h4 = (_source4).cf42;
        let _95_cf42 = _94___mcc_h4;
        let _96_cf41 = _93___mcc_h3;
        return _module.D8.create_DC24(_96_cf41, _95_cf42);
      } else {
        let _97___mcc_h5 = (_source4).cf37;
        let _98_cf37 = _97___mcc_h5;
        return _module.D8.create_DC24(new _dafny.CodePoint('p'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(!(true),true));
      }
    };
    static fm67(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_module.D27.create_DC68(function () {
  let _coll23 = new _dafny.Map();
  for (const _compr_23 of _dafny.IntegerRange(new BigNumber(-31), new BigNumber(806))) {
    let _99_v0 = _compr_23;
    if (((new BigNumber(-31)).isLessThanOrEqualTo(_99_v0)) && ((_99_v0).isLessThan(new BigNumber(806)))) {
      _coll23.push([_module.__default.safeDivisionInt(_99_v0, new BigNumber(-805)),_module.D8.create_DC24(new _dafny.CodePoint('h'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe(false,false))]);
    }
  }
  return _coll23;
}())).dtor_cf109,((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(855),false)).length))).isLessThan(new BigNumber(530)));
    };
    static m0(p0, globalState) {
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.of();
      let _100_v0;
      let _nw0 = Array((new BigNumber(27)).toNumber()).fill(false);
      _100_v0 = _nw0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_100_v0).length))) {
        let _101_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_101_i0)) && ((_101_i0).isLessThan(new BigNumber((_100_v0).length))))) {
          (_100_v0)[(_101_i0)] = (_module.D1.create_DC4(true, true)).dtor_cf6;
        }
      }
      let _102_v1;
      let _nw1 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _102_v1 = _nw1;
      let _103_v2;
      _103_v2 = new BigNumber(-554);
      let _104_v3;
      _104_v3 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_103_v2));
      let _105_v4;
      _105_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_104_v3).cardinality()),_102_v1);
      let _106_v5;
      let _nw2 = Array((new BigNumber(4)).toNumber());
      _nw2[(_dafny.ZERO).toNumber()] = _102_v1;
      _nw2[(_dafny.ONE).toNumber()] = ((_module.__default.fm3(_103_v2, globalState)) ? (_102_v1) : (_102_v1));
      _nw2[(new BigNumber(2)).toNumber()] = _102_v1;
      _nw2[(new BigNumber(3)).toNumber()] = (((_105_v4).contains(_103_v2)) ? ((_105_v4).get(_103_v2)) : (_102_v1));
      _106_v5 = _nw2;
      _106_v5 = _106_v5;
      let _107_v6;
      let _nw3 = Array((new BigNumber(2)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = _100_v0;
      _nw3[(_dafny.ONE).toNumber()] = _100_v0;
      _107_v6 = _nw3;
      let _108_v7;
      _108_v7 = true;
      let _index0 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
      (_100_v0)[_index0] = _108_v7;
      let _109_v8;
      _109_v8 = new _dafny.CodePoint('r'.codePointAt(0));
      let _110_v9;
      _110_v9 = _module.D2.create_DC6(_109_v8, _108_v7, _108_v7, _103_v2);
      let _111_v10;
      _111_v10 = _dafny.Set.fromElements(_103_v2);
      let _112_v11;
      _112_v11 = _dafny.Seq.of(_108_v7, _108_v7);
      let _pat_let_tv0 = _108_v7;
      let _pat_let_tv1 = _109_v8;
      let _pat_let_tv2 = _108_v7;
      let _pat_let_tv3 = _109_v8;
      let _pat_let_tv4 = _108_v7;
      let _pat_let_tv5 = _103_v2;
      let _pat_let_tv6 = _109_v8;
      let _113_v12;
      let _nw4 = Array((new BigNumber(26)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = _110_v9;
      _nw4[(_dafny.ONE).toNumber()] = _110_v9;
      _nw4[(new BigNumber(2)).toNumber()] = function (_pat_let0_0) {
        return function (_114_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_115_dt__update_hcf10_h0) {
              return function (_pat_let2_0) {
                return function (_116_dt__update_hcf9_h0) {
                  return function (_pat_let3_0) {
                    return function (_117_dt__update_hcf11_h0) {
                      return _module.D2.create_DC6(_116_dt__update_hcf9_h0, _115_dt__update_hcf10_h0, _117_dt__update_hcf11_h0, (_114_dt__update__tmp_h0).dtor_cf12);
                    }(_pat_let3_0);
                  }(_pat_let_tv2);
                }(_pat_let2_0);
              }(_pat_let_tv1);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_module.D2.create_DC6(_109_v8, true, _108_v7, new BigNumber((_111_v10).length)));
      _nw4[(new BigNumber(3)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(4)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(5)).toNumber()] = _module.D2.create_DC6(_109_v8, true, false, new BigNumber(-709));
      _nw4[(new BigNumber(6)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(7)).toNumber()] = _module.D2.create_DC6(_109_v8, _108_v7, _108_v7, new BigNumber((_112_v11).length));
      _nw4[(new BigNumber(8)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(9)).toNumber()] = function (_pat_let4_0) {
        return function (_118_dt__update__tmp_h1) {
          return function (_pat_let5_0) {
            return function (_119_dt__update_hcf9_h1) {
              return function (_pat_let6_0) {
                return function (_120_dt__update_hcf11_h1) {
                  return _module.D2.create_DC6(_119_dt__update_hcf9_h1, (_118_dt__update__tmp_h1).dtor_cf10, _120_dt__update_hcf11_h1, (_118_dt__update__tmp_h1).dtor_cf12);
                }(_pat_let6_0);
              }(_pat_let_tv4);
            }(_pat_let5_0);
          }(_pat_let_tv3);
        }(_pat_let4_0);
      }(_110_v9);
      _nw4[(new BigNumber(10)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(11)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(12)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(13)).toNumber()] = ((false) ? (_110_v9) : (_110_v9));
      _nw4[(new BigNumber(14)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(15)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(16)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(17)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(18)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(19)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(20)).toNumber()] = _module.D2.create_DC6(_109_v8, _108_v7, _108_v7, _103_v2);
      _nw4[(new BigNumber(21)).toNumber()] = _module.D2.create_DC6(_109_v8, _108_v7, !(_108_v7), _103_v2);
      _nw4[(new BigNumber(22)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(23)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(24)).toNumber()] = _110_v9;
      _nw4[(new BigNumber(25)).toNumber()] = function (_pat_let7_0) {
        return function (_121_dt__update__tmp_h2) {
          return function (_pat_let8_0) {
            return function (_122_dt__update_hcf12_h0) {
              return function (_pat_let9_0) {
                return function (_123_dt__update_hcf9_h2) {
                  return _module.D2.create_DC6(_123_dt__update_hcf9_h2, (_121_dt__update__tmp_h2).dtor_cf10, (_121_dt__update__tmp_h2).dtor_cf11, _122_dt__update_hcf12_h0);
                }(_pat_let9_0);
              }(_pat_let_tv6);
            }(_pat_let8_0);
          }(_pat_let_tv5);
        }(_pat_let7_0);
      }(_module.__default.fm10(_108_v7, _108_v7, _108_v7, globalState));
      _113_v12 = _nw4;
      let _index1 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_113_v12).length));
      (_113_v12)[_index1] = _110_v9;
      let _124_v13;
      _124_v13 = _dafny.Map.Empty.slice().updateUnsafe(_109_v8,false);
      let _125_v15;
      _125_v15 = _dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0)));
      let _126_v16;
      _126_v16 = _dafny.Map.Empty.slice().updateUnsafe(_108_v7,false);
      let _127_v17;
      _127_v17 = _dafny.Seq.UnicodeFromString("rmwdor");
      let _128_v18;
      _128_v18 = _dafny.Seq.of(new BigNumber((_127_v17).length), _103_v2, new BigNumber((_127_v17).length));
      let _129_v19;
      _129_v19 = _module.D4.create_DC12(_104_v3);
      let _pat_let_tv7 = _108_v7;
      let _index2 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
      let _index3 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_113_v12).length));
      let _rhs0 = _107_v6;
      let _rhs1 = (((_125_v15).IsProperSubsetOf(function () {
        let _coll24 = new _dafny.Set();
        for (const _compr_24 of (_124_v13).Keys.Elements) {
          let _130_v14 = _compr_24;
          if ((_124_v13).contains(_130_v14)) {
            _coll24.add(_130_v14);
          }
        }
        return _coll24;
      }())) ? (_dafny.areEqual(_dafny.Seq.of(_103_v2, _103_v2, new BigNumber((_126_v16).length), _103_v2, new BigNumber((_126_v16).length)), _128_v18)) : (false));
      let _rhs2 = function (_source5) {
        if (_source5.is_DC13) {
          let _131___mcc_h0 = (_source5).cf21;
          let _132___mcc_h1 = (_source5).cf22;
          let _133_cf22 = _132___mcc_h1;
          let _134_cf21 = _131___mcc_h0;
          return _134_cf21;
        } else {
          let _135___mcc_h2 = (_source5).cf20;
          let _136_cf20 = _135___mcc_h2;
          return _pat_let_tv7;
        }
      }(_129_v19);
      let _rhs3 = _103_v2;
      let _rhs4 = _module.D2.create_DC6(((_108_v7) ? (new _dafny.CodePoint('x'.codePointAt(0))) : (_109_v8)), _108_v7, _108_v7, _103_v2);
      let _lhs0 = _100_v0;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
      let _lhs2 = _113_v12;
      let _lhs3 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_113_v12).length));
      _107_v6 = _rhs0;
      _lhs0[_lhs1] = _rhs1;
      r0 = _rhs2;
      _103_v2 = _rhs3;
      _lhs2[_lhs3] = _rhs4;
      let _137_v20;
      _137_v20 = _dafny.Map.Empty.slice().updateUnsafe(_103_v2,_103_v2);
      _103_v2 = new BigNumber((((_137_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(364),_103_v2))).Merge(function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of (_128_v18).Elements) {
          let _138_v21 = _compr_25;
          if (_dafny.Seq.contains(_128_v18, _138_v21)) {
            _coll25.push([(_138_v21).minus(_103_v2),_103_v2]);
          }
        }
        return _coll25;
      }())).length);
      _103_v2 = new BigNumber(-925);
      let _139_v22;
      _139_v22 = _dafny.MultiSet.fromElements(_108_v7);
      let _140_v23;
      _140_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_112_v11).length),_139_v22);
      let _141_v24;
      _141_v24 = _module.D1.create_DC3(_108_v7, _module.__default.safeModuloInt(_103_v2, new BigNumber((_140_v23).length)), _103_v2);
      let _source6 = _141_v24;
      if (_source6.is_DC3) {
        let _142___mcc_h3 = (_source6).cf3;
        let _143___mcc_h4 = (_source6).cf4;
        let _144___mcc_h5 = (_source6).cf5;
        let _145_cf5 = _144___mcc_h5;
        let _146_cf4 = _143___mcc_h4;
        let _147_cf3 = _142___mcc_h3;
        let _source7 = _module.D3.create_DC9(_145_cf5, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]);
        if (_source7.is_DC9) {
          let _148___mcc_h9 = (_source7).cf15;
          let _149___mcc_h10 = (_source7).cf16;
          let _150_cf16 = _149___mcc_h10;
          let _151_cf15 = _148___mcc_h9;
          let _152_v25;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_102_v1);
          _152_v25 = _nw5;
          _145_cf5 = _103_v2;
          let _153_v26;
          _153_v26 = _module.D4.create_DC13(_108_v7, _128_v18);
          let _154_v27;
          _154_v27 = _module.D4.create_DC13(false, (_153_v26).dtor_cf22);
          let _155_v28;
          let _nw6 = Array((new BigNumber(14)).toNumber());
          _nw6[(_dafny.ZERO).toNumber()] = _154_v27;
          _nw6[(_dafny.ONE).toNumber()] = _module.D4.create_DC13(_150_cf16, _128_v18);
          _nw6[(new BigNumber(2)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(3)).toNumber()] = _module.__default.fm65(false, _147_cf3, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], false, globalState);
          _nw6[(new BigNumber(4)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(5)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(6)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(7)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(8)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(9)).toNumber()] = _154_v27;
          _nw6[(new BigNumber(10)).toNumber()] = _module.D4.create_DC13(_150_cf16, _128_v18);
          _nw6[(new BigNumber(11)).toNumber()] = _module.D4.create_DC13((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _128_v18);
          _nw6[(new BigNumber(12)).toNumber()] = _153_v26;
          _nw6[(new BigNumber(13)).toNumber()] = _154_v27;
          _155_v28 = _nw6;
          let _index4 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_155_v28).length));
          (_155_v28)[_index4] = _153_v26;
          let _156_v29;
          _156_v29 = _dafny.Map.Empty.slice().updateUnsafe(_145_cf5,_147_cf3);
          let _index5 = _module.__default.safeIndex(new BigNumber(945), new BigNumber((_155_v28).length));
          (_155_v28)[_index5] = _module.D4.create_DC13((new BigNumber((_156_v29).length)).isLessThan(_145_cf5), _dafny.Seq.Concat(_128_v18, _dafny.Seq.of(_151_cf15)));
          let _157_v30;
          _157_v30 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_108_v7);
          let _158_v31;
          let _nw7 = new _module.C5();
          _nw7.__ctor((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], (_157_v30).Merge(_157_v30), _module.__default.fm2(_108_v7, globalState));
          _158_v31 = _nw7;
          _158_v31 = _158_v31;
        } else if (_source7.is_DC10) {
          let _159___mcc_h11 = (_source7).cf17;
          let _160___mcc_h12 = (_source7).cf18;
          let _161_cf18 = _160___mcc_h12;
          let _162_cf17 = _159___mcc_h11;
          let _163_v32;
          _163_v32 = _dafny.Seq.of(_dafny.Seq.of((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]), _112_v11, _112_v11);
          let _164_v33;
          _164_v33 = _module.D16.create_DC41(_163_v32);
          let _165_v34;
          _165_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_147_cf3, globalState),!(_108_v7));
          let _166_v36;
          _166_v36 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_145_cf5);
          let _167_v37;
          _167_v37 = _dafny.Map.Empty.slice().updateUnsafe((_128_v18)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_147_cf3,new BigNumber((_127_v17).length))).length), new BigNumber((_128_v18).length))],_108_v7);
          let _168_v38;
          let _nw8 = new _module.C8();
          _nw8.__ctor(_dafny.Seq.Concat(_163_v32, (_164_v33).dtor_cf69), (((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]) ? (!(_147_cf3)) : ((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))])), (_165_v34).Merge(function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_166_v36).Keys.Elements) {
              let _169_v35 = _compr_26;
              if ((_166_v36).contains(_169_v35)) {
                _coll26.push([_169_v35,(_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]]);
              }
            }
            return _coll26;
          }()), _dafny.Seq.UnicodeFromString("gtqsgwsq"), (new BigNumber((_167_v37).length)).minus(_162_cf17), _100_v0);
          _168_v38 = _nw8;
          let _rhs5 = (((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]) ? (_168_v38) : (_168_v38));
          let _rhs6 = _147_cf3;
          _168_v38 = _rhs5;
          _108_v7 = _rhs6;
          _103_v2 = _module.__default.safeModuloInt(_161_cf18, _103_v2);
          let _170_v39;
          let _nw9 = new _module.C5();
          _nw9.__ctor(_108_v7, _165_v34, _dafny.Seq.Concat((_168_v38).f2, (_168_v38).f2));
          _170_v39 = _nw9;
          _170_v39 = _170_v39;
          _147_cf3 = (((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]) ? (_108_v7) : (!(false) || (_108_v7)));
        } else if (_source7.is_DC8) {
          let _171___mcc_h13 = (_source7).cf14;
          let _172_cf14 = _171___mcc_h13;
          let _173_v40;
          let _nw10 = Array((new BigNumber(15)).toNumber()).fill(_module.D13.Default());
          _173_v40 = _nw10;
          let _174_v41;
          _174_v41 = _module.D22.create_DC54(_173_v40);
          let _175_v42;
          _175_v42 = _module.D22.create_DC56(_174_v41);
          let _176_v43;
          _176_v43 = _dafny.Map.Empty.slice().updateUnsafe((_103_v2).minus(_145_cf5),_175_v42);
          let _177_v44;
          _177_v44 = _dafny.Map.Empty.slice().updateUnsafe(_103_v2,_176_v43);
          let _rhs7 = (_103_v2).minus(new BigNumber((_127_v17).length));
          let _rhs8 = ((((_177_v44).contains((_dafny.ZERO).minus(_103_v2))) ? ((_177_v44).get((_dafny.ZERO).minus(_103_v2))) : (_176_v43))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(333)), ((_178_v2) => function (_179_i1) {
            return _178_v2;
          })(_103_v2))).length),_175_v42));
          _103_v2 = _rhs7;
          _176_v43 = _rhs8;
          let _index6 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_102_v1).length));
          (_102_v1)[_index6] = _146_cf4;
          let _index7 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_102_v1).length));
          (_102_v1)[_index7] = _103_v2;
          let _180_v45;
          _180_v45 = _dafny.Map.Empty.slice().updateUnsafe(_108_v7,_145_cf5);
          let _181_v46;
          let _init0 = function (_182_i2) {
            return _module.__default.safeModuloInt(_182_i2, new BigNumber(-154));
          };
          let _nw11 = Array((new BigNumber(12)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw11.length); _i0_0++) {
            _nw11[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _181_v46 = _nw11;
          let _index8 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_102_v1).length));
          (_102_v1)[_index8] = ((((_108_v7) ? (_108_v7) : ((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]))) ? (_146_cf4) : (new BigNumber((((_180_v45).update((_112_v11)[_module.__default.safeIndex((_dafny.ZERO).minus(_146_cf4), new BigNumber((_112_v11).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_181_v46, _181_v46)).length))))).update((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _146_cf4)).length)));
          let _183_v47;
          _183_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), ((_184_v8) => function (_185_i3) {
            return _184_v8;
          })(_109_v8)),_147_cf3);
          let _186_v48;
          let _nw12 = new _module.C3();
          _nw12.__ctor(new BigNumber(821), _100_v0, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _183_v47, _dafny.Seq.UnicodeFromString("dsqriyfwh"));
          _186_v48 = _nw12;
          let _187_v49;
          _187_v49 = _dafny.MultiSet.fromElements(_186_v48);
          let _188_v50;
          _188_v50 = _dafny.Map.Empty.slice().updateUnsafe(_146_cf4,_dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_189_v8) => function (_190_i4) {
            return _189_v8;
          })(_109_v8)));
          let _191_v51;
          _191_v51 = _dafny.Seq.of(_186_v48, _186_v48, _186_v48);
          let _index9 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
          let _rhs9 = (_172_cf14).IsSubsetOf(_172_cf14);
          let _rhs10 = _147_cf3;
          let _rhs11 = (((_188_v50).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_146_cf4)))) ? ((_188_v50).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_146_cf4)))) : (_127_v17));
          let _rhs12 = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_191_v51, _dafny.Seq.of(_186_v48)));
          let _lhs4 = _100_v0;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
          r0 = _rhs9;
          _lhs4[_lhs5] = _rhs10;
          _127_v17 = _rhs11;
          _187_v49 = _rhs12;
        } else {
          let _192___mcc_h14 = (_source7).cf19;
          let _193_cf19 = _192___mcc_h14;
          _103_v2 = (_dafny.ZERO).minus((_103_v2).plus(_module.__default.safeModuloInt(_146_cf4, _145_cf5)));
          let _194_v52;
          _194_v52 = _module.D3.create_DC9(_146_cf4, _147_cf3);
          _145_cf5 = (_194_v52).dtor_cf15;
          _147_cf3 = _147_cf3;
          _145_cf5 = _145_cf5;
        }
        if (_module.__default.fm3((new BigNumber(171)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_195_v8) => function (_196_i5) {
          return _195_v8;
        })(_109_v8))).length)), globalState)) {
          _108_v7 = (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))];
          let _index10 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
          (_100_v0)[_index10] = _147_cf3;
          _126_v16 = (_126_v16).update(_108_v7, false);
          let _197_v53;
          let _init1 = ((_198_v15, _199_cf5, _200_v7, _201_v2, _202_cf4, _203_v17, _204_v8, _205_v16, _206_v0) => function (_207_i6) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_module.D13.create_DC35(new BigNumber((_198_v15).length), _199_cf5, _200_v7, _201_v2, _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(_202_cf4,_203_v17)))).dtor_cf58),_module.D8.create_DC24(_204_v8, _205_v16)),(_206_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_206_v0).length))])).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_199_cf5,_module.D8.create_DC24(_204_v8, _205_v16)),_200_v7));
          })(_125_v15, _145_cf5, _108_v7, _103_v2, _146_cf4, _127_v17, _109_v8, _126_v16, _100_v0);
          let _nw13 = Array((new BigNumber(29)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw13.length); _i0_1++) {
            _nw13[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _197_v53 = _nw13;
          let _208_v54;
          _208_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(476),_module.__default.fm66(_109_v8, _146_cf4, new BigNumber((_dafny.Set.fromElements(_109_v8, _109_v8)).length), globalState));
          let _209_v55;
          _209_v55 = _dafny.Map.Empty.slice().updateUnsafe(_208_v54,(_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]);
          let _index11 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_197_v53).length));
          (_197_v53)[_index11] = _209_v55;
          let _index12 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_197_v53).length));
          (_197_v53)[_index12] = (_module.__default.fm67(false, globalState)).Merge(_209_v55);
          r0 = (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))];
        } else {
          let _210_v56;
          let _nw14 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _210_v56 = _nw14;
          let _211_v57;
          let _nw15 = Array((new BigNumber(6)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw15[(_dafny.ONE).toNumber()] = _109_v8;
          _nw15[(new BigNumber(2)).toNumber()] = _109_v8;
          _nw15[(new BigNumber(3)).toNumber()] = _109_v8;
          _nw15[(new BigNumber(4)).toNumber()] = _109_v8;
          _nw15[(new BigNumber(5)).toNumber()] = _109_v8;
          _211_v57 = _nw15;
          let _212_v58;
          _212_v58 = _dafny.Seq.of(_211_v57);
          let _213_v59;
          _213_v59 = _212_v58;
          let _index13 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_210_v56).length));
          (_210_v56)[_index13] = (_213_v59);
          let _214_v60;
          _214_v60 = _dafny.Set.fromElements(_108_v7);
          let _index14 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_210_v56).length));
          let _rhs13 = ((new BigNumber((_128_v18).length)).minus(new BigNumber((_214_v60).length))).multipliedBy(((_module.__default.fm3(_103_v2, globalState)) ? (new BigNumber(157)) : ((((_139_v22).contains(true)) ? ((_139_v22).get(true)) : (_146_cf4)))));
          let _rhs14 = (_dafny.ZERO).minus(_145_cf5);
          let _rhs15 = _145_cf5;
          let _rhs16 = _dafny.Seq.Concat(_212_v58, _212_v58);
          let _lhs6 = _210_v56;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_210_v56).length));
          _145_cf5 = _rhs13;
          _145_cf5 = _rhs14;
          _103_v2 = _rhs15;
          _lhs6[_lhs7] = _rhs16;
          _127_v17 = _127_v17;
          let _215_v61;
          _215_v61 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,(_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]);
          let _216_v62;
          let _nw16 = new _module.C1();
          _nw16.__ctor();
          _216_v62 = _nw16;
          let _217_v63;
          _217_v63 = _dafny.Seq.of(_216_v62, _216_v62, _216_v62);
          let _218_v64;
          let _nw17 = new _module.C5();
          _nw17.__ctor(!(_147_cf3), (_215_v61).Merge(_215_v61), _dafny.Seq.update(_dafny.Seq.Concat(_127_v17, _dafny.Seq.update(_127_v17, _module.__default.safeIndex(new BigNumber((_217_v63).length), new BigNumber((_127_v17).length)), _109_v8)), _module.__default.safeIndex((_dafny.ZERO).minus(_103_v2), new BigNumber((_dafny.Seq.Concat(_127_v17, _dafny.Seq.update(_127_v17, _module.__default.safeIndex(new BigNumber((_217_v63).length), new BigNumber((_127_v17).length)), _109_v8))).length)), new _dafny.CodePoint('s'.codePointAt(0))));
          _218_v64 = _nw17;
          let _219_v65;
          _219_v65 = _dafny.Seq.of(_218_v64, _218_v64, _218_v64);
          _218_v64 = (((_218_v64).fm11((_218_v64).f15, globalState)) ? (_218_v64) : ((_219_v65)[_module.__default.safeIndex(new BigNumber((_module.__default.fm2(_147_cf3, globalState)).length), new BigNumber((_219_v65).length))]));
          _147_cf3 = !((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]);
          let _220_v66;
          _220_v66 = _dafny.MultiSet.fromElements(_109_v8, _109_v8);
          let _221_v67;
          let _nw18 = new _module.C3();
          _nw18.__ctor((_dafny.ZERO).minus(_103_v2), _100_v0, !_dafny.Seq.contains(_128_v18, _146_cf4), _module.__default.fm62(_103_v2, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], globalState), _dafny.Seq.update(_127_v17, _module.__default.safeIndex(_103_v2, new BigNumber((_127_v17).length)), _109_v8));
          _221_v67 = _nw18;
          let _rhs17 = _dafny.MultiSet.fromElements(_109_v8);
          let _rhs18 = _108_v7;
          let _rhs19 = (((_108_v7) ? ((_221_v67).f4) : (_103_v2))).minus(new BigNumber(900));
          let _rhs20 = _221_v67;
          let _rhs21 = _103_v2;
          _220_v66 = _rhs17;
          _147_cf3 = _rhs18;
          _103_v2 = _rhs19;
          _221_v67 = _rhs20;
          _103_v2 = _rhs21;
        }
        let _222_v68;
        _222_v68 = _module.D4.create_DC13(_147_cf3, _128_v18);
        let _source8 = _222_v68;
        if (_source8.is_DC13) {
          let _223___mcc_h15 = (_source8).cf21;
          let _224___mcc_h16 = (_source8).cf22;
          let _225_cf22 = _224___mcc_h16;
          let _226_cf21 = _223___mcc_h15;
          _128_v18 = _225_cf22;
          _103_v2 = _146_cf4;
          _108_v7 = _module.__default.fm3((_128_v18)[_module.__default.safeIndex(_146_cf4, new BigNumber((_128_v18).length))], globalState);
          let _227_v69;
          _227_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_128_v18).length),_225_cf22);
          let _228_v70;
          _228_v70 = _dafny.Map.Empty.slice().updateUnsafe(_226_cf21,new BigNumber((_227_v69).length));
          let _index15 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_102_v1).length));
          (_102_v1)[_index15] = (((_228_v70).contains(_147_cf3)) ? ((_228_v70).get(_147_cf3)) : (_146_cf4));
          let _index16 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_102_v1).length));
          (_102_v1)[_index16] = new BigNumber((_225_cf22).length);
        } else {
          let _229___mcc_h17 = (_source8).cf20;
          let _230_cf20 = _229___mcc_h17;
          _103_v2 = (_146_cf4).plus(_103_v2);
          let _231_v71;
          _231_v71 = _dafny.Map.Empty.slice().updateUnsafe(_147_cf3,_145_cf5);
          let _232_v72;
          _232_v72 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_231_v71);
          _146_cf4 = new BigNumber((_232_v72).length);
          let _233_v73;
          _233_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_145_cf5),new _dafny.CodePoint('v'.codePointAt(0)));
          _233_v73 = (_233_v73).update(_103_v2, _109_v8);
          let _234_v74;
          _234_v74 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nfgd"),!(_108_v7));
          let _235_v75;
          let _nw19 = new _module.C11();
          _nw19.__ctor(_147_cf3, _146_cf4, _100_v0, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _234_v74, _127_v17);
          _235_v75 = _nw19;
          let _236_v76;
          _236_v76 = _dafny.MultiSet.fromElements(_235_v75, _235_v75, _235_v75, _235_v75);
          let _237_v77;
          _237_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_127_v17).length),_236_v76);
          _231_v71 = (_231_v71).update((_236_v76).IsSubsetOf((((_237_v77).contains(_145_cf5)) ? ((_237_v77).get(_145_cf5)) : (_236_v76))), _146_cf4);
        }
        let _238_v78;
        _238_v78 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_111_v10).length),new BigNumber((_dafny.Seq.of(false, _147_cf3)).length)),_103_v2);
        let _239_v79;
        _239_v79 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_108_v7);
        let _240_v80;
        let _nw20 = new _module.C10();
        _nw20.__ctor(new BigNumber((_104_v3).cardinality()), (_dafny.ZERO).minus(new BigNumber((_238_v78).length)), _100_v0, _147_cf3, _239_v79, _127_v17);
        _240_v80 = _nw20;
        let _241_v81;
        _241_v81 = _dafny.Set.fromElements(_240_v80);
        let _242_v82;
        _242_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_241_v81).length),_147_cf3);
        let _243_v83;
        _243_v83 = _dafny.Map.Empty.slice().updateUnsafe(_242_v82,_102_v1);
        let _244_v84;
        let _nw21 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _244_v84 = _nw21;
        let _index17 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_244_v84).length));
        (_244_v84)[_index17] = _109_v8;
        let _245_v86;
        _245_v86 = _dafny.Seq.of(function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of (_128_v18).Elements) {
            let _246_v85 = _compr_27;
            if (_dafny.Seq.contains(_128_v18, _246_v85)) {
              _coll27.add((_246_v85).minus(new BigNumber((_dafny.Seq.UnicodeFromString("onh")).length)));
            }
          }
          return _coll27;
        }(), (_111_v10).Union(_dafny.Set.fromElements(_145_cf5, new BigNumber((_127_v17).length), _146_cf4, new BigNumber(969))), _111_v10);
        let _index18 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_244_v84).length));
        let _rhs22 = (_245_v86)[_module.__default.safeIndex(_103_v2, new BigNumber((_245_v86).length))];
        let _rhs23 = _243_v83;
        let _rhs24 = _103_v2;
        let _rhs25 = _module.__default.fm38(!(_147_cf3), _240_v80.f3, _145_cf5, globalState);
        let _rhs26 = _240_v80;
        let _lhs8 = _244_v84;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_244_v84).length));
        _111_v10 = _rhs22;
        _243_v83 = _rhs23;
        _146_cf4 = _rhs24;
        _lhs8[_lhs9] = _rhs25;
        _240_v80 = _rhs26;
      } else if (_source6.is_DC4) {
        let _247___mcc_h6 = (_source6).cf6;
        let _248___mcc_h7 = (_source6).cf7;
        let _249_cf7 = _248___mcc_h7;
        let _250_cf6 = _247___mcc_h6;
        let _251_v87;
        _251_v87 = _module.D23.create_DC58();
        let _252_v88;
        _252_v88 = _dafny.MultiSet.fromElements(_251_v87, _module.D23.create_DC58());
        _252_v88 = _252_v88;
        let _253_v89;
        _253_v89 = _dafny.Map.Empty.slice().updateUnsafe(_127_v17,_250_cf6);
        let _254_v90;
        let _nw22 = new _module.C11();
        _nw22.__ctor(_249_cf7, _103_v2, _100_v0, _module.__default.fm3(_103_v2, globalState), (_253_v89).Merge(_253_v89), _127_v17);
        _254_v90 = _nw22;
        let _255_v91;
        _255_v91 = _dafny.MultiSet.fromElements(_137_v20, _137_v20);
        _255_v91 = (_255_v91).Intersect((_dafny.MultiSet.fromElements(function () {
          let _coll28 = new _dafny.Map();
          for (const _compr_28 of _dafny.IntegerRange(new BigNumber(-393), new BigNumber(105))) {
            let _256_v92 = _compr_28;
            if (((new BigNumber(-393)).isLessThanOrEqualTo(_256_v92)) && ((_256_v92).isLessThan(new BigNumber(105)))) {
              _coll28.push([(_256_v92).minus(_103_v2),_103_v2]);
            }
          }
          return _coll28;
        }())).Intersect(_255_v91));
        _103_v2 = _103_v2;
      } else {
        let _257___mcc_h8 = (_source6).cf2;
        let _258_cf2 = _257___mcc_h8;
        _103_v2 = _103_v2;
        let _259_v93;
        _259_v93 = _dafny.Map.Empty.slice().updateUnsafe(_103_v2,_127_v17);
        let _260_v94;
        _260_v94 = _module.D11.create_DC30(_259_v93);
        let _261_v95;
        _261_v95 = _module.D13.create_DC35(_103_v2, _103_v2, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], new BigNumber((_112_v11).length), _260_v94);
        let _262_v96;
        _262_v96 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_258_cf2, _module.__default.safeIndex(_103_v2, new BigNumber((_258_cf2).length)), _109_v8),true);
        let _263_v97;
        let _nw23 = new _module.C11();
        _nw23.__ctor((_104_v3).IsProperSubsetOf(_dafny.MultiSet.fromElements(_module.__default.fm0(globalState), (_261_v95).dtor_cf60)), _103_v2, _100_v0, false, (_262_v96).Merge(_262_v96), _258_cf2);
        _263_v97 = _nw23;
        let _264_v98;
        _264_v98 = _module.D6.create_DC18(_103_v2, _103_v2, (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _109_v8, new BigNumber(227));
        let _pat_let_tv8 = _100_v0;
        let _pat_let_tv9 = _100_v0;
        let _265_v99;
        _265_v99 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let10_0) {
          return function (_266_dt__update__tmp_h3) {
            return function (_pat_let11_0) {
              return function (_267_dt__update_hcf21_h0) {
                return _module.D4.create_DC13(_267_dt__update_hcf21_h0, (_266_dt__update__tmp_h3).dtor_cf22);
              }(_pat_let11_0);
            }((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_pat_let_tv8).length))]);
          }(_pat_let10_0);
        }(_module.D4.create_DC13(!((_263_v97).f6), _128_v18)),_258_cf2);
        let _268_v100;
        _268_v100 = _module.D4.create_DC13((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], _128_v18);
        let _269_v101;
        let _nw24 = Array((new BigNumber(16)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_127_v17, (_263_v97).fm14(_108_v7, (_263_v97).f6, _109_v8, _103_v2, globalState));
        _nw24[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_127_v17, _module.__default.safeIndex(_103_v2, new BigNumber((_127_v17).length)), _109_v8), _127_v17);
        _nw24[(new BigNumber(2)).toNumber()] = _127_v17;
        _nw24[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat((_263_v97).fm14((_263_v97).f6, (_110_v9).dtor_cf10, _109_v8, _103_v2, globalState), _dafny.Seq.of((_127_v17)[_module.__default.safeIndex(_103_v2, new BigNumber((_127_v17).length))], _109_v8, _109_v8));
        _nw24[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_127_v17, _258_cf2);
        _nw24[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_258_cf2, _dafny.Seq.of(new _dafny.CodePoint('c'.codePointAt(0))));
        _nw24[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_258_cf2, _module.__default.safeIndex(_103_v2, new BigNumber((_258_cf2).length)), (_264_v98).dtor_cf32);
        _nw24[(new BigNumber(7)).toNumber()] = _127_v17;
        _nw24[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_127_v17, _dafny.Seq.update(_dafny.Seq.of(_module.__default.fm38(true, _108_v7, _103_v2, globalState)), _module.__default.safeIndex(new BigNumber((_127_v17).length), new BigNumber((_dafny.Seq.of(_module.__default.fm38(true, _108_v7, _103_v2, globalState))).length)), new _dafny.CodePoint('y'.codePointAt(0))));
        _nw24[(new BigNumber(9)).toNumber()] = (((_265_v99).contains(_268_v100)) ? ((_265_v99).get(_268_v100)) : (_127_v17));
        _nw24[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), ((_270_v8) => function (_271_i7) {
          return _270_v8;
        })(_109_v8));
        _nw24[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_109_v8);
        _nw24[(new BigNumber(12)).toNumber()] = _127_v17;
        _nw24[(new BigNumber(13)).toNumber()] = _258_cf2;
        _nw24[(new BigNumber(14)).toNumber()] = _127_v17;
        _nw24[(new BigNumber(15)).toNumber()] = _127_v17;
        _269_v101 = _nw24;
        let _index19 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_269_v101).length));
        (_269_v101)[_index19] = _127_v17;
        let _272_v103;
        _272_v103 = _dafny.MultiSet.fromElements(_109_v8);
        let _index20 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
        let _index21 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_269_v101).length));
        let _rhs27 = (function () {
          let _coll29 = new _dafny.Set();
          for (const _compr_29 of (_dafny.MultiSet.fromElements(new BigNumber(594))).Elements) {
            let _273_v102 = _compr_29;
            if ((_dafny.MultiSet.fromElements(new BigNumber(594))).contains(_273_v102)) {
              _coll29.add(_module.__default.safeDivisionInt(_273_v102, new BigNumber(-109)));
            }
          }
          return _coll29;
        }()).IsDisjointFrom(_111_v10);
        let _rhs28 = ((_dafny.ZERO).minus((_103_v2).plus(_103_v2))).minus(new BigNumber((_dafny.Seq.of((_263_v97).f6, (_263_v97).f6)).length));
        let _rhs29 = (_dafny.ZERO).minus(((!_dafny.Seq.contains(_258_cf2, _109_v8)) ? (new BigNumber((_272_v103).cardinality())) : (_103_v2)));
        let _rhs30 = _dafny.Seq.Concat(_127_v17, _dafny.Seq.Concat(_258_cf2, (_263_v97).fm14((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], (_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))], new _dafny.CodePoint('b'.codePointAt(0)), _103_v2, globalState)));
        let _rhs31 = _103_v2;
        let _lhs10 = _100_v0;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length));
        let _lhs12 = _269_v101;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(38), new BigNumber((_269_v101).length));
        _lhs10[_lhs11] = _rhs27;
        _103_v2 = _rhs28;
        _103_v2 = _rhs29;
        _lhs12[_lhs13] = _rhs30;
        _103_v2 = _rhs31;
        if ((_263_v97).f6) {
          let _274_v104;
          _274_v104 = _dafny.Map.Empty.slice().updateUnsafe(((_108_v7) ? (_module.__default.fm22(true, _103_v2, globalState)) : (new _dafny.CodePoint('m'.codePointAt(0)))),_module.__default.safeModuloInt(_103_v2, _103_v2));
          _274_v104 = (_274_v104).update(_109_v8, _103_v2);
          let _rhs32 = _103_v2;
          let _rhs33 = _269_v101;
          _103_v2 = _rhs32;
          _269_v101 = _rhs33;
          let _index22 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_102_v1).length));
          (_102_v1)[_index22] = _103_v2;
          let _275_v105;
          let _nw25 = new _module.C7();
          _nw25.__ctor(_103_v2, _100_v0, false, _262_v96, _dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), ((_276_v8) => function (_277_i8) {
            return _276_v8;
          })(_109_v8)));
          _275_v105 = _nw25;
          let _278_v106;
          _278_v106 = _dafny.Map.Empty.slice().updateUnsafe(_103_v2,_108_v7);
          let _279_v107;
          _279_v107 = _dafny.Set.fromElements(_275_v105, _275_v105, (_module.D25.create_DC62(_275_v105, new BigNumber((_278_v106).length), (_263_v97).fm11((_263_v97).f6, globalState))).dtor_cf94);
          let _280_v108;
          _280_v108 = _dafny.Seq.of(_279_v107);
          let _index23 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_102_v1).length));
          (_102_v1)[_index23] = new BigNumber((_280_v108).length);
          let _281_v109;
          let _nw26 = new _module.C8();
          _nw26.__ctor(_dafny.Seq.of(_112_v11, _112_v11, _112_v11, _112_v11), (_263_v97).f6, _262_v96, _127_v17, (_102_v1)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_102_v1).length))], _100_v0);
          _281_v109 = _nw26;
          let _282_v110;
          _282_v110 = _dafny.Seq.of(_281_v109);
          let _283_v111;
          _283_v111 = _module.D19.create_DC48(_282_v110);
          let _284_v112;
          let _nw27 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _284_v112 = _nw27;
          let _index24 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_284_v112).length));
          (_284_v112)[_index24] = _109_v8;
          let _index25 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_284_v112).length));
          let _rhs34 = _283_v111;
          let _rhs35 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(449)), ((_285_v106) => function (_286_i9) {
            return new BigNumber((_285_v106).length);
          })(_278_v106));
          let _rhs36 = _109_v8;
          let _rhs37 = (_263_v97).f6;
          let _lhs14 = _284_v112;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_284_v112).length));
          _283_v111 = _rhs34;
          _128_v18 = _rhs35;
          _lhs14[_lhs15] = _rhs36;
          r1 = _rhs37;
          let _287_v113;
          let _nw28 = Array((new BigNumber(16)).toNumber()).fill(_module.D2.Default());
          _287_v113 = _nw28;
          let _288_v114;
          let _out0;
          _out0 = (_263_v97).m2(_287_v113, globalState);
          _288_v114 = _out0;
        } else {
          _103_v2 = ((new BigNumber(((_269_v101)[_module.__default.safeIndex(new BigNumber(38), new BigNumber((_269_v101).length))]).length)).minus(_103_v2)).plus(new BigNumber((_dafny.Seq.Concat(_112_v11, _112_v11)).length));
          let _289_v115;
          _289_v115 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),!((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]));
          r1 = !((((_289_v115).contains(_103_v2)) ? ((_289_v115).get(_103_v2)) : (_108_v7))) || ((_100_v0)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_100_v0).length))]);
          r1 = _module.__default.fm3(_103_v2, globalState);
          _137_v20 = (_137_v20).update(_103_v2, ((_dafny.ZERO).minus((_dafny.ZERO).minus(_103_v2))).minus(_103_v2));
          let _290_v116;
          let _nw29 = new _module.C0();
          _nw29.__ctor(_102_v1);
          _290_v116 = _nw29;
        }
      }
      r0 = true;
      let _291_v117;
      _291_v117 = _dafny.MultiSet.fromElements(_128_v18, _128_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), ((_292_v2) => function (_293_i10) {
        return _292_v2;
      })(_103_v2)));
      r1 = (new BigNumber((_291_v117).cardinality())).isEqualTo(_103_v2);
      r2 = _module.__default.fm7(_108_v7, _108_v7, globalState);
      let _294_v118;
      _294_v118 = _dafny.Seq.of((_137_v20).update(_103_v2, (_dafny.ZERO).minus(_103_v2)), _137_v20, _137_v20, _137_v20, _dafny.Map.Empty.slice().updateUnsafe(_103_v2,_103_v2));
      r3 = _294_v118;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _295_globalState;
      let _nw30 = new _module.GlobalState();
      _nw30.__ctor(true);
      _295_globalState = _nw30;
      let _296_v0;
      _296_v0 = true;
      let _297_v1;
      let _init2 = function (_298_i0) {
        return !(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(409)), function (_299_i1) {
          return new BigNumber((_dafny.Seq.UnicodeFromString("ejsfewk")).length);
        }), _dafny.Seq.of(new BigNumber(902), new BigNumber(812), new BigNumber(389), new BigNumber(649))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(485), new BigNumber(593))))));
      };
      let _nw31 = Array((new BigNumber(15)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw31.length); _i0_2++) {
        _nw31[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _297_v1 = _nw31;
      let _300_v2;
      _300_v2 = new BigNumber(59);
      let _rhs38 = (_300_v2).isLessThan(_module.__default.fm0(_295_globalState));
      let _rhs39 = (!(!(false))) && ((_300_v2).isLessThanOrEqualTo(_300_v2));
      let _rhs40 = _297_v1;
      let _rhs41 = ((_296_v0) ? (_296_v0) : (_296_v0));
      let _rhs42 = (_300_v2).minus(_300_v2);
      _296_v0 = _rhs38;
      _296_v0 = _rhs39;
      _297_v1 = _rhs40;
      _296_v0 = _rhs41;
      _300_v2 = _rhs42;
      let _301_v3;
      let _302_v4;
      let _303_v5;
      let _304_v6;
      let _out1;
      let _out2;
      let _out3;
      let _out4;
      let _outcollector0 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe(_297_v1,_296_v0), _295_globalState);
      _out1 = _outcollector0[0];
      _out2 = _outcollector0[1];
      _out3 = _outcollector0[2];
      _out4 = _outcollector0[3];
      _301_v3 = _out1;
      _302_v4 = _out2;
      _303_v5 = _out3;
      _304_v6 = _out4;
      let _305_v7;
      let _init3 = ((_306_v2) => function (_307_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_306_v2,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), function (_308_i3) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), ((_309_v2) => function (_310_i4) {
          return _309_v2;
        })(_306_v2))).length),_306_v2));
      })(_300_v2);
      let _nw32 = Array((new BigNumber(2)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw32.length); _i0_3++) {
        _nw32[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _305_v7 = _nw32;
      let _311_v8;
      _311_v8 = _dafny.Seq.of(_303_v5);
      let _312_v9;
      _312_v9 = _dafny.MultiSet.fromElements(_296_v0, _301_v3);
      let _313_v10;
      _313_v10 = _dafny.Seq.of(_312_v9);
      let _314_v11;
      _314_v11 = new _dafny.CodePoint('p'.codePointAt(0));
      let _index26 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_305_v7).length));
      (_305_v7)[_index26] = _module.__default.fm1((_311_v8)[_module.__default.safeIndex(new BigNumber(((_313_v10)[_module.__default.safeIndex(_300_v2, new BigNumber((_313_v10).length))]).cardinality()), new BigNumber((_311_v8).length))], _300_v2, _300_v2, _314_v11, _295_globalState);
      let _315_v12;
      _315_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_314_v11, _314_v11),_296_v0);
      let _316_v13;
      _316_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_315_v12).length),_300_v2);
      let _317_v14;
      _317_v14 = _module.D0.create_DC0(_300_v2);
      let _index27 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_305_v7).length));
      (_305_v7)[_index27] = (_316_v13).Merge((_316_v13).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_295_globalState),(_317_v14).dtor_cf0)));
      _302_v4 = !((_300_v2).isLessThan(_300_v2));
      let _318_v15;
      let _init4 = ((_319_v2) => function (_320_i5) {
        return (_320_i5).minus(_319_v2);
      })(_300_v2);
      let _nw33 = Array((new BigNumber(5)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw33.length); _i0_4++) {
        _nw33[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _318_v15 = _nw33;
      let _index28 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
      (_318_v15)[_index28] = (_module.__default.fm0(_295_globalState)).multipliedBy(new BigNumber((_dafny.Seq.of(_302_v4)).length));
      let _pat_let_tv10 = _312_v9;
      let _pat_let_tv11 = _301_v3;
      let _pat_let_tv12 = _312_v9;
      let _pat_let_tv13 = _301_v3;
      let _pat_let_tv14 = _300_v2;
      let _pat_let_tv15 = _300_v2;
      let _index29 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
      let _rhs43 = _318_v15;
      let _rhs44 = function (_source9) {
        if (_source9.is_DC1) {
          let _321___mcc_h0 = (_source9).cf1;
          let _322_cf1 = _321___mcc_h0;
          if ((_pat_let_tv10).contains(_pat_let_tv11)) {
            return (_pat_let_tv12).get(_pat_let_tv13);
          } else {
            return (_dafny.ZERO).minus(_pat_let_tv14);
          }
        } else {
          let _323___mcc_h1 = (_source9).cf0;
          let _324_cf0 = _323___mcc_h1;
          return _pat_let_tv15;
        }
      }(_317_v14);
      let _rhs45 = _module.__default.safeModuloInt(_300_v2, _300_v2);
      let _rhs46 = _317_v14;
      let _rhs47 = _300_v2;
      let _lhs16 = _318_v15;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
      _318_v15 = _rhs43;
      _300_v2 = _rhs44;
      _300_v2 = _rhs45;
      _317_v14 = _rhs46;
      _lhs16[_lhs17] = _rhs47;
      let _325_i6;
      _325_i6 = _dafny.ZERO;
      L0: {
        while ((((_315_v12).contains(_dafny.Seq.UnicodeFromString("yjwie"))) ? ((_315_v12).get(_dafny.Seq.UnicodeFromString("yjwie"))) : (_302_v4))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_325_i6)) {
              break L0;
            }
            _325_i6 = (_325_i6).plus(_dafny.ONE);
            _300_v2 = (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))];
            let _326_v16;
            _326_v16 = _dafny.Map.Empty.slice().updateUnsafe(_297_v1,_302_v4);
            let _327_v17;
            let _328_v18;
            let _329_v19;
            let _330_v20;
            let _out5;
            let _out6;
            let _out7;
            let _out8;
            let _outcollector1 = _module.__default.m0(_326_v16, _295_globalState);
            _out5 = _outcollector1[0];
            _out6 = _outcollector1[1];
            _out7 = _outcollector1[2];
            _out8 = _outcollector1[3];
            _327_v17 = _out5;
            _328_v18 = _out6;
            _329_v19 = _out7;
            _330_v20 = _out8;
            let _331_v21;
            _331_v21 = _dafny.Set.fromElements(_300_v2);
            let _index30 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
            (_318_v15)[_index30] = (new BigNumber((((_328_v18) ? (_329_v19) : (_303_v5))).length)).plus((new BigNumber((_331_v21).length)).minus(_300_v2));
            let _332_v22;
            _332_v22 = _dafny.Seq.UnicodeFromString("wtuvm");
            _332_v22 = _dafny.Seq.Concat(_332_v22, _332_v22);
          }
        }
      }
      _300_v2 = _module.__default.safeDivisionInt((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _300_v2);
      let _333_v23;
      _333_v23 = _dafny.Seq.UnicodeFromString("xyfj");
      let _334_v24;
      _334_v24 = _dafny.Seq.of(new BigNumber((_333_v23).length));
      let _335_v25;
      _335_v25 = _dafny.Map.Empty.slice().updateUnsafe(_301_v3,new BigNumber((_334_v24).length));
      let _336_v26;
      _336_v26 = _dafny.Seq.of(_335_v25);
      let _337_v27;
      _337_v27 = _dafny.Map.Empty.slice().updateUnsafe((_300_v2).isLessThanOrEqualTo(new BigNumber(150)),_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))],_302_v4)).length), new BigNumber(((_336_v26)[_module.__default.safeIndex(_300_v2, new BigNumber((_336_v26).length))]).length)));
      _335_v25 = _337_v27;
      let _338_v28;
      _338_v28 = _dafny.Map.Empty.slice().updateUnsafe((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))],_333_v23);
      _333_v23 = _dafny.Seq.Concat(_333_v23, _dafny.Seq.update((((_338_v28).contains(_300_v2)) ? ((_338_v28).get(_300_v2)) : (_module.__default.fm2(!(_302_v4), _295_globalState))), _module.__default.safeIndex((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], new BigNumber(((((_338_v28).contains(_300_v2)) ? ((_338_v28).get(_300_v2)) : (_module.__default.fm2(!(_302_v4), _295_globalState)))).length)), new _dafny.CodePoint('q'.codePointAt(0))));
      let _339_i7;
      _339_i7 = _dafny.ZERO;
      L1: {
        while (_302_v4) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_339_i7)) {
              break L1;
            }
            _339_i7 = (_339_i7).plus(_dafny.ONE);
            let _340_v29;
            let _nw34 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
            _340_v29 = _nw34;
            let _rhs48 = _dafny.Seq.update(_303_v5, _module.__default.safeIndex(((((_337_v27).contains(false)) ? ((_337_v27).get(false)) : ((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))).plus(new BigNumber(796)), new BigNumber((_303_v5).length)), _302_v4);
            let _rhs49 = _340_v29;
            _303_v5 = _rhs48;
            _340_v29 = _rhs49;
            let _341_v30;
            _341_v30 = _module.D0.create_DC1(_300_v2);
            let _source10 = _341_v30;
            if (_source10.is_DC1) {
              let _342___mcc_h2 = (_source10).cf1;
              let _343_cf1 = _342___mcc_h2;
              _302_v4 = !(_301_v3);
              _343_cf1 = _300_v2;
              let _344_v31;
              _344_v31 = _module.D1.create_DC2(_dafny.Seq.UnicodeFromString("hrs"));
              let _345_v32;
              let _nw35 = Array((new BigNumber(23)).toNumber());
              _nw35[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("gkmnvysfk");
              _nw35[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_333_v23, _333_v23);
              _nw35[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_333_v23, _333_v23);
              _nw35[(new BigNumber(3)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(4)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(5)).toNumber()] = _dafny.Seq.update((_344_v31).dtor_cf2, _module.__default.safeIndex((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], new BigNumber(((_344_v31).dtor_cf2).length)), _314_v11);
              _nw35[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_344_v31).dtor_cf2, _333_v23);
              _nw35[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(385)), ((_346_v11) => function (_347_i8) {
                return _346_v11;
              })(_314_v11));
              _nw35[(new BigNumber(8)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_333_v23, _dafny.Seq.UnicodeFromString("up"));
              _nw35[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_333_v23, _333_v23);
              _nw35[(new BigNumber(11)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(12)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(13)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(14)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(15)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(16)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(17)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(18)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_348_v11) => function (_349_i9) {
                return _348_v11;
              })(_314_v11)), _dafny.Seq.UnicodeFromString("xpfqtw"));
              _nw35[(new BigNumber(20)).toNumber()] = _333_v23;
              _nw35[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("jqpxpwr");
              _nw35[(new BigNumber(22)).toNumber()] = _333_v23;
              _345_v32 = _nw35;
              let _index31 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_345_v32).length));
              (_345_v32)[_index31] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-950)), ((_350_v11) => function (_351_i10) {
                return _350_v11;
              })(_314_v11));
              let _index32 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_345_v32).length));
              (_345_v32)[_index32] = _333_v23;
              let _352_v33;
              _352_v33 = _dafny.Seq.of(_334_v24, ((_301_v3) ? (_334_v24) : (_334_v24)), _334_v24);
              let _353_v34;
              _353_v34 = _dafny.Map.Empty.slice().updateUnsafe(_296_v0,_dafny.Seq.UnicodeFromString("hdb"));
              _352_v33 = _dafny.Seq.update(_dafny.Seq.Concat(_352_v33, _352_v33), _module.__default.safeIndex(new BigNumber((_353_v34).length), new BigNumber((_dafny.Seq.Concat(_352_v33, _352_v33)).length)), _334_v24);
            } else {
              let _354___mcc_h3 = (_source10).cf0;
              let _355_cf0 = _354___mcc_h3;
              let _356_v35;
              _356_v35 = _dafny.Map.Empty.slice().updateUnsafe(_296_v0,_301_v3);
              let _357_v36;
              _357_v36 = _dafny.Map.Empty.slice().updateUnsafe((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))],_297_v1);
              let _358_v37;
              _358_v37 = _dafny.Seq.of(_333_v23);
              let _index33 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              let _rhs50 = new BigNumber((_357_v36).length);
              let _rhs51 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_358_v37, _333_v23),_module.__default.fm3(_355_cf0, _295_globalState));
              let _rhs52 = (_module.__default.safeModuloInt(new BigNumber(215), (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))])).plus(_300_v2);
              let _lhs18 = _318_v15;
              let _lhs19 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              _lhs18[_lhs19] = _rhs50;
              _356_v35 = _rhs51;
              _355_cf0 = _rhs52;
              let _359_v38;
              _359_v38 = _dafny.Map.Empty.slice().updateUnsafe(_334_v24,_module.__default.fm0(_295_globalState));
              let _360_v39;
              _360_v39 = _module.D2.create_DC5(_337_v27);
              let _361_v40;
              _361_v40 = _dafny.Set.fromElements(false);
              let _362_v41;
              _362_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_361_v40).length),true);
              let _363_v42;
              let _nw36 = Array((new BigNumber(24)).toNumber());
              _nw36[(_dafny.ZERO).toNumber()] = _335_v25;
              _nw36[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_301_v3,_300_v2)).Merge((((_360_v39).dtor_cf8).update(_296_v0, new BigNumber(500))).update(_296_v0, new BigNumber((_361_v40).length)));
              _nw36[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_302_v4,(_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]);
              _nw36[(new BigNumber(3)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(4)).toNumber()] = (_337_v27).Merge((_336_v26)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_336_v26).length))]);
              _nw36[(new BigNumber(5)).toNumber()] = (_337_v27).update(!(_302_v4), (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]);
              _nw36[(new BigNumber(6)).toNumber()] = _module.__default.fm4(_300_v2, (((_362_v41).contains((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))])) ? ((_362_v41).get((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))])) : (_module.__default.fm3((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _295_globalState))), _295_globalState);
              _nw36[(new BigNumber(7)).toNumber()] = _337_v27;
              _nw36[(new BigNumber(8)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(9)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_303_v5).length));
              _nw36[(new BigNumber(11)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(12)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(13)).toNumber()] = _337_v27;
              _nw36[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_303_v5)[_module.__default.safeIndex((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], new BigNumber((_303_v5).length))],new BigNumber(((_305_v7)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_305_v7).length))]).length))).Merge(_335_v25);
              _nw36[(new BigNumber(15)).toNumber()] = _337_v27;
              _nw36[(new BigNumber(16)).toNumber()] = _module.__default.fm4(_300_v2, _296_v0, _295_globalState);
              _nw36[(new BigNumber(17)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(18)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(19)).toNumber()] = _337_v27;
              _nw36[(new BigNumber(20)).toNumber()] = _335_v25;
              _nw36[(new BigNumber(21)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_302_v4,(_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))])).Merge(_337_v27);
              _nw36[(new BigNumber(22)).toNumber()] = _337_v27;
              _nw36[(new BigNumber(23)).toNumber()] = _335_v25;
              _363_v42 = _nw36;
              let _index34 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_363_v42).length));
              (_363_v42)[_index34] = (_337_v27).Merge(_337_v27);
              let _364_v43;
              _364_v43 = _dafny.Seq.of(_359_v38);
              let _365_v44;
              _365_v44 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))],_296_v0),new BigNumber(169));
              let _index35 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_363_v42).length));
              let _index36 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              let _index37 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              let _rhs53 = ((_359_v38).Merge((_364_v43)[_module.__default.safeIndex(new BigNumber(-648), new BigNumber((_364_v43).length))])).Merge((_359_v38).Merge(_359_v38));
              let _rhs54 = _356_v35;
              let _rhs55 = ((_337_v27).update(_302_v4, _300_v2)).Merge((_337_v27).update(!(_302_v4), _355_cf0));
              let _rhs56 = new BigNumber((_365_v44).length);
              let _rhs57 = _module.__default.safeModuloInt(_355_cf0, (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]);
              let _lhs20 = _363_v42;
              let _lhs21 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_363_v42).length));
              let _lhs22 = _318_v15;
              let _lhs23 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              let _lhs24 = _318_v15;
              let _lhs25 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              _359_v38 = _rhs53;
              _356_v35 = _rhs54;
              _lhs20[_lhs21] = _rhs55;
              _lhs22[_lhs23] = _rhs56;
              _lhs24[_lhs25] = _rhs57;
              let _index38 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
              (_318_v15)[_index38] = _300_v2;
              let _366_v45;
              let _nw37 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
              _366_v45 = _nw37;
              let _367_v46;
              _367_v46 = _dafny.Seq.of(_297_v1, _297_v1);
              let _index39 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_366_v45).length));
              (_366_v45)[_index39] = ((_302_v4) ? (_367_v46) : (_dafny.Seq.of(_297_v1)));
              let _index40 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_366_v45).length));
              (_366_v45)[_index40] = _dafny.Seq.Concat(_367_v46, _367_v46);
            }
            _297_v1 = _297_v1;
            _302_v4 = _301_v3;
          }
        }
      }
      let _368_v47;
      _368_v47 = _module.D1.create_DC4(_296_v0, _dafny.Seq.contains(_334_v24, new BigNumber((_333_v23).length)));
      _368_v47 = _module.D1.create_DC4(_dafny.Seq.IsProperPrefixOf(_303_v5, _303_v5), _301_v3);
      let _369_v48;
      let _init5 = ((_370_v3, _371_v15) => function (_372_i11) {
        return _module.D1.create_DC3(_370_v3, (_371_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_371_v15).length))], (_371_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_371_v15).length))]);
      })(_301_v3, _318_v15);
      let _nw38 = Array((new BigNumber(24)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw38.length); _i0_5++) {
        _nw38[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _369_v48 = _nw38;
      _369_v48 = _369_v48;
      let _373_i12;
      _373_i12 = _dafny.ZERO;
      L2: {
        while (_296_v0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_373_i12)) {
              break L2;
            }
            _373_i12 = (_373_i12).plus(_dafny.ONE);
            let _374_v49;
            _374_v49 = _dafny.MultiSet.fromElements(_300_v2);
            let _375_v50;
            _375_v50 = _dafny.Map.Empty.slice().updateUnsafe(_374_v49,_333_v23);
            _375_v50 = (_dafny.Map.Empty.slice().updateUnsafe(_374_v49,_333_v23)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_374_v49,_333_v23));
            _300_v2 = (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))];
            _300_v2 = (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))];
            let _376_v51;
            _376_v51 = _dafny.Map.Empty.slice().updateUnsafe(_297_v1,_296_v0);
            let _377_v52;
            let _378_v53;
            let _379_v54;
            let _380_v55;
            let _out9;
            let _out10;
            let _out11;
            let _out12;
            let _outcollector2 = _module.__default.m0(_376_v51, _295_globalState);
            _out9 = _outcollector2[0];
            _out10 = _outcollector2[1];
            _out11 = _outcollector2[2];
            _out12 = _outcollector2[3];
            _377_v52 = _out9;
            _378_v53 = _out10;
            _379_v54 = _out11;
            _380_v55 = _out12;
          }
        }
      }
      let _hi0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(_300_v2), _300_v2));
      for (let _381_i13 = new BigNumber(-33); _381_i13.isLessThan(_hi0); _381_i13 = _381_i13.plus(_dafny.ONE)) {
        let _index41 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_297_v1).length));
        (_297_v1)[_index41] = _302_v4;
        let _382_v56;
        _382_v56 = _module.D3.create_DC8(_dafny.Set.fromElements(_302_v4, _301_v3, _302_v4));
        let _index42 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_297_v1).length));
        (_297_v1)[_index42] = _module.__default.fm3(new BigNumber(((_382_v56).dtor_cf14).length), _295_globalState);
        let _index43 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
        (_318_v15)[_index43] = ((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]).plus(_381_i13);
        _300_v2 = _300_v2;
        let _383_v57;
        _383_v57 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(262)), ((_384_v2) => function (_385_i14) {
          return _384_v2;
        })(_300_v2)),_333_v23);
        _300_v2 = new BigNumber((_dafny.Seq.update((((_383_v57).contains(_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))) ? ((_383_v57).get(_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(45)), ((_386_v11) => function (_387_i15) {
          return _386_v11;
        })(_314_v11)))), _module.__default.safeIndex(_381_i13, new BigNumber(((((_383_v57).contains(_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))) ? ((_383_v57).get(_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(45)), ((_388_v11) => function (_389_i15) {
          return _388_v11;
        })(_314_v11))))).length)), _314_v11)).length);
      }
      let _index44 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
      (_297_v1)[_index44] = _296_v0;
      let _index45 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
      (_297_v1)[_index45] = _302_v4;
      let _hi1 = _module.__default.safeModuloInt(_300_v2, (_dafny.ZERO).minus(_300_v2));
      for (let _390_i16 = (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]; _390_i16.isLessThan(_hi1); _390_i16 = _390_i16.plus(_dafny.ONE)) {
        let _index46 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
        (_297_v1)[_index46] = (_301_v3) || (_301_v3);
        let _391_v58;
        _391_v58 = _dafny.Map.Empty.slice().updateUnsafe(_301_v3,_dafny.Map.Empty.slice().updateUnsafe(_296_v0,_390_i16));
        let _source11 = _module.D2.create_DC5((_335_v25).Merge((((_391_v58).contains(_296_v0)) ? ((_391_v58).get(_296_v0)) : (_335_v25))));
        if (_source11.is_DC6) {
          let _392___mcc_h4 = (_source11).cf9;
          let _393___mcc_h5 = (_source11).cf10;
          let _394___mcc_h6 = (_source11).cf11;
          let _395___mcc_h7 = (_source11).cf12;
          let _396_cf12 = _395___mcc_h7;
          let _397_cf11 = _394___mcc_h6;
          let _398_cf10 = _393___mcc_h5;
          let _399_cf9 = _392___mcc_h4;
          let _index47 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          let _rhs58 = new BigNumber((_module.__default.fm5(_dafny.Map.Empty.slice().updateUnsafe(_333_v23,_390_i16), _295_globalState)).cardinality());
          let _rhs59 = (_module.D3.create_DC9(_396_cf12, (_303_v5)[_module.__default.safeIndex((((_316_v13).contains(_396_cf12)) ? ((_316_v13).get(_396_cf12)) : (_390_i16)), new BigNumber((_303_v5).length))])).dtor_cf16;
          let _lhs26 = _318_v15;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          _lhs26[_lhs27] = _rhs58;
          _398_cf10 = _rhs59;
          let _index48 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
          (_297_v1)[_index48] = _module.__default.fm3(_module.__default.safeDivisionInt(_390_i16, _396_cf12), _295_globalState);
          let _400_v59;
          _400_v59 = _dafny.Set.fromElements((_dafny.ZERO).minus(_396_cf12), _module.__default.safeModuloInt(_300_v2, new BigNumber((_module.__default.fm6(_397_cf11, _314_v11, _295_globalState)).length)));
          _400_v59 = _400_v59;
          _334_v24 = _334_v24;
        } else if (_source11.is_DC5) {
          let _401___mcc_h8 = (_source11).cf8;
          let _402_cf8 = _401___mcc_h8;
          let _index49 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          (_318_v15)[_index49] = new BigNumber(147);
          let _403_v60;
          _403_v60 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_297_v1)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length))])).length), _390_i16);
          _316_v13 = (_316_v13).update((((_335_v25).contains(true)) ? ((_335_v25).get(true)) : (new BigNumber((_403_v60).length))), _module.__default.safeDivisionInt(_390_i16, (((_402_cf8).contains(_302_v4)) ? ((_402_cf8).get(_302_v4)) : ((_dafny.ZERO).minus(_300_v2)))));
          let _404_v61;
          let _init6 = ((_405_v5) => function (_406_i17) {
            return _405_v5;
          })(_303_v5);
          let _nw39 = Array((new BigNumber(8)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw39.length); _i0_6++) {
            _nw39[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _404_v61 = _nw39;
          let _407_v62;
          _407_v62 = _dafny.Map.Empty.slice().updateUnsafe((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))],_303_v5);
          let _index50 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_404_v61).length));
          (_404_v61)[_index50] = (((_407_v62).contains(_300_v2)) ? ((_407_v62).get(_300_v2)) : (_303_v5));
          let _index51 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_404_v61).length));
          (_404_v61)[_index51] = _dafny.Seq.Concat(_module.__default.fm7(_302_v4, true, _295_globalState), _303_v5);
          let _408_v63;
          _408_v63 = _dafny.Map.Empty.slice().updateUnsafe(_334_v24,true);
          let _index52 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          let _index53 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
          let _rhs60 = ((_296_v0) ? (new BigNumber((_333_v23).length)) : (new BigNumber((_408_v63).length)));
          let _rhs61 = ((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]).isLessThan((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]);
          let _lhs28 = _318_v15;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          let _lhs30 = _297_v1;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length));
          _lhs28[_lhs29] = _rhs60;
          _lhs30[_lhs31] = _rhs61;
        } else {
          let _409___mcc_h9 = (_source11).cf13;
          let _410_cf13 = _409___mcc_h9;
          let _411_v64;
          _411_v64 = _dafny.Map.Empty.slice().updateUnsafe(_301_v3,_302_v4);
          let _412_v65;
          _412_v65 = _dafny.Map.Empty.slice().updateUnsafe(_411_v64,_390_i16);
          _300_v2 = _module.__default.safeModuloInt((((_412_v65).contains(_411_v64)) ? ((_412_v65).get(_411_v64)) : (_390_i16)), _module.__default.safeModuloInt(_300_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(746)), ((_413_v11) => function (_414_i18) {
            return _413_v11;
          })(_314_v11))).length)));
          let _415_v66;
          let _init7 = ((_416_v23) => function (_417_i19) {
            return (_module.D1.create_DC2(_416_v23)).dtor_cf2;
          })(_333_v23);
          let _nw40 = Array((new BigNumber(26)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw40.length); _i0_7++) {
            _nw40[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _415_v66 = _nw40;
          let _nw41 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _415_v66 = _nw41;
          let _index54 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length));
          (_318_v15)[_index54] = (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))];
          let _418_v67;
          _418_v67 = _dafny.Seq.of(_318_v15, _318_v15, _318_v15, ((_296_v0) ? (_318_v15) : (_318_v15)));
          _418_v67 = _418_v67;
        }
        let _419_v68;
        _419_v68 = _dafny.MultiSet.fromElements((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _300_v2);
        let _420_v69;
        _420_v69 = _dafny.Map.Empty.slice().updateUnsafe(_314_v11,new BigNumber((_dafny.Seq.UnicodeFromString("tmmd")).length));
        let _421_v70;
        _421_v70 = _module.D4.create_DC12(_419_v68);
        let _422_v71;
        _422_v71 = _dafny.Map.Empty.slice().updateUnsafe(_314_v11,(_419_v68).update(_390_i16, _module.__default.abs(_390_i16)));
        let _423_v72;
        _423_v72 = _dafny.Map.Empty.slice().updateUnsafe(_301_v3,_dafny.MultiSet.FromArray(_dafny.Seq.of(_390_i16, (_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))])));
        let _424_v73;
        let _nw42 = Array((new BigNumber(29)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _419_v68;
        _nw42[(_dafny.ONE).toNumber()] = _419_v68;
        _nw42[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_300_v2);
        _nw42[(new BigNumber(3)).toNumber()] = (_419_v68).Intersect(_419_v68);
        _nw42[(new BigNumber(4)).toNumber()] = (_419_v68).Difference(_dafny.MultiSet.FromArray(_334_v24));
        _nw42[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]);
        _nw42[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.update(_334_v24, _module.__default.safeIndex(new BigNumber((_312_v9).cardinality()), new BigNumber((_334_v24).length)), new BigNumber((_334_v24).length)));
        _nw42[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((((_420_v69).contains(_314_v11)) ? ((_420_v69).get(_314_v11)) : (new BigNumber((_333_v23).length))));
        _nw42[(new BigNumber(8)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_334_v24);
        _nw42[(new BigNumber(10)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(11)).toNumber()] = (_421_v70).dtor_cf20;
        _nw42[(new BigNumber(12)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(13)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(14)).toNumber()] = (_419_v68).Intersect((_dafny.MultiSet.FromArray(_334_v24)).update(_390_i16, _module.__default.abs(_390_i16)));
        _nw42[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(((((_305_v7)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_305_v7).length))]).contains(new BigNumber((_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _300_v2)).length))) ? (((_305_v7)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_305_v7).length))]).get(new BigNumber((_dafny.Seq.of((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _300_v2)).length))) : (new BigNumber((_303_v5).length)))))).Difference(_dafny.MultiSet.fromElements((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _390_i16));
        _nw42[(new BigNumber(16)).toNumber()] = _module.__default.fm8((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _333_v23, _295_globalState);
        _nw42[(new BigNumber(17)).toNumber()] = (_419_v68).Intersect(_419_v68);
        _nw42[(new BigNumber(18)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.FromArray(_334_v24);
        _nw42[(new BigNumber(20)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(21)).toNumber()] = _dafny.MultiSet.FromArray(_334_v24);
        _nw42[(new BigNumber(22)).toNumber()] = (_419_v68).Difference(_419_v68);
        _nw42[(new BigNumber(23)).toNumber()] = (((_422_v71).contains(_314_v11)) ? ((_422_v71).get(_314_v11)) : ((_419_v68).update((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))], _module.__default.abs((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]))));
        _nw42[(new BigNumber(24)).toNumber()] = (((_423_v72).contains(true)) ? ((_423_v72).get(true)) : (_419_v68));
        _nw42[(new BigNumber(25)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(26)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_334_v24, _334_v24));
        _nw42[(new BigNumber(27)).toNumber()] = _419_v68;
        _nw42[(new BigNumber(28)).toNumber()] = _419_v68;
        _424_v73 = _nw42;
        _424_v73 = _424_v73;
        let _425_v74;
        _425_v74 = _dafny.Map.Empty.slice().updateUnsafe(_296_v0,((_318_v15)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_318_v15).length))]).isEqualTo(_module.__default.fm0(_295_globalState)));
        let _rhs62 = (_297_v1)[_module.__default.safeIndex(new BigNumber(709), new BigNumber((_297_v1).length))];
        let _rhs63 = _module.__default.fm9(_295_globalState);
        _296_v0 = _rhs62;
        _425_v74 = _rhs63;
      }
      process.stdout.write(_dafny.toString((_295_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_296_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_300_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_301_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_302_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_303_v5, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_304_v6, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-554),new BigNumber(-554)).updateUnsafe(_dafny.ZERO,_dafny.ZERO), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-554),new BigNumber(-554)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-554),new BigNumber(-554)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-554),new BigNumber(-554)), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.ZERO)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_305_v7)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(734)).updateUnsafe(new BigNumber(503),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_305_v7)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ZERO).updateUnsafe(new BigNumber(9),_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_311_v8, _dafny.Seq.of(_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v9).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_313_v10, _dafny.Seq.of(_dafny.MultiSet.fromElements(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_314_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_315_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_316_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_317_v14).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_318_v15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_333_v23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_334_v24, _dafny.Seq.of(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_335_v25).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_336_v26, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_337_v27).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_338_v28).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Seq.UnicodeFromString("xyfj")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_339_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_368_v47).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_368_v47).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ZERO]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ZERO]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ZERO]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ONE]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ONE]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[_dafny.ONE]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(2)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(2)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(2)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(3)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(3)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(3)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(4)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(4)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(4)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(5)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(5)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(5)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(6)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(6)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(6)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(7)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(7)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(7)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(8)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(8)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(8)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(9)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(9)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(9)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(10)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(10)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(10)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(11)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(11)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(11)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(12)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(12)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(12)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(13)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(13)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(13)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(14)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(14)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(14)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(15)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(15)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(15)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(16)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(16)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(16)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(17)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(17)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(17)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(18)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(18)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(18)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(19)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(19)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(19)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(20)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(20)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(20)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(21)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(21)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(21)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(22)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(22)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(22)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(23)]).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(23)]).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_369_v48)[new BigNumber(23)]).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_373_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC3(cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf6, cf7) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + this.cf2.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6 && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf9, cf10, cf11, cf12) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D2(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(new _dafny.CodePoint('D'.codePointAt(0)), false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC10(cf17, cf18) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D3(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(false, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf24, cf25, cf26) {
      let $dt = new D5(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D5(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && this.cf26 === other.cf26;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.ZERO, false, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf28) {
      let $dt = new D6(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC18(cf29, cf30, cf31, cf32, cf33) {
      let $dt = new D6(1);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D6(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC16(cf27) {
      let $dt = new D6(3);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC20(cf35) {
      let $dt = new D6(4);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get is_DC20() { return this.$tag === 4; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf36) {
      let $dt = new D7(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf38, cf39, cf40) {
      let $dt = new D8(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC24(cf41, cf42) {
      let $dt = new D8(1);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC22(cf37) {
      let $dt = new D8(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf44) {
      let $dt = new D9(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC27(cf45) {
      let $dt = new D9(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC25(cf43) {
      let $dt = new D9(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf43 === other.cf43;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf47, cf48) {
      let $dt = new D10(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC28(cf46) {
      let $dt = new D10(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf50) {
      let $dt = new D11(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC30(cf49) {
      let $dt = new D11(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf51) {
      let $dt = new D12(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf53, cf54, cf55, cf56) {
      let $dt = new D13(0);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC35(cf57, cf58, cf59, cf60, cf61) {
      let $dt = new D13(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC33(cf52) {
      let $dt = new D13(2);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC36(cf62) {
      let $dt = new D13(3);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get is_DC36() { return this.$tag === 3; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + this.cf56.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58) && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60) && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf52 === other.cf52;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC34(false, false, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf64) {
      let $dt = new D14(0);
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC39(cf65, cf66, cf67) {
      let $dt = new D14(1);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC37(cf63) {
      let $dt = new D14(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC38" + "(" + this.cf64.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf65, other.cf65) && _dafny.areEqual(this.cf66, other.cf66) && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf63 === other.cf63;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC38(_dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC40(cf68) {
      let $dt = new D15(0);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf70, cf71, cf72) {
      let $dt = new D16(0);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC41(cf69) {
      let $dt = new D16(1);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf70 === other.cf70 && this.cf71 === other.cf71 && this.cf72 === other.cf72;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf69, other.cf69);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC42(false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC43(cf73) {
      let $dt = new D17(0);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC45(cf75, cf76, cf77) {
      let $dt = new D18(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC46() {
      let $dt = new D18(1);
      return $dt;
    }
    static create_DC44(cf74) {
      let $dt = new D18(2);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC47(cf78) {
      let $dt = new D18(3);
      $dt.cf78 = cf78;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf78() { return this.cf78; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC45" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC46";
      } else if (this.$tag === 2) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf78) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf78, other.cf78);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC45(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49(cf80, cf81) {
      let $dt = new D19(0);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC50(cf82) {
      let $dt = new D19(1);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC48(cf79) {
      let $dt = new D19(2);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC49" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC50" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC49(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC51(cf83) {
      let $dt = new D20(1);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC52";
      } else if (this.$tag === 1) {
        return "D20.DC51" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC52();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf84) {
      let $dt = new D21(0);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf86, cf87, cf88) {
      let $dt = new D22(0);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC54(cf85) {
      let $dt = new D22(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC56(cf89) {
      let $dt = new D22(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC56() { return this.$tag === 2; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC55" + "(" + this.cf86.toVerbatimString(true) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC54" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf86, other.cf86) && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf85 === other.cf85;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC55(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58() {
      let $dt = new D23(0);
      return $dt;
    }
    static create_DC57(cf90) {
      let $dt = new D23(1);
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC59(cf91) {
      let $dt = new D23(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC59() { return this.$tag === 2; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC58";
      } else if (this.$tag === 1) {
        return "D23.DC57" + "(" + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC59" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf90 === other.cf90;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf91, other.cf91);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC58();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC60(cf92) {
      let $dt = new D24(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC60" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62(cf94, cf95, cf96) {
      let $dt = new D25(0);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC63(cf97, cf98, cf99, cf100) {
      let $dt = new D25(1);
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC61(cf93) {
      let $dt = new D25(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC64(cf101) {
      let $dt = new D25(3);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get is_DC64() { return this.$tag === 3; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC62" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC63" + "(" + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ", " + _dafny.toString(this.cf99) + ", " + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC61" + "(" + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 3) {
        return "D25.DC64" + "(" + _dafny.toString(this.cf101) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf94 === other.cf94 && _dafny.areEqual(this.cf95, other.cf95) && this.cf96 === other.cf96;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf97, other.cf97) && this.cf98 === other.cf98 && this.cf99 === other.cf99 && this.cf100 === other.cf100;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf93 === other.cf93;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf101, other.cf101);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC62(null, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC66(cf103, cf104, cf105, cf106, cf107) {
      let $dt = new D26(0);
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC65(cf102) {
      let $dt = new D26(1);
      $dt.cf102 = cf102;
      return $dt;
    }
    static create_DC67(cf108) {
      let $dt = new D26(2);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC66" + "(" + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC65" + "(" + _dafny.toString(this.cf102) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC67" + "(" + _dafny.toString(this.cf108) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf103, other.cf103) && this.cf104 === other.cf104 && _dafny.areEqual(this.cf105, other.cf105) && this.cf106 === other.cf106 && _dafny.areEqual(this.cf107, other.cf107);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf102, other.cf102);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC66(_dafny.Map.Empty, false, _dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC69(cf110, cf111) {
      let $dt = new D27(0);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      return $dt;
    }
    static create_DC68(cf109) {
      let $dt = new D27(1);
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC70(cf112) {
      let $dt = new D27(2);
      $dt.cf112 = cf112;
      return $dt;
    }
    get is_DC69() { return this.$tag === 0; }
    get is_DC68() { return this.$tag === 1; }
    get is_DC70() { return this.$tag === 2; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf112() { return this.cf112; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC69" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ")";
      } else if (this.$tag === 1) {
        return "D27.DC68" + "(" + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 2) {
        return "D27.DC70" + "(" + _dafny.toString(this.cf112) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf110, other.cf110) && _dafny.areEqual(this.cf111, other.cf111);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf109, other.cf109);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf112, other.cf112);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D27.create_DC69(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this._f0 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f14 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14) {
      let _this = this;
      (_this)._f14 = f14;
      return;
    }
    fm29(p0, p1, p2, globalState) {
      let _this = this;
      let _source12 = _module.D1.create_DC3(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(false, false), _dafny.MultiSet.fromElements(true))).length),_dafny.Seq.of(false))).length), new BigNumber(337));
      if (_source12.is_DC3) {
        let _426___mcc_h0 = (_source12).cf3;
        let _427___mcc_h1 = (_source12).cf4;
        let _428___mcc_h2 = (_source12).cf5;
        let _429_cf5 = _428___mcc_h2;
        let _430_cf4 = _427___mcc_h1;
        let _431_cf3 = _426___mcc_h0;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else if (_source12.is_DC4) {
        let _432___mcc_h3 = (_source12).cf6;
        let _433___mcc_h4 = (_source12).cf7;
        let _434_cf7 = _433___mcc_h4;
        let _435_cf6 = _432___mcc_h3;
        return new _dafny.CodePoint('f'.codePointAt(0));
      } else {
        let _436___mcc_h5 = (_source12).cf2;
        let _437_cf2 = _436___mcc_h5;
        return (_module.D6.create_DC18(new BigNumber((function () {
  let _coll30 = new _dafny.Set();
  for (const _compr_30 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(201),false)).Keys.Elements) {
    let _438_v0 = _compr_30;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(201),false)).contains(_438_v0)) {
      _coll30.add((_438_v0).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(723)), function (_439_i0) {
        return new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), function (_440_i1) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        })).length),true)).length))).length);
      })).length)));
    }
  }
  return _coll30;
}()).length), new BigNumber(226), false, new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber(571))).dtor_cf32;
      }
    };
    fm30(globalState) {
      let _this = this;
      return !(true) || (!(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(238), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(673),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))).cardinality())).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-31),new BigNumber(-428))).length))));
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm28(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),false)).length))).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-498)),!(!(false)))).length));
    };
    m12(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = _module.D5.Default();
      let _441_v0;
      let _init8 = ((_442_p0) => function (_443_i0) {
        return _442_p0;
      })(p0);
      let _nw43 = Array((new BigNumber(3)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw43.length); _i0_8++) {
        _nw43[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _441_v0 = _nw43;
      let _444_v1;
      _444_v1 = _dafny.Map.Empty.slice().updateUnsafe(_441_v0,p0);
      let _445_v2;
      let _446_v3;
      let _447_v4;
      let _448_v5;
      let _out13;
      let _out14;
      let _out15;
      let _out16;
      let _outcollector3 = _module.__default.m0(_444_v1, globalState);
      _out13 = _outcollector3[0];
      _out14 = _outcollector3[1];
      _out15 = _outcollector3[2];
      _out16 = _outcollector3[3];
      _445_v2 = _out13;
      _446_v3 = _out14;
      _447_v4 = _out15;
      _448_v5 = _out16;
      let _449_i1;
      _449_i1 = _dafny.ZERO;
      L3: {
        while (p0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_449_i1)) {
              break L3;
            }
            _449_i1 = (_449_i1).plus(_dafny.ONE);
            let _450_v6;
            _450_v6 = new BigNumber(534);
            r0 = _450_v6;
            let _451_v7;
            _451_v7 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("odqlsiwfo")).length));
            let _452_v8;
            _452_v8 = _dafny.MultiSet.fromElements(p0);
            let _453_v9;
            _453_v9 = _dafny.Seq.of(_452_v8);
            let _454_v10;
            _454_v10 = new _dafny.CodePoint('g'.codePointAt(0));
            let _455_v11;
            _455_v11 = _dafny.MultiSet.fromElements(_454_v10);
            let _456_v13;
            let _nw44 = Array((new BigNumber(21)).toNumber());
            _nw44[(_dafny.ZERO).toNumber()] = new BigNumber((_451_v7).length);
            _nw44[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_450_v6);
            _nw44[(new BigNumber(2)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(3)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(4)).toNumber()] = new BigNumber(((_453_v9)[_module.__default.safeIndex(_450_v6, new BigNumber((_453_v9).length))]).cardinality());
            _nw44[(new BigNumber(5)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(6)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(7)).toNumber()] = new BigNumber(693);
            _nw44[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_451_v7)[_module.__default.safeIndex(_450_v6, new BigNumber((_451_v7).length))]);
            _nw44[(new BigNumber(9)).toNumber()] = new BigNumber((_452_v8).cardinality());
            _nw44[(new BigNumber(10)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(11)).toNumber()] = (((_455_v11).contains(_454_v10)) ? ((_455_v11).get(_454_v10)) : (new BigNumber((function () {
              let _coll31 = new _dafny.Set();
              for (const _compr_31 of _dafny.IntegerRange(new BigNumber(666), new BigNumber(789))) {
                let _457_v12 = _compr_31;
                if (((new BigNumber(666)).isLessThanOrEqualTo(_457_v12)) && ((_457_v12).isLessThan(new BigNumber(789)))) {
                  _coll31.add(_module.__default.safeModuloInt(_457_v12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("spwpy"))).length)));
                }
              }
              return _coll31;
            }()).length)));
            _nw44[(new BigNumber(12)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(13)).toNumber()] = new BigNumber(347);
            _nw44[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_450_v6);
            _nw44[(new BigNumber(15)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(16)).toNumber()] = new BigNumber(367);
            _nw44[(new BigNumber(17)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ledgplr")).length);
            _nw44[(new BigNumber(19)).toNumber()] = _450_v6;
            _nw44[(new BigNumber(20)).toNumber()] = new BigNumber((_452_v8).cardinality());
            _456_v13 = _nw44;
            let _458_v14;
            let _nw45 = new _module.C0();
            _nw45.__ctor(_456_v13);
            _458_v14 = _nw45;
            let _459_v15;
            _459_v15 = _dafny.MultiSet.fromElements(_450_v6, _450_v6);
            let _460_v16;
            _460_v16 = _dafny.Set.fromElements(_459_v15);
            if ((new BigNumber(938)).isLessThanOrEqualTo((new BigNumber((_460_v16).length)).minus(_module.__default.fm0(globalState)))) {
              let _461_v17;
              let _nw46 = new _module.C0();
              _nw46.__ctor((_458_v14).f14);
              _461_v17 = _nw46;
              let _462_v18;
              _462_v18 = _dafny.Seq.UnicodeFromString("hv");
              let _index55 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_441_v0).length));
              (_441_v0)[_index55] = _dafny.Seq.contains(_462_v18, _454_v10);
              let _index56 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_441_v0).length));
              (_441_v0)[_index56] = _446_v3;
              let _index57 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_441_v0).length));
              (_441_v0)[_index57] = true;
              _454_v10 = _454_v10;
              let _463_v19;
              let _nw47 = new _module.C0();
              _nw47.__ctor((_461_v17).f14);
              _463_v19 = _nw47;
            } else {
              _450_v6 = (_450_v6).multipliedBy(_module.__default.safeModuloInt(_450_v6, _450_v6));
              _451_v7 = _451_v7;
              let _464_v20;
              let _nw48 = new _module.C0();
              _nw48.__ctor((_458_v14).f14);
              _464_v20 = _nw48;
              let _465_v21;
              _465_v21 = _dafny.Seq.UnicodeFromString("dxrlpmtg");
              let _466_v22;
              _466_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_465_v21);
              _465_v21 = _dafny.Seq.Concat((((_466_v22).contains(_446_v3)) ? ((_466_v22).get(_446_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), function (_467_i2) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), function (_468_i3) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              }));
              let _469_v23;
              _469_v23 = _dafny.Map.Empty.slice().updateUnsafe(_446_v3,_451_v7);
              let _470_v24;
              let _nw49 = Array((new BigNumber(13)).toNumber());
              _nw49[(_dafny.ZERO).toNumber()] = _451_v7;
              _nw49[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_451_v7, _module.__default.fm31(_450_v6, _445_v2, new BigNumber(917), _454_v10, globalState));
              _nw49[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(new BigNumber(760));
              _nw49[(new BigNumber(3)).toNumber()] = (((_469_v23).contains(_445_v2)) ? ((_469_v23).get(_445_v2)) : (_451_v7));
              _nw49[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_450_v6);
              _nw49[(new BigNumber(5)).toNumber()] = _451_v7;
              _nw49[(new BigNumber(6)).toNumber()] = _451_v7;
              _nw49[(new BigNumber(7)).toNumber()] = (((_469_v23).contains(_445_v2)) ? ((_469_v23).get(_445_v2)) : (_dafny.Seq.update(_451_v7, _module.__default.safeIndex(_450_v6, new BigNumber((_451_v7).length)), (_dafny.ZERO).minus(_450_v6))));
              _nw49[(new BigNumber(8)).toNumber()] = _451_v7;
              _nw49[(new BigNumber(9)).toNumber()] = _451_v7;
              _nw49[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(new BigNumber(778), _450_v6);
              _nw49[(new BigNumber(11)).toNumber()] = _451_v7;
              _nw49[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-269)), ((_471_v6) => function (_472_i4) {
                return _471_v6;
              })(_450_v6));
              _470_v24 = _nw49;
              let _index58 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_470_v24).length));
              (_470_v24)[_index58] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), ((_473_p0, _474_v3) => function (_475_i5) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_473_p0,_474_v3)).length);
              })(p0, _446_v3));
              let _index59 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_470_v24).length));
              (_470_v24)[_index59] = _451_v7;
            }
            if (_446_v3) {
              _446_v3 = !(_446_v3);
              _456_v13 = (_458_v14).f14;
              _441_v0 = _441_v0;
              r0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_450_v6, ((_446_v3) ? (_450_v6) : (_module.__default.fm0(globalState)))));
              let _476_v25;
              _476_v25 = _dafny.Seq.UnicodeFromString("mqxkm");
              _459_v15 = (_module.__default.fm8(new BigNumber(904), _476_v25, globalState)).update(new BigNumber((_476_v25).length), _module.__default.abs(_450_v6));
            } else {
              let _477_v26;
              _477_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              _477_v26 = _module.__default.fm9(globalState);
              _450_v6 = (_450_v6).multipliedBy(_450_v6);
              let _478_v27;
              _478_v27 = _dafny.Seq.UnicodeFromString("x");
              let _479_v28;
              _479_v28 = _dafny.Map.Empty.slice().updateUnsafe(_441_v0,new BigNumber((_dafny.Seq.Concat(_module.__default.fm2(false, globalState), _478_v27)).length));
              let _480_v29;
              _480_v29 = _module.D1.create_DC2(_478_v27);
              let _rhs64 = _479_v28;
              let _rhs65 = ((_dafny.ZERO).minus(_450_v6)).isLessThan(new BigNumber(((_453_v9)[_module.__default.safeIndex(_450_v6, new BigNumber((_453_v9).length))]).cardinality()));
              let _rhs66 = _dafny.Seq.Concat(_dafny.Seq.Concat(_478_v27, (_480_v29).dtor_cf2), _478_v27);
              _479_v28 = _rhs64;
              r2 = _rhs65;
              _478_v27 = _rhs66;
              let _481_v30;
              _481_v30 = _dafny.Seq.of(_478_v27, _478_v27);
              let _482_v31;
              _482_v31 = _dafny.Map.Empty.slice().updateUnsafe(_450_v6,_module.__default.fm2(_446_v3, globalState));
              let _483_v32;
              _483_v32 = _dafny.Map.Empty.slice().updateUnsafe(_458_v14,_dafny.Seq.Create(_module.__default.abs(new BigNumber(180)), ((_484_v10) => function (_485_i7) {
                return _484_v10;
              })(_454_v10)));
              let _486_v33;
              let _nw50 = Array((new BigNumber(24)).toNumber());
              _nw50[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("bsprls");
              _nw50[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_478_v27, (_481_v30)[_module.__default.safeIndex((_dafny.ZERO).minus(_450_v6), new BigNumber((_481_v30).length))]);
              _nw50[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fviebtams"), _478_v27);
              _nw50[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_478_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-198)), ((_487_v10) => function (_488_i6) {
                return _487_v10;
              })(_454_v10)));
              _nw50[(new BigNumber(4)).toNumber()] = (((_482_v31).contains((_dafny.ZERO).minus(_450_v6))) ? ((_482_v31).get((_dafny.ZERO).minus(_450_v6))) : (_dafny.Seq.UnicodeFromString("bqw")));
              _nw50[(new BigNumber(5)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(6)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(7)).toNumber()] = (((_483_v32).contains(_458_v14)) ? ((_483_v32).get(_458_v14)) : (_478_v27));
              _nw50[(new BigNumber(8)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(9)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(10)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(11)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(12)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), function (_489_i8) {
                return new _dafny.CodePoint('h'.codePointAt(0));
              });
              _nw50[(new BigNumber(14)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nfx"), _478_v27);
              _nw50[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(137)), function (_490_i9) {
                return new _dafny.CodePoint('f'.codePointAt(0));
              });
              _nw50[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_478_v27, _478_v27);
              _nw50[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("obxuhqq");
              _nw50[(new BigNumber(19)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(20)).toNumber()] = _478_v27;
              _nw50[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_478_v27, _dafny.Seq.update((((_482_v31).contains(_450_v6)) ? ((_482_v31).get(_450_v6)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), ((_491_v10) => function (_492_i10) {
                return _491_v10;
              })(_454_v10)))), _module.__default.safeIndex(_450_v6, new BigNumber(((((_482_v31).contains(_450_v6)) ? ((_482_v31).get(_450_v6)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), ((_493_v10) => function (_494_i10) {
                return _493_v10;
              })(_454_v10))))).length)), _454_v10));
              _nw50[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm2(_446_v3, globalState), _478_v27);
              _nw50[(new BigNumber(23)).toNumber()] = _478_v27;
              _486_v33 = _nw50;
              let _index60 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_486_v33).length));
              (_486_v33)[_index60] = _478_v27;
              let _495_v34;
              _495_v34 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(true), _452_v8);
              let _index61 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_486_v33).length));
              (_486_v33)[_index61] = _dafny.Seq.update(_478_v27, _module.__default.safeIndex(new BigNumber((_495_v34).length), new BigNumber((_478_v27).length)), new _dafny.CodePoint('s'.codePointAt(0)));
              let _496_v35;
              _496_v35 = _dafny.Map.Empty.slice().updateUnsafe((_486_v33)[_module.__default.safeIndex(new BigNumber(904), new BigNumber((_486_v33).length))],_450_v6);
              let _497_v36;
              _497_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_446_v3) ? (_445_v2) : (true)),_496_v35);
              _497_v36 = (_497_v36).update(!(new BigNumber(568)).isEqualTo(new BigNumber((_478_v27).length)), _496_v35);
            }
          }
        }
      }
      let _498_v37;
      _498_v37 = _dafny.Seq.UnicodeFromString("ipd");
      let _499_v38;
      _499_v38 = new BigNumber(611);
      let _500_v39;
      _500_v39 = new _dafny.CodePoint('i'.codePointAt(0));
      _498_v37 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm2(_445_v2, globalState), _498_v37), _dafny.Seq.Concat(_498_v37, _498_v37)), _module.__default.safeIndex(_499_v38, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm2(_445_v2, globalState), _498_v37), _dafny.Seq.Concat(_498_v37, _498_v37))).length)), ((p0) ? (_500_v39) : (_500_v39)));
      let _501_i11;
      _501_i11 = _dafny.ZERO;
      L4: {
        while (p0) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_501_i11)) {
              break L4;
            }
            _501_i11 = (_501_i11).plus(_dafny.ONE);
            let _502_v40;
            _502_v40 = _dafny.Seq.of(_499_v38);
            let _rhs67 = _502_v40;
            let _rhs68 = ((false) ? ((_499_v38).minus(_499_v38)) : (_499_v38));
            _502_v40 = _rhs67;
            _499_v38 = _rhs68;
            let _503_v41;
            _503_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,_498_v37);
            _503_v41 = (_503_v41).update((_447_v4)[_module.__default.safeIndex(_499_v38, new BigNumber((_447_v4).length))], _498_v37);
            let _504_v42;
            _504_v42 = _dafny.MultiSet.fromElements(_441_v0, _441_v0);
            let _rhs69 = _502_v40;
            let _rhs70 = new BigNumber((((_dafny.MultiSet.fromElements(_441_v0, _441_v0, _441_v0, _441_v0)).Difference(_504_v42)).Difference((_504_v42).Union(_504_v42))).cardinality());
            _502_v40 = _rhs69;
            _499_v38 = _rhs70;
            let _505_v43;
            let _init9 = ((_506_v38) => function (_507_i12) {
              return (_507_i12).plus(_506_v38);
            })(_499_v38);
            let _nw51 = Array((new BigNumber(8)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw51.length); _i0_9++) {
              _nw51[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _505_v43 = _nw51;
            let _508_v44;
            let _nw52 = new _module.C0();
            _nw52.__ctor(_505_v43);
            _508_v44 = _nw52;
          }
        }
      }
      let _509_v45;
      _509_v45 = _dafny.Map.Empty.slice().updateUnsafe(_499_v38,new BigNumber(368));
      let _hi2 = new BigNumber((_509_v45).length);
      for (let _510_i13 = (_dafny.ZERO).minus(((false) ? (_499_v38) : (_499_v38))); _510_i13.isLessThan(_hi2); _510_i13 = _510_i13.plus(_dafny.ONE)) {
        let _511_v46;
        _511_v46 = _dafny.Set.fromElements(_441_v0);
        let _512_v47;
        let _nw53 = Array((new BigNumber(7)).toNumber());
        _nw53[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_511_v46).length));
        _nw53[(_dafny.ONE).toNumber()] = new BigNumber(326);
        _nw53[(new BigNumber(2)).toNumber()] = _module.__default.fm0(globalState);
        _nw53[(new BigNumber(3)).toNumber()] = _499_v38;
        _nw53[(new BigNumber(4)).toNumber()] = _499_v38;
        _nw53[(new BigNumber(5)).toNumber()] = _499_v38;
        _nw53[(new BigNumber(6)).toNumber()] = new BigNumber(-701);
        _512_v47 = _nw53;
        let _513_v48;
        let _nw54 = new _module.C0();
        _nw54.__ctor(_512_v47);
        _513_v48 = _nw54;
        let _514_v49;
        let _nw55 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _514_v49 = _nw55;
        _514_v49 = _514_v49;
        let _515_v50;
        _515_v50 = _dafny.MultiSet.fromElements(_500_v39);
        let _516_v51;
        _516_v51 = _dafny.Map.Empty.slice().updateUnsafe(_446_v3,new BigNumber((_515_v50).cardinality()));
        _516_v51 = (_516_v51).update(!(_445_v2), _499_v38);
        if (_445_v2) {
          let _517_v52;
          _517_v52 = _dafny.Map.Empty.slice().updateUnsafe(_510_i13,_500_v39);
          let _518_v53;
          _518_v53 = _module.D3.create_DC9((_dafny.ZERO).minus(new BigNumber((_517_v52).length)), _446_v3);
          let _pat_let_tv16 = _498_v37;
          let _pat_let_tv17 = globalState;
          _445_v2 = ((_module.__default.fm32(_510_i13, new BigNumber(420), _499_v38, _498_v37, globalState)).update(function (_pat_let12_0) {
            return function (_519_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_520_dt__update_hcf16_h0) {
                  return _module.D3.create_DC9((_519_dt__update__tmp_h0).dtor_cf15, _520_dt__update_hcf16_h0);
                }(_pat_let13_0);
              }(_module.__default.fm3(new BigNumber((_pat_let_tv16).length), _pat_let_tv17));
            }(_pat_let12_0);
          }(_518_v53), _510_i13)).contains(_518_v53);
          let _521_v54;
          _521_v54 = _dafny.Map.Empty.slice().updateUnsafe(_510_i13,p0);
          let _522_v55;
          _522_v55 = _module.D1.create_DC4((((_521_v54).contains(_499_v38)) ? ((_521_v54).get(_499_v38)) : (_module.__default.fm3(_499_v38, globalState))), _446_v3);
          _522_v55 = _522_v55;
          _499_v38 = _499_v38;
          r2 = _445_v2;
          let _523_v56;
          let _nw56 = Array((new BigNumber(20)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _441_v0;
          _nw56[(_dafny.ONE).toNumber()] = _441_v0;
          _nw56[(new BigNumber(2)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(3)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(4)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(5)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(6)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(7)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(8)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(9)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(10)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(11)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(12)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(13)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(14)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(15)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(16)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(17)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(18)).toNumber()] = _441_v0;
          _nw56[(new BigNumber(19)).toNumber()] = _441_v0;
          _523_v56 = _nw56;
          _523_v56 = _523_v56;
        } else {
          let _init10 = ((_524_v39) => function (_525_i14) {
            return _524_v39;
          })(_500_v39);
          let _nw57 = Array((new BigNumber(11)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw57.length); _i0_10++) {
            _nw57[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _514_v49 = _nw57;
          let _526_v57;
          _526_v57 = _dafny.Seq.of(new BigNumber((_498_v37).length));
          let _index62 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_513_v48).f14).length));
          ((_513_v48).f14)[_index62] = (_this).fm28(p0, new BigNumber((_526_v57).length), globalState);
          let _527_v58;
          _527_v58 = _module.D6.create_DC17(p0);
          let _index63 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_513_v48).f14).length));
          let _rhs71 = _509_v45;
          let _rhs72 = (new BigNumber(-354)).minus(new BigNumber((_dafny.Seq.Concat(_498_v37, _498_v37)).length));
          let _rhs73 = _446_v3;
          let _rhs74 = _527_v58;
          let _lhs32 = (_513_v48).f14;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(374), new BigNumber(((_513_v48).f14).length));
          _509_v45 = _rhs71;
          _lhs32[_lhs33] = _rhs72;
          _446_v3 = _rhs73;
          _527_v58 = _rhs74;
          _445_v2 = _446_v3;
          let _528_v59;
          _528_v59 = _dafny.Set.fromElements(function (_pat_let14_0) {
            return function (_529_dt__update__tmp_h1) {
              return function (_pat_let15_0) {
                return function (_530_dt__update_hcf28_h0) {
                  return _module.D6.create_DC17(_530_dt__update_hcf28_h0);
                }(_pat_let15_0);
              }(false);
            }(_pat_let14_0);
          }(_527_v58), _527_v58, _527_v58, _527_v58);
          _528_v59 = ((_528_v59).Difference(_528_v59)).Difference(_528_v59);
          r0 = _499_v38;
        }
      }
      let _531_v60;
      _531_v60 = _dafny.MultiSet.fromElements(!(_446_v3), _446_v3, p0);
      let _hi3 = (((_531_v60).contains(_445_v2)) ? ((_531_v60).get(_445_v2)) : (_499_v38));
      for (let _532_i15 = _499_v38; _532_i15.isLessThan(_hi3); _532_i15 = _532_i15.plus(_dafny.ONE)) {
        r0 = _532_i15;
        let _533_v61;
        _533_v61 = _module.D0.create_DC1(new BigNumber(334));
        let _source13 = _533_v61;
        if (_source13.is_DC1) {
          let _534___mcc_h0 = (_source13).cf1;
          let _535_cf1 = _534___mcc_h0;
          r0 = _module.__default.safeModuloInt(_499_v38, (_532_i15).plus(_535_cf1));
          let _536_v62;
          let _nw58 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _536_v62 = _nw58;
          let _537_v63;
          _537_v63 = _dafny.Map.Empty.slice().updateUnsafe(((_445_v2) ? (_536_v62) : (_536_v62)),_445_v2);
          _537_v63 = (_537_v63).update(_536_v62, false);
          _500_v39 = _500_v39;
          let _538_v64;
          _538_v64 = _dafny.Set.fromElements(_499_v38, _499_v38, _532_i15, _532_i15);
          let _539_v65;
          _539_v65 = _dafny.Seq.of(_499_v38);
          let _540_v66;
          _540_v66 = _module.D3.create_DC9(_499_v38, _446_v3);
          let _541_v67;
          _541_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_539_v65).length),_540_v66);
          let _542_v69;
          _542_v69 = _dafny.Seq.of(_538_v64, function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (_541_v67).Keys.Elements) {
              let _543_v68 = _compr_32;
              if ((_541_v67).contains(_543_v68)) {
                _coll32.add((_543_v68).minus(new BigNumber(96)));
              }
            }
            return _coll32;
          }(), _538_v64);
          let _544_v70;
          _544_v70 = _dafny.Set.fromElements(_538_v64, _538_v64, _module.__default.fm33(new BigNumber((_dafny.Seq.of(_446_v3, !(_445_v2))).length), globalState));
          let _545_v71;
          _545_v71 = _dafny.Set.fromElements(new BigNumber(((_542_v69)[_module.__default.safeIndex(_499_v38, new BigNumber((_542_v69).length))]).length), new BigNumber((_544_v70).length));
          r2 = (_545_v71).contains(_499_v38);
        } else {
          let _546___mcc_h1 = (_source13).cf0;
          let _547_cf0 = _546___mcc_h1;
          let _548_v72;
          _548_v72 = _dafny.Map.Empty.slice().updateUnsafe(_445_v2,p0);
          _548_v72 = (_548_v72).update(p0, !(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_446_v3, _446_v3), _447_v4))));
          let _549_v73;
          _549_v73 = _dafny.Seq.of(_498_v37, _498_v37, _498_v37, _498_v37, _498_v37);
          let _550_v74;
          _550_v74 = _dafny.Map.Empty.slice().updateUnsafe((_447_v4)[_module.__default.safeIndex(_547_cf0, new BigNumber((_447_v4).length))],_441_v0);
          let _551_v75;
          let _nw59 = Array((new BigNumber(24)).toNumber());
          _nw59[(_dafny.ZERO).toNumber()] = _547_cf0;
          _nw59[(_dafny.ONE).toNumber()] = _499_v38;
          _nw59[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_532_i15, new BigNumber((_549_v73).length));
          _nw59[(new BigNumber(3)).toNumber()] = (new BigNumber((_dafny.Seq.UnicodeFromString("x")).length)).plus(_547_cf0);
          _nw59[(new BigNumber(4)).toNumber()] = (_532_i15).multipliedBy(_532_i15);
          _nw59[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(988)), ((_552_v39) => function (_553_i16) {
            return _552_v39;
          })(_500_v39)), _498_v37)).length);
          _nw59[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm4(_499_v38, _446_v3, globalState)).length);
          _nw59[(new BigNumber(7)).toNumber()] = new BigNumber((_498_v37).length);
          _nw59[(new BigNumber(8)).toNumber()] = _499_v38;
          _nw59[(new BigNumber(9)).toNumber()] = (_547_cf0).minus(new BigNumber((_447_v4).length));
          _nw59[(new BigNumber(10)).toNumber()] = _532_i15;
          _nw59[(new BigNumber(11)).toNumber()] = _532_i15;
          _nw59[(new BigNumber(12)).toNumber()] = ((_445_v2) ? (_499_v38) : (_547_cf0));
          _nw59[(new BigNumber(13)).toNumber()] = _499_v38;
          _nw59[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_532_i15);
          _nw59[(new BigNumber(15)).toNumber()] = _499_v38;
          _nw59[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_548_v72).length), _532_i15);
          _nw59[(new BigNumber(17)).toNumber()] = new BigNumber((_550_v74).length);
          _nw59[(new BigNumber(18)).toNumber()] = _532_i15;
          _nw59[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_498_v37).length)), _499_v38)).cardinality());
          _nw59[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_532_i15);
          _nw59[(new BigNumber(21)).toNumber()] = _499_v38;
          _nw59[(new BigNumber(22)).toNumber()] = _499_v38;
          _nw59[(new BigNumber(23)).toNumber()] = _547_cf0;
          _551_v75 = _nw59;
          let _index64 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_551_v75).length));
          (_551_v75)[_index64] = _532_i15;
          let _index65 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_551_v75).length));
          (_551_v75)[_index65] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_498_v37, _498_v37), _dafny.Seq.UnicodeFromString("sp")), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_498_v37, _498_v37), _dafny.Seq.UnicodeFromString("sp"))).length)), _500_v39)).length);
          let _554_v76;
          _554_v76 = _dafny.MultiSet.fromElements((_this).fm28(_445_v2, new BigNumber(-524), globalState));
          let _555_v77;
          _555_v77 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_500_v39,(((_554_v76).contains(_499_v38)) ? ((_554_v76).get(_499_v38)) : (_547_cf0)))).length));
          let _556_v78;
          _556_v78 = _dafny.MultiSet.fromElements(_547_cf0, new BigNumber((_555_v77).length), _532_i15);
          r2 = ((_556_v78).Union(_554_v76)).IsProperSubsetOf(_dafny.MultiSet.fromElements((_551_v75)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_551_v75).length))]));
          let _557_v79;
          _557_v79 = _dafny.Set.fromElements(_445_v2);
          let _rhs75 = _532_i15;
          let _rhs76 = _module.__default.fm3((_this).fm28(p0, new BigNumber((_dafny.Seq.UnicodeFromString("ihpnbc")).length), globalState), globalState);
          let _rhs77 = !(((_446_v3) ? (_557_v79) : (_dafny.Set.fromElements(_446_v3, !(false))))).equals((_dafny.Set.fromElements(_module.__default.fm3((_551_v75)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_551_v75).length))], globalState))).Difference(_557_v79));
          r0 = _rhs75;
          _445_v2 = _rhs76;
          _445_v2 = _rhs77;
        }
        let _558_v80;
        let _nw60 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _558_v80 = _nw60;
        let _index66 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_558_v80).length));
        (_558_v80)[_index66] = (new BigNumber((_498_v37).length)).minus(_499_v38);
        let _index67 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_558_v80).length));
        (_558_v80)[_index67] = _module.__default.safeDivisionInt((((_531_v60).contains(_446_v3)) ? ((_531_v60).get(_446_v3)) : (_499_v38)), _499_v38);
        if (false) {
          let _index68 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_441_v0).length));
          (_441_v0)[_index68] = _446_v3;
          let _index69 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((_441_v0).length));
          (_441_v0)[_index69] = p0;
          let _559_v81;
          _559_v81 = _dafny.MultiSet.fromElements(_532_i15);
          r1 = _module.__default.fm3(new BigNumber((_559_v81).cardinality()), globalState);
          let _560_v82;
          let _nw61 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
          _560_v82 = _nw61;
          let _561_v83;
          let _nw62 = Array((new BigNumber(3)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _500_v39;
          _nw62[(_dafny.ONE).toNumber()] = _500_v39;
          _nw62[(new BigNumber(2)).toNumber()] = _500_v39;
          _561_v83 = _nw62;
          let _562_v84;
          _562_v84 = _dafny.Set.fromElements(_561_v83, _561_v83, _561_v83);
          let _index70 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_560_v82).length));
          (_560_v82)[_index70] = (_562_v84).Union(_562_v84);
          let _index71 = _module.__default.safeIndex(new BigNumber(374), new BigNumber((_560_v82).length));
          (_560_v82)[_index71] = _562_v84;
          let _563_v85;
          _563_v85 = _dafny.Set.fromElements(!(!(!((_441_v0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((_441_v0).length))]))), !(_446_v3), _445_v2);
          _563_v85 = ((_module.D3.create_DC8(_563_v85)).dtor_cf14).Intersect(_module.__default.fm6(_446_v3, _500_v39, globalState));
          let _564_v86;
          _564_v86 = _dafny.Seq.of((_dafny.ZERO).minus(_532_i15), (_558_v80)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_558_v80).length))], new BigNumber((_447_v4).length));
          let _index72 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_558_v80).length));
          (_558_v80)[_index72] = (new BigNumber((_564_v86).length)).minus(new BigNumber(715));
        } else {
          _558_v80 = _558_v80;
          let _565_v87;
          _565_v87 = _dafny.MultiSet.fromElements(_558_v80, _558_v80, _558_v80, _558_v80, _558_v80);
          _565_v87 = _565_v87;
          _498_v37 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(316)), ((_566_v39) => function (_567_i17) {
            return _566_v39;
          })(_500_v39));
          let _568_v88;
          _568_v88 = _dafny.Seq.of(_498_v37);
          let _569_v89;
          _569_v89 = _dafny.Map.Empty.slice().updateUnsafe(_499_v38,_dafny.Seq.Concat(_568_v88, _568_v88));
          let _570_v90;
          _570_v90 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.Set.fromElements(p0)),_568_v88);
          let _571_v91;
          _571_v91 = _module.D3.create_DC8(_dafny.Set.fromElements(_446_v3));
          _569_v89 = (_569_v89).update(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_446_v3), _447_v4)).length), ((false) ? ((((_570_v90).contains(_571_v91)) ? ((_570_v90).get(_571_v91)) : (_568_v88))) : (_568_v88)));
          r2 = _445_v2;
        }
      }
      r0 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_499_v38, _499_v38), _module.__default.safeModuloInt(_499_v38, _499_v38));
      let _572_v92;
      _572_v92 = _dafny.Set.fromElements(_446_v3);
      r1 = (((_445_v2) ? (_module.__default.fm6(_445_v2, new _dafny.CodePoint('y'.codePointAt(0)), globalState)) : (_572_v92))).IsProperSubsetOf(_dafny.Set.fromElements(true, _445_v2));
      r2 = true;
      let _573_v93;
      _573_v93 = _module.D5.create_DC15(_499_v38, p0, _441_v0);
      r3 = _573_v93;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    __ctor(f13, f3, f1, f2) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus((_this).f13);
    };
    fm11(p0, globalState) {
      let _this = this;
      return _this.f3;
    };
    fm27(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt((_this).f13, (_this).f13)).minus(_module.__default.safeDivisionInt((_this).f13, (_dafny.ZERO).minus((_this).f13)));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _574_v0;
      let _nw63 = Array((new BigNumber(2)).toNumber()).fill(false);
      _574_v0 = _nw63;
      let _index73 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length));
      (_574_v0)[_index73] = _module.__default.fm3((_this).f13, globalState);
      let _index74 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length));
      (_574_v0)[_index74] = !(_this.f3) || (_this.f3);
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_574_v0).length))) {
        let _575_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_575_i0)) && ((_575_i0).isLessThan(new BigNumber((_574_v0).length))))) {
          (_574_v0)[(_575_i0)] = false;
        }
      }
      r1 = (_this).f13;
      let _576_v1;
      _576_v1 = _module.D6.create_DC17(_this.f3);
      let _pat_let_tv18 = p1;
      let _pat_let_tv19 = p1;
      let _pat_let_tv20 = p0;
      r2 = function (_source14) {
        if (_source14.is_DC17) {
          let _577___mcc_h0 = (_source14).cf28;
          let _578_cf28 = _577___mcc_h0;
          return _pat_let_tv18;
        } else if (_source14.is_DC18) {
          let _579___mcc_h1 = (_source14).cf29;
          let _580___mcc_h2 = (_source14).cf30;
          let _581___mcc_h3 = (_source14).cf31;
          let _582___mcc_h4 = (_source14).cf32;
          let _583___mcc_h5 = (_source14).cf33;
          let _584_cf33 = _583___mcc_h5;
          let _585_cf32 = _582___mcc_h4;
          let _586_cf31 = _581___mcc_h3;
          let _587_cf30 = _580___mcc_h2;
          let _588_cf29 = _579___mcc_h1;
          return _587_cf30;
        } else if (_source14.is_DC19) {
          let _589___mcc_h6 = (_source14).cf34;
          let _590_cf34 = _589___mcc_h6;
          return _pat_let_tv19;
        } else if (_source14.is_DC16) {
          let _591___mcc_h7 = (_source14).cf27;
          let _592_cf27 = _591___mcc_h7;
          return _module.__default.safeDivisionInt(new BigNumber((function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(499), new BigNumber(733))) {
              let _593_v2 = _compr_33;
              if (((new BigNumber(499)).isLessThanOrEqualTo(_593_v2)) && ((_593_v2).isLessThan(new BigNumber(733)))) {
                _coll33.push([(_593_v2).minus((_this).f13),_this.f3]);
              }
            }
            return _coll33;
          }()).length), _pat_let_tv20);
        } else {
          let _594___mcc_h8 = (_source14).cf35;
          let _595_cf35 = _594___mcc_h8;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_this).f2, (_this).f2)).length));
        }
      }(_576_v1);
      let _596_v3;
      _596_v3 = _dafny.Seq.UnicodeFromString("axj");
      _596_v3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), function (_597_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      });
      if ((new BigNumber(192)).isEqualTo(p1)) {
        let _598_v4;
        _598_v4 = new _dafny.CodePoint('m'.codePointAt(0));
        let _599_v5;
        _599_v5 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,p1));
        let _600_v6;
        _600_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_599_v5).cardinality()),p0);
        let _601_v7;
        _601_v7 = _module.D6.create_DC18(new BigNumber((_600_v6).length), new BigNumber((_596_v3).length), _this.f3, _598_v4, new BigNumber((_600_v6).length));
        _598_v4 = (_601_v7).dtor_cf32;
        let _602_v8;
        _602_v8 = _module.D3.create_DC10(new BigNumber(151), p0);
        let _603_v9;
        _603_v9 = _module.D3.create_DC11(_602_v8);
        let _604_v10;
        _604_v10 = _dafny.Seq.of(_602_v8, _602_v8);
        let _605_v11;
        _605_v11 = _dafny.Seq.of((_this).f2, _dafny.Seq.UnicodeFromString("vyygrb"));
        let _606_v12;
        _606_v12 = _module.D3.create_DC11((_604_v10)[_module.__default.safeIndex((_this).fm27(p0, (_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))], _dafny.MultiSet.FromArray(_605_v11), globalState), new BigNumber((_604_v10).length))]);
        let _607_v13;
        _607_v13 = _dafny.Set.fromElements(_606_v12, _module.D3.create_DC11(_602_v8), _606_v12, _606_v12);
        _607_v13 = (((_this).fm11((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))], globalState)) ? (((true) ? (_607_v13) : (_607_v13))) : ((((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))]) ? (_607_v13) : (_607_v13))));
        (_this).f3 = _this.f3;
        _596_v3 = _596_v3;
        r1 = p1;
      } else {
        let _index75 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length));
        (_574_v0)[_index75] = _dafny.areEqual((_this).f2, _596_v3);
        r1 = _module.__default.safeModuloInt(p0, p1);
        (_this).f3 = !(p1).isEqualTo(p1);
        if (_this.f3) {
          let _608_v14;
          _608_v14 = _module.D1.create_DC2(_dafny.Seq.UnicodeFromString("scdvhk"));
          let _609_v15;
          _609_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,function (_pat_let16_0) {
            return function (_610_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_611_dt__update_hcf2_h0) {
                  return _module.D1.create_DC2(_611_dt__update_hcf2_h0);
                }(_pat_let17_0);
              }((_this).f2);
            }(_pat_let16_0);
          }(_608_v14));
          (_this).f3 = !((_609_v15).equals(_609_v15)) || ((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))]);
          let _612_v16;
          _612_v16 = _dafny.Seq.of(new BigNumber(330), p0);
          let _613_v17;
          _613_v17 = _module.D4.create_DC13(_this.f3, ((!((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))])) ? (_612_v16) : (_612_v16)));
          let _pat_let_tv21 = _574_v0;
          let _pat_let_tv22 = _574_v0;
          let _pat_let_tv23 = _574_v0;
          let _pat_let_tv24 = _574_v0;
          _613_v17 = function (_pat_let18_0) {
            return function (_616_dt__update__tmp_h2) {
              return function (_pat_let21_0) {
                return function (_617_dt__update_hcf21_h1) {
                  return _module.D4.create_DC13(_617_dt__update_hcf21_h1, (_616_dt__update__tmp_h2).dtor_cf22);
                }(_pat_let21_0);
              }((_pat_let_tv24)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_pat_let_tv23).length))]);
            }(_pat_let18_0);
          }(function (_pat_let19_0) {
            return function (_614_dt__update__tmp_h1) {
              return function (_pat_let20_0) {
                return function (_615_dt__update_hcf21_h0) {
                  return _module.D4.create_DC13(_615_dt__update_hcf21_h0, (_614_dt__update__tmp_h1).dtor_cf22);
                }(_pat_let20_0);
              }((_pat_let_tv22)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_pat_let_tv21).length))]);
            }(_pat_let19_0);
          }(_613_v17));
          r2 = p0;
          let _618_v18;
          let _init11 = ((_619_p0) => function (_620_i2) {
            return (_620_i2).multipliedBy(_619_p0);
          })(p0);
          let _nw64 = Array((new BigNumber(25)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw64.length); _i0_11++) {
            _nw64[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _618_v18 = _nw64;
          let _621_v19;
          _621_v19 = _dafny.Map.Empty.slice().updateUnsafe((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))],_618_v18);
          _621_v19 = _621_v19;
          r2 = (p1).minus((_this).f13);
        } else {
          let _622_v20;
          _622_v20 = _module.D2.create_DC5(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f13));
          let _623_v21;
          _623_v21 = _module.D2.create_DC7(_622_v20);
          let _624_v22;
          _624_v22 = _module.D2.create_DC7(_623_v21);
          _624_v22 = _624_v22;
          let _625_v23;
          _625_v23 = _module.D5.create_DC15(p0, _this.f3, _574_v0);
          let _pat_let_tv25 = _574_v0;
          _625_v23 = function (_pat_let22_0) {
            return function (_628_dt__update__tmp_h4) {
              return function (_pat_let25_0) {
                return function (_629_dt__update_hcf25_h0) {
                  return _module.D5.create_DC15((_628_dt__update__tmp_h4).dtor_cf24, _629_dt__update_hcf25_h0, (_628_dt__update__tmp_h4).dtor_cf26);
                }(_pat_let25_0);
              }(_this.f3);
            }(_pat_let22_0);
          }(function (_pat_let23_0) {
            return function (_626_dt__update__tmp_h3) {
              return function (_pat_let24_0) {
                return function (_627_dt__update_hcf26_h0) {
                  return _module.D5.create_DC15((_626_dt__update__tmp_h3).dtor_cf24, (_626_dt__update__tmp_h3).dtor_cf25, _627_dt__update_hcf26_h0);
                }(_pat_let24_0);
              }(_pat_let_tv25);
            }(_pat_let23_0);
          }(_625_v23));
          let _630_v24;
          let _nw65 = new _module.C1();
          _nw65.__ctor();
          _630_v24 = _nw65;
          let _631_v25;
          _631_v25 = _module.D0.create_DC1(new BigNumber(981));
          let _632_v26;
          _632_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_631_v25).dtor_cf1);
          let _633_v27;
          _633_v27 = _dafny.MultiSet.fromElements(_module.__default.fm5(_632_v26, globalState));
          let _634_v28;
          _634_v28 = _module.D0.create_DC0(p1);
          let _635_v29;
          _635_v29 = _dafny.MultiSet.fromElements(!(_this.f3));
          let _636_v30;
          _636_v30 = _dafny.Seq.of(true, (_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))]);
          let _637_v31;
          _637_v31 = _module.D8.create_DC22(_635_v29);
          _633_v27 = ((true) ? (_module.__default.fm34((_634_v28).dtor_cf0, (_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))], globalState)) : (_dafny.MultiSet.fromElements(_635_v29, _dafny.MultiSet.FromArray(_636_v30), (_637_v31).dtor_cf37)));
          let _638_v32;
          _638_v32 = _dafny.Set.fromElements((_this).f13, p1, p1);
          let _639_v33;
          let _nw66 = Array((new BigNumber(8)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_638_v32).length));
          _nw66[(_dafny.ONE).toNumber()] = (_this).f13;
          _nw66[(new BigNumber(2)).toNumber()] = p1;
          _nw66[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw66[(new BigNumber(4)).toNumber()] = p1;
          _nw66[(new BigNumber(5)).toNumber()] = p0;
          _nw66[(new BigNumber(6)).toNumber()] = p0;
          _nw66[(new BigNumber(7)).toNumber()] = (new BigNumber(609)).plus(p1);
          _639_v33 = _nw66;
          let _index76 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_639_v33).length));
          (_639_v33)[_index76] = new BigNumber(382);
          let _index77 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_639_v33).length));
          (_639_v33)[_index77] = (p0).minus(p1);
        }
        let _640_v34;
        _640_v34 = _dafny.Seq.of(_596_v3, _596_v3, (_this).f2);
        let _641_v35;
        _641_v35 = _dafny.Seq.of(p0, (_this).f13);
        let _642_v36;
        _642_v36 = new _dafny.CodePoint('u'.codePointAt(0));
        let _643_v37;
        _643_v37 = _dafny.Set.fromElements((_this).f13, p1, new BigNumber((_640_v34).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).fm12(p0, (_this).f13, (_641_v35)[_module.__default.safeIndex(p1, new BigNumber((_641_v35).length))], false, globalState),_module.__default.fm31(new BigNumber(((_this).f2).length), _this.f3, p1, _642_v36, globalState))).length), p1);
        let _644_v38;
        _644_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_642_v36);
        r1 = (((_643_v37).IsProperSubsetOf(_643_v37)) ? ((_this).f13) : ((((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))]) ? ((_this).f13) : (new BigNumber((_644_v38).length)))));
      }
      let _645_v39;
      _645_v39 = _dafny.Map.Empty.slice().updateUnsafe(true,!((_574_v0)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_574_v0).length))]));
      r0 = (_645_v39).Merge(_645_v39);
      r1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState)));
      r2 = p1;
      return [r0, r1, r2];
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return _this.f3;
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f2;
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm11(p0, globalState) {
      let _this = this;
      return _this.f3;
    };
    fm25(p0, globalState) {
      let _this = this;
      return (_this).f2;
    };
    fm26(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f4).isLessThan(new BigNumber(230));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _646_v0;
      _646_v0 = _dafny.Set.fromElements(p1);
      let _index78 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index78] = (_646_v0).IsSubsetOf(_dafny.Set.fromElements(p1));
      let _index79 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index79] = (p0).isEqualTo((_this).fm12(p0, p1, (_this).f4, _this.f3, globalState));
      r2 = p0;
      let _hi4 = new BigNumber(99);
      for (let _647_i0 = p0; _647_i0.isLessThan(_hi4); _647_i0 = _647_i0.plus(_dafny.ONE)) {
        let _648_v1;
        let _nw67 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _648_v1 = _nw67;
        let _index80 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_648_v1).length));
        (_648_v1)[_index80] = _647_i0;
        let _index81 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_648_v1).length));
        (_648_v1)[_index81] = (_this).f4;
        _648_v1 = _648_v1;
        let _649_v2;
        _649_v2 = _dafny.MultiSet.fromElements(_this.f3);
        let _650_v3;
        _650_v3 = _dafny.Map.Empty.slice().updateUnsafe((_649_v2).IsSubsetOf(_649_v2),((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (_this.f3) : (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])));
        let _651_v4;
        _651_v4 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
        let _652_v5;
        _652_v5 = _dafny.Seq.of(_647_i0, p0);
        _650_v3 = (_650_v3).update((_651_v4).IsProperSubsetOf((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), (_dafny.ZERO).minus(_647_i0), new BigNumber((_652_v5).length), p0, p1)).update((_this).fm12(p1, p1, new BigNumber(((_this).f2).length), _this.f3, globalState), _module.__default.abs(new BigNumber((_dafny.MultiSet.fromElements(_this.f3)).cardinality())))), _this.f3);
        let _653_v6;
        _653_v6 = new _dafny.CodePoint('k'.codePointAt(0));
        let _index82 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index82] = !_dafny.Seq.contains((_this).f2, _653_v6);
      }
      let _654_i1;
      _654_i1 = _dafny.ZERO;
      L5: {
        while (!(!(_this.f3))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_654_i1)) {
              break L5;
            }
            _654_i1 = (_654_i1).plus(_dafny.ONE);
            let _655_v7;
            _655_v7 = _dafny.Map.Empty.slice().updateUnsafe((p1).plus(new BigNumber(360)),(_this).f4);
            let _656_v8;
            _656_v8 = _dafny.Seq.of((_this).f4);
            let _657_v9;
            _657_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber((_dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _this.f3)).length));
            let _658_v10;
            _658_v10 = _dafny.Set.fromElements(_module.D2.create_DC5(_657_v9));
            _655_v7 = (_655_v7).update((_656_v8)[_module.__default.safeIndex((_this).f4, new BigNumber((_656_v8).length))], new BigNumber((_658_v10).length));
            let _index83 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
            ((_this).f5)[_index83] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
            let _659_v11;
            _659_v11 = new _dafny.CodePoint('n'.codePointAt(0));
            _659_v11 = _659_v11;
            let _660_v12;
            let _nw68 = Array((new BigNumber(6)).toNumber());
            _nw68[(_dafny.ZERO).toNumber()] = false;
            _nw68[(_dafny.ONE).toNumber()] = _this.f3;
            _nw68[(new BigNumber(2)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
            _nw68[(new BigNumber(3)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
            _nw68[(new BigNumber(4)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
            _nw68[(new BigNumber(5)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
            _660_v12 = _nw68;
            let _661_v13;
            _661_v13 = _dafny.Set.fromElements(((_this.f3) ? (_660_v12) : (_660_v12)), _660_v12, _660_v12, _660_v12);
            _661_v13 = _dafny.Set.fromElements(_660_v12);
          }
        }
      }
      let _662_i2;
      _662_i2 = _dafny.ZERO;
      L6: {
        while (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_662_i2)) {
              break L6;
            }
            _662_i2 = (_662_i2).plus(_dafny.ONE);
            let _663_v14;
            _663_v14 = _module.D0.create_DC0((_this).f4);
            let _664_v15;
            _664_v15 = new _dafny.CodePoint('e'.codePointAt(0));
            let _665_v16;
            _665_v16 = _module.D6.create_DC18(new BigNumber(((_this).f2).length), (_663_v14).dtor_cf0, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _664_v15, (_this).f4);
            let _666_v17;
            _666_v17 = _dafny.MultiSet.fromElements(new BigNumber(340), p0, (_this).f4, (_this).f4);
            if ((((_665_v16).dtor_cf31) ? ((_dafny.MultiSet.fromElements((_this).f4)).IsSubsetOf(_666_v17)) : (_this.f3))) {
              let _667_v18;
              let _nw69 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
              _667_v18 = _nw69;
              let _668_v19;
              _668_v19 = _dafny.MultiSet.fromElements(_this.f3);
              let _index84 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_667_v18).length));
              (_667_v18)[_index84] = (((_668_v19).contains(false)) ? ((_668_v19).get(false)) : (new BigNumber(((_this).f2).length)));
              let _669_v20;
              _669_v20 = _dafny.Seq.of((_this).fm11(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], globalState), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
              let _index85 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_667_v18).length));
              (_667_v18)[_index85] = (_dafny.ZERO).minus(((((_669_v20)[_module.__default.safeIndex(p0, new BigNumber((_669_v20).length))]) ? (p0) : (p1))).plus(_module.__default.safeModuloInt(p1, _module.__default.fm0(globalState))));
              let _index86 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index86] = (_this).fm13(globalState);
              let _670_v21;
              _670_v21 = _dafny.Map.Empty.slice().updateUnsafe(_667_v18,!(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]));
              _670_v21 = (_670_v21).update(_667_v18, _this.f3);
              let _671_v22;
              _671_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
              _671_v22 = (_671_v22).update(_this.f3, (p0).minus((_this).f4));
              _670_v21 = (_dafny.Map.Empty.slice().updateUnsafe(_667_v18,true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_667_v18,false)).Merge(_670_v21));
            } else {
              r1 = _module.__default.fm0(globalState);
              let _672_v24;
              _672_v24 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll34 = new _dafny.Map();
                for (const _compr_34 of (_666_v17).Elements) {
                  let _673_v23 = _compr_34;
                  if ((_666_v17).contains(_673_v23)) {
                    _coll34.push([(_673_v23).multipliedBy((_this).f4),_this.f3]);
                  }
                }
                return _coll34;
              }(),p0);
              let _rhs78 = ((new BigNumber(((_this).f2).length)).plus(p0)).multipliedBy((_this).f4);
              let _rhs79 = (_672_v24).Merge(_672_v24);
              let _rhs80 = p1;
              r2 = _rhs78;
              _672_v24 = _rhs79;
              r1 = _rhs80;
              let _674_v25;
              let _nw70 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _674_v25 = _nw70;
              let _675_v26;
              let _nw71 = Array((new BigNumber(12)).toNumber()).fill([]);
              _675_v26 = _nw71;
              let _676_v27;
              let _nw72 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
              _676_v27 = _nw72;
              let _index87 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_676_v27).length));
              (_676_v27)[_index87] = new BigNumber(((_646_v0).Union(_646_v0)).length);
              let _677_v28;
              _677_v28 = _dafny.Seq.of(_674_v25);
              let _678_v29;
              _678_v29 = _dafny.Seq.of((_this).fm14(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _module.__default.fm3(p1, globalState), _664_v15, p0, globalState));
              let _679_v30;
              _679_v30 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))],_this.f3);
              let _index88 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              let _index89 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_676_v27).length));
              let _rhs81 = _this.f3;
              let _rhs82 = (_677_v28)[_module.__default.safeIndex(new BigNumber(759), new BigNumber((_677_v28).length))];
              let _rhs83 = _675_v26;
              let _rhs84 = ((_this).f4).minus(new BigNumber((_678_v29).length));
              let _rhs85 = ((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3)).Merge(_679_v30)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3));
              let _lhs34 = (_this).f5;
              let _lhs35 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              let _lhs36 = _676_v27;
              let _lhs37 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_676_v27).length));
              _lhs34[_lhs35] = _rhs81;
              _674_v25 = _rhs82;
              _675_v26 = _rhs83;
              _lhs36[_lhs37] = _rhs84;
              r0 = _rhs85;
              let _680_v31;
              _680_v31 = _dafny.MultiSet.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
              let _index90 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index90] = !(_680_v31).equals(_680_v31);
              let _681_v32;
              _681_v32 = _dafny.Seq.UnicodeFromString("jfebqujas");
              let _index91 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              let _rhs86 = _681_v32;
              let _rhs87 = _this.f3;
              let _rhs88 = p0;
              let _lhs38 = (_this).f5;
              let _lhs39 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              _681_v32 = _rhs86;
              _lhs38[_lhs39] = _rhs87;
              r1 = _rhs88;
            }
            let _682_v33;
            let _init12 = ((_683_p1) => function (_684_i3) {
              return (_684_i3).multipliedBy(_683_p1);
            })(p1);
            let _nw73 = Array((new BigNumber(17)).toNumber());
            for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw73.length); _i0_12++) {
              _nw73[_i0_12] = _init12(new BigNumber(_i0_12));
            }
            _682_v33 = _nw73;
            let _685_v34;
            let _nw74 = new _module.C0();
            _nw74.__ctor(_682_v33);
            _685_v34 = _nw74;
            let _686_v35;
            _686_v35 = _dafny.Seq.of((((_666_v17).contains(p1)) ? ((_666_v17).get(p1)) : (p1)));
            if (!(!(!(p1).isEqualTo(((_686_v35)[_module.__default.safeIndex(new BigNumber((_666_v17).cardinality()), new BigNumber((_686_v35).length))]).minus((_this).f4))))) {
              let _687_v36;
              _687_v36 = _dafny.Seq.of((_this).f1, (_this).f1, (_this).f1, (_this).f1);
              let _688_v37;
              _688_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber(332));
              let _689_v38;
              let _nw75 = new _module.C2();
              _nw75.__ctor(((_this).f4).plus((_dafny.ZERO).minus((_this).f4)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], (_687_v36)[_module.__default.safeIndex(p0, new BigNumber((_687_v36).length))], _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-691)), ((_690_v15) => function (_691_i4) {
                return _690_v15;
              })(_664_v15)), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("gsrcir"), _module.__default.safeIndex(new BigNumber((_688_v37).length), new BigNumber((_dafny.Seq.UnicodeFromString("gsrcir")).length)), _664_v15)));
              _689_v38 = _nw75;
              let _index92 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
              ((_this).f5)[_index92] = ((_this.f3) ? (((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) : ((_this).fm26((_this).f4, _this.f3, false, globalState)))) : (_this.f3));
              let _692_v39;
              _692_v39 = _dafny.Seq.UnicodeFromString("ynnlqps");
              let _rhs89 = _692_v39;
              let _rhs90 = (_dafny.ZERO).minus((_this).f4);
              let _rhs91 = _664_v15;
              _692_v39 = _rhs89;
              r1 = _rhs90;
              _664_v15 = _rhs91;
              (_this).f3 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
              let _693_v40;
              _693_v40 = _dafny.Map.Empty.slice().updateUnsafe((_689_v38).f13,_module.__default.safeModuloInt((_this).f4, p0));
              _693_v40 = (_693_v40).update(((_this).f4).plus(p0), (new BigNumber((_686_v35).length)).minus((_689_v38).f13));
            } else {
              r2 = p1;
              r1 = _module.__default.safeModuloInt((_this).f4, p1);
              let _index93 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_682_v33).length));
              (_682_v33)[_index93] = (_this).f4;
              let _index94 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_682_v33).length));
              (_682_v33)[_index94] = new BigNumber(234);
              r2 = (_this).f4;
              let _index95 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_682_v33).length));
              (_682_v33)[_index95] = (_682_v33)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_682_v33).length))];
            }
            let _694_v41;
            _694_v41 = _dafny.Set.fromElements((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) && (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]));
            _694_v41 = (_694_v41).Union((_694_v41).Union(_694_v41));
          }
        }
      }
      if (_this.f3) {
        let _695_v42;
        let _init13 = function (_696_i5) {
          return ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
        };
        let _nw76 = Array((new BigNumber(3)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw76.length); _i0_13++) {
          _nw76[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _695_v42 = _nw76;
        let _697_v43;
        _697_v43 = _dafny.Map.Empty.slice().updateUnsafe(_695_v42,(_this).f4);
        _697_v43 = (_697_v43).update(_695_v42, p0);
        let _698_v44;
        _698_v44 = _dafny.MultiSet.fromElements(_this.f3, true);
        let _699_v45;
        _699_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-745),_this.f3),_this.f3);
        let _700_v46;
        _700_v46 = _dafny.Seq.of(_699_v45);
        let _701_v47;
        _701_v47 = _dafny.Seq.of(new BigNumber(((_this).f2).length), p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0)).length), p1, new BigNumber(((_700_v46)[_module.__default.safeIndex(p1, new BigNumber((_700_v46).length))]).length));
        let _702_v48;
        _702_v48 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f4);
        let _703_v49;
        _703_v49 = _dafny.Seq.of(new BigNumber(((_this).f2).length), new BigNumber((_701_v47).length), new BigNumber((_702_v48).length));
        let _704_v50;
        _704_v50 = _dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
        let _705_v51;
        _705_v51 = _module.D1.create_DC3(_this.f3, new BigNumber(823), p0);
        let _706_v52;
        _706_v52 = _dafny.MultiSet.fromElements((_705_v51).dtor_cf4, (_this).f4, new BigNumber(933), new BigNumber(-450), new BigNumber((_dafny.Seq.UnicodeFromString("il")).length));
        let _707_v53;
        _707_v53 = _dafny.Map.Empty.slice().updateUnsafe(_698_v44,(_this).f4);
        let _708_v54;
        _708_v54 = _dafny.Seq.of(_this.f3, _this.f3);
        let _709_v55;
        let _nw77 = Array((new BigNumber(24)).toNumber());
        _nw77[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(-741));
        _nw77[(_dafny.ONE).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(2)).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(3)).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(4)).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(5)).toNumber()] = p1;
        _nw77[(new BigNumber(6)).toNumber()] = new BigNumber(((_this).f2).length);
        _nw77[(new BigNumber(7)).toNumber()] = (new BigNumber(-569)).plus(p1);
        _nw77[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(p0, p0);
        _nw77[(new BigNumber(9)).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(10)).toNumber()] = p1;
        _nw77[(new BigNumber(11)).toNumber()] = new BigNumber((_698_v44).cardinality());
        _nw77[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4));
        _nw77[(new BigNumber(13)).toNumber()] = (new BigNumber((_701_v47).length)).plus(p0);
        _nw77[(new BigNumber(14)).toNumber()] = p0;
        _nw77[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_704_v50).length)), new BigNumber((_706_v52).cardinality()));
        _nw77[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt((_this).f4, (((_707_v53).contains(_698_v44)) ? ((_707_v53).get(_698_v44)) : (p0)));
        _nw77[(new BigNumber(17)).toNumber()] = (_this).f4;
        _nw77[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((p0).plus(p1));
        _nw77[(new BigNumber(19)).toNumber()] = p0;
        _nw77[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-374), new BigNumber((_704_v50).length));
        _nw77[(new BigNumber(21)).toNumber()] = p0;
        _nw77[(new BigNumber(22)).toNumber()] = p0;
        _nw77[(new BigNumber(23)).toNumber()] = new BigNumber((_708_v54).length);
        _709_v55 = _nw77;
        _709_v55 = _709_v55;
        let _index96 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
        (_709_v55)[_index96] = (_this).f4;
        let _index97 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
        (_709_v55)[_index97] = ((!(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])) ? (_module.__default.fm0(globalState)) : (p0));
        if (_this.f3) {
          _708_v54 = _708_v54;
          (_this).f3 = (_module.__default.safeModuloInt(p0, new BigNumber(-580))).isLessThan(_module.__default.safeDivisionInt(p1, (_this).fm12(p0, new BigNumber(202), (_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], globalState)));
          r1 = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
          let _710_v56;
          _710_v56 = new _dafny.CodePoint('l'.codePointAt(0));
          _710_v56 = _710_v56;
          let _711_v57;
          let _nw78 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Set.Empty);
          _711_v57 = _nw78;
          let _712_v58;
          _712_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_711_v57);
          _712_v58 = (_712_v58).update(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _711_v57);
        } else {
          let _713_v59;
          let _nw79 = new _module.C1();
          _nw79.__ctor();
          _713_v59 = _nw79;
          let _index98 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index98] = false;
          let _714_v60;
          _714_v60 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_695_v42);
          (_this).f3 = (_714_v60).equals(_714_v60);
          let _715_v61;
          _715_v61 = _module.D1.create_DC2(_module.__default.fm2(_this.f3, globalState));
          let _716_v62;
          _716_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(38),((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
          let _717_v63;
          _717_v63 = _dafny.Map.Empty.slice().updateUnsafe(_716_v62,_dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), function (_718_i6) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          }));
          let _719_v64;
          _719_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_717_v63).length),(_this).f4);
          let _720_v65;
          _720_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
          let _721_v66;
          _721_v66 = new _dafny.CodePoint('b'.codePointAt(0));
          let _722_v67;
          _722_v67 = _dafny.Seq.of(_701_v47);
          let _723_v68;
          _723_v68 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))],_701_v47);
          let _724_v69;
          let _nw80 = Array((new BigNumber(23)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = _703_v49;
          _nw80[(_dafny.ONE).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus((_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))]));
          _nw80[(new BigNumber(2)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(new BigNumber(((_715_v61).dtor_cf2).length)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.of(new BigNumber(((_715_v61).dtor_cf2).length))).length)), p1);
          _nw80[(new BigNumber(4)).toNumber()] = _701_v47;
          _nw80[(new BigNumber(5)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_703_v49, _701_v47);
          _nw80[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_703_v49, _703_v49);
          _nw80[(new BigNumber(8)).toNumber()] = _module.__default.fm31((((_719_v64).contains(p0)) ? ((_719_v64).get(p0)) : (new BigNumber((_720_v65).length))), !(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]), new BigNumber((_706_v52).cardinality()), _721_v66, globalState);
          _nw80[(new BigNumber(9)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(10)).toNumber()] = _module.__default.fm31(p1, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], (_dafny.ZERO).minus(p1), _721_v66, globalState);
          _nw80[(new BigNumber(11)).toNumber()] = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), ((_725_v55) => function (_726_i7) {
            return (_725_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_725_v55).length))];
          })(_709_v55))) : (_701_v47));
          _nw80[(new BigNumber(12)).toNumber()] = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (_703_v49) : (_703_v49));
          _nw80[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_727_v50) => function (_728_i8) {
            return new BigNumber((_727_v50).length);
          })(_704_v50));
          _nw80[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat((_722_v67)[_module.__default.safeIndex((_this).f4, new BigNumber((_722_v67).length))], _701_v47);
          _nw80[(new BigNumber(15)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(16)).toNumber()] = (((_723_v68).contains(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])) ? ((_723_v68).get(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])) : (_703_v49));
          _nw80[(new BigNumber(17)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(p1));
          _nw80[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p1), _703_v49);
          _nw80[(new BigNumber(19)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(20)).toNumber()] = _701_v47;
          _nw80[(new BigNumber(21)).toNumber()] = _703_v49;
          _nw80[(new BigNumber(22)).toNumber()] = _701_v47;
          _724_v69 = _nw80;
          let _index99 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_724_v69).length));
          (_724_v69)[_index99] = _dafny.Seq.update(_dafny.Seq.of(p1, (((_702_v48).contains((_708_v54)[_module.__default.safeIndex((_this).f4, new BigNumber((_708_v54).length))])) ? ((_702_v48).get((_708_v54)[_module.__default.safeIndex((_this).f4, new BigNumber((_708_v54).length))])) : (p0)), (_this).f4, _module.__default.fm0(globalState)), _module.__default.safeIndex((_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))], new BigNumber((_dafny.Seq.of(p1, (((_702_v48).contains((_708_v54)[_module.__default.safeIndex((_this).f4, new BigNumber((_708_v54).length))])) ? ((_702_v48).get((_708_v54)[_module.__default.safeIndex((_this).f4, new BigNumber((_708_v54).length))])) : (p0)), (_this).f4, _module.__default.fm0(globalState))).length)), p1);
          let _index100 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_724_v69).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
          let _index102 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
          let _index103 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          let _rhs92 = (_722_v67)[_module.__default.safeIndex(p1, new BigNumber((_722_v67).length))];
          let _rhs93 = _721_v66;
          let _rhs94 = (new BigNumber(((_716_v62).Merge(_716_v62)).length)).plus(((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (new BigNumber((_708_v54).length)) : ((_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))])));
          let _rhs95 = (new BigNumber(556)).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(740), (_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))]));
          let _rhs96 = _this.f3;
          let _lhs40 = _724_v69;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_724_v69).length));
          let _lhs42 = _709_v55;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
          let _lhs44 = _709_v55;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
          let _lhs46 = (_this).f5;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          _lhs40[_lhs41] = _rhs92;
          _721_v66 = _rhs93;
          _lhs42[_lhs43] = _rhs94;
          _lhs44[_lhs45] = _rhs95;
          _lhs46[_lhs47] = _rhs96;
          let _729_v70;
          let _nw81 = Array((new BigNumber(18)).toNumber()).fill(_module.D4.Default());
          _729_v70 = _nw81;
          let _730_v71;
          _730_v71 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_729_v70);
          let _731_v72;
          _731_v72 = _dafny.Seq.of(_729_v70, _729_v70, _729_v70);
          let _732_v73;
          _732_v73 = _dafny.Map.Empty.slice().updateUnsafe(p1,_729_v70);
          let _733_v74;
          _733_v74 = _module.D9.create_DC25(_729_v70);
          let _734_v75;
          let _nw82 = Array((new BigNumber(23)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = _729_v70;
          _nw82[(_dafny.ONE).toNumber()] = _729_v70;
          _nw82[(new BigNumber(2)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(3)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(4)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(5)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(6)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(7)).toNumber()] = (((_730_v71).contains(_this.f3)) ? ((_730_v71).get(_this.f3)) : (_729_v70));
          _nw82[(new BigNumber(8)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(9)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(10)).toNumber()] = (_731_v72)[_module.__default.safeIndex((_this).f4, new BigNumber((_731_v72).length))];
          _nw82[(new BigNumber(11)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(12)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(13)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(14)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(15)).toNumber()] = (((_732_v73).contains((_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))])) ? ((_732_v73).get((_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))])) : ((_733_v74).dtor_cf43));
          _nw82[(new BigNumber(16)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(17)).toNumber()] = (_module.D9.create_DC25(_729_v70)).dtor_cf43;
          _nw82[(new BigNumber(18)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(19)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(20)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(21)).toNumber()] = _729_v70;
          _nw82[(new BigNumber(22)).toNumber()] = _729_v70;
          _734_v75 = _nw82;
          let _index104 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_734_v75).length));
          (_734_v75)[_index104] = _729_v70;
          let _index105 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_734_v75).length));
          let _rhs97 = _646_v0;
          let _rhs98 = _729_v70;
          let _rhs99 = _this.f3;
          let _rhs100 = _695_v42;
          let _lhs48 = _734_v75;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_734_v75).length));
          let _lhs50 = _this;
          _646_v0 = _rhs97;
          _lhs48[_lhs49] = _rhs98;
          _lhs50.f3 = _rhs99;
          _695_v42 = _rhs100;
        }
        let _index106 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length));
        (_709_v55)[_index106] = (_709_v55)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_709_v55).length))];
      } else {
        let _735_v76;
        let _nw83 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _735_v76 = _nw83;
        let _index107 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
        (_735_v76)[_index107] = p1;
        let _736_v77;
        _736_v77 = _dafny.MultiSet.fromElements((_this).fm13(globalState));
        let _737_v78;
        _737_v78 = _dafny.Set.fromElements(_736_v77);
        let _index108 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
        (_735_v76)[_index108] = _module.__default.safeModuloInt(new BigNumber((((true) ? (_737_v78) : (_737_v78))).length), p0);
        if (((_this.f3) === (false)) === ((p0).isEqualTo((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))]))) {
          let _738_v79;
          _738_v79 = _dafny.Seq.of((_this).f4, p1, p0);
          _738_v79 = _738_v79;
          let _739_v80;
          _739_v80 = _dafny.MultiSet.fromElements(p1, new BigNumber(617));
          let _index109 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          let _index110 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          let _rhs101 = (_739_v80).Union(_739_v80);
          let _rhs102 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))).plus((_this).f4);
          let _rhs103 = p0;
          let _lhs51 = _735_v76;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          let _lhs53 = _735_v76;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          _739_v80 = _rhs101;
          _lhs51[_lhs52] = _rhs102;
          _lhs53[_lhs54] = _rhs103;
          let _740_v81;
          let _nw84 = Array((new BigNumber(10)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          _nw84[(_dafny.ONE).toNumber()] = (p1).isLessThan(p1);
          _nw84[(new BigNumber(2)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          _nw84[(new BigNumber(3)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          _nw84[(new BigNumber(4)).toNumber()] = (_this.f3) === (_this.f3);
          _nw84[(new BigNumber(5)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          _nw84[(new BigNumber(6)).toNumber()] = true;
          _nw84[(new BigNumber(7)).toNumber()] = (_module.__default.fm3(new BigNumber(((_this).f2).length), globalState)) || (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
          _nw84[(new BigNumber(8)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          _nw84[(new BigNumber(9)).toNumber()] = ((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))]).isLessThan(new BigNumber(-703));
          _740_v81 = _nw84;
          let _741_v82;
          _741_v82 = _dafny.Set.fromElements(_this.f3, !(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]), _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _this.f3);
          let _742_v83;
          _742_v83 = _dafny.Seq.of(((_this.f3) ? (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) : (_this.f3)), !(_741_v82).contains(_this.f3));
          let _743_v84;
          let _nw85 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _743_v84 = _nw85;
          let _744_v85;
          _744_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_738_v79);
          let _index111 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_743_v84).length));
          (_743_v84)[_index111] = _744_v85;
          let _745_v86;
          _745_v86 = new _dafny.CodePoint('c'.codePointAt(0));
          let _index112 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_743_v84).length));
          let _rhs104 = (new BigNumber((_738_v79).length)).multipliedBy(p1);
          let _rhs105 = _740_v81;
          let _rhs106 = _dafny.Seq.update(_dafny.Seq.Concat(_742_v83, _dafny.Seq.update(_dafny.Seq.Concat(_742_v83, _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], false, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])), _module.__default.safeIndex(new BigNumber((_module.__default.fm31((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], (_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], _745_v86, globalState)).length), new BigNumber((_dafny.Seq.Concat(_742_v83, _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], false, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]))).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])), _module.__default.safeIndex((_this).f4, new BigNumber((_dafny.Seq.Concat(_742_v83, _dafny.Seq.update(_dafny.Seq.Concat(_742_v83, _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], false, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])), _module.__default.safeIndex(new BigNumber((_module.__default.fm31((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], (_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], _745_v86, globalState)).length), new BigNumber((_dafny.Seq.Concat(_742_v83, _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], false, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]))).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]))).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
          let _rhs107 = _744_v85;
          let _lhs55 = _743_v84;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_743_v84).length));
          r1 = _rhs104;
          _740_v81 = _rhs105;
          _742_v83 = _rhs106;
          _lhs55[_lhs56] = _rhs107;
          let _746_v87;
          _746_v87 = _module.D5.create_DC15(new BigNumber(((_this).f2).length), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], _740_v81);
          let _747_v88;
          _747_v88 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_746_v87).dtor_cf26);
          _740_v81 = (((_747_v88).contains(p0)) ? ((_747_v88).get(p0)) : (_740_v81));
          let _init14 = ((_748_p0) => function (_749_i9) {
            return (_749_i9).plus(_748_p0);
          })(p0);
          let _nw86 = Array((new BigNumber(14)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw86.length); _i0_14++) {
            _nw86[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _735_v76 = _nw86;
        } else {
          let _index113 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index113] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))];
          let _750_v89;
          let _nw87 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _750_v89 = _nw87;
          let _751_v90;
          _751_v90 = new _dafny.CodePoint('i'.codePointAt(0));
          let _index114 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_750_v89).length));
          (_750_v89)[_index114] = _751_v90;
          let _index115 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_750_v89).length));
          let _index116 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          let _rhs108 = ((!(((_this).f4).isLessThanOrEqualTo(new BigNumber(737)))) ? (_751_v90) : (((((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) ? (_751_v90) : (_751_v90))));
          let _rhs109 = p0;
          let _lhs57 = _750_v89;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_750_v89).length));
          let _lhs59 = _735_v76;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length));
          _lhs57[_lhs58] = _rhs108;
          _lhs59[_lhs60] = _rhs109;
          let _752_v91;
          _752_v91 = _dafny.Seq.of(_this.f3, _this.f3);
          let _753_v92;
          _753_v92 = _dafny.Seq.of((_752_v91)[_module.__default.safeIndex((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], new BigNumber((_752_v91).length))], ((_this.f3) ? (_this.f3) : (((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))])));
          _753_v92 = _module.__default.fm7(!(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]) || (true), ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], globalState);
          let _754_v93;
          let _nw88 = new _module.C1();
          _nw88.__ctor();
          _754_v93 = _nw88;
          let _755_v94;
          _755_v94 = _dafny.Seq.UnicodeFromString("mmwbhci");
          let _756_v95;
          _756_v95 = _dafny.MultiSet.fromElements(p1, (_this).f4);
          let _757_v96;
          _757_v96 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("puaoww"));
          let _758_v97;
          _758_v97 = _dafny.Seq.of((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))]);
          let _759_v98;
          _759_v98 = _module.D2.create_DC6((_755_v94)[_module.__default.safeIndex((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], new BigNumber((_755_v94).length))], _this.f3, _this.f3, new BigNumber(((_757_v96)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_758_v97).length)), new BigNumber((_757_v96).length))]).length));
          let _760_v99;
          _760_v99 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))],p0);
          let _761_v100;
          _761_v100 = _dafny.Seq.of(new BigNumber(572), (_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], new BigNumber((_dafny.Seq.Concat(_module.__default.fm31((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], new BigNumber((_756_v95).cardinality()), (_759_v98).dtor_cf9, globalState), _758_v97)).length), ((((_760_v99).contains(true)) ? ((_760_v99).get(true)) : ((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))]))).plus((_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))]), new BigNumber(363));
          let _rhs110 = _755_v94;
          let _rhs111 = _dafny.Seq.Concat(_755_v94, (_this).f2);
          let _rhs112 = _dafny.Seq.Concat(_761_v100, _761_v100);
          let _rhs113 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(p0).minus(new BigNumber((_761_v100).length)));
          let _rhs114 = (_750_v89)[_module.__default.safeIndex(new BigNumber(837), new BigNumber((_750_v89).length))];
          _755_v94 = _rhs110;
          _755_v94 = _rhs111;
          _761_v100 = _rhs112;
          _760_v99 = _rhs113;
          _751_v90 = _rhs114;
        }
        r1 = p1;
        r1 = (_735_v76)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_735_v76).length))];
        let _762_v101;
        _762_v101 = new _dafny.CodePoint('j'.codePointAt(0));
        let _763_v102;
        _763_v102 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('d'.codePointAt(0)), _762_v101, _762_v101);
        _763_v102 = (_763_v102).Difference(_763_v102);
      }
      let _764_v103;
      _764_v103 = _dafny.MultiSet.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))]);
      let _765_v104;
      _765_v104 = _module.D8.create_DC23((((_764_v103).contains(true)) ? ((_764_v103).get(true)) : (p0)), (_this).f4, _this.f3);
      let _766_v105;
      _766_v105 = _dafny.Map.Empty.slice().updateUnsafe((_765_v104).dtor_cf40,_this.f3);
      r0 = ((_766_v105).Merge(_766_v105)).update(_this.f3, !_dafny.areEqual((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(416)), function (_767_i10) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })));
      r1 = (_this).f4;
      r2 = p0;
      return [r0, r1, r2];
    }
    m10(p0, globalState) {
      let _this = this;
      let r0 = _module.D5.Default();
      let r1 = false;
      if (_this.f3) {
        let _768_v0;
        _768_v0 = new _dafny.CodePoint('o'.codePointAt(0));
        _768_v0 = _768_v0;
        let _769_v1;
        let _nw89 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _769_v1 = _nw89;
        _769_v1 = _769_v1;
        let _770_v2;
        _770_v2 = _dafny.Seq.UnicodeFromString("g");
        _770_v2 = (_this).f2;
        let _771_v3;
        _771_v3 = new BigNumber(-585);
        let _rhs115 = _this.f3;
        let _rhs116 = _771_v3;
        let _lhs61 = _this;
        _lhs61.f3 = _rhs115;
        _771_v3 = _rhs116;
        let _772_v4;
        let _nw90 = new _module.C1();
        _nw90.__ctor();
        _772_v4 = _nw90;
      } else {
        let _773_v5;
        _773_v5 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        _773_v5 = (_773_v5).update(!((p0) || (_this.f3)), !(((_this).f4).isLessThanOrEqualTo((_this).f4)));
        let _774_v6;
        _774_v6 = _dafny.Set.fromElements(false, p0, p0);
        let _index117 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index117] = (_774_v6).IsSubsetOf(_774_v6);
        let _775_v7;
        _775_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f4, (_this).f4),true);
        let _index118 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
        let _rhs117 = ((_this.f3) ? (_dafny.areEqual((_this).f2, (_this).f2)) : (_this.f3));
        let _rhs118 = _775_v7;
        let _rhs119 = true;
        let _lhs62 = (_this).f5;
        let _lhs63 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
        _lhs62[_lhs63] = _rhs117;
        _775_v7 = _rhs118;
        r1 = _rhs119;
        let _776_v8;
        _776_v8 = _dafny.Seq.UnicodeFromString("w");
        _776_v8 = _776_v8;
        let _777_v9;
        let _nw91 = new _module.C2();
        _nw91.__ctor(_module.__default.safeDivisionInt((_this).f4, (_dafny.ZERO).minus((_this).f4)), true, (_this).f1, _776_v8);
        _777_v9 = _nw91;
        let _778_v10;
        _778_v10 = _module.D0.create_DC1((_this).f4);
        let _source15 = _778_v10;
        if (_source15.is_DC1) {
          let _779___mcc_h0 = (_source15).cf1;
          let _780_cf1 = _779___mcc_h0;
          let _781_v11;
          let _nw92 = Array((new BigNumber(3)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("mhq");
          _nw92[(_dafny.ONE).toNumber()] = (_this).f2;
          _nw92[(new BigNumber(2)).toNumber()] = (_this).f2;
          _781_v11 = _nw92;
          _781_v11 = _781_v11;
          _776_v8 = (_this).f2;
          let _782_v12;
          let _init15 = function (_783_i0) {
            return (_783_i0).minus(new BigNumber(792));
          };
          let _nw93 = Array((new BigNumber(28)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw93.length); _i0_15++) {
            _nw93[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _782_v12 = _nw93;
          let _784_v13;
          let _nw94 = new _module.C0();
          _nw94.__ctor(_782_v12);
          _784_v13 = _nw94;
          let _index119 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_782_v12).length));
          (_782_v12)[_index119] = _module.__default.fm0(globalState);
          let _index120 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_782_v12).length));
          (_782_v12)[_index120] = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), new BigNumber((_775_v7).length));
        } else {
          let _785___mcc_h1 = (_source15).cf0;
          let _786_cf0 = _785___mcc_h1;
          let _index121 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index121] = _this.f3;
          let _787_v14;
          _787_v14 = _dafny.Map.Empty.slice().updateUnsafe((_777_v9).f13,new _dafny.CodePoint('l'.codePointAt(0)));
          let _788_v15;
          _788_v15 = _module.D8.create_DC24(new _dafny.CodePoint('q'.codePointAt(0)), _773_v5);
          let _789_v16;
          _789_v16 = _dafny.Seq.of(_788_v15);
          let _790_v17;
          _790_v17 = _dafny.Map.Empty.slice().updateUnsafe(_787_v14,_789_v16);
          let _791_v18;
          let _nw95 = Array((new BigNumber(27)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = p0;
          _nw95[(_dafny.ONE).toNumber()] = false;
          _nw95[(new BigNumber(2)).toNumber()] = p0;
          _nw95[(new BigNumber(3)).toNumber()] = true;
          _nw95[(new BigNumber(4)).toNumber()] = p0;
          _nw95[(new BigNumber(5)).toNumber()] = true;
          _nw95[(new BigNumber(6)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(7)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))];
          _nw95[(new BigNumber(8)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))];
          _nw95[(new BigNumber(9)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(10)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(11)).toNumber()] = true;
          _nw95[(new BigNumber(12)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(13)).toNumber()] = false;
          _nw95[(new BigNumber(14)).toNumber()] = !(_this.f3);
          _nw95[(new BigNumber(15)).toNumber()] = _module.__default.fm3(_786_cf0, globalState);
          _nw95[(new BigNumber(16)).toNumber()] = p0;
          _nw95[(new BigNumber(17)).toNumber()] = true;
          _nw95[(new BigNumber(18)).toNumber()] = false;
          _nw95[(new BigNumber(19)).toNumber()] = true;
          _nw95[(new BigNumber(20)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(21)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(22)).toNumber()] = !(false);
          _nw95[(new BigNumber(23)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))];
          _nw95[(new BigNumber(24)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))];
          _nw95[(new BigNumber(25)).toNumber()] = _this.f3;
          _nw95[(new BigNumber(26)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))];
          _791_v18 = _nw95;
          let _792_v19;
          _792_v19 = _module.D5.create_DC15(new BigNumber(573), p0, _791_v18);
          let _793_v20;
          _793_v20 = _dafny.Seq.of(p0, false);
          let _794_v21;
          _794_v21 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_793_v20).length));
          let _795_v22;
          _795_v22 = _dafny.Seq.of(_794_v21, _794_v21, _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f4));
          let _796_v24;
          _796_v24 = _dafny.Set.fromElements(_794_v21, _794_v21, _794_v21);
          let _797_v25;
          _797_v25 = new _dafny.CodePoint('e'.codePointAt(0));
          let _798_v26;
          _798_v26 = _dafny.Seq.of(_module.__default.fm2(((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))], globalState), _776_v8);
          let _799_v27;
          _799_v27 = _dafny.Set.fromElements((_this).fm25(!((_module.D2.create_DC6(_797_v25, p0, (_this).fm13(globalState), (_this).f4)).dtor_cf10), globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(560)), function (_800_i1) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          }), (_798_v26)[_module.__default.safeIndex((_this).f4, new BigNumber((_798_v26).length))]);
          let _801_v28;
          let _nw96 = Array((new BigNumber(23)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _this.f3;
          _nw96[(_dafny.ONE).toNumber()] = _this.f3;
          _nw96[(new BigNumber(2)).toNumber()] = (((_773_v5).contains(((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))])) ? ((_773_v5).get(((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))])) : (p0));
          _nw96[(new BigNumber(3)).toNumber()] = (((_775_v7).contains((_777_v9).f13)) ? ((_775_v7).get((_777_v9).f13)) : (p0));
          _nw96[(new BigNumber(4)).toNumber()] = (!(false)) === (p0);
          _nw96[(new BigNumber(5)).toNumber()] = _this.f3;
          _nw96[(new BigNumber(6)).toNumber()] = (_792_v19).dtor_cf25;
          _nw96[(new BigNumber(7)).toNumber()] = p0;
          _nw96[(new BigNumber(8)).toNumber()] = p0;
          _nw96[(new BigNumber(9)).toNumber()] = (((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))]) || (p0);
          _nw96[(new BigNumber(10)).toNumber()] = _this.f3;
          _nw96[(new BigNumber(11)).toNumber()] = _dafny.areEqual(_776_v8, _776_v8);
          _nw96[(new BigNumber(12)).toNumber()] = _this.f3;
          _nw96[(new BigNumber(13)).toNumber()] = p0;
          _nw96[(new BigNumber(14)).toNumber()] = !(_this.f3);
          _nw96[(new BigNumber(15)).toNumber()] = !_dafny.areEqual((_this).f2, _776_v8);
          _nw96[(new BigNumber(16)).toNumber()] = p0;
          _nw96[(new BigNumber(17)).toNumber()] = !((_this).fm11(p0, globalState));
          _nw96[(new BigNumber(18)).toNumber()] = (_796_v24).IsSubsetOf(function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_795_v22).Elements) {
              let _802_v23 = _compr_35;
              if (_dafny.Seq.contains(_795_v22, _802_v23)) {
                _coll35.add(_802_v23);
              }
            }
            return _coll35;
          }());
          _nw96[(new BigNumber(19)).toNumber()] = !(false) || (((_this).f5)[_module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length))]);
          _nw96[(new BigNumber(20)).toNumber()] = (_799_v27).IsProperSubsetOf(_799_v27);
          _nw96[(new BigNumber(21)).toNumber()] = true;
          _nw96[(new BigNumber(22)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_774_v6).length))).isLessThanOrEqualTo((_777_v9).f13);
          _801_v28 = _nw96;
          let _index122 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
          let _rhs120 = _790_v17;
          let _rhs121 = !(_775_v7).contains((_this).f4);
          let _rhs122 = _801_v28;
          let _lhs64 = (_this).f5;
          let _lhs65 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
          _790_v17 = _rhs120;
          _lhs64[_lhs65] = _rhs121;
          _801_v28 = _rhs122;
          let _index123 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
          let _rhs123 = ((_777_v9).f13).isLessThan((_this).f4);
          let _rhs124 = _dafny.Map.Empty.slice().updateUnsafe((_777_v9).f13,_797_v25);
          let _lhs66 = (_this).f5;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(786), new BigNumber(((_this).f5).length));
          _lhs66[_lhs67] = _rhs123;
          _787_v14 = _rhs124;
          let _803_v29;
          _803_v29 = _dafny.MultiSet.fromElements((_777_v9).f13, (_dafny.ZERO).minus((_777_v9).f13));
          let _804_v30;
          _804_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_803_v29).contains(new BigNumber(433))) ? ((_803_v29).get(new BigNumber(433))) : ((_dafny.ZERO).minus(_786_cf0))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(295)), ((_805_v25) => function (_806_i2) {
            return _805_v25;
          })(_797_v25)));
          _786_cf0 = (((_777_v9).f13).plus(new BigNumber((_804_v30).length))).minus((_this).f4);
        }
      }
      let _807_v31;
      _807_v31 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f4));
      if ((new BigNumber((_807_v31).length)).isLessThanOrEqualTo(((_this).f4).multipliedBy((_this).f4))) {
        if (_this.f3) {
          let _808_v32;
          _808_v32 = _module.D5.create_DC15((_this).f4, p0, (_this).f5);
          let _809_v33;
          _809_v33 = _dafny.Map.Empty.slice().updateUnsafe((_807_v31)[_module.__default.safeIndex((_this).f4, new BigNumber((_807_v31).length))],(_808_v32).dtor_cf25);
          _809_v33 = (_809_v33).update((_this).f4, _this.f3);
          let _810_v34;
          _810_v34 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0));
          let _811_v35;
          _811_v35 = new _dafny.CodePoint('d'.codePointAt(0));
          let _812_v36;
          let _nw97 = new _module.C2();
          _nw97.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-83)), ((_813_v33) => function (_814_i3) {
            return _813_v33;
          })(_809_v33)), _810_v34)).length), p0, ((_this).f1).update((_this).fm14(p0, true, new _dafny.CodePoint('y'.codePointAt(0)), (_this).f4, globalState), _this.f3), _dafny.Seq.update((_this).f2, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f4), new BigNumber(((_this).f2).length)), _811_v35));
          _812_v36 = _nw97;
          let _815_v37;
          _815_v37 = _module.D6.create_DC18((_this).f4, (_this).f4, p0, _811_v35, (_dafny.ZERO).minus((_this).f4));
          let _pat_let_tv26 = _811_v35;
          _811_v35 = (function (_pat_let26_0) {
            return function (_816_dt__update__tmp_h0) {
              return function (_pat_let27_0) {
                return function (_817_dt__update_hcf32_h0) {
                  return function (_pat_let28_0) {
                    return function (_818_dt__update_hcf30_h0) {
                      return function (_pat_let29_0) {
                        return function (_819_dt__update_hcf31_h0) {
                          return _module.D6.create_DC18((_816_dt__update__tmp_h0).dtor_cf29, _818_dt__update_hcf30_h0, _819_dt__update_hcf31_h0, _817_dt__update_hcf32_h0, (_816_dt__update__tmp_h0).dtor_cf33);
                        }(_pat_let29_0);
                      }(_this.f3);
                    }(_pat_let28_0);
                  }((_this).f4);
                }(_pat_let27_0);
              }(_pat_let_tv26);
            }(_pat_let26_0);
          }(_815_v37)).dtor_cf32;
          _807_v31 = _807_v31;
          let _820_v38;
          _820_v38 = _dafny.Seq.UnicodeFromString("kqd");
          _820_v38 = _820_v38;
        } else {
          let _821_v39;
          let _nw98 = Array((new BigNumber(25)).toNumber()).fill([]);
          _821_v39 = _nw98;
          let _index124 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_821_v39).length));
          (_821_v39)[_index124] = (_this).f5;
          let _index125 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_821_v39).length));
          (_821_v39)[_index125] = (_this).f5;
          let _822_v40;
          _822_v40 = new BigNumber(729);
          let _823_v41;
          let _nw99 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
          _823_v41 = _nw99;
          let _824_v42;
          let _init16 = ((_825_p0) => function (_826_i4) {
            return _module.__default.safeDivisionInt(_826_i4, new BigNumber((_dafny.MultiSet.fromElements(_this.f3, _825_p0, _this.f3)).cardinality()));
          })(p0);
          let _nw100 = Array((new BigNumber(18)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw100.length); _i0_16++) {
            _nw100[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _824_v42 = _nw100;
          let _index126 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_824_v42).length));
          (_824_v42)[_index126] = _module.__default.fm0(globalState);
          let _index127 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_824_v42).length));
          let _rhs125 = ((_this).fm12(_822_v40, new BigNumber(760), _822_v40, _this.f3, globalState)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(931)), function (_827_i5) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })).length));
          let _rhs126 = ((_this.f3) ? (_823_v41) : (((p0) ? (_823_v41) : (_823_v41))));
          let _rhs127 = _822_v40;
          let _lhs68 = _824_v42;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_824_v42).length));
          _822_v40 = _rhs125;
          _823_v41 = _rhs126;
          _lhs68[_lhs69] = _rhs127;
          let _828_v43;
          let _nw101 = new _module.C0();
          _nw101.__ctor(_824_v42);
          _828_v43 = _nw101;
          let _829_v44;
          _829_v44 = new _dafny.CodePoint('r'.codePointAt(0));
          let _830_v45;
          _830_v45 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f2);
          let _831_v46;
          _831_v46 = _module.D3.create_DC10(_822_v40, (_824_v42)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_824_v42).length))]);
          let _832_v47;
          _832_v47 = _module.D3.create_DC11(_831_v46);
          let _833_v48;
          _833_v48 = _dafny.Map.Empty.slice().updateUnsafe(_832_v47,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), function (_834_i6) {
            return (_this).f4;
          })).length));
          _829_v44 = (_828_v43).fm29(_830_v45, _833_v48, new BigNumber((_dafny.Seq.of((_this).f5, (_this).f5, (_821_v39)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_821_v39).length))], (_this).f5, (_821_v39)[_module.__default.safeIndex(new BigNumber(302), new BigNumber((_821_v39).length))])).length), globalState);
          let _835_v49;
          _835_v49 = _dafny.Seq.of(_this.f3);
          let _index128 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_824_v42).length));
          (_824_v42)[_index128] = (_dafny.ZERO).minus(((_dafny.Seq.IsPrefixOf((_this).f2, (_this).f2)) ? (((p0) ? (new BigNumber((_835_v49).length)) : (_822_v40))) : ((_this).f4)));
        }
        r1 = true;
        let _836_v50;
        let _nw102 = new _module.C1();
        _nw102.__ctor();
        _836_v50 = _nw102;
        if (_this.f3) {
          let _837_v51;
          _837_v51 = _dafny.Seq.of(_this.f3, _this.f3);
          let _838_v52;
          _838_v52 = _dafny.Set.fromElements(_837_v51);
          let _839_v53;
          let _nw103 = Array((new BigNumber(7)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = (_this).f4;
          _nw103[(_dafny.ONE).toNumber()] = new BigNumber(449);
          _nw103[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f4);
          _nw103[(new BigNumber(3)).toNumber()] = new BigNumber((_838_v52).length);
          _nw103[(new BigNumber(4)).toNumber()] = (_807_v31)[_module.__default.safeIndex((_this).f4, new BigNumber((_807_v31).length))];
          _nw103[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw103[(new BigNumber(6)).toNumber()] = (_this).f4;
          _839_v53 = _nw103;
          let _840_v54;
          let _nw104 = new _module.C0();
          _nw104.__ctor(((p0) ? (_839_v53) : (_839_v53)));
          _840_v54 = _nw104;
          let _841_v55;
          _841_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_dafny.Seq.UnicodeFromString("kaukntjia"));
          _841_v55 = (_841_v55).update((_dafny.ZERO).minus((_this).f4), (_this).f2);
          r1 = !(p0);
          let _842_v56;
          _842_v56 = new BigNumber(-984);
          let _843_v57;
          _843_v57 = _dafny.MultiSet.fromElements(_this.f3, p0);
          let _844_v58;
          _844_v58 = _module.D3.create_DC9(_842_v56, _this.f3);
          _842_v56 = ((new BigNumber((_843_v57).cardinality())).multipliedBy((_this).f4)).minus(((!((_844_v58).dtor_cf16)) ? (new BigNumber((_837_v51).length)) : (_842_v56)));
          _842_v56 = (_this).f4;
        } else {
          let _845_v59;
          _845_v59 = new BigNumber(283);
          _845_v59 = (_dafny.ZERO).minus((((_this).f4).minus((_this).f4)).plus(_845_v59));
          let _846_v60;
          let _nw105 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _846_v60 = _nw105;
          let _847_v61;
          _847_v61 = _dafny.MultiSet.fromElements(_845_v59, (_this).f4);
          let _index129 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_846_v60).length));
          (_846_v60)[_index129] = _847_v61;
          let _848_v62;
          _848_v62 = _dafny.MultiSet.fromElements(p0);
          let _849_v63;
          _849_v63 = new _dafny.CodePoint('r'.codePointAt(0));
          let _850_v64;
          _850_v64 = _module.D2.create_DC6(_849_v63, p0, p0, _845_v59);
          let _851_v65;
          _851_v65 = _dafny.Seq.of(_807_v31, _dafny.Seq.Concat(_807_v31, _module.__default.fm31(_845_v59, p0, new BigNumber((_dafny.Set.fromElements(p0)).length), (_850_v64).dtor_cf9, globalState)));
          let _852_v66;
          _852_v66 = _module.D0.create_DC1((_this).f4);
          let _853_v67;
          _853_v67 = _dafny.Seq.of(_this.f3);
          let _index130 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_846_v60).length));
          let _rhs128 = (_848_v62).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(p0, !(false))));
          let _rhs129 = (_851_v65)[_module.__default.safeIndex((_845_v59).minus((((_848_v62).contains(p0)) ? ((_848_v62).get(p0)) : ((_this).f4))), new BigNumber((_851_v65).length))];
          let _rhs130 = (_dafny.MultiSet.FromArray(_853_v67)).contains(((_this).f4).isLessThan((_852_v66).dtor_cf1));
          let _rhs131 = _847_v61;
          let _lhs70 = _this;
          let _lhs71 = _this;
          let _lhs72 = _846_v60;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_846_v60).length));
          _lhs70.f3 = _rhs128;
          _807_v31 = _rhs129;
          _lhs71.f3 = _rhs130;
          _lhs72[_lhs73] = _rhs131;
          let _854_v68;
          let _init17 = ((_855_p0) => function (_856_i7) {
            return _855_p0;
          })(p0);
          let _nw106 = Array((new BigNumber(4)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw106.length); _i0_17++) {
            _nw106[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _854_v68 = _nw106;
          let _857_v69;
          _857_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
          _854_v68 = (((_857_v69).contains((_845_v59).isEqualTo((_this).f4))) ? ((_857_v69).get((_845_v59).isEqualTo((_this).f4))) : ((_this).f5));
          let _858_v70;
          _858_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,new BigNumber(-549));
          _858_v70 = (_858_v70).update(_dafny.Seq.Create(_module.__default.abs(_dafny.ONE), ((_859_v63) => function (_860_i8) {
            return _859_v63;
          })(_849_v63)), (_807_v31)[_module.__default.safeIndex((_this).f4, new BigNumber((_807_v31).length))]);
          (_this).f3 = (_845_v59).isLessThan(((_this).f4).multipliedBy(_845_v59));
        }
        if (((!(_this.f3)) ? (_this.f3) : (_this.f3))) {
          let _861_v71;
          let _nw107 = Array((_dafny.ONE).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = (_this).f4;
          _861_v71 = _nw107;
          let _862_v72;
          let _nw108 = new _module.C0();
          _nw108.__ctor(_861_v71);
          _862_v72 = _nw108;
          let _863_v73;
          _863_v73 = new BigNumber(949);
          let _864_v74;
          _864_v74 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
          let _865_v75;
          _865_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,!(_this.f3));
          _863_v73 = (_dafny.ZERO).minus(((!(((_dafny.ZERO).minus((((_864_v74).contains(p0)) ? ((_864_v74).get(p0)) : (_863_v73)))).isLessThanOrEqualTo((_dafny.ZERO).minus(_863_v73)))) ? (new BigNumber((_865_v75).length)) : (new BigNumber(694))));
          let _866_v76;
          _866_v76 = new _dafny.CodePoint('f'.codePointAt(0));
          _866_v76 = new _dafny.CodePoint('k'.codePointAt(0));
          (_this).f3 = (_863_v73).isEqualTo(_module.__default.safeDivisionInt(_863_v73, _863_v73));
          _863_v73 = _863_v73;
        } else {
          let _867_v77;
          let _nw109 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _867_v77 = _nw109;
          let _868_v78;
          _868_v78 = _module.D5.create_DC14(_867_v77);
          let _pat_let_tv27 = _867_v77;
          r0 = function (_pat_let30_0) {
            return function (_869_dt__update__tmp_h1) {
              return function (_pat_let31_0) {
                return function (_870_dt__update_hcf23_h0) {
                  return _module.D5.create_DC14(_870_dt__update_hcf23_h0);
                }(_pat_let31_0);
              }(_pat_let_tv27);
            }(_pat_let30_0);
          }(_868_v78);
          r1 = ((!(p0)) ? (p0) : (_this.f3));
          let _871_v79;
          let _nw110 = new _module.C0();
          _nw110.__ctor(_867_v77);
          _871_v79 = _nw110;
          let _872_v80;
          _872_v80 = _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("q"), _dafny.Seq.UnicodeFromString("bpkogcfjm")));
          _872_v80 = _dafny.Seq.of(p0, !(p0), !(!(false)));
          (_this).f3 = _this.f3;
        }
      } else {
        let _873_v81;
        _873_v81 = _dafny.MultiSet.fromElements(false);
        let _874_v82;
        _874_v82 = new _dafny.CodePoint('a'.codePointAt(0));
        let _875_v83;
        _875_v83 = _dafny.Map.Empty.slice().updateUnsafe(_807_v31,_874_v82);
        let _876_v84;
        _876_v84 = _dafny.Seq.of(_this.f3);
        let _877_v85;
        _877_v85 = _module.D8.create_DC23((_this).f4, _module.__default.fm0(globalState), _this.f3);
        let _878_v86;
        let _nw111 = Array((new BigNumber(25)).toNumber());
        _nw111[(_dafny.ZERO).toNumber()] = !(p0);
        _nw111[(_dafny.ONE).toNumber()] = _this.f3;
        _nw111[(new BigNumber(2)).toNumber()] = true;
        _nw111[(new BigNumber(3)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(4)).toNumber()] = !(p0);
        _nw111[(new BigNumber(5)).toNumber()] = !((_873_v81).IsDisjointFrom(_873_v81));
        _nw111[(new BigNumber(6)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(7)).toNumber()] = !(_this.f3);
        _nw111[(new BigNumber(8)).toNumber()] = p0;
        _nw111[(new BigNumber(9)).toNumber()] = p0;
        _nw111[(new BigNumber(10)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(11)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(12)).toNumber()] = !_dafny.Seq.contains((_this).f2, (((_875_v83).contains(_module.__default.fm31(new BigNumber(975), p0, (_this).f4, _874_v82, globalState))) ? ((_875_v83).get(_module.__default.fm31(new BigNumber(975), p0, (_this).f4, _874_v82, globalState))) : (_874_v82)));
        _nw111[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_876_v84, _876_v84);
        _nw111[(new BigNumber(14)).toNumber()] = (_877_v85).dtor_cf40;
        _nw111[(new BigNumber(15)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_this).f4), _807_v31);
        _nw111[(new BigNumber(16)).toNumber()] = p0;
        _nw111[(new BigNumber(17)).toNumber()] = p0;
        _nw111[(new BigNumber(18)).toNumber()] = p0;
        _nw111[(new BigNumber(19)).toNumber()] = true;
        _nw111[(new BigNumber(20)).toNumber()] = !(p0);
        _nw111[(new BigNumber(21)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(22)).toNumber()] = !(_this.f3);
        _nw111[(new BigNumber(23)).toNumber()] = _this.f3;
        _nw111[(new BigNumber(24)).toNumber()] = _this.f3;
        _878_v86 = _nw111;
        _878_v86 = _878_v86;
        r1 = !((_this).f4).isEqualTo((_this).f4);
        let _879_v87;
        _879_v87 = new BigNumber(810);
        let _880_v88;
        _880_v88 = _module.D5.create_DC15(_879_v87, _this.f3, (_this).f5);
        let _rhs132 = _module.__default.fm0(globalState);
        let _rhs133 = (_880_v88).dtor_cf26;
        let _rhs134 = _879_v87;
        _879_v87 = _rhs132;
        _878_v86 = _rhs133;
        _879_v87 = _rhs134;
        let _881_v89;
        let _nw112 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _881_v89 = _nw112;
        let _882_v90;
        _882_v90 = _module.D0.create_DC0(new BigNumber(((_this).f2).length));
        let _index131 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length));
        (_881_v89)[_index131] = (_882_v90).dtor_cf0;
        let _index132 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length));
        (_881_v89)[_index132] = _879_v87;
        if (_this.f3) {
          let _883_v92;
          let _init18 = ((_884_p0, _885_v82, _886_v87) => function (_887_i9) {
            return function () {
              let _coll36 = new _dafny.Set();
              for (const _compr_36 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_884_p0,_885_v82),_886_v87)).Keys.Elements) {
                let _888_v91 = _compr_36;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_884_p0,_885_v82),_886_v87)).contains(_888_v91)) {
                  _coll36.add(_888_v91);
                }
              }
              return _coll36;
            }();
          })(p0, _874_v82, _879_v87);
          let _nw113 = Array((new BigNumber(26)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw113.length); _i0_18++) {
            _nw113[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _883_v92 = _nw113;
          let _889_v93;
          _889_v93 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_874_v82);
          let _890_v94;
          _890_v94 = _dafny.Set.fromElements(_889_v93, _889_v93);
          let _index133 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_883_v92).length));
          (_883_v92)[_index133] = _890_v94;
          let _index134 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_883_v92).length));
          (_883_v92)[_index134] = _890_v94;
          let _891_v95;
          _891_v95 = _dafny.Seq.of(_881_v89);
          _891_v95 = _dafny.Seq.of(_881_v89);
          r1 = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of((_this).f4, (_881_v89)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-759)), function (_892_i10) {
            return new BigNumber(323);
          })), _807_v31);
          let _index135 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index135] = false;
          let _893_v96;
          _893_v96 = _dafny.MultiSet.fromElements(_876_v84, _876_v84);
          let _index136 = _module.__default.safeIndex(new BigNumber(254), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index136] = ((p0) ? ((!(_this.f3)) || (_this.f3)) : ((_893_v96).IsProperSubsetOf(_893_v96)));
          let _894_v97;
          let _init19 = function (_895_i11) {
            return _dafny.Seq.Concat((_this).f2, _dafny.Seq.UnicodeFromString("icurqkgo"));
          };
          let _nw114 = Array((new BigNumber(9)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw114.length); _i0_19++) {
            _nw114[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _894_v97 = _nw114;
          let _index137 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_894_v97).length));
          (_894_v97)[_index137] = (_this).f2;
          let _index138 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_894_v97).length));
          (_894_v97)[_index138] = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f2, _dafny.Seq.UnicodeFromString("hvoau")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), function (_896_i12) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }));
        } else {
          let _index139 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index139] = true;
          let _index140 = _module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index140] = p0;
          let _897_v98;
          let _nw115 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _897_v98 = _nw115;
          let _898_v99;
          _898_v99 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_module.__default.fm0(globalState));
          let _899_v100;
          _899_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_898_v99).length),_897_v98);
          let _900_v101;
          _900_v101 = _dafny.Seq.of(_807_v31);
          let _901_v102;
          let _nw116 = Array((new BigNumber(2)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _897_v98;
          _nw116[(_dafny.ONE).toNumber()] = (((_899_v100).contains(new BigNumber((_900_v101).length))) ? ((_899_v100).get(new BigNumber((_900_v101).length))) : (_897_v98));
          _901_v102 = _nw116;
          let _index141 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_901_v102).length));
          (_901_v102)[_index141] = _897_v98;
          let _index142 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_901_v102).length));
          (_901_v102)[_index142] = _897_v98;
          let _902_v103;
          _902_v103 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f4, (_881_v89)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length))]),(_this).f2);
          _902_v103 = _902_v103;
          let _903_v104;
          _903_v104 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm3((_881_v89)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length))], globalState)),((_this).f5)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f5).length))]);
          _903_v104 = (_903_v104).update(p0, ((_this).f5)[_module.__default.safeIndex(new BigNumber(310), new BigNumber(((_this).f5).length))]);
          let _index143 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_881_v89).length));
          (_881_v89)[_index143] = (_this).f4;
        }
      }
      if (false) {
        let _904_v105;
        _904_v105 = new _dafny.CodePoint('g'.codePointAt(0));
        let _905_v106;
        let _nw117 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _905_v106 = _nw117;
        let _906_v107;
        _906_v107 = _dafny.Map.Empty.slice().updateUnsafe(_904_v105,_905_v106);
        let _907_v108;
        _907_v108 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_906_v107).contains(_904_v105)) ? ((_906_v107).get(_904_v105)) : (_905_v106)));
        let _908_v109;
        let _nw118 = new _module.C0();
        _nw118.__ctor((((_907_v108).contains(_this.f3)) ? ((_907_v108).get(_this.f3)) : (_905_v106)));
        _908_v109 = _nw118;
        let _909_v110;
        _909_v110 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_807_v31).length),(_this).f4);
        let _910_v111;
        _910_v111 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_this).f2).length));
        let _911_v112;
        _911_v112 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_910_v111).length),p0);
        let _912_v113;
        _912_v113 = _module.D6.create_DC16(_911_v112);
        let _913_v114;
        _913_v114 = _dafny.Map.Empty.slice().updateUnsafe(_912_v113,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),(_this).f4));
        let _914_v115;
        _914_v115 = _dafny.Set.fromElements(_909_v110, (_909_v110).Merge(_909_v110), (((_913_v114).contains(_912_v113)) ? ((_913_v114).get(_912_v113)) : (_909_v110)));
        _914_v115 = _914_v115;
        let _915_v116;
        _915_v116 = _module.D3.create_DC9((_this).f4, p0);
        let _source16 = function (_pat_let32_0) {
          return function (_916_dt__update__tmp_h2) {
            return function (_pat_let33_0) {
              return function (_917_dt__update_hcf16_h0) {
                return _module.D3.create_DC9((_916_dt__update__tmp_h2).dtor_cf15, _917_dt__update_hcf16_h0);
              }(_pat_let33_0);
            }(!(_this.f3));
          }(_pat_let32_0);
        }(_915_v116);
        if (_source16.is_DC9) {
          let _918___mcc_h2 = (_source16).cf15;
          let _919___mcc_h3 = (_source16).cf16;
          let _920_cf16 = _919___mcc_h3;
          let _921_cf15 = _918___mcc_h2;
          let _index144 = _module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index144] = (new BigNumber(573)).isLessThanOrEqualTo(_921_cf15);
          let _922_v117;
          _922_v117 = _dafny.MultiSet.fromElements(_920_cf16, p0, _920_cf16);
          let _index145 = _module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index145] = !(_922_v117).equals(_922_v117);
          let _923_v118;
          _923_v118 = _dafny.Set.fromElements(_904_v105);
          let _924_v119;
          _924_v119 = _module.D6.create_DC17(_this.f3);
          let _925_v120;
          _925_v120 = _dafny.Map.Empty.slice().updateUnsafe((_923_v118).IsSubsetOf(_923_v118),_924_v119);
          _925_v120 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (_920_cf16) : ((_this).fm11(((_this).f5)[_module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length))], globalState))),_924_v119);
          let _index146 = _module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length));
          let _rhs135 = _this.f3;
          let _rhs136 = ((p0) ? ((_dafny.ZERO).minus(_module.__default.fm0(globalState))) : (_921_cf15));
          let _lhs74 = (_this).f5;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length));
          _lhs74[_lhs75] = _rhs135;
          _921_cf15 = _rhs136;
          let _926_v121;
          let _init20 = ((_927_v105) => function (_928_i13) {
            return _927_v105;
          })(_904_v105);
          let _nw119 = Array((_dafny.ONE).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw119.length); _i0_20++) {
            _nw119[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _926_v121 = _nw119;
          let _index147 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_926_v121).length));
          (_926_v121)[_index147] = _904_v105;
          let _929_v122;
          _929_v122 = _dafny.Map.Empty.slice().updateUnsafe(p0,((_this).f5)[_module.__default.safeIndex(new BigNumber(210), new BigNumber(((_this).f5).length))]);
          let _930_v123;
          _930_v123 = _module.D8.create_DC24(_904_v105, _929_v122);
          let _pat_let_tv28 = _929_v122;
          let _index148 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_926_v121).length));
          (_926_v121)[_index148] = (function (_pat_let34_0) {
            return function (_931_dt__update__tmp_h3) {
              return function (_pat_let35_0) {
                return function (_932_dt__update_hcf42_h0) {
                  return _module.D8.create_DC24((_931_dt__update__tmp_h3).dtor_cf41, _932_dt__update_hcf42_h0);
                }(_pat_let35_0);
              }(_pat_let_tv28);
            }(_pat_let34_0);
          }(_930_v123)).dtor_cf41;
        } else if (_source16.is_DC10) {
          let _933___mcc_h4 = (_source16).cf17;
          let _934___mcc_h5 = (_source16).cf18;
          let _935_cf18 = _934___mcc_h5;
          let _936_cf17 = _933___mcc_h4;
          (_this).f3 = _module.__default.fm3((_dafny.ZERO).minus((_dafny.ZERO).minus(_936_cf17)), globalState);
          let _index149 = _module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index149] = !(p0);
          let _index150 = _module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index150] = p0;
          let _index151 = _module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index151] = !(_this.f3);
          let _937_v124;
          _937_v124 = _dafny.Seq.of(_this.f3);
          let _938_v125;
          _938_v125 = _dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))]);
          let _939_v126;
          _939_v126 = _module.D2.create_DC6(_904_v105, _this.f3, p0, (_this).f4);
          let _940_v127;
          _940_v127 = _module.D2.create_DC6((_939_v126).dtor_cf9, ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))], p0, new BigNumber(379));
          let _941_v128;
          let _nw120 = Array((new BigNumber(24)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = (new BigNumber((_807_v31).length)).isLessThan(_936_cf17);
          _nw120[(_dafny.ONE).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))];
          _nw120[(new BigNumber(2)).toNumber()] = _this.f3;
          _nw120[(new BigNumber(3)).toNumber()] = (((_911_v112).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_936_cf17)))) ? ((_911_v112).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_936_cf17)))) : (p0));
          _nw120[(new BigNumber(4)).toNumber()] = _this.f3;
          _nw120[(new BigNumber(5)).toNumber()] = (_937_v124)[_module.__default.safeIndex((_this).f4, new BigNumber((_937_v124).length))];
          _nw120[(new BigNumber(6)).toNumber()] = !(!(!(_938_v125).equals(_938_v125)));
          _nw120[(new BigNumber(7)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))];
          _nw120[(new BigNumber(8)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))];
          _nw120[(new BigNumber(9)).toNumber()] = (((_911_v112).contains((_this).f4)) ? ((_911_v112).get((_this).f4)) : (p0));
          _nw120[(new BigNumber(10)).toNumber()] = false;
          _nw120[(new BigNumber(11)).toNumber()] = _this.f3;
          _nw120[(new BigNumber(12)).toNumber()] = (_940_v127).dtor_cf10;
          _nw120[(new BigNumber(13)).toNumber()] = true;
          _nw120[(new BigNumber(14)).toNumber()] = p0;
          _nw120[(new BigNumber(15)).toNumber()] = p0;
          _nw120[(new BigNumber(16)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))];
          _nw120[(new BigNumber(17)).toNumber()] = !_dafny.areEqual((_this).f2, (_this).f2);
          _nw120[(new BigNumber(18)).toNumber()] = p0;
          _nw120[(new BigNumber(19)).toNumber()] = p0;
          _nw120[(new BigNumber(20)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(519), new BigNumber(((_this).f5).length))];
          _nw120[(new BigNumber(21)).toNumber()] = (((_937_v124)[_module.__default.safeIndex((_this).f4, new BigNumber((_937_v124).length))]) ? (_this.f3) : (false));
          _nw120[(new BigNumber(22)).toNumber()] = p0;
          _nw120[(new BigNumber(23)).toNumber()] = ((_this).f4).isEqualTo(new BigNumber(321));
          _941_v128 = _nw120;
          _941_v128 = _941_v128;
        } else if (_source16.is_DC8) {
          let _942___mcc_h6 = (_source16).cf14;
          let _943_cf14 = _942___mcc_h6;
          let _944_v129;
          _944_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,p0);
          let _945_v130;
          let _946_v131;
          let _947_v132;
          let _948_v133;
          let _out17;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector4 = _module.__default.m0(_944_v129, globalState);
          _out17 = _outcollector4[0];
          _out18 = _outcollector4[1];
          _out19 = _outcollector4[2];
          _out20 = _outcollector4[3];
          _945_v130 = _out17;
          _946_v131 = _out18;
          _947_v132 = _out19;
          _948_v133 = _out20;
          let _949_v134;
          let _nw121 = new _module.C1();
          _nw121.__ctor();
          _949_v134 = _nw121;
          let _950_v135;
          let _nw122 = new _module.C0();
          _nw122.__ctor(_905_v106);
          _950_v135 = _nw122;
          let _index152 = _module.__default.safeIndex(new BigNumber(19), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index152] = (_module.__default.fm0(globalState)).plus((_this).f4);
          let _index153 = _module.__default.safeIndex(new BigNumber(19), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index153] = (((_this).f4).multipliedBy((_this).f4)).plus((_this).f4);
        } else {
          let _951___mcc_h7 = (_source16).cf19;
          let _952_cf19 = _951___mcc_h7;
          let _953_v136;
          _953_v136 = _dafny.Seq.of(p0, p0);
          (_this).f3 = (_module.D6.create_DC18(new BigNumber(334), (_dafny.ZERO).minus(new BigNumber((_953_v136).length)), p0, _904_v105, (_this).f4)).dtor_cf31;
          let _954_v137;
          _954_v137 = _dafny.Set.fromElements(_this.f3);
          let _955_v138;
          _955_v138 = _dafny.Seq.of(_954_v137, (_954_v137).Difference(_954_v137));
          let _956_v139;
          _956_v139 = _dafny.Map.Empty.slice().updateUnsafe(p0,_954_v137);
          _955_v138 = _dafny.Seq.of(((((_956_v139).contains(_this.f3)) ? ((_956_v139).get(_this.f3)) : (_954_v137))).Intersect(_954_v137));
          (_this).f3 = p0;
          let _957_v140;
          _957_v140 = new BigNumber(-13);
          let _958_v141;
          _958_v141 = _module.D2.create_DC5(_910_v111);
          let _959_v142;
          _959_v142 = _module.D2.create_DC7(_958_v141);
          _957_v140 = ((_957_v140).minus((_this).f4)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_959_v142, _959_v142, _959_v142, _959_v142, _959_v142)).length)));
        }
        let _960_v143;
        _960_v143 = _dafny.Set.fromElements(p0);
        let _961_v144;
        _961_v144 = _module.D3.create_DC8(_960_v143);
        let _source17 = _961_v144;
        if (_source17.is_DC9) {
          let _962___mcc_h8 = (_source17).cf15;
          let _963___mcc_h9 = (_source17).cf16;
          let _964_cf16 = _963___mcc_h9;
          let _965_cf15 = _962___mcc_h8;
          _965_cf15 = _965_cf15;
          let _966_v145;
          _966_v145 = _module.D2.create_DC5(_910_v111);
          let _967_v146;
          _967_v146 = _dafny.Map.Empty.slice().updateUnsafe((_908_v109).f14,((p0) ? (_966_v145) : (_module.D2.create_DC5(_910_v111))));
          _967_v146 = (_967_v146).update((_908_v109).f14, _966_v145);
          _965_cf15 = _965_cf15;
          _904_v105 = _904_v105;
        } else if (_source17.is_DC10) {
          let _968___mcc_h10 = (_source17).cf17;
          let _969___mcc_h11 = (_source17).cf18;
          let _970_cf18 = _969___mcc_h11;
          let _971_cf17 = _968___mcc_h10;
          r1 = _module.__default.fm3((_this).f4, globalState);
          _971_cf17 = (_dafny.ZERO).minus(_971_cf17);
          (_this).f3 = p0;
          let _972_v147;
          _972_v147 = _module.D10.create_DC28(_909_v110);
          _970_cf18 = (_dafny.ZERO).minus(((_970_cf18).plus((_this).f4)).minus((_dafny.ZERO).minus(new BigNumber((((_this.f3) ? (_909_v110) : ((_972_v147).dtor_cf46))).length))));
        } else if (_source17.is_DC8) {
          let _973___mcc_h12 = (_source17).cf14;
          let _974_cf14 = _973___mcc_h12;
          let _975_v148;
          _975_v148 = _dafny.Set.fromElements((_this).f4, (_this).f4, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-725)), ((_976_v110) => function (_977_i14) {
            return new BigNumber((_976_v110).length);
          })(_909_v110))).length));
          let _978_v149;
          _978_v149 = _module.D6.create_DC19(((_this.f3) ? (new BigNumber((_975_v148).length)) : ((_this).f4)));
          _978_v149 = _module.D6.create_DC19((_this).f4);
          let _979_v150;
          let _nw123 = Array((new BigNumber(8)).toNumber()).fill(_module.D4.Default());
          _979_v150 = _nw123;
          let _980_v151;
          _980_v151 = _module.D9.create_DC25(_979_v150);
          let _981_v152;
          _981_v152 = new BigNumber(740);
          let _rhs137 = true;
          let _rhs138 = !(false);
          let _rhs139 = (_this).fm13(globalState);
          let _rhs140 = _module.D9.create_DC25(_979_v150);
          let _rhs141 = (_dafny.ZERO).minus((_this).f4);
          let _lhs76 = _this;
          r1 = _rhs137;
          r1 = _rhs138;
          _lhs76.f3 = _rhs139;
          _980_v151 = _rhs140;
          _981_v152 = _rhs141;
          let _index154 = _module.__default.safeIndex(new BigNumber(604), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index154] = (_dafny.ZERO).minus(_981_v152);
          let _index155 = _module.__default.safeIndex(new BigNumber(604), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index155] = (((_dafny.ZERO).minus((_this).f4)).plus(new BigNumber((_807_v31).length))).plus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3)).length), _981_v152));
          let _982_v153;
          _982_v153 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f3) ? (_dafny.Seq.UnicodeFromString("f")) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), ((_983_v105) => function (_984_i15) {
            return _983_v105;
          })(_904_v105)), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), ((_985_v105) => function (_986_i15) {
            return _985_v105;
          })(_904_v105))).length)), _904_v105))),(_this).f2);
          _982_v153 = (_982_v153).update((_this).f2, (_this).f2);
        } else {
          let _987___mcc_h13 = (_source17).cf19;
          let _988_cf19 = _987___mcc_h13;
          let _989_v154;
          _989_v154 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("cckx"),(_this).f4);
          let _index156 = _module.__default.safeIndex(new BigNumber(913), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index156] = ((((_989_v154).contains(_dafny.Seq.UnicodeFromString("rn"))) ? ((_989_v154).get(_dafny.Seq.UnicodeFromString("rn"))) : ((_this).f4))).minus((_this).f4);
          let _index157 = _module.__default.safeIndex(new BigNumber(913), new BigNumber(((_908_v109).f14).length));
          ((_908_v109).f14)[_index157] = ((_this).f4).plus((_this).f4);
          r1 = (_this).fm11(p0, globalState);
          let _990_v155;
          _990_v155 = _dafny.Seq.UnicodeFromString("ndc");
          _990_v155 = _dafny.Seq.Concat(_dafny.Seq.Concat(_990_v155, (_this).f2), _dafny.Seq.Concat((_this).f2, _module.__default.fm2(p0, globalState)));
          let _991_v156;
          let _992_v157;
          let _993_v158;
          let _994_v159;
          let _out21;
          let _out22;
          let _out23;
          let _out24;
          let _outcollector5 = _module.__default.m0(_dafny.Map.Empty.slice().updateUnsafe((_this).f5,p0), globalState);
          _out21 = _outcollector5[0];
          _out22 = _outcollector5[1];
          _out23 = _outcollector5[2];
          _out24 = _outcollector5[3];
          _991_v156 = _out21;
          _992_v157 = _out22;
          _993_v158 = _out23;
          _994_v159 = _out24;
        }
        let _995_v160;
        _995_v160 = new BigNumber(811);
        _995_v160 = (_this).fm12((_this).f4, new BigNumber(-323), _995_v160, _this.f3, globalState);
      } else {
        let _996_v161;
        _996_v161 = new BigNumber(356);
        let _997_v162;
        _997_v162 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f4);
        let _998_v163;
        _998_v163 = _dafny.Map.Empty.slice().updateUnsafe(_996_v161,_997_v162);
        let _rhs142 = _module.__default.fm0(globalState);
        let _rhs143 = (_996_v161).minus(new BigNumber(((((_998_v163).contains((_this).f4)) ? ((_998_v163).get((_this).f4)) : (_997_v162))).length));
        let _rhs144 = ((_this).f4).plus(_996_v161);
        _996_v161 = _rhs142;
        _996_v161 = _rhs143;
        _996_v161 = _rhs144;
        let _999_v164;
        _999_v164 = _dafny.Seq.UnicodeFromString("soixpidow");
        let _rhs145 = _999_v164;
        let _rhs146 = (_this).f4;
        _999_v164 = _rhs145;
        _996_v161 = _rhs146;
        (_this).f3 = _this.f3;
        let _1000_v165;
        _1000_v165 = _dafny.Seq.of(p0);
        (_this).f3 = _dafny.Seq.IsPrefixOf(_1000_v165, _dafny.Seq.Concat(_1000_v165, _dafny.Seq.of(p0, !(p0))));
        let _1001_v166;
        let _init21 = function (_1002_i16) {
          return (_1002_i16).plus(new BigNumber(-944));
        };
        let _nw124 = Array((new BigNumber(8)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw124.length); _i0_21++) {
          _nw124[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _1001_v166 = _nw124;
        let _index158 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1001_v166).length));
        (_1001_v166)[_index158] = (new BigNumber(453)).plus((_this).f4);
        let _index159 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1001_v166).length));
        (_1001_v166)[_index159] = (_this).f4;
      }
      let _1003_v167;
      _1003_v167 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,false);
      let _1004_v168;
      _1004_v168 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1003_v167);
      let _1005_v169;
      let _nw125 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _1005_v169 = _nw125;
      let _1006_v170;
      _1006_v170 = _dafny.Map.Empty.slice().updateUnsafe((((_1004_v168).contains(new BigNumber(((_this).f2).length))) ? ((_1004_v168).get(new BigNumber(((_this).f2).length))) : (_dafny.Map.Empty.slice().updateUnsafe(!(_this.f3),true))),_1005_v169);
      let _1007_v171;
      _1007_v171 = _dafny.Seq.of(_1003_v167);
      _1006_v170 = (_1006_v170).update((_1003_v167).Merge((_1007_v171)[_module.__default.safeIndex((_this).f4, new BigNumber((_1007_v171).length))]), _1005_v169);
      let _1008_v172;
      _1008_v172 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,p0);
      let _1009_v173;
      _1009_v173 = _dafny.Seq.of(_module.D6.create_DC17(_this.f3));
      let _1010_v174;
      _1010_v174 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1009_v173).length),new BigNumber(868));
      (_this).f3 = (((_1008_v172).contains(new BigNumber((_1010_v174).length))) ? ((_1008_v172).get(new BigNumber((_1010_v174).length))) : (_this.f3));
      let _1011_v175;
      _1011_v175 = _dafny.Set.fromElements((_this).f5, (_this).f5);
      let _1012_v176;
      _1012_v176 = _dafny.Map.Empty.slice().updateUnsafe((_1011_v175).Difference(_1011_v175),_module.__default.safeModuloInt(new BigNumber((_1003_v167).length), _module.__default.fm0(globalState)));
      _1012_v176 = (_1012_v176).update(_1011_v175, (new BigNumber((_1010_v174).length)).multipliedBy((_this).f4));
      let _1013_v177;
      _1013_v177 = _module.D5.create_DC14(_1005_v169);
      r0 = _1013_v177;
      let _1014_v178;
      _1014_v178 = _dafny.Set.fromElements((_this).fm12((_807_v31)[_module.__default.safeIndex((_this).f4, new BigNumber((_807_v31).length))], new BigNumber((_dafny.Seq.UnicodeFromString("vhfa")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(867)), function (_1015_i17) {
        return (_this).f4;
      })).length), _this.f3, globalState), (_this).f4, (_this).f4, (_dafny.ZERO).minus((_this).f4));
      let _1016_v179;
      _1016_v179 = _dafny.Map.Empty.slice().updateUnsafe(_1014_v178,_1011_v175);
      let _1017_v180;
      _1017_v180 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),(((_1016_v179).contains(_1014_v178)) ? ((_1016_v179).get(_1014_v178)) : (_dafny.Set.fromElements((_this).f5))));
      r1 = (((((_1017_v180).contains((_this).f4)) ? ((_1017_v180).get((_this).f4)) : (_1011_v175))).Intersect(_1011_v175)).equals(_1011_v175);
      return [r0, r1];
    }
    m11(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      r0 = _module.__default.fm0(globalState);
      let _1018_i0;
      _1018_i0 = _dafny.ZERO;
      L7: {
        while (_this.f3) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1018_i0)) {
              break L7;
            }
            _1018_i0 = (_1018_i0).plus(_dafny.ONE);
            let _1019_v0;
            _1019_v0 = _dafny.Seq.of(p1, (_this).f4);
            let _1020_v1;
            _1020_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_this.f3, _this.f3),(_this).f2);
            let _1021_v3;
            _1021_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_dafny.ZERO).minus(new BigNumber((_1019_v0).length)));
            let _1022_v4;
            _1022_v4 = _dafny.Seq.of(_module.__default.fm3(new BigNumber(544), globalState));
            let _1023_v5;
            let _nw126 = Array((new BigNumber(19)).toNumber());
            _nw126[(_dafny.ZERO).toNumber()] = p1;
            _nw126[(_dafny.ONE).toNumber()] = (_1019_v0)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_1019_v0).length))];
            _nw126[(new BigNumber(2)).toNumber()] = (_this).f4;
            _nw126[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p1);
            _nw126[(new BigNumber(4)).toNumber()] = ((_this.f3) ? ((_dafny.ZERO).minus(p1)) : (new BigNumber((_1019_v0).length)));
            _nw126[(new BigNumber(5)).toNumber()] = new BigNumber((_1019_v0).length);
            _nw126[(new BigNumber(6)).toNumber()] = p1;
            _nw126[(new BigNumber(7)).toNumber()] = (new BigNumber(784)).plus(p1);
            _nw126[(new BigNumber(8)).toNumber()] = (_this).f4;
            _nw126[(new BigNumber(9)).toNumber()] = p1;
            _nw126[(new BigNumber(10)).toNumber()] = (_this).f4;
            _nw126[(new BigNumber(11)).toNumber()] = new BigNumber((function () {
              let _coll37 = new _dafny.Set();
              for (const _compr_37 of (_1020_v1).Keys.Elements) {
                let _1024_v2 = _compr_37;
                if ((_1020_v1).contains(_1024_v2)) {
                  _coll37.add(_1024_v2);
                }
              }
              return _coll37;
            }()).length);
            _nw126[(new BigNumber(12)).toNumber()] = (((_1021_v3).contains((_this).f2)) ? ((_1021_v3).get((_this).f2)) : (_module.__default.fm0(globalState)));
            _nw126[(new BigNumber(13)).toNumber()] = p1;
            _nw126[(new BigNumber(14)).toNumber()] = (p1).minus(p1);
            _nw126[(new BigNumber(15)).toNumber()] = new BigNumber((_1022_v4).length);
            _nw126[(new BigNumber(16)).toNumber()] = (_this).f4;
            _nw126[(new BigNumber(17)).toNumber()] = (_this).f4;
            _nw126[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(p1, (_1019_v0)[_module.__default.safeIndex(p1, new BigNumber((_1019_v0).length))]);
            _1023_v5 = _nw126;
            let _index160 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1023_v5).length));
            (_1023_v5)[_index160] = (_this).f4;
            let _index161 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_1023_v5).length));
            (_1023_v5)[_index161] = (((_dafny.ZERO).minus((_this).f4)).multipliedBy(p1)).multipliedBy(new BigNumber(193));
            let _1025_v6;
            _1025_v6 = _dafny.MultiSet.fromElements((_1023_v5)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_1023_v5).length))], (_this).f4);
            (_this).f3 = (_1025_v6).contains((_module.__default.fm35(_this.f3, globalState)).dtor_cf1);
            (_this).f3 = _this.f3;
            r0 = (_1023_v5)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_1023_v5).length))];
          }
        }
      }
      let _1026_v7;
      _1026_v7 = new _dafny.CodePoint('i'.codePointAt(0));
      _1026_v7 = ((!(_this.f3)) ? (_1026_v7) : (_1026_v7));
      let _1027_v8;
      _1027_v8 = _dafny.MultiSet.fromElements(p1, (_this).f4, p1, (_this).f4);
      let _1028_v9;
      _1028_v9 = _module.D3.create_DC9(new BigNumber((_1027_v8).cardinality()), true);
      let _1029_v10;
      _1029_v10 = _module.D6.create_DC19((_this).f4);
      let _1030_v11;
      _1030_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v9,(_1029_v10).dtor_cf34);
      let _source18 = _module.__default.fm36((_1030_v11).Merge(_1030_v11), globalState);
      if (_source18.is_DC17) {
        let _1031___mcc_h0 = (_source18).cf28;
        let _1032_cf28 = _1031___mcc_h0;
        let _1033_v12;
        let _nw127 = Array((new BigNumber(2)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(p1);
        _nw127[(_dafny.ONE).toNumber()] = _1027_v8;
        _1033_v12 = _nw127;
        let _1034_v13;
        _1034_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v12,new BigNumber(267));
        _1034_v13 = (_1034_v13).update(_1033_v12, (_this).f4);
        let _1035_v14;
        let _nw128 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1035_v14 = _nw128;
        let _1036_v15;
        _1036_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1035_v14);
        let _1037_v16;
        _1037_v16 = _dafny.Seq.of(_1035_v14, _1035_v14, _1035_v14, _1035_v14, _1035_v14);
        let _1038_v17;
        let _nw129 = Array((new BigNumber(19)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = _1035_v14;
        _nw129[(_dafny.ONE).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(2)).toNumber()] = (((_1036_v15).contains((_1028_v9).dtor_cf16)) ? ((_1036_v15).get((_1028_v9).dtor_cf16)) : (_1035_v14));
        _nw129[(new BigNumber(3)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(4)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(5)).toNumber()] = (_1037_v16)[_module.__default.safeIndex((_this).f4, new BigNumber((_1037_v16).length))];
        _nw129[(new BigNumber(6)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(7)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(8)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(9)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(10)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(11)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(12)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(13)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(14)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(15)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(16)).toNumber()] = _1035_v14;
        _nw129[(new BigNumber(17)).toNumber()] = (((_1036_v15).contains(!(false))) ? ((_1036_v15).get(!(false))) : (_1035_v14));
        _nw129[(new BigNumber(18)).toNumber()] = _1035_v14;
        _1038_v17 = _nw129;
        let _index162 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1038_v17).length));
        (_1038_v17)[_index162] = _1035_v14;
        let _index163 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1038_v17).length));
        let _rhs147 = _1035_v14;
        let _rhs148 = (_this).f4;
        let _rhs149 = _1032_cf28;
        let _lhs77 = _1038_v17;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1038_v17).length));
        _lhs77[_lhs78] = _rhs147;
        r0 = _rhs148;
        _1032_cf28 = _rhs149;
        r0 = new BigNumber(911);
        let _1039_v18;
        _1039_v18 = _dafny.Seq.UnicodeFromString("orqcpbsg");
        _1039_v18 = (_this).f2;
      } else if (_source18.is_DC18) {
        let _1040___mcc_h1 = (_source18).cf29;
        let _1041___mcc_h2 = (_source18).cf30;
        let _1042___mcc_h3 = (_source18).cf31;
        let _1043___mcc_h4 = (_source18).cf32;
        let _1044___mcc_h5 = (_source18).cf33;
        let _1045_cf33 = _1044___mcc_h5;
        let _1046_cf32 = _1043___mcc_h4;
        let _1047_cf31 = _1042___mcc_h3;
        let _1048_cf30 = _1041___mcc_h2;
        let _1049_cf29 = _1040___mcc_h1;
        let _1050_v19;
        _1050_v19 = _dafny.Seq.of((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), ((_1051_v7) => function (_1052_i1) {
          return _1051_v7;
        })(_1026_v7)), _dafny.Seq.UnicodeFromString("saoqwp"), (_this).f2, _dafny.Seq.UnicodeFromString("ao"));
        _1047_cf31 = _dafny.Seq.contains((_1050_v19)[_module.__default.safeIndex((_this).f4, new BigNumber((_1050_v19).length))], _1026_v7);
        let _1053_v20;
        _1053_v20 = _dafny.MultiSet.fromElements(_1047_cf31, _1047_cf31);
        _1026_v7 = (((_1053_v20).IsDisjointFrom(_dafny.MultiSet.fromElements(_this.f3))) ? (_1026_v7) : (_1026_v7));
        let _1054_v21;
        _1054_v21 = _dafny.Seq.UnicodeFromString("dm");
        _1054_v21 = _dafny.Seq.UnicodeFromString("iljigkhs");
        let _1055_v22;
        _1055_v22 = _dafny.Seq.of(!(!(_this.f3)), _1047_cf31, _1047_cf31);
        (_this).f3 = _dafny.areEqual(_1054_v21, _dafny.Seq.Concat(_1054_v21, _dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber((_1055_v22).length), new BigNumber(((_this).f2).length)), _1026_v7)));
      } else if (_source18.is_DC19) {
        let _1056___mcc_h6 = (_source18).cf34;
        let _1057_cf34 = _1056___mcc_h6;
        let _1058_v23;
        _1058_v23 = _dafny.Seq.of(_this.f3, _this.f3);
        if ((_1058_v23)[_module.__default.safeIndex(new BigNumber(-324), new BigNumber((_1058_v23).length))]) {
          (_this).f3 = _this.f3;
          (_this).f3 = (_this.f3) === (_this.f3);
          let _1059_v24;
          _1059_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_this).f2, _module.__default.safeIndex(_1057_cf34, new BigNumber(((_this).f2).length)), ((_this).f2)[_module.__default.safeIndex((_this).f4, new BigNumber(((_this).f2).length))]),(_this).f4);
          _1059_v24 = (_1059_v24).update((_this).f2, (_this).f4);
          let _1060_v25;
          _1060_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_1058_v23)[_module.__default.safeIndex(p1, new BigNumber((_1058_v23).length))]);
          _1060_v25 = (_1060_v25).update(_this.f3, false);
          let _1061_v26;
          _1061_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1060_v25,_this.f3);
          _1061_v26 = (_1061_v26).update(_1060_v25, _this.f3);
        } else {
          let _1062_v27;
          _1062_v27 = _dafny.Seq.UnicodeFromString("brwwhg");
          _1062_v27 = _dafny.Seq.update(_module.__default.fm2(_this.f3, globalState), _module.__default.safeIndex(_1057_cf34, new BigNumber((_module.__default.fm2(_this.f3, globalState)).length)), _1026_v7);
          let _1063_v28;
          let _nw130 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1063_v28 = _nw130;
          let _1064_v29;
          let _nw131 = new _module.C0();
          _nw131.__ctor(_1063_v28);
          _1064_v29 = _nw131;
          let _index164 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length));
          ((_1064_v29).f14)[_index164] = (_this).f4;
          let _1065_v30;
          _1065_v30 = _module.D1.create_DC2(_1062_v27);
          let _index165 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length));
          let _rhs150 = _module.__default.safeDivisionInt((_this).f4, p1);
          let _rhs151 = _this.f3;
          let _rhs152 = _dafny.Seq.IsProperPrefixOf(_1062_v27, (_1065_v30).dtor_cf2);
          let _rhs153 = (_this).f4;
          let _lhs79 = _this;
          let _lhs80 = _this;
          let _lhs81 = (_1064_v29).f14;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length));
          r0 = _rhs150;
          _lhs79.f3 = _rhs151;
          _lhs80.f3 = _rhs152;
          _lhs81[_lhs82] = _rhs153;
          let _1066_v31;
          _1066_v31 = _dafny.Seq.of(p1, _1057_cf34, (_this).f4);
          let _1067_v32;
          _1067_v32 = _module.D4.create_DC13(_this.f3, _dafny.Seq.Concat(_1066_v31, _1066_v31));
          _1067_v32 = _1067_v32;
          let _1068_v33;
          _1068_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f2);
          let _1069_v34;
          _1069_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1027_v8).cardinality()),_1062_v27);
          let _1070_v36;
          _1070_v36 = _dafny.Set.fromElements(true, !(_this.f3), _this.f3);
          let _1071_v37;
          _1071_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1057_cf34,_1069_v34);
          let _1072_v40;
          _1072_v40 = _module.D11.create_DC30(_1069_v34);
          let _1073_v41;
          _1073_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1069_v34);
          let _1074_v43;
          _1074_v43 = _dafny.Set.fromElements(p1);
          let _1075_v44;
          let _nw132 = Array((new BigNumber(25)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f2).length),(((_1068_v33).contains(_this.f3)) ? ((_1068_v33).get(_this.f3)) : ((_this).f2)))).update((_this).f4, (_this).f2);
          _nw132[(_dafny.ONE).toNumber()] = (_1069_v34).Merge(function () {
            let _coll38 = new _dafny.Map();
            for (const _compr_38 of (_1066_v31).Elements) {
              let _1076_v35 = _compr_38;
              if (_dafny.Seq.contains(_1066_v31, _1076_v35)) {
                _coll38.push([(_1076_v35).plus((_1029_v10).dtor_cf34),_dafny.Seq.UnicodeFromString("ivjrwehm")]);
              }
            }
            return _coll38;
          }());
          _nw132[(new BigNumber(2)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(3)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(4)).toNumber()] = _module.__default.fm37(((_1064_v29).f14)[_module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length))], p1, ((_1064_v29).f14)[_module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length))], _1070_v36, globalState);
          _nw132[(new BigNumber(5)).toNumber()] = (_1069_v34).update(new BigNumber((_dafny.MultiSet.fromElements(_1057_cf34)).cardinality()), _1062_v27);
          _nw132[(new BigNumber(6)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(7)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.UnicodeFromString("kqqsxqjwf"));
          _nw132[(new BigNumber(9)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(10)).toNumber()] = ((((_1071_v37).contains((_this).f4)) ? ((_1071_v37).get((_this).f4)) : (_1069_v34))).Merge(_1069_v34);
          _nw132[(new BigNumber(11)).toNumber()] = (_1069_v34).Merge(_1069_v34);
          _nw132[(new BigNumber(12)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(13)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(14)).toNumber()] = (_1069_v34).Merge(_1069_v34);
          _nw132[(new BigNumber(15)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(16)).toNumber()] = function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of _dafny.IntegerRange(new BigNumber(232), new BigNumber(565))) {
              let _1077_v38 = _compr_39;
              if (((new BigNumber(232)).isLessThanOrEqualTo(_1077_v38)) && ((_1077_v38).isLessThan(new BigNumber(565)))) {
                _coll39.push([(_1077_v38).minus(new BigNumber((function () {
                  let _coll40 = new _dafny.Set();
                  for (const _compr_40 of _dafny.IntegerRange(new BigNumber(83), new BigNumber(585))) {
                    let _1078_v39 = _compr_40;
                    if (((new BigNumber(83)).isLessThanOrEqualTo(_1078_v39)) && ((_1078_v39).isLessThan(new BigNumber(585)))) {
                      _coll40.add(_module.__default.safeModuloInt(_1078_v39, ((_1064_v29).f14)[_module.__default.safeIndex(new BigNumber(576), new BigNumber(((_1064_v29).f14).length))]));
                    }
                  }
                  return _coll40;
                }()).length)),_dafny.Seq.UnicodeFromString("rhbt")]);
              }
            }
            return _coll39;
          }();
          _nw132[(new BigNumber(17)).toNumber()] = (_1069_v34).update(p1, (_this).f2);
          _nw132[(new BigNumber(18)).toNumber()] = (_1069_v34).update((_this).fm12(_1057_cf34, _1057_cf34, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), ((_1079_v7) => function (_1080_i2) {
            return _1079_v7;
          })(_1026_v7))).length), _this.f3, globalState), _dafny.Seq.UnicodeFromString("ekvspg"));
          _nw132[(new BigNumber(19)).toNumber()] = ((_1072_v40).dtor_cf49).Merge((((_1073_v41).contains(_this.f3)) ? ((_1073_v41).get(_this.f3)) : (_1069_v34)));
          _nw132[(new BigNumber(20)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(21)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(22)).toNumber()] = _1069_v34;
          _nw132[(new BigNumber(23)).toNumber()] = (function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_1074_v43).Elements) {
              let _1081_v42 = _compr_41;
              if ((_1074_v43).contains(_1081_v42)) {
                _coll41.push([_module.__default.safeModuloInt(_1081_v42, p1),(_this).f2]);
              }
            }
            return _coll41;
          }()).Merge(_1069_v34);
          _nw132[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f2);
          _1075_v44 = _nw132;
          let _index166 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_1075_v44).length));
          (_1075_v44)[_index166] = (_1069_v34).update(p1, _dafny.Seq.UnicodeFromString("onhe"));
          let _1082_v45;
          let _nw133 = Array((new BigNumber(25)).toNumber()).fill(_module.D10.Default());
          _1082_v45 = _nw133;
          let _1083_v46;
          _1083_v46 = _dafny.MultiSet.fromElements(false);
          let _index167 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_1075_v44).length));
          let _rhs154 = new BigNumber(-907);
          let _rhs155 = _1069_v34;
          let _rhs156 = (_this).f2;
          let _rhs157 = _1082_v45;
          let _rhs158 = (_1083_v46).Intersect((_1083_v46).Intersect(_dafny.MultiSet.fromElements(_this.f3, _this.f3)));
          let _lhs83 = _1075_v44;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_1075_v44).length));
          r0 = _rhs154;
          _lhs83[_lhs84] = _rhs155;
          _1062_v27 = _rhs156;
          _1082_v45 = _rhs157;
          _1083_v46 = _rhs158;
        }
        let _1084_v47;
        _1084_v47 = _dafny.Set.fromElements(_1026_v7, _1026_v7);
        let _index168 = _module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index168] = (_1084_v47).IsSubsetOf(_1084_v47);
        let _index169 = _module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index169] = (_this.f3) === (_this.f3);
        let _1085_v49;
        _1085_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1057_cf34,((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))]);
        let _1086_v50;
        _1086_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of _dafny.IntegerRange(new BigNumber(117), new BigNumber(-55))) {
            let _1087_v48 = _compr_42;
            if (((new BigNumber(117)).isLessThanOrEqualTo(_1087_v48)) && ((_1087_v48).isLessThan(new BigNumber(-55)))) {
              _coll42.push([(_1087_v48).minus(p1),true]);
            }
          }
          return _coll42;
        }()).Merge(_1085_v49)).length),p1);
        let _1088_v51;
        _1088_v51 = _module.D8.create_DC23(_1057_cf34, p1, ((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))]);
        _1086_v50 = (_1086_v50).update(_1057_cf34, ((((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))]) ? ((_this).f4) : ((_1088_v51).dtor_cf38)));
        let _1089_v52;
        _1089_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(541),_dafny.Seq.UnicodeFromString("sxbfhfglq"));
        let _1090_v53;
        _1090_v53 = _module.D1.create_DC3(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], new BigNumber(((((_1089_v52).contains(p1)) ? ((_1089_v52).get(p1)) : ((_this).f2))).length), new BigNumber((_dafny.Seq.Concat((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-886)), function (_1091_i3) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}))).length));
        let _source19 = _1090_v53;
        if (_source19.is_DC3) {
          let _1092___mcc_h9 = (_source19).cf3;
          let _1093___mcc_h10 = (_source19).cf4;
          let _1094___mcc_h11 = (_source19).cf5;
          let _1095_cf5 = _1094___mcc_h11;
          let _1096_cf4 = _1093___mcc_h10;
          let _1097_cf3 = _1092___mcc_h9;
          let _1098_v54;
          _1098_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1095_cf5);
          let _1099_v55;
          _1099_v55 = _dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))]);
          let _index170 = _module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index170] = ((((_1098_v54).contains((_this).fm26((_this).f4, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], globalState))) ? ((_1098_v54).get((_this).fm26((_this).f4, _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], globalState))) : (new BigNumber((_1099_v55).length)))).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)));
          let _1100_v56;
          _1100_v56 = _dafny.Seq.of((_this).f4);
          _1086_v50 = (_1086_v50).update((_1057_cf34).minus(new BigNumber((_dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))])).length)), _module.__default.safeModuloInt((_1100_v56)[_module.__default.safeIndex((_this).f4, new BigNumber((_1100_v56).length))], (_dafny.ZERO).minus((_this).f4)));
          let _1101_v57;
          _1101_v57 = _dafny.Set.fromElements(_1057_cf34, _1057_cf34);
          r0 = new BigNumber((((_dafny.Set.fromElements(_1096_cf4)).Union(_1101_v57)).Intersect(_1101_v57)).length);
          let _1102_v58;
          _1102_v58 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_1103_v7) => function (_1104_i4) {
            return _1103_v7;
          })(_1026_v7)));
          let _index171 = _module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index171] = !(_1102_v58).contains((_this).f2);
        } else if (_source19.is_DC4) {
          let _1105___mcc_h12 = (_source19).cf6;
          let _1106___mcc_h13 = (_source19).cf7;
          let _1107_cf7 = _1106___mcc_h13;
          let _1108_cf6 = _1105___mcc_h12;
          let _1109_v59;
          _1109_v59 = _dafny.Seq.of(new BigNumber(((_this).f2).length));
          _1108_cf6 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1109_v59, _1109_v59), _1109_v59));
          r0 = _module.__default.safeModuloInt(p1, new BigNumber((_1027_v8).cardinality()));
          (_this).f3 = !(_this.f3);
          let _1110_v60;
          let _nw134 = new _module.C1();
          _nw134.__ctor();
          _1110_v60 = _nw134;
        } else {
          let _1111___mcc_h14 = (_source19).cf2;
          let _1112_cf2 = _1111___mcc_h14;
          let _1113_v61;
          _1113_v61 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f4).multipliedBy(_1057_cf34),_module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))])).length),_1112_cf2)));
          let _1114_v62;
          _1114_v62 = _module.D11.create_DC30(_1089_v52);
          _1113_v61 = (_1113_v61).update(p1, _1114_v62);
          let _1115_v63;
          let _nw135 = Array((new BigNumber(26)).toNumber()).fill([]);
          _1115_v63 = _nw135;
          let _1116_v64;
          let _init22 = ((_1117_v7) => function (_1118_i5) {
            return _1117_v7;
          })(_1026_v7);
          let _nw136 = Array((new BigNumber(19)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw136.length); _i0_22++) {
            _nw136[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1116_v64 = _nw136;
          let _index172 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_1115_v63).length));
          (_1115_v63)[_index172] = _1116_v64;
          let _index173 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_1115_v63).length));
          (_1115_v63)[_index173] = _1116_v64;
          let _1119_v65;
          _1119_v65 = _dafny.Seq.of((_this).f4, (_this).f4);
          let _1120_v66;
          let _nw137 = Array((new BigNumber(13)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1120_v66 = _nw137;
          let _1121_v67;
          _1121_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))],_1119_v65),_1120_v66);
          let _1122_v68;
          _1122_v68 = _dafny.Map.Empty.slice().updateUnsafe(false,_1119_v65);
          _1121_v67 = (_1121_v67).update(((_1122_v68).update(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], _1119_v65)).Merge((_1122_v68).update(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], _1119_v65)), _1120_v66);
          _1119_v65 = _dafny.Seq.of((_this).f4, _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))], _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(217), new BigNumber(((_this).f5).length))])).length), (_this).f4));
        }
      } else if (_source18.is_DC16) {
        let _1123___mcc_h7 = (_source18).cf27;
        let _1124_cf27 = _1123___mcc_h7;
        (_this).f3 = !(!(!(_this.f3)));
        let _1125_v69;
        let _nw138 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1125_v69 = _nw138;
        let _1126_v70;
        _1126_v70 = _dafny.Set.fromElements(_this.f3);
        let _1127_v71;
        _1127_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1026_v7,_1125_v69);
        let _rhs159 = !((_1126_v70).Intersect(_1126_v70)).contains(_this.f3);
        let _rhs160 = (((_1127_v71).contains(_1026_v7)) ? ((_1127_v71).get(_1026_v7)) : (_1125_v69));
        let _rhs161 = _1026_v7;
        let _rhs162 = (_module.__default.safeModuloInt(p1, new BigNumber((_1027_v8).cardinality()))).plus((_dafny.ZERO).minus(new BigNumber(((_this).f2).length)));
        let _lhs85 = _this;
        _lhs85.f3 = _rhs159;
        _1125_v69 = _rhs160;
        _1026_v7 = _rhs161;
        r0 = _rhs162;
        let _1128_v72;
        let _nw139 = new _module.C2();
        _nw139.__ctor(p1, _this.f3, _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3), (_this).f2);
        _1128_v72 = _nw139;
        let _1129_v73;
        _1129_v73 = _module.D6.create_DC17(_this.f3);
        let _index174 = _module.__default.safeIndex(new BigNumber(282), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index174] = (_1129_v73).dtor_cf28;
        let _index175 = _module.__default.safeIndex(new BigNumber(282), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index175] = _this.f3;
      } else {
        let _1130___mcc_h8 = (_source18).cf35;
        let _1131_cf35 = _1130___mcc_h8;
        let _1132_v74;
        _1132_v74 = _dafny.MultiSet.fromElements((_this).f2);
        if ((_1132_v74).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f2))) {
          let _1133_v75;
          _1133_v75 = _dafny.Seq.of(_module.__default.safeDivisionInt(p1, p1));
          _1133_v75 = _dafny.Seq.Concat(_1133_v75, _1133_v75);
          let _1134_v76;
          _1134_v76 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
          _1134_v76 = (_1134_v76).update((_this).f4, _this.f3);
          let _1135_v77;
          let _init23 = ((_1136_p1) => function (_1137_i6) {
            return _module.__default.safeDivisionInt(_1137_i6, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_this.f3)).length),_1136_p1)).length));
          })(p1);
          let _nw140 = Array((new BigNumber(5)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw140.length); _i0_23++) {
            _nw140[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1135_v77 = _nw140;
          let _1138_v78;
          let _nw141 = new _module.C0();
          _nw141.__ctor(_1135_v77);
          _1138_v78 = _nw141;
          let _1139_v79;
          _1139_v79 = _dafny.Seq.UnicodeFromString("xui");
          let _rhs163 = _1138_v78;
          let _rhs164 = _1139_v79;
          let _rhs165 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4));
          _1138_v78 = _rhs163;
          _1139_v79 = _rhs164;
          r0 = _rhs165;
          let _1140_v80;
          _1140_v80 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3));
          let _1141_v81;
          let _nw142 = new _module.C2();
          _nw142.__ctor(p1, _this.f3, (_1140_v80)[_module.__default.safeIndex((_this).f4, new BigNumber((_1140_v80).length))], (_this).f2);
          _1141_v81 = _nw142;
          let _1142_v82;
          let _nw143 = new _module.C1();
          _nw143.__ctor();
          _1142_v82 = _nw143;
        } else {
          let _1143_v83;
          _1143_v83 = _module.D4.create_DC13(_this.f3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(738)), function (_1144_i7) {
  return (_this).f4;
}));
          let _1145_v84;
          _1145_v84 = _dafny.Seq.of((_this).f4);
          let _1146_v85;
          _1146_v85 = _dafny.Set.fromElements(((!(_this.f3)) ? (_1143_v83) : (_1143_v83)), _module.D4.create_DC13(_this.f3, _1145_v84));
          let _1147_v86;
          _1147_v86 = _dafny.Map.Empty.slice().updateUnsafe((_module.D6.create_DC18(p1, (_this).f4, _this.f3, _1026_v7, (_dafny.ZERO).minus(p1))).dtor_cf31,_1146_v85);
          _1146_v85 = (((_1147_v86).contains(_this.f3)) ? ((_1147_v86).get(_this.f3)) : ((_1146_v85).Intersect(_1146_v85)));
          let _rhs166 = (p1).isLessThanOrEqualTo(p1);
          let _rhs167 = _this.f3;
          let _rhs168 = (_this).fm12((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f4, (_dafny.ZERO).minus(new BigNumber(((_this).fm14(_this.f3, _this.f3, _1026_v7, new BigNumber(((_this).f2).length), globalState)).length)))), (_this).f4, new BigNumber((_dafny.MultiSet.fromElements(!(_this.f3))).cardinality()), _this.f3, globalState);
          let _lhs86 = _this;
          let _lhs87 = _this;
          _lhs86.f3 = _rhs166;
          _lhs87.f3 = _rhs167;
          r0 = _rhs168;
          let _1148_v87;
          let _nw144 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1148_v87 = _nw144;
          let _1149_v88;
          let _nw145 = new _module.C0();
          _nw145.__ctor(_1148_v87);
          _1149_v88 = _nw145;
          (_this).f3 = _this.f3;
          let _1150_v89;
          let _init24 = ((_1151_v7) => function (_1152_i8) {
            return _1151_v7;
          })(_1026_v7);
          let _nw146 = Array((new BigNumber(24)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw146.length); _i0_24++) {
            _nw146[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1150_v89 = _nw146;
          let _1153_v90;
          _1153_v90 = _dafny.MultiSet.fromElements(_1150_v89);
          (_this).f3 = ((_1153_v90).update(_1150_v89, _module.__default.abs((_this).f4))).IsDisjointFrom(_1153_v90);
        }
        r0 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), ((_1154_v7) => function (_1155_i9) {
          return _1154_v7;
        })(_1026_v7))), _module.__default.safeIndex((_this).f4, new BigNumber((_dafny.Seq.Concat((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), ((_1156_v7) => function (_1157_i9) {
          return _1156_v7;
        })(_1026_v7)))).length)), _1026_v7)).length));
        r0 = (_this).f4;
        if ((_this.f3) === (_this.f3)) {
          let _1158_v91;
          _1158_v91 = _dafny.Seq.of((_this).f4);
          let _1159_v92;
          let _nw147 = Array((new BigNumber(24)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = _1026_v7;
          _nw147[(_dafny.ONE).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(2)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(3)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(4)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(5)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(6)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(7)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(8)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(9)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(10)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(11)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(12)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(13)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(14)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(15)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(16)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(17)).toNumber()] = _module.__default.fm38(_this.f3, false, (_this).f4, globalState);
          _nw147[(new BigNumber(18)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(19)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(20)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(21)).toNumber()] = _1026_v7;
          _nw147[(new BigNumber(22)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
          _nw147[(new BigNumber(23)).toNumber()] = _1026_v7;
          _1159_v92 = _nw147;
          let _1160_v93;
          _1160_v93 = _dafny.Seq.of(_1159_v92, _1159_v92);
          let _1161_v94;
          _1161_v94 = _module.D4.create_DC13(false, _dafny.Seq.update(_1158_v91, _module.__default.safeIndex((_this).f4, new BigNumber((_1158_v91).length)), new BigNumber((_1160_v93).length)));
          let _1162_v95;
          let _nw148 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1162_v95 = _nw148;
          let _1163_v96;
          let _nw149 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1163_v96 = _nw149;
          let _index176 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1162_v95).length));
          (_1162_v95)[_index176] = _1163_v96;
          let _1164_v97;
          _1164_v97 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber(((_this).f2).length));
          let _index177 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1162_v95).length));
          let _rhs169 = _module.D4.create_DC13(false, _1158_v91);
          let _rhs170 = _this.f3;
          let _rhs171 = ((_this.f3) ? (p1) : ((((_1164_v97).contains(!(_this.f3))) ? ((_1164_v97).get(!(_this.f3))) : ((_this).f4))));
          let _rhs172 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), ((_1165_v7) => function (_1166_i10) {
            return _1165_v7;
          })(_1026_v7)), _1026_v7);
          let _rhs173 = _1163_v96;
          let _lhs88 = _this;
          let _lhs89 = _this;
          let _lhs90 = _1162_v95;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1162_v95).length));
          _1161_v94 = _rhs169;
          _lhs88.f3 = _rhs170;
          r0 = _rhs171;
          _lhs89.f3 = _rhs172;
          _lhs90[_lhs91] = _rhs173;
          let _1167_v98;
          let _nw150 = Array((new BigNumber(27)).toNumber()).fill(false);
          _1167_v98 = _nw150;
          let _index178 = _module.__default.safeIndex(new BigNumber(568), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index178] = !(_1027_v8).equals((_1027_v8).update(p1, _module.__default.abs((_this).f4)));
          let _index179 = _module.__default.safeIndex(new BigNumber(568), new BigNumber(((_this).f5).length));
          let _rhs174 = (_this).f4;
          let _rhs175 = _module.__default.fm0(globalState);
          let _rhs176 = _1167_v98;
          let _rhs177 = !((_this).f4).isEqualTo((_this).f4);
          let _rhs178 = _this.f3;
          let _lhs92 = _this;
          let _lhs93 = (_this).f5;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(568), new BigNumber(((_this).f5).length));
          r0 = _rhs174;
          r0 = _rhs175;
          _1167_v98 = _rhs176;
          _lhs92.f3 = _rhs177;
          _lhs93[_lhs94] = _rhs178;
          (_this).f3 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(568), new BigNumber(((_this).f5).length))];
          r0 = (((_1027_v8).contains((_this).f4)) ? ((_1027_v8).get((_this).f4)) : (p1));
          let _index180 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1162_v95).length));
          (_1162_v95)[_index180] = (_1162_v95)[_module.__default.safeIndex(new BigNumber(752), new BigNumber((_1162_v95).length))];
        } else {
          let _1168_v99;
          _1168_v99 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), ((_1169_v7) => function (_1170_i11) {
            return _1169_v7;
          })(_1026_v7)),p1);
          _1168_v99 = (_1168_v99).update((_this).f2, new BigNumber(-930));
          (_this).f3 = _this.f3;
          let _1171_v100;
          _1171_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,!(_this.f3));
          let _1172_v101;
          _1172_v101 = _dafny.Seq.of(_this.f3);
          let _1173_v102;
          _1173_v102 = _dafny.Seq.of(new BigNumber((_1172_v101).length));
          let _1174_v103;
          let _nw151 = Array((new BigNumber(15)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = new BigNumber((_1171_v100).length);
          _nw151[(_dafny.ONE).toNumber()] = p1;
          _nw151[(new BigNumber(2)).toNumber()] = p1;
          _nw151[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1173_v102).length), new BigNumber(991));
          _nw151[(new BigNumber(4)).toNumber()] = p1;
          _nw151[(new BigNumber(5)).toNumber()] = p1;
          _nw151[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, (_dafny.ZERO).minus((_this).f4)));
          _nw151[(new BigNumber(7)).toNumber()] = (_this).f4;
          _nw151[(new BigNumber(8)).toNumber()] = ((_this).f4).plus(new BigNumber(-855));
          _nw151[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt((_this).f4, p1);
          _nw151[(new BigNumber(10)).toNumber()] = (_module.__default.fm39(globalState)).dtor_cf15;
          _nw151[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus(p1)).plus(p1);
          _nw151[(new BigNumber(12)).toNumber()] = (_this).f4;
          _nw151[(new BigNumber(13)).toNumber()] = p1;
          _nw151[(new BigNumber(14)).toNumber()] = (_this).f4;
          _1174_v103 = _nw151;
          let _index181 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1174_v103).length));
          (_1174_v103)[_index181] = ((_this).f4).multipliedBy(new BigNumber((_module.__default.fm2(_this.f3, globalState)).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1174_v103).length));
          (_1174_v103)[_index182] = p1;
          let _1175_v104;
          _1175_v104 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
          let _1176_v105;
          _1176_v105 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let36_0) {
            return function (_1177_dt__update__tmp_h0) {
              return function (_pat_let37_0) {
                return function (_1178_dt__update_hcf25_h0) {
                  return _module.D5.create_DC15((_1177_dt__update__tmp_h0).dtor_cf24, _1178_dt__update_hcf25_h0, (_1177_dt__update__tmp_h0).dtor_cf26);
                }(_pat_let37_0);
              }(_this.f3);
            }(_pat_let36_0);
          }(_module.D5.create_DC15((_this).f4, true, (_this).f5))).dtor_cf25,_1175_v104);
          let _1179_v106;
          _1179_v106 = _dafny.Set.fromElements(_this.f3);
          let _1180_v107;
          _1180_v107 = _module.D6.create_DC18((_this).f4, new BigNumber((_1179_v106).length), _this.f3, _1026_v7, p1);
          let _1181_v108;
          _1181_v108 = _dafny.MultiSet.fromElements(true, (_1180_v107).dtor_cf31);
          let _1182_v109;
          _1182_v109 = _1172_v101;
          let _1183_v110;
          _1183_v110 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1176_v105);
          let _1184_v111;
          _1184_v111 = _dafny.Seq.of((((_1183_v110).contains(_this.f3)) ? ((_1183_v110).get(_this.f3)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1175_v104))));
          let _rhs179 = _this.f3;
          let _rhs180 = ((_dafny.MultiSet.FromArray((_1182_v109))).Difference(_1181_v108)).IsProperSubsetOf(_1181_v108);
          let _rhs181 = ((_this.f3) ? ((_1184_v111)[_module.__default.safeIndex((((_1027_v8).contains((_this).f4)) ? ((_1027_v8).get((_this).f4)) : (new BigNumber(662))), new BigNumber((_1184_v111).length))]) : (_1176_v105));
          let _lhs95 = _this;
          let _lhs96 = _this;
          _lhs95.f3 = _rhs179;
          _lhs96.f3 = _rhs180;
          _1176_v105 = _rhs181;
          let _1185_v112;
          _1185_v112 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1173_v102, _1173_v102),(_this.f3) === (!(_this.f3)));
          _1185_v112 = (_1185_v112).update(_dafny.Seq.Concat(_dafny.Seq.of((_1174_v103)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1174_v103).length))]), _1173_v102), _this.f3);
        }
      }
      (_this).f3 = _this.f3;
      let _1186_v113;
      _1186_v113 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p1);
      let _1187_v114;
      _1187_v114 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(((_1186_v113).contains(_this.f3)) ? ((_1186_v113).get(_this.f3)) : (p1)));
      let _pat_let_tv29 = p1;
      let _source20 = function (_source21) {
        if (_source21.is_DC6) {
          let _1188___mcc_h18 = (_source21).cf9;
          let _1189___mcc_h19 = (_source21).cf10;
          let _1190___mcc_h20 = (_source21).cf11;
          let _1191___mcc_h21 = (_source21).cf12;
          let _1192_cf12 = _1191___mcc_h21;
          let _1193_cf11 = _1190___mcc_h20;
          let _1194_cf10 = _1189___mcc_h19;
          let _1195_cf9 = _1188___mcc_h18;
          return _module.D9.create_DC27(_1192_cf12);
        } else if (_source21.is_DC5) {
          let _1196___mcc_h22 = (_source21).cf8;
          let _1197_cf8 = _1196___mcc_h22;
          return _module.D9.create_DC27(new BigNumber((_dafny.Set.fromElements(_this.f3, _this.f3)).length));
        } else {
          let _1198___mcc_h23 = (_source21).cf13;
          let _1199_cf13 = _1198___mcc_h23;
          return _module.D9.create_DC27(_pat_let_tv29);
        }
      }(_module.D2.create_DC5(_1186_v113));
      if (_source20.is_DC26) {
        let _1200___mcc_h15 = (_source20).cf44;
        let _1201_cf44 = _1200___mcc_h15;
        let _1202_v115;
        let _nw152 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
        _1202_v115 = _nw152;
        let _1203_v116;
        _1203_v116 = _dafny.Seq.of(_1202_v115, _1202_v115, _1202_v115, _1202_v115);
        _1202_v115 = (_1203_v116)[_module.__default.safeIndex(p1, new BigNumber((_1203_v116).length))];
        if (!(_1201_cf44).isEqualTo((_dafny.ZERO).minus((p1).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState)))))) {
          let _1204_v117;
          _1204_v117 = _module.D3.create_DC9(p1, (_1028_v9).dtor_cf16);
          let _1205_v118;
          _1205_v118 = _module.D3.create_DC11(_1204_v117);
          let _pat_let_tv30 = _1204_v117;
          _1205_v118 = function (_pat_let38_0) {
            return function (_1206_dt__update__tmp_h1) {
              return function (_pat_let39_0) {
                return function (_1207_dt__update_hcf19_h0) {
                  return _module.D3.create_DC11(_1207_dt__update_hcf19_h0);
                }(_pat_let39_0);
              }(_pat_let_tv30);
            }(_pat_let38_0);
          }(_module.__default.fm40(new BigNumber((_1027_v8).cardinality()), _this.f3, new BigNumber(281), p1, globalState));
          let _1208_v119;
          _1208_v119 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(((_this).f2).length)));
          let _1209_v120;
          let _nw153 = new _module.C2();
          _nw153.__ctor(p1, !(!(new BigNumber(401)).isEqualTo((_1208_v119)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), ((_1210_v7) => function (_1211_i12) {
            return _1210_v7;
          })(_1026_v7))).length), new BigNumber((_1208_v119).length))])), (_this).f1, (_this).f2);
          _1209_v120 = _nw153;
          let _1212_v121;
          _1212_v121 = _dafny.Set.fromElements(_1201_cf44, (_1201_cf44).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_1209_v120).f13))), _1201_cf44);
          _1212_v121 = _1212_v121;
          r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1201_cf44), (_this).f4);
          let _1213_v122;
          _1213_v122 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
          let _1214_v123;
          _1214_v123 = _module.D6.create_DC16(_1213_v122);
          let _1215_v124;
          _1215_v124 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC20(_1214_v123),(_dafny.ZERO).minus((_this).f4));
          let _1216_v125;
          _1216_v125 = _dafny.MultiSet.fromElements(_1215_v124);
          let _1217_v126;
          _1217_v126 = _module.D6.create_DC20(_1214_v123);
          let _1218_v127;
          _1218_v127 = _dafny.Seq.of(_1215_v124, _dafny.Map.Empty.slice().updateUnsafe(_1217_v126,(_this).f4), _1215_v124, _1215_v124, _1215_v124);
          let _1219_v128;
          _1219_v128 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1218_v127);
          let _1220_v129;
          let _nw154 = Array((new BigNumber(9)).toNumber());
          _nw154[(_dafny.ZERO).toNumber()] = _1216_v125;
          _nw154[(_dafny.ONE).toNumber()] = _1216_v125;
          _nw154[(new BigNumber(2)).toNumber()] = _1216_v125;
          _nw154[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray((((_1219_v128).contains(_this.f3)) ? ((_1219_v128).get(_this.f3)) : (_1218_v127)));
          _nw154[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_1215_v124, _1215_v124));
          _nw154[(new BigNumber(5)).toNumber()] = _module.__default.fm41((_this).f4, _1201_cf44, globalState);
          _nw154[(new BigNumber(6)).toNumber()] = (_1216_v125).update(_1215_v124, _module.__default.abs((_1209_v120).f13));
          _nw154[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.FromArray(_1218_v127)).Intersect(_1216_v125);
          _nw154[(new BigNumber(8)).toNumber()] = _1216_v125;
          _1220_v129 = _nw154;
          let _index183 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1220_v129).length));
          (_1220_v129)[_index183] = _1216_v125;
          let _index184 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1220_v129).length));
          (_1220_v129)[_index184] = _1216_v125;
        } else {
          (_this).f3 = (new BigNumber(76)).isLessThan((_this).f4);
          _1026_v7 = new _dafny.CodePoint('y'.codePointAt(0));
          (_this).f3 = false;
          let _1221_v130;
          _1221_v130 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
          let _1222_v131;
          let _1223_v132;
          let _1224_v133;
          let _1225_v134;
          let _out25;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector6 = _module.__default.m0((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3)).Merge((_1221_v130).update((_this).f5, _this.f3)), globalState);
          _out25 = _outcollector6[0];
          _out26 = _outcollector6[1];
          _out27 = _outcollector6[2];
          _out28 = _outcollector6[3];
          _1222_v131 = _out25;
          _1223_v132 = _out26;
          _1224_v133 = _out27;
          _1225_v134 = _out28;
          r0 = _module.__default.safeModuloInt((_1029_v10).dtor_cf34, (_dafny.ZERO).minus(new BigNumber(((_this).f2).length)));
        }
        _1201_cf44 = _module.__default.fm0(globalState);
        _1201_cf44 = (_this).f4;
      } else if (_source20.is_DC27) {
        let _1226___mcc_h16 = (_source20).cf45;
        let _1227_cf45 = _1226___mcc_h16;
        let _1228_v135;
        _1228_v135 = _dafny.Seq.of((_this).f4);
        let _1229_v136;
        _1229_v136 = _dafny.Set.fromElements(_1228_v135);
        _1229_v136 = (((_this.f3) ? (_1229_v136) : (_1229_v136))).Union(_1229_v136);
        let _1230_v137;
        _1230_v137 = _dafny.Seq.UnicodeFromString("gv");
        _1230_v137 = _dafny.Seq.UnicodeFromString("pceontkto");
        let _1231_v138;
        _1231_v138 = _module.D11.create_DC31((_this).f4);
        let _pat_let_tv31 = p1;
        let _source22 = function (_pat_let40_0) {
          return function (_1232_dt__update__tmp_h2) {
            return function (_pat_let41_0) {
              return function (_1233_dt__update_hcf50_h0) {
                return _module.D11.create_DC31(_1233_dt__update_hcf50_h0);
              }(_pat_let41_0);
            }(_pat_let_tv31);
          }(_pat_let40_0);
        }(_1231_v138);
        if (_source22.is_DC31) {
          let _1234___mcc_h24 = (_source22).cf50;
          let _1235_cf50 = _1234___mcc_h24;
          (_this).f3 = _this.f3;
          let _1236_v139;
          _1236_v139 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm38(false, _this.f3, (((_1027_v8).contains(new BigNumber((_1187_v114).length))) ? ((_1027_v8).get(new BigNumber((_1187_v114).length))) : ((_this).f4)), globalState),(_this).f4);
          let _1237_v140;
          let _nw155 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1237_v140 = _nw155;
          let _1238_v141;
          let _nw156 = Array((new BigNumber(5)).toNumber()).fill([]);
          _1238_v141 = _nw156;
          let _index185 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_1237_v140).length));
          (_1237_v140)[_index185] = _1238_v141;
          let _1239_v142;
          _1239_v142 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f3) || (true),_1238_v141);
          let _index186 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_1237_v140).length));
          let _rhs182 = (_module.__default.fm42(globalState)).Merge(_1236_v139);
          let _rhs183 = (((_1239_v142).contains(_this.f3)) ? ((_1239_v142).get(_this.f3)) : (_1238_v141));
          let _lhs97 = _1237_v140;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_1237_v140).length));
          _1236_v139 = _rhs182;
          _lhs97[_lhs98] = _rhs183;
          let _1240_v143;
          let _nw157 = new _module.C2();
          _nw157.__ctor(new BigNumber(689), !((_dafny.ZERO).minus(_1235_cf50)).isEqualTo(_1235_cf50), (_this).f1, _1230_v137);
          _1240_v143 = _nw157;
          let _1241_v144;
          _1241_v144 = _dafny.Map.Empty.slice().updateUnsafe(_1026_v7,_dafny.Seq.UnicodeFromString("dlbjot"));
          _1230_v137 = (((_1241_v144).contains(_1026_v7)) ? ((_1241_v144).get(_1026_v7)) : ((_this).f2));
        } else {
          let _1242___mcc_h25 = (_source22).cf49;
          let _1243_cf49 = _1242___mcc_h25;
          let _1244_v145;
          let _nw158 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1244_v145 = _nw158;
          let _index187 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1244_v145).length));
          (_1244_v145)[_index187] = _dafny.Seq.UnicodeFromString("p");
          let _index188 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1244_v145).length));
          (_1244_v145)[_index188] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-345)), ((_1245_v7) => function (_1246_i13) {
            return _1245_v7;
          })(_1026_v7));
          let _1247_v146;
          let _init25 = function (_1248_i14) {
            return _module.__default.safeModuloInt(_1248_i14, (_this).f4);
          };
          let _nw159 = Array((new BigNumber(20)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw159.length); _i0_25++) {
            _nw159[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1247_v146 = _nw159;
          let _index189 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1247_v146).length));
          (_1247_v146)[_index189] = p1;
          let _index190 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_1247_v146).length));
          (_1247_v146)[_index190] = _1227_cf45;
          _1227_cf45 = p1;
          let _1249_v147;
          let _nw160 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
          _1249_v147 = _nw160;
          let _1250_v148;
          _1250_v148 = _dafny.Seq.of(_this.f3, _this.f3, _this.f3, _this.f3, _this.f3);
          let _index191 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1249_v147).length));
          (_1249_v147)[_index191] = _1250_v148;
          let _index192 = _module.__default.safeIndex(new BigNumber(245), new BigNumber((_1249_v147).length));
          (_1249_v147)[_index192] = _module.__default.fm43(globalState);
        }
        _1227_cf45 = p1;
      } else {
        let _1251___mcc_h17 = (_source20).cf43;
        let _1252_cf43 = _1251___mcc_h17;
        let _1253_v149;
        _1253_v149 = _dafny.Set.fromElements(_this.f3);
        if (((_1253_v149).Intersect(_1253_v149)).IsDisjointFrom(((_this.f3) ? (_dafny.Set.fromElements(!(false), _this.f3)) : (_1253_v149)))) {
          let _1254_v150;
          _1254_v150 = _dafny.MultiSet.fromElements(!(_this.f3), _this.f3);
          (_this).f3 = (_dafny.Set.fromElements(_1254_v150)).IsProperSubsetOf(_module.__default.fm44(globalState));
          let _1255_v151;
          _1255_v151 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f2)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber(((_this).f2).length))],((_this.f3) ? (new _dafny.CodePoint('x'.codePointAt(0))) : (_1026_v7)));
          _1255_v151 = (_1255_v151).update(_1026_v7, ((_this.f3) ? (_1026_v7) : (_1026_v7)));
          let _1256_v152;
          _1256_v152 = _dafny.Seq.of(p1);
          let _1257_v153;
          _1257_v153 = _dafny.Seq.of(_1256_v152);
          let _1258_v154;
          _1258_v154 = _module.D9.create_DC26((_this).f4);
          let _1259_v155;
          let _nw161 = Array((new BigNumber(22)).toNumber());
          _nw161[(_dafny.ZERO).toNumber()] = new BigNumber(155);
          _nw161[(_dafny.ONE).toNumber()] = p1;
          _nw161[(new BigNumber(2)).toNumber()] = p1;
          _nw161[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(80),_this.f3)).length);
          _nw161[(new BigNumber(4)).toNumber()] = new BigNumber(681);
          _nw161[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("kbqdrprp")).length);
          _nw161[(new BigNumber(6)).toNumber()] = p1;
          _nw161[(new BigNumber(7)).toNumber()] = p1;
          _nw161[(new BigNumber(8)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(9)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(10)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(11)).toNumber()] = p1;
          _nw161[(new BigNumber(12)).toNumber()] = new BigNumber(161);
          _nw161[(new BigNumber(13)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(14)).toNumber()] = new BigNumber((_1257_v153).length);
          _nw161[(new BigNumber(15)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(16)).toNumber()] = p1;
          _nw161[(new BigNumber(17)).toNumber()] = (_this).f4;
          _nw161[(new BigNumber(18)).toNumber()] = p1;
          _nw161[(new BigNumber(19)).toNumber()] = (_1258_v154).dtor_cf44;
          _nw161[(new BigNumber(20)).toNumber()] = p1;
          _nw161[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.of(_this.f3)).length);
          _1259_v155 = _nw161;
          let _1260_v156;
          _1260_v156 = _dafny.Seq.of(_1259_v155);
          let _rhs184 = ((_this).f4).minus((_this).f4);
          let _rhs185 = !_dafny.areEqual(_1260_v156, _1260_v156);
          let _rhs186 = (p1).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p1));
          let _lhs99 = _this;
          let _lhs100 = _this;
          r0 = _rhs184;
          _lhs99.f3 = _rhs185;
          _lhs100.f3 = _rhs186;
          let _1261_v157;
          _1261_v157 = _dafny.Set.fromElements(p1);
          (_this).f3 = (_1261_v157).IsDisjointFrom(_1261_v157);
          let _1262_v158;
          _1262_v158 = _dafny.Seq.of(_this.f3, _this.f3);
          let _rhs187 = ((_this).f4).isLessThan(_module.__default.safeDivisionInt(p1, new BigNumber((_1262_v158).length)));
          let _rhs188 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f4, (_this).f4), _1256_v152);
          let _lhs101 = _this;
          _lhs101.f3 = _rhs187;
          _1256_v152 = _rhs188;
        } else {
          let _1263_v159;
          let _nw162 = new _module.C2();
          _nw162.__ctor((_dafny.ZERO).minus((_this).f4), _this.f3, (_this).f1, (_this).f2);
          _1263_v159 = _nw162;
          let _1264_v160;
          _1264_v160 = _dafny.MultiSet.fromElements((_this).f2);
          _1264_v160 = _module.__default.fm45((_this).f4, _this.f3, globalState);
          let _rhs189 = _module.__default.safeDivisionInt((_this).f4, _module.__default.safeModuloInt(new BigNumber((_1253_v149).length), p1));
          let _rhs190 = p1;
          let _rhs191 = (_module.D10.create_DC29(_this.f3, _this.f3)).dtor_cf47;
          let _rhs192 = (_module.__default.safeDivisionInt(p1, new BigNumber((_1253_v149).length))).isLessThan((_this).f4);
          let _lhs102 = _this;
          let _lhs103 = _this;
          r0 = _rhs189;
          r0 = _rhs190;
          _lhs102.f3 = _rhs191;
          _lhs103.f3 = _rhs192;
          let _1265_v161;
          _1265_v161 = _module.D8.create_DC23(p1, (_this).f4, _this.f3);
          let _1266_v162;
          _1266_v162 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
          let _1267_v163;
          _1267_v163 = _dafny.MultiSet.fromElements(_this.f3, _this.f3, _this.f3, _this.f3);
          let _1268_v164;
          _1268_v164 = _dafny.Set.fromElements((_this).f4, new BigNumber(951), (_this).f4);
          let _1269_v165;
          _1269_v165 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v164,p1);
          let _1270_v166;
          _1270_v166 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4)));
          let _1271_v167;
          let _nw163 = new _module.C1();
          _nw163.__ctor();
          _1271_v167 = _nw163;
          let _1272_v168;
          _1272_v168 = _dafny.MultiSet.fromElements(_1271_v167, _1271_v167);
          let _pat_let_tv32 = p1;
          let _pat_let_tv33 = _1269_v165;
          let _pat_let_tv34 = p1;
          let _pat_let_tv35 = _1263_v159;
          let _pat_let_tv36 = _1263_v159;
          let _1273_v169;
          let _nw164 = Array((new BigNumber(23)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = _1265_v161;
          _nw164[(_dafny.ONE).toNumber()] = _module.D8.create_DC23(p1, new BigNumber(((_1266_v162).update(p1, _this.f3)).length), _this.f3);
          _nw164[(new BigNumber(2)).toNumber()] = _module.D8.create_DC23(new BigNumber((_1267_v163).cardinality()), p1, _module.__default.fm3(p1, globalState));
          _nw164[(new BigNumber(3)).toNumber()] = _module.D8.create_DC23((_1263_v159).f13, p1, _this.f3);
          _nw164[(new BigNumber(4)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(5)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(6)).toNumber()] = function (_pat_let42_0) {
            return function (_1274_dt__update__tmp_h3) {
              return function (_pat_let43_0) {
                return function (_1275_dt__update_hcf38_h0) {
                  return function (_pat_let44_0) {
                    return function (_1276_dt__update_hcf40_h0) {
                      return _module.D8.create_DC23(_1275_dt__update_hcf38_h0, (_1274_dt__update__tmp_h3).dtor_cf39, _1276_dt__update_hcf40_h0);
                    }(_pat_let44_0);
                  }(_this.f3);
                }(_pat_let43_0);
              }(_pat_let_tv32);
            }(_pat_let42_0);
          }(_1265_v161);
          _nw164[(new BigNumber(7)).toNumber()] = function (_pat_let45_0) {
            return function (_1277_dt__update__tmp_h4) {
              return function (_pat_let46_0) {
                return function (_1278_dt__update_hcf38_h1) {
                  return function (_pat_let47_0) {
                    return function (_1279_dt__update_hcf40_h1) {
                      return _module.D8.create_DC23(_1278_dt__update_hcf38_h1, (_1277_dt__update__tmp_h4).dtor_cf39, _1279_dt__update_hcf40_h1);
                    }(_pat_let47_0);
                  }(_this.f3);
                }(_pat_let46_0);
              }(new BigNumber((_pat_let_tv33).length));
            }(_pat_let45_0);
          }(_1265_v161);
          _nw164[(new BigNumber(8)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(9)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(10)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(11)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(12)).toNumber()] = _module.D8.create_DC23(new BigNumber(714), p1, _this.f3);
          _nw164[(new BigNumber(13)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(14)).toNumber()] = function (_pat_let48_0) {
            return function (_1280_dt__update__tmp_h5) {
              return function (_pat_let49_0) {
                return function (_1281_dt__update_hcf40_h2) {
                  return function (_pat_let50_0) {
                    return function (_1282_dt__update_hcf39_h0) {
                      return _module.D8.create_DC23((_1280_dt__update__tmp_h5).dtor_cf38, _1282_dt__update_hcf39_h0, _1281_dt__update_hcf40_h2);
                    }(_pat_let50_0);
                  }(_pat_let_tv34);
                }(_pat_let49_0);
              }(_this.f3);
            }(_pat_let48_0);
          }(_module.D8.create_DC23(new BigNumber((_dafny.MultiSet.fromElements(_1027_v8)).cardinality()), (_this).f4, _this.f3));
          _nw164[(new BigNumber(15)).toNumber()] = _module.D8.create_DC23(new BigNumber((_1270_v166).length), (_this).f4, _this.f3);
          _nw164[(new BigNumber(16)).toNumber()] = function (_pat_let51_0) {
            return function (_1283_dt__update__tmp_h6) {
              return function (_pat_let52_0) {
                return function (_1284_dt__update_hcf40_h3) {
                  return _module.D8.create_DC23((_1283_dt__update__tmp_h6).dtor_cf38, (_1283_dt__update__tmp_h6).dtor_cf39, _1284_dt__update_hcf40_h3);
                }(_pat_let52_0);
              }(false);
            }(_pat_let51_0);
          }(_1265_v161);
          _nw164[(new BigNumber(17)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(18)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(19)).toNumber()] = _1265_v161;
          _nw164[(new BigNumber(20)).toNumber()] = ((_this.f3) ? (_1265_v161) : (function (_pat_let53_0) {
            return function (_1285_dt__update__tmp_h7) {
              return function (_pat_let54_0) {
                return function (_1286_dt__update_hcf38_h2) {
                  return function (_pat_let55_0) {
                    return function (_1287_dt__update_hcf40_h4) {
                      return _module.D8.create_DC23(_1286_dt__update_hcf38_h2, (_1285_dt__update__tmp_h7).dtor_cf39, _1287_dt__update_hcf40_h4);
                    }(_pat_let55_0);
                  }(_this.f3);
                }(_pat_let54_0);
              }((_pat_let_tv35).f13);
            }(_pat_let53_0);
          }(_1265_v161)));
          _nw164[(new BigNumber(21)).toNumber()] = ((!(_this.f3)) ? (function (_pat_let56_0) {
            return function (_1288_dt__update__tmp_h8) {
              return function (_pat_let57_0) {
                return function (_1289_dt__update_hcf38_h3) {
                  return _module.D8.create_DC23(_1289_dt__update_hcf38_h3, (_1288_dt__update__tmp_h8).dtor_cf39, (_1288_dt__update__tmp_h8).dtor_cf40);
                }(_pat_let57_0);
              }((_pat_let_tv36).f13);
            }(_pat_let56_0);
          }(_1265_v161)) : (_1265_v161));
          _nw164[(new BigNumber(22)).toNumber()] = _module.D8.create_DC23((((_1272_v168).contains(_1271_v167)) ? ((_1272_v168).get(_1271_v167)) : ((_1263_v159).f13)), (_1263_v159).f13, _this.f3);
          _1273_v169 = _nw164;
          let _index193 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1273_v169).length));
          (_1273_v169)[_index193] = _1265_v161;
          let _index194 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_1273_v169).length));
          (_1273_v169)[_index194] = _module.D8.create_DC23((_1263_v159).f13, ((_this.f3) ? (p1) : ((_dafny.ZERO).minus((_1263_v159).f13))), _this.f3);
          let _1290_v170;
          _1290_v170 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_this).f4);
          let _1291_v171;
          _1291_v171 = _dafny.MultiSet.fromElements(_1027_v8, _dafny.MultiSet.fromElements(new BigNumber((_1290_v170).length), (_this).f4));
          _1291_v171 = _module.__default.fm46(globalState);
        }
        (_this).f3 = _this.f3;
        let _1292_v172;
        _1292_v172 = _dafny.Seq.of(!(_this.f3));
        let _1293_v173;
        _1293_v173 = _dafny.Seq.of(_1292_v172);
        r0 = (_module.D9.create_DC27(new BigNumber((_1293_v173).length))).dtor_cf45;
        let _1294_v174;
        let _nw165 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1294_v174 = _nw165;
        let _index195 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1294_v174).length));
        (_1294_v174)[_index195] = new _dafny.CodePoint('l'.codePointAt(0));
        let _1295_v175;
        _1295_v175 = _dafny.MultiSet.fromElements(_1292_v172);
        let _1296_v176;
        _1296_v176 = _module.D11.create_DC31((_dafny.ZERO).minus(new BigNumber((_1295_v175).cardinality())));
        let _1297_v177;
        _1297_v177 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f4));
        let _1298_v178;
        _1298_v178 = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus((_dafny.ZERO).minus(p1)));
        let _1299_v179;
        _1299_v179 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new _dafny.CodePoint('h'.codePointAt(0)));
        let _1300_v181;
        let _nw166 = Array((new BigNumber(12)).toNumber());
        _nw166[(_dafny.ZERO).toNumber()] = p1;
        _nw166[(_dafny.ONE).toNumber()] = (_1297_v177)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f4), new BigNumber((_1297_v177).length))];
        _nw166[(new BigNumber(2)).toNumber()] = p1;
        _nw166[(new BigNumber(3)).toNumber()] = (_this).f4;
        _nw166[(new BigNumber(4)).toNumber()] = p1;
        _nw166[(new BigNumber(5)).toNumber()] = new BigNumber((_1298_v178).length);
        _nw166[(new BigNumber(6)).toNumber()] = new BigNumber((_1299_v179).length);
        _nw166[(new BigNumber(7)).toNumber()] = new BigNumber((_1299_v179).length);
        _nw166[(new BigNumber(8)).toNumber()] = p1;
        _nw166[(new BigNumber(9)).toNumber()] = new BigNumber((_module.__default.fm2(_this.f3, globalState)).length);
        _nw166[(new BigNumber(10)).toNumber()] = new BigNumber((function () {
          let _coll43 = new _dafny.Set();
          for (const _compr_43 of _dafny.IntegerRange(new BigNumber(344), new BigNumber(-998))) {
            let _1301_v180 = _compr_43;
            if (((new BigNumber(344)).isLessThanOrEqualTo(_1301_v180)) && ((_1301_v180).isLessThan(new BigNumber(-998)))) {
              _coll43.add((_1301_v180).minus(p1));
            }
          }
          return _coll43;
        }()).length);
        _nw166[(new BigNumber(11)).toNumber()] = (_this).f4;
        _1300_v181 = _nw166;
        let _1302_v182;
        _1302_v182 = _dafny.Map.Empty.slice().updateUnsafe(_1296_v176,_1300_v181);
        let _1303_v183;
        _1303_v183 = _dafny.Seq.of(_1297_v177);
        let _1304_v184;
        _1304_v184 = _module.D6.create_DC18((_this).f4, (_this).f4, _module.__default.fm3(new BigNumber((_1303_v183).length), globalState), _1026_v7, (_this).f4);
        let _index196 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1294_v174).length));
        let _rhs193 = !(_1302_v182).equals((_dafny.Map.Empty.slice().updateUnsafe(_1296_v176,_1300_v181)).Merge(_1302_v182));
        let _rhs194 = (_1304_v184).dtor_cf32;
        let _rhs195 = _this.f3;
        let _lhs104 = _this;
        let _lhs105 = _1294_v174;
        let _lhs106 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_1294_v174).length));
        let _lhs107 = _this;
        _lhs104.f3 = _rhs193;
        _lhs105[_lhs106] = _rhs194;
        _lhs107.f3 = _rhs195;
      }
      r0 = new BigNumber(349);
      return r0;
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    __ctor(f3, f1, f2) {
      let _this = this;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(642), new BigNumber((((!(_this.f3)) ? ((_this).f2) : (_dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber(356), new BigNumber(((_this).f2).length)), new _dafny.CodePoint('f'.codePointAt(0)))))).length));
    };
    fm11(p0, globalState) {
      let _this = this;
      return !(!(!(_this.f3)) || (_this.f3));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _hi5 = p0;
      for (let _1305_i0 = (_module.D2.create_DC6(_module.__default.fm24(_this.f3, _this.f3, p0, globalState), _this.f3, true, new BigNumber(755))).dtor_cf12; _1305_i0.isLessThan(_hi5); _1305_i0 = _1305_i0.plus(_dafny.ONE)) {
        r2 = new BigNumber(((_this).f2).length);
        let _1306_v0;
        _1306_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
        let _1307_v1;
        _1307_v1 = _dafny.MultiSet.fromElements(_this.f3, _this.f3);
        let _1308_v2;
        _1308_v2 = _dafny.Seq.of(_this.f3);
        let _1309_v3;
        let _nw167 = Array((new BigNumber(29)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = (_this).fm11(_this.f3, globalState);
        _nw167[(_dafny.ONE).toNumber()] = _this.f3;
        _nw167[(new BigNumber(2)).toNumber()] = (_1308_v2)[_module.__default.safeIndex(p0, new BigNumber((_1308_v2).length))];
        _nw167[(new BigNumber(3)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(4)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(5)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(6)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(7)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(8)).toNumber()] = true;
        _nw167[(new BigNumber(9)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(10)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(11)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(12)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(13)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(14)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(15)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(16)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(17)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(18)).toNumber()] = false;
        _nw167[(new BigNumber(19)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(20)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(21)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(22)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(23)).toNumber()] = !(_this.f3);
        _nw167[(new BigNumber(24)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(25)).toNumber()] = (_1308_v2)[_module.__default.safeIndex(p1, new BigNumber((_1308_v2).length))];
        _nw167[(new BigNumber(26)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(27)).toNumber()] = _this.f3;
        _nw167[(new BigNumber(28)).toNumber()] = _this.f3;
        _1309_v3 = _nw167;
        let _1310_v4;
        let _nw168 = new _module.C3();
        _nw168.__ctor(new BigNumber((_1307_v1).cardinality()), _1309_v3, _this.f3, (_this).f1, (_this).f2);
        _1310_v4 = _nw168;
        let _1311_v5;
        _1311_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber((_dafny.MultiSet.fromElements(p1, _module.__default.fm0(globalState))).cardinality()));
        let _1312_v6;
        _1312_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1310_v4,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1311_v5).length),_1310_v4.f3)).length));
        let _1313_v7;
        _1313_v7 = _dafny.Set.fromElements(_this.f3, _this.f3, _1310_v4.f3);
        let _1314_v8;
        let _nw169 = Array((new BigNumber(28)).toNumber());
        _nw169[(_dafny.ZERO).toNumber()] = p0;
        _nw169[(_dafny.ONE).toNumber()] = (new BigNumber((_1306_v0).length)).plus(p0);
        _nw169[(new BigNumber(2)).toNumber()] = new BigNumber(((_this).f2).length);
        _nw169[(new BigNumber(3)).toNumber()] = p1;
        _nw169[(new BigNumber(4)).toNumber()] = _1305_i0;
        _nw169[(new BigNumber(5)).toNumber()] = new BigNumber(911);
        _nw169[(new BigNumber(6)).toNumber()] = p0;
        _nw169[(new BigNumber(7)).toNumber()] = p1;
        _nw169[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber(((_this).f2).length));
        _nw169[(new BigNumber(9)).toNumber()] = _1305_i0;
        _nw169[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw169[(new BigNumber(11)).toNumber()] = p0;
        _nw169[(new BigNumber(12)).toNumber()] = _1305_i0;
        _nw169[(new BigNumber(13)).toNumber()] = new BigNumber(((_this).f2).length);
        _nw169[(new BigNumber(14)).toNumber()] = new BigNumber((_1307_v1).cardinality());
        _nw169[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(((_this).f2).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0)).length));
        _nw169[(new BigNumber(16)).toNumber()] = (p1).minus(new BigNumber(250));
        _nw169[(new BigNumber(17)).toNumber()] = (_this).fm12(new BigNumber((_1312_v6).length), _1305_i0, new BigNumber((_dafny.Seq.UnicodeFromString("x")).length), _1310_v4.f3, globalState);
        _nw169[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt((_1310_v4).f4, (_1310_v4).f4);
        _nw169[(new BigNumber(19)).toNumber()] = p0;
        _nw169[(new BigNumber(20)).toNumber()] = p1;
        _nw169[(new BigNumber(21)).toNumber()] = (_1310_v4).f4;
        _nw169[(new BigNumber(22)).toNumber()] = p0;
        _nw169[(new BigNumber(23)).toNumber()] = (_1310_v4).f4;
        _nw169[(new BigNumber(24)).toNumber()] = ((_1310_v4).fm12(_1305_i0, (_1310_v4).f4, _1305_i0, _1310_v4.f3, globalState)).plus(p1);
        _nw169[(new BigNumber(25)).toNumber()] = (_1310_v4).f4;
        _nw169[(new BigNumber(26)).toNumber()] = new BigNumber(((_1313_v7).Union(_1313_v7)).length);
        _nw169[(new BigNumber(27)).toNumber()] = (((_1307_v1).contains(true)) ? ((_1307_v1).get(true)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), function (_1315_i1) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length)));
        _1314_v8 = _nw169;
        let _index197 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1314_v8).length));
        (_1314_v8)[_index197] = p0;
        let _index198 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1314_v8).length));
        (_1314_v8)[_index198] = ((_1310_v4).f4).minus(p0);
        let _1316_v9;
        let _nw170 = new _module.C3();
        _nw170.__ctor(p1, (_1310_v4).f5, (p0).isEqualTo((_1310_v4).f4), (_this).f1, (_this).f2);
        _1316_v9 = _nw170;
        let _1317_v10;
        _1317_v10 = _dafny.MultiSet.fromElements(_1314_v8);
        (_1310_v4).f3 = (_1317_v10).IsSubsetOf(_1317_v10);
      }
      let _1318_v11;
      let _nw171 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _1318_v11 = _nw171;
      let _1319_v12;
      _1319_v12 = _dafny.Seq.of(_1318_v11);
      _1319_v12 = _dafny.Seq.Concat(_1319_v12, _dafny.Seq.of(_1318_v11));
      (_this).f3 = !(false);
      let _1320_i2;
      _1320_i2 = _dafny.ZERO;
      L8: {
        while (!(_this.f3)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1320_i2)) {
              break L8;
            }
            _1320_i2 = (_1320_i2).plus(_dafny.ONE);
            let _1321_v13;
            let _nw172 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
            _1321_v13 = _nw172;
            let _1322_v14;
            let _nw173 = Array((new BigNumber(15)).toNumber());
            _nw173[(_dafny.ZERO).toNumber()] = _this.f3;
            _nw173[(_dafny.ONE).toNumber()] = true;
            _nw173[(new BigNumber(2)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(3)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(4)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(5)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(6)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(7)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(8)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(9)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(10)).toNumber()] = false;
            _nw173[(new BigNumber(11)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(12)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(13)).toNumber()] = _this.f3;
            _nw173[(new BigNumber(14)).toNumber()] = _this.f3;
            _1322_v14 = _nw173;
            let _1323_v15;
            _1323_v15 = new _dafny.CodePoint('o'.codePointAt(0));
            let _index199 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1321_v13).length));
            (_1321_v13)[_index199] = (_dafny.Map.Empty.slice().updateUnsafe(_1322_v14,_1323_v15)).update(_1322_v14, new _dafny.CodePoint('a'.codePointAt(0)));
            let _1324_v16;
            _1324_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1322_v14,_1323_v15);
            let _index200 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_1321_v13).length));
            (_1321_v13)[_index200] = _1324_v16;
            (_this).f3 = _this.f3;
            let _1325_v17;
            _1325_v17 = _dafny.Seq.of(_this.f3, !(_this.f3), _this.f3);
            let _1326_v18;
            _1326_v18 = _1325_v17;
            let _source23 = _1326_v18;
            let _1327___mcc_h0 = _source23;
            let _1328_cf51 = _1327___mcc_h0;
            let _1329_v19;
            _1329_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
            let _1330_v20;
            _1330_v20 = _dafny.MultiSet.fromElements(p1);
            let _1331_v21;
            _1331_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1330_v20,p1);
            _1329_v19 = (_1329_v19).update(_module.__default.safeModuloInt(p0, new BigNumber((_1331_v21).length)), _this.f3);
            _1328_cf51 = _dafny.Seq.Concat(_dafny.Seq.of(_this.f3, _this.f3), _dafny.Seq.Concat(_dafny.Seq.of(_this.f3, _this.f3), _1325_v17));
            let _1332_v22;
            _1332_v22 = _dafny.MultiSet.fromElements(_this.f3, false);
            let _1333_v23;
            _1333_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1318_v11,_1318_v11);
            let _index201 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1318_v11).length));
            (_1318_v11)[_index201] = p0;
            let _1334_v24;
            _1334_v24 = _dafny.Seq.of(_1333_v23, _1333_v23);
            let _index202 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1318_v11).length));
            let _rhs196 = _1318_v11;
            let _rhs197 = _1332_v22;
            let _rhs198 = (_1334_v24)[_module.__default.safeIndex(p1, new BigNumber((_1334_v24).length))];
            let _rhs199 = (p0).plus(p1);
            let _rhs200 = new BigNumber((_dafny.Seq.UnicodeFromString("wxvovuu")).length);
            let _lhs108 = _1318_v11;
            let _lhs109 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1318_v11).length));
            _1318_v11 = _rhs196;
            _1332_v22 = _rhs197;
            _1333_v23 = _rhs198;
            r1 = _rhs199;
            _lhs108[_lhs109] = _rhs200;
            _1330_v20 = (((_dafny.MultiSet.fromElements(new BigNumber(((_this).f2).length))).update(new BigNumber(((_this).f2).length), _module.__default.abs(p0))).Difference(_1330_v20)).Intersect(_1330_v20);
            let _index203 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1322_v14).length));
            (_1322_v14)[_index203] = _this.f3;
            let _1335_v25;
            _1335_v25 = _dafny.MultiSet.fromElements(_this.f3);
            let _1336_v26;
            _1336_v26 = _dafny.Seq.of(new BigNumber((_1335_v25).cardinality()));
            let _1337_v27;
            _1337_v27 = _dafny.Seq.of(_1336_v26, _1336_v26, _1336_v26);
            let _1338_v28;
            _1338_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), ((_1339_p0) => function (_1340_i3) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_1341_p0) => function (_1342_i4) {
                return _1341_p0;
              })(_1339_p0));
            })(p0)), _1337_v27),!(_this.f3));
            let _index204 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_1322_v14).length));
            (_1322_v14)[_index204] = (((_1338_v28).contains(_dafny.Seq.Concat(_1337_v27, _1337_v27))) ? ((_1338_v28).get(_dafny.Seq.Concat(_1337_v27, _1337_v27))) : (_this.f3));
          }
        }
      }
      let _1343_v29;
      _1343_v29 = new _dafny.CodePoint('w'.codePointAt(0));
      let _1344_v30;
      _1344_v30 = _module.D2.create_DC6(_1343_v29, _this.f3, _this.f3, p1);
      let _source24 = ((_this.f3) ? (_1344_v30) : (_1344_v30));
      if (_source24.is_DC6) {
        let _1345___mcc_h1 = (_source24).cf9;
        let _1346___mcc_h2 = (_source24).cf10;
        let _1347___mcc_h3 = (_source24).cf11;
        let _1348___mcc_h4 = (_source24).cf12;
        let _1349_cf12 = _1348___mcc_h4;
        let _1350_cf11 = _1347___mcc_h3;
        let _1351_cf10 = _1346___mcc_h2;
        let _1352_cf9 = _1345___mcc_h1;
        let _1353_v31;
        _1353_v31 = _dafny.Seq.of((_this).f2, (_this).f2, (_this).f2);
        _1353_v31 = _1353_v31;
        let _1354_v32;
        _1354_v32 = _dafny.MultiSet.fromElements(_1351_cf10, _1351_cf10);
        r2 = (((_1354_v32).contains(_1350_cf11)) ? ((_1354_v32).get(_1350_cf11)) : (new BigNumber(250)));
        let _1355_v33;
        _1355_v33 = _dafny.Set.fromElements(_this.f3, _1350_cf11);
        let _1356_v34;
        _1356_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1350_cf11,p0);
        _1349_cf12 = (new BigNumber((_1355_v33).length)).minus((((_1356_v34).contains(_1350_cf11)) ? ((_1356_v34).get(_1350_cf11)) : (p1)));
        let _1357_v35;
        _1357_v35 = _module.D5.create_DC14((_1319_v12)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_1319_v12).length))]);
        let _source25 = _1357_v35;
        if (_source25.is_DC15) {
          let _1358___mcc_h7 = (_source25).cf24;
          let _1359___mcc_h8 = (_source25).cf25;
          let _1360___mcc_h9 = (_source25).cf26;
          let _1361_cf26 = _1360___mcc_h9;
          let _1362_cf25 = _1359___mcc_h8;
          let _1363_cf24 = _1358___mcc_h7;
          let _1364_v36;
          let _nw174 = Array((new BigNumber(6)).toNumber()).fill(_module.D10.Default());
          _1364_v36 = _nw174;
          let _1365_v37;
          _1365_v37 = _dafny.Seq.of(_1364_v36);
          let _1366_v38;
          _1366_v38 = _dafny.MultiSet.fromElements((_1365_v37)[_module.__default.safeIndex(new BigNumber((_1356_v34).length), new BigNumber((_1365_v37).length))], _1364_v36, _1364_v36, _1364_v36);
          _1366_v38 = _1366_v38;
          _1350_cf11 = (_1351_cf10) === (_this.f3);
          r2 = ((_1350_cf11) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1349_cf12, new BigNumber(292))))) : (_module.__default.safeModuloInt((((_1354_v32).contains(_1351_cf10)) ? ((_1354_v32).get(_1351_cf10)) : (_1363_cf24)), p0)));
          let _index205 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index205] = p1;
          let _1367_v39;
          _1367_v39 = _dafny.Seq.of(_1350_cf11);
          let _1368_v40;
          _1368_v40 = _dafny.Map.Empty.slice().updateUnsafe(((_1351_cf10) ? (_1361_cf26) : (_1361_cf26)),_dafny.Seq.contains(_1367_v39, _this.f3));
          let _1369_v41;
          _1369_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1350_cf11,_this.f3);
          let _1370_v42;
          _1370_v42 = _module.D8.create_DC24(_1343_v29, _1369_v41);
          let _1371_v43;
          _1371_v43 = _dafny.Seq.of(_1370_v42);
          let _1372_v44;
          _1372_v44 = _dafny.Seq.of(p0);
          let _index206 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1318_v11).length));
          let _rhs201 = _1349_cf12;
          let _rhs202 = _1349_cf12;
          let _rhs203 = _1343_v29;
          let _rhs204 = _1368_v40;
          let _rhs205 = !(new BigNumber((_dafny.Seq.Concat(_1371_v43, _module.__default.fm47(new BigNumber((_1372_v44).length), new BigNumber(480), new BigNumber((_1367_v39).length), true, globalState))).length)).isEqualTo(_1363_cf24);
          let _lhs110 = _1318_v11;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(471), new BigNumber((_1318_v11).length));
          _lhs110[_lhs111] = _rhs201;
          _1349_cf12 = _rhs202;
          _1352_cf9 = _rhs203;
          _1368_v40 = _rhs204;
          _1350_cf11 = _rhs205;
        } else {
          let _1373___mcc_h10 = (_source25).cf23;
          let _1374_cf23 = _1373___mcc_h10;
          let _1375_v45;
          let _nw175 = Array((new BigNumber(6)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = _this.f3;
          _nw175[(_dafny.ONE).toNumber()] = _this.f3;
          _nw175[(new BigNumber(2)).toNumber()] = _1350_cf11;
          _nw175[(new BigNumber(3)).toNumber()] = _this.f3;
          _nw175[(new BigNumber(4)).toNumber()] = _this.f3;
          _nw175[(new BigNumber(5)).toNumber()] = _1351_cf10;
          _1375_v45 = _nw175;
          let _1376_v46;
          let _nw176 = new _module.C3();
          _nw176.__ctor((_dafny.ZERO).minus(_module.__default.safeModuloInt(_1349_cf12, p1)), _1375_v45, _this.f3, (_this).f1, (_this).f2);
          _1376_v46 = _nw176;
          let _1377_v47;
          _1377_v47 = _dafny.Seq.of(new BigNumber(((_this).f2).length), _1349_cf12);
          _1356_v34 = (_1356_v34).update(!_dafny.areEqual(_1377_v47, _dafny.Seq.of(_1349_cf12)), _1349_cf12);
          _1351_cf10 = _this.f3;
          _1351_cf10 = !(_this.f3) || (_1350_cf11);
        }
      } else if (_source24.is_DC5) {
        let _1378___mcc_h5 = (_source24).cf8;
        let _1379_cf8 = _1378___mcc_h5;
        let _1380_v48;
        let _init26 = function (_1381_i5) {
          return _this.f3;
        };
        let _nw177 = Array((new BigNumber(18)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw177.length); _i0_26++) {
          _nw177[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1380_v48 = _nw177;
        let _1382_v49;
        _1382_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1380_v48);
        r2 = new BigNumber((_1382_v49).length);
        let _1383_v50;
        _1383_v50 = _dafny.Set.fromElements(_this.f3);
        r2 = (new BigNumber(((_1383_v50).Difference(_1383_v50)).length)).plus(((_dafny.ZERO).minus(new BigNumber(-932))).minus(_module.__default.fm0(globalState)));
        _1318_v11 = _1318_v11;
        let _1384_v51;
        _1384_v51 = _dafny.Seq.of(p1);
        let _1385_v52;
        _1385_v52 = _dafny.Seq.of(p1, new BigNumber((_1384_v51).length));
        r1 = (new BigNumber(-71)).plus((new BigNumber((_1385_v52).length)).multipliedBy(new BigNumber(((_this).f2).length)));
      } else {
        let _1386___mcc_h6 = (_source24).cf13;
        let _1387_cf13 = _1386___mcc_h6;
        r2 = p1;
        (_this).f3 = !(true) || (_this.f3);
        if ((p1).isLessThan((_dafny.ZERO).minus((p1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ppfqtk")).length))))) {
          let _1388_v53;
          _1388_v53 = _dafny.Seq.UnicodeFromString("cjsmxnhxl");
          _1388_v53 = (_this).f2;
          let _index207 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index207] = (_dafny.ZERO).minus(p1);
          let _index208 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index208] = p0;
          (_this).m9((_this.f3) === (_this.f3), _this.f3, _this.f3, globalState);
          let _1389_v54;
          _1389_v54 = _dafny.Seq.of(new BigNumber((_1388_v53).length), p0);
          let _1390_v55;
          _1390_v55 = _dafny.Set.fromElements(new BigNumber(161));
          let _1391_v56;
          _1391_v56 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_dafny.Set.fromElements(p0, (_1389_v54)[_module.__default.safeIndex(p1, new BigNumber((_1389_v54).length))]), _1390_v55, _1390_v55)).Difference(_dafny.Set.fromElements(_1390_v55, _1390_v55)),(_dafny.ZERO).minus(_module.__default.fm0(globalState)));
          let _1392_v58;
          _1392_v58 = _dafny.Set.fromElements(function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_1389_v54).Elements) {
              let _1393_v57 = _compr_44;
              if (_dafny.Seq.contains(_1389_v54, _1393_v57)) {
                _coll44.add((_1393_v57).multipliedBy(new BigNumber(815)));
              }
            }
            return _coll44;
          }(), _1390_v55);
          _1391_v56 = (_1391_v56).update(_1392_v58, p0);
          let _index209 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index209] = p1;
        } else {
          let _1394_v59;
          _1394_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
          (_this).m9(false, _this.f3, !(_1394_v59).contains(p0), globalState);
          let _1395_v60;
          _1395_v60 = _dafny.Seq.of(p1, p0, (_dafny.ZERO).minus(p1), (p0).minus(p0));
          _1395_v60 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(706)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-234)), ((_1396_p1) => function (_1397_i6) {
            return _1396_p1;
          })(p1)));
          let _1398_v61;
          let _nw178 = Array((_dafny.ONE).toNumber()).fill(false);
          _1398_v61 = _nw178;
          let _index210 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1398_v61).length));
          (_1398_v61)[_index210] = _this.f3;
          let _1399_v62;
          _1399_v62 = _dafny.MultiSet.fromElements(_this.f3, _this.f3, _this.f3);
          let _index211 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1398_v61).length));
          let _rhs206 = (p0).multipliedBy(p1);
          let _rhs207 = ((_dafny.ZERO).minus(new BigNumber((_1399_v62).cardinality()))).isLessThanOrEqualTo(_module.__default.safeModuloInt(p1, p0));
          let _rhs208 = !(_this.f3);
          let _lhs112 = _this;
          let _lhs113 = _1398_v61;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1398_v61).length));
          r2 = _rhs206;
          _lhs112.f3 = _rhs207;
          _lhs113[_lhs114] = _rhs208;
          let _1400_v63;
          _1400_v63 = _dafny.MultiSet.fromElements(new BigNumber(706), p0, new BigNumber((_dafny.Set.fromElements(p0)).length));
          let _1401_v64;
          _1401_v64 = _module.D11.create_DC31(new BigNumber((_1400_v63).cardinality()));
          let _1402_v65;
          _1402_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,false);
          let _1403_v66;
          _1403_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1401_v64,_module.D2.create_DC6(_1343_v29, false, _this.f3, new BigNumber((_1402_v65).length)));
          r2 = new BigNumber((_1403_v66).length);
          let _1404_v67;
          _1404_v67 = _dafny.Set.fromElements(_this.f3);
          let _index212 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index212] = (_dafny.ZERO).minus((new BigNumber((_1404_v67).length)).minus(new BigNumber(((_this).f2).length)));
          let _index213 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_1318_v11).length));
          (_1318_v11)[_index213] = new BigNumber(688);
        }
        _1343_v29 = ((_this).f2)[_module.__default.safeIndex(p0, new BigNumber(((_this).f2).length))];
      }
      let _1405_v68;
      _1405_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
      let _1406_v69;
      _1406_v69 = _module.D6.create_DC16(_1405_v68);
      let _pat_let_tv37 = p0;
      let _pat_let_tv38 = p0;
      let _pat_let_tv39 = _1405_v68;
      r2 = function (_source26) {
        if (_source26.is_DC17) {
          let _1407___mcc_h11 = (_source26).cf28;
          let _1408_cf28 = _1407___mcc_h11;
          return new BigNumber(706);
        } else if (_source26.is_DC18) {
          let _1409___mcc_h12 = (_source26).cf29;
          let _1410___mcc_h13 = (_source26).cf30;
          let _1411___mcc_h14 = (_source26).cf31;
          let _1412___mcc_h15 = (_source26).cf32;
          let _1413___mcc_h16 = (_source26).cf33;
          let _1414_cf33 = _1413___mcc_h16;
          let _1415_cf32 = _1412___mcc_h15;
          let _1416_cf31 = _1411___mcc_h14;
          let _1417_cf30 = _1410___mcc_h13;
          let _1418_cf29 = _1409___mcc_h12;
          return _1417_cf30;
        } else if (_source26.is_DC19) {
          let _1419___mcc_h17 = (_source26).cf34;
          let _1420_cf34 = _1419___mcc_h17;
          return _pat_let_tv37;
        } else if (_source26.is_DC16) {
          let _1421___mcc_h18 = (_source26).cf27;
          let _1422_cf27 = _1421___mcc_h18;
          return new BigNumber(-982);
        } else {
          let _1423___mcc_h19 = (_source26).cf35;
          let _1424_cf35 = _1423___mcc_h19;
          return _pat_let_tv38;
        }
      }(function (_pat_let58_0) {
        return function (_1425_dt__update__tmp_h0) {
          return function (_pat_let59_0) {
            return function (_1426_dt__update_hcf27_h0) {
              return _module.D6.create_DC16(_1426_dt__update_hcf27_h0);
            }(_pat_let59_0);
          }(_pat_let_tv39);
        }(_pat_let58_0);
      }(_1406_v69));
      let _1427_v70;
      _1427_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
      r0 = (_1427_v70).Merge((_1427_v70).Merge(_1427_v70));
      let _1428_v71;
      _1428_v71 = _dafny.MultiSet.fromElements(false);
      r1 = _module.__default.safeDivisionInt((p1).plus(new BigNumber((_1428_v71).cardinality())), p0);
      r2 = _module.__default.safeModuloInt((p0).multipliedBy(p1), _module.__default.fm0(globalState));
      return [r0, r1, r2];
    }
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let _1429_v0;
      _1429_v0 = new _dafny.CodePoint('t'.codePointAt(0));
      let _1430_v1;
      _1430_v1 = _dafny.MultiSet.fromElements(p2, _this.f3, false);
      _1429_v0 = (((_1430_v1).IsDisjointFrom(_1430_v1)) ? (_1429_v0) : (_1429_v0));
      let _1431_v2;
      _1431_v2 = new BigNumber(941);
      let _1432_v3;
      _1432_v3 = _dafny.Seq.of(p2);
      let _1433_v4;
      _1433_v4 = _dafny.Seq.of(_1431_v2);
      let _1434_v5;
      let _nw179 = new _module.C2();
      _nw179.__ctor(new BigNumber(((_this).f2).length), p1, (_this).f1, (_this).f2);
      _1434_v5 = _nw179;
      let _1435_v6;
      _1435_v6 = _dafny.Seq.of(_1434_v5, _1434_v5);
      let _1436_v7;
      let _nw180 = Array((new BigNumber(19)).toNumber());
      _nw180[(_dafny.ZERO).toNumber()] = p2;
      _nw180[(_dafny.ONE).toNumber()] = p2;
      _nw180[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.fromElements(_1431_v2)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1431_v2));
      _nw180[(new BigNumber(3)).toNumber()] = true;
      _nw180[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_dafny.Seq.of(true), _1432_v3);
      _nw180[(new BigNumber(5)).toNumber()] = ((!(p1)) ? (_this.f3) : (false));
      _nw180[(new BigNumber(6)).toNumber()] = p1;
      _nw180[(new BigNumber(7)).toNumber()] = !(p2);
      _nw180[(new BigNumber(8)).toNumber()] = (_1432_v3)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_module.__default.fm0(globalState))).length), new BigNumber((_1432_v3).length))];
      _nw180[(new BigNumber(9)).toNumber()] = p1;
      _nw180[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-431)), ((_1437_v0) => function (_1438_i0) {
        return _1437_v0;
      })(_1429_v0)), (_this).f2);
      _nw180[(new BigNumber(11)).toNumber()] = (_1431_v2).isEqualTo((_1433_v4)[_module.__default.safeIndex(_1431_v2, new BigNumber((_1433_v4).length))]);
      _nw180[(new BigNumber(12)).toNumber()] = p2;
      _nw180[(new BigNumber(13)).toNumber()] = _dafny.Seq.contains(_1435_v6, _1434_v5);
      _nw180[(new BigNumber(14)).toNumber()] = ((p2) ? ((_1432_v3)[_module.__default.safeIndex(_1431_v2, new BigNumber((_1432_v3).length))]) : (p0));
      _nw180[(new BigNumber(15)).toNumber()] = p0;
      _nw180[(new BigNumber(16)).toNumber()] = _this.f3;
      _nw180[(new BigNumber(17)).toNumber()] = _dafny.Seq.IsPrefixOf((_this).f2, (_this).f2);
      _nw180[(new BigNumber(18)).toNumber()] = _1434_v5.f3;
      _1436_v7 = _nw180;
      _1436_v7 = _1436_v7;
      let _1439_v8;
      let _init27 = function (_1440_i1) {
        return _module.D9.create_DC26(new BigNumber(892));
      };
      let _nw181 = Array((new BigNumber(11)).toNumber());
      for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw181.length); _i0_27++) {
        _nw181[_i0_27] = _init27(new BigNumber(_i0_27));
      }
      _1439_v8 = _nw181;
      let _1441_v9;
      _1441_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7);
      let _1442_v10;
      _1442_v10 = _dafny.Seq.of(_1436_v7);
      let _1443_v11;
      _1443_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v2,_1436_v7);
      let _1444_v12;
      _1444_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1431_v2);
      let _1445_v13;
      _1445_v13 = _dafny.Seq.of(_1441_v9);
      let _1446_v14;
      let _nw182 = Array((new BigNumber(25)).toNumber());
      _nw182[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7);
      _nw182[(_dafny.ONE).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(2)).toNumber()] = (_1441_v9).update(_1439_v8, _1436_v7);
      _nw182[(new BigNumber(3)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7);
      _nw182[(new BigNumber(5)).toNumber()] = (_1441_v9).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7));
      _nw182[(new BigNumber(6)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(7)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(8)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(9)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(10)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(11)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(12)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(13)).toNumber()] = (_1441_v9).Merge(_1441_v9);
      _nw182[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,(_1442_v10)[_module.__default.safeIndex(new BigNumber((_1430_v1).cardinality()), new BigNumber((_1442_v10).length))]);
      _nw182[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7);
      _nw182[(new BigNumber(16)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(17)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(18)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(19)).toNumber()] = (_1441_v9).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7));
      _nw182[(new BigNumber(20)).toNumber()] = _1441_v9;
      _nw182[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1439_v8,(((_1443_v11).contains((_dafny.ZERO).minus(new BigNumber((_1444_v12).length)))) ? ((_1443_v11).get((_dafny.ZERO).minus(new BigNumber((_1444_v12).length)))) : (_1436_v7)));
      _nw182[(new BigNumber(22)).toNumber()] = ((_1445_v13)[_module.__default.safeIndex(_1431_v2, new BigNumber((_1445_v13).length))]).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1439_v8,_1436_v7)).update(_1439_v8, _1436_v7));
      _nw182[(new BigNumber(23)).toNumber()] = (_1441_v9).Merge(_1441_v9);
      _nw182[(new BigNumber(24)).toNumber()] = _1441_v9;
      _1446_v14 = _nw182;
      let _index214 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_1446_v14).length));
      (_1446_v14)[_index214] = _1441_v9;
      let _1447_v15;
      _1447_v15 = _dafny.MultiSet.fromElements(_1431_v2, _1431_v2);
      let _1448_v16;
      _1448_v16 = _dafny.Set.fromElements(p0);
      let _1449_v17;
      _1449_v17 = _module.D3.create_DC8(_1448_v16);
      let _pat_let_tv40 = _1434_v5;
      let _pat_let_tv41 = _1434_v5;
      let _pat_let_tv42 = _1434_v5;
      let _pat_let_tv43 = _1434_v5;
      let _index215 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_1446_v14).length));
      let _rhs209 = _dafny.Seq.IsProperPrefixOf(_1433_v4, _dafny.Seq.update(_1433_v4, _module.__default.safeIndex(_1431_v2, new BigNumber((_1433_v4).length)), _1431_v2));
      let _rhs210 = (_dafny.ZERO).minus((((_1447_v15).contains((_1431_v2).minus(_1431_v2))) ? ((_1447_v15).get((_1431_v2).minus(_1431_v2))) : ((_1431_v2).plus(_1431_v2))));
      let _rhs211 = p0;
      let _rhs212 = !(function (_source27) {
        if (_source27.is_DC9) {
          let _1450___mcc_h0 = (_source27).cf15;
          let _1451___mcc_h1 = (_source27).cf16;
          let _1452_cf16 = _1451___mcc_h1;
          let _1453_cf15 = _1450___mcc_h0;
          return !(_pat_let_tv40.f3);
        } else if (_source27.is_DC10) {
          let _1454___mcc_h2 = (_source27).cf17;
          let _1455___mcc_h3 = (_source27).cf18;
          let _1456_cf18 = _1455___mcc_h3;
          let _1457_cf17 = _1454___mcc_h2;
          return _pat_let_tv41.f3;
        } else if (_source27.is_DC8) {
          let _1458___mcc_h4 = (_source27).cf14;
          let _1459_cf14 = _1458___mcc_h4;
          return _pat_let_tv42.f3;
        } else {
          let _1460___mcc_h5 = (_source27).cf19;
          let _1461_cf19 = _1460___mcc_h5;
          return _pat_let_tv43.f3;
        }
      }(_1449_v17));
      let _rhs213 = _1441_v9;
      let _lhs115 = _1434_v5;
      let _lhs116 = _this;
      let _lhs117 = _this;
      let _lhs118 = _1446_v14;
      let _lhs119 = _module.__default.safeIndex(new BigNumber(890), new BigNumber((_1446_v14).length));
      _lhs115.f3 = _rhs209;
      _1431_v2 = _rhs210;
      _lhs116.f3 = _rhs211;
      _lhs117.f3 = _rhs212;
      _lhs118[_lhs119] = _rhs213;
      let _1462_i2;
      _1462_i2 = _dafny.ZERO;
      L9: {
        while (_dafny.areEqual((_this).f2, (_this).f2)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1462_i2)) {
              break L9;
            }
            _1462_i2 = (_1462_i2).plus(_dafny.ONE);
            let _1463_v18;
            _1463_v18 = _module.D6.create_DC17(_1434_v5.f3);
            let _1464_v19;
            _1464_v19 = _module.D6.create_DC20(_1463_v18);
            let _1465_v20;
            _1465_v20 = _module.D6.create_DC20(_1463_v18);
            let _1466_v21;
            _1466_v21 = _module.D6.create_DC20(_1464_v19);
            let _source28 = _1466_v21;
            if (_source28.is_DC17) {
              let _1467___mcc_h6 = (_source28).cf28;
              let _1468_cf28 = _1467___mcc_h6;
              let _1469_v22;
              _1469_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v2,_1434_v5.f3);
              let _1470_v23;
              _1470_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v2,_1431_v2);
              let _1471_v24;
              _1471_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_1470_v23).contains(_1431_v2)) ? ((_1470_v23).get(_1431_v2)) : ((((_1444_v12).contains(p1)) ? ((_1444_v12).get(p1)) : (new BigNumber(((_1434_v5).f2).length))))),_1444_v12);
              (_1434_v5).f3 = (_1471_v24).contains(new BigNumber((_1469_v22).length));
              let _1472_v25;
              let _nw183 = new _module.C2();
              _nw183.__ctor((_1431_v2).minus(new BigNumber(-24)), _1434_v5.f3, (_1434_v5).f1, (_1434_v5).f2);
              _1472_v25 = _nw183;
              let _1473_v26;
              _1473_v26 = _dafny.Set.fromElements(_1431_v2, _1431_v2);
              let _1474_v27;
              _1474_v27 = _module.D6.create_DC18(new BigNumber((_1473_v26).length), (_1472_v25).f13, true, _1429_v0, (_1472_v25).f13);
              let _1475_v28;
              _1475_v28 = _dafny.Seq.of(_1474_v27);
              let _1476_v29;
              _1476_v29 = _dafny.Seq.of(_dafny.Seq.Concat(_1475_v28, _1475_v28), _1475_v28, _dafny.Seq.of(_1474_v27, _module.D6.create_DC18(new BigNumber(282), (_1472_v25).f13, _1468_cf28, _1429_v0, (_1472_v25).f13)), _dafny.Seq.Concat(_module.__default.fm48(_1434_v5.f3, _1431_v2, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(815)), ((_1477_v27) => function (_1478_i3) {
                return _1477_v27;
              })(_1474_v27))));
              _1476_v29 = _dafny.Seq.Concat(_1476_v29, _dafny.Seq.of(_1475_v28));
              let _1479_v30;
              _1479_v30 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.update(_dafny.Seq.of(new BigNumber(932), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), function (_1480_i5) {
                return new _dafny.CodePoint('l'.codePointAt(0));
              })).length)), _module.__default.safeIndex(_1431_v2, new BigNumber((_dafny.Seq.of(new BigNumber(932), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(72)), function (_1481_i5) {
                return new _dafny.CodePoint('l'.codePointAt(0));
              })).length))).length)), (_1472_v25).f13));
              let _1482_v31;
              _1482_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1434_v5.f3);
              let _1483_v32;
              _1483_v32 = _module.D9.create_DC27(new BigNumber((_dafny.Seq.UnicodeFromString("ina")).length));
              let _1484_v33;
              _1484_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1483_v32,!(_1434_v5.f3));
              let _1485_v34;
              let _nw184 = Array((new BigNumber(25)).toNumber());
              _nw184[(_dafny.ZERO).toNumber()] = _1433_v4;
              _nw184[(_dafny.ONE).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(2)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1433_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), ((_1486_v4, _1487_v25) => function (_1488_i4) {
                return (_dafny.ZERO).minus((_1486_v4)[_module.__default.safeIndex((_1487_v25).f13, new BigNumber((_1486_v4).length))]);
              })(_1433_v4, _1472_v25)));
              _nw184[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1433_v4, (((_1479_v30).contains(_1434_v5.f3)) ? ((_1479_v30).get(_1434_v5.f3)) : (_1433_v4)));
              _nw184[(new BigNumber(5)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(6)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_1431_v2, _1431_v2, (((_1430_v1).contains(_1434_v5.f3)) ? ((_1430_v1).get(_1434_v5.f3)) : (new BigNumber((_1482_v31).length))), (_1472_v25).f13, (_1472_v25).f13);
              _nw184[(new BigNumber(8)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(9)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), ((_1489_v5, _1490_v23, _1491_v25) => function (_1492_i6) {
                return (((_1490_v23).contains(new BigNumber(((_1489_v5).f2).length))) ? ((_1490_v23).get(new BigNumber(((_1489_v5).f2).length))) : ((_1491_v25).f13));
              })(_1434_v5, _1470_v23, _1472_v25));
              _nw184[(new BigNumber(11)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(672)), ((_1493_v25) => function (_1494_i7) {
                return (_1493_v25).f13;
              })(_1472_v25));
              _nw184[(new BigNumber(13)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(14)).toNumber()] = ((_1468_cf28) ? (_dafny.Seq.of((_1434_v5).fm12(_1431_v2, (_1472_v25).f13, new BigNumber(((_this).f2).length), _1434_v5.f3, globalState))) : (_dafny.Seq.of((_1472_v25).f13)));
              _nw184[(new BigNumber(15)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(16)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(17)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(18)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(19)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(20)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(21)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(22)).toNumber()] = _1433_v4;
              _nw184[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_1433_v4, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), ((_1495_v16) => function (_1496_i8) {
                return new BigNumber((_1495_v16).length);
              })(_1448_v16)), _module.__default.safeIndex(new BigNumber(-923), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(819)), ((_1497_v16) => function (_1498_i8) {
                return new BigNumber((_1497_v16).length);
              })(_1448_v16))).length)), new BigNumber((_1484_v33).length)));
              _nw184[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_1433_v4, _1433_v4);
              _1485_v34 = _nw184;
              let _index216 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1485_v34).length));
              (_1485_v34)[_index216] = _1433_v4;
              let _index217 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_1485_v34).length));
              (_1485_v34)[_index217] = _module.__default.fm31((new BigNumber(254)).plus(_module.__default.fm0(globalState)), p1, _module.__default.safeModuloInt((_1472_v25).f13, (_1472_v25).f13), _1429_v0, globalState);
            } else if (_source28.is_DC18) {
              let _1499___mcc_h7 = (_source28).cf29;
              let _1500___mcc_h8 = (_source28).cf30;
              let _1501___mcc_h9 = (_source28).cf31;
              let _1502___mcc_h10 = (_source28).cf32;
              let _1503___mcc_h11 = (_source28).cf33;
              let _1504_cf33 = _1503___mcc_h11;
              let _1505_cf32 = _1502___mcc_h10;
              let _1506_cf31 = _1501___mcc_h9;
              let _1507_cf30 = _1500___mcc_h8;
              let _1508_cf29 = _1499___mcc_h7;
              (_1434_v5).f3 = _1434_v5.f3;
              _1431_v2 = ((p1) ? (_1504_cf33) : (_1508_cf29));
              let _1509_v35;
              _1509_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.update(_1433_v4, _module.__default.safeIndex(_1431_v2, new BigNumber((_1433_v4).length)), _1508_cf29)),_1434_v5.f3);
              let _1510_v36;
              _1510_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_1434_v5.f3),_1509_v35);
              _1510_v36 = (_1510_v36).update(_1434_v5.f3, (_1509_v35).update(_dafny.MultiSet.fromElements(new BigNumber(160), _1431_v2), !(p1)));
              let _1511_v37;
              _1511_v37 = _dafny.Seq.of(_1447_v15, _1447_v15, _1447_v15, _1447_v15);
              _1447_v15 = (_1511_v37)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1504_cf33, (_dafny.ZERO).minus(_1504_cf33)), new BigNumber((_1511_v37).length))];
            } else if (_source28.is_DC19) {
              let _1512___mcc_h12 = (_source28).cf34;
              let _1513_cf34 = _1512___mcc_h12;
              _1513_cf34 = _1513_cf34;
              _1431_v2 = _module.__default.fm0(globalState);
              let _1514_v38;
              _1514_v38 = _module.D9.create_DC26(_1431_v2);
              (_1434_v5).f3 = (!((_1514_v38).dtor_cf44).isEqualTo(_1513_cf34)) || ((_1432_v3)[_module.__default.safeIndex(_1513_cf34, new BigNumber((_1432_v3).length))]);
              let _1515_v39;
              _1515_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v2,true);
              _1515_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1430_v1).cardinality()),p2);
            } else if (_source28.is_DC16) {
              let _1516___mcc_h13 = (_source28).cf27;
              let _1517_cf27 = _1516___mcc_h13;
              (_1434_v5).f3 = ((_1447_v15).update(new BigNumber((_1444_v12).length), _module.__default.abs(_1431_v2))).IsSubsetOf((_1447_v15).Difference(_1447_v15));
              _1431_v2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_1431_v2).minus(_1431_v2), new BigNumber((_1517_cf27).length)));
              let _1518_v40;
              let _init28 = function (_1519_i9) {
                return (_1519_i9).multipliedBy(new BigNumber(((_this).f2).length));
              };
              let _nw185 = Array((new BigNumber(28)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw185.length); _i0_28++) {
                _nw185[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _1518_v40 = _nw185;
              let _index218 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1518_v40).length));
              (_1518_v40)[_index218] = _module.__default.safeDivisionInt(_1431_v2, _1431_v2);
              let _index219 = _module.__default.safeIndex(new BigNumber(730), new BigNumber((_1518_v40).length));
              (_1518_v40)[_index219] = (_1431_v2).minus(_1431_v2);
              let _1520_v41;
              let _nw186 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1520_v41 = _nw186;
              let _index220 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_1520_v41).length));
              (_1520_v41)[_index220] = _dafny.Seq.Concat((_this).f2, _module.__default.fm2(p2, globalState));
              let _index221 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_1520_v41).length));
              (_1520_v41)[_index221] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("tyiwhk"), _module.__default.safeIndex(new BigNumber((_1517_cf27).length), new BigNumber((_dafny.Seq.UnicodeFromString("tyiwhk")).length)), _1429_v0), _module.__default.safeIndex(_1431_v2, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("tyiwhk"), _module.__default.safeIndex(new BigNumber((_1517_cf27).length), new BigNumber((_dafny.Seq.UnicodeFromString("tyiwhk")).length)), _1429_v0)).length)), ((!(_1434_v5.f3)) ? (_1429_v0) : (_1429_v0)));
            } else {
              let _1521___mcc_h14 = (_source28).cf35;
              let _1522_cf35 = _1521___mcc_h14;
              let _1523_v42;
              _1523_v42 = _module.D1.create_DC3(p1, _1431_v2, _1431_v2);
              (_1434_v5).f3 = (new BigNumber(454)).isLessThan((_dafny.ZERO).minus(((!(_1434_v5.f3)) ? ((_1433_v4)[_module.__default.safeIndex(new BigNumber(497), new BigNumber((_1433_v4).length))]) : ((_1523_v42).dtor_cf4))));
              let _1524_v44;
              _1524_v44 = _module.D6.create_DC17(!(p0));
              let _1525_v45;
              _1525_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1524_v44,_1431_v2);
              let _1526_v46;
              _1526_v46 = _module.D3.create_DC10(_1431_v2, (_dafny.ZERO).minus(_1431_v2));
              let _1527_v47;
              let _nw187 = Array((new BigNumber(28)).toNumber());
              _nw187[(_dafny.ZERO).toNumber()] = _1431_v2;
              _nw187[(_dafny.ONE).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(2)).toNumber()] = new BigNumber((_1430_v1).cardinality());
              _nw187[(new BigNumber(3)).toNumber()] = new BigNumber(319);
              _nw187[(new BigNumber(4)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(5)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(6)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(7)).toNumber()] = new BigNumber((function () {
                let _coll45 = new _dafny.Map();
                for (const _compr_45 of (_1525_v45).Keys.Elements) {
                  let _1528_v43 = _compr_45;
                  if ((_1525_v45).contains(_1528_v43)) {
                    _coll45.push([_1528_v43,_1429_v0]);
                  }
                }
                return _coll45;
              }()).length);
              _nw187[(new BigNumber(8)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(9)).toNumber()] = new BigNumber(((_1434_v5).f2).length);
              _nw187[(new BigNumber(10)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(11)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(12)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((((_1444_v12).contains(_1434_v5.f3)) ? ((_1444_v12).get(_1434_v5.f3)) : (_1431_v2)));
              _nw187[(new BigNumber(14)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(15)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(16)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(17)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(18)).toNumber()] = (_1434_v5).fm12(_1431_v2, new BigNumber(((_1434_v5).f2).length), _1431_v2, _1434_v5.f3, globalState);
              _nw187[(new BigNumber(19)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(20)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(21)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(22)).toNumber()] = new BigNumber((_1430_v1).cardinality());
              _nw187[(new BigNumber(23)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(24)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(25)).toNumber()] = _1431_v2;
              _nw187[(new BigNumber(26)).toNumber()] = (_1526_v46).dtor_cf17;
              _nw187[(new BigNumber(27)).toNumber()] = _1431_v2;
              _1527_v47 = _nw187;
              let _1529_v48;
              let _nw188 = new _module.C0();
              _nw188.__ctor(_1527_v47);
              _1529_v48 = _nw188;
              _1436_v7 = _1436_v7;
              _1431_v2 = (_this).fm12(new BigNumber(-956), (_1431_v2).minus(_1431_v2), _1431_v2, _this.f3, globalState);
            }
            let _1530_v49;
            _1530_v49 = _module.D2.create_DC6(new _dafny.CodePoint('v'.codePointAt(0)), p0, _1434_v5.f3, _1431_v2);
            let _source29 = _1530_v49;
            if (_source29.is_DC6) {
              let _1531___mcc_h15 = (_source29).cf9;
              let _1532___mcc_h16 = (_source29).cf10;
              let _1533___mcc_h17 = (_source29).cf11;
              let _1534___mcc_h18 = (_source29).cf12;
              let _1535_cf12 = _1534___mcc_h18;
              let _1536_cf11 = _1533___mcc_h17;
              let _1537_cf10 = _1532___mcc_h16;
              let _1538_cf9 = _1531___mcc_h15;
              let _1539_v50;
              _1539_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1537_cf10,_1538_cf9);
              let _1540_v51;
              _1540_v51 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1539_v50).length)).minus(new BigNumber(120)),_1433_v4);
              _1433_v4 = (((_1540_v51).contains(new BigNumber((_1442_v10).length))) ? ((_1540_v51).get(new BigNumber((_1442_v10).length))) : (_1433_v4));
              let _rhs214 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(179), ((p0) ? (_1431_v2) : (new BigNumber(249)))));
              let _rhs215 = ((((_1447_v15).contains(new BigNumber((_1432_v3).length))) ? ((_1447_v15).get(new BigNumber((_1432_v3).length))) : (_module.__default.fm0(globalState)))).plus(((_1434_v5.f3) ? (new BigNumber(-47)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(8)), ((_1541_v5) => function (_1542_i10) {
                return new BigNumber(((_1541_v5).f2).length);
              })(_1434_v5))).length))));
              _1535_cf12 = _rhs214;
              _1431_v2 = _rhs215;
              _1431_v2 = new BigNumber(-84);
              let _1543_v52;
              _1543_v52 = _dafny.Set.fromElements(_1431_v2);
              let _rhs216 = _1537_cf10;
              let _rhs217 = true;
              let _rhs218 = _1433_v4;
              let _rhs219 = p2;
              let _rhs220 = (((_1543_v52).contains(_1431_v2)) ? (p2) : (((false) ? (p0) : (!(_1434_v5.f3)))));
              let _lhs120 = _1434_v5;
              let _lhs121 = _this;
              let _lhs122 = _1434_v5;
              let _lhs123 = _1434_v5;
              _lhs120.f3 = _rhs216;
              _lhs121.f3 = _rhs217;
              _1433_v4 = _rhs218;
              _lhs122.f3 = _rhs219;
              _lhs123.f3 = _rhs220;
            } else if (_source29.is_DC5) {
              let _1544___mcc_h19 = (_source29).cf8;
              let _1545_cf8 = _1544___mcc_h19;
              (_1434_v5).f3 = !_dafny.Seq.contains(_dafny.Seq.Concat((_this).f2, (_1434_v5).f2), _1429_v0);
              let _1546_v53;
              _1546_v53 = _module.D11.create_DC31(_1431_v2);
              let _1547_v54;
              _1547_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm49(_1546_v53, new BigNumber((_1430_v1).cardinality()), _1431_v2, globalState),_1431_v2);
              let _1548_v55;
              _1548_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1547_v54,p0);
              _1548_v55 = (_1548_v55).update((_1547_v54).Merge(_1547_v54), (false) || (_this.f3));
              _1545_cf8 = (_1545_cf8).update(p0, _1431_v2);
              let _rhs221 = p0;
              let _rhs222 = !(!(((_1434_v5).fm11(!((_this).fm11(_module.__default.fm3(_1431_v2, globalState), globalState)), globalState)) && (_1434_v5.f3)));
              let _lhs124 = _1434_v5;
              let _lhs125 = _1434_v5;
              _lhs124.f3 = _rhs221;
              _lhs125.f3 = _rhs222;
            } else {
              let _1549___mcc_h20 = (_source29).cf13;
              let _1550_cf13 = _1549___mcc_h20;
              let _1551_v56;
              let _nw189 = new _module.C2();
              _nw189.__ctor(_1431_v2, false, (_1434_v5).f1, (_1434_v5).f2);
              _1551_v56 = _nw189;
              let _1552_v57;
              let _nw190 = new _module.C1();
              _nw190.__ctor();
              _1552_v57 = _nw190;
              let _1553_v58;
              let _nw191 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.of());
              _1553_v58 = _nw191;
              let _rhs223 = (_1551_v56).f13;
              let _rhs224 = p2;
              let _rhs225 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1431_v2), _1431_v2);
              let _rhs226 = ((_1434_v5.f3) ? (_1553_v58) : (_1553_v58));
              let _rhs227 = p0;
              let _lhs126 = _1434_v5;
              let _lhs127 = _1434_v5;
              _1431_v2 = _rhs223;
              _lhs126.f3 = _rhs224;
              _1431_v2 = _rhs225;
              _1553_v58 = _rhs226;
              _lhs127.f3 = _rhs227;
              (_1434_v5).f3 = p1;
            }
            _1431_v2 = _1431_v2;
            if (false) {
              let _1554_v59;
              let _init29 = ((_1555_v2) => function (_1556_i11) {
                return (_1556_i11).multipliedBy(_1555_v2);
              })(_1431_v2);
              let _nw192 = Array((new BigNumber(22)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw192.length); _i0_29++) {
                _nw192[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _1554_v59 = _nw192;
              let _index222 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1554_v59).length));
              (_1554_v59)[_index222] = _1431_v2;
              let _index223 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1554_v59).length));
              (_1554_v59)[_index223] = _1431_v2;
              (_this).f3 = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), ((_1557_v5) => function (_1558_i12) {
                return (_1557_v5).f2;
              })(_1434_v5)), (_1434_v5).f2);
              let _rhs228 = (_1431_v2).minus((_1554_v59)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_1554_v59).length))]);
              let _rhs229 = p1;
              let _rhs230 = _1429_v0;
              let _lhs128 = _1434_v5;
              _1431_v2 = _rhs228;
              _lhs128.f3 = _rhs229;
              _1429_v0 = _rhs230;
              let _1559_v60;
              _1559_v60 = _dafny.Seq.UnicodeFromString("cjg");
              _1559_v60 = (_1434_v5).f2;
              let _1560_v61;
              _1560_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1431_v2,(_1432_v3)[_module.__default.safeIndex((_dafny.ZERO).minus((_1554_v59)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_1554_v59).length))]), new BigNumber((_1432_v3).length))]);
              _1560_v61 = _1560_v61;
            } else {
              let _1561_v62;
              let _init30 = ((_1562_v0, _1563_v5, _1564_v2) => function (_1565_i13) {
                return (_1565_i13).multipliedBy((_module.D2.create_DC6(_1562_v0, _1563_v5.f3, _1563_v5.f3, _1564_v2)).dtor_cf12);
              })(_1429_v0, _1434_v5, _1431_v2);
              let _nw193 = Array((new BigNumber(27)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw193.length); _i0_30++) {
                _nw193[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1561_v62 = _nw193;
              let _1566_v63;
              _1566_v63 = _dafny.Set.fromElements(new BigNumber(-299));
              let _index224 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1561_v62).length));
              (_1561_v62)[_index224] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1566_v63).length),p1)).length);
              let _index225 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1561_v62).length));
              (_1561_v62)[_index225] = _1431_v2;
              let _index226 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1561_v62).length));
              (_1561_v62)[_index226] = (new BigNumber(771)).multipliedBy(((_1434_v5.f3) ? ((_1561_v62)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_1561_v62).length))]) : (_1431_v2)));
              let _1567_v64;
              let _nw194 = new _module.C1();
              _nw194.__ctor();
              _1567_v64 = _nw194;
              _1466_v21 = _module.D6.create_DC20(_1464_v19);
              let _1568_v65;
              let _nw195 = new _module.C2();
              _nw195.__ctor(_1431_v2, p0, (_1434_v5).f1, _dafny.Seq.update((_1434_v5).f2, _module.__default.safeIndex(_1431_v2, new BigNumber(((_1434_v5).f2).length)), _1429_v0));
              _1568_v65 = _nw195;
              let _1569_v66;
              let _nw196 = Array((new BigNumber(16)).toNumber());
              _nw196[(_dafny.ZERO).toNumber()] = ((_1434_v5.f3) ? (_1568_v65) : (_1568_v65));
              _nw196[(_dafny.ONE).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(2)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(3)).toNumber()] = ((p0) ? (_1568_v65) : (_1568_v65));
              _nw196[(new BigNumber(4)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(5)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(6)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(7)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(8)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(9)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(10)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(11)).toNumber()] = (_module.D13.create_DC33(_1568_v65)).dtor_cf52;
              _nw196[(new BigNumber(12)).toNumber()] = ((p0) ? (_1568_v65) : (_1568_v65));
              _nw196[(new BigNumber(13)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(14)).toNumber()] = _1568_v65;
              _nw196[(new BigNumber(15)).toNumber()] = _1568_v65;
              _1569_v66 = _nw196;
              let _index227 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1569_v66).length));
              (_1569_v66)[_index227] = _1568_v65;
              let _index228 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1569_v66).length));
              let _rhs231 = _1434_v5.f3;
              let _rhs232 = _1568_v65;
              let _lhs129 = _1434_v5;
              let _lhs130 = _1569_v66;
              let _lhs131 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1569_v66).length));
              _lhs129.f3 = _rhs231;
              _lhs130[_lhs131] = _rhs232;
            }
          }
        }
      }
      let _1570_i14;
      _1570_i14 = _dafny.ZERO;
      L10: {
        while (_1434_v5.f3) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1570_i14)) {
              break L10;
            }
            _1570_i14 = (_1570_i14).plus(_dafny.ONE);
            let _1571_v67;
            _1571_v67 = _dafny.Seq.UnicodeFromString("aw");
            _1571_v67 = _dafny.Seq.Concat((_1434_v5).f2, _dafny.Seq.Concat(_1571_v67, (_1434_v5).f2));
            _1431_v2 = _module.__default.fm0(globalState);
            let _1572_v68;
            _1572_v68 = _module.D5.create_DC15(_1431_v2, _1434_v5.f3, _1436_v7);
            (_1434_v5).f3 = ((_1572_v68).dtor_cf24).isLessThan(_1431_v2);
            (_this).f3 = p0;
          }
        }
      }
      let _hi6 = (_dafny.ZERO).minus(_1431_v2);
      for (let _1573_i15 = (_1431_v2).multipliedBy(_1431_v2); _1573_i15.isLessThan(_hi6); _1573_i15 = _1573_i15.plus(_dafny.ONE)) {
        let _1574_v69;
        let _nw197 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1574_v69 = _nw197;
        let _index229 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1574_v69).length));
        (_1574_v69)[_index229] = _1573_i15;
        let _1575_v70;
        _1575_v70 = _module.D1.create_DC3(_1434_v5.f3, _1573_i15, (_this).fm12(_1573_i15, _1431_v2, _1431_v2, p1, globalState));
        let _index230 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1574_v69).length));
        (_1574_v69)[_index230] = ((new BigNumber(((_1434_v5).f2).length)).plus((_1575_v70).dtor_cf5)).plus(_1431_v2);
        (_this).f3 = (_1434_v5.f3) || (_1434_v5.f3);
        let _1576_v71;
        _1576_v71 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_1436_v7);
        let _rhs233 = (_1573_i15).isLessThanOrEqualTo(_1573_i15);
        let _rhs234 = (((_1576_v71).contains(_1429_v0)) ? ((_1576_v71).get(_1429_v0)) : (_1436_v7));
        let _rhs235 = p2;
        let _rhs236 = (!(false) || (p0)) === (!(_1434_v5.f3));
        let _lhs132 = _this;
        let _lhs133 = _1434_v5;
        let _lhs134 = _this;
        _lhs132.f3 = _rhs233;
        _1436_v7 = _rhs234;
        _lhs133.f3 = _rhs235;
        _lhs134.f3 = _rhs236;
        _1431_v2 = (_dafny.ZERO).minus((((new BigNumber(-781)).isLessThanOrEqualTo(_1431_v2)) ? ((_1574_v69)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_1574_v69).length))]) : (_module.__default.safeModuloInt((_1574_v69)[_module.__default.safeIndex(new BigNumber(90), new BigNumber((_1574_v69).length))], _1573_i15))));
      }
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f15 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    __ctor(f15, f1, f2) {
      let _this = this;
      (_this)._f15 = f15;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm11(p0, globalState) {
      let _this = this;
      return _dafny.areEqual((((_this).f15) ? (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-315),new BigNumber(-858))).length)), new BigNumber(((_this).f2).length), new BigNumber(204))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(361)), function (_1577_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xwmv")).length),new BigNumber(((_this).f2).length))).length));
      }))), _dafny.Seq.of(new BigNumber(-970)));
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
      this.f12 = _dafny.ZERO;
      this._f11 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f11, f12, f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(_this.f3)).Union(_dafny.MultiSet.fromElements(_this.f3, _this.f3))).IsSubsetOf(_dafny.MultiSet.fromElements(_this.f3, _this.f3, _this.f3, _this.f3));
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iweafjioe"), (_this).f2), _dafny.Seq.UnicodeFromString("kruhckdc"));
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm11(p0, globalState) {
      let _this = this;
      if (_this.f3) {
        return _this.f3;
      } else {
        return (true) || (_this.f3);
      }
    };
    fm21(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(215)), function (_1578_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length), _this.f12);
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      if (_this.f3) {
        let _1579_v0;
        _1579_v0 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1580_v2;
        _1580_v2 = _module.D3.create_DC10(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), function (_1581_i0) {
  return (_this).f11;
})).length), new BigNumber((function () {
  let _coll46 = new _dafny.Set();
  for (const _compr_46 of _dafny.IntegerRange(new BigNumber(493), new BigNumber(7))) {
    let _1582_v1 = _compr_46;
    if (((new BigNumber(493)).isLessThanOrEqualTo(_1582_v1)) && ((_1582_v1).isLessThan(new BigNumber(7)))) {
      _coll46.add(_module.__default.safeModuloInt(_1582_v1, p1));
    }
  }
  return _coll46;
}()).length));
        let _1583_v3;
        _1583_v3 = _module.D3.create_DC11(_1580_v2);
        let _1584_v4;
        _1584_v4 = _dafny.MultiSet.fromElements(_1583_v3);
        let _1585_v5;
        let _init31 = ((_1586_p0) => function (_1587_i1) {
          return _module.__default.safeModuloInt(_1587_i1, (_dafny.ZERO).minus(_1586_p0));
        })(p0);
        let _nw198 = Array((new BigNumber(5)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw198.length); _i0_31++) {
          _nw198[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1585_v5 = _nw198;
        let _1588_v6;
        _1588_v6 = _dafny.MultiSet.fromElements(_this.f3, _this.f3);
        let _1589_v7;
        _1589_v7 = _dafny.Seq.of(_this.f3);
        let _rhs237 = _module.__default.fm22(_this.f3, new BigNumber(((_this).f2).length), globalState);
        let _rhs238 = _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(_module.__default.fm23(new BigNumber((_1588_v6).cardinality()), (_this).f11, _this.f12, _1579_v0, globalState), _1583_v3, _1583_v3, _1583_v3, _1583_v3), _module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f12, (_this).fm12((_this).f4, _this.f12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f11)).length), _this.f3, globalState))), new BigNumber((_dafny.Seq.of(_module.__default.fm23(new BigNumber((_1588_v6).cardinality()), (_this).f11, _this.f12, _1579_v0, globalState), _1583_v3, _1583_v3, _1583_v3, _1583_v3)).length)), _1583_v3));
        let _rhs239 = (new BigNumber((_1589_v7).length)).multipliedBy(_module.__default.safeModuloInt((_this).f4, (_this).f11));
        let _rhs240 = _1585_v5;
        let _rhs241 = _this.f3;
        let _lhs135 = _this;
        _1579_v0 = _rhs237;
        _1584_v4 = _rhs238;
        r1 = _rhs239;
        _1585_v5 = _rhs240;
        _lhs135.f3 = _rhs241;
        (_this).f3 = _this.f3;
        (_this).f3 = !(_this.f3);
        let _1590_v8;
        _1590_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
        let _1591_v9;
        _1591_v9 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_module.__default.fm2(_this.f3, globalState)).length)));
        let _1592_v10;
        _1592_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_1591_v9);
        let _1593_v11;
        _1593_v11 = _dafny.Set.fromElements(_this.f3);
        let _1594_v12;
        _1594_v12 = _dafny.Seq.of(new BigNumber(922));
        _1590_v8 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1592_v10).update(new BigNumber(-641), _dafny.MultiSet.fromElements(new BigNumber((_1593_v11).length), new BigNumber(495), _this.f12, _this.f12, (_1594_v12)[_module.__default.safeIndex((_this).f11, new BigNumber((_1594_v12).length))]))).length),false)).Merge(_1590_v8);
        let _1595_v13;
        _1595_v13 = _module.D5.create_DC14(_1585_v5);
        _1595_v13 = ((_this.f3) ? (_1595_v13) : (_1595_v13));
      } else {
        let _1596_v14;
        _1596_v14 = _module.D3.create_DC9(new BigNumber(((_this).f2).length), _this.f3);
        let _index231 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index231] = (_1596_v14).dtor_cf16;
        let _1597_v15;
        _1597_v15 = _dafny.MultiSet.fromElements((_this).f4);
        let _1598_v16;
        _1598_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), function (_1599_i2) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f11)).length);
        }), _dafny.Seq.of(new BigNumber((_1597_v15).cardinality()), (_this).f11, p0)),_this.f3);
        let _index232 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index232] = (((_1598_v16).contains(_this.f3)) ? ((_1598_v16).get(_this.f3)) : (_this.f3));
        let _1600_v17;
        _1600_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm8(p0, (_this).f2, globalState)).cardinality()),_1598_v16);
        let _1601_v18;
        _1601_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber(708));
        let _1602_v19;
        _1602_v19 = _module.D1.create_DC3(((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))], p0, (_this).f11);
        let _1603_v20;
        _1603_v20 = _dafny.Seq.of((_1602_v19).dtor_cf3);
        let _1604_v21;
        let _nw199 = Array((new BigNumber(9)).toNumber());
        _nw199[(_dafny.ZERO).toNumber()] = (_this).f11;
        _nw199[(_dafny.ONE).toNumber()] = _module.__default.fm0(globalState);
        _nw199[(new BigNumber(2)).toNumber()] = _this.f12;
        _nw199[(new BigNumber(3)).toNumber()] = new BigNumber((_1601_v18).length);
        _nw199[(new BigNumber(4)).toNumber()] = new BigNumber((_1603_v20).length);
        _nw199[(new BigNumber(5)).toNumber()] = new BigNumber((_module.__default.fm9(globalState)).length);
        _nw199[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw199[(new BigNumber(7)).toNumber()] = new BigNumber(-422);
        _nw199[(new BigNumber(8)).toNumber()] = _this.f12;
        _1604_v21 = _nw199;
        let _1605_v22;
        _1605_v22 = _dafny.Set.fromElements(_1604_v21, _1604_v21);
        let _1606_v23;
        _1606_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1600_v17,_1605_v22);
        _1606_v23 = (_1606_v23).update((_1600_v17).Merge(function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of _dafny.IntegerRange(new BigNumber(984), new BigNumber(273))) {
            let _1607_v24 = _compr_47;
            if (((new BigNumber(984)).isLessThanOrEqualTo(_1607_v24)) && ((_1607_v24).isLessThan(new BigNumber(273)))) {
              _coll47.push([(_1607_v24).multipliedBy(new BigNumber((_dafny.Seq.of(p0)).length)),_1598_v16]);
            }
          }
          return _coll47;
        }()), _1605_v22);
        if (_this.f3) {
          let _1608_v25;
          let _init32 = ((_1609_p0) => function (_1610_i3) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(304)), ((_1611_p0) => function (_1612_i4) {
              return ((_this).f2)[_module.__default.safeIndex(_1611_p0, new BigNumber(((_this).f2).length))];
            })(_1609_p0));
          })(p0);
          let _nw200 = Array((new BigNumber(13)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw200.length); _i0_32++) {
            _nw200[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1608_v25 = _nw200;
          let _index233 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_1608_v25).length));
          (_1608_v25)[_index233] = _dafny.Seq.Concat((_this).f2, (_this).f2);
          let _index234 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_1608_v25).length));
          (_1608_v25)[_index234] = (_this).f2;
          r1 = (_module.__default.safeDivisionInt((_this).fm12(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p1)).length), new BigNumber(((_this).f2).length), ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))], globalState), (_this).f4)).multipliedBy((_this).f4);
          let _index235 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index235] = _this.f3;
          (_this).f3 = (_this).fm11(_this.f3, globalState);
          let _1613_v26;
          _1613_v26 = _dafny.Set.fromElements(_1598_v16);
          let _1614_v27;
          _1614_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1598_v16,p0);
          let _index236 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index236] = (_1613_v26).IsDisjointFrom(function () {
            let _coll48 = new _dafny.Set();
            for (const _compr_48 of (_1614_v27).Keys.Elements) {
              let _1615_v28 = _compr_48;
              if ((_1614_v27).contains(_1615_v28)) {
                _coll48.add(_1615_v28);
              }
            }
            return _coll48;
          }());
        } else {
          let _index237 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index237] = _this.f3;
          let _1616_v29;
          let _nw201 = Array((new BigNumber(22)).toNumber());
          _nw201[(_dafny.ZERO).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(_dafny.ONE).toNumber()] = _this.f3;
          _nw201[(new BigNumber(2)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(3)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(4)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(5)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(6)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(7)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(8)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(9)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(10)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(11)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(12)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(13)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(14)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(15)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(16)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _nw201[(new BigNumber(17)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(18)).toNumber()] = _module.__default.fm3(p0, globalState);
          _nw201[(new BigNumber(19)).toNumber()] = _this.f3;
          _nw201[(new BigNumber(20)).toNumber()] = !(!(_this.f3));
          _nw201[(new BigNumber(21)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))];
          _1616_v29 = _nw201;
          let _1617_v30;
          _1617_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1616_v29,_this.f3);
          let _1618_v31;
          let _1619_v32;
          let _1620_v33;
          let _1621_v34;
          let _out29;
          let _out30;
          let _out31;
          let _out32;
          let _outcollector7 = _module.__default.m0(_1617_v30, globalState);
          _out29 = _outcollector7[0];
          _out30 = _outcollector7[1];
          _out31 = _outcollector7[2];
          _out32 = _outcollector7[3];
          _1618_v31 = _out29;
          _1619_v32 = _out30;
          _1620_v33 = _out31;
          _1621_v34 = _out32;
          let _1622_v35;
          _1622_v35 = _module.D3.create_DC10(p1, new BigNumber(((_this).f2).length));
          _1622_v35 = _1622_v35;
          let _1623_v36;
          _1623_v36 = _module.D6.create_DC17(((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))]);
          let _1624_v37;
          _1624_v37 = _dafny.MultiSet.fromElements(!(((_this).f5)[_module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length))]), _1619_v32, (_1623_v36).dtor_cf28);
          _1624_v37 = _1624_v37;
          let _1625_v38;
          _1625_v38 = _module.D5.create_DC14(_1604_v21);
          let _1626_v39;
          let _nw202 = new _module.C0();
          _nw202.__ctor((_1625_v38).dtor_cf23);
          _1626_v39 = _nw202;
        }
        let _1627_v40;
        let _nw203 = new _module.C1();
        _nw203.__ctor();
        _1627_v40 = _nw203;
        let _index238 = _module.__default.safeIndex(new BigNumber(703), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index238] = (_this).fm13(globalState);
      }
      let _1628_v41;
      _1628_v41 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1629_v42;
      _1629_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1628_v41,((true) ? (new BigNumber(-615)) : (p1)));
      _1629_v42 = (_1629_v42).update(((_this).f2)[_module.__default.safeIndex((_this).f4, new BigNumber(((_this).f2).length))], (_this).f4);
      let _1630_v43;
      let _nw204 = new _module.C2();
      _nw204.__ctor((_this).f4, _this.f3, (_this).f1, (_this).f2);
      _1630_v43 = _nw204;
      let _1631_v44;
      _1631_v44 = _module.D13.create_DC33(_1630_v43);
      let _source30 = _1631_v44;
      if (_source30.is_DC34) {
        let _1632___mcc_h0 = (_source30).cf53;
        let _1633___mcc_h1 = (_source30).cf54;
        let _1634___mcc_h2 = (_source30).cf55;
        let _1635___mcc_h3 = (_source30).cf56;
        let _1636_cf56 = _1635___mcc_h3;
        let _1637_cf55 = _1634___mcc_h2;
        let _1638_cf54 = _1633___mcc_h1;
        let _1639_cf53 = _1632___mcc_h0;
        let _index239 = _module.__default.safeIndex(new BigNumber(504), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index239] = _1638_cf54;
        let _index240 = _module.__default.safeIndex(new BigNumber(504), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index240] = !_dafny.Seq.contains((_this).f2, _1628_v41);
        let _1640_v45;
        _1640_v45 = _dafny.Set.fromElements(_1628_v41);
        let _1641_v46;
        _1641_v46 = _dafny.Seq.of(_1640_v45);
        let _1642_v47;
        _1642_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1637_cf55,new BigNumber(((_1641_v46)[_module.__default.safeIndex(p1, new BigNumber((_1641_v46).length))]).length));
        _1642_v47 = (_1642_v47).update(_module.__default.fm0(globalState), (_this.f12).minus(_1637_cf55));
        _1628_v41 = _1628_v41;
        (_this).f3 = _this.f3;
      } else if (_source30.is_DC35) {
        let _1643___mcc_h4 = (_source30).cf57;
        let _1644___mcc_h5 = (_source30).cf58;
        let _1645___mcc_h6 = (_source30).cf59;
        let _1646___mcc_h7 = (_source30).cf60;
        let _1647___mcc_h8 = (_source30).cf61;
        let _1648_cf61 = _1647___mcc_h8;
        let _1649_cf60 = _1646___mcc_h7;
        let _1650_cf59 = _1645___mcc_h6;
        let _1651_cf58 = _1644___mcc_h5;
        let _1652_cf57 = _1643___mcc_h4;
        (_this).f3 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat((_this).f2, (_this).f2), (_this).f2);
        let _1653_v48;
        let _nw205 = new _module.C4();
        _nw205.__ctor(_dafny.Seq.contains((_this).f2, _1628_v41), (_this).f1, (_this).f2);
        _1653_v48 = _nw205;
        (_this).f3 = _1650_cf59;
        let _1654_v49;
        let _nw206 = new _module.C4();
        _nw206.__ctor(!(_this.f3), (_this).f1, (_this).f2);
        _1654_v49 = _nw206;
      } else if (_source30.is_DC33) {
        let _1655___mcc_h9 = (_source30).cf52;
        let _1656_cf52 = _1655___mcc_h9;
        let _1657_v50;
        _1657_v50 = _dafny.Seq.UnicodeFromString("agpskpj");
        _1657_v50 = _1657_v50;
        let _1658_v51;
        let _nw207 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1658_v51 = _nw207;
        _1658_v51 = _1658_v51;
        (_this).f3 = (_1630_v43).fm11(_this.f3, globalState);
        let _1659_v52;
        _1659_v52 = _module.D6.create_DC18(_this.f12, _this.f12, _this.f3, _1628_v41, _this.f12);
        let _1660_v53;
        _1660_v53 = _module.D6.create_DC20(_1659_v52);
        _1660_v53 = _1660_v53;
      } else {
        let _1661___mcc_h10 = (_source30).cf62;
        let _1662_cf62 = _1661___mcc_h10;
        (_this).f12 = new BigNumber(-574);
        (_this).f3 = _this.f3;
        (_this).f12 = (_this).f11;
        let _1663_v54;
        _1663_v54 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        let _1664_v55;
        _1664_v55 = _dafny.MultiSet.fromElements((_this).f2);
        (_this).f3 = (p1).isEqualTo(((_this.f3) ? (new BigNumber((_1663_v54).length)) : ((_dafny.ZERO).minus((_1630_v43).fm27(p1, _this.f3, _1664_v55, globalState)))));
      }
      r1 = _module.__default.safeDivisionInt(p0, _this.f12);
      r1 = new BigNumber(575);
      (_this).f3 = (_this.f3) || (_this.f3);
      let _1665_v56;
      _1665_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
      r0 = _1665_v56;
      r1 = (_dafny.ZERO).minus((p1).plus((_1630_v43).f13));
      r2 = (_module.__default.safeModuloInt(p0, (_this).f11)).multipliedBy((_1630_v43).f13);
      return [r0, r1, r2];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      (_this).f12 = _module.__default.safeModuloInt(p2, _module.__default.safeDivisionInt((_this).f11, new BigNumber(612)));
      let _1666_v0;
      _1666_v0 = _module.D6.create_DC19((_this).f11);
      let _1667_v1;
      _1667_v1 = _module.D13.create_DC34(p0, _this.f3, ((_this.f3) ? ((_1666_v0).dtor_cf34) : (_this.f12)), _module.__default.fm2((_this).fm13(globalState), globalState));
      let _source31 = _1667_v1;
      if (_source31.is_DC34) {
        let _1668___mcc_h0 = (_source31).cf53;
        let _1669___mcc_h1 = (_source31).cf54;
        let _1670___mcc_h2 = (_source31).cf55;
        let _1671___mcc_h3 = (_source31).cf56;
        let _1672_cf56 = _1671___mcc_h3;
        let _1673_cf55 = _1670___mcc_h2;
        let _1674_cf54 = _1669___mcc_h1;
        let _1675_cf53 = _1668___mcc_h0;
        let _1676_v2;
        _1676_v2 = new _dafny.CodePoint('j'.codePointAt(0));
        _1672_cf56 = _dafny.Seq.update(_dafny.Seq.Concat((_this).f2, _dafny.Seq.Concat(_1672_cf56, (_this).f2)), _module.__default.safeIndex((_this).f4, new BigNumber((_dafny.Seq.Concat((_this).f2, _dafny.Seq.Concat(_1672_cf56, (_this).f2))).length)), _1676_v2);
        if (_1675_cf53) {
          let _1677_v3;
          _1677_v3 = _module.D2.create_DC6(_1676_v2, p0, _1675_cf53, _1673_cf55);
          let _1678_v4;
          _1678_v4 = _module.D5.create_DC15((new BigNumber(-841)).multipliedBy((_this).fm12(_this.f12, (_this).f11, _this.f12, (_1677_v3).dtor_cf10, globalState)), _1675_cf53, (_this).f5);
          let _1679_v5;
          _1679_v5 = _dafny.Seq.of((_this).f5, (_this).f5, (_this).f5, (_this).f5);
          _1678_v4 = function (_pat_let60_0) {
            return function (_1680_dt__update__tmp_h0) {
              return function (_pat_let61_0) {
                return function (_1681_dt__update_hcf24_h0) {
                  return _module.D5.create_DC15(_1681_dt__update_hcf24_h0, (_1680_dt__update__tmp_h0).dtor_cf25, (_1680_dt__update__tmp_h0).dtor_cf26);
                }(_pat_let61_0);
              }((_this).f4);
            }(_pat_let60_0);
          }(_module.D5.create_DC15(_module.__default.fm0(globalState), _1675_cf53, (_1679_v5)[_module.__default.safeIndex(_this.f12, new BigNumber((_1679_v5).length))]));
          (_this).f3 = _1675_cf53;
          let _1682_v6;
          _1682_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
          let _1683_v7;
          let _1684_v8;
          let _1685_v9;
          let _1686_v10;
          let _out33;
          let _out34;
          let _out35;
          let _out36;
          let _outcollector8 = _module.__default.m0((_1682_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1675_cf53)), globalState);
          _out33 = _outcollector8[0];
          _out34 = _outcollector8[1];
          _out35 = _outcollector8[2];
          _out36 = _outcollector8[3];
          _1683_v7 = _out33;
          _1684_v8 = _out34;
          _1685_v9 = _out35;
          _1686_v10 = _out36;
          let _1687_v11;
          let _nw208 = new _module.C1();
          _nw208.__ctor();
          _1687_v11 = _nw208;
          let _1688_v12;
          let _init33 = ((_1689_cf56) => function (_1690_i0) {
            return (_module.D1.create_DC2(_1689_cf56)).dtor_cf2;
          })(_1672_cf56);
          let _nw209 = Array((new BigNumber(14)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw209.length); _i0_33++) {
            _nw209[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1688_v12 = _nw209;
          let _nw210 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1688_v12 = _nw210;
        } else {
          let _1691_v13;
          let _nw211 = new _module.C4();
          _nw211.__ctor(_1675_cf53, (_this).f1, _dafny.Seq.UnicodeFromString("wttr"));
          _1691_v13 = _nw211;
          let _1692_v14;
          _1692_v14 = _dafny.Seq.of(_1673_cf55, _1673_cf55);
          (_this).f12 = (_1692_v14)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_1692_v14).length))];
          let _1693_v15;
          let _nw212 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
          _1693_v15 = _nw212;
          let _1694_v16;
          _1694_v16 = _dafny.MultiSet.fromElements(_1693_v15);
          let _1695_v17;
          _1695_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(((_1694_v16).contains(_1693_v15)) ? ((_1694_v16).get(_1693_v15)) : (new BigNumber((_dafny.Seq.update(_1692_v14, _module.__default.safeIndex(_this.f12, new BigNumber((_1692_v14).length)), _this.f12)).length))));
          let _1696_v18;
          _1696_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3);
          _1673_cf55 = (_module.__default.safeModuloInt(_1673_cf55, (((_1695_v17).contains(_1672_cf56)) ? ((_1695_v17).get(_1672_cf56)) : (_1673_cf55)))).plus(new BigNumber((_1696_v18).length));
          let _1697_v19;
          let _nw213 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _1697_v19 = _nw213;
          let _1698_v20;
          _1698_v20 = _dafny.MultiSet.fromElements(p0);
          let _1699_v21;
          _1699_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1698_v20,p0);
          let _index241 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_1697_v19).length));
          (_1697_v19)[_index241] = (_1699_v21).update(_1698_v20, p0);
          let _1700_v23;
          _1700_v23 = _dafny.Seq.of(_1698_v20);
          let _index242 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_1697_v19).length));
          let _rhs242 = (_dafny.ZERO).minus(_1673_cf55);
          let _rhs243 = p2;
          let _rhs244 = p2;
          let _rhs245 = function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_dafny.Seq.Concat(_1700_v23, _1700_v23)).Elements) {
              let _1701_v22 = _compr_49;
              if (_dafny.Seq.contains(_dafny.Seq.Concat(_1700_v23, _1700_v23), _1701_v22)) {
                _coll49.push([_1701_v22,_dafny.Seq.IsProperPrefixOf(_1672_cf56, (_this).f2)]);
              }
            }
            return _coll49;
          }();
          let _lhs136 = _this;
          let _lhs137 = _1697_v19;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_1697_v19).length));
          _1673_cf55 = _rhs242;
          _1673_cf55 = _rhs243;
          _lhs136.f12 = _rhs244;
          _lhs137[_lhs138] = _rhs245;
          let _1702_v24;
          let _nw214 = new _module.C1();
          _nw214.__ctor();
          _1702_v24 = _nw214;
        }
        _1667_v1 = _1667_v1;
        let _1703_v25;
        _1703_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3);
        _1703_v25 = function () {
          let _coll50 = new _dafny.Map();
          for (const _compr_50 of (_dafny.Map.Empty.slice().updateUnsafe((_this).fm14(p0, p0, new _dafny.CodePoint('w'.codePointAt(0)), (_this).f4, globalState),_this.f12)).Keys.Elements) {
            let _1704_v26 = _compr_50;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_this).fm14(p0, p0, new _dafny.CodePoint('w'.codePointAt(0)), (_this).f4, globalState),_this.f12)).contains(_1704_v26)) {
              _coll50.push([_1704_v26,(_module.D3.create_DC9(_1673_cf55, p0)).dtor_cf16]);
            }
          }
          return _coll50;
        }();
      } else if (_source31.is_DC35) {
        let _1705___mcc_h4 = (_source31).cf57;
        let _1706___mcc_h5 = (_source31).cf58;
        let _1707___mcc_h6 = (_source31).cf59;
        let _1708___mcc_h7 = (_source31).cf60;
        let _1709___mcc_h8 = (_source31).cf61;
        let _1710_cf61 = _1709___mcc_h8;
        let _1711_cf60 = _1708___mcc_h7;
        let _1712_cf59 = _1707___mcc_h6;
        let _1713_cf58 = _1706___mcc_h5;
        let _1714_cf57 = _1705___mcc_h4;
        let _index243 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index243] = _1712_cf59;
        let _1715_v27;
        _1715_v27 = _dafny.MultiSet.fromElements(_this.f3);
        let _1716_v28;
        _1716_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),(((p1).contains(p2)) ? ((p1).get(p2)) : (_module.__default.fm0(globalState))));
        let _1717_v29;
        _1717_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1716_v28).length),_1715_v27);
        let _1718_v30;
        _1718_v30 = _dafny.Seq.of(_1715_v27, _1715_v27);
        let _1719_v31;
        let _nw215 = Array((new BigNumber(13)).toNumber());
        _nw215[(_dafny.ZERO).toNumber()] = _1715_v27;
        _nw215[(_dafny.ONE).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(2)).toNumber()] = (_1715_v27).Intersect(_1715_v27);
        _nw215[(new BigNumber(3)).toNumber()] = (((_1717_v29).contains((_this).f4)) ? ((_1717_v29).get((_this).f4)) : (_1715_v27));
        _nw215[(new BigNumber(4)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(5)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(6)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(7)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(8)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(9)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_1712_cf59);
        _nw215[(new BigNumber(11)).toNumber()] = _1715_v27;
        _nw215[(new BigNumber(12)).toNumber()] = (_1715_v27).Union((_1718_v30)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1718_v30).length))]);
        _1719_v31 = _nw215;
        let _index244 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1719_v31).length));
        (_1719_v31)[_index244] = _1715_v27;
        let _1720_v32;
        let _init34 = ((_1721_p2) => function (_1722_i1) {
          return _module.__default.safeDivisionInt(_1722_i1, _1721_p2);
        })(p2);
        let _nw216 = Array((new BigNumber(17)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw216.length); _i0_34++) {
          _nw216[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1720_v32 = _nw216;
        let _index245 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1720_v32).length));
        (_1720_v32)[_index245] = _module.__default.fm0(globalState);
        let _1723_v33;
        _1723_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ktq"),_1711_cf60);
        let _1724_v34;
        _1724_v34 = new _dafny.CodePoint('f'.codePointAt(0));
        let _index246 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
        let _index247 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1719_v31).length));
        let _index248 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1720_v32).length));
        let _rhs246 = (_this).fm13(globalState);
        let _rhs247 = !(_1712_cf59);
        let _rhs248 = (_1715_v27).Intersect(_module.__default.fm5((_1723_v33).update((_this).fm14(_1712_cf59, _this.f3, _1724_v34, _this.f12, globalState), (_this).f11), globalState));
        let _rhs249 = p2;
        let _lhs139 = (_this).f5;
        let _lhs140 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
        let _lhs141 = _this;
        let _lhs142 = _1719_v31;
        let _lhs143 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_1719_v31).length));
        let _lhs144 = _1720_v32;
        let _lhs145 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1720_v32).length));
        _lhs139[_lhs140] = _rhs246;
        _lhs141.f3 = _rhs247;
        _lhs142[_lhs143] = _rhs248;
        _lhs144[_lhs145] = _rhs249;
        let _1725_v35;
        let _nw217 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1725_v35 = _nw217;
        let _1726_v37;
        let _nw218 = new _module.C3();
        _nw218.__ctor((_this).f11, _1725_v35, !(((_this).f5)[_module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length))]), ((_this).f1).Merge(function () {
          let _coll51 = new _dafny.Map();
          for (const _compr_51 of (_1723_v33).Keys.Elements) {
            let _1727_v36 = _compr_51;
            if ((_1723_v33).contains(_1727_v36)) {
              _coll51.push([_1727_v36,p0]);
            }
          }
          return _coll51;
        }()), _dafny.Seq.Concat((_this).f2, (_this).f2));
        _1726_v37 = _nw218;
        let _1728_v38;
        _1728_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1711_cf60,p0);
        let _1729_v39;
        _1729_v39 = _dafny.Seq.of(_1728_v38);
        let _1730_v40;
        _1730_v40 = _dafny.Seq.of(_1711_cf60, new BigNumber((_1729_v39).length), p2, _1711_cf60, (_this).f4);
        let _pat_let_tv44 = _1730_v40;
        let _pat_let_tv45 = _1730_v40;
        let _source32 = function (_pat_let62_0) {
          return function (_1731_dt__update__tmp_h1) {
            return function (_pat_let63_0) {
              return function (_1732_dt__update_hcf24_h1) {
                return _module.D5.create_DC15(_1732_dt__update_hcf24_h1, (_1731_dt__update__tmp_h1).dtor_cf25, (_1731_dt__update__tmp_h1).dtor_cf26);
              }(_pat_let63_0);
            }((_pat_let_tv44)[_module.__default.safeIndex(new BigNumber(511), new BigNumber((_pat_let_tv45).length))]);
          }(_pat_let62_0);
        }(_module.D5.create_DC15(_1714_cf57, (_1726_v37).fm11(!((_1726_v37).fm13(globalState)), globalState), _1725_v35));
        if (_source32.is_DC15) {
          let _1733___mcc_h11 = (_source32).cf24;
          let _1734___mcc_h12 = (_source32).cf25;
          let _1735___mcc_h13 = (_source32).cf26;
          let _1736_cf26 = _1735___mcc_h13;
          let _1737_cf25 = _1734___mcc_h12;
          let _1738_cf24 = _1733___mcc_h11;
          let _1739_v42;
          _1739_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f11, _1714_cf57),function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of (_dafny.Set.fromElements(p1, p1, p1, p1, p1)).Elements) {
              let _1740_v41 = _compr_52;
              if ((_dafny.Set.fromElements(p1, p1, p1, p1, p1)).contains(_1740_v41)) {
                _coll52.add(_1740_v41);
              }
            }
            return _coll52;
          }());
          _1739_v42 = (_1739_v42).update((_this).f11, (function () {
            let _coll53 = new _dafny.Set();
            for (const _compr_53 of (_dafny.Set.fromElements(p1, p1)).Elements) {
              let _1741_v43 = _compr_53;
              if ((_dafny.Set.fromElements(p1, p1)).contains(_1741_v43)) {
                _coll53.add(_1741_v43);
              }
            }
            return _coll53;
          }()).Union(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_1730_v40))));
          let _index249 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index249] = _this.f3;
          _1716_v28 = _1716_v28;
          let _1742_v44;
          let _nw219 = new _module.C0();
          _nw219.__ctor(_1720_v32);
          _1742_v44 = _nw219;
        } else {
          let _1743___mcc_h14 = (_source32).cf23;
          let _1744_cf23 = _1743___mcc_h14;
          let _1745_v45;
          let _nw220 = new _module.C0();
          _nw220.__ctor(_1720_v32);
          _1745_v45 = _nw220;
          let _index250 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index250] = p0;
          let _index251 = _module.__default.safeIndex(new BigNumber(830), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index251] = p0;
          (_this).f3 = _1712_cf59;
        }
        let _index252 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_1720_v32).length));
        (_1720_v32)[_index252] = p2;
      } else if (_source31.is_DC33) {
        let _1746___mcc_h9 = (_source31).cf52;
        let _1747_cf52 = _1746___mcc_h9;
        let _rhs250 = (_this).f4;
        let _rhs251 = !(true);
        let _rhs252 = (_1747_cf52).f13;
        let _lhs146 = _this;
        let _lhs147 = _this;
        let _lhs148 = _this;
        _lhs146.f12 = _rhs250;
        _lhs147.f3 = _rhs251;
        _lhs148.f12 = _rhs252;
        (_this).f12 = (_module.__default.safeModuloInt((_this).f4, (_1747_cf52).f13)).multipliedBy(p2);
        let _index253 = _module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index253] = p0;
        let _1748_v46;
        _1748_v46 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1749_v47;
        _1749_v47 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f3);
        let _1750_v48;
        _1750_v48 = _dafny.Seq.of((_1747_cf52).f13);
        let _1751_v49;
        _1751_v49 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1750_v48), p1, _dafny.MultiSet.fromElements(new BigNumber(((_this).f2).length), (_dafny.ZERO).minus((_this).f4)), _dafny.MultiSet.fromElements((_1747_cf52).f13));
        let _1752_v50;
        _1752_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1751_v49)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_1751_v49).length))],p0);
        let _1753_v51;
        _1753_v51 = _module.D4.create_DC13(_this.f3, _1750_v48);
        let _1754_v52;
        _1754_v52 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f3);
        let _index254 = _module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length));
        let _rhs253 = !(!(new BigNumber(238)).isEqualTo((new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex(_this.f12, new BigNumber(((_this).f2).length)), _1748_v46)).length)).multipliedBy(new BigNumber(((_1749_v47).update(new BigNumber(804), p0)).length))));
        let _rhs254 = false;
        let _rhs255 = (_dafny.ZERO).minus(((_this).f4).minus(_this.f12));
        let _rhs256 = _this.f3;
        let _rhs257 = !((((_1752_v50).contains(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update((_1753_v51).dtor_cf22, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f4)).length), new BigNumber(((_1753_v51).dtor_cf22).length)), (_1747_cf52).f13), _1750_v48)))) ? ((_1752_v50).get(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update((_1753_v51).dtor_cf22, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f4)).length), new BigNumber(((_1753_v51).dtor_cf22).length)), (_1747_cf52).f13), _1750_v48)))) : (!((_1754_v52).update(p0, p0)).equals(_1754_v52))));
        let _lhs149 = _this;
        let _lhs150 = (_this).f5;
        let _lhs151 = _module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length));
        let _lhs152 = _this;
        let _lhs153 = _this;
        let _lhs154 = _this;
        _lhs149.f3 = _rhs253;
        _lhs150[_lhs151] = _rhs254;
        _lhs152.f12 = _rhs255;
        _lhs153.f3 = _rhs256;
        _lhs154.f3 = _rhs257;
        let _1755_v53;
        let _nw221 = Array((new BigNumber(16)).toNumber());
        _nw221[(_dafny.ZERO).toNumber()] = false;
        _nw221[(_dafny.ONE).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length))];
        _nw221[(new BigNumber(2)).toNumber()] = p0;
        _nw221[(new BigNumber(3)).toNumber()] = _this.f3;
        _nw221[(new BigNumber(4)).toNumber()] = p0;
        _nw221[(new BigNumber(5)).toNumber()] = p0;
        _nw221[(new BigNumber(6)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length))];
        _nw221[(new BigNumber(7)).toNumber()] = _this.f3;
        _nw221[(new BigNumber(8)).toNumber()] = p0;
        _nw221[(new BigNumber(9)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length))];
        _nw221[(new BigNumber(10)).toNumber()] = p0;
        _nw221[(new BigNumber(11)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length))];
        _nw221[(new BigNumber(12)).toNumber()] = p0;
        _nw221[(new BigNumber(13)).toNumber()] = true;
        _nw221[(new BigNumber(14)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length))];
        _nw221[(new BigNumber(15)).toNumber()] = !(p0);
        _1755_v53 = _nw221;
        let _1756_v54;
        _1756_v54 = _dafny.Seq.of(_1755_v53);
        let _index255 = _module.__default.safeIndex(new BigNumber(837), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index255] = _dafny.areEqual(_dafny.Seq.Concat(_1756_v54, _1756_v54), _dafny.Seq.update(_dafny.Seq.Concat(_1756_v54, _1756_v54), _module.__default.safeIndex(_this.f12, new BigNumber((_dafny.Seq.Concat(_1756_v54, _1756_v54)).length)), _1755_v53));
      } else {
        let _1757___mcc_h10 = (_source31).cf62;
        let _1758_cf62 = _1757___mcc_h10;
        if (_this.f3) {
          (_this).f12 = (_this).fm21(_this.f3, _this.f12, globalState);
          (_this).f12 = _module.__default.fm0(globalState);
          let _index256 = _module.__default.safeIndex(new BigNumber(822), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index256] = _this.f3;
          let _1759_v55;
          _1759_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_this.f3);
          let _1760_v56;
          _1760_v56 = _dafny.Set.fromElements(p0);
          let _index257 = _module.__default.safeIndex(new BigNumber(822), new BigNumber(((_this).f5).length));
          let _rhs258 = (_dafny.ZERO).minus(new BigNumber(((_1759_v55).Merge(_1759_v55)).length));
          let _rhs259 = (_this).fm21(_this.f3, p2, globalState);
          let _rhs260 = _this.f3;
          let _rhs261 = (_1760_v56).equals(_1760_v56);
          let _lhs155 = _this;
          let _lhs156 = _this;
          let _lhs157 = (_this).f5;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(822), new BigNumber(((_this).f5).length));
          let _lhs159 = _this;
          _lhs155.f12 = _rhs258;
          _lhs156.f12 = _rhs259;
          _lhs157[_lhs158] = _rhs260;
          _lhs159.f3 = _rhs261;
          let _index258 = _module.__default.safeIndex(new BigNumber(822), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index258] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(822), new BigNumber(((_this).f5).length))];
          (_this).f12 = new BigNumber((_1759_v55).length);
        } else {
          let _1761_v57;
          _1761_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
          let _1762_v58;
          let _1763_v59;
          let _1764_v60;
          let _1765_v61;
          let _out37;
          let _out38;
          let _out39;
          let _out40;
          let _outcollector9 = _module.__default.m0(_1761_v57, globalState);
          _out37 = _outcollector9[0];
          _out38 = _outcollector9[1];
          _out39 = _outcollector9[2];
          _out40 = _outcollector9[3];
          _1762_v58 = _out37;
          _1763_v59 = _out38;
          _1764_v60 = _out39;
          _1765_v61 = _out40;
          (_this).f3 = _1762_v58;
          let _1766_v62;
          let _nw222 = new _module.C4();
          _nw222.__ctor(_1763_v59, (_this).f1, _dafny.Seq.UnicodeFromString("dph"));
          _1766_v62 = _nw222;
          let _1767_v63;
          _1767_v63 = _dafny.Set.fromElements(p2);
          let _1768_v64;
          _1768_v64 = _dafny.Seq.of((_this).f4, _this.f12, (_this).f11, (_dafny.ZERO).minus((_this).f4), (new BigNumber((_1767_v63).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,p0)).length)));
          (_this).f12 = new BigNumber((_1768_v64).length);
          let _1769_v65;
          let _init35 = ((_1770_v58) => function (_1771_i2) {
            return _1770_v58;
          })(_1762_v58);
          let _nw223 = Array((new BigNumber(16)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw223.length); _i0_35++) {
            _nw223[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1769_v65 = _nw223;
          let _1772_v66;
          let _nw224 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1772_v66 = _nw224;
          let _1773_v67;
          _1773_v67 = _module.D14.create_DC37(_1772_v66);
          let _1774_v68;
          let _nw225 = Array((new BigNumber(26)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1772_v66;
          _nw225[(_dafny.ONE).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(2)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(3)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(4)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(5)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(6)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(7)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(8)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(9)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(10)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(11)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(12)).toNumber()] = (_1773_v67).dtor_cf63;
          _nw225[(new BigNumber(13)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(14)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(15)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(16)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(17)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(18)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(19)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(20)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(21)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(22)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(23)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(24)).toNumber()] = _1772_v66;
          _nw225[(new BigNumber(25)).toNumber()] = _1772_v66;
          _1774_v68 = _nw225;
          let _1775_v69;
          _1775_v69 = _1774_v68;
          let _1776_v70;
          _1776_v70 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs262 = !_dafny.Seq.contains(_dafny.Seq.of(_1776_v70, _1776_v70), new _dafny.CodePoint('h'.codePointAt(0)));
          let _rhs263 = (_this).f5;
          let _rhs264 = _1775_v69;
          _1762_v58 = _rhs262;
          _1769_v65 = _rhs263;
          _1775_v69 = _rhs264;
        }
        let _1777_v71;
        _1777_v71 = _dafny.Set.fromElements(_this.f3);
        let _1778_v72;
        _1778_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1777_v71);
        let _1779_v73;
        let _nw226 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1779_v73 = _nw226;
        let _1780_v74;
        _1780_v74 = _dafny.Seq.of(_1779_v73);
        let _1781_v75;
        let _nw227 = Array((new BigNumber(26)).toNumber());
        _nw227[(_dafny.ZERO).toNumber()] = _1777_v71;
        _nw227[(_dafny.ONE).toNumber()] = (_1777_v71).Difference(_1777_v71);
        _nw227[(new BigNumber(2)).toNumber()] = (_1777_v71).Intersect(_dafny.Set.fromElements(_this.f3));
        _nw227[(new BigNumber(3)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(4)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(5)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(6)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(7)).toNumber()] = (_1777_v71).Intersect(_1777_v71);
        _nw227[(new BigNumber(8)).toNumber()] = (_1777_v71).Union(_1777_v71);
        _nw227[(new BigNumber(9)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(10)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(11)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(_this.f3, p0);
        _nw227[(new BigNumber(13)).toNumber()] = (_1777_v71).Difference(_1777_v71);
        _nw227[(new BigNumber(14)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(15)).toNumber()] = (_1777_v71).Union(_1777_v71);
        _nw227[(new BigNumber(16)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(17)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(18)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(19)).toNumber()] = _dafny.Set.fromElements(_this.f3, _this.f3);
        _nw227[(new BigNumber(20)).toNumber()] = (_1777_v71).Union(_1777_v71);
        _nw227[(new BigNumber(21)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(22)).toNumber()] = (((_1778_v72).contains(new BigNumber((_1780_v74).length))) ? ((_1778_v72).get(new BigNumber((_1780_v74).length))) : (_dafny.Set.fromElements(p0)));
        _nw227[(new BigNumber(23)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(24)).toNumber()] = _1777_v71;
        _nw227[(new BigNumber(25)).toNumber()] = _1777_v71;
        _1781_v75 = _nw227;
        let _index259 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1781_v75).length));
        (_1781_v75)[_index259] = _dafny.Set.fromElements(p0);
        let _index260 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_1781_v75).length));
        (_1781_v75)[_index260] = (_1777_v71).Intersect((_1777_v71).Difference(_dafny.Set.fromElements(p0)));
        _1777_v71 = _1777_v71;
        let _1782_v76;
        _1782_v76 = _module.D3.create_DC8((_1781_v75)[_module.__default.safeIndex(new BigNumber(793), new BigNumber((_1781_v75).length))]);
        let _1783_v77;
        _1783_v77 = _module.D3.create_DC11(_1782_v76);
        let _1784_v78;
        _1784_v78 = _dafny.Seq.of(_this.f3, !(p0), !(p0));
        let _1785_v79;
        _1785_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1784_v78);
        _1783_v77 = _module.__default.fm40(new BigNumber(((_this).f2).length), _this.f3, new BigNumber(((((_1785_v79).contains(_this.f3)) ? ((_1785_v79).get(_this.f3)) : (_1784_v78))).length), (_dafny.ZERO).minus((_this.f12).plus(_this.f12)), globalState);
      }
      if (_this.f3) {
        let _1786_v80;
        _1786_v80 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1787_v81;
        _1787_v81 = _dafny.Set.fromElements(p0, _this.f3, _this.f3, p0, false);
        let _1788_v82;
        _1788_v82 = _module.D3.create_DC8(_1787_v81);
        let _1789_v83;
        _1789_v83 = _dafny.Set.fromElements(_module.D3.create_DC8(_1787_v81), _1788_v82);
        let _1790_v84;
        _1790_v84 = _dafny.Seq.of(_this.f3, p0);
        _1786_v80 = _module.__default.fm38(!(!((_1789_v83).IsDisjointFrom(_1789_v83))), _dafny.Seq.IsProperPrefixOf(_1790_v84, _1790_v84), new BigNumber(769), globalState);
        let _1791_v85;
        let _nw228 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1791_v85 = _nw228;
        let _index261 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        (_1791_v85)[_index261] = (_this).f4;
        let _index262 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        (_1791_v85)[_index262] = _this.f12;
        let _index263 = _module.__default.safeIndex(new BigNumber(705), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index263] = p0;
        let _index264 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        let _index265 = _module.__default.safeIndex(new BigNumber(705), new BigNumber(((_this).f5).length));
        let _index266 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        let _rhs265 = p2;
        let _rhs266 = false;
        let _rhs267 = (_this.f12).multipliedBy(_module.__default.safeModuloInt(new BigNumber(807), _module.__default.fm0(globalState)));
        let _lhs160 = _1791_v85;
        let _lhs161 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        let _lhs162 = (_this).f5;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(705), new BigNumber(((_this).f5).length));
        let _lhs164 = _1791_v85;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        _lhs160[_lhs161] = _rhs265;
        _lhs162[_lhs163] = _rhs266;
        _lhs164[_lhs165] = _rhs267;
        (_this).f3 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(705), new BigNumber(((_this).f5).length))];
        let _index267 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        let _rhs268 = new BigNumber(240);
        let _rhs269 = _1791_v85;
        let _lhs166 = _1791_v85;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(951), new BigNumber((_1791_v85).length));
        _lhs166[_lhs167] = _rhs268;
        _1791_v85 = _rhs269;
      } else {
        let _1792_v86;
        _1792_v86 = _dafny.Set.fromElements(p0);
        let _1793_v87;
        _1793_v87 = _module.D3.create_DC8(_1792_v86);
        let _1794_v88;
        _1794_v88 = _module.D3.create_DC11(_module.D3.create_DC11(_1793_v87));
        let _source33 = _1794_v88;
        if (_source33.is_DC9) {
          let _1795___mcc_h15 = (_source33).cf15;
          let _1796___mcc_h16 = (_source33).cf16;
          let _1797_cf16 = _1796___mcc_h16;
          let _1798_cf15 = _1795___mcc_h15;
          let _1799_v89;
          _1799_v89 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_1798_cf15);
          let _1800_v90;
          _1800_v90 = _dafny.Seq.of(new BigNumber(((_1799_v89).update(_this.f12, _this.f12)).length));
          let _rhs270 = ((new BigNumber(((_this).f2).length)).multipliedBy((_1800_v90)[_module.__default.safeIndex(_this.f12, new BigNumber((_1800_v90).length))])).isEqualTo(_module.__default.safeDivisionInt(p2, new BigNumber(((_this).f2).length)));
          let _rhs271 = _module.__default.fm0(globalState);
          let _rhs272 = _module.__default.safeDivisionInt(new BigNumber(((_1792_v86).Union(_1792_v86)).length), (_this).f4);
          _1797_cf16 = _rhs270;
          _1798_cf15 = _rhs271;
          _1798_cf15 = _rhs272;
          (_this).f3 = false;
          let _1801_v91;
          let _nw229 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1801_v91 = _nw229;
          let _1802_v92;
          _1802_v92 = _module.D14.create_DC38((_this).f2);
          let _index268 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1801_v91).length));
          (_1801_v91)[_index268] = (_1802_v92).dtor_cf64;
          let _1803_v93;
          _1803_v93 = _dafny.Seq.of(_this.f3, _this.f3);
          let _1804_v94;
          _1804_v94 = new _dafny.CodePoint('n'.codePointAt(0));
          let _index269 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1801_v91).length));
          (_1801_v91)[_index269] = _dafny.Seq.Concat((_this).f2, _dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber((_1803_v93).length), new BigNumber(((_this).f2).length)), _1804_v94));
          let _index270 = _module.__default.safeIndex(new BigNumber(766), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index270] = _this.f3;
          let _1805_v95;
          let _init36 = function (_1806_i3) {
            return _module.__default.safeModuloInt(_1806_i3, (_this).f11);
          };
          let _nw230 = Array((new BigNumber(18)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw230.length); _i0_36++) {
            _nw230[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1805_v95 = _nw230;
          let _index271 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1805_v95).length));
          (_1805_v95)[_index271] = p2;
          let _index272 = _module.__default.safeIndex(new BigNumber(766), new BigNumber(((_this).f5).length));
          let _index273 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1805_v95).length));
          let _index274 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1801_v91).length));
          let _rhs273 = false;
          let _rhs274 = !(true);
          let _rhs275 = _1798_cf15;
          let _rhs276 = (_this).f2;
          let _lhs168 = (_this).f5;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(766), new BigNumber(((_this).f5).length));
          let _lhs170 = _1805_v95;
          let _lhs171 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_1805_v95).length));
          let _lhs172 = _1801_v91;
          let _lhs173 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_1801_v91).length));
          _lhs168[_lhs169] = _rhs273;
          _1797_cf16 = _rhs274;
          _lhs170[_lhs171] = _rhs275;
          _lhs172[_lhs173] = _rhs276;
        } else if (_source33.is_DC10) {
          let _1807___mcc_h17 = (_source33).cf17;
          let _1808___mcc_h18 = (_source33).cf18;
          let _1809_cf18 = _1808___mcc_h18;
          let _1810_cf17 = _1807___mcc_h17;
          (_this).f3 = _this.f3;
          let _1811_v96;
          let _nw231 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1811_v96 = _nw231;
          let _1812_v97;
          _1812_v97 = _dafny.Set.fromElements(_this.f12);
          let _rhs277 = _1811_v96;
          let _rhs278 = _1810_cf17;
          let _rhs279 = _1812_v97;
          _1811_v96 = _rhs277;
          _1810_cf17 = _rhs278;
          _1812_v97 = _rhs279;
          let _index275 = _module.__default.safeIndex(new BigNumber(682), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index275] = p0;
          let _1813_v98;
          _1813_v98 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), function (_1814_i4) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }));
          let _1815_v99;
          _1815_v99 = _dafny.Seq.of((_this).f2);
          let _index276 = _module.__default.safeIndex(new BigNumber(682), new BigNumber(((_this).f5).length));
          let _rhs280 = (_this).fm12(_module.__default.fm0(globalState), _1809_cf18, _1809_cf18, p0, globalState);
          let _rhs281 = ((_this.f3) ? (_1809_cf18) : ((_this).fm21(_this.f3, (_this).f4, globalState)));
          let _rhs282 = p0;
          let _rhs283 = ((_this).f11).isLessThan(new BigNumber(((_1813_v98).Union(_dafny.MultiSet.FromArray(_1815_v99))).cardinality()));
          let _lhs174 = (_this).f5;
          let _lhs175 = _module.__default.safeIndex(new BigNumber(682), new BigNumber(((_this).f5).length));
          let _lhs176 = _this;
          _1810_cf17 = _rhs280;
          _1809_cf18 = _rhs281;
          _lhs174[_lhs175] = _rhs282;
          _lhs176.f3 = _rhs283;
          let _1816_v100;
          let _nw232 = Array((new BigNumber(24)).toNumber()).fill(false);
          _1816_v100 = _nw232;
          let _1817_v101;
          _1817_v101 = _dafny.Set.fromElements(_1816_v100);
          let _1818_v102;
          _1818_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1810_cf17,((_this).f5)[_module.__default.safeIndex(new BigNumber(682), new BigNumber(((_this).f5).length))]);
          let _1819_v103;
          _1819_v103 = _dafny.Set.fromElements(_1818_v102);
          let _1820_v104;
          _1820_v104 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1809_cf18,p0),_1810_cf17);
          let _1821_v106;
          _1821_v106 = _dafny.Seq.of(_1819_v103, function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_1820_v104).Keys.Elements) {
              let _1822_v105 = _compr_54;
              if ((_1820_v104).contains(_1822_v105)) {
                _coll54.add(_1822_v105);
              }
            }
            return _coll54;
          }(), _1819_v103);
          let _1823_v107;
          _1823_v107 = _dafny.MultiSet.fromElements(_this.f3);
          let _1824_v108;
          _1824_v108 = _dafny.Map.Empty.slice().updateUnsafe(_1818_v102,_1823_v107);
          let _1825_v110;
          let _nw233 = Array((new BigNumber(10)).toNumber());
          _nw233[(_dafny.ZERO).toNumber()] = (_1819_v103).Intersect(_1819_v103);
          _nw233[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_1818_v102);
          _nw233[(new BigNumber(2)).toNumber()] = (_1821_v106)[_module.__default.safeIndex(_1809_cf18, new BigNumber((_1821_v106).length))];
          _nw233[(new BigNumber(3)).toNumber()] = _1819_v103;
          _nw233[(new BigNumber(4)).toNumber()] = _1819_v103;
          _nw233[(new BigNumber(5)).toNumber()] = _1819_v103;
          _nw233[(new BigNumber(6)).toNumber()] = _1819_v103;
          _nw233[(new BigNumber(7)).toNumber()] = function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of (_1824_v108).Keys.Elements) {
              let _1826_v109 = _compr_55;
              if ((_1824_v108).contains(_1826_v109)) {
                _coll55.add(_1826_v109);
              }
            }
            return _coll55;
          }();
          _nw233[(new BigNumber(8)).toNumber()] = _1819_v103;
          _nw233[(new BigNumber(9)).toNumber()] = _1819_v103;
          _1825_v110 = _nw233;
          let _index277 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1825_v110).length));
          (_1825_v110)[_index277] = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1809_cf18,p0), _1818_v102, _1818_v102, _1818_v102);
          let _index278 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1825_v110).length));
          let _rhs284 = _1817_v101;
          let _rhs285 = (_1819_v103).Union(_dafny.Set.fromElements(_1818_v102, _dafny.Map.Empty.slice().updateUnsafe(_1809_cf18,true), _1818_v102, _1818_v102));
          let _lhs177 = _1825_v110;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1825_v110).length));
          _1817_v101 = _rhs284;
          _lhs177[_lhs178] = _rhs285;
        } else if (_source33.is_DC8) {
          let _1827___mcc_h19 = (_source33).cf14;
          let _1828_cf14 = _1827___mcc_h19;
          (_this).f3 = false;
          let _1829_v111;
          _1829_v111 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f11);
          let _1830_v112;
          _1830_v112 = _dafny.MultiSet.fromElements((_this).f4);
          let _index279 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index279] = _this.f3;
          let _1831_v113;
          _1831_v113 = _module.D6.create_DC17(p0);
          let _1832_v114;
          _1832_v114 = _dafny.Set.fromElements(new BigNumber(258), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-838)), function (_1833_i5) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length));
          let _index280 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_this).f5).length));
          let _rhs286 = ((_1829_v111).Merge(_1829_v111)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(749)));
          let _rhs287 = (_1831_v113).dtor_cf28;
          let _rhs288 = ((_dafny.ZERO).minus((_this).f11)).isEqualTo(_module.__default.fm0(globalState));
          let _rhs289 = (((_1832_v114).IsProperSubsetOf(_1832_v114)) ? (p1) : ((_1830_v112).Union(p1)));
          let _rhs290 = _this.f3;
          let _lhs179 = _this;
          let _lhs180 = _this;
          let _lhs181 = (_this).f5;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(869), new BigNumber(((_this).f5).length));
          _1829_v111 = _rhs286;
          _lhs179.f3 = _rhs287;
          _lhs180.f3 = _rhs288;
          _1830_v112 = _rhs289;
          _lhs181[_lhs182] = _rhs290;
          (_this).f12 = (_this).f4;
          let _1834_v115;
          _1834_v115 = _dafny.Seq.of((_this).f11);
          let _1835_v116;
          _1835_v116 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1836_v117;
          _1836_v117 = _dafny.Seq.of(new BigNumber((_1834_v115).length), (_this).f11, (_this).f11, new BigNumber(((_this).fm14(_this.f3, _this.f3, _1835_v116, (_this).f4, globalState)).length));
          let _1837_v118;
          _1837_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_1836_v117);
          (_this).f12 = new BigNumber(((((_1837_v118).contains((_this).f2)) ? ((_1837_v118).get((_this).f2)) : (_1834_v115))).length);
        } else {
          let _1838___mcc_h20 = (_source33).cf19;
          let _1839_cf19 = _1838___mcc_h20;
          let _1840_v119;
          let _nw234 = Array((new BigNumber(5)).toNumber());
          _nw234[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw234[(_dafny.ONE).toNumber()] = _this.f12;
          _nw234[(new BigNumber(2)).toNumber()] = (_this).f4;
          _nw234[(new BigNumber(3)).toNumber()] = _this.f12;
          _nw234[(new BigNumber(4)).toNumber()] = new BigNumber(317);
          _1840_v119 = _nw234;
          let _1841_v120;
          let _nw235 = new _module.C0();
          _nw235.__ctor(_1840_v119);
          _1841_v120 = _nw235;
          _1841_v120 = _1841_v120;
          let _index281 = _module.__default.safeIndex(new BigNumber(730), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index281] = (p1).IsProperSubsetOf(p1);
          let _index282 = _module.__default.safeIndex(new BigNumber(730), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index282] = !(!(p2).isEqualTo((_this).f11));
          (_this).f12 = (_this).f4;
          let _1842_v121;
          _1842_v121 = new _dafny.CodePoint('o'.codePointAt(0));
          _1842_v121 = new _dafny.CodePoint('g'.codePointAt(0));
        }
        let _1843_v122;
        let _init37 = function (_1844_i6) {
          return (_1844_i6).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(133))).length))).length)));
        };
        let _nw236 = Array((new BigNumber(17)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw236.length); _i0_37++) {
          _nw236[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1843_v122 = _nw236;
        _1843_v122 = _1843_v122;
        let _1845_v123;
        _1845_v123 = _dafny.MultiSet.fromElements(false);
        (_this).f3 = (_1845_v123).IsProperSubsetOf(_1845_v123);
        let _1846_v124;
        let _nw237 = new _module.C2();
        _nw237.__ctor((_this).f4, !(true), (_this).f1, (_this).f2);
        _1846_v124 = _nw237;
        let _1847_v125;
        _1847_v125 = _dafny.Seq.of(_1846_v124, _1846_v124, _1846_v124);
        let _1848_v126;
        _1848_v126 = _dafny.Map.Empty.slice().updateUnsafe(_1847_v125,p1);
        let _rhs291 = _1848_v126;
        let _rhs292 = _this.f3;
        let _lhs183 = _this;
        _1848_v126 = _rhs291;
        _lhs183.f3 = _rhs292;
        let _1849_v127;
        let _nw238 = Array((new BigNumber(6)).toNumber()).fill([]);
        _1849_v127 = _nw238;
        let _1850_v128;
        _1850_v128 = _module.D0.create_DC0((_1846_v124).f13);
        let _1851_v129;
        let _nw239 = Array((new BigNumber(5)).toNumber());
        _nw239[(_dafny.ZERO).toNumber()] = _1850_v128;
        _nw239[(_dafny.ONE).toNumber()] = _module.__default.fm50(_this.f12, p2, _this.f3, globalState);
        _nw239[(new BigNumber(2)).toNumber()] = _module.D0.create_DC0((_1846_v124).f13);
        _nw239[(new BigNumber(3)).toNumber()] = _1850_v128;
        _nw239[(new BigNumber(4)).toNumber()] = _1850_v128;
        _1851_v129 = _nw239;
        let _index283 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1849_v127).length));
        (_1849_v127)[_index283] = _1851_v129;
        let _index284 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1849_v127).length));
        (_1849_v127)[_index284] = _1851_v129;
      }
      let _1852_i7;
      _1852_i7 = _dafny.ZERO;
      L11: {
        while (_this.f3) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1852_i7)) {
              break L11;
            }
            _1852_i7 = (_1852_i7).plus(_dafny.ONE);
            if (_this.f3) {
              let _1853_v130;
              let _out41;
              _out41 = (_this).m8(p0, globalState);
              _1853_v130 = _out41;
              (_this).f12 = (_this).f4;
              (_this).f12 = (_this.f12).minus(((p0) ? (new BigNumber(((_this).f2).length)) : (p2)));
              let _1854_v131;
              _1854_v131 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f4);
              _1854_v131 = (_1854_v131).update((p2).isLessThan(new BigNumber(505)), (_this.f12).minus((_this).f11));
              let _1855_v132;
              let _init38 = function (_1856_i8) {
                return _module.__default.safeDivisionInt(_1856_i8, _this.f12);
              };
              let _nw240 = Array((new BigNumber(7)).toNumber());
              for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw240.length); _i0_38++) {
                _nw240[_i0_38] = _init38(new BigNumber(_i0_38));
              }
              _1855_v132 = _nw240;
              _1855_v132 = _1855_v132;
            } else {
              let _1857_v133;
              _1857_v133 = _dafny.Seq.of(_this.f3, _this.f3);
              let _rhs293 = _module.__default.fm7(true, _this.f3, globalState);
              let _rhs294 = false;
              let _lhs184 = _this;
              _1857_v133 = _rhs293;
              _lhs184.f3 = _rhs294;
              (_this).f12 = (_dafny.ZERO).minus(p2);
              let _1858_v134;
              _1858_v134 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_dafny.Seq.Create(_module.__default.abs(new BigNumber(631)), function (_1859_i9) {
                return new _dafny.CodePoint('r'.codePointAt(0));
              }));
              let _1860_v135;
              let _nw241 = new _module.C3();
              _nw241.__ctor(p2, (_this).f5, _this.f3, (_this).f1, (((_1858_v134).contains(_this.f12)) ? ((_1858_v134).get(_this.f12)) : (_dafny.Seq.UnicodeFromString("kvxnj"))));
              _1860_v135 = _nw241;
              let _1861_v136;
              let _nw242 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
              _1861_v136 = _nw242;
              let _index285 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1861_v136).length));
              (_1861_v136)[_index285] = p2;
              let _index286 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_1861_v136).length));
              (_1861_v136)[_index286] = _module.__default.fm0(globalState);
              let _1862_v137;
              _1862_v137 = _dafny.Set.fromElements(p2);
              (_this).f3 = !(_1862_v137).contains(_this.f12);
            }
            let _1863_v138;
            _1863_v138 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
            let _rhs295 = (_1863_v138).Merge(_1863_v138);
            let _rhs296 = new BigNumber(76);
            let _lhs185 = _this;
            _1863_v138 = _rhs295;
            _lhs185.f12 = _rhs296;
            let _1864_v139;
            _1864_v139 = _module.D1.create_DC2((_this).f2);
            let _source34 = _1864_v139;
            if (_source34.is_DC3) {
              let _1865___mcc_h21 = (_source34).cf3;
              let _1866___mcc_h22 = (_source34).cf4;
              let _1867___mcc_h23 = (_source34).cf5;
              let _1868_cf5 = _1867___mcc_h23;
              let _1869_cf4 = _1866___mcc_h22;
              let _1870_cf3 = _1865___mcc_h21;
              let _1871_v140;
              _1871_v140 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_this).f4, new BigNumber((_dafny.MultiSet.fromElements((_this).f2)).cardinality())));
              _1871_v140 = _dafny.Set.fromElements((_this).f4);
              let _1872_v141;
              _1872_v141 = _dafny.Seq.of(_1868_cf5, (_this).f11, (_this).f4);
              _1872_v141 = _1872_v141;
              let _1873_v142;
              let _nw243 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _1873_v142 = _nw243;
              let _index287 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_1873_v142).length));
              (_1873_v142)[_index287] = (_this).f4;
              let _index288 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_1873_v142).length));
              (_1873_v142)[_index288] = (_dafny.ZERO).minus(_1869_cf4);
              _1873_v142 = _1873_v142;
            } else if (_source34.is_DC4) {
              let _1874___mcc_h24 = (_source34).cf6;
              let _1875___mcc_h25 = (_source34).cf7;
              let _1876_cf7 = _1875___mcc_h25;
              let _1877_cf6 = _1874___mcc_h24;
              _1877_cf6 = _this.f3;
              let _1878_v143;
              _1878_v143 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).fm21(p0, (_this).f11, globalState));
              _1878_v143 = (_1878_v143).update((_this).f5, ((_1877_cf6) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_1877_cf6)).length))) : (p2)));
              (_this).f12 = ((_this).f11).minus(p2);
              let _1879_v144;
              _1879_v144 = new _dafny.CodePoint('f'.codePointAt(0));
              let _1880_v145;
              let _nw244 = Array((new BigNumber(15)).toNumber());
              _nw244[(_dafny.ZERO).toNumber()] = (_this).f2;
              _nw244[(_dafny.ONE).toNumber()] = ((_1876_cf7) ? ((_this).f2) : ((_this).f2));
              _nw244[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_1879_v144, _1879_v144);
              _nw244[(new BigNumber(3)).toNumber()] = (_this).f2;
              _nw244[(new BigNumber(4)).toNumber()] = _dafny.Seq.update((_this).f2, _module.__default.safeIndex((_this).f11, new BigNumber(((_this).f2).length)), _module.__default.fm24(p0, p0, _this.f12, globalState));
              _nw244[(new BigNumber(5)).toNumber()] = (_this).f2;
              _nw244[(new BigNumber(6)).toNumber()] = ((_module.__default.fm3((_this).f11, globalState)) ? ((_this).f2) : ((_this).f2));
              _nw244[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat((_this).f2, (_this).f2);
              _nw244[(new BigNumber(8)).toNumber()] = (_this).f2;
              _nw244[(new BigNumber(9)).toNumber()] = (_this).f2;
              _nw244[(new BigNumber(10)).toNumber()] = (_this).f2;
              _nw244[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat((_this).f2, (_this).f2);
              _nw244[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1879_v144), (_this).f2);
              _nw244[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), ((_1881_v144) => function (_1882_i10) {
                return _1881_v144;
              })(_1879_v144));
              _nw244[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat((_this).f2, _dafny.Seq.of(_1879_v144)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Concat((_this).f2, _dafny.Seq.of(_1879_v144))).length)), _1879_v144);
              _1880_v145 = _nw244;
              let _index289 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1880_v145).length));
              (_1880_v145)[_index289] = (_this).f2;
              let _index290 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1880_v145).length));
              (_1880_v145)[_index290] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(501)), function (_1883_i11) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              });
            } else {
              let _1884___mcc_h26 = (_source34).cf2;
              let _1885_cf2 = _1884___mcc_h26;
              let _1886_v146;
              _1886_v146 = _dafny.Set.fromElements((_this).f2);
              (_this).f12 = (_this).fm21((_1886_v146).contains((_this).f2), _this.f12, globalState);
              let _1887_v148;
              _1887_v148 = _dafny.Seq.of(_1885_cf2);
              let _1888_v149;
              _1888_v149 = new _dafny.CodePoint('p'.codePointAt(0));
              let _1889_v150;
              let _nw245 = new _module.C3();
              _nw245.__ctor(_this.f12, (_this).f5, _this.f3, function () {
                let _coll56 = new _dafny.Map();
                for (const _compr_56 of (_1887_v148).Elements) {
                  let _1890_v147 = _compr_56;
                  if (_dafny.Seq.contains(_1887_v148, _1890_v147)) {
                    _coll56.push([_1890_v147,false]);
                  }
                }
                return _coll56;
              }(), _dafny.Seq.update(_1885_cf2, _module.__default.safeIndex((_this).f11, new BigNumber((_1885_cf2).length)), _1888_v149));
              _1889_v150 = _nw245;
              let _1891_v151;
              _1891_v151 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
              let _1892_v152;
              _1892_v152 = _dafny.Map.Empty.slice().updateUnsafe(_1888_v149,_1891_v151);
              let _1893_v153;
              _1893_v153 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_1891_v151);
              let _1894_v154;
              _1894_v154 = _dafny.Set.fromElements(p1);
              let _1895_v155;
              _1895_v155 = _dafny.Seq.of(_1891_v151);
              let _1896_v156;
              let _nw246 = Array((new BigNumber(22)).toNumber());
              _nw246[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_module.D6.create_DC18((_this).f11, p2, (_1889_v150).fm26(new BigNumber(-455), p0, p0, globalState), _1888_v149, _this.f12)).dtor_cf31,_this.f3);
              _nw246[(_dafny.ONE).toNumber()] = (_1891_v151).Merge(_1891_v151);
              _nw246[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,true)).Merge(_1891_v151);
              _nw246[(new BigNumber(3)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(4)).toNumber()] = _module.__default.fm9(globalState);
              _nw246[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
              _nw246[(new BigNumber(6)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(7)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).fm11(p0, globalState),true)).Merge(_1891_v151);
              _nw246[(new BigNumber(8)).toNumber()] = ((!(_this.f3)) ? (_1891_v151) : ((((_1892_v152).contains(_1888_v149)) ? ((_1892_v152).get(_1888_v149)) : (_1891_v151))));
              _nw246[(new BigNumber(9)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(10)).toNumber()] = (((_1893_v153).contains(new BigNumber(((_this).fm14(_this.f3, false, _1888_v149, new BigNumber((_1894_v154).length), globalState)).length))) ? ((_1893_v153).get(new BigNumber(((_this).fm14(_this.f3, false, _1888_v149, new BigNumber((_1894_v154).length), globalState)).length))) : ((_1895_v155)[_module.__default.safeIndex(p2, new BigNumber((_1895_v155).length))]));
              _nw246[(new BigNumber(11)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(12)).toNumber()] = (_1891_v151).Merge(_1891_v151);
              _nw246[(new BigNumber(13)).toNumber()] = (_1891_v151).Merge(_1891_v151);
              _nw246[(new BigNumber(14)).toNumber()] = (_1895_v155)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f11), new BigNumber((_1895_v155).length))];
              _nw246[(new BigNumber(15)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(16)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(17)).toNumber()] = (_1895_v155)[_module.__default.safeIndex((_this).f11, new BigNumber((_1895_v155).length))];
              _nw246[(new BigNumber(18)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(19)).toNumber()] = _1891_v151;
              _nw246[(new BigNumber(20)).toNumber()] = (_1891_v151).Merge(_1891_v151);
              _nw246[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
              _1896_v156 = _nw246;
              _1896_v156 = _1896_v156;
              let _1897_v157;
              _1897_v157 = _dafny.Set.fromElements(_this.f12);
              let _1898_v158;
              _1898_v158 = _1897_v157;
              let _1899_v159;
              _1899_v159 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_1897_v157).IsDisjointFrom((_1898_v158)));
              let _1900_v160;
              _1900_v160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_1885_cf2);
              let _1901_v161;
              _1901_v161 = _dafny.Seq.of(_this.f12);
              let _1902_v162;
              _1902_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1888_v149,(_this).f11);
              _1899_v159 = (_1899_v159).update(new BigNumber((_1900_v160).length), ((_dafny.ZERO).minus((((_1902_v162).contains(_1888_v149)) ? ((_1902_v162).get(_1888_v149)) : ((_dafny.ZERO).minus((_this).f4))))).isLessThanOrEqualTo(new BigNumber((_1901_v161).length)));
            }
            (_this).f3 = p0;
          }
        }
      }
      let _1903_v163;
      _1903_v163 = _dafny.Seq.of((_this).f11);
      let _1904_v164;
      _1904_v164 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_this).fm21(p0, (_this).f11, globalState));
      let _1905_v165;
      _1905_v165 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), ((_1906_p0) => function (_1907_i13) {
        return new BigNumber((_dafny.Set.fromElements(_1906_p0, _1906_p0)).length);
      })(p0)));
      let _1908_v166;
      _1908_v166 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1905_v165)[_module.__default.safeIndex((_this).f4, new BigNumber((_1905_v165).length))]);
      let _1909_v167;
      let _nw247 = Array((new BigNumber(26)).toNumber());
      _nw247[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_this.f12, (_this).f4);
      _nw247[(_dafny.ONE).toNumber()] = p1;
      _nw247[(new BigNumber(2)).toNumber()] = p1;
      _nw247[(new BigNumber(3)).toNumber()] = (p1).Difference(_module.__default.fm8(p2, (_this).f2, globalState));
      _nw247[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_1903_v163);
      _nw247[(new BigNumber(5)).toNumber()] = p1;
      _nw247[(new BigNumber(6)).toNumber()] = p1;
      _nw247[(new BigNumber(7)).toNumber()] = p1;
      _nw247[(new BigNumber(8)).toNumber()] = p1;
      _nw247[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements((_this).f11, (_this).fm12(_this.f12, p2, (_this).f11, !(true), globalState), (_this).f11, (_this).f4, (_this).f11);
      _nw247[(new BigNumber(10)).toNumber()] = (p1).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_dafny.ZERO).minus((_this).f4))).length), (_this).fm12(_this.f12, p2, (_this).f11, _this.f3, globalState), new BigNumber(588)));
      _nw247[(new BigNumber(11)).toNumber()] = (p1).Union(p1);
      _nw247[(new BigNumber(12)).toNumber()] = (p1).update(_this.f12, _module.__default.abs((_this).f11));
      _nw247[(new BigNumber(13)).toNumber()] = (_dafny.MultiSet.FromArray(_1903_v163)).Intersect(p1);
      _nw247[(new BigNumber(14)).toNumber()] = p1;
      _nw247[(new BigNumber(15)).toNumber()] = p1;
      _nw247[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_1904_v164).length))).Difference(p1);
      _nw247[(new BigNumber(17)).toNumber()] = (p1).Intersect(p1);
      _nw247[(new BigNumber(18)).toNumber()] = p1;
      _nw247[(new BigNumber(19)).toNumber()] = ((_dafny.MultiSet.FromArray((((_1908_v166).contains(_this.f12)) ? ((_1908_v166).get(_this.f12)) : (_1903_v163)))).update((_this).f4, _module.__default.abs(new BigNumber(333)))).Difference(p1);
      _nw247[(new BigNumber(20)).toNumber()] = (p1).Union((p1).update((_dafny.ZERO).minus((_this).f11), _module.__default.abs(_this.f12)));
      _nw247[(new BigNumber(21)).toNumber()] = (p1).Difference(p1);
      _nw247[(new BigNumber(22)).toNumber()] = p1;
      _nw247[(new BigNumber(23)).toNumber()] = p1;
      _nw247[(new BigNumber(24)).toNumber()] = p1;
      _nw247[(new BigNumber(25)).toNumber()] = p1;
      _1909_v167 = _nw247;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1909_v167).length))) {
        let _1910_i12 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1910_i12)) && ((_1910_i12).isLessThan(new BigNumber((_1909_v167).length))))) {
          (_1909_v167)[(_1910_i12)] = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(p2, (_this).f4), (_this).f4, (_this).f11, (_this).f11);
        }
      }
      (_this).f3 = ((_this).f11).isLessThan(new BigNumber(((_this).f2).length));
      return;
    }
    m8(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _1911_v0;
      _1911_v0 = _module.D0.create_DC0((_this).f4);
      let _1912_v1;
      _1912_v1 = _dafny.Map.Empty.slice().updateUnsafe((_1911_v0).dtor_cf0,new BigNumber((_dafny.Seq.UnicodeFromString("xx")).length));
      let _1913_v2;
      _1913_v2 = _dafny.MultiSet.fromElements(_1912_v1, _1912_v1, _1912_v1);
      _1913_v2 = ((_1913_v2).Difference(_1913_v2)).Union((_1913_v2).update(function () {
        let _coll57 = new _dafny.Map();
        for (const _compr_57 of _dafny.IntegerRange(new BigNumber(478), new BigNumber(344))) {
          let _1914_v3 = _compr_57;
          if (((new BigNumber(478)).isLessThanOrEqualTo(_1914_v3)) && ((_1914_v3).isLessThan(new BigNumber(344)))) {
            _coll57.push([_module.__default.safeDivisionInt(_1914_v3, new BigNumber(116)),(_this).f4]);
          }
        }
        return _coll57;
      }(), _module.__default.abs(new BigNumber(601))));
      let _1915_v4;
      _1915_v4 = _dafny.Seq.of((_this).f5, (_this).f5);
      let _1916_v5;
      let _init39 = function (_1917_i1) {
        return (_1917_i1).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f11)).cardinality()));
      };
      let _nw248 = Array((new BigNumber(23)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw248.length); _i0_39++) {
        _nw248[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1916_v5 = _nw248;
      let _1918_v6;
      _1918_v6 = _dafny.Set.fromElements(_1916_v5);
      let _1919_v7;
      let _nw249 = Array((new BigNumber(22)).toNumber());
      _nw249[(_dafny.ZERO).toNumber()] = (_this).fm12(_this.f12, new BigNumber(-859), (_this).f4, p0, globalState);
      _nw249[(_dafny.ONE).toNumber()] = new BigNumber(163);
      _nw249[(new BigNumber(2)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(3)).toNumber()] = new BigNumber(727);
      _nw249[(new BigNumber(4)).toNumber()] = _module.__default.fm0(globalState);
      _nw249[(new BigNumber(5)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(6)).toNumber()] = new BigNumber(674);
      _nw249[(new BigNumber(7)).toNumber()] = (_this).f4;
      _nw249[(new BigNumber(8)).toNumber()] = (_this).f4;
      _nw249[(new BigNumber(9)).toNumber()] = (_this).f4;
      _nw249[(new BigNumber(10)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1915_v4).length));
      _nw249[(new BigNumber(12)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_1920_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length);
      _nw249[(new BigNumber(14)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(15)).toNumber()] = new BigNumber((_module.__default.fm8(new BigNumber((_1918_v6).length), (_this).f2, globalState)).cardinality());
      _nw249[(new BigNumber(16)).toNumber()] = new BigNumber(-538);
      _nw249[(new BigNumber(17)).toNumber()] = (_this).f4;
      _nw249[(new BigNumber(18)).toNumber()] = new BigNumber(((_this).f2).length);
      _nw249[(new BigNumber(19)).toNumber()] = _this.f12;
      _nw249[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f12))).cardinality());
      _nw249[(new BigNumber(21)).toNumber()] = (_this).f11;
      _1919_v7 = _nw249;
      let _1921_v8;
      let _nw250 = new _module.C0();
      _nw250.__ctor(_1919_v7);
      _1921_v8 = _nw250;
      let _1922_v9;
      _1922_v9 = _dafny.MultiSet.fromElements((_this).f11, _this.f12);
      let _1923_v10;
      _1923_v10 = _dafny.Seq.of(_this.f3);
      let _1924_v11;
      _1924_v11 = _module.D14.create_DC39(new BigNumber(326), ((((_1922_v9).contains((_this).f11)) ? ((_1922_v9).get((_this).f11)) : (new BigNumber((_1923_v10).length)))).plus((_this).f11), (_this).f4);
      _1924_v11 = _1924_v11;
      let _1925_v12;
      _1925_v12 = _dafny.Seq.of(new BigNumber((_1923_v10).length));
      if (!_dafny.areEqual(_1925_v12, _1925_v12)) {
        let _1926_v13;
        _1926_v13 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1927_v14;
        let _nw251 = new _module.C4();
        _nw251.__ctor(_this.f3, _dafny.Map.Empty.slice().updateUnsafe((_this).fm14(p0, _this.f3, _1926_v13, new BigNumber(((_this).f2).length), globalState),_this.f3), (_this).f2);
        _1927_v14 = _nw251;
        let _1928_v15;
        _1928_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f5);
        _1928_v15 = (_1928_v15).update(true, (_this).f5);
        let _1929_v16;
        _1929_v16 = _dafny.Seq.UnicodeFromString("ugutx");
        _1929_v16 = (_this).fm14((p0) && (_this.f3), p0, _1926_v13, (_this).f4, globalState);
        let _1930_v17;
        let _nw252 = new _module.C0();
        _nw252.__ctor((_1921_v8).f14);
        _1930_v17 = _nw252;
        if (p0) {
          let _1931_v18;
          _1931_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
          let _1932_v20;
          _1932_v20 = _dafny.Seq.of(_1929_v16, _module.__default.fm2(p0, globalState), (_this).f2, _1929_v16, _1929_v16);
          let _1933_v21;
          let _nw253 = new _module.C4();
          _nw253.__ctor(!_dafny.Seq.contains(_1925_v12, new BigNumber((_dafny.Seq.of(new BigNumber((_1931_v18).length), new BigNumber((_1925_v12).length))).length)), ((_this).f1).Merge(function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_1932_v20).Elements) {
              let _1934_v19 = _compr_58;
              if (_dafny.Seq.contains(_1932_v20, _1934_v19)) {
                _coll58.push([_1934_v19,true]);
              }
            }
            return _coll58;
          }()), _dafny.Seq.UnicodeFromString("ykpsh"));
          _1933_v21 = _nw253;
          let _1935_v22;
          _1935_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_dafny.Seq.Concat((_this).f2, (_this).f2));
          _1929_v16 = (((_1935_v22).contains((_this).f5)) ? ((_1935_v22).get((_this).f5)) : (_1929_v16));
          let _1936_v23;
          _1936_v23 = _dafny.MultiSet.fromElements(_1929_v16);
          let _1937_v24;
          _1937_v24 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),(_this).f4);
          let _1938_v25;
          _1938_v25 = _dafny.Seq.of(_1937_v24, _1937_v24);
          let _1939_v26;
          _1939_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1936_v23).Intersect(_1936_v23),_1938_v25);
          _1939_v26 = (_1939_v26).update((_1936_v23).update(_1929_v16, _module.__default.abs(_this.f12)), _1938_v25);
          (_this).f12 = ((_1933_v21).fm12((_module.__default.fm10(_this.f3, _this.f3, _this.f3, globalState)).dtor_cf12, _this.f12, new BigNumber((_1922_v9).cardinality()), _this.f3, globalState)).minus(new BigNumber((_1929_v16).length));
          (_this).f12 = _module.__default.safeModuloInt(new BigNumber(((_this).f2).length), (_module.D3.create_DC9(new BigNumber(643), _module.__default.fm3((_this).fm21(p0, (_this).f11, globalState), globalState))).dtor_cf15);
        } else {
          r0 = (!(!(_this.f3))) && (!(p0) || (p0));
          let _1940_v27;
          _1940_v27 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_this.f3)).cardinality()));
          let _1941_v28;
          _1941_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1940_v27,((p0) ? (new BigNumber(912)) : ((_this).f11)));
          let _1942_v29;
          _1942_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1930_v17,_1940_v27);
          let _1943_v30;
          _1943_v30 = _dafny.Map.Empty.slice().updateUnsafe((((_1912_v1).contains((_this).f11)) ? ((_1912_v1).get((_this).f11)) : (_this.f12)),!(_this.f3));
          let _1944_v31;
          _1944_v31 = _dafny.MultiSet.fromElements(_1943_v30);
          let _1945_v32;
          _1945_v32 = _module.D3.create_DC9((_this).f11, p0);
          _1941_v28 = (_1941_v28).update((_1940_v27).Union((((_1942_v29).contains(_1930_v17)) ? ((_1942_v29).get(_1930_v17)) : (_dafny.Set.fromElements(new BigNumber(834), new BigNumber((_1944_v31).cardinality()))))), (_1945_v32).dtor_cf15);
          let _1946_v33;
          _1946_v33 = _module.D6.create_DC16(_1943_v30);
          let _1947_v34;
          _1947_v34 = _module.D6.create_DC20(_1946_v33);
          let _1948_v35;
          _1948_v35 = _module.D6.create_DC20((_1947_v34).dtor_cf35);
          let _1949_v36;
          let _nw254 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1949_v36 = _nw254;
          let _1950_v37;
          let _nw255 = Array((new BigNumber(18)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = _1916_v5;
          _nw255[(_dafny.ONE).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(2)).toNumber()] = (_1930_v17).f14;
          _nw255[(new BigNumber(3)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(4)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(5)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(6)).toNumber()] = (_1930_v17).f14;
          _nw255[(new BigNumber(7)).toNumber()] = (_1930_v17).f14;
          _nw255[(new BigNumber(8)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(9)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(10)).toNumber()] = (_1921_v8).f14;
          _nw255[(new BigNumber(11)).toNumber()] = _1916_v5;
          _nw255[(new BigNumber(12)).toNumber()] = _1916_v5;
          _nw255[(new BigNumber(13)).toNumber()] = _1916_v5;
          _nw255[(new BigNumber(14)).toNumber()] = _1916_v5;
          _nw255[(new BigNumber(15)).toNumber()] = (_1930_v17).f14;
          _nw255[(new BigNumber(16)).toNumber()] = _1919_v7;
          _nw255[(new BigNumber(17)).toNumber()] = (_1930_v17).f14;
          _1950_v37 = _nw255;
          let _index291 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1949_v36).length));
          (_1949_v36)[_index291] = _1950_v37;
          let _index292 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1949_v36).length));
          let _rhs297 = _1948_v35;
          let _rhs298 = _1950_v37;
          let _rhs299 = (_this).f2;
          let _lhs186 = _1949_v36;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_1949_v36).length));
          _1947_v34 = _rhs297;
          _lhs186[_lhs187] = _rhs298;
          _1929_v16 = _rhs299;
          let _1951_v38;
          let _nw256 = new _module.C4();
          _nw256.__ctor(p0, (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("q"),!(_this.f3))).Merge((_this).f1), _1929_v16);
          _1951_v38 = _nw256;
          let _1952_v39;
          _1952_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1929_v16,(_this).f11);
          let _rhs300 = ((_this.f12).isLessThanOrEqualTo(_this.f12)) || (p0);
          let _rhs301 = _1951_v38;
          let _rhs302 = (((_1951_v38.f3) && (!(_1951_v38.f3))) ? (((_this).f11).isLessThanOrEqualTo(_this.f12)) : ((_module.__default.fm5(_module.__default.fm51((_this).f11, _this.f12, globalState), globalState)).IsProperSubsetOf(_module.__default.fm5(_1952_v39, globalState))));
          let _lhs188 = _this;
          let _lhs189 = _this;
          _lhs188.f3 = _rhs300;
          _1951_v38 = _rhs301;
          _lhs189.f3 = _rhs302;
          _1919_v7 = ((p0) ? ((_1921_v8).f14) : ((_1930_v17).f14));
        }
      } else {
        let _1953_v40;
        let _nw257 = new _module.C5();
        _nw257.__ctor(p0, (_this).f1, (_this).f2);
        _1953_v40 = _nw257;
        let _1954_v41;
        _1954_v41 = _dafny.Seq.of(_1953_v40, _1953_v40, _1953_v40);
        let _1955_v42;
        let _init40 = function (_1956_i2) {
          return !(_this.f3);
        };
        let _nw258 = Array((new BigNumber(19)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw258.length); _i0_40++) {
          _nw258[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1955_v42 = _nw258;
        let _rhs303 = _1954_v41;
        let _rhs304 = (_this).f5;
        _1954_v41 = _rhs303;
        _1955_v42 = _rhs304;
        if (!(p0) || (p0)) {
          let _1957_v43;
          _1957_v43 = _dafny.MultiSet.fromElements(!(_this.f3));
          let _1958_v44;
          let _nw259 = new _module.C2();
          _nw259.__ctor((((_1957_v43).contains(_this.f3)) ? ((_1957_v43).get(_this.f3)) : (_this.f12)), p0, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wppe"),_this.f3), _module.__default.fm2(_this.f3, globalState));
          _1958_v44 = _nw259;
          _1957_v43 = _1957_v43;
          let _1959_v45;
          _1959_v45 = _dafny.Set.fromElements(_this.f12);
          _1959_v45 = ((_this.f3) ? (_dafny.Set.fromElements((_this).f4, _module.__default.fm0(globalState), (_1958_v44).f13, new BigNumber(318), (_this).f4)) : (_1959_v45));
          r0 = _this.f3;
          let _1960_v46;
          _1960_v46 = _dafny.Seq.UnicodeFromString("u");
          _1960_v46 = _dafny.Seq.Concat((_1953_v40).f2, _dafny.Seq.Concat((_this).f2, (_1953_v40).f2));
        } else {
          let _1961_v47;
          let _init41 = ((_1962_p0) => function (_1963_i3) {
            return ((_1962_p0) ? (_module.D1.create_DC3(_1962_p0, _this.f12, new BigNumber(757))) : (_module.D1.create_DC3(_this.f3, (_this).f11, (_this).f4)));
          })(p0);
          let _nw260 = Array((new BigNumber(4)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw260.length); _i0_41++) {
            _nw260[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1961_v47 = _nw260;
          let _1964_v48;
          _1964_v48 = _module.D1.create_DC3(_this.f3, (_this).f11, (_this).f4);
          let _index293 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_1961_v47).length));
          (_1961_v47)[_index293] = _1964_v48;
          let _1965_v49;
          let _nw261 = new _module.C3();
          _nw261.__ctor((_this).f11, (_this).f5, _this.f3, (_1953_v40).f1, _dafny.Seq.UnicodeFromString("ctsrc"));
          _1965_v49 = _nw261;
          let _1966_v50;
          _1966_v50 = _dafny.MultiSet.fromElements(_1965_v49, _1965_v49);
          let _1967_v51;
          _1967_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1966_v50).cardinality()));
          let _1968_v52;
          _1968_v52 = _dafny.Set.fromElements((_this).f4, (_dafny.ZERO).minus((_this).f11), new BigNumber((_1967_v51).length), (_this).f11, (_this).f11);
          let _index294 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_1961_v47).length));
          let _rhs305 = !((_this).fm13(globalState));
          let _rhs306 = !(p0);
          let _rhs307 = ((_this).f11).plus(new BigNumber((_1968_v52).length));
          let _rhs308 = _1964_v48;
          let _rhs309 = ((_1968_v52).Union(_dafny.Set.fromElements((_this).f11, _this.f12))).IsProperSubsetOf(_dafny.Set.fromElements((_this).f4, (_this).f4, new BigNumber((_1968_v52).length)));
          let _lhs190 = _this;
          let _lhs191 = _this;
          let _lhs192 = _1961_v47;
          let _lhs193 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_1961_v47).length));
          let _lhs194 = _this;
          r0 = _rhs305;
          _lhs190.f3 = _rhs306;
          _lhs191.f12 = _rhs307;
          _lhs192[_lhs193] = _rhs308;
          _lhs194.f3 = _rhs309;
          let _1969_v53;
          let _nw262 = new _module.C2();
          _nw262.__ctor((_this).f4, p0, (_1953_v40).f1, (_1953_v40).f2);
          _1969_v53 = _nw262;
          let _1970_v54;
          _1970_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1969_v53,_1955_v42);
          let _1971_v55;
          _1971_v55 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f11).plus(new BigNumber((_1970_v54).length)),(_1953_v40).f2);
          _1971_v55 = _1971_v55;
          (_this).f3 = _this.f3;
          r0 = _this.f3;
          (_this).f12 = ((_1969_v53).f13).plus((_1969_v53).f13);
        }
        let _1972_v56;
        _1972_v56 = _dafny.Set.fromElements(new BigNumber(909));
        let _1973_v57;
        _1973_v57 = _dafny.Set.fromElements(_1972_v56, _1972_v56);
        _1973_v57 = _1973_v57;
        let _1974_v58;
        _1974_v58 = _dafny.Seq.of(_1919_v7);
        _1916_v5 = (_1974_v58)[_module.__default.safeIndex(_this.f12, new BigNumber((_1974_v58).length))];
        let _1975_v59;
        let _nw263 = new _module.C5();
        _nw263.__ctor(_this.f3, (_this).f1, _dafny.Seq.UnicodeFromString("ncifpw"));
        _1975_v59 = _nw263;
      }
      let _1976_v60;
      let _nw264 = Array((new BigNumber(2)).toNumber());
      _nw264[(_dafny.ZERO).toNumber()] = p0;
      _nw264[(_dafny.ONE).toNumber()] = p0;
      _1976_v60 = _nw264;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1976_v60).length))) {
        let _1977_i4 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1977_i4)) && ((_1977_i4).isLessThan(new BigNumber((_1976_v60).length))))) {
          (_1976_v60)[(_1977_i4)] = !((_this.f3) || ((_module.D10.create_DC29(p0, p0)).dtor_cf48));
        }
      }
      let _1978_v61;
      _1978_v61 = _dafny.Set.fromElements(_this.f3);
      let _1979_v62;
      _1979_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_1978_v61);
      let _1980_v63;
      _1980_v63 = _module.D3.create_DC8((((_1979_v62).contains((_this).f4)) ? ((_1979_v62).get((_this).f4)) : (_1978_v61)));
      let _pat_let_tv46 = p0;
      let _pat_let_tv47 = p0;
      let _pat_let_tv48 = p0;
      let _rhs310 = function (_source35) {
        if (_source35.is_DC9) {
          let _1981___mcc_h0 = (_source35).cf15;
          let _1982___mcc_h1 = (_source35).cf16;
          let _1983_cf16 = _1982___mcc_h1;
          let _1984_cf15 = _1981___mcc_h0;
          return ((_this).f11).isLessThanOrEqualTo((_this).f4);
        } else if (_source35.is_DC10) {
          let _1985___mcc_h2 = (_source35).cf17;
          let _1986___mcc_h3 = (_source35).cf18;
          let _1987_cf18 = _1986___mcc_h3;
          let _1988_cf17 = _1985___mcc_h2;
          return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv46,(_this).f11)).length)).isLessThan(_1988_cf17);
        } else if (_source35.is_DC8) {
          let _1989___mcc_h4 = (_source35).cf14;
          let _1990_cf14 = _1989___mcc_h4;
          return _pat_let_tv47;
        } else {
          let _1991___mcc_h5 = (_source35).cf19;
          let _1992_cf19 = _1991___mcc_h5;
          return _pat_let_tv48;
        }
      }(_1980_v63);
      let _rhs311 = _module.__default.safeModuloInt((_this).f11, new BigNumber(((_this).f2).length));
      let _lhs195 = _this;
      r0 = _rhs310;
      _lhs195.f12 = _rhs311;
      r0 = p0;
      return r0;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new BigNumber(((_this).f2).length), (_this).f4, (_this).f4)).Union(_dafny.MultiSet.fromElements((_this).f4))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f4, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f4, (_this).f4),_this.f3)).length))))));
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pvsoenjvi"), (_this).f2);
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm11(p0, globalState) {
      let _this = this;
      return (function () {
        let _coll59 = new _dafny.Set();
        for (const _compr_59 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_1993_i0) {
          return _module.D1.create_DC4(_this.f3, _this.f3);
        })).Elements) {
          let _1994_v0 = _compr_59;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_1993_i0) {
            return _module.D1.create_DC4(_this.f3, _this.f3);
          }), _1994_v0)) {
            _coll59.add(_1994_v0);
          }
        }
        return _coll59;
      }()).IsSubsetOf(_dafny.Set.fromElements(_module.D1.create_DC4(_this.f3, _this.f3), _module.D1.create_DC4(true, true)));
    };
    fm18(p0, p1, globalState) {
      let _this = this;
      if (_this.f3) {
        if (_this.f3) {
          return _module.D4.create_DC12(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f4, (_this).f4)));
        } else {
          return _module.D4.create_DC12(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-10)), function (_1995_i0) {
  return (_dafny.ZERO).minus((_this).f4);
})).length)));
        }
      } else {
        return _module.D4.create_DC12(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true, _this.f3)).cardinality()),false))).length), (_this).f4));
      }
    };
    fm19(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      r2 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, p1))).plus(_module.__default.fm0(globalState)));
      let _1996_v0;
      _1996_v0 = new _dafny.CodePoint('g'.codePointAt(0));
      let _1997_v1;
      _1997_v1 = _dafny.Set.fromElements(p1);
      let _1998_v2;
      _1998_v2 = _dafny.Seq.of(new BigNumber((_1997_v1).length));
      let _1999_v3;
      _1999_v3 = _dafny.Seq.of(_this.f3, _this.f3);
      let _2000_v4;
      _2000_v4 = _module.D3.create_DC9(new BigNumber(333), _this.f3);
      let _2001_v5;
      _2001_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v0,!(_this.f3));
      let _2002_v6;
      _2002_v6 = _dafny.MultiSet.fromElements(_this.f3, _this.f3, _this.f3);
      let _2003_v7;
      let _nw265 = Array((new BigNumber(26)).toNumber());
      _nw265[(_dafny.ZERO).toNumber()] = (_this).f4;
      _nw265[(_dafny.ONE).toNumber()] = p1;
      _nw265[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f2).length));
      _nw265[(new BigNumber(3)).toNumber()] = (_module.D2.create_DC6(_1996_v0, _this.f3, false, p0)).dtor_cf12;
      _nw265[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus(p0)).plus(p1);
      _nw265[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-932), new BigNumber((_1998_v2).length));
      _nw265[(new BigNumber(6)).toNumber()] = ((_this).f4).multipliedBy(p1);
      _nw265[(new BigNumber(7)).toNumber()] = new BigNumber(((_this).f2).length);
      _nw265[(new BigNumber(8)).toNumber()] = (_this).f4;
      _nw265[(new BigNumber(9)).toNumber()] = new BigNumber(923);
      _nw265[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber((_1998_v2).length));
      _nw265[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1997_v1).Difference(_1997_v1)).length));
      _nw265[(new BigNumber(12)).toNumber()] = (_this).f4;
      _nw265[(new BigNumber(13)).toNumber()] = p0;
      _nw265[(new BigNumber(14)).toNumber()] = (((_1999_v3)[_module.__default.safeIndex(new BigNumber((_1999_v3).length), new BigNumber((_1999_v3).length))]) ? (p0) : ((_this).f4));
      _nw265[(new BigNumber(15)).toNumber()] = new BigNumber((_1998_v2).length);
      _nw265[(new BigNumber(16)).toNumber()] = (_1998_v2)[_module.__default.safeIndex((_this).fm19((_this).f4, (_dafny.ZERO).minus(p0), _module.D3.create_DC11(_2000_v4), _1998_v2, globalState), new BigNumber((_1998_v2).length))];
      _nw265[(new BigNumber(17)).toNumber()] = p1;
      _nw265[(new BigNumber(18)).toNumber()] = ((_this.f3) ? (p1) : ((_dafny.ZERO).minus((_this).f4)));
      _nw265[(new BigNumber(19)).toNumber()] = new BigNumber((_2001_v5).length);
      _nw265[(new BigNumber(20)).toNumber()] = (_this).f4;
      _nw265[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(88), (_dafny.ZERO).minus(new BigNumber((_2002_v6).cardinality())));
      _nw265[(new BigNumber(22)).toNumber()] = (_this).f4;
      _nw265[(new BigNumber(23)).toNumber()] = p1;
      _nw265[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
      _nw265[(new BigNumber(25)).toNumber()] = (p1).minus(p0);
      _2003_v7 = _nw265;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2003_v7).length))) {
        let _2004_i0 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2004_i0)) && ((_2004_i0).isLessThan(new BigNumber((_2003_v7).length))))) {
          (_2003_v7)[(_2004_i0)] = (_2004_i0).plus(new BigNumber(((_dafny.MultiSet.fromElements((_1998_v2)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_1998_v2).length))])).Union(_dafny.MultiSet.fromElements(p1, p1, new BigNumber(((_this).f2).length), p1, p0))).cardinality()));
        }
      }
      let _hi7 = p1;
      for (let _2005_i1 = _module.__default.safeModuloInt(p1, p0); _2005_i1.isLessThan(_hi7); _2005_i1 = _2005_i1.plus(_dafny.ONE)) {
        (_this).f3 = (_module.__default.safeDivisionInt(new BigNumber(227), _2005_i1)).isLessThan((_this).f4);
        let _2006_v8;
        _2006_v8 = _dafny.Seq.of(_module.__default.fm7(_this.f3, _this.f3, globalState), _1999_v3);
        _1999_v3 = _dafny.Seq.Concat((_2006_v8)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex((_this).f4, new BigNumber(((_this).f2).length)), _1996_v0)).length), new BigNumber((_2006_v8).length))], _1999_v3);
        let _2007_v9;
        let _nw266 = Array((new BigNumber(17)).toNumber()).fill(false);
        _2007_v9 = _nw266;
        let _rhs312 = _dafny.Seq.of(((true) ? (true) : (_this.f3)));
        let _rhs313 = ((_this.f3) ? (_2007_v9) : ((_this).f5));
        let _rhs314 = new BigNumber(((_this).f2).length);
        let _rhs315 = _this.f3;
        let _lhs196 = _this;
        _1999_v3 = _rhs312;
        _2007_v9 = _rhs313;
        r2 = _rhs314;
        _lhs196.f3 = _rhs315;
        if ((_this.f3) || (_this.f3)) {
          (_this).f3 = _this.f3;
          let _index295 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2003_v7).length));
          (_2003_v7)[_index295] = ((_this.f3) ? (p0) : ((_this).f4));
          let _index296 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2003_v7).length));
          let _rhs316 = (((new BigNumber((_1998_v2).length)).isLessThan((_this).f4)) ? (new BigNumber(26)) : (p1));
          let _rhs317 = (p1).plus(new BigNumber(((_dafny.Set.fromElements(_this.f3)).Difference(_dafny.Set.fromElements(_this.f3))).length));
          let _lhs197 = _2003_v7;
          let _lhs198 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_2003_v7).length));
          _lhs197[_lhs198] = _rhs316;
          r1 = _rhs317;
          let _2008_v10;
          _2008_v10 = _dafny.Seq.UnicodeFromString("phdcd");
          _2008_v10 = _2008_v10;
          let _2009_v11;
          _2009_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_1998_v2);
          r2 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2009_v11).length), p1))).multipliedBy((p1).multipliedBy((_2003_v7)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_2003_v7).length))]));
          let _2010_v12;
          _2010_v12 = _dafny.Seq.of(_2000_v4);
          let _2011_v13;
          _2011_v13 = _module.D3.create_DC11((_2010_v12)[_module.__default.safeIndex(p1, new BigNumber((_2010_v12).length))]);
          (_this).f3 = (_this.f3) && ((_module.__default.fm0(globalState)).isLessThan((_this).fm19(p0, (_2003_v7)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_2003_v7).length))], _2011_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_2012_v7) => function (_2013_i2) {
            return (_2012_v7)[_module.__default.safeIndex(new BigNumber(25), new BigNumber((_2012_v7).length))];
          })(_2003_v7)), globalState)));
        } else {
          let _index297 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length));
          (_2007_v9)[_index297] = _this.f3;
          let _2014_v14;
          _2014_v14 = _dafny.Seq.of(_1998_v2, _1998_v2, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f2).length)), _dafny.Seq.of(_2005_i1)), _module.__default.safeIndex((_dafny.ZERO).minus(_2005_i1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_this).f2).length)), _dafny.Seq.of(_2005_i1))).length)), (_this).fm19(_2005_i1, _2005_i1, _module.D3.create_DC11(_module.D3.create_DC11(_2000_v4)), _1998_v2, globalState)), _dafny.Seq.of(p0, p0, p1), _1998_v2);
          let _index298 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length));
          let _rhs318 = _this.f3;
          let _rhs319 = _2003_v7;
          let _rhs320 = _this.f3;
          let _rhs321 = (_this).fm13(globalState);
          let _rhs322 = new BigNumber(((_2014_v14)[_module.__default.safeIndex(_2005_i1, new BigNumber((_2014_v14).length))]).length);
          let _lhs199 = _2007_v9;
          let _lhs200 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length));
          let _lhs201 = _this;
          let _lhs202 = _this;
          _lhs199[_lhs200] = _rhs318;
          _2003_v7 = _rhs319;
          _lhs201.f3 = _rhs320;
          _lhs202.f3 = _rhs321;
          r2 = _rhs322;
          r1 = p1;
          let _index299 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_2003_v7).length));
          (_2003_v7)[_index299] = _2005_i1;
          let _index300 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_2003_v7).length));
          let _index301 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length));
          let _rhs323 = ((_this).f4).minus(p0);
          let _rhs324 = _module.__default.safeModuloInt((_2005_i1).multipliedBy(p1), p0);
          let _rhs325 = _dafny.Seq.IsProperPrefixOf(_1999_v3, _1999_v3);
          let _lhs203 = _2003_v7;
          let _lhs204 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_2003_v7).length));
          let _lhs205 = _2007_v9;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length));
          r1 = _rhs323;
          _lhs203[_lhs204] = _rhs324;
          _lhs205[_lhs206] = _rhs325;
          let _2015_v15;
          let _nw267 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2015_v15 = _nw267;
          _2015_v15 = (((_2007_v9)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length))]) ? (_2015_v15) : (_2015_v15));
          let _2016_v16;
          _2016_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2007_v9)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length))]);
          let _2017_v17;
          _2017_v17 = _dafny.Map.Empty.slice().updateUnsafe(((((_2016_v16).contains((_1998_v2)[_module.__default.safeIndex(p1, new BigNumber((_1998_v2).length))])) ? ((_2016_v16).get((_1998_v2)[_module.__default.safeIndex(p1, new BigNumber((_1998_v2).length))])) : ((_2007_v9)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_2007_v9).length))]))) || (_this.f3),(_this).f2);
          r2 = new BigNumber((_2017_v17).length);
        }
      }
      let _2018_v18;
      _2018_v18 = _module.D5.create_DC15((_this).f4, !(_this.f3), (_this).f5);
      let _2019_v19;
      _2019_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f5);
      let _2020_v20;
      _2020_v20 = _module.D3.create_DC9(new BigNumber(((_this).f2).length), (_this).fm13(globalState));
      let _2021_v21;
      _2021_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_2019_v19).length), (_dafny.ZERO).minus(p1), (_2020_v20).dtor_cf15, new BigNumber((_2002_v6).cardinality()))).cardinality()),(_dafny.ZERO).minus(p0));
      let _2022_v22;
      _2022_v22 = _dafny.Map.Empty.slice().updateUnsafe((_2018_v18).dtor_cf25,_2021_v21);
      _2022_v22 = (_2022_v22).update(_this.f3, _2021_v21);
      let _hi8 = _dafny.ZERO;
      for (let _2023_i3 = new BigNumber(374); _2023_i3.isLessThan(_hi8); _2023_i3 = _2023_i3.plus(_dafny.ONE)) {
        let _2024_v23;
        _2024_v23 = _module.D3.create_DC10((_dafny.ZERO).minus(p0), (_dafny.ZERO).minus(p1));
        let _pat_let_tv49 = p1;
        let _source36 = function (_pat_let64_0) {
          return function (_2025_dt__update__tmp_h0) {
            return function (_pat_let65_0) {
              return function (_2026_dt__update_hcf18_h0) {
                return _module.D3.create_DC10((_2025_dt__update__tmp_h0).dtor_cf17, _2026_dt__update_hcf18_h0);
              }(_pat_let65_0);
            }(_pat_let_tv49);
          }(_pat_let64_0);
        }(_2024_v23);
        if (_source36.is_DC9) {
          let _2027___mcc_h0 = (_source36).cf15;
          let _2028___mcc_h1 = (_source36).cf16;
          let _2029_cf16 = _2028___mcc_h1;
          let _2030_cf15 = _2027___mcc_h0;
          _1999_v3 = _1999_v3;
          let _2031_v24;
          _2031_v24 = _dafny.Seq.UnicodeFromString("bcmpvem");
          _2031_v24 = (_this).f2;
          r2 = (p0).minus((_this).f4);
          r2 = (_this).f4;
        } else if (_source36.is_DC10) {
          let _2032___mcc_h2 = (_source36).cf17;
          let _2033___mcc_h3 = (_source36).cf18;
          let _2034_cf18 = _2033___mcc_h3;
          let _2035_cf17 = _2032___mcc_h2;
          _2003_v7 = _2003_v7;
          let _2036_v25;
          _2036_v25 = _dafny.Map.Empty.slice().updateUnsafe(_2003_v7,_this.f3);
          _2036_v25 = (_2036_v25).update(_2003_v7, _this.f3);
          (_this).f3 = _this.f3;
          _2019_v19 = (_2019_v19).update(_2034_cf18, (_this).f5);
        } else if (_source36.is_DC8) {
          let _2037___mcc_h4 = (_source36).cf14;
          let _2038_cf14 = _2037___mcc_h4;
          _1996_v0 = _1996_v0;
          _2003_v7 = _2003_v7;
          let _2039_v26;
          _2039_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1996_v0,_2023_i3);
          let _rhs326 = (_dafny.ZERO).minus(((_2023_i3).minus((_this).f4)).plus(_module.__default.safeDivisionInt(_2023_i3, p0)));
          let _rhs327 = ((_2002_v6).IsDisjointFrom((_2002_v6).update(false, _module.__default.abs(new BigNumber((_dafny.Set.fromElements(_2023_i3, (((_2039_v26).contains(new _dafny.CodePoint('i'.codePointAt(0)))) ? ((_2039_v26).get(new _dafny.CodePoint('i'.codePointAt(0)))) : (_2023_i3)), _2023_i3)).length))))) && (_this.f3);
          let _lhs207 = _this;
          r2 = _rhs326;
          _lhs207.f3 = _rhs327;
          let _2040_v27;
          _2040_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,!(true));
          let _2041_v28;
          let _2042_v29;
          let _2043_v30;
          let _2044_v31;
          let _out42;
          let _out43;
          let _out44;
          let _out45;
          let _outcollector10 = _module.__default.m0((_2040_v27).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_1999_v3)[_module.__default.safeIndex(new BigNumber((_2002_v6).cardinality()), new BigNumber((_1999_v3).length))])), globalState);
          _out42 = _outcollector10[0];
          _out43 = _outcollector10[1];
          _out44 = _outcollector10[2];
          _out45 = _outcollector10[3];
          _2041_v28 = _out42;
          _2042_v29 = _out43;
          _2043_v30 = _out44;
          _2044_v31 = _out45;
        } else {
          let _2045___mcc_h5 = (_source36).cf19;
          let _2046_cf19 = _2045___mcc_h5;
          let _2047_v32;
          let _nw268 = Array((new BigNumber(12)).toNumber());
          _nw268[(_dafny.ZERO).toNumber()] = ((_this.f3) ? (_1998_v2) : (_1998_v2));
          _nw268[(_dafny.ONE).toNumber()] = _dafny.Seq.of(p0, (_dafny.ZERO).minus(p1), (_dafny.ZERO).minus(new BigNumber((_2021_v21).length)));
          _nw268[(new BigNumber(2)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(3)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(4)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(5)).toNumber()] = _module.__default.fm20(_this.f3, _this.f3, _this.f3, _1999_v3, globalState);
          _nw268[(new BigNumber(6)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(7)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1998_v2, _1998_v2);
          _nw268[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1998_v2, _1998_v2);
          _nw268[(new BigNumber(10)).toNumber()] = _1998_v2;
          _nw268[(new BigNumber(11)).toNumber()] = _1998_v2;
          _2047_v32 = _nw268;
          let _nw269 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _2047_v32 = _nw269;
          let _2048_v33;
          let _init42 = ((_2049_p0) => function (_2050_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2049_p0));
          })(p0);
          let _nw270 = Array((new BigNumber(26)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw270.length); _i0_42++) {
            _nw270[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _2048_v33 = _nw270;
          let _2051_v34;
          _2051_v34 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p1);
          let _index302 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2048_v33).length));
          (_2048_v33)[_index302] = _2051_v34;
          let _2052_v35;
          _2052_v35 = _dafny.MultiSet.fromElements(_1996_v0);
          let _index303 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2048_v33).length));
          (_2048_v33)[_index303] = _module.__default.fm4((((_2052_v35).contains(_1996_v0)) ? ((_2052_v35).get(_1996_v0)) : ((_this).f4)), !(_2002_v6).contains(_this.f3), globalState);
          (_this).f3 = _this.f3;
          let _index304 = _module.__default.safeIndex(new BigNumber(168), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index304] = _this.f3;
          let _2053_v36;
          _2053_v36 = _dafny.Seq.of(_2003_v7);
          let _index305 = _module.__default.safeIndex(new BigNumber(168), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index305] = _dafny.Seq.IsProperPrefixOf(_2053_v36, _2053_v36);
        }
        if (_dafny.Seq.contains(((_this.f3) ? (_1998_v2) : (_1998_v2)), (_this).f4)) {
          (_this).f3 = !(true);
          let _index306 = _module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index306] = _this.f3;
          let _index307 = _module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index307] = (_module.__default.safeModuloInt(p1, (_this).f4)).isEqualTo(((true) ? (_2023_i3) : (_2023_i3)));
          let _2054_v37;
          _2054_v37 = _dafny.Seq.UnicodeFromString("cst");
          _2054_v37 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("djhll"), _2054_v37), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("djhll"), _2054_v37)).length)), _1996_v0);
          let _2055_v38;
          _2055_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2023_i3,(_2021_v21).update(p0, p0));
          let _2056_v39;
          let _nw271 = new _module.C5();
          _nw271.__ctor(!(((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))]), (_this).f1, _dafny.Seq.update(_dafny.Seq.Concat(_2054_v37, _2054_v37), _module.__default.safeIndex(new BigNumber((_2055_v38).length), new BigNumber((_dafny.Seq.Concat(_2054_v37, _2054_v37)).length)), _1996_v0));
          _2056_v39 = _nw271;
          let _2057_v40;
          _2057_v40 = _dafny.Set.fromElements(_module.__default.fm22(((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))], new BigNumber(725), globalState));
          let _2058_v41;
          let _nw272 = Array((new BigNumber(10)).toNumber());
          _nw272[(_dafny.ZERO).toNumber()] = !(_module.__default.fm3(p0, globalState));
          _nw272[(_dafny.ONE).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))];
          _nw272[(new BigNumber(2)).toNumber()] = _this.f3;
          _nw272[(new BigNumber(3)).toNumber()] = _this.f3;
          _nw272[(new BigNumber(4)).toNumber()] = (_2056_v39).fm11((_2056_v39).f15, globalState);
          _nw272[(new BigNumber(5)).toNumber()] = !(_2057_v40).contains(_1996_v0);
          _nw272[(new BigNumber(6)).toNumber()] = (_this.f3) === (((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))]);
          _nw272[(new BigNumber(7)).toNumber()] = (_2056_v39).f15;
          _nw272[(new BigNumber(8)).toNumber()] = !(_this.f3);
          _nw272[(new BigNumber(9)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))];
          _2058_v41 = _nw272;
          let _2059_v42;
          _2059_v42 = _module.D6.create_DC17(_this.f3);
          let _2060_v43;
          let _nw273 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2060_v43 = _nw273;
          let _2061_v44;
          _2061_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2060_v43,true);
          let _nw274 = Array((new BigNumber(13)).toNumber());
          _nw274[(_dafny.ZERO).toNumber()] = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-276)), function (_2062_i6) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("abxt")).length);
          }), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), ((_2063_v0) => function (_2064_i5) {
            return _2063_v0;
          })(_1996_v0))).length));
          _nw274[(_dafny.ONE).toNumber()] = (_2056_v39).f15;
          _nw274[(new BigNumber(2)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))];
          _nw274[(new BigNumber(3)).toNumber()] = _this.f3;
          _nw274[(new BigNumber(4)).toNumber()] = (false) || (!(_this.f3));
          _nw274[(new BigNumber(5)).toNumber()] = (_2059_v42).dtor_cf28;
          _nw274[(new BigNumber(6)).toNumber()] = !((_2056_v39).f15);
          _nw274[(new BigNumber(7)).toNumber()] = true;
          _nw274[(new BigNumber(8)).toNumber()] = ((_this.f3) ? (_this.f3) : (_this.f3));
          _nw274[(new BigNumber(9)).toNumber()] = (_2056_v39).f15;
          _nw274[(new BigNumber(10)).toNumber()] = (_2056_v39).f15;
          _nw274[(new BigNumber(11)).toNumber()] = (_2056_v39).f15;
          _nw274[(new BigNumber(12)).toNumber()] = (((_2061_v44).contains(_2060_v43)) ? ((_2061_v44).get(_2060_v43)) : ((_2056_v39).fm11(((_this).f5)[_module.__default.safeIndex(new BigNumber(321), new BigNumber(((_this).f5).length))], globalState)));
          _2058_v41 = _nw274;
        } else {
          (_this).f3 = _this.f3;
          _1996_v0 = _1996_v0;
          _2003_v7 = _2003_v7;
          let _2065_v45;
          _2065_v45 = _dafny.MultiSet.fromElements(p0);
          r2 = _module.__default.safeDivisionInt(new BigNumber((_2065_v45).cardinality()), ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2021_v21,_this.f3)).length))).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), ((_2066_v0) => function (_2067_i7) {
            return _2066_v0;
          })(_1996_v0))).length)));
          let _2068_v46;
          _2068_v46 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f3));
          let _2069_v47;
          _2069_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f2);
          let _2070_v48;
          let _nw275 = Array((new BigNumber(3)).toNumber());
          _nw275[(_dafny.ZERO).toNumber()] = _2068_v46;
          _nw275[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_2071_i8) {
            return _dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f3));
          }), _module.__default.safeIndex(new BigNumber((_2069_v47).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_2072_i8) {
            return _dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f3));
          })).length)), _2002_v6);
          _nw275[(new BigNumber(2)).toNumber()] = _2068_v46;
          _2070_v48 = _nw275;
          let _index308 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_2070_v48).length));
          (_2070_v48)[_index308] = _2068_v46;
          let _index309 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_2070_v48).length));
          let _rhs328 = _dafny.Seq.Concat(_module.__default.fm52((_this).fm13(globalState), new BigNumber(775), _this.f3, globalState), _dafny.Seq.Concat(_2068_v46, _2068_v46));
          let _rhs329 = _2065_v45;
          let _lhs208 = _2070_v48;
          let _lhs209 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_2070_v48).length));
          _lhs208[_lhs209] = _rhs328;
          _2065_v45 = _rhs329;
        }
        let _2073_v49;
        _2073_v49 = _dafny.Set.fromElements(_this.f3);
        let _2074_v50;
        _2074_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2073_v49);
        let _2075_v51;
        let _nw276 = new _module.C6();
        _nw276.__ctor((new BigNumber(821)).plus(p0), new BigNumber((_2074_v50).length), _module.__default.safeDivisionInt(p0, p0), (_this).f5, _this.f3, (_this).f1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("decv"), _dafny.Seq.update(_module.__default.fm2((_this).fm11(_this.f3, globalState), globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm2((_this).fm11(_this.f3, globalState), globalState)).length)), _1996_v0)));
        _2075_v51 = _nw276;
        let _2076_v52;
        _2076_v52 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
        (_this).f3 = (_2023_i3).isLessThan((_2075_v51.f12).multipliedBy(new BigNumber(((_2076_v52).update(_this.f3, _this.f3)).length)));
      }
      r2 = (_this).fm12((_this).f4, (_this).f4, new BigNumber(479), _this.f3, globalState);
      r0 = _module.__default.fm9(globalState);
      r1 = (p1).minus((_this).f4);
      r2 = p0;
      return [r0, r1, r2];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
      this._f10 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f10, f3, f1, f2, f4, f5) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_module.D3.create_DC9((_this).f4, _this.f3)).dtor_cf16) {
        return (_this).f4;
      } else {
        return (_this).f4;
      }
    };
    fm11(p0, globalState) {
      let _this = this;
      let _source37 = ((_this.f3) ? (_module.D0.create_DC1((_this).f4)) : (_module.D0.create_DC1(new BigNumber(((_this).f10).length))));
      if (_source37.is_DC1) {
        let _2077___mcc_h0 = (_source37).cf1;
        let _2078_cf1 = _2077___mcc_h0;
        return _this.f3;
      } else {
        let _2079___mcc_h1 = (_source37).cf0;
        let _2080_cf0 = _2079___mcc_h1;
        return !((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f3))).IsSubsetOf(_dafny.MultiSet.fromElements(_this.f3)));
      }
    };
    fm13(globalState) {
      let _this = this;
      return !(((_module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),_this.f3))).dtor_cf27).Merge(function () {
        let _coll60 = new _dafny.Map();
        for (const _compr_60 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), function (_2081_i0) {
          return (_dafny.ZERO).minus((_this).f4);
        })).Elements) {
          let _2082_v0 = _compr_60;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), function (_2081_i0) {
            return (_dafny.ZERO).minus((_this).f4);
          }), _2082_v0)) {
            _coll60.push([(_2082_v0).plus((_this).f4),_this.f3]);
          }
        }
        return _coll60;
      }())).contains((_this).f4);
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((_this).f2, _dafny.Seq.UnicodeFromString("s"));
    };
    fm16(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.update((_this).f2, _module.__default.safeIndex(((_this.f3) ? ((_this).f4) : (new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_this.f3))).length))), new BigNumber(((_this).f2).length)), new _dafny.CodePoint('h'.codePointAt(0)));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      if (!(function (_source38) {
        if (_source38.is_DC9) {
          let _2083___mcc_h0 = (_source38).cf15;
          let _2084___mcc_h1 = (_source38).cf16;
          let _2085_cf16 = _2084___mcc_h1;
          let _2086_cf15 = _2083___mcc_h0;
          return _2085_cf16;
        } else if (_source38.is_DC10) {
          let _2087___mcc_h2 = (_source38).cf17;
          let _2088___mcc_h3 = (_source38).cf18;
          let _2089_cf18 = _2088___mcc_h3;
          let _2090_cf17 = _2087___mcc_h2;
          return _this.f3;
        } else if (_source38.is_DC8) {
          let _2091___mcc_h4 = (_source38).cf14;
          let _2092_cf14 = _2091___mcc_h4;
          return (true) === (false);
        } else {
          let _2093___mcc_h5 = (_source38).cf19;
          let _2094_cf19 = _2093___mcc_h5;
          return _this.f3;
        }
      }(_module.__default.fm17(p1, globalState)))) {
        let _2095_v0;
        _2095_v0 = _dafny.Seq.of(p0, p1);
        r2 = (_2095_v0)[_module.__default.safeIndex((_this).f4, new BigNumber((_2095_v0).length))];
        (_this).f3 = _this.f3;
        let _2096_v1;
        _2096_v1 = _module.D5.create_DC15(p1, _this.f3, (_this).f5);
        _2096_v1 = _module.D5.create_DC15(p0, !((p0).isLessThanOrEqualTo(p0)), (_this).f5);
        let _2097_v2;
        _2097_v2 = _dafny.Seq.of(_this.f3, _this.f3);
        let _2098_v3;
        _2098_v3 = new _dafny.CodePoint('y'.codePointAt(0));
        let _2099_v4;
        _2099_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2098_v3,p1);
        let _2100_v5;
        _2100_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
        let _2101_v6;
        _2101_v6 = _dafny.MultiSet.fromElements(p1, new BigNumber(549));
        let _2102_v7;
        let _nw277 = Array((new BigNumber(21)).toNumber());
        _nw277[(_dafny.ZERO).toNumber()] = (_2097_v2)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_2097_v2).length))];
        _nw277[(_dafny.ONE).toNumber()] = _this.f3;
        _nw277[(new BigNumber(2)).toNumber()] = ((_this.f3) ? (!(_this.f3)) : (_this.f3));
        _nw277[(new BigNumber(3)).toNumber()] = _this.f3;
        _nw277[(new BigNumber(4)).toNumber()] = !(_this.f3);
        _nw277[(new BigNumber(5)).toNumber()] = _this.f3;
        _nw277[(new BigNumber(6)).toNumber()] = (new BigNumber((_2097_v2).length)).isLessThanOrEqualTo((_dafny.ZERO).minus((((_2099_v4).contains(_2098_v3)) ? ((_2099_v4).get(_2098_v3)) : (p0))));
        _nw277[(new BigNumber(7)).toNumber()] = _this.f3;
        _nw277[(new BigNumber(8)).toNumber()] = _this.f3;
        _nw277[(new BigNumber(9)).toNumber()] = !(_this.f3) || (_this.f3);
        _nw277[(new BigNumber(10)).toNumber()] = _this.f3;
        _nw277[(new BigNumber(11)).toNumber()] = !_dafny.Seq.contains(_2095_v0, p0);
        _nw277[(new BigNumber(12)).toNumber()] = false;
        _nw277[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.update((_this).f2, _module.__default.safeIndex(p0, new BigNumber(((_this).f2).length)), _2098_v3), (_this).f2);
        _nw277[(new BigNumber(14)).toNumber()] = !(_this.f3) || (_this.f3);
        _nw277[(new BigNumber(15)).toNumber()] = ((_this.f3) ? (_this.f3) : (_this.f3));
        _nw277[(new BigNumber(16)).toNumber()] = !(_this.f3) || (_this.f3);
        _nw277[(new BigNumber(17)).toNumber()] = ((_this).fm12(new BigNumber(871), p0, p0, _this.f3, globalState)).isLessThanOrEqualTo(new BigNumber((_2100_v5).length));
        _nw277[(new BigNumber(18)).toNumber()] = false;
        _nw277[(new BigNumber(19)).toNumber()] = true;
        _nw277[(new BigNumber(20)).toNumber()] = !(_2101_v6).contains(new BigNumber(659));
        _2102_v7 = _nw277;
        _2102_v7 = ((_this.f3) ? (((_this.f3) ? (_2102_v7) : ((_this).f5))) : ((_this).f5));
        if ((p1).isLessThan((_this).f4)) {
          let _2103_v8;
          _2103_v8 = _dafny.Seq.of(((_this).f2)[_module.__default.safeIndex(p1, new BigNumber(((_this).f2).length))]);
          _2103_v8 = _dafny.Seq.Concat(_2103_v8, (_this).f2);
          let _2104_v9;
          let _nw278 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
          _2104_v9 = _nw278;
          let _2105_v10;
          _2105_v10 = _dafny.Set.fromElements(p1, new BigNumber((_dafny.Seq.UnicodeFromString("fsar")).length));
          let _index310 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_2104_v9).length));
          (_2104_v9)[_index310] = _2105_v10;
          let _index311 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_2104_v9).length));
          (_2104_v9)[_index311] = _dafny.Set.fromElements(p0, (_this).fm12((_this).fm12(new BigNumber((_2095_v0).length), (_this).f4, (_this).f4, _this.f3, globalState), p1, (_this).f4, _this.f3, globalState));
          let _index312 = _module.__default.safeIndex(new BigNumber(839), new BigNumber((_2104_v9).length));
          (_2104_v9)[_index312] = (_2104_v9)[_module.__default.safeIndex(new BigNumber(839), new BigNumber((_2104_v9).length))];
          let _2106_v11;
          let _nw279 = Array((new BigNumber(8)).toNumber());
          _nw279[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw279[(_dafny.ONE).toNumber()] = _2098_v3;
          _nw279[(new BigNumber(2)).toNumber()] = _2098_v3;
          _nw279[(new BigNumber(3)).toNumber()] = _2098_v3;
          _nw279[(new BigNumber(4)).toNumber()] = _2098_v3;
          _nw279[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
          _nw279[(new BigNumber(6)).toNumber()] = _2098_v3;
          _nw279[(new BigNumber(7)).toNumber()] = _2098_v3;
          _2106_v11 = _nw279;
          let _2107_v12;
          _2107_v12 = _dafny.Seq.of(_2106_v11, _2106_v11);
          let _2108_v13;
          let _nw280 = Array((new BigNumber(3)).toNumber());
          _nw280[(_dafny.ZERO).toNumber()] = (_2107_v12)[_module.__default.safeIndex(p0, new BigNumber((_2107_v12).length))];
          _nw280[(_dafny.ONE).toNumber()] = _2106_v11;
          _nw280[(new BigNumber(2)).toNumber()] = _2106_v11;
          _2108_v13 = _nw280;
          let _2109_v14;
          _2109_v14 = _dafny.Set.fromElements(!((_this).fm13(globalState)));
          let _2110_v15;
          _2110_v15 = _2108_v13;
          let _rhs330 = new BigNumber((_2109_v14).length);
          let _rhs331 = !(_this.f3);
          let _rhs332 = (_2110_v15);
          let _rhs333 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(380)), function (_2111_i0) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }), _dafny.Seq.update((_this).f2, _module.__default.safeIndex((_this).f4, new BigNumber(((_this).f2).length)), _2098_v3)), _dafny.Seq.Concat((_this).f2, _2103_v8));
          let _lhs210 = _this;
          r1 = _rhs330;
          _lhs210.f3 = _rhs331;
          _2108_v13 = _rhs332;
          _2103_v8 = _rhs333;
          let _2112_v16;
          _2112_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).fm14(_this.f3, _module.__default.fm3(new BigNumber((_2095_v0).length), globalState), _2098_v3, p1, globalState));
          let _rhs334 = (_this).f4;
          let _rhs335 = ((!(_this.f3) || (_this.f3)) ? (_2103_v8) : (_dafny.Seq.Concat((((_2112_v16).contains(_2102_v7)) ? ((_2112_v16).get(_2102_v7)) : (_2103_v8)), (_this).f2)));
          let _rhs336 = p1;
          let _rhs337 = _this.f3;
          let _lhs211 = _this;
          r2 = _rhs334;
          _2103_v8 = _rhs335;
          r2 = _rhs336;
          _lhs211.f3 = _rhs337;
        } else {
          let _2113_v17;
          _2113_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2102_v7,p0);
          _2113_v17 = (_2113_v17).update(_2102_v7, (_2095_v0)[_module.__default.safeIndex((_this).f4, new BigNumber((_2095_v0).length))]);
          r1 = (_this).f4;
          (_this).f3 = _this.f3;
          let _2114_v18;
          _2114_v18 = _dafny.Seq.UnicodeFromString("xildsfr");
          _2114_v18 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jogglsdwh"), (_this).f2);
          let _2115_v19;
          let _nw281 = new _module.C2();
          _nw281.__ctor(((_this.f3) ? (p0) : (p1)), _this.f3, (_this).f1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), ((_2116_v3) => function (_2117_i1) {
            return _2116_v3;
          })(_2098_v3)));
          _2115_v19 = _nw281;
        }
      } else {
        let _index313 = _module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index313] = !((p1).isLessThan(p0));
        let _2118_v20;
        _2118_v20 = _dafny.Set.fromElements(new BigNumber(((_this).f2).length), p1, p1);
        let _2119_v21;
        _2119_v21 = _dafny.MultiSet.fromElements(_2118_v20);
        let _index314 = _module.__default.safeIndex(new BigNumber(75), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index314] = (_dafny.MultiSet.fromElements(_2118_v20)).IsProperSubsetOf(_2119_v21);
        let _2120_v22;
        _2120_v22 = _dafny.Set.fromElements(!(_this.f3));
        let _index315 = _module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length));
        let _index316 = _module.__default.safeIndex(new BigNumber(75), new BigNumber(((_this).f5).length));
        let _rhs338 = (p1).plus(new BigNumber(62));
        let _rhs339 = p1;
        let _rhs340 = !(_this.f3);
        let _rhs341 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_this).f2, _dafny.Seq.UnicodeFromString("ggplncq")), _module.__default.fm2(true, globalState));
        let _rhs342 = _module.__default.fm3(new BigNumber((_2120_v22).length), globalState);
        let _lhs212 = (_this).f5;
        let _lhs213 = _module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length));
        let _lhs214 = (_this).f5;
        let _lhs215 = _module.__default.safeIndex(new BigNumber(75), new BigNumber(((_this).f5).length));
        let _lhs216 = _this;
        r2 = _rhs338;
        r2 = _rhs339;
        _lhs212[_lhs213] = _rhs340;
        _lhs214[_lhs215] = _rhs341;
        _lhs216.f3 = _rhs342;
        let _2121_v23;
        _2121_v23 = _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))]);
        let _2122_v24;
        let _init43 = function (_2123_i2) {
          return false;
        };
        let _nw282 = Array((new BigNumber(25)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw282.length); _i0_43++) {
          _nw282[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _2122_v24 = _nw282;
        let _2124_v25;
        _2124_v25 = _dafny.MultiSet.fromElements(p0, new BigNumber(((_this).f2).length), p1, p0, p0);
        let _2125_v26;
        _2125_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2122_v24,_2124_v25);
        let _2126_v27;
        let _nw283 = Array((new BigNumber(18)).toNumber());
        _nw283[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2121_v23, _2121_v23);
        _nw283[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_2121_v23, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((((_2125_v26).contains(_2122_v24)) ? ((_2125_v26).get(_2122_v24)) : (_2124_v25))).cardinality())), new BigNumber((_2121_v23).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))]), _2121_v23);
        _nw283[(new BigNumber(2)).toNumber()] = _2121_v23;
        _nw283[(new BigNumber(3)).toNumber()] = _2121_v23;
        _nw283[(new BigNumber(4)).toNumber()] = _module.__default.fm7(_this.f3, (_2121_v23)[_module.__default.safeIndex(p0, new BigNumber((_2121_v23).length))], globalState);
        _nw283[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))]);
        _nw283[(new BigNumber(6)).toNumber()] = _module.__default.fm7(((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))], _this.f3, globalState);
        _nw283[(new BigNumber(7)).toNumber()] = ((_this).f10)[_module.__default.safeIndex(new BigNumber(396), new BigNumber(((_this).f10).length))];
        _nw283[(new BigNumber(8)).toNumber()] = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))]) ? (_2121_v23) : (_2121_v23));
        _nw283[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2121_v23, _2121_v23);
        _nw283[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_2121_v23, _module.__default.safeIndex((_this).f4, new BigNumber((_2121_v23).length)), ((_this).f5)[_module.__default.safeIndex(new BigNumber(407), new BigNumber(((_this).f5).length))]);
        _nw283[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_2121_v23, _2121_v23), _module.__default.safeIndex(new BigNumber(-153), new BigNumber((_dafny.Seq.Concat(_2121_v23, _2121_v23)).length)), _this.f3);
        _nw283[(new BigNumber(12)).toNumber()] = _2121_v23;
        _nw283[(new BigNumber(13)).toNumber()] = _2121_v23;
        _nw283[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_2121_v23, _2121_v23);
        _nw283[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_this.f3, false);
        _nw283[(new BigNumber(16)).toNumber()] = _2121_v23;
        _nw283[(new BigNumber(17)).toNumber()] = _2121_v23;
        _2126_v27 = _nw283;
        _2126_v27 = _2126_v27;
        let _2127_v28;
        let _nw284 = Array((new BigNumber(3)).toNumber());
        _nw284[(_dafny.ZERO).toNumber()] = _2122_v24;
        _nw284[(_dafny.ONE).toNumber()] = _2122_v24;
        _nw284[(new BigNumber(2)).toNumber()] = _2122_v24;
        _2127_v28 = _nw284;
        let _index317 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_2127_v28).length));
        (_2127_v28)[_index317] = _2122_v24;
        let _index318 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_2127_v28).length));
        (_2127_v28)[_index318] = _2122_v24;
        let _2128_v29;
        let _nw285 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _2128_v29 = _nw285;
        let _2129_v30;
        let _nw286 = new _module.C0();
        _nw286.__ctor(_2128_v29);
        _2129_v30 = _nw286;
        r1 = new BigNumber(((_this).f2).length);
      }
      let _index319 = _module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index319] = _this.f3;
      let _2130_v31;
      _2130_v31 = _module.D11.create_DC31((_this).f4);
      let _pat_let_tv50 = p1;
      let _pat_let_tv51 = p1;
      let _index320 = _module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length));
      let _rhs343 = p1;
      let _rhs344 = _this.f3;
      let _rhs345 = function (_source39) {
        if (_source39.is_DC31) {
          let _2131___mcc_h6 = (_source39).cf50;
          let _2132_cf50 = _2131___mcc_h6;
          return !(_this.f3) || (_this.f3);
        } else {
          let _2133___mcc_h7 = (_source39).cf49;
          let _2134_cf49 = _2133___mcc_h7;
          return (_pat_let_tv50).isLessThan(_pat_let_tv51);
        }
      }(_2130_v31);
      let _rhs346 = p1;
      let _rhs347 = false;
      let _lhs217 = _this;
      let _lhs218 = (_this).f5;
      let _lhs219 = _module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length));
      let _lhs220 = _this;
      r2 = _rhs343;
      _lhs217.f3 = _rhs344;
      _lhs218[_lhs219] = _rhs345;
      r1 = _rhs346;
      _lhs220.f3 = _rhs347;
      let _2135_v32;
      let _nw287 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2135_v32 = _nw287;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2135_v32).length))) {
        let _2136_i3 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2136_i3)) && ((_2136_i3).isLessThan(new BigNumber((_2135_v32).length))))) {
          (_2135_v32)[(_2136_i3)] = _dafny.Seq.UnicodeFromString("ihcix");
        }
      }
      r1 = (_dafny.ZERO).minus(new BigNumber(((_this).f2).length));
      let _2137_v33;
      let _nw288 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _2137_v33 = _nw288;
      let _index321 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_2137_v33).length));
      (_2137_v33)[_index321] = p0;
      let _index322 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_2137_v33).length));
      (_2137_v33)[_index322] = (_dafny.ZERO).minus((p0).minus(_module.__default.safeDivisionInt((_this).f4, p1)));
      let _index323 = _module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index323] = _this.f3;
      let _2138_v34;
      _2138_v34 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length))],_this.f3);
      r0 = (_2138_v34).Merge(_2138_v34);
      let _2139_v35;
      _2139_v35 = _dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(108), new BigNumber(((_this).f5).length))]);
      r1 = new BigNumber((((_2139_v35).Intersect(_2139_v35)).Intersect((_2139_v35).Intersect(_2139_v35))).length);
      r2 = p1;
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let _2140_v0;
      _2140_v0 = new _dafny.CodePoint('m'.codePointAt(0));
      let _2141_v1;
      _2141_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f4);
      let _2142_v2;
      _2142_v2 = _module.D6.create_DC18(p2, p2, _this.f3, _2140_v0, (((_2141_v1).contains(p2)) ? ((_2141_v1).get(p2)) : (p2)));
      let _2143_v3;
      _2143_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2142_v2,(_module.D6.create_DC18(new BigNumber(489), (_this).f4, _this.f3, _2140_v0, (_this).f4)).dtor_cf32);
      _2143_v3 = (_2143_v3).update(_2142_v2, _2140_v0);
      let _2144_v4;
      _2144_v4 = _dafny.Seq.UnicodeFromString("vtawg");
      _2144_v4 = ((_this.f3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(804)), ((_2145_v4, _2146_p2) => function (_2147_i0) {
        return (_2145_v4)[_module.__default.safeIndex(_2146_p2, new BigNumber((_2145_v4).length))];
      })(_2144_v4, p2))) : ((_this).f2));
      let _2148_v5;
      let _nw289 = Array((_dafny.ONE).toNumber());
      _nw289[(_dafny.ZERO).toNumber()] = p2;
      _2148_v5 = _nw289;
      let _2149_v6;
      _2149_v6 = _module.D5.create_DC14(_2148_v5);
      let _2150_v7;
      _2150_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v6,false);
      if ((((_2150_v7).contains(_2149_v6)) ? ((_2150_v7).get(_2149_v6)) : (_this.f3))) {
        let _2151_v9;
        let _init44 = function (_2152_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(_this.f3,function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of _dafny.IntegerRange(new BigNumber(4), new BigNumber(213))) {
              let _2153_v8 = _compr_61;
              if (((new BigNumber(4)).isLessThanOrEqualTo(_2153_v8)) && ((_2153_v8).isLessThan(new BigNumber(213)))) {
                _coll61.add((_2153_v8).plus(new BigNumber(502)));
              }
            }
            return _coll61;
          }());
        };
        let _nw290 = Array((new BigNumber(13)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw290.length); _i0_44++) {
          _nw290[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _2151_v9 = _nw290;
        let _2154_v10;
        _2154_v10 = _dafny.MultiSet.fromElements(p2);
        let _2155_v11;
        _2155_v11 = _dafny.Seq.of(_2154_v10);
        let _2156_v12;
        _2156_v12 = _dafny.Set.fromElements((_this).f4, p2, (_this).f4, (_this).f4);
        let _2157_v13;
        _2157_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(new BigNumber((_2155_v11).length), globalState),_2156_v12);
        let _index324 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_2151_v9).length));
        (_2151_v9)[_index324] = _2157_v13;
        let _2158_v14;
        _2158_v14 = _dafny.Set.fromElements(_this.f3, _this.f3);
        let _index325 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_2151_v9).length));
        let _rhs348 = _2157_v13;
        let _rhs349 = ((_2158_v14).Union(_2158_v14)).IsSubsetOf(_2158_v14);
        let _lhs221 = _2151_v9;
        let _lhs222 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_2151_v9).length));
        let _lhs223 = _this;
        _lhs221[_lhs222] = _rhs348;
        _lhs223.f3 = _rhs349;
        let _index326 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index326] = _this.f3;
        let _index327 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index327] = (_this).fm11((new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_this.f3)).update((_this).f4, _this.f3)).length)).isLessThan(new BigNumber(-305)), globalState);
        let _2159_v15;
        _2159_v15 = _module.D8.create_DC23(new BigNumber((_2141_v1).length), p2, ((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))]);
        let _index328 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index328] = (_this).fm11((_2159_v15).dtor_cf40, globalState);
        if (((_dafny.Set.fromElements(!(false), _this.f3, _this.f3)).Intersect(_2158_v14)).IsSubsetOf(_2158_v14)) {
          let _index329 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index329] = p2;
          let _index330 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index330] = (_this).f4;
          let _2160_v16;
          _2160_v16 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))],false);
          let _2161_v17;
          _2161_v17 = _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))], true, (((_2160_v16).contains(_this.f3)) ? ((_2160_v16).get(_this.f3)) : (_this.f3)), _this.f3);
          let _2162_v18;
          _2162_v18 = _dafny.MultiSet.fromElements(_2161_v17, _dafny.Seq.of(_this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))], _this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))]));
          let _index331 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index331] = new BigNumber((_2162_v18).cardinality());
          let _2163_v19;
          let _nw291 = Array((new BigNumber(12)).toNumber()).fill(false);
          _2163_v19 = _nw291;
          _2163_v19 = _2163_v19;
          let _2164_v20;
          let _nw292 = new _module.C4();
          _nw292.__ctor((_2154_v10).IsDisjointFrom(_2154_v10), _dafny.Map.Empty.slice().updateUnsafe((_this).f2,true), _dafny.Seq.Concat(_2144_v4, _2144_v4));
          _2164_v20 = _nw292;
          let _2165_v21;
          _2165_v21 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))],_2141_v1);
          let _index332 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          let _index333 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          let _rhs350 = _module.__default.safeModuloInt(new BigNumber((_2165_v21).length), new BigNumber(-125));
          let _rhs351 = _2140_v0;
          let _rhs352 = (_this).f4;
          let _lhs224 = _2148_v5;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          let _lhs226 = _2148_v5;
          let _lhs227 = _module.__default.safeIndex(new BigNumber(973), new BigNumber((_2148_v5).length));
          _lhs224[_lhs225] = _rhs350;
          _2140_v0 = _rhs351;
          _lhs226[_lhs227] = _rhs352;
        } else {
          let _2166_v22;
          _2166_v22 = new BigNumber(816);
          _2166_v22 = (_dafny.ZERO).minus(((((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))]) ? ((_dafny.ZERO).minus((_this).f4)) : (((_dafny.ZERO).minus(p2)).plus(_2166_v22))));
          let _2167_v23;
          _2167_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2148_v5,(_this).f2);
          let _2168_v24;
          _2168_v24 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2167_v23);
          _2168_v24 = (_2168_v24).update(_this.f3, _dafny.Map.Empty.slice().updateUnsafe(_2148_v5,(_this).f2));
          let _2169_v25;
          _2169_v25 = _dafny.MultiSet.fromElements(_this.f3, ((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))]);
          let _2170_v26;
          let _nw293 = new _module.C2();
          _nw293.__ctor((_this).f4, _this.f3, (_this).f1, (_this).f2);
          _2170_v26 = _nw293;
          let _2171_v27;
          _2171_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2170_v26,_this.f3);
          let _2172_v28;
          _2172_v28 = _module.D13.create_DC33(_2170_v26);
          let _2173_v29;
          _2173_v29 = _dafny.Seq.of(!(_this.f3), _this.f3, (((_2171_v27).contains((_module.D13.create_DC33((_2172_v28).dtor_cf52)).dtor_cf52)) ? ((_2171_v27).get((_module.D13.create_DC33((_2172_v28).dtor_cf52)).dtor_cf52)) : ((_this).fm13(globalState))));
          let _index334 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length));
          let _rhs353 = new BigNumber(239);
          let _rhs354 = (_dafny.MultiSet.FromArray(_2173_v29)).Union((_2169_v25).Union(_2169_v25));
          let _rhs355 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2144_v4, _2144_v4), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-897)), ((_2174_v0) => function (_2175_i2) {
            return _2174_v0;
          })(_2140_v0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_2176_v0) => function (_2177_i3) {
            return _2176_v0;
          })(_2140_v0))));
          let _rhs356 = p2;
          let _rhs357 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length))];
          let _lhs228 = (_this).f5;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_this).f5).length));
          _2166_v22 = _rhs353;
          _2169_v25 = _rhs354;
          _2144_v4 = _rhs355;
          _2166_v22 = _rhs356;
          _lhs228[_lhs229] = _rhs357;
          let _index335 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index335] = ((_2170_v26).f13).multipliedBy(p2);
          let _index336 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index336] = (_2170_v26).f13;
          let _index337 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2148_v5).length));
          (_2148_v5)[_index337] = p2;
        }
        let _2178_v30;
        _2178_v30 = new BigNumber(868);
        _2178_v30 = _2178_v30;
      } else {
        let _2179_v31;
        _2179_v31 = _dafny.Seq.of(_this.f3, (_2142_v2).dtor_cf31);
        (_this).f3 = !((_2179_v31)[_module.__default.safeIndex((_this).f4, new BigNumber((_2179_v31).length))]);
        let _2180_v32;
        _2180_v32 = _dafny.Seq.of(new BigNumber((_2141_v1).length), (_this).f4);
        let _2181_v33;
        _2181_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("wftn"));
        let _2182_v34;
        _2182_v34 = _module.D11.create_DC30(_2181_v33);
        let _2183_v35;
        _2183_v35 = _module.D13.create_DC35(new BigNumber(18), (_2180_v32)[_module.__default.safeIndex(p2, new BigNumber((_2180_v32).length))], _this.f3, new BigNumber((_dafny.Seq.update(_2144_v4, _module.__default.safeIndex(p2, new BigNumber((_2144_v4).length)), _2140_v0)).length), _2182_v34);
        let _2184_v36;
        _2184_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2183_v35,_2140_v0);
        let _rhs358 = _2140_v0;
        let _rhs359 = (_module.__default.fm53((_this).f4, globalState)).Merge(_2184_v36);
        let _rhs360 = _module.__default.fm3(p2, globalState);
        let _lhs230 = _this;
        _2140_v0 = _rhs358;
        _2184_v36 = _rhs359;
        _lhs230.f3 = _rhs360;
        (_this).f3 = _this.f3;
        let _2185_v37;
        _2185_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f3);
        let _2186_v38;
        _2186_v38 = _dafny.MultiSet.fromElements(_2144_v4, (_this).f2);
        if ((((_2185_v37).contains((new BigNumber((_2186_v38).cardinality())).minus(new BigNumber(674)))) ? ((_2185_v37).get((new BigNumber((_2186_v38).cardinality())).minus(new BigNumber(674)))) : (_this.f3))) {
          let _index338 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index338] = true;
          let _index339 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index339] = ((function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(387), new BigNumber(160))) {
              let _2187_v39 = _compr_62;
              if (((new BigNumber(387)).isLessThanOrEqualTo(_2187_v39)) && ((_2187_v39).isLessThan(new BigNumber(160)))) {
                _coll62.push([_module.__default.safeModuloInt(_2187_v39, p2),(_this).f4]);
              }
            }
            return _coll62;
          }()).Merge(_2141_v1)).equals(_2141_v1);
          let _2188_v40;
          _2188_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p2);
          let _2189_v41;
          _2189_v41 = _dafny.Set.fromElements(_module.__default.fm54(_2188_v40, _this.f3, globalState));
          let _2190_v42;
          _2190_v42 = _dafny.Map.Empty.slice().updateUnsafe(false,_2189_v41);
          let _2191_v43;
          let _init45 = function (_2192_i4) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fswednf"), _dafny.Seq.UnicodeFromString("hs"));
          };
          let _nw294 = Array((new BigNumber(26)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw294.length); _i0_45++) {
            _nw294[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _2191_v43 = _nw294;
          let _index340 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          let _index341 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          let _rhs361 = (_2190_v42).update((true) || (!(false)), (_2189_v41).Difference(_2189_v41));
          let _rhs362 = _2185_v37;
          let _rhs363 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length))];
          let _rhs364 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length))];
          let _rhs365 = _2191_v43;
          let _lhs231 = (_this).f5;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          let _lhs233 = (_this).f5;
          let _lhs234 = _module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length));
          _2190_v42 = _rhs361;
          _2185_v37 = _rhs362;
          _lhs231[_lhs232] = _rhs363;
          _lhs233[_lhs234] = _rhs364;
          _2191_v43 = _rhs365;
          let _2193_v44;
          let _init46 = function (_2194_i5) {
            return ((_this).f5)[_module.__default.safeIndex(new BigNumber(617), new BigNumber(((_this).f5).length))];
          };
          let _nw295 = Array((new BigNumber(3)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw295.length); _i0_46++) {
            _nw295[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _2193_v44 = _nw295;
          let _2195_v45;
          _2195_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2193_v44,false);
          let _2196_v46;
          let _2197_v47;
          let _2198_v48;
          let _2199_v49;
          let _out46;
          let _out47;
          let _out48;
          let _out49;
          let _outcollector11 = _module.__default.m0(_2195_v45, globalState);
          _out46 = _outcollector11[0];
          _out47 = _outcollector11[1];
          _out48 = _outcollector11[2];
          _out49 = _outcollector11[3];
          _2196_v46 = _out46;
          _2197_v47 = _out47;
          _2198_v48 = _out48;
          _2199_v49 = _out49;
          let _2200_v50;
          _2200_v50 = _module.D2.create_DC6(_2140_v0, _2197_v47, _this.f3, (_this).f4);
          let _2201_v51;
          _2201_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p2, new BigNumber((_dafny.Seq.UnicodeFromString("k")).length)),_2200_v50);
          _2201_v51 = (_2201_v51).update(_2180_v32, _2200_v50);
          let _2202_v52;
          let _2203_v53;
          let _2204_v54;
          let _2205_v55;
          let _out50;
          let _out51;
          let _out52;
          let _out53;
          let _outcollector12 = _module.__default.m0(_2195_v45, globalState);
          _out50 = _outcollector12[0];
          _out51 = _outcollector12[1];
          _out52 = _outcollector12[2];
          _out53 = _outcollector12[3];
          _2202_v52 = _out50;
          _2203_v53 = _out51;
          _2204_v54 = _out52;
          _2205_v55 = _out53;
        } else {
          let _2206_v56;
          let _init47 = function (_2207_i6) {
            return !(_this.f3) || (_this.f3);
          };
          let _nw296 = Array((new BigNumber(26)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw296.length); _i0_47++) {
            _nw296[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _2206_v56 = _nw296;
          let _2208_v57;
          _2208_v57 = _dafny.Seq.of(_2206_v56, _2206_v56, (_this).f5, (_this).f5);
          let _2209_v58;
          _2209_v58 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f4);
          let _rhs366 = _this.f3;
          let _rhs367 = (_2208_v57)[_module.__default.safeIndex(((_this).f4).plus(new BigNumber((_2209_v58).length)), new BigNumber((_2208_v57).length))];
          let _lhs235 = _this;
          _lhs235.f3 = _rhs366;
          _2206_v56 = _rhs367;
          let _2210_v59;
          _2210_v59 = new BigNumber(411);
          _2210_v59 = (_dafny.ZERO).minus(p2);
          (_this).f3 = _this.f3;
          _2209_v58 = (_2209_v58).update(false, (_dafny.ZERO).minus((_this).f4));
          let _2211_v60;
          _2211_v60 = _dafny.MultiSet.fromElements((_this).f4, new BigNumber((_2180_v32).length));
          let _2212_v61;
          _2212_v61 = _module.D2.create_DC7(_module.D2.create_DC7(_module.__default.fm10(_this.f3, _this.f3, _this.f3, globalState)));
          let _2213_v62;
          _2213_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2211_v60,_2212_v61);
          _2213_v62 = (_2213_v62).update((_dafny.MultiSet.fromElements(_2210_v59, p2)).Difference(_2211_v60), _2212_v61);
        }
        let _2214_v63;
        _2214_v63 = new BigNumber(-289);
        let _2215_v64;
        _2215_v64 = _dafny.Set.fromElements((_this).f4);
        let _2216_v65;
        _2216_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f2).length),_2215_v64);
        let _2217_v66;
        _2217_v66 = _dafny.MultiSet.fromElements(_this.f3);
        _2214_v63 = new BigNumber(((((_2216_v65).contains(new BigNumber((((_this.f3) ? (_2217_v66) : (_2217_v66))).cardinality()))) ? ((_2216_v65).get(new BigNumber((((_this.f3) ? (_2217_v66) : (_2217_v66))).cardinality()))) : (_2215_v64))).length);
      }
      if (_this.f3) {
        let _2218_v67;
        _2218_v67 = _dafny.Set.fromElements((_this).f4);
        (_this).f3 = (_2218_v67).IsProperSubsetOf(_2218_v67);
        _2144_v4 = (_this).fm14(!(_this.f3), ((_dafny.ZERO).minus(new BigNumber(((_this).f2).length))).isLessThan(new BigNumber(132)), _2140_v0, (_this).f4, globalState);
        (_this).f3 = _this.f3;
        _2140_v0 = _2140_v0;
        let _2219_v68;
        _2219_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3);
        let _2220_v69;
        let _nw297 = new _module.C3();
        _nw297.__ctor((((_2141_v1).contains(new BigNumber((_dafny.Seq.update(_2144_v4, _module.__default.safeIndex(new BigNumber((_2144_v4).length), new BigNumber((_2144_v4).length)), _2140_v0)).length))) ? ((_2141_v1).get(new BigNumber((_dafny.Seq.update(_2144_v4, _module.__default.safeIndex(new BigNumber((_2144_v4).length), new BigNumber((_2144_v4).length)), _2140_v0)).length))) : ((_this).f4)), (_this).f5, !((_dafny.ZERO).minus((_this).f4)).isEqualTo(new BigNumber((_2219_v68).length)), (_this).f1, _2144_v4);
        _2220_v69 = _nw297;
      } else {
        let _2221_v70;
        _2221_v70 = _dafny.Set.fromElements(_this.f3);
        let _2222_v71;
        _2222_v71 = _module.D14.create_DC39(p2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), ((_2223_v0) => function (_2224_i7) {
  return _2223_v0;
})(_2140_v0)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), ((_2225_v0) => function (_2226_i7) {
  return _2225_v0;
})(_2140_v0))).length)), _2140_v0), _2144_v4)).length)), new BigNumber((_2221_v70).length));
        let _source40 = _2222_v71;
        if (_source40.is_DC38) {
          let _2227___mcc_h0 = (_source40).cf64;
          let _2228_cf64 = _2227___mcc_h0;
          let _2229_v72;
          _2229_v72 = new BigNumber(-855);
          _2229_v72 = _2229_v72;
          _2140_v0 = _2140_v0;
          _2141_v1 = (_2141_v1).update(_2229_v72, p2);
          (_this).f3 = (_2221_v70).IsProperSubsetOf(_2221_v70);
        } else if (_source40.is_DC39) {
          let _2230___mcc_h1 = (_source40).cf65;
          let _2231___mcc_h2 = (_source40).cf66;
          let _2232___mcc_h3 = (_source40).cf67;
          let _2233_cf67 = _2232___mcc_h3;
          let _2234_cf66 = _2231___mcc_h2;
          let _2235_cf65 = _2230___mcc_h1;
          let _2236_v73;
          _2236_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
          let _2237_v74;
          let _2238_v75;
          let _2239_v76;
          let _2240_v77;
          let _out54;
          let _out55;
          let _out56;
          let _out57;
          let _outcollector13 = _module.__default.m0(_2236_v73, globalState);
          _out54 = _outcollector13[0];
          _out55 = _outcollector13[1];
          _out56 = _outcollector13[2];
          _out57 = _outcollector13[3];
          _2237_v74 = _out54;
          _2238_v75 = _out55;
          _2239_v76 = _out56;
          _2240_v77 = _out57;
          _2237_v74 = !(_2238_v75);
          (_this).f3 = _2237_v74;
          let _index342 = _module.__default.safeIndex(new BigNumber(759), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index342] = ((!(_2237_v74)) ? (_2238_v75) : (_this.f3));
          let _index343 = _module.__default.safeIndex(new BigNumber(759), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index343] = (true) && (_2237_v74);
        } else {
          let _2241___mcc_h4 = (_source40).cf63;
          let _2242_cf63 = _2241___mcc_h4;
          (_this).f3 = _this.f3;
          _2148_v5 = _2148_v5;
          let _2243_v78;
          let _nw298 = new _module.C1();
          _nw298.__ctor();
          _2243_v78 = _nw298;
          let _2244_v79;
          let _nw299 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2244_v79 = _nw299;
          let _2245_v80;
          _2245_v80 = _dafny.Map.Empty.slice().updateUnsafe(_2244_v79,((_this).f4).isLessThan(new BigNumber((_dafny.Seq.of((_this).f4)).length)));
          let _2246_v81;
          _2246_v81 = _dafny.MultiSet.fromElements(_this.f3, _this.f3, _this.f3, _this.f3, _this.f3);
          _2245_v80 = (_2245_v80).update(_2244_v79, !(_2246_v81).contains(_this.f3));
        }
        let _2247_v82;
        _2247_v82 = _module.D2.create_DC6(_2140_v0, _this.f3, _this.f3, new BigNumber((_dafny.Seq.of(_this.f3)).length));
        _2247_v82 = _2247_v82;
        let _2248_v83;
        let _nw300 = Array((new BigNumber(26)).toNumber());
        _2248_v83 = _nw300;
        _2248_v83 = _2248_v83;
        let _2249_v84;
        _2249_v84 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f4, new BigNumber((_2144_v4).length)),_2148_v5);
        r0 = (((_2249_v84).contains(p2)) ? ((_2249_v84).get(p2)) : (_2148_v5));
        let _index344 = _module.__default.safeIndex(new BigNumber(969), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index344] = _dafny.Seq.IsProperPrefixOf((_this).f2, _2144_v4);
        let _index345 = _module.__default.safeIndex(new BigNumber(969), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index345] = (_module.__default.safeModuloInt(p2, p2)).isLessThanOrEqualTo(p2);
      }
      let _2250_v85;
      _2250_v85 = _dafny.Map.Empty.slice().updateUnsafe(_2148_v5,_this.f3);
      let _index346 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index346] = (((_2250_v85).contains(_2148_v5)) ? ((_2250_v85).get(_2148_v5)) : (_this.f3));
      let _2251_v86;
      _2251_v86 = _module.D6.create_DC17(_this.f3);
      let _index347 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index347] = (_2251_v86).dtor_cf28;
      let _2252_v87;
      _2252_v87 = _dafny.Seq.of(new BigNumber((_2144_v4).length));
      let _2253_v88;
      _2253_v88 = _dafny.MultiSet.fromElements(!((new BigNumber((_2252_v87).length)).isLessThan((_this).f4)), _this.f3);
      let _2254_v89;
      let _nw301 = Array((new BigNumber(25)).toNumber()).fill([]);
      _2254_v89 = _nw301;
      let _2255_v90;
      let _init48 = ((_2256_v0) => function (_2257_i8) {
        return _2256_v0;
      })(_2140_v0);
      let _nw302 = Array((new BigNumber(2)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw302.length); _i0_48++) {
        _nw302[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _2255_v90 = _nw302;
      let _index348 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_2254_v89).length));
      (_2254_v89)[_index348] = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f5).length))]) ? (_2255_v90) : (_2255_v90));
      let _index349 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f5).length));
      let _index350 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_2254_v89).length));
      let _rhs368 = (_this).fm13(globalState);
      let _rhs369 = _2253_v88;
      let _rhs370 = _2255_v90;
      let _rhs371 = _this.f3;
      let _lhs236 = (_this).f5;
      let _lhs237 = _module.__default.safeIndex(new BigNumber(241), new BigNumber(((_this).f5).length));
      let _lhs238 = _2254_v89;
      let _lhs239 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_2254_v89).length));
      let _lhs240 = _this;
      _lhs236[_lhs237] = _rhs368;
      _2253_v88 = _rhs369;
      _lhs238[_lhs239] = _rhs370;
      _lhs240.f3 = _rhs371;
      r0 = _2148_v5;
      return r0;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
      this._f8 = [];
      this._f9 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f8, f9, f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return !(((_dafny.Set.fromElements(_dafny.Set.fromElements(true, _this.f3))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(_this.f3, false, _this.f3, _this.f3, _this.f3), _dafny.Set.fromElements(true)))).IsDisjointFrom(_dafny.Set.fromElements(_dafny.Set.fromElements(_this.f3), _dafny.Set.fromElements(false))));
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-895)), function (_2258_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      });
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_dafny.ZERO).minus(((_this).f4).plus((_this).f4)), (_this).f4);
    };
    fm11(p0, globalState) {
      let _this = this;
      return _this.f3;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _2259_v0;
      _2259_v0 = _dafny.Set.fromElements(new BigNumber(((_this).f2).length));
      let _2260_i0;
      _2260_i0 = _dafny.ZERO;
      L12: {
        while ((_2259_v0).IsDisjointFrom((_2259_v0).Union(_2259_v0))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2260_i0)) {
              break L12;
            }
            _2260_i0 = (_2260_i0).plus(_dafny.ONE);
            let _2261_v1;
            _2261_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,true);
            (_this).f3 = ((p1).minus(new BigNumber(322))).isEqualTo((new BigNumber(((_this).f2).length)).minus(new BigNumber((_2261_v1).length)));
            let _index351 = _module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f5).length));
            ((_this).f5)[_index351] = _this.f3;
            let _index352 = _module.__default.safeIndex(new BigNumber(587), new BigNumber(((_this).f5).length));
            ((_this).f5)[_index352] = _this.f3;
            r2 = ((_this.f3) ? ((_this).f4) : ((p0).multipliedBy(p1)));
            let _2262_v2;
            let _nw303 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2262_v2 = _nw303;
            let _2263_v3;
            _2263_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2262_v2,(_this).fm11(true, globalState));
            _2263_v3 = (_2263_v3).update(_2262_v2, (_2259_v0).contains(p0));
          }
        }
      }
      if ((_this).fm11(((_this).f4).isLessThan(p0), globalState)) {
        let _2264_v4;
        let _nw304 = Array((new BigNumber(20)).toNumber()).fill([]);
        _2264_v4 = _nw304;
        let _rhs372 = p0;
        let _rhs373 = _2264_v4;
        r2 = _rhs372;
        _2264_v4 = _rhs373;
        r2 = p0;
        let _2265_v5;
        _2265_v5 = _dafny.Seq.UnicodeFromString("bqtdfit");
        _2265_v5 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm2(_this.f3, globalState), (_this).f2), _2265_v5);
        let _2266_v6;
        _2266_v6 = _dafny.MultiSet.fromElements(new BigNumber(838), (_this).f4, new BigNumber(((_this).f2).length));
        let _2267_v7;
        _2267_v7 = _module.D4.create_DC12(_2266_v6);
        let _source41 = _2267_v7;
        if (_source41.is_DC13) {
          let _2268___mcc_h0 = (_source41).cf21;
          let _2269___mcc_h1 = (_source41).cf22;
          let _2270_cf22 = _2269___mcc_h1;
          let _2271_cf21 = _2268___mcc_h0;
          let _2272_v8;
          _2272_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-162),p0);
          let _2273_v9;
          _2273_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_2265_v5);
          let _2274_v10;
          _2274_v10 = _dafny.Seq.of(_2271_cf21);
          let _2275_v12;
          _2275_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("krvvypne"),p0);
          let _2276_v13;
          let _nw305 = Array((new BigNumber(23)).toNumber());
          _nw305[(_dafny.ZERO).toNumber()] = p0;
          _nw305[(_dafny.ONE).toNumber()] = p0;
          _nw305[(new BigNumber(2)).toNumber()] = (_this).f4;
          _nw305[(new BigNumber(3)).toNumber()] = p1;
          _nw305[(new BigNumber(4)).toNumber()] = p0;
          _nw305[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw305[(new BigNumber(6)).toNumber()] = p0;
          _nw305[(new BigNumber(7)).toNumber()] = p1;
          _nw305[(new BigNumber(8)).toNumber()] = p0;
          _nw305[(new BigNumber(9)).toNumber()] = p0;
          _nw305[(new BigNumber(10)).toNumber()] = (_this).f4;
          _nw305[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p0, p1, p0, p1, (_this).f4)).cardinality());
          _nw305[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw305[(new BigNumber(13)).toNumber()] = (((_2272_v8).contains((_this).f4)) ? ((_2272_v8).get((_this).f4)) : (p1));
          _nw305[(new BigNumber(14)).toNumber()] = p1;
          _nw305[(new BigNumber(15)).toNumber()] = p1;
          _nw305[(new BigNumber(16)).toNumber()] = new BigNumber((_2273_v9).length);
          _nw305[(new BigNumber(17)).toNumber()] = (_this).f4;
          _nw305[(new BigNumber(18)).toNumber()] = p0;
          _nw305[(new BigNumber(19)).toNumber()] = new BigNumber((_2274_v10).length);
          _nw305[(new BigNumber(20)).toNumber()] = new BigNumber((_2272_v8).length);
          _nw305[(new BigNumber(21)).toNumber()] = p0;
          _nw305[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of (_2275_v12).Keys.Elements) {
              let _2277_v11 = _compr_63;
              if ((_2275_v12).contains(_2277_v11)) {
                _coll63.push([_2277_v11,_2271_cf21]);
              }
            }
            return _coll63;
          }()).length)));
          _2276_v13 = _nw305;
          let _2278_v14;
          _2278_v14 = _module.D5.create_DC14(_2276_v13);
          let _pat_let_tv52 = _2276_v13;
          let _2279_v15;
          _2279_v15 = _dafny.Seq.of(_2276_v13, _2276_v13, _2276_v13, (function (_pat_let66_0) {
            return function (_2280_dt__update__tmp_h0) {
              return function (_pat_let67_0) {
                return function (_2281_dt__update_hcf23_h0) {
                  return _module.D5.create_DC14(_2281_dt__update_hcf23_h0);
                }(_pat_let67_0);
              }(_pat_let_tv52);
            }(_pat_let66_0);
          }(_2278_v14)).dtor_cf23);
          let _2282_v16;
          _2282_v16 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_2265_v5).length));
          let _2283_v17;
          _2283_v17 = _module.D2.create_DC5(_2282_v16);
          let _2284_v18;
          _2284_v18 = _dafny.Set.fromElements(_2283_v17);
          let _2285_v19;
          _2285_v19 = _module.D1.create_DC3(_2271_cf21, new BigNumber((_dafny.Seq.UnicodeFromString("jvsmecttx")).length), new BigNumber((_2284_v18).length));
          let _2286_v20;
          _2286_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f4, new BigNumber((_2265_v5).length))),(_2285_v19).dtor_cf3);
          let _index353 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_2276_v13).length));
          (_2276_v13)[_index353] = new BigNumber((_2270_cf22).length);
          let _index354 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_2276_v13).length));
          let _rhs374 = _dafny.Seq.update(_2279_v15, _module.__default.safeIndex((_dafny.ZERO).minus((new BigNumber(((function () {
            let _coll64 = new _dafny.Map();
            for (const _compr_64 of _dafny.IntegerRange(new BigNumber(608), new BigNumber(28))) {
              let _2287_v21 = _compr_64;
              if (((new BigNumber(608)).isLessThanOrEqualTo(_2287_v21)) && ((_2287_v21).isLessThan(new BigNumber(28)))) {
                _coll64.push([(_2287_v21).plus((((_2282_v16).contains(_this.f3)) ? ((_2282_v16).get(_this.f3)) : ((_this).f4))),_this.f3]);
              }
            }
            return _coll64;
          }()).update((_this).f4, _module.__default.fm3(new BigNumber((_2259_v0).length), globalState))).length)).multipliedBy(p0)), new BigNumber((_2279_v15).length)), _2276_v13);
          let _rhs375 = _2286_v20;
          let _rhs376 = _module.__default.fm0(globalState);
          let _lhs241 = _2276_v13;
          let _lhs242 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_2276_v13).length));
          _2279_v15 = _rhs374;
          _2286_v20 = _rhs375;
          _lhs241[_lhs242] = _rhs376;
          let _2288_v22;
          _2288_v22 = _dafny.MultiSet.fromElements(_2271_cf21, !(_2271_cf21), _2271_cf21);
          let _2289_v23;
          _2289_v23 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_2271_cf21) || (_this.f3)),_2288_v22);
          _2289_v23 = _2289_v23;
          let _2290_v24;
          let _2291_v25;
          let _out58;
          let _out59;
          let _outcollector14 = (_this).m5(globalState);
          _out58 = _outcollector14[0];
          _out59 = _outcollector14[1];
          _2290_v24 = _out58;
          _2291_v25 = _out59;
          let _2292_v26;
          let _nw306 = new _module.C0();
          _nw306.__ctor(_2276_v13);
          _2292_v26 = _nw306;
        } else {
          let _2293___mcc_h2 = (_source41).cf20;
          let _2294_cf20 = _2293___mcc_h2;
          let _2295_v27;
          _2295_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2265_v5);
          let _2296_v28;
          _2296_v28 = _dafny.Set.fromElements(_this.f3, _this.f3);
          let _2297_v29;
          _2297_v29 = _dafny.Seq.of((_this).fm13(globalState), _this.f3);
          let _2298_v30;
          let _nw307 = new _module.C5();
          _nw307.__ctor(true, (_this).f1, (_this).f2);
          _2298_v30 = _nw307;
          let _2299_v31;
          _2299_v31 = _dafny.MultiSet.fromElements(_this.f3);
          let _2300_v32;
          _2300_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2298_v30,new BigNumber((_2299_v31).cardinality()));
          let _2301_v33;
          let _nw308 = Array((new BigNumber(15)).toNumber());
          _nw308[(_dafny.ZERO).toNumber()] = p0;
          _nw308[(_dafny.ONE).toNumber()] = p1;
          _nw308[(new BigNumber(2)).toNumber()] = p0;
          _nw308[(new BigNumber(3)).toNumber()] = p1;
          _nw308[(new BigNumber(4)).toNumber()] = new BigNumber((_2295_v27).length);
          _nw308[(new BigNumber(5)).toNumber()] = new BigNumber(451);
          _nw308[(new BigNumber(6)).toNumber()] = new BigNumber((_2296_v28).length);
          _nw308[(new BigNumber(7)).toNumber()] = p0;
          _nw308[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(p1);
          _nw308[(new BigNumber(9)).toNumber()] = (_this).f4;
          _nw308[(new BigNumber(10)).toNumber()] = new BigNumber((_2294_cf20).cardinality());
          _nw308[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_2302_i1) {
            return (_this).f4;
          })).length);
          _nw308[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), function (_2303_i2) {
            return (_this).f4;
          })).length);
          _nw308[(new BigNumber(13)).toNumber()] = new BigNumber((_2297_v29).length);
          _nw308[(new BigNumber(14)).toNumber()] = new BigNumber(((_2300_v32).update(_2298_v30, p1)).length);
          _2301_v33 = _nw308;
          let _2304_v34;
          let _nw309 = new _module.C0();
          _nw309.__ctor(_2301_v33);
          _2304_v34 = _nw309;
          (_this).f3 = !(((_this.f3) ? (_this.f3) : (true))) || ((_module.__default.fm0(globalState)).isLessThan((_this).f4));
          let _index355 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_2301_v33).length));
          (_2301_v33)[_index355] = (_this).f4;
          let _index356 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_2301_v33).length));
          (_2301_v33)[_index356] = new BigNumber(942);
          let _2305_v35;
          _2305_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2264_v4);
          _2305_v35 = (_2305_v35).update((new BigNumber((_dafny.Seq.of(p1, p0)).length)).isLessThanOrEqualTo(new BigNumber((_module.__default.fm2(_this.f3, globalState)).length)), _2264_v4);
        }
        let _2306_v36;
        let _nw310 = Array((new BigNumber(26)).toNumber()).fill([]);
        _2306_v36 = _nw310;
        let _2307_v37;
        _2307_v37 = _dafny.Seq.of(_2306_v36, _2306_v36);
        let _2308_v38;
        _2308_v38 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2309_v39;
        _2309_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2308_v38,_this.f3);
        let _2310_v41;
        let _nw311 = Array((new BigNumber(5)).toNumber());
        _nw311[(_dafny.ZERO).toNumber()] = _2306_v36;
        _nw311[(_dafny.ONE).toNumber()] = _2306_v36;
        _nw311[(new BigNumber(2)).toNumber()] = _2306_v36;
        _nw311[(new BigNumber(3)).toNumber()] = (_2307_v37)[_module.__default.safeIndex(new BigNumber((function () {
          let _coll65 = new _dafny.Set();
          for (const _compr_65 of ((_2309_v39).update(new _dafny.CodePoint('v'.codePointAt(0)), _this.f3)).Keys.Elements) {
            let _2311_v40 = _compr_65;
            if (((_2309_v39).update(new _dafny.CodePoint('v'.codePointAt(0)), _this.f3)).contains(_2311_v40)) {
              _coll65.add(_2311_v40);
            }
          }
          return _coll65;
        }()).length), new BigNumber((_2307_v37).length))];
        _nw311[(new BigNumber(4)).toNumber()] = _2306_v36;
        _2310_v41 = _nw311;
        let _index357 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_2310_v41).length));
        (_2310_v41)[_index357] = _2306_v36;
        let _index358 = _module.__default.safeIndex(new BigNumber(196), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index358] = _this.f3;
        let _index359 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_this).f8).length));
        ((_this).f8)[_index359] = _this.f3;
        let _2312_v42;
        _2312_v42 = _dafny.Seq.of(_this.f3, _this.f3);
        let _index360 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_2310_v41).length));
        let _index361 = _module.__default.safeIndex(new BigNumber(196), new BigNumber(((_this).f5).length));
        let _index362 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_this).f8).length));
        let _rhs377 = _2306_v36;
        let _rhs378 = _this.f3;
        let _rhs379 = _this.f3;
        let _rhs380 = (_2312_v42)[_module.__default.safeIndex(new BigNumber(((_2266_v6).update(new BigNumber(245), _module.__default.abs(p0))).cardinality()), new BigNumber((_2312_v42).length))];
        let _rhs381 = (p0).minus((p1).plus((_this).f4));
        let _lhs243 = _2310_v41;
        let _lhs244 = _module.__default.safeIndex(new BigNumber(238), new BigNumber((_2310_v41).length));
        let _lhs245 = (_this).f5;
        let _lhs246 = _module.__default.safeIndex(new BigNumber(196), new BigNumber(((_this).f5).length));
        let _lhs247 = (_this).f8;
        let _lhs248 = _module.__default.safeIndex(new BigNumber(695), new BigNumber(((_this).f8).length));
        let _lhs249 = _this;
        _lhs243[_lhs244] = _rhs377;
        _lhs245[_lhs246] = _rhs378;
        _lhs247[_lhs248] = _rhs379;
        _lhs249.f3 = _rhs380;
        r1 = _rhs381;
      } else {
        r1 = ((_dafny.ZERO).minus((_this).f4)).minus((_this).f4);
        let _2313_v43;
        let _nw312 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _2313_v43 = _nw312;
        let _index363 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
        (_2313_v43)[_index363] = (_this).f4;
        let _2314_v44;
        _2314_v44 = _dafny.Set.fromElements(true);
        let _index364 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
        (_2313_v43)[_index364] = (_module.__default.safeDivisionInt(p1, p1)).multipliedBy(new BigNumber((_2314_v44).length));
        let _2315_v45;
        _2315_v45 = new _dafny.CodePoint('m'.codePointAt(0));
        let _2316_v46;
        _2316_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
        let _2317_v47;
        _2317_v47 = _module.D8.create_DC24(_2315_v45, (_2316_v46).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3)));
        let _source42 = _2317_v47;
        if (_source42.is_DC23) {
          let _2318___mcc_h3 = (_source42).cf38;
          let _2319___mcc_h4 = (_source42).cf39;
          let _2320___mcc_h5 = (_source42).cf40;
          let _2321_cf40 = _2320___mcc_h5;
          let _2322_cf39 = _2319___mcc_h4;
          let _2323_cf38 = _2318___mcc_h3;
          let _2324_v48;
          _2324_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-766),_this.f3);
          _2321_cf40 = _module.__default.fm3(_module.__default.safeDivisionInt(new BigNumber((_2324_v48).length), p0), globalState);
          let _2325_v49;
          let _2326_v50;
          let _out60;
          let _out61;
          let _outcollector15 = (_this).m5(globalState);
          _out60 = _outcollector15[0];
          _out61 = _outcollector15[1];
          _2325_v49 = _out60;
          _2326_v50 = _out61;
          _2326_v50 = !(_2325_v49);
          let _2327_v51;
          _2327_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2313_v43,((_2325_v49) ? ((_this).f4) : ((_this).f4)));
          let _2328_v52;
          _2328_v52 = _dafny.Seq.of(_2313_v43);
          _2327_v51 = (_2327_v51).update((_2328_v52)[_module.__default.safeIndex((_this).f4, new BigNumber((_2328_v52).length))], (_dafny.ZERO).minus(p0));
        } else if (_source42.is_DC24) {
          let _2329___mcc_h6 = (_source42).cf41;
          let _2330___mcc_h7 = (_source42).cf42;
          let _2331_cf42 = _2330___mcc_h7;
          let _2332_cf41 = _2329___mcc_h6;
          let _2333_v53;
          _2333_v53 = _dafny.Seq.of(p0, (_this).f4, p1);
          _2333_v53 = _dafny.Seq.Concat(_2333_v53, _2333_v53);
          let _index365 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index365] = false;
          let _index366 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index366] = _module.__default.fm3((_this).f4, globalState);
          let _2334_v54;
          _2334_v54 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))],(_this).f2);
          let _index367 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          let _index368 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
          let _index369 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          let _rhs382 = !(_this.f3) || ((_2314_v44).IsSubsetOf(_2314_v44));
          let _rhs383 = _module.__default.fm22(((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))], (new BigNumber(((((_2334_v54).contains(true)) ? ((_2334_v54).get(true)) : (_dafny.Seq.UnicodeFromString("k")))).length)).plus((_this).f4), globalState);
          let _rhs384 = _module.__default.fm3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2332_cf41,(_this).f4)).length), globalState);
          let _rhs385 = (_2313_v43)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length))];
          let _rhs386 = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))]) ? (false) : (((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))]));
          let _lhs250 = _this;
          let _lhs251 = (_this).f5;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          let _lhs253 = _2313_v43;
          let _lhs254 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
          let _lhs255 = (_this).f5;
          let _lhs256 = _module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length));
          _lhs250.f3 = _rhs382;
          _2332_cf41 = _rhs383;
          _lhs251[_lhs252] = _rhs384;
          _lhs253[_lhs254] = _rhs385;
          _lhs255[_lhs256] = _rhs386;
          let _2335_v55;
          _2335_v55 = _module.D9.create_DC27(new BigNumber(-385));
          let _2336_v56;
          _2336_v56 = _dafny.Seq.of(_this.f3, _this.f3, _module.__default.fm3((_2313_v43)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length))], globalState));
          let _2337_v57;
          _2337_v57 = _dafny.Seq.of(_2336_v56, _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(246), new BigNumber(((_this).f5).length))]), _2336_v56, _2336_v56, _2336_v56);
          let _2338_v58;
          let _nw313 = new _module.C8();
          _nw313.__ctor(_2337_v57, _this.f3, (_this).f1, _dafny.Seq.UnicodeFromString("imqca"), (_this).f4, (_this).f8);
          _2338_v58 = _nw313;
          let _2339_v59;
          _2339_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2335_v55,_2338_v58);
          _2339_v59 = (_2339_v59).update(_2335_v55, _2338_v58);
        } else {
          let _2340___mcc_h8 = (_source42).cf37;
          let _2341_cf37 = _2340___mcc_h8;
          let _2342_v60;
          _2342_v60 = _dafny.Seq.of(new BigNumber((_2316_v46).length), p0);
          _2316_v46 = (_2316_v46).update(_this.f3, (_2259_v0).IsSubsetOf(_dafny.Set.fromElements(p0, new BigNumber((_2341_cf37).cardinality()), p1, new BigNumber((_2342_v60).length), p0)));
          _2313_v43 = _2313_v43;
          r2 = (_module.__default.safeDivisionInt(p1, (_this).f4)).minus(_module.__default.safeModuloInt((_2313_v43)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length))], p1));
          let _index370 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
          (_2313_v43)[_index370] = _module.__default.fm0(globalState);
        }
        r2 = (_this).f4;
        let _index371 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length));
        (_2313_v43)[_index371] = (_2313_v43)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_2313_v43).length))];
      }
      let _2343_i3;
      _2343_i3 = _dafny.ZERO;
      L13: {
        while ((p1).isLessThan(p0)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2343_i3)) {
              break L13;
            }
            _2343_i3 = (_2343_i3).plus(_dafny.ONE);
            (_this).f3 = _this.f3;
            let _2344_v61;
            let _nw314 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
            _2344_v61 = _nw314;
            let _2345_v62;
            _2345_v62 = _dafny.Seq.of(_this.f3);
            let _rhs387 = !(true);
            let _rhs388 = (_this).f4;
            let _rhs389 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_2345_v62, _module.__default.safeIndex((_this).f4, new BigNumber((_2345_v62).length)), true), _dafny.Seq.Concat(_2345_v62, _module.__default.fm7(_this.f3, _this.f3, globalState)));
            let _rhs390 = _2344_v61;
            let _lhs257 = _this;
            let _lhs258 = _this;
            _lhs257.f3 = _rhs387;
            r2 = _rhs388;
            _lhs258.f3 = _rhs389;
            _2344_v61 = _rhs390;
            let _2346_v63;
            _2346_v63 = _dafny.Seq.of((_this).f2, (_this).f2);
            let _2347_v64;
            _2347_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe((_2346_v63)[_module.__default.safeIndex(p1, new BigNumber((_2346_v63).length))],_this.f3));
            let _2348_v65;
            let _nw315 = new _module.C8();
            _nw315.__ctor(_dafny.Seq.of(_2345_v62, _module.__default.fm7(_this.f3, (_module.__default.fm10(true, _this.f3, _this.f3, globalState)).dtor_cf11, globalState)), _this.f3, (((_2347_v64).contains(p0)) ? ((_2347_v64).get(p0)) : ((_this).f1)), ((_this.f3) ? ((_this).f2) : ((_this).f2)), new BigNumber(451), (_this).f8);
            _2348_v65 = _nw315;
            _2348_v65 = _2348_v65;
            let _2349_v66;
            _2349_v66 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f4);
            let _2350_v67;
            _2350_v67 = _module.D2.create_DC5(_2349_v66);
            let _2351_v68;
            _2351_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2349_v66,_2350_v67);
            _2351_v68 = (_2351_v68).update(_2349_v66, _module.__default.fm55(_this.f3, (_this).f2, _dafny.Seq.UnicodeFromString("xfxldc"), globalState));
          }
        }
      }
      if (_this.f3) {
        let _2352_v69;
        _2352_v69 = _dafny.MultiSet.fromElements(p1, _module.__default.fm0(globalState));
        _2352_v69 = _2352_v69;
        let _index372 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
        ((_this).f8)[_index372] = _this.f3;
        let _index373 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
        ((_this).f8)[_index373] = _this.f3;
        let _2353_v70;
        _2353_v70 = _dafny.Seq.of(p1);
        if (!(_dafny.Seq.IsPrefixOf(_2353_v70, _2353_v70))) {
          let _2354_v71;
          let _nw316 = new _module.C5();
          _nw316.__ctor(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))], (_this).f1, _dafny.Seq.Concat((_this).f2, _dafny.Seq.UnicodeFromString("xlopb")));
          _2354_v71 = _nw316;
          let _2355_v72;
          _2355_v72 = new _dafny.CodePoint('e'.codePointAt(0));
          let _index374 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          let _rhs391 = (_2354_v71).f15;
          let _rhs392 = _2355_v72;
          let _rhs393 = (_2354_v71).fm11(false, globalState);
          let _lhs259 = (_this).f8;
          let _lhs260 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          let _lhs261 = _this;
          _lhs259[_lhs260] = _rhs391;
          _2355_v72 = _rhs392;
          _lhs261.f3 = _rhs393;
          r1 = (_dafny.ZERO).minus(((!(!(_this.f3))) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p0)))) : (_module.__default.safeDivisionInt((_this).f4, p1))));
          let _2356_v73;
          let _init49 = ((_2357_v71) => function (_2358_i4) {
            return !((_dafny.Set.fromElements(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))], (_2357_v71).f15)).IsProperSubsetOf(_dafny.Set.fromElements(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))])));
          })(_2354_v71);
          let _nw317 = Array((new BigNumber(4)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw317.length); _i0_49++) {
            _nw317[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _2356_v73 = _nw317;
          let _2359_v74;
          let _init50 = ((_2360_p1) => function (_2361_i5) {
            return (_2361_i5).multipliedBy((_module.D8.create_DC23(new BigNumber(392), _2360_p1, _this.f3)).dtor_cf38);
          })(p1);
          let _nw318 = Array((new BigNumber(29)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw318.length); _i0_50++) {
            _nw318[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _2359_v74 = _nw318;
          let _2362_v75;
          _2362_v75 = _module.D5.create_DC14(_2359_v74);
          let _2363_v76;
          _2363_v76 = _dafny.Map.Empty.slice().updateUnsafe(((((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]) ? (_2359_v74) : ((_2362_v75).dtor_cf23)),_module.__default.safeDivisionInt((_this).f4, p0));
          let _2364_v77;
          _2364_v77 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))],new BigNumber((_2259_v0).length));
          let _2365_v78;
          _2365_v78 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
          let _2366_v79;
          _2366_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2259_v0).length),(((_2365_v78).contains(p0)) ? ((_2365_v78).get(p0)) : ((_2354_v71).f15)));
          let _index375 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          let _index376 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          let _rhs394 = true;
          let _rhs395 = (_2354_v71).f15;
          let _rhs396 = (_this).f5;
          let _rhs397 = (_2363_v76).update(_2359_v74, _module.__default.safeModuloInt(new BigNumber((_2364_v77).length), new BigNumber(((_2366_v79).update(p1, ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))])).length)));
          let _rhs398 = (_dafny.ZERO).minus(new BigNumber(-889));
          let _lhs262 = (_this).f8;
          let _lhs263 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          let _lhs264 = (_this).f8;
          let _lhs265 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
          _lhs262[_lhs263] = _rhs394;
          _lhs264[_lhs265] = _rhs395;
          _2356_v73 = _rhs396;
          _2363_v76 = _rhs397;
          r2 = _rhs398;
          let _2367_v80;
          _2367_v80 = _dafny.Seq.of(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))], ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          let _2368_v81;
          _2368_v81 = _dafny.Seq.of(_2367_v80);
          let _2369_v82;
          _2369_v82 = _module.D16.create_DC41(_2368_v81);
          let _2370_v83;
          _2370_v83 = _module.D5.create_DC15(p1, false, (_this).f5);
          let _2371_v84;
          _2371_v84 = _dafny.Map.Empty.slice().updateUnsafe((_2354_v71).f15,(_2370_v83).dtor_cf26);
          let _2372_v85;
          let _nw319 = new _module.C8();
          _nw319.__ctor((_2369_v82).dtor_cf69, !(_this.f3) || (_this.f3), ((_this).f1).Merge((_this).f1), _dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber(290), new BigNumber(((_this).f2).length)), _2355_v72), p0, (((_2371_v84).contains(_this.f3)) ? ((_2371_v84).get(_this.f3)) : (_2356_v73)));
          _2372_v85 = _nw319;
        } else {
          let _2373_v86;
          _2373_v86 = _dafny.Seq.of(_module.__default.fm3(new BigNumber(853), globalState), ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          let _2374_v87;
          _2374_v87 = new _dafny.CodePoint('o'.codePointAt(0));
          let _2375_v88;
          _2375_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          let _2376_v89;
          _2376_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2374_v87,new BigNumber((_2375_v88).length));
          let _2377_v90;
          _2377_v90 = _dafny.MultiSet.fromElements(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))], _this.f3);
          let _2378_v91;
          _2378_v91 = _module.D8.create_DC22(_2377_v90);
          let _2379_v92;
          _2379_v92 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f3);
          let _2380_v93;
          _2380_v93 = _dafny.Map.Empty.slice().updateUnsafe(_2374_v87,_this.f3);
          let _2381_v94;
          let _nw320 = Array((new BigNumber(28)).toNumber());
          _nw320[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f2).length));
          _nw320[(_dafny.ONE).toNumber()] = p1;
          _nw320[(new BigNumber(2)).toNumber()] = p1;
          _nw320[(new BigNumber(3)).toNumber()] = p0;
          _nw320[(new BigNumber(4)).toNumber()] = new BigNumber((_2379_v92).length);
          _nw320[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(6)).toNumber()] = p0;
          _nw320[(new BigNumber(7)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(8)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(9)).toNumber()] = p1;
          _nw320[(new BigNumber(10)).toNumber()] = p0;
          _nw320[(new BigNumber(11)).toNumber()] = p1;
          _nw320[(new BigNumber(12)).toNumber()] = new BigNumber(477);
          _nw320[(new BigNumber(13)).toNumber()] = p0;
          _nw320[(new BigNumber(14)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2353_v70).length));
          _nw320[(new BigNumber(16)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(17)).toNumber()] = p0;
          _nw320[(new BigNumber(18)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(19)).toNumber()] = new BigNumber(-705);
          _nw320[(new BigNumber(20)).toNumber()] = new BigNumber((_2380_v93).length);
          _nw320[(new BigNumber(21)).toNumber()] = p1;
          _nw320[(new BigNumber(22)).toNumber()] = p0;
          _nw320[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw320[(new BigNumber(24)).toNumber()] = p1;
          _nw320[(new BigNumber(25)).toNumber()] = (_this).f4;
          _nw320[(new BigNumber(26)).toNumber()] = new BigNumber((_2379_v92).length);
          _nw320[(new BigNumber(27)).toNumber()] = (((_2352_v69).contains(new BigNumber(((_this).f2).length))) ? ((_2352_v69).get(new BigNumber(((_this).f2).length))) : ((_dafny.ZERO).minus(p1)));
          _2381_v94 = _nw320;
          let _2382_v95;
          _2382_v95 = _dafny.Map.Empty.slice().updateUnsafe(false,_2381_v94);
          let _2383_v96;
          let _nw321 = Array((new BigNumber(26)).toNumber());
          _nw321[(_dafny.ZERO).toNumber()] = (_this).fm13(globalState);
          _nw321[(_dafny.ONE).toNumber()] = !(!(_this.f3));
          _nw321[(new BigNumber(2)).toNumber()] = (p1).isEqualTo(p1);
          _nw321[(new BigNumber(3)).toNumber()] = ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))];
          _nw321[(new BigNumber(4)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(5)).toNumber()] = (((_2373_v86)[_module.__default.safeIndex(p0, new BigNumber((_2373_v86).length))]) ? (_this.f3) : (((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]));
          _nw321[(new BigNumber(6)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(7)).toNumber()] = !(_this.f3) || (((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          _nw321[(new BigNumber(8)).toNumber()] = ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))];
          _nw321[(new BigNumber(9)).toNumber()] = !(_dafny.Seq.IsPrefixOf(_2353_v70, _dafny.Seq.of((((_2376_v89).contains(_2374_v87)) ? ((_2376_v89).get(_2374_v87)) : ((_this).f4)), p0, new BigNumber(109), new BigNumber((((_2378_v91).dtor_cf37).update(_this.f3, _module.__default.abs(p0))).cardinality()))));
          _nw321[(new BigNumber(10)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(11)).toNumber()] = !(!(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]) || (_this.f3));
          _nw321[(new BigNumber(12)).toNumber()] = (((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]) || (false);
          _nw321[(new BigNumber(13)).toNumber()] = ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))];
          _nw321[(new BigNumber(14)).toNumber()] = (p0).isLessThan(new BigNumber(-29));
          _nw321[(new BigNumber(15)).toNumber()] = (new BigNumber(909)).isEqualTo(new BigNumber((_2382_v95).length));
          _nw321[(new BigNumber(16)).toNumber()] = (((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]) === (_this.f3);
          _nw321[(new BigNumber(17)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(18)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(19)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(20)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(21)).toNumber()] = !(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          _nw321[(new BigNumber(22)).toNumber()] = _this.f3;
          _nw321[(new BigNumber(23)).toNumber()] = _module.__default.fm3(new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-803)), new BigNumber(((_this).f2).length)), _2374_v87)).length), globalState);
          _nw321[(new BigNumber(24)).toNumber()] = !(_this.f3);
          _nw321[(new BigNumber(25)).toNumber()] = _this.f3;
          _2383_v96 = _nw321;
          _2383_v96 = _2383_v96;
          let _2384_v97;
          _2384_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_module.D8.create_DC23(p1, p0, !(((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))])));
          let _2385_v98;
          _2385_v98 = _module.D8.create_DC23((_this).f4, p1, ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))]);
          _2384_v97 = (_2384_v97).update((_this).f4, _2385_v98);
          let _2386_v99;
          _2386_v99 = _dafny.Seq.of(_dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f2).length)), _2374_v87));
          let _2387_v100;
          _2387_v100 = _dafny.Map.Empty.slice().updateUnsafe(_2386_v99,p1);
          _2387_v100 = (_2387_v100).update(_2386_v99, new BigNumber(524));
          let _2388_v101;
          _2388_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((_this).f2, (_this).f2),p1);
          _2388_v101 = (_2388_v101).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(969)), function (_2389_i6) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), (_dafny.ZERO).minus(p1));
          (_this).f3 = ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))];
        }
        let _index377 = _module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length));
        ((_this).f8)[_index377] = ((_this).f8)[_module.__default.safeIndex(new BigNumber(15), new BigNumber(((_this).f8).length))];
        let _2390_v102;
        _2390_v102 = _dafny.Seq.of((_this).f5, (_this).f5, (_this).f5);
        _2390_v102 = _dafny.Seq.Concat(_2390_v102, _dafny.Seq.Concat(_2390_v102, _2390_v102));
      } else {
        let _2391_v103;
        let _nw322 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _2391_v103 = _nw322;
        let _2392_v104;
        let _nw323 = new _module.C0();
        _nw323.__ctor(_2391_v103);
        _2392_v104 = _nw323;
        r2 = (p0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("aegsjgofj")).length));
        (_this).f3 = (_this).fm13(globalState);
        let _2393_v105;
        _2393_v105 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f5);
        let _2394_v106;
        _2394_v106 = _module.D5.create_DC15(p0, _this.f3, (_this).f8);
        _2393_v105 = (_2393_v105).update(_this.f3, (_2394_v106).dtor_cf26);
        r2 = _module.__default.fm0(globalState);
      }
      r1 = p0;
      let _2395_v107;
      let _nw324 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2395_v107 = _nw324;
      let _2396_v108;
      _2396_v108 = _dafny.MultiSet.fromElements(_2395_v107);
      let _2397_v109;
      _2397_v109 = _dafny.Map.Empty.slice().updateUnsafe(_2396_v108,!(_module.__default.fm3(p0, globalState)));
      let _2398_v110;
      _2398_v110 = _dafny.Seq.of(false);
      let _2399_v111;
      let _nw325 = new _module.C6();
      _nw325.__ctor(p1, _module.__default.safeModuloInt(p1, new BigNumber(843)), p1, (_this).f5, (((_2397_v109).contains(_2396_v108)) ? ((_2397_v109).get(_2396_v108)) : ((_2398_v110)[_module.__default.safeIndex((_this).f4, new BigNumber((_2398_v110).length))])), (_dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3)).Merge((_this).f1), (_this).f2);
      _2399_v111 = _nw325;
      let _2400_v112;
      _2400_v112 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f3);
      r0 = ((_2400_v112).Merge(_2400_v112)).Merge(_module.__default.fm9(globalState));
      r1 = (p1).multipliedBy((p1).minus((_this).f4));
      r2 = (_this).f4;
      return [r0, r1, r2];
    }
    m5(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      if (false) {
        r0 = _this.f3;
        (_this).f3 = _this.f3;
        let _2401_v0;
        _2401_v0 = _dafny.Seq.of(_this.f3, _this.f3, _this.f3, _this.f3, _this.f3);
        let _2402_v1;
        _2402_v1 = _dafny.Seq.of(_2401_v0, _2401_v0);
        let _2403_v2;
        _2403_v2 = _module.D1.create_DC2((_this).f2);
        let _2404_v3;
        _2404_v3 = _dafny.Set.fromElements((_this).f2, (_2403_v2).dtor_cf2, (_this).f2);
        let _2405_v4;
        _2405_v4 = _module.D13.create_DC34(!(_this.f3), _this.f3, (_this).f4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), function (_2406_i1) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}));
        let _2407_v5;
        let _nw326 = new _module.C8();
        _nw326.__ctor(_dafny.Seq.Concat(_2402_v1, _2402_v1), (_2404_v3).IsDisjointFrom(_2404_v3), (_this).f1, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(696)), function (_2408_i0) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        }), (_2405_v4).dtor_cf56), (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4)), (_this).f5);
        _2407_v5 = _nw326;
        let _2409_v6;
        _2409_v6 = _module.D11.create_DC31((_this).f4);
        let _2410_v7;
        _2410_v7 = new BigNumber(479);
        let _2411_v8;
        let _nw327 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _2411_v8 = _nw327;
        let _2412_v9;
        _2412_v9 = _dafny.MultiSet.fromElements(_2411_v8);
        let _2413_v10;
        _2413_v10 = _dafny.MultiSet.fromElements(_this.f3);
        let _2414_v11;
        _2414_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _2415_v12;
        _2415_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
        let _2416_v13;
        let _nw328 = Array((new BigNumber(12)).toNumber());
        _nw328[(_dafny.ZERO).toNumber()] = new BigNumber(-526);
        _nw328[(_dafny.ONE).toNumber()] = (_2407_v5).fm12((_this).f4, _2410_v7, (_this).f4, false, globalState);
        _nw328[(new BigNumber(2)).toNumber()] = _2410_v7;
        _nw328[(new BigNumber(3)).toNumber()] = (_this).f4;
        _nw328[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_this).f4);
        _nw328[(new BigNumber(5)).toNumber()] = _2410_v7;
        _nw328[(new BigNumber(6)).toNumber()] = _2410_v7;
        _nw328[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((((_2412_v9).contains(_2411_v8)) ? ((_2412_v9).get(_2411_v8)) : ((((_2413_v10).contains(_this.f3)) ? ((_2413_v10).get(_this.f3)) : (_2410_v7)))));
        _nw328[(new BigNumber(8)).toNumber()] = _2410_v7;
        _nw328[(new BigNumber(9)).toNumber()] = (_this).f4;
        _nw328[(new BigNumber(10)).toNumber()] = (new BigNumber(((_module.D8.create_DC24(_2414_v11, _2415_v12)).dtor_cf42).length)).plus(_2410_v7);
        _nw328[(new BigNumber(11)).toNumber()] = new BigNumber((_2415_v12).length);
        _2416_v13 = _nw328;
        let _index378 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2411_v8).length));
        (_2411_v8)[_index378] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_2417_i2) {
          return _module.D1.create_DC4(_this.f3, _this.f3);
        }), _dafny.Seq.of(_module.D1.create_DC4(_this.f3, _this.f3)))).length);
        let _index379 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2411_v8).length));
        let _rhs399 = (((_2413_v10).IsDisjointFrom(_2413_v10)) ? (_module.__default.fm56(_2410_v7, globalState)) : (_2409_v6));
        let _rhs400 = _module.__default.safeModuloInt((_this).f4, (_this).f4);
        let _rhs401 = _2410_v7;
        let _rhs402 = _2410_v7;
        let _lhs266 = _2411_v8;
        let _lhs267 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2411_v8).length));
        _2409_v6 = _rhs399;
        _2410_v7 = _rhs400;
        _2410_v7 = _rhs401;
        _lhs266[_lhs267] = _rhs402;
        let _2418_v14;
        let _nw329 = new _module.C8();
        _nw329.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm7(_this.f3, _this.f3, globalState), _2401_v0, _2401_v0, _2401_v0), _dafny.Seq.update(_dafny.Seq.of(_2401_v0), _module.__default.safeIndex(_2410_v7, new BigNumber((_dafny.Seq.of(_2401_v0)).length)), _2401_v0)), (_2401_v0)[_module.__default.safeIndex((_2411_v8)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_2411_v8).length))], new BigNumber((_2401_v0).length))], (_this).f1, (_this).f2, (_2411_v8)[_module.__default.safeIndex(new BigNumber(71), new BigNumber((_2411_v8).length))], (_this).f5);
        _2418_v14 = _nw329;
      } else {
        let _2419_v15;
        let _nw330 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _2419_v15 = _nw330;
        let _index380 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2419_v15).length));
        (_2419_v15)[_index380] = (_this).f4;
        let _index381 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2419_v15).length));
        (_2419_v15)[_index381] = (_this).f4;
        let _2420_v16;
        _2420_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3);
        let _2421_v17;
        let _nw331 = new _module.C7();
        _nw331.__ctor((_this).f4, (_this).f5, (((_2420_v16).contains((_2419_v15)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_2419_v15).length))])) ? ((_2420_v16).get((_2419_v15)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_2419_v15).length))])) : (_this.f3)), ((_this).f1).Merge((_this).f1), (_this).f2);
        _2421_v17 = _nw331;
        let _2422_v18;
        _2422_v18 = _dafny.Set.fromElements(_this.f3);
        let _2423_v19;
        _2423_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2422_v18).length),(_this).f5);
        let _2424_v20;
        _2424_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2423_v19,_this.f3);
        _2424_v20 = (_2424_v20).update((_2423_v19).Merge(_2423_v19), (_this.f3) || (_this.f3));
        let _2425_v21;
        _2425_v21 = new _dafny.CodePoint('l'.codePointAt(0));
        let _index382 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_2419_v15).length));
        (_2419_v15)[_index382] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm57(_2425_v21, globalState)).length));
        r0 = ((_this.f3) ? (_this.f3) : (_this.f3));
      }
      let _2426_v22;
      let _init51 = function (_2427_i4) {
        return (_2427_i4).plus((_this).f4);
      };
      let _nw332 = Array((new BigNumber(4)).toNumber());
      for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw332.length); _i0_51++) {
        _nw332[_i0_51] = _init51(new BigNumber(_i0_51));
      }
      _2426_v22 = _nw332;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2426_v22).length))) {
        let _2428_i3 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2428_i3)) && ((_2428_i3).isLessThan(new BigNumber((_2426_v22).length))))) {
          (_2426_v22)[(_2428_i3)] = _module.__default.safeDivisionInt(_2428_i3, ((_this).f4).multipliedBy((_this).f4));
        }
      }
      let _2429_v23;
      let _nw333 = new _module.C6();
      _nw333.__ctor((_this).f4, (_this).f4, _module.__default.fm0(globalState), (_this).f5, true, ((_this).f1).Merge((_this).f1), (_this).f2);
      _2429_v23 = _nw333;
      let _2430_v24;
      let _init52 = function (_2431_i5) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("htylxj"), (_this).f2);
      };
      let _nw334 = Array((new BigNumber(5)).toNumber());
      for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw334.length); _i0_52++) {
        _nw334[_i0_52] = _init52(new BigNumber(_i0_52));
      }
      _2430_v24 = _nw334;
      let _rhs403 = (_2429_v23).fm13(globalState);
      let _rhs404 = _2429_v23;
      let _rhs405 = _this.f3;
      let _rhs406 = _2430_v24;
      let _lhs268 = _this;
      r0 = _rhs403;
      _2429_v23 = _rhs404;
      _lhs268.f3 = _rhs405;
      _2430_v24 = _rhs406;
      let _2432_v25;
      let _nw335 = new _module.C5();
      _nw335.__ctor(_this.f3, ((_this).f1).Merge((_this).f1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), function (_2433_i6) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }));
      _2432_v25 = _nw335;
      let _2434_v26;
      _2434_v26 = new _dafny.CodePoint('p'.codePointAt(0));
      _2434_v26 = _2434_v26;
      let _2435_v27;
      _2435_v27 = _dafny.Seq.of((_2432_v25).f15);
      let _2436_v28;
      _2436_v28 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f3),(_2435_v27)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_2435_v27).length))]);
      let _2437_v29;
      _2437_v29 = _dafny.Seq.of(new BigNumber(838), (_this).fm12(new BigNumber((_2436_v28).length), _2429_v23.f12, _2429_v23.f12, (_2432_v25).f15, globalState), _2429_v23.f12, _2429_v23.f12, new BigNumber(867));
      (_2429_v23).f12 = (_dafny.ZERO).minus(((_2437_v29)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_2437_v29).length))]).plus(_module.__default.safeModuloInt((_this).f4, (_this).f4)));
      let _2438_v30;
      _2438_v30 = _module.D4.create_DC13((_2432_v25).f15, _2437_v29);
      let _pat_let_tv53 = _2432_v25;
      r0 = (function (_pat_let68_0) {
        return function (_2439_dt__update__tmp_h0) {
          return function (_pat_let69_0) {
            return function (_2440_dt__update_hcf21_h0) {
              return _module.D4.create_DC13(_2440_dt__update_hcf21_h0, (_2439_dt__update__tmp_h0).dtor_cf22);
            }(_pat_let69_0);
          }((_pat_let_tv53).f15);
        }(_pat_let68_0);
      }(_2438_v30)).dtor_cf21;
      r1 = true;
      return [r0, r1];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
      this._f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f7, f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return _this.f3;
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat((_this).f2, (_this).f2);
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f4;
    };
    fm11(p0, globalState) {
      let _this = this;
      return ((_this).f4).isLessThan(new BigNumber(-530));
    };
    fm15(p0, p1, globalState) {
      let _this = this;
      return ((_this).f7).minus((_this).f7);
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _2441_v0;
      _2441_v0 = _dafny.Seq.of((_this).fm11(_this.f3, globalState));
      let _2442_v1;
      _2442_v1 = _module.D3.create_DC10(_module.__default.fm0(globalState), p0);
      let _2443_v2;
      _2443_v2 = _dafny.MultiSet.fromElements(_this.f3);
      let _rhs407 = _dafny.Seq.update(((_this.f3) ? (_dafny.Seq.Concat(_dafny.Seq.of(_this.f3, _this.f3, false), _dafny.Seq.of(_this.f3, true, _this.f3))) : (_dafny.Seq.of(_this.f3))), _module.__default.safeIndex(new BigNumber(420), new BigNumber((((_this.f3) ? (_dafny.Seq.Concat(_dafny.Seq.of(_this.f3, _this.f3, false), _dafny.Seq.of(_this.f3, true, _this.f3))) : (_dafny.Seq.of(_this.f3)))).length)), (_2443_v2).IsSubsetOf(_2443_v2));
      let _rhs408 = _2442_v1;
      let _rhs409 = _module.__default.safeDivisionInt(new BigNumber(((_this).f2).length), (_this).f7);
      _2441_v0 = _rhs407;
      _2442_v1 = _rhs408;
      r1 = _rhs409;
      let _2444_v3;
      _2444_v3 = _dafny.Set.fromElements(_this.f3);
      let _2445_v4;
      _2445_v4 = _dafny.Seq.of((_this).f4);
      let _2446_v5;
      _2446_v5 = _dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(552), new BigNumber((_2444_v3).length)), _2445_v4);
      if ((_this.f3) === ((_2446_v5).equals(_dafny.MultiSet.fromElements(_2445_v4, _2445_v4, _2445_v4)))) {
        let _2447_v6;
        _2447_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm11(_this.f3, globalState),_2441_v0);
        _2447_v6 = _2447_v6;
        let _2448_v7;
        _2448_v7 = _module.D3.create_DC9(p1, _this.f3);
        let _2449_v8;
        _2449_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_2448_v7);
        let _2450_v9;
        _2450_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f4);
        r2 = (_this).fm12(_module.__default.safeDivisionInt((_this).f7, p1), (_this).fm15(_dafny.Map.Empty.slice().updateUnsafe(_2449_v8,p0), _2450_v9, globalState), p1, ((_this.f3) ? (_this.f3) : (_this.f3)), globalState);
        let _2451_v10;
        _2451_v10 = _dafny.MultiSet.fromElements((_this).f5);
        let _2452_v11;
        let _nw336 = Array((new BigNumber(4)).toNumber());
        _nw336[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f2).length));
        _nw336[(_dafny.ONE).toNumber()] = (_this).f4;
        _nw336[(new BigNumber(2)).toNumber()] = new BigNumber((_2451_v10).cardinality());
        _nw336[(new BigNumber(3)).toNumber()] = p0;
        _2452_v11 = _nw336;
        _2452_v11 = _2452_v11;
        let _2453_v12;
        let _nw337 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
        _2453_v12 = _nw337;
        let _2454_v13;
        _2454_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,false);
        let _2455_v14;
        _2455_v14 = _module.D0.create_DC0(new BigNumber(((_2454_v13).update(new BigNumber((_dafny.Seq.UnicodeFromString("j")).length), _this.f3)).length));
        let _2456_v15;
        _2456_v15 = _dafny.Seq.of(_2455_v14);
        let _index383 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2453_v12).length));
        (_2453_v12)[_index383] = _2456_v15;
        let _index384 = _module.__default.safeIndex(new BigNumber(145), new BigNumber((_2453_v12).length));
        (_2453_v12)[_index384] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(702)), ((_2457_v2, _2458_p1) => function (_2459_i0) {
          return _module.D0.create_DC0(new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber((_2457_v2).cardinality()), new BigNumber(((_this).f2).length)), (_module.D2.create_DC6(new _dafny.CodePoint('r'.codePointAt(0)), _this.f3, true, _2458_p1)).dtor_cf9)).length));
        })(_2443_v2, p1));
        let _2460_v16;
        _2460_v16 = _dafny.Seq.UnicodeFromString("hccwe");
        let _2461_v17;
        _2461_v17 = new _dafny.CodePoint('j'.codePointAt(0));
        let _2462_v18;
        _2462_v18 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("eqaqunld"), _module.__default.safeIndex(new BigNumber(846), new BigNumber((_dafny.Seq.UnicodeFromString("eqaqunld")).length)), _2461_v17), _2460_v16, _2460_v16);
        let _2463_v19;
        _2463_v19 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f7);
        let _2464_v20;
        _2464_v20 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(p0, (((_2463_v19).contains(_this.f3)) ? ((_2463_v19).get(_this.f3)) : (p0))),p0);
        let _2465_v21;
        _2465_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2461_v17,new BigNumber(-156));
        let _2466_v22;
        _2466_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2465_v21,(_this).f4);
        let _rhs410 = _dafny.Seq.update((_2462_v18)[_module.__default.safeIndex((((_2464_v20).contains(_2442_v1)) ? ((_2464_v20).get(_2442_v1)) : (new BigNumber(919))), new BigNumber((_2462_v18).length))], _module.__default.safeIndex((_this).f7, new BigNumber(((_2462_v18)[_module.__default.safeIndex((((_2464_v20).contains(_2442_v1)) ? ((_2464_v20).get(_2442_v1)) : (new BigNumber(919))), new BigNumber((_2462_v18).length))]).length)), _2461_v17);
        let _rhs411 = ((((_2466_v22).contains(_dafny.Map.Empty.slice().updateUnsafe(_2461_v17,(_this).f7))) ? ((_2466_v22).get(_dafny.Map.Empty.slice().updateUnsafe(_2461_v17,(_this).f7))) : ((_this).f7))).multipliedBy((_this).f7);
        _2460_v16 = _rhs410;
        r1 = _rhs411;
      } else {
        let _2467_v23;
        let _nw338 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _2467_v23 = _nw338;
        let _2468_v24;
        _2468_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-264),_2467_v23);
        let _2469_v25;
        _2469_v25 = _dafny.MultiSet.fromElements(new BigNumber((_2468_v24).length));
        let _2470_v26;
        _2470_v26 = new _dafny.CodePoint('f'.codePointAt(0));
        let _2471_v27;
        _2471_v27 = _dafny.Set.fromElements(_2470_v26, _2470_v26, _2470_v26);
        r2 = ((_this).f7).multipliedBy((_this).fm12(new BigNumber((_2469_v25).cardinality()), new BigNumber((_2471_v27).length), p0, _this.f3, globalState));
        (_this).f3 = _this.f3;
        let _2472_v28;
        let _nw339 = new _module.C7();
        _nw339.__ctor(p1, (_this).f5, _this.f3, (_this).f1, (_this).f2);
        _2472_v28 = _nw339;
        let _2473_v29;
        _2473_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
        let _2474_v30;
        let _2475_v31;
        let _2476_v32;
        let _2477_v33;
        let _out62;
        let _out63;
        let _out64;
        let _out65;
        let _outcollector16 = _module.__default.m0(_2473_v29, globalState);
        _out62 = _outcollector16[0];
        _out63 = _outcollector16[1];
        _out64 = _outcollector16[2];
        _out65 = _outcollector16[3];
        _2474_v30 = _out62;
        _2475_v31 = _out63;
        _2476_v32 = _out64;
        _2477_v33 = _out65;
        let _2478_v34;
        _2478_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v23,(_this).f5);
        let _2479_v35;
        _2479_v35 = _module.D0.create_DC1(new BigNumber(696));
        let _2480_v37;
        _2480_v37 = _dafny.Seq.of((_this).f2);
        let _2481_v38;
        let _nw340 = new _module.C9();
        _nw340.__ctor((_this).f5, _2478_v34, (_2479_v35).dtor_cf1, (_this).f5, _2475_v31, (function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of (_2480_v37).Elements) {
            let _2482_v36 = _compr_66;
            if (_dafny.Seq.contains(_2480_v37, _2482_v36)) {
              _coll66.push([_2482_v36,_2475_v31]);
            }
          }
          return _coll66;
        }()).update((_this).f2, _this.f3), (_this).f2);
        _2481_v38 = _nw340;
      }
      let _2483_v39;
      _2483_v39 = _dafny.Seq.of(_2441_v0);
      let _2484_v40;
      _2484_v40 = _2441_v0;
      let _2485_v41;
      let _nw341 = Array((new BigNumber(9)).toNumber());
      _nw341[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_this.f3);
      _nw341[(_dafny.ONE).toNumber()] = _2441_v0;
      _nw341[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_this.f3, _this.f3);
      _nw341[(new BigNumber(3)).toNumber()] = _2441_v0;
      _nw341[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_this.f3);
      _nw341[(new BigNumber(5)).toNumber()] = _2441_v0;
      _nw341[(new BigNumber(6)).toNumber()] = (_2483_v39)[_module.__default.safeIndex((_this).f7, new BigNumber((_2483_v39).length))];
      _nw341[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2441_v0, _2441_v0);
      _nw341[(new BigNumber(8)).toNumber()] = (_2484_v40);
      _2485_v41 = _nw341;
      let _index385 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length));
      (_2485_v41)[_index385] = _2441_v0;
      let _index386 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length));
      (_2485_v41)[_index386] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f3), _2441_v0);
      let _2486_v42;
      _2486_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f1);
      let _2487_v43;
      _2487_v43 = _dafny.Set.fromElements((_this).f7);
      let _2488_v45;
      _2488_v45 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("gpqhqgii"));
      let _2489_v46;
      let _nw342 = new _module.C6();
      _nw342.__ctor(p1, _module.__default.safeDivisionInt(new BigNumber(-575), (_this).f7), p0, (_this).f5, _this.f3, (((_2486_v42).contains(new BigNumber((_2487_v43).length))) ? ((_2486_v42).get(new BigNumber((_2487_v43).length))) : (function () {
        let _coll67 = new _dafny.Map();
        for (const _compr_67 of (_2488_v45).Elements) {
          let _2490_v44 = _compr_67;
          if (_dafny.Seq.contains(_2488_v45, _2490_v44)) {
            _coll67.push([_2490_v44,_this.f3]);
          }
        }
        return _coll67;
      }())), (_this).f2);
      _2489_v46 = _nw342;
      if (false) {
        let _2491_v47;
        _2491_v47 = _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_2489_v46.f12,_this.f3));
        let _2492_v48;
        _2492_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(157),_this.f3);
        let _2493_v50;
        _2493_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber((_2443_v2).cardinality()));
        let _pat_let_tv54 = _2492_v48;
        let _pat_let_tv55 = _2492_v48;
        let _2494_v51;
        let _nw343 = Array((new BigNumber(19)).toNumber());
        _nw343[(_dafny.ZERO).toNumber()] = _2491_v47;
        _nw343[(_dafny.ONE).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(2)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(3)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(4)).toNumber()] = function (_pat_let70_0) {
          return function (_2495_dt__update__tmp_h0) {
            return function (_pat_let71_0) {
              return function (_2496_dt__update_hcf27_h0) {
                return _module.D6.create_DC16(_2496_dt__update_hcf27_h0);
              }(_pat_let71_0);
            }(_pat_let_tv54);
          }(_pat_let70_0);
        }(_module.D6.create_DC16((_2492_v48).update(new BigNumber(440), _this.f3)));
        _nw343[(new BigNumber(5)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(6)).toNumber()] = _module.D6.create_DC16(function () {
  let _coll68 = new _dafny.Map();
  for (const _compr_68 of _dafny.IntegerRange(new BigNumber(828), new BigNumber(777))) {
    let _2497_v49 = _compr_68;
    if (((new BigNumber(828)).isLessThanOrEqualTo(_2497_v49)) && ((_2497_v49).isLessThan(new BigNumber(777)))) {
      _coll68.push([_module.__default.safeDivisionInt(_2497_v49, _2489_v46.f12),_this.f3]);
    }
  }
  return _coll68;
}());
        _nw343[(new BigNumber(7)).toNumber()] = function (_pat_let72_0) {
          return function (_2498_dt__update__tmp_h1) {
            return function (_pat_let73_0) {
              return function (_2499_dt__update_hcf27_h1) {
                return _module.D6.create_DC16(_2499_dt__update_hcf27_h1);
              }(_pat_let73_0);
            }(_pat_let_tv55);
          }(_pat_let72_0);
        }(_module.D6.create_DC16(_2492_v48));
        _nw343[(new BigNumber(8)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(9)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(10)).toNumber()] = _module.D6.create_DC16(_2492_v48);
        _nw343[(new BigNumber(11)).toNumber()] = _module.__default.fm54(_2493_v50, true, globalState);
        _nw343[(new BigNumber(12)).toNumber()] = _module.D6.create_DC16(_dafny.Map.Empty.slice().updateUnsafe((_2489_v46).f11,_this.f3));
        _nw343[(new BigNumber(13)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(14)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(15)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(16)).toNumber()] = _2491_v47;
        _nw343[(new BigNumber(17)).toNumber()] = _module.D6.create_DC16(_2492_v48);
        _nw343[(new BigNumber(18)).toNumber()] = _2491_v47;
        _2494_v51 = _nw343;
        let _2500_v52;
        _2500_v52 = _dafny.Seq.of(_2494_v51);
        let _2501_v53;
        _2501_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2445_v4);
        let _2502_v55;
        let _nw344 = Array((new BigNumber(19)).toNumber());
        _nw344[(_dafny.ZERO).toNumber()] = (_this).f4;
        _nw344[(_dafny.ONE).toNumber()] = p1;
        _nw344[(new BigNumber(2)).toNumber()] = (p1).multipliedBy(new BigNumber(73));
        _nw344[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_2500_v52, _2500_v52)).length);
        _nw344[(new BigNumber(4)).toNumber()] = new BigNumber(((((_2501_v53).contains(_this.f3)) ? ((_2501_v53).get(_this.f3)) : (_2445_v4))).length);
        _nw344[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), ((_2503_v0) => function (_2504_i1) {
          return _2503_v0;
        })(_2441_v0))).length));
        _nw344[(new BigNumber(6)).toNumber()] = (_module.D9.create_DC26(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_2489_v46).f11)).update(_this.f3, p0)).length))).dtor_cf44;
        _nw344[(new BigNumber(7)).toNumber()] = (_2489_v46).fm21(_this.f3, (_2445_v4)[_module.__default.safeIndex((_this).f7, new BigNumber((_2445_v4).length))], globalState);
        _nw344[(new BigNumber(8)).toNumber()] = (_this).f7;
        _nw344[(new BigNumber(9)).toNumber()] = (_2489_v46).f11;
        _nw344[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt((_2489_v46).f11, (_dafny.ZERO).minus((_2489_v46).f11));
        _nw344[(new BigNumber(11)).toNumber()] = (_2489_v46).f11;
        _nw344[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt((_2489_v46).f11, (_2489_v46).f11);
        _nw344[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_2493_v50).update(_this.f3, new BigNumber(807))).length)));
        _nw344[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_2489_v46.f12);
        _nw344[(new BigNumber(15)).toNumber()] = _2489_v46.f12;
        _nw344[(new BigNumber(16)).toNumber()] = new BigNumber(-920);
        _nw344[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt((_this).f7, new BigNumber((_2441_v0).length));
        _nw344[(new BigNumber(18)).toNumber()] = new BigNumber((function () {
          let _coll69 = new _dafny.Map();
          for (const _compr_69 of _dafny.IntegerRange(new BigNumber(-611), new BigNumber(128))) {
            let _2505_v54 = _compr_69;
            if (((new BigNumber(-611)).isLessThanOrEqualTo(_2505_v54)) && ((_2505_v54).isLessThan(new BigNumber(128)))) {
              _coll69.push([(_2505_v54).plus((_this).f7),_2443_v2]);
            }
          }
          return _coll69;
        }()).length);
        _2502_v55 = _nw344;
        let _index387 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_2502_v55).length));
        (_2502_v55)[_index387] = p0;
        let _index388 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_2502_v55).length));
        (_2502_v55)[_index388] = (_this).fm12(_2489_v46.f12, (_2489_v46).f11, _module.__default.fm0(globalState), !(_this.f3) || (_this.f3), globalState);
        if (!(false)) {
          r2 = _module.__default.safeDivisionInt((_2489_v46).f11, _2489_v46.f12);
          (_this).f3 = (false) || (!(_this.f3));
          let _2506_v56;
          let _nw345 = new _module.C8();
          _nw345.__ctor(_2483_v39, _this.f3, (_this).f1, _dafny.Seq.UnicodeFromString("e"), new BigNumber((_2493_v50).length), (_this).f5);
          _2506_v56 = _nw345;
          let _2507_v57;
          _2507_v57 = _dafny.Seq.of(_2506_v56, _2506_v56, _2506_v56);
          _2507_v57 = _2507_v57;
          let _2508_v58;
          _2508_v58 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_2489_v46.f12)),(_2489_v46).f11);
          _2508_v58 = _2508_v58;
          let _index389 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length));
          let _rhs412 = (_2502_v55)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_2502_v55).length))];
          let _rhs413 = _dafny.Seq.Concat((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))], _dafny.Seq.update(_module.__default.fm7(_this.f3, !(_this.f3), globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm7(_this.f3, !(_this.f3), globalState)).length)), !(_this.f3)));
          let _rhs414 = _2441_v0;
          let _lhs269 = _2485_v41;
          let _lhs270 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length));
          r2 = _rhs412;
          _lhs269[_lhs270] = _rhs413;
          _2441_v0 = _rhs414;
        } else {
          let _2509_v59;
          _2509_v59 = new _dafny.CodePoint('q'.codePointAt(0));
          let _2510_v60;
          let _nw346 = new _module.C3();
          _nw346.__ctor(new BigNumber(((_this).f2).length), (_this).f5, _this.f3, ((_this).f1).update((_this).f2, true), (_2489_v46).fm14(_this.f3, _this.f3, _2509_v59, (_this).f7, globalState));
          _2510_v60 = _nw346;
          (_2489_v46).f12 = (_2445_v4)[_module.__default.safeIndex(p0, new BigNumber((_2445_v4).length))];
          let _2511_v61;
          _2511_v61 = _dafny.Seq.UnicodeFromString("mah");
          let _rhs415 = (new BigNumber(((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))]).length)).minus((_2489_v46).f11);
          let _rhs416 = (_this).f2;
          let _rhs417 = p1;
          r1 = _rhs415;
          _2511_v61 = _rhs416;
          r1 = _rhs417;
          let _2512_v62;
          _2512_v62 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
          let _2513_v63;
          _2513_v63 = _module.D1.create_DC4((((_2512_v62).contains(_this.f3)) ? ((_2512_v62).get(_this.f3)) : (_this.f3)), !(_this.f3));
          let _2514_v64;
          _2514_v64 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2513_v63);
          _2514_v64 = ((_2514_v64).Merge(_2514_v64)).Merge(_2514_v64);
          (_this).f3 = ((_this).f7).isLessThan((_dafny.ZERO).minus(((_this.f3) ? (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), ((_2515_v59) => function (_2516_i2) {
            return _2515_v59;
          })(_2509_v59))).length)) : ((_2502_v55)[_module.__default.safeIndex(new BigNumber(614), new BigNumber((_2502_v55).length))]))));
        }
        let _2517_v65;
        _2517_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_module.__default.fm7(_this.f3, _this.f3, globalState));
        (_2489_v46).f12 = new BigNumber(((_2517_v65).update(true, (_2483_v39)[_module.__default.safeIndex(new BigNumber(-614), new BigNumber((_2483_v39).length))])).length);
        (_this).f3 = _this.f3;
        _2485_v41 = _2485_v41;
      } else {
        _2441_v0 = (_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))];
        (_this).f3 = _this.f3;
        if (_this.f3) {
          let _2518_v66;
          _2518_v66 = _dafny.Seq.UnicodeFromString("hwuuwcna");
          _2518_v66 = _2518_v66;
          let _2519_v67;
          _2519_v67 = new _dafny.CodePoint('a'.codePointAt(0));
          (_2489_v46).f12 = new BigNumber(((_2489_v46).fm14((_2489_v46.f12).isLessThan(new BigNumber((_module.__default.fm20(_this.f3, !(_this.f3), _this.f3, _dafny.Seq.update(_2441_v0, _module.__default.safeIndex((_2489_v46).f11, new BigNumber((_2441_v0).length)), true), globalState)).length)), _this.f3, _2519_v67, (_this).f4, globalState)).length);
          _2518_v66 = _2518_v66;
          _2519_v67 = _2519_v67;
          let _2520_v68;
          _2520_v68 = _2487_v43;
          let _2521_v69;
          _2521_v69 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2520_v68);
          let _2522_v70;
          _2522_v70 = _dafny.MultiSet.fromElements(new BigNumber((_2487_v43).length));
          let _2523_v71;
          let _nw347 = new _module.C2();
          _nw347.__ctor((_dafny.ZERO).minus((_2489_v46).f11), _this.f3, _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3), _2518_v66);
          _2523_v71 = _nw347;
          let _2524_v72;
          _2524_v72 = _module.D13.create_DC33(_2523_v71);
          let _2525_v73;
          _2525_v73 = _dafny.Seq.of((_this).f5, (_this).f5);
          let _2526_v74;
          let _init53 = ((_2527_v46) => function (_2528_i3) {
            return _module.__default.safeDivisionInt(_2528_i3, (_2527_v46).f11);
          })(_2489_v46);
          let _nw348 = Array((new BigNumber(29)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw348.length); _i0_53++) {
            _nw348[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _2526_v74 = _nw348;
          let _2529_v75;
          _2529_v75 = _dafny.Map.Empty.slice().updateUnsafe(_2526_v74,(_this).f5);
          let _2530_v76;
          let _nw349 = new _module.C9();
          _nw349.__ctor((_2525_v73)[_module.__default.safeIndex(_2489_v46.f12, new BigNumber((_2525_v73).length))], _2529_v75, (_2445_v4)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_2445_v4).length))], (_this).f5, _this.f3, (_this).f1, _dafny.Seq.UnicodeFromString("dlgym"));
          _2530_v76 = _nw349;
          let _2531_v77;
          _2531_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2524_v72,_2530_v76);
          let _2532_v78;
          let _nw350 = Array((new BigNumber(17)).toNumber());
          _nw350[(_dafny.ZERO).toNumber()] = (_this).f4;
          _nw350[(_dafny.ONE).toNumber()] = new BigNumber(((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))]).length);
          _nw350[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_2489_v46.f12);
          _nw350[(new BigNumber(3)).toNumber()] = p0;
          _nw350[(new BigNumber(4)).toNumber()] = new BigNumber((_2518_v66).length);
          _nw350[(new BigNumber(5)).toNumber()] = (((_2522_v70).contains((_2489_v46).f11)) ? ((_2522_v70).get((_2489_v46).f11)) : (new BigNumber(832)));
          _nw350[(new BigNumber(6)).toNumber()] = new BigNumber(922);
          _nw350[(new BigNumber(7)).toNumber()] = (_2489_v46).f11;
          _nw350[(new BigNumber(8)).toNumber()] = new BigNumber(844);
          _nw350[(new BigNumber(9)).toNumber()] = new BigNumber(-112);
          _nw350[(new BigNumber(10)).toNumber()] = p0;
          _nw350[(new BigNumber(11)).toNumber()] = _2489_v46.f12;
          _nw350[(new BigNumber(12)).toNumber()] = new BigNumber((_2518_v66).length);
          _nw350[(new BigNumber(13)).toNumber()] = new BigNumber((_2531_v77).length);
          _nw350[(new BigNumber(14)).toNumber()] = _2489_v46.f12;
          _nw350[(new BigNumber(15)).toNumber()] = (_2489_v46).f11;
          _nw350[(new BigNumber(16)).toNumber()] = new BigNumber(((_this).f2).length);
          _2532_v78 = _nw350;
          let _2533_v79;
          _2533_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2445_v4,_2526_v74);
          let _2534_v80;
          let _nw351 = Array((new BigNumber(27)).toNumber());
          _nw351[(_dafny.ZERO).toNumber()] = _this.f3;
          _nw351[(_dafny.ONE).toNumber()] = (_2487_v43).IsDisjointFrom(_2487_v43);
          _nw351[(new BigNumber(2)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(3)).toNumber()] = false;
          _nw351[(new BigNumber(4)).toNumber()] = !((_2443_v2).IsSubsetOf(_2443_v2));
          _nw351[(new BigNumber(5)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(6)).toNumber()] = false;
          _nw351[(new BigNumber(7)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(8)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(9)).toNumber()] = (new BigNumber((_module.__default.fm6(_this.f3, _2519_v67, globalState)).length)).isLessThanOrEqualTo(p1);
          _nw351[(new BigNumber(10)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(11)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(12)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2445_v4, _2445_v4);
          _nw351[(new BigNumber(14)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(15)).toNumber()] = false;
          _nw351[(new BigNumber(16)).toNumber()] = (p1).isLessThanOrEqualTo(new BigNumber((_2521_v69).length));
          _nw351[(new BigNumber(17)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(18)).toNumber()] = (_2533_v79).contains(_dafny.Seq.of(_2489_v46.f12));
          _nw351[(new BigNumber(19)).toNumber()] = (_2522_v70).IsSubsetOf(_2522_v70);
          _nw351[(new BigNumber(20)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(21)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(22)).toNumber()] = (_2523_v71).fm11(_this.f3, globalState);
          _nw351[(new BigNumber(23)).toNumber()] = !_dafny.areEqual((_this).f2, (_this).f2);
          _nw351[(new BigNumber(24)).toNumber()] = _this.f3;
          _nw351[(new BigNumber(25)).toNumber()] = (_2489_v46.f12).isLessThanOrEqualTo(_2489_v46.f12);
          _nw351[(new BigNumber(26)).toNumber()] = _this.f3;
          _2534_v80 = _nw351;
          _2534_v80 = (_2530_v76).f8;
        } else {
          let _2535_v81;
          let _nw352 = new _module.C1();
          _nw352.__ctor();
          _2535_v81 = _nw352;
          _2535_v81 = _2535_v81;
          let _index390 = _module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index390] = (_2441_v0)[_module.__default.safeIndex((_2489_v46).f11, new BigNumber((_2441_v0).length))];
          let _index391 = _module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index391] = (_2489_v46).fm11(_this.f3, globalState);
          r2 = _2489_v46.f12;
          let _2536_v82;
          _2536_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f7);
          let _2537_v83;
          _2537_v83 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length))],((_this).f5)[_module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length))]);
          let _2538_v84;
          _2538_v84 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((_this).f2).length));
          let _2539_v85;
          _2539_v85 = _dafny.Map.Empty.slice().updateUnsafe((_2489_v46).f11,_this.f3);
          let _2540_v86;
          _2540_v86 = _dafny.MultiSet.fromElements((_2489_v46).f11, new BigNumber((_2539_v85).length), (_2489_v46).f11, (_2489_v46).f11, _2489_v46.f12);
          let _2541_v87;
          _2541_v87 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length))],_2540_v86);
          let _2542_v88;
          let _nw353 = Array((new BigNumber(17)).toNumber());
          _nw353[(_dafny.ZERO).toNumber()] = new BigNumber(((_module.__default.fm33(new BigNumber((_dafny.MultiSet.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length))])).cardinality()), globalState))).length);
          _nw353[(_dafny.ONE).toNumber()] = (_2489_v46).f11;
          _nw353[(new BigNumber(2)).toNumber()] = _2489_v46.f12;
          _nw353[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw353[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((((_2536_v82).contains((_this).f4)) ? ((_2536_v82).get((_this).f4)) : (new BigNumber((_2444_v3).length))));
          _nw353[(new BigNumber(5)).toNumber()] = p1;
          _nw353[(new BigNumber(6)).toNumber()] = p0;
          _nw353[(new BigNumber(7)).toNumber()] = new BigNumber(((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))]).length);
          _nw353[(new BigNumber(8)).toNumber()] = new BigNumber(484);
          _nw353[(new BigNumber(9)).toNumber()] = new BigNumber(34);
          _nw353[(new BigNumber(10)).toNumber()] = (((_2443_v2).contains(false)) ? ((_2443_v2).get(false)) : (new BigNumber((_2537_v83).length)));
          _nw353[(new BigNumber(11)).toNumber()] = _2489_v46.f12;
          _nw353[(new BigNumber(12)).toNumber()] = new BigNumber((_2538_v84).length);
          _nw353[(new BigNumber(13)).toNumber()] = _2489_v46.f12;
          _nw353[(new BigNumber(14)).toNumber()] = new BigNumber(291);
          _nw353[(new BigNumber(15)).toNumber()] = _2489_v46.f12;
          _nw353[(new BigNumber(16)).toNumber()] = new BigNumber((_2541_v87).length);
          _2542_v88 = _nw353;
          let _2543_v89;
          _2543_v89 = new _dafny.CodePoint('g'.codePointAt(0));
          _2483_v39 = _module.__default.fm58(_module.__default.fm31(_module.__default.fm0(globalState), ((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))])[_module.__default.safeIndex((_this).f7, new BigNumber(((_2485_v41)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_2485_v41).length))]).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2542_v88,new BigNumber((_2536_v82).length))).length), _2543_v89, globalState), _2543_v89, _2538_v84, globalState);
          let _index392 = _module.__default.safeIndex(new BigNumber(726), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index392] = ((_2489_v46).f11).isLessThan((_dafny.ZERO).minus(p0));
        }
        let _2544_v90;
        _2544_v90 = new _dafny.CodePoint('w'.codePointAt(0));
        let _2545_v91;
        _2545_v91 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex((_this).fm12((_2489_v46).f11, _2489_v46.f12, p0, _this.f3, globalState), new BigNumber(((_this).f2).length)), _2544_v90)).length));
        _2545_v91 = _2545_v91;
        let _2546_v92;
        _2546_v92 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3);
        (_this).f3 = (_this.f3) || ((_2546_v92).equals(_2546_v92));
      }
      if (true) {
        r1 = _2489_v46.f12;
        let _2547_v93;
        let _nw354 = new _module.C1();
        _nw354.__ctor();
        _2547_v93 = _nw354;
        let _2548_v94;
        _2548_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2547_v93,false);
        let _2549_v95;
        _2549_v95 = _dafny.Seq.of(_2548_v94);
        let _2550_v96;
        let _nw355 = new _module.C2();
        _nw355.__ctor(_module.__default.safeModuloInt(new BigNumber(-309), p0), true, (_dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3)).Merge((_this).f1), (_this).f2);
        _2550_v96 = _nw355;
        let _rhs418 = _2550_v96.f3;
        let _rhs419 = (_2489_v46.f12).isLessThanOrEqualTo((_this).f4);
        let _rhs420 = _2549_v95;
        let _rhs421 = _2550_v96;
        let _lhs271 = _this;
        let _lhs272 = _this;
        _lhs271.f3 = _rhs418;
        _lhs272.f3 = _rhs419;
        _2549_v95 = _rhs420;
        _2550_v96 = _rhs421;
        let _2551_v97;
        let _init54 = ((_2552_p1) => function (_2553_i4) {
          return (_2553_i4).multipliedBy(_2552_p1);
        })(p1);
        let _nw356 = Array((new BigNumber(8)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw356.length); _i0_54++) {
          _nw356[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _2551_v97 = _nw356;
        let _2554_v98;
        _2554_v98 = _dafny.Seq.UnicodeFromString("hbi");
        let _index393 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length));
        (_2551_v97)[_index393] = new BigNumber(-246);
        let _index394 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length));
        let _rhs422 = _2551_v97;
        let _rhs423 = _dafny.Seq.Concat(_2554_v98, ((_2550_v96.f3) ? (_2554_v98) : ((_this).f2)));
        let _rhs424 = _2489_v46.f12;
        let _rhs425 = _this.f3;
        let _lhs273 = _2551_v97;
        let _lhs274 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length));
        let _lhs275 = _this;
        _2551_v97 = _rhs422;
        _2554_v98 = _rhs423;
        _lhs273[_lhs274] = _rhs424;
        _lhs275.f3 = _rhs425;
        r1 = (_this).fm12(_module.__default.safeModuloInt(new BigNumber(375), (_this).f4), _module.__default.safeDivisionInt((_this).f7, _2489_v46.f12), (new BigNumber(448)).minus((_2551_v97)[_module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length))]), _2550_v96.f3, globalState);
        if (!(!(_2550_v96.f3))) {
          (_2489_v46).f12 = p0;
          let _2555_v99;
          _2555_v99 = new _dafny.CodePoint('i'.codePointAt(0));
          let _2556_v100;
          let _nw357 = new _module.C3();
          _nw357.__ctor(p1, (_this).f5, _2550_v96.f3, (((_2550_v96).f1).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), function (_2557_i5) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }), _2550_v96.f3)).Merge((_2550_v96).f1), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), function (_2558_i6) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), _module.__default.safeIndex((_2489_v46).f11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(686)), function (_2559_i6) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length)), _2555_v99));
          _2556_v100 = _nw357;
          let _index395 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length));
          (_2551_v97)[_index395] = new BigNumber((((_this.f3) ? (_2554_v98) : ((_2550_v96).f2))).length);
          let _2560_v101;
          _2560_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-199),p0);
          let _2561_v102;
          _2561_v102 = _dafny.Map.Empty.slice().updateUnsafe((_2489_v46).f11,_2560_v101);
          let _2562_v103;
          _2562_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f5);
          _2561_v102 = (_2561_v102).update(new BigNumber((_2562_v103).length), _2560_v101);
          let _index396 = _module.__default.safeIndex(new BigNumber(812), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index396] = _module.__default.fm3((_2489_v46).f11, globalState);
          let _index397 = _module.__default.safeIndex(new BigNumber(812), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index397] = ((_2550_v96.f3) ? (_this.f3) : (true));
        } else {
          (_2550_v96).f3 = !(_this.f3);
          let _2563_v104;
          let _nw358 = new _module.C2();
          _nw358.__ctor((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(303)), function (_2564_i7) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })).length)), _2550_v96.f3, (_this).f1, _2554_v98);
          _2563_v104 = _nw358;
          let _2565_v105;
          _2565_v105 = _module.D13.create_DC36(_module.D13.create_DC33(_2563_v104));
          _2565_v105 = _2565_v105;
          let _2566_v106;
          _2566_v106 = new _dafny.CodePoint('i'.codePointAt(0));
          _2566_v106 = _2566_v106;
          let _index398 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_2551_v97).length));
          (_2551_v97)[_index398] = (p0).multipliedBy(p1);
          let _2567_v107;
          let _nw359 = Array((new BigNumber(24)).toNumber()).fill(false);
          _2567_v107 = _nw359;
          _2567_v107 = (_this).f5;
        }
      } else {
        (_this).f3 = _this.f3;
        let _2568_v108;
        _2568_v108 = new _dafny.CodePoint('p'.codePointAt(0));
        let _2569_v109;
        _2569_v109 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new BigNumber((_module.__default.fm31(_2489_v46.f12, _this.f3, p1, _2568_v108, globalState)).length));
        let _2570_v110;
        _2570_v110 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f2);
        let _2571_v111;
        let _nw360 = Array((new BigNumber(26)).toNumber());
        _nw360[(_dafny.ZERO).toNumber()] = _2569_v109;
        _nw360[(_dafny.ONE).toNumber()] = (_2569_v109).update(_this.f3, p1);
        _nw360[(new BigNumber(2)).toNumber()] = (_2569_v109).update(_this.f3, (_this).fm12((_this).f7, new BigNumber(((((_2570_v110).contains(_this.f3)) ? ((_2570_v110).get(_this.f3)) : ((_this).f2))).length), p0, _this.f3, globalState));
        _nw360[(new BigNumber(3)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(4)).toNumber()] = (_2569_v109).update(_this.f3, (_this).f7);
        _nw360[(new BigNumber(5)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(6)).toNumber()] = (_2569_v109).Merge(_2569_v109);
        _nw360[(new BigNumber(7)).toNumber()] = (_2569_v109).Merge(_2569_v109);
        _nw360[(new BigNumber(8)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(9)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(10)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(11)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(12)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(13)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_2441_v0)[_module.__default.safeIndex((_this).f4, new BigNumber((_2441_v0).length))],(_this).f4);
        _nw360[(new BigNumber(15)).toNumber()] = (_2569_v109).Merge(_2569_v109);
        _nw360[(new BigNumber(16)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(17)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_module.__default.fm0(globalState));
        _nw360[(new BigNumber(19)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(20)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(21)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(false,_2489_v46.f12)).Merge(_2569_v109);
        _nw360[(new BigNumber(22)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(23)).toNumber()] = _2569_v109;
        _nw360[(new BigNumber(24)).toNumber()] = (((_2489_v46).fm13(globalState)) ? (_2569_v109) : (_dafny.Map.Empty.slice().updateUnsafe(false,(_2489_v46).f11)));
        _nw360[(new BigNumber(25)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_dafny.ZERO).minus(new BigNumber(((_this).f2).length)))).Merge(_2569_v109);
        _2571_v111 = _nw360;
        let _index399 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_2571_v111).length));
        (_2571_v111)[_index399] = (_2569_v109).Merge(_2569_v109);
        let _index400 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_2571_v111).length));
        (_2571_v111)[_index400] = _2569_v109;
        r1 = new BigNumber(-593);
        let _2572_v112;
        let _nw361 = new _module.C5();
        _nw361.__ctor(false, (_this).f1, (_this).f2);
        _2572_v112 = _nw361;
        let _2573_v113;
        _2573_v113 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).f2).length),_2572_v112);
        let _2574_v114;
        _2574_v114 = _dafny.Set.fromElements(_2573_v113);
        let _2575_v115;
        _2575_v115 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_this).f4)).length));
        let _2576_v116;
        _2576_v116 = _dafny.Map.Empty.slice().updateUnsafe((_2574_v114).Intersect(_2574_v114),new BigNumber((_2575_v115).cardinality()));
        r1 = (((_2576_v116).contains(_2574_v114)) ? ((_2576_v116).get(_2574_v114)) : (_2489_v46.f12));
        let _2577_v117;
        _2577_v117 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(495),_module.__default.fm59(new BigNumber(346), globalState));
        let _2578_v118;
        let _init55 = ((_2579_v46) => function (_2580_i8) {
          return _module.__default.safeModuloInt(_2580_i8, _2579_v46.f12);
        })(_2489_v46);
        let _nw362 = Array((new BigNumber(24)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw362.length); _i0_55++) {
          _nw362[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2578_v118 = _nw362;
        let _2581_v119;
        _2581_v119 = _dafny.Map.Empty.slice().updateUnsafe(_2578_v118,(_this).f5);
        let _2582_v120;
        let _nw363 = new _module.C9();
        _nw363.__ctor((((_2572_v112).f15) ? ((_this).f5) : ((_this).f5)), (_2581_v119).Merge(_2581_v119), _module.__default.safeModuloInt(new BigNumber((_2487_v43).length), p1), (_this).f5, (_2441_v0)[_module.__default.safeIndex((_2489_v46).f11, new BigNumber((_2441_v0).length))], (_this).f1, _dafny.Seq.Concat((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(276)), ((_2583_v108) => function (_2584_i9) {
          return _2583_v108;
        })(_2568_v108))));
        _2582_v120 = _nw363;
        let _rhs426 = _2577_v117;
        let _rhs427 = (_2489_v46).f11;
        let _rhs428 = (_this.f3) === (_this.f3);
        let _rhs429 = _this.f3;
        let _rhs430 = _2582_v120;
        let _lhs276 = _this;
        let _lhs277 = _this;
        _2577_v117 = _rhs426;
        r2 = _rhs427;
        _lhs276.f3 = _rhs428;
        _lhs277.f3 = _rhs429;
        _2582_v120 = _rhs430;
      }
      let _2585_v121;
      _2585_v121 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3);
      r0 = (((_2585_v121).update(_this.f3, _this.f3)).Merge(_2585_v121)).update(_this.f3, !(false) || (_this.f3));
      let _2586_v122;
      _2586_v122 = _dafny.MultiSet.fromElements(p0);
      r1 = ((!(_this.f3) || (_this.f3)) ? ((((_this).fm11(_this.f3, globalState)) ? (p1) : (new BigNumber((_2586_v122).cardinality())))) : (_module.__default.safeDivisionInt((_this).f4, _2489_v46.f12)));
      let _2587_v123;
      _2587_v123 = _module.D16.create_DC42(_this.f3, _this.f3, _this.f3);
      let _pat_let_tv56 = p1;
      r2 = function (_source43) {
        if (_source43.is_DC42) {
          let _2588___mcc_h0 = (_source43).cf70;
          let _2589___mcc_h1 = (_source43).cf71;
          let _2590___mcc_h2 = (_source43).cf72;
          let _2591_cf72 = _2590___mcc_h2;
          let _2592_cf71 = _2589___mcc_h1;
          let _2593_cf70 = _2588___mcc_h0;
          return new BigNumber(((_this).f2).length);
        } else {
          let _2594___mcc_h3 = (_source43).cf69;
          let _2595_cf69 = _2594___mcc_h3;
          return _pat_let_tv56;
        }
      }(_2587_v123);
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = false;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = [];
      let _2596_v0;
      _2596_v0 = _dafny.MultiSet.fromElements((_this).f7, new BigNumber(((_this).f2).length));
      let _2597_v2;
      _2597_v2 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_this.f3,_this.f3)).update(_this.f3, false)).length)));
      let _2598_i0;
      _2598_i0 = _dafny.ZERO;
      L14: {
        while ((function () {
          let _coll70 = new _dafny.Set();
          for (const _compr_70 of (_2596_v0).Elements) {
            let _2604_v1 = _compr_70;
            if ((_2596_v0).contains(_2604_v1)) {
              _coll70.add((_2604_v1).minus(new BigNumber(811)));
            }
          }
          return _coll70;
        }()).IsProperSubsetOf(_2597_v2)) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2598_i0)) {
              break L14;
            }
            _2598_i0 = (_2598_i0).plus(_dafny.ONE);
            let _2599_v3;
            let _nw364 = new _module.C5();
            _nw364.__ctor(_this.f3, (_this).f1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-749)), function (_2600_i1) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            }));
            _2599_v3 = _nw364;
            let _2601_v4;
            _2601_v4 = _dafny.Seq.of(_this.f3, _this.f3, _this.f3);
            let _2602_v5;
            _2602_v5 = _dafny.Seq.of(_2601_v4);
            _2601_v4 = (_2602_v5)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(751), (_this).f4), new BigNumber((_2602_v5).length))];
            r1 = true;
            let _2603_v6;
            _2603_v6 = _module.D3.create_DC9(new BigNumber(546), !(!(_this.f3)));
            (_this).f3 = (_2603_v6).dtor_cf16;
          }
        }
      }
      let _2605_v7;
      _2605_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,false);
      let _2606_v8;
      _2606_v8 = new BigNumber(154);
      let _2607_v9;
      _2607_v9 = _dafny.MultiSet.fromElements(_this.f3, false);
      let _rhs431 = _2605_v7;
      let _rhs432 = (_this).f4;
      let _rhs433 = _module.__default.safeModuloInt((new BigNumber((_2596_v0).cardinality())).plus((_this).f4), (_this).f7);
      let _rhs434 = ((_2607_v9).Union(_2607_v9)).IsSubsetOf(_dafny.MultiSet.fromElements(_this.f3, _this.f3));
      _2605_v7 = _rhs431;
      _2606_v8 = _rhs432;
      _2606_v8 = _rhs433;
      r1 = _rhs434;
      let _2608_v10;
      _2608_v10 = _module.D8.create_DC23((_this).f7, (_this).f4, false);
      let _2609_v11;
      _2609_v11 = _dafny.Seq.of(_2608_v10);
      if (!_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_2608_v10, _2608_v10), _2609_v11), _2609_v11)) {
        let _2610_v12;
        _2610_v12 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2606_v8);
        let _2611_v13;
        let _init56 = function (_2612_i2) {
          return _module.__default.safeDivisionInt(_2612_i2, (_this).f4);
        };
        let _nw365 = Array((new BigNumber(29)).toNumber());
        for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw365.length); _i0_56++) {
          _nw365[_i0_56] = _init56(new BigNumber(_i0_56));
        }
        _2611_v13 = _nw365;
        let _2613_v14;
        _2613_v14 = _module.D5.create_DC14(_2611_v13);
        let _2614_v15;
        _2614_v15 = _dafny.Seq.of(_2613_v14, _2613_v14);
        if (!((((_2610_v12).contains(_this.f3)) ? ((_2610_v12).get(_this.f3)) : ((_this).f7))).isEqualTo(new BigNumber((_2614_v15).length))) {
          let _2615_v16;
          _2615_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_module.__default.fm9(globalState));
          let _rhs435 = _2615_v16;
          let _rhs436 = _2614_v15;
          _2615_v16 = _rhs435;
          _2614_v15 = _rhs436;
          r1 = _this.f3;
          let _2616_v17;
          _2616_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3);
          let _2617_v18;
          let _2618_v19;
          let _2619_v20;
          let _2620_v21;
          let _out66;
          let _out67;
          let _out68;
          let _out69;
          let _outcollector17 = _module.__default.m0(_2616_v17, globalState);
          _out66 = _outcollector17[0];
          _out67 = _outcollector17[1];
          _out68 = _outcollector17[2];
          _out69 = _outcollector17[3];
          _2617_v18 = _out66;
          _2618_v19 = _out67;
          _2619_v20 = _out68;
          _2620_v21 = _out69;
          let _2621_v22;
          _2621_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(_2618_v19),_2611_v13);
          let _2622_v23;
          let _nw366 = new _module.C0();
          _nw366.__ctor((((_2621_v22).contains(_this.f3)) ? ((_2621_v22).get(_this.f3)) : (_2611_v13)));
          _2622_v23 = _nw366;
          r2 = _dafny.Seq.Concat((_this).f2, (_this).f2);
        } else {
          let _index401 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_2611_v13).length));
          (_2611_v13)[_index401] = (_this).f4;
          let _index402 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_2611_v13).length));
          let _rhs437 = _this.f3;
          let _rhs438 = ((_this.f3) ? (_2596_v0) : (_2596_v0));
          let _rhs439 = (_this).f4;
          let _rhs440 = (_2606_v8).multipliedBy(new BigNumber(((_this).f2).length));
          let _rhs441 = _this.f3;
          let _lhs278 = _this;
          let _lhs279 = _2611_v13;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_2611_v13).length));
          _lhs278.f3 = _rhs437;
          _2596_v0 = _rhs438;
          _2606_v8 = _rhs439;
          _lhs279[_lhs280] = _rhs440;
          r1 = _rhs441;
          let _2623_v25;
          _2623_v25 = _dafny.Seq.of(p0, p0);
          let _2624_v26;
          _2624_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f4),function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of (_2623_v25).Elements) {
              let _2625_v24 = _compr_71;
              if (_dafny.Seq.contains(_2623_v25, _2625_v24)) {
                _coll71.push([_2625_v24,_2606_v8]);
              }
            }
            return _coll71;
          }());
          let _2626_v27;
          _2626_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_2610_v12);
          let _2627_v28;
          _2627_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2626_v27,(_this).f4);
          let _2628_v29;
          _2628_v29 = p0;
          let _2629_v30;
          _2629_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2628_v29,new BigNumber(103));
          _2624_v26 = (_2624_v26).update(((_this).f7).minus((((_2627_v28).contains(_2626_v27)) ? ((_2627_v28).get(_2626_v27)) : ((_this).f7))), (_2629_v30).update(_2628_v29, (_2611_v13)[_module.__default.safeIndex(new BigNumber(407), new BigNumber((_2611_v13).length))]));
          let _2630_v31;
          _2630_v31 = _2597_v2;
          let _2631_v32;
          _2631_v32 = _dafny.Set.fromElements(_2630_v31);
          let _2632_v33;
          _2632_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber((_2631_v32).length));
          let _2633_v34;
          _2633_v34 = _module.D5.create_DC15((_2611_v13)[_module.__default.safeIndex(new BigNumber(407), new BigNumber((_2611_v13).length))], _this.f3, (_this).f5);
          let _2634_v37;
          _2634_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-224)), function (_2635_i3) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          }),_2606_v8);
          let _2636_v38;
          let _nw367 = new _module.C7();
          _nw367.__ctor((_this).f7, (function (_pat_let74_0) {
            return function (_2637_dt__update__tmp_h0) {
              return function (_pat_let75_0) {
                return function (_2638_dt__update_hcf25_h0) {
                  return _module.D5.create_DC15((_2637_dt__update__tmp_h0).dtor_cf24, _2638_dt__update_hcf25_h0, (_2637_dt__update__tmp_h0).dtor_cf26);
                }(_pat_let75_0);
              }(_this.f3);
            }(_pat_let74_0);
          }(_2633_v34)).dtor_cf26, _this.f3, function () {
            let _coll72 = new _dafny.Map();
            for (const _compr_72 of (function () {
              let _coll73 = new _dafny.Map();
              for (const _compr_73 of (_2634_v37).Keys.Elements) {
                let _2639_v36 = _compr_73;
                if ((_2634_v37).contains(_2639_v36)) {
                  _coll73.push([_2639_v36,(_this).f7]);
                }
              }
              return _coll73;
            }()).Keys.Elements) {
              let _2640_v35 = _compr_72;
              if ((function () {
                let _coll74 = new _dafny.Map();
                for (const _compr_74 of (_2634_v37).Keys.Elements) {
                  let _2639_v36 = _compr_74;
                  if ((_2634_v37).contains(_2639_v36)) {
                    _coll74.push([_2639_v36,(_this).f7]);
                  }
                }
                return _coll74;
              }()).contains(_2640_v35)) {
                _coll72.push([_2640_v35,!(_this.f3)]);
              }
            }
            return _coll72;
          }(), (_this).f2);
          _2636_v38 = _nw367;
          let _2641_v39;
          _2641_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2636_v38);
          let _rhs442 = _dafny.Set.fromElements((new BigNumber((_2632_v33).length)).plus(_2606_v8));
          let _rhs443 = _this.f3;
          let _rhs444 = (((new BigNumber((_2641_v39).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_2596_v0).cardinality())))) ? (!((p0)[_module.__default.safeIndex((_2636_v38).f4, new BigNumber((p0).length))])) : (_this.f3));
          let _lhs281 = _this;
          let _lhs282 = _this;
          _2597_v2 = _rhs442;
          _lhs281.f3 = _rhs443;
          _lhs282.f3 = _rhs444;
          let _2642_v40;
          _2642_v40 = _dafny.Set.fromElements(!(_2636_v38.f3));
          _2642_v40 = _2642_v40;
          let _2643_v41;
          let _nw368 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _2643_v41 = _nw368;
          let _2644_v42;
          let _nw369 = new _module.C0();
          _nw369.__ctor(_2643_v41);
          _2644_v42 = _nw369;
        }
        _2606_v8 = (_2606_v8).plus(new BigNumber(((_this).f2).length));
        r2 = _dafny.Seq.UnicodeFromString("bv");
        _2606_v8 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,_this.f3)).update((_this).f5, _this.f3)).length);
        let _2645_v43;
        _2645_v43 = _dafny.Seq.of(_2611_v13, _2611_v13, _2611_v13, _2611_v13);
        let _2646_v44;
        _2646_v44 = _dafny.Seq.of((_this).f7);
        _2606_v8 = (new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_2611_v13, _2611_v13, (_2645_v43)[_module.__default.safeIndex((_this).f4, new BigNumber((_2645_v43).length))]), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2611_v13,_dafny.Seq.update(_2646_v44, _module.__default.safeIndex(_2606_v8, new BigNumber((_2646_v44).length)), (_this).f4))).length), new BigNumber((_dafny.Seq.of(_2611_v13, _2611_v13, (_2645_v43)[_module.__default.safeIndex((_this).f4, new BigNumber((_2645_v43).length))])).length)), _2611_v13)).length)).multipliedBy(_module.__default.safeDivisionInt((_2646_v44)[_module.__default.safeIndex((_this).f4, new BigNumber((_2646_v44).length))], _2606_v8));
      } else {
        let _2647_v45;
        _2647_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,_2606_v8);
        if (!(_module.__default.fm5(_2647_v45, globalState)).equals(_2607_v9)) {
          let _2648_v46;
          let _nw370 = new _module.C2();
          _nw370.__ctor((_this).f7, _this.f3, (_this).f1, (_this).f2);
          _2648_v46 = _nw370;
          let _2649_v47;
          _2649_v47 = _module.D13.create_DC33(_2648_v46);
          let _2650_v48;
          _2650_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2649_v47,(_2648_v46).f13);
          let _2651_v49;
          _2651_v49 = new _dafny.CodePoint('r'.codePointAt(0));
          let _2652_v50;
          _2652_v50 = _module.D13.create_DC34(_this.f3, _this.f3, (_2648_v46).f13, _dafny.Seq.UnicodeFromString("tmtl"));
          let _2653_v51;
          let _init57 = ((_2654_v46) => function (_2655_i4) {
            return (_2655_i4).multipliedBy((_2654_v46).f13);
          })(_2648_v46);
          let _nw371 = Array((new BigNumber(12)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw371.length); _i0_57++) {
            _nw371[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2653_v51 = _nw371;
          let _2656_v52;
          let _nw372 = new _module.C0();
          _nw372.__ctor(_2653_v51);
          _2656_v52 = _nw372;
          let _2657_v53;
          _2657_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2656_v52,_this.f3);
          let _2658_v54;
          _2658_v54 = _dafny.Seq.of((_2648_v46).f13, (_this).f7, (_this).f4);
          let _2659_v55;
          _2659_v55 = _dafny.Set.fromElements((_this).f2);
          let _2660_v56;
          let _nw373 = Array((new BigNumber(15)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f4));
          _nw373[(_dafny.ONE).toNumber()] = new BigNumber((_2650_v48).length);
          _nw373[(new BigNumber(2)).toNumber()] = new BigNumber(((_this).fm14(_this.f3, _this.f3, _2651_v49, _2606_v8, globalState)).length);
          _nw373[(new BigNumber(3)).toNumber()] = (_this).f7;
          _nw373[(new BigNumber(4)).toNumber()] = new BigNumber(((_2652_v50).dtor_cf56).length);
          _nw373[(new BigNumber(5)).toNumber()] = (_this).f4;
          _nw373[(new BigNumber(6)).toNumber()] = (new BigNumber((_2657_v53).length)).plus(_2606_v8);
          _nw373[(new BigNumber(7)).toNumber()] = (_2658_v54)[_module.__default.safeIndex((_this).f7, new BigNumber((_2658_v54).length))];
          _nw373[(new BigNumber(8)).toNumber()] = ((_2648_v46).f13).plus((_this).f4);
          _nw373[(new BigNumber(9)).toNumber()] = (new BigNumber((_2659_v55).length)).minus((_2648_v46).f13);
          _nw373[(new BigNumber(10)).toNumber()] = _2606_v8;
          _nw373[(new BigNumber(11)).toNumber()] = (_2606_v8).minus(new BigNumber(((_this).f2).length));
          _nw373[(new BigNumber(12)).toNumber()] = (_this).f4;
          _nw373[(new BigNumber(13)).toNumber()] = (_2648_v46).f13;
          _nw373[(new BigNumber(14)).toNumber()] = new BigNumber(-352);
          _2660_v56 = _nw373;
          let _index403 = _module.__default.safeIndex(new BigNumber(833), new BigNumber(((_2656_v52).f14).length));
          ((_2656_v52).f14)[_index403] = new BigNumber(((_this).f2).length);
          let _2661_v58;
          _2661_v58 = _dafny.Seq.of(function () {
            let _coll75 = new _dafny.Set();
            for (const _compr_75 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), ((_2662_v8) => function (_2663_i5) {
              return _2662_v8;
            })(_2606_v8))).Elements) {
              let _2664_v57 = _compr_75;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), ((_2665_v8) => function (_2663_i5) {
                return _2665_v8;
              })(_2606_v8)), _2664_v57)) {
                _coll75.add(_module.__default.safeModuloInt(_2664_v57, new BigNumber(-509)));
              }
            }
            return _coll75;
          }(), (_2597_v2).Union(_2597_v2), _module.__default.fm33(new BigNumber((_2658_v54).length), globalState));
          let _index404 = _module.__default.safeIndex(new BigNumber(833), new BigNumber(((_2656_v52).f14).length));
          ((_2656_v52).f14)[_index404] = new BigNumber((_2661_v58).length);
          let _2666_v59;
          _2666_v59 = p0;
          let _2667_v60;
          _2667_v60 = _dafny.Seq.of((_2666_v59));
          let _2668_v61;
          let _nw374 = new _module.C8();
          _nw374.__ctor(_2667_v60, (p0)[_module.__default.safeIndex(((_2656_v52).f14)[_module.__default.safeIndex(new BigNumber(833), new BigNumber(((_2656_v52).f14).length))], new BigNumber((p0).length))], (_this).f1, (_this).f2, ((_this).f7).multipliedBy(new BigNumber((_2597_v2).length)), (_this).f5);
          _2668_v61 = _nw374;
          let _rhs445 = (_2656_v52).f14;
          let _rhs446 = _2668_v61;
          let _rhs447 = (((new BigNumber(((_2668_v61).f2).length)).isLessThan(new BigNumber(871))) ? (new BigNumber((_2596_v0).cardinality())) : (new BigNumber(((_2668_v61).f2).length)));
          let _rhs448 = _2668_v61.f3;
          _2660_v56 = _rhs445;
          _2668_v61 = _rhs446;
          _2606_v8 = _rhs447;
          r1 = _rhs448;
          let _2669_v62;
          let _nw375 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2669_v62 = _nw375;
          let _index405 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_2669_v62).length));
          (_2669_v62)[_index405] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-3)), ((_2670_v49) => function (_2671_i6) {
            return _2670_v49;
          })(_2651_v49));
          let _index406 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_2669_v62).length));
          (_2669_v62)[_index406] = (_2668_v61).f2;
          let _2672_v63;
          _2672_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2658_v54,_module.__default.fm3(_2606_v8, globalState));
          let _2673_v65;
          _2673_v65 = _dafny.Set.fromElements(_2658_v54, _dafny.Seq.of((_this).f4));
          let _2674_v67;
          _2674_v67 = _dafny.Seq.of(_2672_v63, _2672_v63, _2672_v63);
          let _2675_v71;
          let _nw376 = Array((new BigNumber(22)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = _2672_v63;
          _nw376[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2658_v54,_2668_v61.f3);
          _nw376[(new BigNumber(2)).toNumber()] = function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of (_2673_v65).Elements) {
              let _2676_v64 = _compr_76;
              if ((_2673_v65).contains(_2676_v64)) {
                _coll76.push([_2676_v64,_this.f3]);
              }
            }
            return _coll76;
          }();
          _nw376[(new BigNumber(3)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(4)).toNumber()] = (_2672_v63).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2658_v54,_this.f3));
          _nw376[(new BigNumber(5)).toNumber()] = (_2672_v63).Merge(_2672_v63);
          _nw376[(new BigNumber(6)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(7)).toNumber()] = (_2672_v63).Merge(_2672_v63);
          _nw376[(new BigNumber(8)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2658_v54,!(_2668_v61.f3))).Merge((_2672_v63).update(_module.__default.fm31(_2606_v8, _this.f3, _module.__default.fm0(globalState), _2651_v49, globalState), false));
          _nw376[(new BigNumber(10)).toNumber()] = (_2672_v63).Merge(function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(364)))).Elements) {
              let _2677_v66 = _compr_77;
              if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(364)))).contains(_2677_v66)) {
                _coll77.push([_2677_v66,_2668_v61.f3]);
              }
            }
            return _coll77;
          }());
          _nw376[(new BigNumber(11)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(12)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(13)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(14)).toNumber()] = (_2674_v67)[_module.__default.safeIndex((_2648_v46).f13, new BigNumber((_2674_v67).length))];
          _nw376[(new BigNumber(15)).toNumber()] = (_2672_v63).Merge(function () {
            let _coll78 = new _dafny.Map();
            for (const _compr_78 of (function () {
              let _coll79 = new _dafny.Map();
              for (const _compr_79 of (_2672_v63).Keys.Elements) {
                let _2678_v69 = _compr_79;
                if ((_2672_v63).contains(_2678_v69)) {
                  _coll79.push([_2678_v69,_2606_v8]);
                }
              }
              return _coll79;
            }()).Keys.Elements) {
              let _2679_v68 = _compr_78;
              if ((function () {
                let _coll80 = new _dafny.Map();
                for (const _compr_80 of (_2672_v63).Keys.Elements) {
                  let _2678_v69 = _compr_80;
                  if ((_2672_v63).contains(_2678_v69)) {
                    _coll80.push([_2678_v69,_2606_v8]);
                  }
                }
                return _coll80;
              }()).contains(_2679_v68)) {
                _coll78.push([_2679_v68,_this.f3]);
              }
            }
            return _coll78;
          }());
          _nw376[(new BigNumber(16)).toNumber()] = function () {
            let _coll81 = new _dafny.Map();
            for (const _compr_81 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), ((_2680_v54) => function (_2681_i7) {
              return _2680_v54;
            })(_2658_v54))).Elements) {
              let _2682_v70 = _compr_81;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), ((_2683_v54) => function (_2681_i7) {
                return _2683_v54;
              })(_2658_v54)), _2682_v70)) {
                _coll81.push([_2682_v70,!(_this.f3)]);
              }
            }
            return _coll81;
          }();
          _nw376[(new BigNumber(17)).toNumber()] = (_2672_v63).Merge((_2672_v63).update(_2658_v54, _2668_v61.f3));
          _nw376[(new BigNumber(18)).toNumber()] = (_2672_v63).Merge(_2672_v63);
          _nw376[(new BigNumber(19)).toNumber()] = (_2672_v63).Merge((_2674_v67)[_module.__default.safeIndex(((_2656_v52).f14)[_module.__default.safeIndex(new BigNumber(833), new BigNumber(((_2656_v52).f14).length))], new BigNumber((_2674_v67).length))]);
          _nw376[(new BigNumber(20)).toNumber()] = _2672_v63;
          _nw376[(new BigNumber(21)).toNumber()] = _2672_v63;
          _2675_v71 = _nw376;
          let _2684_v72;
          _2684_v72 = _dafny.Seq.of(_2658_v54, _2658_v54);
          let _index407 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2675_v71).length));
          (_2675_v71)[_index407] = ((_2672_v63).update((_2684_v72)[_module.__default.safeIndex((_this).f4, new BigNumber((_2684_v72).length))], _this.f3)).Merge(_2672_v63);
          let _2685_v74;
          _2685_v74 = _dafny.MultiSet.fromElements(_2658_v54, _2658_v54);
          let _index408 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_2675_v71).length));
          (_2675_v71)[_index408] = function () {
            let _coll82 = new _dafny.Map();
            for (const _compr_82 of ((_2685_v74).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_2658_v54, _2658_v54)))).Elements) {
              let _2686_v73 = _compr_82;
              if (((_2685_v74).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_2658_v54, _2658_v54)))).contains(_2686_v73)) {
                _coll82.push([_2686_v73,(_2652_v50).dtor_cf53]);
              }
            }
            return _coll82;
          }();
          let _2687_v75;
          let _nw377 = Array((new BigNumber(8)).toNumber());
          _2687_v75 = _nw377;
          let _2688_v76;
          let _nw378 = new _module.C5();
          _nw378.__ctor(_this.f3, (_2668_v61).f1, (_this).f2);
          _2688_v76 = _nw378;
          let _index409 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_2687_v75).length));
          (_2687_v75)[_index409] = _2688_v76;
          let _index410 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_2687_v75).length));
          let _nw379 = new _module.C5();
          _nw379.__ctor(_2668_v61.f3, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wcotffvl"),!(_2668_v61.f3)), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qpehooaj"), (_this).fm14((_2688_v76).f15, true, _2651_v49, ((_2656_v52).f14)[_module.__default.safeIndex(new BigNumber(833), new BigNumber(((_2656_v52).f14).length))], globalState)));
          (_2687_v75)[_index410] = _nw379;
        } else {
          _2606_v8 = _module.__default.safeModuloInt((_this).f4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f5), _dafny.Seq.of((_this).f5, (_this).f5, (_this).f5))).length));
          let _2689_v77;
          _2689_v77 = _module.D13.create_DC34(_this.f3, (((_2605_v7).contains(_this.f3)) ? ((_2605_v7).get(_this.f3)) : (false)), (_this).f7, (_this).f2);
          let _2690_v78;
          let _nw380 = new _module.C2();
          _nw380.__ctor(_2606_v8, _this.f3, (_this).f1, (_2689_v77).dtor_cf56);
          _2690_v78 = _nw380;
          let _2691_v79;
          let _nw381 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _2691_v79 = _nw381;
          let _index411 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2691_v79).length));
          (_2691_v79)[_index411] = _2606_v8;
          let _index412 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2691_v79).length));
          let _rhs449 = _2690_v78;
          let _rhs450 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f7));
          let _lhs283 = _2691_v79;
          let _lhs284 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2691_v79).length));
          _2690_v78 = _rhs449;
          _lhs283[_lhs284] = _rhs450;
          let _2692_v80;
          let _nw382 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2692_v80 = _nw382;
          _2692_v80 = _2692_v80;
          let _2693_v81;
          let _nw383 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2693_v81 = _nw383;
          let _2694_v82;
          _2694_v82 = new _dafny.CodePoint('j'.codePointAt(0));
          let _index413 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2693_v81).length));
          (_2693_v81)[_index413] = _2694_v82;
          let _index414 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2693_v81).length));
          let _rhs451 = _2694_v82;
          let _rhs452 = _2693_v81;
          let _rhs453 = (p0)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((p0).length))];
          let _rhs454 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f2, (_this).f2), _dafny.Seq.Concat((_this).f2, (_this).f2));
          let _lhs285 = _2693_v81;
          let _lhs286 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_2693_v81).length));
          _lhs285[_lhs286] = _rhs451;
          _2693_v81 = _rhs452;
          r1 = _rhs453;
          r2 = _rhs454;
          let _2695_v83;
          _2695_v83 = _module.D14.create_DC37(_2693_v81);
          let _index415 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2691_v79).length));
          let _rhs455 = _2695_v83;
          let _rhs456 = (_2690_v78).f13;
          let _lhs287 = _2691_v79;
          let _lhs288 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_2691_v79).length));
          _2695_v83 = _rhs455;
          _lhs287[_lhs288] = _rhs456;
        }
        (_this).f3 = (((_this).f4).minus(_2606_v8)).isLessThan(new BigNumber((_2607_v9).cardinality()));
        if (!(((new BigNumber(((_this).f2).length)).multipliedBy(new BigNumber(-120))).isLessThanOrEqualTo((_dafny.ZERO).minus((_this).f4)))) {
          let _2696_v84;
          _2696_v84 = _dafny.Seq.of((_2606_v8).isLessThan((_this).f4));
          _2696_v84 = _2696_v84;
          (_this).f3 = (_2606_v8).isLessThan(new BigNumber((_2607_v9).cardinality()));
          let _2697_v85;
          _2697_v85 = new _dafny.CodePoint('x'.codePointAt(0));
          _2697_v85 = _2697_v85;
          let _2698_v86;
          let _nw384 = new _module.C3();
          _nw384.__ctor((_this).f4, (_this).f5, !(new BigNumber(320)).isEqualTo(_2606_v8), (_this).f1, (_this).f2);
          _2698_v86 = _nw384;
          let _2699_v87;
          let _nw385 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2699_v87 = _nw385;
          _2699_v87 = (_this).f5;
        } else {
          _2606_v8 = (_this).f7;
          let _2700_v88;
          _2700_v88 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,new _dafny.CodePoint('o'.codePointAt(0)));
          let _2701_v89;
          _2701_v89 = new _dafny.CodePoint('l'.codePointAt(0));
          _2700_v88 = (_2700_v88).update(((_this).f7).isEqualTo((_this).f4), _2701_v89);
          let _2702_v90;
          _2702_v90 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f3) ? ((_this).f2) : ((_this).f2)),(_module.__default.fm60(globalState)));
          _2702_v90 = (_2702_v90).Merge(_2702_v90);
          let _2703_v91;
          _2703_v91 = _dafny.Seq.of(_module.__default.fm9(globalState), _module.__default.fm9(globalState));
          let _2704_v92;
          _2704_v92 = _dafny.Map.Empty.slice().updateUnsafe(_2703_v91,_this.f3);
          _2704_v92 = (_2704_v92).update(_2703_v91, _this.f3);
          _2606_v8 = (_this).f4;
        }
        r2 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f2, (_this).f2), _dafny.Seq.UnicodeFromString("ghwey"));
        let _2705_v93;
        let _nw386 = Array((_dafny.ONE).toNumber());
        _2705_v93 = _nw386;
        _2705_v93 = _2705_v93;
      }
      let _2706_i8;
      _2706_i8 = _dafny.ZERO;
      L15: {
        while (_this.f3) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2706_i8)) {
              break L15;
            }
            _2706_i8 = (_2706_i8).plus(_dafny.ONE);
            let _2707_v94;
            let _nw387 = Array((new BigNumber(23)).toNumber());
            _2707_v94 = _nw387;
            let _2708_v95;
            let _nw388 = new _module.C3();
            _nw388.__ctor((_this).f7, (_this).f5, _this.f3, (_this).f1, (_this).f2);
            _2708_v95 = _nw388;
            let _index416 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2707_v94).length));
            (_2707_v94)[_index416] = ((_this.f3) ? (_2708_v95) : (_2708_v95));
            let _index417 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_2707_v94).length));
            (_2707_v94)[_index417] = _2708_v95;
            r1 = _this.f3;
            let _2709_v96;
            _2709_v96 = new _dafny.CodePoint('q'.codePointAt(0));
            let _2710_v97;
            _2710_v97 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,_2709_v96);
            let _2711_v98;
            _2711_v98 = _module.D2.create_DC6(_2709_v96, _this.f3, true, new BigNumber(516));
            _2710_v97 = (_2710_v97).update(!((_this.f3) === (true)), (_2711_v98).dtor_cf9);
            let _2712_v99;
            _2712_v99 = _dafny.Set.fromElements((_this).f7);
            let _2713_v100;
            _2713_v100 = _dafny.Set.fromElements(_2709_v96);
            let _2714_v101;
            _2714_v101 = _2713_v100;
            let _rhs457 = _2712_v99;
            let _rhs458 = _2714_v101;
            _2712_v99 = _rhs457;
            _2714_v101 = _rhs458;
          }
        }
      }
      _2606_v8 = (((_this).f7).multipliedBy((_this).f7)).minus(((_this.f3) ? ((_this).f7) : ((_this).f4)));
      let _2715_v102;
      _2715_v102 = _module.D13.create_DC34(_this.f3, _this.f3, (_this).f7, (_this).f2);
      let _2716_v103;
      _2716_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2715_v102,_this.f3);
      let _2717_v105;
      _2717_v105 = _dafny.Map.Empty.slice().updateUnsafe(_2606_v8,(_this).f4);
      let _2718_v106;
      _2718_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2717_v105,_2716_v103);
      let _2719_v107;
      _2719_v107 = _module.D18.create_DC44(_2716_v103);
      let _2720_v108;
      let _nw389 = Array((new BigNumber(24)).toNumber());
      _nw389[(_dafny.ZERO).toNumber()] = _2716_v103;
      _nw389[(_dafny.ONE).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(2)).toNumber()] = function () {
        let _coll83 = new _dafny.Map();
        for (const _compr_83 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_2715_v102))).Elements) {
          let _2721_v104 = _compr_83;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_2715_v102))).contains(_2721_v104)) {
            _coll83.push([_2721_v104,!(_this.f3)]);
          }
        }
        return _coll83;
      }();
      _nw389[(new BigNumber(3)).toNumber()] = (_2716_v103).Merge((((_2718_v106).contains(_2717_v105)) ? ((_2718_v106).get(_2717_v105)) : (_2716_v103)));
      _nw389[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2715_v102,false);
      _nw389[(new BigNumber(5)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(6)).toNumber()] = (_2716_v103).update(_2715_v102, _this.f3);
      _nw389[(new BigNumber(7)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(8)).toNumber()] = (_2716_v103).Merge(_2716_v103);
      _nw389[(new BigNumber(9)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(10)).toNumber()] = (_2716_v103).update(_2715_v102, _this.f3);
      _nw389[(new BigNumber(11)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(12)).toNumber()] = (_2716_v103).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2715_v102,_this.f3));
      _nw389[(new BigNumber(13)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(14)).toNumber()] = ((_this.f3) ? ((_2719_v107).dtor_cf74) : (_2716_v103));
      _nw389[(new BigNumber(15)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(16)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(17)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2715_v102,!(_this.f3));
      _nw389[(new BigNumber(19)).toNumber()] = (_2716_v103).Merge(_2716_v103);
      _nw389[(new BigNumber(20)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(21)).toNumber()] = _2716_v103;
      _nw389[(new BigNumber(22)).toNumber()] = (_2716_v103).Merge((_dafny.Map.Empty.slice().updateUnsafe(_2715_v102,_this.f3)).update(_2715_v102, _this.f3));
      _nw389[(new BigNumber(23)).toNumber()] = _2716_v103;
      _2720_v108 = _nw389;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2720_v108).length))) {
        let _2722_i9 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2722_i9)) && ((_2722_i9).isLessThan(new BigNumber((_2720_v108).length))))) {
          (_2720_v108)[(_2722_i9)] = _2716_v103;
        }
      }
      let _2723_v109;
      _2723_v109 = new _dafny.CodePoint('v'.codePointAt(0));
      r0 = _module.__default.fm31((_this).f7, _this.f3, _module.__default.safeModuloInt(new BigNumber(794), new BigNumber(979)), _2723_v109, globalState);
      r1 = !(_module.__default.fm3(_2606_v8, globalState)) || (true);
      r2 = _dafny.Seq.Concat((_this).f2, (_this).f2);
      let _2724_v110;
      let _nw390 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _2724_v110 = _nw390;
      r3 = _2724_v110;
      return [r0, r1, r2, r3];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _2725_v0;
      _2725_v0 = new _dafny.CodePoint('k'.codePointAt(0));
      _2725_v0 = _2725_v0;
      let _2726_v1;
      _2726_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,p0);
      let _hi9 = new BigNumber((_2726_v1).length);
      for (let _2727_i0 = (_this).f7; _2727_i0.isLessThan(_hi9); _2727_i0 = _2727_i0.plus(_dafny.ONE)) {
        let _2728_v2;
        _2728_v2 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(_this.f3, new _dafny.CodePoint('j'.codePointAt(0)), globalState),_this.f3);
        let _index418 = _module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index418] = _this.f3;
        let _index419 = _module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length));
        let _rhs459 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p0, (_this).f7), (_this).f7);
        let _rhs460 = _2728_v2;
        let _rhs461 = _this.f3;
        let _lhs289 = (_this).f5;
        let _lhs290 = _module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length));
        r0 = _rhs459;
        _2728_v2 = _rhs460;
        _lhs289[_lhs290] = _rhs461;
        _2725_v0 = _2725_v0;
        let _2729_v4;
        _2729_v4 = _dafny.Seq.of(_module.__default.fm0(globalState), new BigNumber((function () {
          let _coll84 = new _dafny.Set();
          for (const _compr_84 of _dafny.IntegerRange(new BigNumber(998), new BigNumber(916))) {
            let _2730_v3 = _compr_84;
            if (((new BigNumber(998)).isLessThanOrEqualTo(_2730_v3)) && ((_2730_v3).isLessThan(new BigNumber(916)))) {
              _coll84.add(_module.__default.safeModuloInt(_2730_v3, p0));
            }
          }
          return _coll84;
        }()).length));
        let _2731_v5;
        _2731_v5 = _dafny.Seq.of(new BigNumber((_2729_v4).length), p0);
        let _2732_v6;
        _2732_v6 = _dafny.Set.fromElements(new BigNumber((_2731_v5).length), _2727_i0);
        let _2733_v7;
        _2733_v7 = _dafny.Seq.of(_this.f3);
        let _2734_v8;
        let _nw391 = Array((new BigNumber(3)).toNumber()).fill(false);
        _2734_v8 = _nw391;
        let _2735_v9;
        let _nw392 = new _module.C8();
        _nw392.__ctor(_dafny.Seq.of(_2733_v7, _2733_v7), ((_this).f5)[_module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length))], (_dafny.Map.Empty.slice().updateUnsafe((_this).f2,((_this).f5)[_module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length))])).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), ((_2736_v0) => function (_2737_i1) {
          return _2736_v0;
        })(_2725_v0)), true), _dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), ((_2738_v0) => function (_2739_i2) {
          return _2738_v0;
        })(_2725_v0)), (_this).f4, _2734_v8);
        _2735_v9 = _nw392;
        let _2740_v10;
        _2740_v10 = _dafny.Seq.of(_2735_v9);
        let _2741_v11;
        _2741_v11 = _module.D19.create_DC48(_2740_v10);
        _2726_v1 = (_2726_v1).update((_2727_i0).isLessThanOrEqualTo((_this).f4), ((((_this).f5)[_module.__default.safeIndex(new BigNumber(337), new BigNumber(((_this).f5).length))]) ? (new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements((_this).f4, _2727_i0), _2732_v6)).length)) : (new BigNumber(((_2741_v11).dtor_cf79).length))));
        let _2742_v12;
        _2742_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f4);
        let _2743_v13;
        _2743_v13 = _dafny.MultiSet.fromElements(_2731_v5, _2731_v5);
        let _2744_v14;
        let _nw393 = Array((new BigNumber(11)).toNumber());
        _nw393[(_dafny.ZERO).toNumber()] = (_this).f4;
        _nw393[(_dafny.ONE).toNumber()] = (_this).f4;
        _nw393[(new BigNumber(2)).toNumber()] = p0;
        _nw393[(new BigNumber(3)).toNumber()] = _2727_i0;
        _nw393[(new BigNumber(4)).toNumber()] = (((_2742_v12).contains((_this).f4)) ? ((_2742_v12).get((_this).f4)) : (new BigNumber(((_2743_v13).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-111)), ((_2745_i0) => function (_2746_i3) {
          return _2745_i0;
        })(_2727_i0)), _module.__default.abs(new BigNumber((_2732_v6).length)))).cardinality())));
        _nw393[(new BigNumber(5)).toNumber()] = _2727_i0;
        _nw393[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        _nw393[(new BigNumber(7)).toNumber()] = _2727_i0;
        _nw393[(new BigNumber(8)).toNumber()] = _2727_i0;
        _nw393[(new BigNumber(9)).toNumber()] = new BigNumber((_2726_v1).length);
        _nw393[(new BigNumber(10)).toNumber()] = new BigNumber(((_this).f2).length);
        _2744_v14 = _nw393;
        let _2747_v15;
        let _nw394 = new _module.C0();
        _nw394.__ctor(_2744_v14);
        _2747_v15 = _nw394;
      }
      let _index420 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length));
      ((_this).f5)[_index420] = _this.f3;
      let _2748_v16;
      _2748_v16 = _dafny.Seq.UnicodeFromString("qfp");
      let _2749_v17;
      let _nw395 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _2749_v17 = _nw395;
      let _2750_v18;
      _2750_v18 = _dafny.MultiSet.fromElements(_2749_v17, _2749_v17, _2749_v17);
      let _index421 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length));
      let _rhs462 = _module.__default.fm22(true, new BigNumber((_2750_v18).cardinality()), globalState);
      let _rhs463 = ((_this).f4).isLessThanOrEqualTo(p0);
      let _rhs464 = _dafny.Seq.Concat(_dafny.Seq.Concat((_this).f2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-451)), ((_2751_v0) => function (_2752_i4) {
        return _2751_v0;
      })(_2725_v0))), (_this).f2);
      let _lhs291 = (_this).f5;
      let _lhs292 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length));
      _2725_v0 = _rhs462;
      _lhs291[_lhs292] = _rhs463;
      _2748_v16 = _rhs464;
      let _2753_v19;
      let _nw396 = Array((new BigNumber(12)).toNumber()).fill(_module.D3.Default());
      _2753_v19 = _nw396;
      let _ingredients0 = [];
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2753_v19).length))) {
        let _2754_i5 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2754_i5)) && ((_2754_i5).isLessThan(new BigNumber((_2753_v19).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_2753_v19, (_2754_i5).toNumber(), _module.D3.create_DC9(p0, _dafny.Seq.contains(_dafny.Seq.of(_module.D3.create_DC8(_dafny.Set.fromElements(!(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))]), ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], true)), _module.D3.create_DC8(_dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))])), _module.D3.create_DC8(_dafny.Set.fromElements(_this.f3))), _module.D3.create_DC8(_dafny.Set.fromElements(false))))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _2755_v20;
      _2755_v20 = _dafny.Seq.of(p0, p0);
      let _hi10 = (_dafny.ZERO).minus((_2755_v20)[_module.__default.safeIndex(new BigNumber((_module.__default.fm7((_module.D1.create_DC4(_this.f3, _this.f3)).dtor_cf6, _this.f3, globalState)).length), new BigNumber((_2755_v20).length))]);
      for (let _2756_i6 = _module.__default.safeModuloInt((_this).f4, (_this).f4); _2756_i6.isLessThan(_hi10); _2756_i6 = _2756_i6.plus(_dafny.ONE)) {
        r0 = _2756_i6;
        let _2757_v21;
        _2757_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,false);
        _2757_v21 = (_2757_v21).update(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], _this.f3);
        let _2758_v22;
        let _nw397 = Array((new BigNumber(4)).toNumber()).fill([]);
        _2758_v22 = _nw397;
        let _rhs465 = _2758_v22;
        let _rhs466 = p0;
        _2758_v22 = _rhs465;
        r0 = _rhs466;
        let _2759_v23;
        _2759_v23 = _module.D10.create_DC29(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], _this.f3);
        let _2760_v24;
        _2760_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2759_v23,((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))]);
        let _2761_v25;
        let _nw398 = Array((new BigNumber(11)).toNumber()).fill(false);
        _2761_v25 = _nw398;
        let _2762_v26;
        _2762_v26 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))],_2761_v25);
        _2760_v24 = (_2760_v24).update(_2759_v23, (_2762_v26).equals(_2762_v26));
      }
      if (((_this).f4).isLessThanOrEqualTo(new BigNumber(642))) {
        let _2763_v27;
        _2763_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2748_v16).length),new BigNumber(535));
        _2763_v27 = ((((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))]) ? (_2763_v27) : (_2763_v27));
        r0 = ((_this).fm12(p0, (_this).f4, (_this).f4, ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], globalState)).plus(((_this).f7).multipliedBy((_this).f7));
        (_this).f3 = _this.f3;
        if (!(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))])) {
          let _2764_v28;
          let _nw399 = new _module.C1();
          _nw399.__ctor();
          _2764_v28 = _nw399;
          let _2765_v29;
          _2765_v29 = _dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], _this.f3);
          let _2766_v30;
          _2766_v30 = _dafny.Seq.of(_2765_v29, _2765_v29);
          let _2767_v31;
          let _nw400 = Array((new BigNumber(9)).toNumber());
          _nw400[(_dafny.ZERO).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(_dafny.ONE).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(2)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(3)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(4)).toNumber()] = _this.f3;
          _nw400[(new BigNumber(5)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(6)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(7)).toNumber()] = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _nw400[(new BigNumber(8)).toNumber()] = true;
          _2767_v31 = _nw400;
          let _2768_v32;
          let _nw401 = new _module.C8();
          _nw401.__ctor(_2766_v30, _this.f3, (_this).f1, (_this).f2, (_2764_v28).fm28(((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))], (_this).f4, globalState), _2767_v31);
          _2768_v32 = _nw401;
          let _2769_v33;
          _2769_v33 = _dafny.Seq.of(_2768_v32);
          let _2770_v34;
          _2770_v34 = _dafny.Seq.of((_2769_v33)[_module.__default.safeIndex((_this).f7, new BigNumber((_2769_v33).length))]);
          let _2771_v35;
          _2771_v35 = _module.D19.create_DC48(_2770_v34);
          let _2772_v36;
          _2772_v36 = _dafny.Seq.of(_dafny.Seq.of(_2768_v32), _2769_v33, _2769_v33, _2769_v33, _2770_v34);
          let _2773_v37;
          let _nw402 = Array((new BigNumber(21)).toNumber());
          _nw402[(_dafny.ZERO).toNumber()] = _2771_v35;
          _nw402[(_dafny.ONE).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(2)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(3)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(4)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(5)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(6)).toNumber()] = _module.D19.create_DC48(_2769_v33);
          _nw402[(new BigNumber(7)).toNumber()] = _module.D19.create_DC48(_2770_v34);
          _nw402[(new BigNumber(8)).toNumber()] = _module.D19.create_DC48(_2769_v33);
          _nw402[(new BigNumber(9)).toNumber()] = _module.D19.create_DC48(_2770_v34);
          _nw402[(new BigNumber(10)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(11)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(12)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(13)).toNumber()] = _module.D19.create_DC48((_2772_v36)[_module.__default.safeIndex((_this).f7, new BigNumber((_2772_v36).length))]);
          _nw402[(new BigNumber(14)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(15)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(16)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(17)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(18)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(19)).toNumber()] = _2771_v35;
          _nw402[(new BigNumber(20)).toNumber()] = _module.D19.create_DC48(_2769_v33);
          _2773_v37 = _nw402;
          let _2774_v38;
          _2774_v38 = _dafny.Seq.of(_2773_v37);
          let _2775_v39;
          let _nw403 = Array((new BigNumber(22)).toNumber());
          _nw403[(_dafny.ZERO).toNumber()] = _2773_v37;
          _nw403[(_dafny.ONE).toNumber()] = (_2774_v38)[_module.__default.safeIndex((_this).f7, new BigNumber((_2774_v38).length))];
          _nw403[(new BigNumber(2)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(3)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(4)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(5)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(6)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(7)).toNumber()] = (_2774_v38)[_module.__default.safeIndex(p0, new BigNumber((_2774_v38).length))];
          _nw403[(new BigNumber(8)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(9)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(10)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(11)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(12)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(13)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(14)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(15)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(16)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(17)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(18)).toNumber()] = (_2774_v38)[_module.__default.safeIndex((_this).f4, new BigNumber((_2774_v38).length))];
          _nw403[(new BigNumber(19)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(20)).toNumber()] = _2773_v37;
          _nw403[(new BigNumber(21)).toNumber()] = _2773_v37;
          _2775_v39 = _nw403;
          let _index422 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2775_v39).length));
          (_2775_v39)[_index422] = _2773_v37;
          let _index423 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_2775_v39).length));
          (_2775_v39)[_index423] = _2773_v37;
          _2765_v29 = (_2765_v29);
          (_this).f3 = _this.f3;
          let _2776_v40;
          let _nw404 = new _module.C0();
          _nw404.__ctor(_2749_v17);
          _2776_v40 = _nw404;
        } else {
          let _2777_v41;
          _2777_v41 = _dafny.Seq.of(_this.f3);
          let _2778_v42;
          let _2779_v43;
          let _2780_v44;
          let _2781_v45;
          let _out70;
          let _out71;
          let _out72;
          let _out73;
          let _outcollector18 = (_this).m3(_2777_v41, globalState);
          _out70 = _outcollector18[0];
          _out71 = _outcollector18[1];
          _out72 = _outcollector18[2];
          _out73 = _outcollector18[3];
          _2778_v42 = _out70;
          _2779_v43 = _out71;
          _2780_v44 = _out72;
          _2781_v45 = _out73;
          let _2782_v46;
          let _nw405 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _2782_v46 = _nw405;
          let _index424 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_2782_v46).length));
          (_2782_v46)[_index424] = _2725_v0;
          let _2783_v47;
          _2783_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2779_v43,_2725_v0);
          let _index425 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_2782_v46).length));
          (_2782_v46)[_index425] = (((_2783_v47).contains(_this.f3)) ? ((_2783_v47).get(_this.f3)) : (new _dafny.CodePoint('v'.codePointAt(0))));
          (_this).f3 = (_this.f3) === (_2779_v43);
          let _rhs467 = _dafny.Seq.of((_this).f7, (_this).f7, p0);
          let _rhs468 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
          _2778_v42 = _rhs467;
          _2779_v43 = _rhs468;
          let _index426 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2781_v45).length));
          (_2781_v45)[_index426] = (_this).f7;
          let _index427 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_2781_v45).length));
          (_2781_v45)[_index427] = p0;
        }
        let _index428 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_2749_v17).length));
        (_2749_v17)[_index428] = (_this).f7;
        let _index429 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_2749_v17).length));
        (_2749_v17)[_index429] = p0;
      } else {
        (_this).f3 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
        r0 = (((_this).f7).multipliedBy(_module.__default.fm0(globalState))).plus((_this).f4);
        r0 = new BigNumber(226);
        r1 = (_this).fm11(_this.f3, globalState);
        let _index430 = _module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index430] = _this.f3;
      }
      r0 = p0;
      r1 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(879), new BigNumber(((_this).f5).length))];
      return [r0, r1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
      this._f3 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.Seq.UnicodeFromString("");
      this._f4 = _dafny.ZERO;
      this._f5 = [];
      this._f6 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    set f3(value) {
      let _this = this;
      _this._f3 = value;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    __ctor(f6, f4, f5, f3, f1, f2) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f3 = f3;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    fm13(globalState) {
      let _this = this;
      return true;
    };
    fm14(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_2784_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), (_this).f2);
    };
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.of(((_this).f4).plus((_dafny.ZERO).minus((_this).f4)), new BigNumber(623), (_this).f4)).length);
    };
    fm11(p0, globalState) {
      let _this = this;
      return (_this).f6;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _2785_v0;
      _2785_v0 = _dafny.Seq.of(_this.f3);
      let _2786_v1;
      _2786_v1 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      r1 = (_module.__default.safeModuloInt(new BigNumber((_2785_v0).length), new BigNumber((_module.__default.fm7((_this).f6, !(_this.f3), globalState)).length))).plus(new BigNumber((_2786_v1).length));
      let _2787_v2;
      _2787_v2 = _dafny.Seq.of(_2785_v0);
      let _2788_v4;
      _2788_v4 = _dafny.Set.fromElements((_this).f2);
      let _2789_v5;
      _2789_v5 = _module.D20.create_DC51(function () {
  let _coll85 = new _dafny.Map();
  for (const _compr_85 of (_2788_v4).Elements) {
    let _2790_v3 = _compr_85;
    if ((_2788_v4).contains(_2790_v3)) {
      _coll85.push([_2790_v3,_this.f3]);
    }
  }
  return _coll85;
}());
      let _2791_v6;
      let _nw406 = new _module.C8();
      _nw406.__ctor(_2787_v2, _this.f3, (((_this).f6) ? ((_this).f1) : (((_2789_v5).dtor_cf83).update(_dafny.Seq.UnicodeFromString("uokvrp"), (_this).f6))), (_this).f2, ((_this).f4).multipliedBy(p0), (_this).f5);
      _2791_v6 = _nw406;
      let _hi11 = (_this).f4;
      for (let _2792_i0 = p1; _2792_i0.isLessThan(_hi11); _2792_i0 = _2792_i0.plus(_dafny.ONE)) {
        let _2793_v7;
        _2793_v7 = new _dafny.CodePoint('b'.codePointAt(0));
        let _2794_v8;
        _2794_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_this.f3);
        let _2795_v9;
        _2795_v9 = _module.D8.create_DC24(_2793_v7, _2794_v8);
        let _2796_v10;
        _2796_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2795_v9,_2785_v0);
        let _2797_v11;
        _2797_v11 = _dafny.Seq.of(_2792_i0, (_this).f4, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(896)), ((_2798_v7) => function (_2799_i1) {
          return _2798_v7;
        })(_2793_v7))).length));
        if (_dafny.Seq.IsPrefixOf((((_2796_v10).contains(_2795_v9)) ? ((_2796_v10).get(_2795_v9)) : (_module.__default.fm7((_this).f6, (_2791_v6).fm11(_this.f3, globalState), globalState))), _dafny.Seq.update(_2785_v0, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber((_2797_v11).length), new BigNumber(((_this).f2).length)), _2793_v7)).length), new BigNumber((_2785_v0).length)), _this.f3))) {
          let _2800_v12;
          let _nw407 = Array((_dafny.ONE).toNumber());
          _nw407[(_dafny.ZERO).toNumber()] = p0;
          _2800_v12 = _nw407;
          let _index431 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_2800_v12).length));
          (_2800_v12)[_index431] = p1;
          let _2801_v13;
          _2801_v13 = _dafny.MultiSet.fromElements(_this.f3);
          let _index432 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_2800_v12).length));
          (_2800_v12)[_index432] = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_2794_v8).length)), (((_this).f6) ? ((((_2801_v13).contains(_this.f3)) ? ((_2801_v13).get(_this.f3)) : (_module.__default.fm0(globalState)))) : (_2792_i0)));
          let _2802_v14;
          _2802_v14 = _module.D14.create_DC39(new BigNumber(((_this).f2).length), _2792_i0, new BigNumber(((_this).f2).length));
          let _2803_v15;
          _2803_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_module.__default.fm31(p1, (_this).f6, new BigNumber(-281), _2793_v7, globalState), _dafny.Seq.of((_2802_v14).dtor_cf67))).length),_2800_v12);
          let _2804_v16;
          let _nw408 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _2804_v16 = _nw408;
          let _2805_v17;
          _2805_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm0(globalState));
          let _index433 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2804_v16).length));
          (_2804_v16)[_index433] = _2805_v17;
          let _2806_v18;
          let _nw409 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _2806_v18 = _nw409;
          let _index434 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2804_v16).length));
          let _rhs469 = _2803_v15;
          let _rhs470 = function () {
            let _coll86 = new _dafny.Map();
            for (const _compr_86 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4)).Keys.Elements) {
              let _2807_v19 = _compr_86;
              if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4)).contains(_2807_v19)) {
                _coll86.push([(_2807_v19).multipliedBy((_this).f4),_2792_i0]);
              }
            }
            return _coll86;
          }();
          let _rhs471 = _2802_v14;
          let _rhs472 = _2806_v18;
          let _rhs473 = p0;
          let _lhs293 = _2804_v16;
          let _lhs294 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_2804_v16).length));
          _2803_v15 = _rhs469;
          _lhs293[_lhs294] = _rhs470;
          _2802_v14 = _rhs471;
          _2806_v18 = _rhs472;
          r1 = _rhs473;
          r1 = _2792_i0;
          let _2808_v20;
          _2808_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2793_v7,p0);
          _2808_v20 = _dafny.Map.Empty.slice().updateUnsafe(_2793_v7,(_this).f4);
          let _index435 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_2800_v12).length));
          (_2800_v12)[_index435] = (_2800_v12)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_2800_v12).length))];
        } else {
          let _2809_v21;
          let _init58 = ((_2810_p1) => function (_2811_i2) {
            return (_2811_i2).plus(_2810_p1);
          })(p1);
          let _nw410 = Array((new BigNumber(11)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw410.length); _i0_58++) {
            _nw410[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2809_v21 = _nw410;
          let _2812_v22;
          let _nw411 = new _module.C0();
          _nw411.__ctor(_2809_v21);
          _2812_v22 = _nw411;
          let _2813_v23;
          _2813_v23 = _dafny.Seq.of(_2812_v22, _2812_v22, _2812_v22, _2812_v22);
          (_this).f3 = (_dafny.Seq.IsPrefixOf(_2813_v23, _2813_v23)) && ((_this).f6);
          let _2814_v24;
          _2814_v24 = _dafny.MultiSet.fromElements(p1);
          let _2815_v25;
          let _nw412 = new _module.C6();
          _nw412.__ctor((p0).plus(new BigNumber((_2814_v24).cardinality())), p0, _2792_i0, (_this).f5, (true) === ((_this).f6), ((_this).f1).Merge((_this).f1), ((false) ? ((_this).f2) : (_dafny.Seq.UnicodeFromString("laxa"))));
          _2815_v25 = _nw412;
          r2 = new BigNumber(((_this).f2).length);
          _2797_v11 = _2797_v11;
          let _2816_v26;
          let _nw413 = new _module.C2();
          _nw413.__ctor(p0, _this.f3, (_this).f1, (_this).f2);
          _2816_v26 = _nw413;
          let _2817_v27;
          _2817_v27 = _dafny.Seq.of(_2816_v26);
          (_this).f3 = _module.__default.fm3(new BigNumber((_2817_v27).length), globalState);
        }
        let _2818_v28;
        let _init59 = function (_2819_i3) {
          return _module.__default.safeDivisionInt(_2819_i3, new BigNumber(34));
        };
        let _nw414 = Array((new BigNumber(2)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw414.length); _i0_59++) {
          _nw414[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        _2818_v28 = _nw414;
        let _2820_v29;
        let _nw415 = new _module.C0();
        _nw415.__ctor(_2818_v28);
        _2820_v29 = _nw415;
        let _2821_v30;
        _2821_v30 = _dafny.MultiSet.fromElements(_2797_v11, _2797_v11, _2797_v11, _2797_v11, _dafny.Seq.of(p1));
        let _rhs474 = _2786_v1;
        let _rhs475 = ((_2821_v30).Intersect(_2821_v30)).IsDisjointFrom(_module.__default.fm61(!(_this.f3), globalState));
        let _rhs476 = _2820_v29;
        let _rhs477 = (_2820_v29).f14;
        let _lhs295 = _this;
        _2786_v1 = _rhs474;
        _lhs295.f3 = _rhs475;
        _2820_v29 = _rhs476;
        _2818_v28 = _rhs477;
        let _2822_v31;
        let _nw416 = new _module.C1();
        _nw416.__ctor();
        _2822_v31 = _nw416;
        let _2823_v32;
        _2823_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2822_v31,p0);
        let _2824_v33;
        _2824_v33 = _dafny.Map.Empty.slice().updateUnsafe((_2820_v29).f14,_2823_v32);
        let _2825_v34;
        _2825_v34 = _dafny.MultiSet.fromElements((_this).f6);
        let _2826_v35;
        _2826_v35 = _dafny.Set.fromElements((_this).f6, (_this).f6);
        let _rhs478 = (_dafny.Set.fromElements(p0, (_this).f4, p1)).IsDisjointFrom(_dafny.Set.fromElements((((_2825_v34).contains(_this.f3)) ? ((_2825_v34).get(_this.f3)) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("lofhhys"), _module.__default.safeIndex(new BigNumber((_2826_v35).length), new BigNumber((_dafny.Seq.UnicodeFromString("lofhhys")).length)), new _dafny.CodePoint('x'.codePointAt(0)))).length)))));
        let _rhs479 = (_2824_v33).Merge((_2824_v33).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2820_v29).f14,_2823_v32)));
        let _rhs480 = new BigNumber((_dafny.Seq.Concat(_2797_v11, _2797_v11)).length);
        let _rhs481 = _this.f3;
        let _rhs482 = (_this).fm12(_module.__default.safeModuloInt(p0, new BigNumber(-521)), new BigNumber((_2826_v35).length), _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(_2792_i0)), !_dafny.Seq.contains((_this).f2, new _dafny.CodePoint('l'.codePointAt(0))), globalState);
        let _lhs296 = _this;
        let _lhs297 = _this;
        _lhs296.f3 = _rhs478;
        _2824_v33 = _rhs479;
        r1 = _rhs480;
        _lhs297.f3 = _rhs481;
        r2 = _rhs482;
        let _2827_v36;
        _2827_v36 = _dafny.Set.fromElements(_2792_i0, _2792_i0);
        let _2828_v37;
        let _nw417 = Array((new BigNumber(6)).toNumber()).fill(false);
        _2828_v37 = _nw417;
        let _index436 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2818_v28).length));
        (_2818_v28)[_index436] = (_this).f4;
        let _2829_v38;
        _2829_v38 = _dafny.Seq.of(_2827_v36);
        let _2830_v39;
        _2830_v39 = _dafny.Seq.of((_this).f5);
        let _index437 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2818_v28).length));
        let _rhs483 = (_2829_v38)[_module.__default.safeIndex(p1, new BigNumber((_2829_v38).length))];
        let _rhs484 = (_2830_v39)[_module.__default.safeIndex(_2792_i0, new BigNumber((_2830_v39).length))];
        let _rhs485 = ((_this).f4).minus(p1);
        let _rhs486 = new BigNumber(328);
        let _lhs298 = _2818_v28;
        let _lhs299 = _module.__default.safeIndex(new BigNumber(720), new BigNumber((_2818_v28).length));
        _2827_v36 = _rhs483;
        _2828_v37 = _rhs484;
        _lhs298[_lhs299] = _rhs485;
        r1 = _rhs486;
      }
      if (_this.f3) {
        (_this).f3 = (_this).f6;
        let _2831_v40;
        let _nw418 = new _module.C3();
        _nw418.__ctor(new BigNumber(333), (_this).f5, _dafny.Seq.IsProperPrefixOf(_2785_v0, _2785_v0), (_this).f1, _dafny.Seq.UnicodeFromString("kyvxifc"));
        _2831_v40 = _nw418;
        _2831_v40 = _2831_v40;
        let _2832_v41;
        let _init60 = ((_2833_v40) => function (_2834_i4) {
          return ((_2833_v40.f3) ? (_dafny.Set.fromElements(_2833_v40.f3)) : (_dafny.Set.fromElements(_2833_v40.f3, true)));
        })(_2831_v40);
        let _nw419 = Array((new BigNumber(2)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw419.length); _i0_60++) {
          _nw419[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2832_v41 = _nw419;
        let _index438 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2832_v41).length));
        (_2832_v41)[_index438] = _dafny.Set.fromElements((_this).f6, (_this).f6, (_2785_v0)[_module.__default.safeIndex(p1, new BigNumber((_2785_v0).length))]);
        let _2835_v42;
        _2835_v42 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2836_v43;
        _2836_v43 = _dafny.Set.fromElements((_this).f6, _this.f3);
        let _index439 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_2832_v41).length));
        (_2832_v41)[_index439] = ((_module.__default.fm6(_2831_v40.f3, _2835_v42, globalState)).Difference(_dafny.Set.fromElements(_this.f3))).Difference(_2836_v43);
        let _index440 = _module.__default.safeIndex(new BigNumber(585), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index440] = _this.f3;
        let _index441 = _module.__default.safeIndex(new BigNumber(585), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index441] = (_this).f6;
        let _2837_v44;
        _2837_v44 = _dafny.Seq.of((_2831_v40).f4);
        let _2838_v45;
        let _nw420 = new _module.C2();
        _nw420.__ctor(p0, true, _module.__default.fm62(new BigNumber((_2837_v44).length), (_this).f6, globalState), (_2831_v40).f2);
        _2838_v45 = _nw420;
        let _2839_v46;
        _2839_v46 = _2838_v45;
        let _2840_v47;
        _2840_v47 = _dafny.Seq.of(_2838_v45, _2838_v45, _2838_v45);
        let _2841_v48;
        _2841_v48 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_2838_v45);
        let _2842_v49;
        let _nw421 = Array((new BigNumber(18)).toNumber());
        _nw421[(_dafny.ZERO).toNumber()] = (_2839_v46);
        _nw421[(_dafny.ONE).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(2)).toNumber()] = (_2838_v45);
        _nw421[(new BigNumber(3)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(4)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(5)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(6)).toNumber()] = (_2840_v47)[_module.__default.safeIndex(p0, new BigNumber((_2840_v47).length))];
        _nw421[(new BigNumber(7)).toNumber()] = (_2838_v45);
        _nw421[(new BigNumber(8)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(9)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(10)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(11)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(12)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(13)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(14)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(15)).toNumber()] = (((_2841_v48).contains(p0)) ? ((_2841_v48).get(p0)) : (_2838_v45));
        _nw421[(new BigNumber(16)).toNumber()] = _2838_v45;
        _nw421[(new BigNumber(17)).toNumber()] = _2838_v45;
        _2842_v49 = _nw421;
        let _index442 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_2842_v49).length));
        (_2842_v49)[_index442] = _2838_v45;
        let _index443 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_2842_v49).length));
        (_2842_v49)[_index443] = _2838_v45;
      } else {
        let _2843_v50;
        _2843_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_this.f3);
        let _2844_v51;
        _2844_v51 = new _dafny.CodePoint('i'.codePointAt(0));
        let _2845_v52;
        let _nw422 = new _module.C5();
        _nw422.__ctor((new BigNumber((_2843_v50).length)).isLessThanOrEqualTo(_module.__default.fm0(globalState)), (_this).f1, _dafny.Seq.update((_this).f2, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f2, (_this).f2)).length), new BigNumber(((_this).f2).length)), _2844_v51));
        _2845_v52 = _nw422;
        _2845_v52 = _2845_v52;
        let _2846_v53;
        _2846_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2844_v51,p0);
        let _2847_v54;
        _2847_v54 = _dafny.MultiSet.fromElements(_this.f3, (_this).f6, true, _this.f3, (_this).f6);
        let _2848_v55;
        _2848_v55 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f6);
        let _2849_v56;
        _2849_v56 = _dafny.MultiSet.fromElements((_2845_v52).f2, (_2845_v52).f2, (_2845_v52).f2);
        let _2850_v57;
        _2850_v57 = _dafny.Seq.of((_this).f4, p1, new BigNumber(-492), (_this).f4, p0);
        let _2851_v58;
        _2851_v58 = _module.D8.create_DC24(_2844_v51, _2848_v55);
        let _2852_v59;
        _2852_v59 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_2851_v58).dtor_cf41);
        let _2853_v60;
        let _nw423 = new _module.C4();
        _nw423.__ctor(_this.f3, (_2845_v52).f1, (_2845_v52).f2);
        _2853_v60 = _nw423;
        let _2854_v61;
        _2854_v61 = _dafny.MultiSet.fromElements(p0, (_this).f4);
        let _2855_v62;
        let _nw424 = Array((new BigNumber(24)).toNumber());
        _nw424[(_dafny.ZERO).toNumber()] = (p1).multipliedBy(p0);
        _nw424[(_dafny.ONE).toNumber()] = (p1).multipliedBy(new BigNumber((_2846_v53).length));
        _nw424[(new BigNumber(2)).toNumber()] = new BigNumber((_2847_v54).cardinality());
        _nw424[(new BigNumber(3)).toNumber()] = p0;
        _nw424[(new BigNumber(4)).toNumber()] = new BigNumber(((_2848_v55).Merge(_2848_v55)).length);
        _nw424[(new BigNumber(5)).toNumber()] = (new BigNumber(242)).minus(p1);
        _nw424[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2849_v56).cardinality()), (_this).f4);
        _nw424[(new BigNumber(7)).toNumber()] = (_2850_v57)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_2850_v57).length))];
        _nw424[(new BigNumber(8)).toNumber()] = (_this).f4;
        _nw424[(new BigNumber(9)).toNumber()] = (_this).f4;
        _nw424[(new BigNumber(10)).toNumber()] = new BigNumber(-279);
        _nw424[(new BigNumber(11)).toNumber()] = p1;
        _nw424[(new BigNumber(12)).toNumber()] = new BigNumber(507);
        _nw424[(new BigNumber(13)).toNumber()] = p0;
        _nw424[(new BigNumber(14)).toNumber()] = new BigNumber((_2852_v59).length);
        _nw424[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2786_v1).length), p0));
        _nw424[(new BigNumber(16)).toNumber()] = ((_this.f3) ? (p0) : ((_this).f4));
        _nw424[(new BigNumber(17)).toNumber()] = p1;
        _nw424[(new BigNumber(18)).toNumber()] = ((_this.f3) ? (new BigNumber(172)) : (new BigNumber((_dafny.MultiSet.fromElements(_2853_v60)).cardinality())));
        _nw424[(new BigNumber(19)).toNumber()] = p0;
        _nw424[(new BigNumber(20)).toNumber()] = (new BigNumber((_2854_v61).cardinality())).plus((_this).f4);
        _nw424[(new BigNumber(21)).toNumber()] = p0;
        _nw424[(new BigNumber(22)).toNumber()] = p0;
        _nw424[(new BigNumber(23)).toNumber()] = p1;
        _2855_v62 = _nw424;
        let _index444 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_2855_v62).length));
        (_2855_v62)[_index444] = p1;
        let _index445 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_2855_v62).length));
        (_2855_v62)[_index445] = p0;
        r1 = p0;
        let _2856_v63;
        _2856_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,new BigNumber(((_2847_v54).Intersect(_2847_v54)).cardinality()));
        _2856_v63 = (_2856_v63).Merge(_2856_v63);
        if (_this.f3) {
          (_this).f3 = (_2791_v6).fm11((_this).f6, globalState);
          _2855_v62 = _2855_v62;
          let _index446 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index446] = (_this).f6;
          let _index447 = _module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index447] = _this.f3;
          let _2857_v64;
          let _nw425 = Array((new BigNumber(16)).toNumber()).fill(false);
          _2857_v64 = _nw425;
          let _2858_v65;
          let _nw426 = new _module.C7();
          _nw426.__ctor((_this).f4, _2857_v64, ((_this).f5)[_module.__default.safeIndex(new BigNumber(620), new BigNumber(((_this).f5).length))], (_2845_v52).f1, (_2845_v52).f2);
          _2858_v65 = _nw426;
          _2786_v1 = (_2786_v1).update(((_this).f4).isLessThanOrEqualTo((_2850_v57)[_module.__default.safeIndex(p1, new BigNumber((_2850_v57).length))]), (_this).f4);
        } else {
          (_this).f3 = ((_dafny.ZERO).minus((_this).f4)).isLessThanOrEqualTo((_this).f4);
          r1 = new BigNumber(452);
          r2 = p1;
          (_this).f3 = (((_2843_v50).contains(p1)) ? ((_2843_v50).get(p1)) : (!(_this.f3) || (true)));
          let _2859_v66;
          _2859_v66 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f2,_this.f3));
          let _2860_v67;
          _2860_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f1);
          let _2861_v68;
          let _nw427 = new _module.C8();
          _nw427.__ctor((_2791_v6).f10, _this.f3, ((_2859_v66)[_module.__default.safeIndex((_this).f4, new BigNumber((_2859_v66).length))]).Merge(((((_2860_v67).contains(_this.f3)) ? ((_2860_v67).get(_this.f3)) : ((_this).f1))).update((_2845_v52).f2, _this.f3)), (_this).f2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_2862_v51) => function (_2863_i5) {
            return _2862_v51;
          })(_2844_v51)), (_2845_v52).f2)).length), (_this).f5);
          _2861_v68 = _nw427;
          let _2864_v69;
          _2864_v69 = _2861_v68;
          _2861_v68 = (_2864_v69);
        }
      }
      if (_this.f3) {
        let _2865_v70;
        _2865_v70 = _module.D11.create_DC31((_this).f4);
        let _rhs487 = (_module.__default.fm49(_2865_v70, new BigNumber(((_this).f2).length), p0, globalState)).dtor_cf28;
        let _rhs488 = ((_this).f4).minus(((_2791_v6).fm12((_dafny.ZERO).minus(p1), p1, new BigNumber(124), true, globalState)).minus(new BigNumber(((_this).f2).length)));
        let _rhs489 = _module.__default.fm0(globalState);
        let _lhs300 = _this;
        _lhs300.f3 = _rhs487;
        r2 = _rhs488;
        r2 = _rhs489;
        let _2866_v71;
        _2866_v71 = _dafny.MultiSet.fromElements(_this.f3);
        let _2867_v72;
        _2867_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_dafny.ZERO).minus(new BigNumber(((_2866_v71).update((_this).f6, _module.__default.abs(p0))).cardinality())));
        if (((_dafny.MultiSet.fromElements(_this.f3, true, (_this).f6)).Difference(_2866_v71)).IsProperSubsetOf((_dafny.MultiSet.FromArray(_2785_v0)).Union(_module.__default.fm5(_2867_v72, globalState)))) {
          let _2868_v73;
          _2868_v73 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs490 = _2868_v73;
          let _rhs491 = ((((_2866_v71).contains((_2785_v0)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_2785_v0).length))])) ? ((_2866_v71).get((_2785_v0)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_2785_v0).length))])) : (p1))).isLessThan((p1).multipliedBy(p1));
          let _lhs301 = _this;
          _2868_v73 = _rhs490;
          _lhs301.f3 = _rhs491;
          (_this).f3 = _this.f3;
          let _2869_v74;
          let _init61 = function (_2870_i6) {
            return _module.__default.safeModuloInt(_2870_i6, (_this).f4);
          };
          let _nw428 = Array((new BigNumber(15)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw428.length); _i0_61++) {
            _nw428[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2869_v74 = _nw428;
          let _2871_v75;
          _2871_v75 = _module.D8.create_DC22(_2866_v71);
          let _2872_v76;
          _2872_v76 = _dafny.Seq.of(_2871_v75);
          let _2873_v77;
          _2873_v77 = _dafny.Set.fromElements(_2872_v76);
          let _2874_v78;
          _2874_v78 = _dafny.Seq.of(new BigNumber((_2873_v77).length), p1);
          let _2875_v79;
          _2875_v79 = _dafny.MultiSet.fromElements(_2869_v74);
          let _nw429 = Array((new BigNumber(17)).toNumber());
          _nw429[(_dafny.ZERO).toNumber()] = p0;
          _nw429[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_2874_v78).length), new BigNumber((_dafny.Seq.UnicodeFromString("gc")).length));
          _nw429[(new BigNumber(2)).toNumber()] = p1;
          _nw429[(new BigNumber(3)).toNumber()] = (((_2875_v79).contains(_2869_v74)) ? ((_2875_v79).get(_2869_v74)) : ((_this).f4));
          _nw429[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_this).f4, new BigNumber(((_this).f2).length));
          _nw429[(new BigNumber(5)).toNumber()] = p1;
          _nw429[(new BigNumber(6)).toNumber()] = p1;
          _nw429[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f4);
          _nw429[(new BigNumber(8)).toNumber()] = new BigNumber(28);
          _nw429[(new BigNumber(9)).toNumber()] = p1;
          _nw429[(new BigNumber(10)).toNumber()] = p1;
          _nw429[(new BigNumber(11)).toNumber()] = (_this).f4;
          _nw429[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_this).f2, (_this).f2)).length);
          _nw429[(new BigNumber(13)).toNumber()] = (_this).f4;
          _nw429[(new BigNumber(14)).toNumber()] = (_this).f4;
          _nw429[(new BigNumber(15)).toNumber()] = p0;
          _nw429[(new BigNumber(16)).toNumber()] = new BigNumber((_module.__default.fm20(true, (_this).f6, (_this).fm11(true, globalState), _2785_v0, globalState)).length);
          _2869_v74 = _nw429;
          let _2876_v80;
          let _nw430 = new _module.C2();
          _nw430.__ctor((_this).f4, (_this).f6, (_this).f1, (_this).f2);
          _2876_v80 = _nw430;
          let _2877_v82;
          let _nw431 = new _module.C10();
          _nw431.__ctor(p0, (_this).f4, (_this).f5, (_this).f6, function () {
            let _coll87 = new _dafny.Map();
            for (const _compr_87 of ((_this).f1).Keys.Elements) {
              let _2878_v81 = _compr_87;
              if (((_this).f1).contains(_2878_v81)) {
                _coll87.push([_2878_v81,(_2785_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("mtww")).length), new BigNumber((_2785_v0).length))]]);
              }
            }
            return _coll87;
          }(), (_this).f2);
          _2877_v82 = _nw431;
          let _2879_v83;
          _2879_v83 = _2877_v82;
          _2879_v83 = _2879_v83;
        } else {
          let _2880_v84;
          let _nw432 = new _module.C6();
          _nw432.__ctor((_this).f4, new BigNumber(781), new BigNumber((_dafny.Seq.of(_this.f3, _this.f3)).length), (_this).f5, !((((_this).f6) ? (true) : (!(false)))), (_this).f1, _dafny.Seq.UnicodeFromString("es"));
          _2880_v84 = _nw432;
          let _2881_v85;
          _2881_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
          let _2882_v86;
          _2882_v86 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_2785_v0)[_module.__default.safeIndex((_this).f4, new BigNumber((_2785_v0).length))],_this.f3), _2881_v85);
          let _2883_v87;
          _2883_v87 = _dafny.MultiSet.fromElements(p0);
          let _2884_v88;
          _2884_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2883_v87).cardinality()),(_this).f6);
          (_this).f3 = (function () {
            let _coll88 = new _dafny.Set();
            for (const _compr_88 of (_module.__default.fm63((((_2884_v88).contains((_2880_v84).f11)) ? ((_2884_v88).get((_2880_v84).f11)) : (_this.f3)), _this.f3, _this.f3, (_this).f6, globalState)).Elements) {
              let _2885_v89 = _compr_88;
              if ((_module.__default.fm63((((_2884_v88).contains((_2880_v84).f11)) ? ((_2884_v88).get((_2880_v84).f11)) : (_this.f3)), _this.f3, _this.f3, (_this).f6, globalState)).contains(_2885_v89)) {
                _coll88.add(_2885_v89);
              }
            }
            return _coll88;
          }()).IsSubsetOf(_2882_v86);
          let _2886_v90;
          let _init62 = ((_2887_v0) => function (_2888_i7) {
            return _module.__default.safeDivisionInt(_2888_i7, new BigNumber((_2887_v0).length));
          })(_2785_v0);
          let _nw433 = Array((new BigNumber(21)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw433.length); _i0_62++) {
            _nw433[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2886_v90 = _nw433;
          let _2889_v91;
          _2889_v91 = _dafny.Seq.of(_2886_v90, _2886_v90, _2886_v90, _2886_v90);
          let _2890_v92;
          let _nw434 = new _module.C0();
          _nw434.__ctor((_2889_v91)[_module.__default.safeIndex((_this).f4, new BigNumber((_2889_v91).length))]);
          _2890_v92 = _nw434;
          let _2891_v93;
          _2891_v93 = new _dafny.CodePoint('x'.codePointAt(0));
          let _2892_v94;
          _2892_v94 = _dafny.Map.Empty.slice().updateUnsafe(_2891_v93,_this.f3);
          let _2893_v96;
          _2893_v96 = _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), _2891_v93, _2891_v93);
          let _index448 = _module.__default.safeIndex(new BigNumber(938), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index448] = !((_2893_v96).IsSubsetOf(function () {
            let _coll89 = new _dafny.Set();
            for (const _compr_89 of (_2892_v94).Keys.Elements) {
              let _2894_v95 = _compr_89;
              if ((_2892_v94).contains(_2894_v95)) {
                _coll89.add(_2894_v95);
              }
            }
            return _coll89;
          }()));
          let _index449 = _module.__default.safeIndex(new BigNumber(938), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index449] = (p0).isLessThanOrEqualTo((_this).f4);
          _2866_v71 = _2866_v71;
        }
        let _2895_v97;
        _2895_v97 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-3), (_this).f4)),(_this).f6);
        let _2896_v98;
        _2896_v98 = _module.D6.create_DC16(_2895_v97);
        _2895_v97 = (_2896_v98).dtor_cf27;
        (_this).f3 = (_this).f6;
        r1 = (_this).fm12(p0, new BigNumber(-259), (_this).f4, !(((true) ? ((_2785_v0)[_module.__default.safeIndex(new BigNumber(((_this).f2).length), new BigNumber((_2785_v0).length))]) : (_this.f3))), globalState);
      } else {
        if (_this.f3) {
          let _2897_v99;
          let _init63 = ((_2898_p0) => function (_2899_i8) {
            return (_2899_i8).multipliedBy((_dafny.ZERO).minus(_2898_p0));
          })(p0);
          let _nw435 = Array((new BigNumber(11)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw435.length); _i0_63++) {
            _nw435[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2897_v99 = _nw435;
          _2897_v99 = _2897_v99;
          let _2900_v100;
          _2900_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,!((_this).f6));
          let _2901_v101;
          let _2902_v102;
          let _2903_v103;
          let _2904_v104;
          let _out74;
          let _out75;
          let _out76;
          let _out77;
          let _outcollector19 = _module.__default.m0((_dafny.Map.Empty.slice().updateUnsafe((_this).f5,(_this).f6)).Merge(_2900_v100), globalState);
          _out74 = _outcollector19[0];
          _out75 = _outcollector19[1];
          _out76 = _outcollector19[2];
          _out77 = _outcollector19[3];
          _2901_v101 = _out74;
          _2902_v102 = _out75;
          _2903_v103 = _out76;
          _2904_v104 = _out77;
          let _2905_v105;
          _2905_v105 = _dafny.Seq.of((p0).plus(p0));
          let _2906_v106;
          _2906_v106 = _dafny.MultiSet.fromElements(_2901_v101);
          let _2907_v107;
          _2907_v107 = _dafny.MultiSet.fromElements((_this).f2);
          let _2908_v108;
          _2908_v108 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2907_v107).cardinality()),_2901_v101);
          let _2909_v109;
          _2909_v109 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length), new BigNumber((_dafny.Seq.update(_2905_v105, _module.__default.safeIndex(p1, new BigNumber((_2905_v105).length)), p1)).length), new BigNumber(754), (((_2906_v106).contains((_this).f6)) ? ((_2906_v106).get((_this).f6)) : (p0)), new BigNumber((_2908_v108).length));
          let _2910_v110;
          _2910_v110 = _dafny.Set.fromElements(_2902_v102);
          let _2911_v111;
          _2911_v111 = _dafny.Set.fromElements(_2910_v110);
          let _2912_v112;
          _2912_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2911_v111,new BigNumber((_2905_v105).length));
          _2905_v105 = _dafny.Seq.of((((_2909_v109).contains(new BigNumber((_2905_v105).length))) ? ((_2909_v109).get(new BigNumber((_2905_v105).length))) : (p1)), (((_2912_v112).contains(_2911_v111)) ? ((_2912_v112).get(_2911_v111)) : (new BigNumber(((_this).f2).length))));
          let _index450 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_2897_v99).length));
          (_2897_v99)[_index450] = (_this).f4;
          let _index451 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_2897_v99).length));
          let _rhs492 = ((_this).f6) || (_2902_v102);
          let _rhs493 = (((_2909_v109).contains(p0)) ? ((_2909_v109).get(p0)) : (p0));
          let _rhs494 = new BigNumber(((_this).f2).length);
          let _lhs302 = _2897_v99;
          let _lhs303 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_2897_v99).length));
          _2902_v102 = _rhs492;
          r2 = _rhs493;
          _lhs302[_lhs303] = _rhs494;
          let _2913_v113;
          let _nw436 = new _module.C5();
          _nw436.__ctor(_2901_v101, _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_2903_v103)[_module.__default.safeIndex(new BigNumber((_2785_v0).length), new BigNumber((_2903_v103).length))]), (_this).f2);
          _2913_v113 = _nw436;
          let _2914_v114;
          _2914_v114 = _dafny.Map.Empty.slice().updateUnsafe(_2786_v1,_2901_v101);
          let _2915_v115;
          _2915_v115 = new _dafny.CodePoint('p'.codePointAt(0));
          let _2916_v116;
          _2916_v116 = _dafny.Map.Empty.slice().updateUnsafe(!(_2901_v101),_module.__default.fm6(_this.f3, _2915_v115, globalState));
          let _2917_v117;
          let _nw437 = Array((new BigNumber(20)).toNumber());
          _nw437[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(!((((_2914_v114).contains(_2786_v1)) ? ((_2914_v114).get(_2786_v1)) : (!((_this).f6)))), _2901_v101, (_this).f6);
          _nw437[(_dafny.ONE).toNumber()] = (_2910_v110).Intersect(_dafny.Set.fromElements(false, _this.f3));
          _nw437[(new BigNumber(2)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(3)).toNumber()] = (_2910_v110).Intersect(_dafny.Set.fromElements(true, _2902_v102, true, !((_this).fm11(_2902_v102, globalState))));
          _nw437[(new BigNumber(4)).toNumber()] = ((((_2916_v116).contains(false)) ? ((_2916_v116).get(false)) : (_2910_v110))).Difference(_2910_v110);
          _nw437[(new BigNumber(5)).toNumber()] = _module.__default.fm6(_2901_v101, ((_this).f2)[_module.__default.safeIndex(new BigNumber(894), new BigNumber(((_this).f2).length))], globalState);
          _nw437[(new BigNumber(6)).toNumber()] = (_2910_v110).Difference(_2910_v110);
          _nw437[(new BigNumber(7)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(8)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_2901_v101);
          _nw437[(new BigNumber(10)).toNumber()] = (_2910_v110).Difference(_2910_v110);
          _nw437[(new BigNumber(11)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(12)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(13)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(14)).toNumber()] = (_2910_v110).Intersect(_2910_v110);
          _nw437[(new BigNumber(15)).toNumber()] = (_2910_v110).Union(_2910_v110);
          _nw437[(new BigNumber(16)).toNumber()] = (_module.__default.fm6(_2902_v102, _2915_v115, globalState)).Intersect(_2910_v110);
          _nw437[(new BigNumber(17)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(18)).toNumber()] = _2910_v110;
          _nw437[(new BigNumber(19)).toNumber()] = _2910_v110;
          _2917_v117 = _nw437;
          let _rhs495 = _2913_v113;
          let _rhs496 = _2917_v117;
          _2913_v113 = _rhs495;
          _2917_v117 = _rhs496;
        } else {
          let _2918_v118;
          let _init64 = function (_2919_i9) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          };
          let _nw438 = Array((new BigNumber(20)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw438.length); _i0_64++) {
            _nw438[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2918_v118 = _nw438;
          let _2920_v119;
          _2920_v119 = _dafny.Map.Empty.slice().updateUnsafe(_this.f3,(_this).f6);
          let _2921_v120;
          _2921_v120 = _dafny.Map.Empty.slice().updateUnsafe(_2918_v118,p0);
          let _index452 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_2918_v118).length));
          (_2918_v118)[_index452] = _module.__default.fm38((((_2920_v119).contains((_this).f6)) ? ((_2920_v119).get((_this).f6)) : (_this.f3)), (_this).f6, new BigNumber(((_2921_v120).update(_2918_v118, new BigNumber(603))).length), globalState);
          let _index453 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((_2918_v118).length));
          (_2918_v118)[_index453] = ((_this).f2)[_module.__default.safeIndex((_this).f4, new BigNumber(((_this).f2).length))];
          let _2922_v121;
          _2922_v121 = _dafny.Seq.UnicodeFromString("oqys");
          let _2923_v122;
          let _nw439 = Array((_dafny.ONE).toNumber());
          _nw439[(_dafny.ZERO).toNumber()] = (_this).f5;
          _2923_v122 = _nw439;
          let _index454 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2923_v122).length));
          (_2923_v122)[_index454] = (_this).f5;
          let _2924_v123;
          _2924_v123 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f3);
          let _2925_v124;
          _2925_v124 = _dafny.Set.fromElements((((_2924_v123).contains((_this).f4)) ? ((_2924_v123).get((_this).f4)) : (_this.f3)));
          let _index455 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2923_v122).length));
          let _rhs497 = (_this).f2;
          let _rhs498 = (_this).f5;
          let _rhs499 = (_2925_v124).IsDisjointFrom((_2925_v124).Intersect(_2925_v124));
          let _lhs304 = _2923_v122;
          let _lhs305 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_2923_v122).length));
          let _lhs306 = _this;
          _2922_v121 = _rhs497;
          _lhs304[_lhs305] = _rhs498;
          _lhs306.f3 = _rhs499;
          (_this).f3 = _this.f3;
          let _2926_v125;
          let _nw440 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _2926_v125 = _nw440;
          let _2927_v126;
          _2927_v126 = _dafny.Seq.of(_2926_v125, _2926_v125, _2926_v125);
          _2926_v125 = (_2927_v126)[_module.__default.safeIndex(p1, new BigNumber((_2927_v126).length))];
          let _2928_v127;
          let _nw441 = new _module.C5();
          _nw441.__ctor((_this).f6, ((_this).f1).update(_2922_v121, _this.f3), _2922_v121);
          _2928_v127 = _nw441;
          let _2929_v128;
          _2929_v128 = _dafny.Set.fromElements((_this).f4);
          let _2930_v129;
          _2930_v129 = _dafny.Seq.of(_2929_v128, _2929_v128);
          let _2931_v130;
          _2931_v130 = _dafny.Set.fromElements(p1, new BigNumber(((_2930_v129)[_module.__default.safeIndex(p1, new BigNumber((_2930_v129).length))]).length));
          let _2932_v131;
          _2932_v131 = _dafny.Map.Empty.slice().updateUnsafe(_2928_v127,_2929_v128);
          let _2933_v132;
          _2933_v132 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f4),_dafny.Seq.UnicodeFromString("s"));
          let _2934_v133;
          _2934_v133 = _module.D11.create_DC30(_2933_v132);
          let _2935_v134;
          _2935_v134 = _module.D13.create_DC36(_module.D13.create_DC35((_this).f4, (_this).f4, (_this).f6, new BigNumber(547), _2934_v133));
          let _2936_v135;
          _2936_v135 = _module.D13.create_DC36(_2935_v134);
          let _2937_v136;
          _2937_v136 = _dafny.Map.Empty.slice().updateUnsafe((((_2932_v131).contains(_2928_v127)) ? ((_2932_v131).get(_2928_v127)) : (_2931_v130)),_2936_v135);
          _2937_v136 = (_2937_v136).update(_dafny.Set.fromElements(new BigNumber(((_this).f2).length)), _2936_v135);
        }
        let _2938_v137;
        _2938_v137 = new _dafny.CodePoint('l'.codePointAt(0));
        let _2939_v138;
        let _nw442 = new _module.C2();
        _nw442.__ctor((p0).minus(_module.__default.fm0(globalState)), !((_this).f6) || (_this.f3), ((_this).f1).update((_this).f2, _this.f3), _dafny.Seq.Concat((_this).f2, (_this).f2));
        _2939_v138 = _nw442;
        let _rhs500 = (_dafny.ZERO).minus(p1);
        let _rhs501 = _2938_v137;
        let _rhs502 = _2939_v138;
        let _rhs503 = !(((((_this).f1).contains((_this).f2)) ? (((_this).f1).get((_this).f2)) : ((_this).f6))) || ((p0).isLessThan((_2939_v138).f13));
        let _lhs307 = _this;
        r1 = _rhs500;
        _2938_v137 = _rhs501;
        _2939_v138 = _rhs502;
        _lhs307.f3 = _rhs503;
        let _index456 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
        ((_this).f5)[_index456] = (_this).f6;
        let _index457 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
        let _rhs504 = (_this).f6;
        let _rhs505 = (_dafny.ZERO).minus(p1);
        let _rhs506 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wwfasmko"), _dafny.Seq.UnicodeFromString("gak"))).length);
        let _rhs507 = p0;
        let _lhs308 = (_this).f5;
        let _lhs309 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
        _lhs308[_lhs309] = _rhs504;
        r1 = _rhs505;
        r1 = _rhs506;
        r2 = _rhs507;
        let _2940_v139;
        _2940_v139 = _module.D1.create_DC3((_2791_v6).fm13(globalState), new BigNumber((_dafny.Seq.Concat(_module.__default.fm7(true, ((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))], globalState), _2785_v0)).length), (_dafny.ZERO).minus((p0).plus((_2939_v138).f13)));
        let _source44 = _2940_v139;
        if (_source44.is_DC3) {
          let _2941___mcc_h0 = (_source44).cf3;
          let _2942___mcc_h1 = (_source44).cf4;
          let _2943___mcc_h2 = (_source44).cf5;
          let _2944_cf5 = _2943___mcc_h2;
          let _2945_cf4 = _2942___mcc_h1;
          let _2946_cf3 = _2941___mcc_h0;
          let _2947_v140;
          let _init65 = ((_2948_cf3) => function (_2949_i10) {
            return ((_module.D8.create_DC22(_dafny.MultiSet.fromElements(_2948_cf3))).dtor_cf37).IsProperSubsetOf(_dafny.MultiSet.fromElements(_2948_cf3));
          })(_2946_cf3);
          let _nw443 = Array((new BigNumber(16)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw443.length); _i0_65++) {
            _nw443[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2947_v140 = _nw443;
          let _nw444 = Array((new BigNumber(29)).toNumber()).fill(false);
          _2947_v140 = _nw444;
          let _index458 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
          ((_this).f5)[_index458] = _this.f3;
          let _2950_v141;
          let _nw445 = Array((new BigNumber(8)).toNumber());
          _nw445[(_dafny.ZERO).toNumber()] = (_2939_v138).f13;
          _nw445[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))]), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))])).length)), _2946_cf3)).length);
          _nw445[(new BigNumber(2)).toNumber()] = _2945_cf4;
          _nw445[(new BigNumber(3)).toNumber()] = p0;
          _nw445[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber(643));
          _nw445[(new BigNumber(5)).toNumber()] = new BigNumber(((_this).f2).length);
          _nw445[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm6(true, new _dafny.CodePoint('e'.codePointAt(0)), globalState)).length);
          _nw445[(new BigNumber(7)).toNumber()] = p0;
          _2950_v141 = _nw445;
          let _2951_v142;
          let _nw446 = Array((new BigNumber(14)).toNumber());
          _nw446[(_dafny.ZERO).toNumber()] = _2938_v137;
          _nw446[(_dafny.ONE).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(2)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(3)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
          _nw446[(new BigNumber(5)).toNumber()] = _module.__default.fm24(((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))], _this.f3, new BigNumber(((_this).f2).length), globalState);
          _nw446[(new BigNumber(6)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(7)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(8)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(9)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(10)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(11)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(12)).toNumber()] = _2938_v137;
          _nw446[(new BigNumber(13)).toNumber()] = _2938_v137;
          _2951_v142 = _nw446;
          let _2952_v143;
          _2952_v143 = _module.D10.create_DC29(_module.__default.fm3(new BigNumber(-574), globalState), _module.__default.fm3(_2944_cf5, globalState));
          let _2953_v144;
          _2953_v144 = _dafny.Map.Empty.slice().updateUnsafe(_2952_v143,_2951_v142);
          let _2954_v145;
          let _nw447 = Array((new BigNumber(22)).toNumber());
          _nw447[(_dafny.ZERO).toNumber()] = _2951_v142;
          _nw447[(_dafny.ONE).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(2)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(3)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(4)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(5)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(6)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(7)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(8)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(9)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(10)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(11)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(12)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(13)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(14)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(15)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(16)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(17)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(18)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(19)).toNumber()] = (((_2953_v144).contains(_2952_v143)) ? ((_2953_v144).get(_2952_v143)) : (_2951_v142));
          _nw447[(new BigNumber(20)).toNumber()] = _2951_v142;
          _nw447[(new BigNumber(21)).toNumber()] = _2951_v142;
          _2954_v145 = _nw447;
          let _2955_v146;
          _2955_v146 = _2954_v145;
          let _rhs508 = ((_this.f3) ? (_2950_v141) : (_2950_v141));
          let _rhs509 = _2955_v146;
          let _rhs510 = ((_2939_v138).f13).plus((_2944_cf5).multipliedBy(p1));
          _2950_v141 = _rhs508;
          _2955_v146 = _rhs509;
          _2944_cf5 = _rhs510;
          let _2956_v147;
          let _nw448 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2956_v147 = _nw448;
          let _rhs511 = _module.__default.safeDivisionInt((_2791_v6).fm12(p1, p0, _2944_cf5, (_this).f6, globalState), new BigNumber(381));
          let _rhs512 = _2956_v147;
          let _rhs513 = _2956_v147;
          let _rhs514 = !(_2786_v1).contains(true);
          let _lhs310 = _this;
          _2944_cf5 = _rhs511;
          _2956_v147 = _rhs512;
          _2956_v147 = _rhs513;
          _lhs310.f3 = _rhs514;
        } else if (_source44.is_DC4) {
          let _2957___mcc_h3 = (_source44).cf6;
          let _2958___mcc_h4 = (_source44).cf7;
          let _2959_cf7 = _2958___mcc_h4;
          let _2960_cf6 = _2957___mcc_h3;
          let _2961_v148;
          _2961_v148 = _dafny.Set.fromElements(((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))], _2959_cf7);
          let _2962_v149;
          _2962_v149 = _dafny.MultiSet.fromElements(_2961_v148, _2961_v148);
          let _2963_v150;
          _2963_v150 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2962_v149).Difference(_2962_v149)).cardinality()),(_2959_cf7) === (((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))]));
          let _2964_v151;
          _2964_v151 = _dafny.Set.fromElements(new BigNumber(835), (_2939_v138).f13);
          _2963_v150 = (_2963_v150).update((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_this).f2).length), (_2939_v138).f13)), (_2964_v151).IsProperSubsetOf(_2964_v151));
          let _2965_v152;
          let _init66 = ((_2966_v137) => function (_2967_i11) {
            return _module.__default.safeDivisionInt(_2967_i11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-637)), ((_2968_v137) => function (_2969_i12) {
              return _2968_v137;
            })(_2966_v137))).length));
          })(_2938_v137);
          let _nw449 = Array((new BigNumber(9)).toNumber());
          for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw449.length); _i0_66++) {
            _nw449[_i0_66] = _init66(new BigNumber(_i0_66));
          }
          _2965_v152 = _nw449;
          let _index459 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_2965_v152).length));
          (_2965_v152)[_index459] = p0;
          let _2970_v153;
          let _nw450 = Array((new BigNumber(18)).toNumber()).fill(false);
          _2970_v153 = _nw450;
          let _2971_v154;
          let _nw451 = new _module.C3();
          _nw451.__ctor(p1, _2970_v153, ((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))], (_this).f1, (_this).f2);
          _2971_v154 = _nw451;
          let _2972_v155;
          _2972_v155 = _dafny.Seq.of(_2971_v154);
          let _2973_v156;
          _2973_v156 = _dafny.Seq.of((_2972_v155)[_module.__default.safeIndex((_this).f4, new BigNumber((_2972_v155).length))]);
          let _2974_v157;
          _2974_v157 = _dafny.Map.Empty.slice().updateUnsafe(_2973_v156,p0);
          let _index460 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_2965_v152).length));
          (_2965_v152)[_index460] = new BigNumber((_2974_v157).length);
          _2959_cf7 = _2959_cf7;
          r1 = (((_2965_v152)[_module.__default.safeIndex(new BigNumber(452), new BigNumber((_2965_v152).length))]).multipliedBy((_this).f4)).plus(((_2939_v138).f13).minus((_dafny.ZERO).minus((_this).f4)));
        } else {
          let _2975___mcc_h5 = (_source44).cf2;
          let _2976_cf2 = _2975___mcc_h5;
          let _2977_v158;
          let _nw452 = Array((new BigNumber(19)).toNumber()).fill(false);
          _2977_v158 = _nw452;
          let _2978_v159;
          let _init67 = ((_2979_cf2) => function (_2980_i13) {
            return _2979_cf2;
          })(_2976_cf2);
          let _nw453 = Array((new BigNumber(8)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw453.length); _i0_67++) {
            _nw453[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2978_v159 = _nw453;
          let _2981_v160;
          _2981_v160 = _dafny.Seq.of(_2977_v158, _2977_v158, _2977_v158);
          let _2982_v161;
          _2982_v161 = _dafny.Map.Empty.slice().updateUnsafe((_2939_v138).f13,_this.f3);
          let _index461 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
          let _rhs515 = (_2981_v160)[_module.__default.safeIndex(p1, new BigNumber((_2981_v160).length))];
          let _rhs516 = _2978_v159;
          let _rhs517 = !((_2982_v161).Merge(_2982_v161)).contains((_2939_v138).f13);
          let _lhs311 = (_this).f5;
          let _lhs312 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
          _2977_v158 = _rhs515;
          _2978_v159 = _rhs516;
          _lhs311[_lhs312] = _rhs517;
          let _2983_v162;
          let _nw454 = new _module.C1();
          _nw454.__ctor();
          _2983_v162 = _nw454;
          let _2984_v163;
          let _nw455 = Array((new BigNumber(25)).toNumber()).fill(_module.D13.Default());
          _2984_v163 = _nw455;
          let _2985_v164;
          let _nw456 = Array((new BigNumber(27)).toNumber()).fill([]);
          _2985_v164 = _nw456;
          let _2986_v165;
          _2986_v165 = _module.D22.create_DC54(_2984_v163);
          let _index462 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
          let _rhs518 = ((_this).f4).plus(p0);
          let _rhs519 = ((_this).f6) === (((_2939_v138).f13).isLessThan(p1));
          let _rhs520 = (_2986_v165).dtor_cf85;
          let _rhs521 = _2985_v164;
          let _lhs313 = (_this).f5;
          let _lhs314 = _module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length));
          r1 = _rhs518;
          _lhs313[_lhs314] = _rhs519;
          _2984_v163 = _rhs520;
          _2985_v164 = _rhs521;
          let _2987_v166;
          let _init68 = ((_2988_v138) => function (_2989_i14) {
            return _module.__default.safeDivisionInt(_2989_i14, (_2988_v138).f13);
          })(_2939_v138);
          let _nw457 = Array((new BigNumber(5)).toNumber());
          for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw457.length); _i0_68++) {
            _nw457[_i0_68] = _init68(new BigNumber(_i0_68));
          }
          _2987_v166 = _nw457;
          let _index463 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_2987_v166).length));
          (_2987_v166)[_index463] = new BigNumber(965);
          let _2990_v167;
          _2990_v167 = _module.D11.create_DC30(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(731),_2976_cf2));
          let _2991_v168;
          _2991_v168 = _module.D13.create_DC35(new BigNumber((_2982_v161).length), (_dafny.ZERO).minus((_2939_v138).f13), false, (_2939_v138).f13, _2990_v167);
          let _2992_v169;
          _2992_v169 = _module.D18.create_DC44(_module.__default.fm64((_this).f6, (_this).f6, (_2991_v168).dtor_cf59, globalState));
          let _index464 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_2987_v166).length));
          let _rhs522 = _module.__default.safeDivisionInt((_2939_v138).f13, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f6)).length));
          let _rhs523 = p1;
          let _rhs524 = _2992_v169;
          let _lhs315 = _2987_v166;
          let _lhs316 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_2987_v166).length));
          _lhs315[_lhs316] = _rhs522;
          r2 = _rhs523;
          _2992_v169 = _rhs524;
        }
        (_this).f3 = ((_this).f5)[_module.__default.safeIndex(new BigNumber(946), new BigNumber(((_this).f5).length))];
      }
      let _hi12 = (_this).f4;
      for (let _2993_i15 = new BigNumber((_dafny.Seq.of((_this).f2)).length); _2993_i15.isLessThan(_hi12); _2993_i15 = _2993_i15.plus(_dafny.ONE)) {
        r2 = ((!(new BigNumber((_2786_v1).length)).isEqualTo(_2993_i15)) ? (((_this).f4).multipliedBy((_this).f4)) : (_2993_i15));
        let _2994_v170;
        let _init69 = function (_2995_i16) {
          return (_2995_i16).plus(new BigNumber(((_this).f2).length));
        };
        let _nw458 = Array((new BigNumber(13)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw458.length); _i0_69++) {
          _nw458[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2994_v170 = _nw458;
        let _2996_v171;
        _2996_v171 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(752),_2994_v170);
        let _2997_v172;
        _2997_v172 = _dafny.Seq.of((_this).f4);
        let _2998_v173;
        _2998_v173 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(-329)),_2994_v170);
        let _2999_v174;
        _2999_v174 = _dafny.Map.Empty.slice().updateUnsafe((((_2996_v171).contains((_2997_v172)[_module.__default.safeIndex(_2993_i15, new BigNumber((_2997_v172).length))])) ? ((_2996_v171).get((_2997_v172)[_module.__default.safeIndex(_2993_i15, new BigNumber((_2997_v172).length))])) : ((((_2998_v173).contains(_module.__default.fm20(false, (_this).f6, (_this).f6, _2785_v0, globalState))) ? ((_2998_v173).get(_module.__default.fm20(false, (_this).f6, (_this).f6, _2785_v0, globalState))) : (_2994_v170)))),(_this).f5);
        let _3000_v175;
        _3000_v175 = _dafny.Set.fromElements(p0, _2993_i15, new BigNumber(220), (_this).f4);
        let _3001_v176;
        let _nw459 = new _module.C9();
        _nw459.__ctor((_this).f5, _2999_v174, new BigNumber(723), (_this).f5, (_this).f6, ((_this).f1).update((_this).f2, (_2785_v0)[_module.__default.safeIndex(new BigNumber((_3000_v175).length), new BigNumber((_2785_v0).length))]), _dafny.Seq.Concat((_this).f2, (_2791_v6).fm16(_2993_i15, (_this).f4, p1, globalState)));
        _3001_v176 = _nw459;
        let _3002_v177;
        let _nw460 = new _module.C5();
        _nw460.__ctor((_this).f6, ((_this).f1).Merge(_module.__default.fm62(p0, !(_this.f3), globalState)), (_this).f2);
        _3002_v177 = _nw460;
        _3002_v177 = _3002_v177;
        let _index465 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_2994_v170).length));
        (_2994_v170)[_index465] = p1;
        let _index466 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_2994_v170).length));
        (_2994_v170)[_index466] = (_3001_v176).fm12(_2993_i15, p0, _module.__default.safeModuloInt(new BigNumber(71), p1), (_module.__default.fm33(p0, globalState)).IsProperSubsetOf(_3000_v175), globalState);
      }
      let _3003_v178;
      _3003_v178 = new _dafny.CodePoint('w'.codePointAt(0));
      let _3004_v179;
      _3004_v179 = _module.D8.create_DC24(_3003_v178, _dafny.Map.Empty.slice().updateUnsafe((_this).fm13(globalState),(_this).f6));
      r0 = (_3004_v179).dtor_cf42;
      r1 = _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_3005_v178) => function (_3006_i17) {
        return _3005_v178;
      })(_3003_v178))).length));
      r2 = (_this).f4;
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _3007_v0;
      let _init70 = function (_3008_i0) {
        return (_3008_i0).minus(new BigNumber((_dafny.Set.fromElements((_this).f6)).length));
      };
      let _nw461 = Array((new BigNumber(17)).toNumber());
      for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw461.length); _i0_70++) {
        _nw461[_i0_70] = _init70(new BigNumber(_i0_70));
      }
      _3007_v0 = _nw461;
      let _index467 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
      (_3007_v0)[_index467] = ((_this).f4).plus((_this).f4);
      let _index468 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
      (_3007_v0)[_index468] = ((_this).f4).plus((_this).f4);
      let _3009_v1;
      _3009_v1 = _module.D16.create_DC42((_this).f6, false, (_this).f6);
      if ((_3009_v1).dtor_cf70) {
        let _3010_v2;
        _3010_v2 = _dafny.MultiSet.fromElements(_this.f3);
        _3010_v2 = (_3010_v2).update((_this).f6, _module.__default.abs((_this).f4));
        let _3011_v3;
        _3011_v3 = _dafny.Seq.of((_this).f4, (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))]);
        let _3012_v4;
        _3012_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_3011_v3);
        let _3013_v5;
        _3013_v5 = _dafny.Map.Empty.slice().updateUnsafe((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))],_3012_v4);
        let _index469 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        (_3007_v0)[_index469] = new BigNumber((((((_3013_v5).contains((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))])) ? ((_3013_v5).get((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))])) : (_3012_v4))).Merge(_3012_v4)).length);
        let _3014_v6;
        _3014_v6 = _module.D22.create_DC55((_this).f2, (_this).f4, _3007_v0);
        (_this).f3 = ((((_this).f1).contains((_3014_v6).dtor_cf86)) ? (((_this).f1).get((_3014_v6).dtor_cf86)) : (_this.f3));
        (_this).f3 = ((_this).f6) && ((_this).f6);
        let _3015_v7;
        _3015_v7 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f4).multipliedBy(new BigNumber(((_this).f2).length)),(_this).f2);
        _3015_v7 = (_3015_v7).update(new BigNumber(-210), _dafny.Seq.UnicodeFromString("feihncdx"));
      } else {
        let _3016_v8;
        _3016_v8 = new _dafny.CodePoint('r'.codePointAt(0));
        let _3017_v9;
        _3017_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,((true) ? ((_this).f2) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("mfkrh"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_dafny.Seq.UnicodeFromString("mfkrh")).length)), _3016_v8))));
        let _3018_v10;
        _3018_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f2);
        _3017_v9 = (_3017_v9).update(_dafny.Seq.Concat((_this).f2, (((_3018_v10).contains((_this).f4)) ? ((_3018_v10).get((_this).f4)) : ((_this).f2))), (_this).f2);
        let _3019_v11;
        _3019_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        let _3020_v12;
        let _nw462 = new _module.C4();
        _nw462.__ctor((((_3019_v11).contains((_this).f6)) ? ((_3019_v11).get((_this).f6)) : (_this.f3)), _module.__default.fm62((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))], (_this).f6, globalState), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uckdk"), _dafny.Seq.Create(_module.__default.abs(_dafny.ONE), ((_3021_v8) => function (_3022_i1) {
          return _3021_v8;
        })(_3016_v8))));
        _3020_v12 = _nw462;
        let _3023_v13;
        _3023_v13 = _dafny.Seq.of((_this).f4, (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))]);
        let _3024_v14;
        _3024_v14 = _dafny.Seq.of(_3023_v13);
        let _3025_v15;
        _3025_v15 = _dafny.Seq.of(((true) ? ((_this).fm13(globalState)) : ((_this).f6)));
        let _index470 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _rhs525 = (_3023_v13)[_module.__default.safeIndex(new BigNumber((_3019_v11).length), new BigNumber((_3023_v13).length))];
        let _rhs526 = (_3025_v15)[_module.__default.safeIndex((_this).f4, new BigNumber((_3025_v15).length))];
        let _rhs527 = _this.f3;
        let _rhs528 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-793)), ((_3026_v13) => function (_3027_i2) {
          return _3026_v13;
        })(_3023_v13));
        let _lhs317 = _3007_v0;
        let _lhs318 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _lhs319 = _this;
        let _lhs320 = _this;
        _lhs317[_lhs318] = _rhs525;
        _lhs319.f3 = _rhs526;
        _lhs320.f3 = _rhs527;
        _3024_v14 = _rhs528;
        let _3028_v16;
        _3028_v16 = _module.__default.fm33(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))]), _module.__default.safeIndex((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))], new BigNumber((_dafny.Seq.of((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))])).length)), (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))])).length), globalState);
        let _3029_v17;
        _3029_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_3024_v14)[_module.__default.safeIndex((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))], new BigNumber((_3024_v14).length))]).length),(_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))]);
        let _3030_v18;
        _3030_v18 = _dafny.Set.fromElements((((_3029_v17).contains(new BigNumber(-718))) ? ((_3029_v17).get(new BigNumber(-718))) : (new BigNumber(((_this).f2).length))));
        _3028_v16 = _3030_v18;
        let _index471 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        (_3007_v0)[_index471] = (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))];
      }
      let _3031_v19;
      _3031_v19 = _dafny.Map.Empty.slice().updateUnsafe(_3007_v0,(_this).f5);
      let _3032_v20;
      let _nw463 = new _module.C9();
      _nw463.__ctor((_this).f5, _3031_v19, new BigNumber(762), (_this).f5, _this.f3, (_this).f1, (_this).f2);
      _3032_v20 = _nw463;
      let _3033_v21;
      _3033_v21 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f6),_3032_v20);
      let _3034_v22;
      _3034_v22 = _dafny.Seq.of((_this).f2);
      let _3035_v23;
      _3035_v23 = _module.D23.create_DC57(_3032_v20);
      _3033_v21 = (_3033_v21).update(_dafny.Seq.IsPrefixOf((_3034_v22)[_module.__default.safeIndex(new BigNumber(-462), new BigNumber((_3034_v22).length))], (_this).f2), (_3035_v23).dtor_cf90);
      let _hi13 = (_this).f4;
      for (let _3036_i3 = (_this).f4; _3036_i3.isLessThan(_hi13); _3036_i3 = _3036_i3.plus(_dafny.ONE)) {
        let _index472 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        (_3007_v0)[_index472] = ((_this).f4).plus(_module.__default.safeModuloInt(new BigNumber(683), _3036_i3));
        let _3037_v24;
        let _nw464 = new _module.C2();
        _nw464.__ctor((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))], !(_module.__default.fm3((_this).f4, globalState)), (_this).f1, (_this).f2);
        _3037_v24 = _nw464;
        let _3038_v25;
        _3038_v25 = _3037_v24;
        let _index473 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _index474 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _index475 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _rhs529 = (_dafny.ZERO).minus(_3036_i3);
        let _rhs530 = ((_dafny.ZERO).minus((_this).f4)).multipliedBy(_module.__default.safeModuloInt((_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))], (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))]));
        let _rhs531 = _3007_v0;
        let _rhs532 = (_3007_v0)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length))];
        let _rhs533 = (_3038_v25);
        let _lhs321 = _3007_v0;
        let _lhs322 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _lhs323 = _3007_v0;
        let _lhs324 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        let _lhs325 = _3007_v0;
        let _lhs326 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        _lhs321[_lhs322] = _rhs529;
        _lhs323[_lhs324] = _rhs530;
        _3007_v0 = _rhs531;
        _lhs325[_lhs326] = _rhs532;
        _3037_v24 = _rhs533;
        let _3039_v27;
        let _nw465 = new _module.C3();
        _nw465.__ctor((_this).f4, (_this).f5, false, function () {
          let _coll90 = new _dafny.Map();
          for (const _compr_90 of (_dafny.MultiSet.fromElements((_this).f2, (_this).f2, (_this).f2, (_3037_v24).f2)).Elements) {
            let _3040_v26 = _compr_90;
            if ((_dafny.MultiSet.fromElements((_this).f2, (_this).f2, (_this).f2, (_3037_v24).f2)).contains(_3040_v26)) {
              _coll90.push([_3040_v26,(_this).f6]);
            }
          }
          return _coll90;
        }(), (_this).f2);
        _3039_v27 = _nw465;
        let _3041_v28;
        _3041_v28 = _dafny.Seq.of(_3039_v27);
        let _3042_v29;
        _3042_v29 = new _dafny.CodePoint('l'.codePointAt(0));
        let _index476 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_3007_v0).length));
        (_3007_v0)[_index476] = (new BigNumber((_3041_v28).length)).plus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_3042_v29,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),(_this).f6))).length));
        let _3043_v30;
        let _nw466 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _3043_v30 = _nw466;
        let _index477 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_3043_v30).length));
        (_3043_v30)[_index477] = _3042_v29;
        let _index478 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_3043_v30).length));
        (_3043_v30)[_index478] = (((_this).f6) ? (_3042_v29) : (_3042_v29));
      }
      let _3044_v31;
      let _nw467 = Array((new BigNumber(6)).toNumber());
      _nw467[(_dafny.ZERO).toNumber()] = (_this).f2;
      _nw467[(_dafny.ONE).toNumber()] = _module.__default.fm2((_this).f6, globalState);
      _nw467[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat((_this).f2, (_this).f2);
      _nw467[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(227)), function (_3045_i4) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), (_this).f2);
      _nw467[(new BigNumber(4)).toNumber()] = (_this).f2;
      _nw467[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_this).f2, (_this).f2);
      _3044_v31 = _nw467;
      let _index479 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_3044_v31).length));
      (_3044_v31)[_index479] = (_this).f2;
      let _index480 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_3044_v31).length));
      (_3044_v31)[_index480] = (_this).f2;
      (_this).f3 = !((_this).f4).isEqualTo(_module.__default.safeModuloInt((_this).f4, (_this).f4));
      let _3046_v32;
      _3046_v32 = _module.D10.create_DC29((_this).f6, !(_this.f3));
      let _pat_let_tv57 = _3044_v31;
      let _pat_let_tv58 = _3044_v31;
      let _pat_let_tv59 = _3044_v31;
      let _pat_let_tv60 = _3044_v31;
      let _pat_let_tv61 = _3007_v0;
      let _pat_let_tv62 = _3007_v0;
      r0 = function (_source45) {
        if (_source45.is_DC29) {
          let _3047___mcc_h0 = (_source45).cf47;
          let _3048___mcc_h1 = (_source45).cf48;
          let _3049_cf48 = _3048___mcc_h1;
          let _3050_cf47 = _3047___mcc_h0;
          return function () {
            let _coll91 = new _dafny.Map();
            for (const _compr_91 of (_dafny.Seq.of(new BigNumber(((_pat_let_tv58)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_pat_let_tv57).length))]).length))).Elements) {
              let _3051_v33 = _compr_91;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(((_pat_let_tv60)[_module.__default.safeIndex(new BigNumber(225), new BigNumber((_pat_let_tv59).length))]).length)), _3051_v33)) {
                _coll91.push([(_3051_v33).minus(new BigNumber((_dafny.Seq.of(_3049_cf48)).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_3050_cf47,true)).length)]);
              }
            }
            return _coll91;
          }();
        } else {
          let _3052___mcc_h2 = (_source45).cf46;
          let _3053_cf46 = _3052___mcc_h2;
          return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(240),_this.f3)).length),(_pat_let_tv62)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_pat_let_tv61).length))])).Merge(_3053_cf46);
        }
      }(_3046_v32);
      return r0;
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
