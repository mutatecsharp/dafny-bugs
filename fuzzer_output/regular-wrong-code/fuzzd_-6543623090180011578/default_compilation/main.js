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
      return !((_dafny.MultiSet.fromElements(true)).IsDisjointFrom(_dafny.MultiSet.fromElements(false, false, !(false), false))) || ((false) || (true));
    };
    static fm3(p0, globalState) {
      if (true) {
        return _dafny.Seq.UnicodeFromString("hgwprn");
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), function (_0_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("tsxqnnw"));
      }
    };
    static fm4(p0, p1, globalState) {
      return (_dafny.ZERO).minus((new BigNumber((((!(true)) ? (_dafny.Seq.UnicodeFromString("fv")) : (_dafny.Seq.UnicodeFromString("uosyjyo")))).length)).multipliedBy((_module.D4.create_DC15(!(false), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).dtor_cf24));
    };
    static fm5(globalState) {
      return ((_dafny.Set.fromElements(!(false), false, false)).Union(_dafny.Set.fromElements(true))).Difference((_dafny.Set.fromElements(true, true)).Intersect(_dafny.Set.fromElements(false, !(true))));
    };
    static fm6(globalState) {
      return _module.D2.create_DC7((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)).minus(new BigNumber(242)));
    };
    static fm7(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC3(false, false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), function (_1_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
}), false, new BigNumber(45))).dtor_cf5,new BigNumber(7))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(177)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-120))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(59), new BigNumber(-991))) {
          let _2_v0 = _compr_0;
          if (((new BigNumber(59)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(-991)))) {
            _coll0.push([_module.__default.safeDivisionInt(_2_v0, new BigNumber(958)),(true) || (true)]);
          }
        }
        return _coll0;
      }();
    };
    static fm9(p0, p1, globalState) {
      return ((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(500), new BigNumber(122))) {
          let _3_v0 = _compr_1;
          if (((new BigNumber(500)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(122)))) {
            _coll1.add((_3_v0).multipliedBy(new BigNumber(-706)));
          }
        }
        return _coll1;
      }()).Difference(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-473), new BigNumber(223))) {
          let _4_v1 = _compr_2;
          if (((new BigNumber(-473)).isLessThanOrEqualTo(_4_v1)) && ((_4_v1).isLessThan(new BigNumber(223)))) {
            _coll2.add(_module.__default.safeModuloInt(_4_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(461),true)).length)));
          }
        }
        return _coll2;
      }())).Union(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-290))));
    };
    static fm10(p0, p1, globalState) {
      if (false) {
        if (false) {
          return _dafny.Seq.of(false, !(!(false)));
        } else {
          return _dafny.Seq.of(false);
        }
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(false, false, (_module.D6.create_DC20(true)).dtor_cf34), _dafny.Seq.of(true));
      }
    };
    static fm11(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC12(_module.D3.create_DC10(false, new BigNumber(524), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(380))).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC11()))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC12(_module.D3.create_DC10(false, new BigNumber(682), new BigNumber((function () {
  let _coll3 = new _dafny.Set();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(337), new BigNumber(918))) {
    let _5_v0 = _compr_3;
    if (((new BigNumber(337)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(918)))) {
      _coll3.add(_module.__default.safeDivisionInt(_5_v0, new BigNumber(99)));
    }
  }
  return _coll3;
}()).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC12(_module.D3.create_DC9(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).cardinality())))))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      let _source0 = ((true) ? (_module.D1.create_DC2(new BigNumber(376))) : (_module.D1.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("arg")).length))));
      if (_source0.is_DC3) {
        let _6___mcc_h0 = (_source0).cf5;
        let _7___mcc_h1 = (_source0).cf6;
        let _8___mcc_h2 = (_source0).cf7;
        let _9___mcc_h3 = (_source0).cf8;
        let _10___mcc_h4 = (_source0).cf9;
        let _11_cf9 = _10___mcc_h4;
        let _12_cf8 = _9___mcc_h3;
        let _13_cf7 = _8___mcc_h2;
        let _14_cf6 = _7___mcc_h1;
        let _15_cf5 = _6___mcc_h0;
        return _module.D3.create_DC12(_module.D3.create_DC10(_12_cf8, _11_cf9, _11_cf9));
      } else if (_source0.is_DC2) {
        let _16___mcc_h5 = (_source0).cf4;
        let _17_cf4 = _16___mcc_h5;
        return _module.D3.create_DC12(_module.D3.create_DC11());
      } else {
        let _18___mcc_h6 = (_source0).cf10;
        let _19_cf10 = _18___mcc_h6;
        return _module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC12(_module.D3.create_DC9(_dafny.Seq.of(new BigNumber(27))))));
      }
    };
    static fm13(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(254),_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("nk")).length), new BigNumber((_dafny.Set.fromElements(false, true)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(240),function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(937), new BigNumber(842))) {
          let _20_v0 = _compr_4;
          if (((new BigNumber(937)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(842)))) {
            _coll4.add(_module.__default.safeModuloInt(_20_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))));
          }
        }
        return _coll4;
      }()));
    };
    static fm14(globalState) {
      return _module.D1.create_DC4(_module.D1.create_DC4(_module.D1.create_DC3(true, false, _dafny.Seq.UnicodeFromString("eewew"), true, new BigNumber(243))));
    };
    static fm15(p0, p1, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false, false, false, false))).Intersect((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false)));
    };
    static fm16(p0, p1, globalState) {
      let _source1 = _module.D6.create_DC20(true);
      if (_source1.is_DC19) {
        let _21___mcc_h0 = (_source1).cf31;
        let _22___mcc_h1 = (_source1).cf32;
        let _23___mcc_h2 = (_source1).cf33;
        let _24_cf33 = _23___mcc_h2;
        let _25_cf32 = _22___mcc_h1;
        let _26_cf31 = _21___mcc_h0;
        return _module.D4.create_DC15(_25_cf32, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_module.D0.create_DC0(_25_cf32)).dtor_cf0)).length));
      } else if (_source1.is_DC20) {
        let _27___mcc_h3 = (_source1).cf34;
        let _28_cf34 = _27___mcc_h3;
        return _module.D4.create_DC15(_28_cf34, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_28_cf34)).length), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).cardinality()))).length)),_dafny.MultiSet.fromElements(new BigNumber(673), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_28_cf34,true)).length)))).length)));
      } else if (_source1.is_DC21) {
        return _module.D4.create_DC15(true, new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of (function () {
    let _coll6 = new _dafny.Set();
    for (const _compr_6 of _dafny.IntegerRange(new BigNumber(572), new BigNumber(240))) {
      let _29_v1 = _compr_6;
      if (((new BigNumber(572)).isLessThanOrEqualTo(_29_v1)) && ((_29_v1).isLessThan(new BigNumber(240)))) {
        _coll6.add((_29_v1).multipliedBy(new BigNumber(414)));
      }
    }
    return _coll6;
  }()).Elements) {
    let _30_v0 = _compr_5;
    if ((function () {
      let _coll7 = new _dafny.Set();
      for (const _compr_7 of _dafny.IntegerRange(new BigNumber(572), new BigNumber(240))) {
        let _31_v1 = _compr_7;
        if (((new BigNumber(572)).isLessThanOrEqualTo(_31_v1)) && ((_31_v1).isLessThan(new BigNumber(240)))) {
          _coll7.add((_31_v1).multipliedBy(new BigNumber(414)));
        }
      }
      return _coll7;
    }()).contains(_30_v0)) {
      _coll5.push([_module.__default.safeModuloInt(_30_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(691))).length))),true]);
    }
  }
  return _coll5;
}()).length))).length)))).length));
      } else {
        let _32___mcc_h4 = (_source1).cf30;
        let _33_cf30 = _32___mcc_h4;
        if (true) {
          return _module.D4.create_DC15(!(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(592))).length));
        } else {
          return _module.D4.create_DC15(true, new BigNumber((_dafny.Seq.UnicodeFromString("ngymo")).length));
        }
      }
    };
    static fm17(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).length)))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(943))).length)), _dafny.Set.fromElements(new BigNumber(-890), new BigNumber((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Seq.of(new BigNumber(445), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(873), new BigNumber(28))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("jfd")).length), new BigNumber((_dafny.Seq.UnicodeFromString("cngsgffno")).length), new BigNumber((_dafny.Seq.UnicodeFromString("giiveerwc")).length))).Elements) {
          let _34_v0 = _compr_8;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(445), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(873), new BigNumber(28))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("jfd")).length), new BigNumber((_dafny.Seq.UnicodeFromString("cngsgffno")).length), new BigNumber((_dafny.Seq.UnicodeFromString("giiveerwc")).length)), _34_v0)) {
            _coll8.push([(_34_v0).plus((_module.D4.create_DC15(false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), function (_35_i0) {
  return _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()));
})).length))).dtor_cf24),new BigNumber(-492)]);
          }
        }
        return _coll8;
      }()).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(868))).length)), _dafny.Set.fromElements(new BigNumber(472)), _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_module.D4.create_DC15(false, new BigNumber(655))).dtor_cf24)).length), new BigNumber(-685), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(233), new BigNumber(864))) {
          let _36_v1 = _compr_9;
          if (((new BigNumber(233)).isLessThanOrEqualTo(_36_v1)) && ((_36_v1).isLessThan(new BigNumber(864)))) {
            _coll9.add(_module.__default.safeModuloInt(_36_v1, new BigNumber(830)));
          }
        }
        return _coll9;
      }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)))).length)),true)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))));
    };
    static fm18(p0, globalState) {
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_37_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-673)), function (_38_i1) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        });
      });
    };
    static m0(p0, globalState) {
      let r0 = false;
      r0 = p0;
      let _39_v0;
      _39_v0 = _module.D0.create_DC0(p0);
      let _40_i0;
      _40_i0 = _dafny.ZERO;
      L0: {
        let _pat_let_tv3 = p0;
        let _pat_let_tv4 = p0;
        while (function (_source3) {
          if (_source3.is_DC0) {
            let _99___mcc_h6 = (_source3).cf0;
            let _100_cf0 = _99___mcc_h6;
            return _pat_let_tv3;
          } else {
            let _101___mcc_h7 = (_source3).cf1;
            let _102___mcc_h8 = (_source3).cf2;
            let _103___mcc_h9 = (_source3).cf3;
            let _104_cf3 = _103___mcc_h9;
            let _105_cf2 = _102___mcc_h8;
            let _106_cf1 = _101___mcc_h7;
            return ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-575)), function (_107_i1) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            })).length))).isEqualTo((_module.D1.create_DC3(_pat_let_tv4, true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(431)), function (_108_i2) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}), false, (_dafny.ZERO).minus(new BigNumber(-332)))).dtor_cf9);
          }
        }(_39_v0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_40_i0)) {
              break L0;
            }
            _40_i0 = (_40_i0).plus(_dafny.ONE);
            r0 = p0;
            let _41_v1;
            _41_v1 = new BigNumber(660);
            let _42_v2;
            _42_v2 = _dafny.Seq.of(_41_v1, _41_v1);
            let _43_v3;
            let _nw0 = Array((new BigNumber(3)).toNumber());
            _nw0[(_dafny.ZERO).toNumber()] = _41_v1;
            _nw0[(_dafny.ONE).toNumber()] = _41_v1;
            _nw0[(new BigNumber(2)).toNumber()] = new BigNumber((_42_v2).length);
            _43_v3 = _nw0;
            let _44_v4;
            _44_v4 = _module.D2.create_DC6(_41_v1, true, _43_v3);
            let _pat_let_tv0 = p0;
            let _pat_let_tv1 = p0;
            let _pat_let_tv2 = _43_v3;
            let _source2 = function (_pat_let0_0) {
              return function (_45_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_46_dt__update_hcf13_h0) {
                    return function (_pat_let2_0) {
                      return function (_47_dt__update_hcf14_h0) {
                        return _module.D2.create_DC6((_45_dt__update__tmp_h0).dtor_cf12, _46_dt__update_hcf13_h0, _47_dt__update_hcf14_h0);
                      }(_pat_let2_0);
                    }(_pat_let_tv2);
                  }(_pat_let1_0);
                }((_pat_let_tv0) === (_pat_let_tv1));
              }(_pat_let0_0);
            }(_44_v4);
            if (_source2.is_DC6) {
              let _48___mcc_h0 = (_source2).cf12;
              let _49___mcc_h1 = (_source2).cf13;
              let _50___mcc_h2 = (_source2).cf14;
              let _51_cf14 = _50___mcc_h2;
              let _52_cf13 = _49___mcc_h1;
              let _53_cf12 = _48___mcc_h0;
              let _54_v5;
              _54_v5 = _dafny.Set.fromElements(false);
              let _55_v6;
              let _nw1 = new _module.C0();
              _nw1.__ctor(_53_cf12, ((p0) ? (_dafny.Set.fromElements(!(p0), _module.__default.fm0(!(p0), p0, _53_cf12, globalState), _52_cf13)) : (_54_v5)));
              _55_v6 = _nw1;
              let _56_v7;
              _56_v7 = _dafny.Seq.of(p0, p0, (_dafny.Set.fromElements(_52_cf13)).contains(p0));
              r0 = (_56_v7)[_module.__default.safeIndex(_53_cf12, new BigNumber((_56_v7).length))];
              let _57_v8;
              let _nw2 = new _module.C0();
              _nw2.__ctor((_55_v6).f6, _54_v5);
              _57_v8 = _nw2;
              let _58_v9;
              _58_v9 = _dafny.Set.fromElements((_55_v6).f6, _41_v1);
              let _59_v10;
              _59_v10 = _dafny.Seq.of(_58_v9);
              let _60_v11;
              _60_v11 = _dafny.Map.Empty.slice().updateUnsafe((_57_v8).f6,_dafny.Seq.UnicodeFromString("upqhkypa"));
              let _61_v12;
              _61_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_60_v11).length),_58_v9);
              let _62_v15;
              _62_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_57_v8).f6);
              let _63_v16;
              let _nw3 = Array((new BigNumber(15)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_59_v10, _dafny.Seq.of((((_61_v12).contains((_55_v6).f6)) ? ((_61_v12).get((_55_v6).f6)) : ((_59_v10)[_module.__default.safeIndex(_53_cf12, new BigNumber((_59_v10).length))]))));
              _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_64_cf12) => function (_65_i3) {
                return function () {
                  let _coll10 = new _dafny.Set();
                  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(239), new BigNumber(328))) {
                    let _66_v13 = _compr_10;
                    if (((new BigNumber(239)).isLessThanOrEqualTo(_66_v13)) && ((_66_v13).isLessThan(new BigNumber(328)))) {
                      _coll10.add(_module.__default.safeModuloInt(_66_v13, _64_cf12));
                    }
                  }
                  return _coll10;
                }();
              })(_53_cf12));
              _nw3[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_59_v10, _module.__default.safeIndex((_57_v8).f6, new BigNumber((_59_v10).length)), _58_v9), _59_v10);
              _nw3[(new BigNumber(3)).toNumber()] = _59_v10;
              _nw3[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_59_v10, _59_v10);
              _nw3[(new BigNumber(5)).toNumber()] = _59_v10;
              _nw3[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), ((_67_v9) => function (_68_i4) {
                return _67_v9;
              })(_58_v9));
              _nw3[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_58_v9), _dafny.Seq.of(_58_v9, _58_v9));
              _nw3[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_59_v10, _module.__default.safeIndex((_55_v6).f6, new BigNumber((_59_v10).length)), function () {
                let _coll11 = new _dafny.Set();
                for (const _compr_11 of _dafny.IntegerRange(new BigNumber(-118), new BigNumber(257))) {
                  let _69_v14 = _compr_11;
                  if (((new BigNumber(-118)).isLessThanOrEqualTo(_69_v14)) && ((_69_v14).isLessThan(new BigNumber(257)))) {
                    _coll11.add(_module.__default.safeDivisionInt(_69_v14, new BigNumber((_56_v7).length)));
                  }
                }
                return _coll11;
              }());
              _nw3[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_59_v10, _59_v10);
              _nw3[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_59_v10, _module.__default.safeIndex(new BigNumber((_62_v15).length), new BigNumber((_59_v10).length)), _58_v9);
              _nw3[(new BigNumber(11)).toNumber()] = _59_v10;
              _nw3[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_59_v10, _59_v10);
              _nw3[(new BigNumber(13)).toNumber()] = _59_v10;
              _nw3[(new BigNumber(14)).toNumber()] = _59_v10;
              _63_v16 = _nw3;
              let _index0 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_63_v16).length));
              (_63_v16)[_index0] = _59_v10;
              let _70_v17;
              _70_v17 = _dafny.MultiSet.fromElements(_53_cf12);
              let _71_v18;
              let _nw4 = new _module.C0();
              _nw4.__ctor(_module.__default.safeDivisionInt((_57_v8).f6, (((_62_v15).contains(p0)) ? ((_62_v15).get(p0)) : (new BigNumber((_70_v17).cardinality())))), (_57_v8).f7);
              _71_v18 = _nw4;
              let _index1 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_63_v16).length));
              let _rhs0 = _59_v10;
              let _rhs1 = _71_v18;
              let _rhs2 = p0;
              let _lhs0 = _63_v16;
              let _lhs1 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_63_v16).length));
              _lhs0[_lhs1] = _rhs0;
              _71_v18 = _rhs1;
              r0 = _rhs2;
            } else if (_source2.is_DC7) {
              let _72___mcc_h3 = (_source2).cf15;
              let _73_cf15 = _72___mcc_h3;
              let _74_v19;
              _74_v19 = _dafny.Set.fromElements(new BigNumber(602));
              r0 = ((p0) ? (!(p0)) : ((_74_v19).IsProperSubsetOf(_74_v19)));
              r0 = _module.__default.fm0(true, p0, _module.__default.safeDivisionInt(_73_cf15, _73_cf15), globalState);
              let _75_v20;
              _75_v20 = _dafny.Set.fromElements(_74_v19);
              let _76_v21;
              _76_v21 = _dafny.Map.Empty.slice().updateUnsafe(((p0) ? (new BigNumber((_75_v20).length)) : ((_dafny.ZERO).minus(_73_cf15))),_41_v1);
              _76_v21 = _76_v21;
              let _77_v22;
              _77_v22 = _dafny.Set.fromElements(false, p0, p0);
              let _78_v23;
              let _nw5 = new _module.C0();
              _nw5.__ctor((_dafny.ZERO).minus(_73_cf15), _77_v22);
              _78_v23 = _nw5;
            } else if (_source2.is_DC5) {
              let _79___mcc_h4 = (_source2).cf11;
              let _80_cf11 = _79___mcc_h4;
              let _81_v24;
              _81_v24 = new _dafny.CodePoint('w'.codePointAt(0));
              let _82_v25;
              _82_v25 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_81_v24);
              _82_v25 = (_82_v25).update(p0, ((false) ? (_81_v24) : (_81_v24)));
              let _83_v26;
              let _nw6 = new _module.C0();
              _nw6.__ctor(_module.__default.safeModuloInt(_41_v1, _41_v1), _dafny.Set.fromElements(p0, p0));
              _83_v26 = _nw6;
              _83_v26 = _83_v26;
              let _84_v27;
              _84_v27 = _dafny.Set.fromElements((_83_v26).f6, _41_v1, new BigNumber(-816), _41_v1);
              let _85_v28;
              let _nw7 = new _module.C0();
              _nw7.__ctor(new BigNumber((_84_v27).length), (_83_v26).f7);
              _85_v28 = _nw7;
              (globalState).f0 = (_41_v1).minus((_83_v26).f6);
            } else {
              let _86___mcc_h5 = (_source2).cf16;
              let _87_cf16 = _86___mcc_h5;
              _41_v1 = new BigNumber((function () {
                let _coll12 = new _dafny.Map();
                for (const _compr_12 of _dafny.IntegerRange(new BigNumber(933), new BigNumber(887))) {
                  let _88_v29 = _compr_12;
                  if (((new BigNumber(933)).isLessThanOrEqualTo(_88_v29)) && ((_88_v29).isLessThan(new BigNumber(887)))) {
                    _coll12.push([_module.__default.safeDivisionInt(_88_v29, new BigNumber((_dafny.Seq.of(!(p0))).length)),p0]);
                  }
                }
                return _coll12;
              }()).length);
              let _89_v31;
              _89_v31 = _dafny.Map.Empty.slice().updateUnsafe(_41_v1,_41_v1);
              let _90_v32;
              _90_v32 = _dafny.Map.Empty.slice().updateUnsafe(_89_v31,_41_v1);
              let _index2 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_43_v3).length));
              (_43_v3)[_index2] = ((p0) ? (_41_v1) : (new BigNumber(((function () {
                let _coll13 = new _dafny.Map();
                for (const _compr_13 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p0, p0)).cardinality()))).Elements) {
                  let _91_v30 = _compr_13;
                  if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p0, p0)).cardinality()))).contains(_91_v30)) {
                    _coll13.push([(_91_v30).multipliedBy(new BigNumber(877)),p0]);
                  }
                }
                return _coll13;
              }()).update(new BigNumber((_90_v32).length), p0)).length)));
              let _index3 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_43_v3).length));
              (_43_v3)[_index3] = _41_v1;
              let _92_v33;
              let _nw8 = Array((new BigNumber(7)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _92_v33 = _nw8;
              _92_v33 = _92_v33;
              let _index4 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_43_v3).length));
              (_43_v3)[_index4] = _module.__default.safeModuloInt((_43_v3)[_module.__default.safeIndex(new BigNumber(142), new BigNumber((_43_v3).length))], (_43_v3)[_module.__default.safeIndex(new BigNumber(142), new BigNumber((_43_v3).length))]);
            }
            let _93_v34;
            let _nw9 = Array((new BigNumber(4)).toNumber()).fill(false);
            _93_v34 = _nw9;
            let _94_v35;
            _94_v35 = _dafny.MultiSet.fromElements((_41_v1).isLessThanOrEqualTo(_41_v1));
            let _95_v36;
            _95_v36 = _dafny.Set.fromElements(_41_v1);
            let _96_v37;
            _96_v37 = _module.D3.create_DC11();
            let _rhs3 = _93_v34;
            let _rhs4 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_41_v1, _41_v1), new BigNumber((((p0) ? (_95_v36) : (_95_v36))).length));
            let _rhs5 = _module.__default.fm15(_96_v37, p0, globalState);
            let _lhs2 = globalState;
            _93_v34 = _rhs3;
            _lhs2.f0 = _rhs4;
            _94_v35 = _rhs5;
            let _97_v38;
            _97_v38 = _dafny.Set.fromElements(true);
            let _98_v39;
            let _nw10 = new _module.C0();
            _nw10.__ctor(_41_v1, _97_v38);
            _98_v39 = _nw10;
          }
        }
      }
      let _109_v40;
      _109_v40 = new BigNumber(-284);
      let _source4 = _module.D4.create_DC15((new BigNumber(305)).isLessThanOrEqualTo(_109_v40), _109_v40);
      if (_source4.is_DC14) {
        r0 = p0;
        r0 = p0;
        r0 = !(p0);
        let _110_v41;
        let _nw11 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _110_v41 = _nw11;
        let _111_v42;
        let _nw12 = Array((new BigNumber(17)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _110_v41;
        _nw12[(_dafny.ONE).toNumber()] = _110_v41;
        _nw12[(new BigNumber(2)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(3)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(4)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(5)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(6)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(7)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(8)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(9)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(10)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(11)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(12)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(13)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(14)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(15)).toNumber()] = _110_v41;
        _nw12[(new BigNumber(16)).toNumber()] = _110_v41;
        _111_v42 = _nw12;
        let _index5 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_111_v42).length));
        (_111_v42)[_index5] = _110_v41;
        let _index6 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_111_v42).length));
        (_111_v42)[_index6] = _110_v41;
      } else if (_source4.is_DC15) {
        let _112___mcc_h10 = (_source4).cf23;
        let _113___mcc_h11 = (_source4).cf24;
        let _114_cf24 = _113___mcc_h11;
        let _115_cf23 = _112___mcc_h10;
        let _116_v43;
        _116_v43 = _dafny.Seq.UnicodeFromString("sficpirr");
        let _117_v44;
        _117_v44 = new _dafny.CodePoint('d'.codePointAt(0));
        let _118_v45;
        _118_v45 = _module.D1.create_DC3(_115_cf23, _module.__default.fm0(p0, _115_cf23, _114_cf24, globalState), _dafny.Seq.update(_dafny.Seq.Concat(_116_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_119_i5) {
  return new _dafny.CodePoint('k'.codePointAt(0));
})), _module.__default.safeIndex(_109_v40, new BigNumber((_dafny.Seq.Concat(_116_v43, _dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_120_i5) {
  return new _dafny.CodePoint('k'.codePointAt(0));
}))).length)), _117_v44), (new BigNumber((_116_v43).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_109_v40)).length)), (_dafny.ZERO).minus(_109_v40));
        let _rhs6 = _118_v45;
        let _rhs7 = _module.__default.fm4(_109_v40, p0, globalState);
        _118_v45 = _rhs6;
        _114_cf24 = _rhs7;
        _116_v43 = _dafny.Seq.UnicodeFromString("m");
        if (false) {
          (globalState).f4 = new BigNumber(206);
          let _121_v46;
          _121_v46 = _dafny.Set.fromElements(false, _115_cf23);
          let _122_v47;
          let _nw13 = new _module.C0();
          _nw13.__ctor(_114_cf24, _121_v46);
          _122_v47 = _nw13;
          let _123_v48;
          _123_v48 = _dafny.Map.Empty.slice().updateUnsafe(_122_v47,_115_cf23);
          let _124_v49;
          _124_v49 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(838),_122_v47);
          _115_cf23 = _module.__default.fm0(p0, !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_123_v48).update((((_124_v49).contains(_114_cf24)) ? ((_124_v49).get(_114_cf24)) : (_122_v47)), p0)).length),_dafny.Map.Empty.slice().updateUnsafe(_115_cf23,(_116_v43)[_module.__default.safeIndex((_122_v47).f6, new BigNumber((_116_v43).length))]))).contains(new BigNumber((_116_v43).length)), (_dafny.ZERO).minus((_122_v47).f6), globalState);
          let _125_v50;
          let _nw14 = new _module.C0();
          _nw14.__ctor((_122_v47).f6, _121_v46);
          _125_v50 = _nw14;
          let _126_v51;
          _126_v51 = _dafny.Set.fromElements(_124_v49);
          _115_cf23 = (_126_v51).IsSubsetOf(_126_v51);
          let _127_v52;
          _127_v52 = _dafny.Map.Empty.slice().updateUnsafe((_122_v47).f6,p0);
          _127_v52 = (_127_v52).update((_dafny.ZERO).minus((_125_v50).f6), p0);
        } else {
          r0 = !(p0);
          let _128_v53;
          _128_v53 = _dafny.Set.fromElements(_115_cf23);
          let _129_v54;
          let _nw15 = new _module.C0();
          _nw15.__ctor(_109_v40, _128_v53);
          _129_v54 = _nw15;
          let _130_v55;
          _130_v55 = _module.D4.create_DC13(_129_v54);
          let _131_v56;
          _131_v56 = _dafny.Map.Empty.slice().updateUnsafe(_130_v55,_115_cf23);
          let _132_v57;
          let _nw16 = new _module.C0();
          _nw16.__ctor(new BigNumber((_131_v56).length), _128_v53);
          _132_v57 = _nw16;
          let _133_v58;
          _133_v58 = _dafny.Seq.of(p0);
          let _134_v59;
          _134_v59 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_116_v43, _116_v43),!_dafny.Seq.contains(_133_v58, p0));
          _134_v59 = _134_v59;
          r0 = _115_cf23;
          let _135_v60;
          let _nw17 = new _module.C0();
          _nw17.__ctor((_dafny.ZERO).minus((_132_v57).f6), (_128_v53).Union(_dafny.Set.fromElements(_115_cf23)));
          _135_v60 = _nw17;
        }
        let _136_v61;
        _136_v61 = _dafny.Set.fromElements(_109_v40);
        let _137_v62;
        _137_v62 = _dafny.Map.Empty.slice().updateUnsafe(_117_v44,_136_v61);
        let _138_v63;
        let _nw18 = Array((new BigNumber(10)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = (new BigNumber((_116_v43).length)).minus(_114_cf24);
        _nw18[(_dafny.ONE).toNumber()] = _114_cf24;
        _nw18[(new BigNumber(2)).toNumber()] = _109_v40;
        _nw18[(new BigNumber(3)).toNumber()] = _109_v40;
        _nw18[(new BigNumber(4)).toNumber()] = new BigNumber(446);
        _nw18[(new BigNumber(5)).toNumber()] = _module.__default.fm4(_114_cf24, p0, globalState);
        _nw18[(new BigNumber(6)).toNumber()] = _module.__default.fm4(_109_v40, _115_cf23, globalState);
        _nw18[(new BigNumber(7)).toNumber()] = _109_v40;
        _nw18[(new BigNumber(8)).toNumber()] = _109_v40;
        _nw18[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_109_v40).minus(new BigNumber((_137_v62).length)));
        _138_v63 = _nw18;
        let _index7 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_138_v63).length));
        (_138_v63)[_index7] = _109_v40;
        let _139_v64;
        let _init0 = ((_140_v43) => function (_141_i6) {
          return _140_v43;
        })(_116_v43);
        let _nw19 = Array((new BigNumber(10)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw19.length); _i0_0++) {
          _nw19[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _139_v64 = _nw19;
        let _index8 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_139_v64).length));
        (_139_v64)[_index8] = _116_v43;
        let _142_v65;
        let _init1 = ((_143_v43) => function (_144_i7) {
          return !(!(_dafny.Seq.IsPrefixOf(_143_v43, _143_v43)));
        })(_116_v43);
        let _nw20 = Array((new BigNumber(11)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw20.length); _i0_1++) {
          _nw20[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _142_v65 = _nw20;
        let _145_v66;
        _145_v66 = _dafny.Seq.of(p0);
        let _index9 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_142_v65).length));
        (_142_v65)[_index9] = _dafny.areEqual(_dafny.Seq.of(p0, _115_cf23, _115_cf23), _145_v66);
        let _index10 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_138_v63).length));
        let _index11 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_139_v64).length));
        let _index12 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_142_v65).length));
        let _rhs8 = _109_v40;
        let _rhs9 = _dafny.Seq.UnicodeFromString("ihfjckhnn");
        let _rhs10 = _109_v40;
        let _rhs11 = ((_114_cf24).multipliedBy(new BigNumber((_116_v43).length))).minus(new BigNumber(205));
        let _rhs12 = p0;
        let _lhs3 = _138_v63;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_138_v63).length));
        let _lhs5 = _139_v64;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_139_v64).length));
        let _lhs7 = globalState;
        let _lhs8 = globalState;
        let _lhs9 = _142_v65;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_142_v65).length));
        _lhs3[_lhs4] = _rhs8;
        _lhs5[_lhs6] = _rhs9;
        _lhs7.f0 = _rhs10;
        _lhs8.f0 = _rhs11;
        _lhs9[_lhs10] = _rhs12;
      } else if (_source4.is_DC16) {
        let _146___mcc_h12 = (_source4).cf25;
        let _147___mcc_h13 = (_source4).cf26;
        let _148___mcc_h14 = (_source4).cf27;
        let _149___mcc_h15 = (_source4).cf28;
        let _150_cf28 = _149___mcc_h15;
        let _151_cf27 = _148___mcc_h14;
        let _152_cf26 = _147___mcc_h13;
        let _153_cf25 = _146___mcc_h12;
        let _154_v67;
        _154_v67 = _dafny.Seq.UnicodeFromString("adyhsx");
        let _155_v68;
        _155_v68 = _dafny.Map.Empty.slice().updateUnsafe(_154_v67,p0);
        let _156_v69;
        _156_v69 = _dafny.MultiSet.fromElements(_109_v40, new BigNumber((_155_v68).length));
        let _157_v70;
        _157_v70 = _dafny.Seq.of(new BigNumber((_154_v67).length), _151_cf27);
        let _158_v71;
        _158_v71 = _module.D2.create_DC7((_157_v70)[_module.__default.safeIndex(new BigNumber(467), new BigNumber((_157_v70).length))]);
        let _159_v72;
        _159_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-151),_158_v71);
        let _160_v73;
        _160_v73 = _dafny.MultiSet.fromElements(((((_156_v69).contains(_109_v40)) ? ((_156_v69).get(_109_v40)) : (new BigNumber((_159_v72).length)))).isLessThanOrEqualTo(_151_cf27));
        let _161_v74;
        let _init2 = ((_162_v70, _163_cf26) => function (_164_i8) {
          return (_164_i8).plus((_162_v70)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), ((_165_cf26) => function (_166_i9) {
            return _165_cf26;
          })(_163_cf26))).length), new BigNumber((_162_v70).length))]);
        })(_157_v70, _152_cf26);
        let _nw21 = Array((new BigNumber(6)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw21.length); _i0_2++) {
          _nw21[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _161_v74 = _nw21;
        let _167_v75;
        _167_v75 = _dafny.Map.Empty.slice().updateUnsafe(_161_v74,p0);
        let _168_v76;
        _168_v76 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(365),_dafny.Seq.UnicodeFromString("csu"));
        (globalState).f0 = (((_160_v73).contains(_module.__default.fm0(p0, (((_167_v75).contains(_161_v74)) ? ((_167_v75).get(_161_v74)) : (p0)), new BigNumber((_168_v76).length), globalState))) ? ((_160_v73).get(_module.__default.fm0(p0, (((_167_v75).contains(_161_v74)) ? ((_167_v75).get(_161_v74)) : (p0)), new BigNumber((_168_v76).length), globalState))) : (_109_v40));
        r0 = p0;
        let _169_v77;
        let _nw22 = Array((new BigNumber(23)).toNumber()).fill(false);
        _169_v77 = _nw22;
        _169_v77 = _169_v77;
        let _index13 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_161_v74).length));
        (_161_v74)[_index13] = (_151_cf27).plus(_109_v40);
        let _index14 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_161_v74).length));
        (_161_v74)[_index14] = _151_cf27;
      } else {
        let _170___mcc_h16 = (_source4).cf22;
        let _171_cf22 = _170___mcc_h16;
        let _172_v78;
        _172_v78 = _module.D3.create_DC10(p0, _109_v40, _109_v40);
        let _source5 = _172_v78;
        if (_source5.is_DC10) {
          let _173___mcc_h17 = (_source5).cf18;
          let _174___mcc_h18 = (_source5).cf19;
          let _175___mcc_h19 = (_source5).cf20;
          let _176_cf20 = _175___mcc_h19;
          let _177_cf19 = _174___mcc_h18;
          let _178_cf18 = _173___mcc_h17;
          let _179_v79;
          _179_v79 = _dafny.Seq.of(_109_v40);
          _178_cf18 = !_dafny.areEqual(_179_v79, _179_v79);
          _178_cf18 = _178_cf18;
          let _180_v80;
          let _nw23 = Array((new BigNumber(9)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = new BigNumber(-215);
          _nw23[(_dafny.ONE).toNumber()] = new BigNumber((_179_v79).length);
          _nw23[(new BigNumber(2)).toNumber()] = (_171_cf22).f6;
          _nw23[(new BigNumber(3)).toNumber()] = (_171_cf22).f6;
          _nw23[(new BigNumber(4)).toNumber()] = (_171_cf22).f6;
          _nw23[(new BigNumber(5)).toNumber()] = _177_cf19;
          _nw23[(new BigNumber(6)).toNumber()] = _176_cf20;
          _nw23[(new BigNumber(7)).toNumber()] = _109_v40;
          _nw23[(new BigNumber(8)).toNumber()] = _109_v40;
          _180_v80 = _nw23;
          let _181_v81;
          _181_v81 = _module.D2.create_DC6(_176_cf20, p0, _180_v80);
          let _182_v82;
          _182_v82 = _dafny.MultiSet.fromElements((_181_v81).dtor_cf12, (_171_cf22).f6);
          let _183_v83;
          _183_v83 = _dafny.Map.Empty.slice().updateUnsafe((_171_cf22).f6,_182_v82);
          (globalState).f4 = (_171_cf22).fm1(((((_183_v83).contains(_109_v40)) ? ((_183_v83).get(_109_v40)) : (_182_v82))).update((_dafny.ZERO).minus((_171_cf22).f6), _module.__default.abs((_171_cf22).f6)), _39_v0, globalState);
          _180_v80 = _180_v80;
        } else if (_source5.is_DC11) {
          let _184_v84;
          _184_v84 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
          let _185_v85;
          _185_v85 = _dafny.Seq.of(_109_v40, (_171_cf22).f6);
          let _186_v86;
          _186_v86 = _dafny.Map.Empty.slice().updateUnsafe((_171_cf22).f6,_185_v85);
          let _187_v87;
          _187_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_184_v84).length),_dafny.Seq.Concat(_185_v85, (((_186_v86).contains(new BigNumber(603))) ? ((_186_v86).get(new BigNumber(603))) : (_185_v85))));
          let _188_v88;
          _188_v88 = new _dafny.CodePoint('j'.codePointAt(0));
          let _189_v89;
          _189_v89 = _dafny.MultiSet.fromElements(p0, p0, !(p0));
          let _190_v90;
          _190_v90 = _dafny.Seq.UnicodeFromString("pcyod");
          let _191_v91;
          _191_v91 = _dafny.Seq.of((_module.__default.fm7(_189_v89, _190_v90, globalState)).update(p0, (_171_cf22).f6));
          _186_v86 = (_186_v86).update((_171_cf22).fm2((_module.__default.fm16(_109_v40, _188_v88, globalState)).dtor_cf23, _185_v85, _185_v85, (_191_v91)[_module.__default.safeIndex(_109_v40, new BigNumber((_191_v91).length))], globalState), _185_v85);
          r0 = !(!(p0));
          let _192_v92;
          _192_v92 = _dafny.Map.Empty.slice().updateUnsafe(_109_v40,_module.__default.fm0(p0, p0, (_171_cf22).f6, globalState));
          let _193_v93;
          _193_v93 = _dafny.Set.fromElements(_109_v40, _109_v40, (_171_cf22).f6, (_171_cf22).f6, _109_v40);
          _192_v92 = (_192_v92).update((_171_cf22).f6, (function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(312), new BigNumber(235))) {
              let _194_v94 = _compr_14;
              if (((new BigNumber(312)).isLessThanOrEqualTo(_194_v94)) && ((_194_v94).isLessThan(new BigNumber(235)))) {
                _coll14.add(_module.__default.safeModuloInt(_194_v94, (_171_cf22).f6));
              }
            }
            return _coll14;
          }()).IsProperSubsetOf(_193_v93));
          (globalState).f4 = (_dafny.ZERO).minus((_171_cf22).f6);
        } else if (_source5.is_DC9) {
          let _195___mcc_h20 = (_source5).cf17;
          let _196_cf17 = _195___mcc_h20;
          let _197_v95;
          let _nw24 = Array((new BigNumber(29)).toNumber()).fill(false);
          _197_v95 = _nw24;
          _197_v95 = _197_v95;
          let _198_v96;
          let _nw25 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _198_v96 = _nw25;
          let _rhs13 = p0;
          let _rhs14 = _198_v96;
          let _rhs15 = _109_v40;
          let _rhs16 = (_171_cf22).f6;
          let _lhs11 = globalState;
          let _lhs12 = globalState;
          r0 = _rhs13;
          _198_v96 = _rhs14;
          _lhs11.f0 = _rhs15;
          _lhs12.f4 = _rhs16;
          let _199_v97;
          let _nw26 = new _module.C0();
          _nw26.__ctor((_171_cf22).f6, (_171_cf22).f7);
          _199_v97 = _nw26;
          let _200_v98;
          _200_v98 = new _dafny.CodePoint('c'.codePointAt(0));
          _200_v98 = _200_v98;
        } else {
          let _201___mcc_h21 = (_source5).cf21;
          let _202_cf21 = _201___mcc_h21;
          let _nw27 = new _module.C0();
          _nw27.__ctor(new BigNumber((function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of _dafny.IntegerRange(new BigNumber(320), new BigNumber(-184))) {
              let _203_v99 = _compr_15;
              if (((new BigNumber(320)).isLessThanOrEqualTo(_203_v99)) && ((_203_v99).isLessThan(new BigNumber(-184)))) {
                _coll15.push([(_203_v99).plus((_171_cf22).f6),p0]);
              }
            }
            return _coll15;
          }()).length), (_171_cf22).f7);
          _171_cf22 = _nw27;
          let _204_v100;
          let _nw28 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _204_v100 = _nw28;
          _204_v100 = _204_v100;
          let _205_v101;
          _205_v101 = _dafny.Seq.of(p0);
          let _206_v102;
          let _nw29 = Array((new BigNumber(15)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = p0;
          _nw29[(_dafny.ONE).toNumber()] = p0;
          _nw29[(new BigNumber(2)).toNumber()] = !(!(p0)) || (p0);
          _nw29[(new BigNumber(3)).toNumber()] = p0;
          _nw29[(new BigNumber(4)).toNumber()] = p0;
          _nw29[(new BigNumber(5)).toNumber()] = (_205_v101)[_module.__default.safeIndex((_171_cf22).f6, new BigNumber((_205_v101).length))];
          _nw29[(new BigNumber(6)).toNumber()] = p0;
          _nw29[(new BigNumber(7)).toNumber()] = (_172_v78).dtor_cf18;
          _nw29[(new BigNumber(8)).toNumber()] = (p0) === (p0);
          _nw29[(new BigNumber(9)).toNumber()] = false;
          _nw29[(new BigNumber(10)).toNumber()] = p0;
          _nw29[(new BigNumber(11)).toNumber()] = p0;
          _nw29[(new BigNumber(12)).toNumber()] = p0;
          _nw29[(new BigNumber(13)).toNumber()] = p0;
          _nw29[(new BigNumber(14)).toNumber()] = _module.__default.fm0(!(p0), false, (_dafny.ZERO).minus(new BigNumber((_205_v101).length)), globalState);
          _206_v102 = _nw29;
          let _index15 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_206_v102).length));
          (_206_v102)[_index15] = p0;
          let _index16 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_206_v102).length));
          (_206_v102)[_index16] = !(p0);
          let _207_v103;
          _207_v103 = _dafny.Seq.UnicodeFromString("pibona");
          let _208_v104;
          _208_v104 = _dafny.Seq.of(_109_v40, new BigNumber((_207_v103).length));
          let _209_v105;
          _209_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, (_206_v102)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_206_v102).length))], new BigNumber((_208_v104).length), globalState),_171_cf22);
          let _210_v106;
          _210_v106 = _209_v105;
          let _211_v107;
          let _nw30 = Array((new BigNumber(3)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, (_206_v102)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_206_v102).length))], (_171_cf22).f6, globalState),_171_cf22);
          _nw30[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(false,_171_cf22);
          _nw30[(new BigNumber(2)).toNumber()] = ((_210_v106)).Merge(_209_v105);
          _211_v107 = _nw30;
          _211_v107 = _211_v107;
        }
        if (!(p0)) {
          let _212_v108;
          _212_v108 = new _dafny.CodePoint('w'.codePointAt(0));
          let _213_v109;
          _213_v109 = _dafny.Map.Empty.slice().updateUnsafe(_212_v108,_171_cf22);
          _213_v109 = (_213_v109).update(((!(p0)) ? (_212_v108) : (_212_v108)), _171_cf22);
          let _214_v110;
          let _nw31 = new _module.C0();
          _nw31.__ctor(((_171_cf22).f6).multipliedBy((_171_cf22).f6), ((_171_cf22).f7).Intersect(_dafny.Set.fromElements(p0)));
          _214_v110 = _nw31;
          (globalState).f4 = (_171_cf22).f6;
          let _215_v111;
          _215_v111 = _dafny.MultiSet.fromElements((_171_cf22).f6, (_171_cf22).f6);
          let _216_v112;
          _216_v112 = _dafny.Seq.of(new BigNumber((_215_v111).cardinality()), new BigNumber(-745));
          let _217_v114;
          _217_v114 = _dafny.Seq.of(new BigNumber((function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of (_216_v112).Elements) {
              let _218_v113 = _compr_16;
              if (_dafny.Seq.contains(_216_v112, _218_v113)) {
                _coll16.add(_module.__default.safeDivisionInt(_218_v113, new BigNumber((_dafny.Seq.UnicodeFromString("yr")).length)));
              }
            }
            return _coll16;
          }()).length));
          _217_v114 = _217_v114;
          r0 = (((_171_cf22).f6).multipliedBy((_171_cf22).f6)).isLessThan(_109_v40);
        } else {
          let _219_v115;
          _219_v115 = _dafny.Set.fromElements(new BigNumber(682));
          _219_v115 = _dafny.Set.fromElements(_109_v40, _109_v40, _109_v40);
          let _220_v116;
          _220_v116 = _dafny.Map.Empty.slice().updateUnsafe(_171_cf22,false);
          let _221_v117;
          _221_v117 = _dafny.Seq.of(p0, false, !(_220_v116).equals(_220_v116));
          _221_v117 = _dafny.Seq.Concat(((false) ? (_221_v117) : (_221_v117)), _221_v117);
          let _222_v118;
          let _nw32 = Array((new BigNumber(17)).toNumber()).fill([]);
          _222_v118 = _nw32;
          _222_v118 = _222_v118;
          let _223_v119;
          let _nw33 = Array((new BigNumber(17)).toNumber()).fill(false);
          _223_v119 = _nw33;
          let _224_v120;
          _224_v120 = new _dafny.CodePoint('y'.codePointAt(0));
          let _225_v122;
          _225_v122 = _dafny.Map.Empty.slice().updateUnsafe(_224_v120,new BigNumber(-9));
          let _index17 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_223_v119).length));
          (_223_v119)[_index17] = (_dafny.Map.Empty.slice().updateUnsafe(_224_v120,p0)).equals(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_225_v122).Keys.Elements) {
              let _226_v121 = _compr_17;
              if ((_225_v122).contains(_226_v121)) {
                _coll17.push([_226_v121,p0]);
              }
            }
            return _coll17;
          }());
          let _index18 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_223_v119).length));
          (_223_v119)[_index18] = p0;
          let _227_v123;
          let _init3 = ((_228_cf22) => function (_229_i10) {
            return (_229_i10).plus((_228_cf22).f6);
          })(_171_cf22);
          let _nw34 = Array((new BigNumber(19)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw34.length); _i0_3++) {
            _nw34[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _227_v123 = _nw34;
          let _230_v124;
          _230_v124 = _dafny.Map.Empty.slice().updateUnsafe(_223_v119,_227_v123);
          let _index19 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_223_v119).length));
          (_223_v119)[_index19] = !(_230_v124).contains(_223_v119);
        }
        let _231_v125;
        _231_v125 = _dafny.Set.fromElements((_171_cf22).f6);
        let _232_v126;
        _232_v126 = _dafny.Map.Empty.slice().updateUnsafe(_231_v125,_109_v40);
        let _233_v128;
        _233_v128 = _dafny.Map.Empty.slice().updateUnsafe(_109_v40,function () {
          let _coll18 = new _dafny.Set();
          for (const _compr_18 of (_232_v126).Keys.Elements) {
            let _234_v127 = _compr_18;
            if ((_232_v126).contains(_234_v127)) {
              _coll18.add(_234_v127);
            }
          }
          return _coll18;
        }());
        let _235_v129;
        _235_v129 = _dafny.Set.fromElements(_231_v125, _231_v125);
        _109_v40 = new BigNumber(((((((_233_v128).contains(new BigNumber(398))) ? ((_233_v128).get(new BigNumber(398))) : (_235_v129))).Difference(_module.__default.fm17(_109_v40, p0, (_dafny.ZERO).minus((_171_cf22).f6), globalState))).Union(_235_v129)).length);
        let _236_v130;
        _236_v130 = _dafny.MultiSet.fromElements((_171_cf22).f6);
        (globalState).f0 = ((_171_cf22).fm1(_236_v130, _39_v0, globalState)).multipliedBy((_171_cf22).f6);
      }
      let _237_v131;
      _237_v131 = _dafny.Seq.UnicodeFromString("xukko");
      if (((_dafny.Seq.IsProperPrefixOf(_237_v131, _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_238_i11) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }))) ? (((_dafny.ZERO).minus(_109_v40)).isEqualTo(_109_v40)) : (((_dafny.ZERO).minus(_109_v40)).isLessThan(_109_v40)))) {
        r0 = p0;
        if ((_109_v40).isLessThanOrEqualTo((_109_v40).minus(_109_v40))) {
          let _239_v132;
          _239_v132 = _module.D4.create_DC15(p0, new BigNumber((_dafny.Seq.UnicodeFromString("rq")).length));
          let _240_v133;
          _240_v133 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, p0, new BigNumber((_237_v131).length), globalState),(_239_v132).dtor_cf23);
          let _241_v134;
          _241_v134 = _dafny.Set.fromElements(_109_v40, _109_v40);
          let _242_v135;
          _242_v135 = _dafny.Set.fromElements(p0, p0);
          let _243_v136;
          let _nw35 = new _module.C0();
          _nw35.__ctor(_109_v40, _242_v135);
          _243_v136 = _nw35;
          let _244_v137;
          _244_v137 = _dafny.Map.Empty.slice().updateUnsafe(_243_v136,_109_v40);
          let _245_v138;
          _245_v138 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true)).length), (((_244_v137).contains(_243_v136)) ? ((_244_v137).get(_243_v136)) : ((_243_v136).f6)));
          _240_v133 = (_240_v133).update((_241_v134).IsDisjointFrom(_dafny.Set.fromElements(_109_v40, new BigNumber((_dafny.Set.fromElements(p0)).length), new BigNumber((_237_v131).length), (((_245_v138).contains(_109_v40)) ? ((_245_v138).get(_109_v40)) : (_109_v40)))), p0);
          let _246_v139;
          _246_v139 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18(_109_v40, globalState),_109_v40);
          let _247_v140;
          _247_v140 = new _dafny.CodePoint('l'.codePointAt(0));
          let _248_v141;
          _248_v141 = _dafny.Map.Empty.slice().updateUnsafe(_109_v40,p0);
          (globalState).f0 = (_dafny.ZERO).minus((((_246_v139).contains(_247_v140)) ? ((_246_v139).get(_247_v140)) : (new BigNumber(((_248_v141).update(_109_v40, p0)).length))));
          let _249_v142;
          let _nw36 = new _module.C0();
          _nw36.__ctor(((_243_v136).f6).multipliedBy(_109_v40), (_243_v136).f7);
          _249_v142 = _nw36;
          _247_v140 = _247_v140;
          let _250_v143;
          _250_v143 = _dafny.MultiSet.fromElements(p0);
          let _251_v144;
          _251_v144 = _dafny.Map.Empty.slice().updateUnsafe(_250_v143,_247_v140);
          let _252_v145;
          _252_v145 = _dafny.Seq.of(p0, p0);
          _251_v144 = (_251_v144).update(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_252_v145, _252_v145)), _247_v140);
        } else {
          r0 = p0;
          let _253_v146;
          let _nw37 = Array((new BigNumber(12)).toNumber());
          _253_v146 = _nw37;
          let _254_v147;
          _254_v147 = _dafny.Set.fromElements(p0);
          let _255_v148;
          let _nw38 = new _module.C0();
          _nw38.__ctor(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_109_v40)).length), _254_v147);
          _255_v148 = _nw38;
          let _index20 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_253_v146).length));
          (_253_v146)[_index20] = _255_v148;
          let _index21 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_253_v146).length));
          let _rhs17 = _255_v148;
          let _rhs18 = _109_v40;
          let _rhs19 = (_255_v148).f6;
          let _rhs20 = _255_v148;
          let _lhs13 = _253_v146;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_253_v146).length));
          let _lhs15 = globalState;
          let _lhs16 = globalState;
          _lhs13[_lhs14] = _rhs17;
          _lhs15.f4 = _rhs18;
          _lhs16.f4 = _rhs19;
          _255_v148 = _rhs20;
          let _256_v149;
          let _init4 = ((_257_v148, _258_v40, _259_p0, _260_v131) => function (_261_i12) {
            return (_dafny.MultiSet.fromElements((_257_v148).f6, (_257_v148).f6)).Intersect(_dafny.MultiSet.fromElements((_257_v148).f6, _258_v40, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_259_p0)).length), _258_v40, new BigNumber((_260_v131).length), _258_v40)).length)));
          })(_255_v148, _109_v40, p0, _237_v131);
          let _nw39 = Array((new BigNumber(7)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw39.length); _i0_4++) {
            _nw39[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _256_v149 = _nw39;
          let _262_v150;
          _262_v150 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_109_v40, (_255_v148).f6)).length), (_dafny.ZERO).minus(_109_v40));
          let _index22 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_256_v149).length));
          (_256_v149)[_index22] = _262_v150;
          let _index23 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_256_v149).length));
          (_256_v149)[_index23] = (((p0) ? (_262_v150) : (_262_v150))).Union(_262_v150);
          let _263_v151;
          _263_v151 = _dafny.MultiSet.fromElements(!(p0));
          (globalState).f0 = ((_255_v148).f6).plus(new BigNumber((_263_v151).cardinality()));
          r0 = _dafny.Seq.IsPrefixOf(_237_v131, _237_v131);
        }
        (globalState).f0 = ((p0) ? ((_dafny.ZERO).minus(_109_v40)) : (_109_v40));
        let _264_v152;
        let _nw40 = new _module.C0();
        _nw40.__ctor(_109_v40, _dafny.Set.fromElements(p0));
        _264_v152 = _nw40;
        let _265_v153;
        _265_v153 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_264_v152).f6);
        let _266_v154;
        _266_v154 = _dafny.Seq.of(_109_v40, (_264_v152).fm2(p0, _dafny.Seq.update(_dafny.Seq.of(_109_v40), _module.__default.safeIndex(_109_v40, new BigNumber((_dafny.Seq.of(_109_v40)).length)), _109_v40), _dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), ((_267_v152) => function (_268_i14) {
          return (_267_v152).f6;
        })(_264_v152)), _265_v153, globalState));
        _109_v40 = _module.__default.safeModuloInt(_109_v40, (_264_v152).fm2(true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(27)), ((_269_v152) => function (_270_i13) {
          return (_269_v152).f6;
        })(_264_v152)), _266_v154, _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_module.__default.fm3(_dafny.Seq.of(p0), globalState)).length)), globalState));
      } else {
        (globalState).f0 = _109_v40;
        let _271_v155;
        let _nw41 = Array((new BigNumber(14)).toNumber()).fill(false);
        _271_v155 = _nw41;
        let _272_v156;
        _272_v156 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_271_v155),_dafny.Seq.UnicodeFromString("x"));
        let _273_v157;
        _273_v157 = _dafny.Seq.of(_271_v155, _271_v155);
        let _274_v158;
        _274_v158 = new _dafny.CodePoint('t'.codePointAt(0));
        let _275_v159;
        _275_v159 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jxrvpul"), _dafny.Seq.update(_237_v131, _module.__default.safeIndex(new BigNumber(((((_272_v156).contains(_dafny.Seq.update(_273_v157, _module.__default.safeIndex(_109_v40, new BigNumber((_273_v157).length)), _271_v155))) ? ((_272_v156).get(_dafny.Seq.update(_273_v157, _module.__default.safeIndex(_109_v40, new BigNumber((_273_v157).length)), _271_v155))) : (_237_v131))).length), new BigNumber((_237_v131).length)), _274_v158));
        let _276_v160;
        _276_v160 = _dafny.Seq.of(p0);
        let _277_v161;
        let _nw42 = Array((new BigNumber(29)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_237_v131, _dafny.Seq.update(_237_v131, _module.__default.safeIndex(_109_v40, new BigNumber((_237_v131).length)), _274_v158));
        _nw42[(new BigNumber(3)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(4)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_237_v131, _dafny.Seq.Create(_module.__default.abs(new BigNumber(328)), ((_278_v131, _279_p0, _280_v158) => function (_281_i15) {
          return (_module.D6.create_DC19(new BigNumber((_278_v131).length), _279_p0, _280_v158)).dtor_cf33;
        })(_237_v131, p0, _274_v158)));
        _nw42[(new BigNumber(6)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(7)).toNumber()] = _module.__default.fm19(false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-351)), ((_282_v40) => function (_283_i16) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_282_v40),true)).length);
        })(_109_v40))).length), globalState);
        _nw42[(new BigNumber(8)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(new BigNumber(11)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(12)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(13)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(14)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_275_v159, _module.__default.safeIndex(new BigNumber((_237_v131).length), new BigNumber((_275_v159).length)), _237_v131);
        _nw42[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qrr")), _275_v159);
        _nw42[(new BigNumber(17)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(18)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_275_v159, _dafny.Seq.update(_275_v159, _module.__default.safeIndex(_109_v40, new BigNumber((_275_v159).length)), _dafny.Seq.UnicodeFromString("igyetaiv")));
        _nw42[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), ((_284_v131) => function (_285_i17) {
          return _284_v131;
        })(_237_v131));
        _nw42[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(_module.__default.fm3(_276_v160, globalState));
        _nw42[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_237_v131, _dafny.Seq.Create(_module.__default.abs(new BigNumber(353)), ((_286_v158) => function (_287_i18) {
          return _286_v158;
        })(_274_v158)));
        _nw42[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_275_v159, _275_v159);
        _nw42[(new BigNumber(25)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(26)).toNumber()] = _275_v159;
        _nw42[(new BigNumber(27)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm19(p0, _109_v40, globalState), _275_v159);
        _nw42[(new BigNumber(28)).toNumber()] = _dafny.Seq.of(_237_v131, _237_v131);
        _277_v161 = _nw42;
        let _index24 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_277_v161).length));
        (_277_v161)[_index24] = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(480)), function (_288_i19) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
        let _index25 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_277_v161).length));
        (_277_v161)[_index25] = _dafny.Seq.Concat(_275_v159, _275_v159);
        let _289_v162;
        let _nw43 = Array((new BigNumber(25)).toNumber()).fill([]);
        _289_v162 = _nw43;
        let _290_v163;
        let _nw44 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _290_v163 = _nw44;
        let _index26 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_289_v162).length));
        (_289_v162)[_index26] = _290_v163;
        let _index27 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_289_v162).length));
        let _nw45 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
        (_289_v162)[_index27] = _nw45;
        let _291_v165;
        let _init5 = ((_292_p0, _293_v131, _294_v158, _295_v40) => function (_296_i20) {
          return _dafny.Seq.of(_module.D1.create_DC3(_292_p0, _292_p0, _293_v131, true, new BigNumber((_293_v131).length)), _module.D1.create_DC3(_292_p0, true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), ((_297_v158) => function (_298_i21) {
  return _297_v158;
})(_294_v158)), _292_p0, _295_v40), _module.D1.create_DC3(_292_p0, _292_p0, _293_v131, !(false), new BigNumber((function () {
  let _coll19 = new _dafny.Set();
  for (const _compr_19 of (_dafny.Seq.of(_294_v158, _294_v158)).Elements) {
    let _299_v164 = _compr_19;
    if (_dafny.Seq.contains(_dafny.Seq.of(_294_v158, _294_v158), _299_v164)) {
      _coll19.add(_299_v164);
    }
  }
  return _coll19;
}()).length)));
        })(p0, _237_v131, _274_v158, _109_v40);
        let _nw46 = Array((new BigNumber(9)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw46.length); _i0_5++) {
          _nw46[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _291_v165 = _nw46;
        let _300_v166;
        _300_v166 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.FromArray(_276_v160));
        let _rhs21 = (new BigNumber((_300_v166).length)).multipliedBy(_109_v40);
        let _rhs22 = _109_v40;
        let _rhs23 = p0;
        let _rhs24 = ((p0) ? (_291_v165) : (_291_v165));
        let _lhs17 = globalState;
        _109_v40 = _rhs21;
        _lhs17.f0 = _rhs22;
        r0 = _rhs23;
        _291_v165 = _rhs24;
        let _301_v167;
        _301_v167 = _module.D2.create_DC7(new BigNumber((_237_v131).length));
        let _302_v168;
        _302_v168 = _dafny.Set.fromElements(p0);
        let _303_v169;
        let _nw47 = new _module.C0();
        _nw47.__ctor((_109_v40).plus((_301_v167).dtor_cf15), _302_v168);
        _303_v169 = _nw47;
      }
      let _304_v170;
      _304_v170 = new _dafny.CodePoint('t'.codePointAt(0));
      let _305_v171;
      _305_v171 = _module.D6.create_DC19(_109_v40, _module.__default.fm0(!(!(p0)), p0, _109_v40, globalState), _304_v170);
      let _pat_let_tv5 = _39_v0;
      let _pat_let_tv6 = _39_v0;
      let _pat_let_tv7 = p0;
      let _pat_let_tv8 = _109_v40;
      let _pat_let_tv9 = _109_v40;
      let _pat_let_tv10 = _109_v40;
      let _pat_let_tv11 = _109_v40;
      let _pat_let_tv12 = _109_v40;
      let _pat_let_tv13 = _109_v40;
      r0 = function (_source6) {
        if (_source6.is_DC19) {
          let _306___mcc_h22 = (_source6).cf31;
          let _307___mcc_h23 = (_source6).cf32;
          let _308___mcc_h24 = (_source6).cf33;
          let _309_cf33 = _308___mcc_h24;
          let _310_cf32 = _307___mcc_h23;
          let _311_cf31 = _306___mcc_h22;
          return _310_cf32;
        } else if (_source6.is_DC20) {
          let _312___mcc_h25 = (_source6).cf34;
          let _313_cf34 = _312___mcc_h25;
          return (new BigNumber((function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_dafny.Seq.of(_pat_let_tv5)).Elements) {
              let _314_v172 = _compr_20;
              if (_dafny.Seq.contains(_dafny.Seq.of(_pat_let_tv6), _314_v172)) {
                _coll20.push([_314_v172,_pat_let_tv7]);
              }
            }
            return _coll20;
          }()).length)).isLessThan(_pat_let_tv8);
        } else if (_source6.is_DC21) {
          return (_dafny.MultiSet.fromElements(new BigNumber(14), _pat_let_tv9)).IsSubsetOf(_dafny.MultiSet.fromElements(_pat_let_tv10, _pat_let_tv11));
        } else {
          let _315___mcc_h26 = (_source6).cf30;
          let _316_cf30 = _315___mcc_h26;
          return (_module.D3.create_DC10(true, _pat_let_tv12, _pat_let_tv13)).dtor_cf18;
        }
      }(_305_v171);
      let _317_i22;
      _317_i22 = _dafny.ZERO;
      L1: {
        while (!(p0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_317_i22)) {
              break L1;
            }
            _317_i22 = (_317_i22).plus(_dafny.ONE);
            r0 = p0;
            r0 = p0;
            (globalState).f0 = _109_v40;
            let _318_v173;
            _318_v173 = _dafny.Set.fromElements(p0, p0);
            let _319_v174;
            let _nw48 = new _module.C0();
            _nw48.__ctor((_dafny.ZERO).minus(_module.__default.fm4(_109_v40, p0, globalState)), _318_v173);
            _319_v174 = _nw48;
          }
        }
      }
      r0 = p0;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _320_v0;
      _320_v0 = new BigNumber(-45);
      let _321_v1;
      _321_v1 = true;
      let _322_v2;
      _322_v2 = _dafny.Seq.of(_321_v1, _321_v1);
      let _323_v3;
      _323_v3 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,_322_v2);
      let _324_globalState;
      let _nw49 = new _module.GlobalState();
      _nw49.__ctor(new BigNumber(977), new BigNumber(424), true, new BigNumber(358), new BigNumber(-429), _323_v3);
      _324_globalState = _nw49;
      let _325_i0;
      _325_i0 = _dafny.ZERO;
      L2: {
        while ((((_321_v1) ? (new BigNumber(548)) : (_320_v0))).isLessThanOrEqualTo(new BigNumber(671))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_325_i0)) {
              break L2;
            }
            _325_i0 = (_325_i0).plus(_dafny.ONE);
            let _326_v4;
            _326_v4 = new _dafny.CodePoint('d'.codePointAt(0));
            let _327_v5;
            _327_v5 = _dafny.Seq.UnicodeFromString("ja");
            _326_v4 = (_327_v5)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_321_v1), _dafny.Seq.of(_321_v1))).length), new BigNumber((_327_v5).length))];
            let _328_v6;
            let _nw50 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _328_v6 = _nw50;
            let _329_v7;
            _329_v7 = _dafny.MultiSet.fromElements(_328_v6, _328_v6, _328_v6);
            let _330_v8;
            _330_v8 = _module.D0.create_DC0(_321_v1);
            let _331_v9;
            let _nw51 = Array((new BigNumber(29)).toNumber());
            _nw51[(_dafny.ZERO).toNumber()] = !(_321_v1);
            _nw51[(_dafny.ONE).toNumber()] = _321_v1;
            _nw51[(new BigNumber(2)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(3)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(4)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(5)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(6)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(7)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(8)).toNumber()] = true;
            _nw51[(new BigNumber(9)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(10)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(11)).toNumber()] = true;
            _nw51[(new BigNumber(12)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(13)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(14)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(15)).toNumber()] = _module.__default.fm0(_321_v1, _321_v1, _320_v0, _324_globalState);
            _nw51[(new BigNumber(16)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(17)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(18)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(19)).toNumber()] = !(_321_v1);
            _nw51[(new BigNumber(20)).toNumber()] = false;
            _nw51[(new BigNumber(21)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(22)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(23)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(24)).toNumber()] = true;
            _nw51[(new BigNumber(25)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(26)).toNumber()] = _321_v1;
            _nw51[(new BigNumber(27)).toNumber()] = (_330_v8).dtor_cf0;
            _nw51[(new BigNumber(28)).toNumber()] = _321_v1;
            _331_v9 = _nw51;
            let _332_v10;
            let _init6 = ((_333_v1) => function (_334_i1) {
              return _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_333_v1)).cardinality()));
            })(_321_v1);
            let _nw52 = Array((new BigNumber(20)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw52.length); _i0_6++) {
              _nw52[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _332_v10 = _nw52;
            let _335_v11;
            _335_v11 = _module.D0.create_DC1(_329_v7, _331_v9, _332_v10);
            let _source7 = _335_v11;
            if (_source7.is_DC0) {
              let _336___mcc_h0 = (_source7).cf0;
              let _337_cf0 = _336___mcc_h0;
              let _338_v12;
              _338_v12 = _dafny.Set.fromElements(_320_v0, _320_v0);
              let _339_v13;
              _339_v13 = _dafny.Seq.of(_338_v12, _338_v12);
              _339_v13 = _339_v13;
              let _340_v14;
              _340_v14 = _dafny.Set.fromElements(true, _321_v1, _321_v1);
              let _341_v15;
              let _nw53 = new _module.C0();
              _nw53.__ctor(_module.__default.safeDivisionInt(_320_v0, _320_v0), (_340_v14).Difference(_340_v14));
              _341_v15 = _nw53;
              let _342_v16;
              _342_v16 = _dafny.Map.Empty.slice().updateUnsafe(_321_v1,_341_v15);
              _342_v16 = (_342_v16).update(_337_cf0, (((_342_v16).contains(_321_v1)) ? ((_342_v16).get(_321_v1)) : (_341_v15)));
              _328_v6 = ((true) ? (_328_v6) : (_328_v6));
            } else {
              let _343___mcc_h1 = (_source7).cf1;
              let _344___mcc_h2 = (_source7).cf2;
              let _345___mcc_h3 = (_source7).cf3;
              let _346_cf3 = _345___mcc_h3;
              let _347_cf2 = _344___mcc_h2;
              let _348_cf1 = _343___mcc_h1;
              _328_v6 = _328_v6;
              let _349_v17;
              _349_v17 = _dafny.MultiSet.fromElements(_321_v1);
              _349_v17 = (_349_v17).Union(_349_v17);
              let _350_v18;
              let _out0;
              _out0 = _module.__default.m0(!(((_321_v1) ? (_321_v1) : (_321_v1))), _324_globalState);
              _350_v18 = _out0;
              _320_v0 = new BigNumber((_dafny.Seq.UnicodeFromString("nroqns")).length);
            }
            (_324_globalState).f4 = _320_v0;
            let _351_v19;
            _351_v19 = _dafny.Set.fromElements(_320_v0, (_320_v0).minus(new BigNumber((_327_v5).length)), new BigNumber((_dafny.Seq.Concat(_327_v5, _327_v5)).length), _320_v0, _module.__default.safeDivisionInt(_320_v0, _320_v0));
            let _index28 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_328_v6).length));
            (_328_v6)[_index28] = _320_v0;
            let _index29 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_328_v6).length));
            let _rhs25 = _321_v1;
            let _rhs26 = _351_v19;
            let _rhs27 = _320_v0;
            let _lhs18 = _328_v6;
            let _lhs19 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_328_v6).length));
            _321_v1 = _rhs25;
            _351_v19 = _rhs26;
            _lhs18[_lhs19] = _rhs27;
          }
        }
      }
      let _352_v20;
      _352_v20 = _dafny.Seq.UnicodeFromString("ss");
      let _353_v21;
      _353_v21 = _dafny.MultiSet.fromElements(_320_v0, _320_v0);
      let _source8 = _module.D0.create_DC0((new BigNumber((_353_v21).cardinality())).isLessThan(new BigNumber((_352_v20).length)));
      if (_source8.is_DC0) {
        let _354___mcc_h4 = (_source8).cf0;
        let _355_cf0 = _354___mcc_h4;
        let _356_v22;
        let _nw54 = Array((new BigNumber(7)).toNumber()).fill(false);
        _356_v22 = _nw54;
        let _index30 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length));
        (_356_v22)[_index30] = (_355_cf0) || (true);
        let _357_v23;
        _357_v23 = new _dafny.CodePoint('s'.codePointAt(0));
        let _358_v24;
        _358_v24 = _dafny.MultiSet.fromElements(_module.__default.fm3(_322_v2, _324_globalState));
        let _index31 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length));
        let _rhs28 = _321_v1;
        let _rhs29 = _357_v23;
        let _rhs30 = _320_v0;
        let _rhs31 = new BigNumber((_358_v24).cardinality());
        let _lhs20 = _356_v22;
        let _lhs21 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length));
        let _lhs22 = _324_globalState;
        let _lhs23 = _324_globalState;
        _lhs20[_lhs21] = _rhs28;
        _357_v23 = _rhs29;
        _lhs22.f4 = _rhs30;
        _lhs23.f0 = _rhs31;
        let _359_v25;
        _359_v25 = _dafny.Map.Empty.slice().updateUnsafe(_355_cf0,_module.__default.safeDivisionInt(_320_v0, _320_v0));
        let _360_v26;
        _360_v26 = _dafny.MultiSet.fromElements((_356_v22)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length))]);
        _359_v25 = (_359_v25).update((_dafny.MultiSet.FromArray(_322_v2)).IsDisjointFrom((_360_v26).update(!(!(true)), _module.__default.abs(_320_v0))), (_dafny.ZERO).minus((_dafny.ZERO).minus(_320_v0)));
        _357_v23 = _357_v23;
        let _361_v27;
        _361_v27 = _dafny.Map.Empty.slice().updateUnsafe((_356_v22)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length))],_dafny.Set.fromElements(_321_v1, (_356_v22)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length))], _321_v1, (_356_v22)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_356_v22).length))], _355_cf0));
        let _362_v28;
        _362_v28 = _dafny.Set.fromElements(false, false, _355_cf0);
        let _363_v29;
        let _nw55 = Array((new BigNumber(22)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = _320_v0;
        _nw55[(_dafny.ONE).toNumber()] = _320_v0;
        _nw55[(new BigNumber(2)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(3)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(4)).toNumber()] = (_320_v0).minus(_module.__default.fm4(_320_v0, _355_cf0, _324_globalState));
        _nw55[(new BigNumber(5)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(6)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_352_v20, _352_v20)).length);
        _nw55[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_320_v0, _320_v0);
        _nw55[(new BigNumber(9)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(10)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(11)).toNumber()] = new BigNumber(233);
        _nw55[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(_320_v0, _320_v0);
        _nw55[(new BigNumber(13)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(14)).toNumber()] = (_320_v0).minus(_320_v0);
        _nw55[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_320_v0));
        _nw55[(new BigNumber(16)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(17)).toNumber()] = (new BigNumber(((_361_v27).update(_321_v1, _362_v28)).length)).multipliedBy(_320_v0);
        _nw55[(new BigNumber(18)).toNumber()] = _module.__default.safeModuloInt(_320_v0, _320_v0);
        _nw55[(new BigNumber(19)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(20)).toNumber()] = _320_v0;
        _nw55[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("oywkaa"), _dafny.Seq.UnicodeFromString("k"))).length);
        _363_v29 = _nw55;
        let _index32 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_363_v29).length));
        (_363_v29)[_index32] = _320_v0;
        let _index33 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_363_v29).length));
        (_363_v29)[_index33] = _module.__default.safeModuloInt(_320_v0, _320_v0);
      } else {
        let _364___mcc_h5 = (_source8).cf1;
        let _365___mcc_h6 = (_source8).cf2;
        let _366___mcc_h7 = (_source8).cf3;
        let _367_cf3 = _366___mcc_h7;
        let _368_cf2 = _365___mcc_h6;
        let _369_cf1 = _364___mcc_h5;
        (_324_globalState).f4 = _320_v0;
        let _370_v30;
        _370_v30 = _dafny.Set.fromElements(_321_v1);
        let _371_v31;
        let _nw56 = new _module.C0();
        _nw56.__ctor(_320_v0, _370_v30);
        _371_v31 = _nw56;
        let _372_v32;
        let _nw57 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _372_v32 = _nw57;
        let _373_v33;
        _373_v33 = _dafny.Map.Empty.slice().updateUnsafe(_371_v31,_372_v32);
        _373_v33 = (_dafny.Map.Empty.slice().updateUnsafe(_371_v31,_372_v32)).Merge(_373_v33);
        let _374_v34;
        _374_v34 = new _dafny.CodePoint('t'.codePointAt(0));
        let _375_v35;
        let _nw58 = Array((new BigNumber(16)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = _372_v32;
        _nw58[(_dafny.ONE).toNumber()] = _372_v32;
        _nw58[(new BigNumber(2)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(3)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(4)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(5)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(6)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(7)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(8)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(9)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(10)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(11)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(12)).toNumber()] = ((_321_v1) ? (_372_v32) : (_372_v32));
        _nw58[(new BigNumber(13)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(14)).toNumber()] = _372_v32;
        _nw58[(new BigNumber(15)).toNumber()] = _372_v32;
        _375_v35 = _nw58;
        let _index34 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_375_v35).length));
        (_375_v35)[_index34] = _372_v32;
        let _376_v36;
        _376_v36 = _dafny.Map.Empty.slice().updateUnsafe(_372_v32,_367_cf3);
        let _377_v37;
        _377_v37 = _module.D0.create_DC1(_369_cf1, _368_cf2, (((_376_v36).contains(_372_v32)) ? ((_376_v36).get(_372_v32)) : (_367_cf3)));
        let _378_v38;
        _378_v38 = _dafny.Seq.of(_368_cf2);
        let _379_v39;
        let _nw59 = Array((new BigNumber(28)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _368_cf2;
        _nw59[(_dafny.ONE).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(2)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(3)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(4)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(5)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(6)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(7)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(8)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(9)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(10)).toNumber()] = ((_321_v1) ? (_368_cf2) : (_368_cf2));
        _nw59[(new BigNumber(11)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(12)).toNumber()] = (_377_v37).dtor_cf2;
        _nw59[(new BigNumber(13)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(14)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(15)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(16)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(17)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(18)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(19)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(20)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(21)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(22)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(23)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(24)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(25)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(26)).toNumber()] = _368_cf2;
        _nw59[(new BigNumber(27)).toNumber()] = (_378_v38)[_module.__default.safeIndex(_320_v0, new BigNumber((_378_v38).length))];
        _379_v39 = _nw59;
        let _380_v40;
        _380_v40 = _dafny.MultiSet.fromElements(false, _321_v1);
        let _index35 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_375_v35).length));
        let _rhs32 = _374_v34;
        let _rhs33 = _372_v32;
        let _rhs34 = _321_v1;
        let _rhs35 = _379_v39;
        let _rhs36 = new BigNumber((((!(_321_v1)) ? (_380_v40) : (((_321_v1) ? ((_380_v40).update(true, _module.__default.abs(new BigNumber(-722)))) : (_380_v40))))).cardinality());
        let _lhs24 = _375_v35;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_375_v35).length));
        _374_v34 = _rhs32;
        _lhs24[_lhs25] = _rhs33;
        _321_v1 = _rhs34;
        _379_v39 = _rhs35;
        _320_v0 = _rhs36;
        let _381_v42;
        _381_v42 = _dafny.Set.fromElements((_dafny.ZERO).minus((_371_v31).f6));
        _321_v1 = ((function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(-485), new BigNumber(640))) {
            let _382_v41 = _compr_21;
            if (((new BigNumber(-485)).isLessThanOrEqualTo(_382_v41)) && ((_382_v41).isLessThan(new BigNumber(640)))) {
              _coll21.add((_382_v41).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_374_v34,_321_v1)).length)));
            }
          }
          return _coll21;
        }()).Intersect(_381_v42)).IsProperSubsetOf(_381_v42);
      }
      let _383_v43;
      _383_v43 = _module.D1.create_DC3(_321_v1, _321_v1, _352_v20, _321_v1, _320_v0);
      let _pat_let_tv14 = _321_v1;
      let _pat_let_tv15 = _321_v1;
      let _pat_let_tv16 = _321_v1;
      let _pat_let_tv17 = _321_v1;
      if (function (_source9) {
        if (_source9.is_DC3) {
          let _384___mcc_h8 = (_source9).cf5;
          let _385___mcc_h9 = (_source9).cf6;
          let _386___mcc_h10 = (_source9).cf7;
          let _387___mcc_h11 = (_source9).cf8;
          let _388___mcc_h12 = (_source9).cf9;
          let _389_cf9 = _388___mcc_h12;
          let _390_cf8 = _387___mcc_h11;
          let _391_cf7 = _386___mcc_h10;
          let _392_cf6 = _385___mcc_h9;
          let _393_cf5 = _384___mcc_h8;
          return _pat_let_tv14;
        } else if (_source9.is_DC2) {
          let _394___mcc_h13 = (_source9).cf4;
          let _395_cf4 = _394___mcc_h13;
          return _pat_let_tv15;
        } else {
          let _396___mcc_h14 = (_source9).cf10;
          let _397_cf10 = _396___mcc_h14;
          if (_pat_let_tv16) {
            return true;
          } else {
            return _pat_let_tv17;
          }
        }
      }(_383_v43)) {
        (_324_globalState).f4 = _320_v0;
        let _398_v44;
        _398_v44 = _dafny.MultiSet.fromElements(_321_v1, _321_v1, _321_v1);
        let _399_v45;
        _399_v45 = _dafny.Map.Empty.slice().updateUnsafe(_352_v20,_321_v1);
        let _400_v46;
        let _nw60 = new _module.C0();
        _nw60.__ctor(new BigNumber((_398_v44).cardinality()), _dafny.Set.fromElements(!((((_399_v45).contains(_352_v20)) ? ((_399_v45).get(_352_v20)) : (_321_v1))), _321_v1));
        _400_v46 = _nw60;
        let _401_v47;
        let _nw61 = Array((new BigNumber(24)).toNumber()).fill(false);
        _401_v47 = _nw61;
        _401_v47 = _401_v47;
        (_324_globalState).f4 = _320_v0;
        let _402_v48;
        _402_v48 = new _dafny.CodePoint('b'.codePointAt(0));
        _399_v45 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_352_v20, _module.__default.safeIndex(_320_v0, new BigNumber((_352_v20).length)), _402_v48),false)).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(896)), ((_403_v48) => function (_404_i2) {
          return _403_v48;
        })(_402_v48)), _321_v1);
      } else {
        let _405_v49;
        _405_v49 = _dafny.MultiSet.fromElements(_321_v1, _321_v1);
        if ((new BigNumber((((_321_v1) ? (_405_v49) : (_405_v49))).cardinality())).isLessThanOrEqualTo(_320_v0)) {
          _321_v1 = !(_321_v1);
          _321_v1 = _321_v1;
          _321_v1 = false;
          let _406_v50;
          let _nw62 = Array((new BigNumber(3)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _320_v0;
          _nw62[(_dafny.ONE).toNumber()] = (new BigNumber(274)).minus((_383_v43).dtor_cf9);
          _nw62[(new BigNumber(2)).toNumber()] = (_320_v0).minus(_320_v0);
          _406_v50 = _nw62;
          let _index36 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_406_v50).length));
          (_406_v50)[_index36] = _320_v0;
          let _407_v51;
          _407_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_320_v0, _321_v1, _324_globalState),_321_v1);
          let _index37 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_406_v50).length));
          (_406_v50)[_index37] = new BigNumber((_407_v51).length);
          let _408_v52;
          let _nw63 = Array((new BigNumber(16)).toNumber()).fill(false);
          _408_v52 = _nw63;
          let _409_v53;
          _409_v53 = _dafny.Map.Empty.slice().updateUnsafe(_405_v49,_408_v52);
          (_324_globalState).f0 = new BigNumber((_409_v53).length);
        } else {
          _321_v1 = false;
          let _410_v54;
          let _nw64 = Array((new BigNumber(7)).toNumber());
          _nw64[(_dafny.ZERO).toNumber()] = (_320_v0).isLessThanOrEqualTo(_320_v0);
          _nw64[(_dafny.ONE).toNumber()] = _321_v1;
          _nw64[(new BigNumber(2)).toNumber()] = _321_v1;
          _nw64[(new BigNumber(3)).toNumber()] = _321_v1;
          _nw64[(new BigNumber(4)).toNumber()] = _321_v1;
          _nw64[(new BigNumber(5)).toNumber()] = _321_v1;
          _nw64[(new BigNumber(6)).toNumber()] = _321_v1;
          _410_v54 = _nw64;
          _410_v54 = ((false) ? (_410_v54) : (_410_v54));
          let _411_v55;
          _411_v55 = _dafny.Seq.of(_320_v0);
          let _412_v56;
          _412_v56 = new _dafny.CodePoint('f'.codePointAt(0));
          let _413_v57;
          _413_v57 = _dafny.Set.fromElements(_412_v56, _412_v56, _412_v56);
          let _414_v58;
          _414_v58 = _dafny.Seq.of(new BigNumber((_411_v55).length), _320_v0, _320_v0, _320_v0, new BigNumber((_413_v57).length));
          let _415_v59;
          let _nw65 = new _module.C0();
          _nw65.__ctor((_411_v55)[_module.__default.safeIndex(_module.__default.fm4(_320_v0, _321_v1, _324_globalState), new BigNumber((_411_v55).length))], _dafny.Set.fromElements(!(_321_v1)));
          _415_v59 = _nw65;
          (_324_globalState).f0 = _320_v0;
          (_324_globalState).f4 = (_320_v0).multipliedBy((_383_v43).dtor_cf9);
        }
        let _416_v60;
        _416_v60 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,(_320_v0).plus(_320_v0));
        let _417_v61;
        _417_v61 = _dafny.Set.fromElements((_dafny.ZERO).minus(_320_v0), _320_v0);
        _416_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_417_v61).length),_320_v0);
        let _418_v62;
        _418_v62 = _dafny.Set.fromElements(!(_321_v1));
        let _419_v63;
        let _nw66 = new _module.C0();
        _nw66.__ctor(_module.__default.safeModuloInt(_320_v0, _320_v0), ((_321_v1) ? (_418_v62) : (_418_v62)));
        _419_v63 = _nw66;
        let _420_v64;
        _420_v64 = new _dafny.CodePoint('u'.codePointAt(0));
        (_324_globalState).f0 = ((!_dafny.Seq.contains(_352_v20, _420_v64)) ? (new BigNumber(177)) : (((_419_v63).f6).minus((_419_v63).f6)));
        if (true) {
          let _421_v65;
          let _out1;
          _out1 = _module.__default.m0(_321_v1, _324_globalState);
          _421_v65 = _out1;
          _321_v1 = _421_v65;
          let _422_v66;
          let _nw67 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _422_v66 = _nw67;
          let _423_v67;
          _423_v67 = _dafny.Seq.of(_422_v66, _422_v66, _422_v66);
          let _pat_let_tv18 = _422_v66;
          let _424_v68;
          let _nw68 = Array((new BigNumber(22)).toNumber());
          _nw68[(_dafny.ZERO).toNumber()] = _422_v66;
          _nw68[(_dafny.ONE).toNumber()] = _422_v66;
          _nw68[(new BigNumber(2)).toNumber()] = (_423_v67)[_module.__default.safeIndex((_419_v63).f6, new BigNumber((_423_v67).length))];
          _nw68[(new BigNumber(3)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(4)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(5)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(6)).toNumber()] = (_423_v67)[_module.__default.safeIndex(_module.__default.fm4((_419_v63).f6, _421_v65, _324_globalState), new BigNumber((_423_v67).length))];
          _nw68[(new BigNumber(7)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(8)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(9)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(10)).toNumber()] = (_423_v67)[_module.__default.safeIndex((_419_v63).f6, new BigNumber((_423_v67).length))];
          _nw68[(new BigNumber(11)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(12)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(13)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(14)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(15)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(16)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(17)).toNumber()] = (function (_pat_let3_0) {
            return function (_425_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_426_dt__update_hcf11_h0) {
                  return _module.D2.create_DC5(_426_dt__update_hcf11_h0);
                }(_pat_let4_0);
              }(_pat_let_tv18);
            }(_pat_let3_0);
          }(_module.D2.create_DC5(_422_v66))).dtor_cf11;
          _nw68[(new BigNumber(18)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(19)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(20)).toNumber()] = _422_v66;
          _nw68[(new BigNumber(21)).toNumber()] = _422_v66;
          _424_v68 = _nw68;
          let _index38 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_424_v68).length));
          (_424_v68)[_index38] = _422_v66;
          let _index39 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_424_v68).length));
          (_424_v68)[_index39] = _422_v66;
          let _427_v69;
          let _nw69 = Array((new BigNumber(28)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _421_v65;
          _nw69[(_dafny.ONE).toNumber()] = _module.__default.fm0(_321_v1, (_module.D1.create_DC3((_383_v43).dtor_cf6, _421_v65, _dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), ((_428_v64) => function (_429_i3) {
  return _428_v64;
})(_420_v64)), _421_v65, new BigNumber(639))).dtor_cf5, (_419_v63).f6, _324_globalState);
          _nw69[(new BigNumber(2)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(3)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(4)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(5)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(6)).toNumber()] = false;
          _nw69[(new BigNumber(7)).toNumber()] = (_383_v43).dtor_cf5;
          _nw69[(new BigNumber(8)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(9)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(10)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(11)).toNumber()] = !(_421_v65);
          _nw69[(new BigNumber(12)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(13)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(14)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(15)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(16)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(17)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(18)).toNumber()] = true;
          _nw69[(new BigNumber(19)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(20)).toNumber()] = _321_v1;
          _nw69[(new BigNumber(21)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(22)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(23)).toNumber()] = true;
          _nw69[(new BigNumber(24)).toNumber()] = false;
          _nw69[(new BigNumber(25)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(26)).toNumber()] = _421_v65;
          _nw69[(new BigNumber(27)).toNumber()] = _321_v1;
          _427_v69 = _nw69;
          let _430_v70;
          _430_v70 = _dafny.Map.Empty.slice().updateUnsafe(_427_v69,_421_v65);
          let _431_v71;
          let _nw70 = new _module.C0();
          _nw70.__ctor(((((_405_v49).contains(_421_v65)) ? ((_405_v49).get(_421_v65)) : ((_419_v63).f6))).multipliedBy(new BigNumber((_430_v70).length)), _418_v62);
          _431_v71 = _nw70;
          let _432_v72;
          _432_v72 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_353_v21).cardinality()));
          (_324_globalState).f4 = new BigNumber((_432_v72).length);
        } else {
          let _433_v73;
          let _nw71 = new _module.C0();
          _nw71.__ctor(new BigNumber(483), (_419_v63).f7);
          _433_v73 = _nw71;
          let _434_v74;
          let _nw72 = Array((new BigNumber(19)).toNumber()).fill(false);
          _434_v74 = _nw72;
          let _435_v75;
          _435_v75 = _dafny.Map.Empty.slice().updateUnsafe(_419_v63,_434_v74);
          let _436_v76;
          _436_v76 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,_434_v74);
          _435_v75 = (_435_v75).update(_433_v73, (((_436_v76).contains((_419_v63).f6)) ? ((_436_v76).get((_419_v63).f6)) : ((((_436_v76).contains(_320_v0)) ? ((_436_v76).get(_320_v0)) : (_434_v74)))));
          let _437_v77;
          let _init7 = ((_438_v63) => function (_439_i4) {
            return _module.__default.safeModuloInt(_439_i4, (_438_v63).f6);
          })(_419_v63);
          let _nw73 = Array((_dafny.ONE).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw73.length); _i0_7++) {
            _nw73[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _437_v77 = _nw73;
          let _index40 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_437_v77).length));
          (_437_v77)[_index40] = (_419_v63).f6;
          let _index41 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_437_v77).length));
          (_437_v77)[_index41] = _320_v0;
          let _440_v78;
          _440_v78 = _dafny.Map.Empty.slice().updateUnsafe((_417_v61).Union(_dafny.Set.fromElements(new BigNumber(858))),_dafny.Seq.update(_352_v20, _module.__default.safeIndex((_433_v73).f6, new BigNumber((_352_v20).length)), _420_v64));
          let _pat_let_tv19 = _321_v1;
          let _pat_let_tv20 = _321_v1;
          let _index42 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_437_v77).length));
          let _index43 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_437_v77).length));
          let _rhs37 = (((_440_v78).contains((_417_v61).Union(_417_v61))) ? ((_440_v78).get((_417_v61).Union(_417_v61))) : ((function (_pat_let5_0) {
            return function (_441_dt__update__tmp_h1) {
              return function (_pat_let6_0) {
                return function (_442_dt__update_hcf6_h0) {
                  return function (_pat_let7_0) {
                    return function (_443_dt__update_hcf8_h0) {
                      return _module.D1.create_DC3((_441_dt__update__tmp_h1).dtor_cf5, _442_dt__update_hcf6_h0, (_441_dt__update__tmp_h1).dtor_cf7, _443_dt__update_hcf8_h0, (_441_dt__update__tmp_h1).dtor_cf9);
                    }(_pat_let7_0);
                  }(_pat_let_tv20);
                }(_pat_let6_0);
              }(_pat_let_tv19);
            }(_pat_let5_0);
          }(_383_v43)).dtor_cf7));
          let _rhs38 = (_419_v63).f6;
          let _rhs39 = _module.__default.fm4(_320_v0, _dafny.areEqual(_352_v20, _352_v20), _324_globalState);
          let _rhs40 = ((_419_v63).f6).plus(_320_v0);
          let _rhs41 = (_419_v63).f6;
          let _lhs26 = _437_v77;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_437_v77).length));
          let _lhs28 = _324_globalState;
          let _lhs29 = _437_v77;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_437_v77).length));
          let _lhs31 = _324_globalState;
          _352_v20 = _rhs37;
          _lhs26[_lhs27] = _rhs38;
          _lhs28.f4 = _rhs39;
          _lhs29[_lhs30] = _rhs40;
          _lhs31.f4 = _rhs41;
          let _444_v79;
          _444_v79 = _dafny.Set.fromElements(_420_v64);
          (_324_globalState).f4 = ((_437_v77)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_437_v77).length))]).multipliedBy(new BigNumber(((_444_v79).Intersect(_444_v79)).length));
          let _445_v80;
          _445_v80 = _dafny.Seq.of((_419_v63).f7, _dafny.Set.fromElements(_321_v1, _321_v1, _321_v1));
          let _446_v81;
          _446_v81 = _dafny.Map.Empty.slice().updateUnsafe(_321_v1,_320_v0);
          let _447_v82;
          _447_v82 = _dafny.Seq.of(new BigNumber(-799));
          let _448_v83;
          let _nw74 = new _module.C0();
          _nw74.__ctor(new BigNumber(-604), (_445_v80)[_module.__default.safeIndex((_dafny.ZERO).minus((((_446_v81).contains(_321_v1)) ? ((_446_v81).get(_321_v1)) : (new BigNumber((_447_v82).length)))), new BigNumber((_445_v80).length))]);
          _448_v83 = _nw74;
        }
      }
      let _449_v84;
      _449_v84 = _dafny.Map.Empty.slice().updateUnsafe(_321_v1,new BigNumber((_352_v20).length));
      let _450_v85;
      _450_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(492),_449_v84);
      let _451_v86;
      _451_v86 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_450_v85).length),_449_v84);
      let _452_v87;
      _452_v87 = _dafny.MultiSet.fromElements(_321_v1, _321_v1, false, !(new BigNumber((_450_v85).length)).isEqualTo((_dafny.ZERO).minus(_320_v0)), !(_321_v1) || (_321_v1));
      _452_v87 = (_452_v87).Difference(_452_v87);
      _321_v1 = _321_v1;
      let _453_v88;
      let _nw75 = Array((new BigNumber(7)).toNumber());
      _453_v88 = _nw75;
      let _454_v89;
      _454_v89 = _dafny.Set.fromElements(_321_v1);
      let _rhs42 = _module.__default.safeModuloInt(new BigNumber((_454_v89).length), (_320_v0).multipliedBy(new BigNumber(578)));
      let _rhs43 = false;
      let _rhs44 = _453_v88;
      let _rhs45 = !(!(_321_v1));
      let _rhs46 = _321_v1;
      let _lhs32 = _324_globalState;
      _lhs32.f0 = _rhs42;
      _321_v1 = _rhs43;
      _453_v88 = _rhs44;
      _321_v1 = _rhs45;
      _321_v1 = _rhs46;
      let _455_v90;
      let _out2;
      _out2 = _module.__default.m0(_321_v1, _324_globalState);
      _455_v90 = _out2;
      let _456_v91;
      _456_v91 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,true);
      let _457_v92;
      _457_v92 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),_320_v0);
      let _458_v93;
      _458_v93 = _module.D1.create_DC2(_320_v0);
      let _459_v94;
      let _nw76 = Array((new BigNumber(8)).toNumber());
      _nw76[(_dafny.ZERO).toNumber()] = _452_v87;
      _nw76[(_dafny.ONE).toNumber()] = (_452_v87).Difference(_452_v87);
      _nw76[(new BigNumber(2)).toNumber()] = _452_v87;
      _nw76[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_module.__default.fm0(!(_321_v1), (((_456_v91).contains((_383_v43).dtor_cf9)) ? ((_456_v91).get((_383_v43).dtor_cf9)) : (_455_v90)), new BigNumber((_457_v92).length), _324_globalState), _321_v1, _321_v1, _321_v1, _321_v1);
      _nw76[(new BigNumber(4)).toNumber()] = (_452_v87).Difference(_452_v87);
      _nw76[(new BigNumber(5)).toNumber()] = (_452_v87).update(_321_v1, _module.__default.abs((_458_v93).dtor_cf4));
      _nw76[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_455_v90, _455_v90);
      _nw76[(new BigNumber(7)).toNumber()] = _452_v87;
      _459_v94 = _nw76;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_459_v94).length))) {
        let _460_i5 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_460_i5)) && ((_460_i5).isLessThan(new BigNumber((_459_v94).length))))) {
          (_459_v94)[(_460_i5)] = _452_v87;
        }
      }
      _352_v20 = _dafny.Seq.Concat(_352_v20, _352_v20);
      let _source10 = _458_v93;
      if (_source10.is_DC3) {
        let _461___mcc_h15 = (_source10).cf5;
        let _462___mcc_h16 = (_source10).cf6;
        let _463___mcc_h17 = (_source10).cf7;
        let _464___mcc_h18 = (_source10).cf8;
        let _465___mcc_h19 = (_source10).cf9;
        let _466_cf9 = _465___mcc_h19;
        let _467_cf8 = _464___mcc_h18;
        let _468_cf7 = _463___mcc_h17;
        let _469_cf6 = _462___mcc_h16;
        let _470_cf5 = _461___mcc_h15;
        let _471_v95;
        _471_v95 = _dafny.Map.Empty.slice().updateUnsafe(_469_cf6,false);
        _471_v95 = (_471_v95).update(!(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(954), new BigNumber(-549))) {
            let _472_v96 = _compr_22;
            if (((new BigNumber(954)).isLessThanOrEqualTo(_472_v96)) && ((_472_v96).isLessThan(new BigNumber(-549)))) {
              _coll22.push([_module.__default.safeModuloInt(_472_v96, new BigNumber((_471_v95).length)),_470_cf5]);
            }
          }
          return _coll22;
        }()).equals(_456_v91), _470_cf5);
        let _473_v97;
        _473_v97 = new _dafny.CodePoint('r'.codePointAt(0));
        _473_v97 = _473_v97;
        (_324_globalState).f4 = (_320_v0).plus(new BigNumber((_468_cf7).length));
        _469_cf6 = (((_471_v95).contains((_467_cf8) === (_467_cf8))) ? ((_471_v95).get((_467_cf8) === (_467_cf8))) : (true));
      } else if (_source10.is_DC2) {
        let _474___mcc_h20 = (_source10).cf4;
        let _475_cf4 = _474___mcc_h20;
        _454_v89 = _module.__default.fm5(_324_globalState);
        let _476_v98;
        _476_v98 = _dafny.Seq.of(new BigNumber((_322_v2).length));
        if (!(!((_476_v98)[_module.__default.safeIndex(_320_v0, new BigNumber((_476_v98).length))]).isEqualTo(_475_cf4))) {
          let _477_v99;
          _477_v99 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-19)), function (_478_i6) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }),_322_v2);
          let _479_v101;
          _479_v101 = new _dafny.CodePoint('d'.codePointAt(0));
          let _480_v102;
          _480_v102 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,_479_v101);
          let _481_v103;
          _481_v103 = _dafny.Seq.of(_477_v99, _477_v99, _477_v99, function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_dafny.MultiSet.fromElements(_352_v20, _352_v20)).Elements) {
              let _482_v100 = _compr_23;
              if ((_dafny.MultiSet.fromElements(_352_v20, _352_v20)).contains(_482_v100)) {
                _coll23.push([_482_v100,_dafny.Seq.of(true)]);
              }
            }
            return _coll23;
          }(), _dafny.Map.Empty.slice().updateUnsafe(_352_v20,_dafny.Seq.update(_322_v2, _module.__default.safeIndex(new BigNumber((_480_v102).length), new BigNumber((_322_v2).length)), _321_v1)));
          let _483_v105;
          _483_v105 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,((_455_v90) ? (((_481_v103)[_module.__default.safeIndex(_320_v0, new BigNumber((_481_v103).length))]).update(_352_v20, _322_v2)) : (function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of (_dafny.Map.Empty.slice().updateUnsafe(_352_v20,_321_v1)).Keys.Elements) {
              let _484_v104 = _compr_24;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_352_v20,_321_v1)).contains(_484_v104)) {
                _coll24.push([_484_v104,_322_v2]);
              }
            }
            return _coll24;
          }())));
          _483_v105 = (_483_v105).update(_475_cf4, ((_477_v99).update(_dafny.Seq.UnicodeFromString("kb"), _322_v2)).Merge(_477_v99));
          _321_v1 = !(_321_v1) || (_321_v1);
          let _485_v106;
          _485_v106 = _dafny.Seq.of(_449_v84);
          let _486_v107;
          _486_v107 = _dafny.Seq.of((_485_v106)[_module.__default.safeIndex(_320_v0, new BigNumber((_485_v106).length))], _449_v84, _449_v84, _449_v84);
          let _487_v108;
          let _nw77 = new _module.C0();
          _nw77.__ctor(new BigNumber((((_486_v107)[_module.__default.safeIndex(new BigNumber((_352_v20).length), new BigNumber((_486_v107).length))]).Merge(_449_v84)).length), _454_v89);
          _487_v108 = _nw77;
          (_324_globalState).f4 = _module.__default.safeDivisionInt(new BigNumber((_352_v20).length), _320_v0);
          let _488_v109;
          _488_v109 = _module.D2.create_DC7((_475_cf4).multipliedBy((_487_v108).f6));
          let _pat_let_tv21 = _487_v108;
          _488_v109 = function (_pat_let8_0) {
            return function (_489_dt__update__tmp_h2) {
              return function (_pat_let9_0) {
                return function (_490_dt__update_hcf15_h0) {
                  return _module.D2.create_DC7(_490_dt__update_hcf15_h0);
                }(_pat_let9_0);
              }((_pat_let_tv21).f6);
            }(_pat_let8_0);
          }(_488_v109);
        } else {
          let _491_v110;
          _491_v110 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-884)), function (_492_i7) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })).length), _320_v0);
          let _493_v111;
          _493_v111 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,(_491_v110).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_455_v90)).length), _320_v0)));
          _493_v111 = (_493_v111).update(_320_v0, _491_v110);
          let _494_v112;
          let _nw78 = Array((new BigNumber(2)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = _320_v0;
          _nw78[(_dafny.ONE).toNumber()] = _475_cf4;
          _494_v112 = _nw78;
          let _index44 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_494_v112).length));
          (_494_v112)[_index44] = (new BigNumber(301)).plus(new BigNumber((_352_v20).length));
          let _495_v113;
          _495_v113 = _dafny.Seq.of(_module.D2.create_DC5(_494_v112));
          let _index45 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_494_v112).length));
          let _rhs47 = (_dafny.ZERO).minus((_module.D2.create_DC6(_320_v0, _321_v1, _494_v112)).dtor_cf12);
          let _rhs48 = _dafny.Seq.of(_module.D2.create_DC5(_494_v112));
          let _rhs49 = _module.__default.safeDivisionInt(new BigNumber((_454_v89).length), _320_v0);
          let _lhs33 = _494_v112;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_494_v112).length));
          _lhs33[_lhs34] = _rhs47;
          _495_v113 = _rhs48;
          _320_v0 = _rhs49;
          let _496_v114;
          let _nw79 = new _module.C0();
          _nw79.__ctor((_320_v0).minus(_475_cf4), _module.__default.fm5(_324_globalState));
          _496_v114 = _nw79;
          let _497_v115;
          let _out3;
          _out3 = _module.__default.m0(_455_v90, _324_globalState);
          _497_v115 = _out3;
          let _498_v116;
          _498_v116 = _dafny.Map.Empty.slice().updateUnsafe(_496_v114,_321_v1);
          _497_v115 = !((((_498_v116).contains(((false) ? (_496_v114) : (_496_v114)))) ? ((_498_v116).get(((false) ? (_496_v114) : (_496_v114)))) : (_497_v115)));
        }
        let _499_v117;
        let _init8 = ((_500_v0) => function (_501_i8) {
          return _module.__default.safeModuloInt(_501_i8, _500_v0);
        })(_320_v0);
        let _nw80 = Array((new BigNumber(8)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw80.length); _i0_8++) {
          _nw80[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _499_v117 = _nw80;
        let _502_v118;
        _502_v118 = _module.D2.create_DC5(_499_v117);
        let _503_v119;
        _503_v119 = _dafny.Map.Empty.slice().updateUnsafe(_499_v117,((_455_v90) ? (_502_v118) : (_502_v118)));
        _503_v119 = (_503_v119).update(_499_v117, _502_v118);
        (_324_globalState).f0 = _module.__default.safeModuloInt(_475_cf4, (_dafny.ZERO).minus(_320_v0));
      } else {
        let _504___mcc_h21 = (_source10).cf10;
        let _505_cf10 = _504___mcc_h21;
        let _506_v120;
        _506_v120 = new _dafny.CodePoint('f'.codePointAt(0));
        let _507_v121;
        _507_v121 = _dafny.MultiSet.fromElements(_506_v120);
        let _508_v122;
        let _nw81 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _508_v122 = _nw81;
        let _509_v123;
        _509_v123 = _dafny.Map.Empty.slice().updateUnsafe(_507_v121,_508_v122);
        _509_v123 = (_509_v123).update(_507_v121, _508_v122);
        (_324_globalState).f0 = (_dafny.ZERO).minus(_320_v0);
        _321_v1 = (_320_v0).isLessThanOrEqualTo(_module.__default.safeModuloInt(_320_v0, _320_v0));
        let _510_v124;
        _510_v124 = _dafny.Set.fromElements(_320_v0, _320_v0, _320_v0);
        (_324_globalState).f4 = (_dafny.ZERO).minus(new BigNumber(((_510_v124).Intersect(_510_v124)).length));
      }
      let _511_v125;
      _511_v125 = _dafny.Seq.of(_320_v0, _320_v0);
      let _512_v126;
      let _nw82 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _512_v126 = _nw82;
      let _513_v127;
      _513_v127 = _module.D2.create_DC6(new BigNumber((_module.__default.fm3(_322_v2, _324_globalState)).length), _dafny.areEqual(_511_v125, _511_v125), _512_v126);
      let _source11 = _513_v127;
      if (_source11.is_DC6) {
        let _514___mcc_h22 = (_source11).cf12;
        let _515___mcc_h23 = (_source11).cf13;
        let _516___mcc_h24 = (_source11).cf14;
        let _517_cf14 = _516___mcc_h24;
        let _518_cf13 = _515___mcc_h23;
        let _519_cf12 = _514___mcc_h22;
        _320_v0 = ((new BigNumber(761)).multipliedBy(_519_cf12)).minus(_519_cf12);
        _321_v1 = _455_v90;
        let _520_v128;
        _520_v128 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_519_cf12, _518_cf13, _324_globalState),_module.__default.safeModuloInt(_320_v0, new BigNumber((_511_v125).length)));
        let _rhs50 = _455_v90;
        let _rhs51 = false;
        let _rhs52 = (_520_v128).Merge(_dafny.Map.Empty.slice().updateUnsafe(_519_cf12,_519_cf12));
        _455_v90 = _rhs50;
        _455_v90 = _rhs51;
        _520_v128 = _rhs52;
        _455_v90 = _321_v1;
      } else if (_source11.is_DC7) {
        let _521___mcc_h25 = (_source11).cf15;
        let _522_cf15 = _521___mcc_h25;
        let _523_v129;
        let _out4;
        _out4 = _module.__default.m0(((_321_v1) ? ((((_456_v91).contains(_522_cf15)) ? ((_456_v91).get(_522_cf15)) : (!(_455_v90)))) : (_321_v1)), _324_globalState);
        _523_v129 = _out4;
        let _524_v130;
        let _nw83 = new _module.C0();
        _nw83.__ctor(_320_v0, _454_v89);
        _524_v130 = _nw83;
        let _index46 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_453_v88).length));
        (_453_v88)[_index46] = _524_v130;
        let _index47 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_453_v88).length));
        let _rhs53 = _524_v130;
        let _rhs54 = _320_v0;
        let _rhs55 = _511_v125;
        let _rhs56 = _524_v130;
        let _lhs35 = _453_v88;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_453_v88).length));
        let _lhs37 = _324_globalState;
        _lhs35[_lhs36] = _rhs53;
        _lhs37.f0 = _rhs54;
        _511_v125 = _rhs55;
        _524_v130 = _rhs56;
        let _525_v131;
        _525_v131 = new _dafny.CodePoint('a'.codePointAt(0));
        _525_v131 = _525_v131;
        let _526_v132;
        _526_v132 = _module.D2.create_DC7(_522_cf15);
        let _527_v133;
        _527_v133 = _dafny.Set.fromElements(_522_cf15);
        let _528_v134;
        let _nw84 = Array((new BigNumber(28)).toNumber());
        _nw84[(_dafny.ZERO).toNumber()] = _526_v132;
        _nw84[(_dafny.ONE).toNumber()] = _526_v132;
        _nw84[(new BigNumber(2)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(3)).toNumber()] = ((_module.__default.fm0(_455_v90, _321_v1, new BigNumber((_527_v133).length), _324_globalState)) ? (_526_v132) : (_module.D2.create_DC7(new BigNumber((_352_v20).length))));
        _nw84[(new BigNumber(4)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(5)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(6)).toNumber()] = _module.D2.create_DC7(new BigNumber(68));
        _nw84[(new BigNumber(7)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(8)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(9)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(10)).toNumber()] = _module.D2.create_DC7(_320_v0);
        _nw84[(new BigNumber(11)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(12)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(13)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(14)).toNumber()] = ((false) ? (_526_v132) : (_526_v132));
        _nw84[(new BigNumber(15)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(16)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(17)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(18)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(19)).toNumber()] = _module.D2.create_DC7(new BigNumber((_322_v2).length));
        _nw84[(new BigNumber(20)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(21)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(22)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(23)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(24)).toNumber()] = ((_455_v90) ? (_526_v132) : (_526_v132));
        _nw84[(new BigNumber(25)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(26)).toNumber()] = _526_v132;
        _nw84[(new BigNumber(27)).toNumber()] = ((_321_v1) ? (_526_v132) : (_526_v132));
        _528_v134 = _nw84;
        let _index48 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_528_v134).length));
        (_528_v134)[_index48] = _526_v132;
        let _index49 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_528_v134).length));
        (_528_v134)[_index49] = _module.__default.fm6(_324_globalState);
      } else if (_source11.is_DC5) {
        let _529___mcc_h26 = (_source11).cf11;
        let _530_cf11 = _529___mcc_h26;
        if (((_321_v1) ? (_455_v90) : (_321_v1))) {
          _455_v90 = ((_dafny.ZERO).minus(_module.__default.fm4(new BigNumber(-120), _321_v1, _324_globalState))).isLessThanOrEqualTo((_dafny.ZERO).minus(_320_v0));
          _452_v87 = (_dafny.MultiSet.fromElements(_321_v1, _455_v90, _321_v1, _321_v1, _455_v90)).update(!(!(_455_v90)), _module.__default.abs(_320_v0));
          (_324_globalState).f4 = (_511_v125)[_module.__default.safeIndex(_320_v0, new BigNumber((_511_v125).length))];
          let _531_v135;
          let _init9 = ((_532_v20) => function (_533_i9) {
            return _532_v20;
          })(_352_v20);
          let _nw85 = Array((new BigNumber(26)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw85.length); _i0_9++) {
            _nw85[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _531_v135 = _nw85;
          let _index50 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_531_v135).length));
          (_531_v135)[_index50] = _dafny.Seq.Concat(_352_v20, _dafny.Seq.UnicodeFromString("lkypk"));
          let _index51 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_531_v135).length));
          (_531_v135)[_index51] = _352_v20;
          _449_v84 = (_449_v84).update(true, _320_v0);
        } else {
          let _534_v136;
          let _out5;
          _out5 = _module.__default.m0(!((_dafny.ZERO).minus(_320_v0)).isEqualTo(new BigNumber((_511_v125).length)), _324_globalState);
          _534_v136 = _out5;
          let _535_v137;
          let _out6;
          _out6 = _module.__default.m0(_321_v1, _324_globalState);
          _535_v137 = _out6;
          _534_v136 = !_dafny.areEqual(_352_v20, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qwjyx"), _352_v20));
          (_324_globalState).f0 = _320_v0;
          let _536_v138;
          let _nw86 = new _module.C0();
          _nw86.__ctor(_320_v0, _454_v89);
          _536_v138 = _nw86;
        }
        (_324_globalState).f4 = new BigNumber(-260);
        let _537_v139;
        let _nw87 = Array((new BigNumber(21)).toNumber()).fill([]);
        _537_v139 = _nw87;
        _537_v139 = _537_v139;
        let _538_v140;
        _538_v140 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,_320_v0);
        (_324_globalState).f0 = (((_538_v140).contains((_dafny.ZERO).minus(_320_v0))) ? ((_538_v140).get((_dafny.ZERO).minus(_320_v0))) : (_320_v0));
      } else {
        let _539___mcc_h27 = (_source11).cf16;
        let _540_cf16 = _539___mcc_h27;
        let _541_v141;
        _541_v141 = _dafny.Map.Empty.slice().updateUnsafe(_455_v90,_321_v1);
        let _542_v142;
        let _nw88 = new _module.C0();
        _nw88.__ctor(_320_v0, _454_v89);
        _542_v142 = _nw88;
        let _543_v143;
        _543_v143 = _dafny.Map.Empty.slice().updateUnsafe(_542_v142,(_542_v142).f6);
        if ((new BigNumber(((_543_v143).Merge(_dafny.Map.Empty.slice().updateUnsafe(_542_v142,(_542_v142).f6))).length)).isLessThan((_dafny.ZERO).minus(((((_353_v21).contains(new BigNumber((_322_v2).length))) ? ((_353_v21).get(new BigNumber((_322_v2).length))) : (new BigNumber((_541_v141).length)))).multipliedBy(_320_v0)))) {
          let _544_v144;
          _544_v144 = _dafny.MultiSet.fromElements(_512_v126);
          let _545_v145;
          let _init10 = function (_546_i10) {
            return true;
          };
          let _nw89 = Array((new BigNumber(21)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw89.length); _i0_10++) {
            _nw89[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _545_v145 = _nw89;
          let _547_v146;
          let _nw90 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
          _547_v146 = _nw90;
          let _548_v147;
          _548_v147 = _module.D0.create_DC1(_544_v144, _545_v145, ((_321_v1) ? (_547_v146) : (_547_v146)));
          _548_v147 = _548_v147;
          (_324_globalState).f4 = (_542_v142).f6;
          _449_v84 = _module.__default.fm7(_dafny.MultiSet.fromElements((((_456_v91).contains(_320_v0)) ? ((_456_v91).get(_320_v0)) : (_455_v90)), false), _352_v20, _324_globalState);
          _321_v1 = _321_v1;
          _449_v84 = (_449_v84).update((_383_v43).dtor_cf6, _320_v0);
        } else {
          let _549_v148;
          _549_v148 = _module.D0.create_DC0(false);
          let _550_v149;
          let _nw91 = new _module.C0();
          _nw91.__ctor((_542_v142).fm1(_dafny.MultiSet.fromElements((_542_v142).f6), _549_v148, _324_globalState), _dafny.Set.fromElements(false, _321_v1));
          _550_v149 = _nw91;
          let _551_v150;
          let _nw92 = new _module.C0();
          _nw92.__ctor((_542_v142).f6, (_454_v89).Intersect((_550_v149).f7));
          _551_v150 = _nw92;
          let _552_v151;
          _552_v151 = _module.D3.create_DC9(_511_v125);
          let _553_v152;
          _553_v152 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_320_v0), _dafny.Seq.update((_552_v151).dtor_cf17, _module.__default.safeIndex(new BigNumber((_449_v84).length), new BigNumber(((_552_v151).dtor_cf17).length)), (_542_v142).f6), _511_v125);
          let _554_v153;
          let _nw93 = new _module.C0();
          _nw93.__ctor((_dafny.ZERO).minus((((_553_v152).contains(_dafny.Seq.update(_511_v125, _module.__default.safeIndex(new BigNumber(122), new BigNumber((_511_v125).length)), new BigNumber((_352_v20).length)))) ? ((_553_v152).get(_dafny.Seq.update(_511_v125, _module.__default.safeIndex(new BigNumber(122), new BigNumber((_511_v125).length)), new BigNumber((_352_v20).length)))) : (new BigNumber(624)))), _dafny.Set.fromElements(_455_v90));
          _554_v153 = _nw93;
          _455_v90 = _455_v90;
          let _555_v154;
          let _out7;
          _out7 = _module.__default.m0(_321_v1, _324_globalState);
          _555_v154 = _out7;
        }
        _455_v90 = !(false) || (!((_452_v87).contains(_module.__default.fm0(false, _321_v1, _320_v0, _324_globalState))));
        let _556_v155;
        _556_v155 = _module.D3.create_DC10(_321_v1, new BigNumber((_module.__default.fm9(new BigNumber(664), new BigNumber((_353_v21).cardinality()), _324_globalState)).length), _320_v0);
        let _557_v156;
        _557_v156 = _module.D3.create_DC12(_556_v155);
        _456_v91 = _module.__default.fm8(_321_v1, _557_v156, _454_v89, _320_v0, _324_globalState);
        if (!(_module.__default.safeDivisionInt((_542_v142).f6, (_542_v142).f6)).isEqualTo(((_542_v142).f6).plus(new BigNumber((_352_v20).length)))) {
          (_324_globalState).f0 = (_542_v142).f6;
          let _558_v157;
          let _out8;
          _out8 = _module.__default.m0(!((_dafny.MultiSet.fromElements(_321_v1, !(true))).IsSubsetOf(_452_v87)), _324_globalState);
          _558_v157 = _out8;
          let _index52 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_453_v88).length));
          (_453_v88)[_index52] = _542_v142;
          let _index53 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_453_v88).length));
          let _rhs57 = (_353_v21).Difference(_353_v21);
          let _rhs58 = ((!(_558_v157) || (_558_v157)) ? ((_320_v0).multipliedBy((_542_v142).f6)) : (new BigNumber((_541_v141).length)));
          let _rhs59 = !(_455_v90);
          let _rhs60 = _542_v142;
          let _lhs38 = _324_globalState;
          let _lhs39 = _453_v88;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_453_v88).length));
          _353_v21 = _rhs57;
          _lhs38.f0 = _rhs58;
          _321_v1 = _rhs59;
          _lhs39[_lhs40] = _rhs60;
          let _559_v158;
          _559_v158 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_module.__default.fm10(new BigNumber((_456_v91).length), (_542_v142).f6, _324_globalState), _dafny.Seq.of(_455_v90)),_module.D3.create_DC12(_556_v155));
          _559_v158 = _module.__default.fm11(_324_globalState);
          let _560_v159;
          let _nw94 = Array((new BigNumber(29)).toNumber()).fill(false);
          _560_v159 = _nw94;
          let _index54 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v159).length));
          (_560_v159)[_index54] = _455_v90;
          let _index55 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v159).length));
          let _rhs61 = _455_v90;
          let _rhs62 = _352_v20;
          let _lhs41 = _560_v159;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_560_v159).length));
          _lhs41[_lhs42] = _rhs61;
          _352_v20 = _rhs62;
        } else {
          _455_v90 = _455_v90;
          _455_v90 = true;
          let _561_v160;
          _561_v160 = _dafny.Map.Empty.slice().updateUnsafe(_322_v2,(_542_v142).f6);
          let _index56 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_512_v126).length));
          (_512_v126)[_index56] = new BigNumber((_561_v160).length);
          let _index57 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_512_v126).length));
          (_512_v126)[_index57] = (_542_v142).f6;
          let _562_v161;
          _562_v161 = _dafny.Map.Empty.slice().updateUnsafe(_542_v142,(((_456_v91).contains((_512_v126)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_512_v126).length))])) ? ((_456_v91).get((_512_v126)[_module.__default.safeIndex(new BigNumber(592), new BigNumber((_512_v126).length))])) : (_321_v1)));
          _562_v161 = (_562_v161).update(_542_v142, _455_v90);
          let _nw95 = new _module.C0();
          _nw95.__ctor(_module.__default.safeModuloInt((_542_v142).f6, _320_v0), _454_v89);
          _542_v142 = _nw95;
        }
      }
      _321_v1 = (_320_v0).isLessThan(new BigNumber(183));
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_512_v126).length))) {
        let _563_i11 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_563_i11)) && ((_563_i11).isLessThan(new BigNumber((_512_v126).length))))) {
          (_512_v126)[(_563_i11)] = (_563_i11).minus(((_455_v90) ? (_320_v0) : (_320_v0)));
        }
      }
      let _564_v162;
      _564_v162 = _module.D3.create_DC9(_511_v125);
      let _source12 = ((((_455_v90) ? (_455_v90) : (!(_455_v90)))) ? (_module.D3.create_DC12(_module.D3.create_DC12(_564_v162))) : (_module.__default.fm12(_321_v1, new BigNumber((_module.__default.fm13(_455_v90, _455_v90, _324_globalState)).length), _455_v90, new BigNumber(49), _324_globalState)));
      if (_source12.is_DC10) {
        let _565___mcc_h28 = (_source12).cf18;
        let _566___mcc_h29 = (_source12).cf19;
        let _567___mcc_h30 = (_source12).cf20;
        let _568_cf20 = _567___mcc_h30;
        let _569_cf19 = _566___mcc_h29;
        let _570_cf18 = _565___mcc_h28;
        let _571_v163;
        _571_v163 = _dafny.Seq.of(_449_v84, _449_v84);
        let _572_v164;
        let _out9;
        _out9 = _module.__default.m0(_dafny.areEqual(_571_v163, _571_v163), _324_globalState);
        _572_v164 = _out9;
        _352_v20 = _dafny.Seq.Concat(_352_v20, _352_v20);
        let _573_v165;
        _573_v165 = _dafny.Seq.of(_322_v2);
        let _574_v166;
        let _nw96 = new _module.C0();
        _nw96.__ctor(new BigNumber((_449_v84).length), _454_v89);
        _574_v166 = _nw96;
        let _575_v167;
        _575_v167 = _dafny.Seq.of(_574_v166, _574_v166);
        let _576_v168;
        _576_v168 = _dafny.Map.Empty.slice().updateUnsafe((_575_v167)[_module.__default.safeIndex(_569_cf19, new BigNumber((_575_v167).length))],_455_v90);
        let _577_v169;
        _577_v169 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_574_v166,!(_572_v164)),_321_v1);
        let _578_v170;
        let _nw97 = Array((new BigNumber(17)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _570_cf18;
        _nw97[(_dafny.ONE).toNumber()] = _321_v1;
        _nw97[(new BigNumber(2)).toNumber()] = _321_v1;
        _nw97[(new BigNumber(3)).toNumber()] = _570_cf18;
        _nw97[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_352_v20, _352_v20);
        _nw97[(new BigNumber(5)).toNumber()] = !(_dafny.Seq.contains((_573_v165)[_module.__default.safeIndex(_569_cf19, new BigNumber((_573_v165).length))], _321_v1));
        _nw97[(new BigNumber(6)).toNumber()] = _321_v1;
        _nw97[(new BigNumber(7)).toNumber()] = !(new BigNumber((_511_v125).length)).isEqualTo(new BigNumber((_449_v84).length));
        _nw97[(new BigNumber(8)).toNumber()] = _455_v90;
        _nw97[(new BigNumber(9)).toNumber()] = !(_568_cf20).isEqualTo(new BigNumber(-329));
        _nw97[(new BigNumber(10)).toNumber()] = _572_v164;
        _nw97[(new BigNumber(11)).toNumber()] = true;
        _nw97[(new BigNumber(12)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_576_v168,_570_cf18)).equals(_577_v169);
        _nw97[(new BigNumber(13)).toNumber()] = ((_321_v1) ? (_570_cf18) : (_455_v90));
        _nw97[(new BigNumber(14)).toNumber()] = _572_v164;
        _nw97[(new BigNumber(15)).toNumber()] = _570_cf18;
        _nw97[(new BigNumber(16)).toNumber()] = _455_v90;
        _578_v170 = _nw97;
        let _index58 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_578_v170).length));
        (_578_v170)[_index58] = _570_cf18;
        let _579_v172;
        _579_v172 = _dafny.Set.fromElements(_322_v2);
        let _580_v173;
        _580_v173 = _dafny.Set.fromElements(new BigNumber((_579_v172).length));
        let _index59 = _module.__default.safeIndex(new BigNumber(75), new BigNumber((_578_v170).length));
        (_578_v170)[_index59] = (_dafny.Set.fromElements((_511_v125)[_module.__default.safeIndex(new BigNumber((_353_v21).cardinality()), new BigNumber((_511_v125).length))], (_574_v166).fm2(_570_cf18, _511_v125, _dafny.Seq.Create(_module.__default.abs(new BigNumber(611)), ((_581_v0) => function (_582_i12) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(908), new BigNumber(-712))) {
              let _583_v171 = _compr_25;
              if (((new BigNumber(908)).isLessThanOrEqualTo(_583_v171)) && ((_583_v171).isLessThan(new BigNumber(-712)))) {
                _coll25.push([(_583_v171).multipliedBy(_581_v0),false]);
              }
            }
            return _coll25;
          }()).length),true)).length);
        })(_320_v0)), _449_v84, _324_globalState))).IsDisjointFrom(_580_v173);
        let _584_v174;
        _584_v174 = new _dafny.CodePoint('a'.codePointAt(0));
        _569_cf19 = new BigNumber((_dafny.Seq.update(_352_v20, _module.__default.safeIndex(((_574_v166).f6).multipliedBy(_320_v0), new BigNumber((_352_v20).length)), _584_v174)).length);
      } else if (_source12.is_DC11) {
        let _585_v175;
        _585_v175 = _module.D3.create_DC9(_dafny.Seq.of(_320_v0));
        let _pat_let_tv22 = _511_v125;
        let _586_v176;
        _586_v176 = _dafny.MultiSet.fromElements(_585_v175, _585_v175, _585_v175, function (_pat_let10_0) {
          return function (_587_dt__update__tmp_h3) {
            return function (_pat_let11_0) {
              return function (_588_dt__update_hcf17_h0) {
                return _module.D3.create_DC9(_588_dt__update_hcf17_h0);
              }(_pat_let11_0);
            }(_pat_let_tv22);
          }(_pat_let10_0);
        }(_585_v175));
        let _index60 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_512_v126).length));
        (_512_v126)[_index60] = (_dafny.ZERO).minus(new BigNumber((_586_v176).cardinality()));
        let _589_v177;
        _589_v177 = _dafny.Map.Empty.slice().updateUnsafe(_321_v1,true);
        let _index61 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_512_v126).length));
        (_512_v126)[_index61] = _module.__default.fm4(_320_v0, (((_589_v177).contains(_321_v1)) ? ((_589_v177).get(_321_v1)) : (_455_v90)), _324_globalState);
        let _590_v178;
        _590_v178 = _dafny.Set.fromElements(_449_v84, _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC0(false)).dtor_cf0,new BigNumber((_322_v2).length)), _449_v84);
        let _591_v179;
        _591_v179 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber(-660), _320_v0), new BigNumber((_590_v178).length), new BigNumber(253));
        _591_v179 = _591_v179;
        let _592_v180;
        let _nw98 = new _module.C0();
        _nw98.__ctor(_module.__default.fm4(new BigNumber((_591_v179).length), _455_v90, _324_globalState), _454_v89);
        _592_v180 = _nw98;
        let _index62 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_512_v126).length));
        (_512_v126)[_index62] = _320_v0;
      } else if (_source12.is_DC9) {
        let _593___mcc_h31 = (_source12).cf17;
        let _594_cf17 = _593___mcc_h31;
        _449_v84 = ((_449_v84).Merge(_449_v84)).Merge(_449_v84);
        if (_321_v1) {
          let _595_v181;
          _595_v181 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_320_v0),new BigNumber((((_455_v90) ? (_352_v20) : (_352_v20))).length));
          _595_v181 = (_595_v181).Merge(_595_v181);
          let _596_v182;
          _596_v182 = new _dafny.CodePoint('g'.codePointAt(0));
          let _597_v183;
          _597_v183 = _dafny.Map.Empty.slice().updateUnsafe(false,_596_v182);
          let _598_v184;
          _598_v184 = _dafny.Map.Empty.slice().updateUnsafe(_352_v20,(((_597_v183).contains(_321_v1)) ? ((_597_v183).get(_321_v1)) : (_596_v182)));
          _598_v184 = (_598_v184).update(_352_v20, _596_v182);
          _320_v0 = (_320_v0).plus(_320_v0);
          let _599_v185;
          _599_v185 = _module.D2.create_DC5(_512_v126);
          let _pat_let_tv23 = _512_v126;
          let _pat_let_tv24 = _512_v126;
          _599_v185 = function (_pat_let12_0) {
            return function (_600_dt__update__tmp_h4) {
              return function (_pat_let13_0) {
                return function (_601_dt__update_hcf11_h1) {
                  return _module.D2.create_DC5(_601_dt__update_hcf11_h1);
                }(_pat_let13_0);
              }(((true) ? (_pat_let_tv23) : (_pat_let_tv24)));
            }(_pat_let12_0);
          }(_599_v185);
          let _602_v186;
          let _nw99 = new _module.C0();
          _nw99.__ctor(_320_v0, _454_v89);
          _602_v186 = _nw99;
          let _index63 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_453_v88).length));
          (_453_v88)[_index63] = _602_v186;
          let _index64 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_453_v88).length));
          (_453_v88)[_index64] = ((_321_v1) ? (_602_v186) : (_602_v186));
        } else {
          let _603_v187;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_320_v0, _dafny.Set.fromElements(false, _321_v1));
          _603_v187 = _nw100;
          _603_v187 = _603_v187;
          _352_v20 = _352_v20;
          let _index65 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_512_v126).length));
          (_512_v126)[_index65] = _320_v0;
          let _index66 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_512_v126).length));
          (_512_v126)[_index66] = _320_v0;
          _603_v187 = _603_v187;
          let _604_v188;
          _604_v188 = _dafny.Seq.of(_352_v20);
          let _index67 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_512_v126).length));
          (_512_v126)[_index67] = ((new BigNumber((_604_v188).length)).multipliedBy(new BigNumber(-818))).multipliedBy(new BigNumber((_452_v87).cardinality()));
        }
        let _605_v189;
        let _nw101 = Array((new BigNumber(26)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = (_322_v2)[_module.__default.safeIndex(_320_v0, new BigNumber((_322_v2).length))];
        _nw101[(_dafny.ONE).toNumber()] = _321_v1;
        _nw101[(new BigNumber(2)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(3)).toNumber()] = _455_v90;
        _nw101[(new BigNumber(4)).toNumber()] = false;
        _nw101[(new BigNumber(5)).toNumber()] = true;
        _nw101[(new BigNumber(6)).toNumber()] = !(false);
        _nw101[(new BigNumber(7)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(8)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(9)).toNumber()] = _455_v90;
        _nw101[(new BigNumber(10)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(11)).toNumber()] = !(_module.__default.fm0(_321_v1, _321_v1, _module.__default.fm4(new BigNumber((_322_v2).length), _321_v1, _324_globalState), _324_globalState));
        _nw101[(new BigNumber(12)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(13)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(14)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(15)).toNumber()] = false;
        _nw101[(new BigNumber(16)).toNumber()] = _455_v90;
        _nw101[(new BigNumber(17)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(18)).toNumber()] = _455_v90;
        _nw101[(new BigNumber(19)).toNumber()] = true;
        _nw101[(new BigNumber(20)).toNumber()] = _455_v90;
        _nw101[(new BigNumber(21)).toNumber()] = _module.__default.fm0(_321_v1, (_383_v43).dtor_cf5, _320_v0, _324_globalState);
        _nw101[(new BigNumber(22)).toNumber()] = !(_321_v1);
        _nw101[(new BigNumber(23)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(24)).toNumber()] = _321_v1;
        _nw101[(new BigNumber(25)).toNumber()] = _321_v1;
        _605_v189 = _nw101;
        let _606_v190;
        _606_v190 = _dafny.Map.Empty.slice().updateUnsafe(_455_v90,_605_v189);
        let _607_v191;
        let _nw102 = Array((new BigNumber(26)).toNumber());
        _nw102[(_dafny.ZERO).toNumber()] = (((_606_v190).contains(_321_v1)) ? ((_606_v190).get(_321_v1)) : (_605_v189));
        _nw102[(_dafny.ONE).toNumber()] = _605_v189;
        _nw102[(new BigNumber(2)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(3)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(4)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(5)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(6)).toNumber()] = ((_455_v90) ? (_605_v189) : (_605_v189));
        _nw102[(new BigNumber(7)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(8)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(9)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(10)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(11)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(12)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(13)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(14)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(15)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(16)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(17)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(18)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(19)).toNumber()] = ((_321_v1) ? (_605_v189) : (_605_v189));
        _nw102[(new BigNumber(20)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(21)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(22)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(23)).toNumber()] = ((_321_v1) ? (_605_v189) : (_605_v189));
        _nw102[(new BigNumber(24)).toNumber()] = _605_v189;
        _nw102[(new BigNumber(25)).toNumber()] = _605_v189;
        _607_v191 = _nw102;
        let _index68 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_607_v191).length));
        (_607_v191)[_index68] = _605_v189;
        let _index69 = _module.__default.safeIndex(new BigNumber(215), new BigNumber((_607_v191).length));
        (_607_v191)[_index69] = _605_v189;
        let _source13 = _458_v93;
        if (_source13.is_DC3) {
          let _608___mcc_h33 = (_source13).cf5;
          let _609___mcc_h34 = (_source13).cf6;
          let _610___mcc_h35 = (_source13).cf7;
          let _611___mcc_h36 = (_source13).cf8;
          let _612___mcc_h37 = (_source13).cf9;
          let _613_cf9 = _612___mcc_h37;
          let _614_cf8 = _611___mcc_h36;
          let _615_cf7 = _610___mcc_h35;
          let _616_cf6 = _609___mcc_h34;
          let _617_cf5 = _608___mcc_h33;
          let _618_v192;
          _618_v192 = new _dafny.CodePoint('v'.codePointAt(0));
          let _619_v193;
          _619_v193 = _dafny.Map.Empty.slice().updateUnsafe(_614_cf8,false);
          let _620_v194;
          let _nw103 = Array((new BigNumber(2)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_614_cf8, (((_619_v193).contains(_455_v90)) ? ((_619_v193).get(_455_v90)) : (_616_cf6)));
          _nw103[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_617_cf5, (_322_v2)[_module.__default.safeIndex(new BigNumber((_511_v125).length), new BigNumber((_322_v2).length))], !(_455_v90));
          _620_v194 = _nw103;
          let _index70 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_620_v194).length));
          (_620_v194)[_index70] = _454_v89;
          let _621_v195;
          _621_v195 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,_module.D3.create_DC11());
          let _index71 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_620_v194).length));
          let _rhs63 = _613_cf9;
          let _rhs64 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_594_cf17, _511_v125))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_320_v0, _320_v0, _320_v0));
          let _rhs65 = _618_v192;
          let _rhs66 = ((!((new BigNumber((_457_v92).length)).isLessThanOrEqualTo(new BigNumber((_621_v195).length)))) ? (_454_v89) : (_454_v89));
          let _lhs43 = _324_globalState;
          let _lhs44 = _620_v194;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(316), new BigNumber((_620_v194).length));
          _lhs43.f0 = _rhs63;
          _614_cf8 = _rhs64;
          _618_v192 = _rhs65;
          _lhs44[_lhs45] = _rhs66;
          let _index72 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_605_v189).length));
          (_605_v189)[_index72] = true;
          let _index73 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_605_v189).length));
          (_605_v189)[_index73] = _module.__default.fm0(_321_v1, _455_v90, _613_cf9, _324_globalState);
          let _622_v196;
          let _nw104 = new _module.C0();
          _nw104.__ctor(_module.__default.safeDivisionInt(_613_cf9, (_dafny.ZERO).minus(new BigNumber((_606_v190).length))), _module.__default.fm5(_324_globalState));
          _622_v196 = _nw104;
          _616_cf6 = true;
        } else if (_source13.is_DC2) {
          let _623___mcc_h38 = (_source13).cf4;
          let _624_cf4 = _623___mcc_h38;
          let _625_v197;
          let _nw105 = new _module.C0();
          _nw105.__ctor(_module.__default.safeModuloInt(_module.__default.fm4(_624_cf4, _321_v1, _324_globalState), (_dafny.ZERO).minus(_320_v0)), ((_455_v90) ? (_dafny.Set.fromElements(_455_v90)) : (_454_v89)));
          _625_v197 = _nw105;
          let _626_v198;
          _626_v198 = _dafny.Map.Empty.slice().updateUnsafe(_455_v90,(_322_v2)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_322_v2).length))]);
          _625_v197 = (((((_626_v198).contains(_455_v90)) ? ((_626_v198).get(_455_v90)) : (_455_v90))) ? (_625_v197) : (_625_v197));
          let _627_v199;
          _627_v199 = _module.D4.create_DC13(_625_v197);
          _625_v197 = (_627_v199).dtor_cf22;
          (_324_globalState).f4 = ((_625_v197).f6).plus(_624_cf4);
          (_324_globalState).f4 = new BigNumber((((_449_v84).Merge(_449_v84)).Merge((_449_v84).Merge(_449_v84))).length);
        } else {
          let _628___mcc_h39 = (_source13).cf10;
          let _629_cf10 = _628___mcc_h39;
          let _630_v200;
          _630_v200 = _dafny.Seq.of(_512_v126);
          let _631_v201;
          _631_v201 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_455_v90, _321_v1, _320_v0, _324_globalState),_630_v200);
          _631_v201 = (_631_v201).update(_455_v90, ((_321_v1) ? (_630_v200) : (_630_v200)));
          (_324_globalState).f4 = (new BigNumber(((_449_v84).Merge((_449_v84).update(_321_v1, _320_v0))).length)).multipliedBy(_module.__default.fm4((_dafny.ZERO).minus(new BigNumber((_352_v20).length)), _module.__default.fm0(_321_v1, _321_v1, _320_v0, _324_globalState), _324_globalState));
          let _632_v202;
          let _nw106 = new _module.C0();
          _nw106.__ctor(new BigNumber(382), _454_v89);
          _632_v202 = _nw106;
          let _633_v203;
          _633_v203 = _dafny.Set.fromElements(_632_v202, _632_v202);
          let _634_v204;
          _634_v204 = new _dafny.CodePoint('r'.codePointAt(0));
          let _635_v205;
          _635_v205 = _module.D4.create_DC16(_633_v203, _634_v204, (_632_v202).f6, _634_v204);
          let _636_v206;
          _636_v206 = _dafny.Map.Empty.slice().updateUnsafe(_320_v0,(_dafny.ZERO).minus(_320_v0));
          let _rhs67 = !(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_632_v202).f6,_632_v202)).length), (_632_v202).f6)).isEqualTo((_dafny.ZERO).minus(_320_v0));
          let _rhs68 = _635_v205;
          let _rhs69 = _636_v206;
          let _rhs70 = _512_v126;
          _455_v90 = _rhs67;
          _635_v205 = _rhs68;
          _636_v206 = _rhs69;
          _512_v126 = _rhs70;
        }
      } else {
        let _637___mcc_h32 = (_source12).cf21;
        let _638_cf21 = _637___mcc_h32;
        let _639_v207;
        let _out10;
        _out10 = _module.__default.m0(!(_455_v90), _324_globalState);
        _639_v207 = _out10;
        (_324_globalState).f0 = _320_v0;
        let _640_v208;
        _640_v208 = _module.D3.create_DC10(!(_455_v90), _320_v0, new BigNumber(-144));
        let _pat_let_tv25 = _320_v0;
        let _641_v209;
        let _nw107 = Array((new BigNumber(23)).toNumber());
        _nw107[(_dafny.ZERO).toNumber()] = _455_v90;
        _nw107[(_dafny.ONE).toNumber()] = _455_v90;
        _nw107[(new BigNumber(2)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(3)).toNumber()] = !(_321_v1);
        _nw107[(new BigNumber(4)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(5)).toNumber()] = false;
        _nw107[(new BigNumber(6)).toNumber()] = (_639_v207) && (!(_639_v207));
        _nw107[(new BigNumber(7)).toNumber()] = _321_v1;
        _nw107[(new BigNumber(8)).toNumber()] = _321_v1;
        _nw107[(new BigNumber(9)).toNumber()] = !(true);
        _nw107[(new BigNumber(10)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(11)).toNumber()] = _321_v1;
        _nw107[(new BigNumber(12)).toNumber()] = _639_v207;
        _nw107[(new BigNumber(13)).toNumber()] = (new BigNumber((_dafny.MultiSet.fromElements(_321_v1)).cardinality())).isLessThan(new BigNumber(-508));
        _nw107[(new BigNumber(14)).toNumber()] = !(_455_v90) || (false);
        _nw107[(new BigNumber(15)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(16)).toNumber()] = _321_v1;
        _nw107[(new BigNumber(17)).toNumber()] = _639_v207;
        _nw107[(new BigNumber(18)).toNumber()] = (function (_pat_let14_0) {
          return function (_642_dt__update__tmp_h5) {
            return function (_pat_let15_0) {
              return function (_643_dt__update_hcf20_h0) {
                return _module.D3.create_DC10((_642_dt__update__tmp_h5).dtor_cf18, (_642_dt__update__tmp_h5).dtor_cf19, _643_dt__update_hcf20_h0);
              }(_pat_let15_0);
            }(_pat_let_tv25);
          }(_pat_let14_0);
        }(_640_v208)).dtor_cf18;
        _nw107[(new BigNumber(19)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(20)).toNumber()] = _455_v90;
        _nw107[(new BigNumber(21)).toNumber()] = _639_v207;
        _nw107[(new BigNumber(22)).toNumber()] = (_322_v2)[_module.__default.safeIndex(new BigNumber(-229), new BigNumber((_322_v2).length))];
        _641_v209 = _nw107;
        let _index74 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_641_v209).length));
        (_641_v209)[_index74] = _455_v90;
        let _index75 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_641_v209).length));
        (_641_v209)[_index75] = _321_v1;
        let _644_v210;
        _644_v210 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_456_v91).length),_320_v0);
        let _645_v212;
        let _nw108 = new _module.C0();
        _nw108.__ctor(new BigNumber((function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe(_644_v210,_322_v2)).Keys.Elements) {
            let _646_v211 = _compr_26;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_644_v210,_322_v2)).contains(_646_v211)) {
              _coll26.add(_646_v211);
            }
          }
          return _coll26;
        }()).length), _dafny.Set.fromElements(_321_v1));
        _645_v212 = _nw108;
      }
      if (!((_452_v87).Union(_452_v87)).equals(_452_v87)) {
        let _647_v213;
        _647_v213 = new _dafny.CodePoint('m'.codePointAt(0));
        _647_v213 = _647_v213;
        _456_v91 = (_456_v91).Merge(_456_v91);
        _352_v20 = _dafny.Seq.UnicodeFromString("qyqejkm");
        let _648_v214;
        _648_v214 = _module.D2.create_DC7(_module.__default.fm4((_dafny.ZERO).minus(_320_v0), _321_v1, _324_globalState));
        let _649_v215;
        _649_v215 = _module.D2.create_DC8(_648_v214);
        let _source14 = _649_v215;
        if (_source14.is_DC6) {
          let _650___mcc_h40 = (_source14).cf12;
          let _651___mcc_h41 = (_source14).cf13;
          let _652___mcc_h42 = (_source14).cf14;
          let _653_cf14 = _652___mcc_h42;
          let _654_cf13 = _651___mcc_h41;
          let _655_cf12 = _650___mcc_h40;
          let _656_v216;
          let _nw109 = Array((new BigNumber(6)).toNumber());
          _nw109[(_dafny.ZERO).toNumber()] = (_320_v0).isEqualTo((_511_v125)[_module.__default.safeIndex(_655_cf12, new BigNumber((_511_v125).length))]);
          _nw109[(_dafny.ONE).toNumber()] = true;
          _nw109[(new BigNumber(2)).toNumber()] = !((((_456_v91).contains(_655_cf12)) ? ((_456_v91).get(_655_cf12)) : (_654_cf13)));
          _nw109[(new BigNumber(3)).toNumber()] = _455_v90;
          _nw109[(new BigNumber(4)).toNumber()] = _321_v1;
          _nw109[(new BigNumber(5)).toNumber()] = _321_v1;
          _656_v216 = _nw109;
          let _index76 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_656_v216).length));
          (_656_v216)[_index76] = _321_v1;
          let _index77 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_656_v216).length));
          (_656_v216)[_index77] = _321_v1;
          let _657_v217;
          let _init11 = ((_658_v20) => function (_659_i13) {
            return _dafny.Seq.Concat(_658_v20, _658_v20);
          })(_352_v20);
          let _nw110 = Array((new BigNumber(10)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw110.length); _i0_11++) {
            _nw110[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _657_v217 = _nw110;
          let _index78 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_657_v217).length));
          (_657_v217)[_index78] = _352_v20;
          let _index79 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_657_v217).length));
          (_657_v217)[_index79] = _dafny.Seq.Concat(_352_v20, _352_v20);
          let _660_v220;
          let _init12 = ((_661_cf12) => function (_662_i14) {
            return (function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of _dafny.IntegerRange(new BigNumber(873), new BigNumber(129))) {
                let _663_v218 = _compr_27;
                if (((new BigNumber(873)).isLessThanOrEqualTo(_663_v218)) && ((_663_v218).isLessThan(new BigNumber(129)))) {
                  _coll27.add((_663_v218).minus(_661_cf12));
                }
              }
              return _coll27;
            }()).Difference(function () {
              let _coll28 = new _dafny.Set();
              for (const _compr_28 of _dafny.IntegerRange(new BigNumber(15), new BigNumber(772))) {
                let _664_v219 = _compr_28;
                if (((new BigNumber(15)).isLessThanOrEqualTo(_664_v219)) && ((_664_v219).isLessThan(new BigNumber(772)))) {
                  _coll28.add((_664_v219).plus(_661_cf12));
                }
              }
              return _coll28;
            }());
          })(_655_cf12);
          let _nw111 = Array((new BigNumber(26)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw111.length); _i0_12++) {
            _nw111[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _660_v220 = _nw111;
          let _665_v221;
          _665_v221 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_655_cf12, new BigNumber((_352_v20).length))).length), _320_v0);
          let _666_v222;
          _666_v222 = _dafny.Seq.of(_454_v89, _454_v89);
          let _667_v223;
          let _nw112 = new _module.C0();
          _nw112.__ctor(_655_cf12, (_666_v222)[_module.__default.safeIndex(new BigNumber(((_657_v217)[_module.__default.safeIndex(new BigNumber(2), new BigNumber((_657_v217).length))]).length), new BigNumber((_666_v222).length))]);
          _667_v223 = _nw112;
          let _668_v224;
          _668_v224 = _dafny.Set.fromElements(_667_v223);
          let _669_v225;
          _669_v225 = _module.D4.create_DC16(_668_v224, _647_v213, _320_v0, _647_v213);
          let _index80 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_660_v220).length));
          (_660_v220)[_index80] = ((true) ? (_665_v221) : (_dafny.Set.fromElements(new BigNumber((_665_v221).length), _655_cf12, _655_cf12, (_669_v225).dtor_cf27, _655_cf12)));
          let _index81 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_660_v220).length));
          (_660_v220)[_index81] = _665_v221;
          let _670_v226;
          _670_v226 = _module.D3.create_DC12(_564_v162);
          let _671_v227;
          _671_v227 = _dafny.Map.Empty.slice().updateUnsafe(_670_v226,_667_v223);
          _671_v227 = (_671_v227).update(_module.D3.create_DC12(_564_v162), _667_v223);
        } else if (_source14.is_DC7) {
          let _672___mcc_h43 = (_source14).cf15;
          let _673_cf15 = _672___mcc_h43;
          let _674_v228;
          let _nw113 = new _module.C0();
          _nw113.__ctor(new BigNumber((_511_v125).length), _454_v89);
          _674_v228 = _nw113;
          let _675_v229;
          _675_v229 = _module.D4.create_DC13(_674_v228);
          let _676_v230;
          _676_v230 = _dafny.Map.Empty.slice().updateUnsafe(_647_v213,_module.__default.fm0(_455_v90, _321_v1, _320_v0, _324_globalState));
          let _677_v231;
          _677_v231 = _module.D1.create_DC3(_321_v1, _module.__default.fm0(_module.__default.fm0(_455_v90, _321_v1, new BigNumber(844), _324_globalState), (((_676_v230).contains(_647_v213)) ? ((_676_v230).get(_647_v213)) : (true)), new BigNumber((_322_v2).length), _324_globalState), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("tim"), _module.__default.safeIndex(new BigNumber((_452_v87).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("tim")).length)), new _dafny.CodePoint('u'.codePointAt(0))), _321_v1, _320_v0);
          let _678_v232;
          _678_v232 = _module.D1.create_DC4(_677_v231);
          let _rhs71 = _module.D4.create_DC13(((_455_v90) ? (_674_v228) : (_674_v228)));
          let _rhs72 = _module.D1.create_DC4(_module.__default.fm14(_324_globalState));
          let _rhs73 = _673_cf15;
          let _lhs46 = _324_globalState;
          _675_v229 = _rhs71;
          _678_v232 = _rhs72;
          _lhs46.f0 = _rhs73;
          let _679_v233;
          let _out11;
          _out11 = _module.__default.m0(!(_321_v1), _324_globalState);
          _679_v233 = _out11;
          _679_v233 = !(((_module.__default.fm0(!(_455_v90), _321_v1, new BigNumber((_511_v125).length), _324_globalState)) ? (_679_v233) : (_321_v1)));
          _320_v0 = new BigNumber(230);
        } else if (_source14.is_DC5) {
          let _680___mcc_h44 = (_source14).cf11;
          let _681_cf11 = _680___mcc_h44;
          (_324_globalState).f0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_320_v0, (((_322_v2)[_module.__default.safeIndex(_module.__default.fm4(_320_v0, _321_v1, _324_globalState), new BigNumber((_322_v2).length))]) ? (_320_v0) : (_320_v0))));
          let _682_v234;
          let _nw114 = Array((new BigNumber(20)).toNumber()).fill(false);
          _682_v234 = _nw114;
          let _index82 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_682_v234).length));
          (_682_v234)[_index82] = _321_v1;
          let _683_v235;
          _683_v235 = _module.D3.create_DC12(_module.D3.create_DC11());
          let _index83 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_682_v234).length));
          let _rhs74 = _455_v90;
          let _rhs75 = _683_v235;
          let _rhs76 = _682_v234;
          let _rhs77 = (_module.__default.fm4(_320_v0, _321_v1, _324_globalState)).isLessThanOrEqualTo(((_module.D2.create_DC6(_320_v0, _455_v90, _512_v126)).dtor_cf12).multipliedBy(_320_v0));
          let _rhs78 = _module.__default.fm0(_455_v90, _321_v1, _320_v0, _324_globalState);
          let _lhs47 = _682_v234;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(321), new BigNumber((_682_v234).length));
          _lhs47[_lhs48] = _rhs74;
          _683_v235 = _rhs75;
          _682_v234 = _rhs76;
          _455_v90 = _rhs77;
          _455_v90 = _rhs78;
          let _684_v236;
          _684_v236 = _module.D3.create_DC10(_455_v90, _320_v0, _320_v0);
          (_324_globalState).f0 = (new BigNumber(722)).plus(((_dafny.ZERO).minus(new BigNumber((_452_v87).cardinality()))).multipliedBy((_684_v236).dtor_cf19));
          (_324_globalState).f0 = ((_320_v0).multipliedBy(_320_v0)).minus(_320_v0);
        } else {
          let _685___mcc_h45 = (_source14).cf16;
          let _686_cf16 = _685___mcc_h45;
          (_324_globalState).f0 = (_320_v0).multipliedBy(_module.__default.safeDivisionInt(_320_v0, (_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.FromArray(_322_v2)).update(_321_v1, _module.__default.abs(_320_v0))).cardinality()))));
          _455_v90 = !(true);
          let _687_v237;
          let _nw115 = new _module.C0();
          _nw115.__ctor(_320_v0, _454_v89);
          _687_v237 = _nw115;
          let _688_v238;
          _688_v238 = _module.D3.create_DC10(_455_v90, new BigNumber(354), (_687_v237).f6);
          _321_v1 = (_688_v238).dtor_cf18;
        }
        let _pat_let_tv26 = _455_v90;
        let _689_v239;
        _689_v239 = _dafny.Seq.of(function (_pat_let16_0) {
          return function (_690_dt__update__tmp_h6) {
            return function (_pat_let17_0) {
              return function (_691_dt__update_hcf23_h0) {
                return _module.D4.create_DC15(_691_dt__update_hcf23_h0, (_690_dt__update__tmp_h6).dtor_cf24);
              }(_pat_let17_0);
            }(_pat_let_tv26);
          }(_pat_let16_0);
        }(_module.D4.create_DC15(_321_v1, _320_v0)), _module.D4.create_DC15(!(_321_v1), new BigNumber(412)));
        (_324_globalState).f0 = (_module.__default.safeModuloInt(new BigNumber((_352_v20).length), _320_v0)).multipliedBy(new BigNumber((_dafny.Seq.Concat(_689_v239, _689_v239)).length));
      } else {
        let _692_v240;
        let _nw116 = new _module.C0();
        _nw116.__ctor(new BigNumber((_352_v20).length), _454_v89);
        _692_v240 = _nw116;
        (_324_globalState).f0 = (_692_v240).f6;
        let _693_v241;
        _693_v241 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_352_v20, _352_v20),(_322_v2)[_module.__default.safeIndex((_692_v240).f6, new BigNumber((_322_v2).length))]);
        _693_v241 = (_693_v241).update(_dafny.Seq.Concat(_352_v20, _352_v20), _455_v90);
        let _694_v242;
        let _nw117 = new _module.C0();
        _nw117.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_457_v92).length)), _511_v125)).length), (_692_v240).f7);
        _694_v242 = _nw117;
        let _695_v243;
        let _out12;
        _out12 = _module.__default.m0(true, _324_globalState);
        _695_v243 = _out12;
      }
      let _696_v244;
      let _nw118 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
      _696_v244 = _nw118;
      let _index84 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_696_v244).length));
      (_696_v244)[_index84] = _dafny.Seq.of(_512_v126, _512_v126);
      let _697_v245;
      _697_v245 = _dafny.Seq.of(_512_v126, _512_v126);
      let _index85 = _module.__default.safeIndex(new BigNumber(834), new BigNumber((_696_v244).length));
      (_696_v244)[_index85] = _697_v245;
      process.stdout.write(_dafny.toString(_320_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_321_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_322_v2, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_323_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-45),_dafny.Seq.of(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_324_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_324_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_324_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_324_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_324_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_324_globalState).f5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-45),_dafny.Seq.of(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_352_v20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_353_v21).equals(_dafny.MultiSet.fromElements(new BigNumber(6), new BigNumber(6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_383_v43).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_383_v43).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_383_v43).dtor_cf7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_383_v43).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_383_v43).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_449_v84).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_450_v85).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(492),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_451_v86).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_452_v87).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_454_v89).equals(_dafny.Set.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_455_v90));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_456_v91).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(6),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_457_v92).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_458_v93).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_459_v94)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_511_v125, _dafny.Seq.of(_dafny.ZERO, _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_512_v126)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_513_v127).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_513_v127).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_513_v127).dtor_cf14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_564_v162).dtor_cf17, _dafny.Seq.of(_dafny.ZERO, _dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_696_v244)[new BigNumber(6)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_697_v245).length)));
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
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(false);
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
    static create_DC3(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC4(cf10) {
      let $dt = new D1(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + this.cf7.toVerbatimString(true) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf5 === other.cf5 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, false, _dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC6(cf12, cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf15) {
      let $dt = new D2(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC5(cf11) {
      let $dt = new D2(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC8(cf16) {
      let $dt = new D2(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf11 === other.cf11;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO, false, []);
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
    static create_DC10(cf18, cf19, cf20) {
      let $dt = new D3(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC11() {
      let $dt = new D3(1);
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D3(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D3(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC12() { return this.$tag === 3; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC14() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC15(cf23, cf24) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC16(cf25, cf26, cf27, cf28) {
      let $dt = new D4(2);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC13(cf22) {
      let $dt = new D4(3);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC13() { return this.$tag === 3; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14";
      } else if (this.$tag === 1) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ")";
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
        return other.$tag === 1 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf22 === other.cf22;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14();
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
    static create_DC17(cf29) {
      let $dt = new D5(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf31, cf32, cf33) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC20(cf34) {
      let $dt = new D6(1);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC21() {
      let $dt = new D6(2);
      return $dt;
    }
    static create_DC18(cf30) {
      let $dt = new D6(3);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21";
      } else if (this.$tag === 3) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC19(_dafny.ZERO, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
      this.f4 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f5 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
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
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6, f7) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), function (_698_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length));
    };
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D1.create_DC2(new BigNumber((_dafny.Seq.UnicodeFromString("m")).length))).dtor_cf4;
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
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
