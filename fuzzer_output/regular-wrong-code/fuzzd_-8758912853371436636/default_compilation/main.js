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
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(617)), function (_0_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("ocmloxxh")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nwbfxj"), _dafny.Seq.UnicodeFromString("ctmkds")));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return (new BigNumber((((false) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(428),true)) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-698),!(false))))).length)).isEqualTo(new BigNumber(((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(907), new BigNumber(423))) {
            let _1_v0 = _compr_1;
            if (((new BigNumber(907)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(423)))) {
              _coll1.push([_module.__default.safeModuloInt(_1_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(849), new BigNumber(974))).cardinality())),new _dafny.CodePoint('v'.codePointAt(0))]);
            }
          }
          return _coll1;
        }()).Keys.Elements) {
          let _2_v1 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(907), new BigNumber(423))) {
              let _1_v0 = _compr_2;
              if (((new BigNumber(907)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(423)))) {
                _coll2.push([_module.__default.safeModuloInt(_1_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(849), new BigNumber(974))).cardinality())),new _dafny.CodePoint('v'.codePointAt(0))]);
              }
            }
            return _coll2;
          }()).contains(_2_v1)) {
            _coll0.add(_module.__default.safeModuloInt(_2_v1, new BigNumber(47)));
          }
        }
        return _coll0;
      }()).Union(_dafny.Set.fromElements(new BigNumber(956), new BigNumber(701)))).length));
    };
    static fm3(p0, globalState) {
      if ((new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_3_i0) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).Elements) {
          let _4_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_3_i0) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }), _4_v0)) {
            _coll3.push([_4_v0,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("vaujf")).length))).length)]);
          }
        }
        return _coll3;
      }()).length)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("uat")).length))))) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      } else if (false) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }
    };
    static fm4(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(-324)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(!(false)))).length),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),!(false))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-421), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(532),false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("rnwxeunx")).length))).length), new BigNumber(-846), new BigNumber(246), new BigNumber(938)))).length)));
    };
    static fm5(p0, globalState) {
      return new BigNumber((_dafny.Seq.Concat(((!(true)) ? (_dafny.Seq.of(true, false, true, true, false)) : (_dafny.Seq.of(true))), _dafny.Seq.of(true))).length);
    };
    static m0(p0, p1, globalState) {
      let _5_i0;
      _5_i0 = _dafny.ZERO;
      L0: {
        while (p1) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_5_i0)) {
              break L0;
            }
            _5_i0 = (_5_i0).plus(_dafny.ONE);
            let _6_v0;
            let _nw0 = Array((new BigNumber(12)).toNumber()).fill(false);
            _6_v0 = _nw0;
            _6_v0 = _6_v0;
            let _7_v1;
            _7_v1 = new BigNumber(35);
            let _8_v2;
            let _init0 = function (_9_i1) {
              return new _dafny.CodePoint('j'.codePointAt(0));
            };
            let _nw1 = Array((new BigNumber(9)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
              _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _8_v2 = _nw1;
            let _10_v3;
            _10_v3 = _dafny.Map.Empty.slice().updateUnsafe(_7_v1,_dafny.Seq.of(_8_v2));
            let _11_v4;
            _11_v4 = _dafny.Map.Empty.slice().updateUnsafe(_7_v1,p1);
            let _12_v5;
            _12_v5 = _dafny.Seq.of(_8_v2, _8_v2);
            let _13_v6;
            let _nw2 = new _module.C0();
            _nw2.__ctor((((_10_v3).contains((_dafny.ZERO).minus(_module.__default.fm5((((_11_v4).contains((_dafny.ZERO).minus(_7_v1))) ? ((_11_v4).get((_dafny.ZERO).minus(_7_v1))) : (p0)), globalState)))) ? ((_10_v3).get((_dafny.ZERO).minus(_module.__default.fm5((((_11_v4).contains((_dafny.ZERO).minus(_7_v1))) ? ((_11_v4).get((_dafny.ZERO).minus(_7_v1))) : (p0)), globalState)))) : (_12_v5)));
            _13_v6 = _nw2;
            let _14_v7;
            _14_v7 = true;
            let _15_v8;
            _15_v8 = _dafny.Seq.UnicodeFromString("k");
            _14_v7 = !((new BigNumber((_15_v8).length)).isLessThanOrEqualTo(_7_v1));
            _14_v7 = p1;
          }
        }
      }
      let _16_v9;
      _16_v9 = _dafny.Seq.UnicodeFromString("peut");
      let _17_v10;
      let _nw3 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _17_v10 = _nw3;
      let _18_v11;
      _18_v11 = _dafny.Seq.of(_17_v10, _17_v10, _17_v10, _17_v10, _17_v10);
      let _19_v12;
      let _nw4 = new _module.C0();
      _nw4.__ctor(_18_v11);
      _19_v12 = _nw4;
      let _20_v13;
      _20_v13 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-877)), function (_21_i2) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), _16_v9),_19_v12);
      _20_v13 = (_20_v13).update(p0, _19_v12);
      let _22_v14;
      _22_v14 = _module.D2.create_DC9(p1);
      let _pat_let_tv0 = _16_v9;
      let _pat_let_tv1 = _16_v9;
      let _pat_let_tv2 = _16_v9;
      let _pat_let_tv3 = _16_v9;
      _16_v9 = function (_source0) {
        if (_source0.is_DC7) {
          let _23___mcc_h0 = (_source0).cf11;
          let _24___mcc_h1 = (_source0).cf12;
          let _25___mcc_h2 = (_source0).cf13;
          let _26___mcc_h3 = (_source0).cf14;
          let _27_cf14 = _26___mcc_h3;
          let _28_cf13 = _25___mcc_h2;
          let _29_cf12 = _24___mcc_h1;
          let _30_cf11 = _23___mcc_h0;
          return _dafny.Seq.UnicodeFromString("ba");
        } else if (_source0.is_DC8) {
          let _31___mcc_h4 = (_source0).cf15;
          let _32___mcc_h5 = (_source0).cf16;
          let _33_cf16 = _32___mcc_h5;
          let _34_cf15 = _31___mcc_h4;
          return _pat_let_tv0;
        } else if (_source0.is_DC9) {
          let _35___mcc_h6 = (_source0).cf17;
          let _36_cf17 = _35___mcc_h6;
          return _dafny.Seq.Concat(_pat_let_tv1, _pat_let_tv2);
        } else if (_source0.is_DC6) {
          let _37___mcc_h7 = (_source0).cf10;
          let _38_cf10 = _37___mcc_h7;
          return _dafny.Seq.Concat(_38_cf10, _pat_let_tv3);
        } else {
          let _39___mcc_h8 = (_source0).cf18;
          let _40_cf18 = _39___mcc_h8;
          return _dafny.Seq.UnicodeFromString("eohfqvsl");
        }
      }(_22_v14);
      _16_v9 = _16_v9;
      if (!(!(p0) || (p0)) || (p0)) {
        let _41_v15;
        _41_v15 = true;
        let _42_v16;
        _42_v16 = new BigNumber(348);
        let _43_v17;
        _43_v17 = _dafny.MultiSet.fromElements(_42_v16);
        _41_v15 = (_43_v17).contains((new BigNumber((_16_v9).length)).multipliedBy(_42_v16));
        _41_v15 = p1;
        _41_v15 = _41_v15;
        let _44_v18;
        let _nw5 = new _module.C0();
        _nw5.__ctor(_18_v11);
        _44_v18 = _nw5;
        let _45_v19;
        _45_v19 = _module.D1.create_DC3(_44_v18);
        let _46_v20;
        _46_v20 = _module.D1.create_DC5(_45_v19);
        let _47_v21;
        _47_v21 = _dafny.Map.Empty.slice().updateUnsafe(_41_v15,_dafny.Seq.of(_46_v20));
        let _48_v22;
        _48_v22 = _dafny.Seq.of(_module.D1.create_DC5(_45_v19), _46_v20, _46_v20, _module.D1.create_DC5(_45_v19), _46_v20);
        let _49_v23;
        _49_v23 = _dafny.Seq.of(p1, p1);
        let _50_v24;
        _50_v24 = _dafny.Seq.of((((_47_v21).contains(!(p0))) ? ((_47_v21).get(!(p0))) : (_48_v22)), _dafny.Seq.Concat(_dafny.Seq.of(_46_v20), _48_v22), (((_49_v23)[_module.__default.safeIndex(_42_v16, new BigNumber((_49_v23).length))]) ? (_48_v22) : (_48_v22)));
        _42_v16 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_50_v24)[_module.__default.safeIndex(_module.__default.fm5(_41_v15, globalState), new BigNumber((_50_v24).length))]).length)));
        _44_v18 = _44_v18;
      } else {
        let _51_v25;
        _51_v25 = new BigNumber(534);
        let _52_v26;
        _52_v26 = true;
        let _53_v27;
        _53_v27 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
        let _rhs0 = _51_v25;
        let _rhs1 = (((p0) || ((((_53_v27).contains(_52_v26)) ? ((_53_v27).get(_52_v26)) : (p0)))) ? (_51_v25) : (_51_v25));
        let _rhs2 = true;
        _51_v25 = _rhs0;
        _51_v25 = _rhs1;
        _52_v26 = _rhs2;
        if (((_52_v26) ? (!(p1) || (p1)) : ((_module.__default.fm5(_52_v26, globalState)).isLessThan(_51_v25)))) {
          let _54_v28;
          let _nw6 = Array((new BigNumber(22)).toNumber());
          _54_v28 = _nw6;
          let _index0 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_54_v28).length));
          (_54_v28)[_index0] = _19_v12;
          let _index1 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_54_v28).length));
          (_54_v28)[_index1] = _19_v12;
          let _55_v29;
          _55_v29 = _dafny.Seq.of(_20_v13, _20_v13);
          _51_v25 = new BigNumber((_55_v29).length);
          _51_v25 = _51_v25;
          _52_v26 = p0;
          let _56_v30;
          let _init1 = ((_57_v25) => function (_58_i3) {
            return _module.__default.safeDivisionInt(_58_i3, _57_v25);
          })(_51_v25);
          let _nw7 = Array((new BigNumber(3)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw7.length); _i0_1++) {
            _nw7[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _56_v30 = _nw7;
          _56_v30 = _56_v30;
        } else {
          _52_v26 = p0;
          let _59_v31;
          let _nw8 = new _module.C0();
          _nw8.__ctor(_dafny.Seq.Concat(_18_v11, _18_v11));
          _59_v31 = _nw8;
          let _60_v32;
          _60_v32 = new _dafny.CodePoint('u'.codePointAt(0));
          let _61_v33;
          _61_v33 = _dafny.MultiSet.fromElements(_60_v32);
          _51_v25 = (((_61_v33).contains(_60_v32)) ? ((_61_v33).get(_60_v32)) : ((_59_v31).fm1(p0, _52_v26, true, p1, globalState)));
          _19_v12 = _19_v12;
          let _62_v34;
          let _nw9 = new _module.C0();
          _nw9.__ctor(_dafny.Seq.of(_17_v10));
          _62_v34 = _nw9;
          _62_v34 = _62_v34;
        }
        let _63_v35;
        _63_v35 = _dafny.Set.fromElements(false);
        let _64_v36;
        _64_v36 = _module.D2.create_DC7(_16_v9, new BigNumber((_63_v35).length), _51_v25, !(true));
        let _65_v37;
        _65_v37 = _dafny.Map.Empty.slice().updateUnsafe(_64_v36,_52_v26);
        _65_v37 = _dafny.Map.Empty.slice().updateUnsafe(_64_v36,_52_v26);
        let _66_v38;
        _66_v38 = new _dafny.CodePoint('t'.codePointAt(0));
        let _67_v39;
        _67_v39 = _module.D0.create_DC1(_52_v26, _module.__default.fm5(_52_v26, globalState), _51_v25, false, _51_v25);
        let _68_v40;
        _68_v40 = _dafny.Map.Empty.slice().updateUnsafe(_66_v38,(_67_v39).dtor_cf3);
        _68_v40 = (_68_v40).update(new _dafny.CodePoint('h'.codePointAt(0)), _51_v25);
        if (_52_v26) {
          _51_v25 = (_51_v25).multipliedBy((new BigNumber((_16_v9).length)).minus(new BigNumber(436)));
          _66_v38 = _66_v38;
          let _69_v41;
          let _nw10 = new _module.C0();
          _nw10.__ctor(_18_v11);
          _69_v41 = _nw10;
          let _70_v42;
          _70_v42 = _dafny.Set.fromElements(_51_v25);
          let _71_v43;
          _71_v43 = _dafny.Seq.of(_69_v41, _69_v41);
          let _72_v44;
          let _nw11 = Array((new BigNumber(15)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_51_v25, (_dafny.ZERO).minus(_51_v25)));
          _nw11[(_dafny.ONE).toNumber()] = (new BigNumber((_70_v42).length)).minus(_51_v25);
          _nw11[(new BigNumber(2)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(3)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(4)).toNumber()] = new BigNumber((_16_v9).length);
          _nw11[(new BigNumber(5)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(6)).toNumber()] = new BigNumber((_71_v43).length);
          _nw11[(new BigNumber(7)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(8)).toNumber()] = (_64_v36).dtor_cf13;
          _nw11[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_16_v9).length));
          _nw11[(new BigNumber(10)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(11)).toNumber()] = (_64_v36).dtor_cf13;
          _nw11[(new BigNumber(12)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(13)).toNumber()] = _51_v25;
          _nw11[(new BigNumber(14)).toNumber()] = _51_v25;
          _72_v44 = _nw11;
          let _index2 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_72_v44).length));
          (_72_v44)[_index2] = _51_v25;
          let _73_v45;
          _73_v45 = _dafny.Seq.of(_52_v26, p0);
          let _74_v46;
          _74_v46 = _module.D4.create_DC14(_73_v45, new _dafny.CodePoint('l'.codePointAt(0)), _51_v25, _51_v25);
          let _75_v47;
          _75_v47 = _dafny.Set.fromElements((_74_v46).dtor_cf22, _66_v38, _66_v38);
          let _76_v48;
          _76_v48 = _module.D3.create_DC11(_75_v47);
          let _pat_let_tv4 = _75_v47;
          let _pat_let_tv5 = _75_v47;
          let _index3 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_72_v44).length));
          let _rhs3 = !(p0) || (true);
          let _rhs4 = new BigNumber(((function (_pat_let0_0) {
            return function (_77_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_79_dt__update_hcf19_h0) {
                  return _module.D3.create_DC11(_79_dt__update_hcf19_h0);
                }(_pat_let1_0);
              }(function () {
                let _coll4 = new _dafny.Set();
                for (const _compr_4 of (_pat_let_tv4).Elements) {
                  let _78_v49 = _compr_4;
                  if ((_pat_let_tv5).contains(_78_v49)) {
                    _coll4.add(_78_v49);
                  }
                }
                return _coll4;
              }());
            }(_pat_let0_0);
          }(_76_v48)).dtor_cf19).length);
          let _lhs0 = _72_v44;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_72_v44).length));
          _52_v26 = _rhs3;
          _lhs0[_lhs1] = _rhs4;
          _73_v45 = _73_v45;
        } else {
          _51_v25 = _51_v25;
          let _80_v50;
          let _nw12 = new _module.C0();
          _nw12.__ctor(_18_v11);
          _80_v50 = _nw12;
          let _81_v51;
          _81_v51 = _module.D2.create_DC8(p0, _80_v50);
          _52_v26 = (((true) ? (_81_v51) : (_81_v51))).dtor_cf15;
          let _82_v52;
          _82_v52 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), _66_v38);
          let _83_v53;
          _83_v53 = _dafny.Map.Empty.slice().updateUnsafe((_67_v39).dtor_cf2,_16_v9);
          let _84_v54;
          _84_v54 = _dafny.Seq.of((_82_v52).IsSubsetOf(_dafny.MultiSet.fromElements(_66_v38, _66_v38)), !(false), (_51_v25).isLessThanOrEqualTo(new BigNumber(((((_83_v53).contains(_51_v25)) ? ((_83_v53).get(_51_v25)) : (_16_v9))).length)));
          _84_v54 = _dafny.Seq.Concat(_84_v54, _dafny.Seq.update(_84_v54, _module.__default.safeIndex(new BigNumber((_16_v9).length), new BigNumber((_84_v54).length)), p1));
          _51_v25 = (_51_v25).plus(_51_v25);
          _51_v25 = _51_v25;
        }
      }
      let _85_v55;
      _85_v55 = new BigNumber(115);
      let _86_v56;
      _86_v56 = _dafny.Set.fromElements(_85_v55, _85_v55);
      let _87_v57;
      _87_v57 = _dafny.Map.Empty.slice().updateUnsafe(_86_v56,_16_v9);
      _87_v57 = (_87_v57).Merge(_87_v57);
      return;
    }
    static Main(__noArgsParameter) {
      let _88_v0;
      _88_v0 = new BigNumber(159);
      let _89_v1;
      _89_v1 = _dafny.Seq.UnicodeFromString("vqomatru");
      let _90_v2;
      _90_v2 = _dafny.Seq.of(_88_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_88_v0)), _88_v0, new BigNumber((_89_v1).length));
      let _91_globalState;
      let _nw13 = new _module.GlobalState();
      _nw13.__ctor(new BigNumber(862), new BigNumber(771), new BigNumber(785), function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_90_v2).Elements) {
          let _92_v3 = _compr_5;
          if (_dafny.Seq.contains(_90_v2, _92_v3)) {
            _coll5.add((_92_v3).plus(new BigNumber(477)));
          }
        }
        return _coll5;
      }(), new BigNumber(-185), new _dafny.CodePoint('u'.codePointAt(0)), false, new BigNumber(385), true);
      _91_globalState = _nw13;
      let _93_v4;
      _93_v4 = _module.D0.create_DC0(new BigNumber((_dafny.Seq.Concat(_module.__default.fm0(_91_globalState), _89_v1)).length));
      let _source1 = _93_v4;
      if (_source1.is_DC0) {
        let _94___mcc_h0 = (_source1).cf0;
        let _95_cf0 = _94___mcc_h0;
        _88_v0 = _95_cf0;
        let _96_v5;
        _96_v5 = false;
        _96_v5 = _96_v5;
        let _97_v6;
        _97_v6 = _dafny.Set.fromElements(_89_v1, _dafny.Seq.UnicodeFromString("erdcbq"), _dafny.Seq.Concat(_89_v1, _89_v1));
        let _98_v7;
        let _nw14 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _98_v7 = _nw14;
        let _index4 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_98_v7).length));
        (_98_v7)[_index4] = _88_v0;
        let _index5 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_98_v7).length));
        let _rhs5 = (_97_v6).Difference(_dafny.Set.fromElements(_89_v1));
        let _rhs6 = _88_v0;
        let _lhs2 = _98_v7;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_98_v7).length));
        _97_v6 = _rhs5;
        _lhs2[_lhs3] = _rhs6;
        let _99_v8;
        _99_v8 = _dafny.Seq.of(_96_v5, (_96_v5) && (_96_v5));
        _96_v5 = (_99_v8)[_module.__default.safeIndex(_95_cf0, new BigNumber((_99_v8).length))];
      } else if (_source1.is_DC1) {
        let _100___mcc_h1 = (_source1).cf1;
        let _101___mcc_h2 = (_source1).cf2;
        let _102___mcc_h3 = (_source1).cf3;
        let _103___mcc_h4 = (_source1).cf4;
        let _104___mcc_h5 = (_source1).cf5;
        let _105_cf5 = _104___mcc_h5;
        let _106_cf4 = _103___mcc_h4;
        let _107_cf3 = _102___mcc_h3;
        let _108_cf2 = _101___mcc_h2;
        let _109_cf1 = _100___mcc_h1;
        let _110_v9;
        _110_v9 = _dafny.Set.fromElements(false);
        let _111_v10;
        _111_v10 = _dafny.Set.fromElements(_110_v9, _110_v9);
        _106_cf4 = !((_111_v10).IsSubsetOf((_111_v10).Union(_111_v10)));
        _module.__default.m0(_109_cf1, _109_cf1, _91_globalState);
        let _112_v11;
        let _nw15 = Array((new BigNumber(23)).toNumber()).fill(false);
        _112_v11 = _nw15;
        let _index6 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_112_v11).length));
        (_112_v11)[_index6] = _109_cf1;
        let _index7 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_112_v11).length));
        (_112_v11)[_index7] = ((_109_cf1) ? ((_108_cf2).isLessThan(_88_v0)) : (!(_106_cf4) || (false)));
        let _113_v12;
        let _nw16 = Array((_dafny.ONE).toNumber());
        _nw16[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _113_v12 = _nw16;
        let _114_v13;
        _114_v13 = _dafny.Seq.of(_113_v12);
        let _115_v14;
        let _nw17 = new _module.C0();
        _nw17.__ctor(_dafny.Seq.update(_114_v13, _module.__default.safeIndex(_105_cf5, new BigNumber((_114_v13).length)), _113_v12));
        _115_v14 = _nw17;
      } else {
        let _116___mcc_h6 = (_source1).cf6;
        let _117_cf6 = _116___mcc_h6;
        let _118_v15;
        _118_v15 = true;
        let _119_v16;
        _119_v16 = _dafny.Map.Empty.slice().updateUnsafe((_118_v15) === (_118_v15),_88_v0);
        _119_v16 = (_119_v16).update(_118_v15, _88_v0);
        let _120_v17;
        _120_v17 = new _dafny.CodePoint('d'.codePointAt(0));
        _118_v15 = _module.__default.fm2(!(_118_v15), _120_v17, _88_v0, ((!(_118_v15)) ? (_118_v15) : (_118_v15)), _91_globalState);
        let _121_v18;
        _121_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_88_v0),(_118_v15) || (_118_v15));
        _121_v18 = (_121_v18).update(new BigNumber(-558), (_118_v15) || (false));
        let _122_v19;
        let _nw18 = Array((_dafny.ONE).toNumber()).fill([]);
        _122_v19 = _nw18;
        let _123_v20;
        let _nw19 = Array((new BigNumber(9)).toNumber()).fill(_module.D0.Default());
        _123_v20 = _nw19;
        let _index8 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_122_v19).length));
        (_122_v19)[_index8] = _123_v20;
        let _index9 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_122_v19).length));
        (_122_v19)[_index9] = _123_v20;
      }
      let _124_v21;
      let _nw20 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
      _124_v21 = _nw20;
      let _125_v22;
      _125_v22 = new _dafny.CodePoint('h'.codePointAt(0));
      let _126_v23;
      let _nw21 = Array((new BigNumber(12)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = _125_v22;
      _nw21[(_dafny.ONE).toNumber()] = _125_v22;
      _nw21[(new BigNumber(2)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(3)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(4)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(5)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(6)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
      _nw21[(new BigNumber(8)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(9)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(10)).toNumber()] = _125_v22;
      _nw21[(new BigNumber(11)).toNumber()] = _125_v22;
      _126_v23 = _nw21;
      let _127_v24;
      _127_v24 = _dafny.Seq.of(_126_v23);
      let _128_v25;
      let _nw22 = new _module.C0();
      _nw22.__ctor(_127_v24);
      _128_v25 = _nw22;
      let _129_v26;
      _129_v26 = _dafny.Set.fromElements(_128_v25, _128_v25);
      let _index10 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_124_v21).length));
      (_124_v21)[_index10] = _129_v26;
      let _index11 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_124_v21).length));
      (_124_v21)[_index11] = _129_v26;
      let _source2 = _93_v4;
      if (_source2.is_DC0) {
        let _130___mcc_h7 = (_source2).cf0;
        let _131_cf0 = _130___mcc_h7;
        _89_v1 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rllcb"), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("t"), _module.__default.safeIndex(_88_v0, new BigNumber((_dafny.Seq.UnicodeFromString("t")).length)), _125_v22));
        let _132_v27;
        let _nw23 = new _module.C0();
        _nw23.__ctor(_dafny.Seq.of(_126_v23, _126_v23));
        _132_v27 = _nw23;
        _132_v27 = _132_v27;
        _88_v0 = _88_v0;
        let _133_v28;
        let _init2 = function (_134_i0) {
          return (false) === (true);
        };
        let _nw24 = Array((new BigNumber(9)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw24.length); _i0_2++) {
          _nw24[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _133_v28 = _nw24;
        let _135_v29;
        _135_v29 = true;
        let _index12 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_133_v28).length));
        (_133_v28)[_index12] = _135_v29;
        let _136_v30;
        _136_v30 = _dafny.Set.fromElements(_135_v29);
        let _index13 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_133_v28).length));
        (_133_v28)[_index13] = ((_135_v29) ? (!(_135_v29) || ((_module.D0.create_DC1(_135_v29, new BigNumber((_129_v26).length), (_dafny.ZERO).minus(_88_v0), _135_v29, new BigNumber(556))).dtor_cf4)) : ((new BigNumber((_136_v30).length)).isLessThanOrEqualTo(_88_v0)));
      } else if (_source2.is_DC1) {
        let _137___mcc_h8 = (_source2).cf1;
        let _138___mcc_h9 = (_source2).cf2;
        let _139___mcc_h10 = (_source2).cf3;
        let _140___mcc_h11 = (_source2).cf4;
        let _141___mcc_h12 = (_source2).cf5;
        let _142_cf5 = _141___mcc_h12;
        let _143_cf4 = _140___mcc_h11;
        let _144_cf3 = _139___mcc_h10;
        let _145_cf2 = _138___mcc_h9;
        let _146_cf1 = _137___mcc_h8;
        let _147_v31;
        let _nw25 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _147_v31 = _nw25;
        let _148_v32;
        let _nw26 = Array((new BigNumber(14)).toNumber()).fill(false);
        _148_v32 = _nw26;
        let _149_v33;
        _149_v33 = _dafny.Map.Empty.slice().updateUnsafe(_147_v31,_148_v32);
        _149_v33 = (_149_v33).update(_147_v31, _148_v32);
        _128_v25 = _128_v25;
        _89_v1 = _dafny.Seq.Concat(_89_v1, _89_v1);
        if (_143_cf4) {
          let _150_v34;
          _150_v34 = _module.D0.create_DC1(false, _142_cf5, _88_v0, _146_cf1, _145_cf2);
          _144_cf3 = (_150_v34).dtor_cf2;
          let _151_v35;
          _151_v35 = _dafny.Seq.of(_148_v32, _148_v32);
          let _152_v36;
          _152_v36 = _dafny.Seq.of((_151_v35)[_module.__default.safeIndex(_144_cf3, new BigNumber((_151_v35).length))], _148_v32, _148_v32);
          let _153_v37;
          _153_v37 = _dafny.Map.Empty.slice().updateUnsafe(_146_cf1,new BigNumber((_dafny.Seq.Concat(_151_v35, _151_v35)).length));
          let _154_v38;
          _154_v38 = _dafny.Seq.of(_146_cf1);
          let _155_v39;
          _155_v39 = _dafny.Map.Empty.slice().updateUnsafe(_89_v1,_154_v38);
          _153_v37 = (_153_v37).update((_dafny.Map.Empty.slice().updateUnsafe(_89_v1,_154_v38)).equals(_155_v39), _88_v0);
          let _index14 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_147_v31).length));
          (_147_v31)[_index14] = ((_143_cf4) ? (_145_cf2) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(_88_v0))));
          let _156_v40;
          _156_v40 = _dafny.Set.fromElements((_128_v25).fm1(_module.__default.fm2(_146_cf1, new _dafny.CodePoint('w'.codePointAt(0)), new BigNumber((_90_v2).length), !(_143_cf4), _91_globalState), _146_cf1, _146_cf1, _146_cf1, _91_globalState), _144_cf3, _145_cf2);
          let _index15 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_147_v31).length));
          (_147_v31)[_index15] = (_88_v0).plus(new BigNumber((_156_v40).length));
          _146_cf1 = (_150_v34).dtor_cf1;
          _146_cf1 = !(!(_module.__default.fm2(_143_cf4, ((_146_cf1) ? (_125_v22) : (_module.__default.fm3((_150_v34).dtor_cf3, _91_globalState))), (_144_cf3).minus((_dafny.ZERO).minus(_88_v0)), _146_cf1, _91_globalState)));
        } else {
          _module.__default.m0((_143_cf4) === (false), _146_cf1, _91_globalState);
          _88_v0 = ((_90_v2)[_module.__default.safeIndex(_142_cf5, new BigNumber((_90_v2).length))]).plus(new BigNumber(-475));
          let _index16 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_148_v32).length));
          (_148_v32)[_index16] = _146_cf1;
          let _index17 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_148_v32).length));
          (_148_v32)[_index17] = (_129_v26).contains(_128_v25);
          let _157_v41;
          let _nw27 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _157_v41 = _nw27;
          _157_v41 = _157_v41;
          let _158_v42;
          let _nw28 = Array((new BigNumber(15)).toNumber()).fill([]);
          _158_v42 = _nw28;
          _158_v42 = _158_v42;
        }
      } else {
        let _159___mcc_h13 = (_source2).cf6;
        let _160_cf6 = _159___mcc_h13;
        let _161_v43;
        _161_v43 = false;
        if (!(_161_v43)) {
          _88_v0 = _88_v0;
          let _162_v44;
          _162_v44 = _dafny.Seq.of(_161_v43, _161_v43, _161_v43, _161_v43, _161_v43);
          let _163_v45;
          _163_v45 = _dafny.MultiSet.fromElements(_module.__default.fm2(_161_v43, _125_v22, _88_v0, !(_161_v43), _91_globalState));
          let _164_v46;
          _164_v46 = _dafny.Map.Empty.slice().updateUnsafe(_88_v0,_88_v0);
          let _165_v47;
          _165_v47 = _module.D0.create_DC1(_161_v43, new BigNumber((_163_v45).cardinality()), new BigNumber((_164_v46).length), _161_v43, _88_v0);
          let _166_v48;
          let _nw29 = new _module.C0();
          _nw29.__ctor(_127_v24);
          _166_v48 = _nw29;
          let _167_v49;
          _167_v49 = _dafny.Map.Empty.slice().updateUnsafe(_166_v48,_161_v43);
          let _168_v50;
          _168_v50 = _dafny.MultiSet.fromElements(_88_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(529)), ((_169_v22) => function (_170_i1) {
            return _169_v22;
          })(_125_v22))).length));
          let _171_v51;
          _171_v51 = _module.D0.create_DC1((_165_v47).dtor_cf4, _88_v0, new BigNumber(((_167_v49).update(_166_v48, _161_v43)).length), _161_v43, (((_168_v50).contains(_88_v0)) ? ((_168_v50).get(_88_v0)) : ((_90_v2)[_module.__default.safeIndex(_88_v0, new BigNumber((_90_v2).length))])));
          let _172_v52;
          _172_v52 = _dafny.MultiSet.fromElements(_128_v25);
          let _173_v53;
          let _nw30 = Array((new BigNumber(12)).toNumber());
          _nw30[(_dafny.ZERO).toNumber()] = _88_v0;
          _nw30[(_dafny.ONE).toNumber()] = new BigNumber((_162_v44).length);
          _nw30[(new BigNumber(2)).toNumber()] = _88_v0;
          _nw30[(new BigNumber(3)).toNumber()] = new BigNumber(916);
          _nw30[(new BigNumber(4)).toNumber()] = (_165_v47).dtor_cf5;
          _nw30[(new BigNumber(5)).toNumber()] = _88_v0;
          _nw30[(new BigNumber(6)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_172_v52).cardinality()), (_dafny.ZERO).minus((_166_v48).fm1(_161_v43, _161_v43, _161_v43, false, _91_globalState)));
          _nw30[(new BigNumber(7)).toNumber()] = new BigNumber((_89_v1).length);
          _nw30[(new BigNumber(8)).toNumber()] = _88_v0;
          _nw30[(new BigNumber(9)).toNumber()] = new BigNumber(810);
          _nw30[(new BigNumber(10)).toNumber()] = _88_v0;
          _nw30[(new BigNumber(11)).toNumber()] = _88_v0;
          _173_v53 = _nw30;
          let _index18 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_173_v53).length));
          (_173_v53)[_index18] = _88_v0;
          let _index19 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_173_v53).length));
          (_173_v53)[_index19] = (_166_v48).fm1((_88_v0).isLessThan(_88_v0), _161_v43, _161_v43, ((_module.__default.fm2(_module.__default.fm2(_161_v43, _125_v22, _88_v0, _161_v43, _91_globalState), _125_v22, _88_v0, _161_v43, _91_globalState)) ? (_module.__default.fm2(_161_v43, _125_v22, new BigNumber((_89_v1).length), _161_v43, _91_globalState)) : (false)), _91_globalState);
          let _index20 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_173_v53).length));
          (_173_v53)[_index20] = (_128_v25).fm1(_161_v43, _module.__default.fm2(_161_v43, _125_v22, new BigNumber(-962), _161_v43, _91_globalState), _161_v43, !((_168_v50).IsSubsetOf(_168_v50)), _91_globalState);
          _88_v0 = new BigNumber((_89_v1).length);
          _161_v43 = !(_161_v43);
        } else {
          let _174_v54;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_127_v24);
          _174_v54 = _nw31;
          let _175_v55;
          _175_v55 = _dafny.Seq.of(_174_v54, _174_v54, _174_v54, _174_v54);
          let _176_v56;
          _176_v56 = _module.D1.create_DC3((_175_v55)[_module.__default.safeIndex(_88_v0, new BigNumber((_175_v55).length))]);
          let _177_v57;
          let _nw32 = Array((new BigNumber(21)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = _174_v54;
          _nw32[(_dafny.ONE).toNumber()] = (_module.D1.create_DC3(_174_v54)).dtor_cf7;
          _nw32[(new BigNumber(2)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(3)).toNumber()] = (_176_v56).dtor_cf7;
          _nw32[(new BigNumber(4)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(5)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(6)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(7)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(8)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(9)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(10)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(11)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(12)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(13)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(14)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(15)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(16)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(17)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(18)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(19)).toNumber()] = _174_v54;
          _nw32[(new BigNumber(20)).toNumber()] = _174_v54;
          _177_v57 = _nw32;
          let _index21 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_177_v57).length));
          (_177_v57)[_index21] = _174_v54;
          let _index22 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_177_v57).length));
          (_177_v57)[_index22] = ((_161_v43) ? (_174_v54) : (((_161_v43) ? (_174_v54) : (_174_v54))));
          let _178_v58;
          _178_v58 = _dafny.Map.Empty.slice().updateUnsafe(_161_v43,false);
          _178_v58 = (_178_v58).update(_161_v43, _161_v43);
          let _179_v59;
          let _nw33 = Array((new BigNumber(14)).toNumber()).fill(false);
          _179_v59 = _nw33;
          _179_v59 = _179_v59;
          let _180_v60;
          let _nw34 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _180_v60 = _nw34;
          let _nw35 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
          _180_v60 = _nw35;
          let _181_v61;
          _181_v61 = _dafny.Seq.of(_161_v43, _161_v43, _161_v43);
          _181_v61 = _181_v61;
        }
        let _182_v62;
        let _nw36 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _182_v62 = _nw36;
        let _index23 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_182_v62).length));
        (_182_v62)[_index23] = _88_v0;
        let _index24 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_182_v62).length));
        (_182_v62)[_index24] = _88_v0;
        _161_v43 = _161_v43;
        _88_v0 = _88_v0;
      }
      let _183_v63;
      _183_v63 = true;
      if (_183_v63) {
        _183_v63 = _183_v63;
        let _184_v64;
        _184_v64 = _dafny.Map.Empty.slice().updateUnsafe(_88_v0,false);
        _88_v0 = ((new BigNumber((_184_v64).length)).multipliedBy(_88_v0)).plus(_88_v0);
        _89_v1 = _module.__default.fm0(_91_globalState);
        if (!(_183_v63)) {
          _183_v63 = _183_v63;
          let _185_v65;
          let _nw37 = Array((new BigNumber(20)).toNumber()).fill(false);
          _185_v65 = _nw37;
          let _186_v66;
          _186_v66 = _dafny.Seq.of(_183_v63);
          let _187_v67;
          _187_v67 = _dafny.Set.fromElements(new BigNumber((_186_v66).length), _88_v0);
          let _188_v68;
          _188_v68 = _module.D0.create_DC1(_183_v63, _88_v0, _88_v0, _183_v63, new BigNumber((_187_v67).length));
          let _index25 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_185_v65).length));
          (_185_v65)[_index25] = (_188_v68).dtor_cf1;
          let _index26 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_185_v65).length));
          (_185_v65)[_index26] = true;
          let _index27 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_185_v65).length));
          (_185_v65)[_index27] = _183_v63;
          let _189_v69;
          _189_v69 = _dafny.Map.Empty.slice().updateUnsafe((_185_v65)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_185_v65).length))],_186_v66);
          _189_v69 = _189_v69;
          let _190_v70;
          let _nw38 = new _module.C0();
          _nw38.__ctor(_127_v24);
          _190_v70 = _nw38;
          let _191_v71;
          _191_v71 = _dafny.Map.Empty.slice().updateUnsafe(_190_v70,_88_v0);
          _191_v71 = (_191_v71).update(_190_v70, _88_v0);
        } else {
          let _192_v72;
          _192_v72 = _dafny.Map.Empty.slice().updateUnsafe(_183_v63,new _dafny.CodePoint('m'.codePointAt(0)));
          _192_v72 = _192_v72;
          let _193_v73;
          _193_v73 = _dafny.MultiSet.fromElements(_183_v63);
          let _194_v74;
          _194_v74 = _dafny.Seq.of(_183_v63);
          let _195_v75;
          let _nw39 = Array((new BigNumber(26)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = _193_v73;
          _nw39[(_dafny.ONE).toNumber()] = _module.__default.fm4(false, _183_v63, _91_globalState);
          _nw39[(new BigNumber(2)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(false, _183_v63, _183_v63, _183_v63, _183_v63);
          _nw39[(new BigNumber(4)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(5)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(6)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(_183_v63);
          _nw39[(new BigNumber(8)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(9)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(10)).toNumber()] = _module.__default.fm4(_183_v63, false, _91_globalState);
          _nw39[(new BigNumber(11)).toNumber()] = (_193_v73).update(_183_v63, _module.__default.abs(_88_v0));
          _nw39[(new BigNumber(12)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(13)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(14)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(15)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(16)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(17)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(_183_v63);
          _nw39[(new BigNumber(19)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(20)).toNumber()] = _dafny.MultiSet.fromElements(_183_v63);
          _nw39[(new BigNumber(21)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(22)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(23)).toNumber()] = _module.__default.fm4(_module.__default.fm2(_183_v63, (_89_v1)[_module.__default.safeIndex(new BigNumber((_184_v64).length), new BigNumber((_89_v1).length))], _88_v0, _183_v63, _91_globalState), (_194_v74)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_194_v74).length))], _91_globalState);
          _nw39[(new BigNumber(24)).toNumber()] = _193_v73;
          _nw39[(new BigNumber(25)).toNumber()] = _193_v73;
          _195_v75 = _nw39;
          let _196_v76;
          _196_v76 = _dafny.Map.Empty.slice().updateUnsafe(_195_v75,!(_183_v63) || (_183_v63));
          _183_v63 = (((_196_v76).contains(_195_v75)) ? ((_196_v76).get(_195_v75)) : (_183_v63));
          _88_v0 = _88_v0;
          let _index28 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_124_v21).length));
          (_124_v21)[_index28] = ((_124_v21)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_124_v21).length))]).Difference(_129_v26);
          let _197_v77;
          let _init3 = ((_198_v1) => function (_199_i2) {
            return _198_v1;
          })(_89_v1);
          let _nw40 = Array((new BigNumber(12)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw40.length); _i0_3++) {
            _nw40[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _197_v77 = _nw40;
          let _index29 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_197_v77).length));
          (_197_v77)[_index29] = _89_v1;
          let _200_v78;
          let _nw41 = Array((new BigNumber(23)).toNumber()).fill(false);
          _200_v78 = _nw41;
          let _index30 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_200_v78).length));
          (_200_v78)[_index30] = !(_183_v63);
          let _index31 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_197_v77).length));
          let _index32 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_200_v78).length));
          let _rhs7 = _89_v1;
          let _rhs8 = (!(_183_v63)) && (_183_v63);
          let _lhs4 = _197_v77;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_197_v77).length));
          let _lhs6 = _200_v78;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_200_v78).length));
          _lhs4[_lhs5] = _rhs7;
          _lhs6[_lhs7] = _rhs8;
        }
        let _201_v79;
        _201_v79 = _dafny.Seq.of(_183_v63);
        _183_v63 = !(_module.__default.fm2(_183_v63, _125_v22, new BigNumber((_90_v2).length), _module.__default.fm2(_183_v63, _125_v22, new BigNumber((_201_v79).length), _183_v63, _91_globalState), _91_globalState)) || (!(!(_183_v63)));
      } else {
        let _index33 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_126_v23).length));
        (_126_v23)[_index33] = (_89_v1)[_module.__default.safeIndex(_88_v0, new BigNumber((_89_v1).length))];
        let _index34 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_126_v23).length));
        let _rhs9 = false;
        let _rhs10 = _88_v0;
        let _rhs11 = _125_v22;
        let _rhs12 = ((_183_v63) ? (_88_v0) : (_88_v0));
        let _rhs13 = true;
        let _lhs8 = _126_v23;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_126_v23).length));
        _183_v63 = _rhs9;
        _88_v0 = _rhs10;
        _lhs8[_lhs9] = _rhs11;
        _88_v0 = _rhs12;
        _183_v63 = _rhs13;
        let _202_v80;
        let _init4 = ((_203_v2, _204_v0) => function (_205_i3) {
          return _dafny.Seq.Concat(_203_v2, _dafny.Seq.of(_204_v0, _204_v0, _204_v0));
        })(_90_v2, _88_v0);
        let _nw42 = Array((new BigNumber(28)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw42.length); _i0_4++) {
          _nw42[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _202_v80 = _nw42;
        _202_v80 = _202_v80;
        let _206_v81;
        _206_v81 = _dafny.Map.Empty.slice().updateUnsafe(_90_v2,_89_v1);
        _206_v81 = _206_v81;
        _183_v63 = _183_v63;
        let _index35 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_126_v23).length));
        (_126_v23)[_index35] = (_126_v23)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_126_v23).length))];
      }
      _module.__default.m0(_183_v63, (_183_v63) === (_183_v63), _91_globalState);
      _183_v63 = _183_v63;
      _183_v63 = _183_v63;
      let _207_v82;
      let _init5 = ((_208_v2) => function (_209_i4) {
        return _208_v2;
      })(_90_v2);
      let _nw43 = Array((_dafny.ONE).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
        _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _207_v82 = _nw43;
      let _210_v83;
      _210_v83 = _dafny.MultiSet.fromElements(_183_v63, _183_v63);
      let _211_v84;
      _211_v84 = _dafny.Map.Empty.slice().updateUnsafe(_207_v82,_210_v83);
      let _212_v85;
      _212_v85 = _dafny.Seq.of((((_211_v84).contains(_207_v82)) ? ((_211_v84).get(_207_v82)) : (_210_v83)), _210_v83, _210_v83);
      _212_v85 = _212_v85;
      let _213_v86;
      let _nw44 = new _module.C0();
      _nw44.__ctor(_127_v24);
      _213_v86 = _nw44;
      _module.__default.m0(_183_v63, _183_v63, _91_globalState);
      let _214_v87;
      let _init6 = ((_215_v2, _216_v0) => function (_217_i5) {
        return _module.__default.safeModuloInt(_217_i5, (_215_v2)[_module.__default.safeIndex(_216_v0, new BigNumber((_215_v2).length))]);
      })(_90_v2, _88_v0);
      let _nw45 = Array((new BigNumber(16)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
        _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _214_v87 = _nw45;
      let _index36 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
      (_214_v87)[_index36] = _88_v0;
      let _218_v88;
      _218_v88 = _dafny.Map.Empty.slice().updateUnsafe(_183_v63,new BigNumber(211));
      let _219_v89;
      _219_v89 = _dafny.Set.fromElements(new BigNumber((_90_v2).length));
      let _index37 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
      (_214_v87)[_index37] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((((_218_v88).contains(_183_v63)) ? ((_218_v88).get(_183_v63)) : (_88_v0)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_219_v89).length),(_128_v25).fm1(false, _183_v63, !(_183_v63), !(_183_v63), _91_globalState))).length)));
      let _index38 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
      (_214_v87)[_index38] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_88_v0, (_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]));
      let _index39 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
      (_214_v87)[_index39] = (_88_v0).minus((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]);
      let _220_v90;
      let _nw46 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _220_v90 = _nw46;
      let _index40 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length));
      (_220_v90)[_index40] = _89_v1;
      let _index41 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length));
      (_220_v90)[_index41] = _89_v1;
      let _221_v91;
      _221_v91 = _dafny.Map.Empty.slice().updateUnsafe((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))],(_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]);
      _221_v91 = ((_183_v63) ? (function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_90_v2).Elements) {
          let _222_v92 = _compr_6;
          if (_dafny.Seq.contains(_90_v2, _222_v92)) {
            _coll6.push([(_222_v92).multipliedBy(_88_v0),(_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]]);
          }
        }
        return _coll6;
      }()) : (_221_v91));
      let _223_v93;
      _223_v93 = _dafny.Set.fromElements(_210_v83, (_210_v83).update(_183_v63, _module.__default.abs((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))])), _dafny.MultiSet.fromElements(_183_v63), _210_v83);
      if ((_dafny.Set.fromElements(_210_v83, _210_v83)).IsDisjointFrom(_223_v93)) {
        let _224_v94;
        _224_v94 = _dafny.Seq.of(_89_v1, _dafny.Seq.UnicodeFromString("rqfejscgi"), _89_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-358)), ((_225_v22) => function (_226_i6) {
          return _225_v22;
        })(_125_v22)));
        let _index42 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length));
        let _index43 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
        let _rhs14 = _dafny.Seq.Concat((_224_v94)[_module.__default.safeIndex((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))], new BigNumber((_224_v94).length))], _dafny.Seq.Concat(_89_v1, (_220_v90)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length))]));
        let _rhs15 = ((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]).minus((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]);
        let _lhs10 = _220_v90;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length));
        let _lhs12 = _214_v87;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
        _lhs10[_lhs11] = _rhs14;
        _lhs12[_lhs13] = _rhs15;
        let _227_v95;
        _227_v95 = _dafny.MultiSet.fromElements(_88_v0);
        let _228_v97;
        _228_v97 = _dafny.Map.Empty.slice().updateUnsafe(_227_v95,_219_v89);
        _module.__default.m0((_dafny.MultiSet.FromArray(_dafny.Seq.of((_213_v86).fm1(_183_v63, _183_v63, _183_v63, _183_v63, _91_globalState), _88_v0))).IsSubsetOf(_227_v95), (function () {
          let _coll7 = new _dafny.Set();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-62), new BigNumber(728))) {
            let _229_v96 = _compr_7;
            if (((new BigNumber(-62)).isLessThanOrEqualTo(_229_v96)) && ((_229_v96).isLessThan(new BigNumber(728)))) {
              _coll7.add((_229_v96).minus(_88_v0));
            }
          }
          return _coll7;
        }()).IsDisjointFrom((((_228_v97).contains(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wbacwiua")).length)))) ? ((_228_v97).get(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("wbacwiua")).length)))) : (_219_v89))), _91_globalState);
        let _index44 = _module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length));
        (_214_v87)[_index44] = (_dafny.ZERO).minus(_88_v0);
        let _230_v98;
        _230_v98 = _module.D2.create_DC7((_220_v90)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length))], new BigNumber((_89_v1).length), (_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))], _183_v63);
        let _231_v99;
        _231_v99 = _dafny.MultiSet.fromElements((_230_v98).dtor_cf11, _89_v1, (_220_v90)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length))]);
        let _232_v100;
        _232_v100 = _dafny.Map.Empty.slice().updateUnsafe(_231_v99,_183_v63);
        _232_v100 = (_232_v100).update((_dafny.MultiSet.fromElements(_89_v1)).Difference(_231_v99), false);
        let _233_v101;
        _233_v101 = _dafny.Map.Empty.slice().updateUnsafe((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))],_183_v63);
        _233_v101 = _dafny.Map.Empty.slice().updateUnsafe((_213_v86).fm1(_183_v63, _183_v63, _183_v63, _183_v63, _91_globalState),_183_v63);
      } else {
        _183_v63 = _183_v63;
        let _234_v102;
        let _nw47 = new _module.C0();
        _nw47.__ctor(_dafny.Seq.Concat(_127_v24, _127_v24));
        _234_v102 = _nw47;
        let _235_v103;
        _235_v103 = _dafny.Seq.of(_183_v63, _183_v63, false, _183_v63);
        let _236_v104;
        _236_v104 = _module.D2.create_DC7(_dafny.Seq.update((_220_v90)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length))], _module.__default.safeIndex(_88_v0, new BigNumber(((_220_v90)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_220_v90).length))]).length)), _125_v22), _88_v0, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_183_v63)).length)), _183_v63);
        let _237_v105;
        _237_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_89_v1).length),_183_v63);
        let _238_v106;
        _238_v106 = _module.D1.create_DC4(_214_v87);
        let _239_v107;
        _239_v107 = _dafny.Seq.of(_214_v87, (_238_v106).dtor_cf8);
        let _rhs16 = (_90_v2)[_module.__default.safeIndex((_234_v102).fm1(_183_v63, (_235_v103)[_module.__default.safeIndex((_236_v104).dtor_cf13, new BigNumber((_235_v103).length))], _183_v63, !(!(_183_v63)), _91_globalState), new BigNumber((_90_v2).length))];
        let _rhs17 = _234_v102;
        let _rhs18 = ((!(_module.__default.fm2(!(_183_v63), _module.__default.fm3(new BigNumber((_237_v105).length), _91_globalState), (_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))], _183_v63, _91_globalState))) ? (!((new BigNumber((_239_v107).length)).isLessThan((_214_v87)[_module.__default.safeIndex(new BigNumber(544), new BigNumber((_214_v87).length))]))) : (_183_v63));
        let _rhs19 = _234_v102;
        _88_v0 = _rhs16;
        _234_v102 = _rhs17;
        _183_v63 = _rhs18;
        _234_v102 = _rhs19;
        _125_v22 = (((_183_v63) === (_183_v63)) ? (_125_v22) : (_125_v22));
        _128_v25 = _128_v25;
        _221_v91 = (_221_v91).update(_88_v0, (_234_v102).fm1(_183_v63, _183_v63, true, _183_v63, _91_globalState));
      }
      process.stdout.write(_dafny.toString(_88_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_89_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_90_v2, _dafny.Seq.of(new BigNumber(159), new BigNumber(159), new BigNumber(159), new BigNumber(8)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f3).equals(_dafny.Set.fromElements(new BigNumber(636), new BigNumber(485)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_v4).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_124_v21)[new BigNumber(9)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_125_v22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_126_v23)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_127_v24).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_129_v26).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_v63));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_207_v82)[_dafny.ZERO], _dafny.Seq.of(new BigNumber(159), new BigNumber(159), new BigNumber(159), new BigNumber(8)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_210_v83).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_211_v84).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_212_v85, _dafny.Seq.of(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v87)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_218_v88).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(211)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v89).equals(_dafny.Set.fromElements(new BigNumber(4)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_220_v90)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_221_v91).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(205110),new BigNumber(2580)).updateUnsafe(new BigNumber(10320),new BigNumber(2580)).updateUnsafe(new BigNumber(159),new BigNumber(1105)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v93).equals(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true), _dafny.MultiSet.fromElements(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true), _dafny.MultiSet.fromElements(true)))));
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
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf6) {
      let $dt = new D0(2);
      $dt.cf6 = cf6;
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
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(_dafny.ZERO);
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
    static create_DC4(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4([]);
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
    static create_DC7(cf11, cf12, cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC8(cf15, cf16) {
      let $dt = new D2(1);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D2(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D2(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC10(cf18) {
      let $dt = new D2(4);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get is_DC10() { return this.$tag === 4; }
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
        return "D2.DC7" + "(" + this.cf11.toVerbatimString(true) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else if (this.$tag === 4) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf17 === other.cf17;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC12() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D3(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC12";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC12();
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
    static create_DC14(cf21, cf22, cf23, cf24) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D4(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(_dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = _dafny.ZERO;
      this._f3 = _dafny.Set.Empty;
      this._f4 = _dafny.ZERO;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f9 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    set f9(value) {
      let _this = this;
      _this._f9 = value;
    };
    __ctor(f9) {
      let _this = this;
      (_this)._f9 = f9;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), function (_240_i0) {
        return new BigNumber(141);
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-967)), function (_241_i1) {
        return new BigNumber(81);
      }))).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true, !(false), (_module.D0.create_DC1(true, new BigNumber(486), new BigNumber((_dafny.Seq.of(true, true)).length), true, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(308)))).cardinality()))).dtor_cf4), _dafny.Seq.of(true))).length)));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
