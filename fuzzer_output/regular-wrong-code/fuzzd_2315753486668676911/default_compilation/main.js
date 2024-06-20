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
    static fm5(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(643)), function (_0_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("svlade")), _dafny.Seq.UnicodeFromString("ls"));
    };
    static fm11(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-106))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(767)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("eoqmvc")).length))));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC9(((true) ? (new _dafny.CodePoint('v'.codePointAt(0))) : (new _dafny.CodePoint('v'.codePointAt(0)))));
    };
    static fm17(p0, globalState) {
      return (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false)))).cardinality())).plus((new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(492), new BigNumber(-555))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(492)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(-555)))) {
            _coll0.add(_module.__default.safeDivisionInt(_1_v0, new BigNumber((function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of _dafny.IntegerRange(new BigNumber(581), new BigNumber(375))) {
                let _2_v1 = _compr_1;
                if (((new BigNumber(581)).isLessThanOrEqualTo(_2_v1)) && ((_2_v1).isLessThan(new BigNumber(375)))) {
                  _coll1.push([_module.__default.safeModuloInt(_2_v1, new BigNumber((_dafny.Seq.UnicodeFromString("vfftcgcr")).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(540)), function (_3_i0) {
                    return new _dafny.CodePoint('s'.codePointAt(0));
                  })).length)]);
                }
              }
              return _coll1;
            }()).length)));
          }
        }
        return _coll0;
      }()).length)).multipliedBy(new BigNumber(431)));
    };
    static fm18(p0, p1, p2, globalState) {
      return function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-246)), function (_4_i0) {
          return new BigNumber(316);
        }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-167)), function (_5_i1) {
          return new BigNumber(-88);
        }))).Elements) {
          let _6_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-246)), function (_4_i0) {
            return new BigNumber(316);
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-167)), function (_5_i1) {
            return new BigNumber(-88);
          })), _6_v0)) {
            _coll2.push([_module.__default.safeModuloInt(_6_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber(453), new BigNumber(243))).length)),!(true)]);
          }
        }
        return _coll2;
      }();
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("hafht");
    };
    static fm21(p0, p1, globalState) {
      return _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("scdrbxig"), false);
    };
    static fm22(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(((!(!(!(false)))) ? (true) : (!(true)))),!(!(false)));
    };
    static fm26(globalState) {
      return true;
    };
    static fm27(globalState) {
      if (true) {
        return _dafny.Seq.of(false);
      } else if (false) {
        return _dafny.Seq.of(false);
      } else {
        return (_module.D14.create_DC36((_module.D14.create_DC36(_dafny.Seq.of(true, false))).dtor_cf53)).dtor_cf53;
      }
    };
    static fm30(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D14.create_DC35(_dafny.Seq.of(function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (function () {
    let _coll4 = new _dafny.Map();
    for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-786), new BigNumber(405))) {
      let _7_v1 = _compr_4;
      if (((new BigNumber(-786)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(405)))) {
        _coll4.push([(_7_v1).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), function (_8_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        })).length)),false]);
      }
    }
    return _coll4;
  }()).Keys.Elements) {
    let _9_v0 = _compr_3;
    if ((function () {
      let _coll5 = new _dafny.Map();
      for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-786), new BigNumber(405))) {
        let _7_v1 = _compr_5;
        if (((new BigNumber(-786)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(405)))) {
          _coll5.push([(_7_v1).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(628)), function (_8_i0) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })).length)),false]);
        }
      }
      return _coll5;
    }()).contains(_9_v0)) {
      _coll3.push([(_9_v0).minus(new BigNumber(696)),false]);
    }
  }
  return _coll3;
}(), function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(227), new BigNumber(175))) {
    let _10_v2 = _compr_6;
    if (((new BigNumber(227)).isLessThanOrEqualTo(_10_v2)) && ((_10_v2).isLessThan(new BigNumber(175)))) {
      _coll6.push([_module.__default.safeModuloInt(_10_v2, new BigNumber((_dafny.Seq.of(new BigNumber(-493))).length)),false]);
    }
  }
  return _coll6;
}()));
      if (_source0.is_DC36) {
        let _11___mcc_h0 = (_source0).cf53;
        let _12_cf53 = _11___mcc_h0;
        return _dafny.Seq.UnicodeFromString("fsqaxwqte");
      } else if (_source0.is_DC35) {
        let _13___mcc_h1 = (_source0).cf52;
        let _14_cf52 = _13___mcc_h1;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dxhjglaic"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-301)), function (_15_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
      } else {
        let _16___mcc_h2 = (_source0).cf54;
        let _17_cf54 = _16___mcc_h2;
        return _dafny.Seq.UnicodeFromString("mplnkdwif");
      }
    };
    static fm31(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(72))).Union(_dafny.MultiSet.fromElements(new BigNumber(77), new BigNumber((_dafny.MultiSet.fromElements(!((_module.D20.create_DC55(new BigNumber((_dafny.Seq.of(new BigNumber(186), new BigNumber((_dafny.Set.fromElements(false, false)).length))).length), new BigNumber(((_module.D15.create_DC38(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_18_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length),true))).dtor_cf55).length), false, true)).dtor_cf87), true, true)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("xmniuvfes")).length), new BigNumber(728), new BigNumber(310)));
    };
    static fm32(p0, p1, p2, globalState) {
      return (function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(23), new BigNumber(808))) {
          let _19_v0 = _compr_7;
          if (((new BigNumber(23)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(808)))) {
            _coll7.add((_19_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("kwgcuel")).length)));
          }
        }
        return _coll7;
      }()).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(-804), new BigNumber(56))) {
          let _20_v1 = _compr_8;
          if (((new BigNumber(-804)).isLessThanOrEqualTo(_20_v1)) && ((_20_v1).isLessThan(new BigNumber(56)))) {
            _coll8.push([(_20_v1).multipliedBy(new BigNumber(867)),false]);
          }
        }
        return _coll8;
      }()).length), new BigNumber(-849), new BigNumber((_dafny.Seq.UnicodeFromString("tm")).length), new BigNumber(84), new BigNumber(249))).length), new BigNumber(843), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)))).length)));
    };
    static fm33(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((false) || (!(!(false))),false);
    };
    static fm34(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(true, false)).Intersect(_dafny.MultiSet.fromElements(false, true))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true))));
    };
    static fm35(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(113)).plus(new BigNumber(-13)),(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-78))).contains(false));
    };
    static fm36(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(false)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(true, false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(false, true))));
    };
    static fm37(globalState) {
      return ((_dafny.Set.fromElements(false, false)).Difference(_dafny.Set.fromElements(false))).Union((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(false, true)));
    };
    static fm38(p0, globalState) {
      let _source1 = _module.D12.create_DC33(true);
      if (_source1.is_DC33) {
        let _21___mcc_h0 = (_source1).cf50;
        let _22_cf50 = _21___mcc_h0;
        return new _dafny.CodePoint('s'.codePointAt(0));
      } else {
        let _23___mcc_h1 = (_source1).cf49;
        let _24_cf49 = _23___mcc_h1;
        if (true) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('t'.codePointAt(0));
        }
      }
    };
    static fm39(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(669), new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(930)), _dafny.MultiSet.fromElements(new BigNumber(897), new BigNumber(342), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(503)), function (_25_i0) {
        return new BigNumber(162);
      })).length)))).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(732))).length),new BigNumber((_dafny.Seq.UnicodeFromString("fnqkq")).length))).length))), ((false) ? (_dafny.Seq.of(new BigNumber(157), new BigNumber((_dafny.Seq.UnicodeFromString("rysmqpp")).length))) : (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-412)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber(193), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ng")).length),false)).length),new _dafny.CodePoint('x'.codePointAt(0)))).length))).length), new BigNumber(-685))).length)))));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC2(((true) ? (new BigNumber((_dafny.Seq.UnicodeFromString("clskkt")).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), function (_26_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length))));
    };
    static fm41(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hjghubcnr")).length),new BigNumber((_dafny.Seq.of(false, !((_module.D19.create_DC52(new BigNumber(272), true, true)).dtor_cf82))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-738)), function (_27_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length),new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cxpjsdcn")).length)))).Elements) {
          let _28_v0 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cxpjsdcn")).length))), _28_v0)) {
            _coll9.add(_module.__default.safeDivisionInt(_28_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(854))).cardinality())));
          }
        }
        return _coll9;
      }()).length)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(448),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-717),new BigNumber(806))));
    };
    static fm42(p0, p1, globalState) {
      if ((!(true)) === (!(true))) {
        return _module.D1.create_DC4(_dafny.MultiSet.fromElements(new BigNumber(-88)));
      } else {
        return _module.D1.create_DC4(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(819))).length)));
      }
    };
    static fm43(p0, globalState) {
      return (_dafny.MultiSet.fromElements(function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(858), new BigNumber(165))) {
          let _29_v0 = _compr_10;
          if (((new BigNumber(858)).isLessThanOrEqualTo(_29_v0)) && ((_29_v0).isLessThan(new BigNumber(165)))) {
            _coll10.add((_29_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()))).length)));
          }
        }
        return _coll10;
      }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("ade")).length))).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("euwroheny")).length)))).Difference((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("icwgdatuh")).length), new BigNumber(632)))).Intersect(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-518),false)).length)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pfnisofy")).length)), _dafny.Set.fromElements(new BigNumber(-276)))));
    };
    static fm44(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf((_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("frlkbglny"), !(false))).dtor_cf3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-612)), function (_30_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })),new BigNumber(698));
    };
    static fm45(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("klqhbvu")).length),false)), _dafny.Seq.of(function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(-567), new BigNumber(334))) {
          let _31_v0 = _compr_11;
          if (((new BigNumber(-567)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(334)))) {
            _coll11.push([(_31_v0).minus(new BigNumber((_dafny.Seq.of(false)).length)),(_module.D12.create_DC33(true)).dtor_cf50]);
          }
        }
        return _coll11;
      }(), function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(393)), function (_32_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(true), _dafny.Seq.of(false))).length));
        })).Elements) {
          let _33_v1 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(393)), function (_32_i0) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(true), _dafny.Seq.of(false))).length));
          }), _33_v1)) {
            _coll12.push([_module.__default.safeModuloInt(_33_v1, (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length))),!(!(true))]);
          }
        }
        return _coll12;
      }()));
    };
    static fm46(p0, globalState) {
      return _module.D16.create_DC42(!(false), true, (true) === (true));
    };
    static fm47(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(29)), function (_34_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(168)), function (_35_i1) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("pabhldsjv"), _dafny.Seq.UnicodeFromString("cksrrod"), _dafny.Seq.UnicodeFromString("mwrut"))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ddnqnu"), _dafny.Seq.UnicodeFromString("nb")))).Intersect((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("itm"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_36_i2) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("juvss")))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(669)), function (_37_i3) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }))));
    };
    static fm48(p0, p1, p2, globalState) {
      return _module.D14.create_DC35(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(582), new BigNumber((_dafny.Seq.of(true, true, !(false))).length))).length),!(false)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(343))).length),false)));
    };
    static fm49(p0, p1, globalState) {
      return _module.D14.create_DC36(_dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(false, false)));
    };
    static fm50(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_module.D2.create_DC6(new BigNumber(-174))))).Union(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("udq")).length)), _module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("ivkmw")).length)))))).Intersect(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D2.create_DC6(new BigNumber(214)), _module.D2.create_DC6(new BigNumber(533)))), _dafny.MultiSet.fromElements(_module.D2.create_DC6(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), function (_38_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dxvtbcyv")).length),true)).length), new BigNumber((_dafny.Seq.of(true)).length))).length)), _module.D2.create_DC6(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), function (_39_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length))), _dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D2.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))), _dafny.MultiSet.FromArray(_dafny.Seq.of(_module.D2.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(_module.D16.create_DC42(false, false, true), _module.D16.create_DC42(!(true), !(false), false), _module.D16.create_DC42(!(false), false, false))).cardinality())), _module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("resvq")).length)), _module.D2.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), _module.D2.create_DC6(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("lhcaqip")).length))).length)), _module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("og")).length)))), _dafny.MultiSet.fromElements(_module.D2.create_DC6(new BigNumber(473)), _module.D2.create_DC6(new BigNumber((function () {
  let _coll13 = new _dafny.Map();
  for (const _compr_13 of _dafny.IntegerRange(new BigNumber(-474), new BigNumber(-262))) {
    let _40_v0 = _compr_13;
    if (((new BigNumber(-474)).isLessThanOrEqualTo(_40_v0)) && ((_40_v0).isLessThan(new BigNumber(-262)))) {
      _coll13.push([(_40_v0).multipliedBy(new BigNumber(357)),true]);
    }
  }
  return _coll13;
}()).length)))));
    };
    static fm51(p0, p1, globalState) {
      return _module.D11.create_DC31(_module.D11.create_DC31(_module.D11.create_DC30(true, new BigNumber(684))));
    };
    static fm52(globalState) {
      return _dafny.Seq.Concat(((true) ? (_dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC28(function () {
  let _coll14 = new _dafny.Set();
  for (const _compr_14 of (_dafny.Seq.of(new BigNumber((function () {
    let _coll15 = new _dafny.Map();
    for (const _compr_15 of _dafny.IntegerRange(new BigNumber(334), new BigNumber(762))) {
      let _41_v0 = _compr_15;
      if (((new BigNumber(334)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(762)))) {
        _coll15.push([_module.__default.safeDivisionInt(_41_v0, new BigNumber(610)),new BigNumber((_dafny.Set.fromElements(!(true))).length)]);
      }
    }
    return _coll15;
  }()).length))).Elements) {
    let _42_v1 = _compr_14;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((function () {
      let _coll16 = new _dafny.Map();
      for (const _compr_16 of _dafny.IntegerRange(new BigNumber(334), new BigNumber(762))) {
        let _41_v0 = _compr_16;
        if (((new BigNumber(334)).isLessThanOrEqualTo(_41_v0)) && ((_41_v0).isLessThan(new BigNumber(762)))) {
          _coll16.push([_module.__default.safeDivisionInt(_41_v0, new BigNumber(610)),new BigNumber((_dafny.Set.fromElements(!(true))).length)]);
        }
      }
      return _coll16;
    }()).length)), _42_v1)) {
      _coll14.add((_42_v1).multipliedBy(new BigNumber((_dafny.Seq.of(true, false, false)).length)));
    }
  }
  return _coll14;
}())), _module.D11.create_DC31(_module.D11.create_DC29()))) : (_dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC29()), _module.D11.create_DC31(_module.D11.create_DC29()), _module.D11.create_DC31(_module.D11.create_DC28(_dafny.Set.fromElements(new BigNumber(-644))))))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), function (_43_i0) {
        return _module.D11.create_DC31(_module.D11.create_DC28(_dafny.Set.fromElements(new BigNumber(695))));
      }), _dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC29()), _module.D11.create_DC31(_module.D11.create_DC31(_module.D11.create_DC30(true, new BigNumber((function () {
  let _coll17 = new _dafny.Map();
  for (const _compr_17 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("fviyvd")).length),new _dafny.CodePoint('h'.codePointAt(0)))).length), new BigNumber(-349))).cardinality()),false)).Keys.Elements) {
    let _44_v2 = _compr_17;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("fviyvd")).length),new _dafny.CodePoint('h'.codePointAt(0)))).length), new BigNumber(-349))).cardinality()),false)).contains(_44_v2)) {
      _coll17.push([(_44_v2).plus(new BigNumber(461)),_dafny.MultiSet.fromElements(true)]);
    }
  }
  return _coll17;
}()).length)))))));
    };
    static fm53(p0, globalState) {
      return _module.D3.create_DC10(new BigNumber(-282), _module.__default.safeDivisionInt(new BigNumber(-116), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),false)).length)));
    };
    static fm54(p0, p1, p2, globalState) {
      let _source2 = _module.D12.create_DC32(_dafny.Set.fromElements(false, false, false));
      if (_source2.is_DC33) {
        let _45___mcc_h0 = (_source2).cf50;
        let _46_cf50 = _45___mcc_h0;
        return _module.D19.create_DC53();
      } else {
        let _47___mcc_h1 = (_source2).cf49;
        let _48_cf49 = _47___mcc_h1;
        return _module.D19.create_DC53();
      }
    };
    static fm55(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),true)), (function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new _dafny.CodePoint('t'.codePointAt(0)))).Keys.Elements) {
          let _49_v0 = _compr_18;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new _dafny.CodePoint('t'.codePointAt(0)))).contains(_49_v0)) {
            _coll18.push([_49_v0,false]);
          }
        }
        return _coll18;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true)));
    };
    static fm56(p0, p1, globalState) {
      return _module.D7.create_DC22((_dafny.Set.fromElements(!(false), false)).IsSubsetOf(_dafny.Set.fromElements(!(false))), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cwvgtdwty"), _dafny.Seq.UnicodeFromString("q"))).length), (_dafny.MultiSet.fromElements(false, true, true, true, true)).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, true, true)), (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(498)), function (_50_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length);
})).length)).multipliedBy(new BigNumber(-797)));
    };
    static fm57(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D7.create_DC22(false, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()), false, new BigNumber(483)), _module.D7.create_DC22(true, new BigNumber(-171), false, new BigNumber(747)))).Difference(_dafny.MultiSet.fromElements(_module.D7.create_DC22(false, new BigNumber(640), false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(219),new BigNumber(-399))).length)), _module.D7.create_DC22(!(!(true)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)),new _dafny.CodePoint('t'.codePointAt(0)))).length))), false, new BigNumber((_dafny.Set.fromElements(false, true, true)).length))))).Intersect(_dafny.MultiSet.fromElements(_module.D7.create_DC22(!(true), new BigNumber(898), false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-133)), function (_51_i0) {
  return new _dafny.CodePoint('h'.codePointAt(0));
})).length))));
    };
    static fm58(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-157),_module.D21.create_DC57(new BigNumber(428), new BigNumber(475), _module.D5.create_DC14(new _dafny.CodePoint('p'.codePointAt(0)), true, new BigNumber((_dafny.Set.fromElements(true)).length), (_module.D18.create_DC49(true, !(false), new _dafny.CodePoint('o'.codePointAt(0)), false)).dtor_cf76)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("irgs")).length),_module.D21.create_DC57(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(315))).cardinality()), new BigNumber(639), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), function (_52_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length)), new BigNumber(680), new BigNumber(30))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("yptolnu")).length), _module.D5.create_DC14(new _dafny.CodePoint('o'.codePointAt(0)), false, new BigNumber((_dafny.Seq.UnicodeFromString("tusd")).length), false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(490),_module.D21.create_DC57(new BigNumber(814), new BigNumber(-888), _module.D5.create_DC14(new _dafny.CodePoint('i'.codePointAt(0)), !(true), new BigNumber(374), true)))));
    };
    static fm59(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-493))), new BigNumber(469))).cardinality()),new BigNumber((_dafny.Seq.UnicodeFromString("xamnu")).length)));
    };
    static fm60(p0, p1, globalState) {
      return _module.D19.create_DC52(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("i")).length), new BigNumber(994)), (_dafny.MultiSet.fromElements(false)).IsSubsetOf(_dafny.MultiSet.fromElements(false)), !(!(false)));
    };
    static Main(__noArgsParameter) {
      let _53_v0;
      _53_v0 = false;
      let _54_v1;
      _54_v1 = new BigNumber(215);
      let _55_v2;
      _55_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(_53_v0),_54_v1);
      let _56_v3;
      _56_v3 = _dafny.Set.fromElements(_54_v1);
      let _57_v4;
      _57_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_53_v0);
      let _58_v5;
      _58_v5 = _dafny.Seq.of(_53_v0);
      let _59_v6;
      _59_v6 = _dafny.MultiSet.fromElements(_58_v5);
      let _60_globalState;
      let _nw0 = new _module.GlobalState();
      _nw0.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_53_v0, _53_v0), _dafny.Seq.of(_53_v0)), true, new _dafny.CodePoint('t'.codePointAt(0)), new BigNumber(51), new BigNumber(503), true, _55_v2, new BigNumber(231), _56_v3, true, true, (_57_v4).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_53_v0),_53_v0)), new BigNumber(38), new BigNumber(113), true, new BigNumber(964), _59_v6, true, new BigNumber(454), true);
      _60_globalState = _nw0;
      let _61_v7;
      let _nw1 = new _module.C4();
      _nw1.__ctor();
      _61_v7 = _nw1;
      let _hi0 = _54_v1;
      for (let _62_i0 = _54_v1; _62_i0.isLessThan(_hi0); _62_i0 = _62_i0.plus(_dafny.ONE)) {
        let _63_v8;
        let _init0 = ((_64_v0) => function (_65_i1) {
          return _64_v0;
        })(_53_v0);
        let _nw2 = Array((new BigNumber(23)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
          _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _63_v8 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
        (_63_v8)[_index0] = !(((_53_v0) ? (true) : (_53_v0)));
        let _66_v9;
        _66_v9 = _dafny.Seq.UnicodeFromString("eabrtymy");
        let _index1 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
        let _rhs0 = !(true) || (_53_v0);
        let _rhs1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_67_i2) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        });
        let _rhs2 = _54_v1;
        let _rhs3 = ((_53_v0) ? ((new BigNumber(((_dafny.MultiSet.FromArray(_58_v5)).update(_53_v0, _module.__default.abs(_62_i0))).cardinality())).isEqualTo(_module.__default.fm17(_54_v1, _60_globalState))) : (_53_v0));
        let _rhs4 = _62_i0;
        let _lhs0 = _63_v8;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
        _lhs0[_lhs1] = _rhs0;
        _66_v9 = _rhs1;
        _54_v1 = _rhs2;
        _53_v0 = _rhs3;
        _54_v1 = _rhs4;
        _58_v5 = _58_v5;
        let _68_v10;
        _68_v10 = _dafny.Seq.of(_66_v9, _66_v9, _66_v9, _66_v9);
        _68_v10 = _68_v10;
        if (true) {
          let _69_v11;
          _69_v11 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_module.D4.create_DC11(_68_v10));
          let _70_v12;
          _70_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_57_v4).length),_63_v8);
          let _71_v13;
          _71_v13 = _dafny.Set.fromElements((((_70_v12).contains(new BigNumber(898))) ? ((_70_v12).get(new BigNumber(898))) : (_63_v8)));
          let _72_v14;
          _72_v14 = _module.D1.create_DC2(_62_i0);
          let _index2 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
          let _rhs5 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(918), new BigNumber((_dafny.Set.fromElements(_54_v1, (_dafny.ZERO).minus(_62_i0), _62_i0)).length)))).plus(_54_v1);
          let _rhs6 = (((_61_v7).fm23(_69_v11, new BigNumber((_66_v9).length), new BigNumber((_71_v13).length), _60_globalState)).multipliedBy(new BigNumber((_66_v9).length))).isLessThanOrEqualTo(new BigNumber(((_56_v3).Intersect(_dafny.Set.fromElements((_72_v14).dtor_cf2))).length));
          let _lhs2 = _60_globalState;
          let _lhs3 = _63_v8;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
          _lhs2.f15 = _rhs5;
          _lhs3[_lhs4] = _rhs6;
          (_60_globalState).f4 = new BigNumber(642);
          let _73_v15;
          _73_v15 = _module.D0.create_DC1(_63_v8);
          let _74_v16;
          _74_v16 = _dafny.Map.Empty.slice().updateUnsafe((_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))],_63_v8);
          _73_v15 = _module.D0.create_DC1((((_74_v16).contains((_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))])) ? ((_74_v16).get((_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))])) : (_63_v8)));
          let _index3 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
          let _rhs7 = (_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))];
          let _rhs8 = ((_53_v0) ? ((_dafny.ZERO).minus(_54_v1)) : (_54_v1));
          let _lhs5 = _63_v8;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length));
          _lhs5[_lhs6] = _rhs7;
          _54_v1 = _rhs8;
          let _75_v17;
          _75_v17 = _dafny.MultiSet.fromElements(_53_v0, (_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))]);
          let _76_v18;
          let _nw3 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _76_v18 = _nw3;
          let _77_v19;
          _77_v19 = _module.D2.create_DC7(false, _75_v17, _76_v18);
          let _78_v20;
          _78_v20 = _module.D2.create_DC8(_77_v19);
          let _79_v21;
          let _nw4 = new _module.C3();
          _nw4.__ctor(_78_v20, (_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))]);
          _79_v21 = _nw4;
          let _80_v22;
          _80_v22 = _dafny.Map.Empty.slice().updateUnsafe(_79_v21,(_62_i0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ojorbrkgu")).length)));
          _80_v22 = ((_79_v21.f25) ? ((_80_v22).update(_79_v21, _62_i0)) : ((_80_v22).Merge(_80_v22)));
        } else {
          let _81_v23;
          _81_v23 = _dafny.MultiSet.fromElements(_62_i0, _62_i0, (new BigNumber(86)).multipliedBy(_54_v1));
          _81_v23 = _81_v23;
          let _82_v24;
          _82_v24 = new _dafny.CodePoint('w'.codePointAt(0));
          let _83_v25;
          _83_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-880),_82_v24);
          _83_v25 = (_83_v25).update(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(953)), ((_84_v1) => function (_85_i3) {
            return _84_v1;
          })(_54_v1))).length), _62_i0), ((!((_63_v8)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_63_v8).length))])) ? (_82_v24) : (_82_v24)));
          (_60_globalState).f15 = _54_v1;
          let _86_v26;
          let _nw5 = new _module.C8();
          _nw5.__ctor();
          _86_v26 = _nw5;
          let _87_v27;
          let _nw6 = Array((new BigNumber(10)).toNumber()).fill([]);
          _87_v27 = _nw6;
          _87_v27 = _87_v27;
        }
      }
      let _88_v28;
      let _nw7 = new _module.C6();
      _nw7.__ctor();
      _88_v28 = _nw7;
      let _89_v29;
      _89_v29 = _dafny.MultiSet.fromElements(_88_v28);
      let _90_v30;
      _90_v30 = _module.D20.create_DC55(_54_v1, _54_v1, _53_v0, !(_dafny.MultiSet.fromElements(_88_v28)).equals(_89_v29));
      let _source3 = _90_v30;
      if (_source3.is_DC55) {
        let _91___mcc_h0 = (_source3).cf85;
        let _92___mcc_h1 = (_source3).cf86;
        let _93___mcc_h2 = (_source3).cf87;
        let _94___mcc_h3 = (_source3).cf88;
        let _95_cf88 = _94___mcc_h3;
        let _96_cf87 = _93___mcc_h2;
        let _97_cf86 = _92___mcc_h1;
        let _98_cf85 = _91___mcc_h0;
        let _99_v31;
        _99_v31 = _dafny.Seq.of(_97_cf86);
        let _100_v32;
        _100_v32 = _dafny.Seq.UnicodeFromString("dnextlf");
        (_60_globalState).f15 = (_54_v1).minus((_99_v31)[_module.__default.safeIndex(new BigNumber((_100_v32).length), new BigNumber((_99_v31).length))]);
        let _101_v33;
        _101_v33 = _module.D16.create_DC42(false, _96_cf87, _95_cf88);
        let _102_v34;
        _102_v34 = _dafny.Seq.of(_101_v33, _101_v33, _101_v33);
        let _rhs9 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(740)), ((_103_v33) => function (_104_i4) {
          return _103_v33;
        })(_101_v33)), _102_v34);
        let _rhs10 = (_99_v31)[_module.__default.safeIndex(_97_cf86, new BigNumber((_99_v31).length))];
        let _rhs11 = (_98_cf85).plus(_97_cf86);
        let _rhs12 = !(_95_cf88) || (_95_cf88);
        _102_v34 = _rhs9;
        _98_cf85 = _rhs10;
        _97_cf86 = _rhs11;
        _53_v0 = _rhs12;
        let _105_v35;
        let _nw8 = new _module.C8();
        _nw8.__ctor();
        _105_v35 = _nw8;
        (_60_globalState).f7 = _98_cf85;
      } else {
        let _106___mcc_h4 = (_source3).cf84;
        let _107_cf84 = _106___mcc_h4;
        let _108_v36;
        let _109_v37;
        let _110_v38;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_88_v28).m9(_54_v1, _53_v0, !(!(_53_v0)), _60_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _108_v36 = _out0;
        _109_v37 = _out1;
        _110_v38 = _out2;
        _110_v38 = _53_v0;
        _53_v0 = _53_v0;
        _54_v1 = (_dafny.ZERO).minus((new BigNumber((_58_v5).length)).minus(((false) ? (_54_v1) : (_108_v36))));
      }
      let _111_v39;
      _111_v39 = _module.D11.create_DC29();
      let _112_v40;
      _112_v40 = _module.D11.create_DC31(_111_v39);
      let _113_v41;
      _113_v41 = _module.D11.create_DC31(_111_v39);
      let _114_v42;
      _114_v42 = _dafny.Seq.of(_113_v41, _113_v41);
      let _115_v43;
      _115_v43 = _module.D19.create_DC51(_114_v42);
      _115_v43 = _115_v43;
      let _116_v44;
      _116_v44 = _dafny.Seq.of((((_55_v2).contains(_53_v0)) ? ((_55_v2).get(_53_v0)) : (_54_v1)));
      (_60_globalState).f15 = new BigNumber((_116_v44).length);
      (_60_globalState).f15 = (new BigNumber(253)).plus(_54_v1);
      let _117_v45;
      let _nw9 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
      _117_v45 = _nw9;
      _117_v45 = _117_v45;
      _53_v0 = !(_53_v0);
      let _118_v46;
      _118_v46 = _dafny.MultiSet.fromElements(_56_v3, _56_v3);
      let _119_v47;
      _119_v47 = _module.D5.create_DC15(_module.D5.create_DC13(_118_v46));
      let _120_v48;
      _120_v48 = _dafny.Seq.of(_119_v47, _119_v47, _119_v47);
      _120_v48 = _120_v48;
      _53_v0 = _53_v0;
      let _121_v49;
      let _nw10 = Array((new BigNumber(10)).toNumber());
      _nw10[(_dafny.ZERO).toNumber()] = _58_v5;
      _nw10[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_53_v0);
      _nw10[(new BigNumber(2)).toNumber()] = _58_v5;
      _nw10[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_53_v0);
      _nw10[(new BigNumber(4)).toNumber()] = _58_v5;
      _nw10[(new BigNumber(5)).toNumber()] = _58_v5;
      _nw10[(new BigNumber(6)).toNumber()] = _58_v5;
      _nw10[(new BigNumber(7)).toNumber()] = _58_v5;
      _nw10[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_58_v5, _58_v5);
      _nw10[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_58_v5, _module.__default.safeIndex(_54_v1, new BigNumber((_58_v5).length)), _53_v0);
      _121_v49 = _nw10;
      let _122_v50;
      _122_v50 = _dafny.Seq.UnicodeFromString("ipyktu");
      let _123_v51;
      let _nw11 = new _module.C12();
      _nw11.__ctor(new BigNumber((_122_v50).length), _dafny.Seq.UnicodeFromString("xofwgsc"));
      _123_v51 = _nw11;
      let _124_v52;
      let _nw12 = Array((new BigNumber(4)).toNumber());
      _nw12[(_dafny.ZERO).toNumber()] = _123_v51;
      _nw12[(_dafny.ONE).toNumber()] = _123_v51;
      _nw12[(new BigNumber(2)).toNumber()] = _123_v51;
      _nw12[(new BigNumber(3)).toNumber()] = _123_v51;
      _124_v52 = _nw12;
      let _125_v53;
      _125_v53 = _dafny.Seq.of(_124_v52, _124_v52, _124_v52, _124_v52);
      let _index4 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_121_v49).length));
      (_121_v49)[_index4] = _dafny.Seq.of(false, (_61_v7).fm8(new BigNumber((_125_v53).length), _53_v0, (_123_v51).f22, _60_globalState));
      let _index5 = _module.__default.safeIndex(new BigNumber(986), new BigNumber((_121_v49).length));
      (_121_v49)[_index5] = _dafny.Seq.Concat(_58_v5, _module.__default.fm27(_60_globalState));
      let _126_v54;
      let _nw13 = new _module.C11();
      _nw13.__ctor();
      _126_v54 = _nw13;
      let _127_v55;
      _127_v55 = _dafny.Seq.of(_126_v54);
      let _128_v56;
      _128_v56 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_127_v55);
      if ((_58_v5)[_module.__default.safeIndex(new BigNumber(((((_128_v56).contains(_53_v0)) ? ((_128_v56).get(_53_v0)) : (_127_v55))).length), new BigNumber((_58_v5).length))]) {
        (_60_globalState).f15 = (_123_v51).f22;
        let _129_v57;
        _129_v57 = _module.D1.create_DC3(_122_v50, _53_v0);
        let _source4 = _129_v57;
        if (_source4.is_DC3) {
          let _130___mcc_h5 = (_source4).cf3;
          let _131___mcc_h6 = (_source4).cf4;
          let _132_cf4 = _131___mcc_h6;
          let _133_cf3 = _130___mcc_h5;
          (_60_globalState).f13 = ((_123_v51).f22).plus(new BigNumber(604));
          let _134_v58;
          let _init1 = ((_135_v51) => function (_136_i5) {
            return (_136_i5).minus((_135_v51).f22);
          })(_123_v51);
          let _nw14 = Array((new BigNumber(6)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
            _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _134_v58 = _nw14;
          let _137_v59;
          _137_v59 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_134_v58);
          (_126_v54).m1(_module.__default.fm19(_132_cf4, _53_v0, _60_globalState), (_137_v59).Merge(_137_v59), _60_globalState);
          (_60_globalState).f13 = (_123_v51).f22;
          let _138_v60;
          _138_v60 = new _dafny.CodePoint('w'.codePointAt(0));
          let _139_v61;
          _139_v61 = _dafny.MultiSet.fromElements(_138_v60);
          _139_v61 = _139_v61;
        } else if (_source4.is_DC4) {
          let _140___mcc_h7 = (_source4).cf5;
          let _141_cf5 = _140___mcc_h7;
          let _142_v62;
          _142_v62 = _dafny.Map.Empty.slice().updateUnsafe(_126_v54,(!(true)) && (_53_v0));
          _53_v0 = (((_142_v62).contains(((_53_v0) ? (_126_v54) : (_126_v54)))) ? ((_142_v62).get(((_53_v0) ? (_126_v54) : (_126_v54)))) : (_53_v0));
          _54_v1 = _module.__default.safeModuloInt(_54_v1, new BigNumber((((_53_v0) ? (_dafny.Seq.UnicodeFromString("cbcyxwik")) : (_122_v50))).length));
          let _143_v63;
          _143_v63 = new _dafny.CodePoint('r'.codePointAt(0));
          let _144_v64;
          let _init2 = ((_145_v0) => function (_146_i6) {
            return _145_v0;
          })(_53_v0);
          let _nw15 = Array((new BigNumber(18)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw15.length); _i0_2++) {
            _nw15[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _144_v64 = _nw15;
          let _index6 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_144_v64).length));
          (_144_v64)[_index6] = _53_v0;
          let _147_v65;
          let _init3 = ((_148_v1) => function (_149_i7) {
            return _module.__default.safeModuloInt(_149_i7, _148_v1);
          })(_54_v1);
          let _nw16 = Array((new BigNumber(18)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw16.length); _i0_3++) {
            _nw16[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _147_v65 = _nw16;
          let _150_v66;
          _150_v66 = _module.D2.create_DC8(_module.D2.create_DC5(_147_v65));
          let _151_v67;
          _151_v67 = _module.D2.create_DC5(_147_v65);
          let _152_v68;
          _152_v68 = _module.D2.create_DC8(_151_v67);
          let _pat_let_tv0 = _152_v68;
          let _153_v69;
          let _nw17 = new _module.C3();
          _nw17.__ctor(function (_pat_let0_0) {
            return function (_154_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_155_dt__update_hcf11_h0) {
                  return _module.D2.create_DC8(_155_dt__update_hcf11_h0);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_150_v66), _53_v0);
          _153_v69 = _nw17;
          let _156_v70;
          _156_v70 = _dafny.Map.Empty.slice().updateUnsafe(_153_v69,(_123_v51).f22);
          let _157_v71;
          _157_v71 = _dafny.Set.fromElements(_53_v0, _53_v0);
          let _158_v72;
          _158_v72 = _dafny.Seq.of(_157_v71);
          let _159_v73;
          let _nw18 = Array((new BigNumber(22)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = (new BigNumber((_156_v70).length)).plus((_dafny.ZERO).minus((_153_v69).fm2((_123_v51).fm3(_153_v69.f25, _60_globalState), _60_globalState)));
          _nw18[(_dafny.ONE).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(2)).toNumber()] = new BigNumber(801);
          _nw18[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jqh"), _dafny.Seq.UnicodeFromString("usr"))).length);
          _nw18[(new BigNumber(4)).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(5)).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(6)).toNumber()] = new BigNumber(152);
          _nw18[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_158_v72).length)).multipliedBy(_54_v1));
          _nw18[(new BigNumber(8)).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(9)).toNumber()] = (new BigNumber((_122_v50).length)).multipliedBy(new BigNumber(-562));
          _nw18[(new BigNumber(10)).toNumber()] = (new BigNumber(((_123_v51).f23).length)).multipliedBy(new BigNumber(-945));
          _nw18[(new BigNumber(11)).toNumber()] = ((_53_v0) ? ((_123_v51).f22) : (_54_v1));
          _nw18[(new BigNumber(12)).toNumber()] = (_116_v44)[_module.__default.safeIndex(new BigNumber((_55_v2).length), new BigNumber((_116_v44).length))];
          _nw18[(new BigNumber(13)).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(14)).toNumber()] = new BigNumber(798);
          _nw18[(new BigNumber(15)).toNumber()] = new BigNumber(-850);
          _nw18[(new BigNumber(16)).toNumber()] = (_123_v51).f22;
          _nw18[(new BigNumber(17)).toNumber()] = ((_123_v51).f22).multipliedBy((_123_v51).f22);
          _nw18[(new BigNumber(18)).toNumber()] = new BigNumber(495);
          _nw18[(new BigNumber(19)).toNumber()] = (_116_v44)[_module.__default.safeIndex((_123_v51).f22, new BigNumber((_116_v44).length))];
          _nw18[(new BigNumber(20)).toNumber()] = _54_v1;
          _nw18[(new BigNumber(21)).toNumber()] = new BigNumber(-503);
          _159_v73 = _nw18;
          let _index7 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_147_v65).length));
          (_147_v65)[_index7] = (_123_v51).f22;
          let _index8 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_144_v64).length));
          let _index9 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_147_v65).length));
          let _rhs13 = _122_v50;
          let _rhs14 = _143_v63;
          let _rhs15 = true;
          let _rhs16 = _153_v69.f25;
          let _rhs17 = _54_v1;
          let _lhs7 = _144_v64;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_144_v64).length));
          let _lhs9 = _147_v65;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_147_v65).length));
          _122_v50 = _rhs13;
          _143_v63 = _rhs14;
          _53_v0 = _rhs15;
          _lhs7[_lhs8] = _rhs16;
          _lhs9[_lhs10] = _rhs17;
          _57_v4 = (_57_v4).update(true, (_144_v64)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_144_v64).length))]);
        } else {
          let _160___mcc_h8 = (_source4).cf2;
          let _161_cf2 = _160___mcc_h8;
          (_60_globalState).f15 = _161_cf2;
          let _162_v74;
          let _nw19 = new _module.C13();
          _nw19.__ctor((_123_v51).f22, !((_88_v28).fm16(new BigNumber(366), _60_globalState)));
          _162_v74 = _nw19;
          let _163_v75;
          _163_v75 = _dafny.MultiSet.fromElements(new BigNumber(((_123_v51).f23).length), _54_v1);
          let _rhs18 = _162_v74;
          let _rhs19 = false;
          let _rhs20 = _module.__default.safeModuloInt(new BigNumber((_163_v75).cardinality()), (_162_v74).f20);
          let _lhs11 = _60_globalState;
          _162_v74 = _rhs18;
          _53_v0 = _rhs19;
          _lhs11.f13 = _rhs20;
          (_123_v51).m2((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_116_v44, _116_v44)).length)), _53_v0, new BigNumber(-853), _60_globalState);
          (_162_v74).f21 = false;
        }
        if (_53_v0) {
          let _164_v76;
          let _nw20 = Array((new BigNumber(3)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = new BigNumber(259);
          _nw20[(_dafny.ONE).toNumber()] = new BigNumber(((_123_v51).f23).length);
          _nw20[(new BigNumber(2)).toNumber()] = (_123_v51).f22;
          _164_v76 = _nw20;
          _164_v76 = _164_v76;
          _55_v2 = (_55_v2).Merge(_55_v2);
          let _165_v77;
          _165_v77 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_164_v76);
          let _166_v78;
          _166_v78 = _module.D23.create_DC63(_165_v77);
          (_126_v54).m1((_123_v51).f23, (_165_v77).Merge((_166_v78).dtor_cf106), _60_globalState);
          let _167_v79;
          _167_v79 = _module.D15.create_DC39(_57_v4, !(_53_v0), new BigNumber((_dafny.MultiSet.fromElements(_54_v1)).cardinality()));
          let _168_v80;
          _168_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(985),_167_v79);
          _168_v80 = _168_v80;
          _164_v76 = _164_v76;
        } else {
          let _169_v81;
          _169_v81 = new _dafny.CodePoint('p'.codePointAt(0));
          _169_v81 = _169_v81;
          let _170_v82;
          _170_v82 = _dafny.Set.fromElements(_53_v0);
          let _171_v83;
          _171_v83 = _module.D12.create_DC32(_170_v82);
          _171_v83 = _module.D12.create_DC32((_170_v82).Union(_170_v82));
          let _172_v84;
          let _nw21 = Array((new BigNumber(14)).toNumber()).fill(false);
          _172_v84 = _nw21;
          let _173_v85;
          _173_v85 = _module.D0.create_DC0(_172_v84);
          let _174_v86;
          _174_v86 = _dafny.Seq.of(_173_v85);
          (_60_globalState).f4 = (new BigNumber((((_53_v0) ? (_174_v86) : (_dafny.Seq.of(_173_v85)))).length)).multipliedBy(_54_v1);
          (_60_globalState).f4 = new BigNumber(857);
          let _175_v87;
          let _nw22 = new _module.C2();
          _nw22.__ctor();
          _175_v87 = _nw22;
          let _176_v88;
          _176_v88 = _dafny.Seq.of(_175_v87, _175_v87);
          (_123_v51).m2(new BigNumber((_dafny.Seq.Concat(_176_v88, _176_v88)).length), _53_v0, (_123_v51).f22, _60_globalState);
        }
        (_60_globalState).f13 = (_dafny.ZERO).minus(_54_v1);
        let _177_v89;
        let _nw23 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _177_v89 = _nw23;
        let _index10 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_177_v89).length));
        (_177_v89)[_index10] = (((_55_v2).contains(!(_53_v0))) ? ((_55_v2).get(!(_53_v0))) : ((_123_v51).f22));
        let _index11 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_177_v89).length));
        (_177_v89)[_index11] = new BigNumber(((_123_v51).f23).length);
      } else {
        let _178_v90;
        _178_v90 = _dafny.Set.fromElements(_53_v0);
        if ((_module.__default.fm37(_60_globalState)).IsSubsetOf(_178_v90)) {
          let _179_v91;
          let _nw24 = new _module.C9();
          _nw24.__ctor();
          _179_v91 = _nw24;
          let _180_v93;
          _180_v93 = _dafny.Map.Empty.slice().updateUnsafe(_179_v91,_dafny.MultiSet.fromElements(_56_v3, function () {
            let _coll19 = new _dafny.Set();
            for (const _compr_19 of _dafny.IntegerRange(new BigNumber(-884), new BigNumber(510))) {
              let _181_v92 = _compr_19;
              if (((new BigNumber(-884)).isLessThanOrEqualTo(_181_v92)) && ((_181_v92).isLessThan(new BigNumber(510)))) {
                _coll19.add((_181_v92).plus(new BigNumber(-656)));
              }
            }
            return _coll19;
          }()));
          _180_v93 = (_180_v93).update(_179_v91, _dafny.MultiSet.fromElements(_56_v3));
          let _182_v94;
          let _nw25 = new _module.C7();
          _nw25.__ctor();
          _182_v94 = _nw25;
          let _183_v95;
          _183_v95 = _dafny.Map.Empty.slice().updateUnsafe(_182_v94,_54_v1);
          (_60_globalState).f15 = (((_183_v95).contains(_182_v94)) ? ((_183_v95).get(_182_v94)) : ((_123_v51).f22));
          let _184_v96;
          _184_v96 = new _dafny.CodePoint('c'.codePointAt(0));
          let _185_v97;
          _185_v97 = _dafny.Map.Empty.slice().updateUnsafe((_123_v51).f22,_184_v96);
          let _186_v98;
          _186_v98 = _module.D19.create_DC52(new BigNumber(((_185_v97).Merge(_185_v97)).length), _53_v0, _53_v0);
          let _187_v99;
          _187_v99 = _dafny.Map.Empty.slice().updateUnsafe(_54_v1,_53_v0);
          _186_v98 = _module.__default.fm60(_187_v99, (_dafny.ZERO).minus(new BigNumber(((_123_v51).f23).length)), _60_globalState);
          (_60_globalState).f15 = ((_123_v51).f22).plus((_123_v51).f22);
          let _188_v100;
          _188_v100 = _dafny.Seq.of((_57_v4).Merge(_57_v4), _57_v4, _57_v4);
          let _189_v101;
          _189_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(896),_dafny.Map.Empty.slice().updateUnsafe(_53_v0,_53_v0));
          let _rhs21 = (_123_v51).f22;
          let _rhs22 = _dafny.Seq.Concat(_188_v100, _dafny.Seq.Concat(_188_v100, _dafny.Seq.of((((_189_v101).contains((_123_v51).f22)) ? ((_189_v101).get((_123_v51).f22)) : (_57_v4)), _57_v4)));
          let _rhs23 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(466), _54_v1), new BigNumber(-194));
          let _lhs12 = _60_globalState;
          let _lhs13 = _60_globalState;
          _lhs12.f4 = _rhs21;
          _188_v100 = _rhs22;
          _lhs13.f7 = _rhs23;
        } else {
          let _190_v102;
          _190_v102 = _dafny.MultiSet.fromElements(_53_v0, _53_v0, _53_v0);
          let _191_v103;
          _191_v103 = _dafny.Map.Empty.slice().updateUnsafe(false,_190_v102);
          let _192_v104;
          let _nw26 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _192_v104 = _nw26;
          let _193_v105;
          _193_v105 = _module.D2.create_DC5(_192_v104);
          (_61_v7).m6(_54_v1, ((((_191_v103).contains(true)) ? ((_191_v103).get(true)) : (_190_v102))).update(!(_53_v0), _module.__default.abs(new BigNumber((_module.__default.fm32((_123_v51).f22, (_123_v51).f22, _dafny.Set.fromElements(_53_v0, (((_57_v4).contains(true)) ? ((_57_v4).get(true)) : (_53_v0)), _53_v0), _60_globalState)).length))), _53_v0, _193_v105, _60_globalState);
          let _194_v106;
          _194_v106 = _dafny.MultiSet.fromElements((_123_v51).f22);
          let _195_v107;
          _195_v107 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_116_v44).length),_194_v106);
          (_123_v51).m2(new BigNumber((_195_v107).length), _53_v0, _54_v1, _60_globalState);
          let _196_v108;
          _196_v108 = _module.D7.create_DC21(new BigNumber(((_123_v51).f23).length), new BigNumber(-681), _53_v0, _53_v0, _53_v0);
          (_123_v51).m2((_123_v51).f22, _53_v0, (_123_v51).fm2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_196_v108).dtor_cf32),(_123_v51).f22)).length), _60_globalState), _60_globalState);
          let _197_v109;
          _197_v109 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_57_v4);
          let _198_v110;
          _198_v110 = _module.D14.create_DC36(_dafny.Seq.of(_53_v0, _53_v0, _53_v0, (_123_v51).fm1(_197_v109, new BigNumber(79), _60_globalState)));
          let _199_v111;
          _199_v111 = _dafny.Set.fromElements(_116_v44, _116_v44, _116_v44, _116_v44);
          let _rhs24 = _module.__default.safeModuloInt(_54_v1, (_123_v51).f22);
          let _rhs25 = _198_v110;
          let _rhs26 = (_199_v111).IsProperSubsetOf(_199_v111);
          let _rhs27 = ((_55_v2).equals(_55_v2)) === ((_59_v6).IsProperSubsetOf(_59_v6));
          let _lhs14 = _60_globalState;
          _lhs14.f7 = _rhs24;
          _198_v110 = _rhs25;
          _53_v0 = _rhs26;
          _53_v0 = _rhs27;
          (_60_globalState).f4 = (_dafny.ZERO).minus((_123_v51).f22);
        }
        (_60_globalState).f15 = ((_123_v51).f22).plus(new BigNumber((_178_v90).length));
        _53_v0 = (_178_v90).IsProperSubsetOf(_178_v90);
        let _200_v112;
        let _201_v113;
        let _202_v114;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = (_88_v28).m9(_module.__default.safeDivisionInt((_123_v51).f22, (_123_v51).f22), false, _53_v0, _60_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _200_v112 = _out3;
        _201_v113 = _out4;
        _202_v114 = _out5;
        _202_v114 = _dafny.Seq.IsProperPrefixOf(_116_v44, _dafny.Seq.Concat(_116_v44, _116_v44));
      }
      let _203_v115;
      let _nw27 = new _module.C5();
      _nw27.__ctor();
      _203_v115 = _nw27;
      let _204_v116;
      _204_v116 = _dafny.Set.fromElements(_203_v115, _203_v115);
      let _205_v117;
      let _nw28 = new _module.C1();
      _nw28.__ctor();
      _205_v117 = _nw28;
      let _206_v118;
      let _nw29 = new _module.C8();
      _nw29.__ctor();
      _206_v118 = _nw29;
      let _207_v119;
      _207_v119 = _dafny.Map.Empty.slice().updateUnsafe(_205_v117,_206_v118);
      let _208_v120;
      let _nw30 = Array((new BigNumber(13)).toNumber());
      _nw30[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_123_v51).f22);
      _nw30[(_dafny.ONE).toNumber()] = (_123_v51).f22;
      _nw30[(new BigNumber(2)).toNumber()] = (_54_v1).minus(new BigNumber((_module.__default.fm5(_60_globalState)).length));
      _nw30[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.Concat((_121_v49)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_121_v49).length))], (_121_v49)[_module.__default.safeIndex(new BigNumber(986), new BigNumber((_121_v49).length))])).length);
      _nw30[(new BigNumber(4)).toNumber()] = _54_v1;
      _nw30[(new BigNumber(5)).toNumber()] = _54_v1;
      _nw30[(new BigNumber(6)).toNumber()] = ((_123_v51).f22).minus((_123_v51).f22);
      _nw30[(new BigNumber(7)).toNumber()] = _54_v1;
      _nw30[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(_54_v1, _54_v1);
      _nw30[(new BigNumber(9)).toNumber()] = _54_v1;
      _nw30[(new BigNumber(10)).toNumber()] = new BigNumber((_204_v116).length);
      _nw30[(new BigNumber(11)).toNumber()] = _54_v1;
      _nw30[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((_207_v119).update(_205_v117, _206_v118)).length), new BigNumber(382));
      _208_v120 = _nw30;
      let _index12 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_208_v120).length));
      (_208_v120)[_index12] = _54_v1;
      let _index13 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_208_v120).length));
      (_208_v120)[_index13] = _module.__default.safeModuloInt((_123_v51).f22, (_123_v51).f22);
      let _209_v121;
      _209_v121 = new _dafny.CodePoint('i'.codePointAt(0));
      let _210_v122;
      _210_v122 = _module.D6.create_DC16(_116_v44);
      let _211_v123;
      _211_v123 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm39((_116_v44)[_module.__default.safeIndex(_54_v1, new BigNumber((_116_v44).length))], _209_v121, (_210_v122).dtor_cf23, _53_v0, _60_globalState),_dafny.Seq.update((_123_v51).f23, _module.__default.safeIndex(_54_v1, new BigNumber(((_123_v51).f23).length)), _209_v121));
      _211_v123 = (_211_v123).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(278)), ((_212_v51) => function (_213_i8) {
        return new BigNumber(((_212_v51).f23).length);
      })(_123_v51)), _dafny.Seq.Concat((_123_v51).f23, (_123_v51).f23));
      let _214_v124;
      _214_v124 = _dafny.Map.Empty.slice().updateUnsafe(_53_v0,_122_v50);
      let _rhs28 = _dafny.Seq.Concat((((_214_v124).contains(_53_v0)) ? ((_214_v124).get(_53_v0)) : (_122_v50)), (_123_v51).f23);
      let _rhs29 = (_123_v51).f22;
      let _lhs15 = _60_globalState;
      _122_v50 = _rhs28;
      _lhs15.f15 = _rhs29;
      _55_v2 = ((_55_v2).Merge(_55_v2)).Merge(_55_v2);
      process.stdout.write(_dafny.toString(_53_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_54_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_55_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_56_v3).equals(_dafny.Set.fromElements(new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_58_v5, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_59_v6).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_60_globalState).f0, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_60_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_globalState).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_60_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_globalState).f8).equals(_dafny.Set.fromElements(new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_globalState).f11).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false).updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_60_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_60_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_globalState).f16).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_60_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_89_v29).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v30).dtor_cf85));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v30).dtor_cf86));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v30).dtor_cf87));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v30).dtor_cf88));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_114_v42, _dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC29()), _module.D11.create_DC31(_module.D11.create_DC29())))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_115_v43).dtor_cf80, _dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC29()), _module.D11.create_DC31(_module.D11.create_DC29())))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_116_v44, _dafny.Seq.of(new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_118_v46).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(215)), _dafny.Set.fromElements(new BigNumber(215))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_119_v47).dtor_cf22).dtor_cf17).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(215)), _dafny.Set.fromElements(new BigNumber(215))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_120_v48, _dafny.Seq.of(_module.D5.create_DC15(_module.D5.create_DC13(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(215)), _dafny.Set.fromElements(new BigNumber(215))))), _module.D5.create_DC15(_module.D5.create_DC13(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(215)), _dafny.Set.fromElements(new BigNumber(215))))), _module.D5.create_DC15(_module.D5.create_DC13(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(215)), _dafny.Set.fromElements(new BigNumber(215)))))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[_dafny.ZERO], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[_dafny.ONE], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(2)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(3)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(4)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(5)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(6)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(7)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(8)], _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v49)[new BigNumber(9)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_122_v50).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_v51).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_123_v51).f23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_124_v52)[_dafny.ZERO]).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_124_v52)[_dafny.ZERO]).f23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_124_v52)[_dafny.ONE]).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_124_v52)[_dafny.ONE]).f23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_124_v52)[new BigNumber(2)]).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_124_v52)[new BigNumber(2)]).f23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_124_v52)[new BigNumber(3)]).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_124_v52)[new BigNumber(3)]).f23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_125_v53).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_127_v55).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_128_v56).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_204_v116).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_207_v119).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v120)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_209_v121));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_210_v122).dtor_cf23, _dafny.Seq.of(new BigNumber(215)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_211_v123).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(669), new BigNumber(2), _dafny.ONE, new BigNumber(412), _dafny.ONE, new BigNumber(193), _dafny.ONE, new BigNumber(2)),_dafny.Seq.UnicodeFromString("xofwgic")).updateUnsafe(_dafny.Seq.of(new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7), new BigNumber(7)),_dafny.Seq.UnicodeFromString("xofwgscxofwgsc")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v124).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("ipyktu")))));
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
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([]);
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
    static create_DC3(cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D1(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + this.cf3.toVerbatimString(true) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC7(cf8, cf9, cf10) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D2(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC8(cf11) {
      let $dt = new D2(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf6 === other.cf6;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC6(_dafny.ZERO);
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
    static create_DC10(cf13, cf14) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC9(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC12(cf16) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC11(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(_dafny.ZERO);
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
    static create_DC14(cf18, cf19, cf20, cf21) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC13(cf17) {
      let $dt = new D5(1);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC15(cf22) {
      let $dt = new D5(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO, false);
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
    static create_DC17(cf24, cf25, cf26) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC16(cf23) {
      let $dt = new D6(1);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC18(cf27) {
      let $dt = new D6(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, false, false);
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
    static create_DC20(cf29) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D7(1);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC22(cf35, cf36, cf37, cf38) {
      let $dt = new D7(2);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D7(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC22" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32 && this.cf33 === other.cf33 && this.cf34 === other.cf34;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC20(_dafny.ZERO);
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
    static create_DC24(cf40, cf41) {
      let $dt = new D8(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC25(cf42) {
      let $dt = new D8(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC23(cf39) {
      let $dt = new D8(2);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(_dafny.ZERO, false);
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
    static create_DC26(cf43) {
      let $dt = new D9(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43;
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf44) {
      let $dt = new D10(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf44) + ")";
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
      return _dafny.Map.Empty;
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
    static create_DC29() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC30(cf46, cf47) {
      let $dt = new D11(1);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC28(cf45) {
      let $dt = new D11(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC31(cf48) {
      let $dt = new D11(3);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC29";
      } else if (this.$tag === 1) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf48) + ")";
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
        return other.$tag === 1 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf48, other.cf48);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC29();
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
    static create_DC33(cf50) {
      let $dt = new D12(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC32(cf49) {
      let $dt = new D12(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC33(false);
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
    static create_DC34(cf51) {
      let $dt = new D13(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC36(cf53) {
      let $dt = new D14(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC35(cf52) {
      let $dt = new D14(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC37(cf54) {
      let $dt = new D14(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC36(_dafny.Seq.of());
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
    static create_DC39(cf56, cf57, cf58) {
      let $dt = new D15(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC38(cf55) {
      let $dt = new D15(1);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC40(cf59) {
      let $dt = new D15(2);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC39(_dafny.Map.Empty, false, _dafny.ZERO);
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
    static create_DC42(cf61, cf62, cf63) {
      let $dt = new D16(0);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC43(cf64, cf65) {
      let $dt = new D16(1);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC41(cf60) {
      let $dt = new D16(2);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC44(cf66) {
      let $dt = new D16(3);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get is_DC44() { return this.$tag === 3; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC43" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf61 === other.cf61 && this.cf62 === other.cf62 && this.cf63 === other.cf63;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC42(false, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf68) {
      let $dt = new D17(0);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC47(cf69, cf70, cf71, cf72, cf73) {
      let $dt = new D17(1);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC45(cf67) {
      let $dt = new D17(2);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf69 === other.cf69 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71) && _dafny.areEqual(this.cf72, other.cf72) && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC46(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49(cf75, cf76, cf77, cf78) {
      let $dt = new D18(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC48(cf74) {
      let $dt = new D18(1);
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC50(cf79) {
      let $dt = new D18(2);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC49" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC48" + "(" + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75 && this.cf76 === other.cf76 && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC49(false, false, new _dafny.CodePoint('D'.codePointAt(0)), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf81, cf82, cf83) {
      let $dt = new D19(0);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC53() {
      let $dt = new D19(1);
      return $dt;
    }
    static create_DC51(cf80) {
      let $dt = new D19(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC52" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC53";
      } else if (this.$tag === 2) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81) && this.cf82 === other.cf82 && this.cf83 === other.cf83;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC52(_dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf85, cf86, cf87, cf88) {
      let $dt = new D20(0);
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC54(cf84) {
      let $dt = new D20(1);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC55" + "(" + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC54" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf85, other.cf85) && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87 && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf84 === other.cf84;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC55(_dafny.ZERO, _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf90, cf91, cf92) {
      let $dt = new D21(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC56(cf89) {
      let $dt = new D21(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC58(cf93) {
      let $dt = new D21(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get is_DC58() { return this.$tag === 2; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC57" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC56" + "(" + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC58" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf90, other.cf90) && _dafny.areEqual(this.cf91, other.cf91) && _dafny.areEqual(this.cf92, other.cf92);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf89 === other.cf89;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC57(_dafny.ZERO, _dafny.ZERO, _module.D5.Default());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC60(cf95, cf96) {
      let $dt = new D22(0);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC61(cf97, cf98, cf99, cf100) {
      let $dt = new D22(1);
      $dt.cf97 = cf97;
      $dt.cf98 = cf98;
      $dt.cf99 = cf99;
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC62(cf101, cf102, cf103, cf104, cf105) {
      let $dt = new D22(2);
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      $dt.cf104 = cf104;
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC59(cf94) {
      let $dt = new D22(3);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get is_DC59() { return this.$tag === 3; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC60" + "(" + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC61" + "(" + _dafny.toString(this.cf97) + ", " + _dafny.toString(this.cf98) + ", " + _dafny.toString(this.cf99) + ", " + this.cf100.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D22.DC62" + "(" + this.cf101.toVerbatimString(true) + ", " + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ", " + _dafny.toString(this.cf104) + ", " + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC59" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf97, other.cf97) && _dafny.areEqual(this.cf98, other.cf98) && _dafny.areEqual(this.cf99, other.cf99) && _dafny.areEqual(this.cf100, other.cf100);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf101, other.cf101) && _dafny.areEqual(this.cf102, other.cf102) && _dafny.areEqual(this.cf103, other.cf103) && _dafny.areEqual(this.cf104, other.cf104) && this.cf105 === other.cf105;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf94, other.cf94);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC60(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC64(cf107, cf108) {
      let $dt = new D23(0);
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      return $dt;
    }
    static create_DC63(cf106) {
      let $dt = new D23(1);
      $dt.cf106 = cf106;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf106() { return this.cf106; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC64" + "(" + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC63" + "(" + _dafny.toString(this.cf106) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf107, other.cf107) && _dafny.areEqual(this.cf108, other.cf108);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf106, other.cf106);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC64(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
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
      this.f4 = _dafny.ZERO;
      this.f7 = _dafny.ZERO;
      this.f13 = _dafny.ZERO;
      this.f15 = _dafny.ZERO;
      this._f0 = _dafny.Seq.of();
      this._f1 = false;
      this._f2 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f3 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = _dafny.Map.Empty;
      this._f8 = _dafny.Set.Empty;
      this._f9 = false;
      this._f10 = false;
      this._f11 = _dafny.Map.Empty;
      this._f12 = _dafny.ZERO;
      this._f14 = false;
      this._f16 = _dafny.MultiSet.Empty;
      this._f17 = false;
      this._f18 = _dafny.ZERO;
      this._f19 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
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
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
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
    fm28(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(852)), function (_215_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("o")), _dafny.Seq.UnicodeFromString("fdqgbyeq"));
    };
    fm29(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D5.create_DC14(new _dafny.CodePoint('n'.codePointAt(0)), false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D7.create_DC21(new BigNumber(739), new BigNumber((function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sosvg")).length),new BigNumber(968))).Keys.Elements) {
    let _216_v0 = _compr_20;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("sosvg")).length),new BigNumber(968))).contains(_216_v0)) {
      _coll20.push([_module.__default.safeModuloInt(_216_v0, new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality())),new BigNumber(-322)]);
    }
  }
  return _coll20;
}()).length), true, !(true), false))).length),new BigNumber(206))).length), !(true))).dtor_cf21;
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
    fm25(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(3), _module.__default.safeDivisionInt(new BigNumber(-692), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_217_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length)));
    };
    m13(p0, p1, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let r1 = _dafny.ZERO;
      let r2 = _module.D7.Default();
      let r3 = _module.D1.Default();
      let _218_v0;
      _218_v0 = true;
      let _219_v1;
      let _init4 = ((_220_p1) => function (_221_i0) {
        return (_221_i0).multipliedBy(_220_p1);
      })(p1);
      let _nw31 = Array((new BigNumber(21)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw31.length); _i0_4++) {
        _nw31[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _219_v1 = _nw31;
      let _index14 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
      (_219_v1)[_index14] = _module.__default.safeDivisionInt(p1, p1);
      let _222_v2;
      _222_v2 = _dafny.Set.fromElements(p1);
      let _index15 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
      let _rhs30 = (_222_v2).IsSubsetOf(_222_v2);
      let _rhs31 = new BigNumber(38);
      let _lhs16 = _219_v1;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
      _218_v0 = _rhs30;
      _lhs16[_lhs17] = _rhs31;
      let _223_v3;
      _223_v3 = _dafny.Seq.of((_dafny.ZERO).minus((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]));
      let _224_v4;
      _224_v4 = _module.D6.create_DC16(_223_v3);
      let _225_v5;
      _225_v5 = _dafny.Map.Empty.slice().updateUnsafe(_218_v0,_218_v0);
      let _index16 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
      (_219_v1)[_index16] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray((_224_v4).dtor_cf23)).cardinality())), ((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]).multipliedBy(new BigNumber((_225_v5).length)));
      let _226_i1;
      _226_i1 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm26(globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_226_i1)) {
              break L0;
            }
            _226_i1 = (_226_i1).plus(_dafny.ONE);
            let _227_v6;
            let _init5 = ((_228_v0) => function (_229_i2) {
              return _228_v0;
            })(_218_v0);
            let _nw32 = Array((new BigNumber(23)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw32.length); _i0_5++) {
              _nw32[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _227_v6 = _nw32;
            let _230_v7;
            _230_v7 = _dafny.Seq.of(_227_v6, _227_v6);
            let _source5 = _module.D7.create_DC19(_dafny.Seq.Concat(_230_v7, _dafny.Seq.update(_230_v7, _module.__default.safeIndex(p1, new BigNumber((_230_v7).length)), _227_v6)));
            if (_source5.is_DC20) {
              let _231___mcc_h0 = (_source5).cf29;
              let _232_cf29 = _231___mcc_h0;
              let _233_v8;
              _233_v8 = _dafny.Seq.UnicodeFromString("slrsv");
              let _234_v9;
              _234_v9 = new _dafny.CodePoint('c'.codePointAt(0));
              let _235_v10;
              _235_v10 = _dafny.Map.Empty.slice().updateUnsafe(!(!(!_dafny.areEqual(_233_v8, _dafny.Seq.update(_233_v8, _module.__default.safeIndex((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))], new BigNumber((_233_v8).length)), _234_v9)))),p1);
              _235_v10 = (_235_v10).update(_218_v0, p1);
              let _236_v11;
              _236_v11 = _dafny.MultiSet.fromElements(_218_v0);
              let _237_v12;
              _237_v12 = _dafny.Seq.of(_218_v0);
              _218_v0 = !((_232_cf29).isLessThan(new BigNumber(((_236_v11).Difference(_dafny.MultiSet.FromArray(_237_v12))).cardinality())));
              _218_v0 = _module.__default.fm26(globalState);
              let _index17 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_227_v6).length));
              (_227_v6)[_index17] = _218_v0;
              let _index18 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_227_v6).length));
              (_227_v6)[_index18] = _218_v0;
            } else if (_source5.is_DC21) {
              let _238___mcc_h1 = (_source5).cf30;
              let _239___mcc_h2 = (_source5).cf31;
              let _240___mcc_h3 = (_source5).cf32;
              let _241___mcc_h4 = (_source5).cf33;
              let _242___mcc_h5 = (_source5).cf34;
              let _243_cf34 = _242___mcc_h5;
              let _244_cf33 = _241___mcc_h4;
              let _245_cf32 = _240___mcc_h3;
              let _246_cf31 = _239___mcc_h2;
              let _247_cf30 = _238___mcc_h1;
              _246_cf31 = _246_cf31;
              let _248_v13;
              _248_v13 = _dafny.Seq.of(_245_cf32);
              let _249_v14;
              let _nw33 = Array((new BigNumber(28)).toNumber());
              _nw33[(_dafny.ZERO).toNumber()] = _248_v13;
              _nw33[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_248_v13, _dafny.Seq.update(_248_v13, _module.__default.safeIndex(_246_cf31, new BigNumber((_248_v13).length)), _243_cf34));
              _nw33[(new BigNumber(2)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_248_v13, _module.__default.safeIndex(p1, new BigNumber((_248_v13).length)), !(_218_v0));
              _nw33[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_218_v0, !(!(_245_cf32)));
              _nw33[(new BigNumber(5)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(6)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(7)).toNumber()] = _module.__default.fm27(globalState);
              _nw33[(new BigNumber(8)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(9)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(10)).toNumber()] = _module.__default.fm27(globalState);
              _nw33[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_248_v13, _module.__default.safeIndex(_246_cf31, new BigNumber((_248_v13).length)), _module.__default.fm26(globalState));
              _nw33[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(true, _245_cf32, _218_v0, _243_cf34), _248_v13);
              _nw33[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_245_cf32);
              _nw33[(new BigNumber(14)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(15)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(16)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(17)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(18)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(19)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_243_cf34, _218_v0);
              _nw33[(new BigNumber(21)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(_244_cf33, _218_v0);
              _nw33[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_244_cf33, _243_cf34, false, true, _245_cf32);
              _nw33[(new BigNumber(24)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(25)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(26)).toNumber()] = _248_v13;
              _nw33[(new BigNumber(27)).toNumber()] = _248_v13;
              _249_v14 = _nw33;
              let _250_v15;
              _250_v15 = _249_v14;
              let _251_v16;
              _251_v16 = _dafny.Seq.of(_dafny.Seq.update(_223_v3, _module.__default.safeIndex(new BigNumber(959), new BigNumber((_223_v3).length)), p1));
              let _index19 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
              let _rhs32 = (p1).isLessThanOrEqualTo(_247_cf30);
              let _rhs33 = (_250_v15);
              let _rhs34 = new BigNumber((_251_v16).length);
              let _rhs35 = (_247_cf30).plus(p1);
              let _lhs18 = globalState;
              let _lhs19 = _219_v1;
              let _lhs20 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
              _244_cf33 = _rhs32;
              _249_v14 = _rhs33;
              _lhs18.f7 = _rhs34;
              _lhs19[_lhs20] = _rhs35;
              let _252_v17;
              _252_v17 = _dafny.MultiSet.fromElements(_247_cf30, (_dafny.ZERO).minus((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]));
              let _253_v18;
              _253_v18 = new _dafny.CodePoint('r'.codePointAt(0));
              let _254_v19;
              _254_v19 = _dafny.Seq.UnicodeFromString("fk");
              let _255_v20;
              _255_v20 = _dafny.Seq.of((_254_v19)[_module.__default.safeIndex(p1, new BigNumber((_254_v19).length))]);
              let _256_v21;
              _256_v21 = _dafny.Map.Empty.slice().updateUnsafe(_252_v17,_module.D5.create_DC14(_253_v18, (_248_v13)[_module.__default.safeIndex(p1, new BigNumber((_248_v13).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_254_v19).length))), _218_v0));
              _256_v21 = _256_v21;
              _250_v15 = _249_v14;
            } else if (_source5.is_DC22) {
              let _257___mcc_h6 = (_source5).cf35;
              let _258___mcc_h7 = (_source5).cf36;
              let _259___mcc_h8 = (_source5).cf37;
              let _260___mcc_h9 = (_source5).cf38;
              let _261_cf38 = _260___mcc_h9;
              let _262_cf37 = _259___mcc_h8;
              let _263_cf36 = _258___mcc_h7;
              let _264_cf35 = _257___mcc_h6;
              _264_cf35 = !(_module.__default.fm26(globalState));
              let _265_v22;
              _265_v22 = _dafny.Set.fromElements(_219_v1, _219_v1);
              let _index20 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
              (_219_v1)[_index20] = new BigNumber(((_265_v22).Intersect((_265_v22).Difference(_265_v22))).length);
              _264_cf35 = _264_cf35;
              _261_cf38 = (_dafny.ZERO).minus(new BigNumber(-589));
            } else {
              let _266___mcc_h10 = (_source5).cf28;
              let _267_cf28 = _266___mcc_h10;
              _218_v0 = _218_v0;
              _218_v0 = _218_v0;
              let _268_v23;
              let _nw34 = new _module.C0();
              _nw34.__ctor();
              _268_v23 = _nw34;
              let _269_v24;
              let _nw35 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _269_v24 = _nw35;
              let _270_v25;
              _270_v25 = _dafny.Seq.UnicodeFromString("yaw");
              let _index21 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_269_v24).length));
              (_269_v24)[_index21] = _270_v25;
              let _271_v26;
              _271_v26 = _dafny.MultiSet.fromElements(_219_v1);
              let _index22 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_269_v24).length));
              let _rhs36 = _dafny.Seq.Concat(_270_v25, _270_v25);
              let _rhs37 = (p1).minus((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
              let _rhs38 = !((_271_v26).IsSubsetOf(_271_v26));
              let _rhs39 = p1;
              let _rhs40 = !(_218_v0);
              let _lhs21 = _269_v24;
              let _lhs22 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_269_v24).length));
              let _lhs23 = globalState;
              let _lhs24 = globalState;
              _lhs21[_lhs22] = _rhs36;
              _lhs23.f7 = _rhs37;
              _218_v0 = _rhs38;
              _lhs24.f4 = _rhs39;
              _218_v0 = _rhs40;
            }
            _227_v6 = _227_v6;
            let _index23 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_227_v6).length));
            (_227_v6)[_index23] = _218_v0;
            let _272_v27;
            let _init6 = function (_273_i3) {
              return _dafny.Seq.of(false);
            };
            let _nw36 = Array((new BigNumber(13)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw36.length); _i0_6++) {
              _nw36[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _272_v27 = _nw36;
            let _274_v28;
            _274_v28 = _dafny.Seq.of(_272_v27);
            let _275_v29;
            _275_v29 = (_274_v28)[_module.__default.safeIndex(p1, new BigNumber((_274_v28).length))];
            let _276_v30;
            _276_v30 = _dafny.Set.fromElements(_275_v29);
            let _277_v31;
            _277_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_276_v30);
            let _index24 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_227_v6).length));
            (_227_v6)[_index24] = (_277_v31).contains(p1);
            let _278_v32;
            _278_v32 = _module.D8.create_DC25((_dafny.ZERO).minus(new BigNumber((_223_v3).length)));
            let _279_v33;
            _279_v33 = _dafny.Seq.UnicodeFromString("cd");
            let _280_v34;
            _280_v34 = _dafny.Seq.of(_218_v0, _218_v0, _218_v0);
            let _rhs41 = _278_v32;
            let _rhs42 = (_dafny.ZERO).minus(new BigNumber(-416));
            let _rhs43 = _219_v1;
            let _rhs44 = (_dafny.ZERO).minus(new BigNumber((_279_v33).length));
            let _rhs45 = (new BigNumber((_280_v34).length)).multipliedBy((p1).multipliedBy((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]));
            let _lhs25 = globalState;
            let _lhs26 = globalState;
            let _lhs27 = globalState;
            _278_v32 = _rhs41;
            _lhs25.f4 = _rhs42;
            _219_v1 = _rhs43;
            _lhs26.f13 = _rhs44;
            _lhs27.f4 = _rhs45;
          }
        }
      }
      let _281_v35;
      _281_v35 = _module.D2.create_DC5(_219_v1);
      let _source6 = _281_v35;
      if (_source6.is_DC6) {
        let _282___mcc_h11 = (_source6).cf7;
        let _283_cf7 = _282___mcc_h11;
        let _284_v36;
        _284_v36 = _dafny.Map.Empty.slice().updateUnsafe((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))],_218_v0);
        let _285_v37;
        _285_v37 = _dafny.Seq.of(!(_218_v0));
        _218_v0 = !(((_dafny.Map.Empty.slice().updateUnsafe(p1,_218_v0)).equals(_284_v36)) === ((_285_v37)[_module.__default.safeIndex(_283_cf7, new BigNumber((_285_v37).length))]));
        let _286_v38;
        _286_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm25(_218_v0, _218_v0, globalState),new _dafny.CodePoint('b'.codePointAt(0)));
        let _287_v39;
        let _nw37 = Array((new BigNumber(20)).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = _285_v37;
        _nw37[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_218_v0), _module.__default.safeIndex(new BigNumber((_286_v38).length), new BigNumber((_dafny.Seq.of(_218_v0)).length)), false);
        _nw37[(new BigNumber(2)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(3)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_218_v0);
        _nw37[(new BigNumber(5)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(6)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(7)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(8)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(9)).toNumber()] = _module.__default.fm27(globalState);
        _nw37[(new BigNumber(10)).toNumber()] = _dafny.Seq.of(_218_v0, _218_v0, _218_v0);
        _nw37[(new BigNumber(11)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(12)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(13)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_218_v0, _module.__default.fm26(globalState));
        _nw37[(new BigNumber(15)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(!(false), true);
        _nw37[(new BigNumber(17)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(18)).toNumber()] = _285_v37;
        _nw37[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_285_v37, _module.__default.safeIndex((_this).fm25(_218_v0, !(_218_v0), globalState), new BigNumber((_285_v37).length)), _218_v0);
        _287_v39 = _nw37;
        let _288_v40;
        _288_v40 = _287_v39;
        let _source7 = _288_v40;
        let _289___mcc_h17 = _source7;
        let _290_cf43 = _289___mcc_h17;
        let _291_v41;
        _291_v41 = _dafny.Seq.UnicodeFromString("owcdswdik");
        let _292_v42;
        _292_v42 = _dafny.Seq.of((_291_v41)[_module.__default.safeIndex((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))], new BigNumber((_291_v41).length))]);
        (globalState).f15 = (_dafny.ZERO).minus(new BigNumber((_291_v41).length));
        let _293_v43;
        _293_v43 = _module.D8.create_DC23(_285_v37);
        let _294_v44;
        _294_v44 = _dafny.Map.Empty.slice().updateUnsafe(_218_v0,_291_v41);
        let _295_v45;
        _295_v45 = _dafny.Map.Empty.slice().updateUnsafe((_294_v44).update(_218_v0, _291_v41),_218_v0);
        let _296_v46;
        _296_v46 = _dafny.MultiSet.fromElements(false, !(_218_v0), _218_v0, (((_295_v45).contains(_294_v44)) ? ((_295_v45).get(_294_v44)) : (_218_v0)));
        let _rhs46 = _293_v43;
        let _rhs47 = _296_v46;
        let _rhs48 = _218_v0;
        _293_v43 = _rhs46;
        _296_v46 = _rhs47;
        _218_v0 = _rhs48;
        (globalState).f7 = (_283_cf7).multipliedBy((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
        let _297_v47;
        _297_v47 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_283_cf7, new BigNumber((_291_v41).length)),_283_cf7);
        let _298_v48;
        _298_v48 = _dafny.MultiSet.fromElements((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
        let _299_v49;
        _299_v49 = _dafny.Set.fromElements(_218_v0);
        let _300_v50;
        _300_v50 = new _dafny.CodePoint('t'.codePointAt(0));
        let _301_v51;
        let _nw38 = Array((new BigNumber(11)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = (_299_v49).IsDisjointFrom(_299_v49);
        _nw38[(_dafny.ONE).toNumber()] = _218_v0;
        _nw38[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements(new BigNumber((_223_v3).length))).IsDisjointFrom(_222_v2);
        _nw38[(new BigNumber(3)).toNumber()] = true;
        _nw38[(new BigNumber(4)).toNumber()] = ((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]).isEqualTo((_this).fm25(false, _218_v0, globalState));
        _nw38[(new BigNumber(5)).toNumber()] = _218_v0;
        _nw38[(new BigNumber(6)).toNumber()] = _218_v0;
        _nw38[(new BigNumber(7)).toNumber()] = _218_v0;
        _nw38[(new BigNumber(8)).toNumber()] = _dafny.Seq.contains(_292_v42, _300_v50);
        _nw38[(new BigNumber(9)).toNumber()] = (!(_218_v0)) && (_218_v0);
        _nw38[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements(_218_v0)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_218_v0, _218_v0));
        _301_v51 = _nw38;
        let _index25 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_301_v51).length));
        (_301_v51)[_index25] = (((_284_v36).contains((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) ? ((_284_v36).get((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) : (true));
        let _index26 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_301_v51).length));
        let _rhs49 = _297_v47;
        let _rhs50 = _298_v48;
        let _rhs51 = !(_218_v0) || (_218_v0);
        let _lhs28 = _301_v51;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_301_v51).length));
        _297_v47 = _rhs49;
        _298_v48 = _rhs50;
        _lhs28[_lhs29] = _rhs51;
        let _302_v52;
        _302_v52 = new _dafny.CodePoint('a'.codePointAt(0));
        let _303_v53;
        _303_v53 = _dafny.Seq.UnicodeFromString("uljkewecy");
        if (_dafny.Seq.contains(_dafny.Seq.Concat(_303_v53, _303_v53), _302_v52)) {
          let _304_v54;
          let _nw39 = Array((new BigNumber(5)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = (((_284_v36).contains(p1)) ? ((_284_v36).get(p1)) : (_218_v0));
          _nw39[(_dafny.ONE).toNumber()] = _218_v0;
          _nw39[(new BigNumber(2)).toNumber()] = _218_v0;
          _nw39[(new BigNumber(3)).toNumber()] = _218_v0;
          _nw39[(new BigNumber(4)).toNumber()] = _218_v0;
          _304_v54 = _nw39;
          let _index27 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          (_304_v54)[_index27] = false;
          let _305_v55;
          let _init7 = ((_306_v0, _307_p1) => function (_308_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(!(_306_v0),_307_p1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_306_v0,_307_p1));
          })(_218_v0, p1);
          let _nw40 = Array((new BigNumber(3)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw40.length); _i0_7++) {
            _nw40[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _305_v55 = _nw40;
          let _index28 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _rhs52 = _218_v0;
          let _rhs53 = _305_v55;
          let _lhs30 = _304_v54;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          _lhs30[_lhs31] = _rhs52;
          _305_v55 = _rhs53;
          let _index29 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _rhs54 = _302_v52;
          let _rhs55 = !(!(_218_v0));
          let _lhs32 = _304_v54;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          _302_v52 = _rhs54;
          _lhs32[_lhs33] = _rhs55;
          let _index30 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _index31 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _index32 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _rhs56 = (_304_v54)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length))];
          let _rhs57 = _283_cf7;
          let _rhs58 = !(!(true) || (_218_v0)) || (_218_v0);
          let _rhs59 = _module.__default.fm26(globalState);
          let _rhs60 = (_304_v54)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length))];
          let _lhs34 = _304_v54;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _lhs36 = globalState;
          let _lhs37 = _304_v54;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          let _lhs39 = _304_v54;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length));
          _lhs34[_lhs35] = _rhs56;
          _lhs36.f7 = _rhs57;
          _lhs37[_lhs38] = _rhs58;
          _218_v0 = _rhs59;
          _lhs39[_lhs40] = _rhs60;
          let _309_v56;
          _309_v56 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.of((_304_v54)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length))]));
          r0 = (((_286_v38).contains((new BigNumber(879)).minus(new BigNumber(((((_309_v56).contains((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) ? ((_309_v56).get((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) : (_285_v37))).length)))) ? ((_286_v38).get((new BigNumber(879)).minus(new BigNumber(((((_309_v56).contains((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) ? ((_309_v56).get((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))])) : (_285_v37))).length)))) : (_302_v52));
          let _310_v57;
          _310_v57 = _dafny.Set.fromElements(_218_v0);
          let _311_v58;
          _311_v58 = _module.D7.create_DC22(_218_v0, _283_cf7, (_285_v37)[_module.__default.safeIndex(p1, new BigNumber((_285_v37).length))], (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
          let _312_v59;
          _312_v59 = _dafny.Seq.of(_311_v58, _311_v58);
          _303_v53 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_303_v53, _module.__default.safeIndex(new BigNumber((_310_v57).length), new BigNumber((_303_v53).length)), _302_v52), _module.__default.fm30(new BigNumber((_312_v59).length), _module.__default.fm26(globalState), function () {
            let _coll21 = new _dafny.Set();
            for (const _compr_21 of (_223_v3).Elements) {
              let _313_v60 = _compr_21;
              if (_dafny.Seq.contains(_223_v3, _313_v60)) {
                _coll21.add(_module.__default.safeDivisionInt(_313_v60, new BigNumber(22)));
              }
            }
            return _coll21;
          }(), (_304_v54)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_304_v54).length))], globalState)), _303_v53);
        } else {
          let _314_v61;
          let _nw41 = Array((new BigNumber(12)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _218_v0;
          _nw41[(_dafny.ONE).toNumber()] = true;
          _nw41[(new BigNumber(2)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(3)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(4)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(5)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(6)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(7)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(8)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(9)).toNumber()] = _218_v0;
          _nw41[(new BigNumber(10)).toNumber()] = true;
          _nw41[(new BigNumber(11)).toNumber()] = _218_v0;
          _314_v61 = _nw41;
          _218_v0 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_314_v61,new BigNumber(292))).length)), _283_cf7)).isLessThanOrEqualTo((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
          (globalState).f7 = (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))];
          let _315_v62;
          let _init8 = ((_316_cf7) => function (_317_i5) {
            return _module.D4.create_DC12(_316_cf7);
          })(_283_cf7);
          let _nw42 = Array((new BigNumber(29)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw42.length); _i0_8++) {
            _nw42[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _315_v62 = _nw42;
          let _318_v63;
          _318_v63 = _module.D4.create_DC12(_283_cf7);
          let _index33 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_315_v62).length));
          (_315_v62)[_index33] = _318_v63;
          let _index34 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_315_v62).length));
          (_315_v62)[_index34] = _318_v63;
          let _319_v64;
          _319_v64 = _dafny.Map.Empty.slice().updateUnsafe((((_225_v5).contains(_218_v0)) ? ((_225_v5).get(_218_v0)) : (_218_v0)),p1);
          let _320_v65;
          _320_v65 = _dafny.MultiSet.fromElements(_218_v0, _218_v0);
          let _index35 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_314_v61).length));
          (_314_v61)[_index35] = (new BigNumber((_320_v65).cardinality())).isEqualTo(new BigNumber(-438));
          let _321_v66;
          _321_v66 = _module.D7.create_DC21(new BigNumber(376), _283_cf7, _218_v0, !(_218_v0), _218_v0);
          let _322_v67;
          _322_v67 = _module.D6.create_DC17((_321_v66).dtor_cf33, _218_v0, _218_v0);
          let _323_v68;
          _323_v68 = _dafny.MultiSet.fromElements(_322_v67);
          let _index36 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_314_v61).length));
          let _rhs61 = (_319_v64).Merge(_319_v64);
          let _rhs62 = _218_v0;
          let _rhs63 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-919)), ((_324_v0) => function (_325_i6) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_324_v0,_324_v0)).length);
          })(_218_v0)), (new BigNumber(105)).minus(new BigNumber((_323_v68).cardinality())));
          let _rhs64 = new BigNumber(-271);
          let _rhs65 = !_dafny.areEqual(_223_v3, _223_v3);
          let _lhs41 = _314_v61;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_314_v61).length));
          let _lhs43 = globalState;
          _319_v64 = _rhs61;
          _lhs41[_lhs42] = _rhs62;
          _218_v0 = _rhs63;
          _lhs43.f15 = _rhs64;
          _218_v0 = _rhs65;
          let _326_v69;
          _326_v69 = _module.D6.create_DC16(_223_v3);
          let _327_v70;
          _327_v70 = _module.D6.create_DC18(_326_v69);
          let _328_v71;
          _328_v71 = _dafny.MultiSet.fromElements(_327_v70);
          let _329_v72;
          _329_v72 = _dafny.Seq.of(_module.D6.create_DC18(_326_v69));
          _328_v71 = ((_328_v71).Difference(_328_v71)).Intersect(_dafny.MultiSet.FromArray(_329_v72));
        }
        r1 = _283_cf7;
      } else if (_source6.is_DC7) {
        let _330___mcc_h12 = (_source6).cf8;
        let _331___mcc_h13 = (_source6).cf9;
        let _332___mcc_h14 = (_source6).cf10;
        let _333_cf10 = _332___mcc_h14;
        let _334_cf9 = _331___mcc_h13;
        let _335_cf8 = _330___mcc_h12;
        let _336_v75;
        let _init9 = ((_337_p1, _338_cf8) => function (_339_i7) {
          return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_337_p1, _337_p1), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_338_cf8, _338_cf8)).length)))).Intersect(function () {
            let _coll22 = new _dafny.Set();
            for (const _compr_22 of (function () {
              let _coll23 = new _dafny.Set();
              for (const _compr_23 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(130)))).Elements) {
                let _340_v73 = _compr_23;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(130))), _340_v73)) {
                  _coll23.add(_340_v73);
                }
              }
              return _coll23;
            }()).Elements) {
              let _341_v74 = _compr_22;
              if ((function () {
                let _coll24 = new _dafny.Set();
                for (const _compr_24 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(130)))).Elements) {
                  let _342_v73 = _compr_24;
                  if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(130))), _342_v73)) {
                    _coll24.add(_342_v73);
                  }
                }
                return _coll24;
              }()).contains(_341_v74)) {
                _coll22.add(_341_v74);
              }
            }
            return _coll22;
          }());
        })(p1, _335_cf8);
        let _nw43 = Array((new BigNumber(2)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw43.length); _i0_9++) {
          _nw43[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _336_v75 = _nw43;
        let _343_v76;
        _343_v76 = _dafny.MultiSet.fromElements(p1, (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]);
        let _344_v77;
        _344_v77 = _dafny.Set.fromElements(_343_v76, _343_v76, _module.__default.fm31((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))], (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))], globalState));
        let _index37 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_336_v75).length));
        (_336_v75)[_index37] = _344_v77;
        let _index38 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_336_v75).length));
        (_336_v75)[_index38] = (_344_v77).Intersect(_344_v77);
        (globalState).f13 = p1;
        let _345_v78;
        _345_v78 = _dafny.Seq.UnicodeFromString("vs");
        let _index39 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
        (_219_v1)[_index39] = new BigNumber((_dafny.Seq.Concat(_345_v78, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pq"), _345_v78))).length);
        if (_218_v0) {
          (globalState).f7 = (_dafny.ZERO).minus(p1);
          let _346_v79;
          let _nw44 = new _module.C0();
          _nw44.__ctor();
          _346_v79 = _nw44;
          (globalState).f15 = new BigNumber((_dafny.Seq.Concat(_345_v78, _dafny.Seq.Concat(_345_v78, _345_v78))).length);
          (globalState).f13 = (_dafny.ZERO).minus(((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]).plus((_this).fm25(_335_cf8, _335_cf8, globalState)));
          let _347_v80;
          _347_v80 = _dafny.Map.Empty.slice().updateUnsafe(_281_v35,_218_v0);
          let _348_v81;
          _348_v81 = _347_v80;
          _347_v80 = (_348_v81);
        } else {
          _219_v1 = _333_cf10;
          let _349_v82;
          let _nw45 = new _module.C0();
          _nw45.__ctor();
          _349_v82 = _nw45;
          let _350_v83;
          _350_v83 = new _dafny.CodePoint('w'.codePointAt(0));
          r1 = new BigNumber((_dafny.Seq.Concat(_345_v78, _dafny.Seq.update(_345_v78, _module.__default.safeIndex(new BigNumber(643), new BigNumber((_345_v78).length)), _350_v83))).length);
          _335_cf8 = _335_cf8;
          let _351_v84;
          let _nw46 = new _module.C0();
          _nw46.__ctor();
          _351_v84 = _nw46;
        }
      } else if (_source6.is_DC5) {
        let _352___mcc_h15 = (_source6).cf6;
        let _353_cf6 = _352___mcc_h15;
        let _354_v85;
        _354_v85 = _dafny.Seq.of((_218_v0) === (_218_v0), _module.__default.fm26(globalState), !(_218_v0) || (_218_v0), true, _218_v0);
        _354_v85 = _dafny.Seq.update(_dafny.Seq.Concat(_354_v85, _354_v85), _module.__default.safeIndex(new BigNumber(782), new BigNumber((_dafny.Seq.Concat(_354_v85, _354_v85)).length)), _218_v0);
        _218_v0 = ((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]).isLessThan(p1);
        _281_v35 = _281_v35;
        _218_v0 = !(_module.__default.safeDivisionInt(p1, p1)).isEqualTo((new BigNumber(114)).multipliedBy(p1));
      } else {
        let _355___mcc_h16 = (_source6).cf11;
        let _356_cf11 = _355___mcc_h16;
        _218_v0 = _module.__default.fm26(globalState);
        let _index40 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length));
        (_219_v1)[_index40] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]), (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))]));
        let _357_v86;
        _357_v86 = _dafny.Seq.UnicodeFromString("wxgofwv");
        let _358_v87;
        _358_v87 = new _dafny.CodePoint('e'.codePointAt(0));
        let _359_v88;
        _359_v88 = _dafny.Map.Empty.slice().updateUnsafe(_218_v0,_358_v87);
        let _360_v89;
        _360_v89 = _dafny.Seq.of(_218_v0);
        let _361_v90;
        _361_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_223_v3).length),_357_v86);
        let _362_v91;
        let _nw47 = Array((new BigNumber(12)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _357_v86;
        _nw47[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_357_v86, _357_v86);
        _nw47[(new BigNumber(2)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(3)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vsd"), _module.__default.fm30(new BigNumber((_222_v2).length), _218_v0, _dafny.Set.fromElements(new BigNumber((_359_v88).length), new BigNumber((_dafny.MultiSet.FromArray(_360_v89)).cardinality())), _218_v0, globalState));
        _nw47[(new BigNumber(5)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(6)).toNumber()] = (((_361_v90).contains(new BigNumber((_222_v2).length))) ? ((_361_v90).get(new BigNumber((_222_v2).length))) : (_357_v86));
        _nw47[(new BigNumber(7)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("otlcts");
        _nw47[(new BigNumber(9)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(10)).toNumber()] = _357_v86;
        _nw47[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("ke");
        _362_v91 = _nw47;
        let _index41 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_362_v91).length));
        (_362_v91)[_index41] = _357_v86;
        let _index42 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_362_v91).length));
        let _rhs66 = _357_v86;
        let _rhs67 = p1;
        let _lhs44 = _362_v91;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(954), new BigNumber((_362_v91).length));
        let _lhs46 = globalState;
        _lhs44[_lhs45] = _rhs66;
        _lhs46.f15 = _rhs67;
        let _363_v92;
        _363_v92 = _module.D7.create_DC20(p1);
        _363_v92 = _363_v92;
      }
      let _364_v93;
      let _nw48 = Array((new BigNumber(20)).toNumber()).fill(false);
      _364_v93 = _nw48;
      let _365_v94;
      _365_v94 = _dafny.Seq.of(_218_v0);
      let _index43 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length));
      (_364_v93)[_index43] = (_365_v94)[_module.__default.safeIndex(p1, new BigNumber((_365_v94).length))];
      let _366_v95;
      _366_v95 = _dafny.MultiSet.fromElements(_218_v0);
      let _367_v96;
      _367_v96 = _dafny.Set.fromElements(_366_v95);
      let _368_v98;
      _368_v98 = _dafny.Map.Empty.slice().updateUnsafe(p1,(((_366_v95).contains(_218_v0)) ? ((_366_v95).get(_218_v0)) : (new BigNumber(909))));
      let _index44 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length));
      let _rhs68 = _218_v0;
      let _rhs69 = (_367_v96).IsDisjointFrom((function () {
        let _coll25 = new _dafny.Set();
        for (const _compr_25 of (_dafny.Seq.of(_366_v95)).Elements) {
          let _369_v97 = _compr_25;
          if (_dafny.Seq.contains(_dafny.Seq.of(_366_v95), _369_v97)) {
            _coll25.add(_369_v97);
          }
        }
        return _coll25;
      }()).Difference(_367_v96));
      let _rhs70 = (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))];
      let _rhs71 = (p1).minus((new BigNumber((_368_v98).length)).plus(new BigNumber(188)));
      let _rhs72 = _218_v0;
      let _lhs47 = globalState;
      let _lhs48 = globalState;
      let _lhs49 = _364_v93;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length));
      _218_v0 = _rhs68;
      _218_v0 = _rhs69;
      _lhs47.f15 = _rhs70;
      _lhs48.f15 = _rhs71;
      _lhs49[_lhs50] = _rhs72;
      let _hi1 = (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))];
      for (let _370_i8 = (_this).fm25(_218_v0, _218_v0, globalState); _370_i8.isLessThan(_hi1); _370_i8 = _370_i8.plus(_dafny.ONE)) {
        _223_v3 = _dafny.Seq.Concat(_223_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(612)), function (_371_i9) {
          return new BigNumber(574);
        }));
        (globalState).f15 = (((_368_v98).contains(new BigNumber(((_222_v2).Intersect(_222_v2)).length))) ? ((_368_v98).get(new BigNumber(((_222_v2).Intersect(_222_v2)).length))) : (((_223_v3)[_module.__default.safeIndex(p1, new BigNumber((_223_v3).length))]).minus(_370_i8)));
        let _372_v99;
        _372_v99 = _dafny.Set.fromElements(_218_v0, _218_v0, _218_v0, !((_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))]), true);
        let _index45 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length));
        let _rhs73 = _dafny.Seq.Concat(_dafny.Seq.update(_365_v94, _module.__default.safeIndex(p1, new BigNumber((_365_v94).length)), (_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))]), _dafny.Seq.Concat(_365_v94, _365_v94));
        let _rhs74 = (_this).fm25(_218_v0, (_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))], globalState);
        let _rhs75 = (_module.__default.fm32(new BigNumber((_365_v94).length), (_219_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_219_v1).length))], _372_v99, globalState)).equals((_module.__default.fm32(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(188)), ((_373_v1) => function (_374_i10) {
          return (_373_v1)[_module.__default.safeIndex(new BigNumber(958), new BigNumber((_373_v1).length))];
        })(_219_v1))).length), new BigNumber(466), _372_v99, globalState)).Difference(_222_v2));
        let _lhs51 = globalState;
        let _lhs52 = _364_v93;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length));
        _365_v94 = _rhs73;
        _lhs51.f7 = _rhs74;
        _lhs52[_lhs53] = _rhs75;
        (globalState).f15 = _module.__default.safeDivisionInt((_this).fm25(!(!(true)), (_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))], globalState), _370_i8);
      }
      let _375_v100;
      _375_v100 = _module.D6.create_DC17(true, (_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))], !((_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))]));
      let _pat_let_tv1 = _364_v93;
      let _pat_let_tv2 = _364_v93;
      r0 = function (_source8) {
        if (_source8.is_DC17) {
          let _376___mcc_h18 = (_source8).cf24;
          let _377___mcc_h19 = (_source8).cf25;
          let _378___mcc_h20 = (_source8).cf26;
          let _379_cf26 = _378___mcc_h20;
          let _380_cf25 = _377___mcc_h19;
          let _381_cf24 = _376___mcc_h18;
          return new _dafny.CodePoint('c'.codePointAt(0));
        } else if (_source8.is_DC16) {
          let _382___mcc_h21 = (_source8).cf23;
          let _383_cf23 = _382___mcc_h21;
          if ((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_pat_let_tv1).length))]) {
            return new _dafny.CodePoint('b'.codePointAt(0));
          } else {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }
        } else {
          let _384___mcc_h22 = (_source8).cf27;
          let _385_cf27 = _384___mcc_h22;
          return new _dafny.CodePoint('x'.codePointAt(0));
        }
      }(_375_v100);
      r1 = p1;
      let _386_v101;
      _386_v101 = _dafny.Seq.of(_364_v93);
      r2 = ((!((_364_v93)[_module.__default.safeIndex(new BigNumber(687), new BigNumber((_364_v93).length))])) ? (_module.D7.create_DC19(_386_v101)) : (p0));
      let _387_v102;
      _387_v102 = _dafny.Seq.UnicodeFromString("lpsnk");
      let _388_v103;
      _388_v103 = _module.D1.create_DC3(_387_v102, _218_v0);
      r3 = (((_218_v0) && (false)) ? (_module.D1.create_DC3((_388_v103).dtor_cf3, _218_v0)) : (_388_v103));
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(p0, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements((_dafny.Set.fromElements(new BigNumber(117))).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality()),new BigNumber(679))).length)))), _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("yjblw"),(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("ywhli"), false)).dtor_cf4)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_module.D7.create_DC20(new BigNumber(898))).dtor_cf29, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-596)),new BigNumber(483))).length))).length), new BigNumber(46))).length))), function () {
        let _coll26 = new _dafny.Set();
        for (const _compr_26 of (_dafny.MultiSet.fromElements(new BigNumber(258), new BigNumber((_dafny.Set.fromElements(false, false, true, true)).length))).Elements) {
          let _389_v0 = _compr_26;
          if ((_dafny.MultiSet.fromElements(new BigNumber(258), new BigNumber((_dafny.Set.fromElements(false, false, true, true)).length))).contains(_389_v0)) {
            _coll26.add(_module.__default.safeModuloInt(_389_v0, new BigNumber((function () {
              let _coll27 = new _dafny.Set();
              for (const _compr_27 of _dafny.IntegerRange(new BigNumber(568), new BigNumber(301))) {
                let _390_v1 = _compr_27;
                if (((new BigNumber(568)).isLessThanOrEqualTo(_390_v1)) && ((_390_v1).isLessThan(new BigNumber(301)))) {
                  _coll27.add((_390_v1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),!(true))).length)));
                }
              }
              return _coll27;
            }()).length)));
          }
        }
        return _coll26;
      }(), (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, !(false))).length))).Difference(_dafny.Set.fromElements(new BigNumber(-746))));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      if (!(true) || (!(false))) {
        return !((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(new BigNumber(818), new BigNumber((_dafny.Seq.UnicodeFromString("fve")).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("aclnkmm")).length)))).contains(_module.D3.create_DC10(new BigNumber((function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of (_dafny.MultiSet.fromElements(new BigNumber(813))).Elements) {
    let _391_v0 = _compr_28;
    if ((_dafny.MultiSet.fromElements(new BigNumber(813))).contains(_391_v0)) {
      _coll28.push([(_391_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(595), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('a'.codePointAt(0)))).length),true)).length))).length))).length))).length)),true]);
    }
  }
  return _coll28;
}()).length), new BigNumber((_dafny.Seq.of(true, true)).length))));
      } else {
        return (new BigNumber(653)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(488)), function (_392_i0) {
          return new BigNumber(782);
        })).length));
      }
    };
    fm1(p0, p1, globalState) {
      let _this = this;
      return true;
    };
    fm2(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(690)), function (_393_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber(-61));
    };
    fm24(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(true);
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _hi2 = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), new BigNumber(693));
      for (let _394_i0 = p0; _394_i0.isLessThan(_hi2); _394_i0 = _394_i0.plus(_dafny.ONE)) {
        let _395_v0;
        _395_v0 = false;
        let _396_v1;
        let _nw49 = Array((new BigNumber(25)).toNumber()).fill(false);
        _396_v1 = _nw49;
        let _397_v2;
        _397_v2 = _dafny.MultiSet.fromElements(_396_v1, _396_v1);
        _395_v0 = !((_397_v2).Intersect(_397_v2)).contains(_396_v1);
        (globalState).f4 = _394_i0;
        let _398_v3;
        let _nw50 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _398_v3 = _nw50;
        let _index46 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_398_v3).length));
        (_398_v3)[_index46] = p0;
        let _399_v4;
        _399_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,_395_v0);
        let _400_v5;
        _400_v5 = _dafny.Map.Empty.slice().updateUnsafe(_395_v0,_399_v4);
        let _index47 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_398_v3).length));
        let _rhs76 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p0, _394_i0), new BigNumber((_dafny.Set.fromElements(_395_v0)).length));
        let _rhs77 = (_this).fm1((_400_v5).Merge((_400_v5).update(p2, _399_v4)), _394_i0, globalState);
        let _lhs54 = _398_v3;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_398_v3).length));
        _lhs54[_lhs55] = _rhs76;
        _395_v0 = _rhs77;
        _395_v0 = !(_395_v0);
      }
      let _401_v6;
      _401_v6 = false;
      let _402_v7;
      _402_v7 = _dafny.Seq.UnicodeFromString("qtrupg");
      _401_v6 = (_module.__default.safeDivisionInt(new BigNumber((_402_v7).length), new BigNumber(-925))).isLessThanOrEqualTo(_module.__default.safeModuloInt(p0, p0));
      let _403_i1;
      _403_i1 = _dafny.ZERO;
      L1: {
        while (_401_v6) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_403_i1)) {
              break L1;
            }
            _403_i1 = (_403_i1).plus(_dafny.ONE);
            _401_v6 = _401_v6;
            let _404_v8;
            _404_v8 = _dafny.Set.fromElements(_401_v6);
            let _405_v9;
            _405_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.fromElements(_404_v8, _404_v8));
            let _406_v10;
            _406_v10 = _dafny.MultiSet.fromElements(_404_v8);
            _405_v9 = (_405_v9).update(p0, _406_v10);
            let _407_v11;
            let _nw51 = new _module.C0();
            _nw51.__ctor();
            _407_v11 = _nw51;
            let _408_v12;
            let _nw52 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Set.Empty);
            _408_v12 = _nw52;
            let _index48 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_408_v12).length));
            (_408_v12)[_index48] = _404_v8;
            let _index49 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_408_v12).length));
            (_408_v12)[_index49] = (_404_v8).Intersect((_404_v8).Intersect(_404_v8));
          }
        }
      }
      let _409_v13;
      let _nw53 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
      _409_v13 = _nw53;
      let _410_v14;
      _410_v14 = _dafny.Set.fromElements(true);
      let _411_v15;
      _411_v15 = _dafny.Set.fromElements(_410_v14, _410_v14, _410_v14);
      let _index50 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_409_v13).length));
      (_409_v13)[_index50] = ((_401_v6) ? (_411_v15) : (_411_v15));
      let _index51 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_409_v13).length));
      (_409_v13)[_index51] = _411_v15;
      let _412_v16;
      let _nw54 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
      _412_v16 = _nw54;
      let _413_v17;
      _413_v17 = _dafny.Set.fromElements(p0, (_this).fm2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,!(p2))).length), globalState));
      let _414_v18;
      _414_v18 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
      let _415_v19;
      _415_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
      let _416_v20;
      _416_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(756),_413_v17);
      let _417_v21;
      _417_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_402_v7, _module.__default.fm30(p0, _401_v6, _413_v17, _401_v6, globalState)),(_414_v18).Merge(_module.__default.fm33(!(p2), (((_415_v19).contains(p0)) ? ((_415_v19).get(p0)) : (_401_v6)), new BigNumber(((((_416_v20).contains(p0)) ? ((_416_v20).get(p0)) : (_413_v17))).length), globalState)));
      let _418_v22;
      let _nw55 = Array((new BigNumber(2)).toNumber()).fill(_dafny.MultiSet.Empty);
      _418_v22 = _nw55;
      let _419_v23;
      let _nw56 = new _module.C0();
      _nw56.__ctor();
      _419_v23 = _nw56;
      let _420_v24;
      _420_v24 = _dafny.Map.Empty.slice().updateUnsafe(_419_v23,p0);
      let _421_v25;
      _421_v25 = _dafny.MultiSet.fromElements(_420_v24);
      let _index52 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_418_v22).length));
      (_418_v22)[_index52] = (_421_v25).Union(_421_v25);
      let _422_v27;
      _422_v27 = _module.D11.create_DC28(_413_v17);
      let _423_v28;
      _423_v28 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(954)), function (_424_i2) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }), _module.__default.fm30(p0, p2, (_422_v27).dtor_cf45, p2, globalState), _402_v7);
      let _index53 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_418_v22).length));
      let _rhs78 = _412_v16;
      let _rhs79 = (function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_423_v28).Elements) {
          let _425_v26 = _compr_29;
          if ((_423_v28).contains(_425_v26)) {
            _coll29.push([_425_v26,_414_v18]);
          }
        }
        return _coll29;
      }()).update(_402_v7, (_414_v18).Merge(_414_v18));
      let _rhs80 = _421_v25;
      let _rhs81 = p0;
      let _rhs82 = _401_v6;
      let _lhs56 = _418_v22;
      let _lhs57 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_418_v22).length));
      let _lhs58 = globalState;
      _412_v16 = _rhs78;
      _417_v21 = _rhs79;
      _lhs56[_lhs57] = _rhs80;
      _lhs58.f4 = _rhs81;
      _401_v6 = _rhs82;
      let _426_v29;
      let _init10 = ((_427_v6) => function (_428_i3) {
        return ((_427_v6) ? (new _dafny.CodePoint('g'.codePointAt(0))) : (new _dafny.CodePoint('e'.codePointAt(0))));
      })(_401_v6);
      let _nw57 = Array((new BigNumber(25)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw57.length); _i0_10++) {
        _nw57[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _426_v29 = _nw57;
      let _429_v30;
      _429_v30 = new _dafny.CodePoint('l'.codePointAt(0));
      let _index54 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_426_v29).length));
      (_426_v29)[_index54] = _429_v30;
      let _430_v31;
      let _nw58 = Array((_dafny.ONE).toNumber());
      _nw58[(_dafny.ZERO).toNumber()] = p2;
      _430_v31 = _nw58;
      let _index55 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_430_v31).length));
      (_430_v31)[_index55] = _401_v6;
      let _index56 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_426_v29).length));
      let _index57 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_430_v31).length));
      let _rhs83 = _429_v30;
      let _rhs84 = _module.D11.create_DC28(_413_v17);
      let _rhs85 = p2;
      let _rhs86 = p0;
      let _rhs87 = p0;
      let _lhs59 = _426_v29;
      let _lhs60 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_426_v29).length));
      let _lhs61 = _430_v31;
      let _lhs62 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_430_v31).length));
      let _lhs63 = globalState;
      let _lhs64 = globalState;
      _lhs59[_lhs60] = _rhs83;
      _422_v27 = _rhs84;
      _lhs61[_lhs62] = _rhs85;
      _lhs63.f7 = _rhs86;
      _lhs64.f13 = _rhs87;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let _431_v0;
      _431_v0 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-77)), function (_432_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
      let _433_v1;
      _433_v1 = _module.D4.create_DC11(_dafny.Seq.Concat(_431_v0, _dafny.Seq.of(p0)));
      let _source9 = _433_v1;
      if (_source9.is_DC12) {
        let _434___mcc_h0 = (_source9).cf16;
        let _435_cf16 = _434___mcc_h0;
        let _436_v2;
        let _init11 = function (_437_i1) {
          return _module.__default.safeDivisionInt(_437_i1, new BigNumber(344));
        };
        let _nw59 = Array((new BigNumber(28)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw59.length); _i0_11++) {
          _nw59[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _436_v2 = _nw59;
        let _index58 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length));
        (_436_v2)[_index58] = _435_cf16;
        let _index59 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length));
        (_436_v2)[_index59] = _435_cf16;
        let _438_v3;
        _438_v3 = false;
        if (_438_v3) {
          (globalState).f4 = (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))];
          let _439_v4;
          _439_v4 = _dafny.Map.Empty.slice().updateUnsafe((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))],(_435_cf16).minus(_435_cf16));
          _439_v4 = _439_v4;
          let _440_v5;
          let _init12 = function (_441_i2) {
            return _module.D3.create_DC9(new _dafny.CodePoint('v'.codePointAt(0)));
          };
          let _nw60 = Array((new BigNumber(28)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw60.length); _i0_12++) {
            _nw60[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _440_v5 = _nw60;
          let _442_v6;
          _442_v6 = new _dafny.CodePoint('m'.codePointAt(0));
          let _443_v7;
          _443_v7 = _module.D3.create_DC9(_442_v6);
          let _index60 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_440_v5).length));
          (_440_v5)[_index60] = ((_438_v3) ? (_443_v7) : (_443_v7));
          let _index61 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_440_v5).length));
          (_440_v5)[_index61] = _443_v7;
          let _rhs88 = _436_v2;
          let _rhs89 = _module.__default.safeDivisionInt((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
          let _lhs65 = globalState;
          _436_v2 = _rhs88;
          _lhs65.f7 = _rhs89;
          _438_v3 = !(!(!(_438_v3)));
        } else {
          (globalState).f13 = (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))];
          let _index62 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length));
          (_436_v2)[_index62] = (_dafny.ZERO).minus((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
          (globalState).f15 = (_dafny.ZERO).minus((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
          let _444_v8;
          _444_v8 = _module.D3.create_DC10((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], (_dafny.ZERO).minus(_435_cf16));
          let _445_v9;
          _445_v9 = _dafny.Map.Empty.slice().updateUnsafe(_444_v8,_438_v3);
          let _446_v10;
          _446_v10 = _dafny.MultiSet.fromElements(new BigNumber((_445_v9).length));
          let _447_v11;
          _447_v11 = _dafny.Seq.of(new BigNumber(720));
          let _448_v12;
          _448_v12 = _module.D1.create_DC4(_446_v10);
          let _449_v13;
          _449_v13 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], _435_cf16));
          let _450_v14;
          _450_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_446_v10).cardinality()),_438_v3);
          let _451_v15;
          let _nw61 = Array((new BigNumber(23)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))])).Difference(_446_v10);
          _nw61[(_dafny.ONE).toNumber()] = _446_v10;
          _nw61[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_447_v11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), ((_452_cf16) => function (_453_i3) {
            return new BigNumber((_dafny.Set.fromElements(_452_cf16)).length);
          })(_435_cf16))));
          _nw61[(new BigNumber(3)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(4)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(5)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(6)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(7)).toNumber()] = (_448_v12).dtor_cf5;
          _nw61[(new BigNumber(8)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(9)).toNumber()] = _module.__default.fm31(_435_cf16, (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], globalState);
          _nw61[(new BigNumber(10)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(11)).toNumber()] = (_449_v13)[_module.__default.safeIndex((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], new BigNumber((_449_v13).length))];
          _nw61[(new BigNumber(12)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(13)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_450_v14).length), (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
          _nw61[(new BigNumber(15)).toNumber()] = (_446_v10).Difference(_446_v10);
          _nw61[(new BigNumber(16)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
          _nw61[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_447_v11);
          _nw61[(new BigNumber(19)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(20)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(21)).toNumber()] = _446_v10;
          _nw61[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], (_dafny.ZERO).minus(new BigNumber((p0).length)), new BigNumber((_447_v11).length), new BigNumber(526));
          _451_v15 = _nw61;
          let _index63 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_451_v15).length));
          (_451_v15)[_index63] = (_446_v10).Difference(_446_v10);
          let _index64 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_451_v15).length));
          let _rhs90 = _444_v8;
          let _rhs91 = _438_v3;
          let _rhs92 = _446_v10;
          let _lhs66 = _451_v15;
          let _lhs67 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_451_v15).length));
          _444_v8 = _rhs90;
          _438_v3 = _rhs91;
          _lhs66[_lhs67] = _rhs92;
          (_this).m12(_438_v3, _438_v3, _436_v2, (_module.D7.create_DC22(_438_v3, (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))], _438_v3, (_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))])).dtor_cf35, globalState);
        }
        let _454_v16;
        _454_v16 = _module.D8.create_DC25((_436_v2)[_module.__default.safeIndex(new BigNumber(354), new BigNumber((_436_v2).length))]);
        _454_v16 = _454_v16;
        let _455_v17;
        _455_v17 = _dafny.Seq.of((new BigNumber((p0).length)).isLessThanOrEqualTo(new BigNumber((p0).length)));
        _438_v3 = (_455_v17)[_module.__default.safeIndex(_435_cf16, new BigNumber((_455_v17).length))];
      } else {
        let _456___mcc_h1 = (_source9).cf15;
        let _457_cf15 = _456___mcc_h1;
        let _458_v18;
        _458_v18 = new BigNumber(757);
        (globalState).f7 = (_this).fm2(_458_v18, globalState);
        let _459_v19;
        _459_v19 = false;
        _459_v19 = (new BigNumber(936)).isEqualTo((_dafny.ZERO).minus((_this).fm2(_458_v18, globalState)));
        let _460_v20;
        _460_v20 = _dafny.Set.fromElements(_459_v19);
        let _461_v21;
        _461_v21 = _dafny.Seq.of(_459_v19);
        let _462_v22;
        _462_v22 = _module.D1.create_DC3(p0, !(false));
        let _463_v23;
        let _nw62 = Array((new BigNumber(10)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _459_v19;
        _nw62[(_dafny.ONE).toNumber()] = _459_v19;
        _nw62[(new BigNumber(2)).toNumber()] = _459_v19;
        _nw62[(new BigNumber(3)).toNumber()] = (_459_v19) === (_459_v19);
        _nw62[(new BigNumber(4)).toNumber()] = ((_459_v19) ? (_459_v19) : (_459_v19));
        _nw62[(new BigNumber(5)).toNumber()] = _459_v19;
        _nw62[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(_459_v19)).IsSubsetOf(_460_v20);
        _nw62[(new BigNumber(7)).toNumber()] = _459_v19;
        _nw62[(new BigNumber(8)).toNumber()] = !(((_461_v21)[_module.__default.safeIndex(_458_v18, new BigNumber((_461_v21).length))]) === (_459_v19));
        _nw62[(new BigNumber(9)).toNumber()] = (_462_v22).dtor_cf4;
        _463_v23 = _nw62;
        _463_v23 = _463_v23;
        if (_459_v19) {
          let _464_v24;
          let _init13 = ((_465_p0) => function (_466_i4) {
            return _dafny.Seq.Concat(_465_p0, _465_p0);
          })(p0);
          let _nw63 = Array((new BigNumber(10)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw63.length); _i0_13++) {
            _nw63[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _464_v24 = _nw63;
          _464_v24 = _464_v24;
          _464_v24 = _464_v24;
          let _467_v25;
          let _nw64 = new _module.C0();
          _nw64.__ctor();
          _467_v25 = _nw64;
          let _468_v26;
          _468_v26 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_458_v18, _458_v18),_458_v18);
          _468_v26 = (_468_v26).update((_this).fm2(_458_v18, globalState), _458_v18);
          _459_v19 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_461_v21, _461_v21), _dafny.Seq.Concat(_module.__default.fm27(globalState), _461_v21));
        } else {
          (globalState).f13 = (new BigNumber(972)).multipliedBy(_458_v18);
          let _469_v27;
          _469_v27 = _dafny.Seq.of(_458_v18, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_459_v19,new BigNumber(377))).length), new BigNumber(735));
          let _470_v28;
          _470_v28 = _dafny.Map.Empty.slice().updateUnsafe(_469_v27,_463_v23);
          let _rhs93 = (_458_v18).plus(_458_v18);
          let _rhs94 = _463_v23;
          let _rhs95 = (_470_v28).contains(_469_v27);
          let _rhs96 = new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tmislcqxm"), _dafny.Seq.UnicodeFromString("jtsjc")))).length);
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          _lhs68.f15 = _rhs93;
          _463_v23 = _rhs94;
          _459_v19 = _rhs95;
          _lhs69.f4 = _rhs96;
          (globalState).f13 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(815)), function (_471_i5) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          })).length);
          let _472_v29;
          let _nw65 = new _module.C0();
          _nw65.__ctor();
          _472_v29 = _nw65;
          let _473_v30;
          let _nw66 = new _module.C1();
          _nw66.__ctor();
          _473_v30 = _nw66;
        }
      }
      let _474_v31;
      let _init14 = function (_475_i6) {
        return (_module.D8.create_DC24(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(959),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('d'.codePointAt(0)))))).cardinality()))).length),new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality()))).length), true)).dtor_cf41;
      };
      let _nw67 = Array((new BigNumber(13)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw67.length); _i0_14++) {
        _nw67[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _474_v31 = _nw67;
      let _476_v32;
      _476_v32 = false;
      let _477_v33;
      _477_v33 = _module.D6.create_DC17(_476_v32, false, true);
      let _478_v34;
      _478_v34 = _dafny.Seq.of(_476_v32, _476_v32, (_477_v33).dtor_cf26);
      let _479_v35;
      _479_v35 = new BigNumber(348);
      let _480_v36;
      _480_v36 = _dafny.Map.Empty.slice().updateUnsafe(true,_479_v35);
      let _481_v37;
      _481_v37 = _dafny.Map.Empty.slice().updateUnsafe((_478_v34)[_module.__default.safeIndex(new BigNumber((_480_v36).length), new BigNumber((_478_v34).length))],_474_v31);
      let _482_v38;
      _482_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_474_v31, (((_481_v37).contains(_476_v32)) ? ((_481_v37).get(_476_v32)) : (_474_v31))),_479_v35);
      let _483_v39;
      _483_v39 = _dafny.MultiSet.fromElements((_this).fm8(new BigNumber((_dafny.MultiSet.FromArray(_module.__default.fm27(globalState))).cardinality()), _476_v32, _479_v35, globalState), !(false));
      (globalState).f15 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_482_v38).contains(_dafny.Seq.of(_474_v31))) ? ((_482_v38).get(_dafny.Seq.of(_474_v31))) : (_module.__default.safeDivisionInt(new BigNumber((_483_v39).cardinality()), _479_v35)))));
      (globalState).f7 = (new BigNumber((_dafny.MultiSet.fromElements(!(_476_v32))).cardinality())).plus(new BigNumber((_dafny.MultiSet.FromArray(p0)).cardinality()));
      let _484_v40;
      let _nw68 = Array((_dafny.ONE).toNumber()).fill([]);
      _484_v40 = _nw68;
      let _485_v41;
      _485_v41 = _dafny.Map.Empty.slice().updateUnsafe(_479_v35,_474_v31);
      let _486_v42;
      let _nw69 = Array((new BigNumber(29)).toNumber());
      _nw69[(_dafny.ZERO).toNumber()] = _474_v31;
      _nw69[(_dafny.ONE).toNumber()] = _474_v31;
      _nw69[(new BigNumber(2)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(3)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(4)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(5)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(6)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(7)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(8)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(9)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(10)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(11)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(12)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(13)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(14)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(15)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(16)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(17)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(18)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(19)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(20)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(21)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(22)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(23)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(24)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(25)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(26)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(27)).toNumber()] = _474_v31;
      _nw69[(new BigNumber(28)).toNumber()] = (((_485_v41).contains(new BigNumber(-38))) ? ((_485_v41).get(new BigNumber(-38))) : (_474_v31));
      _486_v42 = _nw69;
      let _487_v43;
      _487_v43 = _module.D0.create_DC1(_474_v31);
      let _index65 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_486_v42).length));
      (_486_v42)[_index65] = (_487_v43).dtor_cf1;
      let _488_v44;
      let _init15 = ((_489_v39) => function (_490_i7) {
        return _489_v39;
      })(_483_v39);
      let _nw70 = Array((new BigNumber(29)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw70.length); _i0_15++) {
        _nw70[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _488_v44 = _nw70;
      let _index66 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_488_v44).length));
      (_488_v44)[_index66] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_476_v32), _478_v34));
      let _491_v45;
      _491_v45 = _dafny.Set.fromElements(_476_v32, _476_v32, false, _476_v32, _476_v32);
      let _492_v46;
      _492_v46 = _dafny.Seq.of(_484_v40, _484_v40, _484_v40, _484_v40, _484_v40);
      let _index67 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_486_v42).length));
      let _index68 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_488_v44).length));
      let _rhs97 = (_492_v46)[_module.__default.safeIndex(_479_v35, new BigNumber((_492_v46).length))];
      let _rhs98 = _474_v31;
      let _rhs99 = _dafny.MultiSet.fromElements(_476_v32, _476_v32);
      let _rhs100 = _491_v45;
      let _lhs70 = _486_v42;
      let _lhs71 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_486_v42).length));
      let _lhs72 = _488_v44;
      let _lhs73 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_488_v44).length));
      _484_v40 = _rhs97;
      _lhs70[_lhs71] = _rhs98;
      _lhs72[_lhs73] = _rhs99;
      _491_v45 = _rhs100;
      let _pat_let_tv3 = p0;
      let _pat_let_tv4 = p0;
      let _pat_let_tv5 = _478_v34;
      let _pat_let_tv6 = _479_v35;
      (globalState).f15 = function (_source10) {
        if (_source10.is_DC12) {
          let _493___mcc_h2 = (_source10).cf16;
          let _494_cf16 = _493___mcc_h2;
          return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))), _dafny.MultiSet.FromArray(_pat_let_tv3), _dafny.MultiSet.FromArray(_pat_let_tv4))).length), new BigNumber((_pat_let_tv5).length), _494_cf16, new BigNumber(561), (_dafny.ZERO).minus(new BigNumber(-846)))).cardinality());
        } else {
          let _495___mcc_h3 = (_source10).cf15;
          let _496_cf15 = _495___mcc_h3;
          return _pat_let_tv6;
        }
      }(_433_v1);
      let _497_v47;
      let _nw71 = Array((new BigNumber(20)).toNumber());
      _nw71[(_dafny.ZERO).toNumber()] = new BigNumber((_478_v34).length);
      _nw71[(_dafny.ONE).toNumber()] = _479_v35;
      _nw71[(new BigNumber(2)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(3)).toNumber()] = (_this).fm2(_479_v35, globalState);
      _nw71[(new BigNumber(4)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(5)).toNumber()] = new BigNumber((p0).length);
      _nw71[(new BigNumber(6)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(7)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(8)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(9)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(10)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(11)).toNumber()] = new BigNumber((_478_v34).length);
      _nw71[(new BigNumber(12)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(13)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(14)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(15)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(16)).toNumber()] = _dafny.ZERO;
      _nw71[(new BigNumber(17)).toNumber()] = _479_v35;
      _nw71[(new BigNumber(18)).toNumber()] = new BigNumber(-971);
      _nw71[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus((((_483_v39).contains(_476_v32)) ? ((_483_v39).get(_476_v32)) : (new BigNumber(-672))));
      _497_v47 = _nw71;
      let _498_v48;
      _498_v48 = _dafny.Map.Empty.slice().updateUnsafe(_497_v47,true);
      _498_v48 = (_498_v48).update(_497_v47, _476_v32);
      return;
    }
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _499_v0;
      _499_v0 = new BigNumber(978);
      let _500_v1;
      _500_v1 = _dafny.Set.fromElements(_499_v0, _499_v0);
      if ((_500_v1).equals(_dafny.Set.fromElements(_499_v0))) {
        let _501_v2;
        let _init16 = ((_502_v0) => function (_503_i0) {
          return (_503_i0).multipliedBy(_502_v0);
        })(_499_v0);
        let _nw72 = Array((new BigNumber(2)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw72.length); _i0_16++) {
          _nw72[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _501_v2 = _nw72;
        let _index69 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length));
        (_501_v2)[_index69] = new BigNumber((_dafny.Seq.of(_501_v2, _501_v2, _501_v2)).length);
        let _index70 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length));
        (_501_v2)[_index70] = (_499_v0).multipliedBy(_499_v0);
        (globalState).f7 = (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))];
        let _504_v3;
        _504_v3 = _module.D2.create_DC7(p3, _module.__default.fm34(_499_v0, _499_v0, p3, globalState), _501_v2);
        let _source11 = _504_v3;
        if (_source11.is_DC6) {
          let _505___mcc_h0 = (_source11).cf7;
          let _506_cf7 = _505___mcc_h0;
          let _507_v4;
          _507_v4 = false;
          _507_v4 = !((_this).fm8(_506_cf7, p3, (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], globalState));
          let _508_v5;
          let _nw73 = new _module.C1();
          _nw73.__ctor();
          _508_v5 = _nw73;
          let _509_v6;
          _509_v6 = _dafny.Map.Empty.slice().updateUnsafe(_500_v1,_499_v0);
          _509_v6 = _509_v6;
          _507_v4 = ((_507_v4) ? (_507_v4) : (_507_v4));
        } else if (_source11.is_DC7) {
          let _510___mcc_h1 = (_source11).cf8;
          let _511___mcc_h2 = (_source11).cf9;
          let _512___mcc_h3 = (_source11).cf10;
          let _513_cf10 = _512___mcc_h3;
          let _514_cf9 = _511___mcc_h2;
          let _515_cf8 = _510___mcc_h1;
          let _516_v7;
          _516_v7 = new _dafny.CodePoint('q'.codePointAt(0));
          let _517_v8;
          _517_v8 = _dafny.MultiSet.fromElements(_499_v0);
          let _518_v9;
          _518_v9 = _dafny.Map.Empty.slice().updateUnsafe(_499_v0,(_dafny.Set.fromElements((((_517_v8).contains(_499_v0)) ? ((_517_v8).get(_499_v0)) : (new BigNumber(286))), new BigNumber(((_module.__default.fm35(globalState)).update((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], _515_cf8)).length))).Union(_500_v1));
          let _519_v10;
          let _nw74 = Array((new BigNumber(22)).toNumber()).fill(false);
          _519_v10 = _nw74;
          let _520_v11;
          _520_v11 = _dafny.Map.Empty.slice().updateUnsafe(_499_v0,p0);
          let _521_v12;
          _521_v12 = _dafny.Seq.UnicodeFromString("jf");
          let _522_v13;
          _522_v13 = _dafny.Map.Empty.slice().updateUnsafe((_517_v8).update(new BigNumber((_521_v12).length), _module.__default.abs((_dafny.ZERO).minus((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]))),_516_v7);
          let _523_v14;
          _523_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_521_v12).length),p3);
          let _rhs101 = ((p3) ? ((((_522_v13).contains(_517_v8)) ? ((_522_v13).get(_517_v8)) : (new _dafny.CodePoint('j'.codePointAt(0))))) : (_516_v7));
          let _rhs102 = _518_v9;
          let _rhs103 = p0;
          let _rhs104 = ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_523_v14).length),p0)).Merge(_520_v11)).Merge(_520_v11);
          let _rhs105 = ((_515_cf8) ? (p3) : (!(p3)));
          _516_v7 = _rhs101;
          _518_v9 = _rhs102;
          _519_v10 = _rhs103;
          _520_v11 = _rhs104;
          _515_cf8 = _rhs105;
          _523_v14 = (_523_v14).update(_499_v0, (p1)[_module.__default.safeIndex(new BigNumber(519), new BigNumber((p1).length))]);
          _515_cf8 = true;
          let _524_v15;
          _524_v15 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jfqn"));
          let _rhs106 = (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))];
          let _rhs107 = _524_v15;
          let _lhs74 = globalState;
          _lhs74.f4 = _rhs106;
          _524_v15 = _rhs107;
        } else if (_source11.is_DC5) {
          let _525___mcc_h4 = (_source11).cf6;
          let _526_cf6 = _525___mcc_h4;
          let _527_v16;
          let _nw75 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _527_v16 = _nw75;
          let _528_v17;
          _528_v17 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
          let _index71 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_527_v16).length));
          (_527_v16)[_index71] = (_528_v17).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,p2));
          let _index72 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_527_v16).length));
          (_527_v16)[_index72] = (_528_v17).Merge((_528_v17).Merge(_module.__default.fm36((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], globalState)));
          _526_cf6 = ((p3) ? ((_504_v3).dtor_cf10) : (((p3) ? (_526_cf6) : (_526_cf6))));
          let _529_v18;
          _529_v18 = _dafny.Set.fromElements(_501_v2, _526_cf6, _526_cf6, _526_cf6);
          let _index73 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((p0).length));
          (p0)[_index73] = (_529_v18).IsProperSubsetOf(_529_v18);
          let _index74 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((p0).length));
          let _rhs108 = ((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]).minus(_module.__default.safeModuloInt((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], new BigNumber(-246)));
          let _rhs109 = p3;
          let _rhs110 = (_dafny.ZERO).minus(_499_v0);
          let _lhs75 = p0;
          let _lhs76 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((p0).length));
          _499_v0 = _rhs108;
          _lhs75[_lhs76] = _rhs109;
          _499_v0 = _rhs110;
          let _530_v19;
          _530_v19 = _module.D2.create_DC5(_501_v2);
          let _531_v20;
          _531_v20 = _dafny.Map.Empty.slice().updateUnsafe(_530_v19,p3);
          let _532_v21;
          let _nw76 = Array((new BigNumber(2)).toNumber());
          _nw76[(_dafny.ZERO).toNumber()] = _531_v20;
          _nw76[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_530_v19,p3);
          _532_v21 = _nw76;
          let _533_v22;
          _533_v22 = _dafny.Map.Empty.slice().updateUnsafe(_530_v19,true);
          let _index75 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_532_v21).length));
          (_532_v21)[_index75] = _533_v22;
          let _534_v23;
          _534_v23 = _dafny.Seq.UnicodeFromString("racjbb");
          let _535_v24;
          _535_v24 = _module.D1.create_DC3(_534_v23, p3);
          let _index76 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_532_v21).length));
          let _rhs111 = _533_v22;
          let _rhs112 = (((p0)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((p0).length))]) ? (_535_v24) : (_535_v24));
          let _rhs113 = (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))];
          let _lhs77 = _532_v21;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_532_v21).length));
          let _lhs79 = globalState;
          _lhs77[_lhs78] = _rhs111;
          _535_v24 = _rhs112;
          _lhs79.f13 = _rhs113;
        } else {
          let _536___mcc_h5 = (_source11).cf11;
          let _537_cf11 = _536___mcc_h5;
          let _538_v25;
          _538_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_dafny.ZERO).minus(_499_v0));
          let _539_v26;
          _539_v26 = _dafny.Seq.of(new BigNumber(331), (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]);
          let _540_v27;
          _540_v27 = _dafny.MultiSet.fromElements((_539_v26)[_module.__default.safeIndex((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], new BigNumber((_539_v26).length))], new BigNumber(86));
          let _541_v28;
          _541_v28 = _dafny.Seq.of(_dafny.MultiSet.fromElements((((_538_v25).contains(p3)) ? ((_538_v25).get(p3)) : ((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]))), _540_v27, _540_v27);
          _541_v28 = _541_v28;
          let _542_v29;
          _542_v29 = _dafny.Map.Empty.slice().updateUnsafe(_499_v0,p3);
          let _index77 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length));
          (_501_v2)[_index77] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(new BigNumber((_542_v29).length))).length), _499_v0)).multipliedBy(((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]).multipliedBy(_499_v0))));
          (globalState).f13 = _module.__default.safeDivisionInt(_499_v0, (_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))]);
          let _index78 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((p0).length));
          (p0)[_index78] = _module.__default.fm26(globalState);
          let _index79 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((p0).length));
          (p0)[_index79] = false;
        }
        let _index80 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length));
        (_501_v2)[_index80] = _module.__default.safeDivisionInt(_499_v0, _499_v0);
        let _543_v30;
        _543_v30 = _dafny.Seq.of(_501_v2);
        let _544_v31;
        _544_v31 = _dafny.Seq.of((_543_v30)[_module.__default.safeIndex((_501_v2)[_module.__default.safeIndex(new BigNumber(436), new BigNumber((_501_v2).length))], new BigNumber((_543_v30).length))], _501_v2, _501_v2);
        _544_v31 = _544_v31;
      } else {
        let _545_v32;
        let _init17 = ((_546_p1, _547_v0) => function (_548_i1) {
          return (_546_p1)[_module.__default.safeIndex(_547_v0, new BigNumber((_546_p1).length))];
        })(p1, _499_v0);
        let _nw77 = Array((new BigNumber(24)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw77.length); _i0_17++) {
          _nw77[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _545_v32 = _nw77;
        let _549_v33;
        let _nw78 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _549_v33 = _nw78;
        let _rhs114 = new BigNumber(-932);
        let _rhs115 = _499_v0;
        let _rhs116 = _545_v32;
        let _rhs117 = _549_v33;
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        _lhs80.f15 = _rhs114;
        _lhs81.f7 = _rhs115;
        _545_v32 = _rhs116;
        _549_v33 = _rhs117;
        let _550_v34;
        _550_v34 = false;
        _550_v34 = p3;
        let _551_v35;
        let _nw79 = new _module.C1();
        _nw79.__ctor();
        _551_v35 = _nw79;
        let _552_v36;
        _552_v36 = _dafny.Seq.of(_499_v0, _499_v0);
        _550_v34 = (((p3) ? (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_552_v36, _module.__default.safeIndex(_499_v0, new BigNumber((_552_v36).length)), _499_v0))).cardinality())) : (_499_v0))).isLessThanOrEqualTo(_499_v0);
        _550_v34 = p3;
      }
      let _553_v37;
      let _nw80 = Array((new BigNumber(18)).toNumber()).fill([]);
      _553_v37 = _nw80;
      let _554_v38;
      _554_v38 = false;
      let _rhs118 = (new BigNumber(784)).plus(new BigNumber(46));
      let _rhs119 = _553_v37;
      let _rhs120 = !(p3);
      let _rhs121 = _554_v38;
      let _rhs122 = _554_v38;
      let _lhs82 = globalState;
      _lhs82.f13 = _rhs118;
      _553_v37 = _rhs119;
      _554_v38 = _rhs120;
      _554_v38 = _rhs121;
      _554_v38 = _rhs122;
      let _hi3 = new BigNumber(-680);
      for (let _555_i2 = _499_v0; _555_i2.isLessThan(_hi3); _555_i2 = _555_i2.plus(_dafny.ONE)) {
        let _556_v39;
        _556_v39 = _module.D0.create_DC1(p0);
        let _557_v40;
        _557_v40 = _dafny.MultiSet.fromElements(_556_v39, _556_v39, _556_v39, _556_v39);
        let _558_v41;
        _558_v41 = _dafny.Map.Empty.slice().updateUnsafe(_555_i2,_557_v40);
        let _559_v42;
        _559_v42 = _dafny.Seq.of(_556_v39);
        _558_v41 = (_558_v41).update(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ayhpbe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(613)), function (_560_i3) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }))).length), _dafny.MultiSet.FromArray(_559_v42));
        let _561_v43;
        let _nw81 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _561_v43 = _nw81;
        let _index81 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length));
        (_561_v43)[_index81] = _555_i2;
        let _index82 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length));
        (_561_v43)[_index82] = _555_i2;
        _554_v38 = (_555_i2).isEqualTo(_555_i2);
        if (p3) {
          let _562_v44;
          _562_v44 = _dafny.Seq.UnicodeFromString("bypilo");
          let _rhs123 = _dafny.Seq.Concat(_562_v44, _562_v44);
          let _rhs124 = (_dafny.ZERO).minus((_561_v43)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length))]);
          let _lhs83 = globalState;
          _562_v44 = _rhs123;
          _lhs83.f15 = _rhs124;
          let _563_v45;
          let _nw82 = new _module.C0();
          _nw82.__ctor();
          _563_v45 = _nw82;
          _554_v38 = !((_561_v43)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length))]).isEqualTo(_499_v0);
          let _564_v46;
          _564_v46 = _dafny.Map.Empty.slice().updateUnsafe(!(p3),p3);
          let _565_v47;
          _565_v47 = _dafny.MultiSet.fromElements(p3, !(_554_v38));
          let _566_v48;
          _566_v48 = _dafny.MultiSet.fromElements(new BigNumber((_565_v47).cardinality()));
          _564_v46 = (_564_v46).update(_554_v38, (_566_v48).IsDisjointFrom(_566_v48));
          let _567_v49;
          let _nw83 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.of());
          _567_v49 = _nw83;
          _567_v49 = _567_v49;
        } else {
          let _568_v50;
          _568_v50 = new _dafny.CodePoint('s'.codePointAt(0));
          let _569_v51;
          _569_v51 = _dafny.Map.Empty.slice().updateUnsafe(_568_v50,p0);
          (globalState).f4 = new BigNumber((_569_v51).length);
          let _570_v52;
          _570_v52 = _dafny.Seq.of(_554_v38, _554_v38, !(_554_v38));
          _570_v52 = _module.__default.fm27(globalState);
          let _571_v53;
          let _init18 = ((_572_p3) => function (_573_i4) {
            return _dafny.Seq.of(_572_p3, _572_p3, !(true));
          })(p3);
          let _nw84 = Array((new BigNumber(15)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw84.length); _i0_18++) {
            _nw84[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _571_v53 = _nw84;
          let _574_v54;
          _574_v54 = _dafny.Map.Empty.slice().updateUnsafe(_571_v53,_554_v38);
          let _575_v55;
          _575_v55 = _dafny.MultiSet.fromElements(_574_v54, _574_v54, _574_v54);
          _554_v38 = !((_575_v55).IsProperSubsetOf(_575_v55));
          let _576_v56;
          _576_v56 = _dafny.Seq.UnicodeFromString("xcbscyt");
          let _577_v57;
          let _nw85 = Array((new BigNumber(3)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("xfaydljs");
          _nw85[(_dafny.ONE).toNumber()] = _576_v56;
          _nw85[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_576_v56, _576_v56);
          _577_v57 = _nw85;
          let _index83 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_577_v57).length));
          (_577_v57)[_index83] = _dafny.Seq.Concat(_576_v56, _576_v56);
          let _578_v58;
          _578_v58 = _dafny.Seq.of(_576_v56);
          let _index84 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_577_v57).length));
          let _rhs125 = _554_v38;
          let _rhs126 = ((((_561_v43)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus((_561_v43)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length))]))) ? (_576_v56) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), ((_579_v50) => function (_580_i5) {
            return _579_v50;
          })(_568_v50))));
          let _rhs127 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_578_v58, _module.__default.safeIndex((_561_v43)[_module.__default.safeIndex(new BigNumber(370), new BigNumber((_561_v43).length))], new BigNumber((_578_v58).length)), _576_v56), _578_v58), _578_v58);
          let _rhs128 = !(p3) || (!(_554_v38));
          let _lhs84 = _577_v57;
          let _lhs85 = _module.__default.safeIndex(new BigNumber(994), new BigNumber((_577_v57).length));
          _554_v38 = _rhs125;
          _lhs84[_lhs85] = _rhs126;
          _578_v58 = _rhs127;
          _554_v38 = _rhs128;
          let _581_v59;
          let _nw86 = new _module.C1();
          _nw86.__ctor();
          _581_v59 = _nw86;
        }
      }
      _554_v38 = !(!(_554_v38)) || (p3);
      let _582_v60;
      _582_v60 = _module.D6.create_DC17(_554_v38, !(p3), _554_v38);
      let _583_i6;
      _583_i6 = _dafny.ZERO;
      L2: {
        while (!(_554_v38) || ((_582_v60).dtor_cf25)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_583_i6)) {
              break L2;
            }
            _583_i6 = (_583_i6).plus(_dafny.ONE);
            let _584_v61;
            _584_v61 = _dafny.Seq.UnicodeFromString("wfjnh");
            _584_v61 = _584_v61;
            _554_v38 = (_499_v0).isLessThanOrEqualTo(_module.__default.safeModuloInt(_499_v0, _499_v0));
            let _585_v62;
            let _init19 = ((_586_v0) => function (_587_i7) {
              return (_587_i7).multipliedBy(_586_v0);
            })(_499_v0);
            let _nw87 = Array((new BigNumber(8)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw87.length); _i0_19++) {
              _nw87[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _585_v62 = _nw87;
            let _index85 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_585_v62).length));
            (_585_v62)[_index85] = (_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(_499_v0)).length)).plus(_499_v0));
            let _index86 = _module.__default.safeIndex(new BigNumber(420), new BigNumber((_585_v62).length));
            (_585_v62)[_index86] = (_dafny.ZERO).minus(_499_v0);
            let _588_v63;
            _588_v63 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements((p1)[_module.__default.safeIndex((_585_v62)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_585_v62).length))], new BigNumber((p1).length))]));
            let _589_v64;
            _589_v64 = _dafny.Map.Empty.slice().updateUnsafe((_585_v62)[_module.__default.safeIndex(new BigNumber(420), new BigNumber((_585_v62).length))],_dafny.Set.fromElements(!(p3)));
            let _590_v65;
            let _nw88 = Array((new BigNumber(21)).toNumber());
            _nw88[(_dafny.ZERO).toNumber()] = (p2).Union(p2);
            _nw88[(_dafny.ONE).toNumber()] = (p2).Difference(p2);
            _nw88[(new BigNumber(2)).toNumber()] = p2;
            _nw88[(new BigNumber(3)).toNumber()] = p2;
            _nw88[(new BigNumber(4)).toNumber()] = p2;
            _nw88[(new BigNumber(5)).toNumber()] = (p2).Difference(_module.__default.fm37(globalState));
            _nw88[(new BigNumber(6)).toNumber()] = p2;
            _nw88[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements(!(p3))).Difference(p2);
            _nw88[(new BigNumber(8)).toNumber()] = (p2).Difference(_module.__default.fm37(globalState));
            _nw88[(new BigNumber(9)).toNumber()] = p2;
            _nw88[(new BigNumber(10)).toNumber()] = p2;
            _nw88[(new BigNumber(11)).toNumber()] = (p2).Difference(p2);
            _nw88[(new BigNumber(12)).toNumber()] = p2;
            _nw88[(new BigNumber(13)).toNumber()] = (p2).Difference((((_588_v63).contains(_554_v38)) ? ((_588_v63).get(_554_v38)) : (_dafny.Set.fromElements(true))));
            _nw88[(new BigNumber(14)).toNumber()] = p2;
            _nw88[(new BigNumber(15)).toNumber()] = p2;
            _nw88[(new BigNumber(16)).toNumber()] = p2;
            _nw88[(new BigNumber(17)).toNumber()] = ((_554_v38) ? (p2) : (p2));
            _nw88[(new BigNumber(18)).toNumber()] = (p2).Intersect(p2);
            _nw88[(new BigNumber(19)).toNumber()] = (p2).Union(p2);
            _nw88[(new BigNumber(20)).toNumber()] = (((_589_v64).contains(_499_v0)) ? ((_589_v64).get(_499_v0)) : (p2));
            _590_v65 = _nw88;
            let _index87 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_590_v65).length));
            (_590_v65)[_index87] = p2;
            let _index88 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_590_v65).length));
            (_590_v65)[_index88] = _dafny.Set.fromElements(_554_v38);
          }
        }
      }
      (globalState).f15 = _499_v0;
      r0 = ((_554_v38) ? (new BigNumber(786)) : (new BigNumber(111)));
      let _591_v66;
      _591_v66 = _module.D4.create_DC11(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("qpuhuc")));
      let _pat_let_tv7 = _499_v0;
      r1 = function (_source12) {
        if (_source12.is_DC12) {
          let _592___mcc_h6 = (_source12).cf16;
          let _593_cf16 = _592___mcc_h6;
          return _593_cf16;
        } else {
          let _594___mcc_h7 = (_source12).cf15;
          let _595_cf15 = _594___mcc_h7;
          return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(204), _pat_let_tv7));
        }
      }(_591_v66);
      r2 = (_dafny.ZERO).minus((_499_v0).multipliedBy(new BigNumber(798)));
      return [r0, r1, r2];
    }
    m12(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _596_v0;
      _596_v0 = _dafny.Set.fromElements(_module.__default.fm38(p1, globalState));
      let _597_v1;
      _597_v1 = new _dafny.CodePoint('q'.codePointAt(0));
      let _598_v2;
      _598_v2 = _dafny.Seq.of(_597_v1);
      let _599_v4;
      _599_v4 = _dafny.Seq.of(function () {
        let _coll30 = new _dafny.Set();
        for (const _compr_30 of (_598_v2).Elements) {
          let _600_v3 = _compr_30;
          if (_dafny.Seq.contains(_598_v2, _600_v3)) {
            _coll30.add(_600_v3);
          }
        }
        return _coll30;
      }(), _596_v0);
      let _601_i0;
      _601_i0 = _dafny.ZERO;
      L3: {
        while (_dafny.areEqual(_dafny.Seq.of(_596_v0), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(221)), ((_662_v0) => function (_663_i1) {
          return _662_v0;
        })(_596_v0)), _599_v4))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_601_i0)) {
              break L3;
            }
            _601_i0 = (_601_i0).plus(_dafny.ONE);
            let _602_v5;
            _602_v5 = new BigNumber(813);
            (globalState).f13 = ((p3) ? ((new BigNumber(72)).plus(_602_v5)) : (_602_v5));
            let _index89 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length));
            (p2)[_index89] = _602_v5;
            let _index90 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length));
            (p2)[_index90] = new BigNumber((function () {
              let _coll31 = new _dafny.Set();
              for (const _compr_31 of _dafny.IntegerRange(new BigNumber(218), new BigNumber(157))) {
                let _603_v6 = _compr_31;
                if (((new BigNumber(218)).isLessThanOrEqualTo(_603_v6)) && ((_603_v6).isLessThan(new BigNumber(157)))) {
                  _coll31.add((_603_v6).plus(new BigNumber((_dafny.MultiSet.fromElements(p1, p0)).cardinality())));
                }
              }
              return _coll31;
            }()).length);
            let _604_v7;
            _604_v7 = _module.D2.create_DC6(_602_v5);
            let _605_v8;
            _605_v8 = _module.D2.create_DC8(_604_v7);
            let _source13 = _605_v8;
            if (_source13.is_DC6) {
              let _606___mcc_h0 = (_source13).cf7;
              let _607_cf7 = _606___mcc_h0;
              let _608_v9;
              _608_v9 = true;
              let _609_v10;
              _609_v10 = _dafny.Map.Empty.slice().updateUnsafe(_602_v5,(_dafny.ZERO).minus(new BigNumber(-213)));
              _608_v9 = (_this).fm24(_602_v5, p0, _609_v10, _607_cf7, globalState);
              _608_v9 = p3;
              let _610_v11;
              let _nw89 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
              _610_v11 = _nw89;
              let _611_v12;
              _611_v12 = _dafny.Seq.of((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]);
              let _index91 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_610_v11).length));
              (_610_v11)[_index91] = _611_v12;
              let _612_v13;
              _612_v13 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))],p0);
              let _index92 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_610_v11).length));
              (_610_v11)[_index92] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), ((_613_cf7) => function (_614_i2) {
                return _613_cf7;
              })(_607_cf7)), ((p0) ? (_611_v12) : (_611_v12))), _module.__default.safeIndex((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), ((_615_cf7) => function (_616_i2) {
                return _615_cf7;
              })(_607_cf7)), ((p0) ? (_611_v12) : (_611_v12)))).length)), new BigNumber((_612_v13).length));
              let _617_v14;
              let _nw90 = Array((new BigNumber(26)).toNumber()).fill(false);
              _617_v14 = _nw90;
              let _618_v15;
              _618_v15 = _module.D0.create_DC1(_617_v14);
              let _619_v16;
              _619_v16 = _dafny.Set.fromElements((_618_v15).dtor_cf1, _617_v14);
              let _rhs129 = (_619_v16).Union(_619_v16);
              let _rhs130 = ((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]).minus(new BigNumber(185));
              let _rhs131 = (_module.__default.safeModuloInt(_607_cf7, _607_cf7)).isLessThanOrEqualTo((_dafny.ZERO).minus(_607_cf7));
              let _rhs132 = _dafny.Seq.IsProperPrefixOf(_598_v2, _dafny.Seq.Concat(_598_v2, _dafny.Seq.UnicodeFromString("sqk")));
              _619_v16 = _rhs129;
              _602_v5 = _rhs130;
              _608_v9 = _rhs131;
              _608_v9 = _rhs132;
            } else if (_source13.is_DC7) {
              let _620___mcc_h1 = (_source13).cf8;
              let _621___mcc_h2 = (_source13).cf9;
              let _622___mcc_h3 = (_source13).cf10;
              let _623_cf10 = _622___mcc_h3;
              let _624_cf9 = _621___mcc_h2;
              let _625_cf8 = _620___mcc_h1;
              (globalState).f13 = (_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]);
              let _626_v17;
              let _init20 = ((_627_p1) => function (_628_i3) {
                return _627_p1;
              })(p1);
              let _nw91 = Array((new BigNumber(13)).toNumber());
              for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw91.length); _i0_20++) {
                _nw91[_i0_20] = _init20(new BigNumber(_i0_20));
              }
              _626_v17 = _nw91;
              let _629_v18;
              _629_v18 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(114)), ((_630_v1) => function (_631_i4) {
                return _630_v1;
              })(_597_v1)), _598_v2);
              let _632_v19;
              _632_v19 = _dafny.Set.fromElements(p3);
              let _633_v20;
              _633_v20 = _dafny.Seq.of(p1);
              let _634_v21;
              _634_v21 = _module.D1.create_DC2(new BigNumber((_633_v20).length));
              let _635_v22;
              _635_v22 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))],_634_v21);
              let _index93 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_626_v17).length));
              (_626_v17)[_index93] = !(!(_629_v18).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wqrdii"), _module.__default.fm30(new BigNumber((_632_v19).length), p1, _dafny.Set.fromElements(_602_v5, new BigNumber((_635_v22).length), _602_v5), p3, globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-842)), ((_636_v1) => function (_637_i5) {
                return _636_v1;
              })(_597_v1)), _598_v2)));
              let _638_v23;
              _638_v23 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
              let _index94 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_626_v17).length));
              (_626_v17)[_index94] = (((_638_v23).contains(p0)) ? ((_638_v23).get(p0)) : (p1));
              _602_v5 = ((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]).minus(_602_v5);
              _598_v2 = (_module.D1.create_DC3(_598_v2, true)).dtor_cf3;
            } else if (_source13.is_DC5) {
              let _639___mcc_h4 = (_source13).cf6;
              let _640_cf6 = _639___mcc_h4;
              let _641_v24;
              let _nw92 = Array((new BigNumber(11)).toNumber()).fill([]);
              _641_v24 = _nw92;
              let _642_v25;
              let _nw93 = Array((new BigNumber(9)).toNumber()).fill(false);
              _642_v25 = _nw93;
              let _index95 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_641_v24).length));
              (_641_v24)[_index95] = _642_v25;
              let _643_v26;
              _643_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_602_v5);
              let _index96 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_641_v24).length));
              let _rhs133 = (p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))];
              let _rhs134 = (new BigNumber((_643_v26).length)).plus((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]);
              let _rhs135 = _642_v25;
              let _rhs136 = (_598_v2)[_module.__default.safeIndex(((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]).minus(new BigNumber((_598_v2).length)), new BigNumber((_598_v2).length))];
              let _rhs137 = _602_v5;
              let _lhs86 = globalState;
              let _lhs87 = globalState;
              let _lhs88 = _641_v24;
              let _lhs89 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_641_v24).length));
              let _lhs90 = globalState;
              _lhs86.f4 = _rhs133;
              _lhs87.f4 = _rhs134;
              _lhs88[_lhs89] = _rhs135;
              _597_v1 = _rhs136;
              _lhs90.f7 = _rhs137;
              let _644_v27;
              _644_v27 = _dafny.Seq.of((_this).fm2(_602_v5, globalState));
              let _index97 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length));
              (p2)[_index97] = (_module.__default.safeModuloInt((_this).fm2((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], globalState), (p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))])).plus((_644_v27)[_module.__default.safeIndex(_602_v5, new BigNumber((_644_v27).length))]);
              let _645_v28;
              _645_v28 = false;
              let _646_v29;
              _646_v29 = _dafny.MultiSet.fromElements(new BigNumber((_598_v2).length), (_dafny.ZERO).minus(_602_v5), _602_v5);
              let _rhs138 = _602_v5;
              let _rhs139 = (_dafny.ZERO).minus(_602_v5);
              let _rhs140 = true;
              let _rhs141 = ((new BigNumber(624)).multipliedBy(_602_v5)).minus(new BigNumber((_646_v29).cardinality()));
              let _rhs142 = _645_v28;
              let _lhs91 = globalState;
              let _lhs92 = globalState;
              _lhs91.f15 = _rhs138;
              _602_v5 = _rhs139;
              _645_v28 = _rhs140;
              _lhs92.f7 = _rhs141;
              _645_v28 = _rhs142;
              let _647_v30;
              _647_v30 = _dafny.Seq.of(p0);
              let _648_v31;
              _648_v31 = _dafny.Set.fromElements(p1, _645_v28);
              let _649_v32;
              let _650_v33;
              let _651_v34;
              let _out6;
              let _out7;
              let _out8;
              let _outcollector2 = (_this).m11((_641_v24)[_module.__default.safeIndex(new BigNumber(964), new BigNumber((_641_v24).length))], _647_v30, _648_v31, (_602_v5).isLessThanOrEqualTo(new BigNumber(-823)), globalState);
              _out6 = _outcollector2[0];
              _out7 = _outcollector2[1];
              _out8 = _outcollector2[2];
              _649_v32 = _out6;
              _650_v33 = _out7;
              _651_v34 = _out8;
            } else {
              let _652___mcc_h5 = (_source13).cf11;
              let _653_cf11 = _652___mcc_h5;
              (globalState).f4 = _602_v5;
              let _654_v35;
              _654_v35 = _dafny.Set.fromElements((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], _602_v5, new BigNumber((_dafny.Seq.UnicodeFromString("l")).length));
              let _655_v36;
              _655_v36 = _dafny.Seq.of(new BigNumber((_596_v0).length));
              let _656_v38;
              _656_v38 = _dafny.Seq.of(_654_v35, _dafny.Set.fromElements((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]), function () {
                let _coll32 = new _dafny.Set();
                for (const _compr_32 of (_655_v36).Elements) {
                  let _657_v37 = _compr_32;
                  if (_dafny.Seq.contains(_655_v36, _657_v37)) {
                    _coll32.add((_657_v37).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC20(new BigNumber(315)),true)).length)));
                  }
                }
                return _coll32;
              }(), _dafny.Set.fromElements((_this).fm2((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], globalState)));
              let _658_v39;
              _658_v39 = _dafny.Seq.of(_dafny.Set.fromElements((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], (p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]), _654_v35, (_656_v38)[_module.__default.safeIndex((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))], new BigNumber((_656_v38).length))], _654_v35);
              (globalState).f13 = new BigNumber(((_658_v39)[_module.__default.safeIndex(((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]).plus((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))]), new BigNumber((_658_v39).length))]).length);
              let _659_v40;
              _659_v40 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length))],_598_v2);
              let _660_v41;
              _660_v41 = _module.D5.create_DC14(_597_v1, p3, (_dafny.ZERO).minus(_602_v5), p0);
              _598_v2 = _dafny.Seq.Concat((((_659_v40).contains(_602_v5)) ? ((_659_v40).get(_602_v5)) : (_598_v2)), _dafny.Seq.update(_598_v2, _module.__default.safeIndex(_602_v5, new BigNumber((_598_v2).length)), (_660_v41).dtor_cf18));
              let _index98 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((p2).length));
              (p2)[_index98] = _module.__default.safeModuloInt(_602_v5, _602_v5);
            }
            let _661_v42;
            _661_v42 = true;
            _661_v42 = p0;
          }
        }
      }
      let _664_v43;
      _664_v43 = new BigNumber(-298);
      let _665_v44;
      _665_v44 = _dafny.Seq.of((_dafny.ZERO).minus(_664_v43), new BigNumber(580), _664_v43);
      let _index99 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((p2).length));
      (p2)[_index99] = new BigNumber((_665_v44).length);
      let _666_v45;
      let _nw94 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
      _666_v45 = _nw94;
      let _667_v46;
      _667_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(632),p1);
      let _668_v47;
      _668_v47 = _dafny.Seq.of(!((((_667_v46).contains(new BigNumber(264))) ? ((_667_v46).get(new BigNumber(264))) : (p1))), p3, p3);
      let _669_v48;
      _669_v48 = _module.D7.create_DC22(p3, new BigNumber((_668_v47).length), p1, new BigNumber((_665_v44).length));
      let _670_v49;
      _670_v49 = _dafny.Seq.of(_669_v48, _669_v48);
      let _671_v50;
      _671_v50 = _dafny.MultiSet.fromElements(_664_v43, new BigNumber((_670_v49).length));
      let _index100 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_666_v45).length));
      (_666_v45)[_index100] = _671_v50;
      let _index101 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((p2).length));
      let _index102 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_666_v45).length));
      let _rhs143 = _664_v43;
      let _rhs144 = _671_v50;
      let _lhs93 = p2;
      let _lhs94 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((p2).length));
      let _lhs95 = _666_v45;
      let _lhs96 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_666_v45).length));
      _lhs93[_lhs94] = _rhs143;
      _lhs95[_lhs96] = _rhs144;
      let _672_v51;
      let _nw95 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _672_v51 = _nw95;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_672_v51).length))) {
        let _673_i6 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_673_i6)) && ((_673_i6).isLessThan(new BigNumber((_672_v51).length))))) {
          (_672_v51)[(_673_i6)] = (_673_i6).multipliedBy(_664_v43);
        }
      }
      (globalState).f15 = (_this).fm2(_664_v43, globalState);
      let _index103 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((p2).length));
      (p2)[_index103] = ((p1) ? (_664_v43) : ((p2)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((p2).length))]));
      let _674_v52;
      _674_v52 = true;
      _674_v52 = p0;
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this.f24 = _module.D2.Default();
      this.f25 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f24, f25) {
      let _this = this;
      (_this).f24 = f24;
      (_this).f25 = f25;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _this.f25;
    };
    fm2(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(((!(_this.f25)) ? (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(328)))).cardinality())) : ((_dafny.ZERO).minus(new BigNumber(-540)))), new BigNumber(714));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _675_v0;
      _675_v0 = new BigNumber(-140);
      let _676_v1;
      _676_v1 = _dafny.Seq.of(_675_v0, _675_v0);
      let _677_v2;
      _677_v2 = _dafny.Map.Empty.slice().updateUnsafe(_675_v0,(_676_v1)[_module.__default.safeIndex(_675_v0, new BigNumber((_676_v1).length))]);
      let _678_v3;
      _678_v3 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f25),_675_v0);
      let _679_v4;
      _679_v4 = new _dafny.CodePoint('k'.codePointAt(0));
      let _680_v5;
      _680_v5 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), _679_v4);
      let _681_v6;
      let _init21 = ((_682_v0) => function (_683_i0) {
        return _module.D1.create_DC4(_dafny.MultiSet.FromArray(_dafny.Seq.of(_682_v0)));
      })(_675_v0);
      let _nw96 = Array((new BigNumber(18)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw96.length); _i0_21++) {
        _nw96[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _681_v6 = _nw96;
      let _684_v7;
      _684_v7 = _dafny.Map.Empty.slice().updateUnsafe(_681_v6,_this.f25);
      let _685_v8;
      _685_v8 = _dafny.Map.Empty.slice().updateUnsafe(_675_v0,_this.f25);
      let _686_v9;
      _686_v9 = _dafny.Seq.of(_this.f25, _this.f25, (((_685_v8).contains(_675_v0)) ? ((_685_v8).get(_675_v0)) : (_this.f25)));
      let _687_v10;
      _687_v10 = _dafny.MultiSet.fromElements(true, _this.f25, _this.f25, _this.f25, _this.f25);
      let _688_v11;
      let _nw97 = Array((new BigNumber(28)).toNumber());
      _nw97[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-941), _675_v0);
      _nw97[(_dafny.ONE).toNumber()] = _675_v0;
      _nw97[(new BigNumber(2)).toNumber()] = ((_this.f25) ? (_675_v0) : (new BigNumber(422)));
      _nw97[(new BigNumber(3)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(4)).toNumber()] = (new BigNumber((_677_v2).length)).multipliedBy(_675_v0);
      _nw97[(new BigNumber(5)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(6)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(7)).toNumber()] = (((_678_v3).contains(!(_this.f25))) ? ((_678_v3).get(!(_this.f25))) : (_675_v0));
      _nw97[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((p0).length), (((_680_v5).contains(_679_v4)) ? ((_680_v5).get(_679_v4)) : (_675_v0)));
      _nw97[(new BigNumber(9)).toNumber()] = new BigNumber(25);
      _nw97[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt(_675_v0, _675_v0);
      _nw97[(new BigNumber(11)).toNumber()] = new BigNumber(-127);
      _nw97[(new BigNumber(12)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(13)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(14)).toNumber()] = new BigNumber(19);
      _nw97[(new BigNumber(15)).toNumber()] = new BigNumber((_684_v7).length);
      _nw97[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_675_v0);
      _nw97[(new BigNumber(17)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(18)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(19)).toNumber()] = new BigNumber((_686_v9).length);
      _nw97[(new BigNumber(20)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(21)).toNumber()] = (_675_v0).plus(new BigNumber(652));
      _nw97[(new BigNumber(22)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(23)).toNumber()] = new BigNumber((_686_v9).length);
      _nw97[(new BigNumber(24)).toNumber()] = (((_687_v10).contains(_this.f25)) ? ((_687_v10).get(_this.f25)) : (_675_v0));
      _nw97[(new BigNumber(25)).toNumber()] = _module.__default.safeModuloInt(_675_v0, _675_v0);
      _nw97[(new BigNumber(26)).toNumber()] = _675_v0;
      _nw97[(new BigNumber(27)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_675_v0), _675_v0);
      _688_v11 = _nw97;
      let _index104 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length));
      (_688_v11)[_index104] = _675_v0;
      let _689_v12;
      let _nw98 = Array((new BigNumber(17)).toNumber());
      _nw98[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_675_v0, new BigNumber(572), _675_v0);
      _nw98[(_dafny.ONE).toNumber()] = _676_v1;
      _nw98[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), ((_690_v0) => function (_691_i1) {
        return _690_v0;
      })(_675_v0));
      _nw98[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(new BigNumber(969));
      _nw98[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_676_v1, _module.__default.safeIndex(_675_v0, new BigNumber((_676_v1).length)), _675_v0);
      _nw98[(new BigNumber(5)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_675_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_this).fm2(new BigNumber((_678_v3).length), globalState))).length));
      _nw98[(new BigNumber(7)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("nht")).length));
      _nw98[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_675_v0, _675_v0);
      _nw98[(new BigNumber(10)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(11)).toNumber()] = (_module.D6.create_DC16(_676_v1)).dtor_cf23;
      _nw98[(new BigNumber(12)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(13)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(14)).toNumber()] = _dafny.Seq.update(_676_v1, _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_676_v1).length)), _675_v0);
      _nw98[(new BigNumber(15)).toNumber()] = _676_v1;
      _nw98[(new BigNumber(16)).toNumber()] = _676_v1;
      _689_v12 = _nw98;
      let _692_v13;
      _692_v13 = _dafny.Seq.of(_689_v12, _689_v12, _689_v12);
      let _693_v14;
      let _nw99 = Array((new BigNumber(15)).toNumber());
      _nw99[(_dafny.ZERO).toNumber()] = ((_this.f25) ? (_689_v12) : (_689_v12));
      _nw99[(_dafny.ONE).toNumber()] = (_692_v13)[_module.__default.safeIndex(_675_v0, new BigNumber((_692_v13).length))];
      _nw99[(new BigNumber(2)).toNumber()] = ((_this.f25) ? (_689_v12) : (_689_v12));
      _nw99[(new BigNumber(3)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(4)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(5)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(6)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(7)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(8)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(9)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(10)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(11)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(12)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(13)).toNumber()] = _689_v12;
      _nw99[(new BigNumber(14)).toNumber()] = _689_v12;
      _693_v14 = _nw99;
      let _index105 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_693_v14).length));
      (_693_v14)[_index105] = _689_v12;
      let _index106 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length));
      let _index107 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_693_v14).length));
      let _rhs145 = new BigNumber(250);
      let _rhs146 = _675_v0;
      let _rhs147 = _689_v12;
      let _rhs148 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_module.D1.create_DC2(_675_v0)).dtor_cf2), _675_v0);
      let _lhs97 = _688_v11;
      let _lhs98 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length));
      let _lhs99 = globalState;
      let _lhs100 = _693_v14;
      let _lhs101 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_693_v14).length));
      _lhs97[_lhs98] = _rhs145;
      _lhs99.f15 = _rhs146;
      _lhs100[_lhs101] = _rhs147;
      _675_v0 = _rhs148;
      _675_v0 = _675_v0;
      let _694_v15;
      let _nw100 = Array((new BigNumber(9)).toNumber()).fill(false);
      _694_v15 = _nw100;
      let _695_v16;
      _695_v16 = _dafny.Seq.of(_694_v15);
      let _696_v17;
      _696_v17 = _module.D7.create_DC19(_695_v16);
      let _rhs149 = _this.f25;
      let _rhs150 = _679_v4;
      let _rhs151 = (_696_v17).dtor_cf28;
      let _lhs102 = _this;
      _lhs102.f25 = _rhs149;
      _679_v4 = _rhs150;
      _695_v16 = _rhs151;
      let _hi4 = ((_this.f25) ? (_675_v0) : ((_688_v11)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length))]));
      for (let _697_i2 = (new BigNumber((_677_v2).length)).minus((_dafny.ZERO).minus(_675_v0)); _697_i2.isLessThan(_hi4); _697_i2 = _697_i2.plus(_dafny.ONE)) {
        let _698_v18;
        let _nw101 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _698_v18 = _nw101;
        let _699_v19;
        _699_v19 = _dafny.Seq.of(_698_v18, _698_v18, _698_v18);
        let _700_v20;
        let _nw102 = Array((new BigNumber(29)).toNumber());
        _nw102[(_dafny.ZERO).toNumber()] = (_699_v19)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_699_v19).length))];
        _nw102[(_dafny.ONE).toNumber()] = _698_v18;
        _nw102[(new BigNumber(2)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(3)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(4)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(5)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(6)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(7)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(8)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(9)).toNumber()] = (_699_v19)[_module.__default.safeIndex((_688_v11)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length))], new BigNumber((_699_v19).length))];
        _nw102[(new BigNumber(10)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(11)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(12)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(13)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(14)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(15)).toNumber()] = ((true) ? (_698_v18) : (_698_v18));
        _nw102[(new BigNumber(16)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(17)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(18)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(19)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(20)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(21)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(22)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(23)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(24)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(25)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(26)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(27)).toNumber()] = _698_v18;
        _nw102[(new BigNumber(28)).toNumber()] = _698_v18;
        _700_v20 = _nw102;
        let _index108 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_700_v20).length));
        (_700_v20)[_index108] = _698_v18;
        let _701_v21;
        _701_v21 = _module.D2.create_DC6((_688_v11)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length))]);
        let _index109 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_700_v20).length));
        let _rhs152 = _698_v18;
        let _rhs153 = _this.f25;
        let _rhs154 = true;
        let _rhs155 = _701_v21;
        let _lhs103 = _700_v20;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(447), new BigNumber((_700_v20).length));
        let _lhs105 = _this;
        let _lhs106 = _this;
        _lhs103[_lhs104] = _rhs152;
        _lhs105.f25 = _rhs153;
        _lhs106.f25 = _rhs154;
        _701_v21 = _rhs155;
        (globalState).f4 = _675_v0;
        (_this).f25 = _this.f25;
        (_this).f25 = !((_675_v0).isEqualTo(new BigNumber((p0).length))) || (_this.f25);
      }
      _675_v0 = (_688_v11)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((_688_v11).length))];
      (_this).f25 = _this.f25;
      return;
    }
    m10(globalState) {
      let _this = this;
      let r0 = _dafny.Set.Empty;
      let r1 = [];
      let r2 = _dafny.Seq.of();
      let _702_v0;
      let _init22 = function (_703_i0) {
        return true;
      };
      let _nw103 = Array((_dafny.ONE).toNumber());
      for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw103.length); _i0_22++) {
        _nw103[_i0_22] = _init22(new BigNumber(_i0_22));
      }
      _702_v0 = _nw103;
      let _index110 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
      (_702_v0)[_index110] = !(((_this.f25) ? (_this.f25) : (_this.f25)));
      let _704_v1;
      _704_v1 = new BigNumber(199);
      let _705_v2;
      _705_v2 = _dafny.Set.fromElements(_704_v1);
      let _706_v3;
      _706_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_704_v1, _704_v1, new BigNumber((_705_v2).length), _704_v1, _704_v1)).Union(_705_v2),_704_v1);
      let _707_v4;
      _707_v4 = new _dafny.CodePoint('v'.codePointAt(0));
      let _708_v5;
      _708_v5 = _dafny.Seq.of(_this.f25);
      let _709_v6;
      _709_v6 = _module.D8.create_DC23(_dafny.Seq.of(_this.f25, _this.f25));
      let _index111 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
      let _rhs156 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_708_v5, (_709_v6).dtor_cf39), _dafny.Seq.of(_this.f25, _this.f25));
      let _rhs157 = ((_dafny.ZERO).minus(_704_v1)).plus(_704_v1);
      let _rhs158 = !(_module.__default.safeModuloInt(_704_v1, _704_v1)).isEqualTo(new BigNumber(-740));
      let _rhs159 = _706_v3;
      let _rhs160 = _707_v4;
      let _lhs107 = _702_v0;
      let _lhs108 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
      let _lhs109 = globalState;
      let _lhs110 = _this;
      _lhs107[_lhs108] = _rhs156;
      _lhs109.f4 = _rhs157;
      _lhs110.f25 = _rhs158;
      _706_v3 = _rhs159;
      _707_v4 = _rhs160;
      let _710_v7;
      let _nw104 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _710_v7 = _nw104;
      let _711_v8;
      let _nw105 = Array((new BigNumber(18)).toNumber());
      _nw105[(_dafny.ZERO).toNumber()] = _710_v7;
      _nw105[(_dafny.ONE).toNumber()] = _710_v7;
      _nw105[(new BigNumber(2)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(3)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(4)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(5)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(6)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(7)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(8)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(9)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(10)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(11)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(12)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(13)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(14)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(15)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(16)).toNumber()] = _710_v7;
      _nw105[(new BigNumber(17)).toNumber()] = _710_v7;
      _711_v8 = _nw105;
      _711_v8 = _711_v8;
      let _712_v9;
      _712_v9 = _dafny.Map.Empty.slice().updateUnsafe(_702_v0,_702_v0);
      let _713_v10;
      _713_v10 = _dafny.Seq.of(_704_v1, new BigNumber((_712_v9).length));
      (_this).f25 = !_dafny.Seq.contains(_713_v10, new BigNumber((((_this.f25) ? (_713_v10) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(159)), ((_714_v1) => function (_715_i1) {
        return _714_v1;
      })(_704_v1))))).length));
      let _716_v11;
      _716_v11 = _dafny.MultiSet.fromElements(_this.f25);
      let _717_v12;
      _717_v12 = _dafny.Seq.UnicodeFromString("bpoxflb");
      let _718_v13;
      _718_v13 = _module.D7.create_DC21((((_716_v11).contains(_this.f25)) ? ((_716_v11).get(_this.f25)) : (new BigNumber((_717_v12).length))), _704_v1, false, !((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]) || (_this.f25), false);
      let _source14 = _718_v13;
      if (_source14.is_DC20) {
        let _719___mcc_h0 = (_source14).cf29;
        let _720_cf29 = _719___mcc_h0;
        if ((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]) {
          let _721_v14;
          let _nw106 = new _module.C1();
          _nw106.__ctor();
          _721_v14 = _nw106;
          let _index112 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length));
          (_710_v7)[_index112] = _module.__default.safeModuloInt(_704_v1, new BigNumber((_708_v5).length));
          let _index113 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          let _index114 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          let _rhs161 = _this.f25;
          let _rhs162 = (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))];
          let _rhs163 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), ((_722_v4) => function (_723_i2) {
            return _722_v4;
          })(_707_v4)), _dafny.Seq.UnicodeFromString("tswns"))).length);
          let _rhs164 = _710_v7;
          let _rhs165 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm39(_704_v1, _707_v4, _713_v10, !(_this.f25), globalState), _dafny.Seq.Concat(_dafny.Seq.of(_720_cf29, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), ((_724_v2) => function (_725_i3) {
            return _module.D11.create_DC28(_724_v2);
          })(_705_v2))).length)), new BigNumber((_717_v12).length)), _713_v10));
          let _lhs111 = _702_v0;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          let _lhs113 = _this;
          let _lhs114 = _710_v7;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length));
          let _lhs116 = _702_v0;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          _lhs111[_lhs112] = _rhs161;
          _lhs113.f25 = _rhs162;
          _lhs114[_lhs115] = _rhs163;
          r1 = _rhs164;
          _lhs116[_lhs117] = _rhs165;
          (_this).f25 = false;
          let _726_v15;
          _726_v15 = _module.D1.create_DC2(new BigNumber((_713_v10).length));
          let _pat_let_tv8 = _704_v1;
          let _727_v16;
          let _nw107 = Array((new BigNumber(20)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = _726_v15;
          _nw107[(_dafny.ONE).toNumber()] = _726_v15;
          _nw107[(new BigNumber(2)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(3)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(4)).toNumber()] = _module.D1.create_DC2(_704_v1);
          _nw107[(new BigNumber(5)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(6)).toNumber()] = _module.D1.create_DC2(_720_cf29);
          _nw107[(new BigNumber(7)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(8)).toNumber()] = _module.D1.create_DC2((_713_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_717_v12, _module.__default.safeIndex(new BigNumber((_705_v2).length), new BigNumber((_717_v12).length)), _707_v4)).length), new BigNumber((_713_v10).length))]);
          _nw107[(new BigNumber(9)).toNumber()] = _module.D1.create_DC2(_720_cf29);
          _nw107[(new BigNumber(10)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(11)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(12)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(13)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(14)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(15)).toNumber()] = function (_pat_let2_0) {
            return function (_728_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_729_dt__update_hcf2_h0) {
                  return _module.D1.create_DC2(_729_dt__update_hcf2_h0);
                }(_pat_let3_0);
              }(_pat_let_tv8);
            }(_pat_let2_0);
          }(_module.D1.create_DC2((_710_v7)[_module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length))]));
          _nw107[(new BigNumber(16)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(17)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(18)).toNumber()] = _726_v15;
          _nw107[(new BigNumber(19)).toNumber()] = _module.__default.fm40((_710_v7)[_module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length))], (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))], _this.f25, _709_v6, globalState);
          _727_v16 = _nw107;
          _727_v16 = _727_v16;
          let _index116 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_710_v7).length));
          (_710_v7)[_index116] = new BigNumber(603);
        } else {
          let _730_v17;
          _730_v17 = _dafny.Map.Empty.slice().updateUnsafe(_705_v2,_this.f25);
          _713_v10 = _dafny.Seq.Concat((((((_730_v17).contains(_705_v2)) ? ((_730_v17).get(_705_v2)) : (_module.__default.fm26(globalState)))) ? (_713_v10) : (_713_v10)), _dafny.Seq.of(_704_v1));
          let _731_v18;
          _731_v18 = _dafny.Map.Empty.slice().updateUnsafe(_704_v1,(_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]);
          let _732_v19;
          _732_v19 = _dafny.Map.Empty.slice().updateUnsafe(_717_v12,_720_cf29);
          _731_v18 = (_731_v18).update(_module.__default.safeModuloInt(new BigNumber((_732_v19).length), _704_v1), true);
          _708_v5 = _module.__default.fm27(globalState);
          let _index117 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_710_v7).length));
          (_710_v7)[_index117] = _704_v1;
          let _index118 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_710_v7).length));
          (_710_v7)[_index118] = _720_cf29;
          let _index119 = _module.__default.safeIndex(new BigNumber(927), new BigNumber((_710_v7).length));
          (_710_v7)[_index119] = _module.__default.safeModuloInt(_720_cf29, _704_v1);
        }
        let _733_v20;
        _733_v20 = _dafny.Set.fromElements(_module.__default.fm41(_707_v4, globalState));
        let _734_v21;
        _734_v21 = _dafny.Map.Empty.slice().updateUnsafe(_733_v20,_707_v4);
        _734_v21 = (_734_v21).update(_733_v20, _707_v4);
        let _735_v22;
        let _nw108 = Array((new BigNumber(12)).toNumber()).fill(_module.D7.Default());
        _735_v22 = _nw108;
        let _736_v23;
        _736_v23 = _module.D3.create_DC10(new BigNumber(808), _720_cf29);
        let _737_v24;
        _737_v24 = _dafny.Map.Empty.slice().updateUnsafe(_735_v22,_736_v23);
        (_this).f25 = (new BigNumber((_737_v24).length)).isEqualTo(_704_v1);
        (globalState).f4 = (new BigNumber((_716_v11).cardinality())).minus(_704_v1);
      } else if (_source14.is_DC21) {
        let _738___mcc_h1 = (_source14).cf30;
        let _739___mcc_h2 = (_source14).cf31;
        let _740___mcc_h3 = (_source14).cf32;
        let _741___mcc_h4 = (_source14).cf33;
        let _742___mcc_h5 = (_source14).cf34;
        let _743_cf34 = _742___mcc_h5;
        let _744_cf33 = _741___mcc_h4;
        let _745_cf32 = _740___mcc_h3;
        let _746_cf31 = _739___mcc_h2;
        let _747_cf30 = _738___mcc_h1;
        let _pat_let_tv9 = _705_v2;
        let _source15 = function (_pat_let4_0) {
          return function (_748_dt__update__tmp_h1) {
            return function (_pat_let5_0) {
              return function (_749_dt__update_hcf45_h0) {
                return _module.D11.create_DC28(_749_dt__update_hcf45_h0);
              }(_pat_let5_0);
            }(_pat_let_tv9);
          }(_pat_let4_0);
        }(_module.D11.create_DC28(_705_v2));
        if (_source15.is_DC29) {
          (globalState).f4 = (_747_cf30).multipliedBy(_747_cf30);
          let _index120 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          (_702_v0)[_index120] = (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))];
          let _750_v25;
          let _nw109 = new _module.C2();
          _nw109.__ctor();
          _750_v25 = _nw109;
          let _751_v26;
          _751_v26 = _dafny.Map.Empty.slice().updateUnsafe(_704_v1,_745_cf32);
          let _752_v27;
          _752_v27 = _module.D7.create_DC20(new BigNumber(316));
          let _rhs166 = _746_cf31;
          let _rhs167 = !_dafny.areEqual(_dafny.Seq.of(true, !(_this.f25)), _dafny.Seq.update(_708_v5, _module.__default.safeIndex((_752_v27).dtor_cf29, new BigNumber((_708_v5).length)), (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]));
          let _rhs168 = _751_v26;
          let _rhs169 = !(_this.f25) || (_743_cf34);
          let _rhs170 = _745_cf32;
          let _lhs118 = globalState;
          _lhs118.f4 = _rhs166;
          _745_cf32 = _rhs167;
          _751_v26 = _rhs168;
          _745_cf32 = _rhs169;
          _745_cf32 = _rhs170;
        } else if (_source15.is_DC30) {
          let _753___mcc_h11 = (_source15).cf46;
          let _754___mcc_h12 = (_source15).cf47;
          let _755_cf47 = _754___mcc_h12;
          let _756_cf46 = _753___mcc_h11;
          let _757_v28;
          _757_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-980),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_743_cf34,_704_v1)).length));
          let _758_v29;
          _758_v29 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_745_cf32,(((_757_v28).contains(_746_cf31)) ? ((_757_v28).get(_746_cf31)) : (_704_v1)))).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), ((_759_v4) => function (_760_i4) {
            return _759_v4;
          })(_707_v4)));
          let _761_v30;
          _761_v30 = _module.D1.create_DC2(_747_cf30);
          let _762_v31;
          _762_v31 = _dafny.Set.fromElements(_761_v30);
          let _763_v32;
          _763_v32 = _dafny.Map.Empty.slice().updateUnsafe((_758_v29).Merge(_758_v29),_762_v31);
          let _764_v33;
          _764_v33 = _dafny.Map.Empty.slice().updateUnsafe(_755_cf47,_762_v31);
          let _765_v34;
          _765_v34 = _dafny.MultiSet.fromElements(_761_v30);
          _763_v32 = (_763_v32).update(_758_v29, (((_764_v33).contains(_747_cf30)) ? ((_764_v33).get(_747_cf30)) : (function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of (_765_v34).Elements) {
              let _766_v35 = _compr_33;
              if ((_765_v34).contains(_766_v35)) {
                _coll33.add(_766_v35);
              }
            }
            return _coll33;
          }())));
          (globalState).f4 = new BigNumber(897);
          let _767_v36;
          let _nw110 = new _module.C2();
          _nw110.__ctor();
          _767_v36 = _nw110;
          let _768_v37;
          _768_v37 = _dafny.Map.Empty.slice().updateUnsafe((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))],true);
          let _769_v38;
          _769_v38 = _dafny.Map.Empty.slice().updateUnsafe(_768_v37,_702_v0);
          let _770_v39;
          let _init23 = ((_771_v5) => function (_772_i5) {
            return _771_v5;
          })(_708_v5);
          let _nw111 = Array((new BigNumber(7)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw111.length); _i0_23++) {
            _nw111[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _770_v39 = _nw111;
          let _index121 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_770_v39).length));
          (_770_v39)[_index121] = _dafny.Seq.Concat(_dafny.Seq.of(_this.f25, true), _708_v5);
          let _index122 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_770_v39).length));
          let _rhs171 = _769_v38;
          let _rhs172 = _module.__default.fm27(globalState);
          let _lhs119 = _770_v39;
          let _lhs120 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_770_v39).length));
          _769_v38 = _rhs171;
          _lhs119[_lhs120] = _rhs172;
        } else if (_source15.is_DC28) {
          let _773___mcc_h13 = (_source15).cf45;
          let _774_cf45 = _773___mcc_h13;
          (globalState).f15 = (_dafny.ZERO).minus(_704_v1);
          (globalState).f13 = _704_v1;
          _745_cf32 = _this.f25;
          (globalState).f7 = _746_cf31;
        } else {
          let _775___mcc_h14 = (_source15).cf48;
          let _776_cf48 = _775___mcc_h14;
          let _777_v40;
          _777_v40 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_dafny.Map.Empty.slice().updateUnsafe(_744_cf33,_744_cf33));
          let _778_v41;
          _778_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm1(_777_v40, new BigNumber((_717_v12).length), globalState),_743_cf34);
          let _779_v42;
          _779_v42 = _dafny.Map.Empty.slice().updateUnsafe(_744_cf33,_778_v41);
          _779_v42 = (_779_v42).update(!((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]), _778_v41);
          (_this).f25 = (_708_v5)[_module.__default.safeIndex(_746_cf31, new BigNumber((_708_v5).length))];
          let _780_v43;
          _780_v43 = _dafny.MultiSet.fromElements(new BigNumber((_708_v5).length));
          let _781_v44;
          _781_v44 = _module.D1.create_DC4(_780_v43);
          _781_v44 = _module.__default.fm42(_747_cf30, new BigNumber((function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of _dafny.IntegerRange(new BigNumber(376), new BigNumber(392))) {
              let _782_v45 = _compr_34;
              if (((new BigNumber(376)).isLessThanOrEqualTo(_782_v45)) && ((_782_v45).isLessThan(new BigNumber(392)))) {
                _coll34.add((_782_v45).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_747_cf30,_747_cf30)).length)));
              }
            }
            return _coll34;
          }()).length), globalState);
          let _783_v46;
          _783_v46 = _module.D0.create_DC0(_702_v0);
          let _784_v47;
          _784_v47 = _dafny.Seq.of(_783_v46);
          (_this).f25 = _dafny.areEqual(_dafny.Seq.Concat(_784_v47, _784_v47), _784_v47);
        }
        _704_v1 = new BigNumber(698);
        let _785_v48;
        _785_v48 = _dafny.MultiSet.fromElements(_705_v2, _705_v2, _705_v2);
        let _source16 = _module.D5.create_DC13((_785_v48).Union(_dafny.MultiSet.fromElements(_705_v2, _705_v2, _705_v2)));
        if (_source16.is_DC14) {
          let _786___mcc_h15 = (_source16).cf18;
          let _787___mcc_h16 = (_source16).cf19;
          let _788___mcc_h17 = (_source16).cf20;
          let _789___mcc_h18 = (_source16).cf21;
          let _790_cf21 = _789___mcc_h18;
          let _791_cf20 = _788___mcc_h17;
          let _792_cf19 = _787___mcc_h16;
          let _793_cf18 = _786___mcc_h15;
          _790_cf21 = _745_cf32;
          let _794_v49;
          let _nw112 = new _module.C0();
          _nw112.__ctor();
          _794_v49 = _nw112;
          let _795_v50;
          let _nw113 = new _module.C2();
          _nw113.__ctor();
          _795_v50 = _nw113;
          (globalState).f4 = (((_716_v11).contains(_792_cf19)) ? ((_716_v11).get(_792_cf19)) : (((_792_cf19) ? (new BigNumber((_705_v2).length)) : (_746_cf31))));
        } else if (_source16.is_DC13) {
          let _796___mcc_h19 = (_source16).cf17;
          let _797_cf17 = _796___mcc_h19;
          let _798_v51;
          _798_v51 = _dafny.MultiSet.fromElements(_746_cf31);
          let _index123 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length));
          (_710_v7)[_index123] = (new BigNumber((_798_v51).cardinality())).plus(new BigNumber((_717_v12).length));
          let _799_v52;
          _799_v52 = _module.D5.create_DC13(_module.__default.fm43((_dafny.ZERO).minus(_746_cf31), globalState));
          let _800_v53;
          _800_v53 = _dafny.Map.Empty.slice().updateUnsafe((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))],_717_v12);
          let _801_v54;
          _801_v54 = _dafny.Seq.of(_717_v12, _717_v12, _717_v12, _717_v12, _717_v12);
          let _802_v55;
          _802_v55 = _dafny.Seq.of(_dafny.Seq.of(_717_v12, _717_v12, (((_800_v53).contains(!((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]))) ? ((_800_v53).get(!((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]))) : (_717_v12)), _717_v12), _dafny.Seq.update(_801_v54, _module.__default.safeIndex(_704_v1, new BigNumber((_801_v54).length)), _717_v12), _801_v54);
          let _index124 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length));
          let _rhs173 = (_747_cf30).multipliedBy(_704_v1);
          let _rhs174 = new BigNumber(((_802_v55)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).fm2(_746_cf31, globalState)), new BigNumber((_802_v55).length))]).length);
          let _rhs175 = _799_v52;
          let _rhs176 = true;
          let _rhs177 = _716_v11;
          let _lhs121 = _710_v7;
          let _lhs122 = _module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length));
          _lhs121[_lhs122] = _rhs173;
          _704_v1 = _rhs174;
          _799_v52 = _rhs175;
          _744_cf33 = _rhs176;
          _716_v11 = _rhs177;
          let _803_v56;
          let _nw114 = Array((new BigNumber(6)).toNumber()).fill(_dafny.MultiSet.Empty);
          _803_v56 = _nw114;
          let _804_v57;
          _804_v57 = _dafny.Set.fromElements(_this.f25, _745_cf32, (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]);
          let _805_v58;
          _805_v58 = _module.D12.create_DC32(_804_v57);
          let _index125 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_803_v56).length));
          (_803_v56)[_index125] = (_dafny.MultiSet.fromElements(_804_v57, (_805_v58).dtor_cf49, _804_v57, _804_v57, _804_v57)).update(_804_v57, _module.__default.abs(_746_cf31));
          let _806_v59;
          _806_v59 = _dafny.MultiSet.fromElements(_804_v57);
          let _807_v60;
          _807_v60 = _dafny.Map.Empty.slice().updateUnsafe(_745_cf32,_747_cf30);
          let _index126 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_803_v56).length));
          let _index127 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          let _rhs178 = (_806_v59).Union((_806_v59).Intersect(_806_v59));
          let _rhs179 = _module.__default.fm26(globalState);
          let _rhs180 = (_713_v10)[_module.__default.safeIndex((((_807_v60).contains(!(!(_743_cf34)))) ? ((_807_v60).get(!(!(_743_cf34)))) : (_747_cf30)), new BigNumber((_713_v10).length))];
          let _lhs123 = _803_v56;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_803_v56).length));
          let _lhs125 = _702_v0;
          let _lhs126 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
          let _lhs127 = globalState;
          _lhs123[_lhs124] = _rhs178;
          _lhs125[_lhs126] = _rhs179;
          _lhs127.f7 = _rhs180;
          (globalState).f15 = _746_cf31;
          let _808_v61;
          _808_v61 = _dafny.Map.Empty.slice().updateUnsafe(_704_v1,_744_cf33);
          let _809_v62;
          _809_v62 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm35(globalState),_746_cf31);
          let _810_v63;
          let _nw115 = Array((new BigNumber(10)).toNumber());
          _nw115[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_713_v10).length), _704_v1);
          _nw115[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_808_v61).length), new BigNumber((_809_v62).length));
          _nw115[(new BigNumber(2)).toNumber()] = _747_cf30;
          _nw115[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_710_v7)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length))]);
          _nw115[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_808_v61).length), _747_cf30);
          _nw115[(new BigNumber(5)).toNumber()] = _747_cf30;
          _nw115[(new BigNumber(6)).toNumber()] = (_710_v7)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length))];
          _nw115[(new BigNumber(7)).toNumber()] = (_710_v7)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length))];
          _nw115[(new BigNumber(8)).toNumber()] = _746_cf31;
          _nw115[(new BigNumber(9)).toNumber()] = (_this).fm2((_710_v7)[_module.__default.safeIndex(new BigNumber(563), new BigNumber((_710_v7).length))], globalState);
          _810_v63 = _nw115;
          r1 = _810_v63;
        } else {
          let _811___mcc_h20 = (_source16).cf22;
          let _812_cf22 = _811___mcc_h20;
          let _813_v64;
          _813_v64 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_717_v12).length));
          let _814_v65;
          _814_v65 = _813_v64;
          let _815_v66;
          _815_v66 = _dafny.MultiSet.fromElements(_813_v64, (_814_v65), (_dafny.Map.Empty.slice().updateUnsafe((_708_v5)[_module.__default.safeIndex(_747_cf30, new BigNumber((_708_v5).length))],_704_v1)).Merge(_module.__default.fm44(_747_cf30, true, globalState)), (_813_v64).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm26(globalState),_747_cf30)));
          (globalState).f4 = (((_815_v66).contains(_813_v64)) ? ((_815_v66).get(_813_v64)) : (_module.__default.safeModuloInt(_704_v1, _704_v1)));
          (globalState).f7 = _704_v1;
          _710_v7 = (_module.D2.create_DC7((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))], _716_v11, _710_v7)).dtor_cf10;
          _743_cf34 = _dafny.Seq.contains(_713_v10, _746_cf31);
        }
        _716_v11 = (_716_v11).update(true, _module.__default.abs(new BigNumber(643)));
      } else if (_source14.is_DC22) {
        let _816___mcc_h6 = (_source14).cf35;
        let _817___mcc_h7 = (_source14).cf36;
        let _818___mcc_h8 = (_source14).cf37;
        let _819___mcc_h9 = (_source14).cf38;
        let _820_cf38 = _819___mcc_h9;
        let _821_cf37 = _818___mcc_h8;
        let _822_cf36 = _817___mcc_h7;
        let _823_cf35 = _816___mcc_h6;
        (_this).f25 = (_716_v11).IsSubsetOf(_716_v11);
        let _824_v67;
        _824_v67 = _dafny.Set.fromElements((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]);
        let _825_v68;
        _825_v68 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_824_v67).length)),_822_cf36);
        let _826_v69;
        _826_v69 = _module.D3.create_DC10(_704_v1, new BigNumber((_825_v68).length));
        let _827_v70;
        _827_v70 = _dafny.Map.Empty.slice().updateUnsafe(_717_v12,_dafny.Seq.UnicodeFromString("de"));
        let _pat_let_tv10 = _704_v1;
        let _828_v71;
        _828_v71 = _dafny.Seq.of(_826_v69, function (_pat_let6_0) {
          return function (_829_dt__update__tmp_h2) {
            return function (_pat_let7_0) {
              return function (_830_dt__update_hcf13_h0) {
                return _module.D3.create_DC10(_830_dt__update_hcf13_h0, (_829_dt__update__tmp_h2).dtor_cf14);
              }(_pat_let7_0);
            }(_pat_let_tv10);
          }(_pat_let6_0);
        }(_826_v69), _826_v69, _826_v69, _module.D3.create_DC10(new BigNumber((_827_v70).length), new BigNumber(-373)));
        let _831_v72;
        _831_v72 = _dafny.Map.Empty.slice().updateUnsafe((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))],_828_v71);
        let _832_v73;
        _832_v73 = _dafny.Map.Empty.slice().updateUnsafe(_823_cf35,_this.f25);
        _831_v72 = _dafny.Map.Empty.slice().updateUnsafe((((_832_v73).contains((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))])) ? ((_832_v73).get((_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))])) : (_823_cf35)),_dafny.Seq.update(_dafny.Seq.of(_826_v69, _826_v69), _module.__default.safeIndex(_822_cf36, new BigNumber((_dafny.Seq.of(_826_v69, _826_v69)).length)), _826_v69));
        _821_cf37 = (_705_v2).IsSubsetOf(_705_v2);
        let _index128 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        (_702_v0)[_index128] = _module.__default.fm26(globalState);
      } else {
        let _833___mcc_h10 = (_source14).cf28;
        let _834_cf28 = _833___mcc_h10;
        let _835_v74;
        _835_v74 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(655),new _dafny.CodePoint('g'.codePointAt(0)));
        let _index129 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        let _index130 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        let _rhs181 = _this.f25;
        let _rhs182 = _835_v74;
        let _rhs183 = (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))];
        let _rhs184 = (_704_v1).minus((_dafny.ZERO).minus(_704_v1));
        let _rhs185 = _710_v7;
        let _lhs128 = _702_v0;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        let _lhs130 = _702_v0;
        let _lhs131 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        _lhs128[_lhs129] = _rhs181;
        _835_v74 = _rhs182;
        _lhs130[_lhs131] = _rhs183;
        _704_v1 = _rhs184;
        _710_v7 = _rhs185;
        _708_v5 = _708_v5;
        (_this).f25 = ((true) ? (_this.f25) : ((_this.f25) === (_this.f25)));
        let _836_v75;
        let _nw116 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _836_v75 = _nw116;
        let _index131 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_836_v75).length));
        (_836_v75)[_index131] = (_835_v74).Merge(_835_v74);
        let _index132 = _module.__default.safeIndex(new BigNumber(530), new BigNumber((_836_v75).length));
        (_836_v75)[_index132] = _835_v74;
      }
      let _837_v76;
      _837_v76 = _dafny.Set.fromElements(_710_v7);
      let _838_v77;
      _838_v77 = _dafny.Map.Empty.slice().updateUnsafe(_717_v12,_704_v1);
      let _hi5 = (((_838_v77).contains(_717_v12)) ? ((_838_v77).get(_717_v12)) : (_704_v1));
      for (let _839_i6 = new BigNumber((_837_v76).length); _839_i6.isLessThan(_hi5); _839_i6 = _839_i6.plus(_dafny.ONE)) {
        let _index133 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        (_702_v0)[_index133] = (_718_v13).dtor_cf34;
        (_this).f25 = !((_839_i6).plus(new BigNumber(211))).isEqualTo(_839_i6);
        let _840_v78;
        let _nw117 = new _module.C2();
        _nw117.__ctor();
        _840_v78 = _nw117;
        let _841_v79;
        _841_v79 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_840_v78);
        let _842_v80;
        _842_v80 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25);
        let _843_v81;
        _843_v81 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_842_v80);
        let _index134 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        let _rhs186 = (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))];
        let _rhs187 = true;
        let _rhs188 = (((_841_v79).contains((_840_v78).fm1(_843_v81, _839_i6, globalState))) ? ((_841_v79).get((_840_v78).fm1(_843_v81, _839_i6, globalState))) : (_840_v78));
        let _lhs132 = _702_v0;
        let _lhs133 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        let _lhs134 = _this;
        _lhs132[_lhs133] = _rhs186;
        _lhs134.f25 = _rhs187;
        _840_v78 = _rhs188;
        let _index135 = _module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length));
        (_702_v0)[_index135] = (_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))];
      }
      let _844_v82;
      let _nw118 = new _module.C1();
      _nw118.__ctor();
      _844_v82 = _nw118;
      r0 = _dafny.Set.fromElements(_704_v1, new BigNumber((_dafny.Seq.UnicodeFromString("pvydf")).length));
      let _845_v83;
      _845_v83 = _dafny.MultiSet.fromElements(_704_v1);
      r1 = ((((_module.__default.fm42(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-269)), function (_846_i7) {
        return new BigNumber(23);
      })).length), _704_v1, globalState)).dtor_cf5).IsDisjointFrom(_845_v83)) ? (_710_v7) : (_710_v7));
      let _847_v84;
      _847_v84 = _dafny.Map.Empty.slice().updateUnsafe(_704_v1,(_702_v0)[_module.__default.safeIndex(new BigNumber(361), new BigNumber((_702_v0).length))]);
      let _848_v86;
      _848_v86 = _dafny.Seq.of(_847_v84, function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of _dafny.IntegerRange(new BigNumber(11), new BigNumber(121))) {
          let _849_v85 = _compr_35;
          if (((new BigNumber(11)).isLessThanOrEqualTo(_849_v85)) && ((_849_v85).isLessThan(new BigNumber(121)))) {
            _coll35.push([(_849_v85).minus(_704_v1),!(true)]);
          }
        }
        return _coll35;
      }(), _dafny.Map.Empty.slice().updateUnsafe(_704_v1,_this.f25), _847_v84, _847_v84);
      let _850_v87;
      _850_v87 = _module.D14.create_DC35(_848_v86);
      r2 = (_850_v87).dtor_cf52;
      return [r0, r1, r2];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(p0, globalState) {
      let _this = this;
      return ((_module.D5.create_DC13(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-472), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.of(false, true)).length))))).dtor_cf17).Intersect((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(-140)), function () {
        let _coll36 = new _dafny.Set();
        for (const _compr_36 of _dafny.IntegerRange(new BigNumber(470), new BigNumber(446))) {
          let _851_v0 = _compr_36;
          if (((new BigNumber(470)).isLessThanOrEqualTo(_851_v0)) && ((_851_v0).isLessThan(new BigNumber(446)))) {
            _coll36.add((_851_v0).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, true, false)).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_852_i0) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })).length))).length)));
          }
        }
        return _coll36;
      }())).Union(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber(710)))));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(new BigNumber(604), (_dafny.ZERO).minus(new BigNumber(-912)), new BigNumber(783))).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(241), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),false)).length), new BigNumber(-235))).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("ecngi")).length)));
    };
    fm23(p0, p1, p2, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(((!(true)) ? (new BigNumber(909)) : (new BigNumber(731))), ((true) ? (new BigNumber((_dafny.Seq.UnicodeFromString("xueolns")).length)) : (new BigNumber((_dafny.Seq.UnicodeFromString("apewrwfc")).length))));
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _853_v0;
      let _nw119 = Array((new BigNumber(5)).toNumber());
      _nw119[(_dafny.ZERO).toNumber()] = p2;
      _nw119[(_dafny.ONE).toNumber()] = p2;
      _nw119[(new BigNumber(2)).toNumber()] = p2;
      _nw119[(new BigNumber(3)).toNumber()] = p2;
      _nw119[(new BigNumber(4)).toNumber()] = p2;
      _853_v0 = _nw119;
      let _854_v1;
      _854_v1 = _module.D0.create_DC1(_853_v0);
      let _source17 = _854_v1;
      if (_source17.is_DC1) {
        let _855___mcc_h0 = (_source17).cf1;
        let _856_cf1 = _855___mcc_h0;
        _856_cf1 = _856_cf1;
        let _857_v2;
        _857_v2 = _dafny.Set.fromElements(p0);
        let _858_v3;
        _858_v3 = _dafny.Seq.of(_857_v2);
        let _859_v4;
        _859_v4 = _dafny.Map.Empty.slice().updateUnsafe(p2,_858_v3);
        let _860_v5;
        _860_v5 = _dafny.Seq.of(p2);
        let _861_v6;
        _861_v6 = _dafny.Seq.UnicodeFromString("wnhu");
        _859_v4 = (_859_v4).update((_860_v5)[_module.__default.safeIndex(new BigNumber((_861_v6).length), new BigNumber((_860_v5).length))], ((p2) ? (_858_v3) : (_858_v3)));
        (globalState).f15 = ((p0).plus(new BigNumber((_861_v6).length))).minus(new BigNumber((_861_v6).length));
        let _862_v7;
        _862_v7 = true;
        _862_v7 = p2;
      } else {
        let _863___mcc_h1 = (_source17).cf0;
        let _864_cf0 = _863___mcc_h1;
        let _865_v8;
        let _init24 = ((_866_p2) => function (_867_i0) {
          return (_dafny.Set.fromElements(_866_p2)).Union(_dafny.Set.fromElements(_866_p2));
        })(p2);
        let _nw120 = Array((new BigNumber(15)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw120.length); _i0_24++) {
          _nw120[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _865_v8 = _nw120;
        let _868_v9;
        _868_v9 = _dafny.Set.fromElements(p2, p2);
        let _index136 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_865_v8).length));
        (_865_v8)[_index136] = _868_v9;
        let _869_v10;
        _869_v10 = _dafny.Map.Empty.slice().updateUnsafe(true,p2);
        let _index137 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_865_v8).length));
        (_865_v8)[_index137] = _dafny.Set.fromElements((((_869_v10).contains(p2)) ? ((_869_v10).get(p2)) : (p2)));
        let _870_v11;
        _870_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
        let _871_v12;
        _871_v12 = true;
        let _872_v13;
        _872_v13 = _dafny.Seq.UnicodeFromString("wihxevenh");
        let _873_v14;
        _873_v14 = _module.D1.create_DC3(_872_v13, p2);
        let _874_v15;
        let _nw121 = new _module.C2();
        _nw121.__ctor();
        _874_v15 = _nw121;
        let _875_v16;
        _875_v16 = _dafny.Set.fromElements(_874_v15);
        let _rhs189 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_873_v14).dtor_cf4);
        let _rhs190 = false;
        let _rhs191 = ((!((_this).fm8(new BigNumber((_875_v16).length), false, p0, globalState)) || (_871_v12)) ? (new BigNumber((_869_v10).length)) : (p0));
        let _lhs135 = globalState;
        _870_v11 = _rhs189;
        _871_v12 = _rhs190;
        _lhs135.f4 = _rhs191;
        let _876_v17;
        _876_v17 = new _dafny.CodePoint('g'.codePointAt(0));
        let _877_v18;
        _877_v18 = _dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), _876_v17);
        let _878_v19;
        _878_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_877_v18).length));
        _878_v19 = _878_v19;
        _869_v10 = ((_869_v10).Merge(_869_v10)).Merge(_869_v10);
      }
      let _879_v20;
      _879_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      _879_v20 = _879_v20;
      let _880_v21;
      _880_v21 = _dafny.Set.fromElements(_module.__default.fm27(globalState));
      if ((_880_v21).IsSubsetOf(_880_v21)) {
        let _881_v22;
        let _nw122 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _881_v22 = _nw122;
        let _index138 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_881_v22).length));
        (_881_v22)[_index138] = _module.__default.fm38(p2, globalState);
        let _882_v23;
        _882_v23 = _dafny.Seq.UnicodeFromString("dnltn");
        let _883_v24;
        _883_v24 = _dafny.Seq.of(p2);
        let _index139 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length));
        (_853_v0)[_index139] = (_883_v24)[_module.__default.safeIndex(new BigNumber((p1).cardinality()), new BigNumber((_883_v24).length))];
        let _884_v25;
        _884_v25 = true;
        let _885_v26;
        _885_v26 = new _dafny.CodePoint('f'.codePointAt(0));
        let _886_v27;
        _886_v27 = _dafny.Seq.of(p0);
        let _index140 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_881_v22).length));
        let _index141 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length));
        let _rhs192 = _885_v26;
        let _rhs193 = _882_v23;
        let _rhs194 = !(!((_this).fm8((_886_v27)[_module.__default.safeIndex(p0, new BigNumber((_886_v27).length))], ((p2) ? (_884_v25) : (_884_v25)), p0, globalState)));
        let _rhs195 = _884_v25;
        let _lhs136 = _881_v22;
        let _lhs137 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_881_v22).length));
        let _lhs138 = _853_v0;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length));
        _lhs136[_lhs137] = _rhs192;
        _882_v23 = _rhs193;
        _lhs138[_lhs139] = _rhs194;
        _884_v25 = _rhs195;
        let _887_v28;
        let _nw123 = Array((new BigNumber(5)).toNumber()).fill(_module.D1.Default());
        _887_v28 = _nw123;
        let _888_v29;
        _888_v29 = _module.D1.create_DC2(p0);
        let _index142 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_887_v28).length));
        (_887_v28)[_index142] = _888_v29;
        let _index143 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_887_v28).length));
        (_887_v28)[_index143] = _888_v29;
        if (p2) {
          _882_v23 = _882_v23;
          let _889_v30;
          _889_v30 = _dafny.Set.fromElements((_module.D11.create_DC30(p2, p0)).dtor_cf47);
          let _890_v31;
          _890_v31 = _dafny.Seq.of(_dafny.Set.fromElements(p0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(494), p0)).cardinality()), p0), _889_v30, _889_v30, _889_v30);
          let _891_v32;
          _891_v32 = _dafny.Set.fromElements(_module.__default.fm26(globalState), !(p2));
          let _rhs196 = false;
          let _rhs197 = _module.__default.fm30(p0, p2, (_890_v31)[_module.__default.safeIndex(p0, new BigNumber((_890_v31).length))], !(!(!((_891_v32).IsProperSubsetOf(_891_v32)))), globalState);
          _884_v25 = _rhs196;
          _882_v23 = _rhs197;
          let _index144 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length));
          (_853_v0)[_index144] = !(p2);
          let _892_v33;
          _892_v33 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("rfny"));
          let _893_v34;
          _893_v34 = _module.D4.create_DC11(_892_v33);
          let _pat_let_tv11 = _892_v33;
          let _894_v35;
          _894_v35 = _dafny.Map.Empty.slice().updateUnsafe((_853_v0)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length))],function (_pat_let8_0) {
            return function (_895_dt__update__tmp_h0) {
              return function (_pat_let9_0) {
                return function (_896_dt__update_hcf15_h0) {
                  return _module.D4.create_DC11(_896_dt__update_hcf15_h0);
                }(_pat_let9_0);
              }(_pat_let_tv11);
            }(_pat_let8_0);
          }(_893_v34));
          let _897_v36;
          let _nw124 = new _module.C2();
          _nw124.__ctor();
          _897_v36 = _nw124;
          let _898_v37;
          _898_v37 = _dafny.Seq.of(_897_v36, _897_v36);
          (globalState).f13 = (_this).fm23(_894_v35, _module.__default.safeDivisionInt(new BigNumber(((_892_v33)[_module.__default.safeIndex(p0, new BigNumber((_892_v33).length))]).length), new BigNumber((_898_v37).length)), p0, globalState);
          _884_v25 = (p0).isLessThan((new BigNumber((p1).cardinality())).minus(p0));
        } else {
          let _899_v38;
          _899_v38 = _dafny.Seq.of(_882_v23);
          let _900_v39;
          _900_v39 = _dafny.Map.Empty.slice().updateUnsafe(p2,_module.D4.create_DC11(_899_v38));
          let _901_v40;
          _901_v40 = _dafny.Set.fromElements(p0, (_this).fm23(_900_v39, new BigNumber((_882_v23).length), new BigNumber(-193), globalState), new BigNumber((_dafny.Seq.of(p0)).length), p0, p0);
          _899_v38 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(950)), ((_902_v23) => function (_903_i1) {
            return _902_v23;
          })(_882_v23)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(950)), ((_904_v23) => function (_905_i1) {
            return _904_v23;
          })(_882_v23))).length)), _module.__default.fm30(new BigNumber((_dafny.Seq.of(new BigNumber(278))).length), !(p2), _901_v40, _884_v25, globalState)), _899_v38);
          (globalState).f4 = new BigNumber(310);
          let _906_v41;
          let _nw125 = new _module.C2();
          _nw125.__ctor();
          _906_v41 = _nw125;
          _882_v23 = _dafny.Seq.UnicodeFromString("bsu");
          let _907_v42;
          let _nw126 = Array((new BigNumber(5)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = (_853_v0)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_853_v0).length))];
          _nw126[(_dafny.ONE).toNumber()] = p2;
          _nw126[(new BigNumber(2)).toNumber()] = _884_v25;
          _nw126[(new BigNumber(3)).toNumber()] = p2;
          _nw126[(new BigNumber(4)).toNumber()] = !(!(p2) || (_884_v25));
          _907_v42 = _nw126;
          let _rhs198 = ((_884_v25) ? (_907_v42) : (_907_v42));
          let _rhs199 = new BigNumber((_883_v24).length);
          let _lhs140 = globalState;
          _907_v42 = _rhs198;
          _lhs140.f15 = _rhs199;
        }
        (globalState).f13 = p0;
        let _908_v43;
        let _nw127 = Array((new BigNumber(27)).toNumber()).fill(_module.D4.Default());
        _908_v43 = _nw127;
        let _909_v44;
        _909_v44 = _module.D4.create_DC12(p0);
        let _index145 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_908_v43).length));
        (_908_v43)[_index145] = _909_v44;
        let _index146 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_908_v43).length));
        (_908_v43)[_index146] = _909_v44;
      } else {
        let _910_v45;
        _910_v45 = new _dafny.CodePoint('y'.codePointAt(0));
        let _911_v46;
        _911_v46 = _dafny.Map.Empty.slice().updateUnsafe(_910_v45,false);
        let _912_v47;
        _912_v47 = _dafny.Seq.UnicodeFromString("qbe");
        _911_v46 = (_911_v46).update(_910_v45, (p0).isLessThanOrEqualTo(new BigNumber((_912_v47).length)));
        (globalState).f13 = (new BigNumber(-853)).multipliedBy(p0);
        let _913_v48;
        _913_v48 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(p0)).isLessThan((_dafny.ZERO).minus(p0)),p2);
        _913_v48 = (_913_v48).update(p2, p2);
        let _914_v49;
        _914_v49 = _dafny.Seq.of(((p2) ? (true) : (p2)));
        let _915_v50;
        _915_v50 = _dafny.Set.fromElements(p0);
        let _916_v51;
        let _nw128 = Array((new BigNumber(3)).toNumber());
        _nw128[(_dafny.ZERO).toNumber()] = new BigNumber(784);
        _nw128[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber((_915_v50).length)));
        _nw128[(new BigNumber(2)).toNumber()] = p0;
        _916_v51 = _nw128;
        let _index147 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_916_v51).length));
        (_916_v51)[_index147] = (new BigNumber(-129)).plus(p0);
        let _index148 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_916_v51).length));
        let _rhs200 = _914_v49;
        let _rhs201 = new BigNumber(886);
        let _lhs141 = _916_v51;
        let _lhs142 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_916_v51).length));
        _914_v49 = _rhs200;
        _lhs141[_lhs142] = _rhs201;
        let _917_v52;
        _917_v52 = false;
        _917_v52 = !(((_916_v51)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_916_v51).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of _dafny.IntegerRange(new BigNumber(745), new BigNumber(-795))) {
            let _918_v53 = _compr_37;
            if (((new BigNumber(745)).isLessThanOrEqualTo(_918_v53)) && ((_918_v53).isLessThan(new BigNumber(-795)))) {
              _coll37.push([_module.__default.safeModuloInt(_918_v53, (_916_v51)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_916_v51).length))]),_912_v47]);
            }
          }
          return _coll37;
        }()).length)))));
      }
      let _919_v54;
      _919_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-152));
      let _920_v55;
      _920_v55 = _module.D4.create_DC12(new BigNumber((_919_v54).length));
      let _921_v56;
      _921_v56 = _dafny.Seq.UnicodeFromString("jpfda");
      let _922_v57;
      _922_v57 = _dafny.Set.fromElements(new BigNumber((_921_v56).length));
      let _923_v58;
      _923_v58 = _dafny.MultiSet.fromElements(_922_v57);
      let _924_v59;
      _924_v59 = _dafny.Set.fromElements(p2, p2);
      let _925_v60;
      _925_v60 = _module.D7.create_DC22(true, p0, p2, p0);
      let _926_v61;
      _926_v61 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0));
      let _927_v62;
      let _nw129 = Array((new BigNumber(23)).toNumber());
      _nw129[(_dafny.ZERO).toNumber()] = p0;
      _nw129[(_dafny.ONE).toNumber()] = new BigNumber(805);
      _nw129[(new BigNumber(2)).toNumber()] = p0;
      _nw129[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_920_v55).dtor_cf16);
      _nw129[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).cardinality()),p2)).length);
      _nw129[(new BigNumber(5)).toNumber()] = p0;
      _nw129[(new BigNumber(6)).toNumber()] = p0;
      _nw129[(new BigNumber(7)).toNumber()] = new BigNumber((_923_v58).cardinality());
      _nw129[(new BigNumber(8)).toNumber()] = p0;
      _nw129[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), ((_928_p0, _929_v20, _930_p2) => function (_931_i3) {
        return _dafny.Seq.of((((_929_v20).contains(_928_p0)) ? ((_929_v20).get(_928_p0)) : (_930_p2)), (_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-685)), function (_932_i4) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}), !(_930_p2))).dtor_cf4, _930_p2);
      })(p0, _879_v20, p2)), _module.__default.safeIndex(new BigNumber((_924_v59).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(658)), ((_933_p0, _934_v20, _935_p2) => function (_936_i3) {
        return _dafny.Seq.of((((_934_v20).contains(_933_p0)) ? ((_934_v20).get(_933_p0)) : (_935_p2)), (_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-685)), function (_937_i4) {
  return new _dafny.CodePoint('t'.codePointAt(0));
}), !(_935_p2))).dtor_cf4, _935_p2);
      })(p0, _879_v20, p2))).length)), _dafny.Seq.of((_925_v60).dtor_cf37, p2))).length);
      _nw129[(new BigNumber(10)).toNumber()] = p0;
      _nw129[(new BigNumber(11)).toNumber()] = p0;
      _nw129[(new BigNumber(12)).toNumber()] = p0;
      _nw129[(new BigNumber(13)).toNumber()] = p0;
      _nw129[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality());
      _nw129[(new BigNumber(15)).toNumber()] = p0;
      _nw129[(new BigNumber(16)).toNumber()] = p0;
      _nw129[(new BigNumber(17)).toNumber()] = p0;
      _nw129[(new BigNumber(18)).toNumber()] = p0;
      _nw129[(new BigNumber(19)).toNumber()] = new BigNumber((_926_v61).cardinality());
      _nw129[(new BigNumber(20)).toNumber()] = p0;
      _nw129[(new BigNumber(21)).toNumber()] = p0;
      _nw129[(new BigNumber(22)).toNumber()] = p0;
      _927_v62 = _nw129;
      let _938_v63;
      _938_v63 = _dafny.Seq.of(_927_v62, _927_v62);
      let _939_i2;
      _939_i2 = _dafny.ZERO;
      L4: {
        while (!(_dafny.areEqual(_938_v63, _938_v63))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_939_i2)) {
              break L4;
            }
            _939_i2 = (_939_i2).plus(_dafny.ONE);
            let _index149 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
            (_927_v62)[_index149] = p0;
            let _940_v64;
            _940_v64 = _module.D14.create_DC35(_module.__default.fm45(globalState));
            let _941_v65;
            _941_v65 = _dafny.Seq.of(_940_v64, _940_v64);
            let _index150 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
            (_927_v62)[_index150] = new BigNumber((((false) ? (_941_v65) : (_941_v65))).length);
            if (p2) {
              (globalState).f4 = p0;
              let _942_v66;
              _942_v66 = true;
              _942_v66 = _942_v66;
              let _943_v67;
              _943_v67 = _dafny.Set.fromElements(_879_v20, _879_v20);
              let _944_v68;
              _944_v68 = _dafny.Seq.of(p2);
              let _rhs202 = !(p2);
              let _rhs203 = _942_v66;
              let _rhs204 = !(!(((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))]).isLessThanOrEqualTo(new BigNumber((_943_v67).length))) || (_dafny.Seq.contains(_944_v68, _942_v66)));
              _942_v66 = _rhs202;
              _942_v66 = _rhs203;
              _942_v66 = _rhs204;
              let _945_v69;
              _945_v69 = _dafny.Seq.of((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], new BigNumber(431), (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))]);
              (globalState).f15 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_945_v69, _module.__default.safeIndex(p0, new BigNumber((_945_v69).length)), new BigNumber((_919_v54).length)), _945_v69), _945_v69)).length);
              let _index151 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_853_v0).length));
              (_853_v0)[_index151] = !(p2);
              let _946_v70;
              let _nw130 = Array((new BigNumber(11)).toNumber()).fill(_module.D6.Default());
              _946_v70 = _nw130;
              let _947_v71;
              _947_v71 = _module.D6.create_DC16(_dafny.Seq.of(new BigNumber((_921_v56).length), p0));
              let _948_v72;
              _948_v72 = _module.D6.create_DC18(_947_v71);
              let _index152 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_946_v70).length));
              (_946_v70)[_index152] = _948_v72;
              let _index153 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_853_v0).length));
              let _index154 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_946_v70).length));
              let _rhs205 = (_944_v68)[_module.__default.safeIndex((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], new BigNumber((_944_v68).length))];
              let _rhs206 = _948_v72;
              let _rhs207 = p2;
              let _lhs143 = _853_v0;
              let _lhs144 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_853_v0).length));
              let _lhs145 = _946_v70;
              let _lhs146 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_946_v70).length));
              _lhs143[_lhs144] = _rhs205;
              _lhs145[_lhs146] = _rhs206;
              _942_v66 = _rhs207;
            } else {
              let _949_v73;
              _949_v73 = true;
              _949_v73 = p2;
              let _950_v74;
              let _init25 = ((_951_v62) => function (_952_i5) {
                return _module.__default.safeDivisionInt(_952_i5, (_951_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_951_v62).length))]);
              })(_927_v62);
              let _nw131 = Array((new BigNumber(23)).toNumber());
              for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw131.length); _i0_25++) {
                _nw131[_i0_25] = _init25(new BigNumber(_i0_25));
              }
              _950_v74 = _nw131;
              let _953_v75;
              _953_v75 = _module.D2.create_DC5(_950_v74);
              let _954_v76;
              _954_v76 = _module.D2.create_DC8(_953_v75);
              let _955_v77;
              let _nw132 = new _module.C3();
              _nw132.__ctor(_954_v76, (p0).isLessThan(new BigNumber((_924_v59).length)));
              _955_v77 = _nw132;
              let _956_v78;
              let _init26 = ((_957_v56) => function (_958_i6) {
                return _957_v56;
              })(_921_v56);
              let _nw133 = Array((new BigNumber(23)).toNumber());
              for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw133.length); _i0_26++) {
                _nw133[_i0_26] = _init26(new BigNumber(_i0_26));
              }
              _956_v78 = _nw133;
              let _index155 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_956_v78).length));
              (_956_v78)[_index155] = _921_v56;
              let _959_v79;
              _959_v79 = _dafny.MultiSet.fromElements((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], new BigNumber(729));
              let _960_v80;
              _960_v80 = _module.D8.create_DC24(p0, !(p2) || ((((_879_v20).contains(new BigNumber((_959_v79).cardinality()))) ? ((_879_v20).get(new BigNumber((_959_v79).cardinality()))) : (_955_v77.f25))));
              let _961_v81;
              _961_v81 = _dafny.Map.Empty.slice().updateUnsafe(((_955_v77.f25) ? (false) : (true)),_853_v0);
              let _index156 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_956_v78).length));
              let _rhs208 = _921_v56;
              let _rhs209 = _960_v80;
              let _rhs210 = (_961_v81).Merge((_961_v81).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,_853_v0)));
              let _lhs147 = _956_v78;
              let _lhs148 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_956_v78).length));
              _lhs147[_lhs148] = _rhs208;
              _960_v80 = _rhs209;
              _961_v81 = _rhs210;
              let _962_v82;
              _962_v82 = _dafny.Map.Empty.slice().updateUnsafe(p1,_950_v74);
              _962_v82 = (_962_v82).update(_module.__default.fm34(p0, p0, _949_v73, globalState), _950_v74);
              (globalState).f13 = (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))];
            }
            let _963_v83;
            _963_v83 = _dafny.Seq.of(_module.__default.fm30((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], p2, _922_v57, true, globalState));
            let _964_v84;
            _964_v84 = _module.D4.create_DC11(_963_v83);
            let _965_v85;
            _965_v85 = _dafny.Map.Empty.slice().updateUnsafe(true,_964_v84);
            let _966_v86;
            _966_v86 = _dafny.Seq.of((_this).fm23(_965_v85, (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], p0, globalState));
            _966_v86 = _dafny.Seq.update(_966_v86, _module.__default.safeIndex(new BigNumber((_module.__default.fm30(p0, p2, _922_v57, p2, globalState)).length), new BigNumber((_966_v86).length)), (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))]);
            let _967_v87;
            _967_v87 = _module.D11.create_DC29();
            let _source18 = _967_v87;
            if (_source18.is_DC29) {
              let _968_v88;
              _968_v88 = false;
              _968_v88 = true;
              (globalState).f7 = ((_dafny.ZERO).minus(p0)).plus((_dafny.ZERO).minus(p0));
              _968_v88 = true;
              let _969_v89;
              let _nw134 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _969_v89 = _nw134;
              let _nw135 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _969_v89 = _nw135;
            } else if (_source18.is_DC30) {
              let _970___mcc_h2 = (_source18).cf46;
              let _971___mcc_h3 = (_source18).cf47;
              let _972_cf47 = _971___mcc_h3;
              let _973_cf46 = _970___mcc_h2;
              let _974_v90;
              _974_v90 = new _dafny.CodePoint('b'.codePointAt(0));
              let _975_v91;
              let _nw136 = Array((new BigNumber(2)).toNumber());
              _nw136[(_dafny.ZERO).toNumber()] = (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))];
              _nw136[(_dafny.ONE).toNumber()] = new BigNumber((_module.__default.fm41(_974_v90, globalState)).length);
              _975_v91 = _nw136;
              let _976_v92;
              _976_v92 = _module.D2.create_DC7(_973_cf46, p1, _975_v91);
              let _977_v93;
              _977_v93 = _module.D2.create_DC8(_976_v92);
              let _pat_let_tv12 = _976_v92;
              let _978_v94;
              let _nw137 = new _module.C3();
              _nw137.__ctor(function (_pat_let10_0) {
                return function (_979_dt__update__tmp_h1) {
                  return function (_pat_let11_0) {
                    return function (_980_dt__update_hcf11_h0) {
                      return _module.D2.create_DC8(_980_dt__update_hcf11_h0);
                    }(_pat_let11_0);
                  }(_pat_let_tv12);
                }(_pat_let10_0);
              }(_977_v93), _973_cf46);
              _978_v94 = _nw137;
              _973_cf46 = p2;
              _974_v90 = _974_v90;
              let _981_v95;
              let _nw138 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
              _981_v95 = _nw138;
              let _982_v96;
              _982_v96 = _module.D7.create_DC21((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], new BigNumber(508), _973_cf46, (_this).fm8(p0, _978_v94.f25, (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], globalState), _978_v94.f25);
              let _index157 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_981_v95).length));
              (_981_v95)[_index157] = (((_982_v96).dtor_cf32) ? (_dafny.Seq.update(_966_v86, _module.__default.safeIndex((_this).fm23(_965_v85, (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], _972_cf47, globalState), new BigNumber((_966_v86).length)), _972_cf47)) : (_966_v86));
              let _983_v97;
              let _nw139 = Array((new BigNumber(17)).toNumber()).fill([]);
              _983_v97 = _nw139;
              let _index158 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_983_v97).length));
              (_983_v97)[_index158] = _853_v0;
              let _984_v98;
              let _nw140 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _984_v98 = _nw140;
              let _index159 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_984_v98).length));
              (_984_v98)[_index159] = _921_v56;
              let _985_v99;
              _985_v99 = _dafny.Map.Empty.slice().updateUnsafe(_853_v0,_921_v56);
              let _986_v100;
              _986_v100 = _dafny.Seq.of(_966_v86, _966_v86);
              let _987_v101;
              _987_v101 = _dafny.Seq.of(_853_v0);
              let _index160 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_981_v95).length));
              let _index161 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_983_v97).length));
              let _index162 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_984_v98).length));
              let _rhs211 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(((_978_v94.f25) ? (_966_v86) : (_966_v86)), _module.__default.safeIndex(_972_cf47, new BigNumber((((_978_v94.f25) ? (_966_v86) : (_966_v86))).length)), _972_cf47), _module.__default.safeIndex(new BigNumber((p1).cardinality()), new BigNumber((_dafny.Seq.update(((_978_v94.f25) ? (_966_v86) : (_966_v86)), _module.__default.safeIndex(_972_cf47, new BigNumber((((_978_v94.f25) ? (_966_v86) : (_966_v86))).length)), _972_cf47)).length)), new BigNumber((_986_v100).length)), _966_v86);
              let _rhs212 = ((true) ? (((p2) ? (_853_v0) : (_853_v0))) : (((_973_cf46) ? (_853_v0) : ((_987_v101)[_module.__default.safeIndex((_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))], new BigNumber((_987_v101).length))]))));
              let _rhs213 = !(_973_cf46) || (_973_cf46);
              let _rhs214 = ((_973_cf46) ? (_dafny.Seq.update(_921_v56, _module.__default.safeIndex(_972_cf47, new BigNumber((_921_v56).length)), _module.__default.fm38(true, globalState))) : (_921_v56));
              let _rhs215 = (_dafny.Map.Empty.slice().updateUnsafe(_853_v0,_921_v56)).Merge(_985_v99);
              let _lhs149 = _981_v95;
              let _lhs150 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_981_v95).length));
              let _lhs151 = _983_v97;
              let _lhs152 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_983_v97).length));
              let _lhs153 = _984_v98;
              let _lhs154 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_984_v98).length));
              _lhs149[_lhs150] = _rhs211;
              _lhs151[_lhs152] = _rhs212;
              _973_cf46 = _rhs213;
              _lhs153[_lhs154] = _rhs214;
              _985_v99 = _rhs215;
            } else if (_source18.is_DC28) {
              let _988___mcc_h4 = (_source18).cf45;
              let _989_cf45 = _988___mcc_h4;
              let _990_v102;
              let _nw141 = Array((new BigNumber(20)).toNumber()).fill(_module.D11.Default());
              _990_v102 = _nw141;
              let _index163 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_990_v102).length));
              (_990_v102)[_index163] = _967_v87;
              let _index164 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_990_v102).length));
              (_990_v102)[_index164] = _967_v87;
              let _991_v103;
              _991_v103 = false;
              let _index165 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
              let _rhs216 = (_924_v59).Difference((_924_v59).Union(_924_v59));
              let _rhs217 = _991_v103;
              let _rhs218 = (p0).plus(((p2) ? (p0) : (p0)));
              let _rhs219 = new BigNumber(153);
              let _lhs155 = globalState;
              let _lhs156 = _927_v62;
              let _lhs157 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
              _924_v59 = _rhs216;
              _991_v103 = _rhs217;
              _lhs155.f15 = _rhs218;
              _lhs156[_lhs157] = _rhs219;
              _991_v103 = p2;
              let _992_v104;
              let _nw142 = new _module.C1();
              _nw142.__ctor();
              _992_v104 = _nw142;
            } else {
              let _993___mcc_h5 = (_source18).cf48;
              let _994_cf48 = _993___mcc_h5;
              _919_v54 = _919_v54;
              (globalState).f4 = p0;
              let _995_v105;
              _995_v105 = false;
              let _996_v106;
              let _init27 = function (_997_i7) {
                return (_997_i7).multipliedBy(new BigNumber(264));
              };
              let _nw143 = Array((new BigNumber(28)).toNumber());
              for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw143.length); _i0_27++) {
                _nw143[_i0_27] = _init27(new BigNumber(_i0_27));
              }
              _996_v106 = _nw143;
              let _998_v107;
              _998_v107 = _module.D2.create_DC5(_996_v106);
              let _index166 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
              let _rhs220 = _module.__default.fm26(globalState);
              let _rhs221 = p0;
              let _rhs222 = (_927_v62)[_module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length))];
              let _rhs223 = p3;
              let _rhs224 = p0;
              let _lhs158 = globalState;
              let _lhs159 = globalState;
              let _lhs160 = _927_v62;
              let _lhs161 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_927_v62).length));
              _995_v105 = _rhs220;
              _lhs158.f13 = _rhs221;
              _lhs159.f13 = _rhs222;
              _998_v107 = _rhs223;
              _lhs160[_lhs161] = _rhs224;
              let _999_v108;
              _999_v108 = new _dafny.CodePoint('p'.codePointAt(0));
              let _1000_v109;
              _1000_v109 = _dafny.Map.Empty.slice().updateUnsafe(_999_v108,_853_v0);
              _1000_v109 = (_1000_v109).update(_999_v108, _853_v0);
            }
          }
        }
      }
      let _1001_v110;
      _1001_v110 = false;
      _1001_v110 = _1001_v110;
      let _hi6 = p0;
      for (let _1002_i8 = ((false) ? (p0) : (p0)); _1002_i8.isLessThan(_hi6); _1002_i8 = _1002_i8.plus(_dafny.ONE)) {
        let _1003_v111;
        let _nw144 = Array((new BigNumber(19)).toNumber()).fill(_module.D1.Default());
        _1003_v111 = _nw144;
        let _index167 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        (_853_v0)[_index167] = p2;
        let _1004_v112;
        _1004_v112 = _module.D2.create_DC7(_1001_v110, p1, _927_v62);
        let _1005_v113;
        _1005_v113 = _dafny.Set.fromElements(_921_v56, _dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), function (_1006_i9) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }), _921_v56, _921_v56, _921_v56);
        let _index168 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        let _rhs225 = ((_dafny.areEqual(_921_v56, _921_v56)) ? (p2) : ((_1004_v112).dtor_cf8));
        let _rhs226 = _1003_v111;
        let _rhs227 = (_1005_v113).IsSubsetOf(_1005_v113);
        let _rhs228 = _1001_v110;
        let _lhs162 = _853_v0;
        let _lhs163 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        _1001_v110 = _rhs225;
        _1003_v111 = _rhs226;
        _lhs162[_lhs163] = _rhs227;
        _1001_v110 = _rhs228;
        let _1007_v114;
        _1007_v114 = _dafny.Seq.of(_921_v56, _921_v56);
        let _1008_v115;
        _1008_v115 = _module.D4.create_DC11(_1007_v114);
        let _1009_v116;
        _1009_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1001_v110,_1008_v115);
        let _1010_v117;
        _1010_v117 = _dafny.MultiSet.fromElements((_this).fm23(_1009_v116, _1002_i8, new BigNumber(330), globalState), (_dafny.ZERO).minus(p0));
        _1001_v110 = (_1010_v117).IsSubsetOf(_1010_v117);
        let _1011_v118;
        _1011_v118 = _dafny.Seq.of(_1001_v110, _module.__default.fm26(globalState), _1001_v110);
        _1011_v118 = _1011_v118;
        let _index169 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        let _index170 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        let _rhs229 = _1001_v110;
        let _rhs230 = (_1011_v118)[_module.__default.safeIndex(p0, new BigNumber((_1011_v118).length))];
        let _lhs164 = _853_v0;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        let _lhs166 = _853_v0;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_853_v0).length));
        _lhs164[_lhs165] = _rhs229;
        _lhs166[_lhs167] = _rhs230;
      }
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(p0, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(180)), _dafny.Set.fromElements(new BigNumber(-596))))).Intersect(_dafny.MultiSet.fromElements(function () {
        let _coll38 = new _dafny.Set();
        for (const _compr_38 of (function () {
          let _coll39 = new _dafny.Set();
          for (const _compr_39 of (function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(741), new BigNumber(602))) {
              let _1012_v0 = _compr_40;
              if (((new BigNumber(741)).isLessThanOrEqualTo(_1012_v0)) && ((_1012_v0).isLessThan(new BigNumber(602)))) {
                _coll40.push([(_1012_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-193))).length)),new BigNumber(587)]);
              }
            }
            return _coll40;
          }()).Keys.Elements) {
            let _1013_v1 = _compr_39;
            if ((function () {
              let _coll41 = new _dafny.Map();
              for (const _compr_41 of _dafny.IntegerRange(new BigNumber(741), new BigNumber(602))) {
                let _1012_v0 = _compr_41;
                if (((new BigNumber(741)).isLessThanOrEqualTo(_1012_v0)) && ((_1012_v0).isLessThan(new BigNumber(602)))) {
                  _coll41.push([(_1012_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-193))).length)),new BigNumber(587)]);
                }
              }
              return _coll41;
            }()).contains(_1013_v1)) {
              _coll39.add(_module.__default.safeDivisionInt(_1013_v1, new BigNumber((_dafny.Seq.UnicodeFromString("we")).length)));
            }
          }
          return _coll39;
        }()).Elements) {
          let _1014_v2 = _compr_38;
          if ((function () {
            let _coll42 = new _dafny.Set();
            for (const _compr_42 of (function () {
              let _coll43 = new _dafny.Map();
              for (const _compr_43 of _dafny.IntegerRange(new BigNumber(741), new BigNumber(602))) {
                let _1012_v0 = _compr_43;
                if (((new BigNumber(741)).isLessThanOrEqualTo(_1012_v0)) && ((_1012_v0).isLessThan(new BigNumber(602)))) {
                  _coll43.push([(_1012_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-193))).length)),new BigNumber(587)]);
                }
              }
              return _coll43;
            }()).Keys.Elements) {
              let _1015_v1 = _compr_42;
              if ((function () {
                let _coll44 = new _dafny.Map();
                for (const _compr_44 of _dafny.IntegerRange(new BigNumber(741), new BigNumber(602))) {
                  let _1012_v0 = _compr_44;
                  if (((new BigNumber(741)).isLessThanOrEqualTo(_1012_v0)) && ((_1012_v0).isLessThan(new BigNumber(602)))) {
                    _coll44.push([(_1012_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-193))).length)),new BigNumber(587)]);
                  }
                }
                return _coll44;
              }()).contains(_1015_v1)) {
                _coll42.add(_module.__default.safeDivisionInt(_1015_v1, new BigNumber((_dafny.Seq.UnicodeFromString("we")).length)));
              }
            }
            return _coll42;
          }()).contains(_1014_v2)) {
            _coll38.add(_module.__default.safeDivisionInt(_1014_v2, new BigNumber(769)));
          }
        }
        return _coll38;
      }(), function () {
        let _coll45 = new _dafny.Set();
        for (const _compr_45 of _dafny.IntegerRange(new BigNumber(374), new BigNumber(894))) {
          let _1016_v3 = _compr_45;
          if (((new BigNumber(374)).isLessThanOrEqualTo(_1016_v3)) && ((_1016_v3).isLessThan(new BigNumber(894)))) {
            _coll45.add((_1016_v3).multipliedBy(new BigNumber(245)));
          }
        }
        return _coll45;
      }(), function () {
        let _coll46 = new _dafny.Set();
        for (const _compr_46 of _dafny.IntegerRange(new BigNumber(666), new BigNumber(713))) {
          let _1017_v4 = _compr_46;
          if (((new BigNumber(666)).isLessThanOrEqualTo(_1017_v4)) && ((_1017_v4).isLessThan(new BigNumber(713)))) {
            _coll46.add((_1017_v4).multipliedBy(new BigNumber(144)));
          }
        }
        return _coll46;
      }()))).Difference((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(412), new BigNumber(183), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(337), new BigNumber(-99))) {
          let _1018_v5 = _compr_47;
          if (((new BigNumber(337)).isLessThanOrEqualTo(_1018_v5)) && ((_1018_v5).isLessThan(new BigNumber(-99)))) {
            _coll47.push([(_1018_v5).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(500))).length)),new BigNumber(853)]);
          }
        }
        return _coll47;
      }()).length),new BigNumber(-161))).length)), function () {
        let _coll48 = new _dafny.Set();
        for (const _compr_48 of (_dafny.MultiSet.fromElements(new BigNumber(215), new BigNumber((_dafny.Seq.UnicodeFromString("h")).length))).Elements) {
          let _1019_v6 = _compr_48;
          if ((_dafny.MultiSet.fromElements(new BigNumber(215), new BigNumber((_dafny.Seq.UnicodeFromString("h")).length))).contains(_1019_v6)) {
            _coll48.add(_module.__default.safeDivisionInt(_1019_v6, new BigNumber(772)));
          }
        }
        return _coll48;
      }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("jghlcq")).length), new BigNumber((_dafny.Set.fromElements(false)).length)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0))))).cardinality())), _dafny.Set.fromElements(new BigNumber(138))))));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))).length)).isLessThan((new BigNumber(99)).multipliedBy(new BigNumber(952)));
    };
    fm20(globalState) {
      let _this = this;
      return !(((_dafny.Seq.IsPrefixOf(_dafny.Seq.of(true, true), _dafny.Seq.of(false))) ? ((!(!(false))) && (false)) : ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-758)))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(785), new BigNumber(488), new BigNumber(825))))));
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      (globalState).f15 = p0;
      if (p2) {
        let _1020_v0;
        _1020_v0 = false;
        let _1021_v1;
        _1021_v1 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1022_v2;
        _1022_v2 = _dafny.Set.fromElements(_1020_v0);
        let _1023_v3;
        let _init28 = ((_1024_p0) => function (_1025_i0) {
          return _module.__default.safeDivisionInt(_1025_i0, _1024_p0);
        })(p0);
        let _nw145 = Array((new BigNumber(28)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw145.length); _i0_28++) {
          _nw145[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1023_v3 = _nw145;
        let _1026_v4;
        let _nw146 = Array((new BigNumber(23)).toNumber());
        _nw146[(_dafny.ZERO).toNumber()] = true;
        _nw146[(_dafny.ONE).toNumber()] = (true) === (_1020_v0);
        _nw146[(new BigNumber(2)).toNumber()] = _1020_v0;
        _nw146[(new BigNumber(3)).toNumber()] = (_module.__default.fm21(new _dafny.CodePoint('m'.codePointAt(0)), _1021_v1, globalState)).dtor_cf4;
        _nw146[(new BigNumber(4)).toNumber()] = p2;
        _nw146[(new BigNumber(5)).toNumber()] = p2;
        _nw146[(new BigNumber(6)).toNumber()] = ((true) ? (_1020_v0) : (true));
        _nw146[(new BigNumber(7)).toNumber()] = _1020_v0;
        _nw146[(new BigNumber(8)).toNumber()] = (_this).fm20(globalState);
        _nw146[(new BigNumber(9)).toNumber()] = (p1).IsDisjointFrom(p1);
        _nw146[(new BigNumber(10)).toNumber()] = p2;
        _nw146[(new BigNumber(11)).toNumber()] = p2;
        _nw146[(new BigNumber(12)).toNumber()] = p2;
        _nw146[(new BigNumber(13)).toNumber()] = (_1022_v2).IsSubsetOf(_1022_v2);
        _nw146[(new BigNumber(14)).toNumber()] = p2;
        _nw146[(new BigNumber(15)).toNumber()] = _1020_v0;
        _nw146[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.fromElements(_1023_v3, _1023_v3)).IsProperSubsetOf((_dafny.MultiSet.fromElements(_1023_v3)).update(_1023_v3, _module.__default.abs((_dafny.ZERO).minus(p0))));
        _nw146[(new BigNumber(17)).toNumber()] = p2;
        _nw146[(new BigNumber(18)).toNumber()] = (_this).fm20(globalState);
        _nw146[(new BigNumber(19)).toNumber()] = false;
        _nw146[(new BigNumber(20)).toNumber()] = !(p2);
        _nw146[(new BigNumber(21)).toNumber()] = p2;
        _nw146[(new BigNumber(22)).toNumber()] = (_this).fm8(p0, p2, p0, globalState);
        _1026_v4 = _nw146;
        let _1027_v5;
        _1027_v5 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("g"));
        let _index171 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length));
        (_1026_v4)[_index171] = (_1027_v5).IsSubsetOf(_1027_v5);
        let _1028_v6;
        _1028_v6 = _dafny.Seq.UnicodeFromString("slclxhejc");
        let _1029_v7;
        _1029_v7 = _dafny.Seq.of(_1028_v6, _dafny.Seq.update(_1028_v6, _module.__default.safeIndex(p0, new BigNumber((_1028_v6).length)), new _dafny.CodePoint('x'.codePointAt(0))));
        let _1030_v8;
        _1030_v8 = _module.D4.create_DC11(_1029_v7);
        let _index172 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length));
        let _rhs231 = p2;
        let _rhs232 = p2;
        let _rhs233 = (_1030_v8).dtor_cf15;
        let _lhs168 = _1026_v4;
        let _lhs169 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length));
        _1020_v0 = _rhs231;
        _lhs168[_lhs169] = _rhs232;
        _1029_v7 = _rhs233;
        let _index173 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length));
        (_1023_v3)[_index173] = p0;
        let _index174 = _module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length));
        (_1023_v3)[_index174] = (((_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))]) ? (p0) : (p0));
        let _index175 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length));
        let _rhs234 = ((_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))]).isLessThan(_module.__default.safeDivisionInt((_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))], (_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))]));
        let _rhs235 = _1021_v1;
        let _lhs170 = _1026_v4;
        let _lhs171 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length));
        _lhs170[_lhs171] = _rhs234;
        _1021_v1 = _rhs235;
        let _1031_v9;
        _1031_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1020_v0);
        let _1032_v10;
        _1032_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))],_1023_v3);
        let _1033_v11;
        _1033_v11 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_1032_v10).length)),_1031_v9);
        let _1034_v12;
        let _nw147 = Array((new BigNumber(17)).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = _1031_v9;
        _nw147[(_dafny.ONE).toNumber()] = (_1031_v9).Merge(_1031_v9);
        _nw147[(new BigNumber(2)).toNumber()] = (_1031_v9).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
        _nw147[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))],_1020_v0);
        _nw147[(new BigNumber(4)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(5)).toNumber()] = (_1031_v9).Merge(_1031_v9);
        _nw147[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))],p2);
        _nw147[(new BigNumber(7)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(8)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p2,(_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))])).update(p2, (_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))]);
        _nw147[(new BigNumber(9)).toNumber()] = (_1031_v9).Merge(_1031_v9);
        _nw147[(new BigNumber(10)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(11)).toNumber()] = (_1031_v9).Merge(_1031_v9);
        _nw147[(new BigNumber(12)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(13)).toNumber()] = ((_1020_v0) ? (_module.__default.fm22(_1029_v7, (_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))], globalState)) : ((((_1033_v11).contains(p0)) ? ((_1033_v11).get(p0)) : (_1031_v9))));
        _nw147[(new BigNumber(14)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(15)).toNumber()] = _1031_v9;
        _nw147[(new BigNumber(16)).toNumber()] = _1031_v9;
        _1034_v12 = _nw147;
        let _index176 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1034_v12).length));
        (_1034_v12)[_index176] = ((false) ? (_dafny.Map.Empty.slice().updateUnsafe(_1020_v0,(_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))])) : (_module.__default.fm22(_1029_v7, p0, globalState)));
        let _index177 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1034_v12).length));
        let _rhs236 = !(p2);
        let _rhs237 = (_1031_v9).Merge(_1031_v9);
        let _rhs238 = (_1026_v4)[_module.__default.safeIndex(new BigNumber(670), new BigNumber((_1026_v4).length))];
        let _lhs172 = _1034_v12;
        let _lhs173 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_1034_v12).length));
        _1020_v0 = _rhs236;
        _lhs172[_lhs173] = _rhs237;
        _1020_v0 = _rhs238;
        let _1035_v13;
        _1035_v13 = _module.D2.create_DC6(((_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))]).plus((_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))]));
        let _pat_let_tv13 = p0;
        _1035_v13 = function (_pat_let12_0) {
          return function (_1036_dt__update__tmp_h0) {
            return function (_pat_let13_0) {
              return function (_1037_dt__update_hcf7_h0) {
                return _module.D2.create_DC6(_1037_dt__update_hcf7_h0);
              }(_pat_let13_0);
            }(_pat_let_tv13);
          }(_pat_let12_0);
        }(_module.D2.create_DC6((_1023_v3)[_module.__default.safeIndex(new BigNumber(876), new BigNumber((_1023_v3).length))]));
      } else {
        let _1038_v14;
        let _nw148 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1038_v14 = _nw148;
        let _1039_v15;
        _1039_v15 = _module.D2.create_DC8(_module.D2.create_DC5(_1038_v14));
        let _1040_v16;
        _1040_v16 = _module.D2.create_DC8(_1039_v15);
        let _pat_let_tv14 = _1039_v15;
        let _1041_v17;
        let _nw149 = new _module.C3();
        _nw149.__ctor(function (_pat_let14_0) {
          return function (_1042_dt__update__tmp_h1) {
            return function (_pat_let15_0) {
              return function (_1043_dt__update_hcf11_h0) {
                return _module.D2.create_DC8(_1043_dt__update_hcf11_h0);
              }(_pat_let15_0);
            }(_pat_let_tv14);
          }(_pat_let14_0);
        }(_1040_v16), p2);
        _1041_v17 = _nw149;
        let _1044_v18;
        _1044_v18 = _dafny.Seq.UnicodeFromString("g");
        let _1045_v19;
        _1045_v19 = _dafny.Set.fromElements(p0);
        let _1046_v20;
        _1046_v20 = _dafny.MultiSet.fromElements(new BigNumber((_1045_v19).length), p0);
        let _1047_v21;
        _1047_v21 = _dafny.Seq.of(p0);
        let _1048_v22;
        let _nw150 = Array((new BigNumber(8)).toNumber());
        _nw150[(_dafny.ZERO).toNumber()] = _1046_v20;
        _nw150[(_dafny.ONE).toNumber()] = _1046_v20;
        _nw150[(new BigNumber(2)).toNumber()] = (_1046_v20).Difference(_dafny.MultiSet.FromArray(_1047_v21));
        _nw150[(new BigNumber(3)).toNumber()] = _1046_v20;
        _nw150[(new BigNumber(4)).toNumber()] = _1046_v20;
        _nw150[(new BigNumber(5)).toNumber()] = (_1046_v20).Union(_1046_v20);
        _nw150[(new BigNumber(6)).toNumber()] = (_1046_v20).Intersect((_1046_v20).update(p0, _module.__default.abs(p0)));
        _nw150[(new BigNumber(7)).toNumber()] = _1046_v20;
        _1048_v22 = _nw150;
        let _1049_v23;
        _1049_v23 = _dafny.Seq.of(_1044_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), function (_1050_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }));
        let _1051_v24;
        _1051_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1041_v17.f25);
        let _1052_v25;
        _1052_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1051_v24).length),_1041_v17.f25);
        let _1053_v26;
        _1053_v26 = _module.D11.create_DC30(_1041_v17.f25, new BigNumber((_1051_v24).length));
        let _1054_v27;
        _1054_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1053_v26).dtor_cf46,p0);
        let _1055_v28;
        _1055_v28 = _dafny.Seq.of(_1041_v17.f25, p2);
        let _1056_v29;
        _1056_v29 = new _dafny.CodePoint('o'.codePointAt(0));
        let _rhs239 = _dafny.Seq.Concat(_1044_v18, (_1049_v23)[_module.__default.safeIndex(p0, new BigNumber((_1049_v23).length))]);
        let _rhs240 = _1048_v22;
        let _rhs241 = new BigNumber((_module.__default.fm39((new BigNumber(((_1054_v27).update(_module.__default.fm26(globalState), new BigNumber(-592))).length)).minus(new BigNumber((_1055_v28).length)), _1056_v29, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), ((_1057_p0) => function (_1058_i2) {
          return _1057_p0;
        })(p0)), _1047_v21), p2, globalState)).length);
        let _rhs242 = p0;
        let _lhs174 = globalState;
        let _lhs175 = globalState;
        _1044_v18 = _rhs239;
        _1048_v22 = _rhs240;
        _lhs174.f4 = _rhs241;
        _lhs175.f13 = _rhs242;
        let _1059_v30;
        _1059_v30 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _1060_v31;
        _1060_v31 = _dafny.MultiSet.fromElements(_1059_v30);
        _1038_v14 = (((_1060_v31).IsProperSubsetOf(_1060_v31)) ? (_1038_v14) : (_1038_v14));
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), ((_1061_v29) => function (_1062_i3) {
          return _1061_v29;
        })(_1056_v29)), _1044_v18), _dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_1063_i4) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        }))) {
          let _index178 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1038_v14).length));
          (_1038_v14)[_index178] = ((_1041_v17).fm2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1041_v17.f25,_dafny.Seq.UnicodeFromString("xs"))).length), globalState)).plus(new BigNumber((_1055_v28).length));
          let _index179 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_1038_v14).length));
          (_1038_v14)[_index179] = p0;
          (_1041_v17).f25 = (_1041_v17.f25) === (_dafny.areEqual(_1055_v28, _dafny.Seq.of(_1041_v17.f25)));
          (_1041_v17).f25 = ((_1038_v14)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1038_v14).length))]).isLessThan((_1038_v14)[_module.__default.safeIndex(new BigNumber(385), new BigNumber((_1038_v14).length))]);
          (_1041_v17).f25 = true;
          _1059_v30 = (_1059_v30).update(false, (_1041_v17.f25) || (p2));
        } else {
          _1056_v29 = _module.__default.fm38((_1041_v17.f25) === (p2), globalState);
          _1056_v29 = _1056_v29;
          let _1064_v32;
          _1064_v32 = _dafny.Set.fromElements((_1055_v28)[_module.__default.safeIndex(p0, new BigNumber((_1055_v28).length))]);
          (globalState).f4 = (_dafny.ZERO).minus((new BigNumber((_1059_v30).length)).plus((new BigNumber((_1044_v18).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_1064_v32).length))).cardinality()))));
          (globalState).f13 = p0;
          (_1041_v17).f25 = ((p2) ? (p2) : (_1041_v17.f25));
        }
        (_1041_v17).f25 = p2;
      }
      let _1065_v33;
      _1065_v33 = _dafny.Seq.UnicodeFromString("phf");
      let _1066_v34;
      _1066_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1065_v33,p2);
      let _1067_i5;
      _1067_i5 = _dafny.ZERO;
      L5: {
        while (!((((_1066_v34).contains(_1065_v33)) ? ((_1066_v34).get(_1065_v33)) : (p2)))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1067_i5)) {
              break L5;
            }
            _1067_i5 = (_1067_i5).plus(_dafny.ONE);
            let _1068_v35;
            _1068_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(175),p0);
            _1068_v35 = (_1068_v35).update(_module.__default.safeModuloInt(p0, p0), p0);
            _1068_v35 = _1068_v35;
            let _1069_v36;
            _1069_v36 = new _dafny.CodePoint('r'.codePointAt(0));
            _1069_v36 = _1069_v36;
            let _1070_v37;
            _1070_v37 = _dafny.Set.fromElements(p2);
            let _1071_v38;
            let _nw151 = Array((new BigNumber(17)).toNumber());
            _nw151[(_dafny.ZERO).toNumber()] = _1070_v37;
            _nw151[(_dafny.ONE).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(2)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(3)).toNumber()] = _module.__default.fm37(globalState);
            _nw151[(new BigNumber(4)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(5)).toNumber()] = (_1070_v37).Union(_module.__default.fm37(globalState));
            _nw151[(new BigNumber(6)).toNumber()] = ((p2) ? (_1070_v37) : (_1070_v37));
            _nw151[(new BigNumber(7)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(8)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(p2);
            _nw151[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(p2);
            _nw151[(new BigNumber(11)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(12)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements(!(p2));
            _nw151[(new BigNumber(14)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(15)).toNumber()] = _1070_v37;
            _nw151[(new BigNumber(16)).toNumber()] = _dafny.Set.fromElements(!(p2), p2, false, p2);
            _1071_v38 = _nw151;
            let _index180 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_1071_v38).length));
            (_1071_v38)[_index180] = _1070_v37;
            let _index181 = _module.__default.safeIndex(new BigNumber(208), new BigNumber((_1071_v38).length));
            (_1071_v38)[_index181] = _1070_v37;
          }
        }
      }
      let _1072_v39;
      let _nw152 = Array((new BigNumber(12)).toNumber()).fill(false);
      _1072_v39 = _nw152;
      let _index182 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1072_v39).length));
      (_1072_v39)[_index182] = true;
      let _1073_v40;
      _1073_v40 = false;
      let _1074_v41;
      _1074_v41 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _1075_v42;
      _1075_v42 = _1074_v41;
      let _index183 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1072_v39).length));
      let _rhs243 = (p0).plus(p0);
      let _rhs244 = _1073_v40;
      let _rhs245 = !(p2);
      let _rhs246 = _1075_v42;
      let _lhs176 = globalState;
      let _lhs177 = _1072_v39;
      let _lhs178 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1072_v39).length));
      _lhs176.f13 = _rhs243;
      _lhs177[_lhs178] = _rhs244;
      _1073_v40 = _rhs245;
      _1075_v42 = _rhs246;
      let _1076_v43;
      let _nw153 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _1076_v43 = _nw153;
      _1076_v43 = _1076_v43;
      let _index184 = _module.__default.safeIndex(new BigNumber(817), new BigNumber((_1072_v39).length));
      (_1072_v39)[_index184] = p2;
      return;
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm15(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(353)), function (_1077_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("vmtigl"))).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ggpkkfhx"), _dafny.Seq.UnicodeFromString("mo"), _dafny.Seq.UnicodeFromString("jhtp"), _dafny.Seq.UnicodeFromString("rhv")))).Union(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-201)), function (_1078_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })));
    };
    fm16(p0, globalState) {
      let _this = this;
      return false;
    };
    m9(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.MultiSet.Empty;
      let r2 = false;
      let _1079_v0;
      _1079_v0 = _dafny.MultiSet.fromElements(p1, (_this).fm16(p0, globalState));
      let _1080_v1;
      let _nw154 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _1080_v1 = _nw154;
      let _1081_v2;
      _1081_v2 = _module.D2.create_DC7(true, _1079_v0, _1080_v1);
      let _1082_v3;
      _1082_v3 = _dafny.MultiSet.fromElements(_1081_v2, _1081_v2, _1081_v2, _1081_v2, _1081_v2);
      let _1083_v4;
      _1083_v4 = _dafny.Seq.of(_1082_v3);
      let _1084_i0;
      _1084_i0 = _dafny.ZERO;
      L6: {
        while (((_1083_v4)[_module.__default.safeIndex(p0, new BigNumber((_1083_v4).length))]).IsSubsetOf(_1082_v3)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1084_i0)) {
              break L6;
            }
            _1084_i0 = (_1084_i0).plus(_dafny.ONE);
            let _1085_v5;
            _1085_v5 = new _dafny.CodePoint('b'.codePointAt(0));
            _1085_v5 = new _dafny.CodePoint('x'.codePointAt(0));
            (globalState).f4 = p0;
            r2 = !(p0).isEqualTo(p0);
            let _1086_v6;
            _1086_v6 = _dafny.Seq.of(_module.__default.fm17(p0, globalState), p0);
            let _1087_v7;
            _1087_v7 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
            let _1088_v8;
            _1088_v8 = _dafny.Set.fromElements((_1086_v6)[_module.__default.safeIndex(new BigNumber((_1087_v7).length), new BigNumber((_1086_v6).length))]);
            r2 = (_1088_v8).IsDisjointFrom(_1088_v8);
          }
        }
      }
      let _1089_v9;
      let _nw155 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
      _1089_v9 = _nw155;
      let _1090_v10;
      _1090_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(p1));
      let _index185 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length));
      (_1089_v9)[_index185] = (_1090_v10).Merge(function () {
        let _coll49 = new _dafny.Map();
        for (const _compr_49 of _dafny.IntegerRange(new BigNumber(693), new BigNumber(711))) {
          let _1091_v11 = _compr_49;
          if (((new BigNumber(693)).isLessThanOrEqualTo(_1091_v11)) && ((_1091_v11).isLessThan(new BigNumber(711)))) {
            _coll49.push([_module.__default.safeModuloInt(_1091_v11, (_dafny.ZERO).minus(p0)),p1]);
          }
        }
        return _coll49;
      }());
      let _1092_v12;
      _1092_v12 = new _dafny.CodePoint('h'.codePointAt(0));
      let _1093_v13;
      _1093_v13 = _dafny.MultiSet.fromElements(_1092_v12, _1092_v12);
      let _1094_v14;
      _1094_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1093_v13);
      let _1095_v15;
      _1095_v15 = _dafny.Seq.UnicodeFromString("yvqumpwp");
      let _index186 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length));
      (_1089_v9)[_index186] = (_module.__default.fm18(p0, new BigNumber(643), (((_1094_v14).contains(new BigNumber((_1095_v15).length))) ? ((_1094_v14).get(new BigNumber((_1095_v15).length))) : (_1093_v13)), globalState)).update(p0, (p2) === (p1));
      let _1096_v16;
      let _nw156 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1096_v16 = _nw156;
      let _index187 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length));
      (_1096_v16)[_index187] = _dafny.Seq.Concat(_module.__default.fm19(p1, p1, globalState), _1095_v15);
      let _1097_v17;
      let _nw157 = Array((new BigNumber(14)).toNumber());
      _nw157[(_dafny.ZERO).toNumber()] = _1080_v1;
      _nw157[(_dafny.ONE).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(2)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(3)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(4)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(5)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(6)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(7)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(8)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(9)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(10)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(11)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(12)).toNumber()] = _1080_v1;
      _nw157[(new BigNumber(13)).toNumber()] = _1080_v1;
      _1097_v17 = _nw157;
      let _index188 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length));
      (_1097_v17)[_index188] = _1080_v1;
      let _1098_v18;
      _1098_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(p0),_1095_v15);
      let _1099_v19;
      _1099_v19 = _module.D1.create_DC2(p0);
      let _1100_v20;
      _1100_v20 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-938)), ((_1101_v12) => function (_1102_i1) {
        return _1101_v12;
      })(_1092_v12)));
      let _1103_v21;
      _1103_v21 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
      let _1104_v22;
      _1104_v22 = _dafny.Map.Empty.slice().updateUnsafe((p2) && (p1),_1080_v1);
      let _index189 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length));
      let _index190 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length));
      let _rhs247 = (p0).plus(p0);
      let _rhs248 = _dafny.Seq.update((((_1098_v18).contains(_1099_v19)) ? ((_1098_v18).get(_1099_v19)) : (_1095_v15)), _module.__default.safeIndex(p0, new BigNumber(((((_1098_v18).contains(_1099_v19)) ? ((_1098_v18).get(_1099_v19)) : (_1095_v15))).length)), _1092_v12);
      let _rhs249 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1095_v15, _1095_v15), (_1100_v20)[_module.__default.safeIndex(new BigNumber((_1103_v21).length), new BigNumber((_1100_v20).length))]);
      let _rhs250 = (((_1104_v22).contains((p0).isLessThan(p0))) ? ((_1104_v22).get((p0).isLessThan(p0))) : (_1080_v1));
      let _lhs179 = globalState;
      let _lhs180 = _1096_v16;
      let _lhs181 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length));
      let _lhs182 = _1097_v17;
      let _lhs183 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length));
      _lhs179.f15 = _rhs247;
      _lhs180[_lhs181] = _rhs248;
      _1095_v15 = _rhs249;
      _lhs182[_lhs183] = _rhs250;
      if (false) {
        let _1105_v23;
        let _nw158 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1105_v23 = _nw158;
        let _index191 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        (_1105_v23)[_index191] = _dafny.Seq.contains(_1095_v15, _1092_v12);
        let _index192 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        (_1105_v23)[_index192] = p1;
        let _1106_v24;
        _1106_v24 = _module.D0.create_DC1(_1105_v23);
        _1106_v24 = _1106_v24;
        let _1107_v25;
        _1107_v25 = _dafny.Seq.of((_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))], false);
        let _1108_v26;
        _1108_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1107_v25,(((_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))]) ? (false) : (p1)));
        let _1109_v27;
        _1109_v27 = _dafny.Seq.of(_dafny.Set.fromElements(_1092_v12));
        let _index193 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        (_1105_v23)[_index193] = (((_1108_v26).contains(_dafny.Seq.of(p1))) ? ((_1108_v26).get(_dafny.Seq.of(p1))) : (((((_1089_v9)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length))]).contains((_dafny.ZERO).minus(new BigNumber((_1109_v27).length)))) ? (((_1089_v9)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length))]).get((_dafny.ZERO).minus(new BigNumber((_1109_v27).length)))) : (p1))));
        let _1110_v28;
        _1110_v28 = _module.D0.create_DC0(_1105_v23);
        let _index194 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        let _index195 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        let _rhs251 = (p0).multipliedBy(p0);
        let _rhs252 = (((_1090_v10).contains(p0)) ? ((_1090_v10).get(p0)) : (p1));
        let _rhs253 = _1110_v28;
        let _rhs254 = !(p2);
        let _rhs255 = _dafny.Seq.Concat(_1107_v25, _1107_v25);
        let _lhs184 = globalState;
        let _lhs185 = _1105_v23;
        let _lhs186 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        let _lhs187 = _1105_v23;
        let _lhs188 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
        _lhs184.f15 = _rhs251;
        _lhs185[_lhs186] = _rhs252;
        _1110_v28 = _rhs253;
        _lhs187[_lhs188] = _rhs254;
        _1107_v25 = _rhs255;
        let _1111_v29;
        _1111_v29 = _module.D1.create_DC3(_1095_v15, p1);
        if (!((_1111_v29).dtor_cf4) || ((_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))])) {
          r2 = (_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))];
          let _1112_v30;
          _1112_v30 = _dafny.Seq.of(_1105_v23);
          let _1113_v31;
          _1113_v31 = _dafny.Set.fromElements(new BigNumber((_1112_v30).length));
          let _1114_v32;
          _1114_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v31,p1);
          let _1115_v33;
          _1115_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          r2 = (((_1114_v32).contains((_dafny.Set.fromElements(p0, (((_1079_v0).contains(p1)) ? ((_1079_v0).get(p1)) : (p0)))).Difference(_dafny.Set.fromElements(new BigNumber((_1115_v33).length), new BigNumber(721))))) ? ((_1114_v32).get((_dafny.Set.fromElements(p0, (((_1079_v0).contains(p1)) ? ((_1079_v0).get(p1)) : (p0)))).Difference(_dafny.Set.fromElements(new BigNumber((_1115_v33).length), new BigNumber(721))))) : (!((_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))])));
          let _index196 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index196] = _module.__default.safeModuloInt(p0, p0);
          let _index197 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index197] = p0;
          _1095_v15 = _dafny.Seq.Concat((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))], _1095_v15);
          _1100_v20 = _1100_v20;
        } else {
          let _1116_v34;
          _1116_v34 = _module.D3.create_DC9(_1092_v12);
          _1092_v12 = (_1116_v34).dtor_cf12;
          let _index198 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
          let _rhs256 = p0;
          let _rhs257 = p0;
          let _rhs258 = p1;
          let _lhs189 = globalState;
          let _lhs190 = globalState;
          let _lhs191 = _1105_v23;
          let _lhs192 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
          _lhs189.f4 = _rhs256;
          _lhs190.f4 = _rhs257;
          _lhs191[_lhs192] = _rhs258;
          let _1117_v35;
          let _nw159 = new _module.C0();
          _nw159.__ctor();
          _1117_v35 = _nw159;
          let _index199 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length));
          (_1105_v23)[_index199] = (_module.__default.safeModuloInt(p0, p0)).isEqualTo(p0);
          let _1118_v36;
          _1118_v36 = _module.D15.create_DC38(_1090_v10);
          let _1119_v37;
          _1119_v37 = _dafny.MultiSet.fromElements((_1118_v36).dtor_cf55, ((p1) ? (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p2)) : (_module.__default.fm35(globalState))), (_1118_v36).dtor_cf55);
          let _rhs259 = _1119_v37;
          let _rhs260 = (_1105_v23)[_module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1105_v23).length))];
          _1119_v37 = _rhs259;
          r2 = _rhs260;
        }
      } else {
        if (_dafny.Seq.IsPrefixOf((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))], _dafny.Seq.update((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))], _module.__default.safeIndex(p0, new BigNumber(((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]).length)), _1092_v12))) {
          (globalState).f15 = p0;
          let _1120_v38;
          let _nw160 = new _module.C4();
          _nw160.__ctor();
          _1120_v38 = _nw160;
          let _1121_v39;
          _1121_v39 = _dafny.Seq.of(p2, (_1120_v38).fm8(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()), p1, p0, globalState), p2, p2);
          _1079_v0 = (_module.__default.fm34(new BigNumber(603), new BigNumber((_1121_v39).length), p2, globalState)).Union(_1079_v0);
          r2 = p2;
          let _1122_v40;
          _1122_v40 = _dafny.Seq.of(_1089_v9);
          let _1123_v41;
          _1123_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1089_v9);
          let _1124_v42;
          _1124_v42 = _dafny.Set.fromElements(false);
          let _1125_v43;
          let _nw161 = Array((new BigNumber(28)).toNumber());
          _nw161[(_dafny.ZERO).toNumber()] = _1089_v9;
          _nw161[(_dafny.ONE).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(2)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(3)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(4)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(5)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(6)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(7)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(8)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(9)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(10)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(11)).toNumber()] = (_1122_v40)[_module.__default.safeIndex(p0, new BigNumber((_1122_v40).length))];
          _nw161[(new BigNumber(12)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(13)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(14)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(15)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(16)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(17)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(18)).toNumber()] = ((p2) ? (_1089_v9) : ((((_1123_v41).contains(new BigNumber((_1124_v42).length))) ? ((_1123_v41).get(new BigNumber((_1124_v42).length))) : ((_1122_v40)[_module.__default.safeIndex(p0, new BigNumber((_1122_v40).length))]))));
          _nw161[(new BigNumber(19)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(20)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(21)).toNumber()] = ((false) ? (_1089_v9) : (_1089_v9));
          _nw161[(new BigNumber(22)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(23)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(24)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(25)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(26)).toNumber()] = _1089_v9;
          _nw161[(new BigNumber(27)).toNumber()] = _1089_v9;
          _1125_v43 = _nw161;
          let _index200 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1125_v43).length));
          (_1125_v43)[_index200] = _1089_v9;
          let _index201 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_1125_v43).length));
          (_1125_v43)[_index201] = _1089_v9;
        } else {
          (globalState).f4 = _module.__default.safeModuloInt(p0, new BigNumber(887));
          r2 = !(!((p0).multipliedBy(new BigNumber(-924))).isEqualTo(_module.__default.safeDivisionInt(p0, p0)));
          let _1126_v44;
          _1126_v44 = _dafny.Seq.of(new BigNumber((_module.__default.fm33(p1, p1, p0, globalState)).length), p0);
          r2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1126_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-375)), function (_1127_i2) {
            return new BigNumber(535);
          })), _1126_v44);
          let _1128_v45;
          _1128_v45 = _dafny.Seq.of(!(!(p2)) || (p2));
          let _1129_v46;
          _1129_v46 = _dafny.Seq.of(_1103_v21);
          let _1130_v47;
          _1130_v47 = _dafny.Seq.of(_1129_v46);
          let _1131_v48;
          _1131_v48 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray((_1130_v47)[_module.__default.safeIndex(new BigNumber(((_1089_v9)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length))]).length), new BigNumber((_1130_v47).length))])).cardinality()));
          let _rhs261 = p0;
          let _rhs262 = _dafny.Seq.of(p2, p1, !((_1131_v48).IsSubsetOf(_1131_v48)));
          let _lhs193 = globalState;
          _lhs193.f15 = _rhs261;
          _1128_v45 = _rhs262;
          r2 = false;
        }
        let _1132_v49;
        _1132_v49 = _dafny.MultiSet.fromElements(p0, p0);
        let _1133_v50;
        _1133_v50 = _dafny.Map.Empty.slice().updateUnsafe(((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))])[_module.__default.safeIndex(p0, new BigNumber(((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]).length))],_1132_v49);
        _1133_v50 = (_1133_v50).update(_1092_v12, (_1132_v49).Union(_1132_v49));
        if (_module.__default.fm26(globalState)) {
          let _index202 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index202] = p0;
          let _index203 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index203] = (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1080_v1,_module.__default.fm26(globalState))).length), p0)).minus(_module.__default.fm17(p0, globalState));
          r2 = p2;
          r2 = (p1) === (((p1) ? (p1) : (p1)));
          let _index204 = _module.__default.safeIndex(new BigNumber(862), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index204] = (p0).multipliedBy((_1080_v1)[_module.__default.safeIndex(new BigNumber(862), new BigNumber((_1080_v1).length))]);
          r2 = true;
        } else {
          let _1134_v51;
          let _nw162 = new _module.C2();
          _nw162.__ctor();
          _1134_v51 = _nw162;
          let _1135_v52;
          let _init29 = ((_1136_p1) => function (_1137_i3) {
            return _1136_p1;
          })(p1);
          let _nw163 = Array((new BigNumber(19)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw163.length); _i0_29++) {
            _nw163[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1135_v52 = _nw163;
          let _1138_v53;
          _1138_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_dafny.Seq.of(_1135_v52)).length));
          (globalState).f13 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1138_v53).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-592)), ((_1139_v12) => function (_1140_i4) {
            return (_module.D3.create_DC9(_1139_v12)).dtor_cf12;
          })(_1092_v12))).length)));
          r2 = p2;
          let _index205 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1135_v52).length));
          (_1135_v52)[_index205] = p1;
          let _index206 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_1135_v52).length));
          (_1135_v52)[_index206] = !(p2) || (p2);
          let _arr0 = (_1097_v17)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length))];
          let _index207 = _module.__default.safeIndex(new BigNumber(200), new BigNumber(((_1097_v17)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length))]).length));
          _arr0[_index207] = new BigNumber((_1095_v15).length);
          let _arr1 = (_1097_v17)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length))];
          let _index208 = _module.__default.safeIndex(new BigNumber(200), new BigNumber(((_1097_v17)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length))]).length));
          _arr1[_index208] = p0;
        }
        if (true) {
          r2 = (((_1103_v21).contains(!(p2))) ? ((_1103_v21).get(!(p2))) : ((_1079_v0).IsProperSubsetOf(_1079_v0)));
          let _1141_v54;
          _1141_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1142_v57;
          let _nw164 = Array((new BigNumber(14)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(387),p0);
          _nw164[(_dafny.ONE).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(2)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(3)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(4)).toNumber()] = (_1141_v54).update((_dafny.ZERO).minus(p0), p0);
          _nw164[(new BigNumber(5)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(6)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(7)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(8)).toNumber()] = ((_1141_v54).update(p0, new BigNumber(733))).Merge(_1141_v54);
          _nw164[(new BigNumber(9)).toNumber()] = ((p2) ? (_1141_v54) : (_1141_v54));
          _nw164[(new BigNumber(10)).toNumber()] = function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of _dafny.IntegerRange(new BigNumber(337), new BigNumber(833))) {
              let _1143_v55 = _compr_50;
              if (((new BigNumber(337)).isLessThanOrEqualTo(_1143_v55)) && ((_1143_v55).isLessThan(new BigNumber(833)))) {
                _coll50.push([(_1143_v55).minus(p0),p0]);
              }
            }
            return _coll50;
          }();
          _nw164[(new BigNumber(11)).toNumber()] = function () {
            let _coll51 = new _dafny.Map();
            for (const _compr_51 of _dafny.IntegerRange(new BigNumber(132), new BigNumber(774))) {
              let _1144_v56 = _compr_51;
              if (((new BigNumber(132)).isLessThanOrEqualTo(_1144_v56)) && ((_1144_v56).isLessThan(new BigNumber(774)))) {
                _coll51.push([_module.__default.safeDivisionInt(_1144_v56, p0),p0]);
              }
            }
            return _coll51;
          }();
          _nw164[(new BigNumber(12)).toNumber()] = _1141_v54;
          _nw164[(new BigNumber(13)).toNumber()] = _1141_v54;
          _1142_v57 = _nw164;
          _1142_v57 = _1142_v57;
          let _1145_v58;
          let _nw165 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1145_v58 = _nw165;
          let _index209 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1145_v58).length));
          (_1145_v58)[_index209] = (_this).fm16(p0, globalState);
          let _1146_v59;
          _1146_v59 = _dafny.Seq.of(false, p1);
          let _1147_v60;
          _1147_v60 = _dafny.Seq.of(_1132_v49, _dafny.MultiSet.fromElements(p0, new BigNumber(209)));
          let _index210 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1145_v58).length));
          (_1145_v58)[_index210] = (_dafny.MultiSet.FromArray(_1147_v60)).contains(((_1132_v49).update(new BigNumber((_1146_v59).length), _module.__default.abs(p0))).Intersect(_dafny.MultiSet.fromElements(p0)));
          let _index211 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1145_v58).length));
          (_1145_v58)[_index211] = p2;
          let _index212 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_1145_v58).length));
          (_1145_v58)[_index212] = p1;
        } else {
          let _index213 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length));
          (_1096_v16)[_index213] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-551)), ((_1148_v12) => function (_1149_i5) {
            return _1148_v12;
          })(_1092_v12)), (_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]);
          let _1150_v61;
          _1150_v61 = _module.D4.create_DC12(p0);
          let _1151_v62;
          _1151_v62 = _dafny.MultiSet.fromElements((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]);
          let _1152_v63;
          _1152_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1150_v61,new BigNumber((_1151_v62).cardinality()));
          _1152_v63 = (_1152_v63).update(_module.D4.create_DC12(p0), p0);
          r2 = p2;
          (globalState).f15 = (_dafny.ZERO).minus(p0);
          let _index214 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1080_v1).length));
          (_1080_v1)[_index214] = p0;
          let _1153_v64;
          let _init30 = ((_1154_p1) => function (_1155_i6) {
            return ((_1154_p1) ? (_1154_p1) : (_1154_p1));
          })(p1);
          let _nw166 = Array((new BigNumber(9)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw166.length); _i0_30++) {
            _nw166[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1153_v64 = _nw166;
          let _index215 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_1153_v64).length));
          (_1153_v64)[_index215] = (p2) === (p1);
          let _1156_v65;
          _1156_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          let _1157_v66;
          _1157_v66 = _dafny.Seq.of(new BigNumber(885), p0);
          let _index216 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1080_v1).length));
          let _index217 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_1153_v64).length));
          let _rhs263 = new BigNumber(((_1156_v65).Merge(_1156_v65)).length);
          let _rhs264 = (_dafny.MultiSet.fromElements(!(p1))).update(_module.__default.fm26(globalState), _module.__default.abs((p0).multipliedBy((_1157_v66)[_module.__default.safeIndex(p0, new BigNumber((_1157_v66).length))])));
          let _rhs265 = (((_1079_v0).update(p1, _module.__default.abs(p0))).Intersect(_1079_v0)).IsSubsetOf(_dafny.MultiSet.fromElements(p1));
          let _rhs266 = p0;
          let _lhs194 = _1080_v1;
          let _lhs195 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_1080_v1).length));
          let _lhs196 = _1153_v64;
          let _lhs197 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_1153_v64).length));
          let _lhs198 = globalState;
          _lhs194[_lhs195] = _rhs263;
          _1079_v0 = _rhs264;
          _lhs196[_lhs197] = _rhs265;
          _lhs198.f4 = _rhs266;
        }
        (globalState).f15 = p0;
      }
      let _1158_v67;
      _1158_v67 = _module.D2.create_DC7(p1, _1079_v0, (_1097_v17)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_1097_v17).length))]);
      let _1159_v68;
      _1159_v68 = _module.D2.create_DC8(_1158_v67);
      let _1160_v69;
      let _nw167 = new _module.C3();
      _nw167.__ctor(_1159_v68, false);
      _1160_v69 = _nw167;
      let _hi7 = p0;
      for (let _1161_i7 = p0; _1161_i7.isLessThan(_hi7); _1161_i7 = _1161_i7.plus(_dafny.ONE)) {
        let _index218 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_1080_v1).length));
        (_1080_v1)[_index218] = new BigNumber(-834);
        let _1162_v70;
        _1162_v70 = _dafny.Set.fromElements(_1161_i7);
        let _1163_v71;
        _1163_v71 = _dafny.Seq.of(_1162_v70);
        let _index219 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_1080_v1).length));
        let _rhs267 = p0;
        let _rhs268 = (_dafny.ZERO).minus(new BigNumber((_1163_v71).length));
        let _lhs199 = _1080_v1;
        let _lhs200 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_1080_v1).length));
        let _lhs201 = globalState;
        _lhs199[_lhs200] = _rhs267;
        _lhs201.f15 = _rhs268;
        r2 = _1160_v69.f25;
        r2 = !(_1160_v69.f25);
        let _1164_v72;
        _1164_v72 = _dafny.MultiSet.fromElements(p0, (_1080_v1)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_1080_v1).length))]);
        (globalState).f4 = (((_1164_v72).contains(p0)) ? ((_1164_v72).get(p0)) : (p0));
      }
      r0 = (_1160_v69).fm2((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber(((_1089_v9)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length))]).length))), globalState);
      let _1165_v73;
      _1165_v73 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1166_v74;
      _1166_v74 = _dafny.Seq.of(_dafny.MultiSet.fromElements((((_1165_v73).contains(p0)) ? ((_1165_v73).get(p0)) : (p0))));
      let _1167_v75;
      _1167_v75 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]).length)));
      let _1168_v76;
      _1168_v76 = _dafny.MultiSet.fromElements(new BigNumber((_1167_v75).length), new BigNumber(((_1096_v16)[_module.__default.safeIndex(new BigNumber(736), new BigNumber((_1096_v16).length))]).length));
      r1 = ((_1166_v74)[_module.__default.safeIndex(new BigNumber(761), new BigNumber((_1166_v74).length))]).Difference((_1168_v76).update(p0, _module.__default.abs(p0)));
      let _1169_v77;
      _1169_v77 = _module.D15.create_DC38((_1089_v9)[_module.__default.safeIndex(new BigNumber(998), new BigNumber((_1089_v9).length))]);
      let _1170_v78;
      _1170_v78 = _module.D15.create_DC40(_1169_v77);
      let _1171_v79;
      _1171_v79 = _module.D15.create_DC40(_1169_v77);
      let _1172_v80;
      _1172_v80 = _module.D15.create_DC40(_1169_v77);
      let _pat_let_tv15 = _1160_v69;
      let _pat_let_tv16 = _1160_v69;
      let _pat_let_tv17 = _1160_v69;
      r2 = function (_source19) {
        if (_source19.is_DC39) {
          let _1173___mcc_h0 = (_source19).cf56;
          let _1174___mcc_h1 = (_source19).cf57;
          let _1175___mcc_h2 = (_source19).cf58;
          let _1176_cf58 = _1175___mcc_h2;
          let _1177_cf57 = _1174___mcc_h1;
          let _1178_cf56 = _1173___mcc_h0;
          return _dafny.areEqual(_dafny.Seq.of(_1177_cf57), _dafny.Seq.of(_pat_let_tv15.f25));
        } else if (_source19.is_DC38) {
          let _1179___mcc_h3 = (_source19).cf55;
          let _1180_cf55 = _1179___mcc_h3;
          return _pat_let_tv16.f25;
        } else {
          let _1181___mcc_h4 = (_source19).cf59;
          let _1182_cf59 = _1181___mcc_h4;
          return _pat_let_tv17.f25;
        }
      }(_1172_v80);
      return [r0, r1, r2];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("bppvmww"), new _dafny.CodePoint('w'.codePointAt(0)));
    };
    fm2(p0, globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), function (_1183_i0) {
        return new BigNumber((((true) ? (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true, !(false), true)).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), function (_1184_i1) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("po")).length));
        })))).length);
      }))).cardinality());
    };
    fm12(globalState) {
      let _this = this;
      return (true) === (!((_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("dtaill"), true)).dtor_cf4));
    };
    fm13(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jkonhsjjq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_1185_i0) {
        return (_module.D3.create_DC9(new _dafny.CodePoint('l'.codePointAt(0)))).dtor_cf12;
      }));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _1186_v0;
      _1186_v0 = true;
      let _1187_v1;
      _1187_v1 = _module.D1.create_DC3(p0, _1186_v0);
      _1186_v0 = (((_1187_v1).dtor_cf4) ? (_1186_v0) : (_1186_v0));
      let _1188_v2;
      _1188_v2 = new BigNumber(792);
      let _1189_v3;
      _1189_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1188_v2,false);
      let _1190_v4;
      let _nw168 = Array((new BigNumber(5)).toNumber());
      _nw168[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_1188_v2);
      _nw168[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_1189_v3).length)));
      _nw168[(new BigNumber(2)).toNumber()] = _1188_v2;
      _nw168[(new BigNumber(3)).toNumber()] = _1188_v2;
      _nw168[(new BigNumber(4)).toNumber()] = _1188_v2;
      _1190_v4 = _nw168;
      let _1191_v5;
      _1191_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1190_v4,new BigNumber((_dafny.Seq.of(_1188_v2)).length));
      _1191_v5 = _1191_v5;
      let _1192_v6;
      _1192_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v0,_1188_v2);
      _1192_v6 = _1192_v6;
      let _1193_v7;
      _1193_v7 = _dafny.MultiSet.fromElements(_1186_v0);
      let _1194_v8;
      _1194_v8 = _dafny.Seq.of(new BigNumber(78));
      let _1195_i0;
      _1195_i0 = _dafny.ZERO;
      L7: {
        while (!(!((((_1186_v0) ? (_1188_v2) : (_1188_v2))).isLessThanOrEqualTo((((_1193_v7).contains(!(_1186_v0))) ? ((_1193_v7).get(!(_1186_v0))) : ((_1194_v8)[_module.__default.safeIndex(_1188_v2, new BigNumber((_1194_v8).length))])))))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1195_i0)) {
              break L7;
            }
            _1195_i0 = (_1195_i0).plus(_dafny.ONE);
            let _1196_v9;
            _1196_v9 = _dafny.Seq.UnicodeFromString("vkph");
            let _1197_v10;
            _1197_v10 = new _dafny.CodePoint('a'.codePointAt(0));
            let _1198_v11;
            _1198_v11 = _dafny.Seq.of(_1193_v7, _1193_v7, _1193_v7);
            let _rhs269 = (_1198_v11)[_module.__default.safeIndex(_1188_v2, new BigNumber((_1198_v11).length))];
            let _rhs270 = ((_1186_v0) ? (_1196_v9) : (_1196_v9));
            let _rhs271 = _1197_v10;
            let _rhs272 = new BigNumber((_1196_v9).length);
            _1193_v7 = _rhs269;
            _1196_v9 = _rhs270;
            _1197_v10 = _rhs271;
            _1188_v2 = _rhs272;
            let _1199_v12;
            _1199_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v0,_dafny.Seq.of(_1188_v2, _1188_v2));
            let _1200_v13;
            _1200_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat((((_1199_v12).contains(_1186_v0)) ? ((_1199_v12).get(_1186_v0)) : (_1194_v8)), _1194_v8),_1186_v0);
            let _1201_v14;
            _1201_v14 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-69)), ((_1202_v10) => function (_1203_i1) {
              return _1202_v10;
            })(_1197_v10))).length), _1188_v2);
            if ((((_1200_v13).contains(_1194_v8)) ? ((_1200_v13).get(_1194_v8)) : ((_1201_v14).IsSubsetOf(_1201_v14)))) {
              let _rhs273 = _module.__default.safeModuloInt((_1188_v2).multipliedBy(_1188_v2), _1188_v2);
              let _rhs274 = (_1194_v8)[_module.__default.safeIndex(_1188_v2, new BigNumber((_1194_v8).length))];
              let _lhs202 = globalState;
              _lhs202.f7 = _rhs273;
              _1188_v2 = _rhs274;
              _1190_v4 = ((_1186_v0) ? (_1190_v4) : (_1190_v4));
              (globalState).f4 = (_dafny.ZERO).minus(_1188_v2);
              (globalState).f15 = (_dafny.ZERO).minus(_1188_v2);
              (globalState).f4 = (_this).fm2(_1188_v2, globalState);
            } else {
              let _1204_v15;
              _1204_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v0,_1186_v0);
              _1186_v0 = (((_1204_v15).contains(false)) ? ((_1204_v15).get(false)) : (_1186_v0));
              let _1205_v16;
              _1205_v16 = _dafny.Seq.of(_1186_v0, _1186_v0);
              let _1206_v17;
              _1206_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v0,_1204_v15);
              let _1207_v18;
              let _nw169 = Array((new BigNumber(27)).toNumber());
              _nw169[(_dafny.ZERO).toNumber()] = _1186_v0;
              _nw169[(_dafny.ONE).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(2)).toNumber()] = false;
              _nw169[(new BigNumber(3)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(4)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(5)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(6)).toNumber()] = !(false);
              _nw169[(new BigNumber(7)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(8)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(9)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(10)).toNumber()] = !(_1186_v0);
              _nw169[(new BigNumber(11)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(12)).toNumber()] = (_1205_v16)[_module.__default.safeIndex(_1188_v2, new BigNumber((_1205_v16).length))];
              _nw169[(new BigNumber(13)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(14)).toNumber()] = true;
              _nw169[(new BigNumber(15)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(16)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(17)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(18)).toNumber()] = (_this).fm1(_1206_v17, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1188_v2,_1188_v2)).length), globalState);
              _nw169[(new BigNumber(19)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(20)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(21)).toNumber()] = (_1205_v16)[_module.__default.safeIndex(_1188_v2, new BigNumber((_1205_v16).length))];
              _nw169[(new BigNumber(22)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(23)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(24)).toNumber()] = _1186_v0;
              _nw169[(new BigNumber(25)).toNumber()] = true;
              _nw169[(new BigNumber(26)).toNumber()] = (((_1189_v3).contains(new BigNumber((_dafny.Set.fromElements(_1188_v2, _1188_v2, _1188_v2)).length))) ? ((_1189_v3).get(new BigNumber((_dafny.Set.fromElements(_1188_v2, _1188_v2, _1188_v2)).length))) : (_1186_v0));
              _1207_v18 = _nw169;
              let _1208_v19;
              _1208_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1207_v18,((_1186_v0) ? (!(!(false))) : (_1186_v0)));
              _1186_v0 = (((_1208_v19).contains(_1207_v18)) ? ((_1208_v19).get(_1207_v18)) : (false));
              let _index220 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length));
              (_1207_v18)[_index220] = (_dafny.Map.Empty.slice().updateUnsafe(false,_1188_v2)).contains((_this).fm1(_1206_v17, _1188_v2, globalState));
              let _1209_v20;
              _1209_v20 = _dafny.Set.fromElements(_1186_v0);
              let _index221 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length));
              (_1207_v18)[_index221] = (_1209_v20).IsDisjointFrom(_1209_v20);
              let _1210_v21;
              _1210_v21 = _dafny.Seq.of(_dafny.Seq.of(_1188_v2), _1194_v8);
              let _rhs275 = _1186_v0;
              let _rhs276 = _1210_v21;
              let _rhs277 = _1188_v2;
              let _rhs278 = (_1206_v17).Merge((_1206_v17).Merge(_1206_v17));
              _1186_v0 = _rhs275;
              _1210_v21 = _rhs276;
              _1188_v2 = _rhs277;
              _1206_v17 = _rhs278;
              let _1211_v22;
              _1211_v22 = _module.D3.create_DC9(_1197_v10);
              let _1212_v23;
              _1212_v23 = _dafny.Seq.of(_1207_v18);
              let _1213_v24;
              let _nw170 = Array((new BigNumber(7)).toNumber());
              _nw170[(_dafny.ZERO).toNumber()] = _1207_v18;
              _nw170[(_dafny.ONE).toNumber()] = _1207_v18;
              _nw170[(new BigNumber(2)).toNumber()] = _1207_v18;
              _nw170[(new BigNumber(3)).toNumber()] = (_1212_v23)[_module.__default.safeIndex(new BigNumber(613), new BigNumber((_1212_v23).length))];
              _nw170[(new BigNumber(4)).toNumber()] = _1207_v18;
              _nw170[(new BigNumber(5)).toNumber()] = _1207_v18;
              _nw170[(new BigNumber(6)).toNumber()] = _1207_v18;
              _1213_v24 = _nw170;
              let _index222 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1213_v24).length));
              (_1213_v24)[_index222] = _1207_v18;
              let _1214_v25;
              _1214_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1186_v0, (_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))], (_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))])).length),_1190_v4);
              let _1215_v26;
              _1215_v26 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ejjgk"));
              let _1216_v28;
              _1216_v28 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll52 = new _dafny.Map();
                for (const _compr_52 of _dafny.IntegerRange(new BigNumber(803), new BigNumber(-816))) {
                  let _1217_v27 = _compr_52;
                  if (((new BigNumber(803)).isLessThanOrEqualTo(_1217_v27)) && ((_1217_v27).isLessThan(new BigNumber(-816)))) {
                    _coll52.push([_module.__default.safeModuloInt(_1217_v27, _1188_v2),_1186_v0]);
                  }
                }
                return _coll52;
              }(),!(!((_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))])));
              let _index223 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length));
              let _index224 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1213_v24).length));
              let _rhs279 = ((new BigNumber((_dafny.MultiSet.fromElements(_1186_v0)).cardinality())).minus(new BigNumber((_1214_v25).length))).isEqualTo((new BigNumber(301)).minus(_1188_v2));
              let _rhs280 = _module.__default.fm14(new BigNumber((_1201_v14).cardinality()), _1215_v26, _1188_v2, (_1216_v28).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1189_v3,(_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))])).update((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1205_v16).length),(_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))])).update(new BigNumber((_dafny.Set.fromElements(new BigNumber(946), _1188_v2)).length), _1186_v0), (_1207_v18)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length))])), globalState);
              let _rhs281 = _1207_v18;
              let _lhs203 = _1207_v18;
              let _lhs204 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1207_v18).length));
              let _lhs205 = _1213_v24;
              let _lhs206 = _module.__default.safeIndex(new BigNumber(49), new BigNumber((_1213_v24).length));
              _lhs203[_lhs204] = _rhs279;
              _1211_v22 = _rhs280;
              _lhs205[_lhs206] = _rhs281;
            }
            let _index225 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length));
            (_1190_v4)[_index225] = new BigNumber(((_1189_v3).update(_1188_v2, !(_1186_v0))).length);
            let _index226 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length));
            (_1190_v4)[_index226] = ((_1186_v0) ? ((_dafny.ZERO).minus(_1188_v2)) : (_1188_v2));
            let _1218_v29;
            _1218_v29 = _module.D1.create_DC2(_1188_v2);
            if (((_1218_v29).dtor_cf2).isLessThan((_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))])) {
              (globalState).f13 = (_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))];
              let _1219_v30;
              let _nw171 = new _module.C0();
              _nw171.__ctor();
              _1219_v30 = _nw171;
              let _1220_v31;
              let _nw172 = new _module.C2();
              _nw172.__ctor();
              _1220_v31 = _nw172;
              let _1221_v32;
              _1221_v32 = _dafny.Seq.of(_1186_v0);
              let _1222_v33;
              _1222_v33 = _module.D11.create_DC30(_1186_v0, (_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))]);
              let _1223_v34;
              _1223_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1221_v32,_1222_v33);
              let _1224_v35;
              _1224_v35 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, _1186_v0),_module.D11.create_DC30(_1186_v0, new BigNumber(364)))).Merge(_1223_v34),_1186_v0);
              _1224_v35 = (_1224_v35).update((_1223_v34).Merge(_1223_v34), _1186_v0);
              (globalState).f13 = (_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))];
            } else {
              let _1225_v36;
              _1225_v36 = _module.D4.create_DC12((_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))]);
              let _pat_let_tv18 = _1190_v4;
              let _pat_let_tv19 = _1190_v4;
              let _1226_v37;
              _1226_v37 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let16_0) {
                return function (_1227_dt__update__tmp_h0) {
                  return function (_pat_let17_0) {
                    return function (_1228_dt__update_hcf16_h0) {
                      return _module.D4.create_DC12(_1228_dt__update_hcf16_h0);
                    }(_pat_let17_0);
                  }((_pat_let_tv19)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_pat_let_tv18).length))]);
                }(_pat_let16_0);
              }(_1225_v36),(_1190_v4)[_module.__default.safeIndex(new BigNumber(706), new BigNumber((_1190_v4).length))]);
              let _rhs282 = (_1226_v37).Merge(_1226_v37);
              let _rhs283 = _1188_v2;
              let _rhs284 = new BigNumber(660);
              let _lhs207 = globalState;
              let _lhs208 = globalState;
              _1226_v37 = _rhs282;
              _lhs207.f13 = _rhs283;
              _lhs208.f4 = _rhs284;
              let _1229_v38;
              _1229_v38 = _dafny.Set.fromElements(!(_1186_v0), !(false));
              let _1230_v39;
              _1230_v39 = _module.D12.create_DC32(((true) ? (_1229_v38) : (_1229_v38)));
              _1230_v39 = _module.D12.create_DC32(_1229_v38);
              (globalState).f7 = new BigNumber(641);
              let _1231_v40;
              let _init31 = ((_1232_v0) => function (_1233_i2) {
                return _dafny.Seq.of(_1232_v0, _1232_v0);
              })(_1186_v0);
              let _nw173 = Array((new BigNumber(15)).toNumber());
              for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw173.length); _i0_31++) {
                _nw173[_i0_31] = _init31(new BigNumber(_i0_31));
              }
              _1231_v40 = _nw173;
              let _1234_v41;
              _1234_v41 = _1231_v40;
              let _1235_v42;
              _1235_v42 = _dafny.Map.Empty.slice().updateUnsafe((_1201_v14).Union(_dafny.MultiSet.fromElements(new BigNumber((_1196_v9).length))),_1234_v41);
              _1235_v42 = (_1235_v42).update(_1201_v14, _1234_v41);
              let _1236_v43;
              let _nw174 = Array((new BigNumber(7)).toNumber()).fill(false);
              _1236_v43 = _nw174;
              let _1237_v44;
              _1237_v44 = _dafny.Seq.of(_1236_v43, _1236_v43, _1236_v43, _1236_v43);
              let _1238_v45;
              _1238_v45 = _module.D7.create_DC19(_1237_v44);
              _1238_v45 = _1238_v45;
            }
          }
        }
      }
      let _1239_v46;
      _1239_v46 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1240_v47;
      let _nw175 = Array((new BigNumber(24)).toNumber());
      _nw175[(_dafny.ZERO).toNumber()] = true;
      _nw175[(_dafny.ONE).toNumber()] = (_1186_v0) && (true);
      _nw175[(new BigNumber(2)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(3)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(4)).toNumber()] = !(_1188_v2).isEqualTo(_1188_v2);
      _nw175[(new BigNumber(5)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(6)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(7)).toNumber()] = !(false);
      _nw175[(new BigNumber(8)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(9)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(10)).toNumber()] = !_dafny.Seq.contains(p0, _1239_v46);
      _nw175[(new BigNumber(11)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(12)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(13)).toNumber()] = (_1186_v0) === (_1186_v0);
      _nw175[(new BigNumber(14)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(15)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(16)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(17)).toNumber()] = _1186_v0;
      _nw175[(new BigNumber(18)).toNumber()] = (((_1189_v3).contains(new BigNumber(-201))) ? ((_1189_v3).get(new BigNumber(-201))) : (!(_1186_v0)));
      _nw175[(new BigNumber(19)).toNumber()] = (new BigNumber(55)).isEqualTo(_1188_v2);
      _nw175[(new BigNumber(20)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1194_v8, _1194_v8);
      _nw175[(new BigNumber(21)).toNumber()] = _module.__default.fm26(globalState);
      _nw175[(new BigNumber(22)).toNumber()] = (_this).fm12(globalState);
      _nw175[(new BigNumber(23)).toNumber()] = _1186_v0;
      _1240_v47 = _nw175;
      let _index227 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1240_v47).length));
      (_1240_v47)[_index227] = !(!(false));
      let _index228 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1240_v47).length));
      (_1240_v47)[_index228] = _1186_v0;
      let _1241_v48;
      _1241_v48 = _dafny.Seq.of(_1186_v0);
      let _index229 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((_1240_v47).length));
      (_1240_v47)[_index229] = (_1188_v2).isEqualTo(new BigNumber((_1241_v48).length));
      return;
    }
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(326)), function (_1242_i0) {
        return new BigNumber((_dafny.Seq.of(true)).length);
      }), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(new BigNumber(404))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("hwtfiavx")).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(766), new BigNumber(187), new BigNumber(520)), _dafny.Seq.of(new BigNumber(474))));
    };
    fm2(p0, globalState) {
      let _this = this;
      if (!((true) || (true))) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("bmw"), !(false))).dtor_cf3, _dafny.Seq.UnicodeFromString("yvx"))).length));
      } else {
        return new BigNumber(231);
      }
    };
    fm7(p0, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(function () {
        let _coll53 = new _dafny.Set();
        for (const _compr_53 of _dafny.IntegerRange(new BigNumber(-228), new BigNumber(300))) {
          let _1243_v0 = _compr_53;
          if (((new BigNumber(-228)).isLessThanOrEqualTo(_1243_v0)) && ((_1243_v0).isLessThan(new BigNumber(300)))) {
            _coll53.add((_1243_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(255))).cardinality()), new BigNumber(129))).cardinality())));
          }
        }
        return _coll53;
      }())).Union((_dafny.MultiSet.fromElements(function () {
        let _coll54 = new _dafny.Set();
        for (const _compr_54 of _dafny.IntegerRange(new BigNumber(-367), new BigNumber(626))) {
          let _1244_v1 = _compr_54;
          if (((new BigNumber(-367)).isLessThanOrEqualTo(_1244_v1)) && ((_1244_v1).isLessThan(new BigNumber(626)))) {
            _coll54.add((_1244_v1).multipliedBy(new BigNumber(122)));
          }
        }
        return _coll54;
      }(), function () {
        let _coll55 = new _dafny.Set();
        for (const _compr_55 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)))).length))).length))).Keys.Elements) {
          let _1245_v2 = _compr_55;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('w'.codePointAt(0)))).length))).length))).contains(_1245_v2)) {
            _coll55.add(_module.__default.safeModuloInt(_1245_v2, new BigNumber(477)));
          }
        }
        return _coll55;
      }())).Difference(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, true)).length)), _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(!(false), true))).cardinality()), new BigNumber(622)), _dafny.Set.fromElements(new BigNumber(18)), function () {
        let _coll56 = new _dafny.Set();
        for (const _compr_56 of (function () {
          let _coll57 = new _dafny.Map();
          for (const _compr_57 of _dafny.IntegerRange(new BigNumber(723), new BigNumber(-208))) {
            let _1246_v3 = _compr_57;
            if (((new BigNumber(723)).isLessThanOrEqualTo(_1246_v3)) && ((_1246_v3).isLessThan(new BigNumber(-208)))) {
              _coll57.push([(_1246_v3).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), function (_1248_i0) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              })).length)),new BigNumber((function () {
                let _coll58 = new _dafny.Map();
                for (const _compr_58 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("xypuautvo"))).Elements) {
                  let _1247_v4 = _compr_58;
                  if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("xypuautvo")), _1247_v4)) {
                    _coll58.push([_1247_v4,false]);
                  }
                }
                return _coll58;
              }()).length)]);
            }
          }
          return _coll57;
        }()).Keys.Elements) {
          let _1249_v5 = _compr_56;
          if ((function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of _dafny.IntegerRange(new BigNumber(723), new BigNumber(-208))) {
              let _1246_v3 = _compr_59;
              if (((new BigNumber(723)).isLessThanOrEqualTo(_1246_v3)) && ((_1246_v3).isLessThan(new BigNumber(-208)))) {
                _coll59.push([(_1246_v3).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(273)), function (_1248_i0) {
                  return new _dafny.CodePoint('a'.codePointAt(0));
                })).length)),new BigNumber((function () {
                  let _coll60 = new _dafny.Map();
                  for (const _compr_60 of (_dafny.Seq.of(_dafny.Seq.UnicodeFromString("xypuautvo"))).Elements) {
                    let _1247_v4 = _compr_60;
                    if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("xypuautvo")), _1247_v4)) {
                      _coll60.push([_1247_v4,false]);
                    }
                  }
                  return _coll60;
                }()).length)]);
              }
            }
            return _coll59;
          }()).contains(_1249_v5)) {
            _coll56.add((_1249_v5).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ubinwfp")).length)));
          }
        }
        return _coll56;
      }(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yhq")).length)))));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(false))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.of(true)))) {
        return true;
      } else {
        return (new BigNumber((_dafny.Seq.UnicodeFromString("ydocfi")).length)).isLessThan(new BigNumber(909));
      }
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _1250_v0;
      _1250_v0 = true;
      _1250_v0 = _1250_v0;
      let _1251_v1;
      let _init32 = ((_1252_v0) => function (_1253_i1) {
        return _1252_v0;
      })(_1250_v0);
      let _nw176 = Array((new BigNumber(6)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw176.length); _i0_32++) {
        _nw176[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1251_v1 = _nw176;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1251_v1).length))) {
        let _1254_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1254_i0)) && ((_1254_i0).isLessThan(new BigNumber((_1251_v1).length))))) {
          (_1251_v1)[(_1254_i0)] = !(_1250_v0) || (!(!(_1250_v0)));
        }
      }
      let _1255_v2;
      _1255_v2 = new BigNumber(551);
      let _1256_v4;
      _1256_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1255_v2,!(_1250_v0));
      let _hi8 = ((_1250_v0) ? (_1255_v2) : (new BigNumber((function () {
        let _coll61 = new _dafny.Map();
        for (const _compr_61 of (_1256_v4).Keys.Elements) {
          let _1257_v3 = _compr_61;
          if ((_1256_v4).contains(_1257_v3)) {
            _coll61.push([_module.__default.safeModuloInt(_1257_v3, _1255_v2),_1255_v2]);
          }
        }
        return _coll61;
      }()).length)));
      for (let _1258_i2 = new BigNumber((_module.__default.fm11(_1250_v0, globalState)).length); _1258_i2.isLessThan(_hi8); _1258_i2 = _1258_i2.plus(_dafny.ONE)) {
        let _1259_v5;
        let _1260_v6;
        let _out9;
        let _out10;
        let _outcollector3 = (_this).m8(globalState);
        _out9 = _outcollector3[0];
        _out10 = _outcollector3[1];
        _1259_v5 = _out9;
        _1260_v6 = _out10;
        let _1261_v7;
        _1261_v7 = _dafny.Seq.of(_1250_v0, _1250_v0);
        _1261_v7 = _dafny.Seq.Concat(_1261_v7, _1261_v7);
        _1250_v0 = (_1250_v0) && (_1250_v0);
        (globalState).f4 = _1259_v5;
      }
      let _1262_v8;
      let _nw177 = new _module.C5();
      _nw177.__ctor();
      _1262_v8 = _nw177;
      _1251_v1 = _1251_v1;
      let _1263_v9;
      _1263_v9 = _dafny.Map.Empty.slice().updateUnsafe(_1251_v1,_1250_v0);
      let _1264_i3;
      _1264_i3 = _dafny.ZERO;
      L8: {
        while ((((_1263_v9).contains(_1251_v1)) ? ((_1263_v9).get(_1251_v1)) : (_1250_v0))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1264_i3)) {
              break L8;
            }
            _1264_i3 = (_1264_i3).plus(_dafny.ONE);
            let _1265_v10;
            let _nw178 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
            _1265_v10 = _nw178;
            let _1266_v11;
            _1266_v11 = _dafny.MultiSet.fromElements(_1265_v10);
            let _1267_v12;
            _1267_v12 = _dafny.Map.Empty.slice().updateUnsafe((_1266_v11).IsSubsetOf((_1266_v11).update(_1265_v10, _module.__default.abs(_1255_v2))),(_1250_v0) === (_1250_v0));
            _1267_v12 = (_1267_v12).update(_1250_v0, !(_1250_v0));
            let _1268_v13;
            _1268_v13 = _module.D7.create_DC21(new BigNumber(637), new BigNumber(-530), _1250_v0, _1250_v0, _1250_v0);
            (globalState).f4 = (_1268_v13).dtor_cf31;
            let _1269_v14;
            let _init33 = ((_1270_v0) => function (_1271_i4) {
              return _dafny.MultiSet.fromElements(_1270_v0);
            })(_1250_v0);
            let _nw179 = Array((new BigNumber(28)).toNumber());
            for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw179.length); _i0_33++) {
              _nw179[_i0_33] = _init33(new BigNumber(_i0_33));
            }
            _1269_v14 = _nw179;
            let _1272_v15;
            _1272_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1269_v14,_1250_v0);
            if ((((_1272_v15).contains(_1269_v14)) ? ((_1272_v15).get(_1269_v14)) : (((_1250_v0) ? (_1250_v0) : (!(_1250_v0)))))) {
              let _1273_v16;
              _1273_v16 = _module.D6.create_DC16(_dafny.Seq.of(_1255_v2, _1255_v2, _1255_v2));
              _1273_v16 = _1273_v16;
              let _1274_v17;
              _1274_v17 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1255_v2), _1255_v2);
              let _index230 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length));
              (_1251_v1)[_index230] = (_1255_v2).isLessThan((((_1274_v17).contains(_1255_v2)) ? ((_1274_v17).get(_1255_v2)) : (_1255_v2)));
              let _index231 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length));
              (_1251_v1)[_index231] = (((_1256_v4).contains(_module.__default.safeModuloInt(_1255_v2, new BigNumber((_dafny.Seq.UnicodeFromString("gmlde")).length)))) ? ((_1256_v4).get(_module.__default.safeModuloInt(_1255_v2, new BigNumber((_dafny.Seq.UnicodeFromString("gmlde")).length)))) : (_1250_v0));
              _1250_v0 = (_1251_v1)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length))];
              let _1275_v18;
              let _nw180 = new _module.C1();
              _nw180.__ctor();
              _1275_v18 = _nw180;
              let _1276_v19;
              _1276_v19 = _module.D15.create_DC38(_1256_v4);
              let _index232 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_1265_v10).length));
              (_1265_v10)[_index232] = (_dafny.ZERO).minus(new BigNumber(((_1276_v19).dtor_cf55).length));
              let _index233 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length));
              let _index234 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_1265_v10).length));
              let _rhs285 = (_1251_v1)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length))];
              let _rhs286 = _1255_v2;
              let _rhs287 = _module.__default.safeDivisionInt(_1255_v2, _1255_v2);
              let _rhs288 = (_1251_v1)[_module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length))];
              let _lhs209 = _1251_v1;
              let _lhs210 = _module.__default.safeIndex(new BigNumber(841), new BigNumber((_1251_v1).length));
              let _lhs211 = globalState;
              let _lhs212 = _1265_v10;
              let _lhs213 = _module.__default.safeIndex(new BigNumber(794), new BigNumber((_1265_v10).length));
              _lhs209[_lhs210] = _rhs285;
              _lhs211.f7 = _rhs286;
              _lhs212[_lhs213] = _rhs287;
              _1250_v0 = _rhs288;
            } else {
              _1250_v0 = _1250_v0;
              let _1277_v20;
              let _nw181 = new _module.C1();
              _nw181.__ctor();
              _1277_v20 = _nw181;
              let _1278_v21;
              _1278_v21 = _dafny.Seq.of(_1277_v20, _1277_v20, _1277_v20, _1277_v20, _1277_v20);
              _1277_v20 = (_1278_v21)[_module.__default.safeIndex(_1255_v2, new BigNumber((_1278_v21).length))];
              _1255_v2 = new BigNumber(328);
              let _1279_v22;
              _1279_v22 = _dafny.Seq.UnicodeFromString("lqpqi");
              let _1280_v23;
              _1280_v23 = _dafny.Seq.of((_1255_v2).plus(_1255_v2));
              let _1281_v24;
              let _nw182 = Array((new BigNumber(12)).toNumber()).fill([]);
              _1281_v24 = _nw182;
              let _index235 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1281_v24).length));
              (_1281_v24)[_index235] = _1251_v1;
              let _1282_v25;
              _1282_v25 = _dafny.Seq.of(_1250_v0);
              let _index236 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1281_v24).length));
              let _rhs289 = (_1277_v20).fm25(_1250_v0, _1250_v0, globalState);
              let _rhs290 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), function (_1283_i5) {
                return (_module.D3.create_DC9(new _dafny.CodePoint('d'.codePointAt(0)))).dtor_cf12;
              });
              let _rhs291 = _dafny.Seq.Concat(_1280_v23, _1280_v23);
              let _rhs292 = _module.__default.safeDivisionInt(_1255_v2, _module.__default.safeModuloInt(_1255_v2, new BigNumber((_1282_v25).length)));
              let _rhs293 = _1251_v1;
              let _lhs214 = globalState;
              let _lhs215 = _1281_v24;
              let _lhs216 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_1281_v24).length));
              _lhs214.f4 = _rhs289;
              _1279_v22 = _rhs290;
              _1280_v23 = _rhs291;
              _1255_v2 = _rhs292;
              _lhs215[_lhs216] = _rhs293;
              let _1284_v26;
              _1284_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v0,_1255_v2);
              _1250_v0 = ((!(_1284_v26).equals(_dafny.Map.Empty.slice().updateUnsafe(_1250_v0,_1255_v2))) ? (false) : (!(_1255_v2).isEqualTo((_dafny.ZERO).minus(_1255_v2))));
            }
            let _1285_v27;
            _1285_v27 = _dafny.MultiSet.fromElements(_1250_v0);
            let _1286_v28;
            _1286_v28 = _module.D2.create_DC7(_1250_v0, _1285_v27, _1265_v10);
            let _1287_v29;
            _1287_v29 = _module.D2.create_DC8(_1286_v28);
            let _1288_v30;
            let _nw183 = new _module.C3();
            _nw183.__ctor(_1287_v29, _1250_v0);
            _1288_v30 = _nw183;
          }
        }
      }
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1289_v0;
      _1289_v0 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).fm2(p0, globalState));
      let _1290_v1;
      _1290_v1 = _dafny.Seq.UnicodeFromString("bqid");
      let _1291_v2;
      _1291_v2 = _module.D4.create_DC11(_dafny.Seq.of(_1290_v1, _1290_v1, _1290_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(321)), function (_1292_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), _1290_v1));
      let _1293_v3;
      _1293_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v0,((p2) ? (_1291_v2) : (_1291_v2)));
      let _1294_v4;
      _1294_v4 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), function (_1295_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-780)), function (_1296_i2) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _1290_v1, _dafny.Seq.UnicodeFromString("qfjnmala"));
      let _1297_v5;
      _1297_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p2);
      let _1298_v6;
      _1298_v6 = _dafny.Seq.of((_1294_v4)[_module.__default.safeIndex(new BigNumber((_1297_v5).length), new BigNumber((_1294_v4).length))], _1290_v1);
      _1293_v3 = (_1293_v3).update(_1289_v0, _module.D4.create_DC11(_1294_v4));
      let _1299_i3;
      _1299_i3 = _dafny.ZERO;
      L9: {
        while (p2) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1299_i3)) {
              break L9;
            }
            _1299_i3 = (_1299_i3).plus(_dafny.ONE);
            let _1300_v7;
            let _nw184 = Array((_dafny.ONE).toNumber()).fill(false);
            _1300_v7 = _nw184;
            let _1301_v8;
            _1301_v8 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,p2));
            let _1302_v9;
            _1302_v9 = _dafny.Seq.of((_module.D16.create_DC41(_1301_v8)).dtor_cf60);
            let _1303_v10;
            _1303_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1300_v7,_dafny.Seq.update((_1302_v9)[_module.__default.safeIndex(p0, new BigNumber((_1302_v9).length))], _module.__default.safeIndex(p0, new BigNumber(((_1302_v9)[_module.__default.safeIndex(p0, new BigNumber((_1302_v9).length))]).length)), _dafny.Map.Empty.slice().updateUnsafe(p2,p2)));
            let _1304_v11;
            _1304_v11 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),p2);
            _1303_v10 = (_1303_v10).update(_1300_v7, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,p2), _1304_v11, _1304_v11));
            let _1305_v12;
            _1305_v12 = true;
            _1305_v12 = (_1305_v12) === (_1305_v12);
            let _1306_v13;
            _1306_v13 = _dafny.Seq.of(p2);
            let _1307_v14;
            let _nw185 = new _module.C0();
            _nw185.__ctor();
            _1307_v14 = _nw185;
            let _1308_v15;
            _1308_v15 = _dafny.Set.fromElements(_1307_v14);
            let _1309_v16;
            _1309_v16 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1290_v1);
            let _1310_v17;
            let _nw186 = Array((new BigNumber(24)).toNumber());
            _nw186[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_1298_v6)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_1298_v6).length))])).length), new BigNumber((_1306_v13).length));
            _nw186[(_dafny.ONE).toNumber()] = new BigNumber(961);
            _nw186[(new BigNumber(2)).toNumber()] = p0;
            _nw186[(new BigNumber(3)).toNumber()] = (new BigNumber(-500)).multipliedBy(p0);
            _nw186[(new BigNumber(4)).toNumber()] = (_this).fm2(new BigNumber(-428), globalState);
            _nw186[(new BigNumber(5)).toNumber()] = p0;
            _nw186[(new BigNumber(6)).toNumber()] = (_this).fm2(p0, globalState);
            _nw186[(new BigNumber(7)).toNumber()] = (p0).minus(p0);
            _nw186[(new BigNumber(8)).toNumber()] = new BigNumber((_1308_v15).length);
            _nw186[(new BigNumber(9)).toNumber()] = p0;
            _nw186[(new BigNumber(10)).toNumber()] = new BigNumber(-819);
            _nw186[(new BigNumber(11)).toNumber()] = p0;
            _nw186[(new BigNumber(12)).toNumber()] = new BigNumber((p1).cardinality());
            _nw186[(new BigNumber(13)).toNumber()] = p0;
            _nw186[(new BigNumber(14)).toNumber()] = new BigNumber((((_1305_v12) ? (_1290_v1) : (_1290_v1))).length);
            _nw186[(new BigNumber(15)).toNumber()] = new BigNumber(((((_1309_v16).contains(_1305_v12)) ? ((_1309_v16).get(_1305_v12)) : (_1290_v1))).length);
            _nw186[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((p1).Intersect(p1)).cardinality()));
            _nw186[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1306_v13, _module.__default.fm27(globalState))).length);
            _nw186[(new BigNumber(18)).toNumber()] = p0;
            _nw186[(new BigNumber(19)).toNumber()] = (p0).multipliedBy(p0);
            _nw186[(new BigNumber(20)).toNumber()] = p0;
            _nw186[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(p0);
            _nw186[(new BigNumber(22)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
            _nw186[(new BigNumber(23)).toNumber()] = p0;
            _1310_v17 = _nw186;
            let _1311_v18;
            _1311_v18 = _module.D8.create_DC25(p0);
            let _pat_let_tv20 = p0;
            let _index237 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1310_v17).length));
            (_1310_v17)[_index237] = (function (_pat_let18_0) {
              return function (_1312_dt__update__tmp_h0) {
                return function (_pat_let19_0) {
                  return function (_1313_dt__update_hcf42_h0) {
                    return _module.D8.create_DC25(_1313_dt__update_hcf42_h0);
                  }(_pat_let19_0);
                }(_pat_let_tv20);
              }(_pat_let18_0);
            }(_1311_v18)).dtor_cf42;
            let _1314_v19;
            let _nw187 = new _module.C5();
            _nw187.__ctor();
            _1314_v19 = _nw187;
            let _index238 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1310_v17).length));
            let _rhs294 = new BigNumber(215);
            let _rhs295 = _1314_v19;
            let _lhs217 = _1310_v17;
            let _lhs218 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((_1310_v17).length));
            _lhs217[_lhs218] = _rhs294;
            _1314_v19 = _rhs295;
            (globalState).f15 = (_1310_v17)[_module.__default.safeIndex(new BigNumber(885), new BigNumber((_1310_v17).length))];
          }
        }
      }
      let _hi9 = p0;
      for (let _1315_i4 = p0; _1315_i4.isLessThan(_hi9); _1315_i4 = _1315_i4.plus(_dafny.ONE)) {
        let _1316_v20;
        _1316_v20 = false;
        let _1317_v21;
        let _nw188 = Array((new BigNumber(15)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = p2;
        _nw188[(_dafny.ONE).toNumber()] = p2;
        _nw188[(new BigNumber(2)).toNumber()] = p2;
        _nw188[(new BigNumber(3)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(4)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(5)).toNumber()] = !(_1316_v20);
        _nw188[(new BigNumber(6)).toNumber()] = _module.__default.fm26(globalState);
        _nw188[(new BigNumber(7)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(8)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(9)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(10)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(11)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(12)).toNumber()] = !(_1316_v20);
        _nw188[(new BigNumber(13)).toNumber()] = _1316_v20;
        _nw188[(new BigNumber(14)).toNumber()] = false;
        _1317_v21 = _nw188;
        let _1318_v22;
        _1318_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1316_v20,_1317_v21);
        let _1319_v23;
        _1319_v23 = _dafny.Set.fromElements(p2);
        let _1320_v24;
        _1320_v24 = _module.D12.create_DC32(_1319_v23);
        let _1321_v25;
        _1321_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_1318_v22).contains(_1316_v20)) ? ((_1318_v22).get(_1316_v20)) : (_1317_v21)),_1320_v24);
        _1316_v20 = !(_1321_v25).equals(_1321_v25);
        let _1322_v26;
        _1322_v26 = _module.D2.create_DC6(p0);
        let _1323_v27;
        let _nw189 = new _module.C3();
        _nw189.__ctor(_module.D2.create_DC8(_1322_v26), p2);
        _1323_v27 = _nw189;
        let _1324_v28;
        _1324_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1290_v1).length),p0);
        let _1325_v29;
        _1325_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1324_v28).length),_1315_i4);
        let _1326_v30;
        _1326_v30 = _dafny.Seq.of((_dafny.ZERO).minus(_1315_i4), _module.__default.fm17(new BigNumber((_1324_v28).length), globalState));
        let _index239 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1317_v21).length));
        (_1317_v21)[_index239] = (_this).fm8(new BigNumber((_1326_v30).length), !(_1323_v27.f25), _1315_i4, globalState);
        let _index240 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1317_v21).length));
        (_1317_v21)[_index240] = !(false);
        let _1327_v31;
        let _init34 = ((_1328_v27) => function (_1329_i5) {
          return _dafny.Map.Empty.slice().updateUnsafe(!(_1328_v27.f25),new BigNumber(851));
        })(_1323_v27);
        let _nw190 = Array((new BigNumber(2)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw190.length); _i0_34++) {
          _nw190[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1327_v31 = _nw190;
        let _index241 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1327_v31).length));
        (_1327_v31)[_index241] = (_1289_v0);
        let _1330_v32;
        _1330_v32 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1331_v33;
        _1331_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1330_v32);
        let _1332_v34;
        let _nw191 = Array((new BigNumber(14)).toNumber());
        _nw191[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw191[(_dafny.ONE).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(2)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(3)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
        _nw191[(new BigNumber(4)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(5)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(6)).toNumber()] = (((_1331_v33).contains(_1315_i4)) ? ((_1331_v33).get(_1315_i4)) : (_1330_v32));
        _nw191[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw191[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
        _nw191[(new BigNumber(9)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(10)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(11)).toNumber()] = ((!(p2)) ? (_1330_v32) : (_1330_v32));
        _nw191[(new BigNumber(12)).toNumber()] = _1330_v32;
        _nw191[(new BigNumber(13)).toNumber()] = _1330_v32;
        _1332_v34 = _nw191;
        let _index242 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1332_v34).length));
        (_1332_v34)[_index242] = _1330_v32;
        let _1333_v35;
        let _nw192 = new _module.C5();
        _nw192.__ctor();
        _1333_v35 = _nw192;
        let _1334_v36;
        _1334_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1315_i4);
        let _1335_v37;
        _1335_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1334_v36).length));
        let _1336_v38;
        _1336_v38 = _dafny.Seq.of(p1, p1, (p1).update(p2, _module.__default.abs(p0)));
        let _index243 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1327_v31).length));
        let _index244 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1332_v34).length));
        let _rhs296 = _dafny.Seq.contains(_1290_v1, _1330_v32);
        let _rhs297 = (_1334_v36).update(!((_1317_v21)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_1317_v21).length))]) || (true), p0);
        let _rhs298 = (_1290_v1)[_module.__default.safeIndex(new BigNumber(((_1336_v38)[_module.__default.safeIndex(new BigNumber((_1319_v23).length), new BigNumber((_1336_v38).length))]).cardinality()), new BigNumber((_1290_v1).length))];
        let _rhs299 = (_1317_v21)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_1317_v21).length))];
        let _rhs300 = _1333_v35;
        let _lhs219 = _1327_v31;
        let _lhs220 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1327_v31).length));
        let _lhs221 = _1332_v34;
        let _lhs222 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_1332_v34).length));
        _1316_v20 = _rhs296;
        _lhs219[_lhs220] = _rhs297;
        _lhs221[_lhs222] = _rhs298;
        _1316_v20 = _rhs299;
        _1333_v35 = _rhs300;
      }
      let _1337_v39;
      _1337_v39 = _dafny.Seq.of(p2);
      let _1338_v40;
      _1338_v40 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1337_v39)[_module.__default.safeIndex(new BigNumber(-696), new BigNumber((_1337_v39).length))]);
      let _1339_v41;
      _1339_v41 = _dafny.MultiSet.fromElements(_module.__default.fm38(false, globalState));
      let _1340_v42;
      _1340_v42 = _dafny.Set.fromElements(new BigNumber((_1339_v41).cardinality()));
      _1338_v40 = (_1338_v40).update(p2, (_1340_v42).IsDisjointFrom(_1340_v42));
      let _1341_v43;
      let _1342_v44;
      let _out11;
      let _out12;
      let _outcollector4 = (_this).m8(globalState);
      _out11 = _outcollector4[0];
      _out12 = _outcollector4[1];
      _1341_v43 = _out11;
      _1342_v44 = _out12;
      let _1343_v45;
      _1343_v45 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1344_v46;
      _1344_v46 = _module.D15.create_DC38(_1297_v5);
      let _pat_let_tv21 = p0;
      let _pat_let_tv22 = p0;
      let _rhs301 = _dafny.Seq.Concat(_1290_v1, _dafny.Seq.update(_1290_v1, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_1342_v44, !(_1342_v44))).cardinality()), new BigNumber((_1290_v1).length)), _1343_v45));
      let _rhs302 = (_dafny.ZERO).minus(p0);
      let _rhs303 = !(!(_1342_v44)) || ((p0).isEqualTo(p0));
      let _rhs304 = function (_source20) {
        if (_source20.is_DC39) {
          let _1345___mcc_h0 = (_source20).cf56;
          let _1346___mcc_h1 = (_source20).cf57;
          let _1347___mcc_h2 = (_source20).cf58;
          let _1348_cf58 = _1347___mcc_h2;
          let _1349_cf57 = _1346___mcc_h1;
          let _1350_cf56 = _1345___mcc_h0;
          return _pat_let_tv21;
        } else if (_source20.is_DC38) {
          let _1351___mcc_h3 = (_source20).cf55;
          let _1352_cf55 = _1351___mcc_h3;
          return _pat_let_tv22;
        } else {
          let _1353___mcc_h4 = (_source20).cf59;
          let _1354_cf59 = _1353___mcc_h4;
          return (new BigNumber(626)).multipliedBy(new BigNumber(818));
        }
      }(_1344_v46);
      let _rhs305 = p0;
      let _lhs223 = globalState;
      let _lhs224 = globalState;
      _1290_v1 = _rhs301;
      _lhs223.f15 = _rhs302;
      _1342_v44 = _rhs303;
      _lhs224.f4 = _rhs304;
      _1341_v43 = _rhs305;
      return;
    }
    m8(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1355_v0;
      _1355_v0 = new BigNumber(919);
      (globalState).f15 = _1355_v0;
      let _1356_v1;
      _1356_v1 = false;
      let _1357_v2;
      _1357_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1356_v1,_1355_v0);
      let _1358_v3;
      let _nw193 = Array((new BigNumber(11)).toNumber());
      _nw193[(_dafny.ZERO).toNumber()] = new BigNumber(-660);
      _nw193[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(_1355_v0)).plus(_1355_v0);
      _nw193[(new BigNumber(2)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(3)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(4)).toNumber()] = new BigNumber(541);
      _nw193[(new BigNumber(5)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(6)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Set.fromElements(_1357_v2)).length);
      _nw193[(new BigNumber(8)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(9)).toNumber()] = _1355_v0;
      _nw193[(new BigNumber(10)).toNumber()] = (_1355_v0).plus(_1355_v0);
      _1358_v3 = _nw193;
      let _1359_v4;
      _1359_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1356_v1,_1356_v1);
      _1358_v3 = ((!(_1359_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_1356_v1))) ? (_1358_v3) : (_1358_v3));
      let _1360_v5;
      _1360_v5 = _module.D6.create_DC17(_1356_v1, _1356_v1, _1356_v1);
      let _pat_let_tv23 = _1355_v0;
      let _pat_let_tv24 = _1356_v1;
      let _pat_let_tv25 = _1356_v1;
      let _pat_let_tv26 = _1355_v0;
      (globalState).f13 = function (_source21) {
        if (_source21.is_DC17) {
          let _1361___mcc_h0 = (_source21).cf24;
          let _1362___mcc_h1 = (_source21).cf25;
          let _1363___mcc_h2 = (_source21).cf26;
          let _1364_cf26 = _1363___mcc_h2;
          let _1365_cf25 = _1362___mcc_h1;
          let _1366_cf24 = _1361___mcc_h0;
          return _pat_let_tv23;
        } else if (_source21.is_DC16) {
          let _1367___mcc_h3 = (_source21).cf23;
          let _1368_cf23 = _1367___mcc_h3;
          return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv24), _dafny.Seq.of(_pat_let_tv25))).length);
        } else {
          let _1369___mcc_h4 = (_source21).cf27;
          let _1370_cf27 = _1369___mcc_h4;
          return _pat_let_tv26;
        }
      }(_1360_v5);
      let _rhs306 = _1356_v1;
      let _rhs307 = _module.__default.fm17(_1355_v0, globalState);
      let _lhs225 = globalState;
      r1 = _rhs306;
      _lhs225.f7 = _rhs307;
      let _1371_v6;
      let _nw194 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _1371_v6 = _nw194;
      let _1372_v7;
      _1372_v7 = new _dafny.CodePoint('k'.codePointAt(0));
      let _index245 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_1371_v6).length));
      (_1371_v6)[_index245] = _1372_v7;
      let _index246 = _module.__default.safeIndex(new BigNumber(972), new BigNumber((_1371_v6).length));
      (_1371_v6)[_index246] = _1372_v7;
      let _1373_v8;
      _1373_v8 = _dafny.MultiSet.fromElements(_1356_v1);
      let _1374_v9;
      let _nw195 = Array((new BigNumber(8)).toNumber());
      _nw195[(_dafny.ZERO).toNumber()] = _1373_v8;
      _nw195[(_dafny.ONE).toNumber()] = (_1373_v8).Intersect(_dafny.MultiSet.fromElements(_1356_v1));
      _nw195[(new BigNumber(2)).toNumber()] = _1373_v8;
      _nw195[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.fromElements(_1356_v1)).Union(_1373_v8);
      _nw195[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_1356_v1);
      _nw195[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.fromElements(_1356_v1)).Union(_1373_v8);
      _nw195[(new BigNumber(6)).toNumber()] = _1373_v8;
      _nw195[(new BigNumber(7)).toNumber()] = (_1373_v8).Difference(_1373_v8);
      _1374_v9 = _nw195;
      let _1375_v10;
      _1375_v10 = _dafny.Seq.of(_1356_v1);
      let _1376_v11;
      _1376_v11 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1375_v10), _dafny.MultiSet.FromArray(_1375_v10));
      let _1377_v12;
      _1377_v12 = _dafny.Set.fromElements(_1356_v1, _1356_v1);
      let _1378_v13;
      _1378_v13 = _dafny.Seq.UnicodeFromString("apmhaajn");
      let _1379_v14;
      _1379_v14 = _dafny.Map.Empty.slice().updateUnsafe((((_1357_v2).contains(_1356_v1)) ? ((_1357_v2).get(_1356_v1)) : (new BigNumber((_1378_v13).length))),_1373_v8);
      let _nw196 = Array((new BigNumber(12)).toNumber());
      _nw196[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_1375_v10);
      _nw196[(_dafny.ONE).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(2)).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(3)).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(4)).toNumber()] = ((_1376_v11)[_module.__default.safeIndex(_1355_v0, new BigNumber((_1376_v11).length))]).Difference(_1373_v8);
      _nw196[(new BigNumber(5)).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(6)).toNumber()] = (_1373_v8).update(_1356_v1, _module.__default.abs(new BigNumber((_1377_v12).length)));
      _nw196[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(_1375_v10);
      _nw196[(new BigNumber(8)).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(9)).toNumber()] = _1373_v8;
      _nw196[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.FromArray(_module.__default.fm27(globalState));
      _nw196[(new BigNumber(11)).toNumber()] = (_1373_v8).Difference((((_1379_v14).contains(_1355_v0)) ? ((_1379_v14).get(_1355_v0)) : (_1373_v8)));
      _1374_v9 = _nw196;
      r0 = _1355_v0;
      let _1380_v15;
      _1380_v15 = _module.D3.create_DC10(_1355_v0, _1355_v0);
      let _pat_let_tv27 = _1356_v1;
      let _pat_let_tv28 = _1356_v1;
      let _pat_let_tv29 = _1356_v1;
      r1 = function (_source22) {
        if (_source22.is_DC10) {
          let _1381___mcc_h5 = (_source22).cf13;
          let _1382___mcc_h6 = (_source22).cf14;
          let _1383_cf14 = _1382___mcc_h6;
          let _1384_cf13 = _1381___mcc_h5;
          return !(_pat_let_tv27) || (_pat_let_tv28);
        } else {
          let _1385___mcc_h7 = (_source22).cf12;
          let _1386_cf12 = _1385___mcc_h7;
          return _pat_let_tv29;
        }
      }(_1380_v15);
      return [r0, r1];
    }
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return (((!(!(true))) ? (_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-16)), function (_1387_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), false)) : (_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("r"), true)))).dtor_cf4;
    };
    fm2(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),false)).length), new BigNumber(330));
    };
    fm7(p0, globalState) {
      let _this = this;
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(-569)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-381))).length)),true)).length), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(6)), function (_1388_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length))).length),new BigNumber(850))).length), new BigNumber(49), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), function (_1389_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("otolbkm")).length))))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())))))).Difference((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber(833)), _dafny.Set.fromElements(new BigNumber(705)), _dafny.Set.fromElements(new BigNumber(909)))).Difference(_dafny.MultiSet.fromElements(function () {
        let _coll62 = new _dafny.Set();
        for (const _compr_62 of (_dafny.Set.fromElements(new BigNumber(829))).Elements) {
          let _1390_v0 = _compr_62;
          if ((_dafny.Set.fromElements(new BigNumber(829))).contains(_1390_v0)) {
            _coll62.add(_module.__default.safeDivisionInt(_1390_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(98)), function (_1391_i2) {
              return new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality());
            })).length), new BigNumber((_dafny.Set.fromElements(!(false))).length))).cardinality())));
          }
        }
        return _coll62;
      }())));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return !((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(390))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(20)))).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qqyrsao")).length),new BigNumber(481)));
    };
    fm9(p0, globalState) {
      let _this = this;
      let _source23 = _module.D2.create_DC6(new BigNumber(-528));
      if (_source23.is_DC6) {
        let _1392___mcc_h0 = (_source23).cf7;
        let _1393_cf7 = _1392___mcc_h0;
        return (_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_1393_cf7)).length), _1393_cf7, _1393_cf7, new BigNumber(338)))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll63 = new _dafny.Map();
          for (const _compr_63 of _dafny.IntegerRange(new BigNumber(484), new BigNumber(492))) {
            let _1394_v0 = _compr_63;
            if (((new BigNumber(484)).isLessThanOrEqualTo(_1394_v0)) && ((_1394_v0).isLessThan(new BigNumber(492)))) {
              _coll63.push([_module.__default.safeDivisionInt(_1394_v0, _1393_cf7),!(false)]);
            }
          }
          return _coll63;
        }()).length)));
      } else if (_source23.is_DC7) {
        let _1395___mcc_h1 = (_source23).cf8;
        let _1396___mcc_h2 = (_source23).cf9;
        let _1397___mcc_h3 = (_source23).cf10;
        let _1398_cf10 = _1397___mcc_h3;
        let _1399_cf9 = _1396___mcc_h2;
        let _1400_cf8 = _1395___mcc_h1;
        return _1400_cf8;
      } else if (_source23.is_DC5) {
        let _1401___mcc_h4 = (_source23).cf6;
        let _1402_cf6 = _1401___mcc_h4;
        return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-84)));
      } else {
        let _1403___mcc_h5 = (_source23).cf11;
        let _1404_cf11 = _1403___mcc_h5;
        return !(!(!(!(true))));
      }
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _1405_v0;
      _1405_v0 = new BigNumber(-428);
      let _1406_v1;
      _1406_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v0,_1405_v0);
      let _1407_v3;
      _1407_v3 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(78));
      let _1408_v4;
      _1408_v4 = _dafny.Seq.of(new BigNumber((_1407_v3).length), _1405_v0);
      let _1409_v10;
      _1409_v10 = false;
      let _1410_v11;
      let _nw197 = Array((new BigNumber(23)).toNumber());
      _nw197[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(_1405_v0, globalState),_1405_v0);
      _nw197[(_dafny.ONE).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(2)).toNumber()] = function () {
        let _coll64 = new _dafny.Map();
        for (const _compr_64 of (_1408_v4).Elements) {
          let _1411_v2 = _compr_64;
          if (_dafny.Seq.contains(_1408_v4, _1411_v2)) {
            _coll64.push([(_1411_v2).plus(_1405_v0),_1405_v0]);
          }
        }
        return _coll64;
      }();
      _nw197[(new BigNumber(3)).toNumber()] = function () {
        let _coll65 = new _dafny.Map();
        for (const _compr_65 of _dafny.IntegerRange(new BigNumber(42), new BigNumber(698))) {
          let _1412_v5 = _compr_65;
          if (((new BigNumber(42)).isLessThanOrEqualTo(_1412_v5)) && ((_1412_v5).isLessThan(new BigNumber(698)))) {
            _coll65.push([_module.__default.safeDivisionInt(_1412_v5, _1405_v0),_1405_v0]);
          }
        }
        return _coll65;
      }();
      _nw197[(new BigNumber(4)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(5)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(6)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(7)).toNumber()] = (_1406_v1).update(_1405_v0, _1405_v0);
      _nw197[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1408_v4)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1408_v4).length))],new BigNumber(-607));
      _nw197[(new BigNumber(9)).toNumber()] = function () {
        let _coll66 = new _dafny.Map();
        for (const _compr_66 of _dafny.IntegerRange(new BigNumber(326), new BigNumber(-776))) {
          let _1413_v6 = _compr_66;
          if (((new BigNumber(326)).isLessThanOrEqualTo(_1413_v6)) && ((_1413_v6).isLessThan(new BigNumber(-776)))) {
            _coll66.push([_module.__default.safeDivisionInt(_1413_v6, _1405_v0),new BigNumber((_1407_v3).length)]);
          }
        }
        return _coll66;
      }();
      _nw197[(new BigNumber(10)).toNumber()] = (_1406_v1).Merge(function () {
        let _coll67 = new _dafny.Map();
        for (const _compr_67 of (_dafny.Seq.of(_1405_v0, _1405_v0)).Elements) {
          let _1414_v7 = _compr_67;
          if (_dafny.Seq.contains(_dafny.Seq.of(_1405_v0, _1405_v0), _1414_v7)) {
            _coll67.push([_module.__default.safeDivisionInt(_1414_v7, _1405_v0),_1405_v0]);
          }
        }
        return _coll67;
      }());
      _nw197[(new BigNumber(11)).toNumber()] = (_1406_v1).Merge(_1406_v1);
      _nw197[(new BigNumber(12)).toNumber()] = (_1406_v1).Merge(_1406_v1);
      _nw197[(new BigNumber(13)).toNumber()] = function () {
        let _coll68 = new _dafny.Map();
        for (const _compr_68 of _dafny.IntegerRange(new BigNumber(-666), new BigNumber(400))) {
          let _1415_v8 = _compr_68;
          if (((new BigNumber(-666)).isLessThanOrEqualTo(_1415_v8)) && ((_1415_v8).isLessThan(new BigNumber(400)))) {
            _coll68.push([_module.__default.safeModuloInt(_1415_v8, new BigNumber(365)),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, true)).cardinality())))]);
          }
        }
        return _coll68;
      }();
      _nw197[(new BigNumber(14)).toNumber()] = (_1406_v1).Merge(_1406_v1);
      _nw197[(new BigNumber(15)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(16)).toNumber()] = (function () {
        let _coll69 = new _dafny.Map();
        for (const _compr_69 of (_dafny.Map.Empty.slice().updateUnsafe(_1405_v0,_1409_v10)).Keys.Elements) {
          let _1416_v9 = _compr_69;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_1405_v0,_1409_v10)).contains(_1416_v9)) {
            _coll69.push([(_1416_v9).multipliedBy(new BigNumber(215)),_1405_v0]);
          }
        }
        return _coll69;
      }()).update(_1405_v0, _1405_v0);
      _nw197[(new BigNumber(17)).toNumber()] = (_1406_v1).Merge(_1406_v1);
      _nw197[(new BigNumber(18)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(19)).toNumber()] = ((_1409_v10) ? (_1406_v1) : (_1406_v1));
      _nw197[(new BigNumber(20)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(21)).toNumber()] = _1406_v1;
      _nw197[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-308),_1405_v0);
      _1410_v11 = _nw197;
      let _index247 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1410_v11).length));
      (_1410_v11)[_index247] = (function () {
        let _coll70 = new _dafny.Map();
        for (const _compr_70 of _dafny.IntegerRange(new BigNumber(537), new BigNumber(936))) {
          let _1417_v12 = _compr_70;
          if (((new BigNumber(537)).isLessThanOrEqualTo(_1417_v12)) && ((_1417_v12).isLessThan(new BigNumber(936)))) {
            _coll70.push([(_1417_v12).multipliedBy(new BigNumber((_1408_v4).length)),_1405_v0]);
          }
        }
        return _coll70;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1405_v0,new BigNumber((p0).length)));
      let _index248 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_1410_v11).length));
      (_1410_v11)[_index248] = (_1406_v1).Merge(((!(_1409_v10)) ? (_1406_v1) : (_1406_v1)));
      _1406_v1 = (_1406_v1).update((_1408_v4)[_module.__default.safeIndex(_1405_v0, new BigNumber((_1408_v4).length))], _1405_v0);
      let _1418_v13;
      _1418_v13 = _dafny.Set.fromElements(_1405_v0, _1405_v0);
      let _hi10 = new BigNumber((_1418_v13).length);
      for (let _1419_i0 = _1405_v0; _1419_i0.isLessThan(_hi10); _1419_i0 = _1419_i0.plus(_dafny.ONE)) {
        let _1420_v14;
        _1420_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_1409_v10);
        if ((((_1420_v14).contains(_1405_v0)) ? ((_1420_v14).get(_1405_v0)) : (_1409_v10))) {
          _1409_v10 = _1409_v10;
          let _1421_v15;
          _1421_v15 = _dafny.MultiSet.fromElements(_1409_v10);
          let _1422_v16;
          let _nw198 = new _module.C8();
          _nw198.__ctor();
          _1422_v16 = _nw198;
          let _1423_v17;
          _1423_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1422_v16,_1409_v10);
          let _1424_v18;
          _1424_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1409_v10,_1423_v17);
          let _1425_v19;
          _1425_v19 = _dafny.Set.fromElements(_1409_v10);
          let _1426_v20;
          let _nw199 = Array((new BigNumber(10)).toNumber());
          _nw199[(_dafny.ZERO).toNumber()] = new BigNumber((_1420_v14).length);
          _nw199[(_dafny.ONE).toNumber()] = _1419_i0;
          _nw199[(new BigNumber(2)).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1409_v10,!(_1409_v10))).update(_1409_v10, _1409_v10)).length);
          _nw199[(new BigNumber(3)).toNumber()] = _1405_v0;
          _nw199[(new BigNumber(4)).toNumber()] = new BigNumber((_1424_v18).length);
          _nw199[(new BigNumber(5)).toNumber()] = new BigNumber(829);
          _nw199[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm32(_1419_i0, _1419_i0, _1425_v19, globalState)).length);
          _nw199[(new BigNumber(7)).toNumber()] = _1405_v0;
          _nw199[(new BigNumber(8)).toNumber()] = _1405_v0;
          _nw199[(new BigNumber(9)).toNumber()] = _1419_i0;
          _1426_v20 = _nw199;
          let _1427_v21;
          _1427_v21 = _module.D2.create_DC7(false, _1421_v15, _1426_v20);
          let _1428_v22;
          _1428_v22 = _dafny.Seq.of(_1409_v10, _1409_v10, _1409_v10, true, _1409_v10);
          let _1429_v23;
          let _nw200 = Array((new BigNumber(22)).toNumber());
          _nw200[(_dafny.ZERO).toNumber()] = _1427_v21;
          _nw200[(_dafny.ONE).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(2)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(3)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(4)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(5)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(6)).toNumber()] = ((_1409_v10) ? (_1427_v21) : (_1427_v21));
          _nw200[(new BigNumber(7)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(8)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(9)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(10)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(11)).toNumber()] = _module.D2.create_DC7(_1409_v10, _dafny.MultiSet.FromArray(_dafny.Seq.update(_1428_v22, _module.__default.safeIndex(_1405_v0, new BigNumber((_1428_v22).length)), _1409_v10)), _1426_v20);
          _nw200[(new BigNumber(12)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(13)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(14)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(15)).toNumber()] = _module.D2.create_DC7(_1409_v10, _1421_v15, _1426_v20);
          _nw200[(new BigNumber(16)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(17)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(18)).toNumber()] = _module.D2.create_DC7(true, _dafny.MultiSet.fromElements(_1409_v10), _1426_v20);
          _nw200[(new BigNumber(19)).toNumber()] = _module.D2.create_DC7(_1409_v10, _1421_v15, _1426_v20);
          _nw200[(new BigNumber(20)).toNumber()] = _1427_v21;
          _nw200[(new BigNumber(21)).toNumber()] = _module.D2.create_DC7(_1409_v10, _module.__default.fm34(new BigNumber((_module.__default.fm19(_1409_v10, !(_1409_v10), globalState)).length), _1405_v0, (((_1420_v14).contains(_1419_i0)) ? ((_1420_v14).get(_1419_i0)) : (_1409_v10)), globalState), _1426_v20);
          _1429_v23 = _nw200;
          let _nw201 = Array((new BigNumber(25)).toNumber()).fill(_module.D2.Default());
          _1429_v23 = _nw201;
          _1409_v10 = (_1425_v19).IsSubsetOf(_1425_v19);
          _1409_v10 = !_dafny.areEqual(_dafny.Seq.Concat(_1408_v4, _1408_v4), _dafny.Seq.Create(_module.__default.abs(new BigNumber(305)), ((_1430_v10, _1431_v3, _1432_i0) => function (_1433_i1) {
            return (((_1431_v3).contains(_1430_v10)) ? ((_1431_v3).get(_1430_v10)) : (_1432_i0));
          })(_1409_v10, _1407_v3, _1419_i0)));
          let _1434_v24;
          _1434_v24 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),_1409_v10);
          let _1435_v25;
          _1435_v25 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1436_v26;
          let _nw202 = Array((new BigNumber(17)).toNumber());
          _nw202[(_dafny.ZERO).toNumber()] = _1409_v10;
          _nw202[(_dafny.ONE).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(2)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(3)).toNumber()] = false;
          _nw202[(new BigNumber(4)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(5)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(6)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(7)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(8)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(9)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(10)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(11)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(12)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(13)).toNumber()] = !(_1409_v10);
          _nw202[(new BigNumber(14)).toNumber()] = _1409_v10;
          _nw202[(new BigNumber(15)).toNumber()] = (((_1434_v24).contains(_1435_v25)) ? ((_1434_v24).get(_1435_v25)) : (false));
          _nw202[(new BigNumber(16)).toNumber()] = _1409_v10;
          _1436_v26 = _nw202;
          let _1437_v27;
          _1437_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1409_v10,_1436_v26);
          _1437_v27 = (_1437_v27).update(_1409_v10, ((true) ? (_1436_v26) : (_1436_v26)));
        } else {
          let _1438_v28;
          let _nw203 = new _module.C2();
          _nw203.__ctor();
          _1438_v28 = _nw203;
          _1438_v28 = _1438_v28;
          (globalState).f4 = _1405_v0;
          let _1439_v29;
          _1439_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1409_v10,p0);
          _1439_v29 = _1439_v29;
          let _1440_v30;
          let _nw204 = new _module.C5();
          _nw204.__ctor();
          _1440_v30 = _nw204;
          _1409_v10 = _1409_v10;
        }
        let _1441_v31;
        _1441_v31 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1442_v32;
        let _nw205 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1442_v32 = _nw205;
        let _1443_v33;
        _1443_v33 = _dafny.Seq.of(_1409_v10, !(false) || (false), _1409_v10);
        let _1444_v34;
        _1444_v34 = _module.D5.create_DC14(_1441_v31, _1409_v10, _1405_v0, _1409_v10);
        let _rhs308 = (_1409_v10) === (_1409_v10);
        let _rhs309 = !((_1443_v33)[_module.__default.safeIndex(_1419_i0, new BigNumber((_1443_v33).length))]);
        let _rhs310 = (_1444_v34).dtor_cf18;
        let _rhs311 = _1442_v32;
        let _rhs312 = (_this).fm2(_module.__default.safeDivisionInt(_1405_v0, _1419_i0), globalState);
        let _lhs226 = globalState;
        _1409_v10 = _rhs308;
        _1409_v10 = _rhs309;
        _1441_v31 = _rhs310;
        _1442_v32 = _rhs311;
        _lhs226.f13 = _rhs312;
        let _1445_v35;
        _1445_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1441_v31),_1442_v32);
        let _1446_v36;
        _1446_v36 = _dafny.MultiSet.fromElements(_1441_v31, _1441_v31, _1441_v31, _1441_v31);
        _1445_v35 = (_1445_v35).update(_1446_v36, _1442_v32);
        let _1447_v37;
        _1447_v37 = _dafny.MultiSet.fromElements(_1405_v0, new BigNumber((_1407_v3).length), _1419_i0);
        let _1448_v38;
        let _nw206 = new _module.C6();
        _nw206.__ctor();
        _1448_v38 = _nw206;
        let _1449_v39;
        _1449_v39 = _dafny.MultiSet.fromElements(_1448_v38, _1448_v38, _1448_v38);
        let _1450_v40;
        let _nw207 = Array((new BigNumber(17)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = new BigNumber((p0).length);
        _nw207[(_dafny.ONE).toNumber()] = new BigNumber(264);
        _nw207[(new BigNumber(2)).toNumber()] = _1419_i0;
        _nw207[(new BigNumber(3)).toNumber()] = new BigNumber((_1447_v37).cardinality());
        _nw207[(new BigNumber(4)).toNumber()] = (_1419_i0).minus((_dafny.ZERO).minus(_1405_v0));
        _nw207[(new BigNumber(5)).toNumber()] = new BigNumber(496);
        _nw207[(new BigNumber(6)).toNumber()] = _1419_i0;
        _nw207[(new BigNumber(7)).toNumber()] = (_this).fm2(_1405_v0, globalState);
        _nw207[(new BigNumber(8)).toNumber()] = _1419_i0;
        _nw207[(new BigNumber(9)).toNumber()] = _1405_v0;
        _nw207[(new BigNumber(10)).toNumber()] = _1405_v0;
        _nw207[(new BigNumber(11)).toNumber()] = _1405_v0;
        _nw207[(new BigNumber(12)).toNumber()] = _1405_v0;
        _nw207[(new BigNumber(13)).toNumber()] = (new BigNumber(210)).plus(new BigNumber((_1447_v37).cardinality()));
        _nw207[(new BigNumber(14)).toNumber()] = _1405_v0;
        _nw207[(new BigNumber(15)).toNumber()] = ((_1409_v10) ? (new BigNumber((_1449_v39).cardinality())) : (_1405_v0));
        _nw207[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(-570), _1405_v0);
        _1450_v40 = _nw207;
        let _index249 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1450_v40).length));
        (_1450_v40)[_index249] = _1405_v0;
        let _index250 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1450_v40).length));
        (_1450_v40)[_index250] = _1405_v0;
      }
      let _1451_v41;
      let _nw208 = Array((new BigNumber(11)).toNumber()).fill(false);
      _1451_v41 = _nw208;
      let _index251 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length));
      (_1451_v41)[_index251] = false;
      let _index252 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length));
      (_1451_v41)[_index252] = _1409_v10;
      let _1452_v42;
      _1452_v42 = _module.D7.create_DC21(_1405_v0, _1405_v0, _1409_v10, (_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))], _1409_v10);
      if (!(((_1452_v42).dtor_cf32) === ((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]))) {
        let _1453_v43;
        let _init35 = ((_1454_v41) => function (_1455_i2) {
          return _dafny.MultiSet.fromElements(!(!(true)), (_1454_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1454_v41).length))]);
        })(_1451_v41);
        let _nw209 = Array((new BigNumber(7)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw209.length); _i0_35++) {
          _nw209[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1453_v43 = _nw209;
        let _index253 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1453_v43).length));
        (_1453_v43)[_index253] = _dafny.MultiSet.fromElements(true);
        let _1456_v44;
        _1456_v44 = _dafny.Set.fromElements((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]);
        let _1457_v45;
        _1457_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1456_v44,(_this).fm9(_1405_v0, globalState));
        let _1458_v46;
        _1458_v46 = _dafny.MultiSet.fromElements((((_1457_v45).contains(_1456_v44)) ? ((_1457_v45).get(_1456_v44)) : ((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))])));
        let _index254 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_1453_v43).length));
        (_1453_v43)[_index254] = _1458_v46;
        let _1459_v47;
        let _nw210 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _1459_v47 = _nw210;
        let _index255 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1459_v47).length));
        (_1459_v47)[_index255] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-664)), function (_1460_i3) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        })).length);
        let _index256 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1459_v47).length));
        (_1459_v47)[_index256] = _1405_v0;
        _1459_v47 = _1459_v47;
        let _1461_v48;
        _1461_v48 = _dafny.Seq.of(_1409_v10);
        let _index257 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length));
        let _rhs313 = true;
        let _rhs314 = _1461_v48;
        let _rhs315 = ((_1456_v44).Difference(_1456_v44)).Difference(_1456_v44);
        let _rhs316 = (_1459_v47)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1459_v47).length))];
        let _lhs227 = _1451_v41;
        let _lhs228 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length));
        let _lhs229 = globalState;
        _lhs227[_lhs228] = _rhs313;
        _1461_v48 = _rhs314;
        _1456_v44 = _rhs315;
        _lhs229.f13 = _rhs316;
        let _1462_v49;
        _1462_v49 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1463_v50;
        _1463_v50 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(_1405_v0, new BigNumber((p0).length)), _1462_v49), p0), _module.__default.fm19(_1409_v10, _1409_v10, globalState), _dafny.Seq.Concat(p0, p0), _dafny.Seq.update(p0, _module.__default.safeIndex((_1459_v47)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1459_v47).length))], new BigNumber((p0).length)), _1462_v49));
        _1463_v50 = _1463_v50;
      } else {
        (globalState).f7 = _1405_v0;
        let _index258 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length));
        (_1451_v41)[_index258] = ((_1409_v10) ? ((_1405_v0).isLessThan(_1405_v0)) : (_1409_v10));
        let _1464_v51;
        let _nw211 = new _module.C8();
        _nw211.__ctor();
        _1464_v51 = _nw211;
        let _1465_v52;
        let _init36 = ((_1466_v0) => function (_1467_i4) {
          return (_1467_i4).minus(_1466_v0);
        })(_1405_v0);
        let _nw212 = Array((new BigNumber(28)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw212.length); _i0_36++) {
          _nw212[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1465_v52 = _nw212;
        let _index259 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length));
        (_1465_v52)[_index259] = _1405_v0;
        let _index260 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length));
        (_1465_v52)[_index260] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.fm17((_this).fm2(_1405_v0, globalState), globalState)).multipliedBy(_module.__default.safeModuloInt(_1405_v0, _1405_v0))));
        if (((_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))]).isLessThan(_module.__default.safeModuloInt((_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))], new BigNumber((p0).length)))) {
          let _1468_v53;
          _1468_v53 = _module.D6.create_DC17((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))], (_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))], (_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]);
          _1409_v10 = (_this).fm8((_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))], ((_1409_v10) ? ((_1468_v53).dtor_cf26) : (_1409_v10)), (_dafny.ZERO).minus(_1405_v0), globalState);
          let _1469_v54;
          let _nw213 = new _module.C2();
          _nw213.__ctor();
          _1469_v54 = _nw213;
          let _1470_v55;
          _1470_v55 = _dafny.MultiSet.fromElements((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]);
          let _1471_v56;
          _1471_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))],_1409_v10);
          let _1472_v57;
          _1472_v57 = _module.D15.create_DC38(_1471_v56);
          let _1473_v58;
          _1473_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1470_v55).Difference(_1470_v55),_1472_v57);
          let _1474_v59;
          _1474_v59 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1475_v60;
          _1475_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1409_v10,_1474_v59);
          _1473_v58 = (_1473_v58).update(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).fm10(_1475_v60, _1405_v0, (_1410_v11)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_1410_v11).length))], globalState), _1409_v10, (_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))], (_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))])), _1472_v57);
          (globalState).f4 = _1405_v0;
          (globalState).f4 = _1405_v0;
        } else {
          _1408_v4 = _1408_v4;
          let _1476_v61;
          let _nw214 = new _module.C4();
          _nw214.__ctor();
          _1476_v61 = _nw214;
          (globalState).f4 = (_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))];
          _1407_v3 = (_1407_v3).update(_1409_v10, (_1465_v52)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1465_v52).length))]);
          let _1477_v62;
          _1477_v62 = _dafny.MultiSet.fromElements(!((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]));
          let _1478_v63;
          _1478_v63 = _module.D2.create_DC5(_1465_v52);
          (_1464_v51).m6(new BigNumber(-905), (_1477_v62).Union(_1477_v62), true, _1478_v63, globalState);
        }
      }
      let _1479_v64;
      let _nw215 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _1479_v64 = _nw215;
      let _1480_v65;
      _1480_v65 = _module.D2.create_DC7((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))], _dafny.MultiSet.fromElements((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))]), _1479_v64);
      let _1481_v66;
      _1481_v66 = _module.D2.create_DC8(_1480_v65);
      let _1482_v67;
      _1482_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v0,_1409_v10);
      let _1483_v68;
      let _nw216 = new _module.C3();
      _nw216.__ctor(_1481_v66, (((_1482_v67).contains(_1405_v0)) ? ((_1482_v67).get(_1405_v0)) : ((_1451_v41)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1451_v41).length))])));
      _1483_v68 = _nw216;
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1484_v0;
      _1484_v0 = _dafny.Set.fromElements(p0);
      let _1485_v1;
      _1485_v1 = _dafny.Seq.of(_1484_v0);
      let _1486_v2;
      _1486_v2 = _dafny.Seq.UnicodeFromString("fwnxp");
      let _1487_v3;
      _1487_v3 = _module.D5.create_DC13(_dafny.MultiSet.fromElements((_1485_v1)[_module.__default.safeIndex(new BigNumber((_1486_v2).length), new BigNumber((_1485_v1).length))]));
      let _source24 = _1487_v3;
      if (_source24.is_DC14) {
        let _1488___mcc_h0 = (_source24).cf18;
        let _1489___mcc_h1 = (_source24).cf19;
        let _1490___mcc_h2 = (_source24).cf20;
        let _1491___mcc_h3 = (_source24).cf21;
        let _1492_cf21 = _1491___mcc_h3;
        let _1493_cf20 = _1490___mcc_h2;
        let _1494_cf19 = _1489___mcc_h1;
        let _1495_cf18 = _1488___mcc_h0;
        let _1496_v4;
        _1496_v4 = _dafny.MultiSet.fromElements(new BigNumber(435));
        let _1497_v5;
        _1497_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1496_v4);
        let _1498_v6;
        _1498_v6 = _dafny.Seq.of(_module.D1.create_DC4((((_1497_v5).contains(p0)) ? ((_1497_v5).get(p0)) : (_1496_v4))));
        _1498_v6 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-20)), ((_1499_p0, _1500_cf20, _1501_cf19, _1502_cf21) => function (_1503_i0) {
          return _module.D1.create_DC4(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_1499_p0, _1500_cf20, _1500_cf20, new BigNumber(-24), new BigNumber((_dafny.Set.fromElements(_1501_cf19, _1502_cf21, !(_1502_cf21))).length))).length), _1499_p0, _1499_p0));
        })(p0, _1493_cf20, _1494_cf19, _1492_cf21));
        let _1504_v7;
        _1504_v7 = _module.D11.create_DC30(p2, _module.__default.fm17((_dafny.ZERO).minus(p0), globalState));
        (globalState).f7 = (_1504_v7).dtor_cf47;
        let _1505_v8;
        let _nw217 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1505_v8 = _nw217;
        let _1506_v9;
        _1506_v9 = _dafny.Seq.of(_1505_v8, _1505_v8, _1505_v8);
        let _1507_v10;
        _1507_v10 = _dafny.Map.Empty.slice().updateUnsafe(false,_1505_v8);
        _1506_v9 = _dafny.Seq.of(_1505_v8, _1505_v8, _1505_v8, (((_1507_v10).contains(_1494_cf19)) ? ((_1507_v10).get(_1494_cf19)) : (_1505_v8)), _1505_v8);
        if (!(!(_1494_cf19) || ((p1).IsProperSubsetOf(p1)))) {
          _1505_v8 = _1505_v8;
          let _1508_v12;
          _1508_v12 = _module.D11.create_DC29();
          let _1509_v13;
          _1509_v13 = _dafny.Set.fromElements(_1508_v12, _1508_v12);
          let _1510_v14;
          _1510_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1494_cf19);
          let _1511_v15;
          _1511_v15 = _module.D12.create_DC32(_dafny.Set.fromElements(p2));
          let _1512_v16;
          _1512_v16 = _dafny.Seq.of(_1511_v15, _1511_v15);
          let _1513_v17;
          _1513_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v13,(_1510_v14).update(new BigNumber((_1512_v16).length), _1492_cf21));
          (globalState).f13 = new BigNumber((function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of (_1513_v17).Keys.Elements) {
              let _1514_v11 = _compr_71;
              if ((_1513_v17).contains(_1514_v11)) {
                _coll71.push([_1514_v11,false]);
              }
            }
            return _coll71;
          }()).length);
          let _1515_v18;
          _1515_v18 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1494_cf19);
          let _1516_v19;
          _1516_v19 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1515_v18);
          let _1517_v20;
          _1517_v20 = _dafny.Seq.of(p2, _1492_cf21, _1492_cf21, ((_1492_cf21) ? (_1492_cf21) : ((_this).fm1(_1516_v19, new BigNumber((_1486_v2).length), globalState))));
          let _1518_v21;
          _1518_v21 = _module.D16.create_DC42(_1494_cf19, _1492_cf21, _1494_cf19);
          let _1519_v22;
          let _nw218 = Array((new BigNumber(12)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = _1518_v21;
          _nw218[(_dafny.ONE).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(2)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(3)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(4)).toNumber()] = _module.D16.create_DC42(p2, (_1518_v21).dtor_cf61, _1494_cf19);
          _nw218[(new BigNumber(5)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(6)).toNumber()] = _module.__default.fm46(p2, globalState);
          _nw218[(new BigNumber(7)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(8)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(9)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(10)).toNumber()] = _1518_v21;
          _nw218[(new BigNumber(11)).toNumber()] = _1518_v21;
          _1519_v22 = _nw218;
          let _index261 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_1519_v22).length));
          (_1519_v22)[_index261] = _1518_v21;
          let _index262 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_1519_v22).length));
          let _rhs317 = _dafny.Seq.Concat(_1517_v20, _dafny.Seq.update(_1517_v20, _module.__default.safeIndex(p0, new BigNumber((_1517_v20).length)), !(_1492_cf21)));
          let _rhs318 = p0;
          let _rhs319 = _1518_v21;
          let _lhs230 = globalState;
          let _lhs231 = _1519_v22;
          let _lhs232 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_1519_v22).length));
          _1517_v20 = _rhs317;
          _lhs230.f13 = _rhs318;
          _lhs231[_lhs232] = _rhs319;
          let _1520_v23;
          let _1521_v24;
          let _out13;
          let _out14;
          let _outcollector5 = (_this).m7(_module.__default.safeModuloInt(_1493_cf20, new BigNumber(821)), _1494_cf19, globalState);
          _out13 = _outcollector5[0];
          _out14 = _outcollector5[1];
          _1520_v23 = _out13;
          _1521_v24 = _out14;
          let _index263 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1505_v8).length));
          (_1505_v8)[_index263] = (_this).fm8(_1493_cf20, _1492_cf21, p0, globalState);
          let _index264 = _module.__default.safeIndex(new BigNumber(900), new BigNumber((_1505_v8).length));
          (_1505_v8)[_index264] = !(false);
        } else {
          let _index265 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1505_v8).length));
          (_1505_v8)[_index265] = _1492_cf21;
          let _1522_v25;
          _1522_v25 = _dafny.Seq.of((_dafny.ZERO).minus(_1493_cf20));
          let _1523_v26;
          _1523_v26 = _dafny.MultiSet.fromElements(_1486_v2, _dafny.Seq.update(_1486_v2, _module.__default.safeIndex(new BigNumber((_1522_v25).length), new BigNumber((_1486_v2).length)), _1495_cf18), _1486_v2);
          let _index266 = _module.__default.safeIndex(new BigNumber(818), new BigNumber((_1505_v8).length));
          (_1505_v8)[_index266] = (_module.__default.fm47(p2, (_dafny.ZERO).minus(_1493_cf20), globalState)).IsProperSubsetOf(_1523_v26);
          _1486_v2 = _dafny.Seq.Concat(_1486_v2, _dafny.Seq.Concat(_1486_v2, _1486_v2));
          let _1524_v27;
          let _nw219 = new _module.C1();
          _nw219.__ctor();
          _1524_v27 = _nw219;
          let _1525_v28;
          _1525_v28 = _dafny.Seq.of(_1524_v27, _1524_v27);
          let _1526_v29;
          _1526_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(133),_1525_v28);
          let _1527_v30;
          _1527_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(771),_1526_v29);
          let _1528_v31;
          _1528_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(_1494_cf19),_1493_cf20);
          let _1529_v32;
          let _nw220 = Array((new BigNumber(23)).toNumber());
          _nw220[(_dafny.ZERO).toNumber()] = (_1526_v29).Merge((_1526_v29).update(_1493_cf20, _1525_v28));
          _nw220[(_dafny.ONE).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(2)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(3)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(4)).toNumber()] = (_1526_v29).update(_1493_cf20, _1525_v28);
          _nw220[(new BigNumber(5)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(6)).toNumber()] = ((_1492_cf21) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1486_v2).length),_1525_v28)) : (_1526_v29));
          _nw220[(new BigNumber(7)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(8)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(9)).toNumber()] = (_1526_v29).Merge(_1526_v29);
          _nw220[(new BigNumber(10)).toNumber()] = (_1526_v29).Merge((((_1527_v30).contains((_1522_v25)[_module.__default.safeIndex(_1493_cf20, new BigNumber((_1522_v25).length))])) ? ((_1527_v30).get((_1522_v25)[_module.__default.safeIndex(_1493_cf20, new BigNumber((_1522_v25).length))])) : (_dafny.Map.Empty.slice().updateUnsafe(_1493_cf20,_dafny.Seq.of(_1524_v27)))));
          _nw220[(new BigNumber(11)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(12)).toNumber()] = (_1526_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1525_v28));
          _nw220[(new BigNumber(13)).toNumber()] = (_1526_v29).Merge(_1526_v29);
          _nw220[(new BigNumber(14)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1525_v28)).update((_dafny.ZERO).minus(_1493_cf20), _1525_v28);
          _nw220[(new BigNumber(16)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(17)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(18)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(19)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_1525_v28)).update(new BigNumber(((_1528_v31).update(_1494_cf19, p0)).length), _1525_v28);
          _nw220[(new BigNumber(20)).toNumber()] = _1526_v29;
          _nw220[(new BigNumber(21)).toNumber()] = (_1526_v29).Merge(_1526_v29);
          _nw220[(new BigNumber(22)).toNumber()] = (_1526_v29).Merge(_1526_v29);
          _1529_v32 = _nw220;
          let _index267 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1529_v32).length));
          (_1529_v32)[_index267] = _1526_v29;
          let _index268 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_1529_v32).length));
          (_1529_v32)[_index268] = _1526_v29;
          (globalState).f4 = p0;
          let _1530_v33;
          _1530_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1492_cf21,_1494_cf19);
          let _1531_v34;
          _1531_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1530_v33).length),_1493_cf20);
          _1531_v34 = (_1531_v34).update((_1493_cf20).plus((_dafny.ZERO).minus(p0)), (new BigNumber(241)).minus(_1493_cf20));
        }
      } else if (_source24.is_DC13) {
        let _1532___mcc_h4 = (_source24).cf17;
        let _1533_cf17 = _1532___mcc_h4;
        let _1534_v35;
        _1534_v35 = false;
        _1534_v35 = p2;
        let _1535_v36;
        _1535_v36 = _dafny.Seq.of(_1486_v2, ((p2) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), function (_1536_i1) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })) : (_1486_v2)), _dafny.Seq.Concat(_1486_v2, _1486_v2), _1486_v2, _1486_v2);
        let _1537_v37;
        let _nw221 = new _module.C2();
        _nw221.__ctor();
        _1537_v37 = _nw221;
        let _1538_v38;
        _1538_v38 = _dafny.Set.fromElements(_1537_v37, _1537_v37, _1537_v37, _1537_v37, _1537_v37);
        let _1539_v39;
        _1539_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1538_v38).length),p0);
        let _1540_v40;
        _1540_v40 = _dafny.Seq.of(p0, p0, new BigNumber((_1539_v39).length), new BigNumber(269));
        let _1541_v41;
        _1541_v41 = _dafny.Seq.of(new BigNumber((_1484_v0).length), (_dafny.ZERO).minus((_1540_v40)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("is")).length), new BigNumber((_1540_v40).length))]), (_dafny.ZERO).minus(p0));
        let _1542_v42;
        _1542_v42 = _dafny.Seq.of(_1540_v40);
        _1486_v2 = (_1535_v36)[_module.__default.safeIndex(new BigNumber((_1542_v42).length), new BigNumber((_1535_v36).length))];
        let _1543_v43;
        let _init37 = ((_1544_v35) => function (_1545_i2) {
          return _1544_v35;
        })(_1534_v35);
        let _nw222 = Array((new BigNumber(2)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw222.length); _i0_37++) {
          _nw222[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1543_v43 = _nw222;
        let _index269 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1543_v43).length));
        (_1543_v43)[_index269] = p2;
        let _index270 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_1543_v43).length));
        (_1543_v43)[_index270] = !(p2);
        let _1546_v44;
        let _nw223 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1546_v44 = _nw223;
        let _1547_v45;
        let _init38 = function (_1548_i3) {
          return (_1548_i3).plus(new BigNumber(-411));
        };
        let _nw224 = Array((new BigNumber(9)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw224.length); _i0_38++) {
          _nw224[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1547_v45 = _nw224;
        let _index271 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1546_v44).length));
        (_1546_v44)[_index271] = _1547_v45;
        let _index272 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1546_v44).length));
        let _rhs320 = _1547_v45;
        let _rhs321 = (_1543_v43)[_module.__default.safeIndex(new BigNumber(747), new BigNumber((_1543_v43).length))];
        let _rhs322 = _1539_v39;
        let _rhs323 = _module.__default.safeModuloInt(new BigNumber((_1486_v2).length), p0);
        let _lhs233 = _1546_v44;
        let _lhs234 = _module.__default.safeIndex(new BigNumber(169), new BigNumber((_1546_v44).length));
        let _lhs235 = globalState;
        _lhs233[_lhs234] = _rhs320;
        _1534_v35 = _rhs321;
        _1539_v39 = _rhs322;
        _lhs235.f7 = _rhs323;
      } else {
        let _1549___mcc_h5 = (_source24).cf22;
        let _1550_cf22 = _1549___mcc_h5;
        if (((((p1).contains(true)) ? ((p1).get(true)) : (p0))).isLessThan(p0)) {
          let _1551_v46;
          _1551_v46 = _dafny.MultiSet.fromElements(new BigNumber(793), p0, p0);
          let _1552_v47;
          _1552_v47 = _dafny.Seq.of((_1551_v46).contains(p0), p2, (_this).fm9(p0, globalState));
          _1552_v47 = _1552_v47;
          (globalState).f13 = p0;
          let _1553_v48;
          let _nw225 = new _module.C6();
          _nw225.__ctor();
          _1553_v48 = _nw225;
          let _1554_v49;
          _1554_v49 = _dafny.Seq.of(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dkgfi"), _1486_v2)).length), p0);
          _1554_v49 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), ((_1555_p0) => function (_1556_i4) {
            return _1555_p0;
          })(p0));
          let _1557_v50;
          _1557_v50 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p2);
          let _1558_v51;
          _1558_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(898),(((_1557_v50).contains((((_1551_v46).contains(p0)) ? ((_1551_v46).get(p0)) : (p0)))) ? ((_1557_v50).get((((_1551_v46).contains(p0)) ? ((_1551_v46).get(p0)) : (p0)))) : ((_1552_v47)[_module.__default.safeIndex(p0, new BigNumber((_1552_v47).length))])));
          let _1559_v52;
          _1559_v52 = _dafny.Seq.of(_1558_v51);
          let _1560_v53;
          _1560_v53 = _module.D14.create_DC35(_1559_v52);
          let _1561_v54;
          _1561_v54 = _dafny.Seq.of(_1560_v53);
          let _1562_v55;
          _1562_v55 = _module.D1.create_DC3(_1486_v2, p2);
          _1561_v54 = _dafny.Seq.Concat(_1561_v54, _dafny.Seq.of(_module.D14.create_DC35(_1559_v52), _module.__default.fm48(!(p2), p2, (_1562_v55).dtor_cf4, globalState)));
        } else {
          let _1563_v56;
          _1563_v56 = false;
          let _1564_v57;
          _1564_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1484_v0);
          _1563_v56 = (_1484_v0).equals((((_1564_v57).contains(p0)) ? ((_1564_v57).get(p0)) : (_1484_v0)));
          _1563_v56 = false;
          (globalState).f7 = (new BigNumber((_1484_v0).length)).multipliedBy(p0);
          let _1565_v58;
          let _nw226 = new _module.C5();
          _nw226.__ctor();
          _1565_v58 = _nw226;
          let _1566_v59;
          _1566_v59 = _dafny.Seq.of(_1563_v56, p2);
          let _1567_v60;
          _1567_v60 = _dafny.Seq.of(p2, (_1566_v59)[_module.__default.safeIndex(new BigNumber((_1566_v59).length), new BigNumber((_1566_v59).length))]);
          let _1568_v61;
          _1568_v61 = _dafny.Set.fromElements(p2);
          let _1569_v62;
          _1569_v62 = _module.D2.create_DC6(_module.__default.safeDivisionInt(p0, (_dafny.ZERO).minus(p0)));
          let _rhs324 = _1486_v2;
          let _rhs325 = _dafny.Seq.Concat(_1566_v59, _dafny.Seq.update(_dafny.Seq.of(_1563_v56, p2), _module.__default.safeIndex(new BigNumber(584), new BigNumber((_dafny.Seq.of(_1563_v56, p2)).length)), p2));
          let _rhs326 = _1568_v61;
          let _rhs327 = _1569_v62;
          _1486_v2 = _rhs324;
          _1566_v59 = _rhs325;
          _1568_v61 = _rhs326;
          _1569_v62 = _rhs327;
        }
        let _1570_v63;
        _1570_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(246));
        let _1571_v64;
        _1571_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1570_v63,p2);
        let _1572_v65;
        _1572_v65 = _module.D17.create_DC45(_1571_v64);
        _1571_v64 = ((!(p2)) ? (_1571_v64) : ((_1571_v64).Merge((_1572_v65).dtor_cf67)));
        let _1573_v66;
        _1573_v66 = true;
        let _1574_v67;
        _1574_v67 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1575_v68;
        _1575_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1574_v67,_1573_v66);
        let _1576_v69;
        let _nw227 = Array((new BigNumber(8)).toNumber());
        _nw227[(_dafny.ZERO).toNumber()] = p2;
        _nw227[(_dafny.ONE).toNumber()] = false;
        _nw227[(new BigNumber(2)).toNumber()] = (_this).fm9(p0, globalState);
        _nw227[(new BigNumber(3)).toNumber()] = true;
        _nw227[(new BigNumber(4)).toNumber()] = _1573_v66;
        _nw227[(new BigNumber(5)).toNumber()] = (((_1575_v68).contains(_1574_v67)) ? ((_1575_v68).get(_1574_v67)) : (p2));
        _nw227[(new BigNumber(6)).toNumber()] = p2;
        _nw227[(new BigNumber(7)).toNumber()] = p2;
        _1576_v69 = _nw227;
        let _1577_v70;
        _1577_v70 = _dafny.Set.fromElements(_1576_v69, _1576_v69, _1576_v69);
        let _1578_v71;
        _1578_v71 = _dafny.Seq.of(_1573_v66, p2, !(_1573_v66));
        let _rhs328 = p2;
        let _rhs329 = (_1578_v71)[_module.__default.safeIndex((p0).minus(new BigNumber((_1570_v63).length)), new BigNumber((_1578_v71).length))];
        let _rhs330 = _1577_v70;
        _1573_v66 = _rhs328;
        _1573_v66 = _rhs329;
        _1577_v70 = _rhs330;
        let _1579_v72;
        _1579_v72 = _module.D14.create_DC37(_module.D14.create_DC36(_1578_v71));
        let _1580_v73;
        _1580_v73 = _module.D14.create_DC36(_1578_v71);
        let _1581_v74;
        _1581_v74 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1573_v66);
        let _pat_let_tv30 = _1580_v73;
        let _pat_let_tv31 = _1580_v73;
        let _1582_v75;
        let _nw228 = Array((new BigNumber(25)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = _1579_v72;
        _nw228[(_dafny.ONE).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(2)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(3)).toNumber()] = function (_pat_let20_0) {
          return function (_1583_dt__update__tmp_h0) {
            return function (_pat_let21_0) {
              return function (_1584_dt__update_hcf54_h0) {
                return _module.D14.create_DC37(_1584_dt__update_hcf54_h0);
              }(_pat_let21_0);
            }(_pat_let_tv30);
          }(_pat_let20_0);
        }(_module.D14.create_DC37(_1580_v73));
        _nw228[(new BigNumber(4)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(5)).toNumber()] = function (_pat_let22_0) {
          return function (_1585_dt__update__tmp_h1) {
            return function (_pat_let23_0) {
              return function (_1586_dt__update_hcf54_h1) {
                return _module.D14.create_DC37(_1586_dt__update_hcf54_h1);
              }(_pat_let23_0);
            }(_pat_let_tv31);
          }(_pat_let22_0);
        }(_1579_v72);
        _nw228[(new BigNumber(6)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(7)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(8)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(9)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(10)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(11)).toNumber()] = _module.D14.create_DC37(_1580_v73);
        _nw228[(new BigNumber(12)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(13)).toNumber()] = _module.D14.create_DC37(_module.__default.fm48(_1573_v66, p2, (((_1581_v74).contains(p0)) ? ((_1581_v74).get(p0)) : (p2)), globalState));
        _nw228[(new BigNumber(14)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(15)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(16)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(17)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(18)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(19)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(20)).toNumber()] = ((p2) ? (_1579_v72) : (_1579_v72));
        _nw228[(new BigNumber(21)).toNumber()] = _module.D14.create_DC37(_1580_v73);
        _nw228[(new BigNumber(22)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(23)).toNumber()] = _1579_v72;
        _nw228[(new BigNumber(24)).toNumber()] = _module.D14.create_DC37(_module.D14.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(969)), ((_1587_v74, _1588_p0, _1589_p2) => function (_1590_i5) {
  return (_1587_v74).update(_1588_p0, _1589_p2);
})(_1581_v74, p0, p2))));
        _1582_v75 = _nw228;
        let _index273 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1582_v75).length));
        (_1582_v75)[_index273] = _1579_v72;
        let _index274 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_1582_v75).length));
        (_1582_v75)[_index274] = _1579_v72;
      }
      if (!(((p0).isLessThanOrEqualTo(p0)) && (!((new BigNumber(807)).isLessThanOrEqualTo(p0))))) {
        _1486_v2 = _dafny.Seq.Concat(_1486_v2, _dafny.Seq.UnicodeFromString("eiseggug"));
        let _1591_v76;
        let _nw229 = new _module.C0();
        _nw229.__ctor();
        _1591_v76 = _nw229;
        let _1592_v77;
        let _init39 = function (_1593_i6) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        };
        let _nw230 = Array((new BigNumber(18)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw230.length); _i0_39++) {
          _nw230[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1592_v77 = _nw230;
        _1592_v77 = _1592_v77;
        if (!((new BigNumber((_1484_v0).length)).isEqualTo(p0))) {
          (globalState).f13 = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), p0);
          let _1594_v78;
          _1594_v78 = _dafny.Seq.of(p2, p2, p2, p2, p2);
          let _1595_v79;
          _1595_v79 = _dafny.Seq.of(p0, new BigNumber((_1594_v78).length));
          let _1596_v80;
          _1596_v80 = _dafny.Seq.of(p0, p0, new BigNumber((_1595_v79).length));
          let _1597_v81;
          _1597_v81 = _dafny.Seq.of(_dafny.Seq.of(p0), _1596_v80);
          let _1598_v82;
          _1598_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1486_v2);
          let _1599_v83;
          _1599_v83 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1600_v84;
          _1600_v84 = _dafny.Set.fromElements(p2);
          let _1601_v85;
          _1601_v85 = _dafny.Seq.of(_1600_v84);
          let _1602_v86;
          _1602_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1601_v85,_1486_v2);
          let _1603_v87;
          _1603_v87 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _1604_v88;
          let _nw231 = new _module.C2();
          _nw231.__ctor();
          _1604_v88 = _nw231;
          let _1605_v89;
          _1605_v89 = _dafny.Seq.of(_module.__default.fm49(p0, new BigNumber((_dafny.Set.fromElements(_1604_v88, _1604_v88)).length), globalState));
          let _1606_v90;
          _1606_v90 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((p1).cardinality()))).update(p2, new BigNumber(-649))).length),_dafny.Set.fromElements(new BigNumber((_1603_v87).length), new BigNumber((_1605_v89).length)));
          let _1607_v91;
          _1607_v91 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1597_v81)[_module.__default.safeIndex(p0, new BigNumber((_1597_v81).length))]);
          let _1608_v92;
          let _nw232 = Array((new BigNumber(27)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = (_1591_v76).fm28(_1597_v81, globalState);
          _nw232[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          _nw232[(new BigNumber(2)).toNumber()] = ((p2) ? (_dafny.Seq.update((((_1598_v82).contains(p2)) ? ((_1598_v82).get(p2)) : (_1486_v2)), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber(((((_1598_v82).contains(p2)) ? ((_1598_v82).get(p2)) : (_1486_v2))).length)), _1599_v83)) : (_1486_v2));
          _nw232[(new BigNumber(3)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(4)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          _nw232[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_1609_v83) => function (_1610_i7) {
            return _1609_v83;
          })(_1599_v83));
          _nw232[(new BigNumber(7)).toNumber()] = (((_1602_v86).contains(_dafny.Seq.of(_dafny.Set.fromElements(p2, p2, p2), _dafny.Set.fromElements(p2, false), _1600_v84, _1600_v84))) ? ((_1602_v86).get(_dafny.Seq.of(_dafny.Set.fromElements(p2, p2, p2), _dafny.Set.fromElements(p2, false), _1600_v84, _1600_v84))) : (_1486_v2));
          _nw232[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("kuauuuwp");
          _nw232[(new BigNumber(9)).toNumber()] = _module.__default.fm30(new BigNumber(535), p2, _dafny.Set.fromElements(new BigNumber((_1606_v90).length), new BigNumber(769)), p2, globalState);
          _nw232[(new BigNumber(10)).toNumber()] = ((p2) ? (_1486_v2) : (_dafny.Seq.UnicodeFromString("eiqjnpdrf")));
          _nw232[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          _nw232[(new BigNumber(12)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1486_v2, (_1591_v76).fm28(_1597_v81, globalState));
          _nw232[(new BigNumber(14)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          _nw232[(new BigNumber(16)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          _nw232[(new BigNumber(18)).toNumber()] = (_1591_v76).fm28(_dafny.Seq.of((((_1607_v91).contains(new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality()))) ? ((_1607_v91).get(new BigNumber((_dafny.MultiSet.fromElements(p2)).cardinality()))) : (_1595_v79)), _1596_v80), globalState);
          _nw232[(new BigNumber(19)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(20)).toNumber()] = (((_1598_v82).contains(p2)) ? ((_1598_v82).get(p2)) : (_1486_v2));
          _nw232[(new BigNumber(21)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(22)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("sxbmntpf");
          _nw232[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ggfo"), _1486_v2);
          _nw232[(new BigNumber(25)).toNumber()] = _1486_v2;
          _nw232[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("kkqchmef");
          _1608_v92 = _nw232;
          let _index275 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1608_v92).length));
          (_1608_v92)[_index275] = _dafny.Seq.Concat(_1486_v2, _1486_v2);
          let _index276 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1608_v92).length));
          (_1608_v92)[_index276] = _1486_v2;
          (globalState).f15 = p0;
          let _1611_v93;
          let _nw233 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _1611_v93 = _nw233;
          let _1612_v94;
          _1612_v94 = _module.D2.create_DC5(_1611_v93);
          _1612_v94 = _module.D2.create_DC5(_1611_v93);
          (globalState).f4 = new BigNumber(493);
        } else {
          let _1613_v95;
          _1613_v95 = true;
          let _1614_v96;
          let _nw234 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1614_v96 = _nw234;
          let _index277 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length));
          (_1614_v96)[_index277] = p0;
          let _1615_v97;
          _1615_v97 = _dafny.Seq.of(p2);
          let _1616_v98;
          _1616_v98 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1613_v95);
          let _1617_v99;
          _1617_v99 = _dafny.Seq.of(new BigNumber(15), new BigNumber(-836));
          let _1618_v100;
          _1618_v100 = _module.D15.create_DC39((_1616_v98).update(_1613_v95, _1613_v95), _1613_v95, new BigNumber((_1617_v99).length));
          let _index278 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length));
          let _rhs331 = (p1).contains(p2);
          let _rhs332 = (p0).plus((new BigNumber((_1616_v98).length)).minus(p0));
          let _rhs333 = _1615_v97;
          let _rhs334 = (_1618_v100).dtor_cf57;
          let _rhs335 = _1613_v95;
          let _lhs236 = _1614_v96;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length));
          _1613_v95 = _rhs331;
          _lhs236[_lhs237] = _rhs332;
          _1615_v97 = _rhs333;
          _1613_v95 = _rhs334;
          _1613_v95 = _rhs335;
          let _1619_v101;
          _1619_v101 = _module.D14.create_DC35(_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_1620_v96, _1621_p2) => function (_1622_i8) {
  return _dafny.Map.Empty.slice().updateUnsafe((_1620_v96)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1620_v96).length))],_1621_p2);
})(_1614_v96, p2)));
          let _1623_v102;
          _1623_v102 = _module.D14.create_DC35((_1619_v101).dtor_cf52);
          let _1624_v103;
          _1624_v103 = _module.D14.create_DC37(_1623_v102);
          _1624_v103 = _1624_v103;
          _1613_v95 = _1613_v95;
          let _1625_v104;
          let _init40 = ((_1626_p1) => function (_1627_i9) {
            return (_1626_p1).IsProperSubsetOf(_1626_p1);
          })(p1);
          let _nw235 = Array((new BigNumber(29)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw235.length); _i0_40++) {
            _nw235[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1625_v104 = _nw235;
          let _index279 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_1625_v104).length));
          (_1625_v104)[_index279] = p2;
          let _1628_v105;
          _1628_v105 = _dafny.MultiSet.fromElements((_1614_v96)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length))], (_1617_v99)[_module.__default.safeIndex((_1614_v96)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length))], new BigNumber((_1617_v99).length))]);
          let _index280 = _module.__default.safeIndex(new BigNumber(47), new BigNumber((_1625_v104).length));
          (_1625_v104)[_index280] = (_1628_v105).IsSubsetOf((_dafny.MultiSet.fromElements((_1614_v96)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length))], p0)).Union(_module.__default.fm31((_1614_v96)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1614_v96).length))], p0, globalState)));
          let _1629_v106;
          _1629_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1617_v99,!((_1625_v104)[_module.__default.safeIndex(new BigNumber(47), new BigNumber((_1625_v104).length))]));
          _1629_v106 = _1629_v106;
        }
        (globalState).f13 = ((_this).fm2(p0, globalState)).minus(p0);
      } else {
        let _1630_v107;
        _1630_v107 = _dafny.Set.fromElements(p2, p2);
        let _1631_v108;
        _1631_v108 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
        let _1632_v109;
        _1632_v109 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        let _1633_v110;
        _1633_v110 = _dafny.Seq.of(!((_this).fm8(new BigNumber((_1630_v107).length), (((_1631_v108).contains(true)) ? ((_1631_v108).get(true)) : (!(!(p2)))), (((_1632_v109).contains(p2)) ? ((_1632_v109).get(p2)) : ((_dafny.ZERO).minus((_this).fm2(p0, globalState)))), globalState)), p2);
        let _1634_v111;
        let _nw236 = Array((new BigNumber(24)).toNumber());
        _nw236[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0));
        _nw236[(_dafny.ONE).toNumber()] = (p0).minus(new BigNumber((_1633_v110).length));
        _nw236[(new BigNumber(2)).toNumber()] = p0;
        _nw236[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw236[(new BigNumber(4)).toNumber()] = p0;
        _nw236[(new BigNumber(5)).toNumber()] = p0;
        _nw236[(new BigNumber(6)).toNumber()] = p0;
        _nw236[(new BigNumber(7)).toNumber()] = new BigNumber(75);
        _nw236[(new BigNumber(8)).toNumber()] = p0;
        _nw236[(new BigNumber(9)).toNumber()] = p0;
        _nw236[(new BigNumber(10)).toNumber()] = p0;
        _nw236[(new BigNumber(11)).toNumber()] = new BigNumber(917);
        _nw236[(new BigNumber(12)).toNumber()] = p0;
        _nw236[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1630_v107).length), p0);
        _nw236[(new BigNumber(14)).toNumber()] = new BigNumber(265);
        _nw236[(new BigNumber(15)).toNumber()] = p0;
        _nw236[(new BigNumber(16)).toNumber()] = p0;
        _nw236[(new BigNumber(17)).toNumber()] = p0;
        _nw236[(new BigNumber(18)).toNumber()] = p0;
        _nw236[(new BigNumber(19)).toNumber()] = p0;
        _nw236[(new BigNumber(20)).toNumber()] = p0;
        _nw236[(new BigNumber(21)).toNumber()] = new BigNumber((p1).cardinality());
        _nw236[(new BigNumber(22)).toNumber()] = (_this).fm2(_module.__default.fm17(new BigNumber((_1486_v2).length), globalState), globalState);
        _nw236[(new BigNumber(23)).toNumber()] = p0;
        _1634_v111 = _nw236;
        _1634_v111 = _1634_v111;
        (globalState).f4 = p0;
        let _1635_v112;
        let _init41 = ((_1636_v110, _1637_p0) => function (_1638_i10) {
          return _module.D12.create_DC32(_dafny.Set.fromElements((_1636_v110)[_module.__default.safeIndex(_1637_p0, new BigNumber((_1636_v110).length))]));
        })(_1633_v110, p0);
        let _nw237 = Array((new BigNumber(28)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw237.length); _i0_41++) {
          _nw237[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1635_v112 = _nw237;
        let _1639_v113;
        _1639_v113 = _module.D12.create_DC32(_dafny.Set.fromElements(p2));
        let _index281 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_1635_v112).length));
        (_1635_v112)[_index281] = _1639_v113;
        let _1640_v114;
        let _nw238 = Array((new BigNumber(15)).toNumber()).fill(_module.D4.Default());
        _1640_v114 = _nw238;
        let _1641_v115;
        _1641_v115 = _module.D4.create_DC11(_dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_1642_v2) => function (_1643_i11) {
  return _1642_v2;
})(_1486_v2)));
        let _index282 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1640_v114).length));
        (_1640_v114)[_index282] = _1641_v115;
        let _index283 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_1635_v112).length));
        let _index284 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1640_v114).length));
        let _rhs336 = ((p2) ? (p0) : (p0));
        let _rhs337 = _1630_v107;
        let _rhs338 = _module.D12.create_DC32(_1630_v107);
        let _rhs339 = p0;
        let _rhs340 = _1641_v115;
        let _lhs238 = globalState;
        let _lhs239 = _1635_v112;
        let _lhs240 = _module.__default.safeIndex(new BigNumber(615), new BigNumber((_1635_v112).length));
        let _lhs241 = globalState;
        let _lhs242 = _1640_v114;
        let _lhs243 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_1640_v114).length));
        _lhs238.f4 = _rhs336;
        _1630_v107 = _rhs337;
        _lhs239[_lhs240] = _rhs338;
        _lhs241.f4 = _rhs339;
        _lhs242[_lhs243] = _rhs340;
        let _1644_v116;
        _1644_v116 = _dafny.Seq.of(p0, p0);
        let _1645_v117;
        _1645_v117 = _dafny.MultiSet.fromElements(p0, new BigNumber((_1644_v116).length), p0, new BigNumber((_1486_v2).length));
        let _1646_v119;
        _1646_v119 = _dafny.MultiSet.fromElements((p0).multipliedBy(p0), new BigNumber(-763), (_dafny.ZERO).minus((((_1645_v117).contains(p0)) ? ((_1645_v117).get(p0)) : ((_dafny.ZERO).minus(p0)))), ((false) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p2),new BigNumber((function () {
          let _coll72 = new _dafny.Map();
          for (const _compr_72 of _dafny.IntegerRange(new BigNumber(436), new BigNumber(908))) {
            let _1647_v118 = _compr_72;
            if (((new BigNumber(436)).isLessThanOrEqualTo(_1647_v118)) && ((_1647_v118).isLessThan(new BigNumber(908)))) {
              _coll72.push([_module.__default.safeDivisionInt(_1647_v118, p0),p2]);
            }
          }
          return _coll72;
        }()).length))).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(815)), function (_1648_i12) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length))), p0);
        _1645_v117 = _1645_v117;
        if (p2) {
          let _1649_v120;
          _1649_v120 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), ((_1650_p0) => function (_1651_i13) {
            return _1650_p0;
          })(p0)));
          let _1652_v121;
          _1652_v121 = _module.D1.create_DC3(_1486_v2, !(!(true)));
          let _1653_v122;
          _1653_v122 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((((_1649_v120).update(_1644_v116, _module.__default.abs(new BigNumber((_1630_v107).length)))).update(_1644_v116, _module.__default.abs(new BigNumber(-51)))).cardinality())),(_1652_v121).dtor_cf3);
          _1653_v122 = (((_1630_v107).IsSubsetOf(_dafny.Set.fromElements(p2))) ? (_dafny.Map.Empty.slice().updateUnsafe(p0,_1486_v2)) : ((_1653_v122).Merge(_1653_v122)));
          let _1654_v123;
          _1654_v123 = true;
          _1654_v123 = (p0).isLessThan(_module.__default.safeDivisionInt(p0, p0));
          let _1655_v124;
          let _nw239 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1655_v124 = _nw239;
          let _1656_v125;
          let _nw240 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1656_v125 = _nw240;
          let _1657_v126;
          _1657_v126 = _dafny.Map.Empty.slice().updateUnsafe(false,_1656_v125);
          let _index285 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1655_v124).length));
          (_1655_v124)[_index285] = (((_1657_v126).contains(true)) ? ((_1657_v126).get(true)) : (_1656_v125));
          let _index286 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1655_v124).length));
          (_1655_v124)[_index286] = _1656_v125;
          let _index287 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1655_v124).length));
          (_1655_v124)[_index287] = _1656_v125;
          let _1658_v127;
          _1658_v127 = _module.D8.create_DC25((p0).minus(p0));
          _1658_v127 = _1658_v127;
        } else {
          let _1659_v128;
          _1659_v128 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1660_v129;
          let _nw241 = Array((new BigNumber(17)).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = _1486_v2;
          _nw241[(_dafny.ONE).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-789)), function (_1661_i14) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          });
          _nw241[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("xxum");
          _nw241[(new BigNumber(4)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(5)).toNumber()] = _module.__default.fm19(p2, !(p2), globalState);
          _nw241[(new BigNumber(6)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(7)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(8)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(9)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("hhvdynhug");
          _nw241[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_1486_v2, _module.__default.safeIndex(_module.__default.fm17((_dafny.ZERO).minus(p0), globalState), new BigNumber((_1486_v2).length)), _1659_v128);
          _nw241[(new BigNumber(12)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_1662_i15) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          });
          _nw241[(new BigNumber(14)).toNumber()] = _1486_v2;
          _nw241[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("m");
          _nw241[(new BigNumber(16)).toNumber()] = _1486_v2;
          _1660_v129 = _nw241;
          let _1663_v130;
          _1663_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1660_v129,(p0).isEqualTo(p0));
          _1663_v130 = (_1663_v130).update(_1660_v129, p2);
          let _1664_v131;
          _1664_v131 = true;
          let _rhs341 = p0;
          let _rhs342 = (_1664_v131) && ((p1).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1664_v131, p2)));
          let _rhs343 = p2;
          let _lhs244 = globalState;
          _lhs244.f4 = _rhs341;
          _1664_v131 = _rhs342;
          _1664_v131 = _rhs343;
          let _1665_v132;
          let _init42 = ((_1666_v131) => function (_1667_i16) {
            return _1666_v131;
          })(_1664_v131);
          let _nw242 = Array((new BigNumber(6)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw242.length); _i0_42++) {
            _nw242[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1665_v132 = _nw242;
          let _1668_v133;
          _1668_v133 = _dafny.Map.Empty.slice().updateUnsafe(!(_1664_v131),_1631_v108);
          let _index288 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1665_v132).length));
          (_1665_v132)[_index288] = (_this).fm1(_1668_v133, new BigNumber((_1633_v110).length), globalState);
          let _index289 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1665_v132).length));
          (_1665_v132)[_index289] = false;
          let _1669_v134;
          let _nw243 = new _module.C7();
          _nw243.__ctor();
          _1669_v134 = _nw243;
          let _1670_v135;
          _1670_v135 = _dafny.Map.Empty.slice().updateUnsafe(_1665_v132,_1665_v132);
          let _1671_v136;
          _1671_v136 = _dafny.Map.Empty.slice().updateUnsafe(_1665_v132,_1670_v135);
          _1670_v135 = (((_1671_v136).contains(_1665_v132)) ? ((_1671_v136).get(_1665_v132)) : (_1670_v135));
        }
      }
      let _1672_v137;
      let _nw244 = Array((new BigNumber(20)).toNumber()).fill(false);
      _1672_v137 = _nw244;
      let _1673_v138;
      _1673_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v137,p2);
      _1673_v138 = (_1673_v138).update(_1672_v137, p2);
      let _hi11 = p0;
      for (let _1674_i17 = _module.__default.safeDivisionInt(p0, p0); _1674_i17.isLessThan(_hi11); _1674_i17 = _1674_i17.plus(_dafny.ONE)) {
        _1672_v137 = _1672_v137;
        let _1675_v139;
        let _nw245 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1675_v139 = _nw245;
        let _1676_v140;
        let _init43 = ((_1677_p0) => function (_1678_i18) {
          return _module.__default.safeModuloInt(_1678_i18, _1677_p0);
        })(p0);
        let _nw246 = Array((new BigNumber(20)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw246.length); _i0_43++) {
          _nw246[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1676_v140 = _nw246;
        let _index290 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1675_v139).length));
        (_1675_v139)[_index290] = _1676_v140;
        let _1679_v141;
        _1679_v141 = _dafny.MultiSet.fromElements(((p2) ? (_1674_i17) : (new BigNumber(-605))), (_1674_i17).multipliedBy(_1674_i17));
        let _1680_v142;
        _1680_v142 = _dafny.Seq.of(_1679_v141, _1679_v141);
        let _1681_v143;
        let _nw247 = new _module.C7();
        _nw247.__ctor();
        _1681_v143 = _nw247;
        let _1682_v144;
        _1682_v144 = _dafny.Seq.of(((_1680_v142)[_module.__default.safeIndex(p0, new BigNumber((_1680_v142).length))]).update(_1674_i17, _module.__default.abs((_module.D17.create_DC47(_1681_v143, true, _1674_i17, _1674_i17, _module.__default.fm17(_1674_i17, globalState))).dtor_cf72)));
        let _index291 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1675_v139).length));
        let _rhs344 = _1676_v140;
        let _rhs345 = new BigNumber((p1).cardinality());
        let _rhs346 = (_1682_v144)[_module.__default.safeIndex(_1674_i17, new BigNumber((_1682_v144).length))];
        let _lhs245 = _1675_v139;
        let _lhs246 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_1675_v139).length));
        let _lhs247 = globalState;
        _lhs245[_lhs246] = _rhs344;
        _lhs247.f4 = _rhs345;
        _1679_v141 = _rhs346;
        if (p2) {
          let _1683_v145;
          _1683_v145 = true;
          let _1684_v146;
          _1684_v146 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1681_v143).fm2(p0, globalState));
          let _1685_v147;
          _1685_v147 = _dafny.Set.fromElements(_1684_v146, _1684_v146);
          _1683_v145 = (_1685_v147).IsSubsetOf(_1685_v147);
          (globalState).f15 = (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-81))).cardinality())).minus(_module.__default.safeModuloInt(new BigNumber((_1679_v141).cardinality()), _1674_i17));
          _1683_v145 = _1683_v145;
          let _1686_v148;
          _1686_v148 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), ((_1687_i17) => function (_1688_i19) {
            return _1687_i17;
          })(_1674_i17)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), ((_1689_p0) => function (_1690_i20) {
            return _1689_p0;
          })(p0)));
          let _1691_v149;
          _1691_v149 = _dafny.Seq.of(p0);
          let _1692_v150;
          _1692_v150 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1693_v151;
          _1693_v151 = _dafny.Set.fromElements((_this).fm9(_1674_i17, globalState));
          let _rhs347 = p2;
          let _rhs348 = _dafny.MultiSet.fromElements(_1691_v149, _module.__default.fm39(_1674_i17, _1692_v150, _1691_v149, _1683_v145, globalState));
          let _rhs349 = !(new BigNumber(751)).isEqualTo(new BigNumber((_1693_v151).length));
          _1683_v145 = _rhs347;
          _1686_v148 = _rhs348;
          _1683_v145 = _rhs349;
          let _1694_v152;
          _1694_v152 = _dafny.Map.Empty.slice().updateUnsafe(_1683_v145,p0);
          let _1695_v153;
          let _nw248 = Array((new BigNumber(13)).toNumber());
          _nw248[(_dafny.ZERO).toNumber()] = _1694_v152;
          _nw248[(_dafny.ONE).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(2)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(3)).toNumber()] = (_1694_v152).update(_1683_v145, _1674_i17);
          _nw248[(new BigNumber(4)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(5)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(6)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(7)).toNumber()] = _module.__default.fm44(_1674_i17, _1683_v145, globalState);
          _nw248[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1683_v145,p0);
          _nw248[(new BigNumber(9)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(10)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(11)).toNumber()] = _1694_v152;
          _nw248[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          _1695_v153 = _nw248;
          let _1696_v154;
          _1696_v154 = _dafny.Map.Empty.slice().updateUnsafe(_1695_v153,p2);
          _1683_v145 = (((_1696_v154).contains(_1695_v153)) ? ((_1696_v154).get(_1695_v153)) : (((p2) ? (false) : (false))));
        } else {
          let _1697_v155;
          _1697_v155 = _module.D1.create_DC4(_1679_v141);
          _1679_v141 = (_1679_v141).Intersect((_1697_v155).dtor_cf5);
          let _1698_v156;
          let _nw249 = new _module.C8();
          _nw249.__ctor();
          _1698_v156 = _nw249;
          let _1699_v157;
          _1699_v157 = _dafny.Set.fromElements(p2, true, p2);
          let _1700_v158;
          _1700_v158 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1699_v157);
          let _1701_v159;
          _1701_v159 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_1702_i21) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }));
          _1700_v158 = (_1700_v158).update(!(_1701_v159).equals(_1701_v159), _1699_v157);
          let _1703_v160;
          _1703_v160 = _module.D1.create_DC2(p0);
          (globalState).f13 = (_module.__default.safeModuloInt(p0, p0)).minus((_1703_v160).dtor_cf2);
          _1699_v157 = _dafny.Set.fromElements(p2, p2, p2, p2, p2);
        }
        let _1704_v161;
        _1704_v161 = new _dafny.CodePoint('u'.codePointAt(0));
        let _1705_v162;
        _1705_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v137,_1704_v161);
        let _1706_v163;
        _1706_v163 = _dafny.Map.Empty.slice().updateUnsafe((p0).minus(p0),(((_1705_v162).contains(_1672_v137)) ? ((_1705_v162).get(_1672_v137)) : (_1704_v161)));
        _1706_v163 = (_1706_v163).update(new BigNumber(115), _1704_v161);
      }
      let _1707_v164;
      let _init44 = function (_1708_i22) {
        return (_1708_i22).plus(new BigNumber(871));
      };
      let _nw250 = Array((new BigNumber(27)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw250.length); _i0_44++) {
        _nw250[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1707_v164 = _nw250;
      _1707_v164 = _1707_v164;
      let _hi12 = p0;
      for (let _1709_i23 = p0; _1709_i23.isLessThan(_hi12); _1709_i23 = _1709_i23.plus(_dafny.ONE)) {
        let _1710_v165;
        let _nw251 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1710_v165 = _nw251;
        let _nw252 = Array((new BigNumber(21)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1710_v165 = _nw252;
        let _1711_v166;
        let _nw253 = new _module.C7();
        _nw253.__ctor();
        _1711_v166 = _nw253;
        _1711_v166 = _1711_v166;
        let _1712_v167;
        _1712_v167 = false;
        _1712_v167 = _1712_v167;
        _1712_v167 = _1712_v167;
      }
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.MultiSet.Empty;
      let _1713_v0;
      let _nw254 = Array((new BigNumber(25)).toNumber()).fill(false);
      _1713_v0 = _nw254;
      let _1714_v1;
      _1714_v1 = _dafny.MultiSet.fromElements(_1713_v0, _1713_v0);
      r0 = (((_1714_v1).contains(_1713_v0)) ? ((_1714_v1).get(_1713_v0)) : (p0));
      let _1715_i0;
      _1715_i0 = _dafny.ZERO;
      L10: {
        while (p1) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1715_i0)) {
              break L10;
            }
            _1715_i0 = (_1715_i0).plus(_dafny.ONE);
            if (true) {
              let _1716_v2;
              let _init45 = function (_1717_i1) {
                return _module.D3.create_DC10(new BigNumber(65), new BigNumber(626));
              };
              let _nw255 = Array((new BigNumber(7)).toNumber());
              for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw255.length); _i0_45++) {
                _nw255[_i0_45] = _init45(new BigNumber(_i0_45));
              }
              _1716_v2 = _nw255;
              let _1718_v3;
              _1718_v3 = _dafny.MultiSet.fromElements(p1);
              let _1719_v4;
              _1719_v4 = _dafny.Seq.of(p1, p1);
              let _1720_v5;
              _1720_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1719_v4).length),p0);
              let _1721_v6;
              _1721_v6 = _dafny.Seq.of(_1720_v5);
              let _1722_v7;
              _1722_v7 = _dafny.Seq.of(new BigNumber(((_1718_v3).update(p1, _module.__default.abs(p0))).cardinality()), (_this).fm2(p0, globalState), new BigNumber((_1721_v6).length));
              let _1723_v8;
              _1723_v8 = _module.D3.create_DC10(new BigNumber((_1722_v7).length), p0);
              let _index292 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1716_v2).length));
              (_1716_v2)[_index292] = _1723_v8;
              let _1724_v9;
              _1724_v9 = _dafny.Set.fromElements(!(p1));
              let _1725_v10;
              _1725_v10 = _dafny.Set.fromElements(_1724_v9, _1724_v9);
              let _pat_let_tv32 = _1720_v5;
              let _pat_let_tv33 = _1725_v10;
              let _pat_let_tv34 = p0;
              let _pat_let_tv35 = _1720_v5;
              let _pat_let_tv36 = _1725_v10;
              let _pat_let_tv37 = p0;
              let _index293 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_1716_v2).length));
              (_1716_v2)[_index293] = function (_pat_let24_0) {
                return function (_1726_dt__update__tmp_h0) {
                  return function (_pat_let25_0) {
                    return function (_1727_dt__update_hcf14_h0) {
                      return _module.D3.create_DC10((_1726_dt__update__tmp_h0).dtor_cf13, _1727_dt__update_hcf14_h0);
                    }(_pat_let25_0);
                  }(_module.__default.safeModuloInt((((_pat_let_tv35).contains(new BigNumber((_pat_let_tv36).length))) ? ((_pat_let_tv32).get(new BigNumber((_pat_let_tv33).length))) : (_pat_let_tv34)), _pat_let_tv37));
                }(_pat_let24_0);
              }(_1723_v8);
              let _1728_v11;
              let _nw256 = new _module.C5();
              _nw256.__ctor();
              _1728_v11 = _nw256;
              (globalState).f15 = _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(((p1) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_1728_v11)).length)) : (new BigNumber((_dafny.Seq.UnicodeFromString("tmqug")).length)))));
              let _1729_v12;
              _1729_v12 = _module.D18.create_DC48(_1720_v5);
              let _1730_v13;
              _1730_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _1731_v14;
              _1731_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1730_v13);
              let _pat_let_tv38 = _1720_v5;
              let _1732_v15;
              _1732_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p0, p0, (_dafny.ZERO).minus(new BigNumber(((function (_pat_let26_0) {
                return function (_1733_dt__update__tmp_h1) {
                  return function (_pat_let27_0) {
                    return function (_1734_dt__update_hcf74_h0) {
                      return _module.D18.create_DC48(_1734_dt__update_hcf74_h0);
                    }(_pat_let27_0);
                  }(_pat_let_tv38);
                }(_pat_let26_0);
              }(_1729_v12)).dtor_cf74).length)))).length),(_this).fm1(_1731_v14, p0, globalState));
              let _1735_v17;
              _1735_v17 = _dafny.Set.fromElements((_dafny.ZERO).minus(p0), p0);
              let _1736_v18;
              _1736_v18 = _dafny.Seq.of(function () {
                let _coll73 = new _dafny.Set();
                for (const _compr_73 of (_1732_v15).Keys.Elements) {
                  let _1737_v16 = _compr_73;
                  if ((_1732_v15).contains(_1737_v16)) {
                    _coll73.add((_1737_v16).minus(new BigNumber(-68)));
                  }
                }
                return _coll73;
              }(), _1735_v17);
              let _1738_v19;
              _1738_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1735_v17,_dafny.Set.fromElements(p0, new BigNumber((_1722_v7).length)));
              let _1739_v21;
              _1739_v21 = _dafny.Seq.of((_1736_v18)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_1736_v18).length))], (((_1738_v19).contains(function () {
                let _coll75 = new _dafny.Set();
                for (const _compr_75 of _dafny.IntegerRange(new BigNumber(594), new BigNumber(392))) {
                  let _1741_v20 = _compr_75;
                  if (((new BigNumber(594)).isLessThanOrEqualTo(_1741_v20)) && ((_1741_v20).isLessThan(new BigNumber(392)))) {
                    _coll75.add((_1741_v20).plus(p0));
                  }
                }
                return _coll75;
              }())) ? ((_1738_v19).get(function () {
                let _coll74 = new _dafny.Set();
                for (const _compr_74 of _dafny.IntegerRange(new BigNumber(594), new BigNumber(392))) {
                  let _1740_v20 = _compr_74;
                  if (((new BigNumber(594)).isLessThanOrEqualTo(_1740_v20)) && ((_1740_v20).isLessThan(new BigNumber(392)))) {
                    _coll74.add((_1740_v20).plus(p0));
                  }
                }
                return _coll74;
              }())) : (_1735_v17)), _dafny.Set.fromElements((_dafny.ZERO).minus(p0), p0, p0), _1735_v17, _1735_v17);
              (globalState).f13 = new BigNumber((_1739_v21).length);
              let _1742_v22;
              let _nw257 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
              _1742_v22 = _nw257;
              let _1743_v23;
              _1743_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v22,p0);
              let _1744_v24;
              _1744_v24 = _dafny.Seq.UnicodeFromString("kxs");
              let _1745_v25;
              _1745_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1744_v24,_1743_v23);
              let _index294 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1713_v0).length));
              (_1713_v0)[_index294] = !(_1743_v23).equals((((_1745_v25).contains(_dafny.Seq.UnicodeFromString("idknhjqof"))) ? ((_1745_v25).get(_dafny.Seq.UnicodeFromString("idknhjqof"))) : (_1743_v23)));
              let _index295 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1713_v0).length));
              (_1713_v0)[_index295] = (_1728_v11).fm20(globalState);
              let _1746_v26;
              _1746_v26 = _module.D2.create_DC6(p0);
              let _1747_v27;
              _1747_v27 = _dafny.MultiSet.fromElements(_1746_v26, _1746_v26, _1746_v26);
              let _1748_v28;
              _1748_v28 = _dafny.MultiSet.fromElements((_1747_v27).update(_1746_v26, _module.__default.abs(p0)), _1747_v27, _1747_v27, _1747_v27, _1747_v27);
              let _1749_v29;
              _1749_v29 = new _dafny.CodePoint('f'.codePointAt(0));
              let _1750_v30;
              _1750_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1713_v0)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1713_v0).length))],_1749_v29);
              let _1751_v31;
              _1751_v31 = _dafny.Seq.of(_1747_v27, (_1747_v27).update(_1746_v26, _module.__default.abs(p0)));
              let _1752_v32;
              _1752_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.FromArray(_1751_v31));
              let _1753_v33;
              let _nw258 = Array((new BigNumber(9)).toNumber());
              _nw258[(_dafny.ZERO).toNumber()] = _1748_v28;
              _nw258[(_dafny.ONE).toNumber()] = ((p1) ? (_1748_v28) : (_1748_v28));
              _nw258[(new BigNumber(2)).toNumber()] = (_1748_v28).Intersect(_1748_v28);
              _nw258[(new BigNumber(3)).toNumber()] = _module.__default.fm50((_1713_v0)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1713_v0).length))], (_this).fm10(_1750_v30, new BigNumber(790), _1720_v5, globalState), p0, (((_1732_v15).contains(p0)) ? ((_1732_v15).get(p0)) : (p1)), globalState);
              _nw258[(new BigNumber(4)).toNumber()] = (((_1752_v32).contains(p0)) ? ((_1752_v32).get(p0)) : (_1748_v28));
              _nw258[(new BigNumber(5)).toNumber()] = _1748_v28;
              _nw258[(new BigNumber(6)).toNumber()] = _1748_v28;
              _nw258[(new BigNumber(7)).toNumber()] = _1748_v28;
              _nw258[(new BigNumber(8)).toNumber()] = (_module.__default.fm50((_1713_v0)[_module.__default.safeIndex(new BigNumber(636), new BigNumber((_1713_v0).length))], (_this).fm9(p0, globalState), p0, p1, globalState)).Union(_1748_v28);
              _1753_v33 = _nw258;
              let _index296 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_1753_v33).length));
              (_1753_v33)[_index296] = _1748_v28;
              let _index297 = _module.__default.safeIndex(new BigNumber(732), new BigNumber((_1753_v33).length));
              (_1753_v33)[_index297] = _1748_v28;
            } else {
              let _1754_v34;
              _1754_v34 = true;
              let _1755_v35;
              _1755_v35 = _dafny.MultiSet.fromElements(p0, new BigNumber(842), (_dafny.ZERO).minus(p0));
              let _1756_v36;
              _1756_v36 = _dafny.Seq.UnicodeFromString("l");
              let _rhs350 = (_dafny.ZERO).minus(p0);
              let _rhs351 = _dafny.Seq.IsPrefixOf(((p1) ? (_1756_v36) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), function (_1757_i2) {
                return new _dafny.CodePoint('b'.codePointAt(0));
              }))), _1756_v36);
              let _rhs352 = ((_1755_v35).Intersect(_1755_v35)).Union(_1755_v35);
              let _lhs248 = globalState;
              _lhs248.f13 = _rhs350;
              _1754_v34 = _rhs351;
              _1755_v35 = _rhs352;
              let _1758_v37;
              _1758_v37 = _dafny.Set.fromElements(p0, p0, p0, p0, p0);
              let _1759_v38;
              _1759_v38 = _dafny.MultiSet.fromElements(_1758_v37);
              let _1760_v39;
              _1760_v39 = _module.D5.create_DC13(_1759_v38);
              let _1761_v40;
              _1761_v40 = _dafny.MultiSet.fromElements(_1760_v39);
              (globalState).f4 = new BigNumber((_1761_v40).cardinality());
              let _index298 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1713_v0).length));
              (_1713_v0)[_index298] = p1;
              let _1762_v41;
              _1762_v41 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),p1);
              let _1763_v42;
              _1763_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1754_v34,_1762_v41);
              let _index299 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1713_v0).length));
              let _rhs353 = _1713_v0;
              let _rhs354 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((p0).multipliedBy((_dafny.ZERO).minus(p0)), _module.__default.fm17(p0, globalState)));
              let _rhs355 = (_this).fm1(_1763_v42, p0, globalState);
              let _lhs249 = globalState;
              let _lhs250 = _1713_v0;
              let _lhs251 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1713_v0).length));
              _1713_v0 = _rhs353;
              _lhs249.f13 = _rhs354;
              _lhs250[_lhs251] = _rhs355;
              _1754_v34 = (_1758_v37).IsDisjointFrom(((false) ? (_1758_v37) : (_dafny.Set.fromElements(p0))));
              let _1764_v43;
              _1764_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(670),_1756_v36);
              let _1765_v44;
              _1765_v44 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1764_v43).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_1756_v36)));
              _1765_v44 = (_1765_v44).update((_1713_v0)[_module.__default.safeIndex(new BigNumber(84), new BigNumber((_1713_v0).length))], function () {
                let _coll76 = new _dafny.Map();
                for (const _compr_76 of (_1755_v35).Elements) {
                  let _1766_v45 = _compr_76;
                  if ((_1755_v35).contains(_1766_v45)) {
                    _coll76.push([(_1766_v45).minus(p0),_dafny.Seq.Create(_module.__default.abs(new BigNumber(124)), function (_1767_i3) {
                      return new _dafny.CodePoint('u'.codePointAt(0));
                    })]);
                  }
                }
                return _coll76;
              }());
            }
            (globalState).f15 = _module.__default.safeModuloInt(p0, p0);
            (globalState).f15 = _module.__default.fm17(p0, globalState);
            let _1768_v46;
            _1768_v46 = true;
            let _1769_v47;
            _1769_v47 = _dafny.Seq.UnicodeFromString("yfrlx");
            let _1770_v48;
            _1770_v48 = _dafny.MultiSet.fromElements(p1);
            let _rhs356 = _1768_v46;
            let _rhs357 = p1;
            let _rhs358 = _1769_v47;
            let _rhs359 = _1770_v48;
            _1768_v46 = _rhs356;
            _1768_v46 = _rhs357;
            _1769_v47 = _rhs358;
            _1770_v48 = _rhs359;
          }
        }
      }
      r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
      let _1771_v49;
      _1771_v49 = _dafny.Set.fromElements(p1, p1);
      let _1772_v50;
      _1772_v50 = _dafny.Set.fromElements(new BigNumber((_1771_v49).length));
      let _1773_v51;
      _1773_v51 = _module.D11.create_DC28(_1772_v50);
      let _1774_v52;
      _1774_v52 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("holgkicqv")).length), p0);
      let _1775_v53;
      _1775_v53 = _dafny.Seq.of(_module.__default.fm51(_1774_v52, p0, globalState));
      let _1776_v54;
      _1776_v54 = _module.D11.create_DC31(_module.D11.create_DC29());
      let _1777_v55;
      _1777_v55 = _module.D19.create_DC51(_1775_v53);
      let _1778_v56;
      _1778_v56 = _dafny.Seq.of(_1775_v53, _1775_v53);
      let _pat_let_tv39 = _1773_v51;
      let _1779_v57;
      let _nw259 = Array((new BigNumber(21)).toNumber());
      _nw259[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_module.D11.create_DC31(_1773_v51));
      _nw259[(_dafny.ONE).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_module.D11.create_DC31(_module.D11.create_DC31(_1773_v51)), _1776_v54, _1776_v54), _1775_v53);
      _nw259[(new BigNumber(3)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_1776_v54);
      _nw259[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(function (_pat_let28_0) {
        return function (_1780_dt__update__tmp_h2) {
          return function (_pat_let29_0) {
            return function (_1781_dt__update_hcf48_h0) {
              return _module.D11.create_DC31(_1781_dt__update_hcf48_h0);
            }(_pat_let29_0);
          }(_pat_let_tv39);
        }(_pat_let28_0);
      }(_1776_v54));
      _nw259[(new BigNumber(6)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(7)).toNumber()] = ((p1) ? (_1775_v53) : ((_1777_v55).dtor_cf80));
      _nw259[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1776_v54), _1775_v53);
      _nw259[(new BigNumber(9)).toNumber()] = (_1778_v56)[_module.__default.safeIndex(p0, new BigNumber((_1778_v56).length))];
      _nw259[(new BigNumber(10)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm52(globalState), _1775_v53);
      _nw259[(new BigNumber(12)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), ((_1782_v54) => function (_1783_i4) {
        return _1782_v54;
      })(_1776_v54));
      _nw259[(new BigNumber(14)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(15)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(16)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-243)), ((_1784_v54) => function (_1785_i5) {
        return _1784_v54;
      })(_1776_v54));
      _nw259[(new BigNumber(18)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(19)).toNumber()] = _1775_v53;
      _nw259[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_1775_v53, _module.__default.safeIndex(p0, new BigNumber((_1775_v53).length)), _module.__default.fm51(_1774_v52, p0, globalState));
      _1779_v57 = _nw259;
      let _index300 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1779_v57).length));
      (_1779_v57)[_index300] = _dafny.Seq.of(_module.D11.create_DC31(_1773_v51), _1776_v54);
      let _1786_v58;
      _1786_v58 = true;
      let _1787_v59;
      _1787_v59 = _module.D0.create_DC0(_1713_v0);
      let _1788_v60;
      _1788_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
      let _1789_v61;
      _1789_v61 = _dafny.Seq.of(_1786_v58, (((_1788_v60).contains(new BigNumber((_dafny.Seq.update(_1774_v52, _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1774_v52).length)), p0)).length))) ? ((_1788_v60).get(new BigNumber((_dafny.Seq.update(_1774_v52, _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1774_v52).length)), p0)).length))) : (_1786_v58)));
      let _1790_v62;
      _1790_v62 = _dafny.Seq.UnicodeFromString("ksdbdq");
      let _1791_v63;
      _1791_v63 = _dafny.Seq.of(_1790_v62, _dafny.Seq.UnicodeFromString("mhqnf"), _module.__default.fm30(p0, _1786_v58, _1772_v50, true, globalState));
      let _1792_v64;
      _1792_v64 = _module.D7.create_DC21(p0, p0, false, _1786_v58, p1);
      let _index301 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1779_v57).length));
      let _rhs360 = (((_1786_v58) ? (_1787_v59) : (_1787_v59))).dtor_cf0;
      let _rhs361 = _dafny.Seq.update(_module.__default.fm52(globalState), _module.__default.safeIndex((new BigNumber((_1789_v61).length)).plus(new BigNumber(-182)), new BigNumber((_module.__default.fm52(globalState)).length)), _1776_v54);
      let _rhs362 = new BigNumber(((_1791_v63)[_module.__default.safeIndex((p0).plus(p0), new BigNumber((_1791_v63).length))]).length);
      let _rhs363 = ((p0).minus(p0)).isLessThan(((_dafny.ZERO).minus((_1792_v64).dtor_cf30)).plus(new BigNumber(-837)));
      let _lhs252 = _1779_v57;
      let _lhs253 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_1779_v57).length));
      let _lhs254 = globalState;
      _1713_v0 = _rhs360;
      _lhs252[_lhs253] = _rhs361;
      _lhs254.f4 = _rhs362;
      _1786_v58 = _rhs363;
      let _1793_i6;
      _1793_i6 = _dafny.ZERO;
      L11: {
        while (p1) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1793_i6)) {
              break L11;
            }
            _1793_i6 = (_1793_i6).plus(_dafny.ONE);
            let _index302 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1713_v0).length));
            (_1713_v0)[_index302] = false;
            let _index303 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1713_v0).length));
            (_1713_v0)[_index303] = !((p0).isLessThanOrEqualTo(p0));
            let _1794_v65;
            _1794_v65 = new _dafny.CodePoint('l'.codePointAt(0));
            let _1795_v66;
            _1795_v66 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
            let _1796_v67;
            _1796_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1786_v58,_1794_v65);
            let _1797_v68;
            _1797_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1790_v62).length));
            let _1798_v69;
            _1798_v69 = _dafny.Map.Empty.slice().updateUnsafe((((_1795_v66).contains(p1)) ? ((_1795_v66).get(p1)) : (p1)),(_this).fm10(_1796_v67, p0, _1797_v68, globalState));
            let _1799_v70;
            _1799_v70 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1798_v69);
            let _index304 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1713_v0).length));
            let _index305 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1713_v0).length));
            let _rhs364 = (p0).isEqualTo(p0);
            let _rhs365 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1786_v58, p1), _1789_v61), _1789_v61);
            let _rhs366 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("aqqe"), _dafny.Seq.update(_dafny.Seq.Concat(_1790_v62, _dafny.Seq.UnicodeFromString("behwcy")), _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_dafny.Seq.Concat(_1790_v62, _dafny.Seq.UnicodeFromString("behwcy"))).length)), _1794_v65));
            let _rhs367 = (_this).fm1(_1799_v70, p0, globalState);
            let _lhs255 = _1713_v0;
            let _lhs256 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1713_v0).length));
            let _lhs257 = _1713_v0;
            let _lhs258 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1713_v0).length));
            _1786_v58 = _rhs364;
            _1789_v61 = _rhs365;
            _lhs255[_lhs256] = _rhs366;
            _lhs257[_lhs258] = _rhs367;
            (globalState).f4 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber(488)));
            let _1800_v71;
            _1800_v71 = _module.D4.create_DC12(((p1) ? ((_dafny.ZERO).minus(new BigNumber((_1790_v62).length))) : (p0)));
            let _source25 = _1800_v71;
            if (_source25.is_DC12) {
              let _1801___mcc_h0 = (_source25).cf16;
              let _1802_cf16 = _1801___mcc_h0;
              let _1803_v72;
              let _nw260 = new _module.C2();
              _nw260.__ctor();
              _1803_v72 = _nw260;
              let _1804_v73;
              let _init46 = ((_1805_cf16) => function (_1806_i7) {
                return _module.__default.safeModuloInt(_1806_i7, _1805_cf16);
              })(_1802_cf16);
              let _nw261 = Array((new BigNumber(19)).toNumber());
              for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw261.length); _i0_46++) {
                _nw261[_i0_46] = _init46(new BigNumber(_i0_46));
              }
              _1804_v73 = _nw261;
              let _index306 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1804_v73).length));
              (_1804_v73)[_index306] = p0;
              let _1807_v74;
              let _nw262 = Array((new BigNumber(27)).toNumber()).fill(false);
              _1807_v74 = _nw262;
              let _1808_v75;
              _1808_v75 = _dafny.Map.Empty.slice().updateUnsafe(_1807_v74,((p1) ? (_1786_v58) : ((_this).fm8(new BigNumber(634), p1, p0, globalState))));
              let _index307 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1804_v73).length));
              (_1804_v73)[_index307] = new BigNumber((_1808_v75).length);
              let _1809_v76;
              let _nw263 = Array((new BigNumber(4)).toNumber()).fill([]);
              _1809_v76 = _nw263;
              let _1810_v77;
              _1810_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1786_v58,_1809_v76);
              let _1811_v78;
              _1811_v78 = _dafny.MultiSet.fromElements(_1802_cf16);
              _1810_v77 = (_1810_v77).update((_1811_v78).IsProperSubsetOf((_1811_v78).update(p0, _module.__default.abs(_1802_cf16))), _1809_v76);
              let _1812_v80;
              _1812_v80 = _module.D15.create_DC38(_dafny.Map.Empty.slice().updateUnsafe(_1802_cf16,_1786_v58));
              let _1813_v81;
              _1813_v81 = _dafny.MultiSet.fromElements(_module.D15.create_DC38(_1788_v60), _module.D15.create_DC38(function () {
  let _coll77 = new _dafny.Map();
  for (const _compr_77 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-584)), ((_1814_v73) => function (_1815_i8) {
    return (_1814_v73)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_1814_v73).length))];
  })(_1804_v73))).Elements) {
    let _1816_v79 = _compr_77;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-584)), ((_1817_v73) => function (_1815_i8) {
      return (_1817_v73)[_module.__default.safeIndex(new BigNumber(409), new BigNumber((_1817_v73).length))];
    })(_1804_v73)), _1816_v79)) {
      _coll77.push([(_1816_v79).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,false)).length)),p1]);
    }
  }
  return _coll77;
}()), _1812_v80);
              let _index308 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_1804_v73).length));
              (_1804_v73)[_index308] = (((_1813_v81).contains(_1812_v80)) ? ((_1813_v81).get(_1812_v80)) : (new BigNumber(((_dafny.Set.fromElements((_1713_v0)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1713_v0).length))], (_1713_v0)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_1713_v0).length))])).Union(_1771_v49)).length)));
            } else {
              let _1818___mcc_h1 = (_source25).cf15;
              let _1819_cf15 = _1818___mcc_h1;
              _1786_v58 = !(!(true));
              let _1820_v82;
              _1820_v82 = _module.D3.create_DC10(p0, p0);
              _1820_v82 = _module.__default.fm53((_this).fm2(p0, globalState), globalState);
              _1786_v58 = _1786_v58;
              let _1821_v83;
              _1821_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm9(p0, globalState),(p0).plus(p0));
              let _1822_v84;
              let _nw264 = Array((new BigNumber(21)).toNumber());
              _nw264[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
              _nw264[(_dafny.ONE).toNumber()] = p0;
              _nw264[(new BigNumber(2)).toNumber()] = (_this).fm2(new BigNumber((_1788_v60).length), globalState);
              _nw264[(new BigNumber(3)).toNumber()] = p0;
              _nw264[(new BigNumber(4)).toNumber()] = p0;
              _nw264[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1774_v52)).cardinality());
              _nw264[(new BigNumber(6)).toNumber()] = p0;
              _nw264[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("wcuryoktk")).length);
              _nw264[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1790_v62).length),_1786_v58)).length);
              _nw264[(new BigNumber(9)).toNumber()] = p0;
              _nw264[(new BigNumber(10)).toNumber()] = p0;
              _nw264[(new BigNumber(11)).toNumber()] = p0;
              _nw264[(new BigNumber(12)).toNumber()] = p0;
              _nw264[(new BigNumber(13)).toNumber()] = p0;
              _nw264[(new BigNumber(14)).toNumber()] = p0;
              _nw264[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Seq.update(_1790_v62, _module.__default.safeIndex(new BigNumber((_1798_v69).length), new BigNumber((_1790_v62).length)), _1794_v65)).length);
              _nw264[(new BigNumber(16)).toNumber()] = new BigNumber((_1790_v62).length);
              _nw264[(new BigNumber(17)).toNumber()] = p0;
              _nw264[(new BigNumber(18)).toNumber()] = p0;
              _nw264[(new BigNumber(19)).toNumber()] = p0;
              _nw264[(new BigNumber(20)).toNumber()] = (_1774_v52)[_module.__default.safeIndex(new BigNumber(364), new BigNumber((_1774_v52).length))];
              _1822_v84 = _nw264;
              let _1823_v85;
              _1823_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1822_v84,_1774_v52);
              let _1824_v86;
              _1824_v86 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(277)), ((_1825_p0) => function (_1826_i9) {
                return _1825_p0;
              })(p0)));
              let _1827_v87;
              _1827_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1774_v52).length),_1789_v61);
              let _1828_v88;
              _1828_v88 = _module.D1.create_DC3(_1790_v62, p1);
              _1821_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf((((_1823_v85).contains(_1822_v84)) ? ((_1823_v85).get(_1822_v84)) : ((_1824_v86)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((((_1827_v87).contains(p0)) ? ((_1827_v87).get(p0)) : (_1789_v61))).length)), new BigNumber((_1824_v86).length))])), _1774_v52),new BigNumber((_dafny.Seq.Concat((_1828_v88).dtor_cf3, _1790_v62)).length));
            }
            (globalState).f7 = ((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber(766)))).multipliedBy(p0);
          }
        }
      }
      let _1829_v89;
      let _nw265 = new _module.C0();
      _nw265.__ctor();
      _1829_v89 = _nw265;
      r0 = (_this).fm2((_dafny.ZERO).minus(p0), globalState);
      let _1830_v90;
      _1830_v90 = _dafny.Set.fromElements(_module.D0.create_DC0(_1713_v0));
      let _1831_v91;
      _1831_v91 = _dafny.MultiSet.fromElements(_1830_v90);
      r1 = ((_1831_v91).Union(_1831_v91)).Union(_1831_v91);
      return [r0, r1];
    }
  };

  $module.C10 = class C10 {
    constructor () {
      this._tname = "_module.C10";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return true;
    };
    fm2(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber(-601));
    };
    fm6(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(function (_source26) {
        if (_source26.is_DC3) {
          let _1832___mcc_h0 = (_source26).cf3;
          let _1833___mcc_h1 = (_source26).cf4;
          let _1834_cf4 = _1833___mcc_h1;
          let _1835_cf3 = _1832___mcc_h0;
          return (_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1834_cf4,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_1834_cf4, _1834_cf4)).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(621), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1834_cf4,_1834_cf4)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(60))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(919)), function (_1836_i0) {
  return new BigNumber(107);
})).length),_1834_cf4)).length), new BigNumber(262))).cardinality())))).length)))).length))).dtor_cf2;
        } else if (_source26.is_DC4) {
          let _1837___mcc_h2 = (_source26).cf5;
          let _1838_cf5 = _1837___mcc_h2;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(false, !(true), true, !(!(false))), _dafny.Seq.of(true, !(false)))).length));
        } else {
          let _1839___mcc_h3 = (_source26).cf2;
          let _1840_cf2 = _1839___mcc_h3;
          return _1840_cf2;
        }
      }(_module.D1.create_DC2(new BigNumber(647))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _1841_v0;
      let _nw266 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _1841_v0 = _nw266;
      let _1842_v1;
      _1842_v1 = _dafny.Seq.of(_1841_v0, _1841_v0);
      let _1843_v2;
      _1843_v2 = new BigNumber(216);
      _1841_v0 = (_1842_v1)[_module.__default.safeIndex(_1843_v2, new BigNumber((_1842_v1).length))];
      (globalState).f4 = (_1843_v2).minus(_1843_v2);
      let _1844_v3;
      _1844_v3 = true;
      let _1845_v4;
      _1845_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1844_v3,_1844_v3);
      _1844_v3 = (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(_1844_v3,_1845_v4), new BigNumber((_1845_v4).length), globalState);
      let _1846_v5;
      _1846_v5 = _dafny.MultiSet.fromElements(_1844_v3);
      let _1847_v6;
      _1847_v6 = _module.D2.create_DC7(true, _1846_v5, _1841_v0);
      let _pat_let_tv40 = _1841_v0;
      let _1848_v7;
      let _nw267 = Array((new BigNumber(9)).toNumber());
      _nw267[(_dafny.ZERO).toNumber()] = _1841_v0;
      _nw267[(_dafny.ONE).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(2)).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(3)).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(4)).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(5)).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(6)).toNumber()] = _1841_v0;
      _nw267[(new BigNumber(7)).toNumber()] = (function (_pat_let30_0) {
        return function (_1849_dt__update__tmp_h0) {
          return function (_pat_let31_0) {
            return function (_1850_dt__update_hcf10_h0) {
              return function (_pat_let32_0) {
                return function (_1851_dt__update_hcf8_h0) {
                  return _module.D2.create_DC7(_1851_dt__update_hcf8_h0, (_1849_dt__update__tmp_h0).dtor_cf9, _1850_dt__update_hcf10_h0);
                }(_pat_let32_0);
              }(false);
            }(_pat_let31_0);
          }(_pat_let_tv40);
        }(_pat_let30_0);
      }(_1847_v6)).dtor_cf10;
      _nw267[(new BigNumber(8)).toNumber()] = _1841_v0;
      _1848_v7 = _nw267;
      let _nw268 = Array((new BigNumber(21)).toNumber()).fill([]);
      _1848_v7 = _nw268;
      let _hi13 = _module.__default.safeDivisionInt(new BigNumber((_1846_v5).cardinality()), (_dafny.ZERO).minus(_1843_v2));
      for (let _1852_i0 = _1843_v2; _1852_i0.isLessThan(_hi13); _1852_i0 = _1852_i0.plus(_dafny.ONE)) {
        let _1853_v8;
        let _nw269 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1853_v8 = _nw269;
        _1853_v8 = _1853_v8;
        let _1854_v9;
        _1854_v9 = new _dafny.CodePoint('j'.codePointAt(0));
        _1854_v9 = _1854_v9;
        let _1855_v10;
        _1855_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1844_v3,_1845_v4);
        let _1856_v11;
        _1856_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1843_v2,(_dafny.ZERO).minus(new BigNumber((_1846_v5).cardinality())));
        let _1857_v12;
        _1857_v12 = _dafny.Map.Empty.slice().updateUnsafe(((_this).fm6(false, (_this).fm1(_1855_v10, _1843_v2, globalState), _1856_v11, new _dafny.CodePoint('w'.codePointAt(0)), globalState)).minus(new BigNumber(446)),(_1846_v5).IsProperSubsetOf(_dafny.MultiSet.fromElements(_1844_v3)));
        _1857_v12 = (_1857_v12).update((_dafny.ZERO).minus(_1852_i0), _1844_v3);
        let _1858_v13;
        let _nw270 = new _module.C4();
        _nw270.__ctor();
        _1858_v13 = _nw270;
      }
      let _1859_v14;
      _1859_v14 = new _dafny.CodePoint('u'.codePointAt(0));
      _1859_v14 = _1859_v14;
      return;
    }
    m5(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let _1860_v0;
      _1860_v0 = _dafny.Set.fromElements(p0, p0, p0);
      let _hi14 = p0;
      for (let _1861_i0 = new BigNumber((_1860_v0).length); _1861_i0.isLessThan(_hi14); _1861_i0 = _1861_i0.plus(_dafny.ONE)) {
        r1 = !(true) || (p2);
        r1 = (_1861_i0).isLessThan(p0);
        let _1862_v1;
        _1862_v1 = _dafny.MultiSet.fromElements(_1861_i0, _1861_i0);
        let _1863_v2;
        _1863_v2 = _dafny.Seq.of(p0);
        r1 = (_1862_v1).IsSubsetOf((_1862_v1).Union(_dafny.MultiSet.FromArray(_1863_v2)));
        let _1864_v3;
        let _init47 = ((_1865_i0) => function (_1866_i1) {
          return (_1866_i1).plus(_1865_i0);
        })(_1861_i0);
        let _nw271 = Array((new BigNumber(26)).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw271.length); _i0_47++) {
          _nw271[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1864_v3 = _nw271;
        let _1867_v4;
        _1867_v4 = _dafny.MultiSet.fromElements(false);
        let _index309 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1864_v3).length));
        (_1864_v3)[_index309] = (((_1862_v1).contains(new BigNumber((_1867_v4).cardinality()))) ? ((_1862_v1).get(new BigNumber((_1867_v4).cardinality()))) : (p0));
        let _1868_v5;
        _1868_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _index310 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1864_v3).length));
        (_1864_v3)[_index310] = (_dafny.ZERO).minus(((_1861_i0).plus(p0)).multipliedBy((((_1868_v5).contains(false)) ? ((_1868_v5).get(false)) : ((_dafny.ZERO).minus(new BigNumber(-617))))));
      }
      let _1869_v6;
      _1869_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
      let _1870_v7;
      _1870_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1869_v6);
      let _1871_i2;
      _1871_i2 = _dafny.ZERO;
      L12: {
        while ((_this).fm1(_1870_v7, p0, globalState)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1871_i2)) {
              break L12;
            }
            _1871_i2 = (_1871_i2).plus(_dafny.ONE);
            let _1872_v8;
            _1872_v8 = _dafny.Seq.of(p1);
            r1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.update(_1872_v8, _module.__default.safeIndex(p0, new BigNumber((_1872_v8).length)), p2), _1872_v8), _1872_v8);
            let _1873_v9;
            _1873_v9 = new _dafny.CodePoint('k'.codePointAt(0));
            r2 = _1873_v9;
            r1 = ((false) ? (p2) : (p1));
            (globalState).f4 = p0;
          }
        }
      }
      let _1874_v10;
      _1874_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _1875_v11;
      _1875_v11 = _dafny.Seq.of(!(!(p2)), p1);
      let _1876_v12;
      _1876_v12 = _module.D12.create_DC33(p1);
      let _1877_v13;
      _1877_v13 = new _dafny.CodePoint('k'.codePointAt(0));
      let _1878_v14;
      _1878_v14 = _module.D11.create_DC30(p1, p0);
      let _1879_v15;
      _1879_v15 = _dafny.Seq.of(new BigNumber(620), p0);
      let _1880_v16;
      _1880_v16 = _dafny.MultiSet.fromElements(!(p2), p1);
      let _1881_v17;
      let _nw272 = new _module.C5();
      _nw272.__ctor();
      _1881_v17 = _nw272;
      let _1882_v18;
      let _nw273 = Array((new BigNumber(27)).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = p1;
      _nw273[(_dafny.ONE).toNumber()] = p2;
      _nw273[(new BigNumber(2)).toNumber()] = (_this).fm1(_1870_v7, new BigNumber((_1874_v10).length), globalState);
      _nw273[(new BigNumber(3)).toNumber()] = p1;
      _nw273[(new BigNumber(4)).toNumber()] = true;
      _nw273[(new BigNumber(5)).toNumber()] = !(p1);
      _nw273[(new BigNumber(6)).toNumber()] = (p0).isLessThan(p0);
      _nw273[(new BigNumber(7)).toNumber()] = p1;
      _nw273[(new BigNumber(8)).toNumber()] = (_this).fm1(_1870_v7, p0, globalState);
      _nw273[(new BigNumber(9)).toNumber()] = p1;
      _nw273[(new BigNumber(10)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1875_v11, _1875_v11);
      _nw273[(new BigNumber(11)).toNumber()] = !(p2);
      _nw273[(new BigNumber(12)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
      _nw273[(new BigNumber(13)).toNumber()] = ((p2) ? (p1) : ((_this).fm1(_1870_v7, p0, globalState)));
      _nw273[(new BigNumber(14)).toNumber()] = (_1876_v12).dtor_cf50;
      _nw273[(new BigNumber(15)).toNumber()] = ((_this).fm6(p2, p1, _1874_v10, _1877_v13, globalState)).isEqualTo((_this).fm6(p2, p1, _dafny.Map.Empty.slice().updateUnsafe((_1878_v14).dtor_cf47,p0), _1877_v13, globalState));
      _nw273[(new BigNumber(16)).toNumber()] = p2;
      _nw273[(new BigNumber(17)).toNumber()] = p2;
      _nw273[(new BigNumber(18)).toNumber()] = p1;
      _nw273[(new BigNumber(19)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1879_v15, _1879_v15);
      _nw273[(new BigNumber(20)).toNumber()] = p1;
      _nw273[(new BigNumber(21)).toNumber()] = p1;
      _nw273[(new BigNumber(22)).toNumber()] = (_1880_v16).IsProperSubsetOf(_1880_v16);
      _nw273[(new BigNumber(23)).toNumber()] = p2;
      _nw273[(new BigNumber(24)).toNumber()] = false;
      _nw273[(new BigNumber(25)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(p2,_1881_v17)).contains(!(p1));
      _nw273[(new BigNumber(26)).toNumber()] = p1;
      _1882_v18 = _nw273;
      _1882_v18 = _1882_v18;
      let _1883_v19;
      _1883_v19 = _dafny.Seq.UnicodeFromString("lpee");
      _1883_v19 = _dafny.Seq.update(_module.__default.fm19(p2, p1, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm19(p2, p1, globalState)).length)), _1877_v13);
      if (false) {
        let _1884_v20;
        let _nw274 = Array((new BigNumber(23)).toNumber()).fill(_module.D19.Default());
        _1884_v20 = _nw274;
        _1884_v20 = _1884_v20;
        (globalState).f15 = p0;
        _1883_v19 = _module.__default.fm30((p0).minus(p0), !(p0).isEqualTo(p0), _1860_v0, ((true) ? (p1) : (p1)), globalState);
        if (p2) {
          r1 = p1;
          let _1885_v21;
          let _nw275 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _1885_v21 = _nw275;
          let _index311 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1885_v21).length));
          (_1885_v21)[_index311] = (_dafny.ZERO).minus(p0);
          let _index312 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1885_v21).length));
          (_1885_v21)[_index312] = p0;
          let _1886_v22;
          _1886_v22 = _dafny.Map.Empty.slice().updateUnsafe(((p2) ? (_1882_v18) : (_1882_v18)),_module.__default.fm17(p0, globalState));
          (globalState).f15 = (((_1886_v22).contains(_1882_v18)) ? ((_1886_v22).get(_1882_v18)) : ((_dafny.ZERO).minus((_1885_v21)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_1885_v21).length))])));
          let _1887_v23;
          let _nw276 = Array((new BigNumber(17)).toNumber());
          _nw276[(_dafny.ZERO).toNumber()] = _1885_v21;
          _nw276[(_dafny.ONE).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(2)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(3)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(4)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(5)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(6)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(7)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(8)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(9)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(10)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(11)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(12)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(13)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(14)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(15)).toNumber()] = _1885_v21;
          _nw276[(new BigNumber(16)).toNumber()] = _1885_v21;
          _1887_v23 = _nw276;
          let _1888_v24;
          _1888_v24 = _module.D20.create_DC54(_1887_v23);
          let _1889_v25;
          let _nw277 = Array((new BigNumber(4)).toNumber());
          _nw277[(_dafny.ZERO).toNumber()] = _1887_v23;
          _nw277[(_dafny.ONE).toNumber()] = _1887_v23;
          _nw277[(new BigNumber(2)).toNumber()] = _1887_v23;
          _nw277[(new BigNumber(3)).toNumber()] = (_1888_v24).dtor_cf84;
          _1889_v25 = _nw277;
          let _index313 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1889_v25).length));
          (_1889_v25)[_index313] = _1887_v23;
          let _index314 = _module.__default.safeIndex(new BigNumber(673), new BigNumber((_1889_v25).length));
          (_1889_v25)[_index314] = _1887_v23;
          let _1890_v26;
          _1890_v26 = _dafny.Set.fromElements(p1);
          let _rhs368 = _1885_v21;
          let _rhs369 = (_1885_v21)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_1885_v21).length))];
          let _rhs370 = (_1890_v26).IsProperSubsetOf(_1890_v26);
          let _rhs371 = _dafny.areEqual(_1883_v19, _module.__default.fm19(false, p1, globalState));
          let _lhs259 = globalState;
          _1885_v21 = _rhs368;
          _lhs259.f13 = _rhs369;
          r1 = _rhs370;
          r1 = _rhs371;
        } else {
          let _1891_v27;
          let _init48 = ((_1892_p1) => function (_1893_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_1892_p1,new BigNumber(789));
          })(p1);
          let _nw278 = Array((new BigNumber(18)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw278.length); _i0_48++) {
            _nw278[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1891_v27 = _nw278;
          let _index315 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1891_v27).length));
          (_1891_v27)[_index315] = (_dafny.Map.Empty.slice().updateUnsafe(p1,p0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_1879_v15).length)));
          let _1894_v28;
          _1894_v28 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _index316 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1891_v27).length));
          let _rhs372 = p1;
          let _rhs373 = (((p1) ? (_1894_v28) : (_dafny.Map.Empty.slice().updateUnsafe(p1,p0)))).Merge(_1894_v28);
          let _rhs374 = p0;
          let _rhs375 = _1877_v13;
          let _lhs260 = _1891_v27;
          let _lhs261 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_1891_v27).length));
          let _lhs262 = globalState;
          r1 = _rhs372;
          _lhs260[_lhs261] = _rhs373;
          _lhs262.f7 = _rhs374;
          _1877_v13 = _rhs375;
          let _1895_v29;
          let _nw279 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _1895_v29 = _nw279;
          let _index317 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_1895_v29).length));
          (_1895_v29)[_index317] = p0;
          let _index318 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_1895_v29).length));
          (_1895_v29)[_index318] = (_1879_v15)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("booqtq")).length), new BigNumber((_1879_v15).length))];
          let _1896_v30;
          let _nw280 = new _module.C8();
          _nw280.__ctor();
          _1896_v30 = _nw280;
          r1 = !((p2) || (p2));
          let _1897_v31;
          _1897_v31 = _dafny.Seq.of(_1882_v18);
          let _1898_v32;
          let _nw281 = Array((new BigNumber(23)).toNumber());
          _nw281[(_dafny.ZERO).toNumber()] = _1882_v18;
          _nw281[(_dafny.ONE).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(2)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(3)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(4)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(5)).toNumber()] = ((p2) ? ((_1897_v31)[_module.__default.safeIndex(new BigNumber((_1874_v10).length), new BigNumber((_1897_v31).length))]) : (_1882_v18));
          _nw281[(new BigNumber(6)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(7)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(8)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(9)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(10)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(11)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(12)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(13)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(14)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(15)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(16)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(17)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(18)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(19)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(20)).toNumber()] = ((p2) ? (_1882_v18) : (_1882_v18));
          _nw281[(new BigNumber(21)).toNumber()] = _1882_v18;
          _nw281[(new BigNumber(22)).toNumber()] = _1882_v18;
          _1898_v32 = _nw281;
          let _index319 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1898_v32).length));
          (_1898_v32)[_index319] = _1882_v18;
          let _index320 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_1898_v32).length));
          (_1898_v32)[_index320] = _1882_v18;
        }
        if (p1) {
          r1 = p2;
          let _1899_v33;
          let _init49 = ((_1900_v0) => function (_1901_i4) {
            return _1900_v0;
          })(_1860_v0);
          let _nw282 = Array((new BigNumber(9)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw282.length); _i0_49++) {
            _nw282[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _1899_v33 = _nw282;
          let _index321 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1899_v33).length));
          (_1899_v33)[_index321] = _dafny.Set.fromElements(p0);
          let _index322 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_1899_v33).length));
          (_1899_v33)[_index322] = _1860_v0;
          let _1902_v34;
          _1902_v34 = _dafny.MultiSet.fromElements(p0);
          r1 = ((_1902_v34).Difference(_1902_v34)).IsSubsetOf(_1902_v34);
          let _1903_v35;
          _1903_v35 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _1904_v36;
          _1904_v36 = _1903_v35;
          let _1905_v37;
          _1905_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1903_v35,_dafny.Seq.Concat(_1883_v19, _dafny.Seq.Create(_module.__default.abs(new BigNumber(976)), function (_1906_i5) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          })));
          _1905_v37 = (_1905_v37).update(_1904_v36, _1883_v19);
          let _1907_v38;
          _1907_v38 = _dafny.Set.fromElements(p1);
          let _1908_v39;
          _1908_v39 = _dafny.Map.Empty.slice().updateUnsafe((p0).multipliedBy(p0),_1907_v38);
          _1908_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1907_v38);
        } else {
          let _1909_v40;
          let _nw283 = new _module.C2();
          _nw283.__ctor();
          _1909_v40 = _nw283;
          let _1910_v41;
          _1910_v41 = _dafny.Seq.of(_1909_v40);
          let _1911_v42;
          _1911_v42 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
          _1909_v40 = (_1910_v41)[_module.__default.safeIndex((((_1911_v42).contains(!(p1))) ? ((_1911_v42).get(!(p1))) : (p0)), new BigNumber((_1910_v41).length))];
          let _1912_v43;
          let _nw284 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1912_v43 = _nw284;
          let _1913_v44;
          _1913_v44 = _module.D21.create_DC56(_1912_v43);
          let _1914_v45;
          let _nw285 = Array((new BigNumber(23)).toNumber());
          _nw285[(_dafny.ZERO).toNumber()] = _1912_v43;
          _nw285[(_dafny.ONE).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(2)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(3)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(4)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(5)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(6)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(7)).toNumber()] = ((p2) ? (_1912_v43) : (_1912_v43));
          _nw285[(new BigNumber(8)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(9)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(10)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(11)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(12)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(13)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(14)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(15)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(16)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(17)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(18)).toNumber()] = (_1913_v44).dtor_cf89;
          _nw285[(new BigNumber(19)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(20)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(21)).toNumber()] = _1912_v43;
          _nw285[(new BigNumber(22)).toNumber()] = _1912_v43;
          _1914_v45 = _nw285;
          let _index323 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1914_v45).length));
          (_1914_v45)[_index323] = _1912_v43;
          let _index324 = _module.__default.safeIndex(new BigNumber(182), new BigNumber((_1914_v45).length));
          (_1914_v45)[_index324] = _1912_v43;
          let _1915_v46;
          _1915_v46 = _module.D5.create_DC14(_1877_v13, !(p2), p0, p1);
          let _1916_v47;
          _1916_v47 = _module.D5.create_DC15(_1915_v46);
          _1916_v47 = _1916_v47;
          let _1917_v48;
          let _nw286 = new _module.C0();
          _nw286.__ctor();
          _1917_v48 = _nw286;
          (globalState).f7 = _module.__default.safeModuloInt(p0, new BigNumber((_1879_v15).length));
        }
      } else {
        let _1918_v49;
        let _nw287 = Array((new BigNumber(2)).toNumber());
        _1918_v49 = _nw287;
        let _1919_v50;
        _1919_v50 = _dafny.MultiSet.fromElements(_1918_v49);
        let _1920_v51;
        _1920_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1860_v0,new BigNumber((_1919_v50).cardinality()));
        (globalState).f7 = (((_1920_v51).contains((_1860_v0).Union(_1860_v0))) ? ((_1920_v51).get((_1860_v0).Union(_1860_v0))) : (new BigNumber(781)));
        r1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_1879_v15, _1879_v15), _1879_v15);
        let _1921_v52;
        let _nw288 = new _module.C7();
        _nw288.__ctor();
        _1921_v52 = _nw288;
        let _1922_v53;
        _1922_v53 = _module.D17.create_DC47(_1921_v52, p1, p0, p0, p0);
        let _1923_v54;
        _1923_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1922_v53,!_dafny.Seq.contains(_1883_v19, new _dafny.CodePoint('h'.codePointAt(0))));
        let _rhs376 = _1882_v18;
        let _rhs377 = _1923_v54;
        _1882_v18 = _rhs376;
        _1923_v54 = _rhs377;
        let _rhs378 = (p0).multipliedBy(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(913)), ((_1924_v13) => function (_1925_i6) {
          return _1924_v13;
        })(_1877_v13)))).length))).multipliedBy((_1879_v15)[_module.__default.safeIndex(p0, new BigNumber((_1879_v15).length))]));
        let _rhs379 = p0;
        let _rhs380 = p1;
        let _lhs263 = globalState;
        let _lhs264 = globalState;
        _lhs263.f7 = _rhs378;
        _lhs264.f15 = _rhs379;
        r1 = _rhs380;
        let _1926_v55;
        let _nw289 = new _module.C4();
        _nw289.__ctor();
        _1926_v55 = _nw289;
      }
      if (p2) {
        (globalState).f7 = ((!(true)) ? (p0) : (p0));
        if (p2) {
          let _1927_v56;
          let _nw290 = new _module.C8();
          _nw290.__ctor();
          _1927_v56 = _nw290;
          let _1928_v57;
          _1928_v57 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1927_v56);
          _1928_v57 = (_1928_v57).update(p2, ((p2) ? (_1927_v56) : (_1927_v56)));
          let _init50 = ((_1929_p2) => function (_1930_i7) {
            return (_1929_p2) === (_1929_p2);
          })(p2);
          let _nw291 = Array((new BigNumber(10)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw291.length); _i0_50++) {
            _nw291[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1882_v18 = _nw291;
          r1 = p1;
          let _1931_v58;
          _1931_v58 = _dafny.Seq.of(_1869_v6);
          let _1932_v59;
          _1932_v59 = _module.D16.create_DC41(_1931_v58);
          let _1933_v60;
          _1933_v60 = _module.D16.create_DC44(_1932_v59);
          _1933_v60 = _1933_v60;
          let _index325 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index325] = _dafny.Seq.IsProperPrefixOf(_1875_v11, _1875_v11);
          let _index326 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index326] = p1;
        } else {
          let _1934_v61;
          let _nw292 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1934_v61 = _nw292;
          let _index327 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1934_v61).length));
          (_1934_v61)[_index327] = (new BigNumber(986)).minus(p0);
          let _1935_v62;
          _1935_v62 = _module.D3.create_DC10(p0, p0);
          let _1936_v63;
          _1936_v63 = _module.D2.create_DC7(p2, _1880_v16, _1934_v61);
          let _1937_v64;
          _1937_v64 = _dafny.Seq.of(_1936_v63, _1936_v63, _1936_v63);
          let _index328 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1934_v61).length));
          let _rhs381 = _1875_v11;
          let _rhs382 = p1;
          let _rhs383 = (_1935_v62).dtor_cf14;
          let _rhs384 = _dafny.areEqual(_1937_v64, _dafny.Seq.Concat(_1937_v64, _1937_v64));
          let _rhs385 = ((p0).multipliedBy(p0)).minus(new BigNumber(-95));
          let _lhs265 = globalState;
          let _lhs266 = _1934_v61;
          let _lhs267 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1934_v61).length));
          _1875_v11 = _rhs381;
          r1 = _rhs382;
          _lhs265.f13 = _rhs383;
          r1 = _rhs384;
          _lhs266[_lhs267] = _rhs385;
          r1 = p2;
          _1934_v61 = _1934_v61;
          let _index329 = _module.__default.safeIndex(new BigNumber(662), new BigNumber((_1934_v61).length));
          (_1934_v61)[_index329] = p0;
          let _1938_v65;
          _1938_v65 = _module.D2.create_DC5(_1934_v61);
          (_1881_v17).m6(((p2) ? (new BigNumber((_dafny.Seq.update(_1875_v11, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(128),p1)).length), new BigNumber((_1875_v11).length)), p1)).length)) : ((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_module.__default.fm19(p2, (_module.D19.create_DC52(p0, p1, p2)).dtor_cf82, globalState)).length)))))), (_1880_v16).Union(_1880_v16), ((p2) ? (p2) : (p1)), _1938_v65, globalState);
        }
        let _1939_v66;
        let _nw293 = new _module.C0();
        _nw293.__ctor();
        _1939_v66 = _nw293;
        let _1940_v67;
        _1940_v67 = _dafny.Set.fromElements(_1860_v0);
        let _1941_v68;
        let _nw294 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1941_v68 = _nw294;
        let _1942_v69;
        _1942_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1941_v68);
        let _1943_v70;
        _1943_v70 = _dafny.Seq.of(_1882_v18, _1882_v18, _1882_v18);
        let _1944_v71;
        _1944_v71 = _dafny.MultiSet.fromElements(p0);
        let _1945_v72;
        _1945_v72 = _module.D19.create_DC52(new BigNumber(297), false, p2);
        let _1946_v73;
        let _nw295 = Array((new BigNumber(27)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = p0;
        _nw295[(_dafny.ONE).toNumber()] = p0;
        _nw295[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Set.fromElements(p0, new BigNumber((_1942_v69).length), p0)).length);
        _nw295[(new BigNumber(3)).toNumber()] = p0;
        _nw295[(new BigNumber(4)).toNumber()] = new BigNumber((_1943_v70).length);
        _nw295[(new BigNumber(5)).toNumber()] = p0;
        _nw295[(new BigNumber(6)).toNumber()] = p0;
        _nw295[(new BigNumber(7)).toNumber()] = new BigNumber((_1883_v19).length);
        _nw295[(new BigNumber(8)).toNumber()] = new BigNumber((_1944_v71).cardinality());
        _nw295[(new BigNumber(9)).toNumber()] = p0;
        _nw295[(new BigNumber(10)).toNumber()] = p0;
        _nw295[(new BigNumber(11)).toNumber()] = (((_1874_v10).contains(new BigNumber((_1883_v19).length))) ? ((_1874_v10).get(new BigNumber((_1883_v19).length))) : (p0));
        _nw295[(new BigNumber(12)).toNumber()] = p0;
        _nw295[(new BigNumber(13)).toNumber()] = new BigNumber(827);
        _nw295[(new BigNumber(14)).toNumber()] = p0;
        _nw295[(new BigNumber(15)).toNumber()] = p0;
        _nw295[(new BigNumber(16)).toNumber()] = p0;
        _nw295[(new BigNumber(17)).toNumber()] = (_1945_v72).dtor_cf81;
        _nw295[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw295[(new BigNumber(19)).toNumber()] = new BigNumber(607);
        _nw295[(new BigNumber(20)).toNumber()] = p0;
        _nw295[(new BigNumber(21)).toNumber()] = new BigNumber((_module.__default.fm31(p0, p0, globalState)).cardinality());
        _nw295[(new BigNumber(22)).toNumber()] = new BigNumber((_1883_v19).length);
        _nw295[(new BigNumber(23)).toNumber()] = p0;
        _nw295[(new BigNumber(24)).toNumber()] = p0;
        _nw295[(new BigNumber(25)).toNumber()] = p0;
        _nw295[(new BigNumber(26)).toNumber()] = p0;
        _1946_v73 = _nw295;
        let _1947_v74;
        _1947_v74 = _dafny.MultiSet.fromElements(_1946_v73);
        let _1948_v75;
        let _nw296 = Array((new BigNumber(4)).toNumber());
        _nw296[(_dafny.ZERO).toNumber()] = _1947_v74;
        _nw296[(_dafny.ONE).toNumber()] = ((_1947_v74).update(_1946_v73, _module.__default.abs(p0))).Difference(_1947_v74);
        _nw296[(new BigNumber(2)).toNumber()] = _1947_v74;
        _nw296[(new BigNumber(3)).toNumber()] = _1947_v74;
        _1948_v75 = _nw296;
        let _index330 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1948_v75).length));
        (_1948_v75)[_index330] = _1947_v74;
        let _index331 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length));
        (_1946_v73)[_index331] = ((_dafny.ZERO).minus(p0)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1875_v11).length)));
        let _index332 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1948_v75).length));
        let _index333 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length));
        let _rhs386 = ((_1940_v67).Union(_1940_v67)).Difference((_1940_v67).Difference(_dafny.Set.fromElements(_1860_v0, _1860_v0)));
        let _rhs387 = (p0).plus(new BigNumber((_1875_v11).length));
        let _rhs388 = ((!(!(p2))) ? ((_1947_v74).Difference(_1947_v74)) : (_1947_v74));
        let _rhs389 = !(p0).isEqualTo(p0);
        let _rhs390 = _module.__default.safeDivisionInt((_1879_v15)[_module.__default.safeIndex(p0, new BigNumber((_1879_v15).length))], p0);
        let _lhs268 = globalState;
        let _lhs269 = _1948_v75;
        let _lhs270 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_1948_v75).length));
        let _lhs271 = _1946_v73;
        let _lhs272 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length));
        _1940_v67 = _rhs386;
        _lhs268.f13 = _rhs387;
        _lhs269[_lhs270] = _rhs388;
        r1 = _rhs389;
        _lhs271[_lhs272] = _rhs390;
        if (!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1875_v11, _dafny.Seq.of(p2, p2)), _dafny.Seq.Concat(_1875_v11, _1875_v11)))) {
          let _1949_v76;
          let _nw297 = new _module.C5();
          _nw297.__ctor();
          _1949_v76 = _nw297;
          (_1949_v76).m6((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))], _module.__default.fm34((_dafny.ZERO).minus(new BigNumber(-881)), (_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))], !(p2), globalState), !(true), _module.D2.create_DC5(_1941_v68), globalState);
          (globalState).f4 = ((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))]).multipliedBy((_dafny.ZERO).minus((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))]));
          let _1950_v77;
          let _nw298 = Array((new BigNumber(6)).toNumber()).fill([]);
          _1950_v77 = _nw298;
          let _rhs391 = ((p2) ? (_1882_v18) : (_1882_v18));
          let _rhs392 = _1950_v77;
          let _rhs393 = (_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))];
          let _lhs273 = globalState;
          _1882_v18 = _rhs391;
          _1950_v77 = _rhs392;
          _lhs273.f13 = _rhs393;
          let _1951_v78;
          _1951_v78 = _dafny.Seq.of(_1941_v68, _1946_v73);
          let _1952_v79;
          _1952_v79 = _dafny.Map.Empty.slice().updateUnsafe((_1951_v78)[_module.__default.safeIndex(p0, new BigNumber((_1951_v78).length))],(_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))]);
          _1952_v79 = _1952_v79;
        } else {
          let _1953_v80;
          let _init51 = ((_1954_p1, _1955_p2, _1956_v73, _1957_p0) => function (_1958_i8) {
            return ((_1954_p1) ? (_dafny.Map.Empty.slice().updateUnsafe(_1955_p2,(_1956_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1956_v73).length))])) : (_dafny.Map.Empty.slice().updateUnsafe(_1954_p1,_1957_p0)));
          })(p1, p2, _1946_v73, p0);
          let _nw299 = Array((new BigNumber(9)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw299.length); _i0_51++) {
            _nw299[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1953_v80 = _nw299;
          let _1959_v81;
          _1959_v81 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _index334 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1953_v80).length));
          (_1953_v80)[_index334] = (_1959_v81).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1945_v72).dtor_cf83,p0));
          let _index335 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_1953_v80).length));
          (_1953_v80)[_index335] = (_1959_v81).Merge(_1959_v81);
          let _1960_v82;
          _1960_v82 = _dafny.Set.fromElements(p2);
          let _1961_v83;
          _1961_v83 = _dafny.Seq.of(_1960_v82, _1960_v82, _1960_v82);
          let _1962_v84;
          _1962_v84 = _module.D5.create_DC14(_1877_v13, false, p0, p2);
          let _rhs394 = false;
          let _rhs395 = (_module.D18.create_DC49(!(p1), p2, (_1962_v84).dtor_cf18, p1)).dtor_cf78;
          let _rhs396 = _1961_v83;
          let _rhs397 = _1879_v15;
          r1 = _rhs394;
          r1 = _rhs395;
          _1961_v83 = _rhs396;
          _1879_v15 = _rhs397;
          let _1963_v85;
          _1963_v85 = _module.D19.create_DC53();
          _1963_v85 = _module.__default.fm54((_1939_v66).fm29(_1875_v11, _1883_v19, p2, p0, globalState), _1869_v6, _1944_v71, globalState);
          let _1964_v86;
          let _nw300 = new _module.C0();
          _nw300.__ctor();
          _1964_v86 = _nw300;
          let _1965_v88;
          _1965_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1877_v13,(_this).fm2(new BigNumber((_1880_v16).cardinality()), globalState));
          let _1966_v89;
          _1966_v89 = _dafny.Seq.of(function () {
            let _coll78 = new _dafny.Map();
            for (const _compr_78 of (_1965_v88).Keys.Elements) {
              let _1967_v87 = _compr_78;
              if ((_1965_v88).contains(_1967_v87)) {
                _coll78.push([_1967_v87,p2]);
              }
            }
            return _coll78;
          }());
          let _1968_v90;
          let _nw301 = new _module.C2();
          _nw301.__ctor();
          _1968_v90 = _nw301;
          let _1969_v91;
          _1969_v91 = _module.D17.create_DC47(_1968_v90, p2, (_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))], new BigNumber((_1883_v19).length), (_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))]);
          let _1970_v92;
          _1970_v92 = _dafny.Seq.of((_1953_v80)[_module.__default.safeIndex(new BigNumber(149), new BigNumber((_1953_v80).length))]);
          let _1971_v93;
          _1971_v93 = _dafny.Map.Empty.slice().updateUnsafe((_1883_v19)[_module.__default.safeIndex(new BigNumber(((_1970_v92)[_module.__default.safeIndex(p0, new BigNumber((_1970_v92).length))]).length), new BigNumber((_1883_v19).length))],(_1881_v17).fm8(p0, false, (_1968_v90).fm2(p0, globalState), globalState));
          let _1972_v94;
          _1972_v94 = _module.D14.create_DC36(_1875_v11);
          let _1973_v95;
          _1973_v95 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(p1));
          let _1974_v96;
          let _nw302 = Array((new BigNumber(12)).toNumber());
          _nw302[(_dafny.ZERO).toNumber()] = _1966_v89;
          _nw302[(_dafny.ONE).toNumber()] = _1966_v89;
          _nw302[(new BigNumber(2)).toNumber()] = _1966_v89;
          _nw302[(new BigNumber(3)).toNumber()] = _1966_v89;
          _nw302[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1966_v89, _module.__default.safeIndex((_1969_v91).dtor_cf73, new BigNumber((_1966_v89).length)), _1971_v93);
          _nw302[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(881)), ((_1975_v93) => function (_1976_i9) {
            return _1975_v93;
          })(_1971_v93));
          _nw302[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(858)), ((_1977_v93) => function (_1978_i10) {
            return _1977_v93;
          })(_1971_v93));
          _nw302[(new BigNumber(7)).toNumber()] = _1966_v89;
          _nw302[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1966_v89, _dafny.Seq.of(_1971_v93, _1971_v93, _1971_v93, _1971_v93));
          _nw302[(new BigNumber(9)).toNumber()] = _module.__default.fm55(p1, p0, p2, _1972_v94, globalState);
          _nw302[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1966_v89, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-814)), ((_1979_v93) => function (_1980_i11) {
            return _1979_v93;
          })(_1971_v93)));
          _nw302[(new BigNumber(11)).toNumber()] = _module.__default.fm55(p1, new BigNumber(((((_1973_v95).contains((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))])) ? ((_1973_v95).get((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))])) : (_1875_v11))).length), p1, _1972_v94, globalState);
          _1974_v96 = _nw302;
          let _index336 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1974_v96).length));
          (_1974_v96)[_index336] = _1966_v89;
          let _1981_v97;
          _1981_v97 = _dafny.Map.Empty.slice().updateUnsafe((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))],_1877_v13);
          let _index337 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1974_v96).length));
          let _rhs398 = (p0).plus((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))]);
          let _rhs399 = p1;
          let _rhs400 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1966_v89, _1966_v89), _dafny.Seq.Concat(_dafny.Seq.update(_1966_v89, _module.__default.safeIndex((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))], new BigNumber((_1966_v89).length)), _dafny.Map.Empty.slice().updateUnsafe(_1877_v13,false)), _dafny.Seq.update(_1966_v89, _module.__default.safeIndex((((_1959_v81).contains(p2)) ? ((_1959_v81).get(p2)) : ((_1946_v73)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_1946_v73).length))])), new BigNumber((_1966_v89).length)), _dafny.Map.Empty.slice().updateUnsafe((((_1981_v97).contains((_dafny.ZERO).minus(new BigNumber((_1880_v16).cardinality())))) ? ((_1981_v97).get((_dafny.ZERO).minus(new BigNumber((_1880_v16).cardinality())))) : (_1877_v13)),p1))));
          let _lhs274 = globalState;
          let _lhs275 = _1974_v96;
          let _lhs276 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_1974_v96).length));
          _lhs274.f15 = _rhs398;
          r1 = _rhs399;
          _lhs275[_lhs276] = _rhs400;
        }
      } else {
        let _1982_v98;
        _1982_v98 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), ((_1983_v13) => function (_1984_i12) {
  return _1983_v13;
})(_1877_v13)), p1);
        let _pat_let_tv41 = _1883_v19;
        let _source27 = function (_pat_let33_0) {
          return function (_1985_dt__update__tmp_h1) {
            return function (_pat_let34_0) {
              return function (_1986_dt__update_hcf3_h0) {
                return _module.D1.create_DC3(_1986_dt__update_hcf3_h0, (_1985_dt__update__tmp_h1).dtor_cf4);
              }(_pat_let34_0);
            }(_pat_let_tv41);
          }(_pat_let33_0);
        }(_1982_v98);
        if (_source27.is_DC3) {
          let _1987___mcc_h0 = (_source27).cf3;
          let _1988___mcc_h1 = (_source27).cf4;
          let _1989_cf4 = _1988___mcc_h1;
          let _1990_cf3 = _1987___mcc_h0;
          let _1991_v99;
          _1991_v99 = _module.D3.create_DC10(p0, new BigNumber((_1880_v16).cardinality()));
          (globalState).f4 = (_1991_v99).dtor_cf13;
          let _1992_v100;
          let _nw303 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1992_v100 = _nw303;
          let _index338 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1992_v100).length));
          (_1992_v100)[_index338] = _1990_cf3;
          let _index339 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1992_v100).length));
          let _rhs401 = !((p1) || (p1));
          let _rhs402 = _dafny.Seq.Concat(_1990_cf3, _1990_cf3);
          let _lhs277 = _1992_v100;
          let _lhs278 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1992_v100).length));
          _1989_cf4 = _rhs401;
          _lhs277[_lhs278] = _rhs402;
          (globalState).f7 = _module.__default.safeDivisionInt((((_1874_v10).contains(new BigNumber(982))) ? ((_1874_v10).get(new BigNumber(982))) : (p0)), (p0).minus(p0));
          r1 = ((_module.D16.create_DC42(_1989_cf4, false, true)).dtor_cf61) === (p2);
        } else if (_source27.is_DC4) {
          let _1993___mcc_h2 = (_source27).cf5;
          let _1994_cf5 = _1993___mcc_h2;
          let _1995_v101;
          _1995_v101 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(p2, p1, p1)).length),(_module.D8.create_DC23(_1875_v11)).dtor_cf39);
          r1 = (new BigNumber(120)).isEqualTo(_module.__default.fm17(new BigNumber((_1995_v101).length), globalState));
          let _1996_v102;
          let _nw304 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _1996_v102 = _nw304;
          let _index340 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1996_v102).length));
          (_1996_v102)[_index340] = new BigNumber((_1883_v19).length);
          let _index341 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1996_v102).length));
          (_1996_v102)[_index341] = p0;
          let _index342 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index342] = p2;
          let _index343 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index343] = (_1875_v11)[_module.__default.safeIndex(_module.__default.safeModuloInt(p0, (_1996_v102)[_module.__default.safeIndex(new BigNumber(566), new BigNumber((_1996_v102).length))]), new BigNumber((_1875_v11).length))];
          let _index344 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index344] = p2;
        } else {
          let _1997___mcc_h3 = (_source27).cf2;
          let _1998_cf2 = _1997___mcc_h3;
          let _1999_v103;
          let _nw305 = new _module.C2();
          _nw305.__ctor();
          _1999_v103 = _nw305;
          r1 = p2;
          let _index345 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index345] = _dafny.Seq.IsPrefixOf(_1883_v19, _1883_v19);
          let _index346 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1882_v18).length));
          let _rhs403 = !(false);
          let _rhs404 = p0;
          let _rhs405 = p1;
          let _rhs406 = !(_1998_cf2).isEqualTo(new BigNumber((_1879_v15).length));
          let _lhs279 = _1882_v18;
          let _lhs280 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1882_v18).length));
          r1 = _rhs403;
          _1998_cf2 = _rhs404;
          _lhs279[_lhs280] = _rhs405;
          r1 = _rhs406;
          let _index347 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_1882_v18).length));
          (_1882_v18)[_index347] = false;
        }
        r1 = p1;
        let _2000_v104;
        _2000_v104 = _dafny.Seq.of(_1880_v16);
        let _2001_v105;
        let _nw306 = Array((new BigNumber(18)).toNumber());
        _nw306[(_dafny.ZERO).toNumber()] = _1880_v16;
        _nw306[(_dafny.ONE).toNumber()] = (_1880_v16).Union(_1880_v16);
        _nw306[(new BigNumber(2)).toNumber()] = _module.__default.fm34(p0, new BigNumber((_1880_v16).cardinality()), p2, globalState);
        _nw306[(new BigNumber(3)).toNumber()] = (_1880_v16).Union(_1880_v16);
        _nw306[(new BigNumber(4)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(5)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(6)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(7)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(8)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(!(p2), !(p1));
        _nw306[(new BigNumber(10)).toNumber()] = (_1880_v16).Intersect(_1880_v16);
        _nw306[(new BigNumber(11)).toNumber()] = (_1880_v16).Intersect(_1880_v16);
        _nw306[(new BigNumber(12)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(13)).toNumber()] = (_1880_v16).Intersect(_1880_v16);
        _nw306[(new BigNumber(14)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(15)).toNumber()] = (_2000_v104)[_module.__default.safeIndex(p0, new BigNumber((_2000_v104).length))];
        _nw306[(new BigNumber(16)).toNumber()] = _1880_v16;
        _nw306[(new BigNumber(17)).toNumber()] = _1880_v16;
        _2001_v105 = _nw306;
        let _rhs407 = _dafny.Seq.Concat(_1883_v19, _dafny.Seq.Concat(_1883_v19, _1883_v19));
        let _rhs408 = _2001_v105;
        let _rhs409 = p0;
        let _lhs281 = globalState;
        _1883_v19 = _rhs407;
        r0 = _rhs408;
        _lhs281.f13 = _rhs409;
        let _2002_v106;
        let _nw307 = new _module.C4();
        _nw307.__ctor();
        _2002_v106 = _nw307;
        let _2003_v107;
        _2003_v107 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,_2002_v106),p2);
        _2003_v107 = _2003_v107;
        let _2004_v108;
        let _nw308 = Array((new BigNumber(13)).toNumber()).fill(_module.D21.Default());
        _2004_v108 = _nw308;
        let _2005_v109;
        let _nw309 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2005_v109 = _nw309;
        let _2006_v110;
        _2006_v110 = _module.D21.create_DC56(_2005_v109);
        let _2007_v111;
        _2007_v111 = _module.D21.create_DC56((_2006_v110).dtor_cf89);
        let _index348 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_2004_v108).length));
        (_2004_v108)[_index348] = _2007_v111;
        let _index349 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_2004_v108).length));
        (_2004_v108)[_index349] = _2006_v110;
      }
      let _2008_v112;
      let _nw310 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
      _2008_v112 = _nw310;
      r0 = _2008_v112;
      r1 = p1;
      r2 = _1877_v13;
      return [r0, r1, r2];
    }
  };

  $module.C11 = class C11 {
    constructor () {
      this._tname = "_module.C11";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return false;
    };
    fm2(p0, globalState) {
      let _this = this;
      return new BigNumber((function (_source28) {
        if (_source28.is_DC3) {
          let _2009___mcc_h0 = (_source28).cf3;
          let _2010___mcc_h1 = (_source28).cf4;
          let _2011_cf4 = _2010___mcc_h1;
          let _2012_cf3 = _2009___mcc_h0;
          return _2012_cf3;
        } else if (_source28.is_DC4) {
          let _2013___mcc_h2 = (_source28).cf5;
          let _2014_cf5 = _2013___mcc_h2;
          return _dafny.Seq.UnicodeFromString("rbfn");
        } else {
          let _2015___mcc_h3 = (_source28).cf2;
          let _2016_cf2 = _2015___mcc_h3;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("aredcecm"), _dafny.Seq.UnicodeFromString("lopn"));
        }
      }(((false) ? (_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("gjiblb"), false)) : (_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("ahql"), true))))).length);
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _2017_v0;
      _2017_v0 = _dafny.Seq.UnicodeFromString("w");
      _2017_v0 = _2017_v0;
      let _2018_i0;
      _2018_i0 = _dafny.ZERO;
      L13: {
        while (false) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2018_i0)) {
              break L13;
            }
            _2018_i0 = (_2018_i0).plus(_dafny.ONE);
            let _2019_v1;
            _2019_v1 = false;
            _2019_v1 = _2019_v1;
            if (true) {
              _2019_v1 = true;
              let _2020_v2;
              _2020_v2 = new _dafny.CodePoint('g'.codePointAt(0));
              let _2021_v3;
              _2021_v3 = new BigNumber(-854);
              let _rhs410 = ((_2019_v1) ? (_2021_v3) : (_2021_v3));
              let _rhs411 = _2020_v2;
              let _lhs282 = globalState;
              _lhs282.f7 = _rhs410;
              _2020_v2 = _rhs411;
              let _2022_v4;
              _2022_v4 = _dafny.Map.Empty.slice().updateUnsafe(_2021_v3,new BigNumber((_dafny.Set.fromElements(new BigNumber(-798))).length));
              let _2023_v5;
              _2023_v5 = _dafny.Seq.of(new BigNumber((_2022_v4).length), _2021_v3);
              let _2024_v6;
              _2024_v6 = _dafny.Set.fromElements(_2021_v3, new BigNumber(-516), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-427)), ((_2025_v2) => function (_2026_i1) {
                return _2025_v2;
              })(_2020_v2))).length));
              let _2027_v7;
              let _nw311 = Array((new BigNumber(5)).toNumber());
              _nw311[(_dafny.ZERO).toNumber()] = _2019_v1;
              _nw311[(_dafny.ONE).toNumber()] = _2019_v1;
              _nw311[(new BigNumber(2)).toNumber()] = _2019_v1;
              _nw311[(new BigNumber(3)).toNumber()] = _2019_v1;
              _nw311[(new BigNumber(4)).toNumber()] = _2019_v1;
              _2027_v7 = _nw311;
              let _2028_v8;
              _2028_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2027_v7,new BigNumber(950));
              let _2029_v9;
              _2029_v9 = _dafny.Seq.of(_2019_v1);
              let _2030_v10;
              let _nw312 = Array((new BigNumber(25)).toNumber());
              _nw312[(_dafny.ZERO).toNumber()] = _2021_v3;
              _nw312[(_dafny.ONE).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(2)).toNumber()] = new BigNumber(591);
              _nw312[(new BigNumber(3)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(4)).toNumber()] = new BigNumber((_2023_v5).length);
              _nw312[(new BigNumber(5)).toNumber()] = new BigNumber((_2024_v6).length);
              _nw312[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_2021_v3);
              _nw312[(new BigNumber(7)).toNumber()] = new BigNumber((_2028_v8).length);
              _nw312[(new BigNumber(8)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(9)).toNumber()] = new BigNumber(321);
              _nw312[(new BigNumber(10)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(11)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(12)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(13)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(14)).toNumber()] = new BigNumber((_2029_v9).length);
              _nw312[(new BigNumber(15)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(16)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(17)).toNumber()] = new BigNumber((p0).length);
              _nw312[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2023_v5).length))));
              _nw312[(new BigNumber(19)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(20)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(21)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(22)).toNumber()] = _2021_v3;
              _nw312[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(_2021_v3);
              _nw312[(new BigNumber(24)).toNumber()] = _2021_v3;
              _2030_v10 = _nw312;
              let _2031_v11;
              _2031_v11 = _dafny.Seq.of(_2030_v10);
              let _rhs412 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), _2017_v0), _dafny.Seq.update(_module.__default.fm5(globalState), _module.__default.safeIndex(new BigNumber((_2023_v5).length), new BigNumber((_module.__default.fm5(globalState)).length)), _2020_v2));
              let _rhs413 = _2021_v3;
              let _rhs414 = _2031_v11;
              let _lhs283 = globalState;
              _2019_v1 = _rhs412;
              _lhs283.f15 = _rhs413;
              _2031_v11 = _rhs414;
              _2017_v0 = p0;
              (globalState).f13 = new BigNumber(-597);
            } else {
              let _2032_v12;
              let _nw313 = Array((new BigNumber(21)).toNumber()).fill(_module.D1.Default());
              _2032_v12 = _nw313;
              let _2033_v13;
              _2033_v13 = _module.D1.create_DC3(p0, true);
              let _index350 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2032_v12).length));
              (_2032_v12)[_index350] = _2033_v13;
              let _index351 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_2032_v12).length));
              (_2032_v12)[_index351] = _2033_v13;
              _2019_v1 = _2019_v1;
              let _2034_v14;
              _2034_v14 = new BigNumber(586);
              let _2035_v15;
              _2035_v15 = _dafny.Map.Empty.slice().updateUnsafe(_2034_v14,_2034_v14);
              (globalState).f13 = (_this).fm2((_2034_v14).plus(new BigNumber((_2035_v15).length)), globalState);
              let _2036_v16;
              _2036_v16 = new _dafny.CodePoint('w'.codePointAt(0));
              let _2037_v17;
              _2037_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2036_v16,p0);
              _2037_v17 = ((_2037_v17).Merge(_2037_v17)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2036_v16,_2017_v0));
              (globalState).f4 = _2034_v14;
            }
            let _2038_v18;
            let _nw314 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _2038_v18 = _nw314;
            let _index352 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2038_v18).length));
            (_2038_v18)[_index352] = new _dafny.CodePoint('l'.codePointAt(0));
            let _2039_v19;
            _2039_v19 = new _dafny.CodePoint('m'.codePointAt(0));
            let _index353 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2038_v18).length));
            (_2038_v18)[_index353] = _2039_v19;
            let _2040_v20;
            _2040_v20 = new BigNumber(534);
            let _2041_v21;
            _2041_v21 = _dafny.Seq.of(new BigNumber(-358), _2040_v20);
            let _2042_v22;
            _2042_v22 = _dafny.MultiSet.fromElements(_2040_v20, _2040_v20, (_dafny.ZERO).minus(_2040_v20), new BigNumber((_2041_v21).length), new BigNumber((_dafny.Seq.of(_2040_v20, _2040_v20)).length));
            let _2043_v23;
            _2043_v23 = _dafny.MultiSet.fromElements(_2019_v1, _2019_v1);
            let _index354 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2038_v18).length));
            let _rhs415 = (_dafny.ZERO).minus(new BigNumber(((_2042_v22).update(_2040_v20, _module.__default.abs(new BigNumber((_2043_v23).cardinality())))).cardinality()));
            let _rhs416 = _2039_v19;
            let _rhs417 = (_2040_v20).plus((_dafny.ZERO).minus(_2040_v20));
            let _rhs418 = (_this).fm2((_2040_v20).multipliedBy(_2040_v20), globalState);
            let _rhs419 = (_2040_v20).isLessThanOrEqualTo(new BigNumber((_2041_v21).length));
            let _lhs284 = globalState;
            let _lhs285 = _2038_v18;
            let _lhs286 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2038_v18).length));
            let _lhs287 = globalState;
            let _lhs288 = globalState;
            _lhs284.f13 = _rhs415;
            _lhs285[_lhs286] = _rhs416;
            _lhs287.f7 = _rhs417;
            _lhs288.f13 = _rhs418;
            _2019_v1 = _rhs419;
          }
        }
      }
      let _2044_v24;
      _2044_v24 = true;
      let _2045_v25;
      _2045_v25 = new BigNumber(261);
      let _2046_v26;
      _2046_v26 = _dafny.MultiSet.fromElements(_2045_v25);
      let _2047_v27;
      _2047_v27 = _module.D1.create_DC4(_2046_v26);
      let _source29 = ((_2044_v24) ? (_2047_v27) : (_2047_v27));
      if (_source29.is_DC3) {
        let _2048___mcc_h0 = (_source29).cf3;
        let _2049___mcc_h1 = (_source29).cf4;
        let _2050_cf4 = _2049___mcc_h1;
        let _2051_cf3 = _2048___mcc_h0;
        let _2052_v28;
        _2052_v28 = _module.D1.create_DC2(new BigNumber((_2046_v26).cardinality()));
        (globalState).f7 = (((_2050_cf4) ? (_module.D1.create_DC2(new BigNumber(495))) : (_2052_v28))).dtor_cf2;
        let _rhs420 = _2045_v25;
        let _rhs421 = _2050_cf4;
        let _lhs289 = globalState;
        _lhs289.f4 = _rhs420;
        _2044_v24 = _rhs421;
        _2050_cf4 = _2050_cf4;
        let _2053_v29;
        let _nw315 = new _module.C10();
        _nw315.__ctor();
        _2053_v29 = _nw315;
        _2053_v29 = _2053_v29;
      } else if (_source29.is_DC4) {
        let _2054___mcc_h2 = (_source29).cf5;
        let _2055_cf5 = _2054___mcc_h2;
        let _2056_v30;
        let _nw316 = Array((new BigNumber(7)).toNumber());
        _nw316[(_dafny.ZERO).toNumber()] = _2044_v24;
        _nw316[(_dafny.ONE).toNumber()] = false;
        _nw316[(new BigNumber(2)).toNumber()] = _2044_v24;
        _nw316[(new BigNumber(3)).toNumber()] = _2044_v24;
        _nw316[(new BigNumber(4)).toNumber()] = _2044_v24;
        _nw316[(new BigNumber(5)).toNumber()] = _2044_v24;
        _nw316[(new BigNumber(6)).toNumber()] = _2044_v24;
        _2056_v30 = _nw316;
        let _2057_v31;
        _2057_v31 = _dafny.Seq.of(_2056_v30, _2056_v30);
        let _2058_v32;
        _2058_v32 = _module.D7.create_DC19(_2057_v31);
        let _2059_v33;
        _2059_v33 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_2058_v32);
        let _2060_v34;
        _2060_v34 = _dafny.Set.fromElements(_2045_v25);
        let _2061_v35;
        let _nw317 = Array((new BigNumber(19)).toNumber());
        _nw317[(_dafny.ZERO).toNumber()] = new BigNumber(223);
        _nw317[(_dafny.ONE).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(2)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(3)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(4)).toNumber()] = (_this).fm2(new BigNumber(704), globalState);
        _nw317[(new BigNumber(5)).toNumber()] = new BigNumber((_2059_v33).length);
        _nw317[(new BigNumber(6)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(7)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(8)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(9)).toNumber()] = new BigNumber(447);
        _nw317[(new BigNumber(10)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(11)).toNumber()] = new BigNumber((_2060_v34).length);
        _nw317[(new BigNumber(12)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(13)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(14)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(15)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(16)).toNumber()] = _2045_v25;
        _nw317[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("hvirhay")).length);
        _nw317[(new BigNumber(18)).toNumber()] = _2045_v25;
        _2061_v35 = _nw317;
        let _2062_v36;
        _2062_v36 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber(918)).multipliedBy(_2045_v25)),_2061_v35);
        _2062_v36 = (_2062_v36).update((_dafny.ZERO).minus(_2045_v25), _2061_v35);
        let _2063_v37;
        let _nw318 = new _module.C9();
        _nw318.__ctor();
        _2063_v37 = _nw318;
        let _2064_v38;
        _2064_v38 = new _dafny.CodePoint('x'.codePointAt(0));
        let _2065_v39;
        _2065_v39 = _dafny.Set.fromElements(_2064_v38, _2064_v38, _2064_v38, _module.__default.fm38(_2044_v24, globalState), new _dafny.CodePoint('a'.codePointAt(0)));
        let _2066_v41;
        _2066_v41 = _dafny.Seq.of((_2065_v39).Union(function () {
          let _coll79 = new _dafny.Set();
          for (const _compr_79 of (_2065_v39).Elements) {
            let _2067_v40 = _compr_79;
            if ((_2065_v39).contains(_2067_v40)) {
              _coll79.add(_2067_v40);
            }
          }
          return _coll79;
        }()));
        _2065_v39 = (_2066_v41)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_2045_v25, _2045_v25), new BigNumber((_2066_v41).length))];
        _2045_v25 = _2045_v25;
      } else {
        let _2068___mcc_h3 = (_source29).cf2;
        let _2069_cf2 = _2068___mcc_h3;
        let _2070_v42;
        let _nw319 = Array((new BigNumber(25)).toNumber()).fill([]);
        _2070_v42 = _nw319;
        let _2071_v43;
        let _nw320 = Array((new BigNumber(5)).toNumber()).fill(false);
        _2071_v43 = _nw320;
        let _index355 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length));
        (_2070_v42)[_index355] = _2071_v43;
        let _index356 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length));
        let _rhs422 = _2069_cf2;
        let _rhs423 = _2071_v43;
        let _rhs424 = true;
        let _lhs290 = _2070_v42;
        let _lhs291 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length));
        _2045_v25 = _rhs422;
        _lhs290[_lhs291] = _rhs423;
        _2044_v24 = _rhs424;
        if (_2044_v24) {
          let _2072_v44;
          _2072_v44 = _dafny.Map.Empty.slice().updateUnsafe(!(_2044_v24) || (_2044_v24),_module.__default.fm5(globalState));
          let _2073_v45;
          _2073_v45 = _module.D1.create_DC3(_2017_v0, _2044_v24);
          _2072_v44 = (_2072_v44).update(_2044_v24, _dafny.Seq.Concat((_2073_v45).dtor_cf3, _2017_v0));
          let _2074_v46;
          _2074_v46 = _dafny.Seq.of(_2044_v24);
          let _2075_v47;
          _2075_v47 = _dafny.Map.Empty.slice().updateUnsafe(_2074_v46,_2044_v24);
          _2075_v47 = (_2075_v47).Merge(_2075_v47);
          let _2076_v48;
          _2076_v48 = _dafny.Seq.of(_2046_v26, _2046_v26);
          let _2077_v49;
          _2077_v49 = _dafny.Seq.of((_2076_v48)[_module.__default.safeIndex(_2045_v25, new BigNumber((_2076_v48).length))]);
          (globalState).f4 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2076_v48, _2076_v48), _2076_v48)).length);
          let _2078_v50;
          let _nw321 = new _module.C10();
          _nw321.__ctor();
          _2078_v50 = _nw321;
          _2078_v50 = _2078_v50;
          (globalState).f15 = _2069_cf2;
        } else {
          let _2079_v51;
          let _init52 = ((_2080_v24) => function (_2081_i2) {
            return _dafny.Set.fromElements(_2080_v24);
          })(_2044_v24);
          let _nw322 = Array((new BigNumber(15)).toNumber());
          for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw322.length); _i0_52++) {
            _nw322[_i0_52] = _init52(new BigNumber(_i0_52));
          }
          _2079_v51 = _nw322;
          _2079_v51 = _2079_v51;
          let _arr2 = (_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))];
          let _index357 = _module.__default.safeIndex(new BigNumber(335), new BigNumber(((_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))]).length));
          _arr2[_index357] = false;
          let _2082_v52;
          let _nw323 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _2082_v52 = _nw323;
          let _2083_v53;
          _2083_v53 = _dafny.Set.fromElements(_2082_v52, _2082_v52);
          let _2084_v54;
          _2084_v54 = _dafny.Seq.of(new BigNumber((_2083_v53).length));
          let _arr3 = (_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))];
          let _index358 = _module.__default.safeIndex(new BigNumber(335), new BigNumber(((_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))]).length));
          _arr3[_index358] = ((_2084_v54)[_module.__default.safeIndex(new BigNumber(984), new BigNumber((_2084_v54).length))]).isLessThan((_2069_cf2).multipliedBy(_2069_cf2));
          let _2085_v55;
          _2085_v55 = new _dafny.CodePoint('i'.codePointAt(0));
          let _2086_v56;
          _2086_v56 = _dafny.MultiSet.fromElements(_2085_v55);
          let _arr4 = (_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))];
          let _index359 = _module.__default.safeIndex(new BigNumber(335), new BigNumber(((_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))]).length));
          _arr4[_index359] = (_2086_v56).IsDisjointFrom(_2086_v56);
          (globalState).f13 = _module.__default.fm17((_dafny.ZERO).minus(_2045_v25), globalState);
          let _arr5 = (_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))];
          let _index360 = _module.__default.safeIndex(new BigNumber(335), new BigNumber(((_2070_v42)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2070_v42).length))]).length));
          _arr5[_index360] = _2044_v24;
        }
        let _2087_v57;
        _2087_v57 = _dafny.Map.Empty.slice().updateUnsafe(_2044_v24,_2044_v24);
        let _2088_v58;
        _2088_v58 = _dafny.Map.Empty.slice().updateUnsafe((((_2087_v57).contains(!(_2044_v24))) ? ((_2087_v57).get(!(_2044_v24))) : (_2044_v24)),_2087_v57);
        let _2089_v59;
        _2089_v59 = _dafny.Seq.of((_this).fm1(_2088_v58, _2069_cf2, globalState));
        let _2090_v60;
        _2090_v60 = _dafny.Map.Empty.slice().updateUnsafe(_2089_v59,_2069_cf2);
        let _2091_v61;
        _2091_v61 = new _dafny.CodePoint('o'.codePointAt(0));
        let _2092_v62;
        _2092_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2091_v61,_2046_v26);
        let _2093_v64;
        _2093_v64 = _dafny.Seq.of(_2092_v62, function () {
          let _coll80 = new _dafny.Map();
          for (const _compr_80 of (_2017_v0).Elements) {
            let _2094_v63 = _compr_80;
            if (_dafny.Seq.contains(_2017_v0, _2094_v63)) {
              _coll80.push([_2094_v63,_2046_v26]);
            }
          }
          return _coll80;
        }(), _2092_v62);
        let _2095_v65;
        _2095_v65 = _dafny.Seq.of(_2069_cf2, (new BigNumber((_2090_v60).length)).minus(_module.__default.fm17(_2045_v25, globalState)), new BigNumber(((_2093_v64)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_2044_v24, _2044_v24)).length), new BigNumber((_2093_v64).length))]).length));
        _2095_v65 = _2095_v65;
        _2044_v24 = (new BigNumber(678)).isLessThanOrEqualTo((_dafny.ZERO).minus(_2069_cf2));
      }
      let _2096_v66;
      let _nw324 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _2096_v66 = _nw324;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2096_v66).length))) {
        let _2097_i3 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2097_i3)) && ((_2097_i3).isLessThan(new BigNumber((_2096_v66).length))))) {
          (_2096_v66)[(_2097_i3)] = (_2097_i3).minus((_2045_v25).plus(_2045_v25));
        }
      }
      let _2098_v67;
      let _nw325 = Array((new BigNumber(4)).toNumber()).fill(false);
      _2098_v67 = _nw325;
      let _2099_v68;
      _2099_v68 = _module.D0.create_DC0(_2098_v67);
      let _pat_let_tv42 = _2098_v67;
      let _pat_let_tv43 = _2098_v67;
      let _2100_v69;
      let _nw326 = Array((new BigNumber(5)).toNumber());
      _nw326[(_dafny.ZERO).toNumber()] = function (_pat_let35_0) {
        return function (_2101_dt__update__tmp_h0) {
          return function (_pat_let36_0) {
            return function (_2102_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_2102_dt__update_hcf0_h0);
            }(_pat_let36_0);
          }(_pat_let_tv42);
        }(_pat_let35_0);
      }(_2099_v68);
      _nw326[(_dafny.ONE).toNumber()] = function (_pat_let37_0) {
        return function (_2103_dt__update__tmp_h1) {
          return function (_pat_let38_0) {
            return function (_2104_dt__update_hcf0_h1) {
              return _module.D0.create_DC0(_2104_dt__update_hcf0_h1);
            }(_pat_let38_0);
          }(_pat_let_tv43);
        }(_pat_let37_0);
      }(_2099_v68);
      _nw326[(new BigNumber(2)).toNumber()] = _2099_v68;
      _nw326[(new BigNumber(3)).toNumber()] = _2099_v68;
      _nw326[(new BigNumber(4)).toNumber()] = _2099_v68;
      _2100_v69 = _nw326;
      let _index361 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_2100_v69).length));
      (_2100_v69)[_index361] = _module.D0.create_DC0(_2098_v67);
      let _2105_v70;
      let _nw327 = Array((new BigNumber(7)).toNumber()).fill([]);
      _2105_v70 = _nw327;
      let _index362 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2105_v70).length));
      (_2105_v70)[_index362] = _2096_v66;
      let _2106_v71;
      _2106_v71 = _module.D6.create_DC17(_2044_v24, !(_2044_v24), _2044_v24);
      let _2107_v72;
      _2107_v72 = _dafny.Seq.of(_2106_v71, _2106_v71);
      let _index363 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2096_v66).length));
      (_2096_v66)[_index363] = new BigNumber((_2107_v72).length);
      let _2108_v73;
      _2108_v73 = _dafny.Map.Empty.slice().updateUnsafe(((_2044_v24) ? (_2096_v66) : (_2096_v66)),_2045_v25);
      let _2109_v74;
      _2109_v74 = _dafny.MultiSet.fromElements(_2044_v24, _2044_v24);
      let _index364 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_2100_v69).length));
      let _index365 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2105_v70).length));
      let _index366 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2096_v66).length));
      let _rhs425 = _2099_v68;
      let _rhs426 = _2096_v66;
      let _rhs427 = _module.__default.safeDivisionInt(((_2044_v24) ? (_2045_v25) : ((_dafny.ZERO).minus(_2045_v25))), _module.__default.safeModuloInt(new BigNumber((_2109_v74).cardinality()), _2045_v25));
      let _rhs428 = _2108_v73;
      let _lhs292 = _2100_v69;
      let _lhs293 = _module.__default.safeIndex(new BigNumber(835), new BigNumber((_2100_v69).length));
      let _lhs294 = _2105_v70;
      let _lhs295 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_2105_v70).length));
      let _lhs296 = _2096_v66;
      let _lhs297 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_2096_v66).length));
      _lhs292[_lhs293] = _rhs425;
      _lhs294[_lhs295] = _rhs426;
      _lhs296[_lhs297] = _rhs427;
      _2108_v73 = _rhs428;
      let _2110_v75;
      let _2111_v76;
      let _2112_v77;
      let _2113_v78;
      let _out15;
      let _out16;
      let _out17;
      let _out18;
      let _outcollector6 = (_this).m3(globalState);
      _out15 = _outcollector6[0];
      _out16 = _outcollector6[1];
      _out17 = _outcollector6[2];
      _out18 = _outcollector6[3];
      _2110_v75 = _out15;
      _2111_v76 = _out16;
      _2112_v77 = _out17;
      _2113_v78 = _out18;
      return;
    }
    m3(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      let r3 = [];
      let _2114_v0;
      let _nw328 = new _module.C9();
      _nw328.__ctor();
      _2114_v0 = _nw328;
      let _2115_v1;
      _2115_v1 = new BigNumber(196);
      let _2116_v2;
      _2116_v2 = _module.D4.create_DC12((_dafny.ZERO).minus(_2115_v1));
      let _2117_v3;
      _2117_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v1,(_module.__default.fm56((_dafny.ZERO).minus(_2115_v1), (_2116_v2).dtor_cf16, globalState)).dtor_cf35);
      _2117_v3 = (_2117_v3).update(_2115_v1, true);
      let _2118_v4;
      _2118_v4 = false;
      let _2119_v5;
      _2119_v5 = _module.D20.create_DC55((_dafny.ZERO).minus(_2115_v1), _2115_v1, _2118_v4, !(_2118_v4));
      _2118_v4 = (_2119_v5).dtor_cf87;
      let _2120_v6;
      _2120_v6 = _dafny.Seq.UnicodeFromString("mxqw");
      let _2121_v7;
      _2121_v7 = _dafny.Map.Empty.slice().updateUnsafe(_2118_v4,_2115_v1);
      let _2122_v8;
      _2122_v8 = _dafny.Seq.of(_2118_v4, (_2114_v0).fm8(new BigNumber((_2121_v7).length), _2118_v4, new BigNumber(887), globalState));
      let _2123_v9;
      _2123_v9 = _dafny.Set.fromElements(_2115_v1);
      let _2124_v10;
      _2124_v10 = _module.D1.create_DC3(_module.__default.fm30(new BigNumber(557), !((_2122_v8)[_module.__default.safeIndex(new BigNumber((_module.__default.fm5(globalState)).length), new BigNumber((_2122_v8).length))]), _2123_v9, false, globalState), false);
      _2120_v6 = (_2124_v10).dtor_cf3;
      let _2125_v11;
      _2125_v11 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v1,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-67)), function (_2126_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _2120_v6)).length));
      _2125_v11 = (_2125_v11).update(_2115_v1, _2115_v1);
      let _2127_v12;
      let _nw329 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _2127_v12 = _nw329;
      let _2128_v13;
      _2128_v13 = _dafny.MultiSet.fromElements(_2118_v4, _2118_v4);
      let _2129_v14;
      _2129_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2128_v13).cardinality()),_2127_v12);
      let _2130_v15;
      let _nw330 = Array((new BigNumber(12)).toNumber());
      _nw330[(_dafny.ZERO).toNumber()] = _2127_v12;
      _nw330[(_dafny.ONE).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(2)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(3)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(4)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(5)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(6)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(7)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(8)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(9)).toNumber()] = (((_2129_v14).contains(new BigNumber(456))) ? ((_2129_v14).get(new BigNumber(456))) : (_2127_v12));
      _nw330[(new BigNumber(10)).toNumber()] = _2127_v12;
      _nw330[(new BigNumber(11)).toNumber()] = _2127_v12;
      _2130_v15 = _nw330;
      let _2131_v16;
      _2131_v16 = _dafny.Seq.of(_2130_v15, _2130_v15, _2130_v15, _2130_v15);
      let _2132_v17;
      let _nw331 = Array((new BigNumber(25)).toNumber());
      _nw331[(_dafny.ZERO).toNumber()] = _2130_v15;
      _nw331[(_dafny.ONE).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(2)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(3)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(4)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(5)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(6)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(7)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(8)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(9)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(10)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(11)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(12)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(13)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(14)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(15)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(16)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(17)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(18)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(19)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(20)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(21)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(22)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(23)).toNumber()] = _2130_v15;
      _nw331[(new BigNumber(24)).toNumber()] = (_2131_v16)[_module.__default.safeIndex((_dafny.ZERO).minus(_2115_v1), new BigNumber((_2131_v16).length))];
      _2132_v17 = _nw331;
      let _index367 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2132_v17).length));
      (_2132_v17)[_index367] = _2130_v15;
      let _index368 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_2132_v17).length));
      (_2132_v17)[_index368] = _2130_v15;
      let _2133_v18;
      let _init53 = function (_2134_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      };
      let _nw332 = Array((new BigNumber(28)).toNumber());
      for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw332.length); _i0_53++) {
        _nw332[_i0_53] = _init53(new BigNumber(_i0_53));
      }
      _2133_v18 = _nw332;
      let _2135_v19;
      _2135_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(_2115_v1, _2118_v4, _2123_v9, !(_2118_v4), globalState),_2133_v18);
      r0 = new BigNumber(((_2135_v19).Merge((_2135_v19).Merge(_2135_v19))).length);
      let _2136_v20;
      _2136_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(_2115_v1, _2115_v1)),_2118_v4);
      let _2137_v21;
      _2137_v21 = _dafny.Seq.of(_2136_v20, _2136_v20);
      r1 = (_2137_v21)[_module.__default.safeIndex((new BigNumber(166)).minus(_2115_v1), new BigNumber((_2137_v21).length))];
      r2 = _2115_v1;
      let _2138_v22;
      let _init54 = ((_2139_v7) => function (_2140_i2) {
        return _2139_v7;
      })(_2121_v7);
      let _nw333 = Array((new BigNumber(23)).toNumber());
      for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw333.length); _i0_54++) {
        _nw333[_i0_54] = _init54(new BigNumber(_i0_54));
      }
      _2138_v22 = _nw333;
      r3 = _2138_v22;
      return [r0, r1, r2, r3];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let _2141_i0;
      _2141_i0 = _dafny.ZERO;
      L14: {
        while (p3) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2141_i0)) {
              break L14;
            }
            _2141_i0 = (_2141_i0).plus(_dafny.ONE);
            let _2142_v0;
            _2142_v0 = true;
            let _2143_v1;
            _2143_v1 = new _dafny.CodePoint('w'.codePointAt(0));
            let _2144_v2;
            _2144_v2 = _module.D5.create_DC14(_2143_v1, p1, p0, p3);
            _2142_v0 = (p1) && ((_2144_v2).dtor_cf21);
            let _2145_v3;
            let _2146_v4;
            let _2147_v5;
            let _2148_v6;
            let _out19;
            let _out20;
            let _out21;
            let _out22;
            let _outcollector7 = (_this).m3(globalState);
            _out19 = _outcollector7[0];
            _out20 = _outcollector7[1];
            _out21 = _outcollector7[2];
            _out22 = _outcollector7[3];
            _2145_v3 = _out19;
            _2146_v4 = _out20;
            _2147_v5 = _out21;
            _2148_v6 = _out22;
            if (true) {
              (globalState).f13 = _2147_v5;
              (globalState).f4 = (_2147_v5).minus((_this).fm2(_2147_v5, globalState));
              let _2149_v7;
              _2149_v7 = _dafny.Seq.UnicodeFromString("sxs");
              let _2150_v8;
              _2150_v8 = _dafny.Map.Empty.slice().updateUnsafe(_2147_v5,new BigNumber(327));
              let _rhs429 = _module.__default.fm17(p0, globalState);
              let _rhs430 = _dafny.Seq.UnicodeFromString("wse");
              let _rhs431 = (_2145_v3).plus((new BigNumber(723)).multipliedBy(_module.__default.fm17(_2145_v3, globalState)));
              let _rhs432 = ((_dafny.ZERO).minus(p0)).multipliedBy(p0);
              let _rhs433 = _2150_v8;
              let _lhs298 = globalState;
              let _lhs299 = globalState;
              let _lhs300 = globalState;
              _lhs298.f4 = _rhs429;
              _2149_v7 = _rhs430;
              _lhs299.f13 = _rhs431;
              _lhs300.f13 = _rhs432;
              _2150_v8 = _rhs433;
              let _2151_v9;
              let _nw334 = new _module.C0();
              _nw334.__ctor();
              _2151_v9 = _nw334;
              let _2152_v10;
              _2152_v10 = _dafny.Seq.of(p1, _2142_v0, p1);
              let _2153_v11;
              _2153_v11 = _dafny.Seq.of(new BigNumber((_2152_v10).length));
              let _2154_v12;
              _2154_v12 = _dafny.Seq.of(_2153_v11);
              let _2155_v13;
              _2155_v13 = _dafny.Seq.of(_2149_v7, _2149_v7);
              let _2156_v14;
              let _nw335 = Array((new BigNumber(21)).toNumber());
              _nw335[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("gif");
              _nw335[(_dafny.ONE).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(2)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_module.__default.fm5(globalState), _module.__default.safeIndex(_2147_v5, new BigNumber((_module.__default.fm5(globalState)).length)), _2143_v1);
              _nw335[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-717)), ((_2157_v1) => function (_2158_i1) {
                return _2157_v1;
              })(_2143_v1));
              _nw335[(new BigNumber(5)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(343)), ((_2159_v1) => function (_2160_i2) {
                return _2159_v1;
              })(_2143_v1));
              _nw335[(new BigNumber(7)).toNumber()] = (_2151_v9).fm28(_2154_v12, globalState);
              _nw335[(new BigNumber(8)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_2149_v7, _module.__default.safeIndex(_2145_v3, new BigNumber((_2149_v7).length)), new _dafny.CodePoint('n'.codePointAt(0))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_2149_v7, _module.__default.safeIndex(_2145_v3, new BigNumber((_2149_v7).length)), new _dafny.CodePoint('n'.codePointAt(0)))).length)), new _dafny.CodePoint('x'.codePointAt(0)));
              _nw335[(new BigNumber(10)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(11)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(12)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(13)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("jc");
              _nw335[(new BigNumber(15)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(16)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(17)).toNumber()] = _2149_v7;
              _nw335[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("befkyx");
              _nw335[(new BigNumber(19)).toNumber()] = (_2155_v13)[_module.__default.safeIndex(_2147_v5, new BigNumber((_2155_v13).length))];
              _nw335[(new BigNumber(20)).toNumber()] = _2149_v7;
              _2156_v14 = _nw335;
              let _2161_v15;
              let _nw336 = Array((new BigNumber(10)).toNumber()).fill(false);
              _2161_v15 = _nw336;
              let _2162_v16;
              _2162_v16 = _dafny.Seq.of(_2161_v15, _2161_v15);
              let _2163_v17;
              _2163_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2156_v14,_2162_v16);
              _2163_v17 = (_2163_v17).update(_2156_v14, _dafny.Seq.Concat(_2162_v16, _2162_v16));
            } else {
              let _2164_v18;
              _2164_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_2147_v5)),_2147_v5);
              (globalState).f4 = (((_2164_v18).contains((_2147_v5).plus(_2145_v3))) ? ((_2164_v18).get((_2147_v5).plus(_2145_v3))) : (p0));
              let _2165_v19;
              _2165_v19 = _dafny.Set.fromElements(!(p1));
              let _2166_v20;
              let _nw337 = Array((new BigNumber(8)).toNumber());
              _nw337[(_dafny.ZERO).toNumber()] = _2142_v0;
              _nw337[(_dafny.ONE).toNumber()] = _2142_v0;
              _nw337[(new BigNumber(2)).toNumber()] = p1;
              _nw337[(new BigNumber(3)).toNumber()] = (_2165_v19).IsDisjointFrom(_2165_v19);
              _nw337[(new BigNumber(4)).toNumber()] = _2142_v0;
              _nw337[(new BigNumber(5)).toNumber()] = p3;
              _nw337[(new BigNumber(6)).toNumber()] = p3;
              _nw337[(new BigNumber(7)).toNumber()] = p1;
              _2166_v20 = _nw337;
              let _index369 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2166_v20).length));
              (_2166_v20)[_index369] = _2142_v0;
              let _index370 = _module.__default.safeIndex(new BigNumber(320), new BigNumber((_2166_v20).length));
              (_2166_v20)[_index370] = p1;
              let _2167_v21;
              _2167_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2142_v0,!(p3));
              _2142_v0 = !(!((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(p1,_2167_v21), p0, globalState)));
              _2145_v3 = p0;
              let _2168_v23;
              _2168_v23 = _dafny.Seq.of(_2147_v5);
              _2164_v18 = ((_2164_v18).Merge(function () {
                let _coll81 = new _dafny.Map();
                for (const _compr_81 of (_2168_v23).Elements) {
                  let _2169_v22 = _compr_81;
                  if (_dafny.Seq.contains(_2168_v23, _2169_v22)) {
                    _coll81.push([(_2169_v22).plus(_2147_v5),_2145_v3]);
                  }
                }
                return _coll81;
              }())).Merge(_2164_v18);
            }
            let _2170_v24;
            let _nw338 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2170_v24 = _nw338;
            let _2171_v25;
            _2171_v25 = _dafny.Seq.UnicodeFromString("ihvfprud");
            let _index371 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_2170_v24).length));
            (_2170_v24)[_index371] = _2171_v25;
            let _index372 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_2170_v24).length));
            (_2170_v24)[_index372] = _dafny.Seq.Concat(_2171_v25, _2171_v25);
          }
        }
      }
      let _2172_v26;
      _2172_v26 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
      let _2173_v27;
      _2173_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2172_v26);
      let _2174_v28;
      _2174_v28 = _dafny.Seq.of((_dafny.ZERO).minus(p0));
      let _2175_v29;
      _2175_v29 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_this).fm1(_2173_v27, new BigNumber((_2174_v28).length), globalState));
      let _2176_v30;
      _2176_v30 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2172_v26);
      let _2177_v31;
      _2177_v31 = _dafny.Seq.of(p1);
      let _2178_v32;
      _2178_v32 = _dafny.MultiSet.fromElements(_2177_v31);
      let _2179_v33;
      _2179_v33 = _dafny.Set.fromElements((_this).fm1(_2173_v27, (_this).fm2(new BigNumber((_2178_v32).cardinality()), globalState), globalState));
      let _2180_v34;
      _2180_v34 = new _dafny.CodePoint('l'.codePointAt(0));
      let _2181_v35;
      _2181_v35 = _dafny.Map.Empty.slice().updateUnsafe(_2180_v34,_dafny.Set.fromElements(p1));
      let _pat_let_tv44 = _2181_v35;
      let _pat_let_tv45 = _2180_v34;
      let _pat_let_tv46 = _2179_v33;
      let _pat_let_tv47 = _2181_v35;
      let _pat_let_tv48 = _2180_v34;
      let _source30 = function (_pat_let39_0) {
        return function (_2182_dt__update__tmp_h0) {
          return function (_pat_let40_0) {
            return function (_2183_dt__update_hcf49_h0) {
              return _module.D12.create_DC32(_2183_dt__update_hcf49_h0);
            }(_pat_let40_0);
          }((((_pat_let_tv47).contains(_pat_let_tv48)) ? ((_pat_let_tv44).get(_pat_let_tv45)) : (_pat_let_tv46)));
        }(_pat_let39_0);
      }(_module.D12.create_DC32(_2179_v33));
      if (_source30.is_DC33) {
        let _2184___mcc_h0 = (_source30).cf50;
        let _2185_cf50 = _2184___mcc_h0;
        (globalState).f13 = new BigNumber(461);
        let _2186_v36;
        _2186_v36 = _module.D3.create_DC10(new BigNumber(468), p0);
        let _source31 = _2186_v36;
        if (_source31.is_DC10) {
          let _2187___mcc_h2 = (_source31).cf13;
          let _2188___mcc_h3 = (_source31).cf14;
          let _2189_cf14 = _2188___mcc_h3;
          let _2190_cf13 = _2187___mcc_h2;
          (globalState).f4 = p0;
          let _2191_v37;
          _2191_v37 = _dafny.Seq.UnicodeFromString("xmyrhykkq");
          let _2192_v38;
          _2192_v38 = _dafny.Map.Empty.slice().updateUnsafe(_2189_cf14,false);
          let _2193_v39;
          let _init55 = ((_2194_cf14) => function (_2195_i3) {
            return (_2195_i3).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), ((_2196_cf14) => function (_2197_i4) {
              return _2196_cf14;
            })(_2194_cf14))).length));
          })(_2189_cf14);
          let _nw339 = Array((new BigNumber(7)).toNumber());
          for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw339.length); _i0_55++) {
            _nw339[_i0_55] = _init55(new BigNumber(_i0_55));
          }
          _2193_v39 = _nw339;
          let _2198_v40;
          _2198_v40 = _module.D2.create_DC5(_2193_v39);
          let _2199_v41;
          _2199_v41 = _dafny.Set.fromElements(_2198_v40);
          let _rhs434 = _module.__default.fm26(globalState);
          let _rhs435 = !(((p1) ? (_2185_cf50) : (!(p3)))) || ((((_2192_v38).contains(new BigNumber((_2179_v33).length))) ? ((_2192_v38).get(new BigNumber((_2179_v33).length))) : (p3)));
          let _rhs436 = !_dafny.Seq.contains(_module.__default.fm19((_2177_v31)[_module.__default.safeIndex(new BigNumber((_2199_v41).length), new BigNumber((_2177_v31).length))], _2185_cf50, globalState), (((_2177_v31)[_module.__default.safeIndex(new BigNumber(29), new BigNumber((_2177_v31).length))]) ? (_2180_v34) : (_2180_v34)));
          let _rhs437 = _2172_v26;
          let _rhs438 = _dafny.Seq.Concat(_2191_v37, _2191_v37);
          _2185_cf50 = _rhs434;
          _2185_cf50 = _rhs435;
          _2185_cf50 = _rhs436;
          _2172_v26 = _rhs437;
          _2191_v37 = _rhs438;
          let _2200_v42;
          let _nw340 = Array((new BigNumber(2)).toNumber()).fill([]);
          _2200_v42 = _nw340;
          let _rhs439 = _2200_v42;
          let _rhs440 = ((p3) ? (!(p1) || (p3)) : (_2185_cf50));
          let _rhs441 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-409)), ((_2201_cf13, _2202_v38) => function (_2203_i5) {
            return (_2201_cf13).minus(new BigNumber((_2202_v38).length));
          })(_2190_cf13, _2192_v38));
          let _rhs442 = _2180_v34;
          let _rhs443 = p3;
          _2200_v42 = _rhs439;
          _2185_cf50 = _rhs440;
          _2174_v28 = _rhs441;
          _2180_v34 = _rhs442;
          _2185_cf50 = _rhs443;
          _2191_v37 = _2191_v37;
        } else {
          let _2204___mcc_h4 = (_source31).cf12;
          let _2205_cf12 = _2204___mcc_h4;
          _2185_cf50 = _module.__default.fm26(globalState);
          let _2206_v43;
          let _nw341 = new _module.C1();
          _nw341.__ctor();
          _2206_v43 = _nw341;
          let _2207_v44;
          _2207_v44 = _dafny.Seq.UnicodeFromString("qu");
          _2207_v44 = _2207_v44;
          _2185_cf50 = (_2179_v33).IsProperSubsetOf(_dafny.Set.fromElements(false));
        }
        let _2208_v45;
        _2208_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(p0, globalState),p0);
        _2208_v45 = (_2208_v45).update(_module.__default.safeModuloInt((_2174_v28)[_module.__default.safeIndex(p0, new BigNumber((_2174_v28).length))], p0), (_dafny.ZERO).minus((_this).fm2(p0, globalState)));
        if (false) {
          _2185_cf50 = !(true);
          let _2209_v46;
          _2209_v46 = _dafny.Seq.UnicodeFromString("auf");
          _2209_v46 = _2209_v46;
          (globalState).f13 = p0;
          let _2210_v47;
          let _nw342 = new _module.C10();
          _nw342.__ctor();
          _2210_v47 = _nw342;
          let _2211_v48;
          _2211_v48 = _dafny.MultiSet.fromElements(new BigNumber(388));
          (globalState).f7 = _module.__default.safeModuloInt(p0, (p0).plus((((_2211_v48).contains(p0)) ? ((_2211_v48).get(p0)) : (p0))));
        } else {
          let _2212_v49;
          _2212_v49 = _dafny.MultiSet.fromElements(!(p3));
          let _2213_v50;
          _2213_v50 = _dafny.MultiSet.fromElements((_2212_v49).update(_2185_cf50, _module.__default.abs(p0)), _dafny.MultiSet.fromElements(p3, p3, p3, _2185_cf50));
          let _2214_v51;
          let _nw343 = Array((new BigNumber(4)).toNumber());
          _nw343[(_dafny.ZERO).toNumber()] = _2185_cf50;
          _nw343[(_dafny.ONE).toNumber()] = (_2213_v50).IsDisjointFrom(_2213_v50);
          _nw343[(new BigNumber(2)).toNumber()] = p1;
          _nw343[(new BigNumber(3)).toNumber()] = p1;
          _2214_v51 = _nw343;
          let _2215_v52;
          _2215_v52 = _module.D6.create_DC16(_dafny.Seq.of(p0));
          let _index373 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          (_2214_v51)[_index373] = _dafny.areEqual((_2215_v52).dtor_cf23, _2174_v28);
          let _index374 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          (_2214_v51)[_index374] = p3;
          let _2216_v53;
          _2216_v53 = _dafny.Set.fromElements((new BigNumber((_2174_v28).length)).plus(p0), new BigNumber(655));
          _2216_v53 = ((_2216_v53).Union(_2216_v53)).Intersect(_2216_v53);
          let _2217_v54;
          _2217_v54 = _module.D7.create_DC22(p3, p0, (_2214_v51)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length))], p0);
          let _2218_v55;
          _2218_v55 = _dafny.MultiSet.fromElements(_2217_v54, _module.D7.create_DC22(p3, p0, p1, p0), _2217_v54, _module.D7.create_DC22(p3, p0, true, p0), _2217_v54);
          let _2219_v56;
          _2219_v56 = _module.D11.create_DC30(_2185_cf50, p0);
          let _2220_v57;
          _2220_v57 = _dafny.Set.fromElements(_module.__default.fm38(_2185_cf50, globalState));
          let _2221_v58;
          _2221_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2180_v34,_2180_v34);
          let _2222_v60;
          _2222_v60 = _module.D5.create_DC13(_module.__default.fm43(p0, globalState));
          let _2223_v61;
          _2223_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2222_v60);
          let _index375 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          let _rhs444 = ((_2220_v57).Difference(function () {
            let _coll82 = new _dafny.Set();
            for (const _compr_82 of (_2221_v58).Keys.Elements) {
              let _2224_v59 = _compr_82;
              if ((_2221_v58).contains(_2224_v59)) {
                _coll82.add(_2224_v59);
              }
            }
            return _coll82;
          }())).IsProperSubsetOf(_dafny.Set.fromElements(_2180_v34));
          let _rhs445 = _module.__default.fm57(p0, _2223_v61, globalState);
          let _rhs446 = _2219_v56;
          let _lhs301 = _2214_v51;
          let _lhs302 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          _lhs301[_lhs302] = _rhs444;
          _2218_v55 = _rhs445;
          _2219_v56 = _rhs446;
          let _2225_v62;
          let _nw344 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _2225_v62 = _nw344;
          let _2226_v63;
          _2226_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2180_v34,p0);
          let _2227_v64;
          _2227_v64 = _dafny.Seq.UnicodeFromString("fhuvugk");
          let _2228_v65;
          let _nw345 = new _module.C6();
          _nw345.__ctor();
          _2228_v65 = _nw345;
          let _2229_v66;
          _2229_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(788),_2228_v65);
          let _2230_v67;
          let _nw346 = Array((new BigNumber(24)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = (((_2208_v45).contains(new BigNumber((_2226_v63).length))) ? ((_2208_v45).get(new BigNumber((_2226_v63).length))) : (p0));
          _nw346[(_dafny.ONE).toNumber()] = p0;
          _nw346[(new BigNumber(2)).toNumber()] = p0;
          _nw346[(new BigNumber(3)).toNumber()] = p0;
          _nw346[(new BigNumber(4)).toNumber()] = p0;
          _nw346[(new BigNumber(5)).toNumber()] = new BigNumber((_2227_v64).length);
          _nw346[(new BigNumber(6)).toNumber()] = p0;
          _nw346[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("le")).length);
          _nw346[(new BigNumber(8)).toNumber()] = p0;
          _nw346[(new BigNumber(9)).toNumber()] = p0;
          _nw346[(new BigNumber(10)).toNumber()] = p0;
          _nw346[(new BigNumber(11)).toNumber()] = new BigNumber((_2229_v66).length);
          _nw346[(new BigNumber(12)).toNumber()] = new BigNumber((_2227_v64).length);
          _nw346[(new BigNumber(13)).toNumber()] = p0;
          _nw346[(new BigNumber(14)).toNumber()] = new BigNumber(720);
          _nw346[(new BigNumber(15)).toNumber()] = p0;
          _nw346[(new BigNumber(16)).toNumber()] = p0;
          _nw346[(new BigNumber(17)).toNumber()] = p0;
          _nw346[(new BigNumber(18)).toNumber()] = p0;
          _nw346[(new BigNumber(19)).toNumber()] = new BigNumber(-383);
          _nw346[(new BigNumber(20)).toNumber()] = p0;
          _nw346[(new BigNumber(21)).toNumber()] = p0;
          _nw346[(new BigNumber(22)).toNumber()] = p0;
          _nw346[(new BigNumber(23)).toNumber()] = p0;
          _2230_v67 = _nw346;
          let _2231_v68;
          _2231_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC5(_2230_v67),_2185_cf50);
          let _index376 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_2225_v62).length));
          (_2225_v62)[_index376] = _2231_v68;
          let _index377 = _module.__default.safeIndex(new BigNumber(942), new BigNumber((_2225_v62).length));
          (_2225_v62)[_index377] = _2231_v68;
          let _2232_v69;
          let _nw347 = Array((new BigNumber(25)).toNumber());
          _nw347[(_dafny.ZERO).toNumber()] = _2177_v31;
          _nw347[(_dafny.ONE).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(2)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(3)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(4)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(5)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(6)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(7)).toNumber()] = _dafny.Seq.of((_2214_v51)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length))]);
          _nw347[(new BigNumber(8)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(9)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(10)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_2177_v31, _module.__default.safeIndex(p0, new BigNumber((_2177_v31).length)), (_2214_v51)[_module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length))]);
          _nw347[(new BigNumber(12)).toNumber()] = _dafny.Seq.of(p3);
          _nw347[(new BigNumber(13)).toNumber()] = _module.__default.fm27(globalState);
          _nw347[(new BigNumber(14)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(15)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(16)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(17)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(18)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(19)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(20)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(21)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(_2185_cf50, true);
          _nw347[(new BigNumber(23)).toNumber()] = _2177_v31;
          _nw347[(new BigNumber(24)).toNumber()] = _2177_v31;
          _2232_v69 = _nw347;
          let _2233_v70;
          _2233_v70 = _2232_v69;
          let _2234_v71;
          _2234_v71 = _dafny.Seq.of(_2230_v67, _2230_v67, _2230_v67, _2230_v67, _2230_v67);
          let _2235_v72;
          _2235_v72 = _dafny.Seq.of((_2234_v71)[_module.__default.safeIndex(new BigNumber(-155), new BigNumber((_2234_v71).length))]);
          let _2236_v73;
          _2236_v73 = _module.D2.create_DC5(_2230_v67);
          let _index378 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          let _rhs447 = _2233_v70;
          let _rhs448 = _dafny.Seq.Concat(_dafny.Seq.of(_2230_v67, (_2236_v73).dtor_cf6), _2234_v71);
          let _rhs449 = (new BigNumber(868)).isLessThan(p0);
          let _rhs450 = (_2228_v65).fm16(p0, globalState);
          let _rhs451 = p1;
          let _lhs303 = _2214_v51;
          let _lhs304 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_2214_v51).length));
          _2233_v70 = _rhs447;
          _2235_v72 = _rhs448;
          _lhs303[_lhs304] = _rhs449;
          _2185_cf50 = _rhs450;
          _2185_cf50 = _rhs451;
        }
      } else {
        let _2237___mcc_h1 = (_source30).cf49;
        let _2238_cf49 = _2237___mcc_h1;
        let _2239_v74;
        let _nw348 = Array((new BigNumber(11)).toNumber());
        _nw348[(_dafny.ZERO).toNumber()] = _2180_v34;
        _nw348[(_dafny.ONE).toNumber()] = _module.__default.fm38(p1, globalState);
        _nw348[(new BigNumber(2)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(3)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
        _nw348[(new BigNumber(5)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(6)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(7)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw348[(new BigNumber(9)).toNumber()] = _2180_v34;
        _nw348[(new BigNumber(10)).toNumber()] = _2180_v34;
        _2239_v74 = _nw348;
        _2239_v74 = _2239_v74;
        let _2240_v75;
        _2240_v75 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _2241_v76;
        let _nw349 = new _module.C9();
        _nw349.__ctor();
        _2241_v76 = _nw349;
        let _2242_v77;
        _2242_v77 = _module.D17.create_DC47(_2241_v76, p3, p0, p0, p0);
        let _2243_v78;
        _2243_v78 = _module.D18.create_DC49(p1, true, _module.__default.fm38(p3, globalState), (_2242_v77).dtor_cf70);
        let _2244_v79;
        let _nw350 = Array((new BigNumber(18)).toNumber());
        _nw350[(_dafny.ZERO).toNumber()] = p3;
        _nw350[(_dafny.ONE).toNumber()] = !(p3);
        _nw350[(new BigNumber(2)).toNumber()] = false;
        _nw350[(new BigNumber(3)).toNumber()] = p1;
        _nw350[(new BigNumber(4)).toNumber()] = p3;
        _nw350[(new BigNumber(5)).toNumber()] = p1;
        _nw350[(new BigNumber(6)).toNumber()] = p1;
        _nw350[(new BigNumber(7)).toNumber()] = p1;
        _nw350[(new BigNumber(8)).toNumber()] = p3;
        _nw350[(new BigNumber(9)).toNumber()] = (((_2240_v75).contains(new BigNumber(627))) ? ((_2240_v75).get(new BigNumber(627))) : (!(false)));
        _nw350[(new BigNumber(10)).toNumber()] = p1;
        _nw350[(new BigNumber(11)).toNumber()] = !((_2243_v78).dtor_cf78);
        _nw350[(new BigNumber(12)).toNumber()] = false;
        _nw350[(new BigNumber(13)).toNumber()] = p1;
        _nw350[(new BigNumber(14)).toNumber()] = p3;
        _nw350[(new BigNumber(15)).toNumber()] = p3;
        _nw350[(new BigNumber(16)).toNumber()] = p1;
        _nw350[(new BigNumber(17)).toNumber()] = p3;
        _2244_v79 = _nw350;
        let _2245_v80;
        _2245_v80 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(595),_2244_v79);
        _2245_v80 = (_2245_v80).update(p0, _2244_v79);
        let _2246_v81;
        let _nw351 = new _module.C8();
        _nw351.__ctor();
        _2246_v81 = _nw351;
        let _2247_v82;
        let _nw352 = Array((new BigNumber(15)).toNumber());
        _nw352[(_dafny.ZERO).toNumber()] = _2244_v79;
        _nw352[(_dafny.ONE).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(2)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(3)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(4)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(5)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(6)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(7)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(8)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(9)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(10)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(11)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(12)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(13)).toNumber()] = _2244_v79;
        _nw352[(new BigNumber(14)).toNumber()] = _2244_v79;
        _2247_v82 = _nw352;
        _2247_v82 = _2247_v82;
      }
      let _2248_v83;
      let _nw353 = new _module.C10();
      _nw353.__ctor();
      _2248_v83 = _nw353;
      let _2249_v84;
      _2249_v84 = false;
      _2249_v84 = _2249_v84;
      let _2250_v85;
      _2250_v85 = _module.D5.create_DC14(_2180_v34, p1, p0, p1);
      let _2251_v86;
      _2251_v86 = _dafny.Seq.UnicodeFromString("juct");
      let _2252_v87;
      let _nw354 = Array((new BigNumber(5)).toNumber());
      _nw354[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_2250_v85).dtor_cf18, (_2251_v86)[_module.__default.safeIndex(p0, new BigNumber((_2251_v86).length))]);
      _nw354[(_dafny.ONE).toNumber()] = _2251_v86;
      _nw354[(new BigNumber(2)).toNumber()] = _module.__default.fm5(globalState);
      _nw354[(new BigNumber(3)).toNumber()] = _2251_v86;
      _nw354[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(_2180_v34, (_2251_v86)[_module.__default.safeIndex(p0, new BigNumber((_2251_v86).length))]), _module.__default.safeIndex(new BigNumber((_2177_v31).length), new BigNumber((_dafny.Seq.of(_2180_v34, (_2251_v86)[_module.__default.safeIndex(p0, new BigNumber((_2251_v86).length))])).length)), _2180_v34);
      _2252_v87 = _nw354;
      let _2253_v88;
      _2253_v88 = _module.D21.create_DC56(_2252_v87);
      let _source32 = _2253_v88;
      if (_source32.is_DC57) {
        let _2254___mcc_h5 = (_source32).cf90;
        let _2255___mcc_h6 = (_source32).cf91;
        let _2256___mcc_h7 = (_source32).cf92;
        let _2257_cf92 = _2256___mcc_h7;
        let _2258_cf91 = _2255___mcc_h6;
        let _2259_cf90 = _2254___mcc_h5;
        let _2260_v89;
        _2260_v89 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v84,new BigNumber(892));
        let _2261_v90;
        _2261_v90 = _dafny.Map.Empty.slice().updateUnsafe((((_2260_v89).contains(p3)) ? ((_2260_v89).get(p3)) : (_2259_cf90)),p1);
        _2261_v90 = (_module.__default.fm35(globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2259_cf90,p1));
        _2180_v34 = new _dafny.CodePoint('r'.codePointAt(0));
        if (_2249_v84) {
          let _2262_v91;
          let _nw355 = new _module.C8();
          _nw355.__ctor();
          _2262_v91 = _nw355;
          let _2263_v92;
          _2263_v92 = _dafny.MultiSet.fromElements(p3, p1);
          let _rhs452 = _2262_v91;
          let _rhs453 = p1;
          let _rhs454 = _2263_v92;
          _2262_v91 = _rhs452;
          _2249_v84 = _rhs453;
          _2263_v92 = _rhs454;
          _2249_v84 = _2249_v84;
          let _2264_v93;
          _2264_v93 = _dafny.MultiSet.fromElements((_2248_v83).fm2(new BigNumber((_2261_v90).length), globalState));
          let _2265_v94;
          _2265_v94 = _dafny.Set.fromElements(new BigNumber((_2264_v93).cardinality()));
          let _2266_v95;
          let _nw356 = Array((new BigNumber(2)).toNumber());
          _nw356[(_dafny.ZERO).toNumber()] = _2249_v84;
          _nw356[(_dafny.ONE).toNumber()] = (_2248_v83).fm1(_2176_v30, new BigNumber((_2265_v94).length), globalState);
          _2266_v95 = _nw356;
          let _2267_v96;
          _2267_v96 = _dafny.Seq.of(_2266_v95);
          let _2268_v97;
          _2268_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2266_v95,_2259_cf90);
          let _rhs455 = _2259_cf90;
          let _rhs456 = _2267_v96;
          let _rhs457 = !(_2268_v97).contains(_2266_v95);
          let _rhs458 = ((_dafny.ZERO).minus(p0)).plus(_2258_cf91);
          let _lhs305 = globalState;
          let _lhs306 = globalState;
          _lhs305.f7 = _rhs455;
          _2267_v96 = _rhs456;
          _2249_v84 = _rhs457;
          _lhs306.f15 = _rhs458;
          let _2269_v98;
          let _nw357 = new _module.C9();
          _nw357.__ctor();
          _2269_v98 = _nw357;
          let _2270_v99;
          let _nw358 = new _module.C6();
          _nw358.__ctor();
          _2270_v99 = _nw358;
          let _2271_v100;
          _2271_v100 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (_2259_cf90) : (new BigNumber((_2251_v86).length))),_2270_v99);
          (globalState).f13 = new BigNumber((_2271_v100).length);
        } else {
          let _2272_v101;
          _2272_v101 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2179_v33).length),false)).length));
          _2174_v28 = _dafny.Seq.Concat(_dafny.Seq.of(_2259_cf90, new BigNumber((_2272_v101).cardinality()), p0), _2174_v28);
          let _2273_v102;
          _2273_v102 = _module.D6.create_DC16(_dafny.Seq.Create(_module.__default.abs(new BigNumber(494)), ((_2274_p0) => function (_2275_i6) {
  return _2274_p0;
})(p0)));
          let _2276_v103;
          _2276_v103 = _module.D6.create_DC18(_2273_v102);
          let _2277_v104;
          _2277_v104 = _module.D6.create_DC18(_2276_v103);
          let _2278_v105;
          _2278_v105 = _module.D6.create_DC18(_2276_v103);
          let _2279_v106;
          _2279_v106 = _module.D6.create_DC18(_2273_v102);
          let _2280_v107;
          _2280_v107 = _module.D6.create_DC18(_2276_v103);
          _2280_v107 = _module.D6.create_DC18(_module.D6.create_DC18(_2276_v103));
          let _2281_v108;
          _2281_v108 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v84,_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2259_cf90),p1));
          let _2282_v109;
          _2282_v109 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2281_v108).length));
          _2282_v109 = (_2282_v109).update(p0, _module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), function (_2283_i7) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length)));
          _2249_v84 = (_2249_v84) && (p3);
          _2179_v33 = (_2179_v33).Union(_2179_v33);
        }
        if (!(p3) || (!(_2249_v84))) {
          let _2284_v111;
          _2284_v111 = _dafny.Set.fromElements(_2258_cf91);
          let _2285_v112;
          _2285_v112 = _dafny.MultiSet.fromElements(function () {
            let _coll83 = new _dafny.Map();
            for (const _compr_83 of (_2284_v111).Elements) {
              let _2286_v110 = _compr_83;
              if ((_2284_v111).contains(_2286_v110)) {
                _coll83.push([(_2286_v110).multipliedBy(new BigNumber((_2251_v86).length)),_module.D21.create_DC57(_2259_cf90, _2258_cf91, _2257_cf92)]);
              }
            }
            return _coll83;
          }());
          let _2287_v113;
          _2287_v113 = _dafny.MultiSet.fromElements(p0);
          let _2288_v114;
          _2288_v114 = _module.D21.create_DC57(new BigNumber((_2287_v113).cardinality()), _2258_cf91, _module.D5.create_DC14(_2180_v34, p1, p0, p1));
          let _2289_v115;
          _2289_v115 = _dafny.Map.Empty.slice().updateUnsafe(_2258_cf91,_2288_v114);
          let _2290_v116;
          _2290_v116 = _dafny.Seq.of(_2289_v115);
          let _2291_v117;
          _2291_v117 = _dafny.Seq.of(_2285_v112, _dafny.MultiSet.FromArray(_2290_v116));
          let _2292_v118;
          _2292_v118 = _dafny.Seq.of(_2285_v112, (_2291_v117)[_module.__default.safeIndex(_2258_cf91, new BigNumber((_2291_v117).length))]);
          let _2293_v119;
          _2293_v119 = _dafny.Seq.of((_2291_v117)[_module.__default.safeIndex(p0, new BigNumber((_2291_v117).length))]);
          let _2294_v120;
          _2294_v120 = _dafny.Set.fromElements(_2180_v34, _2180_v34, _2180_v34);
          let _2295_v121;
          _2295_v121 = _module.D22.create_DC59(_2285_v112);
          let _pat_let_tv49 = _2285_v112;
          let _2296_v122;
          let _nw359 = Array((new BigNumber(29)).toNumber());
          _nw359[(_dafny.ZERO).toNumber()] = _2285_v112;
          _nw359[(_dafny.ONE).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(2)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(3)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(4)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(5)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(6)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(7)).toNumber()] = (_2285_v112).Difference((_2292_v118)[_module.__default.safeIndex(p0, new BigNumber((_2292_v118).length))]);
          _nw359[(new BigNumber(8)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(9)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.update(_2290_v116, _module.__default.safeIndex((_2248_v83).fm6(p1, p3, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-245),p0), _module.__default.fm38(p3, globalState), globalState), new BigNumber((_2290_v116).length)), _2289_v115))).Union(_2285_v112);
          _nw359[(new BigNumber(11)).toNumber()] = (_2291_v117)[_module.__default.safeIndex(new BigNumber(780), new BigNumber((_2291_v117).length))];
          _nw359[(new BigNumber(12)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(13)).toNumber()] = (_2285_v112).Intersect(_2285_v112);
          _nw359[(new BigNumber(14)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(15)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(16)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(17)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(18)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(19)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(20)).toNumber()] = (_2293_v119)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_2293_v119).length))];
          _nw359[(new BigNumber(21)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements((_2290_v116)[_module.__default.safeIndex(new BigNumber((_2294_v120).length), new BigNumber((_2290_v116).length))], _module.__default.fm58(_2249_v84, _2180_v34, globalState));
          _nw359[(new BigNumber(23)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(24)).toNumber()] = (function (_pat_let41_0) {
            return function (_2297_dt__update__tmp_h1) {
              return function (_pat_let42_0) {
                return function (_2298_dt__update_hcf94_h0) {
                  return _module.D22.create_DC59(_2298_dt__update_hcf94_h0);
                }(_pat_let42_0);
              }(_pat_let_tv49);
            }(_pat_let41_0);
          }(_2295_v121)).dtor_cf94;
          _nw359[(new BigNumber(25)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(26)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(27)).toNumber()] = _2285_v112;
          _nw359[(new BigNumber(28)).toNumber()] = (_2285_v112).Intersect(_2285_v112);
          _2296_v122 = _nw359;
          let _index379 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_2296_v122).length));
          (_2296_v122)[_index379] = _2285_v112;
          let _index380 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_2296_v122).length));
          (_2296_v122)[_index380] = _2285_v112;
          _2249_v84 = _2249_v84;
          _2249_v84 = _2249_v84;
          let _2299_v123;
          let _init56 = ((_2300_cf91) => function (_2301_i8) {
            return (_2301_i8).plus((_dafny.ZERO).minus(_2300_cf91));
          })(_2258_cf91);
          let _nw360 = Array((new BigNumber(28)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw360.length); _i0_56++) {
            _nw360[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2299_v123 = _nw360;
          let _index381 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2299_v123).length));
          (_2299_v123)[_index381] = p0;
          let _2302_v124;
          _2302_v124 = _dafny.MultiSet.fromElements(p1, !(_2249_v84), false);
          let _2303_v125;
          _2303_v125 = _dafny.Map.Empty.slice().updateUnsafe(_2251_v86,_2259_cf90);
          let _index382 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2299_v123).length));
          let _rhs459 = _module.__default.safeDivisionInt(_2258_cf91, new BigNumber((_2287_v113).cardinality()));
          let _rhs460 = new BigNumber((_2302_v124).cardinality());
          let _rhs461 = (((_2303_v125).contains(_2251_v86)) ? ((_2303_v125).get(_2251_v86)) : (_module.__default.safeDivisionInt(p0, new BigNumber((_2251_v86).length))));
          let _lhs307 = globalState;
          let _lhs308 = _2299_v123;
          let _lhs309 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2299_v123).length));
          let _lhs310 = globalState;
          _lhs307.f7 = _rhs459;
          _lhs308[_lhs309] = _rhs460;
          _lhs310.f13 = _rhs461;
          let _2304_v126;
          _2304_v126 = _dafny.Seq.of(_2284_v111, _2284_v111, (_2284_v111).Difference(_dafny.Set.fromElements(new BigNumber(-54), new BigNumber(672), _2259_cf90, _2258_cf91)));
          _2304_v126 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2304_v126, _2304_v126), _2304_v126);
        } else {
          let _2305_v127;
          let _init57 = ((_2306_v34) => function (_2307_i9) {
            return _2306_v34;
          })(_2180_v34);
          let _nw361 = Array((new BigNumber(5)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw361.length); _i0_57++) {
            _nw361[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2305_v127 = _nw361;
          let _2308_v128;
          _2308_v128 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2305_v127);
          _2308_v128 = (_2308_v128).update(p1, _2305_v127);
          let _2309_v129;
          _2309_v129 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_2259_cf90, _2259_cf90),_dafny.Seq.Concat(_2174_v28, _2174_v28));
          _2309_v129 = (_2309_v129).update((_2259_cf90).multipliedBy(p0), _2174_v28);
          _2249_v84 = p1;
          let _2310_v130;
          let _nw362 = new _module.C6();
          _nw362.__ctor();
          _2310_v130 = _nw362;
          _2174_v28 = _dafny.Seq.Concat(_2174_v28, _2174_v28);
        }
      } else if (_source32.is_DC56) {
        let _2311___mcc_h8 = (_source32).cf89;
        let _2312_cf89 = _2311___mcc_h8;
        (globalState).f7 = p0;
        _2249_v84 = p3;
        let _2313_v131;
        let _nw363 = new _module.C9();
        _nw363.__ctor();
        _2313_v131 = _nw363;
        let _2314_v132;
        let _init58 = ((_2315_p1) => function (_2316_i10) {
          return _2315_p1;
        })(p1);
        let _nw364 = Array((new BigNumber(23)).toNumber());
        for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw364.length); _i0_58++) {
          _nw364[_i0_58] = _init58(new BigNumber(_i0_58));
        }
        _2314_v132 = _nw364;
        let _2317_v133;
        _2317_v133 = _module.D0.create_DC0(_2314_v132);
        let _2318_v134;
        _2318_v134 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(415),_2314_v132);
        let _2319_v135;
        _2319_v135 = _dafny.Map.Empty.slice().updateUnsafe(_2249_v84,_2314_v132);
        let _2320_v136;
        _2320_v136 = _dafny.Seq.of(_2314_v132);
        let _2321_v137;
        let _nw365 = Array((new BigNumber(27)).toNumber());
        _nw365[(_dafny.ZERO).toNumber()] = _2314_v132;
        _nw365[(_dafny.ONE).toNumber()] = (_2317_v133).dtor_cf0;
        _nw365[(new BigNumber(2)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(3)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(4)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(5)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(6)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(7)).toNumber()] = (((_2318_v134).contains(p0)) ? ((_2318_v134).get(p0)) : (_2314_v132));
        _nw365[(new BigNumber(8)).toNumber()] = (((_2319_v135).contains(p3)) ? ((_2319_v135).get(p3)) : (_2314_v132));
        _nw365[(new BigNumber(9)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(10)).toNumber()] = ((p1) ? (_2314_v132) : (_2314_v132));
        _nw365[(new BigNumber(11)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(12)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(13)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(14)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(15)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(16)).toNumber()] = (_2320_v136)[_module.__default.safeIndex(p0, new BigNumber((_2320_v136).length))];
        _nw365[(new BigNumber(17)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(18)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(19)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(20)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(21)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(22)).toNumber()] = (((_2319_v135).contains(true)) ? ((_2319_v135).get(true)) : ((_2320_v136)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_2322_i11) {
          return new BigNumber(-388);
        })).length), new BigNumber((_2320_v136).length))]));
        _nw365[(new BigNumber(23)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(24)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(25)).toNumber()] = _2314_v132;
        _nw365[(new BigNumber(26)).toNumber()] = _2314_v132;
        _2321_v137 = _nw365;
        let _2323_v138;
        _2323_v138 = _module.D0.create_DC1(_2314_v132);
        let _index383 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_2321_v137).length));
        (_2321_v137)[_index383] = (_2323_v138).dtor_cf1;
        let _index384 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_2321_v137).length));
        let _init59 = ((_2324_p1) => function (_2325_i12) {
          return _2324_p1;
        })(p1);
        let _nw366 = Array((new BigNumber(22)).toNumber());
        for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw366.length); _i0_59++) {
          _nw366[_i0_59] = _init59(new BigNumber(_i0_59));
        }
        (_2321_v137)[_index384] = _nw366;
      } else {
        let _2326___mcc_h9 = (_source32).cf93;
        let _2327_cf93 = _2326___mcc_h9;
        (globalState).f4 = p0;
        let _2328_v139;
        let _init60 = ((_2329_p0) => function (_2330_i13) {
          return (_2330_i13).multipliedBy(_2329_p0);
        })(p0);
        let _nw367 = Array((new BigNumber(8)).toNumber());
        for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw367.length); _i0_60++) {
          _nw367[_i0_60] = _init60(new BigNumber(_i0_60));
        }
        _2328_v139 = _nw367;
        let _index385 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_2328_v139).length));
        (_2328_v139)[_index385] = p0;
        let _2331_v140;
        _2331_v140 = _dafny.Map.Empty.slice().updateUnsafe(_2328_v139,(_dafny.ZERO).minus(p0));
        let _index386 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_2328_v139).length));
        let _rhs462 = _module.__default.safeDivisionInt(p0, new BigNumber((((_2331_v140).update(_2328_v139, p0)).update(_2328_v139, new BigNumber((_2177_v31).length))).length));
        let _rhs463 = (p0).minus(p0);
        let _lhs311 = _2328_v139;
        let _lhs312 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_2328_v139).length));
        let _lhs313 = globalState;
        _lhs311[_lhs312] = _rhs462;
        _lhs313.f7 = _rhs463;
        _2180_v34 = new _dafny.CodePoint('b'.codePointAt(0));
        let _2332_v141;
        _2332_v141 = _dafny.Set.fromElements(_2180_v34);
        let _2333_v142;
        _2333_v142 = _module.D22.create_DC61(p0, new BigNumber(128), (_this).fm2((_2328_v139)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_2328_v139).length))], globalState), _dafny.Seq.update(_2251_v86, _module.__default.safeIndex(new BigNumber((_2332_v141).length), new BigNumber((_2251_v86).length)), _2180_v34));
        let _2334_v143;
        _2334_v143 = _dafny.MultiSet.fromElements(_2333_v142);
        (globalState).f13 = (_dafny.ZERO).minus(((_2328_v139)[_module.__default.safeIndex(new BigNumber(63), new BigNumber((_2328_v139).length))]).minus(new BigNumber((_2334_v143).cardinality())));
      }
      (globalState).f7 = p0;
      let _2335_v144;
      _2335_v144 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
      r0 = _2335_v144;
      return r0;
    }
  };

  $module.C12 = class C12 {
    constructor () {
      this._tname = "_module.C12";
      this._f22 = _dafny.ZERO;
      this._f23 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f22, f23) {
      let _this = this;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return ((_this).f22).isLessThanOrEqualTo((_this).f22);
    };
    fm2(p0, globalState) {
      let _this = this;
      return ((_this).f22).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()));
    };
    fm3(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_this).f22)).length), _module.__default.safeModuloInt((_this).f22, new BigNumber(561)));
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false, true))).cardinality()))).Union(_dafny.Set.fromElements((_this).f22, new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f22)).length)))).Intersect((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f22, (_this).f22)).cardinality()), (_this).f22)).Intersect(function () {
        let _coll84 = new _dafny.Set();
        for (const _compr_84 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f22,false)).Keys.Elements) {
          let _2336_v0 = _compr_84;
          if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f22,false)).contains(_2336_v0)) {
            _coll84.add((_2336_v0).multipliedBy(new BigNumber(-991)));
          }
        }
        return _coll84;
      }()));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let _2337_v0;
      _2337_v0 = true;
      (globalState).f4 = (((_2337_v0) ? ((_this).f22) : ((_this).f22))).multipliedBy((_this).f22);
      let _2338_i0;
      _2338_i0 = _dafny.ZERO;
      L15: {
        while (_2337_v0) {
          C15: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2338_i0)) {
              break L15;
            }
            _2338_i0 = (_2338_i0).plus(_dafny.ONE);
            let _2339_v1;
            _2339_v1 = new _dafny.CodePoint('r'.codePointAt(0));
            _2339_v1 = (p0)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f22), new BigNumber((p0).length))];
            let _2340_v2;
            let _nw368 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
            _2340_v2 = _nw368;
            _2340_v2 = _2340_v2;
            let _2341_v3;
            _2341_v3 = _dafny.Set.fromElements((_this).f22, (_this).fm2((_this).f22, globalState));
            let _2342_v4;
            _2342_v4 = _dafny.Map.Empty.slice().updateUnsafe((_2341_v3).IsProperSubsetOf(_2341_v3),_2337_v0);
            if ((((_2342_v4).contains(_2337_v0)) ? ((_2342_v4).get(_2337_v0)) : (false))) {
              let _2343_v5;
              _2343_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2337_v0,_2342_v4);
              let _2344_v6;
              _2344_v6 = _dafny.Seq.of(_2337_v0, _2337_v0, _2337_v0);
              let _2345_v7;
              let _nw369 = Array((new BigNumber(6)).toNumber());
              _nw369[(_dafny.ZERO).toNumber()] = (_this).fm1(_2343_v5, (_this).f22, globalState);
              _nw369[(_dafny.ONE).toNumber()] = (_2344_v6)[_module.__default.safeIndex(new BigNumber(732), new BigNumber((_2344_v6).length))];
              _nw369[(new BigNumber(2)).toNumber()] = (new BigNumber(873)).isLessThanOrEqualTo((_this).f22);
              _nw369[(new BigNumber(3)).toNumber()] = true;
              _nw369[(new BigNumber(4)).toNumber()] = _dafny.areEqual(p0, p0);
              _nw369[(new BigNumber(5)).toNumber()] = _2337_v0;
              _2345_v7 = _nw369;
              let _index387 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2345_v7).length));
              (_2345_v7)[_index387] = _2337_v0;
              let _index388 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2345_v7).length));
              (_2345_v7)[_index388] = _2337_v0;
              (globalState).f7 = (_this).f22;
              (globalState).f13 = (_this).fm2((_this).f22, globalState);
              let _index389 = _module.__default.safeIndex(new BigNumber(274), new BigNumber((_2345_v7).length));
              (_2345_v7)[_index389] = !_dafny.Seq.contains((_this).f23, _2339_v1);
              let _nw370 = Array((new BigNumber(10)).toNumber()).fill(false);
              _2345_v7 = _nw370;
            } else {
              let _2346_v8;
              _2346_v8 = _dafny.Seq.UnicodeFromString("hlrwshec");
              _2346_v8 = ((((_this).f22).isLessThan(new BigNumber(-642))) ? ((_this).f23) : (_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("umwm"), _dafny.Seq.UnicodeFromString("npllbkx")), _module.__default.safeIndex((_dafny.ZERO).minus((_this).f22), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("umwm"), _dafny.Seq.UnicodeFromString("npllbkx"))).length)), _2339_v1)));
              let _2347_v9;
              let _nw371 = Array((new BigNumber(11)).toNumber()).fill([]);
              _2347_v9 = _nw371;
              let _2348_v10;
              let _init61 = ((_2349_v0) => function (_2350_i1) {
                return _dafny.Map.Empty.slice().updateUnsafe(_2349_v0,new BigNumber(-494));
              })(_2337_v0);
              let _nw372 = Array((new BigNumber(20)).toNumber());
              for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw372.length); _i0_61++) {
                _nw372[_i0_61] = _init61(new BigNumber(_i0_61));
              }
              _2348_v10 = _nw372;
              let _index390 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2347_v9).length));
              (_2347_v9)[_index390] = _2348_v10;
              let _2351_v11;
              let _init62 = ((_2352_v0) => function (_2353_i2) {
                return _2352_v0;
              })(_2337_v0);
              let _nw373 = Array((new BigNumber(6)).toNumber());
              for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw373.length); _i0_62++) {
                _nw373[_i0_62] = _init62(new BigNumber(_i0_62));
              }
              _2351_v11 = _nw373;
              let _2354_v12;
              _2354_v12 = _dafny.Seq.of(_2351_v11);
              let _index391 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2347_v9).length));
              let _rhs464 = _2348_v10;
              let _rhs465 = _dafny.Seq.UnicodeFromString("oexfuyyh");
              let _rhs466 = new BigNumber((_2354_v12).length);
              let _lhs314 = _2347_v9;
              let _lhs315 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_2347_v9).length));
              let _lhs316 = globalState;
              _lhs314[_lhs315] = _rhs464;
              _2346_v8 = _rhs465;
              _lhs316.f7 = _rhs466;
              _2346_v8 = p0;
              let _2355_v13;
              let _nw374 = new _module.C6();
              _nw374.__ctor();
              _2355_v13 = _nw374;
              (globalState).f7 = ((_this).f22).multipliedBy((_this).f22);
            }
            _2337_v0 = !(_2337_v0);
          }
        }
      }
      let _2356_v14;
      _2356_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_2337_v0);
      let _2357_v15;
      _2357_v15 = _module.D14.create_DC37(_module.D14.create_DC35(_dafny.Seq.of(_2356_v14)));
      let _2358_v16;
      _2358_v16 = _dafny.Seq.of(_2337_v0, _2337_v0, _2337_v0);
      _2357_v15 = _module.D14.create_DC37(_module.D14.create_DC36(_2358_v16));
      let _hi15 = (_this).f22;
      for (let _2359_i3 = (_this).f22; _2359_i3.isLessThan(_hi15); _2359_i3 = _2359_i3.plus(_dafny.ONE)) {
        if ((_2359_i3).isLessThan((_this).f22)) {
          let _2360_v17;
          let _init63 = ((_2361_i3) => function (_2362_i4) {
            return _module.__default.safeDivisionInt(_2362_i4, _2361_i3);
          })(_2359_i3);
          let _nw375 = Array((new BigNumber(7)).toNumber());
          for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw375.length); _i0_63++) {
            _nw375[_i0_63] = _init63(new BigNumber(_i0_63));
          }
          _2360_v17 = _nw375;
          _2360_v17 = _2360_v17;
          _2337_v0 = (_2337_v0) === (_2337_v0);
          let _2363_v18;
          let _nw376 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2363_v18 = _nw376;
          let _2364_v19;
          _2364_v19 = new _dafny.CodePoint('a'.codePointAt(0));
          let _index392 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2363_v18).length));
          (_2363_v18)[_index392] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("spf"), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.UnicodeFromString("spf")).length)), _2364_v19);
          let _2365_v20;
          _2365_v20 = _dafny.Seq.UnicodeFromString("s");
          let _index393 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2363_v18).length));
          let _rhs467 = _2360_v17;
          let _rhs468 = (_this).f23;
          let _rhs469 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(p0, p0), _module.__default.safeIndex(((_this).f22).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_2366_v19) => function (_2367_i5) {
            return _2366_v19;
          })(_2364_v19))).length)), new BigNumber((_dafny.Seq.Concat(p0, p0)).length)), _2364_v19), _module.__default.safeIndex(_module.__default.safeDivisionInt(_2359_i3, new BigNumber(534)), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(p0, p0), _module.__default.safeIndex(((_this).f22).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_2368_v19) => function (_2369_i5) {
            return _2368_v19;
          })(_2364_v19))).length)), new BigNumber((_dafny.Seq.Concat(p0, p0)).length)), _2364_v19)).length)), _2364_v19), _module.__default.safeIndex(((_2337_v0) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_2358_v16).length)))) : (_2359_i3)), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(p0, p0), _module.__default.safeIndex(((_this).f22).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_2370_v19) => function (_2371_i5) {
            return _2370_v19;
          })(_2364_v19))).length)), new BigNumber((_dafny.Seq.Concat(p0, p0)).length)), _2364_v19), _module.__default.safeIndex(_module.__default.safeDivisionInt(_2359_i3, new BigNumber(534)), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(p0, p0), _module.__default.safeIndex(((_this).f22).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_2372_v19) => function (_2373_i5) {
            return _2372_v19;
          })(_2364_v19))).length)), new BigNumber((_dafny.Seq.Concat(p0, p0)).length)), _2364_v19)).length)), _2364_v19)).length)), _2364_v19);
          let _lhs317 = _2363_v18;
          let _lhs318 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_2363_v18).length));
          _2360_v17 = _rhs467;
          _lhs317[_lhs318] = _rhs468;
          _2365_v20 = _rhs469;
          let _2374_v21;
          _2374_v21 = _module.D8.create_DC25((_this).f22);
          _2374_v21 = function (_pat_let43_0) {
            return function (_2377_dt__update__tmp_h1) {
              return function (_pat_let46_0) {
                return function (_2378_dt__update_hcf42_h1) {
                  return _module.D8.create_DC25(_2378_dt__update_hcf42_h1);
                }(_pat_let46_0);
              }((_2359_i3).minus((_this).f22));
            }(_pat_let43_0);
          }(function (_pat_let44_0) {
            return function (_2375_dt__update__tmp_h0) {
              return function (_pat_let45_0) {
                return function (_2376_dt__update_hcf42_h0) {
                  return _module.D8.create_DC25(_2376_dt__update_hcf42_h0);
                }(_pat_let45_0);
              }(_2359_i3);
            }(_pat_let44_0);
          }(_2374_v21));
          _2364_v19 = _2364_v19;
        } else {
          let _2379_v22;
          _2379_v22 = _dafny.MultiSet.fromElements(_2337_v0);
          (globalState).f13 = new BigNumber((_2379_v22).cardinality());
          let _2380_v23;
          let _nw377 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _2380_v23 = _nw377;
          let _index394 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length));
          (_2380_v23)[_index394] = _2359_i3;
          let _2381_v24;
          _2381_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2337_v0,_2337_v0);
          let _2382_v25;
          _2382_v25 = _module.D15.create_DC39(_2381_v24, _2337_v0, new BigNumber((_2379_v22).cardinality()));
          let _pat_let_tv50 = _2337_v0;
          let _index395 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length));
          (_2380_v23)[_index395] = ((_2337_v0) ? ((function (_pat_let47_0) {
            return function (_2383_dt__update__tmp_h2) {
              return function (_pat_let48_0) {
                return function (_2384_dt__update_hcf57_h0) {
                  return _module.D15.create_DC39((_2383_dt__update__tmp_h2).dtor_cf56, _2384_dt__update_hcf57_h0, (_2383_dt__update__tmp_h2).dtor_cf58);
                }(_pat_let48_0);
              }(!(_pat_let_tv50));
            }(_pat_let47_0);
          }(_2382_v25)).dtor_cf58) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f22).multipliedBy(_2359_i3)))));
          let _2385_v26;
          let _nw378 = new _module.C2();
          _nw378.__ctor();
          _2385_v26 = _nw378;
          _2337_v0 = !(_2337_v0) || ((_2385_v26).fm8(new BigNumber(24), _2337_v0, _2359_i3, globalState));
          let _2386_v27;
          let _nw379 = Array((new BigNumber(16)).toNumber()).fill([]);
          _2386_v27 = _nw379;
          let _2387_v28;
          let _nw380 = Array((new BigNumber(18)).toNumber());
          _nw380[(_dafny.ZERO).toNumber()] = _2337_v0;
          _nw380[(_dafny.ONE).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(2)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(3)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(4)).toNumber()] = !(_2337_v0);
          _nw380[(new BigNumber(5)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(6)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(7)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(8)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(9)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(10)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(11)).toNumber()] = _module.__default.fm26(globalState);
          _nw380[(new BigNumber(12)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(13)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(14)).toNumber()] = (_2385_v26).fm8((_2380_v23)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length))], _2337_v0, _module.__default.fm17(new BigNumber(57), globalState), globalState);
          _nw380[(new BigNumber(15)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(16)).toNumber()] = _2337_v0;
          _nw380[(new BigNumber(17)).toNumber()] = (((_2356_v14).contains(_2359_i3)) ? ((_2356_v14).get(_2359_i3)) : (_2337_v0));
          _2387_v28 = _nw380;
          let _index396 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_2386_v27).length));
          (_2386_v27)[_index396] = _2387_v28;
          let _2388_v29;
          _2388_v29 = new _dafny.CodePoint('a'.codePointAt(0));
          let _2389_v30;
          _2389_v30 = _dafny.MultiSet.fromElements(_2388_v29, _2388_v29);
          let _2390_v31;
          _2390_v31 = _module.D7.create_DC22(_2337_v0, (_2380_v23)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length))], _2337_v0, (_2380_v23)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length))]);
          let _index397 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_2386_v27).length));
          let _nw381 = Array((new BigNumber(22)).toNumber());
          _nw381[(_dafny.ZERO).toNumber()] = (_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(true,_2381_v24), (_dafny.ZERO).minus(_2359_i3), globalState);
          _nw381[(_dafny.ONE).toNumber()] = (_2337_v0) || (_2337_v0);
          _nw381[(new BigNumber(2)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(3)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(4)).toNumber()] = ((_2337_v0) ? (_2337_v0) : (_2337_v0));
          _nw381[(new BigNumber(5)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(6)).toNumber()] = (_2389_v30).IsSubsetOf(_2389_v30);
          _nw381[(new BigNumber(7)).toNumber()] = true;
          _nw381[(new BigNumber(8)).toNumber()] = !(true) || (_2337_v0);
          _nw381[(new BigNumber(9)).toNumber()] = ((_2337_v0) ? (_2337_v0) : (_2337_v0));
          _nw381[(new BigNumber(10)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(11)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of((_2380_v23)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_2380_v23).length))]))).contains((_dafny.ZERO).minus(_2359_i3));
          _nw381[(new BigNumber(13)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(14)).toNumber()] = !(_2337_v0);
          _nw381[(new BigNumber(15)).toNumber()] = (((_2356_v14).contains(new BigNumber((_module.__default.fm35(globalState)).length))) ? ((_2356_v14).get(new BigNumber((_module.__default.fm35(globalState)).length))) : ((_2390_v31).dtor_cf37));
          _nw381[(new BigNumber(16)).toNumber()] = false;
          _nw381[(new BigNumber(17)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(18)).toNumber()] = _2337_v0;
          _nw381[(new BigNumber(19)).toNumber()] = true;
          _nw381[(new BigNumber(20)).toNumber()] = false;
          _nw381[(new BigNumber(21)).toNumber()] = !(_2337_v0);
          (_2386_v27)[_index397] = _nw381;
        }
        let _2391_v32;
        _2391_v32 = _dafny.MultiSet.fromElements(false, _2337_v0);
        let _2392_v33;
        let _nw382 = Array((new BigNumber(2)).toNumber());
        _nw382[(_dafny.ZERO).toNumber()] = _2337_v0;
        _nw382[(_dafny.ONE).toNumber()] = (_dafny.MultiSet.fromElements(_2337_v0)).IsProperSubsetOf(_2391_v32);
        _2392_v33 = _nw382;
        let _index398 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length));
        (_2392_v33)[_index398] = !(_2337_v0) || (_2337_v0);
        let _index399 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length));
        (_2392_v33)[_index399] = _2337_v0;
        let _2393_v34;
        let _nw383 = Array((new BigNumber(14)).toNumber()).fill([]);
        _2393_v34 = _nw383;
        let _2394_v35;
        let _nw384 = Array((new BigNumber(19)).toNumber());
        _nw384[(_dafny.ZERO).toNumber()] = new BigNumber(684);
        _nw384[(_dafny.ONE).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(2)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(3)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(4)).toNumber()] = new BigNumber(871);
        _nw384[(new BigNumber(5)).toNumber()] = new BigNumber((_module.__default.fm27(globalState)).length);
        _nw384[(new BigNumber(6)).toNumber()] = _2359_i3;
        _nw384[(new BigNumber(7)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(8)).toNumber()] = _2359_i3;
        _nw384[(new BigNumber(9)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(10)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(11)).toNumber()] = new BigNumber((_2358_v16).length);
        _nw384[(new BigNumber(12)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(13)).toNumber()] = _2359_i3;
        _nw384[(new BigNumber(14)).toNumber()] = _2359_i3;
        _nw384[(new BigNumber(15)).toNumber()] = new BigNumber((_2391_v32).cardinality());
        _nw384[(new BigNumber(16)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(17)).toNumber()] = (_this).f22;
        _nw384[(new BigNumber(18)).toNumber()] = (_this).f22;
        _2394_v35 = _nw384;
        let _2395_v36;
        _2395_v36 = _module.D2.create_DC5(_2394_v35);
        let _2396_v38;
        _2396_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_module.__default.fm30(_module.__default.fm17(_2359_i3, globalState), (_2392_v33)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length))], function () {
          let _coll85 = new _dafny.Set();
          for (const _compr_85 of _dafny.IntegerRange(new BigNumber(303), new BigNumber(-84))) {
            let _2397_v37 = _compr_85;
            if (((new BigNumber(303)).isLessThanOrEqualTo(_2397_v37)) && ((_2397_v37).isLessThan(new BigNumber(-84)))) {
              _coll85.add((_2397_v37).multipliedBy(_2359_i3));
            }
          }
          return _coll85;
        }(), _2337_v0, globalState));
        let _2398_v39;
        _2398_v39 = _dafny.Seq.of(_2396_v38, _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_dafny.Seq.UnicodeFromString("gpjnbt")), _2396_v38);
        let _2399_v40;
        _2399_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2398_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f22), new BigNumber((_2398_v39).length))]).length),_2394_v35);
        let _2400_v41;
        let _nw385 = Array((new BigNumber(12)).toNumber());
        _nw385[(_dafny.ZERO).toNumber()] = (_2395_v36).dtor_cf6;
        _nw385[(_dafny.ONE).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(2)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(3)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(4)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(5)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(6)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(7)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(8)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(9)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(10)).toNumber()] = _2394_v35;
        _nw385[(new BigNumber(11)).toNumber()] = (((_2399_v40).contains(_2359_i3)) ? ((_2399_v40).get(_2359_i3)) : (_2394_v35));
        _2400_v41 = _nw385;
        let _index400 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_2393_v34).length));
        (_2393_v34)[_index400] = _2400_v41;
        let _index401 = _module.__default.safeIndex(new BigNumber(647), new BigNumber((_2393_v34).length));
        (_2393_v34)[_index401] = _2400_v41;
        if (_dafny.Seq.IsPrefixOf(p0, p0)) {
          (globalState).f15 = (_this).f22;
          let _2401_v42;
          _2401_v42 = _dafny.Seq.UnicodeFromString("ks");
          _2401_v42 = (((_2392_v33)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length))]) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-693)), function (_2402_i6) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          })) : (_dafny.Seq.Concat(_2401_v42, p0)));
          let _2403_v43;
          let _nw386 = Array((new BigNumber(26)).toNumber()).fill([]);
          _2403_v43 = _nw386;
          let _index402 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2403_v43).length));
          (_2403_v43)[_index402] = _2392_v33;
          let _2404_v44;
          let _nw387 = new _module.C2();
          _nw387.__ctor();
          _2404_v44 = _nw387;
          let _2405_v45;
          _2405_v45 = _dafny.Map.Empty.slice().updateUnsafe(_2404_v44,_2392_v33);
          let _index403 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_2403_v43).length));
          (_2403_v43)[_index403] = (((_2405_v45).contains(_2404_v44)) ? ((_2405_v45).get(_2404_v44)) : (_2392_v33));
          let _2406_v46;
          _2406_v46 = _dafny.Set.fromElements((_2392_v33)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length))]);
          _2356_v14 = (_2356_v14).update(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f22), _2359_i3), (_2406_v46).IsDisjointFrom(_2406_v46));
          _2406_v46 = (_2406_v46).Union(_2406_v46);
        } else {
          let _2407_v47;
          let _nw388 = new _module.C11();
          _nw388.__ctor();
          _2407_v47 = _nw388;
          let _2408_v48;
          _2408_v48 = _dafny.Map.Empty.slice().updateUnsafe(_2407_v47,(_2392_v33)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length))]);
          let _2409_v49;
          _2409_v49 = _dafny.Seq.of(_2408_v48, _dafny.Map.Empty.slice().updateUnsafe(_2407_v47,(_2392_v33)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length))]));
          _2337_v0 = !(!_dafny.areEqual(_2409_v49, _2409_v49));
          let _index404 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length));
          (_2392_v33)[_index404] = !(_2337_v0);
          let _2410_v50;
          let _nw389 = new _module.C2();
          _nw389.__ctor();
          _2410_v50 = _nw389;
          let _2411_v51;
          _2411_v51 = _dafny.Seq.UnicodeFromString("jhounbr");
          let _2412_v52;
          _2412_v52 = _module.D0.create_DC1(_2392_v33);
          let _index405 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length));
          let _rhs470 = (_dafny.ZERO).minus(_2359_i3);
          let _rhs471 = _2337_v0;
          let _rhs472 = (((_2396_v38).contains(new BigNumber(((_this).f23).length))) ? ((_2396_v38).get(new BigNumber(((_this).f23).length))) : ((_this).f23));
          let _rhs473 = _2412_v52;
          let _lhs319 = globalState;
          let _lhs320 = _2392_v33;
          let _lhs321 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_2392_v33).length));
          _lhs319.f15 = _rhs470;
          _lhs320[_lhs321] = _rhs471;
          _2411_v51 = _rhs472;
          _2412_v52 = _rhs473;
          let _2413_v53;
          _2413_v53 = _dafny.Seq.of(_module.__default.safeModuloInt((_this).f22, _2359_i3));
          let _2414_v54;
          _2414_v54 = _dafny.Seq.of(_2411_v51);
          let _rhs474 = new BigNumber((_2356_v14).length);
          let _rhs475 = _2413_v53;
          let _rhs476 = p0;
          let _rhs477 = _2414_v54;
          let _lhs322 = globalState;
          _lhs322.f4 = _rhs474;
          _2413_v53 = _rhs475;
          _2411_v51 = _rhs476;
          _2414_v54 = _rhs477;
        }
      }
      let _2415_v55;
      let _nw390 = new _module.C0();
      _nw390.__ctor();
      _2415_v55 = _nw390;
      let _2416_v56;
      let _init64 = function (_2417_i7) {
        return (_2417_i7).minus((_this).f22);
      };
      let _nw391 = Array((new BigNumber(13)).toNumber());
      for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw391.length); _i0_64++) {
        _nw391[_i0_64] = _init64(new BigNumber(_i0_64));
      }
      _2416_v56 = _nw391;
      let _index406 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_2416_v56).length));
      (_2416_v56)[_index406] = (_this).f22;
      let _2418_v57;
      let _nw392 = new _module.C2();
      _nw392.__ctor();
      _2418_v57 = _nw392;
      let _2419_v58;
      _2419_v58 = new _dafny.CodePoint('n'.codePointAt(0));
      let _2420_v59;
      _2420_v59 = _dafny.Set.fromElements(_2337_v0);
      let _2421_v60;
      _2421_v60 = _module.D5.create_DC14(_2419_v58, _2337_v0, new BigNumber((_2420_v59).length), _2337_v0);
      let _2422_v61;
      _2422_v61 = _dafny.Map.Empty.slice().updateUnsafe(_2421_v60,(_this).f22);
      let _2423_v62;
      _2423_v62 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber((_2422_v61).length));
      let _index407 = _module.__default.safeIndex(new BigNumber(366), new BigNumber((_2416_v56).length));
      (_2416_v56)[_index407] = (new BigNumber((_dafny.Seq.of(_2418_v57, _2418_v57)).length)).minus((((_2423_v62).contains(_2419_v58)) ? ((_2423_v62).get(_2419_v58)) : ((_this).f22)));
      return;
    }
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let _2424_v0;
      _2424_v0 = new _dafny.CodePoint('k'.codePointAt(0));
      let _2425_v1;
      _2425_v1 = _dafny.Set.fromElements((_this).f23, _dafny.Seq.update((_this).f23, _module.__default.safeIndex((_this).f22, new BigNumber(((_this).f23).length)), _2424_v0));
      if (((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("jieavydw"), _module.__default.fm19(p1, p1, globalState))).Union(_2425_v1)).contains((_this).f23)) {
        let _2426_v2;
        _2426_v2 = _dafny.Seq.of(p1, !(!(!(p1))));
        let _2427_v3;
        _2427_v3 = _dafny.Seq.of((_2426_v2)[_module.__default.safeIndex(new BigNumber((_2426_v2).length), new BigNumber((_2426_v2).length))]);
        let _2428_v4;
        _2428_v4 = _module.D18.create_DC49(p1, p1, _2424_v0, (_2427_v3)[_module.__default.safeIndex(p2, new BigNumber((_2427_v3).length))]);
        let _source33 = _2428_v4;
        if (_source33.is_DC49) {
          let _2429___mcc_h0 = (_source33).cf75;
          let _2430___mcc_h1 = (_source33).cf76;
          let _2431___mcc_h2 = (_source33).cf77;
          let _2432___mcc_h3 = (_source33).cf78;
          let _2433_cf78 = _2432___mcc_h3;
          let _2434_cf77 = _2431___mcc_h2;
          let _2435_cf76 = _2430___mcc_h1;
          let _2436_cf75 = _2429___mcc_h0;
          let _2437_v5;
          let _nw393 = Array((new BigNumber(21)).toNumber()).fill(_module.D5.Default());
          _2437_v5 = _nw393;
          let _2438_v6;
          _2438_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2433_cf78);
          let _2439_v7;
          _2439_v7 = _dafny.Set.fromElements(new BigNumber((_2438_v6).length));
          let _2440_v9;
          _2440_v9 = _dafny.MultiSet.fromElements(_2439_v7, function () {
            let _coll86 = new _dafny.Set();
            for (const _compr_86 of _dafny.IntegerRange(new BigNumber(713), new BigNumber(165))) {
              let _2441_v8 = _compr_86;
              if (((new BigNumber(713)).isLessThanOrEqualTo(_2441_v8)) && ((_2441_v8).isLessThan(new BigNumber(165)))) {
                _coll86.add((_2441_v8).multipliedBy(p0));
              }
            }
            return _coll86;
          }());
          let _2442_v10;
          _2442_v10 = _module.D5.create_DC13(_2440_v9);
          let _index408 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_2437_v5).length));
          (_2437_v5)[_index408] = _2442_v10;
          let _pat_let_tv51 = _2439_v7;
          let _pat_let_tv52 = _2439_v7;
          let _pat_let_tv53 = _2440_v9;
          let _index409 = _module.__default.safeIndex(new BigNumber(443), new BigNumber((_2437_v5).length));
          (_2437_v5)[_index409] = function (_pat_let49_0) {
            return function (_2443_dt__update__tmp_h0) {
              return function (_pat_let50_0) {
                return function (_2444_dt__update_hcf17_h0) {
                  return _module.D5.create_DC13(_2444_dt__update_hcf17_h0);
                }(_pat_let50_0);
              }((_dafny.MultiSet.fromElements(_pat_let_tv51, _pat_let_tv52)).Difference(_pat_let_tv53));
            }(_pat_let49_0);
          }(_2442_v10);
          let _2445_v11;
          _2445_v11 = _dafny.MultiSet.fromElements(false);
          let _2446_v12;
          let _nw394 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _2446_v12 = _nw394;
          let _2447_v13;
          _2447_v13 = _module.D2.create_DC8(_module.D2.create_DC7((((_2438_v6).contains((_this).f22)) ? ((_2438_v6).get((_this).f22)) : (_2436_cf75)), _2445_v11, _2446_v12));
          let _2448_v14;
          let _nw395 = new _module.C3();
          _nw395.__ctor(_2447_v13, false);
          _2448_v14 = _nw395;
          let _rhs478 = (p2).minus((_this).f22);
          let _rhs479 = (new BigNumber((_dafny.Seq.Concat(_2426_v2, _dafny.Seq.of(_2448_v14.f25))).length)).isLessThan(p0);
          let _rhs480 = _2448_v14;
          let _lhs323 = globalState;
          _lhs323.f4 = _rhs478;
          _2433_cf78 = _rhs479;
          _2448_v14 = _rhs480;
          _2433_cf78 = !(_2436_cf75);
          let _2449_v15;
          _2449_v15 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_2427_v3)).cardinality()), (_this).f22);
          let _2450_v16;
          _2450_v16 = _dafny.Seq.of(new BigNumber((_2449_v15).cardinality()), new BigNumber(((_this).f23).length));
          (globalState).f15 = new BigNumber((_2450_v16).length);
        } else if (_source33.is_DC48) {
          let _2451___mcc_h4 = (_source33).cf74;
          let _2452_cf74 = _2451___mcc_h4;
          let _2453_v17;
          let _nw396 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.of());
          _2453_v17 = _nw396;
          let _2454_v18;
          let _init65 = ((_2455_v0) => function (_2456_i0) {
            return _2455_v0;
          })(_2424_v0);
          let _nw397 = Array((new BigNumber(22)).toNumber());
          for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw397.length); _i0_65++) {
            _nw397[_i0_65] = _init65(new BigNumber(_i0_65));
          }
          _2454_v18 = _nw397;
          let _2457_v19;
          _2457_v19 = _dafny.Seq.of(_2454_v18);
          let _index410 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_2453_v17).length));
          (_2453_v17)[_index410] = _2457_v19;
          let _2458_v20;
          _2458_v20 = _dafny.Seq.of(_2457_v19);
          let _2459_v21;
          _2459_v21 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(_2454_v18, _2454_v18, _2454_v18), _2457_v19), _dafny.Seq.update((_2458_v20)[_module.__default.safeIndex(p0, new BigNumber((_2458_v20).length))], _module.__default.safeIndex((_dafny.ZERO).minus((_this).f22), new BigNumber(((_2458_v20)[_module.__default.safeIndex(p0, new BigNumber((_2458_v20).length))]).length)), _2454_v18), _2457_v19);
          let _index411 = _module.__default.safeIndex(new BigNumber(764), new BigNumber((_2453_v17).length));
          (_2453_v17)[_index411] = (_2459_v21)[_module.__default.safeIndex((new BigNumber(((_this).f23).length)).minus((_dafny.ZERO).minus((_this).f22)), new BigNumber((_2459_v21).length))];
          let _2460_v22;
          _2460_v22 = _dafny.MultiSet.fromElements(p1, p1);
          let _2461_v23;
          _2461_v23 = _module.D1.create_DC2((((_2460_v22).contains(p1)) ? ((_2460_v22).get(p1)) : (new BigNumber(143))));
          let _2462_v24;
          _2462_v24 = _module.D20.create_DC55(p2, (_this).f22, p1, p1);
          let _2463_v25;
          _2463_v25 = _module.D8.create_DC24(p0, false);
          let _2464_v26;
          _2464_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length),_dafny.Map.Empty.slice().updateUnsafe(p1,_2463_v25));
          let _2465_v27;
          _2465_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2463_v25);
          let _2466_v28;
          _2466_v28 = _dafny.Set.fromElements(!(!(p1)), !(p1));
          let _2467_v29;
          let _nw398 = Array((new BigNumber(11)).toNumber());
          _nw398[(_dafny.ZERO).toNumber()] = (_this).f22;
          _nw398[(_dafny.ONE).toNumber()] = p0;
          _nw398[(new BigNumber(2)).toNumber()] = new BigNumber(678);
          _nw398[(new BigNumber(3)).toNumber()] = p2;
          _nw398[(new BigNumber(4)).toNumber()] = (_this).f22;
          _nw398[(new BigNumber(5)).toNumber()] = (_2462_v24).dtor_cf85;
          _nw398[(new BigNumber(6)).toNumber()] = new BigNumber(((((_2464_v26).contains(p2)) ? ((_2464_v26).get(p2)) : (_2465_v27))).length);
          _nw398[(new BigNumber(7)).toNumber()] = new BigNumber((_2460_v22).cardinality());
          _nw398[(new BigNumber(8)).toNumber()] = new BigNumber((_2466_v28).length);
          _nw398[(new BigNumber(9)).toNumber()] = p2;
          _nw398[(new BigNumber(10)).toNumber()] = new BigNumber(-907);
          _2467_v29 = _nw398;
          let _2468_v30;
          _2468_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2467_v29,_dafny.MultiSet.fromElements(p1));
          let _2469_v31;
          let _nw399 = Array((new BigNumber(21)).toNumber());
          _nw399[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.update(_2426_v2, _module.__default.safeIndex(new BigNumber((_module.__default.fm19(p1, p1, globalState)).length), new BigNumber((_2426_v2).length)), p1))).Union(_2460_v22);
          _nw399[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(!(p1), p1);
          _nw399[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements(p1);
          _nw399[(new BigNumber(3)).toNumber()] = (_2460_v22).update(false, _module.__default.abs(p0));
          _nw399[(new BigNumber(4)).toNumber()] = (_2460_v22).Union(_module.__default.fm34((_this).f22, p0, p1, globalState));
          _nw399[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_2427_v3);
          _nw399[(new BigNumber(6)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(7)).toNumber()] = (_2460_v22).update(false, _module.__default.abs(p2));
          _nw399[(new BigNumber(8)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(9)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(10)).toNumber()] = _dafny.MultiSet.fromElements(!(p1));
          _nw399[(new BigNumber(11)).toNumber()] = (_2460_v22).Union(_2460_v22);
          _nw399[(new BigNumber(12)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(13)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(14)).toNumber()] = (_2460_v22).Union(_2460_v22);
          _nw399[(new BigNumber(15)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(16)).toNumber()] = ((p1) ? (_2460_v22) : (_2460_v22));
          _nw399[(new BigNumber(17)).toNumber()] = _2460_v22;
          _nw399[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(p1);
          _nw399[(new BigNumber(19)).toNumber()] = ((p1) ? (_2460_v22) : (_2460_v22));
          _nw399[(new BigNumber(20)).toNumber()] = (((_2468_v30).contains(_2467_v29)) ? ((_2468_v30).get(_2467_v29)) : (_2460_v22));
          _2469_v31 = _nw399;
          let _2470_v32;
          _2470_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2427_v3)[_module.__default.safeIndex(p0, new BigNumber((_2427_v3).length))]);
          let _index412 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_2454_v18).length));
          (_2454_v18)[_index412] = ((_this).f23)[_module.__default.safeIndex((_this).f22, new BigNumber(((_this).f23).length))];
          let _2471_v33;
          _2471_v33 = _dafny.Seq.of(_2469_v31, _2469_v31);
          let _index413 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_2454_v18).length));
          let _rhs481 = _2461_v23;
          let _rhs482 = (_2471_v33)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_2471_v33).length))];
          let _rhs483 = _2470_v32;
          let _rhs484 = _2424_v0;
          let _rhs485 = (_dafny.ZERO).minus(p2);
          let _lhs324 = _2454_v18;
          let _lhs325 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_2454_v18).length));
          let _lhs326 = globalState;
          _2461_v23 = _rhs481;
          _2469_v31 = _rhs482;
          _2470_v32 = _rhs483;
          _lhs324[_lhs325] = _rhs484;
          _lhs326.f13 = _rhs485;
          (globalState).f15 = (_dafny.ZERO).minus(p0);
          let _2472_v34;
          _2472_v34 = _dafny.Seq.of(p0, p2);
          (globalState).f15 = (_2472_v34)[_module.__default.safeIndex(new BigNumber(219), new BigNumber((_2472_v34).length))];
        } else {
          let _2473___mcc_h5 = (_source33).cf79;
          let _2474_cf79 = _2473___mcc_h5;
          let _2475_v35;
          _2475_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _2476_v36;
          _2476_v36 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2475_v35);
          let _2477_v37;
          let _nw400 = Array((new BigNumber(20)).toNumber());
          _nw400[(_dafny.ZERO).toNumber()] = _2424_v0;
          _nw400[(_dafny.ONE).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(2)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(3)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(4)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(5)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(6)).toNumber()] = ((p1) ? (_2424_v0) : (_2424_v0));
          _nw400[(new BigNumber(7)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(8)).toNumber()] = _module.__default.fm38(p1, globalState);
          _nw400[(new BigNumber(9)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(10)).toNumber()] = (((_this).fm1(_2476_v36, new BigNumber((_2426_v2).length), globalState)) ? (_2424_v0) : (_2424_v0));
          _nw400[(new BigNumber(11)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(12)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(13)).toNumber()] = ((p1) ? (_2424_v0) : (_2424_v0));
          _nw400[(new BigNumber(14)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(15)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(16)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(17)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(18)).toNumber()] = _2424_v0;
          _nw400[(new BigNumber(19)).toNumber()] = _2424_v0;
          _2477_v37 = _nw400;
          _2477_v37 = _2477_v37;
          let _2478_v38;
          _2478_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f22),p2);
          let _2479_v39;
          _2479_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_module.D18.create_DC48(_2478_v38));
          let _2480_v41;
          _2480_v41 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-696)), ((_2481_v0) => function (_2482_i1) {
            return _2481_v0;
          })(_2424_v0)));
          (globalState).f15 = ((_this).f22).multipliedBy(new BigNumber(((_2479_v39).Merge(function () {
            let _coll87 = new _dafny.Map();
            for (const _compr_87 of (_2480_v41).Elements) {
              let _2483_v40 = _compr_87;
              if ((_2480_v41).contains(_2483_v40)) {
                _coll87.push([_2483_v40,_module.D18.create_DC48(_2478_v38)]);
              }
            }
            return _coll87;
          }())).length));
          let _2484_v42;
          _2484_v42 = true;
          _2484_v42 = _2484_v42;
          let _2485_v43;
          _2485_v43 = _dafny.Set.fromElements(p1);
          let _2486_v44;
          _2486_v44 = _dafny.Seq.of((_this).f22, new BigNumber((_2485_v43).length), p2);
          let _2487_v45;
          _2487_v45 = _module.D6.create_DC17((_this).fm1(_2476_v36, new BigNumber((_2486_v44).length), globalState), !(_2484_v42), p1);
          _2484_v42 = (_2487_v45).dtor_cf24;
        }
        let _2488_v46;
        _2488_v46 = _dafny.Set.fromElements(!(p1), p1);
        let _2489_v47;
        _2489_v47 = _dafny.Seq.of(p2, new BigNumber((_2488_v46).length), (_this).f22);
        let _2490_v48;
        _2490_v48 = _dafny.Seq.of(_2489_v47, _2489_v47);
        _2490_v48 = _dafny.Seq.update(_2490_v48, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p0))).length), new BigNumber((_2490_v48).length)), _2489_v47);
        (globalState).f7 = _module.__default.fm17(p0, globalState);
        _2427_v3 = _2427_v3;
        let _2491_v49;
        let _init66 = ((_2492_p1) => function (_2493_i2) {
          return _2492_p1;
        })(p1);
        let _nw401 = Array((new BigNumber(29)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw401.length); _i0_66++) {
          _nw401[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2491_v49 = _nw401;
        let _index414 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_2491_v49).length));
        (_2491_v49)[_index414] = p1;
        let _index415 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_2491_v49).length));
        (_2491_v49)[_index415] = p1;
      } else {
        let _2494_v50;
        let _init67 = ((_2495_p1) => function (_2496_i3) {
          return _2495_p1;
        })(p1);
        let _nw402 = Array((new BigNumber(5)).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw402.length); _i0_67++) {
          _nw402[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2494_v50 = _nw402;
        _2494_v50 = _2494_v50;
        let _2497_v51;
        _2497_v51 = true;
        let _2498_v52;
        let _nw403 = new _module.C11();
        _nw403.__ctor();
        _2498_v52 = _nw403;
        let _2499_v53;
        _2499_v53 = _dafny.Set.fromElements(_2498_v52, _2498_v52);
        let _2500_v54;
        _2500_v54 = _module.D7.create_DC22(p1, p0, p1, new BigNumber((_2499_v53).length));
        let _2501_v55;
        _2501_v55 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2497_v51);
        let _2502_v56;
        _2502_v56 = _module.D7.create_DC22(!((_2500_v54).dtor_cf35), (_dafny.ZERO).minus(p0), (((_2501_v55).contains(_2497_v51)) ? ((_2501_v55).get(_2497_v51)) : (true)), p0);
        _2497_v51 = (_2502_v56).dtor_cf37;
        let _2503_v57;
        _2503_v57 = _dafny.MultiSet.fromElements(_2497_v51);
        _2497_v51 = !(!(_2503_v57).contains(p1));
        let _2504_v58;
        let _nw404 = new _module.C8();
        _nw404.__ctor();
        _2504_v58 = _nw404;
        _2494_v50 = _2494_v50;
      }
      let _2505_v59;
      let _nw405 = new _module.C4();
      _nw405.__ctor();
      _2505_v59 = _nw405;
      let _2506_v61;
      _2506_v61 = _dafny.MultiSet.fromElements((_this).f22, p0, (_this).f22);
      let _2507_v62;
      _2507_v62 = _dafny.Seq.of(function () {
        let _coll88 = new _dafny.Map();
        for (const _compr_88 of (_2506_v61).Elements) {
          let _2508_v60 = _compr_88;
          if ((_2506_v61).contains(_2508_v60)) {
            _coll88.push([_module.__default.safeDivisionInt(_2508_v60, new BigNumber(13)),p1]);
          }
        }
        return _coll88;
      }());
      let _2509_v63;
      _2509_v63 = _module.D14.create_DC35(_2507_v62);
      let _2510_i4;
      _2510_i4 = _dafny.ZERO;
      L16: {
        let _pat_let_tv54 = p1;
        let _pat_let_tv55 = p1;
        let _pat_let_tv56 = p1;
        while (function (_source35) {
          if (_source35.is_DC36) {
            let _2583___mcc_h18 = (_source35).cf53;
            let _2584_cf53 = _2583___mcc_h18;
            return _pat_let_tv54;
          } else if (_source35.is_DC35) {
            let _2585___mcc_h19 = (_source35).cf52;
            let _2586_cf52 = _2585___mcc_h19;
            return _pat_let_tv55;
          } else {
            let _2587___mcc_h20 = (_source35).cf54;
            let _2588_cf54 = _2587___mcc_h20;
            return !(_pat_let_tv56);
          }
        }(_2509_v63)) {
          C16: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2510_i4)) {
              break L16;
            }
            _2510_i4 = (_2510_i4).plus(_dafny.ONE);
            let _2511_v64;
            let _nw406 = new _module.C8();
            _nw406.__ctor();
            _2511_v64 = _nw406;
            let _2512_v65;
            _2512_v65 = _module.D22.create_DC60((_this).f22, (_this).f22);
            let _source34 = _2512_v65;
            if (_source34.is_DC60) {
              let _2513___mcc_h6 = (_source34).cf95;
              let _2514___mcc_h7 = (_source34).cf96;
              let _2515_cf96 = _2514___mcc_h7;
              let _2516_cf95 = _2513___mcc_h6;
              let _2517_v66;
              let _nw407 = new _module.C9();
              _nw407.__ctor();
              _2517_v66 = _nw407;
              let _2518_v67;
              _2518_v67 = true;
              _2518_v67 = true;
              let _2519_v68;
              let _init68 = ((_2520_p1) => function (_2521_i5) {
                return _2520_p1;
              })(p1);
              let _nw408 = Array((_dafny.ONE).toNumber());
              for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw408.length); _i0_68++) {
                _nw408[_i0_68] = _init68(new BigNumber(_i0_68));
              }
              _2519_v68 = _nw408;
              _2519_v68 = _2519_v68;
              (globalState).f7 = _2516_cf95;
            } else if (_source34.is_DC61) {
              let _2522___mcc_h8 = (_source34).cf97;
              let _2523___mcc_h9 = (_source34).cf98;
              let _2524___mcc_h10 = (_source34).cf99;
              let _2525___mcc_h11 = (_source34).cf100;
              let _2526_cf100 = _2525___mcc_h11;
              let _2527_cf99 = _2524___mcc_h10;
              let _2528_cf98 = _2523___mcc_h9;
              let _2529_cf97 = _2522___mcc_h8;
              let _2530_v69;
              _2530_v69 = true;
              _2530_v69 = p1;
              let _2531_v70;
              let _init69 = ((_2532_v0) => function (_2533_i6) {
                return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-459)), ((_2534_v0) => function (_2535_i7) {
                  return _2534_v0;
                })(_2532_v0)), _dafny.Seq.UnicodeFromString("claiu"));
              })(_2424_v0);
              let _nw409 = Array((new BigNumber(19)).toNumber());
              for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw409.length); _i0_69++) {
                _nw409[_i0_69] = _init69(new BigNumber(_i0_69));
              }
              _2531_v70 = _nw409;
              _2531_v70 = _2531_v70;
              let _2536_v71;
              _2536_v71 = _dafny.MultiSet.fromElements(p1);
              let _2537_v72;
              _2537_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2527_cf99,p2);
              let _2538_v73;
              _2538_v73 = _dafny.Seq.of(_2530_v69);
              let _2539_v74;
              _2539_v74 = _dafny.Set.fromElements((_this).f22);
              let _2540_v75;
              _2540_v75 = _module.D22.create_DC62(_2526_cf100, p2, new BigNumber((_2539_v74).length), _2529_cf97, p1);
              let _2541_v76;
              let _nw410 = Array((new BigNumber(23)).toNumber());
              _nw410[(_dafny.ZERO).toNumber()] = (((_2537_v72).contains(p2)) ? ((_2537_v72).get(p2)) : (_2527_cf99));
              _nw410[(_dafny.ONE).toNumber()] = p2;
              _nw410[(new BigNumber(2)).toNumber()] = new BigNumber((_2538_v73).length);
              _nw410[(new BigNumber(3)).toNumber()] = new BigNumber(343);
              _nw410[(new BigNumber(4)).toNumber()] = (_this).f22;
              _nw410[(new BigNumber(5)).toNumber()] = _2528_cf98;
              _nw410[(new BigNumber(6)).toNumber()] = p0;
              _nw410[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_2540_v75)).length);
              _nw410[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(p2);
              _nw410[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_2529_cf97);
              _nw410[(new BigNumber(10)).toNumber()] = _2528_cf98;
              _nw410[(new BigNumber(11)).toNumber()] = p2;
              _nw410[(new BigNumber(12)).toNumber()] = new BigNumber(-744);
              _nw410[(new BigNumber(13)).toNumber()] = _2529_cf97;
              _nw410[(new BigNumber(14)).toNumber()] = p0;
              _nw410[(new BigNumber(15)).toNumber()] = (_this).f22;
              _nw410[(new BigNumber(16)).toNumber()] = (_this).f22;
              _nw410[(new BigNumber(17)).toNumber()] = new BigNumber(661);
              _nw410[(new BigNumber(18)).toNumber()] = p0;
              _nw410[(new BigNumber(19)).toNumber()] = (_this).fm3(p1, globalState);
              _nw410[(new BigNumber(20)).toNumber()] = (_this).f22;
              _nw410[(new BigNumber(21)).toNumber()] = _2528_cf98;
              _nw410[(new BigNumber(22)).toNumber()] = new BigNumber((_2526_cf100).length);
              _2541_v76 = _nw410;
              let _2542_v77;
              _2542_v77 = _module.D2.create_DC5(_2541_v76);
              (_2511_v64).m6(new BigNumber(((_this).f23).length), _2536_v71, ((_2530_v69) ? (p1) : (p1)), _2542_v77, globalState);
              _2424_v0 = _2424_v0;
            } else if (_source34.is_DC62) {
              let _2543___mcc_h12 = (_source34).cf101;
              let _2544___mcc_h13 = (_source34).cf102;
              let _2545___mcc_h14 = (_source34).cf103;
              let _2546___mcc_h15 = (_source34).cf104;
              let _2547___mcc_h16 = (_source34).cf105;
              let _2548_cf105 = _2547___mcc_h16;
              let _2549_cf104 = _2546___mcc_h15;
              let _2550_cf103 = _2545___mcc_h14;
              let _2551_cf102 = _2544___mcc_h13;
              let _2552_cf101 = _2543___mcc_h12;
              let _2553_v78;
              _2553_v78 = _dafny.Map.Empty.slice().updateUnsafe(((p1) ? (p1) : (p1)),(_this).f22);
              _2553_v78 = _2553_v78;
              let _2554_v79;
              let _init70 = ((_2555_p0) => function (_2556_i8) {
                return _module.__default.safeModuloInt(_2556_i8, _2555_p0);
              })(p0);
              let _nw411 = Array((new BigNumber(21)).toNumber());
              for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw411.length); _i0_70++) {
                _nw411[_i0_70] = _init70(new BigNumber(_i0_70));
              }
              _2554_v79 = _nw411;
              let _index416 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_2554_v79).length));
              (_2554_v79)[_index416] = _2551_cf102;
              let _index417 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_2554_v79).length));
              (_2554_v79)[_index417] = new BigNumber(205);
              let _2557_v80;
              _2557_v80 = _module.D2.create_DC5(_2554_v79);
              let _2558_v81;
              _2558_v81 = _module.D2.create_DC8(_2557_v80);
              _2558_v81 = _2558_v81;
              let _2559_v82;
              _2559_v82 = _module.D22.create_DC61(_2551_cf102, new BigNumber((_dafny.Seq.UnicodeFromString("nwpuonys")).length), new BigNumber((_dafny.Seq.UnicodeFromString("goqrmbm")).length), _2552_cf101);
              let _2560_v83;
              _2560_v83 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
              let _2561_v84;
              _2561_v84 = _dafny.Map.Empty.slice().updateUnsafe(_2560_v83,_2552_cf101);
              let _2562_v85;
              _2562_v85 = _dafny.Map.Empty.slice().updateUnsafe((_2554_v79)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_2554_v79).length))],_2552_cf101);
              let _2563_v86;
              let _nw412 = Array((new BigNumber(18)).toNumber());
              _nw412[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2552_cf101, (_this).f23);
              _nw412[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((_this).f23, _2552_cf101);
              _nw412[(new BigNumber(2)).toNumber()] = (_this).f23;
              _nw412[(new BigNumber(3)).toNumber()] = (_2559_v82).dtor_cf100;
              _nw412[(new BigNumber(4)).toNumber()] = (_this).f23;
              _nw412[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_2559_v82).dtor_cf100, _2552_cf101);
              _nw412[(new BigNumber(6)).toNumber()] = _2552_cf101;
              _nw412[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_2552_cf101, _module.__default.safeIndex((_2554_v79)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_2554_v79).length))], new BigNumber((_2552_cf101).length)), new _dafny.CodePoint('o'.codePointAt(0)));
              _nw412[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_2552_cf101, _module.__default.safeIndex(new BigNumber(-44), new BigNumber((_2552_cf101).length)), _2424_v0);
              _nw412[(new BigNumber(9)).toNumber()] = _2552_cf101;
              _nw412[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat((((_2561_v84).contains(_2560_v83)) ? ((_2561_v84).get(_2560_v83)) : ((_this).f23)), (_this).f23);
              _nw412[(new BigNumber(11)).toNumber()] = (_this).f23;
              _nw412[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("s");
              _nw412[(new BigNumber(13)).toNumber()] = (_this).f23;
              _nw412[(new BigNumber(14)).toNumber()] = _2552_cf101;
              _nw412[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("jfd");
              _nw412[(new BigNumber(16)).toNumber()] = (((_2562_v85).contains((_2511_v64).fm2((_dafny.ZERO).minus(_2550_cf103), globalState))) ? ((_2562_v85).get((_2511_v64).fm2((_dafny.ZERO).minus(_2550_cf103), globalState))) : ((_this).f23));
              _nw412[(new BigNumber(17)).toNumber()] = (_this).f23;
              _2563_v86 = _nw412;
              let _index418 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2563_v86).length));
              (_2563_v86)[_index418] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bnimhcmi"), (_this).f23);
              let _2564_v87;
              _2564_v87 = _dafny.Seq.of(_2548_cf105);
              let _2565_v88;
              let _init71 = function (_2566_i9) {
                return false;
              };
              let _nw413 = Array((new BigNumber(19)).toNumber());
              for (let _i0_71 = 0; _i0_71 < new BigNumber(_nw413.length); _i0_71++) {
                _nw413[_i0_71] = _init71(new BigNumber(_i0_71));
              }
              _2565_v88 = _nw413;
              let _2567_v89;
              _2567_v89 = _module.D0.create_DC1(_2565_v88);
              let _2568_v90;
              _2568_v90 = _dafny.Map.Empty.slice().updateUnsafe(_2567_v89,_2564_v87);
              let _2569_v91;
              let _nw414 = new _module.C5();
              _nw414.__ctor();
              _2569_v91 = _nw414;
              let _2570_v92;
              _2570_v92 = _dafny.Map.Empty.slice().updateUnsafe(_2569_v91,_2564_v87);
              let _2571_v93;
              let _nw415 = Array((new BigNumber(26)).toNumber());
              _nw415[(_dafny.ZERO).toNumber()] = _2564_v87;
              _nw415[(_dafny.ONE).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_2548_cf105, _2548_cf105);
              _nw415[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_2548_cf105);
              _nw415[(new BigNumber(4)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((_2505_v59).fm8(_2549_cf104, _2548_cf105, new BigNumber(302), globalState));
              _nw415[(new BigNumber(6)).toNumber()] = (_module.D14.create_DC36(_dafny.Seq.of(p1, _2548_cf105))).dtor_cf53;
              _nw415[(new BigNumber(7)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(p1);
              _nw415[(new BigNumber(9)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(10)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(11)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(12)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(13)).toNumber()] = (((_2568_v90).contains(_2567_v89)) ? ((_2568_v90).get(_2567_v89)) : (_2564_v87));
              _nw415[(new BigNumber(14)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_2564_v87, _module.__default.safeIndex(_2549_cf104, new BigNumber((_2564_v87).length)), _2548_cf105);
              _nw415[(new BigNumber(16)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(17)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(18)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(19)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(20)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(21)).toNumber()] = _module.__default.fm27(globalState);
              _nw415[(new BigNumber(22)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(23)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(24)).toNumber()] = _2564_v87;
              _nw415[(new BigNumber(25)).toNumber()] = (((_2570_v92).contains(_2569_v91)) ? ((_2570_v92).get(_2569_v91)) : (_2564_v87));
              _2571_v93 = _nw415;
              let _2572_v94;
              _2572_v94 = _dafny.Seq.of(_2571_v93, _2571_v93, _2571_v93);
              let _2573_v95;
              let _nw416 = Array((new BigNumber(17)).toNumber());
              _nw416[(_dafny.ZERO).toNumber()] = _2571_v93;
              _nw416[(_dafny.ONE).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(2)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(3)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(4)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(5)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(6)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(7)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(8)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(9)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(10)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(11)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(12)).toNumber()] = (_2572_v94)[_module.__default.safeIndex(_2550_cf103, new BigNumber((_2572_v94).length))];
              _nw416[(new BigNumber(13)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(14)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(15)).toNumber()] = _2571_v93;
              _nw416[(new BigNumber(16)).toNumber()] = _2571_v93;
              _2573_v95 = _nw416;
              let _index419 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_2573_v95).length));
              (_2573_v95)[_index419] = _2571_v93;
              let _index420 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2563_v86).length));
              let _index421 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_2573_v95).length));
              let _rhs486 = (_this).f23;
              let _rhs487 = ((p1) ? (_2571_v93) : (_2571_v93));
              let _rhs488 = (_dafny.ZERO).minus((_2549_cf104).minus((_2554_v79)[_module.__default.safeIndex(new BigNumber(350), new BigNumber((_2554_v79).length))]));
              let _rhs489 = _2424_v0;
              let _lhs327 = _2563_v86;
              let _lhs328 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_2563_v86).length));
              let _lhs329 = _2573_v95;
              let _lhs330 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_2573_v95).length));
              let _lhs331 = globalState;
              _lhs327[_lhs328] = _rhs486;
              _lhs329[_lhs330] = _rhs487;
              _lhs331.f15 = _rhs488;
              _2424_v0 = _rhs489;
            } else {
              let _2574___mcc_h17 = (_source34).cf94;
              let _2575_cf94 = _2574___mcc_h17;
              let _2576_v96;
              _2576_v96 = _module.D8.create_DC25(((_dafny.ZERO).minus(p2)).multipliedBy((_this).f22));
              _2576_v96 = _2576_v96;
              (globalState).f13 = ((p1) ? (p2) : (p0));
              let _2577_v97;
              _2577_v97 = false;
              _2577_v97 = _2577_v97;
              let _2578_v98;
              _2578_v98 = _dafny.Seq.UnicodeFromString("vdybm");
              let _rhs490 = (_this).f23;
              let _rhs491 = ((_2506_v61).Intersect(_module.__default.fm31((_this).f22, (_this).f22, globalState))).Difference((_dafny.MultiSet.fromElements((_this).f22, p0, new BigNumber(-8))).Union(_dafny.MultiSet.fromElements(new BigNumber((_2578_v98).length))));
              _2578_v98 = _rhs490;
              _2506_v61 = _rhs491;
            }
            let _2579_v99;
            _2579_v99 = true;
            _2579_v99 = _2579_v99;
            let _2580_v100;
            _2580_v100 = _dafny.Seq.UnicodeFromString("kkiv");
            _2580_v100 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-119)), ((_2581_v0) => function (_2582_i10) {
              return _2581_v0;
            })(_2424_v0));
          }
        }
      }
      let _2589_v101;
      let _nw417 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _2589_v101 = _nw417;
      let _2590_v102;
      _2590_v102 = _dafny.Seq.UnicodeFromString("fwc");
      let _rhs492 = _2589_v101;
      let _rhs493 = (_this).f23;
      _2589_v101 = _rhs492;
      _2590_v102 = _rhs493;
      let _2591_v103;
      _2591_v103 = _module.D6.create_DC17(p1, true, p1);
      let _2592_i11;
      _2592_i11 = _dafny.ZERO;
      L17: {
        while ((((p1) ? (_2591_v103) : (_2591_v103))).dtor_cf26) {
          C17: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2592_i11)) {
              break L17;
            }
            _2592_i11 = (_2592_i11).plus(_dafny.ONE);
            let _2593_v104;
            _2593_v104 = _dafny.MultiSet.fromElements(p1, !(p1), ((p1) ? (p1) : (p1)), false, p1);
            let _2594_v105;
            _2594_v105 = _dafny.Set.fromElements(!(!(p1)));
            let _rhs494 = (_dafny.MultiSet.fromElements(p1)).Intersect(_dafny.MultiSet.fromElements(p1, p1));
            let _rhs495 = _2589_v101;
            let _rhs496 = (_2594_v105).Union(_dafny.Set.fromElements(p1));
            _2593_v104 = _rhs494;
            _2589_v101 = _rhs495;
            _2594_v105 = _rhs496;
            if (false) {
              (globalState).f13 = p0;
              let _2595_v106;
              let _init72 = ((_2596_p1) => function (_2597_i12) {
                return _2596_p1;
              })(p1);
              let _nw418 = Array((new BigNumber(11)).toNumber());
              for (let _i0_72 = 0; _i0_72 < new BigNumber(_nw418.length); _i0_72++) {
                _nw418[_i0_72] = _init72(new BigNumber(_i0_72));
              }
              _2595_v106 = _nw418;
              let _2598_v107;
              _2598_v107 = _dafny.Map.Empty.slice().updateUnsafe(_2594_v105,p1);
              let _index422 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2595_v106).length));
              (_2595_v106)[_index422] = ((_2505_v59).fm8(new BigNumber((_2598_v107).length), p1, (_dafny.ZERO).minus(p2), globalState)) || (p1);
              let _2599_v108;
              _2599_v108 = true;
              let _2600_v109;
              _2600_v109 = _dafny.Seq.of(p2, p0);
              let _2601_v110;
              _2601_v110 = _module.D1.create_DC3(_dafny.Seq.UnicodeFromString("ohan"), false);
              let _index423 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2595_v106).length));
              let _rhs497 = (_module.__default.safeDivisionInt((_this).f22, new BigNumber((_2590_v102).length))).multipliedBy(new BigNumber((_2600_v109).length));
              let _rhs498 = true;
              let _rhs499 = _module.__default.fm30(p2, ((_this).f22).isLessThanOrEqualTo(p2), (_this).fm4(_2590_v102, p2, _2599_v108, _2601_v110, globalState), _2599_v108, globalState);
              let _rhs500 = (p0).isLessThanOrEqualTo(p0);
              let _rhs501 = p1;
              let _lhs332 = globalState;
              let _lhs333 = _2595_v106;
              let _lhs334 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_2595_v106).length));
              _lhs332.f13 = _rhs497;
              _lhs333[_lhs334] = _rhs498;
              _2590_v102 = _rhs499;
              _2599_v108 = _rhs500;
              _2599_v108 = _rhs501;
              let _2602_v111;
              _2602_v111 = _module.D14.create_DC35(_2507_v62);
              let _2603_v112;
              _2603_v112 = _module.D14.create_DC37(_2602_v111);
              let _2604_v113;
              _2604_v113 = _module.D14.create_DC37(_2602_v111);
              let _2605_v114;
              _2605_v114 = _module.D14.create_DC37(_2603_v112);
              let _2606_v115;
              _2606_v115 = _dafny.Map.Empty.slice().updateUnsafe(_2605_v114,p2);
              let _2607_v116;
              let _nw419 = new _module.C9();
              _nw419.__ctor();
              _2607_v116 = _nw419;
              let _2608_v117;
              _2608_v117 = _module.D17.create_DC47(_2607_v116, _2599_v108, new BigNumber((_2506_v61).cardinality()), (_2607_v116).fm2(p2, globalState), (_this).f22);
              let _2609_v118;
              _2609_v118 = _dafny.Map.Empty.slice().updateUnsafe(_2599_v108,new BigNumber(708));
              (globalState).f15 = (_module.__default.safeDivisionInt((((_2606_v115).contains(_2605_v114)) ? ((_2606_v115).get(_2605_v114)) : (p0)), (_this).f22)).minus(((_2608_v117).dtor_cf72).minus((((_2609_v118).contains(_2599_v108)) ? ((_2609_v118).get(_2599_v108)) : (p2))));
              let _index424 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_2589_v101).length));
              (_2589_v101)[_index424] = p2;
              let _index425 = _module.__default.safeIndex(new BigNumber(368), new BigNumber((_2589_v101).length));
              (_2589_v101)[_index425] = p0;
              let _2610_v119;
              let _nw420 = new _module.C9();
              _nw420.__ctor();
              _2610_v119 = _nw420;
              let _2611_v120;
              _2611_v120 = _dafny.MultiSet.fromElements(_2610_v119, _2610_v119, _2610_v119, _2610_v119, _2610_v119);
              _2611_v120 = _dafny.MultiSet.fromElements(_2610_v119, _2610_v119);
            } else {
              let _2612_v121;
              _2612_v121 = _dafny.Seq.of(p1);
              _2612_v121 = ((p1) ? (_2612_v121) : (_2612_v121));
              let _2613_v122;
              _2613_v122 = false;
              let _rhs502 = !(_2613_v122) || (_2613_v122);
              let _rhs503 = _2593_v104;
              _2613_v122 = _rhs502;
              _2593_v104 = _rhs503;
              let _2614_v123;
              _2614_v123 = _dafny.Seq.of((_this).f23);
              _2590_v102 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update((_2614_v123)[_module.__default.safeIndex((_this).f22, new BigNumber((_2614_v123).length))], _module.__default.safeIndex(p2, new BigNumber(((_2614_v123)[_module.__default.safeIndex((_this).f22, new BigNumber((_2614_v123).length))]).length)), _2424_v0), _module.__default.safeIndex((_this).f22, new BigNumber((_dafny.Seq.update((_2614_v123)[_module.__default.safeIndex((_this).f22, new BigNumber((_2614_v123).length))], _module.__default.safeIndex(p2, new BigNumber(((_2614_v123)[_module.__default.safeIndex((_this).f22, new BigNumber((_2614_v123).length))]).length)), _2424_v0)).length)), _2424_v0), _2590_v102), _dafny.Seq.UnicodeFromString("rxdtrqap"));
              _2589_v101 = _2589_v101;
              let _index426 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2589_v101).length));
              (_2589_v101)[_index426] = (_dafny.ZERO).minus(p0);
              let _index427 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_2589_v101).length));
              (_2589_v101)[_index427] = p0;
            }
            (globalState).f13 = new BigNumber((_module.__default.fm59(_2424_v0, p1, globalState)).length);
            let _2615_v124;
            let _nw421 = new _module.C1();
            _nw421.__ctor();
            _2615_v124 = _nw421;
          }
        }
      }
      let _2616_v125;
      let _nw422 = new _module.C1();
      _nw422.__ctor();
      _2616_v125 = _nw422;
      return;
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };

  $module.C13 = class C13 {
    constructor () {
      this._tname = "_module.C13";
      this.f21 = false;
      this._f20 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f20, f21) {
      let _this = this;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Map.Empty.slice().updateUnsafe((_this).f20,(_this).f20)).Merge((function () {
        let _coll89 = new _dafny.Map();
        for (const _compr_89 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(472))) {
          let _2617_v0 = _compr_89;
          if (((new BigNumber(180)).isLessThanOrEqualTo(_2617_v0)) && ((_2617_v0).isLessThan(new BigNumber(472)))) {
            _coll89.push([_module.__default.safeModuloInt(_2617_v0, (_this).f20),(_this).f20]);
          }
        }
        return _coll89;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f20,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f21,_this.f21)).length))));
    };
    m0(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      if (p2) {
        let _2618_v0;
        let _nw423 = Array((new BigNumber(16)).toNumber()).fill(false);
        _2618_v0 = _nw423;
        let _2619_v1;
        _2619_v1 = _module.D0.create_DC0(_2618_v0);
        let _2620_v2;
        _2620_v2 = _dafny.Seq.of(_2618_v0, _2618_v0, _2618_v0, _2618_v0);
        let _pat_let_tv57 = _2618_v0;
        let _2621_v3;
        _2621_v3 = _dafny.Seq.of(_2618_v0, (function (_pat_let51_0) {
          return function (_2622_dt__update__tmp_h0) {
            return function (_pat_let52_0) {
              return function (_2623_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_2623_dt__update_hcf0_h0);
              }(_pat_let52_0);
            }(_pat_let_tv57);
          }(_pat_let51_0);
        }(_2619_v1)).dtor_cf0, (_2620_v2)[_module.__default.safeIndex((_module.D1.create_DC2(new BigNumber((_dafny.MultiSet.fromElements((_this).f20)).cardinality()))).dtor_cf2, new BigNumber((_2620_v2).length))], _2618_v0);
        _2621_v3 = _dafny.Seq.Concat(_2620_v2, _2621_v3);
        let _2624_v4;
        _2624_v4 = new _dafny.CodePoint('c'.codePointAt(0));
        let _2625_v5;
        _2625_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2624_v4,_2624_v4);
        _2625_v5 = (_2625_v5).update(_2624_v4, _2624_v4);
        let _2626_v6;
        let _nw424 = new _module.C7();
        _nw424.__ctor();
        _2626_v6 = _nw424;
        (_this).f21 = _this.f21;
        let _2627_v7;
        _2627_v7 = _dafny.Seq.UnicodeFromString("dj");
        let _2628_v8;
        let _nw425 = new _module.C12();
        _nw425.__ctor(new BigNumber(-287), _2627_v7);
        _2628_v8 = _nw425;
      } else {
        (globalState).f13 = (_dafny.ZERO).minus((_this).f20);
        let _2629_v9;
        let _nw426 = new _module.C2();
        _nw426.__ctor();
        _2629_v9 = _nw426;
        let _2630_v10;
        _2630_v10 = _dafny.MultiSet.fromElements(false, false, p1, p2);
        let _2631_v11;
        _2631_v11 = _dafny.Set.fromElements(p2, p2);
        let _2632_v12;
        _2632_v12 = _dafny.Seq.of(_this.f21, p1);
        _2630_v10 = _module.__default.fm34((new BigNumber((_2631_v11).length)).multipliedBy((_this).f20), new BigNumber((_dafny.MultiSet.FromArray(_2632_v12)).cardinality()), (_this.f21) || ((_2632_v12)[_module.__default.safeIndex((_this).f20, new BigNumber((_2632_v12).length))]), globalState);
        let _2633_v13;
        let _init73 = function (_2634_i0) {
          return (_2634_i0).plus(new BigNumber(439));
        };
        let _nw427 = Array((new BigNumber(13)).toNumber());
        for (let _i0_73 = 0; _i0_73 < new BigNumber(_nw427.length); _i0_73++) {
          _nw427[_i0_73] = _init73(new BigNumber(_i0_73));
        }
        _2633_v13 = _nw427;
        let _2635_v14;
        _2635_v14 = _dafny.Set.fromElements(_2633_v13, _2633_v13, _2633_v13, _2633_v13);
        (_this).f21 = ((_2635_v14).Intersect(_2635_v14)).IsProperSubsetOf((_dafny.Set.fromElements(_2633_v13, _2633_v13)).Intersect(_2635_v14));
        let _2636_v15;
        _2636_v15 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(978)).isEqualTo(new BigNumber((_2632_v12).length)),(_2631_v11).Intersect(_2631_v11));
        _2636_v15 = _2636_v15;
      }
      let _2637_v16;
      _2637_v16 = _dafny.Seq.of(p1);
      let _2638_v17;
      _2638_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f21,_dafny.Seq.of(false, p1, p1, p0, p0));
      let _2639_v18;
      let _nw428 = Array((new BigNumber(6)).toNumber());
      _nw428[(_dafny.ZERO).toNumber()] = _2637_v16;
      _nw428[(_dafny.ONE).toNumber()] = _2637_v16;
      _nw428[(new BigNumber(2)).toNumber()] = (((_2638_v17).contains(_this.f21)) ? ((_2638_v17).get(_this.f21)) : (_2637_v16));
      _nw428[(new BigNumber(3)).toNumber()] = _module.__default.fm27(globalState);
      _nw428[(new BigNumber(4)).toNumber()] = (_module.D8.create_DC23(_2637_v16)).dtor_cf39;
      _nw428[(new BigNumber(5)).toNumber()] = _2637_v16;
      _2639_v18 = _nw428;
      let _2640_v19;
      _2640_v19 = _dafny.MultiSet.fromElements((_this).f20, (_this).f20, (_this).f20);
      let _2641_v20;
      _2641_v20 = _dafny.MultiSet.fromElements(new BigNumber(((_2640_v19).update((_this).f20, _module.__default.abs((_this).f20))).cardinality()));
      let _2642_v21;
      _2642_v21 = _dafny.Seq.of(_2640_v19, _2641_v20);
      let _2643_v22;
      _2643_v22 = _dafny.Seq.UnicodeFromString("onoohtgc");
      let _2644_v23;
      _2644_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,false);
      let _2645_v24;
      let _nw429 = Array((new BigNumber(22)).toNumber());
      _nw429[(_dafny.ZERO).toNumber()] = p2;
      _nw429[(_dafny.ONE).toNumber()] = p1;
      _nw429[(new BigNumber(2)).toNumber()] = _module.__default.fm26(globalState);
      _nw429[(new BigNumber(3)).toNumber()] = _dafny.areEqual(_2642_v21, _2642_v21);
      _nw429[(new BigNumber(4)).toNumber()] = p2;
      _nw429[(new BigNumber(5)).toNumber()] = p1;
      _nw429[(new BigNumber(6)).toNumber()] = _this.f21;
      _nw429[(new BigNumber(7)).toNumber()] = !(!(p2) || (p0));
      _nw429[(new BigNumber(8)).toNumber()] = p2;
      _nw429[(new BigNumber(9)).toNumber()] = p0;
      _nw429[(new BigNumber(10)).toNumber()] = p0;
      _nw429[(new BigNumber(11)).toNumber()] = (_2637_v16)[_module.__default.safeIndex(new BigNumber((_2643_v22).length), new BigNumber((_2637_v16).length))];
      _nw429[(new BigNumber(12)).toNumber()] = _module.__default.fm26(globalState);
      _nw429[(new BigNumber(13)).toNumber()] = ((!(p2)) ? (p2) : ((((_2644_v23).contains((_this).f20)) ? ((_2644_v23).get((_this).f20)) : (p1))));
      _nw429[(new BigNumber(14)).toNumber()] = p2;
      _nw429[(new BigNumber(15)).toNumber()] = p0;
      _nw429[(new BigNumber(16)).toNumber()] = _this.f21;
      _nw429[(new BigNumber(17)).toNumber()] = true;
      _nw429[(new BigNumber(18)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_2643_v22, _2643_v22);
      _nw429[(new BigNumber(19)).toNumber()] = p2;
      _nw429[(new BigNumber(20)).toNumber()] = p2;
      _nw429[(new BigNumber(21)).toNumber()] = (p2) && (_this.f21);
      _2645_v24 = _nw429;
      let _index428 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2645_v24).length));
      (_2645_v24)[_index428] = ((_this.f21) ? (p0) : (_this.f21));
      let _2646_v25;
      _2646_v25 = new _dafny.CodePoint('l'.codePointAt(0));
      let _pat_let_tv58 = _2643_v22;
      let _pat_let_tv59 = _2643_v22;
      let _index429 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2645_v24).length));
      let _rhs504 = function (_source36) {
        if (_source36.is_DC20) {
          let _2647___mcc_h0 = (_source36).cf29;
          let _2648_cf29 = _2647___mcc_h0;
          return new BigNumber(881);
        } else if (_source36.is_DC21) {
          let _2649___mcc_h1 = (_source36).cf30;
          let _2650___mcc_h2 = (_source36).cf31;
          let _2651___mcc_h3 = (_source36).cf32;
          let _2652___mcc_h4 = (_source36).cf33;
          let _2653___mcc_h5 = (_source36).cf34;
          let _2654_cf34 = _2653___mcc_h5;
          let _2655_cf33 = _2652___mcc_h4;
          let _2656_cf32 = _2651___mcc_h3;
          let _2657_cf31 = _2650___mcc_h2;
          let _2658_cf30 = _2649___mcc_h1;
          return new BigNumber(-379);
        } else if (_source36.is_DC22) {
          let _2659___mcc_h6 = (_source36).cf35;
          let _2660___mcc_h7 = (_source36).cf36;
          let _2661___mcc_h8 = (_source36).cf37;
          let _2662___mcc_h9 = (_source36).cf38;
          let _2663_cf38 = _2662___mcc_h9;
          let _2664_cf37 = _2661___mcc_h8;
          let _2665_cf36 = _2660___mcc_h7;
          let _2666_cf35 = _2659___mcc_h6;
          return new BigNumber((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(170)), function (_2667_i1) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("ug"), _pat_let_tv58, (_module.D22.create_DC62(_pat_let_tv59, (_this).f20, _2663_cf38, _2663_cf38, false)).dtor_cf101)).length);
        } else {
          let _2668___mcc_h10 = (_source36).cf28;
          let _2669_cf28 = _2668___mcc_h10;
          return (_this).f20;
        }
      }(_module.D7.create_DC21((_this).f20, (_this).f20, (_2637_v16)[_module.__default.safeIndex((_this).f20, new BigNumber((_2637_v16).length))], _this.f21, false));
      let _rhs505 = _2639_v18;
      let _rhs506 = !(_dafny.Seq.contains(_2643_v22, _2646_v25));
      let _rhs507 = (_this).f20;
      let _lhs335 = globalState;
      let _lhs336 = _2645_v24;
      let _lhs337 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_2645_v24).length));
      let _lhs338 = globalState;
      _lhs335.f15 = _rhs504;
      _2639_v18 = _rhs505;
      _lhs336[_lhs337] = _rhs506;
      _lhs338.f7 = _rhs507;
      let _2670_v26;
      _2670_v26 = _dafny.MultiSet.fromElements(_this.f21);
      let _2671_v27;
      _2671_v27 = _dafny.Seq.of(new BigNumber((_2670_v26).cardinality()));
      let _2672_v28;
      _2672_v28 = _dafny.Map.Empty.slice().updateUnsafe(p2,_this.f21);
      let _2673_v29;
      _2673_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2671_v27).length),new BigNumber((_2672_v28).length));
      let _2674_v30;
      _2674_v30 = _module.D18.create_DC48(_2673_v29);
      let _2675_v31;
      _2675_v31 = _dafny.Seq.of(_2674_v30, _2674_v30);
      let _2676_v32;
      let _nw430 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _2676_v32 = _nw430;
      let _2677_v33;
      _2677_v33 = _module.D2.create_DC7(p2, _dafny.MultiSet.fromElements(true, !(p0)), _2676_v32);
      let _2678_v34;
      _2678_v34 = _module.D2.create_DC8(_2677_v33);
      let _2679_v35;
      let _nw431 = new _module.C3();
      _nw431.__ctor(_2678_v34, false);
      _2679_v35 = _nw431;
      let _2680_v36;
      _2680_v36 = _dafny.Set.fromElements(_2679_v35);
      let _2681_v37;
      _2681_v37 = _dafny.Set.fromElements(_2679_v35.f25);
      let _2682_v38;
      _2682_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f20,_2637_v16);
      let _2683_v39;
      _2683_v39 = _dafny.Map.Empty.slice().updateUnsafe(_2646_v25,_2682_v38);
      let _2684_v40;
      let _nw432 = Array((new BigNumber(23)).toNumber());
      _nw432[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_2643_v22).length))).length);
      _nw432[(_dafny.ONE).toNumber()] = (((_2670_v26).contains(p1)) ? ((_2670_v26).get(p1)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_2675_v31)).length)));
      _nw432[(new BigNumber(2)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(3)).toNumber()] = new BigNumber((_2680_v36).length);
      _nw432[(new BigNumber(4)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(5)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(6)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(7)).toNumber()] = new BigNumber((_2641_v20).cardinality());
      _nw432[(new BigNumber(8)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(9)).toNumber()] = new BigNumber((_2671_v27).length);
      _nw432[(new BigNumber(10)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(11)).toNumber()] = new BigNumber((_2681_v37).length);
      _nw432[(new BigNumber(12)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(13)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(14)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(15)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(16)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(17)).toNumber()] = new BigNumber((_2672_v28).length);
      _nw432[(new BigNumber(18)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(19)).toNumber()] = new BigNumber(((((_2683_v39).contains(_module.__default.fm38(_2679_v35.f25, globalState))) ? ((_2683_v39).get(_module.__default.fm38(_2679_v35.f25, globalState))) : (_2682_v38))).length);
      _nw432[(new BigNumber(20)).toNumber()] = new BigNumber(587);
      _nw432[(new BigNumber(21)).toNumber()] = (_this).f20;
      _nw432[(new BigNumber(22)).toNumber()] = _module.__default.fm17((_this).f20, globalState);
      _2684_v40 = _nw432;
      let _2685_v41;
      _2685_v41 = _module.D2.create_DC5(_2676_v32);
      _2685_v41 = _2685_v41;
      let _2686_v42;
      let _nw433 = new _module.C10();
      _nw433.__ctor();
      _2686_v42 = _nw433;
      let _2687_v43;
      _2687_v43 = _dafny.Set.fromElements(new BigNumber(190), (_this).f20);
      let _2688_v44;
      let _nw434 = Array((new BigNumber(25)).toNumber());
      _nw434[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("uxr");
      _nw434[(_dafny.ONE).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("hlbyiwl");
      _nw434[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_2643_v22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_2689_v25) => function (_2690_i2) {
        return _2689_v25;
      })(_2646_v25)));
      _nw434[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(488)), ((_2691_v25) => function (_2692_i3) {
        return _2691_v25;
      })(_2646_v25));
      _nw434[(new BigNumber(5)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(6)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2643_v22, _2643_v22);
      _nw434[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_2643_v22, _module.__default.safeIndex((_this).f20, new BigNumber((_2643_v22).length)), _2646_v25);
      _nw434[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2643_v22, _2643_v22);
      _nw434[(new BigNumber(10)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(858)), function (_2693_i4) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      });
      _nw434[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(757)), ((_2694_v25) => function (_2695_i5) {
        return _2694_v25;
      })(_2646_v25));
      _nw434[(new BigNumber(13)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-511)), ((_2696_v25) => function (_2697_i6) {
        return _2696_v25;
      })(_2646_v25)), _2643_v22);
      _nw434[(new BigNumber(15)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(30)), ((_2698_v25) => function (_2699_i7) {
        return _2698_v25;
      })(_2646_v25)), _2643_v22);
      _nw434[(new BigNumber(17)).toNumber()] = _module.__default.fm30((_this).f20, false, _2687_v43, p0, globalState);
      _nw434[(new BigNumber(18)).toNumber()] = (((_2645_v24)[_module.__default.safeIndex(new BigNumber(895), new BigNumber((_2645_v24).length))]) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), ((_2700_v25) => function (_2701_i8) {
        return _2700_v25;
      })(_2646_v25))) : (_2643_v22));
      _nw434[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("ihj");
      _nw434[(new BigNumber(20)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(21)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xfq"), _2643_v22);
      _nw434[(new BigNumber(23)).toNumber()] = _2643_v22;
      _nw434[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), ((_2702_v25) => function (_2703_i9) {
        return _2702_v25;
      })(_2646_v25)), _2643_v22);
      _2688_v44 = _nw434;
      let _index430 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_2688_v44).length));
      (_2688_v44)[_index430] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("idcvyt"), _dafny.Seq.update(_2643_v22, _module.__default.safeIndex((_this).f20, new BigNumber((_2643_v22).length)), _module.__default.fm38(_this.f21, globalState)));
      let _index431 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_2688_v44).length));
      (_2688_v44)[_index431] = _dafny.Seq.Concat(_2643_v22, _2643_v22);
      let _2704_v45;
      let _init74 = ((_2705_v23) => function (_2706_i11) {
        return _2705_v23;
      })(_2644_v23);
      let _nw435 = Array((new BigNumber(8)).toNumber());
      for (let _i0_74 = 0; _i0_74 < new BigNumber(_nw435.length); _i0_74++) {
        _nw435[_i0_74] = _init74(new BigNumber(_i0_74));
      }
      _2704_v45 = _nw435;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2704_v45).length))) {
        let _2707_i10 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2707_i10)) && ((_2707_i10).isLessThan(new BigNumber((_2704_v45).length))))) {
          (_2704_v45)[(_2707_i10)] = _2644_v23;
        }
      }
      r0 = _2646_v25;
      return r0;
    }
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
