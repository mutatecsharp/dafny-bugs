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
      return (_module.D0.create_DC2(false, false, new BigNumber(-941))).dtor_cf4;
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("csbalwk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-85)), function (_0_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(394)), function (_1_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }));
    };
    static fm2(globalState) {
      if (!(!(true)) || (false)) {
        return (_dafny.Set.fromElements(new BigNumber(-71))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(-563)));
      } else {
        return false;
      }
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, !(!(!(false)))), _dafny.Seq.of(true)), _dafny.Seq.of(!(!(true)), true));
    };
    static fm7(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true))));
    };
    static fm8(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(450),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-488),false));
    };
    static fm9(globalState) {
      return ((_dafny.Set.fromElements(false, !(true))).Difference(_dafny.Set.fromElements(false, true))).Intersect(_dafny.Set.fromElements(false));
    };
    static fm10(p0, p1, globalState) {
      if (!(!(false))) {
        if (true) {
          return _module.D0.create_DC2(false, true, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(981),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(758)), function (_2_i0) {
  return new _dafny.CodePoint('a'.codePointAt(0));
})).length))).length)));
        } else {
          return _module.D0.create_DC2(!(false), (_module.D3.create_DC8(new BigNumber(431), new BigNumber(-307), true)).dtor_cf19, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()));
        }
      } else {
        return _module.D0.create_DC2((_module.D0.create_DC2(!(true), true, new BigNumber((_dafny.Seq.of(new BigNumber(943))).length))).dtor_cf3, true, new BigNumber(567));
      }
    };
    static fm11(p0, globalState) {
      return _module.D2.create_DC6(_dafny.Seq.UnicodeFromString("ft"), (new BigNumber((_dafny.Set.fromElements(true)).length)).plus((_module.D3.create_DC8(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true, true)).length))).length), new BigNumber(798), !(!(true)))).dtor_cf17), (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).IsProperSubsetOf(function () {
  let _coll0 = new _dafny.Set();
  for (const _compr_0 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,true))).Elements) {
    let _3_v0 = _compr_0;
    if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,true))).contains(_3_v0)) {
      _coll0.add(_3_v0);
    }
  }
  return _coll0;
}()), new BigNumber((((true) ? (_dafny.MultiSet.fromElements(true, true)) : (_dafny.MultiSet.fromElements((_module.D3.create_DC8(new BigNumber(171), new BigNumber((_dafny.Seq.UnicodeFromString("otasatp")).length), false)).dtor_cf19, false, false)))).cardinality()), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))),_dafny.Seq.of(false))).length), new BigNumber(118))).IsDisjointFrom(_dafny.MultiSet.fromElements(new BigNumber(587))));
    };
    static fm12(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(239));
    };
    static fm13(p0, p1, p2, globalState) {
      if (!((false) === (!(!(true))))) {
        return _module.D4.create_DC9(function () {
  let _coll1 = new _dafny.Map();
  for (const _compr_1 of (_dafny.MultiSet.fromElements(_module.D3.create_DC8(new BigNumber(764), new BigNumber(168), !(false)))).Elements) {
    let _4_v0 = _compr_1;
    if ((_dafny.MultiSet.fromElements(_module.D3.create_DC8(new BigNumber(764), new BigNumber(168), !(false)))).contains(_4_v0)) {
      _coll1.push([_4_v0,_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0)))]);
    }
  }
  return _coll1;
}());
      } else if (true) {
        return _module.D4.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(new BigNumber(-664), new BigNumber((_dafny.Seq.of(new BigNumber(90))).length), true),_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)))));
      } else {
        return _module.D4.create_DC9(_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-837),true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('i'.codePointAt(0)))).length))).length), true),_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), function (_5_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})));
      }
    };
    static fm14(p0, p1, globalState) {
      return _dafny.Set.fromElements(new BigNumber(305));
    };
    static fm15(p0, p1, globalState) {
      return _module.D0.create_DC0(_dafny.Seq.of(!(false), true));
    };
    static fm16(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(!(true), !((!(!((_module.D0.create_DC2(true, false, new BigNumber(735))).dtor_cf2))) && (false)));
    };
    static fm17(p0, p1, p2, globalState) {
      let _source0 = ((!(false)) ? (_module.D2.create_DC6(_dafny.Seq.UnicodeFromString("mjffy"), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false, false, !(false)))).length), false)) : (_module.D2.create_DC6(_dafny.Seq.UnicodeFromString("l"), new BigNumber(139), !(true), new BigNumber((_dafny.Seq.UnicodeFromString("dop")).length), false)));
      if (_source0.is_DC6) {
        let _6___mcc_h0 = (_source0).cf11;
        let _7___mcc_h1 = (_source0).cf12;
        let _8___mcc_h2 = (_source0).cf13;
        let _9___mcc_h3 = (_source0).cf14;
        let _10___mcc_h4 = (_source0).cf15;
        let _11_cf15 = _10___mcc_h4;
        let _12_cf14 = _9___mcc_h3;
        let _13_cf13 = _8___mcc_h2;
        let _14_cf12 = _7___mcc_h1;
        let _15_cf11 = _6___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), ((_16_cf14) => function (_17_i0) {
          return function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(790), new BigNumber(970))) {
              let _18_v0 = _compr_2;
              if (((new BigNumber(790)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(970)))) {
                _coll2.push([_module.__default.safeDivisionInt(_18_v0, _16_cf14),new BigNumber(857)]);
              }
            }
            return _coll2;
          }();
        })(_12_cf14)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(701)), ((_19_cf12) => function (_20_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-290),_19_cf12);
        })(_14_cf12)));
      } else {
        let _21___mcc_h5 = (_source0).cf10;
        let _22_cf10 = _21___mcc_h5;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(980)), function (_23_i2) {
          return (function () {
            let _coll3 = new _dafny.Map();
            for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-768), new BigNumber(394))) {
              let _24_v1 = _compr_3;
              if (((new BigNumber(-768)).isLessThanOrEqualTo(_24_v1)) && ((_24_v1).isLessThan(new BigNumber(394)))) {
                _coll3.push([(_24_v1).multipliedBy(new BigNumber(-880)),new BigNumber(586)]);
              }
            }
            return _coll3;
          }());
        });
      }
    };
    static m0(p0, p1, p2, p3, globalState) {
      if (p0) {
        let _25_v0;
        _25_v0 = true;
        let _26_v1;
        _26_v1 = _dafny.Seq.of(p3);
        _25_v0 = !_dafny.Seq.contains(_26_v1, p1);
        let _27_v2;
        _27_v2 = _module.D3.create_DC8(new BigNumber(-73), p3, _25_v0);
        let _28_v3;
        let _nw0 = new _module.C0();
        _nw0.__ctor(_27_v2);
        _28_v3 = _nw0;
        let _29_v4;
        _29_v4 = _dafny.Seq.of(_25_v0);
        let _30_v5;
        _30_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_29_v4).length),new BigNumber((_29_v4).length));
        let _31_v6;
        _31_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_dafny.MultiSet.fromElements(p0, p2)).update(p0, _module.__default.abs(p3))).cardinality()));
        let _32_v7;
        let _nw1 = Array((new BigNumber(13)).toNumber()).fill(false);
        _32_v7 = _nw1;
        let _33_v8;
        _33_v8 = _module.D4.create_DC12(_25_v0, _dafny.Seq.of(_30_v5, _30_v5, _30_v5, _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_31_v6).length)), _30_v5), p1, _32_v7, p0);
        let _34_v9;
        _34_v9 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_25_v0,(_33_v8).dtor_cf28), _31_v6, _31_v6);
        let _35_v10;
        _35_v10 = _module.D4.create_DC11(_28_v3, (((_34_v9).contains(_dafny.Map.Empty.slice().updateUnsafe(_25_v0,(_28_v3).fm6(p0, p1, !(true), globalState)))) ? ((_34_v9).get(_dafny.Map.Empty.slice().updateUnsafe(_25_v0,(_28_v3).fm6(p0, p1, !(true), globalState)))) : (p3)), p3);
        _35_v10 = _35_v10;
        let _index0 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length));
        (_32_v7)[_index0] = false;
        let _index1 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length));
        (_32_v7)[_index1] = _module.__default.fm2(globalState);
        let _index2 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length));
        (_32_v7)[_index2] = p0;
        if ((p2) && (_module.__default.fm2(globalState))) {
          let _36_v11;
          let _nw2 = new _module.C1();
          _nw2.__ctor();
          _36_v11 = _nw2;
          let _37_v12;
          let _nw3 = Array((new BigNumber(2)).toNumber()).fill([]);
          _37_v12 = _nw3;
          let _38_v13;
          _38_v13 = new _dafny.CodePoint('j'.codePointAt(0));
          let _39_v14;
          _39_v14 = _dafny.Map.Empty.slice().updateUnsafe(_38_v13,_29_v4);
          let _40_v15;
          _40_v15 = _module.D0.create_DC0(_dafny.Seq.of(!(p2)));
          let _41_v16;
          let _nw4 = Array((new BigNumber(16)).toNumber());
          _nw4[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p0);
          _nw4[(_dafny.ONE).toNumber()] = _29_v4;
          _nw4[(new BigNumber(2)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(p2, (_32_v7)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length))]);
          _nw4[(new BigNumber(4)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(5)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(false, true, _25_v0);
          _nw4[(new BigNumber(7)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(true);
          _nw4[(new BigNumber(9)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(10)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_25_v0, (_32_v7)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length))]);
          _nw4[(new BigNumber(12)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(13)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(14)).toNumber()] = _29_v4;
          _nw4[(new BigNumber(15)).toNumber()] = (((_39_v14).contains(_38_v13)) ? ((_39_v14).get(_38_v13)) : ((_40_v15).dtor_cf0));
          _41_v16 = _nw4;
          let _index3 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_37_v12).length));
          (_37_v12)[_index3] = ((!((_32_v7)[_module.__default.safeIndex(new BigNumber(74), new BigNumber((_32_v7).length))])) ? (_41_v16) : (_41_v16));
          let _index4 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_37_v12).length));
          (_37_v12)[_index4] = _41_v16;
          let _42_v17;
          _42_v17 = _dafny.Seq.UnicodeFromString("d");
          let _43_v18;
          _43_v18 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_42_v17).length)).isLessThan(p1),new _dafny.CodePoint('i'.codePointAt(0)));
          let _44_v19;
          _44_v19 = _38_v13;
          _43_v18 = (_43_v18).update(p2, (_44_v19));
          _27_v2 = _27_v2;
          let _45_v20;
          let _nw5 = new _module.C0();
          _nw5.__ctor(_27_v2);
          _45_v20 = _nw5;
        } else {
          let _46_v21;
          let _nw6 = new _module.C1();
          _nw6.__ctor();
          _46_v21 = _nw6;
          let _47_v22;
          let _nw7 = new _module.C1();
          _nw7.__ctor();
          _47_v22 = _nw7;
          (globalState).f0 = p3;
          let _48_v23;
          _48_v23 = _dafny.Seq.UnicodeFromString("mmk");
          _48_v23 = _dafny.Seq.update(_48_v23, _module.__default.safeIndex(p1, new BigNumber((_48_v23).length)), (_48_v23)[_module.__default.safeIndex(p1, new BigNumber((_48_v23).length))]);
          let _49_v24;
          _49_v24 = _dafny.Seq.of(_28_v3);
          let _50_v25;
          _50_v25 = _module.D1.create_DC4(p3, !(p0), true, false);
          let _51_v26;
          let _nw8 = Array((new BigNumber(19)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = _28_v3;
          _nw8[(_dafny.ONE).toNumber()] = _28_v3;
          _nw8[(new BigNumber(2)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(3)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(4)).toNumber()] = (_49_v24)[_module.__default.safeIndex((_50_v25).dtor_cf6, new BigNumber((_49_v24).length))];
          _nw8[(new BigNumber(5)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(6)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(7)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(8)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(9)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(10)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(11)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(12)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(13)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(14)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(15)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(16)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(17)).toNumber()] = _28_v3;
          _nw8[(new BigNumber(18)).toNumber()] = _28_v3;
          _51_v26 = _nw8;
          let _52_v27;
          _52_v27 = _dafny.Map.Empty.slice().updateUnsafe(_51_v26,_25_v0);
          _52_v27 = (_52_v27).update(_51_v26, p2);
        }
      } else {
        let _53_v28;
        _53_v28 = _dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC8(p3, p3, p0)).dtor_cf18,new BigNumber((_dafny.Seq.UnicodeFromString("slgt")).length));
        let _54_v29;
        _54_v29 = _dafny.MultiSet.fromElements(p0, (_dafny.Map.Empty.slice().updateUnsafe(_53_v28,_module.D1.create_DC4(p1, p2, !(p2), false))).contains(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(334),p3)), p2, p0, p0);
        _54_v29 = (_54_v29).Difference(_54_v29);
        if (p2) {
          (globalState).f0 = new BigNumber(70);
          let _55_v30;
          let _nw9 = new _module.C0();
          _nw9.__ctor(_module.D3.create_DC8(p3, _module.__default.fm0(globalState), p2));
          _55_v30 = _nw9;
          let _56_v31;
          _56_v31 = new _dafny.CodePoint('c'.codePointAt(0));
          let _57_v32;
          let _nw10 = Array((new BigNumber(15)).toNumber()).fill(false);
          _57_v32 = _nw10;
          let _58_v33;
          _58_v33 = _dafny.Map.Empty.slice().updateUnsafe(_56_v31,_57_v32);
          _58_v33 = (_58_v33).update(_56_v31, _57_v32);
          let _59_v34;
          _59_v34 = true;
          _59_v34 = !((_59_v34) || (p2));
          let _60_v35;
          _60_v35 = _dafny.Seq.of(!(p0), p0);
          let _61_v36;
          _61_v36 = _module.D4.create_DC9(_dafny.Map.Empty.slice().updateUnsafe((_55_v30).f14,_dafny.Seq.of(_56_v31, _56_v31)));
          let _62_v37;
          _62_v37 = _dafny.Seq.UnicodeFromString("osydgy");
          let _63_v38;
          _63_v38 = _dafny.Map.Empty.slice().updateUnsafe((_55_v30).f14,_62_v37);
          let _64_v39;
          _64_v39 = _dafny.MultiSet.fromElements(p1, new BigNumber((_module.__default.fm1(p0, (_dafny.ZERO).minus(p1), globalState)).length), p1);
          let _pat_let_tv0 = _63_v38;
          let _pat_let_tv1 = _63_v38;
          let _pat_let_tv2 = _63_v38;
          let _pat_let_tv3 = _63_v38;
          let _pat_let_tv4 = _63_v38;
          let _65_v40;
          let _nw11 = Array((new BigNumber(20)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _61_v36;
          _nw11[(_dafny.ONE).toNumber()] = _61_v36;
          _nw11[(new BigNumber(2)).toNumber()] = _module.__default.fm13(p3, _62_v37, _59_v34, globalState);
          _nw11[(new BigNumber(3)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(4)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(5)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(6)).toNumber()] = _module.D4.create_DC9(_63_v38);
          _nw11[(new BigNumber(7)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(8)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(9)).toNumber()] = _module.D4.create_DC9(_63_v38);
          _nw11[(new BigNumber(10)).toNumber()] = ((true) ? (_61_v36) : (_61_v36));
          _nw11[(new BigNumber(11)).toNumber()] = function (_pat_let0_0) {
            return function (_68_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_69_dt__update_hcf20_h1) {
                  return _module.D4.create_DC9(_69_dt__update_hcf20_h1);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let0_0);
          }(function (_pat_let1_0) {
            return function (_66_dt__update__tmp_h0) {
              return function (_pat_let2_0) {
                return function (_67_dt__update_hcf20_h0) {
                  return _module.D4.create_DC9(_67_dt__update_hcf20_h0);
                }(_pat_let2_0);
              }(_pat_let_tv0);
            }(_pat_let1_0);
          }(_61_v36));
          _nw11[(new BigNumber(12)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(13)).toNumber()] = function (_pat_let4_0) {
            return function (_72_dt__update__tmp_h3) {
              return function (_pat_let7_0) {
                return function (_73_dt__update_hcf20_h3) {
                  return _module.D4.create_DC9(_73_dt__update_hcf20_h3);
                }(_pat_let7_0);
              }(_pat_let_tv3);
            }(_pat_let4_0);
          }(function (_pat_let5_0) {
            return function (_70_dt__update__tmp_h2) {
              return function (_pat_let6_0) {
                return function (_71_dt__update_hcf20_h2) {
                  return _module.D4.create_DC9(_71_dt__update_hcf20_h2);
                }(_pat_let6_0);
              }(_pat_let_tv2);
            }(_pat_let5_0);
          }(_61_v36));
          _nw11[(new BigNumber(14)).toNumber()] = ((p0) ? (function (_pat_let8_0) {
            return function (_74_dt__update__tmp_h4) {
              return function (_pat_let9_0) {
                return function (_75_dt__update_hcf20_h4) {
                  return _module.D4.create_DC9(_75_dt__update_hcf20_h4);
                }(_pat_let9_0);
              }(_pat_let_tv4);
            }(_pat_let8_0);
          }(_61_v36)) : (_61_v36));
          _nw11[(new BigNumber(15)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(16)).toNumber()] = _61_v36;
          _nw11[(new BigNumber(17)).toNumber()] = _module.D4.create_DC9(_63_v38);
          _nw11[(new BigNumber(18)).toNumber()] = _module.D4.create_DC9(_63_v38);
          _nw11[(new BigNumber(19)).toNumber()] = _module.__default.fm13(new BigNumber((_64_v39).cardinality()), _62_v37, _59_v34, globalState);
          _65_v40 = _nw11;
          let _index5 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_65_v40).length));
          (_65_v40)[_index5] = _61_v36;
          let _76_v41;
          _76_v41 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm1(p0, new BigNumber((_62_v37).length), globalState));
          let _77_v42;
          let _nw12 = new _module.C1();
          _nw12.__ctor();
          _77_v42 = _nw12;
          let _78_v43;
          _78_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,_77_v42);
          let _79_v44;
          _79_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_78_v43).length),_59_v34);
          let _pat_let_tv5 = _63_v38;
          let _index6 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_65_v40).length));
          let _rhs0 = _module.__default.fm3(p1, _module.__default.safeModuloInt(p1, p3), (new BigNumber((_76_v41).length)).minus(p1), ((p0) ? (_79_v44) : (_79_v44)), globalState);
          let _rhs1 = ((p0) ? (function (_pat_let10_0) {
            return function (_80_dt__update__tmp_h5) {
              return function (_pat_let11_0) {
                return function (_81_dt__update_hcf20_h5) {
                  return _module.D4.create_DC9(_81_dt__update_hcf20_h5);
                }(_pat_let11_0);
              }(_pat_let_tv5);
            }(_pat_let10_0);
          }(_61_v36)) : (_61_v36));
          let _lhs0 = _65_v40;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_65_v40).length));
          _60_v35 = _rhs0;
          _lhs0[_lhs1] = _rhs1;
        } else {
          let _82_v45;
          _82_v45 = true;
          let _83_v46;
          let _nw13 = Array((new BigNumber(11)).toNumber()).fill([]);
          _83_v46 = _nw13;
          let _84_v47;
          let _init0 = ((_85_p3) => function (_86_i0) {
            return (_86_i0).minus((_dafny.ZERO).minus(_85_p3));
          })(p3);
          let _nw14 = Array((new BigNumber(8)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw14.length); _i0_0++) {
            _nw14[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _84_v47 = _nw14;
          let _index7 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_83_v46).length));
          (_83_v46)[_index7] = _84_v47;
          let _87_v48;
          _87_v48 = _module.D3.create_DC8(p1, p3, p0);
          let _88_v49;
          let _nw15 = new _module.C0();
          _nw15.__ctor(_87_v48);
          _88_v49 = _nw15;
          let _89_v50;
          _89_v50 = _dafny.Seq.of(_88_v49);
          let _90_v51;
          _90_v51 = _dafny.Seq.of(p2, p2, p2);
          let _91_v52;
          _91_v52 = _dafny.Set.fromElements(p1, new BigNumber((_90_v51).length));
          let _92_v53;
          _92_v53 = _dafny.Seq.of(_91_v52);
          let _93_v54;
          _93_v54 = _module.D3.create_DC8(p3, new BigNumber((_dafny.Seq.update(_89_v50, _module.__default.safeIndex(new BigNumber((_92_v53).length), new BigNumber((_89_v50).length)), _88_v49)).length), p0);
          let _94_v55;
          let _nw16 = Array((new BigNumber(12)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _82_v45;
          _nw16[(_dafny.ONE).toNumber()] = p2;
          _nw16[(new BigNumber(2)).toNumber()] = _82_v45;
          _nw16[(new BigNumber(3)).toNumber()] = p2;
          _nw16[(new BigNumber(4)).toNumber()] = _82_v45;
          _nw16[(new BigNumber(5)).toNumber()] = p0;
          _nw16[(new BigNumber(6)).toNumber()] = _82_v45;
          _nw16[(new BigNumber(7)).toNumber()] = _82_v45;
          _nw16[(new BigNumber(8)).toNumber()] = (_87_v48).dtor_cf19;
          _nw16[(new BigNumber(9)).toNumber()] = !(p2);
          _nw16[(new BigNumber(10)).toNumber()] = p0;
          _nw16[(new BigNumber(11)).toNumber()] = p0;
          _94_v55 = _nw16;
          let _95_v56;
          _95_v56 = _dafny.Seq.of(_94_v55);
          let _96_v57;
          _96_v57 = _dafny.Map.Empty.slice().updateUnsafe(_94_v55,_94_v55);
          let _97_v58;
          let _nw17 = new _module.C1();
          _nw17.__ctor();
          _97_v58 = _nw17;
          let _98_v59;
          _98_v59 = _dafny.Map.Empty.slice().updateUnsafe(_97_v58,_94_v55);
          let _99_v60;
          let _nw18 = Array((new BigNumber(12)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (_95_v56)[_module.__default.safeIndex(p3, new BigNumber((_95_v56).length))];
          _nw18[(_dafny.ONE).toNumber()] = _94_v55;
          _nw18[(new BigNumber(2)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(3)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(4)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(5)).toNumber()] = (((_96_v57).contains(_94_v55)) ? ((_96_v57).get(_94_v55)) : (_94_v55));
          _nw18[(new BigNumber(6)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(7)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(8)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(9)).toNumber()] = _94_v55;
          _nw18[(new BigNumber(10)).toNumber()] = ((true) ? (_94_v55) : (_94_v55));
          _nw18[(new BigNumber(11)).toNumber()] = (((_98_v59).contains(_97_v58)) ? ((_98_v59).get(_97_v58)) : (_94_v55));
          _99_v60 = _nw18;
          let _index8 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_99_v60).length));
          (_99_v60)[_index8] = _94_v55;
          let _index9 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_83_v46).length));
          let _index10 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_99_v60).length));
          let _rhs2 = _82_v45;
          let _rhs3 = _84_v47;
          let _rhs4 = _94_v55;
          let _rhs5 = p2;
          let _lhs2 = _83_v46;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_83_v46).length));
          let _lhs4 = _99_v60;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(247), new BigNumber((_99_v60).length));
          _82_v45 = _rhs2;
          _lhs2[_lhs3] = _rhs3;
          _lhs4[_lhs5] = _rhs4;
          _82_v45 = _rhs5;
          let _100_v61;
          let _init1 = ((_101_v52) => function (_102_i1) {
            return _101_v52;
          })(_91_v52);
          let _nw19 = Array((new BigNumber(26)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw19.length); _i0_1++) {
            _nw19[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _100_v61 = _nw19;
          let _index11 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_100_v61).length));
          (_100_v61)[_index11] = _91_v52;
          let _103_v62;
          _103_v62 = new _dafny.CodePoint('e'.codePointAt(0));
          let _index12 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_100_v61).length));
          (_100_v61)[_index12] = ((_module.__default.fm14(_103_v62, (_dafny.ZERO).minus(p1), globalState)).Intersect(_91_v52)).Union(_91_v52);
          let _104_v63;
          _104_v63 = _dafny.MultiSet.fromElements(p3);
          let _105_v64;
          _105_v64 = _dafny.MultiSet.fromElements(_module.__default.fm15((_dafny.ZERO).minus(new BigNumber((_104_v63).cardinality())), p1, globalState), _module.D0.create_DC0(_90_v51));
          let _106_v65;
          _106_v65 = _dafny.Map.Empty.slice().updateUnsafe(!(_82_v45) || (p0),_105_v64);
          _106_v65 = (_106_v65).update(!(p0), _105_v64);
          _82_v45 = (_90_v51)[_module.__default.safeIndex(p3, new BigNumber((_90_v51).length))];
          let _107_v66;
          _107_v66 = _dafny.Seq.of(_54_v29);
          _107_v66 = ((p0) ? (_107_v66) : (_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(p0)), _54_v29, _54_v29, _module.__default.fm16(p3, p1, globalState), _54_v29)));
        }
        let _108_v67;
        _108_v67 = _dafny.Seq.UnicodeFromString("iwvnrgqkx");
        let _109_v68;
        _109_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("jgktoekf")).length),true);
        (globalState).f0 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_108_v67).length), _module.__default.safeModuloInt(new BigNumber(489), new BigNumber((_109_v68).length))));
        let _110_v69;
        _110_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,_54_v29);
        let _111_v70;
        _111_v70 = _dafny.Seq.of(p0);
        _110_v69 = (_110_v69).update(p3, _dafny.MultiSet.FromArray(_111_v70));
        let _112_v71;
        _112_v71 = true;
        _112_v71 = (p3).isLessThanOrEqualTo(_module.__default.safeModuloInt(p1, p3));
      }
      (globalState).f0 = p1;
      let _113_v72;
      _113_v72 = _dafny.Seq.UnicodeFromString("oncu");
      let _114_v73;
      _114_v73 = _dafny.Seq.of(new BigNumber(3), new BigNumber(236));
      let _hi0 = (_dafny.ZERO).minus((_114_v73)[_module.__default.safeIndex(p1, new BigNumber((_114_v73).length))]);
      for (let _115_i2 = (_dafny.ZERO).minus(new BigNumber((_113_v72).length)); _115_i2.isLessThan(_hi0); _115_i2 = _115_i2.plus(_dafny.ONE)) {
        let _116_v74;
        _116_v74 = _dafny.Map.Empty.slice().updateUnsafe(p3,_113_v72);
        let _117_v75;
        _117_v75 = new _dafny.CodePoint('i'.codePointAt(0));
        let _118_v76;
        _118_v76 = _dafny.Map.Empty.slice().updateUnsafe(_116_v74,_117_v75);
        _118_v76 = (_118_v76).update(_116_v74, _117_v75);
        _113_v72 = _113_v72;
        (globalState).f0 = p3;
        let _119_v77;
        let _init2 = ((_120_p1) => function (_121_i3) {
          return (_121_i3).minus(_120_p1);
        })(p1);
        let _nw20 = Array((new BigNumber(7)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
          _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _119_v77 = _nw20;
        let _index13 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_119_v77).length));
        (_119_v77)[_index13] = _115_i2;
        let _122_v78;
        _122_v78 = _dafny.MultiSet.fromElements(p3, p3);
        let _123_v79;
        _123_v79 = _dafny.Seq.of(_122_v78);
        let _124_v80;
        let _nw21 = new _module.C1();
        _nw21.__ctor();
        _124_v80 = _nw21;
        let _125_v81;
        _125_v81 = _dafny.Map.Empty.slice().updateUnsafe(_124_v80,_124_v80);
        let _index14 = _module.__default.safeIndex(new BigNumber(883), new BigNumber((_119_v77).length));
        (_119_v77)[_index14] = (_dafny.ZERO).minus((p3).minus(((p0) ? (new BigNumber((_123_v79).length)) : ((_dafny.ZERO).minus(new BigNumber((_125_v81).length))))));
      }
      (globalState).f0 = p1;
      let _126_v82;
      let _nw22 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
      _126_v82 = _nw22;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_126_v82).length))) {
        let _127_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_127_i4)) && ((_127_i4).isLessThan(new BigNumber((_126_v82).length))))) {
          (_126_v82)[(_127_i4)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_114_v73, _114_v73), _114_v73);
        }
      }
      let _128_v83;
      let _nw23 = Array((new BigNumber(27)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = false;
      _nw23[(_dafny.ONE).toNumber()] = p0;
      _nw23[(new BigNumber(2)).toNumber()] = p2;
      _nw23[(new BigNumber(3)).toNumber()] = !(p0);
      _nw23[(new BigNumber(4)).toNumber()] = p2;
      _nw23[(new BigNumber(5)).toNumber()] = p0;
      _nw23[(new BigNumber(6)).toNumber()] = p0;
      _nw23[(new BigNumber(7)).toNumber()] = p2;
      _nw23[(new BigNumber(8)).toNumber()] = p2;
      _nw23[(new BigNumber(9)).toNumber()] = p0;
      _nw23[(new BigNumber(10)).toNumber()] = p0;
      _nw23[(new BigNumber(11)).toNumber()] = p0;
      _nw23[(new BigNumber(12)).toNumber()] = p0;
      _nw23[(new BigNumber(13)).toNumber()] = p2;
      _nw23[(new BigNumber(14)).toNumber()] = p2;
      _nw23[(new BigNumber(15)).toNumber()] = p0;
      _nw23[(new BigNumber(16)).toNumber()] = p2;
      _nw23[(new BigNumber(17)).toNumber()] = _module.__default.fm2(globalState);
      _nw23[(new BigNumber(18)).toNumber()] = p0;
      _nw23[(new BigNumber(19)).toNumber()] = p2;
      _nw23[(new BigNumber(20)).toNumber()] = p0;
      _nw23[(new BigNumber(21)).toNumber()] = p0;
      _nw23[(new BigNumber(22)).toNumber()] = p2;
      _nw23[(new BigNumber(23)).toNumber()] = p0;
      _nw23[(new BigNumber(24)).toNumber()] = !(p0);
      _nw23[(new BigNumber(25)).toNumber()] = p0;
      _nw23[(new BigNumber(26)).toNumber()] = p0;
      _128_v83 = _nw23;
      let _129_v84;
      _129_v84 = _module.D4.create_DC12(p2, _module.__default.fm17(p3, p2, false, globalState), p1, _128_v83, p2);
      let _130_v85;
      _130_v85 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_dafny.ZERO).minus(_module.__default.fm0(globalState)));
      let _131_v86;
      _131_v86 = _dafny.Seq.of(_130_v85);
      let _pat_let_tv6 = _128_v83;
      let _pat_let_tv7 = _114_v73;
      let _pat_let_tv8 = _131_v86;
      let _132_v87;
      let _nw24 = Array((_dafny.ONE).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = function (_pat_let12_0) {
        return function (_133_dt__update__tmp_h6) {
          return function (_pat_let13_0) {
            return function (_134_dt__update_hcf29_h0) {
              return function (_pat_let14_0) {
                return function (_135_dt__update_hcf28_h0) {
                  return function (_pat_let15_0) {
                    return function (_136_dt__update_hcf27_h0) {
                      return _module.D4.create_DC12((_133_dt__update__tmp_h6).dtor_cf26, _136_dt__update_hcf27_h0, _135_dt__update_hcf28_h0, _134_dt__update_hcf29_h0, (_133_dt__update__tmp_h6).dtor_cf30);
                    }(_pat_let15_0);
                  }(_pat_let_tv8);
                }(_pat_let14_0);
              }(new BigNumber((_pat_let_tv7).length));
            }(_pat_let13_0);
          }(_pat_let_tv6);
        }(_pat_let12_0);
      }(_129_v84);
      _132_v87 = _nw24;
      let _index15 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_132_v87).length));
      (_132_v87)[_index15] = _129_v84;
      let _index16 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_132_v87).length));
      (_132_v87)[_index16] = _129_v84;
      return;
    }
    static Main(__noArgsParameter) {
      let _137_v0;
      _137_v0 = true;
      let _138_v1;
      _138_v1 = _dafny.Set.fromElements(_137_v0, _137_v0, _137_v0);
      let _139_v2;
      let _init3 = ((_140_v0) => function (_141_i0) {
        return (_module.D0.create_DC0(_dafny.Seq.of(_140_v0, _140_v0))).dtor_cf0;
      })(_137_v0);
      let _nw25 = Array((_dafny.ONE).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw25.length); _i0_3++) {
        _nw25[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _139_v2 = _nw25;
      let _142_v3;
      _142_v3 = new BigNumber(-730);
      let _143_v4;
      _143_v4 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_142_v3);
      let _144_v5;
      let _nw26 = Array((new BigNumber(21)).toNumber());
      _nw26[(_dafny.ZERO).toNumber()] = _143_v4;
      _nw26[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_142_v3);
      _nw26[(new BigNumber(2)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(3)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(4)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_142_v3);
      _nw26[(new BigNumber(6)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(7)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(8)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(9)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_142_v3);
      _nw26[(new BigNumber(11)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(12)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(13)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(14)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(15)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(16)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(17)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(18)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(19)).toNumber()] = _143_v4;
      _nw26[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_142_v3);
      _144_v5 = _nw26;
      let _145_globalState;
      let _nw27 = new _module.GlobalState();
      _nw27.__ctor(new BigNumber(317), (_dafny.Set.fromElements(_137_v0)).Difference(_138_v1), _139_v2, _144_v5, new _dafny.CodePoint('u'.codePointAt(0)), _dafny.Seq.UnicodeFromString("tinqkjdyt"), new BigNumber(892), new BigNumber(437), true, new _dafny.CodePoint('o'.codePointAt(0)), false, new BigNumber(-34), false, false);
      _145_globalState = _nw27;
      let _146_v6;
      _146_v6 = _dafny.Seq.UnicodeFromString("whhg");
      let _hi1 = new BigNumber((_146_v6).length);
      for (let _147_i1 = new BigNumber((_138_v1).length); _147_i1.isLessThan(_hi1); _147_i1 = _147_i1.plus(_dafny.ONE)) {
        let _148_v7;
        let _nw28 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _148_v7 = _nw28;
        let _149_v8;
        _149_v8 = _dafny.Seq.of(_148_v7);
        let _150_v9;
        _150_v9 = _dafny.Map.Empty.slice().updateUnsafe((_142_v3).plus(new BigNumber((_146_v6).length)),_149_v8);
        _150_v9 = (_150_v9).update((_dafny.ZERO).minus(_module.__default.fm0(_145_globalState)), _149_v8);
        _142_v3 = (_dafny.ZERO).minus(_142_v3);
        let _151_v10;
        _151_v10 = _dafny.Seq.of(_142_v3, _147_i1);
        if ((_142_v3).isLessThan(_module.__default.safeDivisionInt(_142_v3, new BigNumber(((_dafny.MultiSet.FromArray(_151_v10)).update(_142_v3, _module.__default.abs(_147_i1))).cardinality())))) {
          let _152_v11;
          _152_v11 = _dafny.Seq.of(_137_v0);
          let _153_v12;
          _153_v12 = _module.D0.create_DC2(_137_v0, _137_v0, (_dafny.ZERO).minus(new BigNumber((_152_v11).length)));
          _142_v3 = (_153_v12).dtor_cf4;
          let _154_v13;
          let _nw29 = Array((new BigNumber(24)).toNumber()).fill(false);
          _154_v13 = _nw29;
          let _155_v14;
          _155_v14 = _dafny.Seq.of(_154_v13);
          let _index17 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_154_v13).length));
          (_154_v13)[_index17] = !_dafny.areEqual(_155_v14, _155_v14);
          let _index18 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_154_v13).length));
          (_154_v13)[_index18] = (_137_v0) && (_137_v0);
          let _156_v15;
          _156_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_146_v6).length),_137_v0);
          _module.__default.m0((_142_v3).isLessThan(_147_i1), new BigNumber((_156_v15).length), (_154_v13)[_module.__default.safeIndex(new BigNumber(49), new BigNumber((_154_v13).length))], _142_v3, _145_globalState);
          let _157_v16;
          _157_v16 = _dafny.MultiSet.fromElements(_147_i1);
          let _158_v17;
          _158_v17 = _dafny.Map.Empty.slice().updateUnsafe(_147_i1,_module.__default.safeModuloInt((((_157_v16).contains(_142_v3)) ? ((_157_v16).get(_142_v3)) : (_147_i1)), _147_i1));
          let _159_v18;
          _159_v18 = new _dafny.CodePoint('i'.codePointAt(0));
          _158_v17 = (_158_v17).update(new BigNumber((_146_v6).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_160_i2) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), _module.__default.safeIndex(_142_v3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_161_i2) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length)), _159_v18)).length));
          (_145_globalState).f0 = _module.__default.safeModuloInt(_147_i1, _147_i1);
        } else {
          _142_v3 = _142_v3;
          _module.__default.m0(!(_137_v0), _142_v3, (_142_v3).isLessThan(_142_v3), _147_i1, _145_globalState);
          _module.__default.m0(_137_v0, _147_i1, false, (_142_v3).multipliedBy(_142_v3), _145_globalState);
          let _162_v19;
          _162_v19 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_146_v6);
          let _163_v20;
          _163_v20 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_dafny.Seq.IsProperPrefixOf(_146_v6, (((_162_v19).contains(_147_i1)) ? ((_162_v19).get(_147_i1)) : (_146_v6))));
          let _164_v21;
          _164_v21 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_137_v0);
          _137_v0 = (((_163_v20).contains(_147_i1)) ? ((_163_v20).get(_147_i1)) : ((((_164_v21).contains(_137_v0)) ? ((_164_v21).get(_137_v0)) : (_137_v0))));
          _163_v20 = _163_v20;
        }
        let _165_v22;
        _165_v22 = _module.D0.create_DC2(_137_v0, _137_v0, _147_i1);
        let _166_v23;
        _166_v23 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,(_165_v22).dtor_cf2);
        _module.__default.m0((((_166_v23).contains(_137_v0)) ? ((_166_v23).get(_137_v0)) : (_137_v0)), _142_v3, _137_v0, (_142_v3).plus((_dafny.ZERO).minus(_142_v3)), _145_globalState);
      }
      _module.__default.m0((_142_v3).isLessThanOrEqualTo(_142_v3), (_142_v3).minus(_142_v3), _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("bffs"), _146_v6), _142_v3, _145_globalState);
      let _hi2 = (_dafny.ZERO).minus(_142_v3);
      for (let _167_i3 = _142_v3; _167_i3.isLessThan(_hi2); _167_i3 = _167_i3.plus(_dafny.ONE)) {
        _137_v0 = _137_v0;
        _module.__default.m0(_137_v0, (_167_i3).minus(_module.__default.fm0(_145_globalState)), false, _142_v3, _145_globalState);
        let _168_v24;
        _168_v24 = _dafny.Seq.of(_137_v0, _137_v0);
        let _169_v25;
        _169_v25 = _dafny.Map.Empty.slice().updateUnsafe(_167_i3,_142_v3);
        let _170_v26;
        _170_v26 = _dafny.Seq.of(_169_v25);
        let _171_v27;
        _171_v27 = new _dafny.CodePoint('g'.codePointAt(0));
        let _172_v28;
        _172_v28 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(_module.__default.fm1(_137_v0, _142_v3, _145_globalState), _module.__default.safeIndex(new BigNumber((_170_v26).length), new BigNumber((_module.__default.fm1(_137_v0, _142_v3, _145_globalState)).length)), _171_v27)).length));
        let _173_v29;
        _173_v29 = _dafny.Map.Empty.slice().updateUnsafe(_167_i3,!(_137_v0));
        _module.__default.m0(_dafny.Seq.contains(_168_v24, _137_v0), (_172_v28)[_module.__default.safeIndex(_142_v3, new BigNumber((_172_v28).length))], (((_173_v29).contains(_167_i3)) ? ((_173_v29).get(_167_i3)) : (_137_v0)), new BigNumber((_dafny.Seq.update(_168_v24, _module.__default.safeIndex(_167_i3, new BigNumber((_168_v24).length)), _137_v0)).length), _145_globalState);
        let _174_v30;
        _174_v30 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_module.__default.fm2(_145_globalState));
        _142_v3 = (new BigNumber((_174_v30).length)).multipliedBy(_167_i3);
      }
      let _175_v31;
      _175_v31 = _module.D0.create_DC1((_142_v3).plus(_142_v3));
      let _source1 = _175_v31;
      if (_source1.is_DC1) {
        let _176___mcc_h0 = (_source1).cf1;
        let _177_cf1 = _176___mcc_h0;
        _137_v0 = _module.__default.fm2(_145_globalState);
        let _178_v32;
        _178_v32 = _dafny.Seq.of(false);
        _178_v32 = _178_v32;
        let _179_v33;
        _179_v33 = _dafny.Map.Empty.slice().updateUnsafe(_177_cf1,_137_v0);
        let _180_v34;
        _180_v34 = _dafny.Seq.of(_142_v3);
        _178_v32 = _dafny.Seq.Concat(_178_v32, _dafny.Seq.Concat(_178_v32, _module.__default.fm3((_dafny.ZERO).minus(_177_cf1), new BigNumber(-816), _177_cf1, (_179_v33).update(new BigNumber((_180_v34).length), _137_v0), _145_globalState)));
        let _181_v35;
        _181_v35 = _module.D0.create_DC2(_137_v0, _137_v0, _142_v3);
        let _182_v36;
        let _nw30 = Array((new BigNumber(12)).toNumber());
        _nw30[(_dafny.ZERO).toNumber()] = _137_v0;
        _nw30[(_dafny.ONE).toNumber()] = !(_137_v0);
        _nw30[(new BigNumber(2)).toNumber()] = (_181_v35).dtor_cf2;
        _nw30[(new BigNumber(3)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(4)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(5)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(6)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(7)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(8)).toNumber()] = false;
        _nw30[(new BigNumber(9)).toNumber()] = _137_v0;
        _nw30[(new BigNumber(10)).toNumber()] = !(!((_181_v35).dtor_cf2));
        _nw30[(new BigNumber(11)).toNumber()] = _137_v0;
        _182_v36 = _nw30;
        let _183_v37;
        _183_v37 = _dafny.MultiSet.fromElements(_182_v36, _182_v36, _182_v36);
        _137_v0 = !(!((_183_v37).contains(_182_v36))) || (_137_v0);
      } else if (_source1.is_DC2) {
        let _184___mcc_h1 = (_source1).cf2;
        let _185___mcc_h2 = (_source1).cf3;
        let _186___mcc_h3 = (_source1).cf4;
        let _187_cf4 = _186___mcc_h3;
        let _188_cf3 = _185___mcc_h2;
        let _189_cf2 = _184___mcc_h1;
        let _190_v38;
        _190_v38 = _dafny.Seq.of(_137_v0, true, _137_v0);
        _module.__default.m0(_dafny.Seq.IsPrefixOf(_190_v38, _190_v38), _142_v3, (_dafny.Map.Empty.slice().updateUnsafe(_188_cf3,(_dafny.ZERO).minus(new BigNumber((_146_v6).length)))).contains(_188_cf3), _142_v3, _145_globalState);
        let _191_v39;
        let _nw31 = new _module.C1();
        _nw31.__ctor();
        _191_v39 = _nw31;
        let _192_v40;
        let _init4 = ((_193_v0, _194_cf3, _195_cf2) => function (_196_i4) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_193_v0,_194_cf3)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC2(_193_v0, _195_cf2, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)))).cardinality()))).cardinality()))).dtor_cf3,_194_cf3));
        })(_137_v0, _188_cf3, _189_cf2);
        let _nw32 = Array((new BigNumber(5)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw32.length); _i0_4++) {
          _nw32[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _192_v40 = _nw32;
        let _index19 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_192_v40).length));
        (_192_v40)[_index19] = _module.__default.fm7(true, _145_globalState);
        let _197_v41;
        _197_v41 = _dafny.Map.Empty.slice().updateUnsafe(_188_cf3,_137_v0);
        let _index20 = _module.__default.safeIndex(new BigNumber(78), new BigNumber((_192_v40).length));
        (_192_v40)[_index20] = (_197_v41).Merge(_module.__default.fm7(_137_v0, _145_globalState));
        let _198_v42;
        let _init5 = ((_199_cf3) => function (_200_i5) {
          return _199_cf3;
        })(_188_cf3);
        let _nw33 = Array((new BigNumber(6)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
          _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _198_v42 = _nw33;
        let _index21 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_198_v42).length));
        (_198_v42)[_index21] = _188_cf3;
        let _index22 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_198_v42).length));
        (_198_v42)[_index22] = (_187_cf4).isLessThan(_142_v3);
        let _index23 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_198_v42).length));
        (_198_v42)[_index23] = (_187_cf4).isLessThan(_142_v3);
        let _index24 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_198_v42).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_198_v42).length));
        let _index26 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_198_v42).length));
        let _rhs6 = (_187_cf4).isLessThanOrEqualTo(_187_cf4);
        let _rhs7 = _188_cf3;
        let _rhs8 = false;
        let _rhs9 = _188_cf3;
        let _lhs6 = _198_v42;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_198_v42).length));
        let _lhs8 = _198_v42;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_198_v42).length));
        let _lhs10 = _198_v42;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_198_v42).length));
        _lhs6[_lhs7] = _rhs6;
        _189_cf2 = _rhs7;
        _lhs8[_lhs9] = _rhs8;
        _lhs10[_lhs11] = _rhs9;
      } else {
        let _201___mcc_h4 = (_source1).cf0;
        let _202_cf0 = _201___mcc_h4;
        _module.__default.m0(_137_v0, _142_v3, _137_v0, (_142_v3).plus(_142_v3), _145_globalState);
        let _203_v43;
        _203_v43 = _dafny.MultiSet.fromElements(_143_v4);
        _module.__default.m0(_137_v0, (new BigNumber((_203_v43).cardinality())).minus(_142_v3), _137_v0, new BigNumber((_146_v6).length), _145_globalState);
        let _204_v44;
        let _nw34 = Array((new BigNumber(15)).toNumber());
        _nw34[(_dafny.ZERO).toNumber()] = _137_v0;
        _nw34[(_dafny.ONE).toNumber()] = _137_v0;
        _nw34[(new BigNumber(2)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(3)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(4)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(5)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(6)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(7)).toNumber()] = _module.__default.fm2(_145_globalState);
        _nw34[(new BigNumber(8)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(9)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(10)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(11)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(12)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(13)).toNumber()] = _137_v0;
        _nw34[(new BigNumber(14)).toNumber()] = _137_v0;
        _204_v44 = _nw34;
        let _205_v45;
        _205_v45 = _dafny.Map.Empty.slice().updateUnsafe(_204_v44,_142_v3);
        let _206_v46;
        _206_v46 = _dafny.Map.Empty.slice().updateUnsafe((_205_v45).Merge(_dafny.Map.Empty.slice().updateUnsafe(_204_v44,_142_v3)),_dafny.Seq.Concat(_146_v6, _146_v6));
        _206_v46 = (_206_v46).update((_205_v45).Merge(_205_v45), _dafny.Seq.UnicodeFromString("jbs"));
        let _207_v47;
        _207_v47 = new _dafny.CodePoint('j'.codePointAt(0));
        (_145_globalState).f9 = _207_v47;
      }
      _137_v0 = _137_v0;
      let _208_v48;
      _208_v48 = _module.D1.create_DC4(new BigNumber(423), _137_v0, _137_v0, _137_v0);
      let _pat_let_tv9 = _142_v3;
      let _pat_let_tv10 = _143_v4;
      if (function (_source2) {
        if (_source2.is_DC4) {
          let _209___mcc_h5 = (_source2).cf6;
          let _210___mcc_h6 = (_source2).cf7;
          let _211___mcc_h7 = (_source2).cf8;
          let _212___mcc_h8 = (_source2).cf9;
          let _213_cf9 = _212___mcc_h8;
          let _214_cf8 = _211___mcc_h7;
          let _215_cf7 = _210___mcc_h6;
          let _216_cf6 = _209___mcc_h5;
          return !((_dafny.Set.fromElements(_pat_let_tv9, new BigNumber((_pat_let_tv10).length))).IsSubsetOf(_dafny.Set.fromElements(_216_cf6)));
        } else {
          let _217___mcc_h9 = (_source2).cf5;
          let _218_cf5 = _217___mcc_h9;
          return false;
        }
      }(_208_v48)) {
        let _219_v49;
        _219_v49 = _dafny.Seq.of(_142_v3, (((_143_v4).contains(_137_v0)) ? ((_143_v4).get(_137_v0)) : (_142_v3)), _142_v3, new BigNumber(390), _142_v3);
        let _rhs10 = _module.__default.safeDivisionInt(_142_v3, new BigNumber((_219_v49).length));
        let _rhs11 = _142_v3;
        let _rhs12 = _142_v3;
        let _lhs12 = _145_globalState;
        let _lhs13 = _145_globalState;
        _lhs12.f0 = _rhs10;
        _lhs13.f0 = _rhs11;
        _142_v3 = _rhs12;
        let _220_v50;
        let _nw35 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _220_v50 = _nw35;
        let _221_v51;
        _221_v51 = _dafny.Seq.of(_220_v50, _220_v50, _220_v50);
        _221_v51 = _221_v51;
        let _index27 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length));
        (_220_v50)[_index27] = _142_v3;
        let _222_v52;
        _222_v52 = _dafny.MultiSet.fromElements(_137_v0);
        let _223_v53;
        _223_v53 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_145_globalState),_142_v3);
        let _index28 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length));
        (_220_v50)[_index28] = (((_222_v52).contains(((_137_v0) ? (_137_v0) : (_137_v0)))) ? ((_222_v52).get(((_137_v0) ? (_137_v0) : (_137_v0)))) : (new BigNumber((_223_v53).length)));
        if (_module.__default.fm2(_145_globalState)) {
          let _224_v54;
          _224_v54 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_137_v0);
          _224_v54 = (_module.__default.fm8(_142_v3, _145_globalState)).Merge(_module.__default.fm8(new BigNumber(473), _145_globalState));
          let _225_v55;
          let _init6 = function (_226_i6) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          };
          let _nw36 = Array((new BigNumber(9)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw36.length); _i0_6++) {
            _nw36[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _225_v55 = _nw36;
          let _227_v56;
          _227_v56 = new _dafny.CodePoint('k'.codePointAt(0));
          let _index29 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_225_v55).length));
          (_225_v55)[_index29] = _227_v56;
          let _index30 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_225_v55).length));
          (_225_v55)[_index30] = new _dafny.CodePoint('g'.codePointAt(0));
          (_145_globalState).f0 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_220_v50)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length))], _142_v3), (_220_v50)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length))]);
          (_145_globalState).f0 = _142_v3;
          let _228_v57;
          let _nw37 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
          _228_v57 = _nw37;
          let _index31 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_228_v57).length));
          (_228_v57)[_index31] = _dafny.Seq.of(_223_v53, _223_v53, _223_v53);
          let _229_v58;
          _229_v58 = _dafny.Seq.of(_223_v53, _223_v53);
          let _index32 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_228_v57).length));
          (_228_v57)[_index32] = _dafny.Seq.Concat(_229_v58, _229_v58);
        } else {
          let _230_v59;
          _230_v59 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_137_v0)),_dafny.Seq.UnicodeFromString("jlxmbhbd"));
          _230_v59 = (_230_v59).Merge((_230_v59).Merge(_dafny.Map.Empty.slice().updateUnsafe(_137_v0,_146_v6)));
          let _231_v60;
          _231_v60 = _dafny.Set.fromElements(_142_v3);
          _231_v60 = _231_v60;
          let _232_v61;
          _232_v61 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_142_v3, _142_v3, _137_v0),_137_v0);
          let _233_v63;
          _233_v63 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm9(_145_globalState),_146_v6);
          let _234_v64;
          _234_v64 = _dafny.Set.fromElements(_module.D1.create_DC3(_233_v63));
          let _235_v65;
          _235_v65 = _module.D3.create_DC8((_220_v50)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length))], new BigNumber((_234_v64).length), _137_v0);
          let _236_v66;
          let _nw38 = Array((new BigNumber(2)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _232_v61;
          _nw38[(_dafny.ONE).toNumber()] = function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.Seq.of(_235_v65)).Elements) {
              let _237_v62 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(_235_v65), _237_v62)) {
                _coll4.push([_237_v62,_137_v0]);
              }
            }
            return _coll4;
          }();
          _236_v66 = _nw38;
          let _index33 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_236_v66).length));
          (_236_v66)[_index33] = _dafny.Map.Empty.slice().updateUnsafe(_235_v65,true);
          let _index34 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_236_v66).length));
          (_236_v66)[_index34] = (_232_v61).Merge(_dafny.Map.Empty.slice().updateUnsafe(_235_v65,_module.__default.fm2(_145_globalState)));
          _137_v0 = true;
          let _238_v67;
          _238_v67 = _module.D0.create_DC2(_137_v0, _137_v0, (_220_v50)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length))]);
          let _pat_let_tv11 = _137_v0;
          let _239_v68;
          let _nw39 = Array((new BigNumber(11)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _238_v67;
          _nw39[(_dafny.ONE).toNumber()] = _238_v67;
          _nw39[(new BigNumber(2)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(3)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(4)).toNumber()] = _module.D0.create_DC2(_137_v0, _137_v0, _module.__default.fm0(_145_globalState));
          _nw39[(new BigNumber(5)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(6)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(7)).toNumber()] = function (_pat_let16_0) {
            return function (_240_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_241_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_241_dt__update_hcf2_h0, (_240_dt__update__tmp_h0).dtor_cf3, (_240_dt__update__tmp_h0).dtor_cf4);
                }(_pat_let17_0);
              }(_pat_let_tv11);
            }(_pat_let16_0);
          }(_238_v67);
          _nw39[(new BigNumber(8)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(9)).toNumber()] = _238_v67;
          _nw39[(new BigNumber(10)).toNumber()] = function (_pat_let18_0) {
            return function (_242_dt__update__tmp_h1) {
              return function (_pat_let19_0) {
                return function (_243_dt__update_hcf3_h0) {
                  return _module.D0.create_DC2((_242_dt__update__tmp_h1).dtor_cf2, _243_dt__update_hcf3_h0, (_242_dt__update__tmp_h1).dtor_cf4);
                }(_pat_let19_0);
              }(false);
            }(_pat_let18_0);
          }(_module.__default.fm10(_137_v0, _dafny.Seq.UnicodeFromString("oysattfu"), _145_globalState));
          _239_v68 = _nw39;
          let _index35 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_239_v68).length));
          (_239_v68)[_index35] = _238_v67;
          let _index36 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_239_v68).length));
          let _rhs13 = (new BigNumber((_138_v1).length)).plus(_142_v3);
          let _rhs14 = new _dafny.CodePoint('e'.codePointAt(0));
          let _rhs15 = _module.__default.fm10(_137_v0, _146_v6, _145_globalState);
          let _rhs16 = (_dafny.ZERO).minus((_220_v50)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_220_v50).length))]);
          let _lhs14 = _145_globalState;
          let _lhs15 = _145_globalState;
          let _lhs16 = _239_v68;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_239_v68).length));
          _lhs14.f0 = _rhs13;
          _lhs15.f4 = _rhs14;
          _lhs16[_lhs17] = _rhs15;
          _142_v3 = _rhs16;
        }
        (_145_globalState).f0 = _142_v3;
      } else {
        (_145_globalState).f0 = new BigNumber(660);
        let _244_v69;
        let _nw40 = Array((new BigNumber(22)).toNumber()).fill(false);
        _244_v69 = _nw40;
        let _index37 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_244_v69).length));
        (_244_v69)[_index37] = !(_137_v0);
        let _index38 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_244_v69).length));
        (_244_v69)[_index38] = (_142_v3).isEqualTo(_142_v3);
        let _245_v70;
        _245_v70 = _dafny.Seq.of(_146_v6);
        let _index39 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_244_v69).length));
        let _rhs17 = !(_142_v3).isEqualTo(_142_v3);
        let _rhs18 = !(!(_142_v3).isEqualTo(_module.__default.safeDivisionInt(_module.__default.fm0(_145_globalState), new BigNumber((_245_v70).length))));
        let _lhs18 = _244_v69;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_244_v69).length));
        _137_v0 = _rhs17;
        _lhs18[_lhs19] = _rhs18;
        _137_v0 = (_142_v3).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), ((_246_v3, _247_v0, _248_v69) => function (_249_i7) {
          return _dafny.Map.Empty.slice().updateUnsafe(_246_v3,new BigNumber((_dafny.MultiSet.fromElements(_247_v0, (_248_v69)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_248_v69).length))])).cardinality()));
        })(_142_v3, _137_v0, _244_v69))).length));
        let _250_v71;
        _250_v71 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_146_v6,_244_v69));
        (_145_globalState).f0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_142_v3), new BigNumber(((_250_v71)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(576)), function (_251_i8) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length), new BigNumber((_250_v71).length))]).length));
      }
      let _252_v72;
      _252_v72 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_137_v0);
      _252_v72 = (_252_v72).Merge(_252_v72);
      let _253_v73;
      _253_v73 = _module.D0.create_DC2(_137_v0, _137_v0, _142_v3);
      let _pat_let_tv12 = _137_v0;
      let _pat_let_tv13 = _142_v3;
      _142_v3 = function (_source3) {
        if (_source3.is_DC1) {
          let _254___mcc_h10 = (_source3).cf1;
          let _255_cf1 = _254___mcc_h10;
          return _255_cf1;
        } else if (_source3.is_DC2) {
          let _256___mcc_h11 = (_source3).cf2;
          let _257___mcc_h12 = (_source3).cf3;
          let _258___mcc_h13 = (_source3).cf4;
          let _259_cf4 = _258___mcc_h13;
          let _260_cf3 = _257___mcc_h12;
          let _261_cf2 = _256___mcc_h11;
          return new BigNumber((_dafny.MultiSet.fromElements(_pat_let_tv12)).cardinality());
        } else {
          let _262___mcc_h14 = (_source3).cf0;
          let _263_cf0 = _262___mcc_h14;
          return (new BigNumber(667)).multipliedBy(_pat_let_tv13);
        }
      }(_253_v73);
      _137_v0 = _137_v0;
      let _264_v74;
      _264_v74 = _dafny.Map.Empty.slice().updateUnsafe(_208_v48,_146_v6);
      let _265_v75;
      _265_v75 = _module.D2.create_DC6(_dafny.Seq.Concat(_146_v6, (((_264_v74).contains(_module.D1.create_DC4((_dafny.ZERO).minus(new BigNumber((_146_v6).length)), _137_v0, _137_v0, _137_v0))) ? ((_264_v74).get(_module.D1.create_DC4((_dafny.ZERO).minus(new BigNumber((_146_v6).length)), _137_v0, _137_v0, _137_v0))) : (_146_v6))), _142_v3, _137_v0, _142_v3, _137_v0);
      _265_v75 = _265_v75;
      let _266_v76;
      _266_v76 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,(_module.D0.create_DC1(_142_v3)).dtor_cf1);
      let _267_v77;
      _267_v77 = _dafny.Seq.of(_142_v3);
      let _268_v78;
      _268_v78 = _dafny.Seq.of(_266_v76, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_146_v6).length),new BigNumber((_267_v77).length)), _266_v76, _266_v76, _266_v76);
      let _269_v79;
      _269_v79 = _dafny.Seq.of(_137_v0);
      let _270_v80;
      _270_v80 = _module.D3.create_DC8(_142_v3, _142_v3, false);
      let _271_v81;
      let _nw41 = new _module.C0();
      _nw41.__ctor(_270_v80);
      _271_v81 = _nw41;
      let _272_v82;
      _272_v82 = _dafny.Seq.of(_271_v81, _271_v81);
      let _273_v83;
      _273_v83 = _dafny.MultiSet.fromElements((_272_v82)[_module.__default.safeIndex(_142_v3, new BigNumber((_272_v82).length))]);
      let _274_v84;
      _274_v84 = _dafny.MultiSet.fromElements(_142_v3, new BigNumber((_273_v83).cardinality()), (_dafny.ZERO).minus(_142_v3));
      let _275_v85;
      let _nw42 = Array((new BigNumber(21)).toNumber());
      _nw42[(_dafny.ZERO).toNumber()] = _137_v0;
      _nw42[(_dafny.ONE).toNumber()] = _137_v0;
      _nw42[(new BigNumber(2)).toNumber()] = true;
      _nw42[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_268_v78, _268_v78);
      _nw42[(new BigNumber(4)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(5)).toNumber()] = !((_138_v1).contains(_137_v0));
      _nw42[(new BigNumber(6)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(7)).toNumber()] = !(_137_v0);
      _nw42[(new BigNumber(8)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(9)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(10)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(11)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ryigcb"), _146_v6, _146_v6)).cardinality())).isEqualTo(_142_v3);
      _nw42[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_269_v79).length))).IsProperSubsetOf(_274_v84);
      _nw42[(new BigNumber(13)).toNumber()] = (_137_v0) || (_137_v0);
      _nw42[(new BigNumber(14)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(15)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(16)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(17)).toNumber()] = (_137_v0) || (_137_v0);
      _nw42[(new BigNumber(18)).toNumber()] = _137_v0;
      _nw42[(new BigNumber(19)).toNumber()] = !(_137_v0);
      _nw42[(new BigNumber(20)).toNumber()] = !(_137_v0);
      _275_v85 = _nw42;
      let _index40 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length));
      (_275_v85)[_index40] = _137_v0;
      let _276_v86;
      _276_v86 = new _dafny.CodePoint('d'.codePointAt(0));
      let _277_v87;
      _277_v87 = _dafny.Map.Empty.slice().updateUnsafe(_276_v86,(_dafny.ZERO).minus((_271_v81).fm6(_137_v0, (_module.D2.create_DC6(_146_v6, _142_v3, (((_252_v72).contains(_142_v3)) ? ((_252_v72).get(_142_v3)) : (_137_v0)), _142_v3, (((_252_v72).contains(new BigNumber(540))) ? ((_252_v72).get(new BigNumber(540))) : (_137_v0)))).dtor_cf12, !(_137_v0), _145_globalState)));
      let _278_v88;
      _278_v88 = _dafny.MultiSet.fromElements(_137_v0);
      let _index41 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length));
      (_275_v85)[_index41] = ((((_277_v87).contains(_276_v86)) ? ((_277_v87).get(_276_v86)) : ((((_278_v88).contains(_137_v0)) ? ((_278_v88).get(_137_v0)) : (_142_v3))))).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("lthcc")).length));
      _module.__default.m0(_137_v0, (_142_v3).multipliedBy(_142_v3), _137_v0, _142_v3, _145_globalState);
      let _279_v89;
      _279_v89 = _dafny.Map.Empty.slice().updateUnsafe(_270_v80,_146_v6);
      let _280_v90;
      _280_v90 = _module.D4.create_DC9(_dafny.Map.Empty.slice().updateUnsafe((_271_v81).f14,_146_v6));
      _279_v89 = (_280_v90).dtor_cf20;
      let _281_v91;
      _281_v91 = _dafny.MultiSet.fromElements(_269_v79);
      let _rhs19 = new BigNumber(((_281_v91).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(_137_v0), _269_v79))).cardinality());
      let _rhs20 = _267_v77;
      let _rhs21 = (_module.__default.fm11(_276_v86, _145_globalState)).dtor_cf14;
      let _lhs20 = _145_globalState;
      let _lhs21 = _145_globalState;
      _lhs20.f0 = _rhs19;
      _267_v77 = _rhs20;
      _lhs21.f0 = _rhs21;
      let _282_i9;
      _282_i9 = _dafny.ZERO;
      L0: {
        while (_137_v0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_282_i9)) {
              break L0;
            }
            _282_i9 = (_282_i9).plus(_dafny.ONE);
            _module.__default.m0((_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))], _142_v3, (_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))], _142_v3, _145_globalState);
            let _283_v92;
            _283_v92 = _module.D0.create_DC0(_dafny.Seq.of((_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))], _137_v0));
            let _source4 = _283_v92;
            if (_source4.is_DC1) {
              let _284___mcc_h15 = (_source4).cf1;
              let _285_cf1 = _284___mcc_h15;
              let _286_v93;
              let _nw43 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
              _286_v93 = _nw43;
              let _index42 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_286_v93).length));
              (_286_v93)[_index42] = _274_v84;
              let _index43 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length));
              let _index44 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_286_v93).length));
              let _rhs22 = _285_cf1;
              let _rhs23 = _137_v0;
              let _rhs24 = new _dafny.CodePoint('o'.codePointAt(0));
              let _rhs25 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_267_v77, _267_v77))).Intersect(_module.__default.fm12(_276_v86, (_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))], _285_cf1, _145_globalState));
              let _rhs26 = _272_v82;
              let _lhs22 = _275_v85;
              let _lhs23 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length));
              let _lhs24 = _145_globalState;
              let _lhs25 = _286_v93;
              let _lhs26 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_286_v93).length));
              _142_v3 = _rhs22;
              _lhs22[_lhs23] = _rhs23;
              _lhs24.f9 = _rhs24;
              _lhs25[_lhs26] = _rhs25;
              _272_v82 = _rhs26;
              let _287_v94;
              let _nw44 = new _module.C1();
              _nw44.__ctor();
              _287_v94 = _nw44;
              let _288_v95;
              let _nw45 = new _module.C0();
              _nw45.__ctor((_271_v81).f14);
              _288_v95 = _nw45;
              let _289_v96;
              let _nw46 = new _module.C0();
              _nw46.__ctor((_271_v81).f14);
              _289_v96 = _nw46;
            } else if (_source4.is_DC2) {
              let _290___mcc_h16 = (_source4).cf2;
              let _291___mcc_h17 = (_source4).cf3;
              let _292___mcc_h18 = (_source4).cf4;
              let _293_cf4 = _292___mcc_h18;
              let _294_cf3 = _291___mcc_h17;
              let _295_cf2 = _290___mcc_h16;
              let _296_v97;
              let _nw47 = Array((_dafny.ONE).toNumber()).fill([]);
              _296_v97 = _nw47;
              let _index45 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_296_v97).length));
              (_296_v97)[_index45] = _275_v85;
              let _297_v98;
              let _nw48 = Array((new BigNumber(2)).toNumber());
              _nw48[(_dafny.ZERO).toNumber()] = _146_v6;
              _nw48[(_dafny.ONE).toNumber()] = _146_v6;
              _297_v98 = _nw48;
              let _index46 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_297_v98).length));
              (_297_v98)[_index46] = _dafny.Seq.Concat(_146_v6, _146_v6);
              let _298_v99;
              let _init7 = ((_299_cf4) => function (_300_i10) {
                return (_300_i10).minus(_299_cf4);
              })(_293_cf4);
              let _nw49 = Array((new BigNumber(24)).toNumber());
              for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw49.length); _i0_7++) {
                _nw49[_i0_7] = _init7(new BigNumber(_i0_7));
              }
              _298_v99 = _nw49;
              let _index47 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_298_v99).length));
              (_298_v99)[_index47] = new BigNumber(-345);
              let _index48 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_296_v97).length));
              let _index49 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_297_v98).length));
              let _index50 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_298_v99).length));
              let _rhs27 = _275_v85;
              let _rhs28 = _146_v6;
              let _rhs29 = _293_cf4;
              let _rhs30 = (_293_cf4).isLessThan(_module.__default.safeModuloInt(_142_v3, _142_v3));
              let _lhs27 = _296_v97;
              let _lhs28 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_296_v97).length));
              let _lhs29 = _297_v98;
              let _lhs30 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_297_v98).length));
              let _lhs31 = _298_v99;
              let _lhs32 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_298_v99).length));
              _lhs27[_lhs28] = _rhs27;
              _lhs29[_lhs30] = _rhs28;
              _lhs31[_lhs32] = _rhs29;
              _295_cf2 = _rhs30;
              let _index51 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_298_v99).length));
              let _rhs31 = _module.__default.fm0(_145_globalState);
              let _rhs32 = !((_module.D4.create_DC12(_137_v0, _268_v78, (_175_v31).dtor_cf1, _275_v85, _137_v0)).dtor_cf26);
              let _lhs33 = _298_v99;
              let _lhs34 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_298_v99).length));
              _lhs33[_lhs34] = _rhs31;
              _294_cf3 = _rhs32;
              let _301_v100;
              let _nw50 = Array((new BigNumber(16)).toNumber());
              _301_v100 = _nw50;
              let _302_v101;
              _302_v101 = _dafny.Map.Empty.slice().updateUnsafe(_294_cf3,_301_v100);
              _302_v101 = _dafny.Map.Empty.slice().updateUnsafe((_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))],_301_v100);
              _142_v3 = _293_cf4;
            } else {
              let _303___mcc_h19 = (_source4).cf0;
              let _304_cf0 = _303___mcc_h19;
              let _305_v102;
              let _nw51 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _305_v102 = _nw51;
              let _306_v103;
              _306_v103 = _dafny.Map.Empty.slice().updateUnsafe(_137_v0,_305_v102);
              let _307_v104;
              _307_v104 = _module.D3.create_DC7((((_306_v103).contains((_269_v79)[_module.__default.safeIndex(_142_v3, new BigNumber((_269_v79).length))])) ? ((_306_v103).get((_269_v79)[_module.__default.safeIndex(_142_v3, new BigNumber((_269_v79).length))])) : (_305_v102)));
              _307_v104 = _307_v104;
              let _308_v105;
              let _nw52 = new _module.C1();
              _nw52.__ctor();
              _308_v105 = _nw52;
              let _309_v106;
              _309_v106 = _dafny.Seq.of(_308_v105);
              let _index52 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length));
              (_275_v85)[_index52] = (_module.__default.safeDivisionInt(new BigNumber((_304_cf0).length), _142_v3)).isLessThan(new BigNumber((_dafny.Seq.Concat(_309_v106, _dafny.Seq.of(_308_v105, _308_v105))).length));
              (_145_globalState).f0 = _module.__default.safeModuloInt(_142_v3, _142_v3);
              (_145_globalState).f0 = _142_v3;
            }
            let _310_v107;
            let _nw53 = new _module.C1();
            _nw53.__ctor();
            _310_v107 = _nw53;
            let _311_v108;
            _311_v108 = _dafny.Seq.of(_310_v107, _310_v107, _310_v107, _310_v107);
            (_145_globalState).f0 = (_142_v3).plus(new BigNumber((_dafny.Seq.Concat(_311_v108, _311_v108)).length));
            (_145_globalState).f0 = (_142_v3).multipliedBy(_142_v3);
          }
        }
      }
      let _312_v109;
      _312_v109 = _dafny.Map.Empty.slice().updateUnsafe(_146_v6,(_275_v85)[_module.__default.safeIndex(new BigNumber(938), new BigNumber((_275_v85).length))]);
      let _313_v110;
      _313_v110 = _dafny.Map.Empty.slice().updateUnsafe(_138_v1,(((_266_v76).contains(_142_v3)) ? ((_266_v76).get(_142_v3)) : (_142_v3)));
      _312_v109 = (_312_v109).update(_146_v6, !(_313_v110).contains(_dafny.Set.fromElements(_137_v0)));
      process.stdout.write(_dafny.toString(_137_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_139_v2)[_dafny.ZERO], _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_142_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_144_v5)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState).f1).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_145_globalState).f2)[_dafny.ZERO], _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_145_globalState.f3)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-730)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_145_globalState).f5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_145_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_146_v6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_175_v31).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v48).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v48).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v48).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v48).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v72).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v73).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v73).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v73).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_264_v74).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC4(new BigNumber(423), true, true, true),_dafny.Seq.UnicodeFromString("whhg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_265_v75).dtor_cf11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v75).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v75).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v75).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v75).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_266_v76).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_267_v77, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_268_v78, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(4),_dafny.ONE), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE), _dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_269_v79, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v80).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v80).dtor_cf18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v80).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_271_v81).f14).dtor_cf17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_271_v81).f14).dtor_cf18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_271_v81).f14).dtor_cf19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_272_v82).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_273_v83).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_274_v84).equals(_dafny.MultiSet.fromElements(_dafny.ONE, _dafny.ONE, new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_275_v85)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_276_v86));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v87).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(615)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_278_v88).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v89).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.ONE, _dafny.ONE, false),_dafny.Seq.UnicodeFromString("whhg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_280_v90).dtor_cf20).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_dafny.ONE, _dafny.ONE, false),_dafny.Seq.UnicodeFromString("whhg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v91).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_282_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v109).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("whhg"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_313_v110).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),_dafny.ONE))));
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
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
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
        return other.$tag === 1 && this.cf2 === other.cf2 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
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
    static create_DC4(cf6, cf7, cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, false, false, false);
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
    static create_DC6(cf11, cf12, cf13, cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC5(cf10) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + this.cf11.toVerbatimString(true) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false, _dafny.ZERO, false);
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
    static create_DC8(cf17, cf18, cf19) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf16 === other.cf16;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC10(cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC11(cf23, cf24, cf25) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC12(cf26, cf27, cf28, cf29, cf30) {
      let $dt = new D4(2);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC9(cf20) {
      let $dt = new D4(3);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && this.cf30 === other.cf30;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(_dafny.Seq.of(), false);
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
    static create_DC13(cf31) {
      let $dt = new D5(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf32) {
      let $dt = new D6(0);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32);
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
          return D6.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f3 = [];
      this.f4 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f9 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f1 = _dafny.Set.Empty;
      this._f2 = [];
      this._f5 = _dafny.Seq.UnicodeFromString("");
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f10 = false;
      this._f11 = _dafny.ZERO;
      this._f12 = false;
      this._f13 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f14 = _module.D3.Default();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14) {
      let _this = this;
      (_this)._f14 = f14;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(new BigNumber(-847), new BigNumber((_dafny.Seq.UnicodeFromString("uqfinko")).length))).length), new BigNumber(15))).minus(new BigNumber(615));
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
    fm4(globalState) {
      let _this = this;
      return ((_module.D1.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true, false, false),_dafny.Seq.UnicodeFromString("saf")))).dtor_cf5).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),false)).Keys.Elements) {
          let _314_v0 = _compr_5;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),false)).contains(_314_v0)) {
            _coll5.push([_314_v0,_dafny.Seq.UnicodeFromString("byiwjxmen")]);
          }
        }
        return _coll5;
      }());
    };
    fm5(globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_315_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }),_dafny.Seq.UnicodeFromString("kipi"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vbvdrrtth"),_dafny.Seq.UnicodeFromString("n")))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC5(_dafny.Seq.UnicodeFromString("rt"))).dtor_cf10,_dafny.Seq.UnicodeFromString("goeperti")));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.Set.Empty;
      let r3 = _dafny.ZERO;
      let _316_v0;
      _316_v0 = new BigNumber(52);
      let _317_v1;
      _317_v1 = _module.D0.create_DC1(_316_v0);
      let _pat_let_tv14 = _316_v0;
      r3 = (function (_pat_let20_0) {
        return function (_318_dt__update__tmp_h0) {
          return function (_pat_let21_0) {
            return function (_319_dt__update_hcf1_h0) {
              return _module.D0.create_DC1(_319_dt__update_hcf1_h0);
            }(_pat_let21_0);
          }(_pat_let_tv14);
        }(_pat_let20_0);
      }(_317_v1)).dtor_cf1;
      let _hi3 = (_316_v0).plus(_316_v0);
      for (let _320_i0 = new BigNumber((p1).length); _320_i0.isLessThan(_hi3); _320_i0 = _320_i0.plus(_dafny.ONE)) {
        r1 = !(p2);
        let _rhs33 = p2;
        let _rhs34 = _320_i0;
        let _rhs35 = p2;
        r1 = _rhs33;
        _316_v0 = _rhs34;
        r1 = _rhs35;
        r1 = p2;
        let _321_v2;
        _321_v2 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(p1, p1)).length));
        let _rhs36 = (_321_v2)[_module.__default.safeIndex(_316_v0, new BigNumber((_321_v2).length))];
        let _rhs37 = _module.__default.fm2(globalState);
        let _lhs35 = globalState;
        _lhs35.f0 = _rhs36;
        r1 = _rhs37;
      }
      let _322_v3;
      let _nw54 = Array((new BigNumber(7)).toNumber()).fill(false);
      _322_v3 = _nw54;
      let _index53 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
      (_322_v3)[_index53] = p2;
      let _index54 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
      (_322_v3)[_index54] = !(!(p2)) || (p2);
      let _index55 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
      let _rhs38 = !((new BigNumber((p1).length)).isLessThanOrEqualTo(_316_v0));
      let _rhs39 = p2;
      let _lhs36 = _322_v3;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
      r1 = _rhs38;
      _lhs36[_lhs37] = _rhs39;
      let _hi4 = _316_v0;
      for (let _323_i1 = (_316_v0).minus(_316_v0); _323_i1.isLessThan(_hi4); _323_i1 = _323_i1.plus(_dafny.ONE)) {
        let _324_v4;
        _324_v4 = _module.D0.create_DC2((_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))], (_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))], _323_i1);
        let _325_v5;
        _325_v5 = _dafny.Seq.of(_324_v4, _324_v4);
        let _326_v6;
        _326_v6 = _dafny.Map.Empty.slice().updateUnsafe(_316_v0,new BigNumber(178));
        let _327_v7;
        _327_v7 = _dafny.MultiSet.fromElements(_326_v6, _326_v6);
        let _328_v8;
        _328_v8 = _dafny.Seq.of(p2, (_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))]);
        let _329_v9;
        _329_v9 = _dafny.Set.fromElements(_328_v8);
        let _330_v10;
        _330_v10 = _dafny.Seq.of(new BigNumber((_325_v5).length), (_dafny.ZERO).minus(new BigNumber((_327_v7).cardinality())), _323_i1, new BigNumber((_329_v9).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
        let _rhs40 = _316_v0;
        let _rhs41 = _dafny.areEqual(_330_v10, _dafny.Seq.Concat(_330_v10, _330_v10));
        let _lhs38 = _322_v3;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
        r3 = _rhs40;
        _lhs38[_lhs39] = _rhs41;
        if (_module.__default.fm2(globalState)) {
          let _331_v11;
          _331_v11 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.fm2(globalState));
          let _332_v12;
          _332_v12 = _dafny.MultiSet.fromElements(_module.__default.fm2(globalState), p2);
          _331_v11 = (_331_v11).update((_dafny.MultiSet.fromElements((_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))])).IsProperSubsetOf(_332_v12), (_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))]);
          _331_v11 = (_331_v11).update(p2, p2);
          let _333_v13;
          let _nw55 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _333_v13 = _nw55;
          let _334_v14;
          _334_v14 = _module.D3.create_DC7(_333_v13);
          let _335_v15;
          _335_v15 = _dafny.Map.Empty.slice().updateUnsafe((_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))],(_334_v14).dtor_cf16);
          let _336_v16;
          _336_v16 = _dafny.Seq.of(_333_v13, _333_v13, _333_v13, _333_v13);
          _335_v15 = (_335_v15).update((_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))], (_336_v16)[_module.__default.safeIndex(_323_i1, new BigNumber((_336_v16).length))]);
          let _337_v17;
          let _nw56 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _337_v17 = _nw56;
          let _index57 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_337_v17).length));
          (_337_v17)[_index57] = _dafny.Seq.Concat(_336_v16, _336_v16);
          let _338_v18;
          _338_v18 = _dafny.Map.Empty.slice().updateUnsafe(_316_v0,p2);
          let _339_v19;
          _339_v19 = _dafny.Seq.of(_338_v18, _338_v18);
          let _index58 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_337_v17).length));
          (_337_v17)[_index58] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_336_v16, _336_v16), _336_v16), _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_339_v19).length), new BigNumber((p1).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_336_v16, _336_v16), _336_v16)).length)), (((_335_v15).contains(p2)) ? ((_335_v15).get(p2)) : (_333_v13)));
          _331_v11 = (_331_v11).update(p2, p2);
        } else {
          _316_v0 = new BigNumber((_328_v8).length);
          let _340_v20;
          _340_v20 = new _dafny.CodePoint('q'.codePointAt(0));
          (globalState).f4 = _340_v20;
          _326_v6 = (_326_v6).update(_323_i1, _316_v0);
          let _index59 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length));
          (_322_v3)[_index59] = _module.__default.fm2(globalState);
          let _341_v21;
          let _nw57 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _341_v21 = _nw57;
          let _342_v22;
          _342_v22 = _dafny.Seq.of(_341_v21, _341_v21);
          let _343_v23;
          _343_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC7((_342_v22)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_342_v22).length))]),_316_v0);
          let _344_v24;
          let _nw58 = Array((_dafny.ONE).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _343_v23;
          _344_v24 = _nw58;
          let _index60 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_344_v24).length));
          (_344_v24)[_index60] = (_343_v23).Merge(_343_v23);
          let _index61 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_344_v24).length));
          (_344_v24)[_index61] = _343_v23;
        }
        (globalState).f0 = _316_v0;
        let _345_v25;
        _345_v25 = _module.D3.create_DC8(_323_i1, _316_v0, (_322_v3)[_module.__default.safeIndex(new BigNumber(729), new BigNumber((_322_v3).length))]);
        let _346_v26;
        let _nw59 = new _module.C0();
        _nw59.__ctor(_345_v25);
        _346_v26 = _nw59;
      }
      (globalState).f0 = _module.__default.safeDivisionInt(_316_v0, (_316_v0).multipliedBy(new BigNumber((p1).length)));
      r0 = (_dafny.ZERO).minus(new BigNumber((p1).length));
      r1 = !(p2);
      let _347_v27;
      _347_v27 = _dafny.Set.fromElements(_316_v0);
      r2 = (_347_v27).Union(_347_v27);
      r3 = _module.__default.safeDivisionInt(((_dafny.ZERO).minus(_316_v0)).minus(_316_v0), _316_v0);
      return [r0, r1, r2, r3];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
