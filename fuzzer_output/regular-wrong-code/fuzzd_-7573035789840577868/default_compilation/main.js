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
      return !(_dafny.areEqual(_dafny.Seq.of(_module.D0.create_DC1(false), _module.D0.create_DC1(!(true))), _dafny.Seq.Concat(_dafny.Seq.of(_module.D0.create_DC1(false)), _dafny.Seq.of(_module.D0.create_DC1(!(!(true)))))));
    };
    static fm1(p0, globalState) {
      return _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("hu"), (_dafny.ZERO).minus(((false) ? (new BigNumber((_dafny.Seq.UnicodeFromString("bgppqsue")).length)) : (new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(850),new BigNumber(630)))).length)))), new BigNumber(-415));
    };
    static fm2(globalState) {
      return _module.D1.create_DC6((_dafny.ZERO).minus((_module.D11.create_DC30(false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false), false, true, false, true)).length))).length), true)).dtor_cf53), new BigNumber((_dafny.Seq.UnicodeFromString("qipvo")).length));
    };
    static fm3(p0, p1, p2, globalState) {
      return (((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))).plus(new BigNumber(384))).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kifiqbht")).length));
    };
    static fm5(p0, globalState) {
      let _source0 = _module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC2(_module.D0.create_DC1(true)))));
      if (_source0.is_DC1) {
        let _0___mcc_h0 = (_source0).cf1;
        let _1_cf1 = _0___mcc_h0;
        return ((_dafny.ZERO).minus(new BigNumber(-50))).minus(new BigNumber(435));
      } else if (_source0.is_DC0) {
        let _2___mcc_h1 = (_source0).cf0;
        let _3_cf0 = _2___mcc_h1;
        return (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(95)), function (_4_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })).length)));
      } else {
        let _5___mcc_h2 = (_source0).cf2;
        let _6_cf2 = _5___mcc_h2;
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false))).length));
      }
    };
    static fm9(p0, p1, p2, globalState) {
      return (((true) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rch")).length))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))))).minus(new BigNumber(((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(161),false)).length), new BigNumber(240), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(!(true), false)).length))).length))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xpetoj")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_7_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length))).length), new BigNumber(90)))).length));
    };
    static fm10(globalState) {
      return ((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(387), new BigNumber(476))) {
          let _8_v0 = _compr_0;
          if (((new BigNumber(387)).isLessThanOrEqualTo(_8_v0)) && ((_8_v0).isLessThan(new BigNumber(476)))) {
            _coll0.push([(_8_v0).multipliedBy(new BigNumber(853)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(758)), function (_9_i0) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length)]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), function (_10_i1) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length)),new BigNumber(784)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(572),new BigNumber((_dafny.Seq.of(new BigNumber(249))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber((_dafny.Seq.UnicodeFromString("bkuqdvcwu")).length))));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC2(_module.D0.create_DC0(false));
    };
    static fm12(p0, p1, p2, globalState) {
      let _source1 = _module.D0.create_DC0(true);
      if (_source1.is_DC1) {
        let _11___mcc_h0 = (_source1).cf1;
        let _12_cf1 = _11___mcc_h0;
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else if (_source1.is_DC0) {
        let _13___mcc_h1 = (_source1).cf0;
        let _14_cf0 = _13___mcc_h1;
        return new _dafny.CodePoint('r'.codePointAt(0));
      } else {
        let _15___mcc_h2 = (_source1).cf2;
        let _16_cf2 = _15___mcc_h2;
        return new _dafny.CodePoint('b'.codePointAt(0));
      }
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-955)), function (_17_i0) {
        return new BigNumber(372);
      }), (_module.D9.create_DC24(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("tqeuo")).length))).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),new BigNumber((_dafny.Set.fromElements(new BigNumber(456), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dst")).length), (_dafny.ZERO).minus(new BigNumber(-802)), new BigNumber(587))).length))).length))).length)))).dtor_cf42), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-557), new BigNumber(-934))).length), new BigNumber(811), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-612), new BigNumber((_dafny.Seq.of(false)).length))).length)))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wggx")).length)), new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)));
    };
    static fm16(p0, p1, globalState) {
      if (true) {
        if (!(false)) {
          return _dafny.Seq.UnicodeFromString("lr");
        } else {
          return _dafny.Seq.UnicodeFromString("mrdlm");
        }
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mtnbe"), _dafny.Seq.UnicodeFromString("sunbm"));
      }
    };
    static fm17(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(990),!((false) === (true)));
    };
    static fm18(p0, p1, p2, globalState) {
      return (function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.MultiSet.fromElements(new BigNumber(510))).Elements) {
          let _18_v0 = _compr_1;
          if ((_dafny.MultiSet.fromElements(new BigNumber(510))).contains(_18_v0)) {
            _coll1.add((_18_v0).minus(new BigNumber((_dafny.Seq.of(true)).length)));
          }
        }
        return _coll1;
      }()).Difference(function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(958),new BigNumber(174))).Keys.Elements) {
          let _19_v1 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(958),new BigNumber(174))).contains(_19_v1)) {
            _coll2.add((_19_v1).minus(new BigNumber(688)));
          }
        }
        return _coll2;
      }());
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.of(((true) ? (new BigNumber(779)) : (new BigNumber(725))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber(-617))).length));
    };
    static fm20(globalState) {
      return _dafny.Set.fromElements((new BigNumber(-391)).isLessThanOrEqualTo(new BigNumber(-355)), true);
    };
    static fm21(globalState) {
      return _module.D1.create_DC5(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_20_i0) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xqjjx"), _dafny.Seq.UnicodeFromString("uvlppdqpi"))).length), ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(712), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(867),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), function (_21_i1) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length),new _dafny.CodePoint('u'.codePointAt(0)))).length))).length)), new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!(true))).length))).Elements) {
    let _22_v0 = _compr_3;
    if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!(true))).length))).contains(_22_v0)) {
      _coll3.push([(_22_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("wcolpleqj"))).length)),!(true)]);
    }
  }
  return _coll3;
}()).length), new BigNumber((_dafny.Seq.of(false, !(false))).length))).length)));
    };
    static fm22(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(false)).IsSubsetOf(_dafny.Set.fromElements(false)),!((_module.D12.create_DC33(new BigNumber(-467), false, _dafny.MultiSet.fromElements(new BigNumber(-662)), false, true)).dtor_cf61));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-322)), function (_23_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length), new BigNumber(119), new BigNumber(382), new BigNumber(970), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false))).length)),false)).length))).length)))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-465),false)).length), new BigNumber(-733), new BigNumber(919)));
    };
    static fm24(p0, p1, globalState) {
      if (!(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fgnsevrhf")).length), new BigNumber(120))).cardinality()),true)).length))).Elements) {
          let _24_v0 = _compr_4;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("fgnsevrhf")).length), new BigNumber(120))).cardinality()),true)).length))).contains(_24_v0)) {
            _coll4.push([_module.__default.safeDivisionInt(_24_v0, new BigNumber(-51)),new BigNumber((_dafny.Seq.UnicodeFromString("uw")).length)]);
          }
        }
        return _coll4;
      }()).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-855),new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(221),(_dafny.ZERO).minus(new BigNumber((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(783)))).Elements) {
          let _25_v1 = _compr_5;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(783)))).contains(_25_v1)) {
            _coll5.push([_25_v1,new BigNumber((_dafny.Seq.of(new BigNumber(365))).length)]);
          }
        }
        return _coll5;
      }()).length))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(535),new BigNumber((_dafny.Seq.of(false)).length)))).length)))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(437))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(785))).length)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("eceq")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(773), new BigNumber(283))) {
            let _26_v2 = _compr_6;
            if (((new BigNumber(773)).isLessThanOrEqualTo(_26_v2)) && ((_26_v2).isLessThan(new BigNumber(283)))) {
              _coll6.push([_module.__default.safeDivisionInt(_26_v2, new BigNumber((_dafny.Seq.of(!(false), true)).length)),true]);
            }
          }
          return _coll6;
        }()).length), new BigNumber(537))).length)));
      }
    };
    static fm25(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tusbtf"), _dafny.Seq.UnicodeFromString("hbix")),new BigNumber(972));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return _module.D6.create_DC20(!((new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(-259))).length))).length))).length)).isLessThan(new BigNumber(654))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(422))).length), false);
    };
    static fm27(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true, false, false));
    };
    static fm28(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(false, !(true))).Difference(_dafny.MultiSet.fromElements(true))).Intersect(_dafny.MultiSet.fromElements(true, true));
    };
    static fm29(p0, p1, globalState) {
      return (function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.Map.Empty.slice().updateUnsafe((function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of _dafny.IntegerRange(new BigNumber(292), new BigNumber(523))) {
              let _27_v2 = _compr_9;
              if (((new BigNumber(292)).isLessThanOrEqualTo(_27_v2)) && ((_27_v2).isLessThan(new BigNumber(523)))) {
                _coll9.push([(_27_v2).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                  let _coll10 = new _dafny.Set();
                  for (const _compr_10 of _dafny.IntegerRange(new BigNumber(10), new BigNumber(160))) {
                    let _28_v3 = _compr_10;
                    if (((new BigNumber(10)).isLessThanOrEqualTo(_28_v3)) && ((_28_v3).isLessThan(new BigNumber(160)))) {
                      _coll10.add((_28_v3).minus(new BigNumber(428)));
                    }
                  }
                  return _coll10;
                }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))).length)),true]);
              }
            }
            return _coll9;
          }()),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)))).Keys.Elements) {
            let _29_v1 = _compr_8;
            if ((_dafny.Map.Empty.slice().updateUnsafe((function () {
              let _coll11 = new _dafny.Map();
              for (const _compr_11 of _dafny.IntegerRange(new BigNumber(292), new BigNumber(523))) {
                let _27_v2 = _compr_11;
                if (((new BigNumber(292)).isLessThanOrEqualTo(_27_v2)) && ((_27_v2).isLessThan(new BigNumber(523)))) {
                  _coll11.push([(_27_v2).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                    let _coll12 = new _dafny.Set();
                    for (const _compr_12 of _dafny.IntegerRange(new BigNumber(10), new BigNumber(160))) {
                      let _30_v3 = _compr_12;
                      if (((new BigNumber(10)).isLessThanOrEqualTo(_30_v3)) && ((_30_v3).isLessThan(new BigNumber(160)))) {
                        _coll12.add((_30_v3).minus(new BigNumber(428)));
                      }
                    }
                    return _coll12;
                  }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))).length)),true]);
                }
              }
              return _coll11;
            }()),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)))).contains(_29_v1)) {
              _coll8.push([_29_v1,new BigNumber(669)]);
            }
          }
          return _coll8;
        }()).Keys.Elements) {
          let _31_v0 = _compr_7;
          if ((function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe((function () {
              let _coll14 = new _dafny.Map();
              for (const _compr_14 of _dafny.IntegerRange(new BigNumber(292), new BigNumber(523))) {
                let _27_v2 = _compr_14;
                if (((new BigNumber(292)).isLessThanOrEqualTo(_27_v2)) && ((_27_v2).isLessThan(new BigNumber(523)))) {
                  _coll14.push([(_27_v2).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                    let _coll15 = new _dafny.Set();
                    for (const _compr_15 of _dafny.IntegerRange(new BigNumber(10), new BigNumber(160))) {
                      let _32_v3 = _compr_15;
                      if (((new BigNumber(10)).isLessThanOrEqualTo(_32_v3)) && ((_32_v3).isLessThan(new BigNumber(160)))) {
                        _coll15.add((_32_v3).minus(new BigNumber(428)));
                      }
                    }
                    return _coll15;
                  }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))).length)),true]);
                }
              }
              return _coll14;
            }()),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)))).Keys.Elements) {
              let _29_v1 = _compr_13;
              if ((_dafny.Map.Empty.slice().updateUnsafe((function () {
                let _coll16 = new _dafny.Map();
                for (const _compr_16 of _dafny.IntegerRange(new BigNumber(292), new BigNumber(523))) {
                  let _27_v2 = _compr_16;
                  if (((new BigNumber(292)).isLessThanOrEqualTo(_27_v2)) && ((_27_v2).isLessThan(new BigNumber(523)))) {
                    _coll16.push([(_27_v2).minus(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
                      let _coll17 = new _dafny.Set();
                      for (const _compr_17 of _dafny.IntegerRange(new BigNumber(10), new BigNumber(160))) {
                        let _33_v3 = _compr_17;
                        if (((new BigNumber(10)).isLessThanOrEqualTo(_33_v3)) && ((_33_v3).isLessThan(new BigNumber(160)))) {
                          _coll17.add((_33_v3).minus(new BigNumber(428)));
                        }
                      }
                      return _coll17;
                    }()).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length))).length)),true]);
                  }
                }
                return _coll16;
              }()),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)))).contains(_29_v1)) {
                _coll13.push([_29_v1,new BigNumber(669)]);
              }
            }
            return _coll13;
          }()).contains(_31_v0)) {
            _coll7.push([_31_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("lukayiqqc")).length))).length))]);
          }
        }
        return _coll7;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true, true, !(false), false)).length),true),new BigNumber(502)));
    };
    static fm30(p0, globalState) {
      if (true) {
        return (function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of (_dafny.Seq.of(_dafny.Seq.of(!(true), true), _dafny.Seq.of(false, !(true), false, false))).Elements) {
            let _34_v0 = _compr_18;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(!(true), true), _dafny.Seq.of(false, !(true), false, false)), _34_v0)) {
              _coll18.push([_34_v0,false]);
            }
          }
          return _coll18;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, !(!(false))),false));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),false);
      }
    };
    static fm31(p0, p1, globalState) {
      return _module.D10.create_DC27((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(false)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _module.D2.create_DC7(_dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.of(false))).length)))));
    };
    static fm33(p0, p1, globalState) {
      return _module.D10.create_DC28(false, (new BigNumber(59)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(933),true)).length)));
    };
    static fm34(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(false, true, true, false), _dafny.Seq.of(true, false), _dafny.Seq.of(!(true)))).Intersect(function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), function (_35_i0) {
          return _dafny.Seq.of(false, false, !(false));
        })).Elements) {
          let _36_v0 = _compr_19;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(609)), function (_35_i0) {
            return _dafny.Seq.of(false, false, !(false));
          }), _36_v0)) {
            _coll19.add(_36_v0);
          }
        }
        return _coll19;
      }());
    };
    static m9(p0, p1, globalState) {
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _37_v0;
      _37_v0 = _dafny.Seq.UnicodeFromString("dcrcctt");
      let _38_v1;
      _38_v1 = new BigNumber(-916);
      let _39_v2;
      _39_v2 = _dafny.Map.Empty.slice().updateUnsafe(_37_v0,_module.__default.fm12(_38_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), function (_40_i1) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _37_v0, globalState));
      let _41_v3;
      _41_v3 = _dafny.Seq.of(p1);
      let _42_v4;
      _42_v4 = _dafny.Set.fromElements((_41_v3)[_module.__default.safeIndex(new BigNumber(875), new BigNumber((_41_v3).length))], p1, p1);
      let _hi0 = new BigNumber((_42_v4).length);
      for (let _43_i0 = new BigNumber((_39_v2).length); _43_i0.isLessThan(_hi0); _43_i0 = _43_i0.plus(_dafny.ONE)) {
        (globalState).f17 = p1;
        let _44_v5;
        _44_v5 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.of(p1), _module.__default.safeIndex(_38_v1, new BigNumber((_dafny.Seq.of(p1)).length)), p0), _41_v3);
        if (!(!(!_dafny.Seq.contains((_44_v5)[_module.__default.safeIndex(_38_v1, new BigNumber((_44_v5).length))], !(false) || (p1))))) {
          let _45_v6;
          _45_v6 = _module.D0.create_DC1(p0);
          let _46_v7;
          let _nw0 = new _module.C2();
          _nw0.__ctor(_45_v6, p1);
          _46_v7 = _nw0;
          let _47_v8;
          let _nw1 = new _module.C5();
          _nw1.__ctor(_43_i0);
          _47_v8 = _nw1;
          _37_v0 = _37_v0;
          let _48_v9;
          let _init0 = ((_49_v8) => function (_50_i2) {
            return _dafny.Set.fromElements((_49_v8).f18);
          })(_47_v8);
          let _nw2 = Array((new BigNumber(9)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
            _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _48_v9 = _nw2;
          let _51_v10;
          _51_v10 = new _dafny.CodePoint('e'.codePointAt(0));
          let _52_v11;
          let _nw3 = Array((new BigNumber(25)).toNumber());
          _nw3[(_dafny.ZERO).toNumber()] = false;
          _nw3[(_dafny.ONE).toNumber()] = false;
          _nw3[(new BigNumber(2)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(3)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(4)).toNumber()] = p1;
          _nw3[(new BigNumber(5)).toNumber()] = !((_41_v3)[_module.__default.safeIndex(_43_i0, new BigNumber((_41_v3).length))]);
          _nw3[(new BigNumber(6)).toNumber()] = p0;
          _nw3[(new BigNumber(7)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(8)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(9)).toNumber()] = false;
          _nw3[(new BigNumber(10)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(11)).toNumber()] = false;
          _nw3[(new BigNumber(12)).toNumber()] = p1;
          _nw3[(new BigNumber(13)).toNumber()] = p1;
          _nw3[(new BigNumber(14)).toNumber()] = p0;
          _nw3[(new BigNumber(15)).toNumber()] = p1;
          _nw3[(new BigNumber(16)).toNumber()] = p0;
          _nw3[(new BigNumber(17)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(18)).toNumber()] = p0;
          _nw3[(new BigNumber(19)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(20)).toNumber()] = true;
          _nw3[(new BigNumber(21)).toNumber()] = _module.__default.fm0(p1, _51_v10, (_46_v7).f21, globalState);
          _nw3[(new BigNumber(22)).toNumber()] = p0;
          _nw3[(new BigNumber(23)).toNumber()] = (_46_v7).f21;
          _nw3[(new BigNumber(24)).toNumber()] = p0;
          _52_v11 = _nw3;
          let _53_v12;
          _53_v12 = _module.D1.create_DC4((_47_v8).f18, p1, _52_v11);
          let _rhs0 = _37_v0;
          let _rhs1 = _48_v9;
          let _rhs2 = _53_v12;
          let _rhs3 = _43_i0;
          _37_v0 = _rhs0;
          _48_v9 = _rhs1;
          _53_v12 = _rhs2;
          _38_v1 = _rhs3;
          let _54_v13;
          let _nw4 = new _module.C0();
          _nw4.__ctor(p1);
          _54_v13 = _nw4;
          let _55_v14;
          let _nw5 = new _module.C4();
          _nw5.__ctor((_46_v7).f21);
          _55_v14 = _nw5;
          let _56_v15;
          _56_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_54_v13, _54_v13, _54_v13)).length)),_55_v14);
          r2 = _module.__default.fm3(new BigNumber(((_56_v15).update(_43_i0, _55_v14)).length), (_dafny.ZERO).minus(_43_i0), p1, globalState);
        } else {
          let _57_v16;
          let _init1 = ((_58_v0) => function (_59_i3) {
            return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hi"), _58_v0);
          })(_37_v0);
          let _nw6 = Array((new BigNumber(26)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw6.length); _i0_1++) {
            _nw6[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _57_v16 = _nw6;
          let _index0 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_57_v16).length));
          (_57_v16)[_index0] = _dafny.Seq.Concat(_37_v0, _37_v0);
          let _60_v17;
          _60_v17 = _module.D0.create_DC1(p1);
          let _61_v18;
          let _nw7 = new _module.C2();
          _nw7.__ctor(_60_v17, p1);
          _61_v18 = _nw7;
          let _62_v19;
          let _init2 = ((_63_v3) => function (_64_i4) {
            return _63_v3;
          })(_41_v3);
          let _nw8 = Array((new BigNumber(14)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw8.length); _i0_2++) {
            _nw8[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _62_v19 = _nw8;
          let _65_v20;
          _65_v20 = _module.D14.create_DC41(_61_v18, _dafny.Seq.UnicodeFromString("sos"), _62_v19);
          let _pat_let_tv0 = _62_v19;
          let _pat_let_tv1 = _61_v18;
          let _66_v21;
          _66_v21 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let0_0) {
            return function (_67_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_68_dt__update_hcf74_h0) {
                  return function (_pat_let2_0) {
                    return function (_69_dt__update_hcf72_h0) {
                      return _module.D14.create_DC41(_69_dt__update_hcf72_h0, (_67_dt__update__tmp_h0).dtor_cf73, _68_dt__update_hcf74_h0);
                    }(_pat_let2_0);
                  }(_pat_let_tv1);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_65_v20)).dtor_cf73,_37_v0);
          let _70_v22;
          _70_v22 = _dafny.Set.fromElements(_38_v1);
          let _index1 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_57_v16).length));
          let _rhs4 = (((_66_v21).contains(_37_v0)) ? ((_66_v21).get(_37_v0)) : (_37_v0));
          let _rhs5 = (_61_v18).f21;
          let _rhs6 = (_61_v18).f21;
          let _rhs7 = ((_70_v22).Union(_70_v22)).IsSubsetOf(_70_v22);
          let _lhs0 = _57_v16;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_57_v16).length));
          let _lhs2 = globalState;
          let _lhs3 = globalState;
          let _lhs4 = globalState;
          _lhs0[_lhs1] = _rhs4;
          _lhs2.f17 = _rhs5;
          _lhs3.f11 = _rhs6;
          _lhs4.f15 = _rhs7;
          let _71_v23;
          _71_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _72_v24;
          _72_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_71_v23);
          _72_v24 = _72_v24;
          let _rhs8 = _module.__default.fm3(_module.__default.safeDivisionInt(_43_i0, _38_v1), (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(148)), function (_73_i5) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          })).length)).minus(_38_v1), p0, globalState);
          let _rhs9 = _43_i0;
          r2 = _rhs8;
          r2 = _rhs9;
          let _74_v25;
          let _nw9 = new _module.C4();
          _nw9.__ctor((_61_v18).f21);
          _74_v25 = _nw9;
          let _75_v26;
          let _out0;
          _out0 = (_61_v18).m5((p0) && (false), _41_v3, globalState);
          _75_v26 = _out0;
        }
        _38_v1 = _module.__default.safeModuloInt(new BigNumber(-352), new BigNumber((_dafny.Seq.of(_43_i0)).length));
        let _76_v27;
        let _init3 = ((_77_p1) => function (_78_i6) {
          return _77_p1;
        })(p1);
        let _nw10 = Array((new BigNumber(3)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw10.length); _i0_3++) {
          _nw10[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _76_v27 = _nw10;
        let _index2 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_76_v27).length));
        (_76_v27)[_index2] = p0;
        let _index3 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_76_v27).length));
        (_76_v27)[_index3] = p0;
      }
      let _79_v28;
      _79_v28 = _dafny.MultiSet.fromElements(p0);
      if ((_79_v28).IsSubsetOf((_dafny.MultiSet.fromElements(p0)).Intersect(_79_v28))) {
        let _80_v29;
        let _nw11 = Array((new BigNumber(24)).toNumber()).fill([]);
        _80_v29 = _nw11;
        let _81_v30;
        _81_v30 = _dafny.MultiSet.fromElements(_38_v1);
        let _82_v31;
        _82_v31 = _dafny.Map.Empty.slice().updateUnsafe(_80_v29,((_dafny.ZERO).minus(new BigNumber((_81_v30).cardinality()))).multipliedBy(_38_v1));
        _82_v31 = (_82_v31).update(_80_v29, _38_v1);
        let _83_v32;
        let _nw12 = new _module.C0();
        _nw12.__ctor(p0);
        _83_v32 = _nw12;
        let _84_v33;
        _84_v33 = _dafny.Set.fromElements(_83_v32);
        _38_v1 = new BigNumber((_84_v33).length);
        let _85_v34;
        let _nw13 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _85_v34 = _nw13;
        let _86_v35;
        let _nw14 = new _module.C4();
        _nw14.__ctor((_83_v32).f23);
        _86_v35 = _nw14;
        let _87_v36;
        _87_v36 = _dafny.Seq.of(_86_v35, _86_v35, _86_v35);
        let _index4 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_85_v34).length));
        (_85_v34)[_index4] = _87_v36;
        let _index5 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_85_v34).length));
        let _rhs10 = (((true) ? (new BigNumber((_dafny.Seq.UnicodeFromString("gvcj")).length)) : ((_dafny.ZERO).minus(_38_v1)))).isLessThan(_38_v1);
        let _rhs11 = !(new BigNumber(429)).isEqualTo(_38_v1);
        let _rhs12 = _87_v36;
        let _lhs5 = globalState;
        let _lhs6 = globalState;
        let _lhs7 = _85_v34;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_85_v34).length));
        _lhs5.f17 = _rhs10;
        _lhs6.f11 = _rhs11;
        _lhs7[_lhs8] = _rhs12;
        let _88_v37;
        _88_v37 = _module.D12.create_DC34();
        let _source2 = (((_83_v32).f23) ? (_88_v37) : (_88_v37));
        if (_source2.is_DC32) {
          let _89___mcc_h0 = (_source2).cf56;
          let _90_cf56 = _89___mcc_h0;
          let _91_v38;
          _91_v38 = _module.D0.create_DC1(p1);
          let _92_v39;
          let _nw15 = new _module.C2();
          _nw15.__ctor(_91_v38, (_83_v32).f23);
          _92_v39 = _nw15;
          r2 = _90_cf56;
          (globalState).f11 = true;
          let _93_v40;
          _93_v40 = _dafny.Map.Empty.slice().updateUnsafe((_92_v39).f21,_90_cf56);
          let _94_v41;
          let _nw16 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _94_v41 = _nw16;
          let _95_v42;
          _95_v42 = _module.D3.create_DC11(_93_v40, _90_cf56, _94_v41);
          let _rhs13 = _38_v1;
          let _rhs14 = (_dafny.ZERO).minus((new BigNumber(((_dafny.MultiSet.fromElements((_83_v32).f23, (_83_v32).f23)).Intersect(_dafny.MultiSet.fromElements((_83_v32).f23, p0))).cardinality())).plus(new BigNumber((_42_v4).length)));
          let _rhs15 = _95_v42;
          _90_cf56 = _rhs13;
          _38_v1 = _rhs14;
          _95_v42 = _rhs15;
        } else if (_source2.is_DC33) {
          let _96___mcc_h1 = (_source2).cf57;
          let _97___mcc_h2 = (_source2).cf58;
          let _98___mcc_h3 = (_source2).cf59;
          let _99___mcc_h4 = (_source2).cf60;
          let _100___mcc_h5 = (_source2).cf61;
          let _101_cf61 = _100___mcc_h5;
          let _102_cf60 = _99___mcc_h4;
          let _103_cf59 = _98___mcc_h3;
          let _104_cf58 = _97___mcc_h2;
          let _105_cf57 = _96___mcc_h1;
          r2 = _38_v1;
          let _106_v43;
          _106_v43 = _dafny.Map.Empty.slice().updateUnsafe(_83_v32,_module.__default.fm20(globalState));
          _42_v4 = ((_42_v4).Difference(_42_v4)).Union((((_106_v43).contains(_83_v32)) ? ((_106_v43).get(_83_v32)) : (_42_v4)));
          let _107_v44;
          _107_v44 = new _dafny.CodePoint('o'.codePointAt(0));
          let _108_v45;
          _108_v45 = _dafny.Map.Empty.slice().updateUnsafe(_107_v44,_107_v44);
          _108_v45 = _108_v45;
          let _109_v46;
          _109_v46 = _module.D11.create_DC29(_42_v4);
          let _110_v47;
          _110_v47 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_105_cf57,new BigNumber(((_109_v46).dtor_cf51).length)));
          let _111_v48;
          _111_v48 = _module.D0.create_DC1(_102_cf60);
          let _112_v49;
          let _nw17 = new _module.C2();
          _nw17.__ctor(_111_v48, _104_cf58);
          _112_v49 = _nw17;
          let _113_v50;
          let _init4 = ((_114_cf60) => function (_115_i7) {
            return _dafny.Seq.of(_114_cf60);
          })(_102_cf60);
          let _nw18 = Array((new BigNumber(9)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw18.length); _i0_4++) {
            _nw18[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _113_v50 = _nw18;
          let _116_v51;
          _116_v51 = _dafny.MultiSet.fromElements(_module.D14.create_DC41(_112_v49, _37_v0, _113_v50));
          let _117_v52;
          _117_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm27(_105_cf57, _105_cf57, new BigNumber(((_110_v47)[_module.__default.safeIndex(new BigNumber((_116_v51).cardinality()), new BigNumber((_110_v47).length))]).length), globalState)).length),new BigNumber(533));
          let _118_v53;
          _118_v53 = _dafny.Set.fromElements(_117_v52);
          let _119_v54;
          _119_v54 = _module.D6.create_DC20((_83_v32).f23, _105_cf57, p0);
          let _120_v55;
          _120_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(72));
          let _121_v56;
          _121_v56 = _dafny.Set.fromElements(_38_v1, new BigNumber((_module.__default.fm23((_83_v32).f23, _120_v55, _38_v1, _107_v44, globalState)).cardinality()));
          let _122_v57;
          _122_v57 = _dafny.MultiSet.fromElements(_120_v55);
          let _123_v58;
          _123_v58 = _dafny.Seq.of(new BigNumber(949));
          let _124_v59;
          let _nw19 = Array((new BigNumber(26)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = (_83_v32).fm14(new BigNumber((_118_v53).length), globalState);
          _nw19[(_dafny.ONE).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(2)).toNumber()] = _38_v1;
          _nw19[(new BigNumber(3)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(4)).toNumber()] = _38_v1;
          _nw19[(new BigNumber(5)).toNumber()] = (_119_v54).dtor_cf37;
          _nw19[(new BigNumber(6)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(_38_v1, _38_v1);
          _nw19[(new BigNumber(8)).toNumber()] = (((_83_v32).f23) ? (_38_v1) : (_38_v1));
          _nw19[(new BigNumber(9)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(10)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(11)).toNumber()] = (_module.D4.create_DC14(p0, _38_v1, _121_v56, _107_v44)).dtor_cf24;
          _nw19[(new BigNumber(12)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(13)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(14)).toNumber()] = _38_v1;
          _nw19[(new BigNumber(15)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(16)).toNumber()] = new BigNumber(((((_112_v49).f21) ? (_dafny.Seq.UnicodeFromString("kuj")) : (_37_v0))).length);
          _nw19[(new BigNumber(17)).toNumber()] = new BigNumber(12);
          _nw19[(new BigNumber(18)).toNumber()] = new BigNumber((((_122_v57).update(_120_v55, _module.__default.abs(new BigNumber((_123_v58).length)))).Union(_122_v57)).cardinality());
          _nw19[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_79_v28).cardinality()));
          _nw19[(new BigNumber(20)).toNumber()] = _38_v1;
          _nw19[(new BigNumber(21)).toNumber()] = new BigNumber((_37_v0).length);
          _nw19[(new BigNumber(22)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(23)).toNumber()] = _105_cf57;
          _nw19[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(_105_cf57);
          _nw19[(new BigNumber(25)).toNumber()] = _38_v1;
          _124_v59 = _nw19;
          let _125_v60;
          _125_v60 = _dafny.Map.Empty.slice().updateUnsafe((_105_cf57).plus(_38_v1),_124_v59);
          let _rhs16 = _38_v1;
          let _rhs17 = (((_125_v60).contains((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_105_cf57, _105_cf57)))) ? ((_125_v60).get((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_105_cf57, _105_cf57)))) : (_124_v59));
          r2 = _rhs16;
          _124_v59 = _rhs17;
        } else if (_source2.is_DC34) {
          let _126_v61;
          _126_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_38_v1);
          let _127_v62;
          let _nw20 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _127_v62 = _nw20;
          let _128_v63;
          _128_v63 = _module.D3.create_DC11(_126_v61, new BigNumber((_37_v0).length), _127_v62);
          let _129_v64;
          let _nw21 = new _module.C1();
          _nw21.__ctor((_128_v63).dtor_cf20);
          _129_v64 = _nw21;
          let _130_v65;
          _130_v65 = _dafny.Seq.of(_129_v64);
          (globalState).f15 = (_38_v1).isLessThan(new BigNumber((_dafny.Seq.Concat(_130_v65, _130_v65)).length));
          let _131_v66;
          _131_v66 = _dafny.Seq.of(_38_v1);
          let _132_v67;
          _132_v67 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), ((_133_p1, _134_v32) => function (_135_i8) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_133_p1,(_134_v32).f23))).length));
          })(p1, _83_v32)), _module.__default.safeIndex(_38_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(513)), ((_136_p1, _137_v32) => function (_138_i8) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_136_p1,(_137_v32).f23))).length));
          })(p1, _83_v32))).length)), _38_v1), _131_v66, _131_v66);
          let _139_v68;
          _139_v68 = _dafny.MultiSet.fromElements((_132_v67)[_module.__default.safeIndex(_38_v1, new BigNumber((_132_v67).length))], _131_v66);
          _139_v68 = ((_139_v68).Union(_139_v68)).Difference((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(648)), ((_140_v1) => function (_141_i9) {
            return _140_v1;
          })(_38_v1)))).Union(_139_v68));
          let _142_v69;
          let _init5 = ((_143_v32) => function (_144_i10) {
            return ((_143_v32).f23) || ((_143_v32).f23);
          })(_83_v32);
          let _nw22 = Array((new BigNumber(25)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw22.length); _i0_5++) {
            _nw22[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _142_v69 = _nw22;
          let _index6 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_142_v69).length));
          (_142_v69)[_index6] = (_83_v32).f23;
          let _145_v70;
          _145_v70 = _dafny.Set.fromElements(_41_v3);
          let _index7 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_142_v69).length));
          let _rhs18 = _86_v35;
          let _rhs19 = (_module.__default.fm34((_131_v66)[_module.__default.safeIndex(_38_v1, new BigNumber((_131_v66).length))], _38_v1, true, globalState)).IsSubsetOf(_145_v70);
          let _rhs20 = p0;
          let _lhs9 = _142_v69;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(922), new BigNumber((_142_v69).length));
          let _lhs11 = globalState;
          _86_v35 = _rhs18;
          _lhs9[_lhs10] = _rhs19;
          _lhs11.f11 = _rhs20;
          let _146_v71;
          _146_v71 = _module.D12.create_DC33(_38_v1, false, _81_v30, p1, p0);
          _38_v1 = (_146_v71).dtor_cf57;
        } else if (_source2.is_DC31) {
          let _147___mcc_h6 = (_source2).cf55;
          let _148_cf55 = _147___mcc_h6;
          let _149_v72;
          let _init6 = ((_150_v1) => function (_151_i11) {
            return _module.__default.safeDivisionInt(_151_i11, _150_v1);
          })(_38_v1);
          let _nw23 = Array((new BigNumber(22)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw23.length); _i0_6++) {
            _nw23[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _149_v72 = _nw23;
          let _152_v73;
          let _nw24 = new _module.C1();
          _nw24.__ctor(((p1) ? (_149_v72) : (_149_v72)));
          _152_v73 = _nw24;
          let _153_v74;
          _153_v74 = _module.D0.create_DC1((_83_v32).f23);
          let _154_v75;
          let _nw25 = new _module.C2();
          _nw25.__ctor(_153_v74, ((_83_v32).f23) === (p1));
          _154_v75 = _nw25;
          _154_v75 = _154_v75;
          (globalState).f11 = (_83_v32).f23;
          let _155_v76;
          _155_v76 = _module.D12.create_DC32(_38_v1);
          let _156_v77;
          _156_v77 = _dafny.Map.Empty.slice().updateUnsafe(_155_v76,_149_v72);
          _149_v72 = (((_156_v77).contains(_155_v76)) ? ((_156_v77).get(_155_v76)) : ((_152_v73).f22));
        } else {
          let _157___mcc_h7 = (_source2).cf62;
          let _158_cf62 = _157___mcc_h7;
          (globalState).f11 = !(p0) || ((_83_v32).f23);
          let _159_v78;
          let _nw26 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _159_v78 = _nw26;
          let _160_v79;
          let _nw27 = new _module.C1();
          _nw27.__ctor(_159_v78);
          _160_v79 = _nw27;
          let _161_v80;
          _161_v80 = _dafny.Map.Empty.slice().updateUnsafe(_160_v79,(_dafny.ZERO).minus(_38_v1));
          _38_v1 = (_dafny.ZERO).minus(new BigNumber(((_161_v80).Merge((_161_v80).update(_160_v79, new BigNumber(161)))).length));
          let _162_v81;
          _162_v81 = new _dafny.CodePoint('x'.codePointAt(0));
          _162_v81 = _162_v81;
          let _163_v82;
          _163_v82 = _dafny.Map.Empty.slice().updateUnsafe(p1,_38_v1);
          _163_v82 = (_163_v82).update((_83_v32).f23, (_dafny.ZERO).minus(_module.__default.safeModuloInt(_38_v1, _38_v1)));
        }
        (globalState).f11 = p1;
      } else {
        let _164_v83;
        let _nw28 = Array((new BigNumber(8)).toNumber()).fill(false);
        _164_v83 = _nw28;
        let _index8 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_164_v83).length));
        (_164_v83)[_index8] = p1;
        let _index9 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_164_v83).length));
        (_164_v83)[_index9] = !((_38_v1).isLessThan(_38_v1));
        let _index10 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_164_v83).length));
        (_164_v83)[_index10] = (_164_v83)[_module.__default.safeIndex(new BigNumber(434), new BigNumber((_164_v83).length))];
        let _165_v84;
        _165_v84 = _dafny.Seq.of(_37_v0, _37_v0, _37_v0, _37_v0);
        let _166_v85;
        _166_v85 = _dafny.Seq.of(_165_v84);
        _165_v84 = (_166_v85)[_module.__default.safeIndex(new BigNumber((_37_v0).length), new BigNumber((_166_v85).length))];
        let _167_v86;
        _167_v86 = _dafny.Seq.of((_dafny.ZERO).minus(_38_v1), new BigNumber(256), new BigNumber((_37_v0).length), _38_v1);
        (globalState).f11 = _dafny.areEqual(_dafny.Seq.update(_module.__default.fm19(_dafny.Seq.of(_38_v1), p1, globalState), _module.__default.safeIndex(_38_v1, new BigNumber((_module.__default.fm19(_dafny.Seq.of(_38_v1), p1, globalState)).length)), _38_v1), _167_v86);
        let _168_v87;
        _168_v87 = _dafny.Map.Empty.slice().updateUnsafe(p1,_38_v1);
        _168_v87 = (_168_v87).Merge((_168_v87).Merge(_168_v87));
      }
      _41_v3 = _41_v3;
      (globalState).f15 = p0;
      let _169_v88;
      let _nw29 = Array((new BigNumber(29)).toNumber());
      _169_v88 = _nw29;
      let _170_v89;
      let _nw30 = Array((new BigNumber(8)).toNumber());
      _nw30[(_dafny.ZERO).toNumber()] = _38_v1;
      _nw30[(_dafny.ONE).toNumber()] = _38_v1;
      _nw30[(new BigNumber(2)).toNumber()] = _38_v1;
      _nw30[(new BigNumber(3)).toNumber()] = _38_v1;
      _nw30[(new BigNumber(4)).toNumber()] = _38_v1;
      _nw30[(new BigNumber(5)).toNumber()] = _38_v1;
      _nw30[(new BigNumber(6)).toNumber()] = _38_v1;
      _nw30[(new BigNumber(7)).toNumber()] = _38_v1;
      _170_v89 = _nw30;
      let _171_v90;
      _171_v90 = _module.D13.create_DC38(_module.__default.fm12(new BigNumber(-248), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-943)), function (_172_i12) {
  return new _dafny.CodePoint('u'.codePointAt(0));
}), _37_v0, globalState), p0, (_dafny.ZERO).minus(_38_v1), _170_v89);
      let _173_v91;
      let _nw31 = new _module.C1();
      _nw31.__ctor((_171_v90).dtor_cf69);
      _173_v91 = _nw31;
      let _index11 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_169_v88).length));
      (_169_v88)[_index11] = _173_v91;
      let _index12 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_169_v88).length));
      (_169_v88)[_index12] = _173_v91;
      (globalState).f11 = (_38_v1).isEqualTo((new BigNumber((_dafny.Seq.of(false, p1)).length)).multipliedBy(_38_v1));
      r0 = !((_38_v1).isLessThan(_38_v1));
      r1 = p1;
      r2 = _38_v1;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _174_v0;
      _174_v0 = true;
      let _175_v1;
      _175_v1 = _dafny.Seq.of(_174_v0, _174_v0);
      let _176_v2;
      _176_v2 = new BigNumber(644);
      let _177_v3;
      let _nw32 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _177_v3 = _nw32;
      let _178_v4;
      let _nw33 = Array((new BigNumber(28)).toNumber());
      _nw33[(_dafny.ZERO).toNumber()] = _177_v3;
      _nw33[(_dafny.ONE).toNumber()] = _177_v3;
      _nw33[(new BigNumber(2)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(3)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(4)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(5)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(6)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(7)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(8)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(9)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(10)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(11)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(12)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(13)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(14)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(15)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(16)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(17)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(18)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(19)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(20)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(21)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(22)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(23)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(24)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(25)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(26)).toNumber()] = _177_v3;
      _nw33[(new BigNumber(27)).toNumber()] = _177_v3;
      _178_v4 = _nw33;
      let _179_v5;
      _179_v5 = _module.D0.create_DC1(_174_v0);
      let _180_v6;
      _180_v6 = _dafny.Seq.UnicodeFromString("df");
      let _181_v7;
      _181_v7 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_176_v2);
      let _182_globalState;
      let _nw34 = new _module.GlobalState();
      _nw34.__ctor(new _dafny.CodePoint('x'.codePointAt(0)), _dafny.Seq.update(_175_v1, _module.__default.safeIndex(_176_v2, new BigNumber((_175_v1).length)), _174_v0), _178_v4, false, _dafny.Seq.Concat(_175_v1, _dafny.Seq.of(_174_v0)), new _dafny.CodePoint('e'.codePointAt(0)), new BigNumber(-835), false, (((_179_v5).dtor_cf1) ? (_180_v6) : (_180_v6)), false, true, false, _181_v7, true, new BigNumber(-305), true, _dafny.Seq.Concat(_175_v1, _175_v1), false);
      _182_globalState = _nw34;
      let _183_v8;
      _183_v8 = new _dafny.CodePoint('c'.codePointAt(0));
      if (_module.__default.fm0(_174_v0, _183_v8, _174_v0, _182_globalState)) {
        let _184_v9;
        _184_v9 = _dafny.Map.Empty.slice().updateUnsafe(false,_174_v0);
        let _185_v10;
        _185_v10 = _dafny.Seq.of(_176_v2);
        let _186_v11;
        _186_v11 = _module.D0.create_DC0(_174_v0);
        let _187_v12;
        let _nw35 = Array((new BigNumber(15)).toNumber());
        _nw35[(_dafny.ZERO).toNumber()] = (_175_v1)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_175_v1).length))];
        _nw35[(_dafny.ONE).toNumber()] = _174_v0;
        _nw35[(new BigNumber(2)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(3)).toNumber()] = false;
        _nw35[(new BigNumber(4)).toNumber()] = false;
        _nw35[(new BigNumber(5)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(6)).toNumber()] = !(_174_v0) || (_174_v0);
        _nw35[(new BigNumber(7)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(8)).toNumber()] = (_184_v9).contains(true);
        _nw35[(new BigNumber(9)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(10)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_185_v10, _185_v10);
        _nw35[(new BigNumber(12)).toNumber()] = _174_v0;
        _nw35[(new BigNumber(13)).toNumber()] = (_module.D0.create_DC1((_186_v11).dtor_cf0)).dtor_cf1;
        _nw35[(new BigNumber(14)).toNumber()] = _174_v0;
        _187_v12 = _nw35;
        let _index13 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_187_v12).length));
        (_187_v12)[_index13] = _174_v0;
        let _index14 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_187_v12).length));
        (_187_v12)[_index14] = _174_v0;
        let _rhs21 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(829)), ((_188_v8) => function (_189_i0) {
          return _188_v8;
        })(_183_v8));
        let _rhs22 = new BigNumber(-916);
        _180_v6 = _rhs21;
        _176_v2 = _rhs22;
        let _index15 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_187_v12).length));
        (_187_v12)[_index15] = (_174_v0) || ((_187_v12)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_187_v12).length))]);
        let _190_v13;
        let _nw36 = new _module.C3();
        _nw36.__ctor();
        _190_v13 = _nw36;
        let _191_v14;
        _191_v14 = _module.D6.create_DC20(_174_v0, _176_v2, _174_v0);
        let _192_v15;
        let _nw37 = new _module.C2();
        _nw37.__ctor(_179_v5, (_191_v14).dtor_cf36);
        _192_v15 = _nw37;
        let _193_v16;
        let _nw38 = Array((new BigNumber(9)).toNumber()).fill([]);
        _193_v16 = _nw38;
        let _index16 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_193_v16).length));
        (_193_v16)[_index16] = ((_174_v0) ? (_187_v12) : (_187_v12));
        let _194_v17;
        let _nw39 = Array((new BigNumber(8)).toNumber());
        _194_v17 = _nw39;
        let _195_v18;
        let _nw40 = new _module.C4();
        _nw40.__ctor((_187_v12)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_187_v12).length))]);
        _195_v18 = _nw40;
        let _index17 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_194_v17).length));
        (_194_v17)[_index17] = _195_v18;
        let _196_v19;
        _196_v19 = _dafny.Seq.of(_192_v15, _192_v15, _192_v15);
        let _index18 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_193_v16).length));
        let _index19 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_194_v17).length));
        let _rhs23 = _187_v12;
        let _rhs24 = (_196_v19)[_module.__default.safeIndex(_176_v2, new BigNumber((_196_v19).length))];
        let _rhs25 = _187_v12;
        let _rhs26 = _174_v0;
        let _rhs27 = _195_v18;
        let _lhs12 = _193_v16;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_193_v16).length));
        let _lhs14 = _182_globalState;
        let _lhs15 = _194_v17;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_194_v17).length));
        _187_v12 = _rhs23;
        _192_v15 = _rhs24;
        _lhs12[_lhs13] = _rhs25;
        _lhs14.f17 = _rhs26;
        _lhs15[_lhs16] = _rhs27;
      } else {
        let _197_v20;
        let _init7 = ((_198_v8) => function (_199_i1) {
          return _198_v8;
        })(_183_v8);
        let _nw41 = Array((new BigNumber(13)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw41.length); _i0_7++) {
          _nw41[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _197_v20 = _nw41;
        let _index20 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length));
        (_197_v20)[_index20] = _183_v8;
        let _index21 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length));
        (_197_v20)[_index21] = _183_v8;
        let _200_v21;
        _200_v21 = _dafny.Map.Empty.slice().updateUnsafe(_174_v0,(_dafny.ZERO).minus(_176_v2));
        let _201_v22;
        _201_v22 = _module.D13.create_DC38((_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))], (new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_202_v8) => function (_203_i2) {
  return _202_v8;
})(_183_v8)), _module.__default.safeIndex(new BigNumber(701), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), ((_204_v8) => function (_205_i2) {
  return _204_v8;
})(_183_v8))).length)), (_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))])).length)).isLessThan(new BigNumber((_200_v21).length)), (_176_v2).multipliedBy((_dafny.ZERO).minus(_176_v2)), _177_v3);
        _201_v22 = _201_v22;
        let _206_v23;
        _206_v23 = _dafny.Seq.of((_176_v2).minus(new BigNumber((_200_v21).length)), (_176_v2).minus(new BigNumber(409)), _module.__default.safeModuloInt(_176_v2, _176_v2), (_176_v2).minus(_module.__default.fm9(_176_v2, _176_v2, false, _182_globalState)));
        _206_v23 = _dafny.Seq.Concat(_206_v23, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(963)), ((_207_v2) => function (_208_i3) {
          return _207_v2;
        })(_176_v2)), _dafny.Seq.of(_176_v2)));
        let _209_v24;
        _209_v24 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_180_v6, _180_v6), _module.__default.safeIndex(_176_v2, new BigNumber((_dafny.Seq.Concat(_180_v6, _180_v6)).length)), (_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))])).length), _176_v2, (new BigNumber((_206_v23).length)).multipliedBy(_176_v2), _176_v2, (_module.D3.create_DC11(_200_v21, _176_v2, _177_v3)).dtor_cf19);
        let _210_v25;
        _210_v25 = _dafny.Seq.of(_209_v24);
        _209_v24 = (_210_v25)[_module.__default.safeIndex(_176_v2, new BigNumber((_210_v25).length))];
        let _211_v26;
        let _nw42 = new _module.C2();
        _nw42.__ctor(_179_v5, _174_v0);
        _211_v26 = _nw42;
        let _212_v27;
        let _init8 = ((_213_v1) => function (_214_i4) {
          return _213_v1;
        })(_175_v1);
        let _nw43 = Array((new BigNumber(11)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw43.length); _i0_8++) {
          _nw43[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _212_v27 = _nw43;
        let _215_v28;
        _215_v28 = _module.D14.create_DC41(_211_v26, _180_v6, _212_v27);
        let _source3 = _215_v28;
        if (_source3.is_DC41) {
          let _216___mcc_h0 = (_source3).cf72;
          let _217___mcc_h1 = (_source3).cf73;
          let _218___mcc_h2 = (_source3).cf74;
          let _219_cf74 = _218___mcc_h2;
          let _220_cf73 = _217___mcc_h1;
          let _221_cf72 = _216___mcc_h0;
          let _222_v29;
          let _nw44 = Array((new BigNumber(21)).toNumber()).fill(_dafny.MultiSet.Empty);
          _222_v29 = _nw44;
          let _223_v31;
          _223_v31 = _module.D4.create_DC14((_221_cf72).f21, _176_v2, function () {
  let _coll20 = new _dafny.Set();
  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(602), new BigNumber(158))) {
    let _224_v30 = _compr_20;
    if (((new BigNumber(602)).isLessThanOrEqualTo(_224_v30)) && ((_224_v30).isLessThan(new BigNumber(158)))) {
      _coll20.add((_224_v30).plus(_176_v2));
    }
  }
  return _coll20;
}(), (_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))]);
          let _225_v32;
          _225_v32 = _module.D10.create_DC28((_211_v26).f21, new BigNumber((_200_v21).length));
          let _226_v33;
          _226_v33 = _dafny.Map.Empty.slice().updateUnsafe((_225_v32).dtor_cf49,(_211_v26).f21);
          let _227_v34;
          _227_v34 = _dafny.Map.Empty.slice().updateUnsafe((_221_cf72).f21,(((_226_v33).contains(_module.__default.fm0((_211_v26).f21, (_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))], (_211_v26).f21, _182_globalState))) ? ((_226_v33).get(_module.__default.fm0((_211_v26).f21, (_197_v20)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_197_v20).length))], (_211_v26).f21, _182_globalState))) : (!((_211_v26).f21))));
          let _228_v35;
          let _nw45 = Array((new BigNumber(19)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = !((_223_v31).dtor_cf23);
          _nw45[(_dafny.ONE).toNumber()] = (_221_cf72).f21;
          _nw45[(new BigNumber(2)).toNumber()] = (_221_cf72).f21;
          _nw45[(new BigNumber(3)).toNumber()] = true;
          _nw45[(new BigNumber(4)).toNumber()] = true;
          _nw45[(new BigNumber(5)).toNumber()] = (_211_v26).f21;
          _nw45[(new BigNumber(6)).toNumber()] = (_211_v26).f21;
          _nw45[(new BigNumber(7)).toNumber()] = !((((_227_v34).contains((_211_v26).f21)) ? ((_227_v34).get((_211_v26).f21)) : ((_211_v26).f21)));
          _nw45[(new BigNumber(8)).toNumber()] = !((_221_cf72).f21);
          _nw45[(new BigNumber(9)).toNumber()] = (_211_v26).f21;
          _nw45[(new BigNumber(10)).toNumber()] = _174_v0;
          _nw45[(new BigNumber(11)).toNumber()] = false;
          _nw45[(new BigNumber(12)).toNumber()] = (_221_cf72).f21;
          _nw45[(new BigNumber(13)).toNumber()] = _174_v0;
          _nw45[(new BigNumber(14)).toNumber()] = (_211_v26).f21;
          _nw45[(new BigNumber(15)).toNumber()] = true;
          _nw45[(new BigNumber(16)).toNumber()] = _174_v0;
          _nw45[(new BigNumber(17)).toNumber()] = !((_211_v26).f21);
          _nw45[(new BigNumber(18)).toNumber()] = !((_211_v26).f21);
          _228_v35 = _nw45;
          let _229_v36;
          _229_v36 = _dafny.MultiSet.fromElements(_228_v35, _228_v35, _228_v35);
          let _index22 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_222_v29).length));
          (_222_v29)[_index22] = _229_v36;
          let _index23 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_222_v29).length));
          (_222_v29)[_index23] = _229_v36;
          let _230_v37;
          _230_v37 = _module.D3.create_DC10(_176_v2);
          let _231_v38;
          _231_v38 = _dafny.MultiSet.fromElements(_230_v37);
          let _232_v39;
          _232_v39 = _dafny.MultiSet.fromElements(!(true), (_211_v26).f21, false, !(_231_v38).equals(_231_v38));
          _232_v39 = _232_v39;
          let _233_v40;
          let _234_v41;
          let _out1;
          let _out2;
          let _outcollector0 = (_221_cf72).m1((_223_v31).dtor_cf23, (((_200_v21).contains((_211_v26).f21)) ? ((_200_v21).get((_211_v26).f21)) : (new BigNumber(-498))), false, (_module.D1.create_DC4(new BigNumber(812), (_211_v26).f21, _228_v35)).dtor_cf4, _182_globalState);
          _out1 = _outcollector0[0];
          _out2 = _outcollector0[1];
          _233_v40 = _out1;
          _234_v41 = _out2;
          let _235_v42;
          let _init9 = ((_236_cf72) => function (_237_i5) {
            return _dafny.MultiSet.fromElements((_236_cf72).f21);
          })(_221_cf72);
          let _nw46 = Array((new BigNumber(27)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw46.length); _i0_9++) {
            _nw46[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _235_v42 = _nw46;
          let _rhs28 = _234_v41;
          let _rhs29 = new _dafny.CodePoint('j'.codePointAt(0));
          let _rhs30 = _235_v42;
          let _lhs17 = _182_globalState;
          _lhs17.f15 = _rhs28;
          _183_v8 = _rhs29;
          _235_v42 = _rhs30;
        } else {
          let _238___mcc_h3 = (_source3).cf71;
          let _239_cf71 = _238___mcc_h3;
          let _240_v43;
          let _nw47 = new _module.C2();
          _nw47.__ctor(_179_v5, (_211_v26).f21);
          _240_v43 = _nw47;
          _240_v43 = ((false) ? (_240_v43) : (_240_v43));
          let _241_v44;
          let _nw48 = Array((new BigNumber(26)).toNumber()).fill(false);
          _241_v44 = _nw48;
          _241_v44 = _241_v44;
          let _242_v45;
          let _243_v46;
          let _244_v47;
          let _245_v48;
          let _out3;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector1 = (_211_v26).m6(_176_v2, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-465)), ((_246_v8) => function (_247_i6) {
            return _246_v8;
          })(_183_v8)), _180_v6), (((_211_v26).f21) ? ((_211_v26).f21) : ((_211_v26).f21)), _182_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _out5 = _outcollector1[2];
          _out6 = _outcollector1[3];
          _242_v45 = _out3;
          _243_v46 = _out4;
          _244_v47 = _out5;
          _245_v48 = _out6;
          let _index24 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_177_v3).length));
          (_177_v3)[_index24] = _176_v2;
          let _index25 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_177_v3).length));
          (_177_v3)[_index25] = _176_v2;
        }
      }
      let _248_v49;
      _248_v49 = _module.D12.create_DC35(_module.D12.create_DC34());
      let _249_v50;
      _249_v50 = _module.D12.create_DC35(_248_v49);
      let _250_v51;
      _250_v51 = _dafny.Map.Empty.slice().updateUnsafe(_249_v50,_176_v2);
      _250_v51 = (_250_v51).update(_249_v50, _176_v2);
      let _251_v52;
      _251_v52 = _dafny.MultiSet.fromElements(!(_174_v0), _174_v0);
      let _252_v53;
      _252_v53 = _dafny.Seq.of(_module.D10.create_DC27(_251_v52));
      let _253_v54;
      _253_v54 = _dafny.Seq.of(_176_v2);
      let _hi1 = (_253_v54)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_253_v54).length))];
      for (let _254_i7 = (_dafny.ZERO).minus(new BigNumber((_252_v53).length)); _254_i7.isLessThan(_hi1); _254_i7 = _254_i7.plus(_dafny.ONE)) {
        let _255_v55;
        _255_v55 = _dafny.Map.Empty.slice().updateUnsafe(_174_v0,_183_v8);
        (_182_globalState).f11 = ((_174_v0) ? (_module.__default.fm0(_174_v0, _module.__default.fm12(new BigNumber((_255_v55).length), _180_v6, _180_v6, _182_globalState), _174_v0, _182_globalState)) : (_174_v0));
        let _256_v56;
        let _257_v57;
        let _258_v58;
        let _out7;
        let _out8;
        let _out9;
        let _outcollector2 = _module.__default.m9(_174_v0, ((_174_v0) ? (false) : (false)), _182_globalState);
        _out7 = _outcollector2[0];
        _out8 = _outcollector2[1];
        _out9 = _outcollector2[2];
        _256_v56 = _out7;
        _257_v57 = _out8;
        _258_v58 = _out9;
        _183_v8 = (_180_v6)[_module.__default.safeIndex(_258_v58, new BigNumber((_180_v6).length))];
        let _259_v59;
        let _nw49 = Array((new BigNumber(13)).toNumber()).fill(false);
        _259_v59 = _nw49;
        let _260_v60;
        _260_v60 = _dafny.Seq.of(_259_v59);
        if (_dafny.Seq.IsProperPrefixOf(_260_v60, _dafny.Seq.Concat(_260_v60, _260_v60))) {
          let _261_v61;
          _261_v61 = _dafny.Map.Empty.slice().updateUnsafe(_183_v8,_module.__default.fm24(_256_v56, _258_v58, _182_globalState));
          let _262_v62;
          _262_v62 = _dafny.Map.Empty.slice().updateUnsafe(_257_v57,_254_i7);
          let _263_v63;
          _263_v63 = _module.D3.create_DC11(_262_v62, new BigNumber(-703), _177_v3);
          let _264_v64;
          let _nw50 = Array((new BigNumber(4)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = (((_261_v61).contains(_183_v8)) ? ((_261_v61).get(_183_v8)) : ((_263_v63).dtor_cf18));
          _nw50[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_174_v0,_176_v2);
          _nw50[(new BigNumber(2)).toNumber()] = _262_v62;
          _nw50[(new BigNumber(3)).toNumber()] = _262_v62;
          _264_v64 = _nw50;
          let _index26 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_264_v64).length));
          (_264_v64)[_index26] = _262_v62;
          let _index27 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_264_v64).length));
          (_264_v64)[_index27] = _262_v62;
          let _265_v65;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_module.__default.fm0(_174_v0, _183_v8, _257_v57, _182_globalState));
          _265_v65 = _nw51;
          let _266_v66;
          _266_v66 = _dafny.Map.Empty.slice().updateUnsafe(_251_v52,_180_v6);
          let _267_v67;
          let _268_v68;
          let _269_v69;
          let _out10;
          let _out11;
          let _out12;
          let _outcollector3 = _module.__default.m9((_265_v65).fm15(_256_v56, _257_v57, (_dafny.ZERO).minus(_258_v58), _262_v62, _182_globalState), _dafny.Seq.IsPrefixOf(_dafny.Seq.update((((_266_v66).contains(_dafny.MultiSet.fromElements(!(_257_v57)))) ? ((_266_v66).get(_dafny.MultiSet.fromElements(!(_257_v57)))) : (_180_v6)), _module.__default.safeIndex(_254_i7, new BigNumber(((((_266_v66).contains(_dafny.MultiSet.fromElements(!(_257_v57)))) ? ((_266_v66).get(_dafny.MultiSet.fromElements(!(_257_v57)))) : (_180_v6))).length)), _183_v8), _180_v6), _182_globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _out12 = _outcollector3[2];
          _267_v67 = _out10;
          _268_v68 = _out11;
          _269_v69 = _out12;
          let _index28 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_177_v3).length));
          (_177_v3)[_index28] = (_258_v58).multipliedBy(_258_v58);
          let _index29 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_177_v3).length));
          (_177_v3)[_index29] = _176_v2;
          let _270_v70;
          let _271_v71;
          let _272_v72;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector4 = _module.__default.m9((_265_v65).f23, _dafny.Seq.IsPrefixOf(_180_v6, _180_v6), _182_globalState);
          _out13 = _outcollector4[0];
          _out14 = _outcollector4[1];
          _out15 = _outcollector4[2];
          _270_v70 = _out13;
          _271_v71 = _out14;
          _272_v72 = _out15;
        } else {
          _176_v2 = (_254_i7).plus(_module.__default.safeDivisionInt(new BigNumber(873), _254_i7));
          _174_v0 = true;
          let _273_v73;
          _273_v73 = _dafny.Map.Empty.slice().updateUnsafe((_254_i7).minus((_dafny.ZERO).minus(_176_v2)),((_dafny.ZERO).minus(_258_v58)).isLessThan(_258_v58));
          _273_v73 = _273_v73;
          (_182_globalState).f15 = _256_v56;
          let _274_v74;
          let _275_v75;
          let _276_v76;
          let _out16;
          let _out17;
          let _out18;
          let _outcollector5 = _module.__default.m9(_174_v0, _256_v56, _182_globalState);
          _out16 = _outcollector5[0];
          _out17 = _outcollector5[1];
          _out18 = _outcollector5[2];
          _274_v74 = _out16;
          _275_v75 = _out17;
          _276_v76 = _out18;
        }
      }
      let _277_v77;
      let _nw52 = new _module.C0();
      _nw52.__ctor(_174_v0);
      _277_v77 = _nw52;
      let _278_v78;
      _278_v78 = _dafny.MultiSet.fromElements(_277_v77);
      let _rhs31 = (_174_v0) === ((_277_v77).f23);
      let _rhs32 = _278_v78;
      let _lhs18 = _182_globalState;
      _lhs18.f11 = _rhs31;
      _278_v78 = _rhs32;
      let _279_v79;
      _279_v79 = _dafny.Map.Empty.slice().updateUnsafe((_277_v77).f23,_176_v2);
      if ((_277_v77).fm15(_174_v0, _174_v0, _176_v2, _279_v79, _182_globalState)) {
        let _280_v80;
        let _init10 = ((_281_v8) => function (_282_i8) {
          return _281_v8;
        })(_183_v8);
        let _nw53 = Array((new BigNumber(19)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw53.length); _i0_10++) {
          _nw53[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _280_v80 = _nw53;
        let _index30 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_280_v80).length));
        (_280_v80)[_index30] = _183_v8;
        let _index31 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_280_v80).length));
        (_280_v80)[_index31] = _183_v8;
        let _283_v81;
        let _nw54 = new _module.C2();
        _nw54.__ctor(_179_v5, (_277_v77).f23);
        _283_v81 = _nw54;
        let _284_v82;
        let _285_v83;
        let _286_v84;
        let _287_v85;
        let _out19;
        let _out20;
        let _out21;
        let _out22;
        let _outcollector6 = (_283_v81).m6(_module.__default.safeDivisionInt(_176_v2, _176_v2), _180_v6, (_277_v77).f23, _182_globalState);
        _out19 = _outcollector6[0];
        _out20 = _outcollector6[1];
        _out21 = _outcollector6[2];
        _out22 = _outcollector6[3];
        _284_v82 = _out19;
        _285_v83 = _out20;
        _286_v84 = _out21;
        _287_v85 = _out22;
        let _288_v86;
        _288_v86 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_dafny.ZERO).minus(_176_v2));
        if (((_288_v86).contains(_183_v8)) || ((_285_v83) || (_284_v82))) {
          let _289_v87;
          _289_v87 = _module.D12.create_DC34();
          let _290_v88;
          _290_v88 = _dafny.Map.Empty.slice().updateUnsafe(_289_v87,_277_v77);
          _290_v88 = (_290_v88).update(_289_v87, _277_v77);
          let _291_v89;
          _291_v89 = _dafny.Set.fromElements(_253_v54, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-160)), ((_292_v2) => function (_293_i9) {
            return _292_v2;
          })(_176_v2)));
          let _294_v90;
          _294_v90 = _dafny.Seq.of(_291_v89, _291_v89, _291_v89);
          (_182_globalState).f17 = (_291_v89).IsProperSubsetOf((_294_v90)[_module.__default.safeIndex(_176_v2, new BigNumber((_294_v90).length))]);
          _176_v2 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_280_v80)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_280_v80).length))],_251_v52)).length);
          let _295_v91;
          let _nw55 = Array((new BigNumber(4)).toNumber()).fill([]);
          _295_v91 = _nw55;
          let _296_v92;
          _296_v92 = _dafny.MultiSet.fromElements(_176_v2, _176_v2);
          let _297_v93;
          _297_v93 = _dafny.Seq.of(_253_v54);
          let _298_v94;
          _298_v94 = _module.D8.create_DC23(_296_v92);
          let _pat_let_tv2 = _296_v92;
          let _299_v95;
          let _nw56 = Array((new BigNumber(5)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _296_v92;
          _nw56[(_dafny.ONE).toNumber()] = _dafny.MultiSet.FromArray((_297_v93)[_module.__default.safeIndex(_176_v2, new BigNumber((_297_v93).length))]);
          _nw56[(new BigNumber(2)).toNumber()] = (function (_pat_let3_0) {
            return function (_300_dt__update__tmp_h0) {
              return function (_pat_let4_0) {
                return function (_301_dt__update_hcf41_h0) {
                  return _module.D8.create_DC23(_301_dt__update_hcf41_h0);
                }(_pat_let4_0);
              }(_pat_let_tv2);
            }(_pat_let3_0);
          }(_298_v94)).dtor_cf41;
          _nw56[(new BigNumber(3)).toNumber()] = _296_v92;
          _nw56[(new BigNumber(4)).toNumber()] = _296_v92;
          _299_v95 = _nw56;
          let _index32 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_295_v91).length));
          (_295_v91)[_index32] = _299_v95;
          let _index33 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_295_v91).length));
          (_295_v91)[_index33] = _299_v95;
          (_182_globalState).f15 = _284_v82;
        } else {
          let _302_v96;
          let _out23;
          _out23 = (_283_v81).m5(!(!(true) || ((_277_v77).f23)), _175_v1, _182_globalState);
          _302_v96 = _out23;
          let _303_v97;
          _303_v97 = _module.D9.create_DC25(_176_v2, (_283_v81).f21, _176_v2, _277_v77);
          let _304_v98;
          _304_v98 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_303_v97);
          let _305_v99;
          let _nw57 = new _module.C4();
          _nw57.__ctor((_283_v81).f21);
          _305_v99 = _nw57;
          let _306_v100;
          _306_v100 = _dafny.Map.Empty.slice().updateUnsafe(_305_v99,_285_v83);
          let _307_v101;
          let _nw58 = Array((new BigNumber(14)).toNumber()).fill(false);
          _307_v101 = _nw58;
          let _308_v102;
          _308_v102 = _dafny.Map.Empty.slice().updateUnsafe(_307_v101,(_277_v77).f23);
          let _309_v103;
          let _310_v104;
          let _311_v105;
          let _312_v106;
          let _out24;
          let _out25;
          let _out26;
          let _out27;
          let _outcollector7 = (_283_v81).m6((new BigNumber((_304_v98).length)).plus(new BigNumber((_306_v100).length)), _dafny.Seq.Concat(_287_v85, _dafny.Seq.UnicodeFromString("yoiqoa")), (_module.D1.create_DC4(new BigNumber((_308_v102).length), (_277_v77).fm15((_283_v81).f21, (_283_v81).f21, new BigNumber(-511), _279_v79, _182_globalState), _307_v101)).dtor_cf5, _182_globalState);
          _out24 = _outcollector7[0];
          _out25 = _outcollector7[1];
          _out26 = _outcollector7[2];
          _out27 = _outcollector7[3];
          _309_v103 = _out24;
          _310_v104 = _out25;
          _311_v105 = _out26;
          _312_v106 = _out27;
          let _313_v107;
          let _nw59 = new _module.C2();
          _nw59.__ctor(_283_v81.f20, false);
          _313_v107 = _nw59;
          let _314_v108;
          let _nw60 = new _module.C2();
          _nw60.__ctor(_179_v5, (_dafny.MultiSet.fromElements(_313_v107, _313_v107)).contains(_313_v107));
          _314_v108 = _nw60;
          _284_v82 = _module.__default.fm0(_285_v83, _183_v8, _302_v96, _182_globalState);
          let _315_v109;
          let _nw61 = new _module.C5();
          _nw61.__ctor(_module.__default.fm5(new BigNumber((_dafny.Seq.of((_283_v81).f21)).length), _182_globalState));
          _315_v109 = _nw61;
        }
        let _316_v110;
        _316_v110 = _dafny.Set.fromElements(_285_v83, _285_v83);
        let _317_v111;
        _317_v111 = _dafny.Seq.of(_316_v110);
        let _318_v112;
        let _319_v113;
        let _320_v114;
        let _321_v115;
        let _out28;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector8 = (_283_v81).m6(new BigNumber((_317_v111).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-938)), function (_322_i10) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), (_283_v81).f21, _182_globalState);
        _out28 = _outcollector8[0];
        _out29 = _outcollector8[1];
        _out30 = _outcollector8[2];
        _out31 = _outcollector8[3];
        _318_v112 = _out28;
        _319_v113 = _out29;
        _320_v114 = _out30;
        _321_v115 = _out31;
      } else {
        let _323_v116;
        let _nw62 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _323_v116 = _nw62;
        _323_v116 = _323_v116;
        (_182_globalState).f17 = !((_277_v77).f23);
        let _324_v117;
        _324_v117 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_279_v79);
        _324_v117 = (_324_v117).update((_176_v2).plus(_176_v2), (_279_v79).Merge(_279_v79));
        _174_v0 = (_277_v77).f23;
        let _index34 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_177_v3).length));
        (_177_v3)[_index34] = _176_v2;
        let _index35 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_177_v3).length));
        (_177_v3)[_index35] = _176_v2;
      }
      let _325_v118;
      let _326_v119;
      let _327_v120;
      let _out32;
      let _out33;
      let _out34;
      let _outcollector9 = _module.__default.m9(_174_v0, _174_v0, _182_globalState);
      _out32 = _outcollector9[0];
      _out33 = _outcollector9[1];
      _out34 = _outcollector9[2];
      _325_v118 = _out32;
      _326_v119 = _out33;
      _327_v120 = _out34;
      let _328_v121;
      _328_v121 = _dafny.Map.Empty.slice().updateUnsafe(true,_325_v118);
      let _329_v122;
      _329_v122 = _module.D13.create_DC36(_dafny.Seq.update(_dafny.Seq.of(_328_v121, _328_v121, _328_v121, _328_v121, _328_v121), _module.__default.safeIndex(_176_v2, new BigNumber((_dafny.Seq.of(_328_v121, _328_v121, _328_v121, _328_v121, _328_v121)).length)), _328_v121));
      let _pat_let_tv3 = _183_v8;
      let _pat_let_tv4 = _183_v8;
      let _pat_let_tv5 = _277_v77;
      let _pat_let_tv6 = _327_v120;
      let _pat_let_tv7 = _327_v120;
      let _pat_let_tv8 = _183_v8;
      let _pat_let_tv9 = _183_v8;
      _183_v8 = function (_source4) {
        if (_source4.is_DC37) {
          let _330___mcc_h4 = (_source4).cf64;
          let _331___mcc_h5 = (_source4).cf65;
          let _332_cf65 = _331___mcc_h5;
          let _333_cf64 = _330___mcc_h4;
          return _pat_let_tv3;
        } else if (_source4.is_DC38) {
          let _334___mcc_h6 = (_source4).cf66;
          let _335___mcc_h7 = (_source4).cf67;
          let _336___mcc_h8 = (_source4).cf68;
          let _337___mcc_h9 = (_source4).cf69;
          let _338_cf69 = _337___mcc_h9;
          let _339_cf68 = _336___mcc_h8;
          let _340_cf67 = _335___mcc_h7;
          let _341_cf66 = _334___mcc_h6;
          return _pat_let_tv4;
        } else if (_source4.is_DC36) {
          let _342___mcc_h10 = (_source4).cf63;
          let _343_cf63 = _342___mcc_h10;
          return (_module.D4.create_DC14((_pat_let_tv5).f23, _pat_let_tv6, _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("mdb")).length), _pat_let_tv7), _pat_let_tv8)).dtor_cf26;
        } else {
          let _344___mcc_h11 = (_source4).cf70;
          let _345_cf70 = _344___mcc_h11;
          return _pat_let_tv9;
        }
      }(_329_v122);
      let _hi2 = _176_v2;
      for (let _346_i11 = _327_v120; _346_i11.isLessThan(_hi2); _346_i11 = _346_i11.plus(_dafny.ONE)) {
        let _347_v123;
        let _nw63 = Array((new BigNumber(16)).toNumber());
        _347_v123 = _nw63;
        let _348_v124;
        let _nw64 = new _module.C2();
        _nw64.__ctor(_179_v5, true);
        _348_v124 = _nw64;
        let _349_v125;
        _349_v125 = _dafny.Map.Empty.slice().updateUnsafe(_346_i11,_348_v124);
        let _index36 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_347_v123).length));
        (_347_v123)[_index36] = (((_349_v125).contains(_176_v2)) ? ((_349_v125).get(_176_v2)) : (_348_v124));
        let _350_v126;
        _350_v126 = _dafny.Seq.of(_348_v124, _348_v124, _348_v124, _348_v124);
        let _index37 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_347_v123).length));
        (_347_v123)[_index37] = (_350_v126)[_module.__default.safeIndex(_346_i11, new BigNumber((_350_v126).length))];
        let _351_v127;
        _351_v127 = _dafny.Seq.of(_329_v122);
        let _352_v128;
        _352_v128 = _dafny.Seq.of(_351_v127);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(112)), ((_353_v122) => function (_354_i12) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(856)), ((_355_v122) => function (_356_i13) {
            return _355_v122;
          })(_353_v122));
        })(_329_v122)), _352_v128), _352_v128)) {
          _174_v0 = (true) || (_325_v118);
          let _357_v129;
          _357_v129 = _dafny.Set.fromElements(_176_v2);
          _357_v129 = (_dafny.Set.fromElements(new BigNumber((_180_v6).length))).Intersect(_357_v129);
          _327_v120 = new BigNumber(371);
          let _358_v130;
          let _359_v131;
          let _360_v132;
          let _361_v133;
          let _out35;
          let _out36;
          let _out37;
          let _out38;
          let _outcollector10 = (_348_v124).m6(_module.__default.safeDivisionInt(new BigNumber((_180_v6).length), new BigNumber((_279_v79).length)), _dafny.Seq.Concat(_180_v6, _dafny.Seq.UnicodeFromString("cqmm")), (_277_v77).f23, _182_globalState);
          _out35 = _outcollector10[0];
          _out36 = _outcollector10[1];
          _out37 = _outcollector10[2];
          _out38 = _outcollector10[3];
          _358_v130 = _out35;
          _359_v131 = _out36;
          _360_v132 = _out37;
          _361_v133 = _out38;
          let _362_v134;
          _362_v134 = _module.D9.create_DC24(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-737)), ((_363_v120) => function (_364_i14) {
  return _363_v120;
})(_327_v120)));
          _253_v54 = (_362_v134).dtor_cf42;
        } else {
          let _365_v135;
          let _366_v136;
          let _367_v137;
          let _out39;
          let _out40;
          let _out41;
          let _outcollector11 = _module.__default.m9((_348_v124).f21, (_277_v77).f23, _182_globalState);
          _out39 = _outcollector11[0];
          _out40 = _outcollector11[1];
          _out41 = _outcollector11[2];
          _365_v135 = _out39;
          _366_v136 = _out40;
          _367_v137 = _out41;
          _180_v6 = _dafny.Seq.update(_180_v6, _module.__default.safeIndex((_dafny.ZERO).minus(_367_v137), new BigNumber((_180_v6).length)), _183_v8);
          let _368_v138;
          _368_v138 = _dafny.Set.fromElements(_180_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-434)), ((_369_v8) => function (_370_i15) {
            return _369_v8;
          })(_183_v8)));
          _327_v120 = (((_277_v77).f23) ? (_346_i11) : (_module.__default.safeModuloInt(new BigNumber((_368_v138).length), _176_v2)));
          (_182_globalState).f15 = _dafny.Seq.IsPrefixOf(((_365_v135) ? (_180_v6) : (_180_v6)), _dafny.Seq.Concat(_180_v6, _180_v6));
          let _371_v139;
          let _nw65 = new _module.C0();
          _nw65.__ctor(true);
          _371_v139 = _nw65;
        }
        let _372_v140;
        let _nw66 = new _module.C4();
        _nw66.__ctor(_325_v118);
        _372_v140 = _nw66;
        let _373_v141;
        _373_v141 = _dafny.Set.fromElements(_372_v140, _372_v140, _372_v140);
        let _374_v142;
        _374_v142 = _module.D12.create_DC31(_373_v141);
        let _pat_let_tv10 = _373_v141;
        let _source5 = function (_pat_let5_0) {
          return function (_375_dt__update__tmp_h1) {
            return function (_pat_let6_0) {
              return function (_376_dt__update_hcf55_h0) {
                return _module.D12.create_DC31(_376_dt__update_hcf55_h0);
              }(_pat_let6_0);
            }(_pat_let_tv10);
          }(_pat_let5_0);
        }(_374_v142);
        if (_source5.is_DC32) {
          let _377___mcc_h12 = (_source5).cf56;
          let _378_cf56 = _377___mcc_h12;
          let _379_v143;
          _379_v143 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_348_v124).f21,_325_v118), _328_v121, _328_v121);
          let _380_v144;
          _380_v144 = _module.D13.create_DC36(_379_v143);
          let _381_v145;
          _381_v145 = _module.D13.create_DC39(_380_v144);
          _381_v145 = _381_v145;
          let _382_v146;
          _382_v146 = _module.D13.create_DC38(_183_v8, (_277_v77).f23, _346_i11, _177_v3);
          let _383_v147;
          let _384_v148;
          let _385_v149;
          let _out42;
          let _out43;
          let _out44;
          let _outcollector12 = _module.__default.m9((((_277_v77).f23) ? (_174_v0) : ((_348_v124).f21)), (_382_v146).dtor_cf67, _182_globalState);
          _out42 = _outcollector12[0];
          _out43 = _outcollector12[1];
          _out44 = _outcollector12[2];
          _383_v147 = _out42;
          _384_v148 = _out43;
          _385_v149 = _out44;
          let _386_v150;
          _386_v150 = _dafny.Set.fromElements(_180_v6);
          _325_v118 = (_386_v150).IsSubsetOf(_386_v150);
          let _387_v151;
          let _nw67 = new _module.C1();
          _nw67.__ctor(_177_v3);
          _387_v151 = _nw67;
        } else if (_source5.is_DC33) {
          let _388___mcc_h13 = (_source5).cf57;
          let _389___mcc_h14 = (_source5).cf58;
          let _390___mcc_h15 = (_source5).cf59;
          let _391___mcc_h16 = (_source5).cf60;
          let _392___mcc_h17 = (_source5).cf61;
          let _393_cf61 = _392___mcc_h17;
          let _394_cf60 = _391___mcc_h16;
          let _395_cf59 = _390___mcc_h15;
          let _396_cf58 = _389___mcc_h14;
          let _397_cf57 = _388___mcc_h13;
          let _398_v152;
          _398_v152 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(727)), ((_399_v8) => function (_400_i16) {
            return _399_v8;
          })(_183_v8)),false);
          _397_cf57 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.fm3(new BigNumber(-604), _327_v120, !((((_398_v152).contains(_180_v6)) ? ((_398_v152).get(_180_v6)) : (false))), _182_globalState)), _176_v2);
          let _401_v153;
          let _nw68 = new _module.C4();
          _nw68.__ctor((_348_v124).f21);
          _401_v153 = _nw68;
          _372_v140 = _372_v140;
          (_182_globalState).f11 = (_277_v77).f23;
        } else if (_source5.is_DC34) {
          let _402_v154;
          _402_v154 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_174_v0);
          _402_v154 = (_402_v154).update(new BigNumber(918), (new BigNumber((_dafny.MultiSet.FromArray(_175_v1)).cardinality())).isLessThan(_346_i11));
          let _403_v155;
          _403_v155 = _module.D13.create_DC37(new BigNumber(-142), _dafny.Seq.of(_253_v54));
          let _404_v156;
          _404_v156 = _dafny.Set.fromElements(_176_v2, _176_v2, new BigNumber(368));
          let _405_v157;
          _405_v157 = _module.D13.create_DC38(_183_v8, _326_v119, _346_i11, _177_v3);
          let _406_v158;
          let _nw69 = Array((new BigNumber(26)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _183_v8;
          _nw69[(_dafny.ONE).toNumber()] = ((!(true)) ? ((_180_v6)[_module.__default.safeIndex(new BigNumber((_404_v156).length), new BigNumber((_180_v6).length))]) : (_183_v8));
          _nw69[(new BigNumber(2)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(3)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(4)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(5)).toNumber()] = (_405_v157).dtor_cf66;
          _nw69[(new BigNumber(6)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(7)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(8)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(9)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(10)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(11)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(12)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(13)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(14)).toNumber()] = (_180_v6)[_module.__default.safeIndex(_346_i11, new BigNumber((_180_v6).length))];
          _nw69[(new BigNumber(15)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(16)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(17)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(18)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(19)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(20)).toNumber()] = _module.__default.fm12(_327_v120, _180_v6, _180_v6, _182_globalState);
          _nw69[(new BigNumber(21)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(22)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(23)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(24)).toNumber()] = _183_v8;
          _nw69[(new BigNumber(25)).toNumber()] = _module.__default.fm12(new BigNumber((_175_v1).length), _180_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), ((_407_v8) => function (_408_i17) {
            return _407_v8;
          })(_183_v8)), _182_globalState);
          _406_v158 = _nw69;
          let _index38 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_406_v158).length));
          (_406_v158)[_index38] = _183_v8;
          let _index39 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_406_v158).length));
          let _rhs33 = _403_v155;
          let _rhs34 = _404_v156;
          let _rhs35 = _327_v120;
          let _rhs36 = _176_v2;
          let _rhs37 = _183_v8;
          let _lhs19 = _406_v158;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_406_v158).length));
          _403_v155 = _rhs33;
          _404_v156 = _rhs34;
          _176_v2 = _rhs35;
          _327_v120 = _rhs36;
          _lhs19[_lhs20] = _rhs37;
          let _index40 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_177_v3).length));
          (_177_v3)[_index40] = _176_v2;
          let _409_v159;
          _409_v159 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_346_i11, new BigNumber((_402_v154).length)));
          let _index41 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_177_v3).length));
          (_177_v3)[_index41] = new BigNumber((_409_v159).cardinality());
          let _410_v160;
          _410_v160 = _dafny.Map.Empty.slice().updateUnsafe(true,_175_v1);
          let _411_v161;
          _411_v161 = _dafny.Seq.of((((_410_v160).contains(true)) ? ((_410_v160).get(true)) : (_175_v1)));
          let _index42 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_177_v3).length));
          (_177_v3)[_index42] = (_176_v2).plus(new BigNumber(((_411_v161)[_module.__default.safeIndex(_module.__default.fm5(_327_v120, _182_globalState), new BigNumber((_411_v161).length))]).length));
        } else if (_source5.is_DC31) {
          let _412___mcc_h18 = (_source5).cf55;
          let _413_cf55 = _412___mcc_h18;
          let _414_v162;
          let _415_v163;
          let _out45;
          let _out46;
          let _outcollector13 = (_372_v140).m1((_348_v124).f21, new BigNumber((_180_v6).length), !(_325_v118), _module.__default.safeModuloInt((((_279_v79).contains(_325_v118)) ? ((_279_v79).get(_325_v118)) : (new BigNumber(-122))), new BigNumber(-560)), _182_globalState);
          _out45 = _outcollector13[0];
          _out46 = _outcollector13[1];
          _414_v162 = _out45;
          _415_v163 = _out46;
          let _rhs38 = (_346_i11).isLessThanOrEqualTo((_346_i11).minus(new BigNumber((_module.__default.fm19(_253_v54, _414_v162, _182_globalState)).length)));
          let _rhs39 = (((_348_v124).f21) ? (new BigNumber(-772)) : ((_dafny.ZERO).minus(_176_v2)));
          let _rhs40 = _326_v119;
          let _rhs41 = (_348_v124).f21;
          let _lhs21 = _182_globalState;
          let _lhs22 = _182_globalState;
          _415_v163 = _rhs38;
          _176_v2 = _rhs39;
          _lhs21.f15 = _rhs40;
          _lhs22.f15 = _rhs41;
          let _416_v164;
          _416_v164 = _module.D3.create_DC11(_279_v79, new BigNumber((_251_v52).cardinality()), _177_v3);
          let _rhs42 = ((_416_v164).dtor_cf18).Merge(_279_v79);
          let _rhs43 = false;
          _279_v79 = _rhs42;
          _415_v163 = _rhs43;
          _180_v6 = _dafny.Seq.Concat(((_415_v163) ? (_180_v6) : (_dafny.Seq.UnicodeFromString("esrnngllh"))), _180_v6);
        } else {
          let _417___mcc_h19 = (_source5).cf62;
          let _418_cf62 = _417___mcc_h19;
          _176_v2 = _327_v120;
          _249_v50 = _249_v50;
          (_182_globalState).f1 = _dafny.Seq.of((_348_v124).f21);
          let _419_v165;
          _419_v165 = _dafny.Set.fromElements(_177_v3);
          let _420_v166;
          _420_v166 = _module.D15.create_DC42(_372_v140);
          let _421_v167;
          _421_v167 = _dafny.Seq.of(_372_v140, _372_v140);
          let _422_v168;
          let _nw70 = Array((new BigNumber(27)).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _372_v140;
          _nw70[(_dafny.ONE).toNumber()] = (_module.D15.create_DC42(_372_v140)).dtor_cf75;
          _nw70[(new BigNumber(2)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(3)).toNumber()] = (_420_v166).dtor_cf75;
          _nw70[(new BigNumber(4)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(5)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(6)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(7)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(8)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(9)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(10)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(11)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(12)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(13)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(14)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(15)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(16)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(17)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(18)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(19)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(20)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(21)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(22)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(23)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(24)).toNumber()] = (_421_v167)[_module.__default.safeIndex(new BigNumber((_module.__default.fm19(_253_v54, (_277_v77).f23, _182_globalState)).length), new BigNumber((_421_v167).length))];
          _nw70[(new BigNumber(25)).toNumber()] = _372_v140;
          _nw70[(new BigNumber(26)).toNumber()] = _372_v140;
          _422_v168 = _nw70;
          let _index43 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_422_v168).length));
          (_422_v168)[_index43] = _372_v140;
          let _423_v169;
          _423_v169 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("tetmmpkc"), _180_v6, _180_v6);
          let _index44 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_422_v168).length));
          let _rhs44 = _419_v165;
          let _rhs45 = _372_v140;
          let _rhs46 = (_423_v169)[_module.__default.safeIndex(_327_v120, new BigNumber((_423_v169).length))];
          let _lhs23 = _422_v168;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(995), new BigNumber((_422_v168).length));
          _419_v165 = _rhs44;
          _lhs23[_lhs24] = _rhs45;
          _180_v6 = _rhs46;
        }
        _176_v2 = _327_v120;
      }
      let _424_v170;
      let _init11 = ((_425_v1) => function (_426_i19) {
        return _dafny.Seq.IsPrefixOf(_425_v1, _425_v1);
      })(_175_v1);
      let _nw71 = Array((new BigNumber(28)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw71.length); _i0_11++) {
        _nw71[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _424_v170 = _nw71;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_424_v170).length))) {
        let _427_i18 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_427_i18)) && ((_427_i18).isLessThan(new BigNumber((_424_v170).length))))) {
          (_424_v170)[(_427_i18)] = false;
        }
      }
      let _428_v171;
      _428_v171 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_174_v0);
      let _hi3 = ((((_279_v79).contains((_277_v77).f23)) ? ((_279_v79).get((_277_v77).f23)) : (_176_v2))).minus(_327_v120);
      for (let _429_i20 = ((_277_v77).fm14(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_428_v171,_327_v120)).length), _182_globalState)).multipliedBy((_dafny.ZERO).minus(_327_v120)); _429_i20.isLessThan(_hi3); _429_i20 = _429_i20.plus(_dafny.ONE)) {
        _327_v120 = (_dafny.ZERO).minus(_327_v120);
        let _430_v172;
        _430_v172 = _module.D10.create_DC27(_251_v52);
        let _source6 = _430_v172;
        if (_source6.is_DC28) {
          let _431___mcc_h20 = (_source6).cf49;
          let _432___mcc_h21 = (_source6).cf50;
          let _433_cf50 = _432___mcc_h21;
          let _434_cf49 = _431___mcc_h20;
          let _435_v173;
          _435_v173 = _dafny.Seq.of(_253_v54);
          let _436_v174;
          _436_v174 = _module.D13.create_DC39(_module.D13.create_DC37(_176_v2, _435_v173));
          let _437_v175;
          _437_v175 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_436_v174);
          _437_v175 = _437_v175;
          let _438_v176;
          _438_v176 = _dafny.Set.fromElements(_177_v3, _177_v3);
          let _439_v177;
          _439_v177 = _dafny.Map.Empty.slice().updateUnsafe(_327_v120,_438_v176);
          let _440_v178;
          let _nw72 = new _module.C2();
          _nw72.__ctor(_179_v5, (_277_v77).f23);
          _440_v178 = _nw72;
          let _441_v179;
          _441_v179 = _dafny.Map.Empty.slice().updateUnsafe(_424_v170,_440_v178);
          _174_v0 = (_277_v77).fm15((_438_v176).IsSubsetOf((((_439_v177).contains(_176_v2)) ? ((_439_v177).get(_176_v2)) : (_438_v176))), _326_v119, _module.__default.safeDivisionInt(new BigNumber((_441_v179).length), _327_v120), _279_v79, _182_globalState);
          _326_v119 = _434_cf49;
          let _442_v180;
          let _nw73 = new _module.C5();
          _nw73.__ctor(_429_i20);
          _442_v180 = _nw73;
        } else {
          let _443___mcc_h22 = (_source6).cf48;
          let _444_cf48 = _443___mcc_h22;
          let _index45 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_424_v170).length));
          (_424_v170)[_index45] = true;
          let _index46 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_424_v170).length));
          (_424_v170)[_index46] = _326_v119;
          _181_v7 = (_181_v7).update(_327_v120, new BigNumber((_180_v6).length));
          _327_v120 = _176_v2;
          let _445_v181;
          let _init12 = ((_446_v0) => function (_447_i21) {
            return _446_v0;
          })(_174_v0);
          let _nw74 = Array((new BigNumber(2)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw74.length); _i0_12++) {
            _nw74[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _445_v181 = _nw74;
          let _448_v182;
          _448_v182 = _module.D1.create_DC4((_dafny.ZERO).minus(_327_v120), (_277_v77).f23, _445_v181);
          let _449_v183;
          _449_v183 = _dafny.Seq.of(_175_v1, _175_v1, _dafny.Seq.of(false, _326_v119));
          let _rhs47 = (_277_v77).f23;
          let _rhs48 = (_327_v120).isEqualTo((_448_v182).dtor_cf4);
          let _rhs49 = _326_v119;
          let _rhs50 = (_module.__default.safeDivisionInt(_176_v2, _327_v120)).minus(_module.__default.safeModuloInt(_327_v120, new BigNumber(((_449_v183)[_module.__default.safeIndex(_429_i20, new BigNumber((_449_v183).length))]).length)));
          let _lhs25 = _182_globalState;
          let _lhs26 = _182_globalState;
          _lhs25.f11 = _rhs47;
          _326_v119 = _rhs48;
          _lhs26.f11 = _rhs49;
          _176_v2 = _rhs50;
        }
        let _450_v184;
        let _nw75 = new _module.C1();
        _nw75.__ctor(_177_v3);
        _450_v184 = _nw75;
        let _index47 = _module.__default.safeIndex(new BigNumber(974), new BigNumber(((_450_v184).f22).length));
        ((_450_v184).f22)[_index47] = _327_v120;
        let _index48 = _module.__default.safeIndex(new BigNumber(974), new BigNumber(((_450_v184).f22).length));
        ((_450_v184).f22)[_index48] = _327_v120;
      }
      let _hi4 = _176_v2;
      for (let _451_i22 = _327_v120; _451_i22.isLessThan(_hi4); _451_i22 = _451_i22.plus(_dafny.ONE)) {
        let _452_v185;
        let _453_v186;
        let _454_v187;
        let _out47;
        let _out48;
        let _out49;
        let _outcollector14 = _module.__default.m9((_277_v77).f23, false, _182_globalState);
        _out47 = _outcollector14[0];
        _out48 = _outcollector14[1];
        _out49 = _outcollector14[2];
        _452_v185 = _out47;
        _453_v186 = _out48;
        _454_v187 = _out49;
        let _455_v188;
        let _nw76 = new _module.C4();
        _nw76.__ctor(((_277_v77).fm15(false, (_277_v77).f23, _454_v187, _279_v79, _182_globalState)) || (_452_v185));
        _455_v188 = _nw76;
        let _456_v189;
        let _nw77 = new _module.C2();
        _nw77.__ctor(_179_v5, (_277_v77).f23);
        _456_v189 = _nw77;
        _181_v7 = (_181_v7).update(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("bmdrnblv")).length), _454_v187), _176_v2);
      }
      let _457_v190;
      let _458_v191;
      let _459_v192;
      let _out50;
      let _out51;
      let _out52;
      let _outcollector15 = _module.__default.m9(_174_v0, true, _182_globalState);
      _out50 = _outcollector15[0];
      _out51 = _outcollector15[1];
      _out52 = _outcollector15[2];
      _457_v190 = _out50;
      _458_v191 = _out51;
      _459_v192 = _out52;
      _328_v121 = (_328_v121).update(_174_v0, !(_326_v119));
      let _460_v193;
      let _nw78 = new _module.C4();
      _nw78.__ctor(_326_v119);
      _460_v193 = _nw78;
      let _461_v194;
      _461_v194 = _dafny.MultiSet.fromElements(_424_v170);
      let _462_v195;
      _462_v195 = _dafny.MultiSet.fromElements(_176_v2);
      let _rhs51 = _dafny.Seq.update(_180_v6, _module.__default.safeIndex((_459_v192).minus(new BigNumber((_dafny.Seq.update(_253_v54, _module.__default.safeIndex(_459_v192, new BigNumber((_253_v54).length)), (((_461_v194).contains(_424_v170)) ? ((_461_v194).get(_424_v170)) : (_176_v2)))).length)), new BigNumber((_180_v6).length)), new _dafny.CodePoint('w'.codePointAt(0)));
      let _rhs52 = new BigNumber(((_462_v195).Union(_462_v195)).cardinality());
      let _rhs53 = _460_v193;
      _180_v6 = _rhs51;
      _459_v192 = _rhs52;
      _460_v193 = _rhs53;
      (_182_globalState).f1 = _dafny.Seq.of(true);
      let _463_v196;
      let _nw79 = new _module.C1();
      _nw79.__ctor(_177_v3);
      _463_v196 = _nw79;
      process.stdout.write(_dafny.toString(_174_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_175_v1, _dafny.Seq.of(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_176_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v5).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_180_v6, _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_181_v7).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(644),new BigNumber(644)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_182_globalState.f1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_182_globalState).f4, _dafny.Seq.of(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_182_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_globalState).f12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(644),new BigNumber(644)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_182_globalState.f16, _dafny.Seq.of(true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_182_globalState.f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_250_v51).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_251_v52).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_252_v53, _dafny.Seq.of(_module.D10.create_DC27(_dafny.MultiSet.fromElements(false, true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_253_v54, _dafny.Seq.of(new BigNumber(-916)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v77).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_278_v78).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v79).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-916)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_v118));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_326_v119));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_327_v120));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_328_v121).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_329_v122).dtor_cf63, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_424_v170)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_428_v171).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_457_v190));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_458_v191));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_459_v192));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_460_v193).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_461_v194).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_462_v195).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
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
    static create_DC2(cf2) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
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
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
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
    static create_DC4(cf4, cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf10, cf11) {
      let $dt = new D1(2);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(3);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + this.cf7.toVerbatimString(true) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && this.cf5 === other.cf5 && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf3 === other.cf3;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.ZERO, false, []);
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
    static create_DC8(cf13, cf14, cf15) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8([], _dafny.ZERO, _dafny.ZERO);
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
    static create_DC10(cf17) {
      let $dt = new D3(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC11(cf18, cf19, cf20) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC9(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO);
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
    static create_DC13(cf22) {
      let $dt = new D4(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC14(cf23, cf24, cf25, cf26) {
      let $dt = new D4(1);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC12(cf21) {
      let $dt = new D4(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC15(cf27) {
      let $dt = new D4(3);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf27) + ")";
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
        return other.$tag === 1 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf21 === other.cf21;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(_dafny.ZERO);
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
    static create_DC17(cf29, cf30, cf31, cf32, cf33) {
      let $dt = new D5(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D5(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC18(cf34) {
      let $dt = new D5(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(false, _module.D2.Default(), false, false, _dafny.ZERO);
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
    static create_DC20(cf36, cf37, cf38) {
      let $dt = new D6(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC19(cf35) {
      let $dt = new D6(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC20(false, _dafny.ZERO, false);
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
    static create_DC21(cf39) {
      let $dt = new D7(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39;
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
    static create_DC23(cf41) {
      let $dt = new D8(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC22(cf40) {
      let $dt = new D8(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf40 === other.cf40;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(_dafny.MultiSet.Empty);
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
    static create_DC25(cf43, cf44, cf45, cf46) {
      let $dt = new D9(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC24(cf42) {
      let $dt = new D9(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC26(cf47) {
      let $dt = new D9(2);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(_dafny.ZERO, false, _dafny.ZERO, null);
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
    static create_DC28(cf49, cf50) {
      let $dt = new D10(0);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC27(cf48) {
      let $dt = new D10(1);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC28(false, _dafny.ZERO);
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
    static create_DC30(cf52, cf53, cf54) {
      let $dt = new D11(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC29(cf51) {
      let $dt = new D11(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53) && this.cf54 === other.cf54;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30(false, _dafny.ZERO, false);
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
    static create_DC32(cf56) {
      let $dt = new D12(0);
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC33(cf57, cf58, cf59, cf60, cf61) {
      let $dt = new D12(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC34() {
      let $dt = new D12(2);
      return $dt;
    }
    static create_DC31(cf55) {
      let $dt = new D12(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC35(cf62) {
      let $dt = new D12(4);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get is_DC35() { return this.$tag === 4; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34";
      } else if (this.$tag === 3) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 4) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && this.cf61 === other.cf61;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(_dafny.ZERO);
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
    static create_DC37(cf64, cf65) {
      let $dt = new D13(0);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC38(cf66, cf67, cf68, cf69) {
      let $dt = new D13(1);
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC36(cf63) {
      let $dt = new D13(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC39(cf70) {
      let $dt = new D13(3);
      $dt.cf70 = cf70;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get is_DC39() { return this.$tag === 3; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf70() { return this.cf70; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf70) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf66, other.cf66) && this.cf67 === other.cf67 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf70, other.cf70);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC37(_dafny.ZERO, _dafny.Seq.of());
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
    static create_DC41(cf72, cf73, cf74) {
      let $dt = new D14(0);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC40(cf71) {
      let $dt = new D14(1);
      $dt.cf71 = cf71;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf71() { return this.cf71; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf72) + ", " + this.cf73.toVerbatimString(true) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf71) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf72 === other.cf72 && _dafny.areEqual(this.cf73, other.cf73) && this.cf74 === other.cf74;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf71 === other.cf71;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC41(null, _dafny.Seq.UnicodeFromString(""), []);
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
    static create_DC43(cf76, cf77) {
      let $dt = new D15(0);
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC44(cf78, cf79) {
      let $dt = new D15(1);
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC42(cf75) {
      let $dt = new D15(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC45(cf80) {
      let $dt = new D15(3);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get is_DC45() { return this.$tag === 3; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC45" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76) && this.cf77 === other.cf77;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf78 === other.cf78 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf75 === other.cf75;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC43(_dafny.Set.Empty, false);
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
    static create_DC46(cf81) {
      let $dt = new D16(0);
      $dt.cf81 = cf81;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get dtor_cf81() { return this.cf81; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf81) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81);
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
          return D16.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = _dafny.Seq.of();
      this.f2 = [];
      this.f11 = false;
      this.f15 = false;
      this.f16 = _dafny.Seq.of();
      this.f17 = false;
      this._f0 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f3 = false;
      this._f4 = _dafny.Seq.of();
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = false;
      this._f10 = false;
      this._f12 = _dafny.Map.Empty;
      this._f13 = false;
      this._f14 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this).f17 = f17;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f23 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f23) {
      let _this = this;
      (_this)._f23 = f23;
      return;
    }
    fm14(p0, globalState) {
      let _this = this;
      if ((((_this).f23) ? ((_this).f23) : (false))) {
        return (new BigNumber((_dafny.Seq.of((_this).f23)).length)).minus(new BigNumber(493));
      } else {
        return (new BigNumber(969)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("opfaxvbk"))).length));
      }
    };
    fm15(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f23) || ((_module.D0.create_DC0((_this).f23)).dtor_cf0);
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f22 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f22) {
      let _this = this;
      (_this)._f22 = f22;
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let _464_v0;
      _464_v0 = _dafny.Seq.of(true);
      let _465_v1;
      _465_v1 = true;
      let _466_v2;
      _466_v2 = _dafny.Seq.UnicodeFromString("nlrnbcijs");
      let _hi5 = p1;
      for (let _467_i0 = (_module.__default.fm9(p0, new BigNumber((_464_v0).length), _465_v1, globalState)).plus((_module.D1.create_DC5(_466_v2, p0, new BigNumber((_466_v2).length))).dtor_cf9); _467_i0.isLessThan(_hi5); _467_i0 = _467_i0.plus(_dafny.ONE)) {
        let _468_v3;
        _468_v3 = new BigNumber(596);
        _468_v3 = _module.__default.safeModuloInt(p1, _468_v3);
        let _469_v5;
        _469_v5 = new _dafny.CodePoint('j'.codePointAt(0));
        let _470_v6;
        _470_v6 = _dafny.Set.fromElements(_469_v5);
        let _471_v7;
        _471_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_465_v1, _465_v1, _465_v1),(_this).f22);
        let _472_v8;
        _472_v8 = _dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_470_v6).Elements) {
            let _473_v4 = _compr_21;
            if ((_470_v6).contains(_473_v4)) {
              _coll21.push([_473_v4,_467_i0]);
            }
          }
          return _coll21;
        }(),new BigNumber((_471_v7).length));
        let _474_v9;
        _474_v9 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber(706));
        _472_v8 = (_472_v8).update(_474_v9, p1);
        let _475_v10;
        let _nw80 = Array((new BigNumber(2)).toNumber()).fill(false);
        _475_v10 = _nw80;
        let _476_v11;
        _476_v11 = _module.D1.create_DC4(p1, _465_v1, _475_v10);
        let _source7 = _476_v11;
        if (_source7.is_DC4) {
          let _477___mcc_h0 = (_source7).cf4;
          let _478___mcc_h1 = (_source7).cf5;
          let _479___mcc_h2 = (_source7).cf6;
          let _480_cf6 = _479___mcc_h2;
          let _481_cf5 = _478___mcc_h1;
          let _482_cf4 = _477___mcc_h0;
          let _483_v12;
          _483_v12 = _dafny.Set.fromElements(p1, p1);
          let _index49 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length));
          (_475_v10)[_index49] = (_dafny.Set.fromElements(p0, _468_v3, _468_v3)).IsSubsetOf(_483_v12);
          let _index50 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length));
          (_475_v10)[_index50] = !(_465_v1);
          let _484_v13;
          _484_v13 = _dafny.MultiSet.fromElements(p1);
          let _485_v14;
          _485_v14 = _module.D2.create_DC7(_484_v13);
          _465_v1 = (new BigNumber(((_485_v14).dtor_cf12).cardinality())).isEqualTo((p1).plus(p1));
          let _486_v15;
          _486_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_466_v2).length));
          let _index51 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length));
          (_475_v10)[_index51] = (_module.__default.safeModuloInt(_482_cf4, new BigNumber((_486_v15).length))).isEqualTo((_468_v3).plus(p1));
          let _487_v16;
          _487_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(624));
          let _488_v17;
          _488_v17 = _module.D0.create_DC1((_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))]);
          let _489_v18;
          _489_v18 = _module.D0.create_DC2(_488_v17);
          let _pat_let_tv11 = _465_v1;
          let _pat_let_tv12 = _475_v10;
          let _pat_let_tv13 = _475_v10;
          let _pat_let_tv14 = globalState;
          let _490_v19;
          _490_v19 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let7_0) {
            return function (_491_dt__update__tmp_h0) {
              return function (_pat_let8_0) {
                return function (_492_dt__update_hcf2_h0) {
                  return _module.D0.create_DC2(_492_dt__update_hcf2_h0);
                }(_pat_let8_0);
              }(_module.__default.fm11(_467_i0, _pat_let_tv11, (_pat_let_tv13)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_pat_let_tv12).length))], _467_i0, _pat_let_tv14));
            }(_pat_let7_0);
          }(_489_v18),_489_v18);
          let _nw81 = Array((new BigNumber(15)).toNumber());
          _nw81[(_dafny.ZERO).toNumber()] = false;
          _nw81[(_dafny.ONE).toNumber()] = (_467_i0).isLessThan(_468_v3);
          _nw81[(new BigNumber(2)).toNumber()] = (_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))];
          _nw81[(new BigNumber(3)).toNumber()] = (!(_465_v1)) === ((_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))]);
          _nw81[(new BigNumber(4)).toNumber()] = (p0).isLessThan(new BigNumber(-192));
          _nw81[(new BigNumber(5)).toNumber()] = ((_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))]) === (_481_cf5);
          _nw81[(new BigNumber(6)).toNumber()] = (_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))];
          _nw81[(new BigNumber(7)).toNumber()] = !(_487_v16).equals(_module.__default.fm10(globalState));
          _nw81[(new BigNumber(8)).toNumber()] = true;
          _nw81[(new BigNumber(9)).toNumber()] = _465_v1;
          _nw81[(new BigNumber(10)).toNumber()] = (_module.__default.fm9(_482_cf4, _467_i0, _481_cf5, globalState)).isLessThan(new BigNumber((_490_v19).length));
          _nw81[(new BigNumber(11)).toNumber()] = _481_cf5;
          _nw81[(new BigNumber(12)).toNumber()] = (_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))];
          _nw81[(new BigNumber(13)).toNumber()] = ((_475_v10)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_475_v10).length))]) || (_465_v1);
          _nw81[(new BigNumber(14)).toNumber()] = ((_465_v1) ? (_481_cf5) : (_481_cf5));
          _480_cf6 = _nw81;
        } else if (_source7.is_DC5) {
          let _493___mcc_h3 = (_source7).cf7;
          let _494___mcc_h4 = (_source7).cf8;
          let _495___mcc_h5 = (_source7).cf9;
          let _496_cf9 = _495___mcc_h5;
          let _497_cf8 = _494___mcc_h4;
          let _498_cf7 = _493___mcc_h3;
          _469_v5 = _469_v5;
          let _499_v20;
          _499_v20 = _dafny.Map.Empty.slice().updateUnsafe(_465_v1,_464_v0);
          _499_v20 = (_499_v20).update(_module.__default.fm0(_465_v1, _469_v5, _465_v1, globalState), _464_v0);
          _498_cf7 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ss"), _module.__default.safeIndex(_467_i0, new BigNumber((_dafny.Seq.UnicodeFromString("ss")).length)), _469_v5);
          let _index52 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_475_v10).length));
          (_475_v10)[_index52] = _465_v1;
          let _500_v21;
          _500_v21 = _dafny.Set.fromElements(p0);
          let _501_v22;
          _501_v22 = _module.D0.create_DC0(_465_v1);
          let _index53 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_475_v10).length));
          (_475_v10)[_index53] = (((_500_v21).IsProperSubsetOf(_500_v21)) ? ((_501_v22).dtor_cf0) : (_465_v1));
        } else if (_source7.is_DC6) {
          let _502___mcc_h6 = (_source7).cf10;
          let _503___mcc_h7 = (_source7).cf11;
          let _504_cf11 = _503___mcc_h7;
          let _505_cf10 = _502___mcc_h6;
          let _506_v23;
          _506_v23 = _dafny.Set.fromElements(_465_v1);
          let _507_v24;
          _507_v24 = _dafny.Map.Empty.slice().updateUnsafe(_469_v5,(_506_v23).IsSubsetOf(_506_v23));
          let _508_v25;
          _508_v25 = _dafny.MultiSet.fromElements(new BigNumber(69), _504_cf11);
          _507_v24 = (_507_v24).update(_469_v5, (_dafny.MultiSet.fromElements(_504_cf11)).IsSubsetOf(_508_v25));
          (globalState).f11 = _465_v1;
          let _509_v26;
          _509_v26 = _dafny.Seq.of(_466_v2);
          let _510_v27;
          _510_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_467_i0);
          _469_v5 = _module.__default.fm12((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, p1))), _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_466_v2, (_509_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(_468_v3), new BigNumber((_509_v26).length))]), _module.__default.safeIndex((((_510_v27).contains((_this).f22)) ? ((_510_v27).get((_this).f22)) : (_504_cf11)), new BigNumber((_dafny.Seq.Concat(_466_v2, (_509_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(_468_v3), new BigNumber((_509_v26).length))])).length)), _469_v5), _module.__default.safeIndex((_dafny.ZERO).minus(_467_i0), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_466_v2, (_509_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(_468_v3), new BigNumber((_509_v26).length))]), _module.__default.safeIndex((((_510_v27).contains((_this).f22)) ? ((_510_v27).get((_this).f22)) : (_504_cf11)), new BigNumber((_dafny.Seq.Concat(_466_v2, (_509_v26)[_module.__default.safeIndex((_dafny.ZERO).minus(_468_v3), new BigNumber((_509_v26).length))])).length)), _469_v5)).length)), _469_v5), (_509_v26)[_module.__default.safeIndex(p0, new BigNumber((_509_v26).length))], globalState);
          (globalState).f11 = _465_v1;
        } else {
          let _511___mcc_h8 = (_source7).cf3;
          let _512_cf3 = _511___mcc_h8;
          let _513_v28;
          let _nw82 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _513_v28 = _nw82;
          let _514_v29;
          _514_v29 = _dafny.Seq.of(_467_i0);
          let _index54 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length));
          (_513_v28)[_index54] = _dafny.Seq.Concat(_dafny.Seq.of(_468_v3, p1), _514_v29);
          let _515_v30;
          let _init13 = ((_516_v1) => function (_517_i1) {
            return _dafny.Set.fromElements(_516_v1, _516_v1);
          })(_465_v1);
          let _nw83 = Array((_dafny.ONE).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw83.length); _i0_13++) {
            _nw83[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _515_v30 = _nw83;
          let _index55 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length));
          let _rhs54 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(266)), ((_518_v5) => function (_519_i2) {
            return _518_v5;
          })(_469_v5));
          let _rhs55 = _514_v29;
          let _rhs56 = _515_v30;
          let _rhs57 = (new BigNumber((_466_v2).length)).plus(((_465_v1) ? (_468_v3) : (new BigNumber(578))));
          let _lhs27 = _513_v28;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length));
          _466_v2 = _rhs54;
          _lhs27[_lhs28] = _rhs55;
          _515_v30 = _rhs56;
          _468_v3 = _rhs57;
          let _520_v31;
          _520_v31 = _dafny.Seq.of((_513_v28)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length))], (_513_v28)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length))]);
          let _521_v32;
          _521_v32 = _dafny.Map.Empty.slice().updateUnsafe(_520_v31,_module.__default.safeDivisionInt(p1, _module.__default.fm9(p0, new BigNumber((_466_v2).length), _465_v1, globalState)));
          let _522_v33;
          _522_v33 = _dafny.Set.fromElements(new BigNumber(464));
          _521_v32 = (_521_v32).update(_dafny.Seq.update(_dafny.Seq.of((_513_v28)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length))], _514_v29), _module.__default.safeIndex(new BigNumber((_522_v33).length), new BigNumber((_dafny.Seq.of((_513_v28)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length))], _514_v29)).length)), (_513_v28)[_module.__default.safeIndex(new BigNumber(882), new BigNumber((_513_v28).length))]), p0);
          _468_v3 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm13(_467_i0, _467_i0, globalState), _520_v31)).length);
          let _523_v34;
          _523_v34 = _dafny.Map.Empty.slice().updateUnsafe(_468_v3,_465_v1);
          let _524_v35;
          _524_v35 = _dafny.Set.fromElements(_523_v34, _523_v34, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-625),_465_v1));
          let _525_v37;
          _525_v37 = _module.D3.create_DC9(_dafny.Set.fromElements(function () {
  let _coll22 = new _dafny.Map();
  for (const _compr_22 of (_514_v29).Elements) {
    let _526_v36 = _compr_22;
    if (_dafny.Seq.contains(_514_v29, _526_v36)) {
      _coll22.push([(_526_v36).multipliedBy(p0),_465_v1]);
    }
  }
  return _coll22;
}(), _523_v34, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_465_v1)).cardinality()),_465_v1)));
          (globalState).f15 = ((_524_v35).Intersect(_524_v35)).IsSubsetOf((_525_v37).dtor_cf16);
        }
        _468_v3 = p1;
      }
      _465_v1 = true;
      let _527_v38;
      _527_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,_465_v1);
      let _528_v39;
      _528_v39 = _dafny.MultiSet.fromElements(_465_v1);
      let _529_v40;
      _529_v40 = _dafny.Seq.of(_528_v39);
      _527_v38 = (_527_v38).update(p1, (new BigNumber(((_529_v40)[_module.__default.safeIndex(p1, new BigNumber((_529_v40).length))]).cardinality())).isLessThan(p1));
      let _530_v41;
      let _init14 = ((_531_p0) => function (_532_i3) {
        return (new BigNumber(-797)).isLessThanOrEqualTo(_531_p0);
      })(p0);
      let _nw84 = Array((new BigNumber(4)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw84.length); _i0_14++) {
        _nw84[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _530_v41 = _nw84;
      let _index56 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_530_v41).length));
      (_530_v41)[_index56] = !(_465_v1) || (_465_v1);
      let _533_v42;
      _533_v42 = _dafny.MultiSet.fromElements(p1);
      let _index57 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_530_v41).length));
      (_530_v41)[_index57] = !(_533_v42).contains(_module.__default.safeDivisionInt(new BigNumber(-892), p1));
      (globalState).f15 = false;
      let _534_i4;
      _534_i4 = _dafny.ZERO;
      L0: {
        while ((_530_v41)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_530_v41).length))]) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_534_i4)) {
              break L0;
            }
            _534_i4 = (_534_i4).plus(_dafny.ONE);
            let _535_v43;
            let _nw85 = new _module.C0();
            _nw85.__ctor(_465_v1);
            _535_v43 = _nw85;
            let _536_v44;
            _536_v44 = new BigNumber(-941);
            let _537_v45;
            _537_v45 = _dafny.Seq.of(p0);
            _536_v44 = new BigNumber((_537_v45).length);
            _536_v44 = p0;
            _465_v1 = _465_v1;
          }
        }
      }
      return;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Map.Empty;
      if (p1) {
        let _538_v0;
        _538_v0 = _dafny.Seq.of(p0, p0, p0);
        let _539_v1;
        _539_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(798));
        let _540_v2;
        _540_v2 = _dafny.Set.fromElements(p0, (((_539_v1).contains(false)) ? ((_539_v1).get(false)) : (p0)));
        let _541_v3;
        _541_v3 = _dafny.Seq.of(p1);
        let _542_v4;
        let _nw86 = new _module.C0();
        _nw86.__ctor(p1);
        _542_v4 = _nw86;
        let _543_v5;
        _543_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,_542_v4);
        let _544_v6;
        _544_v6 = new _dafny.CodePoint('m'.codePointAt(0));
        let _545_v7;
        let _init15 = ((_546_v3) => function (_547_i0) {
          return _546_v3;
        })(_541_v3);
        let _nw87 = Array((new BigNumber(4)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw87.length); _i0_15++) {
          _nw87[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _545_v7 = _nw87;
        let _548_v8;
        _548_v8 = _module.D1.create_DC3(_545_v7);
        let _549_v9;
        _549_v9 = _dafny.Map.Empty.slice().updateUnsafe(_548_v8,p1);
        let _550_v10;
        _550_v10 = _dafny.Map.Empty.slice().updateUnsafe((_543_v5).update((_542_v4).f23, _542_v4),_module.__default.fm0((_542_v4).f23, _544_v6, (((_549_v9).contains(_module.D1.create_DC3(_545_v7))) ? ((_549_v9).get(_module.D1.create_DC3(_545_v7))) : ((_542_v4).f23)), globalState));
        let _551_v11;
        _551_v11 = _dafny.Seq.UnicodeFromString("ufps");
        let _552_v12;
        _552_v12 = _dafny.MultiSet.fromElements((_542_v4).fm14(p0, globalState));
        let _553_v13;
        _553_v13 = _dafny.Seq.of(_552_v12);
        let _554_v14;
        let _nw88 = Array((new BigNumber(24)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = p1;
        _nw88[(_dafny.ONE).toNumber()] = p1;
        _nw88[(new BigNumber(2)).toNumber()] = (p1) && (p1);
        _nw88[(new BigNumber(3)).toNumber()] = p1;
        _nw88[(new BigNumber(4)).toNumber()] = true;
        _nw88[(new BigNumber(5)).toNumber()] = !(_dafny.Seq.contains(_538_v0, p0));
        _nw88[(new BigNumber(6)).toNumber()] = p1;
        _nw88[(new BigNumber(7)).toNumber()] = p1;
        _nw88[(new BigNumber(8)).toNumber()] = p1;
        _nw88[(new BigNumber(9)).toNumber()] = !(_540_v2).equals(_dafny.Set.fromElements(p0, p0, p0));
        _nw88[(new BigNumber(10)).toNumber()] = p1;
        _nw88[(new BigNumber(11)).toNumber()] = p1;
        _nw88[(new BigNumber(12)).toNumber()] = (_541_v3)[_module.__default.safeIndex(p0, new BigNumber((_541_v3).length))];
        _nw88[(new BigNumber(13)).toNumber()] = p1;
        _nw88[(new BigNumber(14)).toNumber()] = p1;
        _nw88[(new BigNumber(15)).toNumber()] = (((_550_v10).contains(_543_v5)) ? ((_550_v10).get(_543_v5)) : (p1));
        _nw88[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsPrefixOf(_538_v0, _538_v0);
        _nw88[(new BigNumber(17)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(514)), ((_555_v6) => function (_556_i1) {
          return _555_v6;
        })(_544_v6)), _544_v6);
        _nw88[(new BigNumber(18)).toNumber()] = !((new BigNumber((_551_v11).length)).isLessThanOrEqualTo(new BigNumber((_553_v13).length)));
        _nw88[(new BigNumber(19)).toNumber()] = p1;
        _nw88[(new BigNumber(20)).toNumber()] = true;
        _nw88[(new BigNumber(21)).toNumber()] = (_module.__default.fm0((_542_v4).f23, (_551_v11)[_module.__default.safeIndex(p0, new BigNumber((_551_v11).length))], p1, globalState)) === ((_542_v4).f23);
        _nw88[(new BigNumber(22)).toNumber()] = p1;
        _nw88[(new BigNumber(23)).toNumber()] = p1;
        _554_v14 = _nw88;
        let _index58 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_554_v14).length));
        (_554_v14)[_index58] = (_542_v4).f23;
        let _557_v15;
        _557_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_542_v4).f23);
        let _558_v16;
        _558_v16 = _dafny.Set.fromElements(_540_v2);
        let _559_v17;
        _559_v17 = _dafny.Set.fromElements(p1);
        let _index59 = _module.__default.safeIndex(new BigNumber(252), new BigNumber((_554_v14).length));
        (_554_v14)[_index59] = !((_542_v4).f23) || ((_dafny.Set.fromElements(p1, (((_557_v15).contains(new BigNumber((_558_v16).length))) ? ((_557_v15).get(new BigNumber((_558_v16).length))) : (false)))).IsDisjointFrom(_559_v17));
        let _560_v18;
        _560_v18 = _module.D0.create_DC0(false);
        _560_v18 = _560_v18;
        let _561_v19;
        let _nw89 = new _module.C0();
        _nw89.__ctor(((p1) ? (!(p1)) : ((_554_v14)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_554_v14).length))])));
        _561_v19 = _nw89;
        _544_v6 = _544_v6;
        let _index60 = _module.__default.safeIndex(new BigNumber(323), new BigNumber(((_this).f22).length));
        ((_this).f22)[_index60] = new BigNumber((_541_v3).length);
        let _index61 = _module.__default.safeIndex(new BigNumber(323), new BigNumber(((_this).f22).length));
        ((_this).f22)[_index61] = (p0).plus(new BigNumber((_dafny.Seq.of(p0, new BigNumber((_dafny.Seq.of((_554_v14)[_module.__default.safeIndex(new BigNumber(252), new BigNumber((_554_v14).length))], (_542_v4).f23)).length))).length));
      } else {
        let _562_v20;
        _562_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _563_v21;
        _563_v21 = _dafny.Set.fromElements(_562_v20);
        let _source8 = _module.D3.create_DC9((_563_v21).Union(_563_v21));
        if (_source8.is_DC10) {
          let _564___mcc_h0 = (_source8).cf17;
          let _565_cf17 = _564___mcc_h0;
          let _566_v22;
          _566_v22 = _dafny.Seq.UnicodeFromString("cnnnwq");
          let _567_v23;
          _567_v23 = _dafny.MultiSet.fromElements(_566_v22);
          let _568_v24;
          _568_v24 = _dafny.Seq.of(_566_v22);
          r2 = ((_dafny.ZERO).minus((_565_cf17).plus(new BigNumber((_567_v23).cardinality())))).plus(new BigNumber((_dafny.Seq.Concat(_568_v24, _568_v24)).length));
          let _569_v25;
          _569_v25 = new _dafny.CodePoint('r'.codePointAt(0));
          _562_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(871),_module.__default.fm0(p1, _569_v25, p1, globalState));
          let _570_v26;
          _570_v26 = _module.D1.create_DC5(_566_v22, _565_cf17, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-709)), ((_571_v25) => function (_572_i2) {
  return _571_v25;
})(_569_v25))).length));
          _565_cf17 = (_570_v26).dtor_cf8;
          let _573_v27;
          _573_v27 = _dafny.Seq.of(!(p1));
          let _574_v28;
          let _nw90 = Array((new BigNumber(18)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _573_v27;
          _nw90[(_dafny.ONE).toNumber()] = _573_v27;
          _nw90[(new BigNumber(2)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(3)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(4)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_573_v27, _module.__default.safeIndex(p0, new BigNumber((_573_v27).length)), p1);
          _nw90[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(!(p1));
          _nw90[(new BigNumber(7)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(8)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(9)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(10)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(11)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(12)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(13)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(14)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(15)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(16)).toNumber()] = _573_v27;
          _nw90[(new BigNumber(17)).toNumber()] = _573_v27;
          _574_v28 = _nw90;
          let _575_v29;
          _575_v29 = _module.D1.create_DC3(_574_v28);
          let _pat_let_tv15 = _574_v28;
          let _pat_let_tv16 = _574_v28;
          let _rhs58 = function (_pat_let9_0) {
            return function (_576_dt__update__tmp_h0) {
              return function (_pat_let10_0) {
                return function (_577_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_577_dt__update_hcf3_h0);
                }(_pat_let10_0);
              }(((false) ? (_pat_let_tv15) : (_pat_let_tv16)));
            }(_pat_let9_0);
          }(_575_v29);
          let _rhs59 = p1;
          let _rhs60 = (p0).minus((_565_cf17).multipliedBy(_565_cf17));
          let _rhs61 = p1;
          let _rhs62 = _566_v22;
          let _lhs29 = globalState;
          let _lhs30 = globalState;
          _575_v29 = _rhs58;
          _lhs29.f15 = _rhs59;
          _565_cf17 = _rhs60;
          _lhs30.f17 = _rhs61;
          _566_v22 = _rhs62;
        } else if (_source8.is_DC11) {
          let _578___mcc_h1 = (_source8).cf18;
          let _579___mcc_h2 = (_source8).cf19;
          let _580___mcc_h3 = (_source8).cf20;
          let _581_cf20 = _580___mcc_h3;
          let _582_cf19 = _579___mcc_h2;
          let _583_cf18 = _578___mcc_h1;
          r1 = p1;
          let _584_v30;
          _584_v30 = new _dafny.CodePoint('x'.codePointAt(0));
          let _585_v31;
          let _nw91 = Array((new BigNumber(9)).toNumber());
          _nw91[(_dafny.ZERO).toNumber()] = _584_v30;
          _nw91[(_dafny.ONE).toNumber()] = _584_v30;
          _nw91[(new BigNumber(2)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(3)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(4)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(5)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(6)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(7)).toNumber()] = _584_v30;
          _nw91[(new BigNumber(8)).toNumber()] = _584_v30;
          _585_v31 = _nw91;
          let _586_v32;
          _586_v32 = _dafny.Seq.UnicodeFromString("urvmnj");
          let _587_v33;
          _587_v33 = _dafny.Map.Empty.slice().updateUnsafe(_581_cf20,_586_v32);
          let _588_v34;
          _588_v34 = _dafny.MultiSet.fromElements(_582_cf19);
          let _nw92 = Array((new BigNumber(24)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw92[(_dafny.ONE).toNumber()] = _584_v30;
          _nw92[(new BigNumber(2)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(3)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(4)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(5)).toNumber()] = _module.__default.fm12(_582_cf19, (((_587_v33).contains((_this).f22)) ? ((_587_v33).get((_this).f22)) : (_586_v32)), _module.__default.fm16(false, new BigNumber((_588_v34).cardinality()), globalState), globalState);
          _nw92[(new BigNumber(6)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(7)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(8)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(9)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw92[(new BigNumber(10)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(11)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(12)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(13)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(14)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(15)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(16)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(17)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(18)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(19)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(20)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(21)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(22)).toNumber()] = _584_v30;
          _nw92[(new BigNumber(23)).toNumber()] = _584_v30;
          _585_v31 = _nw92;
          let _589_v35;
          let _nw93 = Array((new BigNumber(12)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = p1;
          _nw93[(_dafny.ONE).toNumber()] = p1;
          _nw93[(new BigNumber(2)).toNumber()] = p1;
          _nw93[(new BigNumber(3)).toNumber()] = p1;
          _nw93[(new BigNumber(4)).toNumber()] = p1;
          _nw93[(new BigNumber(5)).toNumber()] = true;
          _nw93[(new BigNumber(6)).toNumber()] = p1;
          _nw93[(new BigNumber(7)).toNumber()] = p1;
          _nw93[(new BigNumber(8)).toNumber()] = p1;
          _nw93[(new BigNumber(9)).toNumber()] = p1;
          _nw93[(new BigNumber(10)).toNumber()] = p1;
          _nw93[(new BigNumber(11)).toNumber()] = p1;
          _589_v35 = _nw93;
          let _590_v36;
          _590_v36 = _module.D1.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-656)), ((_591_v30) => function (_592_i3) {
  return _591_v30;
})(_584_v30)), _582_cf19, _582_cf19);
          let _593_v37;
          _593_v37 = _dafny.Seq.of(_586_v32, _586_v32, _586_v32, _586_v32, _586_v32);
          let _594_v38;
          _594_v38 = _dafny.MultiSet.fromElements(p1);
          let _595_v39;
          _595_v39 = _dafny.Seq.of(true);
          let _596_v40;
          _596_v40 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),_595_v39);
          let _597_v41;
          let _nw94 = Array((new BigNumber(26)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = p1;
          _nw94[(_dafny.ONE).toNumber()] = p1;
          _nw94[(new BigNumber(2)).toNumber()] = p1;
          _nw94[(new BigNumber(3)).toNumber()] = !(_module.__default.fm0((_module.D1.create_DC4(_582_cf19, p1, _589_v35)).dtor_cf5, _584_v30, p1, globalState));
          _nw94[(new BigNumber(4)).toNumber()] = p1;
          _nw94[(new BigNumber(5)).toNumber()] = !(p1) || (p1);
          _nw94[(new BigNumber(6)).toNumber()] = false;
          _nw94[(new BigNumber(7)).toNumber()] = p1;
          _nw94[(new BigNumber(8)).toNumber()] = false;
          _nw94[(new BigNumber(9)).toNumber()] = false;
          _nw94[(new BigNumber(10)).toNumber()] = (_582_cf19).isLessThan(new BigNumber(((_590_v36).dtor_cf7).length));
          _nw94[(new BigNumber(11)).toNumber()] = p1;
          _nw94[(new BigNumber(12)).toNumber()] = !(p1);
          _nw94[(new BigNumber(13)).toNumber()] = _module.__default.fm0(p1, _584_v30, _module.__default.fm0(p1, _584_v30, p1, globalState), globalState);
          _nw94[(new BigNumber(14)).toNumber()] = p1;
          _nw94[(new BigNumber(15)).toNumber()] = _dafny.areEqual(_593_v37, _593_v37);
          _nw94[(new BigNumber(16)).toNumber()] = p1;
          _nw94[(new BigNumber(17)).toNumber()] = p1;
          _nw94[(new BigNumber(18)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-100)), ((_598_v30) => function (_599_i4) {
            return _598_v30;
          })(_584_v30)), _586_v32);
          _nw94[(new BigNumber(19)).toNumber()] = p1;
          _nw94[(new BigNumber(20)).toNumber()] = (_594_v38).equals(_594_v38);
          _nw94[(new BigNumber(21)).toNumber()] = p1;
          _nw94[(new BigNumber(22)).toNumber()] = p1;
          _nw94[(new BigNumber(23)).toNumber()] = !(p1);
          _nw94[(new BigNumber(24)).toNumber()] = p1;
          _nw94[(new BigNumber(25)).toNumber()] = _dafny.areEqual((((_596_v40).contains(p1)) ? ((_596_v40).get(p1)) : (_595_v39)), _dafny.Seq.update(_595_v39, _module.__default.safeIndex(p0, new BigNumber((_595_v39).length)), false));
          _597_v41 = _nw94;
          let _index62 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_597_v41).length));
          (_597_v41)[_index62] = p1;
          let _index63 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_597_v41).length));
          (_597_v41)[_index63] = _module.__default.fm0(p1, _584_v30, p1, globalState);
          r2 = _module.__default.fm9(new BigNumber(-77), _582_cf19, (_597_v41)[_module.__default.safeIndex(new BigNumber(852), new BigNumber((_597_v41).length))], globalState);
        } else {
          let _600___mcc_h4 = (_source8).cf16;
          let _601_cf16 = _600___mcc_h4;
          let _602_v42;
          let _nw95 = Array((new BigNumber(22)).toNumber()).fill(false);
          _602_v42 = _nw95;
          _602_v42 = ((p1) ? (_602_v42) : (_602_v42));
          let _index64 = _module.__default.safeIndex(new BigNumber(525), new BigNumber(((_this).f22).length));
          ((_this).f22)[_index64] = p0;
          let _index65 = _module.__default.safeIndex(new BigNumber(525), new BigNumber(((_this).f22).length));
          ((_this).f22)[_index65] = new BigNumber(276);
          let _603_v43;
          _603_v43 = new _dafny.CodePoint('a'.codePointAt(0));
          (globalState).f15 = _module.__default.fm0(((false) ? (p1) : (p1)), _603_v43, p1, globalState);
          let _604_v44;
          _604_v44 = _dafny.Seq.of(((_this).f22)[_module.__default.safeIndex(new BigNumber(525), new BigNumber(((_this).f22).length))]);
          let _index66 = _module.__default.safeIndex(new BigNumber(525), new BigNumber(((_this).f22).length));
          ((_this).f22)[_index66] = _module.__default.safeDivisionInt(((p1) ? ((_604_v44)[_module.__default.safeIndex(p0, new BigNumber((_604_v44).length))]) : ((_dafny.ZERO).minus(p0))), ((_this).f22)[_module.__default.safeIndex(new BigNumber(525), new BigNumber(((_this).f22).length))]);
        }
        let _605_v45;
        let _nw96 = Array((new BigNumber(3)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = (p0).isLessThanOrEqualTo(new BigNumber(667));
        _nw96[(_dafny.ONE).toNumber()] = !((p0).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-190)), function (_606_i5) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).length)));
        _nw96[(new BigNumber(2)).toNumber()] = p1;
        _605_v45 = _nw96;
        let _index67 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_605_v45).length));
        (_605_v45)[_index67] = true;
        let _607_v46;
        _607_v46 = _dafny.Seq.of(new BigNumber((_module.__default.fm16(!(p1), p0, globalState)).length));
        let _index68 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_605_v45).length));
        (_605_v45)[_index68] = (_dafny.Seq.contains(_607_v46, new BigNumber(543))) === (p1);
        let _608_v47;
        let _nw97 = Array((new BigNumber(6)).toNumber()).fill([]);
        _608_v47 = _nw97;
        _608_v47 = _608_v47;
        let _609_v48;
        _609_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,(_605_v45)[_module.__default.safeIndex(new BigNumber(123), new BigNumber((_605_v45).length))]);
        let _610_v49;
        _610_v49 = _module.D0.create_DC0((p0).isLessThanOrEqualTo(p0));
        let _pat_let_tv17 = _605_v45;
        let _pat_let_tv18 = _605_v45;
        let _rhs63 = p0;
        let _rhs64 = (_609_v48).Merge(_609_v48);
        let _rhs65 = function (_pat_let11_0) {
          return function (_611_dt__update__tmp_h1) {
            return function (_pat_let12_0) {
              return function (_612_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_612_dt__update_hcf0_h0);
              }(_pat_let12_0);
            }((_pat_let_tv18)[_module.__default.safeIndex(new BigNumber(123), new BigNumber((_pat_let_tv17).length))]);
          }(_pat_let11_0);
        }(_module.D0.create_DC0(p1));
        let _rhs66 = p0;
        let _rhs67 = p0;
        r2 = _rhs63;
        _609_v48 = _rhs64;
        _610_v49 = _rhs65;
        r2 = _rhs66;
        r2 = _rhs67;
        let _613_v50;
        _613_v50 = _dafny.Seq.of((_605_v45)[_module.__default.safeIndex(new BigNumber(123), new BigNumber((_605_v45).length))], p1);
        let _614_v51;
        _614_v51 = _dafny.MultiSet.fromElements(_613_v50, _dafny.Seq.of((_605_v45)[_module.__default.safeIndex(new BigNumber(123), new BigNumber((_605_v45).length))]));
        r2 = ((((_614_v51).update(_613_v50, _module.__default.abs(p0))).IsSubsetOf((_614_v51).update(_613_v50, _module.__default.abs(p0)))) ? (p0) : (p0));
      }
      let _615_v52;
      let _nw98 = new _module.C0();
      _nw98.__ctor(!(p1) || (p1));
      _615_v52 = _nw98;
      let _616_v53;
      _616_v53 = _dafny.Seq.UnicodeFromString("fnymofdaa");
      let _617_v54;
      _617_v54 = _dafny.MultiSet.fromElements(p1, p1);
      let _618_v55;
      _618_v55 = _module.D1.create_DC6(p0, new BigNumber((_617_v54).cardinality()));
      let _pat_let_tv19 = _616_v53;
      let _pat_let_tv20 = _616_v53;
      let _pat_let_tv21 = _616_v53;
      let _pat_let_tv22 = _616_v53;
      let _pat_let_tv23 = _616_v53;
      _616_v53 = function (_source9) {
        if (_source9.is_DC4) {
          let _619___mcc_h5 = (_source9).cf4;
          let _620___mcc_h6 = (_source9).cf5;
          let _621___mcc_h7 = (_source9).cf6;
          let _622_cf6 = _621___mcc_h7;
          let _623_cf5 = _620___mcc_h6;
          let _624_cf4 = _619___mcc_h5;
          return _dafny.Seq.Concat(_pat_let_tv19, _dafny.Seq.UnicodeFromString("wgsdhwirx"));
        } else if (_source9.is_DC5) {
          let _625___mcc_h8 = (_source9).cf7;
          let _626___mcc_h9 = (_source9).cf8;
          let _627___mcc_h10 = (_source9).cf9;
          let _628_cf9 = _627___mcc_h10;
          let _629_cf8 = _626___mcc_h9;
          let _630_cf7 = _625___mcc_h8;
          return _pat_let_tv20;
        } else if (_source9.is_DC6) {
          let _631___mcc_h11 = (_source9).cf10;
          let _632___mcc_h12 = (_source9).cf11;
          let _633_cf11 = _632___mcc_h12;
          let _634_cf10 = _631___mcc_h11;
          return _dafny.Seq.Concat(_pat_let_tv21, _pat_let_tv22);
        } else {
          let _635___mcc_h13 = (_source9).cf3;
          let _636_cf3 = _635___mcc_h13;
          return _pat_let_tv23;
        }
      }(_618_v55);
      let _637_v56;
      _637_v56 = _dafny.Map.Empty.slice().updateUnsafe(p1,_616_v53);
      let _638_v57;
      _638_v57 = _module.D0.create_DC1(p1);
      let _639_v58;
      _639_v58 = _dafny.Set.fromElements(_638_v57);
      let _640_v59;
      _640_v59 = _module.D1.create_DC5((((_637_v56).contains(true)) ? ((_637_v56).get(true)) : (_616_v53)), p0, (_dafny.ZERO).minus(new BigNumber((_639_v58).length)));
      let _641_v60;
      _641_v60 = _module.D3.create_DC10(p0);
      let _642_v61;
      _642_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f22);
      _640_v59 = _module.D1.create_DC5(_616_v53, (_641_v60).dtor_cf17, (p0).minus(new BigNumber((_642_v61).length)));
      let _643_v62;
      let _init16 = ((_644_p0) => function (_645_i7) {
        return _module.__default.safeDivisionInt(_645_i7, new BigNumber((_dafny.MultiSet.fromElements(_644_p0, _644_p0, _644_p0, _644_p0, _644_p0)).cardinality()));
      })(p0);
      let _nw99 = Array((new BigNumber(11)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw99.length); _i0_16++) {
        _nw99[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _643_v62 = _nw99;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_643_v62).length))) {
        let _646_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_646_i6)) && ((_646_i6).isLessThan(new BigNumber((_643_v62).length))))) {
          (_643_v62)[(_646_i6)] = _module.__default.safeDivisionInt(_646_i6, new BigNumber((_617_v54).cardinality()));
        }
      }
      let _647_v63;
      _647_v63 = new _dafny.CodePoint('i'.codePointAt(0));
      let _648_v64;
      _648_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,_647_v63);
      _648_v64 = (_648_v64).update(p0, _647_v63);
      let _649_v65;
      _649_v65 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0));
      r0 = _649_v65;
      let _650_v66;
      _650_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_615_v52).f23);
      let _651_v67;
      _651_v67 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _652_v68;
      _652_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,(((_650_v66).contains(p1)) ? ((_650_v66).get(p1)) : ((_615_v52).fm15(false, (_615_v52).f23, p0, _651_v67, globalState))));
      r1 = (((_652_v68).contains(((((_651_v67).contains(p1)) ? ((_651_v67).get(p1)) : (p0))).plus(new BigNumber(462)))) ? ((_652_v68).get(((((_651_v67).contains(p1)) ? ((_651_v67).get(p1)) : (p0))).plus(new BigNumber(462)))) : (p1));
      r2 = p0;
      let _653_v70;
      let _nw100 = new _module.C0();
      _nw100.__ctor((_615_v52).f23);
      _653_v70 = _nw100;
      let _654_v71;
      _654_v71 = _dafny.Seq.of(_653_v70);
      let _655_v73;
      _655_v73 = _dafny.MultiSet.fromElements(p0);
      let _656_v74;
      _656_v74 = _dafny.Seq.of(p0, new BigNumber((_655_v73).cardinality()), p0);
      let _657_v75;
      _657_v75 = _dafny.Seq.of(new BigNumber((_654_v71).length), p0, new BigNumber((function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_656_v74).Elements) {
          let _658_v72 = _compr_23;
          if (_dafny.Seq.contains(_656_v74, _658_v72)) {
            _coll23.push([(_658_v72).multipliedBy(new BigNumber((_dafny.Seq.of((_615_v52).f23)).length)),new BigNumber(96)]);
          }
        }
        return _coll23;
      }()).length), p0);
      r3 = function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_dafny.MultiSet.FromArray(_657_v75)).Elements) {
          let _659_v69 = _compr_24;
          if ((_dafny.MultiSet.FromArray(_657_v75)).contains(_659_v69)) {
            _coll24.push([_module.__default.safeDivisionInt(_659_v69, p0),(_615_v52).f23]);
          }
        }
        return _coll24;
      }();
      return [r0, r1, r2, r3];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f20 = _module.D0.Default();
      this._f21 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f20, f21) {
      let _this = this;
      (_this).f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    fm8(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qr"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tglyx"), _dafny.Seq.UnicodeFromString("x")));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _660_i0;
      _660_i0 = _dafny.ZERO;
      L1: {
        while (p2) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_660_i0)) {
              break L1;
            }
            _660_i0 = (_660_i0).plus(_dafny.ONE);
            (globalState).f11 = !(p0);
            let _661_v0;
            _661_v0 = _dafny.Seq.UnicodeFromString("qletf");
            let _rhs68 = p2;
            let _rhs69 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_661_v0, _661_v0), _661_v0);
            let _lhs31 = globalState;
            let _lhs32 = globalState;
            _lhs31.f15 = _rhs68;
            _lhs32.f15 = _rhs69;
            let _662_v1;
            _662_v1 = _dafny.MultiSet.fromElements(new BigNumber(328));
            let _663_v2;
            _663_v2 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(456)), function (_664_i1) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            }));
            (globalState).f17 = ((_662_v1).update(p3, _module.__default.abs(new BigNumber((_663_v2).length)))).IsDisjointFrom(_662_v1);
            let _665_v3;
            let _nw101 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
            _665_v3 = _nw101;
            let _666_v4;
            let _nw102 = new _module.C1();
            _nw102.__ctor(_665_v3);
            _666_v4 = _nw102;
            let _667_v5;
            _667_v5 = _module.D4.create_DC12(_666_v4);
            let _668_v6;
            _668_v6 = _dafny.Seq.of(_666_v4, _666_v4);
            let _669_v7;
            _669_v7 = _dafny.Map.Empty.slice().updateUnsafe(p3,_665_v3);
            let _670_v8;
            _670_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_669_v7).contains(p1)) ? ((_669_v7).get(p1)) : (_665_v3)),p1);
            let _671_v9;
            let _nw103 = Array((new BigNumber(12)).toNumber());
            _nw103[(_dafny.ZERO).toNumber()] = (((_this).f21) ? (_666_v4) : ((_667_v5).dtor_cf21));
            _nw103[(_dafny.ONE).toNumber()] = _666_v4;
            _nw103[(new BigNumber(2)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(3)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(4)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(5)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(6)).toNumber()] = (_668_v6)[_module.__default.safeIndex((((_670_v8).contains(_665_v3)) ? ((_670_v8).get(_665_v3)) : (p1)), new BigNumber((_668_v6).length))];
            _nw103[(new BigNumber(7)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(8)).toNumber()] = (((_this).f21) ? (_666_v4) : (_666_v4));
            _nw103[(new BigNumber(9)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(10)).toNumber()] = _666_v4;
            _nw103[(new BigNumber(11)).toNumber()] = _666_v4;
            _671_v9 = _nw103;
            let _index69 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_671_v9).length));
            (_671_v9)[_index69] = (_668_v6)[_module.__default.safeIndex(new BigNumber(199), new BigNumber((_668_v6).length))];
            let _index70 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_671_v9).length));
            (_671_v9)[_index70] = _666_v4;
          }
        }
      }
      let _672_v10;
      _672_v10 = new BigNumber(-235);
      let _673_v11;
      _673_v11 = _dafny.Seq.UnicodeFromString("p");
      let _674_v12;
      _674_v12 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f21);
      _672_v10 = _module.__default.fm9((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_673_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(278)), function (_675_i2) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }))).length)), _module.__default.safeModuloInt(p3, _672_v10), (((_674_v12).contains(p3)) ? ((_674_v12).get(p3)) : (p0)), globalState);
      if (!(!(!(!(p2)))) || ((((_this).f21) ? ((_this).f21) : ((_this).f21)))) {
        let _676_v13;
        _676_v13 = _dafny.Seq.of(_672_v10);
        _672_v10 = (p3).multipliedBy((new BigNumber((_676_v13).length)).minus(p3));
        _672_v10 = _module.__default.fm9(_672_v10, p3, p0, globalState);
        (globalState).f11 = !((_this).f21);
        let _677_v14;
        let _init17 = function (_678_i3) {
          return (_this).f21;
        };
        let _nw104 = Array((new BigNumber(20)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw104.length); _i0_17++) {
          _nw104[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _677_v14 = _nw104;
        let _679_v15;
        _679_v15 = _dafny.Seq.of(_677_v14);
        _679_v15 = _679_v15;
        _672_v10 = (_dafny.ZERO).minus(p3);
      } else {
        let _680_v16;
        let _nw105 = Array((new BigNumber(21)).toNumber()).fill(false);
        _680_v16 = _nw105;
        let _index71 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length));
        (_680_v16)[_index71] = (_672_v10).isLessThan(new BigNumber((_673_v11).length));
        let _index72 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length));
        (_680_v16)[_index72] = p2;
        _680_v16 = _680_v16;
        if (p2) {
          let _681_v17;
          let _init18 = ((_682_v12) => function (_683_i4) {
            return _682_v12;
          })(_674_v12);
          let _nw106 = Array((new BigNumber(15)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw106.length); _i0_18++) {
            _nw106[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _681_v17 = _nw106;
          let _684_v18;
          _684_v18 = _dafny.Seq.of(!(p2), p0);
          let _index73 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_681_v17).length));
          (_681_v17)[_index73] = _module.__default.fm17((_684_v18)[_module.__default.safeIndex(p3, new BigNumber((_684_v18).length))], (_this).f21, globalState);
          let _index74 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_681_v17).length));
          (_681_v17)[_index74] = _674_v12;
          let _685_v19;
          let _nw107 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _685_v19 = _nw107;
          let _686_v20;
          _686_v20 = _dafny.Set.fromElements((_this).f21);
          let _index75 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_685_v19).length));
          (_685_v19)[_index75] = new BigNumber(((_686_v20).Union(_686_v20)).length);
          let _687_v21;
          _687_v21 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index76 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_685_v19).length));
          (_685_v19)[_index76] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_673_v11, _module.__default.safeIndex(p1, new BigNumber((_673_v11).length)), _687_v21), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pddyfcmi"), _673_v11))).length);
          let _688_v22;
          _688_v22 = _dafny.MultiSet.fromElements(p1, p3);
          (globalState).f17 = (new BigNumber(((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_module.D1.create_DC5(_673_v11, _672_v10, (_685_v19)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_685_v19).length))])).dtor_cf7).length))), p1)).Union(_688_v22)).cardinality())).isLessThan((_685_v19)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_685_v19).length))]);
          (globalState).f15 = !(p1).isEqualTo(new BigNumber((((!((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))])) ? (_673_v11) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(84)), ((_689_v21) => function (_690_i5) {
            return _689_v21;
          })(_687_v21))))).length));
          (_this).f20 = _this.f20;
        } else {
          let _691_v23;
          _691_v23 = _dafny.Seq.of((((_674_v12).contains(p3)) ? ((_674_v12).get(p3)) : ((_this).f21)), p2);
          let _692_v24;
          _692_v24 = _module.D4.create_DC14((_691_v23)[_module.__default.safeIndex(p3, new BigNumber((_691_v23).length))], (_dafny.ZERO).minus(_672_v10), _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p1)).cardinality()), new BigNumber(-869), new BigNumber((_674_v12).length)), new _dafny.CodePoint('l'.codePointAt(0)));
          let _693_v25;
          _693_v25 = _dafny.Set.fromElements(_680_v16);
          let _694_v26;
          _694_v26 = _dafny.Map.Empty.slice().updateUnsafe((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))],_693_v25);
          let _rhs70 = _692_v24;
          let _rhs71 = _672_v10;
          let _rhs72 = (((p0) ? (_693_v25) : ((((_694_v26).contains((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))])) ? ((_694_v26).get((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))])) : (_693_v25))))).Union(((p2) ? (_693_v25) : (_693_v25)));
          _692_v24 = _rhs70;
          _672_v10 = _rhs71;
          _693_v25 = _rhs72;
          let _695_v27;
          _695_v27 = new _dafny.CodePoint('w'.codePointAt(0));
          let _696_v28;
          _696_v28 = _dafny.MultiSet.fromElements(_695_v27);
          _696_v28 = _696_v28;
          (globalState).f17 = p2;
          _672_v10 = _module.__default.safeDivisionInt(_672_v10, (new BigNumber((function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of _dafny.IntegerRange(new BigNumber(447), new BigNumber(-428))) {
              let _697_v29 = _compr_25;
              if (((new BigNumber(447)).isLessThanOrEqualTo(_697_v29)) && ((_697_v29).isLessThan(new BigNumber(-428)))) {
                _coll25.push([_module.__default.safeModuloInt(_697_v29, new BigNumber((_dafny.Seq.of((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))])).length)),_691_v23]);
              }
            }
            return _coll25;
          }()).length)).multipliedBy(p1));
          let _698_v30;
          _698_v30 = _dafny.Set.fromElements(new BigNumber(-433));
          let _699_v31;
          _699_v31 = _dafny.Set.fromElements(p3, new BigNumber((_698_v30).length), p3, _module.__default.fm9(_672_v10, p1, p0, globalState));
          (globalState).f11 = (_699_v31).IsSubsetOf(_dafny.Set.fromElements(_module.__default.fm9(_672_v10, p1, p0, globalState)));
        }
        let _index77 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length));
        (_680_v16)[_index77] = !(p3).isEqualTo(p1);
        let _700_v32;
        let _nw108 = new _module.C0();
        _nw108.__ctor((_680_v16)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_680_v16).length))]);
        _700_v32 = _nw108;
      }
      if (p2) {
        let _701_v33;
        let _init19 = ((_702_p3) => function (_703_i6) {
          return _module.__default.safeDivisionInt(_703_i6, _702_p3);
        })(p3);
        let _nw109 = Array((new BigNumber(17)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw109.length); _i0_19++) {
          _nw109[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _701_v33 = _nw109;
        let _704_v34;
        _704_v34 = _dafny.Seq.of(p1, _672_v10);
        let _705_v35;
        _705_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_this).f21);
        let _706_v36;
        _706_v36 = _dafny.Map.Empty.slice().updateUnsafe(_704_v34,_705_v35);
        let _index78 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        (_701_v33)[_index78] = new BigNumber((_706_v36).length);
        let _index79 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        (_701_v33)[_index79] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p3, (_dafny.ZERO).minus(_672_v10))), ((!((_this).f21)) ? (_672_v10) : (_672_v10)));
        let _707_v37;
        let _nw110 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
        _707_v37 = _nw110;
        let _708_v38;
        _708_v38 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("yyioc"));
        let _709_v39;
        _709_v39 = _dafny.Set.fromElements(p0);
        let _710_v40;
        _710_v40 = _module.D4.create_DC13(_672_v10);
        let _index80 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        let _index81 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        let _rhs73 = p1;
        let _rhs74 = _707_v37;
        let _rhs75 = ((((_708_v38).contains(_673_v11)) ? ((_708_v38).get(_673_v11)) : (new BigNumber((_709_v39).length)))).minus((_710_v40).dtor_cf22);
        let _rhs76 = _673_v11;
        let _lhs33 = _701_v33;
        let _lhs34 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        let _lhs35 = _701_v33;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length));
        _lhs33[_lhs34] = _rhs73;
        _707_v37 = _rhs74;
        _lhs35[_lhs36] = _rhs75;
        _673_v11 = _rhs76;
        let _711_v41;
        let _init20 = ((_712_v11) => function (_713_i7) {
          return _712_v11;
        })(_673_v11);
        let _nw111 = Array((new BigNumber(20)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw111.length); _i0_20++) {
          _nw111[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _711_v41 = _nw111;
        let _index82 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_711_v41).length));
        (_711_v41)[_index82] = _673_v11;
        let _714_v42;
        _714_v42 = _dafny.Map.Empty.slice().updateUnsafe(_672_v10,p1);
        let _715_v43;
        _715_v43 = _dafny.Set.fromElements((((_714_v42).contains(p1)) ? ((_714_v42).get(p1)) : (new BigNumber(615))));
        let _716_v44;
        _716_v44 = new _dafny.CodePoint('v'.codePointAt(0));
        let _index83 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_711_v41).length));
        let _rhs77 = _dafny.Seq.update(_dafny.Seq.Concat(_673_v11, _673_v11), _module.__default.safeIndex(_module.__default.fm9(p3, p3, false, globalState), new BigNumber((_dafny.Seq.Concat(_673_v11, _673_v11)).length)), _716_v44);
        let _rhs78 = _dafny.Set.fromElements((p1).plus((_701_v33)[_module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length))]), p1, new BigNumber((_704_v34).length));
        let _lhs37 = _711_v41;
        let _lhs38 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_711_v41).length));
        _lhs37[_lhs38] = _rhs77;
        _715_v43 = _rhs78;
        let _717_v45;
        _717_v45 = _module.D1.create_DC6(p3, (_701_v33)[_module.__default.safeIndex(new BigNumber(322), new BigNumber((_701_v33).length))]);
        _717_v45 = _module.D1.create_DC6(p3, p3);
        let _718_v46;
        let _nw112 = new _module.C0();
        _nw112.__ctor((_this).f21);
        _718_v46 = _nw112;
      } else {
        _672_v10 = new BigNumber(458);
        _672_v10 = (_672_v10).minus(p3);
        let _719_v47;
        let _nw113 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _719_v47 = _nw113;
        let _720_v48;
        let _nw114 = new _module.C1();
        _nw114.__ctor(_719_v47);
        _720_v48 = _nw114;
        let _721_v49;
        _721_v49 = _dafny.Map.Empty.slice().updateUnsafe(_720_v48,p0);
        let _722_v50;
        _722_v50 = _dafny.Map.Empty.slice().updateUnsafe((((_721_v49).contains(_720_v48)) ? ((_721_v49).get(_720_v48)) : ((_this).f21)),(_this).f21);
        _722_v50 = (_722_v50).update((((_674_v12).contains(_672_v10)) ? ((_674_v12).get(_672_v10)) : (p0)), p2);
        _672_v10 = p1;
        let _723_v51;
        _723_v51 = _dafny.Seq.of((_this).f21, !_dafny.areEqual(_673_v11, _673_v11), p0);
        let _rhs79 = !((_723_v51)[_module.__default.safeIndex(p1, new BigNumber((_723_v51).length))]);
        let _rhs80 = p0;
        let _rhs81 = p2;
        let _lhs39 = globalState;
        let _lhs40 = globalState;
        let _lhs41 = globalState;
        _lhs39.f15 = _rhs79;
        _lhs40.f11 = _rhs80;
        _lhs41.f15 = _rhs81;
      }
      if (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-184)), function (_724_i8) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _673_v11)) {
        let _725_v52;
        _725_v52 = _dafny.Set.fromElements(p3);
        _672_v10 = (((_725_v52).equals(_dafny.Set.fromElements(new BigNumber((_674_v12).length)))) ? (p1) : ((_672_v10).plus(p3)));
        _674_v12 = (_674_v12).update(new BigNumber(707), (_module.__default.fm18((_this).f21, _672_v10, p3, globalState)).IsSubsetOf(_725_v52));
        _672_v10 = _672_v10;
        _672_v10 = _672_v10;
        let _726_v53;
        _726_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p1);
        let _727_v54;
        _727_v54 = _dafny.Seq.of(p2, (_this).f21);
        let _728_v55;
        _728_v55 = _dafny.Seq.of(p1);
        let _729_v56;
        let _nw115 = Array((new BigNumber(20)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = new BigNumber((_727_v54).length);
        _nw115[(_dafny.ONE).toNumber()] = p3;
        _nw115[(new BigNumber(2)).toNumber()] = _672_v10;
        _nw115[(new BigNumber(3)).toNumber()] = p1;
        _nw115[(new BigNumber(4)).toNumber()] = p3;
        _nw115[(new BigNumber(5)).toNumber()] = (_728_v55)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f21, true, false, p2)).length), new BigNumber((_728_v55).length))];
        _nw115[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw115[(new BigNumber(7)).toNumber()] = _672_v10;
        _nw115[(new BigNumber(8)).toNumber()] = _672_v10;
        _nw115[(new BigNumber(9)).toNumber()] = p1;
        _nw115[(new BigNumber(10)).toNumber()] = (_728_v55)[_module.__default.safeIndex(p1, new BigNumber((_728_v55).length))];
        _nw115[(new BigNumber(11)).toNumber()] = new BigNumber((_674_v12).length);
        _nw115[(new BigNumber(12)).toNumber()] = p1;
        _nw115[(new BigNumber(13)).toNumber()] = p3;
        _nw115[(new BigNumber(14)).toNumber()] = new BigNumber((_728_v55).length);
        _nw115[(new BigNumber(15)).toNumber()] = p1;
        _nw115[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p3);
        _nw115[(new BigNumber(17)).toNumber()] = p3;
        _nw115[(new BigNumber(18)).toNumber()] = _672_v10;
        _nw115[(new BigNumber(19)).toNumber()] = _672_v10;
        _729_v56 = _nw115;
        let _730_v57;
        _730_v57 = _dafny.MultiSet.fromElements(false, p2, p2, p2);
        let _731_v58;
        _731_v58 = _module.D4.create_DC14((_this).f21, _672_v10, _725_v52, (_673_v11)[_module.__default.safeIndex(new BigNumber((_730_v57).cardinality()), new BigNumber((_673_v11).length))]);
        let _732_v59;
        _732_v59 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC11(_726_v53, new BigNumber(919), _729_v56),_731_v58);
        _672_v10 = new BigNumber((_732_v59).length);
      } else {
        let _733_v60;
        let _init21 = ((_734_v10, _735_p3, _736_p2) => function (_737_i9) {
          return (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_734_v10,new _dafny.CodePoint('s'.codePointAt(0)))).length), _735_p3, new BigNumber((_dafny.Seq.of(_736_p2)).length), _735_p3)).IsDisjointFrom(_dafny.Set.fromElements(_735_p3));
        })(_672_v10, p3, p2);
        let _nw116 = Array((new BigNumber(26)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw116.length); _i0_21++) {
          _nw116[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _733_v60 = _nw116;
        _733_v60 = _733_v60;
        (globalState).f15 = (p0) && (p2);
        if ((_672_v10).isLessThan(_672_v10)) {
          let _738_v61;
          _738_v61 = new _dafny.CodePoint('x'.codePointAt(0));
          (globalState).f17 = !_dafny.Seq.contains(_673_v11, _738_v61);
          let _739_v62;
          let _init22 = ((_740_p3) => function (_741_i10) {
            return _module.D4.create_DC13(_740_p3);
          })(p3);
          let _nw117 = Array((new BigNumber(4)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw117.length); _i0_22++) {
            _nw117[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _739_v62 = _nw117;
          let _742_v63;
          _742_v63 = _dafny.Map.Empty.slice().updateUnsafe(p3,_739_v62);
          let _743_v64;
          let _nw118 = Array((new BigNumber(21)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _739_v62;
          _nw118[(_dafny.ONE).toNumber()] = _739_v62;
          _nw118[(new BigNumber(2)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(3)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(4)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(5)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(6)).toNumber()] = ((p0) ? (_739_v62) : (_739_v62));
          _nw118[(new BigNumber(7)).toNumber()] = (((_this).f21) ? (_739_v62) : (_739_v62));
          _nw118[(new BigNumber(8)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(9)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(10)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(11)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(12)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(13)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(14)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(15)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(16)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(17)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(18)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(19)).toNumber()] = _739_v62;
          _nw118[(new BigNumber(20)).toNumber()] = (((_742_v63).contains(p1)) ? ((_742_v63).get(p1)) : (_739_v62));
          _743_v64 = _nw118;
          _743_v64 = ((p0) ? (_743_v64) : (_743_v64));
          let _744_v65;
          let _nw119 = new _module.C0();
          _nw119.__ctor(p2);
          _744_v65 = _nw119;
          let _745_v66;
          _745_v66 = _dafny.Map.Empty.slice().updateUnsafe(p2,_744_v65);
          let _746_v67;
          _746_v67 = _module.D4.create_DC12((((_745_v66).contains((_this).f21)) ? ((_745_v66).get((_this).f21)) : (_744_v65)));
          let _pat_let_tv24 = _744_v65;
          _746_v67 = ((p2) ? (_746_v67) : (function (_pat_let13_0) {
            return function (_747_dt__update__tmp_h0) {
              return function (_pat_let14_0) {
                return function (_748_dt__update_hcf21_h0) {
                  return _module.D4.create_DC12(_748_dt__update_hcf21_h0);
                }(_pat_let14_0);
              }(_pat_let_tv24);
            }(_pat_let13_0);
          }(_746_v67)));
          let _index84 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_733_v60).length));
          (_733_v60)[_index84] = p0;
          let _749_v68;
          let _nw120 = Array((new BigNumber(4)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _672_v10;
          _nw120[(_dafny.ONE).toNumber()] = p1;
          _nw120[(new BigNumber(2)).toNumber()] = new BigNumber(9);
          _nw120[(new BigNumber(3)).toNumber()] = p3;
          _749_v68 = _nw120;
          let _750_v69;
          _750_v69 = _dafny.Map.Empty.slice().updateUnsafe(_733_v60,_749_v68);
          let _751_v70;
          _751_v70 = _dafny.Seq.of(p1);
          let _752_v71;
          _752_v71 = _dafny.Set.fromElements(true);
          let _753_v72;
          let _nw121 = Array((new BigNumber(24)).toNumber());
          _nw121[(_dafny.ZERO).toNumber()] = new BigNumber(505);
          _nw121[(_dafny.ONE).toNumber()] = p1;
          _nw121[(new BigNumber(2)).toNumber()] = p3;
          _nw121[(new BigNumber(3)).toNumber()] = new BigNumber((_750_v69).length);
          _nw121[(new BigNumber(4)).toNumber()] = new BigNumber(499);
          _nw121[(new BigNumber(5)).toNumber()] = new BigNumber((_751_v70).length);
          _nw121[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(634)), ((_754_v61) => function (_755_i11) {
            return _754_v61;
          })(_738_v61))).length);
          _nw121[(new BigNumber(7)).toNumber()] = _672_v10;
          _nw121[(new BigNumber(8)).toNumber()] = _672_v10;
          _nw121[(new BigNumber(9)).toNumber()] = _672_v10;
          _nw121[(new BigNumber(10)).toNumber()] = new BigNumber((_752_v71).length);
          _nw121[(new BigNumber(11)).toNumber()] = p1;
          _nw121[(new BigNumber(12)).toNumber()] = p1;
          _nw121[(new BigNumber(13)).toNumber()] = p1;
          _nw121[(new BigNumber(14)).toNumber()] = new BigNumber(235);
          _nw121[(new BigNumber(15)).toNumber()] = _module.__default.fm9(new BigNumber((_673_v11).length), p3, true, globalState);
          _nw121[(new BigNumber(16)).toNumber()] = _672_v10;
          _nw121[(new BigNumber(17)).toNumber()] = p3;
          _nw121[(new BigNumber(18)).toNumber()] = new BigNumber(181);
          _nw121[(new BigNumber(19)).toNumber()] = new BigNumber(772);
          _nw121[(new BigNumber(20)).toNumber()] = new BigNumber((_673_v11).length);
          _nw121[(new BigNumber(21)).toNumber()] = _672_v10;
          _nw121[(new BigNumber(22)).toNumber()] = p3;
          _nw121[(new BigNumber(23)).toNumber()] = new BigNumber(-85);
          _753_v72 = _nw121;
          let _756_v73;
          let _nw122 = Array((new BigNumber(3)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _753_v72;
          _nw122[(_dafny.ONE).toNumber()] = _749_v68;
          _nw122[(new BigNumber(2)).toNumber()] = _753_v72;
          _756_v73 = _nw122;
          let _index85 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_733_v60).length));
          let _rhs82 = _756_v73;
          let _rhs83 = p2;
          let _lhs42 = globalState;
          let _lhs43 = _733_v60;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_733_v60).length));
          _lhs42.f2 = _rhs82;
          _lhs43[_lhs44] = _rhs83;
          let _757_v74;
          _757_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p3);
          _757_v74 = (_757_v74).update(p2, p1);
        } else {
          _672_v10 = (_dafny.ZERO).minus(p3);
          (globalState).f11 = !_dafny.areEqual(_673_v11, _673_v11);
          (globalState).f17 = !((_this).f21);
          let _758_v75;
          _758_v75 = new _dafny.CodePoint('b'.codePointAt(0));
          let _759_v76;
          let _init23 = ((_760_p1) => function (_761_i12) {
            return (_761_i12).multipliedBy(_760_p1);
          })(p1);
          let _nw123 = Array((new BigNumber(2)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw123.length); _i0_23++) {
            _nw123[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _759_v76 = _nw123;
          let _index86 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_759_v76).length));
          (_759_v76)[_index86] = ((p0) ? (p1) : (new BigNumber(185)));
          let _762_v77;
          _762_v77 = _dafny.Set.fromElements(p2);
          let _763_v78;
          _763_v78 = _dafny.Seq.of(p0, p0, p2, p2, p0);
          let _index87 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_759_v76).length));
          let _rhs84 = false;
          let _rhs85 = new _dafny.CodePoint('q'.codePointAt(0));
          let _rhs86 = _672_v10;
          let _rhs87 = _module.__default.fm9((_module.__default.fm9(new BigNumber((_673_v11).length), p1, p2, globalState)).plus(new BigNumber((_762_v77).length)), _module.__default.safeDivisionInt(_672_v10, new BigNumber(636)), !(p3).isEqualTo(new BigNumber((_763_v78).length)), globalState);
          let _rhs88 = new BigNumber(958);
          let _lhs45 = globalState;
          let _lhs46 = _759_v76;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_759_v76).length));
          _lhs45.f17 = _rhs84;
          _758_v75 = _rhs85;
          _672_v10 = _rhs86;
          _lhs46[_lhs47] = _rhs87;
          _672_v10 = _rhs88;
          (globalState).f15 = p0;
        }
        let _764_v79;
        let _nw124 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _764_v79 = _nw124;
        let _765_v80;
        let _nw125 = new _module.C1();
        _nw125.__ctor(_764_v79);
        _765_v80 = _nw125;
        let _index88 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_733_v60).length));
        (_733_v60)[_index88] = !((_this).f21);
        let _index89 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_733_v60).length));
        (_733_v60)[_index89] = !((_this).f21);
      }
      let _766_v81;
      _766_v81 = _dafny.MultiSet.fromElements(true);
      let _hi6 = new BigNumber((_766_v81).cardinality());
      for (let _767_i13 = new BigNumber(-878); _767_i13.isLessThan(_hi6); _767_i13 = _767_i13.plus(_dafny.ONE)) {
        let _768_v82;
        _768_v82 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f21)).length));
        let _769_v83;
        _769_v83 = _dafny.Set.fromElements(p3);
        let _770_v84;
        _770_v84 = new _dafny.CodePoint('p'.codePointAt(0));
        let _771_v85;
        _771_v85 = _module.D4.create_DC14((_this).f21, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(477)), function (_772_i14) {
  return new _dafny.CodePoint('r'.codePointAt(0));
})).length), _767_i13, _672_v10)).length), _769_v83, _770_v84);
        let _773_v86;
        _773_v86 = _dafny.Map.Empty.slice().updateUnsafe((((_768_v82).contains(p2)) ? ((_768_v82).get(p2)) : ((_771_v85).dtor_cf24)),new BigNumber(264));
        _672_v10 = (_672_v10).minus(new BigNumber((((_773_v86).update(new BigNumber(603), p3)).Merge(_773_v86)).length));
        _674_v12 = (_674_v12).update(p1, !(_773_v86).contains(_672_v10));
        let _774_v87;
        _774_v87 = _dafny.Set.fromElements(p0);
        let _775_v88;
        _775_v88 = _dafny.Seq.of(_774_v87);
        _672_v10 = (p1).multipliedBy(new BigNumber((_dafny.Seq.Concat(_775_v88, _775_v88)).length));
        let _776_v89;
        let _nw126 = Array((new BigNumber(16)).toNumber()).fill([]);
        _776_v89 = _nw126;
        let _777_v90;
        _777_v90 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f21);
        let _778_v91;
        _778_v91 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_767_i13,p2),false);
        let _779_v93;
        _779_v93 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)))).length));
        let _780_v94;
        let _nw127 = Array((new BigNumber(29)).toNumber());
        _nw127[(_dafny.ZERO).toNumber()] = p0;
        _nw127[(_dafny.ONE).toNumber()] = p2;
        _nw127[(new BigNumber(2)).toNumber()] = p0;
        _nw127[(new BigNumber(3)).toNumber()] = (_this).f21;
        _nw127[(new BigNumber(4)).toNumber()] = p0;
        _nw127[(new BigNumber(5)).toNumber()] = p0;
        _nw127[(new BigNumber(6)).toNumber()] = !((_this).f21);
        _nw127[(new BigNumber(7)).toNumber()] = false;
        _nw127[(new BigNumber(8)).toNumber()] = (_this).f21;
        _nw127[(new BigNumber(9)).toNumber()] = p0;
        _nw127[(new BigNumber(10)).toNumber()] = false;
        _nw127[(new BigNumber(11)).toNumber()] = !(p0);
        _nw127[(new BigNumber(12)).toNumber()] = (_this).f21;
        _nw127[(new BigNumber(13)).toNumber()] = p2;
        _nw127[(new BigNumber(14)).toNumber()] = p0;
        _nw127[(new BigNumber(15)).toNumber()] = p0;
        _nw127[(new BigNumber(16)).toNumber()] = _module.__default.fm0(p0, _770_v84, (_this).f21, globalState);
        _nw127[(new BigNumber(17)).toNumber()] = _module.__default.fm0((_this).f21, _770_v84, (_this).f21, globalState);
        _nw127[(new BigNumber(18)).toNumber()] = p2;
        _nw127[(new BigNumber(19)).toNumber()] = p2;
        _nw127[(new BigNumber(20)).toNumber()] = (_this).f21;
        _nw127[(new BigNumber(21)).toNumber()] = !((((_777_v90).contains((_this).f21)) ? ((_777_v90).get((_this).f21)) : ((((_778_v91).contains(function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_779_v93).Elements) {
            let _782_v92 = _compr_27;
            if (_dafny.Seq.contains(_779_v93, _782_v92)) {
              _coll27.push([_module.__default.safeModuloInt(_782_v92, _767_i13),(_this).f21]);
            }
          }
          return _coll27;
        }())) ? ((_778_v91).get(function () {
          let _coll26 = new _dafny.Map();
          for (const _compr_26 of (_779_v93).Elements) {
            let _781_v92 = _compr_26;
            if (_dafny.Seq.contains(_779_v93, _781_v92)) {
              _coll26.push([_module.__default.safeModuloInt(_781_v92, _767_i13),(_this).f21]);
            }
          }
          return _coll26;
        }())) : (p0)))));
        _nw127[(new BigNumber(22)).toNumber()] = p2;
        _nw127[(new BigNumber(23)).toNumber()] = true;
        _nw127[(new BigNumber(24)).toNumber()] = p2;
        _nw127[(new BigNumber(25)).toNumber()] = p0;
        _nw127[(new BigNumber(26)).toNumber()] = p0;
        _nw127[(new BigNumber(27)).toNumber()] = (_this).f21;
        _nw127[(new BigNumber(28)).toNumber()] = (_this).f21;
        _780_v94 = _nw127;
        let _index90 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_776_v89).length));
        (_776_v89)[_index90] = _780_v94;
        let _index91 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_776_v89).length));
        (_776_v89)[_index91] = (_module.D1.create_DC4(_672_v10, (_this).f21, _780_v94)).dtor_cf6;
      }
      r0 = p2;
      let _783_v95;
      _783_v95 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _784_v96;
      _784_v96 = new _dafny.CodePoint('h'.codePointAt(0));
      let _785_v97;
      _785_v97 = _module.D3.create_DC10(_672_v10);
      let _786_v98;
      _786_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_module.__default.fm11(_672_v10, (((_783_v95).contains(_module.__default.fm0(p0, _784_v96, p2, globalState))) ? ((_783_v95).get(_module.__default.fm0(p0, _784_v96, p2, globalState))) : ((_this).f21)), p2, (_785_v97).dtor_cf17, globalState));
      r1 = !(((new BigNumber((_786_v98).length)).minus(p3)).isLessThan((_dafny.ZERO).minus(p3)));
      return [r0, r1];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _787_v0;
      _787_v0 = _module.D4.create_DC13(new BigNumber((_dafny.Seq.of((_this).f21)).length));
      let _788_v1;
      _788_v1 = new BigNumber(992);
      let _789_v2;
      _789_v2 = new _dafny.CodePoint('e'.codePointAt(0));
      let _790_v3;
      _790_v3 = _module.D4.create_DC14(false, (_dafny.ZERO).minus(_788_v1), _dafny.Set.fromElements(_788_v1), _789_v2);
      let _791_v4;
      _791_v4 = _dafny.Seq.of(_788_v1);
      let _792_v5;
      _792_v5 = _module.D4.create_DC14(!((_this).f21), _788_v1, _dafny.Set.fromElements(_788_v1, _788_v1, (_790_v3).dtor_cf24, new BigNumber((_module.__default.fm19(_791_v4, true, globalState)).length)), _789_v2);
      let _pat_let_tv25 = _787_v0;
      let _pat_let_tv26 = _788_v1;
      let _pat_let_tv27 = _787_v0;
      let _pat_let_tv28 = _787_v0;
      _787_v0 = function (_source10) {
        if (_source10.is_DC13) {
          let _793___mcc_h0 = (_source10).cf22;
          let _794_cf22 = _793___mcc_h0;
          let _795_dt__update__tmp_h0 = _pat_let_tv25;
          let _796_dt__update_hcf22_h0 = _pat_let_tv26;
          return _module.D4.create_DC13(_796_dt__update_hcf22_h0);
        } else if (_source10.is_DC14) {
          let _797___mcc_h1 = (_source10).cf23;
          let _798___mcc_h2 = (_source10).cf24;
          let _799___mcc_h3 = (_source10).cf25;
          let _800___mcc_h4 = (_source10).cf26;
          let _801_cf26 = _800___mcc_h4;
          let _802_cf25 = _799___mcc_h3;
          let _803_cf24 = _798___mcc_h2;
          let _804_cf23 = _797___mcc_h1;
          return _module.D4.create_DC13(_803_cf24);
        } else if (_source10.is_DC12) {
          let _805___mcc_h5 = (_source10).cf21;
          let _806_cf21 = _805___mcc_h5;
          return _pat_let_tv27;
        } else {
          let _807___mcc_h6 = (_source10).cf27;
          let _808_cf27 = _807___mcc_h6;
          return _pat_let_tv28;
        }
      }(_790_v3);
      if ((_this).f21) {
        let _809_v6;
        _809_v6 = _dafny.Set.fromElements(new BigNumber((p1).length), _788_v1, _788_v1);
        let _810_v7;
        _810_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_809_v6).IsProperSubsetOf(_809_v6)));
        _788_v1 = new BigNumber((_810_v7).length);
        let _811_v9;
        let _nw128 = Array((new BigNumber(28)).toNumber()).fill(false);
        _811_v9 = _nw128;
        let _812_v10;
        _812_v10 = _module.D1.create_DC4((_dafny.ZERO).minus(_788_v1), false, _811_v9);
        let _813_v11;
        _813_v11 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,(_812_v10).dtor_cf5);
        let _814_v12;
        _814_v12 = _module.D3.create_DC9(_dafny.Set.fromElements(function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(527), new BigNumber(468))) {
    let _815_v8 = _compr_28;
    if (((new BigNumber(527)).isLessThanOrEqualTo(_815_v8)) && ((_815_v8).isLessThan(new BigNumber(468)))) {
      _coll28.push([(_815_v8).plus(_788_v1),(_this).f21]);
    }
  }
  return _coll28;
}(), _813_v11));
        let _816_v13;
        _816_v13 = _dafny.Set.fromElements(_813_v11, _813_v11);
        _814_v12 = _module.D3.create_DC9(_816_v13);
        let _817_v14;
        let _nw129 = Array((new BigNumber(5)).toNumber());
        _nw129[(_dafny.ZERO).toNumber()] = _788_v1;
        _nw129[(_dafny.ONE).toNumber()] = new BigNumber(418);
        _nw129[(new BigNumber(2)).toNumber()] = _788_v1;
        _nw129[(new BigNumber(3)).toNumber()] = _788_v1;
        _nw129[(new BigNumber(4)).toNumber()] = _788_v1;
        _817_v14 = _nw129;
        let _818_v15;
        let _nw130 = new _module.C1();
        _nw130.__ctor(_817_v14);
        _818_v15 = _nw130;
        let _819_v16;
        _819_v16 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,_788_v1);
        let _820_v17;
        _820_v17 = _dafny.Seq.UnicodeFromString("csvwaexvd");
        _813_v11 = (_813_v11).update((_dafny.ZERO).minus((((_819_v16).contains(_788_v1)) ? ((_819_v16).get(_788_v1)) : (new BigNumber((_820_v17).length)))), !(p0));
        let _821_v18;
        _821_v18 = _dafny.Set.fromElements(p0);
        _821_v18 = _821_v18;
      } else {
        let _822_v19;
        _822_v19 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_788_v1)).update(p0, _788_v1)).length));
        let _823_v20;
        let _nw131 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.of());
        _823_v20 = _nw131;
        let _824_v21;
        _824_v21 = _dafny.Seq.of(_823_v20);
        let _825_v22;
        _825_v22 = _module.D1.create_DC3((_824_v21)[_module.__default.safeIndex(new BigNumber((_module.__default.fm20(globalState)).length), new BigNumber((_824_v21).length))]);
        let _826_v23;
        _826_v23 = _dafny.MultiSet.fromElements(_825_v22);
        _822_v19 = (_822_v19).update(_module.__default.safeModuloInt(_788_v1, _788_v1), (new BigNumber((_826_v23).cardinality())).minus(_788_v1));
        _788_v1 = ((_dafny.ZERO).minus(_788_v1)).plus((_788_v1).minus(new BigNumber((p1).length)));
        let _827_v24;
        _827_v24 = _dafny.Seq.of(_789_v2, _789_v2, _789_v2);
        let _828_v25;
        let _nw132 = new _module.C0();
        _nw132.__ctor(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("qgaehayk"), (_827_v24)[_module.__default.safeIndex(_788_v1, new BigNumber((_827_v24).length))]));
        _828_v25 = _nw132;
        _828_v25 = _828_v25;
        _827_v24 = _dafny.Seq.Concat(_827_v24, _827_v24);
        let _rhs89 = _module.__default.safeDivisionInt(_module.__default.fm9(_788_v1, new BigNumber((p1).length), (_this).f21, globalState), _module.__default.fm9(_788_v1, (_791_v4)[_module.__default.safeIndex(_788_v1, new BigNumber((_791_v4).length))], p0, globalState));
        let _rhs90 = ((p0) ? (_827_v24) : (_827_v24));
        _788_v1 = _rhs89;
        _827_v24 = _rhs90;
      }
      let _829_i0;
      _829_i0 = _dafny.ZERO;
      L2: {
        while ((_this).f21) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_829_i0)) {
              break L2;
            }
            _829_i0 = (_829_i0).plus(_dafny.ONE);
            if (((p0) ? ((_this).f21) : (!((_this).f21)))) {
              (globalState).f17 = !((_this).f21);
              let _830_v26;
              let _init24 = ((_831_p0) => function (_832_i1) {
                return _831_p0;
              })(p0);
              let _nw133 = Array((new BigNumber(23)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw133.length); _i0_24++) {
                _nw133[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _830_v26 = _nw133;
              let _833_v27;
              _833_v27 = _dafny.Map.Empty.slice().updateUnsafe(_830_v26,_module.__default.fm9((_dafny.ZERO).minus(_788_v1), new BigNumber((_dafny.Seq.of(p0)).length), p0, globalState));
              _833_v27 = (_833_v27).update(_830_v26, (_788_v1).minus(_788_v1));
              let _834_v28;
              let _nw134 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _834_v28 = _nw134;
              _834_v28 = _834_v28;
              let _835_v29;
              _835_v29 = _dafny.Set.fromElements(new BigNumber((_791_v4).length), _788_v1);
              _788_v1 = _module.__default.fm9(new BigNumber(((_dafny.Set.fromElements(_788_v1)).Union(_835_v29)).length), _module.__default.safeDivisionInt(_788_v1, _788_v1), !((_this).f21), globalState);
              (globalState).f15 = (_this).f21;
            } else {
              let _836_v30;
              let _nw135 = Array((new BigNumber(11)).toNumber()).fill(false);
              _836_v30 = _nw135;
              let _index92 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_836_v30).length));
              (_836_v30)[_index92] = p0;
              let _837_v31;
              _837_v31 = _dafny.MultiSet.fromElements(_836_v30);
              let _index93 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_836_v30).length));
              (_836_v30)[_index93] = _module.__default.fm0(!((_this).f21), _789_v2, (_837_v31).IsSubsetOf(_837_v31), globalState);
              let _838_v32;
              _838_v32 = _dafny.Seq.UnicodeFromString("cb");
              _838_v32 = _dafny.Seq.Concat(_838_v32, _dafny.Seq.Concat(_838_v32, _838_v32));
              _788_v1 = _788_v1;
              let _839_v33;
              let _nw136 = new _module.C0();
              _nw136.__ctor((_this).f21);
              _839_v33 = _nw136;
              let _840_v34;
              let _nw137 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
              _840_v34 = _nw137;
              let _index94 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_840_v34).length));
              (_840_v34)[_index94] = _788_v1;
              let _index95 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_840_v34).length));
              (_840_v34)[_index95] = _788_v1;
            }
            if ((p0) && ((_this).f21)) {
              let _841_v35;
              let _nw138 = new _module.C0();
              _nw138.__ctor((_this).f21);
              _841_v35 = _nw138;
              (globalState).f17 = (_841_v35).f23;
              _789_v2 = _789_v2;
              let _842_v36;
              let _nw139 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _842_v36 = _nw139;
              let _843_v37;
              _843_v37 = _dafny.Seq.UnicodeFromString("nbewbhblg");
              let _index96 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_842_v36).length));
              (_842_v36)[_index96] = _843_v37;
              let _index97 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_842_v36).length));
              (_842_v36)[_index97] = _dafny.Seq.UnicodeFromString("jksnhhduo");
              let _844_v38;
              _844_v38 = _dafny.Set.fromElements(new BigNumber(478));
              _788_v1 = new BigNumber((((_dafny.Set.fromElements(new BigNumber(((_842_v36)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_842_v36).length))]).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(_788_v1)))).Difference(_844_v38)).Intersect(function () {
                let _coll29 = new _dafny.Set();
                for (const _compr_29 of _dafny.IntegerRange(new BigNumber(-957), new BigNumber(919))) {
                  let _845_v39 = _compr_29;
                  if (((new BigNumber(-957)).isLessThanOrEqualTo(_845_v39)) && ((_845_v39).isLessThan(new BigNumber(919)))) {
                    _coll29.add((_845_v39).multipliedBy(_788_v1));
                  }
                }
                return _coll29;
              }())).length);
            } else {
              let _846_v40;
              _846_v40 = _dafny.Seq.UnicodeFromString("p");
              let _847_v41;
              _847_v41 = _module.D1.create_DC5(_846_v40, (_dafny.ZERO).minus(_788_v1), _788_v1);
              _788_v1 = new BigNumber((((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), ((_848_v1) => function (_849_i2) {
                return _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("dyabgd"), _848_v1, _848_v1);
              })(_788_v1))) : (_dafny.Seq.of(_847_v41, _847_v41)))).length);
              _788_v1 = (new BigNumber((_791_v4).length)).multipliedBy(_788_v1);
              let _850_v42;
              let _nw140 = new _module.C0();
              _nw140.__ctor((new BigNumber((_dafny.Seq.update(_846_v40, _module.__default.safeIndex(_788_v1, new BigNumber((_846_v40).length)), _789_v2)).length)).isLessThan(_788_v1));
              _850_v42 = _nw140;
              let _851_v43;
              _851_v43 = _dafny.Seq.of(_846_v40);
              _788_v1 = new BigNumber(((((_850_v42).f23) ? (_dafny.Seq.Concat(_846_v40, _dafny.Seq.update((_851_v43)[_module.__default.safeIndex(_788_v1, new BigNumber((_851_v43).length))], _module.__default.safeIndex(_788_v1, new BigNumber(((_851_v43)[_module.__default.safeIndex(_788_v1, new BigNumber((_851_v43).length))]).length)), _789_v2))) : (_846_v40))).length);
              let _rhs91 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_788_v1), new BigNumber(((_this).fm8(p0, globalState)).length));
              let _rhs92 = (p1)[_module.__default.safeIndex(new BigNumber(518), new BigNumber((p1).length))];
              let _rhs93 = (((_850_v42).f23) ? (new BigNumber(-637)) : (_788_v1));
              let _lhs48 = globalState;
              _788_v1 = _rhs91;
              _lhs48.f15 = _rhs92;
              _788_v1 = _rhs93;
            }
            let _852_v44;
            let _init25 = function (_853_i3) {
              return (_this).f21;
            };
            let _nw141 = Array((new BigNumber(22)).toNumber());
            for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw141.length); _i0_25++) {
              _nw141[_i0_25] = _init25(new BigNumber(_i0_25));
            }
            _852_v44 = _nw141;
            _852_v44 = _852_v44;
            (globalState).f11 = (_this).f21;
          }
        }
      }
      let _854_i4;
      _854_i4 = _dafny.ZERO;
      L3: {
        while (!(_788_v1).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(_788_v1, _788_v1, _788_v1)).cardinality()))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_854_i4)) {
              break L3;
            }
            _854_i4 = (_854_i4).plus(_dafny.ONE);
            (globalState).f17 = (_this).f21;
            let _855_v45;
            let _init26 = ((_856_v4) => function (_857_i5) {
              return _856_v4;
            })(_791_v4);
            let _nw142 = Array((new BigNumber(13)).toNumber());
            for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw142.length); _i0_26++) {
              _nw142[_i0_26] = _init26(new BigNumber(_i0_26));
            }
            _855_v45 = _nw142;
            let _index98 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_855_v45).length));
            (_855_v45)[_index98] = _module.__default.fm19(_791_v4, p0, globalState);
            let _index99 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_855_v45).length));
            (_855_v45)[_index99] = _dafny.Seq.Concat(_791_v4, _791_v4);
            let _rhs94 = !_dafny.Seq.contains((_855_v45)[_module.__default.safeIndex(new BigNumber(158), new BigNumber((_855_v45).length))], _788_v1);
            let _rhs95 = new BigNumber(459);
            let _lhs49 = globalState;
            _lhs49.f11 = _rhs94;
            _788_v1 = _rhs95;
            _788_v1 = _module.__default.fm9(_788_v1, (((_this).f21) ? (_788_v1) : (_788_v1)), p0, globalState);
          }
        }
      }
      let _858_v46;
      let _nw143 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _858_v46 = _nw143;
      let _859_v47;
      _859_v47 = _module.D1.create_DC3(_858_v46);
      let _860_v48;
      let _nw144 = Array((new BigNumber(7)).toNumber());
      _nw144[(_dafny.ZERO).toNumber()] = _858_v46;
      _nw144[(_dafny.ONE).toNumber()] = _858_v46;
      _nw144[(new BigNumber(2)).toNumber()] = _858_v46;
      _nw144[(new BigNumber(3)).toNumber()] = _858_v46;
      _nw144[(new BigNumber(4)).toNumber()] = _858_v46;
      _nw144[(new BigNumber(5)).toNumber()] = _858_v46;
      _nw144[(new BigNumber(6)).toNumber()] = (_859_v47).dtor_cf3;
      _860_v48 = _nw144;
      let _index100 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_860_v48).length));
      (_860_v48)[_index100] = _858_v46;
      let _index101 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_860_v48).length));
      (_860_v48)[_index101] = _858_v46;
      if ((_this).f21) {
        let _861_v49;
        _861_v49 = _dafny.Seq.UnicodeFromString("trlsw");
        _788_v1 = (new BigNumber((_861_v49).length)).plus(new BigNumber((_dafny.Seq.Concat(_791_v4, _791_v4)).length));
        if (!((p1)[_module.__default.safeIndex(_module.__default.fm9(_788_v1, new BigNumber((p1).length), !((_this).f21), globalState), new BigNumber((p1).length))])) {
          let _862_v50;
          let _nw145 = Array((new BigNumber(5)).toNumber()).fill(false);
          _862_v50 = _nw145;
          let _index102 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_862_v50).length));
          (_862_v50)[_index102] = p0;
          let _index103 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_862_v50).length));
          let _rhs96 = (_this).f21;
          let _rhs97 = _788_v1;
          let _lhs50 = _862_v50;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_862_v50).length));
          _lhs50[_lhs51] = _rhs96;
          _788_v1 = _rhs97;
          _788_v1 = (_dafny.ZERO).minus(_788_v1);
          _862_v50 = _862_v50;
          let _863_v51;
          _863_v51 = _dafny.MultiSet.fromElements(p0, (_862_v50)[_module.__default.safeIndex(new BigNumber(474), new BigNumber((_862_v50).length))], (_this).f21);
          let _index104 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_862_v50).length));
          (_862_v50)[_index104] = (_863_v51).IsSubsetOf(_863_v51);
          _788_v1 = _module.__default.safeModuloInt(new BigNumber(468), _788_v1);
        } else {
          _861_v49 = (_module.__default.fm21(globalState)).dtor_cf7;
          let _864_v52;
          let _nw146 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _864_v52 = _nw146;
          let _865_v53;
          _865_v53 = _dafny.MultiSet.fromElements(_864_v52);
          (globalState).f15 = (((_dafny.MultiSet.fromElements(_864_v52, _864_v52)).IsSubsetOf(_865_v53)) ? ((_this).f21) : ((_this).f21));
          let _866_v54;
          _866_v54 = _dafny.Set.fromElements(_dafny.Seq.of(p0, p0));
          let _867_v55;
          _867_v55 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qn"));
          let _868_v56;
          _868_v56 = _dafny.Seq.of(_867_v55);
          let _869_v57;
          let _nw147 = Array((new BigNumber(20)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = p0;
          _nw147[(_dafny.ONE).toNumber()] = (_this).f21;
          _nw147[(new BigNumber(2)).toNumber()] = false;
          _nw147[(new BigNumber(3)).toNumber()] = (_866_v54).IsProperSubsetOf(_866_v54);
          _nw147[(new BigNumber(4)).toNumber()] = _module.__default.fm0((_this).f21, _789_v2, (_this).f21, globalState);
          _nw147[(new BigNumber(5)).toNumber()] = p0;
          _nw147[(new BigNumber(6)).toNumber()] = p0;
          _nw147[(new BigNumber(7)).toNumber()] = (_this).f21;
          _nw147[(new BigNumber(8)).toNumber()] = (_867_v55).IsDisjointFrom((_868_v56)[_module.__default.safeIndex(_788_v1, new BigNumber((_868_v56).length))]);
          _nw147[(new BigNumber(9)).toNumber()] = !(!(false));
          _nw147[(new BigNumber(10)).toNumber()] = (_this).f21;
          _nw147[(new BigNumber(11)).toNumber()] = p0;
          _nw147[(new BigNumber(12)).toNumber()] = p0;
          _nw147[(new BigNumber(13)).toNumber()] = (_this).f21;
          _nw147[(new BigNumber(14)).toNumber()] = true;
          _nw147[(new BigNumber(15)).toNumber()] = p0;
          _nw147[(new BigNumber(16)).toNumber()] = p0;
          _nw147[(new BigNumber(17)).toNumber()] = !((_this).f21);
          _nw147[(new BigNumber(18)).toNumber()] = _module.__default.fm0((_this).f21, _789_v2, p0, globalState);
          _nw147[(new BigNumber(19)).toNumber()] = true;
          _869_v57 = _nw147;
          let _index105 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_869_v57).length));
          (_869_v57)[_index105] = p0;
          let _index106 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_869_v57).length));
          (_869_v57)[_index106] = ((_788_v1).plus(new BigNumber(-440))).isLessThanOrEqualTo((_dafny.ZERO).minus(_788_v1));
          let _870_v58;
          _870_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,p0);
          let _871_v59;
          _871_v59 = _dafny.Map.Empty.slice().updateUnsafe(true,(_870_v58).Merge((_870_v58).update((_869_v57)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_869_v57).length))], (_this).f21)));
          _871_v59 = (_871_v59).update((_this).f21, (_module.__default.fm22(globalState)).update(p0, (_this).f21));
          let _872_v60;
          _872_v60 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,new BigNumber((_861_v49).length));
          let _index107 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_869_v57).length));
          (_869_v57)[_index107] = !(!(_872_v60).contains(_788_v1));
        }
        (globalState).f15 = _module.__default.fm0(!(_788_v1).isEqualTo(_788_v1), _789_v2, (_this).f21, globalState);
        let _873_v61;
        _873_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_788_v1);
        _873_v61 = (_873_v61).update(_dafny.Seq.Concat(p1, p1), (_788_v1).plus(_788_v1));
        let _874_v62;
        _874_v62 = _dafny.MultiSet.fromElements(_788_v1);
        let _875_v63;
        _875_v63 = _dafny.Set.fromElements(_874_v62, _dafny.MultiSet.FromArray(_791_v4), _874_v62, _874_v62);
        let _876_v64;
        _876_v64 = _dafny.Seq.of(_dafny.Set.fromElements(_874_v62, _874_v62), _875_v63, _875_v63, _875_v63, _875_v63);
        let _877_v65;
        _877_v65 = _dafny.MultiSet.fromElements(((_876_v64)[_module.__default.safeIndex(_788_v1, new BigNumber((_876_v64).length))]).IsProperSubsetOf(_875_v63), (_this).f21, p0, (_this).f21);
        _877_v65 = _877_v65;
      } else {
        if (!((_module.__default.fm9(_788_v1, _788_v1, (_this).f21, globalState)).minus(_788_v1)).isEqualTo((((_this).f21) ? (_788_v1) : (_788_v1)))) {
          let _878_v66;
          _878_v66 = _dafny.Seq.UnicodeFromString("upxiueqmn");
          let _879_v67;
          _879_v67 = _dafny.Map.Empty.slice().updateUnsafe(_792_v5,_878_v66);
          _879_v67 = (_879_v67).update(_792_v5, _dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), ((_880_v2) => function (_881_i6) {
            return _880_v2;
          })(_789_v2)));
          let _882_v68;
          _882_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_788_v1);
          let _883_v69;
          _883_v69 = _dafny.MultiSet.fromElements(_788_v1);
          let _884_v70;
          _884_v70 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_788_v1), _dafny.MultiSet.fromElements(new BigNumber((_878_v66).length)), _883_v69, _883_v69, _dafny.MultiSet.FromArray(_791_v4));
          _882_v68 = (_882_v68).update((_module.__default.fm23((_this).f21, _882_v68, _module.__default.fm9(_788_v1, _788_v1, p0, globalState), _789_v2, globalState)).IsProperSubsetOf((_884_v70)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_883_v69).cardinality())), new BigNumber((_884_v70).length))]), _788_v1);
          (globalState).f11 = _module.__default.fm0((_this).f21, new _dafny.CodePoint('x'.codePointAt(0)), false, globalState);
          (globalState).f11 = p0;
          let _885_v71;
          let _init27 = ((_886_v1, _887_v4) => function (_888_i7) {
            return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(_886_v1), _887_v4)).length),_886_v1);
          })(_788_v1, _791_v4);
          let _nw148 = Array((new BigNumber(16)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw148.length); _i0_27++) {
            _nw148[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _885_v71 = _nw148;
          let _889_v72;
          _889_v72 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,new BigNumber((_883_v69).cardinality()));
          let _index108 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_885_v71).length));
          (_885_v71)[_index108] = _889_v72;
          let _index109 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_885_v71).length));
          (_885_v71)[_index109] = _889_v72;
        } else {
          let _890_v73;
          let _nw149 = Array((new BigNumber(22)).toNumber()).fill(false);
          _890_v73 = _nw149;
          let _891_v74;
          _891_v74 = _dafny.Set.fromElements(_890_v73);
          let _892_v75;
          _892_v75 = _dafny.Seq.of(_890_v73, _890_v73);
          let _893_v76;
          let _nw150 = Array((new BigNumber(16)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _891_v74;
          _nw150[(_dafny.ONE).toNumber()] = (_891_v74).Intersect(_891_v74);
          _nw150[(new BigNumber(2)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(3)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(4)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(5)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(_890_v73, _890_v73, _890_v73, _890_v73, (_892_v75)[_module.__default.safeIndex(_788_v1, new BigNumber((_892_v75).length))]);
          _nw150[(new BigNumber(7)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(8)).toNumber()] = (_891_v74).Difference(_891_v74);
          _nw150[(new BigNumber(9)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(10)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_890_v73);
          _nw150[(new BigNumber(12)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(13)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(14)).toNumber()] = _891_v74;
          _nw150[(new BigNumber(15)).toNumber()] = _dafny.Set.fromElements(_890_v73);
          _893_v76 = _nw150;
          let _index110 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_893_v76).length));
          (_893_v76)[_index110] = _891_v74;
          let _index111 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_893_v76).length));
          (_893_v76)[_index111] = _891_v74;
          _788_v1 = _788_v1;
          let _894_v77;
          let _nw151 = new _module.C0();
          _nw151.__ctor(false);
          _894_v77 = _nw151;
          let _index112 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_890_v73).length));
          (_890_v73)[_index112] = (_this).f21;
          let _index113 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_890_v73).length));
          (_890_v73)[_index113] = true;
          let _895_v78;
          let _nw152 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _895_v78 = _nw152;
          let _896_v79;
          _896_v79 = _dafny.MultiSet.fromElements(_788_v1, _788_v1);
          let _897_v80;
          _897_v80 = _dafny.Seq.of(_896_v79, _896_v79);
          let _898_v81;
          _898_v81 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,p0);
          let _899_v82;
          _899_v82 = _dafny.Set.fromElements(_788_v1, new BigNumber(((_896_v79).update((_791_v4)[_module.__default.safeIndex(new BigNumber((_791_v4).length), new BigNumber((_791_v4).length))], _module.__default.abs(_788_v1))).cardinality()), new BigNumber(((_896_v79).Intersect((_897_v80)[_module.__default.safeIndex(new BigNumber((_898_v81).length), new BigNumber((_897_v80).length))])).cardinality()));
          let _900_v83;
          _900_v83 = _dafny.MultiSet.fromElements(p1, p1, p1);
          let _901_v84;
          _901_v84 = _dafny.Map.Empty.slice().updateUnsafe(_900_v83,_788_v1);
          let _902_v85;
          _902_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_788_v1);
          let _index114 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_890_v73).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_890_v73).length));
          let _rhs98 = p0;
          let _rhs99 = _module.__default.fm0(!((_894_v77).fm15((_894_v77).f23, p0, (((_901_v84).contains((_900_v83).update(_dafny.Seq.of(p0), _module.__default.abs(_788_v1)))) ? ((_901_v84).get((_900_v83).update(_dafny.Seq.of(p0), _module.__default.abs(_788_v1)))) : (_788_v1)), _module.__default.fm24((_894_v77).fm15(p0, p0, new BigNumber((_791_v4).length), _902_v85, globalState), new BigNumber((_899_v82).length), globalState), globalState)), (_792_v5).dtor_cf26, true, globalState);
          let _rhs100 = _895_v78;
          let _rhs101 = (_899_v82).Union((function () {
            let _coll30 = new _dafny.Set();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(961), new BigNumber(442))) {
              let _903_v86 = _compr_30;
              if (((new BigNumber(961)).isLessThanOrEqualTo(_903_v86)) && ((_903_v86).isLessThan(new BigNumber(442)))) {
                _coll30.add((_903_v86).multipliedBy(_788_v1));
              }
            }
            return _coll30;
          }()).Difference(_dafny.Set.fromElements(_788_v1, new BigNumber(282), _788_v1)));
          let _lhs52 = _890_v73;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(427), new BigNumber((_890_v73).length));
          let _lhs54 = _890_v73;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_890_v73).length));
          _lhs52[_lhs53] = _rhs98;
          _lhs54[_lhs55] = _rhs99;
          _895_v78 = _rhs100;
          _899_v82 = _rhs101;
          let _904_v87;
          _904_v87 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_789_v2),_788_v1);
          let _905_v88;
          _905_v88 = _dafny.Seq.of(_789_v2, _789_v2);
          _904_v87 = (_904_v87).update(_905_v88, _788_v1);
        }
        let _906_v89;
        _906_v89 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,(_this).f21);
        let _907_v90;
        _907_v90 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_788_v1,(_this).f21));
        let _908_v91;
        _908_v91 = _dafny.Set.fromElements(_906_v89, (_907_v90)[_module.__default.safeIndex(_788_v1, new BigNumber((_907_v90).length))]);
        let _909_v92;
        _909_v92 = _module.D3.create_DC9((_908_v91).Difference(_908_v91));
        _909_v92 = _909_v92;
        let _910_v93;
        let _nw153 = new _module.C0();
        _nw153.__ctor((_this).f21);
        _910_v93 = _nw153;
        let _911_v94;
        _911_v94 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(_788_v1));
        let _912_v95;
        let _nw154 = Array((new BigNumber(12)).toNumber());
        _nw154[(_dafny.ZERO).toNumber()] = (_this).f21;
        _nw154[(_dafny.ONE).toNumber()] = p0;
        _nw154[(new BigNumber(2)).toNumber()] = p0;
        _nw154[(new BigNumber(3)).toNumber()] = ((_910_v93).fm14((_dafny.ZERO).minus(_788_v1), globalState)).isEqualTo(_788_v1);
        _nw154[(new BigNumber(4)).toNumber()] = p0;
        _nw154[(new BigNumber(5)).toNumber()] = (_this).f21;
        _nw154[(new BigNumber(6)).toNumber()] = p0;
        _nw154[(new BigNumber(7)).toNumber()] = false;
        _nw154[(new BigNumber(8)).toNumber()] = (_this).f21;
        _nw154[(new BigNumber(9)).toNumber()] = (_this).f21;
        _nw154[(new BigNumber(10)).toNumber()] = (_910_v93).f23;
        _nw154[(new BigNumber(11)).toNumber()] = (_910_v93).fm15(!((_910_v93).f23), !((_910_v93).f23), (_910_v93).fm14(_788_v1, globalState), _911_v94, globalState);
        _912_v95 = _nw154;
        _912_v95 = (((_this).f21) ? (_912_v95) : (((p0) ? (_912_v95) : (_912_v95))));
        let _source11 = _this.f20;
        if (_source11.is_DC1) {
          let _913___mcc_h7 = (_source11).cf1;
          let _914_cf1 = _913___mcc_h7;
          let _index116 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_912_v95).length));
          (_912_v95)[_index116] = p0;
          let _index117 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_912_v95).length));
          (_912_v95)[_index117] = (_910_v93).fm15((_910_v93).f23, ((_this).f21) || ((_910_v93).f23), new BigNumber(395), _911_v94, globalState);
          let _915_v96;
          _915_v96 = _dafny.Seq.UnicodeFromString("aln");
          let _916_v97;
          _916_v97 = _dafny.Map.Empty.slice().updateUnsafe(_788_v1,_915_v96);
          (globalState).f15 = _dafny.Seq.IsPrefixOf(_915_v96, (((_916_v97).contains(_788_v1)) ? ((_916_v97).get(_788_v1)) : (_915_v96)));
          let _917_v98;
          let _nw155 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _917_v98 = _nw155;
          _917_v98 = _917_v98;
          let _918_v99;
          let _nw156 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _918_v99 = _nw156;
          let _index118 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_918_v99).length));
          (_918_v99)[_index118] = _788_v1;
          let _919_v100;
          _919_v100 = _dafny.MultiSet.fromElements(_788_v1, _788_v1, _788_v1);
          let _index119 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_918_v99).length));
          let _rhs102 = _module.__default.safeModuloInt(_788_v1, _788_v1);
          let _rhs103 = _788_v1;
          let _rhs104 = (_919_v100).Difference(_919_v100);
          let _lhs56 = _918_v99;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_918_v99).length));
          _lhs56[_lhs57] = _rhs102;
          _788_v1 = _rhs103;
          _919_v100 = _rhs104;
        } else if (_source11.is_DC0) {
          let _920___mcc_h8 = (_source11).cf0;
          let _921_cf0 = _920___mcc_h8;
          let _922_v101;
          let _nw157 = Array((new BigNumber(25)).toNumber()).fill([]);
          _922_v101 = _nw157;
          let _923_v102;
          _923_v102 = _dafny.Seq.UnicodeFromString("uajufq");
          let _924_v103;
          let _nw158 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _924_v103 = _nw158;
          let _925_v104;
          _925_v104 = _dafny.Map.Empty.slice().updateUnsafe(_923_v102,_924_v103);
          let _index120 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_922_v101).length));
          (_922_v101)[_index120] = (((_925_v104).contains(_923_v102)) ? ((_925_v104).get(_923_v102)) : (_924_v103));
          let _index121 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_922_v101).length));
          (_922_v101)[_index121] = _924_v103;
          (globalState).f17 = (_792_v5).dtor_cf23;
          (globalState).f15 = (_910_v93).f23;
          _788_v1 = _788_v1;
        } else {
          let _926___mcc_h9 = (_source11).cf2;
          let _927_cf2 = _926___mcc_h9;
          r0 = !(((_this).f21) && (!((_910_v93).f23)));
          let _928_v105;
          let _nw159 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _928_v105 = _nw159;
          let _929_v106;
          _929_v106 = _module.D3.create_DC11(_911_v94, _788_v1, _928_v105);
          let _930_v107;
          _930_v107 = _dafny.MultiSet.fromElements(_module.D3.create_DC11(_911_v94, _788_v1, _928_v105), _929_v106, _929_v106, _929_v106);
          let _index122 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_912_v95).length));
          (_912_v95)[_index122] = (_930_v107).IsProperSubsetOf(_930_v107);
          let _931_v108;
          _931_v108 = _dafny.Seq.UnicodeFromString("vpii");
          let _932_v109;
          _932_v109 = _dafny.Set.fromElements((_910_v93).f23, _module.__default.fm0(p0, _789_v2, (_910_v93).f23, globalState));
          let _933_v110;
          _933_v110 = _dafny.Map.Empty.slice().updateUnsafe(_931_v108,_932_v109);
          let _index123 = _module.__default.safeIndex(new BigNumber(2), new BigNumber((_912_v95).length));
          (_912_v95)[_index123] = ((_933_v110).Merge(_933_v110)).equals(_933_v110);
          let _934_v111;
          _934_v111 = _module.D1.create_DC5(_931_v108, _788_v1, (_dafny.ZERO).minus(_module.__default.fm9(_788_v1, _788_v1, p0, globalState)));
          let _935_v113;
          _935_v113 = _dafny.MultiSet.fromElements(new BigNumber(740));
          let _936_v114;
          _936_v114 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_788_v1, new BigNumber((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_935_v113).Elements) {
              let _937_v112 = _compr_31;
              if ((_935_v113).contains(_937_v112)) {
                _coll31.push([(_937_v112).multipliedBy(_788_v1),p0]);
              }
            }
            return _coll31;
          }()).length)), _788_v1);
          let _rhs105 = _934_v111;
          let _rhs106 = _788_v1;
          let _rhs107 = _935_v113;
          let _rhs108 = ((true) ? (new BigNumber(446)) : (new BigNumber((_935_v113).cardinality())));
          _934_v111 = _rhs105;
          _788_v1 = _rhs106;
          _935_v113 = _rhs107;
          _788_v1 = _rhs108;
          let _rhs109 = _789_v2;
          let _rhs110 = (_dafny.ZERO).minus(_788_v1);
          let _rhs111 = _788_v1;
          let _rhs112 = _789_v2;
          let _rhs113 = _928_v105;
          _789_v2 = _rhs109;
          _788_v1 = _rhs110;
          _788_v1 = _rhs111;
          _789_v2 = _rhs112;
          _928_v105 = _rhs113;
        }
      }
      r0 = !(!((_this).f21));
      return r0;
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _938_i0;
      _938_i0 = _dafny.ZERO;
      L4: {
        while (p2) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_938_i0)) {
              break L4;
            }
            _938_i0 = (_938_i0).plus(_dafny.ONE);
            if ((_this).f21) {
              let _939_v0;
              _939_v0 = _module.D3.create_DC10(p0);
              let _940_v1;
              _940_v1 = _dafny.Seq.of(false, (_this).f21);
              let _941_v2;
              _941_v2 = _dafny.Map.Empty.slice().updateUnsafe(p2,_940_v1);
              let _942_v3;
              _942_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_941_v2);
              let _943_v4;
              _943_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm9((_939_v0).dtor_cf17, p0, p2, globalState),(((_942_v3).contains(p0)) ? ((_942_v3).get(p0)) : (_941_v2)));
              let _944_v5;
              _944_v5 = _dafny.Seq.of(p0);
              let _945_v6;
              _945_v6 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(154))).length));
              let _946_v7;
              _946_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_945_v6).length));
              let _947_v8;
              _947_v8 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
              let _948_v9;
              _948_v9 = _module.D1.create_DC5(p1, p0, new BigNumber((p1).length));
              let _949_v10;
              let _nw160 = Array((new BigNumber(22)).toNumber());
              _nw160[(_dafny.ZERO).toNumber()] = p0;
              _nw160[(_dafny.ONE).toNumber()] = p0;
              _nw160[(new BigNumber(2)).toNumber()] = p0;
              _nw160[(new BigNumber(3)).toNumber()] = new BigNumber(((((_942_v3).contains(new BigNumber((p1).length))) ? ((_942_v3).get(new BigNumber((p1).length))) : (_941_v2))).length);
              _nw160[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements((_this).f21)).cardinality());
              _nw160[(new BigNumber(5)).toNumber()] = p0;
              _nw160[(new BigNumber(6)).toNumber()] = p0;
              _nw160[(new BigNumber(7)).toNumber()] = p0;
              _nw160[(new BigNumber(8)).toNumber()] = new BigNumber((p1).length);
              _nw160[(new BigNumber(9)).toNumber()] = p0;
              _nw160[(new BigNumber(10)).toNumber()] = p0;
              _nw160[(new BigNumber(11)).toNumber()] = p0;
              _nw160[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_944_v5)[_module.__default.safeIndex(_module.__default.fm9(new BigNumber((_946_v7).length), p0, (_this).f21, globalState), new BigNumber((_944_v5).length))]);
              _nw160[(new BigNumber(13)).toNumber()] = p0;
              _nw160[(new BigNumber(14)).toNumber()] = new BigNumber((_940_v1).length);
              _nw160[(new BigNumber(15)).toNumber()] = p0;
              _nw160[(new BigNumber(16)).toNumber()] = new BigNumber((_947_v8).length);
              _nw160[(new BigNumber(17)).toNumber()] = p0;
              _nw160[(new BigNumber(18)).toNumber()] = (_948_v9).dtor_cf9;
              _nw160[(new BigNumber(19)).toNumber()] = p0;
              _nw160[(new BigNumber(20)).toNumber()] = p0;
              _nw160[(new BigNumber(21)).toNumber()] = new BigNumber(659);
              _949_v10 = _nw160;
              let _950_v11;
              let _nw161 = new _module.C1();
              _nw161.__ctor(_949_v10);
              _950_v11 = _nw161;
              let _951_v12;
              let _nw162 = Array((new BigNumber(18)).toNumber()).fill(_module.D3.Default());
              _951_v12 = _nw162;
              let _952_v13;
              let _init28 = ((_953_p2) => function (_954_i1) {
                return _953_p2;
              })(p2);
              let _nw163 = Array((new BigNumber(14)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw163.length); _i0_28++) {
                _nw163[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _952_v13 = _nw163;
              let _index124 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_952_v13).length));
              (_952_v13)[_index124] = (_this).f21;
              let _index125 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_952_v13).length));
              let _rhs114 = _951_v12;
              let _rhs115 = (_this).f21;
              let _lhs58 = _952_v13;
              let _lhs59 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_952_v13).length));
              _951_v12 = _rhs114;
              _lhs58[_lhs59] = _rhs115;
              let _955_v14;
              let _nw164 = new _module.C1();
              _nw164.__ctor((_950_v11).f22);
              _955_v14 = _nw164;
              _952_v13 = _952_v13;
              r2 = p2;
            } else {
              let _956_v15;
              _956_v15 = new BigNumber(714);
              let _957_v16;
              _957_v16 = _dafny.Seq.of(!(p2));
              let _958_v17;
              _958_v17 = _dafny.Seq.of(_957_v16);
              let _959_v18;
              _959_v18 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), ((_960_v16) => function (_961_i2) {
                return _960_v16;
              })(_957_v16)), _958_v17);
              let _962_v19;
              _962_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), function (_963_i3) {
                return new _dafny.CodePoint('r'.codePointAt(0));
              }));
              let _964_v20;
              _964_v20 = new _dafny.CodePoint('c'.codePointAt(0));
              let _965_v21;
              _965_v21 = _dafny.MultiSet.fromElements(_956_v15);
              let _966_v22;
              _966_v22 = _dafny.Seq.of(p0, _956_v15);
              let _967_v23;
              _967_v23 = _dafny.Map.Empty.slice().updateUnsafe(_956_v15,new BigNumber(300));
              let _968_v24;
              _968_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_967_v23).length),_956_v15);
              let _969_v25;
              let _nw165 = Array((new BigNumber(22)).toNumber());
              _nw165[(_dafny.ZERO).toNumber()] = p0;
              _nw165[(_dafny.ONE).toNumber()] = p0;
              _nw165[(new BigNumber(2)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(3)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(4)).toNumber()] = new BigNumber((_module.__default.fm16(_module.__default.fm0(_module.__default.fm0(p2, _964_v20, false, globalState), _964_v20, p2, globalState), p0, globalState)).length);
              _nw165[(new BigNumber(5)).toNumber()] = _module.__default.fm9(p0, new BigNumber((_957_v16).length), p2, globalState);
              _nw165[(new BigNumber(6)).toNumber()] = p0;
              _nw165[(new BigNumber(7)).toNumber()] = new BigNumber((_965_v21).cardinality());
              _nw165[(new BigNumber(8)).toNumber()] = p0;
              _nw165[(new BigNumber(9)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(10)).toNumber()] = new BigNumber(418);
              _nw165[(new BigNumber(11)).toNumber()] = p0;
              _nw165[(new BigNumber(12)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(13)).toNumber()] = p0;
              _nw165[(new BigNumber(14)).toNumber()] = new BigNumber((_966_v22).length);
              _nw165[(new BigNumber(15)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(16)).toNumber()] = new BigNumber(594);
              _nw165[(new BigNumber(17)).toNumber()] = new BigNumber((_968_v24).length);
              _nw165[(new BigNumber(18)).toNumber()] = p0;
              _nw165[(new BigNumber(19)).toNumber()] = new BigNumber(284);
              _nw165[(new BigNumber(20)).toNumber()] = _956_v15;
              _nw165[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f21,p0)).length);
              _969_v25 = _nw165;
              let _970_v26;
              _970_v26 = _dafny.Seq.of(p1);
              let _971_v27;
              _971_v27 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_972_v20) => function (_973_i4) {
                return _972_v20;
              })(_964_v20))).length));
              let _974_v28;
              let _nw166 = Array((new BigNumber(22)).toNumber());
              _nw166[(_dafny.ZERO).toNumber()] = new BigNumber((_957_v16).length);
              _nw166[(_dafny.ONE).toNumber()] = _956_v15;
              _nw166[(new BigNumber(2)).toNumber()] = p0;
              _nw166[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p0);
              _nw166[(new BigNumber(4)).toNumber()] = p0;
              _nw166[(new BigNumber(5)).toNumber()] = p0;
              _nw166[(new BigNumber(6)).toNumber()] = p0;
              _nw166[(new BigNumber(7)).toNumber()] = _956_v15;
              _nw166[(new BigNumber(8)).toNumber()] = p0;
              _nw166[(new BigNumber(9)).toNumber()] = new BigNumber(((_959_v18)[_module.__default.safeIndex(_956_v15, new BigNumber((_959_v18).length))]).length);
              _nw166[(new BigNumber(10)).toNumber()] = new BigNumber((_962_v19).length);
              _nw166[(new BigNumber(11)).toNumber()] = p0;
              _nw166[(new BigNumber(12)).toNumber()] = _956_v15;
              _nw166[(new BigNumber(13)).toNumber()] = new BigNumber((p1).length);
              _nw166[(new BigNumber(14)).toNumber()] = _956_v15;
              _nw166[(new BigNumber(15)).toNumber()] = _module.__default.fm9(new BigNumber(83), _956_v15, (_this).f21, globalState);
              _nw166[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(_969_v25, _969_v25)).length);
              _nw166[(new BigNumber(17)).toNumber()] = new BigNumber(((_970_v26)[_module.__default.safeIndex(p0, new BigNumber((_970_v26).length))]).length);
              _nw166[(new BigNumber(18)).toNumber()] = _956_v15;
              _nw166[(new BigNumber(19)).toNumber()] = new BigNumber((_957_v16).length);
              _nw166[(new BigNumber(20)).toNumber()] = new BigNumber((_971_v27).length);
              _nw166[(new BigNumber(21)).toNumber()] = new BigNumber((p1).length);
              _974_v28 = _nw166;
              _956_v15 = (_module.D3.create_DC11(_dafny.Map.Empty.slice().updateUnsafe((_this).f21,_956_v15), _956_v15, _969_v25)).dtor_cf19;
              let _975_v29;
              _975_v29 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_956_v15);
              _956_v15 = new BigNumber((_dafny.Seq.of(_966_v22, _dafny.Seq.of(_956_v15, _956_v15), _dafny.Seq.of(new BigNumber(544), new BigNumber((_975_v29).length)))).length);
              (globalState).f15 = (_this).f21;
              _956_v15 = _module.__default.safeDivisionInt((_966_v22)[_module.__default.safeIndex(_956_v15, new BigNumber((_966_v22).length))], p0);
              let _976_v30;
              _976_v30 = _dafny.Set.fromElements((_this).f21);
              _956_v15 = (new BigNumber((_976_v30).length)).multipliedBy(_956_v15);
            }
            let _977_v31;
            let _nw167 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _977_v31 = _nw167;
            let _978_v32;
            _978_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,_977_v31);
            let _979_v33;
            _979_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(p1, p1),_978_v32);
            _979_v33 = (_979_v33).update(_dafny.Seq.UnicodeFromString("ubtj"), _978_v32);
            let _980_v34;
            let _init29 = ((_981_p0) => function (_982_i5) {
              return _module.__default.safeModuloInt(_982_i5, _981_p0);
            })(p0);
            let _nw168 = Array((new BigNumber(28)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw168.length); _i0_29++) {
              _nw168[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _980_v34 = _nw168;
            let _983_v35;
            _983_v35 = _dafny.Seq.of(_980_v34);
            let _984_v36;
            let _nw169 = new _module.C1();
            _nw169.__ctor(_980_v34);
            _984_v36 = _nw169;
            let _985_v37;
            _985_v37 = _module.D5.create_DC16(_984_v36);
            let _986_v38;
            _986_v38 = _dafny.Seq.of((_985_v37).dtor_cf28);
            let _987_v39;
            let _nw170 = new _module.C1();
            _nw170.__ctor((_983_v35)[_module.__default.safeIndex(new BigNumber((_986_v38).length), new BigNumber((_983_v35).length))]);
            _987_v39 = _nw170;
            let _988_v40;
            _988_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_980_v34);
            _988_v40 = (_988_v40).update(((_this).f21) && (p2), (_987_v39).f22);
          }
        }
      }
      let _989_i6;
      _989_i6 = _dafny.ZERO;
      L5: {
        while (p2) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_989_i6)) {
              break L5;
            }
            _989_i6 = (_989_i6).plus(_dafny.ONE);
            let _990_v41;
            let _nw171 = Array((new BigNumber(23)).toNumber()).fill(false);
            _990_v41 = _nw171;
            let _991_v42;
            _991_v42 = _module.D1.create_DC4(new BigNumber(451), !((_this).f21), _990_v41);
            let _source12 = _991_v42;
            if (_source12.is_DC4) {
              let _992___mcc_h0 = (_source12).cf4;
              let _993___mcc_h1 = (_source12).cf5;
              let _994___mcc_h2 = (_source12).cf6;
              let _995_cf6 = _994___mcc_h2;
              let _996_cf5 = _993___mcc_h1;
              let _997_cf4 = _992___mcc_h0;
              r2 = p2;
              let _998_v43;
              let _nw172 = new _module.C0();
              _nw172.__ctor(p2);
              _998_v43 = _nw172;
              let _999_v44;
              let _nw173 = new _module.C0();
              _nw173.__ctor(!((_this).f21));
              _999_v44 = _nw173;
              let _1000_v45;
              _1000_v45 = new _dafny.CodePoint('b'.codePointAt(0));
              let _1001_v46;
              _1001_v46 = _dafny.Map.Empty.slice().updateUnsafe((((_999_v44).f23) ? (p1) : (_dafny.Seq.update(p1, _module.__default.safeIndex((_dafny.ZERO).minus(_997_cf4), new BigNumber((p1).length)), _1000_v45))),new BigNumber((_dafny.Set.fromElements((_this).f21, _996_cf5, (_999_v44).f23, (_this).f21)).length));
              let _1002_v47;
              _1002_v47 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,(_998_v43).f23);
              _1001_v46 = _module.__default.fm25((_998_v43).f23, (_1002_v47).equals(_1002_v47), globalState);
            } else if (_source12.is_DC5) {
              let _1003___mcc_h3 = (_source12).cf7;
              let _1004___mcc_h4 = (_source12).cf8;
              let _1005___mcc_h5 = (_source12).cf9;
              let _1006_cf9 = _1005___mcc_h5;
              let _1007_cf8 = _1004___mcc_h4;
              let _1008_cf7 = _1003___mcc_h3;
              let _1009_v48;
              _1009_v48 = _dafny.Seq.of(new BigNumber((p1).length));
              let _1010_v49;
              let _nw174 = Array((new BigNumber(17)).toNumber());
              _nw174[(_dafny.ZERO).toNumber()] = _1007_cf8;
              _nw174[(_dafny.ONE).toNumber()] = (p0).multipliedBy(_1006_cf9);
              _nw174[(new BigNumber(2)).toNumber()] = _1007_cf8;
              _nw174[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_1006_cf9, p0);
              _nw174[(new BigNumber(4)).toNumber()] = _1006_cf9;
              _nw174[(new BigNumber(5)).toNumber()] = p0;
              _nw174[(new BigNumber(6)).toNumber()] = p0;
              _nw174[(new BigNumber(7)).toNumber()] = _1007_cf8;
              _nw174[(new BigNumber(8)).toNumber()] = new BigNumber(723);
              _nw174[(new BigNumber(9)).toNumber()] = new BigNumber(785);
              _nw174[(new BigNumber(10)).toNumber()] = _1006_cf9;
              _nw174[(new BigNumber(11)).toNumber()] = p0;
              _nw174[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Set.fromElements((_1009_v48)[_module.__default.safeIndex(_1007_cf8, new BigNumber((_1009_v48).length))], _1007_cf8, _1006_cf9, p0)).length);
              _nw174[(new BigNumber(13)).toNumber()] = _1007_cf8;
              _nw174[(new BigNumber(14)).toNumber()] = (_1009_v48)[_module.__default.safeIndex(p0, new BigNumber((_1009_v48).length))];
              _nw174[(new BigNumber(15)).toNumber()] = p0;
              _nw174[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
              _1010_v49 = _nw174;
              let _index126 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1010_v49).length));
              (_1010_v49)[_index126] = _1006_cf9;
              let _1011_v50;
              _1011_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1010_v49,_1010_v49);
              let _index127 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1010_v49).length));
              (_1010_v49)[_index127] = new BigNumber((_1011_v50).length);
              (globalState).f11 = (_this).f21;
              let _1012_v51;
              _1012_v51 = _dafny.Seq.of((_this).f21);
              let _index128 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_990_v41).length));
              (_990_v41)[_index128] = (_1012_v51)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_this).f21, (_this).f21, (_1012_v51)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("xlt")).length), new BigNumber((_1012_v51).length))])).length), new BigNumber((_1012_v51).length))];
              let _1013_v52;
              let _nw175 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1013_v52 = _nw175;
              let _1014_v53;
              _1014_v53 = new _dafny.CodePoint('h'.codePointAt(0));
              let _index129 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1013_v52).length));
              (_1013_v52)[_index129] = _1014_v53;
              let _1015_v54;
              _1015_v54 = _dafny.Set.fromElements(!((_this).f21));
              let _1016_v55;
              let _nw176 = new _module.C0();
              _nw176.__ctor((_1015_v54).IsSubsetOf(_1015_v54));
              _1016_v55 = _nw176;
              let _index130 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_990_v41).length));
              let _index131 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1013_v52).length));
              let _rhs116 = p2;
              let _rhs117 = _1006_cf9;
              let _rhs118 = (_1015_v54).IsSubsetOf((_dafny.Set.fromElements((_this).f21)).Difference(_1015_v54));
              let _rhs119 = _1014_v53;
              let _rhs120 = _1016_v55;
              let _lhs60 = _990_v41;
              let _lhs61 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_990_v41).length));
              let _lhs62 = _1013_v52;
              let _lhs63 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_1013_v52).length));
              _lhs60[_lhs61] = _rhs116;
              _1006_cf9 = _rhs117;
              r0 = _rhs118;
              _lhs62[_lhs63] = _rhs119;
              _1016_v55 = _rhs120;
              let _index132 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1010_v49).length));
              (_1010_v49)[_index132] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_1006_cf9), _1007_cf8);
            } else if (_source12.is_DC6) {
              let _1017___mcc_h6 = (_source12).cf10;
              let _1018___mcc_h7 = (_source12).cf11;
              let _1019_cf11 = _1018___mcc_h7;
              let _1020_cf10 = _1017___mcc_h6;
              _1020_cf10 = _1019_cf11;
              let _1021_v56;
              _1021_v56 = _dafny.MultiSet.fromElements(p0, p0);
              let _1022_v57;
              _1022_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1021_v56,_1020_cf10);
              r1 = (_1022_v57).equals(_1022_v57);
              let _1023_v58;
              _1023_v58 = _dafny.Seq.of(_990_v41, _990_v41, _990_v41, _990_v41, _990_v41);
              r2 = _dafny.Seq.contains(_1023_v58, _990_v41);
              r3 = p1;
            } else {
              let _1024___mcc_h8 = (_source12).cf3;
              let _1025_cf3 = _1024___mcc_h8;
              let _1026_v59;
              _1026_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _1027_v60;
              _1027_v60 = _dafny.MultiSet.fromElements(_1026_v59);
              let _1028_v61;
              _1028_v61 = _module.D6.create_DC19(_dafny.MultiSet.fromElements(_1026_v59, _1026_v59));
              _1027_v60 = (_1027_v60).Difference((_1028_v61).dtor_cf35);
              let _1029_v62;
              let _nw177 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
              _1029_v62 = _nw177;
              let _1030_v63;
              _1030_v63 = _dafny.Set.fromElements(new BigNumber((p1).length), p0, p0);
              let _1031_v64;
              _1031_v64 = _dafny.Seq.of(new BigNumber((_1030_v63).length), p0);
              let _index133 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1029_v62).length));
              (_1029_v62)[_index133] = _1031_v64;
              let _index134 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_1029_v62).length));
              (_1029_v62)[_index134] = _1031_v64;
              let _1032_v65;
              let _nw178 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
              _1032_v65 = _nw178;
              let _index135 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1032_v65).length));
              (_1032_v65)[_index135] = new BigNumber((_dafny.Seq.Concat(p1, _module.__default.fm16((_this).f21, p0, globalState))).length);
              let _1033_v66;
              _1033_v66 = _module.D6.create_DC20((_this).f21, p0, true);
              let _index136 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1032_v65).length));
              let _rhs121 = p2;
              let _rhs122 = (_1033_v66).dtor_cf36;
              let _rhs123 = (_dafny.ZERO).minus(p0);
              let _lhs64 = globalState;
              let _lhs65 = globalState;
              let _lhs66 = _1032_v65;
              let _lhs67 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_1032_v65).length));
              _lhs64.f17 = _rhs121;
              _lhs65.f17 = _rhs122;
              _lhs66[_lhs67] = _rhs123;
              let _1034_v67;
              _1034_v67 = _dafny.Seq.of((((_this).f21) ? (p2) : ((_this).f21)), (_this).f21, (_this).f21, p2);
              (globalState).f16 = _dafny.Seq.update(_1034_v67, _module.__default.safeIndex(p0, new BigNumber((_1034_v67).length)), (_this).f21);
            }
            let _1035_v68;
            _1035_v68 = _dafny.Set.fromElements(!((_this).f21), p2, p2);
            let _1036_v69;
            _1036_v69 = _dafny.Seq.of(_module.__default.fm20(globalState));
            let _1037_v70;
            _1037_v70 = new BigNumber(616);
            let _rhs124 = (_1035_v68).Intersect(_dafny.Set.fromElements(true));
            let _rhs125 = _dafny.Seq.Concat(p1, p1);
            let _rhs126 = _dafny.Seq.Concat(_1036_v69, _dafny.Seq.of(_1035_v68, _1035_v68));
            let _rhs127 = new BigNumber(396);
            _1035_v68 = _rhs124;
            r3 = _rhs125;
            _1036_v69 = _rhs126;
            _1037_v70 = _rhs127;
            let _1038_v71;
            _1038_v71 = new _dafny.CodePoint('w'.codePointAt(0));
            let _1039_v72;
            _1039_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1037_v70,_module.__default.fm0((_this).f21, _1038_v71, p2, globalState));
            let _1040_v73;
            _1040_v73 = _dafny.MultiSet.fromElements((_this).f21);
            let _1041_v74;
            let _nw179 = Array((new BigNumber(27)).toNumber());
            _nw179[(_dafny.ZERO).toNumber()] = p0;
            _nw179[(_dafny.ONE).toNumber()] = p0;
            _nw179[(new BigNumber(2)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(3)).toNumber()] = p0;
            _nw179[(new BigNumber(4)).toNumber()] = new BigNumber((p1).length);
            _nw179[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw179[(new BigNumber(6)).toNumber()] = p0;
            _nw179[(new BigNumber(7)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(8)).toNumber()] = new BigNumber(443);
            _nw179[(new BigNumber(9)).toNumber()] = new BigNumber(-109);
            _nw179[(new BigNumber(10)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(11)).toNumber()] = new BigNumber(882);
            _nw179[(new BigNumber(12)).toNumber()] = new BigNumber(55);
            _nw179[(new BigNumber(13)).toNumber()] = p0;
            _nw179[(new BigNumber(14)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(15)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(16)).toNumber()] = _module.__default.fm9(p0, new BigNumber((_1035_v68).length), (_this).f21, globalState);
            _nw179[(new BigNumber(17)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(18)).toNumber()] = new BigNumber((_1039_v72).length);
            _nw179[(new BigNumber(19)).toNumber()] = p0;
            _nw179[(new BigNumber(20)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(21)).toNumber()] = _1037_v70;
            _nw179[(new BigNumber(22)).toNumber()] = new BigNumber((_1040_v73).cardinality());
            _nw179[(new BigNumber(23)).toNumber()] = p0;
            _nw179[(new BigNumber(24)).toNumber()] = p0;
            _nw179[(new BigNumber(25)).toNumber()] = p0;
            _nw179[(new BigNumber(26)).toNumber()] = _1037_v70;
            _1041_v74 = _nw179;
            let _1042_v75;
            let _nw180 = new _module.C1();
            _nw180.__ctor(_1041_v74);
            _1042_v75 = _nw180;
            let _1043_v76;
            let _nw181 = new _module.C0();
            _nw181.__ctor((_this).f21);
            _1043_v76 = _nw181;
            let _1044_v77;
            _1044_v77 = _module.D4.create_DC12(_1043_v76);
            let _1045_v78;
            _1045_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1040_v73).cardinality()),_1043_v76);
            let _1046_v79;
            _1046_v79 = _module.D2.create_DC8(_990_v41, p0, p0);
            let _1047_v80;
            _1047_v80 = _dafny.Map.Empty.slice().updateUnsafe(_1046_v79,_1043_v76);
            let _1048_v81;
            _1048_v81 = _dafny.Seq.of(_1037_v70);
            let _1049_v82;
            _1049_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1048_v81);
            let _1050_v83;
            let _nw182 = Array((new BigNumber(25)).toNumber());
            _nw182[(_dafny.ZERO).toNumber()] = _1043_v76;
            _nw182[(_dafny.ONE).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(2)).toNumber()] = (_1044_v77).dtor_cf21;
            _nw182[(new BigNumber(3)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(4)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(5)).toNumber()] = (_1044_v77).dtor_cf21;
            _nw182[(new BigNumber(6)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(7)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(8)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(9)).toNumber()] = (((_1045_v78).contains(p0)) ? ((_1045_v78).get(p0)) : (_1043_v76));
            _nw182[(new BigNumber(10)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(11)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(12)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(13)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(14)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(15)).toNumber()] = (_1044_v77).dtor_cf21;
            _nw182[(new BigNumber(16)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(17)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(18)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(19)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(20)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(21)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(22)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(23)).toNumber()] = _1043_v76;
            _nw182[(new BigNumber(24)).toNumber()] = (((_1047_v80).contains(_module.D2.create_DC8(_990_v41, new BigNumber((_1049_v82).length), (_module.__default.fm26(false, new _dafny.CodePoint('m'.codePointAt(0)), (_this).f21, p2, globalState)).dtor_cf37))) ? ((_1047_v80).get(_module.D2.create_DC8(_990_v41, new BigNumber((_1049_v82).length), (_module.__default.fm26(false, new _dafny.CodePoint('m'.codePointAt(0)), (_this).f21, p2, globalState)).dtor_cf37))) : (_1043_v76));
            _1050_v83 = _nw182;
            let _index137 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1050_v83).length));
            (_1050_v83)[_index137] = _1043_v76;
            let _1051_v84;
            _1051_v84 = _dafny.Seq.of(p2);
            let _1052_v85;
            _1052_v85 = _dafny.MultiSet.fromElements(p0, new BigNumber(60), p0);
            let _index138 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1050_v83).length));
            let _rhs128 = _1042_v75;
            let _rhs129 = (_1037_v70).isLessThan(_1037_v70);
            let _rhs130 = _1043_v76;
            let _rhs131 = !(((true) ? ((_dafny.MultiSet.fromElements(p0, new BigNumber(115), new BigNumber((_1051_v84).length))).IsSubsetOf(_1052_v85)) : ((_this).f21)));
            let _lhs68 = globalState;
            let _lhs69 = _1050_v83;
            let _lhs70 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_1050_v83).length));
            let _lhs71 = globalState;
            _1042_v75 = _rhs128;
            _lhs68.f11 = _rhs129;
            _lhs69[_lhs70] = _rhs130;
            _lhs71.f11 = _rhs131;
            let _1053_v86;
            let _nw183 = Array((new BigNumber(2)).toNumber());
            _nw183[(_dafny.ZERO).toNumber()] = _1038_v71;
            _nw183[(_dafny.ONE).toNumber()] = _1038_v71;
            _1053_v86 = _nw183;
            _1053_v86 = _1053_v86;
          }
        }
      }
      if (!(_dafny.areEqual(_dafny.Seq.Concat(p1, p1), p1))) {
        let _1054_v87;
        _1054_v87 = _dafny.Set.fromElements(p0, p0);
        let _1055_v88;
        _1055_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1054_v87);
        let _1056_v89;
        _1056_v89 = _dafny.Seq.of((_this).f21);
        let _1057_v90;
        let _out53;
        _out53 = (_this).m5((_1054_v87).equals((((_1055_v88).contains(false)) ? ((_1055_v88).get(false)) : (_1054_v87))), _1056_v89, globalState);
        _1057_v90 = _out53;
        (globalState).f11 = true;
        let _1058_v91;
        let _init30 = function (_1059_i7) {
          return (_1059_i7).multipliedBy(new BigNumber(-579));
        };
        let _nw184 = Array((new BigNumber(7)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw184.length); _i0_30++) {
          _nw184[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1058_v91 = _nw184;
        let _1060_v92;
        let _nw185 = new _module.C1();
        _nw185.__ctor(_1058_v91);
        _1060_v92 = _nw185;
        let _1061_v93;
        _1061_v93 = _dafny.MultiSet.fromElements(p0, p0, p0);
        let _1062_v94;
        _1062_v94 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1063_v95;
        _1063_v95 = new _dafny.CodePoint('h'.codePointAt(0));
        _1061_v93 = (_1061_v93).Intersect(_module.__default.fm23(_1057_v90, (_1062_v94).update(p2, p0), p0, _1063_v95, globalState));
        let _1064_v96;
        _1064_v96 = _dafny.Map.Empty.slice().updateUnsafe((((_1061_v93).contains(p0)) ? ((_1061_v93).get(p0)) : (p0)),p0);
        let _1065_v97;
        _1065_v97 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_module.__default.fm16(_1057_v90, p0, globalState))).length)), p0);
        _1064_v96 = (_1064_v96).update((_1065_v97)[_module.__default.safeIndex(p0, new BigNumber((_1065_v97).length))], _module.__default.safeModuloInt(p0, p0));
      } else {
        let _1066_v98;
        let _nw186 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1066_v98 = _nw186;
        let _index139 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length));
        (_1066_v98)[_index139] = p0;
        let _index140 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length));
        (_1066_v98)[_index140] = p0;
        (globalState).f15 = (_this).f21;
        let _index141 = _module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length));
        (_1066_v98)[_index141] = ((_dafny.ZERO).minus(p0)).plus((_1066_v98)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length))]);
        let _1067_v99;
        _1067_v99 = _module.D6.create_DC20((_this).f21, (_1066_v98)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length))], (_this).f21);
        (globalState).f17 = (_1067_v99).dtor_cf36;
        let _1068_v100;
        let _nw187 = Array((new BigNumber(15)).toNumber()).fill(false);
        _1068_v100 = _nw187;
        let _1069_v101;
        _1069_v101 = _dafny.Seq.of(_1068_v100, _1068_v100);
        let _1070_v102;
        _1070_v102 = _dafny.Set.fromElements(_1068_v100, _1068_v100);
        let _1071_v103;
        _1071_v103 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1072_v104;
        _1072_v104 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1071_v103).length),(_1066_v98)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length))]);
        let _1073_v106;
        _1073_v106 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1074_v107;
        _1074_v107 = _module.D4.create_DC14(false, p0, function () {
  let _coll32 = new _dafny.Set();
  for (const _compr_32 of (_1072_v104).Keys.Elements) {
    let _1075_v105 = _compr_32;
    if ((_1072_v104).contains(_1075_v105)) {
      _coll32.add((_1075_v105).plus(new BigNumber((_dafny.Seq.UnicodeFromString("lp")).length)));
    }
  }
  return _coll32;
}(), _1073_v106);
        let _1076_v108;
        _1076_v108 = _dafny.Map.Empty.slice().updateUnsafe((_1070_v102).IsSubsetOf(_dafny.Set.fromElements((_1069_v101)[_module.__default.safeIndex((_1066_v98)[_module.__default.safeIndex(new BigNumber(151), new BigNumber((_1066_v98).length))], new BigNumber((_1069_v101).length))])),_module.__default.fm12((_1074_v107).dtor_cf24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-558)), function (_1077_i8) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(464)), ((_1078_v106) => function (_1079_i9) {
          return _1078_v106;
        })(_1073_v106)), globalState));
        _1076_v108 = (_1076_v108).update((_this).f21, _1073_v106);
      }
      let _1080_i10;
      _1080_i10 = _dafny.ZERO;
      L6: {
        while (!(p0).isEqualTo(p0)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1080_i10)) {
              break L6;
            }
            _1080_i10 = (_1080_i10).plus(_dafny.ONE);
            let _1081_v109;
            _1081_v109 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f21);
            let _1082_v110;
            _1082_v110 = _dafny.Set.fromElements(_1081_v109);
            let _1083_v111;
            _1083_v111 = _module.D3.create_DC9(_1082_v110);
            _1083_v111 = _1083_v111;
            let _1084_v112;
            let _nw188 = Array((new BigNumber(8)).toNumber()).fill([]);
            _1084_v112 = _nw188;
            let _1085_v113;
            _1085_v113 = new _dafny.CodePoint('i'.codePointAt(0));
            let _1086_v114;
            _1086_v114 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),p1);
            let _1087_v115;
            _1087_v115 = _dafny.Seq.of(p0, new BigNumber((_dafny.Seq.of(p0)).length));
            let _1088_v116;
            _1088_v116 = _module.D3.create_DC10(p0);
            let _1089_v117;
            _1089_v117 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1087_v115).length),_1088_v116);
            let _1090_v118;
            let _nw189 = Array((new BigNumber(21)).toNumber());
            _nw189[(_dafny.ZERO).toNumber()] = p1;
            _nw189[(_dafny.ONE).toNumber()] = p1;
            _nw189[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("snqcwt");
            _nw189[(new BigNumber(3)).toNumber()] = p1;
            _nw189[(new BigNumber(4)).toNumber()] = p1;
            _nw189[(new BigNumber(5)).toNumber()] = p1;
            _nw189[(new BigNumber(6)).toNumber()] = p1;
            _nw189[(new BigNumber(7)).toNumber()] = p1;
            _nw189[(new BigNumber(8)).toNumber()] = p1;
            _nw189[(new BigNumber(9)).toNumber()] = p1;
            _nw189[(new BigNumber(10)).toNumber()] = p1;
            _nw189[(new BigNumber(11)).toNumber()] = p1;
            _nw189[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("cx");
            _nw189[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(p0, new BigNumber((p1).length)), _1085_v113);
            _nw189[(new BigNumber(14)).toNumber()] = p1;
            _nw189[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("ygblryv");
            _nw189[(new BigNumber(16)).toNumber()] = (((_1086_v114).contains(new BigNumber((_1089_v117).length))) ? ((_1086_v114).get(new BigNumber((_1089_v117).length))) : (p1));
            _nw189[(new BigNumber(17)).toNumber()] = p1;
            _nw189[(new BigNumber(18)).toNumber()] = p1;
            _nw189[(new BigNumber(19)).toNumber()] = p1;
            _nw189[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-367)), ((_1091_v113) => function (_1092_i11) {
              return _1091_v113;
            })(_1085_v113));
            _1090_v118 = _nw189;
            let _1093_v119;
            _1093_v119 = _1090_v118;
            let _1094_v120;
            _1094_v120 = _dafny.Seq.of((_1090_v118));
            let _index142 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1084_v112).length));
            (_1084_v112)[_index142] = (_1094_v120)[_module.__default.safeIndex(new BigNumber(215), new BigNumber((_1094_v120).length))];
            let _1095_v121;
            _1095_v121 = _dafny.MultiSet.fromElements(p2);
            let _index143 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1084_v112).length));
            let _rhs132 = _1090_v118;
            let _rhs133 = (_module.__default.safeModuloInt(p0, p0)).isLessThanOrEqualTo(new BigNumber(419));
            let _rhs134 = (_1095_v121).Difference((_dafny.MultiSet.fromElements(p2)).update(false, _module.__default.abs(p0)));
            let _lhs72 = _1084_v112;
            let _lhs73 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_1084_v112).length));
            let _lhs74 = globalState;
            _lhs72[_lhs73] = _rhs132;
            _lhs74.f15 = _rhs133;
            _1095_v121 = _rhs134;
            let _1096_v122;
            _1096_v122 = _dafny.Seq.of(p2);
            let _1097_v123;
            let _nw190 = Array((new BigNumber(13)).toNumber());
            _nw190[(_dafny.ZERO).toNumber()] = _1096_v122;
            _nw190[(_dafny.ONE).toNumber()] = _dafny.Seq.of(p2, p2);
            _nw190[(new BigNumber(2)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(3)).toNumber()] = ((p2) ? (_1096_v122) : (_1096_v122));
            _nw190[(new BigNumber(4)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1096_v122, _dafny.Seq.of(true, (_this).f21));
            _nw190[(new BigNumber(6)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(7)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(8)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(9)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(10)).toNumber()] = _1096_v122;
            _nw190[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1096_v122, _1096_v122);
            _nw190[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1096_v122, _1096_v122);
            _1097_v123 = _nw190;
            let _index144 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1097_v123).length));
            (_1097_v123)[_index144] = _dafny.Seq.Concat(_1096_v122, _module.__default.fm27(p0, p0, (_dafny.ZERO).minus(p0), globalState));
            let _index145 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_1097_v123).length));
            (_1097_v123)[_index145] = _dafny.Seq.update(_dafny.Seq.of(p2), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of(p2)).length)), p2);
            let _1098_v124;
            let _init31 = ((_1099_p0) => function (_1100_i12) {
              return _module.__default.safeModuloInt(_1100_i12, _1099_p0);
            })(p0);
            let _nw191 = Array((new BigNumber(7)).toNumber());
            for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw191.length); _i0_31++) {
              _nw191[_i0_31] = _init31(new BigNumber(_i0_31));
            }
            _1098_v124 = _nw191;
            let _1101_v125;
            _1101_v125 = _dafny.Seq.of(_1098_v124);
            _1101_v125 = _dafny.Seq.of(_1098_v124);
          }
        }
      }
      let _hi7 = p0;
      for (let _1102_i13 = p0; _1102_i13.isLessThan(_hi7); _1102_i13 = _1102_i13.plus(_dafny.ONE)) {
        let _1103_v126;
        let _init32 = function (_1104_i14) {
          return _dafny.Seq.of((_this).f21, (_this).f21);
        };
        let _nw192 = Array((new BigNumber(24)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw192.length); _i0_32++) {
          _nw192[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1103_v126 = _nw192;
        let _1105_v127;
        _1105_v127 = _dafny.Seq.of((_this).f21);
        let _index146 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1103_v126).length));
        (_1103_v126)[_index146] = _dafny.Seq.Concat(_1105_v127, _1105_v127);
        let _1106_v128;
        let _nw193 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1106_v128 = _nw193;
        let _1107_v129;
        _1107_v129 = new BigNumber(808);
        let _1108_v130;
        _1108_v130 = _module.D8.create_DC22(_1106_v128);
        let _index147 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1103_v126).length));
        let _rhs135 = !((_module.__default.fm9(p0, p0, false, globalState)).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber(-12), (_dafny.ZERO).minus(_1107_v129))));
        let _rhs136 = _1105_v127;
        let _rhs137 = (_1108_v130).dtor_cf40;
        let _rhs138 = (p0).minus(_1107_v129);
        let _rhs139 = _1102_i13;
        let _lhs75 = _1103_v126;
        let _lhs76 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_1103_v126).length));
        r1 = _rhs135;
        _lhs75[_lhs76] = _rhs136;
        _1106_v128 = _rhs137;
        _1107_v129 = _rhs138;
        _1107_v129 = _rhs139;
        let _1109_v131;
        _1109_v131 = _dafny.Map.Empty.slice().updateUnsafe((_this).f21,_1102_i13);
        let _1110_v133;
        let _init33 = ((_1111_v131, _1112_p0) => function (_1113_i15) {
          return (function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(-447), new BigNumber(915))) {
              let _1114_v132 = _compr_33;
              if (((new BigNumber(-447)).isLessThanOrEqualTo(_1114_v132)) && ((_1114_v132).isLessThan(new BigNumber(915)))) {
                _coll33.push([(_1114_v132).minus(_1112_p0),_1112_p0]);
              }
            }
            return _coll33;
          }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(515),new BigNumber((_1111_v131).length)));
        })(_1109_v131, p0);
        let _nw194 = Array((new BigNumber(15)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw194.length); _i0_33++) {
          _nw194[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1110_v133 = _nw194;
        let _1115_v134;
        _1115_v134 = _dafny.MultiSet.fromElements(p2);
        let _1116_v135;
        _1116_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1115_v134).cardinality()),new BigNumber((p1).length));
        let _index148 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_1110_v133).length));
        (_1110_v133)[_index148] = _1116_v135;
        let _index149 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_1110_v133).length));
        let _rhs140 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, _module.__default.fm9(new BigNumber(682), new BigNumber((p1).length), (_this).f21, globalState)));
        let _rhs141 = _1109_v131;
        let _rhs142 = _module.__default.fm9(new BigNumber(799), p0, (_dafny.MultiSet.fromElements(!(false), (_this).f21)).IsProperSubsetOf(_module.__default.fm28(_1102_i13, p2, globalState)), globalState);
        let _rhs143 = _module.__default.fm10(globalState);
        let _lhs77 = _1110_v133;
        let _lhs78 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_1110_v133).length));
        _1107_v129 = _rhs140;
        _1109_v131 = _rhs141;
        _1107_v129 = _rhs142;
        _lhs77[_lhs78] = _rhs143;
        let _1117_v136;
        _1117_v136 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_1118_i16) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).length),(_this).f21);
        let _1119_v137;
        _1119_v137 = _dafny.Seq.of(_1102_i13);
        let _1120_v138;
        _1120_v138 = _dafny.Set.fromElements((_this).f21, p2);
        let _1121_v139;
        _1121_v139 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1122_v140;
        _1122_v140 = _dafny.Map.Empty.slice().updateUnsafe(_1121_v139,_1121_v139);
        let _1123_v141;
        _1123_v141 = _dafny.MultiSet.fromElements(new BigNumber((_1122_v140).length), _1107_v129, _1102_i13, _1107_v129, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,_1107_v129)).length));
        let _1124_v142;
        let _nw195 = Array((new BigNumber(27)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = p0;
        _nw195[(_dafny.ONE).toNumber()] = p0;
        _nw195[(new BigNumber(2)).toNumber()] = _1102_i13;
        _nw195[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(_1107_v129, new BigNumber((p1).length)), new _dafny.CodePoint('t'.codePointAt(0)))).length);
        _nw195[(new BigNumber(4)).toNumber()] = new BigNumber(-593);
        _nw195[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw195[(new BigNumber(6)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(7)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(8)).toNumber()] = new BigNumber((_1117_v136).length);
        _nw195[(new BigNumber(9)).toNumber()] = new BigNumber((_1119_v137).length);
        _nw195[(new BigNumber(10)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(11)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(12)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(13)).toNumber()] = p0;
        _nw195[(new BigNumber(14)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(15)).toNumber()] = new BigNumber((_1119_v137).length);
        _nw195[(new BigNumber(16)).toNumber()] = _module.__default.fm9(_1107_v129, new BigNumber((p1).length), p2, globalState);
        _nw195[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(p1)).cardinality());
        _nw195[(new BigNumber(18)).toNumber()] = p0;
        _nw195[(new BigNumber(19)).toNumber()] = p0;
        _nw195[(new BigNumber(20)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(21)).toNumber()] = p0;
        _nw195[(new BigNumber(22)).toNumber()] = _1102_i13;
        _nw195[(new BigNumber(23)).toNumber()] = new BigNumber((_1120_v138).length);
        _nw195[(new BigNumber(24)).toNumber()] = (((_1123_v141).contains(new BigNumber((p1).length))) ? ((_1123_v141).get(new BigNumber((p1).length))) : (_1107_v129));
        _nw195[(new BigNumber(25)).toNumber()] = _1107_v129;
        _nw195[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f21)).length);
        _1124_v142 = _nw195;
        let _1125_v143;
        _1125_v143 = _dafny.MultiSet.fromElements(_1124_v142);
        let _1126_v144;
        _1126_v144 = _dafny.Seq.of(_1125_v143);
        _1107_v129 = new BigNumber(((_1126_v144)[_module.__default.safeIndex(p0, new BigNumber((_1126_v144).length))]).cardinality());
        let _1127_v145;
        let _nw196 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1127_v145 = _nw196;
        let _1128_v146;
        _1128_v146 = _dafny.Seq.of(_1127_v145, _1127_v145);
        let _1129_v147;
        let _nw197 = Array((new BigNumber(26)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = _1127_v145;
        _nw197[(_dafny.ONE).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(2)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(3)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(4)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(5)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(6)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(7)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(8)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(9)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(10)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(11)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(12)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(13)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(14)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(15)).toNumber()] = (_1128_v146)[_module.__default.safeIndex(p0, new BigNumber((_1128_v146).length))];
        _nw197[(new BigNumber(16)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(17)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(18)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(19)).toNumber()] = (((_this).f21) ? (_1127_v145) : (_1127_v145));
        _nw197[(new BigNumber(20)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(21)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(22)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(23)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(24)).toNumber()] = _1127_v145;
        _nw197[(new BigNumber(25)).toNumber()] = _1127_v145;
        _1129_v147 = _nw197;
        let _1130_v148;
        _1130_v148 = _dafny.Seq.of((_1103_v126)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1103_v126).length))], _dafny.Seq.Concat(_1105_v127, (_1103_v126)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_1103_v126).length))]));
        let _1131_v149;
        _1131_v149 = _dafny.Seq.of(_1124_v142);
        let _1132_v150;
        _1132_v150 = _dafny.Seq.of(_1124_v142, (_1131_v149)[_module.__default.safeIndex(_1107_v129, new BigNumber((_1131_v149).length))], _1124_v142, _1124_v142, _1124_v142);
        let _1133_v151;
        let _nw198 = new _module.C1();
        _nw198.__ctor((_1132_v150)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1132_v150).length))]);
        _1133_v151 = _nw198;
        let _1134_v152;
        _1134_v152 = _module.D5.create_DC16(_1133_v151);
        let _1135_v153;
        let _nw199 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1135_v153 = _nw199;
        let _rhs144 = _1129_v147;
        let _rhs145 = ((p2) ? (_1130_v148) : (_1130_v148));
        let _rhs146 = _module.D5.create_DC16(_1133_v151);
        let _rhs147 = _1135_v153;
        _1129_v147 = _rhs144;
        _1130_v148 = _rhs145;
        _1134_v152 = _rhs146;
        _1135_v153 = _rhs147;
      }
      let _1136_v154;
      _1136_v154 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),(_this).f21);
      _1136_v154 = (_1136_v154).update(!(p2), (_this).f21);
      let _1137_v155;
      _1137_v155 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1138_v156;
      _1138_v156 = _dafny.MultiSet.fromElements(_1137_v155);
      let _1139_v157;
      _1139_v157 = _dafny.Seq.of(new BigNumber((_1138_v156).cardinality()));
      let _1140_v158;
      _1140_v158 = _module.D0.create_DC0((_this).f21);
      let _1141_v159;
      _1141_v159 = _module.D6.create_DC20(p2, new BigNumber(642), p2);
      let _1142_v160;
      let _nw200 = Array((new BigNumber(22)).toNumber());
      _nw200[(_dafny.ZERO).toNumber()] = _module.__default.fm0(p2, _1137_v155, true, globalState);
      _nw200[(_dafny.ONE).toNumber()] = p2;
      _nw200[(new BigNumber(2)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(3)).toNumber()] = p2;
      _nw200[(new BigNumber(4)).toNumber()] = p2;
      _nw200[(new BigNumber(5)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(6)).toNumber()] = false;
      _nw200[(new BigNumber(7)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(8)).toNumber()] = (_1140_v158).dtor_cf0;
      _nw200[(new BigNumber(9)).toNumber()] = p2;
      _nw200[(new BigNumber(10)).toNumber()] = _module.__default.fm0(false, _1137_v155, (_1141_v159).dtor_cf36, globalState);
      _nw200[(new BigNumber(11)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(12)).toNumber()] = p2;
      _nw200[(new BigNumber(13)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(14)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(15)).toNumber()] = !((_this).f21);
      _nw200[(new BigNumber(16)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(17)).toNumber()] = false;
      _nw200[(new BigNumber(18)).toNumber()] = false;
      _nw200[(new BigNumber(19)).toNumber()] = false;
      _nw200[(new BigNumber(20)).toNumber()] = (_this).f21;
      _nw200[(new BigNumber(21)).toNumber()] = (_this).f21;
      _1142_v160 = _nw200;
      let _pat_let_tv29 = p0;
      let _pat_let_tv30 = p0;
      let _1143_v161;
      let _nw201 = Array((new BigNumber(19)).toNumber());
      _nw201[(_dafny.ZERO).toNumber()] = p0;
      _nw201[(_dafny.ONE).toNumber()] = p0;
      _nw201[(new BigNumber(2)).toNumber()] = p0;
      _nw201[(new BigNumber(3)).toNumber()] = p0;
      _nw201[(new BigNumber(4)).toNumber()] = p0;
      _nw201[(new BigNumber(5)).toNumber()] = p0;
      _nw201[(new BigNumber(6)).toNumber()] = new BigNumber(130);
      _nw201[(new BigNumber(7)).toNumber()] = p0;
      _nw201[(new BigNumber(8)).toNumber()] = p0;
      _nw201[(new BigNumber(9)).toNumber()] = p0;
      _nw201[(new BigNumber(10)).toNumber()] = new BigNumber(714);
      _nw201[(new BigNumber(11)).toNumber()] = new BigNumber((_1139_v157).length);
      _nw201[(new BigNumber(12)).toNumber()] = (function (_pat_let15_0) {
        return function (_1144_dt__update__tmp_h1) {
          return function (_pat_let16_0) {
            return function (_1145_dt__update_hcf15_h0) {
              return function (_pat_let17_0) {
                return function (_1146_dt__update_hcf14_h0) {
                  return _module.D2.create_DC8((_1144_dt__update__tmp_h1).dtor_cf13, _1146_dt__update_hcf14_h0, _1145_dt__update_hcf15_h0);
                }(_pat_let17_0);
              }(_pat_let_tv30);
            }(_pat_let16_0);
          }((_dafny.ZERO).minus(_pat_let_tv29));
        }(_pat_let15_0);
      }(_module.D2.create_DC8(_1142_v160, p0, new BigNumber((_1136_v154).length)))).dtor_cf15;
      _nw201[(new BigNumber(13)).toNumber()] = p0;
      _nw201[(new BigNumber(14)).toNumber()] = p0;
      _nw201[(new BigNumber(15)).toNumber()] = p0;
      _nw201[(new BigNumber(16)).toNumber()] = new BigNumber(20);
      _nw201[(new BigNumber(17)).toNumber()] = p0;
      _nw201[(new BigNumber(18)).toNumber()] = p0;
      _1143_v161 = _nw201;
      let _1147_v162;
      let _nw202 = new _module.C1();
      _nw202.__ctor(_1143_v161);
      _1147_v162 = _nw202;
      let _1148_v163;
      _1148_v163 = _dafny.Set.fromElements(_1147_v162);
      let _1149_v164;
      _1149_v164 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1148_v163);
      r0 = !(((_1148_v163).equals((((_1149_v164).contains(p1)) ? ((_1149_v164).get(p1)) : (_1148_v163)))) || ((p0).isLessThanOrEqualTo(new BigNumber(248))));
      r1 = _module.__default.fm0(p2, _1137_v155, (p0).isLessThan(p0), globalState);
      r2 = !((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, p0)))).isEqualTo(p0);
      r3 = p1;
      return [r0, r1, r2, r3];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)).plus(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(!(true))).length), new BigNumber((_dafny.Seq.UnicodeFromString("b")).length), new BigNumber(697), new BigNumber(758), new BigNumber((_dafny.Seq.UnicodeFromString("dssq")).length))).length)))).cardinality())))).isEqualTo((new BigNumber(414)).minus(new BigNumber(985)));
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      let _source13 = _module.D0.create_DC0(true);
      if (_source13.is_DC1) {
        let _1150___mcc_h0 = (_source13).cf1;
        let _1151_cf1 = _1150___mcc_h0;
        return _1151_cf1;
      } else if (_source13.is_DC0) {
        let _1152___mcc_h1 = (_source13).cf0;
        let _1153_cf0 = _1152___mcc_h1;
        return _1153_cf0;
      } else {
        let _1154___mcc_h2 = (_source13).cf2;
        let _1155_cf2 = _1154___mcc_h2;
        return false;
      }
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Map.Empty;
      let r3 = _dafny.Seq.UnicodeFromString("");
      (globalState).f11 = true;
      let _1156_i0;
      _1156_i0 = _dafny.ZERO;
      L7: {
        while (p1) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1156_i0)) {
              break L7;
            }
            _1156_i0 = (_1156_i0).plus(_dafny.ONE);
            let _1157_v0;
            let _init34 = ((_1158_p0) => function (_1159_i1) {
              return (_1159_i1).multipliedBy(_1158_p0);
            })(p0);
            let _nw203 = Array((new BigNumber(23)).toNumber());
            for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw203.length); _i0_34++) {
              _nw203[_i0_34] = _init34(new BigNumber(_i0_34));
            }
            _1157_v0 = _nw203;
            let _1160_v1;
            let _nw204 = new _module.C1();
            _nw204.__ctor(_1157_v0);
            _1160_v1 = _nw204;
            let _1161_v2;
            _1161_v2 = _dafny.Seq.of(p2, p2);
            let _1162_v3;
            _1162_v3 = _dafny.Set.fromElements(p0, p0);
            let _1163_v4;
            _1163_v4 = _dafny.Seq.of(_1162_v3);
            let _1164_v5;
            _1164_v5 = _dafny.Seq.of((_1161_v2)[_module.__default.safeIndex(new BigNumber(((_1163_v4)[_module.__default.safeIndex(p0, new BigNumber((_1163_v4).length))]).length), new BigNumber((_1161_v2).length))], p2);
            let _1165_v6;
            _1165_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-984)),_1164_v5);
            let _1166_v7;
            _1166_v7 = _dafny.Seq.of(p1);
            let _index150 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
            (_1157_v0)[_index150] = (new BigNumber(((((_1165_v6).contains(p0)) ? ((_1165_v6).get(p0)) : (_dafny.Seq.update(_1164_v5, _module.__default.safeIndex(p0, new BigNumber((_1164_v5).length)), p2)))).length)).plus(new BigNumber((_dafny.MultiSet.FromArray(_1166_v7)).cardinality()));
            let _index151 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
            (_1157_v0)[_index151] = (p0).multipliedBy(p0);
            let _1167_v8;
            let _nw205 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1167_v8 = _nw205;
            let _1168_v9;
            _1168_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1167_v8,!(true));
            if ((((_1168_v9).contains(_1167_v8)) ? ((_1168_v9).get(_1167_v8)) : (((true) ? (p1) : (p1))))) {
              (globalState).f15 = (_this).fm7((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], p1, (p0).multipliedBy(new BigNumber(-584)), globalState);
              let _index152 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length));
              (p2)[_index152] = ((false) ? (p1) : (false));
              let _1169_v10;
              _1169_v10 = _dafny.Seq.UnicodeFromString("wngf");
              let _index153 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length));
              let _rhs148 = _1169_v10;
              let _rhs149 = p1;
              let _lhs79 = p2;
              let _lhs80 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length));
              r3 = _rhs148;
              _lhs79[_lhs80] = _rhs149;
              let _1170_v11;
              _1170_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1169_v10).length),p1);
              let _1171_v12;
              _1171_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _1172_v13;
              _1172_v13 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length))],(((_1170_v11).contains(p0)) ? ((_1170_v11).get(p0)) : ((_this).fm6(_1171_v12, globalState))));
              let _1173_v14;
              _1173_v14 = _module.D6.create_DC20(_dafny.Seq.IsProperPrefixOf(_1169_v10, _dafny.Seq.update(_1169_v10, _module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_1169_v10).length)), new _dafny.CodePoint('v'.codePointAt(0)))), (_dafny.ZERO).minus((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]), (((_1172_v13).contains((p2)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length))])) ? ((_1172_v13).get((p2)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length))])) : (p1)));
              _1173_v14 = _module.D6.create_DC20(true, _module.__default.fm9(p0, new BigNumber((_1169_v10).length), (p2)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length))], globalState), p1);
              (globalState).f11 = p1;
              let _1174_v15;
              let _init35 = ((_1175_p1) => function (_1176_i2) {
                return _1175_p1;
              })(p1);
              let _nw206 = Array((new BigNumber(25)).toNumber());
              for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw206.length); _i0_35++) {
                _nw206[_i0_35] = _init35(new BigNumber(_i0_35));
              }
              _1174_v15 = _nw206;
              let _1177_v16;
              let _init36 = ((_1178_v10) => function (_1179_i3) {
                return _dafny.Seq.of(new BigNumber((_1178_v10).length));
              })(_1169_v10);
              let _nw207 = Array((new BigNumber(28)).toNumber());
              for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw207.length); _i0_36++) {
                _nw207[_i0_36] = _init36(new BigNumber(_i0_36));
              }
              _1177_v16 = _nw207;
              let _1180_v17;
              _1180_v17 = _dafny.MultiSet.fromElements(p0);
              let _index154 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1177_v16).length));
              (_1177_v16)[_index154] = _dafny.Seq.of(new BigNumber((_1180_v17).cardinality()));
              let _1181_v18;
              _1181_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.__default.fm16(p1, p0, globalState), _module.__default.fm16((p2)[_module.__default.safeIndex(new BigNumber(346), new BigNumber((p2).length))], new BigNumber((_dafny.Seq.UnicodeFromString("cfbueqff")).length), globalState))),_1174_v15);
              let _1182_v19;
              _1182_v19 = _dafny.MultiSet.fromElements(_1169_v10, _1169_v10, _1169_v10, _1169_v10, _1169_v10);
              let _1183_v20;
              _1183_v20 = _dafny.Seq.of((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], p0, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])));
              let _index155 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1177_v16).length));
              let _index156 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              let _rhs150 = (((_1181_v18).contains(_1182_v19)) ? ((_1181_v18).get(_1182_v19)) : (_1174_v15));
              let _rhs151 = _1183_v20;
              let _rhs152 = p0;
              let _lhs81 = _1177_v16;
              let _lhs82 = _module.__default.safeIndex(new BigNumber(988), new BigNumber((_1177_v16).length));
              let _lhs83 = _1157_v0;
              let _lhs84 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              _1174_v15 = _rhs150;
              _lhs81[_lhs82] = _rhs151;
              _lhs83[_lhs84] = _rhs152;
            } else {
              let _index157 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              (_1157_v0)[_index157] = p0;
              let _1184_v21;
              let _nw208 = new _module.C0();
              _nw208.__ctor(((_module.__default.fm0(p1, p3, p1, globalState)) ? (p1) : (p1)));
              _1184_v21 = _nw208;
              let _1185_v22;
              _1185_v22 = _dafny.Seq.UnicodeFromString("xergafjq");
              r0 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_1184_v21).f23, p1, _module.__default.fm0(p1, (_1185_v22)[_module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_1185_v22).length))], p1, globalState)), _1166_v7), _module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_1184_v21).f23, p1, _module.__default.fm0(p1, (_1185_v22)[_module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_1185_v22).length))], p1, globalState)), _1166_v7)).length)), p1);
              let _1186_v23;
              _1186_v23 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])).length), (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], p0);
              r1 = ((((_1186_v23).contains((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])) ? ((_1186_v23).get((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])) : (p0))).minus((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _1187_v24;
              _1187_v24 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1160_v1);
              _1187_v24 = (_1187_v24).update((_module.D4.create_DC14(p1, (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], _1162_v3, new _dafny.CodePoint('s'.codePointAt(0)))).dtor_cf26, _1160_v1);
            }
            if (p1) {
              let _1188_v25;
              _1188_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1162_v3);
              r1 = (new BigNumber((((((_1188_v25).contains((_dafny.ZERO).minus((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]))) ? ((_1188_v25).get((_dafny.ZERO).minus((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]))) : (_1162_v3))).Intersect(_1162_v3)).length)).multipliedBy((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _1189_v26;
              _1189_v26 = _dafny.MultiSet.fromElements(p1);
              let _index158 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              let _rhs153 = _1189_v26;
              let _rhs154 = ((p1) ? ((p1) && (!(!(p1)))) : (p1));
              let _rhs155 = (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))];
              let _lhs85 = globalState;
              let _lhs86 = _1157_v0;
              let _lhs87 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              _1189_v26 = _rhs153;
              _lhs85.f15 = _rhs154;
              _lhs86[_lhs87] = _rhs155;
              let _1190_v27;
              _1190_v27 = _dafny.MultiSet.fromElements(p0);
              let _1191_v28;
              _1191_v28 = _module.D6.create_DC20(!(!(p1)) || (true), ((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]).plus(new BigNumber((_1190_v27).cardinality())), p1);
              let _1192_v29;
              _1192_v29 = _dafny.Set.fromElements(p1);
              let _index159 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              let _rhs156 = (((_1192_v29).IsDisjointFrom(_1192_v29)) ? (_1191_v28) : (_1191_v28));
              let _rhs157 = p1;
              let _rhs158 = p1;
              let _rhs159 = _module.__default.safeDivisionInt((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _lhs88 = globalState;
              let _lhs89 = globalState;
              let _lhs90 = _1157_v0;
              let _lhs91 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              _1191_v28 = _rhs156;
              _lhs88.f15 = _rhs157;
              _lhs89.f11 = _rhs158;
              _lhs90[_lhs91] = _rhs159;
              let _1193_v30;
              _1193_v30 = _dafny.Seq.UnicodeFromString("gghwthy");
              let _1194_v31;
              _1194_v31 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pqjkhiid"), _dafny.Seq.update(_1193_v30, _module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_1193_v30).length)), p3));
              let _1195_v32;
              _1195_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              _1194_v31 = (((_this).fm6(_1195_v32, globalState)) ? (_1194_v31) : (_dafny.Seq.of(_1193_v30)));
              r1 = new BigNumber(277);
            } else {
              let _index160 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              (_1157_v0)[_index160] = (_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))];
              let _1196_v33;
              _1196_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _1197_v35;
              _1197_v35 = _dafny.Seq.UnicodeFromString("t");
              let _1198_v36;
              _1198_v36 = _dafny.MultiSet.fromElements(_1196_v33, function () {
                let _coll34 = new _dafny.Map();
                for (const _compr_34 of _dafny.IntegerRange(new BigNumber(617), new BigNumber(957))) {
                  let _1199_v34 = _compr_34;
                  if (((new BigNumber(617)).isLessThanOrEqualTo(_1199_v34)) && ((_1199_v34).isLessThan(new BigNumber(957)))) {
                    _coll34.push([(_1199_v34).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("uvscymnab")).length)),p0]);
                  }
                }
                return _coll34;
              }(), (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1197_v35).length),(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])).Merge(_1196_v33));
              let _1200_v37;
              _1200_v37 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))]);
              let _1201_v38;
              _1201_v38 = _dafny.Seq.of((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], p0, _module.__default.fm9(new BigNumber((_1200_v37).length), p0, p1, globalState));
              let _index161 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              let _rhs160 = false;
              let _rhs161 = (_1162_v3).IsSubsetOf(_1162_v3);
              let _rhs162 = (_1160_v1).f22;
              let _rhs163 = (((_1198_v36).contains((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(54),(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])).Merge(_1196_v33))) ? ((_1198_v36).get((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(54),(_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))])).Merge(_1196_v33))) : ((_1201_v38)[_module.__default.safeIndex((_1157_v0)[_module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length))], new BigNumber((_1201_v38).length))]));
              let _rhs164 = p0;
              let _lhs92 = globalState;
              let _lhs93 = globalState;
              let _lhs94 = _1157_v0;
              let _lhs95 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_1157_v0).length));
              _lhs92.f11 = _rhs160;
              _lhs93.f11 = _rhs161;
              _1157_v0 = _rhs162;
              r1 = _rhs163;
              _lhs94[_lhs95] = _rhs164;
              let _1202_v39;
              _1202_v39 = _module.D9.create_DC24(_1201_v38);
              let _rhs165 = !(new BigNumber(216)).isEqualTo((new BigNumber(((_1202_v39).dtor_cf42).length)).minus(new BigNumber(970)));
              let _rhs166 = !((false) === (false));
              let _lhs96 = globalState;
              let _lhs97 = globalState;
              _lhs96.f15 = _rhs165;
              _lhs97.f11 = _rhs166;
              let _1203_v40;
              _1203_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1201_v38,p1);
              _1203_v40 = (_1203_v40).update(_1201_v38, p1);
              let _index162 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((p2).length));
              (p2)[_index162] = p1;
              let _1204_v41;
              let _init37 = ((_1205_v35) => function (_1206_i4) {
                return _1205_v35;
              })(_1197_v35);
              let _nw209 = Array((new BigNumber(24)).toNumber());
              for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw209.length); _i0_37++) {
                _nw209[_i0_37] = _init37(new BigNumber(_i0_37));
              }
              _1204_v41 = _nw209;
              let _index163 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1204_v41).length));
              (_1204_v41)[_index163] = _module.__default.fm16(true, p0, globalState);
              let _1207_v42;
              let _nw210 = Array((new BigNumber(5)).toNumber());
              _nw210[(_dafny.ZERO).toNumber()] = (_1160_v1).f22;
              _nw210[(_dafny.ONE).toNumber()] = (_1160_v1).f22;
              _nw210[(new BigNumber(2)).toNumber()] = _1157_v0;
              _nw210[(new BigNumber(3)).toNumber()] = (_1160_v1).f22;
              _nw210[(new BigNumber(4)).toNumber()] = (_1160_v1).f22;
              _1207_v42 = _nw210;
              let _index164 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((p2).length));
              let _index165 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1204_v41).length));
              let _rhs167 = p1;
              let _rhs168 = ((p1) ? (_1207_v42) : (_1207_v42));
              let _rhs169 = _1197_v35;
              let _lhs98 = p2;
              let _lhs99 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((p2).length));
              let _lhs100 = globalState;
              let _lhs101 = _1204_v41;
              let _lhs102 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1204_v41).length));
              _lhs98[_lhs99] = _rhs167;
              _lhs100.f2 = _rhs168;
              _lhs101[_lhs102] = _rhs169;
            }
          }
        }
      }
      let _1208_v43;
      _1208_v43 = _dafny.MultiSet.fromElements(new BigNumber(640));
      _1208_v43 = _1208_v43;
      let _1209_v46;
      let _init38 = ((_1210_p0, _1211_p1) => function (_1212_i5) {
        return (function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_dafny.Seq.of(_1210_p0)).Elements) {
            let _1213_v44 = _compr_35;
            if (_dafny.Seq.contains(_dafny.Seq.of(_1210_p0), _1213_v44)) {
              _coll35.push([(_1213_v44).multipliedBy(_1210_p0),_1211_p1]);
            }
          }
          return _coll35;
        }()).Merge(function () {
          let _coll36 = new _dafny.Map();
          for (const _compr_36 of (_dafny.MultiSet.fromElements(_1210_p0)).Elements) {
            let _1214_v45 = _compr_36;
            if ((_dafny.MultiSet.fromElements(_1210_p0)).contains(_1214_v45)) {
              _coll36.push([(_1214_v45).minus(_1210_p0),_1211_p1]);
            }
          }
          return _coll36;
        }());
      })(p0, p1);
      let _nw211 = Array((new BigNumber(23)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw211.length); _i0_38++) {
        _nw211[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1209_v46 = _nw211;
      let _nw212 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
      _1209_v46 = _nw212;
      let _1215_v47;
      _1215_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm9(p0, new BigNumber(383), !(p1), globalState));
      let _1216_v48;
      _1216_v48 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
      _1215_v47 = (_1215_v47).update(new BigNumber(((_1216_v48).Merge(_1216_v48)).length), p0);
      let _hi8 = p0;
      for (let _1217_i6 = p0; _1217_i6.isLessThan(_hi8); _1217_i6 = _1217_i6.plus(_dafny.ONE)) {
        (globalState).f11 = p1;
        (globalState).f11 = (p1) && (p1);
        let _1218_v49;
        _1218_v49 = _module.D6.create_DC20(false, p0, p1);
        let _1219_v50;
        _1219_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _pat_let_tv31 = p1;
        let _pat_let_tv32 = _1219_v50;
        let _1220_v51;
        _1220_v51 = _dafny.Seq.of(_1218_v49, function (_pat_let18_0) {
          return function (_1221_dt__update__tmp_h0) {
            return function (_pat_let19_0) {
              return function (_1222_dt__update_hcf36_h0) {
                return function (_pat_let20_0) {
                  return function (_1223_dt__update_hcf37_h0) {
                    return _module.D6.create_DC20(_1222_dt__update_hcf36_h0, _1223_dt__update_hcf37_h0, (_1221_dt__update__tmp_h0).dtor_cf38);
                  }(_pat_let20_0);
                }(new BigNumber((_pat_let_tv32).length));
              }(_pat_let19_0);
            }(!(_pat_let_tv31));
          }(_pat_let18_0);
        }(_1218_v49));
        let _1224_v52;
        _1224_v52 = new _dafny.CodePoint('j'.codePointAt(0));
        let _rhs170 = _1220_v51;
        let _rhs171 = p0;
        let _rhs172 = ((true) ? (_1224_v52) : (_1224_v52));
        _1220_v51 = _rhs170;
        r1 = _rhs171;
        _1224_v52 = _rhs172;
        let _1225_v53;
        _1225_v53 = _module.D0.create_DC1(true);
        let _pat_let_tv33 = p1;
        let _1226_v54;
        let _nw213 = new _module.C2();
        _nw213.__ctor(function (_pat_let21_0) {
          return function (_1227_dt__update__tmp_h1) {
            return function (_pat_let22_0) {
              return function (_1228_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_1228_dt__update_hcf1_h0);
              }(_pat_let22_0);
            }(!(_pat_let_tv33));
          }(_pat_let21_0);
        }(_1225_v53), !((p1) && (p1)));
        _1226_v54 = _nw213;
      }
      let _1229_v55;
      _1229_v55 = _dafny.Seq.of(p1, p1, false, true, p1);
      r0 = _dafny.Seq.Concat(_1229_v55, _1229_v55);
      r1 = p0;
      let _1230_v56;
      _1230_v56 = _dafny.Seq.of(_1216_v48);
      r2 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_1230_v56).length), p0),(p1) === (!(p1)));
      let _1231_v57;
      _1231_v57 = _dafny.Seq.UnicodeFromString("suhxbcocc");
      r3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1231_v57, _1231_v57), _1231_v57);
      return [r0, r1, r2, r3];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f19 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor(f19) {
      let _this = this;
      (_this)._f19 = f19;
      return;
    }
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f19;
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1232_v0;
      let _nw214 = Array((new BigNumber(7)).toNumber()).fill(false);
      _1232_v0 = _nw214;
      let _1233_v1;
      _1233_v1 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1234_v2;
      _1234_v2 = _dafny.Seq.of(_1233_v1);
      let _1235_v3;
      _1235_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1234_v2,_1232_v0);
      let _1236_v4;
      _1236_v4 = _dafny.Seq.of(_1232_v0);
      let _1237_v5;
      let _nw215 = Array((new BigNumber(19)).toNumber());
      _nw215[(_dafny.ZERO).toNumber()] = _1232_v0;
      _nw215[(_dafny.ONE).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(2)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(3)).toNumber()] = ((p0) ? (_1232_v0) : (_1232_v0));
      _nw215[(new BigNumber(4)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(5)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(6)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(7)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(8)).toNumber()] = (((_1235_v3).contains(_dafny.Seq.of(_1233_v1))) ? ((_1235_v3).get(_dafny.Seq.of(_1233_v1))) : (_1232_v0));
      _nw215[(new BigNumber(9)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(10)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(11)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(12)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(13)).toNumber()] = (_1236_v4)[_module.__default.safeIndex(new BigNumber(-8), new BigNumber((_1236_v4).length))];
      _nw215[(new BigNumber(14)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(15)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(16)).toNumber()] = (_module.D1.create_DC4(new BigNumber(51), p2, _1232_v0)).dtor_cf6;
      _nw215[(new BigNumber(17)).toNumber()] = _1232_v0;
      _nw215[(new BigNumber(18)).toNumber()] = _1232_v0;
      _1237_v5 = _nw215;
      let _index166 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length));
      (_1237_v5)[_index166] = _1232_v0;
      let _1238_v6;
      _1238_v6 = _dafny.Seq.of(_1234_v2);
      let _1239_v7;
      _1239_v7 = _dafny.MultiSet.fromElements(_1233_v1);
      let _1240_v8;
      _1240_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_1239_v7).contains(_1233_v1)) ? ((_1239_v7).get(_1233_v1)) : (p1)),p2);
      let _index167 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length));
      let _rhs173 = _1232_v0;
      let _rhs174 = _module.__default.fm0(!(!(p1).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1233_v1)).length))), _1233_v1, !(_1240_v8).equals(_1240_v8), globalState);
      let _rhs175 = _1238_v6;
      let _lhs103 = _1237_v5;
      let _lhs104 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length));
      let _lhs105 = globalState;
      _lhs103[_lhs104] = _rhs173;
      _lhs105.f15 = _rhs174;
      _1238_v6 = _rhs175;
      if (!((_this).f19)) {
        let _index168 = _module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length));
        (_1237_v5)[_index168] = (_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))];
        let _1241_v9;
        _1241_v9 = _dafny.MultiSet.fromElements(p1, _dafny.ONE);
        (globalState).f11 = (_1241_v9).IsDisjointFrom(_1241_v9);
        let _1242_v10;
        _1242_v10 = new BigNumber(938);
        let _1243_v11;
        _1243_v11 = _dafny.Seq.of(p2);
        _1242_v10 = _module.__default.fm5(_module.__default.safeDivisionInt(new BigNumber((_1243_v11).length), p1), globalState);
        _1234_v2 = _1234_v2;
        let _1244_v12;
        _1244_v12 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1232_v0);
        let _1245_v13;
        _1245_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        _1244_v12 = (_1244_v12).update(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_1243_v11, _module.__default.safeIndex(new BigNumber((_1245_v13).length), new BigNumber((_1243_v11).length)), p0)).length), p3), (_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))]);
      } else {
        if ((_this).f19) {
          let _1246_v14;
          _1246_v14 = _dafny.Seq.of(p1, p1);
          let _1247_v15;
          let _nw216 = Array((new BigNumber(11)).toNumber());
          _nw216[(_dafny.ZERO).toNumber()] = p3;
          _nw216[(_dafny.ONE).toNumber()] = (_1246_v14)[_module.__default.safeIndex(p3, new BigNumber((_1246_v14).length))];
          _nw216[(new BigNumber(2)).toNumber()] = p1;
          _nw216[(new BigNumber(3)).toNumber()] = p3;
          _nw216[(new BigNumber(4)).toNumber()] = p3;
          _nw216[(new BigNumber(5)).toNumber()] = p3;
          _nw216[(new BigNumber(6)).toNumber()] = p1;
          _nw216[(new BigNumber(7)).toNumber()] = p3;
          _nw216[(new BigNumber(8)).toNumber()] = p3;
          _nw216[(new BigNumber(9)).toNumber()] = p1;
          _nw216[(new BigNumber(10)).toNumber()] = new BigNumber((_1238_v6).length);
          _1247_v15 = _nw216;
          let _1248_v16;
          _1248_v16 = _dafny.Map.Empty.slice().updateUnsafe((p2) || ((_this).f19),_1247_v15);
          _1248_v16 = (_1248_v16).Merge((_1248_v16).Merge(_1248_v16));
          let _1249_v17;
          _1249_v17 = new BigNumber(919);
          let _1250_v18;
          _1250_v18 = _dafny.MultiSet.fromElements(_1237_v5);
          _1249_v17 = (((_1250_v18).contains(_1237_v5)) ? ((_1250_v18).get(_1237_v5)) : (p3));
          let _1251_v19;
          _1251_v19 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_dafny.ZERO).minus(p1));
          _1251_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1249_v17,_module.__default.safeModuloInt(p1, _1249_v17));
          let _1252_v20;
          let _nw217 = new _module.C3();
          _nw217.__ctor();
          _1252_v20 = _nw217;
          let _1253_v21;
          _1253_v21 = _dafny.Seq.of(p0, p2);
          let _rhs176 = !(p2);
          let _rhs177 = !((_this).f19) || (p0);
          let _rhs178 = (p3).isLessThan((((_1251_v19).contains(new BigNumber((_dafny.MultiSet.FromArray(_1253_v21)).cardinality()))) ? ((_1251_v19).get(new BigNumber((_dafny.MultiSet.FromArray(_1253_v21)).cardinality()))) : (_1249_v17)));
          let _rhs179 = p2;
          let _rhs180 = p1;
          let _lhs106 = globalState;
          let _lhs107 = globalState;
          let _lhs108 = globalState;
          r1 = _rhs176;
          _lhs106.f15 = _rhs177;
          _lhs107.f17 = _rhs178;
          _lhs108.f17 = _rhs179;
          _1249_v17 = _rhs180;
        } else {
          let _1254_v22;
          let _nw218 = Array((_dafny.ONE).toNumber()).fill(_dafny.MultiSet.Empty);
          _1254_v22 = _nw218;
          let _index169 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1254_v22).length));
          (_1254_v22)[_index169] = _dafny.MultiSet.fromElements(p0);
          let _1255_v23;
          _1255_v23 = _dafny.MultiSet.fromElements(p2);
          let _1256_v24;
          _1256_v24 = _module.D10.create_DC27(_1255_v23);
          let _index170 = _module.__default.safeIndex(new BigNumber(114), new BigNumber((_1254_v22).length));
          (_1254_v22)[_index170] = (((_this).f19) ? ((_1256_v24).dtor_cf48) : (_1255_v23));
          let _1257_v25;
          _1257_v25 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.__default.safeDivisionInt(new BigNumber((_1234_v2).length), (_dafny.ZERO).minus(new BigNumber((_1240_v8).length))));
          let _1258_v26;
          _1258_v26 = _dafny.MultiSet.fromElements(p1);
          _1257_v25 = (_1257_v25).update(!((new BigNumber((_1258_v26).cardinality())).isLessThanOrEqualTo(p3)), (_dafny.ZERO).minus(new BigNumber((_1257_v25).length)));
          _1240_v8 = (_1240_v8).update(p1, p2);
          r0 = (_this).f19;
          let _1259_v27;
          let _nw219 = Array((new BigNumber(21)).toNumber());
          _1259_v27 = _nw219;
          _1259_v27 = _1259_v27;
        }
        let _1260_v28;
        _1260_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1233_v1,p3);
        let _1261_v29;
        _1261_v29 = _dafny.Set.fromElements(_module.__default.fm9(p1, p3, p2, globalState));
        let _1262_v30;
        _1262_v30 = _module.D4.create_DC14(p0, (((_1260_v28).contains(_1233_v1)) ? ((_1260_v28).get(_1233_v1)) : (p3)), _1261_v29, _1233_v1);
        let _1263_v31;
        _1263_v31 = _module.D6.create_DC20(_module.__default.fm0(p0, (_1262_v30).dtor_cf26, p0, globalState), p1, (_this).f19);
        if (!((function (_pat_let23_0) {
          return function (_1264_dt__update__tmp_h0) {
            return function (_pat_let24_0) {
              return function (_1265_dt__update_hcf37_h0) {
                return _module.D6.create_DC20((_1264_dt__update__tmp_h0).dtor_cf36, _1265_dt__update_hcf37_h0, (_1264_dt__update__tmp_h0).dtor_cf38);
              }(_pat_let24_0);
            }(new BigNumber(115));
          }(_pat_let23_0);
        }(_1263_v31)).dtor_cf38)) {
          let _1266_v32;
          _1266_v32 = new BigNumber(203);
          let _1267_v33;
          _1267_v33 = _dafny.Seq.of(p3);
          _1266_v32 = (_1267_v33)[_module.__default.safeIndex(p3, new BigNumber((_1267_v33).length))];
          r0 = p2;
          (globalState).f15 = !(p2);
          let _1268_v34;
          let _nw220 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _1268_v34 = _nw220;
          let _index171 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length));
          (_1268_v34)[_index171] = p3;
          let _1269_v35;
          _1269_v35 = _module.D1.create_DC4((_dafny.ZERO).minus(p1), (_this).f19, _1232_v0);
          let _1270_v36;
          _1270_v36 = _dafny.Set.fromElements(false, (_this).f19, (_this).f19, p2, (_this).f19);
          let _1271_v37;
          _1271_v37 = _module.D2.create_DC8((_1269_v35).dtor_cf6, new BigNumber(770), new BigNumber((_1270_v36).length));
          let _index172 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length));
          let _rhs181 = _1266_v32;
          let _rhs182 = (_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))];
          let _rhs183 = (_1271_v37).dtor_cf13;
          let _lhs109 = _1268_v34;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length));
          _lhs109[_lhs110] = _rhs181;
          _1232_v0 = _rhs182;
          _1232_v0 = _rhs183;
          let _1272_v38;
          _1272_v38 = _module.D1.create_DC5(_1234_v2, (_dafny.ZERO).minus(p3), p3);
          let _1273_v39;
          _1273_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1234_v2);
          let _1274_v40;
          let _nw221 = Array((new BigNumber(28)).toNumber());
          _nw221[(_dafny.ZERO).toNumber()] = _1234_v2;
          _nw221[(_dafny.ONE).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(2)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(3)).toNumber()] = ((true) ? (_1234_v2) : (_1234_v2));
          _nw221[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_1234_v2, _1234_v2);
          _nw221[(new BigNumber(5)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1234_v2, (_1272_v38).dtor_cf7);
          _nw221[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1234_v2, _1234_v2);
          _nw221[(new BigNumber(8)).toNumber()] = (_1238_v6)[_module.__default.safeIndex(p3, new BigNumber((_1238_v6).length))];
          _nw221[(new BigNumber(9)).toNumber()] = (_1272_v38).dtor_cf7;
          _nw221[(new BigNumber(10)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1234_v2, (_1238_v6)[_module.__default.safeIndex(p1, new BigNumber((_1238_v6).length))]);
          _nw221[(new BigNumber(12)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("fpbix");
          _nw221[(new BigNumber(14)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(15)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_1275_v1) => function (_1276_i0) {
            return _1275_v1;
          })(_1233_v1)), _1234_v2);
          _nw221[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), ((_1277_v1) => function (_1278_i1) {
            return _1277_v1;
          })(_1233_v1)), _module.__default.safeIndex(new BigNumber(398), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), ((_1279_v1) => function (_1280_i1) {
            return _1279_v1;
          })(_1233_v1))).length)), _1233_v1), (((_1273_v39).contains(p2)) ? ((_1273_v39).get(p2)) : (_dafny.Seq.UnicodeFromString("noveiuir"))));
          _nw221[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm16(!(p2), (_1268_v34)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length))], globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), ((_1281_v1) => function (_1282_i2) {
            return _1281_v1;
          })(_1233_v1))), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.Concat(_module.__default.fm16(!(p2), (_1268_v34)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length))], globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), ((_1283_v1) => function (_1284_i2) {
            return _1283_v1;
          })(_1233_v1)))).length)), _1233_v1);
          _nw221[(new BigNumber(19)).toNumber()] = ((p2) ? (_1234_v2) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(898)), ((_1285_v1) => function (_1286_i3) {
            return _1285_v1;
          })(_1233_v1))));
          _nw221[(new BigNumber(20)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("vafiqo");
          _nw221[(new BigNumber(22)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(23)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-200)), ((_1287_v1) => function (_1288_i4) {
            return _1287_v1;
          })(_1233_v1));
          _nw221[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_1234_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(717)), ((_1289_v1) => function (_1290_i5) {
            return _1289_v1;
          })(_1233_v1)));
          _nw221[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_1234_v2, _1234_v2);
          _nw221[(new BigNumber(26)).toNumber()] = _1234_v2;
          _nw221[(new BigNumber(27)).toNumber()] = _1234_v2;
          _1274_v40 = _nw221;
          let _index173 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1274_v40).length));
          (_1274_v40)[_index173] = _1234_v2;
          let _1291_v41;
          _1291_v41 = _dafny.MultiSet.fromElements(p2, (_1266_v32).isLessThanOrEqualTo((_1268_v34)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length))]), p2);
          let _1292_v42;
          let _nw222 = new _module.C0();
          _nw222.__ctor(p0);
          _1292_v42 = _nw222;
          let _1293_v43;
          _1293_v43 = _module.D9.create_DC25(p3, p0, p1, _1292_v42);
          let _1294_v44;
          _1294_v44 = _dafny.MultiSet.fromElements((_1293_v43).dtor_cf45, new BigNumber((_module.__default.fm16((_1292_v42).f23, (_dafny.ZERO).minus(new BigNumber(-311)), globalState)).length));
          let _index174 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1274_v40).length));
          let _index175 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length));
          let _rhs184 = _1270_v36;
          let _rhs185 = _1234_v2;
          let _rhs186 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber((_1294_v44).cardinality()), new BigNumber((_1234_v2).length)), new BigNumber((_1261_v29).length));
          let _rhs187 = (_1291_v41).update((_1270_v36).IsProperSubsetOf(_1270_v36), _module.__default.abs(_1266_v32));
          let _rhs188 = (((_1240_v8).contains(((_1268_v34)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length))]).plus(p1))) ? ((_1240_v8).get(((_1268_v34)[_module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length))]).plus(p1))) : (p0));
          let _lhs111 = _1274_v40;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1274_v40).length));
          let _lhs113 = _1268_v34;
          let _lhs114 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_1268_v34).length));
          let _lhs115 = globalState;
          _1270_v36 = _rhs184;
          _lhs111[_lhs112] = _rhs185;
          _lhs113[_lhs114] = _rhs186;
          _1291_v41 = _rhs187;
          _lhs115.f15 = _rhs188;
        } else {
          let _1295_v45;
          _1295_v45 = _module.D4.create_DC13(p1);
          let _1296_v46;
          _1296_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1295_v45);
          _1296_v46 = (_1296_v46).update(p0, _1295_v45);
          let _1297_v47;
          let _nw223 = new _module.C3();
          _nw223.__ctor();
          _1297_v47 = _nw223;
          let _1298_v48;
          _1298_v48 = _dafny.ONE;
          let _1299_v49;
          let _nw224 = new _module.C2();
          _nw224.__ctor(_module.D0.create_DC1(p0), p2);
          _1299_v49 = _nw224;
          let _rhs189 = _1298_v48;
          let _rhs190 = _1299_v49;
          let _rhs191 = (_this).f19;
          let _lhs116 = globalState;
          _1298_v48 = _rhs189;
          _1299_v49 = _rhs190;
          _lhs116.f17 = _rhs191;
          let _1300_v50;
          _1300_v50 = _dafny.Set.fromElements((_this).f19, p0, p0);
          let _1301_v51;
          let _nw225 = Array((new BigNumber(27)).toNumber());
          _nw225[(_dafny.ZERO).toNumber()] = _1298_v48;
          _nw225[(_dafny.ONE).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(2)).toNumber()] = p1;
          _nw225[(new BigNumber(3)).toNumber()] = p1;
          _nw225[(new BigNumber(4)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(5)).toNumber()] = new BigNumber(242);
          _nw225[(new BigNumber(6)).toNumber()] = p3;
          _nw225[(new BigNumber(7)).toNumber()] = p1;
          _nw225[(new BigNumber(8)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(9)).toNumber()] = p3;
          _nw225[(new BigNumber(10)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_1298_v48);
          _nw225[(new BigNumber(12)).toNumber()] = p1;
          _nw225[(new BigNumber(13)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(14)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(15)).toNumber()] = p3;
          _nw225[(new BigNumber(16)).toNumber()] = new BigNumber((_1234_v2).length);
          _nw225[(new BigNumber(17)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(18)).toNumber()] = new BigNumber((_1300_v50).length);
          _nw225[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.of(p0, false)).length);
          _nw225[(new BigNumber(20)).toNumber()] = p1;
          _nw225[(new BigNumber(21)).toNumber()] = p3;
          _nw225[(new BigNumber(22)).toNumber()] = _1298_v48;
          _nw225[(new BigNumber(23)).toNumber()] = p1;
          _nw225[(new BigNumber(24)).toNumber()] = p3;
          _nw225[(new BigNumber(25)).toNumber()] = p1;
          _nw225[(new BigNumber(26)).toNumber()] = p3;
          _1301_v51 = _nw225;
          let _1302_v52;
          let _nw226 = new _module.C1();
          _nw226.__ctor(_1301_v51);
          _1302_v52 = _nw226;
          let _1303_v53;
          _1303_v53 = _dafny.Seq.of(_1302_v52);
          _1298_v48 = new BigNumber((_1303_v53).length);
          _1240_v8 = (_1240_v8).update(_1298_v48, _module.__default.fm0((_this).f19, new _dafny.CodePoint('a'.codePointAt(0)), (_this).f19, globalState));
        }
        _1234_v2 = _1234_v2;
        _1233_v1 = new _dafny.CodePoint('x'.codePointAt(0));
        if ((p3).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(p3)))) {
          let _1304_v54;
          let _init39 = ((_1305_v1) => function (_1306_i6) {
            return _1305_v1;
          })(_1233_v1);
          let _nw227 = Array((new BigNumber(20)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw227.length); _i0_39++) {
            _nw227[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1304_v54 = _nw227;
          let _1307_v55;
          _1307_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1304_v54,_module.__default.fm5(p3, globalState));
          let _1308_v56;
          _1308_v56 = _dafny.Map.Empty.slice().updateUnsafe(false,p3);
          let _1309_v57;
          _1309_v57 = _dafny.Map.Empty.slice().updateUnsafe((((_1308_v56).contains((_this).f19)) ? ((_1308_v56).get((_this).f19)) : (p1)),_1304_v54);
          let _1310_v58;
          _1310_v58 = _dafny.Seq.of(false);
          _1307_v55 = (_1307_v55).update((((_1309_v57).contains(p1)) ? ((_1309_v57).get(p1)) : (_1304_v54)), new BigNumber((_1310_v58).length));
          let _1311_v59;
          _1311_v59 = _dafny.Seq.of(p3, p1, p3);
          let _1312_v60;
          _1312_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_dafny.Seq.update(_1311_v59, _module.__default.safeIndex(p3, new BigNumber((_1311_v59).length)), new BigNumber((_dafny.Seq.UnicodeFromString("vbydlk")).length)));
          let _1313_v61;
          _1313_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(704),new BigNumber(((((_1312_v60).contains(p2)) ? ((_1312_v60).get(p2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), ((_1314_v2) => function (_1315_i7) {
            return new BigNumber((_1314_v2).length);
          })(_1234_v2))))).length));
          let _1316_v62;
          _1316_v62 = _dafny.Set.fromElements(p2);
          let _1317_v63;
          _1317_v63 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
          let _1318_v64;
          _1318_v64 = _module.D9.create_DC24(_module.__default.fm19(_1311_v59, (((_1317_v63).contains(p0)) ? ((_1317_v63).get(p0)) : ((_1310_v58)[_module.__default.safeIndex(p3, new BigNumber((_1310_v58).length))])), globalState));
          let _1319_v65;
          _1319_v65 = _module.D9.create_DC26(_1318_v64);
          let _1320_v66;
          _1320_v66 = _dafny.MultiSet.fromElements(_1319_v65, _1319_v65);
          let _1321_v67;
          let _nw228 = Array((new BigNumber(19)).toNumber());
          _nw228[(_dafny.ZERO).toNumber()] = new BigNumber((_1310_v58).length);
          _nw228[(_dafny.ONE).toNumber()] = p3;
          _nw228[(new BigNumber(2)).toNumber()] = p1;
          _nw228[(new BigNumber(3)).toNumber()] = new BigNumber(870);
          _nw228[(new BigNumber(4)).toNumber()] = p1;
          _nw228[(new BigNumber(5)).toNumber()] = p1;
          _nw228[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p3);
          _nw228[(new BigNumber(7)).toNumber()] = p3;
          _nw228[(new BigNumber(8)).toNumber()] = new BigNumber(738);
          _nw228[(new BigNumber(9)).toNumber()] = p3;
          _nw228[(new BigNumber(10)).toNumber()] = p3;
          _nw228[(new BigNumber(11)).toNumber()] = p1;
          _nw228[(new BigNumber(12)).toNumber()] = p3;
          _nw228[(new BigNumber(13)).toNumber()] = p3;
          _nw228[(new BigNumber(14)).toNumber()] = new BigNumber(-922);
          _nw228[(new BigNumber(15)).toNumber()] = p1;
          _nw228[(new BigNumber(16)).toNumber()] = p1;
          _nw228[(new BigNumber(17)).toNumber()] = p1;
          _nw228[(new BigNumber(18)).toNumber()] = p1;
          _1321_v67 = _nw228;
          let _1322_v68;
          _1322_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1311_v59).length),_1321_v67);
          let _1323_v69;
          _1323_v69 = _dafny.Seq.of((((_1322_v68).contains(_module.__default.fm9(new BigNumber((_1261_v29).length), p1, (_this).f19, globalState))) ? ((_1322_v68).get(_module.__default.fm9(new BigNumber((_1261_v29).length), p1, (_this).f19, globalState))) : (_1321_v67)));
          let _1324_v70;
          let _nw229 = Array((new BigNumber(23)).toNumber());
          _nw229[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(((((_1313_v61).contains(p1)) ? ((_1313_v61).get(p1)) : (p3))).multipliedBy(new BigNumber(-157)));
          _nw229[(_dafny.ONE).toNumber()] = p1;
          _nw229[(new BigNumber(2)).toNumber()] = (new BigNumber((_1261_v29).length)).multipliedBy(p3);
          _nw229[(new BigNumber(3)).toNumber()] = p1;
          _nw229[(new BigNumber(4)).toNumber()] = _module.__default.fm9(p1, new BigNumber(81), p2, globalState);
          _nw229[(new BigNumber(5)).toNumber()] = p1;
          _nw229[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(p3, p3);
          _nw229[(new BigNumber(7)).toNumber()] = new BigNumber(-328);
          _nw229[(new BigNumber(8)).toNumber()] = p1;
          _nw229[(new BigNumber(9)).toNumber()] = ((p0) ? (_module.__default.fm9(p1, p3, p2, globalState)) : (new BigNumber((_1316_v62).length)));
          _nw229[(new BigNumber(10)).toNumber()] = new BigNumber((((p2) ? (_1308_v56) : (_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-174))))).length);
          _nw229[(new BigNumber(11)).toNumber()] = (p1).minus(p1);
          _nw229[(new BigNumber(12)).toNumber()] = p3;
          _nw229[(new BigNumber(13)).toNumber()] = p1;
          _nw229[(new BigNumber(14)).toNumber()] = new BigNumber(258);
          _nw229[(new BigNumber(15)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1320_v66,p1)).length)).multipliedBy(p1);
          _nw229[(new BigNumber(16)).toNumber()] = _module.__default.fm9(p1, (_dafny.ZERO).minus(p3), (_this).f19, globalState);
          _nw229[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(p1, new BigNumber(589));
          _nw229[(new BigNumber(18)).toNumber()] = p1;
          _nw229[(new BigNumber(19)).toNumber()] = (new BigNumber(730)).plus(p1);
          _nw229[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), p3);
          _nw229[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1323_v69, _1323_v69)).length);
          _nw229[(new BigNumber(22)).toNumber()] = p1;
          _1324_v70 = _nw229;
          let _index176 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          (_1321_v67)[_index176] = p1;
          let _1325_v71;
          _1325_v71 = new BigNumber(358);
          let _index177 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          let _rhs192 = !(p2);
          let _rhs193 = p3;
          let _rhs194 = _1325_v71;
          let _lhs117 = globalState;
          let _lhs118 = _1321_v67;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          _lhs117.f11 = _rhs192;
          _lhs118[_lhs119] = _rhs193;
          _1325_v71 = _rhs194;
          let _1326_v72;
          _1326_v72 = _module.D4.create_DC13(new BigNumber((_module.__default.fm20(globalState)).length));
          let _1327_v73;
          _1327_v73 = _dafny.Map.Empty.slice().updateUnsafe((_1310_v58)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1310_v58).length))],_1326_v72);
          _1327_v73 = (_1327_v73).update(true, _1326_v72);
          let _1328_v74;
          let _nw230 = new _module.C1();
          _nw230.__ctor(_1321_v67);
          _1328_v74 = _nw230;
          let _1329_v75;
          _1329_v75 = _dafny.Seq.of(_1328_v74);
          let _1330_v76;
          _1330_v76 = _dafny.MultiSet.fromElements(!(p2), p2, !(false), false, (_this).f19);
          let _index178 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          let _index179 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          let _rhs195 = (_dafny.ZERO).minus((p3).multipliedBy(new BigNumber((_dafny.Seq.Concat(_1329_v75, _1329_v75)).length)));
          let _rhs196 = ((_dafny.ZERO).minus((((_1330_v76).contains(p0)) ? ((_1330_v76).get(p0)) : (_1325_v71)))).plus((_1321_v67)[_module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length))]);
          let _lhs120 = _1321_v67;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          let _lhs122 = _1321_v67;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_1321_v67).length));
          _lhs120[_lhs121] = _rhs195;
          _lhs122[_lhs123] = _rhs196;
          r0 = (p3).isEqualTo(_1325_v71);
        } else {
          let _1331_v78;
          _1331_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_1234_v2)).cardinality()),p1);
          let _1332_v79;
          _1332_v79 = _dafny.Map.Empty.slice().updateUnsafe((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))],function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_1331_v78).Keys.Elements) {
              let _1333_v77 = _compr_37;
              if ((_1331_v78).contains(_1333_v77)) {
                _coll37.push([_module.__default.safeDivisionInt(_1333_v77, p1),p3]);
              }
            }
            return _coll37;
          }());
          _1234_v2 = _dafny.Seq.update(_module.__default.fm16(p2, new BigNumber((_1332_v79).length), globalState), _module.__default.safeIndex(p3, new BigNumber((_module.__default.fm16(p2, new BigNumber((_1332_v79).length), globalState)).length)), (_1234_v2)[_module.__default.safeIndex(p1, new BigNumber((_1234_v2).length))]);
          let _1334_v80;
          _1334_v80 = _module.D0.create_DC1(p2);
          let _1335_v81;
          let _nw231 = new _module.C2();
          _nw231.__ctor(_1334_v80, p2);
          _1335_v81 = _nw231;
          let _1336_v82;
          _1336_v82 = _dafny.Seq.of((_1335_v81).f21);
          let _1337_v83;
          _1337_v83 = _dafny.Set.fromElements(_1336_v82, _1336_v82, _1336_v82);
          let _nw232 = new _module.C2();
          _nw232.__ctor(_1334_v80, (_1337_v83).IsSubsetOf(_1337_v83));
          _1335_v81 = _nw232;
          let _1338_v84;
          let _nw233 = new _module.C0();
          _nw233.__ctor(!((_this).f19));
          _1338_v84 = _nw233;
          let _index180 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1232_v0).length));
          (_1232_v0)[_index180] = (_1335_v81).f21;
          let _1339_v85;
          _1339_v85 = new BigNumber(-626);
          let _1340_v86;
          _1340_v86 = _dafny.Set.fromElements(p0);
          let _index181 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1232_v0).length));
          let _rhs197 = (_module.__default.safeDivisionInt(p3, _1339_v85)).isLessThanOrEqualTo(new BigNumber((_1340_v86).length));
          let _rhs198 = p1;
          let _lhs124 = _1232_v0;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1232_v0).length));
          _lhs124[_lhs125] = _rhs197;
          _1339_v85 = _rhs198;
          r0 = ((_dafny.Set.fromElements((_1338_v84).f23)).Union(_1340_v86)).IsSubsetOf((_1340_v86).Difference(_1340_v86));
        }
      }
      let _1341_v87;
      let _nw234 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Set.Empty);
      _1341_v87 = _nw234;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1341_v87).length))) {
        let _1342_i8 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1342_i8)) && ((_1342_i8).isLessThan(new BigNumber((_1341_v87).length))))) {
          (_1341_v87)[(_1342_i8)] = (_module.D11.create_DC29(_dafny.Set.fromElements((_this).f19))).dtor_cf51;
        }
      }
      let _arr0 = (_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))];
      let _index182 = _module.__default.safeIndex(new BigNumber(820), new BigNumber(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))]).length));
      _arr0[_index182] = (new BigNumber((_1234_v2).length)).isEqualTo(p1);
      let _arr1 = (_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))];
      let _index183 = _module.__default.safeIndex(new BigNumber(820), new BigNumber(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))]).length));
      _arr1[_index183] = p2;
      let _hi9 = (new BigNumber(818)).plus(new BigNumber(162));
      for (let _1343_i9 = p3; _1343_i9.isLessThan(_hi9); _1343_i9 = _1343_i9.plus(_dafny.ONE)) {
        let _1344_v88;
        let _nw235 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
        _1344_v88 = _nw235;
        let _1345_v89;
        _1345_v89 = _dafny.Seq.of(p0, _module.__default.fm0(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))])[_module.__default.safeIndex(new BigNumber(820), new BigNumber(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))]).length))], new _dafny.CodePoint('a'.codePointAt(0)), p0, globalState));
        let _index184 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1344_v88).length));
        (_1344_v88)[_index184] = _module.__default.safeModuloInt(new BigNumber((_1345_v89).length), p3);
        let _index185 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_1344_v88).length));
        (_1344_v88)[_index185] = (p1).minus(((p2) ? (p1) : (_1343_i9)));
        let _1346_v90;
        _1346_v90 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1343_i9);
        let _1347_v91;
        _1347_v91 = _dafny.Seq.of(p3);
        let _1348_v92;
        _1348_v92 = _module.D1.create_DC5(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-248)), ((_1349_v1) => function (_1350_i10) {
  return _1349_v1;
})(_1233_v1)), (_1344_v88)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1344_v88).length))], p1);
        let _1351_v93;
        _1351_v93 = _module.D3.create_DC11(_1346_v90, _module.__default.safeModuloInt((_1347_v91)[_module.__default.safeIndex(new BigNumber(((_1348_v92).dtor_cf7).length), new BigNumber((_1347_v91).length))], p1), _1344_v88);
        let _pat_let_tv34 = _1346_v90;
        _1351_v93 = function (_pat_let25_0) {
          return function (_1352_dt__update__tmp_h1) {
            return function (_pat_let26_0) {
              return function (_1353_dt__update_hcf18_h0) {
                return _module.D3.create_DC11(_1353_dt__update_hcf18_h0, (_1352_dt__update__tmp_h1).dtor_cf19, (_1352_dt__update__tmp_h1).dtor_cf20);
              }(_pat_let26_0);
            }(_pat_let_tv34);
          }(_pat_let25_0);
        }(_module.D3.create_DC11(_1346_v90, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,(_1344_v88)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_1344_v88).length))])).length), _1344_v88));
        _1347_v91 = _1347_v91;
        let _1354_v94;
        _1354_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1233_v1);
        _1354_v94 = _1354_v94;
      }
      let _1355_v95;
      _1355_v95 = _dafny.Seq.of((_this).f19);
      let _1356_v96;
      _1356_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1355_v95,p1);
      let _1357_i11;
      _1357_i11 = _dafny.ZERO;
      L8: {
        while (!(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0),p3)).equals(_1356_v96)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1357_i11)) {
              break L8;
            }
            _1357_i11 = (_1357_i11).plus(_dafny.ONE);
            _1355_v95 = _1355_v95;
            let _1358_v97;
            _1358_v97 = _dafny.Set.fromElements(p0);
            let _1359_v98;
            _1359_v98 = _dafny.MultiSet.fromElements((_this).f19);
            let _1360_v99;
            let _nw236 = Array((new BigNumber(9)).toNumber());
            _nw236[(_dafny.ZERO).toNumber()] = p1;
            _nw236[(_dafny.ONE).toNumber()] = p1;
            _nw236[(new BigNumber(2)).toNumber()] = new BigNumber((_1239_v7).cardinality());
            _nw236[(new BigNumber(3)).toNumber()] = _module.__default.fm5(p1, globalState);
            _nw236[(new BigNumber(4)).toNumber()] = p1;
            _nw236[(new BigNumber(5)).toNumber()] = (p3).plus(_module.__default.fm9(new BigNumber((_1358_v97).length), p1, (_this).f19, globalState));
            _nw236[(new BigNumber(6)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(p0)).Union(_1359_v98)).cardinality());
            _nw236[(new BigNumber(7)).toNumber()] = p3;
            _nw236[(new BigNumber(8)).toNumber()] = p1;
            _1360_v99 = _nw236;
            let _index186 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1360_v99).length));
            (_1360_v99)[_index186] = (new BigNumber(655)).minus(p3);
            let _index187 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1360_v99).length));
            (_1360_v99)[_index187] = p3;
            let _1361_v100;
            _1361_v100 = _module.D5.create_DC17(p2, _module.D2.create_DC8((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))], p1, new BigNumber(-273)), true, (_this).f19, p1);
            let _index188 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1360_v99).length));
            let _index189 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1360_v99).length));
            let _rhs199 = !(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))])[_module.__default.safeIndex(new BigNumber(820), new BigNumber(((_1237_v5)[_module.__default.safeIndex(new BigNumber(100), new BigNumber((_1237_v5).length))]).length))]);
            let _rhs200 = (_1361_v100).dtor_cf33;
            let _rhs201 = (((_dafny.ZERO).minus(new BigNumber((_1355_v95).length))).multipliedBy(p1)).minus(p1);
            let _lhs126 = _1360_v99;
            let _lhs127 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_1360_v99).length));
            let _lhs128 = _1360_v99;
            let _lhs129 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1360_v99).length));
            r1 = _rhs199;
            _lhs126[_lhs127] = _rhs200;
            _lhs128[_lhs129] = _rhs201;
            let _1362_v101;
            let _nw237 = new _module.C3();
            _nw237.__ctor();
            _1362_v101 = _nw237;
            let _1363_v102;
            _1363_v102 = _dafny.Seq.of(_1360_v99, _1360_v99);
            _1363_v102 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1363_v102, _1363_v102), _1363_v102);
          }
        }
      }
      r0 = (_this).f19;
      let _1364_v103;
      let _nw238 = new _module.C3();
      _nw238.__ctor();
      _1364_v103 = _nw238;
      let _1365_v104;
      _1365_v104 = _dafny.Seq.of(_1364_v103);
      r1 = _module.__default.fm0((_dafny.MultiSet.fromElements(_1364_v103)).IsSubsetOf(_dafny.MultiSet.FromArray(_1365_v104)), ((true) ? (_1233_v1) : (_module.__default.fm12(p3, _1234_v2, _1234_v2, globalState))), p2, globalState);
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let _1366_v0;
      _1366_v0 = _dafny.MultiSet.fromElements((_this).f19);
      let _1367_v1;
      _1367_v1 = new BigNumber(931);
      let _1368_v2;
      _1368_v2 = _module.D10.create_DC27((_1366_v0).Union((_1366_v0).update(!((_this).f19), _module.__default.abs(_1367_v1))));
      let _1369_v3;
      _1369_v3 = _dafny.Seq.UnicodeFromString("yrd");
      let _1370_v4;
      _1370_v4 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1371_v5;
      _1371_v5 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("wbgn"),_1370_v4);
      let _1372_v6;
      _1372_v6 = _dafny.Seq.of(!(_module.__default.fm0((_this).f19, _1370_v4, (_this).f19, globalState)));
      let _1373_v7;
      let _init40 = ((_1374_v1) => function (_1375_i1) {
        return (_1375_i1).plus(_1374_v1);
      })(_1367_v1);
      let _nw239 = Array((new BigNumber(28)).toNumber());
      for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw239.length); _i0_40++) {
        _nw239[_i0_40] = _init40(new BigNumber(_i0_40));
      }
      _1373_v7 = _nw239;
      let _1376_v8;
      let _nw240 = new _module.C1();
      _nw240.__ctor(_1373_v7);
      _1376_v8 = _nw240;
      let _1377_v9;
      _1377_v9 = _dafny.Seq.of(_1376_v8, _1376_v8);
      let _1378_v10;
      _1378_v10 = _dafny.Set.fromElements((_this).f19);
      let _1379_v11;
      _1379_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1367_v1,_1378_v10);
      let _1380_v12;
      let _nw241 = Array((new BigNumber(20)).toNumber());
      _nw241[(_dafny.ZERO).toNumber()] = _1370_v4;
      _nw241[(_dafny.ONE).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(2)).toNumber()] = (((_this).f19) ? ((((_1371_v5).contains(_dafny.Seq.UnicodeFromString("apnh"))) ? ((_1371_v5).get(_dafny.Seq.UnicodeFromString("apnh"))) : (_1370_v4))) : (_1370_v4));
      _nw241[(new BigNumber(3)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(4)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(5)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(6)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(7)).toNumber()] = _module.__default.fm12(new BigNumber((_1372_v6).length), _dafny.Seq.Create(_module.__default.abs(new BigNumber(223)), function (_1381_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("jxjyg"), globalState);
      _nw241[(new BigNumber(8)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(9)).toNumber()] = (_1369_v3)[_module.__default.safeIndex(_module.__default.fm5(new BigNumber((_1377_v9).length), globalState), new BigNumber((_1369_v3).length))];
      _nw241[(new BigNumber(10)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(11)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(12)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(13)).toNumber()] = _module.__default.fm12(_1367_v1, _dafny.Seq.UnicodeFromString("vnd"), _1369_v3, globalState);
      _nw241[(new BigNumber(14)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(15)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(16)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(17)).toNumber()] = _module.__default.fm12(_1367_v1, _dafny.Seq.UnicodeFromString("op"), _1369_v3, globalState);
      _nw241[(new BigNumber(18)).toNumber()] = _1370_v4;
      _nw241[(new BigNumber(19)).toNumber()] = _module.__default.fm12(new BigNumber((_1379_v11).length), _dafny.Seq.UnicodeFromString("hdoqfle"), _1369_v3, globalState);
      _1380_v12 = _nw241;
      let _rhs202 = _module.D10.create_DC27(_1366_v0);
      let _rhs203 = (((_1366_v0).contains((_this).f19)) ? ((_1366_v0).get((_this).f19)) : (_module.__default.fm5(new BigNumber((_1372_v6).length), globalState)));
      let _rhs204 = _dafny.Seq.Concat(_1369_v3, _1369_v3);
      let _rhs205 = _1380_v12;
      _1368_v2 = _rhs202;
      _1367_v1 = _rhs203;
      _1369_v3 = _rhs204;
      _1380_v12 = _rhs205;
      let _1382_v13;
      _1382_v13 = _module.D10.create_DC28((_this).f19, _1367_v1);
      let _1383_v14;
      _1383_v14 = _dafny.Seq.of((_dafny.ZERO).minus(_1367_v1), (_1382_v13).dtor_cf50, ((_dafny.ZERO).minus(_1367_v1)).multipliedBy(new BigNumber((_1369_v3).length)), _1367_v1);
      _1383_v14 = _dafny.Seq.update((((_this).f19) ? (_1383_v14) : (_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_1383_v14, _module.__default.safeIndex(_1367_v1, new BigNumber((_1383_v14).length)), new BigNumber((_1369_v3).length)), _module.__default.safeIndex(_1367_v1, new BigNumber((_dafny.Seq.update(_1383_v14, _module.__default.safeIndex(_1367_v1, new BigNumber((_1383_v14).length)), new BigNumber((_1369_v3).length))).length)), new BigNumber((_module.__default.fm29(_1370_v4, new BigNumber(420), globalState)).length)), _1383_v14))), _module.__default.safeIndex(_1367_v1, new BigNumber(((((_this).f19) ? (_1383_v14) : (_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_1383_v14, _module.__default.safeIndex(_1367_v1, new BigNumber((_1383_v14).length)), new BigNumber((_1369_v3).length)), _module.__default.safeIndex(_1367_v1, new BigNumber((_dafny.Seq.update(_1383_v14, _module.__default.safeIndex(_1367_v1, new BigNumber((_1383_v14).length)), new BigNumber((_1369_v3).length))).length)), new BigNumber((_module.__default.fm29(_1370_v4, new BigNumber(420), globalState)).length)), _1383_v14)))).length)), _1367_v1);
      let _1384_v15;
      _1384_v15 = _dafny.MultiSet.fromElements(_1367_v1);
      let _1385_v16;
      _1385_v16 = _module.D2.create_DC7(_1384_v15);
      let _pat_let_tv35 = _1367_v1;
      _1367_v1 = (_dafny.ZERO).minus(function (_source14) {
        if (_source14.is_DC8) {
          let _1386___mcc_h0 = (_source14).cf13;
          let _1387___mcc_h1 = (_source14).cf14;
          let _1388___mcc_h2 = (_source14).cf15;
          let _1389_cf15 = _1388___mcc_h2;
          let _1390_cf14 = _1387___mcc_h1;
          let _1391_cf13 = _1386___mcc_h0;
          if ((_this).f19) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("yxsw")).length);
          } else {
            return _1390_cf14;
          }
        } else {
          let _1392___mcc_h3 = (_source14).cf12;
          let _1393_cf12 = _1392___mcc_h3;
          return _pat_let_tv35;
        }
      }(_1385_v16));
      let _1394_i2;
      _1394_i2 = _dafny.ZERO;
      L9: {
        while (false) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1394_i2)) {
              break L9;
            }
            _1394_i2 = (_1394_i2).plus(_dafny.ONE);
            let _rhs206 = ((_this).f19) && ((_this).f19);
            let _rhs207 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeDivisionInt(_1367_v1, _1367_v1)).plus(_1367_v1)));
            let _rhs208 = !((_this).f19) || (!(_module.__default.fm0((_this).f19, new _dafny.CodePoint('l'.codePointAt(0)), (_this).f19, globalState)));
            let _rhs209 = (_this).f19;
            let _rhs210 = (_this).f19;
            let _lhs130 = globalState;
            let _lhs131 = globalState;
            let _lhs132 = globalState;
            let _lhs133 = globalState;
            _lhs130.f11 = _rhs206;
            _1367_v1 = _rhs207;
            _lhs131.f11 = _rhs208;
            _lhs132.f17 = _rhs209;
            _lhs133.f15 = _rhs210;
            let _1395_v17;
            let _nw242 = new _module.C3();
            _nw242.__ctor();
            _1395_v17 = _nw242;
            _1366_v0 = _dafny.MultiSet.fromElements(((_this).f19) === ((_this).f19));
            let _1396_v18;
            let _init41 = function (_1397_i3) {
              return (_this).f19;
            };
            let _nw243 = Array((new BigNumber(11)).toNumber());
            for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw243.length); _i0_41++) {
              _nw243[_i0_41] = _init41(new BigNumber(_i0_41));
            }
            _1396_v18 = _nw243;
            let _1398_v19;
            _1398_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_1367_v1);
            let _1399_v20;
            _1399_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_1398_v19).contains((_this).f19)) ? ((_1398_v19).get((_this).f19)) : (new BigNumber((_dafny.Seq.update(_1369_v3, _module.__default.safeIndex(_1367_v1, new BigNumber((_1369_v3).length)), _1370_v4)).length))),false);
            let _1400_v21;
            _1400_v21 = _dafny.Set.fromElements(_1367_v1);
            let _1401_v22;
            _1401_v22 = _module.D4.create_DC14((_this).f19, _1367_v1, _1400_v21, _1370_v4);
            let _1402_v23;
            _1402_v23 = _dafny.MultiSet.fromElements((_1401_v22).dtor_cf26);
            let _index190 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1396_v18).length));
            (_1396_v18)[_index190] = (((_1399_v20).contains(new BigNumber((_1402_v23).cardinality()))) ? ((_1399_v20).get(new BigNumber((_1402_v23).cardinality()))) : ((_this).f19));
            let _index191 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1396_v18).length));
            (_1396_v18)[_index191] = (_this).fm4((_this).f19, _dafny.Map.Empty.slice().updateUnsafe(_1369_v3,_1398_v19), _1370_v4, _1367_v1, globalState);
          }
        }
      }
      (globalState).f17 = !((_this).f19);
      let _1403_v24;
      _1403_v24 = _module.D9.create_DC24(_1383_v14);
      let _1404_v25;
      _1404_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_module.D9.create_DC26(_1403_v24));
      let _1405_v26;
      let _nw244 = Array((new BigNumber(8)).toNumber());
      _nw244[(_dafny.ZERO).toNumber()] = _1404_v25;
      _nw244[(_dafny.ONE).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(2)).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(3)).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(4)).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(5)).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(6)).toNumber()] = _1404_v25;
      _nw244[(new BigNumber(7)).toNumber()] = (_1404_v25).Merge(_1404_v25);
      _1405_v26 = _nw244;
      let _index192 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_1405_v26).length));
      (_1405_v26)[_index192] = _1404_v25;
      let _index193 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_1405_v26).length));
      (_1405_v26)[_index193] = _1404_v25;
      return;
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let _1406_v0;
      _1406_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,(_this).f19);
      if ((((_this).f19) ? ((((_1406_v0).contains((_this).f19)) ? ((_1406_v0).get((_this).f19)) : ((((_1406_v0).contains(p1)) ? ((_1406_v0).get(p1)) : ((_this).f19))))) : (p1))) {
        let _1407_v1;
        _1407_v1 = _dafny.Seq.of(p1);
        let _1408_v2;
        _1408_v2 = _dafny.Seq.UnicodeFromString("ghvjg");
        let _1409_v3;
        _1409_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_1408_v2).length));
        let _1410_v4;
        _1410_v4 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f19),_1409_v3);
        let _1411_v5;
        _1411_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1408_v2,(((_1410_v4).contains((((_1406_v0).contains((_this).f19)) ? ((_1406_v0).get((_this).f19)) : ((_this).f19)))) ? ((_1410_v4).get((((_1406_v0).contains((_this).f19)) ? ((_1406_v0).get((_this).f19)) : ((_this).f19)))) : (_1409_v3)));
        let _1412_v6;
        _1412_v6 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1413_v7;
        _1413_v7 = _module.D1.create_DC6(p0, new BigNumber(943));
        let _1414_v8;
        _1414_v8 = _module.D11.create_DC30(p1, (_1413_v7).dtor_cf11, (_this).f19);
        let _1415_v9;
        _1415_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1407_v1,p1);
        let _1416_v11;
        _1416_v11 = _dafny.MultiSet.fromElements(_1407_v1);
        let _1417_v14;
        _1417_v14 = _dafny.MultiSet.fromElements(true, (_this).f19);
        let _1418_v15;
        _1418_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1417_v14).cardinality()),(_this).f19);
        let _1419_v16;
        _1419_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1418_v15);
        let _1420_v17;
        _1420_v17 = _dafny.Seq.of(new BigNumber(((((_1419_v16).contains(p0)) ? ((_1419_v16).get(p0)) : (_1418_v15))).length), p0, _module.__default.fm5(new BigNumber((_1417_v14).cardinality()), globalState), (_dafny.ZERO).minus(p0));
        let _1421_v18;
        _1421_v18 = _dafny.Seq.of((_1420_v17)[_module.__default.safeIndex(p0, new BigNumber((_1420_v17).length))], p0, p0);
        let _1422_v20;
        _1422_v20 = _dafny.Seq.of(_1407_v1);
        let _1423_v21;
        let _nw245 = Array((new BigNumber(15)).toNumber());
        _nw245[(_dafny.ZERO).toNumber()] = (((_dafny.Map.Empty.slice().updateUnsafe(_1407_v1,p1)).update(_dafny.Seq.of(p1), (_this).fm4((_this).f19, _1411_v5, _1412_v6, p0, globalState))).update(_1407_v1, (_1414_v8).dtor_cf52)).update(_dafny.Seq.of((_this).f19), (_this).f19);
        _nw245[(_dafny.ONE).toNumber()] = _1415_v9;
        _nw245[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, p1),true);
        _nw245[(new BigNumber(3)).toNumber()] = (function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of (_1416_v11).Elements) {
            let _1424_v10 = _compr_38;
            if ((_1416_v11).contains(_1424_v10)) {
              _coll38.push([_1424_v10,!(p1)]);
            }
          }
          return _coll38;
        }()).Merge(_1415_v9);
        _nw245[(new BigNumber(4)).toNumber()] = (_1415_v9).Merge(_1415_v9);
        _nw245[(new BigNumber(5)).toNumber()] = _1415_v9;
        _nw245[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm27(p0, p0, p0, globalState),p1);
        _nw245[(new BigNumber(7)).toNumber()] = _1415_v9;
        _nw245[(new BigNumber(8)).toNumber()] = function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of (_dafny.Seq.of(_dafny.Seq.of((_this).f19, !(p1)))).Elements) {
            let _1425_v12 = _compr_39;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of((_this).f19, !(p1))), _1425_v12)) {
              _coll39.push([_1425_v12,(_this).f19]);
            }
          }
          return _coll39;
        }();
        _nw245[(new BigNumber(9)).toNumber()] = _1415_v9;
        _nw245[(new BigNumber(10)).toNumber()] = function () {
          let _coll40 = new _dafny.Map();
          for (const _compr_40 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),(_this).f19)).Keys.Elements) {
            let _1426_v13 = _compr_40;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p1),(_this).f19)).contains(_1426_v13)) {
              _coll40.push([_1426_v13,p1]);
            }
          }
          return _coll40;
        }();
        _nw245[(new BigNumber(11)).toNumber()] = _module.__default.fm30(_1421_v18, globalState);
        _nw245[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1407_v1,_module.__default.fm0((_this).f19, _1412_v6, p1, globalState))).Merge(_1415_v9);
        _nw245[(new BigNumber(13)).toNumber()] = function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of (_1422_v20).Elements) {
            let _1427_v19 = _compr_41;
            if (_dafny.Seq.contains(_1422_v20, _1427_v19)) {
              _coll41.push([_1427_v19,(_this).f19]);
            }
          }
          return _coll41;
        }();
        _nw245[(new BigNumber(14)).toNumber()] = (_1415_v9).Merge(_1415_v9);
        _1423_v21 = _nw245;
        let _index194 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1423_v21).length));
        (_1423_v21)[_index194] = _1415_v9;
        let _index195 = _module.__default.safeIndex(new BigNumber(336), new BigNumber((_1423_v21).length));
        (_1423_v21)[_index195] = _1415_v9;
        if (_dafny.Seq.IsProperPrefixOf(_1407_v1, _1407_v1)) {
          let _1428_v22;
          let _nw246 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _1428_v22 = _nw246;
          let _1429_v23;
          _1429_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1428_v22,(_1420_v17)[_module.__default.safeIndex(new BigNumber(-865), new BigNumber((_1420_v17).length))]);
          _1429_v23 = (_1429_v23).update(_1428_v22, p0);
          let _1430_v24;
          _1430_v24 = new BigNumber(47);
          let _1431_v25;
          _1431_v25 = _dafny.MultiSet.fromElements(_module.D10.create_DC27(_1417_v14));
          _1430_v24 = (((_1431_v25).contains(_module.__default.fm31(true, (_1421_v18)[_module.__default.safeIndex((_dafny.ZERO).minus(_1430_v24), new BigNumber((_1421_v18).length))], globalState))) ? ((_1431_v25).get(_module.__default.fm31(true, (_1421_v18)[_module.__default.safeIndex((_dafny.ZERO).minus(_1430_v24), new BigNumber((_1421_v18).length))], globalState))) : (_1430_v24));
          let _1432_v26;
          let _nw247 = Array((new BigNumber(5)).toNumber()).fill(false);
          _1432_v26 = _nw247;
          let _1433_v27;
          _1433_v27 = _dafny.MultiSet.fromElements(_1432_v26);
          let _1434_v28;
          _1434_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1433_v27,(_dafny.ZERO).minus((_module.__default.fm5(p0, globalState)).multipliedBy(_1430_v24)));
          _1434_v28 = (_1434_v28).update(_1433_v27, new BigNumber(592));
          let _1435_v29;
          _1435_v29 = _module.D9.create_DC26(_module.D9.create_DC24(_1421_v18));
          let _1436_v30;
          let _nw248 = new _module.C0();
          _nw248.__ctor(p1);
          _1436_v30 = _nw248;
          let _1437_v31;
          _1437_v31 = _module.D9.create_DC26(_module.D9.create_DC26(_module.D9.create_DC25(_1430_v24, (_this).f19, _1430_v24, _1436_v30)));
          let _pat_let_tv36 = _1437_v31;
          let _1438_v32;
          let _nw249 = Array((new BigNumber(9)).toNumber());
          _nw249[(_dafny.ZERO).toNumber()] = _1435_v29;
          _nw249[(_dafny.ONE).toNumber()] = _module.D9.create_DC26(_module.D9.create_DC26(_1437_v31));
          _nw249[(new BigNumber(2)).toNumber()] = _1435_v29;
          _nw249[(new BigNumber(3)).toNumber()] = _1435_v29;
          _nw249[(new BigNumber(4)).toNumber()] = _module.D9.create_DC26(_module.D9.create_DC24(_1420_v17));
          _nw249[(new BigNumber(5)).toNumber()] = ((!(p1)) ? (_1435_v29) : (_1435_v29));
          _nw249[(new BigNumber(6)).toNumber()] = _1435_v29;
          _nw249[(new BigNumber(7)).toNumber()] = _1435_v29;
          _nw249[(new BigNumber(8)).toNumber()] = function (_pat_let27_0) {
            return function (_1439_dt__update__tmp_h0) {
              return function (_pat_let28_0) {
                return function (_1440_dt__update_hcf47_h0) {
                  return _module.D9.create_DC26(_1440_dt__update_hcf47_h0);
                }(_pat_let28_0);
              }(_pat_let_tv36);
            }(_pat_let27_0);
          }(_1435_v29);
          _1438_v32 = _nw249;
          let _index196 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1438_v32).length));
          (_1438_v32)[_index196] = _1435_v29;
          let _1441_v33;
          _1441_v33 = _module.D4.create_DC14((_this).f19, p0, _module.__default.fm18((_1436_v30).f23, new BigNumber((_1420_v17).length), new BigNumber(-821), globalState), _1412_v6);
          let _1442_v34;
          _1442_v34 = _dafny.Set.fromElements((_1420_v17)[_module.__default.safeIndex(_1430_v24, new BigNumber((_1420_v17).length))]);
          let _1443_v35;
          _1443_v35 = _dafny.Set.fromElements((_1441_v33).dtor_cf25, _dafny.Set.fromElements(new BigNumber((_1408_v2).length)), _1442_v34, _1442_v34, _dafny.Set.fromElements((_1414_v8).dtor_cf53, p0));
          let _index197 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1438_v32).length));
          let _rhs211 = p0;
          let _rhs212 = _module.D9.create_DC26(_1437_v31);
          let _rhs213 = !((_1443_v35).equals(_1443_v35));
          let _lhs134 = _1438_v32;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_1438_v32).length));
          let _lhs136 = globalState;
          _1430_v24 = _rhs211;
          _lhs134[_lhs135] = _rhs212;
          _lhs136.f15 = _rhs213;
          let _index198 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1428_v22).length));
          (_1428_v22)[_index198] = _module.__default.safeDivisionInt(_1430_v24, new BigNumber((_1407_v1).length));
          let _1444_v36;
          _1444_v36 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)));
          let _1445_v37;
          _1445_v37 = _module.D4.create_DC13((_dafny.ZERO).minus(new BigNumber((_1408_v2).length)));
          let _1446_v38;
          _1446_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1445_v37,(_dafny.ZERO).minus(_1430_v24));
          let _index199 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1428_v22).length));
          let _rhs214 = ((_1436_v30).f23) || ((_1444_v36).IsProperSubsetOf(_1444_v36));
          let _rhs215 = (_1436_v30).f23;
          let _rhs216 = (((_1446_v38).contains(_1445_v37)) ? ((_1446_v38).get(_1445_v37)) : (_1430_v24));
          let _lhs137 = globalState;
          let _lhs138 = globalState;
          let _lhs139 = _1428_v22;
          let _lhs140 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_1428_v22).length));
          _lhs137.f17 = _rhs214;
          _lhs138.f17 = _rhs215;
          _lhs139[_lhs140] = _rhs216;
        } else {
          let _1447_v39;
          _1447_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1414_v8,!(p1));
          let _1448_v40;
          _1448_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1407_v1,(_1447_v39).update(_1414_v8, p1));
          let _1449_v41;
          let _nw250 = Array((new BigNumber(20)).toNumber());
          _nw250[(_dafny.ZERO).toNumber()] = !((_this).f19);
          _nw250[(_dafny.ONE).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(2)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(3)).toNumber()] = !((_this).f19) || (!(!((_this).f19)));
          _nw250[(new BigNumber(4)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(5)).toNumber()] = !(!(_1448_v40).contains(_1407_v1));
          _nw250[(new BigNumber(6)).toNumber()] = p1;
          _nw250[(new BigNumber(7)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(8)).toNumber()] = p1;
          _nw250[(new BigNumber(9)).toNumber()] = p1;
          _nw250[(new BigNumber(10)).toNumber()] = (((_1406_v0).contains(p1)) ? ((_1406_v0).get(p1)) : (true));
          _nw250[(new BigNumber(11)).toNumber()] = true;
          _nw250[(new BigNumber(12)).toNumber()] = p1;
          _nw250[(new BigNumber(13)).toNumber()] = (((_this).fm4((_this).fm4((_this).f19, _1411_v5, _1412_v6, (_dafny.ZERO).minus(p0), globalState), _dafny.Map.Empty.slice().updateUnsafe(_1408_v2,_1409_v3), _1412_v6, new BigNumber((_dafny.Seq.of(_1412_v6)).length), globalState)) ? ((_this).f19) : (true));
          _nw250[(new BigNumber(14)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(15)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(16)).toNumber()] = (_this).f19;
          _nw250[(new BigNumber(17)).toNumber()] = (((_1418_v15).contains(p0)) ? ((_1418_v15).get(p0)) : (false));
          _nw250[(new BigNumber(18)).toNumber()] = p1;
          _nw250[(new BigNumber(19)).toNumber()] = p1;
          _1449_v41 = _nw250;
          let _1450_v42;
          _1450_v42 = _module.D0.create_DC1((_this).f19);
          let _1451_v43;
          let _nw251 = new _module.C2();
          _nw251.__ctor(_1450_v42, p1);
          _1451_v43 = _nw251;
          let _1452_v44;
          _1452_v44 = _module.D12.create_DC31(_dafny.Set.fromElements(_1451_v43, _1451_v43));
          let _index200 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_1449_v41).length));
          (_1449_v41)[_index200] = !((_1452_v44).dtor_cf55).contains(_1451_v43);
          let _1453_v45;
          _1453_v45 = new BigNumber(579);
          let _1454_v46;
          let _init42 = ((_1455_v45) => function (_1456_i0) {
            return _module.__default.safeModuloInt(_1456_i0, _1455_v45);
          })(_1453_v45);
          let _nw252 = Array((new BigNumber(8)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw252.length); _i0_42++) {
            _nw252[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1454_v46 = _nw252;
          let _1457_v47;
          _1457_v47 = _module.D3.create_DC11(_dafny.Map.Empty.slice().updateUnsafe((_this).f19,new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)), (_dafny.ZERO).minus(_1453_v45), _1454_v46);
          let _index201 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_1449_v41).length));
          let _rhs217 = _1420_v17;
          let _rhs218 = (_1407_v1)[_module.__default.safeIndex((_1457_v47).dtor_cf19, new BigNumber((_1407_v1).length))];
          let _rhs219 = p0;
          let _lhs141 = _1449_v41;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(987), new BigNumber((_1449_v41).length));
          _1420_v17 = _rhs217;
          _lhs141[_lhs142] = _rhs218;
          _1453_v45 = _rhs219;
          let _1458_v48;
          _1458_v48 = _dafny.Set.fromElements(_1453_v45);
          _1453_v45 = new BigNumber(((_1458_v48).Difference(_1458_v48)).length);
          _1453_v45 = (_1453_v45).multipliedBy(new BigNumber((((!(p1)) ? (_1408_v2) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_1459_v6) => function (_1460_i1) {
            return _1459_v6;
          })(_1412_v6))))).length));
          (globalState).f17 = !((_1449_v41)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_1449_v41).length))]);
          let _1461_v49;
          let _nw253 = new _module.C2();
          _nw253.__ctor(_1450_v42, (_1449_v41)[_module.__default.safeIndex(new BigNumber(987), new BigNumber((_1449_v41).length))]);
          _1461_v49 = _nw253;
        }
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1408_v2, _module.__default.fm16((_this).f19, p0, globalState)), _1408_v2)) {
          let _1462_v50;
          let _init43 = function (_1463_i2) {
            return (_1463_i2).minus(new BigNumber(-313));
          };
          let _nw254 = Array((new BigNumber(7)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw254.length); _i0_43++) {
            _nw254[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1462_v50 = _nw254;
          let _index202 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length));
          (_1462_v50)[_index202] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(626)), ((_1464_v6) => function (_1465_i3) {
            return _1464_v6;
          })(_1412_v6))).length);
          let _index203 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1462_v50).length));
          (_1462_v50)[_index203] = new BigNumber((_1408_v2).length);
          let _1466_v51;
          _1466_v51 = new BigNumber(-819);
          let _1467_v52;
          _1467_v52 = _dafny.Set.fromElements(p0, p0);
          let _1468_v53;
          _1468_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1412_v6,(_this).f19);
          let _index204 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length));
          let _index205 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1462_v50).length));
          let _rhs220 = new BigNumber((_1406_v0).length);
          let _rhs221 = _dafny.Seq.Concat(_1408_v2, _module.__default.fm16(_module.__default.fm0((_this).f19, _1412_v6, (_this).f19, globalState), p0, globalState));
          let _rhs222 = p0;
          let _rhs223 = _module.__default.safeDivisionInt((new BigNumber((_1467_v52).length)).plus(_1466_v51), new BigNumber(((_1468_v53).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1412_v6,!(!(!(p1)))))).length));
          let _lhs143 = _1462_v50;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length));
          let _lhs145 = _1462_v50;
          let _lhs146 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1462_v50).length));
          _lhs143[_lhs144] = _rhs220;
          _1408_v2 = _rhs221;
          _lhs145[_lhs146] = _rhs222;
          _1466_v51 = _rhs223;
          let _index206 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length));
          (_1462_v50)[_index206] = (((_1466_v51).isLessThanOrEqualTo((_1462_v50)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length))])) ? (_1466_v51) : (_module.__default.fm5(p0, globalState)));
          let _nw255 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _1462_v50 = _nw255;
          let _1469_v54;
          _1469_v54 = _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("iteqm"), p0, (_1462_v50)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length))]);
          let _1470_v55;
          _1470_v55 = _dafny.Set.fromElements(_1408_v2, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fofwaytk"), _1408_v2), (_1469_v54).dtor_cf7);
          _1470_v55 = _dafny.Set.fromElements(_1408_v2, _1408_v2, _dafny.Seq.UnicodeFromString("frcpgrc"));
          let _1471_v56;
          _1471_v56 = _module.D2.create_DC7(_dafny.MultiSet.FromArray(_1420_v17));
          let _1472_v57;
          _1472_v57 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_1412_v6);
          let _1473_v58;
          _1473_v58 = _module.D4.create_DC14(p1, _1466_v51, _1467_v52, _1412_v6);
          _1471_v56 = _module.__default.fm32(p1, (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1472_v57).length), new BigNumber(905))), (_1462_v50)[_module.__default.safeIndex(new BigNumber(241), new BigNumber((_1462_v50).length))], _module.__default.fm0(!((_this).f19), (_1473_v58).dtor_cf26, p1, globalState), globalState);
        } else {
          let _1474_v59;
          _1474_v59 = new BigNumber(-135);
          _1474_v59 = p0;
          let _1475_v60;
          _1475_v60 = _dafny.Set.fromElements(_1474_v59);
          let _1476_v61;
          _1476_v61 = _module.D4.create_DC14(true, _1474_v59, _1475_v60, _1412_v6);
          let _1477_v62;
          _1477_v62 = _dafny.Set.fromElements(_module.__default.fm0(!(p1), (_1476_v61).dtor_cf26, p1, globalState), (_this).f19, p1);
          let _1478_v63;
          let _nw256 = Array((new BigNumber(2)).toNumber());
          _nw256[(_dafny.ZERO).toNumber()] = p0;
          _nw256[(_dafny.ONE).toNumber()] = new BigNumber((_1477_v62).length);
          _1478_v63 = _nw256;
          let _1479_v64;
          _1479_v64 = _dafny.Seq.of((((_this).f19) ? (_1478_v63) : (_1478_v63)), _1478_v63, _1478_v63);
          let _1480_v65;
          _1480_v65 = _module.D3.create_DC10(new BigNumber((_1418_v15).length));
          let _1481_v66;
          _1481_v66 = _dafny.MultiSet.fromElements(_1480_v65, _1480_v65);
          let _rhs224 = _dafny.Seq.of(_1478_v63, _1478_v63, _1478_v63, _1478_v63, _1478_v63);
          let _rhs225 = new BigNumber(((((_this).f19) ? (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("je"), _1408_v2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_1482_i4) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })))).length);
          let _rhs226 = ((p1) ? ((_1407_v1)[_module.__default.safeIndex((_1420_v17)[_module.__default.safeIndex((((_1481_v66).contains(_1480_v65)) ? ((_1481_v66).get(_1480_v65)) : (_1474_v59)), new BigNumber((_1420_v17).length))], new BigNumber((_1407_v1).length))]) : (p1));
          let _lhs147 = globalState;
          _1479_v64 = _rhs224;
          _1474_v59 = _rhs225;
          _lhs147.f17 = _rhs226;
          _1474_v59 = ((!_dafny.areEqual(_1408_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(26)), ((_1483_v6) => function (_1484_i5) {
            return _1483_v6;
          })(_1412_v6)))) ? (new BigNumber(((_1406_v0).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f19,true))).length)) : (new BigNumber(-108)));
          let _1485_v67;
          _1485_v67 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(!(p1), (_this).f19)).IsSubsetOf(_module.__default.fm28(new BigNumber((_1408_v2).length), true, globalState)),_module.__default.fm33((_this).f19, true, globalState));
          let _1486_v68;
          _1486_v68 = _module.D10.create_DC28(p1, _1474_v59);
          _1485_v67 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1486_v68);
          _1412_v6 = _1412_v6;
        }
        (globalState).f15 = (_this).f19;
        let _1487_v69;
        _1487_v69 = _module.D0.create_DC1((_this).f19);
        let _1488_v70;
        let _nw257 = new _module.C2();
        _nw257.__ctor(_1487_v69, p1);
        _1488_v70 = _nw257;
        let _1489_v71;
        _1489_v71 = _module.D12.create_DC31((_dafny.Set.fromElements(_1488_v70, _1488_v70)).Difference(_dafny.Set.fromElements(_1488_v70)));
        _1489_v71 = _1489_v71;
      } else {
        let _1490_v72;
        _1490_v72 = new _dafny.CodePoint('r'.codePointAt(0));
        _1490_v72 = (((_this).f19) ? (_1490_v72) : (_1490_v72));
        let _1491_v73;
        _1491_v73 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _1492_v74;
        let _nw258 = Array((new BigNumber(14)).toNumber()).fill(false);
        _1492_v74 = _nw258;
        let _1493_v75;
        _1493_v75 = _module.D1.create_DC4(p0, (_this).f19, _1492_v74);
        let _1494_v76;
        _1494_v76 = _dafny.Seq.UnicodeFromString("ird");
        let _1495_v77;
        _1495_v77 = _dafny.Map.Empty.slice().updateUnsafe((_1493_v75).dtor_cf5,_1494_v76);
        _1491_v73 = (_1491_v73).update(false, new BigNumber((_1495_v77).length));
        _1406_v0 = (_1406_v0).update(p1, (!(_module.__default.fm0((_this).f19, _1490_v72, p1, globalState))) || (false));
        r0 = ((_this).f19) || ((_this).f19);
        let _1496_v78;
        _1496_v78 = _module.D0.create_DC1(true);
        let _1497_v79;
        let _nw259 = new _module.C2();
        _nw259.__ctor(_1496_v78, (_this).f19);
        _1497_v79 = _nw259;
        let _1498_v80;
        _1498_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-116),_1497_v79);
        let _1499_v81;
        _1499_v81 = _dafny.Set.fromElements(_1498_v80);
        let _1500_v82;
        _1500_v82 = _dafny.MultiSet.fromElements(p0);
        let _1501_v83;
        _1501_v83 = _module.D8.create_DC23(_1500_v82);
        let _1502_v84;
        _1502_v84 = _dafny.Seq.of((_1497_v79).f21, p1);
        let _1503_v85;
        _1503_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1497_v79).f21);
        let _1504_v86;
        let _nw260 = new _module.C3();
        _nw260.__ctor();
        _1504_v86 = _nw260;
        let _1505_v87;
        _1505_v87 = _dafny.Seq.of(_1501_v83, _module.D8.create_DC23((_module.__default.fm23((_1502_v84)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1502_v84).length))], ((_dafny.Map.Empty.slice().updateUnsafe((_1497_v79).f21,new BigNumber((_1494_v76).length))).update((_1497_v79).f21, new BigNumber((_1503_v85).length))).update(!((_1497_v79).f21), p0), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1504_v86)).length)), _1490_v72, globalState)).update(new BigNumber(-748), _module.__default.abs(p0))), _module.D8.create_DC23(_1500_v82), _1501_v83);
        let _1506_v88;
        _1506_v88 = _dafny.Map.Empty.slice().updateUnsafe((_1499_v81).IsSubsetOf(_1499_v81),_1505_v87);
        _1506_v88 = (_1506_v88).Merge(_1506_v88);
      }
      let _1507_v89;
      _1507_v89 = new _dafny.CodePoint('y'.codePointAt(0));
      let _1508_v90;
      _1508_v90 = new BigNumber(975);
      let _1509_v91;
      let _nw261 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1509_v91 = _nw261;
      let _1510_v92;
      _1510_v92 = _module.D12.create_DC32(_1508_v90);
      let _1511_v93;
      _1511_v93 = _dafny.Seq.of((_this).f19, _module.__default.fm0((_this).f19, _1507_v89, (_this).f19, globalState), (_this).f19, false, true);
      let _1512_v94;
      _1512_v94 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1508_v90);
      let _1513_v95;
      _1513_v95 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(p1)).length));
      let _1514_v96;
      _1514_v96 = _dafny.Seq.of((p0).multipliedBy(_1508_v90), p0, (_1510_v92).dtor_cf56, (((_1511_v93)[_module.__default.safeIndex((((_1512_v94).contains((_this).f19)) ? ((_1512_v94).get((_this).f19)) : (new BigNumber(589))), new BigNumber((_1511_v93).length))]) ? (p0) : (_1508_v90)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(389)), ((_1515_v90) => function (_1516_i6) {
        return _1515_v90;
      })(_1508_v90)), _1513_v95)).length));
      let _1517_v97;
      let _nw262 = Array((new BigNumber(2)).toNumber());
      _nw262[(_dafny.ZERO).toNumber()] = _1511_v93;
      _nw262[(_dafny.ONE).toNumber()] = _1511_v93;
      _1517_v97 = _nw262;
      let _1518_v98;
      _1518_v98 = _module.D1.create_DC3(_1517_v97);
      let _1519_v99;
      _1519_v99 = _dafny.MultiSet.fromElements(_1518_v98);
      let _1520_v100;
      _1520_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1519_v99);
      let _1521_v101;
      _1521_v101 = _dafny.Map.Empty.slice().updateUnsafe(_1520_v100,_module.__default.safeDivisionInt(p0, _1508_v90));
      let _1522_v102;
      _1522_v102 = _dafny.Map.Empty.slice().updateUnsafe(!((_dafny.ZERO).minus(p0)).isEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(p1, p1)).length), _1508_v90)).length)),_1509_v91);
      let _1523_v103;
      _1523_v103 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1512_v94).length),_1513_v95);
      let _rhs227 = _1507_v89;
      let _rhs228 = (((_1521_v101).contains(_1520_v100)) ? ((_1521_v101).get(_1520_v100)) : (p0));
      let _rhs229 = (((_1522_v102).contains((_this).f19)) ? ((_1522_v102).get((_this).f19)) : (_1509_v91));
      let _rhs230 = p1;
      let _rhs231 = _dafny.Seq.Concat((((_1523_v103).contains(_1508_v90)) ? ((_1523_v103).get(_1508_v90)) : (_1514_v96)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_1524_p0) => function (_1525_i7) {
        return _1524_p0;
      })(p0)));
      _1507_v89 = _rhs227;
      _1508_v90 = _rhs228;
      _1509_v91 = _rhs229;
      r0 = _rhs230;
      _1514_v96 = _rhs231;
      let _1526_v104;
      _1526_v104 = _dafny.Seq.of(_1406_v0, (_1406_v0).Merge(_1406_v0));
      let _1527_v105;
      _1527_v105 = _module.D13.create_DC36(_dafny.Seq.of(_1406_v0));
      _1526_v104 = _dafny.Seq.update((((p0).isLessThan(p0)) ? ((_1527_v105).dtor_cf63) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_1528_v0) => function (_1529_i8) {
        return _1528_v0;
      })(_1406_v0)))), _module.__default.safeIndex((_1508_v90).plus((_dafny.ZERO).minus(_1508_v90)), new BigNumber(((((p0).isLessThan(p0)) ? ((_1527_v105).dtor_cf63) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_1530_v0) => function (_1531_i8) {
        return _1530_v0;
      })(_1406_v0))))).length)), _1406_v0);
      let _1532_v106;
      _1532_v106 = _module.D0.create_DC1((_this).f19);
      let _1533_v107;
      let _nw263 = new _module.C2();
      _nw263.__ctor(((!((_this).f19)) ? (_1532_v106) : (_1532_v106)), ((!(p1)) ? (true) : ((_this).f19)));
      _1533_v107 = _nw263;
      let _1534_v108;
      let _nw264 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _1534_v108 = _nw264;
      let _1535_v109;
      _1535_v109 = _module.D13.create_DC38(_1507_v89, (_1533_v107).f21, _1508_v90, _1534_v108);
      let _1536_v110;
      let _nw265 = Array((new BigNumber(5)).toNumber());
      _nw265[(_dafny.ZERO).toNumber()] = _1535_v109;
      _nw265[(_dafny.ONE).toNumber()] = _1535_v109;
      _nw265[(new BigNumber(2)).toNumber()] = _1535_v109;
      _nw265[(new BigNumber(3)).toNumber()] = _1535_v109;
      _nw265[(new BigNumber(4)).toNumber()] = (((_1533_v107).f21) ? (_module.D13.create_DC38(_1507_v89, p1, _1508_v90, _1534_v108)) : (_1535_v109));
      _1536_v110 = _nw265;
      let _1537_v111;
      _1537_v111 = _module.D14.create_DC40(_1536_v110);
      _1536_v110 = (_1537_v111).dtor_cf71;
      let _1538_i9;
      _1538_i9 = _dafny.ZERO;
      L10: {
        while (!(_dafny.Map.Empty.slice().updateUnsafe(_1508_v90,_1508_v90)).contains(_module.__default.safeModuloInt(p0, _1508_v90))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1538_i9)) {
              break L10;
            }
            _1538_i9 = (_1538_i9).plus(_dafny.ONE);
            let _1539_v112;
            _1539_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1507_v89,_1507_v89);
            let _1540_v113;
            _1540_v113 = _dafny.Seq.of(_1507_v89, new _dafny.CodePoint('k'.codePointAt(0)), _1507_v89, _1507_v89, new _dafny.CodePoint('e'.codePointAt(0)));
            _1539_v112 = (_1539_v112).update((_1540_v113)[_module.__default.safeIndex(new BigNumber(14), new BigNumber((_1540_v113).length))], _1507_v89);
            let _rhs232 = !(false) || ((_1533_v107).f21);
            let _rhs233 = _1533_v107;
            let _lhs148 = globalState;
            _lhs148.f17 = _rhs232;
            _1533_v107 = _rhs233;
            let _index207 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_1534_v108).length));
            (_1534_v108)[_index207] = p0;
            let _index208 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1509_v91).length));
            (_1509_v91)[_index208] = p1;
            let _1541_v114;
            _1541_v114 = _dafny.Set.fromElements(_1507_v89, _1507_v89, _1507_v89);
            let _index209 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_1534_v108).length));
            let _index210 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1509_v91).length));
            let _rhs234 = ((!((_1541_v114).IsSubsetOf(_1541_v114))) ? (p0) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lxxgp"), _1540_v113), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lxxgp"), _1540_v113)).length)), _1507_v89)).length)));
            let _rhs235 = (_this).f19;
            let _lhs149 = _1534_v108;
            let _lhs150 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_1534_v108).length));
            let _lhs151 = _1509_v91;
            let _lhs152 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_1509_v91).length));
            _lhs149[_lhs150] = _rhs234;
            _lhs151[_lhs152] = _rhs235;
            _1508_v90 = new BigNumber((_dafny.Seq.Concat(_1540_v113, _1540_v113)).length);
          }
        }
      }
      r0 = (_1533_v107.f20).dtor_cf1;
      return r0;
    }
    get f19() {
      let _this = this;
      return _this._f19;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f18) {
      let _this = this;
      (_this)._f18 = f18;
      return;
    }
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1542_v0;
      _1542_v0 = true;
      let _1543_v1;
      _1543_v1 = _dafny.Seq.of(_1542_v0);
      if (_dafny.Seq.IsProperPrefixOf(_1543_v1, _1543_v1)) {
        let _1544_v2;
        let _init44 = ((_1545_v1) => function (_1546_i0) {
          return _1545_v1;
        })(_1543_v1);
        let _nw266 = Array((new BigNumber(26)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw266.length); _i0_44++) {
          _nw266[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1544_v2 = _nw266;
        let _1547_v3;
        _1547_v3 = _module.D1.create_DC3(_1544_v2);
        _1544_v2 = (_1547_v3).dtor_cf3;
        let _1548_v4;
        let _init45 = function (_1549_i1) {
          return _module.D1.create_DC5(_dafny.Seq.UnicodeFromString("odrkr"), (_this).f18, new BigNumber((_dafny.Seq.UnicodeFromString("tbdrju")).length));
        };
        let _nw267 = Array((new BigNumber(25)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw267.length); _i0_45++) {
          _nw267[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1548_v4 = _nw267;
        let _1550_v5;
        _1550_v5 = _dafny.Seq.UnicodeFromString("bnr");
        let _index211 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1548_v4).length));
        (_1548_v4)[_index211] = _module.D1.create_DC5(_1550_v5, (_this).f18, (_this).f18);
        let _1551_v6;
        _1551_v6 = _dafny.Seq.of(new BigNumber(797));
        let _index212 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1548_v4).length));
        let _rhs236 = _module.__default.fm1((_this).f18, globalState);
        let _rhs237 = (_dafny.MultiSet.fromElements((_this).f18)).IsProperSubsetOf(_dafny.MultiSet.FromArray(_1551_v6));
        let _rhs238 = (p0).IsDisjointFrom(p0);
        let _rhs239 = (_1543_v1)[_module.__default.safeIndex(((_this).f18).multipliedBy(new BigNumber((_1551_v6).length)), new BigNumber((_1543_v1).length))];
        let _lhs153 = _1548_v4;
        let _lhs154 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_1548_v4).length));
        let _lhs155 = globalState;
        let _lhs156 = globalState;
        _lhs153[_lhs154] = _rhs236;
        _1542_v0 = _rhs237;
        _lhs155.f15 = _rhs238;
        _lhs156.f11 = _rhs239;
        let _1552_v7;
        _1552_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
        _1552_v7 = (_1552_v7).update((_this).f18, (new BigNumber(166)).minus((_this).f18));
        let _1553_v8;
        let _nw268 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1553_v8 = _nw268;
        let _1554_v9;
        _1554_v9 = _module.D1.create_DC6((_this).f18, _module.__default.safeDivisionInt((_this).f18, (_this).f18));
        let _1555_v10;
        _1555_v10 = _module.D0.create_DC1(_1542_v0);
        let _1556_v11;
        _1556_v11 = _module.D0.create_DC2(_1555_v10);
        let _rhs240 = _1553_v8;
        let _rhs241 = ((false) ? (((_this).f18).multipliedBy(new BigNumber((_1543_v1).length))) : ((_this).f18));
        let _rhs242 = (_this).f18;
        let _rhs243 = _module.__default.fm2(globalState);
        let _rhs244 = _1556_v11;
        _1553_v8 = _rhs240;
        r0 = _rhs241;
        r0 = _rhs242;
        _1554_v9 = _rhs243;
        _1556_v11 = _rhs244;
        let _1557_v12;
        _1557_v12 = _dafny.Seq.of(_1550_v5, _1550_v5, _1550_v5, _1550_v5, _1550_v5);
        if (!(!_dafny.areEqual(_1557_v12, _1557_v12))) {
          let _1558_v13;
          _1558_v13 = _dafny.Seq.of(_dafny.Seq.of(_1542_v0));
          let _1559_v14;
          _1559_v14 = _dafny.MultiSet.fromElements(_1542_v0);
          let _1560_v15;
          _1560_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v0,new BigNumber((_1559_v14).cardinality()));
          let _1561_v16;
          _1561_v16 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(228)), ((_1562_v5) => function (_1563_i2) {
            return (_1562_v5)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_1562_v5).length))];
          })(_1550_v5))).length));
          let _1564_v17;
          let _nw269 = Array((new BigNumber(23)).toNumber());
          _nw269[(_dafny.ZERO).toNumber()] = (_this).f18;
          _nw269[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((_1558_v13)[_module.__default.safeIndex((_this).f18, new BigNumber((_1558_v13).length))]).length), new BigNumber((_1559_v14).cardinality()));
          _nw269[(new BigNumber(3)).toNumber()] = new BigNumber((p0).length);
          _nw269[(new BigNumber(4)).toNumber()] = ((_this).f18).plus((_this).f18);
          _nw269[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(6)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1551_v6).length), new BigNumber(-560));
          _nw269[(new BigNumber(8)).toNumber()] = new BigNumber(871);
          _nw269[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).f18);
          _nw269[(new BigNumber(10)).toNumber()] = new BigNumber(231);
          _nw269[(new BigNumber(11)).toNumber()] = ((_dafny.ZERO).minus((_this).f18)).multipliedBy((_this).f18);
          _nw269[(new BigNumber(12)).toNumber()] = (((_1560_v15).contains(_1542_v0)) ? ((_1560_v15).get(_1542_v0)) : ((_this).f18));
          _nw269[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_this).f18);
          _nw269[(new BigNumber(14)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(15)).toNumber()] = new BigNumber((((_1561_v16).update(new BigNumber((_1552_v7).length), _module.__default.abs(new BigNumber(655)))).Intersect(_1561_v16)).cardinality());
          _nw269[(new BigNumber(16)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(17)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(18)).toNumber()] = _module.__default.fm3((_1551_v6)[_module.__default.safeIndex((_this).f18, new BigNumber((_1551_v6).length))], (_this).f18, _1542_v0, globalState);
          _nw269[(new BigNumber(19)).toNumber()] = (_module.__default.fm3((_this).f18, (_this).f18, _1542_v0, globalState)).multipliedBy((_this).f18);
          _nw269[(new BigNumber(20)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(21)).toNumber()] = (_this).f18;
          _nw269[(new BigNumber(22)).toNumber()] = new BigNumber((_1550_v5).length);
          _1564_v17 = _nw269;
          let _1565_v18;
          _1565_v18 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1566_v19;
          _1566_v19 = _dafny.MultiSet.fromElements(_1550_v5, _dafny.Seq.of(_1565_v18));
          let _index213 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1564_v17).length));
          (_1564_v17)[_index213] = ((((_1566_v19).contains(_1550_v5)) ? ((_1566_v19).get(_1550_v5)) : ((_this).f18))).plus((_this).f18);
          let _index214 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length));
          (_1553_v8)[_index214] = _1542_v0;
          let _1567_v20;
          _1567_v20 = _module.D0.create_DC1(_1542_v0);
          let _index215 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1564_v17).length));
          let _index216 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length));
          let _rhs245 = (_dafny.ZERO).minus((_this).f18);
          let _rhs246 = _1542_v0;
          let _rhs247 = _1567_v20;
          let _lhs157 = _1564_v17;
          let _lhs158 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_1564_v17).length));
          let _lhs159 = _1553_v8;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length));
          _lhs157[_lhs158] = _rhs245;
          _lhs159[_lhs160] = _rhs246;
          _1567_v20 = _rhs247;
          (globalState).f11 = _dafny.areEqual(((_1542_v0) ? (_1543_v1) : (_dafny.Seq.of(true, _1542_v0))), _1543_v1);
          let _1568_v21;
          let _nw270 = new _module.C0();
          _nw270.__ctor(_1542_v0);
          _1568_v21 = _nw270;
          _1550_v5 = _dafny.Seq.Concat(_1550_v5, _1550_v5);
          let _1569_v22;
          let _nw271 = new _module.C2();
          _nw271.__ctor(_module.D0.create_DC1(_1542_v0), (((_1568_v21).f23) ? ((_1553_v8)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length))]) : ((_1553_v8)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length))])));
          _1569_v22 = _nw271;
          let _nw272 = new _module.C4();
          _nw272.__ctor((_1553_v8)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_1553_v8).length))]);
          _1569_v22 = _nw272;
        } else {
          (globalState).f17 = (p0).IsSubsetOf(p0);
          let _1570_v23;
          _1570_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1542_v0),_1542_v0);
          r0 = _module.__default.fm9((_this).f18, _module.__default.fm5(new BigNumber((_1551_v6).length), globalState), (((_1570_v23).contains(_module.__default.fm20(globalState))) ? ((_1570_v23).get(_module.__default.fm20(globalState))) : (true)), globalState);
          let _1571_v24;
          _1571_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1542_v0);
          let _1572_v25;
          _1572_v25 = new _dafny.CodePoint('n'.codePointAt(0));
          (globalState).f15 = ((_module.__default.fm0((((_1571_v24).contains((_this).f18)) ? ((_1571_v24).get((_this).f18)) : (_1542_v0)), _1572_v25, _1542_v0, globalState)) ? (_1542_v0) : (_module.__default.fm0(_1542_v0, new _dafny.CodePoint('q'.codePointAt(0)), _1542_v0, globalState)));
          _1542_v0 = ((_this).f18).isLessThan((_dafny.ZERO).minus((_this).f18));
          (globalState).f16 = _1543_v1;
        }
      } else {
        if (_1542_v0) {
          let _1573_v26;
          _1573_v26 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1574_v27;
          let _init46 = ((_1575_v0) => function (_1576_i3) {
            return _1575_v0;
          })(_1542_v0);
          let _nw273 = Array((new BigNumber(13)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw273.length); _i0_46++) {
            _nw273[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1574_v27 = _nw273;
          let _index217 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1574_v27).length));
          (_1574_v27)[_index217] = _1542_v0;
          let _index218 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1574_v27).length));
          let _rhs248 = new _dafny.CodePoint('o'.codePointAt(0));
          let _rhs249 = !(_1542_v0);
          let _lhs161 = _1574_v27;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(67), new BigNumber((_1574_v27).length));
          _1573_v26 = _rhs248;
          _lhs161[_lhs162] = _rhs249;
          r0 = (_this).f18;
          r0 = (_this).f18;
          _1573_v26 = _1573_v26;
          let _1577_v28;
          _1577_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f18, (_this).f18),(_dafny.ZERO).minus((_this).f18));
          let _1578_v29;
          let _init47 = ((_1579_v0, _1580_p0, _1581_v26) => function (_1582_i4) {
            return (_module.D4.create_DC14(_1579_v0, (_this).f18, _1580_p0, _1581_v26)).dtor_cf26;
          })(_1542_v0, p0, _1573_v26);
          let _nw274 = Array((new BigNumber(18)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw274.length); _i0_47++) {
            _nw274[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1578_v29 = _nw274;
          let _1583_v30;
          _1583_v30 = _dafny.Seq.of(_1578_v29, _1578_v29);
          _1577_v28 = (_1577_v28).update(_module.__default.safeModuloInt((_this).f18, new BigNumber((_1583_v30).length)), (_this).f18);
        } else {
          let _1584_v31;
          _1584_v31 = _dafny.Seq.UnicodeFromString("xm");
          let _1585_v32;
          let _nw275 = Array((new BigNumber(5)).toNumber());
          _nw275[(_dafny.ZERO).toNumber()] = _1542_v0;
          _nw275[(_dafny.ONE).toNumber()] = _1542_v0;
          _nw275[(new BigNumber(2)).toNumber()] = _1542_v0;
          _nw275[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1584_v31, _dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_1586_i5) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }));
          _nw275[(new BigNumber(4)).toNumber()] = !((_1543_v1)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f18), new BigNumber((_1543_v1).length))]) || (_1542_v0);
          _1585_v32 = _nw275;
          let _index219 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1585_v32).length));
          (_1585_v32)[_index219] = _1542_v0;
          let _index220 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1585_v32).length));
          (_1585_v32)[_index220] = true;
          let _1587_v33;
          _1587_v33 = _dafny.MultiSet.fromElements((_this).f18);
          let _1588_v34;
          _1588_v34 = _module.D2.create_DC7(_1587_v33);
          let _1589_v35;
          _1589_v35 = _dafny.Seq.of(_1588_v34, _1588_v34, _1588_v34, _1588_v34);
          let _index221 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1585_v32).length));
          (_1585_v32)[_index221] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), function (_1590_i6) {
            return _module.D2.create_DC7(_dafny.MultiSet.fromElements((_this).f18));
          }), _1589_v35);
          r0 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1584_v31, _1584_v31)).length));
          let _1591_v36;
          _1591_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
          let _1592_v37;
          let _nw276 = Array((new BigNumber(14)).toNumber());
          _nw276[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1584_v31, _dafny.Seq.UnicodeFromString("kkptvhk"))).length);
          _nw276[(_dafny.ONE).toNumber()] = (_this).f18;
          _nw276[(new BigNumber(2)).toNumber()] = (_this).f18;
          _nw276[(new BigNumber(3)).toNumber()] = new BigNumber((_1591_v36).length);
          _nw276[(new BigNumber(4)).toNumber()] = (_this).f18;
          _nw276[(new BigNumber(5)).toNumber()] = (_this).f18;
          _nw276[(new BigNumber(6)).toNumber()] = (new BigNumber((_1584_v31).length)).multipliedBy(new BigNumber(151));
          _nw276[(new BigNumber(7)).toNumber()] = ((_this).f18).multipliedBy((_this).f18);
          _nw276[(new BigNumber(8)).toNumber()] = new BigNumber((_1591_v36).length);
          _nw276[(new BigNumber(9)).toNumber()] = new BigNumber(453);
          _nw276[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ddlyqlq")).length);
          _nw276[(new BigNumber(11)).toNumber()] = (_this).f18;
          _nw276[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt((_this).f18, (_this).f18);
          _nw276[(new BigNumber(13)).toNumber()] = (_this).f18;
          _1592_v37 = _nw276;
          let _index222 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1592_v37).length));
          (_1592_v37)[_index222] = (_this).f18;
          let _index223 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_1592_v37).length));
          (_1592_v37)[_index223] = new BigNumber(476);
          let _1593_v38;
          _1593_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1585_v32)[_module.__default.safeIndex(new BigNumber(131), new BigNumber((_1585_v32).length))],new BigNumber(-546));
          let _1594_v39;
          _1594_v39 = _module.D3.create_DC11(_1593_v38, (_this).f18, _1592_v37);
          _1594_v39 = _module.D3.create_DC11(_1593_v38, (_this).f18, _1592_v37);
        }
        r0 = (_this).f18;
        let _1595_v40;
        _1595_v40 = _dafny.MultiSet.fromElements(_1542_v0);
        r0 = _module.__default.safeDivisionInt(_module.__default.fm5(new BigNumber(((_module.D10.create_DC27(_1595_v40)).dtor_cf48).cardinality()), globalState), (_this).f18);
        let _1596_v41;
        _1596_v41 = _dafny.Seq.UnicodeFromString("datkd");
        _1596_v41 = _dafny.Seq.Concat(_1596_v41, _1596_v41);
        let _1597_v42;
        let _nw277 = new _module.C4();
        _nw277.__ctor(_1542_v0);
        _1597_v42 = _nw277;
        let _1598_v43;
        _1598_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1597_v42,(_this).f18);
        let _1599_v44;
        _1599_v44 = _module.D0.create_DC1(_1542_v0);
        let _1600_v45;
        let _nw278 = new _module.C2();
        _nw278.__ctor(_1599_v44, _1542_v0);
        _1600_v45 = _nw278;
        let _1601_v46;
        let _nw279 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
        _1601_v46 = _nw279;
        let _1602_v47;
        _1602_v47 = _module.D14.create_DC41(_1600_v45, _1596_v41, _1601_v46);
        let _1603_v48;
        _1603_v48 = _module.D14.create_DC41((_1602_v47).dtor_cf72, _1596_v41, _1601_v46);
        let _1604_v49;
        _1604_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1603_v48,(_this).f18);
        let _1605_v50;
        _1605_v50 = _dafny.Seq.of((_this).f18);
        let _1606_v51;
        _1606_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v0,new BigNumber((_dafny.Seq.UnicodeFromString("jdetgprt")).length));
        let _1607_v52;
        _1607_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1596_v41,(_1606_v51).update((_1597_v42).f19, (_this).f18));
        let _1608_v53;
        _1608_v53 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1609_v54;
        let _nw280 = Array((new BigNumber(20)).toNumber());
        _nw280[(_dafny.ZERO).toNumber()] = _module.__default.fm5((_this).f18, globalState);
        _nw280[(_dafny.ONE).toNumber()] = (((_1598_v43).contains(_1597_v42)) ? ((_1598_v43).get(_1597_v42)) : ((_this).f18));
        _nw280[(new BigNumber(2)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(3)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(4)).toNumber()] = ((_this).f18).plus((_this).f18);
        _nw280[(new BigNumber(5)).toNumber()] = new BigNumber(((p0).Intersect(p0)).length);
        _nw280[(new BigNumber(6)).toNumber()] = (((_1604_v49).contains(_1602_v47)) ? ((_1604_v49).get(_1602_v47)) : ((_this).f18));
        _nw280[(new BigNumber(7)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(8)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(9)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(10)).toNumber()] = _module.__default.fm5(new BigNumber(28), globalState);
        _nw280[(new BigNumber(11)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(12)).toNumber()] = (new BigNumber(667)).plus((_this).f18);
        _nw280[(new BigNumber(13)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(14)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(15)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(16)).toNumber()] = new BigNumber((_module.__default.fm19(_1605_v50, _1542_v0, globalState)).length);
        _nw280[(new BigNumber(17)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(18)).toNumber()] = (_this).f18;
        _nw280[(new BigNumber(19)).toNumber()] = (((_1597_v42).fm4((_1597_v42).f19, _1607_v52, _1608_v53, (_this).f18, globalState)) ? ((_dafny.ZERO).minus((_this).f18)) : ((_this).f18));
        _1609_v54 = _nw280;
        let _index224 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1609_v54).length));
        (_1609_v54)[_index224] = new BigNumber(611);
        let _index225 = _module.__default.safeIndex(new BigNumber(546), new BigNumber((_1609_v54).length));
        (_1609_v54)[_index225] = (_this).f18;
      }
      let _1610_v55;
      _1610_v55 = _dafny.Seq.UnicodeFromString("ftlcnhg");
      let _1611_v56;
      _1611_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_this).f18);
      let _1612_v57;
      _1612_v57 = _module.D0.create_DC1(_1542_v0);
      let _1613_v58;
      _1613_v58 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1614_v59;
      _1614_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1612_v57,_1613_v58);
      let _1615_v60;
      _1615_v60 = _dafny.Map.Empty.slice().updateUnsafe((_1543_v1)[_module.__default.safeIndex((_this).f18, new BigNumber((_1543_v1).length))],(_this).f18);
      let _1616_v61;
      _1616_v61 = _dafny.MultiSet.fromElements(_1542_v0);
      let _1617_v62;
      _1617_v62 = _dafny.Set.fromElements(_1616_v61, _1616_v61);
      let _1618_v63;
      let _nw281 = Array((new BigNumber(23)).toNumber());
      _nw281[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_1619_i7) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length);
      _nw281[(_dafny.ONE).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(2)).toNumber()] = new BigNumber((_1610_v55).length);
      _nw281[(new BigNumber(3)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(4)).toNumber()] = new BigNumber(306);
      _nw281[(new BigNumber(5)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(6)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(7)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_this).f18);
      _nw281[(new BigNumber(9)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(10)).toNumber()] = new BigNumber(((_1611_v56).update((_this).f18, (_this).f18)).length);
      _nw281[(new BigNumber(11)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(12)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1614_v59).length));
      _nw281[(new BigNumber(14)).toNumber()] = (((_1615_v60).contains(false)) ? ((_1615_v60).get(false)) : ((_this).f18));
      _nw281[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1617_v62).length));
      _nw281[(new BigNumber(16)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(17)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(18)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(19)).toNumber()] = new BigNumber((_1610_v55).length);
      _nw281[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f18);
      _nw281[(new BigNumber(21)).toNumber()] = (_this).f18;
      _nw281[(new BigNumber(22)).toNumber()] = (_this).f18;
      _1618_v63 = _nw281;
      let _1620_v64;
      _1620_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1618_v63,_1610_v55);
      _1620_v64 = _1620_v64;
      let _1621_i8;
      _1621_i8 = _dafny.ZERO;
      L11: {
        while (_1542_v0) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1621_i8)) {
              break L11;
            }
            _1621_i8 = (_1621_i8).plus(_dafny.ONE);
            let _1622_v65;
            _1622_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1542_v0);
            let _1623_v66;
            _1623_v66 = _dafny.Set.fromElements((_module.__default.fm17(_1542_v0, _1542_v0, globalState)).update((_this).f18, _module.__default.fm0(_1542_v0, new _dafny.CodePoint('i'.codePointAt(0)), _1542_v0, globalState)), _1622_v65);
            let _source15 = _module.D3.create_DC9(_1623_v66);
            if (_source15.is_DC10) {
              let _1624___mcc_h0 = (_source15).cf17;
              let _1625_cf17 = _1624___mcc_h0;
              (globalState).f11 = ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_1625_cf17)))).isLessThanOrEqualTo((_this).f18);
              let _1626_v67;
              let _nw282 = Array((new BigNumber(17)).toNumber()).fill(false);
              _1626_v67 = _nw282;
              _1626_v67 = _1626_v67;
              let _index226 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1626_v67).length));
              (_1626_v67)[_index226] = _1542_v0;
              let _1627_v68;
              let _nw283 = Array((new BigNumber(11)).toNumber());
              _nw283[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(508)), ((_1628_v58) => function (_1629_i9) {
                return _1628_v58;
              })(_1613_v58));
              _nw283[(_dafny.ONE).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("mk");
              _nw283[(new BigNumber(3)).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(4)).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(5)).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("qpit");
              _nw283[(new BigNumber(7)).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(8)).toNumber()] = _1610_v55;
              _nw283[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_1610_v55, _module.__default.safeIndex(_1625_cf17, new BigNumber((_1610_v55).length)), _1613_v58);
              _nw283[(new BigNumber(10)).toNumber()] = _1610_v55;
              _1627_v68 = _nw283;
              let _1630_v69;
              let _nw284 = Array((new BigNumber(17)).toNumber());
              _nw284[(_dafny.ZERO).toNumber()] = _1627_v68;
              _nw284[(_dafny.ONE).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(2)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(3)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(4)).toNumber()] = ((_1542_v0) ? (_1627_v68) : (_1627_v68));
              _nw284[(new BigNumber(5)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(6)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(7)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(8)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(9)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(10)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(11)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(12)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(13)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(14)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(15)).toNumber()] = _1627_v68;
              _nw284[(new BigNumber(16)).toNumber()] = ((_1542_v0) ? (_1627_v68) : (_1627_v68));
              _1630_v69 = _nw284;
              let _index227 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_1630_v69).length));
              (_1630_v69)[_index227] = _1627_v68;
              let _index228 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1626_v67).length));
              let _index229 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_1630_v69).length));
              let _rhs250 = (_1543_v1)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f18, (_this).f18), new BigNumber((_1543_v1).length))];
              let _rhs251 = _1627_v68;
              let _lhs163 = _1626_v67;
              let _lhs164 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_1626_v67).length));
              let _lhs165 = _1630_v69;
              let _lhs166 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_1630_v69).length));
              _lhs163[_lhs164] = _rhs250;
              _lhs165[_lhs166] = _rhs251;
              _1611_v56 = (_1611_v56).update(_1625_cf17, new BigNumber((_1610_v55).length));
            } else if (_source15.is_DC11) {
              let _1631___mcc_h1 = (_source15).cf18;
              let _1632___mcc_h2 = (_source15).cf19;
              let _1633___mcc_h3 = (_source15).cf20;
              let _1634_cf20 = _1633___mcc_h3;
              let _1635_cf19 = _1632___mcc_h2;
              let _1636_cf18 = _1631___mcc_h1;
              _1635_cf19 = _module.__default.fm9((_this).f18, (_this).f18, _1542_v0, globalState);
              let _1637_v70;
              _1637_v70 = _dafny.MultiSet.fromElements(_1635_cf19);
              let _1638_v71;
              _1638_v71 = _dafny.Set.fromElements(_1637_v70);
              (globalState).f15 = (((_1622_v65).contains(((_this).f18).plus(_1635_cf19))) ? ((_1622_v65).get(((_this).f18).plus(_1635_cf19))) : ((_1638_v71).IsSubsetOf(_1638_v71)));
              _1613_v58 = _1613_v58;
              _1542_v0 = _1542_v0;
            } else {
              let _1639___mcc_h4 = (_source15).cf16;
              let _1640_cf16 = _1639___mcc_h4;
              (globalState).f17 = _1542_v0;
              let _1641_v72;
              let _nw285 = new _module.C2();
              _nw285.__ctor(_1612_v57, _1542_v0);
              _1641_v72 = _nw285;
              let _1642_v73;
              let _nw286 = Array((new BigNumber(14)).toNumber());
              _nw286[(_dafny.ZERO).toNumber()] = _1641_v72;
              _nw286[(_dafny.ONE).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(2)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(3)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(4)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(5)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(6)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(7)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(8)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(9)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(10)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(11)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(12)).toNumber()] = _1641_v72;
              _nw286[(new BigNumber(13)).toNumber()] = _1641_v72;
              _1642_v73 = _nw286;
              let _index230 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1642_v73).length));
              (_1642_v73)[_index230] = _1641_v72;
              let _index231 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1642_v73).length));
              (_1642_v73)[_index231] = _1641_v72;
              (globalState).f15 = ((_1542_v0) ? (_1542_v0) : (_1542_v0));
              let _index232 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1618_v63).length));
              (_1618_v63)[_index232] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1610_v55, _1610_v55)).length));
              let _index233 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1618_v63).length));
              (_1618_v63)[_index233] = _module.__default.safeModuloInt((_this).f18, (_this).f18);
            }
            _1618_v63 = _1618_v63;
            let _1643_v74;
            let _nw287 = Array((new BigNumber(13)).toNumber());
            _nw287[(_dafny.ZERO).toNumber()] = _1542_v0;
            _nw287[(_dafny.ONE).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(2)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(3)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(4)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(5)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(6)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(7)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(8)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(9)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(10)).toNumber()] = _1542_v0;
            _nw287[(new BigNumber(11)).toNumber()] = !(_1542_v0);
            _nw287[(new BigNumber(12)).toNumber()] = _1542_v0;
            _1643_v74 = _nw287;
            let _1644_v75;
            _1644_v75 = _dafny.Map.Empty.slice().updateUnsafe((((_1616_v61).contains(_1542_v0)) ? ((_1616_v61).get(_1542_v0)) : ((_this).f18)),_1643_v74);
            let _1645_v76;
            let _nw288 = Array((new BigNumber(29)).toNumber());
            _1645_v76 = _nw288;
            let _1646_v77;
            let _nw289 = new _module.C4();
            _nw289.__ctor(_1542_v0);
            _1646_v77 = _nw289;
            let _index234 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1645_v76).length));
            (_1645_v76)[_index234] = _1646_v77;
            let _1647_v78;
            _1647_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1542_v0,_1613_v58);
            let _index235 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1645_v76).length));
            let _rhs252 = (_1644_v75).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f18,_1643_v74));
            let _rhs253 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm16(_module.__default.fm0(_1542_v0, (((_1647_v78).contains((_1646_v77).f19)) ? ((_1647_v78).get((_1646_v77).f19)) : (_1613_v58)), (_1646_v77).f19, globalState), (_this).f18, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), ((_1648_v58) => function (_1649_i10) {
              return _1648_v58;
            })(_1613_v58)));
            let _rhs254 = (_dafny.ZERO).minus((_this).f18);
            let _rhs255 = _1646_v77;
            let _lhs167 = _1645_v76;
            let _lhs168 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1645_v76).length));
            _1644_v75 = _rhs252;
            _1542_v0 = _rhs253;
            r0 = _rhs254;
            _lhs167[_lhs168] = _rhs255;
            let _1650_v79;
            _1650_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.D4.create_DC13((_this).f18),_1542_v0);
            (globalState).f15 = !(!((p0).Union(p0)).contains(new BigNumber((_1650_v79).length)));
          }
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1618_v63).length))) {
        let _1651_i11 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1651_i11)) && ((_1651_i11).isLessThan(new BigNumber((_1618_v63).length))))) {
          (_1618_v63)[(_1651_i11)] = _module.__default.safeModuloInt(_1651_i11, (_module.D11.create_DC30(false, (_this).f18, _1542_v0)).dtor_cf53);
        }
      }
      _1610_v55 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("up"), _dafny.Seq.UnicodeFromString("calngaa"));
      (globalState).f17 = !(!(_1542_v0)) || (_1542_v0);
      r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f18), (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f18, (_this).f18)));
      return r0;
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
