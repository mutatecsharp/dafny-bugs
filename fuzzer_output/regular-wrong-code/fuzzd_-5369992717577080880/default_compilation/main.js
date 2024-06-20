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
      if ((_module.D4.create_DC15(new BigNumber((_dafny.Seq.UnicodeFromString("kuw")).length), true)).dtor_cf24) {
        return _module.D0.create_DC0(new BigNumber(-581));
      } else {
        return _module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(412),true)).length)));
      }
    };
    static fm1(p0, p1, globalState) {
      return _module.D0.create_DC2(true, ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(16))).length)))).plus((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of _dafny.IntegerRange(new BigNumber(904), new BigNumber(148))) {
    let _0_v0 = _compr_0;
    if (((new BigNumber(904)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(148)))) {
      _coll0.push([(_0_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(635),true))).length)),true]);
    }
  }
  return _coll0;
}()).length))), (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dh")).length)))).Union(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-400))))));
    };
    static fm2(globalState) {
      return !(false) || (false);
    };
    static fm3(p0, p1, p2, globalState) {
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, (_module.D0.create_DC2(true, new BigNumber(163), _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(713)), function (_1_i0) {
  return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-701),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true)).length))).length));
})).length))))).dtor_cf2)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("gxdaqqlt")).length))));
    };
    static fm10(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(true), true)), _dafny.Seq.of((_module.D5.create_DC18(new BigNumber(-672), !(!(true)))).dtor_cf28, false));
    };
    static fm11(p0, p1, globalState) {
      if ((_dafny.MultiSet.fromElements(true)).equals(_dafny.MultiSet.fromElements(false, true))) {
        return _dafny.Seq.UnicodeFromString("yo");
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jjeyx"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(600)), function (_2_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(false)).Intersect((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(!((_module.D0.create_DC1(true)).dtor_cf1))));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(987)), function (_3_i0) {
        return new BigNumber(398);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xrmaqmxx")).length)))), ((!(false)) ? (_dafny.Seq.of(new BigNumber(42), new BigNumber(627))) : (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), function (_4_i1) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        })).Elements) {
          let _5_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), function (_4_i1) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), _5_v0)) {
            _coll1.push([_5_v0,true]);
          }
        }
        return _coll1;
      }())).length), new BigNumber(-641), new BigNumber(944)))));
    };
    static fm14(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(true)), _dafny.Map.Empty.slice().updateUnsafe(false,(_module.D0.create_DC1((_module.D2.create_DC8(true, new BigNumber(488), false)).dtor_cf7)).dtor_cf1)), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,!(true))));
    };
    static fm15(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_6_i0) {
        return _dafny.Seq.UnicodeFromString("v");
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), function (_7_i1) {
        return _dafny.Seq.UnicodeFromString("hchl");
      })), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-593)), function (_8_i2) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(278)), function (_9_i3) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        });
      }), _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(990)), function (_10_i4) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }))));
    };
    static fm16(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("mlvvv")).length),false)).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(480), new BigNumber(-305))) {
          let _11_v0 = _compr_2;
          if (((new BigNumber(480)).isLessThanOrEqualTo(_11_v0)) && ((_11_v0).isLessThan(new BigNumber(-305)))) {
            _coll2.push([(_11_v0).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length)),true]);
          }
        }
        return _coll2;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, !(true), true, false)).length),!(true)));
    };
    static fm17(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber(837)), (_dafny.ZERO).minus((new BigNumber(437)).plus(new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-549), new BigNumber(598))) {
          let _12_v0 = _compr_3;
          if (((new BigNumber(-549)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(598)))) {
            _coll3.push([(_12_v0).plus((_dafny.ZERO).minus(new BigNumber(-485))),false]);
          }
        }
        return _coll3;
      }()).length))), new BigNumber(368), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(!(false), true)).length)));
    };
    static fm18(globalState) {
      if ((new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length),new BigNumber(356))).Keys.Elements) {
            let _13_v1 = _compr_5;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length),new BigNumber(356))).contains(_13_v1)) {
              _coll5.add(_module.__default.safeDivisionInt(_13_v1, new BigNumber(-91)));
            }
          }
          return _coll5;
        }()).Elements) {
          let _14_v0 = _compr_4;
          if ((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length),new BigNumber(356))).Keys.Elements) {
              let _15_v1 = _compr_6;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("d")).length),new BigNumber(356))).contains(_15_v1)) {
                _coll6.add(_module.__default.safeDivisionInt(_15_v1, new BigNumber(-91)));
              }
            }
            return _coll6;
          }()).contains(_14_v0)) {
            _coll4.push([(_14_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(false, !(true), false, true)).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(972)), function (_16_i0) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            })]);
          }
        }
        return _coll4;
      }()).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("delwixswv")).length)))) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }
    };
    static fm19(p0, globalState) {
      return _dafny.Seq.Concat((_module.D7.create_DC25(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-252)), function (_17_i0) {
  return _module.D3.create_DC9(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))));
}))).dtor_cf41, _dafny.Seq.of(_module.D3.create_DC9(_dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), function (_18_i1) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})), _module.D3.create_DC9(_dafny.Seq.of(new _dafny.CodePoint('u'.codePointAt(0))))));
    };
    static fm20(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(599)), function (_19_i0) {
        return function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dxqwlk")).length),new _dafny.CodePoint('w'.codePointAt(0)))).Keys.Elements) {
            let _20_v0 = _compr_7;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dxqwlk")).length),new _dafny.CodePoint('w'.codePointAt(0)))).contains(_20_v0)) {
              _coll7.push([(_20_v0).multipliedBy(new BigNumber(757)),true]);
            }
          }
          return _coll7;
        }();
      }));
    };
    static fm21(p0, globalState) {
      return _module.D6.create_DC20((_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let _hi0 = p0;
      for (let _21_i0 = p0; _21_i0.isLessThan(_hi0); _21_i0 = _21_i0.plus(_dafny.ONE)) {
        let _22_v0;
        _22_v0 = new _dafny.CodePoint('x'.codePointAt(0));
        let _23_v1;
        _23_v1 = _dafny.Map.Empty.slice().updateUnsafe(_22_v0,_22_v0);
        let _24_v2;
        _24_v2 = _dafny.Map.Empty.slice().updateUnsafe(_23_v1,p2);
        let _index0 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length));
        (p1)[_index0] = p0;
        let _25_v3;
        _25_v3 = _dafny.Seq.UnicodeFromString("nodnggll");
        let _index1 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length));
        let _rhs0 = _24_v2;
        let _rhs1 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(246)), ((_26_v0) => function (_27_i1) {
          return _26_v0;
        })(_22_v0))).length);
        let _rhs2 = new BigNumber((_25_v3).length);
        let _lhs0 = p1;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length));
        _24_v2 = _rhs0;
        r0 = _rhs1;
        _lhs0[_lhs1] = _rhs2;
        let _28_v4;
        _28_v4 = _module.D5.create_DC18(p0, p2);
        let _pat_let_tv0 = p2;
        let _pat_let_tv1 = p0;
        let _pat_let_tv2 = p2;
        let _index2 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length));
        let _rhs3 = p0;
        let _rhs4 = _module.__default.safeDivisionInt((_28_v4).dtor_cf27, (p1)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length))]);
        let _rhs5 = function (_pat_let0_0) {
          return function (_33_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_34_dt__update_hcf28_h1) {
                return _module.D5.create_DC18((_33_dt__update__tmp_h2).dtor_cf27, _34_dt__update_hcf28_h1);
              }(_pat_let5_0);
            }(_pat_let_tv2);
          }(_pat_let0_0);
        }(function (_pat_let1_0) {
          return function (_31_dt__update__tmp_h1) {
            return function (_pat_let4_0) {
              return function (_32_dt__update_hcf27_h0) {
                return _module.D5.create_DC18(_32_dt__update_hcf27_h0, (_31_dt__update__tmp_h1).dtor_cf28);
              }(_pat_let4_0);
            }((_dafny.ZERO).minus(_pat_let_tv1));
          }(_pat_let1_0);
        }(function (_pat_let2_0) {
          return function (_29_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_30_dt__update_hcf28_h0) {
                return _module.D5.create_DC18((_29_dt__update__tmp_h0).dtor_cf27, _30_dt__update_hcf28_h0);
              }(_pat_let3_0);
            }(_pat_let_tv0);
          }(_pat_let2_0);
        }(_28_v4)));
        let _lhs2 = p1;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length));
        _lhs2[_lhs3] = _rhs3;
        r0 = _rhs4;
        _28_v4 = _rhs5;
        let _35_v5;
        let _nw0 = new _module.C1();
        _nw0.__ctor(p2, (p1)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((p1).length))], p2);
        _35_v5 = _nw0;
        let _36_v6;
        let _nw1 = Array((new BigNumber(5)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = _35_v5;
        _nw1[(_dafny.ONE).toNumber()] = _35_v5;
        _nw1[(new BigNumber(2)).toNumber()] = _35_v5;
        _nw1[(new BigNumber(3)).toNumber()] = _35_v5;
        _nw1[(new BigNumber(4)).toNumber()] = _35_v5;
        _36_v6 = _nw1;
        let _index3 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_36_v6).length));
        (_36_v6)[_index3] = _35_v5;
        let _index4 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_36_v6).length));
        (_36_v6)[_index4] = _35_v5;
        (globalState).f3 = !(_dafny.Seq.contains(_25_v3, _22_v0));
      }
      let _hi1 = p0;
      for (let _37_i2 = p0; _37_i2.isLessThan(_hi1); _37_i2 = _37_i2.plus(_dafny.ONE)) {
        let _38_v7;
        let _nw2 = Array((new BigNumber(14)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = p2;
        _nw2[(_dafny.ONE).toNumber()] = true;
        _nw2[(new BigNumber(2)).toNumber()] = p2;
        _nw2[(new BigNumber(3)).toNumber()] = p2;
        _nw2[(new BigNumber(4)).toNumber()] = _module.__default.fm2(globalState);
        _nw2[(new BigNumber(5)).toNumber()] = p2;
        _nw2[(new BigNumber(6)).toNumber()] = p2;
        _nw2[(new BigNumber(7)).toNumber()] = p2;
        _nw2[(new BigNumber(8)).toNumber()] = p2;
        _nw2[(new BigNumber(9)).toNumber()] = p2;
        _nw2[(new BigNumber(10)).toNumber()] = p2;
        _nw2[(new BigNumber(11)).toNumber()] = true;
        _nw2[(new BigNumber(12)).toNumber()] = p2;
        _nw2[(new BigNumber(13)).toNumber()] = p2;
        _38_v7 = _nw2;
        let _39_v8;
        _39_v8 = _dafny.Seq.of(_38_v7, _38_v7);
        let _40_v9;
        let _nw3 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
        _40_v9 = _nw3;
        let _41_v10;
        _41_v10 = _dafny.Map.Empty.slice().updateUnsafe(_39_v8,_40_v9);
        _41_v10 = (_41_v10).update(_dafny.Seq.Concat(_39_v8, _39_v8), _40_v9);
        let _42_v11;
        _42_v11 = _dafny.Seq.of(_37_i2);
        let _43_v12;
        _43_v12 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0), p0, (_42_v11)[_module.__default.safeIndex(_37_i2, new BigNumber((_42_v11).length))], p0, _37_i2);
        let _44_v13;
        _44_v13 = _module.D3.create_DC10(new BigNumber((_43_v12).length), p1, p2);
        let _source0 = _44_v13;
        if (_source0.is_DC10) {
          let _45___mcc_h0 = (_source0).cf11;
          let _46___mcc_h1 = (_source0).cf12;
          let _47___mcc_h2 = (_source0).cf13;
          let _48_cf13 = _47___mcc_h2;
          let _49_cf12 = _46___mcc_h1;
          let _50_cf11 = _45___mcc_h0;
          let _51_v14;
          _51_v14 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _index5 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_49_cf12).length));
          (_49_cf12)[_index5] = ((p2) ? (_50_cf11) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_51_v14).length))).length)));
          let _index6 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_49_cf12).length));
          (_49_cf12)[_index6] = ((false) ? ((_37_i2).multipliedBy(p0)) : (p0));
          let _52_v15;
          let _nw4 = new _module.C2();
          _nw4.__ctor(true, _48_cf13);
          _52_v15 = _nw4;
          _52_v15 = _52_v15;
          let _index7 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_38_v7).length));
          (_38_v7)[_index7] = _52_v15.f10;
          let _index8 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_38_v7).length));
          (_38_v7)[_index8] = p2;
          let _53_v16;
          let _nw5 = Array((new BigNumber(14)).toNumber()).fill(false);
          _53_v16 = _nw5;
          let _54_v17;
          _54_v17 = _module.D0.create_DC1(_52_v15.f10);
          let _55_v18;
          _55_v18 = _dafny.Set.fromElements((_54_v17).dtor_cf1, _52_v15.f10);
          let _56_v19;
          let _57_v20;
          let _58_v21;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_52_v15).m3((p0).minus((_49_cf12)[_module.__default.safeIndex(new BigNumber(352), new BigNumber((_49_cf12).length))]), _53_v16, true, _55_v18, globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _56_v19 = _out0;
          _57_v20 = _out1;
          _58_v21 = _out2;
        } else if (_source0.is_DC11) {
          let _59___mcc_h3 = (_source0).cf14;
          let _60___mcc_h4 = (_source0).cf15;
          let _61___mcc_h5 = (_source0).cf16;
          let _62___mcc_h6 = (_source0).cf17;
          let _63_cf17 = _62___mcc_h6;
          let _64_cf16 = _61___mcc_h5;
          let _65_cf15 = _60___mcc_h4;
          let _66_cf14 = _59___mcc_h3;
          let _67_v23;
          _67_v23 = _dafny.MultiSet.fromElements(function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-756), new BigNumber(716))) {
              let _68_v22 = _compr_8;
              if (((new BigNumber(-756)).isLessThanOrEqualTo(_68_v22)) && ((_68_v22).isLessThan(new BigNumber(716)))) {
                _coll8.push([_module.__default.safeModuloInt(_68_v22, new BigNumber((_65_cf15).length)),!(p2)]);
              }
            }
            return _coll8;
          }());
          let _69_v24;
          _69_v24 = _module.D2.create_DC8((_63_cf17).f13, _64_cf16, (_63_cf17).f13);
          let _pat_let_tv3 = _63_cf17;
          let _pat_let_tv4 = _64_cf16;
          let _pat_let_tv5 = _63_cf17;
          let _rhs6 = (function (_pat_let6_0) {
            return function (_73_dt__update__tmp_h4) {
              return function (_pat_let10_0) {
                return function (_74_dt__update_hcf7_h1) {
                  return _module.D2.create_DC8(_74_dt__update_hcf7_h1, (_73_dt__update__tmp_h4).dtor_cf8, (_73_dt__update__tmp_h4).dtor_cf9);
                }(_pat_let10_0);
              }((_pat_let_tv5).f13);
            }(_pat_let6_0);
          }(function (_pat_let7_0) {
            return function (_70_dt__update__tmp_h3) {
              return function (_pat_let8_0) {
                return function (_71_dt__update_hcf7_h0) {
                  return function (_pat_let9_0) {
                    return function (_72_dt__update_hcf8_h0) {
                      return _module.D2.create_DC8(_71_dt__update_hcf7_h0, _72_dt__update_hcf8_h0, (_70_dt__update__tmp_h3).dtor_cf9);
                    }(_pat_let9_0);
                  }(_pat_let_tv4);
                }(_pat_let8_0);
              }((_pat_let_tv3).f13);
            }(_pat_let7_0);
          }(_69_v24))).dtor_cf7;
          let _rhs7 = _module.__default.fm20((_63_cf17).f13, _42_v11, (_64_cf16).isEqualTo(_63_cf17.f14), globalState);
          let _rhs8 = _dafny.Seq.IsProperPrefixOf(_65_cf15, _dafny.Seq.Concat(_65_cf15, _65_cf15));
          let _lhs4 = globalState;
          let _lhs5 = globalState;
          _lhs4.f3 = _rhs6;
          _67_v23 = _rhs7;
          _lhs5.f3 = _rhs8;
          let _75_v25;
          let _nw6 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _75_v25 = _nw6;
          _75_v25 = p1;
          _64_cf16 = (((((_63_cf17).f13) ? (p2) : ((_63_cf17).f13))) ? (_37_i2) : (_66_cf14));
          let _76_v26;
          _76_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_63_cf17).f13);
          let _77_v27;
          _77_v27 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_76_v26).length)).isLessThanOrEqualTo(_66_cf14),!(p2));
          let _78_v28;
          _78_v28 = _dafny.MultiSet.fromElements(p2);
          _77_v27 = (_77_v27).update(!(_78_v28).contains(p2), (_44_v13).dtor_cf13);
        } else if (_source0.is_DC9) {
          let _79___mcc_h7 = (_source0).cf10;
          let _80_cf10 = _79___mcc_h7;
          let _81_v29;
          _81_v29 = _dafny.Set.fromElements(p2, p2);
          (globalState).f0 = !(((p0).multipliedBy(new BigNumber((_81_v29).length))).isLessThanOrEqualTo((_42_v11)[_module.__default.safeIndex(p0, new BigNumber((_42_v11).length))]));
          let _82_v30;
          let _init0 = ((_83_p2) => function (_84_i3) {
            return ((_83_p2) ? (new _dafny.CodePoint('c'.codePointAt(0))) : (new _dafny.CodePoint('x'.codePointAt(0))));
          })(p2);
          let _nw7 = Array((new BigNumber(27)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
            _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _82_v30 = _nw7;
          let _index9 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_82_v30).length));
          (_82_v30)[_index9] = new _dafny.CodePoint('q'.codePointAt(0));
          let _85_v31;
          _85_v31 = new _dafny.CodePoint('u'.codePointAt(0));
          let _index10 = _module.__default.safeIndex(new BigNumber(82), new BigNumber((_82_v30).length));
          (_82_v30)[_index10] = _85_v31;
          _40_v9 = _40_v9;
          let _index11 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((p1).length));
          (p1)[_index11] = p0;
          let _index12 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((p1).length));
          (p1)[_index12] = p0;
        } else {
          let _86___mcc_h8 = (_source0).cf18;
          let _87_cf18 = _86___mcc_h8;
          _38_v7 = _38_v7;
          let _index13 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_38_v7).length));
          (_38_v7)[_index13] = !(p2);
          let _88_v32;
          _88_v32 = _dafny.MultiSet.fromElements(p2, p2);
          let _89_v33;
          _89_v33 = _dafny.Seq.of(_88_v32);
          let _index14 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_38_v7).length));
          (_38_v7)[_index14] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_89_v33, _89_v33), _89_v33);
          r0 = p0;
          let _90_v34;
          let _nw8 = Array((new BigNumber(23)).toNumber()).fill([]);
          _90_v34 = _nw8;
          let _91_v35;
          let _nw9 = Array((new BigNumber(21)).toNumber()).fill(false);
          _91_v35 = _nw9;
          let _index15 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_90_v34).length));
          (_90_v34)[_index15] = _91_v35;
          let _92_v36;
          _92_v36 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index16 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((_90_v34).length));
          (_90_v34)[_index16] = (_39_v8)[_module.__default.safeIndex(_module.__default.fm3(_92_v36, true, p0, globalState), new BigNumber((_39_v8).length))];
        }
        let _93_v37;
        _93_v37 = _module.D4.create_DC15((p0).plus(_37_i2), p2);
        let _rhs9 = p2;
        let _rhs10 = _93_v37;
        let _lhs6 = globalState;
        _lhs6.f3 = _rhs9;
        _93_v37 = _rhs10;
        let _94_v38;
        _94_v38 = _dafny.Seq.of(p2, p2);
        let _95_v39;
        let _nw10 = new _module.C2();
        _nw10.__ctor((_94_v38)[_module.__default.safeIndex(new BigNumber((_94_v38).length), new BigNumber((_94_v38).length))], p2);
        _95_v39 = _nw10;
      }
      let _hi2 = p0;
      for (let _96_i4 = (p0).plus(p0); _96_i4.isLessThan(_hi2); _96_i4 = _96_i4.plus(_dafny.ONE)) {
        (globalState).f0 = !(p2);
        r0 = (_96_i4).minus(new BigNumber((_module.__default.fm11(false, p2, globalState)).length));
        (globalState).f0 = (((!(!(p2))) ? (p2) : (p2))) && (!(!(_96_i4).isEqualTo(p0)));
        (globalState).f3 = p2;
      }
      let _hi3 = p0;
      for (let _97_i5 = p0; _97_i5.isLessThan(_hi3); _97_i5 = _97_i5.plus(_dafny.ONE)) {
        let _98_v40;
        _98_v40 = _dafny.Seq.of(p2, p2, p2);
        let _99_v41;
        let _nw11 = new _module.C1();
        _nw11.__ctor(!(p2) || (_module.__default.fm2(globalState)), _module.__default.safeModuloInt(new BigNumber((_98_v40).length), _97_i5), ((p2) ? (p2) : (p2)));
        _99_v41 = _nw11;
        _99_v41 = _99_v41;
        let _100_v42;
        let _init1 = ((_101_v41) => function (_102_i6) {
          return (_101_v41).f11;
        })(_99_v41);
        let _nw12 = Array((new BigNumber(15)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
          _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _100_v42 = _nw12;
        let _103_v43;
        _103_v43 = _dafny.Map.Empty.slice().updateUnsafe(_100_v42,(_99_v41).f11);
        _103_v43 = (_103_v43).update(_100_v42, p2);
        let _104_v44;
        _104_v44 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(p0));
        (globalState).f3 = ((_104_v44).IsSubsetOf(_104_v44)) === ((_99_v41).f11);
        let _105_v45;
        let _nw13 = new _module.C3();
        _nw13.__ctor((_99_v41).f11);
        _105_v45 = _nw13;
        let _106_v46;
        _106_v46 = _dafny.Map.Empty.slice().updateUnsafe(_100_v42,_dafny.Map.Empty.slice().updateUnsafe(_105_v45,(_99_v41).f11));
        let _107_v47;
        _107_v47 = _dafny.Map.Empty.slice().updateUnsafe(_105_v45,(_99_v41).f11);
        _106_v46 = (_106_v46).update(_100_v42, _107_v47);
      }
      let _108_i7;
      _108_i7 = _dafny.ZERO;
      L0: {
        while (!(!(p2))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_108_i7)) {
              break L0;
            }
            _108_i7 = (_108_i7).plus(_dafny.ONE);
            (globalState).f0 = p2;
            let _109_v48;
            _109_v48 = _dafny.MultiSet.fromElements(p0);
            let _110_v49;
            _110_v49 = _dafny.MultiSet.fromElements(_109_v48);
            let _111_v50;
            _111_v50 = _module.D0.create_DC2(p2, p0, _110_v49);
            let _112_v51;
            let _nw14 = Array((new BigNumber(21)).toNumber());
            _nw14[(_dafny.ZERO).toNumber()] = p2;
            _nw14[(_dafny.ONE).toNumber()] = p2;
            _nw14[(new BigNumber(2)).toNumber()] = p2;
            _nw14[(new BigNumber(3)).toNumber()] = p2;
            _nw14[(new BigNumber(4)).toNumber()] = p2;
            _nw14[(new BigNumber(5)).toNumber()] = p2;
            _nw14[(new BigNumber(6)).toNumber()] = p2;
            _nw14[(new BigNumber(7)).toNumber()] = p2;
            _nw14[(new BigNumber(8)).toNumber()] = p2;
            _nw14[(new BigNumber(9)).toNumber()] = p2;
            _nw14[(new BigNumber(10)).toNumber()] = p2;
            _nw14[(new BigNumber(11)).toNumber()] = p2;
            _nw14[(new BigNumber(12)).toNumber()] = p2;
            _nw14[(new BigNumber(13)).toNumber()] = p2;
            _nw14[(new BigNumber(14)).toNumber()] = !(p2);
            _nw14[(new BigNumber(15)).toNumber()] = (_111_v50).dtor_cf2;
            _nw14[(new BigNumber(16)).toNumber()] = p2;
            _nw14[(new BigNumber(17)).toNumber()] = true;
            _nw14[(new BigNumber(18)).toNumber()] = p2;
            _nw14[(new BigNumber(19)).toNumber()] = p2;
            _nw14[(new BigNumber(20)).toNumber()] = p2;
            _112_v51 = _nw14;
            let _113_v52;
            _113_v52 = _dafny.Seq.of(_112_v51, _112_v51, _112_v51);
            let _114_v53;
            _114_v53 = _dafny.Seq.UnicodeFromString("nw");
            let _115_v54;
            _115_v54 = _dafny.Seq.of(p0, new BigNumber((_dafny.Seq.Concat(_113_v52, _dafny.Seq.update(_113_v52, _module.__default.safeIndex(new BigNumber((_114_v53).length), new BigNumber((_113_v52).length)), _112_v51))).length));
            let _116_v55;
            _116_v55 = _dafny.Seq.of(_115_v54);
            _115_v54 = (_116_v55)[_module.__default.safeIndex(p0, new BigNumber((_116_v55).length))];
            let _117_v56;
            _117_v56 = _dafny.Seq.of(p2, p2);
            if (_dafny.areEqual(_module.__default.fm10(globalState), _117_v56)) {
              (globalState).f0 = p2;
              let _118_v57;
              _118_v57 = _dafny.Map.Empty.slice().updateUnsafe((p0).isEqualTo(p0),p2);
              let _119_v58;
              _119_v58 = new _dafny.CodePoint('k'.codePointAt(0));
              let _120_v59;
              _120_v59 = _dafny.Map.Empty.slice().updateUnsafe(p2,_119_v58);
              let _rhs11 = _112_v51;
              let _rhs12 = _module.__default.fm13(p2, ((p2) ? (p2) : (p2)), p2, p0, globalState);
              let _rhs13 = (_module.__default.fm21(p2, globalState)).dtor_cf30;
              let _rhs14 = p2;
              let _rhs15 = _120_v59;
              let _lhs7 = globalState;
              _112_v51 = _rhs11;
              _115_v54 = _rhs12;
              _118_v57 = _rhs13;
              _lhs7.f3 = _rhs14;
              _120_v59 = _rhs15;
              let _121_v60;
              let _nw15 = Array((new BigNumber(26)).toNumber()).fill([]);
              _121_v60 = _nw15;
              let _122_v61;
              let _nw16 = Array((new BigNumber(7)).toNumber());
              _122_v61 = _nw16;
              let _index17 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_121_v60).length));
              (_121_v60)[_index17] = _122_v61;
              let _index18 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_121_v60).length));
              (_121_v60)[_index18] = _122_v61;
              let _123_v62;
              _123_v62 = _dafny.MultiSet.fromElements(p2, p2, p2, p2, p2);
              let _124_v63;
              _124_v63 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(p0, new BigNumber((_123_v62).cardinality()))).length));
              let _125_v64;
              let _nw17 = new _module.C0();
              _nw17.__ctor(true, (_115_v54)[_module.__default.safeIndex(p0, new BigNumber((_115_v54).length))], (_124_v63).IsProperSubsetOf(_124_v63));
              _125_v64 = _nw17;
              _125_v64 = _125_v64;
              let _126_v65;
              _126_v65 = _dafny.Set.fromElements(_109_v48, _109_v48, (_dafny.MultiSet.FromArray(_115_v54)).Union(_109_v48));
              let _127_v66;
              let _nw18 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
              _127_v66 = _nw18;
              let _128_v67;
              _128_v67 = _dafny.Map.Empty.slice().updateUnsafe((_125_v64).f13,_117_v56);
              let _index19 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_127_v66).length));
              (_127_v66)[_index19] = _128_v67;
              let _index20 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_127_v66).length));
              let _rhs16 = _module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p0));
              let _rhs17 = _126_v65;
              let _rhs18 = _dafny.Seq.Concat(_114_v53, _114_v53);
              let _rhs19 = _dafny.MultiSet.fromElements(_125_v64.f14, _module.__default.safeModuloInt(_125_v64.f14, _125_v64.f14));
              let _rhs20 = ((_128_v67).Merge(_128_v67)).Merge(_128_v67);
              let _lhs8 = _125_v64;
              let _lhs9 = _127_v66;
              let _lhs10 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_127_v66).length));
              _lhs8.f14 = _rhs16;
              _126_v65 = _rhs17;
              _114_v53 = _rhs18;
              _109_v48 = _rhs19;
              _lhs9[_lhs10] = _rhs20;
            } else {
              let _129_v68;
              _129_v68 = _dafny.MultiSet.fromElements(_115_v54, _dafny.Seq.of(p0));
              (globalState).f0 = !((new BigNumber((_129_v68).cardinality())).isLessThanOrEqualTo(p0));
              _115_v54 = _115_v54;
              r0 = (p0).multipliedBy(p0);
              let _130_v69;
              _130_v69 = new _dafny.CodePoint('l'.codePointAt(0));
              let _index21 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((p1).length));
              (p1)[_index21] = (new BigNumber((_114_v53).length)).plus(_module.__default.fm3(_130_v69, true, new BigNumber((_114_v53).length), globalState));
              let _index22 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((p1).length));
              (p1)[_index22] = p0;
              let _131_v70;
              let _nw19 = new _module.C0();
              _nw19.__ctor(false, p0, p2);
              _131_v70 = _nw19;
              let _132_v71;
              let _nw20 = new _module.C3();
              _nw20.__ctor((_131_v70).f13);
              _132_v71 = _nw20;
              let _133_v72;
              _133_v72 = _dafny.Map.Empty.slice().updateUnsafe(_131_v70,_132_v71);
              let _134_v73;
              _134_v73 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_133_v72).Merge(_dafny.Map.Empty.slice().updateUnsafe(_131_v70,_132_v71)));
              _134_v73 = (_134_v73).update(!((_131_v70).f13), _dafny.Map.Empty.slice().updateUnsafe(_131_v70,_132_v71));
            }
            let _index23 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_112_v51).length));
            (_112_v51)[_index23] = p2;
            let _index24 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_112_v51).length));
            (_112_v51)[_index24] = ((p0).isLessThanOrEqualTo(p0)) || (p2);
          }
        }
      }
      let _hi4 = p0;
      for (let _135_i8 = new BigNumber(702); _135_i8.isLessThan(_hi4); _135_i8 = _135_i8.plus(_dafny.ONE)) {
        (globalState).f0 = (_135_i8).isLessThanOrEqualTo(_135_i8);
        let _136_v74;
        let _nw21 = Array((new BigNumber(4)).toNumber()).fill(false);
        _136_v74 = _nw21;
        let _index25 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_136_v74).length));
        (_136_v74)[_index25] = p2;
        let _index26 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_136_v74).length));
        (_136_v74)[_index26] = false;
        let _137_v75;
        let _nw22 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
        _137_v75 = _nw22;
        let _138_v76;
        _138_v76 = _dafny.Seq.of(p2);
        let _139_v77;
        _139_v77 = _dafny.Set.fromElements(p2, p2, p2);
        let _index27 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_137_v75).length));
        (_137_v75)[_index27] = (_dafny.Set.fromElements(p2, p2, (_138_v76)[_module.__default.safeIndex(p0, new BigNumber((_138_v76).length))], p2)).Union(_139_v77);
        let _140_v78;
        let _nw23 = new _module.C1();
        _nw23.__ctor(p2, _module.__default.safeModuloInt(_135_i8, p0), p2);
        _140_v78 = _nw23;
        let _141_v79;
        _141_v79 = _dafny.Set.fromElements(_135_i8, p0, p0);
        let _142_v80;
        _142_v80 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("hdndrwawf")).length));
        let _143_v81;
        _143_v81 = _dafny.Map.Empty.slice().updateUnsafe(((((_142_v80).contains(_135_i8)) ? ((_142_v80).get(_135_i8)) : (_135_i8))).minus(_135_i8),(p0).multipliedBy(new BigNumber(-532)));
        let _144_v82;
        _144_v82 = _dafny.Seq.UnicodeFromString("koccov");
        let _145_v83;
        _145_v83 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of(true, true, p2, (_140_v78).f9));
        let _146_v84;
        let _nw24 = new _module.C2();
        _nw24.__ctor((_140_v78).f9, (_140_v78).f9);
        _146_v84 = _nw24;
        let _147_v85;
        _147_v85 = _dafny.MultiSet.fromElements(_146_v84);
        let _index28 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_136_v74).length));
        let _index29 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_136_v74).length));
        let _index30 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_137_v75).length));
        let _rhs21 = (_141_v79).IsProperSubsetOf(_141_v79);
        let _rhs22 = p2;
        let _rhs23 = (((_143_v81).contains(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-199)), new BigNumber((_144_v82).length)))) ? ((_143_v81).get(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-199)), new BigNumber((_144_v82).length)))) : (((_140_v78).fm4(_145_v83, p2, (_140_v78).f9, p0, globalState)).minus(new BigNumber((_147_v85).cardinality()))));
        let _rhs24 = _139_v77;
        let _rhs25 = _140_v78;
        let _lhs11 = _136_v74;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_136_v74).length));
        let _lhs13 = _136_v74;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_136_v74).length));
        let _lhs15 = _137_v75;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_137_v75).length));
        _lhs11[_lhs12] = _rhs21;
        _lhs13[_lhs14] = _rhs22;
        r0 = _rhs23;
        _lhs15[_lhs16] = _rhs24;
        _140_v78 = _rhs25;
        let _148_v86;
        let _nw25 = Array((new BigNumber(8)).toNumber()).fill([]);
        _148_v86 = _nw25;
        let _149_v87;
        _149_v87 = _module.D4.create_DC13(_148_v86);
        let _150_v88;
        _150_v88 = _module.D4.create_DC16(_149_v87);
        _150_v88 = _150_v88;
        let _index31 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_136_v74).length));
        (_136_v74)[_index31] = (_135_i8).isEqualTo(new BigNumber(-200));
      }
      r0 = _module.__default.safeDivisionInt((p0).multipliedBy(p0), p0);
      return r0;
    }
    static Main(__noArgsParameter) {
      let _151_v0;
      _151_v0 = new BigNumber(636);
      let _152_v1;
      _152_v1 = _dafny.Seq.of(_151_v0, _151_v0);
      let _153_v2;
      _153_v2 = true;
      let _154_v3;
      _154_v3 = _dafny.Map.Empty.slice().updateUnsafe(_153_v2,_152_v1);
      let _155_v4;
      _155_v4 = _dafny.MultiSet.fromElements(_152_v1, _152_v1, (((_154_v3).contains(!(false))) ? ((_154_v3).get(!(false))) : (_dafny.Seq.of(_151_v0))), _152_v1, _152_v1);
      let _156_v5;
      let _nw26 = Array((new BigNumber(26)).toNumber()).fill(false);
      _156_v5 = _nw26;
      let _157_v6;
      _157_v6 = _dafny.MultiSet.fromElements(_156_v5, _156_v5);
      let _158_v7;
      _158_v7 = _module.D0.create_DC0((_dafny.ZERO).minus(_151_v0));
      let _159_v8;
      _159_v8 = _dafny.Seq.UnicodeFromString("iyuhsbbd");
      let _160_v9;
      _160_v9 = _dafny.Map.Empty.slice().updateUnsafe(_153_v2,_151_v0);
      let _161_v10;
      let _nw27 = Array((new BigNumber(17)).toNumber());
      _nw27[(_dafny.ZERO).toNumber()] = _151_v0;
      _nw27[(_dafny.ONE).toNumber()] = _151_v0;
      _nw27[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_153_v2,new BigNumber((_dafny.Seq.UnicodeFromString("yqsopcbd")).length))).length);
      _nw27[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_151_v0);
      _nw27[(new BigNumber(4)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(5)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(6)).toNumber()] = new BigNumber((_157_v6).cardinality());
      _nw27[(new BigNumber(7)).toNumber()] = (_158_v7).dtor_cf0;
      _nw27[(new BigNumber(8)).toNumber()] = new BigNumber((_152_v1).length);
      _nw27[(new BigNumber(9)).toNumber()] = new BigNumber((_159_v8).length);
      _nw27[(new BigNumber(10)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(11)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(12)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(13)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(14)).toNumber()] = _151_v0;
      _nw27[(new BigNumber(15)).toNumber()] = new BigNumber((_160_v9).length);
      _nw27[(new BigNumber(16)).toNumber()] = _151_v0;
      _161_v10 = _nw27;
      let _162_globalState;
      let _nw28 = new _module.GlobalState();
      _nw28.__ctor(true, new BigNumber(463), true, false, false, _155_v4, _161_v10, new _dafny.CodePoint('s'.codePointAt(0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(798)), function (_163_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }));
      _162_globalState = _nw28;
      let _164_v11;
      _164_v11 = new _dafny.CodePoint('p'.codePointAt(0));
      let _pat_let_tv6 = _159_v8;
      let _pat_let_tv7 = _159_v8;
      let _pat_let_tv8 = _159_v8;
      let _pat_let_tv9 = _159_v8;
      let _pat_let_tv10 = _159_v8;
      let _pat_let_tv11 = _159_v8;
      let _pat_let_tv12 = _159_v8;
      let _pat_let_tv13 = _159_v8;
      let _pat_let_tv14 = _159_v8;
      let _pat_let_tv15 = _159_v8;
      let _pat_let_tv16 = _159_v8;
      let _pat_let_tv17 = _159_v8;
      _159_v8 = _dafny.Seq.update(function (_source1) {
        if (_source1.is_DC1) {
          let _165___mcc_h0 = (_source1).cf1;
          let _166_cf1 = _165___mcc_h0;
          return _pat_let_tv6;
        } else if (_source1.is_DC2) {
          let _167___mcc_h1 = (_source1).cf2;
          let _168___mcc_h2 = (_source1).cf3;
          let _169___mcc_h3 = (_source1).cf4;
          let _170_cf4 = _169___mcc_h3;
          let _171_cf3 = _168___mcc_h2;
          let _172_cf2 = _167___mcc_h1;
          return _dafny.Seq.Concat(_pat_let_tv7, _pat_let_tv8);
        } else if (_source1.is_DC3) {
          return _dafny.Seq.Concat(_pat_let_tv9, _pat_let_tv10);
        } else {
          let _173___mcc_h4 = (_source1).cf0;
          let _174_cf0 = _173___mcc_h4;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qbtq"), _pat_let_tv11);
        }
      }(_module.__default.fm0(_153_v2, _162_globalState)), _module.__default.safeIndex(_151_v0, new BigNumber((function (_source2) {
        if (_source2.is_DC1) {
          let _175___mcc_h5 = (_source2).cf1;
          let _176_cf1 = _175___mcc_h5;
          return _pat_let_tv12;
        } else if (_source2.is_DC2) {
          let _177___mcc_h6 = (_source2).cf2;
          let _178___mcc_h7 = (_source2).cf3;
          let _179___mcc_h8 = (_source2).cf4;
          let _180_cf4 = _179___mcc_h8;
          let _181_cf3 = _178___mcc_h7;
          let _182_cf2 = _177___mcc_h6;
          return _dafny.Seq.Concat(_pat_let_tv13, _pat_let_tv14);
        } else if (_source2.is_DC3) {
          return _dafny.Seq.Concat(_pat_let_tv15, _pat_let_tv16);
        } else {
          let _183___mcc_h9 = (_source2).cf0;
          let _184_cf0 = _183___mcc_h9;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qbtq"), _pat_let_tv17);
        }
      }(_module.__default.fm0(_153_v2, _162_globalState))).length)), _164_v11);
      let _hi5 = _151_v0;
      for (let _185_i1 = _module.__default.safeModuloInt(_151_v0, _151_v0); _185_i1.isLessThan(_hi5); _185_i1 = _185_i1.plus(_dafny.ONE)) {
        let _source3 = _module.__default.fm1(_159_v8, true, _162_globalState);
        if (_source3.is_DC1) {
          let _186___mcc_h10 = (_source3).cf1;
          let _187_cf1 = _186___mcc_h10;
          let _index32 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_156_v5).length));
          (_156_v5)[_index32] = _187_cf1;
          let _index33 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_156_v5).length));
          (_156_v5)[_index33] = _module.__default.fm2(_162_globalState);
          let _188_v12;
          _188_v12 = _dafny.Map.Empty.slice().updateUnsafe(_151_v0,_159_v8);
          let _189_v13;
          _189_v13 = _module.D0.create_DC1(!(_188_v12).equals(_188_v12));
          _189_v13 = _189_v13;
          let _190_v14;
          _190_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(413),_151_v0);
          let _191_v15;
          _191_v15 = _dafny.Map.Empty.slice().updateUnsafe(_185_i1,_190_v14);
          let _192_v16;
          _192_v16 = _dafny.Set.fromElements(new BigNumber((_159_v8).length));
          _151_v0 = _module.__default.safeDivisionInt((_152_v1)[_module.__default.safeIndex(new BigNumber((((((_191_v15).contains(_151_v0)) ? ((_191_v15).get(_151_v0)) : (_190_v14))).update(new BigNumber((_192_v16).length), _185_i1)).length), new BigNumber((_152_v1).length))], _151_v0);
          let _193_v17;
          _193_v17 = _dafny.MultiSet.fromElements((_156_v5)[_module.__default.safeIndex(new BigNumber(185), new BigNumber((_156_v5).length))], (_156_v5)[_module.__default.safeIndex(new BigNumber(185), new BigNumber((_156_v5).length))]);
          let _194_v18;
          _194_v18 = _dafny.Seq.of((_193_v17).IsSubsetOf(_193_v17));
          _151_v0 = (_dafny.ZERO).minus(new BigNumber((_194_v18).length));
        } else if (_source3.is_DC2) {
          let _195___mcc_h11 = (_source3).cf2;
          let _196___mcc_h12 = (_source3).cf3;
          let _197___mcc_h13 = (_source3).cf4;
          let _198_cf4 = _197___mcc_h13;
          let _199_cf3 = _196___mcc_h12;
          let _200_cf2 = _195___mcc_h11;
          let _201_v19;
          _201_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_200_cf2, _200_cf2)).length),_200_cf2);
          let _202_v20;
          let _out3;
          _out3 = _module.__default.m0(_199_cf3, _161_v10, _153_v2, (_201_v19).Merge(_201_v19), _162_globalState);
          _202_v20 = _out3;
          let _nw29 = Array((new BigNumber(20)).toNumber()).fill(false);
          _156_v5 = _nw29;
          (_162_globalState).f7 = _164_v11;
          (_162_globalState).f0 = (new BigNumber(-274)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_153_v2,(_152_v1)[_module.__default.safeIndex(_202_v20, new BigNumber((_152_v1).length))])).length));
        } else if (_source3.is_DC3) {
          let _203_v21;
          _203_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(363),_153_v2);
          let _204_v22;
          _204_v22 = _dafny.Seq.of(_153_v2);
          let _205_v23;
          _205_v23 = _dafny.Set.fromElements(_151_v0, _185_i1, new BigNumber((_204_v22).length));
          let _206_v24;
          let _out4;
          _out4 = _module.__default.m0((_185_i1).multipliedBy(_151_v0), _161_v10, _153_v2, (_203_v21).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm3(_164_v11, _153_v2, new BigNumber((_205_v23).length), _162_globalState),_153_v2)), _162_globalState);
          _206_v24 = _out4;
          _153_v2 = (_204_v22)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_206_v24)), new BigNumber((_204_v22).length))];
          (_162_globalState).f0 = _153_v2;
          (_162_globalState).f3 = _module.__default.fm2(_162_globalState);
        } else {
          let _207___mcc_h14 = (_source3).cf0;
          let _208_cf0 = _207___mcc_h14;
          let _index34 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_161_v10).length));
          (_161_v10)[_index34] = _151_v0;
          let _index35 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_161_v10).length));
          (_161_v10)[_index35] = new BigNumber((_159_v8).length);
          let _209_v25;
          let _nw30 = new _module.C3();
          _nw30.__ctor(_153_v2);
          _209_v25 = _nw30;
          let _210_v26;
          _210_v26 = _dafny.MultiSet.fromElements(_151_v0);
          let _211_v27;
          _211_v27 = _dafny.Seq.of(_153_v2);
          let _212_v28;
          _212_v28 = _dafny.Set.fromElements((_211_v27)[_module.__default.safeIndex(_151_v0, new BigNumber((_211_v27).length))], _153_v2, _153_v2, _153_v2);
          let _213_v29;
          let _nw31 = Array((new BigNumber(23)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = (_module.__default.fm17(_152_v1, _153_v2, _162_globalState)).Difference(_210_v26);
          _nw31[(_dafny.ONE).toNumber()] = _210_v26;
          _nw31[(new BigNumber(2)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(3)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(4)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(5)).toNumber()] = (_210_v26).Intersect(_210_v26);
          _nw31[(new BigNumber(6)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(7)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(8)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(9)).toNumber()] = (_210_v26).Difference(_210_v26);
          _nw31[(new BigNumber(10)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_151_v0, new BigNumber((_dafny.Seq.UnicodeFromString("xpfylyv")).length));
          _nw31[(new BigNumber(12)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(13)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(14)).toNumber()] = (_210_v26).Union(_210_v26);
          _nw31[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm13(_153_v2, _153_v2, _153_v2, new BigNumber((_212_v28).length), _162_globalState));
          _nw31[(new BigNumber(16)).toNumber()] = (_210_v26).Intersect(_dafny.MultiSet.fromElements(_208_cf0));
          _nw31[(new BigNumber(17)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(18)).toNumber()] = _module.__default.fm17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(424)), ((_214_cf0) => function (_215_i2) {
            return _214_cf0;
          })(_208_cf0)), _153_v2, _162_globalState);
          _nw31[(new BigNumber(19)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(20)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(21)).toNumber()] = _210_v26;
          _nw31[(new BigNumber(22)).toNumber()] = _210_v26;
          _213_v29 = _nw31;
          let _index36 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_213_v29).length));
          (_213_v29)[_index36] = (_210_v26).Intersect(_210_v26);
          let _index37 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_161_v10).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_213_v29).length));
          let _rhs26 = _153_v2;
          let _rhs27 = _185_i1;
          let _rhs28 = _210_v26;
          let _rhs29 = _module.__default.fm2(_162_globalState);
          let _lhs17 = _162_globalState;
          let _lhs18 = _161_v10;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_161_v10).length));
          let _lhs20 = _213_v29;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_213_v29).length));
          let _lhs22 = _162_globalState;
          _lhs17.f3 = _rhs26;
          _lhs18[_lhs19] = _rhs27;
          _lhs20[_lhs21] = _rhs28;
          _lhs22.f0 = _rhs29;
          let _index39 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_161_v10).length));
          (_161_v10)[_index39] = _208_cf0;
        }
        let _216_v30;
        _216_v30 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_153_v2)),_156_v5);
        _216_v30 = (_216_v30).update(_153_v2, _156_v5);
        let _217_v31;
        _217_v31 = _dafny.Set.fromElements(_153_v2, _153_v2, false);
        _217_v31 = _dafny.Set.fromElements(_153_v2, _153_v2);
        let _218_v32;
        let _nw32 = new _module.C3();
        _nw32.__ctor((_151_v0).isLessThanOrEqualTo(_185_i1));
        _218_v32 = _nw32;
      }
      (_162_globalState).f0 = _153_v2;
      let _219_v33;
      _219_v33 = _dafny.Set.fromElements(true, false, _153_v2);
      let _hi6 = _151_v0;
      for (let _220_i3 = new BigNumber((_219_v33).length); _220_i3.isLessThan(_hi6); _220_i3 = _220_i3.plus(_dafny.ONE)) {
        let _index40 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length));
        (_156_v5)[_index40] = _module.__default.fm2(_162_globalState);
        let _index41 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length));
        (_156_v5)[_index41] = _153_v2;
        let _index42 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length));
        (_161_v10)[_index42] = (_dafny.ZERO).minus(_151_v0);
        let _221_v34;
        _221_v34 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(_220_i3)), _151_v0, _220_i3, (_dafny.ZERO).minus(_220_i3));
        let _index43 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length));
        (_161_v10)[_index43] = _module.__default.fm3(_164_v11, _153_v2, new BigNumber((_221_v34).length), _162_globalState);
        let _index44 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length));
        (_161_v10)[_index44] = _220_i3;
        if ((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))]) {
          (_162_globalState).f3 = ((_153_v2) ? ((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))]) : (_153_v2));
          _151_v0 = _151_v0;
          let _222_v35;
          let _nw33 = new _module.C0();
          _nw33.__ctor((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))], (_161_v10)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length))], _153_v2);
          _222_v35 = _nw33;
          let _223_v36;
          _223_v36 = _module.D3.create_DC11(_151_v0, _dafny.Seq.UnicodeFromString("f"), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_153_v2,_220_i3)).length), _222_v35);
          let _224_v37;
          let _nw34 = new _module.C1();
          _nw34.__ctor(false, (_223_v36).dtor_cf14, (_222_v35).f13);
          _224_v37 = _nw34;
          let _225_v38;
          _225_v38 = _dafny.Map.Empty.slice().updateUnsafe(_224_v37,new BigNumber((_159_v8).length));
          let _index45 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length));
          (_161_v10)[_index45] = ((_153_v2) ? ((((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))]) ? ((_161_v10)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length))]) : (new BigNumber((_225_v38).length)))) : (_222_v35.f14));
          (_222_v35).f14 = _222_v35.f14;
          let _226_v39;
          let _nw35 = new _module.C1();
          _nw35.__ctor(_153_v2, (_161_v10)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length))], !(!(_153_v2)));
          _226_v39 = _nw35;
          let _227_v40;
          _227_v40 = _module.D4.create_DC15(_222_v35.f14, _153_v2);
          let _nw36 = new _module.C2();
          _nw36.__ctor((_227_v40).dtor_cf24, (_226_v39).f9);
          _226_v39 = _nw36;
        } else {
          _151_v0 = _151_v0;
          let _228_v41;
          let _init2 = ((_229_v0) => function (_230_i4) {
            return _module.__default.safeModuloInt(_230_i4, _229_v0);
          })(_151_v0);
          let _nw37 = Array((new BigNumber(13)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw37.length); _i0_2++) {
            _nw37[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _228_v41 = _nw37;
          let _231_v42;
          _231_v42 = _module.D0.create_DC1((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))]);
          let _232_v43;
          _232_v43 = _dafny.Map.Empty.slice().updateUnsafe(_151_v0,(_231_v42).dtor_cf1);
          let _233_v44;
          let _out5;
          _out5 = _module.__default.m0((_161_v10)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_161_v10).length))], _228_v41, !((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))]), _232_v43, _162_globalState);
          _233_v44 = _out5;
          let _234_v45;
          _234_v45 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_162_globalState),_158_v7);
          let _235_v46;
          _235_v46 = _dafny.Seq.of((_156_v5)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_156_v5).length))], _153_v2);
          _234_v45 = ((_234_v45).update(_235_v46, _module.D0.create_DC0(_220_i3))).update(_module.__default.fm10(_162_globalState), _158_v7);
          _151_v0 = new BigNumber(-74);
          (_162_globalState).f7 = _164_v11;
        }
      }
      let _236_v47;
      _236_v47 = _dafny.Seq.of(_153_v2);
      (_162_globalState).f3 = _dafny.Seq.IsPrefixOf(_236_v47, _236_v47);
      let _237_v48;
      let _nw38 = new _module.C0();
      _nw38.__ctor(_153_v2, _151_v0, _153_v2);
      _237_v48 = _nw38;
      if (((_153_v2) ? (_153_v2) : ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_156_v5,_dafny.Set.fromElements(_237_v48, _237_v48))).length)).isLessThan(_151_v0)))) {
        let _index46 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length));
        (_161_v10)[_index46] = _237_v48.f14;
        let _index47 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length));
        (_161_v10)[_index47] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_237_v48.f14), (_237_v48.f14).plus(new BigNumber(-739)));
        let _index48 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_156_v5).length));
        (_156_v5)[_index48] = (_237_v48).f13;
        let _index49 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_156_v5).length));
        (_156_v5)[_index49] = true;
        let _238_v49;
        let _nw39 = new _module.C3();
        _nw39.__ctor(((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))]).isLessThan((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))]));
        _238_v49 = _nw39;
        _238_v49 = _238_v49;
        if (!(new BigNumber(26)).isEqualTo(_module.__default.safeModuloInt(_237_v48.f14, _237_v48.f14))) {
          let _239_v50;
          let _nw40 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _239_v50 = _nw40;
          let _240_v51;
          _240_v51 = _dafny.Seq.of(_239_v50, _239_v50, _239_v50);
          let _241_v52;
          _241_v52 = _dafny.Seq.of((_240_v51)[_module.__default.safeIndex(new BigNumber((_159_v8).length), new BigNumber((_240_v51).length))], _239_v50);
          let _242_v53;
          let _nw41 = Array((new BigNumber(8)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _239_v50;
          _nw41[(_dafny.ONE).toNumber()] = _239_v50;
          _nw41[(new BigNumber(2)).toNumber()] = ((_153_v2) ? (_239_v50) : (_239_v50));
          _nw41[(new BigNumber(3)).toNumber()] = (_241_v52)[_module.__default.safeIndex((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))], new BigNumber((_241_v52).length))];
          _nw41[(new BigNumber(4)).toNumber()] = _239_v50;
          _nw41[(new BigNumber(5)).toNumber()] = _239_v50;
          _nw41[(new BigNumber(6)).toNumber()] = _239_v50;
          _nw41[(new BigNumber(7)).toNumber()] = _239_v50;
          _242_v53 = _nw41;
          let _index50 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_242_v53).length));
          (_242_v53)[_index50] = _239_v50;
          let _index51 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_242_v53).length));
          let _nw42 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          (_242_v53)[_index51] = _nw42;
          let _index52 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length));
          (_161_v10)[_index52] = _237_v48.f14;
          let _243_v54;
          _243_v54 = _dafny.Map.Empty.slice().updateUnsafe((_156_v5)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_156_v5).length))],_239_v50);
          _243_v54 = ((!(_module.__default.fm2(_162_globalState)) || (_153_v2)) ? (_243_v54) : ((_243_v54).update(true, _239_v50)));
          let _244_v55;
          _244_v55 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))], _151_v0));
          let _245_v56;
          _245_v56 = _dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,new BigNumber((_160_v9).length));
          let _246_v57;
          _246_v57 = _dafny.Seq.of(_236_v47);
          let _index53 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length));
          (_161_v10)[_index53] = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_244_v55).contains((((_245_v56).contains((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))])) ? ((_245_v56).get((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))])) : (_237_v48.f14)))) ? ((_244_v55).get((((_245_v56).contains((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))])) ? ((_245_v56).get((_161_v10)[_module.__default.safeIndex(new BigNumber(648), new BigNumber((_161_v10).length))])) : (_237_v48.f14)))) : (new BigNumber((_dafny.Seq.Concat(_236_v47, (_246_v57)[_module.__default.safeIndex((_dafny.ZERO).minus(_151_v0), new BigNumber((_246_v57).length))])).length)))));
          let _247_v58;
          _247_v58 = _dafny.Map.Empty.slice().updateUnsafe((_237_v48).f13,(_237_v48).f13);
          let _248_v59;
          let _nw43 = new _module.C2();
          _nw43.__ctor(!(_247_v58).equals(_247_v58), (_237_v48).f13);
          _248_v59 = _nw43;
          let _249_v60;
          _249_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_162_globalState),_248_v59);
          _248_v59 = ((_153_v2) ? (_248_v59) : ((((_249_v60).contains((_237_v48).f13)) ? ((_249_v60).get((_237_v48).f13)) : (_248_v59))));
        } else {
          let _250_v61;
          _250_v61 = _dafny.Map.Empty.slice().updateUnsafe(false,(_237_v48).f13);
          let _251_v62;
          _251_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_250_v61).length),(_237_v48).f13);
          let _252_v63;
          _252_v63 = _dafny.Map.Empty.slice().updateUnsafe(_160_v9,!(_251_v62).equals(_251_v62));
          _252_v63 = _252_v63;
          _237_v48 = _237_v48;
          let _253_v64;
          let _nw44 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _253_v64 = _nw44;
          let _254_v65;
          let _nw45 = new _module.C1();
          _nw45.__ctor(true, (_dafny.ZERO).minus(_237_v48.f14), (_156_v5)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_156_v5).length))]);
          _254_v65 = _nw45;
          let _255_v66;
          _255_v66 = _dafny.Map.Empty.slice().updateUnsafe(_253_v64,_254_v65);
          _255_v66 = (_255_v66).update(_253_v64, _254_v65);
          (_237_v48).f14 = _module.__default.safeModuloInt(new BigNumber(778), _237_v48.f14);
          let _256_v67;
          _256_v67 = _dafny.MultiSet.fromElements((_254_v65).f12);
          let _257_v68;
          let _nw46 = new _module.C0();
          _nw46.__ctor((_256_v67).contains(new BigNumber(550)), new BigNumber(114), (_238_v49).f9);
          _257_v68 = _nw46;
        }
        let _258_v69;
        let _nw47 = new _module.C2();
        _nw47.__ctor(_dafny.Seq.contains(_159_v8, _164_v11), (_156_v5)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_156_v5).length))]);
        _258_v69 = _nw47;
        _258_v69 = _258_v69;
      } else {
        let _index54 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_156_v5).length));
        (_156_v5)[_index54] = (_237_v48).f13;
        let _index55 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_156_v5).length));
        (_156_v5)[_index55] = (_237_v48).f13;
        (_162_globalState).f3 = (_237_v48).f13;
        let _259_v70;
        _259_v70 = _module.D0.create_DC3();
        let _source4 = (((_237_v48).f13) ? (_259_v70) : (_259_v70));
        if (_source4.is_DC1) {
          let _260___mcc_h15 = (_source4).cf1;
          let _261_cf1 = _260___mcc_h15;
          let _262_v71;
          let _nw48 = new _module.C1();
          _nw48.__ctor((_237_v48).f13, _237_v48.f14, !((_237_v48).f13));
          _262_v71 = _nw48;
          let _263_v72;
          _263_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(186),_262_v71);
          let _264_v73;
          _264_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_263_v72).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(366),_262_v71))).length),_236_v47);
          let _rhs30 = (_264_v73).Merge(_264_v73);
          let _rhs31 = (_dafny.ZERO).minus(_237_v48.f14);
          let _lhs23 = _237_v48;
          _264_v73 = _rhs30;
          _lhs23.f14 = _rhs31;
          _159_v8 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("p"), _module.__default.safeIndex((_262_v71).f12, new BigNumber((_dafny.Seq.UnicodeFromString("p")).length)), _module.__default.fm18(_162_globalState));
          let _265_v74;
          let _init3 = ((_266_v8) => function (_267_i5) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), ((_268_v8) => function (_269_i6) {
              return _module.D3.create_DC9(_268_v8);
            })(_266_v8));
          })(_159_v8);
          let _nw49 = Array((new BigNumber(25)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw49.length); _i0_3++) {
            _nw49[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _265_v74 = _nw49;
          let _270_v75;
          _270_v75 = _dafny.Map.Empty.slice().updateUnsafe((_262_v71).f12,new BigNumber(-277));
          let _271_v76;
          _271_v76 = _module.D3.create_DC9(_159_v8);
          let _272_v77;
          _272_v77 = _dafny.Seq.of(_271_v76, _271_v76, _271_v76);
          let _index56 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_265_v74).length));
          (_265_v74)[_index56] = _dafny.Seq.Concat(_module.__default.fm19(_270_v75, _162_globalState), _272_v77);
          let _index57 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_265_v74).length));
          (_265_v74)[_index57] = _272_v77;
          let _273_v79;
          let _out6;
          _out6 = _module.__default.m0(_237_v48.f14, _161_v10, !_dafny.areEqual(_159_v8, _dafny.Seq.UnicodeFromString("gretekrl")), function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(515), new BigNumber(497))) {
              let _274_v78 = _compr_9;
              if (((new BigNumber(515)).isLessThanOrEqualTo(_274_v78)) && ((_274_v78).isLessThan(new BigNumber(497)))) {
                _coll9.push([(_274_v78).minus(_237_v48.f14),false]);
              }
            }
            return _coll9;
          }(), _162_globalState);
          _273_v79 = _out6;
        } else if (_source4.is_DC2) {
          let _275___mcc_h16 = (_source4).cf2;
          let _276___mcc_h17 = (_source4).cf3;
          let _277___mcc_h18 = (_source4).cf4;
          let _278_cf4 = _277___mcc_h18;
          let _279_cf3 = _276___mcc_h17;
          let _280_cf2 = _275___mcc_h16;
          let _281_v80;
          let _nw50 = new _module.C1();
          _nw50.__ctor((_156_v5)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_156_v5).length))], (_279_cf3).multipliedBy(_279_cf3), (_237_v48).f13);
          _281_v80 = _nw50;
          _160_v9 = (_160_v9).update((_281_v80).f11, new BigNumber((_160_v9).length));
          let _282_v81;
          let _nw51 = new _module.C2();
          _nw51.__ctor((_237_v48).f13, _dafny.Seq.IsPrefixOf(_152_v1, _152_v1));
          _282_v81 = _nw51;
          let _nw52 = new _module.C2();
          _nw52.__ctor((_151_v0).isEqualTo(_237_v48.f14), (_279_cf3).isLessThanOrEqualTo((_281_v80).f12));
          _282_v81 = _nw52;
          let _rhs32 = _module.__default.fm18(_162_globalState);
          let _rhs33 = new BigNumber(712);
          let _lhs24 = _162_globalState;
          _lhs24.f7 = _rhs32;
          _279_cf3 = _rhs33;
        } else if (_source4.is_DC3) {
          (_237_v48).f14 = _237_v48.f14;
          let _283_v82;
          _283_v82 = _dafny.Map.Empty.slice().updateUnsafe(_153_v2,!_dafny.areEqual(_159_v8, _159_v8));
          let _284_v83;
          let _nw53 = Array((new BigNumber(29)).toNumber()).fill(false);
          _284_v83 = _nw53;
          let _285_v84;
          _285_v84 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_284_v83),_159_v8);
          let _286_v85;
          _286_v85 = _dafny.Map.Empty.slice().updateUnsafe((((_285_v84).contains(_module.D2.create_DC7(_284_v83))) ? ((_285_v84).get(_module.D2.create_DC7(_284_v83))) : (_159_v8)),_236_v47);
          let _rhs34 = _283_v82;
          let _rhs35 = (_286_v85).Merge(_286_v85);
          let _rhs36 = _151_v0;
          let _lhs25 = _237_v48;
          _283_v82 = _rhs34;
          _286_v85 = _rhs35;
          _lhs25.f14 = _rhs36;
          let _287_v86;
          _287_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_159_v8).length),_151_v0);
          let _288_v87;
          _288_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_287_v86).update(new BigNumber((_159_v8).length), _237_v48.f14)).length),_module.__default.fm2(_162_globalState));
          let _289_v88;
          let _out7;
          _out7 = _module.__default.m0((_dafny.ZERO).minus(new BigNumber((_159_v8).length)), _161_v10, _153_v2, _288_v87, _162_globalState);
          _289_v88 = _out7;
          let _290_v89;
          let _out8;
          _out8 = _module.__default.m0((_dafny.ZERO).minus((_151_v0).plus(new BigNumber(514))), _161_v10, (_237_v48).f13, _288_v87, _162_globalState);
          _290_v89 = _out8;
        } else {
          let _291___mcc_h19 = (_source4).cf0;
          let _292_cf0 = _291___mcc_h19;
          (_162_globalState).f0 = (_237_v48).f13;
          let _293_v90;
          let _init4 = ((_294_v11, _295_v8) => function (_296_i7) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-815)), ((_297_v11) => function (_298_i8) {
              return _297_v11;
            })(_294_v11)), _295_v8);
          })(_164_v11, _159_v8);
          let _nw54 = Array((new BigNumber(27)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw54.length); _i0_4++) {
            _nw54[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _293_v90 = _nw54;
          _293_v90 = _293_v90;
          let _299_v91;
          let _nw55 = new _module.C0();
          _nw55.__ctor(!((_237_v48).f13), _237_v48.f14, _153_v2);
          _299_v91 = _nw55;
          let _300_v92;
          let _nw56 = new _module.C1();
          _nw56.__ctor(true, _237_v48.f14, (_237_v48).f13);
          _300_v92 = _nw56;
          let _301_v93;
          _301_v93 = _dafny.Seq.of(_300_v92);
          _153_v2 = ((((_237_v48).f13) ? (new BigNumber(428)) : (new BigNumber((_301_v93).length)))).isLessThan(_151_v0);
        }
        let _302_v94;
        let _nw57 = Array((new BigNumber(15)).toNumber()).fill(_dafny.MultiSet.Empty);
        _302_v94 = _nw57;
        let _index58 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_302_v94).length));
        (_302_v94)[_index58] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_236_v47, _module.__default.safeIndex(_151_v0, new BigNumber((_236_v47).length)), (_237_v48).f13), _236_v47));
        let _303_v95;
        _303_v95 = _dafny.MultiSet.fromElements(true);
        let _index59 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_302_v94).length));
        (_302_v94)[_index59] = _303_v95;
        _151_v0 = _237_v48.f14;
      }
      let _304_v96;
      _304_v96 = _dafny.MultiSet.fromElements(_151_v0);
      _151_v0 = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_237_v48).f13,new BigNumber((_304_v96).cardinality()))).length)).plus(new BigNumber(771));
      if (((_153_v2) ? ((_module.__default.fm2(_162_globalState)) === ((_237_v48).f13)) : (_153_v2))) {
        (_237_v48).f14 = _237_v48.f14;
        let _305_v97;
        _305_v97 = _module.D0.create_DC3();
        let _306_v98;
        _306_v98 = _dafny.Set.fromElements(_305_v97, _305_v97);
        (_162_globalState).f3 = (new BigNumber((_304_v96).cardinality())).isEqualTo(new BigNumber((_306_v98).length));
        let _307_v99;
        let _nw58 = new _module.C2();
        _nw58.__ctor((_dafny.Set.fromElements(_153_v2)).IsProperSubsetOf(_219_v33), (_237_v48).f13);
        _307_v99 = _nw58;
        _307_v99 = _307_v99;
        let _308_v100;
        let _nw59 = new _module.C2();
        _nw59.__ctor(false, _153_v2);
        _308_v100 = _nw59;
        _237_v48 = _237_v48;
      } else {
        (_162_globalState).f0 = (_237_v48).f13;
        (_162_globalState).f7 = _164_v11;
        _152_v1 = _dafny.Seq.of((_module.__default.fm3(_164_v11, (_237_v48).f13, _module.__default.fm3(_164_v11, (_237_v48).f13, _237_v48.f14, _162_globalState), _162_globalState)).minus(_237_v48.f14), new BigNumber(978), _151_v0);
        if ((_237_v48).f13) {
          let _309_v101;
          let _nw60 = new _module.C3();
          _nw60.__ctor(false);
          _309_v101 = _nw60;
          let _310_v102;
          _310_v102 = _dafny.Seq.of(_161_v10, _161_v10, _161_v10);
          (_237_v48).f14 = (new BigNumber((_310_v102).length)).plus((_151_v0).plus(_module.__default.fm3(_164_v11, (_237_v48).f13, _237_v48.f14, _162_globalState)));
          let _311_v103;
          _311_v103 = _dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,(_237_v48).f13);
          let _312_v104;
          let _out9;
          _out9 = _module.__default.m0(_237_v48.f14, _161_v10, (_237_v48).f13, (_dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,(_237_v48).f13)).Merge(_311_v103), _162_globalState);
          _312_v104 = _out9;
          (_237_v48).f14 = _237_v48.f14;
          let _313_v105;
          let _nw61 = new _module.C3();
          _nw61.__ctor((_237_v48).f13);
          _313_v105 = _nw61;
        } else {
          let _314_v106;
          let _nw62 = new _module.C3();
          _nw62.__ctor(_153_v2);
          _314_v106 = _nw62;
          let _315_v107;
          _315_v107 = _dafny.Seq.of(_237_v48);
          let _316_v108;
          _316_v108 = _module.D3.create_DC11(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_314_v106, _314_v106), _module.__default.safeIndex((_dafny.ZERO).minus(_151_v0), new BigNumber((_dafny.Seq.of(_314_v106, _314_v106)).length)), _314_v106)).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(310)), ((_317_v11) => function (_318_i9) {
  return _317_v11;
})(_164_v11)), _237_v48.f14, (_315_v107)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_315_v107).length))]);
          _316_v108 = _316_v108;
          _314_v106 = (_module.D5.create_DC17(_314_v106)).dtor_cf26;
          let _319_v109;
          let _nw63 = new _module.C3();
          _nw63.__ctor((_237_v48).f13);
          _319_v109 = _nw63;
          let _index60 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_156_v5).length));
          (_156_v5)[_index60] = (_237_v48).f13;
          let _index61 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_156_v5).length));
          (_156_v5)[_index61] = (_237_v48).f13;
          let _index62 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_156_v5).length));
          let _index63 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_156_v5).length));
          let _rhs37 = _319_v109;
          let _rhs38 = !(((_153_v2) ? ((_237_v48).f13) : ((_237_v48).f13))) || ((_237_v48).f13);
          let _rhs39 = (_237_v48).f13;
          let _rhs40 = (_237_v48).f13;
          let _lhs26 = _156_v5;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_156_v5).length));
          let _lhs28 = _156_v5;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_156_v5).length));
          _319_v109 = _rhs37;
          _lhs26[_lhs27] = _rhs38;
          _153_v2 = _rhs39;
          _lhs28[_lhs29] = _rhs40;
          (_237_v48).f14 = (new BigNumber((_dafny.Seq.Concat(_152_v1, _152_v1)).length)).multipliedBy(_module.__default.safeDivisionInt((_152_v1)[_module.__default.safeIndex(_237_v48.f14, new BigNumber((_152_v1).length))], _237_v48.f14));
          let _320_v110;
          _320_v110 = _module.D4.create_DC15(new BigNumber(718), true);
          let _321_v111;
          _321_v111 = _dafny.Map.Empty.slice().updateUnsafe((((_320_v110).dtor_cf24) ? (_151_v0) : (new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of((_237_v48).f13, (_319_v109).f9))).length))),(_237_v48).f13);
          let _rhs41 = _161_v10;
          let _rhs42 = _321_v111;
          let _rhs43 = new BigNumber((_219_v33).length);
          let _rhs44 = _164_v11;
          let _lhs30 = _162_globalState;
          _161_v10 = _rhs41;
          _321_v111 = _rhs42;
          _151_v0 = _rhs43;
          _lhs30.f7 = _rhs44;
        }
        _156_v5 = _156_v5;
      }
      let _322_v112;
      let _nw64 = new _module.C3();
      _nw64.__ctor(_153_v2);
      _322_v112 = _nw64;
      let _323_v113;
      _323_v113 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_219_v33).length)));
      let _324_v114;
      _324_v114 = _module.D5.create_DC18(new BigNumber((_323_v113).length), (_237_v48).f13);
      _151_v0 = ((!(((_324_v114).dtor_cf27).isLessThanOrEqualTo(_237_v48.f14))) ? (_151_v0) : (_151_v0));
      if (false) {
        let _325_v115;
        _325_v115 = _dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,_237_v48.f14);
        let _326_v116;
        _326_v116 = _dafny.Map.Empty.slice().updateUnsafe(_151_v0,(_237_v48).f13);
        let _327_v117;
        let _out10;
        _out10 = _module.__default.m0(new BigNumber((_dafny.MultiSet.fromElements(_325_v115, _dafny.Map.Empty.slice().updateUnsafe(_151_v0,_237_v48.f14))).cardinality()), _161_v10, !(_323_v113).equals(_323_v113), _326_v116, _162_globalState);
        _327_v117 = _out10;
        let _328_v118;
        _328_v118 = _module.D0.create_DC3();
        _328_v118 = _328_v118;
        let _329_v119;
        _329_v119 = _dafny.Seq.of(_161_v10, _161_v10);
        let _330_v120;
        let _nw65 = Array((new BigNumber(22)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = _161_v10;
        _nw65[(_dafny.ONE).toNumber()] = _161_v10;
        _nw65[(new BigNumber(2)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(3)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(4)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(5)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(6)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(7)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(8)).toNumber()] = (_329_v119)[_module.__default.safeIndex(_237_v48.f14, new BigNumber((_329_v119).length))];
        _nw65[(new BigNumber(9)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(10)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(11)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(12)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(13)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(14)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(15)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(16)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(17)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(18)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(19)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(20)).toNumber()] = _161_v10;
        _nw65[(new BigNumber(21)).toNumber()] = (_329_v119)[_module.__default.safeIndex(_151_v0, new BigNumber((_329_v119).length))];
        _330_v120 = _nw65;
        let _index64 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_330_v120).length));
        (_330_v120)[_index64] = (_329_v119)[_module.__default.safeIndex(_327_v117, new BigNumber((_329_v119).length))];
        let _index65 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_330_v120).length));
        (_330_v120)[_index65] = _161_v10;
        let _331_v121;
        _331_v121 = _dafny.MultiSet.fromElements(_237_v48);
        let _332_v122;
        let _nw66 = new _module.C2();
        _nw66.__ctor(false, (_237_v48).f13);
        _332_v122 = _nw66;
        let _333_v123;
        _333_v123 = _dafny.Seq.of(_332_v122);
        let _334_v124;
        _334_v124 = _dafny.Map.Empty.slice().updateUnsafe(_331_v121,(_333_v123)[_module.__default.safeIndex(_327_v117, new BigNumber((_333_v123).length))]);
        _334_v124 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_237_v48),_332_v122);
        let _335_v125;
        let _nw67 = new _module.C3();
        _nw67.__ctor((_237_v48).f13);
        _335_v125 = _nw67;
        let _336_v126;
        let _nw68 = new _module.C0();
        _nw68.__ctor((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_335_v125,(_335_v125).f9)).length)).isLessThan(new BigNumber((_dafny.Seq.update(_236_v47, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,_327_v117)).length), new BigNumber((_236_v47).length)), (_237_v48).f13)).length)), _237_v48.f14, true);
        _336_v126 = _nw68;
      } else {
        let _337_v127;
        _337_v127 = _dafny.MultiSet.fromElements(!((_237_v48).f13), (_237_v48).f13, !(true));
        (_162_globalState).f3 = !(!((true) === ((_237_v48).f13)) || ((_337_v127).IsDisjointFrom(_module.__default.fm12(!((_237_v48).f13), (_237_v48).f13, _237_v48.f14, (_237_v48).f13, _162_globalState))));
        let _338_v128;
        _338_v128 = _dafny.Map.Empty.slice().updateUnsafe((_237_v48).f13,true);
        _338_v128 = (_338_v128).update((_237_v48).f13, (_237_v48).f13);
        let _339_v129;
        _339_v129 = _module.D3.create_DC10(_237_v48.f14, _161_v10, false);
        _161_v10 = (_339_v129).dtor_cf12;
        let _340_v130;
        let _nw69 = new _module.C3();
        _nw69.__ctor(!(_237_v48.f14).isEqualTo(_237_v48.f14));
        _340_v130 = _nw69;
        let _341_v131;
        _341_v131 = _dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,_dafny.areEqual(_159_v8, _dafny.Seq.update(_159_v8, _module.__default.safeIndex(_237_v48.f14, new BigNumber((_159_v8).length)), _164_v11)));
        _341_v131 = (_341_v131).update((_237_v48.f14).minus(_237_v48.f14), _153_v2);
      }
      let _hi7 = new BigNumber((_219_v33).length);
      for (let _342_i10 = _151_v0; _342_i10.isLessThan(_hi7); _342_i10 = _342_i10.plus(_dafny.ONE)) {
        let _343_v132;
        _343_v132 = _module.D4.create_DC15(new BigNumber(631), _153_v2);
        let _344_v133;
        _344_v133 = _dafny.Map.Empty.slice().updateUnsafe(_343_v132,_156_v5);
        let _345_v134;
        _345_v134 = _module.D2.create_DC7((((_344_v133).contains(_343_v132)) ? ((_344_v133).get(_343_v132)) : (_156_v5)));
        let _source5 = _345_v134;
        if (_source5.is_DC8) {
          let _346___mcc_h20 = (_source5).cf7;
          let _347___mcc_h21 = (_source5).cf8;
          let _348___mcc_h22 = (_source5).cf9;
          let _349_cf9 = _348___mcc_h22;
          let _350_cf8 = _347___mcc_h21;
          let _351_cf7 = _346___mcc_h20;
          let _352_v135;
          let _nw70 = Array((new BigNumber(5)).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _164_v11;
          _nw70[(_dafny.ONE).toNumber()] = _164_v11;
          _nw70[(new BigNumber(2)).toNumber()] = _164_v11;
          _nw70[(new BigNumber(3)).toNumber()] = _164_v11;
          _nw70[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _352_v135 = _nw70;
          let _353_v136;
          _353_v136 = _dafny.Set.fromElements(_352_v135);
          let _354_v137;
          let _355_v138;
          let _356_v139;
          let _out11;
          let _out12;
          let _out13;
          let _outcollector1 = (_322_v112).m1(_353_v136, _162_globalState);
          _out11 = _outcollector1[0];
          _out12 = _outcollector1[1];
          _out13 = _outcollector1[2];
          _354_v137 = _out11;
          _355_v138 = _out12;
          _356_v139 = _out13;
          let _index66 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_161_v10).length));
          (_161_v10)[_index66] = _151_v0;
          let _index67 = _module.__default.safeIndex(new BigNumber(960), new BigNumber((_161_v10).length));
          (_161_v10)[_index67] = (new BigNumber(353)).multipliedBy((((_155_v4).contains(_152_v1)) ? ((_155_v4).get(_152_v1)) : (_342_i10)));
          _153_v2 = !(!((_237_v48).f13) || ((_324_v114).dtor_cf28)) || ((_237_v48).f13);
          (_237_v48).f14 = _237_v48.f14;
        } else {
          let _357___mcc_h23 = (_source5).cf6;
          let _358_cf6 = _357___mcc_h23;
          (_162_globalState).f0 = _153_v2;
          (_162_globalState).f0 = _153_v2;
          let _359_v140;
          _359_v140 = _dafny.Seq.of(_161_v10, _161_v10, _161_v10, _161_v10);
          let _360_v141;
          _360_v141 = _dafny.Map.Empty.slice().updateUnsafe(_164_v11,_159_v8);
          let _361_v142;
          let _out14;
          _out14 = _module.__default.m0((new BigNumber(209)).minus(_237_v48.f14), (_359_v140)[_module.__default.safeIndex(_237_v48.f14, new BigNumber((_359_v140).length))], !((_237_v48).f13), _module.__default.fm16(_360_v141, _342_i10, _162_globalState), _162_globalState);
          _361_v142 = _out14;
          (_237_v48).f14 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_237_v48.f14).minus(_151_v0), _module.__default.safeModuloInt(new BigNumber(684), (_dafny.ZERO).minus(_237_v48.f14))));
        }
        let _362_v143;
        _362_v143 = _dafny.Map.Empty.slice().updateUnsafe(_342_i10,_237_v48.f14);
        _362_v143 = _362_v143;
        let _363_v144;
        let _init5 = ((_364_v11) => function (_365_i11) {
          return _364_v11;
        })(_164_v11);
        let _nw71 = Array((new BigNumber(25)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw71.length); _i0_5++) {
          _nw71[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _363_v144 = _nw71;
        let _366_v145;
        _366_v145 = _dafny.Set.fromElements(_363_v144);
        let _367_v146;
        let _368_v147;
        let _369_v148;
        let _out15;
        let _out16;
        let _out17;
        let _outcollector2 = (_322_v112).m1(_366_v145, _162_globalState);
        _out15 = _outcollector2[0];
        _out16 = _outcollector2[1];
        _out17 = _outcollector2[2];
        _367_v146 = _out15;
        _368_v147 = _out16;
        _369_v148 = _out17;
        let _index68 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_156_v5).length));
        (_156_v5)[_index68] = (_237_v48).f13;
        let _370_v149;
        _370_v149 = _dafny.Seq.of(_237_v48);
        let _index69 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_156_v5).length));
        (_156_v5)[_index69] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_370_v149, _370_v149), _dafny.Seq.Concat(_370_v149, _370_v149));
      }
      let _371_v150;
      _371_v150 = _module.D5.create_DC19(_219_v33);
      let _372_v151;
      _372_v151 = _dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5);
      let _373_v152;
      let _nw72 = Array((new BigNumber(27)).toNumber());
      _nw72[(_dafny.ZERO).toNumber()] = _372_v151;
      _nw72[(_dafny.ONE).toNumber()] = _372_v151;
      _nw72[(new BigNumber(2)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(3)).toNumber()] = (_372_v151).Merge(_372_v151);
      _nw72[(new BigNumber(4)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(5)).toNumber()] = (_372_v151).update(_371_v150, _156_v5);
      _nw72[(new BigNumber(6)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(7)).toNumber()] = (_372_v151).Merge(_dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5));
      _nw72[(new BigNumber(8)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5)).Merge(_372_v151);
      _nw72[(new BigNumber(10)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(11)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(12)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(13)).toNumber()] = (_372_v151).Merge(_372_v151);
      _nw72[(new BigNumber(14)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5);
      _nw72[(new BigNumber(16)).toNumber()] = (_372_v151).Merge(_372_v151);
      _nw72[(new BigNumber(17)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(18)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5)).Merge(_372_v151);
      _nw72[(new BigNumber(19)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(20)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(21)).toNumber()] = ((_372_v151).update(_371_v150, _156_v5)).Merge(_372_v151);
      _nw72[(new BigNumber(22)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(23)).toNumber()] = (_372_v151).Merge(_372_v151);
      _nw72[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_371_v150,_156_v5);
      _nw72[(new BigNumber(25)).toNumber()] = _372_v151;
      _nw72[(new BigNumber(26)).toNumber()] = _372_v151;
      _373_v152 = _nw72;
      let _index70 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_373_v152).length));
      (_373_v152)[_index70] = _372_v151;
      let _index71 = _module.__default.safeIndex(new BigNumber(310), new BigNumber((_373_v152).length));
      (_373_v152)[_index71] = _372_v151;
      let _374_v153;
      _374_v153 = _dafny.MultiSet.fromElements(_153_v2);
      let _375_v154;
      _375_v154 = _dafny.Map.Empty.slice().updateUnsafe(!((_module.D0.create_DC1(false)).dtor_cf1),((_237_v48).f13) === ((_237_v48).f13));
      let _rhs45 = (_219_v33).IsSubsetOf(_219_v33);
      let _rhs46 = (((_375_v154).contains(_module.__default.fm2(_162_globalState))) ? ((_375_v154).get(_module.__default.fm2(_162_globalState))) : (!((_237_v48).f13)));
      let _rhs47 = (_dafny.MultiSet.fromElements(_153_v2, (_237_v48).f13)).Difference(_374_v153);
      let _rhs48 = _156_v5;
      let _rhs49 = _237_v48.f14;
      let _lhs31 = _162_globalState;
      let _lhs32 = _237_v48;
      _153_v2 = _rhs45;
      _lhs31.f3 = _rhs46;
      _374_v153 = _rhs47;
      _156_v5 = _rhs48;
      _lhs32.f14 = _rhs49;
      if (_153_v2) {
        let _376_v155;
        _376_v155 = _dafny.MultiSet.fromElements(_164_v11);
        let _377_v156;
        let _nw73 = new _module.C2();
        _nw73.__ctor(((((_376_v155).contains(_164_v11)) ? ((_376_v155).get(_164_v11)) : (_151_v0))).isLessThan(_237_v48.f14), (_237_v48).f13);
        _377_v156 = _nw73;
        let _index72 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_161_v10).length));
        (_161_v10)[_index72] = _237_v48.f14;
        let _index73 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_161_v10).length));
        (_161_v10)[_index73] = new BigNumber(((_323_v113).Difference(_323_v113)).length);
        let _378_v157;
        let _nw74 = new _module.C1();
        _nw74.__ctor((_237_v48).f13, _237_v48.f14, (_237_v48).f13);
        _378_v157 = _nw74;
        let _379_v158;
        _379_v158 = _dafny.MultiSet.fromElements(_378_v157, _378_v157);
        let _380_v159;
        let _nw75 = new _module.C0();
        _nw75.__ctor((_237_v48).f13, (_dafny.ZERO).minus(new BigNumber((_379_v158).cardinality())), !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_161_v10)[_module.__default.safeIndex(new BigNumber(152), new BigNumber((_161_v10).length))],_156_v5)).length)).isEqualTo(_237_v48.f14));
        _380_v159 = _nw75;
        let _381_v160;
        let _nw76 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _381_v160 = _nw76;
        let _382_v161;
        _382_v161 = _dafny.Map.Empty.slice().updateUnsafe(_381_v160,(_378_v157).f9);
        let _383_v162;
        let _nw77 = new _module.C1();
        _nw77.__ctor((((_382_v161).contains(_381_v160)) ? ((_382_v161).get(_381_v160)) : ((_380_v159).f13)), new BigNumber((_module.__default.fm11(_module.__default.fm2(_162_globalState), (_237_v48).f13, _162_globalState)).length), _module.__default.fm2(_162_globalState));
        _383_v162 = _nw77;
        let _384_v163;
        _384_v163 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_159_v8).length),_237_v48.f14);
        let _385_v164;
        _385_v164 = _dafny.Seq.of(_383_v162, _383_v162, _383_v162);
        let _rhs50 = (_385_v164)[_module.__default.safeIndex((_161_v10)[_module.__default.safeIndex(new BigNumber(152), new BigNumber((_161_v10).length))], new BigNumber((_385_v164).length))];
        let _rhs51 = _module.__default.fm2(_162_globalState);
        let _rhs52 = _dafny.Map.Empty.slice().updateUnsafe(_151_v0,new BigNumber((_159_v8).length));
        let _lhs33 = _377_v156;
        _383_v162 = _rhs50;
        _lhs33.f10 = _rhs51;
        _384_v163 = _rhs52;
        let _386_v165;
        _386_v165 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_237_v48.f14,_237_v48.f14),_380_v159.f14);
        (_237_v48).f14 = (((_386_v165).contains(_384_v163)) ? ((_386_v165).get(_384_v163)) : (_237_v48.f14));
      } else {
        _160_v9 = (_160_v9).update(_153_v2, _237_v48.f14);
        let _387_v166;
        _387_v166 = _dafny.Set.fromElements(_156_v5, _156_v5);
        _387_v166 = _387_v166;
        (_162_globalState).f3 = !(_153_v2) || ((_237_v48).f13);
        (_237_v48).f14 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(522)), ((_388_v11) => function (_389_i12) {
          return _388_v11;
        })(_164_v11))).length);
        let _390_v167;
        let _nw78 = new _module.C1();
        _nw78.__ctor((_237_v48).f13, _237_v48.f14, _153_v2);
        _390_v167 = _nw78;
        let _391_v168;
        let _nw79 = Array((new BigNumber(3)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = _390_v167;
        _nw79[(_dafny.ONE).toNumber()] = _390_v167;
        _nw79[(new BigNumber(2)).toNumber()] = _390_v167;
        _391_v168 = _nw79;
        let _392_v169;
        _392_v169 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((_237_v48).f13, (_237_v48).f13, (_237_v48).f13, (_237_v48).f13));
        let _index74 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_156_v5).length));
        (_156_v5)[_index74] = _153_v2;
        let _393_v170;
        let _nw80 = new _module.C3();
        _nw80.__ctor((_237_v48).f13);
        _393_v170 = _nw80;
        let _index75 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_156_v5).length));
        let _rhs53 = _391_v168;
        let _rhs54 = (_392_v169).Intersect((_392_v169).Union(_dafny.MultiSet.fromElements(_374_v153)));
        let _rhs55 = (((_237_v48).f13) ? ((_237_v48).f13) : ((_390_v167).fm8((_393_v170).f9, _162_globalState)));
        let _rhs56 = _393_v170;
        let _lhs34 = _156_v5;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_156_v5).length));
        _391_v168 = _rhs53;
        _392_v169 = _rhs54;
        _lhs34[_lhs35] = _rhs55;
        _393_v170 = _rhs56;
      }
      let _hi8 = _237_v48.f14;
      for (let _394_i13 = (_dafny.ZERO).minus(_237_v48.f14); _394_i13.isLessThan(_hi8); _394_i13 = _394_i13.plus(_dafny.ONE)) {
        _153_v2 = !((_237_v48).f13) || ((_237_v48).f13);
        let _index76 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_156_v5).length));
        (_156_v5)[_index76] = (_237_v48).f13;
        let _index77 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_156_v5).length));
        (_156_v5)[_index77] = (((_237_v48).f13) ? (false) : ((_237_v48).f13));
        let _395_v171;
        _395_v171 = _dafny.Map.Empty.slice().updateUnsafe(true,_236_v47);
        let _396_v172;
        _396_v172 = _dafny.Map.Empty.slice().updateUnsafe(_394_i13,(_156_v5)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_156_v5).length))]);
        let _nw81 = Array((new BigNumber(9)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = new BigNumber(-787);
        _nw81[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_237_v48.f14));
        _nw81[(new BigNumber(2)).toNumber()] = (_237_v48).fm4(_395_v171, _153_v2, (_237_v48).f13, (_322_v112).fm4(_395_v171, (_237_v48).f13, false, _237_v48.f14, _162_globalState), _162_globalState);
        _nw81[(new BigNumber(3)).toNumber()] = _237_v48.f14;
        _nw81[(new BigNumber(4)).toNumber()] = (_237_v48.f14).minus(new BigNumber((_374_v153).cardinality()));
        _nw81[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_237_v48.f14));
        _nw81[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_375_v154).contains((_156_v5)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_156_v5).length))])) ? ((_375_v154).get((_156_v5)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_156_v5).length))])) : (_module.__default.fm2(_162_globalState))),_396_v172)).length);
        _nw81[(new BigNumber(7)).toNumber()] = _394_i13;
        _nw81[(new BigNumber(8)).toNumber()] = _237_v48.f14;
        _161_v10 = _nw81;
        (_162_globalState).f7 = _module.__default.fm18(_162_globalState);
      }
      process.stdout.write(_dafny.toString(_151_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_152_v1, _dafny.Seq.of(new BigNumber(636), new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_153_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_154_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(636), new BigNumber(636))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v4).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_157_v6).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v7).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_159_v8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_161_v10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f5).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636)), _dafny.Seq.of(new BigNumber(636), new BigNumber(636))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_162_globalState).f6)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_162_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_162_globalState).f8, _dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_164_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v33).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_236_v47, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_237_v48).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_237_v48.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_304_v96).equals(_dafny.MultiSet.fromElements(new BigNumber(636)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_323_v113).equals(_dafny.Set.fromElements(new BigNumber(-2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_324_v114).dtor_cf27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_324_v114).dtor_cf28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_371_v150).dtor_cf29).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_372_v151).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[_dafny.ZERO]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[_dafny.ONE]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(2)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(3)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(4)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(5)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(6)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(7)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(8)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(9)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(10)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(11)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(12)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(13)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(14)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(15)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(16)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(17)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(18)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(19)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(20)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(21)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(22)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(23)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(24)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(25)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_373_v152)[new BigNumber(26)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_374_v153).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_375_v154).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
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
    static create_DC3() {
      let $dt = new D0(2);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
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
        return "D0.DC3";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false);
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
    static create_DC5() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC6() {
      let $dt = new D1(1);
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D1(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5";
      } else if (this.$tag === 1) {
        return "D1.DC6";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5();
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
    static create_DC8(cf7, cf8, cf9) {
      let $dt = new D2(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC7(cf6) {
      let $dt = new D2(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(false, _dafny.ZERO, false);
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
    static create_DC10(cf11, cf12, cf13) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC11(cf14, cf15, cf16, cf17) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf10) {
      let $dt = new D3(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D3(3);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf14) + ", " + this.cf15.toVerbatimString(true) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO, [], false);
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
    static create_DC14(cf20, cf21, cf22) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC15(cf23, cf24) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D4(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC16(cf25) {
      let $dt = new D4(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf19 === other.cf19;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC18(cf27, cf28) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC19(cf29) {
      let $dt = new D5(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC17(cf26) {
      let $dt = new D5(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf26 === other.cf26;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC18(_dafny.ZERO, false);
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
    static create_DC21(cf31, cf32, cf33, cf34, cf35) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC22(cf36) {
      let $dt = new D6(1);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC23(cf37, cf38, cf39) {
      let $dt = new D6(2);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC20(cf30) {
      let $dt = new D6(3);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D6(4);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC20() { return this.$tag === 3; }
    get is_DC24() { return this.$tag === 4; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC21" + "(" + this.cf31.toVerbatimString(true) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC22" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf37 === other.cf37 && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC21(_dafny.Seq.UnicodeFromString(""), false, null, _dafny.Map.Empty, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC26(cf42) {
      let $dt = new D7(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC25(cf41) {
      let $dt = new D7(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC26" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC25" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC26(_dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
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
      this.f3 = false;
      this.f7 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f4 = false;
      this._f5 = _dafny.MultiSet.Empty;
      this._f6 = [];
      this._f8 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f9 = false;
      this.f14 = _dafny.ZERO;
      this._f13 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f13, f14, f9) {
      let _this = this;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f9 = f9;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_this.f14, (_dafny.ZERO).minus(new BigNumber(-993)));
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f9 = false;
      this._f11 = false;
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f11, f12, f9) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f9 = f9;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this).f12).multipliedBy((_this).f12);
    };
    fm8(p0, globalState) {
      let _this = this;
      return ((_this).f12).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f12,!(false))).length), (_this).f12));
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f11;
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _397_v0;
      let _nw82 = Array((new BigNumber(15)).toNumber()).fill(_module.D0.Default());
      _397_v0 = _nw82;
      let _398_v1;
      _398_v1 = _module.D0.create_DC3();
      let _index78 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length));
      (_397_v0)[_index78] = _398_v1;
      let _399_v2;
      let _init6 = function (_400_i0) {
        return (_this).f11;
      };
      let _nw83 = Array((new BigNumber(24)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw83.length); _i0_6++) {
        _nw83[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _399_v2 = _nw83;
      let _index79 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
      (_399_v2)[_index79] = ((_this).f12).isLessThan((_this).f12);
      let _401_v3;
      _401_v3 = new BigNumber(229);
      let _402_v4;
      _402_v4 = _dafny.MultiSet.fromElements((_this).f9, (_this).f9, (_this).f9);
      let _403_v5;
      _403_v5 = _dafny.Seq.of(_402_v4);
      let _index80 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length));
      let _index81 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
      let _rhs57 = (_this).fm8((_this).f9, globalState);
      let _rhs58 = _398_v1;
      let _rhs59 = (_this).f9;
      let _rhs60 = new BigNumber(((_403_v5)[_module.__default.safeIndex(((_this).f12).minus(new BigNumber(-958)), new BigNumber((_403_v5).length))]).cardinality());
      let _lhs36 = globalState;
      let _lhs37 = _397_v0;
      let _lhs38 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length));
      let _lhs39 = _399_v2;
      let _lhs40 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
      _lhs36.f0 = _rhs57;
      _lhs37[_lhs38] = _rhs58;
      _lhs39[_lhs40] = _rhs59;
      _401_v3 = _rhs60;
      let _source6 = (((_this).f11) ? ((_397_v0)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length))]) : (_398_v1));
      if (_source6.is_DC1) {
        let _404___mcc_h0 = (_source6).cf1;
        let _405_cf1 = _404___mcc_h0;
        let _406_v6;
        let _nw84 = new _module.C0();
        _nw84.__ctor((_this).f11, p1, (_this).f9);
        _406_v6 = _nw84;
        let _407_v7;
        let _nw85 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _407_v7 = _nw85;
        let _index82 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_407_v7).length));
        (_407_v7)[_index82] = (_401_v3).minus(new BigNumber(316));
        let _index83 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_407_v7).length));
        (_407_v7)[_index83] = _406_v6.f14;
        let _408_v8;
        _408_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
        let _409_v9;
        let _out18;
        _out18 = _module.__default.m0(_406_v6.f14, _407_v7, (_this).fm9((_407_v7)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_407_v7).length))], (_406_v6).f13, (_this).f12, globalState), _408_v8, globalState);
        _409_v9 = _out18;
        (globalState).f0 = (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))];
      } else if (_source6.is_DC2) {
        let _410___mcc_h1 = (_source6).cf2;
        let _411___mcc_h2 = (_source6).cf3;
        let _412___mcc_h3 = (_source6).cf4;
        let _413_cf4 = _412___mcc_h3;
        let _414_cf3 = _411___mcc_h2;
        let _415_cf2 = _410___mcc_h1;
        _415_cf2 = (_this).f11;
        _399_v2 = _399_v2;
        let _416_v10;
        _416_v10 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_402_v4),(_this).f11);
        let _417_v11;
        _417_v11 = _dafny.Seq.of(true);
        let _418_v12;
        _418_v12 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_417_v11), _402_v4);
        _416_v10 = (_416_v10).update(_418_v12, !((_this).f9));
        let _419_v13;
        _419_v13 = _module.D0.create_DC1(_dafny.Seq.IsPrefixOf(_417_v11, _417_v11));
        let _source7 = _419_v13;
        if (_source7.is_DC1) {
          let _420___mcc_h5 = (_source7).cf1;
          let _421_cf1 = _420___mcc_h5;
          _401_v3 = _414_cf3;
          (globalState).f0 = (_417_v11)[_module.__default.safeIndex((_this).f12, new BigNumber((_417_v11).length))];
          let _422_v14;
          _422_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_415_cf2);
          let _423_v15;
          _423_v15 = _dafny.Seq.of(_417_v11, _417_v11, _dafny.Seq.of((((_422_v14).contains(new BigNumber(874))) ? ((_422_v14).get(new BigNumber(874))) : ((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]))));
          let _424_v16;
          let _nw86 = Array((new BigNumber(4)).toNumber());
          _nw86[(_dafny.ZERO).toNumber()] = _417_v11;
          _nw86[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_417_v11, _417_v11);
          _nw86[(new BigNumber(2)).toNumber()] = _417_v11;
          _nw86[(new BigNumber(3)).toNumber()] = (_423_v15)[_module.__default.safeIndex(_414_cf3, new BigNumber((_423_v15).length))];
          _424_v16 = _nw86;
          _424_v16 = _424_v16;
          let _rhs61 = _422_v14;
          let _rhs62 = ((((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_415_cf2)).length)))) : ((_this).f12))).multipliedBy((_this).f12);
          let _rhs63 = (((_this).fm8((_this).f11, globalState)) ? (_399_v2) : (_399_v2));
          _422_v14 = _rhs61;
          _414_cf3 = _rhs62;
          _399_v2 = _rhs63;
        } else if (_source7.is_DC2) {
          let _425___mcc_h6 = (_source7).cf2;
          let _426___mcc_h7 = (_source7).cf3;
          let _427___mcc_h8 = (_source7).cf4;
          let _428_cf4 = _427___mcc_h8;
          let _429_cf3 = _426___mcc_h7;
          let _430_cf2 = _425___mcc_h6;
          let _431_v17;
          _431_v17 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f9),(_this).f11);
          let _432_v18;
          _432_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_401_v3);
          let _index84 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          let _index85 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          let _rhs64 = ((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) === ((_this).f9);
          let _rhs65 = (_417_v11)[_module.__default.safeIndex(new BigNumber((_431_v17).length), new BigNumber((_417_v11).length))];
          let _rhs66 = (_this).f11;
          let _rhs67 = _401_v3;
          let _rhs68 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(false), _module.__default.safeIndex(_401_v3, new BigNumber((_dafny.Seq.of(false)).length)), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]), _417_v11), (((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) ? (_dafny.Seq.of(_430_cf2, (_this).f9, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))])) : (_dafny.Seq.update(_dafny.Seq.update(_module.__default.fm10(globalState), _module.__default.safeIndex(new BigNumber((_432_v18).length), new BigNumber((_module.__default.fm10(globalState)).length)), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]), _module.__default.safeIndex(_401_v3, new BigNumber((_dafny.Seq.update(_module.__default.fm10(globalState), _module.__default.safeIndex(new BigNumber((_432_v18).length), new BigNumber((_module.__default.fm10(globalState)).length)), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))])).length)), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]))));
          let _lhs41 = globalState;
          let _lhs42 = _399_v2;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          let _lhs44 = _399_v2;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          _lhs41.f3 = _rhs64;
          _lhs42[_lhs43] = _rhs65;
          _lhs44[_lhs45] = _rhs66;
          _429_cf3 = _rhs67;
          _417_v11 = _rhs68;
          (globalState).f0 = _430_cf2;
          _401_v3 = p1;
          let _433_v19;
          _433_v19 = _dafny.Map.Empty.slice().updateUnsafe(_430_cf2,p2);
          let _434_v21;
          _434_v21 = _dafny.Set.fromElements(_401_v3, (_this).f12, new BigNumber((function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of _dafny.IntegerRange(new BigNumber(-922), new BigNumber(289))) {
              let _435_v20 = _compr_10;
              if (((new BigNumber(-922)).isLessThanOrEqualTo(_435_v20)) && ((_435_v20).isLessThan(new BigNumber(289)))) {
                _coll10.push([(_435_v20).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("aee"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(640)), function (_436_i1) {
                  return new _dafny.CodePoint('b'.codePointAt(0));
                }))).length)),(_this).f9]);
              }
            }
            return _coll10;
          }()).length), _429_cf3, (_this).f12);
          let _437_v22;
          _437_v22 = _dafny.Seq.of(_434_v21);
          _433_v19 = (_433_v19).update((_434_v21).IsProperSubsetOf((_437_v22)[_module.__default.safeIndex((_this).f12, new BigNumber((_437_v22).length))]), _dafny.Seq.Create(_module.__default.abs(new BigNumber(913)), ((_438_v3) => function (_439_i2) {
            return _438_v3;
          })(_401_v3)));
        } else if (_source7.is_DC3) {
          let _440_v23;
          _440_v23 = new _dafny.CodePoint('q'.codePointAt(0));
          let _441_v24;
          _441_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_417_v11);
          let _442_v25;
          _442_v25 = _dafny.Map.Empty.slice().updateUnsafe(_440_v23,(_this).fm4(_441_v24, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_this).fm8(_415_cf2, globalState), (_this).f12, globalState));
          _442_v25 = (_442_v25).update(_440_v23, (_this).f12);
          let _443_v26;
          _443_v26 = _dafny.Map.Empty.slice().updateUnsafe((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))],!(_415_cf2));
          _443_v26 = (_443_v26).update((_this).fm8((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], globalState), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
          let _444_v27;
          let _nw87 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
          _444_v27 = _nw87;
          _444_v27 = _444_v27;
          let _445_v28;
          _445_v28 = _dafny.Seq.of(_401_v3, (_this).f12, _module.__default.safeModuloInt(_401_v3, (_this).f12));
          _445_v28 = p2;
        } else {
          let _446___mcc_h9 = (_source7).cf0;
          let _447_cf0 = _446___mcc_h9;
          let _448_v29;
          _448_v29 = _dafny.Map.Empty.slice().updateUnsafe((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))],_414_cf3);
          let _449_v30;
          _449_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_448_v29);
          let _450_v31;
          _450_v31 = _module.D0.create_DC2((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], new BigNumber((_449_v30).length), _413_cf4);
          let _451_v32;
          _451_v32 = _dafny.MultiSet.fromElements(_450_v31, _450_v31);
          let _452_v33;
          _452_v33 = _dafny.Map.Empty.slice().updateUnsafe(_414_cf3,_415_cf2);
          let _453_v34;
          _453_v34 = _dafny.Set.fromElements(_414_cf3, new BigNumber((_452_v33).length));
          let _454_v35;
          _454_v35 = new _dafny.CodePoint('k'.codePointAt(0));
          let _455_v36;
          _455_v36 = _module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_456_p1) => function (_457_i3) {
  return _456_p1;
})(p1)));
          let _rhs69 = (_451_v32).IsProperSubsetOf((((_this).f9) ? (_451_v32) : (_451_v32)));
          let _rhs70 = ((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) === ((_453_v34).equals(_453_v34));
          let _rhs71 = _454_v35;
          let _rhs72 = (_this).f9;
          let _rhs73 = new BigNumber(((_455_v36).dtor_cf5).length);
          let _lhs46 = globalState;
          let _lhs47 = globalState;
          let _lhs48 = globalState;
          _415_cf2 = _rhs69;
          _lhs46.f3 = _rhs70;
          _lhs47.f7 = _rhs71;
          _lhs48.f3 = _rhs72;
          _414_cf3 = _rhs73;
          let _458_v37;
          _458_v37 = _dafny.Set.fromElements(_399_v2, _399_v2, _399_v2);
          let _459_v38;
          _459_v38 = _dafny.Seq.of(_458_v37, _458_v37, _458_v37, (_458_v37).Union(_458_v37));
          let _460_v39;
          _460_v39 = _dafny.Seq.UnicodeFromString("rrmm");
          _458_v37 = (_459_v38)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm3((_460_v39)[_module.__default.safeIndex(_414_cf3, new BigNumber((_460_v39).length))], (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], new BigNumber(-349), globalState)), new BigNumber((_459_v38).length))];
          let _index86 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          (_399_v2)[_index86] = (_450_v31).dtor_cf2;
          let _461_v40;
          _461_v40 = _dafny.Map.Empty.slice().updateUnsafe((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))],_dafny.Seq.Concat(_module.__default.fm11((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], globalState), _460_v39));
          _447_cf0 = new BigNumber((_461_v40).length);
        }
      } else if (_source6.is_DC3) {
        let _462_v41;
        _462_v41 = _dafny.MultiSet.fromElements(_401_v3, p1, _401_v3);
        let _463_v42;
        _463_v42 = _module.D0.create_DC1((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
        let _464_v43;
        let _nw88 = new _module.C0();
        _nw88.__ctor((_this).fm9(p1, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], p1, globalState), (((_this).f11) ? ((((_462_v41).contains(new BigNumber(973))) ? ((_462_v41).get(new BigNumber(973))) : (_401_v3))) : (_401_v3)), (_463_v42).dtor_cf1);
        _464_v43 = _nw88;
        _464_v43 = _464_v43;
        let _465_v44;
        _465_v44 = _dafny.Seq.UnicodeFromString("h");
        let _466_v45;
        _466_v45 = _dafny.Map.Empty.slice().updateUnsafe(_465_v44,_464_v43.f14);
        let _467_v46;
        _467_v46 = _dafny.Map.Empty.slice().updateUnsafe((_466_v45).Merge(_466_v45),(_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
        _467_v46 = (_467_v46).update(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("lh"),_401_v3), (_this).f11);
        if (((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("hgfq"), _465_v44)) ? ((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) : ((_this).f11))) {
          let _468_v47;
          _468_v47 = _dafny.Seq.of(true);
          let _469_v48;
          _469_v48 = _dafny.Map.Empty.slice().updateUnsafe(false,(_468_v47)[_module.__default.safeIndex((_this).f12, new BigNumber((_468_v47).length))]);
          let _470_v49;
          let _nw89 = Array((new BigNumber(16)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = new BigNumber(((_462_v41).Difference(_462_v41)).cardinality());
          _nw89[(_dafny.ONE).toNumber()] = (((_464_v43).f13) ? (_401_v3) : (new BigNumber((_465_v44).length)));
          _nw89[(new BigNumber(2)).toNumber()] = _401_v3;
          _nw89[(new BigNumber(3)).toNumber()] = new BigNumber((_465_v44).length);
          _nw89[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(_401_v3);
          _nw89[(new BigNumber(5)).toNumber()] = new BigNumber(-221);
          _nw89[(new BigNumber(6)).toNumber()] = (((_464_v43).f13) ? (_464_v43.f14) : (_464_v43.f14));
          _nw89[(new BigNumber(7)).toNumber()] = new BigNumber((_465_v44).length);
          _nw89[(new BigNumber(8)).toNumber()] = _464_v43.f14;
          _nw89[(new BigNumber(9)).toNumber()] = _464_v43.f14;
          _nw89[(new BigNumber(10)).toNumber()] = _464_v43.f14;
          _nw89[(new BigNumber(11)).toNumber()] = _464_v43.f14;
          _nw89[(new BigNumber(12)).toNumber()] = (((_this).f11) ? (new BigNumber((_469_v48).length)) : (new BigNumber((_dafny.Set.fromElements(_464_v43, _464_v43)).length)));
          _nw89[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_464_v43.f14);
          _nw89[(new BigNumber(14)).toNumber()] = p1;
          _nw89[(new BigNumber(15)).toNumber()] = new BigNumber(728);
          _470_v49 = _nw89;
          let _index87 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_470_v49).length));
          (_470_v49)[_index87] = ((_this).f12).minus(p1);
          let _index88 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_470_v49).length));
          (_470_v49)[_index88] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_464_v43.f14, new BigNumber((_dafny.Seq.Concat(_465_v44, _465_v44)).length)));
          let _471_v50;
          _471_v50 = _module.D2.create_DC7(_399_v2);
          let _472_v51;
          _472_v51 = _dafny.Seq.of(_399_v2, _399_v2, (_471_v50).dtor_cf6);
          _472_v51 = (((_this).f11) ? ((((_this).f9) ? (_472_v51) : (_472_v51))) : (_472_v51));
          let _473_v52;
          _473_v52 = _dafny.MultiSet.fromElements(_462_v41, _462_v41, _462_v41);
          let _474_v53;
          _474_v53 = _module.D0.create_DC2((_464_v43).f13, p1, _473_v52);
          let _475_v54;
          _475_v54 = _module.D2.create_DC8((_474_v53).dtor_cf2, (_this).f12, (_464_v43).f13);
          let _476_v55;
          let _nw90 = new _module.C0();
          _nw90.__ctor((_this).f11, (_475_v54).dtor_cf8, (_464_v43).f13);
          _476_v55 = _nw90;
          let _477_v56;
          _477_v56 = new _dafny.CodePoint('i'.codePointAt(0));
          (globalState).f7 = _477_v56;
          (_464_v43).f14 = _464_v43.f14;
        } else {
          let _478_v57;
          _478_v57 = _dafny.Map.Empty.slice().updateUnsafe(_464_v43.f14,_module.__default.safeModuloInt(new BigNumber(186), _464_v43.f14));
          let _479_v58;
          _479_v58 = new _dafny.CodePoint('r'.codePointAt(0));
          let _480_v59;
          _480_v59 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_module.__default.fm12(true, (_464_v43).f13, _464_v43.f14, true, globalState));
          let _481_v60;
          _481_v60 = _dafny.Seq.of((_this).f11, (_this).f11, (_464_v43).f13, (_dafny.Map.Empty.slice().updateUnsafe(_479_v58,_402_v4)).equals(_480_v59));
          let _482_v61;
          _482_v61 = _dafny.Map.Empty.slice().updateUnsafe((_464_v43).f13,_479_v58);
          let _483_v62;
          _483_v62 = _dafny.Set.fromElements((_this).fm8((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], globalState));
          let _index89 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          let _rhs74 = !((_481_v60)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_464_v43.f14, p1), new BigNumber((_481_v60).length))]);
          let _rhs75 = (new BigNumber(((_482_v61).Merge(_482_v61)).length)).plus(new BigNumber((_483_v62).length));
          let _rhs76 = (_464_v43).f13;
          let _rhs77 = _dafny.Map.Empty.slice().updateUnsafe(p1,(new BigNumber((_465_v44).length)).multipliedBy(_464_v43.f14));
          let _lhs49 = globalState;
          let _lhs50 = _464_v43;
          let _lhs51 = _399_v2;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
          _lhs49.f0 = _rhs74;
          _lhs50.f14 = _rhs75;
          _lhs51[_lhs52] = _rhs76;
          _478_v57 = _rhs77;
          (_464_v43).f14 = _464_v43.f14;
          (_464_v43).f14 = (_this).f12;
          _465_v44 = _465_v44;
          let _484_v63;
          _484_v63 = _dafny.Map.Empty.slice().updateUnsafe(_464_v43.f14,_dafny.Seq.UnicodeFromString("xncydgy"));
          let _485_v64;
          _485_v64 = _module.D0.create_DC0(_464_v43.f14);
          _484_v63 = _dafny.Map.Empty.slice().updateUnsafe((_485_v64).dtor_cf0,_module.__default.fm11((_464_v43).f13, (_this).f11, globalState));
        }
        _465_v44 = _dafny.Seq.Concat(_465_v44, _465_v44);
      } else {
        let _486___mcc_h4 = (_source6).cf0;
        let _487_cf0 = _486___mcc_h4;
        let _488_v65;
        _488_v65 = new _dafny.CodePoint('j'.codePointAt(0));
        let _489_v66;
        _489_v66 = _dafny.Seq.UnicodeFromString("p");
        let _490_v67;
        _490_v67 = _dafny.MultiSet.fromElements(_module.__default.fm3(_488_v65, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_this).f12, globalState), new BigNumber((_489_v66).length));
        let _491_v68;
        _491_v68 = _module.D0.create_DC0(_487_cf0);
        _401_v3 = ((((_this).f9) ? (_module.D0.create_DC0((((_490_v67).contains(new BigNumber(278))) ? ((_490_v67).get(new BigNumber(278))) : ((_this).f12)))) : (_491_v68))).dtor_cf0;
        let _492_v69;
        let _nw91 = new _module.C0();
        _nw91.__ctor(false, (((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_487_cf0)))) : (_401_v3)), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
        _492_v69 = _nw91;
        let _493_v70;
        _493_v70 = _dafny.Set.fromElements(true);
        _493_v70 = (_493_v70).Union(_493_v70);
        let _494_v71;
        _494_v71 = _dafny.Map.Empty.slice().updateUnsafe((_492_v69).f13,_492_v69.f14);
        let _495_v72;
        _495_v72 = _module.D3.create_DC11(_487_cf0, _489_v66, new BigNumber((_494_v71).length), _492_v69);
        let _496_v73;
        _496_v73 = _dafny.Seq.of((_495_v72).dtor_cf15, _489_v66, _489_v66);
        let _497_v74;
        _497_v74 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(692)), ((_498_v3) => function (_499_i4) {
          return _498_v3;
        })(_401_v3)), p2),_496_v73);
        _496_v73 = (((_497_v74).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-175)), ((_502_p1) => function (_503_i5) {
          return _502_p1;
        })(p1)))) ? ((_497_v74).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-175)), ((_500_p1) => function (_501_i5) {
          return _500_p1;
        })(p1)))) : (_496_v73));
      }
      let _504_v75;
      let _init7 = function (_505_i6) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      };
      let _nw92 = Array((new BigNumber(21)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw92.length); _i0_7++) {
        _nw92[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _504_v75 = _nw92;
      let _506_v76;
      _506_v76 = new _dafny.CodePoint('e'.codePointAt(0));
      let _index90 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_504_v75).length));
      (_504_v75)[_index90] = _506_v76;
      let _index91 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_504_v75).length));
      (_504_v75)[_index91] = _506_v76;
      let _507_i7;
      _507_i7 = _dafny.ZERO;
      L1: {
        while ((_this).f9) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_507_i7)) {
              break L1;
            }
            _507_i7 = (_507_i7).plus(_dafny.ONE);
            let _508_v77;
            _508_v77 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_401_v3)).length));
            let _509_v78;
            _509_v78 = _dafny.MultiSet.fromElements(_508_v77);
            let _510_v79;
            _510_v79 = _module.D0.create_DC2((_this).f11, p1, _509_v78);
            let _511_v80;
            let _nw93 = new _module.C0();
            _nw93.__ctor(((_this).f9) && ((_this).f9), (new BigNumber(-293)).plus((_510_v79).dtor_cf3), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
            _511_v80 = _nw93;
            let _512_v81;
            let _nw94 = Array((_dafny.ONE).toNumber());
            _nw94[(_dafny.ZERO).toNumber()] = p2;
            _512_v81 = _nw94;
            let _index92 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length));
            (_512_v81)[_index92] = _module.__default.fm13((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], true, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], _511_v80.f14, globalState);
            let _513_v82;
            _513_v82 = _dafny.Seq.of((_this).f9);
            let _index93 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length));
            let _rhs78 = _module.__default.fm3((_504_v75)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_504_v75).length))], !((_511_v80).f13), new BigNumber((_dafny.Seq.Concat(_513_v82, _513_v82)).length), globalState);
            let _rhs79 = _401_v3;
            let _rhs80 = p2;
            let _lhs53 = _511_v80;
            let _lhs54 = _511_v80;
            let _lhs55 = _512_v81;
            let _lhs56 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length));
            _lhs53.f14 = _rhs78;
            _lhs54.f14 = _rhs79;
            _lhs55[_lhs56] = _rhs80;
            let _514_v83;
            _514_v83 = _module.D1.create_DC6();
            let _source8 = (((_511_v80).f13) ? (_514_v83) : (_514_v83));
            if (_source8.is_DC5) {
              let _515_v84;
              _515_v84 = _dafny.Map.Empty.slice().updateUnsafe(_399_v2,_511_v80.f14);
              let _516_v85;
              _516_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_513_v82).length),_511_v80.f14);
              _515_v84 = (_515_v84).update(_399_v2, new BigNumber(((_516_v85).Merge(_516_v85)).length));
              let _517_v86;
              let _nw95 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _517_v86 = _nw95;
              let _index94 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_517_v86).length));
              (_517_v86)[_index94] = _module.__default.safeDivisionInt(new BigNumber(950), _511_v80.f14);
              let _index95 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_517_v86).length));
              (_517_v86)[_index95] = _module.__default.safeDivisionInt(_511_v80.f14, ((true) ? (_511_v80.f14) : (p1)));
              let _518_v87;
              _518_v87 = _dafny.Map.Empty.slice().updateUnsafe(true,_506_v76);
              let _519_v88;
              _519_v88 = _module.D0.create_DC0(p1);
              let _520_v89;
              _520_v89 = _dafny.Map.Empty.slice().updateUnsafe((((_518_v87).contains((_this).f11)) ? ((_518_v87).get((_this).f11)) : (_506_v76)),_519_v88);
              _520_v89 = (_520_v89).update((_504_v75)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_504_v75).length))], _519_v88);
              let _index96 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_517_v86).length));
              (_517_v86)[_index96] = ((_this).f12).multipliedBy(p1);
            } else if (_source8.is_DC6) {
              let _521_v90;
              _521_v90 = _dafny.Seq.UnicodeFromString("kpamgbonn");
              let _522_v91;
              _522_v91 = _dafny.Set.fromElements((_dafny.MultiSet.fromElements((_511_v80).f13)).IsDisjointFrom(_402_v4), (_511_v80).f13, (_511_v80).f13);
              let _523_v92;
              _523_v92 = _dafny.Set.fromElements((_dafny.ZERO).minus(_511_v80.f14));
              let _524_v93;
              _524_v93 = _dafny.Seq.of(_523_v92, _523_v92, _523_v92, _523_v92);
              let _525_v94;
              _525_v94 = _dafny.Map.Empty.slice().updateUnsafe((_511_v80).f13,_522_v91);
              let _rhs81 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_521_v90, _521_v90), _module.__default.safeIndex(new BigNumber(753), new BigNumber((_dafny.Seq.Concat(_521_v90, _521_v90)).length)), _506_v76), _521_v90);
              let _rhs82 = (_dafny.MultiSet.FromArray((((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]) ? (_524_v93) : (_524_v93)))).IsSubsetOf(_dafny.MultiSet.fromElements(_523_v92, _dafny.Set.fromElements(p1)));
              let _rhs83 = ((((_525_v94).contains(true)) ? ((_525_v94).get(true)) : (_dafny.Set.fromElements((_this).f11, (_this).f9)))).Difference(_522_v91);
              let _lhs57 = globalState;
              _521_v90 = _rhs81;
              _lhs57.f3 = _rhs82;
              _522_v91 = _rhs83;
              let _526_v95;
              _526_v95 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_511_v80).f13);
              (globalState).f0 = ((_511_v80).f13) || ((((_526_v95).contains((_dafny.ZERO).minus(p1))) ? ((_526_v95).get((_dafny.ZERO).minus(p1))) : ((_this).f9)));
              (globalState).f7 = _506_v76;
              let _527_v96;
              _527_v96 = _dafny.Seq.of(_521_v90, _dafny.Seq.UnicodeFromString("wx"), _521_v90);
              let _528_v97;
              let _nw96 = new _module.C0();
              _nw96.__ctor(_dafny.Seq.IsProperPrefixOf(_527_v96, _527_v96), p1, (_511_v80).f13);
              _528_v97 = _nw96;
              let _index97 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length));
              let _rhs84 = _399_v2;
              let _rhs85 = _528_v97;
              let _rhs86 = !(function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of _dafny.IntegerRange(new BigNumber(736), new BigNumber(-955))) {
                  let _529_v98 = _compr_11;
                  if (((new BigNumber(736)).isLessThanOrEqualTo(_529_v98)) && ((_529_v98).isLessThan(new BigNumber(-955)))) {
                    _coll11.push([_module.__default.safeDivisionInt(_529_v98, _511_v80.f14),function () {
                      let _coll12 = new _dafny.Map();
                      for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_this).f12),new BigNumber((function () {
                        let _coll13 = new _dafny.Map();
                        for (const _compr_13 of ((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))]).Elements) {
                          let _530_v100 = _compr_13;
                          if (_dafny.Seq.contains((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))], _530_v100)) {
                            _coll13.push([(_530_v100).plus((_dafny.ZERO).minus(p1)),_511_v80.f14]);
                          }
                        }
                        return _coll13;
                      }()).length))).Keys.Elements) {
                        let _531_v99 = _compr_12;
                        if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC0((_this).f12),new BigNumber((function () {
                          let _coll14 = new _dafny.Map();
                          for (const _compr_14 of ((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))]).Elements) {
                            let _530_v100 = _compr_14;
                            if (_dafny.Seq.contains((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))], _530_v100)) {
                              _coll14.push([(_530_v100).plus((_dafny.ZERO).minus(p1)),_511_v80.f14]);
                            }
                          }
                          return _coll14;
                        }()).length))).contains(_531_v99)) {
                          _coll12.push([_531_v99,_401_v3]);
                        }
                      }
                      return _coll12;
                    }()]);
                  }
                }
                return _coll11;
              }()).contains(new BigNumber((_522_v91).length));
              let _rhs87 = (_397_v0)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length))];
              let _rhs88 = ((_this).f11) && ((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
              let _lhs58 = globalState;
              let _lhs59 = _397_v0;
              let _lhs60 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_397_v0).length));
              let _lhs61 = globalState;
              _399_v2 = _rhs84;
              _528_v97 = _rhs85;
              _lhs58.f0 = _rhs86;
              _lhs59[_lhs60] = _rhs87;
              _lhs61.f0 = _rhs88;
            } else {
              let _532___mcc_h10 = (_source8).cf5;
              let _533_cf5 = _532___mcc_h10;
              (globalState).f0 = (_this).f9;
              let _534_v101;
              let _nw97 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
              _534_v101 = _nw97;
              let _index98 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_534_v101).length));
              (_534_v101)[_index98] = _module.__default.safeModuloInt(p1, _511_v80.f14);
              let _index99 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_534_v101).length));
              (_534_v101)[_index99] = _401_v3;
              _401_v3 = _module.__default.safeModuloInt(new BigNumber(148), _401_v3);
              let _535_v102;
              _535_v102 = _module.D0.create_DC1((_this).fm8(true, globalState));
              let _536_v103;
              _536_v103 = _dafny.Seq.of((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))]);
              let _537_v104;
              _537_v104 = _dafny.Map.Empty.slice().updateUnsafe(_513_v82,_511_v80.f14);
              let _538_v105;
              _538_v105 = _dafny.Set.fromElements((_dafny.ZERO).minus((_511_v80.f14).multipliedBy(new BigNumber((_536_v103).length))), new BigNumber((_537_v104).length));
              let _pat_let_tv18 = _511_v80;
              let _rhs89 = function (_pat_let11_0) {
                return function (_539_dt__update__tmp_h0) {
                  return function (_pat_let12_0) {
                    return function (_540_dt__update_hcf1_h0) {
                      return _module.D0.create_DC1(_540_dt__update_hcf1_h0);
                    }(_pat_let12_0);
                  }((_pat_let_tv18).f13);
                }(_pat_let11_0);
              }(_535_v102);
              let _rhs90 = _538_v105;
              _535_v102 = _rhs89;
              _538_v105 = _rhs90;
            }
            let _541_v106;
            _541_v106 = _dafny.Seq.of(_506_v76);
            let _542_v107;
            _542_v107 = _module.D3.create_DC9(_541_v106);
            let _543_v108;
            _543_v108 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_541_v106);
            let _index100 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
            let _index101 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
            let _rhs91 = ((_511_v80.f14).plus(new BigNumber(((_542_v107).dtor_cf10).length))).isLessThan(_511_v80.f14);
            let _rhs92 = !(_543_v108).equals(_543_v108);
            let _rhs93 = p1;
            let _rhs94 = (new BigNumber((_dafny.Seq.UnicodeFromString("mm")).length)).minus(new BigNumber(((_512_v81)[_module.__default.safeIndex(new BigNumber(207), new BigNumber((_512_v81).length))]).length));
            let _lhs62 = _399_v2;
            let _lhs63 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
            let _lhs64 = _399_v2;
            let _lhs65 = _module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length));
            _lhs62[_lhs63] = _rhs91;
            _lhs64[_lhs65] = _rhs92;
            _401_v3 = _rhs93;
            _401_v3 = _rhs94;
          }
        }
      }
      let _544_v109;
      _544_v109 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,p1);
      let _545_v110;
      _545_v110 = _dafny.MultiSet.fromElements(new BigNumber(336), (_dafny.ZERO).minus((_this).f12));
      let _546_v111;
      _546_v111 = _dafny.Seq.of(_545_v110);
      let _547_v112;
      let _nw98 = Array((new BigNumber(21)).toNumber());
      _nw98[(_dafny.ZERO).toNumber()] = new BigNumber((_544_v109).length);
      _nw98[(_dafny.ONE).toNumber()] = _401_v3;
      _nw98[(new BigNumber(2)).toNumber()] = p1;
      _nw98[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).f12);
      _nw98[(new BigNumber(4)).toNumber()] = (_this).f12;
      _nw98[(new BigNumber(5)).toNumber()] = (((_402_v4).contains((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))])) ? ((_402_v4).get((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))])) : (_401_v3));
      _nw98[(new BigNumber(6)).toNumber()] = (_this).f12;
      _nw98[(new BigNumber(7)).toNumber()] = new BigNumber(759);
      _nw98[(new BigNumber(8)).toNumber()] = _401_v3;
      _nw98[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p1);
      _nw98[(new BigNumber(10)).toNumber()] = p1;
      _nw98[(new BigNumber(11)).toNumber()] = p1;
      _nw98[(new BigNumber(12)).toNumber()] = new BigNumber(-591);
      _nw98[(new BigNumber(13)).toNumber()] = new BigNumber(((_546_v111)[_module.__default.safeIndex(p1, new BigNumber((_546_v111).length))]).cardinality());
      _nw98[(new BigNumber(14)).toNumber()] = _401_v3;
      _nw98[(new BigNumber(15)).toNumber()] = p1;
      _nw98[(new BigNumber(16)).toNumber()] = _401_v3;
      _nw98[(new BigNumber(17)).toNumber()] = (_this).f12;
      _nw98[(new BigNumber(18)).toNumber()] = new BigNumber((p2).length);
      _nw98[(new BigNumber(19)).toNumber()] = (_this).f12;
      _nw98[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f12);
      _547_v112 = _nw98;
      let _548_v113;
      _548_v113 = _dafny.Map.Empty.slice().updateUnsafe(_547_v112,(_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
      _548_v113 = (_548_v113).update(_547_v112, !(!(_module.__default.fm2(globalState)) || ((_this).f11)));
      let _549_v114;
      _549_v114 = _module.D2.create_DC7(_399_v2);
      let _source9 = _549_v114;
      if (_source9.is_DC8) {
        let _550___mcc_h11 = (_source9).cf7;
        let _551___mcc_h12 = (_source9).cf8;
        let _552___mcc_h13 = (_source9).cf9;
        let _553_cf9 = _552___mcc_h13;
        let _554_cf8 = _551___mcc_h12;
        let _555_cf7 = _550___mcc_h11;
        _554_cf8 = (_this).f12;
        let _556_v115;
        _556_v115 = _dafny.Seq.of(false, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
        let _557_v116;
        _557_v116 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_556_v115);
        let _558_v117;
        let _nw99 = new _module.C0();
        _nw99.__ctor(_553_cf9, _module.__default.safeModuloInt((_this).fm4(_557_v116, (_this).f9, (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_this).f12, globalState), _554_cf8), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]);
        _558_v117 = _nw99;
        _555_cf7 = _module.__default.fm2(globalState);
        let _559_v118;
        _559_v118 = _dafny.Set.fromElements((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (new BigNumber(336)).isLessThanOrEqualTo(p1), (_this).f9);
        _559_v118 = (_559_v118).Union(_dafny.Set.fromElements(!(_553_cf9), (_this).f11, (_558_v117).f13, !((_558_v117).f13), (_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))]));
      } else {
        let _560___mcc_h14 = (_source9).cf6;
        let _561_cf6 = _560___mcc_h14;
        let _562_v119;
        let _nw100 = new _module.C0();
        _nw100.__ctor(false, new BigNumber((_module.__default.fm11((_399_v2)[_module.__default.safeIndex(new BigNumber(328), new BigNumber((_399_v2).length))], (_this).f11, globalState)).length), (_this).f9);
        _562_v119 = _nw100;
        let _563_v120;
        _563_v120 = _dafny.Seq.UnicodeFromString("ji");
        let _564_v121;
        _564_v121 = _module.D3.create_DC11(_562_v119.f14, _563_v120, _562_v119.f14, _562_v119);
        let _565_v122;
        _565_v122 = _module.D3.create_DC11((_dafny.ZERO).minus(_562_v119.f14), _563_v120, _562_v119.f14, (_564_v121).dtor_cf17);
        _562_v119 = (_564_v121).dtor_cf17;
        let _566_v123;
        let _nw101 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
        _566_v123 = _nw101;
        let _index102 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_566_v123).length));
        (_566_v123)[_index102] = (_module.D1.create_DC4(p2)).dtor_cf5;
        let _index103 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_566_v123).length));
        (_566_v123)[_index103] = _dafny.Seq.of(new BigNumber((_563_v120).length));
        _563_v120 = _563_v120;
        let _567_v124;
        _567_v124 = _dafny.Seq.of(!(_401_v3).isEqualTo(new BigNumber(282)), (_this).f11, (_562_v119).f13);
        (globalState).f3 = (_567_v124)[_module.__default.safeIndex((p1).plus(p1), new BigNumber((_567_v124).length))];
      }
      r0 = (_504_v75)[_module.__default.safeIndex(new BigNumber(535), new BigNumber((_504_v75).length))];
      return r0;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f9 = false;
      this.f10 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f10, f9) {
      let _this = this;
      (_this).f10 = f10;
      (_this)._f9 = f9;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      if ((_this).f9) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f9, !((_this).f9))).length), new BigNumber(550), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ofopjn")).length))).length), new BigNumber(983)),(_this).f9)).length);
      } else {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ibh"), _dafny.Seq.UnicodeFromString("fwhmbs"))).length);
      }
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(266);
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      let _source10 = _module.D0.create_DC2((_this).f9, new BigNumber((_dafny.Seq.UnicodeFromString("ekckaebw")).length), _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(690)), function (_568_i0) {
  return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(268), new BigNumber(682)));
})));
      if (_source10.is_DC1) {
        let _569___mcc_h0 = (_source10).cf1;
        let _570_cf1 = _569___mcc_h0;
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(558),_this.f10)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-997),(_this).f9));
      } else if (_source10.is_DC2) {
        let _571___mcc_h1 = (_source10).cf2;
        let _572___mcc_h2 = (_source10).cf3;
        let _573___mcc_h3 = (_source10).cf4;
        let _574_cf4 = _573___mcc_h3;
        let _575_cf3 = _572___mcc_h2;
        let _576_cf2 = _571___mcc_h1;
        return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_575_cf3),_this.f10);
      } else if (_source10.is_DC3) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f10, (_this).f9, true)).cardinality()),_this.f10)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-643)),_this.f10));
      } else {
        let _577___mcc_h4 = (_source10).cf0;
        let _578_cf0 = _577___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(_578_cf0,(_this).f9);
      }
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let _579_v0;
      _579_v0 = _dafny.Seq.UnicodeFromString("wqqru");
      _579_v0 = _dafny.Seq.UnicodeFromString("mysvhxlt");
      let _580_v1;
      _580_v1 = _dafny.Seq.of(_module.__default.fm2(globalState));
      let _581_v2;
      _581_v2 = new BigNumber(579);
      let _582_v3;
      _582_v3 = _dafny.Map.Empty.slice().updateUnsafe(_581_v2,new BigNumber(21));
      let _583_v4;
      _583_v4 = _dafny.MultiSet.fromElements(_this.f10, (_580_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_580_v1, _module.__default.safeIndex(new BigNumber((_582_v3).length), new BigNumber((_580_v1).length)), (_this).f9)).length), new BigNumber((_580_v1).length))], p0);
      let _584_v5;
      _584_v5 = _dafny.MultiSet.fromElements(_581_v2);
      let _585_v6;
      let _nw102 = new _module.C0();
      _nw102.__ctor(_this.f10, ((((_583_v4).contains(p0)) ? ((_583_v4).get(p0)) : (_581_v2))).plus(_581_v2), (_584_v5).IsProperSubsetOf(_584_v5));
      _585_v6 = _nw102;
      _585_v6 = ((!((_585_v6).f9)) ? (_585_v6) : (_585_v6));
      (_this).f10 = _this.f10;
      let _586_v7;
      _586_v7 = _dafny.Set.fromElements((_585_v6).f9);
      _581_v2 = (_585_v6).fm4(_dafny.Map.Empty.slice().updateUnsafe(p0,_580_v1), (_581_v2).isLessThanOrEqualTo(_581_v2), p0, new BigNumber((_586_v7).length), globalState);
      let _587_v8;
      _587_v8 = _dafny.Seq.of(_module.__default.fm11(true, (_585_v6).f9, globalState));
      (globalState).f0 = (new BigNumber((_587_v8).length)).isLessThan(new BigNumber(748));
      let _hi9 = _581_v2;
      for (let _588_i0 = _581_v2; _588_i0.isLessThan(_hi9); _588_i0 = _588_i0.plus(_dafny.ONE)) {
        (globalState).f0 = !(_this.f10);
        _581_v2 = ((_module.__default.fm3((p2)[_module.__default.safeIndex(_581_v2, new BigNumber((p2).length))], p0, _581_v2, globalState)).plus(_588_i0)).multipliedBy(new BigNumber((p2).length));
        let _589_v9;
        let _nw103 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _589_v9 = _nw103;
        let _index104 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_589_v9).length));
        (_589_v9)[_index104] = _588_i0;
        let _index105 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_589_v9).length));
        (_589_v9)[_index105] = _588_i0;
        let _590_v10;
        _590_v10 = _dafny.Map.Empty.slice().updateUnsafe(_585_v6,true);
        let _591_v11;
        let _nw104 = Array((new BigNumber(6)).toNumber());
        _nw104[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_585_v6,p1);
        _nw104[(_dafny.ONE).toNumber()] = _590_v10;
        _nw104[(new BigNumber(2)).toNumber()] = ((true) ? (_590_v10) : (_590_v10));
        _nw104[(new BigNumber(3)).toNumber()] = _590_v10;
        _nw104[(new BigNumber(4)).toNumber()] = _590_v10;
        _nw104[(new BigNumber(5)).toNumber()] = _590_v10;
        _591_v11 = _nw104;
        let _index106 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_591_v11).length));
        (_591_v11)[_index106] = _590_v10;
        let _592_v12;
        let _nw105 = Array((new BigNumber(2)).toNumber());
        _nw105[(_dafny.ZERO).toNumber()] = !_dafny.Seq.contains(_580_v1, _this.f10);
        _nw105[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(p2, _579_v0);
        _592_v12 = _nw105;
        let _index107 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_592_v12).length));
        (_592_v12)[_index107] = _module.__default.fm2(globalState);
        let _593_v13;
        _593_v13 = _dafny.Seq.of(new BigNumber(171), _581_v2);
        let _594_v14;
        _594_v14 = _dafny.Map.Empty.slice().updateUnsafe(_593_v13,_dafny.Seq.Create(_module.__default.abs(new BigNumber(390)), ((_595_v0) => function (_596_i1) {
          return new BigNumber((_595_v0).length);
        })(_579_v0)));
        let _index108 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_591_v11).length));
        let _index109 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_592_v12).length));
        let _rhs95 = _590_v10;
        let _rhs96 = (_585_v6).f9;
        let _rhs97 = _580_v1;
        let _rhs98 = (_this.f10) || ((_585_v6).f9);
        let _rhs99 = (new BigNumber((_594_v14).length)).minus(_588_i0);
        let _lhs66 = _591_v11;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_591_v11).length));
        let _lhs68 = _592_v12;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((_592_v12).length));
        _lhs66[_lhs67] = _rhs95;
        r1 = _rhs96;
        _580_v1 = _rhs97;
        _lhs68[_lhs69] = _rhs98;
        _581_v2 = _rhs99;
      }
      r0 = _586_v7;
      r1 = ((_this.f10) ? (_this.f10) : ((_this).f9));
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.Map.Empty;
      let _597_v1;
      let _init8 = ((_598_p0, _599_p2) => function (_600_i0) {
        return (function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of _dafny.IntegerRange(new BigNumber(427), new BigNumber(24))) {
            let _601_v0 = _compr_15;
            if (((new BigNumber(427)).isLessThanOrEqualTo(_601_v0)) && ((_601_v0).isLessThan(new BigNumber(24)))) {
              _coll15.add((_601_v0).plus(new BigNumber(786)));
            }
          }
          return _coll15;
        }()).IsProperSubsetOf(_dafny.Set.fromElements(_598_p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(407)), function (_602_i1) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })).length), new BigNumber(514), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_599_p2,_this.f10)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-939)), function (_603_i2) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })).length)));
      })(p0, p2);
      let _nw106 = Array((new BigNumber(25)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw106.length); _i0_8++) {
        _nw106[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _597_v1 = _nw106;
      let _604_v2;
      _604_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
      let _605_v3;
      _605_v3 = _dafny.MultiSet.fromElements(new BigNumber((_604_v2).length));
      let _606_v4;
      _606_v4 = _dafny.Seq.of(p0, new BigNumber((_605_v3).cardinality()), new BigNumber(333), p0);
      let _607_v5;
      _607_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_605_v3).contains((_606_v4)[_module.__default.safeIndex(p0, new BigNumber((_606_v4).length))])) ? ((_605_v3).get((_606_v4)[_module.__default.safeIndex(p0, new BigNumber((_606_v4).length))])) : (p0)),true);
      let _608_v6;
      _608_v6 = _dafny.Seq.UnicodeFromString("isc");
      let _609_v7;
      _609_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_607_v5).length),_608_v6);
      let _610_v8;
      _610_v8 = _dafny.MultiSet.fromElements(((_this.f10) ? (_608_v6) : (_608_v6)), _dafny.Seq.UnicodeFromString("djiib"), _608_v6);
      let _611_v9;
      _611_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f10,_609_v7);
      let _612_v10;
      _612_v10 = _dafny.Seq.of(_this.f10, _this.f10);
      let _613_v11;
      _613_v11 = _module.D2.create_DC8(p2, new BigNumber(618), (_this).f9);
      let _pat_let_tv19 = _610_v8;
      let _pat_let_tv20 = _610_v8;
      let _rhs100 = p1;
      let _rhs101 = (_this).f9;
      let _rhs102 = p2;
      let _rhs103 = (_609_v7).Merge((((_611_v9).contains((_612_v10)[_module.__default.safeIndex(p0, new BigNumber((_612_v10).length))])) ? ((_611_v9).get((_612_v10)[_module.__default.safeIndex(p0, new BigNumber((_612_v10).length))])) : (_609_v7)));
      let _rhs104 = function (_source11) {
        if (_source11.is_DC8) {
          let _614___mcc_h0 = (_source11).cf7;
          let _615___mcc_h1 = (_source11).cf8;
          let _616___mcc_h2 = (_source11).cf9;
          let _617_cf9 = _616___mcc_h2;
          let _618_cf8 = _615___mcc_h1;
          let _619_cf7 = _614___mcc_h0;
          return _pat_let_tv19;
        } else {
          let _620___mcc_h3 = (_source11).cf6;
          let _621_cf6 = _620___mcc_h3;
          return _pat_let_tv20;
        }
      }(_613_v11);
      _597_v1 = _rhs100;
      r0 = _rhs101;
      r1 = _rhs102;
      _609_v7 = _rhs103;
      _610_v8 = _rhs104;
      let _622_v12;
      let _nw107 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _622_v12 = _nw107;
      _622_v12 = (_module.D3.create_DC10(p0, _622_v12, (_612_v10)[_module.__default.safeIndex(p0, new BigNumber((_612_v10).length))])).dtor_cf12;
      let _623_v13;
      _623_v13 = _dafny.Seq.of(_597_v1, p1, p1, _597_v1, _597_v1);
      let _624_v14;
      let _nw108 = new _module.C0();
      _nw108.__ctor(true, new BigNumber((_623_v13).length), (_this.f10) || (p2));
      _624_v14 = _nw108;
      let _625_v15;
      _625_v15 = new _dafny.CodePoint('k'.codePointAt(0));
      let _pat_let_tv21 = _624_v14;
      let _pat_let_tv22 = _624_v14;
      let _pat_let_tv23 = _624_v14;
      let _pat_let_tv24 = _624_v14;
      (_this).f10 = function (_source12) {
        if (_source12.is_DC10) {
          let _626___mcc_h4 = (_source12).cf11;
          let _627___mcc_h5 = (_source12).cf12;
          let _628___mcc_h6 = (_source12).cf13;
          let _629_cf13 = _628___mcc_h6;
          let _630_cf12 = _627___mcc_h5;
          let _631_cf11 = _626___mcc_h4;
          return true;
        } else if (_source12.is_DC11) {
          let _632___mcc_h7 = (_source12).cf14;
          let _633___mcc_h8 = (_source12).cf15;
          let _634___mcc_h9 = (_source12).cf16;
          let _635___mcc_h10 = (_source12).cf17;
          let _636_cf17 = _635___mcc_h10;
          let _637_cf16 = _634___mcc_h9;
          let _638_cf15 = _633___mcc_h8;
          let _639_cf14 = _632___mcc_h7;
          return (_pat_let_tv21).f13;
        } else if (_source12.is_DC9) {
          let _640___mcc_h11 = (_source12).cf10;
          let _641_cf10 = _640___mcc_h11;
          return (_this).f9;
        } else {
          let _642___mcc_h12 = (_source12).cf18;
          let _643_cf18 = _642___mcc_h12;
          return !((_dafny.Set.fromElements(_pat_let_tv22.f14)).IsProperSubsetOf(_dafny.Set.fromElements(_pat_let_tv23.f14, (_dafny.ZERO).minus(_pat_let_tv24.f14))));
        }
      }(_module.D3.create_DC9(_dafny.Seq.of(_625_v15, _625_v15)));
      let _644_v16;
      let _nw109 = Array((new BigNumber(15)).toNumber());
      _nw109[(_dafny.ZERO).toNumber()] = p1;
      _nw109[(_dafny.ONE).toNumber()] = p1;
      _nw109[(new BigNumber(2)).toNumber()] = _597_v1;
      _nw109[(new BigNumber(3)).toNumber()] = p1;
      _nw109[(new BigNumber(4)).toNumber()] = p1;
      _nw109[(new BigNumber(5)).toNumber()] = p1;
      _nw109[(new BigNumber(6)).toNumber()] = p1;
      _nw109[(new BigNumber(7)).toNumber()] = p1;
      _nw109[(new BigNumber(8)).toNumber()] = p1;
      _nw109[(new BigNumber(9)).toNumber()] = p1;
      _nw109[(new BigNumber(10)).toNumber()] = p1;
      _nw109[(new BigNumber(11)).toNumber()] = p1;
      _nw109[(new BigNumber(12)).toNumber()] = p1;
      _nw109[(new BigNumber(13)).toNumber()] = p1;
      _nw109[(new BigNumber(14)).toNumber()] = _597_v1;
      _644_v16 = _nw109;
      let _645_v17;
      _645_v17 = _module.D4.create_DC13(_644_v16);
      let _646_v18;
      let _nw110 = Array((new BigNumber(28)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = _644_v16;
      _nw110[(_dafny.ONE).toNumber()] = _644_v16;
      _nw110[(new BigNumber(2)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(3)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(4)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(5)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(6)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(7)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(8)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(9)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(10)).toNumber()] = (_645_v17).dtor_cf19;
      _nw110[(new BigNumber(11)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(12)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(13)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(14)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(15)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(16)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(17)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(18)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(19)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(20)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(21)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(22)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(23)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(24)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(25)).toNumber()] = _644_v16;
      _nw110[(new BigNumber(26)).toNumber()] = (_645_v17).dtor_cf19;
      _nw110[(new BigNumber(27)).toNumber()] = _644_v16;
      _646_v18 = _nw110;
      let _index110 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_646_v18).length));
      (_646_v18)[_index110] = _644_v16;
      let _index111 = _module.__default.safeIndex(new BigNumber(833), new BigNumber((_646_v18).length));
      (_646_v18)[_index111] = _644_v16;
      if ((new BigNumber((_dafny.Seq.Concat(_608_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), ((_647_v15) => function (_648_i3) {
        return _647_v15;
      })(_625_v15)))).length)).isLessThan(new BigNumber((_module.__default.fm10(globalState)).length))) {
        (globalState).f7 = _625_v15;
        let _649_v19;
        _649_v19 = _dafny.MultiSet.fromElements(p2);
        let _650_v20;
        _650_v20 = _dafny.Map.Empty.slice().updateUnsafe(_624_v14,_649_v19);
        let _651_v21;
        _651_v21 = _dafny.Map.Empty.slice().updateUnsafe(_650_v20,_module.__default.safeModuloInt(new BigNumber(66), new BigNumber(853)));
        _651_v21 = (_651_v21).update(_650_v20, _module.__default.safeDivisionInt(_624_v14.f14, _624_v14.f14));
        let _652_v22;
        let _init9 = ((_653_v2) => function (_654_i4) {
          return _653_v2;
        })(_604_v2);
        let _nw111 = Array((new BigNumber(12)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw111.length); _i0_9++) {
          _nw111[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _652_v22 = _nw111;
        _652_v22 = _652_v22;
        let _655_v23;
        let _nw112 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
        _655_v23 = _nw112;
        let _656_v24;
        _656_v24 = _dafny.MultiSet.fromElements(_597_v1);
        let _index112 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_655_v23).length));
        (_655_v23)[_index112] = (_656_v24).Difference(_dafny.MultiSet.fromElements(p1));
        let _index113 = _module.__default.safeIndex(new BigNumber(718), new BigNumber((_655_v23).length));
        (_655_v23)[_index113] = _656_v24;
        let _657_v25;
        _657_v25 = _module.D1.create_DC4(_606_v4);
        let _source13 = _657_v25;
        if (_source13.is_DC5) {
          (_624_v14).f14 = new BigNumber(421);
          let _658_v26;
          _658_v26 = _dafny.Seq.of(_604_v2);
          let _659_v27;
          _659_v27 = _dafny.Map.Empty.slice().updateUnsafe(_622_v12,_module.__default.fm3(_625_v15, true, p0, globalState));
          let _rhs105 = _624_v14.f14;
          let _rhs106 = _dafny.Seq.update(_module.__default.fm14(new BigNumber(-443), globalState), _module.__default.safeIndex((new BigNumber((_dafny.Seq.of((_624_v14).f13)).length)).minus(_624_v14.f14), new BigNumber((_module.__default.fm14(new BigNumber(-443), globalState)).length)), _604_v2);
          let _rhs107 = !(_659_v27).equals((_659_v27).Merge(_659_v27));
          let _rhs108 = _625_v15;
          let _rhs109 = p1;
          let _lhs70 = _624_v14;
          let _lhs71 = globalState;
          _lhs70.f14 = _rhs105;
          _658_v26 = _rhs106;
          r0 = _rhs107;
          _lhs71.f7 = _rhs108;
          _597_v1 = _rhs109;
          let _660_v28;
          let _nw113 = new _module.C1();
          _nw113.__ctor((_624_v14).f13, (_this).fm6((_624_v14).f13, (_this).f9, globalState), p2);
          _660_v28 = _nw113;
          _660_v28 = _660_v28;
          let _661_v29;
          let _nw114 = Array((new BigNumber(29)).toNumber()).fill(_module.D2.Default());
          _661_v29 = _nw114;
          let _index114 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_661_v29).length));
          (_661_v29)[_index114] = _613_v11;
          let _index115 = _module.__default.safeIndex(new BigNumber(523), new BigNumber((_661_v29).length));
          (_661_v29)[_index115] = _module.D2.create_DC8(false, (_660_v28).f12, (_624_v14).f13);
        } else if (_source13.is_DC6) {
          (_this).f10 = ((_this.f10) ? (!(p2)) : ((_624_v14).f13));
          let _662_v30;
          _662_v30 = _dafny.Seq.of(_608_v6, _608_v6);
          let _663_v31;
          _663_v31 = _dafny.Seq.of(_610_v8);
          let _664_v32;
          let _nw115 = Array((new BigNumber(9)).toNumber());
          _nw115[(_dafny.ZERO).toNumber()] = _610_v8;
          _nw115[(_dafny.ONE).toNumber()] = (_610_v8).update(_608_v6, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("gnmnf")).length)));
          _nw115[(new BigNumber(2)).toNumber()] = _610_v8;
          _nw115[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.FromArray(_662_v30)).Intersect(_610_v8);
          _nw115[(new BigNumber(4)).toNumber()] = (_610_v8).Union(_610_v8);
          _nw115[(new BigNumber(5)).toNumber()] = (_610_v8).Difference(_610_v8);
          _nw115[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm15((_this).f9, globalState));
          _nw115[(new BigNumber(7)).toNumber()] = (_dafny.MultiSet.fromElements(_608_v6)).Intersect((_663_v31)[_module.__default.safeIndex(p0, new BigNumber((_663_v31).length))]);
          _nw115[(new BigNumber(8)).toNumber()] = (_610_v8).Union(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("aonb")));
          _664_v32 = _nw115;
          let _index116 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_664_v32).length));
          (_664_v32)[_index116] = _610_v8;
          let _index117 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_597_v1).length));
          (_597_v1)[_index117] = (new BigNumber((_607_v5).length)).isLessThan((_dafny.ZERO).minus(_624_v14.f14));
          let _665_v33;
          _665_v33 = _dafny.Set.fromElements(_624_v14.f14, new BigNumber(141));
          let _index118 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_664_v32).length));
          let _index119 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_597_v1).length));
          let _rhs110 = _dafny.MultiSet.fromElements((_665_v33).IsDisjointFrom(_665_v33));
          let _rhs111 = !(p2);
          let _rhs112 = p0;
          let _rhs113 = _610_v8;
          let _rhs114 = (_624_v14).f13;
          let _lhs72 = globalState;
          let _lhs73 = _624_v14;
          let _lhs74 = _664_v32;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_664_v32).length));
          let _lhs76 = _597_v1;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_597_v1).length));
          _649_v19 = _rhs110;
          _lhs72.f0 = _rhs111;
          _lhs73.f14 = _rhs112;
          _lhs74[_lhs75] = _rhs113;
          _lhs76[_lhs77] = _rhs114;
          let _index120 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_622_v12).length));
          (_622_v12)[_index120] = p0;
          let _666_v34;
          _666_v34 = _dafny.MultiSet.fromElements(_605_v3);
          let _667_v35;
          let _nw116 = new _module.C1();
          _nw116.__ctor(_this.f10, _module.__default.fm3(_625_v15, (_624_v14).f13, new BigNumber(-346), globalState), (_666_v34).IsProperSubsetOf((_666_v34).update(_605_v3, _module.__default.abs(new BigNumber(-782)))));
          _667_v35 = _nw116;
          let _668_v36;
          _668_v36 = _dafny.Map.Empty.slice().updateUnsafe(_649_v19,_606_v4);
          let _index121 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_622_v12).length));
          let _rhs115 = new BigNumber(((((_668_v36).contains(_dafny.MultiSet.FromArray(((_this.f10) ? (_612_v10) : (_612_v10))))) ? ((_668_v36).get(_dafny.MultiSet.FromArray(((_this.f10) ? (_612_v10) : (_612_v10))))) : (_dafny.Seq.Concat((_657_v25).dtor_cf5, _606_v4)))).length);
          let _rhs116 = ((_624_v14).f13) && ((_597_v1)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_597_v1).length))]);
          let _rhs117 = new BigNumber(243);
          let _rhs118 = _667_v35;
          let _rhs119 = !((_624_v14.f14).isLessThan(_module.__default.safeDivisionInt(p0, new BigNumber(858))));
          let _lhs78 = _624_v14;
          let _lhs79 = _622_v12;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_622_v12).length));
          let _lhs81 = globalState;
          _lhs78.f14 = _rhs115;
          r0 = _rhs116;
          _lhs79[_lhs80] = _rhs117;
          _667_v35 = _rhs118;
          _lhs81.f0 = _rhs119;
          let _669_v37;
          _669_v37 = _dafny.Seq.of(_606_v4, _606_v4);
          let _670_v38;
          _670_v38 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f10) || (false),_dafny.Seq.Concat(_dafny.Seq.update(_669_v37, _module.__default.safeIndex(p0, new BigNumber((_669_v37).length)), _606_v4), _dafny.Seq.update(_dafny.Seq.update(_669_v37, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_669_v37).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_671_i5) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("yqssgr")).length);
          })), _module.__default.safeIndex(new BigNumber((_606_v4).length), new BigNumber((_dafny.Seq.update(_669_v37, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_669_v37).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(184)), function (_672_i5) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("yqssgr")).length);
          }))).length)), _606_v4)));
          _670_v38 = (_670_v38).update(!(((_597_v1)[_module.__default.safeIndex(new BigNumber(321), new BigNumber((_597_v1).length))]) && ((_667_v35).f9)), _669_v37);
        } else {
          let _673___mcc_h13 = (_source13).cf5;
          let _674_cf5 = _673___mcc_h13;
          let _675_v39;
          _675_v39 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.of((_624_v14).f13));
          (_624_v14).f14 = _module.__default.safeModuloInt((_624_v14.f14).minus(p0), (_this).fm4(_675_v39, (_624_v14).f13, _this.f10, _624_v14.f14, globalState));
          _597_v1 = _597_v1;
          let _676_v40;
          _676_v40 = _dafny.Set.fromElements(p1, _597_v1);
          _676_v40 = _676_v40;
          (_624_v14).f14 = (_dafny.ZERO).minus(p0);
        }
      } else {
        _607_v5 = (_607_v5).update(p0, (_624_v14).f13);
        let _677_v41;
        _677_v41 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),(_624_v14).f13));
        (globalState).f0 = ((!(false)) ? (true) : (_dafny.areEqual(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!((_624_v14).f13),(_612_v10)[_module.__default.safeIndex(p0, new BigNumber((_612_v10).length))]), _604_v2, _604_v2), _677_v41)));
        let _678_v42;
        let _nw117 = new _module.C0();
        _nw117.__ctor((_624_v14).f13, p0, true);
        _678_v42 = _nw117;
        let _679_v43;
        _679_v43 = _dafny.Map.Empty.slice().updateUnsafe(_624_v14.f14,_606_v4);
        let _680_v44;
        _680_v44 = _dafny.MultiSet.fromElements(_this.f10, (_624_v14).f13, _this.f10, !(_this.f10));
        _679_v43 = (_679_v43).update(((((_680_v44).contains((_624_v14).f13)) ? ((_680_v44).get((_624_v14).f13)) : ((_dafny.ZERO).minus(_624_v14.f14)))).multipliedBy(_module.__default.fm3(new _dafny.CodePoint('e'.codePointAt(0)), false, _624_v14.f14, globalState)), _dafny.Seq.Concat(_module.__default.fm13((_this).f9, (_678_v42).f13, (_624_v14).f13, new BigNumber((_608_v6).length), globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-594)), ((_681_v42) => function (_682_i6) {
          return _681_v42.f14;
        })(_678_v42))));
        (globalState).f7 = _625_v15;
      }
      r0 = (_624_v14).f13;
      let _683_v45;
      _683_v45 = _module.D0.create_DC0((_dafny.ZERO).minus(p0));
      let _pat_let_tv25 = _612_v10;
      let _pat_let_tv26 = _624_v14;
      let _pat_let_tv27 = _612_v10;
      let _pat_let_tv28 = _624_v14;
      let _pat_let_tv29 = p2;
      r1 = function (_source14) {
        if (_source14.is_DC1) {
          let _684___mcc_h14 = (_source14).cf1;
          let _685_cf1 = _684___mcc_h14;
          return (_pat_let_tv25)[_module.__default.safeIndex(_pat_let_tv26.f14, new BigNumber((_pat_let_tv27).length))];
        } else if (_source14.is_DC2) {
          let _686___mcc_h15 = (_source14).cf2;
          let _687___mcc_h16 = (_source14).cf3;
          let _688___mcc_h17 = (_source14).cf4;
          let _689_cf4 = _688___mcc_h17;
          let _690_cf3 = _687___mcc_h16;
          let _691_cf2 = _686___mcc_h15;
          return (_690_cf3).isLessThanOrEqualTo(_pat_let_tv28.f14);
        } else if (_source14.is_DC3) {
          return _this.f10;
        } else {
          let _692___mcc_h18 = (_source14).cf0;
          let _693_cf0 = _692___mcc_h18;
          return _pat_let_tv29;
        }
      }(_683_v45);
      let _694_v46;
      _694_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,_622_v12);
      r2 = (_694_v46).update((_dafny.ZERO).minus(_624_v14.f14), _622_v12);
      return [r0, r1, r2];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f9) {
      let _this = this;
      (_this)._f9 = f9;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ouhfhsh")).length), new BigNumber(-999))).cardinality()))).minus((new BigNumber(636)).multipliedBy(new BigNumber(-559)));
    };
    fm5(globalState) {
      let _this = this;
      return new _dafny.CodePoint('a'.codePointAt(0));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D0.Default();
      let r2 = _dafny.ZERO;
      let _695_v0;
      _695_v0 = new BigNumber(-2);
      let _696_v1;
      _696_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_695_v0).minus(new BigNumber(757)));
      let _697_v2;
      let _init10 = ((_698_v0) => function (_699_i0) {
        return (_699_i0).minus(_698_v0);
      })(_695_v0);
      let _nw118 = Array((new BigNumber(20)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw118.length); _i0_10++) {
        _nw118[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _697_v2 = _nw118;
      let _700_v3;
      _700_v3 = _dafny.Set.fromElements(_697_v2, _697_v2, _697_v2, _697_v2);
      _696_v1 = (_696_v1).update((_700_v3).IsProperSubsetOf(_dafny.Set.fromElements(_697_v2, _697_v2, _697_v2)), _695_v0);
      let _701_v4;
      _701_v4 = _dafny.Seq.of(_695_v0, _695_v0, _695_v0, _695_v0, _695_v0);
      let _702_v5;
      _702_v5 = _dafny.Seq.UnicodeFromString("ps");
      let _703_v6;
      _703_v6 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_695_v0), _695_v0);
      let _704_v7;
      _704_v7 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_695_v0), _dafny.MultiSet.fromElements(new BigNumber((_702_v5).length), _695_v0), _703_v6);
      let _705_v8;
      _705_v8 = _dafny.Map.Empty.slice().updateUnsafe(_701_v4,(_module.D0.create_DC2((_this).f9, _695_v0, _704_v7)).dtor_cf3);
      let _706_v9;
      _706_v9 = _dafny.MultiSet.fromElements(new BigNumber((_705_v8).length));
      let _707_v10;
      let _nw119 = Array((new BigNumber(6)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = _703_v6;
      _nw119[(_dafny.ONE).toNumber()] = _706_v9;
      _nw119[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(_695_v0);
      _nw119[(new BigNumber(3)).toNumber()] = _703_v6;
      _nw119[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.FromArray(_701_v4);
      _nw119[(new BigNumber(5)).toNumber()] = (_706_v9).Difference(_703_v6);
      _707_v10 = _nw119;
      let _index122 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_707_v10).length));
      (_707_v10)[_index122] = (((_this).f9) ? (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_695_v0), _695_v0)) : (_703_v6));
      let _708_v11;
      _708_v11 = _dafny.Map.Empty.slice().updateUnsafe(_695_v0,_706_v9);
      let _index123 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_707_v10).length));
      (_707_v10)[_index123] = (_706_v9).Intersect((((_708_v11).contains(_695_v0)) ? ((_708_v11).get(_695_v0)) : (_706_v9)));
      let _709_i1;
      _709_i1 = _dafny.ZERO;
      L2: {
        while ((_this).f9) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_709_i1)) {
              break L2;
            }
            _709_i1 = (_709_i1).plus(_dafny.ONE);
            (globalState).f3 = _module.__default.fm2(globalState);
            let _710_v12;
            let _nw120 = new _module.C1();
            _nw120.__ctor((_this).f9, _695_v0, (_this).f9);
            _710_v12 = _nw120;
            let _711_v13;
            _711_v13 = new _dafny.CodePoint('g'.codePointAt(0));
            let _712_v14;
            let _nw121 = new _module.C0();
            _nw121.__ctor((_this).f9, _695_v0, (_this).f9);
            _712_v14 = _nw121;
            let _713_v15;
            _713_v15 = _dafny.Map.Empty.slice().updateUnsafe(_712_v14,(_710_v12).f11);
            let _714_v16;
            let _nw122 = new _module.C1();
            _nw122.__ctor((_dafny.MultiSet.fromElements(_695_v0)).contains(_module.__default.fm3(_711_v13, (_this).f9, (_710_v12).f12, globalState)), new BigNumber((_dafny.MultiSet.fromElements(_695_v0, (_710_v12).f12)).cardinality()), (((_713_v15).contains(_712_v14)) ? ((_713_v15).get(_712_v14)) : ((_this).f9)));
            _714_v16 = _nw122;
            let _715_v17;
            _715_v17 = _dafny.Map.Empty.slice().updateUnsafe((_710_v12).f12,_710_v12);
            let _716_v18;
            _716_v18 = _dafny.Seq.of(_710_v12, _714_v16, _710_v12);
            _714_v16 = (((_715_v17).contains((_714_v16).f12)) ? ((_715_v17).get((_714_v16).f12)) : ((_716_v18)[_module.__default.safeIndex((_710_v12).f12, new BigNumber((_716_v18).length))]));
            let _717_v19;
            _717_v19 = _dafny.MultiSet.fromElements(!((_710_v12).f11), (_this).f9, (_this).f9, (_712_v14).f9);
            let _718_v20;
            let _nw123 = new _module.C2();
            _nw123.__ctor((_717_v19).IsDisjointFrom(_717_v19), !_dafny.Seq.contains(_702_v5, _711_v13));
            _718_v20 = _nw123;
          }
        }
      }
      let _hi10 = _695_v0;
      for (let _719_i2 = _695_v0; _719_i2.isLessThan(_hi10); _719_i2 = _719_i2.plus(_dafny.ONE)) {
        r0 = new BigNumber((_dafny.Seq.Concat(_702_v5, _702_v5)).length);
        let _720_v21;
        _720_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_702_v5);
        _720_v21 = (_720_v21).update(!(_module.__default.fm2(globalState)) || ((_this).f9), _dafny.Seq.Concat(_702_v5, _702_v5));
        let _721_v22;
        _721_v22 = _dafny.Map.Empty.slice().updateUnsafe(_702_v5,!(false));
        let _722_v23;
        _722_v23 = _dafny.Set.fromElements((_this).f9, (_this).f9, (_this).f9, (_this).f9, (_this).f9);
        let _723_v24;
        _723_v24 = _dafny.Map.Empty.slice().updateUnsafe(_721_v22,(_722_v23).Intersect(_dafny.Set.fromElements((_this).f9, (_this).f9, (_this).f9, (_this).f9, !((_this).f9))));
        _723_v24 = (_723_v24).update(_dafny.Map.Empty.slice().updateUnsafe(_702_v5,(_this).f9), (_722_v23).Union(_722_v23));
        (globalState).f3 = (_this).f9;
      }
      let _724_v25;
      _724_v25 = new _dafny.CodePoint('t'.codePointAt(0));
      let _725_v26;
      _725_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_724_v25);
      let _726_v27;
      _726_v27 = _dafny.Map.Empty.slice().updateUnsafe(_695_v0,new BigNumber((_725_v26).length));
      _726_v27 = (_726_v27).update(_695_v0, _module.__default.safeDivisionInt(new BigNumber((_703_v6).cardinality()), _695_v0));
      if ((_this).f9) {
        let _727_v28;
        let _nw124 = Array((new BigNumber(3)).toNumber());
        _nw124[(_dafny.ZERO).toNumber()] = _697_v2;
        _nw124[(_dafny.ONE).toNumber()] = _697_v2;
        _nw124[(new BigNumber(2)).toNumber()] = _697_v2;
        _727_v28 = _nw124;
        _727_v28 = _727_v28;
        let _728_v29;
        _728_v29 = _dafny.Seq.of(_696_v1);
        let _729_v30;
        _729_v30 = _module.D2.create_DC8((_this).f9, (((_this).f9) ? (_695_v0) : (new BigNumber((_726_v27).length))), _dafny.Seq.IsProperPrefixOf(_728_v29, _728_v29));
        let _source15 = _729_v30;
        if (_source15.is_DC8) {
          let _730___mcc_h0 = (_source15).cf7;
          let _731___mcc_h1 = (_source15).cf8;
          let _732___mcc_h2 = (_source15).cf9;
          let _733_cf9 = _732___mcc_h2;
          let _734_cf8 = _731___mcc_h1;
          let _735_cf7 = _730___mcc_h0;
          (globalState).f3 = (_this).f9;
          let _736_v31;
          _736_v31 = _dafny.MultiSet.fromElements(!(true), _733_cf9);
          (globalState).f3 = (_module.__default.safeDivisionInt(_695_v0, new BigNumber((_736_v31).cardinality()))).isLessThanOrEqualTo((_dafny.ZERO).minus((_734_cf8).multipliedBy(_695_v0)));
          (globalState).f0 = (_this).f9;
          let _737_v32;
          _737_v32 = _dafny.Map.Empty.slice().updateUnsafe(_724_v25,_dafny.Seq.UnicodeFromString("uathenfi"));
          let _738_v33;
          let _out19;
          _out19 = _module.__default.m0(((false) ? (_734_cf8) : (new BigNumber((_701_v4).length))), ((_733_cf9) ? (_697_v2) : (_697_v2)), _735_cf7, _module.__default.fm16(_737_v32, _734_cf8, globalState), globalState);
          _738_v33 = _out19;
        } else {
          let _739___mcc_h3 = (_source15).cf6;
          let _740_cf6 = _739___mcc_h3;
          let _741_v34;
          _741_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_695_v0),(_this).f9);
          let _742_v35;
          let _out20;
          _out20 = _module.__default.m0(_module.__default.fm3(_724_v25, (_this).f9, _695_v0, globalState), _697_v2, false, _741_v34, globalState);
          _742_v35 = _out20;
          let _743_v36;
          let _nw125 = new _module.C0();
          _nw125.__ctor(!_dafny.Seq.contains(_702_v5, _724_v25), _695_v0, ((_this).f9) || (!(false)));
          _743_v36 = _nw125;
          let _744_v37;
          _744_v37 = _dafny.Map.Empty.slice().updateUnsafe(_702_v5,_695_v0);
          let _745_v38;
          _745_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(970),_697_v2)).length),_702_v5);
          let _rhs120 = _743_v36;
          let _rhs121 = !(_745_v38).equals(_745_v38);
          let _rhs122 = new BigNumber((_701_v4).length);
          let _rhs123 = _744_v37;
          let _lhs82 = globalState;
          _743_v36 = _rhs120;
          _lhs82.f3 = _rhs121;
          r2 = _rhs122;
          _744_v37 = _rhs123;
          r2 = _module.__default.safeModuloInt(_695_v0, (new BigNumber((_701_v4).length)).plus(_695_v0));
          _702_v5 = _dafny.Seq.UnicodeFromString("kmguhhp");
        }
        let _746_v39;
        let _nw126 = new _module.C0();
        _nw126.__ctor((_695_v0).isLessThan(new BigNumber(-788)), _695_v0, (_this).f9);
        _746_v39 = _nw126;
        _746_v39 = _746_v39;
        let _747_v40;
        _747_v40 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f9);
        (globalState).f3 = (((_747_v40).contains(!((_746_v39).f13))) ? ((_747_v40).get(!((_746_v39).f13))) : (false));
        let _748_v41;
        _748_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_707_v10)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_707_v10).length))]).cardinality()),(((_746_v39).f13) ? ((_this).fm5(globalState)) : (_724_v25)));
        _748_v41 = (_748_v41).update(_695_v0, _724_v25);
      } else {
        let _749_v42;
        _749_v42 = _dafny.Map.Empty.slice().updateUnsafe(_724_v25,_module.__default.safeModuloInt(_695_v0, _695_v0));
        _749_v42 = _749_v42;
        _702_v5 = _702_v5;
        _695_v0 = (new BigNumber(150)).multipliedBy(_695_v0);
        r2 = _695_v0;
        let _750_v43;
        let _nw127 = Array((new BigNumber(18)).toNumber()).fill([]);
        _750_v43 = _nw127;
        let _751_v44;
        let _init11 = ((_752_v25) => function (_753_i3) {
          return _752_v25;
        })(_724_v25);
        let _nw128 = Array((new BigNumber(8)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw128.length); _i0_11++) {
          _nw128[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _751_v44 = _nw128;
        let _index124 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_750_v43).length));
        (_750_v43)[_index124] = _751_v44;
        let _index125 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_750_v43).length));
        (_750_v43)[_index125] = _751_v44;
      }
      r0 = ((_dafny.ZERO).minus(_695_v0)).plus(_module.__default.safeDivisionInt(_695_v0, _695_v0));
      let _754_v45;
      _754_v45 = _module.D1.create_DC5();
      let _pat_let_tv30 = _695_v0;
      let _pat_let_tv31 = _695_v0;
      let _pat_let_tv32 = _726_v27;
      let _pat_let_tv33 = _695_v0;
      let _pat_let_tv34 = _695_v0;
      let _pat_let_tv35 = _726_v27;
      let _pat_let_tv36 = _695_v0;
      r1 = function (_source16) {
        if (_source16.is_DC5) {
          return _module.D0.create_DC0(_pat_let_tv30);
        } else if (_source16.is_DC6) {
          return _module.D0.create_DC0(_pat_let_tv31);
        } else {
          let _755___mcc_h4 = (_source16).cf5;
          let _756_cf5 = _755___mcc_h4;
          return _module.D0.create_DC0(new BigNumber((_dafny.Set.fromElements((((_pat_let_tv35).contains(_pat_let_tv36)) ? ((_pat_let_tv32).get(_pat_let_tv33)) : (_pat_let_tv34)))).length));
        }
      }(_754_v45);
      let _757_v46;
      _757_v46 = _dafny.Set.fromElements((_this).f9, true, true);
      r2 = new BigNumber(((_757_v46).Intersect(_757_v46)).length);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
