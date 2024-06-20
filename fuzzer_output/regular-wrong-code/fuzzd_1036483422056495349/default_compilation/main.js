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
    static fm0(p0, p1, p2, p3, globalState) {
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(997)), _dafny.Seq.of(new BigNumber(-776), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-746), new BigNumber(499))).cardinality()), new BigNumber(549)))).length), new BigNumber(38));
    };
    static fm1(p0, p1, p2, globalState) {
      return (_module.D1.create_DC4(!(false), false, true)).dtor_cf6;
    };
    static fm2(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(39), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(633), (_dafny.ZERO).minus(new BigNumber(-268)))).cardinality())), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-431)), function (_0_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)));
    };
    static fm3(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vjsfret"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bm"), _dafny.Seq.UnicodeFromString("pwkap")));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), function (_1_i0) {
        return new BigNumber(-185);
      })).length),new BigNumber(886))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tysocxty")).length),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())));
    };
    static fm7(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false, !(false), !(false), false)), _dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false, true, !(true), true))));
    };
    static fm8(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(542), new BigNumber(145)),false)).Merge((_module.D13.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(531), new BigNumber((_dafny.Seq.UnicodeFromString("pq")).length)),true))).dtor_cf45);
    };
    static fm9(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(413)).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(603))).cardinality())),_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(686)), _dafny.Set.fromElements(new BigNumber(815)))).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-851)), function (_2_i0) {
        return _module.D11.create_DC30(false);
      })).length)), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())));
    };
    static fm10(p0, globalState) {
      return _module.D2.create_DC5(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uvfc")).length),new BigNumber((_dafny.Seq.of(true)).length)));
    };
    static fm11(globalState) {
      return _module.D1.create_DC3((new BigNumber((_dafny.MultiSet.fromElements(false, !(false))).cardinality())).minus(new BigNumber((_dafny.Seq.UnicodeFromString("epbbayn")).length)));
    };
    static fm12(globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ap")).length));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements((new BigNumber(587)).isLessThan(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), function (_3_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).Elements) {
          let _4_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(139)), function (_3_i0) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          }), _4_v0)) {
            _coll0.push([_4_v0,true]);
          }
        }
        return _coll0;
      }()).length)), ((false) ? (true) : (true)), (false) || (false));
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('h'.codePointAt(0)));
    };
    static fm15(p0, p1, globalState) {
      if ((new BigNumber(231)).isLessThan(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(973)), function (_5_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length));
      })).length))).length))) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }
    };
    static fm16(p0, globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("pyah"));
    };
    static fm17(globalState) {
      if (false) {
        return _module.D6.create_DC17(new BigNumber(380), true);
      } else {
        return _module.D6.create_DC17(new BigNumber(-419), true);
      }
    };
    static fm18(p0, globalState) {
      return _module.D5.create_DC13((new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(854),new _dafny.CodePoint('q'.codePointAt(0)))).length)));
    };
    static fm19(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D1.create_DC3(new BigNumber((_dafny.Seq.of(new BigNumber(338))).length)), _module.D1.create_DC3(new BigNumber(553)), _module.D1.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)), _module.D1.create_DC3(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-257)))).cardinality())), _module.D1.create_DC3(new BigNumber((_dafny.Seq.of(new BigNumber(295), new BigNumber(780))).length)))).Intersect(_dafny.MultiSet.fromElements(_module.D1.create_DC3(new BigNumber(567))))).Intersect((_dafny.MultiSet.fromElements(_module.D1.create_DC3(new BigNumber(602)), _module.D1.create_DC3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D1.create_DC3(new BigNumber((_dafny.Seq.UnicodeFromString("htapck")).length)))).length)), _module.D1.create_DC3(new BigNumber(-138)))).Intersect(_dafny.MultiSet.fromElements(_module.D1.create_DC3(new BigNumber(967)))));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let _hi0 = p0;
      for (let _6_i0 = _module.__default.safeDivisionInt(p0, p0); _6_i0.isLessThan(_hi0); _6_i0 = _6_i0.plus(_dafny.ONE)) {
        let _7_v0;
        _7_v0 = false;
        let _8_v1;
        _8_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-239),_7_v0);
        let _9_v2;
        _9_v2 = _module.D0.create_DC0(_7_v0);
        let _10_v3;
        _10_v3 = _dafny.Set.fromElements(_6_i0);
        let _11_v4;
        let _nw0 = new _module.C1();
        _nw0.__ctor(_7_v0, (p1).minus(new BigNumber(((_8_v1).update(p0, (_9_v2).dtor_cf0)).length)), p0, (_10_v3).IsProperSubsetOf(_dafny.Set.fromElements(p0, p0)));
        _11_v4 = _nw0;
        r0 = _module.__default.safeModuloInt(p1, p0);
        let _12_v6;
        _12_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(947), new BigNumber(518))) {
            let _13_v5 = _compr_1;
            if (((new BigNumber(947)).isLessThanOrEqualTo(_13_v5)) && ((_13_v5).isLessThan(new BigNumber(518)))) {
              _coll1.add(_module.__default.safeDivisionInt(_13_v5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_14_i1) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              })).length)));
            }
          }
          return _coll1;
        }());
        let _15_v7;
        _15_v7 = _dafny.Map.Empty.slice().updateUnsafe(_11_v4.f21,new BigNumber((_12_v6).length));
        let _16_v8;
        _16_v8 = _dafny.Seq.of(new BigNumber(-998), p1);
        let _17_v9;
        _17_v9 = _module.D5.create_DC13((_16_v8)[_module.__default.safeIndex(_6_i0, new BigNumber((_16_v8).length))]);
        let _18_v10;
        _18_v10 = _dafny.Set.fromElements(_17_v9);
        let _19_v11;
        _19_v11 = _dafny.Set.fromElements(_18_v10);
        let _20_v12;
        _20_v12 = _dafny.Map.Empty.slice().updateUnsafe(_15_v7,_19_v11);
        let _21_v13;
        _21_v13 = _dafny.MultiSet.fromElements(_11_v4.f21);
        _20_v12 = (_20_v12).update(_dafny.Map.Empty.slice().updateUnsafe(_7_v0,new BigNumber((_21_v13).cardinality())), _19_v11);
        let _22_v14;
        let _nw1 = new _module.C1();
        _nw1.__ctor(_7_v0, new BigNumber(350), p1, _7_v0);
        _22_v14 = _nw1;
      }
      let _23_v15;
      _23_v15 = false;
      let _24_v16;
      _24_v16 = _dafny.Seq.of(_23_v15, _23_v15);
      if (!((_24_v16)[_module.__default.safeIndex(p1, new BigNumber((_24_v16).length))])) {
        let _25_v17;
        _25_v17 = new _dafny.CodePoint('m'.codePointAt(0));
        let _26_v18;
        _26_v18 = _dafny.Map.Empty.slice().updateUnsafe(_23_v15,new BigNumber(837));
        let _27_v19;
        _27_v19 = _dafny.MultiSet.fromElements(p0, (((_26_v18).contains(_23_v15)) ? ((_26_v18).get(_23_v15)) : (p1)));
        let _28_v20;
        _28_v20 = _module.D2.create_DC6(_27_v19);
        let _29_v21;
        _29_v21 = _module.D2.create_DC7(_28_v20);
        _25_v17 = _module.__default.fm15(_29_v21, !(_23_v15), globalState);
        let _30_v22;
        let _nw2 = Array((new BigNumber(2)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _25_v17;
        _nw2[(_dafny.ONE).toNumber()] = _25_v17;
        _30_v22 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_30_v22).length));
        (_30_v22)[_index0] = new _dafny.CodePoint('a'.codePointAt(0));
        let _31_v23;
        _31_v23 = _dafny.MultiSet.fromElements(_23_v15, _23_v15);
        let _32_v24;
        _32_v24 = _dafny.Seq.of(_31_v23);
        let _33_v25;
        _33_v25 = _dafny.Seq.of(p0);
        let _index1 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_30_v22).length));
        let _rhs0 = !_dafny.areEqual(_dafny.Seq.of(new BigNumber((_32_v24).length)), _33_v25);
        let _rhs1 = !(!(_31_v23).contains(_23_v15));
        let _rhs2 = _25_v17;
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = _30_v22;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_30_v22).length));
        _lhs0.f13 = _rhs0;
        _lhs1.f13 = _rhs1;
        _lhs2[_lhs3] = _rhs2;
        let _34_v26;
        let _nw3 = new _module.C1();
        _nw3.__ctor(!(_23_v15), p0, p0, _23_v15);
        _34_v26 = _nw3;
        let _35_v27;
        _35_v27 = _dafny.Seq.UnicodeFromString("w");
        let _36_v28;
        _36_v28 = _dafny.MultiSet.fromElements(_35_v27);
        let _37_v29;
        _37_v29 = _module.D0.create_DC1((_34_v26).f19, new BigNumber((_24_v16).length), _23_v15);
        let _38_v30;
        _38_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_35_v27);
        let _39_v31;
        let _nw4 = Array((new BigNumber(21)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), function (_40_i2) {
          return _dafny.Seq.UnicodeFromString("nfnc");
        })).length);
        _nw4[(_dafny.ONE).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(2)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(3)).toNumber()] = (_34_v26).f19;
        _nw4[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_24_v16, _24_v16)).length);
        _nw4[(new BigNumber(5)).toNumber()] = (new BigNumber(-199)).minus(new BigNumber((_35_v27).length));
        _nw4[(new BigNumber(6)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(7)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(579), (_34_v26).f19), _33_v25)).length);
        _nw4[(new BigNumber(9)).toNumber()] = new BigNumber(((_36_v28).Intersect(_module.__default.fm16((_37_v29).dtor_cf3, globalState))).cardinality());
        _nw4[(new BigNumber(10)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(11)).toNumber()] = new BigNumber((_35_v27).length);
        _nw4[(new BigNumber(12)).toNumber()] = new BigNumber((_38_v30).length);
        _nw4[(new BigNumber(13)).toNumber()] = (_34_v26).f19;
        _nw4[(new BigNumber(14)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("b")).length);
        _nw4[(new BigNumber(16)).toNumber()] = p1;
        _nw4[(new BigNumber(17)).toNumber()] = _34_v26.f20;
        _nw4[(new BigNumber(18)).toNumber()] = new BigNumber(((_module.__default.fm12(globalState)).update(p1, _module.__default.abs(p1))).cardinality());
        _nw4[(new BigNumber(19)).toNumber()] = new BigNumber(-579);
        _nw4[(new BigNumber(20)).toNumber()] = _34_v26.f20;
        _39_v31 = _nw4;
        let _index2 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_39_v31).length));
        (_39_v31)[_index2] = new BigNumber(-751);
        let _41_v32;
        let _nw5 = new _module.C0();
        _nw5.__ctor(_module.D2.create_DC6(_27_v19), new BigNumber(-841), (_34_v26).f18);
        _41_v32 = _nw5;
        let _42_v33;
        _42_v33 = _dafny.Map.Empty.slice().updateUnsafe(_41_v32,p1);
        let _index3 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_39_v31).length));
        let _rhs3 = _23_v15;
        let _rhs4 = _dafny.Seq.of((((_42_v33).contains(_41_v32)) ? ((_42_v33).get(_41_v32)) : (p1)), (_34_v26).f19, ((_23_v15) ? (p1) : (_34_v26.f20)), (_34_v26).f19);
        let _rhs5 = _34_v26;
        let _rhs6 = (p0).plus(p0);
        let _lhs4 = globalState;
        let _lhs5 = _39_v31;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(980), new BigNumber((_39_v31).length));
        _lhs4.f13 = _rhs3;
        _33_v25 = _rhs4;
        _34_v26 = _rhs5;
        _lhs5[_lhs6] = _rhs6;
        let _43_v34;
        let _nw6 = Array((new BigNumber(4)).toNumber());
        _43_v34 = _nw6;
        let _44_v35;
        _44_v35 = _dafny.Map.Empty.slice().updateUnsafe(_43_v34,(_34_v26).f18);
        _44_v35 = (_44_v35).update(_43_v34, _23_v15);
        let _45_v36;
        _45_v36 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_33_v25, _module.__default.safeIndex((_34_v26).f19, new BigNumber((_33_v25).length)), p1)).length));
        (_34_v26).f20 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_33_v25).length), _module.__default.safeDivisionInt(p1, new BigNumber((_45_v36).length)))));
      } else {
        (globalState).f13 = _23_v15;
        let _46_v37;
        let _nw7 = Array((new BigNumber(23)).toNumber());
        _46_v37 = _nw7;
        let _47_v38;
        _47_v38 = _module.D6.create_DC16(p1);
        let _index4 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_46_v37).length));
        (_46_v37)[_index4] = _47_v38;
        let _index5 = _module.__default.safeIndex(new BigNumber(761), new BigNumber((_46_v37).length));
        (_46_v37)[_index5] = _module.D6.create_DC16(p1);
        let _48_v39;
        let _nw8 = Array((new BigNumber(9)).toNumber());
        _48_v39 = _nw8;
        let _49_v40;
        _49_v40 = _dafny.MultiSet.fromElements(p1, p1);
        let _50_v41;
        _50_v41 = _module.D2.create_DC6(_49_v40);
        let _51_v42;
        let _nw9 = new _module.C0();
        _nw9.__ctor(_50_v41, p1, _23_v15);
        _51_v42 = _nw9;
        let _52_v43;
        _52_v43 = _module.D6.create_DC15(_51_v42);
        let _index6 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_48_v39).length));
        (_48_v39)[_index6] = _52_v43;
        let _index7 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_48_v39).length));
        (_48_v39)[_index7] = _52_v43;
        (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_24_v16, _24_v16), _24_v16);
        let _53_v44;
        let _nw10 = new _module.C0();
        _nw10.__ctor(_module.D2.create_DC6(_49_v40), new BigNumber((_dafny.Seq.UnicodeFromString("uhrsgbmu")).length), _23_v15);
        _53_v44 = _nw10;
        let _54_v45;
        _54_v45 = _dafny.Seq.of(_53_v44);
        let _55_v46;
        let _nw11 = Array((new BigNumber(4)).toNumber());
        _nw11[(_dafny.ZERO).toNumber()] = _54_v45;
        _nw11[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_54_v45, _54_v45);
        _nw11[(new BigNumber(2)).toNumber()] = _54_v45;
        _nw11[(new BigNumber(3)).toNumber()] = _54_v45;
        _55_v46 = _nw11;
        let _index8 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_55_v46).length));
        (_55_v46)[_index8] = _54_v45;
        let _56_v47;
        _56_v47 = _dafny.Seq.of(_53_v44);
        let _index9 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_55_v46).length));
        (_55_v46)[_index9] = ((((_51_v42).f18) ? (_56_v47) : (_54_v45)));
      }
      let _57_v48;
      let _nw12 = new _module.C1();
      _nw12.__ctor(_23_v15, p1, (_dafny.ZERO).minus(p1), _23_v15);
      _57_v48 = _nw12;
      let _58_v49;
      _58_v49 = _module.D4.create_DC9(_57_v48);
      let _59_v50;
      _59_v50 = _dafny.Seq.of((_58_v49).dtor_cf13, _57_v48);
      let _60_v51;
      let _nw13 = Array((new BigNumber(20)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _57_v48;
      _nw13[(_dafny.ONE).toNumber()] = _57_v48;
      _nw13[(new BigNumber(2)).toNumber()] = (_59_v50)[_module.__default.safeIndex(p1, new BigNumber((_59_v50).length))];
      _nw13[(new BigNumber(3)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(4)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(5)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(6)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(7)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(8)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(9)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(10)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(11)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(12)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(13)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(14)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(15)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(16)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(17)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(18)).toNumber()] = _57_v48;
      _nw13[(new BigNumber(19)).toNumber()] = _57_v48;
      _60_v51 = _nw13;
      let _61_v52;
      _61_v52 = _dafny.Seq.of(_60_v51, _60_v51, _60_v51);
      let _62_v53;
      _62_v53 = _module.D6.create_DC16(p1);
      let _63_v54;
      _63_v54 = _dafny.MultiSet.fromElements(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(507)), ((_64_p0) => function (_65_i3) {
        return _64_p0;
      })(p0))).length));
      let _66_v55;
      _66_v55 = _module.D2.create_DC6(_63_v54);
      let _67_v56;
      let _nw14 = new _module.C0();
      _nw14.__ctor(_66_v55, p1, _23_v15);
      _67_v56 = _nw14;
      let _68_v57;
      _68_v57 = _dafny.Seq.of(_67_v56);
      let _69_v58;
      _69_v58 = _dafny.Seq.of(new BigNumber(479), new BigNumber((_dafny.Set.fromElements((_62_v53).dtor_cf22, new BigNumber((_68_v57).length))).length));
      let _rhs7 = (_61_v52)[_module.__default.safeIndex(((_69_v58)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_69_v58).length))]).plus(p0), new BigNumber((_61_v52).length))];
      let _rhs8 = p0;
      let _rhs9 = (_67_v56).f23;
      let _rhs10 = _module.__default.fm1(p0, (_57_v48.f21) || (_23_v15), (_dafny.MultiSet.fromElements(_57_v48.f21)).Intersect(_dafny.MultiSet.fromElements(!(_23_v15))), globalState);
      let _lhs7 = globalState;
      let _lhs8 = globalState;
      let _lhs9 = globalState;
      _60_v51 = _rhs7;
      _lhs7.f9 = _rhs8;
      _lhs8.f5 = _rhs9;
      _lhs9.f13 = _rhs10;
      if ((new BigNumber(36)).isLessThanOrEqualTo(_module.__default.safeDivisionInt((_67_v56).f23, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm9(_57_v48.f21, true, globalState)).length))))) {
        let _70_v59;
        _70_v59 = _dafny.Set.fromElements(_23_v15, _57_v48.f21, _23_v15);
        let _71_v60;
        _71_v60 = _module.D9.create_DC23(_70_v59);
        let _pat_let_tv0 = _57_v48;
        let _pat_let_tv1 = _57_v48;
        _70_v59 = (_70_v59).Union((function (_pat_let0_0) {
          return function (_72_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_73_dt__update_hcf32_h0) {
                return _module.D9.create_DC23(_73_dt__update_hcf32_h0);
              }(_pat_let1_0);
            }(_dafny.Set.fromElements(_pat_let_tv0.f21, _pat_let_tv1.f21));
          }(_pat_let0_0);
        }(_71_v60)).dtor_cf32);
        (globalState).f3 = _module.__default.safeModuloInt(_module.__default.fm0(new BigNumber(929), p1, new BigNumber(503), p0, globalState), new BigNumber(263));
        let _74_v61;
        let _nw15 = new _module.C1();
        _nw15.__ctor(_57_v48.f21, p0, new BigNumber(537), !((_57_v48).fm5(false, new BigNumber(-758), globalState)));
        _74_v61 = _nw15;
        let _75_v62;
        _75_v62 = _dafny.Map.Empty.slice().updateUnsafe(_67_v56,_70_v59);
        let _76_v63;
        _76_v63 = _module.D1.create_DC3((new BigNumber((_75_v62).length)).plus(new BigNumber((_dafny.Seq.of(_74_v61.f21, _57_v48.f21, _74_v61.f21)).length)));
        let _source0 = _76_v63;
        if (_source0.is_DC3) {
          let _77___mcc_h0 = (_source0).cf5;
          let _78_cf5 = _77___mcc_h0;
          let _79_v64;
          _79_v64 = _dafny.Map.Empty.slice().updateUnsafe((_67_v56).f23,_74_v61.f21);
          _79_v64 = (_79_v64).update(p0, _74_v61.f21);
          let _80_v65;
          let _nw16 = Array((new BigNumber(6)).toNumber()).fill(false);
          _80_v65 = _nw16;
          let _81_v66;
          let _nw17 = Array((_dafny.ONE).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = (_67_v56).f23;
          _81_v66 = _nw17;
          let _82_v67;
          _82_v67 = _dafny.Map.Empty.slice().updateUnsafe(_80_v65,_81_v66);
          let _83_v68;
          _83_v68 = _module.D7.create_DC18(_82_v67);
          _83_v68 = _module.D7.create_DC18((_82_v67).Merge(_82_v67));
          _23_v15 = !(((_67_v56).f23).isLessThanOrEqualTo((_67_v56).f23));
          (_74_v61).f21 = !(_57_v48.f21);
        } else if (_source0.is_DC4) {
          let _84___mcc_h1 = (_source0).cf6;
          let _85___mcc_h2 = (_source0).cf7;
          let _86___mcc_h3 = (_source0).cf8;
          let _87_cf8 = _86___mcc_h3;
          let _88_cf7 = _85___mcc_h2;
          let _89_cf6 = _84___mcc_h1;
          let _90_v70;
          let _init0 = ((_91_cf7, _92_v48, _93_v56) => function (_94_i4) {
            return ((_91_cf7) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nqlnui")).length),_92_v48.f21)) : (function () {
              let _coll2 = new _dafny.Map();
              for (const _compr_2 of _dafny.IntegerRange(new BigNumber(418), new BigNumber(981))) {
                let _95_v69 = _compr_2;
                if (((new BigNumber(418)).isLessThanOrEqualTo(_95_v69)) && ((_95_v69).isLessThan(new BigNumber(981)))) {
                  _coll2.push([(_95_v69).minus((_93_v56).f23),_92_v48.f21]);
                }
              }
              return _coll2;
            }()));
          })(_88_cf7, _57_v48, _67_v56);
          let _nw18 = Array((new BigNumber(29)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw18.length); _i0_0++) {
            _nw18[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _90_v70 = _nw18;
          let _96_v71;
          _96_v71 = _dafny.Map.Empty.slice().updateUnsafe((_67_v56).f23,_57_v48.f21);
          let _index10 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_90_v70).length));
          (_90_v70)[_index10] = _96_v71;
          let _index11 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_90_v70).length));
          (_90_v70)[_index11] = _96_v71;
          let _97_v72;
          _97_v72 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,_57_v48.f21);
          let _98_v73;
          let _nw19 = new _module.C1();
          _nw19.__ctor(_57_v48.f21, (new BigNumber(615)).multipliedBy(p1), new BigNumber((_97_v72).length), _57_v48.f21);
          _98_v73 = _nw19;
          (globalState).f9 = (_67_v56).f23;
          let _99_v74;
          _99_v74 = _dafny.Seq.UnicodeFromString("qaky");
          (globalState).f4 = _99_v74;
        } else {
          let _100___mcc_h4 = (_source0).cf4;
          let _101_cf4 = _100___mcc_h4;
          let _102_v75;
          let _103_v76;
          let _104_v77;
          let _105_v78;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_57_v48).m1(_57_v48.f21, _74_v61.f21, globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _102_v75 = _out0;
          _103_v76 = _out1;
          _104_v77 = _out2;
          _105_v78 = _out3;
          let _106_v79;
          let _nw20 = Array((new BigNumber(15)).toNumber()).fill(false);
          _106_v79 = _nw20;
          let _107_v80;
          _107_v80 = _dafny.Seq.UnicodeFromString("p");
          let _108_v81;
          _108_v81 = _dafny.Map.Empty.slice().updateUnsafe(_104_v77,new BigNumber((_24_v16).length));
          let _109_v82;
          _109_v82 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,_108_v81);
          let _110_v83;
          let _nw21 = Array((new BigNumber(17)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = (_67_v56).f23;
          _nw21[(_dafny.ONE).toNumber()] = p0;
          _nw21[(new BigNumber(2)).toNumber()] = (_67_v56).f23;
          _nw21[(new BigNumber(3)).toNumber()] = new BigNumber((_107_v80).length);
          _nw21[(new BigNumber(4)).toNumber()] = (_67_v56).f23;
          _nw21[(new BigNumber(5)).toNumber()] = new BigNumber((_109_v82).length);
          _nw21[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_67_v56).f23);
          _nw21[(new BigNumber(7)).toNumber()] = new BigNumber((_69_v58).length);
          _nw21[(new BigNumber(8)).toNumber()] = (_67_v56).f23;
          _nw21[(new BigNumber(9)).toNumber()] = _104_v77;
          _nw21[(new BigNumber(10)).toNumber()] = (_67_v56).f23;
          _nw21[(new BigNumber(11)).toNumber()] = p0;
          _nw21[(new BigNumber(12)).toNumber()] = p0;
          _nw21[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), function (_111_i5) {
            return new _dafny.CodePoint('p'.codePointAt(0));
          })).length);
          _nw21[(new BigNumber(14)).toNumber()] = _104_v77;
          _nw21[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_67_v56).f23);
          _nw21[(new BigNumber(16)).toNumber()] = _104_v77;
          _110_v83 = _nw21;
          let _112_v84;
          _112_v84 = _dafny.Map.Empty.slice().updateUnsafe(_106_v79,_110_v83);
          let _113_v85;
          _113_v85 = _module.D7.create_DC18(_112_v84);
          _113_v85 = _113_v85;
          let _114_v86;
          _114_v86 = _module.D6.create_DC17((_67_v56).f23, true);
          let _115_v87;
          _115_v87 = _dafny.Set.fromElements(_114_v86);
          let _116_v88;
          _116_v88 = _dafny.Map.Empty.slice().updateUnsafe(_104_v77,_115_v87);
          _116_v88 = (_116_v88).update(_104_v77, _dafny.Set.fromElements(_114_v86, _module.__default.fm17(globalState), _module.D6.create_DC17(p1, _23_v15)));
          let _117_v89;
          _117_v89 = _dafny.Map.Empty.slice().updateUnsafe(_110_v83,_107_v80);
          (globalState).f5 = new BigNumber(((_117_v89).Merge((_117_v89).update(_110_v83, _107_v80))).length);
        }
        r0 = (_67_v56).f23;
      } else {
        let _118_v90;
        let _nw22 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _118_v90 = _nw22;
        let _index12 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length));
        (_118_v90)[_index12] = new BigNumber(250);
        let _index13 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length));
        (_118_v90)[_index13] = new BigNumber((_24_v16).length);
        _57_v48 = _57_v48;
        let _119_v91;
        _119_v91 = _dafny.Map.Empty.slice().updateUnsafe((_67_v56).f23,p1);
        let _120_v93;
        _120_v93 = _module.D9.create_DC24(!(_119_v91).equals(function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(665), new BigNumber(12))) {
    let _121_v92 = _compr_3;
    if (((new BigNumber(665)).isLessThanOrEqualTo(_121_v92)) && ((_121_v92).isLessThan(new BigNumber(12)))) {
      _coll3.push([(_121_v92).minus(new BigNumber(330)),(_67_v56).f23]);
    }
  }
  return _coll3;
}()), (_dafny.ZERO).minus((_67_v56).f23), !(_57_v48.f21));
        let _source1 = _120_v93;
        if (_source1.is_DC24) {
          let _122___mcc_h5 = (_source1).cf33;
          let _123___mcc_h6 = (_source1).cf34;
          let _124___mcc_h7 = (_source1).cf35;
          let _125_cf35 = _124___mcc_h7;
          let _126_cf34 = _123___mcc_h6;
          let _127_cf33 = _122___mcc_h5;
          let _128_v94;
          let _init1 = ((_129_v48) => function (_130_i6) {
            return _129_v48.f21;
          })(_57_v48);
          let _nw23 = Array((new BigNumber(27)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw23.length); _i0_1++) {
            _nw23[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _128_v94 = _nw23;
          let _131_v95;
          _131_v95 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,_128_v94);
          let _132_v96;
          let _nw24 = Array((new BigNumber(4)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _128_v94;
          _nw24[(_dafny.ONE).toNumber()] = _128_v94;
          _nw24[(new BigNumber(2)).toNumber()] = (((_131_v95).contains(_57_v48.f21)) ? ((_131_v95).get(_57_v48.f21)) : (_128_v94));
          _nw24[(new BigNumber(3)).toNumber()] = _128_v94;
          _132_v96 = _nw24;
          let _index14 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_132_v96).length));
          (_132_v96)[_index14] = _128_v94;
          let _133_v97;
          _133_v97 = _dafny.Seq.of((_module.D10.create_DC25(_128_v94)).dtor_cf36);
          let _134_v98;
          _134_v98 = _dafny.Set.fromElements(_57_v48.f21);
          let _index15 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_132_v96).length));
          (_132_v96)[_index15] = (_133_v97)[_module.__default.safeIndex(new BigNumber((_134_v98).length), new BigNumber((_133_v97).length))];
          let _135_v99;
          _135_v99 = _dafny.Seq.UnicodeFromString("qggfdxe");
          let _136_v100;
          _136_v100 = _dafny.MultiSet.fromElements(true);
          let _rhs11 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(663)), ((_137_p0) => function (_138_i7) {
            return _137_p0;
          })(p0));
          let _rhs12 = (_dafny.ZERO).minus(_126_cf34);
          let _rhs13 = (_24_v16)[_module.__default.safeIndex(((_118_v90)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length))]).minus(new BigNumber((_135_v99).length)), new BigNumber((_24_v16).length))];
          let _rhs14 = (_118_v90)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length))];
          let _rhs15 = (((_136_v100).contains((new BigNumber((_dafny.Seq.UnicodeFromString("ursmty")).length)).isLessThan((_67_v56).f23))) ? ((_136_v100).get((new BigNumber((_dafny.Seq.UnicodeFromString("ursmty")).length)).isLessThan((_67_v56).f23))) : (((_67_v56).f23).multipliedBy((_67_v56).f23)));
          let _lhs10 = globalState;
          let _lhs11 = _57_v48;
          let _lhs12 = globalState;
          _69_v58 = _rhs11;
          _lhs10.f3 = _rhs12;
          _lhs11.f21 = _rhs13;
          _lhs12.f3 = _rhs14;
          _126_cf34 = _rhs15;
          let _139_v101;
          _139_v101 = _dafny.Map.Empty.slice().updateUnsafe(_127_cf33,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_67_v56).f23)).cardinality()),_57_v48)).length));
          _125_cf35 = ((new BigNumber((_139_v101).length)).minus(new BigNumber((_136_v100).cardinality()))).isLessThanOrEqualTo((_118_v90)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length))]);
          let _140_v102;
          _140_v102 = _dafny.Map.Empty.slice().updateUnsafe((_118_v90)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length))],_135_v99);
          let _index16 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length));
          let _rhs16 = _132_v96;
          let _rhs17 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_133_v97, _133_v97)).length), (_dafny.ZERO).minus(new BigNumber(((_140_v102).Merge(function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(410), new BigNumber(451))) {
              let _141_v103 = _compr_4;
              if (((new BigNumber(410)).isLessThanOrEqualTo(_141_v103)) && ((_141_v103).isLessThan(new BigNumber(451)))) {
                _coll4.push([(_141_v103).plus(new BigNumber(169)),_135_v99]);
              }
            }
            return _coll4;
          }())).length)));
          let _lhs13 = _118_v90;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length));
          _132_v96 = _rhs16;
          _lhs13[_lhs14] = _rhs17;
        } else {
          let _142___mcc_h8 = (_source1).cf32;
          let _143_cf32 = _142___mcc_h8;
          let _144_v104;
          let _init2 = ((_145_v56, _146_v48) => function (_147_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC13((_145_v56).f23),_146_v48.f21);
          })(_67_v56, _57_v48);
          let _nw25 = Array((new BigNumber(12)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw25.length); _i0_2++) {
            _nw25[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _144_v104 = _nw25;
          let _148_v105;
          _148_v105 = _module.D5.create_DC13(p0);
          let _149_v106;
          _149_v106 = _dafny.Map.Empty.slice().updateUnsafe(_148_v105,_57_v48.f21);
          let _index17 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_144_v104).length));
          (_144_v104)[_index17] = _149_v106;
          let _150_v108;
          _150_v108 = _dafny.Seq.of(_148_v105, _module.__default.fm18(_23_v15, globalState), _148_v105);
          let _pat_let_tv2 = _67_v56;
          let _index18 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_144_v104).length));
          let _rhs18 = _57_v48.f21;
          let _rhs19 = _57_v48.f21;
          let _rhs20 = (function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of (_150_v108).Elements) {
              let _151_v107 = _compr_5;
              if (_dafny.Seq.contains(_150_v108, _151_v107)) {
                _coll5.push([_151_v107,_23_v15]);
              }
            }
            return _coll5;
          }()).update(function (_pat_let2_0) {
            return function (_152_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_153_dt__update_hcf19_h0) {
                  return _module.D5.create_DC13(_153_dt__update_hcf19_h0);
                }(_pat_let3_0);
              }((_pat_let_tv2).f23);
            }(_pat_let2_0);
          }(_148_v105), _57_v48.f21);
          let _lhs15 = globalState;
          let _lhs16 = _144_v104;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(534), new BigNumber((_144_v104).length));
          _lhs15.f13 = _rhs18;
          _23_v15 = _rhs19;
          _lhs16[_lhs17] = _rhs20;
          let _nw26 = new _module.C1();
          _nw26.__ctor(_57_v48.f21, (_67_v56).f23, (_67_v56).f23, ((_57_v48.f21) ? (_57_v48.f21) : (_57_v48.f21)));
          _57_v48 = _nw26;
          (globalState).f13 = _57_v48.f21;
          (globalState).f9 = (_67_v56).f23;
        }
        let _154_v109;
        _154_v109 = _dafny.Set.fromElements(_67_v56, _67_v56, _67_v56);
        let _index19 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length));
        (_118_v90)[_index19] = (_dafny.ZERO).minus(new BigNumber((_154_v109).length));
        if (_23_v15) {
          let _155_v110;
          _155_v110 = _dafny.Map.Empty.slice().updateUnsafe((_67_v56).f23,_57_v48);
          _155_v110 = (_155_v110).Merge(_155_v110);
          let _156_v111;
          _156_v111 = _dafny.Map.Empty.slice().updateUnsafe(_23_v15,p0);
          _156_v111 = _156_v111;
          (globalState).f13 = _57_v48.f21;
          let _157_v112;
          let _nw27 = new _module.C1();
          _nw27.__ctor(_57_v48.f21, (_67_v56).f23, (_67_v56).f23, true);
          _157_v112 = _nw27;
          let _158_v113;
          _158_v113 = new _dafny.CodePoint('t'.codePointAt(0));
          let _159_v114;
          _159_v114 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,_57_v48);
          let _160_v115;
          _160_v115 = _dafny.MultiSet.fromElements(_118_v90, _118_v90);
          let _rhs21 = _157_v112;
          let _rhs22 = !(_57_v48.f21) || ((_157_v112).f18);
          let _rhs23 = _158_v113;
          let _rhs24 = (((_159_v114).contains((_dafny.MultiSet.fromElements(_118_v90)).IsProperSubsetOf(_160_v115))) ? ((_159_v114).get((_dafny.MultiSet.fromElements(_118_v90)).IsProperSubsetOf(_160_v115))) : (_57_v48));
          let _lhs18 = globalState;
          _157_v112 = _rhs21;
          _lhs18.f13 = _rhs22;
          _158_v113 = _rhs23;
          _57_v48 = _rhs24;
          let _161_v116;
          let _nw28 = Array((new BigNumber(6)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _158_v113;
          _nw28[(_dafny.ONE).toNumber()] = _158_v113;
          _nw28[(new BigNumber(2)).toNumber()] = _158_v113;
          _nw28[(new BigNumber(3)).toNumber()] = _158_v113;
          _nw28[(new BigNumber(4)).toNumber()] = _158_v113;
          _nw28[(new BigNumber(5)).toNumber()] = _158_v113;
          _161_v116 = _nw28;
          let _index20 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_161_v116).length));
          (_161_v116)[_index20] = _158_v113;
          let _162_v117;
          _162_v117 = _module.D2.create_DC6(_dafny.MultiSet.fromElements(new BigNumber(789)));
          let _163_v118;
          _163_v118 = _module.D2.create_DC7(_162_v117);
          let _index21 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_161_v116).length));
          (_161_v116)[_index21] = _module.__default.fm15(_163_v118, (_57_v48).fm5((_157_v112).f18, _module.__default.fm0(new BigNumber(760), p0, (_67_v56).f23, new BigNumber(189), globalState), globalState), globalState);
        } else {
          let _164_v119;
          _164_v119 = new _dafny.CodePoint('k'.codePointAt(0));
          let _165_v120;
          _165_v120 = _module.D2.create_DC6(_63_v54);
          let _166_v121;
          _166_v121 = _module.D2.create_DC7(_module.D2.create_DC7(_165_v120));
          let _167_v122;
          let _nw29 = Array((new BigNumber(6)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _164_v119;
          _nw29[(_dafny.ONE).toNumber()] = _164_v119;
          _nw29[(new BigNumber(2)).toNumber()] = _module.__default.fm15(_166_v121, _57_v48.f21, globalState);
          _nw29[(new BigNumber(3)).toNumber()] = _164_v119;
          _nw29[(new BigNumber(4)).toNumber()] = ((_57_v48.f21) ? (_164_v119) : (_164_v119));
          _nw29[(new BigNumber(5)).toNumber()] = _164_v119;
          _167_v122 = _nw29;
          let _index22 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_167_v122).length));
          (_167_v122)[_index22] = _164_v119;
          let _168_v123;
          _168_v123 = _dafny.Map.Empty.slice().updateUnsafe(true,_57_v48);
          let _169_v124;
          let _nw30 = new _module.C1();
          _nw30.__ctor((_57_v48).fm5(_57_v48.f21, p1, globalState), new BigNumber((_dafny.MultiSet.FromArray(_24_v16)).cardinality()), new BigNumber((_168_v123).length), _57_v48.f21);
          _169_v124 = _nw30;
          let _170_v125;
          _170_v125 = _module.D10.create_DC26(_169_v124);
          let _171_v126;
          let _nw31 = Array((new BigNumber(4)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _170_v125;
          _nw31[(_dafny.ONE).toNumber()] = _module.D10.create_DC26(_169_v124);
          _nw31[(new BigNumber(2)).toNumber()] = _170_v125;
          _nw31[(new BigNumber(3)).toNumber()] = _module.D10.create_DC26(_169_v124);
          _171_v126 = _nw31;
          let _172_v127;
          let _nw32 = Array((new BigNumber(3)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = _118_v90;
          _nw32[(_dafny.ONE).toNumber()] = _118_v90;
          _nw32[(new BigNumber(2)).toNumber()] = _118_v90;
          _172_v127 = _nw32;
          let _index23 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_172_v127).length));
          (_172_v127)[_index23] = _118_v90;
          let _173_v128;
          _173_v128 = _module.D11.create_DC29(_118_v90);
          let _index24 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_167_v122).length));
          let _index25 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_172_v127).length));
          let _rhs25 = _69_v58;
          let _rhs26 = _164_v119;
          let _rhs27 = _171_v126;
          let _rhs28 = (_173_v128).dtor_cf40;
          let _lhs19 = _167_v122;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_167_v122).length));
          let _lhs21 = _172_v127;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_172_v127).length));
          _69_v58 = _rhs25;
          _lhs19[_lhs20] = _rhs26;
          _171_v126 = _rhs27;
          _lhs21[_lhs22] = _rhs28;
          let _174_v129;
          let _175_v130;
          let _176_v131;
          let _177_v132;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_169_v124).m1((_24_v16)[_module.__default.safeIndex(_169_v124.f20, new BigNumber((_24_v16).length))], _23_v15, globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _174_v129 = _out4;
          _175_v130 = _out5;
          _176_v131 = _out6;
          _177_v132 = _out7;
          let _178_v133;
          _178_v133 = _dafny.MultiSet.fromElements(_57_v48.f21, (_169_v124).f18);
          let _179_v134;
          _179_v134 = _dafny.Seq.UnicodeFromString("rihyj");
          let _180_v135;
          _180_v135 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_178_v133).cardinality()), new BigNumber(600), new BigNumber(-288), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), ((_181_v119) => function (_182_i9) {
            return _181_v119;
          })(_164_v119))).length), globalState),_dafny.Seq.Concat(_dafny.Seq.update(_179_v134, _module.__default.safeIndex((_118_v90)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_118_v90).length))], new BigNumber((_179_v134).length)), (_167_v122)[_module.__default.safeIndex(new BigNumber(751), new BigNumber((_167_v122).length))]), _179_v134));
          _180_v135 = (_180_v135).update((((_178_v133).contains(_57_v48.f21)) ? ((_178_v133).get(_57_v48.f21)) : (_176_v131)), _179_v134);
          let _183_v136;
          let _nw33 = new _module.C0();
          _nw33.__ctor(_67_v56.f22, p1, _175_v130);
          _183_v136 = _nw33;
          let _184_v137;
          _184_v137 = _module.D6.create_DC15(_183_v136);
          let _185_v138;
          _185_v138 = _dafny.Set.fromElements((_184_v137).dtor_cf21);
          let _186_v139;
          _186_v139 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_185_v138).length),_57_v48.f21);
          (globalState).f3 = ((_dafny.ZERO).minus((_69_v58)[_module.__default.safeIndex(new BigNumber((_186_v139).length), new BigNumber((_69_v58).length))])).plus(_176_v131);
          let _187_v140;
          let _init3 = ((_188_v124) => function (_189_i10) {
            return !(!((_188_v124).f18));
          })(_169_v124);
          let _nw34 = Array((new BigNumber(15)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw34.length); _i0_3++) {
            _nw34[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _187_v140 = _nw34;
          let _190_v141;
          _190_v141 = _dafny.Map.Empty.slice().updateUnsafe(_187_v140,new BigNumber((_119_v91).length));
          let _191_v142;
          _191_v142 = _dafny.Map.Empty.slice().updateUnsafe(false,_190_v141);
          let _rhs29 = (_24_v16)[_module.__default.safeIndex(_169_v124.f20, new BigNumber((_24_v16).length))];
          let _rhs30 = _24_v16;
          let _rhs31 = (_190_v141).equals(((((_191_v142).contains(_57_v48.f21)) ? ((_191_v142).get(_57_v48.f21)) : (_dafny.Map.Empty.slice().updateUnsafe(_187_v140,_177_v132)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_187_v140,_module.__default.fm0(_176_v131, (_67_v56).f23, new BigNumber((_24_v16).length), _169_v124.f20, globalState))));
          let _rhs32 = (_183_v136).f18;
          let _rhs33 = (_174_v129).minus((_67_v56).f23);
          let _lhs23 = _57_v48;
          let _lhs24 = _57_v48;
          let _lhs25 = globalState;
          let _lhs26 = globalState;
          _lhs23.f21 = _rhs29;
          _24_v16 = _rhs30;
          _lhs24.f21 = _rhs31;
          _lhs25.f13 = _rhs32;
          _lhs26.f9 = _rhs33;
        }
      }
      let _192_v143;
      _192_v143 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(366));
      let _193_v144;
      _193_v144 = _dafny.MultiSet.fromElements(_192_v143, _192_v143);
      let _194_v145;
      _194_v145 = _module.D7.create_DC20(_193_v144);
      let _pat_let_tv3 = _57_v48;
      let _pat_let_tv4 = _193_v144;
      let _pat_let_tv5 = _193_v144;
      (_57_v48).f21 = function (_source2) {
        if (_source2.is_DC19) {
          let _195___mcc_h9 = (_source2).cf26;
          let _196___mcc_h10 = (_source2).cf27;
          let _197___mcc_h11 = (_source2).cf28;
          let _198_cf28 = _197___mcc_h11;
          let _199_cf27 = _196___mcc_h10;
          let _200_cf26 = _195___mcc_h9;
          return _198_cf28;
        } else if (_source2.is_DC20) {
          let _201___mcc_h12 = (_source2).cf29;
          let _202_cf29 = _201___mcc_h12;
          return true;
        } else if (_source2.is_DC18) {
          let _203___mcc_h13 = (_source2).cf25;
          let _204_cf25 = _203___mcc_h13;
          return _pat_let_tv3.f21;
        } else {
          let _205___mcc_h14 = (_source2).cf30;
          let _206_cf30 = _205___mcc_h14;
          return (_pat_let_tv4).IsSubsetOf(_pat_let_tv5);
        }
      }(_194_v145);
      let _207_v146;
      _207_v146 = _dafny.Seq.UnicodeFromString("w");
      if (!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ev"), _dafny.Seq.UnicodeFromString("xlhryx")), _207_v146))) {
        if (_57_v48.f21) {
          let _208_v147;
          _208_v147 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), !(!(_57_v48.f21)), _dafny.MultiSet.fromElements(_23_v15), globalState),new BigNumber((_69_v58).length));
          let _209_v148;
          let _nw35 = new _module.C1();
          _nw35.__ctor(_23_v15, p1, (_69_v58)[_module.__default.safeIndex((_67_v56).f23, new BigNumber((_69_v58).length))], false);
          _209_v148 = _nw35;
          let _210_v149;
          _210_v149 = _module.D10.create_DC26(_209_v148);
          let _211_v150;
          _211_v150 = _dafny.Map.Empty.slice().updateUnsafe(_208_v147,_210_v149);
          _211_v150 = (_211_v150).update(_208_v147, _210_v149);
          let _212_v151;
          _212_v151 = _module.D4.create_DC10(_57_v48.f21, (((_192_v143).contains((_dafny.ZERO).minus(p1))) ? ((_192_v143).get((_dafny.ZERO).minus(p1))) : (new BigNumber(-796))), p1);
          let _213_v152;
          _213_v152 = _module.D1.create_DC3((_212_v151).dtor_cf15);
          let _214_v153;
          _214_v153 = _dafny.MultiSet.fromElements(_213_v152, _213_v152);
          let _215_v154;
          _215_v154 = _module.D9.create_DC24(false, (_209_v148).f19, _57_v48.f21);
          let _rhs34 = ((_57_v48.f21) ? (((_209_v148).f19).minus(p1)) : ((_209_v148).f19));
          let _rhs35 = (_215_v154).dtor_cf35;
          let _rhs36 = _module.__default.fm19(_209_v148.f20, new BigNumber((_dafny.Seq.Concat(_69_v58, _69_v58)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-197)), ((_216_v56) => function (_217_i11) {
            return (_216_v56).f23;
          })(_67_v56))).length), globalState);
          let _rhs37 = _57_v48.f21;
          let _rhs38 = _57_v48.f21;
          let _lhs27 = globalState;
          let _lhs28 = _57_v48;
          let _lhs29 = globalState;
          let _lhs30 = _57_v48;
          _lhs27.f3 = _rhs34;
          _lhs28.f21 = _rhs35;
          _214_v153 = _rhs36;
          _lhs29.f13 = _rhs37;
          _lhs30.f21 = _rhs38;
          _57_v48 = _57_v48;
          let _218_v155;
          let _init4 = function (_219_i12) {
            return _module.__default.safeDivisionInt(_219_i12, new BigNumber(110));
          };
          let _nw36 = Array((new BigNumber(15)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw36.length); _i0_4++) {
            _nw36[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _218_v155 = _nw36;
          let _index26 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_218_v155).length));
          (_218_v155)[_index26] = p0;
          let _index27 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_218_v155).length));
          (_218_v155)[_index27] = (new BigNumber(-685)).plus(_209_v148.f20);
          (globalState).f13 = (_209_v148).fm5(_57_v48.f21, (_209_v148).f19, globalState);
        } else {
          _57_v48 = _57_v48;
          (globalState).f13 = _23_v15;
          let _220_v156;
          _220_v156 = _dafny.MultiSet.fromElements(_23_v15, _57_v48.f21);
          let _221_v157;
          let _nw37 = new _module.C1();
          _nw37.__ctor((_220_v156).equals(_220_v156), (_dafny.ZERO).minus(((_67_v56).f23).multipliedBy(new BigNumber((_207_v146).length))), p0, ((_57_v48.f21) ? (_57_v48.f21) : (_57_v48.f21)));
          _221_v157 = _nw37;
          let _222_v158;
          let _nw38 = Array((new BigNumber(3)).toNumber()).fill([]);
          _222_v158 = _nw38;
          _222_v158 = _222_v158;
          let _223_v159;
          _223_v159 = new _dafny.CodePoint('k'.codePointAt(0));
          let _224_v160;
          _224_v160 = _dafny.Set.fromElements(_57_v48, _221_v157, _57_v48, _57_v48, _57_v48);
          _223_v159 = ((_57_v48.f21) ? (_223_v159) : (((_221_v157.f21) ? ((_207_v146)[_module.__default.safeIndex(new BigNumber((_224_v160).length), new BigNumber((_207_v146).length))]) : (_223_v159))));
        }
        (globalState).f9 = (((new BigNumber(263)).isLessThan(new BigNumber((_207_v146).length))) ? (p0) : ((_67_v56).f23));
        if ((p1).isLessThanOrEqualTo(new BigNumber(914))) {
          let _225_v161;
          let _226_v162;
          let _227_v163;
          let _228_v164;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = (_57_v48).m1((_24_v16)[_module.__default.safeIndex((_67_v56).f23, new BigNumber((_24_v16).length))], true, globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _225_v161 = _out8;
          _226_v162 = _out9;
          _227_v163 = _out10;
          _228_v164 = _out11;
          let _229_v165;
          _229_v165 = new _dafny.CodePoint('x'.codePointAt(0));
          _229_v165 = _229_v165;
          let _230_v166;
          _230_v166 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,new BigNumber((_207_v146).length));
          let _231_v167;
          _231_v167 = _dafny.MultiSet.fromElements(_57_v48.f21, _23_v15);
          let _232_v168;
          let _nw39 = new _module.C1();
          _nw39.__ctor(_module.__default.fm1(new BigNumber((_230_v166).length), _57_v48.f21, _231_v167, globalState), (_dafny.ZERO).minus((_67_v56).f23), p0, _module.__default.fm1((_67_v56).f23, _226_v162, _231_v167, globalState));
          _232_v168 = _nw39;
          let _233_v169;
          let _init5 = ((_234_v164, _235_v15) => function (_236_i13) {
            return (_236_i13).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_234_v164,_235_v15)).length));
          })(_228_v164, _23_v15);
          let _nw40 = Array((new BigNumber(6)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
            _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _233_v169 = _nw40;
          let _index28 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_233_v169).length));
          (_233_v169)[_index28] = (_67_v56).f23;
          let _237_v170;
          _237_v170 = _dafny.MultiSet.fromElements(_69_v58);
          let _238_v171;
          let _init6 = function (_239_i14) {
            return false;
          };
          let _nw41 = Array((new BigNumber(7)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw41.length); _i0_6++) {
            _nw41[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _238_v171 = _nw41;
          let _240_v172;
          _240_v172 = _dafny.Map.Empty.slice().updateUnsafe(_238_v171,_57_v48.f21);
          let _241_v173;
          _241_v173 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(_23_v15)), _231_v167, _231_v167, _231_v167, _231_v167);
          let _index29 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_233_v169).length));
          (_233_v169)[_index29] = (new BigNumber((_237_v170).cardinality())).minus(_module.__default.fm0(new BigNumber((_207_v146).length), new BigNumber(785), (_dafny.ZERO).minus(new BigNumber((_240_v172).length)), new BigNumber((_241_v173).length), globalState));
          let _242_v174;
          let _nw42 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Set.Empty);
          _242_v174 = _nw42;
          let _243_v175;
          _243_v175 = _dafny.Set.fromElements((_67_v56).f23);
          let _index30 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_242_v174).length));
          (_242_v174)[_index30] = _243_v175;
          let _index31 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_242_v174).length));
          (_242_v174)[_index31] = (_243_v175).Difference(_243_v175);
        } else {
          let _244_v176;
          let _nw43 = Array((new BigNumber(7)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = !(!((_69_v58)[_module.__default.safeIndex(new BigNumber((_207_v146).length), new BigNumber((_69_v58).length))]).isEqualTo((_67_v56).f23));
          _nw43[(_dafny.ONE).toNumber()] = !(_57_v48.f21);
          _nw43[(new BigNumber(2)).toNumber()] = _57_v48.f21;
          _nw43[(new BigNumber(3)).toNumber()] = _57_v48.f21;
          _nw43[(new BigNumber(4)).toNumber()] = _57_v48.f21;
          _nw43[(new BigNumber(5)).toNumber()] = _23_v15;
          _nw43[(new BigNumber(6)).toNumber()] = _57_v48.f21;
          _244_v176 = _nw43;
          let _index32 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_244_v176).length));
          (_244_v176)[_index32] = !(_57_v48.f21) || (_57_v48.f21);
          let _index33 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_244_v176).length));
          (_244_v176)[_index33] = _23_v15;
          let _245_v177;
          let _nw44 = new _module.C0();
          _nw44.__ctor(_module.D2.create_DC6(_63_v54), (new BigNumber((_207_v146).length)).plus((_67_v56).f23), _57_v48.f21);
          _245_v177 = _nw44;
          let _246_v178;
          _246_v178 = _module.D5.create_DC13((_67_v56).f23);
          let _pat_let_tv6 = _245_v177;
          let _pat_let_tv7 = p0;
          let _pat_let_tv8 = _67_v56;
          let _pat_let_tv9 = _67_v56;
          let _pat_let_tv10 = globalState;
          _246_v178 = function (_pat_let4_0) {
            return function (_247_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_248_dt__update_hcf19_h1) {
                  return _module.D5.create_DC13(_248_dt__update_hcf19_h1);
                }(_pat_let5_0);
              }((_dafny.ZERO).minus(_module.__default.fm0((_pat_let_tv6).f23, _pat_let_tv7, (_pat_let_tv8).f23, (_pat_let_tv9).f23, _pat_let_tv10)));
            }(_pat_let4_0);
          }(_246_v178);
          let _249_v179;
          _249_v179 = _dafny.Seq.of(_245_v177);
          _249_v179 = _249_v179;
          let _250_v180;
          let _251_v181;
          let _252_v182;
          let _253_v183;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = (_57_v48).m1(_23_v15, _57_v48.f21, globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _250_v180 = _out12;
          _251_v181 = _out13;
          _252_v182 = _out14;
          _253_v183 = _out15;
        }
        let _254_v184;
        _254_v184 = _dafny.MultiSet.fromElements(_57_v48.f21);
        let _255_v185;
        _255_v185 = _dafny.Set.fromElements(_57_v48.f21, _23_v15);
        let _256_v186;
        let _nw45 = new _module.C1();
        _nw45.__ctor(!((_module.__default.fm1((_67_v56).f23, _57_v48.f21, _254_v184, globalState)) === (_57_v48.f21)), (new BigNumber(366)).plus(p1), (_67_v56).f23, (_255_v185).IsProperSubsetOf(_255_v185));
        _256_v186 = _nw45;
        let _257_v187;
        _257_v187 = _module.D6.create_DC17((((_63_v54).contains(new BigNumber(123))) ? ((_63_v54).get(new BigNumber(123))) : ((_67_v56).f23)), _256_v186.f21);
        let _258_v188;
        let _nw46 = new _module.C1();
        _nw46.__ctor(_57_v48.f21, (_257_v187).dtor_cf23, p1, (new BigNumber((_207_v146).length)).isEqualTo(p0));
        _258_v188 = _nw46;
      } else {
        _23_v15 = (_57_v48.f21) || (true);
        if (_57_v48.f21) {
          let _259_v189;
          _259_v189 = _dafny.MultiSet.fromElements(_57_v48.f21);
          let _260_v190;
          let _nw47 = new _module.C1();
          _nw47.__ctor(_57_v48.f21, _module.__default.safeModuloInt(new BigNumber((_259_v189).cardinality()), p0), (_67_v56).f23, !(_57_v48.f21));
          _260_v190 = _nw47;
          (_57_v48).f21 = false;
          let _261_v191;
          _261_v191 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,p1);
          let _262_v192;
          _262_v192 = _dafny.Map.Empty.slice().updateUnsafe((_57_v48).fm5(_57_v48.f21, new BigNumber(308), globalState),_261_v191);
          _262_v192 = (_262_v192).update(_57_v48.f21, _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,p0));
          (_260_v190).f21 = _260_v190.f21;
          let _263_v193;
          _263_v193 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_module.D6.create_DC14(_67_v56)).dtor_cf20);
          let _264_v194;
          let _nw48 = new _module.C1();
          _nw48.__ctor(false, (_67_v56).f23, new BigNumber(((_259_v189).update(_57_v48.f21, _module.__default.abs((_67_v56).f23))).cardinality()), _57_v48.f21);
          _264_v194 = _nw48;
          let _265_v195;
          _265_v195 = _dafny.Map.Empty.slice().updateUnsafe(_264_v194,_57_v48.f21);
          let _266_v196;
          _266_v196 = _dafny.Map.Empty.slice().updateUnsafe(_67_v56,(_68_v57)[_module.__default.safeIndex((_67_v56).f23, new BigNumber((_68_v57).length))]);
          let _267_v197;
          let _nw49 = Array((new BigNumber(19)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = (((_263_v193).contains(new BigNumber((_265_v195).length))) ? ((_263_v193).get(new BigNumber((_265_v195).length))) : (_67_v56));
          _nw49[(_dafny.ONE).toNumber()] = _67_v56;
          _nw49[(new BigNumber(2)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(3)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(4)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(5)).toNumber()] = (((_266_v196).contains(_67_v56)) ? ((_266_v196).get(_67_v56)) : (_67_v56));
          _nw49[(new BigNumber(6)).toNumber()] = ((_57_v48.f21) ? (_67_v56) : (_67_v56));
          _nw49[(new BigNumber(7)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(8)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(9)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(10)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(11)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(12)).toNumber()] = (_68_v57)[_module.__default.safeIndex((_264_v194).f19, new BigNumber((_68_v57).length))];
          _nw49[(new BigNumber(13)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(14)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(15)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(16)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(17)).toNumber()] = _67_v56;
          _nw49[(new BigNumber(18)).toNumber()] = _67_v56;
          _267_v197 = _nw49;
          let _index34 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_267_v197).length));
          (_267_v197)[_index34] = _67_v56;
          let _index35 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_267_v197).length));
          let _nw50 = new _module.C0();
          _nw50.__ctor(_67_v56.f22, (_264_v194).f19, !(_23_v15));
          (_267_v197)[_index35] = _nw50;
        } else {
          let _268_v198;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_66_v55, p1, _57_v48.f21);
          _268_v198 = _nw51;
          _268_v198 = _268_v198;
          (globalState).f3 = (_dafny.ZERO).minus((_67_v56).f23);
          let _269_v199;
          _269_v199 = _dafny.Map.Empty.slice().updateUnsafe(p0,_69_v58);
          _269_v199 = (_269_v199).update(_module.__default.safeModuloInt(p0, (_67_v56).f23), _dafny.Seq.update(_dafny.Seq.of((_67_v56).f23), _module.__default.safeIndex((_67_v56).f23, new BigNumber((_dafny.Seq.of((_67_v56).f23)).length)), p1));
          _69_v58 = _dafny.Seq.of((_67_v56).f23, (_67_v56).f23);
          let _270_v200;
          let _init7 = ((_271_v48, _272_v198) => function (_273_i15) {
            return (_dafny.MultiSet.fromElements(_271_v48.f21, _271_v48.f21)).IsSubsetOf(_dafny.MultiSet.fromElements((_272_v198).f18));
          })(_57_v48, _268_v198);
          let _nw52 = Array((new BigNumber(11)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw52.length); _i0_7++) {
            _nw52[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _270_v200 = _nw52;
          let _index36 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_270_v200).length));
          (_270_v200)[_index36] = _57_v48.f21;
          let _274_v201;
          _274_v201 = _dafny.Set.fromElements(_270_v200);
          let _275_v202;
          _275_v202 = _274_v201;
          let _276_v203;
          let _nw53 = new _module.C1();
          _nw53.__ctor(_57_v48.f21, p1, (_67_v56).f23, true);
          _276_v203 = _nw53;
          let _277_v204;
          _277_v204 = _dafny.Map.Empty.slice().updateUnsafe(_276_v203,(_276_v203).f19);
          let _278_v205;
          _278_v205 = _dafny.Seq.of(_276_v203, _276_v203, _276_v203);
          let _279_v206;
          _279_v206 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,(_277_v204).update((_278_v205)[_module.__default.safeIndex((_67_v56).f23, new BigNumber((_278_v205).length))], p1));
          let _280_v207;
          _280_v207 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_279_v206).contains(_57_v48.f21)) ? ((_279_v206).get(_57_v48.f21)) : (_277_v204))).length),_192_v143);
          let _index37 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_270_v200).length));
          let _rhs39 = !((((_57_v48.f21) ? (_dafny.Set.fromElements(_270_v200, _270_v200, _270_v200, _270_v200, _270_v200)) : ((_275_v202)))).IsDisjointFrom((_dafny.Set.fromElements(_270_v200)).Difference(_274_v201)));
          let _rhs40 = p1;
          let _rhs41 = !((new BigNumber((_280_v207).length)).isLessThanOrEqualTo(_module.__default.safeModuloInt(p1, p0)));
          let _lhs31 = _270_v200;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_270_v200).length));
          let _lhs33 = globalState;
          let _lhs34 = globalState;
          _lhs31[_lhs32] = _rhs39;
          _lhs33.f5 = _rhs40;
          _lhs34.f13 = _rhs41;
        }
        let _281_v208;
        _281_v208 = _module.D2.create_DC5(((_57_v48.f21) ? (_192_v143) : (_192_v143)));
        let _rhs42 = _192_v143;
        let _rhs43 = _module.D2.create_DC5(_192_v143);
        let _rhs44 = _23_v15;
        let _lhs35 = globalState;
        _192_v143 = _rhs42;
        _281_v208 = _rhs43;
        _lhs35.f13 = _rhs44;
        let _282_v209;
        let _init8 = ((_283_v16) => function (_284_i16) {
          return (_284_i16).plus(new BigNumber((_283_v16).length));
        })(_24_v16);
        let _nw54 = Array((new BigNumber(10)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw54.length); _i0_8++) {
          _nw54[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _282_v209 = _nw54;
        let _index38 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length));
        (_282_v209)[_index38] = p0;
        let _285_v210;
        let _init9 = ((_286_v146) => function (_287_i17) {
          return _286_v146;
        })(_207_v146);
        let _nw55 = Array((new BigNumber(18)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw55.length); _i0_9++) {
          _nw55[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _285_v210 = _nw55;
        let _index39 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_285_v210).length));
        (_285_v210)[_index39] = _207_v146;
        let _288_v211;
        let _init10 = ((_289_v48) => function (_290_i18) {
          return _289_v48.f21;
        })(_57_v48);
        let _nw56 = Array((new BigNumber(26)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw56.length); _i0_10++) {
          _nw56[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _288_v211 = _nw56;
        let _index40 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_288_v211).length));
        (_288_v211)[_index40] = _dafny.Seq.IsProperPrefixOf(_69_v58, _69_v58);
        let _291_v212;
        _291_v212 = _dafny.Set.fromElements(_57_v48.f21);
        let _292_v213;
        let _nw57 = new _module.C0();
        _nw57.__ctor(_67_v56.f22, new BigNumber((_291_v212).length), !(_23_v15));
        _292_v213 = _nw57;
        let _293_v214;
        _293_v214 = _dafny.Set.fromElements(_292_v213, _292_v213);
        let _294_v215;
        _294_v215 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index41 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length));
        let _index42 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_285_v210).length));
        let _index43 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_288_v211).length));
        let _rhs45 = p1;
        let _rhs46 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_295_i19) {
          return ((false) ? (new _dafny.CodePoint('v'.codePointAt(0))) : (new _dafny.CodePoint('e'.codePointAt(0))));
        })).length);
        let _rhs47 = _dafny.Seq.update(_207_v146, _module.__default.safeIndex(p1, new BigNumber((_207_v146).length)), _294_v215);
        let _rhs48 = _57_v48.f21;
        let _rhs49 = ((_dafny.Seq.contains(_207_v146, _294_v215)) ? (_dafny.Set.fromElements(_292_v213, _292_v213, _292_v213)) : (_dafny.Set.fromElements(_292_v213)));
        let _lhs36 = _282_v209;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length));
        let _lhs38 = globalState;
        let _lhs39 = _285_v210;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_285_v210).length));
        let _lhs41 = _288_v211;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(784), new BigNumber((_288_v211).length));
        _lhs36[_lhs37] = _rhs45;
        _lhs38.f5 = _rhs46;
        _lhs39[_lhs40] = _rhs47;
        _lhs41[_lhs42] = _rhs48;
        _293_v214 = _rhs49;
        let _296_v216;
        _296_v216 = _dafny.Map.Empty.slice().updateUnsafe(_57_v48.f21,_57_v48.f21);
        let _297_v217;
        _297_v217 = _module.D1.create_DC4(!((_292_v213).f18), (((_24_v16)[_module.__default.safeIndex((_67_v56).f23, new BigNumber((_24_v16).length))]) ? ((_57_v48).fm5(_23_v15, (_67_v56).f23, globalState)) : (_57_v48.f21)), (new BigNumber((_296_v216).length)).isLessThan(p0));
        let _index44 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length));
        let _rhs50 = (_282_v209)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length))];
        let _rhs51 = new BigNumber((_dafny.Seq.of(new BigNumber(((_285_v210)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_285_v210).length))]).length))).length);
        let _rhs52 = _297_v217;
        let _lhs43 = globalState;
        let _lhs44 = _282_v209;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_282_v209).length));
        _lhs43.f5 = _rhs50;
        _lhs44[_lhs45] = _rhs51;
        _297_v217 = _rhs52;
      }
      let _298_v218;
      let _nw58 = new _module.C0();
      _nw58.__ctor(_66_v55, p0, _57_v48.f21);
      _298_v218 = _nw58;
      let _299_v219;
      _299_v219 = _dafny.Map.Empty.slice().updateUnsafe(_67_v56,_298_v218);
      r0 = new BigNumber((((_299_v219).update(_67_v56, _298_v218)).Merge(_299_v219)).length);
      let _300_v220;
      let _init11 = ((_301_v56) => function (_302_i20) {
        return (_302_i20).minus((_301_v56).f23);
      })(_67_v56);
      let _nw59 = Array((new BigNumber(23)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw59.length); _i0_11++) {
        _nw59[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _300_v220 = _nw59;
      let _303_v221;
      _303_v221 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(690),_300_v220);
      let _304_v222;
      _304_v222 = _dafny.MultiSet.fromElements(_57_v48.f21);
      r1 = (_303_v221).update(new BigNumber((_304_v222).cardinality()), _300_v220);
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _305_v0;
      let _nw60 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _305_v0 = _nw60;
      let _306_v1;
      let _nw61 = Array((new BigNumber(2)).toNumber()).fill(false);
      _306_v1 = _nw61;
      let _307_v2;
      _307_v2 = _dafny.Seq.UnicodeFromString("mqrmmy");
      let _308_v3;
      _308_v3 = false;
      let _309_globalState;
      let _nw62 = new _module.GlobalState();
      _nw62.__ctor(_305_v0, true, _306_v1, _dafny.ZERO, _dafny.Seq.UnicodeFromString("okaofhw"), new BigNumber(-352), new BigNumber(893), _307_v2, new BigNumber(868), new BigNumber(685), true, new BigNumber(229), false, false, new BigNumber(157), false, true, ((_308_v3) ? (_306_v1) : (_306_v1)));
      _309_globalState = _nw62;
      let _310_v4;
      let _init12 = ((_311_v2) => function (_312_i0) {
        return _module.__default.safeDivisionInt(_312_i0, new BigNumber((_311_v2).length));
      })(_307_v2);
      let _nw63 = Array((new BigNumber(14)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw63.length); _i0_12++) {
        _nw63[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _310_v4 = _nw63;
      let _313_v5;
      _313_v5 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_310_v4);
      let _314_v6;
      _314_v6 = _dafny.Seq.of((((_313_v5).contains(_308_v3)) ? ((_313_v5).get(_308_v3)) : (_310_v4)), _310_v4, _310_v4);
      _314_v6 = _dafny.Seq.Concat(_314_v6, _314_v6);
      let _315_v7;
      _315_v7 = new BigNumber(604);
      (_309_globalState).f9 = ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_308_v3)).length)))).multipliedBy(_module.__default.fm0(new BigNumber((_307_v2).length), new BigNumber((_307_v2).length), _315_v7, _315_v7, _309_globalState));
      let _316_v8;
      _316_v8 = _dafny.MultiSet.fromElements(_310_v4, _310_v4);
      if (!(((_dafny.MultiSet.fromElements(_310_v4)).update(_310_v4, _module.__default.abs(_315_v7))).equals((_316_v8).Union(_316_v8)))) {
        let _317_v9;
        let _init13 = ((_318_v7) => function (_319_i1) {
          return _dafny.Set.fromElements(_318_v7);
        })(_315_v7);
        let _nw64 = Array((new BigNumber(24)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw64.length); _i0_13++) {
          _nw64[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _317_v9 = _nw64;
        let _320_v10;
        _320_v10 = new _dafny.CodePoint('f'.codePointAt(0));
        let _321_v11;
        _321_v11 = _dafny.Seq.of(_308_v3, _308_v3, _308_v3, !(_315_v7).isEqualTo(_315_v7), (_315_v7).isLessThan(new BigNumber((_307_v2).length)));
        let _rhs53 = (((_dafny.ZERO).minus(_315_v7)).multipliedBy((_dafny.ZERO).minus(_315_v7))).minus(_315_v7);
        let _rhs54 = (((_315_v7).isLessThanOrEqualTo(_315_v7)) ? (_317_v9) : (_317_v9));
        let _rhs55 = _320_v10;
        let _rhs56 = _305_v0;
        let _rhs57 = ((_308_v3) ? (_dafny.Seq.of(_308_v3, _308_v3, _module.__default.fm1(_315_v7, _308_v3, _dafny.MultiSet.fromElements(_308_v3), _309_globalState))) : (_321_v11));
        let _lhs46 = _309_globalState;
        _lhs46.f9 = _rhs53;
        _317_v9 = _rhs54;
        _320_v10 = _rhs55;
        _305_v0 = _rhs56;
        _321_v11 = _rhs57;
        let _322_v12;
        _322_v12 = _dafny.MultiSet.fromElements(_308_v3);
        let _323_v13;
        _323_v13 = _dafny.MultiSet.fromElements(_306_v1);
        let _324_v14;
        let _325_v15;
        let _out16;
        let _out17;
        let _outcollector4 = _module.__default.m0((((_322_v12).contains(_308_v3)) ? ((_322_v12).get(_308_v3)) : (new BigNumber((_323_v13).cardinality()))), _315_v7, _309_globalState);
        _out16 = _outcollector4[0];
        _out17 = _outcollector4[1];
        _324_v14 = _out16;
        _325_v15 = _out17;
        let _326_v16;
        _326_v16 = _dafny.Set.fromElements(_308_v3, _308_v3);
        let _327_v17;
        _327_v17 = _dafny.Seq.of(new BigNumber(615));
        let _328_v18;
        _328_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_315_v7),(_321_v11)[_module.__default.safeIndex(_324_v14, new BigNumber((_321_v11).length))]);
        let _329_v19;
        _329_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_326_v16).length), (_327_v17)[_module.__default.safeIndex(_315_v7, new BigNumber((_327_v17).length))], _315_v7, new BigNumber((_328_v18).length), _309_globalState),_308_v3);
        let _330_v20;
        _330_v20 = _dafny.Seq.of(new BigNumber((_329_v19).length));
        let _331_v21;
        _331_v21 = _dafny.Seq.of(_315_v7, _module.__default.safeModuloInt(new BigNumber((_327_v17).length), _324_v14), _315_v7);
        let _332_v22;
        _332_v22 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_308_v3);
        let _333_v23;
        _333_v23 = _dafny.MultiSet.fromElements(_315_v7);
        let _334_v24;
        _334_v24 = _dafny.Seq.of(_333_v23);
        let _335_v25;
        _335_v25 = _dafny.Seq.of(_334_v24, _334_v24);
        let _336_v26;
        _336_v26 = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_315_v7);
        let _rhs58 = _module.__default.fm2(_320_v10, _308_v3, _309_globalState);
        let _rhs59 = ((!(new BigNumber(568)).isEqualTo(new BigNumber((_332_v22).length))) ? (_315_v7) : (new BigNumber((_dafny.Seq.Concat((_335_v25)[_module.__default.safeIndex((((_336_v26).contains(new BigNumber(515))) ? ((_336_v26).get(new BigNumber(515))) : (_324_v14)), new BigNumber((_335_v25).length))], _334_v24)).length)));
        _330_v20 = _rhs58;
        _324_v14 = _rhs59;
        (_309_globalState).f13 = _308_v3;
        (_309_globalState).f3 = _324_v14;
      } else {
        if ((false) === (_308_v3)) {
          (_309_globalState).f2 = _306_v1;
          let _337_v27;
          _337_v27 = new _dafny.CodePoint('r'.codePointAt(0));
          let _338_v28;
          _338_v28 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_337_v27);
          let _339_v29;
          _339_v29 = _dafny.Seq.of(_338_v28);
          (_309_globalState).f5 = new BigNumber((_dafny.Seq.Concat(_339_v29, _339_v29)).length);
          let _340_v30;
          _340_v30 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC0(_308_v3)).dtor_cf0,_315_v7);
          let _341_v31;
          _341_v31 = _dafny.Seq.of(_308_v3, _308_v3);
          let _rhs60 = _340_v30;
          let _rhs61 = _dafny.Seq.Concat(((false) ? (_341_v31) : (_341_v31)), _341_v31);
          let _rhs62 = _308_v3;
          let _lhs47 = _309_globalState;
          _340_v30 = _rhs60;
          _341_v31 = _rhs61;
          _lhs47.f13 = _rhs62;
          (_309_globalState).f3 = new BigNumber(176);
          let _342_v32;
          let _nw65 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Set.Empty);
          _342_v32 = _nw65;
          let _343_v33;
          _343_v33 = _dafny.Set.fromElements(_315_v7, _315_v7, _315_v7, new BigNumber(-714), (_dafny.ZERO).minus((_dafny.ZERO).minus(_315_v7)));
          let _index45 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_342_v32).length));
          (_342_v32)[_index45] = _343_v33;
          let _index46 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_342_v32).length));
          (_342_v32)[_index46] = _343_v33;
        } else {
          let _344_v34;
          _344_v34 = _dafny.Seq.of(_308_v3);
          let _345_v35;
          _345_v35 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_308_v3);
          let _346_v36;
          _346_v36 = _dafny.Seq.of(_dafny.Seq.Concat(_344_v34, _344_v34), _dafny.Seq.Concat((_module.D1.create_DC2(_dafny.Seq.of(_308_v3, _308_v3))).dtor_cf4, _dafny.Seq.of((((_345_v35).contains(_308_v3)) ? ((_345_v35).get(_308_v3)) : (_308_v3)))));
          _344_v34 = (_346_v36)[_module.__default.safeIndex(_315_v7, new BigNumber((_346_v36).length))];
          let _347_v37;
          _347_v37 = _dafny.Map.Empty.slice().updateUnsafe((_315_v7).plus(_315_v7),_308_v3);
          _347_v37 = (_347_v37).update(_315_v7, !_dafny.areEqual(_307_v2, _307_v2));
          (_309_globalState).f9 = (_dafny.ZERO).minus(((_308_v3) ? ((_dafny.ZERO).minus((_315_v7).multipliedBy(_315_v7))) : (_315_v7)));
          let _index47 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_306_v1).length));
          (_306_v1)[_index47] = false;
          let _index48 = _module.__default.safeIndex(new BigNumber(886), new BigNumber((_306_v1).length));
          (_306_v1)[_index48] = ((_308_v3) ? (_308_v3) : (false));
          let _348_v38;
          _348_v38 = _module.D1.create_DC4(_308_v3, (_306_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_306_v1).length))], _308_v3);
          let _349_v39;
          _349_v39 = _dafny.MultiSet.fromElements(_348_v38, _348_v38);
          let _index49 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_310_v4).length));
          (_310_v4)[_index49] = (((_349_v39).contains(_348_v38)) ? ((_349_v39).get(_348_v38)) : (_315_v7));
          let _350_v40;
          let _init14 = ((_351_v1, _352_v7, _353_v2) => function (_354_i2) {
            return _dafny.Seq.update((((_351_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_351_v1).length))]) ? (_dafny.Seq.UnicodeFromString("ghaimy")) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nvj"), _module.__default.safeIndex((_dafny.ZERO).minus(_352_v7), new BigNumber((_dafny.Seq.UnicodeFromString("nvj")).length)), new _dafny.CodePoint('k'.codePointAt(0))))), _module.__default.safeIndex(new BigNumber((_353_v2).length), new BigNumber(((((_351_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_351_v1).length))]) ? (_dafny.Seq.UnicodeFromString("ghaimy")) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nvj"), _module.__default.safeIndex((_dafny.ZERO).minus(_352_v7), new BigNumber((_dafny.Seq.UnicodeFromString("nvj")).length)), new _dafny.CodePoint('k'.codePointAt(0)))))).length)), new _dafny.CodePoint('l'.codePointAt(0)));
          })(_306_v1, _315_v7, _307_v2);
          let _nw66 = Array((new BigNumber(22)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw66.length); _i0_14++) {
            _nw66[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _350_v40 = _nw66;
          let _index50 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_350_v40).length));
          (_350_v40)[_index50] = _307_v2;
          let _355_v41;
          _355_v41 = _dafny.MultiSet.fromElements((_306_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_306_v1).length))]);
          let _356_v42;
          _356_v42 = _dafny.Map.Empty.slice().updateUnsafe(false,_307_v2);
          let _357_v43;
          _357_v43 = _dafny.Seq.of(_dafny.Seq.Concat(_307_v2, _307_v2), (((_356_v42).contains(_308_v3)) ? ((_356_v42).get(_308_v3)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_358_i3) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }))), _307_v2);
          let _index51 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_310_v4).length));
          let _index52 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_350_v40).length));
          let _rhs63 = _module.__default.fm1(_315_v7, _dafny.Seq.IsProperPrefixOf(_307_v2, _307_v2), (_dafny.MultiSet.fromElements(_308_v3, (_306_v1)[_module.__default.safeIndex(new BigNumber(886), new BigNumber((_306_v1).length))])).Difference(_355_v41), _309_globalState);
          let _rhs64 = _module.__default.safeDivisionInt(_315_v7, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_307_v2).length)), _315_v7));
          let _rhs65 = _308_v3;
          let _rhs66 = (_357_v43)[_module.__default.safeIndex(_315_v7, new BigNumber((_357_v43).length))];
          let _lhs48 = _309_globalState;
          let _lhs49 = _310_v4;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_310_v4).length));
          let _lhs51 = _350_v40;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_350_v40).length));
          _lhs48.f13 = _rhs63;
          _lhs49[_lhs50] = _rhs64;
          _308_v3 = _rhs65;
          _lhs51[_lhs52] = _rhs66;
        }
        let _359_v44;
        _359_v44 = _dafny.Seq.of((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_315_v7, _315_v7)));
        _359_v44 = ((!(((_308_v3) ? (!(_308_v3)) : (_308_v3)))) ? (_359_v44) : (_dafny.Seq.Concat(_359_v44, _359_v44)));
        if ((((new BigNumber(810)).isLessThan(_315_v7)) ? (!(_308_v3) || (_308_v3)) : ((_315_v7).isLessThan(_315_v7)))) {
          let _360_v45;
          _360_v45 = _dafny.MultiSet.fromElements(_308_v3);
          let _361_v46;
          _361_v46 = _dafny.Set.fromElements(true, _module.__default.fm1(_315_v7, _308_v3, _360_v45, _309_globalState));
          let _index53 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_310_v4).length));
          (_310_v4)[_index53] = _module.__default.fm0(new BigNumber((_361_v46).length), _315_v7, _315_v7, _315_v7, _309_globalState);
          let _362_v47;
          _362_v47 = _dafny.Set.fromElements(_307_v2, _307_v2, _module.__default.fm3(_309_globalState), _307_v2, _307_v2);
          let _index54 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_310_v4).length));
          (_310_v4)[_index54] = (_dafny.ZERO).minus(_module.__default.fm0(_315_v7, new BigNumber(780), _315_v7, new BigNumber((((_308_v3) ? (_362_v47) : (_362_v47))).length), _309_globalState));
          (_309_globalState).f13 = _308_v3;
          _308_v3 = !(_308_v3);
          let _363_v48;
          _363_v48 = new _dafny.CodePoint('b'.codePointAt(0));
          let _rhs67 = _308_v3;
          let _rhs68 = _315_v7;
          let _rhs69 = _363_v48;
          let _lhs53 = _309_globalState;
          let _lhs54 = _309_globalState;
          _lhs53.f13 = _rhs67;
          _lhs54.f9 = _rhs68;
          _363_v48 = _rhs69;
          let _364_v49;
          let _365_v50;
          let _out18;
          let _out19;
          let _outcollector5 = _module.__default.m0(_315_v7, _315_v7, _309_globalState);
          _out18 = _outcollector5[0];
          _out19 = _outcollector5[1];
          _364_v49 = _out18;
          _365_v50 = _out19;
        } else {
          let _366_v51;
          _366_v51 = new _dafny.CodePoint('a'.codePointAt(0));
          _366_v51 = ((_308_v3) ? ((_307_v2)[_module.__default.safeIndex(_315_v7, new BigNumber((_307_v2).length))]) : (((_308_v3) ? (_366_v51) : (_366_v51))));
          let _367_v52;
          _367_v52 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,new BigNumber(628));
          let _368_v53;
          _368_v53 = _dafny.MultiSet.fromElements(_367_v52, _367_v52, (_367_v52).update(_308_v3, _315_v7));
          (_309_globalState).f13 = (_368_v53).IsDisjointFrom((_368_v53).Union(_368_v53));
          let _369_v54;
          _369_v54 = _dafny.MultiSet.fromElements(_306_v1, _306_v1);
          let _370_v55;
          _370_v55 = _dafny.MultiSet.fromElements(_308_v3, _308_v3);
          (_309_globalState).f5 = (new BigNumber(((_369_v54).Union((_369_v54).update(_306_v1, _module.__default.abs(_315_v7)))).cardinality())).plus((new BigNumber((_367_v52).length)).minus(new BigNumber((_370_v55).cardinality())));
          let _371_v56;
          _371_v56 = _dafny.Seq.of(!(_308_v3));
          let _372_v57;
          _372_v57 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_315_v7, _315_v7, _315_v7, new BigNumber((_371_v56).length), _309_globalState),_315_v7);
          let _373_v59;
          _373_v59 = _dafny.MultiSet.fromElements(_315_v7, _315_v7);
          let _374_v60;
          _374_v60 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(248),_315_v7));
          let _375_v61;
          _375_v61 = _module.D2.create_DC5(_dafny.Map.Empty.slice().updateUnsafe(_315_v7,(((_367_v52).contains(_308_v3)) ? ((_367_v52).get(_308_v3)) : (new BigNumber(-956)))));
          let _376_v64;
          let _nw67 = Array((new BigNumber(27)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _372_v57;
          _nw67[(_dafny.ONE).toNumber()] = _372_v57;
          _nw67[(new BigNumber(2)).toNumber()] = (_372_v57).Merge(_372_v57);
          _nw67[(new BigNumber(3)).toNumber()] = function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_373_v59).Elements) {
              let _377_v58 = _compr_6;
              if ((_373_v59).contains(_377_v58)) {
                _coll6.push([_module.__default.safeDivisionInt(_377_v58, _315_v7),(_dafny.ZERO).minus(_315_v7)]);
              }
            }
            return _coll6;
          }();
          _nw67[(new BigNumber(4)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(5)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(6)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(7)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(8)).toNumber()] = (_372_v57).Merge(_372_v57);
          _nw67[(new BigNumber(9)).toNumber()] = (_374_v60)[_module.__default.safeIndex(new BigNumber(498), new BigNumber((_374_v60).length))];
          _nw67[(new BigNumber(10)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(11)).toNumber()] = (_372_v57).Merge(_372_v57);
          _nw67[(new BigNumber(12)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_315_v7);
          _nw67[(new BigNumber(14)).toNumber()] = ((_308_v3) ? (_372_v57) : ((_374_v60)[_module.__default.safeIndex(_315_v7, new BigNumber((_374_v60).length))]));
          _nw67[(new BigNumber(15)).toNumber()] = (_375_v61).dtor_cf9;
          _nw67[(new BigNumber(16)).toNumber()] = (_372_v57).Merge(_372_v57);
          _nw67[(new BigNumber(17)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_307_v2).length),_module.__default.fm0(_315_v7, new BigNumber((_372_v57).length), _315_v7, new BigNumber((_307_v2).length), _309_globalState));
          _nw67[(new BigNumber(19)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(20)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(21)).toNumber()] = _module.__default.fm4(_dafny.Seq.UnicodeFromString("n"), true, _308_v3, _308_v3, _309_globalState);
          _nw67[(new BigNumber(22)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(23)).toNumber()] = function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-970), new BigNumber(932))) {
              let _378_v62 = _compr_7;
              if (((new BigNumber(-970)).isLessThanOrEqualTo(_378_v62)) && ((_378_v62).isLessThan(new BigNumber(932)))) {
                _coll7.push([(_378_v62).minus(_315_v7),new BigNumber((_dafny.Seq.of(new BigNumber(723))).length)]);
              }
            }
            return _coll7;
          }();
          _nw67[(new BigNumber(24)).toNumber()] = function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-721), new BigNumber(577))) {
              let _379_v63 = _compr_8;
              if (((new BigNumber(-721)).isLessThanOrEqualTo(_379_v63)) && ((_379_v63).isLessThan(new BigNumber(577)))) {
                _coll8.push([(_379_v63).multipliedBy(new BigNumber((_dafny.Set.fromElements(_308_v3)).length)),_315_v7]);
              }
            }
            return _coll8;
          }();
          _nw67[(new BigNumber(25)).toNumber()] = _372_v57;
          _nw67[(new BigNumber(26)).toNumber()] = ((false) ? (_module.__default.fm4(_307_v2, _308_v3, !(false), _308_v3, _309_globalState)) : ((_dafny.Map.Empty.slice().updateUnsafe(_315_v7,new BigNumber((_307_v2).length))).update(_315_v7, new BigNumber((_367_v52).length))));
          _376_v64 = _nw67;
          _376_v64 = _376_v64;
          let _380_v65;
          let _nw68 = new _module.C1();
          _nw68.__ctor(_308_v3, (_315_v7).multipliedBy(_315_v7), new BigNumber(-522), _308_v3);
          _380_v65 = _nw68;
        }
        let _index55 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_310_v4).length));
        (_310_v4)[_index55] = _315_v7;
        let _index56 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_310_v4).length));
        (_310_v4)[_index56] = _module.__default.safeDivisionInt(_315_v7, new BigNumber(425));
        let _381_v66;
        _381_v66 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_315_v7);
        let _382_v67;
        _382_v67 = _dafny.Seq.of(_308_v3, _308_v3);
        let _383_v68;
        let _384_v69;
        let _out20;
        let _out21;
        let _outcollector6 = _module.__default.m0(_module.__default.fm0((_310_v4)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_310_v4).length))], (_dafny.ZERO).minus((_310_v4)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_310_v4).length))]), _module.__default.fm0(new BigNumber(((_381_v66)).length), _315_v7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_385_v4) => function (_386_i4) {
          return (_385_v4)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_385_v4).length))];
        })(_310_v4))).length), new BigNumber((_382_v67).length), _309_globalState), new BigNumber((_307_v2).length), _309_globalState), new BigNumber(659), _309_globalState);
        _out20 = _outcollector6[0];
        _out21 = _outcollector6[1];
        _383_v68 = _out20;
        _384_v69 = _out21;
      }
      let _387_v70;
      _387_v70 = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_315_v7);
      (_309_globalState).f3 = ((new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_387_v70).Keys.Elements) {
          let _388_v71 = _compr_9;
          if ((_387_v70).contains(_388_v71)) {
            _coll9.add(_module.__default.safeModuloInt(_388_v71, new BigNumber(240)));
          }
        }
        return _coll9;
      }()).length)).multipliedBy(_315_v7)).plus(_315_v7);
      (_309_globalState).f9 = new BigNumber(-648);
      let _389_v72;
      _389_v72 = _module.D0.create_DC0(true);
      let _source3 = _389_v72;
      if (_source3.is_DC1) {
        let _390___mcc_h0 = (_source3).cf1;
        let _391___mcc_h1 = (_source3).cf2;
        let _392___mcc_h2 = (_source3).cf3;
        let _393_cf3 = _392___mcc_h2;
        let _394_cf2 = _391___mcc_h1;
        let _395_cf1 = _390___mcc_h0;
        let _396_v73;
        _396_v73 = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_308_v3);
        let _index57 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_310_v4).length));
        (_310_v4)[_index57] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements((((_396_v73).contains(new BigNumber(415))) ? ((_396_v73).get(new BigNumber(415))) : (_393_cf3)))).length), _394_cf2);
        let _397_v74;
        _397_v74 = _dafny.MultiSet.fromElements(_308_v3, _308_v3, _308_v3, _393_cf3);
        let _398_v75;
        _398_v75 = _dafny.Seq.of(_393_cf3, _308_v3);
        let _399_v76;
        _399_v76 = _dafny.Seq.of(_394_cf2);
        let _400_v77;
        _400_v77 = _dafny.Seq.of(_399_v76);
        let _index58 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_310_v4).length));
        (_310_v4)[_index58] = (((_dafny.MultiSet.FromArray(_398_v75)).IsSubsetOf(_397_v74)) ? (_module.__default.safeModuloInt(_395_cf1, new BigNumber(750))) : ((new BigNumber(987)).minus((_dafny.ZERO).minus(new BigNumber(((_400_v77)[_module.__default.safeIndex(_395_cf1, new BigNumber((_400_v77).length))]).length)))));
        let _401_v78;
        _401_v78 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_308_v3);
        let _402_v79;
        let _403_v80;
        let _out22;
        let _out23;
        let _outcollector7 = _module.__default.m0((_dafny.ZERO).minus((new BigNumber((_397_v74).cardinality())).minus(new BigNumber((_401_v78).length))), new BigNumber((_dafny.Seq.UnicodeFromString("lmuinekm")).length), _309_globalState);
        _out22 = _outcollector7[0];
        _out23 = _outcollector7[1];
        _402_v79 = _out22;
        _403_v80 = _out23;
        let _404_v81;
        _404_v81 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_315_v7);
        let _405_v82;
        _405_v82 = _dafny.Map.Empty.slice().updateUnsafe(_404_v81,_308_v3);
        _398_v75 = _dafny.Seq.update(_398_v75, _module.__default.safeIndex((_310_v4)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_310_v4).length))], new BigNumber((_398_v75).length)), (_402_v79).isEqualTo(new BigNumber((_405_v82).length)));
        _401_v78 = (_401_v78).update(false, false);
      } else {
        let _406___mcc_h3 = (_source3).cf0;
        let _407_cf0 = _406___mcc_h3;
        (_309_globalState).f4 = _307_v2;
        (_309_globalState).f4 = _dafny.Seq.Concat(_307_v2, _307_v2);
        let _408_v83;
        let _nw69 = new _module.C1();
        _nw69.__ctor(!(_407_cf0), _315_v7, _315_v7, _407_cf0);
        _408_v83 = _nw69;
        let _409_v84;
        _409_v84 = _dafny.Map.Empty.slice().updateUnsafe(_408_v83,_315_v7);
        let _410_v85;
        _410_v85 = _dafny.Seq.of((_315_v7).minus(_315_v7), _315_v7, new BigNumber(516), new BigNumber((_409_v84).length), _408_v83.f20);
        _410_v85 = _410_v85;
        let _411_v86;
        _411_v86 = _module.D1.create_DC3((_408_v83).f19);
        let _412_v87;
        _412_v87 = _dafny.Map.Empty.slice().updateUnsafe(false,_411_v86);
        let _413_v88;
        _413_v88 = _dafny.Seq.of(_412_v87, _412_v87, _412_v87);
        (_309_globalState).f13 = (_408_v83).fm5(false, new BigNumber((_413_v88).length), _309_globalState);
      }
      let _414_v89;
      let _nw70 = new _module.C1();
      _nw70.__ctor(false, new BigNumber((_module.__default.fm8(_309_globalState)).length), _315_v7, !(_308_v3));
      _414_v89 = _nw70;
      let _415_v90;
      _415_v90 = _dafny.Map.Empty.slice().updateUnsafe(_414_v89,_307_v2);
      let _416_v91;
      _416_v91 = _module.D4.create_DC9(_414_v89);
      _415_v90 = (_415_v90).update((_416_v91).dtor_cf13, _307_v2);
      let _417_v92;
      _417_v92 = new _dafny.CodePoint('t'.codePointAt(0));
      let _418_v93;
      let _nw71 = new _module.C1();
      _nw71.__ctor(_308_v3, _module.__default.safeModuloInt(_315_v7, _315_v7), _315_v7, _414_v89.f21);
      _418_v93 = _nw71;
      let _419_v94;
      _419_v94 = _dafny.Seq.of(!((_418_v93).f18), _414_v89.f21);
      let _420_v95;
      _420_v95 = _dafny.Map.Empty.slice().updateUnsafe(_419_v94,_308_v3);
      let _421_v96;
      _421_v96 = _dafny.Map.Empty.slice().updateUnsafe(_305_v0,(_418_v93).f18);
      let _422_v97;
      _422_v97 = _dafny.MultiSet.fromElements((((_420_v95).contains(_419_v94)) ? ((_420_v95).get(_419_v94)) : (_414_v89.f21)), _308_v3, (((_421_v96).contains(_305_v0)) ? ((_421_v96).get(_305_v0)) : (false)));
      let _423_v98;
      _423_v98 = _dafny.Map.Empty.slice().updateUnsafe(_308_v3,_308_v3);
      let _424_v99;
      _424_v99 = _dafny.Seq.of(_423_v98, _dafny.Map.Empty.slice().updateUnsafe((_418_v93).f18,_308_v3));
      let _425_v100;
      _425_v100 = _dafny.Seq.of(new BigNumber(-186));
      let _426_v101;
      _426_v101 = _dafny.Map.Empty.slice().updateUnsafe((_418_v93).f18,(_418_v93).f19);
      let _427_v102;
      _427_v102 = _module.D0.create_DC1(_315_v7, (_425_v100)[_module.__default.safeIndex(new BigNumber((_426_v101).length), new BigNumber((_425_v100).length))], _308_v3);
      let _428_v103;
      _428_v103 = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_418_v93);
      let _rhs70 = new _dafny.CodePoint('s'.codePointAt(0));
      let _rhs71 = ((false) ? (_307_v2) : (_dafny.Seq.UnicodeFromString("nd")));
      let _rhs72 = (_418_v93).fm5(_dafny.areEqual(_424_v99, _dafny.Seq.update(_424_v99, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_417_v92,new BigNumber((_387_v70).length))).length), new BigNumber((_424_v99).length)), _module.__default.fm9(_414_v89.f21, (_418_v93).f18, _309_globalState))), (_427_v102).dtor_cf2, _309_globalState);
      let _rhs73 = (((_428_v103).contains(_module.__default.safeDivisionInt(_418_v93.f20, _315_v7))) ? ((_428_v103).get(_module.__default.safeDivisionInt(_418_v93.f20, _315_v7))) : (_418_v93));
      let _rhs74 = _422_v97;
      _417_v92 = _rhs70;
      _307_v2 = _rhs71;
      _308_v3 = _rhs72;
      _418_v93 = _rhs73;
      _422_v97 = _rhs74;
      let _429_v104;
      _429_v104 = _dafny.MultiSet.fromElements(_414_v89);
      (_414_v89).f21 = (_308_v3) && (_module.__default.fm1((((_429_v104).contains(_414_v89)) ? ((_429_v104).get(_414_v89)) : ((_dafny.ZERO).minus(_418_v93.f20))), _308_v3, _422_v97, _309_globalState));
      let _430_v105;
      _430_v105 = _module.D2.create_DC5(_387_v70);
      _430_v105 = _module.__default.fm10(_315_v7, _309_globalState);
      if (!_dafny.areEqual(_module.__default.fm2(_417_v92, _308_v3, _309_globalState), _dafny.Seq.Concat(_425_v100, _425_v100))) {
        let _431_v106;
        let _nw72 = new _module.C1();
        _nw72.__ctor((_418_v93).f18, (_418_v93).f19, _315_v7, _414_v89.f21);
        _431_v106 = _nw72;
        (_309_globalState).f13 = (_418_v93).f18;
        if ((_418_v93).fm5(_308_v3, _315_v7, _309_globalState)) {
          (_309_globalState).f3 = (_dafny.ZERO).minus((_418_v93).f19);
          (_414_v89).f21 = false;
          (_414_v89).f21 = _431_v106.f21;
          let _432_v107;
          _432_v107 = _dafny.MultiSet.fromElements(_418_v93.f20);
          let _433_v108;
          _433_v108 = _module.D2.create_DC6(_432_v107);
          let _434_v109;
          _434_v109 = _dafny.Set.fromElements((_418_v93).f19, (_418_v93).f19, _315_v7);
          let _435_v110;
          let _nw73 = new _module.C0();
          _nw73.__ctor(_module.D2.create_DC6(_dafny.MultiSet.fromElements(new BigNumber((_434_v109).length), (_418_v93).f19, _418_v93.f20)), (_418_v93).f19, (_418_v93).f18);
          _435_v110 = _nw73;
          let _436_v111;
          _436_v111 = _dafny.Seq.of(_435_v110, _435_v110, _435_v110, _435_v110, _435_v110);
          let _437_v112;
          let _nw74 = new _module.C0();
          _nw74.__ctor(_433_v108, _418_v93.f20, (_414_v89).fm5(_414_v89.f21, new BigNumber((_436_v111).length), _309_globalState));
          _437_v112 = _nw74;
          let _438_v113;
          _438_v113 = _dafny.Seq.of(_437_v112);
          let _439_v114;
          let _nw75 = new _module.C1();
          _nw75.__ctor(_414_v89.f21, new BigNumber((_dafny.Seq.Concat(_438_v113, _438_v113)).length), (_315_v7).multipliedBy((_418_v93).f19), ((_module.__default.fm11(_309_globalState)).dtor_cf5).isLessThan((_437_v112).f23));
          _439_v114 = _nw75;
          let _index59 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_310_v4).length));
          (_310_v4)[_index59] = _418_v93.f20;
          let _index60 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_310_v4).length));
          let _rhs75 = (_418_v93.f20).multipliedBy(_418_v93.f20);
          let _rhs76 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber(-158), new BigNumber(654)), _418_v93.f20);
          let _rhs77 = (_dafny.Set.fromElements((_418_v93).f19, new BigNumber((_422_v97).cardinality()))).IsDisjointFrom((_434_v109).Intersect(_434_v109));
          let _rhs78 = _305_v0;
          let _rhs79 = false;
          let _lhs55 = _310_v4;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_310_v4).length));
          let _lhs57 = _309_globalState;
          let _lhs58 = _309_globalState;
          let _lhs59 = _309_globalState;
          _lhs55[_lhs56] = _rhs75;
          _lhs57.f9 = _rhs76;
          _lhs58.f13 = _rhs77;
          _305_v0 = _rhs78;
          _lhs59.f13 = _rhs79;
        } else {
          let _440_v115;
          let _init15 = ((_441_v101) => function (_442_i5) {
            return _441_v101;
          })(_426_v101);
          let _nw76 = Array((new BigNumber(21)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw76.length); _i0_15++) {
            _nw76[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _440_v115 = _nw76;
          let _443_v116;
          _443_v116 = _dafny.Map.Empty.slice().updateUnsafe(_419_v94,_440_v115);
          _443_v116 = (_443_v116).update(_dafny.Seq.of((_418_v93).f18, (_418_v93).f18, (_418_v93).f18, _414_v89.f21, !(_431_v106.f21)), _440_v115);
          let _444_v117;
          _444_v117 = _module.D1.create_DC2(_dafny.Seq.of((_418_v93).f18, _414_v89.f21, _431_v106.f21));
          let _445_v118;
          _445_v118 = _module.D2.create_DC6(_dafny.MultiSet.fromElements(new BigNumber(115), (_418_v93).f19));
          let _446_v119;
          _446_v119 = _dafny.MultiSet.fromElements((_418_v93).f19);
          let _pat_let_tv11 = _446_v119;
          let _447_v120;
          let _nw77 = new _module.C0();
          _nw77.__ctor(function (_pat_let6_0) {
            return function (_448_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_449_dt__update_hcf10_h0) {
                  return _module.D2.create_DC6(_449_dt__update_hcf10_h0);
                }(_pat_let7_0);
              }(_pat_let_tv11);
            }(_pat_let6_0);
          }(_445_v118), (_418_v93).f19, _431_v106.f21);
          _447_v120 = _nw77;
          let _450_v121;
          _450_v121 = _dafny.Map.Empty.slice().updateUnsafe(_447_v120,_418_v93.f20);
          let _451_v122;
          _451_v122 = _dafny.Map.Empty.slice().updateUnsafe(_444_v117,_450_v121);
          _451_v122 = _451_v122;
          _447_v120 = _447_v120;
          let _452_v123;
          _452_v123 = _dafny.Map.Empty.slice().updateUnsafe(_431_v106,_418_v93.f20);
          let _453_v124;
          let _454_v125;
          let _out24;
          let _out25;
          let _outcollector8 = (_418_v93).m2((_446_v119).Difference(_dafny.MultiSet.fromElements(_315_v7, _418_v93.f20)), (((_452_v123).contains(_414_v89)) ? ((_452_v123).get(_414_v89)) : ((_418_v93).f19)), _306_v1, _307_v2, _309_globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _453_v124 = _out24;
          _454_v125 = _out25;
          let _455_v126;
          _455_v126 = _dafny.Map.Empty.slice().updateUnsafe(_414_v89.f21,_306_v1);
          _455_v126 = (_455_v126).update(_414_v89.f21, _306_v1);
        }
        let _456_v127;
        _456_v127 = _dafny.Set.fromElements(_306_v1, _306_v1);
        if ((((_418_v93).f19).multipliedBy(new BigNumber((_456_v127).length))).isEqualTo(_module.__default.safeDivisionInt((_418_v93).f19, (_418_v93).f19))) {
          let _457_v128;
          let _nw78 = Array((new BigNumber(18)).toNumber());
          _457_v128 = _nw78;
          let _458_v129;
          _458_v129 = _dafny.Map.Empty.slice().updateUnsafe(_457_v128,_310_v4);
          let _459_v130;
          _459_v130 = _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)));
          let _460_v131;
          _460_v131 = _dafny.MultiSet.fromElements(_module.__default.fm0(_418_v93.f20, _418_v93.f20, (_418_v93).f19, new BigNumber((_459_v130).length), _309_globalState));
          let _461_v132;
          let _nw79 = new _module.C1();
          _nw79.__ctor(_414_v89.f21, (_418_v93).f19, _module.__default.fm0(_418_v93.f20, (_418_v93).f19, new BigNumber((_458_v129).length), _315_v7, _309_globalState), (_module.D0.create_DC1(new BigNumber((_460_v131).cardinality()), (_418_v93).f19, _431_v106.f21)).dtor_cf3);
          _461_v132 = _nw79;
          (_309_globalState).f9 = _315_v7;
          let _index61 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_306_v1).length));
          (_306_v1)[_index61] = !(_414_v89.f21);
          let _index62 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_306_v1).length));
          (_306_v1)[_index62] = ((new BigNumber((_dafny.Seq.UnicodeFromString("jfbbox")).length)).plus(_418_v93.f20)).isLessThan((_418_v93).f19);
          _387_v70 = (_387_v70).update(_418_v93.f20, (_dafny.ZERO).minus(new BigNumber((_423_v98).length)));
          _307_v2 = (((_415_v90).contains(_431_v106)) ? ((_415_v90).get(_431_v106)) : (_307_v2));
        } else {
          (_309_globalState).f5 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_418_v93.f20, new BigNumber(((_422_v97).Difference(_422_v97)).cardinality())));
          (_414_v89).f21 = !(new BigNumber((_dafny.MultiSet.fromElements((_418_v93).f19, _418_v93.f20)).cardinality())).isEqualTo((_dafny.ZERO).minus(_315_v7));
          _308_v3 = (_387_v70).contains(_418_v93.f20);
          let _462_v133;
          let _463_v134;
          let _464_v135;
          let _465_v136;
          let _out26;
          let _out27;
          let _out28;
          let _out29;
          let _outcollector9 = (_414_v89).m1(_431_v106.f21, (_431_v106.f21) || (false), _309_globalState);
          _out26 = _outcollector9[0];
          _out27 = _outcollector9[1];
          _out28 = _outcollector9[2];
          _out29 = _outcollector9[3];
          _462_v133 = _out26;
          _463_v134 = _out27;
          _464_v135 = _out28;
          _465_v136 = _out29;
          let _rhs80 = (_418_v93).f19;
          let _rhs81 = (_dafny.ZERO).minus(_418_v93.f20);
          let _lhs60 = _309_globalState;
          _464_v135 = _rhs80;
          _lhs60.f3 = _rhs81;
        }
        (_418_v93).f20 = new BigNumber(-525);
      } else {
        let _466_v137;
        _466_v137 = _dafny.Seq.of((((_418_v93).f18) ? (_419_v94) : (_dafny.Seq.of((_418_v93).f18))), _dafny.Seq.of(_414_v89.f21));
        let _467_v138;
        let _nw80 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _467_v138 = _nw80;
        let _index63 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length));
        (_467_v138)[_index63] = _dafny.Seq.Concat(_419_v94, _419_v94);
        let _index64 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length));
        let _rhs82 = (((_418_v93).f19).plus((_418_v93).f19)).plus(_315_v7);
        let _rhs83 = _414_v89.f21;
        let _rhs84 = _466_v137;
        let _rhs85 = _419_v94;
        let _rhs86 = ((!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), ((_468_v92) => function (_469_i6) {
          return _468_v92;
        })(_417_v92)), _307_v2)) ? (((_418_v93).f19).isLessThanOrEqualTo(_418_v93.f20)) : (((_418_v93).f18) || ((_418_v93).f18)));
        let _lhs61 = _309_globalState;
        let _lhs62 = _414_v89;
        let _lhs63 = _467_v138;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length));
        let _lhs65 = _414_v89;
        _lhs61.f9 = _rhs82;
        _lhs62.f21 = _rhs83;
        _466_v137 = _rhs84;
        _lhs63[_lhs64] = _rhs85;
        _lhs65.f21 = _rhs86;
        let _470_v139;
        _470_v139 = _module.D1.create_DC3((_418_v93).f19);
        let _pat_let_tv12 = _418_v93;
        _470_v139 = function (_pat_let8_0) {
          return function (_471_dt__update__tmp_h1) {
            return function (_pat_let9_0) {
              return function (_472_dt__update_hcf5_h0) {
                return _module.D1.create_DC3(_472_dt__update_hcf5_h0);
              }(_pat_let9_0);
            }(_pat_let_tv12.f20);
          }(_pat_let8_0);
        }(_470_v139);
        let _index65 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length));
        let _rhs87 = ((_414_v89.f21) ? (new BigNumber((_425_v100).length)) : ((_418_v93).f19));
        let _rhs88 = _dafny.Seq.IsProperPrefixOf(_307_v2, _307_v2);
        let _rhs89 = _dafny.Seq.Concat(_419_v94, _dafny.Seq.Concat((_467_v138)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length))], _419_v94));
        let _rhs90 = _307_v2;
        let _lhs66 = _309_globalState;
        let _lhs67 = _414_v89;
        let _lhs68 = _467_v138;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_467_v138).length));
        _lhs66.f3 = _rhs87;
        _lhs67.f21 = _rhs88;
        _lhs68[_lhs69] = _rhs89;
        _307_v2 = _rhs90;
        if (_414_v89.f21) {
          let _index66 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_306_v1).length));
          (_306_v1)[_index66] = !((_418_v93).f18);
          let _index67 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_306_v1).length));
          (_306_v1)[_index67] = _308_v3;
          let _473_v140;
          let _nw81 = Array((new BigNumber(4)).toNumber()).fill(false);
          _473_v140 = _nw81;
          (_309_globalState).f2 = _473_v140;
          let _474_v141;
          let _init16 = ((_475_v102) => function (_476_i7) {
            return _475_v102;
          })(_427_v102);
          let _nw82 = Array((new BigNumber(25)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw82.length); _i0_16++) {
            _nw82[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _474_v141 = _nw82;
          let _index68 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_474_v141).length));
          (_474_v141)[_index68] = _427_v102;
          let _index69 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_474_v141).length));
          (_474_v141)[_index69] = _module.D0.create_DC1((_418_v93).f19, _module.__default.fm0((_418_v93).f19, (_418_v93).f19, _418_v93.f20, (_418_v93).f19, _309_globalState), (new BigNumber((_387_v70).length)).isLessThanOrEqualTo(_418_v93.f20));
          let _477_v142;
          let _nw83 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
          _477_v142 = _nw83;
          _477_v142 = _477_v142;
          let _index70 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_310_v4).length));
          (_310_v4)[_index70] = (_418_v93).f19;
          let _index71 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_310_v4).length));
          (_310_v4)[_index71] = new BigNumber((_dafny.Seq.Concat(_424_v99, _424_v99)).length);
        } else {
          (_309_globalState).f3 = (_418_v93).f19;
          let _478_v143;
          let _479_v144;
          let _out30;
          let _out31;
          let _outcollector10 = _module.__default.m0(_418_v93.f20, new BigNumber((_module.__default.fm2(_417_v92, _414_v89.f21, _309_globalState)).length), _309_globalState);
          _out30 = _outcollector10[0];
          _out31 = _outcollector10[1];
          _478_v143 = _out30;
          _479_v144 = _out31;
          (_309_globalState).f9 = _478_v143;
          let _rhs91 = _414_v89.f21;
          let _rhs92 = _310_v4;
          let _rhs93 = _478_v143;
          let _lhs70 = _309_globalState;
          _308_v3 = _rhs91;
          _310_v4 = _rhs92;
          _lhs70.f3 = _rhs93;
          let _index72 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_306_v1).length));
          (_306_v1)[_index72] = (_418_v93).f18;
          let _index73 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_306_v1).length));
          (_306_v1)[_index73] = _414_v89.f21;
        }
        let _480_v145;
        _480_v145 = _dafny.Map.Empty.slice().updateUnsafe(_419_v94,(_dafny.ZERO).minus((_418_v93).f19));
        let _481_v146;
        _481_v146 = _dafny.MultiSet.fromElements(new BigNumber((_425_v100).length), new BigNumber((_480_v145).length));
        let _482_v147;
        let _483_v148;
        let _out32;
        let _out33;
        let _outcollector11 = (_414_v89).m2(_481_v146, (_418_v93).f19, _306_v1, _module.__default.fm3(_309_globalState), _309_globalState);
        _out32 = _outcollector11[0];
        _out33 = _outcollector11[1];
        _482_v147 = _out32;
        _483_v148 = _out33;
      }
      let _index74 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
      (_310_v4)[_index74] = (_418_v93).f19;
      let _484_v149;
      let _init17 = ((_485_v101) => function (_486_i8) {
        return _485_v101;
      })(_426_v101);
      let _nw84 = Array((new BigNumber(28)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw84.length); _i0_17++) {
        _nw84[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _484_v149 = _nw84;
      let _487_v150;
      _487_v150 = _426_v101;
      let _index75 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_484_v149).length));
      (_484_v149)[_index75] = ((!(_414_v89.f21)) ? (_487_v150) : (_487_v150));
      let _index76 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
      let _index77 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_484_v149).length));
      let _rhs94 = _418_v93.f20;
      let _rhs95 = _487_v150;
      let _rhs96 = (_module.__default.fm0(_418_v93.f20, (_dafny.ZERO).minus(_418_v93.f20), _module.__default.fm0(new BigNumber(-511), _315_v7, _418_v93.f20, _315_v7, _309_globalState), _418_v93.f20, _309_globalState)).isLessThan(((_308_v3) ? (new BigNumber((_419_v94).length)) : (_418_v93.f20)));
      let _lhs71 = _310_v4;
      let _lhs72 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
      let _lhs73 = _484_v149;
      let _lhs74 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_484_v149).length));
      let _lhs75 = _414_v89;
      _lhs71[_lhs72] = _rhs94;
      _lhs73[_lhs74] = _rhs95;
      _lhs75.f21 = _rhs96;
      let _pat_let_tv13 = _417_v92;
      let _pat_let_tv14 = _417_v92;
      let _pat_let_tv15 = _417_v92;
      _417_v92 = function (_source4) {
        if (_source4.is_DC6) {
          let _488___mcc_h4 = (_source4).cf10;
          let _489_cf10 = _488___mcc_h4;
          return _pat_let_tv13;
        } else if (_source4.is_DC5) {
          let _490___mcc_h5 = (_source4).cf9;
          let _491_cf9 = _490___mcc_h5;
          return _pat_let_tv14;
        } else {
          let _492___mcc_h6 = (_source4).cf11;
          let _493_cf11 = _492___mcc_h6;
          return _pat_let_tv15;
        }
      }(((true) ? (_430_v105) : (_430_v105)));
      (_309_globalState).f3 = new BigNumber(440);
      let _494_v151;
      let _495_v152;
      let _out34;
      let _out35;
      let _outcollector12 = (_414_v89).m2(_module.__default.fm12(_309_globalState), (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _306_v1, ((_414_v89.f21) ? (_dafny.Seq.UnicodeFromString("xmgvn")) : (_307_v2)), _309_globalState);
      _out34 = _outcollector12[0];
      _out35 = _outcollector12[1];
      _494_v151 = _out34;
      _495_v152 = _out35;
      if (((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))]).isLessThan(_418_v93.f20)) {
        let _496_v153;
        _496_v153 = _dafny.MultiSet.fromElements(_418_v93.f20);
        let _497_v154;
        _497_v154 = _dafny.MultiSet.fromElements((_418_v93).f19, _module.__default.fm0(new BigNumber((_496_v153).cardinality()), (_418_v93).f19, (_dafny.ZERO).minus(new BigNumber((_307_v2).length)), (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _309_globalState));
        (_418_v93).f20 = _module.__default.fm0(_315_v7, new BigNumber(790), (((_497_v154).contains(new BigNumber((_425_v100).length))) ? ((_497_v154).get(new BigNumber((_425_v100).length))) : (_module.__default.fm0(_315_v7, _315_v7, new BigNumber(250), new BigNumber((_425_v100).length), _309_globalState))), (_418_v93).f19, _309_globalState);
        let _498_v155;
        _498_v155 = _dafny.Map.Empty.slice().updateUnsafe(_414_v89,_414_v89.f21);
        let _499_v156;
        _499_v156 = _module.D1.create_DC3(new BigNumber((_495_v152).length));
        let _source5 = (((_498_v155).equals(_498_v155)) ? (_499_v156) : (_499_v156));
        if (_source5.is_DC3) {
          let _500___mcc_h7 = (_source5).cf5;
          let _501_cf5 = _500___mcc_h7;
          let _502_v157;
          _502_v157 = _module.D5.create_DC12(_418_v93);
          let _503_v158;
          _503_v158 = _dafny.Map.Empty.slice().updateUnsafe(_414_v89.f21,_418_v93);
          let _504_v159;
          let _nw85 = Array((new BigNumber(18)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _418_v93;
          _nw85[(_dafny.ONE).toNumber()] = _418_v93;
          _nw85[(new BigNumber(2)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(3)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(4)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(5)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(6)).toNumber()] = (_502_v157).dtor_cf18;
          _nw85[(new BigNumber(7)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(8)).toNumber()] = (((_503_v158).contains(_414_v89.f21)) ? ((_503_v158).get(_414_v89.f21)) : (_418_v93));
          _nw85[(new BigNumber(9)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(10)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(11)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(12)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(13)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(14)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(15)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(16)).toNumber()] = _418_v93;
          _nw85[(new BigNumber(17)).toNumber()] = ((_414_v89.f21) ? (_418_v93) : (_418_v93));
          _504_v159 = _nw85;
          let _rhs97 = _504_v159;
          let _rhs98 = _307_v2;
          let _rhs99 = _414_v89;
          let _lhs76 = _309_globalState;
          _504_v159 = _rhs97;
          _lhs76.f4 = _rhs98;
          _414_v89 = _rhs99;
          (_309_globalState).f3 = _501_cf5;
          (_414_v89).f21 = _308_v3;
          (_309_globalState).f4 = ((_414_v89.f21) ? (_307_v2) : (_307_v2));
        } else if (_source5.is_DC4) {
          let _505___mcc_h8 = (_source5).cf6;
          let _506___mcc_h9 = (_source5).cf7;
          let _507___mcc_h10 = (_source5).cf8;
          let _508_cf8 = _507___mcc_h10;
          let _509_cf7 = _506___mcc_h9;
          let _510_cf6 = _505___mcc_h8;
          _426_v101 = _dafny.Map.Empty.slice().updateUnsafe((_418_v93).f18,(_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))]);
          _509_cf7 = (_418_v93).f18;
          _510_cf6 = _509_cf7;
          let _511_v160;
          _511_v160 = _dafny.Seq.of(_307_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(410)), ((_512_v92) => function (_513_i9) {
            return _512_v92;
          })(_417_v92)), _307_v2);
          _419_v94 = _dafny.Seq.update(_dafny.Seq.of((_308_v3) === (_414_v89.f21), _308_v3), _module.__default.safeIndex(new BigNumber((_511_v160).length), new BigNumber((_dafny.Seq.of((_308_v3) === (_414_v89.f21), _308_v3)).length)), (new BigNumber((_module.__default.fm13((_418_v93).f18, (((_423_v98).contains(_308_v3)) ? ((_423_v98).get(_308_v3)) : (_508_cf8)), _418_v93.f20, new BigNumber((_dafny.Seq.of(new BigNumber(389))).length), _309_globalState)).length)).isLessThan(_418_v93.f20));
        } else {
          let _514___mcc_h11 = (_source5).cf4;
          let _515_cf4 = _514___mcc_h11;
          (_414_v89).f21 = !(_dafny.Seq.IsPrefixOf(_307_v2, _307_v2));
          let _index78 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          let _index79 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          let _rhs100 = _418_v93.f20;
          let _rhs101 = _389_v72;
          let _rhs102 = (_418_v93.f20).minus(_315_v7);
          let _lhs77 = _310_v4;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          let _lhs79 = _310_v4;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          _lhs77[_lhs78] = _rhs100;
          _389_v72 = _rhs101;
          _lhs79[_lhs80] = _rhs102;
          let _index80 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          (_310_v4)[_index80] = ((_dafny.ZERO).minus((_418_v93).f19)).plus(_418_v93.f20);
          let _516_v161;
          _516_v161 = _dafny.Seq.of(_307_v2);
          let _index81 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          let _rhs103 = _310_v4;
          let _rhs104 = _414_v89.f21;
          let _rhs105 = _516_v161;
          let _rhs106 = new BigNumber(88);
          let _rhs107 = _417_v92;
          let _lhs81 = _310_v4;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          _310_v4 = _rhs103;
          _308_v3 = _rhs104;
          _516_v161 = _rhs105;
          _lhs81[_lhs82] = _rhs106;
          _417_v92 = _rhs107;
        }
        let _index82 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_306_v1).length));
        (_306_v1)[_index82] = (!(_414_v89.f21)) === ((_418_v93).f18);
        let _index83 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_306_v1).length));
        (_306_v1)[_index83] = _414_v89.f21;
        let _517_v162;
        let _nw86 = new _module.C1();
        _nw86.__ctor(((true) ? ((_418_v93).f18) : ((_module.D0.create_DC0(_308_v3)).dtor_cf0)), _418_v93.f20, _418_v93.f20, _414_v89.f21);
        _517_v162 = _nw86;
        if (((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))]).isLessThan(new BigNumber((_307_v2).length))) {
          (_309_globalState).f5 = _315_v7;
          (_418_v93).f20 = ((_418_v93).f19).multipliedBy(_418_v93.f20);
          let _518_v163;
          _518_v163 = _module.D2.create_DC6(_497_v154);
          let _519_v164;
          _519_v164 = _module.D2.create_DC6((_518_v163).dtor_cf10);
          let _520_v165;
          let _nw87 = new _module.C0();
          _nw87.__ctor(_519_v164, _418_v93.f20, (_418_v93).f18);
          _520_v165 = _nw87;
          let _521_v166;
          _521_v166 = _dafny.Seq.of(_520_v165, _520_v165, _520_v165);
          let _522_v167;
          _522_v167 = _dafny.Map.Empty.slice().updateUnsafe((_521_v166)[_module.__default.safeIndex((_418_v93).f19, new BigNumber((_521_v166).length))],_414_v89.f21);
          _522_v167 = (_522_v167).update(_520_v165, (_418_v93).f18);
          let _523_v168;
          _523_v168 = _module.D2.create_DC5(_387_v70);
          let _524_v169;
          _524_v169 = _module.D2.create_DC7(_523_v168);
          let _525_v170;
          _525_v170 = _module.D2.create_DC7(_523_v168);
          let _526_v171;
          _526_v171 = _dafny.Map.Empty.slice().updateUnsafe(_525_v170,_dafny.Map.Empty.slice().updateUnsafe(_517_v162.f21,new BigNumber((_307_v2).length)));
          _526_v171 = (_526_v171).update(((false) ? (_525_v170) : (_525_v170)), _426_v101);
          let _527_v172;
          _527_v172 = _dafny.Map.Empty.slice().updateUnsafe(_418_v93.f20,(_418_v93).f18);
          _527_v172 = (_527_v172).update(_module.__default.safeDivisionInt(_418_v93.f20, (_418_v93).f19), !(!(_module.__default.fm1((_418_v93).f19, _517_v162.f21, (_dafny.MultiSet.fromElements((_418_v93).f18)).update((_418_v93).fm5(_517_v162.f21, _418_v93.f20, _309_globalState), _module.__default.abs((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))])), _309_globalState))));
        } else {
          let _528_v173;
          let _nw88 = Array((new BigNumber(7)).toNumber()).fill(false);
          _528_v173 = _nw88;
          let _529_v174;
          let _530_v175;
          let _out36;
          let _out37;
          let _outcollector13 = (_418_v93).m2((_497_v154).update((_dafny.ZERO).minus(_418_v93.f20), _module.__default.abs(_418_v93.f20)), new BigNumber((_dafny.Seq.of(_414_v89.f21)).length), _528_v173, _dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_531_v92) => function (_532_i10) {
            return _531_v92;
          })(_417_v92)), _309_globalState);
          _out36 = _outcollector13[0];
          _out37 = _outcollector13[1];
          _529_v174 = _out36;
          _530_v175 = _out37;
          let _533_v176;
          let _nw89 = Array((new BigNumber(15)).toNumber());
          _533_v176 = _nw89;
          let _index84 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_533_v176).length));
          (_533_v176)[_index84] = _414_v89;
          let _index85 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_533_v176).length));
          (_533_v176)[_index85] = _517_v162;
          let _534_v177;
          _534_v177 = _dafny.Map.Empty.slice().updateUnsafe((((_427_v102).dtor_cf3) ? (_307_v2) : (_307_v2)),_414_v89.f21);
          _534_v177 = (_534_v177).update(_dafny.Seq.Concat(_307_v2, _307_v2), _414_v89.f21);
          let _535_v178;
          _535_v178 = _module.D2.create_DC6(_module.__default.fm12(_309_globalState));
          let _536_v179;
          _536_v179 = _dafny.Map.Empty.slice().updateUnsafe(_315_v7,_dafny.MultiSet.FromArray(_419_v94));
          let _537_v180;
          let _nw90 = new _module.C0();
          _nw90.__ctor(_module.D2.create_DC6(_497_v154), (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _module.__default.fm1(new BigNumber((_307_v2).length), _414_v89.f21, (((_536_v179).contains(_315_v7)) ? ((_536_v179).get(_315_v7)) : (_422_v97)), _309_globalState));
          _537_v180 = _nw90;
          let _538_v181;
          _538_v181 = _dafny.Map.Empty.slice().updateUnsafe(_537_v180,!((_418_v93).f18));
          let _539_v182;
          let _nw91 = new _module.C0();
          _nw91.__ctor(_535_v178, (new BigNumber((_538_v181).length)).minus(new BigNumber((_422_v97).cardinality())), !(_414_v89.f21));
          _539_v182 = _nw91;
          _539_v182 = _539_v182;
          (_309_globalState).f13 = _517_v162.f21;
        }
      } else {
        if (((_414_v89.f21) ? (_308_v3) : (false))) {
          let _540_v183;
          let _541_v184;
          let _out38;
          let _out39;
          let _outcollector14 = (_414_v89).m2(_dafny.MultiSet.fromElements(new BigNumber(608)), _418_v93.f20, _306_v1, _307_v2, _309_globalState);
          _out38 = _outcollector14[0];
          _out39 = _outcollector14[1];
          _540_v183 = _out38;
          _541_v184 = _out39;
          let _init18 = ((_542_v89) => function (_543_i11) {
            return _542_v89.f21;
          })(_414_v89);
          let _nw92 = Array((new BigNumber(14)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw92.length); _i0_18++) {
            _nw92[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _306_v1 = _nw92;
          let _544_v185;
          _544_v185 = _dafny.Map.Empty.slice().updateUnsafe(_419_v94,_425_v100);
          _544_v185 = (_544_v185).update(_dafny.Seq.update(_dafny.Seq.of(_308_v3, _414_v89.f21, true, _308_v3, _414_v89.f21), _module.__default.safeIndex(new BigNumber(816), new BigNumber((_dafny.Seq.of(_308_v3, _414_v89.f21, true, _308_v3, _414_v89.f21)).length)), _414_v89.f21), _425_v100);
          let _545_v186;
          _545_v186 = _dafny.Set.fromElements((_418_v93).f18);
          let _546_v187;
          _546_v187 = _dafny.MultiSet.fromElements(_418_v93.f20, new BigNumber((_545_v186).length));
          let _pat_let_tv16 = _546_v187;
          let _547_v188;
          let _nw93 = new _module.C0();
          _nw93.__ctor(function (_pat_let10_0) {
            return function (_550_dt__update__tmp_h3) {
              return function (_pat_let13_0) {
                return function (_551_dt__update_hcf10_h2) {
                  return _module.D2.create_DC6(_551_dt__update_hcf10_h2);
                }(_pat_let13_0);
              }(_dafny.MultiSet.fromElements(new BigNumber(286)));
            }(_pat_let10_0);
          }(function (_pat_let11_0) {
            return function (_548_dt__update__tmp_h2) {
              return function (_pat_let12_0) {
                return function (_549_dt__update_hcf10_h1) {
                  return _module.D2.create_DC6(_549_dt__update_hcf10_h1);
                }(_pat_let12_0);
              }(_pat_let_tv16);
            }(_pat_let11_0);
          }(_module.D2.create_DC6(_546_v187))), (_418_v93).f19, (_418_v93).f18);
          _547_v188 = _nw93;
          let _552_v189;
          _552_v189 = _dafny.Set.fromElements((_module.D6.create_DC14(_547_v188)).dtor_cf20);
          let _553_v190;
          _553_v190 = _dafny.MultiSet.fromElements((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], (_418_v93).f19, new BigNumber((_552_v189).length), new BigNumber(360));
          let _554_v191;
          let _555_v192;
          let _out40;
          let _out41;
          let _outcollector15 = (_414_v89).m2(_546_v187, _418_v93.f20, _306_v1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xduc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-917)), ((_556_v92) => function (_557_i12) {
            return _556_v92;
          })(_417_v92))), _309_globalState);
          _out40 = _outcollector15[0];
          _out41 = _outcollector15[1];
          _554_v191 = _out40;
          _555_v192 = _out41;
          let _558_v193;
          _558_v193 = _dafny.Map.Empty.slice().updateUnsafe((_418_v93).f18,_417_v92);
          let _559_v194;
          _559_v194 = _dafny.Map.Empty.slice().updateUnsafe(_414_v89.f21,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("qafu"), _module.__default.safeIndex((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], new BigNumber((_dafny.Seq.UnicodeFromString("qafu")).length)), _417_v92));
          let _560_v195;
          let _561_v196;
          let _out42;
          let _out43;
          let _outcollector16 = (_414_v89).m2((_dafny.MultiSet.fromElements(_418_v93.f20, _module.__default.fm0((_418_v93).f19, (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _module.__default.fm0(new BigNumber(81), _315_v7, (_418_v93).f19, (_418_v93).f19, _309_globalState), (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _309_globalState))).Intersect(_553_v190), _module.__default.safeDivisionInt(_module.__default.fm0(new BigNumber((_555_v192).length), new BigNumber(892), _module.__default.fm0((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], (_418_v93).f19, new BigNumber((_552_v189).length), _418_v93.f20, _309_globalState), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_418_v93).f19,true)).length), _309_globalState), new BigNumber((_558_v193).length)), _306_v1, (((_559_v194).contains((_418_v93).f18)) ? ((_559_v194).get((_418_v93).f18)) : (_307_v2)), _309_globalState);
          _out42 = _outcollector16[0];
          _out43 = _outcollector16[1];
          _560_v195 = _out42;
          _561_v196 = _out43;
        } else {
          let _562_v197;
          _562_v197 = _dafny.Set.fromElements(_417_v92);
          let _pat_let_tv17 = _562_v197;
          let _pat_let_tv18 = _414_v89;
          let _pat_let_tv19 = _418_v93;
          let _pat_let_tv20 = _418_v93;
          let _pat_let_tv21 = _309_globalState;
          let _pat_let_tv22 = _418_v93;
          _427_v102 = function (_pat_let14_0) {
            return function (_563_dt__update__tmp_h4) {
              return function (_pat_let15_0) {
                return function (_564_dt__update_hcf3_h0) {
                  return function (_pat_let16_0) {
                    return function (_565_dt__update_hcf1_h0) {
                      return _module.D0.create_DC1(_565_dt__update_hcf1_h0, (_563_dt__update__tmp_h4).dtor_cf2, _564_dt__update_hcf3_h0);
                    }(_pat_let16_0);
                  }(_pat_let_tv22.f20);
                }(_pat_let15_0);
              }((_pat_let_tv17).IsSubsetOf(_module.__default.fm14(_pat_let_tv18.f21, (_pat_let_tv19).f18, (_pat_let_tv20).f19, _pat_let_tv21)));
            }(_pat_let14_0);
          }(_427_v102);
          _426_v101 = (_426_v101).update(false, (_418_v93).f19);
          let _566_v198;
          _566_v198 = _dafny.MultiSet.fromElements(new BigNumber((_495_v152).length), new BigNumber(122));
          let _567_v200;
          _567_v200 = _dafny.MultiSet.fromElements(_307_v2, _307_v2);
          let _568_v201;
          let _569_v202;
          let _out44;
          let _out45;
          let _outcollector17 = (_414_v89).m2(_566_v198, new BigNumber((function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of (_567_v200).Elements) {
              let _570_v199 = _compr_10;
              if ((_567_v200).contains(_570_v199)) {
                _coll10.push([_570_v199,(_418_v93).f18]);
              }
            }
            return _coll10;
          }()).length), _306_v1, _307_v2, _309_globalState);
          _out44 = _outcollector17[0];
          _out45 = _outcollector17[1];
          _568_v201 = _out44;
          _569_v202 = _out45;
          let _571_v203;
          _571_v203 = _dafny.Map.Empty.slice().updateUnsafe(_306_v1,_310_v4);
          let _572_v204;
          _572_v204 = _module.D7.create_DC18(_571_v203);
          _571_v203 = ((_572_v204).dtor_cf25).Merge((_571_v203).Merge(_dafny.Map.Empty.slice().updateUnsafe(_306_v1,_310_v4)));
          (_414_v89).f21 = ((_module.__default.fm7((_418_v93).f19, (_418_v93).f19, (_418_v93).f18, _309_globalState)).Intersect(_422_v97)).IsSubsetOf(_dafny.MultiSet.fromElements(_308_v3, (_418_v93).f18));
        }
        let _573_v205;
        _573_v205 = _dafny.Seq.of(_306_v1, _306_v1, _306_v1);
        let _574_v206;
        let _nw94 = Array((new BigNumber(15)).toNumber());
        _nw94[(_dafny.ZERO).toNumber()] = (_573_v205)[_module.__default.safeIndex(new BigNumber((_422_v97).cardinality()), new BigNumber((_573_v205).length))];
        _nw94[(_dafny.ONE).toNumber()] = _306_v1;
        _nw94[(new BigNumber(2)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(3)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(4)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(5)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(6)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(7)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(8)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(9)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(10)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(11)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(12)).toNumber()] = (_573_v205)[_module.__default.safeIndex(_418_v93.f20, new BigNumber((_573_v205).length))];
        _nw94[(new BigNumber(13)).toNumber()] = _306_v1;
        _nw94[(new BigNumber(14)).toNumber()] = _306_v1;
        _574_v206 = _nw94;
        _574_v206 = _574_v206;
        let _575_v207;
        _575_v207 = _module.D6.create_DC17((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], _308_v3);
        let _rhs108 = _418_v93.f20;
        let _rhs109 = (((_418_v93).f19).multipliedBy((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))])).isLessThan(new BigNumber(925));
        let _rhs110 = _315_v7;
        let _rhs111 = (_575_v207).dtor_cf23;
        let _lhs83 = _418_v93;
        let _lhs84 = _309_globalState;
        let _lhs85 = _309_globalState;
        let _lhs86 = _309_globalState;
        _lhs83.f20 = _rhs108;
        _lhs84.f13 = _rhs109;
        _lhs85.f5 = _rhs110;
        _lhs86.f9 = _rhs111;
        if (!(_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.update(_307_v2, _module.__default.safeIndex((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], new BigNumber((_307_v2).length)), _417_v92), _dafny.Seq.UnicodeFromString("fa")), _dafny.Seq.UnicodeFromString("a")))) {
          (_309_globalState).f13 = !((_module.D1.create_DC3(_315_v7)).dtor_cf5).isEqualTo(_418_v93.f20);
          (_309_globalState).f9 = _418_v93.f20;
          let _576_v208;
          let _577_v209;
          let _578_v210;
          let _579_v211;
          let _out46;
          let _out47;
          let _out48;
          let _out49;
          let _outcollector18 = (_418_v93).m1(_414_v89.f21, (_418_v93).f18, _309_globalState);
          _out46 = _outcollector18[0];
          _out47 = _outcollector18[1];
          _out48 = _outcollector18[2];
          _out49 = _outcollector18[3];
          _576_v208 = _out46;
          _577_v209 = _out47;
          _578_v210 = _out48;
          _579_v211 = _out49;
          let _580_v212;
          _580_v212 = _module.D7.create_DC19(_414_v89.f21, _417_v92, _577_v209);
          let _581_v213;
          let _nw95 = new _module.C1();
          _nw95.__ctor(_414_v89.f21, (_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))], (_418_v93).f19, (_580_v212).dtor_cf28);
          _581_v213 = _nw95;
          let _582_v214;
          _582_v214 = _dafny.MultiSet.fromElements((_418_v93).f19);
          let _583_v215;
          _583_v215 = _module.D2.create_DC6(_582_v214);
          let _584_v216;
          let _nw96 = new _module.C0();
          _nw96.__ctor(_583_v215, (_418_v93).f19, _308_v3);
          _584_v216 = _nw96;
          let _585_v217;
          _585_v217 = _dafny.Map.Empty.slice().updateUnsafe(_584_v216,new BigNumber(-175));
          let _586_v218;
          _586_v218 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(888),_307_v2);
          let _587_v219;
          _587_v219 = _dafny.MultiSet.fromElements(_585_v217, _dafny.Map.Empty.slice().updateUnsafe(_584_v216,new BigNumber((_586_v218).length)));
          let _588_v220;
          let _nw97 = new _module.C1();
          _nw97.__ctor((_418_v93).f18, _315_v7, (new BigNumber((_module.__default.fm13(_581_v213.f21, (_418_v93).f18, (((_587_v219).contains(_585_v217)) ? ((_587_v219).get(_585_v217)) : (_418_v93.f20)), (_dafny.ZERO).minus((_584_v216).f23), _309_globalState)).length)).minus(new BigNumber((_dafny.MultiSet.FromArray(_425_v100)).cardinality())), _414_v89.f21);
          _588_v220 = _nw97;
        } else {
          (_309_globalState).f9 = ((_module.__default.fm1((_425_v100)[_module.__default.safeIndex(_418_v93.f20, new BigNumber((_425_v100).length))], _308_v3, _422_v97, _309_globalState)) ? ((_418_v93).f19) : ((_418_v93).f19));
          let _index86 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length));
          (_310_v4)[_index86] = ((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))]).plus(_418_v93.f20);
          (_309_globalState).f5 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_418_v93.f20),function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of _dafny.IntegerRange(new BigNumber(800), new BigNumber(926))) {
              let _589_v221 = _compr_11;
              if (((new BigNumber(800)).isLessThanOrEqualTo(_589_v221)) && ((_589_v221).isLessThan(new BigNumber(926)))) {
                _coll11.push([(_589_v221).multipliedBy((_418_v93).f19),_315_v7]);
              }
            }
            return _coll11;
          }())).length);
          let _590_v222;
          _590_v222 = _dafny.MultiSet.fromElements(new BigNumber((_495_v152).length));
          let _591_v223;
          _591_v223 = _module.D2.create_DC6(_590_v222);
          let _592_v224;
          let _nw98 = new _module.C0();
          _nw98.__ctor(_591_v223, _418_v93.f20, true);
          _592_v224 = _nw98;
          let _593_v225;
          _593_v225 = _dafny.Seq.of(_592_v224, _592_v224);
          let _594_v226;
          let _nw99 = new _module.C0();
          _nw99.__ctor(_591_v223, new BigNumber((_593_v225).length), _308_v3);
          _594_v226 = _nw99;
          let _595_v227;
          _595_v227 = _dafny.Set.fromElements(_594_v226);
          (_309_globalState).f3 = new BigNumber((_595_v227).length);
          let _596_v228;
          _596_v228 = _dafny.Set.fromElements(_307_v2, _dafny.Seq.UnicodeFromString("njxobfr"), _307_v2, _307_v2);
          let _597_v230;
          _597_v230 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_596_v228).Elements) {
              let _598_v229 = _compr_12;
              if ((_596_v228).contains(_598_v229)) {
                _coll12.add(_598_v229);
              }
            }
            return _coll12;
          }(),_310_v4);
          _597_v230 = (_597_v230).update(_596_v228, _310_v4);
        }
        let _599_v231;
        _599_v231 = _dafny.MultiSet.fromElements(_418_v93.f20, (_418_v93.f20).plus((_418_v93).f19), new BigNumber((_425_v100).length));
        _599_v231 = ((_599_v231).Difference(_599_v231)).update(new BigNumber((_419_v94).length), _module.__default.abs((_310_v4)[_module.__default.safeIndex(new BigNumber(150), new BigNumber((_310_v4).length))]));
      }
      process.stdout.write(_dafny.toString((_306_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_307_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_308_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState.f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_309_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_309_globalState.f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_309_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_309_globalState).f7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_309_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_309_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_309_globalState).f17)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_310_v4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_313_v5).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_314_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_315_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_316_v8).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_387_v70).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(604),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v72).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_414_v89.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_415_v90).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_416_v91).dtor_cf13.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_416_v91).dtor_cf13).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_416_v91).dtor_cf13.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_416_v91).dtor_cf13).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_417_v92));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_418_v93).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_418_v93.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_418_v93).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_419_v94, _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_420_v95).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_421_v96).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_422_v97).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_423_v98).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_424_v99, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_425_v100, _dafny.Seq.of(new BigNumber(-186)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_426_v101).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_427_v102).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_427_v102).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_427_v102).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_428_v103).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_429_v104).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_430_v105).dtor_cf9).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(4),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[_dafny.ZERO])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[_dafny.ONE])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(2)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(3)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(4)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(5)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(6)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(7)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(8)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(9)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(10)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(11)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(12)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(13)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(14)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(15)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(16)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(17)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(18)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(19)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(20)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(21)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(22)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(23)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(24)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(25)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(26)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_484_v149)[new BigNumber(27)])).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_487_v150)).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.ZERO))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[_dafny.ZERO]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[_dafny.ONE]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(2)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(3)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(4)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(5)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(6)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(7)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(8)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(9)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(10)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(11)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(12)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(13)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(14)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(15)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(16)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(17)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(18)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(19)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(20)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_494_v151)[new BigNumber(21)]).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_495_v152).equals(_dafny.Set.fromElements(new BigNumber(2), new BigNumber(-525)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC4(cf6, cf7, cf8) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.ZERO);
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
    static create_DC6(cf10) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D2(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.MultiSet.Empty);
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
    static create_DC8(cf12) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf14, cf15, cf16) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf13) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC11(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC13(cf19) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_dafny.ZERO);
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
    static create_DC15(cf21) {
      let $dt = new D6(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC16(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC17(cf23, cf24) {
      let $dt = new D6(2);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D6(3);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf20 === other.cf20;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(null);
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
    static create_DC19(cf26, cf27, cf28) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC20(cf29) {
      let $dt = new D7(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC18(cf25) {
      let $dt = new D7(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC21(cf30) {
      let $dt = new D7(3);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19(false, new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC22(cf31) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf31) + ")";
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
    static create_DC24(cf33, cf34, cf35) {
      let $dt = new D9(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC23(cf32) {
      let $dt = new D9(1);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(false, _dafny.ZERO, false);
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
    static create_DC26(cf37) {
      let $dt = new D10(0);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC27(cf38) {
      let $dt = new D10(1);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC25(cf36) {
      let $dt = new D10(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC28(cf39) {
      let $dt = new D10(3);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get is_DC28() { return this.$tag === 3; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC25" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf37 === other.cf37;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf36 === other.cf36;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC26(null);
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
    static create_DC30(cf41) {
      let $dt = new D11(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC31(cf42) {
      let $dt = new D11(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC29(cf40) {
      let $dt = new D11(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC32(cf43) {
      let $dt = new D11(3);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get is_DC32() { return this.$tag === 3; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf40 === other.cf40;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30(false);
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
    static create_DC33(cf44) {
      let $dt = new D12(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf46, cf47) {
      let $dt = new D13(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC34(cf45) {
      let $dt = new D13(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC36(cf48) {
      let $dt = new D13(2);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC34() { return this.$tag === 1; }
    get is_DC36() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC35(new _dafny.CodePoint('D'.codePointAt(0)), _module.D9.Default());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
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
      this.f2 = [];
      this.f3 = _dafny.ZERO;
      this.f4 = _dafny.Seq.UnicodeFromString("");
      this.f5 = _dafny.ZERO;
      this.f9 = _dafny.ZERO;
      this.f13 = false;
      this._f0 = [];
      this._f1 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Seq.UnicodeFromString("");
      this._f8 = _dafny.ZERO;
      this._f10 = false;
      this._f11 = _dafny.ZERO;
      this._f12 = false;
      this._f14 = _dafny.ZERO;
      this._f15 = false;
      this._f16 = false;
      this._f17 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
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
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f18 = false;
      this.f22 = _module.D2.Default();
      this._f23 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f22, f23, f18) {
      let _this = this;
      (_this).f22 = f22;
      (_this)._f23 = f23;
      (_this)._f18 = f18;
      return;
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f20 = _dafny.ZERO;
      this._f19 = _dafny.ZERO;
      this._f18 = false;
      this.f21 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
    set f20(value) {
      let _this = this;
      _this._f20 = value;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    __ctor(f21, f19, f20, f18) {
      let _this = this;
      (_this).f21 = f21;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f18 = f18;
      return;
    }
    fm5(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(new BigNumber(354), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), function (_600_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      })).length), _this.f20)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(444))))).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f19));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(989),_this.f20)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_this.f20),_this.f20))).Merge((function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.Seq.of((_dafny.ZERO).minus(_this.f20), (_this).f19)).Elements) {
          let _601_v0 = _compr_13;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(_this.f20), (_this).f19), _601_v0)) {
            _coll13.push([(_601_v0).minus((_this).f19),(_this).f19]);
          }
        }
        return _coll13;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f20,new BigNumber((_dafny.Seq.of(_this.f21)).length))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _602_v0;
      let _init19 = function (_603_i0) {
        return (_603_i0).minus((_this).f19);
      };
      let _nw100 = Array((new BigNumber(8)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw100.length); _i0_19++) {
        _nw100[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _602_v0 = _nw100;
      let _index87 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length));
      (_602_v0)[_index87] = _module.__default.safeModuloInt(_this.f20, _this.f20);
      let _604_v1;
      _604_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f20,_this.f20);
      let _605_v2;
      _605_v2 = _dafny.MultiSet.fromElements(_this.f20);
      let _606_v3;
      _606_v3 = _module.D2.create_DC6(_605_v2);
      let _607_v4;
      _607_v4 = _dafny.MultiSet.fromElements(_this.f21);
      let _608_v5;
      let _nw101 = new _module.C0();
      _nw101.__ctor(_606_v3, (((_607_v4).contains(!(p1))) ? ((_607_v4).get(!(p1))) : (_this.f20)), (_this).f18);
      _608_v5 = _nw101;
      let _609_v6;
      _609_v6 = _dafny.Map.Empty.slice().updateUnsafe(_608_v5,_602_v0);
      let _610_v7;
      _610_v7 = _dafny.Seq.UnicodeFromString("oelkbtu");
      let _index88 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length));
      (_602_v0)[_index88] = ((((_604_v1).contains((_this).f19)) ? ((_604_v1).get((_this).f19)) : (new BigNumber(((_609_v6).update(_608_v5, _602_v0)).length)))).minus(new BigNumber((_610_v7).length));
      let _611_v8;
      let _nw102 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
      _611_v8 = _nw102;
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_611_v8).length))) {
        let _612_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_612_i1)) && ((_612_i1).isLessThan(new BigNumber((_611_v8).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_611_v8, (_612_i1).toNumber(), (_605_v2).update(_this.f20, _module.__default.abs((_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))]))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      let _hi1 = new BigNumber((_610_v7).length);
      for (let _613_i2 = new BigNumber(719); _613_i2.isLessThan(_hi1); _613_i2 = _613_i2.plus(_dafny.ONE)) {
        let _614_v9;
        _614_v9 = _dafny.Set.fromElements(p0, p0, true);
        let _615_v10;
        _615_v10 = _dafny.MultiSet.fromElements(_614_v9, _614_v9, _614_v9);
        (globalState).f5 = (((_604_v1).contains((((_615_v10).contains(_614_v9)) ? ((_615_v10).get(_614_v9)) : ((_this).f19)))) ? ((_604_v1).get((((_615_v10).contains(_614_v9)) ? ((_615_v10).get(_614_v9)) : ((_this).f19)))) : ((((_604_v1).contains(_this.f20)) ? ((_604_v1).get(_this.f20)) : (_613_i2))));
        let _616_v11;
        _616_v11 = _module.D1.create_DC3(new BigNumber(980));
        let _617_v12;
        _617_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_608_v5).f18) ? ((_this).f18) : ((_this).f18)),(_dafny.ZERO).minus((_616_v11).dtor_cf5));
        let _618_v13;
        _618_v13 = _dafny.Map.Empty.slice().updateUnsafe((_608_v5).f18,_module.__default.fm1((_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))], (_608_v5).f18, _607_v4, globalState));
        let _619_v14;
        _619_v14 = _dafny.Map.Empty.slice().updateUnsafe(_608_v5,(_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))]);
        let _620_v15;
        _620_v15 = _dafny.Map.Empty.slice().updateUnsafe(_605_v2,_619_v14);
        _617_v12 = (_617_v12).update((_this).f18, _module.__default.fm0(_613_i2, _613_i2, new BigNumber((_618_v13).length), new BigNumber((_620_v15).length), globalState));
        r0 = (_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))];
        let _index89 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length));
        (_602_v0)[_index89] = (_616_v11).dtor_cf5;
      }
      let _621_v16;
      _621_v16 = _module.D0.create_DC0((_this).f18);
      let _622_v17;
      _622_v17 = _dafny.Seq.of((_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))], _this.f20);
      let _623_v18;
      let _nw103 = Array((new BigNumber(11)).toNumber());
      _nw103[(_dafny.ZERO).toNumber()] = true;
      _nw103[(_dafny.ONE).toNumber()] = false;
      _nw103[(new BigNumber(2)).toNumber()] = p1;
      _nw103[(new BigNumber(3)).toNumber()] = !((_621_v16).dtor_cf0);
      _nw103[(new BigNumber(4)).toNumber()] = !((_this.f20).isEqualTo(_this.f20));
      _nw103[(new BigNumber(5)).toNumber()] = p1;
      _nw103[(new BigNumber(6)).toNumber()] = !(false);
      _nw103[(new BigNumber(7)).toNumber()] = p0;
      _nw103[(new BigNumber(8)).toNumber()] = (_608_v5).f18;
      _nw103[(new BigNumber(9)).toNumber()] = p1;
      _nw103[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_624_v0) => function (_625_i3) {
        return (_dafny.ZERO).minus((_624_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_624_v0).length))]);
      })(_602_v0)), _622_v17);
      _623_v18 = _nw103;
      let _index90 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_623_v18).length));
      (_623_v18)[_index90] = (_this).fm5(false, _this.f20, globalState);
      let _index91 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_623_v18).length));
      (_623_v18)[_index91] = (((_607_v4).IsSubsetOf(_607_v4)) ? (p0) : (p0));
      _607_v4 = (_dafny.MultiSet.fromElements(true, p0)).Union((_module.__default.fm7((_this).f19, (_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))], p1, globalState)).Union(_607_v4));
      let _626_v19;
      _626_v19 = _dafny.Seq.of((_608_v5).f18, (_623_v18)[_module.__default.safeIndex(new BigNumber(650), new BigNumber((_623_v18).length))]);
      r3 = ((((_this).f18) ? (_this.f20) : ((((_605_v2).contains(new BigNumber((_626_v19).length))) ? ((_605_v2).get(new BigNumber((_626_v19).length))) : (_this.f20))))).minus((_this).f19);
      r0 = (_this).f19;
      let _627_v20;
      _627_v20 = _dafny.Seq.of(_608_v5);
      r1 = _dafny.Seq.contains(_627_v20, _608_v5);
      r2 = (_602_v0)[_module.__default.safeIndex(new BigNumber(449), new BigNumber((_602_v0).length))];
      r3 = _this.f20;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Set.Empty;
      let _628_v0;
      let _nw104 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _628_v0 = _nw104;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_628_v0).length))) {
        let _629_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_629_i0)) && ((_629_i0).isLessThan(new BigNumber((_628_v0).length))))) {
          (_628_v0)[(_629_i0)] = (_629_i0).minus(_this.f20);
        }
      }
      let _630_v1;
      _630_v1 = _module.D2.create_DC6(p0);
      let _631_v2;
      _631_v2 = _dafny.Seq.of(p1);
      let _pat_let_tv23 = _630_v1;
      let _632_v3;
      _632_v3 = _dafny.MultiSet.fromElements(_630_v1, _630_v1, function (_pat_let17_0) {
        return function (_633_dt__update__tmp_h0) {
          return function (_pat_let18_0) {
            return function (_634_dt__update_hcf10_h0) {
              return _module.D2.create_DC6(_634_dt__update_hcf10_h0);
            }(_pat_let18_0);
          }((_pat_let_tv23).dtor_cf10);
        }(_pat_let17_0);
      }(_module.D2.create_DC6(_dafny.MultiSet.FromArray(_631_v2))));
      let _635_i1;
      _635_i1 = _dafny.ZERO;
      L0: {
        while ((((_this).f18) ? (_this.f21) : ((_632_v3).IsDisjointFrom(_632_v3)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_635_i1)) {
              break L0;
            }
            _635_i1 = (_635_i1).plus(_dafny.ONE);
            let _636_v6;
            let _init20 = function (_637_i2) {
              return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(977),_this.f20),_dafny.Map.Empty.slice().updateUnsafe((_this).f19,_this.f21))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
                let _coll14 = new _dafny.Set();
                for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-686), new BigNumber(230))) {
                  let _638_v4 = _compr_14;
                  if (((new BigNumber(-686)).isLessThanOrEqualTo(_638_v4)) && ((_638_v4).isLessThan(new BigNumber(230)))) {
                    _coll14.add((_638_v4).minus(_this.f20));
                  }
                }
                return _coll14;
              }()).length),_this.f20),function () {
                let _coll15 = new _dafny.Map();
                for (const _compr_15 of _dafny.IntegerRange(new BigNumber(955), new BigNumber(-330))) {
                  let _639_v5 = _compr_15;
                  if (((new BigNumber(955)).isLessThanOrEqualTo(_639_v5)) && ((_639_v5).isLessThan(new BigNumber(-330)))) {
                    _coll15.push([(_639_v5).minus(new BigNumber(32)),false]);
                  }
                }
                return _coll15;
              }()));
            };
            let _nw105 = Array((new BigNumber(29)).toNumber());
            for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw105.length); _i0_20++) {
              _nw105[_i0_20] = _init20(new BigNumber(_i0_20));
            }
            _636_v6 = _nw105;
            let _640_v7;
            _640_v7 = new _dafny.CodePoint('e'.codePointAt(0));
            let _641_v8;
            _641_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_640_v7)).length),new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("yh"), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.UnicodeFromString("yh")).length)), _640_v7)).length));
            let _642_v9;
            _642_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f19,_this.f21);
            let _643_v10;
            _643_v10 = _dafny.Map.Empty.slice().updateUnsafe(_641_v8,_642_v9);
            let _index92 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_636_v6).length));
            (_636_v6)[_index92] = _643_v10;
            let _index93 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_636_v6).length));
            (_636_v6)[_index93] = ((_this.f21) ? (((true) ? (_643_v10) : (_643_v10))) : (function () {
              let _coll16 = new _dafny.Map();
              for (const _compr_16 of (_dafny.Set.fromElements(_641_v8)).Elements) {
                let _644_v11 = _compr_16;
                if ((_dafny.Set.fromElements(_641_v8)).contains(_644_v11)) {
                  _coll16.push([_644_v11,_642_v9]);
                }
              }
              return _coll16;
            }()));
            if (_dafny.Seq.contains(_module.__default.fm3(globalState), _640_v7)) {
              (globalState).f5 = p1;
              let _645_v12;
              _645_v12 = _dafny.Set.fromElements((_this).f19);
              let _index94 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((p2).length));
              (p2)[_index94] = (_645_v12).IsSubsetOf(_645_v12);
              let _646_v13;
              _646_v13 = _module.D0.create_DC0((_this).f18);
              let _index95 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((p2).length));
              (p2)[_index95] = (_646_v13).dtor_cf0;
              (globalState).f5 = (_dafny.ZERO).minus(_this.f20);
              let _647_v14;
              _647_v14 = _dafny.Set.fromElements(_this.f21, _this.f21);
              _647_v14 = (_647_v14).Intersect((_647_v14).Intersect(_647_v14));
              let _648_v15;
              _648_v15 = _dafny.Map.Empty.slice().updateUnsafe(_631_v2,(_this).f19);
              _648_v15 = (_648_v15).update(_631_v2, (new BigNumber(706)).multipliedBy((_dafny.ZERO).minus(new BigNumber((p3).length))));
            } else {
              (globalState).f13 = !((_this).f18);
              let _649_v16;
              _649_v16 = _dafny.Set.fromElements(_this.f21, _this.f21, (_this).f18);
              let _650_v17;
              _650_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(811),_649_v16);
              _650_v17 = _650_v17;
              let _651_v18;
              _651_v18 = _module.D1.create_DC4((_this).f18, false, _this.f21);
              (globalState).f13 = (((_651_v18).dtor_cf7) ? ((_this).f18) : (_this.f21));
              let _652_v19;
              _652_v19 = _dafny.Set.fromElements(_this.f20, (_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), _module.__default.fm0(p1, new BigNumber(-806), p1, p1, globalState));
              let _653_v20;
              let _nw106 = new _module.C0();
              _nw106.__ctor(_module.D2.create_DC6(p0), (_this).f19, _this.f21);
              _653_v20 = _nw106;
              let _654_v21;
              _654_v21 = _dafny.Map.Empty.slice().updateUnsafe(_652_v19,_653_v20);
              _654_v21 = (_654_v21).update((_652_v19).Difference(_652_v19), _653_v20);
              let _655_v22;
              _655_v22 = _dafny.Seq.of(p3);
              let _656_v23;
              _656_v23 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.IsProperPrefixOf(_655_v22, _655_v22));
              _656_v23 = _656_v23;
            }
            let _index96 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_628_v0).length));
            (_628_v0)[_index96] = _module.__default.safeDivisionInt((_this).f19, _this.f20);
            let _index97 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_628_v0).length));
            (_628_v0)[_index97] = (p1).multipliedBy((_this).f19);
            let _657_v24;
            let _nw107 = new _module.C0();
            _nw107.__ctor(_630_v1, (_this).f19, (_this).f18);
            _657_v24 = _nw107;
          }
        }
      }
      let _658_v25;
      let _nw108 = new _module.C0();
      _nw108.__ctor(_module.D2.create_DC6(p0), p1, _this.f21);
      _658_v25 = _nw108;
      let _659_v26;
      _659_v26 = _dafny.Set.fromElements(_658_v25, _658_v25);
      let _index98 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_628_v0).length));
      (_628_v0)[_index98] = (new BigNumber((_659_v26).length)).minus(new BigNumber(825));
      let _index99 = _module.__default.safeIndex(new BigNumber(658), new BigNumber((_628_v0).length));
      (_628_v0)[_index99] = p1;
      let _660_v27;
      _660_v27 = new _dafny.CodePoint('y'.codePointAt(0));
      let _661_v28;
      _661_v28 = _dafny.Set.fromElements(_660_v27, _660_v27, _660_v27);
      let _662_v29;
      _662_v29 = _dafny.MultiSet.fromElements(_661_v28);
      let _663_v30;
      _663_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f18,(_662_v29).update(_661_v28, _module.__default.abs(new BigNumber((_631_v2).length))));
      let _664_v31;
      _664_v31 = _dafny.Seq.of((((_663_v30).contains(_this.f21)) ? ((_663_v30).get(_this.f21)) : (_662_v29)), _dafny.MultiSet.fromElements(_661_v28, _dafny.Set.fromElements(_660_v27)));
      _662_v29 = (_664_v31)[_module.__default.safeIndex((new BigNumber((p3).length)).multipliedBy(_this.f20), new BigNumber((_664_v31).length))];
      let _665_v32;
      let _nw109 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _665_v32 = _nw109;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_665_v32).length))) {
        let _666_i3 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_666_i3)) && ((_666_i3).isLessThan(new BigNumber((_665_v32).length))))) {
          (_665_v32)[(_666_i3)] = _dafny.Seq.Concat(p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-898)), ((_667_v27) => function (_668_i4) {
            return _667_v27;
          })(_660_v27)));
        }
      }
      _628_v0 = _628_v0;
      let _669_v33;
      let _init21 = ((_670_v25) => function (_671_i5) {
        return _670_v25.f22;
      })(_658_v25);
      let _nw110 = Array((new BigNumber(22)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw110.length); _i0_21++) {
        _nw110[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _669_v33 = _nw110;
      r0 = _669_v33;
      r1 = _dafny.Set.fromElements((_this).f19, (_658_v25).f23);
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
