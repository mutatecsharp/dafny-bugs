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
      return !(((new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).length)).multipliedBy(new BigNumber(((_module.D6.create_DC14(_dafny.Seq.of(true, false))).dtor_cf20).length))).isLessThanOrEqualTo(new BigNumber(567)));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of (_dafny.Seq.of(new BigNumber(151))).Elements) {
            let _0_v1 = _compr_1;
            if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(151)), _0_v1)) {
              _coll1.push([(_0_v1).minus(new BigNumber(829)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), function (_1_i0) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              })).length))).length), new BigNumber(523))).length)]);
            }
          }
          return _coll1;
        }()).Keys.Elements) {
          let _2_v0 = _compr_0;
          if ((function () {
            let _coll2 = new _dafny.Map();
            for (const _compr_2 of (_dafny.Seq.of(new BigNumber(151))).Elements) {
              let _0_v1 = _compr_2;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(151)), _0_v1)) {
                _coll2.push([(_0_v1).minus(new BigNumber(829)),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), function (_1_i0) {
                  return new _dafny.CodePoint('y'.codePointAt(0));
                })).length))).length), new BigNumber(523))).length)]);
              }
            }
            return _coll2;
          }()).contains(_2_v0)) {
            _coll0.push([(_2_v0).minus(new BigNumber(542)),new BigNumber(921)]);
          }
        }
        return _coll0;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(627),new BigNumber(464)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(638)))).length),new BigNumber(126)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-220),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(487)), function (_3_i1) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(false))).length);
      })).length)))).length), new BigNumber(-511))).plus(new BigNumber(535));
    };
    static fm4(globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber(462))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("rnvmueh")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(614),_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), function (_4_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })))).length))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(239))).length)), new BigNumber(-650), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-585)), _dafny.Seq.of(new BigNumber(913), new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("f")).length)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(301),false)).length)))).length)))));
    };
    static fm5(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-620),new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).length))).length),true))).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(963), new BigNumber(663))) {
          let _5_v0 = _compr_3;
          if (((new BigNumber(963)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(663)))) {
            _coll3.push([_module.__default.safeModuloInt(_5_v0, new BigNumber(441)),false]);
          }
        }
        return _coll3;
      }());
    };
    static fm6(p0, p1, globalState) {
      return (((true) ? (_dafny.MultiSet.fromElements(false, false, true)) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false, false, !(false), false))))).Union(_dafny.MultiSet.fromElements(!(true)));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,false);
    };
    static fm8(p0, globalState) {
      return _dafny.Set.fromElements(((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qjl")).length),new BigNumber(-53))).length))).isLessThanOrEqualTo(new BigNumber(90)));
    };
    static fm9(globalState) {
      return _dafny.Seq.UnicodeFromString("usw");
    };
    static fm10(p0, p1, p2, globalState) {
      return _dafny.Seq.of(!(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(335), new BigNumber(539))) {
          let _6_v0 = _compr_4;
          if (((new BigNumber(335)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(539)))) {
            _coll4.push([(_6_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length))),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_7_i0) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length))]);
          }
        }
        return _coll4;
      }()).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(186),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))), !_dafny.Seq.contains(_dafny.Seq.of(true, true, !(true), true, true), true));
    };
    static fm11(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xqmuuiom"), _dafny.Seq.UnicodeFromString("sjc"), _dafny.Seq.UnicodeFromString("f"))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ofdbvnl")))).Union((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("bc"))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), function (_8_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(42)), function (_9_i1) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        });
      }))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('v'.codePointAt(0))))).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('c'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('s'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('s'.codePointAt(0)))));
    };
    static m0(globalState) {
      let _10_v0;
      _10_v0 = false;
      let _11_v1;
      _11_v1 = new BigNumber(-798);
      let _12_v2;
      _12_v2 = _dafny.Map.Empty.slice().updateUnsafe(_10_v0,_10_v0);
      let _13_v3;
      _13_v3 = _dafny.Seq.of(!(_10_v0));
      let _14_v4;
      _14_v4 = _dafny.Set.fromElements(_11_v1, (_dafny.ZERO).minus(_11_v1));
      let _15_v5;
      _15_v5 = _dafny.Set.fromElements(_10_v0);
      let _16_v6;
      let _nw0 = Array((new BigNumber(22)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _11_v1;
      _nw0[(_dafny.ONE).toNumber()] = new BigNumber((_12_v2).length);
      _nw0[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_11_v1);
      _nw0[(new BigNumber(3)).toNumber()] = new BigNumber((_13_v3).length);
      _nw0[(new BigNumber(4)).toNumber()] = new BigNumber((_14_v4).length);
      _nw0[(new BigNumber(5)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(6)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(7)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(8)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(9)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(10)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(11)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(12)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(13)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(14)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(15)).toNumber()] = new BigNumber((_15_v5).length);
      _nw0[(new BigNumber(16)).toNumber()] = new BigNumber((_13_v3).length);
      _nw0[(new BigNumber(17)).toNumber()] = new BigNumber(-292);
      _nw0[(new BigNumber(18)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(19)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(20)).toNumber()] = _11_v1;
      _nw0[(new BigNumber(21)).toNumber()] = _11_v1;
      _16_v6 = _nw0;
      let _17_v7;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_10_v0, ((_10_v0) ? (_16_v6) : (_16_v6)));
      _17_v7 = _nw1;
      let _18_i0;
      _18_i0 = _dafny.ZERO;
      L0: {
        while ((_17_v7).f28) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_18_i0)) {
              break L0;
            }
            _18_i0 = (_18_i0).plus(_dafny.ONE);
            if (_module.__default.fm0(globalState)) {
              let _19_v8;
              _19_v8 = new _dafny.CodePoint('d'.codePointAt(0));
              _19_v8 = ((_10_v0) ? (_19_v8) : (_19_v8));
              let _20_v9;
              let _nw2 = Array((new BigNumber(19)).toNumber());
              _nw2[(_dafny.ZERO).toNumber()] = _10_v0;
              _nw2[(_dafny.ONE).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(2)).toNumber()] = _10_v0;
              _nw2[(new BigNumber(3)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(4)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(5)).toNumber()] = (_17_v7).fm2((_17_v7).f28, globalState);
              _nw2[(new BigNumber(6)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(7)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(8)).toNumber()] = !((_17_v7).f28);
              _nw2[(new BigNumber(9)).toNumber()] = _10_v0;
              _nw2[(new BigNumber(10)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(11)).toNumber()] = _10_v0;
              _nw2[(new BigNumber(12)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(13)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(14)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(15)).toNumber()] = (_17_v7).f28;
              _nw2[(new BigNumber(16)).toNumber()] = !((_17_v7).f28);
              _nw2[(new BigNumber(17)).toNumber()] = _10_v0;
              _nw2[(new BigNumber(18)).toNumber()] = (_17_v7).f28;
              _20_v9 = _nw2;
              let _21_v10;
              _21_v10 = _20_v9;
              let _22_v11;
              _22_v11 = _dafny.Seq.of(_20_v9, _20_v9, _20_v9, _20_v9, _20_v9);
              let _23_v12;
              _23_v12 = _dafny.Seq.UnicodeFromString("btlenkt");
              let _24_v13;
              _24_v13 = _dafny.Set.fromElements(_20_v9, (_21_v10), _20_v9, (_22_v11)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_23_v12).length)), new BigNumber((_22_v11).length))]);
              let _25_v14;
              _25_v14 = _module.D0.create_DC1((_17_v7).f28, _11_v1);
              let _26_v15;
              _26_v15 = _module.D0.create_DC2(_25_v14);
              let _27_v16;
              _27_v16 = _dafny.MultiSet.fromElements(_10_v0, !((_17_v7).f28));
              let _28_v17;
              _28_v17 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_20_v9);
              let _29_v18;
              _29_v18 = _dafny.Seq.of(_dafny.Set.fromElements((((_28_v17).contains(new BigNumber((_dafny.Set.fromElements(_11_v1, _11_v1, new BigNumber((_15_v5).length), _11_v1)).length))) ? ((_28_v17).get(new BigNumber((_dafny.Set.fromElements(_11_v1, _11_v1, new BigNumber((_15_v5).length), _11_v1)).length))) : (_20_v9))), _24_v13, (_24_v13).Difference(_24_v13), _dafny.Set.fromElements(_20_v9), (_24_v13).Union(_24_v13));
              let _30_v19;
              _30_v19 = _module.D1.create_DC4((_17_v7).f28);
              let _rhs0 = (_29_v18)[_module.__default.safeIndex(_11_v1, new BigNumber((_29_v18).length))];
              let _rhs1 = _26_v15;
              let _rhs2 = _20_v9;
              let _rhs3 = ((_27_v16).Union(_27_v16)).Intersect(_module.__default.fm6((_30_v19).dtor_cf5, _11_v1, globalState));
              _24_v13 = _rhs0;
              _26_v15 = _rhs1;
              _20_v9 = _rhs2;
              _27_v16 = _rhs3;
              let _31_v20;
              let _nw3 = new _module.C0();
              _nw3.__ctor((_17_v7).f28, (_17_v7).f29);
              _31_v20 = _nw3;
              let _32_v21;
              _32_v21 = _dafny.Map.Empty.slice().updateUnsafe((_17_v7).f28,_13_v3);
              let _33_v22;
              _33_v22 = _dafny.Seq.of(_dafny.Seq.of((_17_v7).f28), _13_v3, _dafny.Seq.Concat((((_32_v21).contains(_10_v0)) ? ((_32_v21).get(_10_v0)) : (_13_v3)), _13_v3), _dafny.Seq.of((_31_v20).f28));
              let _34_v23;
              _34_v23 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_11_v1);
              let _35_v24;
              _35_v24 = _module.D3.create_DC10((_17_v7).fm2((_17_v7).f28, globalState), new BigNumber((_34_v23).length), !((_17_v7).f28), new BigNumber(282), (_17_v7).f29);
              let _index0 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_20_v9).length));
              (_20_v9)[_index0] = (_35_v24).dtor_cf14;
              let _index1 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_20_v9).length));
              let _rhs4 = false;
              let _rhs5 = ((_11_v1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_11_v1,new BigNumber(-924))).length))).isLessThan(_11_v1);
              let _rhs6 = _33_v22;
              let _rhs7 = !((_17_v7).f28);
              let _lhs0 = globalState;
              let _lhs1 = globalState;
              let _lhs2 = _20_v9;
              let _lhs3 = _module.__default.safeIndex(new BigNumber(970), new BigNumber((_20_v9).length));
              _lhs0.f0 = _rhs4;
              _lhs1.f0 = _rhs5;
              _33_v22 = _rhs6;
              _lhs2[_lhs3] = _rhs7;
              let _36_v25;
              _36_v25 = _module.D2.create_DC6(_27_v16);
              let _37_v26;
              _37_v26 = _dafny.Map.Empty.slice().updateUnsafe((((_31_v20).fm2((_17_v7).f28, globalState)) ? (_36_v25) : (_36_v25)),new BigNumber(841));
              _37_v26 = (_37_v26).update(_module.D2.create_DC6(_27_v16), _11_v1);
            } else {
              let _38_v27;
              _38_v27 = _dafny.MultiSet.fromElements((_11_v1).isLessThanOrEqualTo(_11_v1), (_11_v1).isEqualTo(_11_v1));
              let _rhs8 = _11_v1;
              let _rhs9 = (((_38_v27).contains((_17_v7).f28)) ? ((_38_v27).get((_17_v7).f28)) : (new BigNumber(591)));
              let _rhs10 = (_17_v7).f29;
              _11_v1 = _rhs8;
              _11_v1 = _rhs9;
              _16_v6 = _rhs10;
              let _39_v28;
              let _nw4 = Array((new BigNumber(25)).toNumber()).fill(false);
              _39_v28 = _nw4;
              let _index2 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_39_v28).length));
              (_39_v28)[_index2] = (_17_v7).f28;
              let _40_v29;
              _40_v29 = new _dafny.CodePoint('o'.codePointAt(0));
              let _41_v30;
              _41_v30 = _dafny.Seq.of(_40_v29, _40_v29, _40_v29);
              let _42_v31;
              _42_v31 = _module.D0.create_DC1(_10_v0, new BigNumber(-832));
              let _index3 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_39_v28).length));
              let _rhs11 = (new BigNumber((_41_v30).length)).plus((_42_v31).dtor_cf2);
              let _rhs12 = (_17_v7).f28;
              let _rhs13 = (((_17_v7).f28) ? ((function () {
                let _coll5 = new _dafny.Set();
                for (const _compr_5 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), ((_43_v1) => function (_44_i1) {
                  return _43_v1;
                })(_11_v1))).Elements) {
                  let _45_v32 = _compr_5;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(944)), ((_46_v1) => function (_44_i1) {
                    return _46_v1;
                  })(_11_v1)), _45_v32)) {
                    _coll5.add(_module.__default.safeModuloInt(_45_v32, new BigNumber((_dafny.Seq.UnicodeFromString("sbrnrtwv")).length)));
                  }
                }
                return _coll5;
              }()).contains(_11_v1)) : ((_11_v1).isLessThan(_11_v1)));
              let _lhs4 = _39_v28;
              let _lhs5 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_39_v28).length));
              _11_v1 = _rhs11;
              _lhs4[_lhs5] = _rhs12;
              _10_v0 = _rhs13;
              _40_v29 = _40_v29;
              let _47_v33;
              _47_v33 = _dafny.MultiSet.fromElements(_17_v7, _17_v7, _17_v7);
              _11_v1 = (((_47_v33).contains(_17_v7)) ? ((_47_v33).get(_17_v7)) : ((_11_v1).plus(_11_v1)));
              let _48_v34;
              _48_v34 = _dafny.Map.Empty.slice().updateUnsafe(_41_v30,_10_v0);
              let _49_v35;
              let _nw5 = Array((new BigNumber(7)).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = (_12_v2).update(_10_v0, false);
              _nw5[(_dafny.ONE).toNumber()] = _12_v2;
              _nw5[(new BigNumber(2)).toNumber()] = _12_v2;
              _nw5[(new BigNumber(3)).toNumber()] = (_12_v2).Merge(_module.__default.fm7((_17_v7).f28, (_17_v7).f28, (_39_v28)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_39_v28).length))], _11_v1, globalState));
              _nw5[(new BigNumber(4)).toNumber()] = _12_v2;
              _nw5[(new BigNumber(5)).toNumber()] = (_module.__default.fm7(_10_v0, _module.__default.fm0(globalState), (_17_v7).f28, new BigNumber((_48_v34).length), globalState)).update(_10_v0, _10_v0);
              _nw5[(new BigNumber(6)).toNumber()] = _12_v2;
              _49_v35 = _nw5;
              let _index4 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_49_v35).length));
              (_49_v35)[_index4] = (_dafny.Map.Empty.slice().updateUnsafe((_17_v7).f28,(_17_v7).fm2((_17_v7).f28, globalState))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_39_v28)[_module.__default.safeIndex(new BigNumber(941), new BigNumber((_39_v28).length))],(_17_v7).f28));
              let _index5 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_49_v35).length));
              (_49_v35)[_index5] = _12_v2;
            }
            if (!(_10_v0) || (!((_17_v7).f28) || (!((_17_v7).f28)))) {
              let _50_v36;
              _50_v36 = _dafny.Seq.UnicodeFromString("vithblii");
              (globalState).f0 = _dafny.Seq.contains(_dafny.Seq.Concat(_50_v36, _50_v36), new _dafny.CodePoint('i'.codePointAt(0)));
              let _51_v37;
              let _init0 = function (_52_i2) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              };
              let _nw6 = Array((new BigNumber(7)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw6.length); _i0_0++) {
                _nw6[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _51_v37 = _nw6;
              let _53_v38;
              _53_v38 = _50_v36;
              let _rhs14 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_54_i3) {
                return new _dafny.CodePoint('h'.codePointAt(0));
              }), _dafny.Seq.Concat((_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), function (_55_i4) {
                return new _dafny.CodePoint('x'.codePointAt(0));
              })), _50_v36));
              let _rhs15 = (((_17_v7).f28) ? (_51_v37) : ((((_17_v7).f28) ? (_51_v37) : (_51_v37))));
              let _lhs6 = globalState;
              _lhs6.f0 = _rhs14;
              _51_v37 = _rhs15;
              _10_v0 = _10_v0;
              _11_v1 = (_module.__default.safeDivisionInt(_11_v1, _11_v1)).minus(_module.__default.fm3(_11_v1, (_17_v7).f28, _11_v1, _dafny.Map.Empty.slice().updateUnsafe(_11_v1,(_17_v7).f28), globalState));
              let _56_v39;
              let _nw7 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
              _56_v39 = _nw7;
              let _57_v40;
              _57_v40 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_17_v7);
              let _58_v41;
              _58_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_57_v40).length),_module.__default.fm8(_11_v1, globalState));
              let _index6 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_56_v39).length));
              (_56_v39)[_index6] = (_dafny.Set.fromElements((_17_v7).f28, _10_v0, _10_v0, (_17_v7).f28, (_17_v7).f28)).Difference((((_58_v41).contains(new BigNumber(330))) ? ((_58_v41).get(new BigNumber(330))) : (_module.__default.fm8(_11_v1, globalState))));
              let _index7 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_56_v39).length));
              (_56_v39)[_index7] = (_15_v5).Difference((_15_v5).Union(_15_v5));
            } else {
              let _59_v42;
              _59_v42 = new _dafny.CodePoint('l'.codePointAt(0));
              let _60_v43;
              _60_v43 = _module.D3.create_DC10(_10_v0, _11_v1, (_17_v7).f28, _11_v1, (_17_v7).f29);
              let _61_v44;
              _61_v44 = _dafny.Map.Empty.slice().updateUnsafe((_60_v43).dtor_cf14,_17_v7);
              let _62_v45;
              _62_v45 = _dafny.Seq.of((((_61_v44).contains(!((_17_v7).fm2((_17_v7).f28, globalState)))) ? ((_61_v44).get(!((_17_v7).fm2((_17_v7).f28, globalState)))) : (_17_v7)), _17_v7, _17_v7);
              let _63_v46;
              _63_v46 = _dafny.Seq.UnicodeFromString("squrxyb");
              let _64_v47;
              _64_v47 = _63_v46;
              let _65_v48;
              _65_v48 = _dafny.Map.Empty.slice().updateUnsafe(_64_v47,false);
              let _66_v49;
              _66_v49 = _module.D6.create_DC13(_59_v42);
              let _pat_let_tv0 = _59_v42;
              let _67_v50;
              _67_v50 = _module.D6.create_DC13((function (_pat_let0_0) {
  return function (_68_dt__update__tmp_h1) {
    return function (_pat_let1_0) {
      return function (_69_dt__update_hcf19_h0) {
        return _module.D6.create_DC13(_69_dt__update_hcf19_h0);
      }(_pat_let1_0);
    }(_pat_let_tv0);
  }(_pat_let0_0);
}(_66_v49)).dtor_cf19);
              let _rhs16 = _11_v1;
              let _rhs17 = (_62_v45)[_module.__default.safeIndex(new BigNumber(((_65_v48).Merge(_65_v48)).length), new BigNumber((_62_v45).length))];
              let _rhs18 = (_66_v49).dtor_cf19;
              _11_v1 = _rhs16;
              _17_v7 = _rhs17;
              _59_v42 = _rhs18;
              let _70_v51;
              let _nw8 = Array((new BigNumber(5)).toNumber());
              _nw8[(_dafny.ZERO).toNumber()] = (_17_v7).f28;
              _nw8[(_dafny.ONE).toNumber()] = (_17_v7).f28;
              _nw8[(new BigNumber(2)).toNumber()] = _10_v0;
              _nw8[(new BigNumber(3)).toNumber()] = (_17_v7).f28;
              _nw8[(new BigNumber(4)).toNumber()] = _10_v0;
              _70_v51 = _nw8;
              let _71_v52;
              let _nw9 = Array((new BigNumber(18)).toNumber());
              _nw9[(_dafny.ZERO).toNumber()] = (_17_v7).f29;
              _nw9[(_dafny.ONE).toNumber()] = _16_v6;
              _nw9[(new BigNumber(2)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(3)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(4)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(5)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(6)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(7)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(8)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(9)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(10)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(11)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(12)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(13)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(14)).toNumber()] = _16_v6;
              _nw9[(new BigNumber(15)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(16)).toNumber()] = (_17_v7).f29;
              _nw9[(new BigNumber(17)).toNumber()] = _16_v6;
              _71_v52 = _nw9;
              let _72_v53;
              _72_v53 = _dafny.Map.Empty.slice().updateUnsafe(_70_v51,_71_v52);
              _72_v53 = (_72_v53).update(_70_v51, _71_v52);
              let _73_v54;
              let _nw10 = new _module.C0();
              _nw10.__ctor(_10_v0, (_17_v7).f29);
              _73_v54 = _nw10;
              let _74_v55;
              _74_v55 = _dafny.MultiSet.fromElements(_11_v1);
              let _75_v56;
              _75_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-735)), ((_76_v42) => function (_77_i5) {
                return _76_v42;
              })(_59_v42)),new BigNumber((_74_v55).cardinality()));
              _75_v56 = (_75_v56).update(_module.__default.fm9(globalState), _11_v1);
              let _78_v57;
              _78_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(275),new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("v"), _63_v46)).length));
              let _79_v58;
              _79_v58 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,(_17_v7).f28);
              let _80_v59;
              _80_v59 = _dafny.Seq.of(_79_v58);
              _78_v57 = (_78_v57).update((((_13_v3)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_13_v3).length))]) ? (_11_v1) : (_module.__default.fm3(_11_v1, (_17_v7).f28, _11_v1, (_80_v59)[_module.__default.safeIndex(new BigNumber((_63_v46).length), new BigNumber((_80_v59).length))], globalState))), _11_v1);
            }
            _11_v1 = (((_17_v7).fm2(_10_v0, globalState)) ? (_11_v1) : (_11_v1));
            (globalState).f0 = (_17_v7).f28;
          }
        }
      }
      let _81_v60;
      _81_v60 = _dafny.MultiSet.fromElements(_10_v0, (_17_v7).f28, _10_v0, (_17_v7).f28);
      let _82_v61;
      let _nw11 = Array((new BigNumber(8)).toNumber()).fill(false);
      _82_v61 = _nw11;
      let _83_v62;
      _83_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(618),_82_v61);
      let _84_v63;
      _84_v63 = _dafny.Seq.of(_83_v62, _83_v62, _83_v62);
      let _rhs19 = (_81_v60).IsDisjointFrom(_81_v60);
      let _rhs20 = true;
      let _rhs21 = (_17_v7).f28;
      let _rhs22 = (new BigNumber((((_84_v63)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_11_v1, _11_v1)).length), new BigNumber((_84_v63).length))]).Merge(_83_v62)).length)).minus(_11_v1);
      let _lhs7 = globalState;
      let _lhs8 = globalState;
      _lhs7.f0 = _rhs19;
      _10_v0 = _rhs20;
      _lhs8.f0 = _rhs21;
      _11_v1 = _rhs22;
      if (!((_17_v7).f28)) {
        _11_v1 = _11_v1;
        (globalState).f0 = (_17_v7).f28;
        let _85_v64;
        _85_v64 = _dafny.Seq.of(_dafny.Seq.of(_10_v0));
        let _86_v65;
        _86_v65 = _module.D0.create_DC1(_10_v0, new BigNumber((_85_v64).length));
        if ((_86_v65).dtor_cf1) {
          let _index8 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_82_v61).length));
          (_82_v61)[_index8] = (_17_v7).f28;
          let _index9 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_82_v61).length));
          (_82_v61)[_index9] = ((_17_v7).f28) === ((_17_v7).f28);
          let _87_v66;
          let _nw12 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _87_v66 = _nw12;
          let _88_v67;
          _88_v67 = _dafny.Set.fromElements(_87_v66);
          (globalState).f0 = (_dafny.Set.fromElements(_87_v66)).IsSubsetOf(_88_v67);
          let _89_v68;
          _89_v68 = _module.D6.create_DC14(_module.__default.fm10(_11_v1, new BigNumber(673), _11_v1, globalState));
          let _90_v69;
          _90_v69 = _module.D1.create_DC3((_17_v7).f29);
          let _91_v70;
          _91_v70 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_89_v68).dtor_cf20).length),_90_v69);
          let _92_v71;
          _92_v71 = _module.D1.create_DC5((((_91_v70).contains(_11_v1)) ? ((_91_v70).get(_11_v1)) : (_90_v69)));
          let _93_v72;
          _93_v72 = _module.D6.create_DC16(_dafny.Seq.of(_92_v71), _dafny.Seq.UnicodeFromString("am"));
          let _94_v73;
          _94_v73 = _dafny.Seq.UnicodeFromString("jrtiwitvl");
          let _pat_let_tv1 = _90_v69;
          let _95_v74;
          _95_v74 = _dafny.Seq.of(function (_pat_let2_0) {
            return function (_96_dt__update__tmp_h2) {
              return function (_pat_let3_0) {
                return function (_97_dt__update_hcf6_h0) {
                  return _module.D1.create_DC5(_97_dt__update_hcf6_h0);
                }(_pat_let3_0);
              }(_pat_let_tv1);
            }(_pat_let2_0);
          }(_92_v71), _92_v71);
          let _98_v75;
          _98_v75 = new _dafny.CodePoint('i'.codePointAt(0));
          let _pat_let_tv2 = _94_v73;
          let _pat_let_tv3 = _95_v74;
          let _pat_let_tv4 = _94_v73;
          let _pat_let_tv5 = _12_v2;
          let _pat_let_tv6 = _94_v73;
          let _pat_let_tv7 = _98_v75;
          let _99_v76;
          let _nw13 = Array((new BigNumber(10)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _93_v72;
          _nw13[(_dafny.ONE).toNumber()] = function (_pat_let4_0) {
            return function (_102_dt__update__tmp_h4) {
              return function (_pat_let7_0) {
                return function (_103_dt__update_hcf24_h0) {
                  return _module.D6.create_DC16(_103_dt__update_hcf24_h0, (_102_dt__update__tmp_h4).dtor_cf25);
                }(_pat_let7_0);
              }(_pat_let_tv3);
            }(_pat_let4_0);
          }(function (_pat_let5_0) {
            return function (_100_dt__update__tmp_h3) {
              return function (_pat_let6_0) {
                return function (_101_dt__update_hcf25_h0) {
                  return _module.D6.create_DC16((_100_dt__update__tmp_h3).dtor_cf24, _101_dt__update_hcf25_h0);
                }(_pat_let6_0);
              }(_pat_let_tv2);
            }(_pat_let5_0);
          }(_93_v72));
          _nw13[(new BigNumber(2)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(3)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(4)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(5)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(6)).toNumber()] = function (_pat_let8_0) {
            return function (_104_dt__update__tmp_h5) {
              return function (_pat_let9_0) {
                return function (_105_dt__update_hcf25_h1) {
                  return _module.D6.create_DC16((_104_dt__update__tmp_h5).dtor_cf24, _105_dt__update_hcf25_h1);
                }(_pat_let9_0);
              }(_dafny.Seq.update(_pat_let_tv4, _module.__default.safeIndex(new BigNumber((_pat_let_tv5).length), new BigNumber((_pat_let_tv6).length)), _pat_let_tv7));
            }(_pat_let8_0);
          }(_93_v72);
          _nw13[(new BigNumber(7)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(8)).toNumber()] = _93_v72;
          _nw13[(new BigNumber(9)).toNumber()] = _93_v72;
          _99_v76 = _nw13;
          let _index10 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_99_v76).length));
          (_99_v76)[_index10] = _module.D6.create_DC16(_95_v74, _94_v73);
          let _index11 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_99_v76).length));
          (_99_v76)[_index11] = (((_17_v7).f28) ? (_93_v72) : (_93_v72));
          (globalState).f0 = (_17_v7).f28;
          _93_v72 = _93_v72;
        } else {
          let _106_v77;
          _106_v77 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("xsxaj"));
          (globalState).f0 = (_106_v77).IsSubsetOf(_106_v77);
          let _107_v78;
          _107_v78 = _dafny.MultiSet.fromElements(new BigNumber(777));
          let _108_v79;
          _108_v79 = _dafny.Map.Empty.slice().updateUnsafe((_11_v1).multipliedBy(new BigNumber((_14_v4).length)),_107_v78);
          _107_v78 = (((_108_v79).contains(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ysstjicth"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-758)), function (_110_i6) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }))).length))) ? ((_108_v79).get(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ysstjicth"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-758)), function (_109_i6) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }))).length))) : (_107_v78));
          let _111_v80;
          _111_v80 = _dafny.Seq.UnicodeFromString("ipw");
          let _112_v81;
          _112_v81 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_111_v80);
          let _rhs23 = _module.__default.fm10((new BigNumber(-680)).plus(_11_v1), _11_v1, _11_v1, globalState);
          let _rhs24 = _dafny.Seq.Concat(_dafny.Seq.Concat(_111_v80, (((_112_v81).contains(_11_v1)) ? ((_112_v81).get(_11_v1)) : (_111_v80))), _111_v80);
          let _rhs25 = _module.__default.safeModuloInt(new BigNumber(491), _11_v1);
          let _rhs26 = (_17_v7).f28;
          let _lhs9 = globalState;
          let _lhs10 = globalState;
          _13_v3 = _rhs23;
          _lhs9.f2 = _rhs24;
          _11_v1 = _rhs25;
          _lhs10.f0 = _rhs26;
          let _index12 = _module.__default.safeIndex(new BigNumber(738), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index12] = _11_v1;
          let _index13 = _module.__default.safeIndex(new BigNumber(738), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index13] = _11_v1;
          let _113_v82;
          _113_v82 = _dafny.Map.Empty.slice().updateUnsafe(_17_v7,(_17_v7).f29);
          _113_v82 = (_113_v82).update(_17_v7, (_17_v7).f29);
        }
        let _114_v83;
        _114_v83 = new _dafny.CodePoint('p'.codePointAt(0));
        _114_v83 = _114_v83;
        let _115_v84;
        let _nw14 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Set.Empty);
        _115_v84 = _nw14;
        let _index14 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_115_v84).length));
        (_115_v84)[_index14] = function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-913), new BigNumber(328))) {
            let _116_v85 = _compr_6;
            if (((new BigNumber(-913)).isLessThanOrEqualTo(_116_v85)) && ((_116_v85).isLessThan(new BigNumber(328)))) {
              _coll6.add((_116_v85).minus(new BigNumber((_dafny.MultiSet.fromElements(_11_v1, _11_v1)).cardinality())));
            }
          }
          return _coll6;
        }();
        let _index15 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_115_v84).length));
        (_115_v84)[_index15] = _14_v4;
      } else {
        let _117_v86;
        _117_v86 = _dafny.Map.Empty.slice().updateUnsafe(_10_v0,_11_v1);
        _11_v1 = ((((_117_v86).contains(!(!((_17_v7).f28)))) ? ((_117_v86).get(!(!((_17_v7).f28)))) : (_11_v1))).plus(_11_v1);
        (globalState).f0 = true;
        let _index16 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length));
        (_82_v61)[_index16] = !((_17_v7).f28) || ((_17_v7).f28);
        let _118_v87;
        _118_v87 = _dafny.Seq.UnicodeFromString("dicp");
        let _119_v88;
        _119_v88 = _dafny.Seq.of(_11_v1);
        let _index17 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length));
        let _rhs27 = _13_v3;
        let _rhs28 = ((_10_v0) ? (new BigNumber((_dafny.Seq.Concat(_118_v87, _dafny.Seq.UnicodeFromString("dqnhuuat"))).length)) : ((_dafny.ZERO).minus((_119_v88)[_module.__default.safeIndex(_11_v1, new BigNumber((_119_v88).length))])));
        let _rhs29 = (_17_v7).f28;
        let _lhs11 = _82_v61;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length));
        _13_v3 = _rhs27;
        _11_v1 = _rhs28;
        _lhs11[_lhs12] = _rhs29;
        let _120_v89;
        let _nw15 = Array((new BigNumber(13)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = (_82_v61)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length))];
        _nw15[(_dafny.ONE).toNumber()] = (_81_v60).IsDisjointFrom(_dafny.MultiSet.fromElements(_10_v0, _10_v0, (_17_v7).f28, _10_v0));
        _nw15[(new BigNumber(2)).toNumber()] = _10_v0;
        _nw15[(new BigNumber(3)).toNumber()] = false;
        _nw15[(new BigNumber(4)).toNumber()] = (_17_v7).f28;
        _nw15[(new BigNumber(5)).toNumber()] = (_82_v61)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length))];
        _nw15[(new BigNumber(6)).toNumber()] = true;
        _nw15[(new BigNumber(7)).toNumber()] = (_17_v7).f28;
        _nw15[(new BigNumber(8)).toNumber()] = (_82_v61)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length))];
        _nw15[(new BigNumber(9)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(672))).equals(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), ((_121_v88) => function (_122_i7) {
          return new BigNumber((_121_v88).length);
        })(_119_v88))));
        _nw15[(new BigNumber(10)).toNumber()] = (_82_v61)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length))];
        _nw15[(new BigNumber(11)).toNumber()] = (_17_v7).f28;
        _nw15[(new BigNumber(12)).toNumber()] = (_17_v7).f28;
        _120_v89 = _nw15;
        _120_v89 = _120_v89;
        let _123_v90;
        _123_v90 = _dafny.Seq.of(_17_v7, _17_v7, _17_v7);
        let _124_v91;
        _124_v91 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,(_82_v61)[_module.__default.safeIndex(new BigNumber(488), new BigNumber((_82_v61).length))]);
        _17_v7 = (_123_v90)[_module.__default.safeIndex(_module.__default.fm3(new BigNumber(502), false, _11_v1, _124_v91, globalState), new BigNumber((_123_v90).length))];
      }
      let _125_v92;
      _125_v92 = _module.D0.create_DC1(((!((_17_v7).f28)) ? ((_17_v7).f28) : (_10_v0)), _11_v1);
      let _source0 = _125_v92;
      if (_source0.is_DC1) {
        let _126___mcc_h0 = (_source0).cf1;
        let _127___mcc_h1 = (_source0).cf2;
        let _128_cf2 = _127___mcc_h1;
        let _129_cf1 = _126___mcc_h0;
        _11_v1 = _11_v1;
        let _130_v93;
        _130_v93 = _dafny.Map.Empty.slice().updateUnsafe(_128_cf2,!(_129_cf1));
        _128_cf2 = (_dafny.ZERO).minus(_module.__default.fm3((_128_cf2).minus(_128_cf2), (_17_v7).f28, _11_v1, _130_v93, globalState));
        let _pat_let_tv8 = _11_v1;
        let _131_v94;
        _131_v94 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let10_0) {
          return function (_132_dt__update__tmp_h6) {
            return function (_pat_let11_0) {
              return function (_133_dt__update_hcf2_h0) {
                return _module.D0.create_DC1((_132_dt__update__tmp_h6).dtor_cf1, _133_dt__update_hcf2_h0);
              }(_pat_let11_0);
            }(_pat_let_tv8);
          }(_pat_let10_0);
        }(_125_v92),_130_v93);
        let _134_v95;
        _134_v95 = _dafny.Seq.of(_14_v4, _14_v4);
        let _135_v96;
        _135_v96 = _dafny.Seq.of(_11_v1, _128_cf2, _11_v1, _11_v1);
        let _rhs30 = (_131_v94).Merge((_131_v94).Merge(_131_v94));
        let _rhs31 = !(((_134_v95)[_module.__default.safeIndex(_128_cf2, new BigNumber((_134_v95).length))]).equals(_14_v4)) || (!((_17_v7).f28));
        let _rhs32 = _module.__default.safeDivisionInt(_128_cf2, (_dafny.ZERO).minus(new BigNumber((_135_v96).length)));
        let _rhs33 = _128_cf2;
        let _lhs13 = globalState;
        _131_v94 = _rhs30;
        _lhs13.f0 = _rhs31;
        _128_cf2 = _rhs32;
        _11_v1 = _rhs33;
        if (!(_11_v1).isEqualTo(_11_v1)) {
          let _136_v97;
          let _nw16 = new _module.C0();
          _nw16.__ctor((_17_v7).f28, (_17_v7).f29);
          _136_v97 = _nw16;
          let _index18 = _module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index18] = (new BigNumber(507)).multipliedBy(_128_cf2);
          let _index19 = _module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index19] = new BigNumber(((_12_v2).Merge(_12_v2)).length);
          let _137_v98;
          _137_v98 = _dafny.Seq.of(_82_v61);
          let _138_v99;
          _138_v99 = _dafny.MultiSet.fromElements(_11_v1);
          let _139_v100;
          _139_v100 = _dafny.Seq.UnicodeFromString("ckgrfoe");
          let _140_v101;
          _140_v101 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_139_v100);
          let _141_v102;
          _141_v102 = _module.D0.create_DC0(_138_v99);
          let _142_v103;
          _142_v103 = _dafny.Map.Empty.slice().updateUnsafe(((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length))],_138_v99);
          let _143_v104;
          let _nw17 = Array((new BigNumber(18)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = (_138_v99).Union(_138_v99);
          _nw17[(_dafny.ONE).toNumber()] = ((_138_v99).update(new BigNumber(187), _module.__default.abs(((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length))]))).Difference(_138_v99);
          _nw17[(new BigNumber(2)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(3)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_139_v100).length), _128_cf2, _11_v1, new BigNumber((_139_v100).length), _11_v1)).Intersect(_138_v99);
          _nw17[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_140_v101).length)))));
          _nw17[(new BigNumber(5)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(6)).toNumber()] = (_dafny.MultiSet.FromArray(_135_v96)).update(_11_v1, _module.__default.abs(_11_v1));
          _nw17[(new BigNumber(7)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(8)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.FromArray(_135_v96);
          _nw17[(new BigNumber(10)).toNumber()] = (_138_v99).Intersect(_138_v99);
          _nw17[(new BigNumber(11)).toNumber()] = (_141_v102).dtor_cf0;
          _nw17[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(_11_v1));
          _nw17[(new BigNumber(13)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(14)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(15)).toNumber()] = _138_v99;
          _nw17[(new BigNumber(16)).toNumber()] = ((((_142_v103).contains(new BigNumber(705))) ? ((_142_v103).get(new BigNumber(705))) : (_138_v99))).update(_128_cf2, _module.__default.abs(_11_v1));
          _nw17[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements(_128_cf2, new BigNumber(-663));
          _143_v104 = _nw17;
          let _index20 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_143_v104).length));
          (_143_v104)[_index20] = _dafny.MultiSet.FromArray(_135_v96);
          let _144_v105;
          _144_v105 = new _dafny.CodePoint('b'.codePointAt(0));
          let _145_v106;
          _145_v106 = _dafny.Map.Empty.slice().updateUnsafe(true,_144_v105);
          let _146_v107;
          _146_v107 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_129_cf1,_144_v105), _dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('k'.codePointAt(0))), (_145_v106).update((_17_v7).f28, _144_v105), _145_v106, (((_136_v97).f28) ? (_145_v106) : (_145_v106)));
          let _147_v108;
          _147_v108 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_139_v100).length),_136_v97);
          let _148_v109;
          _148_v109 = _module.D7.create_DC17((((_147_v108).contains(new BigNumber((_13_v3).length))) ? ((_147_v108).get(new BigNumber((_13_v3).length))) : (_17_v7)));
          let _149_v110;
          _149_v110 = _dafny.Seq.of((_148_v109).dtor_cf26);
          let _150_v111;
          _150_v111 = _dafny.Seq.of(_139_v100, _139_v100);
          let _151_v112;
          _151_v112 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("f"), _139_v100, _139_v100, _dafny.Seq.UnicodeFromString("a"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(55)), ((_152_v105) => function (_153_i8) {
            return _152_v105;
          })(_144_v105)));
          let _154_v113;
          _154_v113 = _module.D7.create_DC18((_17_v7).f28);
          let _index21 = _module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length));
          let _index22 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_143_v104).length));
          let _rhs34 = _137_v98;
          let _rhs35 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_149_v110, _dafny.Seq.of(_136_v97, _136_v97, _136_v97)), _module.__default.safeIndex(new BigNumber(381), new BigNumber((_dafny.Seq.Concat(_149_v110, _dafny.Seq.of(_136_v97, _136_v97, _136_v97))).length)), _136_v97)).length);
          let _rhs36 = _138_v99;
          let _rhs37 = ((((_136_v97).f28) ? (_module.__default.fm11(new BigNumber(-642), new _dafny.CodePoint('a'.codePointAt(0)), globalState)) : (_151_v112))).IsProperSubsetOf((_dafny.MultiSet.FromArray(_150_v111)).Difference(_151_v112));
          let _rhs38 = _module.__default.fm12(_11_v1, ((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length))], (((_81_v60).contains((_136_v97).f28)) ? ((_81_v60).get((_136_v97).f28)) : (((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length))])), _154_v113, globalState);
          let _lhs14 = (_17_v7).f29;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length));
          let _lhs16 = _143_v104;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_143_v104).length));
          let _lhs18 = globalState;
          _137_v98 = _rhs34;
          _lhs14[_lhs15] = _rhs35;
          _lhs16[_lhs17] = _rhs36;
          _lhs18.f0 = _rhs37;
          _146_v107 = _rhs38;
          let _index23 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_82_v61).length));
          (_82_v61)[_index23] = (_17_v7).f28;
          let _index24 = _module.__default.safeIndex(new BigNumber(991), new BigNumber((_82_v61).length));
          (_82_v61)[_index24] = (_17_v7).f28;
          let _index25 = _module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index25] = (_128_cf2).plus((_11_v1).multipliedBy(((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(804), new BigNumber(((_17_v7).f29).length))]));
        } else {
          _130_v93 = (_130_v93).update(_11_v1, (_17_v7).f28);
          _17_v7 = _17_v7;
          let _155_v114;
          _155_v114 = _82_v61;
          let _156_v115;
          let _nw18 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _156_v115 = _nw18;
          let _index26 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_156_v115).length));
          (_156_v115)[_index26] = new _dafny.CodePoint('k'.codePointAt(0));
          let _157_v116;
          _157_v116 = new _dafny.CodePoint('p'.codePointAt(0));
          let _index27 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_156_v115).length));
          let _rhs39 = _82_v61;
          let _rhs40 = _157_v116;
          let _lhs19 = _156_v115;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(346), new BigNumber((_156_v115).length));
          _155_v114 = _rhs39;
          _lhs19[_lhs20] = _rhs40;
          let _index28 = _module.__default.safeIndex(new BigNumber(767), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index28] = new BigNumber(21);
          let _158_v117;
          _158_v117 = _dafny.MultiSet.fromElements(new BigNumber(-324));
          let _index29 = _module.__default.safeIndex(new BigNumber(767), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index29] = (_dafny.ZERO).minus((((_158_v117).contains(_11_v1)) ? ((_158_v117).get(_11_v1)) : (_11_v1)));
          let _159_v118;
          _159_v118 = _dafny.Seq.of(_82_v61, (((_17_v7).f28) ? (_82_v61) : (_82_v61)), _82_v61);
          let _index30 = _module.__default.safeIndex(new BigNumber(767), new BigNumber(((_17_v7).f29).length));
          let _rhs41 = _11_v1;
          let _rhs42 = (_159_v118)[_module.__default.safeIndex(((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(767), new BigNumber(((_17_v7).f29).length))], new BigNumber((_159_v118).length))];
          let _rhs43 = _11_v1;
          let _lhs21 = (_17_v7).f29;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(767), new BigNumber(((_17_v7).f29).length));
          _lhs21[_lhs22] = _rhs41;
          _82_v61 = _rhs42;
          _128_cf2 = _rhs43;
        }
      } else if (_source0.is_DC0) {
        let _160___mcc_h2 = (_source0).cf0;
        let _161_cf0 = _160___mcc_h2;
        _11_v1 = _module.__default.fm3((_11_v1).plus(_11_v1), (_17_v7).f28, _11_v1, _dafny.Map.Empty.slice().updateUnsafe(_11_v1,(_17_v7).f28), globalState);
        let _162_v119;
        _162_v119 = _module.D6.create_DC15((_17_v7).f28, (_17_v7).f28, _11_v1);
        _11_v1 = (_162_v119).dtor_cf23;
        _11_v1 = _module.__default.safeDivisionInt(_11_v1, _11_v1);
        let _163_v120;
        _163_v120 = _dafny.Seq.UnicodeFromString("sudyiqf");
        (globalState).f2 = _dafny.Seq.Concat(_163_v120, _dafny.Seq.Concat(_163_v120, _163_v120));
      } else {
        let _164___mcc_h3 = (_source0).cf3;
        let _165_cf3 = _164___mcc_h3;
        let _166_v121;
        _166_v121 = _module.D1.create_DC5(_module.D1.create_DC3((_17_v7).f29));
        let _167_v122;
        _167_v122 = _module.D1.create_DC5(_166_v121);
        let _168_v123;
        _168_v123 = _module.D1.create_DC5(_167_v122);
        let _169_v124;
        _169_v124 = _dafny.Seq.of(_168_v123);
        let _170_v125;
        _170_v125 = _dafny.Seq.UnicodeFromString("yxedcgbw");
        let _171_v126;
        _171_v126 = _module.D6.create_DC16(_169_v124, _170_v125);
        let _source1 = _171_v126;
        if (_source1.is_DC14) {
          let _172___mcc_h4 = (_source1).cf20;
          let _173_cf20 = _172___mcc_h4;
          let _174_v127;
          _174_v127 = _dafny.Map.Empty.slice().updateUnsafe((_17_v7).f28,new BigNumber((_dafny.Seq.UnicodeFromString("kmsp")).length));
          _11_v1 = new BigNumber((((_dafny.Map.Empty.slice().updateUnsafe((_17_v7).f28,_11_v1)).Merge(_174_v127)).Merge(_174_v127)).length);
          let _index31 = _module.__default.safeIndex(new BigNumber(878), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index31] = _11_v1;
          let _index32 = _module.__default.safeIndex(new BigNumber(878), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index32] = _11_v1;
          _11_v1 = (((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(878), new BigNumber(((_17_v7).f29).length))]).minus(((_17_v7).f29)[_module.__default.safeIndex(new BigNumber(878), new BigNumber(((_17_v7).f29).length))]);
          let _175_v128;
          _175_v128 = _dafny.Seq.of(_170_v125);
          (globalState).f2 = (_175_v128)[_module.__default.safeIndex(_11_v1, new BigNumber((_175_v128).length))];
        } else if (_source1.is_DC15) {
          let _176___mcc_h5 = (_source1).cf21;
          let _177___mcc_h6 = (_source1).cf22;
          let _178___mcc_h7 = (_source1).cf23;
          let _179_cf23 = _178___mcc_h7;
          let _180_cf22 = _177___mcc_h6;
          let _181_cf21 = _176___mcc_h5;
          let _index33 = _module.__default.safeIndex(new BigNumber(646), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index33] = new BigNumber(-981);
          let _index34 = _module.__default.safeIndex(new BigNumber(646), new BigNumber(((_17_v7).f29).length));
          ((_17_v7).f29)[_index34] = _11_v1;
          (globalState).f2 = _170_v125;
          (globalState).f0 = true;
          (globalState).f0 = !((_17_v7).f28);
        } else if (_source1.is_DC16) {
          let _182___mcc_h8 = (_source1).cf24;
          let _183___mcc_h9 = (_source1).cf25;
          let _184_cf25 = _183___mcc_h9;
          let _185_cf24 = _182___mcc_h8;
          let _186_v129;
          _186_v129 = _dafny.Set.fromElements(_17_v7);
          let _187_v130;
          _187_v130 = _dafny.Map.Empty.slice().updateUnsafe(_186_v129,(_17_v7).f28);
          let _188_v131;
          let _nw19 = Array((new BigNumber(13)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_186_v129,(_17_v7).f28);
          _nw19[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_17_v7),!((_17_v7).f28))).Merge(_187_v130);
          _nw19[(new BigNumber(2)).toNumber()] = (_187_v130).Merge(_187_v130);
          _nw19[(new BigNumber(3)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(4)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(5)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(6)).toNumber()] = (_187_v130).Merge(_187_v130);
          _nw19[(new BigNumber(7)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(8)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_186_v129,(_17_v7).f28)).update(_dafny.Set.fromElements(_17_v7), (_17_v7).f28);
          _nw19[(new BigNumber(9)).toNumber()] = (_187_v130).update(_186_v129, (_17_v7).f28);
          _nw19[(new BigNumber(10)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(11)).toNumber()] = _187_v130;
          _nw19[(new BigNumber(12)).toNumber()] = _187_v130;
          _188_v131 = _nw19;
          let _index35 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_188_v131).length));
          (_188_v131)[_index35] = (_187_v130).Merge(_187_v130);
          let _189_v132;
          _189_v132 = _module.D7.create_DC18((_17_v7).f28);
          let _190_v133;
          _190_v133 = _module.D7.create_DC19(_189_v132);
          let _191_v134;
          _191_v134 = _module.D7.create_DC19(_189_v132);
          let _index36 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_188_v131).length));
          let _rhs44 = _187_v130;
          let _rhs45 = ((_10_v0) ? (_17_v7) : (_17_v7));
          let _rhs46 = _module.__default.fm9(globalState);
          let _rhs47 = _191_v134;
          let _lhs23 = _188_v131;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_188_v131).length));
          let _lhs25 = globalState;
          _lhs23[_lhs24] = _rhs44;
          _17_v7 = _rhs45;
          _lhs25.f2 = _rhs46;
          _191_v134 = _rhs47;
          _14_v4 = _14_v4;
          _11_v1 = _module.__default.safeDivisionInt(new BigNumber((_184_cf25).length), _11_v1);
          let _192_v135;
          _192_v135 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_11_v1, _11_v1)),(_17_v7).f28);
          let _193_v136;
          _193_v136 = _dafny.Seq.of(_11_v1);
          _192_v135 = (_192_v135).update(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(294)), _193_v136, _193_v136)).length), (_17_v7).f28);
        } else {
          let _194___mcc_h10 = (_source1).cf19;
          let _195_cf19 = _194___mcc_h10;
          (globalState).f2 = _dafny.Seq.Concat(_170_v125, (_171_v126).dtor_cf25);
          let _196_v137;
          let _nw20 = new _module.C0();
          _nw20.__ctor((_17_v7).f28, _16_v6);
          _196_v137 = _nw20;
          let _197_v138;
          _197_v138 = _module.D1.create_DC4((_11_v1).isLessThanOrEqualTo(_11_v1));
          let _rhs48 = _module.__default.safeModuloInt(_11_v1, _11_v1);
          let _rhs49 = _197_v138;
          let _rhs50 = (_17_v7).fm2((_17_v7).f28, globalState);
          let _rhs51 = (!((_17_v7).f28)) || ((_196_v137).f28);
          let _rhs52 = (_196_v137).f29;
          let _lhs26 = globalState;
          let _lhs27 = globalState;
          _11_v1 = _rhs48;
          _197_v138 = _rhs49;
          _lhs26.f0 = _rhs50;
          _lhs27.f0 = _rhs51;
          _16_v6 = _rhs52;
          _11_v1 = (_11_v1).minus((_dafny.ZERO).minus(_11_v1));
        }
        let _198_v139;
        _198_v139 = new _dafny.CodePoint('g'.codePointAt(0));
        let _199_v140;
        let _nw21 = Array((new BigNumber(2)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = _198_v139;
        _nw21[(_dafny.ONE).toNumber()] = _198_v139;
        _199_v140 = _nw21;
        let _200_v141;
        _200_v141 = _dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140);
        let _201_v142;
        _201_v142 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140));
        let _202_v143;
        _202_v143 = _dafny.Seq.of(_17_v7);
        let _203_v144;
        _203_v144 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(597)), function (_204_i9) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length), new BigNumber((_202_v143).length));
        let _205_v145;
        _205_v145 = _module.D8.create_DC20(_15_v5);
        let _206_v146;
        let _nw22 = Array((new BigNumber(24)).toNumber());
        _nw22[(_dafny.ZERO).toNumber()] = _200_v141;
        _nw22[(_dafny.ONE).toNumber()] = _200_v141;
        _nw22[(new BigNumber(2)).toNumber()] = (_201_v142)[_module.__default.safeIndex((((_203_v144).contains(_11_v1)) ? ((_203_v144).get(_11_v1)) : (new BigNumber(((_205_v145).dtor_cf29).length))), new BigNumber((_201_v142).length))];
        _nw22[(new BigNumber(3)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(4)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140);
        _nw22[(new BigNumber(6)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(7)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(8)).toNumber()] = (_200_v141).Merge(_200_v141);
        _nw22[(new BigNumber(9)).toNumber()] = (_201_v142)[_module.__default.safeIndex(new BigNumber(-425), new BigNumber((_201_v142).length))];
        _nw22[(new BigNumber(10)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(11)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(12)).toNumber()] = (_201_v142)[_module.__default.safeIndex(new BigNumber(267), new BigNumber((_201_v142).length))];
        _nw22[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140)).Merge(_200_v141);
        _nw22[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140);
        _nw22[(new BigNumber(15)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(16)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(17)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(18)).toNumber()] = (_200_v141).update(_198_v139, _199_v140);
        _nw22[(new BigNumber(19)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(20)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(21)).toNumber()] = _200_v141;
        _nw22[(new BigNumber(22)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140)).Merge((_200_v141).update(_198_v139, _199_v140));
        _nw22[(new BigNumber(23)).toNumber()] = ((_200_v141).update(_198_v139, _199_v140)).Merge(_200_v141);
        _206_v146 = _nw22;
        let _index37 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_206_v146).length));
        (_206_v146)[_index37] = _200_v141;
        let _index38 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_206_v146).length));
        (_206_v146)[_index38] = (_200_v141).Merge(_dafny.Map.Empty.slice().updateUnsafe(_198_v139,_199_v140));
        let _207_v147;
        _207_v147 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,(_17_v7).f28);
        _11_v1 = _module.__default.safeModuloInt(_11_v1, (_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_207_v147).length)).minus(_11_v1))));
        let _index39 = _module.__default.safeIndex(new BigNumber(649), new BigNumber(((_17_v7).f29).length));
        ((_17_v7).f29)[_index39] = new BigNumber(346);
        let _index40 = _module.__default.safeIndex(new BigNumber(649), new BigNumber(((_17_v7).f29).length));
        ((_17_v7).f29)[_index40] = (_11_v1).minus((_125_v92).dtor_cf2);
      }
      let _208_v148;
      _208_v148 = _dafny.Map.Empty.slice().updateUnsafe(_11_v1,_11_v1);
      let _209_v149;
      _209_v149 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_208_v148, _208_v148),(_17_v7).f28);
      let _210_v150;
      _210_v150 = new _dafny.CodePoint('q'.codePointAt(0));
      let _211_v151;
      _211_v151 = _dafny.Map.Empty.slice().updateUnsafe((_17_v7).f28,_210_v150);
      _209_v149 = (_209_v149).update(_dafny.Seq.of((_208_v148).update(_11_v1, new BigNumber((_211_v151).length))), (_17_v7).f28);
      return;
    }
    static Main(__noArgsParameter) {
      let _212_v0;
      _212_v0 = false;
      let _213_v1;
      _213_v1 = _dafny.Map.Empty.slice().updateUnsafe(_212_v0,_212_v0);
      let _214_v2;
      _214_v2 = _dafny.Set.fromElements(_212_v0);
      let _215_v3;
      _215_v3 = _dafny.Seq.UnicodeFromString("sk");
      let _216_v4;
      _216_v4 = _dafny.Seq.of(_212_v0, _212_v0, _212_v0, _212_v0);
      let _217_v5;
      _217_v5 = new BigNumber(228);
      let _218_v6;
      _218_v6 = _dafny.Seq.of(_217_v5);
      let _219_v7;
      _219_v7 = _dafny.Set.fromElements(new BigNumber((_218_v6).length));
      let _220_v8;
      _220_v8 = _dafny.Seq.of((_dafny.ZERO).minus(_217_v5), new BigNumber((_219_v7).length));
      let _221_v9;
      _221_v9 = new _dafny.CodePoint('v'.codePointAt(0));
      let _222_v10;
      _222_v10 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_217_v5),_221_v9);
      let _223_v11;
      _223_v11 = _dafny.MultiSet.fromElements(_222_v10);
      let _224_v12;
      let _init1 = ((_225_v5) => function (_226_i0) {
        return (_226_i0).plus(_225_v5);
      })(_217_v5);
      let _nw23 = Array((new BigNumber(13)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw23.length); _i0_1++) {
        _nw23[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _224_v12 = _nw23;
      let _227_v13;
      _227_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ybllip"),_dafny.MultiSet.fromElements(_217_v5, _217_v5));
      let _228_v15;
      _228_v15 = _dafny.Seq.of(_218_v6);
      let _229_v16;
      _229_v16 = _module.D0.create_DC1(_212_v0, new BigNumber((_213_v1).length));
      let _230_v17;
      _230_v17 = _module.D0.create_DC2(_229_v16);
      let _231_v18;
      _231_v18 = _dafny.MultiSet.fromElements(_230_v17, _230_v17);
      let _232_v19;
      _232_v19 = _dafny.MultiSet.fromElements(_231_v18);
      let _233_v20;
      _233_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_232_v19).contains(_231_v18)) ? ((_232_v19).get(_231_v18)) : (_217_v5)),_212_v0);
      let _234_globalState;
      let _nw24 = new _module.GlobalState();
      _nw24.__ctor(true, (_dafny.Set.fromElements((((_213_v1).contains(_212_v0)) ? ((_213_v1).get(_212_v0)) : (_212_v0)))).Difference(_214_v2), _215_v3, _216_v4, _220_v8, new BigNumber(-929), new BigNumber(851), true, _223_v11, new BigNumber(90), new BigNumber(-909), new BigNumber(917), new BigNumber(336), _224_v12, new BigNumber(206), new BigNumber(664), new BigNumber(-432), (_227_v13).Merge(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-191)), function (_235_i1) {
          return _dafny.Seq.UnicodeFromString("incqt");
        }), _module.__default.safeIndex(_217_v5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-191)), function (_236_i1) {
          return _dafny.Seq.UnicodeFromString("incqt");
        })).length)), _215_v3)).Elements) {
          let _237_v14 = _compr_7;
          if (_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-191)), function (_235_i1) {
            return _dafny.Seq.UnicodeFromString("incqt");
          }), _module.__default.safeIndex(_217_v5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-191)), function (_236_i1) {
            return _dafny.Seq.UnicodeFromString("incqt");
          })).length)), _215_v3), _237_v14)) {
            _coll7.push([_237_v14,(_module.D0.create_DC0(_dafny.MultiSet.fromElements(_217_v5))).dtor_cf0]);
          }
        }
        return _coll7;
      }()), true, _228_v15, _dafny.Seq.update(_215_v3, _module.__default.safeIndex(_217_v5, new BigNumber((_215_v3).length)), _221_v9), new BigNumber(-661), new BigNumber(-89), _dafny.Seq.of(_217_v5), true, new BigNumber(849), false, _233_v20);
      _234_globalState = _nw24;
      _module.__default.m0(_234_globalState);
      if ((new BigNumber((_216_v4).length)).isEqualTo(new BigNumber((_233_v20).length))) {
        _217_v5 = new BigNumber((_215_v3).length);
        (_234_globalState).f0 = !(!(_module.__default.fm0(_234_globalState)));
        let _238_v21;
        let _nw25 = new _module.C0();
        _nw25.__ctor(_212_v0, _224_v12);
        _238_v21 = _nw25;
        let _239_v22;
        _239_v22 = _module.D0.create_DC1(_212_v0, _module.__default.fm3(_217_v5, _212_v0, _module.__default.fm3(_217_v5, (_238_v21).f28, new BigNumber(544), _233_v20, _234_globalState), _233_v20, _234_globalState));
        let _240_v23;
        _240_v23 = _module.D1.create_DC3(_224_v12);
        let _241_v24;
        let _nw26 = new _module.C0();
        _nw26.__ctor((_239_v22).dtor_cf1, (_240_v23).dtor_cf4);
        _241_v24 = _nw26;
        _241_v24 = _241_v24;
        _221_v9 = (_241_v24).fm1(_module.D0.create_DC1((_238_v21).f28, _217_v5), _217_v5, _217_v5, new BigNumber((_215_v3).length), _234_globalState);
      } else {
        let _index41 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_224_v12).length));
        (_224_v12)[_index41] = new BigNumber((_219_v7).length);
        let _index42 = _module.__default.safeIndex(new BigNumber(59), new BigNumber((_224_v12).length));
        (_224_v12)[_index42] = (_dafny.ZERO).minus(new BigNumber(((_213_v1).Merge((_213_v1).update(_212_v0, _212_v0))).length));
        let _242_v25;
        let _init2 = ((_243_v8) => function (_244_i2) {
          return _module.__default.safeModuloInt(_244_i2, new BigNumber((_243_v8).length));
        })(_220_v8);
        let _nw27 = Array((new BigNumber(21)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw27.length); _i0_2++) {
          _nw27[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _242_v25 = _nw27;
        let _245_v26;
        let _nw28 = new _module.C0();
        _nw28.__ctor(_212_v0, _242_v25);
        _245_v26 = _nw28;
        (_234_globalState).f0 = (_245_v26).f28;
        _217_v5 = (_224_v12)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_224_v12).length))];
        let _246_v27;
        let _nw29 = Array((new BigNumber(11)).toNumber()).fill(false);
        _246_v27 = _nw29;
        let _247_v29;
        _247_v29 = _dafny.MultiSet.fromElements(_233_v20, _dafny.Map.Empty.slice().updateUnsafe((_224_v12)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_224_v12).length))],(_245_v26).f28), function () {
          let _coll8 = new _dafny.Map();
          for (const _compr_8 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-392)), ((_248_v26) => function (_249_i3) {
            return new BigNumber(((_module.D2.create_DC6(_dafny.MultiSet.fromElements((_248_v26).f28))).dtor_cf7).cardinality());
          })(_245_v26))).Elements) {
            let _250_v28 = _compr_8;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-392)), ((_251_v26) => function (_249_i3) {
              return new BigNumber(((_module.D2.create_DC6(_dafny.MultiSet.fromElements((_251_v26).f28))).dtor_cf7).cardinality());
            })(_245_v26)), _250_v28)) {
              _coll8.push([_module.__default.safeDivisionInt(_250_v28, _217_v5),(_245_v26).f28]);
            }
          }
          return _coll8;
        }(), _233_v20, _233_v20);
        let _index43 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_246_v27).length));
        (_246_v27)[_index43] = (new BigNumber((_247_v29).cardinality())).isLessThan((_224_v12)[_module.__default.safeIndex(new BigNumber(59), new BigNumber((_224_v12).length))]);
        let _index44 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_246_v27).length));
        (_246_v27)[_index44] = (_245_v26).f28;
      }
      _212_v0 = _212_v0;
      _217_v5 = _217_v5;
      let _252_v30;
      _252_v30 = _dafny.Seq.of(_215_v3, _215_v3);
      _252_v30 = _dafny.Seq.Concat(_252_v30, _dafny.Seq.Concat(_252_v30, _252_v30));
      let _253_v31;
      _253_v31 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_217_v5), (_dafny.ZERO).minus(_217_v5)));
      _253_v31 = (_module.__default.fm4(_234_globalState)).Difference(_253_v31);
      let _254_v32;
      _254_v32 = _module.D3.create_DC9(_233_v20);
      (_234_globalState).f0 = (((_217_v5).isLessThanOrEqualTo(_module.__default.fm3((_dafny.ZERO).minus(_217_v5), _212_v0, _217_v5, (_254_v32).dtor_cf11, _234_globalState))) ? (!((_dafny.MultiSet.FromArray(_dafny.Seq.of(_217_v5))).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_212_v0,true)).length)))) : (!(_212_v0) || (!(_212_v0))));
      let _255_v33;
      let _init3 = function (_256_i4) {
        return _dafny.Seq.UnicodeFromString("skjjyean");
      };
      let _nw30 = Array((new BigNumber(20)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw30.length); _i0_3++) {
        _nw30[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _255_v33 = _nw30;
      let _257_v34;
      _257_v34 = _dafny.Map.Empty.slice().updateUnsafe(_255_v33,_217_v5);
      _257_v34 = (_257_v34).update(_255_v33, _217_v5);
      _217_v5 = (_217_v5).plus(new BigNumber((_215_v3).length));
      (_234_globalState).f0 = _212_v0;
      (_234_globalState).f0 = !((((_213_v1).contains(_212_v0)) ? ((_213_v1).get(_212_v0)) : ((!(_212_v0)) === (_212_v0))));
      let _258_v35;
      _258_v35 = _dafny.Map.Empty.slice().updateUnsafe(_212_v0,_224_v12);
      let _259_v36;
      let _nw31 = new _module.C0();
      _nw31.__ctor(_212_v0, _224_v12);
      _259_v36 = _nw31;
      let _260_v37;
      _260_v37 = _dafny.Set.fromElements(_259_v36, _259_v36, _259_v36);
      let _261_v38;
      _261_v38 = _dafny.Map.Empty.slice().updateUnsafe(_258_v35,(_260_v37).IsSubsetOf(_260_v37));
      let _262_v39;
      _262_v39 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_259_v36).f28,_224_v12));
      let _263_v40;
      _263_v40 = _dafny.Seq.of((_262_v39)[_module.__default.safeIndex(_217_v5, new BigNumber((_262_v39).length))]);
      _261_v38 = (_261_v38).update(((_263_v40)[_module.__default.safeIndex(_217_v5, new BigNumber((_263_v40).length))]).Merge(_258_v35), true);
      let _264_v41;
      _264_v41 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_212_v0,(_259_v36).f28)).update(true, _212_v0),_259_v36);
      let _hi0 = new BigNumber((_264_v41).length);
      for (let _265_i5 = new BigNumber((_dafny.Seq.Concat(_252_v30, _252_v30)).length); _265_i5.isLessThan(_hi0); _265_i5 = _265_i5.plus(_dafny.ONE)) {
        (_234_globalState).f0 = !(!(!(_219_v7).equals(_219_v7)));
        (_234_globalState).f0 = (_259_v36).f28;
        let _266_v42;
        _266_v42 = _dafny.MultiSet.fromElements(_259_v36);
        let _267_v43;
        let _nw32 = new _module.C0();
        _nw32.__ctor((_dafny.MultiSet.fromElements(_259_v36)).IsSubsetOf(_266_v42), _224_v12);
        _267_v43 = _nw32;
        let _268_v44;
        _268_v44 = _dafny.Map.Empty.slice().updateUnsafe(_265_i5,_dafny.Seq.of((_259_v36).f28, _212_v0, !((_259_v36).f28)));
        _216_v4 = _dafny.Seq.update((((_268_v44).contains(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_267_v43,_217_v5)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_259_v36,new BigNumber(-199))).update(_267_v43, _265_i5))).length))) ? ((_268_v44).get(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_267_v43,_217_v5)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_259_v36,new BigNumber(-199))).update(_267_v43, _265_i5))).length))) : (_dafny.Seq.of(!((_259_v36).f28)))), _module.__default.safeIndex((_265_i5).plus(_217_v5), new BigNumber(((((_268_v44).contains(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_267_v43,_217_v5)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_259_v36,new BigNumber(-199))).update(_267_v43, _265_i5))).length))) ? ((_268_v44).get(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_267_v43,_217_v5)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_259_v36,new BigNumber(-199))).update(_267_v43, _265_i5))).length))) : (_dafny.Seq.of(!((_259_v36).f28))))).length)), (_259_v36).f28);
      }
      _254_v32 = _254_v32;
      let _269_v45;
      _269_v45 = _module.D3.create_DC10((_259_v36).f28, new BigNumber(281), _212_v0, _217_v5, (_259_v36).f29);
      _217_v5 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("iy")).length), _module.__default.fm3(_217_v5, (_259_v36).f28, new BigNumber((_dafny.Seq.of(!((_269_v45).dtor_cf14))).length), _module.__default.fm5(_221_v9, _234_globalState), _234_globalState));
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_255_v33).length))) {
        let _270_i6 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_270_i6)) && ((_270_i6).isLessThan(new BigNumber((_255_v33).length))))) {
          (_255_v33)[(_270_i6)] = _dafny.Seq.UnicodeFromString("lnqf");
        }
      }
      process.stdout.write(_dafny.toString(_212_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v1).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_214_v2).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_215_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_216_v4, _dafny.Seq.of(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_217_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_218_v6, _dafny.Seq.of(new BigNumber(228)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v7).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_220_v8, _dafny.Seq.of(new BigNumber(-228), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_221_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v10).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-228),new _dafny.CodePoint('v'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v11).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-228),new _dafny.CodePoint('v'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_224_v12)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ybllip"),_dafny.MultiSet.fromElements(new BigNumber(228), new BigNumber(228))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_228_v15, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(228))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v16).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_229_v16).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_v17).dtor_cf3).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_230_v17).dtor_cf3).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v18).equals(_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC1(false, _dafny.ONE)), _module.D0.create_DC2(_module.D0.create_DC1(false, _dafny.ONE))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_232_v19).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_module.D0.create_DC2(_module.D0.create_DC1(false, _dafny.ONE)), _module.D0.create_DC2(_module.D0.create_DC1(false, _dafny.ONE)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_233_v20).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_234_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f1).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_234_globalState.f2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_234_globalState).f3, _dafny.Seq.of(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_234_globalState.f4, _dafny.Seq.of(new BigNumber(-228), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f8).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-228),new _dafny.CodePoint('v'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState.f17).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ybllip"),_dafny.MultiSet.fromElements(new BigNumber(228), new BigNumber(228))).updateUnsafe(_dafny.Seq.UnicodeFromString("incqt"),_dafny.MultiSet.fromElements(new BigNumber(228))).updateUnsafe(_dafny.Seq.UnicodeFromString("sk"),_dafny.MultiSet.fromElements(new BigNumber(228))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_234_globalState).f19, _dafny.Seq.of(_dafny.Seq.of(new BigNumber(228))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_234_globalState).f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_234_globalState).f23, _dafny.Seq.of(new BigNumber(228)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_234_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_234_globalState).f27).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_252_v30, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("sk"), _dafny.Seq.UnicodeFromString("sk"), _dafny.Seq.UnicodeFromString("sk"), _dafny.Seq.UnicodeFromString("sk"), _dafny.Seq.UnicodeFromString("sk"), _dafny.Seq.UnicodeFromString("sk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v31).equals(_dafny.MultiSet.fromElements(new BigNumber(462), new BigNumber(-1), new BigNumber(-650)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_254_v32).dtor_cf11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_255_v33)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_257_v34).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_258_v35).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_259_v36).f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_259_v36).f29)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_260_v37).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_261_v38).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_262_v39).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_263_v40).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_264_v41).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v45).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v45).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v45).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v45).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v45).dtor_cf16)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf3) {
      let $dt = new D0(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, _dafny.ZERO);
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
    static create_DC4(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 1 && this.cf4 === other.cf4;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false);
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
    static create_DC7(cf8, cf9) {
      let $dt = new D2(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf7) {
      let $dt = new D2(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D2(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC10(cf12, cf13, cf14, cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf11) {
      let $dt = new D3(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, _dafny.ZERO, false, _dafny.ZERO, []);
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
    static create_DC11(cf17) {
      let $dt = new D4(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17;
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf18) {
      let $dt = new D5(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + this.cf18.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC14(cf20) {
      let $dt = new D6(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC15(cf21, cf22, cf23) {
      let $dt = new D6(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC16(cf24, cf25) {
      let $dt = new D6(2);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D6(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get is_DC13() { return this.$tag === 3; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf24) + ", " + this.cf25.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24) && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC14(_dafny.Seq.of());
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
    static create_DC18(cf27) {
      let $dt = new D7(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC17(cf26) {
      let $dt = new D7(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC19(cf28) {
      let $dt = new D7(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf26 === other.cf26;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18(false);
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
    static create_DC21(cf30, cf31) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC20(cf29) {
      let $dt = new D8(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(_dafny.ZERO, false);
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
      this.f0 = false;
      this.f2 = _dafny.Seq.UnicodeFromString("");
      this.f4 = _dafny.Seq.of();
      this.f17 = _dafny.Map.Empty;
      this._f1 = _dafny.Set.Empty;
      this._f3 = _dafny.Seq.of();
      this._f5 = _dafny.ZERO;
      this._f6 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.MultiSet.Empty;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f11 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f13 = [];
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f16 = _dafny.ZERO;
      this._f18 = false;
      this._f19 = _dafny.Seq.of();
      this._f20 = _dafny.Seq.UnicodeFromString("");
      this._f21 = _dafny.ZERO;
      this._f22 = _dafny.ZERO;
      this._f23 = _dafny.Seq.of();
      this._f24 = false;
      this._f25 = _dafny.ZERO;
      this._f26 = false;
      this._f27 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this).f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f27 = f27;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
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
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f28 = false;
      this._f29 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f28, f29) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!(true)) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }
    };
    fm2(p0, globalState) {
      let _this = this;
      return (_this).f28;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    get f29() {
      let _this = this;
      return _this._f29;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
