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
    static fm0(p0, p1, p2, globalState) {
      return new BigNumber(-184);
    };
    static fm3(p0, globalState) {
      let _source0 = _module.D2.create_DC7(false, false, true);
      if (_source0.is_DC7) {
        let _0___mcc_h0 = (_source0).cf11;
        let _1___mcc_h1 = (_source0).cf12;
        let _2___mcc_h2 = (_source0).cf13;
        let _3_cf13 = _2___mcc_h2;
        let _4_cf12 = _1___mcc_h1;
        let _5_cf11 = _0___mcc_h0;
        return _module.D1.create_DC4(_dafny.Seq.UnicodeFromString("ljgfacm"));
      } else {
        let _6___mcc_h3 = (_source0).cf10;
        let _7_cf10 = _6___mcc_h3;
        if (!(false)) {
          return _module.D1.create_DC4(_dafny.Seq.UnicodeFromString("g"));
        } else {
          return _module.D1.create_DC4(_dafny.Seq.UnicodeFromString("dewv"));
        }
      }
    };
    static fm4(p0, p1, p2, globalState) {
      return true;
    };
    static fm6(p0, p1, globalState) {
      return _dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-430))).length)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-548), new BigNumber(187))) {
          let _8_v0 = _compr_0;
          if (((new BigNumber(-548)).isLessThanOrEqualTo(_8_v0)) && ((_8_v0).isLessThan(new BigNumber(187)))) {
            _coll0.push([(_8_v0).multipliedBy(new BigNumber(953)),new BigNumber(-372)]);
          }
        }
        return _coll0;
      }())).length))).length)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)));
    };
    static fm7(globalState) {
      return _module.D0.create_DC2(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("sfekhaf")).length), new BigNumber(774), new BigNumber(-567), new BigNumber(402)), new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), (_dafny.ZERO).minus(new BigNumber(-165)))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)))).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-3)), function (_9_i0) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length), new BigNumber(740)))).cardinality()), false);
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(853)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ks")).length))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(121), new BigNumber(889), new BigNumber((_dafny.Seq.UnicodeFromString("xikkpo")).length), new BigNumber((_dafny.Seq.of(new BigNumber(258))).length), new BigNumber((_dafny.MultiSet.fromElements(!(!(true)), true, true, !(true))).cardinality())), _dafny.Seq.of(new BigNumber(379))));
    };
    static fm9(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(832), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_10_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length), (_dafny.ZERO).minus(new BigNumber(-61))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(510)), function (_11_i1) {
        return new BigNumber(258);
      })));
    };
    static fm10(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, !(true)), _dafny.Seq.of((_module.D0.create_DC1(_dafny.Set.fromElements(_dafny.Set.fromElements(true, !(true))), true)).dtor_cf3)), _dafny.Seq.of((_module.D0.create_DC2(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ajxt")).length)), new BigNumber(-490)), new BigNumber(836), true)).dtor_cf6));
    };
    static fm11(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0))))).Intersect(((_module.D7.create_DC18(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))))).dtor_cf26).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))));
    };
    static fm12(p0, p1, globalState) {
      if (!(false)) {
        return _module.D4.create_DC13(_module.D4.create_DC12(new BigNumber(60), new _dafny.CodePoint('e'.codePointAt(0)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(115)), function (_12_i0) {
  return new _dafny.CodePoint('t'.codePointAt(0));
})).length))));
      } else {
        return _module.D4.create_DC13(_module.D4.create_DC12(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-739),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(380))).length), new BigNumber((_dafny.Seq.UnicodeFromString("bntttie")).length))).length))).length), new _dafny.CodePoint('g'.codePointAt(0)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gix")).length))).cardinality()))));
      }
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("kem")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("cqlomrr"), _dafny.Seq.UnicodeFromString("hxb"))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("dugneonhq")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("awqmcmgf"))));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), ((false) ? (new _dafny.CodePoint('d'.codePointAt(0))) : (new _dafny.CodePoint('g'.codePointAt(0)))), new _dafny.CodePoint('j'.codePointAt(0)));
    };
    static fm15(p0, p1, globalState) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    };
    static fm16(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(false), false)).length))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)))).length))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(951))).length), new BigNumber(849)))));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat((_module.D8.create_DC20(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("aapfmbeu")).length))), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("lytyvmod")).length))))).dtor_cf28, _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("rdnvcfxx")).length))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),false)).Keys.Elements) {
          let _13_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(897),false)).contains(_13_v0)) {
            _coll1.push([(_13_v0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_14_i0) {
              return new BigNumber(303);
            })).length)),new _dafny.CodePoint('t'.codePointAt(0))]);
          }
        }
        return _coll1;
      }()).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gwpfdjks")).length))).length)), _dafny.Seq.of(new BigNumber(764)), (_module.D9.create_DC24(_dafny.Seq.of(new BigNumber(369), new BigNumber(-732)))).dtor_cf36, _dafny.Seq.of(new BigNumber(961), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(812), new BigNumber(50)))));
    };
    static fm18(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dh"), _dafny.Seq.UnicodeFromString("nrki"));
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let r1 = [];
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      r3 = p0;
      r2 = p0;
      let _15_v0;
      _15_v0 = true;
      let _16_i0;
      _16_i0 = _dafny.ZERO;
      L0: {
        while (_15_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_16_i0)) {
              break L0;
            }
            _16_i0 = (_16_i0).plus(_dafny.ONE);
            if (_15_v0) {
              r3 = (p0).minus((p0).multipliedBy(p0));
              let _17_v1;
              _17_v1 = _dafny.Seq.UnicodeFromString("ym");
              let _18_v2;
              _18_v2 = _dafny.Set.fromElements(new BigNumber((_17_v1).length));
              let _19_v3;
              _19_v3 = _dafny.Map.Empty.slice().updateUnsafe((_15_v0) === (_15_v0),_18_v2);
              _19_v3 = (_19_v3).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_15_v0),function () {
                let _coll2 = new _dafny.Set();
                for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-838), new BigNumber(345))) {
                  let _20_v4 = _compr_2;
                  if (((new BigNumber(-838)).isLessThanOrEqualTo(_20_v4)) && ((_20_v4).isLessThan(new BigNumber(345)))) {
                    _coll2.add((_20_v4).plus(p0));
                  }
                }
                return _coll2;
              }()));
              let _21_v5;
              let _nw0 = Array((new BigNumber(5)).toNumber()).fill(_module.D4.Default());
              _21_v5 = _nw0;
              let _22_v6;
              _22_v6 = new _dafny.CodePoint('t'.codePointAt(0));
              let _23_v7;
              _23_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(10),new BigNumber(-844));
              let _24_v8;
              _24_v8 = _dafny.MultiSet.fromElements(p0);
              let _25_v9;
              _25_v9 = _module.D4.create_DC12(p0, _module.__default.fm15(_22_v6, _23_v7, globalState), _24_v8);
              let _index0 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_21_v5).length));
              (_21_v5)[_index0] = _25_v9;
              let _index1 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_21_v5).length));
              (_21_v5)[_index1] = _25_v9;
              let _26_v10;
              _26_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(768),_17_v1);
              let _27_v11;
              _27_v11 = _module.D0.create_DC3(_22_v6);
              _26_v10 = (_26_v10).update(p0, _module.__default.fm18((_27_v11).dtor_cf7, globalState));
              let _28_v12;
              let _nw1 = new _module.C1();
              _nw1.__ctor(((_15_v0) ? (_15_v0) : (_15_v0)), new BigNumber(-333));
              _28_v12 = _nw1;
              let _nw2 = new _module.C1();
              _nw2.__ctor((p0).isLessThanOrEqualTo(new BigNumber(667)), p0);
              _28_v12 = _nw2;
            } else {
              let _29_v13;
              _29_v13 = new _dafny.CodePoint('u'.codePointAt(0));
              (globalState).f4 = _29_v13;
              let _30_v14;
              let _nw3 = Array((new BigNumber(7)).toNumber()).fill([]);
              _30_v14 = _nw3;
              let _31_v15;
              _31_v15 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_30_v14);
              _31_v15 = (_31_v15).update((_15_v0) || (_15_v0), _30_v14);
              r2 = p0;
              let _32_v16;
              let _nw4 = Array((new BigNumber(29)).toNumber()).fill(false);
              _32_v16 = _nw4;
              _32_v16 = _32_v16;
              let _33_v17;
              _33_v17 = _dafny.Seq.of(p0);
              r0 = _dafny.Seq.contains(_33_v17, p0);
            }
            let _34_v18;
            let _nw5 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
            _34_v18 = _nw5;
            let _35_v19;
            _35_v19 = _module.D0.create_DC0(p0, _15_v0);
            let _36_v20;
            _36_v20 = _dafny.MultiSet.fromElements((_35_v19).dtor_cf0, p0);
            let _37_v21;
            _37_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_36_v20).cardinality()),_34_v18);
            let _38_v22;
            _38_v22 = _dafny.Seq.of(_34_v18);
            let _39_v23;
            _39_v23 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_34_v18);
            let _40_v24;
            let _nw6 = Array((new BigNumber(24)).toNumber());
            _nw6[(_dafny.ZERO).toNumber()] = _34_v18;
            _nw6[(_dafny.ONE).toNumber()] = _34_v18;
            _nw6[(new BigNumber(2)).toNumber()] = (((_37_v21).contains(p0)) ? ((_37_v21).get(p0)) : (_34_v18));
            _nw6[(new BigNumber(3)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(4)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(5)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(6)).toNumber()] = (_38_v22)[_module.__default.safeIndex(p0, new BigNumber((_38_v22).length))];
            _nw6[(new BigNumber(7)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(8)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(9)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(10)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(11)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(12)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(13)).toNumber()] = (((_39_v23).contains(_15_v0)) ? ((_39_v23).get(_15_v0)) : (_34_v18));
            _nw6[(new BigNumber(14)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(15)).toNumber()] = (((_37_v21).contains(p0)) ? ((_37_v21).get(p0)) : (_34_v18));
            _nw6[(new BigNumber(16)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(17)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(18)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(19)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(20)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(21)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(22)).toNumber()] = _34_v18;
            _nw6[(new BigNumber(23)).toNumber()] = _34_v18;
            _40_v24 = _nw6;
            let _41_v25;
            _41_v25 = _module.D5.create_DC14(_40_v24);
            _41_v25 = _41_v25;
            if (false) {
              let _42_v26;
              _42_v26 = _dafny.Seq.UnicodeFromString("esfxlsqkn");
              r3 = _module.__default.safeModuloInt(new BigNumber((_42_v26).length), p0);
              let _43_v27;
              let _init0 = ((_44_v0) => function (_45_i1) {
                return _44_v0;
              })(_15_v0);
              let _nw7 = Array((new BigNumber(12)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
                _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _43_v27 = _nw7;
              let _index2 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_43_v27).length));
              (_43_v27)[_index2] = (_15_v0) && (_module.__default.fm6((p1)[_module.__default.safeIndex(p0, new BigNumber((p1).length))], p0, globalState));
              let _index3 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_43_v27).length));
              (_43_v27)[_index3] = (false) || (false);
              r3 = p0;
              let _46_v28;
              _46_v28 = new _dafny.CodePoint('b'.codePointAt(0));
              let _47_v29;
              _47_v29 = _dafny.Map.Empty.slice().updateUnsafe(true,_46_v28);
              let _48_v30;
              let _nw8 = new _module.C0();
              _nw8.__ctor((_47_v29).equals(_47_v29));
              _48_v30 = _nw8;
              let _49_v31;
              _49_v31 = _module.D0.create_DC3(_46_v28);
              let _50_v32;
              let _nw9 = Array((new BigNumber(16)).toNumber());
              _nw9[(_dafny.ZERO).toNumber()] = _module.D0.create_DC3(_46_v28);
              _nw9[(_dafny.ONE).toNumber()] = _module.D0.create_DC3(_46_v28);
              _nw9[(new BigNumber(2)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(3)).toNumber()] = _module.D0.create_DC3(_46_v28);
              _nw9[(new BigNumber(4)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(5)).toNumber()] = _module.D0.create_DC3(_46_v28);
              _nw9[(new BigNumber(6)).toNumber()] = _module.D0.create_DC3(new _dafny.CodePoint('c'.codePointAt(0)));
              _nw9[(new BigNumber(7)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(8)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(9)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(10)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(11)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(12)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(13)).toNumber()] = _module.D0.create_DC3(_46_v28);
              _nw9[(new BigNumber(14)).toNumber()] = _49_v31;
              _nw9[(new BigNumber(15)).toNumber()] = _49_v31;
              _50_v32 = _nw9;
              let _51_v33;
              let _nw10 = Array((new BigNumber(9)).toNumber());
              _nw10[(_dafny.ZERO).toNumber()] = _50_v32;
              _nw10[(_dafny.ONE).toNumber()] = _50_v32;
              _nw10[(new BigNumber(2)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(3)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(4)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(5)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(6)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(7)).toNumber()] = _50_v32;
              _nw10[(new BigNumber(8)).toNumber()] = _50_v32;
              _51_v33 = _nw10;
              let _index4 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_51_v33).length));
              (_51_v33)[_index4] = _50_v32;
              let _index5 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_51_v33).length));
              let _rhs0 = _48_v30;
              let _rhs1 = _50_v32;
              let _lhs0 = _51_v33;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_51_v33).length));
              _48_v30 = _rhs0;
              _lhs0[_lhs1] = _rhs1;
              let _index6 = _module.__default.safeIndex(new BigNumber(92), new BigNumber((_43_v27).length));
              (_43_v27)[_index6] = !((_43_v27)[_module.__default.safeIndex(new BigNumber(92), new BigNumber((_43_v27).length))]);
            } else {
              r0 = _15_v0;
              let _index7 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length));
              (_34_v18)[_index7] = p0;
              let _index8 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length));
              (_34_v18)[_index8] = p0;
              let _52_v34;
              _52_v34 = _dafny.Seq.of(new BigNumber(479), _module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe(true,_15_v0), _15_v0, _15_v0, globalState));
              let _53_v35;
              _53_v35 = _dafny.Set.fromElements(_15_v0, _15_v0);
              let _54_v36;
              _54_v36 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_module.__default.fm4(_15_v0, _53_v35, p0, globalState));
              let _55_v37;
              _55_v37 = _dafny.Seq.of((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))], (new BigNumber((_52_v34).length)).minus(_module.__default.fm0(_54_v36, _15_v0, _15_v0, globalState)), _module.__default.safeModuloInt((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))], p0));
              _55_v37 = _dafny.Seq.Concat(_55_v37, _52_v34);
              let _56_v38;
              _56_v38 = new _dafny.CodePoint('g'.codePointAt(0));
              let _57_v39;
              _57_v39 = _dafny.Seq.UnicodeFromString("sir");
              let _58_v40;
              _58_v40 = _dafny.Set.fromElements(_36_v20, _36_v20);
              let _59_v41;
              let _nw11 = Array((new BigNumber(17)).toNumber());
              _nw11[(_dafny.ZERO).toNumber()] = _15_v0;
              _nw11[(_dafny.ONE).toNumber()] = _15_v0;
              _nw11[(new BigNumber(2)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(3)).toNumber()] = !(!(!_dafny.Seq.contains(_57_v39, _56_v38)));
              _nw11[(new BigNumber(4)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(5)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(6)).toNumber()] = (p0).isLessThan(_module.__default.fm0(_54_v36, _15_v0, _15_v0, globalState));
              _nw11[(new BigNumber(7)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(8)).toNumber()] = (p1)[_module.__default.safeIndex(p0, new BigNumber((p1).length))];
              _nw11[(new BigNumber(9)).toNumber()] = ((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))]).isLessThanOrEqualTo((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))]);
              _nw11[(new BigNumber(10)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(11)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(12)).toNumber()] = (p0).isLessThanOrEqualTo(new BigNumber(77));
              _nw11[(new BigNumber(13)).toNumber()] = (_58_v40).contains(_36_v20);
              _nw11[(new BigNumber(14)).toNumber()] = _15_v0;
              _nw11[(new BigNumber(15)).toNumber()] = (_15_v0) || (!((p1)[_module.__default.safeIndex(p0, new BigNumber((p1).length))]));
              _nw11[(new BigNumber(16)).toNumber()] = !(_15_v0);
              _59_v41 = _nw11;
              let _index9 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length));
              let _rhs2 = p0;
              let _rhs3 = _module.__default.fm0(_54_v36, !(!(_15_v0)), _15_v0, globalState);
              let _rhs4 = _59_v41;
              let _rhs5 = ((((_15_v0) ? (_15_v0) : ((p1)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((p1).length))]))) ? (_module.__default.fm6(_15_v0, (_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))], globalState)) : (((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))]).isLessThan((_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))])));
              let _lhs2 = _34_v18;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length));
              _lhs2[_lhs3] = _rhs2;
              r3 = _rhs3;
              _59_v41 = _rhs4;
              _15_v0 = _rhs5;
              let _index10 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length));
              (_34_v18)[_index10] = (_34_v18)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_34_v18).length))];
            }
            let _60_v42;
            _60_v42 = new _dafny.CodePoint('x'.codePointAt(0));
            let _61_v43;
            _61_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
            let _62_v44;
            _62_v44 = _dafny.Seq.of(_module.__default.fm15(_60_v42, _61_v43, globalState));
            let _63_v45;
            _63_v45 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_62_v44);
            let _64_v46;
            _64_v46 = _dafny.Set.fromElements(!(_63_v45).contains(false), (p0).isLessThan(p0), _15_v0);
            _64_v46 = ((_64_v46).Union(_64_v46)).Union(_64_v46);
          }
        }
      }
      let _65_v47;
      let _nw12 = Array((new BigNumber(3)).toNumber());
      _65_v47 = _nw12;
      _65_v47 = _65_v47;
      let _66_v48;
      _66_v48 = _dafny.Seq.UnicodeFromString("mns");
      let _67_v49;
      _67_v49 = _module.D1.create_DC4(_66_v48);
      let _source1 = _67_v49;
      if (_source1.is_DC5) {
        let _68___mcc_h0 = (_source1).cf9;
        let _69_cf9 = _68___mcc_h0;
        let _70_v50;
        _70_v50 = _dafny.Set.fromElements(_69_cf9);
        let _71_v51;
        _71_v51 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_69_cf9);
        let _72_v52;
        _72_v52 = _dafny.Set.fromElements(_module.__default.fm4(_15_v0, _70_v50, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_71_v51).length), p0, p0, p0, p0)).cardinality()), globalState));
        if (((_15_v0) ? ((_72_v52).IsSubsetOf(_70_v50)) : (_15_v0))) {
          let _73_v53;
          _73_v53 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_15_v0, _69_cf9), p1, _dafny.Seq.of(_15_v0, _69_cf9), _dafny.Seq.of(_69_cf9, _69_cf9, _15_v0, _69_cf9, !(_69_cf9)));
          let _74_v54;
          _74_v54 = _dafny.MultiSet.fromElements(_73_v53, _dafny.MultiSet.fromElements(p1));
          let _75_v55;
          _75_v55 = _dafny.MultiSet.fromElements(true);
          let _76_v56;
          _76_v56 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0), (((_74_v54).contains(_73_v53)) ? ((_74_v54).get(_73_v53)) : ((((_75_v55).contains(_15_v0)) ? ((_75_v55).get(_15_v0)) : (p0)))), p0);
          _76_v56 = _76_v56;
          let _77_v57;
          _77_v57 = new _dafny.CodePoint('c'.codePointAt(0));
          let _78_v58;
          _78_v58 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_66_v48, _module.__default.safeIndex(p0, new BigNumber((_66_v48).length)), _77_v57),p0);
          let _79_v59;
          _79_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_module.__default.fm18(_77_v57, globalState), _66_v48),new BigNumber((_78_v58).length));
          let _80_v60;
          _80_v60 = _dafny.Set.fromElements(p0);
          let _81_v61;
          _81_v61 = _dafny.Set.fromElements(p0, new BigNumber((_dafny.MultiSet.fromElements(_80_v60, _dafny.Set.fromElements(p0))).cardinality()));
          _79_v59 = (_79_v59).update(((_69_cf9) ? (_66_v48) : (_dafny.Seq.UnicodeFromString("bmdwy"))), new BigNumber((_80_v60).length));
          let _82_v62;
          let _nw13 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _82_v62 = _nw13;
          let _index11 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_82_v62).length));
          (_82_v62)[_index11] = p0;
          let _index12 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_82_v62).length));
          (_82_v62)[_index12] = (p0).minus(_module.__default.safeDivisionInt(new BigNumber((_66_v48).length), (_dafny.ZERO).minus(new BigNumber((_75_v55).cardinality()))));
          r2 = p0;
          let _83_v63;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_15_v0);
          _83_v63 = _nw14;
          let _84_v64;
          _84_v64 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_83_v63);
          _84_v64 = (_84_v64).update((_82_v62)[_module.__default.safeIndex(new BigNumber(948), new BigNumber((_82_v62).length))], _83_v63);
        } else {
          let _85_v65;
          let _init1 = function (_86_i2) {
            return (_86_i2).plus(new BigNumber(310));
          };
          let _nw15 = Array((new BigNumber(5)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw15.length); _i0_1++) {
            _nw15[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _85_v65 = _nw15;
          let _index13 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_85_v65).length));
          (_85_v65)[_index13] = p0;
          let _87_v66;
          _87_v66 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_69_cf9);
          let _88_v67;
          _88_v67 = _dafny.Map.Empty.slice().updateUnsafe(_66_v48,_87_v66);
          let _index14 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_85_v65).length));
          let _rhs6 = _module.__default.fm6(_15_v0, (_dafny.ZERO).minus(_module.__default.fm0((((_88_v67).contains(_66_v48)) ? ((_88_v67).get(_66_v48)) : (_87_v66)), (p1)[_module.__default.safeIndex(new BigNumber((_66_v48).length), new BigNumber((p1).length))], _15_v0, globalState)), globalState);
          let _rhs7 = (_dafny.ZERO).minus(p0);
          let _rhs8 = !(_69_cf9);
          let _rhs9 = ((p0).plus(_module.__default.fm0(_dafny.Map.Empty.slice().updateUnsafe(!(_15_v0),!(true)), (((_87_v66).contains(_15_v0)) ? ((_87_v66).get(_15_v0)) : (_69_cf9)), _15_v0, globalState))).minus(p0);
          let _lhs4 = _85_v65;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_85_v65).length));
          r0 = _rhs6;
          _lhs4[_lhs5] = _rhs7;
          r0 = _rhs8;
          r2 = _rhs9;
          let _89_v68;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_15_v0);
          _89_v68 = _nw16;
          let _rhs10 = _module.__default.safeModuloInt(p0, new BigNumber(-220));
          let _rhs11 = _15_v0;
          let _rhs12 = _89_v68;
          r2 = _rhs10;
          _15_v0 = _rhs11;
          _89_v68 = _rhs12;
          let _90_v69;
          let _init2 = ((_91_v0) => function (_92_i3) {
            return _91_v0;
          })(_15_v0);
          let _nw17 = Array((new BigNumber(7)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw17.length); _i0_2++) {
            _nw17[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _90_v69 = _nw17;
          let _index15 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_90_v69).length));
          (_90_v69)[_index15] = _69_cf9;
          let _index16 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_90_v69).length));
          let _rhs13 = _69_cf9;
          let _rhs14 = _15_v0;
          let _lhs6 = _90_v69;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_90_v69).length));
          _69_cf9 = _rhs13;
          _lhs6[_lhs7] = _rhs14;
          r3 = (p0).minus(_module.__default.safeDivisionInt(p0, p0));
          let _index17 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_90_v69).length));
          (_90_v69)[_index17] = !(_69_cf9);
        }
        let _93_v70;
        let _nw18 = Array((new BigNumber(26)).toNumber()).fill(false);
        _93_v70 = _nw18;
        let _index18 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_93_v70).length));
        (_93_v70)[_index18] = _69_cf9;
        let _index19 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_93_v70).length));
        (_93_v70)[_index19] = (p1)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber(609), new BigNumber((_66_v48).length)), new BigNumber((p1).length))];
        let _94_v71;
        let _nw19 = new _module.C1();
        _nw19.__ctor((_93_v70)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_93_v70).length))], p0);
        _94_v71 = _nw19;
        let _95_v72;
        _95_v72 = _dafny.Map.Empty.slice().updateUnsafe(_94_v71,p0);
        let _96_v73;
        _96_v73 = _dafny.Map.Empty.slice().updateUnsafe(_69_cf9,new BigNumber((_95_v72).length));
        let _97_v74;
        _97_v74 = _dafny.Seq.of((_94_v71).f15);
        let _98_v75;
        let _nw20 = new _module.C1();
        _nw20.__ctor(_module.__default.fm6(_15_v0, new BigNumber((_96_v73).length), globalState), new BigNumber((_97_v74).length));
        _98_v75 = _nw20;
        let _99_v76;
        _99_v76 = _dafny.Set.fromElements(_93_v70);
        let _100_v77;
        _100_v77 = _dafny.Map.Empty.slice().updateUnsafe(_99_v76,_96_v73);
        _100_v77 = (_100_v77).update(_99_v76, _96_v73);
      } else {
        let _101___mcc_h1 = (_source1).cf8;
        let _102_cf8 = _101___mcc_h1;
        let _103_v78;
        _103_v78 = _dafny.Seq.of(p0);
        let _104_v79;
        _104_v79 = _dafny.Map.Empty.slice().updateUnsafe(p0,_103_v78);
        let _105_v80;
        _105_v80 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_15_v0);
        _103_v78 = _dafny.Seq.Concat(_dafny.Seq.update((((_104_v79).contains(p0)) ? ((_104_v79).get(p0)) : (_103_v78)), _module.__default.safeIndex(_module.__default.fm0(_105_v80, _15_v0, _15_v0, globalState), new BigNumber(((((_104_v79).contains(p0)) ? ((_104_v79).get(p0)) : (_103_v78))).length)), p0), _dafny.Seq.Concat(_103_v78, _103_v78));
        let _106_v81;
        let _nw21 = new _module.C1();
        _nw21.__ctor((new BigNumber((_dafny.Seq.UnicodeFromString("sm")).length)).isLessThan(p0), p0);
        _106_v81 = _nw21;
        let _107_v82;
        _107_v82 = _dafny.Map.Empty.slice().updateUnsafe((_106_v81).f15,_15_v0);
        let _108_v83;
        _108_v83 = _dafny.Map.Empty.slice().updateUnsafe((_106_v81).f15,_107_v82);
        let _109_v84;
        _109_v84 = _dafny.Map.Empty.slice().updateUnsafe(((_106_v81).f15).minus(new BigNumber(((((_108_v83).contains(p0)) ? ((_108_v83).get(p0)) : (_107_v82))).length)),(_106_v81).f14);
        _107_v82 = (_107_v82).update(p0, _15_v0);
        let _110_v85;
        _110_v85 = new _dafny.CodePoint('y'.codePointAt(0));
        let _111_v86;
        let _nw22 = new _module.C0();
        _nw22.__ctor((_106_v81).f14);
        _111_v86 = _nw22;
        let _112_v87;
        _112_v87 = _dafny.Map.Empty.slice().updateUnsafe(_15_v0,_110_v85);
        let _113_v88;
        _113_v88 = _dafny.Map.Empty.slice().updateUnsafe(_111_v86,_112_v87);
        let _114_v89;
        _114_v89 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(((((_113_v88).contains(_111_v86)) ? ((_113_v88).get(_111_v86)) : (_112_v87))).length)));
        let _115_v90;
        _115_v90 = _module.D4.create_DC12((_106_v81).f15, _110_v85, _114_v89);
        r3 = (new BigNumber(((_115_v90).dtor_cf20).cardinality())).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((_103_v78)[_module.__default.safeIndex((_106_v81).f15, new BigNumber((_103_v78).length))])));
      }
      r0 = false;
      let _116_v91;
      _116_v91 = new _dafny.CodePoint('l'.codePointAt(0));
      let _117_v92;
      _117_v92 = _dafny.MultiSet.fromElements(new BigNumber((p1).length));
      let _118_v93;
      _118_v93 = _module.D4.create_DC12(p0, _116_v91, _117_v92);
      let _pat_let_tv0 = _15_v0;
      let _pat_let_tv1 = p1;
      let _pat_let_tv2 = p1;
      let _pat_let_tv3 = _15_v0;
      let _pat_let_tv4 = _15_v0;
      r0 = !(function (_source2) {
        if (_source2.is_DC12) {
          let _119___mcc_h2 = (_source2).cf18;
          let _120___mcc_h3 = (_source2).cf19;
          let _121___mcc_h4 = (_source2).cf20;
          let _122_cf20 = _121___mcc_h4;
          let _123_cf19 = _120___mcc_h3;
          let _124_cf18 = _119___mcc_h2;
          return !(_pat_let_tv0) || (!((_pat_let_tv1)[_module.__default.safeIndex(_124_cf18, new BigNumber((_pat_let_tv2).length))]));
        } else if (_source2.is_DC11) {
          let _125___mcc_h5 = (_source2).cf17;
          let _126_cf17 = _125___mcc_h5;
          return _pat_let_tv3;
        } else {
          let _127___mcc_h6 = (_source2).cf21;
          let _128_cf21 = _127___mcc_h6;
          return _pat_let_tv4;
        }
      }(_118_v93));
      let _129_v94;
      _129_v94 = _dafny.Seq.of(p0, p0, p0, p0, p0);
      let _130_v95;
      _130_v95 = _module.D0.create_DC0(p0, _15_v0);
      let _pat_let_tv5 = p0;
      let _131_v96;
      let _nw23 = Array((new BigNumber(14)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = _130_v95;
      _nw23[(_dafny.ONE).toNumber()] = _module.D0.create_DC0(p0, _15_v0);
      _nw23[(new BigNumber(2)).toNumber()] = _module.D0.create_DC0(p0, _module.__default.fm6(_15_v0, p0, globalState));
      _nw23[(new BigNumber(3)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(4)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(5)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(6)).toNumber()] = _module.D0.create_DC0(p0, _module.__default.fm6(_15_v0, p0, globalState));
      _nw23[(new BigNumber(7)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(8)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(9)).toNumber()] = function (_pat_let0_0) {
        return function (_132_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_133_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_133_dt__update_hcf0_h0, (_132_dt__update__tmp_h0).dtor_cf1);
            }(_pat_let1_0);
          }(_pat_let_tv5);
        }(_pat_let0_0);
      }(_130_v95);
      _nw23[(new BigNumber(10)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(11)).toNumber()] = _130_v95;
      _nw23[(new BigNumber(12)).toNumber()] = _module.D0.create_DC0(p0, _15_v0);
      _nw23[(new BigNumber(13)).toNumber()] = _130_v95;
      _131_v96 = _nw23;
      r1 = ((((!(_15_v0)) ? (_module.__default.fm4(_15_v0, _dafny.Set.fromElements(_15_v0, _15_v0), new BigNumber((_129_v94).length), globalState)) : (_15_v0))) ? (((_15_v0) ? (_131_v96) : (_131_v96))) : (_131_v96));
      r2 = ((_15_v0) ? ((p0).minus(new BigNumber((_117_v92).cardinality()))) : (p0));
      r3 = p0;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _134_v0;
      _134_v0 = new BigNumber(-250);
      let _135_v1;
      _135_v1 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_134_v0));
      let _136_v2;
      let _nw24 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _136_v2 = _nw24;
      let _137_v3;
      _137_v3 = _dafny.Set.fromElements(_136_v2);
      let _138_v4;
      _138_v4 = _dafny.Seq.of(_134_v0);
      let _139_v5;
      _139_v5 = _dafny.Set.fromElements(_138_v4);
      let _140_v6;
      _140_v6 = false;
      let _141_v7;
      _141_v7 = _dafny.MultiSet.fromElements(_140_v6, _140_v6);
      let _142_v8;
      _142_v8 = _dafny.MultiSet.fromElements(_141_v7);
      let _143_v9;
      _143_v9 = _dafny.Seq.of(!(_140_v6));
      let _144_v10;
      _144_v10 = _dafny.Seq.UnicodeFromString("jxxycbwg");
      let _145_v11;
      _145_v11 = _dafny.Map.Empty.slice().updateUnsafe(_140_v6,true);
      let _146_v12;
      _146_v12 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,new BigNumber((_dafny.Seq.UnicodeFromString("ue")).length));
      let _147_v13;
      let _nw25 = Array((new BigNumber(8)).toNumber());
      _nw25[(_dafny.ZERO).toNumber()] = _134_v0;
      _nw25[(_dafny.ONE).toNumber()] = new BigNumber((_144_v10).length);
      _nw25[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_140_v6,_145_v11)).length);
      _nw25[(new BigNumber(3)).toNumber()] = _134_v0;
      _nw25[(new BigNumber(4)).toNumber()] = new BigNumber((_138_v4).length);
      _nw25[(new BigNumber(5)).toNumber()] = new BigNumber((_144_v10).length);
      _nw25[(new BigNumber(6)).toNumber()] = new BigNumber((_146_v12).length);
      _nw25[(new BigNumber(7)).toNumber()] = _134_v0;
      _147_v13 = _nw25;
      let _148_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor(new BigNumber(870), new BigNumber(-430), _135_v1, false, new _dafny.CodePoint('x'.codePointAt(0)), _137_v3, (_139_v5).Union(_139_v5), new _dafny.CodePoint('u'.codePointAt(0)), _142_v8, false, new BigNumber(270), _143_v9, _147_v13, true);
      _148_globalState = _nw26;
      let _index20 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
      (_147_v13)[_index20] = _134_v0;
      let _index21 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
      (_147_v13)[_index21] = _module.__default.safeModuloInt(((_dafny.ZERO).minus(new BigNumber((_144_v10).length))).multipliedBy(new BigNumber(427)), (_dafny.ZERO).minus(_134_v0));
      let _149_v14;
      _149_v14 = _dafny.Map.Empty.slice().updateUnsafe((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))],_140_v6);
      let _150_v15;
      _150_v15 = _dafny.Set.fromElements(new BigNumber((_149_v14).length));
      let _151_v16;
      _151_v16 = _module.D0.create_DC2(_150_v15, _134_v0, _140_v6);
      let _source3 = (((_151_v16).dtor_cf6) ? (_151_v16) : (_151_v16));
      if (_source3.is_DC0) {
        let _152___mcc_h0 = (_source3).cf0;
        let _153___mcc_h1 = (_source3).cf1;
        let _154_cf1 = _153___mcc_h1;
        let _155_cf0 = _152___mcc_h0;
        let _156_v17;
        let _nw27 = Array((new BigNumber(5)).toNumber()).fill(false);
        _156_v17 = _nw27;
        let _index22 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_156_v17).length));
        (_156_v17)[_index22] = false;
        let _157_v18;
        _157_v18 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_144_v10, _144_v10));
        let _index23 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_156_v17).length));
        (_156_v17)[_index23] = _140_v6;
        let _index24 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_156_v17).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        let _index26 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_156_v17).length));
        let _rhs15 = _154_cf1;
        let _rhs16 = _module.__default.safeModuloInt(_155_cf0, _134_v0);
        let _rhs17 = (new BigNumber((_144_v10).length)).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt(_134_v0, _155_cf0)));
        let _rhs18 = _157_v18;
        let _rhs19 = (_134_v0).isEqualTo(_module.__default.safeModuloInt(_134_v0, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]));
        let _lhs8 = _156_v17;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_156_v17).length));
        let _lhs10 = _147_v13;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        let _lhs12 = _156_v17;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_156_v17).length));
        _lhs8[_lhs9] = _rhs15;
        _lhs10[_lhs11] = _rhs16;
        _155_cf0 = _rhs17;
        _157_v18 = _rhs18;
        _lhs12[_lhs13] = _rhs19;
        if (!(_154_cf1)) {
          let _158_v19;
          let _159_v20;
          let _160_v21;
          let _161_v22;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = _module.__default.m0(_module.__default.safeModuloInt(new BigNumber((_144_v10).length), new BigNumber(-6)), _143_v9, _148_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _158_v19 = _out0;
          _159_v20 = _out1;
          _160_v21 = _out2;
          _161_v22 = _out3;
          let _162_v23;
          let _163_v24;
          let _164_v25;
          let _165_v26;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0((_module.__default.fm0(_145_v11, _154_cf1, false, _148_globalState)).multipliedBy(new BigNumber(-280)), _143_v9, _148_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _162_v23 = _out4;
          _163_v24 = _out5;
          _164_v25 = _out6;
          _165_v26 = _out7;
          let _166_v27;
          _166_v27 = _dafny.MultiSet.fromElements(_146_v12, _146_v12, _146_v12);
          _166_v27 = _166_v27;
          _161_v22 = _164_v25;
          let _index27 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index27] = _161_v22;
        } else {
          let _167_v28;
          _167_v28 = new _dafny.CodePoint('l'.codePointAt(0));
          let _168_v29;
          _168_v29 = _dafny.MultiSet.fromElements(_167_v28);
          let _169_v30;
          _169_v30 = _dafny.Map.Empty.slice().updateUnsafe(_168_v29,_144_v10);
          let _index28 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index28] = (new BigNumber(((_149_v14).update((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], (_156_v17)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_156_v17).length))])).length)).plus(new BigNumber((((_154_cf1) ? (_169_v30) : (_169_v30))).length));
          let _170_v31;
          let _171_v32;
          let _172_v33;
          let _173_v34;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0((_dafny.ZERO).minus(new BigNumber((_139_v5).length)), _143_v9, _148_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _170_v31 = _out8;
          _171_v32 = _out9;
          _172_v33 = _out10;
          _173_v34 = _out11;
          let _174_v35;
          let _nw28 = new _module.C0();
          _nw28.__ctor(_154_cf1);
          _174_v35 = _nw28;
          let _175_v36;
          _175_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_140_v6),_155_cf0);
          let _176_v37;
          _176_v37 = _dafny.Map.Empty.slice().updateUnsafe(_175_v36,_144_v10);
          let _rhs20 = _170_v31;
          let _rhs21 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_173_v34, new BigNumber(((((_176_v37).contains(_175_v36)) ? ((_176_v37).get(_175_v36)) : (_144_v10))).length)));
          _154_cf1 = _rhs20;
          _155_cf0 = _rhs21;
          let _177_v38;
          _177_v38 = _dafny.Set.fromElements(_145_v11);
          let _178_v39;
          let _nw29 = new _module.C0();
          _nw29.__ctor((_177_v38).IsSubsetOf(_177_v38));
          _178_v39 = _nw29;
        }
        let _index29 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index29] = new BigNumber(534);
        let _179_v40;
        _179_v40 = _module.D2.create_DC6(_dafny.Seq.of((_156_v17)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_156_v17).length))]));
        let _180_v41;
        _180_v41 = _dafny.Map.Empty.slice().updateUnsafe((_155_cf0).multipliedBy(new BigNumber((_138_v4).length)),_179_v40);
        let _181_v42;
        _181_v42 = new _dafny.CodePoint('r'.codePointAt(0));
        let _rhs22 = new BigNumber((_180_v41).length);
        let _rhs23 = _181_v42;
        let _rhs24 = _module.__default.safeDivisionInt(_155_cf0, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
        let _lhs14 = _148_globalState;
        _155_cf0 = _rhs22;
        _lhs14.f4 = _rhs23;
        _155_cf0 = _rhs24;
      } else if (_source3.is_DC1) {
        let _182___mcc_h2 = (_source3).cf2;
        let _183___mcc_h3 = (_source3).cf3;
        let _184_cf3 = _183___mcc_h3;
        let _185_cf2 = _182___mcc_h2;
        let _186_v43;
        _186_v43 = _dafny.Set.fromElements(!(_140_v6));
        if (!(((_186_v43).Difference(_186_v43)).IsDisjointFrom(_dafny.Set.fromElements(!(_140_v6), _184_cf3)))) {
          let _pat_let_tv6 = _150_v15;
          let _187_v44;
          let _nw30 = Array((new BigNumber(12)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = ((_184_cf3) ? (_151_v16) : (_151_v16));
          _nw30[(_dafny.ONE).toNumber()] = _151_v16;
          _nw30[(new BigNumber(2)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(3)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(4)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(5)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(6)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(7)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(8)).toNumber()] = _module.D0.create_DC2(_150_v15, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], _140_v6);
          _nw30[(new BigNumber(9)).toNumber()] = _module.D0.create_DC2(_150_v15, _module.__default.fm0(_145_v11, _184_cf3, _140_v6, _148_globalState), false);
          _nw30[(new BigNumber(10)).toNumber()] = _151_v16;
          _nw30[(new BigNumber(11)).toNumber()] = function (_pat_let2_0) {
            return function (_188_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_189_dt__update_hcf4_h0) {
                  return _module.D0.create_DC2(_189_dt__update_hcf4_h0, (_188_dt__update__tmp_h0).dtor_cf5, (_188_dt__update__tmp_h0).dtor_cf6);
                }(_pat_let3_0);
              }(_pat_let_tv6);
            }(_pat_let2_0);
          }(_module.D0.create_DC2(_150_v15, _134_v0, _184_cf3));
          _187_v44 = _nw30;
          _187_v44 = _187_v44;
          let _index30 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index30] = (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))];
          _134_v0 = _module.__default.safeModuloInt((_134_v0).minus(_134_v0), (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          _144_v10 = _144_v10;
          let _190_v45;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_module.__default.fm4(_140_v6, _186_v43, new BigNumber((_dafny.MultiSet.fromElements((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))])).cardinality()), _148_globalState));
          _190_v45 = _nw31;
        } else {
          let _191_v46;
          _191_v46 = _module.D1.create_DC5(((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).isLessThanOrEqualTo((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]));
          _191_v46 = _191_v46;
          _134_v0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(((_149_v14).update(_134_v0, _140_v6)).length), _134_v0));
          let _192_v47;
          let _nw32 = new _module.C1();
          _nw32.__ctor(_184_cf3, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          _192_v47 = _nw32;
          let _193_v48;
          _193_v48 = _dafny.Map.Empty.slice().updateUnsafe(_192_v47,(_144_v10)[_module.__default.safeIndex(_134_v0, new BigNumber((_144_v10).length))]);
          let _194_v49;
          _194_v49 = _module.D3.create_DC8(_192_v47);
          let _195_v50;
          _195_v50 = new _dafny.CodePoint('w'.codePointAt(0));
          let _index31 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          let _rhs25 = (((_193_v48).contains((_194_v49).dtor_cf14)) ? ((_193_v48).get((_194_v49).dtor_cf14)) : (_195_v50));
          let _rhs26 = (_192_v47).f15;
          let _rhs27 = _module.__default.fm6(((_module.__default.fm6(_module.__default.fm4(!(_184_cf3), _186_v43, new BigNumber((_138_v4).length), _148_globalState), (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], _148_globalState)) ? ((_192_v47).f14) : ((_192_v47).f14)), _134_v0, _148_globalState);
          let _rhs28 = _module.__default.safeDivisionInt((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], (_192_v47).f15);
          let _lhs15 = _148_globalState;
          let _lhs16 = _147_v13;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          _lhs15.f4 = _rhs25;
          _134_v0 = _rhs26;
          _140_v6 = _rhs27;
          _lhs16[_lhs17] = _rhs28;
          (_192_v47).m1(_148_globalState);
          _134_v0 = (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))];
        }
        let _196_v51;
        _196_v51 = _module.D1.create_DC4(_144_v10);
        let _pat_let_tv7 = _144_v10;
        let _pat_let_tv8 = _144_v10;
        let _197_v52;
        let _nw33 = Array((new BigNumber(7)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _module.D1.create_DC4(_144_v10);
        _nw33[(_dafny.ONE).toNumber()] = _196_v51;
        _nw33[(new BigNumber(2)).toNumber()] = _196_v51;
        _nw33[(new BigNumber(3)).toNumber()] = _196_v51;
        _nw33[(new BigNumber(4)).toNumber()] = function (_pat_let4_0) {
          return function (_198_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_199_dt__update_hcf8_h0) {
                return _module.D1.create_DC4(_199_dt__update_hcf8_h0);
              }(_pat_let5_0);
            }(_pat_let_tv7);
          }(_pat_let4_0);
        }(_196_v51);
        _nw33[(new BigNumber(5)).toNumber()] = _196_v51;
        _nw33[(new BigNumber(6)).toNumber()] = function (_pat_let6_0) {
          return function (_200_dt__update__tmp_h2) {
            return function (_pat_let7_0) {
              return function (_201_dt__update_hcf8_h1) {
                return _module.D1.create_DC4(_201_dt__update_hcf8_h1);
              }(_pat_let7_0);
            }(_pat_let_tv8);
          }(_pat_let6_0);
        }(_196_v51);
        _197_v52 = _nw33;
        let _202_v53;
        let _nw34 = Array((new BigNumber(29)).toNumber()).fill(false);
        _202_v53 = _nw34;
        let _rhs29 = _197_v52;
        let _rhs30 = _147_v13;
        let _rhs31 = _202_v53;
        _197_v52 = _rhs29;
        _147_v13 = _rhs30;
        _202_v53 = _rhs31;
        let _203_v54;
        let _204_v55;
        let _205_v56;
        let _206_v57;
        let _out12;
        let _out13;
        let _out14;
        let _out15;
        let _outcollector3 = _module.__default.m0((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], _143_v9, _148_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _out15 = _outcollector3[3];
        _203_v54 = _out12;
        _204_v55 = _out13;
        _205_v56 = _out14;
        _206_v57 = _out15;
        let _207_v58;
        _207_v58 = _dafny.Seq.of(_143_v9, _143_v9);
        let _208_v59;
        let _209_v60;
        let _210_v61;
        let _211_v62;
        let _out16;
        let _out17;
        let _out18;
        let _out19;
        let _outcollector4 = _module.__default.m0(new BigNumber(-643), (_207_v58)[_module.__default.safeIndex(_206_v57, new BigNumber((_207_v58).length))], _148_globalState);
        _out16 = _outcollector4[0];
        _out17 = _outcollector4[1];
        _out18 = _outcollector4[2];
        _out19 = _outcollector4[3];
        _208_v59 = _out16;
        _209_v60 = _out17;
        _210_v61 = _out18;
        _211_v62 = _out19;
      } else if (_source3.is_DC2) {
        let _212___mcc_h4 = (_source3).cf4;
        let _213___mcc_h5 = (_source3).cf5;
        let _214___mcc_h6 = (_source3).cf6;
        let _215_cf6 = _214___mcc_h6;
        let _216_cf5 = _213___mcc_h5;
        let _217_cf4 = _212___mcc_h4;
        let _218_v63;
        _218_v63 = _dafny.MultiSet.fromElements(_143_v9, _143_v9, _143_v9);
        let _219_v64;
        _219_v64 = _module.D0.create_DC0(new BigNumber((_143_v9).length), _140_v6);
        let _220_v65;
        _220_v65 = _dafny.Seq.of(_143_v9);
        let _221_v66;
        _221_v66 = _dafny.Map.Empty.slice().updateUnsafe((_219_v64).dtor_cf1,_220_v65);
        let _222_v67;
        let _nw35 = Array((new BigNumber(8)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = _218_v63;
        _nw35[(_dafny.ONE).toNumber()] = (_218_v63).Union((_218_v63).update(_143_v9, _module.__default.abs(new BigNumber(596))));
        _nw35[(new BigNumber(2)).toNumber()] = ((_215_cf6) ? (_218_v63) : (_218_v63));
        _nw35[(new BigNumber(3)).toNumber()] = _218_v63;
        _nw35[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_143_v9, _143_v9);
        _nw35[(new BigNumber(5)).toNumber()] = _218_v63;
        _nw35[(new BigNumber(6)).toNumber()] = ((_215_cf6) ? (_dafny.MultiSet.FromArray((((_221_v66).contains(_215_cf6)) ? ((_221_v66).get(_215_cf6)) : (_220_v65)))) : (_218_v63));
        _nw35[(new BigNumber(7)).toNumber()] = _218_v63;
        _222_v67 = _nw35;
        let _index32 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_222_v67).length));
        (_222_v67)[_index32] = _dafny.MultiSet.fromElements(_143_v9, _143_v9, _143_v9);
        let _223_v68;
        _223_v68 = _dafny.Seq.of(_218_v63, (_218_v63).Union(_218_v63), (_218_v63).Union(_218_v63), _218_v63, _dafny.MultiSet.fromElements(_143_v9, _143_v9));
        let _index33 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_222_v67).length));
        (_222_v67)[_index33] = (_223_v68)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_138_v4)[_module.__default.safeIndex(new BigNumber(504), new BigNumber((_138_v4).length))], new BigNumber((_dafny.Seq.of(false)).length)), new BigNumber((_223_v68).length))];
        if (_140_v6) {
          let _224_v69;
          let _nw36 = new _module.C0();
          _nw36.__ctor(_module.__default.fm6(_140_v6, new BigNumber(-901), _148_globalState));
          _224_v69 = _nw36;
          _144_v10 = (_224_v69).fm5(_215_cf6, _216_cf5, _module.__default.safeDivisionInt(new BigNumber((_145_v11).length), _216_cf5), (_224_v69).f16, _148_globalState);
          let _225_v70;
          _225_v70 = _module.D4.create_DC11(_139_v5);
          let _226_v71;
          let _nw37 = new _module.C1();
          _nw37.__ctor(_module.__default.fm6(!(true), new BigNumber(999), _148_globalState), new BigNumber(((_225_v70).dtor_cf17).length));
          _226_v71 = _nw37;
          (_224_v69).m1(_148_globalState);
          let _227_v72;
          let _nw38 = new _module.C1();
          _nw38.__ctor(_dafny.Seq.IsProperPrefixOf(_144_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), function (_228_i0) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          })), new BigNumber((_145_v11).length));
          _227_v72 = _nw38;
        } else {
          let _229_v73;
          let _nw39 = new _module.C0();
          _nw39.__ctor(_215_cf6);
          _229_v73 = _nw39;
          let _230_v74;
          _230_v74 = _dafny.Seq.of(_229_v73, _229_v73);
          let _index34 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          let _rhs32 = ((_151_v16).dtor_cf5).plus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          let _rhs33 = _229_v73;
          let _rhs34 = _dafny.Seq.Concat(_138_v4, _dafny.Seq.Concat(_138_v4, _138_v4));
          let _rhs35 = _216_cf5;
          let _rhs36 = _dafny.Seq.of(_229_v73, _229_v73, _229_v73);
          let _lhs18 = _147_v13;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          _216_cf5 = _rhs32;
          _229_v73 = _rhs33;
          _138_v4 = _rhs34;
          _lhs18[_lhs19] = _rhs35;
          _230_v74 = _rhs36;
          let _index35 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index35] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_216_cf5), _138_v4)).length);
          _215_cf6 = _215_cf6;
          _215_cf6 = !(!(_140_v6)) || ((_229_v73).f16);
          _134_v0 = _module.__default.safeModuloInt(_134_v0, _134_v0);
        }
        let _index36 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index36] = _134_v0;
        if (_215_cf6) {
          _134_v0 = (_219_v64).dtor_cf0;
          let _231_v75;
          let _init3 = ((_232_v6) => function (_233_i1) {
            return !(_232_v6);
          })(_140_v6);
          let _nw40 = Array((new BigNumber(23)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw40.length); _i0_3++) {
            _nw40[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _231_v75 = _nw40;
          let _index37 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_231_v75).length));
          (_231_v75)[_index37] = _215_cf6;
          let _index38 = _module.__default.safeIndex(new BigNumber(380), new BigNumber((_231_v75).length));
          (_231_v75)[_index38] = !(_140_v6);
          _143_v9 = _dafny.Seq.update(_143_v9, _module.__default.safeIndex(_216_cf5, new BigNumber((_143_v9).length)), false);
          _140_v6 = (((_149_v14).contains(new BigNumber(157))) ? ((_149_v14).get(new BigNumber(157))) : ((_231_v75)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_231_v75).length))]));
          let _234_v76;
          _234_v76 = new _dafny.CodePoint('j'.codePointAt(0));
          let _235_v77;
          _235_v77 = _dafny.Seq.of(_144_v10);
          let _236_v78;
          let _nw41 = Array((new BigNumber(10)).toNumber()).fill(_module.D0.Default());
          _236_v78 = _nw41;
          let _237_v79;
          _237_v79 = _dafny.Map.Empty.slice().updateUnsafe(_134_v0,_141_v7);
          let _238_v80;
          _238_v80 = _dafny.Map.Empty.slice().updateUnsafe(_236_v78,_237_v79);
          let _rhs37 = _234_v76;
          let _rhs38 = _231_v75;
          let _rhs39 = _dafny.Seq.Concat(_144_v10, (_235_v77)[_module.__default.safeIndex(_216_cf5, new BigNumber((_235_v77).length))]);
          let _rhs40 = ((((_238_v80).contains(_236_v78)) ? ((_238_v80).get(_236_v78)) : (_237_v79))).contains(new BigNumber((_dafny.Seq.of((_231_v75)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_231_v75).length))])).length));
          let _rhs41 = (_231_v75)[_module.__default.safeIndex(new BigNumber(380), new BigNumber((_231_v75).length))];
          let _lhs20 = _148_globalState;
          _lhs20.f4 = _rhs37;
          _231_v75 = _rhs38;
          _144_v10 = _rhs39;
          _140_v6 = _rhs40;
          _140_v6 = _rhs41;
        } else {
          let _239_v81;
          _239_v81 = new _dafny.CodePoint('h'.codePointAt(0));
          let _index39 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_136_v2).length));
          (_136_v2)[_index39] = _239_v81;
          let _index40 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_136_v2).length));
          (_136_v2)[_index40] = _239_v81;
          _216_cf5 = (_dafny.ZERO).minus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          let _240_v82;
          _240_v82 = _dafny.MultiSet.fromElements(new BigNumber((_149_v14).length), (_dafny.ZERO).minus(_216_cf5));
          let _index41 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          let _rhs42 = _147_v13;
          let _rhs43 = !(false);
          let _rhs44 = _216_cf5;
          let _rhs45 = _240_v82;
          let _lhs21 = _148_globalState;
          let _lhs22 = _147_v13;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          _lhs21.f12 = _rhs42;
          _140_v6 = _rhs43;
          _lhs22[_lhs23] = _rhs44;
          _240_v82 = _rhs45;
          let _241_v83;
          let _242_v84;
          let _243_v85;
          let _244_v86;
          let _out20;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector5 = _module.__default.m0((_dafny.ZERO).minus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]), _module.__default.fm10(!(_215_cf6), (_136_v2)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_136_v2).length))], _134_v0, _148_globalState), _148_globalState);
          _out20 = _outcollector5[0];
          _out21 = _outcollector5[1];
          _out22 = _outcollector5[2];
          _out23 = _outcollector5[3];
          _241_v83 = _out20;
          _242_v84 = _out21;
          _243_v85 = _out22;
          _244_v86 = _out23;
          _140_v6 = (_140_v6) === (true);
        }
      } else {
        let _245___mcc_h7 = (_source3).cf7;
        let _246_cf7 = _245___mcc_h7;
        let _247_v87;
        let _nw42 = Array((new BigNumber(8)).toNumber());
        _247_v87 = _nw42;
        _247_v87 = _247_v87;
        _140_v6 = (_134_v0).isLessThan(_134_v0);
        let _248_v88;
        _248_v88 = _dafny.Map.Empty.slice().updateUnsafe((_143_v9)[_module.__default.safeIndex((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], new BigNumber((_143_v9).length))],((_140_v6) ? (_149_v14) : (_149_v14)));
        _248_v88 = (_248_v88).Merge(_248_v88);
        _140_v6 = !((_module.__default.safeDivisionInt(_134_v0, _134_v0)).isLessThan((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]));
      }
      _140_v6 = (_140_v6) || (!((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).isEqualTo((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]));
      let _249_i2;
      _249_i2 = _dafny.ZERO;
      L1: {
        while (_140_v6) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_249_i2)) {
              break L1;
            }
            _249_i2 = (_249_i2).plus(_dafny.ONE);
            _146_v12 = (_146_v12).update((_134_v0).multipliedBy(_134_v0), ((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).minus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]));
            _140_v6 = (((_140_v6) ? (_149_v14) : (_149_v14))).contains(_134_v0);
            let _250_v89;
            let _init4 = ((_251_v0, _252_v13) => function (_253_i3) {
              return (_251_v0).isLessThan((_252_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_252_v13).length))]);
            })(_134_v0, _147_v13);
            let _nw43 = Array((new BigNumber(11)).toNumber());
            for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw43.length); _i0_4++) {
              _nw43[_i0_4] = _init4(new BigNumber(_i0_4));
            }
            _250_v89 = _nw43;
            let _init5 = ((_254_v6) => function (_255_i4) {
              return _254_v6;
            })(_140_v6);
            let _nw44 = Array((new BigNumber(16)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw44.length); _i0_5++) {
              _nw44[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _250_v89 = _nw44;
            let _256_v90;
            let _nw45 = new _module.C1();
            _nw45.__ctor(_140_v6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(561)), function (_257_i5) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            })).length));
            _256_v90 = _nw45;
            let _index42 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
            let _rhs46 = _256_v90;
            let _rhs47 = new BigNumber(509);
            let _lhs24 = _147_v13;
            let _lhs25 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
            _256_v90 = _rhs46;
            _lhs24[_lhs25] = _rhs47;
          }
        }
      }
      let _258_v91;
      let _init6 = ((_259_v6) => function (_260_i6) {
        return _259_v6;
      })(_140_v6);
      let _nw46 = Array((new BigNumber(8)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw46.length); _i0_6++) {
        _nw46[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _258_v91 = _nw46;
      let _261_v92;
      let _nw47 = new _module.C0();
      _nw47.__ctor(false);
      _261_v92 = _nw47;
      let _262_v93;
      let _nw48 = Array((new BigNumber(2)).toNumber());
      _nw48[(_dafny.ZERO).toNumber()] = ((_140_v6) ? (_261_v92) : (_261_v92));
      _nw48[(_dafny.ONE).toNumber()] = _261_v92;
      _262_v93 = _nw48;
      let _index43 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length));
      (_262_v93)[_index43] = _261_v92;
      let _263_v94;
      let _nw49 = new _module.C1();
      _nw49.__ctor(_140_v6, _134_v0);
      _263_v94 = _nw49;
      let _264_v95;
      _264_v95 = _dafny.Seq.of(_263_v94, _263_v94);
      let _265_v96;
      _265_v96 = new _dafny.CodePoint('c'.codePointAt(0));
      let _266_v97;
      _266_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_264_v95).length),_265_v96);
      let _index44 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length));
      let _rhs48 = _258_v91;
      let _rhs49 = _261_v92;
      let _rhs50 = (((_266_v97).contains(_134_v0)) ? ((_266_v97).get(_134_v0)) : ((_144_v10)[_module.__default.safeIndex((_138_v4)[_module.__default.safeIndex(new BigNumber(-258), new BigNumber((_138_v4).length))], new BigNumber((_144_v10).length))]));
      let _lhs26 = _262_v93;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length));
      let _lhs28 = _148_globalState;
      _258_v91 = _rhs48;
      _lhs26[_lhs27] = _rhs49;
      _lhs28.f4 = _rhs50;
      (_261_v92).m1(_148_globalState);
      let _source4 = _151_v16;
      if (_source4.is_DC0) {
        let _267___mcc_h8 = (_source4).cf0;
        let _268___mcc_h9 = (_source4).cf1;
        let _269_cf1 = _268___mcc_h9;
        let _270_cf0 = _267___mcc_h8;
        _134_v0 = _270_cf0;
        let _271_v98;
        let _272_v99;
        let _273_v100;
        let _274_v101;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0((_263_v94).f15, _dafny.Seq.Concat(_143_v9, _dafny.Seq.of(_269_cf1, !(_269_cf1), _140_v6, (_263_v94).f14, (_263_v94).f14)), _148_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _271_v98 = _out24;
        _272_v99 = _out25;
        _273_v100 = _out26;
        _274_v101 = _out27;
        let _index45 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index45] = _module.__default.safeDivisionInt(new BigNumber(391), _270_cf0);
        _144_v10 = _dafny.Seq.UnicodeFromString("fxrlkl");
      } else if (_source4.is_DC1) {
        let _275___mcc_h10 = (_source4).cf2;
        let _276___mcc_h11 = (_source4).cf3;
        let _277_cf3 = _276___mcc_h11;
        let _278_cf2 = _275___mcc_h10;
        let _279_v102;
        let _nw50 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
        _279_v102 = _nw50;
        let _280_v103;
        _280_v103 = _dafny.Set.fromElements(false, (_263_v94).f14);
        let _index46 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_279_v102).length));
        (_279_v102)[_index46] = (_dafny.Set.fromElements(_277_cf3, (_261_v92).f16)).Difference(_280_v103);
        let _index47 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_279_v102).length));
        (_279_v102)[_index47] = (((_263_v94).f14) ? (_280_v103) : (_280_v103));
        let _index48 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index48] = (_263_v94).f15;
        let _281_v104;
        _281_v104 = _module.D1.create_DC5(true);
        let _282_v105;
        _282_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(_265_v96),(_263_v94).f15);
        let _283_v106;
        _283_v106 = _module.D0.create_DC3(_265_v96);
        let _284_v107;
        let _nw51 = new _module.C1();
        _nw51.__ctor((_281_v104).dtor_cf9, (((_282_v105).contains(_283_v106)) ? ((_282_v105).get(_283_v106)) : ((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))])));
        _284_v107 = _nw51;
        let _285_v108;
        let _nw52 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _285_v108 = _nw52;
        let _286_v109;
        let _nw53 = new _module.C0();
        _nw53.__ctor((_261_v92).f16);
        _286_v109 = _nw53;
        let _287_v110;
        _287_v110 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_263_v94).f15),_286_v109);
        let _index49 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_285_v108).length));
        (_285_v108)[_index49] = _287_v110;
        let _index50 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_285_v108).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        let _rhs51 = (_263_v94).f15;
        let _rhs52 = new BigNumber((function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_dafny.Seq.of(_144_v10, _144_v10)).Elements) {
            let _288_v111 = _compr_3;
            if (_dafny.Seq.contains(_dafny.Seq.of(_144_v10, _144_v10), _288_v111)) {
              _coll3.push([_288_v111,_module.D0.create_DC1(_278_cf2, (_263_v94).f14)]);
            }
          }
          return _coll3;
        }()).length);
        let _rhs53 = ((_287_v110).Merge(_287_v110)).Merge(_287_v110);
        let _rhs54 = new BigNumber((_141_v7).cardinality());
        let _lhs29 = _285_v108;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_285_v108).length));
        let _lhs31 = _147_v13;
        let _lhs32 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        _134_v0 = _rhs51;
        _134_v0 = _rhs52;
        _lhs29[_lhs30] = _rhs53;
        _lhs31[_lhs32] = _rhs54;
      } else if (_source4.is_DC2) {
        let _289___mcc_h12 = (_source4).cf4;
        let _290___mcc_h13 = (_source4).cf5;
        let _291___mcc_h14 = (_source4).cf6;
        let _292_cf6 = _291___mcc_h14;
        let _293_cf5 = _290___mcc_h13;
        let _294_cf4 = _289___mcc_h12;
        let _295_v112;
        let _nw54 = new _module.C0();
        _nw54.__ctor(_140_v6);
        _295_v112 = _nw54;
        _140_v6 = !(((_295_v112).f16) && ((_261_v92).f16));
        _140_v6 = ((_263_v94).f15).isLessThanOrEqualTo(new BigNumber((_module.__default.fm9(_265_v96, _134_v0, new BigNumber(904), _dafny.Map.Empty.slice().updateUnsafe(!((_263_v94).f14),_293_cf5), _148_globalState)).length));
        let _296_v113;
        let _nw55 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
        _296_v113 = _nw55;
        let _297_v114;
        _297_v114 = _dafny.Map.Empty.slice().updateUnsafe(_144_v10,_147_v13);
        let _index52 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_296_v113).length));
        (_296_v113)[_index52] = _297_v114;
        let _index53 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_296_v113).length));
        (_296_v113)[_index53] = (_dafny.Map.Empty.slice().updateUnsafe(_144_v10,_147_v13)).Merge((_297_v114).Merge(_297_v114));
      } else {
        let _298___mcc_h15 = (_source4).cf7;
        let _299_cf7 = _298___mcc_h15;
        let _index54 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index54] = (_263_v94).f15;
        if ((_263_v94).f14) {
          let _300_v115;
          _300_v115 = _module.D3.create_DC10(_module.D3.create_DC9(_299_cf7));
          let _pat_let_tv9 = _263_v94;
          let _301_v116;
          _301_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_261_v92).f16, (_143_v9)[_module.__default.safeIndex(new BigNumber((_264_v95).length), new BigNumber((_143_v9).length))])).length),function (_pat_let8_0) {
            return function (_302_dt__update__tmp_h3) {
              return function (_pat_let9_0) {
                return function (_303_dt__update_hcf16_h0) {
                  return _module.D3.create_DC10(_303_dt__update_hcf16_h0);
                }(_pat_let9_0);
              }(_module.D3.create_DC8(_pat_let_tv9));
            }(_pat_let8_0);
          }(_300_v115));
          _301_v116 = (_301_v116).update((_263_v94).f15, _module.D3.create_DC10(_module.D3.create_DC9(_265_v96)));
          let _304_v117;
          let _nw56 = new _module.C1();
          _nw56.__ctor((_261_v92).f16, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          _304_v117 = _nw56;
          _304_v117 = _304_v117;
          let _305_v118;
          _305_v118 = _module.D3.create_DC8(_263_v94);
          _305_v118 = _305_v118;
          _144_v10 = _144_v10;
          let _306_v119;
          _306_v119 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6((_261_v92).f16, (_263_v94).f15, _148_globalState),_143_v9);
          _306_v119 = _306_v119;
        } else {
          let _307_v120;
          _307_v120 = _dafny.MultiSet.fromElements(_299_cf7);
          let _308_v121;
          _308_v121 = _dafny.Map.Empty.slice().updateUnsafe(_140_v6,_307_v120);
          let _309_v122;
          _309_v122 = _dafny.Set.fromElements(false, !((_263_v94).f14));
          let _310_v123;
          _310_v123 = _dafny.Seq.of((((_308_v121).contains((_261_v92).f16)) ? ((_308_v121).get((_261_v92).f16)) : ((_307_v120).update(new _dafny.CodePoint('j'.codePointAt(0)), _module.__default.abs(new BigNumber((_309_v122).length))))));
          let _index55 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index55] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Concat(_310_v123, _310_v123)).length), (_263_v94).f15);
          _134_v0 = (_263_v94).f15;
          let _311_v124;
          _311_v124 = _module.D1.create_DC5(_140_v6);
          let _312_v125;
          _312_v125 = _dafny.Map.Empty.slice().updateUnsafe((_263_v94).f15,_262_v93);
          let _index56 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length));
          let _rhs55 = _147_v13;
          let _rhs56 = _311_v124;
          let _rhs57 = (((_145_v11).contains((((_145_v11).contains((_261_v92).f16)) ? ((_145_v11).get((_261_v92).f16)) : ((_261_v92).f16)))) ? ((_145_v11).get((((_145_v11).contains((_261_v92).f16)) ? ((_145_v11).get((_261_v92).f16)) : ((_261_v92).f16)))) : ((_261_v92).f16));
          let _rhs58 = (_312_v125).Merge((_312_v125).update(_134_v0, _262_v93));
          let _rhs59 = _261_v92;
          let _lhs33 = _148_globalState;
          let _lhs34 = _262_v93;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length));
          _lhs33.f12 = _rhs55;
          _311_v124 = _rhs56;
          _140_v6 = _rhs57;
          _312_v125 = _rhs58;
          _lhs34[_lhs35] = _rhs59;
          _134_v0 = _134_v0;
          let _index57 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          let _rhs60 = (_module.__default.safeDivisionInt((_263_v94).f15, _134_v0)).plus(_module.__default.safeDivisionInt((_263_v94).f15, (_263_v94).f15));
          let _rhs61 = _145_v11;
          let _rhs62 = _299_cf7;
          let _rhs63 = _dafny.Seq.Concat(_144_v10, _144_v10);
          let _lhs36 = _147_v13;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          _lhs36[_lhs37] = _rhs60;
          _145_v11 = _rhs61;
          _299_cf7 = _rhs62;
          _144_v10 = _rhs63;
        }
        _147_v13 = _147_v13;
        _140_v6 = _dafny.Seq.IsProperPrefixOf(_138_v4, _138_v4);
      }
      let _313_v126;
      _313_v126 = _dafny.MultiSet.fromElements(_265_v96);
      _313_v126 = ((_313_v126).Intersect(_module.__default.fm11(_dafny.Seq.UnicodeFromString("avyxie"), _148_globalState))).update(_265_v96, _module.__default.abs(((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).multipliedBy((_263_v94).f15)));
      _140_v6 = (new BigNumber((_144_v10).length)).isLessThan((_263_v94).f15);
      let _pat_let_tv10 = _151_v16;
      let _pat_let_tv11 = _150_v15;
      let _pat_let_tv12 = _263_v94;
      let _pat_let_tv13 = _151_v16;
      let _source5 = function (_source6) {
        if (_source6.is_DC12) {
          let _314___mcc_h24 = (_source6).cf18;
          let _315___mcc_h25 = (_source6).cf19;
          let _316___mcc_h26 = (_source6).cf20;
          let _317_cf20 = _316___mcc_h26;
          let _318_cf19 = _315___mcc_h25;
          let _319_cf18 = _314___mcc_h24;
          return _pat_let_tv10;
        } else if (_source6.is_DC11) {
          let _320___mcc_h27 = (_source6).cf17;
          let _321_cf17 = _320___mcc_h27;
          return _module.D0.create_DC2(_pat_let_tv11, (_pat_let_tv12).f15, false);
        } else {
          let _322___mcc_h28 = (_source6).cf21;
          let _323_cf21 = _322___mcc_h28;
          return _pat_let_tv13;
        }
      }(_module.__default.fm12(_140_v6, (_263_v94).f15, _148_globalState));
      if (_source5.is_DC0) {
        let _324___mcc_h16 = (_source5).cf0;
        let _325___mcc_h17 = (_source5).cf1;
        let _326_cf1 = _325___mcc_h17;
        let _327_cf0 = _324___mcc_h16;
        let _328_v127;
        let _nw57 = new _module.C1();
        _nw57.__ctor((_261_v92).f16, (_263_v94).f15);
        _328_v127 = _nw57;
        _140_v6 = _140_v6;
        if ((_328_v127).f14) {
          let _329_v128;
          let _nw58 = new _module.C1();
          _nw58.__ctor(true, (_263_v94).f15);
          _329_v128 = _nw58;
          let _330_v129;
          _330_v129 = _dafny.MultiSet.fromElements((_263_v94).f15, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]);
          let _331_v130;
          _331_v130 = _module.D4.create_DC12((_263_v94).f15, _265_v96, _330_v129);
          let _332_v131;
          _332_v131 = _dafny.Map.Empty.slice().updateUnsafe(_138_v4,_331_v130);
          let _333_v132;
          _333_v132 = _dafny.Map.Empty.slice().updateUnsafe(_329_v128,new BigNumber((_332_v131).length));
          _326_cf1 = (_333_v132).equals(_333_v132);
          let _334_v133;
          _334_v133 = _dafny.Map.Empty.slice().updateUnsafe((_328_v127).f14,_143_v9);
          let _335_v134;
          _335_v134 = _dafny.Map.Empty.slice().updateUnsafe(_334_v133,_147_v13);
          _147_v13 = (((_335_v134).contains((_334_v133).update((_263_v94).f14, _143_v9))) ? ((_335_v134).get((_334_v133).update((_263_v94).f14, _143_v9))) : (((!(_module.__default.fm6((_328_v127).f14, (_263_v94).f15, _148_globalState))) ? (_147_v13) : (_147_v13))));
          let _index58 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index58] = ((_dafny.ZERO).minus((_263_v94).f15)).multipliedBy((_263_v94).f15);
          let _336_v135;
          let _nw59 = Array((new BigNumber(28)).toNumber()).fill(_module.D3.Default());
          _336_v135 = _nw59;
          let _337_v136;
          _337_v136 = _module.D3.create_DC8(_263_v94);
          let _index59 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_336_v135).length));
          (_336_v135)[_index59] = _337_v136;
          let _index60 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_336_v135).length));
          (_336_v135)[_index60] = _module.D3.create_DC8(_328_v127);
          (_148_globalState).f4 = new _dafny.CodePoint('w'.codePointAt(0));
        } else {
          let _index61 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
          (_147_v13)[_index61] = (_263_v94).fm1(_144_v10, (_328_v127).f15, (_263_v94).f15, _148_globalState);
          (_263_v94).m1(_148_globalState);
          _140_v6 = _326_cf1;
          let _index62 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_258_v91).length));
          (_258_v91)[_index62] = !(true);
          let _338_v137;
          _338_v137 = _dafny.MultiSet.fromElements((_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))], (_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))], _261_v92, _261_v92, (_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))]);
          let _index63 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_258_v91).length));
          (_258_v91)[_index63] = (_338_v137).IsDisjointFrom(_338_v137);
          let _339_v138;
          _339_v138 = _dafny.Set.fromElements((_258_v91)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_258_v91).length))], (_261_v92).f16);
          _140_v6 = _module.__default.fm4((_263_v94).f14, _339_v138, (_138_v4)[_module.__default.safeIndex((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))], new BigNumber((_138_v4).length))], _148_globalState);
        }
        let _340_v139;
        _340_v139 = _dafny.Map.Empty.slice().updateUnsafe(_141_v7,(((_263_v94).f14) ? ((_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))]) : (_261_v92)));
        let _341_v140;
        _341_v140 = _dafny.Seq.of(_dafny.Seq.Concat(_144_v10, _144_v10), _dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_342_v96) => function (_343_i7) {
          return _342_v96;
        })(_265_v96)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), ((_344_v96) => function (_345_i8) {
          return _344_v96;
        })(_265_v96)));
        let _346_v141;
        _346_v141 = _dafny.Set.fromElements(_140_v6, _326_cf1, _326_cf1, (_261_v92).f16);
        let _347_v142;
        _347_v142 = _dafny.Map.Empty.slice().updateUnsafe((_261_v92).f16,(_263_v94).f15);
        let _348_v143;
        _348_v143 = _dafny.Set.fromElements(_346_v141, _346_v141, _346_v141, _module.__default.fm9(_265_v96, _327_cf0, (_263_v94).f15, (_347_v142).update(_140_v6, _327_cf0), _148_globalState), _346_v141);
        let _rhs64 = (_150_v15).IsSubsetOf(_dafny.Set.fromElements((_263_v94).f15));
        let _rhs65 = _module.__default.fm6((false) && ((_328_v127).f14), _327_cf0, _148_globalState);
        let _rhs66 = _dafny.Map.Empty.slice().updateUnsafe(_141_v7,(_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))]);
        let _rhs67 = (_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))];
        let _rhs68 = _module.__default.fm13(_module.D0.create_DC1(_348_v143, true), (_143_v9)[_module.__default.safeIndex(new BigNumber(-196), new BigNumber((_143_v9).length))], _148_globalState);
        _326_cf1 = _rhs64;
        _140_v6 = _rhs65;
        _340_v139 = _rhs66;
        _261_v92 = _rhs67;
        _341_v140 = _rhs68;
      } else if (_source5.is_DC1) {
        let _349___mcc_h18 = (_source5).cf2;
        let _350___mcc_h19 = (_source5).cf3;
        let _351_cf3 = _350___mcc_h19;
        let _352_cf2 = _349___mcc_h18;
        let _353_v144;
        let _nw60 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _353_v144 = _nw60;
        let _index64 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_353_v144).length));
        (_353_v144)[_index64] = _144_v10;
        let _index65 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_353_v144).length));
        (_353_v144)[_index65] = _144_v10;
        let _index66 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
        (_147_v13)[_index66] = (_263_v94).f15;
        (_263_v94).m1(_148_globalState);
        let _354_v145;
        _354_v145 = _dafny.Map.Empty.slice().updateUnsafe(_258_v91,new BigNumber((_144_v10).length));
        _354_v145 = (_354_v145).update(_258_v91, (_263_v94).f15);
      } else if (_source5.is_DC2) {
        let _355___mcc_h20 = (_source5).cf4;
        let _356___mcc_h21 = (_source5).cf5;
        let _357___mcc_h22 = (_source5).cf6;
        let _358_cf6 = _357___mcc_h22;
        let _359_cf5 = _356___mcc_h21;
        let _360_cf4 = _355___mcc_h20;
        _140_v6 = (_263_v94).f14;
        _140_v6 = false;
        let _361_v146;
        let _nw61 = new _module.C1();
        _nw61.__ctor(!(((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).isLessThan(_359_cf5)), _module.__default.safeDivisionInt(_134_v0, (_263_v94).f15));
        _361_v146 = _nw61;
        let _362_v147;
        let _363_v148;
        let _364_v149;
        let _365_v150;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector7 = _module.__default.m0((((_263_v94).f14) ? (_134_v0) : (new BigNumber(524))), _143_v9, _148_globalState);
        _out28 = _outcollector7[0];
        _out29 = _outcollector7[1];
        _out30 = _outcollector7[2];
        _out31 = _outcollector7[3];
        _362_v147 = _out28;
        _363_v148 = _out29;
        _364_v149 = _out30;
        _365_v150 = _out31;
      } else {
        let _366___mcc_h23 = (_source5).cf7;
        let _367_cf7 = _366___mcc_h23;
        _263_v94 = _263_v94;
        let _index67 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length));
        (_258_v91)[_index67] = true;
        let _368_v151;
        _368_v151 = _dafny.Set.fromElements(_367_cf7);
        let _369_v153;
        _369_v153 = _dafny.Seq.of(_368_v151, _368_v151);
        let _370_v155;
        let _nw62 = Array((new BigNumber(20)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _368_v151;
        _nw62[(_dafny.ONE).toNumber()] = _368_v151;
        _nw62[(new BigNumber(2)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(3)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_367_cf7);
        _nw62[(new BigNumber(5)).toNumber()] = _module.__default.fm14(new BigNumber(456), _367_cf7, (_263_v94).f15, (_263_v94).f15, _148_globalState);
        _nw62[(new BigNumber(6)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(7)).toNumber()] = function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of ((_369_v153)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_369_v153).length))]).Elements) {
              let _371_v152 = _compr_5;
              if (((_369_v153)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_369_v153).length))]).contains(_371_v152)) {
                _coll5.push([_371_v152,_367_cf7]);
              }
            }
            return _coll5;
          }()).Keys.Elements) {
            let _372_v154 = _compr_4;
            if ((function () {
              let _coll6 = new _dafny.Map();
              for (const _compr_6 of ((_369_v153)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_369_v153).length))]).Elements) {
                let _371_v152 = _compr_6;
                if (((_369_v153)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_369_v153).length))]).contains(_371_v152)) {
                  _coll6.push([_371_v152,_367_cf7]);
                }
              }
              return _coll6;
            }()).contains(_372_v154)) {
              _coll4.add(_372_v154);
            }
          }
          return _coll4;
        }();
        _nw62[(new BigNumber(8)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_module.__default.fm15(new _dafny.CodePoint('k'.codePointAt(0)), _dafny.Map.Empty.slice().updateUnsafe((_263_v94).f15,_134_v0), _148_globalState));
        _nw62[(new BigNumber(10)).toNumber()] = (((_261_v92).f16) ? (_368_v151) : (_368_v151));
        _nw62[(new BigNumber(11)).toNumber()] = ((!(false)) ? (_368_v151) : (_368_v151));
        _nw62[(new BigNumber(12)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(13)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(14)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(15)).toNumber()] = (_368_v151).Union(_368_v151);
        _nw62[(new BigNumber(16)).toNumber()] = _368_v151;
        _nw62[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(_265_v96, _367_cf7);
        _nw62[(new BigNumber(18)).toNumber()] = _module.__default.fm14(_134_v0, _367_cf7, new BigNumber((_141_v7).cardinality()), new BigNumber(42), _148_globalState);
        _nw62[(new BigNumber(19)).toNumber()] = (_368_v151).Union(_368_v151);
        _370_v155 = _nw62;
        let _373_v156;
        _373_v156 = _dafny.Map.Empty.slice().updateUnsafe(_143_v9,true);
        let _374_v157;
        _374_v157 = _dafny.Seq.of((_262_v93)[_module.__default.safeIndex(new BigNumber(538), new BigNumber((_262_v93).length))], _261_v92);
        let _375_v158;
        let _nw63 = new _module.C1();
        _nw63.__ctor(_module.__default.fm6((_263_v94).f14, (_263_v94).f15, _148_globalState), new BigNumber((_dafny.Seq.of(new BigNumber(819), new BigNumber((_373_v156).length), new BigNumber((_374_v157).length), _module.__default.fm0(_145_v11, !(false), (_263_v94).f14, _148_globalState))).length));
        _375_v158 = _nw63;
        let _376_v159;
        _376_v159 = _dafny.Map.Empty.slice().updateUnsafe((((_263_v94).f14) ? (_263_v94) : (_263_v94)),_375_v158);
        let _377_v160;
        _377_v160 = _dafny.Seq.of(_370_v155);
        let _index68 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length));
        let _rhs69 = _dafny.Seq.Concat(_dafny.Seq.of((_263_v94).f14), _143_v9);
        let _rhs70 = (_143_v9)[_module.__default.safeIndex(_134_v0, new BigNumber((_143_v9).length))];
        let _rhs71 = (_261_v92).f16;
        let _rhs72 = new BigNumber((_376_v159).length);
        let _rhs73 = (_377_v160)[_module.__default.safeIndex(_134_v0, new BigNumber((_377_v160).length))];
        let _lhs38 = _258_v91;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length));
        _143_v9 = _rhs69;
        _lhs38[_lhs39] = _rhs70;
        _140_v6 = _rhs71;
        _134_v0 = _rhs72;
        _370_v155 = _rhs73;
        let _378_v161;
        _378_v161 = _dafny.Set.fromElements((_261_v92).f16, _140_v6, (_258_v91)[_module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length))], (_261_v92).f16, _140_v6);
        _140_v6 = _module.__default.fm4((_140_v6) || ((_261_v92).f16), _378_v161, _module.__default.safeModuloInt(_134_v0, new BigNumber(-752)), _148_globalState);
        let _379_v162;
        _379_v162 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_263_v94).f14, !((_258_v91)[_module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length))]), (_263_v94).f14, (_261_v92).f16, (_258_v91)[_module.__default.safeIndex(new BigNumber(615), new BigNumber((_258_v91).length))])).length), _134_v0);
        let _380_v163;
        _380_v163 = _dafny.Map.Empty.slice().updateUnsafe(_367_cf7,(_263_v94).f14);
        let _381_v164;
        _381_v164 = _dafny.Seq.of(_379_v162);
        let _382_v165;
        let _nw64 = Array((new BigNumber(13)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _379_v162;
        _nw64[(_dafny.ONE).toNumber()] = (_module.__default.fm16(_140_v6, (((_149_v14).contains(new BigNumber((_380_v163).length))) ? ((_149_v14).get(new BigNumber((_380_v163).length))) : ((((_145_v11).contains((_261_v92).f16)) ? ((_145_v11).get((_261_v92).f16)) : ((_263_v94).f14)))), _148_globalState)).Union(_379_v162);
        _nw64[(new BigNumber(2)).toNumber()] = (_379_v162).Difference(_dafny.MultiSet.fromElements(new BigNumber(-260)));
        _nw64[(new BigNumber(3)).toNumber()] = _module.__default.fm16((_263_v94).f14, (_263_v94).f14, _148_globalState);
        _nw64[(new BigNumber(4)).toNumber()] = (_379_v162).Difference(_dafny.MultiSet.fromElements(new BigNumber(-871)));
        _nw64[(new BigNumber(5)).toNumber()] = (_381_v164)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_381_v164).length))];
        _nw64[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm0(_145_v11, (_261_v92).f16, (_261_v92).f16, _148_globalState));
        _nw64[(new BigNumber(7)).toNumber()] = (_379_v162).Difference(_379_v162);
        _nw64[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements((_263_v94).f15)).Intersect(_379_v162);
        _nw64[(new BigNumber(9)).toNumber()] = (_381_v164)[_module.__default.safeIndex(_134_v0, new BigNumber((_381_v164).length))];
        _nw64[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(_134_v0);
        _nw64[(new BigNumber(11)).toNumber()] = (_379_v162).Union(_379_v162);
        _nw64[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_143_v9).length), _134_v0);
        _382_v165 = _nw64;
        let _index69 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_382_v165).length));
        (_382_v165)[_index69] = _379_v162;
        let _index70 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_382_v165).length));
        (_382_v165)[_index70] = (_379_v162).Intersect(_379_v162);
      }
      let _383_i9;
      _383_i9 = _dafny.ZERO;
      L2: {
        while (!(_140_v6)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_383_i9)) {
              break L2;
            }
            _383_i9 = (_383_i9).plus(_dafny.ONE);
            let _384_v166;
            _384_v166 = _module.D3.create_DC9(_265_v96);
            let _385_v167;
            _385_v167 = _module.D3.create_DC10(_384_v166);
            let _386_v168;
            _386_v168 = _dafny.Set.fromElements(_module.D3.create_DC10(_384_v166), _385_v167);
            let _387_v169;
            _387_v169 = _dafny.Seq.of(_386_v168, _386_v168, _386_v168);
            let _index71 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
            (_147_v13)[_index71] = new BigNumber(((_387_v169)[_module.__default.safeIndex(((true) ? ((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]) : ((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))])), new BigNumber((_387_v169).length))]).length);
            let _388_v170;
            let _nw65 = Array((new BigNumber(19)).toNumber()).fill([]);
            _388_v170 = _nw65;
            let _389_v171;
            _389_v171 = _module.D5.create_DC14(_388_v170);
            _388_v170 = (((_140_v6) ? (_389_v171) : (_389_v171))).dtor_cf22;
            let _390_v172;
            let _nw66 = new _module.C1();
            _nw66.__ctor(_140_v6, (_263_v94).f15);
            _390_v172 = _nw66;
            let _391_v173;
            _391_v173 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(_140_v6, _141_v7, new BigNumber((_dafny.Seq.update(_144_v10, _module.__default.safeIndex((_390_v172).f15, new BigNumber((_144_v10).length)), _265_v96)).length), _144_v10, _148_globalState),_258_v91);
            let _392_v174;
            _392_v174 = _dafny.Seq.of(_138_v4, _dafny.Seq.of((_263_v94).f15, (_dafny.ZERO).minus(new BigNumber((_144_v10).length))));
            _391_v173 = (_391_v173).update(_392_v174, _258_v91);
          }
        }
      }
      let _393_i10;
      _393_i10 = _dafny.ZERO;
      L3: {
        while ((_140_v6) && (false)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_393_i10)) {
              break L3;
            }
            _393_i10 = (_393_i10).plus(_dafny.ONE);
            let _394_v175;
            _394_v175 = _dafny.Map.Empty.slice().updateUnsafe((_263_v94).f14,_134_v0);
            _394_v175 = ((_394_v175).Merge((_394_v175).update(true, (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]))).Merge(_394_v175);
            (_261_v92).m1(_148_globalState);
            let _395_v176;
            let _nw67 = new _module.C0();
            _nw67.__ctor((_261_v92).f16);
            _395_v176 = _nw67;
            let _396_v178;
            _396_v178 = _module.D0.create_DC1(function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-916)), ((_397_v92) => function (_398_i11) {
    return _dafny.Set.fromElements((_397_v92).f16, (_397_v92).f16, !((_397_v92).f16));
  })(_261_v92))).Elements) {
    let _399_v177 = _compr_7;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-916)), ((_400_v92) => function (_398_i11) {
      return _dafny.Set.fromElements((_400_v92).f16, (_400_v92).f16, !((_400_v92).f16));
    })(_261_v92)), _399_v177)) {
      _coll7.add(_399_v177);
    }
  }
  return _coll7;
}(), (_263_v94).f14);
            let _index72 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_258_v91).length));
            (_258_v91)[_index72] = (_396_v178).dtor_cf3;
            let _index73 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_258_v91).length));
            (_258_v91)[_index73] = !(_394_v175).contains((_263_v94).f14);
          }
        }
      }
      let _401_v179;
      let _nw68 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
      _401_v179 = _nw68;
      let _index74 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_401_v179).length));
      (_401_v179)[_index74] = function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_138_v4).Elements) {
          let _402_v180 = _compr_8;
          if (_dafny.Seq.contains(_138_v4, _402_v180)) {
            _coll8.push([(_402_v180).plus((_263_v94).f15),(_263_v94).f14]);
          }
        }
        return _coll8;
      }();
      let _index75 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_401_v179).length));
      (_401_v179)[_index75] = _149_v14;
      (_261_v92).m1(_148_globalState);
      let _403_i12;
      _403_i12 = _dafny.ZERO;
      L4: {
        while (!((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]).isEqualTo(new BigNumber((_dafny.Seq.Concat(_144_v10, _dafny.Seq.update(_144_v10, _module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))])), new BigNumber((_144_v10).length)), _265_v96))).length))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_403_i12)) {
              break L4;
            }
            _403_i12 = (_403_i12).plus(_dafny.ONE);
            let _404_v181;
            let _nw69 = new _module.C1();
            _nw69.__ctor(_140_v6, (_263_v94).f15);
            _404_v181 = _nw69;
            let _405_v182;
            _405_v182 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(784),_404_v181);
            let _406_v183;
            _406_v183 = _module.D3.create_DC8(_263_v94);
            let _407_v184;
            let _nw70 = Array((new BigNumber(3)).toNumber());
            _nw70[(_dafny.ZERO).toNumber()] = _406_v183;
            _nw70[(_dafny.ONE).toNumber()] = _406_v183;
            _nw70[(new BigNumber(2)).toNumber()] = _406_v183;
            _407_v184 = _nw70;
            let _408_v185;
            _408_v185 = _dafny.Map.Empty.slice().updateUnsafe((((_405_v182).contains(_134_v0)) ? ((_405_v182).get(_134_v0)) : (_404_v181)),_407_v184);
            let _409_v186;
            _409_v186 = _dafny.Map.Empty.slice().updateUnsafe(_407_v184,_147_v13);
            _140_v6 = ((_409_v186).Merge(_dafny.Map.Empty.slice().updateUnsafe(_407_v184,_147_v13))).contains((((_408_v185).contains(_404_v181)) ? ((_408_v185).get(_404_v181)) : (_407_v184)));
            if ((_263_v94).f14) {
              _140_v6 = (_263_v94).f14;
              _140_v6 = (_261_v92).f16;
              let _410_v187;
              let _nw71 = new _module.C1();
              _nw71.__ctor((_261_v92).f16, (_263_v94).f15);
              _410_v187 = _nw71;
              _140_v6 = _module.__default.fm6((_263_v94).f14, (_263_v94).f15, _148_globalState);
              let _411_v188;
              _411_v188 = _dafny.Set.fromElements((_263_v94).f14);
              let _412_v189;
              _412_v189 = _dafny.MultiSet.fromElements(_411_v188);
              let _index76 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_258_v91).length));
              (_258_v91)[_index76] = (_412_v189).IsProperSubsetOf(_412_v189);
              let _index77 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_258_v91).length));
              (_258_v91)[_index77] = _140_v6;
            } else {
              (_148_globalState).f4 = _265_v96;
              _147_v13 = _147_v13;
              let _413_v190;
              _413_v190 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.of((_263_v94).f14), _module.__default.safeIndex((_138_v4)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_138_v4).length))], new BigNumber((_dafny.Seq.of((_263_v94).f14)).length)), _140_v6), _143_v9, _dafny.Seq.Concat(_143_v9, _143_v9), _143_v9);
              _413_v190 = _dafny.Seq.Concat(_413_v190, _413_v190);
              let _index78 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_258_v91).length));
              (_258_v91)[_index78] = (_263_v94).f14;
              let _414_v191;
              _414_v191 = _dafny.Set.fromElements(_140_v6, (_263_v94).f14);
              let _index79 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_258_v91).length));
              (_258_v91)[_index79] = ((_140_v6) ? (((_263_v94).f15).isLessThan((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))])) : (_module.__default.fm4((_263_v94).f14, _414_v191, _134_v0, _148_globalState)));
              _405_v182 = (_405_v182).update((_134_v0).plus((_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))]), _404_v181);
            }
            if ((_261_v92).f16) {
              let _415_v192;
              _415_v192 = _dafny.Seq.of(_404_v181, _404_v181, _404_v181, _404_v181, _404_v181);
              let _416_v193;
              let _nw72 = Array((_dafny.ONE).toNumber());
              _nw72[(_dafny.ZERO).toNumber()] = _415_v192;
              _416_v193 = _nw72;
              let _index80 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_416_v193).length));
              (_416_v193)[_index80] = _415_v192;
              let _index81 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_416_v193).length));
              (_416_v193)[_index81] = _415_v192;
              _144_v10 = _dafny.Seq.Concat(_144_v10, _144_v10);
              let _index82 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
              (_147_v13)[_index82] = ((_263_v94).f15).multipliedBy(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nadw")).length),(_263_v94).f14)).Merge((_401_v179)[_module.__default.safeIndex(new BigNumber(534), new BigNumber((_401_v179).length))])).length));
              let _417_v194;
              let _nw73 = Array((new BigNumber(9)).toNumber());
              _nw73[(_dafny.ZERO).toNumber()] = _404_v181;
              _nw73[(_dafny.ONE).toNumber()] = _404_v181;
              _nw73[(new BigNumber(2)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(3)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(4)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(5)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(6)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(7)).toNumber()] = _404_v181;
              _nw73[(new BigNumber(8)).toNumber()] = _404_v181;
              _417_v194 = _nw73;
              _417_v194 = _417_v194;
              let _418_v195;
              let _nw74 = new _module.C1();
              _nw74.__ctor((_261_v92).f16, ((false) ? ((_263_v94).f15) : ((_dafny.ZERO).minus((_404_v181).fm1(_144_v10, new BigNumber(-329), _134_v0, _148_globalState)))));
              _418_v195 = _nw74;
            } else {
              let _419_v196;
              _419_v196 = _dafny.Set.fromElements(_404_v181, _404_v181);
              let _420_v197;
              _420_v197 = _dafny.Set.fromElements((_419_v196));
              let _421_v198;
              _421_v198 = _dafny.Set.fromElements(_404_v181);
              let _422_v199;
              _422_v199 = _dafny.Seq.of(_420_v197);
              let _423_v200;
              let _nw75 = Array((new BigNumber(25)).toNumber());
              _nw75[(_dafny.ZERO).toNumber()] = _420_v197;
              _nw75[(_dafny.ONE).toNumber()] = _420_v197;
              _nw75[(new BigNumber(2)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(3)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(4)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(5)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(6)).toNumber()] = (_420_v197).Intersect(_420_v197);
              _nw75[(new BigNumber(7)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(8)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(9)).toNumber()] = (_420_v197).Intersect(_420_v197);
              _nw75[(new BigNumber(10)).toNumber()] = (_420_v197).Difference(_dafny.Set.fromElements(_421_v198));
              _nw75[(new BigNumber(11)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(12)).toNumber()] = (_420_v197).Difference(_420_v197);
              _nw75[(new BigNumber(13)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(14)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(15)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(16)).toNumber()] = (_420_v197).Difference(_420_v197);
              _nw75[(new BigNumber(17)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(18)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(19)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(20)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(21)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(22)).toNumber()] = _420_v197;
              _nw75[(new BigNumber(23)).toNumber()] = (_422_v199)[_module.__default.safeIndex((_263_v94).f15, new BigNumber((_422_v199).length))];
              _nw75[(new BigNumber(24)).toNumber()] = _420_v197;
              _423_v200 = _nw75;
              let _index83 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_423_v200).length));
              (_423_v200)[_index83] = _420_v197;
              let _index84 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_423_v200).length));
              (_423_v200)[_index84] = _420_v197;
              _140_v6 = (_261_v92).f16;
              (_261_v92).m1(_148_globalState);
              let _424_v201;
              _424_v201 = _module.D0.create_DC0((_263_v94).f15, (_263_v94).f14);
              let _425_v202;
              _425_v202 = _dafny.Map.Empty.slice().updateUnsafe((_263_v94).f14,_dafny.Map.Empty.slice().updateUnsafe((_263_v94).f15,(_261_v92).f16));
              _424_v201 = _module.D0.create_DC0(new BigNumber((((_425_v202).update(false, _149_v14)).Merge(_425_v202)).length), !((_261_v92).f16) || (_140_v6));
              let _index85 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length));
              (_147_v13)[_index85] = (_147_v13)[_module.__default.safeIndex(new BigNumber(749), new BigNumber((_147_v13).length))];
            }
            let _426_v203;
            _426_v203 = _dafny.Set.fromElements(_258_v91, _258_v91);
            _426_v203 = (_426_v203).Intersect(_426_v203);
          }
        }
      }
      let _427_v204;
      _427_v204 = _module.D1.create_DC5(_140_v6);
      let _pat_let_tv14 = _145_v11;
      let _pat_let_tv15 = _263_v94;
      let _pat_let_tv16 = _145_v11;
      let _pat_let_tv17 = _263_v94;
      let _pat_let_tv18 = _263_v94;
      _140_v6 = function (_source7) {
        if (_source7.is_DC5) {
          let _428___mcc_h29 = (_source7).cf9;
          let _429_cf9 = _428___mcc_h29;
          return _429_cf9;
        } else {
          let _430___mcc_h30 = (_source7).cf8;
          let _431_cf8 = _430___mcc_h30;
          if ((_pat_let_tv14).contains((_pat_let_tv15).f14)) {
            return (_pat_let_tv16).get((_pat_let_tv17).f14);
          } else {
            return (_pat_let_tv18).f14;
          }
        }
      }(_427_v204);
      process.stdout.write(_dafny.toString(_134_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v1).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-250))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v2)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_137_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_138_v4, _dafny.Seq.of(new BigNumber(-250), new BigNumber(-250), new BigNumber(-250)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v5).equals(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-250))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_140_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v7).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_142_v8).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_143_v9, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_144_v10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-250),new BigNumber(2)).updateUnsafe(_dafny.ZERO,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_148_globalState).f2).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(-250))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_148_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_148_globalState.f5).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f6).equals(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-250))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_148_globalState).f8).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_148_globalState).f11, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState.f12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_149_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(84),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_150_v15).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_151_v16).dtor_cf4).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v16).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v16).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_249_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v91)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_261_v92).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_262_v93)[_dafny.ZERO]).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_262_v93)[_dafny.ONE]).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v94).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_263_v94).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_264_v95).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_265_v96));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_266_v97).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),new _dafny.CodePoint('c'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_313_v126).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_383_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_393_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_401_v179)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(84),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_403_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_427_v204).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC1(cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4, cf5, cf6) {
      let $dt = new D0(2);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D0(3);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0) && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO, false);
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
    static create_DC5(cf9) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + this.cf8.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false);
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
    static create_DC7(cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(false, false, false);
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
    static create_DC9(cf15) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC12(cf18, cf19, cf20) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D4(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC13(cf21) {
      let $dt = new D4(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.MultiSet.Empty);
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
    static create_DC15(cf23, cf24) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC16() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC14(cf22) {
      let $dt = new D5(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16";
      } else if (this.$tag === 2) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.Set.Empty, _dafny.Seq.of());
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
    static create_DC17(cf25) {
      let $dt = new D6(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf27) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(false);
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
    static create_DC21(cf29, cf30, cf31) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC22(cf32, cf33, cf34) {
      let $dt = new D8(1);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC20(cf28) {
      let $dt = new D8(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC23(cf35) {
      let $dt = new D8(3);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Set.Empty);
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
    static create_DC25(cf37, cf38, cf39) {
      let $dt = new D9(0);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC26(cf40, cf41) {
      let $dt = new D9(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC24(cf36) {
      let $dt = new D9(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC27(cf42) {
      let $dt = new D9(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get is_DC27() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf37 === other.cf37 && this.cf38 === other.cf38 && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f5 = _dafny.Set.Empty;
      this.f6 = _dafny.Set.Empty;
      this.f12 = [];
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.MultiSet.Empty;
      this._f3 = false;
      this._f7 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f8 = _dafny.MultiSet.Empty;
      this._f9 = false;
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.Seq.of();
      this._f13 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
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
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f16 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f16) {
      let _this = this;
      (_this)._f16 = f16;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.fromElements((_this).f16, !((_this).f16))).cardinality());
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return (function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)))).Elements) {
          let _432_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))), _432_v0)) {
            _coll9.push([_432_v0,_dafny.Seq.UnicodeFromString("tmqlcma")]);
          }
        }
        return _coll9;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_dafny.Seq.UnicodeFromString("qky"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_433_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      }))));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tv"), (((_this).f16) ? (_dafny.Seq.UnicodeFromString("usgh")) : (_dafny.Seq.UnicodeFromString("xjrsoedx"))));
    };
    m1(globalState) {
      let _this = this;
      let _434_v0;
      let _init7 = function (_435_i0) {
        return (_this).f16;
      };
      let _nw76 = Array((new BigNumber(4)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw76.length); _i0_7++) {
        _nw76[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _434_v0 = _nw76;
      let _index86 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length));
      (_434_v0)[_index86] = (_this).f16;
      let _436_v1;
      _436_v1 = _module.D1.create_DC5((_this).f16);
      let _index87 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length));
      (_434_v0)[_index87] = function (_source8) {
        if (_source8.is_DC5) {
          let _437___mcc_h0 = (_source8).cf9;
          let _438_cf9 = _437___mcc_h0;
          return (new BigNumber((_dafny.Set.fromElements((_this).f16)).length)).isEqualTo(new BigNumber(444));
        } else {
          let _439___mcc_h1 = (_source8).cf8;
          let _440_cf8 = _439___mcc_h1;
          return (_this).f16;
        }
      }(_436_v1);
      let _441_v2;
      _441_v2 = new BigNumber(96);
      let _442_v3;
      _442_v3 = _dafny.Set.fromElements(_441_v2, _441_v2);
      let _443_v4;
      _443_v4 = _dafny.Seq.of(_442_v3);
      let _444_v5;
      _444_v5 = _dafny.Seq.UnicodeFromString("oobutohe");
      let _445_v6;
      _445_v6 = _dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(_441_v2)));
      let _446_i1;
      _446_i1 = _dafny.ZERO;
      L5: {
        while ((new BigNumber((((_443_v4)[_module.__default.safeIndex(new BigNumber((_444_v5).length), new BigNumber((_443_v4).length))]).Union(_442_v3)).length)).isLessThanOrEqualTo(new BigNumber((_445_v6).length))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_446_i1)) {
              break L5;
            }
            _446_i1 = (_446_i1).plus(_dafny.ONE);
            let _447_v7;
            let _init8 = ((_448_v0) => function (_449_i2) {
              return _dafny.Seq.Concat(_dafny.Seq.of((_this).f16, (_this).f16), _dafny.Seq.of((_448_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_448_v0).length))]));
            })(_434_v0);
            let _nw77 = Array((new BigNumber(14)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw77.length); _i0_8++) {
              _nw77[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _447_v7 = _nw77;
            let _450_v8;
            _450_v8 = _dafny.Seq.of((_this).f16, false, (_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))], !((_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]), _module.__default.fm6((_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))], _441_v2, globalState));
            let _index88 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_447_v7).length));
            (_447_v7)[_index88] = _dafny.Seq.Concat(_450_v8, _450_v8);
            let _index89 = _module.__default.safeIndex(new BigNumber(22), new BigNumber((_447_v7).length));
            (_447_v7)[_index89] = _dafny.Seq.update(_450_v8, _module.__default.safeIndex(_441_v2, new BigNumber((_450_v8).length)), (_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]);
            let _451_v9;
            let _452_v10;
            let _453_v11;
            let _454_v12;
            let _out32;
            let _out33;
            let _out34;
            let _out35;
            let _outcollector8 = _module.__default.m0((_dafny.ZERO).minus((_441_v2).multipliedBy(_441_v2)), (_447_v7)[_module.__default.safeIndex(new BigNumber(22), new BigNumber((_447_v7).length))], globalState);
            _out32 = _outcollector8[0];
            _out33 = _outcollector8[1];
            _out34 = _outcollector8[2];
            _out35 = _outcollector8[3];
            _451_v9 = _out32;
            _452_v10 = _out33;
            _453_v11 = _out34;
            _454_v12 = _out35;
            let _455_v13;
            let _nw78 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
            _455_v13 = _nw78;
            let _456_v14;
            _456_v14 = new _dafny.CodePoint('f'.codePointAt(0));
            let _457_v15;
            _457_v15 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wluaxmjb"),_456_v14);
            let _index90 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_455_v13).length));
            (_455_v13)[_index90] = (_441_v2).multipliedBy(new BigNumber((_457_v15).length));
            let _index91 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_455_v13).length));
            (_455_v13)[_index91] = _module.__default.safeModuloInt(_441_v2, new BigNumber((_dafny.Seq.Concat(_445_v6, _445_v6)).length));
            let _458_v16;
            _458_v16 = _dafny.MultiSet.fromElements(new BigNumber((_444_v5).length), (_455_v13)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_455_v13).length))], new BigNumber(876));
            let _pat_let_tv19 = _458_v16;
            let _source9 = function (_pat_let10_0) {
              return function (_459_dt__update__tmp_h0) {
                return function (_pat_let11_0) {
                  return function (_460_dt__update_hcf5_h0) {
                    return _module.D0.create_DC2((_459_dt__update__tmp_h0).dtor_cf4, _460_dt__update_hcf5_h0, (_459_dt__update__tmp_h0).dtor_cf6);
                  }(_pat_let11_0);
                }(new BigNumber((_pat_let_tv19).cardinality()));
              }(_pat_let10_0);
            }(_module.__default.fm7(globalState));
            if (_source9.is_DC0) {
              let _461___mcc_h2 = (_source9).cf0;
              let _462___mcc_h3 = (_source9).cf1;
              let _463_cf1 = _462___mcc_h3;
              let _464_cf0 = _461___mcc_h2;
              let _index92 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length));
              (_434_v0)[_index92] = _451_v9;
              let _465_v17;
              _465_v17 = _dafny.Map.Empty.slice().updateUnsafe(_455_v13,!(_451_v9));
              _465_v17 = (_465_v17).update(_455_v13, _module.__default.fm6(_451_v9, new BigNumber((_442_v3).length), globalState));
              let _466_v18;
              _466_v18 = _module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements(_441_v2, new BigNumber(682))).length), _451_v9);
              _463_cf1 = !(!(!(!((_466_v18).dtor_cf1))));
              let _467_v19;
              let _468_v20;
              let _469_v21;
              let _470_v22;
              let _out36;
              let _out37;
              let _out38;
              let _out39;
              let _outcollector9 = _module.__default.m0(((_455_v13)[_module.__default.safeIndex(new BigNumber(462), new BigNumber((_455_v13).length))]).minus(_464_cf0), _dafny.Seq.of((_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]), globalState);
              _out36 = _outcollector9[0];
              _out37 = _outcollector9[1];
              _out38 = _outcollector9[2];
              _out39 = _outcollector9[3];
              _467_v19 = _out36;
              _468_v20 = _out37;
              _469_v21 = _out38;
              _470_v22 = _out39;
            } else if (_source9.is_DC1) {
              let _471___mcc_h4 = (_source9).cf2;
              let _472___mcc_h5 = (_source9).cf3;
              let _473_cf3 = _472___mcc_h5;
              let _474_cf2 = _471___mcc_h4;
              _441_v2 = new BigNumber(868);
              let _475_v23;
              let _init9 = ((_476_v14, _477_v0) => function (_478_i3) {
                return _dafny.Map.Empty.slice().updateUnsafe(_476_v14,(_477_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_477_v0).length))]);
              })(_456_v14, _434_v0);
              let _nw79 = Array((new BigNumber(7)).toNumber());
              for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw79.length); _i0_9++) {
                _nw79[_i0_9] = _init9(new BigNumber(_i0_9));
              }
              _475_v23 = _nw79;
              _475_v23 = _475_v23;
              let _479_v24;
              _479_v24 = _module.D0.create_DC3(_456_v14);
              _479_v24 = _479_v24;
              (globalState).f12 = _455_v13;
            } else if (_source9.is_DC2) {
              let _480___mcc_h6 = (_source9).cf4;
              let _481___mcc_h7 = (_source9).cf5;
              let _482___mcc_h8 = (_source9).cf6;
              let _483_cf6 = _482___mcc_h8;
              let _484_cf5 = _481___mcc_h7;
              let _485_cf4 = _480___mcc_h6;
              _444_v5 = _dafny.Seq.UnicodeFromString("qf");
              let _index93 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_455_v13).length));
              (_455_v13)[_index93] = _453_v11;
              let _486_v25;
              let _nw80 = Array((new BigNumber(2)).toNumber());
              _nw80[(_dafny.ZERO).toNumber()] = _456_v14;
              _nw80[(_dafny.ONE).toNumber()] = _456_v14;
              _486_v25 = _nw80;
              let _index94 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_486_v25).length));
              (_486_v25)[_index94] = _456_v14;
              let _index95 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_486_v25).length));
              (_486_v25)[_index95] = _456_v14;
              let _487_v26;
              _487_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_dafny.Set.fromElements(new BigNumber((_444_v5).length), _441_v2, _484_cf5));
              let _488_v27;
              _488_v27 = _module.D0.create_DC2(_485_cf4, new BigNumber((_450_v8).length), _451_v9);
              let _489_v28;
              _489_v28 = _module.D0.create_DC2((((_487_v26).contains(_451_v9)) ? ((_487_v26).get(_451_v9)) : (_442_v3)), (_488_v27).dtor_cf5, true);
              _489_v28 = _module.__default.fm7(globalState);
            } else {
              let _490___mcc_h9 = (_source9).cf7;
              let _491_cf7 = _490___mcc_h9;
              let _492_v29;
              let _493_v30;
              let _494_v31;
              let _495_v32;
              let _out40;
              let _out41;
              let _out42;
              let _out43;
              let _outcollector10 = _module.__default.m0(_453_v11, _dafny.Seq.of((_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]), globalState);
              _out40 = _outcollector10[0];
              _out41 = _outcollector10[1];
              _out42 = _outcollector10[2];
              _out43 = _outcollector10[3];
              _492_v29 = _out40;
              _493_v30 = _out41;
              _494_v31 = _out42;
              _495_v32 = _out43;
              _491_cf7 = (_444_v5)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_444_v5).length))];
              let _496_v33;
              _496_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-902),((!(_451_v9)) ? ((_this).f16) : ((_this).f16)));
              _496_v33 = (_496_v33).Merge(_496_v33);
              _434_v0 = _434_v0;
            }
          }
        }
      }
      let _497_v34;
      let _nw81 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _497_v34 = _nw81;
      let _index96 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length));
      (_497_v34)[_index96] = _441_v2;
      let _index97 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length));
      (_497_v34)[_index97] = (_this).fm1(_444_v5, _441_v2, (((_this).f16) ? (new BigNumber((_444_v5).length)) : (_441_v2)), globalState);
      let _498_v35;
      _498_v35 = new _dafny.CodePoint('f'.codePointAt(0));
      let _499_v36;
      _499_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f16,_498_v35);
      (globalState).f4 = (((_499_v36).contains((_this).f16)) ? ((_499_v36).get((_this).f16)) : (new _dafny.CodePoint('v'.codePointAt(0))));
      let _500_v37;
      _500_v37 = _dafny.Map.Empty.slice().updateUnsafe(((_497_v34)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length))]).isLessThan((_497_v34)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length))]),_441_v2);
      let _501_v38;
      _501_v38 = _dafny.Map.Empty.slice().updateUnsafe(_442_v3,new BigNumber(-717));
      let _502_v39;
      _502_v39 = _module.D0.create_DC2(_442_v3, (_497_v34)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length))], (_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]);
      let _503_v40;
      _503_v40 = _dafny.Map.Empty.slice().updateUnsafe(_441_v2,(_434_v0)[_module.__default.safeIndex(new BigNumber(848), new BigNumber((_434_v0).length))]);
      _500_v37 = (_500_v37).update((_this).f16, (((_this).f16) ? ((_497_v34)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length))]) : ((((_501_v38).contains((_502_v39).dtor_cf4)) ? ((_501_v38).get((_502_v39).dtor_cf4)) : (new BigNumber((_503_v40).length))))));
      let _index98 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length));
      (_497_v34)[_index98] = (_497_v34)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_497_v34).length))];
      return;
    }
    get f16() {
      let _this = this;
      return _this._f16;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f14 = false;
      this._f15 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f14, f15) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      return;
    }
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((((_this).f14) ? (_dafny.Seq.of(_dafny.Seq.of((_this).f15, (_this).f15, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f14)).length), (_this).f15))) : (_dafny.Seq.of(_dafny.Seq.of((_this).f15, (_this).f15)))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber(512)), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f14)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f15)).length))))).length));
    };
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return (function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(861)), function (_504_i0) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }))).Keys.Elements) {
          let _505_v0 = _compr_10;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(861)), function (_504_i0) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }))).contains(_505_v0)) {
            _coll10.push([_505_v0,_dafny.Seq.UnicodeFromString("evyrhbeq")]);
          }
        }
        return _coll10;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),_dafny.Seq.UnicodeFromString("ycuwefvef"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), function (_506_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }))));
    };
    m1(globalState) {
      let _this = this;
      let _507_v0;
      _507_v0 = _dafny.Seq.UnicodeFromString("qqptj");
      let _508_v1;
      _508_v1 = _dafny.MultiSet.fromElements(_507_v0, _dafny.Seq.Concat(_507_v0, _507_v0), _507_v0);
      let _509_v2;
      _509_v2 = _dafny.Seq.of(_507_v0, _507_v0, (_module.__default.fm3((_this).f15, globalState)).dtor_cf8);
      _508_v1 = ((_dafny.MultiSet.FromArray(_509_v2)).Intersect(_508_v1)).Union(_508_v1);
      _509_v2 = _509_v2;
      let _hi0 = (_this).f15;
      for (let _510_i0 = (_this).f15; _510_i0.isLessThan(_hi0); _510_i0 = _510_i0.plus(_dafny.ONE)) {
        let _511_v3;
        _511_v3 = _dafny.Set.fromElements((_this).f14, (_this).f14, (_this).f14);
        let _512_v4;
        _512_v4 = _dafny.MultiSet.fromElements(new BigNumber((_507_v0).length));
        let _513_v5;
        _513_v5 = new _dafny.CodePoint('r'.codePointAt(0));
        let _514_v6;
        _514_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_513_v5);
        let _515_v7;
        _515_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14);
        let _516_v8;
        _516_v8 = _dafny.Seq.of(true, (_this).f14);
        let _517_v9;
        let _nw82 = Array((new BigNumber(28)).toNumber());
        _nw82[(_dafny.ZERO).toNumber()] = (_510_i0).isLessThan(_510_i0);
        _nw82[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(3)).toNumber()] = true;
        _nw82[(new BigNumber(4)).toNumber()] = (_module.D0.create_DC2(_dafny.Set.fromElements((_this).f15), new BigNumber(180), true)).dtor_cf6;
        _nw82[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(6)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(9)).toNumber()] = (_511_v3).IsSubsetOf(_511_v3);
        _nw82[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(11)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(13)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(14)).toNumber()] = ((((_512_v4).contains(new BigNumber(488))) ? ((_512_v4).get(new BigNumber(488))) : (new BigNumber((_514_v6).length)))).isLessThan(_510_i0);
        _nw82[(new BigNumber(15)).toNumber()] = !(!((_this).f14)) || ((_this).f14);
        _nw82[(new BigNumber(16)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(17)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(18)).toNumber()] = (_510_i0).isLessThan(new BigNumber((_507_v0).length));
        _nw82[(new BigNumber(19)).toNumber()] = !(!(new BigNumber((_515_v7).length)).isEqualTo(new BigNumber(759)));
        _nw82[(new BigNumber(20)).toNumber()] = (_this).f14;
        _nw82[(new BigNumber(21)).toNumber()] = (_512_v4).IsProperSubsetOf((_512_v4).update(_510_i0, _module.__default.abs((_this).f15)));
        _nw82[(new BigNumber(22)).toNumber()] = !((((_this).f14) ? ((_this).f14) : ((_this).f14)));
        _nw82[(new BigNumber(23)).toNumber()] = !_dafny.areEqual(_516_v8, _516_v8);
        _nw82[(new BigNumber(24)).toNumber()] = !((_this).f14);
        _nw82[(new BigNumber(25)).toNumber()] = _dafny.Seq.contains(_507_v0, _513_v5);
        _nw82[(new BigNumber(26)).toNumber()] = _dafny.Seq.IsPrefixOf(_516_v8, _516_v8);
        _nw82[(new BigNumber(27)).toNumber()] = (_this).f14;
        _517_v9 = _nw82;
        let _index99 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
        (_517_v9)[_index99] = (_510_i0).isLessThanOrEqualTo((_this).f15);
        let _index100 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
        (_517_v9)[_index100] = (((_515_v7).contains((_510_i0).isEqualTo((_this).f15))) ? ((_515_v7).get((_510_i0).isEqualTo((_this).f15))) : (((_this).f14) || ((_this).f14)));
        let _518_v11;
        let _init10 = ((_519_i0, _520_v3) => function (_521_i1) {
          return function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_dafny.Set.fromElements((_this).f15, _519_i0, (_dafny.ZERO).minus(new BigNumber((_520_v3).length)), _519_i0, (_this).f15)).Elements) {
              let _522_v10 = _compr_11;
              if ((_dafny.Set.fromElements((_this).f15, _519_i0, (_dafny.ZERO).minus(new BigNumber((_520_v3).length)), _519_i0, (_this).f15)).contains(_522_v10)) {
                _coll11.push([_module.__default.safeModuloInt(_522_v10, new BigNumber((_dafny.MultiSet.fromElements((_this).f14, !((_this).f14), (_this).f14, (_this).f14, (_this).f14)).cardinality())),new BigNumber(-420)]);
              }
            }
            return _coll11;
          }();
        })(_510_i0, _511_v3);
        let _nw83 = Array((new BigNumber(22)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw83.length); _i0_10++) {
          _nw83[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _518_v11 = _nw83;
        let _init11 = ((_523_i0) => function (_524_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_523_i0,(_module.D0.create_DC0((_this).f15, (_this).f14)).dtor_cf0);
        })(_510_i0);
        let _nw84 = Array((new BigNumber(19)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw84.length); _i0_11++) {
          _nw84[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _518_v11 = _nw84;
        let _525_v12;
        _525_v12 = _dafny.Seq.of((_this).f15);
        let _526_v13;
        _526_v13 = _module.D0.create_DC0(new BigNumber((_525_v12).length), !((_this).f14));
        let _source10 = function (_pat_let12_0) {
          return function (_527_dt__update__tmp_h0) {
            return function (_pat_let13_0) {
              return function (_528_dt__update_hcf1_h0) {
                return _module.D0.create_DC0((_527_dt__update__tmp_h0).dtor_cf0, _528_dt__update_hcf1_h0);
              }(_pat_let13_0);
            }((_this).f14);
          }(_pat_let12_0);
        }(_526_v13);
        if (_source10.is_DC0) {
          let _529___mcc_h0 = (_source10).cf0;
          let _530___mcc_h1 = (_source10).cf1;
          let _531_cf1 = _530___mcc_h1;
          let _532_cf0 = _529___mcc_h0;
          _531_cf1 = (((_517_v9)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length))]) ? (!(_module.__default.fm4((_this).f14, _511_v3, (_this).f15, globalState))) : ((_this).f14));
          (globalState).f4 = _513_v5;
          let _533_v14;
          let _nw85 = new _module.C0();
          _nw85.__ctor((_511_v3).IsProperSubsetOf(_511_v3));
          _533_v14 = _nw85;
          let _index101 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
          (_517_v9)[_index101] = (_510_i0).isLessThanOrEqualTo(_510_i0);
        } else if (_source10.is_DC1) {
          let _534___mcc_h2 = (_source10).cf2;
          let _535___mcc_h3 = (_source10).cf3;
          let _536_cf3 = _535___mcc_h3;
          let _537_cf2 = _534___mcc_h2;
          let _538_v15;
          _538_v15 = new BigNumber(931);
          let _539_v16;
          let _nw86 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _539_v16 = _nw86;
          let _index102 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_539_v16).length));
          (_539_v16)[_index102] = _515_v7;
          let _index103 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_539_v16).length));
          let _index104 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
          let _rhs74 = (_525_v12)[_module.__default.safeIndex((_this).f15, new BigNumber((_525_v12).length))];
          let _rhs75 = (_515_v7).Merge((_515_v7).Merge(_515_v7));
          let _rhs76 = (_module.__default.safeModuloInt(_538_v15, _510_i0)).multipliedBy(_510_i0);
          let _rhs77 = (((_this).f14) ? ((_517_v9)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length))]) : (!((_this).f14)));
          let _lhs40 = _539_v16;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_539_v16).length));
          let _lhs42 = _517_v9;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
          _538_v15 = _rhs74;
          _lhs40[_lhs41] = _rhs75;
          _538_v15 = _rhs76;
          _lhs42[_lhs43] = _rhs77;
          let _540_v17;
          _540_v17 = _dafny.Map.Empty.slice().updateUnsafe(_510_i0,_dafny.Set.fromElements(false));
          let _541_v18;
          _541_v18 = _dafny.Map.Empty.slice().updateUnsafe(_540_v17,(_this).f15);
          _541_v18 = (_541_v18).update(_540_v17, new BigNumber(-424));
          let _542_v19;
          _542_v19 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_510_i0),(_this).f15);
          _538_v15 = _module.__default.safeDivisionInt(((_this).f15).multipliedBy((_this).f15), new BigNumber((_542_v19).length));
          _536_cf3 = !(_536_cf3);
        } else if (_source10.is_DC2) {
          let _543___mcc_h4 = (_source10).cf4;
          let _544___mcc_h5 = (_source10).cf5;
          let _545___mcc_h6 = (_source10).cf6;
          let _546_cf6 = _545___mcc_h6;
          let _547_cf5 = _544___mcc_h5;
          let _548_cf4 = _543___mcc_h4;
          _546_cf6 = !((_548_cf4).equals(_548_cf4));
          (globalState).f4 = _513_v5;
          let _index105 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
          (_517_v9)[_index105] = !((_this).f15).isEqualTo((_this).fm1(_507_v0, (_this).f15, new BigNumber(602), globalState));
          let _549_v20;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_517_v9)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length))]);
          _549_v20 = _nw87;
          let _550_v21;
          _550_v21 = _dafny.Seq.of(_549_v20, _549_v20, _549_v20);
          _546_cf6 = ((!((_this).f15).isEqualTo(new BigNumber((_550_v21).length))) ? ((_549_v20).f16) : (false));
        } else {
          let _551___mcc_h7 = (_source10).cf7;
          let _552_cf7 = _551___mcc_h7;
          _507_v0 = _507_v0;
          let _index106 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length));
          (_517_v9)[_index106] = (_this).f14;
          let _553_v22;
          _553_v22 = _dafny.MultiSet.fromElements((((_517_v9)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length))]) ? ((_517_v9)[_module.__default.safeIndex(new BigNumber(466), new BigNumber((_517_v9).length))]) : (true)));
          _553_v22 = _553_v22;
          _517_v9 = _517_v9;
        }
        let _554_v23;
        _554_v23 = _dafny.Set.fromElements(_511_v3);
        let _555_v24;
        _555_v24 = _module.D0.create_DC1((_554_v23).Difference(_554_v23), (_this).f14);
        _555_v24 = _555_v24;
      }
      let _556_v25;
      _556_v25 = new BigNumber(253);
      _556_v25 = _556_v25;
      let _557_v26;
      _557_v26 = _dafny.Set.fromElements(_556_v25, _556_v25);
      let _558_v27;
      _558_v27 = _dafny.Map.Empty.slice().updateUnsafe(_556_v25,_dafny.Set.fromElements((_this).f14, (_this).f14));
      let _559_v28;
      _559_v28 = _dafny.Set.fromElements((_this).f14, (_this).f14, (_this).f14);
      let _560_v29;
      _560_v29 = _module.D0.create_DC2(_557_v26, new BigNumber(((_dafny.Set.fromElements((_this).f14)).Union((((_558_v27).contains((_this).f15)) ? ((_558_v27).get((_this).f15)) : (_559_v28)))).length), (_this).f14);
      let _source11 = _560_v29;
      if (_source11.is_DC0) {
        let _561___mcc_h8 = (_source11).cf0;
        let _562___mcc_h9 = (_source11).cf1;
        let _563_cf1 = _562___mcc_h9;
        let _564_cf0 = _561___mcc_h8;
        let _565_v30;
        _565_v30 = _dafny.MultiSet.fromElements((_this).f15, _556_v25);
        let _566_v31;
        _566_v31 = _dafny.Map.Empty.slice().updateUnsafe((_565_v30).Difference(_565_v30),_556_v25);
        _566_v31 = (_566_v31).Merge((_566_v31).Merge(_566_v31));
        if ((_563_cf1) || (((_563_cf1) ? (_563_cf1) : (_563_cf1)))) {
          _556_v25 = (_this).f15;
          let _567_v32;
          _567_v32 = _dafny.MultiSet.fromElements(_563_cf1, false);
          _564_cf0 = new BigNumber(((_567_v32).Intersect(_567_v32)).cardinality());
          let _568_v33;
          _568_v33 = _dafny.Seq.of(_563_cf1);
          let _569_v34;
          _569_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_507_v0).length),(_this).f14);
          let _570_v35;
          let _nw88 = Array((new BigNumber(25)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = !(_563_cf1);
          _nw88[(_dafny.ONE).toNumber()] = _563_cf1;
          _nw88[(new BigNumber(2)).toNumber()] = _563_cf1;
          _nw88[(new BigNumber(3)).toNumber()] = (_this).f14;
          _nw88[(new BigNumber(4)).toNumber()] = ((_563_cf1) ? ((_568_v33)[_module.__default.safeIndex(_564_cf0, new BigNumber((_568_v33).length))]) : (false));
          _nw88[(new BigNumber(5)).toNumber()] = (_this).f14;
          _nw88[(new BigNumber(6)).toNumber()] = _563_cf1;
          _nw88[(new BigNumber(7)).toNumber()] = false;
          _nw88[(new BigNumber(8)).toNumber()] = !(!((_this).f15).isEqualTo((_dafny.ZERO).minus((_this).f15)));
          _nw88[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_507_v0, _dafny.Seq.UnicodeFromString("wmjh"));
          _nw88[(new BigNumber(10)).toNumber()] = false;
          _nw88[(new BigNumber(11)).toNumber()] = !(_563_cf1);
          _nw88[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw88[(new BigNumber(13)).toNumber()] = !(_563_cf1) || ((_this).f14);
          _nw88[(new BigNumber(14)).toNumber()] = (_565_v30).IsProperSubsetOf(_565_v30);
          _nw88[(new BigNumber(15)).toNumber()] = false;
          _nw88[(new BigNumber(16)).toNumber()] = (((_569_v34).contains(_564_cf0)) ? ((_569_v34).get(_564_cf0)) : (true));
          _nw88[(new BigNumber(17)).toNumber()] = (_567_v32).IsSubsetOf(_567_v32);
          _nw88[(new BigNumber(18)).toNumber()] = _563_cf1;
          _nw88[(new BigNumber(19)).toNumber()] = (_this).f14;
          _nw88[(new BigNumber(20)).toNumber()] = (_this).f14;
          _nw88[(new BigNumber(21)).toNumber()] = true;
          _nw88[(new BigNumber(22)).toNumber()] = (((_this).f14) ? (_563_cf1) : (true));
          _nw88[(new BigNumber(23)).toNumber()] = !(_563_cf1);
          _nw88[(new BigNumber(24)).toNumber()] = _563_cf1;
          _570_v35 = _nw88;
          let _571_v36;
          _571_v36 = new _dafny.CodePoint('s'.codePointAt(0));
          let _572_v37;
          _572_v37 = _dafny.Set.fromElements(_571_v36);
          let _nw89 = Array((new BigNumber(15)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _563_cf1;
          _nw89[(_dafny.ONE).toNumber()] = !(!((_this).f14));
          _nw89[(new BigNumber(2)).toNumber()] = (_this).f14;
          _nw89[(new BigNumber(3)).toNumber()] = _563_cf1;
          _nw89[(new BigNumber(4)).toNumber()] = false;
          _nw89[(new BigNumber(5)).toNumber()] = (_563_cf1) && ((_this).f14);
          _nw89[(new BigNumber(6)).toNumber()] = !((_this).f14);
          _nw89[(new BigNumber(7)).toNumber()] = (true) === ((_this).f14);
          _nw89[(new BigNumber(8)).toNumber()] = _563_cf1;
          _nw89[(new BigNumber(9)).toNumber()] = !(_563_cf1);
          _nw89[(new BigNumber(10)).toNumber()] = true;
          _nw89[(new BigNumber(11)).toNumber()] = (_this).f14;
          _nw89[(new BigNumber(12)).toNumber()] = _563_cf1;
          _nw89[(new BigNumber(13)).toNumber()] = (_556_v25).isLessThan(new BigNumber((_572_v37).length));
          _nw89[(new BigNumber(14)).toNumber()] = false;
          _570_v35 = _nw89;
          let _573_v38;
          let _nw90 = new _module.C0();
          _nw90.__ctor(_563_cf1);
          _573_v38 = _nw90;
          let _index107 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_570_v35).length));
          (_570_v35)[_index107] = _563_cf1;
          let _index108 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_570_v35).length));
          (_570_v35)[_index108] = !(_563_cf1);
        } else {
          let _574_v39;
          _574_v39 = _dafny.MultiSet.fromElements(_560_v29);
          let _575_v40;
          _575_v40 = _dafny.Seq.of(_module.D0.create_DC2(_557_v26, _564_cf0, _563_cf1), _560_v29);
          let _576_v41;
          _576_v41 = _dafny.Set.fromElements(_574_v39, _574_v39, (_574_v39).Union(_dafny.MultiSet.FromArray(_575_v40)), _dafny.MultiSet.fromElements(_560_v29), _dafny.MultiSet.fromElements(_560_v29));
          let _577_v42;
          _577_v42 = _dafny.Map.Empty.slice().updateUnsafe(_563_cf1,_563_cf1);
          let _rhs78 = _dafny.Set.fromElements((_574_v39).Difference(_574_v39));
          let _rhs79 = _507_v0;
          let _rhs80 = _module.__default.fm0((_577_v42).Merge(_577_v42), (_563_cf1) || ((_this).f14), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_578_i3) {
            return _module.D0.create_DC3(new _dafny.CodePoint('k'.codePointAt(0)));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-891)), function (_579_i4) {
            return _module.D0.create_DC3(new _dafny.CodePoint('m'.codePointAt(0)));
          })), globalState);
          let _rhs81 = _559_v28;
          let _rhs82 = (_this).f14;
          _576_v41 = _rhs78;
          _507_v0 = _rhs79;
          _564_cf0 = _rhs80;
          _559_v28 = _rhs81;
          _563_cf1 = _rhs82;
          _563_cf1 = !((_module.__default.fm0(_577_v42, !(!((_this).f14)), _563_cf1, globalState)).isEqualTo(_module.__default.fm0(_577_v42, (_this).f14, (_this).f14, globalState)));
          _564_cf0 = _module.__default.safeModuloInt((_this).f15, new BigNumber((_module.__default.fm8((_this).f14, (_this).f14, (_this).f14, _563_cf1, globalState)).length));
          _563_cf1 = (_this).f14;
          _564_cf0 = (_this).f15;
        }
        let _580_v43;
        let _nw91 = Array((new BigNumber(26)).toNumber()).fill(false);
        _580_v43 = _nw91;
        let _init12 = function (_581_i5) {
          return (_this).f14;
        };
        let _nw92 = Array((new BigNumber(7)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw92.length); _i0_12++) {
          _nw92[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _580_v43 = _nw92;
        let _582_v44;
        let _nw93 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _582_v44 = _nw93;
        let _index109 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_582_v44).length));
        (_582_v44)[_index109] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-152)), ((_583_v0, _584_v25) => function (_585_i6) {
          return (_583_v0)[_module.__default.safeIndex(_584_v25, new BigNumber((_583_v0).length))];
        })(_507_v0, _556_v25))).length);
        let _index110 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_582_v44).length));
        (_582_v44)[_index110] = (_dafny.ZERO).minus((((_this).f14) ? (_564_cf0) : (((_563_cf1) ? (new BigNumber((_dafny.Seq.of((_this).f14)).length)) : (_564_cf0)))));
      } else if (_source11.is_DC1) {
        let _586___mcc_h10 = (_source11).cf2;
        let _587___mcc_h11 = (_source11).cf3;
        let _588_cf3 = _587___mcc_h11;
        let _589_cf2 = _586___mcc_h10;
        let _590_v45;
        _590_v45 = new _dafny.CodePoint('u'.codePointAt(0));
        let _591_v46;
        let _nw94 = new _module.C0();
        _nw94.__ctor(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), function (_592_i7) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _dafny.Seq.update(_507_v0, _module.__default.safeIndex((_this).f15, new BigNumber((_507_v0).length)), _590_v45)));
        _591_v46 = _nw94;
        let _593_v47;
        let _nw95 = Array((new BigNumber(10)).toNumber()).fill(false);
        _593_v47 = _nw95;
        _593_v47 = _593_v47;
        let _index111 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_593_v47).length));
        (_593_v47)[_index111] = (_591_v46).f16;
        let _index112 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_593_v47).length));
        (_593_v47)[_index112] = (_591_v46).f16;
        let _594_v48;
        let _nw96 = new _module.C0();
        _nw96.__ctor((_591_v46).f16);
        _594_v48 = _nw96;
        let _595_v49;
        _595_v49 = _dafny.Seq.of(_594_v48);
        _595_v49 = _595_v49;
      } else if (_source11.is_DC2) {
        let _596___mcc_h12 = (_source11).cf4;
        let _597___mcc_h13 = (_source11).cf5;
        let _598___mcc_h14 = (_source11).cf6;
        let _599_cf6 = _598___mcc_h14;
        let _600_cf5 = _597___mcc_h13;
        let _601_cf4 = _596___mcc_h12;
        let _602_v50;
        _602_v50 = _dafny.Seq.of(_556_v25, new BigNumber((_557_v26).length), _600_cf5, _600_cf5);
        let _603_v51;
        _603_v51 = _dafny.Seq.of(false, _module.__default.fm4(!((_this).f14), _559_v28, (_this).f15, globalState));
        let _604_v52;
        _604_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_603_v51);
        let _605_v53;
        let _nw97 = Array((new BigNumber(15)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = (_this).f15;
        _nw97[(_dafny.ONE).toNumber()] = _600_cf5;
        _nw97[(new BigNumber(2)).toNumber()] = _600_cf5;
        _nw97[(new BigNumber(3)).toNumber()] = (((_this).f14) ? (_556_v25) : ((_this).f15));
        _nw97[(new BigNumber(4)).toNumber()] = _556_v25;
        _nw97[(new BigNumber(5)).toNumber()] = new BigNumber(918);
        _nw97[(new BigNumber(6)).toNumber()] = (_this).f15;
        _nw97[(new BigNumber(7)).toNumber()] = new BigNumber(718);
        _nw97[(new BigNumber(8)).toNumber()] = (_556_v25).minus(_600_cf5);
        _nw97[(new BigNumber(9)).toNumber()] = new BigNumber((_602_v50).length);
        _nw97[(new BigNumber(10)).toNumber()] = (_this).f15;
        _nw97[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_556_v25);
        _nw97[(new BigNumber(12)).toNumber()] = new BigNumber(67);
        _nw97[(new BigNumber(13)).toNumber()] = (new BigNumber((_604_v52).length)).minus(_556_v25);
        _nw97[(new BigNumber(14)).toNumber()] = _556_v25;
        _605_v53 = _nw97;
        let _index113 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length));
        (_605_v53)[_index113] = (_this).f15;
        let _606_v54;
        let _nw98 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _606_v54 = _nw98;
        let _index114 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length));
        (_606_v54)[_index114] = _507_v0;
        let _607_v55;
        _607_v55 = new _dafny.CodePoint('d'.codePointAt(0));
        let _608_v56;
        _608_v56 = _dafny.MultiSet.fromElements(_607_v55);
        let _609_v57;
        _609_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_556_v25);
        let _index115 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length));
        let _index116 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length));
        let _rhs83 = (_608_v56).IsDisjointFrom((_608_v56).update(_607_v55, _module.__default.abs(_600_cf5)));
        let _rhs84 = (_module.__default.fm9(_607_v55, (_this).f15, new BigNumber((_dafny.Seq.of(_599_cf6)).length), _609_v57, globalState)).Difference(_559_v28);
        let _rhs85 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_507_v0, _507_v0), _dafny.Seq.Concat(_509_v2, _509_v2))).length));
        let _rhs86 = _507_v0;
        let _lhs44 = _605_v53;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length));
        let _lhs46 = _606_v54;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length));
        _599_cf6 = _rhs83;
        _559_v28 = _rhs84;
        _lhs44[_lhs45] = _rhs85;
        _lhs46[_lhs47] = _rhs86;
        let _610_v58;
        _610_v58 = _dafny.Map.Empty.slice().updateUnsafe(_605_v53,(_this).f14);
        let _611_v59;
        let _nw99 = Array((new BigNumber(23)).toNumber());
        _nw99[(_dafny.ZERO).toNumber()] = _599_cf6;
        _nw99[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(3)).toNumber()] = (((_610_v58).contains(_605_v53)) ? ((_610_v58).get(_605_v53)) : (!(_599_cf6)));
        _nw99[(new BigNumber(4)).toNumber()] = !(true);
        _nw99[(new BigNumber(5)).toNumber()] = _599_cf6;
        _nw99[(new BigNumber(6)).toNumber()] = _module.__default.fm4(true, _559_v28, new BigNumber(((_606_v54)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length))]).length), globalState);
        _nw99[(new BigNumber(7)).toNumber()] = !(!(!_dafny.Seq.contains(_602_v50, _600_cf5)));
        _nw99[(new BigNumber(8)).toNumber()] = _599_cf6;
        _nw99[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(10)).toNumber()] = false;
        _nw99[(new BigNumber(11)).toNumber()] = _599_cf6;
        _nw99[(new BigNumber(12)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(13)).toNumber()] = (_600_cf5).isEqualTo((_605_v53)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length))]);
        _nw99[(new BigNumber(14)).toNumber()] = !_dafny.areEqual((_606_v54)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length))], (_606_v54)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length))]);
        _nw99[(new BigNumber(15)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(16)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(17)).toNumber()] = ((_599_cf6) ? (_599_cf6) : ((_this).f14));
        _nw99[(new BigNumber(18)).toNumber()] = !(_599_cf6) || ((_this).f14);
        _nw99[(new BigNumber(19)).toNumber()] = (_this).f14;
        _nw99[(new BigNumber(20)).toNumber()] = false;
        _nw99[(new BigNumber(21)).toNumber()] = _599_cf6;
        _nw99[(new BigNumber(22)).toNumber()] = _599_cf6;
        _611_v59 = _nw99;
        let _index117 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_611_v59).length));
        (_611_v59)[_index117] = (_this).f14;
        let _index118 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_611_v59).length));
        (_611_v59)[_index118] = _module.__default.fm4((_this).f14, _559_v28, (_this).f15, globalState);
        let _612_v60;
        _612_v60 = _module.D2.create_DC6(_dafny.Seq.of(_599_cf6));
        let _613_v61;
        _613_v61 = _dafny.MultiSet.fromElements(!((_this).f14));
        let _index119 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length));
        let _index120 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_611_v59).length));
        let _index121 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_611_v59).length));
        let _rhs87 = function (_pat_let14_0) {
          return function (_614_dt__update__tmp_h1) {
            return function (_pat_let15_0) {
              return function (_615_dt__update_hcf6_h0) {
                return function (_pat_let16_0) {
                  return function (_616_dt__update_hcf5_h0) {
                    return _module.D0.create_DC2((_614_dt__update__tmp_h1).dtor_cf4, _616_dt__update_hcf5_h0, _615_dt__update_hcf6_h0);
                  }(_pat_let16_0);
                }(new BigNumber(360));
              }(_pat_let15_0);
            }(!((_this).f14) || ((_this).f14));
          }(_pat_let14_0);
        }(_560_v29);
        let _rhs88 = (_this).f15;
        let _rhs89 = false;
        let _rhs90 = true;
        let _rhs91 = (((_dafny.MultiSet.FromArray((_612_v60).dtor_cf10)).IsProperSubsetOf(_613_v61)) ? ((_613_v61).IsProperSubsetOf(_613_v61)) : (!(false)));
        let _lhs48 = _605_v53;
        let _lhs49 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_605_v53).length));
        let _lhs50 = _611_v59;
        let _lhs51 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_611_v59).length));
        let _lhs52 = _611_v59;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_611_v59).length));
        _560_v29 = _rhs87;
        _lhs48[_lhs49] = _rhs88;
        _599_cf6 = _rhs89;
        _lhs50[_lhs51] = _rhs90;
        _lhs52[_lhs53] = _rhs91;
        let _617_v62;
        let _nw100 = Array((new BigNumber(16)).toNumber());
        _617_v62 = _nw100;
        let _618_v63;
        let _nw101 = new _module.C0();
        _nw101.__ctor(!(true));
        _618_v63 = _nw101;
        let _index122 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_617_v62).length));
        (_617_v62)[_index122] = _618_v63;
        let _index123 = _module.__default.safeIndex(new BigNumber(851), new BigNumber((_617_v62).length));
        (_617_v62)[_index123] = _618_v63;
        let _619_v64;
        _619_v64 = _dafny.Map.Empty.slice().updateUnsafe((_611_v59)[_module.__default.safeIndex(new BigNumber(274), new BigNumber((_611_v59).length))],(_606_v54)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length))]);
        _600_cf5 = new BigNumber((((_619_v64).update(_599_cf6, (_606_v54)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_606_v54).length))])).Merge((_619_v64).Merge(_619_v64))).length);
      } else {
        let _620___mcc_h15 = (_source11).cf7;
        let _621_cf7 = _620___mcc_h15;
        let _622_v65;
        _622_v65 = _module.D2.create_DC7(_module.__default.fm6((_this).f14, _556_v25, globalState), (_this).f14, !(new BigNumber((_dafny.Seq.UnicodeFromString("lawxj")).length)).isEqualTo(new BigNumber(819)));
        _622_v65 = _622_v65;
        if (!(_module.__default.safeDivisionInt(_556_v25, _556_v25)).isEqualTo((_this).f15)) {
          let _623_v66;
          let _nw102 = new _module.C0();
          _nw102.__ctor((_this).f14);
          _623_v66 = _nw102;
          let _624_v67;
          _624_v67 = _dafny.Seq.of((_this).f14, !(!((_this).f14) || ((_623_v66).f16)), (_623_v66).f16, !((_dafny.ZERO).minus((_this).f15)).isEqualTo((_this).f15));
          let _625_v68;
          let _nw103 = Array((new BigNumber(20)).toNumber()).fill(false);
          _625_v68 = _nw103;
          let _index124 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_625_v68).length));
          (_625_v68)[_index124] = (_this).f14;
          let _index125 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_625_v68).length));
          let _rhs92 = _624_v67;
          let _rhs93 = ((_this).f15).isLessThanOrEqualTo(_556_v25);
          let _rhs94 = _623_v66;
          let _rhs95 = _625_v68;
          let _lhs54 = _625_v68;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(777), new BigNumber((_625_v68).length));
          _624_v67 = _rhs92;
          _lhs54[_lhs55] = _rhs93;
          _623_v66 = _rhs94;
          _625_v68 = _rhs95;
          let _626_v69;
          let _nw104 = new _module.C0();
          _nw104.__ctor((_625_v68)[_module.__default.safeIndex(new BigNumber(777), new BigNumber((_625_v68).length))]);
          _626_v69 = _nw104;
          let _627_v70;
          let _nw105 = new _module.C0();
          _nw105.__ctor(_dafny.Seq.IsPrefixOf((_626_v69).fm5(false, _556_v25, new BigNumber(43), (_this).f14, globalState), _507_v0));
          _627_v70 = _nw105;
          let _628_v71;
          _628_v71 = _module.D0.create_DC0((_this).f15, (_626_v69).f16);
          let _629_v72;
          let _nw106 = new _module.C0();
          _nw106.__ctor(((_628_v71).dtor_cf1) || (_module.__default.fm6((_627_v70).f16, _556_v25, globalState)));
          _629_v72 = _nw106;
        } else {
          _556_v25 = (((_this).f14) ? ((_this).f15) : (_module.__default.safeModuloInt(_556_v25, _556_v25)));
          let _630_v73;
          _630_v73 = _dafny.Map.Empty.slice().updateUnsafe(_556_v25,_556_v25);
          _630_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f15);
          let _631_v74;
          let _nw107 = new _module.C0();
          _nw107.__ctor(false);
          _631_v74 = _nw107;
          let _632_v75;
          _632_v75 = _dafny.Seq.of((_this).f14, _module.__default.fm4((_631_v74).f16, _559_v28, new BigNumber(231), globalState));
          let _633_v76;
          let _634_v77;
          let _635_v78;
          let _636_v79;
          let _out44;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector11 = _module.__default.m0(_556_v25, _632_v75, globalState);
          _out44 = _outcollector11[0];
          _out45 = _outcollector11[1];
          _out46 = _outcollector11[2];
          _out47 = _outcollector11[3];
          _633_v76 = _out44;
          _634_v77 = _out45;
          _635_v78 = _out46;
          _636_v79 = _out47;
          let _637_v80;
          _637_v80 = _module.D0.create_DC3(_621_cf7);
          _637_v80 = _637_v80;
        }
        let _638_v81;
        _638_v81 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f15);
        let _639_v82;
        _639_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f14);
        let _640_v83;
        let _nw108 = new _module.C0();
        _nw108.__ctor((_this).f14);
        _640_v83 = _nw108;
        let _641_v84;
        _641_v84 = _dafny.Map.Empty.slice().updateUnsafe(_556_v25,_640_v83);
        let _642_v85;
        let _nw109 = Array((new BigNumber(16)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = !((_this).f14) || ((_this).f14);
        _nw109[(_dafny.ONE).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(2)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(3)).toNumber()] = !((_this).f14) || (_module.__default.fm4((_this).f14, _559_v28, new BigNumber((_638_v81).length), globalState));
        _nw109[(new BigNumber(4)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(5)).toNumber()] = !((_this).f14);
        _nw109[(new BigNumber(6)).toNumber()] = (((_639_v82).contains((_this).f14)) ? ((_639_v82).get((_this).f14)) : ((_this).f14));
        _nw109[(new BigNumber(7)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(8)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(9)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(11)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(12)).toNumber()] = (new BigNumber((_641_v84).length)).isLessThanOrEqualTo(new BigNumber(-202));
        _nw109[(new BigNumber(13)).toNumber()] = (_this).f14;
        _nw109[(new BigNumber(14)).toNumber()] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(438)), ((_643_cf7) => function (_644_i8) {
          return _643_cf7;
        })(_621_cf7))).length)).isLessThan((_dafny.ZERO).minus((_this).f15));
        _nw109[(new BigNumber(15)).toNumber()] = !((_640_v83).f16);
        _642_v85 = _nw109;
        let _index126 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_642_v85).length));
        (_642_v85)[_index126] = (_this).f14;
        let _index127 = _module.__default.safeIndex(new BigNumber(51), new BigNumber((_642_v85).length));
        (_642_v85)[_index127] = (_this).f14;
        let _645_v86;
        _645_v86 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(true, _559_v28, new BigNumber(-291), globalState),(((_642_v85)[_module.__default.safeIndex(new BigNumber(51), new BigNumber((_642_v85).length))]) ? (_dafny.Seq.UnicodeFromString("kyo")) : (_507_v0)));
        _645_v86 = (_645_v86).update((_this).f14, _dafny.Seq.Concat(_507_v0, _507_v0));
      }
      let _646_v87;
      _646_v87 = false;
      let _647_v88;
      let _init13 = ((_648_v87) => function (_649_i9) {
        return _648_v87;
      })(_646_v87);
      let _nw110 = Array((new BigNumber(23)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw110.length); _i0_13++) {
        _nw110[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _647_v88 = _nw110;
      let _650_v89;
      _650_v89 = _dafny.MultiSet.fromElements(_647_v88, _647_v88);
      let _651_v90;
      _651_v90 = _dafny.MultiSet.fromElements((_this).fm1(_507_v0, (_this).f15, _556_v25, globalState));
      _646_v87 = ((_650_v89).update(_647_v88, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_651_v90).cardinality()))))).IsSubsetOf(_650_v89);
      return;
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
