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
    static fm4(p0, p1, globalState) {
      if ((new BigNumber((_dafny.Seq.UnicodeFromString("piggfikm")).length)).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, false, !(false), false)).length), new BigNumber(-825), (_dafny.ZERO).minus(new BigNumber(-343)), new BigNumber(185))).length))) {
        return _dafny.Seq.UnicodeFromString("rehg");
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hm"), _dafny.Seq.UnicodeFromString("clvp"));
      }
    };
    static fm7(p0, p1, p2, globalState) {
      if (_dafny.areEqual(_dafny.Seq.UnicodeFromString("bxb"), _dafny.Seq.UnicodeFromString("p"))) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }
    };
    static fm9(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kivrkb"), _dafny.Seq.UnicodeFromString("fqfafs"));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _module.D5.create_DC11(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(231),true)).Merge(function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-259),!(true))).length), new BigNumber(54), new BigNumber(207), new BigNumber(841), new BigNumber((_dafny.Set.fromElements(new BigNumber(774))).length))).Elements) {
    let _0_v0 = _compr_0;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-259),!(true))).length), new BigNumber(54), new BigNumber(207), new BigNumber(841), new BigNumber((_dafny.Set.fromElements(new BigNumber(774))).length)), _0_v0)) {
      _coll0.push([(_0_v0).minus(new BigNumber(998)),!(!(!(!(false))))]);
    }
  }
  return _coll0;
}())).length), (_module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ekmhsfs"))).dtor_cf10);
    };
    static fm12(p0, p1, globalState) {
      return (_module.D16.create_DC41(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("tfnuqeu")))).dtor_cf72;
    };
    static fm13(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(240))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("bjdopi")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(false)).length))));
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Set.fromElements(new BigNumber(100))).length))).length)), new BigNumber(-881), new BigNumber(473), new BigNumber((_dafny.Set.fromElements(true)).length));
    };
    static fm15(globalState) {
      return _module.D2.create_DC4(true);
    };
    static fm16(p0, p1, p2, globalState) {
      return _dafny.Seq.of(((true) ? (_module.D4.create_DC7(_dafny.Seq.UnicodeFromString("gwsfpeakx"))) : (_module.D4.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), function (_1_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})))));
    };
    static fm17(p0, p1, p2, globalState) {
      return _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ysmgae"));
    };
    static fm18(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yriwev")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(908)),!(true))).length), new BigNumber(640))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(182), new BigNumber((_dafny.Seq.of(true, !(!(true)), false, true)).length)));
    };
    static fm19(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(false, true)).Difference(_dafny.MultiSet.fromElements(true))).Difference(_dafny.MultiSet.fromElements(true, !(false), !(true), false, false));
    };
    static fm20(globalState) {
      return function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(589), new BigNumber(859))) {
          let _2_v0 = _compr_1;
          if (((new BigNumber(589)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(859)))) {
            _coll1.push([(_2_v0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("nqxdyp")).length)),true]);
          }
        }
        return _coll1;
      }();
    };
    static fm21(p0, globalState) {
      return function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(-882))).length),new BigNumber((_dafny.Seq.of(true)).length)))).Elements) {
          let _3_v0 = _compr_2;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(-882))).length),new BigNumber((_dafny.Seq.of(true)).length)))).contains(_3_v0)) {
            _coll2.push([_3_v0,(_module.D3.create_DC6(!(true), new BigNumber(563), new BigNumber(882))).dtor_cf8]);
          }
        }
        return _coll2;
      }();
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge((_module.D10.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(true,true))).dtor_cf42)).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false)));
    };
    static fm23(p0, p1, p2, globalState) {
      if (!(!(!(!(true))))) {
        return (function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new BigNumber(152))).length))).length), new BigNumber(114), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_4_i0) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length))).Elements) {
            let _5_v0 = _compr_3;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new BigNumber(152))).length))).length), new BigNumber(114), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_4_i0) {
              return new _dafny.CodePoint('i'.codePointAt(0));
            })).length))).contains(_5_v0)) {
              _coll3.add((_5_v0).minus(new BigNumber((_dafny.Set.fromElements(!(true))).length)));
            }
          }
          return _coll3;
        }()).Intersect(function () {
          let _coll4 = new _dafny.Set();
          for (const _compr_4 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("chjhmkgm")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber(875))).Elements) {
            let _6_v1 = _compr_4;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("chjhmkgm")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber(875)), _6_v1)) {
              _coll4.add((_6_v1).plus(new BigNumber(-795)));
            }
          }
          return _coll4;
        }());
      } else {
        return _dafny.Set.fromElements(new BigNumber((function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(887), new BigNumber(374))) {
            let _7_v2 = _compr_5;
            if (((new BigNumber(887)).isLessThanOrEqualTo(_7_v2)) && ((_7_v2).isLessThan(new BigNumber(374)))) {
              _coll5.add((_7_v2).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))).cardinality())));
            }
          }
          return _coll5;
        }()).length));
      }
    };
    static fm24(globalState) {
      return _module.D7.create_DC16(_module.D7.create_DC15());
    };
    static fm25(p0, p1, p2, globalState) {
      return _module.D12.create_DC30(new BigNumber(378), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)))).cardinality()));
    };
    static fm26(p0, p1, globalState) {
      return _dafny.Set.fromElements(((!(true)) ? (true) : (false)));
    };
    static fm27(globalState) {
      let _source0 = _module.D8.create_DC18(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),new BigNumber(798))).length)),new BigNumber(95))).length),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(702))).length))).cardinality()))), true, new BigNumber(138));
      if (_source0.is_DC18) {
        let _8___mcc_h0 = (_source0).cf28;
        let _9___mcc_h1 = (_source0).cf29;
        let _10___mcc_h2 = (_source0).cf30;
        let _11_cf30 = _10___mcc_h2;
        let _12_cf29 = _9___mcc_h1;
        let _13_cf28 = _8___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_12_cf29), _dafny.Seq.of(_12_cf29, _12_cf29, _12_cf29));
      } else if (_source0.is_DC19) {
        let _14___mcc_h3 = (_source0).cf31;
        let _15_cf31 = _14___mcc_h3;
        return _dafny.Seq.Concat(_dafny.Seq.of(_15_cf31), _dafny.Seq.of(_15_cf31, _15_cf31));
      } else {
        let _16___mcc_h4 = (_source0).cf27;
        let _17_cf27 = _16___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.of(false, true, !(true)), _dafny.Seq.of(true));
      }
    };
    static m4(p0, p1, p2, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _18_v0;
      _18_v0 = _dafny.MultiSet.fromElements(p0);
      let _19_v1;
      _19_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _20_v2;
      let _nw0 = new _module.C3();
      _nw0.__ctor((((_19_v1).contains(p1)) ? ((_19_v1).get(p1)) : (p1)), p1);
      _20_v2 = _nw0;
      let _21_v3;
      _21_v3 = _dafny.Set.fromElements(_20_v2);
      let _22_v4;
      _22_v4 = _dafny.Seq.of(_20_v2.f8);
      let _23_v5;
      _23_v5 = _dafny.Seq.of(new BigNumber((_22_v4).length));
      let _24_v6;
      _24_v6 = _module.D0.create_DC1(p1, _dafny.Seq.of(p1, (_22_v4)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_22_v4).length))]), false);
      let _25_v7;
      let _nw1 = new _module.C3();
      _nw1.__ctor(p1, p1);
      _25_v7 = _nw1;
      let _26_v8;
      _26_v8 = _dafny.Map.Empty.slice().updateUnsafe(_25_v7,p0);
      let _27_v9;
      _27_v9 = _module.D9.create_DC22(p0, new BigNumber((_23_v5).length), p1, (_dafny.ZERO).minus((_20_v2).fm2(_24_v6, globalState)), _26_v8);
      let _28_v10;
      _28_v10 = _dafny.MultiSet.fromElements((((_18_v0).contains(new BigNumber((_21_v3).length))) ? ((_18_v0).get(new BigNumber((_21_v3).length))) : (new BigNumber(-700))), (_23_v5)[_module.__default.safeIndex((_27_v9).dtor_cf35, new BigNumber((_23_v5).length))]);
      let _29_v11;
      let _nw2 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
      _29_v11 = _nw2;
      let _30_v12;
      _30_v12 = _dafny.Set.fromElements(p0);
      let _31_v13;
      _31_v13 = _dafny.Seq.of(_30_v12, _30_v12, _30_v12, _30_v12);
      let _index0 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_29_v11).length));
      (_29_v11)[_index0] = (_31_v13)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_31_v13).length))];
      let _32_v14;
      _32_v14 = _dafny.Map.Empty.slice().updateUnsafe((_25_v7).f9,p0);
      let _33_v15;
      _33_v15 = _module.D5.create_DC11(new BigNumber((_32_v14).length), p2);
      let _34_v16;
      _34_v16 = new _dafny.CodePoint('m'.codePointAt(0));
      let _pat_let_tv0 = p0;
      let _pat_let_tv1 = _34_v16;
      let _35_v17;
      let _nw3 = Array((new BigNumber(20)).toNumber());
      _nw3[(_dafny.ZERO).toNumber()] = p0;
      _nw3[(_dafny.ONE).toNumber()] = p0;
      _nw3[(new BigNumber(2)).toNumber()] = (new BigNumber(772)).plus(p0);
      _nw3[(new BigNumber(3)).toNumber()] = p0;
      _nw3[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("en")).length);
      _nw3[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((p0).minus(p0));
      _nw3[(new BigNumber(6)).toNumber()] = p0;
      _nw3[(new BigNumber(7)).toNumber()] = p0;
      _nw3[(new BigNumber(8)).toNumber()] = (p0).plus(p0);
      _nw3[(new BigNumber(9)).toNumber()] = p0;
      _nw3[(new BigNumber(10)).toNumber()] = new BigNumber(141);
      _nw3[(new BigNumber(11)).toNumber()] = (_20_v2).fm2(_24_v6, globalState);
      _nw3[(new BigNumber(12)).toNumber()] = (_23_v5)[_module.__default.safeIndex((function (_pat_let0_0) {
        return function (_36_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_39_dt__update_hcf19_h0) {
              return _module.D5.create_DC11((_36_dt__update__tmp_h0).dtor_cf18, _39_dt__update_hcf19_h0);
            }(_pat_let1_0);
          }(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-596)), function (_37_i0) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _module.__default.safeIndex(_pat_let_tv0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-596)), function (_38_i0) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          })).length)), _pat_let_tv1));
        }(_pat_let0_0);
      }(_33_v15)).dtor_cf18, new BigNumber((_23_v5).length))];
      _nw3[(new BigNumber(13)).toNumber()] = p0;
      _nw3[(new BigNumber(14)).toNumber()] = p0;
      _nw3[(new BigNumber(15)).toNumber()] = p0;
      _nw3[(new BigNumber(16)).toNumber()] = (p0).plus(p0);
      _nw3[(new BigNumber(17)).toNumber()] = p0;
      _nw3[(new BigNumber(18)).toNumber()] = p0;
      _nw3[(new BigNumber(19)).toNumber()] = (p0).multipliedBy(p0);
      _35_v17 = _nw3;
      let _index1 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_29_v11).length));
      let _rhs0 = _dafny.MultiSet.fromElements(p0);
      let _rhs1 = (_30_v12).Intersect(_30_v12);
      let _rhs2 = _35_v17;
      let _rhs3 = _module.__default.safeModuloInt((_23_v5)[_module.__default.safeIndex(p0, new BigNumber((_23_v5).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_25_v7).f9), _22_v4)).length));
      let _lhs0 = _29_v11;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_29_v11).length));
      _18_v0 = _rhs0;
      _lhs0[_lhs1] = _rhs1;
      _35_v17 = _rhs2;
      r0 = _rhs3;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_35_v17).length))) {
        let _40_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_40_i1)) && ((_40_i1).isLessThan(new BigNumber((_35_v17).length))))) {
          (_35_v17)[(_40_i1)] = (_40_i1).multipliedBy(p0);
        }
      }
      let _pat_let_tv2 = _25_v7;
      (_20_v2).f8 = function (_source1) {
        if (_source1.is_DC0) {
          return (_pat_let_tv2).f9;
        } else {
          let _41___mcc_h0 = (_source1).cf0;
          let _42___mcc_h1 = (_source1).cf1;
          let _43___mcc_h2 = (_source1).cf2;
          let _44_cf2 = _43___mcc_h2;
          let _45_cf1 = _42___mcc_h1;
          let _46_cf0 = _41___mcc_h0;
          return !(false);
        }
      }(_24_v6);
      let _hi0 = (p0).multipliedBy(p0);
      for (let _47_i2 = p0; _47_i2.isLessThan(_hi0); _47_i2 = _47_i2.plus(_dafny.ONE)) {
        let _48_v18;
        let _49_v19;
        let _50_v20;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_25_v7).m1(_dafny.Seq.Concat(p2, p2), (_dafny.MultiSet.fromElements(_20_v2.f8, p1)).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of((_20_v2).fm0(globalState), _20_v2.f8))), globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _48_v18 = _out0;
        _49_v19 = _out1;
        _50_v20 = _out2;
        let _51_v21;
        _51_v21 = _dafny.Map.Empty.slice().updateUnsafe(_48_v18,_20_v2.f8);
        r1 = (p0).isLessThan(new BigNumber((_51_v21).length));
        let _52_v22;
        _52_v22 = _module.D4.create_DC7(_module.__default.fm9((_25_v7).f9, _20_v2.f8, globalState));
        let _53_v23;
        _53_v23 = _dafny.Seq.of(_52_v22, _52_v22, _52_v22, _52_v22, _52_v22);
        let _54_v24;
        let _nw4 = new _module.C1();
        _nw4.__ctor(_34_v16, _dafny.Set.fromElements(_53_v23, _53_v23), _dafny.Seq.contains(_23_v5, _47_i2), _20_v2.f8);
        _54_v24 = _nw4;
        let _55_v25;
        let _nw5 = new _module.C3();
        _nw5.__ctor(_50_v20, _20_v2.f8);
        _55_v25 = _nw5;
        let _nw6 = new _module.C3();
        _nw6.__ctor(!(_20_v2.f8) || (_50_v20), (((_25_v7).f9) ? (_50_v20) : (_20_v2.f8)));
        _55_v25 = _nw6;
      }
      let _56_v26;
      let _init0 = ((_57_v2, _58_v7) => function (_59_i4) {
        return _dafny.Set.fromElements(_57_v2.f8, !(_57_v2.f8), (_58_v7).f9);
      })(_20_v2, _25_v7);
      let _nw7 = Array((new BigNumber(2)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
        _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _56_v26 = _nw7;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_56_v26).length))) {
        let _60_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_60_i3)) && ((_60_i3).isLessThan(new BigNumber((_56_v26).length))))) {
          (_56_v26)[(_60_i3)] = _dafny.Set.fromElements((_25_v7).f9, _20_v2.f8);
        }
      }
      let _hi1 = (_dafny.ZERO).minus((p0).plus(p0));
      for (let _61_i5 = (p0).multipliedBy(p0); _61_i5.isLessThan(_hi1); _61_i5 = _61_i5.plus(_dafny.ONE)) {
        r0 = (new BigNumber(672)).multipliedBy(new BigNumber((p2).length));
        let _62_v27;
        _62_v27 = _dafny.MultiSet.fromElements(_20_v2.f8, true);
        let _63_v28;
        _63_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_62_v27);
        r0 = _module.__default.safeModuloInt(new BigNumber((_63_v28).length), (_20_v2).fm2(_24_v6, globalState));
        let _64_v29;
        _64_v29 = _module.D5.create_DC10(_35_v17);
        let _65_v30;
        _65_v30 = _dafny.MultiSet.fromElements((_64_v29).dtor_cf17);
        _65_v30 = (_65_v30).Union((_dafny.MultiSet.fromElements(_35_v17, _35_v17)).Intersect(_65_v30));
        let _66_v31;
        let _init1 = ((_67_v2) => function (_68_i6) {
          return !(_67_v2.f8);
        })(_20_v2);
        let _nw8 = Array((new BigNumber(17)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw8.length); _i0_1++) {
          _nw8[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _66_v31 = _nw8;
        let _index2 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_66_v31).length));
        (_66_v31)[_index2] = (_25_v7).f9;
        let _index3 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_66_v31).length));
        (_66_v31)[_index3] = _20_v2.f8;
      }
      r0 = p0;
      let _69_v32;
      _69_v32 = _module.D12.create_DC30(p0, p0);
      let _70_v33;
      _70_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.D12.create_DC32(_69_v32),p2);
      let _71_v34;
      _71_v34 = _dafny.Set.fromElements((((_70_v33).contains(_module.D12.create_DC32(_69_v32))) ? ((_70_v33).get(_module.D12.create_DC32(_69_v32))) : (p2)));
      r1 = ((_20_v2.f8) ? (_20_v2.f8) : ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("kl"))).IsDisjointFrom(_71_v34)));
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _72_v0;
      _72_v0 = _dafny.Seq.UnicodeFromString("xl");
      let _73_v1;
      let _init2 = function (_74_i0) {
        return true;
      };
      let _nw9 = Array((new BigNumber(18)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw9.length); _i0_2++) {
        _nw9[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _73_v1 = _nw9;
      let _75_globalState;
      let _nw10 = new _module.GlobalState();
      _nw10.__ctor(new _dafny.CodePoint('i'.codePointAt(0)), new BigNumber(114), new BigNumber(-993), new BigNumber(495), _72_v0, new BigNumber(-616), _73_v1, new BigNumber(-825));
      _75_globalState = _nw10;
      let _76_v2;
      _76_v2 = false;
      let _77_v3;
      _77_v3 = _dafny.Seq.of(_76_v2, _76_v2, _76_v2, _76_v2);
      let _pat_let_tv3 = _77_v3;
      let _source2 = function (_source3) {
        if (_source3.is_DC0) {
          return _module.D0.create_DC0();
        } else {
          let _78___mcc_h3 = (_source3).cf0;
          let _79___mcc_h4 = (_source3).cf1;
          let _80___mcc_h5 = (_source3).cf2;
          let _81_cf2 = _80___mcc_h5;
          let _82_cf1 = _79___mcc_h4;
          let _83_cf0 = _78___mcc_h3;
          return _module.D0.create_DC0();
        }
      }(function (_pat_let2_0) {
        return function (_84_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_85_dt__update_hcf1_h0) {
              return _module.D0.create_DC1((_84_dt__update__tmp_h0).dtor_cf0, _85_dt__update_hcf1_h0, (_84_dt__update__tmp_h0).dtor_cf2);
            }(_pat_let3_0);
          }(_pat_let_tv3);
        }(_pat_let2_0);
      }(_module.D0.create_DC1(_76_v2, _77_v3, _76_v2)));
      if (_source2.is_DC0) {
        if (_76_v2) {
          let _86_v4;
          let _nw11 = new _module.C3();
          _nw11.__ctor(_76_v2, _76_v2);
          _86_v4 = _nw11;
          let _87_v5;
          _87_v5 = new _dafny.CodePoint('p'.codePointAt(0));
          let _rhs4 = new _dafny.CodePoint('c'.codePointAt(0));
          let _rhs5 = _76_v2;
          let _rhs6 = _77_v3;
          let _rhs7 = _87_v5;
          _87_v5 = _rhs4;
          _76_v2 = _rhs5;
          _77_v3 = _rhs6;
          _87_v5 = _rhs7;
          let _88_v6;
          let _nw12 = Array((new BigNumber(6)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_72_v0, _72_v0);
          _nw12[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("fblhxtr");
          _nw12[(new BigNumber(2)).toNumber()] = _72_v0;
          _nw12[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_72_v0, _72_v0);
          _nw12[(new BigNumber(4)).toNumber()] = _72_v0;
          _nw12[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("whybwpehq");
          _88_v6 = _nw12;
          let _89_v7;
          _89_v7 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,_72_v0);
          let _90_v8;
          _90_v8 = new BigNumber(873);
          let _91_v9;
          _91_v9 = _dafny.Seq.of(_90_v8);
          let _92_v10;
          _92_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_91_v9).length),_72_v0);
          let _index4 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_88_v6).length));
          (_88_v6)[_index4] = (((_89_v7).contains(_76_v2)) ? ((_89_v7).get(_76_v2)) : ((((_92_v10).contains(_90_v8)) ? ((_92_v10).get(_90_v8)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(621)), ((_93_v5) => function (_94_i1) {
            return _93_v5;
          })(_87_v5))))));
          let _index5 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_88_v6).length));
          (_88_v6)[_index5] = _72_v0;
          let _95_v11;
          _95_v11 = _module.D9.create_DC21(_90_v8, _90_v8);
          let _96_v12;
          _96_v12 = _dafny.Map.Empty.slice().updateUnsafe(_95_v11,_90_v8);
          let _97_v13;
          _97_v13 = _module.D0.create_DC1(_76_v2, _77_v3, _76_v2);
          let _98_v14;
          let _nw13 = Array((new BigNumber(11)).toNumber()).fill([]);
          _98_v14 = _nw13;
          let _99_v15;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_98_v14);
          _99_v15 = _nw14;
          let _100_v16;
          _100_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_72_v0).length),_99_v15);
          let _101_v17;
          _101_v17 = _dafny.Seq.of(_100_v16);
          let _102_v18;
          _102_v18 = _dafny.Set.fromElements(_73_v1, _73_v1);
          let _103_v19;
          _103_v19 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,_90_v8);
          let _104_v20;
          _104_v20 = _dafny.Map.Empty.slice().updateUnsafe(_90_v8,(_86_v4).fm2(_97_v13, _75_globalState));
          let _105_v21;
          _105_v21 = _dafny.Map.Empty.slice().updateUnsafe(_90_v8,true);
          let _106_v22;
          _106_v22 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,!(_76_v2));
          let _107_v23;
          _107_v23 = _dafny.MultiSet.fromElements(_106_v22, _106_v22);
          let _108_v24;
          _108_v24 = _module.D5.create_DC11(new BigNumber((_107_v23).cardinality()), (_88_v6)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_88_v6).length))]);
          let _109_v25;
          let _init3 = ((_110_v5) => function (_111_i2) {
            return _110_v5;
          })(_87_v5);
          let _nw15 = Array((new BigNumber(4)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw15.length); _i0_3++) {
            _nw15[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _109_v25 = _nw15;
          let _112_v26;
          _112_v26 = _dafny.Map.Empty.slice().updateUnsafe(_109_v25,_76_v2);
          let _pat_let_tv4 = _90_v8;
          let _pat_let_tv5 = _90_v8;
          let _113_v27;
          let _nw16 = Array((new BigNumber(29)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _90_v8;
          _nw16[(_dafny.ONE).toNumber()] = _90_v8;
          _nw16[(new BigNumber(2)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(3)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(4)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(5)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(6)).toNumber()] = (((_96_v12).contains(function (_pat_let6_0) {
            return function (_116_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_117_dt__update_hcf33_h0) {
                  return _module.D9.create_DC21(_117_dt__update_hcf33_h0, (_116_dt__update__tmp_h1).dtor_cf34);
                }(_pat_let7_0);
              }(_pat_let_tv5);
            }(_pat_let6_0);
          }(_95_v11))) ? ((_96_v12).get(function (_pat_let4_0) {
            return function (_114_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_115_dt__update_hcf33_h1) {
                  return _module.D9.create_DC21(_115_dt__update_hcf33_h1, (_114_dt__update__tmp_h2).dtor_cf34);
                }(_pat_let5_0);
              }(_pat_let_tv4);
            }(_pat_let4_0);
          }(_95_v11))) : (_90_v8));
          _nw16[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_86_v4).fm2(_97_v13, _75_globalState), _90_v8);
          _nw16[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_90_v8, (_86_v4).fm2(_module.D0.create_DC1(_76_v2, _77_v3, _76_v2), _75_globalState));
          _nw16[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_101_v17, _101_v17)).length);
          _nw16[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_102_v18).length));
          _nw16[(new BigNumber(11)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(12)).toNumber()] = ((_76_v2) ? (_90_v8) : (_90_v8));
          _nw16[(new BigNumber(13)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(14)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(15)).toNumber()] = new BigNumber((_module.__default.fm21(new BigNumber((_103_v19).length), _75_globalState)).length);
          _nw16[(new BigNumber(16)).toNumber()] = (_90_v8).plus(_90_v8);
          _nw16[(new BigNumber(17)).toNumber()] = (_86_v4).fm2(_97_v13, _75_globalState);
          _nw16[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Seq.update(_91_v9, _module.__default.safeIndex(_90_v8, new BigNumber((_91_v9).length)), _90_v8)).length);
          _nw16[(new BigNumber(19)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(20)).toNumber()] = (((_104_v20).contains(new BigNumber((_72_v0).length))) ? ((_104_v20).get(new BigNumber((_72_v0).length))) : (_90_v8));
          _nw16[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_105_v21).length)).plus(new BigNumber((_77_v3).length)));
          _nw16[(new BigNumber(22)).toNumber()] = (new BigNumber(-632)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_76_v2,_90_v8)).length));
          _nw16[(new BigNumber(23)).toNumber()] = new BigNumber(-302);
          _nw16[(new BigNumber(24)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(25)).toNumber()] = _90_v8;
          _nw16[(new BigNumber(26)).toNumber()] = ((_76_v2) ? ((_108_v24).dtor_cf18) : (new BigNumber((_112_v26).length)));
          _nw16[(new BigNumber(27)).toNumber()] = (_dafny.ZERO).minus((_90_v8).multipliedBy(_90_v8));
          _nw16[(new BigNumber(28)).toNumber()] = _90_v8;
          _113_v27 = _nw16;
          let _index6 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_113_v27).length));
          (_113_v27)[_index6] = (_dafny.ZERO).minus(_90_v8);
          let _118_v28;
          _118_v28 = _dafny.MultiSet.fromElements((_86_v4).fm2(_97_v13, _75_globalState), new BigNumber(334), _90_v8, _module.__default.safeModuloInt(_90_v8, _90_v8), _90_v8);
          let _index7 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_113_v27).length));
          (_113_v27)[_index7] = (((_118_v28).contains((_86_v4).fm2(_module.D0.create_DC1(_76_v2, _dafny.Seq.of(_76_v2), _76_v2), _75_globalState))) ? ((_118_v28).get((_86_v4).fm2(_module.D0.create_DC1(_76_v2, _dafny.Seq.of(_76_v2), _76_v2), _75_globalState))) : (new BigNumber(-339)));
          let _index8 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_73_v1).length));
          (_73_v1)[_index8] = _76_v2;
          let _119_v29;
          _119_v29 = _dafny.MultiSet.fromElements((_99_v15).fm0(_75_globalState));
          let _index9 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_73_v1).length));
          (_73_v1)[_index9] = ((((_119_v29).contains(_76_v2)) ? ((_119_v29).get(_76_v2)) : (_90_v8))).isLessThanOrEqualTo((_113_v27)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_113_v27).length))]);
        } else {
          let _120_v30;
          _120_v30 = new BigNumber(-507);
          let _121_v31;
          _121_v31 = _dafny.Seq.of((_dafny.ZERO).minus(_120_v30));
          let _122_v32;
          _122_v32 = new _dafny.CodePoint('e'.codePointAt(0));
          _72_v0 = _dafny.Seq.update(_72_v0, _module.__default.safeIndex((_121_v31)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_121_v31).length))], new BigNumber((_72_v0).length)), _122_v32);
          let _123_v33;
          _123_v33 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,_73_v1);
          _73_v1 = (((_123_v33).contains(_76_v2)) ? ((_123_v33).get(_76_v2)) : (_73_v1));
          let _124_v34;
          _124_v34 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), ((_125_v32) => function (_126_i3) {
            return _125_v32;
          })(_122_v32)), _72_v0);
          _124_v34 = _124_v34;
          let _127_v35;
          _127_v35 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,new BigNumber((_72_v0).length));
          _127_v35 = (_127_v35).update(_76_v2, _120_v30);
          _120_v30 = (_120_v30).multipliedBy(_120_v30);
        }
        let _128_v36;
        _128_v36 = new BigNumber(115);
        _128_v36 = _128_v36;
        let _129_v37;
        _129_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(44),_76_v2);
        _129_v37 = (_129_v37).update(_128_v36, ((_76_v2) ? (false) : (_76_v2)));
        let _130_v38;
        _130_v38 = new _dafny.CodePoint('c'.codePointAt(0));
        let _131_v40;
        _131_v40 = _module.D4.create_DC7(_72_v0);
        let _132_v41;
        _132_v41 = _dafny.Seq.of(_131_v40, _131_v40, _131_v40, _131_v40, _131_v40);
        let _133_v42;
        _133_v42 = _dafny.Set.fromElements(_132_v41);
        let _134_v43;
        let _nw17 = new _module.C1();
        _nw17.__ctor(_130_v38, (function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-383)), ((_135_v38) => function (_136_i4) {
            return _dafny.Seq.of(_module.D4.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), ((_137_v38) => function (_138_i5) {
  return _137_v38;
})(_135_v38))));
          })(_130_v38))).Elements) {
            let _139_v39 = _compr_6;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-383)), ((_140_v38) => function (_136_i4) {
              return _dafny.Seq.of(_module.D4.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), ((_141_v38) => function (_138_i5) {
  return _141_v38;
})(_140_v38))));
            })(_130_v38)), _139_v39)) {
              _coll6.add(_139_v39);
            }
          }
          return _coll6;
        }()).Union(_133_v42), _76_v2, _76_v2);
        _134_v43 = _nw17;
        _134_v43 = _134_v43;
      } else {
        let _142___mcc_h0 = (_source2).cf0;
        let _143___mcc_h1 = (_source2).cf1;
        let _144___mcc_h2 = (_source2).cf2;
        let _145_cf2 = _144___mcc_h2;
        let _146_cf1 = _143___mcc_h1;
        let _147_cf0 = _142___mcc_h0;
        let _148_v44;
        _148_v44 = new BigNumber(948);
        _147_cf0 = ((new BigNumber((_72_v0).length)).isLessThanOrEqualTo(_148_v44)) && (_145_cf2);
        let _149_v45;
        _149_v45 = _module.D4.create_DC8(_module.__default.safeModuloInt((_dafny.ZERO).minus(_148_v44), _148_v44), _145_cf2);
        let _source4 = _149_v45;
        if (_source4.is_DC8) {
          let _150___mcc_h6 = (_source4).cf11;
          let _151___mcc_h7 = (_source4).cf12;
          let _152_cf12 = _151___mcc_h7;
          let _153_cf11 = _150___mcc_h6;
          let _154_v46;
          _154_v46 = _dafny.Map.Empty.slice().updateUnsafe(_145_cf2,_145_cf2);
          let _155_v47;
          _155_v47 = _dafny.Seq.of(_154_v46);
          let _156_v48;
          _156_v48 = _dafny.MultiSet.fromElements(_147_cf0);
          let _157_v49;
          _157_v49 = _module.D3.create_DC6(_147_cf0, _153_cf11, new BigNumber((_156_v48).cardinality()));
          let _158_v50;
          let _nw18 = Array((new BigNumber(13)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (_154_v46).Merge(_154_v46);
          _nw18[(_dafny.ONE).toNumber()] = (_154_v46).Merge(_154_v46);
          _nw18[(new BigNumber(2)).toNumber()] = _154_v46;
          _nw18[(new BigNumber(3)).toNumber()] = (_155_v47)[_module.__default.safeIndex(_148_v44, new BigNumber((_155_v47).length))];
          _nw18[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,(_157_v49).dtor_cf7);
          _nw18[(new BigNumber(5)).toNumber()] = (_154_v46).Merge(_154_v46);
          _nw18[(new BigNumber(6)).toNumber()] = _154_v46;
          _nw18[(new BigNumber(7)).toNumber()] = ((_154_v46).update(_76_v2, _145_cf2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_76_v2,!(_147_cf0)));
          _nw18[(new BigNumber(8)).toNumber()] = _154_v46;
          _nw18[(new BigNumber(9)).toNumber()] = (_154_v46).Merge(_154_v46);
          _nw18[(new BigNumber(10)).toNumber()] = (_154_v46).Merge(_154_v46);
          _nw18[(new BigNumber(11)).toNumber()] = _module.__default.fm22(_153_cf11, new BigNumber((_72_v0).length), new BigNumber(240), false, _75_globalState);
          _nw18[(new BigNumber(12)).toNumber()] = (_154_v46).update(_152_cf12, _152_cf12);
          _158_v50 = _nw18;
          let _index10 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_158_v50).length));
          (_158_v50)[_index10] = _154_v46;
          let _159_v51;
          let _nw19 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _159_v51 = _nw19;
          let _160_v52;
          _160_v52 = _module.D10.create_DC24(_154_v46);
          let _index11 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_158_v50).length));
          let _rhs8 = (_160_v52).dtor_cf42;
          let _rhs9 = _159_v51;
          let _lhs2 = _158_v50;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_158_v50).length));
          _lhs2[_lhs3] = _rhs8;
          _159_v51 = _rhs9;
          let _161_v53;
          _161_v53 = _dafny.Seq.of(_77_v3, _77_v3, _dafny.Seq.of(_147_cf0), _77_v3, _dafny.Seq.of(true));
          let _162_v54;
          _162_v54 = _module.D11.create_DC26(_161_v53);
          _148_v44 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_148_v44, (_dafny.ZERO).minus(_153_cf11)), new BigNumber(((_162_v54).dtor_cf44).length));
          let _163_v55;
          _163_v55 = _dafny.Set.fromElements(new BigNumber((_72_v0).length), new BigNumber((_77_v3).length));
          _148_v44 = (_153_cf11).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),_163_v55)).length)));
          let _index12 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_159_v51).length));
          (_159_v51)[_index12] = _148_v44;
          let _index13 = _module.__default.safeIndex(new BigNumber(66), new BigNumber((_159_v51).length));
          (_159_v51)[_index13] = (((_145_cf2) ? (_148_v44) : (_148_v44))).plus(_148_v44);
        } else if (_source4.is_DC9) {
          let _164___mcc_h8 = (_source4).cf13;
          let _165___mcc_h9 = (_source4).cf14;
          let _166___mcc_h10 = (_source4).cf15;
          let _167___mcc_h11 = (_source4).cf16;
          let _168_cf16 = _167___mcc_h11;
          let _169_cf15 = _166___mcc_h10;
          let _170_cf14 = _165___mcc_h9;
          let _171_cf13 = _164___mcc_h8;
          _171_cf13 = _171_cf13;
          let _172_v56;
          _172_v56 = _dafny.Set.fromElements(_147_cf0, _76_v2);
          let _173_v57;
          _173_v57 = _dafny.MultiSet.fromElements(new BigNumber((_172_v56).length));
          let _174_v58;
          _174_v58 = _dafny.MultiSet.fromElements(_173_v57);
          let _175_v59;
          _175_v59 = _dafny.Set.fromElements(new BigNumber((_174_v58).cardinality()));
          let _index14 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_73_v1).length));
          (_73_v1)[_index14] = ((_171_cf13) ? (_76_v2) : (false));
          let _176_v60;
          _176_v60 = _module.D2.create_DC3(_170_cf14);
          let _177_v61;
          let _nw20 = new _module.C2();
          _nw20.__ctor((_168_cf16).isEqualTo((_176_v60).dtor_cf4), _76_v2);
          _177_v61 = _nw20;
          let _index15 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_73_v1).length));
          let _rhs10 = _module.__default.fm23(_171_cf13, _dafny.Set.fromElements(_72_v0), _168_cf16, _75_globalState);
          let _rhs11 = (_148_v44).isLessThanOrEqualTo(new BigNumber((_72_v0).length));
          let _rhs12 = _171_cf13;
          let _rhs13 = _170_cf14;
          let _rhs14 = _177_v61;
          let _lhs4 = _73_v1;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_73_v1).length));
          _175_v59 = _rhs10;
          _76_v2 = _rhs11;
          _lhs4[_lhs5] = _rhs12;
          _170_cf14 = _rhs13;
          _177_v61 = _rhs14;
          let _178_v62;
          let _nw21 = new _module.C2();
          _nw21.__ctor(!((_173_v57).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm14(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_179_i6) {
            return new _dafny.CodePoint('i'.codePointAt(0));
          })).length), _72_v0, _76_v2, _75_globalState)).length)))), (_73_v1)[_module.__default.safeIndex(new BigNumber(685), new BigNumber((_73_v1).length))]);
          _178_v62 = _nw21;
          let _180_v63;
          let _nw22 = new _module.C2();
          _nw22.__ctor(((_76_v2) ? (_147_cf0) : ((_178_v62).fm0(_75_globalState))), _76_v2);
          _180_v63 = _nw22;
        } else {
          let _181___mcc_h12 = (_source4).cf10;
          let _182_cf10 = _181___mcc_h12;
          _148_v44 = ((_148_v44).minus(new BigNumber(-862))).minus(_148_v44);
          _145_cf2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), function (_183_i7) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("hduwoi"));
          let _184_v64;
          _184_v64 = _dafny.Map.Empty.slice().updateUnsafe(_148_v44,(new BigNumber((_146_cf1).length)).plus(new BigNumber(859)));
          _148_v44 = (_dafny.ZERO).minus((((_184_v64).contains(new BigNumber((_182_cf10).length))) ? ((_184_v64).get(new BigNumber((_182_cf10).length))) : (_148_v44)));
          let _185_v65;
          let _nw23 = Array((new BigNumber(4)).toNumber()).fill([]);
          _185_v65 = _nw23;
          let _186_v66;
          let _nw24 = new _module.C3();
          _nw24.__ctor(_76_v2, _145_cf2);
          _186_v66 = _nw24;
          let _187_v67;
          _187_v67 = _module.D12.create_DC29(_186_v66);
          let _188_v68;
          let _nw25 = Array((new BigNumber(28)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = _186_v66;
          _nw25[(_dafny.ONE).toNumber()] = _186_v66;
          _nw25[(new BigNumber(2)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(3)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(4)).toNumber()] = (_187_v67).dtor_cf49;
          _nw25[(new BigNumber(5)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(6)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(7)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(8)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(9)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(10)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(11)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(12)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(13)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(14)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(15)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(16)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(17)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(18)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(19)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(20)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(21)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(22)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(23)).toNumber()] = (_module.D12.create_DC29(_186_v66)).dtor_cf49;
          _nw25[(new BigNumber(24)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(25)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(26)).toNumber()] = _186_v66;
          _nw25[(new BigNumber(27)).toNumber()] = _186_v66;
          _188_v68 = _nw25;
          let _index16 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_185_v65).length));
          (_185_v65)[_index16] = _188_v68;
          let _index17 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_185_v65).length));
          (_185_v65)[_index17] = _188_v68;
        }
        _148_v44 = ((_148_v44).minus(_148_v44)).multipliedBy(_module.__default.safeModuloInt(_148_v44, _148_v44));
        if (((_76_v2) ? (_145_cf2) : (_145_cf2))) {
          let _189_v69;
          let _nw26 = Array((new BigNumber(16)).toNumber()).fill([]);
          _189_v69 = _nw26;
          let _190_v70;
          let _nw27 = new _module.C0();
          _nw27.__ctor(_189_v69);
          _190_v70 = _nw27;
          let _191_v71;
          let _nw28 = new _module.C0();
          _nw28.__ctor(_189_v69);
          _191_v71 = _nw28;
          let _192_v72;
          let _nw29 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _192_v72 = _nw29;
          let _index18 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_192_v72).length));
          (_192_v72)[_index18] = (_dafny.ZERO).minus(new BigNumber((_77_v3).length));
          let _193_v73;
          _193_v73 = _dafny.Set.fromElements(new BigNumber(426));
          let _194_v74;
          _194_v74 = _dafny.Seq.of(_148_v44, _148_v44);
          let _index19 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_192_v72).length));
          let _rhs15 = (_193_v73).IsSubsetOf(_193_v73);
          let _rhs16 = _module.__default.safeDivisionInt(((_dafny.ZERO).minus(_148_v44)).plus(_148_v44), ((_dafny.ZERO).minus(new BigNumber((_194_v74).length))).minus(_148_v44));
          let _lhs6 = _192_v72;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_192_v72).length));
          _145_cf2 = _rhs15;
          _lhs6[_lhs7] = _rhs16;
          let _195_v75;
          let _196_v76;
          let _out3;
          let _out4;
          let _outcollector1 = _module.__default.m4(_148_v44, ((_145_cf2) ? (_147_cf0) : (_76_v2)), _72_v0, _75_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _195_v75 = _out3;
          _196_v76 = _out4;
          let _197_v77;
          _197_v77 = _dafny.Set.fromElements(_76_v2);
          _147_cf0 = ((_197_v77).Intersect(_197_v77)).IsSubsetOf(_197_v77);
        } else {
          let _198_v78;
          let _nw30 = new _module.C3();
          _nw30.__ctor(_147_cf0, _145_cf2);
          _198_v78 = _nw30;
          let _199_v79;
          _199_v79 = _module.D12.create_DC29(_198_v78);
          let _pat_let_tv6 = _198_v78;
          let _200_v80;
          let _nw31 = Array((new BigNumber(19)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _199_v79;
          _nw31[(_dafny.ONE).toNumber()] = _199_v79;
          _nw31[(new BigNumber(2)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(3)).toNumber()] = ((_76_v2) ? (_module.D12.create_DC29(_198_v78)) : (_199_v79));
          _nw31[(new BigNumber(4)).toNumber()] = ((_145_cf2) ? (_199_v79) : (_module.D12.create_DC29(_198_v78)));
          _nw31[(new BigNumber(5)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(6)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(7)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(8)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(9)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(10)).toNumber()] = function (_pat_let8_0) {
            return function (_201_dt__update__tmp_h3) {
              return function (_pat_let9_0) {
                return function (_202_dt__update_hcf49_h0) {
                  return _module.D12.create_DC29(_202_dt__update_hcf49_h0);
                }(_pat_let9_0);
              }(_pat_let_tv6);
            }(_pat_let8_0);
          }(_199_v79);
          _nw31[(new BigNumber(11)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(12)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(13)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(14)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(15)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(16)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(17)).toNumber()] = _199_v79;
          _nw31[(new BigNumber(18)).toNumber()] = _199_v79;
          _200_v80 = _nw31;
          _200_v80 = _200_v80;
          let _203_v81;
          _203_v81 = _dafny.Map.Empty.slice().updateUnsafe(_148_v44,_module.__default.fm24(_75_globalState));
          let _204_v82;
          _204_v82 = _dafny.Set.fromElements(_145_cf2);
          let _205_v83;
          _205_v83 = _module.D7.create_DC14(_204_v82);
          _203_v81 = (_203_v81).update((_dafny.ZERO).minus(_148_v44), _module.D7.create_DC16(_205_v83));
          let _206_v84;
          let _207_v85;
          let _out5;
          let _out6;
          let _outcollector2 = _module.__default.m4(_module.__default.safeDivisionInt(_148_v44, _148_v44), false, _72_v0, _75_globalState);
          _out5 = _outcollector2[0];
          _out6 = _outcollector2[1];
          _206_v84 = _out5;
          _207_v85 = _out6;
          let _208_v86;
          let _nw32 = Array((new BigNumber(2)).toNumber()).fill(_dafny.MultiSet.Empty);
          _208_v86 = _nw32;
          _208_v86 = _208_v86;
          let _index20 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_73_v1).length));
          (_73_v1)[_index20] = _207_v85;
          let _209_v87;
          _209_v87 = _dafny.Set.fromElements(_206_v84, _148_v44);
          let _index21 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_73_v1).length));
          (_73_v1)[_index21] = (_209_v87).IsProperSubsetOf(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-626), new BigNumber(343))) {
              let _210_v88 = _compr_7;
              if (((new BigNumber(-626)).isLessThanOrEqualTo(_210_v88)) && ((_210_v88).isLessThan(new BigNumber(343)))) {
                _coll7.add(_module.__default.safeModuloInt(_210_v88, _148_v44));
              }
            }
            return _coll7;
          }());
        }
      }
      let _211_v89;
      _211_v89 = new BigNumber(-166);
      let _212_v90;
      let _213_v91;
      let _out7;
      let _out8;
      let _outcollector3 = _module.__default.m4(_module.__default.safeModuloInt(new BigNumber(827), _211_v89), _76_v2, _module.__default.fm4(_211_v89, new BigNumber(610), _75_globalState), _75_globalState);
      _out7 = _outcollector3[0];
      _out8 = _outcollector3[1];
      _212_v90 = _out7;
      _213_v91 = _out8;
      let _214_v92;
      let _nw33 = Array((new BigNumber(11)).toNumber()).fill([]);
      _214_v92 = _nw33;
      let _215_v93;
      let _nw34 = new _module.C0();
      _nw34.__ctor(_214_v92);
      _215_v93 = _nw34;
      let _216_v94;
      _216_v94 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,_211_v89);
      let _217_v95;
      _217_v95 = _dafny.MultiSet.fromElements(new BigNumber((_216_v94).length));
      let _218_v96;
      _218_v96 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,new BigNumber((_217_v95).cardinality()));
      let _219_v97;
      _219_v97 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_216_v94).Merge(_216_v94)).length),(_211_v89).multipliedBy(_212_v90));
      let _220_v98;
      let _nw35 = new _module.C3();
      _nw35.__ctor(_76_v2, _76_v2);
      _220_v98 = _nw35;
      let _221_v99;
      _221_v99 = _module.D12.create_DC29(_220_v98);
      let _222_v100;
      _222_v100 = _module.D12.create_DC32(_221_v99);
      let _223_v101;
      let _nw36 = Array((new BigNumber(8)).toNumber());
      _nw36[(_dafny.ZERO).toNumber()] = _222_v100;
      _nw36[(_dafny.ONE).toNumber()] = _222_v100;
      _nw36[(new BigNumber(2)).toNumber()] = _222_v100;
      _nw36[(new BigNumber(3)).toNumber()] = _222_v100;
      _nw36[(new BigNumber(4)).toNumber()] = _222_v100;
      _nw36[(new BigNumber(5)).toNumber()] = _module.D12.create_DC32(_module.D12.create_DC29(_220_v98));
      _nw36[(new BigNumber(6)).toNumber()] = _222_v100;
      _nw36[(new BigNumber(7)).toNumber()] = _module.D12.create_DC32(_221_v99);
      _223_v101 = _nw36;
      let _index22 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_223_v101).length));
      (_223_v101)[_index22] = _222_v100;
      let _224_v102;
      let _init4 = ((_225_v90) => function (_226_i8) {
        return _module.__default.safeDivisionInt(_226_i8, _225_v90);
      })(_212_v90);
      let _nw37 = Array((new BigNumber(9)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
        _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _224_v102 = _nw37;
      let _index23 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_223_v101).length));
      let _rhs17 = _219_v97;
      let _rhs18 = _222_v100;
      let _rhs19 = _224_v102;
      let _rhs20 = !((((_213_v91) ? (_212_v90) : (_211_v89))).isLessThanOrEqualTo(_212_v90));
      let _rhs21 = (_dafny.ZERO).minus(_212_v90);
      let _lhs8 = _223_v101;
      let _lhs9 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_223_v101).length));
      _219_v97 = _rhs17;
      _lhs8[_lhs9] = _rhs18;
      _224_v102 = _rhs19;
      _213_v91 = _rhs20;
      _212_v90 = _rhs21;
      let _227_v103;
      _227_v103 = _dafny.Seq.of(_212_v90);
      _227_v103 = _227_v103;
      let _index24 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
      (_73_v1)[_index24] = _dafny.Seq.IsProperPrefixOf(_72_v0, _dafny.Seq.UnicodeFromString("rggpcxlop"));
      let _228_v104;
      _228_v104 = new _dafny.CodePoint('q'.codePointAt(0));
      let _229_v105;
      let _nw38 = Array((new BigNumber(8)).toNumber());
      _nw38[(_dafny.ZERO).toNumber()] = _module.__default.fm4(_212_v90, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-737)), ((_230_v91, _231_v2) => function (_232_i9) {
        return _dafny.Seq.of(_230_v91, _231_v2, _230_v91, _230_v91);
      })(_213_v91, _76_v2))).length), _75_globalState);
      _nw38[(_dafny.ONE).toNumber()] = _72_v0;
      _nw38[(new BigNumber(2)).toNumber()] = ((_213_v91) ? (_72_v0) : (_72_v0));
      _nw38[(new BigNumber(3)).toNumber()] = _72_v0;
      _nw38[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("hw");
      _nw38[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(((_213_v91) ? (_72_v0) : (_72_v0)), _module.__default.safeIndex(_211_v89, new BigNumber((((_213_v91) ? (_72_v0) : (_72_v0))).length)), _228_v104);
      _nw38[(new BigNumber(6)).toNumber()] = _72_v0;
      _nw38[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(589)), function (_233_i10) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      });
      _229_v105 = _nw38;
      let _index25 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
      (_229_v105)[_index25] = _72_v0;
      let _234_v106;
      _234_v106 = _dafny.Map.Empty.slice().updateUnsafe(_73_v1,_76_v2);
      let _235_v107;
      _235_v107 = _dafny.Map.Empty.slice().updateUnsafe(_213_v91,_76_v2);
      let _pat_let_tv7 = _211_v89;
      let _pat_let_tv8 = _235_v107;
      let _index26 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
      let _index27 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
      let _rhs22 = ((((_219_v97).contains(_211_v89)) ? ((_219_v97).get(_211_v89)) : ((_dafny.ZERO).minus(new BigNumber((_234_v106).length))))).isLessThanOrEqualTo((_dafny.ZERO).minus(((_213_v91) ? ((_dafny.ZERO).minus(_211_v89)) : (_211_v89))));
      let _rhs23 = ((_76_v2) ? (_211_v89) : (((_module.__default.fm25(_212_v90, _211_v89, new BigNumber(781), _75_globalState)).dtor_cf51).plus(_212_v90)));
      let _rhs24 = function (_source5) {
        if (_source5.is_DC25) {
          let _236___mcc_h13 = (_source5).cf43;
          let _237_cf43 = _236___mcc_h13;
          return false;
        } else {
          let _238___mcc_h14 = (_source5).cf42;
          let _239_cf42 = _238___mcc_h14;
          return (_pat_let_tv7).isEqualTo(new BigNumber(833));
        }
      }(function (_pat_let10_0) {
        return function (_240_dt__update__tmp_h4) {
          return function (_pat_let11_0) {
            return function (_241_dt__update_hcf42_h0) {
              return _module.D10.create_DC24(_241_dt__update_hcf42_h0);
            }(_pat_let11_0);
          }(_pat_let_tv8);
        }(_pat_let10_0);
      }(_module.D10.create_DC24(_235_v107)));
      let _rhs25 = _72_v0;
      let _lhs10 = _73_v1;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
      let _lhs12 = _229_v105;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
      _213_v91 = _rhs22;
      _212_v90 = _rhs23;
      _lhs10[_lhs11] = _rhs24;
      _lhs12[_lhs13] = _rhs25;
      let _242_v108;
      _242_v108 = _module.D4.create_DC7((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]);
      let _243_v109;
      _243_v109 = _dafny.Set.fromElements(_dafny.Seq.of(_242_v108, _242_v108, _242_v108, _242_v108));
      let _244_v110;
      _244_v110 = _dafny.Seq.of(_220_v98);
      let _245_v111;
      let _nw39 = new _module.C1();
      _nw39.__ctor(_228_v104, (_243_v109).Union(_243_v109), _dafny.areEqual(_244_v110, _244_v110), _213_v91);
      _245_v111 = _nw39;
      let _246_i11;
      _246_i11 = _dafny.ZERO;
      L0: {
        while (_213_v91) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_246_i11)) {
              break L0;
            }
            _246_i11 = (_246_i11).plus(_dafny.ONE);
            let _index28 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
            (_229_v105)[_index28] = _dafny.Seq.Concat((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _72_v0);
            if ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]) {
              let _247_v112;
              _247_v112 = _dafny.MultiSet.fromElements(!(!(_76_v2)), (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
              let _248_v113;
              _248_v113 = _module.D0.create_DC0();
              (_220_v98).m0(_247_v112, _248_v113, _76_v2, _75_globalState);
              _76_v2 = _76_v2;
              let _249_v114;
              _249_v114 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(637),(_245_v111).f10);
              let _250_v115;
              _250_v115 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,_249_v114);
              _213_v91 = (_250_v115).contains(_211_v89);
              let _251_v116;
              let _nw40 = new _module.C0();
              _nw40.__ctor((_215_v93).f12);
              _251_v116 = _nw40;
              let _252_v117;
              _252_v117 = _dafny.Set.fromElements((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _213_v91);
              let _253_v118;
              _253_v118 = _dafny.Seq.of(_252_v117);
              let _254_v119;
              _254_v119 = _module.D7.create_DC14((_253_v118)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_255_v89) => function (_256_i12) {
  return _255_v89;
})(_211_v89))).length), new BigNumber((_253_v118).length))]);
              let _pat_let_tv9 = _252_v117;
              _254_v119 = function (_pat_let12_0) {
                return function (_257_dt__update__tmp_h5) {
                  return function (_pat_let13_0) {
                    return function (_258_dt__update_hcf25_h0) {
                      return _module.D7.create_DC14(_258_dt__update_hcf25_h0);
                    }(_pat_let13_0);
                  }(_pat_let_tv9);
                }(_pat_let12_0);
              }(_254_v119);
            } else {
              let _259_v120;
              _259_v120 = _module.D7.create_DC14(_dafny.Set.fromElements((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]));
              let _260_v121;
              _260_v121 = _module.D7.create_DC16(_259_v120);
              _260_v121 = _module.D7.create_DC16(_259_v120);
              let _261_v122;
              let _nw41 = Array((new BigNumber(5)).toNumber()).fill(_module.D3.Default());
              _261_v122 = _nw41;
              let _262_v123;
              _262_v123 = _module.D3.create_DC6(_76_v2, new BigNumber(839), _212_v90);
              let _index29 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_261_v122).length));
              (_261_v122)[_index29] = _262_v123;
              let _263_v124;
              _263_v124 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(835)), function (_264_i13) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              })).length), _211_v89);
              let _265_v125;
              _265_v125 = _module.D0.create_DC1(!(!(_263_v124).equals(_263_v124)), _77_v3, (_213_v91) || (true));
              let _index30 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_261_v122).length));
              let _rhs26 = _262_v123;
              let _rhs27 = _265_v125;
              let _rhs28 = _228_v104;
              let _lhs14 = _261_v122;
              let _lhs15 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_261_v122).length));
              _lhs14[_lhs15] = _rhs26;
              _265_v125 = _rhs27;
              _228_v104 = _rhs28;
              let _266_v126;
              _266_v126 = _dafny.Set.fromElements(_76_v2, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
              let _267_v127;
              _267_v127 = _module.D2.create_DC3((((_217_v95).contains(new BigNumber((_266_v126).length))) ? ((_217_v95).get(new BigNumber((_266_v126).length))) : (_212_v90)));
              _211_v89 = (_267_v127).dtor_cf4;
              _211_v89 = _212_v90;
              _211_v89 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm9((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _75_globalState), _module.__default.fm4(_212_v90, new BigNumber(921), _75_globalState))).length);
            }
            let _268_v128;
            _268_v128 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), ((_269_v90) => function (_270_i14) {
              return _269_v90;
            })(_212_v90))).length),_217_v95);
            let _271_v129;
            _271_v129 = _module.D11.create_DC27(true, _268_v128);
            _213_v91 = !(!((_271_v129).dtor_cf45));
            let _272_v130;
            _272_v130 = _dafny.MultiSet.fromElements(_220_v98, _220_v98, _220_v98, _220_v98);
            let _273_v131;
            _273_v131 = _module.D12.create_DC30(new BigNumber(885), (_dafny.ZERO).minus(new BigNumber((_272_v130).cardinality())));
            _216_v94 = (_216_v94).update((_213_v91) === ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]), ((_273_v131).dtor_cf51).plus(_211_v89));
          }
        }
      }
      let _hi2 = ((_dafny.ZERO).minus(new BigNumber((_72_v0).length))).plus(new BigNumber(203));
      for (let _274_i15 = new BigNumber(929); _274_i15.isLessThan(_hi2); _274_i15 = _274_i15.plus(_dafny.ONE)) {
        let _275_v133;
        _275_v133 = _dafny.Set.fromElements(new BigNumber((_72_v0).length), new BigNumber((function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of _dafny.IntegerRange(new BigNumber(868), new BigNumber(615))) {
            let _276_v132 = _compr_8;
            if (((new BigNumber(868)).isLessThanOrEqualTo(_276_v132)) && ((_276_v132).isLessThan(new BigNumber(615)))) {
              _coll8.push([_module.__default.safeDivisionInt(_276_v132, _211_v89),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_77_v3).length))))).cardinality())]);
            }
          }
          return _coll8;
        }()).length));
        if ((_dafny.Set.fromElements(_212_v90)).IsProperSubsetOf(_275_v133)) {
          let _277_v134;
          let _nw42 = new _module.C2();
          _nw42.__ctor((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _76_v2);
          _277_v134 = _nw42;
          let _278_v135;
          let _nw43 = Array((new BigNumber(4)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _277_v134;
          _nw43[(_dafny.ONE).toNumber()] = _277_v134;
          _nw43[(new BigNumber(2)).toNumber()] = _277_v134;
          _nw43[(new BigNumber(3)).toNumber()] = _277_v134;
          _278_v135 = _nw43;
          let _index31 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_278_v135).length));
          (_278_v135)[_index31] = _277_v134;
          let _279_v136;
          _279_v136 = _dafny.Map.Empty.slice().updateUnsafe(_76_v2,_277_v134);
          let _index32 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_278_v135).length));
          (_278_v135)[_index32] = (((_279_v136).contains(false)) ? ((_279_v136).get(false)) : (_277_v134));
          _228_v104 = (_245_v111).f10;
          _212_v90 = (_211_v89).minus(_module.__default.safeDivisionInt(_274_i15, _274_i15));
          let _index33 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index33] = (_215_v93).fm0(_75_globalState);
          let _280_v137;
          let _281_v138;
          let _out9;
          let _out10;
          let _outcollector4 = (_220_v98).m2(_277_v134.f8, _75_globalState);
          _out9 = _outcollector4[0];
          _out10 = _outcollector4[1];
          _280_v137 = _out9;
          _281_v138 = _out10;
        } else {
          _216_v94 = (_216_v94).update((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _274_i15);
          _228_v104 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index34 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
          (_229_v105)[_index34] = _dafny.Seq.UnicodeFromString("cubgc");
          let _282_v139;
          let _init5 = ((_283_v2, _284_v3) => function (_285_i16) {
            return (_dafny.MultiSet.fromElements(!(_283_v2))).Intersect(_dafny.MultiSet.FromArray(_284_v3));
          })(_76_v2, _77_v3);
          let _nw44 = Array((new BigNumber(24)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw44.length); _i0_5++) {
            _nw44[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _282_v139 = _nw44;
          let _286_v140;
          _286_v140 = _dafny.Set.fromElements(_213_v91);
          let _287_v141;
          _287_v141 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_219_v97).length)),_286_v140);
          let _rhs29 = _dafny.Seq.UnicodeFromString("ovh");
          let _rhs30 = new BigNumber(((_287_v141).update(_275_v133, (_286_v140).Union(_286_v140))).length);
          let _rhs31 = _282_v139;
          _72_v0 = _rhs29;
          _212_v90 = _rhs30;
          _282_v139 = _rhs31;
          let _288_v142;
          let _nw45 = new _module.C0();
          _nw45.__ctor((_215_v93).f12);
          _288_v142 = _nw45;
        }
        let _289_v143;
        _289_v143 = _module.D0.create_DC1(_76_v2, _77_v3, false);
        let _290_v144;
        let _291_v145;
        let _out11;
        let _out12;
        let _outcollector5 = (_220_v98).m2((_220_v98).fm1((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _289_v143, _75_globalState), _75_globalState);
        _out11 = _outcollector5[0];
        _out12 = _outcollector5[1];
        _290_v144 = _out11;
        _291_v145 = _out12;
        _291_v145 = _76_v2;
        _291_v145 = (_212_v90).isLessThanOrEqualTo(_290_v144);
      }
      _76_v2 = (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))];
      let _source6 = _module.D11.create_DC27(_76_v2, function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(421), new BigNumber(465))) {
    let _292_v146 = _compr_9;
    if (((new BigNumber(421)).isLessThanOrEqualTo(_292_v146)) && ((_292_v146).isLessThan(new BigNumber(465)))) {
      _coll9.push([(_292_v146).minus(_212_v90),_217_v95]);
    }
  }
  return _coll9;
}());
      if (_source6.is_DC27) {
        let _293___mcc_h15 = (_source6).cf45;
        let _294___mcc_h16 = (_source6).cf46;
        let _295_cf46 = _294___mcc_h16;
        let _296_cf45 = _293___mcc_h15;
        if (_76_v2) {
          _212_v90 = _211_v89;
          _242_v108 = _242_v108;
          _220_v98 = _220_v98;
          _211_v89 = _212_v90;
          let _297_v147;
          let _nw46 = new _module.C2();
          _nw46.__ctor(_213_v91, _76_v2);
          _297_v147 = _nw46;
          let _298_v148;
          _298_v148 = _dafny.MultiSet.fromElements((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]);
          let _299_v149;
          _299_v149 = _module.D0.create_DC1(_76_v2, _77_v3, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
          let _300_v150;
          _300_v150 = _dafny.Map.Empty.slice().updateUnsafe((_245_v111).f10,_211_v89);
          let _index35 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          let _rhs32 = (_298_v148).IsProperSubsetOf(_298_v148);
          let _rhs33 = _297_v147;
          let _rhs34 = (_245_v111).fm2(_299_v149, _75_globalState);
          let _rhs35 = (_77_v3)[_module.__default.safeIndex((((_300_v150).contains((_245_v111).f10)) ? ((_300_v150).get((_245_v111).f10)) : (_212_v90)), new BigNumber((_77_v3).length))];
          let _rhs36 = (_297_v147).fm2(_299_v149, _75_globalState);
          let _lhs16 = _73_v1;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          _lhs16[_lhs17] = _rhs32;
          _297_v147 = _rhs33;
          _211_v89 = _rhs34;
          _213_v91 = _rhs35;
          _211_v89 = _rhs36;
        } else {
          let _301_v151;
          _301_v151 = _dafny.Seq.of(_224_v102, _224_v102, _224_v102, _224_v102, _224_v102);
          _296_cf45 = _dafny.Seq.IsProperPrefixOf(_301_v151, _301_v151);
          _76_v2 = !(_213_v91);
          let _index36 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_224_v102).length));
          (_224_v102)[_index36] = new BigNumber(((_dafny.MultiSet.FromArray(_227_v103)).Union(_217_v95)).cardinality());
          let _302_v152;
          _302_v152 = _dafny.Seq.of(_242_v108);
          let _303_v153;
          let _nw47 = new _module.C1();
          _nw47.__ctor(_228_v104, _dafny.Set.fromElements(_302_v152, _302_v152, _302_v152), !_dafny.Seq.contains(_77_v3, true), _76_v2);
          _303_v153 = _nw47;
          let _304_v154;
          _304_v154 = _dafny.Seq.of(_303_v153);
          let _index37 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_224_v102).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          let _rhs37 = new BigNumber(296);
          let _rhs38 = ((_213_v91) ? ((_304_v154)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("xj")).length), new BigNumber((_304_v154).length))]) : (_303_v153));
          let _rhs39 = (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))];
          let _lhs18 = _224_v102;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_224_v102).length));
          let _lhs20 = _73_v1;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          _lhs18[_lhs19] = _rhs37;
          _303_v153 = _rhs38;
          _lhs20[_lhs21] = _rhs39;
          let _305_v155;
          _305_v155 = _dafny.MultiSet.fromElements(_228_v104, (_245_v111).f10);
          let _index39 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index39] = (_305_v155).IsProperSubsetOf(_305_v155);
          let _306_v156;
          let _307_v157;
          let _out13;
          let _out14;
          let _outcollector6 = _module.__default.m4(_211_v89, _296_cf45, (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _75_globalState);
          _out13 = _outcollector6[0];
          _out14 = _outcollector6[1];
          _306_v156 = _out13;
          _307_v157 = _out14;
        }
        let _index40 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        (_73_v1)[_index40] = (_245_v111).fm0(_75_globalState);
        let _308_v158;
        _308_v158 = _228_v104;
        let _309_v159;
        _309_v159 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,(_308_v158));
        let _rhs40 = _211_v89;
        let _rhs41 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_211_v89, new BigNumber((_309_v159).length)), (_212_v90).plus(_211_v89));
        _211_v89 = _rhs40;
        _211_v89 = _rhs41;
        if (_213_v91) {
          let _310_v160;
          let _nw48 = new _module.C0();
          _nw48.__ctor((_215_v93).f12);
          _310_v160 = _nw48;
          let _311_v161;
          let _nw49 = new _module.C1();
          _nw49.__ctor((_245_v111).f10, (_245_v111).f11, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _296_cf45);
          _311_v161 = _nw49;
          let _index41 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index41] = ((_213_v91) ? (_76_v2) : ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]));
          let _index42 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index42] = !(_213_v91);
          _76_v2 = false;
        } else {
          let _312_v162;
          let _313_v163;
          let _out15;
          let _out16;
          let _outcollector7 = (_220_v98).m2(!(((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]) === ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])), _75_globalState);
          _out15 = _outcollector7[0];
          _out16 = _outcollector7[1];
          _312_v162 = _out15;
          _313_v163 = _out16;
          _76_v2 = true;
          let _314_v164;
          _314_v164 = _dafny.MultiSet.fromElements(_313_v163, _213_v91);
          let _315_v165;
          _315_v165 = _module.D0.create_DC0();
          (_245_v111).m0(_314_v164, _315_v165, ((_76_v2) ? (_313_v163) : (_296_cf45)), _75_globalState);
          let _316_v166;
          let _init6 = ((_317_v1) => function (_318_i17) {
            return _module.D8.create_DC19((_317_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_317_v1).length))]);
          })(_73_v1);
          let _nw50 = Array((new BigNumber(20)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw50.length); _i0_6++) {
            _nw50[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _316_v166 = _nw50;
          let _319_v167;
          _319_v167 = _module.D8.create_DC19(_296_cf45);
          let _index43 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_316_v166).length));
          (_316_v166)[_index43] = _319_v167;
          let _pat_let_tv10 = _213_v91;
          let _index44 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_316_v166).length));
          (_316_v166)[_index44] = function (_pat_let14_0) {
            return function (_320_dt__update__tmp_h6) {
              return function (_pat_let15_0) {
                return function (_321_dt__update_hcf31_h0) {
                  return _module.D8.create_DC19(_321_dt__update_hcf31_h0);
                }(_pat_let15_0);
              }(_pat_let_tv10);
            }(_pat_let14_0);
          }(_319_v167);
          _211_v89 = new BigNumber(((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]).length);
        }
      } else if (_source6.is_DC28) {
        let _322___mcc_h17 = (_source6).cf47;
        let _323___mcc_h18 = (_source6).cf48;
        let _324_cf48 = _323___mcc_h18;
        let _325_cf47 = _322___mcc_h17;
        let _326_v168;
        let _nw51 = Array((new BigNumber(19)).toNumber()).fill([]);
        _326_v168 = _nw51;
        let _327_v169;
        let _nw52 = Array((new BigNumber(11)).toNumber());
        _nw52[(_dafny.ZERO).toNumber()] = ((_213_v91) ? (_326_v168) : (_326_v168));
        _nw52[(_dafny.ONE).toNumber()] = _326_v168;
        _nw52[(new BigNumber(2)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(3)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(4)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(5)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(6)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(7)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(8)).toNumber()] = ((_213_v91) ? (_326_v168) : (_326_v168));
        _nw52[(new BigNumber(9)).toNumber()] = _326_v168;
        _nw52[(new BigNumber(10)).toNumber()] = _326_v168;
        _327_v169 = _nw52;
        let _index45 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_327_v169).length));
        (_327_v169)[_index45] = _326_v168;
        let _328_v170;
        _328_v170 = _module.D13.create_DC33(_326_v168);
        let _index46 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_327_v169).length));
        (_327_v169)[_index46] = (_328_v170).dtor_cf56;
        let _329_v171;
        _329_v171 = _dafny.Seq.of(_dafny.Seq.of(_213_v91));
        let _330_v172;
        _330_v172 = _module.D11.create_DC26(_329_v171);
        let _331_v173;
        let _nw53 = new _module.C2();
        _nw53.__ctor((_module.D13.create_DC34(_330_v172, (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])).dtor_cf59, (_212_v90).isLessThan(_211_v89));
        _331_v173 = _nw53;
        _219_v97 = (_219_v97).update(_325_cf47, (_dafny.ZERO).minus(_212_v90));
        let _332_v174;
        _332_v174 = _module.D0.create_DC1(_76_v2, _77_v3, _213_v91);
        let _333_v175;
        let _nw54 = new _module.C2();
        _nw54.__ctor(_76_v2, (_77_v3)[_module.__default.safeIndex((_220_v98).fm2(_332_v174, _75_globalState), new BigNumber((_77_v3).length))]);
        _333_v175 = _nw54;
      } else {
        let _334___mcc_h19 = (_source6).cf44;
        let _335_cf44 = _334___mcc_h19;
        let _rhs42 = (_dafny.ZERO).minus(_212_v90);
        let _rhs43 = _211_v89;
        let _rhs44 = new BigNumber(373);
        _211_v89 = _rhs42;
        _211_v89 = _rhs43;
        _211_v89 = _rhs44;
        let _336_v176;
        _336_v176 = _dafny.Set.fromElements((_211_v89).minus((_dafny.ZERO).minus(_212_v90)));
        let _337_v177;
        let _nw55 = new _module.C3();
        _nw55.__ctor(_76_v2, _213_v91);
        _337_v177 = _nw55;
        let _338_v178;
        _338_v178 = _module.D9.create_DC22(new BigNumber((_216_v94).length), _212_v90, _76_v2, _211_v89, _dafny.Map.Empty.slice().updateUnsafe((_module.D14.create_DC36(_337_v177)).dtor_cf61,_211_v89));
        let _pat_let_tv11 = _212_v90;
        let _339_v179;
        _339_v179 = _dafny.Seq.of(_338_v178, _338_v178, function (_pat_let16_0) {
          return function (_340_dt__update__tmp_h7) {
            return function (_pat_let17_0) {
              return function (_341_dt__update_hcf36_h0) {
                return _module.D9.create_DC22((_340_dt__update__tmp_h7).dtor_cf35, _341_dt__update_hcf36_h0, (_340_dt__update__tmp_h7).dtor_cf37, (_340_dt__update__tmp_h7).dtor_cf38, (_340_dt__update__tmp_h7).dtor_cf39);
              }(_pat_let17_0);
            }(_pat_let_tv11);
          }(_pat_let16_0);
        }(_338_v178), _338_v178);
        let _342_v180;
        _342_v180 = _dafny.Set.fromElements((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]);
        let _343_v182;
        _343_v182 = _dafny.Map.Empty.slice().updateUnsafe(_337_v177,_76_v2);
        let _index47 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
        let _rhs45 = _module.__default.fm23(_76_v2, function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_342_v180).Elements) {
            let _344_v181 = _compr_10;
            if ((_342_v180).contains(_344_v181)) {
              _coll10.add(_344_v181);
            }
          }
          return _coll10;
        }(), (_211_v89).minus(new BigNumber((_343_v182).length)), _75_globalState);
        let _rhs46 = _339_v179;
        let _rhs47 = _module.__default.safeDivisionInt(new BigNumber((_77_v3).length), (_dafny.ZERO).minus(_211_v89));
        let _rhs48 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qwk"), (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]);
        let _lhs22 = _229_v105;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
        _336_v176 = _rhs45;
        _339_v179 = _rhs46;
        _211_v89 = _rhs47;
        _lhs22[_lhs23] = _rhs48;
        let _345_v183;
        let _nw56 = new _module.C2();
        _nw56.__ctor((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _76_v2);
        _345_v183 = _nw56;
        let _346_v184;
        _346_v184 = _dafny.MultiSet.fromElements(_345_v183);
        let _347_v185;
        _347_v185 = _module.D0.create_DC1((_346_v184).IsSubsetOf(_346_v184), _77_v3, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
        let _348_v186;
        _348_v186 = _dafny.MultiSet.fromElements(_228_v104);
        let _349_v187;
        _349_v187 = _dafny.Seq.of(_345_v183, _345_v183, _345_v183, _345_v183, _345_v183);
        let _pat_let_tv12 = _76_v2;
        let _index48 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _rhs49 = true;
        let _rhs50 = !(!((_337_v177).f9) || ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])) || ((_348_v186).IsProperSubsetOf(_dafny.MultiSet.fromElements(_module.__default.fm7(_211_v89, (_337_v177).f9, _212_v90, _75_globalState))));
        let _rhs51 = function (_pat_let18_0) {
          return function (_350_dt__update__tmp_h8) {
            return function (_pat_let19_0) {
              return function (_351_dt__update_hcf0_h0) {
                return _module.D0.create_DC1(_351_dt__update_hcf0_h0, (_350_dt__update__tmp_h8).dtor_cf1, (_350_dt__update__tmp_h8).dtor_cf2);
              }(_pat_let19_0);
            }(_pat_let_tv12);
          }(_pat_let18_0);
        }(_347_v185);
        let _rhs52 = _211_v89;
        let _rhs53 = (_349_v187)[_module.__default.safeIndex(_211_v89, new BigNumber((_349_v187).length))];
        let _lhs24 = _73_v1;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        _lhs24[_lhs25] = _rhs49;
        _213_v91 = _rhs50;
        _347_v185 = _rhs51;
        _212_v90 = _rhs52;
        _345_v183 = _rhs53;
        let _352_v188;
        let _nw57 = Array((new BigNumber(16)).toNumber()).fill([]);
        _352_v188 = _nw57;
        let _index49 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_352_v188).length));
        (_352_v188)[_index49] = _229_v105;
        let _353_v189;
        _353_v189 = _dafny.Seq.of(_224_v102);
        let _354_v190;
        _354_v190 = _dafny.Map.Empty.slice().updateUnsafe(_73_v1,_229_v105);
        let _index50 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_352_v188).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _rhs54 = (((_354_v190).contains(_73_v1)) ? ((_354_v190).get(_73_v1)) : (_229_v105));
        let _rhs55 = _76_v2;
        let _rhs56 = (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))];
        let _rhs57 = _353_v189;
        let _lhs26 = _352_v188;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_352_v188).length));
        let _lhs28 = _73_v1;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        _lhs26[_lhs27] = _rhs54;
        _lhs28[_lhs29] = _rhs55;
        _213_v91 = _rhs56;
        _353_v189 = _rhs57;
      }
      let _index52 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
      (_73_v1)[_index52] = _213_v91;
      let _355_v191;
      _355_v191 = _module.D6.create_DC13(_213_v91, true, (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _213_v91);
      let _pat_let_tv13 = _73_v1;
      let _pat_let_tv14 = _73_v1;
      let _pat_let_tv15 = _213_v91;
      let _source7 = function (_pat_let20_0) {
        return function (_356_dt__update__tmp_h9) {
          return function (_pat_let21_0) {
            return function (_357_dt__update_hcf22_h0) {
              return function (_pat_let22_0) {
                return function (_358_dt__update_hcf24_h0) {
                  return _module.D6.create_DC13((_356_dt__update__tmp_h9).dtor_cf21, _357_dt__update_hcf22_h0, (_356_dt__update__tmp_h9).dtor_cf23, _358_dt__update_hcf24_h0);
                }(_pat_let22_0);
              }(_pat_let_tv15);
            }(_pat_let21_0);
          }((_pat_let_tv14)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_pat_let_tv13).length))]);
        }(_pat_let20_0);
      }(_355_v191);
      if (_source7.is_DC13) {
        let _359___mcc_h20 = (_source7).cf21;
        let _360___mcc_h21 = (_source7).cf22;
        let _361___mcc_h22 = (_source7).cf23;
        let _362___mcc_h23 = (_source7).cf24;
        let _363_cf24 = _362___mcc_h23;
        let _364_cf23 = _361___mcc_h22;
        let _365_cf22 = _360___mcc_h21;
        let _366_cf21 = _359___mcc_h20;
        let _index53 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _rhs58 = !(_366_cf21);
        let _rhs59 = !(_76_v2);
        let _rhs60 = (_220_v98).fm2(_module.D0.create_DC1(!((_77_v3)[_module.__default.safeIndex(_212_v90, new BigNumber((_77_v3).length))]), _77_v3, _76_v2), _75_globalState);
        let _lhs30 = _73_v1;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        _lhs30[_lhs31] = _rhs58;
        _365_cf22 = _rhs59;
        _211_v89 = _rhs60;
        let _367_v192;
        let _nw58 = new _module.C2();
        _nw58.__ctor(_76_v2, _366_cf21);
        _367_v192 = _nw58;
        _367_v192 = ((_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-465)), ((_368_v3) => function (_369_i18) {
          return _368_v3;
        })(_77_v3)), _dafny.Seq.of(_77_v3, _77_v3))) ? (_367_v192) : (_367_v192));
        _213_v91 = _365_cf22;
        _365_cf22 = !(_366_cf21) || ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
      } else {
        let _370___mcc_h24 = (_source7).cf20;
        let _371_cf20 = _370___mcc_h24;
        _212_v90 = (_dafny.ZERO).minus(_212_v90);
        _76_v2 = !(!((_220_v98).fm0(_75_globalState)));
        let _372_v193;
        let _nw59 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _372_v193 = _nw59;
        let _373_v194;
        _373_v194 = _dafny.Seq.of(_372_v193);
        _372_v193 = (_373_v194)[_module.__default.safeIndex(_212_v90, new BigNumber((_373_v194).length))];
        _212_v90 = _212_v90;
      }
      let _374_v195;
      _374_v195 = _module.D6.create_DC12((_215_v93).f12);
      _374_v195 = _374_v195;
      let _375_v196;
      _375_v196 = _dafny.MultiSet.fromElements(_213_v91, _76_v2, _213_v91);
      let _376_v197;
      _376_v197 = _dafny.Seq.of(_73_v1, _73_v1, _73_v1);
      let _377_v198;
      _377_v198 = _dafny.Set.fromElements(new BigNumber(155), _212_v90, new BigNumber((_375_v196).cardinality()), new BigNumber(436), new BigNumber((_376_v197).length));
      let _378_v199;
      _378_v199 = _module.D8.create_DC17(_377_v198);
      let _pat_let_tv16 = _76_v2;
      let _pat_let_tv17 = _217_v95;
      let _pat_let_tv18 = _217_v95;
      let _pat_let_tv19 = _219_v97;
      let _pat_let_tv20 = _235_v107;
      if (function (_source8) {
        if (_source8.is_DC18) {
          let _379___mcc_h29 = (_source8).cf28;
          let _380___mcc_h30 = (_source8).cf29;
          let _381___mcc_h31 = (_source8).cf30;
          let _382_cf30 = _381___mcc_h31;
          let _383_cf29 = _380___mcc_h30;
          let _384_cf28 = _379___mcc_h29;
          return _pat_let_tv16;
        } else if (_source8.is_DC19) {
          let _385___mcc_h32 = (_source8).cf31;
          let _386_cf31 = _385___mcc_h32;
          return (_pat_let_tv17).IsDisjointFrom(_pat_let_tv18);
        } else {
          let _387___mcc_h33 = (_source8).cf27;
          let _388_cf27 = _387___mcc_h33;
          return !(!((new BigNumber((_pat_let_tv19).length)).isLessThanOrEqualTo(new BigNumber(((_module.D10.create_DC24(_pat_let_tv20)).dtor_cf42).length))));
        }
      }((((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]) ? (_378_v199) : (_378_v199)))) {
        let _index54 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
        let _rhs61 = _72_v0;
        let _rhs62 = (new BigNumber((_dafny.Seq.Concat(_227_v103, _227_v103)).length)).multipliedBy(_211_v89);
        let _lhs32 = _229_v105;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
        _lhs32[_lhs33] = _rhs61;
        _212_v90 = _rhs62;
        let _389_v200;
        _389_v200 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-452),_213_v91);
        let _390_v201;
        _390_v201 = _dafny.MultiSet.fromElements(_389_v200);
        let _rhs63 = _212_v90;
        let _rhs64 = (_211_v89).plus(new BigNumber(-447));
        let _rhs65 = _module.__default.fm14(_212_v90, (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _76_v2, _75_globalState);
        let _rhs66 = (_dafny.ZERO).minus(new BigNumber(((((_213_v91) ? (_390_v201) : (_390_v201))).Difference(_390_v201)).cardinality()));
        _211_v89 = _rhs63;
        _212_v90 = _rhs64;
        _227_v103 = _rhs65;
        _212_v90 = _rhs66;
        let _391_v202;
        _391_v202 = _dafny.Map.Empty.slice().updateUnsafe(_245_v111,_213_v91);
        _391_v202 = (_391_v202).update(_245_v111, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
        _73_v1 = _73_v1;
        _211_v89 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_module.__default.fm4(_211_v89, _212_v90, _75_globalState), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-753)), ((_392_v111) => function (_393_i19) {
          return (_392_v111).f10;
        })(_245_v111)), _72_v0))).length));
      } else {
        let _source9 = _module.D3.create_DC6((_module.D6.create_DC13(false, _213_v91, _72_v0, _213_v91)).dtor_cf22, (_dafny.ZERO).minus(new BigNumber((_216_v94).length)), _212_v90);
        if (_source9.is_DC6) {
          let _394___mcc_h25 = (_source9).cf7;
          let _395___mcc_h26 = (_source9).cf8;
          let _396___mcc_h27 = (_source9).cf9;
          let _397_cf9 = _396___mcc_h27;
          let _398_cf8 = _395___mcc_h26;
          let _399_cf7 = _394___mcc_h25;
          let _400_v203;
          let _nw60 = new _module.C3();
          _nw60.__ctor(_76_v2, (_77_v3)[_module.__default.safeIndex(new BigNumber((_77_v3).length), new BigNumber((_77_v3).length))]);
          _400_v203 = _nw60;
          let _401_v204;
          let _402_v205;
          let _403_v206;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector8 = (_220_v98).m1(_72_v0, _399_cf7, _75_globalState);
          _out17 = _outcollector8[0];
          _out18 = _outcollector8[1];
          _out19 = _outcollector8[2];
          _401_v204 = _out17;
          _402_v205 = _out18;
          _403_v206 = _out19;
          _76_v2 = ((_397_cf9).plus(_397_cf9)).isLessThanOrEqualTo(_212_v90);
          let _404_v207;
          let _nw61 = new _module.C2();
          _nw61.__ctor((_403_v206) || (true), _403_v206);
          _404_v207 = _nw61;
        } else {
          let _405___mcc_h28 = (_source9).cf6;
          let _406_cf6 = _405___mcc_h28;
          let _407_v208;
          let _nw62 = new _module.C3();
          _nw62.__ctor((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _213_v91);
          _407_v208 = _nw62;
          let _408_v209;
          _408_v209 = _module.D14.create_DC36(_407_v208);
          let _409_v210;
          _409_v210 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_211_v89, _211_v89),_408_v209);
          _409_v210 = (_409_v210).update((((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]) ? (_212_v90) : (_212_v90)), _408_v209);
          let _index55 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_224_v102).length));
          (_224_v102)[_index55] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_410_v111) => function (_411_i20) {
            return (_410_v111).f10;
          })(_245_v111)), _module.__default.safeIndex(_212_v90, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), ((_412_v111) => function (_413_i20) {
            return (_412_v111).f10;
          })(_245_v111))).length)), _228_v104)).length);
          let _index56 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_224_v102).length));
          (_224_v102)[_index56] = _211_v89;
          _213_v91 = ((_dafny.ZERO).minus(_211_v89)).isLessThanOrEqualTo((_212_v90).multipliedBy(new BigNumber(929)));
          let _414_v211;
          _414_v211 = _dafny.Map.Empty.slice().updateUnsafe((_407_v208).f9,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-383)), ((_415_v111) => function (_416_i21) {
            return (_415_v111).f10;
          })(_245_v111)));
          let _index57 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_224_v102).length));
          (_224_v102)[_index57] = new BigNumber(((_414_v211).update(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), ((_417_v111) => function (_418_i22) {
            return (_417_v111).f10;
          })(_245_v111)), (_245_v111).f10), _dafny.Seq.Concat((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _72_v0))).length);
        }
        let _419_v212;
        let _420_v213;
        let _421_v214;
        let _out20;
        let _out21;
        let _out22;
        let _outcollector9 = (_245_v111).m1((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _76_v2, _75_globalState);
        _out20 = _outcollector9[0];
        _out21 = _outcollector9[1];
        _out22 = _outcollector9[2];
        _419_v212 = _out20;
        _420_v213 = _out21;
        _421_v214 = _out22;
        let _422_v215;
        _422_v215 = _module.D12.create_DC31(_76_v2, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _212_v90);
        let _index58 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _index59 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _rhs67 = _module.__default.safeDivisionInt(_420_v213, _212_v90);
        let _rhs68 = _214_v92;
        let _rhs69 = (_422_v215).dtor_cf53;
        let _rhs70 = _421_v214;
        let _lhs34 = _73_v1;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _lhs36 = _73_v1;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        _211_v89 = _rhs67;
        _214_v92 = _rhs68;
        _lhs34[_lhs35] = _rhs69;
        _lhs36[_lhs37] = _rhs70;
        let _423_v216;
        _423_v216 = _module.D0.create_DC0();
        (_220_v98).m0(_375_v196, _423_v216, _76_v2, _75_globalState);
        let _index60 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_224_v102).length));
        (_224_v102)[_index60] = _212_v90;
        let _index61 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_224_v102).length));
        (_224_v102)[_index61] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(390), _420_v213), _211_v89);
      }
      let _424_v217;
      _424_v217 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,_217_v95);
      let _425_v218;
      _425_v218 = _module.D11.create_DC27((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _424_v217);
      let _source10 = _425_v218;
      if (_source10.is_DC27) {
        let _426___mcc_h34 = (_source10).cf45;
        let _427___mcc_h35 = (_source10).cf46;
        let _428_cf46 = _427___mcc_h35;
        let _429_cf45 = _426___mcc_h34;
        _212_v90 = _212_v90;
        let _index62 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        (_73_v1)[_index62] = _429_cf45;
        let _430_v219;
        _430_v219 = _module.D12.create_DC31(_429_cf45, _429_cf45, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]), _dafny.Seq.of((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _76_v2))).length));
        let _rhs71 = (_212_v90).isEqualTo(_212_v90);
        let _rhs72 = (_dafny.ZERO).minus((_211_v89).minus(_module.__default.safeModuloInt(_212_v90, new BigNumber((_219_v97).length))));
        let _rhs73 = _430_v219;
        _429_cf45 = _rhs71;
        _212_v90 = _rhs72;
        _430_v219 = _rhs73;
        _235_v107 = (_235_v107).update((_213_v91) && ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]), (_dafny.MultiSet.fromElements(new BigNumber(-973))).IsDisjointFrom(_dafny.MultiSet.fromElements(_211_v89)));
      } else if (_source10.is_DC28) {
        let _431___mcc_h36 = (_source10).cf47;
        let _432___mcc_h37 = (_source10).cf48;
        let _433_cf48 = _432___mcc_h37;
        let _434_cf47 = _431___mcc_h36;
        let _index63 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length));
        (_224_v102)[_index63] = (_dafny.ZERO).minus((_211_v89).multipliedBy(_211_v89));
        let _index64 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _index65 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length));
        let _rhs74 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.of(_211_v89, _434_cf47), _227_v103), _227_v103);
        let _rhs75 = ((_76_v2) ? (_211_v89) : (_211_v89));
        let _rhs76 = (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))];
        let _lhs38 = _73_v1;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _lhs40 = _224_v102;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length));
        _lhs38[_lhs39] = _rhs74;
        _lhs40[_lhs41] = _rhs75;
        _76_v2 = _rhs76;
        let _435_v220;
        _435_v220 = _module.D12.create_DC31(_76_v2, _76_v2, new BigNumber(624));
        let _source11 = _435_v220;
        if (_source11.is_DC30) {
          let _436___mcc_h39 = (_source11).cf50;
          let _437___mcc_h40 = (_source11).cf51;
          let _438_cf51 = _437___mcc_h40;
          let _439_cf50 = _436___mcc_h39;
          let _440_v221;
          let _nw63 = new _module.C2();
          _nw63.__ctor(_76_v2, (((_235_v107).contains(_213_v91)) ? ((_235_v107).get(_213_v91)) : (_76_v2)));
          _440_v221 = _nw63;
          let _441_v222;
          _441_v222 = _dafny.Map.Empty.slice().updateUnsafe(_440_v221,_433_cf48);
          _441_v222 = (_441_v222).update(_440_v221, _433_cf48);
          let _index66 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index66] = true;
          _438_cf51 = _438_cf51;
          let _442_v223;
          let _nw64 = new _module.C3();
          _nw64.__ctor(true, _76_v2);
          _442_v223 = _nw64;
        } else if (_source11.is_DC31) {
          let _443___mcc_h41 = (_source11).cf52;
          let _444___mcc_h42 = (_source11).cf53;
          let _445___mcc_h43 = (_source11).cf54;
          let _446_cf54 = _445___mcc_h43;
          let _447_cf53 = _444___mcc_h42;
          let _448_cf52 = _443___mcc_h41;
          let _index67 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index67] = (_dafny.MultiSet.fromElements(new BigNumber(((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]).length))).IsSubsetOf((_217_v95).Union(_dafny.MultiSet.FromArray(_227_v103)));
          let _449_v224;
          let _450_v225;
          let _out23;
          let _out24;
          let _outcollector10 = _module.__default.m4(new BigNumber((_dafny.Seq.of(_211_v89)).length), _213_v91, _dafny.Seq.UnicodeFromString("xexs"), _75_globalState);
          _out23 = _outcollector10[0];
          _out24 = _outcollector10[1];
          _449_v224 = _out23;
          _450_v225 = _out24;
          let _451_v226;
          let _452_v227;
          let _out25;
          let _out26;
          let _outcollector11 = _module.__default.m4(_212_v90, ((_dafny.ZERO).minus(_211_v89)).isLessThan(_449_v224), (_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _75_globalState);
          _out25 = _outcollector11[0];
          _out26 = _outcollector11[1];
          _451_v226 = _out25;
          _452_v227 = _out26;
          _216_v94 = (_216_v94).update(_213_v91, (_449_v224).minus((_dafny.ZERO).minus(_434_cf47)));
        } else if (_source11.is_DC29) {
          let _453___mcc_h44 = (_source11).cf49;
          let _454_cf49 = _453___mcc_h44;
          let _455_v228;
          let _nw65 = new _module.C3();
          _nw65.__ctor(_76_v2, _76_v2);
          _455_v228 = _nw65;
          _455_v228 = _455_v228;
          let _456_v229;
          _456_v229 = _module.D9.create_DC21(new BigNumber(808), (_dafny.ZERO).minus((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]));
          let _457_v230;
          let _nw66 = Array((new BigNumber(26)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _211_v89;
          _nw66[(_dafny.ONE).toNumber()] = new BigNumber(585);
          _nw66[(new BigNumber(2)).toNumber()] = (_212_v90).plus((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]);
          _nw66[(new BigNumber(3)).toNumber()] = _212_v90;
          _nw66[(new BigNumber(4)).toNumber()] = (((_219_v97).contains(_211_v89)) ? ((_219_v97).get(_211_v89)) : (_211_v89));
          _nw66[(new BigNumber(5)).toNumber()] = _434_cf47;
          _nw66[(new BigNumber(6)).toNumber()] = (_211_v89).plus(_211_v89);
          _nw66[(new BigNumber(7)).toNumber()] = new BigNumber(((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]).length);
          _nw66[(new BigNumber(8)).toNumber()] = ((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]).plus(_211_v89);
          _nw66[(new BigNumber(9)).toNumber()] = _434_cf47;
          _nw66[(new BigNumber(10)).toNumber()] = _212_v90;
          _nw66[(new BigNumber(11)).toNumber()] = new BigNumber((_72_v0).length);
          _nw66[(new BigNumber(12)).toNumber()] = ((_dafny.ZERO).minus(_212_v90)).minus(_212_v90);
          _nw66[(new BigNumber(13)).toNumber()] = (_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))];
          _nw66[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]);
          _nw66[(new BigNumber(15)).toNumber()] = new BigNumber((_77_v3).length);
          _nw66[(new BigNumber(16)).toNumber()] = _212_v90;
          _nw66[(new BigNumber(17)).toNumber()] = _434_cf47;
          _nw66[(new BigNumber(18)).toNumber()] = _434_cf47;
          _nw66[(new BigNumber(19)).toNumber()] = ((_456_v229).dtor_cf33).multipliedBy(_211_v89);
          _nw66[(new BigNumber(20)).toNumber()] = (_211_v89).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), ((_458_v111) => function (_459_i23) {
            return (_458_v111).f10;
          })(_245_v111))).length));
          _nw66[(new BigNumber(21)).toNumber()] = ((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]).minus(_212_v90);
          _nw66[(new BigNumber(22)).toNumber()] = new BigNumber(-503);
          _nw66[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]);
          _nw66[(new BigNumber(24)).toNumber()] = _434_cf47;
          _nw66[(new BigNumber(25)).toNumber()] = (_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))];
          _457_v230 = _nw66;
          let _rhs77 = _457_v230;
          let _rhs78 = _212_v90;
          _457_v230 = _rhs77;
          _211_v89 = _rhs78;
          let _460_v231;
          let _nw67 = new _module.C3();
          _nw67.__ctor(_455_v228.f8, (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
          _460_v231 = _nw67;
          let _461_v232;
          _461_v232 = _dafny.Seq.of(_242_v108);
          let _nw68 = new _module.C1();
          _nw68.__ctor((_245_v111).f10, ((_245_v111).f11).Difference(_dafny.Set.fromElements(_461_v232, _461_v232)), (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], (_460_v231).f9);
          _460_v231 = _nw68;
          let _462_v233;
          _462_v233 = _module.D0.create_DC1(_76_v2, _77_v3, _455_v228.f8);
          _76_v2 = ((!(_455_v228.f8)) || ((_245_v111).fm1((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _462_v233, _75_globalState))) === ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
        } else {
          let _463___mcc_h45 = (_source11).cf55;
          let _464_cf55 = _463___mcc_h45;
          let _465_v234;
          let _nw69 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
          _465_v234 = _nw69;
          let _466_v235;
          let _nw70 = new _module.C3();
          _nw70.__ctor((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))], _213_v91);
          _466_v235 = _nw70;
          let _467_v236;
          _467_v236 = _module.D9.create_DC23(_466_v235, _76_v2);
          let _pat_let_tv21 = _76_v2;
          let _468_v237;
          let _nw71 = Array((new BigNumber(15)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = _467_v236;
          _nw71[(_dafny.ONE).toNumber()] = _467_v236;
          _nw71[(new BigNumber(2)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(3)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(4)).toNumber()] = _module.D9.create_DC23(_466_v235, !(_76_v2));
          _nw71[(new BigNumber(5)).toNumber()] = function (_pat_let23_0) {
            return function (_469_dt__update__tmp_h10) {
              return function (_pat_let24_0) {
                return function (_470_dt__update_hcf41_h0) {
                  return _module.D9.create_DC23((_469_dt__update__tmp_h10).dtor_cf40, _470_dt__update_hcf41_h0);
                }(_pat_let24_0);
              }(_pat_let_tv21);
            }(_pat_let23_0);
          }(_467_v236);
          _nw71[(new BigNumber(6)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(7)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(8)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(9)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(10)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(11)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(12)).toNumber()] = _467_v236;
          _nw71[(new BigNumber(13)).toNumber()] = _module.D9.create_DC23(_466_v235, !((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]));
          _nw71[(new BigNumber(14)).toNumber()] = _467_v236;
          _468_v237 = _nw71;
          let _471_v238;
          _471_v238 = _dafny.Map.Empty.slice().updateUnsafe(_212_v90,_220_v98);
          let _472_v239;
          _472_v239 = _dafny.Map.Empty.slice().updateUnsafe(_468_v237,new BigNumber(((_471_v238).update(new BigNumber((_77_v3).length), _220_v98)).length));
          let _index68 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_465_v234).length));
          (_465_v234)[_index68] = _472_v239;
          let _index69 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_465_v234).length));
          (_465_v234)[_index69] = (_472_v239).Merge(_472_v239);
          let _index70 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length));
          (_224_v102)[_index70] = (_227_v103)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(330), _434_cf47), new BigNumber((_227_v103).length))];
          let _index71 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length));
          (_224_v102)[_index71] = ((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]).plus((_dafny.ZERO).minus((_224_v102)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_224_v102).length))]));
          let _473_v240;
          _473_v240 = _module.D0.create_DC1(_466_v235.f8, _77_v3, _213_v91);
          let _474_v241;
          let _475_v242;
          let _476_v243;
          let _out27;
          let _out28;
          let _out29;
          let _outcollector12 = (_220_v98).m1((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], (_220_v98).fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(275)), ((_477_v111) => function (_478_i24) {
            return (_477_v111).f10;
          })(_245_v111)), _473_v240, _75_globalState), _75_globalState);
          _out27 = _outcollector12[0];
          _out28 = _outcollector12[1];
          _out29 = _outcollector12[2];
          _474_v241 = _out27;
          _475_v242 = _out28;
          _476_v243 = _out29;
        }
        let _479_v244;
        let _nw72 = new _module.C2();
        _nw72.__ctor((!((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])) && (_213_v91), _76_v2);
        _479_v244 = _nw72;
        let _index72 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        (_73_v1)[_index72] = _76_v2;
      } else {
        let _480___mcc_h38 = (_source10).cf44;
        let _481_cf44 = _480___mcc_h38;
        let _482_v245;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_214_v92);
        _482_v245 = _nw73;
        let _483_v246;
        _483_v246 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_72_v0, _72_v0)).length),(_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
        let _index73 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        let _rhs79 = _76_v2;
        let _rhs80 = _482_v245;
        let _rhs81 = (((_483_v246).contains(_211_v89)) ? ((_483_v246).get(_211_v89)) : ((_213_v91) === (_213_v91)));
        let _rhs82 = (_212_v90).plus(new BigNumber((_module.__default.fm26(_216_v94, _212_v90, _75_globalState)).length));
        let _lhs42 = _73_v1;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
        _213_v91 = _rhs79;
        _482_v245 = _rhs80;
        _lhs42[_lhs43] = _rhs81;
        _212_v90 = _rhs82;
        let _484_v247;
        _484_v247 = _module.D0.create_DC1((_245_v111).fm0(_75_globalState), _module.__default.fm27(_75_globalState), (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
        if ((_215_v93).fm1(_dafny.Seq.Concat(_72_v0, _72_v0), _484_v247, _75_globalState)) {
          let _index74 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index74] = (_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))];
          _212_v90 = new BigNumber(135);
          _212_v90 = (_211_v89).multipliedBy(new BigNumber(((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]).length));
          let _485_v248;
          let _nw74 = new _module.C1();
          _nw74.__ctor((_245_v111).f10, (_245_v111).f11, _213_v91, true);
          _485_v248 = _nw74;
          let _486_v249;
          _486_v249 = _dafny.Seq.of(_485_v248);
          let _487_v250;
          _487_v250 = _dafny.Map.Empty.slice().updateUnsafe((_486_v249)[_module.__default.safeIndex(_211_v89, new BigNumber((_486_v249).length))],!(false));
          _487_v250 = (_487_v250).update(_485_v248, false);
          let _index75 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_224_v102).length));
          (_224_v102)[_index75] = new BigNumber((_dafny.Seq.UnicodeFromString("mnwoqpjhn")).length);
          let _index76 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_224_v102).length));
          let _index77 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          let _rhs83 = _module.__default.safeDivisionInt(_212_v90, _211_v89);
          let _rhs84 = (_485_v248).f9;
          let _rhs85 = (_211_v89).multipliedBy(_211_v89);
          let _lhs44 = _224_v102;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_224_v102).length));
          let _lhs46 = _73_v1;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          _lhs44[_lhs45] = _rhs83;
          _lhs46[_lhs47] = _rhs84;
          _211_v89 = _rhs85;
        } else {
          _212_v90 = _211_v89;
          _213_v91 = _213_v91;
          let _488_v251;
          let _nw75 = Array((new BigNumber(26)).toNumber()).fill(_dafny.MultiSet.Empty);
          _488_v251 = _nw75;
          let _index78 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_488_v251).length));
          (_488_v251)[_index78] = _375_v196;
          let _489_v252;
          _489_v252 = (_245_v111).f10;
          let _index79 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_488_v251).length));
          (_488_v251)[_index79] = _module.__default.fm19(_76_v2, (_245_v111).f10, (new BigNumber(589)).plus(new BigNumber(((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))]).length)), _75_globalState);
          let _490_v253;
          _490_v253 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_module.__default.fm4(_211_v89, _211_v89, _75_globalState), _module.__default.safeIndex(_212_v90, new BigNumber((_module.__default.fm4(_211_v89, _211_v89, _75_globalState)).length)), (_245_v111).f10),(_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]);
          let _491_v254;
          _491_v254 = _dafny.Map.Empty.slice().updateUnsafe((_211_v89).multipliedBy(new BigNumber(420)),_490_v253);
          _491_v254 = (_491_v254).update(_module.__default.safeModuloInt(_212_v90, new BigNumber((_dafny.Seq.UnicodeFromString("tbnxyymh")).length)), (((_491_v254).contains(new BigNumber((_219_v97).length))) ? ((_491_v254).get(new BigNumber((_219_v97).length))) : (_490_v253)));
          let _492_v255;
          let _493_v256;
          let _494_v257;
          let _out30;
          let _out31;
          let _out32;
          let _outcollector13 = (_220_v98).m1(_dafny.Seq.Concat((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))], _72_v0), !((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))]) || (false), _75_globalState);
          _out30 = _outcollector13[0];
          _out31 = _outcollector13[1];
          _out32 = _outcollector13[2];
          _492_v255 = _out30;
          _493_v256 = _out31;
          _494_v257 = _out32;
        }
        let _495_v258;
        _495_v258 = _dafny.MultiSet.fromElements(_224_v102);
        let _496_v259;
        let _nw76 = new _module.C1();
        _nw76.__ctor((_245_v111).f10, (_245_v111).f11, !((_dafny.MultiSet.fromElements(_224_v102)).IsDisjointFrom(_495_v258)), true);
        _496_v259 = _nw76;
        let _497_v260;
        _497_v260 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,_243_v109);
        let _nw77 = new _module.C1();
        _nw77.__ctor((_245_v111).f10, ((((_497_v260).contains(_212_v90)) ? ((_497_v260).get(_212_v90)) : ((((_497_v260).contains(_212_v90)) ? ((_497_v260).get(_212_v90)) : ((_245_v111).f11))))).Intersect(_243_v109), _76_v2, _213_v91);
        _496_v259 = _nw77;
        if (_76_v2) {
          let _498_v261;
          _498_v261 = _dafny.Set.fromElements(_224_v102);
          let _499_v262;
          let _nw78 = new _module.C2();
          _nw78.__ctor(_213_v91, (_498_v261).IsSubsetOf(_498_v261));
          _499_v262 = _nw78;
          let _500_v263;
          _500_v263 = _dafny.MultiSet.fromElements(_72_v0);
          let _501_v264;
          _501_v264 = _module.D15.create_DC38(_499_v262);
          let _502_v265;
          _502_v265 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_377_v198).length)).isEqualTo(new BigNumber(775)),(_501_v264).dtor_cf66);
          let _503_v266;
          _503_v266 = _dafny.Map.Empty.slice().updateUnsafe(_224_v102,_72_v0);
          let _index80 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          let _rhs86 = (_500_v263).IsDisjointFrom(_500_v263);
          let _rhs87 = (((_502_v265).contains(((_496_v259.f8) ? (_213_v91) : ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])))) ? ((_502_v265).get(((_496_v259.f8) ? (_213_v91) : ((_73_v1)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length))])))) : (_499_v262));
          let _rhs88 = _dafny.Seq.Concat(_72_v0, (((_503_v266).contains(_224_v102)) ? ((_503_v266).get(_224_v102)) : ((_229_v105)[_module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length))])));
          let _lhs48 = _73_v1;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          _lhs48[_lhs49] = _rhs86;
          _499_v262 = _rhs87;
          _72_v0 = _rhs88;
          let _504_v267;
          _504_v267 = _dafny.Seq.of(_217_v95);
          let _505_v268;
          _505_v268 = _dafny.Set.fromElements(((_504_v267)[_module.__default.safeIndex(_211_v89, new BigNumber((_504_v267).length))]).IsSubsetOf(_217_v95), _76_v2);
          _505_v268 = _505_v268;
          let _index81 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
          let _rhs89 = _211_v89;
          let _rhs90 = _72_v0;
          let _lhs50 = _229_v105;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
          _211_v89 = _rhs89;
          _lhs50[_lhs51] = _rhs90;
          let _index82 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_229_v105).length));
          (_229_v105)[_index82] = _72_v0;
          let _506_v269;
          _506_v269 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,_245_v111);
          _506_v269 = _506_v269;
        } else {
          let _507_v270;
          _507_v270 = _dafny.Set.fromElements(_496_v259.f8);
          let _508_v271;
          let _nw79 = new _module.C1();
          _nw79.__ctor(_module.__default.fm7(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_245_v111,new BigNumber((_377_v198).length))).length), _76_v2, _211_v89, _75_globalState), ((_245_v111).f11).Union((_245_v111).f11), (_507_v270).IsProperSubsetOf(_507_v270), _213_v91);
          _508_v271 = _nw79;
          let _index83 = _module.__default.safeIndex(new BigNumber(370), new BigNumber(((_215_v93).f12).length));
          ((_215_v93).f12)[_index83] = (_module.D5.create_DC10(_224_v102)).dtor_cf17;
          let _index84 = _module.__default.safeIndex(new BigNumber(370), new BigNumber(((_215_v93).f12).length));
          ((_215_v93).f12)[_index84] = _224_v102;
          let _509_v272;
          _509_v272 = _dafny.Map.Empty.slice().updateUnsafe(_211_v89,_377_v198);
          _483_v246 = (_483_v246).update(new BigNumber(929), (_377_v198).IsProperSubsetOf((((_509_v272).contains(_212_v90)) ? ((_509_v272).get(_212_v90)) : (_dafny.Set.fromElements(_212_v90, new BigNumber((_377_v198).length), _212_v90)))));
          (_496_v259).f8 = _76_v2;
          let _index85 = _module.__default.safeIndex(new BigNumber(875), new BigNumber((_73_v1).length));
          (_73_v1)[_index85] = !(!(_496_v259.f8)) || (_496_v259.f8);
        }
      }
      process.stdout.write((_72_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_73_v1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_75_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_75_globalState).f6)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_76_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_77_v3, _dafny.Seq.of(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_211_v89));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_212_v90));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_213_v91));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v94).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(14)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_217_v95).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_218_v96).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v97).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(-27058)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_221_v99).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_221_v99).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_222_v100).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_222_v100).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[_dafny.ZERO]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[_dafny.ZERO]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[_dafny.ONE]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[_dafny.ONE]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(2)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(2)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(3)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(3)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(4)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(4)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(5)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(5)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(6)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(6)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_223_v101)[new BigNumber(7)]).dtor_cf55).dtor_cf49.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_223_v101)[new BigNumber(7)]).dtor_cf55).dtor_cf49).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v102)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_227_v103, _dafny.Seq.of(_dafny.ONE, new BigNumber(-881), new BigNumber(473), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_v104));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_229_v105)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_229_v105)[new BigNumber(7)], _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_234_v106).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v107).equals(_dafny.Map.Empty.slice().updateUnsafe(false,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_242_v108).dtor_cf10).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_243_v109).equals(_dafny.Set.fromElements(_dafny.Seq.of(_module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_244_v110).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_245_v111).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_245_v111).f11).equals(_dafny.Set.fromElements(_dafny.Seq.of(_module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")), _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("xe")))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_i11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v191).dtor_cf21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v191).dtor_cf22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_355_v191).dtor_cf23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v191).dtor_cf24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_375_v196).equals(_dafny.MultiSet.fromElements(true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_376_v197).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_377_v198).equals(_dafny.Set.fromElements(new BigNumber(155), new BigNumber(-161), new BigNumber(3), new BigNumber(436)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_378_v199).dtor_cf27).equals(_dafny.Set.fromElements(new BigNumber(155), new BigNumber(-161), new BigNumber(3), new BigNumber(436)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v217).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-761),_dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_425_v218).dtor_cf45));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_425_v218).dtor_cf46).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-761),_dafny.MultiSet.fromElements(_dafny.ONE)))));
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
    static create_DC1(cf0, cf1, cf2) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
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
        return other.$tag === 1 && this.cf0 === other.cf0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2;
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
    static create_DC2(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC4(cf5) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D2(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(false);
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
    static create_DC6(cf7, cf8, cf9) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D3(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC8(cf11, cf12) {
      let $dt = new D4(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC9(cf13, cf14, cf15, cf16) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC7(cf10) {
      let $dt = new D4(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC7" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC8(_dafny.ZERO, false);
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
    static create_DC11(cf18, cf19) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC10(cf17) {
      let $dt = new D5(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf18) + ", " + this.cf19.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf17 === other.cf17;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC11(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC13(cf21, cf22, cf23, cf24) {
      let $dt = new D6(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC12(cf20) {
      let $dt = new D6(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + this.cf23.toVerbatimString(true) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(false, false, _dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC15() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC15";
      } else if (this.$tag === 1) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC15();
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
    static create_DC18(cf28, cf29, cf30) {
      let $dt = new D8(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC19(cf31) {
      let $dt = new D8(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC17(cf27) {
      let $dt = new D8(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf31 === other.cf31;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC18(_dafny.Map.Empty, false, _dafny.ZERO);
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
    static create_DC21(cf33, cf34) {
      let $dt = new D9(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC22(cf35, cf36, cf37, cf38, cf39) {
      let $dt = new D9(1);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC23(cf40, cf41) {
      let $dt = new D9(2);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC20(cf32) {
      let $dt = new D9(3);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf40 === other.cf40 && this.cf41 === other.cf41;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC21(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC25(cf43) {
      let $dt = new D10(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC24(cf42) {
      let $dt = new D10(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC25" + "(" + this.cf43.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC24" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC25(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC27(cf45, cf46) {
      let $dt = new D11(0);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC28(cf47, cf48) {
      let $dt = new D11(1);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC26(cf44) {
      let $dt = new D11(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47) && this.cf48 === other.cf48;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC27(false, _dafny.Map.Empty);
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
    static create_DC30(cf50, cf51) {
      let $dt = new D12(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC31(cf52, cf53, cf54) {
      let $dt = new D12(1);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC29(cf49) {
      let $dt = new D12(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC32(cf55) {
      let $dt = new D12(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50) && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf52 === other.cf52 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf49 === other.cf49;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC30(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC34(cf57, cf58, cf59) {
      let $dt = new D13(0);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC33(cf56) {
      let $dt = new D13(1);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC35(cf60) {
      let $dt = new D13(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf57) + ", " + this.cf58.toVerbatimString(true) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf57, other.cf57) && _dafny.areEqual(this.cf58, other.cf58) && this.cf59 === other.cf59;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf56 === other.cf56;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC34(_module.D11.Default(), _dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC37(cf62, cf63, cf64, cf65) {
      let $dt = new D14(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC36(cf61) {
      let $dt = new D14(1);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64) && this.cf65 === other.cf65;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf61 === other.cf61;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC37(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO, null);
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
    static create_DC39(cf67, cf68, cf69) {
      let $dt = new D15(0);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC40(cf70, cf71) {
      let $dt = new D15(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC38(cf66) {
      let $dt = new D15(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf67 === other.cf67 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf66 === other.cf66;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC39(false, _dafny.ZERO, false);
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
    static create_DC42(cf73, cf74) {
      let $dt = new D16(0);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC41(cf72) {
      let $dt = new D16(1);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73) && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC42(_dafny.Seq.of(), _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
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
      this._f0 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.Seq.UnicodeFromString("");
      this._f5 = _dafny.ZERO;
      this._f6 = [];
      this._f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f12) {
      let _this = this;
      (_this)._f12 = f12;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (new BigNumber(402)).isLessThan(new BigNumber(-807));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eacmud"), _dafny.Seq.UnicodeFromString("h")), ((!(false)) ? (_dafny.Seq.UnicodeFromString("gdjemsj")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), function (_510_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }))));
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return ((function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(846), new BigNumber(41))) {
          let _511_v0 = _compr_11;
          if (((new BigNumber(846)).isLessThanOrEqualTo(_511_v0)) && ((_511_v0).isLessThan(new BigNumber(41)))) {
            _coll11.push([(_511_v0).minus(new BigNumber(450)),new BigNumber(550)]);
          }
        }
        return _coll11;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-609),(_dafny.ZERO).minus(new BigNumber(-626))))).Merge(((true) ? (function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of _dafny.IntegerRange(new BigNumber(942), new BigNumber(845))) {
          let _512_v1 = _compr_12;
          if (((new BigNumber(942)).isLessThanOrEqualTo(_512_v1)) && ((_512_v1).isLessThan(new BigNumber(845)))) {
            _coll12.push([_module.__default.safeDivisionInt(_512_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(275)), function (_513_i0) {
              return new BigNumber(939);
            })).length)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(528), new BigNumber(570)))).cardinality())]);
          }
        }
        return _coll12;
      }()) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(299),new BigNumber((_dafny.Seq.of(new BigNumber(849))).length)))));
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f8 = false;
      this._f9 = false;
      this._f10 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f11 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T2, _module.T0, _module.T1];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    set f8(value) {
      let _this = this;
      _this._f8 = value;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f10, f11, f9, f8) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f9 = f9;
      (_this)._f8 = f8;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return _this.f8;
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return (function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("awbrbsnwd"))).Elements) {
          let _514_v0 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("awbrbsnwd")), _514_v0)) {
            _coll13.push([_514_v0,(_this).f9]);
          }
        }
        return _coll13;
      }()).contains(_dafny.Seq.UnicodeFromString("sapwe"));
    };
    fm2(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((new BigNumber(348)).plus(new BigNumber(959)), (new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality())).plus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), function (_515_i0) {
        return new BigNumber(-831);
      }))).cardinality())));
    };
    fm8(globalState) {
      let _this = this;
      return _this.f8;
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let _516_v0;
      let _nw80 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
      _516_v0 = _nw80;
      let _517_v1;
      _517_v1 = _dafny.Seq.of(p2, p2, (_this).f9, _this.f8, true);
      let _index86 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_516_v0).length));
      (_516_v0)[_index86] = _517_v1;
      let _518_v2;
      _518_v2 = _module.D2.create_DC4(false);
      let _index87 = _module.__default.safeIndex(new BigNumber(467), new BigNumber((_516_v0).length));
      (_516_v0)[_index87] = _dafny.Seq.Concat(_517_v1, _dafny.Seq.of(_this.f8, (_518_v2).dtor_cf5));
      let _519_v3;
      _519_v3 = new BigNumber(-928);
      let _520_v4;
      _520_v4 = _module.D0.create_DC1(_this.f8, _517_v1, p2);
      let _521_v5;
      _521_v5 = _dafny.Seq.UnicodeFromString("riwfnhqa");
      let _hi3 = new BigNumber((_521_v5).length);
      for (let _522_i0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((p2) ? (_519_v3) : ((_this).fm2(_520_v4, globalState))))); _522_i0.isLessThan(_hi3); _522_i0 = _522_i0.plus(_dafny.ONE)) {
        let _523_v6;
        _523_v6 = _dafny.Map.Empty.slice().updateUnsafe(_519_v3,_521_v5);
        let _524_v7;
        _524_v7 = _dafny.Seq.of(_521_v5);
        let _525_v8;
        let _nw81 = Array((new BigNumber(25)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = _521_v5;
        _nw81[(_dafny.ONE).toNumber()] = _module.__default.fm9(!(_this.f8), _this.f8, globalState);
        _nw81[(new BigNumber(2)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(3)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(4)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(5)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(6)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(7)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat((((_523_v6).contains((_this).fm2(_520_v4, globalState))) ? ((_523_v6).get((_this).fm2(_520_v4, globalState))) : (_521_v5)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(430)), function (_526_i1) {
          return (_this).f10;
        }));
        _nw81[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("fltsycqgp");
        _nw81[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_521_v5, _521_v5);
        _nw81[(new BigNumber(11)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(12)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(13)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("rboifjmne");
        _nw81[(new BigNumber(15)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_521_v5, _dafny.Seq.UnicodeFromString("mk"));
        _nw81[(new BigNumber(17)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("nvmycqm"), _module.__default.safeIndex(_522_i0, new BigNumber((_dafny.Seq.UnicodeFromString("nvmycqm")).length)), (_this).f10);
        _nw81[(new BigNumber(19)).toNumber()] = _521_v5;
        _nw81[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("tjel");
        _nw81[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_521_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), function (_527_i2) {
          return (_this).f10;
        }));
        _nw81[(new BigNumber(22)).toNumber()] = (_524_v7)[_module.__default.safeIndex((_this).fm2(_520_v4, globalState), new BigNumber((_524_v7).length))];
        _nw81[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_528_i3) {
          return (_this).f10;
        });
        _nw81[(new BigNumber(24)).toNumber()] = _521_v5;
        _525_v8 = _nw81;
        let _index88 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_525_v8).length));
        (_525_v8)[_index88] = _521_v5;
        let _index89 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_525_v8).length));
        (_525_v8)[_index89] = _dafny.Seq.update(_521_v5, _module.__default.safeIndex(_522_i0, new BigNumber((_521_v5).length)), new _dafny.CodePoint('n'.codePointAt(0)));
        (_this).f8 = !(p2);
        (_this).f8 = !((_this).f9);
        (_this).f8 = (_this).fm8(globalState);
      }
      let _529_v9;
      _529_v9 = _dafny.Set.fromElements(_519_v3);
      (_this).f8 = (_529_v9).IsProperSubsetOf(_529_v9);
      let _530_i4;
      _530_i4 = _dafny.ZERO;
      L1: {
        while ((_519_v3).isLessThanOrEqualTo(_519_v3)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_530_i4)) {
              break L1;
            }
            _530_i4 = (_530_i4).plus(_dafny.ONE);
            let _531_v10;
            _531_v10 = _dafny.Set.fromElements(_this.f8, false);
            let _532_v11;
            _532_v11 = _dafny.Seq.of(_531_v10);
            let _533_v12;
            _533_v12 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("yial"), _521_v5)),_532_v11);
            _533_v12 = (_533_v12).update(_this.f8, _532_v11);
            let _534_v13;
            let _nw82 = Array((new BigNumber(18)).toNumber()).fill(false);
            _534_v13 = _nw82;
            let _535_v14;
            _535_v14 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
            let _index90 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_534_v13).length));
            (_534_v13)[_index90] = (((_535_v14).contains(false)) ? ((_535_v14).get(false)) : ((_this).fm8(globalState)));
            let _index91 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_534_v13).length));
            (_534_v13)[_index91] = (_this).fm0(globalState);
            let _index92 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_534_v13).length));
            (_534_v13)[_index92] = (_this).f9;
            (_this).f8 = _this.f8;
          }
        }
      }
      if (_this.f8) {
        let _536_v15;
        let _init7 = function (_537_i5) {
          return _this.f8;
        };
        let _nw83 = Array((new BigNumber(22)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw83.length); _i0_7++) {
          _nw83[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _536_v15 = _nw83;
        _536_v15 = _536_v15;
        let _538_v16;
        let _nw84 = Array((new BigNumber(8)).toNumber());
        _nw84[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("atblghtma");
        _nw84[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(133)), function (_539_i6) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        });
        _nw84[(new BigNumber(2)).toNumber()] = _521_v5;
        _nw84[(new BigNumber(3)).toNumber()] = _521_v5;
        _nw84[(new BigNumber(4)).toNumber()] = _521_v5;
        _nw84[(new BigNumber(5)).toNumber()] = _521_v5;
        _nw84[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("hvtuvijkn");
        _nw84[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_521_v5, _521_v5);
        _538_v16 = _nw84;
        let _540_v17;
        _540_v17 = _dafny.Seq.of(_521_v5);
        let _index93 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length));
        (_538_v16)[_index93] = (_540_v17)[_module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.fromElements(_this.f8, !(p2), _this.f8)).update(!(_this.f8), _module.__default.abs(_519_v3))).cardinality()), new BigNumber((_540_v17).length))];
        let _541_v18;
        _541_v18 = _dafny.Map.Empty.slice().updateUnsafe(_519_v3,_519_v3);
        let _index94 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length));
        (_538_v16)[_index94] = _dafny.Seq.update(((p2) ? ((_540_v17)[_module.__default.safeIndex((((_541_v18).contains(_519_v3)) ? ((_541_v18).get(_519_v3)) : (_519_v3)), new BigNumber((_540_v17).length))]) : (_dafny.Seq.Concat(_521_v5, _521_v5))), _module.__default.safeIndex((new BigNumber(754)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), function (_542_i7) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length)), new BigNumber((((p2) ? ((_540_v17)[_module.__default.safeIndex((((_541_v18).contains(_519_v3)) ? ((_541_v18).get(_519_v3)) : (_519_v3)), new BigNumber((_540_v17).length))]) : (_dafny.Seq.Concat(_521_v5, _521_v5)))).length)), (_this).f10);
        let _index95 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length));
        (_538_v16)[_index95] = _dafny.Seq.Concat(_dafny.Seq.Concat((_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))], (_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))]), (_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))]);
        let _543_v19;
        _543_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f8);
        let _544_v20;
        _544_v20 = _dafny.MultiSet.fromElements(new BigNumber(((_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))]).length), _519_v3, new BigNumber((p0).cardinality()), _519_v3);
        _543_v19 = (_543_v19).update(_this.f8, !(_544_v20).contains(new BigNumber((_543_v19).length)));
        if (!(!((_this).f9)) || (_this.f8)) {
          (_this).f8 = ((new BigNumber(766)).plus(new BigNumber(((_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))]).length))).isLessThan(_module.__default.safeDivisionInt(_519_v3, _519_v3));
          let _545_v21;
          _545_v21 = _dafny.Seq.of(_519_v3, _519_v3, _519_v3);
          let _546_v22;
          _546_v22 = _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("arjmj"));
          let _547_v23;
          _547_v23 = _dafny.Map.Empty.slice().updateUnsafe(_546_v22,_519_v3);
          let _548_v24;
          let _nw85 = Array((new BigNumber(9)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _519_v3;
          _nw85[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_519_v3);
          _nw85[(new BigNumber(2)).toNumber()] = new BigNumber((_545_v21).length);
          _nw85[(new BigNumber(3)).toNumber()] = (_519_v3).minus(new BigNumber(898));
          _nw85[(new BigNumber(4)).toNumber()] = (_this).fm2(_520_v4, globalState);
          _nw85[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_519_v3, new BigNumber((_dafny.Seq.of(_519_v3, (_this).fm2(_520_v4, globalState))).length));
          _nw85[(new BigNumber(6)).toNumber()] = _519_v3;
          _nw85[(new BigNumber(7)).toNumber()] = _519_v3;
          _nw85[(new BigNumber(8)).toNumber()] = (_519_v3).minus((((_547_v23).contains(_546_v22)) ? ((_547_v23).get(_546_v22)) : ((_this).fm2(_module.D0.create_DC1(_this.f8, (_516_v0)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_516_v0).length))], !(_this.f8)), globalState))));
          _548_v24 = _nw85;
          let _index96 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_548_v24).length));
          (_548_v24)[_index96] = _519_v3;
          let _549_v25;
          _549_v25 = _dafny.Set.fromElements(_this.f8);
          let _index97 = _module.__default.safeIndex(new BigNumber(585), new BigNumber((_548_v24).length));
          (_548_v24)[_index97] = new BigNumber((_549_v25).length);
          let _index98 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_536_v15).length));
          (_536_v15)[_index98] = true;
          let _index99 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_536_v15).length));
          (_536_v15)[_index99] = (_this).f9;
          let _550_v26;
          _550_v26 = _dafny.Seq.of(_545_v21, _545_v21);
          _550_v26 = _550_v26;
          let _551_v27;
          let _nw86 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _551_v27 = _nw86;
          let _552_v28;
          _552_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_548_v24)[_module.__default.safeIndex(new BigNumber(585), new BigNumber((_548_v24).length))]);
          let _index100 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_551_v27).length));
          (_551_v27)[_index100] = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_552_v28);
          let _553_v30;
          _553_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_552_v28);
          let _index101 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_551_v27).length));
          (_551_v27)[_index101] = (function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), function (_554_i8) {
              return (_this).f10;
            })).Elements) {
              let _555_v29 = _compr_14;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-42)), function (_554_i8) {
                return (_this).f10;
              }), _555_v29)) {
                _coll14.push([_555_v29,_552_v28]);
              }
            }
            return _coll14;
          }()).Merge(_553_v30);
        } else {
          let _556_v31;
          _556_v31 = _dafny.Map.Empty.slice().updateUnsafe(_521_v5,_519_v3);
          _556_v31 = (_556_v31).update((_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))], _519_v3);
          _518_v2 = _518_v2;
          let _557_v32;
          let _nw87 = Array((new BigNumber(8)).toNumber()).fill([]);
          _557_v32 = _nw87;
          let _558_v33;
          let _nw88 = new _module.C0();
          _nw88.__ctor(_557_v32);
          _558_v33 = _nw88;
          let _559_v34;
          let _init8 = function (_560_i9) {
            return (_this).f10;
          };
          let _nw89 = Array((new BigNumber(16)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw89.length); _i0_8++) {
            _nw89[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _559_v34 = _nw89;
          let _index102 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_559_v34).length));
          (_559_v34)[_index102] = (_this).f10;
          let _561_v35;
          _561_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_541_v18).length),_521_v5);
          let _index103 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_559_v34).length));
          let _rhs91 = (_this).f10;
          let _rhs92 = _536_v15;
          let _rhs93 = (_519_v3).isLessThan(_519_v3);
          let _rhs94 = (_558_v33).fm1(_dafny.Seq.Concat((_538_v16)[_module.__default.safeIndex(new BigNumber(788), new BigNumber((_538_v16).length))], _dafny.Seq.update((((_561_v35).contains((_dafny.ZERO).minus(_519_v3))) ? ((_561_v35).get((_dafny.ZERO).minus(_519_v3))) : ((_540_v17)[_module.__default.safeIndex(_519_v3, new BigNumber((_540_v17).length))])), _module.__default.safeIndex((_this).fm2(_module.D0.create_DC1(!(!(p2)), (_516_v0)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_516_v0).length))], _this.f8), globalState), new BigNumber(((((_561_v35).contains((_dafny.ZERO).minus(_519_v3))) ? ((_561_v35).get((_dafny.ZERO).minus(_519_v3))) : ((_540_v17)[_module.__default.safeIndex(_519_v3, new BigNumber((_540_v17).length))]))).length)), (_this).f10)), _520_v4, globalState);
          let _lhs52 = _559_v34;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(42), new BigNumber((_559_v34).length));
          let _lhs54 = _this;
          let _lhs55 = _this;
          _lhs52[_lhs53] = _rhs91;
          _536_v15 = _rhs92;
          _lhs54.f8 = _rhs93;
          _lhs55.f8 = _rhs94;
          (_this).f8 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_this.f8, (_this).f9), (_516_v0)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_516_v0).length))]);
        }
      } else {
        let _562_v36;
        let _nw90 = Array((new BigNumber(18)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = (_519_v3).plus(_519_v3);
        _nw90[(_dafny.ONE).toNumber()] = _519_v3;
        _nw90[(new BigNumber(2)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(3)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(4)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(330), _519_v3);
        _nw90[(new BigNumber(6)).toNumber()] = new BigNumber(408);
        _nw90[(new BigNumber(7)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(8)).toNumber()] = (_this).fm2(_module.D0.create_DC1(false, _517_v1, (_this).f9), globalState);
        _nw90[(new BigNumber(9)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(10)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(11)).toNumber()] = (new BigNumber(392)).minus(_519_v3);
        _nw90[(new BigNumber(12)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(13)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(14)).toNumber()] = new BigNumber((_521_v5).length);
        _nw90[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(_519_v3, _519_v3);
        _nw90[(new BigNumber(16)).toNumber()] = _519_v3;
        _nw90[(new BigNumber(17)).toNumber()] = _519_v3;
        _562_v36 = _nw90;
        _562_v36 = _562_v36;
        _519_v3 = _519_v3;
        let _563_v37;
        _563_v37 = _module.D5.create_DC10(_562_v36);
        let _564_v38;
        _564_v38 = _dafny.Seq.of(_562_v36);
        let _565_v39;
        let _nw91 = Array((new BigNumber(19)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = _562_v36;
        _nw91[(_dafny.ONE).toNumber()] = _562_v36;
        _nw91[(new BigNumber(2)).toNumber()] = (_563_v37).dtor_cf17;
        _nw91[(new BigNumber(3)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(4)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(5)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(6)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(7)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(8)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(9)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(10)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(11)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(12)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(13)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(14)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(15)).toNumber()] = (_564_v38)[_module.__default.safeIndex(_519_v3, new BigNumber((_564_v38).length))];
        _nw91[(new BigNumber(16)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(17)).toNumber()] = _562_v36;
        _nw91[(new BigNumber(18)).toNumber()] = _562_v36;
        _565_v39 = _nw91;
        let _566_v40;
        let _nw92 = new _module.C0();
        _nw92.__ctor(_565_v39);
        _566_v40 = _nw92;
        let _index104 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_562_v36).length));
        (_562_v36)[_index104] = _519_v3;
        let _index105 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_562_v36).length));
        let _rhs95 = (((_this).fm2(_520_v4, globalState)).plus(_519_v3)).plus(_module.__default.safeDivisionInt(_519_v3, _519_v3));
        let _rhs96 = _519_v3;
        let _rhs97 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("gucbasaxj")).length), _519_v3);
        let _rhs98 = _521_v5;
        let _lhs56 = _562_v36;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_562_v36).length));
        _lhs56[_lhs57] = _rhs95;
        _519_v3 = _rhs96;
        _519_v3 = _rhs97;
        _521_v5 = _rhs98;
        let _567_v41;
        let _nw93 = new _module.C0();
        _nw93.__ctor((_566_v40).f12);
        _567_v41 = _nw93;
        _567_v41 = _566_v40;
      }
      let _568_v42;
      _568_v42 = _module.D3.create_DC6((_this).f9, _519_v3, new BigNumber((_521_v5).length));
      let _569_v43;
      _569_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_568_v42);
      let _570_v44;
      _570_v44 = _dafny.MultiSet.fromElements(_519_v3);
      _569_v43 = (_569_v43).update(((p2) ? (p2) : (_this.f8)), _module.D3.create_DC6((_this).f9, new BigNumber(((_570_v44).update(_519_v3, _module.__default.abs(_519_v3))).cardinality()), new BigNumber(307)));
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _571_v0;
      let _nw94 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
      _571_v0 = _nw94;
      let _572_v1;
      _572_v1 = new BigNumber(802);
      let _573_v2;
      _573_v2 = _dafny.Set.fromElements(_572_v1);
      let _index106 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_571_v0).length));
      (_571_v0)[_index106] = _573_v2;
      let _574_v3;
      _574_v3 = _dafny.Set.fromElements((_this).f10, (_this).f10);
      let _index107 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_571_v0).length));
      (_571_v0)[_index107] = _dafny.Set.fromElements((_572_v1).plus(new BigNumber((_574_v3).length)));
      let _hi4 = new BigNumber((_dafny.Seq.of(new BigNumber(337), (_dafny.ZERO).minus(_572_v1))).length);
      for (let _575_i0 = _572_v1; _575_i0.isLessThan(_hi4); _575_i0 = _575_i0.plus(_dafny.ONE)) {
        if ((_this).fm0(globalState)) {
          let _576_v4;
          _576_v4 = _module.D5.create_DC11(_575_i0, p0);
          let _577_v5;
          _577_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f10);
          _576_v4 = _module.__default.fm11((_577_v5).Merge(_577_v5), (_this).f9, !(((false) ? (p1) : (_this.f8))), _this.f8, globalState);
          let _578_v6;
          let _nw95 = Array((new BigNumber(4)).toNumber()).fill(false);
          _578_v6 = _nw95;
          let _index108 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length));
          (_578_v6)[_index108] = (_this).f9;
          let _579_v7;
          _579_v7 = _dafny.Seq.of(_577_v5);
          let _580_v8;
          _580_v8 = _dafny.MultiSet.fromElements(p1);
          let _581_v9;
          let _nw96 = Array((_dafny.ONE).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _575_i0;
          _581_v9 = _nw96;
          let _582_v10;
          _582_v10 = _dafny.MultiSet.fromElements(_581_v9, _581_v9);
          let _583_v11;
          _583_v11 = _module.D5.create_DC10(_581_v9);
          let _584_v12;
          _584_v12 = _dafny.Seq.of(_582_v10, _dafny.MultiSet.fromElements(_581_v9, _581_v9, (_583_v11).dtor_cf17, _581_v9, _581_v9), _582_v10);
          let _585_v13;
          _585_v13 = _dafny.Seq.of(p1);
          let _586_v14;
          _586_v14 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_585_v13, _module.__default.safeIndex(_575_i0, new BigNumber((_585_v13).length)), p1)).length), _572_v1, _575_i0);
          let _587_v15;
          _587_v15 = _dafny.Map.Empty.slice().updateUnsafe(_575_i0,p1);
          let _588_v16;
          _588_v16 = _module.D0.create_DC1(p1, _585_v13, _this.f8);
          let _589_v17;
          _589_v17 = _dafny.MultiSet.fromElements((_this).f10);
          let _590_v18;
          let _nw97 = Array((new BigNumber(29)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = new BigNumber(((_579_v7)[_module.__default.safeIndex(_572_v1, new BigNumber((_579_v7).length))]).length);
          _nw97[(_dafny.ONE).toNumber()] = (_575_i0).plus(_572_v1);
          _nw97[(new BigNumber(2)).toNumber()] = (((_580_v8).contains(p1)) ? ((_580_v8).get(p1)) : (_572_v1));
          _nw97[(new BigNumber(3)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(4)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(5)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(6)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(7)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(8)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(9)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(10)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(11)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(12)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(13)).toNumber()] = new BigNumber(((_584_v12)[_module.__default.safeIndex(new BigNumber((_586_v14).length), new BigNumber((_584_v12).length))]).cardinality());
          _nw97[(new BigNumber(14)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(15)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(16)).toNumber()] = new BigNumber(-560);
          _nw97[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_587_v15).length), _575_i0);
          _nw97[(new BigNumber(18)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(19)).toNumber()] = new BigNumber((_586_v14).length);
          _nw97[(new BigNumber(20)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(21)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(22)).toNumber()] = _575_i0;
          _nw97[(new BigNumber(23)).toNumber()] = _572_v1;
          _nw97[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus((_572_v1).plus(_572_v1));
          _nw97[(new BigNumber(25)).toNumber()] = _module.__default.safeModuloInt(_575_i0, _575_i0);
          _nw97[(new BigNumber(26)).toNumber()] = (_this).fm2(function (_pat_let25_0) {
            return function (_591_dt__update__tmp_h0) {
              return function (_pat_let26_0) {
                return function (_592_dt__update_hcf0_h0) {
                  return _module.D0.create_DC1(_592_dt__update_hcf0_h0, (_591_dt__update__tmp_h0).dtor_cf1, (_591_dt__update__tmp_h0).dtor_cf2);
                }(_pat_let26_0);
              }((_this).f9);
            }(_pat_let25_0);
          }(_588_v16), globalState);
          _nw97[(new BigNumber(27)).toNumber()] = (_575_i0).minus((((_589_v17).contains((_this).f10)) ? ((_589_v17).get((_this).f10)) : (_575_i0)));
          _nw97[(new BigNumber(28)).toNumber()] = new BigNumber(-577);
          _590_v18 = _nw97;
          let _index109 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_581_v9).length));
          (_581_v9)[_index109] = new BigNumber((_dafny.Seq.Concat(_586_v14, _586_v14)).length);
          let _index110 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length));
          let _index111 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_581_v9).length));
          let _rhs99 = _module.__default.fm9((_575_i0).isEqualTo(_575_i0), true, globalState);
          let _rhs100 = p1;
          let _rhs101 = (_575_i0).minus(_575_i0);
          let _lhs58 = _578_v6;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length));
          let _lhs60 = _581_v9;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_581_v9).length));
          r0 = _rhs99;
          _lhs58[_lhs59] = _rhs100;
          _lhs60[_lhs61] = _rhs101;
          let _593_v19;
          _593_v19 = _module.D4.create_DC8(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_594_i1) {
  return new _dafny.CodePoint('b'.codePointAt(0));
})).length), (_this).f9);
          let _index112 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length));
          (_578_v6)[_index112] = (_593_v19).dtor_cf12;
          let _595_v20;
          _595_v20 = _dafny.Seq.of(p0, p0, p0);
          let _596_v21;
          _596_v21 = _dafny.Map.Empty.slice().updateUnsafe(_595_v20,(_578_v6)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length))]);
          let _597_v22;
          _597_v22 = _module.D3.create_DC5(_586_v14);
          let _index113 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_578_v6).length));
          (_578_v6)[_index113] = (((_596_v21).contains(_module.__default.fm12(p1, _597_v22, globalState))) ? ((_596_v21).get(_module.__default.fm12(p1, _597_v22, globalState))) : ((_585_v13)[_module.__default.safeIndex((_this).fm2(_588_v16, globalState), new BigNumber((_585_v13).length))]));
          r2 = ((_581_v9)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_581_v9).length))]).isLessThan((_581_v9)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_581_v9).length))]);
        } else {
          let _598_v23;
          _598_v23 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(p0, p0), _dafny.Seq.Concat(p0, p0));
          _598_v23 = (_598_v23).Difference(((p1) ? (_dafny.MultiSet.fromElements(p0, p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(376)), function (_599_i2) {
            return (_this).f10;
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), function (_600_i3) {
            return (_this).f10;
          }))) : (_dafny.MultiSet.fromElements(p0, p0, p0))));
          let _601_v24;
          let _nw98 = Array((new BigNumber(11)).toNumber()).fill([]);
          _601_v24 = _nw98;
          let _602_v25;
          let _nw99 = new _module.C0();
          _nw99.__ctor(_601_v24);
          _602_v25 = _nw99;
          let _603_v26;
          _603_v26 = _dafny.Seq.of(_602_v25, _602_v25, _602_v25, _602_v25, _602_v25);
          let _604_v27;
          _604_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(-991));
          let _605_v28;
          _605_v28 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_this.f8,_572_v1), _604_v27);
          let _606_v29;
          _606_v29 = _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber((_603_v26).length))).multipliedBy(_572_v1), _575_i0, _module.__default.safeModuloInt(_575_i0, _572_v1), _575_i0, (_575_i0).multipliedBy(new BigNumber(((_605_v28)[_module.__default.safeIndex(_575_i0, new BigNumber((_605_v28).length))]).length)));
          let _rhs102 = new BigNumber((_606_v29).length);
          let _rhs103 = p1;
          _572_v1 = _rhs102;
          r2 = _rhs103;
          let _607_v30;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_601_v24);
          _607_v30 = _nw100;
          let _608_v31;
          _608_v31 = _dafny.MultiSet.fromElements(_572_v1);
          _608_v31 = (_608_v31).Intersect(_608_v31);
          let _609_v32;
          let _nw101 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _609_v32 = _nw101;
          let _610_v33;
          _610_v33 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),_609_v32);
          _610_v33 = (_610_v33).update(new _dafny.CodePoint('h'.codePointAt(0)), _609_v32);
        }
        let _611_v34;
        _611_v34 = _dafny.Map.Empty.slice().updateUnsafe(true,_572_v1);
        let _612_v35;
        _612_v35 = _dafny.Seq.of(_611_v34, _611_v34, _module.__default.fm13(globalState));
        _612_v35 = _612_v35;
        let _613_v36;
        let _nw102 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _613_v36 = _nw102;
        let _614_v37;
        let _nw103 = Array((new BigNumber(4)).toNumber());
        _nw103[(_dafny.ZERO).toNumber()] = _613_v36;
        _nw103[(_dafny.ONE).toNumber()] = _613_v36;
        _nw103[(new BigNumber(2)).toNumber()] = _613_v36;
        _nw103[(new BigNumber(3)).toNumber()] = _613_v36;
        _614_v37 = _nw103;
        let _615_v38;
        _615_v38 = _module.D6.create_DC12(_614_v37);
        let _616_v39;
        let _nw104 = new _module.C0();
        _nw104.__ctor((_615_v38).dtor_cf20);
        _616_v39 = _nw104;
        let _rhs104 = !(!(_this.f8));
        let _rhs105 = p1;
        let _lhs62 = _this;
        let _lhs63 = _this;
        _lhs62.f8 = _rhs104;
        _lhs63.f8 = _rhs105;
      }
      let _617_v40;
      let _init9 = ((_618_v1) => function (_619_i4) {
        return _dafny.Seq.Concat((_module.D3.create_DC5(_dafny.Seq.of(_618_v1, new BigNumber((_dafny.Set.fromElements(_this.f8)).length)))).dtor_cf6, (_module.D3.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(869)), ((_620_v1) => function (_621_i5) {
  return _620_v1;
})(_618_v1)))).dtor_cf6);
      })(_572_v1);
      let _nw105 = Array((new BigNumber(20)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw105.length); _i0_9++) {
        _nw105[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _617_v40 = _nw105;
      let _index114 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_617_v40).length));
      (_617_v40)[_index114] = _module.__default.fm14(_572_v1, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), function (_622_i6) {
        return (_this).f10;
      }), _module.__default.safeIndex(new BigNumber(347), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(104)), function (_623_i6) {
        return (_this).f10;
      })).length)), (_this).f10), _this.f8, globalState);
      let _624_v41;
      _624_v41 = _dafny.MultiSet.fromElements(false);
      let _625_v42;
      _625_v42 = _dafny.Map.Empty.slice().updateUnsafe(_572_v1,_624_v41);
      let _626_v43;
      let _nw106 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _626_v43 = _nw106;
      let _627_v44;
      _627_v44 = _module.D5.create_DC10(_626_v43);
      let _628_v45;
      _628_v45 = _dafny.MultiSet.fromElements(_627_v44, _module.D5.create_DC10(_626_v43));
      let _629_v46;
      _629_v46 = _dafny.Seq.of((_572_v1).plus(new BigNumber(235)), _572_v1, new BigNumber(-143), (_572_v1).minus(new BigNumber(((((_625_v42).contains(new BigNumber((_628_v45).cardinality()))) ? ((_625_v42).get(new BigNumber((_628_v45).cardinality()))) : (_624_v41))).cardinality())), new BigNumber(388));
      let _index115 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_617_v40).length));
      (_617_v40)[_index115] = _629_v46;
      let _630_v47;
      let _nw107 = Array((new BigNumber(11)).toNumber());
      _nw107[(_dafny.ZERO).toNumber()] = p1;
      _nw107[(_dafny.ONE).toNumber()] = p1;
      _nw107[(new BigNumber(2)).toNumber()] = _this.f8;
      _nw107[(new BigNumber(3)).toNumber()] = false;
      _nw107[(new BigNumber(4)).toNumber()] = p1;
      _nw107[(new BigNumber(5)).toNumber()] = _this.f8;
      _nw107[(new BigNumber(6)).toNumber()] = _this.f8;
      _nw107[(new BigNumber(7)).toNumber()] = _this.f8;
      _nw107[(new BigNumber(8)).toNumber()] = !_dafny.areEqual((_617_v40)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_617_v40).length))], (_617_v40)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_617_v40).length))]);
      _nw107[(new BigNumber(9)).toNumber()] = (true) === ((_this).f9);
      _nw107[(new BigNumber(10)).toNumber()] = p1;
      _630_v47 = _nw107;
      let _index116 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length));
      (_630_v47)[_index116] = !_dafny.Seq.contains((_617_v40)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_617_v40).length))], _572_v1);
      let _index117 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length));
      (_630_v47)[_index117] = _dafny.Seq.contains(p0, (_this).f10);
      let _hi5 = _572_v1;
      for (let _631_i7 = _572_v1; _631_i7.isLessThan(_hi5); _631_i7 = _631_i7.plus(_dafny.ONE)) {
        let _632_v48;
        _632_v48 = _dafny.Seq.of((_630_v47)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length))], p1, !((_630_v47)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length))]), (_630_v47)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length))], (_630_v47)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length))]);
        let _633_v49;
        _633_v49 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_632_v48).length)).isLessThanOrEqualTo(_572_v1),_631_i7);
        let _634_v50;
        _634_v50 = _module.D0.create_DC1((_this).f9, _632_v48, true);
        _633_v49 = (_633_v49).update((_this).f9, _module.__default.safeDivisionInt((_this).fm2(_634_v50, globalState), new BigNumber(59)));
        let _635_v51;
        let _nw108 = Array((new BigNumber(20)).toNumber());
        _nw108[(_dafny.ZERO).toNumber()] = _626_v43;
        _nw108[(_dafny.ONE).toNumber()] = _626_v43;
        _nw108[(new BigNumber(2)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(3)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(4)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(5)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(6)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(7)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(8)).toNumber()] = (_627_v44).dtor_cf17;
        _nw108[(new BigNumber(9)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(10)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(11)).toNumber()] = (_627_v44).dtor_cf17;
        _nw108[(new BigNumber(12)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(13)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(14)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(15)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(16)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(17)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(18)).toNumber()] = _626_v43;
        _nw108[(new BigNumber(19)).toNumber()] = _626_v43;
        _635_v51 = _nw108;
        let _636_v52;
        let _nw109 = new _module.C0();
        _nw109.__ctor(_635_v51);
        _636_v52 = _nw109;
        let _637_v53;
        let _nw110 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
        _637_v53 = _nw110;
        let _638_v54;
        _638_v54 = _dafny.Map.Empty.slice().updateUnsafe(_631_i7,_this.f8);
        let _index118 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_637_v53).length));
        (_637_v53)[_index118] = _638_v54;
        let _index119 = _module.__default.safeIndex(new BigNumber(167), new BigNumber((_637_v53).length));
        (_637_v53)[_index119] = function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(796), new BigNumber(249))) {
            let _639_v55 = _compr_15;
            if (((new BigNumber(796)).isLessThanOrEqualTo(_639_v55)) && ((_639_v55).isLessThan(new BigNumber(249)))) {
              _coll15.push([(_639_v55).multipliedBy(_572_v1),(_this).f9]);
            }
          }
          return _coll15;
        }();
        r2 = _this.f8;
      }
      _572_v1 = new BigNumber(421);
      r0 = _dafny.Seq.Concat(p0, p0);
      r1 = _572_v1;
      r2 = (_630_v47)[_module.__default.safeIndex(new BigNumber(30), new BigNumber((_630_v47).length))];
      return [r0, r1, r2];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f8 = false;
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    set f8(value) {
      let _this = this;
      _this._f8 = value;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f8, f9) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return new BigNumber(((((_this.f8) ? (_module.D3.create_DC5(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("efhmg")).length), new BigNumber((_dafny.Set.fromElements(!((_this).f9), _this.f8)).length), new BigNumber((_dafny.Seq.UnicodeFromString("jpesa")).length)))) : (_module.D3.create_DC5(_dafny.Seq.of(new BigNumber(989), new BigNumber(490)))))).dtor_cf6).length);
    };
    fm0(globalState) {
      let _this = this;
      return !(new BigNumber(775)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-788), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f8,new BigNumber(718))).length))),_this.f8)).length), new BigNumber((_dafny.Seq.of(new BigNumber(319))).length))).cardinality()))));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return (_this).f9;
    };
    fm5(p0, p1, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements((_this).f9, (_this).f9, (_this).f9, (_this).f9);
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(470), new BigNumber(464))), new BigNumber(-7));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let _640_v0;
      _640_v0 = new BigNumber(195);
      let _641_v1;
      _641_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(194),_module.__default.fm7(_640_v0, _this.f8, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), ((_642_v0) => function (_643_i0) {
        return _642_v0;
      })(_640_v0))).length), globalState));
      let _644_v2;
      _644_v2 = new _dafny.CodePoint('w'.codePointAt(0));
      _641_v1 = (_641_v1).update(_640_v0, _644_v2);
      _640_v0 = (_640_v0).minus((_640_v0).plus(_640_v0));
      let _645_v3;
      _645_v3 = _dafny.Map.Empty.slice().updateUnsafe(_640_v0,_640_v0);
      let _646_v4;
      _646_v4 = _module.D3.create_DC6((_this).f9, (new BigNumber((_645_v3).length)).multipliedBy(_640_v0), _640_v0);
      let _source12 = _646_v4;
      if (_source12.is_DC6) {
        let _647___mcc_h0 = (_source12).cf7;
        let _648___mcc_h1 = (_source12).cf8;
        let _649___mcc_h2 = (_source12).cf9;
        let _650_cf9 = _649___mcc_h2;
        let _651_cf8 = _648___mcc_h1;
        let _652_cf7 = _647___mcc_h0;
        _644_v2 = (((_this).f9) ? (_644_v2) : (_644_v2));
        let _653_v5;
        let _init10 = ((_654_cf8) => function (_655_i1) {
          return !(_654_cf8).isEqualTo(_654_cf8);
        })(_651_cf8);
        let _nw111 = Array((new BigNumber(27)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw111.length); _i0_10++) {
          _nw111[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _653_v5 = _nw111;
        let _rhs106 = (_this).f9;
        let _rhs107 = (_module.__default.safeDivisionInt(_640_v0, _650_cf9)).multipliedBy(_650_cf9);
        let _rhs108 = _653_v5;
        _652_cf7 = _rhs106;
        _640_v0 = _rhs107;
        _653_v5 = _rhs108;
        _650_cf9 = (_640_v0).plus((_650_cf9).minus(_640_v0));
        _653_v5 = _653_v5;
      } else {
        let _656___mcc_h3 = (_source12).cf6;
        let _657_cf6 = _656___mcc_h3;
        let _658_v6;
        let _nw112 = Array((new BigNumber(6)).toNumber());
        _nw112[(_dafny.ZERO).toNumber()] = !(_640_v0).isEqualTo(_640_v0);
        _nw112[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw112[(new BigNumber(2)).toNumber()] = (_this.f8) === (!(_this.f8));
        _nw112[(new BigNumber(3)).toNumber()] = true;
        _nw112[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw112[(new BigNumber(5)).toNumber()] = true;
        _658_v6 = _nw112;
        let _index120 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_658_v6).length));
        (_658_v6)[_index120] = p2;
        let _index121 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_658_v6).length));
        (_658_v6)[_index121] = p2;
        _640_v0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_640_v0), _640_v0);
        let _659_v7;
        _659_v7 = _dafny.Seq.UnicodeFromString("tfsokpds");
        let _660_v8;
        _660_v8 = _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ussmtod"));
        let _661_v9;
        _661_v9 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ypcmbt"), _dafny.Seq.Concat(_659_v7, (_660_v8).dtor_cf10));
        _659_v7 = (_661_v9)[_module.__default.safeIndex(_640_v0, new BigNumber((_661_v9).length))];
        let _662_v10;
        _662_v10 = _dafny.Seq.of(_660_v8, _module.D4.create_DC7(_659_v7), _660_v8);
        let _663_v11;
        _663_v11 = _dafny.Set.fromElements(_662_v10);
        let _664_v12;
        let _nw113 = new _module.C1();
        _nw113.__ctor(_644_v2, _663_v11, (_this).f9, (_this).f9);
        _664_v12 = _nw113;
      }
      (_this).f8 = (_this).f9;
      let _665_v13;
      _665_v13 = _dafny.Map.Empty.slice().updateUnsafe(_640_v0,p2);
      let _666_v14;
      _666_v14 = _dafny.Seq.of(_665_v13, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()),p2));
      (_this).f8 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_666_v14, _666_v14), _666_v14);
      let _667_v15;
      let _init11 = function (_668_i2) {
        return _module.__default.safeModuloInt(_668_i2, new BigNumber(540));
      };
      let _nw114 = Array((new BigNumber(2)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw114.length); _i0_11++) {
        _nw114[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _667_v15 = _nw114;
      let _source13 = _module.D5.create_DC10(_667_v15);
      if (_source13.is_DC11) {
        let _669___mcc_h4 = (_source13).cf18;
        let _670___mcc_h5 = (_source13).cf19;
        let _671_cf19 = _670___mcc_h5;
        let _672_cf18 = _669___mcc_h4;
        let _673_v16;
        let _nw115 = Array((new BigNumber(29)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = p2;
        _nw115[(_dafny.ONE).toNumber()] = !(!(_this.f8) || ((_this).f9));
        _nw115[(new BigNumber(2)).toNumber()] = p2;
        _nw115[(new BigNumber(3)).toNumber()] = p2;
        _nw115[(new BigNumber(4)).toNumber()] = p2;
        _nw115[(new BigNumber(5)).toNumber()] = p2;
        _nw115[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw115[(new BigNumber(7)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(8)).toNumber()] = p2;
        _nw115[(new BigNumber(9)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(10)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus(_640_v0)).isEqualTo((_dafny.ZERO).minus(_672_cf18));
        _nw115[(new BigNumber(12)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(13)).toNumber()] = (_this).f9;
        _nw115[(new BigNumber(14)).toNumber()] = (_this).f9;
        _nw115[(new BigNumber(15)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(16)).toNumber()] = ((_this).fm6(_672_cf18, _672_cf18, globalState)).isLessThanOrEqualTo(_672_cf18);
        _nw115[(new BigNumber(17)).toNumber()] = p2;
        _nw115[(new BigNumber(18)).toNumber()] = !((_this).f9) || (_this.f8);
        _nw115[(new BigNumber(19)).toNumber()] = p2;
        _nw115[(new BigNumber(20)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(21)).toNumber()] = p2;
        _nw115[(new BigNumber(22)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(23)).toNumber()] = !(_this.f8);
        _nw115[(new BigNumber(24)).toNumber()] = p2;
        _nw115[(new BigNumber(25)).toNumber()] = p2;
        _nw115[(new BigNumber(26)).toNumber()] = _this.f8;
        _nw115[(new BigNumber(27)).toNumber()] = (new BigNumber(552)).isLessThan(_672_cf18);
        _nw115[(new BigNumber(28)).toNumber()] = false;
        _673_v16 = _nw115;
        let _index122 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length));
        (_673_v16)[_index122] = p2;
        let _674_v17;
        _674_v17 = _module.D4.create_DC8(_640_v0, (_this).f9);
        let _index123 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length));
        (_673_v16)[_index123] = (_674_v17).dtor_cf12;
        let _index124 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length));
        let _rhs109 = _640_v0;
        let _rhs110 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_672_cf18).multipliedBy(_640_v0),((p2) ? ((_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))]) : (_this.f8)))).length);
        let _rhs111 = (_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))];
        let _lhs64 = _673_v16;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length));
        _672_cf18 = _rhs109;
        _640_v0 = _rhs110;
        _lhs64[_lhs65] = _rhs111;
        let _675_v18;
        _675_v18 = _dafny.MultiSet.fromElements(_640_v0, (_dafny.ZERO).minus(_672_cf18));
        let _676_v19;
        _676_v19 = _dafny.Map.Empty.slice().updateUnsafe(_672_cf18,_675_v18);
        let _677_v20;
        _677_v20 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,!((_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))]));
        let _678_v21;
        _678_v21 = _dafny.Map.Empty.slice().updateUnsafe(_677_v20,_644_v2);
        let _679_v22;
        _679_v22 = _dafny.Map.Empty.slice().updateUnsafe((_675_v18).Union((((_676_v19).contains(new BigNumber((_678_v21).length))) ? ((_676_v19).get(new BigNumber((_678_v21).length))) : (_675_v18))),_667_v15);
        _679_v22 = (_679_v22).update(_675_v18, _667_v15);
        let _680_v23;
        _680_v23 = _dafny.Set.fromElements((_640_v0).plus(_640_v0), _672_cf18, _640_v0, _672_cf18);
        let _681_v24;
        _681_v24 = _dafny.Seq.of(_this.f8);
        let _682_v25;
        _682_v25 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm15(globalState)).dtor_cf5,_681_v24);
        let _683_v26;
        _683_v26 = _dafny.Seq.of(_module.__default.safeModuloInt(_672_cf18, new BigNumber(894)));
        let _684_v27;
        _684_v27 = _dafny.Map.Empty.slice().updateUnsafe((_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))],_680_v23);
        let _685_v28;
        _685_v28 = _dafny.Map.Empty.slice().updateUnsafe(_644_v2,_672_cf18);
        let _686_v29;
        _686_v29 = _dafny.Map.Empty.slice().updateUnsafe(_672_cf18,_685_v28);
        let _687_v30;
        _687_v30 = _dafny.Seq.of(_682_v25);
        let _rhs112 = ((((_684_v27).contains((_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))])) ? ((_684_v27).get((_673_v16)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_673_v16).length))])) : (_dafny.Set.fromElements(new BigNumber((_686_v29).length))))).Intersect(_680_v23);
        let _rhs113 = (_687_v30)[_module.__default.safeIndex((_dafny.ZERO).minus(_640_v0), new BigNumber((_687_v30).length))];
        let _rhs114 = _683_v26;
        let _rhs115 = _module.__default.safeModuloInt(_640_v0, _640_v0);
        _680_v23 = _rhs112;
        _682_v25 = _rhs113;
        _683_v26 = _rhs114;
        _672_cf18 = _rhs115;
      } else {
        let _688___mcc_h6 = (_source13).cf17;
        let _689_cf17 = _688___mcc_h6;
        let _690_v31;
        _690_v31 = _dafny.MultiSet.fromElements(_640_v0);
        let _691_v32;
        _691_v32 = _dafny.MultiSet.fromElements((p2) === (_this.f8), ((_this).f9) === ((_this).f9), p2, true, !((new BigNumber((_690_v31).cardinality())).isEqualTo(_640_v0)));
        let _692_v33;
        _692_v33 = _dafny.Set.fromElements(p2);
        let _693_v34;
        _693_v34 = _module.D7.create_DC14(_692_v33);
        let _694_v35;
        _694_v35 = _dafny.Seq.of(((_693_v34).dtor_cf25).IsDisjointFrom(_692_v33));
        let _rhs116 = new BigNumber(((_690_v31).update(_640_v0, _module.__default.abs(_640_v0))).cardinality());
        let _rhs117 = _dafny.MultiSet.FromArray(_694_v35);
        _640_v0 = _rhs116;
        _691_v32 = _rhs117;
        let _695_v36;
        _695_v36 = _dafny.Seq.UnicodeFromString("tpt");
        let _index125 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length));
        (_667_v15)[_index125] = new BigNumber((_dafny.Seq.Concat(_695_v36, _695_v36)).length);
        let _index126 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length));
        let _rhs118 = _module.__default.safeDivisionInt(((false) ? (_640_v0) : (_640_v0)), _module.__default.safeModuloInt(_640_v0, _640_v0));
        let _rhs119 = _640_v0;
        let _lhs66 = _667_v15;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length));
        _640_v0 = _rhs118;
        _lhs66[_lhs67] = _rhs119;
        let _696_v37;
        _696_v37 = _module.D7.create_DC14(_692_v33);
        let _697_v38;
        _697_v38 = _module.D7.create_DC16(_696_v37);
        let _698_v39;
        _698_v39 = _dafny.Seq.of(_696_v37);
        let _699_v40;
        _699_v40 = _module.D7.create_DC16((_698_v39)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f9, _this.f8)).length), new BigNumber((_698_v39).length))]);
        _699_v40 = _699_v40;
        if (false) {
          let _index127 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length));
          (_667_v15)[_index127] = (_640_v0).multipliedBy(_640_v0);
          let _700_v41;
          _700_v41 = _dafny.Map.Empty.slice().updateUnsafe(((_667_v15)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length))]).isLessThanOrEqualTo(_640_v0),_689_cf17);
          _700_v41 = (_700_v41).update(p2, _689_cf17);
          let _index128 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length));
          (_667_v15)[_index128] = (_667_v15)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length))];
          (_this).f8 = (_this).f9;
          let _701_v42;
          _701_v42 = _dafny.Seq.of(_689_cf17);
          let _702_v43;
          _702_v43 = _dafny.Map.Empty.slice().updateUnsafe(_695_v36,_689_cf17);
          let _703_v44;
          let _nw116 = Array((new BigNumber(18)).toNumber());
          _nw116[(_dafny.ZERO).toNumber()] = _689_cf17;
          _nw116[(_dafny.ONE).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(2)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(3)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(4)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(5)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(6)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(7)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(8)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(9)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(10)).toNumber()] = (_701_v42)[_module.__default.safeIndex((_667_v15)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length))], new BigNumber((_701_v42).length))];
          _nw116[(new BigNumber(11)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(12)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(13)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(14)).toNumber()] = (((_702_v43).contains(_695_v36)) ? ((_702_v43).get(_695_v36)) : (_689_cf17));
          _nw116[(new BigNumber(15)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(16)).toNumber()] = _689_cf17;
          _nw116[(new BigNumber(17)).toNumber()] = _689_cf17;
          _703_v44 = _nw116;
          let _704_v45;
          let _nw117 = new _module.C0();
          _nw117.__ctor(_703_v44);
          _704_v45 = _nw117;
        } else {
          _640_v0 = (_667_v15)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length))];
          let _705_v46;
          _705_v46 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("u"));
          let _706_v47;
          _706_v47 = _module.D0.create_DC1(_this.f8, _694_v35, p2);
          (_this).f8 = (_this).fm1(_dafny.Seq.Concat(_695_v36, (_705_v46)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_705_v46).length))]), function (_pat_let27_0) {
            return function (_707_dt__update__tmp_h0) {
              return function (_pat_let28_0) {
                return function (_708_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1((_707_dt__update__tmp_h0).dtor_cf0, (_707_dt__update__tmp_h0).dtor_cf1, _708_dt__update_hcf2_h0);
                }(_pat_let28_0);
              }((_this).f9);
            }(_pat_let27_0);
          }(_706_v47), globalState);
          _690_v31 = _690_v31;
          let _709_v48;
          let _init12 = function (_710_i3) {
            return _this.f8;
          };
          let _nw118 = Array((new BigNumber(22)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw118.length); _i0_12++) {
            _nw118[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _709_v48 = _nw118;
          let _index129 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_709_v48).length));
          (_709_v48)[_index129] = _this.f8;
          let _index130 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_709_v48).length));
          (_709_v48)[_index130] = (_this).f9;
          let _711_v49;
          _711_v49 = _dafny.Seq.of(_module.D4.create_DC7(_695_v36));
          let _712_v50;
          _712_v50 = _dafny.Set.fromElements(_711_v49, _711_v49);
          let _713_v51;
          let _nw119 = new _module.C1();
          _nw119.__ctor(_644_v2, _712_v50, false, !(p2));
          _713_v51 = _nw119;
          let _714_v52;
          let _nw120 = Array((new BigNumber(14)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _713_v51;
          _nw120[(_dafny.ONE).toNumber()] = _713_v51;
          _nw120[(new BigNumber(2)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(3)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(4)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(5)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(6)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(7)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(8)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(9)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(10)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(11)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(12)).toNumber()] = _713_v51;
          _nw120[(new BigNumber(13)).toNumber()] = _713_v51;
          _714_v52 = _nw120;
          let _715_v53;
          _715_v53 = _dafny.Map.Empty.slice().updateUnsafe(_695_v36,_714_v52);
          let _716_v54;
          _716_v54 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.update(_695_v36, _module.__default.safeIndex((_667_v15)[_module.__default.safeIndex(new BigNumber(72), new BigNumber((_667_v15).length))], new BigNumber((_695_v36).length)), _644_v2));
          _715_v53 = (_715_v53).update(_dafny.Seq.Concat((((_716_v54).contains(p2)) ? ((_716_v54).get(p2)) : (_695_v36)), _695_v36), _714_v52);
        }
      }
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _717_v0;
      let _nw121 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _717_v0 = _nw121;
      let _718_v1;
      _718_v1 = new BigNumber(345);
      let _index131 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length));
      (_717_v0)[_index131] = _718_v1;
      let _index132 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length));
      (_717_v0)[_index132] = (_this).fm6(_718_v1, (_718_v1).plus(_718_v1), globalState);
      let _719_v2;
      _719_v2 = _module.D5.create_DC10(_717_v0);
      let _source14 = _719_v2;
      if (_source14.is_DC11) {
        let _720___mcc_h0 = (_source14).cf18;
        let _721___mcc_h1 = (_source14).cf19;
        let _722_cf19 = _721___mcc_h1;
        let _723_cf18 = _720___mcc_h0;
        let _index133 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length));
        (_717_v0)[_index133] = _718_v1;
        r1 = new BigNumber(590);
        let _724_v3;
        let _nw122 = Array((new BigNumber(13)).toNumber()).fill(false);
        _724_v3 = _nw122;
        let _index134 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_724_v3).length));
        (_724_v3)[_index134] = ((_this.f8) ? (_this.f8) : (p1));
        let _index135 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_724_v3).length));
        (_724_v3)[_index135] = (_this).f9;
        _717_v0 = _717_v0;
      } else {
        let _725___mcc_h2 = (_source14).cf17;
        let _726_cf17 = _725___mcc_h2;
        let _727_v4;
        _727_v4 = new _dafny.CodePoint('o'.codePointAt(0));
        let _728_v5;
        _728_v5 = _module.D4.create_DC7(p0);
        let _729_v6;
        _729_v6 = _dafny.Seq.of(_module.D4.create_DC7(p0), _728_v5, _728_v5, _728_v5, _728_v5);
        let _730_v7;
        let _nw123 = Array((new BigNumber(28)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = _726_cf17;
        _nw123[(_dafny.ONE).toNumber()] = _717_v0;
        _nw123[(new BigNumber(2)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(3)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(4)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(5)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(6)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(7)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(8)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(9)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(10)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(11)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(12)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(13)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(14)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(15)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(16)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(17)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(18)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(19)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(20)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(21)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(22)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(23)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(24)).toNumber()] = _726_cf17;
        _nw123[(new BigNumber(25)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(26)).toNumber()] = _717_v0;
        _nw123[(new BigNumber(27)).toNumber()] = _726_cf17;
        _730_v7 = _nw123;
        let _731_v8;
        let _nw124 = new _module.C0();
        _nw124.__ctor(_730_v7);
        _731_v8 = _nw124;
        let _732_v9;
        _732_v9 = _dafny.Map.Empty.slice().updateUnsafe(_731_v8,_this.f8);
        let _733_v10;
        _733_v10 = _dafny.Set.fromElements(_dafny.Seq.of(_728_v5), _729_v6, _module.__default.fm16(false, _718_v1, new BigNumber((_732_v9).length), globalState), _729_v6);
        let _734_v11;
        let _nw125 = new _module.C1();
        _nw125.__ctor(_727_v4, _733_v10, ((_this).f9) && (_this.f8), _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-972)), ((_735_v4) => function (_736_i0) {
          return _735_v4;
        })(_727_v4)), p0));
        _734_v11 = _nw125;
        (_this).f8 = true;
        _726_cf17 = _726_cf17;
        r1 = (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))];
      }
      let _737_v12;
      _737_v12 = _dafny.Seq.of(_717_v0, _717_v0, _717_v0, _717_v0, _717_v0);
      let _738_v13;
      _738_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_717_v0);
      let _739_v14;
      let _nw126 = Array((new BigNumber(14)).toNumber());
      _nw126[(_dafny.ZERO).toNumber()] = _717_v0;
      _nw126[(_dafny.ONE).toNumber()] = _717_v0;
      _nw126[(new BigNumber(2)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(3)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(4)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(5)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(6)).toNumber()] = (_737_v12)[_module.__default.safeIndex(_718_v1, new BigNumber((_737_v12).length))];
      _nw126[(new BigNumber(7)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(8)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(9)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(10)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(11)).toNumber()] = (((_738_v13).contains(_this.f8)) ? ((_738_v13).get(_this.f8)) : ((_719_v2).dtor_cf17));
      _nw126[(new BigNumber(12)).toNumber()] = _717_v0;
      _nw126[(new BigNumber(13)).toNumber()] = _717_v0;
      _739_v14 = _nw126;
      let _740_v15;
      let _nw127 = new _module.C0();
      _nw127.__ctor(_739_v14);
      _740_v15 = _nw127;
      let _741_v16;
      _741_v16 = _dafny.MultiSet.fromElements(p1);
      let _742_v17;
      _742_v17 = _dafny.Seq.of((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]);
      let _743_v18;
      _743_v18 = _dafny.Set.fromElements(_this.f8, (_this).f9, (_this).f9, _this.f8, p1);
      let _744_v19;
      _744_v19 = _module.D7.create_DC14(_743_v18);
      let _745_v20;
      _745_v20 = _dafny.Map.Empty.slice().updateUnsafe((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))],p1);
      let _746_v21;
      _746_v21 = _dafny.Seq.of((_this).f9);
      let _747_v22;
      _747_v22 = _module.D0.create_DC1(p1, _746_v21, (_746_v21)[_module.__default.safeIndex((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], new BigNumber((_746_v21).length))]);
      let _748_v23;
      let _nw128 = Array((new BigNumber(27)).toNumber());
      _nw128[(_dafny.ZERO).toNumber()] = (_740_v15).fm0(globalState);
      _nw128[(_dafny.ONE).toNumber()] = p1;
      _nw128[(new BigNumber(2)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(3)).toNumber()] = p1;
      _nw128[(new BigNumber(4)).toNumber()] = true;
      _nw128[(new BigNumber(5)).toNumber()] = _this.f8;
      _nw128[(new BigNumber(6)).toNumber()] = _this.f8;
      _nw128[(new BigNumber(7)).toNumber()] = p1;
      _nw128[(new BigNumber(8)).toNumber()] = p1;
      _nw128[(new BigNumber(9)).toNumber()] = p1;
      _nw128[(new BigNumber(10)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(11)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(12)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(13)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(14)).toNumber()] = p1;
      _nw128[(new BigNumber(15)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(16)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(17)).toNumber()] = _this.f8;
      _nw128[(new BigNumber(18)).toNumber()] = false;
      _nw128[(new BigNumber(19)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(20)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(21)).toNumber()] = _this.f8;
      _nw128[(new BigNumber(22)).toNumber()] = p1;
      _nw128[(new BigNumber(23)).toNumber()] = p1;
      _nw128[(new BigNumber(24)).toNumber()] = (_this).f9;
      _nw128[(new BigNumber(25)).toNumber()] = p1;
      _nw128[(new BigNumber(26)).toNumber()] = p1;
      _748_v23 = _nw128;
      let _749_v24;
      _749_v24 = _dafny.Seq.of(_748_v23, _748_v23);
      let _750_v25;
      _750_v25 = _dafny.Map.Empty.slice().updateUnsafe((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))],(_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]);
      let _751_v26;
      let _nw129 = Array((new BigNumber(26)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = (_718_v1).isLessThan((((_741_v16).contains(_this.f8)) ? ((_741_v16).get(_this.f8)) : ((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))])));
      _nw129[(_dafny.ONE).toNumber()] = _this.f8;
      _nw129[(new BigNumber(2)).toNumber()] = (_this).fm0(globalState);
      _nw129[(new BigNumber(3)).toNumber()] = (_this).f9;
      _nw129[(new BigNumber(4)).toNumber()] = (new BigNumber(558)).isLessThanOrEqualTo(_718_v1);
      _nw129[(new BigNumber(5)).toNumber()] = (_this).f9;
      _nw129[(new BigNumber(6)).toNumber()] = (_this).f9;
      _nw129[(new BigNumber(7)).toNumber()] = ((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]).isLessThanOrEqualTo(new BigNumber((_742_v17).length));
      _nw129[(new BigNumber(8)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(p1, (_this).f9)).IsDisjointFrom((_744_v19).dtor_cf25);
      _nw129[(new BigNumber(10)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(11)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(12)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(13)).toNumber()] = p1;
      _nw129[(new BigNumber(14)).toNumber()] = false;
      _nw129[(new BigNumber(15)).toNumber()] = p1;
      _nw129[(new BigNumber(16)).toNumber()] = !((_718_v1).isLessThan(new BigNumber((_745_v20).length)));
      _nw129[(new BigNumber(17)).toNumber()] = (_740_v15).fm1(p0, _747_v22, globalState);
      _nw129[(new BigNumber(18)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(19)).toNumber()] = !((_750_v25).contains(new BigNumber((_dafny.Seq.of(new BigNumber((_749_v24).length), (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))])).length)));
      _nw129[(new BigNumber(20)).toNumber()] = (_this).f9;
      _nw129[(new BigNumber(21)).toNumber()] = p1;
      _nw129[(new BigNumber(22)).toNumber()] = p1;
      _nw129[(new BigNumber(23)).toNumber()] = !(_this.f8) || ((_this).f9);
      _nw129[(new BigNumber(24)).toNumber()] = _this.f8;
      _nw129[(new BigNumber(25)).toNumber()] = p1;
      _751_v26 = _nw129;
      let _index136 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length));
      (_748_v23)[_index136] = p1;
      let _752_v27;
      _752_v27 = _dafny.Set.fromElements(_718_v1);
      let _753_v28;
      _753_v28 = _module.D8.create_DC17(_752_v27);
      let _index137 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length));
      let _rhs120 = (_dafny.Set.fromElements(new BigNumber(120))).IsProperSubsetOf(((_753_v28).dtor_cf27).Difference(_752_v27));
      let _rhs121 = _740_v15;
      let _rhs122 = (_dafny.ZERO).minus((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]);
      let _rhs123 = p1;
      let _rhs124 = !((_module.__default.safeModuloInt(_718_v1, (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))])).isLessThanOrEqualTo(new BigNumber((_742_v17).length)));
      let _lhs68 = _748_v23;
      let _lhs69 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length));
      r2 = _rhs120;
      _740_v15 = _rhs121;
      _718_v1 = _rhs122;
      _lhs68[_lhs69] = _rhs123;
      r2 = _rhs124;
      let _754_v29;
      _754_v29 = _module.D2.create_DC4((_this).f9);
      let _755_i1;
      _755_i1 = _dafny.ZERO;
      L2: {
        let _pat_let_tv22 = _748_v23;
        let _pat_let_tv23 = _748_v23;
        let _pat_let_tv24 = _748_v23;
        let _pat_let_tv25 = _748_v23;
        let _pat_let_tv26 = _748_v23;
        let _pat_let_tv27 = _748_v23;
        let _pat_let_tv28 = p1;
        let _pat_let_tv29 = p1;
        let _pat_let_tv30 = _748_v23;
        let _pat_let_tv31 = _748_v23;
        while (function (_source15) {
          if (_source15.is_DC4) {
            let _768___mcc_h3 = (_source15).cf5;
            let _769_cf5 = _768___mcc_h3;
            if ((_pat_let_tv23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_pat_let_tv22).length))]) {
              return (_module.D8.create_DC19((_this).f9)).dtor_cf31;
            } else {
              return (_module.D6.create_DC13((_pat_let_tv25)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_pat_let_tv24).length))], (_this).f9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-939)), function (_770_i2) {
  return new _dafny.CodePoint('u'.codePointAt(0));
}), _769_cf5)).dtor_cf21;
            }
          } else {
            let _771___mcc_h4 = (_source15).cf4;
            let _772_cf4 = _771___mcc_h4;
            if (!((_pat_let_tv27)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_pat_let_tv26).length))])) {
              return _pat_let_tv28;
            } else {
              return _pat_let_tv29;
            }
          }
        }(function (_pat_let29_0) {
          return function (_773_dt__update__tmp_h0) {
            return function (_pat_let30_0) {
              return function (_774_dt__update_hcf5_h0) {
                return _module.D2.create_DC4(_774_dt__update_hcf5_h0);
              }(_pat_let30_0);
            }(!((_pat_let_tv31)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_pat_let_tv30).length))]));
          }(_pat_let29_0);
        }(_754_v29))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_755_i1)) {
              break L2;
            }
            _755_i1 = (_755_i1).plus(_dafny.ONE);
            let _index138 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length));
            (_717_v0)[_index138] = new BigNumber(-765);
            (_this).f8 = p1;
            let _756_v30;
            _756_v30 = new _dafny.CodePoint('m'.codePointAt(0));
            let _757_v31;
            _757_v31 = _module.D4.create_DC7(_dafny.Seq.Create(_module.__default.abs(new BigNumber(887)), ((_758_v30) => function (_759_i4) {
  return _758_v30;
})(_756_v30)));
            let _760_v32;
            _760_v32 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_761_p0) => function (_762_i3) {
              return _module.D4.create_DC7(_761_p0);
            })(p0)), _module.__default.safeIndex((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_763_p0) => function (_764_i3) {
              return _module.D4.create_DC7(_763_p0);
            })(p0))).length)), _757_v31));
            let _765_v33;
            let _nw130 = new _module.C1();
            _nw130.__ctor(_756_v30, (_760_v32).Difference(_760_v32), (_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))], (_this).f9);
            _765_v33 = _nw130;
            _765_v33 = _765_v33;
            _756_v30 = ((!(p1) || ((_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))])) ? (_module.__default.fm7((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], (_this).f9, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), ((_766_v30) => function (_767_i5) {
              return _766_v30;
            })(_756_v30))).length), globalState)) : (((true) ? (_756_v30) : (_module.__default.fm7((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], (_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))], new BigNumber((_752_v27).length), globalState)))));
          }
        }
      }
      let _775_i6;
      _775_i6 = _dafny.ZERO;
      L3: {
        while (!(p1)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_775_i6)) {
              break L3;
            }
            _775_i6 = (_775_i6).plus(_dafny.ONE);
            let _776_v34;
            _776_v34 = _dafny.Seq.of(_746_v21);
            let _777_v35;
            _777_v35 = _module.D4.create_DC7(p0);
            let _778_v36;
            _778_v36 = _dafny.Map.Empty.slice().updateUnsafe(_776_v34,_777_v35);
            _778_v36 = (_778_v36).update(_776_v34, _module.__default.fm17((_this).f9, _this.f8, (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], globalState));
            let _779_v37;
            let _nw131 = Array((new BigNumber(26)).toNumber()).fill([]);
            _779_v37 = _nw131;
            let _index139 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_779_v37).length));
            (_779_v37)[_index139] = _751_v26;
            let _780_v38;
            let _nw132 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _780_v38 = _nw132;
            let _index140 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length));
            (_780_v38)[_index140] = _dafny.Seq.Concat(p0, _dafny.Seq.UnicodeFromString("iudnfs"));
            let _index141 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_779_v37).length));
            let _index142 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length));
            let _rhs125 = (_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))];
            let _rhs126 = _748_v23;
            let _rhs127 = p0;
            let _rhs128 = p0;
            let _lhs70 = _779_v37;
            let _lhs71 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_779_v37).length));
            let _lhs72 = _780_v38;
            let _lhs73 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length));
            r2 = _rhs125;
            _lhs70[_lhs71] = _rhs126;
            _lhs72[_lhs73] = _rhs127;
            r0 = _rhs128;
            if (((((p1) ? (true) : ((_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))]))) ? ((_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))]) : (!(_718_v1).isEqualTo((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))])))) {
              let _index143 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length));
              (_717_v0)[_index143] = ((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]).minus((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))]);
              let _index144 = _module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length));
              (_748_v23)[_index144] = (_this).fm0(globalState);
              let _781_v39;
              _781_v39 = _dafny.Set.fromElements(p0);
              _781_v39 = _781_v39;
              r1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_742_v17).length), new BigNumber((_743_v18).length))), new BigNumber(15));
              let _782_v40;
              let _nw133 = new _module.C0();
              _nw133.__ctor(_739_v14);
              _782_v40 = _nw133;
            } else {
              (_this).f8 = ((((_this).f9) ? (_743_v18) : (_743_v18))).IsSubsetOf(_743_v18);
              let _783_v41;
              _783_v41 = _dafny.Set.fromElements(_742_v17, _dafny.Seq.Concat(_742_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_784_v1) => function (_785_i7) {
                return _784_v1;
              })(_718_v1))), _dafny.Seq.of((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], _718_v1));
              let _786_v42;
              _786_v42 = new _dafny.CodePoint('h'.codePointAt(0));
              let _787_v43;
              _787_v43 = _786_v42;
              let _788_v44;
              let _nw134 = Array((new BigNumber(29)).toNumber());
              _nw134[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
              _nw134[(_dafny.ONE).toNumber()] = (((_this).f9) ? (_786_v42) : ((_787_v43)));
              _nw134[(new BigNumber(2)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(3)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(4)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(5)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(6)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(7)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(8)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
              _nw134[(new BigNumber(10)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(11)).toNumber()] = ((_780_v38)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length))])[_module.__default.safeIndex(new BigNumber(240), new BigNumber(((_780_v38)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length))]).length))];
              _nw134[(new BigNumber(12)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(13)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(14)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(15)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(16)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(17)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(18)).toNumber()] = _module.__default.fm7((_this).fm6(_718_v1, new BigNumber((_module.__default.fm18(_this.f8, (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], (_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], globalState)).cardinality()), globalState), (_this).fm1((_780_v38)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_780_v38).length))], _747_v22, globalState), _718_v1, globalState);
              _nw134[(new BigNumber(19)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(20)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(21)).toNumber()] = _module.__default.fm7((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], p1, (_this).fm2(_module.D0.create_DC1(false, _dafny.Seq.of((_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))]), _this.f8), globalState), globalState);
              _nw134[(new BigNumber(22)).toNumber()] = (_786_v42);
              _nw134[(new BigNumber(23)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(24)).toNumber()] = (_787_v43);
              _nw134[(new BigNumber(25)).toNumber()] = _786_v42;
              _nw134[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
              _nw134[(new BigNumber(27)).toNumber()] = _module.__default.fm7(new BigNumber(254), p1, (_dafny.ZERO).minus(_718_v1), globalState);
              _nw134[(new BigNumber(28)).toNumber()] = ((p1) ? (_786_v42) : (_786_v42));
              _788_v44 = _nw134;
              let _index145 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_788_v44).length));
              (_788_v44)[_index145] = _786_v42;
              let _789_v45;
              _789_v45 = _module.D8.create_DC19(p1);
              let _790_v46;
              let _nw135 = Array((new BigNumber(10)).toNumber()).fill(_module.D6.Default());
              _790_v46 = _nw135;
              let _791_v47;
              _791_v47 = _dafny.Map.Empty.slice().updateUnsafe(_789_v45,_790_v46);
              let _index146 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_788_v44).length));
              let _rhs129 = !(true);
              let _rhs130 = ((_dafny.Set.fromElements(_742_v17, _dafny.Seq.update(_742_v17, _module.__default.safeIndex((_717_v0)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_717_v0).length))], new BigNumber((_742_v17).length)), new BigNumber(858)), _742_v17)).Difference(_783_v41)).Difference((_783_v41).Intersect(_783_v41));
              let _rhs131 = new _dafny.CodePoint('t'.codePointAt(0));
              let _rhs132 = (_791_v47).update(function (_pat_let31_0) {
                return function (_792_dt__update__tmp_h1) {
                  return function (_pat_let32_0) {
                    return function (_793_dt__update_hcf31_h0) {
                      return _module.D8.create_DC19(_793_dt__update_hcf31_h0);
                    }(_pat_let32_0);
                  }((_this).f9);
                }(_pat_let31_0);
              }(_789_v45), _790_v46);
              let _lhs74 = _this;
              let _lhs75 = _788_v44;
              let _lhs76 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_788_v44).length));
              _lhs74.f8 = _rhs129;
              _783_v41 = _rhs130;
              _lhs75[_lhs76] = _rhs131;
              _791_v47 = _rhs132;
              let _794_v48;
              let _nw136 = new _module.C0();
              _nw136.__ctor((_740_v15).f12);
              _794_v48 = _nw136;
              (_this).f8 = (_752_v27).IsSubsetOf(_752_v27);
              _718_v1 = (_this).fm2(_747_v22, globalState);
            }
            _717_v0 = _717_v0;
          }
        }
      }
      let _795_v49;
      let _nw137 = new _module.C0();
      _nw137.__ctor((_740_v15).f12);
      _795_v49 = _nw137;
      r0 = p0;
      r1 = new BigNumber((_746_v21).length);
      r2 = (_748_v23)[_module.__default.safeIndex(new BigNumber(815), new BigNumber((_748_v23).length))];
      return [r0, r1, r2];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _796_v0;
      _796_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,(_this).f9);
      let _797_v1;
      _797_v1 = _module.D5.create_DC11(new BigNumber((_796_v0).length), p0);
      let _798_v2;
      _798_v2 = _module.D5.create_DC11(p1, _dafny.Seq.update((_797_v1).dtor_cf19, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber(((_797_v1).dtor_cf19).length)), _module.__default.fm7((_dafny.ZERO).minus(p1), _this.f8, p1, globalState)));
      _797_v1 = _798_v2;
      let _799_v3;
      _799_v3 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_this).f9, _this.f8));
      let _hi6 = (((_this).f9) ? (p1) : (p1));
      for (let _800_i0 = new BigNumber((_799_v3).cardinality()); _800_i0.isLessThan(_hi6); _800_i0 = _800_i0.plus(_dafny.ONE)) {
        let _801_v4;
        _801_v4 = _dafny.MultiSet.fromElements(p0, p0);
        let _802_v5;
        let _nw138 = Array((new BigNumber(9)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(p0, p0, p0, _module.__default.fm9(_this.f8, !((_this).f9), globalState)))).IsDisjointFrom(_801_v4);
        _nw138[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw138[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw138[(new BigNumber(3)).toNumber()] = (false) || (_this.f8);
        _nw138[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw138[(new BigNumber(5)).toNumber()] = _this.f8;
        _nw138[(new BigNumber(6)).toNumber()] = _this.f8;
        _nw138[(new BigNumber(7)).toNumber()] = _this.f8;
        _nw138[(new BigNumber(8)).toNumber()] = (_this).f9;
        _802_v5 = _nw138;
        let _index147 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length));
        (_802_v5)[_index147] = _this.f8;
        let _index148 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length));
        (_802_v5)[_index148] = !_dafny.areEqual(p0, _dafny.Seq.UnicodeFromString("kggm"));
        let _803_v6;
        _803_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.IsPrefixOf(p0, p0), _this.f8, _this.f8);
        let _804_v7;
        _804_v7 = _dafny.Seq.of((_802_v5)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length))]);
        let _index149 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length));
        let _rhs133 = _800_i0;
        let _rhs134 = new BigNumber((_803_v6).cardinality());
        let _rhs135 = !_dafny.areEqual(_dafny.Seq.of((_802_v5)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length))]), _dafny.Seq.Concat(_804_v7, _804_v7));
        let _lhs77 = _802_v5;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length));
        r0 = _rhs133;
        r0 = _rhs134;
        _lhs77[_lhs78] = _rhs135;
        let _805_v8;
        _805_v8 = new _dafny.CodePoint('o'.codePointAt(0));
        _805_v8 = _805_v8;
        let _index150 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_802_v5).length));
        (_802_v5)[_index150] = (_this).f9;
      }
      let _806_v9;
      _806_v9 = _dafny.MultiSet.fromElements(p1);
      let _807_i1;
      _807_i1 = _dafny.ZERO;
      L4: {
        while (((_806_v9).Difference(_806_v9)).IsSubsetOf((_806_v9).update(p1, _module.__default.abs(p1)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_807_i1)) {
              break L4;
            }
            _807_i1 = (_807_i1).plus(_dafny.ONE);
            r1 = !(_this.f8);
            let _808_v10;
            _808_v10 = _dafny.Seq.UnicodeFromString("xkd");
            let _809_v11;
            _809_v11 = _dafny.Seq.of(p1);
            let _810_v12;
            _810_v12 = _module.D3.create_DC5(_809_v11);
            let _811_v13;
            let _nw139 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
            _811_v13 = _nw139;
            let _812_v14;
            let _init13 = function (_813_i2) {
              return (_813_i2).plus(new BigNumber((_dafny.Seq.of((_this).f9, _this.f8)).length));
            };
            let _nw140 = Array((new BigNumber(25)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw140.length); _i0_13++) {
              _nw140[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _812_v14 = _nw140;
            let _814_v15;
            _814_v15 = _dafny.Seq.of(_812_v14, _812_v14);
            let _index151 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_811_v13).length));
            (_811_v13)[_index151] = _dafny.Seq.Concat(_814_v15, _814_v15);
            let _815_v16;
            _815_v16 = new _dafny.CodePoint('g'.codePointAt(0));
            let _816_v17;
            _816_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_815_v16);
            let _index152 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_811_v13).length));
            let _rhs136 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("rdheeg"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("rdheeg")).length)), (((_816_v17).contains(_this.f8)) ? ((_816_v17).get(_this.f8)) : (new _dafny.CodePoint('c'.codePointAt(0)))));
            let _rhs137 = _810_v12;
            let _rhs138 = _814_v15;
            let _rhs139 = p1;
            let _rhs140 = (_this).fm0(globalState);
            let _lhs79 = _811_v13;
            let _lhs80 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_811_v13).length));
            let _lhs81 = _this;
            _808_v10 = _rhs136;
            _810_v12 = _rhs137;
            _lhs79[_lhs80] = _rhs138;
            r0 = _rhs139;
            _lhs81.f8 = _rhs140;
            let _817_v18;
            let _init14 = ((_818_p1) => function (_819_i3) {
              return (_dafny.Map.Empty.slice().updateUnsafe(_818_p1,new BigNumber(215))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_818_p1,_818_p1));
            })(p1);
            let _nw141 = Array((new BigNumber(7)).toNumber());
            for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw141.length); _i0_14++) {
              _nw141[_i0_14] = _init14(new BigNumber(_i0_14));
            }
            _817_v18 = _nw141;
            _817_v18 = _817_v18;
            let _820_v19;
            _820_v19 = _module.D4.create_DC7(_808_v10);
            let _821_v20;
            _821_v20 = _dafny.Seq.of(_820_v19, _820_v19);
            let _822_v21;
            _822_v21 = _dafny.Set.fromElements(_821_v20, _821_v20);
            let _823_v22;
            let _nw142 = new _module.C1();
            _nw142.__ctor(_815_v16, _822_v21, _this.f8, (_this).f9);
            _823_v22 = _nw142;
            let _824_v23;
            _824_v23 = _dafny.Map.Empty.slice().updateUnsafe(_823_v22,p1);
            let _825_v24;
            _825_v24 = _dafny.Seq.of((_824_v23).Merge(_824_v23), (_824_v23).update(_823_v22, p1), _dafny.Map.Empty.slice().updateUnsafe(_823_v22,p1), _824_v23, _824_v23);
            _824_v23 = (_825_v24)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), ((_826_p1) => function (_827_i4) {
              return _826_p1;
            })(p1))).length), new BigNumber((_825_v24).length))];
          }
        }
      }
      let _828_v25;
      let _init15 = function (_829_i6) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      };
      let _nw143 = Array((new BigNumber(29)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw143.length); _i0_15++) {
        _nw143[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _828_v25 = _nw143;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_828_v25).length))) {
        let _830_i5 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_830_i5)) && ((_830_i5).isLessThan(new BigNumber((_828_v25).length))))) {
          (_828_v25)[(_830_i5)] = new _dafny.CodePoint('c'.codePointAt(0));
        }
      }
      let _831_v26;
      let _nw144 = Array((new BigNumber(15)).toNumber()).fill([]);
      _831_v26 = _nw144;
      let _832_v27;
      let _nw145 = new _module.C0();
      _nw145.__ctor(_831_v26);
      _832_v27 = _nw145;
      let _833_v28;
      _833_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_832_v27, _832_v27, _832_v27, _832_v27, _832_v27)).length),(_this).f9);
      _833_v28 = (_833_v28).update(_module.__default.safeModuloInt(p1, new BigNumber(499)), (_this).f9);
      r0 = p1;
      r0 = p1;
      r1 = true;
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f8 = false;
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
    set f8(value) {
      let _this = this;
      _this._f8 = value;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f8, f9) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return ((new BigNumber(-710)).multipliedBy(new BigNumber((_dafny.Set.fromElements(true)).length))).multipliedBy(new BigNumber(433));
    };
    fm0(globalState) {
      let _this = this;
      return (function () {
        let _coll16 = new _dafny.Set();
        for (const _compr_16 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-88)), new BigNumber((_dafny.Seq.of(_this.f8, _this.f8)).length))).Elements) {
          let _834_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-88)), new BigNumber((_dafny.Seq.of(_this.f8, _this.f8)).length)), _834_v0)) {
            _coll16.add(_module.__default.safeModuloInt(_834_v0, new BigNumber(7)));
          }
        }
        return _coll16;
      }()).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f9, _this.f8)).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f8,new BigNumber((function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(57), new BigNumber(809))) {
          let _835_v1 = _compr_17;
          if (((new BigNumber(57)).isLessThanOrEqualTo(_835_v1)) && ((_835_v1).isLessThan(new BigNumber(809)))) {
            _coll17.push([(_835_v1).multipliedBy(new BigNumber(-161)),_this.f8]);
          }
        }
        return _coll17;
      }()).length))).length), new BigNumber(944), new BigNumber(-645)));
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return (_this).f9;
    };
    fm3(globalState) {
      let _this = this;
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(976), new BigNumber((_dafny.Seq.UnicodeFromString("ceyl")).length), new BigNumber(563)))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), function (_836_i0) {
        return new BigNumber(201);
      })).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(443), new BigNumber(200), new BigNumber((function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (_dafny.MultiSet.fromElements(new BigNumber(748), new BigNumber(142))).Elements) {
          let _837_v0 = _compr_18;
          if ((_dafny.MultiSet.fromElements(new BigNumber(748), new BigNumber(142))).contains(_837_v0)) {
            _coll18.add((_837_v0).plus(new BigNumber((_dafny.Seq.of((new _dafny.CodePoint('o'.codePointAt(0))))).length)));
          }
        }
        return _coll18;
      }()).length))).length)));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let _838_v0;
      let _nw146 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
      _838_v0 = _nw146;
      let _index153 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_838_v0).length));
      (_838_v0)[_index153] = p0;
      let _839_v1;
      _839_v1 = _dafny.Seq.of(p0, _dafny.MultiSet.fromElements(_this.f8, p2, p2, p2, (_this).f9), p0, p0, p0);
      let _840_v2;
      _840_v2 = new BigNumber(475);
      let _index154 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_838_v0).length));
      (_838_v0)[_index154] = ((((_this).f9) ? (p0) : (_dafny.MultiSet.fromElements(false, _this.f8)))).Intersect((_839_v1)[_module.__default.safeIndex(_840_v2, new BigNumber((_839_v1).length))]);
      let _841_v3;
      _841_v3 = _dafny.Seq.of(_840_v2);
      (_this).f8 = !(!((new BigNumber((_841_v3).length)).isEqualTo(_840_v2)));
      let _842_v4;
      _842_v4 = _dafny.Seq.UnicodeFromString("ks");
      let _843_v5;
      _843_v5 = _module.D2.create_DC3((_841_v3)[_module.__default.safeIndex(_840_v2, new BigNumber((_841_v3).length))]);
      let _rhs141 = _module.__default.fm4(_840_v2, _840_v2, globalState);
      let _rhs142 = (_dafny.ZERO).minus(((_843_v5).dtor_cf4).multipliedBy((_dafny.ZERO).minus(_840_v2)));
      let _rhs143 = _this.f8;
      let _lhs82 = _this;
      _842_v4 = _rhs141;
      _840_v2 = _rhs142;
      _lhs82.f8 = _rhs143;
      let _844_v6;
      let _init16 = function (_845_i0) {
        return false;
      };
      let _nw147 = Array((new BigNumber(7)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw147.length); _i0_16++) {
        _nw147[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _844_v6 = _nw147;
      _844_v6 = _844_v6;
      let _index155 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_844_v6).length));
      (_844_v6)[_index155] = p2;
      let _846_v7;
      _846_v7 = _dafny.Seq.of(p2);
      let _index156 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_844_v6).length));
      (_844_v6)[_index156] = (_846_v7)[_module.__default.safeIndex((_dafny.ZERO).minus((_840_v2).plus(_840_v2)), new BigNumber((_846_v7).length))];
      let _index157 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_844_v6).length));
      let _rhs144 = (new BigNumber((_dafny.Seq.UnicodeFromString("fvnhbyxkp")).length)).isEqualTo(_840_v2);
      let _rhs145 = (_840_v2).plus(new BigNumber(637));
      let _lhs83 = _844_v6;
      let _lhs84 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_844_v6).length));
      _lhs83[_lhs84] = _rhs144;
      _840_v2 = _rhs145;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _847_v0;
      let _nw148 = new _module.C2();
      _nw148.__ctor((_this).f9, !(p1));
      _847_v0 = _nw148;
      let _848_v1;
      _848_v1 = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,!(p1));
      let _849_v2;
      _849_v2 = _dafny.Map.Empty.slice().updateUnsafe((_847_v0).f9,_this.f8);
      let _850_v3;
      _850_v3 = new BigNumber(995);
      let _851_v4;
      _851_v4 = _dafny.MultiSet.fromElements(new BigNumber((_849_v2).length), _850_v3);
      let _852_v5;
      _852_v5 = _dafny.Seq.of(_851_v4);
      let _hi7 = new BigNumber((_852_v5).length);
      for (let _853_i0 = new BigNumber((_848_v1).length); _853_i0.isLessThan(_hi7); _853_i0 = _853_i0.plus(_dafny.ONE)) {
        let _854_v6;
        _854_v6 = _dafny.Seq.of(_this.f8, true, false);
        let _855_v7;
        let _nw149 = Array((new BigNumber(9)).toNumber());
        _nw149[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw149[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw149[(new BigNumber(2)).toNumber()] = false;
        _nw149[(new BigNumber(3)).toNumber()] = (_847_v0).f9;
        _nw149[(new BigNumber(4)).toNumber()] = (_854_v6)[_module.__default.safeIndex(_850_v3, new BigNumber((_854_v6).length))];
        _nw149[(new BigNumber(5)).toNumber()] = (_this).fm0(globalState);
        _nw149[(new BigNumber(6)).toNumber()] = (((_849_v2).contains((_847_v0).f9)) ? ((_849_v2).get((_847_v0).f9)) : (p1));
        _nw149[(new BigNumber(7)).toNumber()] = p1;
        _nw149[(new BigNumber(8)).toNumber()] = true;
        _855_v7 = _nw149;
        let _856_v8;
        _856_v8 = _dafny.Map.Empty.slice().updateUnsafe(_855_v7,p0);
        _856_v8 = (_856_v8).update(_855_v7, p0);
        r1 = (_853_i0).multipliedBy(_853_i0);
        let _857_v9;
        _857_v9 = _dafny.Seq.of(new BigNumber(875), _850_v3);
        let _858_v10;
        _858_v10 = _module.D3.create_DC5(_857_v9);
        let _source16 = _858_v10;
        if (_source16.is_DC6) {
          let _859___mcc_h0 = (_source16).cf7;
          let _860___mcc_h1 = (_source16).cf8;
          let _861___mcc_h2 = (_source16).cf9;
          let _862_cf9 = _861___mcc_h2;
          let _863_cf8 = _860___mcc_h1;
          let _864_cf7 = _859___mcc_h0;
          let _865_v11;
          let _init17 = function (_866_i1) {
            return (_866_i1).plus(new BigNumber(-678));
          };
          let _nw150 = Array((new BigNumber(23)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw150.length); _i0_17++) {
            _nw150[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _865_v11 = _nw150;
          _865_v11 = _865_v11;
          let _867_v12;
          _867_v12 = new _dafny.CodePoint('u'.codePointAt(0));
          let _868_v13;
          _868_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(532),_867_v12);
          _868_v13 = (_868_v13).update(_850_v3, _867_v12);
          let _869_v14;
          _869_v14 = _dafny.Map.Empty.slice().updateUnsafe(_863_cf8,!(_this.f8));
          let _870_v15;
          _870_v15 = _module.D0.create_DC1((((_869_v14).contains(new BigNumber((p0).length))) ? ((_869_v14).get(new BigNumber((p0).length))) : (p1)), _854_v6, (_847_v0).f9);
          _863_cf8 = (_this).fm2(_870_v15, globalState);
          let _871_v16;
          _871_v16 = _dafny.MultiSet.fromElements((_847_v0).f9, (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), ((_872_v2) => function (_873_i2) {
            return _872_v2;
          })(_849_v2))).length)).isEqualTo(new BigNumber((p0).length)));
          _871_v16 = _871_v16;
        } else {
          let _874___mcc_h3 = (_source16).cf6;
          let _875_cf6 = _874___mcc_h3;
          let _876_v17;
          _876_v17 = _dafny.Map.Empty.slice().updateUnsafe(_847_v0,_857_v9);
          r2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_875_cf6, (((_876_v17).contains(_847_v0)) ? ((_876_v17).get(_847_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-723)), ((_877_i0) => function (_878_i3) {
            return (_dafny.ZERO).minus(_877_i0);
          })(_853_i0))))), _857_v9);
          _850_v3 = _module.__default.safeModuloInt(new BigNumber(((_dafny.MultiSet.fromElements(_853_i0)).Union(_851_v4)).cardinality()), _853_i0);
          let _879_v18;
          _879_v18 = _dafny.MultiSet.fromElements(p1, _this.f8, _this.f8, (_this).f9, _this.f8);
          r1 = (_dafny.ZERO).minus((_module.D2.create_DC3((_857_v9)[_module.__default.safeIndex(new BigNumber((_879_v18).cardinality()), new BigNumber((_857_v9).length))])).dtor_cf4);
          let _880_v19;
          _880_v19 = _dafny.Set.fromElements(true, (_847_v0).f9);
          let _881_v20;
          _881_v20 = _module.D7.create_DC14(_880_v19);
          let _882_v21;
          _882_v21 = _module.D7.create_DC16(_881_v20);
          _882_v21 = _882_v21;
        }
        r2 = (_this).f9;
      }
      let _883_v22;
      let _init18 = ((_884_v3) => function (_885_i4) {
        return _module.__default.safeDivisionInt(_885_i4, _884_v3);
      })(_850_v3);
      let _nw151 = Array((_dafny.ONE).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw151.length); _i0_18++) {
        _nw151[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _883_v22 = _nw151;
      let _886_v23;
      _886_v23 = _dafny.Seq.of(_883_v22, _883_v22);
      let _887_v24;
      _887_v24 = _dafny.Seq.of(_886_v23, _886_v23);
      _886_v23 = (_887_v24)[_module.__default.safeIndex(((_this.f8) ? (_850_v3) : (_850_v3)), new BigNumber((_887_v24).length))];
      r1 = new BigNumber(292);
      r1 = (_dafny.ZERO).minus(_850_v3);
      let _888_v25;
      _888_v25 = _dafny.Seq.of((_this).f9, (_847_v0).f9);
      let _889_v26;
      _889_v26 = _module.D0.create_DC1(_this.f8, _888_v25, p1);
      let _890_v27;
      let _nw152 = Array((new BigNumber(17)).toNumber());
      _nw152[(_dafny.ZERO).toNumber()] = _883_v22;
      _nw152[(_dafny.ONE).toNumber()] = _883_v22;
      _nw152[(new BigNumber(2)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(3)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(4)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(5)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(6)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(7)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(8)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(9)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(10)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(11)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(12)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(13)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(14)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(15)).toNumber()] = _883_v22;
      _nw152[(new BigNumber(16)).toNumber()] = _883_v22;
      _890_v27 = _nw152;
      let _891_v28;
      let _nw153 = new _module.C0();
      _nw153.__ctor(_890_v27);
      _891_v28 = _nw153;
      let _892_v29;
      _892_v29 = _dafny.MultiSet.fromElements(_891_v28);
      let _893_v30;
      _893_v30 = _module.D2.create_DC4(p1);
      let _894_v31;
      _894_v31 = _dafny.MultiSet.fromElements(_module.D5.create_DC10(_883_v22));
      let _895_v32;
      let _nw154 = Array((new BigNumber(11)).toNumber());
      _nw154[(_dafny.ZERO).toNumber()] = _this.f8;
      _nw154[(_dafny.ONE).toNumber()] = (((_this).f9) ? (_this.f8) : (_this.f8));
      _nw154[(new BigNumber(2)).toNumber()] = (_847_v0).f9;
      _nw154[(new BigNumber(3)).toNumber()] = (((_this).fm1(_dafny.Seq.UnicodeFromString("uk"), _889_v26, globalState)) ? (!((_this).f9)) : ((_this).fm1(p0, _889_v26, globalState)));
      _nw154[(new BigNumber(4)).toNumber()] = (_847_v0).f9;
      _nw154[(new BigNumber(5)).toNumber()] = _this.f8;
      _nw154[(new BigNumber(6)).toNumber()] = (_892_v29).IsSubsetOf(_892_v29);
      _nw154[(new BigNumber(7)).toNumber()] = (_this).f9;
      _nw154[(new BigNumber(8)).toNumber()] = !((_893_v30).dtor_cf5);
      _nw154[(new BigNumber(9)).toNumber()] = _this.f8;
      _nw154[(new BigNumber(10)).toNumber()] = !(new BigNumber((_894_v31).cardinality())).isEqualTo(new BigNumber(653));
      _895_v32 = _nw154;
      let _index158 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_895_v32).length));
      (_895_v32)[_index158] = _this.f8;
      let _index159 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_895_v32).length));
      (_895_v32)[_index159] = true;
      (_this).f8 = (_this).f9;
      r0 = p0;
      let _896_v33;
      _896_v33 = _dafny.MultiSet.fromElements(_this.f8);
      r1 = (_850_v3).minus((((_896_v33).contains(p1)) ? ((_896_v33).get(p1)) : (_850_v3)));
      r2 = (((_this).f9) === ((_891_v28).fm0(globalState))) === ((new BigNumber(154)).isLessThanOrEqualTo(_850_v3));
      return [r0, r1, r2];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _897_v0;
      _897_v0 = new BigNumber(-16);
      let _hi8 = _897_v0;
      for (let _898_i0 = _897_v0; _898_i0.isLessThan(_hi8); _898_i0 = _898_i0.plus(_dafny.ONE)) {
        let _899_v1;
        let _nw155 = Array((new BigNumber(18)).toNumber()).fill(false);
        _899_v1 = _nw155;
        let _index160 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_899_v1).length));
        (_899_v1)[_index160] = (_this).f9;
        let _index161 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_899_v1).length));
        (_899_v1)[_index161] = p0;
        let _900_v2;
        _900_v2 = _dafny.Seq.UnicodeFromString("ikpfgysl");
        let _901_v3;
        let _nw156 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _901_v3 = _nw156;
        let _902_v4;
        _902_v4 = _dafny.Map.Empty.slice().updateUnsafe(_900_v2,_901_v3);
        _902_v4 = (_902_v4).update(_900_v2, _901_v3);
        let _index162 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_899_v1).length));
        (_899_v1)[_index162] = p0;
        let _index163 = _module.__default.safeIndex(new BigNumber(297), new BigNumber((_899_v1).length));
        (_899_v1)[_index163] = (_899_v1)[_module.__default.safeIndex(new BigNumber(297), new BigNumber((_899_v1).length))];
      }
      if ((_this).f9) {
        let _903_v5;
        let _nw157 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _903_v5 = _nw157;
        let _904_v6;
        let _nw158 = Array((new BigNumber(21)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = p0;
        _nw158[(_dafny.ONE).toNumber()] = !(p0);
        _nw158[(new BigNumber(2)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(3)).toNumber()] = false;
        _nw158[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw158[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw158[(new BigNumber(6)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(7)).toNumber()] = p0;
        _nw158[(new BigNumber(8)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(9)).toNumber()] = false;
        _nw158[(new BigNumber(10)).toNumber()] = false;
        _nw158[(new BigNumber(11)).toNumber()] = p0;
        _nw158[(new BigNumber(12)).toNumber()] = (_this).f9;
        _nw158[(new BigNumber(13)).toNumber()] = !((_this).f9);
        _nw158[(new BigNumber(14)).toNumber()] = true;
        _nw158[(new BigNumber(15)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(16)).toNumber()] = (_this).f9;
        _nw158[(new BigNumber(17)).toNumber()] = !(true);
        _nw158[(new BigNumber(18)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(19)).toNumber()] = _this.f8;
        _nw158[(new BigNumber(20)).toNumber()] = _this.f8;
        _904_v6 = _nw158;
        let _905_v7;
        _905_v7 = _dafny.Seq.of(_904_v6);
        let _906_v8;
        _906_v8 = _dafny.Map.Empty.slice().updateUnsafe(_903_v5,_dafny.Seq.contains(_905_v7, _904_v6));
        let _907_v9;
        _907_v9 = _dafny.Set.fromElements(p0, p0);
        let _908_v10;
        _908_v10 = _module.D4.create_DC8(_897_v0, (_this).f9);
        let _909_v11;
        _909_v11 = _dafny.Seq.of(_908_v10, _908_v10, _908_v10, _908_v10);
        _906_v8 = (_906_v8).update(_903_v5, !(new BigNumber((_907_v9).length)).isEqualTo(new BigNumber((_909_v11).length)));
        let _910_v12;
        _910_v12 = _dafny.Map.Empty.slice().updateUnsafe(_897_v0,_897_v0);
        let _911_v13;
        _911_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_897_v0);
        let _912_v14;
        _912_v14 = _dafny.Set.fromElements((((_910_v12).contains(_897_v0)) ? ((_910_v12).get(_897_v0)) : (new BigNumber((_911_v13).length))), new BigNumber(-441), _897_v0);
        if ((_dafny.Set.fromElements(new BigNumber(498))).IsSubsetOf(_912_v14)) {
          _897_v0 = (_dafny.ZERO).minus(_897_v0);
          let _913_v15;
          _913_v15 = _dafny.Seq.UnicodeFromString("niqye");
          let _914_v16;
          _914_v16 = _dafny.Seq.of(p0, (_this).f9);
          let _915_v17;
          _915_v17 = _module.D0.create_DC1((_this).f9, _914_v16, true);
          let _916_v18;
          let _nw159 = new _module.C2();
          _nw159.__ctor(((_this).f9) || ((_this).fm1(_913_v15, _915_v17, globalState)), false);
          _916_v18 = _nw159;
          _916_v18 = _916_v18;
          let _index164 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_903_v5).length));
          (_903_v5)[_index164] = _897_v0;
          let _index165 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_903_v5).length));
          (_903_v5)[_index165] = _897_v0;
          let _917_v19;
          _917_v19 = _dafny.Seq.of((_903_v5)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_903_v5).length))]);
          _917_v19 = _917_v19;
          r0 = (_903_v5)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_903_v5).length))];
        } else {
          let _index166 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_903_v5).length));
          (_903_v5)[_index166] = _897_v0;
          let _index167 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_903_v5).length));
          (_903_v5)[_index167] = _897_v0;
          let _918_v20;
          _918_v20 = new _dafny.CodePoint('m'.codePointAt(0));
          let _919_v21;
          _919_v21 = _dafny.Seq.of(_918_v20, _918_v20);
          let _920_v23;
          let _nw160 = new _module.C1();
          _nw160.__ctor((_919_v21)[_module.__default.safeIndex((_dafny.ZERO).minus(_897_v0), new BigNumber((_919_v21).length))], function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_921_i1) {
              return _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ymfwvwo"));
            }))).Elements) {
              let _922_v22 = _compr_19;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_921_i1) {
                return _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ymfwvwo"));
              }))).contains(_922_v22)) {
                _coll19.add(_922_v22);
              }
            }
            return _coll19;
          }(), p0, !(!((_this).f9) || (_this.f8)));
          _920_v23 = _nw160;
          let _923_v24;
          let _init19 = ((_924_v5) => function (_925_i2) {
            return (_925_i2).minus((_924_v5)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_924_v5).length))]);
          })(_903_v5);
          let _nw161 = Array((new BigNumber(22)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw161.length); _i0_19++) {
            _nw161[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _923_v24 = _nw161;
          let _926_v25;
          _926_v25 = _dafny.MultiSet.fromElements((_this).f9, (_this).f9, (_this).f9, !((((_906_v8).contains(_923_v24)) ? ((_906_v8).get(_923_v24)) : (p0))), _this.f8);
          (_920_v23).m0(_926_v25, _module.D0.create_DC0(), (((_this).f9) ? (p0) : (p0)), globalState);
          r1 = ((_903_v5)[_module.__default.safeIndex(new BigNumber(68), new BigNumber((_903_v5).length))]).isLessThanOrEqualTo(_897_v0);
          let _927_v26;
          let _nw162 = Array((new BigNumber(4)).toNumber()).fill([]);
          _927_v26 = _nw162;
          let _928_v27;
          let _init20 = ((_929_v23) => function (_930_i3) {
            return (_929_v23).f10;
          })(_920_v23);
          let _nw163 = Array((new BigNumber(8)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw163.length); _i0_20++) {
            _nw163[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _928_v27 = _nw163;
          let _index168 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_927_v26).length));
          (_927_v26)[_index168] = _928_v27;
          let _931_v28;
          _931_v28 = _dafny.Seq.of(_this.f8);
          let _932_v29;
          _932_v29 = _module.D0.create_DC1((_this).f9, _931_v28, p0);
          let _index169 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_927_v26).length));
          (_927_v26)[_index169] = (((_920_v23).fm1(_919_v21, _932_v29, globalState)) ? (_928_v27) : (_928_v27));
        }
        let _933_v30;
        let _nw164 = Array((new BigNumber(14)).toNumber()).fill([]);
        _933_v30 = _nw164;
        let _934_v31;
        _934_v31 = _module.D4.create_DC9(p0, new BigNumber(-727), _904_v6, _897_v0);
        let _935_v32;
        _935_v32 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,_904_v6);
        let _936_v33;
        let _nw165 = Array((new BigNumber(29)).toNumber());
        _nw165[(_dafny.ZERO).toNumber()] = _904_v6;
        _nw165[(_dafny.ONE).toNumber()] = _904_v6;
        _nw165[(new BigNumber(2)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(3)).toNumber()] = (_934_v31).dtor_cf15;
        _nw165[(new BigNumber(4)).toNumber()] = (((_935_v32).contains((_this).f9)) ? ((_935_v32).get((_this).f9)) : (_904_v6));
        _nw165[(new BigNumber(5)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(6)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(7)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(8)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(9)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(10)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(11)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(12)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(13)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(14)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(15)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(16)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(17)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(18)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(19)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(20)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(21)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(22)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(23)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(24)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(25)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(26)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(27)).toNumber()] = _904_v6;
        _nw165[(new BigNumber(28)).toNumber()] = _904_v6;
        _936_v33 = _nw165;
        let _index170 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_933_v30).length));
        (_933_v30)[_index170] = _936_v33;
        let _937_v34;
        _937_v34 = _dafny.Seq.of(p0);
        let _938_v35;
        _938_v35 = _dafny.MultiSet.fromElements(_897_v0, _897_v0);
        let _index171 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_933_v30).length));
        let _rhs146 = _936_v33;
        let _rhs147 = new BigNumber((_937_v34).length);
        let _rhs148 = ((_938_v35).Union(_938_v35)).IsProperSubsetOf((_dafny.MultiSet.fromElements(_897_v0)).Intersect(_938_v35));
        let _lhs85 = _933_v30;
        let _lhs86 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_933_v30).length));
        let _lhs87 = _this;
        _lhs85[_lhs86] = _rhs146;
        r0 = _rhs147;
        _lhs87.f8 = _rhs148;
        r0 = _897_v0;
        r1 = (_this).f9;
      } else {
        let _939_v36;
        _939_v36 = new _dafny.CodePoint('h'.codePointAt(0));
        let _940_v37;
        _940_v37 = _dafny.Seq.UnicodeFromString("piygtl");
        let _941_v38;
        _941_v38 = _dafny.Map.Empty.slice().updateUnsafe(_939_v36,_940_v37);
        _941_v38 = (_941_v38).update(_939_v36, _dafny.Seq.Concat(_940_v37, _940_v37));
        let _942_v39;
        _942_v39 = _dafny.Seq.of(new BigNumber(-155));
        let _943_v40;
        _943_v40 = _dafny.Seq.of(p0);
        let _944_v41;
        _944_v41 = _dafny.Map.Empty.slice().updateUnsafe((_942_v39)[_module.__default.safeIndex(_897_v0, new BigNumber((_942_v39).length))],_module.D4.create_DC8(new BigNumber((_943_v40).length), _this.f8));
        let _945_v42;
        _945_v42 = _dafny.MultiSet.fromElements(p0);
        let _946_v43;
        _946_v43 = _module.D4.create_DC8(new BigNumber((_945_v42).cardinality()), p0);
        _944_v41 = (_944_v41).update(new BigNumber((_943_v40).length), _946_v43);
        _897_v0 = _897_v0;
        let _947_v44;
        let _nw166 = Array((new BigNumber(29)).toNumber()).fill(false);
        _947_v44 = _nw166;
        _947_v44 = _947_v44;
        let _948_v45;
        _948_v45 = _939_v36;
        _945_v42 = (_945_v42).Difference(_module.__default.fm19((_this).f9, _948_v45, (_942_v39)[_module.__default.safeIndex(_897_v0, new BigNumber((_942_v39).length))], globalState));
      }
      let _949_v46;
      _949_v46 = _dafny.Set.fromElements(_897_v0, _897_v0);
      let _950_v47;
      _950_v47 = _dafny.Seq.UnicodeFromString("ev");
      let _951_v48;
      _951_v48 = _dafny.Seq.of(_949_v46, _dafny.Set.fromElements(new BigNumber((_950_v47).length), _897_v0), _949_v46);
      let _952_v49;
      _952_v49 = _module.D8.create_DC17((_951_v48)[_module.__default.safeIndex(_897_v0, new BigNumber((_951_v48).length))]);
      let _source17 = _952_v49;
      if (_source17.is_DC18) {
        let _953___mcc_h0 = (_source17).cf28;
        let _954___mcc_h1 = (_source17).cf29;
        let _955___mcc_h2 = (_source17).cf30;
        let _956_cf30 = _955___mcc_h2;
        let _957_cf29 = _954___mcc_h1;
        let _958_cf28 = _953___mcc_h0;
        let _959_v50;
        _959_v50 = _dafny.Map.Empty.slice().updateUnsafe(_897_v0,(_this).f9);
        let _960_v51;
        let _nw167 = new _module.C2();
        _nw167.__ctor(_957_cf29, _957_cf29);
        _960_v51 = _nw167;
        let _961_v52;
        _961_v52 = _dafny.Set.fromElements(_950_v47, _950_v47);
        let _962_v53;
        _962_v53 = _dafny.MultiSet.fromElements(_956_cf30);
        let _963_v54;
        _963_v54 = _module.D9.create_DC20(_962_v53);
        let _964_v55;
        _964_v55 = _module.D9.create_DC20(((_963_v54).dtor_cf32).update(_956_cf30, _module.__default.abs(new BigNumber(726))));
        let _rhs149 = !((((_module.D9.create_DC20((_962_v53).update(new BigNumber((_962_v53).cardinality()), _module.__default.abs(_956_cf30)))).dtor_cf32).Union((_964_v55).dtor_cf32)).IsDisjointFrom((_962_v53).Intersect(_962_v53)));
        let _rhs150 = (_959_v50).Merge((_959_v50).update(_897_v0, (_this).f9));
        let _rhs151 = ((_957_cf29) ? (_960_v51) : (_960_v51));
        let _rhs152 = _960_v51;
        let _rhs153 = (_961_v52).Difference(_961_v52);
        let _lhs88 = _this;
        _lhs88.f8 = _rhs149;
        _959_v50 = _rhs150;
        _960_v51 = _rhs151;
        _960_v51 = _rhs152;
        _961_v52 = _rhs153;
        let _965_v56;
        _965_v56 = new _dafny.CodePoint('c'.codePointAt(0));
        let _966_v57;
        _966_v57 = _module.D4.create_DC7(_950_v47);
        let _967_v58;
        _967_v58 = _dafny.Seq.of(_966_v57);
        let _968_v59;
        _968_v59 = _dafny.Map.Empty.slice().updateUnsafe(_962_v53,_967_v58);
        let _pat_let_tv32 = _965_v56;
        let _969_v60;
        _969_v60 = _dafny.Set.fromElements(_967_v58, (((_968_v59).contains(_962_v53)) ? ((_968_v59).get(_962_v53)) : (_dafny.Seq.of(_module.D4.create_DC7(_950_v47), _966_v57, _966_v57, _966_v57, function (_pat_let33_0) {
          return function (_970_dt__update__tmp_h0) {
            return function (_pat_let34_0) {
              return function (_973_dt__update_hcf10_h0) {
                return _module.D4.create_DC7(_973_dt__update_hcf10_h0);
              }(_pat_let34_0);
            }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), ((_971_v56) => function (_972_i4) {
              return _971_v56;
            })(_pat_let_tv32)));
          }(_pat_let33_0);
        }(_module.D4.create_DC7(_950_v47))))));
        let _974_v61;
        _974_v61 = _module.D4.create_DC8(_897_v0, true);
        let _975_v62;
        let _nw168 = new _module.C1();
        _nw168.__ctor(_965_v56, _969_v60, (_this).f9, (_974_v61).dtor_cf12);
        _975_v62 = _nw168;
        _956_cf30 = _897_v0;
        let _976_v63;
        let _nw169 = Array((new BigNumber(15)).toNumber()).fill(_module.D8.Default());
        _976_v63 = _nw169;
        let _977_v64;
        let _init21 = ((_978_v0, _979_v53, _980_cf30) => function (_981_i5) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_978_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f8,new BigNumber((_979_v53).cardinality()))).length), _980_cf30, _980_cf30), _dafny.Seq.of(_978_v0));
        })(_897_v0, _962_v53, _956_cf30);
        let _nw170 = Array((_dafny.ONE).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw170.length); _i0_21++) {
          _nw170[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _977_v64 = _nw170;
        let _982_v65;
        _982_v65 = _dafny.Seq.of(_960_v51);
        let _rhs154 = ((_957_cf29) ? (_976_v63) : (_976_v63));
        let _rhs155 = _977_v64;
        let _rhs156 = _module.__default.safeDivisionInt(_897_v0, new BigNumber(((_962_v53).Intersect(_962_v53)).cardinality()));
        let _rhs157 = _956_cf30;
        let _rhs158 = new BigNumber((_982_v65).length);
        _976_v63 = _rhs154;
        _977_v64 = _rhs155;
        _897_v0 = _rhs156;
        _897_v0 = _rhs157;
        _956_cf30 = _rhs158;
      } else if (_source17.is_DC19) {
        let _983___mcc_h3 = (_source17).cf31;
        let _984_cf31 = _983___mcc_h3;
        let _985_v66;
        _985_v66 = _dafny.Seq.of(new BigNumber((_950_v47).length), _897_v0, new BigNumber((_950_v47).length), _897_v0, new BigNumber(26));
        let _986_v68;
        _986_v68 = _dafny.Seq.of(p0, true);
        let _987_v70;
        let _init22 = ((_988_v0) => function (_989_i6) {
          return (_989_i6).plus(_988_v0);
        })(_897_v0);
        let _nw171 = Array((new BigNumber(12)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw171.length); _i0_22++) {
          _nw171[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _987_v70 = _nw171;
        let _990_v71;
        _990_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f8),_dafny.Set.fromElements(_897_v0));
        let _991_v72;
        let _nw172 = Array((new BigNumber(23)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = (function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of (_985_v66).Elements) {
            let _992_v67 = _compr_20;
            if (_dafny.Seq.contains(_985_v66, _992_v67)) {
              _coll20.add(_module.__default.safeDivisionInt(_992_v67, (_module.D2.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("ieuojgwcd")).length))).dtor_cf4));
            }
          }
          return _coll20;
        }()).Intersect(_949_v46);
        _nw172[(_dafny.ONE).toNumber()] = _949_v46;
        _nw172[(new BigNumber(2)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_986_v68).length), new BigNumber((_950_v47).length));
        _nw172[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_897_v0);
        _nw172[(new BigNumber(5)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(6)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(7)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(8)).toNumber()] = (_949_v46).Difference(_949_v46);
        _nw172[(new BigNumber(9)).toNumber()] = ((_951_v48)[_module.__default.safeIndex(_897_v0, new BigNumber((_951_v48).length))]).Intersect(_949_v46);
        _nw172[(new BigNumber(10)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(11)).toNumber()] = function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(622), new BigNumber(777))) {
            let _993_v69 = _compr_21;
            if (((new BigNumber(622)).isLessThanOrEqualTo(_993_v69)) && ((_993_v69).isLessThan(new BigNumber(777)))) {
              _coll21.add((_993_v69).plus((_module.D3.create_DC6(p0, new BigNumber((_dafny.MultiSet.fromElements(p0, (_this).f9)).cardinality()), _897_v0)).dtor_cf8));
            }
          }
          return _coll21;
        }();
        _nw172[(new BigNumber(12)).toNumber()] = (_949_v46).Union(_949_v46);
        _nw172[(new BigNumber(13)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(14)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(15)).toNumber()] = (_dafny.Set.fromElements(_897_v0)).Union(_dafny.Set.fromElements(_897_v0, _897_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_987_v70,(_dafny.ZERO).minus(_897_v0))).length)));
        _nw172[(new BigNumber(16)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(17)).toNumber()] = (((_990_v71).contains(_984_cf31)) ? ((_990_v71).get(_984_cf31)) : (_949_v46));
        _nw172[(new BigNumber(18)).toNumber()] = (_949_v46).Intersect(_949_v46);
        _nw172[(new BigNumber(19)).toNumber()] = (_949_v46).Intersect(_dafny.Set.fromElements(new BigNumber(-300)));
        _nw172[(new BigNumber(20)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(21)).toNumber()] = _949_v46;
        _nw172[(new BigNumber(22)).toNumber()] = _949_v46;
        _991_v72 = _nw172;
        let _index172 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_991_v72).length));
        (_991_v72)[_index172] = _dafny.Set.fromElements(_897_v0);
        let _index173 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_991_v72).length));
        (_991_v72)[_index173] = _dafny.Set.fromElements((_897_v0).minus(_897_v0));
        let _994_v73;
        let _nw173 = Array((new BigNumber(5)).toNumber()).fill(false);
        _994_v73 = _nw173;
        _994_v73 = _994_v73;
        _994_v73 = _994_v73;
        let _index174 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_987_v70).length));
        (_987_v70)[_index174] = _897_v0;
        let _index175 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_987_v70).length));
        (_987_v70)[_index175] = _897_v0;
      } else {
        let _995___mcc_h4 = (_source17).cf27;
        let _996_cf27 = _995___mcc_h4;
        let _997_v74;
        let _init23 = ((_998_v0) => function (_999_i7) {
          return !((_module.D4.create_DC8(_998_v0, (_this).f9)).dtor_cf11).isEqualTo(_998_v0);
        })(_897_v0);
        let _nw174 = Array((new BigNumber(4)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw174.length); _i0_23++) {
          _nw174[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _997_v74 = _nw174;
        let _index176 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_997_v74).length));
        (_997_v74)[_index176] = _this.f8;
        let _1000_v75;
        _1000_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm0(globalState),(_this).f9);
        let _index177 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_997_v74).length));
        (_997_v74)[_index177] = (((_1000_v75).contains((_this).f9)) ? ((_1000_v75).get((_this).f9)) : (_this.f8));
        let _index178 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_997_v74).length));
        (_997_v74)[_index178] = p0;
        let _rhs159 = (_dafny.ZERO).minus(_897_v0);
        let _rhs160 = _897_v0;
        _897_v0 = _rhs159;
        _897_v0 = _rhs160;
        let _1001_v76;
        let _nw175 = Array((new BigNumber(16)).toNumber()).fill(_module.D9.Default());
        _1001_v76 = _nw175;
        let _1002_v77;
        _1002_v77 = _module.D4.create_DC7(_950_v47);
        let _1003_v78;
        _1003_v78 = _dafny.Seq.of(_1002_v77, function (_pat_let35_0) {
          return function (_1004_dt__update__tmp_h1) {
            return function (_pat_let36_0) {
              return function (_1006_dt__update_hcf10_h1) {
                return _module.D4.create_DC7(_1006_dt__update_hcf10_h1);
              }(_pat_let36_0);
            }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(219)), function (_1005_i8) {
              return new _dafny.CodePoint('i'.codePointAt(0));
            }));
          }(_pat_let35_0);
        }(_1002_v77));
        let _1007_v79;
        _1007_v79 = _dafny.Set.fromElements(_1003_v78);
        let _1008_v80;
        let _nw176 = new _module.C1();
        _nw176.__ctor((_950_v47)[_module.__default.safeIndex(_897_v0, new BigNumber((_950_v47).length))], _1007_v79, _this.f8, (_997_v74)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_997_v74).length))]);
        _1008_v80 = _nw176;
        let _1009_v81;
        _1009_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1008_v80,_897_v0);
        let _1010_v82;
        _1010_v82 = _module.D9.create_DC22(_897_v0, _897_v0, p0, _897_v0, _1009_v81);
        let _index179 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1001_v76).length));
        (_1001_v76)[_index179] = _1010_v82;
        let _index180 = _module.__default.safeIndex(new BigNumber(159), new BigNumber((_1001_v76).length));
        (_1001_v76)[_index180] = _module.D9.create_DC22(_897_v0, (_897_v0).multipliedBy(_897_v0), p0, new BigNumber((_module.__default.fm20(globalState)).length), (_1009_v81).Merge((_1009_v81).update(_1008_v80, _897_v0)));
      }
      let _1011_v83;
      _1011_v83 = _module.D8.create_DC19(_this.f8);
      let _pat_let_tv33 = _950_v47;
      _897_v0 = function (_source18) {
        if (_source18.is_DC18) {
          let _1012___mcc_h5 = (_source18).cf28;
          let _1013___mcc_h6 = (_source18).cf29;
          let _1014___mcc_h7 = (_source18).cf30;
          let _1015_cf30 = _1014___mcc_h7;
          let _1016_cf29 = _1013___mcc_h6;
          let _1017_cf28 = _1012___mcc_h5;
          return (_1015_cf30).multipliedBy(new BigNumber((_dafny.Seq.of(_1015_cf30, _1015_cf30)).length));
        } else if (_source18.is_DC19) {
          let _1018___mcc_h8 = (_source18).cf31;
          let _1019_cf31 = _1018___mcc_h8;
          return (_dafny.ZERO).minus(new BigNumber((_pat_let_tv33).length));
        } else {
          let _1020___mcc_h9 = (_source18).cf27;
          let _1021_cf27 = _1020___mcc_h9;
          return new BigNumber(100);
        }
      }(_1011_v83);
      let _1022_v84;
      _1022_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_897_v0);
      let _1023_v85;
      _1023_v85 = _dafny.Seq.of(_1022_v84, _module.__default.fm13(globalState), _1022_v84);
      let _1024_i9;
      _1024_i9 = _dafny.ZERO;
      L5: {
        while ((new BigNumber((_1023_v85).length)).isEqualTo(_897_v0)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1024_i9)) {
              break L5;
            }
            _1024_i9 = (_1024_i9).plus(_dafny.ONE);
            let _1025_v86;
            _1025_v86 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_897_v0), _897_v0);
            let _1026_v87;
            _1026_v87 = _dafny.Set.fromElements(_this.f8, (_1025_v86).IsSubsetOf(_1025_v86), (_this).f9, p0);
            let _1027_v88;
            _1027_v88 = _module.D0.create_DC1(p0, _dafny.Seq.of((_this).f9, (_this).f9), !(p0));
            let _rhs161 = _897_v0;
            let _rhs162 = _950_v47;
            let _rhs163 = _897_v0;
            let _rhs164 = (_this).fm2(((p0) ? (_1027_v88) : (function (_pat_let37_0) {
              return function (_1028_dt__update__tmp_h2) {
                return function (_pat_let38_0) {
                  return function (_1029_dt__update_hcf0_h0) {
                    return _module.D0.create_DC1(_1029_dt__update_hcf0_h0, (_1028_dt__update__tmp_h2).dtor_cf1, (_1028_dt__update__tmp_h2).dtor_cf2);
                  }(_pat_let38_0);
                }((_this).f9);
              }(_pat_let37_0);
            }(_1027_v88))), globalState);
            let _rhs165 = _1026_v87;
            r0 = _rhs161;
            _950_v47 = _rhs162;
            _897_v0 = _rhs163;
            _897_v0 = _rhs164;
            _1026_v87 = _rhs165;
            if (true) {
              let _1030_v89;
              _1030_v89 = _dafny.Seq.of(_this.f8);
              r1 = (_1030_v89)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_897_v0, (_dafny.ZERO).minus(_897_v0)), new BigNumber((_1030_v89).length))];
              (_this).f8 = (_this).f9;
              let _1031_v90;
              let _nw177 = Array((new BigNumber(5)).toNumber()).fill([]);
              _1031_v90 = _nw177;
              let _nw178 = Array((new BigNumber(18)).toNumber()).fill([]);
              _1031_v90 = _nw178;
              let _1032_v91;
              _1032_v91 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_module.__default.fm7(_897_v0, (_this).f9, (_this).fm2(_1027_v88, globalState), globalState));
              let _1033_v92;
              _1033_v92 = _dafny.Set.fromElements(_1032_v91);
              let _1034_v93;
              _1034_v93 = _dafny.Seq.of((_this).fm2(_1027_v88, globalState));
              let _1035_v94;
              _1035_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v92,(((_this).f9) ? (new BigNumber((_1034_v93).length)) : (_897_v0)));
              let _1036_v95;
              _1036_v95 = _dafny.Map.Empty.slice().updateUnsafe(_this.f8,true);
              _1035_v94 = (_1035_v94).update(_dafny.Set.fromElements(_1032_v91, _1032_v91, _1032_v91), new BigNumber((_1036_v95).length));
              let _1037_v96;
              let _nw179 = new _module.C0();
              _nw179.__ctor(_1031_v90);
              _1037_v96 = _nw179;
            } else {
              let _1038_v97;
              _1038_v97 = _module.D4.create_DC7(_dafny.Seq.UnicodeFromString("ki"));
              let _pat_let_tv34 = _950_v47;
              let _1039_v98;
              _1039_v98 = _dafny.Seq.of(_1038_v97, _1038_v97, function (_pat_let39_0) {
                return function (_1040_dt__update__tmp_h3) {
                  return function (_pat_let40_0) {
                    return function (_1041_dt__update_hcf10_h2) {
                      return _module.D4.create_DC7(_1041_dt__update_hcf10_h2);
                    }(_pat_let40_0);
                  }(_pat_let_tv34);
                }(_pat_let39_0);
              }(_1038_v97), _1038_v97, _1038_v97);
              let _1042_v99;
              let _nw180 = new _module.C1();
              _nw180.__ctor(new _dafny.CodePoint('q'.codePointAt(0)), _dafny.Set.fromElements(_1039_v98), (_1026_v87).equals(_1026_v87), !(_this.f8) || (p0));
              _1042_v99 = _nw180;
              let _1043_v100;
              let _nw181 = Array((new BigNumber(23)).toNumber()).fill([]);
              _1043_v100 = _nw181;
              let _1044_v101;
              let _nw182 = Array((new BigNumber(17)).toNumber());
              _nw182[(_dafny.ZERO).toNumber()] = false;
              _nw182[(_dafny.ONE).toNumber()] = _this.f8;
              _nw182[(new BigNumber(2)).toNumber()] = (_this).f9;
              _nw182[(new BigNumber(3)).toNumber()] = false;
              _nw182[(new BigNumber(4)).toNumber()] = false;
              _nw182[(new BigNumber(5)).toNumber()] = p0;
              _nw182[(new BigNumber(6)).toNumber()] = (_this).f9;
              _nw182[(new BigNumber(7)).toNumber()] = _this.f8;
              _nw182[(new BigNumber(8)).toNumber()] = (_this).f9;
              _nw182[(new BigNumber(9)).toNumber()] = _this.f8;
              _nw182[(new BigNumber(10)).toNumber()] = p0;
              _nw182[(new BigNumber(11)).toNumber()] = _this.f8;
              _nw182[(new BigNumber(12)).toNumber()] = true;
              _nw182[(new BigNumber(13)).toNumber()] = _this.f8;
              _nw182[(new BigNumber(14)).toNumber()] = (_this).f9;
              _nw182[(new BigNumber(15)).toNumber()] = p0;
              _nw182[(new BigNumber(16)).toNumber()] = (_this).f9;
              _1044_v101 = _nw182;
              let _1045_v102;
              _1045_v102 = _dafny.Seq.of(_1044_v101, _1044_v101, _1044_v101);
              let _index181 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1043_v100).length));
              (_1043_v100)[_index181] = (_1045_v102)[_module.__default.safeIndex(_897_v0, new BigNumber((_1045_v102).length))];
              let _1046_v103;
              _1046_v103 = _dafny.Seq.of(_897_v0);
              let _1047_v104;
              let _init24 = ((_1048_v0) => function (_1049_i10) {
                return (_1049_i10).multipliedBy(_1048_v0);
              })(_897_v0);
              let _nw183 = Array((new BigNumber(9)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw183.length); _i0_24++) {
                _nw183[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _1047_v104 = _nw183;
              let _1050_v105;
              let _nw184 = Array((new BigNumber(2)).toNumber());
              _nw184[(_dafny.ZERO).toNumber()] = _1047_v104;
              _nw184[(_dafny.ONE).toNumber()] = _1047_v104;
              _1050_v105 = _nw184;
              let _1051_v106;
              let _nw185 = new _module.C0();
              _nw185.__ctor(_1050_v105);
              _1051_v106 = _nw185;
              let _1052_v107;
              _1052_v107 = _dafny.Map.Empty.slice().updateUnsafe(_1051_v106,(_this).f9);
              let _1053_v108;
              _1053_v108 = _module.D4.create_DC9(p0, new BigNumber((_1052_v107).length), _1044_v101, _897_v0);
              let _1054_v109;
              _1054_v109 = _dafny.Map.Empty.slice().updateUnsafe(_897_v0,(_1053_v108).dtor_cf15);
              let _index182 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1043_v100).length));
              let _rhs166 = (((_1054_v109).contains((_897_v0).plus(_897_v0))) ? ((_1054_v109).get((_897_v0).plus(_897_v0))) : (_1044_v101));
              let _rhs167 = _1044_v101;
              let _rhs168 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-424)), ((_1055_v0) => function (_1056_i11) {
                return (_dafny.ZERO).minus(_1055_v0);
              })(_897_v0)), _1046_v103), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_897_v0)).length)));
              let _rhs169 = _950_v47;
              let _lhs89 = _1043_v100;
              let _lhs90 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_1043_v100).length));
              _lhs89[_lhs90] = _rhs166;
              _1044_v101 = _rhs167;
              _1046_v103 = _rhs168;
              _950_v47 = _rhs169;
              let _1057_v110;
              let _nw186 = new _module.C1();
              _nw186.__ctor((_1042_v99).f10, (_1042_v99).f11, (p0) && (true), !(false));
              _1057_v110 = _nw186;
              let _1058_v111;
              _1058_v111 = new _dafny.CodePoint('f'.codePointAt(0));
              let _index183 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1047_v104).length));
              (_1047_v104)[_index183] = (_1042_v99).fm2(_1027_v88, globalState);
              let _index184 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1047_v104).length));
              let _rhs170 = _1047_v104;
              let _rhs171 = _1058_v111;
              let _rhs172 = (_897_v0).minus(_897_v0);
              let _lhs91 = _1047_v104;
              let _lhs92 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1047_v104).length));
              _1047_v104 = _rhs170;
              _1058_v111 = _rhs171;
              _lhs91[_lhs92] = _rhs172;
              let _index185 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_1047_v104).length));
              (_1047_v104)[_index185] = _module.__default.safeModuloInt((_1042_v99).fm2(_1027_v88, globalState), (_1047_v104)[_module.__default.safeIndex(new BigNumber(9), new BigNumber((_1047_v104).length))]);
            }
            let _1059_v112;
            let _nw187 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
            _1059_v112 = _nw187;
            let _nw188 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
            _1059_v112 = _nw188;
            let _1060_v113;
            _1060_v113 = new _dafny.CodePoint('e'.codePointAt(0));
            let _1061_v114;
            _1061_v114 = _dafny.Seq.of((_this).f9, p0, (_this).f9);
            r1 = (p0) && ((_this).fm1(_dafny.Seq.update(_950_v47, _module.__default.safeIndex(_897_v0, new BigNumber((_950_v47).length)), _1060_v113), _module.D0.create_DC1(false, _1061_v114, _this.f8), globalState));
          }
        }
      }
      let _1062_v115;
      _1062_v115 = _dafny.MultiSet.fromElements(_1022_v84, _1022_v84);
      if ((_1062_v115).IsProperSubsetOf(_1062_v115)) {
        (_this).f8 = !(_this.f8);
        let _1063_v116;
        _1063_v116 = _dafny.Seq.of(p0, _this.f8, (_this).f9);
        let _1064_v117;
        _1064_v117 = _module.D0.create_DC1(true, _1063_v116, p0);
        r0 = (_dafny.ZERO).minus((_this).fm2(_1064_v117, globalState));
        let _1065_v118;
        let _nw189 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _1065_v118 = _nw189;
        let _1066_v119;
        _1066_v119 = _dafny.Seq.of(_1065_v118, _1065_v118, _1065_v118, _1065_v118, _1065_v118);
        let _1067_v120;
        let _nw190 = Array((new BigNumber(17)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = _1065_v118;
        _nw190[(_dafny.ONE).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(2)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(3)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(4)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(5)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(6)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(7)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(8)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(9)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(10)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(11)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(12)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(13)).toNumber()] = (_1066_v119)[_module.__default.safeIndex((_dafny.ZERO).minus(_897_v0), new BigNumber((_1066_v119).length))];
        _nw190[(new BigNumber(14)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(15)).toNumber()] = _1065_v118;
        _nw190[(new BigNumber(16)).toNumber()] = _1065_v118;
        _1067_v120 = _nw190;
        let _1068_v121;
        let _nw191 = new _module.C0();
        _nw191.__ctor(_1067_v120);
        _1068_v121 = _nw191;
        let _1069_v122;
        _1069_v122 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f8, (_this).f9),_this.f8);
        _1069_v122 = _1069_v122;
        let _1070_v123;
        _1070_v123 = new _dafny.CodePoint('n'.codePointAt(0));
        let _rhs173 = (_this).fm2(_1064_v117, globalState);
        let _rhs174 = ((p0) ? (_1070_v123) : (_1070_v123));
        let _rhs175 = !(_this.f8) || (_this.f8);
        let _rhs176 = (((_897_v0).isEqualTo(_897_v0)) ? (_897_v0) : (_897_v0));
        let _rhs177 = (_this).f9;
        let _lhs93 = _this;
        _897_v0 = _rhs173;
        _1070_v123 = _rhs174;
        r1 = _rhs175;
        r0 = _rhs176;
        _lhs93.f8 = _rhs177;
      } else {
        let _1071_v124;
        let _nw192 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _1071_v124 = _nw192;
        let _index186 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1071_v124).length));
        (_1071_v124)[_index186] = _897_v0;
        let _1072_v125;
        _1072_v125 = _dafny.Seq.of((_this).f9, (_this).f9);
        let _1073_v126;
        _1073_v126 = _module.D0.create_DC1(_this.f8, _1072_v125, (_this).f9);
        let _1074_v127;
        _1074_v127 = _module.D0.create_DC1((_this).fm1(_dafny.Seq.UnicodeFromString("k"), _1073_v126, globalState), _1072_v125, p0);
        let _index187 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_1071_v124).length));
        (_1071_v124)[_index187] = (_this).fm2(function (_pat_let41_0) {
          return function (_1075_dt__update__tmp_h4) {
            return function (_pat_let42_0) {
              return function (_1076_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_1075_dt__update__tmp_h4).dtor_cf0, (_1075_dt__update__tmp_h4).dtor_cf1, _1076_dt__update_hcf2_h0);
              }(_pat_let42_0);
            }((_this).f9);
          }(_pat_let41_0);
        }(_1074_v127), globalState);
        let _1077_v128;
        let _nw193 = Array((new BigNumber(24)).toNumber()).fill([]);
        _1077_v128 = _nw193;
        let _1078_v129;
        let _nw194 = new _module.C0();
        _nw194.__ctor(_1077_v128);
        _1078_v129 = _nw194;
        let _1079_v130;
        _1079_v130 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1078_v129);
        _1079_v130 = (_1079_v130).update(p0, _1078_v129);
        (_this).f8 = _this.f8;
        r0 = (_1071_v124)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_1071_v124).length))];
        let _1080_v131;
        _1080_v131 = _dafny.Set.fromElements((_this).f9);
        _897_v0 = new BigNumber(((_1080_v131).Union(_1080_v131)).length);
      }
      let _1081_v132;
      _1081_v132 = new _dafny.CodePoint('h'.codePointAt(0));
      r0 = _module.__default.safeDivisionInt(_897_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(967)), function (_1082_i12) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _module.__default.safeIndex(new BigNumber(-265), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(967)), function (_1083_i12) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length)), _1081_v132)).length));
      r1 = _this.f8;
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
