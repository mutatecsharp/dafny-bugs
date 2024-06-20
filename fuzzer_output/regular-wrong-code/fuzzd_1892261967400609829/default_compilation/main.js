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
      if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true, false, false)).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(252)), function (_0_i1) {
        return new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality());
      })).length)))).IsSubsetOf(_dafny.MultiSet.fromElements(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(294), new BigNumber(982))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(294)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(982)))) {
            _coll0.push([(_1_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(549)), function (_2_i0) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).length);
            }))).cardinality()))).length)),new BigNumber((_dafny.Set.fromElements(false)).length)]);
          }
        }
        return _coll0;
      }()))) {
        return true;
      } else {
        return false;
      }
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return _module.__default.safeModuloInt(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).length), (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-146))).cardinality())).multipliedBy(new BigNumber(-601)));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("t");
    };
    static fm3(p0, p1, globalState) {
      return _dafny.Set.fromElements(true, true, (_dafny.MultiSet.fromElements(new BigNumber(693), new BigNumber(689))).IsProperSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()), new BigNumber((_dafny.Seq.of(true, true)).length))));
    };
    static fm4(p0, p1, globalState) {
      return new _dafny.CodePoint('p'.codePointAt(0));
    };
    static fm6(p0, globalState) {
      return _dafny.Seq.of(!((!(true)) === (false)), !(!(true)), _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.Set.fromElements((_module.D8.create_DC20(false, new BigNumber(962))).dtor_cf26, new BigNumber(-800))), _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-845)))), false, (new BigNumber(782)).isEqualTo(new BigNumber(905)));
    };
    static fm7(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(522))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!((_module.D3.create_DC9(true, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("lihiauqtn")).length),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(654))).length), new BigNumber(-799))).length)))).dtor_cf13),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(631),false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber(690))).length))));
    };
    static fm8(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC9(false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(70),new BigNumber(280)))).dtor_cf13,new BigNumber(299))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-53)), function (_3_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())));
      })), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Set.fromElements(true)).length))));
    };
    static fm9(p0, p1, globalState) {
      return function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-674)), function (_4_i0) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }),!(true))).Keys.Elements) {
          let _5_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-674)), function (_4_i0) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }),!(true))).contains(_5_v0)) {
            _coll1.add(_5_v0);
          }
        }
        return _coll1;
      }();
    };
    static fm10(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-122),_module.D3.create_DC9(false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).length))).length),(_dafny.ZERO).minus(new BigNumber(-32)))));
    };
    static fm11(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(333)), function (_6_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("hqgfpgnr")).length);
      });
    };
    static fm12(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(68),new BigNumber(47))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),new BigNumber((_dafny.Seq.of(new BigNumber(79))).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-498)),(_module.D0.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).dtor_cf3)).Merge(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(310), new BigNumber(-947))) {
          let _7_v0 = _compr_2;
          if (((new BigNumber(310)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(-947)))) {
            _coll2.push([(_7_v0).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).cardinality())),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cjm")).length))]);
          }
        }
        return _coll2;
      }()));
    };
    static fm13(globalState) {
      return _module.D2.create_DC4((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_module.D3.create_DC9(false, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("iqbk")).length)))).cardinality()),new BigNumber((_dafny.Seq.of(false, false, true, true, true)).length))))).cardinality())))).cardinality())))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-7),new BigNumber((_dafny.Seq.UnicodeFromString("dnpmcdw")).length))));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('y'.codePointAt(0)))).length)).minus(new BigNumber(96)),new _dafny.CodePoint('o'.codePointAt(0)));
    };
    static m0(p0, globalState) {
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = [];
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _8_i0;
      _8_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_8_i0)) {
              break L0;
            }
            _8_i0 = (_8_i0).plus(_dafny.ONE);
            let _9_v0;
            let _nw0 = new _module.C0();
            _nw0.__ctor();
            _9_v0 = _nw0;
            (globalState).f3 = !(p0);
            let _10_v1;
            _10_v1 = _dafny.Seq.of(p0);
            (globalState).f3 = ((p0) ? (p0) : ((_dafny.MultiSet.FromArray(_10_v1)).IsProperSubsetOf(_dafny.MultiSet.fromElements(p0))));
            let _11_v2;
            _11_v2 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jbq"), _dafny.Seq.UnicodeFromString("jjrtckkdd")));
            let _12_v3;
            _12_v3 = new BigNumber(346);
            _11_v2 = _module.__default.fm9(p0, _12_v3, globalState);
          }
        }
      }
      let _13_v4;
      _13_v4 = new BigNumber(346);
      let _14_v5;
      _14_v5 = _module.D0.create_DC2(_13_v4, _13_v4);
      let _15_v6;
      _15_v6 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("dfufllfmk")).length), _13_v4);
      (globalState).f22 = ((_14_v5).dtor_cf2).minus(new BigNumber((_15_v6).length));
      let _16_v7;
      _16_v7 = _module.D2.create_DC6();
      _16_v7 = _16_v7;
      let _source0 = _module.D8.create_DC20(p0, _13_v4);
      if (_source0.is_DC20) {
        let _17___mcc_h0 = (_source0).cf25;
        let _18___mcc_h1 = (_source0).cf26;
        let _19_cf26 = _18___mcc_h1;
        let _20_cf25 = _17___mcc_h0;
        let _21_v8;
        let _init0 = ((_22_p0) => function (_23_i1) {
          return _22_p0;
        })(p0);
        let _nw1 = Array((new BigNumber(18)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _21_v8 = _nw1;
        let _24_v9;
        _24_v9 = _dafny.Map.Empty.slice().updateUnsafe(_13_v4,_13_v4);
        let _25_v10;
        _25_v10 = _dafny.Set.fromElements(_24_v9, (_24_v9).update(new BigNumber((_dafny.Seq.of(_19_cf26)).length), _19_cf26), _24_v9, _24_v9, _24_v9);
        let _index0 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_21_v8).length));
        (_21_v8)[_index0] = (_25_v10).equals(_25_v10);
        let _index1 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_21_v8).length));
        (_21_v8)[_index1] = p0;
        let _26_v11;
        _26_v11 = new _dafny.CodePoint('l'.codePointAt(0));
        let _27_v12;
        _27_v12 = _dafny.MultiSet.fromElements(_13_v4, _13_v4, _19_cf26, _19_cf26);
        let _28_v13;
        _28_v13 = _dafny.Map.Empty.slice().updateUnsafe(_26_v11,(_19_cf26).multipliedBy(new BigNumber((_27_v12).cardinality())));
        _28_v13 = (_28_v13).update(_26_v11, ((p0) ? (_13_v4) : (_13_v4)));
        let _29_v14;
        _29_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_13_v4).isEqualTo(new BigNumber(584)),_module.__default.safeModuloInt(_19_cf26, _13_v4));
        _29_v14 = (_29_v14).Merge(_dafny.Map.Empty.slice().updateUnsafe((_21_v8)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_21_v8).length))],_19_cf26));
        let _30_v15;
        _30_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_21_v8)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_21_v8).length))]);
        _30_v15 = (_30_v15).update((_21_v8)[_module.__default.safeIndex(new BigNumber(554), new BigNumber((_21_v8).length))], p0);
      } else if (_source0.is_DC19) {
        let _31___mcc_h2 = (_source0).cf24;
        let _32_cf24 = _31___mcc_h2;
        (globalState).f14 = ((false) ? (p0) : (((p0) ? (p0) : (p0))));
        let _33_v16;
        _33_v16 = new _dafny.CodePoint('a'.codePointAt(0));
        let _34_v17;
        _34_v17 = _dafny.Map.Empty.slice().updateUnsafe(_33_v16,new _dafny.CodePoint('k'.codePointAt(0)));
        let _35_v18;
        _35_v18 = _dafny.Set.fromElements(p0);
        if ((_13_v4).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_34_v17).length), new BigNumber((_35_v18).length)))) {
          let _36_v19;
          let _nw2 = new _module.C0();
          _nw2.__ctor();
          _36_v19 = _nw2;
          let _37_v20;
          _37_v20 = _dafny.Seq.UnicodeFromString("ntnbowsn");
          (globalState).f21 = _37_v20;
          let _38_v21;
          let _nw3 = new _module.C0();
          _nw3.__ctor();
          _38_v21 = _nw3;
          _33_v16 = _33_v16;
          (globalState).f22 = _13_v4;
        } else {
          (globalState).f14 = p0;
          let _39_v22;
          _39_v22 = _dafny.MultiSet.fromElements(p0);
          _39_v22 = _39_v22;
          (globalState).f25 = (_13_v4).multipliedBy(new BigNumber(-892));
          let _40_v23;
          let _nw4 = Array((new BigNumber(21)).toNumber());
          _40_v23 = _nw4;
          let _41_v24;
          let _nw5 = new _module.C0();
          _nw5.__ctor();
          _41_v24 = _nw5;
          let _42_v25;
          _42_v25 = _dafny.Map.Empty.slice().updateUnsafe(_41_v24,_40_v23);
          let _43_v26;
          _43_v26 = _dafny.Seq.of(_40_v23, (((_42_v25).contains(_41_v24)) ? ((_42_v25).get(_41_v24)) : (_40_v23)));
          (globalState).f3 = (new BigNumber((_dafny.Seq.Concat(_43_v26, _dafny.Seq.of(_40_v23))).length)).isLessThanOrEqualTo(_13_v4);
          let _44_v27;
          _44_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _45_v28;
          _45_v28 = _module.D0.create_DC1(p0);
          let _46_v29;
          _46_v29 = _dafny.Seq.of(_45_v28);
          let _47_v30;
          _47_v30 = _dafny.Map.Empty.slice().updateUnsafe(_44_v27,_46_v29);
          let _48_v31;
          _48_v31 = _dafny.Map.Empty.slice().updateUnsafe((_module.D5.create_DC14(_47_v30, _13_v4)).dtor_cf20,new BigNumber(309));
          let _49_v32;
          _49_v32 = _module.D3.create_DC9(p0, _48_v31);
          let _50_v33;
          _50_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), ((_51_v16) => function (_52_i2) {
            return _51_v16;
          })(_33_v16))).length),_49_v32);
          let _53_v34;
          let _nw6 = Array((new BigNumber(24)).toNumber()).fill(false);
          _53_v34 = _nw6;
          let _54_v35;
          _54_v35 = _module.D4.create_DC11(_50_v33, _53_v34);
          _54_v35 = _54_v35;
        }
        let _55_v36;
        _55_v36 = _dafny.Seq.UnicodeFromString("llmnukb");
        let _56_v37;
        _56_v37 = _dafny.Map.Empty.slice().updateUnsafe(_33_v16,_55_v36);
        let _source1 = _module.D0.create_DC1(_module.__default.fm0(new BigNumber(((((_56_v37).contains(_33_v16)) ? ((_56_v37).get(_33_v16)) : (_55_v36))).length), p0, globalState));
        if (_source1.is_DC1) {
          let _57___mcc_h4 = (_source1).cf1;
          let _58_cf1 = _57___mcc_h4;
          (globalState).f24 = true;
          (globalState).f9 = (((_58_cf1) ? (_13_v4) : (new BigNumber((_dafny.Seq.UnicodeFromString("iwdi")).length)))).multipliedBy(_13_v4);
          let _59_v38;
          _59_v38 = _dafny.Seq.of(_58_cf1);
          let _60_v39;
          _60_v39 = _dafny.Seq.of(_dafny.Seq.Concat(_59_v38, _dafny.Seq.of(p0, p0)));
          r0 = (_60_v39)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("bkf")).length), new BigNumber((_60_v39).length))];
          let _61_v40;
          _61_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_13_v4,_59_v38)).length),new BigNumber(177));
          let _62_v41;
          _62_v41 = _module.D3.create_DC9(_module.__default.fm0(_13_v4, _58_cf1, globalState), (_61_v40).Merge(_61_v40));
          _62_v41 = _62_v41;
        } else if (_source1.is_DC2) {
          let _63___mcc_h5 = (_source1).cf2;
          let _64___mcc_h6 = (_source1).cf3;
          let _65_cf3 = _64___mcc_h6;
          let _66_cf2 = _63___mcc_h5;
          let _67_v42;
          let _nw7 = new _module.C0();
          _nw7.__ctor();
          _67_v42 = _nw7;
          let _68_v43;
          _68_v43 = _dafny.Set.fromElements(_33_v16, _33_v16, _33_v16);
          let _69_v44;
          let _nw8 = Array((new BigNumber(7)).toNumber());
          _nw8[(_dafny.ZERO).toNumber()] = _68_v43;
          _nw8[(_dafny.ONE).toNumber()] = _68_v43;
          _nw8[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_33_v16, _33_v16, _33_v16, _33_v16);
          _nw8[(new BigNumber(3)).toNumber()] = _68_v43;
          _nw8[(new BigNumber(4)).toNumber()] = ((!(p0)) ? (_68_v43) : (_68_v43));
          _nw8[(new BigNumber(5)).toNumber()] = _68_v43;
          _nw8[(new BigNumber(6)).toNumber()] = _68_v43;
          _69_v44 = _nw8;
          let _70_v45;
          let _init1 = ((_71_v36) => function (_72_i3) {
            return (_72_i3).multipliedBy(new BigNumber((_71_v36).length));
          })(_55_v36);
          let _nw9 = Array((new BigNumber(16)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw9.length); _i0_1++) {
            _nw9[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _70_v45 = _nw9;
          let _index2 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_70_v45).length));
          (_70_v45)[_index2] = new BigNumber(551);
          let _73_v46;
          let _init2 = ((_74_p0) => function (_75_i4) {
            return _74_p0;
          })(p0);
          let _nw10 = Array((new BigNumber(5)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw10.length); _i0_2++) {
            _nw10[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _73_v46 = _nw10;
          let _76_v47;
          _76_v47 = _module.D4.create_DC11(_module.__default.fm10(_33_v16, new BigNumber((_15_v6).length), globalState), _73_v46);
          let _77_v48;
          _77_v48 = _dafny.Map.Empty.slice().updateUnsafe(_76_v47,p0);
          let _78_v49;
          let _nw11 = Array((new BigNumber(9)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _77_v48;
          _nw11[(_dafny.ONE).toNumber()] = _77_v48;
          _nw11[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_76_v47,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_76_v47,p0));
          _nw11[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_76_v47,p0);
          _nw11[(new BigNumber(4)).toNumber()] = (_77_v48).Merge((_dafny.Map.Empty.slice().updateUnsafe(_76_v47,p0)).update(_76_v47, p0));
          _nw11[(new BigNumber(5)).toNumber()] = _77_v48;
          _nw11[(new BigNumber(6)).toNumber()] = _77_v48;
          _nw11[(new BigNumber(7)).toNumber()] = (_77_v48).Merge(_77_v48);
          _nw11[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_76_v47,p0);
          _78_v49 = _nw11;
          let _index3 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_78_v49).length));
          (_78_v49)[_index3] = _77_v48;
          let _index4 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_70_v45).length));
          let _index5 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_78_v49).length));
          let _rhs0 = _69_v44;
          let _rhs1 = _65_cf3;
          let _rhs2 = !(!(((p0) ? (p0) : (p0))));
          let _rhs3 = (_77_v48).Merge((_77_v48).Merge(_77_v48));
          let _rhs4 = p0;
          let _lhs0 = _70_v45;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_70_v45).length));
          let _lhs2 = globalState;
          let _lhs3 = _78_v49;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_78_v49).length));
          let _lhs5 = globalState;
          _69_v44 = _rhs0;
          _lhs0[_lhs1] = _rhs1;
          _lhs2.f24 = _rhs2;
          _lhs3[_lhs4] = _rhs3;
          _lhs5.f3 = _rhs4;
          let _79_v50;
          _79_v50 = _dafny.Set.fromElements(_73_v46, _73_v46);
          let _rhs5 = (_79_v50).Intersect(((p0) ? (_79_v50) : (_dafny.Set.fromElements(_73_v46))));
          let _rhs6 = _33_v16;
          let _rhs7 = (_70_v45)[_module.__default.safeIndex(new BigNumber(389), new BigNumber((_70_v45).length))];
          let _lhs6 = globalState;
          _79_v50 = _rhs5;
          r3 = _rhs6;
          _lhs6.f9 = _rhs7;
          let _80_v51;
          _80_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,_66_cf2);
          let _81_v52;
          _81_v52 = _dafny.Map.Empty.slice().updateUnsafe(_80_v51,((p0) ? (!(p0)) : (!(p0))));
          _81_v52 = (_81_v52).update(_80_v51, p0);
        } else {
          let _82___mcc_h7 = (_source1).cf0;
          let _83_cf0 = _82___mcc_h7;
          let _84_v53;
          let _nw12 = new _module.C0();
          _nw12.__ctor();
          _84_v53 = _nw12;
          (globalState).f21 = _dafny.Seq.UnicodeFromString("ukla");
          let _85_v54;
          _85_v54 = _dafny.Seq.of(_module.__default.fm11(p0, globalState), _module.__default.fm11(p0, globalState), _15_v6);
          r1 = (new BigNumber((_module.__default.fm2((_84_v53).fm5(globalState), new BigNumber((_15_v6).length), p0, _13_v4, globalState)).length)).multipliedBy(new BigNumber((_85_v54).length));
          (globalState).f9 = _13_v4;
        }
        let _86_v55;
        let _nw13 = new _module.C0();
        _nw13.__ctor();
        _86_v55 = _nw13;
        let _87_v56;
        let _nw14 = Array((_dafny.ONE).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = _86_v55;
        _87_v56 = _nw14;
        let _index6 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_87_v56).length));
        (_87_v56)[_index6] = _86_v55;
        let _index7 = _module.__default.safeIndex(new BigNumber(359), new BigNumber((_87_v56).length));
        (_87_v56)[_index7] = _86_v55;
      } else {
        let _88___mcc_h3 = (_source0).cf27;
        let _89_cf27 = _88___mcc_h3;
        let _90_v57;
        let _nw15 = new _module.C0();
        _nw15.__ctor();
        _90_v57 = _nw15;
        r1 = new BigNumber(788);
        let _91_v58;
        _91_v58 = _dafny.Seq.of(!(p0));
        let _92_v59;
        _92_v59 = _dafny.Map.Empty.slice().updateUnsafe(_91_v58,p0);
        _92_v59 = _92_v59;
        let _93_v60;
        let _nw16 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _93_v60 = _nw16;
        let _94_v61;
        _94_v61 = new _dafny.CodePoint('x'.codePointAt(0));
        let _index8 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_93_v60).length));
        (_93_v60)[_index8] = ((p0) ? (_94_v61) : (_94_v61));
        let _95_v62;
        _95_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _96_v63;
        _96_v63 = _module.D0.create_DC1(p0);
        let _97_v64;
        _97_v64 = _dafny.Seq.of(_96_v63, _96_v63);
        let _98_v65;
        _98_v65 = _dafny.Map.Empty.slice().updateUnsafe(_95_v62,_97_v64);
        let _99_v66;
        _99_v66 = _module.D5.create_DC14((_98_v65).update(_95_v62, _97_v64), _13_v4);
        let _index9 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_93_v60).length));
        let _rhs8 = (_99_v66).dtor_cf20;
        let _rhs9 = !(!(p0));
        let _rhs10 = new _dafny.CodePoint('c'.codePointAt(0));
        let _rhs11 = ((_dafny.ZERO).minus(_13_v4)).plus(_13_v4);
        let _lhs7 = globalState;
        let _lhs8 = _93_v60;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_93_v60).length));
        let _lhs10 = globalState;
        r1 = _rhs8;
        _lhs7.f24 = _rhs9;
        _lhs8[_lhs9] = _rhs10;
        _lhs10.f25 = _rhs11;
      }
      let _100_v67;
      let _nw17 = new _module.C0();
      _nw17.__ctor();
      _100_v67 = _nw17;
      let _101_v68;
      _101_v68 = _dafny.Map.Empty.slice().updateUnsafe(_100_v67,(_100_v67).fm5(globalState));
      let _102_v69;
      _102_v69 = _dafny.Map.Empty.slice().updateUnsafe((((_101_v68).contains(_100_v67)) ? ((_101_v68).get(_100_v67)) : (new BigNumber(86))),_13_v4);
      let _rhs12 = _13_v4;
      let _rhs13 = (_102_v69).Merge((_102_v69).Merge(_102_v69));
      _13_v4 = _rhs12;
      _102_v69 = _rhs13;
      let _hi0 = _13_v4;
      for (let _103_i5 = _13_v4; _103_i5.isLessThan(_hi0); _103_i5 = _103_i5.plus(_dafny.ONE)) {
        let _104_v70;
        let _nw18 = Array((new BigNumber(2)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = p0;
        _nw18[(_dafny.ONE).toNumber()] = _module.__default.fm0(_13_v4, true, globalState);
        _104_v70 = _nw18;
        let _index10 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_104_v70).length));
        (_104_v70)[_index10] = false;
        let _105_v71;
        _105_v71 = _dafny.MultiSet.fromElements(p0);
        let _index11 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_104_v70).length));
        (_104_v70)[_index11] = _module.__default.fm0(_103_i5, (new BigNumber(417)).isEqualTo((((_105_v71).contains(p0)) ? ((_105_v71).get(p0)) : (_103_i5))), globalState);
        (globalState).f14 = (_104_v70)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_104_v70).length))];
        let _106_v72;
        _106_v72 = _dafny.Seq.of(p0, !(p0), p0, p0);
        let _107_v73;
        _107_v73 = _dafny.Seq.UnicodeFromString("fjmbsow");
        (globalState).f9 = (((_14_v5).dtor_cf2).minus(new BigNumber((_106_v72).length))).multipliedBy(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_108_i6) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _107_v73)).length));
        let _109_v74;
        _109_v74 = _module.D2.create_DC4(_module.__default.fm12(_106_v72, (_104_v70)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_104_v70).length))], globalState));
        _109_v74 = _module.__default.fm13(globalState);
      }
      let _110_v75;
      _110_v75 = _dafny.Seq.of(p0, !(p0), false, false);
      let _111_v76;
      _111_v76 = _dafny.Seq.of(_110_v75);
      let _112_v77;
      _112_v77 = _dafny.MultiSet.fromElements(_13_v4, _13_v4);
      r0 = (_111_v76)[_module.__default.safeIndex((((_112_v77).contains((_dafny.ZERO).minus(_13_v4))) ? ((_112_v77).get((_dafny.ZERO).minus(_13_v4))) : (new BigNumber((_module.__default.fm14(_dafny.Seq.Create(_module.__default.abs(new BigNumber(798)), ((_113_v4) => function (_114_i7) {
        return _113_v4;
      })(_13_v4)), _13_v4, _13_v4, _112_v77, globalState)).length))), new BigNumber((_111_v76).length))];
      r1 = _13_v4;
      let _115_v78;
      let _init3 = ((_116_v4) => function (_117_i8) {
        return _module.__default.safeModuloInt(_117_i8, _116_v4);
      })(_13_v4);
      let _nw19 = Array((new BigNumber(8)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw19.length); _i0_3++) {
        _nw19[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _115_v78 = _nw19;
      let _118_v79;
      _118_v79 = _dafny.Seq.of(_115_v78, _115_v78, _115_v78, _115_v78);
      r2 = (_118_v79)[_module.__default.safeIndex(_13_v4, new BigNumber((_118_v79).length))];
      let _119_v80;
      _119_v80 = new _dafny.CodePoint('f'.codePointAt(0));
      r3 = _119_v80;
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _120_v0;
      _120_v0 = new BigNumber(691);
      let _121_v1;
      _121_v1 = _dafny.Set.fromElements(_120_v0);
      let _122_v2;
      _122_v2 = _dafny.Map.Empty.slice().updateUnsafe(_121_v1,_120_v0);
      let _123_v3;
      _123_v3 = _dafny.Seq.of(_120_v0, _120_v0);
      let _124_v4;
      _124_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_120_v0)).length),_123_v3);
      let _125_v5;
      let _nw20 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _125_v5 = _nw20;
      let _126_v6;
      _126_v6 = _module.D0.create_DC0(_125_v5);
      let _127_v7;
      let _nw21 = Array((new BigNumber(4)).toNumber());
      _nw21[(_dafny.ZERO).toNumber()] = _125_v5;
      _nw21[(_dafny.ONE).toNumber()] = _125_v5;
      _nw21[(new BigNumber(2)).toNumber()] = _125_v5;
      _nw21[(new BigNumber(3)).toNumber()] = (_126_v6).dtor_cf0;
      _127_v7 = _nw21;
      let _128_v8;
      _128_v8 = _dafny.MultiSet.fromElements(new BigNumber((_123_v3).length));
      let _129_v9;
      _129_v9 = _dafny.Seq.UnicodeFromString("px");
      let _130_v10;
      _130_v10 = _dafny.Map.Empty.slice().updateUnsafe(_127_v7,_dafny.Set.fromElements(new BigNumber((_128_v8).cardinality()), _120_v0, new BigNumber((_129_v9).length), _120_v0));
      let _131_v11;
      let _init4 = function (_132_i0) {
        return (new _dafny.CodePoint('j'.codePointAt(0)));
      };
      let _nw22 = Array((new BigNumber(23)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw22.length); _i0_4++) {
        _nw22[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _131_v11 = _nw22;
      let _133_v12;
      _133_v12 = true;
      let _134_v13;
      _134_v13 = _dafny.Map.Empty.slice().updateUnsafe(_133_v12,_133_v12);
      let _135_v14;
      let _nw23 = Array((new BigNumber(2)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("eqobbdls");
      _nw23[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("rgyeuu");
      _135_v14 = _nw23;
      let _136_v15;
      _136_v15 = _dafny.Set.fromElements(_123_v3);
      let _137_globalState;
      let _nw24 = new _module.GlobalState();
      _nw24.__ctor((_122_v2).Merge(_122_v2), _121_v1, new BigNumber(445), false, new BigNumber(988), new _dafny.CodePoint('l'.codePointAt(0)), _124_v4, new BigNumber(496), new BigNumber(-233), new BigNumber(938), (((_130_v10).contains(_127_v7)) ? ((_130_v10).get(_127_v7)) : (_121_v1)), new BigNumber(987), false, true, false, new BigNumber(99), false, _131_v11, new BigNumber(161), _134_v13, _135_v14, _dafny.Seq.UnicodeFromString("txxtamr"), new BigNumber(577), _136_v15, false, new BigNumber(181));
      _137_globalState = _nw24;
      let _138_v16;
      _138_v16 = new _dafny.CodePoint('j'.codePointAt(0));
      let _139_v17;
      _139_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_133_v12,_120_v0),_138_v16);
      _139_v17 = _139_v17;
      let _140_i1;
      _140_i1 = _dafny.ZERO;
      L1: {
        while (_133_v12) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_140_i1)) {
              break L1;
            }
            _140_i1 = (_140_i1).plus(_dafny.ONE);
            let _141_v18;
            _141_v18 = _dafny.Seq.of(_133_v12);
            (_137_globalState).f14 = (_141_v18)[_module.__default.safeIndex(_120_v0, new BigNumber((_141_v18).length))];
            let _142_v19;
            _142_v19 = _module.D0.create_DC1(_133_v12);
            let _143_v20;
            _143_v20 = _dafny.Map.Empty.slice().updateUnsafe(_142_v19,new BigNumber((_128_v8).cardinality()));
            let _144_v21;
            _144_v21 = _dafny.Map.Empty.slice().updateUnsafe(_120_v0,_133_v12);
            let _145_v22;
            _145_v22 = _dafny.Set.fromElements(_133_v12);
            let _146_v23;
            _146_v23 = _dafny.Map.Empty.slice().updateUnsafe((((_144_v21).contains(new BigNumber((_145_v22).length))) ? ((_144_v21).get(new BigNumber((_145_v22).length))) : (_133_v12)),_120_v0);
            _143_v20 = (_143_v20).update(_142_v19, new BigNumber((_146_v23).length));
            let _index12 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_125_v5).length));
            (_125_v5)[_index12] = new BigNumber((_129_v9).length);
            let _147_v24;
            _147_v24 = _dafny.Map.Empty.slice().updateUnsafe((_141_v18)[_module.__default.safeIndex(new BigNumber((_141_v18).length), new BigNumber((_141_v18).length))],_138_v16);
            let _148_v25;
            _148_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_147_v24).length),_125_v5);
            let _index13 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_125_v5).length));
            (_125_v5)[_index13] = (_120_v0).plus(new BigNumber(((_148_v25).Merge(_148_v25)).length));
            let _149_v26;
            _149_v26 = _dafny.Map.Empty.slice().updateUnsafe(_138_v16,_133_v12);
            let _150_v27;
            _150_v27 = _138_v16;
            _149_v26 = (_149_v26).update(_150_v27, _133_v12);
          }
        }
      }
      let _151_v28;
      let _152_v29;
      let _153_v30;
      let _154_v31;
      let _out0;
      let _out1;
      let _out2;
      let _out3;
      let _outcollector0 = _module.__default.m0(_module.__default.fm0(new BigNumber((_dafny.Seq.UnicodeFromString("gb")).length), _133_v12, _137_globalState), _137_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _out3 = _outcollector0[3];
      _151_v28 = _out0;
      _152_v29 = _out1;
      _153_v30 = _out2;
      _154_v31 = _out3;
      (_137_globalState).f22 = (_dafny.ZERO).minus(_152_v29);
      if (_133_v12) {
        (_137_globalState).f14 = _133_v12;
        let _155_v32;
        _155_v32 = _module.D0.create_DC1(_133_v12);
        let _source2 = _155_v32;
        if (_source2.is_DC1) {
          let _156___mcc_h0 = (_source2).cf1;
          let _157_cf1 = _156___mcc_h0;
          let _158_v33;
          let _nw25 = Array((new BigNumber(8)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = !(_152_v29).isEqualTo(_152_v29);
          _nw25[(_dafny.ONE).toNumber()] = true;
          _nw25[(new BigNumber(2)).toNumber()] = _133_v12;
          _nw25[(new BigNumber(3)).toNumber()] = _133_v12;
          _nw25[(new BigNumber(4)).toNumber()] = (_121_v1).IsDisjointFrom(_121_v1);
          _nw25[(new BigNumber(5)).toNumber()] = _157_cf1;
          _nw25[(new BigNumber(6)).toNumber()] = (new BigNumber(492)).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_120_v0)));
          _nw25[(new BigNumber(7)).toNumber()] = _157_cf1;
          _158_v33 = _nw25;
          _158_v33 = _158_v33;
          (_137_globalState).f14 = (!(_133_v12)) || (_157_cf1);
          (_137_globalState).f22 = _152_v29;
          let _index14 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_125_v5).length));
          (_125_v5)[_index14] = _152_v29;
          let _index15 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_125_v5).length));
          (_125_v5)[_index15] = new BigNumber(882);
        } else if (_source2.is_DC2) {
          let _159___mcc_h1 = (_source2).cf2;
          let _160___mcc_h2 = (_source2).cf3;
          let _161_cf3 = _160___mcc_h2;
          let _162_cf2 = _159___mcc_h1;
          let _163_v34;
          _163_v34 = _dafny.MultiSet.fromElements(_133_v12, !(_133_v12));
          let _164_v36;
          _164_v36 = _dafny.Map.Empty.slice().updateUnsafe(_162_cf2,_120_v0);
          let _165_v37;
          _165_v37 = _dafny.Map.Empty.slice().updateUnsafe(_133_v12,_module.__default.fm1(new BigNumber(716), (_151_v28)[_module.__default.safeIndex(_162_cf2, new BigNumber((_151_v28).length))], _129_v9, _133_v12, _137_globalState));
          let _166_v40;
          let _nw26 = Array((new BigNumber(19)).toNumber());
          _nw26[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_163_v34).cardinality()),_161_cf3)).Merge(function () {
            let _coll3 = new _dafny.Map();
            for (const _compr_3 of _dafny.IntegerRange(new BigNumber(717), new BigNumber(364))) {
              let _167_v35 = _compr_3;
              if (((new BigNumber(717)).isLessThanOrEqualTo(_167_v35)) && ((_167_v35).isLessThan(new BigNumber(364)))) {
                _coll3.push([(_167_v35).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_133_v12,_162_cf2),true)).length)),new BigNumber((_dafny.MultiSet.fromElements(_162_cf2, (_dafny.ZERO).minus(_120_v0))).cardinality())]);
              }
            }
            return _coll3;
          }());
          _nw26[(_dafny.ONE).toNumber()] = _164_v36;
          _nw26[(new BigNumber(2)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(3)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_162_cf2,_152_v29)).Merge(_164_v36);
          _nw26[(new BigNumber(5)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(6)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(7)).toNumber()] = (_164_v36).update(new BigNumber((_165_v37).length), _161_cf3);
          _nw26[(new BigNumber(8)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(9)).toNumber()] = ((_133_v12) ? (_164_v36) : ((_module.D2.create_DC4(_164_v36)).dtor_cf5));
          _nw26[(new BigNumber(10)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(11)).toNumber()] = function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-473), new BigNumber(31))) {
              let _168_v38 = _compr_4;
              if (((new BigNumber(-473)).isLessThanOrEqualTo(_168_v38)) && ((_168_v38).isLessThan(new BigNumber(31)))) {
                _coll4.push([(_168_v38).plus((_dafny.ZERO).minus(_152_v29)),(_dafny.ZERO).minus(_162_cf2)]);
              }
            }
            return _coll4;
          }();
          _nw26[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_152_v29,_162_cf2)).Merge(_164_v36);
          _nw26[(new BigNumber(13)).toNumber()] = (_164_v36).Merge(_164_v36);
          _nw26[(new BigNumber(14)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(15)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(16)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(17)).toNumber()] = _164_v36;
          _nw26[(new BigNumber(18)).toNumber()] = (function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of ((_164_v36).update(new BigNumber((_151_v28).length), _161_cf3)).Keys.Elements) {
              let _169_v39 = _compr_5;
              if (((_164_v36).update(new BigNumber((_151_v28).length), _161_cf3)).contains(_169_v39)) {
                _coll5.push([_module.__default.safeModuloInt(_169_v39, new BigNumber(990)),new BigNumber((_151_v28).length)]);
              }
            }
            return _coll5;
          }()).update(_162_cf2, (_dafny.ZERO).minus(new BigNumber(-994)));
          _166_v40 = _nw26;
          let _170_v41;
          _170_v41 = _dafny.Map.Empty.slice().updateUnsafe(_152_v29,_155_v32);
          let _index16 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_153_v30).length));
          (_153_v30)[_index16] = _module.__default.fm1(new BigNumber((_170_v41).length), _133_v12, _dafny.Seq.UnicodeFromString("vjvcgp"), _133_v12, _137_globalState);
          let _index17 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_153_v30).length));
          let _rhs14 = _162_cf2;
          let _rhs15 = _162_cf2;
          let _rhs16 = _166_v40;
          let _rhs17 = !(_133_v12);
          let _rhs18 = _120_v0;
          let _lhs11 = _137_globalState;
          let _lhs12 = _137_globalState;
          let _lhs13 = _153_v30;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_153_v30).length));
          _152_v29 = _rhs14;
          _lhs11.f22 = _rhs15;
          _166_v40 = _rhs16;
          _lhs12.f14 = _rhs17;
          _lhs13[_lhs14] = _rhs18;
          let _171_v42;
          _171_v42 = _dafny.Set.fromElements((_164_v36).Merge((_164_v36).update(new BigNumber(12), (_153_v30)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_153_v30).length))])));
          let _172_v43;
          _172_v43 = _dafny.Map.Empty.slice().updateUnsafe(_164_v36,_133_v12);
          let _pat_let_tv0 = _171_v42;
          let _pat_let_tv1 = _171_v42;
          _171_v42 = ((function (_pat_let0_0) {
            return function (_175_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_176_dt__update_hcf12_h1) {
                  return _module.D3.create_DC8(_176_dt__update_hcf12_h1);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let0_0);
          }(function (_pat_let1_0) {
            return function (_173_dt__update__tmp_h0) {
              return function (_pat_let2_0) {
                return function (_174_dt__update_hcf12_h0) {
                  return _module.D3.create_DC8(_174_dt__update_hcf12_h0);
                }(_pat_let2_0);
              }(_pat_let_tv0);
            }(_pat_let1_0);
          }(_module.D3.create_DC8(_dafny.Set.fromElements(_164_v36))))).dtor_cf12).Intersect((_171_v42).Intersect(function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_172_v43).Keys.Elements) {
              let _177_v44 = _compr_6;
              if ((_172_v43).contains(_177_v44)) {
                _coll6.add(_177_v44);
              }
            }
            return _coll6;
          }()));
          (_137_globalState).f22 = _module.__default.fm1(_module.__default.safeModuloInt(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), (_dafny.ZERO).minus((_dafny.ZERO).minus(_152_v29))), _133_v12, _module.__default.fm2(_152_v29, _120_v0, (_151_v28)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(933)), ((_178_v29) => function (_179_i2) {
            return _178_v29;
          })(_152_v29))).length), new BigNumber((_151_v28).length))], (_dafny.ZERO).minus(_161_cf3), _137_globalState), _133_v12, _137_globalState);
          let _180_v45;
          _180_v45 = _dafny.Set.fromElements(true, _133_v12);
          _180_v45 = ((_180_v45).Difference(_module.__default.fm3((_153_v30)[_module.__default.safeIndex(new BigNumber(329), new BigNumber((_153_v30).length))], _133_v12, _137_globalState))).Union((_dafny.Set.fromElements(_133_v12, _133_v12)).Difference(_180_v45));
        } else {
          let _181___mcc_h3 = (_source2).cf0;
          let _182_cf0 = _181___mcc_h3;
          (_137_globalState).f25 = new BigNumber((_121_v1).length);
          (_137_globalState).f9 = _120_v0;
          let _183_v46;
          let _184_v47;
          let _185_v48;
          let _186_v49;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = _module.__default.m0(!(_133_v12), _137_globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _183_v46 = _out4;
          _184_v47 = _out5;
          _185_v48 = _out6;
          _186_v49 = _out7;
          let _187_v50;
          _187_v50 = _dafny.MultiSet.fromElements(_133_v12, _133_v12, _133_v12, _133_v12, _133_v12);
          let _188_v51;
          _188_v51 = _138_v16;
          let _rhs19 = (new BigNumber((_187_v50).cardinality())).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_152_v29, new BigNumber((_123_v3).length)))));
          let _rhs20 = _module.__default.fm4(_188_v51, _dafny.Seq.of(true), _137_globalState);
          _184_v47 = _rhs19;
          _186_v49 = _rhs20;
        }
        let _189_v52;
        _189_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_120_v0),_120_v0);
        let _190_v54;
        _190_v54 = _dafny.Seq.of(_123_v3, _dafny.Seq.update(_123_v3, _module.__default.safeIndex(_120_v0, new BigNumber((_123_v3).length)), _120_v0));
        _189_v52 = (_189_v52).Merge((function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_190_v54).Elements) {
            let _191_v53 = _compr_7;
            if (_dafny.Seq.contains(_190_v54, _191_v53)) {
              _coll7.push([_191_v53,_120_v0]);
            }
          }
          return _coll7;
        }()).Merge(_189_v52));
        let _192_v55;
        _192_v55 = _dafny.MultiSet.fromElements(_133_v12, _133_v12, true, _133_v12);
        let _193_v56;
        _193_v56 = _module.D0.create_DC2(new BigNumber(623), _module.__default.fm1(_152_v29, _133_v12, _129_v9, _133_v12, _137_globalState));
        let _source3 = _module.D2.create_DC5((_121_v1).IsProperSubsetOf(_121_v1), ((((_192_v55).contains(false)) ? ((_192_v55).get(false)) : ((_193_v56).dtor_cf2))).isLessThan(_152_v29), true, _152_v29, (_module.__default.fm1(_152_v29, false, _module.__default.fm2(_152_v29, _152_v29, _133_v12, _120_v0, _137_globalState), _133_v12, _137_globalState)).isLessThan(_120_v0));
        if (_source3.is_DC5) {
          let _194___mcc_h4 = (_source3).cf6;
          let _195___mcc_h5 = (_source3).cf7;
          let _196___mcc_h6 = (_source3).cf8;
          let _197___mcc_h7 = (_source3).cf9;
          let _198___mcc_h8 = (_source3).cf10;
          let _199_cf10 = _198___mcc_h8;
          let _200_cf9 = _197___mcc_h7;
          let _201_cf8 = _196___mcc_h6;
          let _202_cf7 = _195___mcc_h5;
          let _203_cf6 = _194___mcc_h4;
          let _204_v57;
          _204_v57 = _dafny.Set.fromElements(_121_v1);
          _202_cf7 = ((_204_v57).Intersect(_204_v57)).IsSubsetOf(_204_v57);
          let _rhs21 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_200_cf9).multipliedBy(_152_v29)));
          let _rhs22 = ((_133_v12) ? (_125_v5) : (((_201_cf8) ? (_125_v5) : (_125_v5))));
          let _rhs23 = _129_v9;
          let _rhs24 = _dafny.Seq.Concat(_module.__default.fm2(_200_cf9, _120_v0, _202_cf7, new BigNumber((_151_v28).length), _137_globalState), _129_v9);
          let _lhs15 = _137_globalState;
          let _lhs16 = _137_globalState;
          _152_v29 = _rhs21;
          _125_v5 = _rhs22;
          _lhs15.f21 = _rhs23;
          _lhs16.f21 = _rhs24;
          let _index18 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_125_v5).length));
          (_125_v5)[_index18] = _120_v0;
          let _index19 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_125_v5).length));
          (_125_v5)[_index19] = _152_v29;
          let _205_v58;
          let _206_v59;
          let _207_v60;
          let _208_v61;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = _module.__default.m0(!(_133_v12), _137_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _205_v58 = _out8;
          _206_v59 = _out9;
          _207_v60 = _out10;
          _208_v61 = _out11;
        } else if (_source3.is_DC6) {
          let _209_v62;
          let _nw27 = new _module.C0();
          _nw27.__ctor();
          _209_v62 = _nw27;
          let _210_v63;
          _210_v63 = _dafny.MultiSet.fromElements(_131_v11);
          let _211_v64;
          _211_v64 = _dafny.Map.Empty.slice().updateUnsafe(_210_v63,_209_v62);
          let _212_v65;
          _212_v65 = _dafny.Seq.of(_209_v62);
          let _213_v66;
          _213_v66 = _dafny.Map.Empty.slice().updateUnsafe(_133_v12,_209_v62);
          let _214_v67;
          let _nw28 = Array((new BigNumber(26)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _209_v62;
          _nw28[(_dafny.ONE).toNumber()] = _209_v62;
          _nw28[(new BigNumber(2)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(3)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(4)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(5)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(6)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(7)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(8)).toNumber()] = ((_module.__default.fm0(_120_v0, _133_v12, _137_globalState)) ? (_209_v62) : (_209_v62));
          _nw28[(new BigNumber(9)).toNumber()] = (((_211_v64).contains(_210_v63)) ? ((_211_v64).get(_210_v63)) : ((_212_v65)[_module.__default.safeIndex(_120_v0, new BigNumber((_212_v65).length))]));
          _nw28[(new BigNumber(10)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(11)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(12)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(13)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(14)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(15)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(16)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(17)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(18)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(19)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(20)).toNumber()] = (((_213_v66).contains(_module.__default.fm0(_120_v0, _133_v12, _137_globalState))) ? ((_213_v66).get(_module.__default.fm0(_120_v0, _133_v12, _137_globalState))) : (_209_v62));
          _nw28[(new BigNumber(21)).toNumber()] = (((_213_v66).contains(_133_v12)) ? ((_213_v66).get(_133_v12)) : (_209_v62));
          _nw28[(new BigNumber(22)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(23)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(24)).toNumber()] = _209_v62;
          _nw28[(new BigNumber(25)).toNumber()] = _209_v62;
          _214_v67 = _nw28;
          let _index20 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_214_v67).length));
          (_214_v67)[_index20] = _209_v62;
          let _index21 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_214_v67).length));
          (_214_v67)[_index21] = ((_133_v12) ? (_209_v62) : (_209_v62));
          let _215_v68;
          let _nw29 = new _module.C0();
          _nw29.__ctor();
          _215_v68 = _nw29;
          let _216_v69;
          let _217_v70;
          let _218_v71;
          let _219_v72;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = _module.__default.m0(_133_v12, _137_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _216_v69 = _out12;
          _217_v70 = _out13;
          _218_v71 = _out14;
          _219_v72 = _out15;
        } else if (_source3.is_DC4) {
          let _220___mcc_h9 = (_source3).cf5;
          let _221_cf5 = _220___mcc_h9;
          let _222_v73;
          let _nw30 = new _module.C0();
          _nw30.__ctor();
          _222_v73 = _nw30;
          let _223_v74;
          _223_v74 = _dafny.Map.Empty.slice().updateUnsafe(_222_v73,_222_v73);
          let _224_v75;
          _224_v75 = _module.D4.create_DC10(_222_v73);
          let _rhs25 = (((_223_v74).contains((_224_v75).dtor_cf15)) ? ((_223_v74).get((_224_v75).dtor_cf15)) : (_222_v73));
          let _rhs26 = (((_192_v55).contains((((_134_v13).contains(_133_v12)) ? ((_134_v13).get(_133_v12)) : (_133_v12)))) ? ((_192_v55).get((((_134_v13).contains(_133_v12)) ? ((_134_v13).get(_133_v12)) : (_133_v12)))) : (_module.__default.safeDivisionInt(_120_v0, _152_v29)));
          let _lhs17 = _137_globalState;
          _222_v73 = _rhs25;
          _lhs17.f22 = _rhs26;
          let _225_v76;
          let _nw31 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _225_v76 = _nw31;
          let _rhs27 = _152_v29;
          let _rhs28 = (new BigNumber(-295)).isLessThanOrEqualTo((_123_v3)[_module.__default.safeIndex(_120_v0, new BigNumber((_123_v3).length))]);
          let _rhs29 = _154_v31;
          let _rhs30 = !(_133_v12);
          let _rhs31 = _225_v76;
          let _lhs18 = _137_globalState;
          let _lhs19 = _137_globalState;
          let _lhs20 = _137_globalState;
          _lhs18.f25 = _rhs27;
          _lhs19.f14 = _rhs28;
          _154_v31 = _rhs29;
          _lhs20.f14 = _rhs30;
          _225_v76 = _rhs31;
          (_137_globalState).f21 = _129_v9;
          let _226_v77;
          let _227_v78;
          let _228_v79;
          let _229_v80;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = _module.__default.m0((_152_v29).isLessThan(new BigNumber(845)), _137_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _226_v77 = _out16;
          _227_v78 = _out17;
          _228_v79 = _out18;
          _229_v80 = _out19;
        } else {
          let _230___mcc_h10 = (_source3).cf11;
          let _231_cf11 = _230___mcc_h10;
          let _232_v81;
          let _nw32 = Array((new BigNumber(11)).toNumber()).fill(false);
          _232_v81 = _nw32;
          let _index22 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_232_v81).length));
          (_232_v81)[_index22] = (_155_v32).dtor_cf1;
          let _233_v82;
          _233_v82 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_module.__default.fm6(_120_v0, _137_globalState)));
          let _index23 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_232_v81).length));
          (_232_v81)[_index23] = (((_233_v82)[_module.__default.safeIndex(_152_v29, new BigNumber((_233_v82).length))]).Intersect((_192_v55).update(_133_v12, _module.__default.abs((_dafny.ZERO).minus(_120_v0))))).IsSubsetOf(((_133_v12) ? (_192_v55) : (_dafny.MultiSet.fromElements(_133_v12, _133_v12))));
          (_137_globalState).f24 = ((_133_v12) ? ((_dafny.MultiSet.fromElements(_152_v29)).IsProperSubsetOf(_128_v8)) : (_133_v12));
          let _234_v83;
          _234_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_123_v3).length),_152_v29);
          let _235_v84;
          _235_v84 = _dafny.Set.fromElements(_133_v12);
          let _236_v85;
          _236_v85 = _dafny.Set.fromElements(_234_v83, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_129_v9).length),_152_v29), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_235_v84).length),_152_v29), _234_v83, _234_v83);
          let _237_v86;
          _237_v86 = _module.D3.create_DC8(_236_v85);
          _237_v86 = _237_v86;
          (_137_globalState).f14 = (_232_v81)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_232_v81).length))];
        }
        let _238_v87;
        let _239_v88;
        let _240_v89;
        let _241_v90;
        let _out20;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector5 = _module.__default.m0(_133_v12, _137_globalState);
        _out20 = _outcollector5[0];
        _out21 = _outcollector5[1];
        _out22 = _outcollector5[2];
        _out23 = _outcollector5[3];
        _238_v87 = _out20;
        _239_v88 = _out21;
        _240_v89 = _out22;
        _241_v90 = _out23;
      } else {
        let _source4 = _154_v31;
        let _242___mcc_h11 = _source4;
        let _243_cf4 = _242___mcc_h11;
        let _244_v91;
        _244_v91 = _dafny.Set.fromElements(_133_v12);
        _244_v91 = _244_v91;
        let _245_v92;
        let _nw33 = Array((new BigNumber(27)).toNumber()).fill(false);
        _245_v92 = _nw33;
        let _index24 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_245_v92).length));
        (_245_v92)[_index24] = _133_v12;
        let _246_v93;
        let _nw34 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.of());
        _246_v93 = _nw34;
        let _247_v94;
        _247_v94 = _dafny.Seq.of(_125_v5, _153_v30);
        let _index25 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_246_v93).length));
        (_246_v93)[_index25] = _247_v94;
        let _248_v95;
        _248_v95 = _dafny.Map.Empty.slice().updateUnsafe(!(_133_v12),new BigNumber((_151_v28).length));
        let _249_v96;
        _249_v96 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_248_v95).length),_133_v12),_153_v30);
        let _250_v97;
        _250_v97 = _dafny.Map.Empty.slice().updateUnsafe(_152_v29,_133_v12);
        let _index26 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_245_v92).length));
        let _index27 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_246_v93).length));
        let _rhs32 = (((_249_v96).contains(_250_v97)) ? ((_249_v96).get(_250_v97)) : (((_133_v12) ? (_125_v5) : (_153_v30))));
        let _rhs33 = _133_v12;
        let _rhs34 = !(_133_v12);
        let _rhs35 = (_123_v3)[_module.__default.safeIndex(_152_v29, new BigNumber((_123_v3).length))];
        let _rhs36 = _dafny.Seq.Concat(_247_v94, _247_v94);
        let _lhs21 = _137_globalState;
        let _lhs22 = _245_v92;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(263), new BigNumber((_245_v92).length));
        let _lhs24 = _137_globalState;
        let _lhs25 = _246_v93;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_246_v93).length));
        _153_v30 = _rhs32;
        _lhs21.f14 = _rhs33;
        _lhs22[_lhs23] = _rhs34;
        _lhs24.f22 = _rhs35;
        _lhs25[_lhs26] = _rhs36;
        let _251_v98;
        let _nw35 = new _module.C0();
        _nw35.__ctor();
        _251_v98 = _nw35;
        let _252_v99;
        let _253_v100;
        let _254_v101;
        let _255_v102;
        let _out24;
        let _out25;
        let _out26;
        let _out27;
        let _outcollector6 = _module.__default.m0(_module.__default.fm0(_120_v0, (_245_v92)[_module.__default.safeIndex(new BigNumber(263), new BigNumber((_245_v92).length))], _137_globalState), _137_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _out26 = _outcollector6[2];
        _out27 = _outcollector6[3];
        _252_v99 = _out24;
        _253_v100 = _out25;
        _254_v101 = _out26;
        _255_v102 = _out27;
        _129_v9 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aml"), _129_v9), _dafny.Seq.Concat(_129_v9, _129_v9));
        (_137_globalState).f14 = ((!(_152_v29).isEqualTo(_120_v0)) ? (_133_v12) : (((_133_v12) ? (_133_v12) : (_133_v12))));
        (_137_globalState).f24 = _133_v12;
        let _256_v103;
        _256_v103 = _dafny.Map.Empty.slice().updateUnsafe(_152_v29,_129_v9);
        _256_v103 = (_256_v103).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_120_v0,!(false))).length), _129_v9);
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_125_v5).length))) {
        let _257_i3 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_257_i3)) && ((_257_i3).isLessThan(new BigNumber((_125_v5).length))))) {
          (_125_v5)[(_257_i3)] = (_257_i3).plus(new BigNumber(110));
        }
      }
      let _258_v104;
      _258_v104 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_152_v29),_154_v31);
      (_137_globalState).f24 = _dafny.Seq.contains(_dafny.Seq.update(_129_v9, _module.__default.safeIndex((_dafny.ZERO).minus(_120_v0), new BigNumber((_129_v9).length)), _154_v31), (((_258_v104).contains(_120_v0)) ? ((_258_v104).get(_120_v0)) : (new _dafny.CodePoint('g'.codePointAt(0)))));
      let _259_i4;
      _259_i4 = _dafny.ZERO;
      L2: {
        while (_133_v12) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_259_i4)) {
              break L2;
            }
            _259_i4 = (_259_i4).plus(_dafny.ONE);
            let _260_v105;
            _260_v105 = _dafny.Seq.of(_125_v5, _125_v5, _125_v5);
            let _index28 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length));
            (_153_v30)[_index28] = (new BigNumber((_dafny.Seq.of((_151_v28)[_module.__default.safeIndex(_152_v29, new BigNumber((_151_v28).length))])).length)).minus(new BigNumber((_260_v105).length));
            let _261_v106;
            _261_v106 = _dafny.Seq.of(_134_v13);
            let _262_v107;
            _262_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_261_v106).length),_133_v12);
            let _263_v108;
            _263_v108 = _dafny.Map.Empty.slice().updateUnsafe(_262_v107,_120_v0);
            let _264_v111;
            _264_v111 = _dafny.Seq.of(_262_v107, _262_v107, function () {
              let _coll8 = new _dafny.Map();
              for (const _compr_8 of (_dafny.MultiSet.fromElements(_152_v29, _120_v0)).Elements) {
                let _265_v110 = _compr_8;
                if ((_dafny.MultiSet.fromElements(_152_v29, _120_v0)).contains(_265_v110)) {
                  _coll8.push([(_265_v110).plus(_120_v0),_133_v12]);
                }
              }
              return _coll8;
            }());
            let _index29 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length));
            let _rhs37 = new BigNumber((((_263_v108).Merge(function () {
              let _coll9 = new _dafny.Map();
              for (const _compr_9 of (_264_v111).Elements) {
                let _266_v109 = _compr_9;
                if (_dafny.Seq.contains(_264_v111, _266_v109)) {
                  _coll9.push([_266_v109,_152_v29]);
                }
              }
              return _coll9;
            }())).Merge(_263_v108)).length);
            let _rhs38 = _152_v29;
            let _lhs27 = _137_globalState;
            let _lhs28 = _153_v30;
            let _lhs29 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length));
            _lhs27.f9 = _rhs37;
            _lhs28[_lhs29] = _rhs38;
            let _267_v112;
            _267_v112 = _dafny.Set.fromElements(_133_v12, _133_v12);
            let _268_v113;
            _268_v113 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_152_v29,_138_v16),true);
            (_137_globalState).f9 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.fm1(new BigNumber((_267_v112).length), _133_v12, _129_v9, _133_v12, _137_globalState)), new BigNumber((_268_v113).length)),(_153_v30)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length))])).length);
            let _index30 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length));
            (_153_v30)[_index30] = _module.__default.safeModuloInt(_152_v29, new BigNumber((_262_v107).length));
            (_137_globalState).f22 = (_dafny.ZERO).minus((_153_v30)[_module.__default.safeIndex(new BigNumber(236), new BigNumber((_153_v30).length))]);
          }
        }
      }
      (_137_globalState).f25 = _152_v29;
      let _hi1 = new BigNumber(63);
      for (let _269_i5 = new BigNumber((_129_v9).length); _269_i5.isLessThan(_hi1); _269_i5 = _269_i5.plus(_dafny.ONE)) {
        let _270_v114;
        _270_v114 = _dafny.Map.Empty.slice().updateUnsafe(_269_i5,_152_v29);
        let _271_v115;
        _271_v115 = _dafny.Set.fromElements(_270_v114);
        let _272_v116;
        _272_v116 = _module.D3.create_DC8(_271_v115);
        let _pat_let_tv2 = _270_v114;
        let _pat_let_tv3 = _270_v114;
        let _pat_let_tv4 = _271_v115;
        let _source5 = function (_pat_let4_0) {
          return function (_273_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_274_dt__update_hcf12_h2) {
                return _module.D3.create_DC8(_274_dt__update_hcf12_h2);
              }(_pat_let5_0);
            }((_dafny.Set.fromElements(_pat_let_tv2, _pat_let_tv3)).Union(_pat_let_tv4));
          }(_pat_let4_0);
        }(_272_v116);
        if (_source5.is_DC9) {
          let _275___mcc_h12 = (_source5).cf13;
          let _276___mcc_h13 = (_source5).cf14;
          let _277_cf14 = _276___mcc_h13;
          let _278_cf13 = _275___mcc_h12;
          let _279_v117;
          _279_v117 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_151_v28, _151_v28),_152_v29);
          _279_v117 = (_279_v117).update(_151_v28, (_dafny.ZERO).minus(_269_i5));
          let _280_v118;
          let _nw36 = Array((new BigNumber(16)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _278_cf13;
          _nw36[(_dafny.ONE).toNumber()] = _133_v12;
          _nw36[(new BigNumber(2)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(3)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(4)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(5)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(6)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(7)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(8)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(9)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(10)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(11)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(12)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(13)).toNumber()] = _133_v12;
          _nw36[(new BigNumber(14)).toNumber()] = _278_cf13;
          _nw36[(new BigNumber(15)).toNumber()] = _278_cf13;
          _280_v118 = _nw36;
          let _281_v119;
          _281_v119 = _dafny.Map.Empty.slice().updateUnsafe(_269_i5,_280_v118);
          (_137_globalState).f3 = !(((!(_278_cf13)) ? (_281_v119) : (_281_v119))).equals(_281_v119);
          let _index31 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_153_v30).length));
          (_153_v30)[_index31] = _module.__default.safeModuloInt(_module.__default.fm1(_152_v29, _133_v12, _129_v9, _278_cf13, _137_globalState), (_dafny.ZERO).minus(_152_v29));
          let _index32 = _module.__default.safeIndex(new BigNumber(920), new BigNumber((_153_v30).length));
          (_153_v30)[_index32] = _120_v0;
          _151_v28 = _module.__default.fm6(new BigNumber((_129_v9).length), _137_globalState);
        } else {
          let _282___mcc_h14 = (_source5).cf12;
          let _283_cf12 = _282___mcc_h14;
          let _284_v120;
          _284_v120 = _dafny.Map.Empty.slice().updateUnsafe(_133_v12,new BigNumber(875));
          _284_v120 = (_284_v120).update(true, ((_133_v12) ? (new BigNumber((_dafny.Seq.of(_120_v0)).length)) : (_152_v29)));
          (_137_globalState).f3 = _133_v12;
          (_137_globalState).f9 = _152_v29;
          let _285_v121;
          let _nw37 = new _module.C0();
          _nw37.__ctor();
          _285_v121 = _nw37;
        }
        if (false) {
          let _286_v122;
          let _287_v123;
          let _288_v124;
          let _289_v125;
          let _out28;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector7 = _module.__default.m0(!(_133_v12), _137_globalState);
          _out28 = _outcollector7[0];
          _out29 = _outcollector7[1];
          _out30 = _outcollector7[2];
          _out31 = _outcollector7[3];
          _286_v122 = _out28;
          _287_v123 = _out29;
          _288_v124 = _out30;
          _289_v125 = _out31;
          let _290_v126;
          _290_v126 = _module.D2.create_DC4(_270_v114);
          let _291_v127;
          _291_v127 = _module.D3.create_DC9(_133_v12, (_290_v126).dtor_cf5);
          let _292_v128;
          _292_v128 = _dafny.Seq.of((_270_v114).update(_120_v0, new BigNumber((_129_v9).length)));
          let _293_v129;
          _293_v129 = _dafny.Map.Empty.slice().updateUnsafe(_120_v0,_module.D3.create_DC9((_291_v127).dtor_cf13, (_292_v128)[_module.__default.safeIndex(_module.__default.fm1(_120_v0, _133_v12, _dafny.Seq.Create(_module.__default.abs(new BigNumber(410)), ((_294_v125) => function (_295_i6) {
  return _294_v125;
})(_289_v125)), _133_v12, _137_globalState), new BigNumber((_292_v128).length))]));
          let _296_v130;
          _296_v130 = _dafny.Map.Empty.slice().updateUnsafe((_293_v129).Merge(_293_v129),_154_v31);
          _296_v130 = (_296_v130).update(_293_v129, _138_v16);
          let _297_v131;
          let _nw38 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _297_v131 = _nw38;
          let _298_v132;
          _298_v132 = _dafny.Seq.of(_123_v3, _123_v3, _dafny.Seq.of(_152_v29, (_dafny.ZERO).minus(_269_i5)));
          let _index33 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_297_v131).length));
          (_297_v131)[_index33] = _dafny.MultiSet.FromArray(_298_v132);
          let _index34 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_297_v131).length));
          (_297_v131)[_index34] = _dafny.MultiSet.fromElements(_dafny.Seq.of(_269_i5, new BigNumber(683), new BigNumber((_123_v3).length)));
          (_137_globalState).f22 = (_module.__default.fm1(_152_v29, _133_v12, _dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_299_v125) => function (_300_i7) {
            return _299_v125;
          })(_289_v125)), _module.__default.fm0(_269_i5, _133_v12, _137_globalState), _137_globalState)).plus(new BigNumber(228));
          let _301_v133;
          let _nw39 = new _module.C0();
          _nw39.__ctor();
          _301_v133 = _nw39;
        } else {
          let _302_v134;
          let _nw40 = Array((new BigNumber(29)).toNumber()).fill([]);
          _302_v134 = _nw40;
          let _303_v135;
          let _init5 = function (_304_i8) {
            return true;
          };
          let _nw41 = Array((new BigNumber(19)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw41.length); _i0_5++) {
            _nw41[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _303_v135 = _nw41;
          let _305_v136;
          let _nw42 = Array((new BigNumber(12)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _303_v135;
          _nw42[(_dafny.ONE).toNumber()] = _303_v135;
          _nw42[(new BigNumber(2)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(3)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(4)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(5)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(6)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(7)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(8)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(9)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(10)).toNumber()] = _303_v135;
          _nw42[(new BigNumber(11)).toNumber()] = _303_v135;
          _305_v136 = _nw42;
          let _index35 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_302_v134).length));
          (_302_v134)[_index35] = _305_v136;
          let _306_v137;
          _306_v137 = _dafny.Map.Empty.slice().updateUnsafe(_125_v5,_125_v5);
          let _index36 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_302_v134).length));
          let _rhs39 = ((!(true)) ? ((_133_v12) && (_module.__default.fm0(new BigNumber((_306_v137).length), _133_v12, _137_globalState))) : (_133_v12));
          let _rhs40 = _305_v136;
          let _rhs41 = _120_v0;
          let _lhs30 = _302_v134;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_302_v134).length));
          let _lhs32 = _137_globalState;
          _133_v12 = _rhs39;
          _lhs30[_lhs31] = _rhs40;
          _lhs32.f25 = _rhs41;
          let _307_v138;
          let _nw43 = new _module.C0();
          _nw43.__ctor();
          _307_v138 = _nw43;
          let _index37 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_125_v5).length));
          (_125_v5)[_index37] = new BigNumber(964);
          let _308_v139;
          _308_v139 = _dafny.Map.Empty.slice().updateUnsafe(_307_v138,_307_v138);
          let _index38 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_125_v5).length));
          let _rhs42 = _module.__default.fm0(_152_v29, _133_v12, _137_globalState);
          let _rhs43 = (((_308_v139).contains(_307_v138)) ? ((_308_v139).get(_307_v138)) : (_307_v138));
          let _rhs44 = _152_v29;
          let _rhs45 = (_123_v3)[_module.__default.safeIndex(_120_v0, new BigNumber((_123_v3).length))];
          let _rhs46 = (_dafny.ZERO).minus(_120_v0);
          let _lhs33 = _137_globalState;
          let _lhs34 = _137_globalState;
          let _lhs35 = _137_globalState;
          let _lhs36 = _125_v5;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_125_v5).length));
          _lhs33.f24 = _rhs42;
          _307_v138 = _rhs43;
          _lhs34.f22 = _rhs44;
          _lhs35.f25 = _rhs45;
          _lhs36[_lhs37] = _rhs46;
          let _309_v140;
          _309_v140 = _module.D4.create_DC10(_307_v138);
          let _310_v141;
          _310_v141 = _dafny.Map.Empty.slice().updateUnsafe(_309_v140,(_125_v5)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_125_v5).length))]);
          _310_v141 = (_310_v141).update(_309_v140, (_125_v5)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_125_v5).length))]);
          let _311_v142;
          let _nw44 = new _module.C0();
          _nw44.__ctor();
          _311_v142 = _nw44;
        }
        let _312_v143;
        let _nw45 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
        _312_v143 = _nw45;
        let _313_v144;
        _313_v144 = _dafny.Map.Empty.slice().updateUnsafe(!(_133_v12),_152_v29);
        let _index39 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_312_v143).length));
        (_312_v143)[_index39] = _313_v144;
        let _314_v145;
        _314_v145 = _module.D5.create_DC12(_313_v144);
        let _pat_let_tv5 = _313_v144;
        let _pat_let_tv6 = _133_v12;
        let _pat_let_tv7 = _313_v144;
        let _index40 = _module.__default.safeIndex(new BigNumber(937), new BigNumber((_312_v143).length));
        (_312_v143)[_index40] = (function (_pat_let6_0) {
          return function (_319_dt__update__tmp_h5) {
            return function (_pat_let11_0) {
              return function (_320_dt__update_hcf18_h2) {
                return _module.D5.create_DC12(_320_dt__update_hcf18_h2);
              }(_pat_let11_0);
            }(_pat_let_tv7);
          }(_pat_let6_0);
        }(function (_pat_let7_0) {
          return function (_317_dt__update__tmp_h4) {
            return function (_pat_let10_0) {
              return function (_318_dt__update_hcf18_h1) {
                return _module.D5.create_DC12(_318_dt__update_hcf18_h1);
              }(_pat_let10_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv6,_269_i5));
          }(_pat_let7_0);
        }(function (_pat_let8_0) {
          return function (_315_dt__update__tmp_h3) {
            return function (_pat_let9_0) {
              return function (_316_dt__update_hcf18_h0) {
                return _module.D5.create_DC12(_316_dt__update_hcf18_h0);
              }(_pat_let9_0);
            }(_pat_let_tv5);
          }(_pat_let8_0);
        }(_314_v145)))).dtor_cf18;
        let _321_v146;
        _321_v146 = _dafny.MultiSet.fromElements(_133_v12, _133_v12);
        let _322_v147;
        _322_v147 = _dafny.Seq.of(_dafny.MultiSet.fromElements(true), _321_v146, _321_v146, _321_v146);
        (_137_globalState).f14 = (((_128_v8).IsSubsetOf(_128_v8)) ? (_133_v12) : (((_322_v147)[_module.__default.safeIndex(_152_v29, new BigNumber((_322_v147).length))]).IsDisjointFrom(_321_v146)));
      }
      let _323_v148;
      let _nw46 = new _module.C0();
      _nw46.__ctor();
      _323_v148 = _nw46;
      let _324_v149;
      _324_v149 = _dafny.Seq.of(_323_v148, _323_v148);
      let _325_v150;
      _325_v150 = _dafny.Map.Empty.slice().updateUnsafe(_120_v0,_129_v9);
      _323_v148 = (_324_v149)[_module.__default.safeIndex(_module.__default.fm1(_152_v29, _133_v12, (((_325_v150).contains(_120_v0)) ? ((_325_v150).get(_120_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-209)), ((_326_v16) => function (_327_i9) {
        return _326_v16;
      })(_138_v16)))), _133_v12, _137_globalState), new BigNumber((_324_v149).length))];
      let _index41 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length));
      (_125_v5)[_index41] = (_dafny.ZERO).minus(_152_v29);
      let _index42 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length));
      let _rhs47 = _120_v0;
      let _rhs48 = _154_v31;
      let _rhs49 = (_120_v0).isLessThanOrEqualTo(_120_v0);
      let _rhs50 = _133_v12;
      let _lhs38 = _125_v5;
      let _lhs39 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length));
      let _lhs40 = _137_globalState;
      let _lhs41 = _137_globalState;
      _lhs38[_lhs39] = _rhs47;
      _138_v16 = _rhs48;
      _lhs40.f3 = _rhs49;
      _lhs41.f24 = _rhs50;
      let _328_v151;
      _328_v151 = _dafny.Map.Empty.slice().updateUnsafe((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))],_133_v12);
      let _329_v152;
      _329_v152 = _dafny.MultiSet.fromElements((((_328_v151).contains(_152_v29)) ? ((_328_v151).get(_152_v29)) : (!(_133_v12))));
      if ((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_151_v28, _151_v28))).IsProperSubsetOf(_329_v152)) {
        (_137_globalState).f24 = true;
        (_137_globalState).f22 = (_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))];
        let _330_v153;
        _330_v153 = _dafny.Map.Empty.slice().updateUnsafe((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))],(_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]);
        let _331_v154;
        _331_v154 = _dafny.Map.Empty.slice().updateUnsafe(_120_v0,_module.D3.create_DC9(false, _330_v153));
        let _332_v155;
        let _nw47 = Array((new BigNumber(15)).toNumber()).fill(false);
        _332_v155 = _nw47;
        let _333_v156;
        _333_v156 = _module.D4.create_DC11(_331_v154, _332_v155);
        let _334_v157;
        _334_v157 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_133_v12),_333_v156);
        let _335_v158;
        _335_v158 = _dafny.Map.Empty.slice().updateUnsafe(_120_v0,new BigNumber((_334_v157).length));
        let _source6 = _module.D2.create_DC4((_335_v158).Merge(_330_v153));
        if (_source6.is_DC5) {
          let _336___mcc_h15 = (_source6).cf6;
          let _337___mcc_h16 = (_source6).cf7;
          let _338___mcc_h17 = (_source6).cf8;
          let _339___mcc_h18 = (_source6).cf9;
          let _340___mcc_h19 = (_source6).cf10;
          let _341_cf10 = _340___mcc_h19;
          let _342_cf9 = _339___mcc_h18;
          let _343_cf8 = _338___mcc_h17;
          let _344_cf7 = _337___mcc_h16;
          let _345_cf6 = _336___mcc_h15;
          let _346_v159;
          let _347_v160;
          let _348_v161;
          let _349_v162;
          let _out32;
          let _out33;
          let _out34;
          let _out35;
          let _outcollector8 = _module.__default.m0(_module.__default.fm0((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], _345_cf6, _137_globalState), _137_globalState);
          _out32 = _outcollector8[0];
          _out33 = _outcollector8[1];
          _out34 = _outcollector8[2];
          _out35 = _outcollector8[3];
          _346_v159 = _out32;
          _347_v160 = _out33;
          _348_v161 = _out34;
          _349_v162 = _out35;
          let _350_v163;
          _350_v163 = _module.D6.create_DC16(_121_v1);
          let _351_v164;
          _351_v164 = _134_v13;
          let _pat_let_tv8 = _121_v1;
          let _rhs51 = (function (_pat_let12_0) {
            return function (_352_dt__update__tmp_h6) {
              return function (_pat_let13_0) {
                return function (_353_dt__update_hcf22_h0) {
                  return _module.D6.create_DC16(_353_dt__update_hcf22_h0);
                }(_pat_let13_0);
              }(_pat_let_tv8);
            }(_pat_let12_0);
          }(_350_v163)).dtor_cf22;
          let _rhs52 = _133_v12;
          let _rhs53 = _324_v149;
          let _rhs54 = new BigNumber(((_351_v164)).length);
          let _lhs42 = _137_globalState;
          let _lhs43 = _137_globalState;
          _121_v1 = _rhs51;
          _lhs42.f24 = _rhs52;
          _324_v149 = _rhs53;
          _lhs43.f9 = _rhs54;
          (_137_globalState).f14 = _133_v12;
          let _354_v165;
          let _nw48 = new _module.C0();
          _nw48.__ctor();
          _354_v165 = _nw48;
        } else if (_source6.is_DC6) {
          let _355_v166;
          let _356_v167;
          let _357_v168;
          let _358_v169;
          let _out36;
          let _out37;
          let _out38;
          let _out39;
          let _outcollector9 = _module.__default.m0(_133_v12, _137_globalState);
          _out36 = _outcollector9[0];
          _out37 = _outcollector9[1];
          _out38 = _outcollector9[2];
          _out39 = _outcollector9[3];
          _355_v166 = _out36;
          _356_v167 = _out37;
          _357_v168 = _out38;
          _358_v169 = _out39;
          let _index43 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_332_v155).length));
          (_332_v155)[_index43] = _133_v12;
          let _index44 = _module.__default.safeIndex(new BigNumber(756), new BigNumber((_332_v155).length));
          (_332_v155)[_index44] = (_152_v29).isLessThanOrEqualTo(new BigNumber((_121_v1).length));
          _151_v28 = _355_v166;
          (_137_globalState).f9 = new BigNumber(102);
        } else if (_source6.is_DC4) {
          let _359___mcc_h20 = (_source6).cf5;
          let _360_cf5 = _359___mcc_h20;
          (_137_globalState).f21 = _129_v9;
          (_137_globalState).f24 = (new BigNumber((_151_v28).length)).isLessThan(_module.__default.safeDivisionInt(_152_v29, (_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]));
          let _361_v170;
          let _nw49 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _361_v170 = _nw49;
          let _362_v171;
          _362_v171 = _dafny.Map.Empty.slice().updateUnsafe(_323_v148,_258_v104);
          let _363_v172;
          _363_v172 = _module.D8.create_DC19((((_362_v171).contains(_323_v148)) ? ((_362_v171).get(_323_v148)) : (_258_v104)));
          let _index45 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_361_v170).length));
          (_361_v170)[_index45] = ((_363_v172).dtor_cf24).Merge(_258_v104);
          let _364_v173;
          _364_v173 = _dafny.Seq.of((_258_v104).update((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], _154_v31), _258_v104, _258_v104);
          let _index46 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_361_v170).length));
          (_361_v170)[_index46] = (_258_v104).Merge((_364_v173)[_module.__default.safeIndex((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], new BigNumber((_364_v173).length))]);
          let _365_v174;
          _365_v174 = _dafny.Map.Empty.slice().updateUnsafe(_151_v28,_dafny.Seq.Concat(_123_v3, _123_v3));
          _365_v174 = (_365_v174).update(_dafny.Seq.Concat(_151_v28, _151_v28), _dafny.Seq.Concat(_123_v3, _123_v3));
        } else {
          let _366___mcc_h21 = (_source6).cf11;
          let _367_cf11 = _366___mcc_h21;
          let _368_v175;
          _368_v175 = _dafny.Seq.of(_module.__default.fm7(_133_v12, _137_globalState));
          (_137_globalState).f14 = _dafny.Seq.IsPrefixOf(_368_v175, _dafny.Seq.Concat(_module.__default.fm8(_137_globalState), _368_v175));
          let _369_v176;
          let _370_v177;
          let _371_v178;
          let _372_v179;
          let _out40;
          let _out41;
          let _out42;
          let _out43;
          let _outcollector10 = _module.__default.m0(_133_v12, _137_globalState);
          _out40 = _outcollector10[0];
          _out41 = _outcollector10[1];
          _out42 = _outcollector10[2];
          _out43 = _outcollector10[3];
          _369_v176 = _out40;
          _370_v177 = _out41;
          _371_v178 = _out42;
          _372_v179 = _out43;
          _154_v31 = _module.__default.fm4(_372_v179, _151_v28, _137_globalState);
          let _rhs55 = (new BigNumber(-121)).multipliedBy((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]);
          let _rhs56 = _133_v12;
          let _lhs44 = _137_globalState;
          let _lhs45 = _137_globalState;
          _lhs44.f25 = _rhs55;
          _lhs45.f24 = _rhs56;
        }
        (_137_globalState).f14 = !(_133_v12);
        if ((_dafny.MultiSet.FromArray(_151_v28)).IsSubsetOf(_329_v152)) {
          _154_v31 = _154_v31;
          _154_v31 = ((false) ? (_138_v16) : (_154_v31));
          _125_v5 = _153_v30;
          let _373_v180;
          let _nw50 = new _module.C0();
          _nw50.__ctor();
          _373_v180 = _nw50;
          let _374_v181;
          _374_v181 = _dafny.Set.fromElements(true, _dafny.Seq.contains(_123_v3, _152_v29), _133_v12, !(_120_v0).isEqualTo(new BigNumber((_151_v28).length)), (_133_v12) === (_133_v12));
          let _index47 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length));
          let _rhs57 = (((_335_v158).contains(_152_v29)) ? ((_335_v158).get(_152_v29)) : ((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]));
          let _rhs58 = _125_v5;
          let _rhs59 = ((_374_v181).Union(_374_v181)).Union(_374_v181);
          let _rhs60 = _120_v0;
          let _rhs61 = _133_v12;
          let _lhs46 = _125_v5;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length));
          let _lhs48 = _137_globalState;
          let _lhs49 = _137_globalState;
          _lhs46[_lhs47] = _rhs57;
          _125_v5 = _rhs58;
          _374_v181 = _rhs59;
          _lhs48.f25 = _rhs60;
          _lhs49.f24 = _rhs61;
        } else {
          let _375_v182;
          let _nw51 = new _module.C0();
          _nw51.__ctor();
          _375_v182 = _nw51;
          (_137_globalState).f25 = (_dafny.ZERO).minus(_152_v29);
          (_137_globalState).f25 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_133_v12, _133_v12), _151_v28)).length);
          (_137_globalState).f25 = _152_v29;
          let _376_v183;
          let _nw52 = new _module.C0();
          _nw52.__ctor();
          _376_v183 = _nw52;
        }
      } else {
        let _377_v185;
        _377_v185 = _dafny.Map.Empty.slice().updateUnsafe(_152_v29,new BigNumber((function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_123_v3).Elements) {
            let _378_v184 = _compr_10;
            if (_dafny.Seq.contains(_123_v3, _378_v184)) {
              _coll10.push([(_378_v184).plus(_120_v0),new BigNumber((_128_v8).cardinality())]);
            }
          }
          return _coll10;
        }()).length));
        let _rhs62 = (_dafny.ZERO).minus(((new BigNumber((_377_v185).length)).plus((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))])).minus(_module.__default.safeDivisionInt(_152_v29, _152_v29)));
        let _rhs63 = _133_v12;
        let _rhs64 = _121_v1;
        let _lhs50 = _137_globalState;
        let _lhs51 = _137_globalState;
        _lhs50.f25 = _rhs62;
        _lhs51.f3 = _rhs63;
        _121_v1 = _rhs64;
        let _379_v186;
        let _380_v187;
        let _381_v188;
        let _382_v189;
        let _out44;
        let _out45;
        let _out46;
        let _out47;
        let _outcollector11 = _module.__default.m0(_133_v12, _137_globalState);
        _out44 = _outcollector11[0];
        _out45 = _outcollector11[1];
        _out46 = _outcollector11[2];
        _out47 = _outcollector11[3];
        _379_v186 = _out44;
        _380_v187 = _out45;
        _381_v188 = _out46;
        _382_v189 = _out47;
        let _383_v190;
        let _384_v191;
        let _385_v192;
        let _386_v193;
        let _out48;
        let _out49;
        let _out50;
        let _out51;
        let _outcollector12 = _module.__default.m0(!(_133_v12) || (_133_v12), _137_globalState);
        _out48 = _outcollector12[0];
        _out49 = _outcollector12[1];
        _out50 = _outcollector12[2];
        _out51 = _outcollector12[3];
        _383_v190 = _out48;
        _384_v191 = _out49;
        _385_v192 = _out50;
        _386_v193 = _out51;
        let _387_v194;
        let _nw53 = Array((new BigNumber(10)).toNumber());
        _387_v194 = _nw53;
        let _index48 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_387_v194).length));
        (_387_v194)[_index48] = _323_v148;
        let _388_v195;
        let _nw54 = Array((new BigNumber(14)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _133_v12;
        _nw54[(_dafny.ONE).toNumber()] = _133_v12;
        _nw54[(new BigNumber(2)).toNumber()] = ((_133_v12) ? (_133_v12) : (_133_v12));
        _nw54[(new BigNumber(3)).toNumber()] = (!(_133_v12)) && (_133_v12);
        _nw54[(new BigNumber(4)).toNumber()] = true;
        _nw54[(new BigNumber(5)).toNumber()] = !(_133_v12);
        _nw54[(new BigNumber(6)).toNumber()] = _133_v12;
        _nw54[(new BigNumber(7)).toNumber()] = _133_v12;
        _nw54[(new BigNumber(8)).toNumber()] = _133_v12;
        _nw54[(new BigNumber(9)).toNumber()] = _133_v12;
        _nw54[(new BigNumber(10)).toNumber()] = _133_v12;
        _nw54[(new BigNumber(11)).toNumber()] = ((_133_v12) ? (_133_v12) : (_133_v12));
        _nw54[(new BigNumber(12)).toNumber()] = !(_152_v29).isEqualTo(_152_v29);
        _nw54[(new BigNumber(13)).toNumber()] = _133_v12;
        _388_v195 = _nw54;
        let _index49 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length));
        (_388_v195)[_index49] = _133_v12;
        let _389_v196;
        _389_v196 = _dafny.Set.fromElements(_129_v9, _129_v9, _dafny.Seq.UnicodeFromString("bxdpjq"), _module.__default.fm2((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], _152_v29, false, (_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], _137_globalState), _129_v9);
        let _index50 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_387_v194).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length));
        let _rhs65 = _323_v148;
        let _rhs66 = _131_v11;
        let _rhs67 = _123_v3;
        let _rhs68 = _dafny.Seq.contains(_dafny.Seq.Concat(_383_v190, _383_v190), !_dafny.areEqual(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(781)), ((_390_v193) => function (_391_i10) {
          return _390_v193;
        })(_386_v193))));
        let _rhs69 = (_389_v196).IsProperSubsetOf(_389_v196);
        let _lhs52 = _387_v194;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_387_v194).length));
        let _lhs54 = _137_globalState;
        let _lhs55 = _388_v195;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length));
        _lhs52[_lhs53] = _rhs65;
        _131_v11 = _rhs66;
        _123_v3 = _rhs67;
        _lhs54.f3 = _rhs68;
        _lhs55[_lhs56] = _rhs69;
        if (false) {
          (_137_globalState).f9 = (_dafny.ZERO).minus(_380_v187);
          (_137_globalState).f25 = (_323_v148).fm5(_137_globalState);
          let _392_v197;
          let _393_v198;
          let _394_v199;
          let _395_v200;
          let _out52;
          let _out53;
          let _out54;
          let _out55;
          let _outcollector13 = _module.__default.m0(_133_v12, _137_globalState);
          _out52 = _outcollector13[0];
          _out53 = _outcollector13[1];
          _out54 = _outcollector13[2];
          _out55 = _outcollector13[3];
          _392_v197 = _out52;
          _393_v198 = _out53;
          _394_v199 = _out54;
          _395_v200 = _out55;
          let _396_v201;
          let _nw55 = new _module.C0();
          _nw55.__ctor();
          _396_v201 = _nw55;
          let _397_v202;
          _397_v202 = _dafny.Map.Empty.slice().updateUnsafe((_388_v195)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length))],_120_v0);
          _397_v202 = _397_v202;
        } else {
          (_137_globalState).f25 = new BigNumber((_383_v190).length);
          (_137_globalState).f24 = (new BigNumber((_123_v3).length)).isLessThanOrEqualTo((_384_v191).plus((_dafny.ZERO).minus(_152_v29)));
          let _398_v203;
          let _nw56 = new _module.C0();
          _nw56.__ctor();
          _398_v203 = _nw56;
          (_137_globalState).f14 = (_388_v195)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length))];
          (_137_globalState).f24 = (_388_v195)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_388_v195).length))];
        }
      }
      let _399_v204;
      _399_v204 = _dafny.Set.fromElements((_151_v28)[_module.__default.safeIndex(new BigNumber((_129_v9).length), new BigNumber((_151_v28).length))], _133_v12, _133_v12);
      if (((_399_v204).Difference(_399_v204)).IsDisjointFrom(_399_v204)) {
        let _400_v205;
        let _nw57 = Array((new BigNumber(17)).toNumber()).fill(false);
        _400_v205 = _nw57;
        let _index52 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_400_v205).length));
        (_400_v205)[_index52] = _133_v12;
        let _401_v206;
        _401_v206 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_328_v151).length),new BigNumber(-212));
        let _402_v207;
        _402_v207 = _module.D2.create_DC4(_401_v206);
        let _index53 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_400_v205).length));
        (_400_v205)[_index53] = _module.__default.fm0(new BigNumber((_dafny.Seq.Concat(_module.__default.fm2((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], (_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))], _133_v12, _152_v29, _137_globalState), _129_v9)).length), !((_402_v207).dtor_cf5).contains((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]), _137_globalState);
        (_137_globalState).f22 = (_dafny.ZERO).minus((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))]);
        let _403_v208;
        let _404_v209;
        let _405_v210;
        let _406_v211;
        let _out56;
        let _out57;
        let _out58;
        let _out59;
        let _outcollector14 = _module.__default.m0((_400_v205)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_400_v205).length))], _137_globalState);
        _out56 = _outcollector14[0];
        _out57 = _outcollector14[1];
        _out58 = _outcollector14[2];
        _out59 = _outcollector14[3];
        _403_v208 = _out56;
        _404_v209 = _out57;
        _405_v210 = _out58;
        _406_v211 = _out59;
        _153_v30 = _405_v210;
        (_137_globalState).f24 = (_400_v205)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_400_v205).length))];
      } else {
        if (!(!(new BigNumber((_151_v28).length)).isEqualTo((_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))])) || ((_151_v28)[_module.__default.safeIndex(_152_v29, new BigNumber((_151_v28).length))])) {
          let _407_v212;
          let _nw58 = Array((new BigNumber(19)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _127_v7;
          _nw58[(_dafny.ONE).toNumber()] = _127_v7;
          _nw58[(new BigNumber(2)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(3)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(4)).toNumber()] = ((_133_v12) ? (_127_v7) : (_127_v7));
          _nw58[(new BigNumber(5)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(6)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(7)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(8)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(9)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(10)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(11)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(12)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(13)).toNumber()] = ((_133_v12) ? (_127_v7) : (_127_v7));
          _nw58[(new BigNumber(14)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(15)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(16)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(17)).toNumber()] = _127_v7;
          _nw58[(new BigNumber(18)).toNumber()] = _127_v7;
          _407_v212 = _nw58;
          let _index54 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_407_v212).length));
          (_407_v212)[_index54] = _127_v7;
          let _index55 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_407_v212).length));
          (_407_v212)[_index55] = _127_v7;
          (_137_globalState).f24 = (_133_v12) && (!(_133_v12) || (_133_v12));
          let _408_v213;
          _408_v213 = _module.D6.create_DC17();
          let _rhs70 = _129_v9;
          let _rhs71 = _dafny.Seq.UnicodeFromString("smrlm");
          let _rhs72 = _408_v213;
          let _lhs57 = _137_globalState;
          let _lhs58 = _137_globalState;
          _lhs57.f21 = _rhs70;
          _lhs58.f21 = _rhs71;
          _408_v213 = _rhs72;
          _153_v30 = _153_v30;
          (_137_globalState).f24 = _module.__default.fm0(_152_v29, _module.__default.fm0(_120_v0, _133_v12, _137_globalState), _137_globalState);
        } else {
          let _409_v214;
          let _410_v215;
          let _411_v216;
          let _412_v217;
          let _out60;
          let _out61;
          let _out62;
          let _out63;
          let _outcollector15 = _module.__default.m0(_133_v12, _137_globalState);
          _out60 = _outcollector15[0];
          _out61 = _outcollector15[1];
          _out62 = _outcollector15[2];
          _out63 = _outcollector15[3];
          _409_v214 = _out60;
          _410_v215 = _out61;
          _411_v216 = _out62;
          _412_v217 = _out63;
          (_137_globalState).f14 = _133_v12;
          let _413_v218;
          _413_v218 = _dafny.Map.Empty.slice().updateUnsafe(_411_v216,false);
          _413_v218 = (_413_v218).Merge(_413_v218);
          let _414_v219;
          let _415_v220;
          let _416_v221;
          let _417_v222;
          let _out64;
          let _out65;
          let _out66;
          let _out67;
          let _outcollector16 = _module.__default.m0(_133_v12, _137_globalState);
          _out64 = _outcollector16[0];
          _out65 = _outcollector16[1];
          _out66 = _outcollector16[2];
          _out67 = _outcollector16[3];
          _414_v219 = _out64;
          _415_v220 = _out65;
          _416_v221 = _out66;
          _417_v222 = _out67;
          (_137_globalState).f21 = _129_v9;
        }
        let _418_v223;
        let _419_v224;
        let _420_v225;
        let _421_v226;
        let _out68;
        let _out69;
        let _out70;
        let _out71;
        let _outcollector17 = _module.__default.m0((_399_v204).IsDisjointFrom(_399_v204), _137_globalState);
        _out68 = _outcollector17[0];
        _out69 = _outcollector17[1];
        _out70 = _outcollector17[2];
        _out71 = _outcollector17[3];
        _418_v223 = _out68;
        _419_v224 = _out69;
        _420_v225 = _out70;
        _421_v226 = _out71;
        let _422_v227;
        let _423_v228;
        let _424_v229;
        let _425_v230;
        let _out72;
        let _out73;
        let _out74;
        let _out75;
        let _outcollector18 = _module.__default.m0(_133_v12, _137_globalState);
        _out72 = _outcollector18[0];
        _out73 = _outcollector18[1];
        _out74 = _outcollector18[2];
        _out75 = _outcollector18[3];
        _422_v227 = _out72;
        _423_v228 = _out73;
        _424_v229 = _out74;
        _425_v230 = _out75;
        (_137_globalState).f24 = _133_v12;
        (_137_globalState).f14 = _133_v12;
      }
      (_137_globalState).f9 = (_125_v5)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_125_v5).length))];
      _120_v0 = (_323_v148).fm5(_137_globalState);
      process.stdout.write(_dafny.toString(_120_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v1).equals(_dafny.Set.fromElements(new BigNumber(691)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(691)),new BigNumber(691)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_123_v3, _dafny.Seq.of(new BigNumber(691), new BigNumber(691)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.of(new BigNumber(691), new BigNumber(691))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_126_v6).dtor_cf0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ZERO])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[_dafny.ONE])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(2)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_v7)[new BigNumber(3)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v8).equals(_dafny.MultiSet.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_129_v9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_130_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v11)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_v12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_135_v14)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_135_v14)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v15).equals(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(691), new BigNumber(691))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState.f0).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(691)),new BigNumber(691)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f1).equals(_dafny.Set.fromElements(new BigNumber(691)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.of(new BigNumber(691), new BigNumber(691))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f10).equals(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(691), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f17)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_137_globalState).f19).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_137_globalState).f20)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_137_globalState).f20)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_137_globalState.f21).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_globalState.f23).equals(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(691), new BigNumber(691))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_globalState.f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_138_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(691)),new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_140_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_151_v28, _dafny.Seq.of(false, true, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_152_v29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v30)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_154_v31));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_258_v104).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-119716),new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_259_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_324_v149).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_325_v150).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(691),_dafny.Seq.UnicodeFromString("px")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_328_v151).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(691),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_329_v152).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_399_v204).equals(_dafny.Set.fromElements(false, true))));
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
    static create_DC2(cf2, cf3) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
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
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
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
    static create_DC3(cf4) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
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
    static create_DC5(cf6, cf7, cf8, cf9, cf10) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D2(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false, false, false, _dafny.ZERO, false);
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
    static create_DC9(cf13, cf14) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC8(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(false, _dafny.Map.Empty);
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
    static create_DC11(cf16, cf17) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.Map.Empty, []);
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
    static create_DC13() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC14(cf19, cf20) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D5(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC15(cf21) {
      let $dt = new D5(3);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC15() { return this.$tag === 3; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf21) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13();
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
    static create_DC17() {
      let $dt = new D6(0);
      return $dt;
    }
    static create_DC16(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17();
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
    static create_DC18(cf23) {
      let $dt = new D7(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf25, cf26) {
      let $dt = new D8(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC19(cf24) {
      let $dt = new D8(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC21(cf27) {
      let $dt = new D8(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.Map.Empty;
      this.f3 = false;
      this.f9 = _dafny.ZERO;
      this.f14 = false;
      this.f21 = _dafny.Seq.UnicodeFromString("");
      this.f22 = _dafny.ZERO;
      this.f23 = _dafny.Set.Empty;
      this.f24 = false;
      this.f25 = _dafny.ZERO;
      this._f1 = _dafny.Set.Empty;
      this._f2 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = _dafny.Map.Empty;
      this._f7 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
      this._f10 = _dafny.Set.Empty;
      this._f11 = _dafny.ZERO;
      this._f12 = false;
      this._f13 = false;
      this._f15 = _dafny.ZERO;
      this._f16 = false;
      this._f17 = [];
      this._f18 = _dafny.ZERO;
      this._f19 = _dafny.Map.Empty;
      this._f20 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this).f22 = f22;
      (_this).f23 = f23;
      (_this).f24 = f24;
      (_this).f25 = f25;
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
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
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
    fm5(globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.UnicodeFromString("rpbp")).length)).minus((new BigNumber(704)).multipliedBy(new BigNumber(515)));
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
