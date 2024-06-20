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
      return new BigNumber(((_dafny.Set.fromElements(!(false))).Union(_dafny.Set.fromElements(false))).length);
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(!(false))).Intersect(((false) ? (_dafny.MultiSet.FromArray(_dafny.Seq.of(false))) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true))))));
    };
    static fm2(p0, p1, globalState) {
      return !(true);
    };
    static fm3(globalState) {
      return (_dafny.Set.fromElements(true, false, false)).Difference((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(false, !(false))));
    };
    static fm4(p0, p1, p2, globalState) {
      return (((true) ? (_module.D5.create_DC14(_dafny.Seq.of((_module.D4.create_DC13(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(642))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), function (_0_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length))).length))).dtor_cf17))) : (_module.D5.create_DC14(_dafny.Seq.of(new BigNumber(675), new BigNumber(827), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(718),new BigNumber(-405))).length)))))).dtor_cf18;
    };
    static fm5(p0, globalState) {
      let _source0 = _module.D2.create_DC8(true, false);
      if (_source0.is_DC7) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ammm"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(832)), function (_1_i0) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }));
      } else if (_source0.is_DC8) {
        let _2___mcc_h0 = (_source0).cf9;
        let _3___mcc_h1 = (_source0).cf10;
        let _4_cf10 = _3___mcc_h1;
        let _5_cf9 = _2___mcc_h0;
        return _dafny.Seq.UnicodeFromString("widm");
      } else {
        let _6___mcc_h2 = (_source0).cf8;
        let _7_cf8 = _6___mcc_h2;
        if (true) {
          return _dafny.Seq.UnicodeFromString("ult");
        } else {
          return _dafny.Seq.UnicodeFromString("krthaq");
        }
      }
    };
    static fm8(globalState) {
      let _source1 = ((!(false)) ? (_module.D5.create_DC14(_dafny.Seq.of(new BigNumber(268)))) : (_module.D5.create_DC14(_dafny.Seq.of(new BigNumber(-267)))));
      if (_source1.is_DC15) {
        let _8___mcc_h0 = (_source1).cf19;
        let _9___mcc_h1 = (_source1).cf20;
        let _10_cf20 = _9___mcc_h1;
        let _11_cf19 = _8___mcc_h0;
        if (_10_cf20) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_10_cf20,_10_cf20)).length))));
        } else {
          return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_dafny.Seq.UnicodeFromString("isbarghd"))).length));
        }
      } else {
        let _12___mcc_h2 = (_source1).cf18;
        let _13_cf18 = _12___mcc_h2;
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(!(false))).length));
      }
    };
    static fm9(p0, p1, p2, globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_14_i0) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        })).Elements) {
          let _15_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), function (_14_i0) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _15_v0)) {
            _coll0.add(_15_v0);
          }
        }
        return _coll0;
      }()).Intersect(_dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(false, (new BigNumber(486)).isLessThan(new BigNumber(-386)), !((function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0)))).Elements) {
          let _16_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('v'.codePointAt(0))), _16_v0)) {
            _coll1.add(_16_v0);
          }
        }
        return _coll1;
      }()).IsSubsetOf(_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0))))), !(((false) ? (false) : (!(!(true))))));
    };
    static fm11(globalState) {
      if ((_dafny.Set.fromElements(true, !(false))).IsSubsetOf(_dafny.Set.fromElements(true, !(false)))) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),!(false));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, true),true));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-464),_dafny.Seq.UnicodeFromString("creunkjsj"))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),_dafny.Seq.UnicodeFromString("ft"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("lvvvveiw")).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(272)), function (_17_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }))));
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _18_v0;
      let _nw0 = Array((new BigNumber(26)).toNumber()).fill(false);
      _18_v0 = _nw0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_18_v0).length))) {
        let _19_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_19_i0)) && ((_19_i0).isLessThan(new BigNumber((_18_v0).length))))) {
          (_18_v0)[(_19_i0)] = false;
        }
      }
      let _20_v1;
      _20_v1 = new BigNumber(555);
      let _21_v2;
      _21_v2 = _module.D0.create_DC0(p0);
      let _22_v3;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_20_v1, p0, _21_v2);
      _22_v3 = _nw1;
      let _23_v4;
      _23_v4 = new _dafny.CodePoint('p'.codePointAt(0));
      let _24_v5;
      _24_v5 = _dafny.Set.fromElements(_23_v4);
      let _25_v6;
      _25_v6 = _dafny.Set.fromElements(_24_v5, _24_v5, _24_v5, _24_v5, _module.__default.fm9((_22_v3).f12, _20_v1, new BigNumber((p1).length), globalState));
      let _hi0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_22_v3,new BigNumber((_25_v6).length))).length);
      for (let _26_i1 = ((p0) ? (_20_v1) : (_20_v1)); _26_i1.isLessThan(_hi0); _26_i1 = _26_i1.plus(_dafny.ONE)) {
        let _27_v7;
        _27_v7 = _dafny.Seq.UnicodeFromString("lotbhuwwl");
        let _28_v8;
        _28_v8 = _dafny.Set.fromElements(new BigNumber((_27_v7).length));
        let _29_v9;
        _29_v9 = _dafny.Seq.of(p1);
        let _30_v10;
        _30_v10 = _dafny.Map.Empty.slice().updateUnsafe((_22_v3).f12,_26_i1);
        let _31_v11;
        _31_v11 = _dafny.MultiSet.fromElements(_18_v0);
        let _32_v12;
        _32_v12 = _31_v11;
        let _33_v13;
        let _nw2 = Array((new BigNumber(22)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = p1;
        _nw2[(_dafny.ONE).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(_26_i1, new BigNumber((p1).length)), !(!(p0)));
        _nw2[(new BigNumber(2)).toNumber()] = p1;
        _nw2[(new BigNumber(3)).toNumber()] = p1;
        _nw2[(new BigNumber(4)).toNumber()] = p1;
        _nw2[(new BigNumber(5)).toNumber()] = p1;
        _nw2[(new BigNumber(6)).toNumber()] = (_29_v9)[_module.__default.safeIndex(_20_v1, new BigNumber((_29_v9).length))];
        _nw2[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(p1, _dafny.Seq.update(_dafny.Seq.of(!(true)), _module.__default.safeIndex((_22_v3).f12, new BigNumber((_dafny.Seq.of(!(true))).length)), p0));
        _nw2[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(p1, p1);
        _nw2[(new BigNumber(9)).toNumber()] = _module.__default.fm10(new BigNumber(880), p0, (_22_v3).f12, _26_i1, globalState);
        _nw2[(new BigNumber(10)).toNumber()] = p1;
        _nw2[(new BigNumber(11)).toNumber()] = p1;
        _nw2[(new BigNumber(12)).toNumber()] = p1;
        _nw2[(new BigNumber(13)).toNumber()] = ((p0) ? (p1) : (_dafny.Seq.of(p0)));
        _nw2[(new BigNumber(14)).toNumber()] = _module.__default.fm10((_22_v3).f12, p0, (((_30_v10).contains(new BigNumber(((_32_v12)).cardinality()))) ? ((_30_v10).get(new BigNumber(((_32_v12)).cardinality()))) : (_26_i1)), _26_i1, globalState);
        _nw2[(new BigNumber(15)).toNumber()] = p1;
        _nw2[(new BigNumber(16)).toNumber()] = ((p0) ? (p1) : (_dafny.Seq.update(p1, _module.__default.safeIndex(_20_v1, new BigNumber((p1).length)), p0)));
        _nw2[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(p0, p0, p0, p0);
        _nw2[(new BigNumber(18)).toNumber()] = ((p0) ? (p1) : (p1));
        _nw2[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(p0, p0, p0, p0, !(p0));
        _nw2[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(_module.__default.fm2(p0, p1, globalState), p0);
        _nw2[(new BigNumber(21)).toNumber()] = p1;
        _33_v13 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length));
        (_33_v13)[_index0] = p1;
        let _index1 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length));
        let _rhs0 = _28_v8;
        let _rhs1 = p0;
        let _rhs2 = p1;
        let _rhs3 = _22_v3;
        let _rhs4 = p0;
        let _lhs0 = globalState;
        let _lhs1 = _33_v13;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length));
        let _lhs3 = globalState;
        _28_v8 = _rhs0;
        _lhs0.f5 = _rhs1;
        _lhs1[_lhs2] = _rhs2;
        _22_v3 = _rhs3;
        _lhs3.f5 = _rhs4;
        let _34_v14;
        _34_v14 = _module.D2.create_DC7();
        let _35_v15;
        _35_v15 = _dafny.Set.fromElements(_34_v14);
        if ((_26_i1).isLessThan(((_22_v3).f12).multipliedBy(new BigNumber((_35_v15).length)))) {
          let _36_v16;
          let _nw3 = Array((new BigNumber(21)).toNumber()).fill([]);
          _36_v16 = _nw3;
          _36_v16 = _36_v16;
          let _37_v17;
          let _nw4 = new _module.C0();
          _nw4.__ctor((new BigNumber((_27_v7).length)).multipliedBy(_26_i1), ((p0) ? (p0) : (p0)), _21_v2);
          _37_v17 = _nw4;
          let _pat_let_tv0 = p0;
          let _pat_let_tv1 = p0;
          let _pat_let_tv2 = p0;
          let _rhs5 = !(!(p0) || (true));
          let _rhs6 = (_26_i1).minus(_module.__default.safeDivisionInt((_37_v17).f12, (_dafny.ZERO).minus(_26_i1)));
          let _rhs7 = function (_pat_let0_0) {
            return function (_38_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_39_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_39_dt__update_hcf0_h0);
                }(_pat_let1_0);
              }(((!(_pat_let_tv2)) ? (_pat_let_tv0) : (_pat_let_tv1)));
            }(_pat_let0_0);
          }(_21_v2);
          let _rhs8 = !(!(p0) || (p0));
          let _lhs4 = globalState;
          let _lhs5 = globalState;
          _lhs4.f8 = _rhs5;
          _20_v1 = _rhs6;
          _21_v2 = _rhs7;
          _lhs5.f8 = _rhs8;
          let _40_v18;
          _40_v18 = _dafny.Seq.of((_22_v3).f12);
          let _41_v19;
          _41_v19 = _dafny.Map.Empty.slice().updateUnsafe(_40_v18,_37_v17);
          let _42_v20;
          _42_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_40_v18,_37_v17)).Merge(_41_v19),(_30_v10).Merge(_30_v10));
          _42_v20 = (_42_v20).update((_41_v19).update(_40_v18, _37_v17), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(485),(_22_v3).f12));
          let _43_v21;
          _43_v21 = _dafny.Seq.of(_37_v17, _22_v3, _37_v17, _22_v3);
          _43_v21 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_37_v17, _22_v3), _43_v21), _43_v21);
        } else {
          let _44_v22;
          let _nw5 = Array((new BigNumber(4)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = new BigNumber(-460);
          _nw5[(_dafny.ONE).toNumber()] = (_22_v3).f12;
          _nw5[(new BigNumber(2)).toNumber()] = _module.__default.safeModuloInt((_22_v3).f12, (_22_v3).f12);
          _nw5[(new BigNumber(3)).toNumber()] = (_22_v3).f12;
          _44_v22 = _nw5;
          _44_v22 = _44_v22;
          let _45_v23;
          _45_v23 = _dafny.MultiSet.fromElements(_module.__default.fm0((_dafny.ZERO).minus(_20_v1), _dafny.Map.Empty.slice().updateUnsafe(p1,p0), globalState));
          let _46_v24;
          _46_v24 = _module.D0.create_DC2(p0, new BigNumber((_45_v23).cardinality()));
          let _47_v25;
          _47_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_33_v13)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length))], _module.__default.safeIndex(new BigNumber(9), new BigNumber(((_33_v13)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length))]).length)), p0),p0);
          let _48_v26;
          _48_v26 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p1,p0), _47_v25, _47_v25);
          let _rhs9 = (_dafny.ZERO).minus((_22_v3).f12);
          let _rhs10 = (((_45_v23).contains((_22_v3).f12)) ? ((_45_v23).get((_22_v3).f12)) : (_module.__default.fm0((_22_v3).f12, (_48_v26)[_module.__default.safeIndex(_26_i1, new BigNumber((_48_v26).length))], globalState)));
          let _rhs11 = _23_v4;
          let _rhs12 = _46_v24;
          _20_v1 = _rhs9;
          _20_v1 = _rhs10;
          _23_v4 = _rhs11;
          _46_v24 = _rhs12;
          (globalState).f5 = true;
          let _49_v27;
          _49_v27 = _dafny.Seq.of(_module.__default.fm0(new BigNumber((_27_v7).length), _module.__default.fm11(globalState), globalState), new BigNumber(((_33_v13)[_module.__default.safeIndex(new BigNumber(907), new BigNumber((_33_v13).length))]).length));
          let _50_v28;
          _50_v28 = _dafny.Map.Empty.slice().updateUnsafe(_23_v4,(_22_v3).f12);
          (globalState).f5 = ((_module.__default.fm12(_30_v10, _28_v8, new BigNumber((_50_v28).length), p0, globalState)).update((_22_v3).f12, _27_v7)).contains((_49_v27)[_module.__default.safeIndex(new BigNumber(571), new BigNumber((_49_v27).length))]);
          let _51_v29;
          let _nw6 = new _module.C0();
          _nw6.__ctor(_20_v1, p0, _21_v2);
          _51_v29 = _nw6;
          let _52_v30;
          _52_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_51_v29);
          let _53_v31;
          _53_v31 = _dafny.Set.fromElements(_52_v30);
          _53_v31 = _53_v31;
        }
        let _54_v32;
        _54_v32 = _dafny.Set.fromElements(p0);
        let _55_v33;
        _55_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(459),_54_v32);
        let _56_v34;
        _56_v34 = _dafny.MultiSet.fromElements(new BigNumber(((_55_v33).Merge(_55_v33)).length), ((_22_v3).f12).plus(_20_v1), _20_v1, new BigNumber(865), (_26_i1).plus((_22_v3).f12));
        _56_v34 = _56_v34;
        _22_v3 = _22_v3;
      }
      let _57_v35;
      _57_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.MultiSet.fromElements(p0));
      let _58_v36;
      _58_v36 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      (globalState).f8 = (_22_v3).fm7(_21_v2, _57_v35, (((((_58_v36).contains(p0)) ? ((_58_v36).get(p0)) : (p0))) ? (_20_v1) : (new BigNumber(609))), globalState);
      let _59_v37;
      let _nw7 = new _module.C0();
      _nw7.__ctor((_dafny.ZERO).minus(new BigNumber(-566)), p0, _21_v2);
      _59_v37 = _nw7;
      let _60_v38;
      _60_v38 = _dafny.Map.Empty.slice().updateUnsafe((_22_v3).f12,_59_v37);
      let _61_v39;
      _61_v39 = _dafny.MultiSet.fromElements(_59_v37, (((_60_v38).contains((_22_v3).f12)) ? ((_60_v38).get((_22_v3).f12)) : (_59_v37)));
      let _62_v40;
      _62_v40 = _dafny.Map.Empty.slice().updateUnsafe(_61_v39,!(!(_20_v1).isEqualTo(new BigNumber(922))));
      (globalState).f8 = (((_62_v40).contains(_61_v39)) ? ((_62_v40).get(_61_v39)) : ((_59_v37).f10));
      let _63_i2;
      _63_i2 = _dafny.ZERO;
      L0: {
        while ((_module.D2.create_DC8((_59_v37).f10, p0)).dtor_cf10) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_63_i2)) {
              break L0;
            }
            _63_i2 = (_63_i2).plus(_dafny.ONE);
            let _rhs13 = new BigNumber(107);
            let _rhs14 = _20_v1;
            let _rhs15 = false;
            let _lhs6 = globalState;
            _20_v1 = _rhs13;
            _20_v1 = _rhs14;
            _lhs6.f8 = _rhs15;
            if (((_22_v3).f12).isEqualTo((_22_v3).f12)) {
              _20_v1 = (_dafny.ZERO).minus(_20_v1);
              let _64_v41;
              _64_v41 = _dafny.Seq.of((_22_v3).f12, _20_v1);
              let _65_v42;
              _65_v42 = _module.D2.create_DC8((_59_v37).f10, (_59_v37).f10);
              let _66_v43;
              _66_v43 = _dafny.MultiSet.fromElements((_59_v37).f10);
              let _67_v44;
              _67_v44 = _dafny.Seq.of(_66_v43);
              let _68_v45;
              let _init0 = ((_69_v1) => function (_70_i3) {
                return (_70_i3).multipliedBy(_69_v1);
              })(_20_v1);
              let _nw8 = Array((new BigNumber(18)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw8.length); _i0_0++) {
                _nw8[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _68_v45 = _nw8;
              let _71_v46;
              _71_v46 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex((_22_v3).f12, new BigNumber((p1).length))],_68_v45);
              let _rhs16 = _dafny.Seq.Concat(_64_v41, (((_59_v37).f10) ? (_64_v41) : (_dafny.Seq.of((_22_v3).f12))));
              let _rhs17 = !(!_dafny.Seq.contains(_67_v44, _66_v43)) || (p0);
              let _rhs18 = (new BigNumber(854)).multipliedBy(((_22_v3).f12).multipliedBy(new BigNumber((_71_v46).length)));
              let _rhs19 = _65_v42;
              let _rhs20 = (((_60_v38).contains((_22_v3).f12)) ? ((_60_v38).get((_22_v3).f12)) : (_59_v37));
              let _lhs7 = globalState;
              _64_v41 = _rhs16;
              _lhs7.f8 = _rhs17;
              _20_v1 = _rhs18;
              _65_v42 = _rhs19;
              _59_v37 = _rhs20;
              let _index2 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_68_v45).length));
              (_68_v45)[_index2] = (_22_v3).f12;
              let _72_v47;
              _72_v47 = _dafny.MultiSet.fromElements(_18_v0);
              let _73_v48;
              _73_v48 = _dafny.Set.fromElements(_64_v41);
              let _74_v49;
              _74_v49 = _dafny.Seq.of(_22_v3);
              let _75_v50;
              _75_v50 = _dafny.Map.Empty.slice().updateUnsafe((_73_v48).IsSubsetOf(_73_v48),_dafny.Seq.Concat(_74_v49, _74_v49));
              let _index3 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_68_v45).length));
              let _rhs21 = (new BigNumber(((_72_v47).Union(_72_v47)).cardinality())).multipliedBy((_22_v3).f12);
              let _rhs22 = (_22_v3).f12;
              let _rhs23 = (_22_v3).f12;
              let _rhs24 = new BigNumber((_75_v50).length);
              let _rhs25 = ((_59_v37).f10) || ((_59_v37).f10);
              let _lhs8 = _68_v45;
              let _lhs9 = _module.__default.safeIndex(new BigNumber(287), new BigNumber((_68_v45).length));
              let _lhs10 = globalState;
              _lhs8[_lhs9] = _rhs21;
              _20_v1 = _rhs22;
              _20_v1 = _rhs23;
              _20_v1 = _rhs24;
              _lhs10.f8 = _rhs25;
              _59_v37 = _59_v37;
              let _76_v51;
              let _init1 = ((_77_v37, _78_p0, _79_v3) => function (_80_i4) {
                return _module.D1.create_DC4((_77_v37).f10, _dafny.Map.Empty.slice().updateUnsafe(_78_p0,(_79_v3).f12), !((_77_v37).f10));
              })(_59_v37, p0, _22_v3);
              let _nw9 = Array((_dafny.ONE).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw9.length); _i0_1++) {
                _nw9[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _76_v51 = _nw9;
              let _81_v52;
              _81_v52 = _dafny.Map.Empty.slice().updateUnsafe((_59_v37).f10,(_22_v3).f12);
              let _82_v53;
              _82_v53 = _module.D1.create_DC4((_59_v37).f10, _81_v52, true);
              let _index4 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_76_v51).length));
              (_76_v51)[_index4] = _82_v53;
              let _index5 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_76_v51).length));
              (_76_v51)[_index5] = _82_v53;
            } else {
              let _83_v54;
              _83_v54 = _module.D0.create_DC1();
              _83_v54 = _83_v54;
              let _84_v55;
              _84_v55 = _module.D2.create_DC8((_59_v37).f10, (_59_v37).f10);
              let _85_v56;
              _85_v56 = _dafny.MultiSet.fromElements((_84_v55).dtor_cf9, (_59_v37).f10, (_59_v37).f10);
              (globalState).f8 = (((_59_v37).f10) ? ((_59_v37).f10) : ((_22_v3).fm7(_module.D0.create_DC0(false), _dafny.Map.Empty.slice().updateUnsafe((_59_v37).f10,_85_v56), (_22_v3).f12, globalState)));
              _59_v37 = _59_v37;
              _59_v37 = _59_v37;
              let _86_v57;
              let _nw10 = Array((new BigNumber(9)).toNumber());
              _86_v57 = _nw10;
              let _87_v58;
              _87_v58 = _dafny.Map.Empty.slice().updateUnsafe((_22_v3).f12,_22_v3);
              let _index6 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_86_v57).length));
              (_86_v57)[_index6] = (((_59_v37).f10) ? (_22_v3) : ((((_87_v58).contains((_22_v3).f12)) ? ((_87_v58).get((_22_v3).f12)) : (_22_v3))));
              let _index7 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_86_v57).length));
              (_86_v57)[_index7] = _22_v3;
            }
            let _88_v59;
            _88_v59 = _module.D0.create_DC2((_59_v37).f10, (_22_v3).f12);
            let _89_v60;
            _89_v60 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_59_v37).f10);
            let _90_v61;
            _90_v61 = _module.D4.create_DC12((_59_v37).f10, _59_v37, true, (_59_v37).f10);
            let _91_v62;
            _91_v62 = _dafny.Map.Empty.slice().updateUnsafe(_22_v3,_90_v61);
            let _92_v63;
            _92_v63 = _dafny.Seq.of(new BigNumber(-463));
            let _93_v64;
            let _nw11 = Array((new BigNumber(19)).toNumber());
            _nw11[(_dafny.ZERO).toNumber()] = new BigNumber(-181);
            _nw11[(_dafny.ONE).toNumber()] = (_22_v3).f12;
            _nw11[(new BigNumber(2)).toNumber()] = new BigNumber(658);
            _nw11[(new BigNumber(3)).toNumber()] = (_88_v59).dtor_cf2;
            _nw11[(new BigNumber(4)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(5)).toNumber()] = (_22_v3).f12;
            _nw11[(new BigNumber(6)).toNumber()] = _module.__default.fm0((_22_v3).f12, _89_v60, globalState);
            _nw11[(new BigNumber(7)).toNumber()] = (_22_v3).f12;
            _nw11[(new BigNumber(8)).toNumber()] = new BigNumber((_91_v62).length);
            _nw11[(new BigNumber(9)).toNumber()] = new BigNumber(-413);
            _nw11[(new BigNumber(10)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(11)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(12)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(13)).toNumber()] = (_22_v3).f12;
            _nw11[(new BigNumber(14)).toNumber()] = new BigNumber(528);
            _nw11[(new BigNumber(15)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(16)).toNumber()] = (_22_v3).f12;
            _nw11[(new BigNumber(17)).toNumber()] = _20_v1;
            _nw11[(new BigNumber(18)).toNumber()] = new BigNumber((_92_v63).length);
            _93_v64 = _nw11;
            let _94_v65;
            _94_v65 = _module.D4.create_DC10(_93_v64);
            let _95_v66;
            _95_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(17),(_94_v65).dtor_cf12);
            _95_v66 = (_95_v66).update(((p0) ? (_20_v1) : (new BigNumber(-648))), _93_v64);
            (globalState).f8 = ((_59_v37).f10) === ((_module.__default.fm2((_59_v37).f10, p1, globalState)) && ((_59_v37).f10));
          }
        }
      }
      let _nw12 = new _module.C0();
      _nw12.__ctor((_22_v3).f12, p0, (_59_v37).f11);
      _59_v37 = _nw12;
      let _96_v67;
      _96_v67 = _dafny.Seq.UnicodeFromString("qf");
      r0 = _dafny.Seq.Concat(_96_v67, _96_v67);
      return r0;
    }
    static Main(__noArgsParameter) {
      let _97_v0;
      _97_v0 = false;
      let _98_v1;
      _98_v1 = _dafny.Seq.of(_97_v0, _97_v0);
      let _99_v2;
      _99_v2 = new BigNumber(147);
      let _100_v3;
      _100_v3 = _dafny.MultiSet.fromElements(_dafny.Seq.of(true), _98_v1, _dafny.Seq.update(_98_v1, _module.__default.safeIndex(new BigNumber(-758), new BigNumber((_98_v1).length)), false), _dafny.Seq.update(_98_v1, _module.__default.safeIndex(_99_v2, new BigNumber((_98_v1).length)), !(_97_v0)));
      let _101_v4;
      _101_v4 = _module.D0.create_DC0(_97_v0);
      let _102_v5;
      let _nw13 = Array((new BigNumber(11)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _97_v0;
      _nw13[(_dafny.ONE).toNumber()] = (_101_v4).dtor_cf0;
      _nw13[(new BigNumber(2)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(3)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(4)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(5)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(6)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(7)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(8)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(9)).toNumber()] = _97_v0;
      _nw13[(new BigNumber(10)).toNumber()] = _97_v0;
      _102_v5 = _nw13;
      let _103_v6;
      _103_v6 = _dafny.Map.Empty.slice().updateUnsafe(_97_v0,_99_v2);
      let _104_v7;
      _104_v7 = _module.D1.create_DC4(_97_v0, _103_v6, _97_v0);
      let _105_v8;
      _105_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(_97_v0))).length),_97_v0);
      let _106_globalState;
      let _nw14 = new _module.GlobalState();
      _nw14.__ctor(_100_v3, _102_v5, new BigNumber(-299), true, new BigNumber(-378), true, (_104_v7).dtor_cf5, _105_v8, false, true);
      _106_globalState = _nw14;
      let _107_v9;
      _107_v9 = _dafny.MultiSet.fromElements(!(_97_v0), _97_v0, _97_v0);
      if (!(((_107_v9).Intersect(_107_v9)).IsSubsetOf(_107_v9))) {
        let _108_v10;
        _108_v10 = _dafny.Map.Empty.slice().updateUnsafe(_98_v1,_97_v0);
        let _109_v11;
        _109_v11 = _dafny.Seq.UnicodeFromString("dsjek");
        _99_v2 = (_module.__default.fm0(_99_v2, _108_v10, _106_globalState)).plus((new BigNumber((_109_v11).length)).plus(_99_v2));
        let _110_v12;
        let _init2 = ((_111_v2) => function (_112_i0) {
          return _module.__default.safeModuloInt(_112_i0, _111_v2);
        })(_99_v2);
        let _nw15 = Array((new BigNumber(28)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw15.length); _i0_2++) {
          _nw15[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _110_v12 = _nw15;
        _110_v12 = _110_v12;
        let _113_v13;
        let _out0;
        _out0 = _module.__default.m0(_97_v0, _98_v1, _106_globalState);
        _113_v13 = _out0;
        let _index8 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_110_v12).length));
        (_110_v12)[_index8] = _99_v2;
        let _114_v14;
        _114_v14 = _dafny.Set.fromElements(_97_v0, _97_v0, false, _97_v0, _97_v0);
        let _115_v15;
        _115_v15 = _dafny.Set.fromElements(new BigNumber((_114_v14).length));
        let _index9 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_110_v12).length));
        (_110_v12)[_index9] = (new BigNumber((_115_v15).length)).plus(_99_v2);
        let _116_v16;
        _116_v16 = _dafny.Map.Empty.slice().updateUnsafe(_97_v0,_97_v0);
        let _117_v17;
        _117_v17 = _dafny.Seq.of(((!((((_116_v16).contains(true)) ? ((_116_v16).get(true)) : (_97_v0)))) ? (new BigNumber(446)) : ((_110_v12)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_110_v12).length))])));
        let _118_v18;
        _118_v18 = _dafny.Map.Empty.slice().updateUnsafe(_97_v0,_117_v17);
        _117_v17 = _dafny.Seq.Concat(_dafny.Seq.Concat(_117_v17, _dafny.Seq.update((((_118_v18).contains(_97_v0)) ? ((_118_v18).get(_97_v0)) : (_117_v17)), _module.__default.safeIndex(_99_v2, new BigNumber(((((_118_v18).contains(_97_v0)) ? ((_118_v18).get(_97_v0)) : (_117_v17))).length)), (_110_v12)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_110_v12).length))])), _dafny.Seq.of((_110_v12)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_110_v12).length))], new BigNumber(584)));
      } else {
        _107_v9 = (_107_v9).Union(_module.__default.fm1(_99_v2, _97_v0, _99_v2, _97_v0, _106_globalState));
        (_106_globalState).f8 = (_module.D1.create_DC4(_97_v0, _103_v6, _97_v0)).dtor_cf6;
        let _119_v19;
        _119_v19 = _module.D1.create_DC3(_103_v6);
        let _120_v20;
        _120_v20 = _module.D1.create_DC5(_119_v19);
        let _index10 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        (_102_v5)[_index10] = false;
        let _index11 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        let _rhs26 = _99_v2;
        let _rhs27 = _105_v8;
        let _rhs28 = _120_v20;
        let _rhs29 = _99_v2;
        let _rhs30 = _97_v0;
        let _lhs11 = _102_v5;
        let _lhs12 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        _99_v2 = _rhs26;
        _105_v8 = _rhs27;
        _120_v20 = _rhs28;
        _99_v2 = _rhs29;
        _lhs11[_lhs12] = _rhs30;
        let _index12 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        (_102_v5)[_index12] = _module.__default.fm2((_102_v5)[_module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length))], _dafny.Seq.Concat(_98_v1, _98_v1), _106_globalState);
        let _index13 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        let _rhs31 = !(_97_v0) || (_97_v0);
        let _rhs32 = (_102_v5)[_module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length))];
        let _lhs13 = _106_globalState;
        let _lhs14 = _102_v5;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_102_v5).length));
        _lhs13.f8 = _rhs31;
        _lhs14[_lhs15] = _rhs32;
      }
      (_106_globalState).f8 = _97_v0;
      _99_v2 = _99_v2;
      let _hi1 = _99_v2;
      for (let _121_i1 = _99_v2; _121_i1.isLessThan(_hi1); _121_i1 = _121_i1.plus(_dafny.ONE)) {
        let _122_v21;
        _122_v21 = _dafny.Set.fromElements((_dafny.ZERO).minus(_99_v2), _121_i1, _121_i1);
        let _123_v22;
        _123_v22 = _dafny.Seq.UnicodeFromString("wgy");
        let _124_v23;
        _124_v23 = _module.D0.create_DC1();
        let _125_v24;
        _125_v24 = _dafny.Map.Empty.slice().updateUnsafe(_123_v22,_124_v23);
        _99_v2 = ((!(_107_v9).equals(_107_v9)) ? (new BigNumber((_122_v21).length)) : ((_121_i1).multipliedBy(new BigNumber((_125_v24).length))));
        let _126_v25;
        _126_v25 = _dafny.Map.Empty.slice().updateUnsafe(_121_i1,_99_v2);
        _107_v9 = (_dafny.MultiSet.fromElements(_97_v0, _97_v0, _97_v0, _97_v0)).update(((((_126_v25).contains(_99_v2)) ? ((_126_v25).get(_99_v2)) : (new BigNumber((_123_v22).length)))).isLessThanOrEqualTo(_121_i1), _module.__default.abs(_121_i1));
        (_106_globalState).f8 = _97_v0;
        if (!(!((_97_v0) === (!(_97_v0)))) || (_97_v0)) {
          let _127_v26;
          _127_v26 = new _dafny.CodePoint('f'.codePointAt(0));
          let _128_v27;
          let _nw16 = Array((new BigNumber(3)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _127_v26;
          _nw16[(_dafny.ONE).toNumber()] = ((!(_97_v0)) ? (new _dafny.CodePoint('y'.codePointAt(0))) : (_127_v26));
          _nw16[(new BigNumber(2)).toNumber()] = _127_v26;
          _128_v27 = _nw16;
          _128_v27 = _128_v27;
          (_106_globalState).f8 = _97_v0;
          _98_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_98_v1, _98_v1), _98_v1);
          let _129_v28;
          let _out1;
          _out1 = _module.__default.m0(_97_v0, _dafny.Seq.Concat(_98_v1, _dafny.Seq.of(!(_97_v0))), _106_globalState);
          _129_v28 = _out1;
          let _index14 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_102_v5).length));
          (_102_v5)[_index14] = true;
          let _130_v29;
          let _init3 = ((_131_v2) => function (_132_i2) {
            return _module.__default.safeModuloInt(_132_i2, _131_v2);
          })(_99_v2);
          let _nw17 = Array((new BigNumber(2)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw17.length); _i0_3++) {
            _nw17[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _130_v29 = _nw17;
          let _index15 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_130_v29).length));
          (_130_v29)[_index15] = (_dafny.ZERO).minus(_99_v2);
          let _133_v30;
          _133_v30 = _module.D2.create_DC6(_module.__default.fm3(_106_globalState));
          let _134_v31;
          _134_v31 = _dafny.Set.fromElements(_97_v0);
          let _135_v32;
          _135_v32 = _dafny.Seq.of(_99_v2, new BigNumber(-223));
          let _index16 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_102_v5).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_130_v29).length));
          let _rhs33 = (((_133_v30).dtor_cf8).Difference(_134_v31)).IsDisjointFrom(((_97_v0) ? (_dafny.Set.fromElements(_97_v0, _97_v0)) : (_134_v31)));
          let _rhs34 = _module.__default.safeDivisionInt(_99_v2, _99_v2);
          let _rhs35 = (_135_v32)[_module.__default.safeIndex((_121_i1).multipliedBy(_121_i1), new BigNumber((_135_v32).length))];
          let _rhs36 = (_dafny.ZERO).minus((_121_i1).plus(new BigNumber(325)));
          let _lhs16 = _102_v5;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_102_v5).length));
          let _lhs18 = _130_v29;
          let _lhs19 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_130_v29).length));
          _lhs16[_lhs17] = _rhs33;
          _99_v2 = _rhs34;
          _99_v2 = _rhs35;
          _lhs18[_lhs19] = _rhs36;
        } else {
          _123_v22 = _123_v22;
          _99_v2 = _121_i1;
          let _136_v33;
          _136_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_97_v0);
          (_106_globalState).f5 = _module.__default.fm2(!(_module.__default.fm0(_121_i1, _136_v33, _106_globalState)).isEqualTo(_99_v2), _98_v1, _106_globalState);
          _101_v4 = _101_v4;
          let _137_v34;
          _137_v34 = new _dafny.CodePoint('i'.codePointAt(0));
          _137_v34 = _137_v34;
        }
      }
      let _138_v36;
      _138_v36 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_99_v2, function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_98_v1,_97_v0)).Keys.Elements) {
          let _139_v35 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_98_v1,_97_v0)).contains(_139_v35)) {
            _coll2.push([_139_v35,_97_v0]);
          }
        }
        return _coll2;
      }(), _106_globalState),_99_v2);
      let _140_v37;
      _140_v37 = _dafny.MultiSet.fromElements(new BigNumber(717));
      let _141_v38;
      _141_v38 = _dafny.Seq.UnicodeFromString("toc");
      let _142_v39;
      _142_v39 = new _dafny.CodePoint('y'.codePointAt(0));
      let _143_v40;
      let _init4 = ((_144_v2) => function (_145_i3) {
        return _module.__default.safeDivisionInt(_145_i3, _144_v2);
      })(_99_v2);
      let _nw18 = Array((new BigNumber(11)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw18.length); _i0_4++) {
        _nw18[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _143_v40 = _nw18;
      let _146_v41;
      _146_v41 = _dafny.Seq.of(_143_v40);
      let _rhs37 = (_138_v36).equals((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(249),(((_140_v37).contains(_99_v2)) ? ((_140_v37).get(_99_v2)) : (new BigNumber((_141_v38).length))))).Merge(_138_v36));
      let _rhs38 = new BigNumber((_module.__default.fm4(_142_v39, _dafny.Seq.Concat(_module.__default.fm5(_142_v39, _106_globalState), _141_v38), _97_v0, _106_globalState)).length);
      let _rhs39 = _97_v0;
      let _rhs40 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_143_v40), _146_v41), _dafny.Seq.Concat(_146_v41, _dafny.Seq.of(_143_v40, _143_v40, _143_v40, _143_v40)))).length));
      let _lhs20 = _106_globalState;
      _97_v0 = _rhs37;
      _99_v2 = _rhs38;
      _lhs20.f8 = _rhs39;
      _99_v2 = _rhs40;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_102_v5).length))) {
        let _147_i4 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_147_i4)) && ((_147_i4).isLessThan(new BigNumber((_102_v5).length))))) {
          (_102_v5)[(_147_i4)] = ((_99_v2).multipliedBy(_99_v2)).isLessThan(_99_v2);
        }
      }
      _141_v38 = _141_v38;
      _99_v2 = _99_v2;
      let _148_v43;
      _148_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_99_v2,_99_v2)).length),function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(728), new BigNumber(918))) {
          let _149_v42 = _compr_3;
          if (((new BigNumber(728)).isLessThanOrEqualTo(_149_v42)) && ((_149_v42).isLessThan(new BigNumber(918)))) {
            _coll3.push([_module.__default.safeDivisionInt(_149_v42, _99_v2),_97_v0]);
          }
        }
        return _coll3;
      }());
      let _index18 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
      (_102_v5)[_index18] = !(_148_v43).contains(new BigNumber(927));
      let _150_v45;
      _150_v45 = _dafny.Set.fromElements(_98_v1);
      let _index19 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
      let _rhs41 = _97_v0;
      let _rhs42 = ((true) ? (((_97_v0) ? (_142_v39) : (_142_v39))) : (_142_v39));
      let _rhs43 = _module.__default.safeModuloInt(_99_v2, _module.__default.safeModuloInt(_99_v2, _module.__default.fm0(_99_v2, function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_150_v45).Elements) {
          let _151_v44 = _compr_4;
          if ((_150_v45).contains(_151_v44)) {
            _coll4.push([_151_v44,_97_v0]);
          }
        }
        return _coll4;
      }(), _106_globalState)));
      let _rhs44 = (new BigNumber(601)).plus(_99_v2);
      let _lhs21 = _102_v5;
      let _lhs22 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
      _lhs21[_lhs22] = _rhs41;
      _142_v39 = _rhs42;
      _99_v2 = _rhs43;
      _99_v2 = _rhs44;
      let _152_v46;
      _152_v46 = _dafny.Map.Empty.slice().updateUnsafe((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))],(_97_v0) === ((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]));
      let _153_v48;
      _153_v48 = _dafny.Seq.of(_99_v2);
      _152_v46 = (_152_v46).update((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_153_v48).Elements) {
          let _154_v47 = _compr_5;
          if (_dafny.Seq.contains(_153_v48, _154_v47)) {
            _coll5.push([(_154_v47).minus(_99_v2),_140_v37]);
          }
        }
        return _coll5;
      }()).contains(new BigNumber((_140_v37).cardinality())), (_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]);
      let _155_v49;
      let _nw19 = new _module.C0();
      _nw19.__ctor(_99_v2, _97_v0, _101_v4);
      _155_v49 = _nw19;
      _99_v2 = ((_155_v49).f12).multipliedBy(((_155_v49).f12).minus((_153_v48)[_module.__default.safeIndex(new BigNumber(130), new BigNumber((_153_v48).length))]));
      let _hi2 = _99_v2;
      for (let _156_i5 = (new BigNumber((_105_v8).length)).plus((_155_v49).f12); _156_i5.isLessThan(_hi2); _156_i5 = _156_i5.plus(_dafny.ONE)) {
        (_106_globalState).f8 = _97_v0;
        let _157_v50;
        let _nw20 = new _module.C0();
        _nw20.__ctor((_dafny.ZERO).minus(((_155_v49).f12).multipliedBy(new BigNumber((_103_v6).length))), (((_105_v8).contains((_155_v49).f12)) ? ((_105_v8).get((_155_v49).f12)) : (true)), _101_v4);
        _157_v50 = _nw20;
        let _index20 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
        (_102_v5)[_index20] = !(((_module.__default.fm8(_106_globalState)).Intersect(_dafny.MultiSet.FromArray(_153_v48))).IsProperSubsetOf(_140_v37));
        let _158_v51;
        _158_v51 = _dafny.Seq.of(_157_v50, _157_v50, _157_v50);
        if ((new BigNumber((_158_v51).length)).isEqualTo(_99_v2)) {
          let _159_v52;
          let _out2;
          _out2 = _module.__default.m0((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))], _98_v1, _106_globalState);
          _159_v52 = _out2;
          let _160_v53;
          let _out3;
          _out3 = _module.__default.m0((_155_v49).fm6(_97_v0, _99_v2, true, _106_globalState), _98_v1, _106_globalState);
          _160_v53 = _out3;
          (_106_globalState).f8 = _97_v0;
          let _rhs45 = _97_v0;
          let _rhs46 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_155_v49).f12), ((_155_v49).f12).minus(new BigNumber(-61)));
          let _lhs23 = _106_globalState;
          _lhs23.f5 = _rhs45;
          _99_v2 = _rhs46;
          let _index21 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
          (_102_v5)[_index21] = _97_v0;
        } else {
          _99_v2 = (_157_v50).f12;
          _99_v2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_99_v2).minus(new BigNumber(458)), (_dafny.ZERO).minus((_155_v49).f12)));
          _142_v39 = _142_v39;
          let _161_v54;
          _161_v54 = _module.D2.create_DC6(_dafny.Set.fromElements((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]));
          let _162_v55;
          _162_v55 = _dafny.Map.Empty.slice().updateUnsafe(_161_v54,(_157_v50).f12);
          _99_v2 = (((_162_v55).contains(_161_v54)) ? ((_162_v55).get(_161_v54)) : (((_155_v49).f12).plus((_155_v49).f12)));
          let _163_v56;
          let _out4;
          _out4 = _module.__default.m0((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))], _98_v1, _106_globalState);
          _163_v56 = _out4;
        }
      }
      let _pat_let_tv3 = _102_v5;
      let _pat_let_tv4 = _102_v5;
      let _index22 = _module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length));
      (_102_v5)[_index22] = (function (_pat_let2_0) {
        return function (_164_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_165_dt__update_hcf10_h0) {
              return _module.D2.create_DC8((_164_dt__update__tmp_h0).dtor_cf9, _165_dt__update_hcf10_h0);
            }(_pat_let3_0);
          }((_pat_let_tv4)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_pat_let_tv3).length))]);
        }(_pat_let2_0);
      }(_module.D2.create_DC8(_97_v0, (_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]))).dtor_cf9;
      let _166_v57;
      let _nw21 = new _module.C0();
      _nw21.__ctor((_155_v49).f12, _97_v0, _module.D0.create_DC0((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]));
      _166_v57 = _nw21;
      let _167_v58;
      let _out5;
      _out5 = _module.__default.m0((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))], _dafny.Seq.of((_102_v5)[_module.__default.safeIndex(new BigNumber(797), new BigNumber((_102_v5).length))]), _106_globalState);
      _167_v58 = _out5;
      process.stdout.write(_dafny.toString(_97_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_98_v1, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_99_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_100_v3).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(false, false), _dafny.Seq.of(false, false), _dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_101_v4).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_102_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(147)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v7).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_104_v7).dtor_cf5).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(147)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_104_v7).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_105_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f0).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(false, false), _dafny.Seq.of(false, false), _dafny.Seq.of(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_106_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f6).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(147)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_106_globalState).f7).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_106_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_107_v9).equals(_dafny.MultiSet.fromElements(true, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v36).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),new BigNumber(147)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_140_v37).equals(_dafny.MultiSet.fromElements(new BigNumber(717)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_141_v38).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_142_v39));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_143_v40)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_146_v41).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_v43).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-104),false).updateUnsafe(new BigNumber(-105),false).updateUnsafe(new BigNumber(-106),false).updateUnsafe(new BigNumber(-107),false).updateUnsafe(new BigNumber(-108),false).updateUnsafe(new BigNumber(-109),false).updateUnsafe(new BigNumber(-110),false).updateUnsafe(new BigNumber(-111),false).updateUnsafe(new BigNumber(-112),false).updateUnsafe(new BigNumber(-113),false).updateUnsafe(new BigNumber(-114),false).updateUnsafe(new BigNumber(-115),false).updateUnsafe(new BigNumber(-116),false).updateUnsafe(new BigNumber(-117),false).updateUnsafe(new BigNumber(-118),false).updateUnsafe(new BigNumber(-119),false).updateUnsafe(new BigNumber(-120),false).updateUnsafe(new BigNumber(-121),false).updateUnsafe(new BigNumber(-122),false).updateUnsafe(new BigNumber(-123),false).updateUnsafe(new BigNumber(-124),false).updateUnsafe(new BigNumber(-125),false).updateUnsafe(new BigNumber(-126),false).updateUnsafe(new BigNumber(-127),false).updateUnsafe(new BigNumber(-128),false).updateUnsafe(new BigNumber(-129),false).updateUnsafe(new BigNumber(-130),false).updateUnsafe(new BigNumber(-131),false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_150_v45).equals(_dafny.Set.fromElements(_dafny.Seq.of(false, false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v46).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_153_v48, _dafny.Seq.of(new BigNumber(594)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v49).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_166_v57).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_167_v58).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC2(cf1, cf2) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
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
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf1 === other.cf1 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC3(cf3) {
      let $dt = new D1(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC5(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(false, _dafny.Map.Empty, false);
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
    static create_DC7() {
      let $dt = new D2(0);
      return $dt;
    }
    static create_DC8(cf9, cf10) {
      let $dt = new D2(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D2(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7";
      } else if (this.$tag === 1) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf8) + ")";
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
        return other.$tag === 1 && this.cf9 === other.cf9 && this.cf10 === other.cf10;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7();
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
    static create_DC9(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC11() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC12(cf13, cf14, cf15, cf16) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC13(cf17) {
      let $dt = new D4(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC10(cf12) {
      let $dt = new D4(3);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf12) + ")";
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
        return other.$tag === 1 && this.cf13 === other.cf13 && this.cf14 === other.cf14 && this.cf15 === other.cf15 && this.cf16 === other.cf16;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf12 === other.cf12;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11();
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
    static create_DC15(cf19, cf20) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC14(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f5 = false;
      this.f8 = false;
      this._f0 = _dafny.MultiSet.Empty;
      this._f1 = [];
      this._f2 = _dafny.ZERO;
      this._f3 = false;
      this._f4 = _dafny.ZERO;
      this._f6 = _dafny.Map.Empty;
      this._f7 = _dafny.Map.Empty;
      this._f9 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
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
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f10 = false;
      this._f11 = _module.D0.Default();
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    __ctor(f12, f10, f11) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      return;
    }
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((_this).f12)).isLessThan((_this).f12);
    };
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return !(false) || (false);
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
