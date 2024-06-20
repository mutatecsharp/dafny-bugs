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
    static fm0(p0, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("cfbosyfj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_0_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("w"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("u"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), function (_1_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), function (_2_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-981)),!(false))).length),_module.D1.create_DC3(_dafny.Seq.of(!(true))))).length))).cardinality())).multipliedBy(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Set.fromElements(function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
              let _coll3 = new _dafny.Map();
              for (const _compr_3 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                let _3_v1 = _compr_3;
                if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                  _coll3.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                }
              }
              return _coll3;
            }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).Keys.Elements) {
              let _4_v2 = _compr_2;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                let _coll4 = new _dafny.Map();
                for (const _compr_4 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                  let _3_v1 = _compr_4;
                  if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                    _coll4.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                  }
                }
                return _coll4;
              }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).contains(_4_v2)) {
                _coll2.add((_4_v2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_5_i1) {
                  return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_6_i2) {
                    return new BigNumber((_dafny.Seq.of(false)).length);
                  })).length));
                })).length)));
              }
            }
            return _coll2;
          }())).Elements) {
            let _7_v0 = _compr_1;
            if ((_dafny.Set.fromElements(function () {
              let _coll5 = new _dafny.Set();
              for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                let _coll6 = new _dafny.Map();
                for (const _compr_6 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                  let _3_v1 = _compr_6;
                  if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                    _coll6.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                  }
                }
                return _coll6;
              }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).Keys.Elements) {
                let _8_v2 = _compr_5;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                  let _coll7 = new _dafny.Map();
                  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                    let _3_v1 = _compr_7;
                    if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                      _coll7.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                    }
                  }
                  return _coll7;
                }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).contains(_8_v2)) {
                  _coll5.add((_8_v2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_5_i1) {
                    return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_6_i2) {
                      return new BigNumber((_dafny.Seq.of(false)).length);
                    })).length));
                  })).length)));
                }
              }
              return _coll5;
            }())).contains(_7_v0)) {
              _coll1.push([_7_v0,new BigNumber(90)]);
            }
          }
          return _coll1;
        }()).length),!(true))).Keys.Elements) {
          let _9_v3 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_dafny.Set.fromElements(function () {
              let _coll9 = new _dafny.Set();
              for (const _compr_9 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                let _coll10 = new _dafny.Map();
                for (const _compr_10 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                  let _3_v1 = _compr_10;
                  if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                    _coll10.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                  }
                }
                return _coll10;
              }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).Keys.Elements) {
                let _10_v2 = _compr_9;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                  let _coll11 = new _dafny.Map();
                  for (const _compr_11 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                    let _3_v1 = _compr_11;
                    if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                      _coll11.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                    }
                  }
                  return _coll11;
                }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).contains(_10_v2)) {
                  _coll9.add((_10_v2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_5_i1) {
                    return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_6_i2) {
                      return new BigNumber((_dafny.Seq.of(false)).length);
                    })).length));
                  })).length)));
                }
              }
              return _coll9;
            }())).Elements) {
              let _7_v0 = _compr_8;
              if ((_dafny.Set.fromElements(function () {
                let _coll12 = new _dafny.Set();
                for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                  let _coll13 = new _dafny.Map();
                  for (const _compr_13 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                    let _3_v1 = _compr_13;
                    if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                      _coll13.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                    }
                  }
                  return _coll13;
                }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).Keys.Elements) {
                  let _11_v2 = _compr_12;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
                    let _coll14 = new _dafny.Map();
                    for (const _compr_14 of _dafny.IntegerRange(new BigNumber(280), new BigNumber(120))) {
                      let _3_v1 = _compr_14;
                      if (((new BigNumber(280)).isLessThanOrEqualTo(_3_v1)) && ((_3_v1).isLessThan(new BigNumber(120)))) {
                        _coll14.push([(_3_v1).multipliedBy(new BigNumber(-118)),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cth")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(340),new BigNumber(90))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true)).length),new BigNumber(622)))).length)]);
                      }
                    }
                    return _coll14;
                  }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-625))).length))).length)))).cardinality()),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-234), new BigNumber((_dafny.Seq.UnicodeFromString("qpmel")).length))).length))).cardinality()))).contains(_11_v2)) {
                    _coll12.add((_11_v2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), function (_5_i1) {
                      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_6_i2) {
                        return new BigNumber((_dafny.Seq.of(false)).length);
                      })).length));
                    })).length)));
                  }
                }
                return _coll12;
              }())).contains(_7_v0)) {
                _coll8.push([_7_v0,new BigNumber(90)]);
              }
            }
            return _coll8;
          }()).length),!(true))).contains(_9_v3)) {
            _coll0.add((_9_v3).plus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-920))).length)));
          }
        }
        return _coll0;
      }()).length)),(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll15 = new _dafny.Set();
        for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_12_i3) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length),new BigNumber((_dafny.Set.fromElements(false)).length))).Keys.Elements) {
          let _13_v4 = _compr_15;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), function (_12_i3) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          })).length),new BigNumber((_dafny.Set.fromElements(false)).length))).contains(_13_v4)) {
            _coll15.add((_13_v4).multipliedBy(new BigNumber(438)));
          }
        }
        return _coll15;
      }()).length),new BigNumber((_dafny.Seq.UnicodeFromString("jiwdrv")).length))).length)).minus(new BigNumber(630)));
    };
    static fm9(p0, p1, p2, globalState) {
      let _source0 = _dafny.Set.fromElements(false);
      let _14___mcc_h0 = _source0;
      let _15_cf23 = _14___mcc_h0;
      return (_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)));
    };
    static fm10(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-621))))).Difference((_dafny.MultiSet.fromElements(new BigNumber(545))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(352), new BigNumber((_dafny.Seq.of(true)).length)))));
    };
    static fm11(p0, p1, p2, globalState) {
      let _source1 = _module.D14.create_DC32(_dafny.MultiSet.fromElements(!(true)), true, false, !(true), new BigNumber(542));
      if (_source1.is_DC32) {
        let _16___mcc_h0 = (_source1).cf56;
        let _17___mcc_h1 = (_source1).cf57;
        let _18___mcc_h2 = (_source1).cf58;
        let _19___mcc_h3 = (_source1).cf59;
        let _20___mcc_h4 = (_source1).cf60;
        let _21_cf60 = _20___mcc_h4;
        let _22_cf59 = _19___mcc_h3;
        let _23_cf58 = _18___mcc_h2;
        let _24_cf57 = _17___mcc_h1;
        let _25_cf56 = _16___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_22_cf59), _dafny.Seq.of(_23_cf58));
      } else {
        let _26___mcc_h5 = (_source1).cf55;
        let _27_cf55 = _26___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.of(!(true), true), _dafny.Seq.of(true));
      }
    };
    static fm14(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('d'.codePointAt(0));
    };
    static fm15(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(835)), function (_28_i0) {
        return new BigNumber((_dafny.Set.fromElements(false)).length);
      }), _dafny.Seq.of(new BigNumber(215))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), function (_29_i1) {
        return new BigNumber(600);
      }));
    };
    static fm18(p0, globalState) {
      return _dafny.Set.fromElements(!(false) || (!(false)));
    };
    static fm19(p0, p1, globalState) {
      return _module.D0.create_DC0((new BigNumber(406)).isEqualTo(new BigNumber(-130)));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC9(new BigNumber(-36), new BigNumber(441), (false) && (false));
    };
    static fm21(p0, globalState) {
      return (_dafny.MultiSet.fromElements(false, false, !(true))).Intersect(_dafny.MultiSet.fromElements(false));
    };
    static fm22(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("vvtodqg");
    };
    static fm23(p0, globalState) {
      let _source2 = _module.D1.create_DC5(new BigNumber(796));
      if (_source2.is_DC4) {
        let _30___mcc_h0 = (_source2).cf7;
        let _31___mcc_h1 = (_source2).cf8;
        let _32___mcc_h2 = (_source2).cf9;
        let _33___mcc_h3 = (_source2).cf10;
        let _34___mcc_h4 = (_source2).cf11;
        let _35_cf11 = _34___mcc_h4;
        let _36_cf10 = _33___mcc_h3;
        let _37_cf9 = _32___mcc_h2;
        let _38_cf8 = _31___mcc_h1;
        let _39_cf7 = _30___mcc_h0;
        if (!(true)) {
          return _dafny.Seq.of(_38_cf8);
        } else {
          return _dafny.Seq.of(new BigNumber((function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of _dafny.IntegerRange(new BigNumber(-531), new BigNumber(718))) {
              let _40_v0 = _compr_16;
              if (((new BigNumber(-531)).isLessThanOrEqualTo(_40_v0)) && ((_40_v0).isLessThan(new BigNumber(718)))) {
                _coll16.push([(_40_v0).multipliedBy(_37_cf9),true]);
              }
            }
            return _coll16;
          }()).length));
        }
      } else if (_source2.is_DC5) {
        let _41___mcc_h5 = (_source2).cf12;
        let _42_cf12 = _41___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.of(_42_cf12, new BigNumber(584), (_dafny.ZERO).minus(_42_cf12), _42_cf12), _dafny.Seq.of(_42_cf12));
      } else {
        let _43___mcc_h6 = (_source2).cf6;
        let _44_cf6 = _43___mcc_h6;
        return _dafny.Seq.of(new BigNumber(107));
      }
    };
    static fm24(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('v'.codePointAt(0));
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(new BigNumber(244), new BigNumber(580)), _dafny.Seq.of(new BigNumber(-662), new BigNumber((_dafny.Seq.of(!(false), true, true, true)).length))), true);
    };
    static fm26(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0)))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(327),true)).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jsiwka")).length),new BigNumber((_dafny.Seq.UnicodeFromString("vjgqe")).length)))).Intersect((_dafny.MultiSet.FromArray(_dafny.Seq.of(function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(334),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(false, true)).length))).length))).Keys.Elements) {
          let _45_v0 = _compr_17;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(334),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(false, true)).length))).length))).contains(_45_v0)) {
            _coll17.push([_module.__default.safeModuloInt(_45_v0, new BigNumber((_dafny.Set.fromElements(true)).length)),new BigNumber(969)]);
          }
        }
        return _coll17;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(243),new BigNumber((_dafny.MultiSet.fromElements((_module.D22.create_DC54(new _dafny.CodePoint('u'.codePointAt(0)), _module.D9.create_DC22(false), true, true, true)).dtor_cf98, new _dafny.CodePoint('x'.codePointAt(0)))).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-42), new BigNumber(599))) {
          let _46_v1 = _compr_18;
          if (((new BigNumber(-42)).isLessThanOrEqualTo(_46_v1)) && ((_46_v1).isLessThan(new BigNumber(599)))) {
            _coll18.push([_module.__default.safeModuloInt(_46_v1, _dafny.ZERO),new BigNumber(335)]);
          }
        }
        return _coll18;
      }()).length))).length))).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(48),new BigNumber(789))))).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(162),new BigNumber((_dafny.Seq.UnicodeFromString("bwfgndon")).length)))));
    };
    static fm27(p0, p1, globalState) {
      let _source3 = _module.D22.create_DC54(new _dafny.CodePoint('r'.codePointAt(0)), _module.D9.create_DC22(true), false, true, !(false));
      if (_source3.is_DC54) {
        let _47___mcc_h0 = (_source3).cf98;
        let _48___mcc_h1 = (_source3).cf99;
        let _49___mcc_h2 = (_source3).cf100;
        let _50___mcc_h3 = (_source3).cf101;
        let _51___mcc_h4 = (_source3).cf102;
        let _52_cf102 = _51___mcc_h4;
        let _53_cf101 = _50___mcc_h3;
        let _54_cf100 = _49___mcc_h2;
        let _55_cf99 = _48___mcc_h1;
        let _56_cf98 = _47___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rqm"),_54_cf100);
      } else {
        let _57___mcc_h5 = (_source3).cf97;
        let _58_cf97 = _57___mcc_h5;
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("smnurn"),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), function (_59_i0) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }),!(true)));
      }
    };
    static fm28(globalState) {
      return (function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_dafny.Seq.of(_module.D0.create_DC1(true, true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).length), !(true)), _module.D0.create_DC1(!(false), !(false), new BigNumber(297), true), _module.D0.create_DC1(false, !(false), new BigNumber(106), true))).Elements) {
          let _60_v0 = _compr_19;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC1(true, true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).length), !(true)), _module.D0.create_DC1(!(false), !(false), new BigNumber(297), true), _module.D0.create_DC1(false, !(false), new BigNumber(106), true)), _60_v0)) {
            _coll19.add(_60_v0);
          }
        }
        return _coll19;
      }()).Difference(_dafny.Set.fromElements(_module.D0.create_DC1(false, !(false), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), false)));
    };
    static fm29(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),new BigNumber((_dafny.Seq.of(false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(667),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jp")).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)))).length),new BigNumber(((_module.D10.create_DC24(_dafny.Seq.UnicodeFromString("lwtjp"), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(227))), true)).dtor_cf44).length)));
    };
    static fm30(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.UnicodeFromString("ojauypo")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-979)));
    };
    static fm31(p0, globalState) {
      return (_module.D21.create_DC49(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(809), new BigNumber(745), new BigNumber((_dafny.Seq.UnicodeFromString("mf")).length))).length),new BigNumber(326))))).dtor_cf86;
    };
    static fm32(p0, p1, globalState) {
      if ((!(false)) && (!(false))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(true, true)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(true));
      }
    };
    static fm33(p0, p1, globalState) {
      return _module.D2.create_DC8();
    };
    static fm34(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ovj")).length), new BigNumber(57)),_dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber(362)))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length)),_dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("vwg")).length),new BigNumber(725))))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_61_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("uxbg")).length);
      }),_dafny.Seq.Create(_module.__default.abs(new BigNumber(364)), function (_62_i1) {
        return _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jlhmrsu")).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length))).length)));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(992))).length),new BigNumber(-660))).length))).length),true)).length),false)).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-63)), function (_63_i2) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length))).length)),new BigNumber((_dafny.Seq.of(true)).length))).length)),_dafny.Seq.of(_module.D2.create_DC6(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_dafny.Set.fromElements(new BigNumber(283))).Elements) {
    let _64_v0 = _compr_20;
    if ((_dafny.Set.fromElements(new BigNumber(283))).contains(_64_v0)) {
      _coll20.push([(_64_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(980)), function (_65_i3) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length))),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-254)))]);
    }
  }
  return _coll20;
}())))));
    };
    static fm35(p0, globalState) {
      if (true) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), function (_66_i0) {
          return _module.D5.create_DC14((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true, true))).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(161)), function (_67_i1) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length);
})).length))).length), true);
        });
      } else {
        return _dafny.Seq.of(_module.D5.create_DC14(new BigNumber(219), new BigNumber((_dafny.Set.fromElements(new BigNumber(274))).length), false));
      }
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-68),new BigNumber(-480)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-255)), function (_68_i0) {
        return _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-248)),new BigNumber(436)));
      }), _dafny.Seq.Concat(_dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(282),new BigNumber(895))), _module.D2.create_DC6(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of _dafny.IntegerRange(new BigNumber(691), new BigNumber(981))) {
    let _69_v0 = _compr_21;
    if (((new BigNumber(691)).isLessThanOrEqualTo(_69_v0)) && ((_69_v0).isLessThan(new BigNumber(981)))) {
      _coll21.push([(_69_v0).minus(new BigNumber(202)),new BigNumber(-438)]);
    }
  }
  return _coll21;
}()), _module.D2.create_DC6(function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of _dafny.IntegerRange(new BigNumber(-193), new BigNumber(618))) {
    let _70_v1 = _compr_22;
    if (((new BigNumber(-193)).isLessThanOrEqualTo(_70_v1)) && ((_70_v1).isLessThan(new BigNumber(618)))) {
      _coll22.push([_module.__default.safeDivisionInt(_70_v1, new BigNumber(560)),new BigNumber((_dafny.Seq.UnicodeFromString("ao")).length)]);
    }
  }
  return _coll22;
}())), _dafny.Seq.of(_module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-649))).cardinality()),new BigNumber(384))))));
    };
    static fm37(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("sqdccsvha")),new _dafny.CodePoint('h'.codePointAt(0)))).Merge((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("nrjkdbhoe")),false)).Keys.Elements) {
          let _71_v0 = _compr_23;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("nrjkdbhoe")),false)).contains(_71_v0)) {
            _coll23.push([_71_v0,new _dafny.CodePoint('n'.codePointAt(0))]);
          }
        }
        return _coll23;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("mxdmgbo")),new _dafny.CodePoint('d'.codePointAt(0)))));
    };
    static fm38(p0, p1, p2, globalState) {
      return _module.D6.create_DC16(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_72_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(135)), function (_73_i1) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})));
    };
    static fm39(globalState) {
      return new BigNumber(479);
    };
    static fm40(p0, p1, p2, globalState) {
      let _source4 = _module.D14.create_DC31(_dafny.Seq.of(_dafny.Set.fromElements(!(true)), _dafny.Set.fromElements(false), _dafny.Set.fromElements(!(true)), _dafny.Set.fromElements((_module.D10.create_DC24(_dafny.Seq.UnicodeFromString("hy"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_74_i0) {
  return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(855));
}), true)).dtor_cf46)));
      if (_source4.is_DC32) {
        let _75___mcc_h0 = (_source4).cf56;
        let _76___mcc_h1 = (_source4).cf57;
        let _77___mcc_h2 = (_source4).cf58;
        let _78___mcc_h3 = (_source4).cf59;
        let _79___mcc_h4 = (_source4).cf60;
        let _80_cf60 = _79___mcc_h4;
        let _81_cf59 = _78___mcc_h3;
        let _82_cf58 = _77___mcc_h2;
        let _83_cf57 = _76___mcc_h1;
        let _84_cf56 = _75___mcc_h0;
        if (_82_cf58) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }
      } else {
        let _85___mcc_h5 = (_source4).cf55;
        let _86_cf55 = _85___mcc_h5;
        if (true) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }
      }
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(false, true), (_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(!(!(true)), false, true, false)), (_dafny.Set.fromElements(true, false, false)).Intersect(_dafny.Set.fromElements((_module.D2.create_DC7(true, true, new BigNumber(-255), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))).cardinality()))).dtor_cf15, true, true, !(false), false)), ((true) ? (_dafny.Set.fromElements(true, false, true, false)) : (_dafny.Set.fromElements(!(!(true))))));
    };
    static fm42(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dagoucki")).length))).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-490)))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('q'.codePointAt(0))))).length), new BigNumber(760)))));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), function (_87_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length), new BigNumber(385))));
    };
    static fm45(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((new BigNumber(-250)).minus(new BigNumber((_dafny.Seq.of(!(true))).length)));
    };
    static fm46(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(528),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(117),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("chdvwyj")).length), new BigNumber((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_dafny.Seq.of(new BigNumber(((_module.D7.create_DC19(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).length), _dafny.MultiSet.fromElements(true), true, _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(794)))).dtor_cf37).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_88_i0) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })).length))).Elements) {
          let _89_v0 = _compr_24;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(((_module.D7.create_DC19(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))).length), _dafny.MultiSet.fromElements(true), true, _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(794)))).dtor_cf37).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_88_i0) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          })).length)), _89_v0)) {
            _coll24.push([(_89_v0).minus(new BigNumber(423)),new BigNumber(909)]);
          }
        }
        return _coll24;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(612))).length)))).cardinality()))).length),!(!(true)))))).Merge((function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of _dafny.IntegerRange(new BigNumber(442), new BigNumber(367))) {
          let _90_v1 = _compr_25;
          if (((new BigNumber(442)).isLessThanOrEqualTo(_90_v1)) && ((_90_v1).isLessThan(new BigNumber(367)))) {
            _coll25.push([(_90_v1).plus(new BigNumber(-352)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(492),false)]);
          }
        }
        return _coll25;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(447)), function (_91_i1) {
        return new BigNumber(465);
      })).length),false)).length),function () {
        let _coll26 = new _dafny.Map();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(812), new BigNumber(557))) {
          let _92_v2 = _compr_26;
          if (((new BigNumber(812)).isLessThanOrEqualTo(_92_v2)) && ((_92_v2).isLessThan(new BigNumber(557)))) {
            _coll26.push([(_92_v2).multipliedBy(new BigNumber(894)),true]);
          }
        }
        return _coll26;
      }())));
    };
    static fm47(p0, p1, p2, p3, globalState) {
      let _source5 = _module.D16.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(610),true));
      if (_source5.is_DC38) {
        let _93___mcc_h0 = (_source5).cf69;
        let _94___mcc_h1 = (_source5).cf70;
        let _95___mcc_h2 = (_source5).cf71;
        let _96_cf71 = _95___mcc_h2;
        let _97_cf70 = _94___mcc_h1;
        let _98_cf69 = _93___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_96_cf71)).length))).length),_96_cf71);
      } else if (_source5.is_DC37) {
        let _99___mcc_h3 = (_source5).cf68;
        let _100_cf68 = _99___mcc_h3;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vtfmeu"),false)).length),true)).Merge(function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_100_cf68).Keys.Elements) {
            let _101_v0 = _compr_27;
            if ((_100_cf68).contains(_101_v0)) {
              _coll27.push([_module.__default.safeModuloInt(_101_v0, new BigNumber(-772)),true]);
            }
          }
          return _coll27;
        }());
      } else {
        let _102___mcc_h4 = (_source5).cf72;
        let _103_cf72 = _102___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tqifsibhn")).length),true);
      }
    };
    static fm48(p0, globalState) {
      return (_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(460)), function (_104_i0) {
        return (_module.D3.create_DC11(new _dafny.CodePoint('t'.codePointAt(0)))).dtor_cf22;
      })));
    };
    static fm49(p0, p1, p2, globalState) {
      return (((true) ? (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll28 = new _dafny.Set();
        for (const _compr_28 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(536))).Keys.Elements) {
          let _105_v0 = _compr_28;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new BigNumber(536))).contains(_105_v0)) {
            _coll28.add(_105_v0);
          }
        }
        return _coll28;
      }()).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(923))).length)))) : (_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(575), new BigNumber(305), new BigNumber(46))), _dafny.MultiSet.fromElements(new BigNumber(-264), new BigNumber(693)), _dafny.MultiSet.fromElements(new BigNumber(632)), _dafny.MultiSet.fromElements(new BigNumber(170)))))).Difference(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(178),!(false))).length))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(false)).length))).length))));
    };
    static fm50(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-384)))).Difference((_dafny.MultiSet.fromElements(function () {
        let _coll29 = new _dafny.Set();
        for (const _compr_29 of (_dafny.Seq.of(new BigNumber(339), new BigNumber(982))).Elements) {
          let _106_v0 = _compr_29;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(339), new BigNumber(982)), _106_v0)) {
            _coll29.add((_106_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
              let _coll30 = new _dafny.Map();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(-368), new BigNumber(100))) {
                let _107_v1 = _compr_30;
                if (((new BigNumber(-368)).isLessThanOrEqualTo(_107_v1)) && ((_107_v1).isLessThan(new BigNumber(100)))) {
                  _coll30.push([_module.__default.safeDivisionInt(_107_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-175)), function (_108_i0) {
                    return new _dafny.CodePoint('i'.codePointAt(0));
                  })).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(932),_dafny.MultiSet.fromElements(new BigNumber(733)))).length))).length))).length)),new BigNumber(988)]);
                }
              }
              return _coll30;
            }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("xnb")).length))).length)));
          }
        }
        return _coll29;
      }())).Intersect(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(246))).length)))));
    };
    static fm51(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((_dafny.ZERO).minus(((false) ? (new BigNumber(443)) : (new BigNumber(851)))), new BigNumber(571), new BigNumber(829));
    };
    static fm52(p0, p1, globalState) {
      return (_dafny.Set.fromElements(!(false))).Union(_dafny.Set.fromElements(false));
    };
    static fm53(p0, p1, p2, p3, globalState) {
      let _source6 = _module.D10.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((function () {
  let _coll31 = new _dafny.Map();
  for (const _compr_31 of (_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-916))).cardinality()), new BigNumber(671)), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), function () {
    let _coll32 = new _dafny.Set();
    for (const _compr_32 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).Elements) {
      let _109_v1 = _compr_32;
      if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)), _109_v1)) {
        _coll32.add(_module.__default.safeModuloInt(_109_v1, new BigNumber((_dafny.Set.fromElements(!(false))).length)));
      }
    }
    return _coll32;
  }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length)))).Elements) {
    let _110_v0 = _compr_31;
    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-916))).cardinality()), new BigNumber(671)), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)), function () {
      let _coll33 = new _dafny.Set();
      for (const _compr_33 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).Elements) {
        let _111_v1 = _compr_33;
        if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)), _111_v1)) {
          _coll33.add(_module.__default.safeModuloInt(_111_v1, new BigNumber((_dafny.Set.fromElements(!(false))).length)));
        }
      }
      return _coll33;
    }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true, true)).length))), _110_v0)) {
      _coll31.push([_110_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(64)), function (_112_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length))]);
    }
  }
  return _coll31;
}()).length), new BigNumber(660)),_dafny.Seq.of(_module.D2.create_DC6(function () {
  let _coll34 = new _dafny.Map();
  for (const _compr_34 of _dafny.IntegerRange(new BigNumber(874), new BigNumber(495))) {
    let _113_v2 = _compr_34;
    if (((new BigNumber(874)).isLessThanOrEqualTo(_113_v2)) && ((_113_v2).isLessThan(new BigNumber(495)))) {
      _coll34.push([(_113_v2).multipliedBy(new BigNumber(-116)),new BigNumber(823)]);
    }
  }
  return _coll34;
}()), _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(461),new BigNumber(387))))));
      if (_source6.is_DC24) {
        let _114___mcc_h0 = (_source6).cf44;
        let _115___mcc_h1 = (_source6).cf45;
        let _116___mcc_h2 = (_source6).cf46;
        let _117_cf46 = _116___mcc_h2;
        let _118_cf45 = _115___mcc_h1;
        let _119_cf44 = _114___mcc_h0;
        return _module.D6.create_DC17(new BigNumber(-489), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_117_cf46,false)).length), new BigNumber(973), new BigNumber(287), new BigNumber(272));
      } else if (_source6.is_DC25) {
        return _module.D6.create_DC17(new BigNumber((_dafny.Seq.UnicodeFromString("tu")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(-503), new BigNumber(26), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("grxj")).length)));
      } else {
        let _120___mcc_h3 = (_source6).cf43;
        let _121_cf43 = _120___mcc_h3;
        if (false) {
          return _module.D6.create_DC17(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(-178), new BigNumber(-617), new BigNumber(884), new BigNumber((function () {
  let _coll35 = new _dafny.Map();
  for (const _compr_35 of (_dafny.Seq.of(function () {
    let _coll36 = new _dafny.Map();
    for (const _compr_36 of _dafny.IntegerRange(new BigNumber(656), new BigNumber(744))) {
      let _122_v4 = _compr_36;
      if (((new BigNumber(656)).isLessThanOrEqualTo(_122_v4)) && ((_122_v4).isLessThan(new BigNumber(744)))) {
        _coll36.push([(_122_v4).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))),false]);
      }
    }
    return _coll36;
  }())).Elements) {
    let _123_v3 = _compr_35;
    if (_dafny.Seq.contains(_dafny.Seq.of(function () {
      let _coll37 = new _dafny.Map();
      for (const _compr_37 of _dafny.IntegerRange(new BigNumber(656), new BigNumber(744))) {
        let _122_v4 = _compr_37;
        if (((new BigNumber(656)).isLessThanOrEqualTo(_122_v4)) && ((_122_v4).isLessThan(new BigNumber(744)))) {
          _coll37.push([(_122_v4).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length))),false]);
        }
      }
      return _coll37;
    }()), _123_v3)) {
      _coll35.push([_123_v3,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-843),!(false))).length),new BigNumber(538))]);
    }
  }
  return _coll35;
}()).length));
        } else {
          return _module.D6.create_DC17(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(264)), function (_124_i1) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vghjbcr")).length))), new BigNumber(-987), new BigNumber((_dafny.Seq.UnicodeFromString("rfqfeded")).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)))).length));
        }
      }
    };
    static fm54(p0, p1, p2, globalState) {
      return _module.D16.create_DC38(((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll38 = new _dafny.Map();
  for (const _compr_38 of _dafny.IntegerRange(new BigNumber(-354), new BigNumber(574))) {
    let _125_v0 = _compr_38;
    if (((new BigNumber(-354)).isLessThanOrEqualTo(_125_v0)) && ((_125_v0).isLessThan(new BigNumber(574)))) {
      _coll38.push([_module.__default.safeModuloInt(_125_v0, new BigNumber(-173)),false]);
    }
  }
  return _coll38;
}()).length))).minus(new BigNumber(443)), ((true) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(432),(_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(614), (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll39 = new _dafny.Map();
  for (const _compr_39 of (_dafny.MultiSet.fromElements(new BigNumber(746), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wwtt")).length)))).Elements) {
    let _126_v1 = _compr_39;
    if ((_dafny.MultiSet.fromElements(new BigNumber(746), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wwtt")).length)))).contains(_126_v1)) {
      _coll39.push([(_126_v1).minus(new BigNumber((_dafny.Seq.of(new BigNumber(730))).length)),_dafny.Seq.of(true)]);
    }
  }
  return _coll39;
}()).length)), new BigNumber((_dafny.Seq.UnicodeFromString("g")).length))).length))),_dafny.Seq.of(true, true, true))).length)) : ((_dafny.ZERO).minus(new BigNumber(-784)))), false);
    };
    static fm55(globalState) {
      return _dafny.Seq.of((function () {
        let _coll40 = new _dafny.Set();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(284), new BigNumber(-98))) {
          let _127_v0 = _compr_40;
          if (((new BigNumber(284)).isLessThanOrEqualTo(_127_v0)) && ((_127_v0).isLessThan(new BigNumber(-98)))) {
            _coll40.add(_module.__default.safeModuloInt(_127_v0, new BigNumber(210)));
          }
        }
        return _coll40;
      }()).Union(function () {
        let _coll41 = new _dafny.Set();
        for (const _compr_41 of _dafny.IntegerRange(new BigNumber(314), new BigNumber(577))) {
          let _128_v1 = _compr_41;
          if (((new BigNumber(314)).isLessThanOrEqualTo(_128_v1)) && ((_128_v1).isLessThan(new BigNumber(577)))) {
            _coll41.add(_module.__default.safeDivisionInt(_128_v1, new BigNumber(-287)));
          }
        }
        return _coll41;
      }()));
    };
    static fm56(globalState) {
      return _module.D3.create_DC11(((!(true)) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('d'.codePointAt(0)))));
    };
    static m0(p0, globalState) {
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.Seq.of();
      let r2 = false;
      let r3 = _dafny.Map.Empty;
      let _129_v0;
      _129_v0 = true;
      let _130_v1;
      _130_v1 = _dafny.Seq.of(_129_v0);
      let _131_v2;
      _131_v2 = new BigNumber(762);
      if ((_130_v1)[_module.__default.safeIndex(_131_v2, new BigNumber((_130_v1).length))]) {
        let _132_v3;
        _132_v3 = _dafny.MultiSet.fromElements(_module.__default.fm39(globalState));
        let _133_v4;
        _133_v4 = _dafny.Map.Empty.slice().updateUnsafe(_132_v3,_129_v0);
        let _134_v5;
        _134_v5 = _module.D0.create_DC0(_129_v0);
        let _135_v6;
        _135_v6 = _dafny.Seq.of(_134_v5, _134_v5);
        let _136_v7;
        _136_v7 = _module.D0.create_DC2((_135_v6)[_module.__default.safeIndex(_131_v2, new BigNumber((_135_v6).length))]);
        _133_v4 = (_133_v4).update(_module.__default.fm10(_131_v2, _131_v2, _136_v7, globalState), _129_v0);
        (globalState).f15 = !(_129_v0);
        let _137_v8;
        let _nw0 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _137_v8 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_137_v8).length));
        (_137_v8)[_index0] = _module.__default.safeDivisionInt(_131_v2, _131_v2);
        let _index1 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_137_v8).length));
        (_137_v8)[_index1] = _module.__default.fm39(globalState);
        (globalState).f22 = _129_v0;
        (globalState).f12 = new BigNumber((_module.__default.fm22((_137_v8)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_137_v8).length))], _129_v0, globalState)).length);
      } else {
        let _138_v9;
        _138_v9 = _dafny.Map.Empty.slice().updateUnsafe(_131_v2,true);
        let _139_v10;
        _139_v10 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,!(_129_v0));
        let _140_v11;
        _140_v11 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,_131_v2);
        let _141_v12;
        _141_v12 = _dafny.Seq.of((((_140_v11).contains(true)) ? ((_140_v11).get(true)) : (_131_v2)));
        let _142_v13;
        _142_v13 = _dafny.Seq.of(_141_v12, _141_v12, _dafny.Seq.of(new BigNumber((_130_v1).length), _131_v2), _141_v12, _141_v12);
        let _143_v14;
        let _nw1 = new _module.C2();
        _nw1.__ctor(_142_v13, _129_v0, _129_v0);
        _143_v14 = _nw1;
        let _144_v15;
        _144_v15 = _dafny.Map.Empty.slice().updateUnsafe((_129_v0) === ((((_138_v9).contains(new BigNumber((_139_v10).length))) ? ((_138_v9).get(new BigNumber((_139_v10).length))) : (_129_v0))),_143_v14);
        let _145_v16;
        _145_v16 = _dafny.MultiSet.fromElements(_129_v0, _129_v0);
        let _146_v17;
        let _nw2 = new _module.C6();
        _nw2.__ctor(new BigNumber((_145_v16).cardinality()));
        _146_v17 = _nw2;
        let _147_v18;
        _147_v18 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,_146_v17);
        let _148_v19;
        _148_v19 = _module.D0.create_DC1(true, _129_v0, (_141_v12)[_module.__default.safeIndex(new BigNumber((_147_v18).length), new BigNumber((_141_v12).length))], !(_129_v0));
        _144_v15 = (_144_v15).update((_148_v19).dtor_cf1, _143_v14);
        (globalState).f2 = _129_v0;
        let _149_v20;
        let _init0 = function (_150_i0) {
          return true;
        };
        let _nw3 = Array((new BigNumber(28)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw3.length); _i0_0++) {
          _nw3[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _149_v20 = _nw3;
        let _index2 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length));
        (_149_v20)[_index2] = _129_v0;
        let _index3 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length));
        (_149_v20)[_index3] = false;
        let _151_v21;
        let _nw4 = Array((new BigNumber(15)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _145_v16;
        _nw4[(_dafny.ONE).toNumber()] = _145_v16;
        _nw4[(new BigNumber(2)).toNumber()] = (_145_v16).Union(_145_v16);
        _nw4[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_130_v1);
        _nw4[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.fromElements(false)).Union((_145_v16).update(_129_v0, _module.__default.abs((_dafny.ZERO).minus(_131_v2))));
        _nw4[(new BigNumber(5)).toNumber()] = _145_v16;
        _nw4[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]);
        _nw4[(new BigNumber(7)).toNumber()] = _module.__default.fm9(true, _146_v17.f40, _129_v0, globalState);
        _nw4[(new BigNumber(8)).toNumber()] = _module.__default.fm21(_146_v17.f40, globalState);
        _nw4[(new BigNumber(9)).toNumber()] = _145_v16;
        _nw4[(new BigNumber(10)).toNumber()] = _145_v16;
        _nw4[(new BigNumber(11)).toNumber()] = (_145_v16).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))])));
        _nw4[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_129_v0, (_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]);
        _nw4[(new BigNumber(13)).toNumber()] = (_145_v16).Difference(_module.__default.fm9(_129_v0, _131_v2, (_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))], globalState));
        _nw4[(new BigNumber(14)).toNumber()] = _145_v16;
        _151_v21 = _nw4;
        let _index4 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_151_v21).length));
        (_151_v21)[_index4] = _145_v16;
        let _index5 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_151_v21).length));
        (_151_v21)[_index5] = _145_v16;
        let _152_v22;
        _152_v22 = _dafny.Seq.of(p0);
        if ((new BigNumber((_dafny.Seq.Concat(_152_v22, _dafny.Seq.of(p0))).length)).isEqualTo(new BigNumber(365))) {
          (globalState).f16 = (new BigNumber(200)).plus(_131_v2);
          let _153_v23;
          _153_v23 = _module.D10.create_DC25();
          _153_v23 = _module.D10.create_DC25();
          let _154_v24;
          let _init1 = ((_155_v17, _156_v2, _157_v20, _158_v0) => function (_159_i1) {
            return _module.__default.safeDivisionInt(_159_i1, new BigNumber((_dafny.Set.fromElements(_155_v17.f40, (_dafny.ZERO).minus(_156_v2), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_157_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_157_v20).length))],_158_v0)).length))).length));
          })(_146_v17, _131_v2, _149_v20, _129_v0);
          let _nw5 = Array((new BigNumber(21)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
            _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _154_v24 = _nw5;
          let _index6 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_154_v24).length));
          (_154_v24)[_index6] = _146_v17.f40;
          let _index7 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_154_v24).length));
          (_154_v24)[_index7] = _module.__default.fm39(globalState);
          let _160_v25;
          _160_v25 = _module.D9.create_DC22((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]);
          let _161_v26;
          _161_v26 = _dafny.Map.Empty.slice().updateUnsafe(_130_v1,_149_v20);
          let _162_v27;
          _162_v27 = _module.D9.create_DC21(_161_v26);
          let _163_v28;
          _163_v28 = _dafny.Seq.of(_162_v27);
          let _index8 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_154_v24).length));
          let _index9 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_154_v24).length));
          let _rhs0 = _146_v17.f40;
          let _rhs1 = new BigNumber((_dafny.Seq.Concat(_141_v12, _141_v12)).length);
          let _rhs2 = (_160_v25).dtor_cf42;
          let _rhs3 = new BigNumber(((((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]) ? (_163_v28) : (_163_v28))).length);
          let _lhs0 = _146_v17;
          let _lhs1 = _154_v24;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_154_v24).length));
          let _lhs3 = _154_v24;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_154_v24).length));
          _lhs0.f40 = _rhs0;
          _lhs1[_lhs2] = _rhs1;
          _129_v0 = _rhs2;
          _lhs3[_lhs4] = _rhs3;
          (globalState).f2 = !_dafny.Seq.contains(_130_v1, (_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]);
          _138_v9 = (_138_v9).Merge(_dafny.Map.Empty.slice().updateUnsafe(_146_v17.f40,(_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]));
        } else {
          let _164_v29;
          _164_v29 = _dafny.Set.fromElements(new BigNumber((_140_v11).length));
          let _165_v30;
          _165_v30 = _module.D9.create_DC22((((_143_v14).fm16(new BigNumber((_152_v22).length), _dafny.Seq.of(_module.D2.create_DC9(_131_v2, new BigNumber((_164_v29).length), _129_v0)), _146_v17.f40, _146_v17.f40, globalState)) ? ((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))]) : (_129_v0)));
          _165_v30 = _165_v30;
          r1 = _130_v1;
          let _166_v31;
          _166_v31 = _module.D2.create_DC9(_146_v17.f40, _131_v2, (((_138_v9).contains(_146_v17.f40)) ? ((_138_v9).get(_146_v17.f40)) : ((_149_v20)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_149_v20).length))])));
          let _167_v32;
          _167_v32 = _dafny.Seq.of(_166_v31);
          let _168_v33;
          _168_v33 = _dafny.Map.Empty.slice().updateUnsafe(_131_v2,_146_v17.f40);
          (globalState).f12 = ((((_129_v0) ? (_129_v0) : ((_143_v14).fm16(_146_v17.f40, _167_v32, new BigNumber((_168_v33).length), _146_v17.f40, globalState)))) ? (new BigNumber(622)) : (_module.__default.safeModuloInt((_143_v14).fm1(globalState), new BigNumber((_dafny.MultiSet.fromElements((_146_v17).fm2(_129_v0, _131_v2, globalState))).cardinality()))));
          (globalState).f16 = (_146_v17).fm1(globalState);
          (globalState).f22 = _dafny.Seq.contains(_152_v22, p0);
        }
      }
      let _169_v34;
      let _init2 = ((_170_v1, _171_v2, _172_v0) => function (_173_i2) {
        return (_dafny.Set.fromElements(true, (_170_v1)[_module.__default.safeIndex(_171_v2, new BigNumber((_170_v1).length))])).Intersect(_dafny.Set.fromElements(true, _172_v0));
      })(_130_v1, _131_v2, _129_v0);
      let _nw6 = Array((new BigNumber(8)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
        _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _169_v34 = _nw6;
      let _174_v35;
      _174_v35 = _dafny.Set.fromElements(_129_v0);
      let _index10 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_169_v34).length));
      (_169_v34)[_index10] = (_174_v35).Union(_174_v35);
      let _175_v36;
      _175_v36 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_129_v0,_131_v2));
      let _176_v37;
      _176_v37 = _module.D21.create_DC52(new BigNumber((_dafny.Seq.of(_129_v0)).length), _175_v36);
      let _177_v38;
      _177_v38 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,_176_v37);
      let _178_v39;
      _178_v39 = _dafny.MultiSet.fromElements(_177_v38, _177_v38, _177_v38, (_177_v38).update(_129_v0, _module.D21.create_DC52(_131_v2, _175_v36)));
      let _179_v40;
      _179_v40 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,_174_v35);
      let _180_v41;
      _180_v41 = _dafny.Seq.of(_177_v38);
      let _index11 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_169_v34).length));
      let _rhs4 = _module.__default.fm18((((_179_v40).contains(_129_v0)) ? ((_179_v40).get(_129_v0)) : (_174_v35)), globalState);
      let _rhs5 = (_178_v39).Intersect(_dafny.MultiSet.FromArray(_180_v41));
      let _lhs5 = _169_v34;
      let _lhs6 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_169_v34).length));
      _lhs5[_lhs6] = _rhs4;
      _178_v39 = _rhs5;
      let _181_v42;
      let _nw7 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _181_v42 = _nw7;
      let _182_v43;
      _182_v43 = _dafny.MultiSet.fromElements(_131_v2);
      let _183_v44;
      _183_v44 = _dafny.MultiSet.fromElements(_182_v43);
      let _rhs6 = _129_v0;
      let _rhs7 = (_dafny.ZERO).minus((new BigNumber(812)).minus(((((_183_v44).contains(_182_v43)) ? ((_183_v44).get(_182_v43)) : (_131_v2))).plus(new BigNumber((_module.__default.fm18(_174_v35, globalState)).length))));
      let _rhs8 = (_131_v2).plus(_131_v2);
      let _rhs9 = _181_v42;
      let _rhs10 = p0;
      let _lhs7 = globalState;
      let _lhs8 = globalState;
      _lhs7.f2 = _rhs6;
      _lhs8.f12 = _rhs7;
      _131_v2 = _rhs8;
      _181_v42 = _rhs9;
      r0 = _rhs10;
      (globalState).f16 = (_dafny.ZERO).minus(_131_v2);
      let _index12 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_169_v34).length));
      (_169_v34)[_index12] = _module.__default.fm52(!(_129_v0) || (true), (_dafny.ZERO).minus((new BigNumber(((_169_v34)[_module.__default.safeIndex(new BigNumber(683), new BigNumber((_169_v34).length))]).length)).minus(_131_v2)), globalState);
      let _hi0 = _131_v2;
      for (let _184_i3 = _131_v2; _184_i3.isLessThan(_hi0); _184_i3 = _184_i3.plus(_dafny.ONE)) {
        (globalState).f2 = _129_v0;
        _131_v2 = _module.__default.safeModuloInt(_131_v2, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(_184_i3)).cardinality()), _131_v2)));
        let _185_v45;
        let _nw8 = new _module.C8();
        _nw8.__ctor(_182_v43);
        _185_v45 = _nw8;
        let _186_v46;
        _186_v46 = _module.D25.create_DC62(_185_v45);
        let _source7 = _module.D25.create_DC64(_186_v46);
        if (_source7.is_DC63) {
          let _187___mcc_h0 = (_source7).cf116;
          let _188___mcc_h1 = (_source7).cf117;
          let _189___mcc_h2 = (_source7).cf118;
          let _190_cf118 = _189___mcc_h2;
          let _191_cf117 = _188___mcc_h1;
          let _192_cf116 = _187___mcc_h0;
          let _193_v47;
          _193_v47 = _module.D15.create_DC35(_129_v0);
          let _194_v48;
          let _nw9 = new _module.C5();
          _nw9.__ctor((_193_v47).dtor_cf65, _191_cf117);
          _194_v48 = _nw9;
          let _195_v49;
          _195_v49 = _dafny.Seq.of(_194_v48);
          let _196_v50;
          _196_v50 = _dafny.Map.Empty.slice().updateUnsafe(_195_v49,_129_v0);
          _196_v50 = (_196_v50).update(_dafny.Seq.update(_195_v49, _module.__default.safeIndex(_131_v2, new BigNumber((_195_v49).length)), _194_v48), _129_v0);
          (globalState).f24 = (_129_v0) && (_129_v0);
          let _197_v51;
          _197_v51 = _dafny.MultiSet.fromElements(true);
          let _198_v52;
          _198_v52 = _dafny.Map.Empty.slice().updateUnsafe((_190_cf118).plus(_184_i3),_197_v51);
          _198_v52 = (_198_v52).update(_190_cf118, _dafny.MultiSet.fromElements(_191_cf117, _129_v0));
          let _199_v53;
          _199_v53 = _dafny.Map.Empty.slice().updateUnsafe(_182_v43,_129_v0);
          (globalState).f15 = (_199_v53).equals(_199_v53);
        } else if (_source7.is_DC62) {
          let _200___mcc_h3 = (_source7).cf115;
          let _201_cf115 = _200___mcc_h3;
          let _202_v54;
          let _nw10 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _202_v54 = _nw10;
          let _203_v55;
          _203_v55 = _module.D3.create_DC10(_202_v54);
          _202_v54 = (_203_v55).dtor_cf21;
          (globalState).f6 = _dafny.Seq.Concat(_module.__default.fm22(_184_i3, _129_v0, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), function (_204_i4) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          }));
          let _205_v56;
          let _init3 = ((_206_v0) => function (_207_i5) {
            return _206_v0;
          })(_129_v0);
          let _nw11 = Array((new BigNumber(7)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw11.length); _i0_3++) {
            _nw11[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _205_v56 = _nw11;
          let _index13 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length));
          (_205_v56)[_index13] = _129_v0;
          let _index14 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length));
          let _rhs11 = _129_v0;
          let _rhs12 = _129_v0;
          let _rhs13 = _129_v0;
          let _rhs14 = _129_v0;
          let _lhs9 = _205_v56;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length));
          let _lhs11 = globalState;
          let _lhs12 = globalState;
          let _lhs13 = globalState;
          _lhs9[_lhs10] = _rhs11;
          _lhs11.f2 = _rhs12;
          _lhs12.f2 = _rhs13;
          _lhs13.f15 = _rhs14;
          let _208_v57;
          _208_v57 = _dafny.MultiSet.fromElements(_202_v54, _202_v54);
          let _209_v58;
          _209_v58 = _dafny.Map.Empty.slice().updateUnsafe(true,_208_v57);
          let _210_v59;
          _210_v59 = _module.D26.create_DC65(_208_v57);
          let _211_v60;
          _211_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hmiq")).length),((_208_v57).update(_202_v54, _module.__default.abs(new BigNumber(((((_209_v58).contains((_205_v56)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length))])) ? ((_209_v58).get((_205_v56)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length))])) : ((_210_v59).dtor_cf120))).cardinality())))).Difference(_dafny.MultiSet.fromElements(_202_v54, _202_v54, _202_v54)));
          let _212_v61;
          _212_v61 = _dafny.Map.Empty.slice().updateUnsafe(_203_v55,_131_v2);
          let _213_v62;
          let _nw12 = new _module.C4();
          _nw12.__ctor(_212_v61, (_205_v56)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_205_v56).length))], _129_v0);
          _213_v62 = _nw12;
          let _214_v63;
          _214_v63 = _module.D16.create_DC38((_dafny.ZERO).minus(_131_v2), _184_i3, _129_v0);
          _211_v60 = (_211_v60).update(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_213_v62,_129_v0)).update(_213_v62, (_213_v62).fm2(true, (_214_v63).dtor_cf69, globalState))).length), _208_v57);
        } else {
          let _215___mcc_h4 = (_source7).cf119;
          let _216_cf119 = _215___mcc_h4;
          let _217_v64;
          let _218_v65;
          let _out0;
          let _out1;
          let _outcollector0 = (_185_v45).m4(globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _217_v64 = _out0;
          _218_v65 = _out1;
          (globalState).f12 = new BigNumber((_module.__default.fm22(new BigNumber(430), _218_v65, globalState)).length);
          (globalState).f16 = _184_i3;
          _217_v64 = _217_v64;
        }
        let _219_v66;
        let _init4 = ((_220_v2) => function (_221_i6) {
          return _module.__default.safeDivisionInt(_221_i6, _220_v2);
        })(_131_v2);
        let _nw13 = Array((new BigNumber(12)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw13.length); _i0_4++) {
          _nw13[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _219_v66 = _nw13;
        let _222_v67;
        _222_v67 = _dafny.Map.Empty.slice().updateUnsafe(_129_v0,new BigNumber(668));
        let _index15 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_219_v66).length));
        (_219_v66)[_index15] = (_131_v2).plus((((_222_v67).contains(_129_v0)) ? ((_222_v67).get(_129_v0)) : (_131_v2)));
        let _index16 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_219_v66).length));
        (_219_v66)[_index16] = ((_dafny.ZERO).minus(_module.__default.fm39(globalState))).plus(_module.__default.safeModuloInt(_module.__default.fm39(globalState), _131_v2));
      }
      r0 = p0;
      r1 = _module.__default.fm25(_129_v0, !(_129_v0) || (_129_v0), p0, globalState);
      r2 = (_131_v2).isLessThan(_131_v2);
      let _223_v68;
      _223_v68 = _dafny.MultiSet.fromElements(true, _129_v0);
      r3 = _dafny.Map.Empty.slice().updateUnsafe((_223_v68).Union(_223_v68),!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_130_v1, _module.__default.safeIndex(_131_v2, new BigNumber((_130_v1).length)), _129_v0), _130_v1)));
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _224_v0;
      _224_v0 = new BigNumber(979);
      let _225_v1;
      _225_v1 = _dafny.Seq.UnicodeFromString("ovdhmu");
      let _226_v2;
      _226_v2 = _dafny.Set.fromElements(_224_v0, (_dafny.ZERO).minus(new BigNumber((_225_v1).length)));
      let _227_v3;
      _227_v3 = _dafny.MultiSet.fromElements(_225_v1);
      let _228_v4;
      let _nw14 = Array((new BigNumber(8)).toNumber()).fill([]);
      _228_v4 = _nw14;
      let _229_v5;
      let _init5 = ((_230_v2) => function (_231_i1) {
        return (_231_i1).minus(new BigNumber((_230_v2).length));
      })(_226_v2);
      let _nw15 = Array((new BigNumber(17)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw15.length); _i0_5++) {
        _nw15[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _229_v5 = _nw15;
      let _232_v6;
      _232_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), function (_233_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }),_229_v5);
      let _234_v8;
      _234_v8 = new _dafny.CodePoint('v'.codePointAt(0));
      let _235_v9;
      _235_v9 = _dafny.Set.fromElements(_234_v8);
      let _236_v10;
      _236_v10 = _dafny.Map.Empty.slice().updateUnsafe(_224_v0,function () {
        let _coll42 = new _dafny.Map();
        for (const _compr_42 of (_235_v9).Elements) {
          let _237_v7 = _compr_42;
          if ((_235_v9).contains(_237_v7)) {
            _coll42.push([_237_v7,new BigNumber(-441)]);
          }
        }
        return _coll42;
      }());
      let _238_v11;
      _238_v11 = _dafny.Map.Empty.slice().updateUnsafe(_234_v8,_224_v0);
      let _239_v12;
      _239_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_236_v10).contains(_224_v0)) ? ((_236_v10).get(_224_v0)) : (_238_v11)),_224_v0);
      let _240_globalState;
      let _nw16 = new _module.GlobalState();
      _nw16.__ctor(_226_v2, _227_v3, true, true, new _dafny.CodePoint('l'.codePointAt(0)), _228_v4, _225_v1, _232_v6, false, true, new BigNumber(139), true, new BigNumber(66), new BigNumber(-610), _229_v5, false, new BigNumber(796), true, true, new _dafny.CodePoint('w'.codePointAt(0)), _225_v1, new BigNumber(942), true, new BigNumber(735), false, _239_v12);
      _240_globalState = _nw16;
      let _241_v14;
      _241_v14 = false;
      let _242_v15;
      _242_v15 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_241_v14,_224_v0));
      let _243_v16;
      _243_v16 = _dafny.Map.Empty.slice().updateUnsafe(false,_224_v0);
      let _244_v17;
      _244_v17 = _dafny.Map.Empty.slice().updateUnsafe(_243_v16,_241_v14);
      (_240_globalState).f12 = new BigNumber(((function () {
        let _coll43 = new _dafny.Map();
        for (const _compr_43 of (_242_v15).Elements) {
          let _245_v13 = _compr_43;
          if ((_242_v15).contains(_245_v13)) {
            _coll43.push([_245_v13,(_module.D0.create_DC1(_241_v14, _241_v14, _224_v0, _241_v14)).dtor_cf2]);
          }
        }
        return _coll43;
      }()).Merge(_244_v17)).length);
      if ((_227_v3).IsProperSubsetOf(_module.__default.fm0(false, _240_globalState))) {
        let _246_v18;
        let _247_v19;
        let _248_v20;
        let _249_v21;
        let _out2;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0(_234_v8, _240_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _out4 = _outcollector1[2];
        _out5 = _outcollector1[3];
        _246_v18 = _out2;
        _247_v19 = _out3;
        _248_v20 = _out4;
        _249_v21 = _out5;
        let _250_v22;
        _250_v22 = _module.D0.create_DC0(_248_v20);
        _250_v22 = _250_v22;
        (_240_globalState).f12 = _224_v0;
        let _251_v23;
        _251_v23 = _dafny.MultiSet.fromElements(new BigNumber(882));
        let _252_v24;
        let _nw17 = Array((new BigNumber(2)).toNumber());
        _nw17[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_224_v0);
        _nw17[(_dafny.ONE).toNumber()] = _251_v23;
        _252_v24 = _nw17;
        let _253_v25;
        _253_v25 = _module.D1.create_DC3(_247_v19);
        let _rhs15 = _241_v14;
        let _rhs16 = (_253_v25).dtor_cf6;
        let _rhs17 = _252_v24;
        let _rhs18 = _246_v18;
        let _lhs14 = _240_globalState;
        _lhs14.f2 = _rhs15;
        _247_v19 = _rhs16;
        _252_v24 = _rhs17;
        _246_v18 = _rhs18;
        let _254_v26;
        let _255_v27;
        let _256_v28;
        let _257_v29;
        let _out6;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector2 = _module.__default.m0(_234_v8, _240_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _out9 = _outcollector2[3];
        _254_v26 = _out6;
        _255_v27 = _out7;
        _256_v28 = _out8;
        _257_v29 = _out9;
      } else {
        (_240_globalState).f15 = (_226_v2).IsDisjointFrom(_226_v2);
        (_240_globalState).f24 = (_241_v14) || (_241_v14);
        let _258_v30;
        let _nw18 = new _module.C3();
        _nw18.__ctor(_dafny.MultiSet.fromElements(_dafny.Seq.of(_224_v0)), true, (_241_v14) && (_241_v14));
        _258_v30 = _nw18;
        let _259_v31;
        _259_v31 = _dafny.Map.Empty.slice().updateUnsafe(_241_v14,true);
        let _260_v32;
        let _nw19 = new _module.C1();
        _nw19.__ctor(_224_v0, new BigNumber((_259_v31).length), _241_v14, false);
        _260_v32 = _nw19;
        let _261_v33;
        _261_v33 = _dafny.Map.Empty.slice().updateUnsafe(_260_v32,_241_v14);
        let _262_v34;
        _262_v34 = _dafny.Seq.of(_261_v33);
        let _263_v35;
        _263_v35 = _dafny.MultiSet.fromElements(new BigNumber((_225_v1).length));
        let _264_v36;
        let _nw20 = new _module.C8();
        _nw20.__ctor(_263_v35);
        _264_v36 = _nw20;
        let _265_v37;
        _265_v37 = _dafny.MultiSet.fromElements(_264_v36, _264_v36, _264_v36, _264_v36);
        let _266_v38;
        _266_v38 = _dafny.Map.Empty.slice().updateUnsafe((_224_v0).plus(new BigNumber(((_262_v34)[_module.__default.safeIndex((_260_v32).f36, new BigNumber((_262_v34).length))]).length)),new BigNumber((_265_v37).cardinality()));
        _266_v38 = (_266_v38).update(_224_v0, (_224_v0).plus(_224_v0));
        let _267_v39;
        _267_v39 = _module.D3.create_DC11(new _dafny.CodePoint('b'.codePointAt(0)));
        let _268_v40;
        _268_v40 = _dafny.Map.Empty.slice().updateUnsafe(_267_v39,_241_v14);
        _268_v40 = (_268_v40).update(_module.__default.fm56(_240_globalState), _241_v14);
      }
      let _269_v41;
      let _init6 = ((_270_v14) => function (_271_i2) {
        return _270_v14;
      })(_241_v14);
      let _nw21 = Array((new BigNumber(4)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw21.length); _i0_6++) {
        _nw21[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _269_v41 = _nw21;
      let _index17 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
      (_269_v41)[_index17] = _241_v14;
      let _index18 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
      (_269_v41)[_index18] = _241_v14;
      let _272_v42;
      _272_v42 = _dafny.Seq.of(_224_v0, _224_v0, _224_v0, _224_v0);
      let _273_v43;
      _273_v43 = _dafny.Map.Empty.slice().updateUnsafe(_269_v41,_241_v14);
      let _274_v44;
      _274_v44 = _dafny.Map.Empty.slice().updateUnsafe((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))],(_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
      let _275_v45;
      _275_v45 = _dafny.Seq.of(_241_v14);
      let _276_v46;
      _276_v46 = _dafny.MultiSet.fromElements(new BigNumber((_275_v45).length), _224_v0, _224_v0);
      let _277_v47;
      let _nw22 = new _module.C9();
      _nw22.__ctor(_241_v14, new BigNumber((_274_v44).length), !((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]), !(_241_v14), _276_v46);
      _277_v47 = _nw22;
      let _278_v48;
      _278_v48 = _dafny.Seq.of(_277_v47);
      let _279_v49;
      _279_v49 = _dafny.MultiSet.fromElements(new BigNumber((_273_v43).length), _224_v0, new BigNumber((_278_v48).length));
      let _280_v50;
      let _nw23 = new _module.C8();
      _nw23.__ctor(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_272_v42, _module.__default.safeIndex(new BigNumber((_279_v49).cardinality()), new BigNumber((_272_v42).length)), new BigNumber(-790))).length)));
      _280_v50 = _nw23;
      let _281_v51;
      let _nw24 = Array((new BigNumber(2)).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = _280_v50;
      _nw24[(_dafny.ONE).toNumber()] = _280_v50;
      _281_v51 = _nw24;
      let _index19 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_281_v51).length));
      (_281_v51)[_index19] = _280_v50;
      let _index20 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_281_v51).length));
      (_281_v51)[_index20] = _280_v50;
      let _282_v52;
      _282_v52 = _dafny.Map.Empty.slice().updateUnsafe(_224_v0,_224_v0);
      let _283_v55;
      let _nw25 = Array((new BigNumber(23)).toNumber());
      _nw25[(_dafny.ZERO).toNumber()] = _282_v52;
      _nw25[(_dafny.ONE).toNumber()] = _282_v52;
      _nw25[(new BigNumber(2)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(3)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_241_v14,_241_v14)).length),_224_v0);
      _nw25[(new BigNumber(5)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(6)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(7)).toNumber()] = function () {
        let _coll44 = new _dafny.Map();
        for (const _compr_44 of (_272_v42).Elements) {
          let _284_v53 = _compr_44;
          if (_dafny.Seq.contains(_272_v42, _284_v53)) {
            _coll44.push([(_284_v53).plus((((_282_v52).contains(new BigNumber(165))) ? ((_282_v52).get(new BigNumber(165))) : (_224_v0))),_224_v0]);
          }
        }
        return _coll44;
      }();
      _nw25[(new BigNumber(8)).toNumber()] = (_282_v52).update(new BigNumber((_225_v1).length), new BigNumber((_dafny.Seq.of((_277_v47).f27, _241_v14)).length));
      _nw25[(new BigNumber(9)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(10)).toNumber()] = (_282_v52).update(_224_v0, new BigNumber((_225_v1).length));
      _nw25[(new BigNumber(11)).toNumber()] = function () {
        let _coll45 = new _dafny.Map();
        for (const _compr_45 of (_279_v49).Elements) {
          let _285_v54 = _compr_45;
          if ((_279_v49).contains(_285_v54)) {
            _coll45.push([(_285_v54).multipliedBy(new BigNumber((_243_v16).length)),_224_v0]);
          }
        }
        return _coll45;
      }();
      _nw25[(new BigNumber(12)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(13)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(14)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(15)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("etqlwqsyc")).length),new BigNumber((_275_v45).length));
      _nw25[(new BigNumber(17)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(18)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(19)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(20)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(21)).toNumber()] = _282_v52;
      _nw25[(new BigNumber(22)).toNumber()] = _282_v52;
      _283_v55 = _nw25;
      let _286_v56;
      _286_v56 = _module.D6.create_DC16(_225_v1);
      let _287_v57;
      _287_v57 = _dafny.Map.Empty.slice().updateUnsafe(_283_v55,new BigNumber(((((_277_v47).f27) ? ((_286_v56).dtor_cf29) : (_225_v1))).length));
      _224_v0 = (((_287_v57).contains(_283_v55)) ? ((_287_v57).get(_283_v55)) : (_224_v0));
      let _288_v58;
      let _nw26 = new _module.C6();
      _nw26.__ctor(_224_v0);
      _288_v58 = _nw26;
      let _289_v59;
      let _290_v60;
      let _out10;
      let _out11;
      let _outcollector3 = (_277_v47).m2((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))], _288_v58, (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))], _240_globalState);
      _out10 = _outcollector3[0];
      _out11 = _outcollector3[1];
      _289_v59 = _out10;
      _290_v60 = _out11;
      let _291_v61;
      let _292_v62;
      let _293_v63;
      let _294_v64;
      let _out12;
      let _out13;
      let _out14;
      let _out15;
      let _outcollector4 = _module.__default.m0(_289_v59, _240_globalState);
      _out12 = _outcollector4[0];
      _out13 = _outcollector4[1];
      _out14 = _outcollector4[2];
      _out15 = _outcollector4[3];
      _291_v61 = _out12;
      _292_v62 = _out13;
      _293_v63 = _out14;
      _294_v64 = _out15;
      let _295_v65;
      _295_v65 = _dafny.MultiSet.fromElements((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
      if (((_dafny.MultiSet.fromElements((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))])).Difference(_295_v65)).IsDisjointFrom((_295_v65).update(_293_v63, _module.__default.abs(_290_v60)))) {
        let _296_v66;
        let _297_v67;
        let _out16;
        let _out17;
        let _outcollector5 = (_277_v47).m2((_277_v47).f27, _288_v58, (_277_v47).f27, _240_globalState);
        _out16 = _outcollector5[0];
        _out17 = _outcollector5[1];
        _296_v66 = _out16;
        _297_v67 = _out17;
        let _index21 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_229_v5).length));
        (_229_v5)[_index21] = _224_v0;
        let _298_v68;
        _298_v68 = _dafny.MultiSet.fromElements(_272_v42);
        let _299_v69;
        let _nw27 = new _module.C3();
        _nw27.__ctor(((true) ? (_dafny.MultiSet.fromElements(_272_v42, _dafny.Seq.of(_297_v67))) : (_298_v68)), (_288_v58).fm2((_277_v47).f27, _224_v0, _240_globalState), (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
        _299_v69 = _nw27;
        let _index22 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_229_v5).length));
        let _rhs19 = (_297_v67).multipliedBy((_224_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("nhcidbx")).length)));
        let _rhs20 = _299_v69;
        let _lhs15 = _229_v5;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_229_v5).length));
        _lhs15[_lhs16] = _rhs19;
        _299_v69 = _rhs20;
        _274_v44 = (_274_v44).update(((_272_v42)[_module.__default.safeIndex(_224_v0, new BigNumber((_272_v42).length))]).isEqualTo((_dafny.ZERO).minus((_229_v5)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_229_v5).length))])), (_277_v47).f26);
        let _300_v70;
        _300_v70 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oujjykne"), _225_v1);
        (_299_v69).m1((_277_v47).f26, _dafny.Seq.update(_300_v70, _module.__default.safeIndex((_dafny.ZERO).minus(_224_v0), new BigNumber((_300_v70).length)), _225_v1), _module.__default.fm25((_277_v47).f26, (_277_v47).f26, _296_v66, _240_globalState), _240_globalState);
        let _301_v71;
        _301_v71 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(370),_269_v41);
        _301_v71 = (_301_v71).update((_229_v5)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_229_v5).length))], _269_v41);
      } else {
        (_240_globalState).f2 = !((_277_v47).f26) || ((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
        let _302_v72;
        _302_v72 = _dafny.Map.Empty.slice().updateUnsafe(_290_v60,_241_v14);
        let _303_v73;
        _303_v73 = _dafny.Map.Empty.slice().updateUnsafe((_277_v47).f26,_302_v72);
        let _304_v75;
        _304_v75 = _dafny.Set.fromElements((_277_v47).f26, (_277_v47).f26);
        let _305_v76;
        _305_v76 = _dafny.Map.Empty.slice().updateUnsafe(_304_v75,(_277_v47).f26);
        let _306_v78;
        let _nw28 = Array((new BigNumber(12)).toNumber());
        _nw28[(_dafny.ZERO).toNumber()] = _302_v72;
        _nw28[(_dafny.ONE).toNumber()] = _302_v72;
        _nw28[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_224_v0,(_277_v47).f27);
        _nw28[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_225_v1).length),(_277_v47).f26)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_272_v42).length),(_277_v47).fm2((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))], _290_v60, _240_globalState)));
        _nw28[(new BigNumber(4)).toNumber()] = (((_303_v73).contains(_293_v63)) ? ((_303_v73).get(_293_v63)) : (_302_v72));
        _nw28[(new BigNumber(5)).toNumber()] = function () {
          let _coll46 = new _dafny.Map();
          for (const _compr_46 of _dafny.IntegerRange(new BigNumber(100), new BigNumber(8))) {
            let _307_v74 = _compr_46;
            if (((new BigNumber(100)).isLessThanOrEqualTo(_307_v74)) && ((_307_v74).isLessThan(new BigNumber(8)))) {
              _coll46.push([(_307_v74).plus(_290_v60),(_277_v47).f27]);
            }
          }
          return _coll46;
        }();
        _nw28[(new BigNumber(6)).toNumber()] = _302_v72;
        _nw28[(new BigNumber(7)).toNumber()] = _302_v72;
        _nw28[(new BigNumber(8)).toNumber()] = _302_v72;
        _nw28[(new BigNumber(9)).toNumber()] = _302_v72;
        _nw28[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_224_v0,(_277_v47).f26)).update(new BigNumber((_305_v76).length), (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
        _nw28[(new BigNumber(11)).toNumber()] = function () {
          let _coll47 = new _dafny.Map();
          for (const _compr_47 of (_279_v49).Elements) {
            let _308_v77 = _compr_47;
            if ((_279_v49).contains(_308_v77)) {
              _coll47.push([_module.__default.safeModuloInt(_308_v77, _290_v60),(_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]]);
            }
          }
          return _coll47;
        }();
        _306_v78 = _nw28;
        let _index23 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_306_v78).length));
        (_306_v78)[_index23] = _302_v72;
        let _index24 = _module.__default.safeIndex(new BigNumber(490), new BigNumber((_306_v78).length));
        (_306_v78)[_index24] = (_302_v72).Merge(_302_v72);
        (_277_v47).m1((_277_v47).f27, _dafny.Seq.of(_225_v1, _225_v1), _292_v62, _240_globalState);
        let _index25 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_229_v5).length));
        (_229_v5)[_index25] = (_224_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("rodtn")).length));
        let _index26 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_229_v5).length));
        (_229_v5)[_index26] = (_224_v0).plus(_224_v0);
        let _309_v79;
        let _nw29 = new _module.C5();
        _nw29.__ctor((_277_v47).f27, (_288_v58).fm2(true, new BigNumber(144), _240_globalState));
        _309_v79 = _nw29;
        let _310_v80;
        _310_v80 = _module.D25.create_DC62((_281_v51)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_281_v51).length))]);
        let _311_v81;
        let _nw30 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _311_v81 = _nw30;
        let _312_v82;
        _312_v82 = _module.D3.create_DC10(_311_v81);
        let _313_v83;
        _313_v83 = _dafny.Map.Empty.slice().updateUnsafe(_312_v82,_224_v0);
        let _314_v84;
        let _nw31 = new _module.C4();
        _nw31.__ctor(_313_v83, (_277_v47).f27, (_277_v47).f27);
        _314_v84 = _nw31;
        let _315_v85;
        _315_v85 = _module.D24.create_DC60(_314_v84, new BigNumber((_225_v1).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_292_v62).length))).length), (_314_v84).f27);
        let _index27 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_229_v5).length));
        let _index28 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_281_v51).length));
        let _rhs21 = (new BigNumber(-755)).minus(new BigNumber(8));
        let _rhs22 = _309_v79;
        let _rhs23 = (_310_v80).dtor_cf115;
        let _rhs24 = _292_v62;
        let _rhs25 = _module.__default.safeDivisionInt(_224_v0, (_315_v85).dtor_cf112);
        let _lhs17 = _229_v5;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_229_v5).length));
        let _lhs19 = _281_v51;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_281_v51).length));
        _lhs17[_lhs18] = _rhs21;
        _309_v79 = _rhs22;
        _lhs19[_lhs20] = _rhs23;
        _292_v62 = _rhs24;
        _224_v0 = _rhs25;
      }
      (_240_globalState).f2 = (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))];
      let _316_v86;
      let _nw32 = Array((new BigNumber(27)).toNumber());
      _316_v86 = _nw32;
      let _317_v87;
      let _nw33 = new _module.C5();
      _nw33.__ctor(_293_v63, _241_v14);
      _317_v87 = _nw33;
      let _index29 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_316_v86).length));
      (_316_v86)[_index29] = _317_v87;
      let _318_v88;
      _318_v88 = _dafny.Seq.of(_288_v58, _288_v58);
      let _index30 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_316_v86).length));
      let _rhs26 = _289_v59;
      let _rhs27 = _224_v0;
      let _rhs28 = _317_v87;
      let _rhs29 = !(_241_v14);
      let _rhs30 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_288_v58, _288_v58, _288_v58, _288_v58, _288_v58), _dafny.Seq.Concat(_318_v88, _318_v88));
      let _lhs21 = _240_globalState;
      let _lhs22 = _240_globalState;
      let _lhs23 = _316_v86;
      let _lhs24 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_316_v86).length));
      let _lhs25 = _240_globalState;
      let _lhs26 = _240_globalState;
      _lhs21.f19 = _rhs26;
      _lhs22.f12 = _rhs27;
      _lhs23[_lhs24] = _rhs28;
      _lhs25.f15 = _rhs29;
      _lhs26.f22 = _rhs30;
      let _319_v89;
      _319_v89 = _dafny.Seq.of(_229_v5, _229_v5, _229_v5);
      _229_v5 = (_319_v89)[_module.__default.safeIndex(new BigNumber((_243_v16).length), new BigNumber((_319_v89).length))];
      let _source8 = _module.D7.create_DC18(_277_v47);
      if (_source8.is_DC19) {
        let _320___mcc_h0 = (_source8).cf36;
        let _321___mcc_h1 = (_source8).cf37;
        let _322___mcc_h2 = (_source8).cf38;
        let _323___mcc_h3 = (_source8).cf39;
        let _324_cf39 = _323___mcc_h3;
        let _325_cf38 = _322___mcc_h2;
        let _326_cf37 = _321___mcc_h1;
        let _327_cf36 = _320___mcc_h0;
        let _index31 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
        (_269_v41)[_index31] = (_288_v58).fm2((_277_v47).f26, _327_cf36, _240_globalState);
        (_317_v87).m7((_276_v46).IsSubsetOf(_dafny.MultiSet.FromArray(_272_v42)), (_277_v47).f26, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_327_cf36, _290_v60)), _276_v46, _240_globalState);
        _229_v5 = _229_v5;
        let _328_v90;
        _328_v90 = _dafny.Map.Empty.slice().updateUnsafe(_229_v5,_229_v5);
        let _329_v91;
        _329_v91 = _dafny.Map.Empty.slice().updateUnsafe(_328_v90,_290_v60);
        _329_v91 = (_329_v91).update(_328_v90, _224_v0);
      } else {
        let _330___mcc_h4 = (_source8).cf35;
        let _331_cf35 = _330___mcc_h4;
        let _332_v92;
        _332_v92 = _dafny.Map.Empty.slice().updateUnsafe(_225_v1,new BigNumber((_225_v1).length));
        let _333_v93;
        _333_v93 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_282_v52).length),(_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
        (_240_globalState).f16 = (((_332_v92).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-115)), ((_336_v8) => function (_337_i3) {
          return _336_v8;
        })(_234_v8)))) ? ((_332_v92).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-115)), ((_334_v8) => function (_335_i3) {
          return _334_v8;
        })(_234_v8)))) : (new BigNumber((_333_v93).length)));
        if (!(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_225_v1, _225_v1))).equals(_dafny.MultiSet.FromArray(_dafny.Seq.of(_291_v61)))) {
          let _338_v94;
          _338_v94 = _module.D0.create_DC0((_331_cf35).f26);
          (_240_globalState).f16 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(753), (_277_v47).fm4(_338_v94, _224_v0, _224_v0, _240_globalState)), new BigNumber(123)));
          (_240_globalState).f24 = (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))];
          let _index32 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
          (_269_v41)[_index32] = (((_331_cf35).f26) ? ((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]) : (!_dafny.Seq.contains(_225_v1, new _dafny.CodePoint('s'.codePointAt(0)))));
          let _index33 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
          (_269_v41)[_index33] = (_331_cf35).f27;
          let _pat_let_tv0 = _277_v47;
          let _339_v95;
          let _out18;
          _out18 = (_317_v87).m8(function (_pat_let0_0) {
            return function (_340_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_341_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_341_dt__update_hcf0_h0);
                }(_pat_let1_0);
              }((_pat_let_tv0).f26);
            }(_pat_let0_0);
          }(_338_v94), (_331_cf35).f27, _295_v65, (_224_v0).multipliedBy(_290_v60), _240_globalState);
          _339_v95 = _out18;
        } else {
          let _342_v96;
          _342_v96 = _dafny.Map.Empty.slice().updateUnsafe((_331_cf35).f27,_275_v45);
          _342_v96 = (_342_v96).update((_277_v47).f27, _275_v45);
          let _343_v97;
          _343_v97 = _dafny.Seq.of(_225_v1, _225_v1, _225_v1, _225_v1);
          let _344_v98;
          let _nw34 = new _module.C9();
          _nw34.__ctor(false, _224_v0, (_277_v47).f27, !_dafny.areEqual((_343_v97)[_module.__default.safeIndex(_224_v0, new BigNumber((_343_v97).length))], _225_v1), _279_v49);
          _344_v98 = _nw34;
          let _345_v99;
          let _nw35 = new _module.C5();
          _nw35.__ctor(!((_331_cf35).f26), (_277_v47).f27);
          _345_v99 = _nw35;
          _345_v99 = _345_v99;
          (_240_globalState).f22 = (_288_v58).fm2(!(_241_v14) || (!((_331_cf35).f26)), _290_v60, _240_globalState);
          let _346_v100;
          _346_v100 = _module.D0.create_DC0(false);
          (_240_globalState).f16 = (_331_cf35).fm4(_346_v100, _module.__default.safeModuloInt(new BigNumber((_275_v45).length), (_272_v42)[_module.__default.safeIndex((_344_v98).f31, new BigNumber((_272_v42).length))]), (((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]) ? (_224_v0) : (_224_v0)), _240_globalState);
        }
        let _347_v101;
        let _nw36 = new _module.C6();
        _nw36.__ctor((_290_v60).multipliedBy(_290_v60));
        _347_v101 = _nw36;
        let _348_v102;
        _348_v102 = _module.D0.create_DC0(false);
        let _349_v103;
        _349_v103 = _module.D5.create_DC14(_224_v0, _347_v101.f40, (_348_v102).dtor_cf0);
        let _350_v104;
        _350_v104 = _module.D5.create_DC15(_349_v103);
        let _pat_let_tv1 = _349_v103;
        _350_v104 = (((((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]) ? ((_331_cf35).f27) : (true))) ? (((_293_v63) ? (function (_pat_let2_0) {
          return function (_351_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_352_dt__update_hcf28_h0) {
                return _module.D5.create_DC15(_352_dt__update_hcf28_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_350_v104)) : (_350_v104))) : (_350_v104));
      }
      let _353_i4;
      _353_i4 = _dafny.ZERO;
      L0: {
        while ((_295_v65).IsSubsetOf((_295_v65).Difference(_295_v65))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_353_i4)) {
              break L0;
            }
            _353_i4 = (_353_i4).plus(_dafny.ONE);
            let _354_v105;
            _354_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(563),(_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
            let _355_v106;
            _355_v106 = _module.D16.create_DC37(_354_v105);
            let _356_v107;
            _356_v107 = _dafny.MultiSet.fromElements((_355_v106).dtor_cf68, _354_v105);
            let _357_v108;
            _357_v108 = _356_v107;
            let _source9 = _357_v108;
            let _358___mcc_h5 = _source9;
            let _359_cf82 = _358___mcc_h5;
            (_240_globalState).f2 = (_277_v47).f27;
            _243_v16 = _243_v16;
            _290_v60 = _module.__default.safeDivisionInt(_224_v0, (_288_v58).fm1(_240_globalState));
            let _360_v109;
            _360_v109 = _dafny.Map.Empty.slice().updateUnsafe(_229_v5,(_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
            let _rhs31 = (_277_v47).f27;
            let _rhs32 = new BigNumber((((_360_v109).Merge(_dafny.Map.Empty.slice().updateUnsafe(_229_v5,(_277_v47).f27))).Merge(_360_v109)).length);
            let _lhs27 = _240_globalState;
            _293_v63 = _rhs31;
            _lhs27.f12 = _rhs32;
            if ((_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]) {
              (_240_globalState).f22 = true;
              let _361_v110;
              let _init7 = ((_362_v1) => function (_363_i5) {
                return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("glxoh"), _362_v1);
              })(_225_v1);
              let _nw37 = Array((_dafny.ONE).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw37.length); _i0_7++) {
                _nw37[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _361_v110 = _nw37;
              let _364_v111;
              _364_v111 = _dafny.Map.Empty.slice().updateUnsafe(_290_v60,_225_v1);
              let _index34 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_361_v110).length));
              (_361_v110)[_index34] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-70)), function (_365_i6) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }), (((_364_v111).contains(_224_v0)) ? ((_364_v111).get(_224_v0)) : (_225_v1)));
              let _index35 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_361_v110).length));
              (_361_v110)[_index35] = _dafny.Seq.Concat(_225_v1, _225_v1);
              let _366_v112;
              _366_v112 = _module.D3.create_DC10(_229_v5);
              let _367_v113;
              _367_v113 = _dafny.Map.Empty.slice().updateUnsafe(_366_v112,_290_v60);
              let _368_v114;
              let _nw38 = new _module.C4();
              _nw38.__ctor(_367_v113, (_277_v47).f27, (_277_v47).f26);
              _368_v114 = _nw38;
              let _369_v115;
              _369_v115 = _dafny.Map.Empty.slice().updateUnsafe(_368_v114,_290_v60);
              _369_v115 = _369_v115;
              (_240_globalState).f6 = _225_v1;
              (_240_globalState).f6 = (_361_v110)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_361_v110).length))];
            } else {
              (_240_globalState).f24 = (_277_v47).f27;
              let _370_v116;
              let _nw39 = new _module.C1();
              _nw39.__ctor(new BigNumber((_272_v42).length), _290_v60, (_226_v2).contains(_224_v0), (_277_v47).f26);
              _370_v116 = _nw39;
              let _371_v117;
              let _init8 = ((_372_v61) => function (_373_i7) {
                return _372_v61;
              })(_291_v61);
              let _nw40 = Array((new BigNumber(8)).toNumber());
              for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw40.length); _i0_8++) {
                _nw40[_i0_8] = _init8(new BigNumber(_i0_8));
              }
              _371_v117 = _nw40;
              let _index36 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_371_v117).length));
              (_371_v117)[_index36] = _289_v59;
              let _index37 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_371_v117).length));
              (_371_v117)[_index37] = _291_v61;
              let _index38 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
              (_269_v41)[_index38] = _241_v14;
              let _374_v118;
              _374_v118 = _dafny.Map.Empty.slice().updateUnsafe((_272_v42)[_module.__default.safeIndex((_dafny.ZERO).minus(_370_v116.f37), new BigNumber((_272_v42).length))],_234_v8);
              _374_v118 = (_374_v118).update(_290_v60, (_371_v117)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((_371_v117).length))]);
            }
            let _375_v119;
            _375_v119 = _dafny.Map.Empty.slice().updateUnsafe(_272_v42,_224_v0);
            let _376_v120;
            _376_v120 = _dafny.Set.fromElements((_316_v86)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_316_v86).length))]);
            let _377_v121;
            _377_v121 = _dafny.Seq.of(_376_v120, _376_v120, _376_v120, _376_v120);
            _375_v119 = (_375_v119).update(_dafny.Seq.Concat(_272_v42, _272_v42), new BigNumber(((_377_v121)[_module.__default.safeIndex(_224_v0, new BigNumber((_377_v121).length))]).length));
            let _378_v122;
            let _nw41 = Array((new BigNumber(8)).toNumber());
            _378_v122 = _nw41;
            _378_v122 = _378_v122;
          }
        }
      }
      let _hi1 = _290_v60;
      for (let _379_i8 = new BigNumber((_dafny.Set.fromElements(false, (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))], (_277_v47).f26, _241_v14)).length); _379_i8.isLessThan(_hi1); _379_i8 = _379_i8.plus(_dafny.ONE)) {
        _269_v41 = _269_v41;
        let _380_v123;
        _380_v123 = _dafny.Set.fromElements((_277_v47).f26, (_277_v47).f26, (_269_v41)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length))]);
        let _381_v124;
        _381_v124 = _module.D1.create_DC5((_dafny.ZERO).minus(new BigNumber((_380_v123).length)));
        let _rhs33 = (_379_i8).minus((_381_v124).dtor_cf12);
        let _rhs34 = _224_v0;
        let _rhs35 = new BigNumber(855);
        let _rhs36 = (_226_v2).IsDisjointFrom(_dafny.Set.fromElements(_379_i8));
        let _lhs28 = _240_globalState;
        let _lhs29 = _240_globalState;
        let _lhs30 = _240_globalState;
        let _lhs31 = _240_globalState;
        _lhs28.f16 = _rhs33;
        _lhs29.f12 = _rhs34;
        _lhs30.f12 = _rhs35;
        _lhs31.f24 = _rhs36;
        let _index39 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
        let _rhs37 = new BigNumber(-455);
        let _rhs38 = (_277_v47).fm2(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_290_v60,(_dafny.ZERO).minus(new BigNumber(-932)))).length), _240_globalState);
        let _lhs32 = _240_globalState;
        let _lhs33 = _269_v41;
        let _lhs34 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
        _lhs32.f12 = _rhs37;
        _lhs33[_lhs34] = _rhs38;
        let _382_v125;
        _382_v125 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_224_v0, new BigNumber((_dafny.Seq.UnicodeFromString("yrpw")).length), _379_i8, (_dafny.ZERO).minus(_379_i8), new BigNumber(123)));
        let _383_v126;
        let _nw42 = new _module.C3();
        _nw42.__ctor(_382_v125, false, !((_277_v47).f27));
        _383_v126 = _nw42;
        let _384_v127;
        _384_v127 = _module.D23.create_DC55(_383_v126);
        _383_v126 = (_384_v127).dtor_cf103;
      }
      let _385_v128;
      let _386_v129;
      let _out19;
      let _out20;
      let _outcollector6 = (_280_v50).m4(_240_globalState);
      _out19 = _outcollector6[0];
      _out20 = _outcollector6[1];
      _385_v128 = _out19;
      _386_v129 = _out20;
      let _index40 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_269_v41).length));
      (_269_v41)[_index40] = !(!(_241_v14));
      process.stdout.write(_dafny.toString(_224_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_225_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_226_v2).equals(_dafny.Set.fromElements(new BigNumber(979), new BigNumber(-6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v3).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ovdhmu")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_232_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_234_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v9).equals(_dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(-441))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_239_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(-441)),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f0).equals(_dafny.Set.fromElements(new BigNumber(979), new BigNumber(-6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f1).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ovdhmu")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_240_globalState.f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_240_globalState).f7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f14)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_240_globalState).f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_globalState).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_240_globalState.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_240_globalState).f25).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(-441)),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_241_v14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_242_v15).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(979))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_244_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(979)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v41)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v41)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v41)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v41)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_272_v42, _dafny.Seq.of(new BigNumber(979), new BigNumber(979), new BigNumber(979), new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_273_v43).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_v44).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_275_v45, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_276_v46).equals(_dafny.MultiSet.fromElements(_dafny.ONE, new BigNumber(979), new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v47).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v47).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_278_v48).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v49).equals(_dafny.MultiSet.fromElements(_dafny.ONE, _dafny.ONE, new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_281_v51)[_dafny.ZERO]).f29).equals(_dafny.MultiSet.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_281_v51)[_dafny.ONE]).f29).equals(_dafny.MultiSet.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v52).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(1958),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)).updateUnsafe(new BigNumber(6),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(979)).updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_283_v55)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(979),new BigNumber(979)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_286_v56).dtor_cf29).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_287_v57).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_v59));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_290_v60));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_291_v61));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_292_v62, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_293_v63));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_294_v64).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true, true, true, true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v65).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_316_v86)[new BigNumber(2)]).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_316_v86)[new BigNumber(2)]).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_318_v88).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_319_v89).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_353_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_385_v128));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_386_v129));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf5) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, _dafny.ZERO, false);
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
    static create_DC4(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC5(cf12) {
      let $dt = new D1(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + this.cf10.toVerbatimString(true) + ", " + this.cf11.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4([], _dafny.ZERO, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC7(cf14, cf15, cf16, cf17) {
      let $dt = new D2(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC9(cf18, cf19, cf20) {
      let $dt = new D2(2);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC6(cf13) {
      let $dt = new D2(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8";
      } else if (this.$tag === 2) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC11(cf22) {
      let $dt = new D3(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC10(cf21) {
      let $dt = new D3(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC11(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC12(cf23) {
      let $dt = new D4(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf25, cf26, cf27) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D5(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf28) {
      let $dt = new D5(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC17(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC16(cf29) {
      let $dt = new D6(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + this.cf29.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC19(cf36, cf37, cf38, cf39) {
      let $dt = new D7(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC18(cf35) {
      let $dt = new D7(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf35 === other.cf35;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(_dafny.ZERO, _dafny.MultiSet.Empty, false, _dafny.Map.Empty);
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
    static create_DC20(cf40) {
      let $dt = new D8(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf42) {
      let $dt = new D9(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC21(cf41) {
      let $dt = new D9(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC22(false);
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
    static create_DC24(cf44, cf45, cf46) {
      let $dt = new D10(0);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC25() {
      let $dt = new D10(1);
      return $dt;
    }
    static create_DC23(cf43) {
      let $dt = new D10(2);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC24" + "(" + this.cf44.toVerbatimString(true) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC25";
      } else if (this.$tag === 2) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC24(_dafny.Seq.UnicodeFromString(""), _dafny.Seq.of(), false);
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
    static create_DC26(cf47) {
      let $dt = new D11(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47;
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
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf49) {
      let $dt = new D12(0);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC27(cf48) {
      let $dt = new D12(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf48 === other.cf48;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC28(_dafny.ZERO);
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
    static create_DC30(cf51, cf52, cf53, cf54) {
      let $dt = new D13(0);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC29(cf50) {
      let $dt = new D13(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53 && this.cf54 === other.cf54;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC30(false, _dafny.ZERO, false, []);
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
    static create_DC32(cf56, cf57, cf58, cf59, cf60) {
      let $dt = new D14(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC31(cf55) {
      let $dt = new D14(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC32" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && this.cf58 === other.cf58 && this.cf59 === other.cf59 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC32(_dafny.MultiSet.Empty, false, false, false, _dafny.ZERO);
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
    static create_DC34(cf62, cf63, cf64) {
      let $dt = new D15(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC35(cf65) {
      let $dt = new D15(1);
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC36(cf66, cf67) {
      let $dt = new D15(2);
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC33(cf61) {
      let $dt = new D15(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC35" + "(" + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf65 === other.cf65;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf66 === other.cf66 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC34(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC38(cf69, cf70, cf71) {
      let $dt = new D16(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC37(cf68) {
      let $dt = new D16(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC39(cf72) {
      let $dt = new D16(2);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC37() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC38" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC37" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC38(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC41(cf74, cf75, cf76) {
      let $dt = new D17(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC42(cf77, cf78, cf79) {
      let $dt = new D17(1);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC43(cf80, cf81) {
      let $dt = new D17(2);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC40(cf73) {
      let $dt = new D17(3);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get is_DC40() { return this.$tag === 3; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf74 === other.cf74 && this.cf75 === other.cf75 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78) && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC41(false, false, _dafny.ZERO);
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
    static create_DC44(cf82) {
      let $dt = new D18(0);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC45(cf83) {
      let $dt = new D19(0);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
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
    static create_DC47() {
      let $dt = new D20(0);
      return $dt;
    }
    static create_DC46(cf84) {
      let $dt = new D20(1);
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC48(cf85) {
      let $dt = new D20(2);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC47";
      } else if (this.$tag === 1) {
        return "D20.DC46" + "(" + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC48" + "(" + _dafny.toString(this.cf85) + ")";
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
        return other.$tag === 1 && this.cf84 === other.cf84;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC47();
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
    static create_DC50(cf87, cf88, cf89, cf90) {
      let $dt = new D21(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC51(cf91, cf92, cf93, cf94) {
      let $dt = new D21(1);
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC52(cf95, cf96) {
      let $dt = new D21(2);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC49(cf86) {
      let $dt = new D21(3);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get is_DC49() { return this.$tag === 3; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC50" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC51" + "(" + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC49" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf87 === other.cf87 && _dafny.areEqual(this.cf88, other.cf88) && this.cf89 === other.cf89 && _dafny.areEqual(this.cf90, other.cf90);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92 && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf86, other.cf86);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC50(false, _dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC54(cf98, cf99, cf100, cf101, cf102) {
      let $dt = new D22(0);
      $dt.cf98 = cf98;
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      return $dt;
    }
    static create_DC53(cf97) {
      let $dt = new D22(1);
      $dt.cf97 = cf97;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf97() { return this.cf97; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC54" + "(" + _dafny.toString(this.cf98) + ", " + _dafny.toString(this.cf99) + ", " + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + _dafny.toString(this.cf102) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC53" + "(" + _dafny.toString(this.cf97) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf98, other.cf98) && _dafny.areEqual(this.cf99, other.cf99) && this.cf100 === other.cf100 && this.cf101 === other.cf101 && this.cf102 === other.cf102;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf97, other.cf97);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC54(new _dafny.CodePoint('D'.codePointAt(0)), _module.D9.Default(), false, false, false);
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
    static create_DC56(cf104, cf105, cf106, cf107) {
      let $dt = new D23(0);
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      return $dt;
    }
    static create_DC57() {
      let $dt = new D23(1);
      return $dt;
    }
    static create_DC55(cf103) {
      let $dt = new D23(2);
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC58(cf108) {
      let $dt = new D23(3);
      $dt.cf108 = cf108;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get is_DC58() { return this.$tag === 3; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf108() { return this.cf108; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC56" + "(" + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ", " + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC57";
      } else if (this.$tag === 2) {
        return "D23.DC55" + "(" + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 3) {
        return "D23.DC58" + "(" + _dafny.toString(this.cf108) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf104, other.cf104) && _dafny.areEqual(this.cf105, other.cf105) && _dafny.areEqual(this.cf106, other.cf106) && this.cf107 === other.cf107;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf103 === other.cf103;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf108, other.cf108);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC56(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC60(cf110, cf111, cf112, cf113) {
      let $dt = new D24(0);
      $dt.cf110 = cf110;
      $dt.cf111 = cf111;
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      return $dt;
    }
    static create_DC59(cf109) {
      let $dt = new D24(1);
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC61(cf114) {
      let $dt = new D24(2);
      $dt.cf114 = cf114;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC59() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get dtor_cf110() { return this.cf110; }
    get dtor_cf111() { return this.cf111; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf114() { return this.cf114; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC60" + "(" + _dafny.toString(this.cf110) + ", " + _dafny.toString(this.cf111) + ", " + _dafny.toString(this.cf112) + ", " + _dafny.toString(this.cf113) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC59" + "(" + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC61" + "(" + _dafny.toString(this.cf114) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf110 === other.cf110 && _dafny.areEqual(this.cf111, other.cf111) && _dafny.areEqual(this.cf112, other.cf112) && this.cf113 === other.cf113;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf109, other.cf109);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf114, other.cf114);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC60(null, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC63(cf116, cf117, cf118) {
      let $dt = new D25(0);
      $dt.cf116 = cf116;
      $dt.cf117 = cf117;
      $dt.cf118 = cf118;
      return $dt;
    }
    static create_DC62(cf115) {
      let $dt = new D25(1);
      $dt.cf115 = cf115;
      return $dt;
    }
    static create_DC64(cf119) {
      let $dt = new D25(2);
      $dt.cf119 = cf119;
      return $dt;
    }
    get is_DC63() { return this.$tag === 0; }
    get is_DC62() { return this.$tag === 1; }
    get is_DC64() { return this.$tag === 2; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf119() { return this.cf119; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC63" + "(" + this.cf116.toVerbatimString(true) + ", " + _dafny.toString(this.cf117) + ", " + _dafny.toString(this.cf118) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC62" + "(" + _dafny.toString(this.cf115) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC64" + "(" + _dafny.toString(this.cf119) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf116, other.cf116) && this.cf117 === other.cf117 && _dafny.areEqual(this.cf118, other.cf118);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf115 === other.cf115;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf119, other.cf119);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC63(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC66(cf121, cf122, cf123) {
      let $dt = new D26(0);
      $dt.cf121 = cf121;
      $dt.cf122 = cf122;
      $dt.cf123 = cf123;
      return $dt;
    }
    static create_DC65(cf120) {
      let $dt = new D26(1);
      $dt.cf120 = cf120;
      return $dt;
    }
    static create_DC67(cf124) {
      let $dt = new D26(2);
      $dt.cf124 = cf124;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get dtor_cf121() { return this.cf121; }
    get dtor_cf122() { return this.cf122; }
    get dtor_cf123() { return this.cf123; }
    get dtor_cf120() { return this.cf120; }
    get dtor_cf124() { return this.cf124; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC66" + "(" + _dafny.toString(this.cf121) + ", " + _dafny.toString(this.cf122) + ", " + _dafny.toString(this.cf123) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC65" + "(" + _dafny.toString(this.cf120) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC67" + "(" + _dafny.toString(this.cf124) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf121, other.cf121) && this.cf122 === other.cf122 && this.cf123 === other.cf123;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf120, other.cf120);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf124, other.cf124);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC66(_dafny.ZERO, false, null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
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

  $module.T3 = class T3 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = false;
      this.f5 = [];
      this.f6 = _dafny.Seq.UnicodeFromString("");
      this.f12 = _dafny.ZERO;
      this.f15 = false;
      this.f16 = _dafny.ZERO;
      this.f19 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f22 = false;
      this.f24 = false;
      this._f0 = _dafny.Set.Empty;
      this._f1 = _dafny.MultiSet.Empty;
      this._f3 = false;
      this._f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f7 = _dafny.Map.Empty;
      this._f8 = false;
      this._f9 = false;
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f13 = _dafny.ZERO;
      this._f14 = [];
      this._f17 = false;
      this._f18 = false;
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f21 = _dafny.ZERO;
      this._f23 = _dafny.ZERO;
      this._f25 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f32 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f32) {
      let _this = this;
      (_this)._f32 = f32;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return (_this).f32;
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return true;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f26 = false;
      this._f27 = false;
      this.f37 = _dafny.ZERO;
      this._f36 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f36, f37, f26, f27) {
      let _this = this;
      (_this)._f36 = f36;
      (_this).f37 = f37;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_this).f26), _dafny.Seq.of((_this).f26)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f27), _dafny.Seq.of((_this).f27, (_this).f26)));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f36;
    };
    fm1(globalState) {
      let _this = this;
      return _this.f37;
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return (_this).f26;
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let _387_i0;
      _387_i0 = _dafny.ZERO;
      L1: {
        while (true) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_387_i0)) {
              break L1;
            }
            _387_i0 = (_387_i0).plus(_dafny.ONE);
            (globalState).f12 = _this.f37;
            (globalState).f24 = !(p2) || ((_this).fm2((_this).f26, _this.f37, globalState));
            let _388_v0;
            let _init9 = function (_389_i1) {
              return _dafny.Seq.of((_this).f36, _this.f37);
            };
            let _nw43 = Array((new BigNumber(10)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw43.length); _i0_9++) {
              _nw43[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _388_v0 = _nw43;
            let _390_v1;
            _390_v1 = _dafny.Seq.of(_this.f37, (_this).f36);
            let _index41 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_388_v0).length));
            (_388_v0)[_index41] = _390_v1;
            let _index42 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_388_v0).length));
            (_388_v0)[_index42] = _390_v1;
            (globalState).f2 = p2;
          }
        }
      }
      let _391_v2;
      _391_v2 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f27),!(p2) || (p2));
      let _392_v3;
      let _nw44 = Array((new BigNumber(5)).toNumber());
      _392_v3 = _nw44;
      let _393_v4;
      let _nw45 = new _module.C0();
      _nw45.__ctor((_this).f27);
      _393_v4 = _nw45;
      let _index43 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_392_v3).length));
      (_392_v3)[_index43] = _393_v4;
      let _394_v5;
      _394_v5 = _dafny.Set.fromElements(_this.f37, new BigNumber(975));
      let _395_v6;
      let _nw46 = Array((new BigNumber(9)).toNumber()).fill(false);
      _395_v6 = _nw46;
      let _396_v7;
      _396_v7 = _dafny.MultiSet.fromElements((_this).f36);
      let _397_v8;
      _397_v8 = _dafny.Map.Empty.slice().updateUnsafe(_395_v6,(((_396_v7).contains((_this).f36)) ? ((_396_v7).get((_this).f36)) : (_this.f37)));
      let _index44 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_392_v3).length));
      let _rhs39 = ((new BigNumber(-338)).minus(new BigNumber((_394_v5).length))).plus((_this).fm4(_module.__default.fm19((_this).f36, new BigNumber((_397_v8).length), globalState), new BigNumber(-144), (_this).f36, globalState));
      let _rhs40 = _391_v2;
      let _rhs41 = _393_v4;
      let _lhs35 = _this;
      let _lhs36 = _392_v3;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_392_v3).length));
      _lhs35.f37 = _rhs39;
      _391_v2 = _rhs40;
      _lhs36[_lhs37] = _rhs41;
      let _398_v9;
      _398_v9 = _module.D2.create_DC7(p2, (_this).f27, new BigNumber(-868), (_this).f36);
      let _399_v10;
      _399_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f37,(_398_v9).dtor_cf17);
      let _400_v11;
      _400_v11 = _module.D2.create_DC6(_399_v10);
      let _hi2 = new BigNumber((_dafny.Set.fromElements(_400_v11)).length);
      for (let _401_i2 = _this.f37; _401_i2.isLessThan(_hi2); _401_i2 = _401_i2.plus(_dafny.ONE)) {
        _394_v5 = _394_v5;
        let _402_v12;
        let _nw47 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _402_v12 = _nw47;
        let _index45 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_402_v12).length));
        (_402_v12)[_index45] = _module.__default.safeModuloInt(new BigNumber(796), (_this).f36);
        let _index46 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_402_v12).length));
        (_402_v12)[_index46] = _module.__default.safeModuloInt(_401_i2, _401_i2);
        let _403_v13;
        _403_v13 = new _dafny.CodePoint('a'.codePointAt(0));
        let _404_v14;
        let _nw48 = Array((new BigNumber(3)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = _403_v13;
        _nw48[(_dafny.ONE).toNumber()] = _403_v13;
        _nw48[(new BigNumber(2)).toNumber()] = _403_v13;
        _404_v14 = _nw48;
        let _405_v15;
        _405_v15 = _dafny.Map.Empty.slice().updateUnsafe((_402_v12)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_402_v12).length))],_404_v14);
        _405_v15 = (_405_v15).update(_401_i2, ((p0) ? (_404_v14) : (_404_v14)));
        let _index47 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_395_v6).length));
        (_395_v6)[_index47] = (_this).f26;
        let _406_v16;
        _406_v16 = _dafny.Seq.UnicodeFromString("tniumeoec");
        let _407_v17;
        _407_v17 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_406_v16).length)),_396_v7);
        let _index48 = _module.__default.safeIndex(new BigNumber(955), new BigNumber((_395_v6).length));
        (_395_v6)[_index48] = ((_dafny.ZERO).minus(new BigNumber((((((_407_v17).contains((_this).f36)) ? ((_407_v17).get((_this).f36)) : (_396_v7))).Intersect(_396_v7)).cardinality()))).isLessThan((_402_v12)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_402_v12).length))]);
      }
      (_this).f37 = _this.f37;
      let _408_v18;
      _408_v18 = _dafny.Seq.of((_this).f26);
      let _409_v19;
      _409_v19 = _dafny.Seq.of(new BigNumber((_408_v18).length));
      let _410_v20;
      _410_v20 = _409_v19;
      let _411_v21;
      _411_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_dafny.Seq.Concat((_410_v20), _dafny.Seq.of(new BigNumber((_408_v18).length))));
      let _412_v22;
      let _init10 = ((_413_v4, _414_v18, _415_p0) => function (_416_i3) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(_module.D2.create_DC9((_this).f36, _this.f37, (_413_v4).f32)), _dafny.MultiSet.fromElements(_module.D2.create_DC9((_this).f36, (_this).f36, true), _module.D2.create_DC9((_this).f36, (_this).f36, (_this).f27), _module.D2.create_DC9(_this.f37, _this.f37, (_413_v4).f32))), _dafny.Seq.of(_dafny.MultiSet.fromElements(_module.D2.create_DC9((_this).f36, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_414_v18),(_413_v4).f32)).length))).length), (_413_v4).f32), _module.D2.create_DC9(new BigNumber((_414_v18).length), _this.f37, _415_p0))));
      })(_393_v4, _408_v18, p0);
      let _nw49 = Array((_dafny.ONE).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw49.length); _i0_10++) {
        _nw49[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _412_v22 = _nw49;
      let _417_v23;
      _417_v23 = _module.D2.create_DC9((_this).f36, _this.f37, (_this).f27);
      let _418_v24;
      _418_v24 = _dafny.MultiSet.fromElements(_417_v23);
      let _419_v25;
      _419_v25 = _dafny.MultiSet.fromElements(p2);
      let _420_v26;
      _420_v26 = _dafny.Seq.of(_418_v24, _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(_417_v23), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_419_v25).cardinality())), new BigNumber((_dafny.Seq.of(_417_v23)).length)), _module.__default.fm20((_this).f36, p2, (_this).f36, (_398_v9).dtor_cf17, globalState))));
      let _index49 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_412_v22).length));
      (_412_v22)[_index49] = _dafny.Seq.Concat(_420_v26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(62)), ((_421_v24) => function (_422_i4) {
        return _421_v24;
      })(_418_v24)));
      let _423_v27;
      _423_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f37,_418_v24);
      let _424_v28;
      _424_v28 = _module.D0.create_DC0((_this).f26);
      let _index50 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_412_v22).length));
      let _rhs42 = _411_v21;
      let _rhs43 = _dafny.Seq.Concat(_dafny.Seq.update(_420_v26, _module.__default.safeIndex(new BigNumber(-716), new BigNumber((_420_v26).length)), (((_423_v27).contains(_this.f37)) ? ((_423_v27).get(_this.f37)) : (_418_v24))), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_420_v26, _module.__default.safeIndex((_this).fm4(_424_v28, _this.f37, (_this).f36, globalState), new BigNumber((_420_v26).length)), _418_v24), _module.__default.safeIndex((_this).f36, new BigNumber((_dafny.Seq.update(_420_v26, _module.__default.safeIndex((_this).fm4(_424_v28, _this.f37, (_this).f36, globalState), new BigNumber((_420_v26).length)), _418_v24)).length)), _418_v24), _dafny.Seq.of(_418_v24, _418_v24)));
      let _lhs38 = _412_v22;
      let _lhs39 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_412_v22).length));
      _411_v21 = _rhs42;
      _lhs38[_lhs39] = _rhs43;
      let _425_v29;
      let _init11 = ((_426_v25) => function (_427_i5) {
        return _426_v25;
      })(_419_v25);
      let _nw50 = Array((new BigNumber(29)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw50.length); _i0_11++) {
        _nw50[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _425_v29 = _nw50;
      _425_v29 = _425_v29;
      let _428_v30;
      _428_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,new _dafny.CodePoint('c'.codePointAt(0)));
      let _429_v31;
      _429_v31 = new _dafny.CodePoint('m'.codePointAt(0));
      r0 = (((_428_v30).contains((_393_v4).f32)) ? ((_428_v30).get((_393_v4).f32)) : (_429_v31));
      let _430_v32;
      _430_v32 = _dafny.Seq.of(_429_v31, _429_v31);
      let _431_v33;
      _431_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_430_v32).length),(_399_v10).update(_this.f37, (_this).f36));
      r1 = new BigNumber((_431_v33).length);
      return [r0, r1];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _432_v0;
      let _nw51 = Array((new BigNumber(3)).toNumber()).fill(false);
      _432_v0 = _nw51;
      let _index51 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length));
      (_432_v0)[_index51] = (_this).f26;
      let _433_v1;
      let _nw52 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _433_v1 = _nw52;
      let _434_v2;
      _434_v2 = _dafny.Set.fromElements((_this).f26);
      let _435_v3;
      _435_v3 = _dafny.Seq.of(_434_v2);
      let _436_v4;
      _436_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_435_v3)[_module.__default.safeIndex(new BigNumber(850), new BigNumber((_435_v3).length))]).length),(_this).f26);
      let _437_v5;
      _437_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_dafny.ZERO).minus(new BigNumber((_436_v4).length)));
      let _438_v6;
      _438_v6 = _dafny.MultiSet.fromElements((((_437_v5).contains(p0)) ? ((_437_v5).get(p0)) : (_this.f37)), _this.f37);
      let _439_v7;
      _439_v7 = _dafny.Seq.UnicodeFromString("k");
      let _440_v8;
      _440_v8 = new _dafny.CodePoint('t'.codePointAt(0));
      let _index52 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      (_433_v1)[_index52] = (((_438_v6).contains(new BigNumber((_dafny.Seq.update(_439_v7, _module.__default.safeIndex(new BigNumber(561), new BigNumber((_439_v7).length)), _440_v8)).length))) ? ((_438_v6).get(new BigNumber((_dafny.Seq.update(_439_v7, _module.__default.safeIndex(new BigNumber(561), new BigNumber((_439_v7).length)), _440_v8)).length))) : ((_this).fm1(globalState)));
      let _441_v9;
      let _init12 = ((_442_v8) => function (_443_i0) {
        return _442_v8;
      })(_440_v8);
      let _nw53 = Array((new BigNumber(24)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw53.length); _i0_12++) {
        _nw53[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _441_v9 = _nw53;
      let _index53 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length));
      let _index54 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _rhs44 = p0;
      let _rhs45 = _this.f37;
      let _rhs46 = _441_v9;
      let _lhs40 = _432_v0;
      let _lhs41 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length));
      let _lhs42 = _433_v1;
      let _lhs43 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      _lhs40[_lhs41] = _rhs44;
      _lhs42[_lhs43] = _rhs45;
      _441_v9 = _rhs46;
      let _444_v10;
      _444_v10 = _dafny.MultiSet.fromElements(p0, (_this).f27);
      let _445_v11;
      _445_v11 = _dafny.Seq.of(_this.f37, new BigNumber((_439_v7).length));
      let _index55 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _index56 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _index57 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _index58 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length));
      let _rhs47 = _this.f37;
      let _rhs48 = new BigNumber((((_444_v10).Union(_444_v10)).Union(_444_v10)).cardinality());
      let _rhs49 = new BigNumber(-718);
      let _rhs50 = ((((p0) ? (p0) : (!(p0)))) ? (!_dafny.areEqual(_445_v11, _445_v11)) : (((((_437_v5).contains((_this).f26)) ? ((_437_v5).get((_this).f26)) : (new BigNumber((_437_v5).length)))).isLessThan(new BigNumber(349))));
      let _rhs51 = (_this).f26;
      let _lhs44 = _433_v1;
      let _lhs45 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _lhs46 = _433_v1;
      let _lhs47 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _lhs48 = _433_v1;
      let _lhs49 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      let _lhs50 = globalState;
      let _lhs51 = _432_v0;
      let _lhs52 = _module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length));
      _lhs44[_lhs45] = _rhs47;
      _lhs46[_lhs47] = _rhs48;
      _lhs48[_lhs49] = _rhs49;
      _lhs50.f22 = _rhs50;
      _lhs51[_lhs52] = _rhs51;
      let _hi3 = (_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))];
      for (let _446_i1 = _this.f37; _446_i1.isLessThan(_hi3); _446_i1 = _446_i1.plus(_dafny.ONE)) {
        let _447_v12;
        let _nw54 = new _module.C0();
        _nw54.__ctor((_444_v10).IsSubsetOf(_module.__default.fm21(_this.f37, globalState)));
        _447_v12 = _nw54;
        let _448_v13;
        _448_v13 = _dafny.Set.fromElements(_this.f37, (_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))], (_this).f36, (_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))]);
        let _449_v14;
        _449_v14 = _dafny.Map.Empty.slice().updateUnsafe(_448_v13,_432_v0);
        _449_v14 = _449_v14;
        let _index59 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
        (_433_v1)[_index59] = _446_i1;
        (globalState).f15 = (p2)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_this.f37, new BigNumber((_434_v2).length)), new BigNumber((p2).length))];
      }
      let _450_v15;
      _450_v15 = _dafny.Map.Empty.slice().updateUnsafe(_434_v2,!((_this).f26));
      _450_v15 = ((p0) ? (_dafny.Map.Empty.slice().updateUnsafe(_434_v2,(_432_v0)[_module.__default.safeIndex(new BigNumber(254), new BigNumber((_432_v0).length))])) : (_450_v15));
      let _hi4 = (_this.f37).multipliedBy((_dafny.ZERO).minus((_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))]));
      for (let _451_i2 = (_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))]; _451_i2.isLessThan(_hi4); _451_i2 = _451_i2.plus(_dafny.ONE)) {
        let _452_v16;
        let _nw55 = new _module.C0();
        _nw55.__ctor((_this).f27);
        _452_v16 = _nw55;
        (globalState).f12 = ((_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))]).plus(_451_i2);
        let _453_v17;
        _453_v17 = _module.D1.create_DC4(_441_v9, (_433_v1)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length))], new BigNumber((_dafny.MultiSet.fromElements(_440_v8, _440_v8)).cardinality()), _module.__default.fm22(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), ((_454_v8) => function (_455_i3) {
  return _454_v8;
})(_440_v8))).length), (_this).f27, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), ((_456_v8) => function (_457_i4) {
  return _456_v8;
})(_440_v8)));
        _453_v17 = (((_this).f26) ? (_453_v17) : (_module.D1.create_DC4(_441_v9, new BigNumber(((_453_v17).dtor_cf11).length), _this.f37, _439_v7, _439_v7)));
        (globalState).f15 = false;
      }
      let _index60 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_433_v1).length));
      (_433_v1)[_index60] = _this.f37;
      return;
    }
    get f36() {
      let _this = this;
      return _this._f36;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f26 = false;
      this._f27 = false;
      this._f35 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f35, f26, f27) {
      let _this = this;
      (_this)._f35 = f35;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll48 = new _dafny.Set();
        for (const _compr_48 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_458_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(589),true)).length));
        })).Elements) {
          let _459_v0 = _compr_48;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_458_i0) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(589),true)).length));
          }), _459_v0)) {
            _coll48.add(_module.__default.safeDivisionInt(_459_v0, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false)).length))).length)));
          }
        }
        return _coll48;
      }()).length))).cardinality()),(_this).f27)).length), new BigNumber((_dafny.Set.fromElements(true)).length)),new BigNumber((_dafny.Seq.of((_this).f26, (_this).f27, (_this).f26)).length))).length))).equals(_dafny.Set.fromElements((_module.D0.create_DC1(false, (_this).f26, new BigNumber(832), (_this).f26)).dtor_cf3))) {
        return _dafny.Seq.of((_this).f26, (_this).f27);
      } else {
        return _dafny.Seq.of((_this).f26, (_this).f26);
      }
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((function (_source10) {
        if (_source10.is_DC7) {
          let _460___mcc_h0 = (_source10).cf14;
          let _461___mcc_h1 = (_source10).cf15;
          let _462___mcc_h2 = (_source10).cf16;
          let _463___mcc_h3 = (_source10).cf17;
          let _464_cf17 = _463___mcc_h3;
          let _465_cf16 = _462___mcc_h2;
          let _466_cf15 = _461___mcc_h1;
          let _467_cf14 = _460___mcc_h0;
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))));
        } else if (_source10.is_DC8) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0))));
        } else if (_source10.is_DC9) {
          let _468___mcc_h4 = (_source10).cf18;
          let _469___mcc_h5 = (_source10).cf19;
          let _470___mcc_h6 = (_source10).cf20;
          let _471_cf20 = _470___mcc_h6;
          let _472_cf19 = _469___mcc_h5;
          let _473_cf18 = _468___mcc_h4;
          return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))));
        } else {
          let _474___mcc_h7 = (_source10).cf13;
          let _475_cf13 = _474___mcc_h7;
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0))));
        }
      }(_module.D2.create_DC8())).length);
    };
    fm1(globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of((_module.D7.create_DC19(new BigNumber((_dafny.Seq.of((_this).f26)).length), _dafny.MultiSet.fromElements((_this).f27), (_this).f27, _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-748)), function (_476_i0) {
  return _dafny.Set.fromElements((_this).f26);
})).length)))).dtor_cf36, new BigNumber(-153)))).cardinality());
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    fm16(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source11 = _module.D6.create_DC17(new BigNumber((_dafny.Seq.UnicodeFromString("uot")).length), new BigNumber(884), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f26))).cardinality()), new BigNumber(-382), new BigNumber((_dafny.Set.fromElements((_this).f26, (_this).f27, true)).length));
      if (_source11.is_DC17) {
        let _477___mcc_h0 = (_source11).cf30;
        let _478___mcc_h1 = (_source11).cf31;
        let _479___mcc_h2 = (_source11).cf32;
        let _480___mcc_h3 = (_source11).cf33;
        let _481___mcc_h4 = (_source11).cf34;
        let _482_cf34 = _481___mcc_h4;
        let _483_cf33 = _480___mcc_h3;
        let _484_cf32 = _479___mcc_h2;
        let _485_cf31 = _478___mcc_h1;
        let _486_cf30 = _477___mcc_h0;
        return (_this).f26;
      } else {
        let _487___mcc_h5 = (_source11).cf29;
        let _488_cf29 = _487___mcc_h5;
        return !((_this).f26);
      }
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f26;
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      r1 = (_this).fm1(globalState);
      let _489_v0;
      _489_v0 = _dafny.Set.fromElements(p0);
      let _490_v1;
      _490_v1 = new BigNumber(-337);
      let _491_v2;
      _491_v2 = _dafny.Seq.UnicodeFromString("ajxqysp");
      let _rhs52 = (_dafny.ZERO).minus(((_490_v1).multipliedBy(new BigNumber((_491_v2).length))).minus(_490_v1));
      let _rhs53 = (_module.__default.fm18(_489_v0, globalState)).Difference(_489_v0);
      let _rhs54 = _490_v1;
      let _lhs53 = globalState;
      _lhs53.f16 = _rhs52;
      _489_v0 = _rhs53;
      r1 = _rhs54;
      let _492_v3;
      _492_v3 = _dafny.Seq.of(p0);
      let _493_v4;
      _493_v4 = _dafny.Seq.of(_dafny.Seq.update(_492_v3, _module.__default.safeIndex(new BigNumber((_492_v3).length), new BigNumber((_492_v3).length)), p2), _dafny.Seq.Concat(_492_v3, _492_v3));
      _493_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_493_v4, _493_v4), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), ((_494_v3) => function (_495_i0) {
        return _494_v3;
      })(_492_v3)), _module.__default.safeIndex(_490_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(567)), ((_496_v3) => function (_497_i0) {
        return _496_v3;
      })(_492_v3))).length)), _492_v3));
      let _498_v5;
      let _nw56 = Array((new BigNumber(18)).toNumber()).fill(false);
      _498_v5 = _nw56;
      let _index61 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_498_v5).length));
      (_498_v5)[_index61] = true;
      let _499_v6;
      _499_v6 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_490_v1, new BigNumber((_491_v2).length), _490_v1, _490_v1),(_this).f26);
      let _500_v7;
      _500_v7 = _dafny.MultiSet.fromElements(_490_v1, new BigNumber((_491_v2).length));
      let _501_v8;
      _501_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_491_v2).length));
      let _502_v9;
      _502_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_501_v8).length),!(false));
      let _index62 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_498_v5).length));
      (_498_v5)[_index62] = (((_499_v6).contains((_500_v7).Union(_dafny.MultiSet.fromElements(new BigNumber((_502_v9).length))))) ? ((_499_v6).get((_500_v7).Union(_dafny.MultiSet.fromElements(new BigNumber((_502_v9).length))))) : ((_this).f26));
      let _503_v10;
      _503_v10 = new _dafny.CodePoint('p'.codePointAt(0));
      let _504_v11;
      _504_v11 = _dafny.Map.Empty.slice().updateUnsafe((_491_v2)[_module.__default.safeIndex(_490_v1, new BigNumber((_491_v2).length))],_503_v10);
      let _505_v12;
      _505_v12 = _dafny.Seq.of(_504_v11);
      _505_v12 = _505_v12;
      let _506_v13;
      _506_v13 = _dafny.Seq.of(new BigNumber(-288), new BigNumber((_500_v7).cardinality()), new BigNumber(180), _490_v1, new BigNumber((_dafny.Seq.UnicodeFromString("flhymnjo")).length));
      let _507_v15;
      let _nw57 = Array((new BigNumber(18)).toNumber());
      _nw57[(_dafny.ZERO).toNumber()] = _506_v13;
      _nw57[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), ((_508_v1) => function (_509_i1) {
        return _508_v1;
      })(_490_v1));
      _nw57[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_506_v13, _506_v13);
      _nw57[(new BigNumber(3)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(4)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(((!((_498_v5)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_498_v5).length))])) ? (_506_v13) : (_506_v13)), _module.__default.safeIndex(_490_v1, new BigNumber((((!((_498_v5)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_498_v5).length))])) ? (_506_v13) : (_506_v13))).length)), _490_v1);
      _nw57[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_506_v13, _506_v13);
      _nw57[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), ((_510_v7) => function (_511_i2) {
        return new BigNumber((_510_v7).cardinality());
      })(_500_v7));
      _nw57[(new BigNumber(8)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_491_v2).length)), _module.__default.safeIndex(new BigNumber((function () {
        let _coll49 = new _dafny.Set();
        for (const _compr_49 of (_504_v11).Keys.Elements) {
          let _512_v14 = _compr_49;
          if ((_504_v11).contains(_512_v14)) {
            _coll49.add(_512_v14);
          }
        }
        return _coll49;
      }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_491_v2).length))).length)), _490_v1);
      _nw57[(new BigNumber(10)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(11)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(12)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(444)), ((_513_v1) => function (_514_i3) {
        return _513_v1;
      })(_490_v1));
      _nw57[(new BigNumber(14)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_506_v13, _module.__default.safeIndex(_490_v1, new BigNumber((_506_v13).length)), _490_v1);
      _nw57[(new BigNumber(16)).toNumber()] = _506_v13;
      _nw57[(new BigNumber(17)).toNumber()] = _506_v13;
      _507_v15 = _nw57;
      let _515_v16;
      _515_v16 = _dafny.Set.fromElements(new BigNumber((_500_v7).cardinality()));
      _507_v15 = (((_515_v16).contains((_dafny.ZERO).minus(_490_v1))) ? (_507_v15) : (_507_v15));
      r0 = _503_v10;
      r1 = ((_dafny.ZERO).minus(_490_v1)).multipliedBy((_dafny.ZERO).minus(_490_v1));
      return [r0, r1];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _516_v0;
      let _nw58 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _516_v0 = _nw58;
      let _517_v1;
      _517_v1 = _dafny.MultiSet.fromElements(_516_v0);
      let _518_v2;
      _518_v2 = new BigNumber(240);
      let _519_v3;
      _519_v3 = _dafny.Map.Empty.slice().updateUnsafe(_518_v2,_516_v0);
      let _520_v4;
      _520_v4 = _module.D0.create_DC0(p0);
      let _521_v5;
      _521_v5 = new _dafny.CodePoint('a'.codePointAt(0));
      let _522_v6;
      _522_v6 = _dafny.MultiSet.fromElements(_521_v5, _521_v5, _521_v5, _521_v5);
      let _523_v7;
      _523_v7 = _dafny.Seq.UnicodeFromString("cudjrcpt");
      let _524_v8;
      _524_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_523_v7).length),_dafny.MultiSet.fromElements(_516_v0, _516_v0, _516_v0, _516_v0, _516_v0));
      let _525_v9;
      let _nw59 = Array((new BigNumber(28)).toNumber());
      _nw59[(_dafny.ZERO).toNumber()] = _517_v1;
      _nw59[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements((((_519_v3).contains((_this).fm4(_520_v4, _518_v2, new BigNumber((_522_v6).cardinality()), globalState))) ? ((_519_v3).get((_this).fm4(_520_v4, _518_v2, new BigNumber((_522_v6).cardinality()), globalState))) : (_516_v0)));
      _nw59[(new BigNumber(2)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(3)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(4)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(5)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(6)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(_516_v0, _516_v0);
      _nw59[(new BigNumber(8)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(9)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(10)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(11)).toNumber()] = ((_517_v1).update(_516_v0, _module.__default.abs(new BigNumber(892)))).Difference(_517_v1);
      _nw59[(new BigNumber(12)).toNumber()] = (_517_v1).Intersect(_517_v1);
      _nw59[(new BigNumber(13)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(14)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(15)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(16)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(17)).toNumber()] = (_517_v1).Union(_517_v1);
      _nw59[(new BigNumber(18)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(19)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(20)).toNumber()] = (_517_v1).Union(_dafny.MultiSet.fromElements(_516_v0, _516_v0, _516_v0, _516_v0, _516_v0));
      _nw59[(new BigNumber(21)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(22)).toNumber()] = (_517_v1).Difference(_517_v1);
      _nw59[(new BigNumber(23)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(24)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(25)).toNumber()] = _517_v1;
      _nw59[(new BigNumber(26)).toNumber()] = (((_524_v8).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_518_v2)))) ? ((_524_v8).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_518_v2)))) : (_dafny.MultiSet.fromElements(_516_v0, _516_v0, _516_v0)));
      _nw59[(new BigNumber(27)).toNumber()] = _517_v1;
      _525_v9 = _nw59;
      let _index63 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_525_v9).length));
      (_525_v9)[_index63] = _517_v1;
      let _index64 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_525_v9).length));
      (_525_v9)[_index64] = _517_v1;
      let _526_i0;
      _526_i0 = _dafny.ZERO;
      L2: {
        while (!((_522_v6).IsProperSubsetOf((_dafny.MultiSet.fromElements(_521_v5)).Union(_dafny.MultiSet.fromElements(_521_v5))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_526_i0)) {
              break L2;
            }
            _526_i0 = (_526_i0).plus(_dafny.ONE);
            let _527_v10;
            _527_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_518_v2);
            _527_v10 = (_527_v10).update((_this).f27, _518_v2);
            let _528_v11;
            let _nw60 = new _module.C0();
            _nw60.__ctor(p0);
            _528_v11 = _nw60;
            let _index65 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_516_v0).length));
            (_516_v0)[_index65] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).fm1(globalState)), _518_v2));
            let _index66 = _module.__default.safeIndex(new BigNumber(433), new BigNumber((_516_v0).length));
            (_516_v0)[_index66] = _518_v2;
            let _529_v12;
            let _init13 = ((_530_v0) => function (_531_i1) {
              return _module.__default.safeModuloInt(_531_i1, (_530_v0)[_module.__default.safeIndex(new BigNumber(433), new BigNumber((_530_v0).length))]);
            })(_516_v0);
            let _nw61 = Array((new BigNumber(23)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw61.length); _i0_13++) {
              _nw61[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _529_v12 = _nw61;
            let _532_v13;
            _532_v13 = _module.D3.create_DC10(_529_v12);
            _532_v13 = _module.D3.create_DC10(_529_v12);
          }
        }
      }
      let _533_v14;
      _533_v14 = _dafny.Set.fromElements(_518_v2);
      let _534_v15;
      _534_v15 = _dafny.MultiSet.fromElements(_533_v14);
      let _535_v16;
      _535_v16 = _dafny.MultiSet.fromElements((_this).fm17(new BigNumber((_523_v7).length), (_dafny.ZERO).minus(_518_v2), _518_v2, _534_v15, globalState), (_this).f27, (_this).f26, !(!(false)));
      (globalState).f12 = ((new BigNumber((_523_v7).length)).multipliedBy((((_535_v16).contains(!(p0))) ? ((_535_v16).get(!(p0))) : (_518_v2)))).minus((_this).fm1(globalState));
      let _536_v17;
      _536_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f26);
      (_this).m12(_518_v2, new BigNumber((_536_v17).length), globalState);
      let _537_v18;
      let _nw62 = new _module.C1();
      _nw62.__ctor(_518_v2, new BigNumber((_523_v7).length), p0, (_this).f26);
      _537_v18 = _nw62;
      let _538_v19;
      _538_v19 = _module.D7.create_DC18(_537_v18);
      let _source12 = _538_v19;
      if (_source12.is_DC19) {
        let _539___mcc_h0 = (_source12).cf36;
        let _540___mcc_h1 = (_source12).cf37;
        let _541___mcc_h2 = (_source12).cf38;
        let _542___mcc_h3 = (_source12).cf39;
        let _543_cf39 = _542___mcc_h3;
        let _544_cf38 = _541___mcc_h2;
        let _545_cf37 = _540___mcc_h1;
        let _546_cf36 = _539___mcc_h0;
        (globalState).f15 = (_537_v18).f26;
        (globalState).f12 = (_518_v2).plus((_518_v2).minus(new BigNumber(648)));
        let _547_v20;
        _547_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f26)).length),_546_cf36);
        let _548_v21;
        let _nw63 = Array((new BigNumber(28)).toNumber()).fill([]);
        _548_v21 = _nw63;
        let _549_v22;
        _549_v22 = _dafny.Map.Empty.slice().updateUnsafe(_548_v21,(_537_v18).f27);
        let _550_v23;
        _550_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_549_v22).contains(_548_v21)) ? ((_549_v22).get(_548_v21)) : ((_537_v18).f27)),_536_v17);
        _546_cf36 = (_537_v18).fm4(_520_v4, (new BigNumber((_547_v20).length)).multipliedBy(_518_v2), new BigNumber((_550_v23).length), globalState);
        _521_v5 = _521_v5;
      } else {
        let _551___mcc_h4 = (_source12).cf35;
        let _552_cf35 = _551___mcc_h4;
        let _553_v24;
        let _init14 = ((_554_v2, _555_v7, _556_v18) => function (_557_i2) {
          return (_module.D5.create_DC14(_554_v2, new BigNumber((_555_v7).length), (_556_v18).f27)).dtor_cf27;
        })(_518_v2, _523_v7, _537_v18);
        let _nw64 = Array((new BigNumber(18)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw64.length); _i0_14++) {
          _nw64[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _553_v24 = _nw64;
        let _rhs55 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("htwwcultw"), _dafny.Seq.UnicodeFromString("jlkne"));
        let _rhs56 = _553_v24;
        let _lhs54 = globalState;
        _lhs54.f6 = _rhs55;
        _553_v24 = _rhs56;
        (globalState).f12 = (_518_v2).multipliedBy(new BigNumber(970));
        let _558_v25;
        let _init15 = ((_559_v14) => function (_560_i3) {
          return _559_v14;
        })(_533_v14);
        let _nw65 = Array((new BigNumber(19)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw65.length); _i0_15++) {
          _nw65[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _558_v25 = _nw65;
        let _index67 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_558_v25).length));
        (_558_v25)[_index67] = function () {
          let _coll50 = new _dafny.Set();
          for (const _compr_50 of (_dafny.MultiSet.FromArray(_module.__default.fm23(_521_v5, globalState))).Elements) {
            let _561_v26 = _compr_50;
            if ((_dafny.MultiSet.FromArray(_module.__default.fm23(_521_v5, globalState))).contains(_561_v26)) {
              _coll50.add(_module.__default.safeDivisionInt(_561_v26, new BigNumber((_dafny.Set.fromElements(false)).length)));
            }
          }
          return _coll50;
        }();
        let _index68 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_558_v25).length));
        (_558_v25)[_index68] = _533_v14;
        (globalState).f16 = (_dafny.ZERO).minus(_518_v2);
      }
      _521_v5 = (((_537_v18).fm2(p0, _518_v2, globalState)) ? (_521_v5) : (((p0) ? (_521_v5) : (_521_v5))));
      return;
    }
    m11(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _562_v0;
      _562_v0 = _module.D0.create_DC0((_this).f26);
      let _563_v1;
      _563_v1 = _dafny.Seq.UnicodeFromString("euqbmr");
      let _564_v2;
      _564_v2 = _dafny.Map.Empty.slice().updateUnsafe(_563_v1,(_this).f27);
      let _565_v3;
      _565_v3 = new BigNumber(671);
      let _566_v4;
      let _nw66 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _566_v4 = _nw66;
      let _567_v5;
      _567_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(function (_pat_let4_0) {
        return function (_568_dt__update__tmp_h0) {
          return function (_pat_let5_0) {
            return function (_569_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_569_dt__update_hcf0_h0);
            }(_pat_let5_0);
          }((_this).f26);
        }(_pat_let4_0);
      }(_562_v0), new BigNumber((_564_v2).length), _565_v3, globalState),_566_v4);
      _567_v5 = (_567_v5).update(_565_v3, _566_v4);
      let _570_v6;
      _570_v6 = _dafny.Map.Empty.slice().updateUnsafe(_563_v1,_565_v3);
      let _hi5 = _565_v3;
      for (let _571_i0 = new BigNumber((_570_v6).length); _571_i0.isLessThan(_hi5); _571_i0 = _571_i0.plus(_dafny.ONE)) {
        let _572_v7;
        _572_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24(_563_v1, (_this).f26, _565_v3, globalState),(_this).f27);
        _572_v7 = (_572_v7).Merge(_572_v7);
        let _573_v8;
        _573_v8 = new _dafny.CodePoint('n'.codePointAt(0));
        let _574_v9;
        let _575_v10;
        let _576_v11;
        let _577_v12;
        let _out21;
        let _out22;
        let _out23;
        let _out24;
        let _outcollector7 = _module.__default.m0(_573_v8, globalState);
        _out21 = _outcollector7[0];
        _out22 = _outcollector7[1];
        _out23 = _outcollector7[2];
        _out24 = _outcollector7[3];
        _574_v9 = _out21;
        _575_v10 = _out22;
        _576_v11 = _out23;
        _577_v12 = _out24;
        let _index69 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_566_v4).length));
        (_566_v4)[_index69] = _565_v3;
        let _578_v13;
        _578_v13 = _module.D6.create_DC16(_dafny.Seq.Concat(_563_v1, _563_v1));
        let _579_v14;
        _579_v14 = _module.D2.create_DC9(_571_i0, _571_i0, (_this).f27);
        let _580_v15;
        _580_v15 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(_565_v3, new BigNumber(201)));
        let _index70 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_566_v4).length));
        let _rhs57 = _571_i0;
        let _rhs58 = _571_i0;
        let _rhs59 = ((!((_579_v14).dtor_cf19).isEqualTo(_565_v3)) ? (_578_v13) : (_module.D6.create_DC16(_563_v1)));
        let _rhs60 = (_this).fm17(_571_i0, (_dafny.ZERO).minus(_571_i0), _565_v3, _580_v15, globalState);
        let _lhs55 = globalState;
        let _lhs56 = _566_v4;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_566_v4).length));
        let _lhs58 = globalState;
        _lhs55.f12 = _rhs57;
        _lhs56[_lhs57] = _rhs58;
        _578_v13 = _rhs59;
        _lhs58.f2 = _rhs60;
        let _581_v16;
        let _init16 = function (_582_i1) {
          return (_this).f27;
        };
        let _nw67 = Array((new BigNumber(13)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw67.length); _i0_16++) {
          _nw67[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _581_v16 = _nw67;
        let _583_v17;
        _583_v17 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f27),(_this).f27);
        _581_v16 = (((((_583_v17).contains(_576_v11)) ? ((_583_v17).get(_576_v11)) : (true))) ? (_581_v16) : (p1));
      }
      let _584_i2;
      _584_i2 = _dafny.ZERO;
      L3: {
        while ((false) === ((_this).f26)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_584_i2)) {
              break L3;
            }
            _584_i2 = (_584_i2).plus(_dafny.ONE);
            let _585_v18;
            _585_v18 = new _dafny.CodePoint('g'.codePointAt(0));
            let _586_v19;
            _586_v19 = _dafny.Set.fromElements((_this).f27);
            let _587_v20;
            _587_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_585_v18, _585_v18), _563_v1)).length),(_586_v19).Intersect(_586_v19));
            _587_v20 = _587_v20;
            (globalState).f22 = (_this).f26;
            (globalState).f12 = _565_v3;
            (globalState).f16 = new BigNumber(308);
          }
        }
      }
      let _588_v21;
      _588_v21 = _dafny.MultiSet.fromElements(_565_v3);
      let _589_v22;
      _589_v22 = _dafny.Set.fromElements((_this).f27, false);
      let _index71 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length));
      (p1)[_index71] = (_588_v21).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_589_v22).length), _565_v3));
      let _590_v23;
      _590_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,false);
      let _591_v24;
      _591_v24 = _dafny.MultiSet.fromElements((((_590_v23).contains((_this).f27)) ? ((_590_v23).get((_this).f27)) : (false)));
      let _592_v25;
      let _init17 = function (_593_i3) {
        return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of((_this).f27)).length));
      };
      let _nw68 = Array((new BigNumber(6)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw68.length); _i0_17++) {
        _nw68[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _592_v25 = _nw68;
      let _594_v26;
      _594_v26 = _dafny.Map.Empty.slice().updateUnsafe(_565_v3,_591_v24);
      let _595_v27;
      _595_v27 = _dafny.Map.Empty.slice().updateUnsafe(_589_v22,_592_v25);
      let _index72 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length));
      let _rhs61 = (_this).f26;
      let _rhs62 = (((_594_v26).contains(_565_v3)) ? ((_594_v26).get(_565_v3)) : (_dafny.MultiSet.fromElements(true, (_this).f27)));
      let _rhs63 = ((_this).f27) && ((_this).f26);
      let _rhs64 = ((true) ? (_592_v25) : ((((_595_v27).contains(_dafny.Set.fromElements((_this).f27))) ? ((_595_v27).get(_dafny.Set.fromElements((_this).f27))) : (_592_v25))));
      let _rhs65 = (_dafny.ZERO).minus(_565_v3);
      let _lhs59 = p1;
      let _lhs60 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length));
      let _lhs61 = globalState;
      _lhs59[_lhs60] = _rhs61;
      _591_v24 = _rhs62;
      _lhs61.f15 = _rhs63;
      _592_v25 = _rhs64;
      _565_v3 = _rhs65;
      let _596_v28;
      _596_v28 = _dafny.Set.fromElements(_565_v3);
      let _597_i4;
      _597_i4 = _dafny.ZERO;
      L4: {
        while (!(_596_v28).contains((_565_v3).plus(_565_v3))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_597_i4)) {
              break L4;
            }
            _597_i4 = (_597_i4).plus(_dafny.ONE);
            let _index73 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_566_v4).length));
            (_566_v4)[_index73] = _565_v3;
            let _598_v29;
            _598_v29 = _589_v22;
            let _599_v30;
            _599_v30 = new _dafny.CodePoint('c'.codePointAt(0));
            let _600_v31;
            _600_v31 = _dafny.Seq.of((_this).f26, false, true);
            let _601_v32;
            _601_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_600_v31);
            let _index74 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_566_v4).length));
            let _rhs66 = ((_this).fm4(_562_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))],_599_v30)).length), new BigNumber((_601_v32).length), globalState)).minus(_565_v3);
            let _rhs67 = _565_v3;
            let _rhs68 = _598_v29;
            let _lhs62 = globalState;
            let _lhs63 = _566_v4;
            let _lhs64 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_566_v4).length));
            _lhs62.f16 = _rhs66;
            _lhs63[_lhs64] = _rhs67;
            _598_v29 = _rhs68;
            let _602_v33;
            let _nw69 = new _module.C1();
            _nw69.__ctor(_565_v3, _565_v3, (_this).f27, (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]);
            _602_v33 = _nw69;
            let _603_v34;
            _603_v34 = _dafny.Seq.of(_602_v33, _602_v33, _602_v33);
            let _rhs69 = _dafny.Seq.IsPrefixOf(_module.__default.fm22(_565_v3, (_this).f26, globalState), _dafny.Seq.Concat(_563_v1, _563_v1));
            let _rhs70 = (_603_v34)[_module.__default.safeIndex(_565_v3, new BigNumber((_603_v34).length))];
            let _rhs71 = _dafny.Seq.contains(_563_v1, _599_v30);
            let _rhs72 = _599_v30;
            let _lhs65 = globalState;
            let _lhs66 = globalState;
            _lhs65.f22 = _rhs69;
            _602_v33 = _rhs70;
            _lhs66.f22 = _rhs71;
            _599_v30 = _rhs72;
            let _604_v35;
            let _nw70 = new _module.C0();
            _nw70.__ctor((p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]);
            _604_v35 = _nw70;
            let _605_v36;
            let _nw71 = new _module.C1();
            _nw71.__ctor((_602_v33).f36, _565_v3, (_604_v35).f32, (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]);
            _605_v36 = _nw71;
          }
        }
      }
      let _hi6 = (_this).fm1(globalState);
      for (let _606_i5 = _565_v3; _606_i5.isLessThan(_hi6); _606_i5 = _606_i5.plus(_dafny.ONE)) {
        (globalState).f19 = (_563_v1)[_module.__default.safeIndex((((_this).f27) ? (_565_v3) : (_606_i5)), new BigNumber((_563_v1).length))];
        let _607_v37;
        _607_v37 = new _dafny.CodePoint('n'.codePointAt(0));
        if ((new BigNumber(-287)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("lbthv"), _module.__default.safeIndex(_606_i5, new BigNumber((_dafny.Seq.UnicodeFromString("lbthv")).length)), _607_v37), _563_v1)).length))) {
          let _608_v38;
          _608_v38 = _dafny.Map.Empty.slice().updateUnsafe(_606_i5,(_this).f27);
          let _609_v39;
          _609_v39 = _module.D2.create_DC9(new BigNumber((_608_v38).length), _565_v3, false);
          let _610_v40;
          _610_v40 = _dafny.Seq.of(_609_v39, _609_v39);
          let _611_v41;
          let _nw72 = new _module.C1();
          _nw72.__ctor(new BigNumber(133), _565_v3, (_this).fm16(_606_i5, _610_v40, new BigNumber((_dafny.Set.fromElements(_606_i5)).length), new BigNumber((_563_v1).length), globalState), (_this).f26);
          _611_v41 = _nw72;
          let _612_v42;
          _612_v42 = _dafny.Map.Empty.slice().updateUnsafe((_611_v41).f36,_dafny.Seq.of(true, (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]));
          let _613_v43;
          _613_v43 = _dafny.Seq.of(true);
          let _614_v44;
          _614_v44 = _dafny.Map.Empty.slice().updateUnsafe((((_612_v42).contains((_611_v41).f36)) ? ((_612_v42).get((_611_v41).f36)) : (_613_v43)),p1);
          let _rhs73 = ((_module.D9.create_DC21(_614_v44)).dtor_cf41).Merge(_614_v44);
          let _rhs74 = ((((_this).f26) ? (_606_i5) : (_565_v3))).multipliedBy(_module.__default.safeDivisionInt((_611_v41).f36, _606_i5));
          let _lhs67 = globalState;
          _614_v44 = _rhs73;
          _lhs67.f12 = _rhs74;
          let _615_v45;
          let _nw73 = new _module.C1();
          _nw73.__ctor(_611_v41.f37, (_611_v41.f37).minus(_606_i5), !((_this).f27), (false) && ((_this).f26));
          _615_v45 = _nw73;
          let _616_v46;
          _616_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_611_v41).f36);
          _590_v23 = (_590_v23).update(!(new BigNumber((_616_v46).length)).isEqualTo(_611_v41.f37), (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]);
          let _617_v47;
          let _nw74 = Array((new BigNumber(21)).toNumber()).fill(_module.D0.Default());
          _617_v47 = _nw74;
          let _618_v48;
          _618_v48 = _dafny.Map.Empty.slice().updateUnsafe(_607_v37,_617_v47);
          _618_v48 = (_618_v48).update(_607_v37, _617_v47);
        } else {
          let _619_v49;
          _619_v49 = _dafny.MultiSet.fromElements(_596_v28);
          let _620_v50;
          _620_v50 = _dafny.Seq.of((_this).fm17(_565_v3, _606_i5, _565_v3, _619_v49, globalState));
          (globalState).f16 = (_this).fm4(_562_v0, new BigNumber((_dafny.Seq.Concat((_this).fm3(_563_v1, _620_v50, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_565_v3,(p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))])).length), globalState), _620_v50)).length), new BigNumber(864), globalState);
          (globalState).f16 = (_dafny.ZERO).minus((_this).fm4(_562_v0, _565_v3, new BigNumber(400), globalState));
          (globalState).f12 = _module.__default.safeModuloInt((_dafny.ZERO).minus((((p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))]) ? (_565_v3) : (_565_v3))), new BigNumber(((_596_v28).Intersect(_596_v28)).length));
          (globalState).f12 = _606_i5;
          (globalState).f24 = (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))];
        }
        let _621_v51;
        _621_v51 = _dafny.Seq.of(!(false), (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))], true);
        let _622_v52;
        _622_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_621_v51, _dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))], (p1)[_module.__default.safeIndex(new BigNumber(156), new BigNumber((p1).length))])),(_this).f26);
        _622_v52 = _622_v52;
        _621_v51 = _dafny.Seq.Concat(_dafny.Seq.of((_this).f27), _dafny.Seq.Concat(_621_v51, _dafny.Seq.of(!((_this).f26))));
      }
      r0 = _dafny.Seq.UnicodeFromString("bmiwf");
      return r0;
    }
    m12(p0, p1, globalState) {
      let _this = this;
      let _623_v0;
      _623_v0 = _dafny.MultiSet.fromElements(p1);
      let _624_v1;
      let _nw75 = new _module.C1();
      _nw75.__ctor(p1, p1, (((_this).f27) ? ((_this).f27) : ((_this).f27)), (_dafny.MultiSet.fromElements(p1, p1)).IsDisjointFrom(_623_v0));
      _624_v1 = _nw75;
      let _625_v2;
      let _init18 = function (_626_i0) {
        return (_this).f26;
      };
      let _nw76 = Array((new BigNumber(17)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw76.length); _i0_18++) {
        _nw76[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _625_v2 = _nw76;
      let _627_v3;
      _627_v3 = _dafny.Seq.of((_this).f26, (_this).f26, (_this).f27);
      let _628_v4;
      _628_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_627_v3);
      let _629_v5;
      _629_v5 = _module.D5.create_DC13(_628_v4);
      let _rhs75 = !(!_dafny.areEqual(_dafny.Seq.of((_this).f27, (_this).f27, (_this).f27, (_this).f27), _dafny.Seq.of((_this).f27, (_this).f26))) || ((_this).f26);
      let _rhs76 = _625_v2;
      let _rhs77 = (_this).f27;
      let _rhs78 = _629_v5;
      let _lhs68 = globalState;
      let _lhs69 = globalState;
      _lhs68.f24 = _rhs75;
      _625_v2 = _rhs76;
      _lhs69.f22 = _rhs77;
      _629_v5 = _rhs78;
      let _630_v6;
      _630_v6 = _dafny.Seq.UnicodeFromString("adlk");
      let _hi7 = _module.__default.safeDivisionInt(_624_v1.f37, (_dafny.ZERO).minus(new BigNumber((_627_v3).length)));
      for (let _631_i1 = new BigNumber((_630_v6).length); _631_i1.isLessThan(_hi7); _631_i1 = _631_i1.plus(_dafny.ONE)) {
        (globalState).f12 = _631_i1;
        let _632_v7;
        let _init19 = ((_633_v1) => function (_634_i2) {
          return _module.__default.safeDivisionInt(_634_i2, (_633_v1).f36);
        })(_624_v1);
        let _nw77 = Array((new BigNumber(12)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw77.length); _i0_19++) {
          _nw77[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _632_v7 = _nw77;
        let _index75 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_632_v7).length));
        (_632_v7)[_index75] = _624_v1.f37;
        let _635_v8;
        _635_v8 = new _dafny.CodePoint('s'.codePointAt(0));
        let _636_v9;
        _636_v9 = _dafny.Seq.of(_625_v2, _625_v2);
        let _index76 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_632_v7).length));
        let _rhs79 = _dafny.Seq.update(_630_v6, _module.__default.safeIndex(new BigNumber((_627_v3).length), new BigNumber((_630_v6).length)), _635_v8);
        let _rhs80 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_636_v9).length)));
        let _rhs81 = (_this).f26;
        let _rhs82 = (_this).f27;
        let _rhs83 = !((_624_v1).fm2((_this).f26, ((true) ? (_624_v1.f37) : (_624_v1.f37)), globalState));
        let _lhs70 = globalState;
        let _lhs71 = _632_v7;
        let _lhs72 = _module.__default.safeIndex(new BigNumber(313), new BigNumber((_632_v7).length));
        let _lhs73 = globalState;
        let _lhs74 = globalState;
        let _lhs75 = globalState;
        _lhs70.f6 = _rhs79;
        _lhs71[_lhs72] = _rhs80;
        _lhs73.f22 = _rhs81;
        _lhs74.f24 = _rhs82;
        _lhs75.f2 = _rhs83;
        let _637_v10;
        _637_v10 = _dafny.Set.fromElements(!((_this).f27), (_this).f27, !(false));
        _637_v10 = ((_637_v10).Intersect(_637_v10)).Intersect((_637_v10).Union(_637_v10));
        let _638_v11;
        let _nw78 = Array((new BigNumber(27)).toNumber());
        _nw78[(_dafny.ZERO).toNumber()] = _625_v2;
        _nw78[(_dafny.ONE).toNumber()] = _625_v2;
        _nw78[(new BigNumber(2)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(3)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(4)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(5)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(6)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(7)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(8)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(9)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(10)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(11)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(12)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(13)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(14)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(15)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(16)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(17)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(18)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(19)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(20)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(21)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(22)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(23)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(24)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(25)).toNumber()] = _625_v2;
        _nw78[(new BigNumber(26)).toNumber()] = _625_v2;
        _638_v11 = _nw78;
        let _639_v12;
        _639_v12 = _dafny.Seq.of(_638_v11, _638_v11, _638_v11);
        let _rhs84 = ((_this).f27) === ((_this).f26);
        let _rhs85 = (_639_v12)[_module.__default.safeIndex((_624_v1).f36, new BigNumber((_639_v12).length))];
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        _lhs76.f22 = _rhs84;
        _lhs77.f5 = _rhs85;
      }
      let _index77 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_625_v2).length));
      (_625_v2)[_index77] = (_this).f27;
      let _index78 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_625_v2).length));
      (_625_v2)[_index78] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(815)), function (_640_i3) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }), _630_v6), _630_v6);
      (globalState).f16 = (_dafny.ZERO).minus(p0);
      (_624_v1).f37 = new BigNumber((_dafny.Seq.Concat(_630_v6, _630_v6)).length);
      return;
    }
    get f35() {
      let _this = this;
      return _this._f35;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f26 = false;
      this._f27 = false;
      this._f34 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f34, f26, f27) {
      let _this = this;
      (_this)._f34 = f34;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm1(globalState) {
      let _this = this;
      if ((_this).f26) {
        return (new BigNumber(((_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("yoqe"))).dtor_cf29).length)).plus(new BigNumber((_dafny.Seq.UnicodeFromString("jmfpgf")).length));
      } else {
        return new BigNumber((_dafny.Seq.of((_this).f27, (_this).f26, (_this).f26)).length);
      }
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    fm12(p0, globalState) {
      let _this = this;
      return ((((_this).f27) ? (new BigNumber(84)) : (new BigNumber(719)))).plus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of((_this).f26)).length), new BigNumber(676)));
    };
    fm13(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(109)), function (_641_i0) {
        return new BigNumber(((_module.D6.create_DC16(_dafny.Seq.UnicodeFromString("r"))).dtor_cf29).length);
      })).length)).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(956),(_this).f26))).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('k'.codePointAt(0)))).length))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _642_v0;
      let _nw79 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _642_v0 = _nw79;
      let _643_v1;
      _643_v1 = new BigNumber(281);
      let _644_v2;
      _644_v2 = _dafny.Map.Empty.slice().updateUnsafe(_643_v1,_642_v0);
      let _645_v3;
      let _nw80 = Array((new BigNumber(20)).toNumber());
      _nw80[(_dafny.ZERO).toNumber()] = _642_v0;
      _nw80[(_dafny.ONE).toNumber()] = _642_v0;
      _nw80[(new BigNumber(2)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(3)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(4)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(5)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(6)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(7)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(8)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(9)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(10)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(11)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(12)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(13)).toNumber()] = (((_644_v2).contains(_643_v1)) ? ((_644_v2).get(_643_v1)) : (_642_v0));
      _nw80[(new BigNumber(14)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(15)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(16)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(17)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(18)).toNumber()] = _642_v0;
      _nw80[(new BigNumber(19)).toNumber()] = _642_v0;
      _645_v3 = _nw80;
      let _index79 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_645_v3).length));
      (_645_v3)[_index79] = _642_v0;
      let _646_v4;
      _646_v4 = _dafny.Seq.of(_642_v0, _642_v0);
      let _index80 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_645_v3).length));
      (_645_v3)[_index80] = (_646_v4)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_646_v4).length))];
      let _647_v5;
      _647_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f26);
      let _648_v6;
      _648_v6 = _dafny.Seq.of(_647_v5, _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f27), _647_v5, _647_v5, _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f26));
      let _649_v7;
      _649_v7 = _dafny.Seq.of(_643_v1, _643_v1, _643_v1);
      let _650_v8;
      _650_v8 = _dafny.Map.Empty.slice().updateUnsafe(_648_v6,_649_v7);
      let _651_v9;
      _651_v9 = _dafny.Seq.UnicodeFromString("qgla");
      let _652_v10;
      _652_v10 = _dafny.Map.Empty.slice().updateUnsafe(_651_v9,_648_v6);
      _650_v8 = (_650_v8).update(_dafny.Seq.Concat((((_652_v10).contains(_651_v9)) ? ((_652_v10).get(_651_v9)) : (_648_v6)), _dafny.Seq.of(_647_v5)), _dafny.Seq.Concat(_649_v7, _649_v7));
      let _653_v11;
      _653_v11 = _dafny.Seq.of((_this).f27, (_this).f26, p0, (_643_v1).isEqualTo(new BigNumber((_651_v9).length)), p0);
      _653_v11 = _dafny.Seq.Concat(_653_v11, _653_v11);
      let _654_v12;
      _654_v12 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0))));
      _643_v1 = (_dafny.ZERO).minus((_this).fm13(_dafny.Seq.Concat(_654_v12, _654_v12), (_651_v9)[_module.__default.safeIndex(_643_v1, new BigNumber((_651_v9).length))], (_643_v1).plus(_643_v1), globalState));
      _651_v9 = _dafny.Seq.Concat(_651_v9, _651_v9);
      let _655_v13;
      _655_v13 = _dafny.Set.fromElements((_this).f27);
      let _656_v14;
      _656_v14 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_655_v13).length)),_643_v1);
      let _hi8 = _643_v1;
      for (let _657_i0 = (_dafny.ZERO).minus((((_656_v14).contains(_643_v1)) ? ((_656_v14).get(_643_v1)) : (_643_v1))); _657_i0.isLessThan(_hi8); _657_i0 = _657_i0.plus(_dafny.ONE)) {
        (globalState).f15 = (_657_i0).isLessThan(_module.__default.safeDivisionInt(_657_i0, new BigNumber((_653_v11).length)));
        let _658_v15;
        _658_v15 = _dafny.Map.Empty.slice().updateUnsafe(_657_i0,(_this).f27);
        let _659_v16;
        _659_v16 = _dafny.MultiSet.fromElements(_643_v1);
        let _660_v17;
        _660_v17 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_658_v15).length)).plus(new BigNumber((_659_v16).cardinality())),_651_v9);
        let _661_v18;
        _661_v18 = new _dafny.CodePoint('e'.codePointAt(0));
        _660_v17 = (_660_v17).update(new BigNumber((_dafny.Seq.of(_module.__default.fm14((_this).f26, _657_i0, _643_v1, globalState), _661_v18, _661_v18, _661_v18)).length), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), function (_662_i1) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _651_v9));
        _643_v1 = new BigNumber((_dafny.Seq.Concat(((p0) ? (_dafny.Seq.of(_657_i0)) : (_dafny.Seq.of(_657_i0))), _dafny.Seq.Concat(_module.__default.fm15(globalState), _649_v7))).length);
        let _663_v19;
        let _nw81 = new _module.C1();
        _nw81.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("ylv")).length), new BigNumber((_651_v9).length), p0, (_this).f27);
        _663_v19 = _nw81;
        let _664_v20;
        _664_v20 = _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC18(_663_v19)).dtor_cf35,_657_i0);
        let _665_v21;
        let _nw82 = Array((new BigNumber(12)).toNumber());
        _nw82[(_dafny.ZERO).toNumber()] = _664_v20;
        _nw82[(_dafny.ONE).toNumber()] = (_664_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe(_663_v19,new BigNumber((_646_v4).length)));
        _nw82[(new BigNumber(2)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(3)).toNumber()] = (_664_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe(_663_v19,_643_v1));
        _nw82[(new BigNumber(4)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(5)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_663_v19,new BigNumber((_659_v16).cardinality()))).Merge(_664_v20);
        _nw82[(new BigNumber(7)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(8)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(9)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(10)).toNumber()] = _664_v20;
        _nw82[(new BigNumber(11)).toNumber()] = _664_v20;
        _665_v21 = _nw82;
        let _index81 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_665_v21).length));
        (_665_v21)[_index81] = _dafny.Map.Empty.slice().updateUnsafe(_663_v19,new BigNumber(196));
        let _index82 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_665_v21).length));
        (_665_v21)[_index82] = (_dafny.Map.Empty.slice().updateUnsafe(_663_v19,_657_i0)).Merge((_664_v20).Merge(_664_v20));
      }
      return;
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f26 = false;
      this._f27 = false;
      this._f33 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f33, f26, f27) {
      let _this = this;
      (_this)._f33 = f33;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm1(globalState) {
      let _this = this;
      let _source13 = _dafny.Set.fromElements((_this).f27);
      let _666___mcc_h0 = _source13;
      let _667_cf23 = _666___mcc_h0;
      return new BigNumber((_dafny.Seq.of(!((_this).f26), (_this).f26, (_this).f26)).length);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return (_this).f27;
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _668_v0;
      _668_v0 = new BigNumber(32);
      (globalState).f12 = _668_v0;
      let _hi9 = _668_v0;
      for (let _669_i0 = _668_v0; _669_i0.isLessThan(_hi9); _669_i0 = _669_i0.plus(_dafny.ONE)) {
        if (true) {
          let _670_v1;
          _670_v1 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_669_i0));
          let _671_v2;
          _671_v2 = _dafny.MultiSet.fromElements((_dafny.MultiSet.fromElements(new BigNumber((p2).length))).IsProperSubsetOf(_670_v1), !((_this).f26) || (p0));
          (globalState).f12 = (((_671_v2).contains((_this).f26)) ? ((_671_v2).get((_this).f26)) : ((_669_i0).plus(_668_v0)));
          let _672_v3;
          let _init20 = ((_673_v0) => function (_674_i1) {
            return _module.D2.create_DC7((_this).f26, (_this).f27, _673_v0, _673_v0);
          })(_668_v0);
          let _nw83 = Array((new BigNumber(13)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw83.length); _i0_20++) {
            _nw83[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _672_v3 = _nw83;
          _672_v3 = _672_v3;
          let _675_v4;
          _675_v4 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,new BigNumber(136));
          (globalState).f12 = (_669_i0).plus((((_675_v4).contains(_669_i0)) ? ((_675_v4).get(_669_i0)) : (new BigNumber((_671_v2).cardinality()))));
          let _676_v5;
          let _nw84 = new _module.C0();
          _nw84.__ctor((_this).f27);
          _676_v5 = _nw84;
          (globalState).f15 = !(!((_669_i0).isLessThanOrEqualTo(_668_v0)));
        } else {
          let _677_v6;
          _677_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p2);
          let _678_v7;
          _678_v7 = _module.D5.create_DC13(_677_v6);
          _677_v6 = ((_678_v7).dtor_cf24).Merge(_677_v6);
          let _679_v8;
          _679_v8 = _module.D1.create_DC3(p2);
          let _680_v9;
          let _nw85 = new _module.C0();
          _nw85.__ctor(p0);
          _680_v9 = _nw85;
          let _681_v10;
          _681_v10 = _dafny.Seq.of(_680_v9);
          let _682_v11;
          _682_v11 = _dafny.Map.Empty.slice().updateUnsafe(_679_v8,_dafny.Seq.contains(_681_v10, _680_v9));
          let _683_v12;
          _683_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f27);
          let _684_v13;
          _684_v13 = _dafny.Seq.of(_682_v11);
          let _685_v15;
          _685_v15 = _dafny.MultiSet.fromElements(_679_v8, _module.D1.create_DC3(p2));
          _682_v11 = (_dafny.Map.Empty.slice().updateUnsafe(_679_v8,(((_683_v12).contains((_680_v9).f32)) ? ((_683_v12).get((_680_v9).f32)) : ((_this).f27)))).Merge(((_684_v13)[_module.__default.safeIndex(_669_i0, new BigNumber((_684_v13).length))]).Merge(function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of (_685_v15).Elements) {
              let _686_v14 = _compr_51;
              if ((_685_v15).contains(_686_v14)) {
                _coll51.push([_686_v14,(_680_v9).f32]);
              }
            }
            return _coll51;
          }()));
          let _rhs86 = _668_v0;
          let _rhs87 = _668_v0;
          let _lhs78 = globalState;
          let _lhs79 = globalState;
          _lhs78.f12 = _rhs86;
          _lhs79.f12 = _rhs87;
          let _687_v16;
          _687_v16 = _module.D5.create_DC14(new BigNumber((p2).length), _668_v0, true);
          let _688_v17;
          let _nw86 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _688_v17 = _nw86;
          let _689_v18;
          _689_v18 = _dafny.Seq.of(_688_v17, _688_v17);
          (globalState).f12 = (((_687_v16).dtor_cf26).plus(new BigNumber((_689_v18).length))).minus(_668_v0);
          let _690_v19;
          _690_v19 = _dafny.Seq.UnicodeFromString("ldus");
          (globalState).f15 = (_669_i0).isLessThan((new BigNumber((_690_v19).length)).multipliedBy(_669_i0));
        }
        let _691_v20;
        let _init21 = ((_692_v0) => function (_693_i2) {
          return (_693_i2).plus((_module.D5.create_DC14(_692_v0, _692_v0, true)).dtor_cf26);
        })(_668_v0);
        let _nw87 = Array((new BigNumber(9)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw87.length); _i0_21++) {
          _nw87[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _691_v20 = _nw87;
        let _694_v21;
        _694_v21 = new _dafny.CodePoint('o'.codePointAt(0));
        let _695_v22;
        _695_v22 = _dafny.Set.fromElements(_694_v21);
        let _696_v23;
        _696_v23 = _dafny.Map.Empty.slice().updateUnsafe(_695_v22,_668_v0);
        let _697_v24;
        _697_v24 = _dafny.Seq.UnicodeFromString("cdyf");
        let _698_v25;
        _698_v25 = _dafny.Seq.of(_668_v0, _669_i0);
        let _699_v26;
        _699_v26 = _dafny.MultiSet.fromElements(_698_v25, _698_v25);
        let _700_v27;
        let _nw88 = new _module.C3();
        _nw88.__ctor(_699_v26, p0, p0);
        _700_v27 = _nw88;
        let _701_v28;
        let _nw89 = Array((new BigNumber(14)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = (_669_i0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("yghsv")).length));
        _nw89[(_dafny.ONE).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_691_v20,new BigNumber((_696_v23).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_691_v20,_669_i0))).length);
        _nw89[(new BigNumber(2)).toNumber()] = _668_v0;
        _nw89[(new BigNumber(3)).toNumber()] = _668_v0;
        _nw89[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_668_v0);
        _nw89[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_669_i0);
        _nw89[(new BigNumber(6)).toNumber()] = _669_i0;
        _nw89[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_697_v24, _697_v24)).length);
        _nw89[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(546)), ((_702_v0) => function (_703_i3) {
          return _702_v0;
        })(_668_v0)),_700_v27)).length);
        _nw89[(new BigNumber(9)).toNumber()] = _669_i0;
        _nw89[(new BigNumber(10)).toNumber()] = _668_v0;
        _nw89[(new BigNumber(11)).toNumber()] = _669_i0;
        _nw89[(new BigNumber(12)).toNumber()] = _669_i0;
        _nw89[(new BigNumber(13)).toNumber()] = _668_v0;
        _701_v28 = _nw89;
        let _704_v29;
        _704_v29 = _module.D3.create_DC10(_701_v28);
        _691_v20 = (_704_v29).dtor_cf21;
        let _705_v30;
        _705_v30 = _module.D7.create_DC19(_669_i0, _dafny.MultiSet.fromElements(!(p0), (_this).f26), false, _dafny.Map.Empty.slice().updateUnsafe(false,_668_v0));
        let _706_v31;
        _706_v31 = _dafny.MultiSet.fromElements((_700_v27).f27);
        let _pat_let_tv2 = _694_v21;
        let _pat_let_tv3 = _706_v31;
        let _707_v32;
        let _nw90 = Array((new BigNumber(12)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = (((_this).f27) ? (_705_v30) : (_705_v30));
        _nw90[(_dafny.ONE).toNumber()] = function (_pat_let6_0) {
          return function (_708_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_711_dt__update_hcf36_h0) {
                return function (_pat_let8_0) {
                  return function (_712_dt__update_hcf38_h0) {
                    return function (_pat_let9_0) {
                      return function (_713_dt__update_hcf37_h0) {
                        return _module.D7.create_DC19(_711_dt__update_hcf36_h0, _713_dt__update_hcf37_h0, _712_dt__update_hcf38_h0, (_708_dt__update__tmp_h0).dtor_cf39);
                      }(_pat_let9_0);
                    }(_pat_let_tv3);
                  }(_pat_let8_0);
                }((_this).f26);
              }(_pat_let7_0);
            }(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_709_v21) => function (_710_i4) {
              return _709_v21;
            })(_pat_let_tv2))).length));
          }(_pat_let6_0);
        }(_705_v30);
        _nw90[(new BigNumber(2)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(3)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(4)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(5)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(6)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(7)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(8)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(9)).toNumber()] = _705_v30;
        _nw90[(new BigNumber(10)).toNumber()] = ((p0) ? (_705_v30) : (_705_v30));
        _nw90[(new BigNumber(11)).toNumber()] = _705_v30;
        _707_v32 = _nw90;
        let _714_v33;
        _714_v33 = _dafny.Map.Empty.slice().updateUnsafe((_700_v27).f26,new BigNumber(637));
        let _index83 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_707_v32).length));
        (_707_v32)[_index83] = _module.D7.create_DC19(_668_v0, _706_v31, p0, _714_v33);
        let _index84 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_707_v32).length));
        (_707_v32)[_index84] = _705_v30;
        let _715_v34;
        let _nw91 = Array((new BigNumber(4)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = (_700_v27).f26;
        _nw91[(_dafny.ONE).toNumber()] = (_700_v27).f27;
        _nw91[(new BigNumber(2)).toNumber()] = false;
        _nw91[(new BigNumber(3)).toNumber()] = (_700_v27).f26;
        _715_v34 = _nw91;
        let _index85 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_715_v34).length));
        (_715_v34)[_index85] = ((true) ? ((_700_v27).fm2((_700_v27).f26, _669_i0, globalState)) : ((_700_v27).f27));
        let _716_v35;
        let _nw92 = Array((new BigNumber(29)).toNumber()).fill(_module.D1.Default());
        _716_v35 = _nw92;
        let _717_v36;
        _717_v36 = _dafny.Map.Empty.slice().updateUnsafe(true,_716_v35);
        let _718_v37;
        _718_v37 = _dafny.Map.Empty.slice().updateUnsafe(_669_i0,_669_i0);
        let _719_v38;
        _719_v38 = _dafny.Map.Empty.slice().updateUnsafe(_717_v36,new BigNumber((_718_v37).length));
        let _index86 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_715_v34).length));
        let _rhs88 = !(!(_668_v0).isEqualTo(_669_i0)) || ((_this).f26);
        let _rhs89 = _dafny.Seq.UnicodeFromString("fxf");
        let _rhs90 = _module.__default.safeModuloInt(_668_v0, _668_v0);
        let _rhs91 = (!(_668_v0).isEqualTo((((_719_v38).contains(_717_v36)) ? ((_719_v38).get(_717_v36)) : ((_698_v25)[_module.__default.safeIndex(_669_i0, new BigNumber((_698_v25).length))])))) || ((_669_i0).isLessThanOrEqualTo(_668_v0));
        let _lhs80 = _715_v34;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_715_v34).length));
        let _lhs82 = globalState;
        let _lhs83 = globalState;
        let _lhs84 = globalState;
        _lhs80[_lhs81] = _rhs88;
        _lhs82.f6 = _rhs89;
        _lhs83.f12 = _rhs90;
        _lhs84.f22 = _rhs91;
      }
      let _720_v39;
      _720_v39 = new _dafny.CodePoint('q'.codePointAt(0));
      let _721_v40;
      _721_v40 = _dafny.Set.fromElements((_this).f26);
      let _722_v41;
      let _nw93 = Array((new BigNumber(21)).toNumber());
      _nw93[(_dafny.ZERO).toNumber()] = _720_v39;
      _nw93[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
      _nw93[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
      _nw93[(new BigNumber(3)).toNumber()] = _module.__default.fm14(p0, _668_v0, new BigNumber((_721_v40).length), globalState);
      _nw93[(new BigNumber(4)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(5)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
      _nw93[(new BigNumber(7)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(8)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(9)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(10)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(11)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(12)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(13)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
      _nw93[(new BigNumber(15)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
      _nw93[(new BigNumber(17)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(18)).toNumber()] = _720_v39;
      _nw93[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
      _nw93[(new BigNumber(20)).toNumber()] = _720_v39;
      _722_v41 = _nw93;
      let _723_v42;
      _723_v42 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,false);
      let _724_v43;
      _724_v43 = _dafny.Seq.of(_723_v42, _723_v42);
      let _725_v44;
      _725_v44 = _module.D2.create_DC7(true, false, _668_v0, new BigNumber((_dafny.MultiSet.FromArray(_724_v43)).cardinality()));
      let _726_v45;
      _726_v45 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_725_v44, _725_v44),_668_v0);
      let _727_v46;
      _727_v46 = _dafny.Seq.of(_725_v44, _725_v44);
      let _728_v47;
      _728_v47 = _dafny.Seq.UnicodeFromString("baa");
      let _729_v48;
      _729_v48 = _module.D1.create_DC4(_722_v41, _668_v0, (((_726_v45).contains(_727_v46)) ? ((_726_v45).get(_727_v46)) : (_668_v0)), _728_v47, _dafny.Seq.update(_728_v47, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_668_v0,_668_v0)).length)), new BigNumber((_728_v47).length)), _720_v39));
      let _source14 = _729_v48;
      if (_source14.is_DC4) {
        let _730___mcc_h0 = (_source14).cf7;
        let _731___mcc_h1 = (_source14).cf8;
        let _732___mcc_h2 = (_source14).cf9;
        let _733___mcc_h3 = (_source14).cf10;
        let _734___mcc_h4 = (_source14).cf11;
        let _735_cf11 = _734___mcc_h4;
        let _736_cf10 = _733___mcc_h3;
        let _737_cf9 = _732___mcc_h2;
        let _738_cf8 = _731___mcc_h1;
        let _739_cf7 = _730___mcc_h0;
        let _740_v49;
        _740_v49 = _dafny.Seq.of(_668_v0, _738_cf8, _737_cf9);
        let _741_v50;
        _741_v50 = _dafny.Seq.of(_740_v49, _740_v49, _dafny.Seq.update(_740_v49, _module.__default.safeIndex(new BigNumber((_740_v49).length), new BigNumber((_740_v49).length)), _738_cf8), _dafny.Seq.of(_668_v0), _740_v49);
        let _742_v51;
        let _nw94 = new _module.C2();
        _nw94.__ctor(_741_v50, false, (_this).f26);
        _742_v51 = _nw94;
        let _743_v52;
        _743_v52 = _dafny.Seq.of(_742_v51);
        (globalState).f12 = (((_this).f26) ? (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex((_dafny.ZERO).minus((_this).fm1(globalState)), new BigNumber((p2).length)), !((_this).f26))).length), _737_cf9)) : (new BigNumber((_dafny.Seq.Concat(_743_v52, _743_v52)).length)));
        let _744_v53;
        _744_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),_738_cf8);
        let _745_v54;
        _745_v54 = _module.D0.create_DC0((_this).f26);
        _744_v53 = (_744_v53).update(_738_cf8, (_742_v51).fm4(_745_v54, _668_v0, _737_cf9, globalState));
        (globalState).f2 = !((_this).f26);
        if (!((_this).f27)) {
          let _746_v55;
          let _nw95 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _746_v55 = _nw95;
          let _index87 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_746_v55).length));
          (_746_v55)[_index87] = _dafny.Seq.Concat(_module.__default.fm15(globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_747_cf8) => function (_748_i5) {
            return _747_cf8;
          })(_738_cf8)));
          let _index88 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_746_v55).length));
          let _rhs92 = (_dafny.ZERO).minus((_742_v51).fm4(_745_v54, _737_cf9, _668_v0, globalState));
          let _rhs93 = (p2)[_module.__default.safeIndex(_668_v0, new BigNumber((p2).length))];
          let _rhs94 = _736_cf10;
          let _rhs95 = _module.__default.fm23(_720_v39, globalState);
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = _746_v55;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_746_v55).length));
          _737_cf9 = _rhs92;
          _lhs85.f15 = _rhs93;
          _lhs86.f6 = _rhs94;
          _lhs87[_lhs88] = _rhs95;
          (globalState).f19 = _720_v39;
          (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.update(_728_v47, _module.__default.safeIndex((_dafny.ZERO).minus(_668_v0), new BigNumber((_728_v47).length)), _720_v39), _dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), ((_749_v39) => function (_750_i6) {
            return _749_v39;
          })(_720_v39)));
          let _751_v56;
          _751_v56 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("riqaxgwa"), _dafny.Seq.UnicodeFromString("jqtgis"));
          (globalState).f24 = (_742_v51).fm2((_751_v56).IsProperSubsetOf(_751_v56), _737_cf9, globalState);
          (globalState).f15 = (_this).f26;
        } else {
          let _752_v57;
          let _nw96 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _752_v57 = _nw96;
          let _753_v58;
          let _nw97 = Array((new BigNumber(11)).toNumber()).fill(_module.D3.Default());
          _753_v58 = _nw97;
          let _754_v59;
          _754_v59 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f26),_739_cf7);
          let _755_v60;
          _755_v60 = _dafny.Seq.of(_739_cf7);
          let _756_v61;
          let _nw98 = Array((new BigNumber(15)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = _722_v41;
          _nw98[(_dafny.ONE).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(2)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(3)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(4)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(5)).toNumber()] = _722_v41;
          _nw98[(new BigNumber(6)).toNumber()] = _722_v41;
          _nw98[(new BigNumber(7)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(8)).toNumber()] = _722_v41;
          _nw98[(new BigNumber(9)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(10)).toNumber()] = ((true) ? (_739_cf7) : (_722_v41));
          _nw98[(new BigNumber(11)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(12)).toNumber()] = (((_754_v59).contains(true)) ? ((_754_v59).get(true)) : (_739_cf7));
          _nw98[(new BigNumber(13)).toNumber()] = _739_cf7;
          _nw98[(new BigNumber(14)).toNumber()] = (_755_v60)[_module.__default.safeIndex(new BigNumber(-535), new BigNumber((_755_v60).length))];
          _756_v61 = _nw98;
          let _index89 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_756_v61).length));
          (_756_v61)[_index89] = _722_v41;
          let _index90 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_756_v61).length));
          let _rhs96 = ((p0) ? (_752_v57) : (_752_v57));
          let _rhs97 = _753_v58;
          let _rhs98 = _722_v41;
          let _lhs89 = _756_v61;
          let _lhs90 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_756_v61).length));
          _752_v57 = _rhs96;
          _753_v58 = _rhs97;
          _lhs89[_lhs90] = _rhs98;
          _737_cf9 = (((_this).f27) ? (_737_cf9) : ((new BigNumber((_736_cf10).length)).multipliedBy(_737_cf9)));
          let _757_v62;
          _757_v62 = _dafny.Set.fromElements(new BigNumber(-955));
          let _758_v63;
          _758_v63 = _module.D2.create_DC9(new BigNumber((_757_v62).length), _668_v0, false);
          let _759_v64;
          _759_v64 = _dafny.MultiSet.fromElements(!((_this).f27), (_this).f26, (_this).f26);
          let _760_v65;
          _760_v65 = _dafny.Seq.of(p2);
          let _761_v66;
          let _nw99 = Array((new BigNumber(12)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = _738_cf8;
          _nw99[(_dafny.ONE).toNumber()] = (_758_v63).dtor_cf18;
          _nw99[(new BigNumber(2)).toNumber()] = new BigNumber((_759_v64).cardinality());
          _nw99[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_737_cf9);
          _nw99[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_this).f27, p0, p0, false)).cardinality());
          _nw99[(new BigNumber(5)).toNumber()] = _668_v0;
          _nw99[(new BigNumber(6)).toNumber()] = _738_cf8;
          _nw99[(new BigNumber(7)).toNumber()] = _738_cf8;
          _nw99[(new BigNumber(8)).toNumber()] = new BigNumber((_735_cf11).length);
          _nw99[(new BigNumber(9)).toNumber()] = _668_v0;
          _nw99[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p2, (_760_v65)[_module.__default.safeIndex((_742_v51).fm1(globalState), new BigNumber((_760_v65).length))])).length);
          _nw99[(new BigNumber(11)).toNumber()] = _668_v0;
          _761_v66 = _nw99;
          let _index91 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_761_v66).length));
          (_761_v66)[_index91] = (_668_v0).multipliedBy(_737_cf9);
          let _index92 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_761_v66).length));
          let _rhs99 = _737_cf9;
          let _rhs100 = _761_v66;
          let _rhs101 = !(!((((_this).f26) ? (p0) : ((((_723_v42).contains(_738_cf8)) ? ((_723_v42).get(_738_cf8)) : (p0)))))) || (((_dafny.ZERO).minus(_737_cf9)).isLessThan(_738_cf8));
          let _rhs102 = p0;
          let _rhs103 = _module.__default.safeModuloInt(_737_cf9, _668_v0);
          let _lhs91 = globalState;
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          let _lhs94 = _761_v66;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_761_v66).length));
          _lhs91.f16 = _rhs99;
          _761_v66 = _rhs100;
          _lhs92.f15 = _rhs101;
          _lhs93.f2 = _rhs102;
          _lhs94[_lhs95] = _rhs103;
          let _762_v67;
          _762_v67 = _dafny.Seq.of(p0);
          _762_v67 = _762_v67;
          let _763_v68;
          _763_v68 = _module.D7.create_DC19(_668_v0, (_759_v64).update((_this).f27, _module.__default.abs(_737_cf9)), (_this).f26, _dafny.Map.Empty.slice().updateUnsafe(p0,_737_cf9));
          let _764_v69;
          _764_v69 = _dafny.MultiSet.fromElements(_738_cf8, (_dafny.ZERO).minus(_737_cf9));
          let _765_v70;
          _765_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_742_v51).fm2(p0, _668_v0, globalState));
          let _766_v71;
          let _nw100 = Array((new BigNumber(26)).toNumber());
          _nw100[(_dafny.ZERO).toNumber()] = false;
          _nw100[(_dafny.ONE).toNumber()] = false;
          _nw100[(new BigNumber(2)).toNumber()] = (_738_cf8).isLessThan(new BigNumber((_module.__default.fm22((((_744_v53).contains(_738_cf8)) ? ((_744_v53).get(_738_cf8)) : (_737_cf9)), (_this).f27, globalState)).length));
          _nw100[(new BigNumber(3)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(4)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(5)).toNumber()] = (_763_v68).dtor_cf38;
          _nw100[(new BigNumber(6)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(7)).toNumber()] = p0;
          _nw100[(new BigNumber(8)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(9)).toNumber()] = _dafny.Seq.IsPrefixOf(_728_v47, _728_v47);
          _nw100[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.FromArray(_740_v49)).IsDisjointFrom(_764_v69);
          _nw100[(new BigNumber(11)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(12)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(13)).toNumber()] = true;
          _nw100[(new BigNumber(14)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_736_cf10, _dafny.Seq.UnicodeFromString("h"));
          _nw100[(new BigNumber(15)).toNumber()] = ((_this).f26) && (p0);
          _nw100[(new BigNumber(16)).toNumber()] = (_this).f26;
          _nw100[(new BigNumber(17)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(18)).toNumber()] = (_this).f26;
          _nw100[(new BigNumber(19)).toNumber()] = (((_765_v70).contains((_this).f26)) ? ((_765_v70).get((_this).f26)) : ((_this).f27));
          _nw100[(new BigNumber(20)).toNumber()] = true;
          _nw100[(new BigNumber(21)).toNumber()] = false;
          _nw100[(new BigNumber(22)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(23)).toNumber()] = (_this).f27;
          _nw100[(new BigNumber(24)).toNumber()] = !((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f27)).equals(_765_v70));
          _nw100[(new BigNumber(25)).toNumber()] = !(p0);
          _766_v71 = _nw100;
          let _index93 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_766_v71).length));
          (_766_v71)[_index93] = false;
          let _767_v72;
          _767_v72 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_761_v66)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_761_v66).length))]);
          let _index94 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_766_v71).length));
          let _rhs104 = p0;
          let _rhs105 = (_this).fm1(globalState);
          let _rhs106 = _720_v39;
          let _rhs107 = ((((_742_v51).fm2(!((_this).f27), (_761_v66)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_761_v66).length))], globalState)) ? (p0) : ((_this).f27))) || (!(new BigNumber((_767_v72).length)).isEqualTo(_668_v0));
          let _lhs96 = globalState;
          let _lhs97 = globalState;
          let _lhs98 = globalState;
          let _lhs99 = _766_v71;
          let _lhs100 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_766_v71).length));
          _lhs96.f22 = _rhs104;
          _lhs97.f16 = _rhs105;
          _lhs98.f19 = _rhs106;
          _lhs99[_lhs100] = _rhs107;
        }
      } else if (_source14.is_DC5) {
        let _768___mcc_h5 = (_source14).cf12;
        let _769_cf12 = _768___mcc_h5;
        if ((_this).f26) {
          let _770_v73;
          _770_v73 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,_720_v39);
          let _771_v74;
          _771_v74 = _dafny.Seq.of(_770_v73);
          (globalState).f15 = (_668_v0).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_771_v74, _771_v74)).length));
          let _772_v75;
          _772_v75 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_769_cf12));
          let _773_v76;
          let _nw101 = new _module.C3();
          _nw101.__ctor(_772_v75, (p2)[_module.__default.safeIndex(_769_cf12, new BigNumber((p2).length))], (_this).f27);
          _773_v76 = _nw101;
          let _774_v77;
          _774_v77 = _dafny.Seq.of(_769_cf12, _668_v0);
          let _775_v78;
          _775_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-971)), ((_776_v0) => function (_777_i7) {
            return new BigNumber((_dafny.Seq.of(_776_v0, _776_v0, (_dafny.ZERO).minus(_776_v0))).length);
          })(_668_v0))).length),(_774_v77)[_module.__default.safeIndex((_773_v76).fm12(new BigNumber((_728_v47).length), globalState), new BigNumber((_774_v77).length))]);
          let _778_v79;
          _778_v79 = _dafny.Map.Empty.slice().updateUnsafe(_773_v76,new BigNumber((_775_v78).length));
          let _779_v80;
          _779_v80 = _dafny.Map.Empty.slice().updateUnsafe(_778_v79,_720_v39);
          let _780_v81;
          _780_v81 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,_775_v78);
          let _781_v82;
          let _nw102 = Array((new BigNumber(19)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(p2, _dafny.Seq.update(p2, _module.__default.safeIndex(_769_cf12, new BigNumber((p2).length)), (_this).f26));
          _nw102[(_dafny.ONE).toNumber()] = p2;
          _nw102[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(!(p0));
          _nw102[(new BigNumber(3)).toNumber()] = p2;
          _nw102[(new BigNumber(4)).toNumber()] = _module.__default.fm25(true, (_this).f26, _720_v39, globalState);
          _nw102[(new BigNumber(5)).toNumber()] = p2;
          _nw102[(new BigNumber(6)).toNumber()] = p2;
          _nw102[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p2, p2);
          _nw102[(new BigNumber(8)).toNumber()] = _module.__default.fm25((_this).f27, (_this).f26, (((_779_v80).contains(_778_v79)) ? ((_779_v80).get(_778_v79)) : (_720_v39)), globalState);
          _nw102[(new BigNumber(9)).toNumber()] = p2;
          _nw102[(new BigNumber(10)).toNumber()] = p2;
          _nw102[(new BigNumber(11)).toNumber()] = p2;
          _nw102[(new BigNumber(12)).toNumber()] = _dafny.Seq.update((((_this).f26) ? (p2) : (_dafny.Seq.of(false))), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_773_v76).fm12(_769_cf12, globalState),(((_780_v81).contains(new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm15(globalState))).cardinality()))) ? ((_780_v81).get(new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm15(globalState))).cardinality()))) : (_dafny.Map.Empty.slice().updateUnsafe(_769_cf12,_668_v0))))).length), new BigNumber(((((_this).f26) ? (p2) : (_dafny.Seq.of(false)))).length)), (_this).f27);
          _nw102[(new BigNumber(13)).toNumber()] = p2;
          _nw102[(new BigNumber(14)).toNumber()] = p2;
          _nw102[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(p2, _module.__default.safeIndex(_769_cf12, new BigNumber((p2).length)), (_this).f27);
          _nw102[(new BigNumber(16)).toNumber()] = p2;
          _nw102[(new BigNumber(17)).toNumber()] = p2;
          _nw102[(new BigNumber(18)).toNumber()] = _module.__default.fm25((_this).f26, (_this).f27, _720_v39, globalState);
          _781_v82 = _nw102;
          let _index95 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_781_v82).length));
          (_781_v82)[_index95] = p2;
          let _index96 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_781_v82).length));
          (_781_v82)[_index96] = p2;
          let _782_v83;
          let _nw103 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _782_v83 = _nw103;
          let _783_v84;
          _783_v84 = _module.D3.create_DC10(_782_v83);
          let _784_v85;
          _784_v85 = _dafny.Seq.of((_783_v84).dtor_cf21);
          let _785_v86;
          _785_v86 = _module.D2.create_DC9(new BigNumber(121), _668_v0, false);
          let _786_v87;
          let _nw104 = Array((new BigNumber(27)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _782_v83;
          _nw104[(_dafny.ONE).toNumber()] = (_784_v85)[_module.__default.safeIndex(_668_v0, new BigNumber((_784_v85).length))];
          _nw104[(new BigNumber(2)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(3)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(4)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(5)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(6)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(7)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(8)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(9)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(10)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(11)).toNumber()] = (((_785_v86).dtor_cf20) ? (_782_v83) : (_782_v83));
          _nw104[(new BigNumber(12)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(13)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(14)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(15)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(16)).toNumber()] = (((_this).f26) ? (_782_v83) : (_782_v83));
          _nw104[(new BigNumber(17)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(18)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(19)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(20)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(21)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(22)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(23)).toNumber()] = (_783_v84).dtor_cf21;
          _nw104[(new BigNumber(24)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(25)).toNumber()] = _782_v83;
          _nw104[(new BigNumber(26)).toNumber()] = _782_v83;
          _786_v87 = _nw104;
          _786_v87 = _786_v87;
          let _787_v88;
          _787_v88 = _dafny.MultiSet.fromElements((_this).f27, (_this).f27);
          let _788_v89;
          let _nw105 = Array((new BigNumber(25)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = !(!(!(p0))) || (!(p0));
          _nw105[(_dafny.ONE).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(2)).toNumber()] = (_this).f26;
          _nw105[(new BigNumber(3)).toNumber()] = p0;
          _nw105[(new BigNumber(4)).toNumber()] = !(p0);
          _nw105[(new BigNumber(5)).toNumber()] = p0;
          _nw105[(new BigNumber(6)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(7)).toNumber()] = (_this).f26;
          _nw105[(new BigNumber(8)).toNumber()] = p0;
          _nw105[(new BigNumber(9)).toNumber()] = false;
          _nw105[(new BigNumber(10)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(11)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(12)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(13)).toNumber()] = !((_787_v88).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).fm2(p0, _668_v0, globalState))));
          _nw105[(new BigNumber(14)).toNumber()] = p0;
          _nw105[(new BigNumber(15)).toNumber()] = (_this).f26;
          _nw105[(new BigNumber(16)).toNumber()] = p0;
          _nw105[(new BigNumber(17)).toNumber()] = p0;
          _nw105[(new BigNumber(18)).toNumber()] = !((_this).f26);
          _nw105[(new BigNumber(19)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(20)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(21)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(22)).toNumber()] = (p0) && (false);
          _nw105[(new BigNumber(23)).toNumber()] = (_this).f27;
          _nw105[(new BigNumber(24)).toNumber()] = (new BigNumber(((_781_v82)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_781_v82).length))]).length)).isLessThan(_668_v0);
          _788_v89 = _nw105;
          let _index97 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_788_v89).length));
          (_788_v89)[_index97] = (_this).fm2(p0, (_dafny.ZERO).minus(_769_cf12), globalState);
          let _index98 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_788_v89).length));
          (_788_v89)[_index98] = p0;
          let _789_v90;
          _789_v90 = _dafny.Seq.of(_dafny.Seq.of((_this).fm2(p0, _668_v0, globalState), (_788_v89)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_788_v89).length))]), p2, (_781_v82)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_781_v82).length))]);
          let _790_v91;
          let _nw106 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _790_v91 = _nw106;
          let _index99 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_790_v91).length));
          (_790_v91)[_index99] = _770_v73;
          let _index100 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_788_v89).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_790_v91).length));
          let _rhs108 = _789_v90;
          let _rhs109 = true;
          let _rhs110 = (((_775_v78).contains(_769_cf12)) ? ((_775_v78).get(_769_cf12)) : (new BigNumber(953)));
          let _rhs111 = _770_v73;
          let _lhs101 = _788_v89;
          let _lhs102 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_788_v89).length));
          let _lhs103 = globalState;
          let _lhs104 = _790_v91;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_790_v91).length));
          _789_v90 = _rhs108;
          _lhs101[_lhs102] = _rhs109;
          _lhs103.f12 = _rhs110;
          _lhs104[_lhs105] = _rhs111;
        } else {
          let _791_v92;
          _791_v92 = _dafny.Map.Empty.slice().updateUnsafe(true,_668_v0);
          let _792_v93;
          _792_v93 = _dafny.Seq.of(_791_v92);
          let _793_v94;
          let _nw107 = Array((new BigNumber(26)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = (_this).f27;
          _nw107[(_dafny.ONE).toNumber()] = p0;
          _nw107[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_728_v47, _module.__default.safeIndex(_769_cf12, new BigNumber((_728_v47).length)), new _dafny.CodePoint('e'.codePointAt(0))), _dafny.Seq.UnicodeFromString("urnk"));
          _nw107[(new BigNumber(3)).toNumber()] = !(p0) || ((_this).fm2((_this).f27, _668_v0, globalState));
          _nw107[(new BigNumber(4)).toNumber()] = (_this).f27;
          _nw107[(new BigNumber(5)).toNumber()] = (_this).fm2(p0, _769_cf12, globalState);
          _nw107[(new BigNumber(6)).toNumber()] = (_this).f27;
          _nw107[(new BigNumber(7)).toNumber()] = !((_this).f27);
          _nw107[(new BigNumber(8)).toNumber()] = !((_module.D5.create_DC14(_668_v0, _769_cf12, (_this).f26)).dtor_cf27);
          _nw107[(new BigNumber(9)).toNumber()] = p0;
          _nw107[(new BigNumber(10)).toNumber()] = (new BigNumber((p2).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_769_cf12));
          _nw107[(new BigNumber(11)).toNumber()] = p0;
          _nw107[(new BigNumber(12)).toNumber()] = (_this).f26;
          _nw107[(new BigNumber(13)).toNumber()] = (_this).f27;
          _nw107[(new BigNumber(14)).toNumber()] = p0;
          _nw107[(new BigNumber(15)).toNumber()] = p0;
          _nw107[(new BigNumber(16)).toNumber()] = (_this).f26;
          _nw107[(new BigNumber(17)).toNumber()] = (_this).f27;
          _nw107[(new BigNumber(18)).toNumber()] = (p2)[_module.__default.safeIndex(_668_v0, new BigNumber((p2).length))];
          _nw107[(new BigNumber(19)).toNumber()] = (_this).f26;
          _nw107[(new BigNumber(20)).toNumber()] = p0;
          _nw107[(new BigNumber(21)).toNumber()] = (_this).f26;
          _nw107[(new BigNumber(22)).toNumber()] = true;
          _nw107[(new BigNumber(23)).toNumber()] = p0;
          _nw107[(new BigNumber(24)).toNumber()] = (_this).f27;
          _nw107[(new BigNumber(25)).toNumber()] = !_dafny.areEqual(_792_v93, _792_v93);
          _793_v94 = _nw107;
          let _index102 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length));
          (_793_v94)[_index102] = (p2)[_module.__default.safeIndex(_668_v0, new BigNumber((p2).length))];
          let _index103 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length));
          (_793_v94)[_index103] = (_this).f27;
          let _794_v95;
          let _nw108 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _794_v95 = _nw108;
          let _795_v96;
          _795_v96 = _dafny.Map.Empty.slice().updateUnsafe(_794_v95,(_793_v94)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length))]);
          (globalState).f22 = (new BigNumber(((_795_v96).Merge(_dafny.Map.Empty.slice().updateUnsafe(_794_v95,p0))).length)).isLessThan(_769_cf12);
          let _796_v97;
          let _nw109 = new _module.C1();
          _nw109.__ctor(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_728_v47, _module.__default.safeIndex(_668_v0, new BigNumber((_728_v47).length)), _720_v39)).length), _668_v0), (_this).fm1(globalState), (_793_v94)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length))], (_this).f27);
          _796_v97 = _nw109;
          let _797_v98;
          _797_v98 = _module.D3.create_DC11(_720_v39);
          let _798_v99;
          _798_v99 = _dafny.Seq.of(new BigNumber((_721_v40).length), (_796_v97).f36);
          let _799_v100;
          _799_v100 = _dafny.Map.Empty.slice().updateUnsafe(_793_v94,_dafny.Seq.IsPrefixOf(_module.__default.fm23((_797_v98).dtor_cf22, globalState), _798_v99));
          _799_v100 = (_799_v100).update(_793_v94, !(false));
          let _800_v101;
          let _nw110 = new _module.C1();
          _nw110.__ctor(new BigNumber((_dafny.Seq.of(false, (_this).f26)).length), _796_v97.f37, (_793_v94)[_module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length))], p0);
          _800_v101 = _nw110;
          let _801_v102;
          _801_v102 = _module.D7.create_DC18(_800_v101);
          let _index104 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_794_v95).length));
          (_794_v95)[_index104] = _668_v0;
          let _index105 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_794_v95).length));
          let _index106 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length));
          let _rhs112 = _801_v102;
          let _rhs113 = _796_v97.f37;
          let _rhs114 = (_this).f26;
          let _rhs115 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_668_v0), _798_v99);
          let _rhs116 = _800_v101;
          let _lhs106 = _794_v95;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_794_v95).length));
          let _lhs108 = _793_v94;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_793_v94).length));
          let _lhs110 = globalState;
          _801_v102 = _rhs112;
          _lhs106[_lhs107] = _rhs113;
          _lhs108[_lhs109] = _rhs114;
          _lhs110.f2 = _rhs115;
          _800_v101 = _rhs116;
        }
        if (!(!(true) || (p0))) {
          (globalState).f24 = !_dafny.Seq.contains(p2, (_this).f27);
          (globalState).f22 = (_this).f26;
          let _rhs117 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_668_v0));
          let _rhs118 = new _dafny.CodePoint('y'.codePointAt(0));
          let _rhs119 = new BigNumber(687);
          let _lhs111 = globalState;
          let _lhs112 = globalState;
          _lhs111.f16 = _rhs117;
          _720_v39 = _rhs118;
          _lhs112.f16 = _rhs119;
          (globalState).f15 = ((_this).f26) || (p0);
          let _802_v103;
          let _init22 = ((_803_cf12, _804_v0) => function (_805_i8) {
            return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_803_cf12,_804_v0));
          })(_769_cf12, _668_v0);
          let _nw111 = Array((new BigNumber(29)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw111.length); _i0_22++) {
            _nw111[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _802_v103 = _nw111;
          let _index107 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_802_v103).length));
          (_802_v103)[_index107] = _module.__default.fm26((_this).fm2((_this).f26, _769_cf12, globalState), p0, p0, globalState);
          let _806_v104;
          _806_v104 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,_769_cf12);
          let _807_v105;
          _807_v105 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_769_cf12,_668_v0), (_806_v104).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-261),_769_cf12)), _806_v104);
          let _index108 = _module.__default.safeIndex(new BigNumber(369), new BigNumber((_802_v103).length));
          (_802_v103)[_index108] = _807_v105;
        } else {
          let _808_v106;
          _808_v106 = _dafny.MultiSet.fromElements(!((_this).fm2(p0, (_this).fm1(globalState), globalState)));
          let _rhs120 = !(!(!((_this).f27))) || ((_668_v0).isEqualTo(_769_cf12));
          let _rhs121 = _808_v106;
          let _lhs113 = globalState;
          _lhs113.f2 = _rhs120;
          _808_v106 = _rhs121;
          let _809_v107;
          let _init23 = ((_810_p2) => function (_811_i9) {
            return _810_p2;
          })(p2);
          let _nw112 = Array((new BigNumber(4)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw112.length); _i0_23++) {
            _nw112[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _809_v107 = _nw112;
          let _index109 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_809_v107).length));
          (_809_v107)[_index109] = p2;
          let _index110 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_809_v107).length));
          let _rhs122 = _769_cf12;
          let _rhs123 = !(_769_cf12).isEqualTo(_668_v0);
          let _rhs124 = _dafny.Seq.Concat(((true) ? (p2) : (p2)), p2);
          let _lhs114 = globalState;
          let _lhs115 = _809_v107;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_809_v107).length));
          _668_v0 = _rhs122;
          _lhs114.f24 = _rhs123;
          _lhs115[_lhs116] = _rhs124;
          let _812_v108;
          _812_v108 = _module.D2.create_DC8();
          let _813_v109;
          _813_v109 = _dafny.Set.fromElements(_812_v108);
          _813_v109 = _813_v109;
          (globalState).f15 = (_this).f26;
          let _814_v110;
          let _init24 = ((_815_p0) => function (_816_i10) {
            return _815_p0;
          })(p0);
          let _nw113 = Array((new BigNumber(17)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw113.length); _i0_24++) {
            _nw113[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _814_v110 = _nw113;
          let _817_v111;
          _817_v111 = _dafny.Map.Empty.slice().updateUnsafe((_809_v107)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_809_v107).length))],_814_v110);
          let _818_v112;
          _818_v112 = _module.D9.create_DC21(_817_v111);
          let _819_v113;
          _819_v113 = _dafny.Seq.of(_818_v112);
          (globalState).f24 = _dafny.Seq.contains(_819_v113, _818_v112);
        }
        let _820_v115;
        let _init25 = ((_821_cf12, _822_v0) => function (_823_i11) {
          return (function () {
            let _coll52 = new _dafny.Set();
            for (const _compr_52 of (_dafny.Set.fromElements(_822_v0)).Elements) {
              let _824_v114 = _compr_52;
              if ((_dafny.Set.fromElements(_822_v0)).contains(_824_v114)) {
                _coll52.add((_824_v114).plus((_dafny.ZERO).minus(new BigNumber(-311))));
              }
            }
            return _coll52;
          }()).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_822_v0,_dafny.Map.Empty.slice().updateUnsafe(_822_v0,_821_cf12))).length)));
        })(_769_cf12, _668_v0);
        let _nw114 = Array((new BigNumber(18)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw114.length); _i0_25++) {
          _nw114[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _820_v115 = _nw114;
        let _825_v116;
        _825_v116 = _module.D1.create_DC3((_module.D1.create_DC3(_dafny.Seq.of(p0))).dtor_cf6);
        let _826_v117;
        _826_v117 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm18(_721_v40, globalState)).length));
        let _827_v118;
        _827_v118 = _dafny.Seq.of(new BigNumber(336));
        let _pat_let_tv4 = p0;
        let _pat_let_tv5 = _826_v117;
        let _rhs125 = ((_this).f27) && ((_this).f26);
        let _rhs126 = _820_v115;
        let _rhs127 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((function (_pat_let10_0) {
          return function (_828_dt__update__tmp_h1) {
            return function (_pat_let11_0) {
              return function (_829_dt__update_hcf15_h0) {
                return function (_pat_let12_0) {
                  return function (_830_dt__update_hcf16_h0) {
                    return _module.D2.create_DC7((_828_dt__update__tmp_h1).dtor_cf14, _829_dt__update_hcf15_h0, _830_dt__update_hcf16_h0, (_828_dt__update__tmp_h1).dtor_cf17);
                  }(_pat_let12_0);
                }(new BigNumber((_pat_let_tv5).cardinality()));
              }(_pat_let11_0);
            }(_pat_let_tv4);
          }(_pat_let10_0);
        }(_725_v44)).dtor_cf17, (_827_v118)[_module.__default.safeIndex(new BigNumber(-276), new BigNumber((_827_v118).length))]));
        let _rhs128 = _825_v116;
        let _lhs117 = globalState;
        let _lhs118 = globalState;
        _lhs117.f15 = _rhs125;
        _820_v115 = _rhs126;
        _lhs118.f12 = _rhs127;
        _825_v116 = _rhs128;
        let _831_v119;
        _831_v119 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
        _831_v119 = (_831_v119).update(p0, p0);
      } else {
        let _832___mcc_h6 = (_source14).cf6;
        let _833_cf6 = _832___mcc_h6;
        (globalState).f19 = _720_v39;
        let _834_v120;
        _834_v120 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("q"),_668_v0);
        (globalState).f12 = new BigNumber((_834_v120).length);
        let _835_v121;
        let _836_v122;
        let _837_v123;
        let _838_v124;
        let _out25;
        let _out26;
        let _out27;
        let _out28;
        let _outcollector8 = _module.__default.m0(_720_v39, globalState);
        _out25 = _outcollector8[0];
        _out26 = _outcollector8[1];
        _out27 = _outcollector8[2];
        _out28 = _outcollector8[3];
        _835_v121 = _out25;
        _836_v122 = _out26;
        _837_v123 = _out27;
        _838_v124 = _out28;
        if (((!((_this).f26)) ? (((p0) ? (false) : ((_this).fm2((_this).f26, _668_v0, globalState)))) : (p0))) {
          let _839_v125;
          _839_v125 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm1(globalState));
          (globalState).f15 = ((_668_v0).multipliedBy((((_839_v125).contains(_837_v123)) ? ((_839_v125).get(_837_v123)) : (new BigNumber(869))))).isLessThanOrEqualTo(new BigNumber(954));
          let _840_v126;
          _840_v126 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,_668_v0);
          _840_v126 = (_840_v126).update(_668_v0, new BigNumber((_836_v122).length));
          _837_v123 = _837_v123;
          let _841_v127;
          let _init26 = function (_842_i12) {
            return (_this).f26;
          };
          let _nw115 = Array((new BigNumber(6)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw115.length); _i0_26++) {
            _nw115[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _841_v127 = _nw115;
          let _843_v128;
          _843_v128 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(243),_841_v127);
          let _844_v129;
          let _nw116 = Array((new BigNumber(23)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _841_v127;
          _nw116[(_dafny.ONE).toNumber()] = (((_843_v128).contains(_668_v0)) ? ((_843_v128).get(_668_v0)) : (_841_v127));
          _nw116[(new BigNumber(2)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(3)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(4)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(5)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(6)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(7)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(8)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(9)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(10)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(11)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(12)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(13)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(14)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(15)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(16)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(17)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(18)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(19)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(20)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(21)).toNumber()] = _841_v127;
          _nw116[(new BigNumber(22)).toNumber()] = _841_v127;
          _844_v129 = _nw116;
          (globalState).f5 = _844_v129;
          let _index111 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_841_v127).length));
          (_841_v127)[_index111] = p0;
          let _index112 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_841_v127).length));
          let _rhs129 = _835_v121;
          let _rhs130 = !(p0);
          let _rhs131 = (_668_v0).multipliedBy((_dafny.ZERO).minus(_668_v0));
          let _rhs132 = _dafny.Seq.Concat(_836_v122, _dafny.Seq.Concat(p2, _836_v122));
          let _rhs133 = (_this).f26;
          let _lhs119 = globalState;
          let _lhs120 = _841_v127;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_841_v127).length));
          let _lhs122 = globalState;
          let _lhs123 = globalState;
          _lhs119.f19 = _rhs129;
          _lhs120[_lhs121] = _rhs130;
          _lhs122.f12 = _rhs131;
          _836_v122 = _rhs132;
          _lhs123.f15 = _rhs133;
        } else {
          let _845_v130;
          _845_v130 = _dafny.Set.fromElements(_668_v0, _668_v0);
          let _846_v131;
          _846_v131 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,new BigNumber((_845_v130).length));
          let _847_v132;
          let _nw117 = Array((new BigNumber(13)).toNumber());
          _nw117[(_dafny.ZERO).toNumber()] = _668_v0;
          _nw117[(_dafny.ONE).toNumber()] = _668_v0;
          _nw117[(new BigNumber(2)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(3)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(4)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(5)).toNumber()] = new BigNumber(250);
          _nw117[(new BigNumber(6)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(7)).toNumber()] = new BigNumber((_846_v131).length);
          _nw117[(new BigNumber(8)).toNumber()] = new BigNumber(492);
          _nw117[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_668_v0);
          _nw117[(new BigNumber(10)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(11)).toNumber()] = _668_v0;
          _nw117[(new BigNumber(12)).toNumber()] = _668_v0;
          _847_v132 = _nw117;
          let _index113 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_847_v132).length));
          (_847_v132)[_index113] = _668_v0;
          let _index114 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_847_v132).length));
          (_847_v132)[_index114] = (_dafny.ZERO).minus((_668_v0).plus(_668_v0));
          (globalState).f24 = false;
          let _848_v133;
          _848_v133 = _module.D7.create_DC19(_668_v0, _module.__default.fm21((_847_v132)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_847_v132).length))], globalState), true, _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_668_v0));
          (globalState).f22 = !((((_this).f26) ? ((new BigNumber((_728_v47).length)).isLessThan(new BigNumber((_728_v47).length))) : ((_848_v133).dtor_cf38)));
          (globalState).f2 = (false) || ((_this).f27);
          let _849_v134;
          _849_v134 = _dafny.Map.Empty.slice().updateUnsafe(_728_v47,(_this).f26);
          let _850_v135;
          _850_v135 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber((_721_v40).length));
          _849_v134 = (_module.__default.fm27(new BigNumber((_850_v135).length), _668_v0, globalState)).Merge((_849_v134).update(_dafny.Seq.UnicodeFromString("bd"), _837_v123));
        }
      }
      let _851_v136;
      _851_v136 = _dafny.Map.Empty.slice().updateUnsafe(_668_v0,_dafny.Seq.UnicodeFromString("ga"));
      let _852_v137;
      _852_v137 = _module.D0.create_DC1((_this).f26, (_this).f26, new BigNumber((_851_v136).length), p0);
      let _853_i13;
      _853_i13 = _dafny.ZERO;
      L5: {
        while ((_module.__default.fm28(globalState)).equals(_dafny.Set.fromElements(_852_v137))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_853_i13)) {
              break L5;
            }
            _853_i13 = (_853_i13).plus(_dafny.ONE);
            (globalState).f12 = _668_v0;
            (globalState).f12 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_668_v0, _668_v0), new BigNumber(((((_this).f27) ? (p2) : (p2))).length));
            let _854_v138;
            let _nw118 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
            _854_v138 = _nw118;
            let _855_v139;
            _855_v139 = _dafny.Seq.of(_854_v138);
            let _rhs134 = _module.__default.safeDivisionInt(_668_v0, new BigNumber(-344));
            let _rhs135 = ((_this).f27) && (_dafny.areEqual(_855_v139, _855_v139));
            let _lhs124 = globalState;
            let _lhs125 = globalState;
            _lhs124.f12 = _rhs134;
            _lhs125.f15 = _rhs135;
            let _856_v140;
            _856_v140 = _module.D5.create_DC14(new BigNumber(803), _668_v0, (_this).f27);
            let _857_v141;
            _857_v141 = _dafny.Seq.of((_856_v140).dtor_cf26, (_this).fm1(globalState));
            _668_v0 = (_857_v141)[_module.__default.safeIndex(new BigNumber(291), new BigNumber((_857_v141).length))];
          }
        }
      }
      let _858_v142;
      let _nw119 = Array((new BigNumber(14)).toNumber()).fill(false);
      _858_v142 = _nw119;
      let _index115 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_858_v142).length));
      (_858_v142)[_index115] = (_this).f27;
      let _index116 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_858_v142).length));
      (_858_v142)[_index116] = (_668_v0).isLessThanOrEqualTo(_668_v0);
      let _index117 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_858_v142).length));
      let _rhs136 = _728_v47;
      let _rhs137 = (_668_v0).minus(_module.__default.safeDivisionInt(new BigNumber(-402), new BigNumber((_721_v40).length)));
      let _rhs138 = p0;
      let _rhs139 = !((_this).f26);
      let _lhs126 = globalState;
      let _lhs127 = globalState;
      let _lhs128 = _858_v142;
      let _lhs129 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_858_v142).length));
      let _lhs130 = globalState;
      _lhs126.f6 = _rhs136;
      _lhs127.f16 = _rhs137;
      _lhs128[_lhs129] = _rhs138;
      _lhs130.f24 = _rhs139;
      return;
    }
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _859_v0;
      let _nw120 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
      _859_v0 = _nw120;
      let _860_v1;
      _860_v1 = new BigNumber(-964);
      let _861_v2;
      _861_v2 = _dafny.Map.Empty.slice().updateUnsafe(_860_v1,_860_v1);
      let _862_v3;
      _862_v3 = _dafny.Map.Empty.slice().updateUnsafe(_860_v1,_861_v2);
      let _index118 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_859_v0).length));
      (_859_v0)[_index118] = _862_v3;
      let _863_v4;
      _863_v4 = _dafny.Seq.of(_860_v1, _860_v1);
      let _864_v5;
      _864_v5 = _dafny.MultiSet.fromElements((_this).fm2(p0, new BigNumber((_863_v4).length), globalState), (_this).f26);
      let _865_v6;
      _865_v6 = _dafny.Seq.of((_864_v5).IsSubsetOf(_864_v5));
      let _866_v7;
      let _nw121 = Array((new BigNumber(29)).toNumber()).fill([]);
      _866_v7 = _nw121;
      let _index119 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_866_v7).length));
      (_866_v7)[_index119] = p1;
      let _index120 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_859_v0).length));
      let _index121 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_866_v7).length));
      let _rhs140 = _862_v3;
      let _rhs141 = _865_v6;
      let _rhs142 = p1;
      let _lhs131 = _859_v0;
      let _lhs132 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_859_v0).length));
      let _lhs133 = _866_v7;
      let _lhs134 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_866_v7).length));
      _lhs131[_lhs132] = _rhs140;
      _865_v6 = _rhs141;
      _lhs133[_lhs134] = _rhs142;
      _860_v1 = _860_v1;
      (globalState).f2 = p0;
      r0 = (_this).f26;
      let _867_v8;
      let _nw122 = Array((new BigNumber(28)).toNumber()).fill([]);
      _867_v8 = _nw122;
      let _868_v9;
      let _nw123 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _868_v9 = _nw123;
      let _869_v10;
      _869_v10 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_860_v1),_868_v9);
      let _index122 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_867_v8).length));
      (_867_v8)[_index122] = (((_869_v10).contains(_860_v1)) ? ((_869_v10).get(_860_v1)) : (_868_v9));
      let _870_v11;
      let _nw124 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _870_v11 = _nw124;
      let _index123 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((p1).length));
      (p1)[_index123] = (_this).f27;
      let _index124 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_867_v8).length));
      let _index125 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((p1).length));
      let _index126 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_866_v7).length));
      let _rhs143 = _868_v9;
      let _rhs144 = _870_v11;
      let _rhs145 = p0;
      let _rhs146 = p1;
      let _lhs135 = _867_v8;
      let _lhs136 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_867_v8).length));
      let _lhs137 = p1;
      let _lhs138 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((p1).length));
      let _lhs139 = _866_v7;
      let _lhs140 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_866_v7).length));
      _lhs135[_lhs136] = _rhs143;
      _870_v11 = _rhs144;
      _lhs137[_lhs138] = _rhs145;
      _lhs139[_lhs140] = _rhs146;
      let _index127 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((p1).length));
      (p1)[_index127] = !((_this).f27);
      let _871_v12;
      _871_v12 = new _dafny.CodePoint('e'.codePointAt(0));
      let _872_v13;
      _872_v13 = _dafny.Map.Empty.slice().updateUnsafe(_860_v1,_871_v12);
      let _873_v14;
      _873_v14 = _dafny.Set.fromElements(_872_v13);
      r0 = !(((_873_v14).Difference(function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of (_873_v14).Elements) {
          let _874_v15 = _compr_53;
          if ((_873_v14).contains(_874_v15)) {
            _coll53.add(_874_v15);
          }
        }
        return _coll53;
      }())).IsSubsetOf(_873_v14));
      r1 = !(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(530)), ((_875_v1) => function (_876_i0) {
        return (_dafny.ZERO).minus(_875_v1);
      })(_860_v1)), _863_v4));
      let _877_v16;
      _877_v16 = _module.D9.create_DC22((_this).f27);
      r2 = (((_877_v16).dtor_cf42) ? ((_dafny.ZERO).minus(_860_v1)) : (new BigNumber((_module.__default.fm15(globalState)).length)));
      let _878_v17;
      _878_v17 = _module.D2.create_DC8();
      let _pat_let_tv6 = p1;
      let _pat_let_tv7 = p1;
      let _pat_let_tv8 = p1;
      let _pat_let_tv9 = p1;
      let _pat_let_tv10 = p2;
      let _pat_let_tv11 = _865_v6;
      let _pat_let_tv12 = _865_v6;
      let _pat_let_tv13 = p2;
      let _pat_let_tv14 = _865_v6;
      let _pat_let_tv15 = _865_v6;
      let _pat_let_tv16 = _860_v1;
      let _pat_let_tv17 = _860_v1;
      let _pat_let_tv18 = _860_v1;
      r3 = function (_source15) {
        if (_source15.is_DC7) {
          let _879___mcc_h0 = (_source15).cf14;
          let _880___mcc_h1 = (_source15).cf15;
          let _881___mcc_h2 = (_source15).cf16;
          let _882___mcc_h3 = (_source15).cf17;
          let _883_cf17 = _882___mcc_h3;
          let _884_cf16 = _881___mcc_h2;
          let _885_cf15 = _880___mcc_h1;
          let _886_cf14 = _879___mcc_h0;
          return (_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(427), new BigNumber((_pat_let_tv6).length))];
        } else if (_source15.is_DC8) {
          return (_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(427), new BigNumber((_pat_let_tv8).length))];
        } else if (_source15.is_DC9) {
          let _887___mcc_h4 = (_source15).cf18;
          let _888___mcc_h5 = (_source15).cf19;
          let _889___mcc_h6 = (_source15).cf20;
          let _890_cf20 = _889___mcc_h6;
          let _891_cf19 = _888___mcc_h5;
          let _892_cf18 = _887___mcc_h4;
          if ((_pat_let_tv10).contains((_pat_let_tv11)[_module.__default.safeIndex(_892_cf18, new BigNumber((_pat_let_tv12).length))])) {
            return (_pat_let_tv13).get((_pat_let_tv14)[_module.__default.safeIndex(_892_cf18, new BigNumber((_pat_let_tv15).length))]);
          } else {
            return (_module.D5.create_DC14(new BigNumber(752), new BigNumber(463), _890_cf20)).dtor_cf27;
          }
        } else {
          let _893___mcc_h7 = (_source15).cf13;
          let _894_cf13 = _893___mcc_h7;
          return (_dafny.Set.fromElements(_pat_let_tv16, _pat_let_tv17)).equals(function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of _dafny.IntegerRange(new BigNumber(-111), new BigNumber(-734))) {
              let _895_v18 = _compr_54;
              if (((new BigNumber(-111)).isLessThanOrEqualTo(_895_v18)) && ((_895_v18).isLessThan(new BigNumber(-734)))) {
                _coll54.add((_895_v18).plus(_pat_let_tv18));
              }
            }
            return _coll54;
          }());
        }
      }(_878_v17);
      return [r0, r1, r2, r3];
    }
    m10(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _896_v0;
      _896_v0 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),!(p0));
      let _897_v1;
      _897_v1 = _dafny.Seq.of(p0, (_this).fm2((_this).f27, new BigNumber(771), globalState), p0);
      let _hi10 = new BigNumber((_dafny.Seq.Concat(_897_v1, _897_v1)).length);
      for (let _898_i0 = new BigNumber((_896_v0).length); _898_i0.isLessThan(_hi10); _898_i0 = _898_i0.plus(_dafny.ONE)) {
        let _899_v2;
        let _init27 = ((_900_v1) => function (_901_i1) {
          return _dafny.Seq.of(_900_v1, _900_v1, _900_v1);
        })(_897_v1);
        let _nw125 = Array((new BigNumber(19)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw125.length); _i0_27++) {
          _nw125[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _899_v2 = _nw125;
        let _index128 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_899_v2).length));
        (_899_v2)[_index128] = _dafny.Seq.Concat(_dafny.Seq.of(_897_v1), _dafny.Seq.of(_897_v1));
        let _902_v3;
        _902_v3 = _dafny.Seq.of(_897_v1, _897_v1);
        let _index129 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_899_v2).length));
        (_899_v2)[_index129] = _dafny.Seq.Concat(_902_v3, _dafny.Seq.of(_dafny.Seq.of(p0), _897_v1));
        let _903_v4;
        _903_v4 = _dafny.Seq.UnicodeFromString("usmgaxt");
        let _904_v5;
        _904_v5 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), ((_905_p1) => function (_906_i2) {
          return _905_p1;
        })(p1))).length), new BigNumber(123), p1, _898_i0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_907_i3) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        })).length));
        let _908_v6;
        _908_v6 = _dafny.Set.fromElements(p1, new BigNumber((_904_v5).cardinality()));
        let _909_v7;
        _909_v7 = _dafny.Map.Empty.slice().updateUnsafe(_908_v6,(_this).f27);
        let _910_v8;
        _910_v8 = _dafny.Seq.of(new BigNumber(-252), new BigNumber((_908_v6).length));
        let _911_v9;
        _911_v9 = _dafny.MultiSet.fromElements((_this).f26);
        let _912_v10;
        let _nw126 = Array((new BigNumber(20)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm29(false, _898_i0, globalState)).length);
        _nw126[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm22(new BigNumber((_903_v4).length), (_897_v1)[_module.__default.safeIndex(p1, new BigNumber((_897_v1).length))], globalState)).length), _898_i0);
        _nw126[(new BigNumber(2)).toNumber()] = new BigNumber((_908_v6).length);
        _nw126[(new BigNumber(3)).toNumber()] = new BigNumber(((_908_v6).Difference(_908_v6)).length);
        _nw126[(new BigNumber(4)).toNumber()] = _898_i0;
        _nw126[(new BigNumber(5)).toNumber()] = (p1).minus(p1);
        _nw126[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_909_v7).Merge(_909_v7)).length));
        _nw126[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(p1, _898_i0);
        _nw126[(new BigNumber(8)).toNumber()] = (_898_i0).plus(p1);
        _nw126[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm30(_898_i0, p0, globalState)).length), new BigNumber((_897_v1).length));
        _nw126[(new BigNumber(10)).toNumber()] = _898_i0;
        _nw126[(new BigNumber(11)).toNumber()] = (_this).fm1(globalState);
        _nw126[(new BigNumber(12)).toNumber()] = _898_i0;
        _nw126[(new BigNumber(13)).toNumber()] = _898_i0;
        _nw126[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((((_this).f27) ? (p1) : (_898_i0)));
        _nw126[(new BigNumber(15)).toNumber()] = p1;
        _nw126[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(_898_i0, (_910_v8)[_module.__default.safeIndex(_898_i0, new BigNumber((_910_v8).length))]);
        _nw126[(new BigNumber(17)).toNumber()] = p1;
        _nw126[(new BigNumber(18)).toNumber()] = new BigNumber((_903_v4).length);
        _nw126[(new BigNumber(19)).toNumber()] = new BigNumber((_911_v9).cardinality());
        _912_v10 = _nw126;
        let _index130 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_912_v10).length));
        (_912_v10)[_index130] = p1;
        let _index131 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_912_v10).length));
        (_912_v10)[_index131] = (p1).minus(_898_i0);
        let _913_v11;
        let _nw127 = new _module.C1();
        _nw127.__ctor((_this).fm1(globalState), new BigNumber((_897_v1).length), (_this).f26, p0);
        _913_v11 = _nw127;
        (_913_v11).f37 = new BigNumber(((_911_v9).Difference(_911_v9)).cardinality());
      }
      let _914_v12;
      _914_v12 = _dafny.MultiSet.fromElements((_this).f26, (_this).f27, (_this).fm2((_this).f27, p1, globalState));
      (globalState).f22 = (_897_v1)[_module.__default.safeIndex(new BigNumber((_914_v12).cardinality()), new BigNumber((_897_v1).length))];
      let _915_v13;
      let _init28 = ((_916_p1, _917_p0) => function (_918_i4) {
        return _dafny.Seq.of((_dafny.ZERO).minus((_module.D5.create_DC14(_916_p1, new BigNumber((_dafny.Seq.of(new BigNumber(-342), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_919_i5) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})).length))).length), _917_p0)).dtor_cf26));
      })(p1, p0);
      let _nw128 = Array((new BigNumber(13)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw128.length); _i0_28++) {
        _nw128[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _915_v13 = _nw128;
      _915_v13 = _915_v13;
      let _920_v14;
      _920_v14 = _dafny.Seq.UnicodeFromString("vgkdg");
      let _921_v16;
      _921_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(602),_dafny.Set.fromElements((_this).f26));
      if (!(((p0) ? (function () {
        let _coll55 = new _dafny.Map();
        for (const _compr_55 of _dafny.IntegerRange(new BigNumber(-648), new BigNumber(519))) {
          let _922_v15 = _compr_55;
          if (((new BigNumber(-648)).isLessThanOrEqualTo(_922_v15)) && ((_922_v15).isLessThan(new BigNumber(519)))) {
            _coll55.push([(_922_v15).multipliedBy(p1),_dafny.Set.fromElements((_this).f26)]);
          }
        }
        return _coll55;
      }()) : (_921_v16))).contains(new BigNumber((_920_v14).length))) {
        let _923_v17;
        let _nw129 = Array((new BigNumber(11)).toNumber()).fill(false);
        _923_v17 = _nw129;
        let _index132 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_923_v17).length));
        (_923_v17)[_index132] = (_this).f26;
        let _index133 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_923_v17).length));
        (_923_v17)[_index133] = !(true);
        let _924_v18;
        _924_v18 = _dafny.Seq.of(_897_v1);
        _924_v18 = _924_v18;
        let _925_v19;
        let _nw130 = new _module.C0();
        _nw130.__ctor((_this).f27);
        _925_v19 = _nw130;
        let _926_v20;
        _926_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ettqukfnq"),_925_v19);
        let _927_v21;
        let _nw131 = Array((new BigNumber(16)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = _925_v19;
        _nw131[(_dafny.ONE).toNumber()] = _925_v19;
        _nw131[(new BigNumber(2)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(3)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(4)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(5)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(6)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(7)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(8)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(9)).toNumber()] = (((_923_v17)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_923_v17).length))]) ? (_925_v19) : (_925_v19));
        _nw131[(new BigNumber(10)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(11)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(12)).toNumber()] = (((_926_v20).contains(_920_v14)) ? ((_926_v20).get(_920_v14)) : (_925_v19));
        _nw131[(new BigNumber(13)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(14)).toNumber()] = _925_v19;
        _nw131[(new BigNumber(15)).toNumber()] = _925_v19;
        _927_v21 = _nw131;
        let _index134 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_927_v21).length));
        (_927_v21)[_index134] = _925_v19;
        let _index135 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_927_v21).length));
        (_927_v21)[_index135] = _925_v19;
        let _928_v22;
        let _nw132 = Array((new BigNumber(4)).toNumber());
        _928_v22 = _nw132;
        let _929_v23;
        _929_v23 = _dafny.Seq.of(p1, p1);
        let _930_v24;
        _930_v24 = _dafny.MultiSet.fromElements(_929_v23, _929_v23, _929_v23);
        let _931_v25;
        let _nw133 = new _module.C3();
        _nw133.__ctor(_930_v24, (_923_v17)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_923_v17).length))], (_this).f26);
        _931_v25 = _nw133;
        let _index136 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_928_v22).length));
        (_928_v22)[_index136] = _931_v25;
        let _index137 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_928_v22).length));
        (_928_v22)[_index137] = _931_v25;
        let _932_v26;
        _932_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        _932_v26 = (_932_v26).update(!_dafny.areEqual(_dafny.Seq.of(_897_v1), _924_v18), (_this).f26);
      } else {
        let _933_v27;
        _933_v27 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f27),p1);
        let _934_v28;
        _934_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_933_v27).length),p1);
        let _935_v29;
        _935_v29 = _dafny.Seq.of(_915_v13, _915_v13);
        let _rhs147 = p1;
        let _rhs148 = new BigNumber((_dafny.Seq.Concat(_897_v1, _dafny.Seq.Concat(_897_v1, _897_v1))).length);
        let _rhs149 = (_this).f26;
        let _rhs150 = ((false) ? (p1) : ((((_934_v28).contains(p1)) ? ((_934_v28).get(p1)) : (p1))));
        let _rhs151 = (_935_v29)[_module.__default.safeIndex((_this).fm1(globalState), new BigNumber((_935_v29).length))];
        let _lhs141 = globalState;
        r2 = _rhs147;
        r2 = _rhs148;
        _lhs141.f22 = _rhs149;
        r2 = _rhs150;
        _915_v13 = _rhs151;
        let _936_v30;
        let _nw134 = Array((new BigNumber(5)).toNumber()).fill(false);
        _936_v30 = _nw134;
        let _937_v31;
        _937_v31 = new _dafny.CodePoint('s'.codePointAt(0));
        let _938_v32;
        _938_v32 = _dafny.Map.Empty.slice().updateUnsafe(_936_v30,_937_v31);
        let _939_v33;
        let _940_v34;
        let _941_v35;
        let _942_v36;
        let _out29;
        let _out30;
        let _out31;
        let _out32;
        let _outcollector9 = _module.__default.m0((((_938_v32).contains(_936_v30)) ? ((_938_v32).get(_936_v30)) : (_937_v31)), globalState);
        _out29 = _outcollector9[0];
        _out30 = _outcollector9[1];
        _out31 = _outcollector9[2];
        _out32 = _outcollector9[3];
        _939_v33 = _out29;
        _940_v34 = _out30;
        _941_v35 = _out31;
        _942_v36 = _out32;
        let _943_v37;
        _943_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,_940_v34);
        let _944_v38;
        let _nw135 = new _module.C1();
        _nw135.__ctor(p1, new BigNumber(((((_943_v37).contains(p1)) ? ((_943_v37).get(p1)) : (_940_v34))).length), (_this).f26, (_this).f26);
        _944_v38 = _nw135;
        let _945_v39;
        let _946_v40;
        let _947_v41;
        let _948_v42;
        let _out33;
        let _out34;
        let _out35;
        let _out36;
        let _outcollector10 = _module.__default.m0(new _dafny.CodePoint('l'.codePointAt(0)), globalState);
        _out33 = _outcollector10[0];
        _out34 = _outcollector10[1];
        _out35 = _outcollector10[2];
        _out36 = _outcollector10[3];
        _945_v39 = _out33;
        _946_v40 = _out34;
        _947_v41 = _out35;
        _948_v42 = _out36;
        if ((_module.D5.create_DC14(_944_v38.f37, new BigNumber((_914_v12).cardinality()), _947_v41)).dtor_cf27) {
          let _949_v43;
          _949_v43 = _dafny.Seq.of((_dafny.ZERO).minus(_944_v38.f37), _944_v38.f37);
          let _950_v44;
          _950_v44 = _dafny.Seq.of(_949_v43, _949_v43);
          let _951_v45;
          _951_v45 = _module.D0.create_DC0((_this).f26);
          let _952_v46;
          _952_v46 = _dafny.Set.fromElements((_944_v38).fm4(_951_v45, new BigNumber((_dafny.Seq.UnicodeFromString("uc")).length), new BigNumber((_946_v40).length), globalState));
          let _953_v47;
          let _nw136 = new _module.C2();
          _nw136.__ctor(_950_v44, ((_941_v35) ? (p0) : (_947_v41)), (_952_v46).IsProperSubsetOf(_952_v46));
          _953_v47 = _nw136;
          let _nw137 = new _module.C2();
          _nw137.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-120)), ((_954_v43) => function (_955_i6) {
            return _954_v43;
          })(_949_v43)), _dafny.Seq.IsPrefixOf(_920_v14, _920_v14), (_953_v47).f26);
          _953_v47 = _nw137;
          _936_v30 = _936_v30;
          (globalState).f2 = (_953_v47).f26;
          let _956_v48;
          _956_v48 = _dafny.Set.fromElements((_this).fm2(!((_953_v47).f27), _944_v38.f37, globalState));
          let _957_v49;
          _957_v49 = _dafny.Map.Empty.slice().updateUnsafe(_939_v33,_944_v38.f37);
          let _958_v50;
          _958_v50 = _dafny.Map.Empty.slice().updateUnsafe((_944_v38).f36,_944_v38);
          let _959_v51;
          _959_v51 = _dafny.Map.Empty.slice().updateUnsafe((_944_v38).f36,_920_v14);
          let _960_v52;
          let _nw138 = Array((new BigNumber(17)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = new BigNumber((_956_v48).length);
          _nw138[(_dafny.ONE).toNumber()] = new BigNumber((_957_v49).length);
          _nw138[(new BigNumber(2)).toNumber()] = (_944_v38).f36;
          _nw138[(new BigNumber(3)).toNumber()] = new BigNumber((_958_v50).length);
          _nw138[(new BigNumber(4)).toNumber()] = new BigNumber(((((_959_v51).contains((_944_v38).f36)) ? ((_959_v51).get((_944_v38).f36)) : (_920_v14))).length);
          _nw138[(new BigNumber(5)).toNumber()] = _944_v38.f37;
          _nw138[(new BigNumber(6)).toNumber()] = (_944_v38).f36;
          _nw138[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("umscmdh")).length);
          _nw138[(new BigNumber(8)).toNumber()] = (_944_v38).f36;
          _nw138[(new BigNumber(9)).toNumber()] = (_944_v38).f36;
          _nw138[(new BigNumber(10)).toNumber()] = _944_v38.f37;
          _nw138[(new BigNumber(11)).toNumber()] = new BigNumber((_module.__default.fm31(p1, globalState)).length);
          _nw138[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_944_v38).f36);
          _nw138[(new BigNumber(13)).toNumber()] = (_944_v38).f36;
          _nw138[(new BigNumber(14)).toNumber()] = _944_v38.f37;
          _nw138[(new BigNumber(15)).toNumber()] = _944_v38.f37;
          _nw138[(new BigNumber(16)).toNumber()] = (_944_v38).f36;
          _960_v52 = _nw138;
          let _961_v53;
          _961_v53 = _dafny.Map.Empty.slice().updateUnsafe(_960_v52,_945_v39);
          _961_v53 = (_961_v53).update(_960_v52, _939_v33);
          r0 = _module.__default.safeDivisionInt(((_947_v41) ? ((_944_v38).f36) : (new BigNumber(((_953_v47).fm3(_920_v14, _897_v1, (_944_v38).f36, globalState)).length))), (_944_v38).f36);
        } else {
          let _962_v54;
          let _nw139 = Array((new BigNumber(7)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt((_944_v38).f36, _944_v38.f37);
          _nw139[(_dafny.ONE).toNumber()] = (_944_v38).f36;
          _nw139[(new BigNumber(2)).toNumber()] = _944_v38.f37;
          _nw139[(new BigNumber(3)).toNumber()] = p1;
          _nw139[(new BigNumber(4)).toNumber()] = _944_v38.f37;
          _nw139[(new BigNumber(5)).toNumber()] = (_944_v38).f36;
          _nw139[(new BigNumber(6)).toNumber()] = (_944_v38.f37).plus((_944_v38).f36);
          _962_v54 = _nw139;
          let _963_v55;
          let _nw140 = Array((new BigNumber(4)).toNumber());
          _nw140[(_dafny.ZERO).toNumber()] = _962_v54;
          _nw140[(_dafny.ONE).toNumber()] = _962_v54;
          _nw140[(new BigNumber(2)).toNumber()] = _962_v54;
          _nw140[(new BigNumber(3)).toNumber()] = _962_v54;
          _963_v55 = _nw140;
          let _rhs152 = _941_v35;
          let _rhs153 = _962_v54;
          let _rhs154 = new BigNumber(89);
          let _rhs155 = _963_v55;
          let _lhs142 = globalState;
          let _lhs143 = _944_v38;
          _lhs142.f2 = _rhs152;
          _962_v54 = _rhs153;
          _lhs143.f37 = _rhs154;
          _963_v55 = _rhs155;
          let _964_v56;
          _964_v56 = _module.D0.create_DC0((_this).f26);
          (globalState).f12 = (_944_v38).fm4(_964_v56, (_dafny.ZERO).minus(((_944_v38).f36).multipliedBy(_944_v38.f37)), _module.__default.safeDivisionInt(_944_v38.f37, (_944_v38).f36), globalState);
          (globalState).f22 = !(true);
          let _965_v57;
          let _nw141 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _965_v57 = _nw141;
          let _966_v58;
          let _nw142 = Array((new BigNumber(5)).toNumber());
          _nw142[(_dafny.ZERO).toNumber()] = _944_v38;
          _nw142[(_dafny.ONE).toNumber()] = _944_v38;
          _nw142[(new BigNumber(2)).toNumber()] = _944_v38;
          _nw142[(new BigNumber(3)).toNumber()] = _944_v38;
          _nw142[(new BigNumber(4)).toNumber()] = _944_v38;
          _966_v58 = _nw142;
          let _index138 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_966_v58).length));
          (_966_v58)[_index138] = _944_v38;
          let _index139 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_966_v58).length));
          let _rhs156 = _965_v57;
          let _rhs157 = (_this).f27;
          let _rhs158 = (new BigNumber(826)).multipliedBy((_944_v38).f36);
          let _rhs159 = (_944_v38).f36;
          let _rhs160 = _944_v38;
          let _lhs144 = globalState;
          let _lhs145 = globalState;
          let _lhs146 = globalState;
          let _lhs147 = _966_v58;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(832), new BigNumber((_966_v58).length));
          _965_v57 = _rhs156;
          _lhs144.f22 = _rhs157;
          _lhs145.f16 = _rhs158;
          _lhs146.f16 = _rhs159;
          _lhs147[_lhs148] = _rhs160;
          (globalState).f22 = (_this).f26;
        }
      }
      let _967_v59;
      let _init29 = function (_968_i7) {
        return (_this).f27;
      };
      let _nw143 = Array((new BigNumber(7)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw143.length); _i0_29++) {
        _nw143[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _967_v59 = _nw143;
      let _969_v60;
      _969_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f27);
      let _970_v61;
      let _971_v62;
      let _972_v63;
      let _973_v64;
      let _out37;
      let _out38;
      let _out39;
      let _out40;
      let _outcollector11 = (_this).m9((_this).f27, _967_v59, (((_this).f26) ? (_969_v60) : (_969_v60)), globalState);
      _out37 = _outcollector11[0];
      _out38 = _outcollector11[1];
      _out39 = _outcollector11[2];
      _out40 = _outcollector11[3];
      _970_v61 = _out37;
      _971_v62 = _out38;
      _972_v63 = _out39;
      _973_v64 = _out40;
      let _974_v65;
      _974_v65 = _dafny.Seq.of(_897_v1);
      _974_v65 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_897_v1), _974_v65), _974_v65);
      r0 = p1;
      r1 = _972_v63;
      r2 = (_dafny.ZERO).minus(p1);
      r3 = _module.__default.fm24(_920_v14, (_this).f26, p1, globalState);
      return [r0, r1, r2, r3];
    }
    get f33() {
      let _this = this;
      return _this._f33;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f26 = false;
      this._f27 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    __ctor(f26, f27) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    fm1(globalState) {
      let _this = this;
      return new BigNumber(-356);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      let _source16 = (((_this).f27) ? (_module.D0.create_DC2(_module.D0.create_DC1((_this).f27, (_this).f27, new BigNumber(756), !(true)))) : (_module.D0.create_DC2(_module.D0.create_DC0((_this).f26))));
      if (_source16.is_DC1) {
        let _975___mcc_h0 = (_source16).cf1;
        let _976___mcc_h1 = (_source16).cf2;
        let _977___mcc_h2 = (_source16).cf3;
        let _978___mcc_h3 = (_source16).cf4;
        let _979_cf4 = _978___mcc_h3;
        let _980_cf3 = _977___mcc_h2;
        let _981_cf2 = _976___mcc_h1;
        let _982_cf1 = _975___mcc_h0;
        return (_this).f27;
      } else if (_source16.is_DC0) {
        let _983___mcc_h4 = (_source16).cf0;
        let _984_cf0 = _983___mcc_h4;
        return (new BigNumber((_dafny.Seq.UnicodeFromString("wrscocei")).length)).isEqualTo(new BigNumber(468));
      } else {
        let _985___mcc_h5 = (_source16).cf5;
        let _986_cf5 = _985___mcc_h5;
        return (_this).f26;
      }
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _987_v0;
      let _nw144 = new _module.C0();
      _nw144.__ctor((_this).f26);
      _987_v0 = _nw144;
      let _988_v2;
      _988_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      let _989_v3;
      _989_v3 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll56 = new _dafny.Map();
        for (const _compr_56 of (_988_v2).Keys.Elements) {
          let _990_v1 = _compr_56;
          if ((_988_v2).contains(_990_v1)) {
            _coll56.push([_990_v1,new BigNumber((_dafny.Seq.UnicodeFromString("fmtrj")).length)]);
          }
        }
        return _coll56;
      }(),p0);
      let _991_v4;
      _991_v4 = new BigNumber(-494);
      let _992_v5;
      _992_v5 = _dafny.Map.Empty.slice().updateUnsafe(p2,_991_v4);
      let _993_v7;
      _993_v7 = _module.D2.create_DC6(function () {
  let _coll57 = new _dafny.Map();
  for (const _compr_57 of _dafny.IntegerRange(new BigNumber(828), new BigNumber(982))) {
    let _994_v6 = _compr_57;
    if (((new BigNumber(828)).isLessThanOrEqualTo(_994_v6)) && ((_994_v6).isLessThan(new BigNumber(982)))) {
      _coll57.push([(_994_v6).minus(_991_v4),_991_v4]);
    }
  }
  return _coll57;
}());
      _989_v3 = (_989_v3).update(_992_v5, !((_991_v4).isLessThan(new BigNumber(((_993_v7).dtor_cf13).length))));
      (globalState).f6 = _dafny.Seq.UnicodeFromString("nbjowp");
      let _995_v8;
      let _nw145 = new _module.C0();
      _nw145.__ctor(!(p0) || ((_this).f27));
      _995_v8 = _nw145;
      let _996_v9;
      _996_v9 = _module.D2.create_DC9(new BigNumber(530), new BigNumber((p2).length), (_this).f26);
      let _source17 = _996_v9;
      if (_source17.is_DC7) {
        let _997___mcc_h0 = (_source17).cf14;
        let _998___mcc_h1 = (_source17).cf15;
        let _999___mcc_h2 = (_source17).cf16;
        let _1000___mcc_h3 = (_source17).cf17;
        let _1001_cf17 = _1000___mcc_h3;
        let _1002_cf16 = _999___mcc_h2;
        let _1003_cf15 = _998___mcc_h1;
        let _1004_cf14 = _997___mcc_h0;
        let _1005_v10;
        let _nw146 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1005_v10 = _nw146;
        let _1006_v11;
        _1006_v11 = _dafny.Seq.UnicodeFromString("m");
        let _index140 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1005_v10).length));
        (_1005_v10)[_index140] = _1006_v11;
        let _index141 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1005_v10).length));
        (_1005_v10)[_index141] = _dafny.Seq.Concat(_1006_v11, _1006_v11);
        (globalState).f22 = (_this).fm2(_1004_cf14, (_dafny.ZERO).minus(_991_v4), globalState);
        let _1007_v12;
        let _nw147 = Array((new BigNumber(20)).toNumber()).fill(false);
        _1007_v12 = _nw147;
        let _index142 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1007_v12).length));
        (_1007_v12)[_index142] = _dafny.Seq.IsPrefixOf((_1005_v10)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_1005_v10).length))], (_1005_v10)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_1005_v10).length))]);
        let _index143 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_1007_v12).length));
        (_1007_v12)[_index143] = false;
        (globalState).f2 = (_this).f27;
      } else if (_source17.is_DC8) {
        (globalState).f15 = (_987_v0).f32;
        let _1008_v13;
        _1008_v13 = _module.D0.create_DC1((_987_v0).f32, (_995_v8).f32, _991_v4, (_995_v8).f32);
        let _1009_v14;
        _1009_v14 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,(_1008_v13).dtor_cf3);
        let _1010_v15;
        _1010_v15 = _dafny.MultiSet.fromElements(_module.__default.fm8(new BigNumber(40), _991_v4, (_987_v0).f32, _991_v4, globalState), _1009_v14, _1009_v14, _1009_v14, _1009_v14);
        let _1011_v16;
        _1011_v16 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1012_v17;
        _1012_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1011_v16);
        let _1013_v18;
        _1013_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1012_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1011_v16))).length),_dafny.MultiSet.fromElements(_1009_v14));
        _1010_v15 = (((_1013_v18).contains(_991_v4)) ? ((_1013_v18).get(_991_v4)) : ((_1010_v15).Union(_1010_v15)));
        let _1014_v19;
        _1014_v19 = _dafny.Seq.UnicodeFromString("sksvlyg");
        (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.update(_1014_v19, _module.__default.safeIndex(_991_v4, new BigNumber((_1014_v19).length)), _1011_v16), _1014_v19);
        let _1015_v20;
        _1015_v20 = _dafny.MultiSet.fromElements((_987_v0).f32, (_987_v0).f32);
        _1015_v20 = _module.__default.fm9(!((_995_v8).f32), _991_v4, (_987_v0).f32, globalState);
      } else if (_source17.is_DC9) {
        let _1016___mcc_h4 = (_source17).cf18;
        let _1017___mcc_h5 = (_source17).cf19;
        let _1018___mcc_h6 = (_source17).cf20;
        let _1019_cf20 = _1018___mcc_h6;
        let _1020_cf19 = _1017___mcc_h5;
        let _1021_cf18 = _1016___mcc_h4;
        let _1022_v21;
        _1022_v21 = _dafny.Seq.UnicodeFromString("bfsobu");
        let _1023_v22;
        _1023_v22 = _dafny.MultiSet.fromElements(true);
        let _1024_v23;
        _1024_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_1020_cf19);
        let _1025_v24;
        _1025_v24 = _dafny.MultiSet.fromElements(new BigNumber((_1023_v22).cardinality()), new BigNumber((_1024_v23).length));
        let _1026_v25;
        _1026_v25 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1027_v26;
        _1027_v26 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1022_v21, _module.__default.safeIndex((((_1025_v24).contains((_this).fm1(globalState))) ? ((_1025_v24).get((_this).fm1(globalState))) : (_1020_cf19)), new BigNumber((_1022_v21).length)), _1026_v25),_1026_v25);
        _1027_v26 = (_1027_v26).update(_dafny.Seq.UnicodeFromString("ja"), _1026_v25);
        (globalState).f16 = _991_v4;
        _1021_cf18 = (_this).fm1(globalState);
        let _1028_v27;
        _1028_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1021_cf18,(_987_v0).f32);
        let _1029_v28;
        _1029_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1028_v27).update(_1021_cf18, (_this).f26),_991_v4);
        _1021_cf18 = new BigNumber((_1029_v28).length);
      } else {
        let _1030___mcc_h7 = (_source17).cf13;
        let _1031_cf13 = _1030___mcc_h7;
        (globalState).f16 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(535)), function (_1032_i0) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })).length);
        let _1033_v29;
        _1033_v29 = _dafny.Seq.UnicodeFromString("cc");
        let _1034_v30;
        _1034_v30 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,_991_v4);
        let _1035_v31;
        _1035_v31 = _dafny.Seq.of(_991_v4, _991_v4, _991_v4);
        let _1036_v32;
        _1036_v32 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,(_995_v8).f32);
        let _1037_v33;
        _1037_v33 = _module.D1.create_DC5(_991_v4);
        let _1038_v34;
        let _nw148 = Array((new BigNumber(21)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = _991_v4;
        _nw148[(_dafny.ONE).toNumber()] = new BigNumber(41);
        _nw148[(new BigNumber(2)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(3)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(4)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(5)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(6)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(7)).toNumber()] = new BigNumber(887);
        _nw148[(new BigNumber(8)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(9)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(10)).toNumber()] = new BigNumber((_1033_v29).length);
        _nw148[(new BigNumber(11)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(12)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(13)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(14)).toNumber()] = (((_1034_v30).contains((_987_v0).f32)) ? ((_1034_v30).get((_987_v0).f32)) : (_991_v4));
        _nw148[(new BigNumber(15)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(16)).toNumber()] = (_1035_v31)[_module.__default.safeIndex(_991_v4, new BigNumber((_1035_v31).length))];
        _nw148[(new BigNumber(17)).toNumber()] = new BigNumber((_1033_v29).length);
        _nw148[(new BigNumber(18)).toNumber()] = _991_v4;
        _nw148[(new BigNumber(19)).toNumber()] = new BigNumber((_1036_v32).length);
        _nw148[(new BigNumber(20)).toNumber()] = (_1037_v33).dtor_cf12;
        _1038_v34 = _nw148;
        let _1039_v35;
        _1039_v35 = _module.D3.create_DC10(_1038_v34);
        let _1040_v36;
        let _nw149 = Array((new BigNumber(25)).toNumber());
        _nw149[(_dafny.ZERO).toNumber()] = _1038_v34;
        _nw149[(_dafny.ONE).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(2)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(3)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(4)).toNumber()] = (_1039_v35).dtor_cf21;
        _nw149[(new BigNumber(5)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(6)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(7)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(8)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(9)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(10)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(11)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(12)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(13)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(14)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(15)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(16)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(17)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(18)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(19)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(20)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(21)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(22)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(23)).toNumber()] = _1038_v34;
        _nw149[(new BigNumber(24)).toNumber()] = _1038_v34;
        _1040_v36 = _nw149;
        _1040_v36 = _1040_v36;
        let _1041_v37;
        _1041_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v29,_991_v4);
        let _1042_v38;
        _1042_v38 = _dafny.MultiSet.fromElements(p0, (_987_v0).fm6(new _dafny.CodePoint('j'.codePointAt(0)), (_995_v8).f32, globalState), (_987_v0).f32, (p2)[_module.__default.safeIndex(new BigNumber(((_1041_v37).update(_1033_v29, _991_v4)).length), new BigNumber((p2).length))], (_995_v8).f32);
        let _1043_v39;
        _1043_v39 = _dafny.MultiSet.fromElements((new BigNumber((_1042_v38).cardinality())).plus(_991_v4), _991_v4);
        let _1044_v40;
        let _nw150 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1044_v40 = _nw150;
        let _1045_v41;
        _1045_v41 = _module.D1.create_DC4(_1044_v40, _991_v4, _991_v4, _1033_v29, _1033_v29);
        let _1046_v42;
        _1046_v42 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_987_v0).f32);
        let _pat_let_tv19 = _1046_v42;
        let _1047_v43;
        _1047_v43 = _dafny.Map.Empty.slice().updateUnsafe((_987_v0).f32,(function (_pat_let13_0) {
          return function (_1048_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_1049_dt__update_hcf8_h0) {
                return _module.D1.create_DC4((_1048_dt__update__tmp_h0).dtor_cf7, _1049_dt__update_hcf8_h0, (_1048_dt__update__tmp_h0).dtor_cf9, (_1048_dt__update__tmp_h0).dtor_cf10, (_1048_dt__update__tmp_h0).dtor_cf11);
              }(_pat_let14_0);
            }(new BigNumber((_pat_let_tv19).length));
          }(_pat_let13_0);
        }(_1045_v41)).dtor_cf10);
        let _1050_v44;
        _1050_v44 = _module.D0.create_DC1((_987_v0).f32, (_this).f26, _991_v4, !((_this).f26));
        _1043_v39 = (_module.__default.fm10(_991_v4, new BigNumber((_1047_v43).length), _module.D0.create_DC2(_1050_v44), globalState)).Union(_1043_v39);
        if ((_this).f26) {
          (globalState).f16 = new BigNumber((((_dafny.MultiSet.fromElements((_this).f26)).Union(_1042_v38)).Difference(_1042_v38)).cardinality());
          (globalState).f16 = _991_v4;
          let _1051_v45;
          let _nw151 = Array((new BigNumber(3)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = p0;
          _nw151[(_dafny.ONE).toNumber()] = (_987_v0).f32;
          _nw151[(new BigNumber(2)).toNumber()] = (_987_v0).f32;
          _1051_v45 = _nw151;
          let _index144 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1051_v45).length));
          (_1051_v45)[_index144] = (_this).f26;
          let _1052_v46;
          _1052_v46 = _module.D0.create_DC0((_987_v0).f32);
          let _index145 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1051_v45).length));
          let _rhs161 = !(p0);
          let _rhs162 = (((_1052_v46).dtor_cf0) ? ((_987_v0).f32) : (!((_this).f27)));
          let _lhs149 = _1051_v45;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_1051_v45).length));
          let _lhs151 = globalState;
          _lhs149[_lhs150] = _rhs161;
          _lhs151.f24 = _rhs162;
          _1038_v34 = _1038_v34;
          (globalState).f12 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_991_v4), _991_v4);
        } else {
          let _1053_v47;
          let _nw152 = Array((new BigNumber(14)).toNumber()).fill(false);
          _1053_v47 = _nw152;
          let _index146 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_1053_v47).length));
          (_1053_v47)[_index146] = (_1043_v39).IsProperSubsetOf(_1043_v39);
          let _1054_v48;
          _1054_v48 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index147 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_1053_v47).length));
          (_1053_v47)[_index147] = (_995_v8).fm6(_1054_v48, (_995_v8).f32, globalState);
          let _1055_v49;
          _1055_v49 = _module.D2.create_DC8();
          let _rhs163 = _1055_v49;
          let _rhs164 = true;
          let _lhs152 = globalState;
          _1055_v49 = _rhs163;
          _lhs152.f24 = _rhs164;
          let _1056_v50;
          _1056_v50 = _module.D0.create_DC0(true);
          let _1057_v51;
          let _out41;
          _out41 = (_this).m8(_1056_v50, (_this).f27, (_dafny.MultiSet.fromElements((_987_v0).f32)).update((_987_v0).f32, _module.__default.abs(_991_v4)), _991_v4, globalState);
          _1057_v51 = _out41;
          let _index148 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1038_v34).length));
          (_1038_v34)[_index148] = _991_v4;
          let _index149 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1038_v34).length));
          (_1038_v34)[_index149] = ((((_1053_v47)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_1053_v47).length))]) ? (_991_v4) : (_991_v4))).minus(new BigNumber((_1033_v29).length));
          _1054_v48 = _1054_v48;
        }
      }
      let _1058_v52;
      _1058_v52 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1059_v53;
      _1059_v53 = _module.D1.create_DC3(_module.__default.fm11((_995_v8).f32, (_987_v0).f32, _1058_v52, globalState));
      let _pat_let_tv20 = p2;
      let _pat_let_tv21 = _991_v4;
      let _pat_let_tv22 = p2;
      let _pat_let_tv23 = _987_v0;
      let _source18 = function (_pat_let15_0) {
        return function (_1060_dt__update__tmp_h1) {
          return function (_pat_let16_0) {
            return function (_1061_dt__update_hcf6_h0) {
              return _module.D1.create_DC3(_1061_dt__update_hcf6_h0);
            }(_pat_let16_0);
          }(_dafny.Seq.update(_pat_let_tv20, _module.__default.safeIndex(_pat_let_tv21, new BigNumber((_pat_let_tv22).length)), (_pat_let_tv23).f32));
        }(_pat_let15_0);
      }(_1059_v53);
      if (_source18.is_DC4) {
        let _1062___mcc_h8 = (_source18).cf7;
        let _1063___mcc_h9 = (_source18).cf8;
        let _1064___mcc_h10 = (_source18).cf9;
        let _1065___mcc_h11 = (_source18).cf10;
        let _1066___mcc_h12 = (_source18).cf11;
        let _1067_cf11 = _1066___mcc_h12;
        let _1068_cf10 = _1065___mcc_h11;
        let _1069_cf9 = _1064___mcc_h10;
        let _1070_cf8 = _1063___mcc_h9;
        let _1071_cf7 = _1062___mcc_h8;
        let _1072_v54;
        let _nw153 = Array((new BigNumber(4)).toNumber()).fill([]);
        _1072_v54 = _nw153;
        let _1073_v55;
        _1073_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1072_v54,(_this).fm1(globalState));
        let _1074_v56;
        _1074_v56 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(true, (_987_v0).f32, (_this).f27));
        _1073_v55 = (_1073_v55).update(_1072_v54, (new BigNumber((_1074_v56).length)).plus(_1070_cf8));
        let _1075_v57;
        _1075_v57 = _dafny.Set.fromElements(!((_995_v8).f32), (_this).f27);
        let _1076_v58;
        _1076_v58 = _1075_v57;
        (globalState).f2 = (_1075_v57).IsProperSubsetOf((_1076_v58));
        _991_v4 = _1069_cf9;
        let _1077_v59;
        _1077_v59 = _module.D2.create_DC8();
        let _source19 = _1077_v59;
        if (_source19.is_DC7) {
          let _1078___mcc_h15 = (_source19).cf14;
          let _1079___mcc_h16 = (_source19).cf15;
          let _1080___mcc_h17 = (_source19).cf16;
          let _1081___mcc_h18 = (_source19).cf17;
          let _1082_cf17 = _1081___mcc_h18;
          let _1083_cf16 = _1080___mcc_h17;
          let _1084_cf15 = _1079___mcc_h16;
          let _1085_cf14 = _1078___mcc_h15;
          let _1086_v60;
          let _init30 = ((_1087_p2) => function (_1088_i1) {
            return _module.__default.safeModuloInt(_1088_i1, new BigNumber((_1087_p2).length));
          })(p2);
          let _nw154 = Array((new BigNumber(18)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw154.length); _i0_30++) {
            _nw154[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1086_v60 = _nw154;
          let _1089_v61;
          _1089_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1086_v60,(_this).fm1(globalState));
          _1089_v61 = (_1089_v61).update(_1086_v60, (_1082_cf17).minus(new BigNumber((_1068_cf10).length)));
          (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(45)), ((_1090_v52) => function (_1091_i2) {
            return _1090_v52;
          })(_1058_v52)), _1067_cf11);
          _995_v8 = _995_v8;
          (globalState).f6 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(830)), ((_1092_v52) => function (_1093_i3) {
            return _1092_v52;
          })(_1058_v52));
        } else if (_source19.is_DC8) {
          let _1094_v62;
          _1094_v62 = _dafny.Seq.of(new BigNumber(((_dafny.MultiSet.fromElements(_1058_v52)).update(_1058_v52, _module.__default.abs(new BigNumber(204)))).cardinality()));
          let _1095_v63;
          let _nw155 = Array((new BigNumber(23)).toNumber()).fill(false);
          _1095_v63 = _nw155;
          let _1096_v64;
          _1096_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1095_v63,_dafny.Seq.update(_1094_v62, _module.__default.safeIndex(_991_v4, new BigNumber((_1094_v62).length)), _1069_cf9));
          let _1097_v65;
          _1097_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1094_v62,(((_1096_v64).contains(_1095_v63)) ? ((_1096_v64).get(_1095_v63)) : (_1094_v62)));
          (globalState).f15 = !(!(_1097_v65).equals(_1097_v65));
          let _1098_v66;
          _1098_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1069_cf9,_1094_v62);
          _1098_v66 = (_1098_v66).update(_991_v4, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_1099_cf9) => function (_1100_i4) {
            return _1099_cf9;
          })(_1069_cf9)), _1094_v62), _module.__default.safeIndex(_991_v4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_1101_cf9) => function (_1102_i4) {
            return _1101_cf9;
          })(_1069_cf9)), _1094_v62)).length)), _991_v4));
          let _1103_v67;
          _1103_v67 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,_1070_cf8);
          _1103_v67 = (_1103_v67).update(new BigNumber(-944), _991_v4);
          (globalState).f2 = (_987_v0).f32;
        } else if (_source19.is_DC9) {
          let _1104___mcc_h19 = (_source19).cf18;
          let _1105___mcc_h20 = (_source19).cf19;
          let _1106___mcc_h21 = (_source19).cf20;
          let _1107_cf20 = _1106___mcc_h21;
          let _1108_cf19 = _1105___mcc_h20;
          let _1109_cf18 = _1104___mcc_h19;
          (globalState).f2 = !(_1107_cf20) || ((_1109_cf18).isLessThan(_1070_cf8));
          let _index150 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1071_cf7).length));
          (_1071_cf7)[_index150] = (((_995_v8).f32) ? (_1058_v52) : (_1058_v52));
          let _1110_v68;
          _1110_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1107_cf20,(_987_v0).f32);
          let _1111_v69;
          _1111_v69 = _dafny.MultiSet.fromElements(_991_v4, (_dafny.ZERO).minus((_1070_cf8).minus(_991_v4)), (_dafny.ZERO).minus(((false) ? (new BigNumber((_1110_v68).length)) : (_1109_cf18))), (_1069_cf9).multipliedBy(new BigNumber(194)));
          let _index151 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1071_cf7).length));
          let _rhs165 = _1058_v52;
          let _rhs166 = ((_1111_v69).Intersect(_1111_v69)).Union((_1111_v69).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_991_v4,p0)).length))));
          let _rhs167 = _1109_cf18;
          let _lhs153 = _1071_cf7;
          let _lhs154 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1071_cf7).length));
          let _lhs155 = globalState;
          _lhs153[_lhs154] = _rhs165;
          _1111_v69 = _rhs166;
          _lhs155.f12 = _rhs167;
          let _1112_v70;
          let _init31 = ((_1113_v4) => function (_1114_i5) {
            return _module.__default.safeDivisionInt(_1114_i5, _1113_v4);
          })(_991_v4);
          let _nw156 = Array((_dafny.ONE).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw156.length); _i0_31++) {
            _nw156[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1112_v70 = _nw156;
          let _1115_v71;
          _1115_v71 = _module.D3.create_DC10(_1112_v70);
          let _1116_v72;
          _1116_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1115_v71,(_this).fm1(globalState));
          let _1117_v73;
          let _nw157 = new _module.C4();
          _nw157.__ctor(_1116_v72, (_995_v8).f32, (_this).f27);
          _1117_v73 = _nw157;
          let _1118_v74;
          _1118_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1070_cf8,_1117_v73);
          let _1119_v75;
          let _nw158 = Array((new BigNumber(15)).toNumber());
          _nw158[(_dafny.ZERO).toNumber()] = _1109_cf18;
          _nw158[(_dafny.ONE).toNumber()] = (_this).fm1(globalState);
          _nw158[(new BigNumber(2)).toNumber()] = new BigNumber((_1075_v57).length);
          _nw158[(new BigNumber(3)).toNumber()] = _1108_cf19;
          _nw158[(new BigNumber(4)).toNumber()] = new BigNumber(507);
          _nw158[(new BigNumber(5)).toNumber()] = new BigNumber((_1068_cf10).length);
          _nw158[(new BigNumber(6)).toNumber()] = _1108_cf19;
          _nw158[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_1070_cf8);
          _nw158[(new BigNumber(8)).toNumber()] = _1069_cf9;
          _nw158[(new BigNumber(9)).toNumber()] = (_module.D0.create_DC1((_987_v0).f32, (_995_v8).f32, new BigNumber((_1068_cf10).length), (_995_v8).f32)).dtor_cf3;
          _nw158[(new BigNumber(10)).toNumber()] = _1108_cf19;
          _nw158[(new BigNumber(11)).toNumber()] = new BigNumber((_1118_v74).length);
          _nw158[(new BigNumber(12)).toNumber()] = _1108_cf19;
          _nw158[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_1070_cf8);
          _nw158[(new BigNumber(14)).toNumber()] = _1108_cf19;
          _1119_v75 = _nw158;
          let _1120_v76;
          _1120_v76 = _dafny.MultiSet.fromElements(_1119_v75, _1119_v75, _1112_v70);
          (globalState).f16 = new BigNumber((_1120_v76).cardinality());
          let _1121_v77;
          let _nw159 = new _module.C1();
          _nw159.__ctor((_1070_cf8).plus(_1069_cf9), (_1069_cf9).minus(_1069_cf9), (_991_v4).isLessThanOrEqualTo(_1070_cf8), (!((_995_v8).f32)) || (_1107_cf20));
          _1121_v77 = _nw159;
        } else {
          let _1122___mcc_h22 = (_source19).cf13;
          let _1123_cf13 = _1122___mcc_h22;
          let _1124_v78;
          let _nw160 = new _module.C1();
          _nw160.__ctor(new BigNumber((p2).length), _1069_cf9, p0, (_this).f27);
          _1124_v78 = _nw160;
          let _1125_v79;
          _1125_v79 = _dafny.Set.fromElements(_1124_v78, _1124_v78, _1124_v78);
          _1125_v79 = (_1125_v79).Difference(_1125_v79);
          (globalState).f22 = (true) && ((_995_v8).f32);
          (globalState).f12 = (_1124_v78).f36;
          (globalState).f6 = _dafny.Seq.Concat(_1067_cf11, (((_987_v0).f32) ? (_1068_cf10) : (_1067_cf11)));
        }
      } else if (_source18.is_DC5) {
        let _1126___mcc_h13 = (_source18).cf12;
        let _1127_cf12 = _1126___mcc_h13;
        let _1128_v80;
        _1128_v80 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,_dafny.Seq.of((_987_v0).f32, (_987_v0).f32));
        let _1129_v81;
        _1129_v81 = _module.D5.create_DC13(((true) ? (_module.__default.fm32((_this).fm1(globalState), _1127_cf12, globalState)) : ((_1128_v80).update(p0, p2))));
        let _source20 = _1129_v81;
        if (_source20.is_DC14) {
          let _1130___mcc_h23 = (_source20).cf25;
          let _1131___mcc_h24 = (_source20).cf26;
          let _1132___mcc_h25 = (_source20).cf27;
          let _1133_cf27 = _1132___mcc_h25;
          let _1134_cf26 = _1131___mcc_h24;
          let _1135_cf25 = _1130___mcc_h23;
          (globalState).f16 = (_dafny.ZERO).minus(_1134_cf26);
          let _rhs168 = _1134_cf26;
          let _rhs169 = new BigNumber(294);
          let _rhs170 = (_987_v0).f32;
          let _lhs156 = globalState;
          let _lhs157 = globalState;
          let _lhs158 = globalState;
          _lhs156.f16 = _rhs168;
          _lhs157.f12 = _rhs169;
          _lhs158.f24 = _rhs170;
          let _1136_v82;
          let _init32 = function (_1137_i6) {
            return _module.D2.create_DC8();
          };
          let _nw161 = Array((_dafny.ONE).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw161.length); _i0_32++) {
            _nw161[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1136_v82 = _nw161;
          let _index152 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1136_v82).length));
          (_1136_v82)[_index152] = _module.D2.create_DC8();
          let _1138_v83;
          _1138_v83 = _dafny.Seq.of(_1058_v52);
          let _1139_v84;
          _1139_v84 = _dafny.Set.fromElements(p0, false);
          let _1140_v85;
          _1140_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1138_v83,new BigNumber((_1139_v84).length));
          let _index153 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_1136_v82).length));
          (_1136_v82)[_index153] = _module.__default.fm33((_dafny.ZERO).minus((((_this).f26) ? (new BigNumber(743)) : (_1135_cf25))), new BigNumber(((_1140_v85).Merge((_1140_v85).update(_1138_v83, _1127_cf12))).length), globalState);
          (globalState).f16 = (_dafny.ZERO).minus(_1135_cf25);
        } else if (_source20.is_DC13) {
          let _1141___mcc_h26 = (_source20).cf24;
          let _1142_cf24 = _1141___mcc_h26;
          let _1143_v86;
          let _nw162 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
          _1143_v86 = _nw162;
          let _1144_v87;
          _1144_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-46),_1127_cf12);
          let _1145_v88;
          _1145_v88 = _dafny.Seq.of(_1144_v87);
          let _index154 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1143_v86).length));
          (_1143_v86)[_index154] = _1145_v88;
          let _index155 = _module.__default.safeIndex(new BigNumber(522), new BigNumber((_1143_v86).length));
          (_1143_v86)[_index155] = _1145_v88;
          let _1146_v89;
          let _nw163 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _1146_v89 = _nw163;
          _1146_v89 = _1146_v89;
          _1127_cf12 = _991_v4;
          let _1147_v90;
          _1147_v90 = _dafny.Seq.UnicodeFromString("priuoxrge");
          let _1148_v91;
          _1148_v91 = _dafny.MultiSet.fromElements((p1)[_module.__default.safeIndex(_991_v4, new BigNumber((p1).length))], _1147_v90);
          _1148_v91 = _1148_v91;
        } else {
          let _1149___mcc_h27 = (_source20).cf28;
          let _1150_cf28 = _1149___mcc_h27;
          let _1151_v92;
          _1151_v92 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_995_v8).f32));
          let _1152_v93;
          let _init33 = ((_1153_v0) => function (_1154_i7) {
            return (_1153_v0).f32;
          })(_987_v0);
          let _nw164 = Array((new BigNumber(13)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw164.length); _i0_33++) {
            _nw164[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1152_v93 = _nw164;
          let _1155_v94;
          _1155_v94 = _dafny.MultiSet.fromElements(_1152_v93, _1152_v93, _1152_v93);
          let _1156_v95;
          _1156_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_991_v4);
          let _1157_v96;
          _1157_v96 = _dafny.Seq.UnicodeFromString("dq");
          let _1158_v97;
          _1158_v97 = _dafny.Seq.of(new BigNumber((_1157_v96).length));
          let _1159_v98;
          _1159_v98 = _dafny.Set.fromElements((_987_v0).f32, (_this).f26, false, (new BigNumber(609)).isEqualTo(new BigNumber(((_1151_v92).update(p2, _module.__default.abs((((_1155_v94).contains(_1152_v93)) ? ((_1155_v94).get(_1152_v93)) : (_991_v4))))).cardinality())), (_987_v0).fm7(_1156_v95, _1158_v97, (_this).fm2((_this).f26, (_dafny.ZERO).minus(new BigNumber((_1157_v96).length)), globalState), globalState));
          let _rhs171 = _1159_v98;
          let _rhs172 = _996_v9;
          let _rhs173 = _1158_v97;
          _1159_v98 = _rhs171;
          _996_v9 = _rhs172;
          _1158_v97 = _rhs173;
          let _1160_v99;
          _1160_v99 = _dafny.MultiSet.fromElements(_1158_v97, _1158_v97, _1158_v97);
          let _1161_v100;
          let _nw165 = new _module.C3();
          _nw165.__ctor((_1160_v99).update(_1158_v97, _module.__default.abs(_991_v4)), (_991_v4).isLessThanOrEqualTo(_1127_cf12), ((_995_v8).f32) === (!((_995_v8).f32)));
          _1161_v100 = _nw165;
          _1158_v97 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1158_v97, _dafny.Seq.of(_1127_cf12, _1127_cf12)), _1158_v97);
          (globalState).f22 = (_995_v8).f32;
        }
        let _source21 = _module.D1.create_DC3(_dafny.Seq.of(true, (_this).f26, (_987_v0).f32));
        if (_source21.is_DC4) {
          let _1162___mcc_h28 = (_source21).cf7;
          let _1163___mcc_h29 = (_source21).cf8;
          let _1164___mcc_h30 = (_source21).cf9;
          let _1165___mcc_h31 = (_source21).cf10;
          let _1166___mcc_h32 = (_source21).cf11;
          let _1167_cf11 = _1166___mcc_h32;
          let _1168_cf10 = _1165___mcc_h31;
          let _1169_cf9 = _1164___mcc_h30;
          let _1170_cf8 = _1163___mcc_h29;
          let _1171_cf7 = _1162___mcc_h28;
          let _1172_v101;
          let _nw166 = Array((new BigNumber(11)).toNumber());
          _nw166[(_dafny.ZERO).toNumber()] = _1170_cf8;
          _nw166[(_dafny.ONE).toNumber()] = new BigNumber(-155);
          _nw166[(new BigNumber(2)).toNumber()] = _1169_cf9;
          _nw166[(new BigNumber(3)).toNumber()] = _1170_cf8;
          _nw166[(new BigNumber(4)).toNumber()] = new BigNumber(272);
          _nw166[(new BigNumber(5)).toNumber()] = new BigNumber(842);
          _nw166[(new BigNumber(6)).toNumber()] = _991_v4;
          _nw166[(new BigNumber(7)).toNumber()] = _991_v4;
          _nw166[(new BigNumber(8)).toNumber()] = _1169_cf9;
          _nw166[(new BigNumber(9)).toNumber()] = new BigNumber((_1167_cf11).length);
          _nw166[(new BigNumber(10)).toNumber()] = _991_v4;
          _1172_v101 = _nw166;
          let _1173_v102;
          _1173_v102 = _module.D3.create_DC10(_1172_v101);
          let _1174_v103;
          _1174_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v102,_991_v4);
          let _1175_v104;
          let _nw167 = new _module.C4();
          _nw167.__ctor(_1174_v103, (_987_v0).f32, (_995_v8).f32);
          _1175_v104 = _nw167;
          let _1176_v105;
          _1176_v105 = _dafny.Seq.of(_1175_v104);
          let _1177_v106;
          _1177_v106 = _dafny.Set.fromElements((_987_v0).f32);
          let _1178_v107;
          _1178_v107 = _dafny.Set.fromElements(new BigNumber((_1177_v106).length));
          let _1179_v108;
          _1179_v108 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), ((_1180_v4) => function (_1181_i8) {
            return _1180_v4;
          })(_991_v4)));
          let _1182_v109;
          let _nw168 = new _module.C3();
          _nw168.__ctor(_1179_v108, (_995_v8).f32, !(p0));
          _1182_v109 = _nw168;
          let _1183_v110;
          _1183_v110 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(new BigNumber((_1176_v105).length), _1170_cf8)).IsSubsetOf(_1178_v107),_1182_v109);
          _1183_v110 = _1183_v110;
          (globalState).f22 = !((new BigNumber((_dafny.Seq.Concat(_1167_cf11, _1167_cf11)).length)).isEqualTo((((_995_v8).f32) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(77)), ((_1184_v52) => function (_1185_i9) {
            return _1184_v52;
          })(_1058_v52))).length))) : (_991_v4))));
          let _1186_v111;
          let _init34 = ((_1187_p2, _1188_v0) => function (_1189_i10) {
            return _dafny.Seq.Concat(_1187_p2, _dafny.Seq.of((_1188_v0).f32, (_this).f27));
          })(p2, _987_v0);
          let _nw169 = Array((new BigNumber(17)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw169.length); _i0_34++) {
            _nw169[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1186_v111 = _nw169;
          let _index156 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1186_v111).length));
          (_1186_v111)[_index156] = _dafny.Seq.Concat(p2, p2);
          let _index157 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_1186_v111).length));
          (_1186_v111)[_index157] = p2;
          let _1190_v112;
          _1190_v112 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), ((_1191_v52) => function (_1192_i11) {
            return _1191_v52;
          })(_1058_v52)),_1170_cf8);
          _1190_v112 = (_1190_v112).update(_dafny.Seq.Concat(_1167_cf11, _1168_cf10), _991_v4);
        } else if (_source21.is_DC5) {
          let _1193___mcc_h33 = (_source21).cf12;
          let _1194_cf12 = _1193___mcc_h33;
          let _1195_v113;
          _1195_v113 = _dafny.Seq.of(_991_v4, _1194_cf12);
          let _1196_v114;
          _1196_v114 = _dafny.Seq.UnicodeFromString("nhtyej");
          let _1197_v115;
          _1197_v115 = _dafny.MultiSet.fromElements(_1196_v114, _1196_v114);
          let _1198_v116;
          _1198_v116 = _dafny.Seq.of(_993_v7, _993_v7, _993_v7);
          let _1199_v117;
          _1199_v117 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1195_v113, _module.__default.safeIndex(new BigNumber((_1197_v115).cardinality()), new BigNumber((_1195_v113).length)), _1194_cf12),_1198_v116);
          let _1200_v118;
          _1200_v118 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,_dafny.Map.Empty.slice().updateUnsafe(_1195_v113,_1198_v116));
          let _1201_v119;
          _1201_v119 = _module.D10.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_1195_v113,_1198_v116));
          let _1202_v120;
          let _nw170 = Array((new BigNumber(23)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _1199_v117;
          _nw170[(_dafny.ONE).toNumber()] = (((_1200_v118).contains(true)) ? ((_1200_v118).get(true)) : (_1199_v117));
          _nw170[(new BigNumber(2)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(3)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(4)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(5)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(6)).toNumber()] = (_1199_v117).update(_module.__default.fm15(globalState), _dafny.Seq.of(_993_v7));
          _nw170[(new BigNumber(7)).toNumber()] = (_1199_v117).Merge(_1199_v117);
          _nw170[(new BigNumber(8)).toNumber()] = (_module.__default.fm34(globalState)).update(_dafny.Seq.of(_991_v4), _1198_v116);
          _nw170[(new BigNumber(9)).toNumber()] = (_1199_v117).Merge(_1199_v117);
          _nw170[(new BigNumber(10)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1195_v113,_1198_v116);
          _nw170[(new BigNumber(12)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(13)).toNumber()] = (_1199_v117).update(_1195_v113, _dafny.Seq.Create(_module.__default.abs(new BigNumber(721)), ((_1203_v7) => function (_1204_i12) {
            return _1203_v7;
          })(_993_v7)));
          _nw170[(new BigNumber(14)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(15)).toNumber()] = (_1201_v119).dtor_cf43;
          _nw170[(new BigNumber(16)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(17)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_1205_cf12) => function (_1206_i13) {
            return _1205_cf12;
          })(_1194_cf12)),_1198_v116);
          _nw170[(new BigNumber(19)).toNumber()] = (_1199_v117).Merge(_1199_v117);
          _nw170[(new BigNumber(20)).toNumber()] = (_1199_v117).update(_dafny.Seq.of(_1194_cf12), _1198_v116);
          _nw170[(new BigNumber(21)).toNumber()] = _1199_v117;
          _nw170[(new BigNumber(22)).toNumber()] = (_1199_v117).Merge(_1199_v117);
          _1202_v120 = _nw170;
          let _index158 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1202_v120).length));
          (_1202_v120)[_index158] = _dafny.Map.Empty.slice().updateUnsafe(_1195_v113,_1198_v116);
          let _1207_v122;
          _1207_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1195_v113,_991_v4);
          let _index159 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1202_v120).length));
          (_1202_v120)[_index159] = ((_1199_v117).Merge((function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_1207_v122).Keys.Elements) {
              let _1208_v121 = _compr_58;
              if ((_1207_v122).contains(_1208_v121)) {
                _coll58.push([_1208_v121,_1198_v116]);
              }
            }
            return _coll58;
          }()).update(_dafny.Seq.of(_1194_cf12), _1198_v116))).Merge(_1199_v117);
          let _rhs174 = (new BigNumber((_dafny.Seq.UnicodeFromString("ertaynatq")).length)).minus(new BigNumber(893));
          let _rhs175 = _1194_cf12;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          _lhs159.f16 = _rhs174;
          _lhs160.f16 = _rhs175;
          (globalState).f15 = (_987_v0).f32;
          let _1209_v123;
          let _nw171 = Array((new BigNumber(12)).toNumber());
          _1209_v123 = _nw171;
          let _1210_v124;
          _1210_v124 = _dafny.MultiSet.fromElements(_1209_v123, _1209_v123);
          let _rhs176 = _1210_v124;
          let _rhs177 = (_996_v9).dtor_cf19;
          let _lhs161 = globalState;
          _1210_v124 = _rhs176;
          _lhs161.f12 = _rhs177;
        } else {
          let _1211___mcc_h34 = (_source21).cf6;
          let _1212_cf6 = _1211___mcc_h34;
          let _1213_v125;
          _1213_v125 = _dafny.Seq.UnicodeFromString("wsdmh");
          let _1214_v126;
          _1214_v126 = _dafny.Seq.of((_this).fm1(globalState));
          let _1215_v127;
          _1215_v127 = _dafny.Map.Empty.slice().updateUnsafe((_987_v0).f32,_991_v4);
          let _1216_v128;
          _1216_v128 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
          let _1217_v129;
          let _nw172 = Array((new BigNumber(15)).toNumber());
          _nw172[(_dafny.ZERO).toNumber()] = (_1214_v126)[_module.__default.safeIndex(_1127_cf12, new BigNumber((_1214_v126).length))];
          _nw172[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_1127_cf12);
          _nw172[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_991_v4);
          _nw172[(new BigNumber(3)).toNumber()] = (_this).fm1(globalState);
          _nw172[(new BigNumber(4)).toNumber()] = _1127_cf12;
          _nw172[(new BigNumber(5)).toNumber()] = new BigNumber((_1215_v127).length);
          _nw172[(new BigNumber(6)).toNumber()] = new BigNumber(698);
          _nw172[(new BigNumber(7)).toNumber()] = _991_v4;
          _nw172[(new BigNumber(8)).toNumber()] = new BigNumber((_1216_v128).length);
          _nw172[(new BigNumber(9)).toNumber()] = _991_v4;
          _nw172[(new BigNumber(10)).toNumber()] = _1127_cf12;
          _nw172[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_991_v4);
          _nw172[(new BigNumber(12)).toNumber()] = _1127_cf12;
          _nw172[(new BigNumber(13)).toNumber()] = _991_v4;
          _nw172[(new BigNumber(14)).toNumber()] = _1127_cf12;
          _1217_v129 = _nw172;
          let _1218_v130;
          _1218_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1217_v129,_1213_v125);
          let _1219_v131;
          let _nw173 = new _module.C1();
          _nw173.__ctor(new BigNumber((_1213_v125).length), (_dafny.ZERO).minus(new BigNumber((_1214_v126).length)), (_1218_v130).contains(_1217_v129), (_995_v8).f32);
          _1219_v131 = _nw173;
          let _1220_v132;
          _1220_v132 = _module.D0.create_DC0((_this).f26);
          (globalState).f22 = !((_987_v0).f32) || ((_this).fm2((_this).f26, (_1219_v131).fm4(_1220_v132, (_1219_v131).f36, _991_v4, globalState), globalState));
          let _1221_v133;
          _1221_v133 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1213_v125, _dafny.Seq.UnicodeFromString("uxenwl")),_1058_v52);
          _1221_v133 = (_1221_v133).update(_1213_v125, _1058_v52);
          _1215_v127 = (_1215_v127).update((_995_v8).fm6(_1058_v52, (_987_v0).f32, globalState), new BigNumber(487));
        }
        if ((_995_v8).f32) {
          let _1222_v134;
          _1222_v134 = _dafny.Seq.UnicodeFromString("ck");
          let _1223_v135;
          _1223_v135 = _dafny.MultiSet.fromElements(new BigNumber((_1222_v134).length));
          let _1224_v136;
          _1224_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1223_v135,(_987_v0).f32);
          let _1225_v137;
          _1225_v137 = _dafny.Set.fromElements((_995_v8).f32, (_987_v0).f32);
          let _1226_v138;
          _1226_v138 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,_1225_v137);
          let _1227_v140;
          let _init35 = function (_1228_i15) {
            return (_this).f26;
          };
          let _nw174 = Array((new BigNumber(15)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw174.length); _i0_35++) {
            _nw174[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1227_v140 = _nw174;
          let _1229_v141;
          _1229_v141 = _dafny.MultiSet.fromElements(_1227_v140);
          let _1230_v142;
          _1230_v142 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,(((_1229_v141).contains(_1227_v140)) ? ((_1229_v141).get(_1227_v140)) : (_1127_cf12)));
          let _1231_v143;
          _1231_v143 = _dafny.Map.Empty.slice().updateUnsafe(_1230_v142,new BigNumber(53));
          let _1232_v144;
          _1232_v144 = _dafny.MultiSet.fromElements((_987_v0).f32);
          let _1233_v145;
          let _nw175 = Array((new BigNumber(24)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = _1127_cf12;
          _nw175[(_dafny.ONE).toNumber()] = new BigNumber((_1224_v136).length);
          _nw175[(new BigNumber(2)).toNumber()] = (_991_v4).multipliedBy(_991_v4);
          _nw175[(new BigNumber(3)).toNumber()] = new BigNumber(((((_1226_v138).contains(new BigNumber(255))) ? ((_1226_v138).get(new BigNumber(255))) : (_dafny.Set.fromElements((_995_v8).f32)))).length);
          _nw175[(new BigNumber(4)).toNumber()] = new BigNumber((function () {
            let _coll59 = new _dafny.Set();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(158), new BigNumber(746))) {
              let _1234_v139 = _compr_59;
              if (((new BigNumber(158)).isLessThanOrEqualTo(_1234_v139)) && ((_1234_v139).isLessThan(new BigNumber(746)))) {
                _coll59.add(_module.__default.safeModuloInt(_1234_v139, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(858)), function (_1235_i14) {
                  return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cuhflyu")).length));
                })).length)));
              }
            }
            return _coll59;
          }()).length);
          _nw175[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_991_v4);
          _nw175[(new BigNumber(6)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(7)).toNumber()] = _991_v4;
          _nw175[(new BigNumber(8)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(9)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(10)).toNumber()] = _991_v4;
          _nw175[(new BigNumber(11)).toNumber()] = (_991_v4).minus(_991_v4);
          _nw175[(new BigNumber(12)).toNumber()] = _991_v4;
          _nw175[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(_1127_cf12, _1127_cf12);
          _nw175[(new BigNumber(14)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(15)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(16)).toNumber()] = (_1127_cf12).minus(new BigNumber((_1231_v143).length));
          _nw175[(new BigNumber(17)).toNumber()] = _1127_cf12;
          _nw175[(new BigNumber(18)).toNumber()] = new BigNumber((_1232_v144).cardinality());
          _nw175[(new BigNumber(19)).toNumber()] = _991_v4;
          _nw175[(new BigNumber(20)).toNumber()] = new BigNumber((_1230_v142).length);
          _nw175[(new BigNumber(21)).toNumber()] = _991_v4;
          _nw175[(new BigNumber(22)).toNumber()] = ((_this).fm1(globalState)).multipliedBy(_991_v4);
          _nw175[(new BigNumber(23)).toNumber()] = ((_dafny.ZERO).minus(_991_v4)).multipliedBy(_1127_cf12);
          _1233_v145 = _nw175;
          let _index160 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length));
          (_1233_v145)[_index160] = _1127_cf12;
          let _index161 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length));
          (_1233_v145)[_index161] = ((!(((_this).f27) === ((_this).f27))) ? (new BigNumber((_1222_v134).length)) : (_991_v4));
          (globalState).f16 = _991_v4;
          let _1236_v146;
          _1236_v146 = _dafny.Map.Empty.slice().updateUnsafe((_987_v0).f32,_1127_cf12);
          let _1237_v147;
          _1237_v147 = _module.D5.create_DC14((_1233_v145)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length))], _1127_cf12, !((_987_v0).f32));
          let _1238_v148;
          _1238_v148 = _dafny.Seq.of(_1237_v147, _1237_v147);
          let _1239_v149;
          _1239_v149 = _dafny.Map.Empty.slice().updateUnsafe((_987_v0).f32,(_995_v8).f32);
          let _1240_v150;
          _1240_v150 = _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.update(_module.__default.fm35((((_1236_v146).contains((_987_v0).f32)) ? ((_1236_v146).get((_987_v0).f32)) : ((_1233_v145)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length))])), globalState), _module.__default.safeIndex(_991_v4, new BigNumber((_module.__default.fm35((((_1236_v146).contains((_987_v0).f32)) ? ((_1236_v146).get((_987_v0).f32)) : ((_1233_v145)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length))])), globalState)).length)), _1237_v147), _1238_v148), !(_1239_v149).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_987_v0).f32)), (_995_v8).f32, (_995_v8).f32);
          _1240_v150 = _dafny.Seq.Concat(p2, _dafny.Seq.Concat(_dafny.Seq.of((_987_v0).f32, (_995_v8).f32), (_1059_v53).dtor_cf6));
          let _1241_v151;
          _1241_v151 = _dafny.MultiSet.fromElements(_1222_v134);
          let _index162 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length));
          (_1233_v145)[_index162] = (((_1241_v151).contains(_1222_v134)) ? ((_1241_v151).get(_1222_v134)) : ((_1233_v145)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length))]));
          let _1242_v152;
          let _nw176 = new _module.C1();
          _nw176.__ctor((_dafny.ZERO).minus((_1233_v145)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_1233_v145).length))]), _1127_cf12, (_987_v0).f32, (_this).f27);
          _1242_v152 = _nw176;
        } else {
          _1127_cf12 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1127_cf12, _1127_cf12));
          let _1243_v153;
          _1243_v153 = _dafny.MultiSet.fromElements(new BigNumber(878), _991_v4);
          let _1244_v154;
          let _out42;
          _out42 = (_this).m8(_module.D0.create_DC0((_987_v0).f32), (_995_v8).f32, _dafny.MultiSet.fromElements((_987_v0).f32, (_987_v0).f32, p0), _module.__default.safeModuloInt(_1127_cf12, new BigNumber(((_1243_v153).update(_1127_cf12, _module.__default.abs(_991_v4))).cardinality())), globalState);
          _1244_v154 = _out42;
          let _1245_v155;
          let _init36 = function (_1246_i16) {
            return (_1246_i16).plus(new BigNumber(121));
          };
          let _nw177 = Array((new BigNumber(29)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw177.length); _i0_36++) {
            _nw177[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1245_v155 = _nw177;
          let _1247_v156;
          _1247_v156 = _module.D3.create_DC10(_1245_v155);
          let _1248_v157;
          _1248_v157 = _dafny.Map.Empty.slice().updateUnsafe(_1247_v156,(_dafny.ZERO).minus((_dafny.ZERO).minus(_991_v4)));
          let _1249_v158;
          let _nw178 = new _module.C4();
          _nw178.__ctor((((_987_v0).f32) ? (_1248_v157) : (_1248_v157)), (_987_v0).f32, (_this).f27);
          _1249_v158 = _nw178;
          (globalState).f2 = (_995_v8).f32;
          let _1250_v159;
          let _nw179 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _1250_v159 = _nw179;
          let _1251_v160;
          _1251_v160 = _dafny.Seq.UnicodeFromString("hnfdufwb");
          let _1252_v161;
          _1252_v161 = _dafny.Map.Empty.slice().updateUnsafe(_1127_cf12,_1251_v160);
          let _1253_v162;
          _1253_v162 = _dafny.Seq.of(_1252_v161, _1252_v161, _1252_v161, _1252_v161);
          let _index163 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1250_v159).length));
          (_1250_v159)[_index163] = ((_1253_v162)[_module.__default.safeIndex(_1127_cf12, new BigNumber((_1253_v162).length))]).Merge(_1252_v161);
          let _1254_v163;
          _1254_v163 = _module.D5.create_DC13((_1128_v80).update((_this).f26, _dafny.Seq.of((_987_v0).f32)));
          let _1255_v164;
          _1255_v164 = _module.D5.create_DC15(_1254_v163);
          let _1256_v165;
          _1256_v165 = _module.D5.create_DC15(_module.D5.create_DC15(_1255_v164));
          let _1257_v166;
          _1257_v166 = _module.D5.create_DC15(_1255_v164);
          let _1258_v167;
          let _nw180 = Array((new BigNumber(6)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = _module.D5.create_DC15(_1254_v163);
          _nw180[(_dafny.ONE).toNumber()] = _1257_v166;
          _nw180[(new BigNumber(2)).toNumber()] = _1257_v166;
          _nw180[(new BigNumber(3)).toNumber()] = _1257_v166;
          _nw180[(new BigNumber(4)).toNumber()] = _1257_v166;
          _nw180[(new BigNumber(5)).toNumber()] = _1257_v166;
          _1258_v167 = _nw180;
          let _index164 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1250_v159).length));
          let _rhs178 = (function () {
            let _coll60 = new _dafny.Map();
            for (const _compr_60 of ((_993_v7).dtor_cf13).Keys.Elements) {
              let _1259_v168 = _compr_60;
              if (((_993_v7).dtor_cf13).contains(_1259_v168)) {
                _coll60.push([(_1259_v168).plus(new BigNumber(725)),_1251_v160]);
              }
            }
            return _coll60;
          }()).Merge((_1252_v161).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_991_v4),_1251_v160)));
          let _rhs179 = p0;
          let _rhs180 = _1258_v167;
          let _rhs181 = !((_987_v0).f32);
          let _lhs162 = _1250_v159;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1250_v159).length));
          let _lhs164 = globalState;
          _lhs162[_lhs163] = _rhs178;
          _1244_v154 = _rhs179;
          _1258_v167 = _rhs180;
          _lhs164.f2 = _rhs181;
        }
        if ((_this).f26) {
          let _1260_v169;
          _1260_v169 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(_1127_cf12)).plus(_1127_cf12),(_995_v8).f32);
          let _1261_v170;
          _1261_v170 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,_991_v4);
          _1260_v169 = (_1260_v169).update(new BigNumber((_1261_v170).length), (_987_v0).f32);
          let _1262_v171;
          let _nw181 = Array((new BigNumber(16)).toNumber()).fill(_module.D0.Default());
          _1262_v171 = _nw181;
          let _1263_v172;
          _1263_v172 = _module.D0.create_DC0((_this).f27);
          let _index165 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1262_v171).length));
          (_1262_v171)[_index165] = _1263_v172;
          let _1264_v173;
          _1264_v173 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_995_v8).f32);
          let _index166 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1262_v171).length));
          let _rhs182 = _module.D0.create_DC0((((_1264_v173).contains((_995_v8).f32)) ? ((_1264_v173).get((_995_v8).f32)) : ((_987_v0).f32)));
          let _rhs183 = _991_v4;
          let _lhs165 = _1262_v171;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1262_v171).length));
          let _lhs167 = globalState;
          _lhs165[_lhs166] = _rhs182;
          _lhs167.f16 = _rhs183;
          let _1265_v174;
          _1265_v174 = _dafny.Seq.of(_1127_cf12, _991_v4);
          let _1266_v175;
          _1266_v175 = _1265_v174;
          let _1267_v176;
          _1267_v176 = _dafny.Seq.of(_1266_v175, _1266_v175);
          let _1268_v177;
          _1268_v177 = _dafny.Seq.UnicodeFromString("iywdtfdjg");
          let _1269_v178;
          _1269_v178 = _module.D6.create_DC17(_1127_cf12, _1127_cf12, _991_v4, new BigNumber((_dafny.Seq.update(_1268_v177, _module.__default.safeIndex(new BigNumber(-75), new BigNumber((_1268_v177).length)), _1058_v52)).length), _991_v4);
          let _1270_v179;
          _1270_v179 = _dafny.Map.Empty.slice().updateUnsafe(_1267_v176,_1269_v178);
          _1270_v179 = _1270_v179;
          let _1271_v180;
          _1271_v180 = _module.D0.create_DC1((_987_v0).f32, (_995_v8).f32, new BigNumber((_1268_v177).length), p0);
          let _1272_v181;
          let _nw182 = new _module.C1();
          _nw182.__ctor(_991_v4, _991_v4, (_this).f26, (_this).f27);
          _1272_v181 = _nw182;
          let _1273_v182;
          _1273_v182 = _dafny.MultiSet.fromElements(_1272_v181);
          let _1274_v183;
          _1274_v183 = _dafny.Set.fromElements((_this).f26, (_995_v8).f32);
          let _1275_v184;
          _1275_v184 = _dafny.Map.Empty.slice().updateUnsafe(true,_1058_v52);
          let _1276_v185;
          _1276_v185 = _module.D5.create_DC14(_1272_v181.f37, new BigNumber((_1268_v177).length), (_995_v8).f32);
          let _1277_v186;
          let _nw183 = Array((new BigNumber(20)).toNumber());
          _nw183[(_dafny.ZERO).toNumber()] = !(!(!(new BigNumber(-994)).isEqualTo(_991_v4)));
          _nw183[(_dafny.ONE).toNumber()] = (_995_v8).f32;
          _nw183[(new BigNumber(2)).toNumber()] = (_this).fm2(true, (_this).fm1(globalState), globalState);
          _nw183[(new BigNumber(3)).toNumber()] = (_995_v8).f32;
          _nw183[(new BigNumber(4)).toNumber()] = (_995_v8).f32;
          _nw183[(new BigNumber(5)).toNumber()] = !((_987_v0).fm7(_1261_v170, _1265_v174, (_this).f27, globalState)) || (true);
          _nw183[(new BigNumber(6)).toNumber()] = !((_987_v0).f32) || ((_1271_v180).dtor_cf1);
          _nw183[(new BigNumber(7)).toNumber()] = (_this).f26;
          _nw183[(new BigNumber(8)).toNumber()] = (_this).f27;
          _nw183[(new BigNumber(9)).toNumber()] = (_995_v8).f32;
          _nw183[(new BigNumber(10)).toNumber()] = ((_987_v0).fm7(_1261_v170, _1265_v174, !(p0), globalState)) && ((_this).f26);
          _nw183[(new BigNumber(11)).toNumber()] = ((((_1273_v182).contains(_1272_v181)) ? ((_1273_v182).get(_1272_v181)) : (new BigNumber((_1265_v174).length)))).isLessThan(new BigNumber((_1274_v183).length));
          _nw183[(new BigNumber(12)).toNumber()] = (_987_v0).fm6(_1058_v52, !((_987_v0).fm7(_1261_v170, _dafny.Seq.of(new BigNumber((_1275_v184).length)), (_987_v0).f32, globalState)), globalState);
          _nw183[(new BigNumber(13)).toNumber()] = true;
          _nw183[(new BigNumber(14)).toNumber()] = !(!((_this).f27)) || ((_987_v0).f32);
          _nw183[(new BigNumber(15)).toNumber()] = (_1276_v185).dtor_cf27;
          _nw183[(new BigNumber(16)).toNumber()] = ((p0) ? ((_995_v8).f32) : (p0));
          _nw183[(new BigNumber(17)).toNumber()] = (!(!(p0))) === ((_987_v0).f32);
          _nw183[(new BigNumber(18)).toNumber()] = (_995_v8).f32;
          _nw183[(new BigNumber(19)).toNumber()] = _dafny.areEqual(p2, _dafny.Seq.of(p0, p0, (_995_v8).f32));
          _1277_v186 = _nw183;
          let _index167 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1277_v186).length));
          (_1277_v186)[_index167] = (_995_v8).f32;
          let _index168 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_1277_v186).length));
          (_1277_v186)[_index168] = !((_987_v0).f32);
          let _1278_v187;
          let _nw184 = Array((new BigNumber(15)).toNumber()).fill(_module.D1.Default());
          _1278_v187 = _nw184;
          let _1279_v188;
          _1279_v188 = _dafny.Seq.of(_1278_v187);
          (globalState).f22 = !_dafny.areEqual(_1279_v188, _1279_v188);
        } else {
          (globalState).f12 = _1127_cf12;
          let _1280_v189;
          _1280_v189 = _dafny.Seq.of(new BigNumber(911), new BigNumber((p2).length));
          let _1281_v190;
          _1281_v190 = _dafny.Seq.of(_1280_v189);
          let _1282_v191;
          let _nw185 = new _module.C2();
          _nw185.__ctor(_1281_v190, (_987_v0).f32, (p2)[_module.__default.safeIndex(_991_v4, new BigNumber((p2).length))]);
          _1282_v191 = _nw185;
          let _1283_v192;
          _1283_v192 = _module.D7.create_DC18(_1282_v191);
          let _1284_v193;
          _1284_v193 = _dafny.MultiSet.fromElements(_1283_v192);
          let _1285_v194;
          _1285_v194 = _dafny.Seq.of(_1284_v193);
          let _1286_v195;
          _1286_v195 = _dafny.Seq.of(_1283_v192, _1283_v192);
          (globalState).f24 = !(((_dafny.MultiSet.FromArray(_1286_v195)).update(_1283_v192, _module.__default.abs(_1127_cf12))).IsSubsetOf((_1285_v194)[_module.__default.safeIndex(new BigNumber(38), new BigNumber((_1285_v194).length))]));
          let _1287_v196;
          let _nw186 = Array((new BigNumber(5)).toNumber()).fill(false);
          _1287_v196 = _nw186;
          let _rhs184 = _1287_v196;
          let _rhs185 = !((_this).f27) || ((_995_v8).f32);
          let _rhs186 = (_this).fm1(globalState);
          let _lhs168 = globalState;
          let _lhs169 = globalState;
          _1287_v196 = _rhs184;
          _lhs168.f2 = _rhs185;
          _lhs169.f16 = _rhs186;
          let _1288_v197;
          let _nw187 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _1288_v197 = _nw187;
          let _index169 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_1288_v197).length));
          (_1288_v197)[_index169] = _1127_cf12;
          let _index170 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_1288_v197).length));
          (_1288_v197)[_index170] = _1127_cf12;
          let _1289_v198;
          _1289_v198 = _dafny.Set.fromElements(_1288_v197);
          _1289_v198 = (_dafny.Set.fromElements(_1288_v197, _1288_v197, _1288_v197, _1288_v197)).Intersect(_dafny.Set.fromElements(_1288_v197, _1288_v197, _1288_v197));
        }
      } else {
        let _1290___mcc_h14 = (_source18).cf6;
        let _1291_cf6 = _1290___mcc_h14;
        let _1292_v199;
        _1292_v199 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,p0);
        (globalState).f15 = (((_1292_v199).contains((_this).f27)) ? ((_1292_v199).get((_this).f27)) : ((p2)[_module.__default.safeIndex(new BigNumber(450), new BigNumber((p2).length))]));
        let _1293_v200;
        _1293_v200 = _dafny.Map.Empty.slice().updateUnsafe((_995_v8).f32,_991_v4);
        let _1294_v201;
        _1294_v201 = _dafny.Seq.of(_991_v4);
        let _1295_v202;
        _1295_v202 = _dafny.Map.Empty.slice().updateUnsafe((_987_v0).fm7(_1293_v200, _1294_v201, (p2)[_module.__default.safeIndex(_991_v4, new BigNumber((p2).length))], globalState),_dafny.Map.Empty.slice().updateUnsafe(_991_v4,new BigNumber(-45)));
        let _1296_v203;
        _1296_v203 = _dafny.Map.Empty.slice().updateUnsafe(_991_v4,new BigNumber(-349));
        let _1297_v204;
        _1297_v204 = _dafny.Map.Empty.slice().updateUnsafe((((_1296_v203).contains(_991_v4)) ? ((_1296_v203).get(_991_v4)) : (new BigNumber(912))),new BigNumber(2));
        let _1298_v205;
        _1298_v205 = _dafny.Seq.UnicodeFromString("gwqodhw");
        let _1299_v206;
        _1299_v206 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("rphevkuq")).length),(_987_v0).fm6(_1058_v52, (_987_v0).f32, globalState));
        let _1300_v207;
        _1300_v207 = _dafny.MultiSet.fromElements(_1294_v201, _1294_v201);
        let _1301_v208;
        let _nw188 = new _module.C3();
        _nw188.__ctor(_1300_v207, (_this).f27, (_987_v0).f32);
        _1301_v208 = _nw188;
        let _1302_v209;
        let _nw189 = Array((new BigNumber(22)).toNumber());
        _nw189[(_dafny.ZERO).toNumber()] = _991_v4;
        _nw189[(_dafny.ONE).toNumber()] = _991_v4;
        _nw189[(new BigNumber(2)).toNumber()] = new BigNumber((_1292_v199).length);
        _nw189[(new BigNumber(3)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(4)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(5)).toNumber()] = new BigNumber((_1299_v206).length);
        _nw189[(new BigNumber(6)).toNumber()] = new BigNumber((_1298_v205).length);
        _nw189[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.of((_987_v0).f32)).length);
        _nw189[(new BigNumber(8)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(9)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(10)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(11)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(12)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(13)).toNumber()] = new BigNumber(138);
        _nw189[(new BigNumber(14)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(15)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(16)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(17)).toNumber()] = new BigNumber(61);
        _nw189[(new BigNumber(18)).toNumber()] = new BigNumber(540);
        _nw189[(new BigNumber(19)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(20)).toNumber()] = _991_v4;
        _nw189[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.of(_1301_v208)).length);
        _1302_v209 = _nw189;
        let _1303_v210;
        _1303_v210 = _module.D3.create_DC10(_1302_v209);
        let _1304_v211;
        let _nw190 = new _module.C4();
        _nw190.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1303_v210,new BigNumber((_1298_v205).length)), (_this).f27, (_995_v8).f32);
        _1304_v211 = _nw190;
        let _1305_v212;
        _1305_v212 = _dafny.Map.Empty.slice().updateUnsafe(_1304_v211,(_this).f27);
        let _1306_v213;
        _1306_v213 = _dafny.Map.Empty.slice().updateUnsafe(_1294_v201,_991_v4);
        let _1307_v214;
        let _nw191 = Array((new BigNumber(29)).toNumber());
        _nw191[(_dafny.ZERO).toNumber()] = _991_v4;
        _nw191[(_dafny.ONE).toNumber()] = new BigNumber(-872);
        _nw191[(new BigNumber(2)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(3)).toNumber()] = new BigNumber(((_1295_v202).Merge(_dafny.Map.Empty.slice().updateUnsafe((_987_v0).f32,_1296_v203))).length);
        _nw191[(new BigNumber(4)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(5)).toNumber()] = new BigNumber(-557);
        _nw191[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-248), _991_v4);
        _nw191[(new BigNumber(7)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(8)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((((_this).f26) ? (_dafny.Seq.update(_1298_v205, _module.__default.safeIndex(new BigNumber((_1292_v199).length), new BigNumber((_1298_v205).length)), _1058_v52)) : (_1298_v205))).length));
        _nw191[(new BigNumber(10)).toNumber()] = ((_dafny.ZERO).minus(_991_v4)).minus(_991_v4);
        _nw191[(new BigNumber(11)).toNumber()] = (new BigNumber(354)).minus((_1294_v201)[_module.__default.safeIndex(_991_v4, new BigNumber((_1294_v201).length))]);
        _nw191[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_991_v4, new BigNumber((_1299_v206).length)));
        _nw191[(new BigNumber(13)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_1298_v205, _module.__default.safeIndex(_991_v4, new BigNumber((_1298_v205).length)), _1058_v52)).length);
        _nw191[(new BigNumber(15)).toNumber()] = new BigNumber(850);
        _nw191[(new BigNumber(16)).toNumber()] = new BigNumber(942);
        _nw191[(new BigNumber(17)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(18)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_1305_v212).Merge(_1305_v212)).length));
        _nw191[(new BigNumber(20)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(21)).toNumber()] = new BigNumber((_1306_v213).length);
        _nw191[(new BigNumber(22)).toNumber()] = (((_995_v8).f32) ? (new BigNumber(196)) : (new BigNumber(664)));
        _nw191[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(_991_v4, _991_v4);
        _nw191[(new BigNumber(24)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(25)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(26)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(27)).toNumber()] = _991_v4;
        _nw191[(new BigNumber(28)).toNumber()] = new BigNumber(480);
        _1307_v214 = _nw191;
        _1307_v214 = _1302_v209;
        _1305_v212 = _1305_v212;
        let _1308_v215;
        let _init37 = function (_1309_i17) {
          return _dafny.Seq.of((_this).f26, (_this).f27);
        };
        let _nw192 = Array((new BigNumber(8)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw192.length); _i0_37++) {
          _nw192[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1308_v215 = _nw192;
        let _index171 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1308_v215).length));
        (_1308_v215)[_index171] = _1291_cf6;
        let _index172 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1308_v215).length));
        (_1308_v215)[_index172] = _dafny.Seq.Concat(p2, _1291_cf6);
      }
      return;
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1310_i0;
      _1310_i0 = _dafny.ZERO;
      L6: {
        while ((_this).f27) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1310_i0)) {
              break L6;
            }
            _1310_i0 = (_1310_i0).plus(_dafny.ONE);
            let _1311_v0;
            let _nw193 = new _module.C1();
            _nw193.__ctor(p2, p2, (_this).f26, p1);
            _1311_v0 = _nw193;
            let _1312_v1;
            let _nw194 = Array((new BigNumber(3)).toNumber());
            _nw194[(_dafny.ZERO).toNumber()] = _1311_v0;
            _nw194[(_dafny.ONE).toNumber()] = _1311_v0;
            _nw194[(new BigNumber(2)).toNumber()] = _1311_v0;
            _1312_v1 = _nw194;
            let _1313_v2;
            _1313_v2 = _dafny.Seq.of(_1312_v1);
            (globalState).f2 = _dafny.Seq.IsPrefixOf(_1313_v2, _1313_v2);
            let _1314_v3;
            _1314_v3 = _dafny.Seq.UnicodeFromString("jn");
            let _1315_v4;
            _1315_v4 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(825)), function (_1316_i1) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            }), _dafny.Seq.UnicodeFromString("yknlqwi")), _1314_v3);
            (_1311_v0).f37 = new BigNumber((_1315_v4).length);
            let _1317_v5;
            _1317_v5 = _dafny.Seq.of(_1314_v3);
            _1317_v5 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1317_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), ((_1318_v3) => function (_1319_i2) {
              return _1318_v3;
            })(_1314_v3))), _1317_v5), _module.__default.safeIndex(_1311_v0.f37, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1317_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), ((_1320_v3) => function (_1321_i2) {
              return _1320_v3;
            })(_1314_v3))), _1317_v5)).length)), _1314_v3);
            let _1322_v6;
            let _nw195 = Array((new BigNumber(5)).toNumber()).fill([]);
            _1322_v6 = _nw195;
            let _1323_v7;
            let _nw196 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
            _1323_v7 = _nw196;
            let _index173 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1322_v6).length));
            (_1322_v6)[_index173] = _1323_v7;
            let _index174 = _module.__default.safeIndex(new BigNumber(961), new BigNumber((_1322_v6).length));
            (_1322_v6)[_index174] = _1323_v7;
          }
        }
      }
      let _1324_v8;
      let _init38 = function (_1325_i3) {
        return (_this).f26;
      };
      let _nw197 = Array((new BigNumber(6)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw197.length); _i0_38++) {
        _nw197[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1324_v8 = _nw197;
      let _index175 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
      (_1324_v8)[_index175] = (_this).f26;
      let _1326_v9;
      let _nw198 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1326_v9 = _nw198;
      let _index176 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1324_v8).length));
      (_1324_v8)[_index176] = (_this).f26;
      let _1327_v10;
      _1327_v10 = _dafny.Seq.of(p2);
      let _1328_v11;
      _1328_v11 = _dafny.MultiSet.fromElements(_1327_v10, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-396)), ((_1329_p2) => function (_1330_i4) {
        return _1329_p2;
      })(p2)), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-396)), ((_1331_p2) => function (_1332_i4) {
        return _1331_p2;
      })(p2))).length)), p2));
      let _index177 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
      let _index178 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1324_v8).length));
      let _rhs187 = false;
      let _rhs188 = (_this).f27;
      let _rhs189 = _1326_v9;
      let _rhs190 = p1;
      let _rhs191 = ((_1328_v11).Union(_1328_v11)).IsSubsetOf(_1328_v11);
      let _lhs170 = globalState;
      let _lhs171 = _1324_v8;
      let _lhs172 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
      let _lhs173 = _1324_v8;
      let _lhs174 = _module.__default.safeIndex(new BigNumber(234), new BigNumber((_1324_v8).length));
      let _lhs175 = globalState;
      _lhs170.f15 = _rhs187;
      _lhs171[_lhs172] = _rhs188;
      _1326_v9 = _rhs189;
      _lhs173[_lhs174] = _rhs190;
      _lhs175.f2 = _rhs191;
      let _hi11 = p2;
      for (let _1333_i5 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).fm1(globalState)), p2); _1333_i5.isLessThan(_hi11); _1333_i5 = _1333_i5.plus(_dafny.ONE)) {
        (globalState).f12 = p2;
        let _1334_v12;
        _1334_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2((_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))], (_this).fm1(globalState), globalState),p2);
        (globalState).f16 = (((_1333_i5).isLessThanOrEqualTo(p2)) ? ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(463)), function (_1335_i6) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        })).length)).plus(new BigNumber((_1334_v12).length))) : (new BigNumber(887)));
        let _1336_v13;
        _1336_v13 = _dafny.Seq.of(!(p1), (_this).f26);
        (globalState).f15 = ((!(!(true))) ? ((_1336_v13)[_module.__default.safeIndex(p2, new BigNumber((_1336_v13).length))]) : (p0));
        if (p1) {
          let _1337_v14;
          _1337_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1338_v15;
          _1338_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1337_v14).length));
          let _1339_v16;
          _1339_v16 = _module.D2.create_DC6(_1338_v15);
          let _1340_v17;
          _1340_v17 = _dafny.Seq.of(_1339_v16, _module.D2.create_DC6(_1338_v15));
          let _1341_v18;
          _1341_v18 = _dafny.MultiSet.fromElements(_1340_v17, _1340_v17, _1340_v17, _1340_v17, _1340_v17);
          let _index179 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
          (_1324_v8)[_index179] = (_1341_v18).IsSubsetOf((((_this).f26) ? (_dafny.MultiSet.fromElements(_dafny.Seq.update(_1340_v17, _module.__default.safeIndex(new BigNumber(700), new BigNumber((_1340_v17).length)), _1339_v16))) : (_module.__default.fm36(p2, (_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))], !((_this).fm2(p1, p2, globalState)), (_this).f26, globalState))));
          let _1342_v19;
          _1342_v19 = _module.D3.create_DC10(_1326_v9);
          let _1343_v20;
          _1343_v20 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1324_v8);
          let _1344_v21;
          _1344_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1342_v19,new BigNumber((_1343_v20).length));
          let _1345_v22;
          let _nw199 = new _module.C4();
          _nw199.__ctor(_1344_v21, p1, p1);
          _1345_v22 = _nw199;
          let _1346_v23;
          _1346_v23 = _dafny.Seq.UnicodeFromString("gegvu");
          (globalState).f6 = _dafny.Seq.Concat(_1346_v23, _1346_v23);
          let _1347_v24;
          _1347_v24 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(117)), ((_1348_p2) => function (_1349_i7) {
            return _1348_p2;
          })(p2)), _dafny.Seq.of(p2), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1338_v15).length)), _1333_i5), _1327_v10, _dafny.Seq.of(_1333_i5, _1333_i5));
          let _1350_v25;
          let _nw200 = new _module.C2();
          _nw200.__ctor(_1347_v24, (_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))], (_this).f27);
          _1350_v25 = _nw200;
          let _1351_v26;
          _1351_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1327_v10,_1350_v25);
          let _1352_v27;
          let _nw201 = Array((new BigNumber(3)).toNumber());
          _nw201[(_dafny.ZERO).toNumber()] = (_1351_v26).Merge(_1351_v26);
          _nw201[(_dafny.ONE).toNumber()] = _1351_v26;
          _nw201[(new BigNumber(2)).toNumber()] = _1351_v26;
          _1352_v27 = _nw201;
          let _index180 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1352_v27).length));
          (_1352_v27)[_index180] = _1351_v26;
          let _index181 = _module.__default.safeIndex(new BigNumber(450), new BigNumber((_1352_v27).length));
          (_1352_v27)[_index181] = _1351_v26;
          let _1353_v28;
          _1353_v28 = _module.D0.create_DC0((_this).f26);
          let _index182 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1326_v9).length));
          (_1326_v9)[_index182] = ((_1350_v25).fm4(_1353_v28, _1333_i5, p2, globalState)).minus(_1333_i5);
          let _index183 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1326_v9).length));
          let _rhs192 = (_dafny.ZERO).minus((((_1336_v13)[_module.__default.safeIndex(_1333_i5, new BigNumber((_1336_v13).length))]) ? (new BigNumber(606)) : ((_1333_i5).minus(p2))));
          let _rhs193 = _1333_i5;
          let _lhs176 = globalState;
          let _lhs177 = _1326_v9;
          let _lhs178 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1326_v9).length));
          _lhs176.f12 = _rhs192;
          _lhs177[_lhs178] = _rhs193;
        } else {
          let _1354_v29;
          let _nw202 = Array((new BigNumber(6)).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = _1326_v9;
          _nw202[(_dafny.ONE).toNumber()] = _1326_v9;
          _nw202[(new BigNumber(2)).toNumber()] = _1326_v9;
          _nw202[(new BigNumber(3)).toNumber()] = _1326_v9;
          _nw202[(new BigNumber(4)).toNumber()] = _1326_v9;
          _nw202[(new BigNumber(5)).toNumber()] = _1326_v9;
          _1354_v29 = _nw202;
          let _index184 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1354_v29).length));
          (_1354_v29)[_index184] = _1326_v9;
          let _index185 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1354_v29).length));
          (_1354_v29)[_index185] = _1326_v9;
          let _1355_v30;
          _1355_v30 = _dafny.Seq.UnicodeFromString("ctfdse");
          let _1356_v31;
          _1356_v31 = _dafny.Seq.of((_1354_v29)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1354_v29).length))], _1326_v9, _1326_v9);
          let _rhs194 = (_1354_v29)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_1354_v29).length))];
          let _rhs195 = _1355_v30;
          let _rhs196 = new BigNumber(812);
          let _rhs197 = (_1356_v31)[_module.__default.safeIndex(p2, new BigNumber((_1356_v31).length))];
          let _lhs179 = globalState;
          let _lhs180 = globalState;
          _1326_v9 = _rhs194;
          _lhs179.f6 = _rhs195;
          _lhs180.f16 = _rhs196;
          _1326_v9 = _rhs197;
          (globalState).f16 = (_this).fm1(globalState);
          let _1357_v32;
          let _nw203 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
          _1357_v32 = _nw203;
          let _1358_v33;
          let _nw204 = Array((new BigNumber(4)).toNumber()).fill(_module.D1.Default());
          _1358_v33 = _nw204;
          let _1359_v34;
          _1359_v34 = _dafny.Seq.of(_1358_v33, _1358_v33);
          let _index186 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1357_v32).length));
          (_1357_v32)[_index186] = _1359_v34;
          let _1360_v35;
          _1360_v35 = _dafny.MultiSet.fromElements(p3, p3);
          let _index187 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1357_v32).length));
          let _rhs198 = _dafny.Seq.Concat(_1359_v34, _1359_v34);
          let _rhs199 = _1333_i5;
          let _rhs200 = (_1360_v35).IsSubsetOf(_1360_v35);
          let _lhs181 = _1357_v32;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_1357_v32).length));
          let _lhs183 = globalState;
          let _lhs184 = globalState;
          _lhs181[_lhs182] = _rhs198;
          _lhs183.f16 = _rhs199;
          _lhs184.f15 = _rhs200;
          let _1361_v36;
          let _nw205 = new _module.C1();
          _nw205.__ctor(p2, _1333_i5, ((_dafny.ZERO).minus(new BigNumber(-125))).isEqualTo(new BigNumber((p3).cardinality())), (_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))]);
          _1361_v36 = _nw205;
        }
      }
      let _index188 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
      (_1324_v8)[_index188] = (_this).f26;
      let _1362_v37;
      _1362_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(946),true);
      (globalState).f24 = (new BigNumber((_1362_v37).length)).isEqualTo(p2);
      if ((_this).f26) {
        let _1363_v38;
        _1363_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p3).cardinality()),p2);
        let _1364_v39;
        _1364_v39 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(p2, new BigNumber((_1363_v38).length)), p2);
        _1364_v39 = (_1364_v39).Difference((_1364_v39).Difference(_1364_v39));
        (globalState).f16 = new BigNumber((_1327_v10).length);
        let _1365_v40;
        _1365_v40 = _dafny.Seq.of((_this).f26);
        _1327_v10 = _dafny.Seq.of(new BigNumber((_1365_v40).length), p2, (new BigNumber(599)).minus(p2));
        (globalState).f22 = p1;
        let _1366_v41;
        let _nw206 = Array((new BigNumber(21)).toNumber()).fill([]);
        _1366_v41 = _nw206;
        _1366_v41 = _1366_v41;
      } else {
        let _index189 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length));
        (_1326_v9)[_index189] = p2;
        let _index190 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length));
        (_1326_v9)[_index190] = p2;
        let _1367_v42;
        _1367_v42 = _module.D0.create_DC0(p0);
        let _1368_v43;
        _1368_v43 = _dafny.Seq.of(true);
        let _1369_v44;
        _1369_v44 = _dafny.MultiSet.fromElements(p1);
        let _1370_v45;
        let _out43;
        _out43 = (_this).m8(_1367_v42, (_1368_v43)[_module.__default.safeIndex((_this).fm1(globalState), new BigNumber((_1368_v43).length))], _1369_v44, (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], globalState);
        _1370_v45 = _out43;
        let _1371_v46;
        _1371_v46 = _dafny.MultiSet.fromElements(_1324_v8, _1324_v8);
        let _1372_v47;
        _1372_v47 = _dafny.Seq.UnicodeFromString("nhfpskf");
        let _1373_v48;
        _1373_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1371_v46).Intersect(_1371_v46)).cardinality()),(((_this).f27) ? (_1372_v47) : (_1372_v47)));
        _1373_v48 = (_1373_v48).update((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("t"), _module.__default.fm22(p2, (_this).fm2(p0, (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], globalState), globalState)));
        let _1374_v49;
        _1374_v49 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))]);
        let _1375_v50;
        _1375_v50 = _module.D7.create_DC19((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], _1369_v44, (_this).f27, _1374_v49);
        if ((_1375_v50).dtor_cf38) {
          let _index191 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length));
          (_1326_v9)[_index191] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          (globalState).f12 = p2;
          let _1376_v51;
          _1376_v51 = _dafny.Map.Empty.slice().updateUnsafe((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))],_dafny.Seq.of((_1375_v50).dtor_cf38));
          let _1377_v52;
          _1377_v52 = new _dafny.CodePoint('h'.codePointAt(0));
          _1376_v51 = (_1376_v51).update((_this).fm1(globalState), _module.__default.fm11((_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))], (_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))], _1377_v52, globalState));
          let _1378_v53;
          let _nw207 = Array((new BigNumber(23)).toNumber()).fill([]);
          _1378_v53 = _nw207;
          let _1379_v54;
          let _nw208 = Array((new BigNumber(25)).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("fbnprm");
          _nw208[(_dafny.ONE).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(2)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(3)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(4)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(5)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("c");
          _nw208[(new BigNumber(7)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(8)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(9)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(10)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(11)).toNumber()] = _module.__default.fm22(p2, _1370_v45, globalState);
          _nw208[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("hluwry");
          _nw208[(new BigNumber(13)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(14)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(15)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(16)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(17)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(18)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(19)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(20)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("pkxwgoox"), _module.__default.safeIndex((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], new BigNumber((_dafny.Seq.UnicodeFromString("pkxwgoox")).length)), _1377_v52);
          _nw208[(new BigNumber(22)).toNumber()] = _1372_v47;
          _nw208[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), ((_1380_v52) => function (_1381_i8) {
            return _1380_v52;
          })(_1377_v52));
          _nw208[(new BigNumber(24)).toNumber()] = _1372_v47;
          _1379_v54 = _nw208;
          let _1382_v55;
          _1382_v55 = _dafny.Seq.of(_1379_v54, _1379_v54);
          let _index192 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1378_v53).length));
          (_1378_v53)[_index192] = (((_this).fm2(true, (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], globalState)) ? (_1379_v54) : ((_1382_v55)[_module.__default.safeIndex(p2, new BigNumber((_1382_v55).length))]));
          let _index193 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1378_v53).length));
          (_1378_v53)[_index193] = _1379_v54;
          let _1383_v56;
          _1383_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_dafny.Seq.of((_this).f27, _1370_v45));
          let _1384_v57;
          _1384_v57 = _module.D5.create_DC13(_1383_v56);
          _1384_v57 = _1384_v57;
        } else {
          let _1385_v58;
          _1385_v58 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1386_v59;
          let _nw209 = Array((new BigNumber(25)).toNumber());
          _nw209[(_dafny.ZERO).toNumber()] = p2;
          _nw209[(_dafny.ONE).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(2)).toNumber()] = p2;
          _nw209[(new BigNumber(3)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(4)).toNumber()] = p2;
          _nw209[(new BigNumber(5)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(6)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))]);
          _nw209[(new BigNumber(8)).toNumber()] = new BigNumber((_1327_v10).length);
          _nw209[(new BigNumber(9)).toNumber()] = p2;
          _nw209[(new BigNumber(10)).toNumber()] = p2;
          _nw209[(new BigNumber(11)).toNumber()] = p2;
          _nw209[(new BigNumber(12)).toNumber()] = new BigNumber(606);
          _nw209[(new BigNumber(13)).toNumber()] = p2;
          _nw209[(new BigNumber(14)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(15)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(16)).toNumber()] = p2;
          _nw209[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.of(_1385_v58)).length);
          _nw209[(new BigNumber(18)).toNumber()] = p2;
          _nw209[(new BigNumber(19)).toNumber()] = p2;
          _nw209[(new BigNumber(20)).toNumber()] = (_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))];
          _nw209[(new BigNumber(21)).toNumber()] = p2;
          _nw209[(new BigNumber(22)).toNumber()] = new BigNumber((_1372_v47).length);
          _nw209[(new BigNumber(23)).toNumber()] = p2;
          _nw209[(new BigNumber(24)).toNumber()] = p2;
          _1386_v59 = _nw209;
          let _1387_v60;
          _1387_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1386_v59,new BigNumber((_1372_v47).length));
          let _1388_v61;
          _1388_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1369_v44,_dafny.Set.fromElements((_this).f27));
          let _1389_v62;
          _1389_v62 = _module.D5.create_DC14(new BigNumber(((((_1388_v61).contains(_1369_v44)) ? ((_1388_v61).get(_1369_v44)) : (_dafny.Set.fromElements((_this).f26, (_this).f26)))).length), p2, (_1368_v43)[_module.__default.safeIndex((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], new BigNumber((_1368_v43).length))]);
          let _index194 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length));
          let _rhs201 = (((_this).f27) ? (new BigNumber(((_1387_v60).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1386_v59,(_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))]))).length)) : (p2));
          let _rhs202 = (_1389_v62).dtor_cf26;
          let _rhs203 = ((!((p2).isLessThanOrEqualTo(p2))) ? (_1370_v45) : ((_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))]));
          let _lhs185 = globalState;
          let _lhs186 = _1326_v9;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length));
          let _lhs188 = globalState;
          _lhs185.f12 = _rhs201;
          _lhs186[_lhs187] = _rhs202;
          _lhs188.f24 = _rhs203;
          let _1390_v63;
          let _nw210 = new _module.C0();
          _nw210.__ctor((_this).f26);
          _1390_v63 = _nw210;
          let _1391_v64;
          _1391_v64 = _dafny.Set.fromElements((_this).f26);
          let _1392_v65;
          _1392_v65 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(262)), ((_1393_v58) => function (_1394_i9) {
            return _module.D3.create_DC11(_1393_v58);
          })(_1385_v58)),_1390_v63);
          let _1395_v66;
          _1395_v66 = _module.D3.create_DC11(_1385_v58);
          let _rhs204 = ((((_this).f27) ? ((_1391_v64)) : (_1391_v64))).contains((_1390_v63).f32);
          let _rhs205 = (((_1392_v65).contains(_dafny.Seq.of(_1395_v66, _module.D3.create_DC11(_1385_v58), _1395_v66, _1395_v66))) ? ((_1392_v65).get(_dafny.Seq.of(_1395_v66, _module.D3.create_DC11(_1385_v58), _1395_v66, _1395_v66))) : (_1390_v63));
          let _lhs189 = globalState;
          _lhs189.f2 = _rhs204;
          _1390_v63 = _rhs205;
          (globalState).f6 = _1372_v47;
          let _1396_v67;
          let _nw211 = Array((new BigNumber(18)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mxwll"), _1372_v47);
          _nw211[(_dafny.ONE).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(2)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_1372_v47, _module.__default.safeIndex((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], new BigNumber((_1372_v47).length)), _1385_v58);
          _nw211[(new BigNumber(4)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(5)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1372_v47, _1372_v47);
          _nw211[(new BigNumber(7)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("lawmo");
          _nw211[(new BigNumber(9)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(10)).toNumber()] = _module.__default.fm22((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], p0, globalState);
          _nw211[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), ((_1397_v58) => function (_1398_i10) {
            return _1397_v58;
          })(_1385_v58));
          _nw211[(new BigNumber(12)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), function (_1399_i11) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          });
          _nw211[(new BigNumber(14)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(15)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(16)).toNumber()] = _1372_v47;
          _nw211[(new BigNumber(17)).toNumber()] = _1372_v47;
          _1396_v67 = _nw211;
          let _init39 = ((_1400_v58) => function (_1401_i12) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), ((_1402_v58) => function (_1403_i13) {
              return _1402_v58;
            })(_1400_v58));
          })(_1385_v58);
          let _nw212 = Array((new BigNumber(3)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw212.length); _i0_39++) {
            _nw212[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1396_v67 = _nw212;
          let _1404_v68;
          _1404_v68 = _1324_v8;
          let _1405_v69;
          _1405_v69 = _dafny.Seq.of(_1324_v8, (_1404_v68), _1324_v8, _1324_v8);
          let _1406_v70;
          _1406_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("hnkb"),_dafny.Seq.of(_1324_v8));
          let _1407_v71;
          let _nw213 = Array((new BigNumber(5)).toNumber());
          _nw213[(_dafny.ZERO).toNumber()] = _1405_v69;
          _nw213[(_dafny.ONE).toNumber()] = _1405_v69;
          _nw213[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1405_v69, _1405_v69);
          _nw213[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1405_v69, _1405_v69);
          _nw213[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1405_v69, _dafny.Seq.update((((_1406_v70).contains(_module.__default.fm22((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], (_1390_v63).fm6(_1385_v58, (_1390_v63).f32, globalState), globalState))) ? ((_1406_v70).get(_module.__default.fm22((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], (_1390_v63).fm6(_1385_v58, (_1390_v63).f32, globalState), globalState))) : (_1405_v69)), _module.__default.safeIndex(p2, new BigNumber(((((_1406_v70).contains(_module.__default.fm22((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], (_1390_v63).fm6(_1385_v58, (_1390_v63).f32, globalState), globalState))) ? ((_1406_v70).get(_module.__default.fm22((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))], (_1390_v63).fm6(_1385_v58, (_1390_v63).f32, globalState), globalState))) : (_1405_v69))).length)), _1324_v8));
          _1407_v71 = _nw213;
          let _index195 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1407_v71).length));
          (_1407_v71)[_index195] = _1405_v69;
          let _index196 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((_1407_v71).length));
          (_1407_v71)[_index196] = _dafny.Seq.Concat(_1405_v69, _dafny.Seq.Concat(_1405_v69, _1405_v69));
        }
        if (p1) {
          let _1408_v72;
          _1408_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1372_v47,_1324_v8);
          _1408_v72 = (_1408_v72).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_1409_i14) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          }), _1324_v8);
          let _1410_v73;
          let _nw214 = new _module.C3();
          _nw214.__ctor(_1328_v11, (p1) === (_1370_v45), !(_1370_v45) || (_1370_v45));
          _1410_v73 = _nw214;
          let _1411_v74;
          let _out44;
          _out44 = (_this).m8(_1367_v42, (_this).f26, _dafny.MultiSet.fromElements((_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))]), _module.__default.safeDivisionInt(new BigNumber(553), new BigNumber((_1372_v47).length)), globalState);
          _1411_v74 = _out44;
          (globalState).f12 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))]), p2);
          let _1412_v75;
          _1412_v75 = new _dafny.CodePoint('h'.codePointAt(0));
          (globalState).f19 = _1412_v75;
        } else {
          let _1413_v76;
          _1413_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1362_v37,p1);
          let _1414_v77;
          let _nw215 = new _module.C1();
          _nw215.__ctor(new BigNumber((_1413_v76).length), p2, p0, false);
          _1414_v77 = _nw215;
          let _1415_v78;
          _1415_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1414_v77,p3);
          let _1416_v79;
          _1416_v79 = _dafny.Map.Empty.slice().updateUnsafe(p2,(((_1415_v78).contains(_1414_v77)) ? ((_1415_v78).get(_1414_v77)) : (_dafny.MultiSet.fromElements((_1414_v77).f36))));
          let _1417_v80;
          let _nw216 = new _module.C0();
          _nw216.__ctor(true);
          _1417_v80 = _nw216;
          let _1418_v81;
          _1418_v81 = _dafny.Seq.of(_1417_v80, _1417_v80);
          let _1419_v82;
          _1419_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1416_v79,_dafny.Seq.Concat(_1418_v81, _1418_v81));
          _1419_v82 = (_1419_v82).update(_1416_v79, _1418_v81);
          _1327_v10 = _dafny.Seq.Concat(_1327_v10, _1327_v10);
          let _1420_v83;
          let _nw217 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1420_v83 = _nw217;
          let _1421_v84;
          _1421_v84 = _module.D6.create_DC16(_dafny.Seq.UnicodeFromString("drgpsi"));
          let _1422_v85;
          _1422_v85 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1423_v86;
          _1423_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1421_v84,_1422_v85);
          let _1424_v87;
          _1424_v87 = _dafny.Set.fromElements(_1423_v86);
          let _1425_v89;
          _1425_v89 = _module.D3.create_DC10(_1420_v83);
          let _1426_v90;
          let _nw218 = new _module.C4();
          _nw218.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_1425_v89,(_dafny.ZERO).minus((_1326_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1326_v9).length))])), p0, true);
          _1426_v90 = _nw218;
          let _1427_v91;
          _1427_v91 = _dafny.Seq.of(_1426_v90, _1426_v90, _1426_v90);
          let _1428_v92;
          _1428_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1427_v91).length),_1414_v77.f37);
          let _1429_v93;
          _1429_v93 = _module.D2.create_DC6(function () {
  let _coll61 = new _dafny.Map();
  for (const _compr_61 of (_1428_v92).Keys.Elements) {
    let _1430_v88 = _compr_61;
    if ((_1428_v92).contains(_1430_v88)) {
      _coll61.push([(_1430_v88).multipliedBy(new BigNumber(-800)),p2]);
    }
  }
  return _coll61;
}());
          let _1431_v94;
          _1431_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1327_v10,_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), ((_1432_v77, _1433_v9) => function (_1434_i15) {
            return _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(_1432_v77.f37,new BigNumber((_dafny.Seq.of(_dafny.Seq.of((_1433_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1433_v9).length))]))).length)));
          })(_1414_v77, _1326_v9)), _module.__default.safeIndex(new BigNumber((_1372_v47).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(719)), ((_1435_v77, _1436_v9) => function (_1437_i15) {
            return _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(_1435_v77.f37,new BigNumber((_dafny.Seq.of(_dafny.Seq.of((_1436_v9)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_1436_v9).length))]))).length)));
          })(_1414_v77, _1326_v9))).length)), _1429_v93));
          let _1438_v95;
          _1438_v95 = _module.D10.create_DC23(_1431_v94);
          let _rhs206 = _1414_v77.f37;
          let _rhs207 = _1420_v83;
          let _rhs208 = _1324_v8;
          let _rhs209 = _dafny.Set.fromElements((_module.__default.fm37(_1370_v45, (_1414_v77).fm1(globalState), p2, _1438_v95, globalState)).Merge(_1423_v86));
          let _rhs210 = new BigNumber(569);
          let _lhs190 = globalState;
          let _lhs191 = globalState;
          _lhs190.f16 = _rhs206;
          _1420_v83 = _rhs207;
          _1324_v8 = _rhs208;
          _1424_v87 = _rhs209;
          _lhs191.f16 = _rhs210;
          let _index197 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
          (_1324_v8)[_index197] = ((_1426_v90).fm1(globalState)).isLessThanOrEqualTo((new BigNumber((_dafny.Seq.of((_1324_v8)[_module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length))])).length)).minus((_1414_v77).f36));
          let _1439_v96;
          _1439_v96 = _dafny.Set.fromElements(new BigNumber((_1372_v47).length));
          let _index198 = _module.__default.safeIndex(new BigNumber(477), new BigNumber((_1324_v8).length));
          (_1324_v8)[_index198] = ((new BigNumber((_1439_v96).length)).minus(p2)).isLessThan(p2);
        }
      }
      return;
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _1440_v0;
      let _nw219 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1440_v0 = _nw219;
      let _1441_v1;
      _1441_v1 = new _dafny.CodePoint('h'.codePointAt(0));
      let _index199 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length));
      (_1440_v0)[_index199] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("qsnm"), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.UnicodeFromString("qsnm")).length)), _1441_v1);
      let _1442_v2;
      _1442_v2 = _dafny.Seq.UnicodeFromString("yxuu");
      let _index200 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length));
      (_1440_v0)[_index200] = _1442_v2;
      let _1443_v3;
      _1443_v3 = _dafny.Seq.of((_this).fm2((_this).f26, p3, globalState), p1);
      (globalState).f12 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_this).f26), _1443_v3)).length);
      (globalState).f16 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
      let _1444_v4;
      _1444_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p3, p3)).length),new BigNumber(427));
      let _1445_v5;
      _1445_v5 = _module.D2.create_DC6(_1444_v4);
      let _1446_v7;
      _1446_v7 = _dafny.Set.fromElements(p3);
      let _pat_let_tv24 = _1446_v7;
      let _pat_let_tv25 = _1446_v7;
      let _pat_let_tv26 = p3;
      let _1447_v8;
      let _nw220 = Array((new BigNumber(18)).toNumber());
      _nw220[(_dafny.ZERO).toNumber()] = _1445_v5;
      _nw220[(_dafny.ONE).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(2)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(3)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(4)).toNumber()] = function (_pat_let17_0) {
        return function (_1448_dt__update__tmp_h0) {
          return function (_pat_let18_0) {
            return function (_1450_dt__update_hcf13_h0) {
              return _module.D2.create_DC6(_1450_dt__update_hcf13_h0);
            }(_pat_let18_0);
          }(function () {
            let _coll62 = new _dafny.Map();
            for (const _compr_62 of (_pat_let_tv24).Elements) {
              let _1449_v6 = _compr_62;
              if ((_pat_let_tv25).contains(_1449_v6)) {
                _coll62.push([_module.__default.safeDivisionInt(_1449_v6, new BigNumber(270)),_pat_let_tv26]);
              }
            }
            return _coll62;
          }());
        }(_pat_let17_0);
      }(_1445_v5);
      _nw220[(new BigNumber(5)).toNumber()] = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(p3,p3));
      _nw220[(new BigNumber(6)).toNumber()] = _module.D2.create_DC6(_1444_v4);
      _nw220[(new BigNumber(7)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(8)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(9)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(10)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(11)).toNumber()] = _module.D2.create_DC6(_1444_v4);
      _nw220[(new BigNumber(12)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(13)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(14)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(15)).toNumber()] = _module.D2.create_DC6(_1444_v4);
      _nw220[(new BigNumber(16)).toNumber()] = _1445_v5;
      _nw220[(new BigNumber(17)).toNumber()] = _module.D2.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(p3,p3));
      _1447_v8 = _nw220;
      let _1451_v9;
      _1451_v9 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(999)), function (_1452_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), (_this).fm5(_1443_v3, _1443_v3, p0, !(false), globalState)),_1447_v8);
      _1451_v9 = (_1451_v9).update((_this).f26, _1447_v8);
      let _hi12 = (_dafny.ZERO).minus(p3);
      for (let _1453_i1 = p3; _1453_i1.isLessThan(_hi12); _1453_i1 = _1453_i1.plus(_dafny.ONE)) {
        let _1454_v10;
        let _nw221 = Array((new BigNumber(28)).toNumber()).fill(false);
        _1454_v10 = _nw221;
        let _1455_v11;
        _1455_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1446_v7);
        let _index201 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length));
        (_1454_v10)[_index201] = !((_dafny.Set.fromElements(p3, _1453_i1)).IsSubsetOf((((_1455_v11).contains((_this).f27)) ? ((_1455_v11).get((_this).f27)) : (_1446_v7))));
        let _1456_v12;
        _1456_v12 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
        let _1457_v13;
        _1457_v13 = _dafny.MultiSet.fromElements(_1454_v10);
        let _index202 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length));
        let _rhs211 = !(_module.__default.safeModuloInt(new BigNumber((_1456_v12).length), _1453_i1)).isEqualTo(new BigNumber(297));
        let _rhs212 = ((_1457_v13).Difference(_dafny.MultiSet.fromElements(_1454_v10, _1454_v10))).IsProperSubsetOf(_1457_v13);
        let _lhs192 = globalState;
        let _lhs193 = _1454_v10;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length));
        _lhs192.f22 = _rhs211;
        _lhs193[_lhs194] = _rhs212;
        (globalState).f16 = p3;
        let _1458_v14;
        _1458_v14 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_module.__default.fm25((_1454_v10)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length))], (_1454_v10)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length))], new _dafny.CodePoint('q'.codePointAt(0)), globalState), _1443_v3));
        _1458_v14 = (_1458_v14).Union(_dafny.MultiSet.fromElements(_1443_v3, _1443_v3, _dafny.Seq.of((_1454_v10)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length))], (_this).fm2(p1, p3, globalState)), _dafny.Seq.of((_this).f26, (_1454_v10)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_1454_v10).length))], (_this).f26)));
        (globalState).f16 = _1453_i1;
      }
      let _1459_v15;
      _1459_v15 = _dafny.Seq.of((_this).fm1(globalState));
      let _1460_i2;
      _1460_i2 = _dafny.ZERO;
      L7: {
        while (((!((_1443_v3)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1459_v15, _module.__default.safeIndex(p3, new BigNumber((_1459_v15).length)), new BigNumber(60))).length), new BigNumber((_1443_v3).length))])) ? ((_this).f26) : (false))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1460_i2)) {
              break L7;
            }
            _1460_i2 = (_1460_i2).plus(_dafny.ONE);
            let _1461_v16;
            _1461_v16 = _module.D0.create_DC1((_this).f27, (_this).f26, new BigNumber((_1443_v3).length), p1);
            let _1462_v17;
            _1462_v17 = _module.D0.create_DC2(_1461_v16);
            let _1463_v18;
            _1463_v18 = _module.D0.create_DC2(_1462_v17);
            let _source22 = _module.D0.create_DC2(_1462_v17);
            if (_source22.is_DC1) {
              let _1464___mcc_h0 = (_source22).cf1;
              let _1465___mcc_h1 = (_source22).cf2;
              let _1466___mcc_h2 = (_source22).cf3;
              let _1467___mcc_h3 = (_source22).cf4;
              let _1468_cf4 = _1467___mcc_h3;
              let _1469_cf3 = _1466___mcc_h2;
              let _1470_cf2 = _1465___mcc_h1;
              let _1471_cf1 = _1464___mcc_h0;
              let _1472_v19;
              let _nw222 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1472_v19 = _nw222;
              let _index203 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_1472_v19).length));
              (_1472_v19)[_index203] = _1441_v1;
              let _index204 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_1472_v19).length));
              (_1472_v19)[_index204] = _1441_v1;
              let _1473_v20;
              _1473_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1469_cf3,!(_1470_cf2) || (true));
              _1473_v20 = (_1473_v20).update(_1469_cf3, (_1469_cf3).isLessThanOrEqualTo(p3));
              (globalState).f24 = (((_this).fm2(p1, p3, globalState)) ? (_1468_cf4) : (true));
              let _1474_v21;
              let _init40 = ((_1475_cf2) => function (_1476_i3) {
                return _1475_cf2;
              })(_1470_cf2);
              let _nw223 = Array((new BigNumber(28)).toNumber());
              for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw223.length); _i0_40++) {
                _nw223[_i0_40] = _init40(new BigNumber(_i0_40));
              }
              _1474_v21 = _nw223;
              let _index205 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1474_v21).length));
              (_1474_v21)[_index205] = _1471_cf1;
              let _index206 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1474_v21).length));
              (_1474_v21)[_index206] = !(_1471_cf1);
            } else if (_source22.is_DC0) {
              let _1477___mcc_h4 = (_source22).cf0;
              let _1478_cf0 = _1477___mcc_h4;
              let _1479_v22;
              let _init41 = ((_1480_p3) => function (_1481_i4) {
                return _module.__default.safeModuloInt(_1481_i4, _1480_p3);
              })(p3);
              let _nw224 = Array((new BigNumber(6)).toNumber());
              for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw224.length); _i0_41++) {
                _nw224[_i0_41] = _init41(new BigNumber(_i0_41));
              }
              _1479_v22 = _nw224;
              let _1482_v23;
              _1482_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_1479_v22);
              let _1483_v24;
              let _nw225 = Array((new BigNumber(4)).toNumber());
              _nw225[(_dafny.ZERO).toNumber()] = _1479_v22;
              _nw225[(_dafny.ONE).toNumber()] = _1479_v22;
              _nw225[(new BigNumber(2)).toNumber()] = (((_1482_v23).contains(!(true))) ? ((_1482_v23).get(!(true))) : (_1479_v22));
              _nw225[(new BigNumber(3)).toNumber()] = _1479_v22;
              _1483_v24 = _nw225;
              let _index207 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_1483_v24).length));
              (_1483_v24)[_index207] = _1479_v22;
              let _index208 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_1483_v24).length));
              (_1483_v24)[_index208] = _1479_v22;
              let _1484_v25;
              _1484_v25 = _dafny.MultiSet.fromElements(p3, p3, p3, p3, p3);
              let _1485_v26;
              let _nw226 = Array((new BigNumber(11)).toNumber());
              _nw226[(_dafny.ZERO).toNumber()] = (p3).isLessThan(new BigNumber(728));
              _nw226[(_dafny.ONE).toNumber()] = _1478_cf0;
              _nw226[(new BigNumber(2)).toNumber()] = _1478_cf0;
              _nw226[(new BigNumber(3)).toNumber()] = _1478_cf0;
              _nw226[(new BigNumber(4)).toNumber()] = !(!_dafny.areEqual((_1440_v0)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length))], _1442_v2));
              _nw226[(new BigNumber(5)).toNumber()] = !((_1484_v25).IsSubsetOf(_1484_v25));
              _nw226[(new BigNumber(6)).toNumber()] = (_this).fm2(false, p3, globalState);
              _nw226[(new BigNumber(7)).toNumber()] = (_this).f26;
              _nw226[(new BigNumber(8)).toNumber()] = (_this).f27;
              _nw226[(new BigNumber(9)).toNumber()] = p1;
              _nw226[(new BigNumber(10)).toNumber()] = true;
              _1485_v26 = _nw226;
              let _index209 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1485_v26).length));
              (_1485_v26)[_index209] = p1;
              let _1486_v27;
              _1486_v27 = _module.D6.create_DC16(_dafny.Seq.UnicodeFromString("oqncb"));
              let _pat_let_tv27 = _1442_v2;
              let _pat_let_tv28 = _1442_v2;
              let _1487_v28;
              let _nw227 = Array((new BigNumber(28)).toNumber());
              _nw227[(_dafny.ZERO).toNumber()] = _module.D6.create_DC16(_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), function (_1488_i5) {
  return new _dafny.CodePoint('b'.codePointAt(0));
}));
              _nw227[(_dafny.ONE).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(2)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(3)).toNumber()] = _module.__default.fm38(true, p3, (_1440_v0)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length))], globalState);
              _nw227[(new BigNumber(4)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(5)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(6)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(7)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(8)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(9)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(10)).toNumber()] = function (_pat_let19_0) {
                return function (_1489_dt__update__tmp_h1) {
                  return function (_pat_let20_0) {
                    return function (_1490_dt__update_hcf29_h0) {
                      return _module.D6.create_DC16(_1490_dt__update_hcf29_h0);
                    }(_pat_let20_0);
                  }(_pat_let_tv27);
                }(_pat_let19_0);
              }(_module.D6.create_DC16(_1442_v2));
              _nw227[(new BigNumber(11)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(12)).toNumber()] = _module.D6.create_DC16(_1442_v2);
              _nw227[(new BigNumber(13)).toNumber()] = function (_pat_let21_0) {
                return function (_1491_dt__update__tmp_h2) {
                  return function (_pat_let22_0) {
                    return function (_1492_dt__update_hcf29_h1) {
                      return _module.D6.create_DC16(_1492_dt__update_hcf29_h1);
                    }(_pat_let22_0);
                  }(_pat_let_tv28);
                }(_pat_let21_0);
              }(_1486_v27);
              _nw227[(new BigNumber(14)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(15)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(16)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(17)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(18)).toNumber()] = _module.D6.create_DC16(_1442_v2);
              _nw227[(new BigNumber(19)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(20)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(21)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(22)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(23)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(24)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(25)).toNumber()] = _module.D6.create_DC16((_1440_v0)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length))]);
              _nw227[(new BigNumber(26)).toNumber()] = _1486_v27;
              _nw227[(new BigNumber(27)).toNumber()] = _1486_v27;
              _1487_v28 = _nw227;
              let _index210 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1487_v28).length));
              (_1487_v28)[_index210] = _1486_v27;
              let _index211 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1485_v26).length));
              let _index212 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1487_v28).length));
              let _rhs213 = !((_this).f27);
              let _rhs214 = _1486_v27;
              let _lhs195 = _1485_v26;
              let _lhs196 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1485_v26).length));
              let _lhs197 = _1487_v28;
              let _lhs198 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1487_v28).length));
              _lhs195[_lhs196] = _rhs213;
              _lhs197[_lhs198] = _rhs214;
              let _1493_v29;
              _1493_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(p3),_1441_v1);
              let _1494_v30;
              _1494_v30 = _dafny.Seq.of(_1493_v29);
              let _1495_v31;
              _1495_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1494_v30, _dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_1496_p3, _1497_v1) => function (_1498_i6) {
                return _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(_1496_p3),_1497_v1);
              })(p3, _1441_v1))),_1443_v3);
              _1495_v31 = (_1495_v31).update(_dafny.Seq.Concat(_dafny.Seq.of(function () {
                let _coll63 = new _dafny.Map();
                for (const _compr_63 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(p3),p3)).Keys.Elements) {
                  let _1499_v32 = _compr_63;
                  if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC5(p3),p3)).contains(_1499_v32)) {
                    _coll63.push([_1499_v32,_1441_v1]);
                  }
                }
                return _coll63;
              }()), _1494_v30), _1443_v3);
              let _pat_let_tv29 = _1440_v0;
              let _pat_let_tv30 = _1440_v0;
              (globalState).f6 = (function (_pat_let23_0) {
                return function (_1500_dt__update__tmp_h3) {
                  return function (_pat_let24_0) {
                    return function (_1501_dt__update_hcf29_h2) {
                      return _module.D6.create_DC16(_1501_dt__update_hcf29_h2);
                    }(_pat_let24_0);
                  }((_pat_let_tv30)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_pat_let_tv29).length))]);
                }(_pat_let23_0);
              }(_1486_v27)).dtor_cf29;
            } else {
              let _1502___mcc_h5 = (_source22).cf5;
              let _1503_cf5 = _1502___mcc_h5;
              let _1504_v33;
              let _init42 = function (_1505_i7) {
                return (_this).f26;
              };
              let _nw228 = Array((new BigNumber(6)).toNumber());
              for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw228.length); _i0_42++) {
                _nw228[_i0_42] = _init42(new BigNumber(_i0_42));
              }
              _1504_v33 = _nw228;
              let _1506_v34;
              _1506_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v33,p3);
              _1506_v34 = (_1506_v34).update(_1504_v33, p3);
              let _1507_v35;
              _1507_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(((_1440_v0)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length))]).length));
              let _1508_v36;
              _1508_v36 = _dafny.MultiSet.fromElements(_1507_v35, _1507_v35);
              let _rhs215 = (_1440_v0)[_module.__default.safeIndex(new BigNumber(371), new BigNumber((_1440_v0).length))];
              let _rhs216 = p1;
              let _rhs217 = (p3).isLessThan(new BigNumber((_1508_v36).cardinality()));
              let _rhs218 = (_this).f27;
              let _rhs219 = _1441_v1;
              let _lhs199 = globalState;
              let _lhs200 = globalState;
              let _lhs201 = globalState;
              let _lhs202 = globalState;
              let _lhs203 = globalState;
              _lhs199.f6 = _rhs215;
              _lhs200.f24 = _rhs216;
              _lhs201.f15 = _rhs217;
              _lhs202.f24 = _rhs218;
              _lhs203.f19 = _rhs219;
              (globalState).f2 = p1;
              (globalState).f16 = p3;
            }
            let _1509_v37;
            let _1510_v38;
            let _1511_v39;
            let _1512_v40;
            let _out45;
            let _out46;
            let _out47;
            let _out48;
            let _outcollector12 = _module.__default.m0(_1441_v1, globalState);
            _out45 = _outcollector12[0];
            _out46 = _outcollector12[1];
            _out47 = _outcollector12[2];
            _out48 = _outcollector12[3];
            _1509_v37 = _out45;
            _1510_v38 = _out46;
            _1511_v39 = _out47;
            _1512_v40 = _out48;
            let _1513_v41;
            let _nw229 = Array((new BigNumber(10)).toNumber()).fill([]);
            _1513_v41 = _nw229;
            let _1514_v42;
            let _nw230 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
            _1514_v42 = _nw230;
            let _index213 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1513_v41).length));
            (_1513_v41)[_index213] = _1514_v42;
            let _1515_v43;
            let _nw231 = new _module.C1();
            _nw231.__ctor(p3, p3, true, false);
            _1515_v43 = _nw231;
            let _1516_v44;
            _1516_v44 = _dafny.Seq.of(_1515_v43);
            let _index214 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1513_v41).length));
            let _rhs220 = ((((_this).f27) || (p1)) ? (_1514_v42) : (_1514_v42));
            let _rhs221 = _1516_v44;
            let _rhs222 = (_dafny.ZERO).minus((p3).plus(p3));
            let _rhs223 = (_1515_v43).f26;
            let _lhs204 = _1513_v41;
            let _lhs205 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_1513_v41).length));
            let _lhs206 = globalState;
            let _lhs207 = globalState;
            _lhs204[_lhs205] = _rhs220;
            _1516_v44 = _rhs221;
            _lhs206.f12 = _rhs222;
            _lhs207.f24 = _rhs223;
            let _1517_v45;
            _1517_v45 = _module.D3.create_DC10((_1513_v41)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_1513_v41).length))]);
            let _1518_v46;
            _1518_v46 = _dafny.Map.Empty.slice().updateUnsafe((_1515_v43).f26,(_this).f26);
            let _1519_v47;
            _1519_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1517_v45,new BigNumber((_1518_v46).length));
            let _1520_v48;
            let _nw232 = new _module.C4();
            _nw232.__ctor(_1519_v47, (_1515_v43).f26, (_this).f26);
            _1520_v48 = _nw232;
            let _1521_v49;
            _1521_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1520_v48,!(false));
            _1521_v49 = (_1521_v49).update(_1520_v48, true);
          }
        }
      }
      r0 = !((_1443_v3)[_module.__default.safeIndex(p3, new BigNumber((_1443_v3).length))]) || ((_this).f27);
      return r0;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f40 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f40) {
      let _this = this;
      (_this).f40 = f40;
      return;
    }
    fm1(globalState) {
      let _this = this;
      return ((_this.f40).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_1522_i0) {
        return _this.f40;
      })).length))).minus(_this.f40);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return ((function () {
        let _coll64 = new _dafny.Set();
        for (const _compr_64 of (_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f40, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length), (_dafny.ZERO).minus(_this.f40)))).Elements) {
          let _1523_v0 = _compr_64;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f40, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(false))).length), (_dafny.ZERO).minus(_this.f40)))).contains(_1523_v0)) {
            _coll64.add(_module.__default.safeModuloInt(_1523_v0, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), false, false, !(false), true))).cardinality())));
          }
        }
        return _coll64;
      }()).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(_this.f40), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_this.f40)).cardinality()))).cardinality()), new BigNumber((_dafny.MultiSet.fromElements(true, !(true))).cardinality()), _this.f40))).equals((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(400)), function (_1524_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), (_dafny.ZERO).minus(_this.f40), _this.f40)).Difference(_dafny.Set.fromElements(_this.f40, new BigNumber((_dafny.Seq.of(function () {
        let _coll65 = new _dafny.Set();
        for (const _compr_65 of _dafny.IntegerRange(new BigNumber(178), new BigNumber(967))) {
          let _1525_v1 = _compr_65;
          if (((new BigNumber(178)).isLessThanOrEqualTo(_1525_v1)) && ((_1525_v1).isLessThan(new BigNumber(967)))) {
            _coll65.add((_1525_v1).multipliedBy(_this.f40));
          }
        }
        return _coll65;
      }())).length))));
    };
    fm44(globalState) {
      let _this = this;
      return (!(!(false))) && (!((!(false)) === (true)));
    };
    m15(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      (globalState).f12 = (_this).fm1(globalState);
      let _1526_v0;
      _1526_v0 = _dafny.MultiSet.fromElements(p2);
      let _1527_v1;
      _1527_v1 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1526_v0).cardinality())).plus(new BigNumber(866)),p2);
      let _1528_v2;
      _1528_v2 = _module.D16.create_DC37(_1527_v1);
      _1527_v1 = (_1527_v1).Merge((_1528_v2).dtor_cf68);
      (globalState).f15 = (p3) && (p2);
      let _1529_v3;
      _1529_v3 = _dafny.Seq.UnicodeFromString("ymvn");
      let _1530_v4;
      _1530_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1529_v3,p2);
      (globalState).f12 = (_dafny.ZERO).minus((((_1526_v0).contains(!(_dafny.Map.Empty.slice().updateUnsafe(_1529_v3,p2)).equals(_1530_v4))) ? ((_1526_v0).get(!(_dafny.Map.Empty.slice().updateUnsafe(_1529_v3,p2)).equals(_1530_v4))) : (p0)));
      let _1531_i0;
      _1531_i0 = _dafny.ZERO;
      L8: {
        while (p2) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1531_i0)) {
              break L8;
            }
            _1531_i0 = (_1531_i0).plus(_dafny.ONE);
            let _1532_v5;
            _1532_v5 = _module.D15.create_DC35(p2);
            _1532_v5 = _1532_v5;
            let _1533_v6;
            let _init43 = function (_1534_i1) {
              return _module.__default.safeDivisionInt(_1534_i1, new BigNumber(669));
            };
            let _nw233 = Array((new BigNumber(25)).toNumber());
            for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw233.length); _i0_43++) {
              _nw233[_i0_43] = _init43(new BigNumber(_i0_43));
            }
            _1533_v6 = _nw233;
            let _index215 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1533_v6).length));
            (_1533_v6)[_index215] = _this.f40;
            let _index216 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1533_v6).length));
            (_1533_v6)[_index216] = new BigNumber(134);
            let _1535_v7;
            let _nw234 = new _module.C1();
            _nw234.__ctor(p0, _this.f40, p2, p2);
            _1535_v7 = _nw234;
            let _1536_v8;
            _1536_v8 = _module.D7.create_DC18(_1535_v7);
            let _pat_let_tv31 = _1535_v7;
            let _pat_let_tv32 = _1535_v7;
            let _1537_v9;
            let _nw235 = Array((new BigNumber(18)).toNumber());
            _nw235[(_dafny.ZERO).toNumber()] = _1536_v8;
            _nw235[(_dafny.ONE).toNumber()] = function (_pat_let25_0) {
              return function (_1538_dt__update__tmp_h0) {
                return function (_pat_let26_0) {
                  return function (_1539_dt__update_hcf35_h0) {
                    return _module.D7.create_DC18(_1539_dt__update_hcf35_h0);
                  }(_pat_let26_0);
                }(_pat_let_tv31);
              }(_pat_let25_0);
            }(_1536_v8);
            _nw235[(new BigNumber(2)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(3)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(4)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(5)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(6)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(7)).toNumber()] = function (_pat_let27_0) {
              return function (_1540_dt__update__tmp_h1) {
                return function (_pat_let28_0) {
                  return function (_1541_dt__update_hcf35_h1) {
                    return _module.D7.create_DC18(_1541_dt__update_hcf35_h1);
                  }(_pat_let28_0);
                }(_pat_let_tv32);
              }(_pat_let27_0);
            }(_module.D7.create_DC18(_1535_v7));
            _nw235[(new BigNumber(8)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(9)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(10)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(11)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(12)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(13)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(14)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(15)).toNumber()] = _module.D7.create_DC18(_1535_v7);
            _nw235[(new BigNumber(16)).toNumber()] = _1536_v8;
            _nw235[(new BigNumber(17)).toNumber()] = _1536_v8;
            _1537_v9 = _nw235;
            let _index217 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1537_v9).length));
            (_1537_v9)[_index217] = _1536_v8;
            let _index218 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_1537_v9).length));
            (_1537_v9)[_index218] = _1536_v8;
            let _1542_v10;
            let _nw236 = Array((new BigNumber(7)).toNumber()).fill(false);
            _1542_v10 = _nw236;
            let _index219 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1542_v10).length));
            (_1542_v10)[_index219] = (_1535_v7).f26;
            let _index220 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_1542_v10).length));
            (_1542_v10)[_index220] = p2;
          }
        }
      }
      let _1543_i2;
      _1543_i2 = _dafny.ZERO;
      L9: {
        while (p2) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1543_i2)) {
              break L9;
            }
            _1543_i2 = (_1543_i2).plus(_dafny.ONE);
            let _1544_v11;
            _1544_v11 = _dafny.Seq.of(p3);
            let _1545_v12;
            _1545_v12 = _dafny.Set.fromElements(_this.f40, _this.f40, p0, p0, p0);
            let _1546_v13;
            let _nw237 = Array((new BigNumber(8)).toNumber());
            _nw237[(_dafny.ZERO).toNumber()] = p3;
            _nw237[(_dafny.ONE).toNumber()] = (_1544_v11)[_module.__default.safeIndex(new BigNumber((_1545_v12).length), new BigNumber((_1544_v11).length))];
            _nw237[(new BigNumber(2)).toNumber()] = true;
            _nw237[(new BigNumber(3)).toNumber()] = (_this).fm2(p2, p0, globalState);
            _nw237[(new BigNumber(4)).toNumber()] = true;
            _nw237[(new BigNumber(5)).toNumber()] = p3;
            _nw237[(new BigNumber(6)).toNumber()] = p2;
            _nw237[(new BigNumber(7)).toNumber()] = p2;
            _1546_v13 = _nw237;
            let _index221 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_1546_v13).length));
            (_1546_v13)[_index221] = (_this).fm44(globalState);
            let _index222 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_1546_v13).length));
            (_1546_v13)[_index222] = (((_this).fm1(globalState)).plus((_dafny.ZERO).minus(p0))).isLessThan(_this.f40);
            r2 = p0;
            let _rhs224 = (_this).fm1(globalState);
            let _rhs225 = (_this).fm1(globalState);
            let _rhs226 = new BigNumber((_dafny.Seq.of(new BigNumber(-197), (_dafny.ZERO).minus((_this).fm1(globalState)), new BigNumber((_1545_v12).length), ((_dafny.ZERO).minus(_this.f40)).plus(_this.f40), new BigNumber(452))).length);
            let _rhs227 = _1529_v3;
            let _lhs208 = globalState;
            let _lhs209 = globalState;
            let _lhs210 = globalState;
            let _lhs211 = globalState;
            _lhs208.f16 = _rhs224;
            _lhs209.f12 = _rhs225;
            _lhs210.f12 = _rhs226;
            _lhs211.f6 = _rhs227;
            let _1547_v14;
            _1547_v14 = _dafny.Seq.of(new BigNumber((_1529_v3).length));
            let _1548_v15;
            _1548_v15 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0, p0), _1547_v14, _1547_v14);
            let _1549_v16;
            let _nw238 = new _module.C3();
            _nw238.__ctor(_1548_v15, (((_1546_v13)[_module.__default.safeIndex(new BigNumber(147), new BigNumber((_1546_v13).length))]) ? (p3) : (true)), (_1546_v13)[_module.__default.safeIndex(new BigNumber(147), new BigNumber((_1546_v13).length))]);
            _1549_v16 = _nw238;
          }
        }
      }
      r0 = _module.D1.create_DC5(_module.__default.safeModuloInt(p0, p0));
      let _1550_v17;
      _1550_v17 = _dafny.Seq.of(p3);
      let _1551_v18;
      _1551_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1526_v0,(_1550_v17)[_module.__default.safeIndex(_this.f40, new BigNumber((_1550_v17).length))]);
      r1 = _1551_v18;
      r2 = (_this.f40).multipliedBy(_this.f40);
      return [r0, r1, r2];
    }
    m16(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _1552_v0;
      _1552_v0 = _dafny.Seq.of(_this.f40);
      let _1553_v1;
      _1553_v1 = _dafny.Seq.of(_1552_v0);
      let _1554_v2;
      _1554_v2 = false;
      let _rhs228 = p0;
      let _rhs229 = _dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm45(_1554_v2, p0, new BigNumber(-164), _1554_v2, globalState), _1552_v0, _1552_v0), _1553_v1);
      let _lhs212 = globalState;
      _lhs212.f16 = _rhs228;
      _1553_v1 = _rhs229;
      let _rhs230 = p0;
      let _rhs231 = _this.f40;
      let _lhs213 = globalState;
      let _lhs214 = globalState;
      _lhs213.f16 = _rhs230;
      _lhs214.f12 = _rhs231;
      let _1555_v3;
      let _nw239 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1555_v3 = _nw239;
      let _index223 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length));
      (_1555_v3)[_index223] = p0;
      let _1556_v4;
      _1556_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f40,_module.__default.fm47(_1554_v2, _1554_v2, _1554_v2, !(_1554_v2), globalState));
      let _1557_v5;
      _1557_v5 = _dafny.Seq.UnicodeFromString("mswelsh");
      let _1558_v6;
      _1558_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1557_v5).length),_1554_v2);
      let _1559_v7;
      _1559_v7 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
      let _1560_v8;
      _1560_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1559_v7,_this.f40);
      let _1561_v9;
      _1561_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1554_v2,_this.f40),new BigNumber((_1560_v8).length));
      let _1562_v10;
      let _nw240 = Array((new BigNumber(4)).toNumber());
      _nw240[(_dafny.ZERO).toNumber()] = _module.__default.fm46(_1554_v2, p0, p0, _1554_v2, globalState);
      _nw240[(_dafny.ONE).toNumber()] = _1556_v4;
      _nw240[(new BigNumber(2)).toNumber()] = (_1556_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1558_v6));
      _nw240[(new BigNumber(3)).toNumber()] = _module.__default.fm46(_1554_v2, (((_1561_v9).contains(_dafny.Map.Empty.slice().updateUnsafe(_1554_v2,new BigNumber((_dafny.Seq.UnicodeFromString("yxbdjm")).length)))) ? ((_1561_v9).get(_dafny.Map.Empty.slice().updateUnsafe(_1554_v2,new BigNumber((_dafny.Seq.UnicodeFromString("yxbdjm")).length)))) : (_this.f40)), p0, _1554_v2, globalState);
      _1562_v10 = _nw240;
      let _index224 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1562_v10).length));
      (_1562_v10)[_index224] = _1556_v4;
      let _1563_v11;
      let _nw241 = Array((new BigNumber(27)).toNumber()).fill(false);
      _1563_v11 = _nw241;
      let _1564_v12;
      _1564_v12 = _dafny.Seq.of(_1563_v11, _1563_v11);
      let _1565_v13;
      _1565_v13 = _dafny.Seq.of(_1564_v12);
      let _1566_v14;
      let _nw242 = Array((new BigNumber(25)).toNumber());
      _nw242[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1564_v12, _1564_v12);
      _nw242[(_dafny.ONE).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(2)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(3)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(4)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(5)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1564_v12, (_1565_v13)[_module.__default.safeIndex(new BigNumber(490), new BigNumber((_1565_v13).length))]);
      _nw242[(new BigNumber(7)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_1563_v11, _1563_v11);
      _nw242[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_1563_v11), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ajrpp")).length))).length), new BigNumber((_dafny.Seq.of(_1563_v11)).length)), _1563_v11), _1564_v12);
      _nw242[(new BigNumber(10)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(11)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(12)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(13)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(14)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_1563_v11, _1563_v11, _1563_v11);
      _nw242[(new BigNumber(16)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_1564_v12, _module.__default.safeIndex(p0, new BigNumber((_1564_v12).length)), _1563_v11);
      _nw242[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1564_v12, _dafny.Seq.update(_1564_v12, _module.__default.safeIndex(_this.f40, new BigNumber((_1564_v12).length)), _1563_v11));
      _nw242[(new BigNumber(19)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(20)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(21)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1564_v12, _1564_v12);
      _nw242[(new BigNumber(23)).toNumber()] = _1564_v12;
      _nw242[(new BigNumber(24)).toNumber()] = _1564_v12;
      _1566_v14 = _nw242;
      let _index225 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1566_v14).length));
      (_1566_v14)[_index225] = _1564_v12;
      let _1567_v15;
      _1567_v15 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1568_v16;
      _1568_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p0),_1567_v15);
      let _1569_v17;
      _1569_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_1568_v16).contains(_1559_v7)) ? ((_1568_v16).get(_1559_v7)) : ((_1557_v5)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1557_v5).length))])),_1554_v2);
      let _index226 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length));
      let _index227 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1562_v10).length));
      let _index228 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1566_v14).length));
      let _rhs232 = p0;
      let _rhs233 = ((_1554_v2) ? (_1554_v2) : (((!(false)) ? (_1554_v2) : (_1554_v2))));
      let _rhs234 = _1556_v4;
      let _rhs235 = _dafny.Seq.update(_1564_v12, _module.__default.safeIndex(p0, new BigNumber((_1564_v12).length)), _1563_v11);
      let _rhs236 = _1569_v17;
      let _lhs215 = _1555_v3;
      let _lhs216 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length));
      let _lhs217 = globalState;
      let _lhs218 = _1562_v10;
      let _lhs219 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1562_v10).length));
      let _lhs220 = _1566_v14;
      let _lhs221 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1566_v14).length));
      _lhs215[_lhs216] = _rhs232;
      _lhs217.f24 = _rhs233;
      _lhs218[_lhs219] = _rhs234;
      _lhs220[_lhs221] = _rhs235;
      _1569_v17 = _rhs236;
      (globalState).f24 = !(!(_1554_v2));
      let _1570_v18;
      _1570_v18 = _module.D6.create_DC16(_dafny.Seq.Create(_module.__default.abs(new BigNumber(977)), ((_1571_v15) => function (_1572_i0) {
  return _1571_v15;
})(_1567_v15)));
      let _pat_let_tv33 = _1554_v2;
      let _pat_let_tv34 = _1554_v2;
      if (function (_source23) {
        if (_source23.is_DC17) {
          let _1573___mcc_h0 = (_source23).cf30;
          let _1574___mcc_h1 = (_source23).cf31;
          let _1575___mcc_h2 = (_source23).cf32;
          let _1576___mcc_h3 = (_source23).cf33;
          let _1577___mcc_h4 = (_source23).cf34;
          let _1578_cf34 = _1577___mcc_h4;
          let _1579_cf33 = _1576___mcc_h3;
          let _1580_cf32 = _1575___mcc_h2;
          let _1581_cf31 = _1574___mcc_h1;
          let _1582_cf30 = _1573___mcc_h0;
          return _pat_let_tv33;
        } else {
          let _1583___mcc_h5 = (_source23).cf29;
          let _1584_cf29 = _1583___mcc_h5;
          return _pat_let_tv34;
        }
      }(((_1554_v2) ? (_1570_v18) : (_1570_v18)))) {
        let _1585_v19;
        _1585_v19 = _dafny.Seq.of(_1554_v2);
        let _index229 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length));
        (_1555_v3)[_index229] = new BigNumber((_dafny.Seq.Concat(_1585_v19, _dafny.Seq.Concat(_dafny.Seq.update(_1585_v19, _module.__default.safeIndex(_this.f40, new BigNumber((_1585_v19).length)), _1554_v2), _1585_v19))).length);
        (globalState).f24 = _1554_v2;
        _1555_v3 = _1555_v3;
        _1555_v3 = _1555_v3;
        let _1586_v20;
        _1586_v20 = _module.D1.create_DC5((_dafny.ZERO).minus(_this.f40));
        let _1587_v21;
        _1587_v21 = _dafny.Set.fromElements(_1586_v20);
        let _1588_v22;
        _1588_v22 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("yeue"), _1557_v5)).length));
        let _rhs237 = (_dafny.Set.fromElements(_1586_v20, _1586_v20, _1586_v20)).Difference(_1587_v21);
        let _rhs238 = _1588_v22;
        let _rhs239 = _1563_v11;
        _1587_v21 = _rhs237;
        _1588_v22 = _rhs238;
        _1563_v11 = _rhs239;
      } else {
        let _1589_v23;
        _1589_v23 = _dafny.Seq.of(_1557_v5);
        let _1590_v24;
        let _nw243 = Array((new BigNumber(2)).toNumber());
        _nw243[(_dafny.ZERO).toNumber()] = _1567_v15;
        _nw243[(_dafny.ONE).toNumber()] = _1567_v15;
        _1590_v24 = _nw243;
        let _1591_v25;
        _1591_v25 = _module.D1.create_DC4(_1590_v24, p0, p0, _1557_v5, _dafny.Seq.UnicodeFromString("cjntvhbdw"));
        let _1592_v26;
        let _nw244 = Array((new BigNumber(20)).toNumber());
        _nw244[(_dafny.ZERO).toNumber()] = _1557_v5;
        _nw244[(_dafny.ONE).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(2)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(3)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(4)).toNumber()] = (_1589_v23)[_module.__default.safeIndex(_this.f40, new BigNumber((_1589_v23).length))];
        _nw244[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("p");
        _nw244[(new BigNumber(6)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(7)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(8)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(9)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("agk");
        _nw244[(new BigNumber(11)).toNumber()] = (_1591_v25).dtor_cf10;
        _nw244[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1557_v5, _1557_v5);
        _nw244[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(566)), ((_1593_v15) => function (_1594_i1) {
          return _1593_v15;
        })(_1567_v15));
        _nw244[(new BigNumber(14)).toNumber()] = _1557_v5;
        _nw244[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("lwu");
        _nw244[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_1557_v5, _1557_v5);
        _nw244[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("gdid");
        _nw244[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1557_v5, _1557_v5);
        _nw244[(new BigNumber(19)).toNumber()] = _1557_v5;
        _1592_v26 = _nw244;
        let _index230 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1592_v26).length));
        (_1592_v26)[_index230] = _1557_v5;
        let _1595_v27;
        _1595_v27 = _dafny.Set.fromElements(_1554_v2);
        let _1596_v28;
        _1596_v28 = _dafny.MultiSet.fromElements(true);
        let _index231 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_1592_v26).length));
        (_1592_v26)[_index231] = _dafny.Seq.Concat(_1557_v5, _dafny.Seq.update(((true) ? (_1557_v5) : (_dafny.Seq.update(_1557_v5, _module.__default.safeIndex(new BigNumber((_1595_v27).length), new BigNumber((_1557_v5).length)), _1567_v15))), _module.__default.safeIndex((((_1596_v28).contains(false)) ? ((_1596_v28).get(false)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1559_v7).cardinality()),_1554_v2)).length))), new BigNumber((((true) ? (_1557_v5) : (_dafny.Seq.update(_1557_v5, _module.__default.safeIndex(new BigNumber((_1595_v27).length), new BigNumber((_1557_v5).length)), _1567_v15)))).length)), _1567_v15));
        let _1597_v29;
        _1597_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v2,_1554_v2);
        (globalState).f22 = !((_1554_v2) || ((_this.f40).isLessThanOrEqualTo(new BigNumber((_1597_v29).length))));
        let _1598_v30;
        _1598_v30 = _module.D5.create_DC14(p0, _this.f40, _1554_v2);
        let _1599_v31;
        let _nw245 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _1599_v31 = _nw245;
        let _1600_v32;
        _1600_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1599_v31,((_1554_v2) ? (_1555_v3) : (_1555_v3)));
        let _1601_v33;
        _1601_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v2,(_dafny.ZERO).minus(_this.f40));
        let _1602_v34;
        _1602_v34 = _module.D7.create_DC19((_dafny.ZERO).minus((_1555_v3)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length))]), _1596_v28, _1554_v2, _dafny.Map.Empty.slice().updateUnsafe(_1554_v2,(_1555_v3)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length))]));
        let _1603_v35;
        _1603_v35 = _module.D7.create_DC19(_this.f40, _1596_v28, _1554_v2, (_1602_v34).dtor_cf39);
        let _pat_let_tv35 = globalState;
        let _rhs240 = (((_1600_v32).contains(_1599_v31)) ? ((_1600_v32).get(_1599_v31)) : (_1555_v3));
        let _rhs241 = ((((_1601_v33).contains(_1554_v2)) ? ((_1601_v33).get(_1554_v2)) : (new BigNumber((_1601_v33).length)))).isLessThanOrEqualTo((_1603_v35).dtor_cf36);
        let _rhs242 = function (_pat_let29_0) {
          return function (_1604_dt__update__tmp_h0) {
            return function (_pat_let30_0) {
              return function (_1605_dt__update_hcf26_h0) {
                return function (_pat_let31_0) {
                  return function (_1606_dt__update_hcf25_h0) {
                    return _module.D5.create_DC14(_1606_dt__update_hcf25_h0, _1605_dt__update_hcf26_h0, (_1604_dt__update__tmp_h0).dtor_cf27);
                  }(_pat_let31_0);
                }(new BigNumber(176));
              }(_pat_let30_0);
            }((_this).fm1(_pat_let_tv35));
          }(_pat_let29_0);
        }(_1598_v30);
        let _lhs222 = globalState;
        _1555_v3 = _rhs240;
        _lhs222.f24 = _rhs241;
        _1598_v30 = _rhs242;
        _1564_v12 = _dafny.Seq.Concat(_1564_v12, (_1566_v14)[_module.__default.safeIndex(new BigNumber(119), new BigNumber((_1566_v14).length))]);
        let _1607_v36;
        let _nw246 = Array((new BigNumber(2)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _1555_v3;
        _nw246[(_dafny.ONE).toNumber()] = _1555_v3;
        _1607_v36 = _nw246;
        let _index232 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1607_v36).length));
        (_1607_v36)[_index232] = _1555_v3;
        let _index233 = _module.__default.safeIndex(new BigNumber(542), new BigNumber((_1607_v36).length));
        (_1607_v36)[_index233] = ((true) ? (_1555_v3) : (_1555_v3));
      }
      let _1608_v37;
      _1608_v37 = _dafny.Set.fromElements((_1555_v3)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length))], new BigNumber((_1557_v5).length), p0, (_1555_v3)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1555_v3).length))]);
      let _1609_v38;
      _1609_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1554_v2,_1608_v37);
      _1609_v38 = (_1609_v38).update(_dafny.Seq.IsPrefixOf(_1557_v5, _dafny.Seq.UnicodeFromString("krmbocy")), _1608_v37);
      let _1610_v39;
      _1610_v39 = _dafny.Seq.of(_1554_v2, _1554_v2, _1554_v2, _1554_v2, _1554_v2);
      r0 = _dafny.Seq.Concat(_1610_v39, _dafny.Seq.of(_1554_v2, _1554_v2, false, _1554_v2));
      return r0;
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this.f38 = false;
      this.f39 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f38, f39) {
      let _this = this;
      (_this).f38 = f38;
      (_this).f39 = f39;
      return;
    }
    fm1(globalState) {
      let _this = this;
      return new BigNumber((function (_source24) {
        let _1611___mcc_h0 = _source24;
        let _1612_cf40 = _1611___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_this.f38, false, _this.f38, _this.f38, true)).length),new BigNumber(428))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(818),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_module.D2.create_DC9(new BigNumber(795), new BigNumber(-873), _this.f38)).dtor_cf20)).length)));
      }(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-746))).length),_this.f38)).length)))).length);
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return (function () {
        let _coll66 = new _dafny.Set();
        for (const _compr_66 of (_dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-201)), new BigNumber((_dafny.Seq.UnicodeFromString("imnioove")).length), new BigNumber(112)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(373))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_this.f38, _this.f38)).cardinality()))), _dafny.MultiSet.fromElements(new BigNumber(987)))).Elements) {
          let _1613_v0 = _compr_66;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-201)), new BigNumber((_dafny.Seq.UnicodeFromString("imnioove")).length), new BigNumber(112)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(373))), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_this.f38, _this.f38)).cardinality()))), _dafny.MultiSet.fromElements(new BigNumber(987))), _1613_v0)) {
            _coll66.add(_1613_v0);
          }
        }
        return _coll66;
      }()).IsProperSubsetOf((_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(699), new BigNumber(-34), new BigNumber(-43), new BigNumber((_dafny.Seq.UnicodeFromString("vqgjwd")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f38),true)).length))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("jhuvf")).length)))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ihdf")).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-567)), function (_1614_i0) {
        return _dafny.Seq.UnicodeFromString("ro");
      })).length), new BigNumber(871))).length)))));
    };
    m13(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _1615_v0;
      _1615_v0 = new BigNumber(616);
      let _1616_v1;
      _1616_v1 = _module.D12.create_DC28(_1615_v0);
      let _1617_v2;
      _1617_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(!(_this.f38), _1615_v0, globalState),_1615_v0);
      let _pat_let_tv38 = _1615_v0;
      let _pat_let_tv39 = _1615_v0;
      let _hi13 = (((_1617_v2).contains(_this.f38)) ? ((_1617_v2).get(_this.f38)) : (_1615_v0));
      for (let _1618_i0 = new BigNumber((_dafny.Seq.of(_1616_v1, _1616_v1, _1616_v1, function (_pat_let36_0) {
        return function (_1648_dt__update__tmp_h1) {
          return function (_pat_let39_0) {
            return function (_1649_dt__update_hcf49_h1) {
              return _module.D12.create_DC28(_1649_dt__update_hcf49_h1);
            }(_pat_let39_0);
          }(_pat_let_tv39);
        }(_pat_let36_0);
      }(function (_pat_let37_0) {
        return function (_1646_dt__update__tmp_h0) {
          return function (_pat_let38_0) {
            return function (_1647_dt__update_hcf49_h0) {
              return _module.D12.create_DC28(_1647_dt__update_hcf49_h0);
            }(_pat_let38_0);
          }(_pat_let_tv38);
        }(_pat_let37_0);
      }(_1616_v1)), _1616_v1)).length); _1618_i0.isLessThan(_hi13); _1618_i0 = _1618_i0.plus(_dafny.ONE)) {
        let _1619_v3;
        _1619_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1615_v0,_1618_i0);
        let _1620_v4;
        _1620_v4 = _module.D0.create_DC1(_this.f38, _this.f38, new BigNumber((_1619_v3).length), _this.f38);
        _1620_v4 = _1620_v4;
        let _1621_v5;
        _1621_v5 = _module.D5.create_DC15(_module.D5.create_DC14(_1618_i0, (_this).fm1(globalState), (_this).fm2(_this.f38, _1615_v0, globalState)));
        let _1622_v6;
        _1622_v6 = _module.D5.create_DC15(_1621_v5);
        let _1623_v7;
        let _nw247 = Array((new BigNumber(8)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = _1622_v6;
        _nw247[(_dafny.ONE).toNumber()] = _module.D5.create_DC15(_1621_v5);
        _nw247[(new BigNumber(2)).toNumber()] = _1622_v6;
        _nw247[(new BigNumber(3)).toNumber()] = _1622_v6;
        _nw247[(new BigNumber(4)).toNumber()] = _1622_v6;
        _nw247[(new BigNumber(5)).toNumber()] = _1622_v6;
        _nw247[(new BigNumber(6)).toNumber()] = _1622_v6;
        _nw247[(new BigNumber(7)).toNumber()] = _1622_v6;
        _1623_v7 = _nw247;
        let _index234 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1623_v7).length));
        (_1623_v7)[_index234] = ((!(!(_this.f38))) ? (_1622_v6) : (_1622_v6));
        let _pat_let_tv36 = _1621_v5;
        let _pat_let_tv37 = _1621_v5;
        let _index235 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_1623_v7).length));
        (_1623_v7)[_index235] = function (_pat_let32_0) {
          return function (_1626_dt__update__tmp_h3) {
            return function (_pat_let35_0) {
              return function (_1627_dt__update_hcf28_h1) {
                return _module.D5.create_DC15(_1627_dt__update_hcf28_h1);
              }(_pat_let35_0);
            }(_pat_let_tv37);
          }(_pat_let32_0);
        }(function (_pat_let33_0) {
          return function (_1624_dt__update__tmp_h2) {
            return function (_pat_let34_0) {
              return function (_1625_dt__update_hcf28_h0) {
                return _module.D5.create_DC15(_1625_dt__update_hcf28_h0);
              }(_pat_let34_0);
            }(_pat_let_tv36);
          }(_pat_let33_0);
        }(_module.D5.create_DC15(_module.D5.create_DC14((_dafny.ZERO).minus(_1615_v0), _1618_i0, _this.f38))));
        let _1628_v8;
        _1628_v8 = _dafny.Seq.of(_1615_v0);
        let _1629_v9;
        _1629_v9 = _dafny.Seq.of(_dafny.Seq.update(_1628_v8, _module.__default.safeIndex(_1618_i0, new BigNumber((_1628_v8).length)), _1615_v0));
        let _1630_v10;
        let _nw248 = new _module.C2();
        _nw248.__ctor(_dafny.Seq.update(_1629_v9, _module.__default.safeIndex(_1615_v0, new BigNumber((_1629_v9).length)), _dafny.Seq.of(_1615_v0, new BigNumber(250), _1615_v0, _1618_i0)), !(_this.f38), _this.f38);
        _1630_v10 = _nw248;
        let _1631_v11;
        _1631_v11 = _dafny.MultiSet.fromElements(new BigNumber(-248));
        let _1632_v12;
        _1632_v12 = _dafny.Seq.of(_this.f38, _this.f38);
        let _1633_v13;
        let _nw249 = Array((new BigNumber(19)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = _1618_i0;
        _nw249[(_dafny.ONE).toNumber()] = new BigNumber((_1631_v11).cardinality());
        _nw249[(new BigNumber(2)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(3)).toNumber()] = new BigNumber(812);
        _nw249[(new BigNumber(4)).toNumber()] = new BigNumber(661);
        _nw249[(new BigNumber(5)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(6)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(7)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(8)).toNumber()] = new BigNumber((_1632_v12).length);
        _nw249[(new BigNumber(9)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(10)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(11)).toNumber()] = _1615_v0;
        _nw249[(new BigNumber(12)).toNumber()] = _1615_v0;
        _nw249[(new BigNumber(13)).toNumber()] = (_1628_v8)[_module.__default.safeIndex(_1618_i0, new BigNumber((_1628_v8).length))];
        _nw249[(new BigNumber(14)).toNumber()] = _1615_v0;
        _nw249[(new BigNumber(15)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(16)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(17)).toNumber()] = _1618_i0;
        _nw249[(new BigNumber(18)).toNumber()] = _1618_i0;
        _1633_v13 = _nw249;
        let _1634_v14;
        _1634_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1633_v13,false);
        let _1635_v15;
        _1635_v15 = _module.D5.create_DC14(_1618_i0, (_dafny.ZERO).minus(new BigNumber((_1632_v12).length)), _this.f38);
        let _1636_v16;
        _1636_v16 = _module.D2.create_DC7(_this.f38, _this.f38, new BigNumber((_1634_v14).length), (_1635_v15).dtor_cf25);
        let _1637_v17;
        let _nw250 = Array((new BigNumber(7)).toNumber()).fill(false);
        _1637_v17 = _nw250;
        let _index236 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1637_v17).length));
        (_1637_v17)[_index236] = _this.f38;
        let _1638_v18;
        _1638_v18 = _dafny.Seq.UnicodeFromString("mhjfmmfl");
        let _1639_v19;
        _1639_v19 = _module.D3.create_DC11((_1638_v18)[_module.__default.safeIndex(_1618_i0, new BigNumber((_1638_v18).length))]);
        let _1640_v20;
        _1640_v20 = _dafny.Map.Empty.slice().updateUnsafe(false,_1639_v19);
        let _1641_v21;
        _1641_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1628_v8,_dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), ((_1642_v3) => function (_1643_i1) {
          return _module.D2.create_DC6(_1642_v3);
        })(_1619_v3)));
        let _1644_v22;
        _1644_v22 = _module.D10.create_DC23(_1641_v21);
        let _1645_v23;
        _1645_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1644_v22,_this.f38);
        let _index237 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1637_v17).length));
        let _rhs243 = ((_this.f38) ? (_1630_v10) : (_1630_v10));
        let _rhs244 = _1636_v16;
        let _rhs245 = new BigNumber((_1640_v20).length);
        let _rhs246 = _1618_i0;
        let _rhs247 = (((_1645_v23).contains(_1644_v22)) ? ((_1645_v23).get(_1644_v22)) : (_this.f38));
        let _lhs223 = globalState;
        let _lhs224 = globalState;
        let _lhs225 = _1637_v17;
        let _lhs226 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1637_v17).length));
        _1630_v10 = _rhs243;
        _1636_v16 = _rhs244;
        _lhs223.f16 = _rhs245;
        _lhs224.f12 = _rhs246;
        _lhs225[_lhs226] = _rhs247;
        let _index238 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_1637_v17).length));
        (_1637_v17)[_index238] = (_1637_v17)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_1637_v17).length))];
      }
      let _1650_v24;
      _1650_v24 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1651_v25;
      let _1652_v26;
      let _1653_v27;
      let _1654_v28;
      let _out49;
      let _out50;
      let _out51;
      let _out52;
      let _outcollector13 = _module.__default.m0(_1650_v24, globalState);
      _out49 = _outcollector13[0];
      _out50 = _outcollector13[1];
      _out51 = _outcollector13[2];
      _out52 = _outcollector13[3];
      _1651_v25 = _out49;
      _1652_v26 = _out50;
      _1653_v27 = _out51;
      _1654_v28 = _out52;
      let _1655_v29;
      _1655_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1651_v25,new BigNumber(369));
      let _rhs248 = !((_1615_v0).isLessThanOrEqualTo(_module.__default.safeDivisionInt(_1615_v0, (_this).fm1(globalState))));
      let _rhs249 = (_module.__default.safeModuloInt(new BigNumber(809), new BigNumber(846))).multipliedBy((((_1655_v29).contains(_module.__default.fm40(_1615_v0, (_this).fm2(_1653_v27, _1615_v0, globalState), _1653_v27, globalState))) ? ((_1655_v29).get(_module.__default.fm40(_1615_v0, (_this).fm2(_1653_v27, _1615_v0, globalState), _1653_v27, globalState))) : (_1615_v0)));
      let _lhs227 = globalState;
      _lhs227.f24 = _rhs248;
      r1 = _rhs249;
      let _1656_v30;
      _1656_v30 = _dafny.MultiSet.fromElements(_1615_v0, _1615_v0);
      let _1657_v31;
      _1657_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1656_v30,_this.f38);
      let _1658_v32;
      _1658_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1615_v0,new BigNumber((_1657_v31).length));
      _1658_v32 = (_1658_v32).update((_dafny.ZERO).minus(_1615_v0), _1615_v0);
      (_this).f38 = _1653_v27;
      let _1659_v33;
      _1659_v33 = _module.D10.create_DC25();
      let _pat_let_tv40 = _1656_v30;
      let _pat_let_tv41 = _1656_v30;
      let _pat_let_tv42 = _1615_v0;
      let _pat_let_tv43 = _1656_v30;
      _1656_v30 = function (_source25) {
        if (_source25.is_DC24) {
          let _1660___mcc_h0 = (_source25).cf44;
          let _1661___mcc_h1 = (_source25).cf45;
          let _1662___mcc_h2 = (_source25).cf46;
          let _1663_cf46 = _1662___mcc_h2;
          let _1664_cf45 = _1661___mcc_h1;
          let _1665_cf44 = _1660___mcc_h0;
          return _pat_let_tv40;
        } else if (_source25.is_DC25) {
          return (_pat_let_tv41).update(new BigNumber((_dafny.Seq.UnicodeFromString("iowvgu")).length), _module.__default.abs(_pat_let_tv42));
        } else {
          let _1666___mcc_h3 = (_source25).cf43;
          let _1667_cf43 = _1666___mcc_h3;
          return _pat_let_tv43;
        }
      }(_1659_v33);
      r0 = _1653_v27;
      r1 = (_1615_v0).multipliedBy(_1615_v0);
      r2 = _1653_v27;
      return [r0, r1, r2];
    }
    m14(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _1668_v0;
      _1668_v0 = new BigNumber(-497);
      let _1669_v1;
      _1669_v1 = _dafny.Seq.of(_1668_v0);
      let _1670_v2;
      let _nw251 = new _module.C2();
      _nw251.__ctor(_dafny.Seq.of(_1669_v1), !(p0), _this.f38);
      _1670_v2 = _nw251;
      let _1671_v3;
      _1671_v3 = _dafny.Seq.of((_1670_v2).fm2(p1, _1668_v0, globalState), p0);
      let _1672_v4;
      _1672_v4 = _module.D1.create_DC3(_dafny.Seq.Concat(_dafny.Seq.of(p1), _1671_v3));
      let _source26 = _1672_v4;
      if (_source26.is_DC4) {
        let _1673___mcc_h0 = (_source26).cf7;
        let _1674___mcc_h1 = (_source26).cf8;
        let _1675___mcc_h2 = (_source26).cf9;
        let _1676___mcc_h3 = (_source26).cf10;
        let _1677___mcc_h4 = (_source26).cf11;
        let _1678_cf11 = _1677___mcc_h4;
        let _1679_cf10 = _1676___mcc_h3;
        let _1680_cf9 = _1675___mcc_h2;
        let _1681_cf8 = _1674___mcc_h1;
        let _1682_cf7 = _1673___mcc_h0;
        let _1683_v5;
        _1683_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).fm2(p1, _1681_cf8, globalState));
        let _1684_v6;
        _1684_v6 = _dafny.Set.fromElements(_this.f38, (((_1683_v5).contains(_this.f38)) ? ((_1683_v5).get(_this.f38)) : (p1)), !(_this.f38), _this.f38, p0);
        let _1685_v7;
        _1685_v7 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (_1684_v6) : (_1684_v6)),p1);
        let _1686_v8;
        _1686_v8 = _1684_v6;
        let _rhs250 = ((_1684_v6).Difference(_1684_v6)).IsSubsetOf((_1684_v6).Intersect(_dafny.Set.fromElements(true)));
        let _rhs251 = _1671_v3;
        let _rhs252 = _1671_v3;
        let _rhs253 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1684_v6).length), _1681_cf8));
        let _rhs254 = (((_1685_v7).contains(_1686_v8)) ? ((_1685_v7).get(_1686_v8)) : (_this.f38));
        let _lhs228 = globalState;
        let _lhs229 = globalState;
        _lhs228.f22 = _rhs250;
        _1671_v3 = _rhs251;
        _1671_v3 = _rhs252;
        _1680_cf9 = _rhs253;
        _lhs229.f22 = _rhs254;
        let _1687_v9;
        _1687_v9 = _dafny.Seq.of(_1684_v6);
        let _1688_v10;
        _1688_v10 = _module.D14.create_DC31(_dafny.Seq.of(_1684_v6));
        let _1689_v11;
        _1689_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of(_1684_v6));
        let _1690_v12;
        let _nw252 = Array((new BigNumber(18)).toNumber());
        _nw252[(_dafny.ZERO).toNumber()] = _1687_v9;
        _nw252[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1687_v9, _dafny.Seq.update(_1687_v9, _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1687_v9).length)), _dafny.Set.fromElements(p0, true)));
        _nw252[(new BigNumber(2)).toNumber()] = _1687_v9;
        _nw252[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1687_v9, _module.__default.safeIndex(_1681_cf8, new BigNumber((_1687_v9).length)), _1684_v6), _1687_v9);
        _nw252[(new BigNumber(4)).toNumber()] = _1687_v9;
        _nw252[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1687_v9, _1687_v9);
        _nw252[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1684_v6, _1684_v6, _1684_v6), _1687_v9);
        _nw252[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_1684_v6);
        _nw252[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1687_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(799)), ((_1691_v6) => function (_1692_i0) {
          return _1691_v6;
        })(_1684_v6)));
        _nw252[(new BigNumber(9)).toNumber()] = _1687_v9;
        _nw252[(new BigNumber(10)).toNumber()] = _1687_v9;
        _nw252[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_1687_v9, _module.__default.safeIndex(_1668_v0, new BigNumber((_1687_v9).length)), _1684_v6);
        _nw252[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat((_1688_v10).dtor_cf55, _1687_v9);
        _nw252[(new BigNumber(13)).toNumber()] = (((_1689_v11).contains(p0)) ? ((_1689_v11).get(p0)) : (_1687_v9));
        _nw252[(new BigNumber(14)).toNumber()] = (_1688_v10).dtor_cf55;
        _nw252[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(785)), ((_1693_v6) => function (_1694_i1) {
          return _1693_v6;
        })(_1684_v6));
        _nw252[(new BigNumber(16)).toNumber()] = ((_this.f38) ? (_1687_v9) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), ((_1695_v6) => function (_1696_i2) {
          return _1695_v6;
        })(_1684_v6))));
        _nw252[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_1687_v9, _1687_v9);
        _1690_v12 = _nw252;
        let _index239 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1690_v12).length));
        (_1690_v12)[_index239] = _1687_v9;
        let _1697_v13;
        _1697_v13 = _module.D3.create_DC11(new _dafny.CodePoint('a'.codePointAt(0)));
        let _1698_v14;
        _1698_v14 = _dafny.Seq.of(_module.__default.fm41(!(_this.f38), _1697_v13, _1680_cf9, p0, globalState));
        let _1699_v15;
        let _nw253 = new _module.C5();
        _nw253.__ctor((_1671_v3)[_module.__default.safeIndex(_1680_cf9, new BigNumber((_1671_v3).length))], true);
        _1699_v15 = _nw253;
        let _1700_v16;
        _1700_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1699_v15,p1);
        let _1701_v17;
        _1701_v17 = _module.D15.create_DC33(_1700_v16);
        let _pat_let_tv44 = _1700_v16;
        let _index240 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1690_v12).length));
        let _rhs255 = _1681_cf8;
        let _rhs256 = _dafny.Seq.Concat(_1687_v9, (_1698_v14)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1679_cf10, _module.__default.safeIndex((_1670_v2).fm4(_module.D0.create_DC0(p1), _1668_v0, new BigNumber(((function (_pat_let40_0) {
          return function (_1702_dt__update__tmp_h0) {
            return function (_pat_let41_0) {
              return function (_1703_dt__update_hcf61_h0) {
                return _module.D15.create_DC33(_1703_dt__update_hcf61_h0);
              }(_pat_let41_0);
            }(_pat_let_tv44);
          }(_pat_let40_0);
        }(_1701_v17)).dtor_cf61).length), globalState), new BigNumber((_1679_cf10).length)), p2)).length), new BigNumber((_1698_v14).length))]);
        let _lhs230 = globalState;
        let _lhs231 = _1690_v12;
        let _lhs232 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1690_v12).length));
        _lhs230.f16 = _rhs255;
        _lhs231[_lhs232] = _rhs256;
        let _1704_v18;
        _1704_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1668_v0,_1680_cf9);
        if (((_dafny.ZERO).minus(((p0) ? (_1681_cf8) : (_1668_v0)))).isEqualTo(new BigNumber(((_1704_v18).update(_1668_v0, _1681_cf8)).length))) {
          _1704_v18 = (_1704_v18).update(new BigNumber(567), _1681_cf8);
          let _1705_v19;
          let _init44 = ((_1706_cf9) => function (_1707_i3) {
            return _module.__default.safeDivisionInt(_1707_i3, (_dafny.ZERO).minus(_1706_cf9));
          })(_1680_cf9);
          let _nw254 = Array((new BigNumber(21)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw254.length); _i0_44++) {
            _nw254[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1705_v19 = _nw254;
          let _index241 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1705_v19).length));
          (_1705_v19)[_index241] = (_1699_v15).fm1(globalState);
          let _index242 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_1705_v19).length));
          (_1705_v19)[_index242] = (_1669_v1)[_module.__default.safeIndex((_1699_v15).fm1(globalState), new BigNumber((_1669_v1).length))];
          let _1708_v20;
          let _nw255 = new _module.C0();
          _nw255.__ctor((_1681_cf8).isLessThanOrEqualTo((_1705_v19)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1705_v19).length))]));
          _1708_v20 = _nw255;
          let _rhs257 = _1708_v20;
          let _rhs258 = _1705_v19;
          let _rhs259 = ((_1670_v2).f35)[_module.__default.safeIndex(_1668_v0, new BigNumber(((_1670_v2).f35).length))];
          _1708_v20 = _rhs257;
          _1705_v19 = _rhs258;
          _1669_v1 = _rhs259;
          let _1709_v21;
          _1709_v21 = _module.D15.create_DC34(_1681_cf8, p0, new BigNumber((_1678_cf11).length));
          _1709_v21 = _module.D15.create_DC34((_1705_v19)[_module.__default.safeIndex(new BigNumber(658), new BigNumber((_1705_v19).length))], false, _module.__default.safeDivisionInt(_1668_v0, _1680_cf9));
          let _1710_v22;
          let _nw256 = new _module.C5();
          _nw256.__ctor(true, (true) || (true));
          _1710_v22 = _nw256;
        } else {
          let _1711_v23;
          let _nw257 = Array((_dafny.ONE).toNumber()).fill(false);
          _1711_v23 = _nw257;
          _1711_v23 = _1711_v23;
          let _1712_v24;
          let _nw258 = Array((new BigNumber(9)).toNumber());
          _nw258[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_1680_cf9, _1681_cf8);
          _nw258[(_dafny.ONE).toNumber()] = ((p1) ? (_1680_cf9) : (_1680_cf9));
          _nw258[(new BigNumber(2)).toNumber()] = _1680_cf9;
          _nw258[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1671_v3).length), _1680_cf9));
          _nw258[(new BigNumber(4)).toNumber()] = new BigNumber(836);
          _nw258[(new BigNumber(5)).toNumber()] = _1681_cf8;
          _nw258[(new BigNumber(6)).toNumber()] = _1668_v0;
          _nw258[(new BigNumber(7)).toNumber()] = ((_dafny.ZERO).minus(_1681_cf8)).multipliedBy(new BigNumber((_1679_cf10).length));
          _nw258[(new BigNumber(8)).toNumber()] = _1680_cf9;
          _1712_v24 = _nw258;
          let _index243 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1712_v24).length));
          (_1712_v24)[_index243] = new BigNumber((_dafny.Seq.update(_1671_v3, _module.__default.safeIndex(_1668_v0, new BigNumber((_1671_v3).length)), _this.f38)).length);
          let _index244 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_1712_v24).length));
          (_1712_v24)[_index244] = _1668_v0;
          let _1713_v25;
          _1713_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1711_v23,_1680_cf9);
          let _1714_v26;
          _1714_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1713_v25).length));
          (globalState).f12 = new BigNumber(((_1714_v26).Merge(((p1) ? (_1714_v26) : (_dafny.Map.Empty.slice().updateUnsafe(p0,(_1712_v24)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_1712_v24).length))]))))).length);
          let _1715_v27;
          _1715_v27 = _dafny.MultiSet.fromElements(p1, !(p1), _this.f38);
          let _1716_v28;
          _1716_v28 = _dafny.Set.fromElements((_1712_v24)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_1712_v24).length))]);
          let _1717_v29;
          _1717_v29 = _module.D2.create_DC9((_1668_v0).minus((_dafny.ZERO).minus(_1668_v0)), (_dafny.ZERO).minus(_1668_v0), (_1716_v28).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_1715_v27).cardinality()))));
          _1717_v29 = _1717_v29;
          (globalState).f16 = _1680_cf9;
        }
        (globalState).f22 = p0;
      } else if (_source26.is_DC5) {
        let _1718___mcc_h5 = (_source26).cf12;
        let _1719_cf12 = _1718___mcc_h5;
        r3 = p1;
        let _1720_v30;
        _1720_v30 = _dafny.Set.fromElements(_1719_cf12, _1668_v0);
        let _1721_v31;
        let _nw259 = Array((new BigNumber(2)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = p2;
        _nw259[(_dafny.ONE).toNumber()] = p2;
        _1721_v31 = _nw259;
        let _1722_v32;
        _1722_v32 = _dafny.Seq.UnicodeFromString("ryndpso");
        let _1723_v33;
        _1723_v33 = _module.D1.create_DC4(_1721_v31, _1719_cf12, _1668_v0, _1722_v32, _1722_v32);
        let _1724_v34;
        _1724_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1725_v35;
        _1725_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1719_cf12,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vi")).length)));
        let _1726_v36;
        _1726_v36 = _dafny.Set.fromElements(_this.f38, p0, true, p1);
        let _1727_v37;
        _1727_v37 = _dafny.MultiSet.fromElements(_1668_v0, _1719_cf12, _1719_cf12, _1719_cf12);
        let _1728_v38;
        let _nw260 = Array((new BigNumber(24)).toNumber());
        _nw260[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_1668_v0);
        _nw260[(_dafny.ONE).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(2)).toNumber()] = (_1723_v33).dtor_cf8;
        _nw260[(new BigNumber(3)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_1668_v0);
        _nw260[(new BigNumber(5)).toNumber()] = _1668_v0;
        _nw260[(new BigNumber(6)).toNumber()] = new BigNumber(((_1724_v34).Merge(_1724_v34)).length);
        _nw260[(new BigNumber(7)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(8)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(9)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(10)).toNumber()] = (_1670_v2).fm1(globalState);
        _nw260[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(_1719_cf12, _1719_cf12);
        _nw260[(new BigNumber(12)).toNumber()] = _1668_v0;
        _nw260[(new BigNumber(13)).toNumber()] = new BigNumber((_module.__default.fm42(_1668_v0, globalState)).cardinality());
        _nw260[(new BigNumber(14)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(15)).toNumber()] = _1668_v0;
        _nw260[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("j")).length);
        _nw260[(new BigNumber(17)).toNumber()] = _1719_cf12;
        _nw260[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_1725_v35).length), new BigNumber((_1726_v36).length));
        _nw260[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_1719_cf12, _1668_v0)).cardinality());
        _nw260[(new BigNumber(20)).toNumber()] = _1668_v0;
        _nw260[(new BigNumber(21)).toNumber()] = (_1668_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1727_v37).cardinality())));
        _nw260[(new BigNumber(22)).toNumber()] = _1668_v0;
        _nw260[(new BigNumber(23)).toNumber()] = _1668_v0;
        _1728_v38 = _nw260;
        let _index245 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1728_v38).length));
        (_1728_v38)[_index245] = new BigNumber((_dafny.Seq.Concat(_1669_v1, _1669_v1)).length);
        let _index246 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1728_v38).length));
        let _rhs260 = (_1720_v30).Difference(_1720_v30);
        let _rhs261 = _module.__default.fm43(_1668_v0, new BigNumber(909), new BigNumber((_dafny.Seq.Concat(_1671_v3, _1671_v3)).length), _dafny.Map.Empty.slice().updateUnsafe(p1,p2), globalState);
        let _rhs262 = (p0) && (!(p0));
        let _rhs263 = _1668_v0;
        let _rhs264 = (_1719_cf12).plus(_1719_cf12);
        let _lhs233 = globalState;
        let _lhs234 = globalState;
        let _lhs235 = _1728_v38;
        let _lhs236 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1728_v38).length));
        _1720_v30 = _rhs260;
        _1669_v1 = _rhs261;
        _lhs233.f15 = _rhs262;
        _lhs234.f16 = _rhs263;
        _lhs235[_lhs236] = _rhs264;
        let _1729_v39;
        let _nw261 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1729_v39 = _nw261;
        _1729_v39 = _1729_v39;
        let _1730_v40;
        _1730_v40 = _module.D0.create_DC0(_this.f38);
        let _pat_let_tv45 = p0;
        _1719_cf12 = (_1670_v2).fm4(function (_pat_let42_0) {
          return function (_1731_dt__update__tmp_h1) {
            return function (_pat_let43_0) {
              return function (_1732_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_1732_dt__update_hcf0_h0);
              }(_pat_let43_0);
            }(_pat_let_tv45);
          }(_pat_let42_0);
        }(_1730_v40), _module.__default.safeModuloInt(_1668_v0, _1719_cf12), (_1719_cf12).minus(_1719_cf12), globalState);
      } else {
        let _1733___mcc_h6 = (_source26).cf6;
        let _1734_cf6 = _1733___mcc_h6;
        let _1735_v41;
        let _nw262 = new _module.C6();
        _nw262.__ctor(_1668_v0);
        _1735_v41 = _nw262;
        let _1736_v42;
        let _1737_v43;
        let _out53;
        let _out54;
        let _outcollector14 = (_1670_v2).m2(!(_1668_v0).isEqualTo(_1668_v0), _1735_v41, _this.f38, globalState);
        _out53 = _outcollector14[0];
        _out54 = _outcollector14[1];
        _1736_v42 = _out53;
        _1737_v43 = _out54;
        if (p1) {
          (globalState).f15 = p1;
          (globalState).f22 = (_1735_v41).fm2(p1, _1668_v0, globalState);
          let _1738_v44;
          _1738_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _1739_v45;
          _1739_v45 = _module.D13.create_DC29(_1738_v44);
          let _1740_v46;
          _1740_v46 = _dafny.Set.fromElements(_1739_v45, _1739_v45);
          (globalState).f15 = (_1740_v46).IsDisjointFrom(_1740_v46);
          let _1741_v47;
          _1741_v47 = _dafny.MultiSet.fromElements(p0);
          let _1742_v48;
          _1742_v48 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1738_v44).length),false));
          (globalState).f24 = (((_dafny.ZERO).minus(new BigNumber((_1741_v47).cardinality()))).isEqualTo(new BigNumber((_1742_v48).length))) === ((new BigNumber((_1734_cf6).length)).isEqualTo(new BigNumber(505)));
          _1668_v0 = _1737_v43;
        } else {
          let _1743_v49;
          _1743_v49 = _dafny.Set.fromElements(_1737_v43);
          let _1744_v50;
          _1744_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1668_v0,p1);
          let _1745_v51;
          _1745_v51 = _dafny.Seq.UnicodeFromString("bam");
          let _1746_v52;
          _1746_v52 = _module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(p0,_1671_v3));
          let _1747_v53;
          _1747_v53 = _module.D5.create_DC15(_1746_v52);
          let _1748_v54;
          _1748_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1747_v53,_1737_v43);
          let _1749_v55;
          let _nw263 = Array((new BigNumber(25)).toNumber());
          _nw263[(_dafny.ZERO).toNumber()] = _1737_v43;
          _nw263[(_dafny.ONE).toNumber()] = new BigNumber((_1743_v49).length);
          _nw263[(new BigNumber(2)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(3)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(4)).toNumber()] = new BigNumber(-377);
          _nw263[(new BigNumber(5)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(6)).toNumber()] = _1737_v43;
          _nw263[(new BigNumber(7)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((((_1744_v50).contains(_1737_v43)) ? ((_1744_v50).get(_1737_v43)) : (_this.f38))),_1737_v43)).length);
          _nw263[(new BigNumber(9)).toNumber()] = (_1735_v41).fm1(globalState);
          _nw263[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_1668_v0);
          _nw263[(new BigNumber(11)).toNumber()] = _1737_v43;
          _nw263[(new BigNumber(12)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(13)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_1737_v43);
          _nw263[(new BigNumber(15)).toNumber()] = _1737_v43;
          _nw263[(new BigNumber(16)).toNumber()] = new BigNumber((_1745_v51).length);
          _nw263[(new BigNumber(17)).toNumber()] = (_1735_v41).fm1(globalState);
          _nw263[(new BigNumber(18)).toNumber()] = new BigNumber((_1748_v54).length);
          _nw263[(new BigNumber(19)).toNumber()] = _1737_v43;
          _nw263[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(_1668_v0);
          _nw263[(new BigNumber(21)).toNumber()] = new BigNumber((_1745_v51).length);
          _nw263[(new BigNumber(22)).toNumber()] = _1668_v0;
          _nw263[(new BigNumber(23)).toNumber()] = _1737_v43;
          _nw263[(new BigNumber(24)).toNumber()] = _1737_v43;
          _1749_v55 = _nw263;
          let _1750_v56;
          _1750_v56 = _dafny.MultiSet.fromElements(_this.f38, false, false);
          let _1751_v57;
          _1751_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1749_v55,_1750_v56);
          _1751_v57 = (_1751_v57).Merge(_1751_v57);
          (globalState).f22 = p0;
          (globalState).f6 = _dafny.Seq.UnicodeFromString("qx");
          let _index247 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_1749_v55).length));
          (_1749_v55)[_index247] = _1668_v0;
          let _index248 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_1749_v55).length));
          (_1749_v55)[_index248] = (_module.D1.create_DC5((_dafny.ZERO).minus(_1737_v43))).dtor_cf12;
          (globalState).f15 = _dafny.Seq.IsPrefixOf(_1745_v51, _dafny.Seq.Create(_module.__default.abs(new BigNumber(657)), function (_1752_i4) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }));
        }
        let _1753_v58;
        let _nw264 = new _module.C1();
        _nw264.__ctor(_1668_v0, _1668_v0, p1, p0);
        _1753_v58 = _nw264;
        let _1754_v59;
        _1754_v59 = _dafny.MultiSet.fromElements(_1753_v58, _1753_v58, _1753_v58, _1753_v58);
        let _1755_v60;
        _1755_v60 = _module.D15.create_DC34(new BigNumber(-234), p1, _1668_v0);
        let _1756_v61;
        _1756_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(944),(_1755_v60).dtor_cf63)).length));
        let _1757_v62;
        _1757_v62 = _dafny.Set.fromElements(true, p1, !((_1753_v58).f27), (_1753_v58).f26, (_1753_v58).f26);
        let _1758_v63;
        _1758_v63 = _dafny.MultiSet.fromElements(p2);
        let _1759_v64;
        let _nw265 = Array((new BigNumber(8)).toNumber());
        _nw265[(_dafny.ZERO).toNumber()] = ((p0) ? (new BigNumber((_1669_v1).length)) : (_1737_v43));
        _nw265[(_dafny.ONE).toNumber()] = _1737_v43;
        _nw265[(new BigNumber(2)).toNumber()] = new BigNumber((_1754_v59).cardinality());
        _nw265[(new BigNumber(3)).toNumber()] = _1737_v43;
        _nw265[(new BigNumber(4)).toNumber()] = new BigNumber((_1734_cf6).length);
        _nw265[(new BigNumber(5)).toNumber()] = new BigNumber(405);
        _nw265[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((((_1756_v61).contains((_1753_v58).f26)) ? ((_1756_v61).get((_1753_v58).f26)) : (new BigNumber((_1757_v62).length))), _1668_v0);
        _nw265[(new BigNumber(7)).toNumber()] = new BigNumber(((_module.__default.fm48(new BigNumber((_dafny.Seq.of(_1668_v0)).length), globalState)).Union(_1758_v63)).cardinality());
        _1759_v64 = _nw265;
        let _index249 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1759_v64).length));
        (_1759_v64)[_index249] = _1737_v43;
        let _1760_v65;
        _1760_v65 = _module.D0.create_DC0((_1734_cf6)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(362)), ((_1761_p2) => function (_1762_i5) {
  return _1761_p2;
})(p2))).length), new BigNumber((_1734_cf6).length))]);
        let _index250 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1759_v64).length));
        let _rhs265 = (_dafny.ZERO).minus((_1670_v2).fm4(_1760_v65, _1737_v43, _1737_v43, globalState));
        let _rhs266 = new BigNumber((_1734_cf6).length);
        let _lhs237 = _1759_v64;
        let _lhs238 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1759_v64).length));
        _lhs237[_lhs238] = _rhs265;
        _1737_v43 = _rhs266;
        let _1763_v66;
        _1763_v66 = _dafny.MultiSet.fromElements(new BigNumber(-658));
        let _1764_v67;
        _1764_v67 = _dafny.MultiSet.fromElements((_1763_v66).Difference(_1763_v66));
        _1764_v67 = _module.__default.fm49((_1763_v66).update(_1737_v43, _module.__default.abs(new BigNumber(-948))), new BigNumber((_dafny.Seq.update(_1734_cf6, _module.__default.safeIndex(_1668_v0, new BigNumber((_1734_cf6).length)), p1)).length), (_1753_v58).f27, globalState);
      }
      let _hi14 = _1668_v0;
      for (let _1765_i6 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(206),_1668_v0)).length); _1765_i6.isLessThan(_hi14); _1765_i6 = _1765_i6.plus(_dafny.ONE)) {
        let _1766_v68;
        _1766_v68 = _dafny.MultiSet.fromElements(_1668_v0);
        let _1767_v69;
        _1767_v69 = _dafny.Set.fromElements(p1, _this.f38);
        let _1768_v70;
        _1768_v70 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1668_v0),p1);
        let _1769_v71;
        _1769_v71 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1766_v68).cardinality()),p1), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1767_v69).length),p0), _1768_v70, _1768_v70);
        let _1770_v72;
        _1770_v72 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f38),_1769_v71);
        r3 = !(((_1769_v71).Intersect((((_1770_v72).contains(p0)) ? ((_1770_v72).get(p0)) : (_1769_v71)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_1668_v0,p1))));
        let _1771_v73;
        _1771_v73 = _dafny.MultiSet.fromElements(true);
        let _1772_v74;
        _1772_v74 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,p0);
        let _1773_v75;
        let _nw266 = Array((new BigNumber(21)).toNumber());
        _nw266[(_dafny.ZERO).toNumber()] = p0;
        _nw266[(_dafny.ONE).toNumber()] = p0;
        _nw266[(new BigNumber(2)).toNumber()] = p1;
        _nw266[(new BigNumber(3)).toNumber()] = true;
        _nw266[(new BigNumber(4)).toNumber()] = p0;
        _nw266[(new BigNumber(5)).toNumber()] = p0;
        _nw266[(new BigNumber(6)).toNumber()] = (_1671_v3)[_module.__default.safeIndex((((_1771_v73).contains(_this.f38)) ? ((_1771_v73).get(_this.f38)) : (_1668_v0)), new BigNumber((_1671_v3).length))];
        _nw266[(new BigNumber(7)).toNumber()] = _this.f38;
        _nw266[(new BigNumber(8)).toNumber()] = _this.f38;
        _nw266[(new BigNumber(9)).toNumber()] = p0;
        _nw266[(new BigNumber(10)).toNumber()] = !(true);
        _nw266[(new BigNumber(11)).toNumber()] = p1;
        _nw266[(new BigNumber(12)).toNumber()] = _this.f38;
        _nw266[(new BigNumber(13)).toNumber()] = _this.f38;
        _nw266[(new BigNumber(14)).toNumber()] = (_1670_v2).fm2((((_1772_v74).contains(_this.f38)) ? ((_1772_v74).get(_this.f38)) : (_this.f38)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-511)), ((_1774_p2) => function (_1775_i7) {
          return _1774_p2;
        })(p2))).length), globalState);
        _nw266[(new BigNumber(15)).toNumber()] = p1;
        _nw266[(new BigNumber(16)).toNumber()] = p1;
        _nw266[(new BigNumber(17)).toNumber()] = p1;
        _nw266[(new BigNumber(18)).toNumber()] = p0;
        _nw266[(new BigNumber(19)).toNumber()] = true;
        _nw266[(new BigNumber(20)).toNumber()] = !((_module.D2.create_DC7(p1, true, _1668_v0, _1668_v0)).dtor_cf14);
        _1773_v75 = _nw266;
        let _1776_v76;
        _1776_v76 = _module.D14.create_DC32(_1771_v73, p1, p0, p1, new BigNumber((_1767_v69).length));
        let _1777_v77;
        _1777_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1776_v76,_this.f38);
        let _index251 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1773_v75).length));
        (_1773_v75)[_index251] = (_1670_v2).fm17(_1668_v0, _1765_i6, new BigNumber((_1777_v77).length), _module.__default.fm50(globalState), globalState);
        let _index252 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1773_v75).length));
        let _rhs267 = p1;
        let _rhs268 = _1765_i6;
        let _lhs239 = _1773_v75;
        let _lhs240 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1773_v75).length));
        let _lhs241 = globalState;
        _lhs239[_lhs240] = _rhs267;
        _lhs241.f16 = _rhs268;
        let _1778_v78;
        let _nw267 = new _module.C6();
        _nw267.__ctor(new BigNumber((_1766_v68).cardinality()));
        _1778_v78 = _nw267;
        let _1779_v79;
        let _1780_v80;
        let _out55;
        let _out56;
        let _outcollector15 = (_1670_v2).m2(p0, _1778_v78, !(((p1) ? (_this.f38) : (p0))), globalState);
        _out55 = _outcollector15[0];
        _out56 = _outcollector15[1];
        _1779_v79 = _out55;
        _1780_v80 = _out56;
        let _1781_v81;
        _1781_v81 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1668_v0);
        let _1782_v82;
        _1782_v82 = _module.D7.create_DC19(_1668_v0, _dafny.MultiSet.fromElements(p1, p0), (_1773_v75)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_1773_v75).length))], _1781_v81);
        let _1783_v83;
        _1783_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p1,_this.f38),(_1782_v82).dtor_cf36);
        let _index253 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1773_v75).length));
        (_1773_v75)[_index253] = (((_1768_v70).contains(_module.__default.safeModuloInt(new BigNumber(641), (_dafny.ZERO).minus(new BigNumber((_1783_v83).length))))) ? ((_1768_v70).get(_module.__default.safeModuloInt(new BigNumber(641), (_dafny.ZERO).minus(new BigNumber((_1783_v83).length))))) : (p0));
      }
      (globalState).f16 = _1668_v0;
      let _1784_v84;
      let _nw268 = Array((new BigNumber(6)).toNumber()).fill(false);
      _1784_v84 = _nw268;
      _1784_v84 = _1784_v84;
      let _1785_v85;
      _1785_v85 = _dafny.Seq.UnicodeFromString("hcjrgaf");
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("tixwcfv"), _1785_v85)) {
        let _1786_v86;
        let _nw269 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1786_v86 = _nw269;
        let _1787_v87;
        _1787_v87 = _module.D3.create_DC10(_1786_v86);
        let _1788_v88;
        _1788_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1787_v87,new BigNumber((_dafny.Seq.of(_1668_v0)).length));
        let _1789_v89;
        let _nw270 = new _module.C4();
        _nw270.__ctor(_1788_v88, p1, _this.f38);
        _1789_v89 = _nw270;
        let _1790_v90;
        _1790_v90 = _dafny.Seq.of(_1785_v85, _1785_v85, _1785_v85, _1785_v85);
        (_1670_v2).m1(((!(!(false))) ? (true) : (_this.f38)), _1790_v90, _dafny.Seq.Concat(_1671_v3, _1671_v3), globalState);
        let _1791_v91;
        _1791_v91 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), function (_1792_i8) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }));
        let _rhs269 = _1791_v91;
        let _rhs270 = true;
        let _rhs271 = _1668_v0;
        let _rhs272 = _dafny.Seq.Concat(_1785_v85, _1785_v85);
        let _lhs242 = _this;
        let _lhs243 = globalState;
        let _lhs244 = globalState;
        _1791_v91 = _rhs269;
        _lhs242.f38 = _rhs270;
        _lhs243.f12 = _rhs271;
        _lhs244.f6 = _rhs272;
        r0 = false;
        let _1793_v92;
        let _nw271 = Array((new BigNumber(9)).toNumber());
        _nw271[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("pnu");
        _nw271[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1785_v85, _dafny.Seq.UnicodeFromString("uxsmk"));
        _nw271[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1785_v85, _1785_v85);
        _nw271[(new BigNumber(3)).toNumber()] = _1785_v85;
        _nw271[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("m");
        _nw271[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("fws");
        _nw271[(new BigNumber(6)).toNumber()] = _1785_v85;
        _nw271[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("cobamhx");
        _nw271[(new BigNumber(8)).toNumber()] = _1785_v85;
        _1793_v92 = _nw271;
        _1793_v92 = _1793_v92;
      } else {
        let _1794_v93;
        _1794_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1668_v0,new BigNumber(223));
        let _1795_v94;
        _1795_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1784_v84,new BigNumber((_dafny.MultiSet.fromElements(false, _this.f38, _this.f38, p1, p1)).cardinality()));
        let _1796_v95;
        _1796_v95 = _dafny.Set.fromElements(_this.f38, !(true), p0, true, false);
        let _1797_v96;
        let _nw272 = Array((new BigNumber(18)).toNumber());
        _nw272[(_dafny.ZERO).toNumber()] = _1668_v0;
        _nw272[(_dafny.ONE).toNumber()] = (new BigNumber((_1794_v93).length)).minus(new BigNumber(647));
        _nw272[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1785_v85).length)));
        _nw272[(new BigNumber(3)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(4)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_1668_v0, _1668_v0);
        _nw272[(new BigNumber(6)).toNumber()] = new BigNumber(834);
        _nw272[(new BigNumber(7)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(8)).toNumber()] = (_1668_v0).plus(_1668_v0);
        _nw272[(new BigNumber(9)).toNumber()] = (((_1795_v94).contains(_1784_v84)) ? ((_1795_v94).get(_1784_v84)) : (_1668_v0));
        _nw272[(new BigNumber(10)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(11)).toNumber()] = ((_module.D16.create_DC38(_1668_v0, _1668_v0, _this.f38)).dtor_cf70).minus((_dafny.ZERO).minus(_1668_v0));
        _nw272[(new BigNumber(12)).toNumber()] = new BigNumber(174);
        _nw272[(new BigNumber(13)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(14)).toNumber()] = new BigNumber((_1796_v95).length);
        _nw272[(new BigNumber(15)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(16)).toNumber()] = _1668_v0;
        _nw272[(new BigNumber(17)).toNumber()] = ((!(_this.f38)) ? (_1668_v0) : (_1668_v0));
        _1797_v96 = _nw272;
        let _index254 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length));
        (_1797_v96)[_index254] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_1668_v0), _1668_v0);
        let _1798_v97;
        _1798_v97 = _module.D0.create_DC0(p0);
        let _1799_v98;
        _1799_v98 = _module.D0.create_DC2(_1798_v97);
        let _1800_v99;
        _1800_v99 = _dafny.Seq.of(_1799_v98);
        let _index255 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length));
        let _rhs273 = _1668_v0;
        let _rhs274 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1668_v0), new BigNumber(-14)), _1668_v0);
        let _rhs275 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1785_v85, _1785_v85)).length))).multipliedBy((_1670_v2).fm1(globalState));
        let _rhs276 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(63)), ((_1801_v98) => function (_1802_i9) {
          return _1801_v98;
        })(_1799_v98));
        let _rhs277 = _1668_v0;
        let _lhs245 = globalState;
        let _lhs246 = _1797_v96;
        let _lhs247 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length));
        let _lhs248 = globalState;
        _1668_v0 = _rhs273;
        _lhs245.f12 = _rhs274;
        _lhs246[_lhs247] = _rhs275;
        _1800_v99 = _rhs276;
        _lhs248.f12 = _rhs277;
        (globalState).f22 = p1;
        let _1803_v100;
        let _nw273 = Array((new BigNumber(15)).toNumber()).fill([]);
        _1803_v100 = _nw273;
        let _index256 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1803_v100).length));
        (_1803_v100)[_index256] = _1797_v96;
        let _index257 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_1803_v100).length));
        (_1803_v100)[_index257] = _1797_v96;
        _1796_v95 = (_1796_v95).Union(_dafny.Set.fromElements(p1));
        let _1804_v101;
        _1804_v101 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]);
        if (((true) ? ((_1804_v101).equals(_1804_v101)) : (_this.f38))) {
          (globalState).f24 = ((_this.f38) && (false)) || (p1);
          (globalState).f15 = _this.f38;
          let _1805_v102;
          _1805_v102 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _1806_v103;
          _1806_v103 = _dafny.Map.Empty.slice().updateUnsafe((_1803_v100)[_module.__default.safeIndex(new BigNumber(144), new BigNumber((_1803_v100).length))],p0);
          _1805_v102 = (_1805_v102).update((((_1806_v103).contains(_1797_v96)) ? ((_1806_v103).get(_1797_v96)) : (p0)), _this.f38);
          _1784_v84 = _1784_v84;
          let _index258 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length));
          (_1797_v96)[_index258] = ((true) ? ((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]) : ((_1668_v0).minus(_1668_v0)));
        } else {
          let _1807_v104;
          _1807_v104 = _dafny.Set.fromElements(((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]).multipliedBy((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]), _1668_v0, new BigNumber(177), new BigNumber(986), (_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]);
          let _1808_v105;
          _1808_v105 = _dafny.Seq.of(_1785_v85, _1785_v85, _1785_v85);
          let _1809_v106;
          _1809_v106 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1808_v105, _1808_v105)),_dafny.Set.fromElements((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))], _1668_v0));
          let _1810_v107;
          _1810_v107 = _dafny.MultiSet.fromElements(_1785_v85);
          _1807_v104 = (((_1809_v106).contains((_1810_v107).update(_1785_v85, _module.__default.abs(_1668_v0)))) ? ((_1809_v106).get((_1810_v107).update(_1785_v85, _module.__default.abs(_1668_v0)))) : (_1807_v104));
          let _1811_v108;
          _1811_v108 = _module.D17.create_DC40((_1670_v2).f35);
          let _1812_v109;
          _1812_v109 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))], _1668_v0));
          let _1813_v110;
          _1813_v110 = _dafny.MultiSet.fromElements(_1669_v1);
          let _1814_v111;
          _1814_v111 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.update((_1811_v108).dtor_cf73, _module.__default.safeIndex((_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))], new BigNumber(((_1811_v108).dtor_cf73).length)), _dafny.Seq.of(new BigNumber(((((_1812_v109).contains(p2)) ? ((_1812_v109).get(p2)) : (_1669_v1))).length)))), _1813_v110, _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), ((_1815_v2, _1816_v0) => function (_1817_i10) {
            return ((_1815_v2).f35)[_module.__default.safeIndex(_1816_v0, new BigNumber(((_1815_v2).f35).length))];
          })(_1670_v2, _1668_v0))), _1813_v110, _1813_v110);
          let _1818_v112;
          let _nw274 = new _module.C3();
          _nw274.__ctor((_1814_v111)[_module.__default.safeIndex(_1668_v0, new BigNumber((_1814_v111).length))], _this.f38, p1);
          _1818_v112 = _nw274;
          let _1819_v113;
          _1819_v113 = _dafny.Seq.of(_1794_v93, _1794_v93);
          _1819_v113 = _1819_v113;
          (globalState).f16 = _module.__default.safeDivisionInt(_1668_v0, (_1797_v96)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_1797_v96).length))]);
          let _1820_v114;
          let _nw275 = new _module.C2();
          _nw275.__ctor((_1670_v2).f35, false, false);
          _1820_v114 = _nw275;
        }
      }
      r0 = p1;
      let _1821_v115;
      _1821_v115 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,_1668_v0);
      let _1822_v116;
      _1822_v116 = _dafny.MultiSet.fromElements(_1821_v115, _1821_v115);
      let _1823_v117;
      _1823_v117 = _module.D1.create_DC5(new BigNumber(226));
      r1 = ((_1822_v116).update(_1821_v115, _module.__default.abs((_1823_v117).dtor_cf12))).Intersect((_dafny.MultiSet.fromElements(_1821_v115)).Difference(_dafny.MultiSet.fromElements(_1821_v115, _1821_v115)));
      let _1824_v118;
      _1824_v118 = _dafny.Map.Empty.slice().updateUnsafe(_this.f38,p2);
      let _1825_v119;
      _1825_v119 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1671_v3, _1671_v3),new BigNumber((((p1) ? (_dafny.Map.Empty.slice().updateUnsafe(p1,p2)) : (_1824_v118))).length));
      r2 = _1825_v119;
      r3 = p1;
      return [r0, r1, r2, r3];
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f29 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T3];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f29) {
      let _this = this;
      (_this)._f29 = f29;
      return;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = [];
      let r3 = undefined;
      (globalState).f16 = p0;
      let _hi15 = p0;
      for (let _1826_i0 = p0; _1826_i0.isLessThan(_hi15); _1826_i0 = _1826_i0.plus(_dafny.ONE)) {
        let _1827_v0;
        _1827_v0 = new _dafny.CodePoint('y'.codePointAt(0));
        if ((_module.__default.safeDivisionInt(p3, new BigNumber(-589))).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(p2, _module.__default.safeIndex(p1, new BigNumber((p2).length)), _1827_v0)).length)))) {
          let _1828_v1;
          _1828_v1 = false;
          let _1829_v3;
          _1829_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1826_i0);
          let _1830_v4;
          _1830_v4 = _dafny.Set.fromElements(new BigNumber(-207), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((function () {
            let _coll67 = new _dafny.Map();
            for (const _compr_67 of (_1829_v3).Keys.Elements) {
              let _1831_v2 = _compr_67;
              if ((_1829_v3).contains(_1831_v2)) {
                _coll67.push([_module.__default.safeDivisionInt(_1831_v2, _1826_i0),_1826_i0]);
              }
            }
            return _coll67;
          }()).length))).length), p3);
          let _1832_v5;
          _1832_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1828_v1,_1830_v4);
          let _1833_v6;
          _1833_v6 = _dafny.Map.Empty.slice().updateUnsafe((_1826_i0).multipliedBy(p0),(_1832_v5).Merge(_1832_v5));
          _1833_v6 = (_1833_v6).update(new BigNumber((p2).length), (_1832_v5).Merge(_1832_v5));
          _1829_v3 = (_1829_v3).update(_module.__default.fm39(globalState), p1);
          let _1834_v7;
          let _init45 = ((_1835_p1) => function (_1836_i1) {
            return (_1836_i1).plus(_1835_p1);
          })(p1);
          let _nw276 = Array((new BigNumber(11)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw276.length); _i0_45++) {
            _nw276[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1834_v7 = _nw276;
          let _index259 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1834_v7).length));
          (_1834_v7)[_index259] = (_dafny.ZERO).minus(p3);
          let _1837_v8;
          let _init46 = ((_1838_v1) => function (_1839_i2) {
            return _1838_v1;
          })(_1828_v1);
          let _nw277 = Array((new BigNumber(7)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw277.length); _i0_46++) {
            _nw277[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1837_v8 = _nw277;
          let _index260 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1837_v8).length));
          (_1837_v8)[_index260] = !((_1828_v1) && (_1828_v1));
          let _1840_v9;
          _1840_v9 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
          let _1841_v10;
          _1841_v10 = _module.D13.create_DC29(_1840_v9);
          let _1842_v11;
          _1842_v11 = _dafny.Seq.of(_1826_i0, new BigNumber((p2).length), _1826_i0, p3, p0);
          let _index261 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1834_v7).length));
          let _index262 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1837_v8).length));
          let _rhs278 = (new BigNumber(109)).multipliedBy((new BigNumber(((_1841_v10).dtor_cf50).length)).multipliedBy(_1826_i0));
          let _rhs279 = _1828_v1;
          let _rhs280 = (p0).multipliedBy(p0);
          let _rhs281 = ((!(_1828_v1)) ? (_dafny.Seq.contains(_1842_v11, p3)) : ((_1828_v1) && (_1828_v1)));
          let _rhs282 = ((_1826_i0).plus(p1)).plus(((_1828_v1) ? (new BigNumber(578)) : (_1826_i0)));
          let _lhs249 = _1834_v7;
          let _lhs250 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_1834_v7).length));
          let _lhs251 = _1837_v8;
          let _lhs252 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_1837_v8).length));
          let _lhs253 = globalState;
          let _lhs254 = globalState;
          let _lhs255 = globalState;
          _lhs249[_lhs250] = _rhs278;
          _lhs251[_lhs252] = _rhs279;
          _lhs253.f16 = _rhs280;
          _lhs254.f15 = _rhs281;
          _lhs255.f12 = _rhs282;
          (globalState).f22 = (_1837_v8)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_1837_v8).length))];
          let _1843_v12;
          let _nw278 = new _module.C0();
          _nw278.__ctor(!((_1837_v8)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_1837_v8).length))]));
          _1843_v12 = _nw278;
        } else {
          let _1844_v13;
          _1844_v13 = true;
          (globalState).f22 = !((_1844_v13) && (_1844_v13));
          (globalState).f16 = p1;
          let _1845_v14;
          _1845_v14 = _dafny.Set.fromElements(_1827_v0);
          let _1846_v15;
          let _nw279 = new _module.C1();
          _nw279.__ctor(p3, _1826_i0, _1844_v13, (new BigNumber((p2).length)).isEqualTo(new BigNumber((_1845_v14).length)));
          _1846_v15 = _nw279;
          let _1847_v16;
          let _nw280 = Array((new BigNumber(21)).toNumber()).fill(false);
          _1847_v16 = _nw280;
          let _index263 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1847_v16).length));
          (_1847_v16)[_index263] = true;
          let _index264 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_1847_v16).length));
          (_1847_v16)[_index264] = !(_1844_v13);
          (globalState).f16 = _1846_v15.f37;
        }
        (globalState).f12 = _1826_i0;
        let _1848_v17;
        _1848_v17 = _dafny.Seq.of(_1826_i0, _module.__default.fm39(globalState));
        (globalState).f16 = new BigNumber((_1848_v17).length);
        (globalState).f19 = _1827_v0;
      }
      let _1849_v18;
      let _init47 = ((_1850_p1) => function (_1851_i3) {
        return (_1851_i3).multipliedBy(_1850_p1);
      })(p1);
      let _nw281 = Array((new BigNumber(27)).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw281.length); _i0_47++) {
        _nw281[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _1849_v18 = _nw281;
      let _1852_v19;
      _1852_v19 = _dafny.Seq.of(_1849_v18, _1849_v18, _1849_v18);
      let _1853_v20;
      _1853_v20 = _module.D3.create_DC10((_1852_v19)[_module.__default.safeIndex(p0, new BigNumber((_1852_v19).length))]);
      let _pat_let_tv46 = _1849_v18;
      let _1854_v21;
      _1854_v21 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let44_0) {
        return function (_1855_dt__update__tmp_h0) {
          return function (_pat_let45_0) {
            return function (_1856_dt__update_hcf21_h0) {
              return _module.D3.create_DC10(_1856_dt__update_hcf21_h0);
            }(_pat_let45_0);
          }(_pat_let_tv46);
        }(_pat_let44_0);
      }(_1853_v20),p1);
      let _1857_v22;
      _1857_v22 = _dafny.Set.fromElements(_1849_v18);
      let _1858_v23;
      _1858_v23 = true;
      let _1859_v24;
      let _nw282 = new _module.C4();
      _nw282.__ctor(_1854_v21, (_1857_v22).IsDisjointFrom(_dafny.Set.fromElements(_1849_v18, _1849_v18, _1849_v18, _1849_v18)), _1858_v23);
      _1859_v24 = _nw282;
      let _1860_v25;
      _1860_v25 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1861_v26;
      let _1862_v27;
      let _1863_v28;
      let _1864_v29;
      let _out57;
      let _out58;
      let _out59;
      let _out60;
      let _outcollector16 = _module.__default.m0(_1860_v25, globalState);
      _out57 = _outcollector16[0];
      _out58 = _outcollector16[1];
      _out59 = _outcollector16[2];
      _out60 = _outcollector16[3];
      _1861_v26 = _out57;
      _1862_v27 = _out58;
      _1863_v28 = _out59;
      _1864_v29 = _out60;
      let _1865_v30;
      _1865_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1858_v23,p1);
      let _hi16 = new BigNumber((_1865_v30).length);
      for (let _1866_i4 = (new BigNumber((p2).length)).minus(p1); _1866_i4.isLessThan(_hi16); _1866_i4 = _1866_i4.plus(_dafny.ONE)) {
        (globalState).f16 = _1866_i4;
        (globalState).f16 = (_1859_v24).fm1(globalState);
        _1862_v27 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1858_v23), _1862_v27), _1862_v27);
        let _1867_v31;
        _1867_v31 = _dafny.Seq.of(p3);
        let _1868_v32;
        _1868_v32 = _dafny.MultiSet.fromElements(_1867_v31);
        let _1869_v33;
        let _nw283 = new _module.C3();
        _nw283.__ctor(_1868_v32, _1863_v28, _1858_v23);
        _1869_v33 = _nw283;
        let _1870_v34;
        _1870_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1869_v33,(_this).f29);
        let _1871_v35;
        _1871_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1870_v34,(p1).isLessThan(new BigNumber(76)));
        _1871_v35 = (_1871_v35).update((_1870_v34).update(_1869_v33, (_this).f29), (_1862_v27)[_module.__default.safeIndex((_1869_v33).fm1(globalState), new BigNumber((_1862_v27).length))]);
      }
      let _1872_v36;
      _1872_v36 = _dafny.Set.fromElements(_1863_v28);
      let _rhs283 = ((_1872_v36).Intersect(_1872_v36)).IsProperSubsetOf(_1872_v36);
      let _rhs284 = p0;
      let _lhs256 = globalState;
      r0 = _rhs283;
      _lhs256.f16 = _rhs284;
      r0 = !(_1858_v23) || (_1858_v23);
      r1 = !((_1858_v23) || (_1858_v23));
      r2 = (_1852_v19)[_module.__default.safeIndex(p0, new BigNumber((_1852_v19).length))];
      let _1873_v38;
      let _nw284 = new _module.C7();
      _nw284.__ctor(_1858_v23, function () {
        let _coll68 = new _dafny.Map();
        for (const _compr_68 of _dafny.IntegerRange(new BigNumber(-329), new BigNumber(305))) {
          let _1874_v37 = _compr_68;
          if (((new BigNumber(-329)).isLessThanOrEqualTo(_1874_v37)) && ((_1874_v37).isLessThan(new BigNumber(305)))) {
            _coll68.push([(_1874_v37).plus(new BigNumber((_dafny.MultiSet.FromArray(_1862_v27)).cardinality())),_1865_v30]);
          }
        }
        return _coll68;
      }());
      _1873_v38 = _nw284;
      r3 = _1873_v38;
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1875_v0;
      _1875_v0 = new BigNumber(290);
      let _1876_v1;
      _1876_v1 = _dafny.Seq.of(_1875_v0);
      let _1877_v2;
      let _nw285 = new _module.C6();
      _nw285.__ctor(_1875_v0);
      _1877_v2 = _nw285;
      let _1878_v3;
      _1878_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1877_v2,_1876_v1);
      let _1879_v4;
      _1879_v4 = _dafny.MultiSet.fromElements(_1876_v1, (((_1878_v3).contains(_1877_v2)) ? ((_1878_v3).get(_1877_v2)) : (_1876_v1)));
      let _1880_v5;
      _1880_v5 = false;
      let _1881_v6;
      let _nw286 = new _module.C3();
      _nw286.__ctor(_1879_v4, false, _1880_v5);
      _1881_v6 = _nw286;
      let _1882_v7;
      _1882_v7 = _dafny.MultiSet.fromElements(_1881_v6, _1881_v6);
      let _1883_v8;
      _1883_v8 = _dafny.Seq.of(_1880_v5, false);
      let _1884_v9;
      _1884_v9 = _module.D6.create_DC17((((_1882_v7).contains(_1881_v6)) ? ((_1882_v7).get(_1881_v6)) : (_1877_v2.f40)), _1877_v2.f40, new BigNumber(51), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1883_v8, _module.__default.safeIndex(_1875_v0, new BigNumber((_1883_v8).length)), _1880_v5))).cardinality()), _1875_v0);
      let _1885_v10;
      _1885_v10 = new _dafny.CodePoint('v'.codePointAt(0));
      let _1886_v11;
      _1886_v11 = _dafny.MultiSet.fromElements(_1885_v10);
      let _1887_v12;
      _1887_v12 = _dafny.Seq.of(_1886_v11);
      let _pat_let_tv47 = _1881_v6;
      let _pat_let_tv48 = _1887_v12;
      let _pat_let_tv49 = _1885_v10;
      let _pat_let_tv50 = globalState;
      let _pat_let_tv51 = _1875_v0;
      let _pat_let_tv52 = _1877_v2;
      let _pat_let_tv53 = _1877_v2;
      let _pat_let_tv54 = _1875_v0;
      let _1888_v13;
      let _nw287 = Array((new BigNumber(23)).toNumber());
      _nw287[(_dafny.ZERO).toNumber()] = _1884_v9;
      _nw287[(_dafny.ONE).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(2)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(3)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(4)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(5)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(6)).toNumber()] = _module.D6.create_DC17(_1875_v0, _1875_v0, (_dafny.ZERO).minus(_1877_v2.f40), _1877_v2.f40, _1877_v2.f40);
      _nw287[(new BigNumber(7)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(8)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(9)).toNumber()] = _module.D6.create_DC17(new BigNumber(195), _1875_v0, _1877_v2.f40, _1875_v0, _1875_v0);
      _nw287[(new BigNumber(10)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(11)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(12)).toNumber()] = function (_pat_let46_0) {
        return function (_1889_dt__update__tmp_h0) {
          return function (_pat_let47_0) {
            return function (_1890_dt__update_hcf30_h0) {
              return function (_pat_let48_0) {
                return function (_1891_dt__update_hcf31_h0) {
                  return _module.D6.create_DC17(_1890_dt__update_hcf30_h0, _1891_dt__update_hcf31_h0, (_1889_dt__update__tmp_h0).dtor_cf32, (_1889_dt__update__tmp_h0).dtor_cf33, (_1889_dt__update__tmp_h0).dtor_cf34);
                }(_pat_let48_0);
              }(new BigNumber(87));
            }(_pat_let47_0);
          }((_pat_let_tv47).fm13(_pat_let_tv48, _pat_let_tv49, new BigNumber(((_this).f29).cardinality()), _pat_let_tv50));
        }(_pat_let46_0);
      }(_1884_v9);
      _nw287[(new BigNumber(13)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(14)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(15)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(16)).toNumber()] = function (_pat_let49_0) {
        return function (_1892_dt__update__tmp_h1) {
          return function (_pat_let50_0) {
            return function (_1893_dt__update_hcf34_h0) {
              return function (_pat_let51_0) {
                return function (_1894_dt__update_hcf33_h0) {
                  return _module.D6.create_DC17((_1892_dt__update__tmp_h1).dtor_cf30, (_1892_dt__update__tmp_h1).dtor_cf31, (_1892_dt__update__tmp_h1).dtor_cf32, _1894_dt__update_hcf33_h0, _1893_dt__update_hcf34_h0);
                }(_pat_let51_0);
              }((_dafny.ZERO).minus(_pat_let_tv52.f40));
            }(_pat_let50_0);
          }(_pat_let_tv51);
        }(_pat_let49_0);
      }(_module.D6.create_DC17(_1875_v0, new BigNumber((_module.__default.fm51(_1875_v0, _1875_v0, _1880_v5, new BigNumber(597), globalState)).length), _1877_v2.f40, _1875_v0, _1877_v2.f40));
      _nw287[(new BigNumber(17)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(18)).toNumber()] = _module.D6.create_DC17(_1877_v2.f40, (_1881_v6).fm13(_1887_v12, _1885_v10, _1877_v2.f40, globalState), _1875_v0, new BigNumber(((_this).f29).cardinality()), _1875_v0);
      _nw287[(new BigNumber(19)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(20)).toNumber()] = function (_pat_let52_0) {
        return function (_1895_dt__update__tmp_h2) {
          return function (_pat_let53_0) {
            return function (_1896_dt__update_hcf33_h1) {
              return function (_pat_let54_0) {
                return function (_1897_dt__update_hcf30_h1) {
                  return _module.D6.create_DC17(_1897_dt__update_hcf30_h1, (_1895_dt__update__tmp_h2).dtor_cf31, (_1895_dt__update__tmp_h2).dtor_cf32, _1896_dt__update_hcf33_h1, (_1895_dt__update__tmp_h2).dtor_cf34);
                }(_pat_let54_0);
              }(_pat_let_tv54);
            }(_pat_let53_0);
          }(_pat_let_tv53.f40);
        }(_pat_let52_0);
      }(_1884_v9);
      _nw287[(new BigNumber(21)).toNumber()] = _1884_v9;
      _nw287[(new BigNumber(22)).toNumber()] = _1884_v9;
      _1888_v13 = _nw287;
      let _index265 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_1888_v13).length));
      (_1888_v13)[_index265] = _module.D6.create_DC17(_1875_v0, _1875_v0, _1875_v0, new BigNumber(66), _1875_v0);
      let _index266 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_1888_v13).length));
      (_1888_v13)[_index266] = _module.D6.create_DC17((new BigNumber((_dafny.Seq.of(_1885_v10, _1885_v10)).length)).plus(_1877_v2.f40), _1877_v2.f40, _1877_v2.f40, _1875_v0, (_dafny.ZERO).minus(_1877_v2.f40));
      (globalState).f12 = _1875_v0;
      let _1898_v14;
      _1898_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_1880_v5),_1875_v0);
      let _1899_v15;
      _1899_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(_1880_v5),_1898_v14);
      _1899_v15 = (_1899_v15).update(_1880_v5, (_dafny.Map.Empty.slice().updateUnsafe(_1880_v5,_1875_v0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1880_v5,_1875_v0)));
      if (!(!(_dafny.Seq.IsPrefixOf(_1883_v8, _dafny.Seq.of(_1880_v5))))) {
        let _1900_v16;
        _1900_v16 = _dafny.Seq.UnicodeFromString("kyvdj");
        let _1901_v17;
        _1901_v17 = _dafny.MultiSet.fromElements(_1880_v5);
        let _1902_v18;
        let _nw288 = Array((new BigNumber(19)).toNumber());
        _nw288[(_dafny.ZERO).toNumber()] = _1877_v2.f40;
        _nw288[(_dafny.ONE).toNumber()] = _1875_v0;
        _nw288[(new BigNumber(2)).toNumber()] = new BigNumber(158);
        _nw288[(new BigNumber(3)).toNumber()] = _1877_v2.f40;
        _nw288[(new BigNumber(4)).toNumber()] = (_1877_v2).fm1(globalState);
        _nw288[(new BigNumber(5)).toNumber()] = _module.__default.fm39(globalState);
        _nw288[(new BigNumber(6)).toNumber()] = new BigNumber((_1900_v16).length);
        _nw288[(new BigNumber(7)).toNumber()] = _1877_v2.f40;
        _nw288[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(_1877_v2.f40, new BigNumber((_1901_v17).cardinality()));
        _nw288[(new BigNumber(9)).toNumber()] = (_1877_v2.f40).plus(_1877_v2.f40);
        _nw288[(new BigNumber(10)).toNumber()] = ((_1877_v2).fm1(globalState)).multipliedBy(_1875_v0);
        _nw288[(new BigNumber(11)).toNumber()] = _1875_v0;
        _nw288[(new BigNumber(12)).toNumber()] = _1877_v2.f40;
        _nw288[(new BigNumber(13)).toNumber()] = (_1877_v2.f40).plus(new BigNumber((_1898_v14).length));
        _nw288[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt(_1875_v0, _1875_v0);
        _nw288[(new BigNumber(15)).toNumber()] = new BigNumber(-336);
        _nw288[(new BigNumber(16)).toNumber()] = new BigNumber(((_1901_v17).Union(_dafny.MultiSet.fromElements(_1880_v5))).cardinality());
        _nw288[(new BigNumber(17)).toNumber()] = new BigNumber(416);
        _nw288[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt((_1876_v1)[_module.__default.safeIndex(new BigNumber(((_this).f29).cardinality()), new BigNumber((_1876_v1).length))], _1875_v0);
        _1902_v18 = _nw288;
        _1902_v18 = _1902_v18;
        let _1903_v19;
        _1903_v19 = _dafny.Map.Empty.slice().updateUnsafe((_1881_v6).fm13(_1887_v12, _1885_v10, _1877_v2.f40, globalState),_1900_v16);
        let _1904_v20;
        _1904_v20 = _dafny.Set.fromElements(_1875_v0, _1877_v2.f40);
        let _1905_v21;
        _1905_v21 = _dafny.Map.Empty.slice().updateUnsafe((((_1903_v19).contains(new BigNumber((_1904_v20).length))) ? ((_1903_v19).get(new BigNumber((_1904_v20).length))) : (_dafny.Seq.UnicodeFromString("xicujy"))),((_1880_v5) ? (new BigNumber((_dafny.MultiSet.FromArray(_1900_v16)).cardinality())) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_1877_v2.f40)))));
        _1905_v21 = (_1905_v21).update(_1900_v16, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("akwbgonr"), _1900_v16)).length));
        let _1906_v22;
        _1906_v22 = _module.D3.create_DC10(_1902_v18);
        let _1907_v23;
        _1907_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1880_v5,(_1906_v22).dtor_cf21);
        let _1908_v24;
        _1908_v24 = _dafny.Map.Empty.slice().updateUnsafe((((_1907_v23).contains(_1880_v5)) ? ((_1907_v23).get(_1880_v5)) : (_1902_v18)),!_dafny.Seq.contains(_dafny.Seq.update(_1900_v16, _module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1900_v16).length)), (_1900_v16)[_module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1900_v16).length))]), _module.__default.fm40(new BigNumber(373), _1880_v5, true, globalState)));
        _1908_v24 = (_1908_v24).update(_1902_v18, (_1877_v2).fm44(globalState));
        let _1909_v25;
        _1909_v25 = _module.D14.create_DC32(_1901_v17, (_1883_v8)[_module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1883_v8).length))], _1880_v5, _1880_v5, _1877_v2.f40);
        let _1910_v26;
        _1910_v26 = _module.D3.create_DC11(_1885_v10);
        let _1911_v27;
        let _nw289 = Array((new BigNumber(27)).toNumber());
        _nw289[(_dafny.ZERO).toNumber()] = _1885_v10;
        _nw289[(_dafny.ONE).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(2)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(3)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(4)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(5)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(6)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(7)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(8)).toNumber()] = (_1910_v26).dtor_cf22;
        _nw289[(new BigNumber(9)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(10)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(11)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(12)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(13)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(14)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(15)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
        _nw289[(new BigNumber(17)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(18)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(19)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(20)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(21)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(22)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(23)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(24)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(25)).toNumber()] = _1885_v10;
        _nw289[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _1911_v27 = _nw289;
        let _1912_v28;
        _1912_v28 = _module.D1.create_DC4(_1911_v27, _1877_v2.f40, _1877_v2.f40, _1900_v16, _1900_v16);
        let _1913_v29;
        _1913_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1875_v0,new BigNumber((_1904_v20).length));
        let _1914_v30;
        _1914_v30 = _dafny.Seq.of(_dafny.Set.fromElements((((_1913_v29).contains(_1877_v2.f40)) ? ((_1913_v29).get(_1877_v2.f40)) : (new BigNumber((_1913_v29).length))), _1877_v2.f40));
        let _1915_v31;
        let _nw290 = Array((new BigNumber(22)).toNumber());
        _nw290[(_dafny.ZERO).toNumber()] = _1880_v5;
        _nw290[(_dafny.ONE).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(2)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(3)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(4)).toNumber()] = false;
        _nw290[(new BigNumber(5)).toNumber()] = !(_1880_v5);
        _nw290[(new BigNumber(6)).toNumber()] = (((_1909_v25).dtor_cf58) ? (_1880_v5) : (false));
        _nw290[(new BigNumber(7)).toNumber()] = false;
        _nw290[(new BigNumber(8)).toNumber()] = _dafny.Seq.IsProperPrefixOf((_1912_v28).dtor_cf10, _dafny.Seq.UnicodeFromString("kyjjbux"));
        _nw290[(new BigNumber(9)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(10)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(11)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(12)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(13)).toNumber()] = (_1880_v5) || (_1880_v5);
        _nw290[(new BigNumber(14)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(15)).toNumber()] = (_1904_v20).IsSubsetOf((_1914_v30)[_module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1914_v30).length))]);
        _nw290[(new BigNumber(16)).toNumber()] = (!(true)) && (_1880_v5);
        _nw290[(new BigNumber(17)).toNumber()] = _1880_v5;
        _nw290[(new BigNumber(18)).toNumber()] = (_1880_v5) && (_1880_v5);
        _nw290[(new BigNumber(19)).toNumber()] = false;
        _nw290[(new BigNumber(20)).toNumber()] = false;
        _nw290[(new BigNumber(21)).toNumber()] = !(_1880_v5);
        _1915_v31 = _nw290;
        let _index267 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_1915_v31).length));
        (_1915_v31)[_index267] = _1880_v5;
        let _index268 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_1915_v31).length));
        (_1915_v31)[_index268] = !(!(_1877_v2.f40).isEqualTo(((_1880_v5) ? ((_dafny.ZERO).minus(_1875_v0)) : (_1875_v0))));
        let _index269 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1902_v18).length));
        (_1902_v18)[_index269] = _1877_v2.f40;
        let _index270 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1902_v18).length));
        (_1902_v18)[_index270] = _1877_v2.f40;
      } else {
        r1 = ((((_1882_v7).contains(_1881_v6)) ? ((_1882_v7).get(_1881_v6)) : (_1875_v0))).isLessThan(_1875_v0);
        let _1916_v32;
        _1916_v32 = _dafny.Seq.UnicodeFromString("vaiuecm");
        (globalState).f6 = _1916_v32;
        let _1917_v33;
        let _nw291 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _1917_v33 = _nw291;
        let _index271 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1917_v33).length));
        (_1917_v33)[_index271] = new BigNumber(-728);
        let _1918_v34;
        _1918_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(480),_dafny.Map.Empty.slice().updateUnsafe(_1880_v5,_1875_v0));
        let _1919_v35;
        let _nw292 = new _module.C7();
        _nw292.__ctor(_1880_v5, _1918_v34);
        _1919_v35 = _nw292;
        let _1920_v36;
        _1920_v36 = _dafny.Set.fromElements(_1919_v35);
        let _index272 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_1917_v33).length));
        (_1917_v33)[_index272] = new BigNumber((((!((_1880_v5) && (_1880_v5))) ? ((_1920_v36).Union(_1920_v36)) : ((_1920_v36).Difference(_1920_v36)))).length);
        let _1921_v37;
        _1921_v37 = _module.D3.create_DC10(_1917_v33);
        let _1922_v38;
        _1922_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v37,_1877_v2.f40);
        let _1923_v39;
        _1923_v39 = _dafny.MultiSet.fromElements(_1880_v5, _1880_v5);
        let _1924_v40;
        _1924_v40 = _module.D14.create_DC32(_1923_v39, _1880_v5, !(_1880_v5), _1880_v5, _1877_v2.f40);
        let _1925_v41;
        let _nw293 = new _module.C1();
        _nw293.__ctor(new BigNumber(((_this).f29).cardinality()), new BigNumber((_module.__default.fm52(_1880_v5, (_1917_v33)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1917_v33).length))], globalState)).length), (_1924_v40).dtor_cf57, _1880_v5);
        _1925_v41 = _nw293;
        let _1926_v42;
        _1926_v42 = _dafny.Seq.of(_1925_v41);
        let _1927_v43;
        let _nw294 = new _module.C4();
        _nw294.__ctor(_1922_v38, _1880_v5, !_dafny.areEqual(_dafny.Seq.update(_1926_v42, _module.__default.safeIndex((_1925_v41).f36, new BigNumber((_1926_v42).length)), _1925_v41), _dafny.Seq.of(_1925_v41, _1925_v41)));
        _1927_v43 = _nw294;
        let _1928_v44;
        _1928_v44 = _dafny.Seq.of(_1916_v32);
        let _1929_v45;
        let _nw295 = Array((new BigNumber(25)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = _1916_v32;
        _nw295[(_dafny.ONE).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(2)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(3)).toNumber()] = (_1928_v44)[_module.__default.safeIndex((_1917_v33)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_1917_v33).length))], new BigNumber((_1928_v44).length))];
        _nw295[(new BigNumber(4)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(5)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(6)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("hiqyfxr");
        _nw295[(new BigNumber(8)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_1916_v32, _module.__default.safeIndex(new BigNumber(-400), new BigNumber((_1916_v32).length)), _1885_v10), _dafny.Seq.UnicodeFromString("mfc"));
        _nw295[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("bsbqonwkb");
        _nw295[(new BigNumber(11)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(12)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(13)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(14)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_1916_v32, _module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1916_v32).length)), _1885_v10);
        _nw295[(new BigNumber(16)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(17)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(18)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(19)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("kasxjomfh");
        _nw295[(new BigNumber(21)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(22)).toNumber()] = (_1928_v44)[_module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1928_v44).length))];
        _nw295[(new BigNumber(23)).toNumber()] = _1916_v32;
        _nw295[(new BigNumber(24)).toNumber()] = _1916_v32;
        _1929_v45 = _nw295;
        let _index273 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1929_v45).length));
        (_1929_v45)[_index273] = _dafny.Seq.Concat(_1916_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(283)), function (_1930_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }));
        let _index274 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_1929_v45).length));
        (_1929_v45)[_index274] = _1916_v32;
      }
      let _1931_v46;
      _1931_v46 = _dafny.Seq.UnicodeFromString("tgaaxgqsq");
      let _1932_v47;
      _1932_v47 = _module.D6.create_DC16(_dafny.Seq.Concat(_1931_v46, _1931_v46));
      let _pat_let_tv55 = _1931_v46;
      _1932_v47 = ((_1880_v5) ? (_1932_v47) : (function (_pat_let55_0) {
        return function (_1933_dt__update__tmp_h3) {
          return function (_pat_let56_0) {
            return function (_1934_dt__update_hcf29_h0) {
              return _module.D6.create_DC16(_1934_dt__update_hcf29_h0);
            }(_pat_let56_0);
          }(_pat_let_tv55);
        }(_pat_let55_0);
      }(_1932_v47)));
      let _hi17 = (_1877_v2.f40).minus(_1875_v0);
      for (let _1935_i1 = (new BigNumber(-597)).multipliedBy(_1875_v0); _1935_i1.isLessThan(_hi17); _1935_i1 = _1935_i1.plus(_dafny.ONE)) {
        _1885_v10 = _1885_v10;
        let _1936_v48;
        let _nw296 = new _module.C0();
        _nw296.__ctor(true);
        _1936_v48 = _nw296;
        (globalState).f22 = !(((_1880_v5) ? (_1880_v5) : ((_1936_v48).f32)));
        let _1937_v49;
        _1937_v49 = _module.D0.create_DC1(true, (_1936_v48).f32, _1877_v2.f40, (_1936_v48).f32);
        let _1938_v50;
        _1938_v50 = _dafny.Set.fromElements(_1937_v49);
        _1938_v50 = (_1938_v50).Difference((_1938_v50).Union(_dafny.Set.fromElements(_1937_v49)));
      }
      r0 = _1880_v5;
      r1 = (_1883_v8)[_module.__default.safeIndex(_1877_v2.f40, new BigNumber((_1883_v8).length))];
      return [r0, r1];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
      this._f26 = false;
      this._f27 = false;
      this._f29 = _dafny.MultiSet.Empty;
      this.f30 = false;
      this._f31 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T2, _module.T3, _module.T1, _module.T0];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
    __ctor(f30, f31, f26, f27, f29) {
      let _this = this;
      (_this).f30 = f30;
      (_this)._f31 = f31;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      (_this)._f29 = f29;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_this).f26, (_this).f26), _dafny.Seq.of((_this).f27)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f26), _dafny.Seq.of((_this).f27)));
    };
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return ((_this).f31).minus((_this).f31);
    };
    fm1(globalState) {
      let _this = this;
      return (_this).f31;
    };
    fm2(p0, p1, globalState) {
      let _this = this;
      return ((_this).f27) === ((_this).f27);
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let _1939_v0;
      let _nw297 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
      _1939_v0 = _nw297;
      let _1940_v1;
      _1940_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_this.f30);
      let _index275 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length));
      (_1939_v0)[_index275] = _1940_v1;
      let _index276 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length));
      let _rhs285 = ((_this.f30) ? ((_dafny.ZERO).minus((_this).f31)) : (_module.__default.safeDivisionInt((_this).f31, (_this).f31)));
      let _rhs286 = (_1940_v1).Merge((_1940_v1).Merge(_1940_v1));
      let _lhs257 = globalState;
      let _lhs258 = _1939_v0;
      let _lhs259 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length));
      _lhs257.f12 = _rhs285;
      _lhs258[_lhs259] = _rhs286;
      let _1941_v2;
      _1941_v2 = _dafny.Seq.of(((((_1939_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length))]).contains((_this).f27)) ? (((_1939_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length))]).get((_this).f27)) : (p0)), !(_this.f30), (_this).fm2(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f31)).length), globalState));
      let _1942_v3;
      _1942_v3 = new _dafny.CodePoint('v'.codePointAt(0));
      let _1943_v4;
      _1943_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1941_v2).length),_1942_v3);
      let _1944_v5;
      _1944_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_this.f30);
      let _1945_v6;
      _1945_v6 = _dafny.Set.fromElements((((_1943_v4).contains(new BigNumber((_1944_v5).length))) ? ((_1943_v4).get(new BigNumber((_1944_v5).length))) : (_1942_v3)), _1942_v3);
      let _1946_i0;
      _1946_i0 = _dafny.ZERO;
      L10: {
        while ((_1945_v6).IsDisjointFrom(_1945_v6)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1946_i0)) {
              break L10;
            }
            _1946_i0 = (_1946_i0).plus(_dafny.ONE);
            let _1947_v7;
            _1947_v7 = _module.D0.create_DC0(p2);
            let _1948_v8;
            _1948_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f31);
            let _pat_let_tv56 = globalState;
            let _pat_let_tv57 = _1948_v8;
            let _source27 = function (_pat_let57_0) {
              return function (_1949_dt__update__tmp_h0) {
                return function (_pat_let58_0) {
                  return function (_1950_dt__update_hcf0_h0) {
                    return _module.D0.create_DC0(_1950_dt__update_hcf0_h0);
                  }(_pat_let58_0);
                }(!(_dafny.Map.Empty.slice().updateUnsafe((_this).fm1(_pat_let_tv56),(_this).f31)).equals(_pat_let_tv57));
              }(_pat_let57_0);
            }(_1947_v7);
            if (_source27.is_DC1) {
              let _1951___mcc_h0 = (_source27).cf1;
              let _1952___mcc_h1 = (_source27).cf2;
              let _1953___mcc_h2 = (_source27).cf3;
              let _1954___mcc_h3 = (_source27).cf4;
              let _1955_cf4 = _1954___mcc_h3;
              let _1956_cf3 = _1953___mcc_h2;
              let _1957_cf2 = _1952___mcc_h1;
              let _1958_cf1 = _1951___mcc_h0;
              (globalState).f22 = !((_1947_v7).dtor_cf0);
              _1943_v4 = _1943_v4;
              let _1959_v9;
              let _init48 = ((_1960_p2) => function (_1961_i1) {
                return _1960_p2;
              })(p2);
              let _nw298 = Array((new BigNumber(14)).toNumber());
              for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw298.length); _i0_48++) {
                _nw298[_i0_48] = _init48(new BigNumber(_i0_48));
              }
              _1959_v9 = _nw298;
              let _index277 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1959_v9).length));
              (_1959_v9)[_index277] = ((true) ? ((_this).f26) : (_this.f30));
              let _index278 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1959_v9).length));
              (_1959_v9)[_index278] = _1955_cf4;
              (globalState).f12 = ((((_this).f29).contains(_1956_cf3)) ? (((_this).f29).get(_1956_cf3)) : ((_this).f31));
            } else if (_source27.is_DC0) {
              let _1962___mcc_h4 = (_source27).cf0;
              let _1963_cf0 = _1962___mcc_h4;
              let _1964_v10;
              let _init49 = ((_1965_p0) => function (_1966_i2) {
                return _1965_p0;
              })(p0);
              let _nw299 = Array((new BigNumber(25)).toNumber());
              for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw299.length); _i0_49++) {
                _nw299[_i0_49] = _init49(new BigNumber(_i0_49));
              }
              _1964_v10 = _nw299;
              let _1967_v11;
              _1967_v11 = _dafny.Seq.of(_1964_v10, _1964_v10, _1964_v10, _1964_v10, _1964_v10);
              (globalState).f12 = (((p0) ? (new BigNumber(692)) : ((_this).f31))).minus(new BigNumber((_dafny.Seq.Concat(_1967_v11, _1967_v11)).length));
              let _1968_v12;
              _1968_v12 = _dafny.Seq.UnicodeFromString("bex");
              let _1969_v13;
              let _nw300 = Array((new BigNumber(24)).toNumber());
              _nw300[(_dafny.ZERO).toNumber()] = (_this).f31;
              _nw300[(_dafny.ONE).toNumber()] = new BigNumber(-773);
              _nw300[(new BigNumber(2)).toNumber()] = new BigNumber((_1968_v12).length);
              _nw300[(new BigNumber(3)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(4)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(5)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(6)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(7)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(8)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(9)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(10)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(11)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(12)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(13)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(14)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(15)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(16)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(17)).toNumber()] = new BigNumber((_1941_v2).length);
              _nw300[(new BigNumber(18)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(19)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(20)).toNumber()] = new BigNumber((_1968_v12).length);
              _nw300[(new BigNumber(21)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(22)).toNumber()] = (_this).f31;
              _nw300[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.of(_1941_v2, _1941_v2)).length);
              _1969_v13 = _nw300;
              let _1970_v14;
              _1970_v14 = _module.D3.create_DC10(_1969_v13);
              let _1971_v15;
              _1971_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1970_v14,(_this).f31);
              let _1972_v16;
              let _nw301 = new _module.C4();
              _nw301.__ctor(_1971_v15, _1963_cf0, !((_this).f26));
              _1972_v16 = _nw301;
              (globalState).f2 = true;
              let _index279 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1969_v13).length));
              (_1969_v13)[_index279] = (_this).f31;
              let _index280 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1969_v13).length));
              (_1969_v13)[_index280] = (_this).f31;
            } else {
              let _1973___mcc_h5 = (_source27).cf5;
              let _1974_cf5 = _1973___mcc_h5;
              let _1975_v17;
              let _nw302 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1975_v17 = _nw302;
              let _1976_v18;
              _1976_v18 = _dafny.Seq.UnicodeFromString("dbooou");
              let _index281 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1975_v17).length));
              (_1975_v17)[_index281] = _1976_v18;
              let _1977_v19;
              _1977_v19 = _dafny.Set.fromElements(_this.f30);
              let _index282 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_1975_v17).length));
              (_1975_v17)[_index282] = (((_1977_v19).IsProperSubsetOf(_1977_v19)) ? (_1976_v18) : (_dafny.Seq.Concat(_1976_v18, _1976_v18)));
              let _1978_v20;
              _1978_v20 = _dafny.Set.fromElements((_this).f31, (_this).f31);
              let _1979_v21;
              let _nw303 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1979_v21 = _nw303;
              let _1980_v22;
              _1980_v22 = _module.D1.create_DC4(_1979_v21, new BigNumber(731), new BigNumber(279), _dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_1981_v3) => function (_1982_i3) {
  return _1981_v3;
})(_1942_v3)), _dafny.Seq.UnicodeFromString("l"));
              (globalState).f12 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_1976_v18).length), new BigNumber((_1978_v20).length))), (_1980_v22).dtor_cf8);
              let _1983_v23;
              _1983_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f31);
              (globalState).f2 = (((_this).f27) ? (false) : ((p1).fm2((_this).fm2(_this.f30, new BigNumber((_1983_v23).length), globalState), (_this).f31, globalState)));
              (globalState).f2 = (_this).fm2(((_this).f31).isLessThanOrEqualTo((_this).f31), ((_this).f31).plus(new BigNumber(111)), globalState);
            }
            let _1984_v24;
            _1984_v24 = _dafny.Seq.of(_1942_v3, _1942_v3);
            let _1985_v25;
            _1985_v25 = _dafny.Seq.of(_1984_v24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), ((_1986_v3) => function (_1987_i4) {
              return _1986_v3;
            })(_1942_v3)));
            _1985_v25 = _1985_v25;
            if (_this.f30) {
              let _1988_v26;
              let _nw304 = Array((new BigNumber(4)).toNumber()).fill(false);
              _1988_v26 = _nw304;
              let _index283 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1988_v26).length));
              (_1988_v26)[_index283] = !((p1).fm2((_this).f26, (_this).f31, globalState));
              let _index284 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1988_v26).length));
              (_1988_v26)[_index284] = (_this).fm2((_this).f26, (_this).f31, globalState);
              let _1989_v27;
              _1989_v27 = _dafny.Seq.of(_1988_v26);
              let _1990_v28;
              let _init50 = function (_1991_i5) {
                return (_1991_i5).plus((_this).f31);
              };
              let _nw305 = Array((new BigNumber(17)).toNumber());
              for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw305.length); _i0_50++) {
                _nw305[_i0_50] = _init50(new BigNumber(_i0_50));
              }
              _1990_v28 = _nw305;
              let _rhs287 = _1989_v27;
              let _rhs288 = (p1).fm2((_1988_v26)[_module.__default.safeIndex(new BigNumber(357), new BigNumber((_1988_v26).length))], (_this).f31, globalState);
              let _rhs289 = _1990_v28;
              let _rhs290 = _module.__default.safeModuloInt((_this).f31, (_this).f31);
              let _lhs260 = globalState;
              let _lhs261 = globalState;
              _1989_v27 = _rhs287;
              _lhs260.f24 = _rhs288;
              _1990_v28 = _rhs289;
              _lhs261.f16 = _rhs290;
              let _1992_v29;
              _1992_v29 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f31)).length));
              let _index285 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1990_v28).length));
              (_1990_v28)[_index285] = ((p0) ? ((_this).f31) : (new BigNumber((_1992_v29).length)));
              let _1993_v30;
              _1993_v30 = _dafny.Seq.of(_1989_v27, _1989_v27);
              let _index286 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1990_v28).length));
              let _rhs291 = (_dafny.ZERO).minus((_this).f31);
              let _rhs292 = (p1).fm2(_dafny.Seq.contains((_1993_v30)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1993_v30).length))], _1988_v26), new BigNumber(22), globalState);
              let _rhs293 = (_this).f31;
              let _rhs294 = _this.f30;
              let _lhs262 = globalState;
              let _lhs263 = _this;
              let _lhs264 = _1990_v28;
              let _lhs265 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1990_v28).length));
              let _lhs266 = globalState;
              _lhs262.f12 = _rhs291;
              _lhs263.f30 = _rhs292;
              _lhs264[_lhs265] = _rhs293;
              _lhs266.f15 = _rhs294;
              _1988_v26 = _1988_v26;
              let _1994_v31;
              _1994_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(_1990_v28),(_1990_v28)[_module.__default.safeIndex(new BigNumber(292), new BigNumber((_1990_v28).length))]);
              let _1995_v32;
              let _nw306 = new _module.C4();
              _nw306.__ctor(_1994_v31, p2, !(p0));
              _1995_v32 = _nw306;
            } else {
              let _1996_v33;
              _1996_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f31);
              let _1997_v34;
              let _1998_v35;
              let _1999_v36;
              let _out61;
              let _out62;
              let _out63;
              let _outcollector17 = (_this).m6(_1996_v33, (_dafny.ZERO).minus((_this).f31), _1942_v3, (_this).f31, globalState);
              _out61 = _outcollector17[0];
              _out62 = _outcollector17[1];
              _out63 = _outcollector17[2];
              _1997_v34 = _out61;
              _1998_v35 = _out62;
              _1999_v36 = _out63;
              _1942_v3 = new _dafny.CodePoint('o'.codePointAt(0));
              let _2000_v37;
              let _init51 = function (_2001_i6) {
                return (_this).f26;
              };
              let _nw307 = Array((new BigNumber(3)).toNumber());
              for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw307.length); _i0_51++) {
                _nw307[_i0_51] = _init51(new BigNumber(_i0_51));
              }
              _2000_v37 = _nw307;
              _2000_v37 = _2000_v37;
              let _2002_v38;
              _2002_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
              let _2003_v39;
              _2003_v39 = _module.D12.create_DC27(p1);
              (globalState).f24 = (((_2002_v38).contains((_2003_v39).dtor_cf48)) ? ((_2002_v38).get((_2003_v39).dtor_cf48)) : (p0));
              (globalState).f24 = (_this).f26;
            }
            let _2004_v40;
            _2004_v40 = _dafny.Set.fromElements(_1940_v1, ((_1939_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1939_v0).length))]).update(p0, (_this).f26), _1940_v1);
            _2004_v40 = (_2004_v40).Union(_2004_v40);
          }
        }
      }
      let _2005_v41;
      _2005_v41 = _module.D12.create_DC28(new BigNumber(229));
      let _source28 = _2005_v41;
      if (_source28.is_DC28) {
        let _2006___mcc_h6 = (_source28).cf49;
        let _2007_cf49 = _2006___mcc_h6;
        (globalState).f22 = ((_this.f30) ? (((_this.f30) ? (_this.f30) : (p2))) : (_this.f30));
        (globalState).f2 = (_1941_v2)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_1941_v2).length))];
        let _2008_v42;
        let _nw308 = new _module.C5();
        _nw308.__ctor(false, false);
        _2008_v42 = _nw308;
        (globalState).f12 = (_this).f31;
      } else {
        let _2009___mcc_h7 = (_source28).cf48;
        let _2010_cf48 = _2009___mcc_h7;
        let _2011_v43;
        let _init52 = function (_2012_i7) {
          return (_2012_i7).plus((_this).f31);
        };
        let _nw309 = Array((new BigNumber(26)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw309.length); _i0_52++) {
          _nw309[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _2011_v43 = _nw309;
        let _2013_v44;
        _2013_v44 = _module.D3.create_DC10(_2011_v43);
        let _2014_v45;
        _2014_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2013_v44,new BigNumber((_dafny.Seq.of(_this.f30)).length));
        let _2015_v46;
        let _nw310 = new _module.C4();
        _nw310.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_2013_v44,(_this).f31)).Merge(_2014_v45), p2, ((p2) ? (p0) : ((_this).f27)));
        _2015_v46 = _nw310;
        if ((_this).f26) {
          let _2016_v47;
          _2016_v47 = _dafny.Set.fromElements((_2015_v46).fm1(globalState), (_this).f31);
          _2016_v47 = _2016_v47;
          let _2017_v48;
          _2017_v48 = _dafny.Seq.UnicodeFromString("vsjg");
          (globalState).f6 = _2017_v48;
          let _2018_v49;
          _2018_v49 = _dafny.Set.fromElements(_2017_v48, _2017_v48, _2017_v48);
          let _2019_v50;
          _2019_v50 = _dafny.MultiSet.fromElements(_1942_v3);
          let _2020_v51;
          _2020_v51 = _module.D0.create_DC0(p2);
          let _2021_v52;
          _2021_v52 = _dafny.Set.fromElements((_1941_v2)[_module.__default.safeIndex((_this).f31, new BigNumber((_1941_v2).length))]);
          let _2022_v53;
          let _nw311 = new _module.C1();
          _nw311.__ctor(new BigNumber((_2017_v48).length), new BigNumber(800), (_this).f27, p0);
          _2022_v53 = _nw311;
          let _2023_v54;
          _2023_v54 = _dafny.Seq.of(_2022_v53, _2022_v53);
          let _2024_v55;
          _2024_v55 = _dafny.Map.Empty.slice().updateUnsafe(_2023_v54,(_this).f31);
          let _2025_v56;
          let _nw312 = Array((new BigNumber(25)).toNumber());
          _nw312[(_dafny.ZERO).toNumber()] = (_2018_v49).IsProperSubsetOf(_2018_v49);
          _nw312[(_dafny.ONE).toNumber()] = (_this).f27;
          _nw312[(new BigNumber(2)).toNumber()] = (_this).f26;
          _nw312[(new BigNumber(3)).toNumber()] = (_2019_v50).IsSubsetOf(_2019_v50);
          _nw312[(new BigNumber(4)).toNumber()] = (_2020_v51).dtor_cf0;
          _nw312[(new BigNumber(5)).toNumber()] = ((_this).f27) || (_this.f30);
          _nw312[(new BigNumber(6)).toNumber()] = ((_this).f31).isLessThanOrEqualTo((_this).f31);
          _nw312[(new BigNumber(7)).toNumber()] = p0;
          _nw312[(new BigNumber(8)).toNumber()] = !((_this).f26);
          _nw312[(new BigNumber(9)).toNumber()] = !((_dafny.Set.fromElements(_this.f30, p2)).IsProperSubsetOf(_2021_v52));
          _nw312[(new BigNumber(10)).toNumber()] = (_2010_cf48).fm2((_this).f27, (_this).f31, globalState);
          _nw312[(new BigNumber(11)).toNumber()] = !(!((_this.f30) || ((_this).f27)));
          _nw312[(new BigNumber(12)).toNumber()] = p0;
          _nw312[(new BigNumber(13)).toNumber()] = p0;
          _nw312[(new BigNumber(14)).toNumber()] = _this.f30;
          _nw312[(new BigNumber(15)).toNumber()] = (_this).fm2((_this).f27, new BigNumber((_2024_v55).length), globalState);
          _nw312[(new BigNumber(16)).toNumber()] = (_2022_v53).f26;
          _nw312[(new BigNumber(17)).toNumber()] = !(((_this).f27) === ((_this).f27));
          _nw312[(new BigNumber(18)).toNumber()] = !((_dafny.Set.fromElements(new BigNumber(379))).IsDisjointFrom(_2016_v47));
          _nw312[(new BigNumber(19)).toNumber()] = p2;
          _nw312[(new BigNumber(20)).toNumber()] = !((_this).fm2((_2022_v53).f26, (_this).f31, globalState)) || ((_this).f27);
          _nw312[(new BigNumber(21)).toNumber()] = (_this).f27;
          _nw312[(new BigNumber(22)).toNumber()] = _this.f30;
          _nw312[(new BigNumber(23)).toNumber()] = (_2022_v53).f26;
          _nw312[(new BigNumber(24)).toNumber()] = (_1941_v2)[_module.__default.safeIndex((_this).f31, new BigNumber((_1941_v2).length))];
          _2025_v56 = _nw312;
          let _index287 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_2025_v56).length));
          (_2025_v56)[_index287] = !((_2015_v46).fm2((_2022_v53).f27, (_this).f31, globalState));
          let _index288 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_2025_v56).length));
          (_2025_v56)[_index288] = (_2022_v53).f26;
          (globalState).f16 = new BigNumber((_dafny.Seq.Concat(_2017_v48, _dafny.Seq.Concat(_2017_v48, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-868)), ((_2026_v3) => function (_2027_i8) {
            return _2026_v3;
          })(_1942_v3))))).length);
          let _2028_v57;
          _2028_v57 = _dafny.Seq.of((_this).f31);
          (globalState).f2 = (_1941_v2)[_module.__default.safeIndex((new BigNumber((_2028_v57).length)).plus((_this).f31), new BigNumber((_1941_v2).length))];
        } else {
          let _2029_v58;
          _2029_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f31);
          let _2030_v59;
          _2030_v59 = _dafny.Seq.of(new BigNumber((_2029_v58).length), (_this).f31);
          let _2031_v60;
          let _nw313 = new _module.C8();
          _nw313.__ctor(_dafny.MultiSet.FromArray(_2030_v59));
          _2031_v60 = _nw313;
          let _2032_v61;
          _2032_v61 = _dafny.Seq.UnicodeFromString("qtfvyyvhs");
          let _2033_v62;
          _2033_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2032_v61,_1941_v2);
          let _2034_v63;
          _2034_v63 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f27) ? (_2031_v60) : (_2031_v60)),_2033_v62);
          _2034_v63 = (_2034_v63).update(_2031_v60, _2033_v62);
          let _index289 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length));
          (_2011_v43)[_index289] = (_this).f31;
          let _index290 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length));
          let _rhs295 = (_this).f31;
          let _rhs296 = true;
          let _rhs297 = (new BigNumber(771)).minus((_this).f31);
          let _rhs298 = (_this).f31;
          let _rhs299 = (_this).f27;
          let _lhs267 = _2011_v43;
          let _lhs268 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length));
          let _lhs269 = globalState;
          let _lhs270 = globalState;
          let _lhs271 = globalState;
          let _lhs272 = globalState;
          _lhs267[_lhs268] = _rhs295;
          _lhs269.f22 = _rhs296;
          _lhs270.f12 = _rhs297;
          _lhs271.f12 = _rhs298;
          _lhs272.f22 = _rhs299;
          let _2035_v64;
          _2035_v64 = _dafny.MultiSet.fromElements((_2011_v43)[_module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length))]);
          _2035_v64 = _2035_v64;
          let _2036_v65;
          _2036_v65 = _dafny.MultiSet.fromElements(_1944_v5);
          let _2037_v66;
          _2037_v66 = _2036_v65;
          let _index291 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length));
          let _rhs300 = (_2015_v46).fm1(globalState);
          let _rhs301 = (new BigNumber(((_2037_v66)).cardinality())).multipliedBy(((_this.f30) ? ((_this).f31) : (new BigNumber((_2030_v59).length))));
          let _rhs302 = p0;
          let _rhs303 = ((_2031_v60).f29).Intersect(((_2031_v60).f29).Difference(_2035_v64));
          let _lhs273 = globalState;
          let _lhs274 = _2011_v43;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_2011_v43).length));
          let _lhs276 = globalState;
          _lhs273.f12 = _rhs300;
          _lhs274[_lhs275] = _rhs301;
          _lhs276.f15 = _rhs302;
          _2035_v64 = _rhs303;
          let _2038_v67;
          _2038_v67 = _dafny.MultiSet.fromElements(_1942_v3, _1942_v3, new _dafny.CodePoint('u'.codePointAt(0)), _1942_v3);
          let _2039_v68;
          let _nw314 = Array((new BigNumber(7)).toNumber());
          _nw314[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements(_1942_v3)).IsSubsetOf(_2038_v67);
          _nw314[(_dafny.ONE).toNumber()] = p0;
          _nw314[(new BigNumber(2)).toNumber()] = !(true) || (p2);
          _nw314[(new BigNumber(3)).toNumber()] = !(!((_this).f26));
          _nw314[(new BigNumber(4)).toNumber()] = p2;
          _nw314[(new BigNumber(5)).toNumber()] = _this.f30;
          _nw314[(new BigNumber(6)).toNumber()] = (p0) || ((_this).f27);
          _2039_v68 = _nw314;
          let _index292 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_2039_v68).length));
          (_2039_v68)[_index292] = (_this).f26;
          let _index293 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_2039_v68).length));
          (_2039_v68)[_index293] = (_this).f27;
        }
        let _2040_v69;
        _2040_v69 = _module.D2.create_DC7(p2, false, (_this).f31, (new BigNumber(((_this).f29).cardinality())).minus((_this).f31));
        let _rhs304 = ((!(p2)) ? (p2) : (p0));
        let _rhs305 = _2040_v69;
        let _lhs277 = globalState;
        _lhs277.f15 = _rhs304;
        _2040_v69 = _rhs305;
        let _2041_v70;
        let _init53 = function (_2042_i9) {
          return (_this).f26;
        };
        let _nw315 = Array((new BigNumber(23)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw315.length); _i0_53++) {
          _nw315[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _2041_v70 = _nw315;
        let _index294 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2041_v70).length));
        (_2041_v70)[_index294] = (_2010_cf48).fm2((_this).f26, (_this).f31, globalState);
        let _index295 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_2041_v70).length));
        (_2041_v70)[_index295] = p0;
      }
      let _2043_v71;
      let _nw316 = Array((new BigNumber(17)).toNumber()).fill(false);
      _2043_v71 = _nw316;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2043_v71).length))) {
        let _2044_i10 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2044_i10)) && ((_2044_i10).isLessThan(new BigNumber((_2043_v71).length))))) {
          (_2043_v71)[(_2044_i10)] = (_module.D16.create_DC38((_this).f31, new BigNumber(928), _this.f30)).dtor_cf71;
        }
      }
      let _hi18 = (_this).f31;
      for (let _2045_i11 = (_this).f31; _2045_i11.isLessThan(_hi18); _2045_i11 = _2045_i11.plus(_dafny.ONE)) {
        let _2046_v72;
        _2046_v72 = _dafny.Seq.UnicodeFromString("eobm");
        (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_2046_v72, _module.__default.safeIndex((_this).f31, new BigNumber((_2046_v72).length)), new _dafny.CodePoint('e'.codePointAt(0))), _2046_v72), _module.__default.safeIndex(_2045_i11, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_2046_v72, _module.__default.safeIndex((_this).f31, new BigNumber((_2046_v72).length)), new _dafny.CodePoint('e'.codePointAt(0))), _2046_v72)).length)), _1942_v3), _2046_v72);
        (globalState).f6 = _2046_v72;
        let _2047_v73;
        _2047_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_dafny.Seq.UnicodeFromString("yxokab"));
        let _2048_v74;
        _2048_v74 = _dafny.Seq.of(_2045_i11, (_this).f31);
        let _2049_v75;
        _2049_v75 = _dafny.Seq.of(_2048_v74);
        let _2050_v76;
        let _nw317 = new _module.C2();
        _nw317.__ctor(_2049_v75, _this.f30, p2);
        _2050_v76 = _nw317;
        let _2051_v77;
        let _nw318 = Array((new BigNumber(3)).toNumber());
        _nw318[(_dafny.ZERO).toNumber()] = _2050_v76;
        _nw318[(_dafny.ONE).toNumber()] = _2050_v76;
        _nw318[(new BigNumber(2)).toNumber()] = _2050_v76;
        _2051_v77 = _nw318;
        let _2052_v78;
        _2052_v78 = _dafny.Seq.of(_2051_v77, _2051_v77);
        let _2053_v79;
        _2053_v79 = _dafny.Set.fromElements((_this).f31, _2045_i11, new BigNumber((_2046_v72).length), new BigNumber((_2047_v73).length), new BigNumber((_dafny.Seq.Concat(_2052_v78, _dafny.Seq.of(_2051_v77, _2051_v77))).length));
        _2053_v79 = (function () {
          let _coll69 = new _dafny.Set();
          for (const _compr_69 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(689))) {
            let _2054_v80 = _compr_69;
            if (((new BigNumber(809)).isLessThanOrEqualTo(_2054_v80)) && ((_2054_v80).isLessThan(new BigNumber(689)))) {
              _coll69.add((_2054_v80).plus((_this).f31));
            }
          }
          return _coll69;
        }()).Intersect(_2053_v79);
        (globalState).f12 = _2045_i11;
      }
      (globalState).f12 = (_this).f31;
      r0 = _1942_v3;
      r1 = (_this).f31;
      return [r0, r1];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let _2055_v0;
      _2055_v0 = _module.D0.create_DC0((_this).f27);
      let _2056_v1;
      _2056_v1 = _dafny.Seq.of((_this).fm4(_2055_v0, (_this).f31, (_this).f31, globalState));
      let _2057_v2;
      _2057_v2 = _dafny.MultiSet.fromElements(_2056_v1, _2056_v1);
      let _2058_v3;
      let _nw319 = new _module.C3();
      _nw319.__ctor(_2057_v2, _this.f30, _this.f30);
      _2058_v3 = _nw319;
      _2058_v3 = _2058_v3;
      let _2059_v4;
      _2059_v4 = new _dafny.CodePoint('s'.codePointAt(0));
      let _2060_v5;
      _2060_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2056_v1);
      let _2061_v6;
      _2061_v6 = _dafny.Map.Empty.slice().updateUnsafe(_2059_v4,new BigNumber(((((_2060_v5).contains((_this).f27)) ? ((_2060_v5).get((_this).f27)) : (_2056_v1))).length));
      let _2062_v7;
      let _init54 = ((_2063_v4) => function (_2064_i1) {
        return _2063_v4;
      })(_2059_v4);
      let _nw320 = Array((new BigNumber(15)).toNumber());
      for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw320.length); _i0_54++) {
        _nw320[_i0_54] = _init54(new BigNumber(_i0_54));
      }
      _2062_v7 = _nw320;
      let _2065_v8;
      _2065_v8 = _dafny.Seq.UnicodeFromString("kbutbaoar");
      let _2066_v9;
      _2066_v9 = _module.D1.create_DC4(_2062_v7, (_this).f31, (_this).f31, _2065_v8, _2065_v8);
      let _hi19 = (_2066_v9).dtor_cf8;
      for (let _2067_i0 = ((_2058_v3).fm12(new BigNumber((_2061_v6).length), globalState)).plus((_this).f31); _2067_i0.isLessThan(_hi19); _2067_i0 = _2067_i0.plus(_dafny.ONE)) {
        _2065_v8 = _dafny.Seq.Concat(_2065_v8, _2065_v8);
        let _2068_v10;
        let _init55 = ((_2069_p0) => function (_2070_i2) {
          return (_2070_i2).multipliedBy(new BigNumber((_dafny.Set.fromElements((_this).f27, (_this).f27, _2069_p0)).length));
        })(p0);
        let _nw321 = Array((new BigNumber(3)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw321.length); _i0_55++) {
          _nw321[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _2068_v10 = _nw321;
        let _index296 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length));
        (_2068_v10)[_index296] = ((_this).f31).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("sw")).length));
        let _index297 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length));
        (_2068_v10)[_index297] = (_this).f31;
        let _2071_v11;
        _2071_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_this).fm3(_2065_v8, p2, (_this).f31, globalState));
        let _2072_v12;
        _2072_v12 = _module.D5.create_DC13((_2071_v11).Merge(_2071_v11));
        let _source29 = _2072_v12;
        if (_source29.is_DC14) {
          let _2073___mcc_h0 = (_source29).cf25;
          let _2074___mcc_h1 = (_source29).cf26;
          let _2075___mcc_h2 = (_source29).cf27;
          let _2076_cf27 = _2075___mcc_h2;
          let _2077_cf26 = _2074___mcc_h1;
          let _2078_cf25 = _2073___mcc_h0;
          let _2079_v13;
          _2079_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),_dafny.Seq.UnicodeFromString("hhvqkm"));
          let _2080_v14;
          _2080_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2078_cf25);
          let _2081_v15;
          _2081_v15 = _dafny.Seq.of(_2080_v14, _2080_v14);
          let _2082_v16;
          _2082_v16 = _module.D10.create_DC24((((_2079_v13).contains(_2067_i0)) ? ((_2079_v13).get(_2067_i0)) : (_2065_v8)), _2081_v15, (_this).f26);
          _2082_v16 = _2082_v16;
          (globalState).f12 = _module.__default.safeModuloInt((_this).f31, _2077_cf26);
          let _2083_v17;
          _2083_v17 = _dafny.Seq.of(_2056_v1, _2056_v1);
          let _2084_v18;
          let _nw322 = new _module.C2();
          _nw322.__ctor(_2083_v17, true, _this.f30);
          _2084_v18 = _nw322;
          let _2085_v19;
          _2085_v19 = _dafny.Map.Empty.slice().updateUnsafe((_2058_v3).fm1(globalState),p0);
          let _2086_v20;
          _2086_v20 = _dafny.MultiSet.fromElements(_2085_v19);
          let _2087_v21;
          _2087_v21 = _2086_v20;
          _2087_v21 = _2086_v20;
        } else if (_source29.is_DC13) {
          let _2088___mcc_h3 = (_source29).cf24;
          let _2089_cf24 = _2088___mcc_h3;
          (globalState).f15 = (_this).f27;
          let _2090_v22;
          _2090_v22 = _dafny.Map.Empty.slice().updateUnsafe(_2056_v1,(_dafny.ZERO).minus((_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))]));
          _2090_v22 = (_2090_v22).update(_2056_v1, (_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))]);
          _2059_v4 = _2059_v4;
          let _2091_v23;
          _2091_v23 = _module.D6.create_DC16(_2065_v8);
          let _2092_v24;
          _2092_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))]);
          let _2093_v25;
          _2093_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_2067_i0), _2056_v1), _module.__default.safeIndex((_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_2067_i0), _2056_v1)).length)), new BigNumber(((_2091_v23).dtor_cf29).length)),!(_2092_v24).equals(_2092_v24));
          let _index298 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length));
          (_2068_v10)[_index298] = new BigNumber((_2093_v25).length);
        } else {
          let _2094___mcc_h4 = (_source29).cf28;
          let _2095_cf28 = _2094___mcc_h4;
          let _2096_v26;
          _2096_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2067_i0,(_this).fm2((p2)[_module.__default.safeIndex((_this).f31, new BigNumber((p2).length))], _2067_i0, globalState));
          let _2097_v27;
          _2097_v27 = _module.D16.create_DC37((_2096_v26).update((_this).f31, p0));
          let _2098_v28;
          _2098_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f31),(_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))]);
          let _2099_v29;
          _2099_v29 = _dafny.Map.Empty.slice().updateUnsafe((_2097_v27).dtor_cf68,_2098_v28);
          let _2100_v31;
          _2100_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(103),(_this).f26),(_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))]);
          let _2101_v32;
          _2101_v32 = _2100_v31;
          _2099_v29 = (_2099_v29).Merge(function () {
            let _coll70 = new _dafny.Map();
            for (const _compr_70 of ((_2101_v32)).Keys.Elements) {
              let _2102_v30 = _compr_70;
              if (((_2101_v32)).contains(_2102_v30)) {
                _coll70.push([_2102_v30,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f26)).length),(_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))])]);
              }
            }
            return _coll70;
          }());
          let _2103_v33;
          _2103_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2065_v8).length),p2);
          let _2104_v34;
          _2104_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2065_v8).length));
          let _2105_v35;
          _2105_v35 = _module.D15.create_DC34((_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))], (_this).f26, (_this).f31);
          let _2106_v36;
          _2106_v36 = _dafny.MultiSet.fromElements(_2105_v35);
          let _2107_v37;
          _2107_v37 = _dafny.Map.Empty.slice().updateUnsafe((((_2103_v33).contains(new BigNumber((_2104_v34).length))) ? ((_2103_v33).get(new BigNumber((_2104_v34).length))) : (p2)),(_2106_v36).Union(_2106_v36));
          _2107_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(p2, _module.__default.safeIndex((_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))], new BigNumber((p2).length)), (_this).f27),_2106_v36);
          let _2108_v38;
          let _nw323 = new _module.C0();
          _nw323.__ctor(true);
          _2108_v38 = _nw323;
          let _2109_v39;
          _2109_v39 = _dafny.Map.Empty.slice().updateUnsafe((_2108_v38).f32,(_this).f29);
          let _2110_v40;
          _2110_v40 = _dafny.MultiSet.fromElements(p0);
          let _rhs306 = ((_this.f30) ? (_2059_v4) : (_2059_v4));
          let _rhs307 = _2108_v38;
          let _rhs308 = (new BigNumber(((_2109_v39).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm2(p0, (_2068_v10)[_module.__default.safeIndex(new BigNumber(745), new BigNumber((_2068_v10).length))], globalState),(_this).f29))).length)).multipliedBy(new BigNumber((_dafny.Seq.of((_this).f26)).length));
          let _rhs309 = _2065_v8;
          let _rhs310 = !((_2110_v40).IsSubsetOf(_2110_v40));
          let _lhs278 = globalState;
          let _lhs279 = globalState;
          let _lhs280 = globalState;
          let _lhs281 = globalState;
          _lhs278.f19 = _rhs306;
          _2108_v38 = _rhs307;
          _lhs279.f16 = _rhs308;
          _lhs280.f6 = _rhs309;
          _lhs281.f15 = _rhs310;
          let _2111_v41;
          let _nw324 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _2111_v41 = _nw324;
          let _2112_v42;
          _2112_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f26);
          let _2113_v43;
          _2113_v43 = _module.D13.create_DC29(_2112_v42);
          let _index299 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2111_v41).length));
          (_2111_v41)[_index299] = ((_2113_v43).dtor_cf50).Merge(_2112_v42);
          let _index300 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2111_v41).length));
          (_2111_v41)[_index300] = _2112_v42;
        }
        let _2114_v44;
        _2114_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2068_v10,(p2)[_module.__default.safeIndex(new BigNumber((p2).length), new BigNumber((p2).length))]);
        _2114_v44 = ((_this.f30) ? (_2114_v44) : (_2114_v44));
      }
      let _2115_v45;
      _2115_v45 = _dafny.MultiSet.fromElements((_this).f26);
      let _2116_v46;
      _2116_v46 = _dafny.Set.fromElements(false, (_this).f27);
      let _2117_v47;
      _2117_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v45,new BigNumber((_2116_v46).length));
      let _2118_v48;
      _2118_v48 = _module.D0.create_DC1(p0, false, new BigNumber((_2065_v8).length), (p2)[_module.__default.safeIndex(new BigNumber((_2117_v47).length), new BigNumber((p2).length))]);
      let _2119_v49;
      _2119_v49 = _module.D0.create_DC2(_2118_v48);
      (globalState).f12 = (_dafny.ZERO).minus(function (_source30) {
        if (_source30.is_DC1) {
          let _2120___mcc_h5 = (_source30).cf1;
          let _2121___mcc_h6 = (_source30).cf2;
          let _2122___mcc_h7 = (_source30).cf3;
          let _2123___mcc_h8 = (_source30).cf4;
          let _2124_cf4 = _2123___mcc_h8;
          let _2125_cf3 = _2122___mcc_h7;
          let _2126_cf2 = _2121___mcc_h6;
          let _2127_cf1 = _2120___mcc_h5;
          return _2125_cf3;
        } else if (_source30.is_DC0) {
          let _2128___mcc_h9 = (_source30).cf0;
          let _2129_cf0 = _2128___mcc_h9;
          return new BigNumber(691);
        } else {
          let _2130___mcc_h10 = (_source30).cf5;
          let _2131_cf5 = _2130___mcc_h10;
          return (_this).f31;
        }
      }(_2119_v49));
      let _pat_let_tv58 = _2115_v45;
      let _2132_v50;
      _2132_v50 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let59_0) {
        return function (_2135_dt__update__tmp_h0) {
          return function (_pat_let60_0) {
            return function (_2136_dt__update_hcf37_h0) {
              return function (_pat_let61_0) {
                return function (_2137_dt__update_hcf36_h0) {
                  return _module.D7.create_DC19(_2137_dt__update_hcf36_h0, _2136_dt__update_hcf37_h0, (_2135_dt__update__tmp_h0).dtor_cf38, (_2135_dt__update__tmp_h0).dtor_cf39);
                }(_pat_let61_0);
              }((_this).f31);
            }(_pat_let60_0);
          }(_pat_let_tv58);
        }(_pat_let59_0);
      }(_module.D7.create_DC19((_this).f31, _2115_v45, (_2058_v3).fm2(false, (_this).f31, globalState), _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(359)), ((_2133_v4) => function (_2134_i4) {
  return _2133_v4;
})(_2059_v4))).length))))).dtor_cf38,(_this).f27);
      let _2138_i3;
      _2138_i3 = _dafny.ZERO;
      L11: {
        while ((((_2132_v50).contains(p0)) ? ((_2132_v50).get(p0)) : ((_this).f27))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2138_i3)) {
              break L11;
            }
            _2138_i3 = (_2138_i3).plus(_dafny.ONE);
            let _2139_v51;
            _2139_v51 = _dafny.Seq.of((_this).f29);
            let _2140_v52;
            _2140_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_2065_v8).length));
            (globalState).f2 = ((_dafny.ZERO).minus((((_this).f27) ? (new BigNumber((_dafny.Set.fromElements((_this).f29, (_2139_v51)[_module.__default.safeIndex((_this).f31, new BigNumber((_2139_v51).length))])).length)) : ((((_2140_v52).contains((_this).f26)) ? ((_2140_v52).get((_this).f26)) : ((_this).f31)))))).isLessThan(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(612)), function (_2141_i5) {
              return (_this).f31;
            })).length), (_dafny.ZERO).minus((_this).f31)));
            let _2142_v53;
            let _nw325 = new _module.C3();
            _nw325.__ctor(_dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f31), _dafny.Seq.update(_2056_v1, _module.__default.safeIndex((_this).f31, new BigNumber((_2056_v1).length)), new BigNumber(-473))), false, p0);
            _2142_v53 = _nw325;
            let _2143_v54;
            _2143_v54 = _module.D20.create_DC46(_2142_v53);
            let _2144_v55;
            let _nw326 = Array((new BigNumber(28)).toNumber());
            _nw326[(_dafny.ZERO).toNumber()] = _2142_v53;
            _nw326[(_dafny.ONE).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(2)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(3)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(4)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(5)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(6)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(7)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(8)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(9)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(10)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(11)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(12)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(13)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(14)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(15)).toNumber()] = (_2143_v54).dtor_cf84;
            _nw326[(new BigNumber(16)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(17)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(18)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(19)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(20)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(21)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(22)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(23)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(24)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(25)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(26)).toNumber()] = _2142_v53;
            _nw326[(new BigNumber(27)).toNumber()] = _2142_v53;
            _2144_v55 = _nw326;
            let _index301 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2144_v55).length));
            (_2144_v55)[_index301] = _2142_v53;
            let _index302 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_2144_v55).length));
            (_2144_v55)[_index302] = _2142_v53;
            let _2145_v56;
            _2145_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_2142_v53).f27);
            let _2146_v57;
            let _nw327 = Array((new BigNumber(3)).toNumber());
            _nw327[(_dafny.ZERO).toNumber()] = (((_2145_v56).contains((_this).f31)) ? ((_2145_v56).get((_this).f31)) : ((_2142_v53).f27));
            _nw327[(_dafny.ONE).toNumber()] = (_this).f26;
            _nw327[(new BigNumber(2)).toNumber()] = (_this).f27;
            _2146_v57 = _nw327;
            let _2147_v58;
            _2147_v58 = _module.D9.create_DC21(_dafny.Map.Empty.slice().updateUnsafe(p2,_2146_v57));
            let _2148_v59;
            let _init56 = ((_2149_v52) => function (_2150_i6) {
              return _2149_v52;
            })(_2140_v52);
            let _nw328 = Array((new BigNumber(20)).toNumber());
            for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw328.length); _i0_56++) {
              _nw328[_i0_56] = _init56(new BigNumber(_i0_56));
            }
            _2148_v59 = _nw328;
            let _2151_v60;
            _2151_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((p2)[_module.__default.safeIndex((_this).f31, new BigNumber((p2).length))]),_2146_v57);
            let _rhs311 = _2058_v3;
            let _rhs312 = _module.D9.create_DC21(_2151_v60);
            let _rhs313 = _2148_v59;
            let _rhs314 = new BigNumber(110);
            let _lhs282 = globalState;
            _2058_v3 = _rhs311;
            _2147_v58 = _rhs312;
            _2148_v59 = _rhs313;
            _lhs282.f12 = _rhs314;
            let _2152_v61;
            _2152_v61 = _module.D1.create_DC5((_this).f31);
            _2152_v61 = _2152_v61;
          }
        }
      }
      let _2153_v62;
      let _init57 = function (_2154_i7) {
        return _module.__default.safeDivisionInt(_2154_i7, (_this).f31);
      };
      let _nw329 = Array((new BigNumber(15)).toNumber());
      for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw329.length); _i0_57++) {
        _nw329[_i0_57] = _init57(new BigNumber(_i0_57));
      }
      _2153_v62 = _nw329;
      let _index303 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2153_v62).length));
      (_2153_v62)[_index303] = (_this).f31;
      let _2155_v63;
      let _nw330 = Array((new BigNumber(3)).toNumber()).fill([]);
      _2155_v63 = _nw330;
      let _index304 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_2155_v63).length));
      (_2155_v63)[_index304] = _2153_v62;
      let _2156_v64;
      _2156_v64 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f27));
      let _index305 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2153_v62).length));
      let _index306 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_2155_v63).length));
      let _rhs315 = (new BigNumber(((_2156_v64)[_module.__default.safeIndex((_this).f31, new BigNumber((_2156_v64).length))]).length)).minus((_this).f31);
      let _rhs316 = (_this).f26;
      let _rhs317 = _2153_v62;
      let _rhs318 = _2153_v62;
      let _lhs283 = _2153_v62;
      let _lhs284 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2153_v62).length));
      let _lhs285 = globalState;
      let _lhs286 = _2155_v63;
      let _lhs287 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_2155_v63).length));
      _lhs283[_lhs284] = _rhs315;
      _lhs285.f15 = _rhs316;
      _lhs286[_lhs287] = _rhs317;
      _2153_v62 = _rhs318;
      (globalState).f15 = (_this).f27;
      return;
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = [];
      let r3 = undefined;
      let _2157_v0;
      _2157_v0 = _dafny.Seq.of(p3);
      (globalState).f16 = (_2157_v0)[_module.__default.safeIndex((_this).f31, new BigNumber((_2157_v0).length))];
      let _hi20 = new BigNumber(-364);
      for (let _2158_i0 = p3; _2158_i0.isLessThan(_hi20); _2158_i0 = _2158_i0.plus(_dafny.ONE)) {
        let _2159_v1;
        let _nw331 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
        _2159_v1 = _nw331;
        _2159_v1 = _2159_v1;
        let _2160_v2;
        _2160_v2 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f31),(_this).f27);
        _2160_v2 = (_2160_v2).update(p0, _this.f30);
        let _2161_v3;
        let _init58 = function (_2162_i1) {
          return false;
        };
        let _nw332 = Array((new BigNumber(9)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw332.length); _i0_58++) {
          _nw332[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2161_v3 = _nw332;
        let _index307 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2161_v3).length));
        (_2161_v3)[_index307] = (_this).f27;
        let _index308 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2161_v3).length));
        let _rhs319 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-29)), function (_2163_i2) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }), p2);
        let _rhs320 = _dafny.areEqual(_dafny.Seq.Concat(p2, _module.__default.fm22((_dafny.ZERO).minus((_this).f31), _this.f30, globalState)), p2);
        let _rhs321 = (_this).f26;
        let _lhs288 = globalState;
        let _lhs289 = globalState;
        let _lhs290 = _2161_v3;
        let _lhs291 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_2161_v3).length));
        _lhs288.f6 = _rhs319;
        _lhs289.f2 = _rhs320;
        _lhs290[_lhs291] = _rhs321;
        (globalState).f16 = new BigNumber(763);
      }
      let _2164_i3;
      _2164_i3 = _dafny.ZERO;
      L12: {
        while (!((p3).isLessThanOrEqualTo(new BigNumber(-399)))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2164_i3)) {
              break L12;
            }
            _2164_i3 = (_2164_i3).plus(_dafny.ONE);
            let _2165_v4;
            _2165_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f30);
            let _2166_v5;
            _2166_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,p1);
            let _2167_v6;
            _2167_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2165_v4).length),_2166_v5);
            let _2168_v7;
            _2168_v7 = _dafny.MultiSet.fromElements((_this).f26);
            let _2169_v8;
            _2169_v8 = new _dafny.CodePoint('o'.codePointAt(0));
            let _2170_v9;
            let _2171_v10;
            let _2172_v11;
            let _out64;
            let _out65;
            let _out66;
            let _outcollector18 = (_this).m6((((_2167_v6).contains(p1)) ? ((_2167_v6).get(p1)) : (_module.__default.fm30(new BigNumber((_2168_v7).cardinality()), true, globalState))), (_module.__default.fm39(globalState)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,p0)).length)), _2169_v8, p1, globalState);
            _out64 = _outcollector18[0];
            _out65 = _outcollector18[1];
            _out66 = _outcollector18[2];
            _2170_v9 = _out64;
            _2171_v10 = _out65;
            _2172_v11 = _out66;
            let _2173_v12;
            let _init59 = ((_2174_p0) => function (_2175_i4) {
              return (_2175_i4).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_this).f31, _2174_p0)).cardinality()));
            })(p0);
            let _nw333 = Array((new BigNumber(5)).toNumber());
            for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw333.length); _i0_59++) {
              _nw333[_i0_59] = _init59(new BigNumber(_i0_59));
            }
            _2173_v12 = _nw333;
            let _2176_v13;
            let _nw334 = new _module.C8();
            _nw334.__ctor((_this).f29);
            _2176_v13 = _nw334;
            let _2177_v14;
            _2177_v14 = _dafny.MultiSet.fromElements(_2176_v13);
            let _2178_v15;
            let _nw335 = Array((new BigNumber(10)).toNumber());
            _nw335[(_dafny.ZERO).toNumber()] = _this.f30;
            _nw335[(_dafny.ONE).toNumber()] = (_this).f26;
            _nw335[(new BigNumber(2)).toNumber()] = (_this).f26;
            _nw335[(new BigNumber(3)).toNumber()] = true;
            _nw335[(new BigNumber(4)).toNumber()] = (_this).fm2((_this).f26, new BigNumber((_2177_v14).cardinality()), globalState);
            _nw335[(new BigNumber(5)).toNumber()] = _2172_v11;
            _nw335[(new BigNumber(6)).toNumber()] = _2172_v11;
            _nw335[(new BigNumber(7)).toNumber()] = _this.f30;
            _nw335[(new BigNumber(8)).toNumber()] = _this.f30;
            _nw335[(new BigNumber(9)).toNumber()] = (_this).f26;
            _2178_v15 = _nw335;
            let _2179_v16;
            _2179_v16 = _dafny.Seq.of(_2178_v15, _2178_v15);
            let _index309 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2173_v12).length));
            (_2173_v12)[_index309] = (((_2168_v7).contains(_2172_v11)) ? ((_2168_v7).get(_2172_v11)) : (new BigNumber((_2179_v16).length)));
            let _2180_v17;
            _2180_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2166_v5).length),new BigNumber((p2).length));
            let _2181_v18;
            _2181_v18 = _dafny.Map.Empty.slice().updateUnsafe(_2172_v11,_2180_v17);
            let _2182_v19;
            _2182_v19 = _module.D21.create_DC49(_2181_v18);
            let _index310 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_2173_v12).length));
            (_2173_v12)[_index310] = (_module.__default.safeModuloInt(_2170_v9, p0)).minus(_module.__default.safeDivisionInt(_2171_v10, new BigNumber(((_2182_v19).dtor_cf86).length)));
            let _index311 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_2178_v15).length));
            (_2178_v15)[_index311] = _2172_v11;
            let _2183_v20;
            _2183_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_2169_v8);
            let _index312 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_2178_v15).length));
            (_2178_v15)[_index312] = (p1).isLessThan(new BigNumber((_2183_v20).length));
            let _index313 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_2178_v15).length));
            (_2178_v15)[_index313] = !(_dafny.MultiSet.fromElements((_this).f27, (_this).fm2((_this).f27, (_2173_v12)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_2173_v12).length))], globalState), false)).contains((_this).f27);
          }
        }
      }
      if ((_this).f27) {
        let _2184_v21;
        let _nw336 = Array((_dafny.ONE).toNumber());
        _nw336[(_dafny.ZERO).toNumber()] = (_this).f26;
        _2184_v21 = _nw336;
        let _2185_v22;
        _2185_v22 = _module.D13.create_DC30(_this.f30, (_this).f31, _this.f30, _2184_v21);
        let _2186_v23;
        _2186_v23 = _dafny.Seq.of(_this.f30, _this.f30);
        let _2187_v24;
        let _nw337 = new _module.C6();
        _nw337.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_2186_v23).length))).length));
        _2187_v24 = _nw337;
        let _2188_v25;
        _2188_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_2187_v24);
        let _2189_v26;
        _2189_v26 = _dafny.Set.fromElements((((_2188_v25).contains((_this).f26)) ? ((_2188_v25).get((_this).f26)) : (_2187_v24)));
        let _2190_v27;
        _2190_v27 = _module.D22.create_DC53(_2189_v26);
        let _2191_v28;
        _2191_v28 = _dafny.MultiSet.fromElements((_this).f27);
        let _2192_v29;
        let _nw338 = Array((new BigNumber(18)).toNumber());
        _nw338[(_dafny.ZERO).toNumber()] = (p0).isLessThanOrEqualTo(p3);
        _nw338[(_dafny.ONE).toNumber()] = (_this).f27;
        _nw338[(new BigNumber(2)).toNumber()] = !((_this).f27) || ((_this).f27);
        _nw338[(new BigNumber(3)).toNumber()] = !((_2185_v22).dtor_cf51);
        _nw338[(new BigNumber(4)).toNumber()] = _this.f30;
        _nw338[(new BigNumber(5)).toNumber()] = (_this).f27;
        _nw338[(new BigNumber(6)).toNumber()] = !((_2190_v27).dtor_cf97).equals(_2189_v26);
        _nw338[(new BigNumber(7)).toNumber()] = (_this).f27;
        _nw338[(new BigNumber(8)).toNumber()] = !((_this).f27) || (_this.f30);
        _nw338[(new BigNumber(9)).toNumber()] = _this.f30;
        _nw338[(new BigNumber(10)).toNumber()] = (_this).f26;
        _nw338[(new BigNumber(11)).toNumber()] = (_this).f26;
        _nw338[(new BigNumber(12)).toNumber()] = ((_this).f26) && ((_this).f27);
        _nw338[(new BigNumber(13)).toNumber()] = !((_2191_v28).IsDisjointFrom(_dafny.MultiSet.FromArray(_2186_v23)));
        _nw338[(new BigNumber(14)).toNumber()] = _this.f30;
        _nw338[(new BigNumber(15)).toNumber()] = ((((_2191_v28).contains(_this.f30)) ? ((_2191_v28).get(_this.f30)) : (p3))).isLessThan(new BigNumber(-253));
        _nw338[(new BigNumber(16)).toNumber()] = (_this).fm2((_this).f26, new BigNumber((p2).length), globalState);
        _nw338[(new BigNumber(17)).toNumber()] = (_this).f27;
        _2192_v29 = _nw338;
        _2192_v29 = _2184_v21;
        let _2193_v30;
        let _init60 = function (_2194_i5) {
          return (_2194_i5).minus((_this).f31);
        };
        let _nw339 = Array((new BigNumber(6)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw339.length); _i0_60++) {
          _nw339[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2193_v30 = _nw339;
        let _2195_v31;
        _2195_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f30);
        let _index314 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
        (_2193_v30)[_index314] = (new BigNumber(952)).plus(new BigNumber((_2195_v31).length));
        let _index315 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
        (_2193_v30)[_index315] = p0;
        let _2196_v32;
        let _nw340 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _2196_v32 = _nw340;
        let _2197_v33;
        _2197_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2193_v30)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length))]);
        let _index316 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2196_v32).length));
        (_2196_v32)[_index316] = (_2197_v33).Merge(_2197_v33);
        let _index317 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length));
        (_2184_v21)[_index317] = !(_this.f30);
        let _2198_v34;
        _2198_v34 = _module.D0.create_DC1((_2186_v23)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_this).f26)).length), new BigNumber((_2186_v23).length))], true, p1, (_this).f26);
        let _2199_v35;
        _2199_v35 = _module.D0.create_DC2(_2198_v34);
        let _index318 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2196_v32).length));
        let _index319 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length));
        let _rhs322 = _2197_v33;
        let _rhs323 = (_this).f27;
        let _rhs324 = ((_2193_v30)[_module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length))]).isLessThanOrEqualTo((((_2191_v28).contains((_this).f26)) ? ((_2191_v28).get((_this).f26)) : (new BigNumber((_2197_v33).length))));
        let _rhs325 = (_this).f31;
        let _rhs326 = !((_dafny.MultiSet.fromElements(p3, p3)).IsSubsetOf(_module.__default.fm10((_this).f31, p3, _module.D0.create_DC2(_2199_v35), globalState)));
        let _lhs292 = _2196_v32;
        let _lhs293 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_2196_v32).length));
        let _lhs294 = _2184_v21;
        let _lhs295 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length));
        let _lhs296 = globalState;
        let _lhs297 = globalState;
        _lhs292[_lhs293] = _rhs322;
        r0 = _rhs323;
        _lhs294[_lhs295] = _rhs324;
        _lhs296.f12 = _rhs325;
        _lhs297.f22 = _rhs326;
        let _2200_v36;
        _2200_v36 = _module.D3.create_DC10(_2193_v30);
        let _2201_v37;
        _2201_v37 = _dafny.Map.Empty.slice().updateUnsafe(_2200_v36,p0);
        let _2202_v38;
        let _nw341 = new _module.C4();
        _nw341.__ctor(_2201_v37, (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))], (_this).f27);
        _2202_v38 = _nw341;
        let _2203_v39;
        _2203_v39 = _dafny.MultiSet.fromElements(_2193_v30);
        let _rhs327 = (_2203_v39).IsSubsetOf(_2203_v39);
        let _rhs328 = (_2157_v0)[_module.__default.safeIndex((_2187_v24).fm1(globalState), new BigNumber((_2157_v0).length))];
        let _rhs329 = (((_this).f27) ? (_2202_v38) : (_2202_v38));
        let _rhs330 = p1;
        let _lhs298 = _this;
        let _lhs299 = globalState;
        let _lhs300 = globalState;
        _lhs298.f30 = _rhs327;
        _lhs299.f12 = _rhs328;
        _2202_v38 = _rhs329;
        _lhs300.f12 = _rhs330;
        let _2204_v40;
        _2204_v40 = _dafny.Set.fromElements((_2187_v24).fm1(globalState));
        let _2205_v41;
        _2205_v41 = _dafny.Set.fromElements(!((_this).f26), (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]);
        let _2206_v42;
        _2206_v42 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2205_v41);
        let _2207_v43;
        _2207_v43 = _module.D21.create_DC50((_this).f26, new BigNumber((_2204_v40).length), (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))], new BigNumber((_2206_v42).length));
        let _source31 = _2207_v43;
        if (_source31.is_DC50) {
          let _2208___mcc_h0 = (_source31).cf87;
          let _2209___mcc_h1 = (_source31).cf88;
          let _2210___mcc_h2 = (_source31).cf89;
          let _2211___mcc_h3 = (_source31).cf90;
          let _2212_cf90 = _2211___mcc_h3;
          let _2213_cf89 = _2210___mcc_h2;
          let _2214_cf88 = _2209___mcc_h1;
          let _2215_cf87 = _2208___mcc_h0;
          let _2216_v44;
          let _nw342 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2216_v44 = _nw342;
          let _index320 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2216_v44).length));
          (_2216_v44)[_index320] = _2184_v21;
          let _index321 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length));
          let _index322 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2216_v44).length));
          let _index323 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          let _rhs331 = _2213_cf89;
          let _rhs332 = _2192_v29;
          let _rhs333 = _2214_cf88;
          let _rhs334 = _dafny.Seq.Concat(p2, p2);
          let _rhs335 = p2;
          let _lhs301 = _2184_v21;
          let _lhs302 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length));
          let _lhs303 = _2216_v44;
          let _lhs304 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_2216_v44).length));
          let _lhs305 = _2193_v30;
          let _lhs306 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          let _lhs307 = globalState;
          let _lhs308 = globalState;
          _lhs301[_lhs302] = _rhs331;
          _lhs303[_lhs304] = _rhs332;
          _lhs305[_lhs306] = _rhs333;
          _lhs307.f6 = _rhs334;
          _lhs308.f6 = _rhs335;
          let _2217_v45;
          let _nw343 = Array((_dafny.ONE).toNumber());
          _2217_v45 = _nw343;
          let _index324 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2217_v45).length));
          (_2217_v45)[_index324] = _2202_v38;
          let _index325 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          let _index326 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2217_v45).length));
          let _rhs336 = !(_dafny.Seq.IsProperPrefixOf(p2, p2));
          let _rhs337 = (_this).f31;
          let _rhs338 = _dafny.Seq.UnicodeFromString("icsw");
          let _rhs339 = _2202_v38;
          let _lhs309 = _2193_v30;
          let _lhs310 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          let _lhs311 = globalState;
          let _lhs312 = _2217_v45;
          let _lhs313 = _module.__default.safeIndex(new BigNumber(519), new BigNumber((_2217_v45).length));
          r0 = _rhs336;
          _lhs309[_lhs310] = _rhs337;
          _lhs311.f6 = _rhs338;
          _lhs312[_lhs313] = _rhs339;
          let _2218_v46;
          let _init61 = ((_2219_p3) => function (_2220_i6) {
            return _dafny.Seq.of(_2219_p3);
          })(p3);
          let _nw344 = Array((new BigNumber(15)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw344.length); _i0_61++) {
            _nw344[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2218_v46 = _nw344;
          let _index327 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2218_v46).length));
          (_2218_v46)[_index327] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(827)), ((_2221_p2) => function (_2222_i7) {
            return (_dafny.ZERO).minus(new BigNumber((_2221_p2).length));
          })(p2));
          let _2223_v47;
          _2223_v47 = _module.D12.create_DC28(p3);
          let _index328 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2218_v46).length));
          (_2218_v46)[_index328] = _dafny.Seq.Concat(_2157_v0, _dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex((_2157_v0)[_module.__default.safeIndex((_2223_v47).dtor_cf49, new BigNumber((_2157_v0).length))], new BigNumber((_dafny.Seq.of(p0)).length)), p3));
          let _2224_v48;
          _2224_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_dafny.Seq.of(true)).length));
          let _2225_v50;
          _2225_v50 = _module.D17.create_DC41(!((function () {
  let _coll71 = new _dafny.Set();
  for (const _compr_71 of (_2204_v40).Elements) {
    let _2226_v49 = _compr_71;
    if ((_2204_v40).contains(_2226_v49)) {
      _coll71.add((_2226_v49).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xxuvuil")).length)));
    }
  }
  return _coll71;
}()).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_2224_v48).length), p3))), (_this).f26, p0);
          _2225_v50 = _2225_v50;
        } else if (_source31.is_DC51) {
          let _2227___mcc_h4 = (_source31).cf91;
          let _2228___mcc_h5 = (_source31).cf92;
          let _2229___mcc_h6 = (_source31).cf93;
          let _2230___mcc_h7 = (_source31).cf94;
          let _2231_cf94 = _2230___mcc_h7;
          let _2232_cf93 = _2229___mcc_h6;
          let _2233_cf92 = _2228___mcc_h5;
          let _2234_cf91 = _2227___mcc_h4;
          let _2235_v51;
          _2235_v51 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(-449));
          let _2236_v52;
          let _nw345 = new _module.C7();
          _nw345.__ctor((_2233_cf92).f27, _dafny.Map.Empty.slice().updateUnsafe(_2234_cf91,_2235_v51));
          _2236_v52 = _nw345;
          let _2237_v53;
          _2237_v53 = _dafny.Map.Empty.slice().updateUnsafe(_2236_v52,p3);
          let _index329 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          (_2193_v30)[_index329] = (_module.D17.create_DC42(new BigNumber((_2237_v53).length), p1, p3)).dtor_cf79;
          (globalState).f6 = ((_dafny.Seq.IsPrefixOf(p2, p2)) ? (((!(_this.f30)) ? (p2) : (_dafny.Seq.UnicodeFromString("hsr")))) : (p2));
          let _index330 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          (_2193_v30)[_index330] = (p1).plus(_2231_cf94);
          let _2238_v54;
          let _nw346 = new _module.C1();
          _nw346.__ctor(_module.__default.safeModuloInt(p3, (_this).f31), _2232_cf93, (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))], (p0).isEqualTo(_2232_cf93));
          _2238_v54 = _nw346;
          let _nw347 = new _module.C1();
          _nw347.__ctor((_2238_v54.f37).multipliedBy(p3), p0, !((_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]), (_2233_cf92).f26);
          _2238_v54 = _nw347;
        } else if (_source31.is_DC52) {
          let _2239___mcc_h8 = (_source31).cf95;
          let _2240___mcc_h9 = (_source31).cf96;
          let _2241_cf96 = _2240___mcc_h9;
          let _2242_cf95 = _2239___mcc_h8;
          (globalState).f12 = p1;
          _2157_v0 = _dafny.Seq.Concat(_2157_v0, _dafny.Seq.update(_dafny.Seq.Concat(_2157_v0, _dafny.Seq.update(_2157_v0, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2157_v0).length)), p0)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_2157_v0, _dafny.Seq.update(_2157_v0, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_2157_v0).length)), p0))).length)), (_this).f31));
          let _2243_v55;
          _2243_v55 = _dafny.MultiSet.fromElements(_2157_v0);
          let _2244_v56;
          let _nw348 = new _module.C3();
          _nw348.__ctor(_2243_v55, true, (_this).f27);
          _2244_v56 = _nw348;
          let _2245_v57;
          let _nw349 = Array((new BigNumber(7)).toNumber()).fill([]);
          _2245_v57 = _nw349;
          let _2246_v58;
          let _init62 = ((_2247_v28) => function (_2248_i8) {
            return _2247_v28;
          })(_2191_v28);
          let _nw350 = Array((new BigNumber(9)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw350.length); _i0_62++) {
            _nw350[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2246_v58 = _nw350;
          let _index331 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2245_v57).length));
          (_2245_v57)[_index331] = _2246_v58;
          let _index332 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2245_v57).length));
          let _rhs340 = (((_2157_v0)[_module.__default.safeIndex(p1, new BigNumber((_2157_v0).length))]).plus(p0)).minus((_dafny.ZERO).minus((_2157_v0)[_module.__default.safeIndex(new BigNumber(-797), new BigNumber((_2157_v0).length))]));
          let _rhs341 = _2246_v58;
          let _rhs342 = _2193_v30;
          let _rhs343 = new BigNumber(((_dafny.Set.fromElements(new BigNumber((_2191_v28).cardinality()))).Intersect(_dafny.Set.fromElements(p1, (_this).f31, p3))).length);
          let _lhs314 = globalState;
          let _lhs315 = _2245_v57;
          let _lhs316 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_2245_v57).length));
          let _lhs317 = globalState;
          _lhs314.f16 = _rhs340;
          _lhs315[_lhs316] = _rhs341;
          _2193_v30 = _rhs342;
          _lhs317.f16 = _rhs343;
        } else {
          let _2249___mcc_h10 = (_source31).cf86;
          let _2250_cf86 = _2249___mcc_h10;
          let _2251_v59;
          let _nw351 = new _module.C4();
          _nw351.__ctor(((_2202_v38).f33).Merge((_2202_v38).f33), !_dafny.areEqual(p2, p2), _this.f30);
          _2251_v59 = _nw351;
          let _2252_v60;
          _2252_v60 = _dafny.Map.Empty.slice().updateUnsafe((_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))],(_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]);
          let _2253_v61;
          let _nw352 = new _module.C4();
          _nw352.__ctor(((_this.f30) ? ((_2202_v38).f33) : ((_2251_v59).f33)), (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))], (((_2252_v60).contains(!((_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]))) ? ((_2252_v60).get(!((_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]))) : ((_this).f26)));
          _2253_v61 = _nw352;
          let _2254_v62;
          let _nw353 = new _module.C4();
          _nw353.__ctor((_2251_v59).f33, (_this).f27, (_this).f26);
          _2254_v62 = _nw353;
          let _2255_v63;
          _2255_v63 = _module.D21.create_DC51((_this).f31, _2254_v62, new BigNumber(875), (_this).f31);
          let _index333 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_2193_v30).length));
          (_2193_v30)[_index333] = (_2255_v63).dtor_cf91;
          let _2256_v64;
          _2256_v64 = new _dafny.CodePoint('r'.codePointAt(0));
          let _2257_v65;
          _2257_v65 = _dafny.Set.fromElements(_2256_v64, _2256_v64);
          let _2258_v66;
          let _nw354 = new _module.C6();
          _nw354.__ctor(new BigNumber((_2257_v65).length));
          _2258_v66 = _nw354;
          let _2259_v67;
          _2259_v67 = _dafny.MultiSet.fromElements(_2258_v66);
          let _2260_v68;
          _2260_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2259_v67,p3);
          let _2261_v69;
          _2261_v69 = _module.D16.create_DC37(_2195_v31);
          let _2262_v71;
          _2262_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_2256_v64);
          let _2263_v72;
          _2263_v72 = _module.D14.create_DC32(_dafny.MultiSet.fromElements((_2254_v62).f27), (_this).f26, (_this).f27, (_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))], new BigNumber(((_2262_v71).update((_2254_v62).f27, _2256_v64)).length));
          let _2264_v73;
          _2264_v73 = _dafny.Map.Empty.slice().updateUnsafe(_2263_v72,(_2254_v62).fm2(_this.f30, p1, globalState));
          let _2265_v74;
          _2265_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_2264_v73);
          let _2266_v75;
          let _nw355 = Array((new BigNumber(6)).toNumber());
          _nw355[(_dafny.ZERO).toNumber()] = (_2195_v31).update(new BigNumber((_2260_v68).length), (_this).f26);
          _nw355[(_dafny.ONE).toNumber()] = _2195_v31;
          _nw355[(new BigNumber(2)).toNumber()] = _2195_v31;
          _nw355[(new BigNumber(3)).toNumber()] = (_2195_v31).Merge((_2261_v69).dtor_cf68);
          _nw355[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_2251_v59).fm1(globalState),false);
          _nw355[(new BigNumber(5)).toNumber()] = function () {
            let _coll72 = new _dafny.Map();
            for (const _compr_72 of (_2265_v74).Keys.Elements) {
              let _2267_v70 = _compr_72;
              if ((_2265_v74).contains(_2267_v70)) {
                _coll72.push([(_2267_v70).plus(p1),(_2254_v62).f26]);
              }
            }
            return _coll72;
          }();
          _2266_v75 = _nw355;
          let _index334 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_2266_v75).length));
          (_2266_v75)[_index334] = _2195_v31;
          let _index335 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_2266_v75).length));
          (_2266_v75)[_index335] = (((_2184_v21)[_module.__default.safeIndex(new BigNumber(441), new BigNumber((_2184_v21).length))]) ? (_2195_v31) : ((_2195_v31).Merge(_2195_v31)));
        }
      } else {
        let _2268_v76;
        _2268_v76 = _dafny.Set.fromElements(_this.f30);
        let _2269_v77;
        _2269_v77 = _module.D6.create_DC17(new BigNumber((_2268_v76).length), p0, p3, p1, (_dafny.ZERO).minus(p0));
        let _2270_v78;
        _2270_v78 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f26);
        let _2271_v79;
        _2271_v79 = _module.D16.create_DC38(new BigNumber((_2270_v78).length), p1, _this.f30);
        let _2272_v80;
        _2272_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,new BigNumber((_dafny.Set.fromElements(_2271_v79)).length));
        let _2273_v81;
        _2273_v81 = _module.D0.create_DC0(_this.f30);
        let _2274_v82;
        _2274_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm4(_2273_v81, p0, (_dafny.ZERO).minus(p3), globalState),_dafny.Seq.UnicodeFromString("k"));
        let _2275_v83;
        _2275_v83 = _dafny.MultiSet.fromElements((_this).f27);
        let _2276_v84;
        _2276_v84 = _dafny.Set.fromElements((_this).f31, (((_2275_v83).contains((_this).fm2((function (_pat_let64_0) {
          return function (_2279_dt__update__tmp_h0) {
            return function (_pat_let65_0) {
              return function (_2280_dt__update_hcf36_h0) {
                return _module.D7.create_DC19(_2280_dt__update_hcf36_h0, (_2279_dt__update__tmp_h0).dtor_cf37, (_2279_dt__update__tmp_h0).dtor_cf38, (_2279_dt__update__tmp_h0).dtor_cf39);
              }(_pat_let65_0);
            }((_this).f31);
          }(_pat_let64_0);
        }(_module.D7.create_DC19(new BigNumber((_dafny.Seq.UnicodeFromString("smafhihgt")).length), _2275_v83, (_this).f26, _2272_v80))).dtor_cf38, p0, globalState))) ? ((_2275_v83).get((_this).fm2((function (_pat_let62_0) {
          return function (_2277_dt__update__tmp_h1) {
            return function (_pat_let63_0) {
              return function (_2278_dt__update_hcf36_h1) {
                return _module.D7.create_DC19(_2278_dt__update_hcf36_h1, (_2277_dt__update__tmp_h1).dtor_cf37, (_2277_dt__update__tmp_h1).dtor_cf38, (_2277_dt__update__tmp_h1).dtor_cf39);
              }(_pat_let63_0);
            }((_this).f31);
          }(_pat_let62_0);
        }(_module.D7.create_DC19(new BigNumber((_dafny.Seq.UnicodeFromString("smafhihgt")).length), _2275_v83, (_this).f26, _2272_v80))).dtor_cf38, p0, globalState))) : (p0)), p0);
        let _2281_v85;
        _2281_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f31);
        _2269_v77 = _module.__default.fm53((((_this).f27) ? (new BigNumber((_dafny.MultiSet.fromElements(p3, p3, new BigNumber((_2272_v80).length))).cardinality())) : (new BigNumber(((((_2274_v82).contains(p0)) ? ((_2274_v82).get(p0)) : (p2))).length))), p1, _2276_v84, _2281_v85, globalState);
        (globalState).f24 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(p2, p2), _dafny.Seq.Concat(p2, p2));
        (globalState).f15 = (_this).f27;
        (globalState).f16 = (p3).plus(p1);
        let _2282_v86;
        _2282_v86 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), ((_2283_p1) => function (_2284_i9) {
          return _2283_p1;
        })(p1)));
        let _2285_v87;
        let _nw356 = new _module.C2();
        _nw356.__ctor(_dafny.Seq.update(_2282_v86, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f31), new BigNumber((_2282_v86).length)), _dafny.Seq.of(p0)), (_this).f27, (_this).f27);
        _2285_v87 = _nw356;
      }
      let _2286_v88;
      let _nw357 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _2286_v88 = _nw357;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2286_v88).length))) {
        let _2287_i10 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2287_i10)) && ((_2287_i10).isLessThan(new BigNumber((_2286_v88).length))))) {
          (_2286_v88)[(_2287_i10)] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_2288_i11) {
            return ((_this.f30) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('e'.codePointAt(0))));
          });
        }
      }
      let _2289_v89;
      _2289_v89 = _module.D0.create_DC0(_this.f30);
      let _2290_v90;
      _2290_v90 = _dafny.Seq.of(!(!(false)));
      let _2291_v91;
      _2291_v91 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,(_this).f31);
      let _2292_v92;
      let _nw358 = Array((new BigNumber(26)).toNumber());
      _nw358[(_dafny.ZERO).toNumber()] = true;
      _nw358[(_dafny.ONE).toNumber()] = (_this).fm2((_this).f26, (_this).f31, globalState);
      _nw358[(new BigNumber(2)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(3)).toNumber()] = (p1).isLessThan(new BigNumber(-543));
      _nw358[(new BigNumber(4)).toNumber()] = !((_this).f27) || (_this.f30);
      _nw358[(new BigNumber(5)).toNumber()] = (_2289_v89).dtor_cf0;
      _nw358[(new BigNumber(6)).toNumber()] = (_2290_v90)[_module.__default.safeIndex(new BigNumber((_2291_v91).length), new BigNumber((_2290_v90).length))];
      _nw358[(new BigNumber(7)).toNumber()] = _dafny.Seq.IsPrefixOf(_2157_v0, _2157_v0);
      _nw358[(new BigNumber(8)).toNumber()] = ((_this).fm1(globalState)).isLessThanOrEqualTo((_dafny.ZERO).minus(p1));
      _nw358[(new BigNumber(9)).toNumber()] = (_this).f27;
      _nw358[(new BigNumber(10)).toNumber()] = (_this).f27;
      _nw358[(new BigNumber(11)).toNumber()] = (_this).f26;
      _nw358[(new BigNumber(12)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("o"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_2293_i13) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }));
      _nw358[(new BigNumber(13)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(14)).toNumber()] = true;
      _nw358[(new BigNumber(15)).toNumber()] = (_this).f27;
      _nw358[(new BigNumber(16)).toNumber()] = (_this).f27;
      _nw358[(new BigNumber(17)).toNumber()] = (_this).f27;
      _nw358[(new BigNumber(18)).toNumber()] = (_this).f26;
      _nw358[(new BigNumber(19)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(20)).toNumber()] = (_this).f26;
      _nw358[(new BigNumber(21)).toNumber()] = !((_this).f27) || (_this.f30);
      _nw358[(new BigNumber(22)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(23)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(24)).toNumber()] = _this.f30;
      _nw358[(new BigNumber(25)).toNumber()] = false;
      _2292_v92 = _nw358;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2292_v92).length))) {
        let _2294_i12 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2294_i12)) && ((_2294_i12).isLessThan(new BigNumber((_2292_v92).length))))) {
          (_2292_v92)[(_2294_i12)] = !((((_this).f27) ? ((_this).f26) : (false)));
        }
      }
      r0 = ((_this).fm2((_this).f27, p0, globalState)) === ((_dafny.MultiSet.FromArray(_2157_v0)).equals((_this).f29));
      let _2295_v93;
      let _nw359 = new _module.C0();
      _nw359.__ctor((_this).f27);
      _2295_v93 = _nw359;
      let _2296_v94;
      _2296_v94 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,_2295_v93);
      r1 = ((_this).f31).isLessThanOrEqualTo(new BigNumber((_2296_v94).length));
      let _2297_v95;
      let _nw360 = Array((new BigNumber(5)).toNumber());
      _nw360[(_dafny.ZERO).toNumber()] = p3;
      _nw360[(_dafny.ONE).toNumber()] = (_this).f31;
      _nw360[(new BigNumber(2)).toNumber()] = (_this).f31;
      _nw360[(new BigNumber(3)).toNumber()] = ((_this.f30) ? ((_this).f31) : ((_this).f31));
      _nw360[(new BigNumber(4)).toNumber()] = (_this).f31;
      _2297_v95 = _nw360;
      r2 = _2297_v95;
      let _2298_v96;
      let _nw361 = new _module.C6();
      _nw361.__ctor(p0);
      _2298_v96 = _nw361;
      r3 = _2298_v96;
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _2299_v0;
      _2299_v0 = _module.D6.create_DC16(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-172)), function (_2300_i1) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}));
      let _2301_i0;
      _2301_i0 = _dafny.ZERO;
      L13: {
        while (function (_source32) {
          if (_source32.is_DC17) {
            let _2313___mcc_h0 = (_source32).cf30;
            let _2314___mcc_h1 = (_source32).cf31;
            let _2315___mcc_h2 = (_source32).cf32;
            let _2316___mcc_h3 = (_source32).cf33;
            let _2317___mcc_h4 = (_source32).cf34;
            let _2318_cf34 = _2317___mcc_h4;
            let _2319_cf33 = _2316___mcc_h3;
            let _2320_cf32 = _2315___mcc_h2;
            let _2321_cf31 = _2314___mcc_h1;
            let _2322_cf30 = _2313___mcc_h0;
            return true;
          } else {
            let _2323___mcc_h5 = (_source32).cf29;
            let _2324_cf29 = _2323___mcc_h5;
            return !(true);
          }
        }(_2299_v0)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2301_i0)) {
              break L13;
            }
            _2301_i0 = (_2301_i0).plus(_dafny.ONE);
            let _2302_v1;
            _2302_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f27);
            let _2303_v2;
            _2303_v2 = _dafny.Seq.of((_2302_v1).Merge(_2302_v1), _2302_v1);
            _2303_v2 = _2303_v2;
            let _2304_v3;
            _2304_v3 = _dafny.Set.fromElements((_this).f31);
            let _2305_v4;
            _2305_v4 = _dafny.Seq.of(new BigNumber(-469), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2304_v3).length),(_this).f26)).length), (_this).f31);
            let _2306_v5;
            _2306_v5 = _dafny.Seq.of(_2305_v4, _2305_v4, _2305_v4);
            let _2307_v6;
            let _nw362 = new _module.C2();
            _nw362.__ctor(_2306_v5, _this.f30, (_this).f26);
            _2307_v6 = _nw362;
            (globalState).f2 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_2305_v4, _2305_v4))).IsSubsetOf(((_this).f29).Union((_this).f29));
            let _2308_v7;
            _2308_v7 = _dafny.Seq.of(_this.f30);
            let _2309_v8;
            _2309_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_2307_v6).fm1(globalState));
            let _2310_v9;
            _2310_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber((_2309_v8).length));
            let _2311_v10;
            _2311_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2308_v7).length),_2310_v9);
            let _2312_v11;
            let _nw363 = new _module.C7();
            _nw363.__ctor(_this.f30, _2311_v10);
            _2312_v11 = _nw363;
          }
        }
      }
      let _2325_v12;
      let _nw364 = new _module.C6();
      _nw364.__ctor((_this).f31);
      _2325_v12 = _nw364;
      let _2326_v13;
      _2326_v13 = _dafny.Set.fromElements(_this.f30);
      _2326_v13 = _2326_v13;
      if ((_this).f27) {
        let _2327_v14;
        _2327_v14 = _module.D16.create_DC39(_module.__default.fm54(true, (_this).f31, _2325_v12.f40, globalState));
        let _source33 = _2327_v14;
        if (_source33.is_DC38) {
          let _2328___mcc_h6 = (_source33).cf69;
          let _2329___mcc_h7 = (_source33).cf70;
          let _2330___mcc_h8 = (_source33).cf71;
          let _2331_cf71 = _2330___mcc_h8;
          let _2332_cf70 = _2329___mcc_h7;
          let _2333_cf69 = _2328___mcc_h6;
          let _2334_v15;
          _2334_v15 = _dafny.Seq.UnicodeFromString("nq");
          let _2335_v16;
          _2335_v16 = _module.D3.create_DC11(_module.__default.fm24(_2334_v15, (_this).f26, _2332_cf70, globalState));
          let _2336_v17;
          let _init63 = function (_2337_i2) {
            return !(false);
          };
          let _nw365 = Array((new BigNumber(27)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw365.length); _i0_63++) {
            _nw365[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2336_v17 = _nw365;
          let _index336 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_2336_v17).length));
          (_2336_v17)[_index336] = _this.f30;
          let _2338_v18;
          _2338_v18 = new _dafny.CodePoint('f'.codePointAt(0));
          let _2339_v19;
          _2339_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2333_cf69,true);
          let _2340_v20;
          let _nw366 = Array((new BigNumber(19)).toNumber());
          _nw366[(_dafny.ZERO).toNumber()] = _2325_v12.f40;
          _nw366[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_2325_v12.f40);
          _nw366[(new BigNumber(2)).toNumber()] = _2325_v12.f40;
          _nw366[(new BigNumber(3)).toNumber()] = new BigNumber((_2339_v19).length);
          _nw366[(new BigNumber(4)).toNumber()] = (_this).f31;
          _nw366[(new BigNumber(5)).toNumber()] = _2325_v12.f40;
          _nw366[(new BigNumber(6)).toNumber()] = new BigNumber(-305);
          _nw366[(new BigNumber(7)).toNumber()] = new BigNumber(635);
          _nw366[(new BigNumber(8)).toNumber()] = _2333_cf69;
          _nw366[(new BigNumber(9)).toNumber()] = _2332_cf70;
          _nw366[(new BigNumber(10)).toNumber()] = _2325_v12.f40;
          _nw366[(new BigNumber(11)).toNumber()] = _2332_cf70;
          _nw366[(new BigNumber(12)).toNumber()] = _2332_cf70;
          _nw366[(new BigNumber(13)).toNumber()] = _2325_v12.f40;
          _nw366[(new BigNumber(14)).toNumber()] = _2333_cf69;
          _nw366[(new BigNumber(15)).toNumber()] = _2325_v12.f40;
          _nw366[(new BigNumber(16)).toNumber()] = (_this).f31;
          _nw366[(new BigNumber(17)).toNumber()] = new BigNumber((_2334_v15).length);
          _nw366[(new BigNumber(18)).toNumber()] = (_this).f31;
          _2340_v20 = _nw366;
          let _2341_v21;
          _2341_v21 = _dafny.MultiSet.fromElements(_2340_v20, _2340_v20);
          let _2342_v22;
          _2342_v22 = _dafny.Set.fromElements(_2333_cf69, (((_2341_v21).contains(_2340_v20)) ? ((_2341_v21).get(_2340_v20)) : (_2325_v12.f40)), _2325_v12.f40);
          let _index337 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_2336_v17).length));
          let _rhs344 = _module.D3.create_DC11(_2338_v18);
          let _rhs345 = (_2326_v13).IsProperSubsetOf((_2326_v13).Difference(_dafny.Set.fromElements(_this.f30, _this.f30, (_this).f26)));
          let _rhs346 = _2325_v12.f40;
          let _rhs347 = new BigNumber((_2342_v22).length);
          let _lhs318 = _2336_v17;
          let _lhs319 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_2336_v17).length));
          let _lhs320 = globalState;
          _2335_v16 = _rhs344;
          _lhs318[_lhs319] = _rhs345;
          _2333_cf69 = _rhs346;
          _lhs320.f16 = _rhs347;
          (_2325_v12).f40 = _2325_v12.f40;
          (globalState).f16 = _2333_cf69;
          let _2343_v23;
          _2343_v23 = _dafny.Seq.of((_this).f31);
          (globalState).f15 = _dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_2343_v23).length)), _2333_cf69);
        } else if (_source33.is_DC37) {
          let _2344___mcc_h9 = (_source33).cf68;
          let _2345_cf68 = _2344___mcc_h9;
          let _2346_v24;
          let _init64 = function (_2347_i3) {
            return (_this).f26;
          };
          let _nw367 = Array((new BigNumber(18)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw367.length); _i0_64++) {
            _nw367[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2346_v24 = _nw367;
          let _2348_v25;
          _2348_v25 = _dafny.Set.fromElements((_this).f31);
          let _2349_v26;
          _2349_v26 = _module.D12.create_DC28(_2325_v12.f40);
          let _2350_v27;
          let _nw368 = Array((new BigNumber(2)).toNumber());
          _nw368[(_dafny.ZERO).toNumber()] = new BigNumber((_2348_v25).length);
          _nw368[(_dafny.ONE).toNumber()] = (_2349_v26).dtor_cf49;
          _2350_v27 = _nw368;
          let _index338 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2346_v24).length));
          (_2346_v24)[_index338] = !((new BigNumber((_dafny.Seq.of(_2350_v27, _2350_v27)).length)).isLessThanOrEqualTo((_this).f31));
          let _index339 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_2346_v24).length));
          (_2346_v24)[_index339] = (_this).f27;
          (globalState).f2 = !((_2325_v12).fm44(globalState));
          let _index340 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2350_v27).length));
          (_2350_v27)[_index340] = _module.__default.safeModuloInt((_this).f31, new BigNumber(-769));
          let _index341 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_2350_v27).length));
          (_2350_v27)[_index341] = _2325_v12.f40;
          (globalState).f2 = ((_2346_v24)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_2346_v24).length))]) === (!(_this.f30));
        } else {
          let _2351___mcc_h10 = (_source33).cf72;
          let _2352_cf72 = _2351___mcc_h10;
          let _2353_v28;
          _2353_v28 = _dafny.Seq.UnicodeFromString("veta");
          let _2354_v29;
          _2354_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_2353_v28).length));
          let _2355_v30;
          _2355_v30 = new _dafny.CodePoint('t'.codePointAt(0));
          let _2356_v31;
          _2356_v31 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)), _2355_v30);
          _2354_v29 = (_2354_v29).update((_2356_v31).IsSubsetOf(_dafny.MultiSet.fromElements(_2355_v30, _2355_v30, _2355_v30)), _module.__default.safeModuloInt(_2325_v12.f40, (_dafny.ZERO).minus(_2325_v12.f40)));
          r1 = !((_2325_v12).fm44(globalState));
          let _2357_v32;
          let _init65 = ((_2358_v12) => function (_2359_i4) {
            return (_2359_i4).plus(_2358_v12.f40);
          })(_2325_v12);
          let _nw369 = Array((new BigNumber(21)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw369.length); _i0_65++) {
            _nw369[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2357_v32 = _nw369;
          let _2360_v33;
          _2360_v33 = _dafny.Seq.of(true, (_this).f26, _this.f30, (_this).f26);
          let _2361_v34;
          _2361_v34 = _dafny.Seq.of(new BigNumber((_2360_v33).length), _2325_v12.f40);
          let _2362_v35;
          _2362_v35 = _dafny.Set.fromElements(_2325_v12.f40, (_this).fm1(globalState));
          let _nw370 = Array((new BigNumber(4)).toNumber());
          _nw370[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_2361_v34)[_module.__default.safeIndex(_module.__default.fm39(globalState), new BigNumber((_2361_v34).length))]);
          _nw370[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_2362_v35).length))).minus((_this).f31);
          _nw370[(new BigNumber(2)).toNumber()] = new BigNumber((_2353_v28).length);
          _nw370[(new BigNumber(3)).toNumber()] = (_this).f31;
          _2357_v32 = _nw370;
          let _2363_v36;
          let _nw371 = Array((new BigNumber(10)).toNumber()).fill(false);
          _2363_v36 = _nw371;
          let _index342 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_2363_v36).length));
          (_2363_v36)[_index342] = _this.f30;
          let _index343 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_2363_v36).length));
          (_2363_v36)[_index343] = (((_this).f26) ? (!((_this).f27)) : ((_this).f27));
        }
        let _2364_v37;
        _2364_v37 = _dafny.Seq.of(_2325_v12.f40, _module.__default.safeModuloInt(_2325_v12.f40, new BigNumber(((_this).f29).cardinality())));
        let _2365_v38;
        _2365_v38 = _dafny.Seq.UnicodeFromString("ivqetbs");
        (globalState).f12 = (_2364_v37)[_module.__default.safeIndex(new BigNumber((_2365_v38).length), new BigNumber((_2364_v37).length))];
        let _2366_v39;
        let _init66 = ((_2367_v38) => function (_2368_i5) {
          return (new BigNumber((_2367_v38).length)).isLessThanOrEqualTo((_this).f31);
        })(_2365_v38);
        let _nw372 = Array((new BigNumber(2)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw372.length); _i0_66++) {
          _nw372[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2366_v39 = _nw372;
        let _index344 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2366_v39).length));
        (_2366_v39)[_index344] = (_2325_v12.f40).isLessThanOrEqualTo(new BigNumber(788));
        let _index345 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2366_v39).length));
        (_2366_v39)[_index345] = (((new BigNumber(((_this).f29).cardinality())).isLessThan((_this).f31)) ? (_this.f30) : ((_this).f26));
        let _2369_v40;
        _2369_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2325_v12.f40,!(_this.f30));
        _2369_v40 = _2369_v40;
        (globalState).f15 = (_2325_v12).fm44(globalState);
      } else {
        let _2370_v41;
        let _nw373 = new _module.C5();
        _nw373.__ctor((_this).f27, !((_this).f26));
        _2370_v41 = _nw373;
        _2370_v41 = _2370_v41;
        (globalState).f16 = _module.__default.safeDivisionInt(_2325_v12.f40, (_2325_v12.f40).plus(new BigNumber(844)));
        let _2371_v42;
        _2371_v42 = _module.D7.create_DC19(_2325_v12.f40, _dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f30, (_this).f26)), (_module.D0.create_DC1((_this).f27, (_this).f26, new BigNumber(344), _this.f30)).dtor_cf4, _module.__default.fm30(_2325_v12.f40, (_this).f27, globalState));
        let _2372_v43;
        _2372_v43 = _dafny.Map.Empty.slice().updateUnsafe((_2325_v12).fm44(globalState),(_this).f31);
        let _pat_let_tv59 = _2372_v43;
        let _2373_v44;
        _2373_v44 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let66_0) {
          return function (_2374_dt__update__tmp_h0) {
            return function (_pat_let67_0) {
              return function (_2375_dt__update_hcf39_h0) {
                return _module.D7.create_DC19((_2374_dt__update__tmp_h0).dtor_cf36, (_2374_dt__update__tmp_h0).dtor_cf37, (_2374_dt__update__tmp_h0).dtor_cf38, _2375_dt__update_hcf39_h0);
              }(_pat_let67_0);
            }(_pat_let_tv59);
          }(_pat_let66_0);
        }(_2371_v42)).dtor_cf39,((_this).f27) || (_this.f30));
        _2373_v44 = (_2373_v44).update((_2372_v43).update(_this.f30, new BigNumber(990)), !(true));
        if ((_this).f26) {
          let _2376_v45;
          let _nw374 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _2376_v45 = _nw374;
          let _index346 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2376_v45).length));
          (_2376_v45)[_index346] = (_this).f31;
          let _index347 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_2376_v45).length));
          (_2376_v45)[_index347] = new BigNumber(537);
          let _2377_v46;
          _2377_v46 = _dafny.Seq.UnicodeFromString("ohnwk");
          let _2378_v47;
          _2378_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-558),_2377_v46);
          let _2379_v48;
          let _nw375 = new _module.C1();
          _nw375.__ctor(new BigNumber(((((_2378_v47).contains(_2325_v12.f40)) ? ((_2378_v47).get(_2325_v12.f40)) : (_2377_v46))).length), _2325_v12.f40, (_this).f27, (_this).f27);
          _2379_v48 = _nw375;
          let _2380_v49;
          _2380_v49 = _dafny.Map.Empty.slice().updateUnsafe((_2379_v48).f36,_2379_v48);
          let _2381_v50;
          _2381_v50 = _dafny.Seq.of(_2379_v48, _2379_v48, (((_2380_v49).contains((_this).f31)) ? ((_2380_v49).get((_this).f31)) : (_2379_v48)), _2379_v48);
          let _2382_v51;
          _2382_v51 = _dafny.Map.Empty.slice().updateUnsafe(_2381_v50,_2379_v48.f37);
          _2382_v51 = _2382_v51;
          let _2383_v52;
          _2383_v52 = _dafny.Seq.of(new BigNumber((_2377_v46).length));
          let _2384_v53;
          _2384_v53 = _dafny.Seq.of(new BigNumber((_2383_v52).length));
          let _2385_v54;
          _2385_v54 = _module.D0.create_DC0(!(true));
          let _2386_v55;
          _2386_v55 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_2383_v52, _module.__default.safeIndex((_this).fm4(_2385_v54, new BigNumber(713), new BigNumber(453), globalState), new BigNumber((_2383_v52).length)), _2325_v12.f40), _dafny.Seq.of(new BigNumber(-911)), _2383_v52);
          let _2387_v56;
          let _nw376 = new _module.C3();
          _nw376.__ctor(_2386_v55, (_this).f27, (_this).f26);
          _2387_v56 = _nw376;
          let _2388_v57;
          _2388_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_2387_v56);
          let _2389_v58;
          _2389_v58 = new _dafny.CodePoint('s'.codePointAt(0));
          let _2390_v59;
          _2390_v59 = _module.D23.create_DC55(_2387_v56);
          _2387_v56 = (((_2388_v57).contains(!_dafny.Seq.contains(_dafny.Seq.update(_2377_v46, _module.__default.safeIndex(_2325_v12.f40, new BigNumber((_2377_v46).length)), _2389_v58), _2389_v58))) ? ((_2388_v57).get(!_dafny.Seq.contains(_dafny.Seq.update(_2377_v46, _module.__default.safeIndex(_2325_v12.f40, new BigNumber((_2377_v46).length)), _2389_v58), _2389_v58))) : ((((_2388_v57).contains((_this).f26)) ? ((_2388_v57).get((_this).f26)) : ((_2390_v59).dtor_cf103))));
          (globalState).f2 = ((_this).f27) && (_this.f30);
          (globalState).f12 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_2376_v45)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_2376_v45).length))], _2325_v12.f40), _2325_v12.f40);
        } else {
          let _2391_v60;
          _2391_v60 = _dafny.MultiSet.fromElements(_2325_v12.f40);
          _2391_v60 = _2391_v60;
          let _2392_v61;
          _2392_v61 = _dafny.MultiSet.fromElements(_module.D1.create_DC3(_dafny.Seq.of((_this).f26)));
          let _2393_v62;
          _2393_v62 = _dafny.Seq.of((_this).f26, (_this).f27);
          let _2394_v63;
          _2394_v63 = _module.D1.create_DC3(_2393_v62);
          let _pat_let_tv60 = _2393_v62;
          let _2395_v64;
          _2395_v64 = _dafny.Seq.of(function (_pat_let68_0) {
            return function (_2396_dt__update__tmp_h1) {
              return function (_pat_let69_0) {
                return function (_2397_dt__update_hcf6_h0) {
                  return _module.D1.create_DC3(_2397_dt__update_hcf6_h0);
                }(_pat_let69_0);
              }(_pat_let_tv60);
            }(_pat_let68_0);
          }(_2394_v63), _2394_v63);
          let _2398_v65;
          _2398_v65 = _dafny.Set.fromElements(new BigNumber((_2393_v62).length));
          let _2399_v66;
          _2399_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2398_v65,new BigNumber(979));
          let _2400_v67;
          _2400_v67 = _module.D24.create_DC59(_2392_v61);
          let _pat_let_tv61 = _2393_v62;
          let _pat_let_tv62 = _2393_v62;
          let _2401_v68;
          let _nw377 = Array((new BigNumber(27)).toNumber());
          _nw377[(_dafny.ZERO).toNumber()] = _2392_v61;
          _nw377[(_dafny.ONE).toNumber()] = (_2392_v61).Union(_2392_v61);
          _nw377[(new BigNumber(2)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(3)).toNumber()] = (_2392_v61).update(_2394_v63, _module.__default.abs(new BigNumber((_2393_v62).length)));
          _nw377[(new BigNumber(4)).toNumber()] = ((_2392_v61).update(_2394_v63, _module.__default.abs(_2325_v12.f40))).Intersect(_2392_v61);
          _nw377[(new BigNumber(5)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(_2394_v63, _2394_v63, _2394_v63))).Union(_2392_v61);
          _nw377[(new BigNumber(7)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements(function (_pat_let70_0) {
            return function (_2402_dt__update__tmp_h2) {
              return function (_pat_let71_0) {
                return function (_2403_dt__update_hcf6_h1) {
                  return _module.D1.create_DC3(_2403_dt__update_hcf6_h1);
                }(_pat_let71_0);
              }(_pat_let_tv61);
            }(_pat_let70_0);
          }(_2394_v63))).Union(_2392_v61);
          _nw377[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_2395_v64);
          _nw377[(new BigNumber(10)).toNumber()] = (_2392_v61).Union(_2392_v61);
          _nw377[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_2394_v63, function (_pat_let72_0) {
            return function (_2404_dt__update__tmp_h3) {
              return function (_pat_let73_0) {
                return function (_2405_dt__update_hcf6_h2) {
                  return _module.D1.create_DC3(_2405_dt__update_hcf6_h2);
                }(_pat_let73_0);
              }(_pat_let_tv62);
            }(_pat_let72_0);
          }(_2394_v63));
          _nw377[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.FromArray(_2395_v64)).Union(_2392_v61);
          _nw377[(new BigNumber(13)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(14)).toNumber()] = (_2392_v61).update(_2394_v63, _module.__default.abs(new BigNumber((_2399_v66).length)));
          _nw377[(new BigNumber(15)).toNumber()] = (_2392_v61).Difference((_2400_v67).dtor_cf109);
          _nw377[(new BigNumber(16)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(17)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(18)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(19)).toNumber()] = (_dafny.MultiSet.fromElements(_2394_v63)).update(_2394_v63, _module.__default.abs((_this).f31));
          _nw377[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements(_2394_v63, _2394_v63, _2394_v63, _2394_v63);
          _nw377[(new BigNumber(21)).toNumber()] = (_2392_v61).Intersect(_2392_v61);
          _nw377[(new BigNumber(22)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(23)).toNumber()] = _dafny.MultiSet.fromElements(_2394_v63, _module.D1.create_DC3(_dafny.Seq.update(_2393_v62, _module.__default.safeIndex((_this).f31, new BigNumber((_2393_v62).length)), (_this).f26)));
          _nw377[(new BigNumber(24)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(25)).toNumber()] = _2392_v61;
          _nw377[(new BigNumber(26)).toNumber()] = _2392_v61;
          _2401_v68 = _nw377;
          let _index348 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2401_v68).length));
          (_2401_v68)[_index348] = _dafny.MultiSet.FromArray(_2395_v64);
          let _2406_v69;
          _2406_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2325_v12.f40,_this.f30)).length),(_this).f26);
          let _index349 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2401_v68).length));
          let _rhs348 = (_2392_v61).update(_module.D1.create_DC3(_2393_v62), _module.__default.abs((((_2370_v41).fm2((_this).f27, new BigNumber((_2406_v69).length), globalState)) ? (_2325_v12.f40) : (new BigNumber(120)))));
          let _rhs349 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_2325_v12.f40, _2325_v12.f40));
          let _rhs350 = _2325_v12.f40;
          let _rhs351 = (_2325_v12.f40).plus(_2325_v12.f40);
          let _rhs352 = (((_this).f27) ? (_module.__default.fm39(globalState)) : (_2325_v12.f40));
          let _lhs321 = _2401_v68;
          let _lhs322 = _module.__default.safeIndex(new BigNumber(641), new BigNumber((_2401_v68).length));
          let _lhs323 = globalState;
          let _lhs324 = globalState;
          let _lhs325 = globalState;
          let _lhs326 = _2325_v12;
          _lhs321[_lhs322] = _rhs348;
          _lhs323.f12 = _rhs349;
          _lhs324.f16 = _rhs350;
          _lhs325.f16 = _rhs351;
          _lhs326.f40 = _rhs352;
          let _2407_v70;
          let _nw378 = Array((new BigNumber(12)).toNumber()).fill([]);
          _2407_v70 = _nw378;
          (globalState).f5 = (((_this).f27) ? (_2407_v70) : (_2407_v70));
          _2393_v62 = (((_this).f27) ? (_2393_v62) : (_dafny.Seq.of(_this.f30)));
          let _2408_v71;
          _2408_v71 = _dafny.MultiSet.fromElements((_this).f26);
          let _2409_v72;
          _2409_v72 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f27, true), _2408_v71, _dafny.MultiSet.fromElements(false));
          let _2410_v73;
          _2410_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_2408_v71);
          let _2411_v74;
          _2411_v74 = _module.D14.create_DC32(_2408_v71, (_this).f27, (_this).f26, _this.f30, new BigNumber(407));
          let _2412_v75;
          let _nw379 = Array((new BigNumber(15)).toNumber());
          _nw379[(_dafny.ZERO).toNumber()] = _2408_v71;
          _nw379[(_dafny.ONE).toNumber()] = _2408_v71;
          _nw379[(new BigNumber(2)).toNumber()] = _2408_v71;
          _nw379[(new BigNumber(3)).toNumber()] = _module.__default.fm21((_this).f31, globalState);
          _nw379[(new BigNumber(4)).toNumber()] = (_2409_v72)[_module.__default.safeIndex(_2325_v12.f40, new BigNumber((_2409_v72).length))];
          _nw379[(new BigNumber(5)).toNumber()] = (_2408_v71).Difference(_2408_v71);
          _nw379[(new BigNumber(6)).toNumber()] = (((_2410_v73).contains(_this.f30)) ? ((_2410_v73).get(_this.f30)) : (_2408_v71));
          _nw379[(new BigNumber(7)).toNumber()] = (_2411_v74).dtor_cf56;
          _nw379[(new BigNumber(8)).toNumber()] = _2408_v71;
          _nw379[(new BigNumber(9)).toNumber()] = _2408_v71;
          _nw379[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray((((_this).f27) ? (_2393_v62) : (_2393_v62)));
          _nw379[(new BigNumber(11)).toNumber()] = _2408_v71;
          _nw379[(new BigNumber(12)).toNumber()] = (_2409_v72)[_module.__default.safeIndex(_2325_v12.f40, new BigNumber((_2409_v72).length))];
          _nw379[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).fm2((_this).f27, new BigNumber(49), globalState)));
          _nw379[(new BigNumber(14)).toNumber()] = _2408_v71;
          _2412_v75 = _nw379;
          let _2413_v76;
          let _init67 = ((_2414_v12) => function (_2415_i6) {
            return (_2415_i6).multipliedBy((_dafny.ZERO).minus(_2414_v12.f40));
          })(_2325_v12);
          let _nw380 = Array((new BigNumber(19)).toNumber());
          for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw380.length); _i0_67++) {
            _nw380[_i0_67] = _init67(new BigNumber(_i0_67));
          }
          _2413_v76 = _nw380;
          let _2416_v77;
          _2416_v77 = _dafny.Seq.of(_2413_v76);
          let _index350 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_2412_v75).length));
          (_2412_v75)[_index350] = _module.__default.fm21(new BigNumber((_2416_v77).length), globalState);
          let _2417_v78;
          _2417_v78 = _dafny.Seq.of((_2325_v12.f40).plus(_2325_v12.f40));
          let _2418_v79;
          _2418_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2325_v12.f40,_2325_v12.f40);
          let _index351 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_2412_v75).length));
          let _rhs353 = ((_this.f30) ? (_2325_v12.f40) : (_module.__default.safeModuloInt((_this).f31, _2325_v12.f40)));
          let _rhs354 = (_2417_v78)[_module.__default.safeIndex((_this).fm1(globalState), new BigNumber((_2417_v78).length))];
          let _rhs355 = _2408_v71;
          let _rhs356 = (_this).fm2((_2325_v12).fm2((_this).f27, _2325_v12.f40, globalState), _module.__default.safeDivisionInt(_2325_v12.f40, new BigNumber((_2418_v79).length)), globalState);
          let _lhs327 = globalState;
          let _lhs328 = _2325_v12;
          let _lhs329 = _2412_v75;
          let _lhs330 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_2412_v75).length));
          let _lhs331 = globalState;
          _lhs327.f16 = _rhs353;
          _lhs328.f40 = _rhs354;
          _lhs329[_lhs330] = _rhs355;
          _lhs331.f22 = _rhs356;
        }
        let _2419_v80;
        _2419_v80 = _dafny.Seq.of(_dafny.Seq.of((_this).f31, (_this).f31));
        let _2420_v81;
        _2420_v81 = _dafny.Seq.of(_2419_v80);
        let _2421_v82;
        let _nw381 = new _module.C2();
        _nw381.__ctor(((_this.f30) ? ((_2420_v81)[_module.__default.safeIndex(_2325_v12.f40, new BigNumber((_2420_v81).length))]) : (_2419_v80)), !(true) || (!((_this).f26)), (_this).f27);
        _2421_v82 = _nw381;
        _2421_v82 = _2421_v82;
      }
      let _2422_v83;
      _2422_v83 = _dafny.Seq.UnicodeFromString("uuuanc");
      let _2423_v84;
      _2423_v84 = _dafny.Seq.of((new BigNumber(884)).plus(new BigNumber((_2422_v83).length)), _2325_v12.f40);
      let _2424_v85;
      _2424_v85 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2425_v86;
      _2425_v86 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,_2424_v85);
      _2423_v84 = _dafny.Seq.Concat(_2423_v84, _dafny.Seq.Concat(_2423_v84, _module.__default.fm43((_this).f31, _2325_v12.f40, new BigNumber((_dafny.Seq.UnicodeFromString("yjmphn")).length), _2425_v86, globalState)));
      let _2426_v87;
      let _nw382 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _2426_v87 = _nw382;
      let _index352 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_2426_v87).length));
      (_2426_v87)[_index352] = (_this).f31;
      let _2427_v88;
      let _init68 = function (_2428_i7) {
        return (_this).f27;
      };
      let _nw383 = Array((new BigNumber(12)).toNumber());
      for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw383.length); _i0_68++) {
        _nw383[_i0_68] = _init68(new BigNumber(_i0_68));
      }
      _2427_v88 = _nw383;
      let _2429_v89;
      _2429_v89 = _dafny.Set.fromElements(_2427_v88);
      let _2430_v90;
      _2430_v90 = _dafny.Seq.of(_2429_v89, _2429_v89, _2429_v89);
      let _2431_v91;
      _2431_v91 = _dafny.Map.Empty.slice().updateUnsafe(_2325_v12.f40,_2429_v89);
      let _index353 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_2426_v87).length));
      (_2426_v87)[_index353] = new BigNumber((((_2430_v90)[_module.__default.safeIndex((_dafny.ZERO).minus(_2325_v12.f40), new BigNumber((_2430_v90).length))]).Union((((_2431_v91).contains(_2325_v12.f40)) ? ((_2431_v91).get(_2325_v12.f40)) : (_2429_v89)))).length);
      let _2432_v92;
      _2432_v92 = _dafny.Seq.of((_this).f26);
      let _2433_v93;
      _2433_v93 = _module.D16.create_DC38(new BigNumber(718), new BigNumber((_2432_v92).length), (_this).f26);
      let _pat_let_tv63 = _2426_v87;
      let _pat_let_tv64 = _2426_v87;
      let _pat_let_tv65 = _2432_v92;
      let _pat_let_tv66 = _2326_v13;
      r0 = function (_source34) {
        if (_source34.is_DC38) {
          let _2434___mcc_h11 = (_source34).cf69;
          let _2435___mcc_h12 = (_source34).cf70;
          let _2436___mcc_h13 = (_source34).cf71;
          let _2437_cf71 = _2436___mcc_h13;
          let _2438_cf70 = _2435___mcc_h12;
          let _2439_cf69 = _2434___mcc_h11;
          return (_this).f26;
        } else if (_source34.is_DC37) {
          let _2440___mcc_h14 = (_source34).cf68;
          let _2441_cf68 = _2440___mcc_h14;
          return !(((_pat_let_tv64)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_pat_let_tv63).length))]).isLessThan(new BigNumber((_pat_let_tv65).length)));
        } else {
          let _2442___mcc_h15 = (_source34).cf72;
          let _2443_cf72 = _2442___mcc_h15;
          return !(!(_pat_let_tv66).equals((_dafny.Set.fromElements(_this.f30, true))));
        }
      }(_2433_v93);
      r1 = (_this).f27;
      return [r0, r1];
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _2444_v0;
      _2444_v0 = _dafny.Seq.UnicodeFromString("flsqn");
      let _2445_v1;
      _2445_v1 = _dafny.Seq.of(new BigNumber((_2444_v0).length), p1);
      let _2446_v2;
      let _nw384 = Array((new BigNumber(14)).toNumber());
      _nw384[(_dafny.ZERO).toNumber()] = p1;
      _nw384[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.of(_this.f30)).length);
      _nw384[(new BigNumber(2)).toNumber()] = (_this).f31;
      _nw384[(new BigNumber(3)).toNumber()] = (_this).f31;
      _nw384[(new BigNumber(4)).toNumber()] = (_this).f31;
      _nw384[(new BigNumber(5)).toNumber()] = new BigNumber((_2445_v1).length);
      _nw384[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f31,p1)).length);
      _nw384[(new BigNumber(7)).toNumber()] = (_this).f31;
      _nw384[(new BigNumber(8)).toNumber()] = p1;
      _nw384[(new BigNumber(9)).toNumber()] = (_this).f31;
      _nw384[(new BigNumber(10)).toNumber()] = p0;
      _nw384[(new BigNumber(11)).toNumber()] = p1;
      _nw384[(new BigNumber(12)).toNumber()] = p1;
      _nw384[(new BigNumber(13)).toNumber()] = p0;
      _2446_v2 = _nw384;
      let _2447_v3;
      _2447_v3 = _dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC10(_2446_v2)).dtor_cf21,p1);
      _2447_v3 = (_2447_v3).update(_2446_v2, (_dafny.ZERO).minus(p0));
      r1 = _module.__default.fm39(globalState);
      let _2448_v4;
      _2448_v4 = _dafny.MultiSet.fromElements(p2);
      let _2449_v5;
      _2449_v5 = _dafny.Set.fromElements(new BigNumber((_2445_v1).length), p1, p0, new BigNumber((_dafny.Seq.UnicodeFromString("ujlp")).length), (_dafny.ZERO).minus(new BigNumber((_2448_v4).cardinality())));
      let _2450_v6;
      _2450_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,_2449_v5);
      let _2451_v7;
      _2451_v7 = _dafny.Seq.of(_dafny.Set.fromElements(p1, new BigNumber(-51), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f27)).length)), _dafny.Set.fromElements(p0), (((_2450_v6).contains((_this).f27)) ? ((_2450_v6).get((_this).f27)) : (_2449_v5)));
      let _2452_v8;
      _2452_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2451_v7,(_this).f26);
      let _2453_i0;
      _2453_i0 = _dafny.ZERO;
      L14: {
        while (!((((_2452_v8).contains(_2451_v7)) ? ((_2452_v8).get(_2451_v7)) : (_this.f30)))) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2453_i0)) {
              break L14;
            }
            _2453_i0 = (_2453_i0).plus(_dafny.ONE);
            let _2454_v9;
            _2454_v9 = _dafny.Seq.of(_this.f30, true);
            let _2455_v10;
            _2455_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_2454_v9, _module.__default.safeIndex(p1, new BigNumber((_2454_v9).length)), true),p0);
            let _2456_v11;
            let _nw385 = new _module.C5();
            _nw385.__ctor((_this).f26, _this.f30);
            _2456_v11 = _nw385;
            let _2457_v12;
            _2457_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_2456_v11);
            let _2458_v13;
            _2458_v13 = _dafny.Set.fromElements(_this.f30);
            let _2459_v14;
            _2459_v14 = _dafny.Set.fromElements(_2458_v13);
            let _2460_v15;
            _2460_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2459_v14).length),(_this).f26);
            let _2461_v16;
            _2461_v16 = _module.D16.create_DC38(new BigNumber((_2457_v12).length), new BigNumber((_2460_v15).length), (_this).f27);
            let _2462_v17;
            _2462_v17 = _dafny.Seq.of(_2461_v16, _2461_v16);
            r1 = (_this).fm4(_module.D0.create_DC0((_this).f26), (((_2455_v10).contains(_dafny.Seq.of((_this).f27, _this.f30))) ? ((_2455_v10).get(_dafny.Seq.of((_this).f27, _this.f30))) : (p1)), _module.__default.safeModuloInt(new BigNumber((_2454_v9).length), new BigNumber((_2462_v17).length)), globalState);
            (globalState).f16 = new BigNumber(154);
            let _2463_v18;
            _2463_v18 = _dafny.Map.Empty.slice().updateUnsafe(((!((_this).f27)) ? ((_this).f26) : (p2)),!((_this).f26) || (_this.f30));
            _2463_v18 = (_2463_v18).update(p2, _this.f30);
            let _2464_v19;
            _2464_v19 = _dafny.MultiSet.fromElements(_2445_v1);
            let _2465_v20;
            let _nw386 = new _module.C3();
            _nw386.__ctor(_2464_v19, true, (_dafny.MultiSet.fromElements(_this.f30)).contains(p2));
            _2465_v20 = _nw386;
          }
        }
      }
      r1 = (_this).f31;
      (globalState).f15 = false;
      if ((_this).fm2(p2, p1, globalState)) {
        let _2466_v21;
        let _nw387 = Array((new BigNumber(6)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _2466_v21 = _nw387;
        let _2467_v22;
        let _init69 = function (_2468_i1) {
          return ((_this.f30) ? (_module.D20.create_DC47()) : (_module.D20.create_DC47()));
        };
        let _nw388 = Array((new BigNumber(16)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw388.length); _i0_69++) {
          _nw388[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2467_v22 = _nw388;
        let _2469_v23;
        _2469_v23 = _module.D20.create_DC47();
        let _index354 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_2467_v22).length));
        (_2467_v22)[_index354] = _2469_v23;
        let _index355 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_2467_v22).length));
        let _rhs357 = _dafny.Seq.Concat(_2444_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-395)), function (_2470_i2) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }));
        let _rhs358 = _2466_v21;
        let _rhs359 = _2446_v2;
        let _rhs360 = _2469_v23;
        let _lhs332 = globalState;
        let _lhs333 = _2467_v22;
        let _lhs334 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_2467_v22).length));
        _lhs332.f6 = _rhs357;
        _2466_v21 = _rhs358;
        _2446_v2 = _rhs359;
        _lhs333[_lhs334] = _rhs360;
        (globalState).f16 = new BigNumber((((_2448_v4).Union(_2448_v4)).Difference(_2448_v4)).cardinality());
        r1 = (_dafny.ZERO).minus(p0);
        let _2471_v24;
        _2471_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2449_v5).length),(_this).f26);
        let _2472_v25;
        _2472_v25 = _module.D0.create_DC0((_this).f27);
        let _2473_v26;
        _2473_v26 = _dafny.Seq.of((_this).f26, (_this).f26);
        if (((_this).fm4(_2472_v25, new BigNumber((_2473_v26).length), p0, globalState)).isLessThan(new BigNumber((((p2) ? (_2471_v24) : (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f26)))).length))) {
          let _2474_v27;
          _2474_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          (globalState).f12 = new BigNumber((((_2474_v27).update(p0, p0)).Merge(_2474_v27)).length);
          let _2475_v28;
          let _nw389 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
          _2475_v28 = _nw389;
          let _index356 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2475_v28).length));
          (_2475_v28)[_index356] = (_this).f29;
          let _2476_v29;
          _2476_v29 = new _dafny.CodePoint('p'.codePointAt(0));
          let _2477_v30;
          _2477_v30 = _module.D6.create_DC16(_2444_v0);
          let _2478_v31;
          let _nw390 = Array((new BigNumber(16)).toNumber());
          _nw390[(_dafny.ZERO).toNumber()] = _2444_v0;
          _nw390[(_dafny.ONE).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(2)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(3)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(4)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(881)), function (_2479_i3) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }), _2444_v0);
          _nw390[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("gixiulo");
          _nw390[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(521)), function (_2480_i4) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          });
          _nw390[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("gxxceu");
          _nw390[(new BigNumber(9)).toNumber()] = ((_this.f30) ? (_dafny.Seq.UnicodeFromString("qkrwt")) : (_2444_v0));
          _nw390[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_2444_v0, _module.__default.safeIndex(new BigNumber((_2444_v0).length), new BigNumber((_2444_v0).length)), _2476_v29);
          _nw390[(new BigNumber(11)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(12)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_2444_v0, _2444_v0);
          _nw390[(new BigNumber(14)).toNumber()] = _2444_v0;
          _nw390[(new BigNumber(15)).toNumber()] = (_2477_v30).dtor_cf29;
          _2478_v31 = _nw390;
          let _index357 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2478_v31).length));
          (_2478_v31)[_index357] = _dafny.Seq.Concat(_2444_v0, _module.__default.fm22(new BigNumber((_dafny.Seq.UnicodeFromString("aj")).length), _this.f30, globalState));
          let _2481_v32;
          _2481_v32 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_2473_v26), ((_2448_v4).update(true, _module.__default.abs((_2445_v1)[_module.__default.safeIndex(new BigNumber(757), new BigNumber((_2445_v1).length))]))).update((_this).f27, _module.__default.abs(p1)));
          let _2482_v33;
          _2482_v33 = _dafny.Seq.of(_module.__default.fm22(p0, (_this).f26, globalState));
          let _index358 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2475_v28).length));
          let _index359 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2478_v31).length));
          let _rhs361 = (_2481_v32)[_module.__default.safeIndex((((_2448_v4).contains((_this).f26)) ? ((_2448_v4).get((_this).f26)) : (p1)), new BigNumber((_2481_v32).length))];
          let _rhs362 = (((_this).f29).Difference(_module.__default.fm42(new BigNumber(((_this).f29).cardinality()), globalState))).Difference((_this).f29);
          let _rhs363 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pyblpod"), _dafny.Seq.Concat((_2482_v33)[_module.__default.safeIndex(p1, new BigNumber((_2482_v33).length))], _2444_v0));
          let _lhs335 = _2475_v28;
          let _lhs336 = _module.__default.safeIndex(new BigNumber(654), new BigNumber((_2475_v28).length));
          let _lhs337 = _2478_v31;
          let _lhs338 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_2478_v31).length));
          _2448_v4 = _rhs361;
          _lhs335[_lhs336] = _rhs362;
          _lhs337[_lhs338] = _rhs363;
          (globalState).f16 = p1;
          let _2483_v34;
          let _nw391 = new _module.C0();
          _nw391.__ctor(_this.f30);
          _2483_v34 = _nw391;
          _2483_v34 = _2483_v34;
          let _2484_v35;
          let _init70 = function (_2485_i5) {
            return (_this).f27;
          };
          let _nw392 = Array((new BigNumber(5)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw392.length); _i0_70++) {
            _nw392[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2484_v35 = _nw392;
          let _2486_v36;
          _2486_v36 = _module.D5.create_DC14(p1, p0, (_2483_v34).f32);
          let _2487_v37;
          let _nw393 = new _module.C1();
          _nw393.__ctor(p0, p0, (_2486_v36).dtor_cf27, (_2474_v27).equals(_2474_v27));
          _2487_v37 = _nw393;
          let _2488_v38;
          _2488_v38 = _dafny.Map.Empty.slice().updateUnsafe((_2483_v34).f32,new BigNumber(-689));
          let _rhs364 = _2484_v35;
          let _rhs365 = _2487_v37;
          let _rhs366 = _dafny.Seq.of((_2483_v34).fm7(_2488_v38, _2445_v1, p2, globalState));
          let _rhs367 = !(_this.f30);
          let _rhs368 = !(_this.f30) || ((false) || (p2));
          let _lhs339 = globalState;
          let _lhs340 = globalState;
          _2484_v35 = _rhs364;
          _2487_v37 = _rhs365;
          _2473_v26 = _rhs366;
          _lhs339.f15 = _rhs367;
          _lhs340.f15 = _rhs368;
        } else {
          let _2489_v39;
          _2489_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,p1);
          let _2490_v40;
          _2490_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f26),_2489_v39);
          _2490_v40 = _2490_v40;
          let _2491_v41;
          _2491_v41 = new _dafny.CodePoint('t'.codePointAt(0));
          let _2492_v42;
          let _2493_v43;
          let _2494_v44;
          let _2495_v45;
          let _out67;
          let _out68;
          let _out69;
          let _out70;
          let _outcollector19 = _module.__default.m0(_2491_v41, globalState);
          _out67 = _outcollector19[0];
          _out68 = _outcollector19[1];
          _out69 = _outcollector19[2];
          _out70 = _outcollector19[3];
          _2492_v42 = _out67;
          _2493_v43 = _out68;
          _2494_v44 = _out69;
          _2495_v45 = _out70;
          let _2496_v46;
          let _nw394 = new _module.C0();
          _nw394.__ctor((p0).isLessThan((_this).f31));
          _2496_v46 = _nw394;
          let _2497_v47;
          let _nw395 = new _module.C2();
          _nw395.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(437)), ((_2498_v1) => function (_2499_i6) {
            return _2498_v1;
          })(_2445_v1)), (_this).f26, (_2496_v46).f32);
          _2497_v47 = _nw395;
          let _2500_v48;
          let _nw396 = new _module.C5();
          _nw396.__ctor((_this).f26, (_this).f27);
          _2500_v48 = _nw396;
          let _2501_v49;
          _2501_v49 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_2444_v0).length)).plus(p0),_2500_v48);
          let _2502_v50;
          _2502_v50 = _dafny.Map.Empty.slice().updateUnsafe((_2496_v46).f32,(_this).f27);
          _2501_v49 = (_2501_v49).update(new BigNumber(((_2502_v50).Merge(_2502_v50)).length), _2500_v48);
        }
        let _2503_v51;
        _2503_v51 = _dafny.Set.fromElements((_this).f26, (_this).f27);
        (globalState).f16 = new BigNumber((_2503_v51).length);
      } else {
        let _2504_v52;
        let _nw397 = new _module.C5();
        _nw397.__ctor((new BigNumber(126)).isLessThanOrEqualTo(p0), ((p2) ? (true) : ((_this).f27)));
        _2504_v52 = _nw397;
        let _2505_v53;
        let _init71 = ((_2506_v0) => function (_2507_i7) {
          return _module.D6.create_DC16(_2506_v0);
        })(_2444_v0);
        let _nw398 = Array((new BigNumber(13)).toNumber());
        for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw398.length); _i0_71++) {
          _nw398[_i0_71] = _init71(new BigNumber(_i0_71));
        }
        _2505_v53 = _nw398;
        let _2508_v54;
        let _init72 = function (_2509_i8) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        };
        let _nw399 = Array((new BigNumber(13)).toNumber());
        for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw399.length); _i0_72++) {
          _nw399[_i0_72] = _init72(new BigNumber(_i0_72));
        }
        _2508_v54 = _nw399;
        let _2510_v55;
        _2510_v55 = _module.D1.create_DC4(_2508_v54, p1, new BigNumber((_2448_v4).cardinality()), _2444_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-896)), function (_2511_i9) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}));
        let _index360 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2505_v53).length));
        (_2505_v53)[_index360] = _module.D6.create_DC16((_2510_v55).dtor_cf11);
        let _2512_v56;
        let _init73 = ((_2513_v0) => function (_2514_i10) {
          return _2513_v0;
        })(_2444_v0);
        let _nw400 = Array((new BigNumber(17)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw400.length); _i0_73++) {
          _nw400[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2512_v56 = _nw400;
        let _index361 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length));
        (_2512_v56)[_index361] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(704)), function (_2515_i11) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        });
        let _2516_v57;
        _2516_v57 = _module.D21.create_DC50(_this.f30, (_dafny.ZERO).minus((_this).f31), _this.f30, (_this).f31);
        let _index362 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2505_v53).length));
        let _index363 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length));
        let _rhs369 = _2504_v52;
        let _rhs370 = (p1).isEqualTo((p1).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-898)))));
        let _rhs371 = _this.f30;
        let _rhs372 = _module.__default.fm38((_2516_v57).dtor_cf89, (_dafny.ZERO).minus((_this).fm1(globalState)), _2444_v0, globalState);
        let _rhs373 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), function (_2517_i12) {
          return ((false) ? (new _dafny.CodePoint('f'.codePointAt(0))) : (new _dafny.CodePoint('e'.codePointAt(0))));
        });
        let _lhs341 = globalState;
        let _lhs342 = globalState;
        let _lhs343 = _2505_v53;
        let _lhs344 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_2505_v53).length));
        let _lhs345 = _2512_v56;
        let _lhs346 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length));
        _2504_v52 = _rhs369;
        _lhs341.f2 = _rhs370;
        _lhs342.f22 = _rhs371;
        _lhs343[_lhs344] = _rhs372;
        _lhs345[_lhs346] = _rhs373;
        let _2518_v58;
        _2518_v58 = _dafny.Set.fromElements((_this).f26);
        let _2519_v59;
        _2519_v59 = _dafny.Seq.of(_2518_v58, _2518_v58);
        let _2520_v60;
        _2520_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2519_v59,_2444_v0);
        (globalState).f6 = (((_2520_v60).contains(_2519_v59)) ? ((_2520_v60).get(_2519_v59)) : (_2444_v0));
        r0 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ofrmiv"), (_2512_v56)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length))]);
        (globalState).f16 = (_this).f31;
        let _2521_v61;
        _2521_v61 = _dafny.Set.fromElements(_module.__default.fm22(p0, _this.f30, globalState), (_2512_v56)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length))], ((p2) ? ((_2512_v56)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length))]) : (_dafny.Seq.UnicodeFromString("x"))));
        let _2522_v62;
        _2522_v62 = _dafny.Map.Empty.slice().updateUnsafe((_2512_v56)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length))],p0);
        let _rhs374 = (_dafny.Set.fromElements((_2512_v56)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_2512_v56).length))])).Union(function () {
          let _coll73 = new _dafny.Set();
          for (const _compr_73 of (_2522_v62).Keys.Elements) {
            let _2523_v63 = _compr_73;
            if ((_2522_v62).contains(_2523_v63)) {
              _coll73.add(_2523_v63);
            }
          }
          return _coll73;
        }());
        let _rhs375 = (_this).f26;
        let _lhs347 = globalState;
        _2521_v61 = _rhs374;
        _lhs347.f2 = _rhs375;
      }
      r0 = _module.__default.fm22(p1, !((p2) === (_this.f30)), globalState);
      r1 = (_dafny.ZERO).minus((_this).fm4(_module.D0.create_DC0((_this).f27), (p0).plus(p1), p1, globalState));
      r2 = p2;
      return [r0, r1, r2];
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _2524_v0;
      _2524_v0 = _dafny.Seq.of((_this).f27, (_this).f26);
      let _hi21 = new BigNumber((_dafny.Seq.Concat(_2524_v0, _2524_v0)).length);
      for (let _2525_i0 = p1; _2525_i0.isLessThan(_hi21); _2525_i0 = _2525_i0.plus(_dafny.ONE)) {
        let _2526_v1;
        let _init74 = ((_2527_v0, _2528_i0, _2529_p1, _2530_p3) => function (_2531_i1) {
          return _dafny.areEqual(_dafny.Seq.of(new BigNumber((_2527_v0).length), _2528_i0), _dafny.Seq.of(_2529_p1, (_this).f31, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(734)), ((_2532_p3) => function (_2533_i2) {
            return _2532_p3;
          })(_2530_p3))).length)));
        })(_2524_v0, _2525_i0, p1, p3);
        let _nw401 = Array((new BigNumber(7)).toNumber());
        for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw401.length); _i0_74++) {
          _nw401[_i0_74] = _init74(new BigNumber(_i0_74));
        }
        _2526_v1 = _nw401;
        let _index364 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length));
        (_2526_v1)[_index364] = ((!((_this).f26)) ? (_this.f30) : ((_this).f26));
        let _2534_v2;
        _2534_v2 = _dafny.Map.Empty.slice().updateUnsafe(p1,!((_this).f26));
        let _index365 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length));
        (_2526_v1)[_index365] = (((_2534_v2).contains(p3)) ? ((_2534_v2).get(p3)) : (true));
        let _2535_v3;
        _2535_v3 = _module.D1.create_DC3(_dafny.Seq.of(!(_this.f30)));
        let _2536_v5;
        _2536_v5 = _dafny.Seq.of(function () {
          let _coll74 = new _dafny.Set();
          for (const _compr_74 of _dafny.IntegerRange(new BigNumber(435), new BigNumber(179))) {
            let _2537_v4 = _compr_74;
            if (((new BigNumber(435)).isLessThanOrEqualTo(_2537_v4)) && ((_2537_v4).isLessThan(new BigNumber(179)))) {
              _coll74.add(_module.__default.safeModuloInt(_2537_v4, new BigNumber((p0).length)));
            }
          }
          return _coll74;
        }());
        let _2538_v6;
        _2538_v6 = _dafny.Set.fromElements(new BigNumber(545), new BigNumber(-539), (_this).fm1(globalState), new BigNumber(747), p3);
        let _2539_v7;
        _2539_v7 = _dafny.Set.fromElements((_this).f26, true, (_this).fm2((_2526_v1)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length))], _2525_i0, globalState));
        let _2540_v8;
        _2540_v8 = _module.D6.create_DC17(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_2526_v1)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length))], (_this).f27),(_this).f31)).length), (_this).f31, p3, new BigNumber(-624), p1);
        let _index366 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length));
        let _rhs376 = !_dafny.areEqual(_module.__default.fm55(globalState), _dafny.Seq.update(_2536_v5, _module.__default.safeIndex(p1, new BigNumber((_2536_v5).length)), _2538_v6));
        let _rhs377 = ((new BigNumber((_2539_v7).length)).plus(new BigNumber((_dafny.Seq.of((_2526_v1)[_module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length))])).length))).isLessThan((_this).f31);
        let _rhs378 = _this.f30;
        let _rhs379 = (_2540_v8).dtor_cf34;
        let _rhs380 = _2535_v3;
        let _lhs348 = globalState;
        let _lhs349 = globalState;
        let _lhs350 = _2526_v1;
        let _lhs351 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_2526_v1).length));
        let _lhs352 = globalState;
        _lhs348.f24 = _rhs376;
        _lhs349.f15 = _rhs377;
        _lhs350[_lhs351] = _rhs378;
        _lhs352.f12 = _rhs379;
        _2535_v3 = _rhs380;
        (globalState).f12 = _module.__default.fm39(globalState);
        (globalState).f12 = _2525_i0;
      }
      let _2541_v9;
      _2541_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,!((_this).fm2(_this.f30, (_this).f31, globalState)));
      let _2542_v10;
      _2542_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2541_v9).length),!((_this).f26));
      let _2543_v11;
      _2543_v11 = _dafny.Seq.UnicodeFromString("tvcq");
      let _2544_v12;
      _2544_v12 = _dafny.Map.Empty.slice().updateUnsafe(((((_2541_v9).contains(new BigNumber(-801))) ? ((_2541_v9).get(new BigNumber(-801))) : ((_this).fm2(_this.f30, p1, globalState)))) || (_this.f30),_2543_v11);
      (globalState).f6 = (((_2544_v12).contains((_this).f27)) ? ((_2544_v12).get((_this).f27)) : (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qntndfgm"), _dafny.Seq.UnicodeFromString("jkyhfrdrq"))));
      let _2545_v13;
      _2545_v13 = _module.D1.create_DC3(_2524_v0);
      if (function (_source35) {
        if (_source35.is_DC4) {
          let _2546___mcc_h0 = (_source35).cf7;
          let _2547___mcc_h1 = (_source35).cf8;
          let _2548___mcc_h2 = (_source35).cf9;
          let _2549___mcc_h3 = (_source35).cf10;
          let _2550___mcc_h4 = (_source35).cf11;
          let _2551_cf11 = _2550___mcc_h4;
          let _2552_cf10 = _2549___mcc_h3;
          let _2553_cf9 = _2548___mcc_h2;
          let _2554_cf8 = _2547___mcc_h1;
          let _2555_cf7 = _2546___mcc_h0;
          return ((_this).f27) === (_this.f30);
        } else if (_source35.is_DC5) {
          let _2556___mcc_h5 = (_source35).cf12;
          let _2557_cf12 = _2556___mcc_h5;
          return (_this).f27;
        } else {
          let _2558___mcc_h6 = (_source35).cf6;
          let _2559_cf6 = _2558___mcc_h6;
          return (_this).f26;
        }
      }(_2545_v13)) {
        if (!(((_dafny.ZERO).minus(p1)).isLessThan((p3).multipliedBy(new BigNumber(737))))) {
          let _2560_v14;
          let _init75 = function (_2561_i3) {
            return _this.f30;
          };
          let _nw402 = Array((new BigNumber(25)).toNumber());
          for (let _i0_75 = 0; _i0_75 < new BigNumber(_nw402.length); _i0_75++) {
            _nw402[_i0_75] = _init75(new BigNumber(_i0_75));
          }
          _2560_v14 = _nw402;
          let _2562_v15;
          _2562_v15 = _2560_v14;
          _2562_v15 = _2562_v15;
          (globalState).f15 = !(((_this).f26) && (_this.f30)) || (((_this).f27) && (_this.f30));
          let _2563_v16;
          _2563_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(-67), p3),p2);
          _2563_v16 = (_2563_v16).update(p1, p2);
          let _2564_v17;
          _2564_v17 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(p3, (_this).f31), p1);
          _2564_v17 = (_this).f29;
          let _rhs381 = _2543_v11;
          let _rhs382 = _2543_v11;
          let _rhs383 = (_this).f31;
          let _rhs384 = p2;
          let _rhs385 = p3;
          let _lhs353 = globalState;
          let _lhs354 = globalState;
          let _lhs355 = globalState;
          let _lhs356 = globalState;
          let _lhs357 = globalState;
          _lhs353.f6 = _rhs381;
          _lhs354.f6 = _rhs382;
          _lhs355.f16 = _rhs383;
          _lhs356.f19 = _rhs384;
          _lhs357.f12 = _rhs385;
        } else {
          let _2565_v18;
          _2565_v18 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p1));
          let _2566_v19;
          let _nw403 = new _module.C3();
          _nw403.__ctor(_2565_v18, (_this).f26, _this.f30);
          _2566_v19 = _nw403;
          let _2567_v20;
          _2567_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2566_v19);
          let _2568_v21;
          _2568_v21 = _dafny.Set.fromElements(_2567_v20, _2567_v20, (_dafny.Map.Empty.slice().updateUnsafe(p3,_2566_v19)).Merge(_2567_v20), _2567_v20);
          let _2569_v22;
          let _nw404 = Array((new BigNumber(6)).toNumber()).fill([]);
          _2569_v22 = _nw404;
          let _2570_v23;
          let _nw405 = Array((new BigNumber(7)).toNumber()).fill(false);
          _2570_v23 = _nw405;
          let _index367 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_2569_v22).length));
          (_2569_v22)[_index367] = _2570_v23;
          let _2571_v24;
          _2571_v24 = _dafny.Seq.of(p0, p0);
          let _2572_v25;
          _2572_v25 = _module.D10.create_DC24(_2543_v11, _2571_v24, _this.f30);
          let _index368 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_2569_v22).length));
          let _rhs386 = _2568_v21;
          let _rhs387 = (_module.__default.safeDivisionInt((_this).f31, (_this).f31)).minus((p3).multipliedBy(p1));
          let _rhs388 = _dafny.Seq.contains((_2572_v25).dtor_cf44, p2);
          let _rhs389 = (_this).f26;
          let _rhs390 = _2570_v23;
          let _lhs358 = globalState;
          let _lhs359 = globalState;
          let _lhs360 = globalState;
          let _lhs361 = _2569_v22;
          let _lhs362 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_2569_v22).length));
          _2568_v21 = _rhs386;
          _lhs358.f12 = _rhs387;
          _lhs359.f15 = _rhs388;
          _lhs360.f15 = _rhs389;
          _lhs361[_lhs362] = _rhs390;
          let _2573_v26;
          _2573_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f27);
          _2573_v26 = (_2573_v26).update(p2, _this.f30);
          let _index369 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_2569_v22).length));
          let _rhs391 = (_this).f31;
          let _rhs392 = (_this).f26;
          let _rhs393 = _2570_v23;
          let _lhs363 = globalState;
          let _lhs364 = globalState;
          let _lhs365 = _2569_v22;
          let _lhs366 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_2569_v22).length));
          _lhs363.f12 = _rhs391;
          _lhs364.f2 = _rhs392;
          _lhs365[_lhs366] = _rhs393;
          (globalState).f15 = !(false);
          (globalState).f16 = p3;
        }
        let _2574_v27;
        let _nw406 = Array((new BigNumber(25)).toNumber()).fill(false);
        _2574_v27 = _nw406;
        _2574_v27 = _2574_v27;
        let _2575_v28;
        _2575_v28 = _dafny.Map.Empty.slice().updateUnsafe(_2543_v11,_dafny.Seq.Create(_module.__default.abs(new BigNumber(932)), ((_2576_p2) => function (_2577_i4) {
          return _2576_p2;
        })(p2)));
        (globalState).f12 = new BigNumber(((_2575_v28).Merge((_2575_v28).update(_2543_v11, _2543_v11))).length);
        let _index370 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_2574_v27).length));
        (_2574_v27)[_index370] = false;
        let _index371 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_2574_v27).length));
        (_2574_v27)[_index371] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("ioovj"), _2543_v11);
        let _2578_v29;
        _2578_v29 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f31), (_this).f31, p1);
        let _2579_v30;
        _2579_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2578_v29,p3);
        let _2580_v32;
        _2580_v32 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), ((_2581_p1) => function (_2582_i5) {
          return _2581_p1;
        })(p1));
        let _2583_v33;
        _2583_v33 = _dafny.Seq.of(_2578_v29, (_2580_v32));
        _2579_v30 = function () {
          let _coll75 = new _dafny.Map();
          for (const _compr_75 of (_2583_v33).Elements) {
            let _2584_v31 = _compr_75;
            if (_dafny.Seq.contains(_2583_v33, _2584_v31)) {
              _coll75.push([_2584_v31,p3]);
            }
          }
          return _coll75;
        }();
      } else {
        (globalState).f24 = (_this).f27;
        let _2585_v34;
        _2585_v34 = _dafny.MultiSet.fromElements(_this.f30);
        let _2586_v35;
        _2586_v35 = _dafny.Seq.of(p1, new BigNumber(72));
        let _2587_v36;
        _2587_v36 = _dafny.Map.Empty.slice().updateUnsafe((_2586_v35)[_module.__default.safeIndex(p3, new BigNumber((_2586_v35).length))],p0);
        let _2588_v37;
        let _nw407 = new _module.C7();
        _nw407.__ctor(true, _2587_v36);
        _2588_v37 = _nw407;
        let _2589_v38;
        _2589_v38 = _dafny.Set.fromElements(p1);
        let _rhs394 = (_2585_v34).update((_this).f26, _module.__default.abs(new BigNumber((_2589_v38).length)));
        let _rhs395 = _dafny.Seq.IsPrefixOf(_2586_v35, _2586_v35);
        let _rhs396 = _2588_v37;
        let _lhs367 = globalState;
        _2585_v34 = _rhs394;
        _lhs367.f22 = _rhs395;
        _2588_v37 = _rhs396;
        let _2590_v39;
        _2590_v39 = _dafny.Seq.of(p0);
        let _2591_v40;
        _2591_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f30,((_dafny.ZERO).minus(p3)).minus((_module.D21.create_DC52((_this).f31, _2590_v39)).dtor_cf95));
        let _2592_v41;
        _2592_v41 = _module.D23.create_DC57();
        let _2593_v42;
        _2593_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.D23.create_DC58(_2592_v41),(_this).f27);
        let _2594_v43;
        _2594_v43 = _module.D23.create_DC58(_2592_v41);
        let _2595_v44;
        _2595_v44 = _module.D10.create_DC24(_2543_v11, _2590_v39, false);
        _2591_v40 = (_2591_v40).update((((_this).f26) ? ((_this).f27) : ((((_2593_v42).contains(_2594_v43)) ? ((_2593_v42).get(_2594_v43)) : ((_this).f26)))), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_2543_v11).length), new BigNumber((_2543_v11).length), new BigNumber(((_2595_v44).dtor_cf44).length), p1, new BigNumber((_2524_v0).length))).cardinality()));
        (globalState).f2 = (_2589_v38).IsProperSubsetOf(_dafny.Set.fromElements((_this).f31, new BigNumber(-14), (_this).f31, new BigNumber(-503), (_this).f31));
        (globalState).f16 = (_this).f31;
      }
      let _2596_v45;
      let _nw408 = new _module.C8();
      _nw408.__ctor(_dafny.MultiSet.fromElements(p3));
      _2596_v45 = _nw408;
      let _2597_v46;
      let _init76 = function (_2598_i7) {
        return _this.f30;
      };
      let _nw409 = Array((new BigNumber(13)).toNumber());
      for (let _i0_76 = 0; _i0_76 < new BigNumber(_nw409.length); _i0_76++) {
        _nw409[_i0_76] = _init76(new BigNumber(_i0_76));
      }
      _2597_v46 = _nw409;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2597_v46).length))) {
        let _2599_i6 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2599_i6)) && ((_2599_i6).isLessThan(new BigNumber((_2597_v46).length))))) {
          (_2597_v46)[(_2599_i6)] = (_this.f30) && ((((_2542_v10).contains(new BigNumber((_2543_v11).length))) ? ((_2542_v10).get(new BigNumber((_2543_v11).length))) : ((_this).f26)));
        }
      }
      let _2600_v47;
      _2600_v47 = _dafny.Seq.of(p3);
      let _2601_v48;
      _2601_v48 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(922), (_this).fm1(globalState)), _dafny.Seq.Concat(_2600_v47, _2600_v47));
      _2601_v48 = _2601_v48;
      r0 = p3;
      r1 = (_this).fm1(globalState);
      r2 = !((p1).isEqualTo(p3));
      return [r0, r1, r2];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
