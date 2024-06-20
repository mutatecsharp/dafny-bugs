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
      return true;
    };
    static fm1(p0, globalState) {
      return new BigNumber((((_dafny.MultiSet.fromElements(new BigNumber(699), new BigNumber(128))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)))).cardinality()), new BigNumber(-261)))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(64), new BigNumber(565))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(64)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(565)))) {
            _coll0.push([(_0_v0).minus(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).length)),true]);
          }
        }
        return _coll0;
      }()).length))).length))).Difference(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(556), new BigNumber(176))) {
          let _1_v1 = _compr_1;
          if (((new BigNumber(556)).isLessThanOrEqualTo(_1_v1)) && ((_1_v1).isLessThan(new BigNumber(176)))) {
            _coll1.push([_module.__default.safeDivisionInt(_1_v1, new BigNumber((_dafny.Seq.of(new BigNumber(992))).length)),true]);
          }
        }
        return _coll1;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(402),new BigNumber(682))).length))))).cardinality());
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("be"), _dafny.Seq.UnicodeFromString("mfixa"));
    };
    static fm5(p0, p1, globalState) {
      let _source0 = _module.D5.create_DC17(new BigNumber(400), false, !(false), !(false), new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()));
      if (_source0.is_DC17) {
        let _2___mcc_h0 = (_source0).cf27;
        let _3___mcc_h1 = (_source0).cf28;
        let _4___mcc_h2 = (_source0).cf29;
        let _5___mcc_h3 = (_source0).cf30;
        let _6___mcc_h4 = (_source0).cf31;
        let _7_cf31 = _6___mcc_h4;
        let _8_cf30 = _5___mcc_h3;
        let _9_cf29 = _4___mcc_h2;
        let _10_cf28 = _3___mcc_h1;
        let _11_cf27 = _2___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_12_i0) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("clhe"), _dafny.Seq.UnicodeFromString("lh")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("ilyfkpfb"), _dafny.Seq.UnicodeFromString("xvy"), _dafny.Seq.UnicodeFromString("tygu")));
      } else if (_source0.is_DC18) {
        let _13___mcc_h5 = (_source0).cf32;
        let _14___mcc_h6 = (_source0).cf33;
        let _15_cf33 = _14___mcc_h6;
        let _16_cf32 = _13___mcc_h5;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("wfng")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jovnaak"), _dafny.Seq.UnicodeFromString("gepycipjg")));
      } else {
        let _17___mcc_h7 = (_source0).cf26;
        let _18_cf26 = _17___mcc_h7;
        return _dafny.Seq.Concat(_dafny.Seq.of(_18_cf26, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_19_i1) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        })), _dafny.Seq.of(_18_cf26));
      }
    };
    static fm8(p0, p1, p2, globalState) {
      if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(2), new BigNumber((_dafny.Seq.UnicodeFromString("kdpoord")).length)))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_20_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length));
        })).length), new BigNumber(625))).Elements) {
          let _21_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), function (_20_i0) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length));
          })).length), new BigNumber(625)), _21_v0)) {
            _coll2.add(_module.__default.safeDivisionInt(_21_v0, new BigNumber(237)));
          }
        }
        return _coll2;
      }()).length), new BigNumber(547)))))) {
        return _dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)));
      } else {
        return (function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of (_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
              let _22_v1 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))), _22_v1)) {
                _coll4.add(_22_v1);
              }
            }
            return _coll4;
          }()).Elements) {
            let _23_v2 = _compr_3;
            if ((function () {
              let _coll5 = new _dafny.Set();
              for (const _compr_5 of (_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
                let _24_v1 = _compr_5;
                if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))), _24_v1)) {
                  _coll5.add(_24_v1);
                }
              }
              return _coll5;
            }()).contains(_23_v2)) {
              _coll3.add(_23_v2);
            }
          }
          return _coll3;
        }()).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('q'.codePointAt(0))));
      }
    };
    static fm9(p0, globalState) {
      return _module.D0.create_DC3((_dafny.MultiSet.fromElements(true, true)).IsSubsetOf(_dafny.MultiSet.fromElements(false)));
    };
    static fm10(globalState) {
      return _dafny.Seq.of((new BigNumber(-108)).plus(new BigNumber((_dafny.Set.fromElements(false)).length)), (new BigNumber((_dafny.Seq.of(false, !(!(false)), false)).length)).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("domjx")).length)), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bgwkdeg")).length))).length));
    };
    static fm11(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true)));
    };
    static fm12(p0, globalState) {
      return (_dafny.MultiSet.fromElements(!(false))).Intersect(_dafny.MultiSet.fromElements(true));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      if (true) {
        if (true) {
          return _module.D0.create_DC0(new BigNumber((function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of (_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).Elements) {
    let _25_v0 = _compr_6;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0))), _25_v0)) {
      _coll6.add(_25_v0);
    }
  }
  return _coll6;
}()).length));
        } else {
          return _module.D0.create_DC0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(589),false)).length));
        }
      } else if (!(true)) {
        return _module.D0.create_DC0(new BigNumber((_dafny.Seq.of(false, true, true)).length));
      } else {
        return _module.D0.create_DC0(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-76), new BigNumber(-81))).cardinality()));
      }
    };
    static fm14(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(98)), function (_26_i0) {
        return _module.D0.create_DC2(false);
      }), _dafny.Seq.of(_module.D0.create_DC2(true), _module.D0.create_DC2(true), _module.D0.create_DC2(false), _module.D0.create_DC2(false))), _dafny.Seq.of(_module.D0.create_DC2(true)));
    };
    static fm15(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(392)), function (_27_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length))).cardinality()))).cardinality()))).length),_dafny.Seq.of(_module.D5.create_DC18(new BigNumber(20), new BigNumber(354)), _module.D5.create_DC18((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)), new BigNumber(87))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(748),_dafny.Seq.of(_module.D5.create_DC18(new BigNumber(73), new BigNumber((_dafny.Seq.of(true)).length)), _module.D5.create_DC18(new BigNumber((_dafny.Seq.UnicodeFromString("dsqcv")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))).cardinality()))).length)), _module.D5.create_DC18(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_28_i1) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-780),true)).length)))));
    };
    static fm16(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(354)), function (_29_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length),!_dafny.areEqual(_dafny.Seq.of(new BigNumber(339), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-818)), function (_30_i1) {
        return _dafny.MultiSet.fromElements(false, false);
      })).length)), _dafny.Seq.of(new BigNumber(740))));
    };
    static fm17(globalState) {
      if ((new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(141))).length))).length)).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('l'.codePointAt(0)))).length))) {
        return _module.D2.create_DC8(new BigNumber(895), new BigNumber(72));
      } else {
        return _module.D2.create_DC8((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),new _dafny.CodePoint('b'.codePointAt(0)))).length)), new BigNumber(-93));
      }
    };
    static fm18(p0, globalState) {
      return ((_dafny.Set.fromElements(false, true, true)).Difference(_dafny.Set.fromElements(false, !(!(true))))).Difference(_dafny.Set.fromElements(true));
    };
    static fm19(p0, globalState) {
      return ((function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(200), new BigNumber(666))) {
          let _31_v0 = _compr_7;
          if (((new BigNumber(200)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(666)))) {
            _coll7.add(_module.__default.safeDivisionInt(_31_v0, new BigNumber(261)));
          }
        }
        return _coll7;
      }()).Union(_dafny.Set.fromElements(new BigNumber(884)))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(true, false, !(!(false)), true, false)).cardinality())));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _module.D6.create_DC20(_dafny.Seq.of(true), !(false) || (!(!(true))));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return _module.D8.create_DC25((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)))).length),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fvyxx")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.MultiSet.fromElements(false))).length))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(265),new BigNumber(982))));
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return (((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kromkpsyu")).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("le")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length))).length)));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      if ((_dafny.MultiSet.fromElements(!(false), true)).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, false))) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(true, true)), _dafny.Seq.of((_module.D1.create_DC4(_dafny.MultiSet.fromElements(false))).dtor_cf4, _dafny.MultiSet.fromElements(!(false), !(true)), _dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(false, false)));
      } else {
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(false, !(true)));
      }
    };
    static m2(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let r3 = false;
      (globalState).f3 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p2, new BigNumber(102))), (_dafny.ZERO).minus(p2));
      let _32_v0;
      _32_v0 = false;
      if (_32_v0) {
        let _33_v1;
        _33_v1 = _dafny.MultiSet.fromElements(new BigNumber(416), p2);
        let _34_v2;
        _34_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_33_v1).cardinality()),true);
        let _35_v3;
        _35_v3 = _dafny.Seq.of(new BigNumber((_34_v2).length), new BigNumber((_module.__default.fm16(globalState)).length));
        let _36_v4;
        _36_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_35_v3).length),_32_v0);
        r3 = (((_34_v2).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(12)), ((_40_p3, _41_p2) => function (_42_i0) {
          return (_40_p3)[_module.__default.safeIndex(_41_p2, new BigNumber((_40_p3).length))];
        })(p3, p2))).length))) ? ((_34_v2).get(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(12)), ((_37_p3, _38_p2) => function (_39_i0) {
          return (_37_p3)[_module.__default.safeIndex(_38_p2, new BigNumber((_37_p3).length))];
        })(p3, p2))).length))) : (_32_v0));
        let _43_v5;
        let _nw0 = Array((new BigNumber(13)).toNumber());
        _nw0[(_dafny.ZERO).toNumber()] = p2;
        _nw0[(_dafny.ONE).toNumber()] = new BigNumber(-94);
        _nw0[(new BigNumber(2)).toNumber()] = p2;
        _nw0[(new BigNumber(3)).toNumber()] = new BigNumber(119);
        _nw0[(new BigNumber(4)).toNumber()] = new BigNumber(30);
        _nw0[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw0[(new BigNumber(6)).toNumber()] = p2;
        _nw0[(new BigNumber(7)).toNumber()] = (p2).multipliedBy(p2);
        _nw0[(new BigNumber(8)).toNumber()] = p2;
        _nw0[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw0[(new BigNumber(10)).toNumber()] = new BigNumber(340);
        _nw0[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(p2);
        _nw0[(new BigNumber(12)).toNumber()] = (_35_v3)[_module.__default.safeIndex(p2, new BigNumber((_35_v3).length))];
        _43_v5 = _nw0;
        let _index0 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_43_v5).length));
        (_43_v5)[_index0] = p2;
        let _44_v6;
        let _init0 = ((_45_v0) => function (_46_i1) {
          return (_dafny.Set.fromElements(_45_v0)).contains(_45_v0);
        })(_32_v0);
        let _nw1 = Array((new BigNumber(4)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _44_v6 = _nw1;
        let _47_v7;
        _47_v7 = _dafny.Seq.of(_44_v6, _44_v6);
        let _index1 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_43_v5).length));
        let _rhs0 = !((_32_v0) && (_32_v0)) || (_32_v0);
        let _rhs1 = p2;
        let _rhs2 = ((_module.D3.create_DC11(p2)).dtor_cf14).multipliedBy(p2);
        let _rhs3 = _44_v6;
        let _rhs4 = (_47_v7)[_module.__default.safeIndex(new BigNumber((p3).length), new BigNumber((_47_v7).length))];
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        let _lhs2 = _43_v5;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_43_v5).length));
        _lhs0.f0 = _rhs0;
        _lhs1.f3 = _rhs1;
        _lhs2[_lhs3] = _rhs2;
        _44_v6 = _rhs3;
        _44_v6 = _rhs4;
        (globalState).f1 = (_43_v5)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_43_v5).length))];
        let _48_v8;
        let _init1 = ((_49_v0) => function (_50_i2) {
          return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_49_v0, false, _49_v0, _49_v0)), _dafny.MultiSet.fromElements(_49_v0)), _dafny.Seq.of(_dafny.MultiSet.fromElements(_49_v0)));
        })(_32_v0);
        let _nw2 = Array((new BigNumber(17)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
          _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _48_v8 = _nw2;
        let _index2 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_48_v8).length));
        (_48_v8)[_index2] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), ((_51_v0) => function (_52_i3) {
          return _dafny.MultiSet.fromElements(_51_v0, _51_v0, _51_v0);
        })(_32_v0));
        let _53_v9;
        _53_v9 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),p3);
        let _54_v10;
        _54_v10 = new _dafny.CodePoint('j'.codePointAt(0));
        let _index3 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_48_v8).length));
        (_48_v8)[_index3] = _module.__default.fm23(new BigNumber(49), new _dafny.CodePoint('y'.codePointAt(0)), (_53_v9).update(_54_v10, p3), p3, globalState);
        let _index4 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_43_v5).length));
        (_43_v5)[_index4] = p2;
      } else {
        let _55_v11;
        _55_v11 = _dafny.Seq.of(p2);
        let _56_v12;
        _56_v12 = _dafny.Map.Empty.slice().updateUnsafe(_55_v11,p2);
        _56_v12 = (_56_v12).update(_55_v11, p2);
        let _57_v13;
        _57_v13 = _module.D1.create_DC5(_32_v0);
        let _source1 = _57_v13;
        if (_source1.is_DC5) {
          let _58___mcc_h0 = (_source1).cf5;
          let _59_cf5 = _58___mcc_h0;
          (globalState).f1 = p2;
          let _60_v14;
          let _nw3 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _60_v14 = _nw3;
          let _61_v15;
          _61_v15 = _dafny.Seq.of(_60_v14);
          let _62_v16;
          let _nw4 = new _module.C0();
          _nw4.__ctor((_61_v15)[_module.__default.safeIndex(new BigNumber(-78), new BigNumber((_61_v15).length))]);
          _62_v16 = _nw4;
          let _63_v17;
          _63_v17 = _dafny.MultiSet.fromElements(p2);
          let _64_v18;
          _64_v18 = _module.D0.create_DC3(_32_v0);
          let _65_v19;
          _65_v19 = new _dafny.CodePoint('k'.codePointAt(0));
          let _66_v20;
          _66_v20 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_65_v19,_59_cf5), _dafny.Map.Empty.slice().updateUnsafe(_65_v19,_59_cf5));
          let _67_v21;
          _67_v21 = _module.D3.create_DC11(p2);
          let _68_v22;
          _68_v22 = _dafny.Seq.of(_67_v21, _67_v21);
          let _69_v23;
          let _nw5 = Array((new BigNumber(26)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = p2;
          _nw5[(_dafny.ONE).toNumber()] = p2;
          _nw5[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw5[(new BigNumber(3)).toNumber()] = p2;
          _nw5[(new BigNumber(4)).toNumber()] = p2;
          _nw5[(new BigNumber(5)).toNumber()] = new BigNumber((_63_v17).cardinality());
          _nw5[(new BigNumber(6)).toNumber()] = new BigNumber((_module.__default.fm11(_59_cf5, _32_v0, _64_v18, globalState)).length);
          _nw5[(new BigNumber(7)).toNumber()] = new BigNumber(473);
          _nw5[(new BigNumber(8)).toNumber()] = p2;
          _nw5[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw5[(new BigNumber(10)).toNumber()] = p2;
          _nw5[(new BigNumber(11)).toNumber()] = p2;
          _nw5[(new BigNumber(12)).toNumber()] = p2;
          _nw5[(new BigNumber(13)).toNumber()] = p2;
          _nw5[(new BigNumber(14)).toNumber()] = new BigNumber(-114);
          _nw5[(new BigNumber(15)).toNumber()] = p2;
          _nw5[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(p2);
          _nw5[(new BigNumber(17)).toNumber()] = new BigNumber(((_66_v20)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p2, new BigNumber(697))).length), new BigNumber((_66_v20).length))]).length);
          _nw5[(new BigNumber(18)).toNumber()] = p2;
          _nw5[(new BigNumber(19)).toNumber()] = p2;
          _nw5[(new BigNumber(20)).toNumber()] = new BigNumber((_68_v22).length);
          _nw5[(new BigNumber(21)).toNumber()] = p2;
          _nw5[(new BigNumber(22)).toNumber()] = _module.__default.fm1(new BigNumber((p3).length), globalState);
          _nw5[(new BigNumber(23)).toNumber()] = (_55_v11)[_module.__default.safeIndex(p2, new BigNumber((_55_v11).length))];
          _nw5[(new BigNumber(24)).toNumber()] = p2;
          _nw5[(new BigNumber(25)).toNumber()] = p2;
          _69_v23 = _nw5;
          let _70_v24;
          _70_v24 = _dafny.Seq.of(_69_v23);
          _70_v24 = _dafny.Seq.Concat(_70_v24, _dafny.Seq.Concat(_70_v24, _70_v24));
          (globalState).f1 = (_dafny.ZERO).minus((new BigNumber(427)).plus(p2));
        } else if (_source1.is_DC4) {
          let _71___mcc_h1 = (_source1).cf4;
          let _72_cf4 = _71___mcc_h1;
          let _73_v25;
          let _init2 = ((_74_p2) => function (_75_i4) {
            return (_75_i4).minus((_module.D6.create_DC22(_74_p2, _74_p2)).dtor_cf42);
          })(p2);
          let _nw6 = Array((new BigNumber(22)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _73_v25 = _nw6;
          let _index5 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length));
          (_73_v25)[_index5] = new BigNumber(183);
          let _76_v26;
          let _nw7 = Array((new BigNumber(27)).toNumber());
          _nw7[(_dafny.ZERO).toNumber()] = _32_v0;
          _nw7[(_dafny.ONE).toNumber()] = _32_v0;
          _nw7[(new BigNumber(2)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(3)).toNumber()] = false;
          _nw7[(new BigNumber(4)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(5)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(6)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(7)).toNumber()] = !(_32_v0);
          _nw7[(new BigNumber(8)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(9)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(10)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(11)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(12)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(13)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(14)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(15)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(16)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(17)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(18)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(19)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(20)).toNumber()] = !(_32_v0);
          _nw7[(new BigNumber(21)).toNumber()] = false;
          _nw7[(new BigNumber(22)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(23)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(24)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(25)).toNumber()] = _32_v0;
          _nw7[(new BigNumber(26)).toNumber()] = _32_v0;
          _76_v26 = _nw7;
          let _77_v27;
          _77_v27 = _dafny.Map.Empty.slice().updateUnsafe(_76_v26,p3);
          let _78_v28;
          _78_v28 = _dafny.Map.Empty.slice().updateUnsafe(_77_v27,false);
          let _79_v29;
          _79_v29 = _dafny.Map.Empty.slice().updateUnsafe((_55_v11)[_module.__default.safeIndex(p2, new BigNumber((_55_v11).length))],new BigNumber((p3).length));
          let _80_v30;
          _80_v30 = _dafny.Map.Empty.slice().updateUnsafe(_79_v29,_77_v27);
          let _index6 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length));
          let _rhs5 = ((_32_v0) ? (p2) : ((new BigNumber((p0).length)).minus(p2)));
          let _rhs6 = p2;
          let _rhs7 = !((((_78_v28).contains((((_80_v30).contains(_79_v29)) ? ((_80_v30).get(_79_v29)) : (_dafny.Map.Empty.slice().updateUnsafe(_76_v26,_dafny.Seq.UnicodeFromString("ft")))))) ? ((_78_v28).get((((_80_v30).contains(_79_v29)) ? ((_80_v30).get(_79_v29)) : (_dafny.Map.Empty.slice().updateUnsafe(_76_v26,_dafny.Seq.UnicodeFromString("ft")))))) : (true)));
          let _rhs8 = p2;
          let _lhs4 = globalState;
          let _lhs5 = globalState;
          let _lhs6 = _73_v25;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length));
          _lhs4.f1 = _rhs5;
          _lhs5.f3 = _rhs6;
          r3 = _rhs7;
          _lhs6[_lhs7] = _rhs8;
          let _81_v31;
          _81_v31 = _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber(347), _module.__default.fm1(p2, globalState)), (_73_v25)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length))], (_73_v25)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length))]);
          let _82_v32;
          _82_v32 = _dafny.Set.fromElements(_32_v0, _32_v0, _32_v0);
          let _rhs9 = new BigNumber(362);
          let _rhs10 = (new BigNumber(-108)).minus(p2);
          let _rhs11 = p0;
          let _rhs12 = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,p2)).length), (_dafny.ZERO).minus(new BigNumber((_82_v32).length)));
          let _lhs8 = globalState;
          r0 = _rhs9;
          _lhs8.f3 = _rhs10;
          _81_v31 = _rhs11;
          r0 = _rhs12;
          let _83_v33;
          _83_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_dafny.ZERO).minus((_73_v25)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length))]), globalState),_dafny.Seq.Concat(p3, p3));
          let _84_v34;
          _84_v34 = _dafny.Seq.of(p3);
          _83_v33 = (_83_v33).update((_dafny.ZERO).minus((new BigNumber(-273)).minus((_dafny.ZERO).minus(p2))), (_84_v34)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_32_v0, _32_v0, _32_v0, _32_v0, _module.__default.fm0((_73_v25)[_module.__default.safeIndex(new BigNumber(415), new BigNumber((_73_v25).length))], globalState))).length), new BigNumber((_84_v34).length))]);
          let _index7 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_76_v26).length));
          (_76_v26)[_index7] = _32_v0;
          let _index8 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_76_v26).length));
          (_76_v26)[_index8] = !(_32_v0);
        } else {
          let _85___mcc_h2 = (_source1).cf6;
          let _86_cf6 = _85___mcc_h2;
          (globalState).f1 = (p2).minus(p2);
          let _87_v35;
          _87_v35 = _dafny.Set.fromElements(p2, p2, p2);
          _87_v35 = _87_v35;
          r3 = _32_v0;
          (globalState).f1 = p2;
        }
        let _88_v36;
        let _nw8 = new _module.C1();
        _nw8.__ctor();
        _88_v36 = _nw8;
        let _89_v37;
        _89_v37 = new _dafny.CodePoint('j'.codePointAt(0));
        let _90_v38;
        _90_v38 = _dafny.Map.Empty.slice().updateUnsafe(_88_v36,_89_v37);
        _90_v38 = (_90_v38).update(_88_v36, _89_v37);
        _88_v36 = _88_v36;
        let _91_v39;
        let _nw9 = new _module.C1();
        _nw9.__ctor();
        _91_v39 = _nw9;
      }
      let _92_v40;
      _92_v40 = _dafny.Map.Empty.slice().updateUnsafe(p2,true);
      let _93_v41;
      _93_v41 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,(p3)[_module.__default.safeIndex(p2, new BigNumber((p3).length))]);
      let _94_v42;
      _94_v42 = _dafny.Seq.of(_32_v0);
      let _95_v43;
      _95_v43 = _dafny.Set.fromElements(false);
      let _96_v44;
      let _nw10 = Array((new BigNumber(11)).toNumber());
      _nw10[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_32_v0);
      _nw10[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(!(_module.__default.fm0(new BigNumber((p3).length), globalState)), _32_v0, (((_92_v40).contains(new BigNumber((_93_v41).length))) ? ((_92_v40).get(new BigNumber((_93_v41).length))) : (_32_v0)), _32_v0)).Union(_dafny.Set.fromElements(_32_v0, (_94_v42)[_module.__default.safeIndex(p2, new BigNumber((_94_v42).length))], _32_v0));
      _nw10[(new BigNumber(2)).toNumber()] = (_95_v43).Difference(_95_v43);
      _nw10[(new BigNumber(3)).toNumber()] = (_95_v43).Union(_95_v43);
      _nw10[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_32_v0, _32_v0, _32_v0);
      _nw10[(new BigNumber(5)).toNumber()] = _95_v43;
      _nw10[(new BigNumber(6)).toNumber()] = _95_v43;
      _nw10[(new BigNumber(7)).toNumber()] = (_dafny.Set.fromElements(_32_v0, _32_v0)).Intersect(_95_v43);
      _nw10[(new BigNumber(8)).toNumber()] = _module.__default.fm18(p2, globalState);
      _nw10[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_32_v0)).Union(_95_v43);
      _nw10[(new BigNumber(10)).toNumber()] = _95_v43;
      _96_v44 = _nw10;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_96_v44).length))) {
        let _97_i5 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_97_i5)) && ((_97_i5).isLessThan(new BigNumber((_96_v44).length))))) {
          (_96_v44)[(_97_i5)] = _dafny.Set.fromElements(_32_v0);
        }
      }
      let _98_v45;
      _98_v45 = new _dafny.CodePoint('i'.codePointAt(0));
      let _99_v46;
      _99_v46 = _module.D5.create_DC16(_dafny.Seq.update((_module.D6.create_DC21(p2, _32_v0, p3, _32_v0)).dtor_cf39, _module.__default.safeIndex(new BigNumber((p3).length), new BigNumber(((_module.D6.create_DC21(p2, _32_v0, p3, _32_v0)).dtor_cf39).length)), _98_v45));
      _99_v46 = _99_v46;
      (globalState).f3 = _module.__default.fm1((p2).minus(p2), globalState);
      let _100_i6;
      _100_i6 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm0(p2, globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_100_i6)) {
              break L0;
            }
            _100_i6 = (_100_i6).plus(_dafny.ONE);
            let _101_v47;
            _101_v47 = _dafny.Seq.of(p2, new BigNumber((_module.__default.fm10(globalState)).length), p2);
            let _102_v48;
            _102_v48 = _module.D5.create_DC17(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_101_v47, _module.__default.safeIndex(p2, new BigNumber((_101_v47).length)), new BigNumber((_95_v43).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(342)), ((_103_p2) => function (_104_i7) {
  return _103_p2;
})(p2))))).cardinality()), _32_v0, ((_32_v0) ? (!(_32_v0)) : (_32_v0)), (p2).isEqualTo(p2), (p2).multipliedBy(p2));
            let _105_v49;
            _105_v49 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_101_v47).length));
            let _rhs13 = p2;
            let _rhs14 = _102_v48;
            let _rhs15 = _105_v49;
            let _lhs9 = globalState;
            _lhs9.f3 = _rhs13;
            _102_v48 = _rhs14;
            _105_v49 = _rhs15;
            if (_32_v0) {
              let _106_v50;
              let _nw11 = Array((new BigNumber(23)).toNumber()).fill(false);
              _106_v50 = _nw11;
              _106_v50 = _106_v50;
              let _107_v51;
              let _nw12 = new _module.C1();
              _nw12.__ctor();
              _107_v51 = _nw12;
              let _108_v52;
              _108_v52 = _dafny.Set.fromElements(_107_v51, _107_v51);
              _101_v47 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(16)), ((_109_p2) => function (_110_i8) {
                return _109_p2;
              })(p2)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_108_v52).length)), _101_v47));
              let _111_v53;
              _111_v53 = _dafny.Map.Empty.slice().updateUnsafe(_107_v51,p2);
              _111_v53 = (_111_v53).update(_107_v51, new BigNumber((p3).length));
              (globalState).f1 = (p2).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_112_i9) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              })).length));
              r3 = _32_v0;
            } else {
              let _113_v54;
              _113_v54 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-486),_94_v42);
              let _114_v55;
              _114_v55 = _module.D6.create_DC20(_94_v42, _32_v0);
              let _115_v56;
              _115_v56 = _dafny.Seq.of(_dafny.Seq.of((((_113_v54).contains(p2)) ? ((_113_v54).get(p2)) : (_94_v42)), _94_v42, (_114_v55).dtor_cf35));
              (globalState).f1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(p3, _module.__default.safeIndex(new BigNumber(((_115_v56)[_module.__default.safeIndex(p2, new BigNumber((_115_v56).length))]).length), new BigNumber((p3).length)), _98_v45)).length));
              let _116_v57;
              let _nw13 = new _module.C1();
              _nw13.__ctor();
              _116_v57 = _nw13;
              let _117_v58;
              let _init3 = ((_118_v0) => function (_119_i10) {
                return _118_v0;
              })(_32_v0);
              let _nw14 = Array((new BigNumber(15)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw14.length); _i0_3++) {
                _nw14[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _117_v58 = _nw14;
              let _120_v59;
              _120_v59 = _dafny.Set.fromElements(_117_v58);
              let _121_v60;
              _121_v60 = _dafny.Map.Empty.slice().updateUnsafe(_32_v0,new BigNumber(-399));
              let _122_v61;
              _122_v61 = _module.D6.create_DC22(p2, p2);
              let _rhs16 = _32_v0;
              let _rhs17 = (p3)[_module.__default.safeIndex(new BigNumber(((_120_v59).Union(_120_v59)).length), new BigNumber((p3).length))];
              let _rhs18 = ((new BigNumber(649)).plus((((_121_v60).contains(false)) ? ((_121_v60).get(false)) : (p2)))).isLessThan((_122_v61).dtor_cf42);
              r1 = _rhs16;
              _98_v45 = _rhs17;
              r1 = _rhs18;
              let _123_v62;
              let _nw15 = new _module.C1();
              _nw15.__ctor();
              _123_v62 = _nw15;
              let _124_v63;
              let _init4 = ((_125_p3) => function (_126_i11) {
                return _125_p3;
              })(p3);
              let _nw16 = Array((new BigNumber(9)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw16.length); _i0_4++) {
                _nw16[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _124_v63 = _nw16;
              let _127_v64;
              let _nw17 = new _module.C0();
              _nw17.__ctor(_124_v63);
              _127_v64 = _nw17;
            }
            (globalState).f1 = _module.__default.fm1(new BigNumber((_dafny.Seq.UnicodeFromString("ywah")).length), globalState);
            if ((_95_v43).IsProperSubsetOf((_module.__default.fm18(p2, globalState)).Intersect(_95_v43))) {
              r2 = (new BigNumber(-131)).isLessThan(new BigNumber(5));
              let _128_v65;
              let _nw18 = new _module.C1();
              _nw18.__ctor();
              _128_v65 = _nw18;
              let _129_v66;
              let _init5 = ((_130_v47, _131_p2, _132_p3, _133_v45) => function (_134_i12) {
                return (_dafny.MultiSet.fromElements(new BigNumber(448), _131_p2, new BigNumber((_dafny.Seq.update(_132_p3, _module.__default.safeIndex(_131_p2, new BigNumber((_132_p3).length)), _133_v45)).length), _131_p2, _131_p2)).IsSubsetOf(_dafny.MultiSet.FromArray(_130_v47));
              })(_101_v47, p2, p3, _98_v45);
              let _nw19 = Array((new BigNumber(5)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw19.length); _i0_5++) {
                _nw19[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _129_v66 = _nw19;
              let _135_v67;
              _135_v67 = _dafny.MultiSet.fromElements(p3, _dafny.Seq.Concat(p3, _module.__default.fm4(true, p2, p2, new _dafny.CodePoint('e'.codePointAt(0)), globalState)));
              let _136_v68;
              let _nw20 = Array((new BigNumber(13)).toNumber());
              _136_v68 = _nw20;
              let _index9 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_136_v68).length));
              (_136_v68)[_index9] = _128_v65;
              let _index10 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_136_v68).length));
              let _rhs19 = _129_v66;
              let _rhs20 = _135_v67;
              let _rhs21 = _128_v65;
              let _lhs10 = _136_v68;
              let _lhs11 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_136_v68).length));
              _129_v66 = _rhs19;
              _135_v67 = _rhs20;
              _lhs10[_lhs11] = _rhs21;
              (globalState).f0 = _dafny.Seq.IsPrefixOf(p3, _dafny.Seq.Concat(p3, _dafny.Seq.update(p3, _module.__default.safeIndex(new BigNumber(902), new BigNumber((p3).length)), _98_v45)));
              let _137_v69;
              let _nw21 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _137_v69 = _nw21;
              let _138_v70;
              let _nw22 = new _module.C0();
              _nw22.__ctor(_137_v69);
              _138_v70 = _nw22;
              let _139_v71;
              _139_v71 = _dafny.MultiSet.fromElements(p2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_138_v70), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.of(_138_v70)).length)), _138_v70)).length)), p2, p2);
              let _140_v72;
              _140_v72 = _module.D0.create_DC3(_32_v0);
              let _141_v73;
              let _out0;
              _out0 = (_128_v65).m1((_module.D0.create_DC3(_32_v0)).dtor_cf3, _dafny.Set.fromElements((_128_v65).fm2(_139_v71, globalState)), _dafny.Seq.Concat(p3, p3), _140_v72, globalState);
              _141_v73 = _out0;
            } else {
              let _142_v74;
              let _nw23 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _142_v74 = _nw23;
              let _143_v75;
              _143_v75 = _dafny.Map.Empty.slice().updateUnsafe(_142_v74,(new BigNumber(236)).minus(p2));
              _143_v75 = (_143_v75).update(_142_v74, p2);
              let _144_v76;
              let _nw24 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _144_v76 = _nw24;
              let _index11 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_144_v76).length));
              (_144_v76)[_index11] = p2;
              let _145_v77;
              _145_v77 = _module.D1.create_DC5(!(!(_32_v0)));
              let _146_v78;
              let _nw25 = Array((new BigNumber(15)).toNumber()).fill(false);
              _146_v78 = _nw25;
              let _index12 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_146_v78).length));
              (_146_v78)[_index12] = _32_v0;
              let _index13 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_144_v76).length));
              let _index14 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_146_v78).length));
              let _rhs22 = p2;
              let _rhs23 = ((_32_v0) ? (new BigNumber(273)) : (new BigNumber(-363)));
              let _rhs24 = _145_v77;
              let _rhs25 = !((p2).isLessThan(new BigNumber(375)));
              let _lhs12 = _144_v76;
              let _lhs13 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_144_v76).length));
              let _lhs14 = globalState;
              let _lhs15 = _146_v78;
              let _lhs16 = _module.__default.safeIndex(new BigNumber(486), new BigNumber((_146_v78).length));
              _lhs12[_lhs13] = _rhs22;
              _lhs14.f3 = _rhs23;
              _145_v77 = _rhs24;
              _lhs15[_lhs16] = _rhs25;
              let _147_v79;
              _147_v79 = _module.D5.create_DC18(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), ((_148_v45) => function (_149_i13) {
  return _148_v45;
})(_98_v45))).length), new BigNumber((p3).length));
              let _150_v80;
              _150_v80 = _dafny.Map.Empty.slice().updateUnsafe(p3,_147_v79);
              _150_v80 = (_150_v80).update(p3, _147_v79);
              let _index15 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_144_v76).length));
              (_144_v76)[_index15] = _module.__default.safeModuloInt(p2, ((_32_v0) ? ((_144_v76)[_module.__default.safeIndex(new BigNumber(186), new BigNumber((_144_v76).length))]) : (p2)));
              let _151_v81;
              let _nw26 = new _module.C0();
              _nw26.__ctor(_142_v74);
              _151_v81 = _nw26;
            }
          }
        }
      }
      let _152_v82;
      let _nw27 = new _module.C1();
      _nw27.__ctor();
      _152_v82 = _nw27;
      let _153_v83;
      _153_v83 = _dafny.Set.fromElements(_152_v82, _152_v82);
      r0 = new BigNumber((_153_v83).length);
      let _154_v84;
      _154_v84 = _dafny.Seq.of(_92_v40);
      let _155_v85;
      _155_v85 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update(p3, _module.__default.safeIndex(new BigNumber(((_154_v84)[_module.__default.safeIndex(new BigNumber(718), new BigNumber((_154_v84).length))]).length), new BigNumber((p3).length)), _98_v45)).length), p2);
      let _156_v86;
      _156_v86 = _dafny.Set.fromElements(_155_v85);
      let _157_v87;
      _157_v87 = _dafny.Seq.of(_155_v85);
      r1 = ((_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), ((_158_v41) => function (_159_i14) {
        return new BigNumber((_158_v41).length);
      })(_93_v41)))).Difference(function () {
        let _coll8 = new _dafny.Set();
        for (const _compr_8 of (_157_v87).Elements) {
          let _160_v88 = _compr_8;
          if (_dafny.Seq.contains(_157_v87, _160_v88)) {
            _coll8.add(_160_v88);
          }
        }
        return _coll8;
      }())).IsSubsetOf(_156_v86);
      r2 = _32_v0;
      let _161_v89;
      let _nw28 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _161_v89 = _nw28;
      let _162_v90;
      let _nw29 = new _module.C0();
      _nw29.__ctor(_161_v89);
      _162_v90 = _nw29;
      let _163_v91;
      _163_v91 = _dafny.Seq.of(_162_v90);
      r3 = _module.__default.fm0(_module.__default.safeDivisionInt(p2, new BigNumber((_163_v91).length)), globalState);
      return [r0, r1, r2, r3];
    }
    static Main(__noArgsParameter) {
      let _164_v0;
      let _init6 = function (_165_i0) {
        return _dafny.MultiSet.FromArray(_dafny.Seq.of(false, true));
      };
      let _nw30 = Array((new BigNumber(13)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw30.length); _i0_6++) {
        _nw30[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _164_v0 = _nw30;
      let _166_v1;
      _166_v1 = false;
      let _167_v2;
      _167_v2 = _dafny.Seq.of(_166_v1, !(_166_v1));
      let _168_globalState;
      let _nw31 = new _module.GlobalState();
      _nw31.__ctor(true, new BigNumber(827), true, new BigNumber(-651), _164_v0, _dafny.Seq.Concat(_167_v2, _167_v2));
      _168_globalState = _nw31;
      let _169_v3;
      _169_v3 = new BigNumber(272);
      let _170_v4;
      _170_v4 = _dafny.Map.Empty.slice().updateUnsafe(_169_v3,_169_v3);
      let _hi0 = (_169_v3).minus((((_170_v4).contains((_dafny.ZERO).minus(_169_v3))) ? ((_170_v4).get((_dafny.ZERO).minus(_169_v3))) : (_169_v3)));
      for (let _171_i1 = _169_v3; _171_i1.isLessThan(_hi0); _171_i1 = _171_i1.plus(_dafny.ONE)) {
        let _172_v5;
        _172_v5 = new _dafny.CodePoint('k'.codePointAt(0));
        let _173_v6;
        _173_v6 = _dafny.MultiSet.fromElements(_172_v5, _172_v5);
        let _174_v7;
        _174_v7 = _dafny.Seq.of(_173_v6, _173_v6, _dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0))));
        let _175_v8;
        _175_v8 = _module.D0.create_DC0(_169_v3);
        _174_v7 = _dafny.Seq.Concat(_174_v7, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_176_v5) => function (_177_i2) {
          return _dafny.MultiSet.fromElements(_176_v5, _176_v5);
        })(_172_v5)), _module.__default.safeIndex((_175_v8).dtor_cf0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_178_v5) => function (_179_i2) {
          return _dafny.MultiSet.fromElements(_178_v5, _178_v5);
        })(_172_v5))).length)), _173_v6));
        _170_v4 = (_170_v4).update(_171_i1, (_171_i1).minus((_dafny.ZERO).minus(_171_i1)));
        let _180_v9;
        _180_v9 = _dafny.Set.fromElements(_module.__default.fm0(new BigNumber((_167_v2).length), _168_globalState), _166_v1);
        let _181_v10;
        _181_v10 = _dafny.MultiSet.fromElements(_180_v9, _180_v9);
        let _182_v11;
        _182_v11 = _module.D0.create_DC2((_181_v10).IsProperSubsetOf(_181_v10));
        let _183_v12;
        _183_v12 = _dafny.Seq.UnicodeFromString("g");
        _182_v11 = ((!(new BigNumber((_dafny.Seq.update(_183_v12, _module.__default.safeIndex(_169_v3, new BigNumber((_183_v12).length)), _172_v5)).length)).isEqualTo(_169_v3)) ? (_182_v11) : (_182_v11));
        (_168_globalState).f3 = _171_i1;
      }
      let _184_v13;
      _184_v13 = _dafny.Map.Empty.slice().updateUnsafe(_166_v1,(_dafny.ZERO).minus(_169_v3));
      let _185_v14;
      _185_v14 = _module.D0.create_DC2(true);
      _184_v13 = (_184_v13).update((_185_v14).dtor_cf2, _169_v3);
      _166_v1 = _166_v1;
      let _186_v15;
      let _init7 = ((_187_v3) => function (_188_i3) {
        return (_188_i3).plus(new BigNumber((_dafny.Set.fromElements(_187_v3)).length));
      })(_169_v3);
      let _nw32 = Array((new BigNumber(20)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw32.length); _i0_7++) {
        _nw32[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _186_v15 = _nw32;
      let _index16 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
      (_186_v15)[_index16] = _169_v3;
      let _index17 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
      (_186_v15)[_index17] = _169_v3;
      let _189_v16;
      _189_v16 = _dafny.Map.Empty.slice().updateUnsafe(_186_v15,new BigNumber(-954));
      let _hi1 = new BigNumber((_189_v16).length);
      for (let _190_i4 = _module.__default.fm1(_169_v3, _168_globalState); _190_i4.isLessThan(_hi1); _190_i4 = _190_i4.plus(_dafny.ONE)) {
        let _191_v17;
        _191_v17 = _dafny.Seq.UnicodeFromString("ldovmaed");
        let _192_v18;
        _192_v18 = _dafny.Map.Empty.slice().updateUnsafe(_169_v3,_dafny.Seq.Concat(_191_v17, _191_v17));
        _192_v18 = (_192_v18).update(_module.__default.safeDivisionInt(_169_v3, _module.__default.fm1(_169_v3, _168_globalState)), _191_v17);
        let _193_v19;
        _193_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_170_v4).length),_166_v1);
        let _194_v20;
        _194_v20 = _dafny.Set.fromElements(_166_v1, _166_v1);
        _166_v1 = (((_193_v19).contains(_190_i4)) ? ((_193_v19).get(_190_i4)) : ((_194_v20).IsProperSubsetOf(_194_v20)));
        if (false) {
          let _195_v21;
          let _nw33 = new _module.C1();
          _nw33.__ctor();
          _195_v21 = _nw33;
          let _196_v22;
          _196_v22 = _dafny.MultiSet.fromElements((((_192_v18).contains((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])) ? ((_192_v18).get((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])) : (_191_v17)));
          let _197_v23;
          let _nw34 = Array((new BigNumber(5)).toNumber()).fill([]);
          _197_v23 = _nw34;
          let _198_v24;
          _198_v24 = _module.D3.create_DC12(_169_v3, _196_v22, new BigNumber((_167_v2).length), _197_v23);
          let _199_v25;
          _199_v25 = _dafny.MultiSet.fromElements(_198_v24, _198_v24);
          let _200_v26;
          let _nw35 = Array((new BigNumber(14)).toNumber()).fill(false);
          _200_v26 = _nw35;
          let _201_v27;
          _201_v27 = _dafny.Map.Empty.slice().updateUnsafe(_169_v3,_200_v26);
          let _202_v28;
          _202_v28 = _dafny.Map.Empty.slice().updateUnsafe(_199_v25,(((_201_v27).contains(_169_v3)) ? ((_201_v27).get(_169_v3)) : (_200_v26)));
          _202_v28 = (_202_v28).update(_199_v25, _200_v26);
          let _203_v29;
          _203_v29 = _dafny.Map.Empty.slice().updateUnsafe(((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]).minus(_169_v3),((_166_v1) ? (_186_v15) : (_186_v15)));
          _186_v15 = (((_203_v29).contains(_190_i4)) ? ((_203_v29).get(_190_i4)) : (_186_v15));
          let _204_v30;
          let _205_v31;
          let _206_v32;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_195_v21).m0(_168_globalState);
          _out1 = _outcollector0[0];
          _out2 = _outcollector0[1];
          _out3 = _outcollector0[2];
          _204_v30 = _out1;
          _205_v31 = _out2;
          _206_v32 = _out3;
          let _207_v33;
          let _init8 = ((_208_v17) => function (_209_i5) {
            return _208_v17;
          })(_191_v17);
          let _nw36 = Array((new BigNumber(29)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw36.length); _i0_8++) {
            _nw36[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _207_v33 = _nw36;
          let _210_v34;
          let _nw37 = new _module.C0();
          _nw37.__ctor(_207_v33);
          _210_v34 = _nw37;
          let _211_v35;
          let _nw38 = Array((new BigNumber(28)).toNumber());
          _nw38[(_dafny.ZERO).toNumber()] = _210_v34;
          _nw38[(_dafny.ONE).toNumber()] = _210_v34;
          _nw38[(new BigNumber(2)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(3)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(4)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(5)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(6)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(7)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(8)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(9)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(10)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(11)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(12)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(13)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(14)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(15)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(16)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(17)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(18)).toNumber()] = ((_166_v1) ? (_210_v34) : (_210_v34));
          _nw38[(new BigNumber(19)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(20)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(21)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(22)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(23)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(24)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(25)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(26)).toNumber()] = _210_v34;
          _nw38[(new BigNumber(27)).toNumber()] = _210_v34;
          _211_v35 = _nw38;
          _211_v35 = _211_v35;
        } else {
          (_168_globalState).f0 = !(!(_166_v1) || (((true) ? (_166_v1) : (_166_v1))));
          let _index18 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
          (_186_v15)[_index18] = new BigNumber(((_193_v19).Merge(_193_v19)).length);
          _166_v1 = !(_166_v1) || (_dafny.Seq.IsPrefixOf(_167_v2, _167_v2));
          (_168_globalState).f0 = _module.__default.fm0((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _168_globalState);
          let _212_v36;
          let _init9 = ((_213_v1) => function (_214_i6) {
            return _213_v1;
          })(_166_v1);
          let _nw39 = Array((new BigNumber(13)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw39.length); _i0_9++) {
            _nw39[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _212_v36 = _nw39;
          _212_v36 = _212_v36;
        }
        let _215_v37;
        _215_v37 = _dafny.Seq.of(_170_v4, (_170_v4).update(new BigNumber((_191_v17).length), _190_i4));
        let _216_v38;
        _216_v38 = _module.D3.create_DC9(_215_v37);
        let _217_v39;
        _217_v39 = _dafny.Seq.of(_216_v38);
        let _218_v40;
        _218_v40 = _dafny.Map.Empty.slice().updateUnsafe(_217_v39,_191_v17);
        let _rhs26 = _218_v40;
        let _rhs27 = _186_v15;
        _218_v40 = _rhs26;
        _186_v15 = _rhs27;
      }
      let _219_v41;
      let _nw40 = Array((new BigNumber(22)).toNumber()).fill(false);
      _219_v41 = _nw40;
      let _index19 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length));
      (_219_v41)[_index19] = _module.__default.fm0((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _168_globalState);
      let _220_v42;
      _220_v42 = _dafny.MultiSet.fromElements(_166_v1);
      let _pat_let_tv0 = _169_v3;
      let _pat_let_tv1 = _186_v15;
      let _pat_let_tv2 = _186_v15;
      let _pat_let_tv3 = _220_v42;
      let _pat_let_tv4 = _169_v3;
      let _pat_let_tv5 = _166_v1;
      let _pat_let_tv6 = _169_v3;
      let _index20 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length));
      (_219_v41)[_index20] = function (_source2) {
        if (_source2.is_DC5) {
          let _221___mcc_h0 = (_source2).cf5;
          let _222_cf5 = _221___mcc_h0;
          return (_dafny.MultiSet.fromElements(_pat_let_tv0, (_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_pat_let_tv1).length))], new BigNumber((_pat_let_tv3).cardinality()))).IsDisjointFrom(_dafny.MultiSet.fromElements(_pat_let_tv4));
        } else if (_source2.is_DC4) {
          let _223___mcc_h1 = (_source2).cf4;
          let _224_cf4 = _223___mcc_h1;
          return _pat_let_tv5;
        } else {
          let _225___mcc_h2 = (_source2).cf6;
          let _226_cf6 = _225___mcc_h2;
          return (new BigNumber(371)).isLessThanOrEqualTo(_pat_let_tv6);
        }
      }(_module.D1.create_DC4(_220_v42));
      let _227_i7;
      _227_i7 = _dafny.ZERO;
      L1: {
        while ((_169_v3).isLessThan(_169_v3)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_227_i7)) {
              break L1;
            }
            _227_i7 = (_227_i7).plus(_dafny.ONE);
            let _228_v43;
            let _nw41 = Array((new BigNumber(2)).toNumber()).fill([]);
            _228_v43 = _nw41;
            let _index21 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_228_v43).length));
            (_228_v43)[_index21] = _219_v41;
            let _229_v44;
            _229_v44 = new _dafny.CodePoint('w'.codePointAt(0));
            let _230_v45;
            _230_v45 = _dafny.MultiSet.fromElements(_229_v44, _229_v44, _229_v44, _229_v44);
            let _index22 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_228_v43).length));
            let _nw42 = Array((new BigNumber(12)).toNumber());
            _nw42[(_dafny.ZERO).toNumber()] = ((_166_v1) ? ((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) : ((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]));
            _nw42[(_dafny.ONE).toNumber()] = _166_v1;
            _nw42[(new BigNumber(2)).toNumber()] = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
            _nw42[(new BigNumber(3)).toNumber()] = _166_v1;
            _nw42[(new BigNumber(4)).toNumber()] = (_230_v45).IsSubsetOf(_230_v45);
            _nw42[(new BigNumber(5)).toNumber()] = ((_166_v1) ? ((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) : (_166_v1));
            _nw42[(new BigNumber(6)).toNumber()] = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
            _nw42[(new BigNumber(7)).toNumber()] = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
            _nw42[(new BigNumber(8)).toNumber()] = _166_v1;
            _nw42[(new BigNumber(9)).toNumber()] = _166_v1;
            _nw42[(new BigNumber(10)).toNumber()] = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
            _nw42[(new BigNumber(11)).toNumber()] = _166_v1;
            (_228_v43)[_index22] = _nw42;
            if (false) {
              (_168_globalState).f0 = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
              let _231_v46;
              _231_v46 = _module.D1.create_DC4(_220_v42);
              let _232_v47;
              _232_v47 = _dafny.Seq.of(_169_v3, _169_v3, _169_v3, _169_v3);
              let _index23 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
              let _rhs28 = _231_v46;
              let _rhs29 = _module.__default.fm1((_dafny.ZERO).minus(new BigNumber((_232_v47).length)), _168_globalState);
              let _rhs30 = _169_v3;
              let _lhs17 = _168_globalState;
              let _lhs18 = _186_v15;
              let _lhs19 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
              _231_v46 = _rhs28;
              _lhs17.f3 = _rhs29;
              _lhs18[_lhs19] = _rhs30;
              let _233_v48;
              _233_v48 = _dafny.Seq.UnicodeFromString("tabouo");
              let _234_v49;
              let _nw43 = new _module.C1();
              _nw43.__ctor();
              _234_v49 = _nw43;
              let _235_v50;
              _235_v50 = _dafny.MultiSet.fromElements(_234_v49);
              let _236_v51;
              _236_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_233_v48).length),(_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              let _rhs31 = _dafny.Seq.Concat(_233_v48, _233_v48);
              let _rhs32 = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
              let _rhs33 = (((_dafny.ZERO).minus(_169_v3)).multipliedBy((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])).isLessThan(_169_v3);
              let _rhs34 = (_236_v51).contains(_module.__default.safeModuloInt(new BigNumber((_233_v48).length), new BigNumber((_235_v50).cardinality())));
              let _rhs35 = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
              let _lhs20 = _168_globalState;
              let _lhs21 = _168_globalState;
              _233_v48 = _rhs31;
              _lhs20.f0 = _rhs32;
              _166_v1 = _rhs33;
              _166_v1 = _rhs34;
              _lhs21.f1 = _rhs35;
              let _237_v52;
              let _nw44 = new _module.C1();
              _nw44.__ctor();
              _237_v52 = _nw44;
              let _238_v53;
              let _239_v54;
              let _240_v55;
              let _out4;
              let _out5;
              let _out6;
              let _outcollector1 = (_234_v49).m0(_168_globalState);
              _out4 = _outcollector1[0];
              _out5 = _outcollector1[1];
              _out6 = _outcollector1[2];
              _238_v53 = _out4;
              _239_v54 = _out5;
              _240_v55 = _out6;
            } else {
              let _241_v56;
              _241_v56 = _dafny.Set.fromElements(_166_v1);
              let _242_v57;
              _242_v57 = _241_v56;
              let _243_v58;
              let _nw45 = Array((new BigNumber(9)).toNumber());
              _nw45[(_dafny.ZERO).toNumber()] = _241_v56;
              _nw45[(_dafny.ONE).toNumber()] = (_242_v57);
              _nw45[(new BigNumber(2)).toNumber()] = _241_v56;
              _nw45[(new BigNumber(3)).toNumber()] = _241_v56;
              _nw45[(new BigNumber(4)).toNumber()] = _module.__default.fm18(new BigNumber(914), _168_globalState);
              _nw45[(new BigNumber(5)).toNumber()] = _241_v56;
              _nw45[(new BigNumber(6)).toNumber()] = (_dafny.Set.fromElements(!((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]), _166_v1)).Intersect(_dafny.Set.fromElements(_166_v1));
              _nw45[(new BigNumber(7)).toNumber()] = (_241_v56).Intersect(_241_v56);
              _nw45[(new BigNumber(8)).toNumber()] = _241_v56;
              _243_v58 = _nw45;
              let _index24 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_243_v58).length));
              (_243_v58)[_index24] = _module.__default.fm18(new BigNumber(359), _168_globalState);
              let _index25 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_243_v58).length));
              (_243_v58)[_index25] = _dafny.Set.fromElements((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              let _244_v59;
              let _nw46 = Array((new BigNumber(27)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _244_v59 = _nw46;
              let _index26 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_244_v59).length));
              (_244_v59)[_index26] = _229_v44;
              let _index27 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_244_v59).length));
              (_244_v59)[_index27] = _229_v44;
              let _245_v60;
              _245_v60 = _dafny.Set.fromElements(_169_v3, (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _169_v3, new BigNumber(221));
              let _246_v61;
              let _init10 = ((_247_v2, _248_v1) => function (_249_i8) {
                return _module.D6.create_DC20(_247_v2, _248_v1);
              })(_167_v2, _166_v1);
              let _nw47 = Array((new BigNumber(23)).toNumber());
              for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw47.length); _i0_10++) {
                _nw47[_i0_10] = _init10(new BigNumber(_i0_10));
              }
              _246_v61 = _nw47;
              let _250_v62;
              _250_v62 = _dafny.Seq.UnicodeFromString("gdhmm");
              let _251_v63;
              let _252_v64;
              let _253_v65;
              let _254_v66;
              let _out7;
              let _out8;
              let _out9;
              let _out10;
              let _outcollector2 = _module.__default.m2(_245_v60, _246_v61, (((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) ? (new BigNumber(((_243_v58)[_module.__default.safeIndex(new BigNumber(606), new BigNumber((_243_v58).length))]).length)) : (_169_v3)), _dafny.Seq.update(_dafny.Seq.Concat(_250_v62, _250_v62), _module.__default.safeIndex((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], new BigNumber((_dafny.Seq.Concat(_250_v62, _250_v62)).length)), (_244_v59)[_module.__default.safeIndex(new BigNumber(403), new BigNumber((_244_v59).length))]), _168_globalState);
              _out7 = _outcollector2[0];
              _out8 = _outcollector2[1];
              _out9 = _outcollector2[2];
              _out10 = _outcollector2[3];
              _251_v63 = _out7;
              _252_v64 = _out8;
              _253_v65 = _out9;
              _254_v66 = _out10;
              let _255_v67;
              let _256_v68;
              let _257_v69;
              let _258_v70;
              let _out11;
              let _out12;
              let _out13;
              let _out14;
              let _outcollector3 = _module.__default.m2((_245_v60).Difference(_module.__default.fm19((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], _168_globalState)), _246_v61, ((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]).minus((_dafny.ZERO).minus(_169_v3)), _250_v62, _168_globalState);
              _out11 = _outcollector3[0];
              _out12 = _outcollector3[1];
              _out13 = _outcollector3[2];
              _out14 = _outcollector3[3];
              _255_v67 = _out11;
              _256_v68 = _out12;
              _257_v69 = _out13;
              _258_v70 = _out14;
              _258_v70 = (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))];
            }
            if (_166_v1) {
              let _259_v71;
              _259_v71 = _dafny.Map.Empty.slice().updateUnsafe(_169_v3,!(_166_v1));
              _259_v71 = (_259_v71).update((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _166_v1);
              let _260_v72;
              let _nw48 = new _module.C1();
              _nw48.__ctor();
              _260_v72 = _nw48;
              let _261_v73;
              _261_v73 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))],true)).length),_260_v72);
              let _262_v74;
              _262_v74 = _module.D6.create_DC20(_167_v2, _166_v1);
              let _263_v75;
              let _nw49 = Array((new BigNumber(24)).toNumber());
              _nw49[(_dafny.ZERO).toNumber()] = _262_v74;
              _nw49[(_dafny.ONE).toNumber()] = _262_v74;
              _nw49[(new BigNumber(2)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(3)).toNumber()] = _module.D6.create_DC20(_167_v2, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              _nw49[(new BigNumber(4)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(5)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(6)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(7)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(8)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(9)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(10)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(11)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(12)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(13)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(14)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(15)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(16)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(17)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(18)).toNumber()] = _module.D6.create_DC20(_167_v2, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              _nw49[(new BigNumber(19)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(20)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(21)).toNumber()] = _262_v74;
              _nw49[(new BigNumber(22)).toNumber()] = _module.D6.create_DC20(_167_v2, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              _nw49[(new BigNumber(23)).toNumber()] = _262_v74;
              _263_v75 = _nw49;
              let _264_v76;
              _264_v76 = _dafny.Seq.UnicodeFromString("quqguci");
              let _265_v77;
              let _266_v78;
              let _267_v79;
              let _268_v80;
              let _out15;
              let _out16;
              let _out17;
              let _out18;
              let _outcollector4 = _module.__default.m2(_dafny.Set.fromElements(new BigNumber((_261_v73).length), (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _169_v3, (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]), (((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) ? (_263_v75) : (_263_v75)), _169_v3, _264_v76, _168_globalState);
              _out15 = _outcollector4[0];
              _out16 = _outcollector4[1];
              _out17 = _outcollector4[2];
              _out18 = _outcollector4[3];
              _265_v77 = _out15;
              _266_v78 = _out16;
              _267_v79 = _out17;
              _268_v80 = _out18;
              let _269_v81;
              _269_v81 = _dafny.Seq.of(new BigNumber(295), _265_v77, _265_v77);
              _267_v79 = !(!_dafny.Seq.contains(_269_v81, (_169_v3).plus((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])));
              let _270_v82;
              _270_v82 = _dafny.Set.fromElements(_265_v77, _169_v3);
              let _271_v83;
              _271_v83 = _dafny.Map.Empty.slice().updateUnsafe(_267_v79,_270_v82);
              let _nw50 = Array((new BigNumber(24)).toNumber());
              _nw50[(_dafny.ZERO).toNumber()] = _265_v77;
              _nw50[(_dafny.ONE).toNumber()] = _265_v77;
              _nw50[(new BigNumber(2)).toNumber()] = new BigNumber(((((_271_v83).contains(_266_v78)) ? ((_271_v83).get(_266_v78)) : (_module.__default.fm19((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], _168_globalState)))).length);
              _nw50[(new BigNumber(3)).toNumber()] = (_169_v3).plus(_169_v3);
              _nw50[(new BigNumber(4)).toNumber()] = _265_v77;
              _nw50[(new BigNumber(5)).toNumber()] = (new BigNumber(338)).minus(new BigNumber(39));
              _nw50[(new BigNumber(6)).toNumber()] = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
              _nw50[(new BigNumber(7)).toNumber()] = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
              _nw50[(new BigNumber(8)).toNumber()] = (_169_v3).multipliedBy(_265_v77);
              _nw50[(new BigNumber(9)).toNumber()] = _169_v3;
              _nw50[(new BigNumber(10)).toNumber()] = (_265_v77).minus(_265_v77);
              _nw50[(new BigNumber(11)).toNumber()] = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
              _nw50[(new BigNumber(12)).toNumber()] = _169_v3;
              _nw50[(new BigNumber(13)).toNumber()] = new BigNumber(525);
              _nw50[(new BigNumber(14)).toNumber()] = (_169_v3).minus(_265_v77);
              _nw50[(new BigNumber(15)).toNumber()] = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
              _nw50[(new BigNumber(16)).toNumber()] = (_265_v77).minus((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]);
              _nw50[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_169_v3);
              _nw50[(new BigNumber(18)).toNumber()] = (((_170_v4).contains(new BigNumber((_220_v42).cardinality()))) ? ((_170_v4).get(new BigNumber((_220_v42).cardinality()))) : ((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]));
              _nw50[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_266_v78,_266_v78)).length);
              _nw50[(new BigNumber(20)).toNumber()] = _169_v3;
              _nw50[(new BigNumber(21)).toNumber()] = _265_v77;
              _nw50[(new BigNumber(22)).toNumber()] = _265_v77;
              _nw50[(new BigNumber(23)).toNumber()] = new BigNumber((_167_v2).length);
              _186_v15 = _nw50;
              _268_v80 = !(_module.__default.fm0(_module.__default.fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(344)), ((_272_v44) => function (_273_i9) {
                return _272_v44;
              })(_229_v44))).length), _168_globalState), _168_globalState));
            } else {
              let _274_v84;
              _274_v84 = _dafny.MultiSet.fromElements(_219_v41, _219_v41);
              (_168_globalState).f0 = !(_274_v84).equals(_274_v84);
              let _275_v85;
              _275_v85 = _module.D2.create_DC8(_module.__default.safeModuloInt(_169_v3, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(31))).cardinality())), new BigNumber(-850));
              _275_v85 = _275_v85;
              let _276_v86;
              _276_v86 = _module.D4.create_DC13((_228_v43)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_228_v43).length))]);
              let _277_v87;
              _277_v87 = _dafny.Map.Empty.slice().updateUnsafe(_166_v1,_276_v86);
              let _278_v88;
              let _init11 = ((_279_v44) => function (_280_i10) {
                return _dafny.Seq.Create(_module.__default.abs(new BigNumber(571)), ((_281_v44) => function (_282_i11) {
                  return _281_v44;
                })(_279_v44));
              })(_229_v44);
              let _nw51 = Array((new BigNumber(13)).toNumber());
              for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw51.length); _i0_11++) {
                _nw51[_i0_11] = _init11(new BigNumber(_i0_11));
              }
              _278_v88 = _nw51;
              let _283_v89;
              let _nw52 = new _module.C0();
              _nw52.__ctor(_278_v88);
              _283_v89 = _nw52;
              let _284_v90;
              _284_v90 = _dafny.Map.Empty.slice().updateUnsafe(_277_v87,_283_v89);
              _284_v90 = (_284_v90).update((_dafny.Map.Empty.slice().updateUnsafe(_166_v1,_276_v86)).Merge(_277_v87), _283_v89);
              let _285_v91;
              _285_v91 = _dafny.Set.fromElements((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], new BigNumber((_dafny.Seq.of(_169_v3)).length), (_dafny.ZERO).minus(_169_v3), (_dafny.ZERO).minus(_169_v3));
              let _286_v92;
              _286_v92 = _module.D6.create_DC20(_167_v2, _166_v1);
              let _287_v93;
              _287_v93 = _dafny.Map.Empty.slice().updateUnsafe(_229_v44,(_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              let _pat_let_tv7 = _219_v41;
              let _pat_let_tv8 = _219_v41;
              let _288_v94;
              let _nw53 = Array((new BigNumber(28)).toNumber());
              _nw53[(_dafny.ZERO).toNumber()] = _286_v92;
              _nw53[(_dafny.ONE).toNumber()] = _286_v92;
              _nw53[(new BigNumber(2)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(3)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(4)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(5)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(6)).toNumber()] = function (_pat_let0_0) {
                return function (_289_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_290_dt__update_hcf35_h0) {
                      return _module.D6.create_DC20(_290_dt__update_hcf35_h0, (_289_dt__update__tmp_h0).dtor_cf36);
                    }(_pat_let1_0);
                  }(_dafny.Seq.of((_pat_let_tv8)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_pat_let_tv7).length))]));
                }(_pat_let0_0);
              }(_286_v92);
              _nw53[(new BigNumber(7)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(8)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(9)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(10)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(11)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(12)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(13)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(14)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(15)).toNumber()] = _module.D6.create_DC20(_167_v2, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              _nw53[(new BigNumber(16)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(17)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(18)).toNumber()] = _module.__default.fm20(_287_v93, _166_v1, (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _168_globalState);
              _nw53[(new BigNumber(19)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(20)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(21)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(22)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(23)).toNumber()] = _module.D6.create_DC20(_dafny.Seq.of(_166_v1, false, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], (_module.D5.create_DC17((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _166_v1, _166_v1, false, (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])).dtor_cf29, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]), _166_v1);
              _nw53[(new BigNumber(24)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(25)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(26)).toNumber()] = _286_v92;
              _nw53[(new BigNumber(27)).toNumber()] = _module.D6.create_DC20(_167_v2, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]);
              _288_v94 = _nw53;
              let _291_v95;
              _291_v95 = _dafny.Seq.UnicodeFromString("sbgug");
              let _292_v96;
              let _293_v97;
              let _294_v98;
              let _295_v99;
              let _out19;
              let _out20;
              let _out21;
              let _out22;
              let _outcollector5 = _module.__default.m2(_285_v91, _288_v94, _169_v3, _291_v95, _168_globalState);
              _out19 = _outcollector5[0];
              _out20 = _outcollector5[1];
              _out21 = _outcollector5[2];
              _out22 = _outcollector5[3];
              _292_v96 = _out19;
              _293_v97 = _out20;
              _294_v98 = _out21;
              _295_v99 = _out22;
              let _296_v100;
              _296_v100 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(_169_v3, _169_v3));
              _296_v100 = _296_v100;
            }
            _229_v44 = _229_v44;
          }
        }
      }
      let _297_v101;
      let _nw54 = new _module.C1();
      _nw54.__ctor();
      _297_v101 = _nw54;
      let _298_v102;
      _298_v102 = _dafny.Seq.of(_297_v101, _297_v101, _297_v101, _297_v101);
      _297_v101 = (_298_v102)[_module.__default.safeIndex(((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]).multipliedBy((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]), new BigNumber((_298_v102).length))];
      let _299_v103;
      _299_v103 = _dafny.Map.Empty.slice().updateUnsafe(_169_v3,_167_v2);
      let _300_v104;
      _300_v104 = _dafny.Seq.UnicodeFromString("fb");
      let _index28 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
      (_186_v15)[_index28] = ((_166_v1) ? (new BigNumber(((_299_v103).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false, false, _166_v1, false)).length),_167_v2))).length)) : (new BigNumber((_300_v104).length)));
      let _301_v105;
      _301_v105 = new _dafny.CodePoint('d'.codePointAt(0));
      let _302_v106;
      _302_v106 = _dafny.Set.fromElements((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]);
      let _rhs36 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_169_v3, _module.__default.fm1(new BigNumber(-808), _168_globalState)));
      let _rhs37 = _dafny.Seq.update(_300_v104, _module.__default.safeIndex(_169_v3, new BigNumber((_300_v104).length)), _301_v105);
      let _rhs38 = (_module.__default.fm21((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _302_v106, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], _168_globalState)).dtor_cf45;
      _169_v3 = _rhs36;
      _300_v104 = _rhs37;
      _170_v4 = _rhs38;
      let _303_v107;
      _303_v107 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xhuyl"));
      let _304_v108;
      _304_v108 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), ((_305_v105) => function (_306_i12) {
        return _305_v105;
      })(_301_v105)));
      _303_v107 = ((_dafny.MultiSet.fromElements(_300_v104, _300_v104)).Intersect(_303_v107)).Union(_dafny.MultiSet.FromArray(_304_v108));
      let _307_v109;
      let _308_v110;
      let _309_v111;
      let _out23;
      let _out24;
      let _out25;
      let _outcollector6 = (_297_v101).m0(_168_globalState);
      _out23 = _outcollector6[0];
      _out24 = _outcollector6[1];
      _out25 = _outcollector6[2];
      _307_v109 = _out23;
      _308_v110 = _out24;
      _309_v111 = _out25;
      let _310_v112;
      let _311_v113;
      let _312_v114;
      let _out26;
      let _out27;
      let _out28;
      let _outcollector7 = (_297_v101).m0(_168_globalState);
      _out26 = _outcollector7[0];
      _out27 = _outcollector7[1];
      _out28 = _outcollector7[2];
      _310_v112 = _out26;
      _311_v113 = _out27;
      _312_v114 = _out28;
      if (_166_v1) {
        _169_v3 = _module.__default.safeDivisionInt((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _169_v3);
        _297_v101 = _297_v101;
        let _313_v115;
        let _nw55 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
        _313_v115 = _nw55;
        let _314_v116;
        _314_v116 = _dafny.Set.fromElements(true);
        let _index29 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_313_v115).length));
        (_313_v115)[_index29] = (_314_v116).Difference(_314_v116);
        let _index30 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_313_v115).length));
        let _index31 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
        let _rhs39 = (_module.__default.fm18(_169_v3, _168_globalState)).Intersect(_314_v116);
        let _rhs40 = _module.__default.safeModuloInt((_module.__default.fm1(_module.__default.fm1((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))], _168_globalState), _168_globalState)).plus((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]), (new BigNumber((_314_v116).length)).minus(new BigNumber((_184_v13).length)));
        let _rhs41 = _module.__default.fm1(_310_v112, _168_globalState);
        let _lhs22 = _313_v115;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(342), new BigNumber((_313_v115).length));
        let _lhs24 = _186_v15;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length));
        let _lhs26 = _168_globalState;
        _lhs22[_lhs23] = _rhs39;
        _lhs24[_lhs25] = _rhs40;
        _lhs26.f3 = _rhs41;
        let _315_v117;
        let _nw56 = Array((new BigNumber(17)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = (_184_v13).update(_166_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(563)), ((_316_v105) => function (_317_i13) {
          return _316_v105;
        })(_301_v105))).length));
        _nw56[(_dafny.ONE).toNumber()] = _184_v13;
        _nw56[(new BigNumber(2)).toNumber()] = _module.__default.fm22(_166_v1, _301_v105, _dafny.Seq.UnicodeFromString("gdnjcodfx"), !(!(_166_v1)), _168_globalState);
        _nw56[(new BigNumber(3)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(4)).toNumber()] = (_184_v13).Merge(_184_v13);
        _nw56[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))],(_dafny.ZERO).minus(new BigNumber((_170_v4).length)));
        _nw56[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_166_v1,(_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])).Merge(_184_v13);
        _nw56[(new BigNumber(7)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(8)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(9)).toNumber()] = (((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) ? (_184_v13) : (_184_v13));
        _nw56[(new BigNumber(10)).toNumber()] = (_184_v13).Merge(_184_v13);
        _nw56[(new BigNumber(11)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(12)).toNumber()] = (_184_v13).Merge(_184_v13);
        _nw56[(new BigNumber(13)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(14)).toNumber()] = _184_v13;
        _nw56[(new BigNumber(15)).toNumber()] = (_184_v13).Merge(_184_v13);
        _nw56[(new BigNumber(16)).toNumber()] = _184_v13;
        _315_v117 = _nw56;
        let _rhs42 = ((function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of (_170_v4).Keys.Elements) {
            let _318_v118 = _compr_9;
            if ((_170_v4).contains(_318_v118)) {
              _coll9.push([(_318_v118).multipliedBy((_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))]),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_169_v3)).cardinality()), _310_v112, _307_v109, _169_v3, (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])).length)]);
            }
          }
          return _coll9;
        }()).Merge(_170_v4)).Merge(_170_v4);
        let _rhs43 = _315_v117;
        _170_v4 = _rhs42;
        _315_v117 = _rhs43;
        _169_v3 = (((_307_v109).isLessThanOrEqualTo(_307_v109)) ? ((_169_v3).multipliedBy(new BigNumber(-212))) : (_310_v112));
      } else {
        let _319_v119;
        _319_v119 = _module.D3.create_DC11(new BigNumber(174));
        let _320_v120;
        _320_v120 = _module.D0.create_DC0((_319_v119).dtor_cf14);
        let _321_v121;
        _321_v121 = _dafny.Seq.of(_310_v112);
        let _322_v122;
        _322_v122 = _dafny.Seq.of(_321_v121);
        let _323_v123;
        _323_v123 = _module.D4.create_DC14(_310_v112, _320_v120, (_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))], _166_v1, (_322_v122)[_module.__default.safeIndex(_307_v109, new BigNumber((_322_v122).length))]);
        let _324_v124;
        _324_v124 = _module.D4.create_DC15(_module.D4.create_DC15(_323_v123));
        _324_v124 = _324_v124;
        (_168_globalState).f1 = _310_v112;
        let _325_v125;
        _325_v125 = _dafny.Seq.of(_186_v15, ((true) ? (_186_v15) : (_186_v15)));
        let _326_v126;
        _326_v126 = _dafny.MultiSet.fromElements(_310_v112);
        let _327_v127;
        _327_v127 = _module.D6.create_DC19(_325_v125);
        let _rhs44 = _297_v101;
        let _rhs45 = _module.__default.fm1((((_326_v126).contains(_169_v3)) ? ((_326_v126).get(_169_v3)) : (_310_v112)), _168_globalState);
        let _rhs46 = _dafny.Seq.Concat(_dafny.Seq.Concat(_325_v125, (_327_v127).dtor_cf34), _dafny.Seq.of(_186_v15, _186_v15));
        let _rhs47 = _311_v113;
        _297_v101 = _rhs44;
        _307_v109 = _rhs45;
        _325_v125 = _rhs46;
        _311_v113 = _rhs47;
        let _328_v128;
        let _nw57 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _328_v128 = _nw57;
        let _index32 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length));
        (_328_v128)[_index32] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), ((_329_v105) => function (_330_i14) {
          return (_module.D0.create_DC1(_329_v105)).dtor_cf1;
        })(_301_v105));
        let _index33 = _module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length));
        (_328_v128)[_index33] = _dafny.Seq.Concat(_dafny.Seq.Concat(_300_v104, _300_v104), _300_v104);
        let _331_v129;
        let _nw58 = Array((new BigNumber(16)).toNumber()).fill([]);
        _331_v129 = _nw58;
        let _332_v130;
        _332_v130 = _dafny.Map.Empty.slice().updateUnsafe((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))],_331_v129);
        let _rhs48 = (_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))];
        let _rhs49 = (((_332_v130).contains(_dafny.Seq.IsPrefixOf((_328_v128)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length))], (_328_v128)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length))]))) ? ((_332_v130).get(_dafny.Seq.IsPrefixOf((_328_v128)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length))], (_328_v128)[_module.__default.safeIndex(new BigNumber(429), new BigNumber((_328_v128).length))]))) : (_331_v129));
        _307_v109 = _rhs48;
        _331_v129 = _rhs49;
      }
      _170_v4 = (_170_v4).update(_310_v112, _307_v109);
      let _333_v131;
      _333_v131 = _dafny.Map.Empty.slice().updateUnsafe(_301_v105,_307_v109);
      let _334_v132;
      _334_v132 = _dafny.Seq.of(_333_v131);
      let _hi2 = new BigNumber(((_334_v132)[_module.__default.safeIndex(_310_v112, new BigNumber((_334_v132).length))]).length);
      for (let _335_i15 = (new BigNumber((_dafny.Seq.UnicodeFromString("jywtmvy")).length)).minus((_dafny.ZERO).minus(new BigNumber(-98))); _335_i15.isLessThan(_hi2); _335_i15 = _335_i15.plus(_dafny.ONE)) {
        let _336_v133;
        _336_v133 = _dafny.MultiSet.fromElements(new BigNumber((_302_v106).length), new BigNumber((_300_v104).length), _335_i15);
        let _index34 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length));
        (_219_v41)[_index34] = (_297_v101).fm2(_336_v133, _168_globalState);
        let _337_v134;
        let _nw59 = Array((new BigNumber(5)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _184_v13;
        _nw59[(_dafny.ONE).toNumber()] = ((_166_v1) ? (_184_v13) : (_dafny.Map.Empty.slice().updateUnsafe(_166_v1,(_186_v15)[_module.__default.safeIndex(new BigNumber(743), new BigNumber((_186_v15).length))])));
        _nw59[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))],_307_v109);
        _nw59[(new BigNumber(3)).toNumber()] = _184_v13;
        _nw59[(new BigNumber(4)).toNumber()] = _184_v13;
        _337_v134 = _nw59;
        let _index35 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_337_v134).length));
        (_337_v134)[_index35] = _184_v13;
        let _index36 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_337_v134).length));
        (_337_v134)[_index36] = _184_v13;
        _166_v1 = !(!(_166_v1));
        let _338_v135;
        let _init12 = ((_339_v104) => function (_340_i16) {
          return _339_v104;
        })(_300_v104);
        let _nw60 = Array((new BigNumber(24)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw60.length); _i0_12++) {
          _nw60[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _338_v135 = _nw60;
        let _341_v136;
        let _nw61 = new _module.C0();
        _nw61.__ctor(_338_v135);
        _341_v136 = _nw61;
        let _342_v137;
        _342_v137 = _dafny.Map.Empty.slice().updateUnsafe(_341_v136,_297_v101);
        let _343_v138;
        _343_v138 = _dafny.Map.Empty.slice().updateUnsafe((_341_v136).f6,_297_v101);
        let _344_v139;
        let _nw62 = Array((new BigNumber(28)).toNumber());
        _nw62[(_dafny.ZERO).toNumber()] = _297_v101;
        _nw62[(_dafny.ONE).toNumber()] = _297_v101;
        _nw62[(new BigNumber(2)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(3)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(4)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(5)).toNumber()] = (((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) ? ((((_342_v137).contains(_341_v136)) ? ((_342_v137).get(_341_v136)) : (_297_v101))) : (_297_v101));
        _nw62[(new BigNumber(6)).toNumber()] = (((_343_v138).contains(_338_v135)) ? ((_343_v138).get(_338_v135)) : (_297_v101));
        _nw62[(new BigNumber(7)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(8)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(9)).toNumber()] = (_298_v102)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_298_v102).length))];
        _nw62[(new BigNumber(10)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(11)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(12)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(13)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(14)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(15)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(16)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(17)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(18)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(19)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(20)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(21)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(22)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(23)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(24)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(25)).toNumber()] = (((_219_v41)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_219_v41).length))]) ? (_297_v101) : (_297_v101));
        _nw62[(new BigNumber(26)).toNumber()] = _297_v101;
        _nw62[(new BigNumber(27)).toNumber()] = _297_v101;
        _344_v139 = _nw62;
        _344_v139 = _344_v139;
      }
      process.stdout.write(_dafny.toString(((_164_v0)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_164_v0)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_166_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_167_v2, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_168_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_168_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_168_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_168_globalState.f4)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_168_globalState).f5, _dafny.Seq.of(false, true, false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_169_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_170_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,new BigNumber(5)).updateUnsafe(_dafny.ONE,new BigNumber(2)).updateUnsafe(new BigNumber(265),new BigNumber(982)).updateUnsafe(new BigNumber(534),new BigNumber(534)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-272)).updateUnsafe(true,new BigNumber(272)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v14).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v15)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_189_v16).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v41)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v42).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_227_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_298_v102).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_299_v103).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(272),_dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_300_v104).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_301_v105));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v106).equals(_dafny.Set.fromElements(new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_303_v107).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_304_v108, _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_307_v109));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_309_v111).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_310_v112));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_312_v114).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_333_v131).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(534)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_334_v132, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(534))))));
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
    static create_DC2(cf2) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D0(2);
      $dt.cf3 = cf3;
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
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf3 === other.cf3;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC5(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf4) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC6(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false);
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
    static create_DC8(cf8, cf9) {
      let $dt = new D2(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC7(cf7) {
      let $dt = new D2(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf7) + ")";
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
        return other.$tag === 1 && this.cf7 === other.cf7;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC11(cf14) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC12(cf15, cf16, cf17, cf18) {
      let $dt = new D3(2);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf10) {
      let $dt = new D3(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf11 === other.cf11 && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, false, false);
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
    static create_DC14(cf20, cf21, cf22, cf23, cf24) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
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
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC14(_dafny.ZERO, _module.D0.Default(), false, false, _dafny.Seq.of());
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
    static create_DC17(cf27, cf28, cf29, cf30, cf31) {
      let $dt = new D5(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC18(cf32, cf33) {
      let $dt = new D5(1);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC16(cf26) {
      let $dt = new D5(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC17" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC16" + "(" + this.cf26.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && this.cf28 === other.cf28 && this.cf29 === other.cf29 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC17(_dafny.ZERO, false, false, false, _dafny.ZERO);
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
    static create_DC20(cf35, cf36) {
      let $dt = new D6(0);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC21(cf37, cf38, cf39, cf40) {
      let $dt = new D6(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC22(cf41, cf42) {
      let $dt = new D6(2);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC19(cf34) {
      let $dt = new D6(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC23(cf43) {
      let $dt = new D6(4);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get is_DC23() { return this.$tag === 4; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ", " + this.cf39.toVerbatimString(true) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC22" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39) && this.cf40 === other.cf40;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC20(_dafny.Seq.of(), false);
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
    static create_DC24(cf44) {
      let $dt = new D7(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC24" + "(" + _dafny.toString(this.cf44) + ")";
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf46, cf47, cf48, cf49, cf50) {
      let $dt = new D8(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC25(cf45) {
      let $dt = new D8(1);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC27(cf51) {
      let $dt = new D8(2);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC26" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC25" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC27" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && this.cf47 === other.cf47 && this.cf48 === other.cf48 && this.cf49 === other.cf49 && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC26(false, [], false, false, false);
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
    static create_DC28(cf52) {
      let $dt = new D9(0);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52);
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
          return D9.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = false;
      this.f1 = _dafny.ZERO;
      this.f3 = _dafny.ZERO;
      this.f4 = [];
      this._f2 = false;
      this._f5 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      return;
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f6 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("uy");
    };
    fm7(p0, globalState) {
      let _this = this;
      return (new BigNumber(725)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, false),true)).length));
    };
    get f6() {
      let _this = this;
      return _this._f6;
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
    fm2(p0, globalState) {
      let _this = this;
      return !(function (_source3) {
        if (_source3.is_DC1) {
          let _345___mcc_h0 = (_source3).cf1;
          let _346_cf1 = _345___mcc_h0;
          return false;
        } else if (_source3.is_DC2) {
          let _347___mcc_h1 = (_source3).cf2;
          let _348_cf2 = _347___mcc_h1;
          return _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(58), new BigNumber(-449)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(773)), function (_349_i0) {
            return new BigNumber(3);
          }));
        } else if (_source3.is_DC3) {
          let _350___mcc_h2 = (_source3).cf3;
          let _351_cf3 = _350___mcc_h2;
          if (_351_cf3) {
            return false;
          } else {
            return _351_cf3;
          }
        } else {
          let _352___mcc_h3 = (_source3).cf0;
          let _353_cf0 = _352___mcc_h3;
          return true;
        }
      }(_module.D0.create_DC0((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hw")).length)))));
    };
    fm3(p0, globalState) {
      let _this = this;
      return new _dafny.CodePoint('c'.codePointAt(0));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _module.D0.Default();
      let _354_v0;
      let _nw63 = Array((new BigNumber(29)).toNumber()).fill(false);
      _354_v0 = _nw63;
      r1 = _354_v0;
      r1 = _354_v0;
      let _355_v1;
      _355_v1 = false;
      if (_355_v1) {
        let _356_v2;
        let _init13 = function (_357_i0) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pl"), _dafny.Seq.UnicodeFromString("lg"));
        };
        let _nw64 = Array((new BigNumber(7)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw64.length); _i0_13++) {
          _nw64[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _356_v2 = _nw64;
        let _358_v3;
        _358_v3 = new BigNumber(640);
        let _359_v4;
        let _nw65 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _359_v4 = _nw65;
        let _360_v5;
        _360_v5 = _dafny.Seq.of(_359_v4, _359_v4, _359_v4);
        let _361_v6;
        _361_v6 = _dafny.Set.fromElements(_358_v3, new BigNumber((_360_v5).length));
        let _362_v7;
        _362_v7 = new _dafny.CodePoint('b'.codePointAt(0));
        let _index37 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length));
        (_356_v2)[_index37] = _module.__default.fm4(_355_v1, _358_v3, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_361_v6,_355_v1)).length), _362_v7, globalState);
        let _363_v8;
        _363_v8 = _dafny.Seq.UnicodeFromString("eottg");
        let _index38 = _module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length));
        (_356_v2)[_index38] = _363_v8;
        let _index39 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_354_v0).length));
        (_354_v0)[_index39] = _355_v1;
        let _364_v9;
        _364_v9 = _dafny.Seq.of(_dafny.Seq.Concat(_363_v8, _363_v8), _dafny.Seq.UnicodeFromString("srkfki"), _dafny.Seq.update(_dafny.Seq.Concat((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))], _363_v8), _module.__default.safeIndex(new BigNumber(((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))]).length), new BigNumber((_dafny.Seq.Concat((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))], _363_v8)).length)), _362_v7), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-93)), function (_365_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), (_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))]);
        let _366_v10;
        _366_v10 = _module.D0.create_DC1(_362_v7);
        let _index40 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_354_v0).length));
        let _rhs50 = _355_v1;
        let _rhs51 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))], _module.__default.fm4(_355_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_354_v0,!(_355_v1))).length), new BigNumber(-828), _362_v7, globalState), (_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))]), _dafny.Seq.update(_dafny.Seq.update(_module.__default.fm5(_358_v3, _358_v3, globalState), _module.__default.safeIndex(_358_v3, new BigNumber((_module.__default.fm5(_358_v3, _358_v3, globalState)).length)), _dafny.Seq.UnicodeFromString("xkrtmgfeb")), _module.__default.safeIndex(_358_v3, new BigNumber((_dafny.Seq.update(_module.__default.fm5(_358_v3, _358_v3, globalState), _module.__default.safeIndex(_358_v3, new BigNumber((_module.__default.fm5(_358_v3, _358_v3, globalState)).length)), _dafny.Seq.UnicodeFromString("xkrtmgfeb"))).length)), _363_v8)), _364_v9);
        let _rhs52 = function (_pat_let2_0) {
          return function (_367_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_368_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_368_dt__update_hcf1_h0);
              }(_pat_let3_0);
            }(new _dafny.CodePoint('g'.codePointAt(0)));
          }(_pat_let2_0);
        }(_366_v10);
        let _lhs27 = _354_v0;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_354_v0).length));
        _lhs27[_lhs28] = _rhs50;
        _364_v9 = _rhs51;
        _366_v10 = _rhs52;
        let _369_v11;
        _369_v11 = _module.D0.create_DC0(new BigNumber(-402));
        let _370_v12;
        _370_v12 = _dafny.Seq.of(_369_v11, _369_v11, _369_v11, _module.D0.create_DC0(_358_v3));
        (globalState).f3 = new BigNumber((_370_v12).length);
        (globalState).f1 = _module.__default.safeDivisionInt(new BigNumber(272), _358_v3);
        let _371_v13;
        _371_v13 = _dafny.Map.Empty.slice().updateUnsafe(_362_v7,(_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))]);
        _371_v13 = (_371_v13).update(_362_v7, _dafny.Seq.update((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))], _module.__default.safeIndex(_358_v3, new BigNumber(((_356_v2)[_module.__default.safeIndex(new BigNumber(275), new BigNumber((_356_v2).length))]).length)), _362_v7));
      } else {
        let _372_v14;
        _372_v14 = new BigNumber(382);
        (globalState).f1 = ((_dafny.ZERO).minus(_372_v14)).multipliedBy(_372_v14);
        let _373_v15;
        _373_v15 = _dafny.Seq.UnicodeFromString("kp");
        if (!(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), function (_374_i2) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), _373_v15)) || (_355_v1)) {
          (globalState).f1 = _372_v14;
          let _375_v16;
          _375_v16 = new _dafny.CodePoint('v'.codePointAt(0));
          let _376_v17;
          _376_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC1(_375_v16)), _module.D0.create_DC1(_375_v16)),_372_v14);
          let _377_v18;
          _377_v18 = _dafny.Map.Empty.slice().updateUnsafe(_355_v1,_dafny.Seq.of(_355_v1));
          let _378_v19;
          _378_v19 = _dafny.Seq.of(_355_v1);
          _376_v17 = (_376_v17).update(_355_v1, (_dafny.ZERO).minus(new BigNumber(((_377_v18).update(_355_v1, _378_v19)).length)));
          let _379_v20;
          let _nw66 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _379_v20 = _nw66;
          let _380_v21;
          let _nw67 = new _module.C0();
          _nw67.__ctor(_379_v20);
          _380_v21 = _nw67;
          let _381_v22;
          let _nw68 = Array((new BigNumber(9)).toNumber()).fill([]);
          _381_v22 = _nw68;
          let _index41 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_381_v22).length));
          (_381_v22)[_index41] = (_380_v21).f6;
          let _382_v23;
          _382_v23 = _dafny.Seq.of((_380_v21).f6);
          let _383_v24;
          _383_v24 = _dafny.Seq.of((_380_v21).f6, (_380_v21).f6, _379_v20, (_382_v23)[_module.__default.safeIndex(new BigNumber((_378_v19).length), new BigNumber((_382_v23).length))], (_380_v21).f6);
          let _index42 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_381_v22).length));
          (_381_v22)[_index42] = (_382_v23)[_module.__default.safeIndex(_372_v14, new BigNumber((_382_v23).length))];
          (globalState).f1 = _372_v14;
        } else {
          let _384_v25;
          let _nw69 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _384_v25 = _nw69;
          let _385_v26;
          let _nw70 = new _module.C0();
          _nw70.__ctor(_384_v25);
          _385_v26 = _nw70;
          let _386_v27;
          let _nw71 = new _module.C0();
          _nw71.__ctor(_384_v25);
          _386_v27 = _nw71;
          let _387_v28;
          _387_v28 = _dafny.Map.Empty.slice().updateUnsafe((_372_v14).isLessThanOrEqualTo(_module.__default.fm1(new BigNumber((_373_v15).length), globalState)),_372_v14);
          let _388_v29;
          _388_v29 = new _dafny.CodePoint('o'.codePointAt(0));
          _387_v28 = (_387_v28).update(!(!((_module.__default.fm8(_355_v1, _355_v1, new BigNumber((_373_v15).length), globalState)).IsProperSubsetOf(_dafny.Set.fromElements(_388_v29)))), _372_v14);
          (globalState).f0 = _355_v1;
          r0 = _372_v14;
        }
        (globalState).f3 = new BigNumber(240);
        let _389_v30;
        _389_v30 = new _dafny.CodePoint('n'.codePointAt(0));
        let _390_v31;
        let _nw72 = Array((new BigNumber(8)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = _module.__default.fm4(_355_v1, _372_v14, _372_v14, _389_v30, globalState);
        _nw72[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("jcokmtins");
        _nw72[(new BigNumber(2)).toNumber()] = _373_v15;
        _nw72[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_373_v15, _module.__default.safeIndex(new BigNumber((_373_v15).length), new BigNumber((_373_v15).length)), _389_v30);
        _nw72[(new BigNumber(4)).toNumber()] = _373_v15;
        _nw72[(new BigNumber(5)).toNumber()] = _373_v15;
        _nw72[(new BigNumber(6)).toNumber()] = _373_v15;
        _nw72[(new BigNumber(7)).toNumber()] = _373_v15;
        _390_v31 = _nw72;
        let _391_v32;
        let _nw73 = new _module.C0();
        _nw73.__ctor(_390_v31);
        _391_v32 = _nw73;
        let _392_v33;
        let _nw74 = Array((new BigNumber(15)).toNumber());
        _nw74[(_dafny.ZERO).toNumber()] = _391_v32;
        _nw74[(_dafny.ONE).toNumber()] = _391_v32;
        _nw74[(new BigNumber(2)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(3)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(4)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(5)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(6)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(7)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(8)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(9)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(10)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(11)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(12)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(13)).toNumber()] = _391_v32;
        _nw74[(new BigNumber(14)).toNumber()] = _391_v32;
        _392_v33 = _nw74;
        let _393_v34;
        _393_v34 = _dafny.Set.fromElements(_392_v33);
        _393_v34 = (_393_v34).Intersect(_393_v34);
        let _394_v35;
        let _nw75 = Array((_dafny.ONE).toNumber());
        _nw75[(_dafny.ZERO).toNumber()] = _372_v14;
        _394_v35 = _nw75;
        let _395_v36;
        _395_v36 = _dafny.Set.fromElements(_372_v14);
        let _index43 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_394_v35).length));
        (_394_v35)[_index43] = new BigNumber(((_395_v36).Difference(_395_v36)).length);
        let _index44 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_394_v35).length));
        (_394_v35)[_index44] = _372_v14;
      }
      let _396_v37;
      _396_v37 = new BigNumber(534);
      let _397_v38;
      _397_v38 = _dafny.MultiSet.fromElements(_355_v1);
      let _398_v39;
      _398_v39 = _dafny.Set.fromElements(_module.D0.create_DC2(_355_v1), _module.D0.create_DC2(_355_v1));
      let _399_v40;
      _399_v40 = _dafny.Seq.of(_398_v39, _398_v39);
      let _400_v41;
      _400_v41 = _dafny.Map.Empty.slice().updateUnsafe(_355_v1,(_dafny.ZERO).minus(_396_v37));
      let _401_v42;
      _401_v42 = _dafny.Seq.UnicodeFromString("cqusjevmf");
      let _402_v43;
      let _nw76 = Array((new BigNumber(20)).toNumber());
      _nw76[(_dafny.ZERO).toNumber()] = new BigNumber(195);
      _nw76[(_dafny.ONE).toNumber()] = _396_v37;
      _nw76[(new BigNumber(2)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(3)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_396_v37, globalState);
      _nw76[(new BigNumber(5)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_397_v38,_396_v37)).length);
      _nw76[(new BigNumber(7)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(8)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(9)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(10)).toNumber()] = new BigNumber((_399_v40).length);
      _nw76[(new BigNumber(11)).toNumber()] = (_396_v37).plus(_396_v37);
      _nw76[(new BigNumber(12)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(13)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(14)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt((((_400_v41).contains(_355_v1)) ? ((_400_v41).get(_355_v1)) : (_396_v37)), new BigNumber((_401_v42).length));
      _nw76[(new BigNumber(16)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(17)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(18)).toNumber()] = _396_v37;
      _nw76[(new BigNumber(19)).toNumber()] = _396_v37;
      _402_v43 = _nw76;
      _402_v43 = _402_v43;
      let _403_v44;
      let _init14 = ((_404_v1) => function (_405_i3) {
        return _module.D0.create_DC3(_404_v1);
      })(_355_v1);
      let _nw77 = Array((new BigNumber(22)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw77.length); _i0_14++) {
        _nw77[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _403_v44 = _nw77;
      let _index45 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_403_v44).length));
      (_403_v44)[_index45] = _module.__default.fm9(new BigNumber((_dafny.MultiSet.fromElements(_396_v37, _396_v37, new BigNumber((_401_v42).length), _396_v37, _396_v37)).cardinality()), globalState);
      let _406_v45;
      _406_v45 = _dafny.Map.Empty.slice().updateUnsafe(_396_v37,_355_v1);
      let _index46 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_403_v44).length));
      let _rhs53 = _355_v1;
      let _rhs54 = _module.D0.create_DC3((_355_v1) || (_355_v1));
      let _rhs55 = (_dafny.ZERO).minus(_396_v37);
      let _rhs56 = (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(342),_355_v1)).Merge(_406_v45)).length));
      let _lhs29 = globalState;
      let _lhs30 = _403_v44;
      let _lhs31 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_403_v44).length));
      let _lhs32 = globalState;
      let _lhs33 = globalState;
      _lhs29.f0 = _rhs53;
      _lhs30[_lhs31] = _rhs54;
      _lhs32.f1 = _rhs55;
      _lhs33.f1 = _rhs56;
      _396_v37 = _396_v37;
      r0 = _396_v37;
      r1 = _354_v0;
      r2 = (_403_v44)[_module.__default.safeIndex(new BigNumber(692), new BigNumber((_403_v44).length))];
      return [r0, r1, r2];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _407_v0;
      let _nw78 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _407_v0 = _nw78;
      let _408_v1;
      let _nw79 = new _module.C0();
      _nw79.__ctor(_407_v0);
      _408_v1 = _nw79;
      let _409_v2;
      _409_v2 = new BigNumber(652);
      let _410_v3;
      _410_v3 = _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_409_v2);
      let _411_v4;
      _411_v4 = _dafny.Seq.of(_408_v1);
      _410_v3 = (_410_v3).update(_408_v1, ((p0) ? (new BigNumber(824)) : (new BigNumber((_411_v4).length))));
      let _412_i0;
      _412_i0 = _dafny.ZERO;
      L2: {
        while (p0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_412_i0)) {
              break L2;
            }
            _412_i0 = (_412_i0).plus(_dafny.ONE);
            let _413_v5;
            _413_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_409_v2),_module.D0.create_DC0(_409_v2));
            let _414_v6;
            _414_v6 = _dafny.Seq.of(_413_v5);
            let _415_v7;
            _415_v7 = _dafny.Seq.of(_409_v2, _409_v2);
            let _416_v8;
            _416_v8 = _dafny.MultiSet.fromElements((_414_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_415_v7)[_module.__default.safeIndex(_409_v2, new BigNumber((_415_v7).length))],_409_v2)).length), new BigNumber((_414_v6).length))]);
            _416_v8 = _416_v8;
            (globalState).f1 = (_409_v2).multipliedBy(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-537)), ((_417_v2) => function (_418_i1) {
              return _417_v2;
            })(_409_v2))).length)), _409_v2));
            let _419_v9;
            let _nw80 = new _module.C0();
            _nw80.__ctor(_407_v0);
            _419_v9 = _nw80;
            let _420_v10;
            let _nw81 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _420_v10 = _nw81;
            let _421_v11;
            _421_v11 = _dafny.Seq.of(p0, p0, false, p0);
            let _index47 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_420_v10).length));
            (_420_v10)[_index47] = new BigNumber((_421_v11).length);
            let _index48 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_420_v10).length));
            (_420_v10)[_index48] = new BigNumber((_421_v11).length);
          }
        }
      }
      let _422_v12;
      let _init15 = ((_423_p0) => function (_424_i2) {
        return _423_p0;
      })(p0);
      let _nw82 = Array((new BigNumber(29)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw82.length); _i0_15++) {
        _nw82[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _422_v12 = _nw82;
      let _425_v13;
      _425_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm10(globalState)).length),_422_v12);
      _425_v13 = (_425_v13).update(_409_v2, ((p0) ? (_422_v12) : (_422_v12)));
      let _426_v14;
      _426_v14 = new _dafny.CodePoint('x'.codePointAt(0));
      let _427_v15;
      _427_v15 = _module.D0.create_DC1(_426_v14);
      let _428_v16;
      _428_v16 = _dafny.Seq.of(_module.__default.fm1(new BigNumber((_dafny.MultiSet.fromElements(_module.D0.create_DC1(_426_v14), _427_v15)).cardinality()), globalState));
      if (_dafny.Seq.contains(_428_v16, (new BigNumber((_dafny.Seq.UnicodeFromString("bbijgu")).length)).multipliedBy(new BigNumber((p2).length)))) {
        let _429_v17;
        _429_v17 = _dafny.Seq.of(p0, p0, false);
        _429_v17 = _dafny.Seq.Concat(_429_v17, ((!(p0)) ? (_429_v17) : (_429_v17)));
        let _430_v18;
        _430_v18 = _module.D0.create_DC2(p0);
        (globalState).f0 = (_430_v18).dtor_cf2;
        let _source4 = _module.D0.create_DC3(_module.__default.fm0(new BigNumber((_428_v16).length), globalState));
        if (_source4.is_DC1) {
          let _431___mcc_h0 = (_source4).cf1;
          let _432_cf1 = _431___mcc_h0;
          let _433_v19;
          let _init16 = ((_434_v2) => function (_435_i3) {
            return (_435_i3).plus(_434_v2);
          })(_409_v2);
          let _nw83 = Array((new BigNumber(5)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw83.length); _i0_16++) {
            _nw83[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _433_v19 = _nw83;
          let _index49 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_433_v19).length));
          (_433_v19)[_index49] = (_dafny.ZERO).minus((_409_v2).minus(new BigNumber(-624)));
          let _436_v20;
          _436_v20 = _dafny.Seq.of(_433_v19);
          let _index50 = _module.__default.safeIndex(new BigNumber(656), new BigNumber((_433_v19).length));
          (_433_v19)[_index50] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_433_v19), _dafny.Seq.Concat(_436_v20, _436_v20))).length);
          let _437_v21;
          _437_v21 = _dafny.Map.Empty.slice().updateUnsafe(_410_v3,p0);
          _437_v21 = _dafny.Map.Empty.slice().updateUnsafe(_410_v3,p0);
          r0 = _module.__default.safeModuloInt((_433_v19)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_433_v19).length))], _409_v2);
          let _index51 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_422_v12).length));
          (_422_v12)[_index51] = p0;
          let _index52 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_422_v12).length));
          let _rhs57 = (_this).fm3(!((_433_v19)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_433_v19).length))]).isEqualTo((_433_v19)[_module.__default.safeIndex(new BigNumber(656), new BigNumber((_433_v19).length))]), globalState);
          let _rhs58 = p0;
          let _rhs59 = p0;
          let _lhs34 = globalState;
          let _lhs35 = _422_v12;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_422_v12).length));
          _432_cf1 = _rhs57;
          _lhs34.f0 = _rhs58;
          _lhs35[_lhs36] = _rhs59;
        } else if (_source4.is_DC2) {
          let _438___mcc_h1 = (_source4).cf2;
          let _439_cf2 = _438___mcc_h1;
          let _440_v22;
          _440_v22 = _dafny.Seq.of(_module.__default.fm11(p0, p0, _module.D0.create_DC3(p0), globalState), _429_v17, _429_v17);
          let _441_v23;
          _441_v23 = _dafny.MultiSet.fromElements(new BigNumber((_440_v22).length));
          let _442_v24;
          _442_v24 = _dafny.Map.Empty.slice().updateUnsafe(_409_v2,new BigNumber((_441_v23).cardinality()));
          _442_v24 = (_442_v24).update(((_dafny.ZERO).minus(_409_v2)).plus(_409_v2), _module.__default.fm1(_409_v2, globalState));
          let _443_v25;
          let _nw84 = new _module.C0();
          _nw84.__ctor((_408_v1).f6);
          _443_v25 = _nw84;
          _426_v14 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs60 = _409_v2;
          let _rhs61 = (_409_v2).minus((_409_v2).multipliedBy(_409_v2));
          let _lhs37 = globalState;
          let _lhs38 = globalState;
          _lhs37.f3 = _rhs60;
          _lhs38.f1 = _rhs61;
        } else if (_source4.is_DC3) {
          let _444___mcc_h2 = (_source4).cf3;
          let _445_cf3 = _444___mcc_h2;
          (globalState).f0 = p0;
          let _446_v26;
          _446_v26 = _dafny.Set.fromElements(_408_v1);
          let _447_v27;
          _447_v27 = _dafny.Set.fromElements(_446_v26);
          let _448_v28;
          _448_v28 = _dafny.MultiSet.fromElements(_445_cf3);
          let _449_v29;
          _449_v29 = _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_448_v28);
          let _450_v30;
          _450_v30 = _module.D1.create_DC4(_448_v28);
          let _451_v31;
          _451_v31 = _module.D2.create_DC7(_408_v1);
          let _452_v32;
          _452_v32 = _dafny.Map.Empty.slice().updateUnsafe(_426_v14,(_451_v31).dtor_cf7);
          let _453_v33;
          let _nw85 = Array((new BigNumber(27)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = _409_v2;
          _nw85[(_dafny.ONE).toNumber()] = (_409_v2).plus((_module.D0.create_DC0(new BigNumber((_dafny.Seq.UnicodeFromString("reo")).length))).dtor_cf0);
          _nw85[(new BigNumber(2)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(3)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(4)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(5)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(6)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(7)).toNumber()] = new BigNumber(((_447_v27).Intersect(_447_v27)).length);
          _nw85[(new BigNumber(8)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(9)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_409_v2), _409_v2);
          _nw85[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_428_v16).length), _409_v2);
          _nw85[(new BigNumber(12)).toNumber()] = ((_428_v16)[_module.__default.safeIndex(_409_v2, new BigNumber((_428_v16).length))]).minus(_409_v2);
          _nw85[(new BigNumber(13)).toNumber()] = _module.__default.safeModuloInt(_409_v2, (((_448_v28).contains(_445_cf3)) ? ((_448_v28).get(_445_cf3)) : (_409_v2)));
          _nw85[(new BigNumber(14)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(15)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_449_v29).contains(_408_v1)) ? ((_449_v29).get(_408_v1)) : ((_450_v30).dtor_cf4)),_422_v12)).length);
          _nw85[(new BigNumber(17)).toNumber()] = _module.__default.fm1(_409_v2, globalState);
          _nw85[(new BigNumber(18)).toNumber()] = _module.__default.safeDivisionInt(_409_v2, _409_v2);
          _nw85[(new BigNumber(19)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(20)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(21)).toNumber()] = new BigNumber(((_452_v32).Merge(_452_v32)).length);
          _nw85[(new BigNumber(22)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(23)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(24)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(25)).toNumber()] = _409_v2;
          _nw85[(new BigNumber(26)).toNumber()] = new BigNumber(814);
          _453_v33 = _nw85;
          let _index53 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_453_v33).length));
          (_453_v33)[_index53] = _409_v2;
          let _index54 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_453_v33).length));
          (_453_v33)[_index54] = _409_v2;
          let _454_v34;
          _454_v34 = _dafny.Seq.UnicodeFromString("hkcpig");
          _454_v34 = (((_409_v2).isLessThanOrEqualTo((_453_v33)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_453_v33).length))])) ? (_454_v34) : (_454_v34));
          let _455_v35;
          _455_v35 = _dafny.MultiSet.fromElements(_428_v16, _dafny.Seq.update(_module.__default.fm10(globalState), _module.__default.safeIndex((_453_v33)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_453_v33).length))], new BigNumber((_module.__default.fm10(globalState)).length)), _409_v2), _428_v16);
          _455_v35 = ((_455_v35).Difference(_455_v35)).Intersect(((p0) ? (_455_v35) : (_dafny.MultiSet.fromElements(_dafny.Seq.update(_428_v16, _module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_428_v16).length)), new BigNumber((_dafny.Seq.of(_445_cf3)).length)), _dafny.Seq.of((_453_v33)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_453_v33).length))], _409_v2), _428_v16))));
        } else {
          let _456___mcc_h3 = (_source4).cf0;
          let _457_cf0 = _456___mcc_h3;
          let _458_v36;
          _458_v36 = _dafny.Set.fromElements(_457_cf0, _409_v2);
          let _459_v37;
          _459_v37 = _dafny.Map.Empty.slice().updateUnsafe(_457_cf0,new BigNumber((_458_v36).length));
          let _460_v38;
          _460_v38 = _dafny.Seq.of(_459_v37);
          let _rhs62 = _module.__default.safeModuloInt(new BigNumber(903), _409_v2);
          let _rhs63 = _dafny.Seq.Concat(_dafny.Seq.of(_408_v1), _411_v4);
          let _rhs64 = (_module.D3.create_DC9(_460_v38)).dtor_cf10;
          let _rhs65 = _422_v12;
          let _lhs39 = globalState;
          _lhs39.f1 = _rhs62;
          _411_v4 = _rhs63;
          _460_v38 = _rhs64;
          _422_v12 = _rhs65;
          let _461_v39;
          let _nw86 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _461_v39 = _nw86;
          let _index55 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_461_v39).length));
          (_461_v39)[_index55] = _426_v14;
          let _462_v40;
          _462_v40 = _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_426_v14);
          let _463_v41;
          _463_v41 = _module.D2.create_DC7(_408_v1);
          let _index56 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_461_v39).length));
          (_461_v39)[_index56] = (((_462_v40).contains((_463_v41).dtor_cf7)) ? ((_462_v40).get((_463_v41).dtor_cf7)) : (_426_v14));
          (globalState).f1 = _409_v2;
          let _464_v42;
          _464_v42 = _dafny.MultiSet.fromElements(p0, p0, !(p0));
          let _465_v43;
          _465_v43 = _module.D1.create_DC4(_dafny.MultiSet.FromArray(_429_v17));
          let _466_v44;
          let _nw87 = Array((new BigNumber(22)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = (_464_v42).Difference(_dafny.MultiSet.fromElements(false));
          _nw87[(_dafny.ONE).toNumber()] = (_464_v42).Union(_dafny.MultiSet.FromArray(_429_v17));
          _nw87[(new BigNumber(2)).toNumber()] = ((p0) ? (_464_v42) : (_464_v42));
          _nw87[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(p0, p0, _module.__default.fm0(_409_v2, globalState));
          _nw87[(new BigNumber(4)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(5)).toNumber()] = (_465_v43).dtor_cf4;
          _nw87[(new BigNumber(6)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(p0);
          _nw87[(new BigNumber(8)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(9)).toNumber()] = (_dafny.MultiSet.FromArray(_429_v17)).Intersect(_464_v42);
          _nw87[(new BigNumber(10)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(11)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(p0, p0);
          _nw87[(new BigNumber(13)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(14)).toNumber()] = (_464_v42).Intersect(_464_v42);
          _nw87[(new BigNumber(15)).toNumber()] = (_464_v42).Union(_464_v42);
          _nw87[(new BigNumber(16)).toNumber()] = (_464_v42).Difference(_dafny.MultiSet.fromElements(true));
          _nw87[(new BigNumber(17)).toNumber()] = _module.__default.fm12(_409_v2, globalState);
          _nw87[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(true);
          _nw87[(new BigNumber(19)).toNumber()] = _464_v42;
          _nw87[(new BigNumber(20)).toNumber()] = (_464_v42).Difference(_464_v42);
          _nw87[(new BigNumber(21)).toNumber()] = (_dafny.MultiSet.fromElements(false, !(p0))).Union(_464_v42);
          _466_v44 = _nw87;
          let _467_v45;
          _467_v45 = _dafny.Map.Empty.slice().updateUnsafe(_408_v1,_464_v42);
          let _index57 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_466_v44).length));
          (_466_v44)[_index57] = (((_467_v45).contains(_408_v1)) ? ((_467_v45).get(_408_v1)) : (_464_v42));
          let _468_v46;
          _468_v46 = _module.D0.create_DC0(_457_cf0);
          let _pat_let_tv9 = _457_cf0;
          let _469_v47;
          let _nw88 = Array((new BigNumber(11)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = _module.D0.create_DC0(_409_v2);
          _nw88[(_dafny.ONE).toNumber()] = _468_v46;
          _nw88[(new BigNumber(2)).toNumber()] = function (_pat_let4_0) {
            return function (_470_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_471_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_471_dt__update_hcf0_h0);
                }(_pat_let5_0);
              }(_pat_let_tv9);
            }(_pat_let4_0);
          }(_468_v46);
          _nw88[(new BigNumber(3)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(4)).toNumber()] = _module.__default.fm13(false, p0, _dafny.Map.Empty.slice().updateUnsafe(_409_v2,_457_cf0), p0, globalState);
          _nw88[(new BigNumber(5)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(6)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(7)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(8)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(9)).toNumber()] = _468_v46;
          _nw88[(new BigNumber(10)).toNumber()] = _module.D0.create_DC0(_409_v2);
          _469_v47 = _nw88;
          let _index58 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_469_v47).length));
          (_469_v47)[_index58] = _468_v46;
          let _472_v48;
          _472_v48 = _module.D0.create_DC3((_457_cf0).isLessThan(_409_v2));
          let _473_v49;
          _473_v49 = _dafny.Map.Empty.slice().updateUnsafe(_409_v2,_dafny.MultiSet.fromElements(!(p0), p0));
          let _index59 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_466_v44).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_469_v47).length));
          let _rhs66 = (((_473_v49).contains(_409_v2)) ? ((_473_v49).get(_409_v2)) : ((_464_v42).Difference(_464_v42)));
          let _rhs67 = _457_cf0;
          let _rhs68 = _468_v46;
          let _rhs69 = _module.D0.create_DC3(p0);
          let _rhs70 = _module.__default.fm1(_409_v2, globalState);
          let _lhs40 = _466_v44;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_466_v44).length));
          let _lhs42 = globalState;
          let _lhs43 = _469_v47;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_469_v47).length));
          _lhs40[_lhs41] = _rhs66;
          _lhs42.f3 = _rhs67;
          _lhs43[_lhs44] = _rhs68;
          _472_v48 = _rhs69;
          r0 = _rhs70;
        }
        if ((_409_v2).isLessThanOrEqualTo(new BigNumber((_429_v17).length))) {
          let _index61 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_422_v12).length));
          (_422_v12)[_index61] = p0;
          let _index62 = _module.__default.safeIndex(new BigNumber(788), new BigNumber((_422_v12).length));
          (_422_v12)[_index62] = p0;
          let _474_v50;
          let _init17 = ((_475_v2) => function (_476_i4) {
            return _module.__default.safeDivisionInt(_476_i4, _475_v2);
          })(_409_v2);
          let _nw89 = Array((new BigNumber(3)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw89.length); _i0_17++) {
            _nw89[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _474_v50 = _nw89;
          let _index63 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_474_v50).length));
          (_474_v50)[_index63] = _409_v2;
          let _index64 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_474_v50).length));
          (_474_v50)[_index64] = (_409_v2).minus(_409_v2);
          (globalState).f0 = !(true);
          _407_v0 = (_408_v1).f6;
          (globalState).f0 = p0;
        } else {
          (globalState).f0 = p0;
          (globalState).f0 = _dafny.areEqual(_dafny.Seq.Concat(_429_v17, _429_v17), _429_v17);
          let _index65 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length));
          (_422_v12)[_index65] = p0;
          let _477_v51;
          _477_v51 = _module.D4.create_DC13(_422_v12);
          let _pat_let_tv10 = _477_v51;
          let _pat_let_tv11 = _422_v12;
          let _index66 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length));
          let _rhs71 = (function (_pat_let6_0) {
            return function (_478_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_481_dt__update_hcf19_h0) {
                  return _module.D4.create_DC13(_481_dt__update_hcf19_h0);
                }(_pat_let7_0);
              }((function (_pat_let8_0) {
                return function (_479_dt__update__tmp_h2) {
                  return function (_pat_let9_0) {
                    return function (_480_dt__update_hcf19_h1) {
                      return _module.D4.create_DC13(_480_dt__update_hcf19_h1);
                    }(_pat_let9_0);
                  }(_pat_let_tv11);
                }(_pat_let8_0);
              }(_pat_let_tv10)).dtor_cf19);
            }(_pat_let6_0);
          }(_477_v51)).dtor_cf19;
          let _rhs72 = (_409_v2).isLessThan(_409_v2);
          let _rhs73 = p0;
          let _lhs45 = _422_v12;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length));
          let _lhs47 = globalState;
          _422_v12 = _rhs71;
          _lhs45[_lhs46] = _rhs72;
          _lhs47.f0 = _rhs73;
          let _index67 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length));
          let _rhs74 = _407_v0;
          let _rhs75 = (_422_v12)[_module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length))];
          let _lhs48 = _422_v12;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_422_v12).length));
          _407_v0 = _rhs74;
          _lhs48[_lhs49] = _rhs75;
          (globalState).f0 = p0;
        }
        let _482_v52;
        _482_v52 = _dafny.Seq.of(_426_v14);
        _482_v52 = _dafny.Seq.of(_426_v14);
      } else {
        (globalState).f1 = ((_dafny.ZERO).minus(_409_v2)).minus(_409_v2);
        (globalState).f0 = p0;
        let _483_v53;
        _483_v53 = _dafny.Map.Empty.slice().updateUnsafe(_409_v2,new BigNumber(547));
        r0 = (((_483_v53).contains(_409_v2)) ? ((_483_v53).get(_409_v2)) : (_409_v2));
        let _484_v54;
        _484_v54 = _dafny.Seq.of(p0);
        let _485_v55;
        _485_v55 = _dafny.Map.Empty.slice().updateUnsafe(_484_v54,_module.__default.fm4(p0, _409_v2, _409_v2, _426_v14, globalState));
        _485_v55 = (_485_v55).update(_module.__default.fm11(p0, p0, p3, globalState), _dafny.Seq.Concat(p2, p2));
        let _index68 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_422_v12).length));
        (_422_v12)[_index68] = true;
        let _index69 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_422_v12).length));
        (_422_v12)[_index69] = !(p0);
      }
      let _486_v56;
      _486_v56 = _module.D0.create_DC2(p0);
      let _487_v57;
      _487_v57 = _dafny.Seq.of(_486_v56, _486_v56);
      let _488_v58;
      let _nw90 = Array((new BigNumber(10)).toNumber());
      _nw90[(_dafny.ZERO).toNumber()] = _487_v57;
      _nw90[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_module.__default.fm14(p0, globalState), _dafny.Seq.of(_486_v56));
      _nw90[(new BigNumber(2)).toNumber()] = _487_v57;
      _nw90[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_487_v57, _487_v57);
      _nw90[(new BigNumber(4)).toNumber()] = _module.__default.fm14(p0, globalState);
      _nw90[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(69)), ((_489_v56) => function (_490_i5) {
        return _489_v56;
      })(_486_v56)), _487_v57);
      _nw90[(new BigNumber(6)).toNumber()] = _487_v57;
      _nw90[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_487_v57, _487_v57);
      _nw90[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_486_v56, _486_v56);
      _nw90[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_487_v57, _487_v57);
      _488_v58 = _nw90;
      let _index70 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_488_v58).length));
      (_488_v58)[_index70] = _487_v57;
      let _index71 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
      (_422_v12)[_index71] = (new BigNumber((_module.__default.fm4(p0, _409_v2, _409_v2, _426_v14, globalState)).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(_409_v2));
      let _491_v59;
      _491_v59 = _dafny.Map.Empty.slice().updateUnsafe(p0,_487_v57);
      let _index72 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_488_v58).length));
      let _index73 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
      let _rhs76 = _module.__default.safeModuloInt(_409_v2, _409_v2);
      let _rhs77 = (((_491_v59).contains(!(p0))) ? ((_491_v59).get(!(p0))) : (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), ((_492_v56) => function (_493_i6) {
        return _492_v56;
      })(_486_v56)), _dafny.Seq.of(_486_v56))));
      let _rhs78 = p0;
      let _rhs79 = (p0) || (_dafny.Seq.IsProperPrefixOf(p2, _dafny.Seq.UnicodeFromString("y")));
      let _lhs50 = globalState;
      let _lhs51 = _488_v58;
      let _lhs52 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_488_v58).length));
      let _lhs53 = _422_v12;
      let _lhs54 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
      let _lhs55 = globalState;
      _lhs50.f3 = _rhs76;
      _lhs51[_lhs52] = _rhs77;
      _lhs53[_lhs54] = _rhs78;
      _lhs55.f0 = _rhs79;
      if ((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) {
        let _494_v60;
        _494_v60 = _dafny.MultiSet.fromElements(p0);
        let _pat_let_tv12 = _494_v60;
        let _pat_let_tv13 = _409_v2;
        _486_v56 = function (_pat_let10_0) {
          return function (_495_dt__update__tmp_h3) {
            return function (_pat_let11_0) {
              return function (_496_dt__update_hcf2_h0) {
                return _module.D0.create_DC2(_496_dt__update_hcf2_h0);
              }(_pat_let11_0);
            }(!((new BigNumber((_pat_let_tv12).cardinality())).isLessThanOrEqualTo(_pat_let_tv13)));
          }(_pat_let10_0);
        }(_486_v56);
        let _497_v61;
        let _init18 = ((_498_p0) => function (_499_i7) {
          return _dafny.Seq.of(_498_p0, _498_p0);
        })(p0);
        let _nw91 = Array((new BigNumber(9)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw91.length); _i0_18++) {
          _nw91[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _497_v61 = _nw91;
        let _500_v62;
        _500_v62 = _dafny.Map.Empty.slice().updateUnsafe(_409_v2,_497_v61);
        _497_v61 = (((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) ? (_497_v61) : ((((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) ? ((((_500_v62).contains(_409_v2)) ? ((_500_v62).get(_409_v2)) : (_497_v61))) : (_497_v61))));
        if ((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) {
          (globalState).f1 = new BigNumber((_dafny.Seq.of(_dafny.Seq.Concat(p2, p2))).length);
          let _501_v63;
          _501_v63 = _module.D2.create_DC7(_408_v1);
          let _rhs80 = _408_v1;
          let _rhs81 = _501_v63;
          _408_v1 = _rhs80;
          _501_v63 = _rhs81;
          let _502_v64;
          _502_v64 = _dafny.Seq.of(true, p0);
          let _503_v65;
          _503_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_502_v64).length),(_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]);
          let _504_v66;
          let _nw92 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _504_v66 = _nw92;
          let _505_v67;
          _505_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_503_v65).length),_504_v66);
          (globalState).f0 = (new BigNumber((_505_v67).length)).isLessThan(new BigNumber(993));
          let _506_v68;
          _506_v68 = _module.D5.create_DC16(_dafny.Seq.UnicodeFromString("xhrajd"));
          (globalState).f3 = ((((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) ? (_409_v2) : (new BigNumber(((_506_v68).dtor_cf26).length)))).minus(_409_v2);
          let _507_v69;
          let _nw93 = new _module.C0();
          _nw93.__ctor((_408_v1).f6);
          _507_v69 = _nw93;
        } else {
          (globalState).f3 = _module.__default.safeDivisionInt(_409_v2, new BigNumber((_module.__default.fm4((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], new BigNumber((_494_v60).cardinality()), _409_v2, _426_v14, globalState)).length));
          let _508_v70;
          _508_v70 = _dafny.Seq.UnicodeFromString("fcqa");
          let _509_v71;
          _509_v71 = _module.D5.create_DC18(new BigNumber((_dafny.Seq.UnicodeFromString("vcxcktchu")).length), _409_v2);
          let _510_v72;
          _510_v72 = _dafny.Seq.of(_module.D5.create_DC18(_409_v2, _409_v2), _509_v71, _509_v71, _module.D5.create_DC18(_409_v2, new BigNumber(-180)), _509_v71);
          let _511_v73;
          _511_v73 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm1(_module.__default.fm1(_409_v2, globalState), globalState)).minus(_409_v2),_510_v72);
          let _512_v74;
          _512_v74 = _dafny.Seq.of(p0, p0);
          let _513_v75;
          _513_v75 = _dafny.MultiSet.fromElements(_409_v2);
          let _index74 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          let _rhs82 = (_408_v1).fm7(new _dafny.CodePoint('j'.codePointAt(0)), globalState);
          let _rhs83 = p2;
          let _rhs84 = new BigNumber(174);
          let _rhs85 = _module.__default.fm15(((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]) === ((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]), _dafny.Seq.Concat(_dafny.Seq.update(_512_v74, _module.__default.safeIndex(new BigNumber((_513_v75).cardinality()), new BigNumber((_512_v74).length)), true), _dafny.Seq.of(p0, p0, p0, (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], !(p0))), _dafny.Seq.Concat(_512_v74, _512_v74), globalState);
          let _rhs86 = p0;
          let _lhs56 = globalState;
          let _lhs57 = globalState;
          let _lhs58 = _422_v12;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          _lhs56.f3 = _rhs82;
          _508_v70 = _rhs83;
          _lhs57.f3 = _rhs84;
          _511_v73 = _rhs85;
          _lhs58[_lhs59] = _rhs86;
          _497_v61 = _497_v61;
          let _514_v76;
          _514_v76 = _dafny.Map.Empty.slice().updateUnsafe(_409_v2,(_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]);
          let _515_v77;
          _515_v77 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_514_v76).length));
          let _516_v78;
          let _nw94 = Array((new BigNumber(19)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = new BigNumber((_515_v77).length);
          _nw94[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(_409_v2);
          _nw94[(new BigNumber(2)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(3)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(p0)).length);
          _nw94[(new BigNumber(5)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(6)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(7)).toNumber()] = new BigNumber((_513_v75).cardinality());
          _nw94[(new BigNumber(8)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_409_v2);
          _nw94[(new BigNumber(10)).toNumber()] = _module.__default.fm1(new BigNumber((_module.__default.fm10(globalState)).length), globalState);
          _nw94[(new BigNumber(11)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(12)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(13)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(14)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(15)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(16)).toNumber()] = _409_v2;
          _nw94[(new BigNumber(17)).toNumber()] = new BigNumber(655);
          _nw94[(new BigNumber(18)).toNumber()] = _409_v2;
          _516_v78 = _nw94;
          let _517_v79;
          _517_v79 = _dafny.Seq.of(_516_v78);
          let _518_v80;
          _518_v80 = _dafny.Set.fromElements(_dafny.Set.fromElements((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], !(p0)));
          let _519_v81;
          _519_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_518_v80).length),_516_v78);
          let _520_v82;
          let _nw95 = Array((new BigNumber(28)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = _517_v79;
          _nw95[(_dafny.ONE).toNumber()] = _517_v79;
          _nw95[(new BigNumber(2)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(3)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(4)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(5)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(6)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_516_v78);
          _nw95[(new BigNumber(8)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(9)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(10)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(11)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(12)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(13)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(14)).toNumber()] = _dafny.Seq.of(_516_v78, _516_v78);
          _nw95[(new BigNumber(15)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(16)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_517_v79, _dafny.Seq.of(_516_v78, _516_v78));
          _nw95[(new BigNumber(18)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(19)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_517_v79, _517_v79), _module.__default.safeIndex(new BigNumber((_module.__default.fm16(globalState)).length), new BigNumber((_dafny.Seq.Concat(_517_v79, _517_v79)).length)), (((_519_v81).contains(new BigNumber(992))) ? ((_519_v81).get(new BigNumber(992))) : (_516_v78)));
          _nw95[(new BigNumber(21)).toNumber()] = _dafny.Seq.update(_517_v79, _module.__default.safeIndex(_409_v2, new BigNumber((_517_v79).length)), _516_v78);
          _nw95[(new BigNumber(22)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_517_v79, _517_v79), _module.__default.safeIndex(new BigNumber((_512_v74).length), new BigNumber((_dafny.Seq.Concat(_517_v79, _517_v79)).length)), _516_v78);
          _nw95[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(_516_v78, _516_v78);
          _nw95[(new BigNumber(24)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(25)).toNumber()] = _517_v79;
          _nw95[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_517_v79, _517_v79);
          _nw95[(new BigNumber(27)).toNumber()] = _517_v79;
          _520_v82 = _nw95;
          let _521_v83;
          _521_v83 = _module.D6.create_DC19(_517_v79);
          let _index75 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_520_v82).length));
          (_520_v82)[_index75] = (_521_v83).dtor_cf34;
          let _index76 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_520_v82).length));
          (_520_v82)[_index76] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_517_v79, _517_v79), _517_v79), _module.__default.safeIndex(_409_v2, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_517_v79, _517_v79), _517_v79)).length)), _516_v78);
          _515_v77 = (_515_v77).update(!(p0), _module.__default.safeModuloInt(_409_v2, _409_v2));
        }
        let _522_v84;
        _522_v84 = _dafny.Seq.of(p0, (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]);
        let _523_v85;
        _523_v85 = _module.D6.create_DC20(_522_v84, p0);
        let _source5 = _523_v85;
        if (_source5.is_DC20) {
          let _524___mcc_h4 = (_source5).cf35;
          let _525___mcc_h5 = (_source5).cf36;
          let _526_cf36 = _525___mcc_h5;
          let _527_cf35 = _524___mcc_h4;
          let _528_v86;
          _528_v86 = _module.D6.create_DC22(((_526_cf36) ? (new BigNumber((_dafny.Seq.UnicodeFromString("nock")).length)) : (_409_v2)), _module.__default.fm1(new BigNumber((p2).length), globalState));
          let _529_v87;
          _529_v87 = _dafny.Set.fromElements(_408_v1);
          let _rhs87 = new BigNumber((_529_v87).length);
          let _rhs88 = (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))];
          let _rhs89 = _module.D6.create_DC22(_module.__default.safeDivisionInt(_409_v2, new BigNumber(-337)), (_dafny.ZERO).minus((new BigNumber((_425_v13).length)).plus(_409_v2)));
          _409_v2 = _rhs87;
          _526_cf36 = _rhs88;
          _528_v86 = _rhs89;
          let _index77 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          (_422_v12)[_index77] = !((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]);
          _409_v2 = _409_v2;
          (globalState).f1 = _409_v2;
        } else if (_source5.is_DC21) {
          let _530___mcc_h6 = (_source5).cf37;
          let _531___mcc_h7 = (_source5).cf38;
          let _532___mcc_h8 = (_source5).cf39;
          let _533___mcc_h9 = (_source5).cf40;
          let _534_cf40 = _533___mcc_h9;
          let _535_cf39 = _532___mcc_h8;
          let _536_cf38 = _531___mcc_h7;
          let _537_cf37 = _530___mcc_h6;
          let _index78 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          (_422_v12)[_index78] = _534_cf40;
          let _index79 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          (_422_v12)[_index79] = _534_cf40;
          _534_cf40 = p0;
          let _538_v88;
          _538_v88 = _dafny.MultiSet.fromElements(_535_cf39, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("wa"), _module.__default.safeIndex(_409_v2, new BigNumber((_dafny.Seq.UnicodeFromString("wa")).length)), new _dafny.CodePoint('o'.codePointAt(0))));
          let _index80 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          (_422_v12)[_index80] = ((_538_v88).Union(_538_v88)).IsSubsetOf(_538_v88);
        } else if (_source5.is_DC22) {
          let _539___mcc_h10 = (_source5).cf41;
          let _540___mcc_h11 = (_source5).cf42;
          let _541_cf42 = _540___mcc_h11;
          let _542_cf41 = _539___mcc_h10;
          let _543_v89;
          _543_v89 = _dafny.Set.fromElements(_541_cf42);
          let _544_v90;
          let _init19 = ((_545_v2) => function (_546_i8) {
            return (_546_i8).plus(_545_v2);
          })(_409_v2);
          let _nw96 = Array((new BigNumber(15)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw96.length); _i0_19++) {
            _nw96[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _544_v90 = _nw96;
          let _547_v91;
          _547_v91 = _dafny.Set.fromElements(_544_v90, _544_v90, _544_v90);
          let _548_v92;
          _548_v92 = _dafny.Map.Empty.slice().updateUnsafe(_543_v89,new BigNumber((_547_v91).length));
          _548_v92 = (_548_v92).update((_543_v89).Intersect(_543_v89), _409_v2);
          let _549_v93;
          let _nw97 = Array((new BigNumber(26)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _410_v3;
          _nw97[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_408_v1,_409_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_408_v1,_409_v2));
          _nw97[(new BigNumber(2)).toNumber()] = (_410_v3).update(_408_v1, _409_v2);
          _nw97[(new BigNumber(3)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(4)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(5)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(6)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(7)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(8)).toNumber()] = (_410_v3).Merge(_410_v3);
          _nw97[(new BigNumber(9)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(10)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(11)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(12)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_411_v4)[_module.__default.safeIndex(_409_v2, new BigNumber((_411_v4).length))],(_dafny.ZERO).minus(new BigNumber((_428_v16).length)));
          _nw97[(new BigNumber(14)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_408_v1,_541_cf42)).Merge((_410_v3).update(_408_v1, _541_cf42));
          _nw97[(new BigNumber(15)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(16)).toNumber()] = (_410_v3).update(_408_v1, _542_cf41);
          _nw97[(new BigNumber(17)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(18)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(19)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(20)).toNumber()] = ((_410_v3).update(_408_v1, _module.__default.fm1(_409_v2, globalState))).Merge(_410_v3);
          _nw97[(new BigNumber(21)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_408_v1,new BigNumber((p1).length));
          _nw97[(new BigNumber(23)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(24)).toNumber()] = _410_v3;
          _nw97[(new BigNumber(25)).toNumber()] = (_410_v3).Merge(_410_v3);
          _549_v93 = _nw97;
          _549_v93 = _549_v93;
          let _550_v94;
          let _nw98 = new _module.C0();
          _nw98.__ctor(((false) ? ((_408_v1).f6) : ((_408_v1).f6)));
          _550_v94 = _nw98;
          let _index81 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_544_v90).length));
          (_544_v90)[_index81] = (new BigNumber(-10)).multipliedBy(_542_cf41);
          let _index82 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_544_v90).length));
          (_544_v90)[_index82] = (_dafny.ZERO).minus((((_522_v84)[_module.__default.safeIndex(new BigNumber(-978), new BigNumber((_522_v84).length))]) ? (_409_v2) : (_409_v2)));
        } else if (_source5.is_DC19) {
          let _551___mcc_h12 = (_source5).cf34;
          let _552_cf34 = _551___mcc_h12;
          let _index83 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
          (_422_v12)[_index83] = p0;
          _494_v60 = _494_v60;
          let _553_v95;
          _553_v95 = _dafny.Set.fromElements(p1);
          (globalState).f1 = ((_409_v2).minus(_409_v2)).multipliedBy((_dafny.ZERO).minus((_409_v2).plus(new BigNumber((_553_v95).length))));
          let _554_v96;
          _554_v96 = _module.D2.create_DC8((_dafny.ZERO).minus(_409_v2), new BigNumber((_522_v84).length));
          let _555_v97;
          _555_v97 = _dafny.Seq.of(_module.__default.fm17(globalState), _module.D2.create_DC8(_409_v2, new BigNumber((p2).length)), _554_v96);
          let _556_v98;
          _556_v98 = _dafny.Set.fromElements(_408_v1);
          let _557_v99;
          _557_v99 = _dafny.Seq.of(_dafny.Seq.update(_555_v97, _module.__default.safeIndex(_409_v2, new BigNumber((_555_v97).length)), _554_v96), _dafny.Seq.of(_554_v96, _module.D2.create_DC8((_dafny.ZERO).minus(new BigNumber((_556_v98).length)), _409_v2), _554_v96), _dafny.Seq.Concat(_dafny.Seq.of(_554_v96), _555_v97), _dafny.Seq.of(_554_v96, _554_v96, _554_v96));
          _557_v99 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(297)), ((_558_v97) => function (_559_i9) {
            return _dafny.Seq.Concat(_558_v97, _558_v97);
          })(_555_v97));
        } else {
          let _560___mcc_h13 = (_source5).cf43;
          let _561_cf43 = _560___mcc_h13;
          (globalState).f3 = _409_v2;
          let _562_v100;
          let _nw99 = new _module.C0();
          _nw99.__ctor(((false) ? ((_408_v1).f6) : ((_408_v1).f6)));
          _562_v100 = _nw99;
          let _index84 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_497_v61).length));
          (_497_v61)[_index84] = _522_v84;
          let _index85 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_497_v61).length));
          (_497_v61)[_index85] = _dafny.Seq.Concat(_522_v84, _dafny.Seq.update(_module.__default.fm11(p0, (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], p3, globalState), _module.__default.safeIndex(_409_v2, new BigNumber((_module.__default.fm11(p0, (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], p3, globalState)).length)), true));
          (globalState).f1 = _409_v2;
        }
        let _563_v101;
        let _564_v102;
        let _565_v103;
        let _out29;
        let _out30;
        let _out31;
        let _outcollector8 = (_this).m0(globalState);
        _out29 = _outcollector8[0];
        _out30 = _outcollector8[1];
        _out31 = _outcollector8[2];
        _563_v101 = _out29;
        _564_v102 = _out30;
        _565_v103 = _out31;
      } else {
        let _566_v104;
        _566_v104 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_426_v14);
        _566_v104 = (_566_v104).update((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))], _426_v14);
        _426_v14 = _426_v14;
        let _567_v105;
        _567_v105 = _dafny.MultiSet.fromElements(!((_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]), (_422_v12)[_module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length))]);
        _567_v105 = _567_v105;
        let _index86 = _module.__default.safeIndex(new BigNumber(231), new BigNumber((_422_v12).length));
        (_422_v12)[_index86] = _module.__default.fm0(_409_v2, globalState);
        let _568_v106;
        let _nw100 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Seq.of());
        _568_v106 = _nw100;
        _568_v106 = _568_v106;
      }
      let _569_v107;
      _569_v107 = _module.D3.create_DC11(_409_v2);
      r0 = (_409_v2).plus((_569_v107).dtor_cf14);
      return r0;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
