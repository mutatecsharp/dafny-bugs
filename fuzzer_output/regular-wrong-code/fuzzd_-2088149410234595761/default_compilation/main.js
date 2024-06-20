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
    static fm0(p0, p1, globalState) {
      return _module.__default.safeModuloInt(new BigNumber(833), new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(565), new BigNumber(642))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(565)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(642)))) {
            _coll0.push([_module.__default.safeDivisionInt(_0_v0, new BigNumber((_dafny.Seq.UnicodeFromString("doerfd")).length)),new BigNumber((function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))))).Elements) {
                let _1_v1 = _compr_1;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0))))).contains(_1_v1)) {
                  _coll1.push([_1_v1,new BigNumber(24)]);
                }
              }
              return _coll1;
            }()).length)]);
          }
        }
        return _coll0;
      }()).length));
    };
    static fm1(globalState) {
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(-329), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))), new BigNumber(696), new BigNumber((_dafny.Set.fromElements(true, true, !(true))).length))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_2_i0) {
        return _dafny.Seq.UnicodeFromString("omngle");
      })).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-597)),!(false))).length))));
    };
    static fm4(p0, globalState) {
      return (((true) ? (new BigNumber(-92)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(((_dafny.Seq.Create(_module.__default.abs(new BigNumber(605)), function (_3_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }))).length)))).cardinality()))).length)))).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("esnepq")).length), new BigNumber(414)));
    };
    static fm5(p0, p1, p2, globalState) {
      return _dafny.Seq.of(false, (new BigNumber((_dafny.Seq.UnicodeFromString("aahphab")).length)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-978)))), !(true), (new BigNumber(-992)).isLessThanOrEqualTo(new BigNumber(-747)), (true) && (!(!(true))));
    };
    static fm6(p0, p1, globalState) {
      let _source0 = _dafny.Seq.UnicodeFromString("jkgbfxm");
      let _4___mcc_h0 = _source0;
      let _5_cf5 = _4___mcc_h0;
      return _5_cf5;
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true, true, !(false), true)).cardinality()), new BigNumber((_dafny.Seq.of(true, true)).length)), _dafny.Seq.of(new BigNumber(-334))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("cudaub")).length)), _dafny.Seq.of(new BigNumber(650))));
    };
    static fm8(globalState) {
      return ((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(true, false, true))).Difference(_dafny.MultiSet.fromElements(true, true));
    };
    static m0(p0, p1, globalState) {
      let r0 = [];
      let r1 = _dafny.ZERO;
      (globalState).f2 = p0;
      let _6_v0;
      _6_v0 = false;
      let _7_v1;
      _7_v1 = _dafny.Seq.of(_6_v0, _6_v0, _6_v0);
      _7_v1 = _7_v1;
      let _8_v2;
      _8_v2 = _dafny.Set.fromElements(p1);
      let _9_i0;
      _9_i0 = _dafny.ZERO;
      L0: {
        while (!(!(((_8_v2).IsSubsetOf(_8_v2)) || (true)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_9_i0)) {
              break L0;
            }
            _9_i0 = (_9_i0).plus(_dafny.ONE);
            let _10_v3;
            _10_v3 = new _dafny.CodePoint('r'.codePointAt(0));
            let _11_v4;
            _11_v4 = _dafny.Map.Empty.slice().updateUnsafe(_6_v0,_10_v3);
            let _12_v5;
            let _nw0 = Array((new BigNumber(15)).toNumber());
            _nw0[(_dafny.ZERO).toNumber()] = _10_v3;
            _nw0[(_dafny.ONE).toNumber()] = _module.__default.fm1(globalState);
            _nw0[(new BigNumber(2)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(3)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(4)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(5)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(6)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(7)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(8)).toNumber()] = ((_6_v0) ? (_10_v3) : (_10_v3));
            _nw0[(new BigNumber(9)).toNumber()] = (((_11_v4).contains(_6_v0)) ? ((_11_v4).get(_6_v0)) : (new _dafny.CodePoint('b'.codePointAt(0))));
            _nw0[(new BigNumber(10)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(11)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(12)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(13)).toNumber()] = _10_v3;
            _nw0[(new BigNumber(14)).toNumber()] = _10_v3;
            _12_v5 = _nw0;
            let _index0 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_12_v5).length));
            (_12_v5)[_index0] = _10_v3;
            let _index1 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_12_v5).length));
            (_12_v5)[_index1] = _10_v3;
            r1 = _module.__default.fm0(_6_v0, p0, globalState);
            let _13_v6;
            let _nw1 = new _module.C0();
            _nw1.__ctor();
            _13_v6 = _nw1;
            let _14_v7;
            _14_v7 = _dafny.MultiSet.fromElements(true, _6_v0, _6_v0);
            _6_v0 = !(new BigNumber(((_14_v7).Union(_14_v7)).cardinality())).isEqualTo(p0);
          }
        }
      }
      let _15_v8;
      _15_v8 = _dafny.Seq.UnicodeFromString("bhpgj");
      let _hi0 = _module.__default.safeDivisionInt(new BigNumber(380), new BigNumber(60));
      for (let _16_i1 = new BigNumber((_15_v8).length); _16_i1.isLessThan(_hi0); _16_i1 = _16_i1.plus(_dafny.ONE)) {
        let _17_v9;
        let _nw2 = Array((_dafny.ONE).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _16_i1;
        _17_v9 = _nw2;
        let _18_v10;
        _18_v10 = _dafny.Seq.of(_17_v9, _17_v9, _17_v9);
        (globalState).f2 = new BigNumber((_18_v10).length);
        _15_v8 = _dafny.Seq.Concat(_15_v8, _15_v8);
        let _19_v11;
        _19_v11 = new _dafny.CodePoint('n'.codePointAt(0));
        _15_v8 = _dafny.Seq.Concat(_15_v8, _dafny.Seq.update(_module.__default.fm6(!(false), _6_v0, globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm6(!(false), _6_v0, globalState)).length)), _19_v11));
        if (_6_v0) {
          let _20_v12;
          _20_v12 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(830)).multipliedBy(_16_i1),_15_v8);
          _20_v12 = (_20_v12).update(p1, _15_v8);
          (globalState).f2 = p1;
          let _21_v13;
          let _nw3 = new _module.C0();
          _nw3.__ctor();
          _21_v13 = _nw3;
          let _22_v14;
          let _nw4 = Array((new BigNumber(24)).toNumber());
          _nw4[(_dafny.ZERO).toNumber()] = _21_v13;
          _nw4[(_dafny.ONE).toNumber()] = _21_v13;
          _nw4[(new BigNumber(2)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(3)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(4)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(5)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(6)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(7)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(8)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(9)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(10)).toNumber()] = ((_6_v0) ? (_21_v13) : (_21_v13));
          _nw4[(new BigNumber(11)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(12)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(13)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(14)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(15)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(16)).toNumber()] = ((_6_v0) ? (_21_v13) : (_21_v13));
          _nw4[(new BigNumber(17)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(18)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(19)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(20)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(21)).toNumber()] = _21_v13;
          _nw4[(new BigNumber(22)).toNumber()] = ((_6_v0) ? (_21_v13) : (_21_v13));
          _nw4[(new BigNumber(23)).toNumber()] = _21_v13;
          _22_v14 = _nw4;
          let _nw5 = Array((new BigNumber(27)).toNumber());
          _22_v14 = _nw5;
          let _23_v15;
          _23_v15 = _dafny.MultiSet.fromElements(_6_v0, _6_v0);
          (globalState).f2 = (((_23_v15).contains(!(!_dafny.Seq.contains(_15_v8, _19_v11)))) ? ((_23_v15).get(!(!_dafny.Seq.contains(_15_v8, _19_v11)))) : (_16_i1));
          let _24_v16;
          _24_v16 = _dafny.Seq.of(p1, _module.__default.safeDivisionInt(p0, p0));
          _24_v16 = _module.__default.fm7(_6_v0, ((_6_v0) ? (new BigNumber(569)) : (_16_i1)), globalState);
        } else {
          let _25_v17;
          let _init0 = ((_26_v8) => function (_27_i2) {
            return _26_v8;
          })(_15_v8);
          let _nw6 = Array((new BigNumber(16)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw6.length); _i0_0++) {
            _nw6[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _25_v17 = _nw6;
          let _28_v18;
          _28_v18 = _module.D0.create_DC2(_6_v0, _25_v17, p0);
          let _29_v19;
          _29_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_28_v18).dtor_cf2);
          let _30_v20;
          _30_v20 = _dafny.MultiSet.fromElements(_16_i1, _16_i1, p1);
          _6_v0 = (((_29_v19).contains(_module.__default.safeModuloInt(_16_i1, new BigNumber((_dafny.Seq.update(_15_v8, _module.__default.safeIndex((((_30_v20).contains(new BigNumber((_15_v8).length))) ? ((_30_v20).get(new BigNumber((_15_v8).length))) : (p0)), new BigNumber((_15_v8).length)), _19_v11)).length)))) ? ((_29_v19).get(_module.__default.safeModuloInt(_16_i1, new BigNumber((_dafny.Seq.update(_15_v8, _module.__default.safeIndex((((_30_v20).contains(new BigNumber((_15_v8).length))) ? ((_30_v20).get(new BigNumber((_15_v8).length))) : (p0)), new BigNumber((_15_v8).length)), _19_v11)).length)))) : (_6_v0));
          let _31_v21;
          _31_v21 = _dafny.Seq.of(_module.__default.fm0(_6_v0, p0, globalState), p1);
          let _32_v22;
          _32_v22 = _dafny.Set.fromElements(_19_v11, _19_v11);
          let _33_v23;
          _33_v23 = _dafny.Seq.of(_31_v21);
          let _34_v24;
          _34_v24 = _module.D0.create_DC1(p0, _dafny.Seq.update(_31_v21, _module.__default.safeIndex(new BigNumber((_32_v22).length), new BigNumber((_31_v21).length)), new BigNumber((_33_v23).length)));
          (globalState).f2 = (_34_v24).dtor_cf0;
          _6_v0 = _6_v0;
          let _35_v25;
          _35_v25 = _module.D0.create_DC0();
          let _rhs0 = _16_i1;
          let _rhs1 = _35_v25;
          let _lhs0 = globalState;
          _lhs0.f2 = _rhs0;
          _35_v25 = _rhs1;
          let _36_v26;
          _36_v26 = _dafny.MultiSet.fromElements(_6_v0);
          _6_v0 = ((_module.__default.fm8(globalState)).Intersect(_36_v26)).IsDisjointFrom((_36_v26).Union(_dafny.MultiSet.fromElements(_6_v0)));
        }
      }
      let _37_v27;
      let _nw7 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _37_v27 = _nw7;
      let _38_v28;
      _38_v28 = _dafny.Seq.of(p0, p0);
      let _39_v29;
      _39_v29 = _module.D0.create_DC2(_6_v0, _37_v27, (_module.D0.create_DC1(p1, _38_v28)).dtor_cf0);
      let _40_i3;
      _40_i3 = _dafny.ZERO;
      L1: {
        while (!((_39_v29).dtor_cf2)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_40_i3)) {
              break L1;
            }
            _40_i3 = (_40_i3).plus(_dafny.ONE);
            let _41_v30;
            let _nw8 = Array((_dafny.ONE).toNumber());
            _41_v30 = _nw8;
            let _42_v31;
            let _nw9 = new _module.C0();
            _nw9.__ctor();
            _42_v31 = _nw9;
            let _index2 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_41_v30).length));
            (_41_v30)[_index2] = _42_v31;
            let _index3 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_41_v30).length));
            (_41_v30)[_index3] = _42_v31;
            let _43_v32;
            _43_v32 = _module.D0.create_DC1(p1, _38_v28);
            (globalState).f2 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_6_v0, _6_v0)).length), (p1).minus((_43_v32).dtor_cf0));
            _38_v28 = _38_v28;
            let _44_v33;
            _44_v33 = _dafny.Map.Empty.slice().updateUnsafe(_38_v28,_dafny.Seq.Concat(_15_v8, _dafny.Seq.UnicodeFromString("odihkxdux")));
            r1 = new BigNumber((_44_v33).length);
          }
        }
      }
      let _45_v35;
      let _init1 = ((_46_v8) => function (_47_i5) {
        return (_dafny.Set.fromElements(new _dafny.CodePoint('b'.codePointAt(0)))).Difference(function () {
          let _coll2 = new _dafny.Set();
          for (const _compr_2 of (_dafny.MultiSet.FromArray(_46_v8)).Elements) {
            let _48_v34 = _compr_2;
            if ((_dafny.MultiSet.FromArray(_46_v8)).contains(_48_v34)) {
              _coll2.add(_48_v34);
            }
          }
          return _coll2;
        }());
      })(_15_v8);
      let _nw10 = Array((new BigNumber(29)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw10.length); _i0_1++) {
        _nw10[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _45_v35 = _nw10;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_45_v35).length))) {
        let _49_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_49_i4)) && ((_49_i4).isLessThan(new BigNumber((_45_v35).length))))) {
          (_45_v35)[(_49_i4)] = (_dafny.Set.fromElements(new _dafny.CodePoint('b'.codePointAt(0)))).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0))));
        }
      }
      let _50_v36;
      let _nw11 = Array((new BigNumber(19)).toNumber()).fill(false);
      _50_v36 = _nw11;
      r0 = _50_v36;
      r1 = p1;
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _51_v0;
      _51_v0 = _dafny.Seq.UnicodeFromString("rhvwcv");
      let _52_globalState;
      let _nw12 = new _module.GlobalState();
      _nw12.__ctor(_dafny.Seq.Concat(_51_v0, _51_v0), true, new BigNumber(411), _51_v0);
      _52_globalState = _nw12;
      let _53_v1;
      _53_v1 = new BigNumber(387);
      (_52_globalState).f2 = _53_v1;
      let _54_v2;
      _54_v2 = false;
      _54_v2 = _54_v2;
      if (_54_v2) {
        let _55_v3;
        let _56_v4;
        let _out0;
        let _out1;
        let _outcollector0 = _module.__default.m0(_53_v1, (_53_v1).minus(_53_v1), _52_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _55_v3 = _out0;
        _56_v4 = _out1;
        let _57_v5;
        _57_v5 = _dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(-8), _56_v4));
        let _58_v6;
        _58_v6 = _dafny.Seq.of(_57_v5, _57_v5);
        _54_v2 = ((_58_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(_56_v4), new BigNumber((_58_v6).length))]).IsProperSubsetOf(_57_v5);
        let _59_v7;
        let _init2 = ((_60_v2, _61_v0) => function (_62_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe(_60_v2,new BigNumber((_61_v0).length));
        })(_54_v2, _51_v0);
        let _nw13 = Array((new BigNumber(16)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw13.length); _i0_2++) {
          _nw13[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _59_v7 = _nw13;
        let _index4 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_55_v3).length));
        (_55_v3)[_index4] = _54_v2;
        let _63_v8;
        _63_v8 = _dafny.Map.Empty.slice().updateUnsafe(_53_v1,_module.__default.safeModuloInt(new BigNumber((_51_v0).length), _56_v4));
        let _64_v9;
        _64_v9 = _dafny.Set.fromElements(_53_v1, _56_v4, _53_v1);
        let _index5 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_55_v3).length));
        let _rhs2 = new BigNumber((_63_v8).length);
        let _rhs3 = _59_v7;
        let _rhs4 = (_64_v9).IsDisjointFrom((_64_v9).Intersect(_64_v9));
        let _lhs1 = _52_globalState;
        let _lhs2 = _55_v3;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_55_v3).length));
        _lhs1.f2 = _rhs2;
        _59_v7 = _rhs3;
        _lhs2[_lhs3] = _rhs4;
        let _65_v10;
        let _66_v11;
        let _out2;
        let _out3;
        let _outcollector1 = _module.__default.m0(_module.__default.safeDivisionInt(_53_v1, _56_v4), _53_v1, _52_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _65_v10 = _out2;
        _66_v11 = _out3;
        _51_v0 = _dafny.Seq.UnicodeFromString("nwstnawx");
      } else {
        let _67_v12;
        let _nw14 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _67_v12 = _nw14;
        let _68_v13;
        _68_v13 = _dafny.Map.Empty.slice().updateUnsafe(_53_v1,_67_v12);
        let _69_v14;
        let _nw15 = Array((_dafny.ONE).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = (((_68_v13).contains(_53_v1)) ? ((_68_v13).get(_53_v1)) : (_67_v12));
        _69_v14 = _nw15;
        let _index6 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_69_v14).length));
        (_69_v14)[_index6] = _67_v12;
        let _index7 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_69_v14).length));
        (_69_v14)[_index7] = _67_v12;
        (_52_globalState).f2 = _53_v1;
        let _70_v15;
        let _71_v16;
        let _out4;
        let _out5;
        let _outcollector2 = _module.__default.m0(_53_v1, (_dafny.ZERO).minus(new BigNumber(-856)), _52_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _70_v15 = _out4;
        _71_v16 = _out5;
        let _index8 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_69_v14).length));
        (_69_v14)[_index8] = _67_v12;
        let _72_v17;
        _72_v17 = _dafny.Map.Empty.slice().updateUnsafe(_70_v15,(_71_v16).isLessThanOrEqualTo(_71_v16));
        _72_v17 = (_72_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(_70_v15,false));
      }
      let _hi1 = _53_v1;
      for (let _73_i1 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("suibphoto"), _51_v0)).length); _73_i1.isLessThan(_hi1); _73_i1 = _73_i1.plus(_dafny.ONE)) {
        let _74_v18;
        let _init3 = ((_75_v1, _76_i1) => function (_77_i2) {
          return !(_75_v1).isEqualTo(_76_i1);
        })(_53_v1, _73_i1);
        let _nw16 = Array((new BigNumber(12)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw16.length); _i0_3++) {
          _nw16[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _74_v18 = _nw16;
        _74_v18 = _74_v18;
        _53_v1 = _73_i1;
        let _78_v19;
        _78_v19 = _dafny.Seq.of(false, _54_v2);
        let _79_v20;
        let _80_v21;
        let _out6;
        let _out7;
        let _outcollector3 = _module.__default.m0((new BigNumber(663)).multipliedBy(_module.__default.fm0(_54_v2, new BigNumber((_78_v19).length), _52_globalState)), _73_i1, _52_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _79_v20 = _out6;
        _80_v21 = _out7;
        let _rhs5 = _54_v2;
        let _rhs6 = new BigNumber(-337);
        let _lhs4 = _52_globalState;
        _54_v2 = _rhs5;
        _lhs4.f2 = _rhs6;
      }
      let _81_v22;
      _81_v22 = _dafny.Set.fromElements(_54_v2, _54_v2, true);
      let _82_v23;
      let _nw17 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _82_v23 = _nw17;
      let _83_v24;
      _83_v24 = _module.D0.create_DC2(!(_81_v22).equals(_81_v22), _82_v23, _53_v1);
      let _source1 = _83_v24;
      if (_source1.is_DC0) {
        let _84_v25;
        _84_v25 = _dafny.Map.Empty.slice().updateUnsafe(_53_v1,_54_v2);
        let _85_v26;
        _85_v26 = _dafny.Seq.of(_54_v2);
        let _86_v27;
        _86_v27 = _dafny.Map.Empty.slice().updateUnsafe(_85_v26,true);
        let _87_v28;
        let _nw18 = Array((new BigNumber(12)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = _53_v1;
        _nw18[(_dafny.ONE).toNumber()] = _module.__default.fm0(_54_v2, _53_v1, _52_globalState);
        _nw18[(new BigNumber(2)).toNumber()] = new BigNumber(-526);
        _nw18[(new BigNumber(3)).toNumber()] = new BigNumber(-912);
        _nw18[(new BigNumber(4)).toNumber()] = _53_v1;
        _nw18[(new BigNumber(5)).toNumber()] = _53_v1;
        _nw18[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-894));
        _nw18[(new BigNumber(7)).toNumber()] = _53_v1;
        _nw18[(new BigNumber(8)).toNumber()] = (_53_v1).minus(new BigNumber((_84_v25).length));
        _nw18[(new BigNumber(9)).toNumber()] = new BigNumber((_86_v27).length);
        _nw18[(new BigNumber(10)).toNumber()] = ((true) ? (_53_v1) : (_53_v1));
        _nw18[(new BigNumber(11)).toNumber()] = _53_v1;
        _87_v28 = _nw18;
        let _88_v29;
        let _init4 = function (_89_i3) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        };
        let _nw19 = Array((new BigNumber(5)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw19.length); _i0_4++) {
          _nw19[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _88_v29 = _nw19;
        let _90_v30;
        _90_v30 = _dafny.Map.Empty.slice().updateUnsafe(_88_v29,_53_v1);
        let _91_v31;
        _91_v31 = _dafny.Map.Empty.slice().updateUnsafe(_53_v1,new BigNumber((_90_v30).length));
        let _92_v32;
        _92_v32 = _dafny.Map.Empty.slice().updateUnsafe(true,_91_v31);
        let _index9 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length));
        (_87_v28)[_index9] = new BigNumber(((((_92_v32).contains(_54_v2)) ? ((_92_v32).get(_54_v2)) : (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(955), new BigNumber(328))) {
            let _93_v33 = _compr_3;
            if (((new BigNumber(955)).isLessThanOrEqualTo(_93_v33)) && ((_93_v33).isLessThan(new BigNumber(328)))) {
              _coll3.push([(_93_v33).multipliedBy(_53_v1),_53_v1]);
            }
          }
          return _coll3;
        }()))).length);
        let _index10 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length));
        (_87_v28)[_index10] = _module.__default.fm0(_54_v2, new BigNumber(348), _52_globalState);
        let _94_v34;
        _94_v34 = new _dafny.CodePoint('o'.codePointAt(0));
        let _95_v35;
        let _nw20 = Array((new BigNumber(4)).toNumber()).fill(false);
        _95_v35 = _nw20;
        let _index11 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_95_v35).length));
        (_95_v35)[_index11] = _54_v2;
        let _96_v36;
        _96_v36 = _dafny.Map.Empty.slice().updateUnsafe(((_54_v2) ? (_83_v24) : (_module.D0.create_DC2(false, _82_v23, _53_v1))),_94_v34);
        let _index12 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length));
        let _index13 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_95_v35).length));
        let _rhs7 = _54_v2;
        let _rhs8 = _module.__default.safeModuloInt((_87_v28)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length))], (_87_v28)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length))]);
        let _rhs9 = (((_96_v36).contains(((false) ? (_module.D0.create_DC2(_54_v2, _82_v23, (_87_v28)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length))])) : (_83_v24)))) ? ((_96_v36).get(((false) ? (_module.D0.create_DC2(_54_v2, _82_v23, (_87_v28)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length))])) : (_83_v24)))) : (_module.__default.fm1(_52_globalState)));
        let _rhs10 = !(((!(_54_v2)) ? (_54_v2) : (_54_v2))) || (_54_v2);
        let _lhs5 = _87_v28;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length));
        let _lhs7 = _95_v35;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(709), new BigNumber((_95_v35).length));
        _54_v2 = _rhs7;
        _lhs5[_lhs6] = _rhs8;
        _94_v34 = _rhs9;
        _lhs7[_lhs8] = _rhs10;
        let _97_v37;
        let _nw21 = new _module.C0();
        _nw21.__ctor();
        _97_v37 = _nw21;
        let _98_v38;
        _98_v38 = _dafny.Seq.of(new BigNumber(-296), _53_v1);
        let _rhs11 = _51_v0;
        let _rhs12 = (_98_v38)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_87_v28)[_module.__default.safeIndex(new BigNumber(377), new BigNumber((_87_v28).length))], _53_v1), new BigNumber((_98_v38).length))];
        let _lhs9 = _52_globalState;
        _51_v0 = _rhs11;
        _lhs9.f2 = _rhs12;
      } else if (_source1.is_DC1) {
        let _99___mcc_h0 = (_source1).cf0;
        let _100___mcc_h1 = (_source1).cf1;
        let _101_cf1 = _100___mcc_h1;
        let _102_cf0 = _99___mcc_h0;
        let _103_v39;
        let _nw22 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _103_v39 = _nw22;
        let _index14 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_103_v39).length));
        (_103_v39)[_index14] = _53_v1;
        let _index15 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_103_v39).length));
        (_103_v39)[_index15] = (_102_cf0).minus(_102_cf0);
        _54_v2 = _54_v2;
        _103_v39 = _103_v39;
        _54_v2 = _54_v2;
      } else {
        let _104___mcc_h2 = (_source1).cf2;
        let _105___mcc_h3 = (_source1).cf3;
        let _106___mcc_h4 = (_source1).cf4;
        let _107_cf4 = _106___mcc_h4;
        let _108_cf3 = _105___mcc_h3;
        let _109_cf2 = _104___mcc_h2;
        _51_v0 = _51_v0;
        _109_cf2 = (_83_v24).dtor_cf2;
        (_52_globalState).f2 = (_53_v1).plus((_53_v1).minus(_53_v1));
        let _110_v41;
        let _nw23 = new _module.C0();
        _nw23.__ctor();
        _110_v41 = _nw23;
        let _111_v42;
        _111_v42 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(927), new BigNumber(-180))) {
            let _112_v40 = _compr_4;
            if (((new BigNumber(927)).isLessThanOrEqualTo(_112_v40)) && ((_112_v40).isLessThan(new BigNumber(-180)))) {
              _coll4.push([_module.__default.safeDivisionInt(_112_v40, _53_v1),_54_v2]);
            }
          }
          return _coll4;
        }(),_110_v41);
        let _113_v43;
        let _114_v44;
        let _out8;
        let _out9;
        let _outcollector4 = _module.__default.m0(new BigNumber((_111_v42).length), (_dafny.ZERO).minus(((_109_cf2) ? (_53_v1) : (new BigNumber(506)))), _52_globalState);
        _out8 = _outcollector4[0];
        _out9 = _outcollector4[1];
        _113_v43 = _out8;
        _114_v44 = _out9;
      }
      let _115_v45;
      let _nw24 = new _module.C0();
      _nw24.__ctor();
      _115_v45 = _nw24;
      let _116_v46;
      _116_v46 = _dafny.MultiSet.fromElements(_115_v45, _115_v45);
      let _117_v47;
      _117_v47 = _dafny.Seq.of(_53_v1, _53_v1, _53_v1, new BigNumber((_116_v46).cardinality()));
      let _118_v48;
      _118_v48 = _dafny.MultiSet.fromElements(new BigNumber(-254), (_117_v47)[_module.__default.safeIndex(new BigNumber((_51_v0).length), new BigNumber((_117_v47).length))], _53_v1);
      let _hi2 = (((_118_v48).contains(new BigNumber(-779))) ? ((_118_v48).get(new BigNumber(-779))) : (_53_v1));
      for (let _119_i4 = _53_v1; _119_i4.isLessThan(_hi2); _119_i4 = _119_i4.plus(_dafny.ONE)) {
        let _120_v49;
        _120_v49 = _dafny.Seq.of(_54_v2);
        let _121_v50;
        _121_v50 = _dafny.Map.Empty.slice().updateUnsafe(_120_v49,_54_v2);
        let _122_v51;
        _122_v51 = new _dafny.CodePoint('s'.codePointAt(0));
        _51_v0 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_123_i5) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _51_v0), _module.__default.safeIndex((new BigNumber((_121_v50).length)).plus(_119_i4), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_124_i5) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _51_v0)).length)), _122_v51), _module.__default.safeIndex((_dafny.ZERO).minus(_53_v1), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_125_i5) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _51_v0), _module.__default.safeIndex((new BigNumber((_121_v50).length)).plus(_119_i4), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_126_i5) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _51_v0)).length)), _122_v51)).length)), new _dafny.CodePoint('l'.codePointAt(0)));
        let _127_v52;
        _127_v52 = _dafny.Map.Empty.slice().updateUnsafe((_53_v1).plus(_119_i4),new BigNumber(807));
        _127_v52 = (_127_v52).update((new BigNumber((_dafny.Set.fromElements(_119_i4, _119_i4, _53_v1)).length)).plus(_53_v1), _module.__default.safeDivisionInt(_119_i4, _53_v1));
        let _128_v53;
        let _nw25 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _128_v53 = _nw25;
        let _129_v54;
        _129_v54 = _dafny.Set.fromElements(_128_v53, _128_v53);
        _54_v2 = (_129_v54).IsSubsetOf(_129_v54);
        _54_v2 = (_120_v49)[_module.__default.safeIndex(_53_v1, new BigNumber((_120_v49).length))];
      }
      let _130_v55;
      let _nw26 = Array((new BigNumber(2)).toNumber()).fill([]);
      _130_v55 = _nw26;
      _130_v55 = _130_v55;
      let _131_v56;
      _131_v56 = _dafny.Seq.of(_54_v2, _54_v2);
      if (_dafny.Seq.contains(_117_v47, (_53_v1).multipliedBy(new BigNumber((_131_v56).length)))) {
        _53_v1 = ((false) ? (_53_v1) : ((_53_v1).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), function (_132_i6) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length))));
        let _133_v57;
        _133_v57 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_52_globalState),_54_v2);
        let _134_v58;
        _134_v58 = new _dafny.CodePoint('a'.codePointAt(0));
        _54_v2 = !((((_133_v57).contains(_134_v58)) ? ((_133_v57).get(_134_v58)) : (!(_dafny.areEqual(_51_v0, _51_v0)))));
        (_52_globalState).f2 = _53_v1;
        _83_v24 = _83_v24;
        let _135_v59;
        _135_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(993),_53_v1);
        _54_v2 = (new BigNumber(774)).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_51_v0).length))), new BigNumber((_135_v59).length)));
      } else {
        let _136_v60;
        let _nw27 = new _module.C0();
        _nw27.__ctor();
        _136_v60 = _nw27;
        _54_v2 = !(new BigNumber((_117_v47).length)).isEqualTo(_53_v1);
        let _137_v61;
        let _138_v62;
        let _out10;
        let _out11;
        let _outcollector5 = _module.__default.m0(_53_v1, _53_v1, _52_globalState);
        _out10 = _outcollector5[0];
        _out11 = _outcollector5[1];
        _137_v61 = _out10;
        _138_v62 = _out11;
        if (_54_v2) {
          let _139_v63;
          let _nw28 = Array((new BigNumber(25)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _module.__default.fm0(!(_54_v2), _53_v1, _52_globalState);
          _nw28[(_dafny.ONE).toNumber()] = _53_v1;
          _nw28[(new BigNumber(2)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(3)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(4)).toNumber()] = (_138_v62).plus(_53_v1);
          _nw28[(new BigNumber(5)).toNumber()] = (_138_v62).minus(_53_v1);
          _nw28[(new BigNumber(6)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(7)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(8)).toNumber()] = (new BigNumber(10)).plus(_53_v1);
          _nw28[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("fatyy")).length);
          _nw28[(new BigNumber(10)).toNumber()] = _138_v62;
          _nw28[(new BigNumber(11)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(12)).toNumber()] = _138_v62;
          _nw28[(new BigNumber(13)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(14)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(15)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(16)).toNumber()] = (_138_v62).multipliedBy((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(_54_v2, new BigNumber((_81_v22).length), _52_globalState))));
          _nw28[(new BigNumber(17)).toNumber()] = (new BigNumber((_131_v56).length)).multipliedBy(new BigNumber(7));
          _nw28[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_138_v62);
          _nw28[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(_138_v62);
          _nw28[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(_138_v62, new BigNumber((_dafny.Seq.UnicodeFromString("iasqrm")).length));
          _nw28[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_138_v62);
          _nw28[(new BigNumber(22)).toNumber()] = _53_v1;
          _nw28[(new BigNumber(23)).toNumber()] = (_117_v47)[_module.__default.safeIndex(_53_v1, new BigNumber((_117_v47).length))];
          _nw28[(new BigNumber(24)).toNumber()] = _138_v62;
          _139_v63 = _nw28;
          _139_v63 = _139_v63;
          _54_v2 = !((_118_v48).IsSubsetOf((_module.__default.fm3(_dafny.Seq.of(_53_v1), _54_v2, _54_v2, _54_v2, _52_globalState)).Difference(_118_v48)));
          let _index16 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_139_v63).length));
          (_139_v63)[_index16] = _53_v1;
          let _index17 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_139_v63).length));
          (_139_v63)[_index17] = new BigNumber((_dafny.Set.fromElements(((_54_v2) ? (_54_v2) : (true)), ((_118_v48).update(_138_v62, _module.__default.abs(_53_v1))).IsSubsetOf(_118_v48), _dafny.Seq.IsPrefixOf(_51_v0, _51_v0))).length);
          let _140_v64;
          _140_v64 = _dafny.Map.Empty.slice().updateUnsafe(_54_v2,_module.__default.fm4(_54_v2, _52_globalState));
          let _141_v65;
          _141_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_54_v2, _138_v62, _52_globalState),_54_v2);
          _54_v2 = (((_140_v64).contains(((_139_v63)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_139_v63).length))]).isLessThanOrEqualTo(new BigNumber((_141_v65).length)))) ? ((_140_v64).get(((_139_v63)[_module.__default.safeIndex(new BigNumber(57), new BigNumber((_139_v63).length))]).isLessThanOrEqualTo(new BigNumber((_141_v65).length)))) : (!(_54_v2)));
          let _142_v66;
          let _nw29 = new _module.C0();
          _nw29.__ctor();
          _142_v66 = _nw29;
        } else {
          let _143_v67;
          let _nw30 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _143_v67 = _nw30;
          let _144_v68;
          let _init5 = ((_145_v1) => function (_146_i7) {
            return (_146_i7).minus(_145_v1);
          })(_53_v1);
          let _nw31 = Array((new BigNumber(8)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw31.length); _i0_5++) {
            _nw31[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _144_v68 = _nw31;
          let _index18 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length));
          (_144_v68)[_index18] = _53_v1;
          let _147_v69;
          _147_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_51_v0).length),new BigNumber((_81_v22).length));
          let _148_v70;
          _148_v70 = _dafny.Map.Empty.slice().updateUnsafe(_137_v61,_147_v69);
          let _149_v71;
          _149_v71 = _dafny.Map.Empty.slice().updateUnsafe(_144_v68,_143_v67);
          let _index19 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length));
          let _rhs13 = !(_148_v70).equals(_148_v70);
          let _rhs14 = _54_v2;
          let _rhs15 = (((_149_v71).contains(_144_v68)) ? ((_149_v71).get(_144_v68)) : (_143_v67));
          let _rhs16 = (_dafny.ZERO).minus(_138_v62);
          let _lhs10 = _144_v68;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length));
          _54_v2 = _rhs13;
          _54_v2 = _rhs14;
          _143_v67 = _rhs15;
          _lhs10[_lhs11] = _rhs16;
          let _150_v72;
          _150_v72 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.update(_131_v56, _module.__default.safeIndex(_53_v1, new BigNumber((_131_v56).length)), _54_v2), _131_v56), _dafny.Seq.Concat(_131_v56, _131_v56));
          _150_v72 = _dafny.Set.fromElements(_131_v56, _dafny.Seq.Concat(_module.__default.fm5((_144_v68)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length))], _54_v2, _54_v2, _52_globalState), _dafny.Seq.of(_54_v2)), _131_v56);
          let _151_v73;
          _151_v73 = new _dafny.CodePoint('a'.codePointAt(0));
          let _152_v74;
          _152_v74 = _dafny.Map.Empty.slice().updateUnsafe(_118_v48,_151_v73);
          _152_v74 = (_152_v74).update(_118_v48, ((false) ? (_151_v73) : (_151_v73)));
          let _153_v75;
          let _nw32 = Array((new BigNumber(24)).toNumber()).fill(_module.D0.Default());
          _153_v75 = _nw32;
          let _index20 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_153_v75).length));
          (_153_v75)[_index20] = _83_v24;
          let _pat_let_tv0 = _82_v23;
          let _index21 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_153_v75).length));
          (_153_v75)[_index21] = function (_pat_let0_0) {
            return function (_154_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_155_dt__update_hcf3_h0) {
                  return _module.D0.create_DC2((_154_dt__update__tmp_h0).dtor_cf2, _155_dt__update_hcf3_h0, (_154_dt__update__tmp_h0).dtor_cf4);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_83_v24);
          let _156_v76;
          let _init6 = function (_157_i8) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          };
          let _nw33 = Array((new BigNumber(24)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw33.length); _i0_6++) {
            _nw33[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _156_v76 = _nw33;
          let _index22 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_156_v76).length));
          (_156_v76)[_index22] = _151_v73;
          let _index23 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_156_v76).length));
          let _rhs17 = _151_v73;
          let _rhs18 = (_118_v48).update((_144_v68)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length))], _module.__default.abs((_144_v68)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length))]));
          let _rhs19 = (_144_v68)[_module.__default.safeIndex(new BigNumber(580), new BigNumber((_144_v68).length))];
          let _lhs12 = _156_v76;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_156_v76).length));
          _lhs12[_lhs13] = _rhs17;
          _118_v48 = _rhs18;
          _138_v62 = _rhs19;
        }
        let _158_v77;
        _158_v77 = _dafny.MultiSet.fromElements(!(_54_v2));
        _158_v77 = _158_v77;
      }
      _115_v45 = _115_v45;
      let _159_v78;
      _159_v78 = new _dafny.CodePoint('i'.codePointAt(0));
      let _rhs20 = _54_v2;
      let _rhs21 = _54_v2;
      let _rhs22 = _159_v78;
      let _rhs23 = new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(_53_v1, _53_v1))).Intersect(((_54_v2) ? (_118_v48) : (_118_v48)))).cardinality());
      _54_v2 = _rhs20;
      _54_v2 = _rhs21;
      _159_v78 = _rhs22;
      _53_v1 = _rhs23;
      let _160_v79;
      _160_v79 = _module.D0.create_DC0();
      let _161_v80;
      let _nw34 = Array((_dafny.ONE).toNumber());
      _nw34[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_53_v1,_160_v79)).length);
      _161_v80 = _nw34;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_161_v80).length))) {
        let _162_i9 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_162_i9)) && ((_162_i9).isLessThan(new BigNumber((_161_v80).length))))) {
          (_161_v80)[(_162_i9)] = _module.__default.safeModuloInt(_162_i9, _53_v1);
        }
      }
      let _163_v81;
      let _nw35 = new _module.C0();
      _nw35.__ctor();
      _163_v81 = _nw35;
      _54_v2 = _54_v2;
      let _164_v82;
      _164_v82 = _dafny.Map.Empty.slice().updateUnsafe(_54_v2,_53_v1);
      if (((_54_v2) ? (!(_164_v82).equals(_164_v82)) : (!(_54_v2)))) {
        let _165_v83;
        let _166_v84;
        let _out12;
        let _out13;
        let _outcollector6 = _module.__default.m0(_53_v1, _53_v1, _52_globalState);
        _out12 = _outcollector6[0];
        _out13 = _outcollector6[1];
        _165_v83 = _out12;
        _166_v84 = _out13;
        _54_v2 = !(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_167_v78) => function (_168_i10) {
          return _167_v78;
        })(_159_v78)), _51_v0));
        let _169_v85;
        let _170_v86;
        let _out14;
        let _out15;
        let _outcollector7 = _module.__default.m0(_module.__default.safeDivisionInt(new BigNumber((_118_v48).cardinality()), _53_v1), _53_v1, _52_globalState);
        _out14 = _outcollector7[0];
        _out15 = _outcollector7[1];
        _169_v85 = _out14;
        _170_v86 = _out15;
        _54_v2 = !_dafny.Seq.contains(_dafny.Seq.Concat(_51_v0, _51_v0), _159_v78);
        let _171_v87;
        let _nw36 = new _module.C0();
        _nw36.__ctor();
        _171_v87 = _nw36;
      } else {
        _53_v1 = _53_v1;
        let _rhs24 = (_dafny.ZERO).minus(_53_v1);
        let _rhs25 = _53_v1;
        let _rhs26 = _164_v82;
        let _lhs14 = _52_globalState;
        _lhs14.f2 = _rhs24;
        _53_v1 = _rhs25;
        _164_v82 = _rhs26;
        let _172_v88;
        let _nw37 = Array((new BigNumber(12)).toNumber()).fill([]);
        _172_v88 = _nw37;
        let _173_v89;
        let _init7 = ((_174_v78) => function (_175_i11) {
          return _174_v78;
        })(_159_v78);
        let _nw38 = Array((new BigNumber(7)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw38.length); _i0_7++) {
          _nw38[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _173_v89 = _nw38;
        let _index24 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_172_v88).length));
        (_172_v88)[_index24] = _173_v89;
        let _index25 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_172_v88).length));
        (_172_v88)[_index25] = _173_v89;
        let _176_v90;
        let _177_v91;
        let _out16;
        let _out17;
        let _outcollector8 = _module.__default.m0(new BigNumber((_51_v0).length), (_53_v1).plus(_53_v1), _52_globalState);
        _out16 = _outcollector8[0];
        _out17 = _outcollector8[1];
        _176_v90 = _out16;
        _177_v91 = _out17;
        _51_v0 = _51_v0;
      }
      let _178_v92;
      let _nw39 = Array((new BigNumber(9)).toNumber()).fill(false);
      _178_v92 = _nw39;
      _178_v92 = _178_v92;
      let _179_v93;
      _179_v93 = _dafny.Seq.of(_161_v80);
      let _180_v94;
      _180_v94 = _dafny.Map.Empty.slice().updateUnsafe(_117_v47,_dafny.Seq.Concat(_179_v93, _179_v93));
      _180_v94 = (_180_v94).update(_117_v47, _179_v93);
      process.stdout.write((_51_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_52_globalState).f0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_52_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_52_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_52_globalState).f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_53_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_54_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_81_v22).equals(_dafny.Set.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_83_v24).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_83_v24).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_116_v46).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_117_v47, _dafny.Seq.of(new BigNumber(386), new BigNumber(386), new BigNumber(386), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v48).equals(_dafny.MultiSet.fromElements(new BigNumber(-254), new BigNumber(386), new BigNumber(386)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_131_v56, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_v78));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v80)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_v82).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_179_v93).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_180_v94).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC1(cf0, cf1) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0) && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf2 === other.cf2 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0();
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
    static create_DC3(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + this.cf5.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.ZERO;
      this._f0 = _dafny.Seq.UnicodeFromString("");
      this._f1 = false;
      this._f3 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm2(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ieh"), _dafny.Seq.UnicodeFromString("ewuihk")), _dafny.Seq.UnicodeFromString("h"));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
