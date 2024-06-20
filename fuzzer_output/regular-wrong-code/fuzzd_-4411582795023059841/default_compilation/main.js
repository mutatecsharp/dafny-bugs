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
      return (new BigNumber(-427)).multipliedBy(new BigNumber(806));
    };
    static fm1(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)))).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false, !(true), true)));
    };
    static fm2(p0, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(366)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-655))).length)), !(new BigNumber(252)).isEqualTo(new BigNumber(907)), !((_dafny.Set.fromElements(new BigNumber(424))).IsSubsetOf(function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_0_i0) {
            return (new BigNumber(-901));
          })).Elements) {
            let _1_v0 = _compr_1;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_0_i0) {
              return (new BigNumber(-901));
            }), _1_v0)) {
              _coll1.push([_module.__default.safeModuloInt(_1_v0, new BigNumber(949)),true]);
            }
          }
          return _coll1;
        }()).Keys.Elements) {
          let _2_v1 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_0_i0) {
              return (new BigNumber(-901));
            })).Elements) {
              let _1_v0 = _compr_2;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_0_i0) {
                return (new BigNumber(-901));
              }), _1_v0)) {
                _coll2.push([_module.__default.safeModuloInt(_1_v0, new BigNumber(949)),true]);
              }
            }
            return _coll2;
          }()).contains(_2_v1)) {
            _coll0.add((_2_v1).minus(new BigNumber(((_module.D1.create_DC3(_dafny.Seq.of(new BigNumber(57), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("i")).length),true)).length))), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length)))).dtor_cf4).length)));
          }
        }
        return _coll0;
      }())), !_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0))), new _dafny.CodePoint('y'.codePointAt(0))), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(403),new BigNumber(754))).contains(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("wcdmhndul"), _dafny.Seq.UnicodeFromString("cf"), _dafny.Seq.UnicodeFromString("vdotcjy"))).length)));
    };
    static fm5(p0, globalState) {
      return _module.D1.create_DC1(!(true));
    };
    static fm6(globalState) {
      return _module.D1.create_DC3(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-626)), _dafny.Seq.of(new BigNumber(-882), new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("alqfd")).length),!(!(!(true)))))).length), new BigNumber((_dafny.Seq.UnicodeFromString("rdcr")).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-427))))), (new BigNumber((_dafny.Seq.UnicodeFromString("vc")).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ikfvgwhey")).length))));
    };
    static fm7(globalState) {
      return _module.D1.create_DC2((new BigNumber((_dafny.Seq.UnicodeFromString("isjymm")).length)).plus(new BigNumber(687)), (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(381), new BigNumber(900), new BigNumber((_dafny.Set.fromElements(!(false), false)).length))));
    };
    static fm8(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_3_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      });
    };
    static fm9(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),false)).Merge((_module.D5.create_DC12(_dafny.Map.Empty.slice().updateUnsafe(false,false))).dtor_cf14);
    };
    static fm10(globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(679), new BigNumber(756), new BigNumber(9))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("apyq")).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber(272))).length), new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))))).Elements) {
          let _4_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))), _4_v0)) {
            _coll3.push([_4_v0,true]);
          }
        }
        return _coll3;
      }()).length)))).Intersect(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("st")).length))).length)),true)).Keys.Elements) {
          let _5_v1 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("st")).length))).length)),true)).contains(_5_v1)) {
            _coll4.add(_module.__default.safeModuloInt(_5_v1, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))));
          }
        }
        return _coll4;
      }());
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('t'.codePointAt(0));
    };
    static fm12(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(364), new BigNumber((_dafny.Set.fromElements(true, !(true))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-803))).cardinality())), new BigNumber(-364)), _dafny.Seq.of(new BigNumber(290), new BigNumber(-495))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)))).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-600), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(929)))).length))).length), new BigNumber(616))));
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _hi0 = new BigNumber(-648);
      for (let _6_i0 = p1; _6_i0.isLessThan(_hi0); _6_i0 = _6_i0.plus(_dafny.ONE)) {
        let _7_v0;
        _7_v0 = _dafny.Seq.of(p1, _module.__default.fm0(globalState), _6_i0);
        r2 = (_7_v0)[_module.__default.safeIndex(_6_i0, new BigNumber((_7_v0).length))];
        if (!(_6_i0).isEqualTo(p1)) {
          let _8_v1;
          _8_v1 = new _dafny.CodePoint('h'.codePointAt(0));
          _8_v1 = _8_v1;
          let _9_v2;
          let _nw0 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _9_v2 = _nw0;
          let _10_v3;
          let _nw1 = Array((new BigNumber(13)).toNumber());
          _nw1[(_dafny.ZERO).toNumber()] = _9_v2;
          _nw1[(_dafny.ONE).toNumber()] = _9_v2;
          _nw1[(new BigNumber(2)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(3)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(4)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(5)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(6)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(7)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(8)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(9)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(10)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(11)).toNumber()] = _9_v2;
          _nw1[(new BigNumber(12)).toNumber()] = _9_v2;
          _10_v3 = _nw1;
          let _11_v4;
          _11_v4 = false;
          let _12_v5;
          _12_v5 = _dafny.Map.Empty.slice().updateUnsafe(_10_v3,_11_v4);
          _12_v5 = (_12_v5).update(_10_v3, _11_v4);
          r0 = _11_v4;
          let _13_v6;
          let _init0 = ((_14_v4) => function (_15_i1) {
            return (_15_i1).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_14_v4,_14_v4)).length)));
          })(_11_v4);
          let _nw2 = Array((new BigNumber(17)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
            _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _13_v6 = _nw2;
          let _index0 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_13_v6).length));
          (_13_v6)[_index0] = (_dafny.ZERO).minus(new BigNumber(-266));
          let _16_v7;
          _16_v7 = _dafny.Seq.of(_11_v4);
          let _17_v8;
          _17_v8 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p1));
          let _18_v9;
          _18_v9 = _dafny.Map.Empty.slice().updateUnsafe(_13_v6,_17_v8);
          let _19_v10;
          _19_v10 = _dafny.Seq.of(_11_v4, (_11_v4) === (_11_v4), (_dafny.MultiSet.fromElements(new BigNumber((_16_v7).length))).IsDisjointFrom((((_18_v9).contains(_13_v6)) ? ((_18_v9).get(_13_v6)) : ((_17_v8).update(_6_i0, _module.__default.abs(_6_i0))))), _11_v4);
          let _index1 = _module.__default.safeIndex(new BigNumber(462), new BigNumber((_13_v6).length));
          (_13_v6)[_index1] = new BigNumber((_19_v10).length);
          let _20_v11;
          _20_v11 = _dafny.Map.Empty.slice().updateUnsafe(_11_v4,_17_v8);
          _20_v11 = (_20_v11).update(_11_v4, _17_v8);
        } else {
          let _21_v12;
          _21_v12 = _dafny.MultiSet.fromElements(p1, _6_i0, _6_i0);
          let _22_v13;
          _22_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_21_v12);
          let _23_v14;
          _23_v14 = new _dafny.CodePoint('f'.codePointAt(0));
          let _24_v15;
          _24_v15 = _dafny.MultiSet.fromElements(_23_v14, new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)));
          _22_v13 = (_22_v13).update(p1, _dafny.MultiSet.fromElements(p1, (((_24_v15).contains(new _dafny.CodePoint('b'.codePointAt(0)))) ? ((_24_v15).get(new _dafny.CodePoint('b'.codePointAt(0)))) : (_6_i0))));
          let _25_v16;
          let _nw3 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _25_v16 = _nw3;
          let _26_v17;
          _26_v17 = _dafny.MultiSet.fromElements(_25_v16);
          _26_v17 = (_26_v17).Union(_dafny.MultiSet.fromElements(_25_v16, _25_v16));
          let _27_v18;
          let _nw4 = new _module.C0();
          _nw4.__ctor(new _dafny.CodePoint('t'.codePointAt(0)));
          _27_v18 = _nw4;
          _27_v18 = _27_v18;
          let _28_v19;
          let _nw5 = Array((new BigNumber(21)).toNumber()).fill(false);
          _28_v19 = _nw5;
          let _29_v20;
          _29_v20 = false;
          let _30_v21;
          _30_v21 = _dafny.Map.Empty.slice().updateUnsafe(_27_v18,!(_29_v20));
          let _index2 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_28_v19).length));
          (_28_v19)[_index2] = !(_30_v21).contains(_27_v18);
          let _index3 = _module.__default.safeIndex(new BigNumber(850), new BigNumber((_28_v19).length));
          (_28_v19)[_index3] = _29_v20;
          let _31_v22;
          let _init1 = ((_32_v20, _33_v19) => function (_34_i2) {
            return (_dafny.Set.fromElements(_32_v20)).Union(_dafny.Set.fromElements((_33_v19)[_module.__default.safeIndex(new BigNumber(850), new BigNumber((_33_v19).length))], _32_v20, (_33_v19)[_module.__default.safeIndex(new BigNumber(850), new BigNumber((_33_v19).length))]));
          })(_29_v20, _28_v19);
          let _nw6 = Array((new BigNumber(19)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
            _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _31_v22 = _nw6;
          _31_v22 = _31_v22;
        }
        let _35_v23;
        _35_v23 = false;
        let _36_v24;
        _36_v24 = _dafny.Map.Empty.slice().updateUnsafe(_35_v23,p1);
        let _37_v25;
        _37_v25 = _dafny.Seq.of(_35_v23);
        let _38_v26;
        _38_v26 = _dafny.Map.Empty.slice().updateUnsafe(_6_i0,new BigNumber(-755));
        let _39_v27;
        _39_v27 = _dafny.Seq.of(_38_v26, _38_v26);
        let _40_v28;
        _40_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_39_v27).length),_7_v0);
        let _41_v29;
        let _nw7 = Array((new BigNumber(27)).toNumber());
        _nw7[(_dafny.ZERO).toNumber()] = p1;
        _nw7[(_dafny.ONE).toNumber()] = _6_i0;
        _nw7[(new BigNumber(2)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(3)).toNumber()] = p1;
        _nw7[(new BigNumber(4)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(5)).toNumber()] = (((_36_v24).contains(_35_v23)) ? ((_36_v24).get(_35_v23)) : (_6_i0));
        _nw7[(new BigNumber(6)).toNumber()] = (new BigNumber((_dafny.Set.fromElements(p1)).length)).plus(_6_i0);
        _nw7[(new BigNumber(7)).toNumber()] = (p1).minus(new BigNumber(206));
        _nw7[(new BigNumber(8)).toNumber()] = p1;
        _nw7[(new BigNumber(9)).toNumber()] = (p1).minus(p1);
        _nw7[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_37_v25, _37_v25)).length);
        _nw7[(new BigNumber(11)).toNumber()] = new BigNumber(-159);
        _nw7[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, p1));
        _nw7[(new BigNumber(13)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(14)).toNumber()] = new BigNumber(((_40_v28).update(p1, _dafny.Seq.update(_7_v0, _module.__default.safeIndex(p1, new BigNumber((_7_v0).length)), _6_i0))).length);
        _nw7[(new BigNumber(15)).toNumber()] = (new BigNumber(750)).minus(_6_i0);
        _nw7[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(p1, p1, (_dafny.ZERO).minus((_dafny.ZERO).minus(p1)))).length);
        _nw7[(new BigNumber(17)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm9(_35_v23, _35_v23, _35_v23, globalState)).length));
        _nw7[(new BigNumber(19)).toNumber()] = p1;
        _nw7[(new BigNumber(20)).toNumber()] = p1;
        _nw7[(new BigNumber(21)).toNumber()] = (_6_i0).minus(_6_i0);
        _nw7[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_6_i0);
        _nw7[(new BigNumber(23)).toNumber()] = (_6_i0).minus(new BigNumber(-539));
        _nw7[(new BigNumber(24)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(25)).toNumber()] = _6_i0;
        _nw7[(new BigNumber(26)).toNumber()] = p1;
        _41_v29 = _nw7;
        let _index4 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_41_v29).length));
        (_41_v29)[_index4] = _6_i0;
        let _42_v30;
        _42_v30 = _dafny.Map.Empty.slice().updateUnsafe(_7_v0,_35_v23);
        let _index5 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_41_v29).length));
        let _rhs0 = _6_i0;
        let _rhs1 = _42_v30;
        let _lhs0 = _41_v29;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_41_v29).length));
        _lhs0[_lhs1] = _rhs0;
        _42_v30 = _rhs1;
        let _43_v31;
        _43_v31 = new _dafny.CodePoint('t'.codePointAt(0));
        let _44_v32;
        let _nw8 = new _module.C0();
        _nw8.__ctor(_43_v31);
        _44_v32 = _nw8;
        let _45_v33;
        _45_v33 = _dafny.MultiSet.fromElements(_44_v32);
        let _46_v34;
        _46_v34 = _dafny.MultiSet.fromElements(new BigNumber((_45_v33).cardinality()), new BigNumber((_7_v0).length));
        let _47_v35;
        _47_v35 = _dafny.Map.Empty.slice().updateUnsafe(_35_v23,!((_37_v25)[_module.__default.safeIndex(new BigNumber((_46_v34).cardinality()), new BigNumber((_37_v25).length))]));
        _47_v35 = (_47_v35).update(_35_v23, _35_v23);
      }
      let _48_v36;
      _48_v36 = new _dafny.CodePoint('w'.codePointAt(0));
      _48_v36 = _48_v36;
      let _49_v37;
      _49_v37 = true;
      let _50_v38;
      _50_v38 = _dafny.Seq.UnicodeFromString("ejbib");
      let _51_v39;
      _51_v39 = _dafny.MultiSet.fromElements(p1, new BigNumber((_50_v38).length));
      let _52_v40;
      _52_v40 = _dafny.Seq.of(_49_v37);
      let _53_v41;
      _53_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,(_52_v40)[_module.__default.safeIndex(p1, new BigNumber((_52_v40).length))]);
      let _54_v42;
      _54_v42 = _dafny.Seq.of(p1, new BigNumber((_dafny.Set.fromElements(_49_v37, _49_v37, _module.__default.fm1(_dafny.Seq.of(_51_v39), _49_v37, globalState))).length), new BigNumber((p0).cardinality()), p1, new BigNumber((_53_v41).length));
      let _55_v43;
      let _nw9 = Array((new BigNumber(13)).toNumber());
      _nw9[(_dafny.ZERO).toNumber()] = p1;
      _nw9[(_dafny.ONE).toNumber()] = p1;
      _nw9[(new BigNumber(2)).toNumber()] = (p1).multipliedBy(p1);
      _nw9[(new BigNumber(3)).toNumber()] = p1;
      _nw9[(new BigNumber(4)).toNumber()] = new BigNumber((_module.__default.fm10(globalState)).length);
      _nw9[(new BigNumber(5)).toNumber()] = p1;
      _nw9[(new BigNumber(6)).toNumber()] = new BigNumber(-831);
      _nw9[(new BigNumber(7)).toNumber()] = (p1).plus(p1);
      _nw9[(new BigNumber(8)).toNumber()] = p1;
      _nw9[(new BigNumber(9)).toNumber()] = (_54_v42)[_module.__default.safeIndex(p1, new BigNumber((_54_v42).length))];
      _nw9[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(p1, (_dafny.ZERO).minus(new BigNumber((_53_v41).length)));
      _nw9[(new BigNumber(11)).toNumber()] = p1;
      _nw9[(new BigNumber(12)).toNumber()] = p1;
      _55_v43 = _nw9;
      let _index6 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length));
      (_55_v43)[_index6] = (p1).multipliedBy(p1);
      let _index7 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length));
      (_55_v43)[_index7] = p1;
      r2 = (_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))];
      let _56_v44;
      let _init2 = ((_57_v37) => function (_58_i3) {
        return _57_v37;
      })(_49_v37);
      let _nw10 = Array((new BigNumber(27)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw10.length); _i0_2++) {
        _nw10[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _56_v44 = _nw10;
      let _index8 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
      (_56_v44)[_index8] = _49_v37;
      let _index9 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length));
      let _index10 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
      let _rhs2 = _module.__default.fm1(_dafny.Seq.of(_51_v39), _49_v37, globalState);
      let _rhs3 = _52_v40;
      let _rhs4 = p1;
      let _rhs5 = p1;
      let _rhs6 = _49_v37;
      let _lhs2 = globalState;
      let _lhs3 = _55_v43;
      let _lhs4 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length));
      let _lhs5 = _56_v44;
      let _lhs6 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
      _lhs2.f0 = _rhs2;
      _52_v40 = _rhs3;
      r2 = _rhs4;
      _lhs3[_lhs4] = _rhs5;
      _lhs5[_lhs6] = _rhs6;
      let _59_v45;
      _59_v45 = _module.D1.create_DC4();
      let _pat_let_tv0 = _55_v43;
      let _pat_let_tv1 = _55_v43;
      let _pat_let_tv2 = _51_v39;
      let _pat_let_tv3 = _51_v39;
      let _pat_let_tv4 = _55_v43;
      let _pat_let_tv5 = _55_v43;
      let _pat_let_tv6 = _49_v37;
      let _pat_let_tv7 = _49_v37;
      let _pat_let_tv8 = _49_v37;
      if (function (_source0) {
        if (_source0.is_DC2) {
          let _60___mcc_h0 = (_source0).cf2;
          let _61___mcc_h1 = (_source0).cf3;
          let _62_cf3 = _61___mcc_h1;
          let _63_cf2 = _60___mcc_h0;
          return _62_cf3;
        } else if (_source0.is_DC3) {
          let _64___mcc_h2 = (_source0).cf4;
          let _65___mcc_h3 = (_source0).cf5;
          let _66_cf5 = _65___mcc_h3;
          let _67_cf4 = _64___mcc_h2;
          return (_dafny.Set.fromElements((_pat_let_tv1)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_pat_let_tv0).length))])).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber((_pat_let_tv2).cardinality()), new BigNumber((_pat_let_tv3).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("sp")).length), new BigNumber((_67_cf4).length), (_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_pat_let_tv4).length))]));
        } else if (_source0.is_DC4) {
          if (true) {
            return (_module.D1.create_DC1(_pat_let_tv6)).dtor_cf1;
          } else {
            return _pat_let_tv7;
          }
        } else {
          let _68___mcc_h4 = (_source0).cf1;
          let _69_cf1 = _68___mcc_h4;
          return _pat_let_tv8;
        }
      }(_59_v45)) {
        _56_v44 = _56_v44;
        _48_v36 = _module.__default.fm11(_dafny.Seq.Concat(_54_v42, _54_v42), p1, new BigNumber(766), (((_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))]) ? (_49_v37) : (_49_v37)), globalState);
        let _70_v46;
        let _nw11 = new _module.C0();
        _nw11.__ctor(_48_v36);
        _70_v46 = _nw11;
        let _index11 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
        (_56_v44)[_index11] = (((_70_v46).fm3(globalState)) ? (_49_v37) : (_49_v37));
        r1 = (p1).minus(new BigNumber(440));
      } else {
        let _index12 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
        (_56_v44)[_index12] = !((_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))]);
        let _71_v47;
        let _nw12 = new _module.C0();
        _nw12.__ctor(_48_v36);
        _71_v47 = _nw12;
        let _72_v48;
        _72_v48 = _dafny.Seq.of(_71_v47);
        let _73_v49;
        _73_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_72_v48, _dafny.Seq.of(_71_v47, _71_v47)),(_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))]);
        _73_v49 = (_73_v49).update(_72_v48, false);
        _49_v37 = (p1).isLessThanOrEqualTo(((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]).multipliedBy(p1));
        if (_49_v37) {
          let _index13 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
          (_56_v44)[_index13] = (_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))];
          let _74_v50;
          let _nw13 = Array((new BigNumber(23)).toNumber());
          _74_v50 = _nw13;
          let _75_v51;
          let _nw14 = new _module.C0();
          _nw14.__ctor(new _dafny.CodePoint('q'.codePointAt(0)));
          _75_v51 = _nw14;
          let _index14 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_74_v50).length));
          (_74_v50)[_index14] = _75_v51;
          let _76_v52;
          _76_v52 = _module.D2.create_DC6(_75_v51);
          let _77_v53;
          _77_v53 = _dafny.Seq.of(_50_v38);
          let _78_v54;
          _78_v54 = _module.D4.create_DC10((_77_v53)[_module.__default.safeIndex((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))], new BigNumber((_77_v53).length))]);
          let _79_v55;
          _79_v55 = _module.D4.create_DC10((_78_v54).dtor_cf11);
          let _80_v56;
          _80_v56 = _dafny.Map.Empty.slice().updateUnsafe(_49_v37,(_78_v54).dtor_cf11);
          let _81_v57;
          _81_v57 = _module.D1.create_DC1((_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))]);
          let _index15 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_74_v50).length));
          let _index16 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
          let _rhs7 = (_76_v52).dtor_cf7;
          let _rhs8 = _dafny.Seq.IsProperPrefixOf((((_80_v56).contains(_49_v37)) ? ((_80_v56).get(_49_v37)) : (_dafny.Seq.UnicodeFromString("pcctiawd"))), _module.__default.fm8(_dafny.Seq.of(p1, (_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]), globalState));
          let _rhs9 = !((_81_v57).dtor_cf1);
          let _lhs7 = _74_v50;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_74_v50).length));
          let _lhs9 = globalState;
          let _lhs10 = _56_v44;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
          _lhs7[_lhs8] = _rhs7;
          _lhs9.f0 = _rhs8;
          _lhs10[_lhs11] = _rhs9;
          let _82_v58;
          _82_v58 = _dafny.Map.Empty.slice().updateUnsafe(_59_v45,(_dafny.ZERO).minus((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]));
          let _83_v59;
          let _nw15 = Array((new BigNumber(26)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_50_v38, _module.__default.fm8(_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), ((_84_p1) => function (_85_i4) {
            return _84_p1;
          })(p1)), globalState));
          _nw15[(_dafny.ONE).toNumber()] = _50_v38;
          _nw15[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), ((_86_v36) => function (_87_i5) {
            return _86_v36;
          })(_48_v36));
          _nw15[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), ((_88_v51) => function (_89_i6) {
            return (_88_v51).f3;
          })(_75_v51)), _module.__default.safeIndex(new BigNumber((_82_v58).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(21)), ((_90_v51) => function (_91_i6) {
            return (_90_v51).f3;
          })(_75_v51))).length)), (_75_v51).f3);
          _nw15[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(103)), ((_92_v51) => function (_93_i7) {
            return (_92_v51).f3;
          })(_75_v51)), _50_v38);
          _nw15[(new BigNumber(5)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(6)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(7)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("w");
          _nw15[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_50_v38, _module.__default.safeIndex((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))], new BigNumber((_50_v38).length)), _48_v36);
          _nw15[(new BigNumber(10)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(11)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(_50_v38, _module.__default.safeIndex((_54_v42)[_module.__default.safeIndex((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))], new BigNumber((_54_v42).length))], new BigNumber((_50_v38).length)), (_75_v51).f3);
          _nw15[(new BigNumber(13)).toNumber()] = _module.__default.fm8(_54_v42, globalState);
          _nw15[(new BigNumber(14)).toNumber()] = ((_49_v37) ? (_50_v38) : (_module.__default.fm8(_module.__default.fm12(_module.__default.fm0(globalState), p1, globalState), globalState)));
          _nw15[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("jumanw");
          _nw15[(new BigNumber(16)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("f");
          _nw15[(new BigNumber(18)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("kyeeyllc");
          _nw15[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("k");
          _nw15[(new BigNumber(21)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(22)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(23)).toNumber()] = _50_v38;
          _nw15[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("hqetjr");
          _nw15[(new BigNumber(25)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), ((_94_v51) => function (_95_i8) {
            return (_94_v51).f3;
          })(_75_v51)), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(392)), ((_96_v51) => function (_97_i9) {
            return (_96_v51).f3;
          })(_75_v51))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), ((_98_v51) => function (_99_i8) {
            return (_98_v51).f3;
          })(_75_v51))).length)), _48_v36);
          _83_v59 = _nw15;
          let _index17 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_83_v59).length));
          (_83_v59)[_index17] = _dafny.Seq.UnicodeFromString("mqdbtkjg");
          let _index18 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_83_v59).length));
          (_83_v59)[_index18] = _50_v38;
          r1 = (_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))];
          let _index19 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length));
          (_56_v44)[_index19] = true;
        } else {
          r2 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_50_v38, _50_v38), _50_v38)).length);
          let _100_v60;
          let _nw16 = Array((new BigNumber(16)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = (((_56_v44)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_56_v44).length))]) ? (_71_v47) : (_71_v47));
          _nw16[(_dafny.ONE).toNumber()] = _71_v47;
          _nw16[(new BigNumber(2)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(3)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(4)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(5)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(6)).toNumber()] = ((!(_49_v37)) ? (_71_v47) : (_71_v47));
          _nw16[(new BigNumber(7)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(8)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(9)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(10)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(11)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(12)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(13)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(14)).toNumber()] = _71_v47;
          _nw16[(new BigNumber(15)).toNumber()] = _71_v47;
          _100_v60 = _nw16;
          _100_v60 = _100_v60;
          _50_v38 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), ((_101_v36) => function (_102_i10) {
            return _101_v36;
          })(_48_v36));
          let _103_v61;
          _103_v61 = _dafny.Map.Empty.slice().updateUnsafe(_49_v37,(_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]);
          let _104_v62;
          _104_v62 = _dafny.Map.Empty.slice().updateUnsafe(p1,_55_v43);
          r2 = new BigNumber(((_103_v61).Merge((_dafny.Map.Empty.slice().updateUnsafe(_49_v37,p1)).update(_49_v37, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_56_v44,new BigNumber((_104_v62).length))).length))))).length);
          r2 = (_dafny.ZERO).minus(((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]).minus((_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]));
        }
      }
      r0 = _49_v37;
      r1 = _module.__default.safeModuloInt(new BigNumber(455), (_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))]);
      r2 = (_55_v43)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_55_v43).length))];
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _105_globalState;
      let _nw17 = new _module.GlobalState();
      _nw17.__ctor(true, true, false);
      _105_globalState = _nw17;
      let _106_v0;
      _106_v0 = _dafny.Seq.UnicodeFromString("jract");
      _106_v0 = _dafny.Seq.Concat(_106_v0, _106_v0);
      let _107_v1;
      _107_v1 = new BigNumber(-27);
      let _108_v2;
      _108_v2 = _dafny.Map.Empty.slice().updateUnsafe(_107_v1,(_dafny.ZERO).minus(_107_v1));
      _108_v2 = (_108_v2).update(_107_v1, _107_v1);
      _107_v1 = _107_v1;
      let _109_v3;
      _109_v3 = true;
      (_105_globalState).f0 = _109_v3;
      (_105_globalState).f0 = true;
      if (_109_v3) {
        let _110_v4;
        let _init3 = function (_111_i0) {
          return true;
        };
        let _nw18 = Array((_dafny.ONE).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw18.length); _i0_3++) {
          _nw18[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _110_v4 = _nw18;
        let _index20 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length));
        (_110_v4)[_index20] = _109_v3;
        let _index21 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length));
        (_110_v4)[_index21] = ((_module.__default.fm0(_105_globalState)).minus(_107_v1)).isLessThanOrEqualTo(_107_v1);
        let _112_v5;
        let _init4 = ((_113_v3) => function (_114_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(_113_v3,!(_113_v3));
        })(_109_v3);
        let _nw19 = Array((new BigNumber(8)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw19.length); _i0_4++) {
          _nw19[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _112_v5 = _nw19;
        let _115_v6;
        _115_v6 = _dafny.MultiSet.fromElements(_107_v1);
        let _116_v7;
        _116_v7 = _dafny.Seq.of(_115_v6);
        let _117_v8;
        _117_v8 = _dafny.MultiSet.fromElements((_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))]);
        let _118_v9;
        _118_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_117_v8).cardinality()),(_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))]);
        let _119_v10;
        _119_v10 = _dafny.Map.Empty.slice().updateUnsafe(_112_v5,((_module.__default.fm1(_116_v7, _109_v3, _105_globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe((((_118_v9).contains(new BigNumber(531))) ? ((_118_v9).get(new BigNumber(531))) : (_109_v3)),_117_v8)) : (_dafny.Map.Empty.slice().updateUnsafe((_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))],_module.__default.fm2(_107_v1, _105_globalState)))));
        let _120_v11;
        _120_v11 = _dafny.Map.Empty.slice().updateUnsafe(true,_117_v8);
        _119_v10 = (_119_v10).update(_112_v5, _120_v11);
        let _121_v12;
        _121_v12 = _dafny.Seq.of(_107_v1);
        let _122_v13;
        _122_v13 = _dafny.Map.Empty.slice().updateUnsafe(_121_v12,_107_v1);
        let _123_v14;
        _123_v14 = _dafny.Set.fromElements(_107_v1, (_dafny.ZERO).minus(_107_v1), _107_v1);
        let _124_v15;
        _124_v15 = (_dafny.ZERO).minus(new BigNumber((_123_v14).length));
        let _125_v16;
        _125_v16 = _dafny.Map.Empty.slice().updateUnsafe(_107_v1,_118_v9);
        let _126_v17;
        _126_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,(_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))]);
        let _127_v18;
        let _nw20 = Array((new BigNumber(17)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = (new BigNumber((_106_v0).length)).multipliedBy(new BigNumber(-941));
        _nw20[(_dafny.ONE).toNumber()] = _107_v1;
        _nw20[(new BigNumber(2)).toNumber()] = _107_v1;
        _nw20[(new BigNumber(3)).toNumber()] = (_107_v1).multipliedBy(_107_v1);
        _nw20[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_122_v13).length), _107_v1);
        _nw20[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_107_v1, _107_v1);
        _nw20[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_107_v1);
        _nw20[(new BigNumber(7)).toNumber()] = new BigNumber(784);
        _nw20[(new BigNumber(8)).toNumber()] = _107_v1;
        _nw20[(new BigNumber(9)).toNumber()] = _107_v1;
        _nw20[(new BigNumber(10)).toNumber()] = (_124_v15);
        _nw20[(new BigNumber(11)).toNumber()] = new BigNumber(-619);
        _nw20[(new BigNumber(12)).toNumber()] = _107_v1;
        _nw20[(new BigNumber(13)).toNumber()] = (_107_v1).plus(new BigNumber((_125_v16).length));
        _nw20[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((((_115_v6).contains(_107_v1)) ? ((_115_v6).get(_107_v1)) : (_107_v1)));
        _nw20[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(_107_v1, new BigNumber((_126_v17).length));
        _nw20[(new BigNumber(16)).toNumber()] = _107_v1;
        _127_v18 = _nw20;
        let _index22 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length));
        (_127_v18)[_index22] = (_107_v1).minus(new BigNumber(888));
        let _index23 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length));
        let _rhs10 = ((((_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))]) ? (_107_v1) : (new BigNumber((_115_v6).cardinality())))).multipliedBy(_107_v1);
        let _rhs11 = ((((_dafny.ZERO).minus(_107_v1)).isLessThanOrEqualTo(_107_v1)) ? (_126_v17) : ((_dafny.Map.Empty.slice().updateUnsafe(!((_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))]),(_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))])).update((_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))], (_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))])));
        let _rhs12 = (new BigNumber(90)).isLessThanOrEqualTo((_124_v15));
        let _lhs12 = _127_v18;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length));
        let _lhs14 = _105_globalState;
        _lhs12[_lhs13] = _rhs10;
        _126_v17 = _rhs11;
        _lhs14.f0 = _rhs12;
        let _index24 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length));
        let _rhs13 = _module.__default.fm0(_105_globalState);
        let _rhs14 = (_110_v4)[_module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length))];
        let _rhs15 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_127_v18)[_module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length))], (_127_v18)[_module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length))]), new BigNumber(772));
        let _lhs15 = _127_v18;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_127_v18).length));
        let _lhs17 = _110_v4;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_110_v4).length));
        _lhs15[_lhs16] = _rhs13;
        _lhs17[_lhs18] = _rhs14;
        _107_v1 = _rhs15;
        let _128_v19;
        _128_v19 = new _dafny.CodePoint('c'.codePointAt(0));
        let _129_v20;
        let _nw21 = new _module.C0();
        _nw21.__ctor(_128_v19);
        _129_v20 = _nw21;
      } else {
        let _130_v21;
        let _nw22 = Array((new BigNumber(3)).toNumber()).fill(false);
        _130_v21 = _nw22;
        let _index26 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length));
        (_130_v21)[_index26] = _109_v3;
        let _index27 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length));
        let _rhs16 = _109_v3;
        let _rhs17 = _module.__default.fm0(_105_globalState);
        let _lhs19 = _130_v21;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length));
        _lhs19[_lhs20] = _rhs16;
        _107_v1 = _rhs17;
        let _131_v22;
        _131_v22 = _dafny.Seq.of(_109_v3);
        let _source1 = _module.__default.fm5(new BigNumber((_131_v22).length), _105_globalState);
        if (_source1.is_DC2) {
          let _132___mcc_h0 = (_source1).cf2;
          let _133___mcc_h1 = (_source1).cf3;
          let _134_cf3 = _133___mcc_h1;
          let _135_cf2 = _132___mcc_h0;
          let _136_v23;
          _136_v23 = _dafny.Set.fromElements(_134_cf3, !(false), _134_cf3);
          let _137_v24;
          _137_v24 = _dafny.Set.fromElements(_130_v21, _130_v21);
          let _138_v25;
          _138_v25 = _dafny.MultiSet.fromElements(new BigNumber((_136_v23).length), new BigNumber((_137_v24).length));
          _134_cf3 = (((_dafny.MultiSet.FromArray(_dafny.Seq.of(_107_v1, _107_v1))).IsSubsetOf(_138_v25)) ? ((_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))]) : (_134_cf3));
          _135_cf2 = (_107_v1).multipliedBy(_135_cf2);
          let _139_v26;
          let _nw23 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _139_v26 = _nw23;
          let _140_v27;
          _140_v27 = _module.D1.create_DC1(_134_cf3);
          let _141_v28;
          _141_v28 = new _dafny.CodePoint('o'.codePointAt(0));
          let _142_v29;
          let _nw24 = new _module.C0();
          _nw24.__ctor(_141_v28);
          _142_v29 = _nw24;
          let _143_v30;
          _143_v30 = _dafny.Map.Empty.slice().updateUnsafe(_140_v27,_142_v29);
          let _index28 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_139_v26).length));
          (_139_v26)[_index28] = new BigNumber((_143_v30).length);
          let _index29 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_139_v26).length));
          (_139_v26)[_index29] = _module.__default.safeDivisionInt(_107_v1, _135_cf2);
          (_105_globalState).f0 = (_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))];
        } else if (_source1.is_DC3) {
          let _144___mcc_h2 = (_source1).cf4;
          let _145___mcc_h3 = (_source1).cf5;
          let _146_cf5 = _145___mcc_h3;
          let _147_cf4 = _144___mcc_h2;
          (_105_globalState).f0 = _109_v3;
          _109_v3 = !(_109_v3);
          _107_v1 = new BigNumber(988);
          _146_cf5 = new BigNumber(-855);
        } else if (_source1.is_DC4) {
          let _148_v31;
          _148_v31 = _dafny.MultiSet.fromElements(_109_v3);
          let _149_v32;
          let _150_v33;
          let _151_v34;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = _module.__default.m0(_148_v31, _107_v1, _105_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _149_v32 = _out0;
          _150_v33 = _out1;
          _151_v34 = _out2;
          _106_v0 = _dafny.Seq.Concat(_106_v0, _106_v0);
          let _152_v35;
          _152_v35 = new _dafny.CodePoint('v'.codePointAt(0));
          let _153_v36;
          let _nw25 = new _module.C0();
          _nw25.__ctor(_152_v35);
          _153_v36 = _nw25;
          let _154_v37;
          _154_v37 = _dafny.Set.fromElements(_151_v34, new BigNumber(308));
          let _155_v38;
          _155_v38 = _dafny.MultiSet.fromElements(new BigNumber((_154_v37).length));
          let _156_v39;
          _156_v39 = _dafny.Seq.of(_155_v38);
          let _157_v40;
          _157_v40 = _dafny.Set.fromElements((_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))], _module.__default.fm1(_156_v39, (_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))], _105_globalState));
          let _index30 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length));
          let _rhs18 = _153_v36;
          let _rhs19 = ((_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))]) === (!(_157_v40).contains(!(_109_v3)));
          let _lhs21 = _130_v21;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length));
          _153_v36 = _rhs18;
          _lhs21[_lhs22] = _rhs19;
          let _158_v41;
          let _nw26 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _158_v41 = _nw26;
          let _index31 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_158_v41).length));
          (_158_v41)[_index31] = _151_v34;
          let _index32 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_158_v41).length));
          let _rhs20 = _152_v35;
          let _rhs21 = _150_v33;
          let _lhs23 = _158_v41;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_158_v41).length));
          _152_v35 = _rhs20;
          _lhs23[_lhs24] = _rhs21;
        } else {
          let _159___mcc_h4 = (_source1).cf1;
          let _160_cf1 = _159___mcc_h4;
          let _161_v42;
          _161_v42 = _dafny.Seq.of(_107_v1, _107_v1, _107_v1, _107_v1, _107_v1);
          let _162_v43;
          _162_v43 = _dafny.Set.fromElements(_dafny.Seq.update(_161_v42, _module.__default.safeIndex(_107_v1, new BigNumber((_161_v42).length)), _107_v1), _161_v42);
          (_105_globalState).f0 = !(_162_v43).equals(_162_v43);
          let _163_v44;
          let _nw27 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _163_v44 = _nw27;
          _163_v44 = (((_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))]) ? (_163_v44) : (_163_v44));
          let _164_v45;
          _164_v45 = _dafny.MultiSet.fromElements(_109_v3, _160_cf1);
          let _165_v46;
          _165_v46 = _module.D2.create_DC5(_164_v45);
          let _166_v47;
          let _167_v48;
          let _168_v49;
          let _out3;
          let _out4;
          let _out5;
          let _outcollector1 = _module.__default.m0(((_165_v46).dtor_cf6).Difference(_164_v45), _107_v1, _105_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _166_v47 = _out3;
          _167_v48 = _out4;
          _168_v49 = _out5;
          let _169_v50;
          _169_v50 = _module.D1.create_DC2(_107_v1, _160_cf1);
          let _rhs22 = _130_v21;
          let _rhs23 = _169_v50;
          _130_v21 = _rhs22;
          _169_v50 = _rhs23;
        }
        _107_v1 = (_dafny.ZERO).minus(_107_v1);
        let _170_v51;
        _170_v51 = _module.D1.create_DC2(new BigNumber((_106_v0).length), (_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))]);
        let _171_v52;
        _171_v52 = _dafny.Map.Empty.slice().updateUnsafe(_107_v1,_dafny.Seq.of(new BigNumber(-327), (_170_v51).dtor_cf2));
        let _172_v54;
        _172_v54 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),_107_v1);
        let _173_v55;
        _173_v55 = _dafny.Seq.of(_131_v22);
        let _174_v56;
        _174_v56 = _dafny.Seq.of(new BigNumber(((_173_v55)[_module.__default.safeIndex(new BigNumber((_106_v0).length), new BigNumber((_173_v55).length))]).length), _107_v1);
        _171_v52 = (_171_v52).update(new BigNumber((function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_172_v54).Keys.Elements) {
            let _175_v53 = _compr_5;
            if ((_172_v54).contains(_175_v53)) {
              _coll5.push([_175_v53,_109_v3]);
            }
          }
          return _coll5;
        }()).length), _174_v56);
        let _176_v57;
        _176_v57 = _dafny.Map.Empty.slice().updateUnsafe(_107_v1,(_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))]);
        let _177_v58;
        _177_v58 = _dafny.MultiSet.fromElements(_109_v3, (_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))], (_130_v21)[_module.__default.safeIndex(new BigNumber(869), new BigNumber((_130_v21).length))], (((_176_v57).contains((_module.__default.fm6(_105_globalState)).dtor_cf5)) ? ((_176_v57).get((_module.__default.fm6(_105_globalState)).dtor_cf5)) : (_109_v3)));
        let _178_v59;
        let _179_v60;
        let _180_v61;
        let _out6;
        let _out7;
        let _out8;
        let _outcollector2 = _module.__default.m0((_177_v58).Union((_177_v58).update(_109_v3, _module.__default.abs(_107_v1))), _module.__default.fm0(_105_globalState), _105_globalState);
        _out6 = _outcollector2[0];
        _out7 = _outcollector2[1];
        _out8 = _outcollector2[2];
        _178_v59 = _out6;
        _179_v60 = _out7;
        _180_v61 = _out8;
      }
      let _hi1 = _107_v1;
      for (let _181_i2 = _107_v1; _181_i2.isLessThan(_hi1); _181_i2 = _181_i2.plus(_dafny.ONE)) {
        let _182_v62;
        _182_v62 = _dafny.MultiSet.fromElements(_109_v3);
        let _183_v63;
        let _184_v64;
        let _185_v65;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector3 = _module.__default.m0(_182_v62, _107_v1, _105_globalState);
        _out9 = _outcollector3[0];
        _out10 = _outcollector3[1];
        _out11 = _outcollector3[2];
        _183_v63 = _out9;
        _184_v64 = _out10;
        _185_v65 = _out11;
        let _186_v66;
        _186_v66 = new _dafny.CodePoint('i'.codePointAt(0));
        let _187_v67;
        let _nw28 = new _module.C0();
        _nw28.__ctor(_186_v66);
        _187_v67 = _nw28;
        let _188_v68;
        _188_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(_187_v67),_module.__default.fm7(_105_globalState));
        let _189_v70;
        _189_v70 = _dafny.Seq.of(_183_v63);
        let _190_v71;
        let _nw29 = Array((new BigNumber(21)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = (_182_v62).Difference(_dafny.MultiSet.fromElements(_183_v63));
        _nw29[(_dafny.ONE).toNumber()] = _module.__default.fm2(new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_dafny.Set.fromElements(_185_v65)).Elements) {
            let _191_v69 = _compr_6;
            if ((_dafny.Set.fromElements(_185_v65)).contains(_191_v69)) {
              _coll6.push([(_191_v69).multipliedBy(new BigNumber((_106_v0).length)),_183_v63]);
            }
          }
          return _coll6;
        }()).length), _105_globalState);
        _nw29[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_189_v70);
        _nw29[(new BigNumber(3)).toNumber()] = (_182_v62).update(_109_v3, _module.__default.abs(_184_v64));
        _nw29[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_183_v63, _109_v3, _109_v3, _109_v3, _183_v63);
        _nw29[(new BigNumber(5)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(6)).toNumber()] = (_module.__default.fm2(_107_v1, _105_globalState)).Intersect(_182_v62);
        _nw29[(new BigNumber(7)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(8)).toNumber()] = (_182_v62).Difference(_182_v62);
        _nw29[(new BigNumber(9)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(10)).toNumber()] = _module.__default.fm2((_dafny.ZERO).minus(new BigNumber((_106_v0).length)), _105_globalState);
        _nw29[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.FromArray(_189_v70);
        _nw29[(new BigNumber(12)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(13)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(14)).toNumber()] = (_182_v62).Difference(_182_v62);
        _nw29[(new BigNumber(15)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(16)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(17)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(_109_v3);
        _nw29[(new BigNumber(19)).toNumber()] = _182_v62;
        _nw29[(new BigNumber(20)).toNumber()] = _182_v62;
        _190_v71 = _nw29;
        let _index33 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_190_v71).length));
        (_190_v71)[_index33] = _182_v62;
        let _192_v72;
        let _nw30 = new _module.C0();
        _nw30.__ctor(_186_v66);
        _192_v72 = _nw30;
        let _193_v73;
        _193_v73 = _module.D2.create_DC6(_187_v67);
        let _194_v74;
        _194_v74 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_184_v64),_192_v72);
        let _195_v75;
        _195_v75 = _module.D1.create_DC2(new BigNumber((_194_v74).length), true);
        let _196_v76;
        _196_v76 = _dafny.Set.fromElements(_184_v64, _107_v1);
        let _index34 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_190_v71).length));
        let _rhs24 = (_184_v64).minus(_184_v64);
        let _rhs25 = (_188_v68).Merge(_dafny.Map.Empty.slice().updateUnsafe(_193_v73,_195_v75));
        let _rhs26 = _module.__default.fm2(_184_v64, _105_globalState);
        let _rhs27 = ((_dafny.ZERO).minus((_185_v65).plus(new BigNumber((_196_v76).length)))).isLessThan(_185_v65);
        let _rhs28 = _192_v72;
        let _lhs25 = _190_v71;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_190_v71).length));
        let _lhs27 = _105_globalState;
        _107_v1 = _rhs24;
        _188_v68 = _rhs25;
        _lhs25[_lhs26] = _rhs26;
        _lhs27.f0 = _rhs27;
        _192_v72 = _rhs28;
        let _197_v77;
        let _nw31 = Array((new BigNumber(13)).toNumber()).fill(false);
        _197_v77 = _nw31;
        let _index35 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_197_v77).length));
        (_197_v77)[_index35] = !(_109_v3);
        let _198_v78;
        _198_v78 = _dafny.Seq.of(_107_v1);
        let _199_v79;
        _199_v79 = _module.D1.create_DC3(_198_v78, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_192_v72)).length)));
        let _200_v80;
        _200_v80 = _dafny.Map.Empty.slice().updateUnsafe(_109_v3,new BigNumber((_dafny.Seq.UnicodeFromString("dqkshokld")).length));
        let _201_v81;
        _201_v81 = _dafny.Map.Empty.slice().updateUnsafe(_187_v67,_109_v3);
        let _202_v82;
        _202_v82 = _dafny.Map.Empty.slice().updateUnsafe(_187_v67,_107_v1);
        let _index36 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_197_v77).length));
        let _rhs29 = (_199_v79).dtor_cf5;
        let _rhs30 = _183_v63;
        let _rhs31 = _dafny.Seq.update(_dafny.Seq.of((_189_v70)[_module.__default.safeIndex((((_200_v80).contains(_183_v63)) ? ((_200_v80).get(_183_v63)) : (new BigNumber((_106_v0).length))), new BigNumber((_189_v70).length))], _183_v63), _module.__default.safeIndex(((_183_v63) ? (_181_i2) : (new BigNumber((_201_v81).length))), new BigNumber((_dafny.Seq.of((_189_v70)[_module.__default.safeIndex((((_200_v80).contains(_183_v63)) ? ((_200_v80).get(_183_v63)) : (new BigNumber((_106_v0).length))), new BigNumber((_189_v70).length))], _183_v63)).length)), !(!((_202_v82).contains(_187_v67))));
        let _rhs32 = !(_184_v64).isEqualTo(new BigNumber(282));
        let _lhs28 = _197_v77;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_197_v77).length));
        let _lhs30 = _105_globalState;
        _185_v65 = _rhs29;
        _lhs28[_lhs29] = _rhs30;
        _189_v70 = _rhs31;
        _lhs30.f0 = _rhs32;
        let _203_v83;
        let _nw32 = new _module.C0();
        _nw32.__ctor(_186_v66);
        _203_v83 = _nw32;
      }
      let _204_v84;
      _204_v84 = new _dafny.CodePoint('m'.codePointAt(0));
      let _205_v85;
      let _init5 = ((_206_v3) => function (_207_i3) {
        return !(_206_v3);
      })(_109_v3);
      let _nw33 = Array((new BigNumber(6)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
        _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _205_v85 = _nw33;
      let _index37 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
      (_205_v85)[_index37] = false;
      let _208_v86;
      _208_v86 = _dafny.MultiSet.fromElements(_109_v3);
      let _209_v87;
      _209_v87 = _dafny.Set.fromElements((((_208_v86).contains(_109_v3)) ? ((_208_v86).get(_109_v3)) : (_107_v1)));
      let _210_v88;
      _210_v88 = _dafny.Seq.of(_209_v87, _209_v87, _209_v87);
      let _211_v89;
      _211_v89 = _dafny.Set.fromElements(_109_v3);
      let _index38 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
      let _rhs33 = _106_v0;
      let _rhs34 = !(_109_v3);
      let _rhs35 = ((_109_v3) ? ((_209_v87).IsSubsetOf((_210_v88)[_module.__default.safeIndex(_107_v1, new BigNumber((_210_v88).length))])) : ((_211_v89).IsProperSubsetOf(_211_v89)));
      let _rhs36 = _204_v84;
      let _rhs37 = !(!(_109_v3) || (_109_v3));
      let _lhs31 = _105_globalState;
      let _lhs32 = _205_v85;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
      _106_v0 = _rhs33;
      _109_v3 = _rhs34;
      _lhs31.f0 = _rhs35;
      _204_v84 = _rhs36;
      _lhs32[_lhs33] = _rhs37;
      let _212_v90;
      _212_v90 = _module.D1.create_DC2(_107_v1, true);
      let _213_v91;
      let _214_v92;
      let _215_v93;
      let _out12;
      let _out13;
      let _out14;
      let _outcollector4 = _module.__default.m0((_module.__default.fm2(_107_v1, _105_globalState)).update((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))], _module.__default.abs(_107_v1)), (_212_v90).dtor_cf2, _105_globalState);
      _out12 = _outcollector4[0];
      _out13 = _outcollector4[1];
      _out14 = _outcollector4[2];
      _213_v91 = _out12;
      _214_v92 = _out13;
      _215_v93 = _out14;
      if (_109_v3) {
        let _216_v94;
        _216_v94 = _dafny.Map.Empty.slice().updateUnsafe(_214_v92,(_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]);
        let _217_v95;
        let _init6 = ((_218_v93) => function (_219_i4) {
          return (_219_i4).plus(_218_v93);
        })(_215_v93);
        let _nw34 = Array((new BigNumber(5)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw34.length); _i0_6++) {
          _nw34[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _217_v95 = _nw34;
        let _220_v96;
        _220_v96 = _dafny.Map.Empty.slice().updateUnsafe(_217_v95,_205_v85);
        let _221_v97;
        _221_v97 = _dafny.Seq.of(_213_v91, !(_213_v91), (((_216_v94).contains(_107_v1)) ? ((_216_v94).get(_107_v1)) : ((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))])), (_220_v96).contains(_217_v95), _213_v91);
        _213_v91 = !((_221_v97)[_module.__default.safeIndex(_215_v93, new BigNumber((_221_v97).length))]);
        let _222_v98;
        let _nw35 = new _module.C0();
        _nw35.__ctor(_204_v84);
        _222_v98 = _nw35;
        _222_v98 = _222_v98;
        let _223_v99;
        _223_v99 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(925)), ((_224_v98) => function (_225_i5) {
          return (_224_v98).f3;
        })(_222_v98))).length),_204_v84);
        let _226_v100;
        _226_v100 = _dafny.Map.Empty.slice().updateUnsafe(_109_v3,new BigNumber((_223_v99).length));
        let _227_v101;
        _227_v101 = _dafny.Map.Empty.slice().updateUnsafe((_221_v97)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_221_v97).length))],_109_v3);
        let _228_v102;
        _228_v102 = _dafny.Map.Empty.slice().updateUnsafe(_215_v93,_226_v100);
        let _229_v103;
        _229_v103 = _dafny.Seq.of(_215_v93, new BigNumber((_226_v100).length));
        let _230_v104;
        _230_v104 = _dafny.Seq.of(_226_v100, _226_v100);
        let _rhs38 = !((((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]) ? ((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]) : (!((((_227_v101).contains((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))])) ? ((_227_v101).get((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))])) : (!((_module.D1.create_DC2(_107_v1, _213_v91)).dtor_cf3)))))));
        let _rhs39 = _204_v84;
        let _rhs40 = (((((_228_v102).contains((_229_v103)[_module.__default.safeIndex(_107_v1, new BigNumber((_229_v103).length))])) ? ((_228_v102).get((_229_v103)[_module.__default.safeIndex(_107_v1, new BigNumber((_229_v103).length))])) : ((_230_v104)[_module.__default.safeIndex(_215_v93, new BigNumber((_230_v104).length))]))).update(_213_v91, _214_v92)).Merge(_226_v100);
        let _lhs34 = _105_globalState;
        _lhs34.f0 = _rhs38;
        _204_v84 = _rhs39;
        _226_v100 = _rhs40;
        _213_v91 = _109_v3;
        _227_v101 = (_227_v101).update(false, !((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]));
      } else {
        let _231_v105;
        let _232_v106;
        let _233_v107;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector5 = _module.__default.m0(_208_v86, _215_v93, _105_globalState);
        _out15 = _outcollector5[0];
        _out16 = _outcollector5[1];
        _out17 = _outcollector5[2];
        _231_v105 = _out15;
        _232_v106 = _out16;
        _233_v107 = _out17;
        let _index39 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
        (_205_v85)[_index39] = (_209_v87).IsDisjointFrom(function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(105), new BigNumber(36))) {
            let _234_v108 = _compr_7;
            if (((new BigNumber(105)).isLessThanOrEqualTo(_234_v108)) && ((_234_v108).isLessThan(new BigNumber(36)))) {
              _coll7.add(_module.__default.safeModuloInt(_234_v108, (((_208_v86).contains(_231_v105)) ? ((_208_v86).get(_231_v105)) : (_107_v1))));
            }
          }
          return _coll7;
        }());
        let _235_v109;
        _235_v109 = _dafny.Seq.of((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]);
        _235_v109 = _235_v109;
        let _236_v110;
        let _init7 = ((_237_v2, _238_v107, _239_v105, _240_v106, _241_v85) => function (_242_i6) {
          return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((((_237_v2).contains(new BigNumber(354))) ? ((_237_v2).get(new BigNumber(354))) : (_238_v107)),_239_v105), _dafny.Map.Empty.slice().updateUnsafe(_240_v106,(_module.D1.create_DC1((_241_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_241_v85).length))])).dtor_cf1));
        })(_108_v2, _233_v107, _231_v105, _232_v106, _205_v85);
        let _nw36 = Array((new BigNumber(26)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw36.length); _i0_7++) {
          _nw36[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _236_v110 = _nw36;
        let _243_v111;
        _243_v111 = _dafny.Map.Empty.slice().updateUnsafe(_233_v107,true);
        let _244_v112;
        _244_v112 = _dafny.MultiSet.fromElements(_243_v111, _243_v111);
        let _index40 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_236_v110).length));
        (_236_v110)[_index40] = (_244_v112).Union(_244_v112);
        let _245_v113;
        let _nw37 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _245_v113 = _nw37;
        let _index41 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
        (_245_v113)[_index41] = _215_v93;
        let _246_v114;
        _246_v114 = _dafny.Seq.of(_215_v93);
        let _index42 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_245_v113).length));
        (_245_v113)[_index42] = ((_109_v3) ? (_233_v107) : (new BigNumber((_246_v114).length)));
        let _index43 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_236_v110).length));
        let _index44 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
        let _index45 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_245_v113).length));
        let _rhs41 = (_244_v112).Intersect(_dafny.MultiSet.fromElements(function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_243_v111).Keys.Elements) {
            let _247_v115 = _compr_8;
            if ((_243_v111).contains(_247_v115)) {
              _coll8.push([(_247_v115).multipliedBy(new BigNumber(-496)),true]);
            }
          }
          return _coll8;
        }(), _243_v111));
        let _rhs42 = !((_208_v86).IsDisjointFrom((_208_v86).update(_109_v3, _module.__default.abs(_107_v1))));
        let _rhs43 = (((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]) ? (_module.__default.fm8(_246_v114, _105_globalState)) : (_dafny.Seq.UnicodeFromString("l")));
        let _rhs44 = _module.__default.safeDivisionInt(_107_v1, _module.__default.safeDivisionInt(_233_v107, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm8(_246_v114, _105_globalState)).length))));
        let _rhs45 = (_214_v92).multipliedBy(_233_v107);
        let _lhs35 = _236_v110;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_236_v110).length));
        let _lhs37 = _105_globalState;
        let _lhs38 = _245_v113;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
        let _lhs40 = _245_v113;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_245_v113).length));
        _lhs35[_lhs36] = _rhs41;
        _lhs37.f0 = _rhs42;
        _106_v0 = _rhs43;
        _lhs38[_lhs39] = _rhs44;
        _lhs40[_lhs41] = _rhs45;
        if (!(!(_109_v3) || (_231_v105))) {
          let _248_v116;
          _248_v116 = _dafny.Map.Empty.slice().updateUnsafe(_243_v111,_208_v86);
          let _249_v117;
          _249_v117 = _module.D1.create_DC1((_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))]);
          let _250_v118;
          let _251_v119;
          let _252_v120;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector6 = _module.__default.m0((_208_v86).Intersect(((((_248_v116).contains(_243_v111)) ? ((_248_v116).get(_243_v111)) : (_dafny.MultiSet.fromElements(_231_v105)))).update((_249_v117).dtor_cf1, _module.__default.abs(new BigNumber((_243_v111).length)))), _233_v107, _105_globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _250_v118 = _out18;
          _251_v119 = _out19;
          _252_v120 = _out20;
          let _253_v121;
          _253_v121 = _module.D1.create_DC4();
          _253_v121 = _module.D1.create_DC4();
          let _index46 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
          (_205_v85)[_index46] = !(((_213_v91) ? ((_212_v90).dtor_cf2) : (_215_v93))).isEqualTo((_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(new BigNumber((_208_v86).cardinality()), _215_v93, _215_v93, new BigNumber((_235_v109).length), (_245_v113)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length))])).length)).minus(_107_v1)));
          let _index47 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          let _rhs46 = _252_v120;
          let _rhs47 = _module.__default.safeDivisionInt(_232_v106, (_214_v92).plus(_215_v93));
          let _lhs42 = _245_v113;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          _lhs42[_lhs43] = _rhs46;
          _252_v120 = _rhs47;
          let _index48 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          (_245_v113)[_index48] = _251_v119;
        } else {
          let _254_v122;
          let _nw38 = new _module.C0();
          _nw38.__ctor(_204_v84);
          _254_v122 = _nw38;
          let _index49 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length));
          (_205_v85)[_index49] = !(_213_v91);
          let _index50 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          (_245_v113)[_index50] = new BigNumber(514);
          let _255_v123;
          let _nw39 = new _module.C0();
          _nw39.__ctor(_204_v84);
          _255_v123 = _nw39;
          let _256_v124;
          let _init8 = ((_257_v91, _258_v0, _259_v84) => function (_260_i7) {
            return ((_257_v91) ? (_258_v0) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(495)), ((_261_v84) => function (_262_i8) {
              return _261_v84;
            })(_259_v84))));
          })(_213_v91, _106_v0, _204_v84);
          let _nw40 = Array((new BigNumber(26)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw40.length); _i0_8++) {
            _nw40[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _256_v124 = _nw40;
          let _263_v125;
          _263_v125 = _module.D3.create_DC8(_256_v124);
          let _index51 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          let _rhs48 = (new BigNumber(633)).multipliedBy(((_246_v114)[_module.__default.safeIndex(_215_v93, new BigNumber((_246_v114).length))]).plus(new BigNumber(-719)));
          let _rhs49 = (_263_v125).dtor_cf9;
          let _rhs50 = _215_v93;
          let _rhs51 = ((((_208_v86).contains(_213_v91)) ? ((_208_v86).get(_213_v91)) : (_module.__default.fm0(_105_globalState)))).isEqualTo(_233_v107);
          let _lhs44 = _245_v113;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_245_v113).length));
          let _lhs46 = _105_globalState;
          _lhs44[_lhs45] = _rhs48;
          _256_v124 = _rhs49;
          _215_v93 = _rhs50;
          _lhs46.f0 = _rhs51;
        }
      }
      let _264_v126;
      _264_v126 = _dafny.Seq.of(_214_v92);
      let _265_v127;
      let _nw41 = new _module.C0();
      _nw41.__ctor(_204_v84);
      _265_v127 = _nw41;
      let _266_v128;
      _266_v128 = _dafny.Map.Empty.slice().updateUnsafe(_214_v92,_265_v127);
      let _267_v129;
      _267_v129 = _dafny.Seq.of((((_266_v128).contains(_215_v93)) ? ((_266_v128).get(_215_v93)) : (_265_v127)), _265_v127);
      _107_v1 = (_107_v1).plus(new BigNumber((_dafny.Seq.Concat(_264_v126, _dafny.Seq.of(new BigNumber(638), _214_v92, new BigNumber((_267_v129).length)))).length));
      _107_v1 = _215_v93;
      _109_v3 = !(_213_v91);
      let _268_i9;
      _268_i9 = _dafny.ZERO;
      L0: {
        while (_dafny.areEqual(_dafny.Seq.update(_264_v126, _module.__default.safeIndex(_215_v93, new BigNumber((_264_v126).length)), _214_v92), _dafny.Seq.Concat(_264_v126, _264_v126))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_268_i9)) {
              break L0;
            }
            _268_i9 = (_268_i9).plus(_dafny.ONE);
            let _269_v130;
            _269_v130 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_215_v93,_213_v91), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_264_v126).length),(_module.D1.create_DC2(new BigNumber(58), (_205_v85)[_module.__default.safeIndex(new BigNumber(445), new BigNumber((_205_v85).length))])).dtor_cf3));
            _269_v130 = _269_v130;
            _215_v93 = (_264_v126)[_module.__default.safeIndex(_107_v1, new BigNumber((_264_v126).length))];
            _214_v92 = _107_v1;
            let _270_v131;
            let _nw42 = new _module.C0();
            _nw42.__ctor(new _dafny.CodePoint('g'.codePointAt(0)));
            _270_v131 = _nw42;
            _270_v131 = _270_v131;
          }
        }
      }
      (_105_globalState).f0 = (_109_v3) && (_213_v91);
      _213_v91 = _109_v3;
      process.stdout.write(_dafny.toString(_105_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_106_v0, _dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_107_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_108_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-27),new BigNumber(-27)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_109_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_204_v84));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v85)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v86).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_209_v87).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_210_v88, _dafny.Seq.of(_dafny.Set.fromElements(_dafny.ONE), _dafny.Set.fromElements(_dafny.ONE), _dafny.Set.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_211_v89).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v90).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_213_v91));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_214_v92));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_215_v93));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_264_v126, _dafny.Seq.of(new BigNumber(455)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_266_v128).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_267_v129).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_268_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.ZERO;
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
    static create_DC2(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC3(cf4, cf5) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4() {
      let $dt = new D1(2);
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(3);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC1() { return this.$tag === 3; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4";
      } else if (this.$tag === 3) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, false);
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
    static create_DC6(cf7) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC7(cf8) {
      let $dt = new D2(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(null);
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
    static create_DC9(cf10) {
      let $dt = new D3(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D3(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf9 === other.cf9;
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
    static create_DC11(cf12, cf13) {
      let $dt = new D4(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC10(cf11) {
      let $dt = new D4(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + this.cf11.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_module.D1.Default(), _dafny.MultiSet.Empty);
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
    static create_DC13(cf15) {
      let $dt = new D5(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC12(cf14) {
      let $dt = new D5(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC14(cf16) {
      let $dt = new D5(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this._f1 = false;
      this._f2 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f3 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f3) {
      let _this = this;
      (_this)._f3 = f3;
      return;
    }
    fm3(globalState) {
      let _this = this;
      let _source2 = new BigNumber((_dafny.Seq.UnicodeFromString("tqw")).length);
      let _271___mcc_h0 = _source2;
      let _272_cf0 = _271___mcc_h0;
      return !(false);
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (_module.D1.create_DC1(true)).dtor_cf1;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
